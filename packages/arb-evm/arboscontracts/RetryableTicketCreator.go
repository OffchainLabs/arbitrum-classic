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

// RetryableTicketCreatorMetaData contains all meta data concerning the RetryableTicketCreator contract.
var RetryableTicketCreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RetryableTicketCreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RetryableTicketCreatorMetaData.ABI instead.
var RetryableTicketCreatorABI = RetryableTicketCreatorMetaData.ABI

// RetryableTicketCreator is an auto generated Go binding around an Ethereum contract.
type RetryableTicketCreator struct {
	RetryableTicketCreatorCaller     // Read-only binding to the contract
	RetryableTicketCreatorTransactor // Write-only binding to the contract
	RetryableTicketCreatorFilterer   // Log filterer for contract events
}

// RetryableTicketCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RetryableTicketCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RetryableTicketCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RetryableTicketCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RetryableTicketCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RetryableTicketCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RetryableTicketCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RetryableTicketCreatorSession struct {
	Contract     *RetryableTicketCreator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RetryableTicketCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RetryableTicketCreatorCallerSession struct {
	Contract *RetryableTicketCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// RetryableTicketCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RetryableTicketCreatorTransactorSession struct {
	Contract     *RetryableTicketCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// RetryableTicketCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RetryableTicketCreatorRaw struct {
	Contract *RetryableTicketCreator // Generic contract binding to access the raw methods on
}

// RetryableTicketCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RetryableTicketCreatorCallerRaw struct {
	Contract *RetryableTicketCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// RetryableTicketCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RetryableTicketCreatorTransactorRaw struct {
	Contract *RetryableTicketCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRetryableTicketCreator creates a new instance of RetryableTicketCreator, bound to a specific deployed contract.
func NewRetryableTicketCreator(address common.Address, backend bind.ContractBackend) (*RetryableTicketCreator, error) {
	contract, err := bindRetryableTicketCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RetryableTicketCreator{RetryableTicketCreatorCaller: RetryableTicketCreatorCaller{contract: contract}, RetryableTicketCreatorTransactor: RetryableTicketCreatorTransactor{contract: contract}, RetryableTicketCreatorFilterer: RetryableTicketCreatorFilterer{contract: contract}}, nil
}

// NewRetryableTicketCreatorCaller creates a new read-only instance of RetryableTicketCreator, bound to a specific deployed contract.
func NewRetryableTicketCreatorCaller(address common.Address, caller bind.ContractCaller) (*RetryableTicketCreatorCaller, error) {
	contract, err := bindRetryableTicketCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RetryableTicketCreatorCaller{contract: contract}, nil
}

// NewRetryableTicketCreatorTransactor creates a new write-only instance of RetryableTicketCreator, bound to a specific deployed contract.
func NewRetryableTicketCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RetryableTicketCreatorTransactor, error) {
	contract, err := bindRetryableTicketCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RetryableTicketCreatorTransactor{contract: contract}, nil
}

// NewRetryableTicketCreatorFilterer creates a new log filterer instance of RetryableTicketCreator, bound to a specific deployed contract.
func NewRetryableTicketCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RetryableTicketCreatorFilterer, error) {
	contract, err := bindRetryableTicketCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RetryableTicketCreatorFilterer{contract: contract}, nil
}

// bindRetryableTicketCreator binds a generic wrapper to an already deployed contract.
func bindRetryableTicketCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RetryableTicketCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RetryableTicketCreator *RetryableTicketCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RetryableTicketCreator.Contract.RetryableTicketCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RetryableTicketCreator *RetryableTicketCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RetryableTicketCreator.Contract.RetryableTicketCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RetryableTicketCreator *RetryableTicketCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RetryableTicketCreator.Contract.RetryableTicketCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RetryableTicketCreator *RetryableTicketCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RetryableTicketCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RetryableTicketCreator *RetryableTicketCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RetryableTicketCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RetryableTicketCreator *RetryableTicketCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RetryableTicketCreator.Contract.contract.Transact(opts, method, params...)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns()
func (_RetryableTicketCreator *RetryableTicketCreatorTransactor) CreateRetryableTicket(opts *bind.TransactOpts, destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _RetryableTicketCreator.contract.Transact(opts, "createRetryableTicket", destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns()
func (_RetryableTicketCreator *RetryableTicketCreatorSession) CreateRetryableTicket(destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _RetryableTicketCreator.Contract.CreateRetryableTicket(&_RetryableTicketCreator.TransactOpts, destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns()
func (_RetryableTicketCreator *RetryableTicketCreatorTransactorSession) CreateRetryableTicket(destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _RetryableTicketCreator.Contract.CreateRetryableTicket(&_RetryableTicketCreator.TransactOpts, destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}
