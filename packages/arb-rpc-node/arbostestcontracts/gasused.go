// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbostestcontracts

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

// GasUsedABI is the input ABI used to generate the binding from.
const GasUsedABI = "[{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[],\"name\":\"noop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sstore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GasUsedFuncSigs maps the 4-byte function signature to its string representation.
var GasUsedFuncSigs = map[string]string{
	"5dfc2e4a": "noop()",
	"703c2d1a": "sstore()",
}

// GasUsedBin is the compiled bytecode used for deploying new contracts.
var GasUsedBin = "0x608060405260878060116000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80635dfc2e4a146037578063703c2d1a14603f575b600080fd5b603d6045565b005b603d6047565b565b60008054600101905556fea265627a7a72315820cedc45ada7118340d59e28e3f93afd5e4a1b321b17b008f6d5002b16848b7a3764736f6c63430005110032"

// DeployGasUsed deploys a new Ethereum contract, binding an instance of GasUsed to it.
func DeployGasUsed(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasUsed, error) {
	parsed, err := abi.JSON(strings.NewReader(GasUsedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GasUsedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasUsed{GasUsedCaller: GasUsedCaller{contract: contract}, GasUsedTransactor: GasUsedTransactor{contract: contract}, GasUsedFilterer: GasUsedFilterer{contract: contract}}, nil
}

// GasUsed is an auto generated Go binding around an Ethereum contract.
type GasUsed struct {
	GasUsedCaller     // Read-only binding to the contract
	GasUsedTransactor // Write-only binding to the contract
	GasUsedFilterer   // Log filterer for contract events
}

// GasUsedCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasUsedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasUsedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasUsedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasUsedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasUsedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasUsedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasUsedSession struct {
	Contract     *GasUsed          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasUsedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasUsedCallerSession struct {
	Contract *GasUsedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GasUsedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasUsedTransactorSession struct {
	Contract     *GasUsedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GasUsedRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasUsedRaw struct {
	Contract *GasUsed // Generic contract binding to access the raw methods on
}

// GasUsedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasUsedCallerRaw struct {
	Contract *GasUsedCaller // Generic read-only contract binding to access the raw methods on
}

// GasUsedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasUsedTransactorRaw struct {
	Contract *GasUsedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasUsed creates a new instance of GasUsed, bound to a specific deployed contract.
func NewGasUsed(address common.Address, backend bind.ContractBackend) (*GasUsed, error) {
	contract, err := bindGasUsed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasUsed{GasUsedCaller: GasUsedCaller{contract: contract}, GasUsedTransactor: GasUsedTransactor{contract: contract}, GasUsedFilterer: GasUsedFilterer{contract: contract}}, nil
}

// NewGasUsedCaller creates a new read-only instance of GasUsed, bound to a specific deployed contract.
func NewGasUsedCaller(address common.Address, caller bind.ContractCaller) (*GasUsedCaller, error) {
	contract, err := bindGasUsed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasUsedCaller{contract: contract}, nil
}

// NewGasUsedTransactor creates a new write-only instance of GasUsed, bound to a specific deployed contract.
func NewGasUsedTransactor(address common.Address, transactor bind.ContractTransactor) (*GasUsedTransactor, error) {
	contract, err := bindGasUsed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasUsedTransactor{contract: contract}, nil
}

// NewGasUsedFilterer creates a new log filterer instance of GasUsed, bound to a specific deployed contract.
func NewGasUsedFilterer(address common.Address, filterer bind.ContractFilterer) (*GasUsedFilterer, error) {
	contract, err := bindGasUsed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasUsedFilterer{contract: contract}, nil
}

// bindGasUsed binds a generic wrapper to an already deployed contract.
func bindGasUsed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasUsedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasUsed *GasUsedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasUsed.Contract.GasUsedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasUsed *GasUsedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasUsed.Contract.GasUsedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasUsed *GasUsedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasUsed.Contract.GasUsedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasUsed *GasUsedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasUsed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasUsed *GasUsedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasUsed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasUsed *GasUsedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasUsed.Contract.contract.Transact(opts, method, params...)
}

// Noop is a paid mutator transaction binding the contract method 0x5dfc2e4a.
//
// Solidity: function noop() returns()
func (_GasUsed *GasUsedTransactor) Noop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasUsed.contract.Transact(opts, "noop")
}

// Noop is a paid mutator transaction binding the contract method 0x5dfc2e4a.
//
// Solidity: function noop() returns()
func (_GasUsed *GasUsedSession) Noop() (*types.Transaction, error) {
	return _GasUsed.Contract.Noop(&_GasUsed.TransactOpts)
}

// Noop is a paid mutator transaction binding the contract method 0x5dfc2e4a.
//
// Solidity: function noop() returns()
func (_GasUsed *GasUsedTransactorSession) Noop() (*types.Transaction, error) {
	return _GasUsed.Contract.Noop(&_GasUsed.TransactOpts)
}

// Sstore is a paid mutator transaction binding the contract method 0x703c2d1a.
//
// Solidity: function sstore() returns()
func (_GasUsed *GasUsedTransactor) Sstore(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasUsed.contract.Transact(opts, "sstore")
}

// Sstore is a paid mutator transaction binding the contract method 0x703c2d1a.
//
// Solidity: function sstore() returns()
func (_GasUsed *GasUsedSession) Sstore() (*types.Transaction, error) {
	return _GasUsed.Contract.Sstore(&_GasUsed.TransactOpts)
}

// Sstore is a paid mutator transaction binding the contract method 0x703c2d1a.
//
// Solidity: function sstore() returns()
func (_GasUsed *GasUsedTransactorSession) Sstore() (*types.Transaction, error) {
	return _GasUsed.Contract.Sstore(&_GasUsed.TransactOpts)
}
