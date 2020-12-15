// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arboscontracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbOwnerABI is the input ABI used to generate the binding from.
const ArbOwnerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"marshalledCode\",\"type\":\"bytes\"}],\"name\":\"continueArbosUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishArbosUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwnerAddr\",\"type\":\"address\"}],\"name\":\"giveOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"startArbosUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbOwnerFuncSigs maps the 4-byte function signature to its string representation.
var ArbOwnerFuncSigs = map[string]string{
	"0b766fec": "continueArbosUpgrade(bytes)",
	"e380002e": "finishArbosUpgrade()",
	"e3a0a148": "giveOwnership(address)",
	"6d92b8e4": "startArbosUpgrade()",
}

// ArbOwner is an auto generated Go binding around an Ethereum contract.
type ArbOwner struct {
	ArbOwnerCaller     // Read-only binding to the contract
	ArbOwnerTransactor // Write-only binding to the contract
	ArbOwnerFilterer   // Log filterer for contract events
}

// ArbOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbOwnerSession struct {
	Contract     *ArbOwner         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbOwnerCallerSession struct {
	Contract *ArbOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbOwnerTransactorSession struct {
	Contract     *ArbOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbOwnerRaw struct {
	Contract *ArbOwner // Generic contract binding to access the raw methods on
}

// ArbOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbOwnerCallerRaw struct {
	Contract *ArbOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// ArbOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbOwnerTransactorRaw struct {
	Contract *ArbOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbOwner creates a new instance of ArbOwner, bound to a specific deployed contract.
func NewArbOwner(address common.Address, backend bind.ContractBackend) (*ArbOwner, error) {
	contract, err := bindArbOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbOwner{ArbOwnerCaller: ArbOwnerCaller{contract: contract}, ArbOwnerTransactor: ArbOwnerTransactor{contract: contract}, ArbOwnerFilterer: ArbOwnerFilterer{contract: contract}}, nil
}

// NewArbOwnerCaller creates a new read-only instance of ArbOwner, bound to a specific deployed contract.
func NewArbOwnerCaller(address common.Address, caller bind.ContractCaller) (*ArbOwnerCaller, error) {
	contract, err := bindArbOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbOwnerCaller{contract: contract}, nil
}

// NewArbOwnerTransactor creates a new write-only instance of ArbOwner, bound to a specific deployed contract.
func NewArbOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbOwnerTransactor, error) {
	contract, err := bindArbOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbOwnerTransactor{contract: contract}, nil
}

// NewArbOwnerFilterer creates a new log filterer instance of ArbOwner, bound to a specific deployed contract.
func NewArbOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbOwnerFilterer, error) {
	contract, err := bindArbOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbOwnerFilterer{contract: contract}, nil
}

// bindArbOwner binds a generic wrapper to an already deployed contract.
func bindArbOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbOwner *ArbOwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbOwner.Contract.ArbOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbOwner *ArbOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.Contract.ArbOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbOwner *ArbOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbOwner.Contract.ArbOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbOwner *ArbOwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbOwner *ArbOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbOwner *ArbOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbOwner.Contract.contract.Transact(opts, method, params...)
}

// ContinueArbosUpgrade is a paid mutator transaction binding the contract method 0x0b766fec.
//
// Solidity: function continueArbosUpgrade(bytes marshalledCode) returns()
func (_ArbOwner *ArbOwnerTransactor) ContinueArbosUpgrade(opts *bind.TransactOpts, marshalledCode []byte) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "continueArbosUpgrade", marshalledCode)
}

// ContinueArbosUpgrade is a paid mutator transaction binding the contract method 0x0b766fec.
//
// Solidity: function continueArbosUpgrade(bytes marshalledCode) returns()
func (_ArbOwner *ArbOwnerSession) ContinueArbosUpgrade(marshalledCode []byte) (*types.Transaction, error) {
	return _ArbOwner.Contract.ContinueArbosUpgrade(&_ArbOwner.TransactOpts, marshalledCode)
}

// ContinueArbosUpgrade is a paid mutator transaction binding the contract method 0x0b766fec.
//
// Solidity: function continueArbosUpgrade(bytes marshalledCode) returns()
func (_ArbOwner *ArbOwnerTransactorSession) ContinueArbosUpgrade(marshalledCode []byte) (*types.Transaction, error) {
	return _ArbOwner.Contract.ContinueArbosUpgrade(&_ArbOwner.TransactOpts, marshalledCode)
}

// FinishArbosUpgrade is a paid mutator transaction binding the contract method 0xe380002e.
//
// Solidity: function finishArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerTransactor) FinishArbosUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "finishArbosUpgrade")
}

// FinishArbosUpgrade is a paid mutator transaction binding the contract method 0xe380002e.
//
// Solidity: function finishArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerSession) FinishArbosUpgrade() (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishArbosUpgrade(&_ArbOwner.TransactOpts)
}

// FinishArbosUpgrade is a paid mutator transaction binding the contract method 0xe380002e.
//
// Solidity: function finishArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerTransactorSession) FinishArbosUpgrade() (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishArbosUpgrade(&_ArbOwner.TransactOpts)
}

// GiveOwnership is a paid mutator transaction binding the contract method 0xe3a0a148.
//
// Solidity: function giveOwnership(address newOwnerAddr) returns()
func (_ArbOwner *ArbOwnerTransactor) GiveOwnership(opts *bind.TransactOpts, newOwnerAddr common.Address) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "giveOwnership", newOwnerAddr)
}

// GiveOwnership is a paid mutator transaction binding the contract method 0xe3a0a148.
//
// Solidity: function giveOwnership(address newOwnerAddr) returns()
func (_ArbOwner *ArbOwnerSession) GiveOwnership(newOwnerAddr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.GiveOwnership(&_ArbOwner.TransactOpts, newOwnerAddr)
}

// GiveOwnership is a paid mutator transaction binding the contract method 0xe3a0a148.
//
// Solidity: function giveOwnership(address newOwnerAddr) returns()
func (_ArbOwner *ArbOwnerTransactorSession) GiveOwnership(newOwnerAddr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.GiveOwnership(&_ArbOwner.TransactOpts, newOwnerAddr)
}

// StartArbosUpgrade is a paid mutator transaction binding the contract method 0x6d92b8e4.
//
// Solidity: function startArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerTransactor) StartArbosUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "startArbosUpgrade")
}

// StartArbosUpgrade is a paid mutator transaction binding the contract method 0x6d92b8e4.
//
// Solidity: function startArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerSession) StartArbosUpgrade() (*types.Transaction, error) {
	return _ArbOwner.Contract.StartArbosUpgrade(&_ArbOwner.TransactOpts)
}

// StartArbosUpgrade is a paid mutator transaction binding the contract method 0x6d92b8e4.
//
// Solidity: function startArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerTransactorSession) StartArbosUpgrade() (*types.Transaction, error) {
	return _ArbOwner.Contract.StartArbosUpgrade(&_ArbOwner.TransactOpts)
}
