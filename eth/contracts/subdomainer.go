// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// SubdomainerABI is the input ABI used to generate the binding from.
const SubdomainerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reverseRes\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subnode\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"subRegister\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"namePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"},{\"name\":\"revAdd\",\"type\":\"address\"},{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"subnode\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SubdomainRegistered\",\"type\":\"event\"}]"

// Subdomainer is an auto generated Go binding around an Ethereum contract.
type Subdomainer struct {
	SubdomainerCaller     // Read-only binding to the contract
	SubdomainerTransactor // Write-only binding to the contract
	SubdomainerFilterer   // Log filterer for contract events
}

// SubdomainerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubdomainerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubdomainerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubdomainerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubdomainerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubdomainerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubdomainerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubdomainerSession struct {
	Contract     *Subdomainer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubdomainerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubdomainerCallerSession struct {
	Contract *SubdomainerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SubdomainerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubdomainerTransactorSession struct {
	Contract     *SubdomainerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SubdomainerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubdomainerRaw struct {
	Contract *Subdomainer // Generic contract binding to access the raw methods on
}

// SubdomainerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubdomainerCallerRaw struct {
	Contract *SubdomainerCaller // Generic read-only contract binding to access the raw methods on
}

// SubdomainerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubdomainerTransactorRaw struct {
	Contract *SubdomainerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubdomainer creates a new instance of Subdomainer, bound to a specific deployed contract.
func NewSubdomainer(address common.Address, backend bind.ContractBackend) (*Subdomainer, error) {
	contract, err := bindSubdomainer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subdomainer{SubdomainerCaller: SubdomainerCaller{contract: contract}, SubdomainerTransactor: SubdomainerTransactor{contract: contract}, SubdomainerFilterer: SubdomainerFilterer{contract: contract}}, nil
}

// NewSubdomainerCaller creates a new read-only instance of Subdomainer, bound to a specific deployed contract.
func NewSubdomainerCaller(address common.Address, caller bind.ContractCaller) (*SubdomainerCaller, error) {
	contract, err := bindSubdomainer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubdomainerCaller{contract: contract}, nil
}

// NewSubdomainerTransactor creates a new write-only instance of Subdomainer, bound to a specific deployed contract.
func NewSubdomainerTransactor(address common.Address, transactor bind.ContractTransactor) (*SubdomainerTransactor, error) {
	contract, err := bindSubdomainer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubdomainerTransactor{contract: contract}, nil
}

// NewSubdomainerFilterer creates a new log filterer instance of Subdomainer, bound to a specific deployed contract.
func NewSubdomainerFilterer(address common.Address, filterer bind.ContractFilterer) (*SubdomainerFilterer, error) {
	contract, err := bindSubdomainer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubdomainerFilterer{contract: contract}, nil
}

// bindSubdomainer binds a generic wrapper to an already deployed contract.
func bindSubdomainer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubdomainerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subdomainer *SubdomainerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Subdomainer.Contract.SubdomainerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subdomainer *SubdomainerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subdomainer.Contract.SubdomainerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subdomainer *SubdomainerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subdomainer.Contract.SubdomainerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subdomainer *SubdomainerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Subdomainer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subdomainer *SubdomainerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subdomainer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subdomainer *SubdomainerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subdomainer.Contract.contract.Transact(opts, method, params...)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_Subdomainer *SubdomainerCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subdomainer.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_Subdomainer *SubdomainerSession) Ens() (common.Address, error) {
	return _Subdomainer.Contract.Ens(&_Subdomainer.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_Subdomainer *SubdomainerCallerSession) Ens() (common.Address, error) {
	return _Subdomainer.Contract.Ens(&_Subdomainer.CallOpts)
}

// NamePrice is a free data retrieval call binding the contract method 0xb22073b6.
//
// Solidity: function namePrice() constant returns(uint256)
func (_Subdomainer *SubdomainerCaller) NamePrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subdomainer.contract.Call(opts, out, "namePrice")
	return *ret0, err
}

// NamePrice is a free data retrieval call binding the contract method 0xb22073b6.
//
// Solidity: function namePrice() constant returns(uint256)
func (_Subdomainer *SubdomainerSession) NamePrice() (*big.Int, error) {
	return _Subdomainer.Contract.NamePrice(&_Subdomainer.CallOpts)
}

// NamePrice is a free data retrieval call binding the contract method 0xb22073b6.
//
// Solidity: function namePrice() constant returns(uint256)
func (_Subdomainer *SubdomainerCallerSession) NamePrice() (*big.Int, error) {
	return _Subdomainer.Contract.NamePrice(&_Subdomainer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subdomainer *SubdomainerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subdomainer.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subdomainer *SubdomainerSession) Owner() (common.Address, error) {
	return _Subdomainer.Contract.Owner(&_Subdomainer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subdomainer *SubdomainerCallerSession) Owner() (common.Address, error) {
	return _Subdomainer.Contract.Owner(&_Subdomainer.CallOpts)
}

// ReverseRes is a free data retrieval call binding the contract method 0x21f60fad.
//
// Solidity: function reverseRes() constant returns(address)
func (_Subdomainer *SubdomainerCaller) ReverseRes(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subdomainer.contract.Call(opts, out, "reverseRes")
	return *ret0, err
}

// ReverseRes is a free data retrieval call binding the contract method 0x21f60fad.
//
// Solidity: function reverseRes() constant returns(address)
func (_Subdomainer *SubdomainerSession) ReverseRes() (common.Address, error) {
	return _Subdomainer.Contract.ReverseRes(&_Subdomainer.CallOpts)
}

// ReverseRes is a free data retrieval call binding the contract method 0x21f60fad.
//
// Solidity: function reverseRes() constant returns(address)
func (_Subdomainer *SubdomainerCallerSession) ReverseRes() (common.Address, error) {
	return _Subdomainer.Contract.ReverseRes(&_Subdomainer.CallOpts)
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_Subdomainer *SubdomainerCaller) RootName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Subdomainer.contract.Call(opts, out, "rootName")
	return *ret0, err
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_Subdomainer *SubdomainerSession) RootName() (string, error) {
	return _Subdomainer.Contract.RootName(&_Subdomainer.CallOpts)
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_Subdomainer *SubdomainerCallerSession) RootName() (string, error) {
	return _Subdomainer.Contract.RootName(&_Subdomainer.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_Subdomainer *SubdomainerCaller) RootNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Subdomainer.contract.Call(opts, out, "rootNode")
	return *ret0, err
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_Subdomainer *SubdomainerSession) RootNode() ([32]byte, error) {
	return _Subdomainer.Contract.RootNode(&_Subdomainer.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_Subdomainer *SubdomainerCallerSession) RootNode() ([32]byte, error) {
	return _Subdomainer.Contract.RootNode(&_Subdomainer.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Subdomainer *SubdomainerTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Subdomainer.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Subdomainer *SubdomainerSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Subdomainer.Contract.ChangeOwner(&_Subdomainer.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Subdomainer *SubdomainerTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Subdomainer.Contract.ChangeOwner(&_Subdomainer.TransactOpts, newOwner)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(newPrice uint256) returns()
func (_Subdomainer *SubdomainerTransactor) ChangePrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _Subdomainer.contract.Transact(opts, "changePrice", newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(newPrice uint256) returns()
func (_Subdomainer *SubdomainerSession) ChangePrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Subdomainer.Contract.ChangePrice(&_Subdomainer.TransactOpts, newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(newPrice uint256) returns()
func (_Subdomainer *SubdomainerTransactorSession) ChangePrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Subdomainer.Contract.ChangePrice(&_Subdomainer.TransactOpts, newPrice)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, _owner address) returns()
func (_Subdomainer *SubdomainerTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _Subdomainer.contract.Transact(opts, "setOwner", node, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, _owner address) returns()
func (_Subdomainer *SubdomainerSession) SetOwner(node [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _Subdomainer.Contract.SetOwner(&_Subdomainer.TransactOpts, node, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, _owner address) returns()
func (_Subdomainer *SubdomainerTransactorSession) SetOwner(node [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _Subdomainer.Contract.SetOwner(&_Subdomainer.TransactOpts, node, _owner)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Subdomainer *SubdomainerTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Subdomainer.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Subdomainer *SubdomainerSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Subdomainer.Contract.SetResolver(&_Subdomainer.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Subdomainer *SubdomainerTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Subdomainer.Contract.SetResolver(&_Subdomainer.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, _owner address) returns()
func (_Subdomainer *SubdomainerTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _Subdomainer.contract.Transact(opts, "setSubnodeOwner", node, label, _owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, _owner address) returns()
func (_Subdomainer *SubdomainerSession) SetSubnodeOwner(node [32]byte, label [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _Subdomainer.Contract.SetSubnodeOwner(&_Subdomainer.TransactOpts, node, label, _owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, _owner address) returns()
func (_Subdomainer *SubdomainerTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _Subdomainer.Contract.SetSubnodeOwner(&_Subdomainer.TransactOpts, node, label, _owner)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Subdomainer *SubdomainerTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Subdomainer.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Subdomainer *SubdomainerSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Subdomainer.Contract.SetTTL(&_Subdomainer.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Subdomainer *SubdomainerTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Subdomainer.Contract.SetTTL(&_Subdomainer.TransactOpts, node, ttl)
}

// SubRegister is a paid mutator transaction binding the contract method 0x99b50c32.
//
// Solidity: function subRegister(subnode bytes32, name string) returns()
func (_Subdomainer *SubdomainerTransactor) SubRegister(opts *bind.TransactOpts, subnode [32]byte, name string) (*types.Transaction, error) {
	return _Subdomainer.contract.Transact(opts, "subRegister", subnode, name)
}

// SubRegister is a paid mutator transaction binding the contract method 0x99b50c32.
//
// Solidity: function subRegister(subnode bytes32, name string) returns()
func (_Subdomainer *SubdomainerSession) SubRegister(subnode [32]byte, name string) (*types.Transaction, error) {
	return _Subdomainer.Contract.SubRegister(&_Subdomainer.TransactOpts, subnode, name)
}

// SubRegister is a paid mutator transaction binding the contract method 0x99b50c32.
//
// Solidity: function subRegister(subnode bytes32, name string) returns()
func (_Subdomainer *SubdomainerTransactorSession) SubRegister(subnode [32]byte, name string) (*types.Transaction, error) {
	return _Subdomainer.Contract.SubRegister(&_Subdomainer.TransactOpts, subnode, name)
}

// SubdomainerOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Subdomainer contract.
type SubdomainerOwnerChangedIterator struct {
	Event *SubdomainerOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubdomainerOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubdomainerOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubdomainerOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubdomainerOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubdomainerOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubdomainerOwnerChanged represents a OwnerChanged event raised by the Subdomainer contract.
type SubdomainerOwnerChanged struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: e OwnerChanged(newOwner address)
func (_Subdomainer *SubdomainerFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SubdomainerOwnerChangedIterator, error) {

	logs, sub, err := _Subdomainer.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SubdomainerOwnerChangedIterator{contract: _Subdomainer.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: e OwnerChanged(newOwner address)
func (_Subdomainer *SubdomainerFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SubdomainerOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Subdomainer.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubdomainerOwnerChanged)
				if err := _Subdomainer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SubdomainerPriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the Subdomainer contract.
type SubdomainerPriceChangedIterator struct {
	Event *SubdomainerPriceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubdomainerPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubdomainerPriceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubdomainerPriceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubdomainerPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubdomainerPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubdomainerPriceChanged represents a PriceChanged event raised by the Subdomainer contract.
type SubdomainerPriceChanged struct {
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0xa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d622.
//
// Solidity: e PriceChanged(newPrice uint256)
func (_Subdomainer *SubdomainerFilterer) FilterPriceChanged(opts *bind.FilterOpts) (*SubdomainerPriceChangedIterator, error) {

	logs, sub, err := _Subdomainer.contract.FilterLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return &SubdomainerPriceChangedIterator{contract: _Subdomainer.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0xa6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d622.
//
// Solidity: e PriceChanged(newPrice uint256)
func (_Subdomainer *SubdomainerFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *SubdomainerPriceChanged) (event.Subscription, error) {

	logs, sub, err := _Subdomainer.contract.WatchLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubdomainerPriceChanged)
				if err := _Subdomainer.contract.UnpackLog(event, "PriceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SubdomainerSubdomainRegisteredIterator is returned from FilterSubdomainRegistered and is used to iterate over the raw logs and unpacked data for SubdomainRegistered events raised by the Subdomainer contract.
type SubdomainerSubdomainRegisteredIterator struct {
	Event *SubdomainerSubdomainRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubdomainerSubdomainRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubdomainerSubdomainRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubdomainerSubdomainRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubdomainerSubdomainRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubdomainerSubdomainRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubdomainerSubdomainRegistered represents a SubdomainRegistered event raised by the Subdomainer contract.
type SubdomainerSubdomainRegistered struct {
	Subnode [32]byte
	Owner   common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubdomainRegistered is a free log retrieval operation binding the contract event 0xc0e2ffa068f06071df9cac08e046d8cf99e02abbb92be4f8fa9ce589221d997e.
//
// Solidity: e SubdomainRegistered(subnode bytes32, owner indexed address, price uint256)
func (_Subdomainer *SubdomainerFilterer) FilterSubdomainRegistered(opts *bind.FilterOpts, owner []common.Address) (*SubdomainerSubdomainRegisteredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Subdomainer.contract.FilterLogs(opts, "SubdomainRegistered", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SubdomainerSubdomainRegisteredIterator{contract: _Subdomainer.contract, event: "SubdomainRegistered", logs: logs, sub: sub}, nil
}

// WatchSubdomainRegistered is a free log subscription operation binding the contract event 0xc0e2ffa068f06071df9cac08e046d8cf99e02abbb92be4f8fa9ce589221d997e.
//
// Solidity: e SubdomainRegistered(subnode bytes32, owner indexed address, price uint256)
func (_Subdomainer *SubdomainerFilterer) WatchSubdomainRegistered(opts *bind.WatchOpts, sink chan<- *SubdomainerSubdomainRegistered, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Subdomainer.contract.WatchLogs(opts, "SubdomainRegistered", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubdomainerSubdomainRegistered)
				if err := _Subdomainer.contract.UnpackLog(event, "SubdomainRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
