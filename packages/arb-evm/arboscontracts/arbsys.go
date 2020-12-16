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

// ArbSysABI is the input ABI used to generate the binding from.
const ArbSysABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ERC721Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawal\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressTable_addressExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressTable_compress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"buf\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"addressTable_decompress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressTable_lookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"addressTable_lookupIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressTable_register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressTable_size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"functionTableGet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"functionTableSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBlsPublicKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"}],\"name\":\"registerBlsKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"buf\",\"type\":\"bytes\"}],\"name\":\"uploadFunctionTable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ArbSysFuncSigs maps the 4-byte function signature to its string representation.
var ArbSysFuncSigs = map[string]string{
	"43683c1b": "addressTable_addressExists(address)",
	"2a890b98": "addressTable_compress(address)",
	"be05f2b9": "addressTable_decompress(bytes,uint256)",
	"8ee24222": "addressTable_lookup(address)",
	"4c352200": "addressTable_lookupIndex(uint256)",
	"aca1dc65": "addressTable_register(address)",
	"03092fdb": "addressTable_size()",
	"d6ed0f56": "functionTableGet(address,uint256)",
	"1221c896": "functionTableSize(address)",
	"647e0e98": "getBlsPublicKey(address)",
	"a169625f": "getStorageAt(address,uint256)",
	"23ca0cd2": "getTransactionCount(address)",
	"1ff548ca": "registerBlsKey(uint256,uint256,uint256,uint256)",
	"b7cf9527": "uploadFunctionTable(bytes)",
	"a1db9782": "withdrawERC20(address,uint256)",
	"f3e414f8": "withdrawERC721(address,uint256)",
	"25e16063": "withdrawEth(address)",
}

// ArbSys is an auto generated Go binding around an Ethereum contract.
type ArbSys struct {
	ArbSysCaller     // Read-only binding to the contract
	ArbSysTransactor // Write-only binding to the contract
	ArbSysFilterer   // Log filterer for contract events
}

// ArbSysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbSysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbSysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbSysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbSysSession struct {
	Contract     *ArbSys           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbSysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbSysCallerSession struct {
	Contract *ArbSysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbSysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbSysTransactorSession struct {
	Contract     *ArbSysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbSysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbSysRaw struct {
	Contract *ArbSys // Generic contract binding to access the raw methods on
}

// ArbSysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbSysCallerRaw struct {
	Contract *ArbSysCaller // Generic read-only contract binding to access the raw methods on
}

// ArbSysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbSysTransactorRaw struct {
	Contract *ArbSysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbSys creates a new instance of ArbSys, bound to a specific deployed contract.
func NewArbSys(address common.Address, backend bind.ContractBackend) (*ArbSys, error) {
	contract, err := bindArbSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbSys{ArbSysCaller: ArbSysCaller{contract: contract}, ArbSysTransactor: ArbSysTransactor{contract: contract}, ArbSysFilterer: ArbSysFilterer{contract: contract}}, nil
}

// NewArbSysCaller creates a new read-only instance of ArbSys, bound to a specific deployed contract.
func NewArbSysCaller(address common.Address, caller bind.ContractCaller) (*ArbSysCaller, error) {
	contract, err := bindArbSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbSysCaller{contract: contract}, nil
}

// NewArbSysTransactor creates a new write-only instance of ArbSys, bound to a specific deployed contract.
func NewArbSysTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbSysTransactor, error) {
	contract, err := bindArbSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbSysTransactor{contract: contract}, nil
}

// NewArbSysFilterer creates a new log filterer instance of ArbSys, bound to a specific deployed contract.
func NewArbSysFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbSysFilterer, error) {
	contract, err := bindArbSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbSysFilterer{contract: contract}, nil
}

// bindArbSys binds a generic wrapper to an already deployed contract.
func bindArbSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbSysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbSys *ArbSysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbSys.Contract.ArbSysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbSys *ArbSysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbSys.Contract.ArbSysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbSys *ArbSysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbSys.Contract.ArbSysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbSys *ArbSysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbSys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbSys *ArbSysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbSys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbSys *ArbSysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbSys.Contract.contract.Transact(opts, method, params...)
}

// AddressTableAddressExists is a free data retrieval call binding the contract method 0x43683c1b.
//
// Solidity: function addressTable_addressExists(address addr) view returns(bool)
func (_ArbSys *ArbSysCaller) AddressTableAddressExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "addressTable_addressExists", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressTableAddressExists is a free data retrieval call binding the contract method 0x43683c1b.
//
// Solidity: function addressTable_addressExists(address addr) view returns(bool)
func (_ArbSys *ArbSysSession) AddressTableAddressExists(addr common.Address) (bool, error) {
	return _ArbSys.Contract.AddressTableAddressExists(&_ArbSys.CallOpts, addr)
}

// AddressTableAddressExists is a free data retrieval call binding the contract method 0x43683c1b.
//
// Solidity: function addressTable_addressExists(address addr) view returns(bool)
func (_ArbSys *ArbSysCallerSession) AddressTableAddressExists(addr common.Address) (bool, error) {
	return _ArbSys.Contract.AddressTableAddressExists(&_ArbSys.CallOpts, addr)
}

// AddressTableCompress is a free data retrieval call binding the contract method 0x2a890b98.
//
// Solidity: function addressTable_compress(address addr) view returns(bytes)
func (_ArbSys *ArbSysCaller) AddressTableCompress(opts *bind.CallOpts, addr common.Address) ([]byte, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "addressTable_compress", addr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AddressTableCompress is a free data retrieval call binding the contract method 0x2a890b98.
//
// Solidity: function addressTable_compress(address addr) view returns(bytes)
func (_ArbSys *ArbSysSession) AddressTableCompress(addr common.Address) ([]byte, error) {
	return _ArbSys.Contract.AddressTableCompress(&_ArbSys.CallOpts, addr)
}

// AddressTableCompress is a free data retrieval call binding the contract method 0x2a890b98.
//
// Solidity: function addressTable_compress(address addr) view returns(bytes)
func (_ArbSys *ArbSysCallerSession) AddressTableCompress(addr common.Address) ([]byte, error) {
	return _ArbSys.Contract.AddressTableCompress(&_ArbSys.CallOpts, addr)
}

// AddressTableDecompress is a free data retrieval call binding the contract method 0xbe05f2b9.
//
// Solidity: function addressTable_decompress(bytes buf, uint256 offset) view returns(address, uint256)
func (_ArbSys *ArbSysCaller) AddressTableDecompress(opts *bind.CallOpts, buf []byte, offset *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "addressTable_decompress", buf, offset)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// AddressTableDecompress is a free data retrieval call binding the contract method 0xbe05f2b9.
//
// Solidity: function addressTable_decompress(bytes buf, uint256 offset) view returns(address, uint256)
func (_ArbSys *ArbSysSession) AddressTableDecompress(buf []byte, offset *big.Int) (common.Address, *big.Int, error) {
	return _ArbSys.Contract.AddressTableDecompress(&_ArbSys.CallOpts, buf, offset)
}

// AddressTableDecompress is a free data retrieval call binding the contract method 0xbe05f2b9.
//
// Solidity: function addressTable_decompress(bytes buf, uint256 offset) view returns(address, uint256)
func (_ArbSys *ArbSysCallerSession) AddressTableDecompress(buf []byte, offset *big.Int) (common.Address, *big.Int, error) {
	return _ArbSys.Contract.AddressTableDecompress(&_ArbSys.CallOpts, buf, offset)
}

// AddressTableLookup is a free data retrieval call binding the contract method 0x8ee24222.
//
// Solidity: function addressTable_lookup(address addr) view returns(uint256)
func (_ArbSys *ArbSysCaller) AddressTableLookup(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "addressTable_lookup", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressTableLookup is a free data retrieval call binding the contract method 0x8ee24222.
//
// Solidity: function addressTable_lookup(address addr) view returns(uint256)
func (_ArbSys *ArbSysSession) AddressTableLookup(addr common.Address) (*big.Int, error) {
	return _ArbSys.Contract.AddressTableLookup(&_ArbSys.CallOpts, addr)
}

// AddressTableLookup is a free data retrieval call binding the contract method 0x8ee24222.
//
// Solidity: function addressTable_lookup(address addr) view returns(uint256)
func (_ArbSys *ArbSysCallerSession) AddressTableLookup(addr common.Address) (*big.Int, error) {
	return _ArbSys.Contract.AddressTableLookup(&_ArbSys.CallOpts, addr)
}

// AddressTableLookupIndex is a free data retrieval call binding the contract method 0x4c352200.
//
// Solidity: function addressTable_lookupIndex(uint256 index) view returns(address)
func (_ArbSys *ArbSysCaller) AddressTableLookupIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "addressTable_lookupIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressTableLookupIndex is a free data retrieval call binding the contract method 0x4c352200.
//
// Solidity: function addressTable_lookupIndex(uint256 index) view returns(address)
func (_ArbSys *ArbSysSession) AddressTableLookupIndex(index *big.Int) (common.Address, error) {
	return _ArbSys.Contract.AddressTableLookupIndex(&_ArbSys.CallOpts, index)
}

// AddressTableLookupIndex is a free data retrieval call binding the contract method 0x4c352200.
//
// Solidity: function addressTable_lookupIndex(uint256 index) view returns(address)
func (_ArbSys *ArbSysCallerSession) AddressTableLookupIndex(index *big.Int) (common.Address, error) {
	return _ArbSys.Contract.AddressTableLookupIndex(&_ArbSys.CallOpts, index)
}

// AddressTableSize is a free data retrieval call binding the contract method 0x03092fdb.
//
// Solidity: function addressTable_size() view returns(uint256)
func (_ArbSys *ArbSysCaller) AddressTableSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "addressTable_size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressTableSize is a free data retrieval call binding the contract method 0x03092fdb.
//
// Solidity: function addressTable_size() view returns(uint256)
func (_ArbSys *ArbSysSession) AddressTableSize() (*big.Int, error) {
	return _ArbSys.Contract.AddressTableSize(&_ArbSys.CallOpts)
}

// AddressTableSize is a free data retrieval call binding the contract method 0x03092fdb.
//
// Solidity: function addressTable_size() view returns(uint256)
func (_ArbSys *ArbSysCallerSession) AddressTableSize() (*big.Int, error) {
	return _ArbSys.Contract.AddressTableSize(&_ArbSys.CallOpts)
}

// FunctionTableGet is a free data retrieval call binding the contract method 0xd6ed0f56.
//
// Solidity: function functionTableGet(address addr, uint256 index) view returns(uint256, bool, uint256)
func (_ArbSys *ArbSysCaller) FunctionTableGet(opts *bind.CallOpts, addr common.Address, index *big.Int) (*big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "functionTableGet", addr, index)

	if err != nil {
		return *new(*big.Int), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// FunctionTableGet is a free data retrieval call binding the contract method 0xd6ed0f56.
//
// Solidity: function functionTableGet(address addr, uint256 index) view returns(uint256, bool, uint256)
func (_ArbSys *ArbSysSession) FunctionTableGet(addr common.Address, index *big.Int) (*big.Int, bool, *big.Int, error) {
	return _ArbSys.Contract.FunctionTableGet(&_ArbSys.CallOpts, addr, index)
}

// FunctionTableGet is a free data retrieval call binding the contract method 0xd6ed0f56.
//
// Solidity: function functionTableGet(address addr, uint256 index) view returns(uint256, bool, uint256)
func (_ArbSys *ArbSysCallerSession) FunctionTableGet(addr common.Address, index *big.Int) (*big.Int, bool, *big.Int, error) {
	return _ArbSys.Contract.FunctionTableGet(&_ArbSys.CallOpts, addr, index)
}

// FunctionTableSize is a free data retrieval call binding the contract method 0x1221c896.
//
// Solidity: function functionTableSize(address addr) view returns(uint256)
func (_ArbSys *ArbSysCaller) FunctionTableSize(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "functionTableSize", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FunctionTableSize is a free data retrieval call binding the contract method 0x1221c896.
//
// Solidity: function functionTableSize(address addr) view returns(uint256)
func (_ArbSys *ArbSysSession) FunctionTableSize(addr common.Address) (*big.Int, error) {
	return _ArbSys.Contract.FunctionTableSize(&_ArbSys.CallOpts, addr)
}

// FunctionTableSize is a free data retrieval call binding the contract method 0x1221c896.
//
// Solidity: function functionTableSize(address addr) view returns(uint256)
func (_ArbSys *ArbSysCallerSession) FunctionTableSize(addr common.Address) (*big.Int, error) {
	return _ArbSys.Contract.FunctionTableSize(&_ArbSys.CallOpts, addr)
}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0x647e0e98.
//
// Solidity: function getBlsPublicKey(address addr) view returns(uint256, uint256, uint256, uint256)
func (_ArbSys *ArbSysCaller) GetBlsPublicKey(opts *bind.CallOpts, addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "getBlsPublicKey", addr)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0x647e0e98.
//
// Solidity: function getBlsPublicKey(address addr) view returns(uint256, uint256, uint256, uint256)
func (_ArbSys *ArbSysSession) GetBlsPublicKey(addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbSys.Contract.GetBlsPublicKey(&_ArbSys.CallOpts, addr)
}

// GetBlsPublicKey is a free data retrieval call binding the contract method 0x647e0e98.
//
// Solidity: function getBlsPublicKey(address addr) view returns(uint256, uint256, uint256, uint256)
func (_ArbSys *ArbSysCallerSession) GetBlsPublicKey(addr common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbSys.Contract.GetBlsPublicKey(&_ArbSys.CallOpts, addr)
}

// GetStorageAt is a free data retrieval call binding the contract method 0xa169625f.
//
// Solidity: function getStorageAt(address account, uint256 index) view returns(uint256)
func (_ArbSys *ArbSysCaller) GetStorageAt(opts *bind.CallOpts, account common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "getStorageAt", account, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0xa169625f.
//
// Solidity: function getStorageAt(address account, uint256 index) view returns(uint256)
func (_ArbSys *ArbSysSession) GetStorageAt(account common.Address, index *big.Int) (*big.Int, error) {
	return _ArbSys.Contract.GetStorageAt(&_ArbSys.CallOpts, account, index)
}

// GetStorageAt is a free data retrieval call binding the contract method 0xa169625f.
//
// Solidity: function getStorageAt(address account, uint256 index) view returns(uint256)
func (_ArbSys *ArbSysCallerSession) GetStorageAt(account common.Address, index *big.Int) (*big.Int, error) {
	return _ArbSys.Contract.GetStorageAt(&_ArbSys.CallOpts, account, index)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x23ca0cd2.
//
// Solidity: function getTransactionCount(address account) view returns(uint256)
func (_ArbSys *ArbSysCaller) GetTransactionCount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "getTransactionCount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x23ca0cd2.
//
// Solidity: function getTransactionCount(address account) view returns(uint256)
func (_ArbSys *ArbSysSession) GetTransactionCount(account common.Address) (*big.Int, error) {
	return _ArbSys.Contract.GetTransactionCount(&_ArbSys.CallOpts, account)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x23ca0cd2.
//
// Solidity: function getTransactionCount(address account) view returns(uint256)
func (_ArbSys *ArbSysCallerSession) GetTransactionCount(account common.Address) (*big.Int, error) {
	return _ArbSys.Contract.GetTransactionCount(&_ArbSys.CallOpts, account)
}

// AddressTableRegister is a paid mutator transaction binding the contract method 0xaca1dc65.
//
// Solidity: function addressTable_register(address addr) returns(uint256)
func (_ArbSys *ArbSysTransactor) AddressTableRegister(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "addressTable_register", addr)
}

// AddressTableRegister is a paid mutator transaction binding the contract method 0xaca1dc65.
//
// Solidity: function addressTable_register(address addr) returns(uint256)
func (_ArbSys *ArbSysSession) AddressTableRegister(addr common.Address) (*types.Transaction, error) {
	return _ArbSys.Contract.AddressTableRegister(&_ArbSys.TransactOpts, addr)
}

// AddressTableRegister is a paid mutator transaction binding the contract method 0xaca1dc65.
//
// Solidity: function addressTable_register(address addr) returns(uint256)
func (_ArbSys *ArbSysTransactorSession) AddressTableRegister(addr common.Address) (*types.Transaction, error) {
	return _ArbSys.Contract.AddressTableRegister(&_ArbSys.TransactOpts, addr)
}

// RegisterBlsKey is a paid mutator transaction binding the contract method 0x1ff548ca.
//
// Solidity: function registerBlsKey(uint256 x0, uint256 x1, uint256 y0, uint256 y1) returns()
func (_ArbSys *ArbSysTransactor) RegisterBlsKey(opts *bind.TransactOpts, x0 *big.Int, x1 *big.Int, y0 *big.Int, y1 *big.Int) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "registerBlsKey", x0, x1, y0, y1)
}

// RegisterBlsKey is a paid mutator transaction binding the contract method 0x1ff548ca.
//
// Solidity: function registerBlsKey(uint256 x0, uint256 x1, uint256 y0, uint256 y1) returns()
func (_ArbSys *ArbSysSession) RegisterBlsKey(x0 *big.Int, x1 *big.Int, y0 *big.Int, y1 *big.Int) (*types.Transaction, error) {
	return _ArbSys.Contract.RegisterBlsKey(&_ArbSys.TransactOpts, x0, x1, y0, y1)
}

// RegisterBlsKey is a paid mutator transaction binding the contract method 0x1ff548ca.
//
// Solidity: function registerBlsKey(uint256 x0, uint256 x1, uint256 y0, uint256 y1) returns()
func (_ArbSys *ArbSysTransactorSession) RegisterBlsKey(x0 *big.Int, x1 *big.Int, y0 *big.Int, y1 *big.Int) (*types.Transaction, error) {
	return _ArbSys.Contract.RegisterBlsKey(&_ArbSys.TransactOpts, x0, x1, y0, y1)
}

// UploadFunctionTable is a paid mutator transaction binding the contract method 0xb7cf9527.
//
// Solidity: function uploadFunctionTable(bytes buf) returns()
func (_ArbSys *ArbSysTransactor) UploadFunctionTable(opts *bind.TransactOpts, buf []byte) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "uploadFunctionTable", buf)
}

// UploadFunctionTable is a paid mutator transaction binding the contract method 0xb7cf9527.
//
// Solidity: function uploadFunctionTable(bytes buf) returns()
func (_ArbSys *ArbSysSession) UploadFunctionTable(buf []byte) (*types.Transaction, error) {
	return _ArbSys.Contract.UploadFunctionTable(&_ArbSys.TransactOpts, buf)
}

// UploadFunctionTable is a paid mutator transaction binding the contract method 0xb7cf9527.
//
// Solidity: function uploadFunctionTable(bytes buf) returns()
func (_ArbSys *ArbSysTransactorSession) UploadFunctionTable(buf []byte) (*types.Transaction, error) {
	return _ArbSys.Contract.UploadFunctionTable(&_ArbSys.TransactOpts, buf)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address dest, uint256 amount) returns()
func (_ArbSys *ArbSysTransactor) WithdrawERC20(opts *bind.TransactOpts, dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "withdrawERC20", dest, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address dest, uint256 amount) returns()
func (_ArbSys *ArbSysSession) WithdrawERC20(dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawERC20(&_ArbSys.TransactOpts, dest, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address dest, uint256 amount) returns()
func (_ArbSys *ArbSysTransactorSession) WithdrawERC20(dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawERC20(&_ArbSys.TransactOpts, dest, amount)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address dest, uint256 id) returns()
func (_ArbSys *ArbSysTransactor) WithdrawERC721(opts *bind.TransactOpts, dest common.Address, id *big.Int) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "withdrawERC721", dest, id)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address dest, uint256 id) returns()
func (_ArbSys *ArbSysSession) WithdrawERC721(dest common.Address, id *big.Int) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawERC721(&_ArbSys.TransactOpts, dest, id)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address dest, uint256 id) returns()
func (_ArbSys *ArbSysTransactorSession) WithdrawERC721(dest common.Address, id *big.Int) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawERC721(&_ArbSys.TransactOpts, dest, id)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address dest) payable returns()
func (_ArbSys *ArbSysTransactor) WithdrawEth(opts *bind.TransactOpts, dest common.Address) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "withdrawEth", dest)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address dest) payable returns()
func (_ArbSys *ArbSysSession) WithdrawEth(dest common.Address) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawEth(&_ArbSys.TransactOpts, dest)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address dest) payable returns()
func (_ArbSys *ArbSysTransactorSession) WithdrawEth(dest common.Address) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawEth(&_ArbSys.TransactOpts, dest)
}

// ArbSysERC20WithdrawalIterator is returned from FilterERC20Withdrawal and is used to iterate over the raw logs and unpacked data for ERC20Withdrawal events raised by the ArbSys contract.
type ArbSysERC20WithdrawalIterator struct {
	Event *ArbSysERC20Withdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbSysERC20WithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbSysERC20Withdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbSysERC20Withdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbSysERC20WithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbSysERC20WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbSysERC20Withdrawal represents a ERC20Withdrawal event raised by the ArbSys contract.
type ArbSysERC20Withdrawal struct {
	DestAddr  common.Address
	TokenAddr common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Withdrawal is a free log retrieval operation binding the contract event 0x2fbb3e8dc2807d6a61feb98fc2120153f77b2c0d25ef6272e1756935dd62c847.
//
// Solidity: event ERC20Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 amount)
func (_ArbSys *ArbSysFilterer) FilterERC20Withdrawal(opts *bind.FilterOpts, destAddr []common.Address, tokenAddr []common.Address) (*ArbSysERC20WithdrawalIterator, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _ArbSys.contract.FilterLogs(opts, "ERC20Withdrawal", destAddrRule, tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &ArbSysERC20WithdrawalIterator{contract: _ArbSys.contract, event: "ERC20Withdrawal", logs: logs, sub: sub}, nil
}

// WatchERC20Withdrawal is a free log subscription operation binding the contract event 0x2fbb3e8dc2807d6a61feb98fc2120153f77b2c0d25ef6272e1756935dd62c847.
//
// Solidity: event ERC20Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 amount)
func (_ArbSys *ArbSysFilterer) WatchERC20Withdrawal(opts *bind.WatchOpts, sink chan<- *ArbSysERC20Withdrawal, destAddr []common.Address, tokenAddr []common.Address) (event.Subscription, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _ArbSys.contract.WatchLogs(opts, "ERC20Withdrawal", destAddrRule, tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbSysERC20Withdrawal)
				if err := _ArbSys.contract.UnpackLog(event, "ERC20Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20Withdrawal is a log parse operation binding the contract event 0x2fbb3e8dc2807d6a61feb98fc2120153f77b2c0d25ef6272e1756935dd62c847.
//
// Solidity: event ERC20Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 amount)
func (_ArbSys *ArbSysFilterer) ParseERC20Withdrawal(log types.Log) (*ArbSysERC20Withdrawal, error) {
	event := new(ArbSysERC20Withdrawal)
	if err := _ArbSys.contract.UnpackLog(event, "ERC20Withdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbSysERC721WithdrawalIterator is returned from FilterERC721Withdrawal and is used to iterate over the raw logs and unpacked data for ERC721Withdrawal events raised by the ArbSys contract.
type ArbSysERC721WithdrawalIterator struct {
	Event *ArbSysERC721Withdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbSysERC721WithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbSysERC721Withdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbSysERC721Withdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbSysERC721WithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbSysERC721WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbSysERC721Withdrawal represents a ERC721Withdrawal event raised by the ArbSys contract.
type ArbSysERC721Withdrawal struct {
	DestAddr  common.Address
	TokenAddr common.Address
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Withdrawal is a free log retrieval operation binding the contract event 0xdce1abc7607cf83c96953308072c045ec35d02fbf5777e7d0ec8b102cd89ff81.
//
// Solidity: event ERC721Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 indexed id)
func (_ArbSys *ArbSysFilterer) FilterERC721Withdrawal(opts *bind.FilterOpts, destAddr []common.Address, tokenAddr []common.Address, id []*big.Int) (*ArbSysERC721WithdrawalIterator, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ArbSys.contract.FilterLogs(opts, "ERC721Withdrawal", destAddrRule, tokenAddrRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ArbSysERC721WithdrawalIterator{contract: _ArbSys.contract, event: "ERC721Withdrawal", logs: logs, sub: sub}, nil
}

// WatchERC721Withdrawal is a free log subscription operation binding the contract event 0xdce1abc7607cf83c96953308072c045ec35d02fbf5777e7d0ec8b102cd89ff81.
//
// Solidity: event ERC721Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 indexed id)
func (_ArbSys *ArbSysFilterer) WatchERC721Withdrawal(opts *bind.WatchOpts, sink chan<- *ArbSysERC721Withdrawal, destAddr []common.Address, tokenAddr []common.Address, id []*big.Int) (event.Subscription, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ArbSys.contract.WatchLogs(opts, "ERC721Withdrawal", destAddrRule, tokenAddrRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbSysERC721Withdrawal)
				if err := _ArbSys.contract.UnpackLog(event, "ERC721Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC721Withdrawal is a log parse operation binding the contract event 0xdce1abc7607cf83c96953308072c045ec35d02fbf5777e7d0ec8b102cd89ff81.
//
// Solidity: event ERC721Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 indexed id)
func (_ArbSys *ArbSysFilterer) ParseERC721Withdrawal(log types.Log) (*ArbSysERC721Withdrawal, error) {
	event := new(ArbSysERC721Withdrawal)
	if err := _ArbSys.contract.UnpackLog(event, "ERC721Withdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbSysEthWithdrawalIterator is returned from FilterEthWithdrawal and is used to iterate over the raw logs and unpacked data for EthWithdrawal events raised by the ArbSys contract.
type ArbSysEthWithdrawalIterator struct {
	Event *ArbSysEthWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbSysEthWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbSysEthWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbSysEthWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbSysEthWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbSysEthWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbSysEthWithdrawal represents a EthWithdrawal event raised by the ArbSys contract.
type ArbSysEthWithdrawal struct {
	DestAddr common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawal is a free log retrieval operation binding the contract event 0xc32d3c7eb0f275cbb5b72b3d3c688269430f30e5b9bb36980396edd9101c615c.
//
// Solidity: event EthWithdrawal(address indexed destAddr, uint256 amount)
func (_ArbSys *ArbSysFilterer) FilterEthWithdrawal(opts *bind.FilterOpts, destAddr []common.Address) (*ArbSysEthWithdrawalIterator, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}

	logs, sub, err := _ArbSys.contract.FilterLogs(opts, "EthWithdrawal", destAddrRule)
	if err != nil {
		return nil, err
	}
	return &ArbSysEthWithdrawalIterator{contract: _ArbSys.contract, event: "EthWithdrawal", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawal is a free log subscription operation binding the contract event 0xc32d3c7eb0f275cbb5b72b3d3c688269430f30e5b9bb36980396edd9101c615c.
//
// Solidity: event EthWithdrawal(address indexed destAddr, uint256 amount)
func (_ArbSys *ArbSysFilterer) WatchEthWithdrawal(opts *bind.WatchOpts, sink chan<- *ArbSysEthWithdrawal, destAddr []common.Address) (event.Subscription, error) {

	var destAddrRule []interface{}
	for _, destAddrItem := range destAddr {
		destAddrRule = append(destAddrRule, destAddrItem)
	}

	logs, sub, err := _ArbSys.contract.WatchLogs(opts, "EthWithdrawal", destAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbSysEthWithdrawal)
				if err := _ArbSys.contract.UnpackLog(event, "EthWithdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthWithdrawal is a log parse operation binding the contract event 0xc32d3c7eb0f275cbb5b72b3d3c688269430f30e5b9bb36980396edd9101c615c.
//
// Solidity: event EthWithdrawal(address indexed destAddr, uint256 amount)
func (_ArbSys *ArbSysFilterer) ParseEthWithdrawal(log types.Log) (*ArbSysEthWithdrawal, error) {
	event := new(ArbSysEthWithdrawal)
	if err := _ArbSys.contract.UnpackLog(event, "EthWithdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}
