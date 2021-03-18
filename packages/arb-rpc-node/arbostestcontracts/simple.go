// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbostestcontracts

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

// SimpleABI is the input ABI used to generate the binding from.
const SimpleABI = "[{\"inputs\":[],\"name\":\"acceptPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"nestedCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rejectPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reverts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SimpleFuncSigs maps the 4-byte function signature to its string representation.
var SimpleFuncSigs = map[string]string{
	"ae0aba8c": "acceptPayment()",
	"267c4ae4": "exists()",
	"9b7c9da3": "nestedCall(uint256)",
	"9436bc1f": "rejectPayment()",
	"3bccbbc9": "reverts()",
}

// SimpleBin is the compiled bytecode used for deploying new contracts.
var SimpleBin = "0x608060405234801561001057600080fd5b506101de806100206000396000f3fe60806040526004361061004e5760003560e01c8063267c4ae4146100955780633bccbbc9146100bc5780639436bc1f146100d15780639b7c9da3146100e6578063ae0aba8c1461011057610090565b36610090576040805162461bcd60e51b815260206004820152600b60248201526a6e6f206465706f7369747360a81b604482015290519081900360640190fd5b005b600080fd5b3480156100a157600080fd5b506100aa610118565b60408051918252519081900360200190f35b3480156100c857600080fd5b5061008e61011d565b3480156100dd57600080fd5b5061008e61015b565b3480156100f257600080fd5b5061008e6004803603602081101561010957600080fd5b503561015d565b61008e61015b565b600a90565b6040805162461bcd60e51b815260206004820152600e60248201526d1d1a1a5cc81a5cc818481d195cdd60921b604482015290519081900360640190fd5b565b60405130908290600081818185875af1925050503d806000811461019d576040519150601f19603f3d011682016040523d82523d6000602084013e6101a2565b606091505b5050505056fea26469706673582212203fd4cf537d89fe44d609541d3dbe6ab55cb6cf4366dbdd2c259902863f9d8c0d64736f6c634300060c0033"

// DeploySimple deploys a new Ethereum contract, binding an instance of Simple to it.
func DeploySimple(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simple, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// Simple is an auto generated Go binding around an Ethereum contract.
type Simple struct {
	SimpleCaller     // Read-only binding to the contract
	SimpleTransactor // Write-only binding to the contract
	SimpleFilterer   // Log filterer for contract events
}

// SimpleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleSession struct {
	Contract     *Simple           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCallerSession struct {
	Contract *SimpleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SimpleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleTransactorSession struct {
	Contract     *SimpleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRaw struct {
	Contract *Simple // Generic contract binding to access the raw methods on
}

// SimpleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCallerRaw struct {
	Contract *SimpleCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleTransactorRaw struct {
	Contract *SimpleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimple creates a new instance of Simple, bound to a specific deployed contract.
func NewSimple(address common.Address, backend bind.ContractBackend) (*Simple, error) {
	contract, err := bindSimple(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// NewSimpleCaller creates a new read-only instance of Simple, bound to a specific deployed contract.
func NewSimpleCaller(address common.Address, caller bind.ContractCaller) (*SimpleCaller, error) {
	contract, err := bindSimple(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCaller{contract: contract}, nil
}

// NewSimpleTransactor creates a new write-only instance of Simple, bound to a specific deployed contract.
func NewSimpleTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTransactor, error) {
	contract, err := bindSimple(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransactor{contract: contract}, nil
}

// NewSimpleFilterer creates a new log filterer instance of Simple, bound to a specific deployed contract.
func NewSimpleFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleFilterer, error) {
	contract, err := bindSimple(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleFilterer{contract: contract}, nil
}

// bindSimple binds a generic wrapper to an already deployed contract.
func bindSimple(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.SimpleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transact(opts, method, params...)
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xae0aba8c.
//
// Solidity: function acceptPayment() payable returns()
func (_Simple *SimpleTransactor) AcceptPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "acceptPayment")
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xae0aba8c.
//
// Solidity: function acceptPayment() payable returns()
func (_Simple *SimpleSession) AcceptPayment() (*types.Transaction, error) {
	return _Simple.Contract.AcceptPayment(&_Simple.TransactOpts)
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xae0aba8c.
//
// Solidity: function acceptPayment() payable returns()
func (_Simple *SimpleTransactorSession) AcceptPayment() (*types.Transaction, error) {
	return _Simple.Contract.AcceptPayment(&_Simple.TransactOpts)
}

// Exists is a paid mutator transaction binding the contract method 0x267c4ae4.
//
// Solidity: function exists() returns(uint256)
func (_Simple *SimpleTransactor) Exists(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "exists")
}

// Exists is a paid mutator transaction binding the contract method 0x267c4ae4.
//
// Solidity: function exists() returns(uint256)
func (_Simple *SimpleSession) Exists() (*types.Transaction, error) {
	return _Simple.Contract.Exists(&_Simple.TransactOpts)
}

// Exists is a paid mutator transaction binding the contract method 0x267c4ae4.
//
// Solidity: function exists() returns(uint256)
func (_Simple *SimpleTransactorSession) Exists() (*types.Transaction, error) {
	return _Simple.Contract.Exists(&_Simple.TransactOpts)
}

// NestedCall is a paid mutator transaction binding the contract method 0x9b7c9da3.
//
// Solidity: function nestedCall(uint256 value) returns()
func (_Simple *SimpleTransactor) NestedCall(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "nestedCall", value)
}

// NestedCall is a paid mutator transaction binding the contract method 0x9b7c9da3.
//
// Solidity: function nestedCall(uint256 value) returns()
func (_Simple *SimpleSession) NestedCall(value *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.NestedCall(&_Simple.TransactOpts, value)
}

// NestedCall is a paid mutator transaction binding the contract method 0x9b7c9da3.
//
// Solidity: function nestedCall(uint256 value) returns()
func (_Simple *SimpleTransactorSession) NestedCall(value *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.NestedCall(&_Simple.TransactOpts, value)
}

// RejectPayment is a paid mutator transaction binding the contract method 0x9436bc1f.
//
// Solidity: function rejectPayment() returns()
func (_Simple *SimpleTransactor) RejectPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "rejectPayment")
}

// RejectPayment is a paid mutator transaction binding the contract method 0x9436bc1f.
//
// Solidity: function rejectPayment() returns()
func (_Simple *SimpleSession) RejectPayment() (*types.Transaction, error) {
	return _Simple.Contract.RejectPayment(&_Simple.TransactOpts)
}

// RejectPayment is a paid mutator transaction binding the contract method 0x9436bc1f.
//
// Solidity: function rejectPayment() returns()
func (_Simple *SimpleTransactorSession) RejectPayment() (*types.Transaction, error) {
	return _Simple.Contract.RejectPayment(&_Simple.TransactOpts)
}

// Reverts is a paid mutator transaction binding the contract method 0x3bccbbc9.
//
// Solidity: function reverts() returns()
func (_Simple *SimpleTransactor) Reverts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "reverts")
}

// Reverts is a paid mutator transaction binding the contract method 0x3bccbbc9.
//
// Solidity: function reverts() returns()
func (_Simple *SimpleSession) Reverts() (*types.Transaction, error) {
	return _Simple.Contract.Reverts(&_Simple.TransactOpts)
}

// Reverts is a paid mutator transaction binding the contract method 0x3bccbbc9.
//
// Solidity: function reverts() returns()
func (_Simple *SimpleTransactorSession) Reverts() (*types.Transaction, error) {
	return _Simple.Contract.Reverts(&_Simple.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simple *SimpleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simple *SimpleSession) Receive() (*types.Transaction, error) {
	return _Simple.Contract.Receive(&_Simple.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simple *SimpleTransactorSession) Receive() (*types.Transaction, error) {
	return _Simple.Contract.Receive(&_Simple.TransactOpts)
}
