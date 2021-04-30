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

// ArbInfoABI is the input ABI used to generate the binding from.
const ArbInfoABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbInfoBin is the compiled bytecode used for deploying new contracts.
var ArbInfoBin = "0x608060405234801561001057600080fd5b50610211806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637e105ce21461003b578063f8b2cb4f146100f8575b600080fd5b61007d6004803603602081101561005157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610150565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100bd5780820151818401526020810190506100a2565b50505050905090810190601f1680156100ea5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61013a6004803603602081101561010e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101ba565b6040518082815260200191505060405180910390f35b60606000823b905060608167ffffffffffffffff8111801561017157600080fd5b506040519080825280601f01601f1916602001820160405280156101a45781602001600182028036833780820191505090505b50905081600060208301863c8092505050919050565b60008173ffffffffffffffffffffffffffffffffffffffff1631905091905056fea2646970667358221220483b2c0267d18e9080bddafc2cd0a2c10afd08ba37b502e758136e701716883a64736f6c634300060b0033"

// DeployArbInfo deploys a new Ethereum contract, binding an instance of ArbInfo to it.
func DeployArbInfo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbInfo, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbInfoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbInfoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbInfo{ArbInfoCaller: ArbInfoCaller{contract: contract}, ArbInfoTransactor: ArbInfoTransactor{contract: contract}, ArbInfoFilterer: ArbInfoFilterer{contract: contract}}, nil
}

// ArbInfo is an auto generated Go binding around an Ethereum contract.
type ArbInfo struct {
	ArbInfoCaller     // Read-only binding to the contract
	ArbInfoTransactor // Write-only binding to the contract
	ArbInfoFilterer   // Log filterer for contract events
}

// ArbInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbInfoSession struct {
	Contract     *ArbInfo          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbInfoCallerSession struct {
	Contract *ArbInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ArbInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbInfoTransactorSession struct {
	Contract     *ArbInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ArbInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbInfoRaw struct {
	Contract *ArbInfo // Generic contract binding to access the raw methods on
}

// ArbInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbInfoCallerRaw struct {
	Contract *ArbInfoCaller // Generic read-only contract binding to access the raw methods on
}

// ArbInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbInfoTransactorRaw struct {
	Contract *ArbInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbInfo creates a new instance of ArbInfo, bound to a specific deployed contract.
func NewArbInfo(address common.Address, backend bind.ContractBackend) (*ArbInfo, error) {
	contract, err := bindArbInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbInfo{ArbInfoCaller: ArbInfoCaller{contract: contract}, ArbInfoTransactor: ArbInfoTransactor{contract: contract}, ArbInfoFilterer: ArbInfoFilterer{contract: contract}}, nil
}

// NewArbInfoCaller creates a new read-only instance of ArbInfo, bound to a specific deployed contract.
func NewArbInfoCaller(address common.Address, caller bind.ContractCaller) (*ArbInfoCaller, error) {
	contract, err := bindArbInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbInfoCaller{contract: contract}, nil
}

// NewArbInfoTransactor creates a new write-only instance of ArbInfo, bound to a specific deployed contract.
func NewArbInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbInfoTransactor, error) {
	contract, err := bindArbInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbInfoTransactor{contract: contract}, nil
}

// NewArbInfoFilterer creates a new log filterer instance of ArbInfo, bound to a specific deployed contract.
func NewArbInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbInfoFilterer, error) {
	contract, err := bindArbInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbInfoFilterer{contract: contract}, nil
}

// bindArbInfo binds a generic wrapper to an already deployed contract.
func bindArbInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbInfo *ArbInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbInfo.Contract.ArbInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbInfo *ArbInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbInfo.Contract.ArbInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbInfo *ArbInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbInfo.Contract.ArbInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbInfo *ArbInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbInfo *ArbInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbInfo *ArbInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbInfo.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256)
func (_ArbInfo *ArbInfoCaller) GetBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArbInfo.contract.Call(opts, &out, "getBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256)
func (_ArbInfo *ArbInfoSession) GetBalance(account common.Address) (*big.Int, error) {
	return _ArbInfo.Contract.GetBalance(&_ArbInfo.CallOpts, account)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256)
func (_ArbInfo *ArbInfoCallerSession) GetBalance(account common.Address) (*big.Int, error) {
	return _ArbInfo.Contract.GetBalance(&_ArbInfo.CallOpts, account)
}

// GetCode is a free data retrieval call binding the contract method 0x7e105ce2.
//
// Solidity: function getCode(address account) view returns(bytes)
func (_ArbInfo *ArbInfoCaller) GetCode(opts *bind.CallOpts, account common.Address) ([]byte, error) {
	var out []interface{}
	err := _ArbInfo.contract.Call(opts, &out, "getCode", account)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCode is a free data retrieval call binding the contract method 0x7e105ce2.
//
// Solidity: function getCode(address account) view returns(bytes)
func (_ArbInfo *ArbInfoSession) GetCode(account common.Address) ([]byte, error) {
	return _ArbInfo.Contract.GetCode(&_ArbInfo.CallOpts, account)
}

// GetCode is a free data retrieval call binding the contract method 0x7e105ce2.
//
// Solidity: function getCode(address account) view returns(bytes)
func (_ArbInfo *ArbInfoCallerSession) GetCode(account common.Address) ([]byte, error) {
	return _ArbInfo.Contract.GetCode(&_ArbInfo.CallOpts, account)
}
