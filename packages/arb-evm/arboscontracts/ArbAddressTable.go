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

// ArbAddressTableABI is the input ABI used to generate the binding from.
const ArbAddressTableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"compress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"buf\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"decompress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"lookupIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbAddressTable is an auto generated Go binding around an Ethereum contract.
type ArbAddressTable struct {
	ArbAddressTableCaller     // Read-only binding to the contract
	ArbAddressTableTransactor // Write-only binding to the contract
	ArbAddressTableFilterer   // Log filterer for contract events
}

// ArbAddressTableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbAddressTableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbAddressTableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbAddressTableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbAddressTableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbAddressTableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbAddressTableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbAddressTableSession struct {
	Contract     *ArbAddressTable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbAddressTableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbAddressTableCallerSession struct {
	Contract *ArbAddressTableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ArbAddressTableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbAddressTableTransactorSession struct {
	Contract     *ArbAddressTableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ArbAddressTableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbAddressTableRaw struct {
	Contract *ArbAddressTable // Generic contract binding to access the raw methods on
}

// ArbAddressTableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbAddressTableCallerRaw struct {
	Contract *ArbAddressTableCaller // Generic read-only contract binding to access the raw methods on
}

// ArbAddressTableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbAddressTableTransactorRaw struct {
	Contract *ArbAddressTableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbAddressTable creates a new instance of ArbAddressTable, bound to a specific deployed contract.
func NewArbAddressTable(address common.Address, backend bind.ContractBackend) (*ArbAddressTable, error) {
	contract, err := bindArbAddressTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbAddressTable{ArbAddressTableCaller: ArbAddressTableCaller{contract: contract}, ArbAddressTableTransactor: ArbAddressTableTransactor{contract: contract}, ArbAddressTableFilterer: ArbAddressTableFilterer{contract: contract}}, nil
}

// NewArbAddressTableCaller creates a new read-only instance of ArbAddressTable, bound to a specific deployed contract.
func NewArbAddressTableCaller(address common.Address, caller bind.ContractCaller) (*ArbAddressTableCaller, error) {
	contract, err := bindArbAddressTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbAddressTableCaller{contract: contract}, nil
}

// NewArbAddressTableTransactor creates a new write-only instance of ArbAddressTable, bound to a specific deployed contract.
func NewArbAddressTableTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbAddressTableTransactor, error) {
	contract, err := bindArbAddressTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbAddressTableTransactor{contract: contract}, nil
}

// NewArbAddressTableFilterer creates a new log filterer instance of ArbAddressTable, bound to a specific deployed contract.
func NewArbAddressTableFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbAddressTableFilterer, error) {
	contract, err := bindArbAddressTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbAddressTableFilterer{contract: contract}, nil
}

// bindArbAddressTable binds a generic wrapper to an already deployed contract.
func bindArbAddressTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbAddressTableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbAddressTable *ArbAddressTableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbAddressTable.Contract.ArbAddressTableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbAddressTable *ArbAddressTableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbAddressTable.Contract.ArbAddressTableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbAddressTable *ArbAddressTableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbAddressTable.Contract.ArbAddressTableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbAddressTable *ArbAddressTableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbAddressTable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbAddressTable *ArbAddressTableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbAddressTable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbAddressTable *ArbAddressTableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbAddressTable.Contract.contract.Transact(opts, method, params...)
}

// AddressExists is a free data retrieval call binding the contract method 0xa5025222.
//
// Solidity: function addressExists(address addr) view returns(bool)
func (_ArbAddressTable *ArbAddressTableCaller) AddressExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ArbAddressTable.contract.Call(opts, &out, "addressExists", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressExists is a free data retrieval call binding the contract method 0xa5025222.
//
// Solidity: function addressExists(address addr) view returns(bool)
func (_ArbAddressTable *ArbAddressTableSession) AddressExists(addr common.Address) (bool, error) {
	return _ArbAddressTable.Contract.AddressExists(&_ArbAddressTable.CallOpts, addr)
}

// AddressExists is a free data retrieval call binding the contract method 0xa5025222.
//
// Solidity: function addressExists(address addr) view returns(bool)
func (_ArbAddressTable *ArbAddressTableCallerSession) AddressExists(addr common.Address) (bool, error) {
	return _ArbAddressTable.Contract.AddressExists(&_ArbAddressTable.CallOpts, addr)
}

// Decompress is a free data retrieval call binding the contract method 0x31862ada.
//
// Solidity: function decompress(bytes buf, uint256 offset) pure returns(address, uint256)
func (_ArbAddressTable *ArbAddressTableCaller) Decompress(opts *bind.CallOpts, buf []byte, offset *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ArbAddressTable.contract.Call(opts, &out, "decompress", buf, offset)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Decompress is a free data retrieval call binding the contract method 0x31862ada.
//
// Solidity: function decompress(bytes buf, uint256 offset) pure returns(address, uint256)
func (_ArbAddressTable *ArbAddressTableSession) Decompress(buf []byte, offset *big.Int) (common.Address, *big.Int, error) {
	return _ArbAddressTable.Contract.Decompress(&_ArbAddressTable.CallOpts, buf, offset)
}

// Decompress is a free data retrieval call binding the contract method 0x31862ada.
//
// Solidity: function decompress(bytes buf, uint256 offset) pure returns(address, uint256)
func (_ArbAddressTable *ArbAddressTableCallerSession) Decompress(buf []byte, offset *big.Int) (common.Address, *big.Int, error) {
	return _ArbAddressTable.Contract.Decompress(&_ArbAddressTable.CallOpts, buf, offset)
}

// Lookup is a free data retrieval call binding the contract method 0xd4b6b5da.
//
// Solidity: function lookup(address addr) view returns(uint256)
func (_ArbAddressTable *ArbAddressTableCaller) Lookup(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArbAddressTable.contract.Call(opts, &out, "lookup", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Lookup is a free data retrieval call binding the contract method 0xd4b6b5da.
//
// Solidity: function lookup(address addr) view returns(uint256)
func (_ArbAddressTable *ArbAddressTableSession) Lookup(addr common.Address) (*big.Int, error) {
	return _ArbAddressTable.Contract.Lookup(&_ArbAddressTable.CallOpts, addr)
}

// Lookup is a free data retrieval call binding the contract method 0xd4b6b5da.
//
// Solidity: function lookup(address addr) view returns(uint256)
func (_ArbAddressTable *ArbAddressTableCallerSession) Lookup(addr common.Address) (*big.Int, error) {
	return _ArbAddressTable.Contract.Lookup(&_ArbAddressTable.CallOpts, addr)
}

// LookupIndex is a free data retrieval call binding the contract method 0x8a186788.
//
// Solidity: function lookupIndex(uint256 index) view returns(address)
func (_ArbAddressTable *ArbAddressTableCaller) LookupIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ArbAddressTable.contract.Call(opts, &out, "lookupIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LookupIndex is a free data retrieval call binding the contract method 0x8a186788.
//
// Solidity: function lookupIndex(uint256 index) view returns(address)
func (_ArbAddressTable *ArbAddressTableSession) LookupIndex(index *big.Int) (common.Address, error) {
	return _ArbAddressTable.Contract.LookupIndex(&_ArbAddressTable.CallOpts, index)
}

// LookupIndex is a free data retrieval call binding the contract method 0x8a186788.
//
// Solidity: function lookupIndex(uint256 index) view returns(address)
func (_ArbAddressTable *ArbAddressTableCallerSession) LookupIndex(index *big.Int) (common.Address, error) {
	return _ArbAddressTable.Contract.LookupIndex(&_ArbAddressTable.CallOpts, index)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_ArbAddressTable *ArbAddressTableCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbAddressTable.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_ArbAddressTable *ArbAddressTableSession) Size() (*big.Int, error) {
	return _ArbAddressTable.Contract.Size(&_ArbAddressTable.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_ArbAddressTable *ArbAddressTableCallerSession) Size() (*big.Int, error) {
	return _ArbAddressTable.Contract.Size(&_ArbAddressTable.CallOpts)
}

// Compress is a paid mutator transaction binding the contract method 0xf6a455a2.
//
// Solidity: function compress(address addr) returns(bytes)
func (_ArbAddressTable *ArbAddressTableTransactor) Compress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ArbAddressTable.contract.Transact(opts, "compress", addr)
}

// Compress is a paid mutator transaction binding the contract method 0xf6a455a2.
//
// Solidity: function compress(address addr) returns(bytes)
func (_ArbAddressTable *ArbAddressTableSession) Compress(addr common.Address) (*types.Transaction, error) {
	return _ArbAddressTable.Contract.Compress(&_ArbAddressTable.TransactOpts, addr)
}

// Compress is a paid mutator transaction binding the contract method 0xf6a455a2.
//
// Solidity: function compress(address addr) returns(bytes)
func (_ArbAddressTable *ArbAddressTableTransactorSession) Compress(addr common.Address) (*types.Transaction, error) {
	return _ArbAddressTable.Contract.Compress(&_ArbAddressTable.TransactOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns(uint256)
func (_ArbAddressTable *ArbAddressTableTransactor) Register(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ArbAddressTable.contract.Transact(opts, "register", addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns(uint256)
func (_ArbAddressTable *ArbAddressTableSession) Register(addr common.Address) (*types.Transaction, error) {
	return _ArbAddressTable.Contract.Register(&_ArbAddressTable.TransactOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address addr) returns(uint256)
func (_ArbAddressTable *ArbAddressTableTransactorSession) Register(addr common.Address) (*types.Transaction, error) {
	return _ArbAddressTable.Contract.Register(&_ArbAddressTable.TransactOpts, addr)
}
