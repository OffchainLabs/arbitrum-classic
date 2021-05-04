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

// ArbRetryableTxABI is the input ABI used to generate the binding from.
const ArbRetryableTxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"LifetimeExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"}],\"name\":\"TicketCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"}],\"name\":\"getBeneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"}],\"name\":\"getKeepalivePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLifetime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"getSubmissionPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"}],\"name\":\"getTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ticketId\",\"type\":\"bytes32\"}],\"name\":\"keepalive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbRetryableTxFuncSigs maps the 4-byte function signature to its string representation.
var ArbRetryableTxFuncSigs = map[string]string{
	"c4d252f5": "cancel(bytes32)",
	"ba20dda4": "getBeneficiary(bytes32)",
	"b16607e5": "getKeepalivePrice(bytes32)",
	"81e6e083": "getLifetime()",
	"f88029dc": "getSubmissionPrice(uint256)",
	"9f1025c6": "getTimeout(bytes32)",
	"f0b21a41": "keepalive(bytes32)",
	"eda1122c": "redeem(bytes32)",
}

// ArbRetryableTx is an auto generated Go binding around an Ethereum contract.
type ArbRetryableTx struct {
	ArbRetryableTxCaller     // Read-only binding to the contract
	ArbRetryableTxTransactor // Write-only binding to the contract
	ArbRetryableTxFilterer   // Log filterer for contract events
}

// ArbRetryableTxCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbRetryableTxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRetryableTxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbRetryableTxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRetryableTxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbRetryableTxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRetryableTxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbRetryableTxSession struct {
	Contract     *ArbRetryableTx   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbRetryableTxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbRetryableTxCallerSession struct {
	Contract *ArbRetryableTxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ArbRetryableTxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbRetryableTxTransactorSession struct {
	Contract     *ArbRetryableTxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ArbRetryableTxRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbRetryableTxRaw struct {
	Contract *ArbRetryableTx // Generic contract binding to access the raw methods on
}

// ArbRetryableTxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbRetryableTxCallerRaw struct {
	Contract *ArbRetryableTxCaller // Generic read-only contract binding to access the raw methods on
}

// ArbRetryableTxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbRetryableTxTransactorRaw struct {
	Contract *ArbRetryableTxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbRetryableTx creates a new instance of ArbRetryableTx, bound to a specific deployed contract.
func NewArbRetryableTx(address common.Address, backend bind.ContractBackend) (*ArbRetryableTx, error) {
	contract, err := bindArbRetryableTx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbRetryableTx{ArbRetryableTxCaller: ArbRetryableTxCaller{contract: contract}, ArbRetryableTxTransactor: ArbRetryableTxTransactor{contract: contract}, ArbRetryableTxFilterer: ArbRetryableTxFilterer{contract: contract}}, nil
}

// NewArbRetryableTxCaller creates a new read-only instance of ArbRetryableTx, bound to a specific deployed contract.
func NewArbRetryableTxCaller(address common.Address, caller bind.ContractCaller) (*ArbRetryableTxCaller, error) {
	contract, err := bindArbRetryableTx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbRetryableTxCaller{contract: contract}, nil
}

// NewArbRetryableTxTransactor creates a new write-only instance of ArbRetryableTx, bound to a specific deployed contract.
func NewArbRetryableTxTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbRetryableTxTransactor, error) {
	contract, err := bindArbRetryableTx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbRetryableTxTransactor{contract: contract}, nil
}

// NewArbRetryableTxFilterer creates a new log filterer instance of ArbRetryableTx, bound to a specific deployed contract.
func NewArbRetryableTxFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbRetryableTxFilterer, error) {
	contract, err := bindArbRetryableTx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbRetryableTxFilterer{contract: contract}, nil
}

// bindArbRetryableTx binds a generic wrapper to an already deployed contract.
func bindArbRetryableTx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbRetryableTxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbRetryableTx *ArbRetryableTxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbRetryableTx.Contract.ArbRetryableTxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbRetryableTx *ArbRetryableTxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.ArbRetryableTxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbRetryableTx *ArbRetryableTxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.ArbRetryableTxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbRetryableTx *ArbRetryableTxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbRetryableTx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbRetryableTx *ArbRetryableTxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbRetryableTx *ArbRetryableTxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.contract.Transact(opts, method, params...)
}

// GetBeneficiary is a free data retrieval call binding the contract method 0xba20dda4.
//
// Solidity: function getBeneficiary(bytes32 ticketId) view returns(address)
func (_ArbRetryableTx *ArbRetryableTxCaller) GetBeneficiary(opts *bind.CallOpts, ticketId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ArbRetryableTx.contract.Call(opts, &out, "getBeneficiary", ticketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBeneficiary is a free data retrieval call binding the contract method 0xba20dda4.
//
// Solidity: function getBeneficiary(bytes32 ticketId) view returns(address)
func (_ArbRetryableTx *ArbRetryableTxSession) GetBeneficiary(ticketId [32]byte) (common.Address, error) {
	return _ArbRetryableTx.Contract.GetBeneficiary(&_ArbRetryableTx.CallOpts, ticketId)
}

// GetBeneficiary is a free data retrieval call binding the contract method 0xba20dda4.
//
// Solidity: function getBeneficiary(bytes32 ticketId) view returns(address)
func (_ArbRetryableTx *ArbRetryableTxCallerSession) GetBeneficiary(ticketId [32]byte) (common.Address, error) {
	return _ArbRetryableTx.Contract.GetBeneficiary(&_ArbRetryableTx.CallOpts, ticketId)
}

// GetKeepalivePrice is a free data retrieval call binding the contract method 0xb16607e5.
//
// Solidity: function getKeepalivePrice(bytes32 ticketId) view returns(uint256, uint256)
func (_ArbRetryableTx *ArbRetryableTxCaller) GetKeepalivePrice(opts *bind.CallOpts, ticketId [32]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbRetryableTx.contract.Call(opts, &out, "getKeepalivePrice", ticketId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetKeepalivePrice is a free data retrieval call binding the contract method 0xb16607e5.
//
// Solidity: function getKeepalivePrice(bytes32 ticketId) view returns(uint256, uint256)
func (_ArbRetryableTx *ArbRetryableTxSession) GetKeepalivePrice(ticketId [32]byte) (*big.Int, *big.Int, error) {
	return _ArbRetryableTx.Contract.GetKeepalivePrice(&_ArbRetryableTx.CallOpts, ticketId)
}

// GetKeepalivePrice is a free data retrieval call binding the contract method 0xb16607e5.
//
// Solidity: function getKeepalivePrice(bytes32 ticketId) view returns(uint256, uint256)
func (_ArbRetryableTx *ArbRetryableTxCallerSession) GetKeepalivePrice(ticketId [32]byte) (*big.Int, *big.Int, error) {
	return _ArbRetryableTx.Contract.GetKeepalivePrice(&_ArbRetryableTx.CallOpts, ticketId)
}

// GetLifetime is a free data retrieval call binding the contract method 0x81e6e083.
//
// Solidity: function getLifetime() view returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxCaller) GetLifetime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbRetryableTx.contract.Call(opts, &out, "getLifetime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLifetime is a free data retrieval call binding the contract method 0x81e6e083.
//
// Solidity: function getLifetime() view returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxSession) GetLifetime() (*big.Int, error) {
	return _ArbRetryableTx.Contract.GetLifetime(&_ArbRetryableTx.CallOpts)
}

// GetLifetime is a free data retrieval call binding the contract method 0x81e6e083.
//
// Solidity: function getLifetime() view returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxCallerSession) GetLifetime() (*big.Int, error) {
	return _ArbRetryableTx.Contract.GetLifetime(&_ArbRetryableTx.CallOpts)
}

// GetSubmissionPrice is a free data retrieval call binding the contract method 0xf88029dc.
//
// Solidity: function getSubmissionPrice(uint256 calldataSize) view returns(uint256, uint256)
func (_ArbRetryableTx *ArbRetryableTxCaller) GetSubmissionPrice(opts *bind.CallOpts, calldataSize *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbRetryableTx.contract.Call(opts, &out, "getSubmissionPrice", calldataSize)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSubmissionPrice is a free data retrieval call binding the contract method 0xf88029dc.
//
// Solidity: function getSubmissionPrice(uint256 calldataSize) view returns(uint256, uint256)
func (_ArbRetryableTx *ArbRetryableTxSession) GetSubmissionPrice(calldataSize *big.Int) (*big.Int, *big.Int, error) {
	return _ArbRetryableTx.Contract.GetSubmissionPrice(&_ArbRetryableTx.CallOpts, calldataSize)
}

// GetSubmissionPrice is a free data retrieval call binding the contract method 0xf88029dc.
//
// Solidity: function getSubmissionPrice(uint256 calldataSize) view returns(uint256, uint256)
func (_ArbRetryableTx *ArbRetryableTxCallerSession) GetSubmissionPrice(calldataSize *big.Int) (*big.Int, *big.Int, error) {
	return _ArbRetryableTx.Contract.GetSubmissionPrice(&_ArbRetryableTx.CallOpts, calldataSize)
}

// GetTimeout is a free data retrieval call binding the contract method 0x9f1025c6.
//
// Solidity: function getTimeout(bytes32 ticketId) view returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxCaller) GetTimeout(opts *bind.CallOpts, ticketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ArbRetryableTx.contract.Call(opts, &out, "getTimeout", ticketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeout is a free data retrieval call binding the contract method 0x9f1025c6.
//
// Solidity: function getTimeout(bytes32 ticketId) view returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxSession) GetTimeout(ticketId [32]byte) (*big.Int, error) {
	return _ArbRetryableTx.Contract.GetTimeout(&_ArbRetryableTx.CallOpts, ticketId)
}

// GetTimeout is a free data retrieval call binding the contract method 0x9f1025c6.
//
// Solidity: function getTimeout(bytes32 ticketId) view returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxCallerSession) GetTimeout(ticketId [32]byte) (*big.Int, error) {
	return _ArbRetryableTx.Contract.GetTimeout(&_ArbRetryableTx.CallOpts, ticketId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 ticketId) returns()
func (_ArbRetryableTx *ArbRetryableTxTransactor) Cancel(opts *bind.TransactOpts, ticketId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.contract.Transact(opts, "cancel", ticketId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 ticketId) returns()
func (_ArbRetryableTx *ArbRetryableTxSession) Cancel(ticketId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.Cancel(&_ArbRetryableTx.TransactOpts, ticketId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 ticketId) returns()
func (_ArbRetryableTx *ArbRetryableTxTransactorSession) Cancel(ticketId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.Cancel(&_ArbRetryableTx.TransactOpts, ticketId)
}

// Keepalive is a paid mutator transaction binding the contract method 0xf0b21a41.
//
// Solidity: function keepalive(bytes32 ticketId) payable returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxTransactor) Keepalive(opts *bind.TransactOpts, ticketId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.contract.Transact(opts, "keepalive", ticketId)
}

// Keepalive is a paid mutator transaction binding the contract method 0xf0b21a41.
//
// Solidity: function keepalive(bytes32 ticketId) payable returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxSession) Keepalive(ticketId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.Keepalive(&_ArbRetryableTx.TransactOpts, ticketId)
}

// Keepalive is a paid mutator transaction binding the contract method 0xf0b21a41.
//
// Solidity: function keepalive(bytes32 ticketId) payable returns(uint256)
func (_ArbRetryableTx *ArbRetryableTxTransactorSession) Keepalive(ticketId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.Keepalive(&_ArbRetryableTx.TransactOpts, ticketId)
}

// Redeem is a paid mutator transaction binding the contract method 0xeda1122c.
//
// Solidity: function redeem(bytes32 txId) returns()
func (_ArbRetryableTx *ArbRetryableTxTransactor) Redeem(opts *bind.TransactOpts, txId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.contract.Transact(opts, "redeem", txId)
}

// Redeem is a paid mutator transaction binding the contract method 0xeda1122c.
//
// Solidity: function redeem(bytes32 txId) returns()
func (_ArbRetryableTx *ArbRetryableTxSession) Redeem(txId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.Redeem(&_ArbRetryableTx.TransactOpts, txId)
}

// Redeem is a paid mutator transaction binding the contract method 0xeda1122c.
//
// Solidity: function redeem(bytes32 txId) returns()
func (_ArbRetryableTx *ArbRetryableTxTransactorSession) Redeem(txId [32]byte) (*types.Transaction, error) {
	return _ArbRetryableTx.Contract.Redeem(&_ArbRetryableTx.TransactOpts, txId)
}

// ArbRetryableTxCanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the ArbRetryableTx contract.
type ArbRetryableTxCanceledIterator struct {
	Event *ArbRetryableTxCanceled // Event containing the contract specifics and raw log

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
func (it *ArbRetryableTxCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRetryableTxCanceled)
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
		it.Event = new(ArbRetryableTxCanceled)
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
func (it *ArbRetryableTxCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRetryableTxCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRetryableTxCanceled represents a Canceled event raised by the ArbRetryableTx contract.
type ArbRetryableTxCanceled struct {
	TicketId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0x134fdd648feeaf30251f0157f9624ef8608ff9a042aad6d13e73f35d21d3f88d.
//
// Solidity: event Canceled(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) FilterCanceled(opts *bind.FilterOpts, ticketId [][32]byte) (*ArbRetryableTxCanceledIterator, error) {

	var ticketIdRule []interface{}
	for _, ticketIdItem := range ticketId {
		ticketIdRule = append(ticketIdRule, ticketIdItem)
	}

	logs, sub, err := _ArbRetryableTx.contract.FilterLogs(opts, "Canceled", ticketIdRule)
	if err != nil {
		return nil, err
	}
	return &ArbRetryableTxCanceledIterator{contract: _ArbRetryableTx.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0x134fdd648feeaf30251f0157f9624ef8608ff9a042aad6d13e73f35d21d3f88d.
//
// Solidity: event Canceled(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *ArbRetryableTxCanceled, ticketId [][32]byte) (event.Subscription, error) {

	var ticketIdRule []interface{}
	for _, ticketIdItem := range ticketId {
		ticketIdRule = append(ticketIdRule, ticketIdItem)
	}

	logs, sub, err := _ArbRetryableTx.contract.WatchLogs(opts, "Canceled", ticketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRetryableTxCanceled)
				if err := _ArbRetryableTx.contract.UnpackLog(event, "Canceled", log); err != nil {
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

// ParseCanceled is a log parse operation binding the contract event 0x134fdd648feeaf30251f0157f9624ef8608ff9a042aad6d13e73f35d21d3f88d.
//
// Solidity: event Canceled(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) ParseCanceled(log types.Log) (*ArbRetryableTxCanceled, error) {
	event := new(ArbRetryableTxCanceled)
	if err := _ArbRetryableTx.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRetryableTxLifetimeExtendedIterator is returned from FilterLifetimeExtended and is used to iterate over the raw logs and unpacked data for LifetimeExtended events raised by the ArbRetryableTx contract.
type ArbRetryableTxLifetimeExtendedIterator struct {
	Event *ArbRetryableTxLifetimeExtended // Event containing the contract specifics and raw log

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
func (it *ArbRetryableTxLifetimeExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRetryableTxLifetimeExtended)
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
		it.Event = new(ArbRetryableTxLifetimeExtended)
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
func (it *ArbRetryableTxLifetimeExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRetryableTxLifetimeExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRetryableTxLifetimeExtended represents a LifetimeExtended event raised by the ArbRetryableTx contract.
type ArbRetryableTxLifetimeExtended struct {
	TicketId   [32]byte
	NewTimeout *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLifetimeExtended is a free log retrieval operation binding the contract event 0xf4c40a5f930e1469fcc053bf25f045253a7bad2fcc9b88c05ec1fca8e2066b83.
//
// Solidity: event LifetimeExtended(bytes32 indexed ticketId, uint256 newTimeout)
func (_ArbRetryableTx *ArbRetryableTxFilterer) FilterLifetimeExtended(opts *bind.FilterOpts, ticketId [][32]byte) (*ArbRetryableTxLifetimeExtendedIterator, error) {

	var ticketIdRule []interface{}
	for _, ticketIdItem := range ticketId {
		ticketIdRule = append(ticketIdRule, ticketIdItem)
	}

	logs, sub, err := _ArbRetryableTx.contract.FilterLogs(opts, "LifetimeExtended", ticketIdRule)
	if err != nil {
		return nil, err
	}
	return &ArbRetryableTxLifetimeExtendedIterator{contract: _ArbRetryableTx.contract, event: "LifetimeExtended", logs: logs, sub: sub}, nil
}

// WatchLifetimeExtended is a free log subscription operation binding the contract event 0xf4c40a5f930e1469fcc053bf25f045253a7bad2fcc9b88c05ec1fca8e2066b83.
//
// Solidity: event LifetimeExtended(bytes32 indexed ticketId, uint256 newTimeout)
func (_ArbRetryableTx *ArbRetryableTxFilterer) WatchLifetimeExtended(opts *bind.WatchOpts, sink chan<- *ArbRetryableTxLifetimeExtended, ticketId [][32]byte) (event.Subscription, error) {

	var ticketIdRule []interface{}
	for _, ticketIdItem := range ticketId {
		ticketIdRule = append(ticketIdRule, ticketIdItem)
	}

	logs, sub, err := _ArbRetryableTx.contract.WatchLogs(opts, "LifetimeExtended", ticketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRetryableTxLifetimeExtended)
				if err := _ArbRetryableTx.contract.UnpackLog(event, "LifetimeExtended", log); err != nil {
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

// ParseLifetimeExtended is a log parse operation binding the contract event 0xf4c40a5f930e1469fcc053bf25f045253a7bad2fcc9b88c05ec1fca8e2066b83.
//
// Solidity: event LifetimeExtended(bytes32 indexed ticketId, uint256 newTimeout)
func (_ArbRetryableTx *ArbRetryableTxFilterer) ParseLifetimeExtended(log types.Log) (*ArbRetryableTxLifetimeExtended, error) {
	event := new(ArbRetryableTxLifetimeExtended)
	if err := _ArbRetryableTx.contract.UnpackLog(event, "LifetimeExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRetryableTxRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the ArbRetryableTx contract.
type ArbRetryableTxRedeemedIterator struct {
	Event *ArbRetryableTxRedeemed // Event containing the contract specifics and raw log

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
func (it *ArbRetryableTxRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRetryableTxRedeemed)
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
		it.Event = new(ArbRetryableTxRedeemed)
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
func (it *ArbRetryableTxRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRetryableTxRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRetryableTxRedeemed represents a Redeemed event raised by the ArbRetryableTx contract.
type ArbRetryableTxRedeemed struct {
	TicketId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x27fc6cca2a0e9eb6f4876c01fc7779b00cdeb7277a770ac2b844db5932449578.
//
// Solidity: event Redeemed(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) FilterRedeemed(opts *bind.FilterOpts, ticketId [][32]byte) (*ArbRetryableTxRedeemedIterator, error) {

	var ticketIdRule []interface{}
	for _, ticketIdItem := range ticketId {
		ticketIdRule = append(ticketIdRule, ticketIdItem)
	}

	logs, sub, err := _ArbRetryableTx.contract.FilterLogs(opts, "Redeemed", ticketIdRule)
	if err != nil {
		return nil, err
	}
	return &ArbRetryableTxRedeemedIterator{contract: _ArbRetryableTx.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x27fc6cca2a0e9eb6f4876c01fc7779b00cdeb7277a770ac2b844db5932449578.
//
// Solidity: event Redeemed(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *ArbRetryableTxRedeemed, ticketId [][32]byte) (event.Subscription, error) {

	var ticketIdRule []interface{}
	for _, ticketIdItem := range ticketId {
		ticketIdRule = append(ticketIdRule, ticketIdItem)
	}

	logs, sub, err := _ArbRetryableTx.contract.WatchLogs(opts, "Redeemed", ticketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRetryableTxRedeemed)
				if err := _ArbRetryableTx.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0x27fc6cca2a0e9eb6f4876c01fc7779b00cdeb7277a770ac2b844db5932449578.
//
// Solidity: event Redeemed(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) ParseRedeemed(log types.Log) (*ArbRetryableTxRedeemed, error) {
	event := new(ArbRetryableTxRedeemed)
	if err := _ArbRetryableTx.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRetryableTxTicketCreatedIterator is returned from FilterTicketCreated and is used to iterate over the raw logs and unpacked data for TicketCreated events raised by the ArbRetryableTx contract.
type ArbRetryableTxTicketCreatedIterator struct {
	Event *ArbRetryableTxTicketCreated // Event containing the contract specifics and raw log

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
func (it *ArbRetryableTxTicketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRetryableTxTicketCreated)
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
		it.Event = new(ArbRetryableTxTicketCreated)
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
func (it *ArbRetryableTxTicketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRetryableTxTicketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRetryableTxTicketCreated represents a TicketCreated event raised by the ArbRetryableTx contract.
type ArbRetryableTxTicketCreated struct {
	TicketId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTicketCreated is a free log retrieval operation binding the contract event 0x7c793cced5743dc5f531bbe2bfb5a9fa3f40adef29231e6ab165c08a29e3dd89.
//
// Solidity: event TicketCreated(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) FilterTicketCreated(opts *bind.FilterOpts, ticketId [][32]byte) (*ArbRetryableTxTicketCreatedIterator, error) {

	var ticketIdRule []interface{}
	for _, ticketIdItem := range ticketId {
		ticketIdRule = append(ticketIdRule, ticketIdItem)
	}

	logs, sub, err := _ArbRetryableTx.contract.FilterLogs(opts, "TicketCreated", ticketIdRule)
	if err != nil {
		return nil, err
	}
	return &ArbRetryableTxTicketCreatedIterator{contract: _ArbRetryableTx.contract, event: "TicketCreated", logs: logs, sub: sub}, nil
}

// WatchTicketCreated is a free log subscription operation binding the contract event 0x7c793cced5743dc5f531bbe2bfb5a9fa3f40adef29231e6ab165c08a29e3dd89.
//
// Solidity: event TicketCreated(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) WatchTicketCreated(opts *bind.WatchOpts, sink chan<- *ArbRetryableTxTicketCreated, ticketId [][32]byte) (event.Subscription, error) {

	var ticketIdRule []interface{}
	for _, ticketIdItem := range ticketId {
		ticketIdRule = append(ticketIdRule, ticketIdItem)
	}

	logs, sub, err := _ArbRetryableTx.contract.WatchLogs(opts, "TicketCreated", ticketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRetryableTxTicketCreated)
				if err := _ArbRetryableTx.contract.UnpackLog(event, "TicketCreated", log); err != nil {
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

// ParseTicketCreated is a log parse operation binding the contract event 0x7c793cced5743dc5f531bbe2bfb5a9fa3f40adef29231e6ab165c08a29e3dd89.
//
// Solidity: event TicketCreated(bytes32 indexed ticketId)
func (_ArbRetryableTx *ArbRetryableTxFilterer) ParseTicketCreated(log types.Log) (*ArbRetryableTxTicketCreated, error) {
	event := new(ArbRetryableTxTicketCreated)
	if err := _ArbRetryableTx.contract.UnpackLog(event, "TicketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
