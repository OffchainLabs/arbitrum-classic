// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goarbitrum

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
const ArbSysABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"blockUpperBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cloneContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentMessageBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentMessageTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timestampUpperBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbSysFuncSigs maps the 4-byte function signature to its string representation.
var ArbSysFuncSigs = map[string]string{
	"4baa4a24": "blockUpperBound()",
	"474ed9c0": "cloneContract(address)",
	"21151d8d": "currentMessageBlock()",
	"f1362091": "currentMessageTimestamp()",
	"23ca0cd2": "getTransactionCount(address)",
	"f5e71ccd": "timestampUpperBound()",
	"a1db9782": "withdrawERC20(address,uint256)",
	"f3e414f8": "withdrawERC721(address,uint256)",
	"1b9a91a4": "withdrawEth(address,uint256)",
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
func (_ArbSys *ArbSysRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ArbSys *ArbSysCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// BlockUpperBound is a free data retrieval call binding the contract method 0x4baa4a24.
//
// Solidity: function blockUpperBound() view returns(uint256)
func (_ArbSys *ArbSysCaller) BlockUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbSys.contract.Call(opts, out, "blockUpperBound")
	return *ret0, err
}

// BlockUpperBound is a free data retrieval call binding the contract method 0x4baa4a24.
//
// Solidity: function blockUpperBound() view returns(uint256)
func (_ArbSys *ArbSysSession) BlockUpperBound() (*big.Int, error) {
	return _ArbSys.Contract.BlockUpperBound(&_ArbSys.CallOpts)
}

// BlockUpperBound is a free data retrieval call binding the contract method 0x4baa4a24.
//
// Solidity: function blockUpperBound() view returns(uint256)
func (_ArbSys *ArbSysCallerSession) BlockUpperBound() (*big.Int, error) {
	return _ArbSys.Contract.BlockUpperBound(&_ArbSys.CallOpts)
}

// CurrentMessageBlock is a free data retrieval call binding the contract method 0x21151d8d.
//
// Solidity: function currentMessageBlock() view returns(uint256)
func (_ArbSys *ArbSysCaller) CurrentMessageBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbSys.contract.Call(opts, out, "currentMessageBlock")
	return *ret0, err
}

// CurrentMessageBlock is a free data retrieval call binding the contract method 0x21151d8d.
//
// Solidity: function currentMessageBlock() view returns(uint256)
func (_ArbSys *ArbSysSession) CurrentMessageBlock() (*big.Int, error) {
	return _ArbSys.Contract.CurrentMessageBlock(&_ArbSys.CallOpts)
}

// CurrentMessageBlock is a free data retrieval call binding the contract method 0x21151d8d.
//
// Solidity: function currentMessageBlock() view returns(uint256)
func (_ArbSys *ArbSysCallerSession) CurrentMessageBlock() (*big.Int, error) {
	return _ArbSys.Contract.CurrentMessageBlock(&_ArbSys.CallOpts)
}

// CurrentMessageTimestamp is a free data retrieval call binding the contract method 0xf1362091.
//
// Solidity: function currentMessageTimestamp() view returns(uint256)
func (_ArbSys *ArbSysCaller) CurrentMessageTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbSys.contract.Call(opts, out, "currentMessageTimestamp")
	return *ret0, err
}

// CurrentMessageTimestamp is a free data retrieval call binding the contract method 0xf1362091.
//
// Solidity: function currentMessageTimestamp() view returns(uint256)
func (_ArbSys *ArbSysSession) CurrentMessageTimestamp() (*big.Int, error) {
	return _ArbSys.Contract.CurrentMessageTimestamp(&_ArbSys.CallOpts)
}

// CurrentMessageTimestamp is a free data retrieval call binding the contract method 0xf1362091.
//
// Solidity: function currentMessageTimestamp() view returns(uint256)
func (_ArbSys *ArbSysCallerSession) CurrentMessageTimestamp() (*big.Int, error) {
	return _ArbSys.Contract.CurrentMessageTimestamp(&_ArbSys.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x23ca0cd2.
//
// Solidity: function getTransactionCount(address account) view returns(uint256)
func (_ArbSys *ArbSysCaller) GetTransactionCount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbSys.contract.Call(opts, out, "getTransactionCount", account)
	return *ret0, err
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

// TimestampUpperBound is a free data retrieval call binding the contract method 0xf5e71ccd.
//
// Solidity: function timestampUpperBound() view returns(uint256)
func (_ArbSys *ArbSysCaller) TimestampUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbSys.contract.Call(opts, out, "timestampUpperBound")
	return *ret0, err
}

// TimestampUpperBound is a free data retrieval call binding the contract method 0xf5e71ccd.
//
// Solidity: function timestampUpperBound() view returns(uint256)
func (_ArbSys *ArbSysSession) TimestampUpperBound() (*big.Int, error) {
	return _ArbSys.Contract.TimestampUpperBound(&_ArbSys.CallOpts)
}

// TimestampUpperBound is a free data retrieval call binding the contract method 0xf5e71ccd.
//
// Solidity: function timestampUpperBound() view returns(uint256)
func (_ArbSys *ArbSysCallerSession) TimestampUpperBound() (*big.Int, error) {
	return _ArbSys.Contract.TimestampUpperBound(&_ArbSys.CallOpts)
}

// CloneContract is a paid mutator transaction binding the contract method 0x474ed9c0.
//
// Solidity: function cloneContract(address account) returns(address)
func (_ArbSys *ArbSysTransactor) CloneContract(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "cloneContract", account)
}

// CloneContract is a paid mutator transaction binding the contract method 0x474ed9c0.
//
// Solidity: function cloneContract(address account) returns(address)
func (_ArbSys *ArbSysSession) CloneContract(account common.Address) (*types.Transaction, error) {
	return _ArbSys.Contract.CloneContract(&_ArbSys.TransactOpts, account)
}

// CloneContract is a paid mutator transaction binding the contract method 0x474ed9c0.
//
// Solidity: function cloneContract(address account) returns(address)
func (_ArbSys *ArbSysTransactorSession) CloneContract(account common.Address) (*types.Transaction, error) {
	return _ArbSys.Contract.CloneContract(&_ArbSys.TransactOpts, account)
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

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address dest, uint256 amount) returns()
func (_ArbSys *ArbSysTransactor) WithdrawEth(opts *bind.TransactOpts, dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "withdrawEth", dest, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address dest, uint256 amount) returns()
func (_ArbSys *ArbSysSession) WithdrawEth(dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawEth(&_ArbSys.TransactOpts, dest, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address dest, uint256 amount) returns()
func (_ArbSys *ArbSysTransactorSession) WithdrawEth(dest common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawEth(&_ArbSys.TransactOpts, dest, amount)
}
