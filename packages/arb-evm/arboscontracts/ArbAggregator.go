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

// ArbAggregatorABI is the input ABI used to generate the binding from.
const ArbAggregatorABI = "[{\"inputs\":[],\"name\":\"getDefaultAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"getFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPreferredAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDefault\",\"type\":\"address\"}],\"name\":\"setDefaultAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prefAgg\",\"type\":\"address\"}],\"name\":\"setPreferredAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbAggregator is an auto generated Go binding around an Ethereum contract.
type ArbAggregator struct {
	ArbAggregatorCaller     // Read-only binding to the contract
	ArbAggregatorTransactor // Write-only binding to the contract
	ArbAggregatorFilterer   // Log filterer for contract events
}

// ArbAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbAggregatorSession struct {
	Contract     *ArbAggregator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbAggregatorCallerSession struct {
	Contract *ArbAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ArbAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbAggregatorTransactorSession struct {
	Contract     *ArbAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ArbAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbAggregatorRaw struct {
	Contract *ArbAggregator // Generic contract binding to access the raw methods on
}

// ArbAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbAggregatorCallerRaw struct {
	Contract *ArbAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// ArbAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbAggregatorTransactorRaw struct {
	Contract *ArbAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbAggregator creates a new instance of ArbAggregator, bound to a specific deployed contract.
func NewArbAggregator(address common.Address, backend bind.ContractBackend) (*ArbAggregator, error) {
	contract, err := bindArbAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbAggregator{ArbAggregatorCaller: ArbAggregatorCaller{contract: contract}, ArbAggregatorTransactor: ArbAggregatorTransactor{contract: contract}, ArbAggregatorFilterer: ArbAggregatorFilterer{contract: contract}}, nil
}

// NewArbAggregatorCaller creates a new read-only instance of ArbAggregator, bound to a specific deployed contract.
func NewArbAggregatorCaller(address common.Address, caller bind.ContractCaller) (*ArbAggregatorCaller, error) {
	contract, err := bindArbAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbAggregatorCaller{contract: contract}, nil
}

// NewArbAggregatorTransactor creates a new write-only instance of ArbAggregator, bound to a specific deployed contract.
func NewArbAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbAggregatorTransactor, error) {
	contract, err := bindArbAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbAggregatorTransactor{contract: contract}, nil
}

// NewArbAggregatorFilterer creates a new log filterer instance of ArbAggregator, bound to a specific deployed contract.
func NewArbAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbAggregatorFilterer, error) {
	contract, err := bindArbAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbAggregatorFilterer{contract: contract}, nil
}

// bindArbAggregator binds a generic wrapper to an already deployed contract.
func bindArbAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbAggregator *ArbAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbAggregator.Contract.ArbAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbAggregator *ArbAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbAggregator.Contract.ArbAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbAggregator *ArbAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbAggregator.Contract.ArbAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbAggregator *ArbAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbAggregator *ArbAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbAggregator *ArbAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbAggregator.Contract.contract.Transact(opts, method, params...)
}

// GetDefaultAggregator is a free data retrieval call binding the contract method 0x875883f2.
//
// Solidity: function getDefaultAggregator() view returns(address)
func (_ArbAggregator *ArbAggregatorCaller) GetDefaultAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbAggregator.contract.Call(opts, &out, "getDefaultAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDefaultAggregator is a free data retrieval call binding the contract method 0x875883f2.
//
// Solidity: function getDefaultAggregator() view returns(address)
func (_ArbAggregator *ArbAggregatorSession) GetDefaultAggregator() (common.Address, error) {
	return _ArbAggregator.Contract.GetDefaultAggregator(&_ArbAggregator.CallOpts)
}

// GetDefaultAggregator is a free data retrieval call binding the contract method 0x875883f2.
//
// Solidity: function getDefaultAggregator() view returns(address)
func (_ArbAggregator *ArbAggregatorCallerSession) GetDefaultAggregator() (common.Address, error) {
	return _ArbAggregator.Contract.GetDefaultAggregator(&_ArbAggregator.CallOpts)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x9c2c5bb5.
//
// Solidity: function getFeeCollector(address aggregator) view returns(address)
func (_ArbAggregator *ArbAggregatorCaller) GetFeeCollector(opts *bind.CallOpts, aggregator common.Address) (common.Address, error) {
	var out []interface{}
	err := _ArbAggregator.contract.Call(opts, &out, "getFeeCollector", aggregator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeCollector is a free data retrieval call binding the contract method 0x9c2c5bb5.
//
// Solidity: function getFeeCollector(address aggregator) view returns(address)
func (_ArbAggregator *ArbAggregatorSession) GetFeeCollector(aggregator common.Address) (common.Address, error) {
	return _ArbAggregator.Contract.GetFeeCollector(&_ArbAggregator.CallOpts, aggregator)
}

// GetFeeCollector is a free data retrieval call binding the contract method 0x9c2c5bb5.
//
// Solidity: function getFeeCollector(address aggregator) view returns(address)
func (_ArbAggregator *ArbAggregatorCallerSession) GetFeeCollector(aggregator common.Address) (common.Address, error) {
	return _ArbAggregator.Contract.GetFeeCollector(&_ArbAggregator.CallOpts, aggregator)
}

// GetPreferredAggregator is a free data retrieval call binding the contract method 0x52f10740.
//
// Solidity: function getPreferredAggregator(address addr) view returns(address, bool)
func (_ArbAggregator *ArbAggregatorCaller) GetPreferredAggregator(opts *bind.CallOpts, addr common.Address) (common.Address, bool, error) {
	var out []interface{}
	err := _ArbAggregator.contract.Call(opts, &out, "getPreferredAggregator", addr)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetPreferredAggregator is a free data retrieval call binding the contract method 0x52f10740.
//
// Solidity: function getPreferredAggregator(address addr) view returns(address, bool)
func (_ArbAggregator *ArbAggregatorSession) GetPreferredAggregator(addr common.Address) (common.Address, bool, error) {
	return _ArbAggregator.Contract.GetPreferredAggregator(&_ArbAggregator.CallOpts, addr)
}

// GetPreferredAggregator is a free data retrieval call binding the contract method 0x52f10740.
//
// Solidity: function getPreferredAggregator(address addr) view returns(address, bool)
func (_ArbAggregator *ArbAggregatorCallerSession) GetPreferredAggregator(addr common.Address) (common.Address, bool, error) {
	return _ArbAggregator.Contract.GetPreferredAggregator(&_ArbAggregator.CallOpts, addr)
}

// SetDefaultAggregator is a paid mutator transaction binding the contract method 0x0ffd6650.
//
// Solidity: function setDefaultAggregator(address newDefault) returns()
func (_ArbAggregator *ArbAggregatorTransactor) SetDefaultAggregator(opts *bind.TransactOpts, newDefault common.Address) (*types.Transaction, error) {
	return _ArbAggregator.contract.Transact(opts, "setDefaultAggregator", newDefault)
}

// SetDefaultAggregator is a paid mutator transaction binding the contract method 0x0ffd6650.
//
// Solidity: function setDefaultAggregator(address newDefault) returns()
func (_ArbAggregator *ArbAggregatorSession) SetDefaultAggregator(newDefault common.Address) (*types.Transaction, error) {
	return _ArbAggregator.Contract.SetDefaultAggregator(&_ArbAggregator.TransactOpts, newDefault)
}

// SetDefaultAggregator is a paid mutator transaction binding the contract method 0x0ffd6650.
//
// Solidity: function setDefaultAggregator(address newDefault) returns()
func (_ArbAggregator *ArbAggregatorTransactorSession) SetDefaultAggregator(newDefault common.Address) (*types.Transaction, error) {
	return _ArbAggregator.Contract.SetDefaultAggregator(&_ArbAggregator.TransactOpts, newDefault)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x29149799.
//
// Solidity: function setFeeCollector(address aggregator, address newFeeCollector) returns()
func (_ArbAggregator *ArbAggregatorTransactor) SetFeeCollector(opts *bind.TransactOpts, aggregator common.Address, newFeeCollector common.Address) (*types.Transaction, error) {
	return _ArbAggregator.contract.Transact(opts, "setFeeCollector", aggregator, newFeeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x29149799.
//
// Solidity: function setFeeCollector(address aggregator, address newFeeCollector) returns()
func (_ArbAggregator *ArbAggregatorSession) SetFeeCollector(aggregator common.Address, newFeeCollector common.Address) (*types.Transaction, error) {
	return _ArbAggregator.Contract.SetFeeCollector(&_ArbAggregator.TransactOpts, aggregator, newFeeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0x29149799.
//
// Solidity: function setFeeCollector(address aggregator, address newFeeCollector) returns()
func (_ArbAggregator *ArbAggregatorTransactorSession) SetFeeCollector(aggregator common.Address, newFeeCollector common.Address) (*types.Transaction, error) {
	return _ArbAggregator.Contract.SetFeeCollector(&_ArbAggregator.TransactOpts, aggregator, newFeeCollector)
}

// SetPreferredAggregator is a paid mutator transaction binding the contract method 0x6e928a6e.
//
// Solidity: function setPreferredAggregator(address prefAgg) returns()
func (_ArbAggregator *ArbAggregatorTransactor) SetPreferredAggregator(opts *bind.TransactOpts, prefAgg common.Address) (*types.Transaction, error) {
	return _ArbAggregator.contract.Transact(opts, "setPreferredAggregator", prefAgg)
}

// SetPreferredAggregator is a paid mutator transaction binding the contract method 0x6e928a6e.
//
// Solidity: function setPreferredAggregator(address prefAgg) returns()
func (_ArbAggregator *ArbAggregatorSession) SetPreferredAggregator(prefAgg common.Address) (*types.Transaction, error) {
	return _ArbAggregator.Contract.SetPreferredAggregator(&_ArbAggregator.TransactOpts, prefAgg)
}

// SetPreferredAggregator is a paid mutator transaction binding the contract method 0x6e928a6e.
//
// Solidity: function setPreferredAggregator(address prefAgg) returns()
func (_ArbAggregator *ArbAggregatorTransactorSession) SetPreferredAggregator(prefAgg common.Address) (*types.Transaction, error) {
	return _ArbAggregator.Contract.SetPreferredAggregator(&_ArbAggregator.TransactOpts, prefAgg)
}
