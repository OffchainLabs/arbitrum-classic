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

// ArbFunctionTableABI is the input ABI used to generate the binding from.
const ArbFunctionTableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"buf\",\"type\":\"bytes\"}],\"name\":\"upload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbFunctionTableFuncSigs maps the 4-byte function signature to its string representation.
var ArbFunctionTableFuncSigs = map[string]string{
	"b464631b": "get(address,uint256)",
	"88987068": "size(address)",
	"ce2ae159": "upload(bytes)",
}

// ArbFunctionTable is an auto generated Go binding around an Ethereum contract.
type ArbFunctionTable struct {
	ArbFunctionTableCaller     // Read-only binding to the contract
	ArbFunctionTableTransactor // Write-only binding to the contract
	ArbFunctionTableFilterer   // Log filterer for contract events
}

// ArbFunctionTableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbFunctionTableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFunctionTableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbFunctionTableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFunctionTableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbFunctionTableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFunctionTableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbFunctionTableSession struct {
	Contract     *ArbFunctionTable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbFunctionTableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbFunctionTableCallerSession struct {
	Contract *ArbFunctionTableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ArbFunctionTableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbFunctionTableTransactorSession struct {
	Contract     *ArbFunctionTableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ArbFunctionTableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbFunctionTableRaw struct {
	Contract *ArbFunctionTable // Generic contract binding to access the raw methods on
}

// ArbFunctionTableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbFunctionTableCallerRaw struct {
	Contract *ArbFunctionTableCaller // Generic read-only contract binding to access the raw methods on
}

// ArbFunctionTableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbFunctionTableTransactorRaw struct {
	Contract *ArbFunctionTableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbFunctionTable creates a new instance of ArbFunctionTable, bound to a specific deployed contract.
func NewArbFunctionTable(address common.Address, backend bind.ContractBackend) (*ArbFunctionTable, error) {
	contract, err := bindArbFunctionTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbFunctionTable{ArbFunctionTableCaller: ArbFunctionTableCaller{contract: contract}, ArbFunctionTableTransactor: ArbFunctionTableTransactor{contract: contract}, ArbFunctionTableFilterer: ArbFunctionTableFilterer{contract: contract}}, nil
}

// NewArbFunctionTableCaller creates a new read-only instance of ArbFunctionTable, bound to a specific deployed contract.
func NewArbFunctionTableCaller(address common.Address, caller bind.ContractCaller) (*ArbFunctionTableCaller, error) {
	contract, err := bindArbFunctionTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbFunctionTableCaller{contract: contract}, nil
}

// NewArbFunctionTableTransactor creates a new write-only instance of ArbFunctionTable, bound to a specific deployed contract.
func NewArbFunctionTableTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbFunctionTableTransactor, error) {
	contract, err := bindArbFunctionTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbFunctionTableTransactor{contract: contract}, nil
}

// NewArbFunctionTableFilterer creates a new log filterer instance of ArbFunctionTable, bound to a specific deployed contract.
func NewArbFunctionTableFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbFunctionTableFilterer, error) {
	contract, err := bindArbFunctionTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbFunctionTableFilterer{contract: contract}, nil
}

// bindArbFunctionTable binds a generic wrapper to an already deployed contract.
func bindArbFunctionTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbFunctionTableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbFunctionTable *ArbFunctionTableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbFunctionTable.Contract.ArbFunctionTableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbFunctionTable *ArbFunctionTableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbFunctionTable.Contract.ArbFunctionTableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbFunctionTable *ArbFunctionTableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbFunctionTable.Contract.ArbFunctionTableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbFunctionTable *ArbFunctionTableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbFunctionTable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbFunctionTable *ArbFunctionTableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbFunctionTable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbFunctionTable *ArbFunctionTableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbFunctionTable.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address addr, uint256 index) view returns(uint256, bool, uint256)
func (_ArbFunctionTable *ArbFunctionTableCaller) Get(opts *bind.CallOpts, addr common.Address, index *big.Int) (*big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _ArbFunctionTable.contract.Call(opts, &out, "get", addr, index)

	if err != nil {
		return *new(*big.Int), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// Get is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address addr, uint256 index) view returns(uint256, bool, uint256)
func (_ArbFunctionTable *ArbFunctionTableSession) Get(addr common.Address, index *big.Int) (*big.Int, bool, *big.Int, error) {
	return _ArbFunctionTable.Contract.Get(&_ArbFunctionTable.CallOpts, addr, index)
}

// Get is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address addr, uint256 index) view returns(uint256, bool, uint256)
func (_ArbFunctionTable *ArbFunctionTableCallerSession) Get(addr common.Address, index *big.Int) (*big.Int, bool, *big.Int, error) {
	return _ArbFunctionTable.Contract.Get(&_ArbFunctionTable.CallOpts, addr, index)
}

// Size is a free data retrieval call binding the contract method 0x88987068.
//
// Solidity: function size(address addr) view returns(uint256)
func (_ArbFunctionTable *ArbFunctionTableCaller) Size(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArbFunctionTable.contract.Call(opts, &out, "size", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x88987068.
//
// Solidity: function size(address addr) view returns(uint256)
func (_ArbFunctionTable *ArbFunctionTableSession) Size(addr common.Address) (*big.Int, error) {
	return _ArbFunctionTable.Contract.Size(&_ArbFunctionTable.CallOpts, addr)
}

// Size is a free data retrieval call binding the contract method 0x88987068.
//
// Solidity: function size(address addr) view returns(uint256)
func (_ArbFunctionTable *ArbFunctionTableCallerSession) Size(addr common.Address) (*big.Int, error) {
	return _ArbFunctionTable.Contract.Size(&_ArbFunctionTable.CallOpts, addr)
}

// Upload is a paid mutator transaction binding the contract method 0xce2ae159.
//
// Solidity: function upload(bytes buf) returns()
func (_ArbFunctionTable *ArbFunctionTableTransactor) Upload(opts *bind.TransactOpts, buf []byte) (*types.Transaction, error) {
	return _ArbFunctionTable.contract.Transact(opts, "upload", buf)
}

// Upload is a paid mutator transaction binding the contract method 0xce2ae159.
//
// Solidity: function upload(bytes buf) returns()
func (_ArbFunctionTable *ArbFunctionTableSession) Upload(buf []byte) (*types.Transaction, error) {
	return _ArbFunctionTable.Contract.Upload(&_ArbFunctionTable.TransactOpts, buf)
}

// Upload is a paid mutator transaction binding the contract method 0xce2ae159.
//
// Solidity: function upload(bytes buf) returns()
func (_ArbFunctionTable *ArbFunctionTableTransactorSession) Upload(buf []byte) (*types.Transaction, error) {
	return _ArbFunctionTable.Contract.Upload(&_ArbFunctionTable.TransactOpts, buf)
}
