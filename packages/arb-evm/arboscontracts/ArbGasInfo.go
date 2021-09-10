// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arboscontracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbGasInfoMetaData contains all meta data concerning the ArbGasInfo contract.
var ArbGasInfoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getGasAccountingParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricesInArbGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"getPricesInArbGasWithAggregator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricesInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"getPricesInWeiWithAggregator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ArbGasInfoABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbGasInfoMetaData.ABI instead.
var ArbGasInfoABI = ArbGasInfoMetaData.ABI

// ArbGasInfo is an auto generated Go binding around an Ethereum contract.
type ArbGasInfo struct {
	ArbGasInfoCaller     // Read-only binding to the contract
	ArbGasInfoTransactor // Write-only binding to the contract
	ArbGasInfoFilterer   // Log filterer for contract events
}

// ArbGasInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbGasInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbGasInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbGasInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbGasInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbGasInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbGasInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbGasInfoSession struct {
	Contract     *ArbGasInfo       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbGasInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbGasInfoCallerSession struct {
	Contract *ArbGasInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbGasInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbGasInfoTransactorSession struct {
	Contract     *ArbGasInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbGasInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbGasInfoRaw struct {
	Contract *ArbGasInfo // Generic contract binding to access the raw methods on
}

// ArbGasInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbGasInfoCallerRaw struct {
	Contract *ArbGasInfoCaller // Generic read-only contract binding to access the raw methods on
}

// ArbGasInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbGasInfoTransactorRaw struct {
	Contract *ArbGasInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbGasInfo creates a new instance of ArbGasInfo, bound to a specific deployed contract.
func NewArbGasInfo(address common.Address, backend bind.ContractBackend) (*ArbGasInfo, error) {
	contract, err := bindArbGasInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbGasInfo{ArbGasInfoCaller: ArbGasInfoCaller{contract: contract}, ArbGasInfoTransactor: ArbGasInfoTransactor{contract: contract}, ArbGasInfoFilterer: ArbGasInfoFilterer{contract: contract}}, nil
}

// NewArbGasInfoCaller creates a new read-only instance of ArbGasInfo, bound to a specific deployed contract.
func NewArbGasInfoCaller(address common.Address, caller bind.ContractCaller) (*ArbGasInfoCaller, error) {
	contract, err := bindArbGasInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbGasInfoCaller{contract: contract}, nil
}

// NewArbGasInfoTransactor creates a new write-only instance of ArbGasInfo, bound to a specific deployed contract.
func NewArbGasInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbGasInfoTransactor, error) {
	contract, err := bindArbGasInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbGasInfoTransactor{contract: contract}, nil
}

// NewArbGasInfoFilterer creates a new log filterer instance of ArbGasInfo, bound to a specific deployed contract.
func NewArbGasInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbGasInfoFilterer, error) {
	contract, err := bindArbGasInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbGasInfoFilterer{contract: contract}, nil
}

// bindArbGasInfo binds a generic wrapper to an already deployed contract.
func bindArbGasInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbGasInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbGasInfo *ArbGasInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbGasInfo.Contract.ArbGasInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbGasInfo *ArbGasInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbGasInfo.Contract.ArbGasInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbGasInfo *ArbGasInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbGasInfo.Contract.ArbGasInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbGasInfo *ArbGasInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbGasInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbGasInfo *ArbGasInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbGasInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbGasInfo *ArbGasInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbGasInfo.Contract.contract.Transact(opts, method, params...)
}

// GetGasAccountingParams is a free data retrieval call binding the contract method 0x612af178.
//
// Solidity: function getGasAccountingParams() view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCaller) GetGasAccountingParams(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbGasInfo.contract.Call(opts, &out, "getGasAccountingParams")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetGasAccountingParams is a free data retrieval call binding the contract method 0x612af178.
//
// Solidity: function getGasAccountingParams() view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoSession) GetGasAccountingParams() (*big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetGasAccountingParams(&_ArbGasInfo.CallOpts)
}

// GetGasAccountingParams is a free data retrieval call binding the contract method 0x612af178.
//
// Solidity: function getGasAccountingParams() view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCallerSession) GetGasAccountingParams() (*big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetGasAccountingParams(&_ArbGasInfo.CallOpts)
}

// GetPricesInArbGas is a free data retrieval call binding the contract method 0x02199f34.
//
// Solidity: function getPricesInArbGas() view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCaller) GetPricesInArbGas(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbGasInfo.contract.Call(opts, &out, "getPricesInArbGas")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPricesInArbGas is a free data retrieval call binding the contract method 0x02199f34.
//
// Solidity: function getPricesInArbGas() view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoSession) GetPricesInArbGas() (*big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetPricesInArbGas(&_ArbGasInfo.CallOpts)
}

// GetPricesInArbGas is a free data retrieval call binding the contract method 0x02199f34.
//
// Solidity: function getPricesInArbGas() view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCallerSession) GetPricesInArbGas() (*big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetPricesInArbGas(&_ArbGasInfo.CallOpts)
}

// GetPricesInArbGasWithAggregator is a free data retrieval call binding the contract method 0x7a1ea732.
//
// Solidity: function getPricesInArbGasWithAggregator(address aggregator) view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCaller) GetPricesInArbGasWithAggregator(opts *bind.CallOpts, aggregator common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbGasInfo.contract.Call(opts, &out, "getPricesInArbGasWithAggregator", aggregator)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPricesInArbGasWithAggregator is a free data retrieval call binding the contract method 0x7a1ea732.
//
// Solidity: function getPricesInArbGasWithAggregator(address aggregator) view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoSession) GetPricesInArbGasWithAggregator(aggregator common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetPricesInArbGasWithAggregator(&_ArbGasInfo.CallOpts, aggregator)
}

// GetPricesInArbGasWithAggregator is a free data retrieval call binding the contract method 0x7a1ea732.
//
// Solidity: function getPricesInArbGasWithAggregator(address aggregator) view returns(uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCallerSession) GetPricesInArbGasWithAggregator(aggregator common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetPricesInArbGasWithAggregator(&_ArbGasInfo.CallOpts, aggregator)
}

// GetPricesInWei is a free data retrieval call binding the contract method 0x41b247a8.
//
// Solidity: function getPricesInWei() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCaller) GetPricesInWei(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbGasInfo.contract.Call(opts, &out, "getPricesInWei")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetPricesInWei is a free data retrieval call binding the contract method 0x41b247a8.
//
// Solidity: function getPricesInWei() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoSession) GetPricesInWei() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetPricesInWei(&_ArbGasInfo.CallOpts)
}

// GetPricesInWei is a free data retrieval call binding the contract method 0x41b247a8.
//
// Solidity: function getPricesInWei() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCallerSession) GetPricesInWei() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetPricesInWei(&_ArbGasInfo.CallOpts)
}

// GetPricesInWeiWithAggregator is a free data retrieval call binding the contract method 0xba9c916e.
//
// Solidity: function getPricesInWeiWithAggregator(address aggregator) view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCaller) GetPricesInWeiWithAggregator(opts *bind.CallOpts, aggregator common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbGasInfo.contract.Call(opts, &out, "getPricesInWeiWithAggregator", aggregator)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetPricesInWeiWithAggregator is a free data retrieval call binding the contract method 0xba9c916e.
//
// Solidity: function getPricesInWeiWithAggregator(address aggregator) view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoSession) GetPricesInWeiWithAggregator(aggregator common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetPricesInWeiWithAggregator(&_ArbGasInfo.CallOpts, aggregator)
}

// GetPricesInWeiWithAggregator is a free data retrieval call binding the contract method 0xba9c916e.
//
// Solidity: function getPricesInWeiWithAggregator(address aggregator) view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbGasInfo *ArbGasInfoCallerSession) GetPricesInWeiWithAggregator(aggregator common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbGasInfo.Contract.GetPricesInWeiWithAggregator(&_ArbGasInfo.CallOpts, aggregator)
}
