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

// ArbInfoMetaData contains all meta data concerning the ArbInfo contract.
var ArbInfoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101ba806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637e105ce21461003b578063f8b2cb4f14610064575b600080fd5b61004e6100493660046100e9565b61008d565b60405161005b9190610119565b60405180910390f35b61007f6100723660046100e9565b6001600160a01b03163190565b60405190815260200161005b565b6060813b60008167ffffffffffffffff8111156100ac576100ac61016e565b6040519080825280601f01601f1916602001820160405280156100d6576020820181803683370190505b50905081600060208301863c9392505050565b6000602082840312156100fb57600080fd5b81356001600160a01b038116811461011257600080fd5b9392505050565b600060208083528351808285015260005b818110156101465785810183015185820160400152820161012a565b81811115610158576000604083870101525b50601f01601f1916929092016040019392505050565b634e487b7160e01b600052604160045260246000fdfea26469706673582212205f4f32f3428485229109817950665b5020c32f258ff8408368e677f22237dc2464736f6c634300080a0033",
}

// ArbInfoABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbInfoMetaData.ABI instead.
var ArbInfoABI = ArbInfoMetaData.ABI

// ArbInfoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArbInfoMetaData.Bin instead.
var ArbInfoBin = ArbInfoMetaData.Bin

// DeployArbInfo deploys a new Ethereum contract, binding an instance of ArbInfo to it.
func DeployArbInfo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbInfo, error) {
	parsed, err := ArbInfoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbInfoBin), backend)
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
