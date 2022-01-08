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
	Bin: "0x608060405234801561001057600080fd5b5061030c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637e105ce21461003b578063f8b2cb4f1461006b575b600080fd5b6100556004803603810190610050919061018b565b61009b565b6040516100629190610251565b60405180910390f35b6100856004803603810190610080919061018b565b610107565b604051610092919061028c565b60405180910390f35b60606000823b905060008167ffffffffffffffff8111156100bf576100be6102a7565b5b6040519080825280601f01601f1916602001820160405280156100f15781602001600182028036833780820191505090505b50905081600060208301863c8092505050919050565b60008173ffffffffffffffffffffffffffffffffffffffff16319050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101588261012d565b9050919050565b6101688161014d565b811461017357600080fd5b50565b6000813590506101858161015f565b92915050565b6000602082840312156101a1576101a0610128565b5b60006101af84828501610176565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156101f25780820151818401526020810190506101d7565b83811115610201576000848401525b50505050565b6000601f19601f8301169050919050565b6000610223826101b8565b61022d81856101c3565b935061023d8185602086016101d4565b61024681610207565b840191505092915050565b6000602082019050818103600083015261026b8184610218565b905092915050565b6000819050919050565b61028681610273565b82525050565b60006020820190506102a1600083018461027d565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea2646970667358221220ab990cd5b96ec9d623d21ab3680a82edf49282914f53d5e1a8828992e82d567b64736f6c634300080a0033",
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
