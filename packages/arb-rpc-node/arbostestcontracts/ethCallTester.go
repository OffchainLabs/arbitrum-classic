// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbostestcontracts

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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"}],\"name\":\"burnArbGas\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAccountInfo\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMarshalledStorage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isEOA\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"code\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"initStorage\",\"type\":\"bytes\"}],\"name\":\"installAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"code\",\"type\":\"bytes\"}],\"name\":\"setCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"key\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Store is a paid mutator transaction binding the contract method 0x2e619f57.
//
// Solidity: function store(address addr, uint32 key, uint32 value) returns()
func (_ArbosTest *ArbosTestTransactor) Store(opts *bind.TransactOpts, addr common.Address, key uint32, value uint32) (*types.Transaction, error) {
	return _ArbosTest.contract.Transact(opts, "store", addr, key, value)
}

// Store is a paid mutator transaction binding the contract method 0x2e619f57.
//
// Solidity: function store(address addr, uint32 key, uint32 value) returns()
func (_ArbosTest *ArbosTestSession) Store(addr common.Address, key uint32, value uint32) (*types.Transaction, error) {
	return _ArbosTest.Contract.Store(&_ArbosTest.TransactOpts, addr, key, value)
}

// Store is a paid mutator transaction binding the contract method 0x2e619f57.
//
// Solidity: function store(address addr, uint32 key, uint32 value) returns()
func (_ArbosTest *ArbosTestTransactorSession) Store(addr common.Address, key uint32, value uint32) (*types.Transaction, error) {
	return _ArbosTest.Contract.Store(&_ArbosTest.TransactOpts, addr, key, value)
}

// EthCallTesterMetaData contains all meta data concerning the EthCallTester contract.
var EthCallTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"failStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"sLoad\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610100600055610146806100266000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806312065fe01461005157806330874d8d1461006b5780635197c7aa146100885780638bb204a114610090575b600080fd5b61005961009a565b60408051918252519081900360200190f35b6100596004803603602081101561008157600080fd5b503561009e565b6100596100a2565b6100986100a8565b005b4790565b5490565b60005490565b60408051632e619f5760e01b8152306004820152600060248201819052630300000060448301529151606992632e619f57926064808201939182900301818387803b1580156100f657600080fd5b505af115801561010a573d6000803e3d6000fd5b5050505056fea264697066735822122089b7777e8ef15ddea1bcf79a6890c4381f424905b7832dfa58dc66dde08a19b764736f6c634300060c0033",
}

// EthCallTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use EthCallTesterMetaData.ABI instead.
var EthCallTesterABI = EthCallTesterMetaData.ABI

// EthCallTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthCallTesterMetaData.Bin instead.
var EthCallTesterBin = EthCallTesterMetaData.Bin

// DeployEthCallTester deploys a new Ethereum contract, binding an instance of EthCallTester to it.
func DeployEthCallTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthCallTester, error) {
	parsed, err := EthCallTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthCallTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCallTester{EthCallTesterCaller: EthCallTesterCaller{contract: contract}, EthCallTesterTransactor: EthCallTesterTransactor{contract: contract}, EthCallTesterFilterer: EthCallTesterFilterer{contract: contract}}, nil
}

// EthCallTester is an auto generated Go binding around an Ethereum contract.
type EthCallTester struct {
	EthCallTesterCaller     // Read-only binding to the contract
	EthCallTesterTransactor // Write-only binding to the contract
	EthCallTesterFilterer   // Log filterer for contract events
}

// EthCallTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCallTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCallTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCallTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCallTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCallTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCallTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCallTesterSession struct {
	Contract     *EthCallTester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthCallTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCallTesterCallerSession struct {
	Contract *EthCallTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EthCallTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCallTesterTransactorSession struct {
	Contract     *EthCallTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EthCallTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCallTesterRaw struct {
	Contract *EthCallTester // Generic contract binding to access the raw methods on
}

// EthCallTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCallTesterCallerRaw struct {
	Contract *EthCallTesterCaller // Generic read-only contract binding to access the raw methods on
}

// EthCallTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCallTesterTransactorRaw struct {
	Contract *EthCallTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCallTester creates a new instance of EthCallTester, bound to a specific deployed contract.
func NewEthCallTester(address common.Address, backend bind.ContractBackend) (*EthCallTester, error) {
	contract, err := bindEthCallTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCallTester{EthCallTesterCaller: EthCallTesterCaller{contract: contract}, EthCallTesterTransactor: EthCallTesterTransactor{contract: contract}, EthCallTesterFilterer: EthCallTesterFilterer{contract: contract}}, nil
}

// NewEthCallTesterCaller creates a new read-only instance of EthCallTester, bound to a specific deployed contract.
func NewEthCallTesterCaller(address common.Address, caller bind.ContractCaller) (*EthCallTesterCaller, error) {
	contract, err := bindEthCallTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCallTesterCaller{contract: contract}, nil
}

// NewEthCallTesterTransactor creates a new write-only instance of EthCallTester, bound to a specific deployed contract.
func NewEthCallTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCallTesterTransactor, error) {
	contract, err := bindEthCallTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCallTesterTransactor{contract: contract}, nil
}

// NewEthCallTesterFilterer creates a new log filterer instance of EthCallTester, bound to a specific deployed contract.
func NewEthCallTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCallTesterFilterer, error) {
	contract, err := bindEthCallTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCallTesterFilterer{contract: contract}, nil
}

// bindEthCallTester binds a generic wrapper to an already deployed contract.
func bindEthCallTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCallTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCallTester *EthCallTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthCallTester.Contract.EthCallTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCallTester *EthCallTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCallTester.Contract.EthCallTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCallTester *EthCallTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCallTester.Contract.EthCallTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCallTester *EthCallTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthCallTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCallTester *EthCallTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCallTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCallTester *EthCallTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCallTester.Contract.contract.Transact(opts, method, params...)
}

// FailStore is a paid mutator transaction binding the contract method 0x8bb204a1.
//
// Solidity: function failStore() returns()
func (_EthCallTester *EthCallTesterTransactor) FailStore(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCallTester.contract.Transact(opts, "failStore")
}

// FailStore is a paid mutator transaction binding the contract method 0x8bb204a1.
//
// Solidity: function failStore() returns()
func (_EthCallTester *EthCallTesterSession) FailStore() (*types.Transaction, error) {
	return _EthCallTester.Contract.FailStore(&_EthCallTester.TransactOpts)
}

// FailStore is a paid mutator transaction binding the contract method 0x8bb204a1.
//
// Solidity: function failStore() returns()
func (_EthCallTester *EthCallTesterTransactorSession) FailStore() (*types.Transaction, error) {
	return _EthCallTester.Contract.FailStore(&_EthCallTester.TransactOpts)
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_EthCallTester *EthCallTesterTransactor) GetBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCallTester.contract.Transact(opts, "getBalance")
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_EthCallTester *EthCallTesterSession) GetBalance() (*types.Transaction, error) {
	return _EthCallTester.Contract.GetBalance(&_EthCallTester.TransactOpts)
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_EthCallTester *EthCallTesterTransactorSession) GetBalance() (*types.Transaction, error) {
	return _EthCallTester.Contract.GetBalance(&_EthCallTester.TransactOpts)
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(uint256)
func (_EthCallTester *EthCallTesterTransactor) GetX(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCallTester.contract.Transact(opts, "getX")
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(uint256)
func (_EthCallTester *EthCallTesterSession) GetX() (*types.Transaction, error) {
	return _EthCallTester.Contract.GetX(&_EthCallTester.TransactOpts)
}

// GetX is a paid mutator transaction binding the contract method 0x5197c7aa.
//
// Solidity: function getX() returns(uint256)
func (_EthCallTester *EthCallTesterTransactorSession) GetX() (*types.Transaction, error) {
	return _EthCallTester.Contract.GetX(&_EthCallTester.TransactOpts)
}

// SLoad is a paid mutator transaction binding the contract method 0x30874d8d.
//
// Solidity: function sLoad(uint256 key) returns(uint256)
func (_EthCallTester *EthCallTesterTransactor) SLoad(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _EthCallTester.contract.Transact(opts, "sLoad", key)
}

// SLoad is a paid mutator transaction binding the contract method 0x30874d8d.
//
// Solidity: function sLoad(uint256 key) returns(uint256)
func (_EthCallTester *EthCallTesterSession) SLoad(key *big.Int) (*types.Transaction, error) {
	return _EthCallTester.Contract.SLoad(&_EthCallTester.TransactOpts, key)
}

// SLoad is a paid mutator transaction binding the contract method 0x30874d8d.
//
// Solidity: function sLoad(uint256 key) returns(uint256)
func (_EthCallTester *EthCallTesterTransactorSession) SLoad(key *big.Int) (*types.Transaction, error) {
	return _EthCallTester.Contract.SLoad(&_EthCallTester.TransactOpts, key)
}
