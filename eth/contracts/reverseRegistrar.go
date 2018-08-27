// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ReverseRegistrarABI is the input ABI used to generate the binding from.
const ReverseRegistrarABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimWithResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"node\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"},{\"name\":\"resolverAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReverseRegistrar is an auto generated Go binding around an Ethereum contract.
type ReverseRegistrar struct {
	ReverseRegistrarCaller     // Read-only binding to the contract
	ReverseRegistrarTransactor // Write-only binding to the contract
	ReverseRegistrarFilterer   // Log filterer for contract events
}

// ReverseRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReverseRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReverseRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReverseRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReverseRegistrarSession struct {
	Contract     *ReverseRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReverseRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReverseRegistrarCallerSession struct {
	Contract *ReverseRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ReverseRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReverseRegistrarTransactorSession struct {
	Contract     *ReverseRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReverseRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReverseRegistrarRaw struct {
	Contract *ReverseRegistrar // Generic contract binding to access the raw methods on
}

// ReverseRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReverseRegistrarCallerRaw struct {
	Contract *ReverseRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// ReverseRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReverseRegistrarTransactorRaw struct {
	Contract *ReverseRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReverseRegistrar creates a new instance of ReverseRegistrar, bound to a specific deployed contract.
func NewReverseRegistrar(address common.Address, backend bind.ContractBackend) (*ReverseRegistrar, error) {
	contract, err := bindReverseRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReverseRegistrar{ReverseRegistrarCaller: ReverseRegistrarCaller{contract: contract}, ReverseRegistrarTransactor: ReverseRegistrarTransactor{contract: contract}, ReverseRegistrarFilterer: ReverseRegistrarFilterer{contract: contract}}, nil
}

// NewReverseRegistrarCaller creates a new read-only instance of ReverseRegistrar, bound to a specific deployed contract.
func NewReverseRegistrarCaller(address common.Address, caller bind.ContractCaller) (*ReverseRegistrarCaller, error) {
	contract, err := bindReverseRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReverseRegistrarCaller{contract: contract}, nil
}

// NewReverseRegistrarTransactor creates a new write-only instance of ReverseRegistrar, bound to a specific deployed contract.
func NewReverseRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*ReverseRegistrarTransactor, error) {
	contract, err := bindReverseRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReverseRegistrarTransactor{contract: contract}, nil
}

// NewReverseRegistrarFilterer creates a new log filterer instance of ReverseRegistrar, bound to a specific deployed contract.
func NewReverseRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*ReverseRegistrarFilterer, error) {
	contract, err := bindReverseRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReverseRegistrarFilterer{contract: contract}, nil
}

// bindReverseRegistrar binds a generic wrapper to an already deployed contract.
func bindReverseRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReverseRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReverseRegistrar *ReverseRegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReverseRegistrar.Contract.ReverseRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReverseRegistrar *ReverseRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.ReverseRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReverseRegistrar *ReverseRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.ReverseRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReverseRegistrar *ReverseRegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReverseRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReverseRegistrar *ReverseRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReverseRegistrar *ReverseRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.contract.Transact(opts, method, params...)
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() constant returns(address)
func (_ReverseRegistrar *ReverseRegistrarCaller) DefaultResolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReverseRegistrar.contract.Call(opts, out, "defaultResolver")
	return *ret0, err
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() constant returns(address)
func (_ReverseRegistrar *ReverseRegistrarSession) DefaultResolver() (common.Address, error) {
	return _ReverseRegistrar.Contract.DefaultResolver(&_ReverseRegistrar.CallOpts)
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() constant returns(address)
func (_ReverseRegistrar *ReverseRegistrarCallerSession) DefaultResolver() (common.Address, error) {
	return _ReverseRegistrar.Contract.DefaultResolver(&_ReverseRegistrar.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseRegistrar *ReverseRegistrarCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReverseRegistrar.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseRegistrar *ReverseRegistrarSession) Ens() (common.Address, error) {
	return _ReverseRegistrar.Contract.Ens(&_ReverseRegistrar.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseRegistrar *ReverseRegistrarCallerSession) Ens() (common.Address, error) {
	return _ReverseRegistrar.Contract.Ens(&_ReverseRegistrar.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(addr address) constant returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarCaller) Node(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ReverseRegistrar.contract.Call(opts, out, "node", addr)
	return *ret0, err
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(addr address) constant returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarSession) Node(addr common.Address) ([32]byte, error) {
	return _ReverseRegistrar.Contract.Node(&_ReverseRegistrar.CallOpts, addr)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(addr address) constant returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarCallerSession) Node(addr common.Address) ([32]byte, error) {
	return _ReverseRegistrar.Contract.Node(&_ReverseRegistrar.CallOpts, addr)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(owner address) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarTransactor) Claim(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ReverseRegistrar.contract.Transact(opts, "claim", owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(owner address) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.Claim(&_ReverseRegistrar.TransactOpts, owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(owner address) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarTransactorSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.Claim(&_ReverseRegistrar.TransactOpts, owner)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(owner address, resolver address) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarTransactor) ClaimWithResolver(opts *bind.TransactOpts, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ReverseRegistrar.contract.Transact(opts, "claimWithResolver", owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(owner address, resolver address) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.ClaimWithResolver(&_ReverseRegistrar.TransactOpts, owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(owner address, resolver address) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarTransactorSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.ClaimWithResolver(&_ReverseRegistrar.TransactOpts, owner, resolver)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name string) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ReverseRegistrar.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name string) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarSession) SetName(name string) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.SetName(&_ReverseRegistrar.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name string) returns(bytes32)
func (_ReverseRegistrar *ReverseRegistrarTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _ReverseRegistrar.Contract.SetName(&_ReverseRegistrar.TransactOpts, name)
}
