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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SubdomainRegistrarABI is the input ABI used to generate the binding from.
const SubdomainRegistrarABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferDomainOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_subnode\",\"type\":\"bytes32\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferNodeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"registerSubdomain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publicResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_publicResolver\",\"type\":\"address\"},{\"name\":\"_rootNode\",\"type\":\"bytes32\"},{\"name\":\"_rootName\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SubdomainRegistrar is an auto generated Go binding around an Ethereum contract.
type SubdomainRegistrar struct {
	SubdomainRegistrarCaller     // Read-only binding to the contract
	SubdomainRegistrarTransactor // Write-only binding to the contract
	SubdomainRegistrarFilterer   // Log filterer for contract events
}

// SubdomainRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubdomainRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubdomainRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubdomainRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubdomainRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubdomainRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubdomainRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubdomainRegistrarSession struct {
	Contract     *SubdomainRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SubdomainRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubdomainRegistrarCallerSession struct {
	Contract *SubdomainRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SubdomainRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubdomainRegistrarTransactorSession struct {
	Contract     *SubdomainRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SubdomainRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubdomainRegistrarRaw struct {
	Contract *SubdomainRegistrar // Generic contract binding to access the raw methods on
}

// SubdomainRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubdomainRegistrarCallerRaw struct {
	Contract *SubdomainRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// SubdomainRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubdomainRegistrarTransactorRaw struct {
	Contract *SubdomainRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubdomainRegistrar creates a new instance of SubdomainRegistrar, bound to a specific deployed contract.
func NewSubdomainRegistrar(address common.Address, backend bind.ContractBackend) (*SubdomainRegistrar, error) {
	contract, err := bindSubdomainRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubdomainRegistrar{SubdomainRegistrarCaller: SubdomainRegistrarCaller{contract: contract}, SubdomainRegistrarTransactor: SubdomainRegistrarTransactor{contract: contract}, SubdomainRegistrarFilterer: SubdomainRegistrarFilterer{contract: contract}}, nil
}

// NewSubdomainRegistrarCaller creates a new read-only instance of SubdomainRegistrar, bound to a specific deployed contract.
func NewSubdomainRegistrarCaller(address common.Address, caller bind.ContractCaller) (*SubdomainRegistrarCaller, error) {
	contract, err := bindSubdomainRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubdomainRegistrarCaller{contract: contract}, nil
}

// NewSubdomainRegistrarTransactor creates a new write-only instance of SubdomainRegistrar, bound to a specific deployed contract.
func NewSubdomainRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*SubdomainRegistrarTransactor, error) {
	contract, err := bindSubdomainRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubdomainRegistrarTransactor{contract: contract}, nil
}

// NewSubdomainRegistrarFilterer creates a new log filterer instance of SubdomainRegistrar, bound to a specific deployed contract.
func NewSubdomainRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*SubdomainRegistrarFilterer, error) {
	contract, err := bindSubdomainRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubdomainRegistrarFilterer{contract: contract}, nil
}

// bindSubdomainRegistrar binds a generic wrapper to an already deployed contract.
func bindSubdomainRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubdomainRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubdomainRegistrar *SubdomainRegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubdomainRegistrar.Contract.SubdomainRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubdomainRegistrar *SubdomainRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.SubdomainRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubdomainRegistrar *SubdomainRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.SubdomainRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubdomainRegistrar *SubdomainRegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubdomainRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubdomainRegistrar *SubdomainRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubdomainRegistrar *SubdomainRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubdomainRegistrar.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarSession) Ens() (common.Address, error) {
	return _SubdomainRegistrar.Contract.Ens(&_SubdomainRegistrar.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarCallerSession) Ens() (common.Address, error) {
	return _SubdomainRegistrar.Contract.Ens(&_SubdomainRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubdomainRegistrar.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarSession) Owner() (common.Address, error) {
	return _SubdomainRegistrar.Contract.Owner(&_SubdomainRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarCallerSession) Owner() (common.Address, error) {
	return _SubdomainRegistrar.Contract.Owner(&_SubdomainRegistrar.CallOpts)
}

// PublicResolver is a free data retrieval call binding the contract method 0xf8256121.
//
// Solidity: function publicResolver() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarCaller) PublicResolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubdomainRegistrar.contract.Call(opts, out, "publicResolver")
	return *ret0, err
}

// PublicResolver is a free data retrieval call binding the contract method 0xf8256121.
//
// Solidity: function publicResolver() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarSession) PublicResolver() (common.Address, error) {
	return _SubdomainRegistrar.Contract.PublicResolver(&_SubdomainRegistrar.CallOpts)
}

// PublicResolver is a free data retrieval call binding the contract method 0xf8256121.
//
// Solidity: function publicResolver() constant returns(address)
func (_SubdomainRegistrar *SubdomainRegistrarCallerSession) PublicResolver() (common.Address, error) {
	return _SubdomainRegistrar.Contract.PublicResolver(&_SubdomainRegistrar.CallOpts)
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_SubdomainRegistrar *SubdomainRegistrarCaller) RootName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SubdomainRegistrar.contract.Call(opts, out, "rootName")
	return *ret0, err
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_SubdomainRegistrar *SubdomainRegistrarSession) RootName() (string, error) {
	return _SubdomainRegistrar.Contract.RootName(&_SubdomainRegistrar.CallOpts)
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_SubdomainRegistrar *SubdomainRegistrarCallerSession) RootName() (string, error) {
	return _SubdomainRegistrar.Contract.RootName(&_SubdomainRegistrar.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_SubdomainRegistrar *SubdomainRegistrarCaller) RootNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SubdomainRegistrar.contract.Call(opts, out, "rootNode")
	return *ret0, err
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_SubdomainRegistrar *SubdomainRegistrarSession) RootNode() ([32]byte, error) {
	return _SubdomainRegistrar.Contract.RootNode(&_SubdomainRegistrar.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_SubdomainRegistrar *SubdomainRegistrarCallerSession) RootNode() ([32]byte, error) {
	return _SubdomainRegistrar.Contract.RootNode(&_SubdomainRegistrar.CallOpts)
}

// RegisterSubdomain is a paid mutator transaction binding the contract method 0x771e60de.
//
// Solidity: function registerSubdomain(_name string) returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactor) RegisterSubdomain(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _SubdomainRegistrar.contract.Transact(opts, "registerSubdomain", _name)
}

// RegisterSubdomain is a paid mutator transaction binding the contract method 0x771e60de.
//
// Solidity: function registerSubdomain(_name string) returns()
func (_SubdomainRegistrar *SubdomainRegistrarSession) RegisterSubdomain(_name string) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.RegisterSubdomain(&_SubdomainRegistrar.TransactOpts, _name)
}

// RegisterSubdomain is a paid mutator transaction binding the contract method 0x771e60de.
//
// Solidity: function registerSubdomain(_name string) returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactorSession) RegisterSubdomain(_name string) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.RegisterSubdomain(&_SubdomainRegistrar.TransactOpts, _name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubdomainRegistrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SubdomainRegistrar *SubdomainRegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.RenounceOwnership(&_SubdomainRegistrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.RenounceOwnership(&_SubdomainRegistrar.TransactOpts)
}

// TransferDomainOwnership is a paid mutator transaction binding the contract method 0x54bfc7e5.
//
// Solidity: function transferDomainOwnership(_to address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactor) TransferDomainOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.contract.Transact(opts, "transferDomainOwnership", _to)
}

// TransferDomainOwnership is a paid mutator transaction binding the contract method 0x54bfc7e5.
//
// Solidity: function transferDomainOwnership(_to address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarSession) TransferDomainOwnership(_to common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.TransferDomainOwnership(&_SubdomainRegistrar.TransactOpts, _to)
}

// TransferDomainOwnership is a paid mutator transaction binding the contract method 0x54bfc7e5.
//
// Solidity: function transferDomainOwnership(_to address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactorSession) TransferDomainOwnership(_to common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.TransferDomainOwnership(&_SubdomainRegistrar.TransactOpts, _to)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x61b93ed9.
//
// Solidity: function transferNodeOwnership(_subnode bytes32, _newOwner address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactor) TransferNodeOwnership(opts *bind.TransactOpts, _subnode [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.contract.Transact(opts, "transferNodeOwnership", _subnode, _newOwner)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x61b93ed9.
//
// Solidity: function transferNodeOwnership(_subnode bytes32, _newOwner address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarSession) TransferNodeOwnership(_subnode [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.TransferNodeOwnership(&_SubdomainRegistrar.TransactOpts, _subnode, _newOwner)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x61b93ed9.
//
// Solidity: function transferNodeOwnership(_subnode bytes32, _newOwner address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactorSession) TransferNodeOwnership(_subnode [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.TransferNodeOwnership(&_SubdomainRegistrar.TransactOpts, _subnode, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.TransferOwnership(&_SubdomainRegistrar.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_SubdomainRegistrar *SubdomainRegistrarTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SubdomainRegistrar.Contract.TransferOwnership(&_SubdomainRegistrar.TransactOpts, _newOwner)
}

// SubdomainRegistrarOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the SubdomainRegistrar contract.
type SubdomainRegistrarOwnershipRenouncedIterator struct {
	Event *SubdomainRegistrarOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *SubdomainRegistrarOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubdomainRegistrarOwnershipRenounced)
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
		it.Event = new(SubdomainRegistrarOwnershipRenounced)
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
func (it *SubdomainRegistrarOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubdomainRegistrarOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubdomainRegistrarOwnershipRenounced represents a OwnershipRenounced event raised by the SubdomainRegistrar contract.
type SubdomainRegistrarOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_SubdomainRegistrar *SubdomainRegistrarFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*SubdomainRegistrarOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _SubdomainRegistrar.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubdomainRegistrarOwnershipRenouncedIterator{contract: _SubdomainRegistrar.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_SubdomainRegistrar *SubdomainRegistrarFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *SubdomainRegistrarOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _SubdomainRegistrar.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubdomainRegistrarOwnershipRenounced)
				if err := _SubdomainRegistrar.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// SubdomainRegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SubdomainRegistrar contract.
type SubdomainRegistrarOwnershipTransferredIterator struct {
	Event *SubdomainRegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SubdomainRegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubdomainRegistrarOwnershipTransferred)
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
		it.Event = new(SubdomainRegistrarOwnershipTransferred)
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
func (it *SubdomainRegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubdomainRegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubdomainRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the SubdomainRegistrar contract.
type SubdomainRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SubdomainRegistrar *SubdomainRegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SubdomainRegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SubdomainRegistrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubdomainRegistrarOwnershipTransferredIterator{contract: _SubdomainRegistrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SubdomainRegistrar *SubdomainRegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SubdomainRegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SubdomainRegistrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubdomainRegistrarOwnershipTransferred)
				if err := _SubdomainRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
