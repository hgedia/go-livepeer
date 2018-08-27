// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RegisterSubdomainABI is the input ABI used to generate the binding from.
const RegisterSubdomainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferDomainOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultReverseResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"transferNodeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerSubdomain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reverseRegistrar\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publicResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_publicResolver\",\"type\":\"address\"},{\"name\":\"_defaultReverseResolver\",\"type\":\"address\"},{\"name\":\"_rootNode\",\"type\":\"bytes32\"},{\"name\":\"_rootName\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RegisterSubdomain is an auto generated Go binding around an Ethereum contract.
type RegisterSubdomain struct {
	RegisterSubdomainCaller     // Read-only binding to the contract
	RegisterSubdomainTransactor // Write-only binding to the contract
	RegisterSubdomainFilterer   // Log filterer for contract events
}

// RegisterSubdomainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegisterSubdomainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterSubdomainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegisterSubdomainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterSubdomainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegisterSubdomainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterSubdomainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegisterSubdomainSession struct {
	Contract     *RegisterSubdomain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RegisterSubdomainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegisterSubdomainCallerSession struct {
	Contract *RegisterSubdomainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RegisterSubdomainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegisterSubdomainTransactorSession struct {
	Contract     *RegisterSubdomainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RegisterSubdomainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegisterSubdomainRaw struct {
	Contract *RegisterSubdomain // Generic contract binding to access the raw methods on
}

// RegisterSubdomainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegisterSubdomainCallerRaw struct {
	Contract *RegisterSubdomainCaller // Generic read-only contract binding to access the raw methods on
}

// RegisterSubdomainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegisterSubdomainTransactorRaw struct {
	Contract *RegisterSubdomainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegisterSubdomain creates a new instance of RegisterSubdomain, bound to a specific deployed contract.
func NewRegisterSubdomain(address common.Address, backend bind.ContractBackend) (*RegisterSubdomain, error) {
	contract, err := bindRegisterSubdomain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegisterSubdomain{RegisterSubdomainCaller: RegisterSubdomainCaller{contract: contract}, RegisterSubdomainTransactor: RegisterSubdomainTransactor{contract: contract}, RegisterSubdomainFilterer: RegisterSubdomainFilterer{contract: contract}}, nil
}

// NewRegisterSubdomainCaller creates a new read-only instance of RegisterSubdomain, bound to a specific deployed contract.
func NewRegisterSubdomainCaller(address common.Address, caller bind.ContractCaller) (*RegisterSubdomainCaller, error) {
	contract, err := bindRegisterSubdomain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterSubdomainCaller{contract: contract}, nil
}

// NewRegisterSubdomainTransactor creates a new write-only instance of RegisterSubdomain, bound to a specific deployed contract.
func NewRegisterSubdomainTransactor(address common.Address, transactor bind.ContractTransactor) (*RegisterSubdomainTransactor, error) {
	contract, err := bindRegisterSubdomain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterSubdomainTransactor{contract: contract}, nil
}

// NewRegisterSubdomainFilterer creates a new log filterer instance of RegisterSubdomain, bound to a specific deployed contract.
func NewRegisterSubdomainFilterer(address common.Address, filterer bind.ContractFilterer) (*RegisterSubdomainFilterer, error) {
	contract, err := bindRegisterSubdomain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegisterSubdomainFilterer{contract: contract}, nil
}

// bindRegisterSubdomain binds a generic wrapper to an already deployed contract.
func bindRegisterSubdomain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegisterSubdomainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegisterSubdomain *RegisterSubdomainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RegisterSubdomain.Contract.RegisterSubdomainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegisterSubdomain *RegisterSubdomainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.RegisterSubdomainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegisterSubdomain *RegisterSubdomainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.RegisterSubdomainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegisterSubdomain *RegisterSubdomainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RegisterSubdomain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegisterSubdomain *RegisterSubdomainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegisterSubdomain *RegisterSubdomainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.contract.Transact(opts, method, params...)
}

// DefaultReverseResolver is a free data retrieval call binding the contract method 0x56299fbc.
//
// Solidity: function defaultReverseResolver() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCaller) DefaultReverseResolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RegisterSubdomain.contract.Call(opts, out, "defaultReverseResolver")
	return *ret0, err
}

// DefaultReverseResolver is a free data retrieval call binding the contract method 0x56299fbc.
//
// Solidity: function defaultReverseResolver() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainSession) DefaultReverseResolver() (common.Address, error) {
	return _RegisterSubdomain.Contract.DefaultReverseResolver(&_RegisterSubdomain.CallOpts)
}

// DefaultReverseResolver is a free data retrieval call binding the contract method 0x56299fbc.
//
// Solidity: function defaultReverseResolver() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCallerSession) DefaultReverseResolver() (common.Address, error) {
	return _RegisterSubdomain.Contract.DefaultReverseResolver(&_RegisterSubdomain.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RegisterSubdomain.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainSession) Ens() (common.Address, error) {
	return _RegisterSubdomain.Contract.Ens(&_RegisterSubdomain.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCallerSession) Ens() (common.Address, error) {
	return _RegisterSubdomain.Contract.Ens(&_RegisterSubdomain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RegisterSubdomain.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainSession) Owner() (common.Address, error) {
	return _RegisterSubdomain.Contract.Owner(&_RegisterSubdomain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCallerSession) Owner() (common.Address, error) {
	return _RegisterSubdomain.Contract.Owner(&_RegisterSubdomain.CallOpts)
}

// PublicResolver is a free data retrieval call binding the contract method 0xf8256121.
//
// Solidity: function publicResolver() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCaller) PublicResolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RegisterSubdomain.contract.Call(opts, out, "publicResolver")
	return *ret0, err
}

// PublicResolver is a free data retrieval call binding the contract method 0xf8256121.
//
// Solidity: function publicResolver() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainSession) PublicResolver() (common.Address, error) {
	return _RegisterSubdomain.Contract.PublicResolver(&_RegisterSubdomain.CallOpts)
}

// PublicResolver is a free data retrieval call binding the contract method 0xf8256121.
//
// Solidity: function publicResolver() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCallerSession) PublicResolver() (common.Address, error) {
	return _RegisterSubdomain.Contract.PublicResolver(&_RegisterSubdomain.CallOpts)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCaller) ReverseRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RegisterSubdomain.contract.Call(opts, out, "reverseRegistrar")
	return *ret0, err
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainSession) ReverseRegistrar() (common.Address, error) {
	return _RegisterSubdomain.Contract.ReverseRegistrar(&_RegisterSubdomain.CallOpts)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() constant returns(address)
func (_RegisterSubdomain *RegisterSubdomainCallerSession) ReverseRegistrar() (common.Address, error) {
	return _RegisterSubdomain.Contract.ReverseRegistrar(&_RegisterSubdomain.CallOpts)
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_RegisterSubdomain *RegisterSubdomainCaller) RootName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RegisterSubdomain.contract.Call(opts, out, "rootName")
	return *ret0, err
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_RegisterSubdomain *RegisterSubdomainSession) RootName() (string, error) {
	return _RegisterSubdomain.Contract.RootName(&_RegisterSubdomain.CallOpts)
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() constant returns(string)
func (_RegisterSubdomain *RegisterSubdomainCallerSession) RootName() (string, error) {
	return _RegisterSubdomain.Contract.RootName(&_RegisterSubdomain.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_RegisterSubdomain *RegisterSubdomainCaller) RootNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RegisterSubdomain.contract.Call(opts, out, "rootNode")
	return *ret0, err
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_RegisterSubdomain *RegisterSubdomainSession) RootNode() ([32]byte, error) {
	return _RegisterSubdomain.Contract.RootNode(&_RegisterSubdomain.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_RegisterSubdomain *RegisterSubdomainCallerSession) RootNode() ([32]byte, error) {
	return _RegisterSubdomain.Contract.RootNode(&_RegisterSubdomain.CallOpts)
}

// RegisterSubdomain is a paid mutator transaction binding the contract method 0x771e60de.
//
// Solidity: function registerSubdomain(name string) returns()
func (_RegisterSubdomain *RegisterSubdomainTransactor) RegisterSubdomain(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _RegisterSubdomain.contract.Transact(opts, "registerSubdomain", name)
}

// RegisterSubdomain is a paid mutator transaction binding the contract method 0x771e60de.
//
// Solidity: function registerSubdomain(name string) returns()
func (_RegisterSubdomain *RegisterSubdomainSession) RegisterSubdomain(name string) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.RegisterSubdomain(&_RegisterSubdomain.TransactOpts, name)
}

// RegisterSubdomain is a paid mutator transaction binding the contract method 0x771e60de.
//
// Solidity: function registerSubdomain(name string) returns()
func (_RegisterSubdomain *RegisterSubdomainTransactorSession) RegisterSubdomain(name string) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.RegisterSubdomain(&_RegisterSubdomain.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegisterSubdomain *RegisterSubdomainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegisterSubdomain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegisterSubdomain *RegisterSubdomainSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.RenounceOwnership(&_RegisterSubdomain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegisterSubdomain *RegisterSubdomainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.RenounceOwnership(&_RegisterSubdomain.TransactOpts)
}

// TransferDomainOwnership is a paid mutator transaction binding the contract method 0x54bfc7e5.
//
// Solidity: function transferDomainOwnership(to address) returns()
func (_RegisterSubdomain *RegisterSubdomainTransactor) TransferDomainOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.contract.Transact(opts, "transferDomainOwnership", to)
}

// TransferDomainOwnership is a paid mutator transaction binding the contract method 0x54bfc7e5.
//
// Solidity: function transferDomainOwnership(to address) returns()
func (_RegisterSubdomain *RegisterSubdomainSession) TransferDomainOwnership(to common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.TransferDomainOwnership(&_RegisterSubdomain.TransactOpts, to)
}

// TransferDomainOwnership is a paid mutator transaction binding the contract method 0x54bfc7e5.
//
// Solidity: function transferDomainOwnership(to address) returns()
func (_RegisterSubdomain *RegisterSubdomainTransactorSession) TransferDomainOwnership(to common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.TransferDomainOwnership(&_RegisterSubdomain.TransactOpts, to)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x61b93ed9.
//
// Solidity: function transferNodeOwnership(node bytes32, owner address) returns()
func (_RegisterSubdomain *RegisterSubdomainTransactor) TransferNodeOwnership(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.contract.Transact(opts, "transferNodeOwnership", node, owner)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x61b93ed9.
//
// Solidity: function transferNodeOwnership(node bytes32, owner address) returns()
func (_RegisterSubdomain *RegisterSubdomainSession) TransferNodeOwnership(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.TransferNodeOwnership(&_RegisterSubdomain.TransactOpts, node, owner)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x61b93ed9.
//
// Solidity: function transferNodeOwnership(node bytes32, owner address) returns()
func (_RegisterSubdomain *RegisterSubdomainTransactorSession) TransferNodeOwnership(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.TransferNodeOwnership(&_RegisterSubdomain.TransactOpts, node, owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RegisterSubdomain *RegisterSubdomainTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RegisterSubdomain *RegisterSubdomainSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.TransferOwnership(&_RegisterSubdomain.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RegisterSubdomain *RegisterSubdomainTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RegisterSubdomain.Contract.TransferOwnership(&_RegisterSubdomain.TransactOpts, _newOwner)
}

// RegisterSubdomainOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RegisterSubdomain contract.
type RegisterSubdomainOwnershipRenouncedIterator struct {
	Event *RegisterSubdomainOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RegisterSubdomainOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterSubdomainOwnershipRenounced)
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
		it.Event = new(RegisterSubdomainOwnershipRenounced)
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
func (it *RegisterSubdomainOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterSubdomainOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterSubdomainOwnershipRenounced represents a OwnershipRenounced event raised by the RegisterSubdomain contract.
type RegisterSubdomainOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RegisterSubdomain *RegisterSubdomainFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RegisterSubdomainOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RegisterSubdomain.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegisterSubdomainOwnershipRenouncedIterator{contract: _RegisterSubdomain.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RegisterSubdomain *RegisterSubdomainFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RegisterSubdomainOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RegisterSubdomain.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterSubdomainOwnershipRenounced)
				if err := _RegisterSubdomain.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RegisterSubdomainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RegisterSubdomain contract.
type RegisterSubdomainOwnershipTransferredIterator struct {
	Event *RegisterSubdomainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegisterSubdomainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterSubdomainOwnershipTransferred)
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
		it.Event = new(RegisterSubdomainOwnershipTransferred)
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
func (it *RegisterSubdomainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterSubdomainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterSubdomainOwnershipTransferred represents a OwnershipTransferred event raised by the RegisterSubdomain contract.
type RegisterSubdomainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RegisterSubdomain *RegisterSubdomainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegisterSubdomainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegisterSubdomain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegisterSubdomainOwnershipTransferredIterator{contract: _RegisterSubdomain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RegisterSubdomain *RegisterSubdomainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegisterSubdomainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegisterSubdomain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterSubdomainOwnershipTransferred)
				if err := _RegisterSubdomain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
