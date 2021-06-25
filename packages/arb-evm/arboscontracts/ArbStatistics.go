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

// ArbStatisticsABI is the input ABI used to generate the binding from.
const ArbStatisticsABI = "[{\"inputs\":[],\"name\":\"getStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbStatistics is an auto generated Go binding around an Ethereum contract.
type ArbStatistics struct {
	ArbStatisticsCaller     // Read-only binding to the contract
	ArbStatisticsTransactor // Write-only binding to the contract
	ArbStatisticsFilterer   // Log filterer for contract events
}

// ArbStatisticsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbStatisticsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbStatisticsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbStatisticsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbStatisticsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbStatisticsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbStatisticsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbStatisticsSession struct {
	Contract     *ArbStatistics    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbStatisticsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbStatisticsCallerSession struct {
	Contract *ArbStatisticsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ArbStatisticsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbStatisticsTransactorSession struct {
	Contract     *ArbStatisticsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ArbStatisticsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbStatisticsRaw struct {
	Contract *ArbStatistics // Generic contract binding to access the raw methods on
}

// ArbStatisticsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbStatisticsCallerRaw struct {
	Contract *ArbStatisticsCaller // Generic read-only contract binding to access the raw methods on
}

// ArbStatisticsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbStatisticsTransactorRaw struct {
	Contract *ArbStatisticsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbStatistics creates a new instance of ArbStatistics, bound to a specific deployed contract.
func NewArbStatistics(address common.Address, backend bind.ContractBackend) (*ArbStatistics, error) {
	contract, err := bindArbStatistics(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbStatistics{ArbStatisticsCaller: ArbStatisticsCaller{contract: contract}, ArbStatisticsTransactor: ArbStatisticsTransactor{contract: contract}, ArbStatisticsFilterer: ArbStatisticsFilterer{contract: contract}}, nil
}

// NewArbStatisticsCaller creates a new read-only instance of ArbStatistics, bound to a specific deployed contract.
func NewArbStatisticsCaller(address common.Address, caller bind.ContractCaller) (*ArbStatisticsCaller, error) {
	contract, err := bindArbStatistics(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbStatisticsCaller{contract: contract}, nil
}

// NewArbStatisticsTransactor creates a new write-only instance of ArbStatistics, bound to a specific deployed contract.
func NewArbStatisticsTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbStatisticsTransactor, error) {
	contract, err := bindArbStatistics(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbStatisticsTransactor{contract: contract}, nil
}

// NewArbStatisticsFilterer creates a new log filterer instance of ArbStatistics, bound to a specific deployed contract.
func NewArbStatisticsFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbStatisticsFilterer, error) {
	contract, err := bindArbStatistics(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbStatisticsFilterer{contract: contract}, nil
}

// bindArbStatistics binds a generic wrapper to an already deployed contract.
func bindArbStatistics(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbStatisticsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbStatistics *ArbStatisticsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbStatistics.Contract.ArbStatisticsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbStatistics *ArbStatisticsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbStatistics.Contract.ArbStatisticsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbStatistics *ArbStatisticsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbStatistics.Contract.ArbStatisticsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbStatistics *ArbStatisticsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbStatistics.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbStatistics *ArbStatisticsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbStatistics.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbStatistics *ArbStatisticsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbStatistics.Contract.contract.Transact(opts, method, params...)
}

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbStatistics *ArbStatisticsCaller) GetStats(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbStatistics.contract.Call(opts, &out, "getStats")

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

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbStatistics *ArbStatisticsSession) GetStats() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbStatistics.Contract.GetStats(&_ArbStatistics.CallOpts)
}

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_ArbStatistics *ArbStatisticsCallerSession) GetStats() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbStatistics.Contract.GetStats(&_ArbStatistics.CallOpts)
}
