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

// ArbBLSMetaData contains all meta data concerning the ArbBLS contract.
var ArbBLSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ArbBLSABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbBLSMetaData.ABI instead.
var ArbBLSABI = ArbBLSMetaData.ABI

// ArbBLS is an auto generated Go binding around an Ethereum contract.
type ArbBLS struct {
	ArbBLSCaller     // Read-only binding to the contract
	ArbBLSTransactor // Write-only binding to the contract
	ArbBLSFilterer   // Log filterer for contract events
}

// ArbBLSCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbBLSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBLSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbBLSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBLSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbBLSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBLSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbBLSSession struct {
	Contract     *ArbBLS           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbBLSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbBLSCallerSession struct {
	Contract *ArbBLSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbBLSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbBLSTransactorSession struct {
	Contract     *ArbBLSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbBLSRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbBLSRaw struct {
	Contract *ArbBLS // Generic contract binding to access the raw methods on
}

// ArbBLSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbBLSCallerRaw struct {
	Contract *ArbBLSCaller // Generic read-only contract binding to access the raw methods on
}

// ArbBLSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbBLSTransactorRaw struct {
	Contract *ArbBLSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbBLS creates a new instance of ArbBLS, bound to a specific deployed contract.
func NewArbBLS(address common.Address, backend bind.ContractBackend) (*ArbBLS, error) {
	contract, err := bindArbBLS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbBLS{ArbBLSCaller: ArbBLSCaller{contract: contract}, ArbBLSTransactor: ArbBLSTransactor{contract: contract}, ArbBLSFilterer: ArbBLSFilterer{contract: contract}}, nil
}

// NewArbBLSCaller creates a new read-only instance of ArbBLS, bound to a specific deployed contract.
func NewArbBLSCaller(address common.Address, caller bind.ContractCaller) (*ArbBLSCaller, error) {
	contract, err := bindArbBLS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbBLSCaller{contract: contract}, nil
}

// NewArbBLSTransactor creates a new write-only instance of ArbBLS, bound to a specific deployed contract.
func NewArbBLSTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbBLSTransactor, error) {
	contract, err := bindArbBLS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbBLSTransactor{contract: contract}, nil
}

// NewArbBLSFilterer creates a new log filterer instance of ArbBLS, bound to a specific deployed contract.
func NewArbBLSFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbBLSFilterer, error) {
	contract, err := bindArbBLS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbBLSFilterer{contract: contract}, nil
}

// bindArbBLS binds a generic wrapper to an already deployed contract.
func bindArbBLS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbBLSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbBLS *ArbBLSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbBLS.Contract.ArbBLSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbBLS *ArbBLSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBLS.Contract.ArbBLSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbBLS *ArbBLSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbBLS.Contract.ArbBLSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbBLS *ArbBLSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbBLS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbBLS *ArbBLSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBLS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbBLS *ArbBLSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbBLS.Contract.contract.Transact(opts, method, params...)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address addr) view returns(uint256, uint256, uint256, uint256)
func (_ArbBLS *ArbBLSCaller) GetPublicKey(opts *bind.CallOpts, addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbBLS.contract.Call(opts, &out, "getPublicKey", addr)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address addr) view returns(uint256, uint256, uint256, uint256)
func (_ArbBLS *ArbBLSSession) GetPublicKey(addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbBLS.Contract.GetPublicKey(&_ArbBLS.CallOpts, addr)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address addr) view returns(uint256, uint256, uint256, uint256)
func (_ArbBLS *ArbBLSCallerSession) GetPublicKey(addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbBLS.Contract.GetPublicKey(&_ArbBLS.CallOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0x375a7c7f.
//
// Solidity: function register(uint256 x0, uint256 x1, uint256 y0, uint256 y1) returns()
func (_ArbBLS *ArbBLSTransactor) Register(opts *bind.TransactOpts, x0 *big.Int, x1 *big.Int, y0 *big.Int, y1 *big.Int) (*types.Transaction, error) {
	return _ArbBLS.contract.Transact(opts, "register", x0, x1, y0, y1)
}

// Register is a paid mutator transaction binding the contract method 0x375a7c7f.
//
// Solidity: function register(uint256 x0, uint256 x1, uint256 y0, uint256 y1) returns()
func (_ArbBLS *ArbBLSSession) Register(x0 *big.Int, x1 *big.Int, y0 *big.Int, y1 *big.Int) (*types.Transaction, error) {
	return _ArbBLS.Contract.Register(&_ArbBLS.TransactOpts, x0, x1, y0, y1)
}

// Register is a paid mutator transaction binding the contract method 0x375a7c7f.
//
// Solidity: function register(uint256 x0, uint256 x1, uint256 y0, uint256 y1) returns()
func (_ArbBLS *ArbBLSTransactorSession) Register(x0 *big.Int, x1 *big.Int, y0 *big.Int, y1 *big.Int) (*types.Transaction, error) {
	return _ArbBLS.Contract.Register(&_ArbBLS.TransactOpts, x0, x1, y0, y1)
}
