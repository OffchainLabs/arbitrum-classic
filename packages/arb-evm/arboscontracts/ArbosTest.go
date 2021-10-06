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

// ArbosTestMetaData contains all meta data concerning the ArbosTest contract.
var ArbosTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"}],\"name\":\"burnArbGas\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAccountInfo\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMarshalledStorage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isEOA\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"code\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"initStorage\",\"type\":\"bytes\"}],\"name\":\"installAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"code\",\"type\":\"bytes\"}],\"name\":\"setCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ArbosTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbosTestMetaData.ABI instead.
var ArbosTestABI = ArbosTestMetaData.ABI

// ArbosTest is an auto generated Go binding around an Ethereum contract.
type ArbosTest struct {
	ArbosTestCaller     // Read-only binding to the contract
	ArbosTestTransactor // Write-only binding to the contract
	ArbosTestFilterer   // Log filterer for contract events
}

// ArbosTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbosTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbosTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbosTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbosTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbosTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbosTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbosTestSession struct {
	Contract     *ArbosTest        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbosTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbosTestCallerSession struct {
	Contract *ArbosTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArbosTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbosTestTransactorSession struct {
	Contract     *ArbosTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArbosTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbosTestRaw struct {
	Contract *ArbosTest // Generic contract binding to access the raw methods on
}

// ArbosTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbosTestCallerRaw struct {
	Contract *ArbosTestCaller // Generic read-only contract binding to access the raw methods on
}

// ArbosTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbosTestTransactorRaw struct {
	Contract *ArbosTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbosTest creates a new instance of ArbosTest, bound to a specific deployed contract.
func NewArbosTest(address common.Address, backend bind.ContractBackend) (*ArbosTest, error) {
	contract, err := bindArbosTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbosTest{ArbosTestCaller: ArbosTestCaller{contract: contract}, ArbosTestTransactor: ArbosTestTransactor{contract: contract}, ArbosTestFilterer: ArbosTestFilterer{contract: contract}}, nil
}

// NewArbosTestCaller creates a new read-only instance of ArbosTest, bound to a specific deployed contract.
func NewArbosTestCaller(address common.Address, caller bind.ContractCaller) (*ArbosTestCaller, error) {
	contract, err := bindArbosTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbosTestCaller{contract: contract}, nil
}

// NewArbosTestTransactor creates a new write-only instance of ArbosTest, bound to a specific deployed contract.
func NewArbosTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbosTestTransactor, error) {
	contract, err := bindArbosTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbosTestTransactor{contract: contract}, nil
}

// NewArbosTestFilterer creates a new log filterer instance of ArbosTest, bound to a specific deployed contract.
func NewArbosTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbosTestFilterer, error) {
	contract, err := bindArbosTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbosTestFilterer{contract: contract}, nil
}

// bindArbosTest binds a generic wrapper to an already deployed contract.
func bindArbosTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbosTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbosTest *ArbosTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbosTest.Contract.ArbosTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbosTest *ArbosTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbosTest.Contract.ArbosTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbosTest *ArbosTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbosTest.Contract.ArbosTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbosTest *ArbosTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbosTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbosTest *ArbosTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbosTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbosTest *ArbosTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbosTest.Contract.contract.Transact(opts, method, params...)
}

// BurnArbGas is a free data retrieval call binding the contract method 0xbb3480f9.
//
// Solidity: function burnArbGas(uint256 gasAmount) view returns()
func (_ArbosTest *ArbosTestCaller) BurnArbGas(opts *bind.CallOpts, gasAmount *big.Int) error {
	var out []interface{}
	err := _ArbosTest.contract.Call(opts, &out, "burnArbGas", gasAmount)

	if err != nil {
		return err
	}

	return err

}

// BurnArbGas is a free data retrieval call binding the contract method 0xbb3480f9.
//
// Solidity: function burnArbGas(uint256 gasAmount) view returns()
func (_ArbosTest *ArbosTestSession) BurnArbGas(gasAmount *big.Int) error {
	return _ArbosTest.Contract.BurnArbGas(&_ArbosTest.CallOpts, gasAmount)
}

// BurnArbGas is a free data retrieval call binding the contract method 0xbb3480f9.
//
// Solidity: function burnArbGas(uint256 gasAmount) view returns()
func (_ArbosTest *ArbosTestCallerSession) BurnArbGas(gasAmount *big.Int) error {
	return _ArbosTest.Contract.BurnArbGas(&_ArbosTest.CallOpts, gasAmount)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address addr) view returns()
func (_ArbosTest *ArbosTestCaller) GetAccountInfo(opts *bind.CallOpts, addr common.Address) error {
	var out []interface{}
	err := _ArbosTest.contract.Call(opts, &out, "getAccountInfo", addr)

	if err != nil {
		return err
	}

	return err

}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address addr) view returns()
func (_ArbosTest *ArbosTestSession) GetAccountInfo(addr common.Address) error {
	return _ArbosTest.Contract.GetAccountInfo(&_ArbosTest.CallOpts, addr)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address addr) view returns()
func (_ArbosTest *ArbosTestCallerSession) GetAccountInfo(addr common.Address) error {
	return _ArbosTest.Contract.GetAccountInfo(&_ArbosTest.CallOpts, addr)
}

// GetMarshalledStorage is a free data retrieval call binding the contract method 0xd56aa31f.
//
// Solidity: function getMarshalledStorage(address addr) view returns()
func (_ArbosTest *ArbosTestCaller) GetMarshalledStorage(opts *bind.CallOpts, addr common.Address) error {
	var out []interface{}
	err := _ArbosTest.contract.Call(opts, &out, "getMarshalledStorage", addr)

	if err != nil {
		return err
	}

	return err

}

// GetMarshalledStorage is a free data retrieval call binding the contract method 0xd56aa31f.
//
// Solidity: function getMarshalledStorage(address addr) view returns()
func (_ArbosTest *ArbosTestSession) GetMarshalledStorage(addr common.Address) error {
	return _ArbosTest.Contract.GetMarshalledStorage(&_ArbosTest.CallOpts, addr)
}

// GetMarshalledStorage is a free data retrieval call binding the contract method 0xd56aa31f.
//
// Solidity: function getMarshalledStorage(address addr) view returns()
func (_ArbosTest *ArbosTestCallerSession) GetMarshalledStorage(addr common.Address) error {
	return _ArbosTest.Contract.GetMarshalledStorage(&_ArbosTest.CallOpts, addr)
}

// InstallAccount is a paid mutator transaction binding the contract method 0xfbe6e022.
//
// Solidity: function installAccount(address addr, bool isEOA, uint256 balance, uint256 nonce, bytes code, bytes initStorage) returns()
func (_ArbosTest *ArbosTestTransactor) InstallAccount(opts *bind.TransactOpts, addr common.Address, isEOA bool, balance *big.Int, nonce *big.Int, code []byte, initStorage []byte) (*types.Transaction, error) {
	return _ArbosTest.contract.Transact(opts, "installAccount", addr, isEOA, balance, nonce, code, initStorage)
}

// InstallAccount is a paid mutator transaction binding the contract method 0xfbe6e022.
//
// Solidity: function installAccount(address addr, bool isEOA, uint256 balance, uint256 nonce, bytes code, bytes initStorage) returns()
func (_ArbosTest *ArbosTestSession) InstallAccount(addr common.Address, isEOA bool, balance *big.Int, nonce *big.Int, code []byte, initStorage []byte) (*types.Transaction, error) {
	return _ArbosTest.Contract.InstallAccount(&_ArbosTest.TransactOpts, addr, isEOA, balance, nonce, code, initStorage)
}

// InstallAccount is a paid mutator transaction binding the contract method 0xfbe6e022.
//
// Solidity: function installAccount(address addr, bool isEOA, uint256 balance, uint256 nonce, bytes code, bytes initStorage) returns()
func (_ArbosTest *ArbosTestTransactorSession) InstallAccount(addr common.Address, isEOA bool, balance *big.Int, nonce *big.Int, code []byte, initStorage []byte) (*types.Transaction, error) {
	return _ArbosTest.Contract.InstallAccount(&_ArbosTest.TransactOpts, addr, isEOA, balance, nonce, code, initStorage)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address addr, uint256 balance) returns()
func (_ArbosTest *ArbosTestTransactor) SetBalance(opts *bind.TransactOpts, addr common.Address, balance *big.Int) (*types.Transaction, error) {
	return _ArbosTest.contract.Transact(opts, "setBalance", addr, balance)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address addr, uint256 balance) returns()
func (_ArbosTest *ArbosTestSession) SetBalance(addr common.Address, balance *big.Int) (*types.Transaction, error) {
	return _ArbosTest.Contract.SetBalance(&_ArbosTest.TransactOpts, addr, balance)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address addr, uint256 balance) returns()
func (_ArbosTest *ArbosTestTransactorSession) SetBalance(addr common.Address, balance *big.Int) (*types.Transaction, error) {
	return _ArbosTest.Contract.SetBalance(&_ArbosTest.TransactOpts, addr, balance)
}

// SetCode is a paid mutator transaction binding the contract method 0x978190f1.
//
// Solidity: function setCode(address addr, bytes code) returns()
func (_ArbosTest *ArbosTestTransactor) SetCode(opts *bind.TransactOpts, addr common.Address, code []byte) (*types.Transaction, error) {
	return _ArbosTest.contract.Transact(opts, "setCode", addr, code)
}

// SetCode is a paid mutator transaction binding the contract method 0x978190f1.
//
// Solidity: function setCode(address addr, bytes code) returns()
func (_ArbosTest *ArbosTestSession) SetCode(addr common.Address, code []byte) (*types.Transaction, error) {
	return _ArbosTest.Contract.SetCode(&_ArbosTest.TransactOpts, addr, code)
}

// SetCode is a paid mutator transaction binding the contract method 0x978190f1.
//
// Solidity: function setCode(address addr, bytes code) returns()
func (_ArbosTest *ArbosTestTransactorSession) SetCode(addr common.Address, code []byte) (*types.Transaction, error) {
	return _ArbosTest.Contract.SetCode(&_ArbosTest.TransactOpts, addr, code)
}

// SetNonce is a paid mutator transaction binding the contract method 0x1d79f325.
//
// Solidity: function setNonce(address addr, uint256 nonce) returns()
func (_ArbosTest *ArbosTestTransactor) SetNonce(opts *bind.TransactOpts, addr common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _ArbosTest.contract.Transact(opts, "setNonce", addr, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0x1d79f325.
//
// Solidity: function setNonce(address addr, uint256 nonce) returns()
func (_ArbosTest *ArbosTestSession) SetNonce(addr common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _ArbosTest.Contract.SetNonce(&_ArbosTest.TransactOpts, addr, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0x1d79f325.
//
// Solidity: function setNonce(address addr, uint256 nonce) returns()
func (_ArbosTest *ArbosTestTransactorSession) SetNonce(addr common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _ArbosTest.Contract.SetNonce(&_ArbosTest.TransactOpts, addr, nonce)
}

// SetState is a paid mutator transaction binding the contract method 0xafefabf7.
//
// Solidity: function setState(address addr, bytes state) returns()
func (_ArbosTest *ArbosTestTransactor) SetState(opts *bind.TransactOpts, addr common.Address, state []byte) (*types.Transaction, error) {
	return _ArbosTest.contract.Transact(opts, "setState", addr, state)
}

// SetState is a paid mutator transaction binding the contract method 0xafefabf7.
//
// Solidity: function setState(address addr, bytes state) returns()
func (_ArbosTest *ArbosTestSession) SetState(addr common.Address, state []byte) (*types.Transaction, error) {
	return _ArbosTest.Contract.SetState(&_ArbosTest.TransactOpts, addr, state)
}

// SetState is a paid mutator transaction binding the contract method 0xafefabf7.
//
// Solidity: function setState(address addr, bytes state) returns()
func (_ArbosTest *ArbosTestTransactorSession) SetState(addr common.Address, state []byte) (*types.Transaction, error) {
	return _ArbosTest.Contract.SetState(&_ArbosTest.TransactOpts, addr, state)
}

// Store is a paid mutator transaction binding the contract method 0xf0c26aff.
//
// Solidity: function store(address addr, uint256 key, uint256 value) returns()
func (_ArbosTest *ArbosTestTransactor) Store(opts *bind.TransactOpts, addr common.Address, key *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ArbosTest.contract.Transact(opts, "store", addr, key, value)
}

// Store is a paid mutator transaction binding the contract method 0xf0c26aff.
//
// Solidity: function store(address addr, uint256 key, uint256 value) returns()
func (_ArbosTest *ArbosTestSession) Store(addr common.Address, key *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ArbosTest.Contract.Store(&_ArbosTest.TransactOpts, addr, key, value)
}

// Store is a paid mutator transaction binding the contract method 0xf0c26aff.
//
// Solidity: function store(address addr, uint256 key, uint256 value) returns()
func (_ArbosTest *ArbosTestTransactorSession) Store(addr common.Address, key *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ArbosTest.Contract.Store(&_ArbosTest.TransactOpts, addr, key, value)
}
