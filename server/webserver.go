package server

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
	"github.com/livepeer/go-livepeer/core"
	"github.com/livepeer/go-livepeer/eth"
	"github.com/livepeer/go-livepeer/net"
	"github.com/livepeer/go-livepeer/types"
)

func (s *LivepeerServer) StartWebserver() {
	//Temporary endpoint just so we can invoke a transcode job.  IRL this should be invoked by transcoders monitoring the smart contract.
	http.HandleFunc("/transcode", func(w http.ResponseWriter, r *http.Request) {
		strmID := r.URL.Query().Get("strmID")
		if strmID == "" {
			http.Error(w, "Need to specify strmID", 500)
			return
		}

		ps := []types.VideoProfile{types.P240p30fps16x9, types.P360p30fps16x9}
		ids, err := s.LivepeerNode.TranscodeAndBroadcast(net.TranscodeConfig{StrmID: strmID, Profiles: ps}, nil)
		if err != nil {
			glog.Errorf("Error transcoding: %v", err)
			http.Error(w, "Error transcoding.", 500)
		}

		vids := make(map[core.StreamID]types.VideoProfile)
		for i, vp := range ps {
			vids[ids[i]] = vp
		}

		sid := core.StreamID(strmID)
		s.LivepeerNode.NotifyBroadcaster(sid.GetNodeID(), sid, vids)
	})

	//Set the broadcast config for creating onchain jobs.
	http.HandleFunc("/setBroadcastConfig", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Query().Get("price")
		if p != "" {
			pi, err := strconv.Atoi(p)
			if err != nil {
				glog.Errorf("Price conversion failed: %v", err)
				return
			}
			BroadcastPrice = big.NewInt(int64(pi))
		}

		j := r.URL.Query().Get("job")
		if j != "" {
			jp := types.VideoProfileLookup[j]
			if jp.Name != "" {
				BroadcastJobVideoProfile = jp
			}
		}

		glog.Infof("Transcode Job Price: %v, Transcode Job Type: %v", BroadcastPrice, BroadcastJobVideoProfile.Name)
	})

	//Activate the transcoder on-chain.
	http.HandleFunc("/activateTranscoder", func(w http.ResponseWriter, r *http.Request) {
		active, err := s.LivepeerNode.Eth.IsActiveTranscoder()
		if err != nil {
			glog.Errorf("Error getting transcoder state: %v", err)
		}
		if active {
			glog.Error("Transcoder is already active")
			return
		}

		//Wait until a fresh round, register transcoder
		err = s.LivepeerNode.Eth.WaitUntilNextRound(eth.ProtocolBlockPerRound)
		if err != nil {
			glog.Errorf("Failed to wait until next round: %v", err)
			return
		}
		if err := eth.CheckRoundAndInit(s.LivepeerNode.Eth, EthRpcTimeout, EthMinedTxTimeout); err != nil {
			glog.Errorf("%v", err)
			return
		}

		tx, err := s.LivepeerNode.Eth.Transcoder(10, 5, big.NewInt(100))
		if err != nil {
			glog.Errorf("Error creating transcoder: %v", err)
			return
		}

		receipt, err := eth.WaitForMinedTx(s.LivepeerNode.Eth.Backend(), EthRpcTimeout, EthMinedTxTimeout, tx.Hash())
		if err != nil {
			glog.Errorf("%v", err)
			return
		}
		if tx.Gas().Cmp(receipt.GasUsed) == 0 {
			glog.Errorf("Client 0 failed transcoder registration")
		}

		printStake(s.LivepeerNode.Eth)
	})

	//Set transcoder config on-chain.
	http.HandleFunc("/setTranscoderConfig", func(w http.ResponseWriter, r *http.Request) {
		fc := r.URL.Query().Get("feecut")
		if fc != "" {
			fci, err := strconv.Atoi(fc)
			if err != nil {
				glog.Errorf("Fee cut conversion failed: %v", err)
				return
			}
			TranscoderFeeCut = uint8(fci)
		}

		rc := r.URL.Query().Get("rewardcut")
		if rc != "" {
			rci, err := strconv.Atoi(rc)
			if err != nil {
				glog.Errorf("Reward cut conversion failed: %v", err)
				return
			}
			TranscoderRewardCut = uint8(rci)
		}

		p := r.URL.Query().Get("price")
		if p != "" {
			pi, err := strconv.Atoi(p)
			if err != nil {
				glog.Errorf("Price conversion failed: %v", err)
				return
			}
			TranscoderSegmentPrice = big.NewInt(int64(pi))
		}

		glog.Infof("Transcoder Fee Cut: %v, Transcoder Reward Cut: %v, Transcoder Segment Price: %v", TranscoderFeeCut, TranscoderRewardCut, TranscoderSegmentPrice)
	})

	//Bond some amount of tokens to a transcoder.
	http.HandleFunc("/bond", func(w http.ResponseWriter, r *http.Request) {
		addrStr := r.URL.Query().Get("addr")
		if addrStr == "" {
			glog.Errorf("Need to provide addr")
			return
		}

		amountStr := r.URL.Query().Get("amount")
		if amountStr == "" {
			glog.Errorf("Need to provide amount")
			return
		}

		addr := common.HexToAddress(addrStr)
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			glog.Errorf("Cannot convert amount: %v", err)
			return
		}

		eth.CheckRoundAndInit(s.LivepeerNode.Eth, EthRpcTimeout, EthMinedTxTimeout)
		tx, err := s.LivepeerNode.Eth.Bond(big.NewInt(int64(amount)), addr)
		if err != nil {
			glog.Errorf("Failed to bond: %v", err)
			return
		}
		receipt, err := eth.WaitForMinedTx(s.LivepeerNode.Eth.Backend(), EthRpcTimeout, EthMinedTxTimeout, tx.Hash())
		if err != nil {
			glog.Errorf("%v", err)
			return
		}
		if tx.Gas().Cmp(receipt.GasUsed) == 0 {
			glog.Errorf("Failed bonding: Ethereum Exception")
		}
	})

	//Print the transcoder's stake
	http.HandleFunc("/printStake", func(w http.ResponseWriter, r *http.Request) {
		printStake(s.LivepeerNode.Eth)
	})

	//Print the current broadcast HLS streamID
	http.HandleFunc("/streamID", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(LastHLSStreamID))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("StreamDB: %v", s.LivepeerNode.StreamDB)))
		w.Write([]byte(fmt.Sprintf("\n\nVideoNetwork: %v", s.LivepeerNode.VideoNetwork)))
	})

}
