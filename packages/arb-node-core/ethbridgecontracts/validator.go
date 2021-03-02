// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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

// ValidatorABI is the input ABI used to generate the binding from.
const ValidatorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"executeTransactions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"returnOldDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"contractIChallenge[]\",\"name\":\"challenges\",\"type\":\"address[]\"}],\"name\":\"timeoutChallenges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorFuncSigs = map[string]string{
	"ce1d571f": "executeTransaction(bytes,address,uint256)",
	"72f45866": "executeTransactions(bytes[],address[],uint256[])",
	"944f4495": "returnOldDeposits(address,address[])",
	"3711fe51": "timeoutChallenges(address,address[])",
}

// ValidatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610628806100326000396000f3fe60806040526004361061003f5760003560e01c80633711fe511461004457806372f4586614610066578063944f449514610079578063ce1d571f14610099575b600080fd5b34801561005057600080fd5b5061006461005f3660046104fa565b6100ac565b005b6100646100743660046103e0565b610134565b34801561008557600080fd5b506100646100943660046104fa565b610249565b6100646100a7366004610476565b6102d5565b8060005b8181101561012d578383828181106100c457fe5b90506020020160208101906100d991906103bd565b6001600160a01b03166370dea79a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561011357600080fd5b505af1925050508015610124575060015b506001016100b0565b5050505050565b6000546001600160a01b031633146101675760405162461bcd60e51b815260040161015e90610571565b60405180910390fd5b8460005b8181101561023f57600086868381811061018157fe5b905060200201602081019061019691906103bd565b6001600160a01b03168585848181106101ab57fe5b905060200201358a8a858181106101be57fe5b90506020028101906101d09190610595565b6040516101de92919061054d565b60006040518083038185875af1925050503d806000811461021b576040519150601f19603f3d011682016040523d82523d6000602084013e610220565b606091505b5050905080610236576040513d806000833e8082fd5b5060010161016b565b5050505050505050565b8060005b8181101561012d57846001600160a01b0316637427be5185858481811061027057fe5b905060200201602081019061028591906103bd565b6040518263ffffffff1660e01b81526004016102a1919061055d565b600060405180830381600087803b1580156102bb57600080fd5b505af19250505080156102cc575060015b5060010161024d565b6000546001600160a01b031633146102ff5760405162461bcd60e51b815260040161015e90610571565b6000826001600160a01b031682868660405161031c92919061054d565b60006040518083038185875af1925050503d8060008114610359576040519150601f19603f3d011682016040523d82523d6000602084013e61035e565b606091505b505090508061012d576040513d806000833e8082fd5b60008083601f840112610385578182fd5b50813567ffffffffffffffff81111561039c578182fd5b60208301915083602080830285010111156103b657600080fd5b9250929050565b6000602082840312156103ce578081fd5b81356103d9816105da565b9392505050565b600080600080600080606087890312156103f8578182fd5b863567ffffffffffffffff8082111561040f578384fd5b61041b8a838b01610374565b90985096506020890135915080821115610433578384fd5b61043f8a838b01610374565b90965094506040890135915080821115610457578384fd5b5061046489828a01610374565b979a9699509497509295939492505050565b6000806000806060858703121561048b578384fd5b843567ffffffffffffffff808211156104a2578586fd5b818701915087601f8301126104b5578586fd5b8135818111156104c3578687fd5b8860208285010111156104d4578687fd5b602092830196509450508501356104ea816105da565b9396929550929360400135925050565b60008060006040848603121561050e578283fd5b8335610519816105da565b9250602084013567ffffffffffffffff811115610534578283fd5b61054086828701610374565b9497909650939450505050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b6020808252600a908201526927a7262cafa7aba722a960b11b604082015260600190565b6000808335601e198436030181126105ab578283fd5b83018035915067ffffffffffffffff8211156105c5578283fd5b6020019150368190038213156103b657600080fd5b6001600160a01b03811681146105ef57600080fd5b5056fea264697066735822122030dee0cdbeed512bf8a5b365371369b98d9d6794d3e8f2d38908a94e22056bc364736f6c634300060c0033"

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorTransactor) ExecuteTransaction(opts *bind.TransactOpts, data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "executeTransaction", data, destination, amount)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorSession) ExecuteTransaction(data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransaction(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorTransactorSession) ExecuteTransaction(data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransaction(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorTransactor) ExecuteTransactions(opts *bind.TransactOpts, data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "executeTransactions", data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorSession) ExecuteTransactions(data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactions(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorTransactorSession) ExecuteTransactions(data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactions(&_Validator.TransactOpts, data, destination, amount)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorTransactor) ReturnOldDeposits(opts *bind.TransactOpts, rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "returnOldDeposits", rollup, stakers)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorSession) ReturnOldDeposits(rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ReturnOldDeposits(&_Validator.TransactOpts, rollup, stakers)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorTransactorSession) ReturnOldDeposits(rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ReturnOldDeposits(&_Validator.TransactOpts, rollup, stakers)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x3711fe51.
//
// Solidity: function timeoutChallenges(address rollup, address[] challenges) returns()
func (_Validator *ValidatorTransactor) TimeoutChallenges(opts *bind.TransactOpts, rollup common.Address, challenges []common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "timeoutChallenges", rollup, challenges)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x3711fe51.
//
// Solidity: function timeoutChallenges(address rollup, address[] challenges) returns()
func (_Validator *ValidatorSession) TimeoutChallenges(rollup common.Address, challenges []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TimeoutChallenges(&_Validator.TransactOpts, rollup, challenges)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x3711fe51.
//
// Solidity: function timeoutChallenges(address rollup, address[] challenges) returns()
func (_Validator *ValidatorTransactorSession) TimeoutChallenges(rollup common.Address, challenges []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TimeoutChallenges(&_Validator.TransactOpts, rollup, challenges)
}
