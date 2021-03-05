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

// IBridgeABI is the input ABI used to generate the binding from.
const IBridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"deliverMessageToInbox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBridgeFuncSigs maps the 4-byte function signature to its string representation.
var IBridgeFuncSigs = map[string]string{
	"ab5d8943": "activeOutbox()",
	"c29372de": "allowedInboxes(address)",
	"413b35bd": "allowedOutboxes(address)",
	"02bbfad1": "deliverMessageToInbox(uint8,address,bytes32)",
	"9e5d4c49": "executeCall(address,uint256,bytes)",
	"d9dd67ab": "inboxAccs(uint256)",
	"3dbcc8d1": "messageCount()",
	"e45b7ce6": "setInbox(address,bool)",
	"cee3d728": "setOutbox(address,bool)",
}

// IBridge is an auto generated Go binding around an Ethereum contract.
type IBridge struct {
	IBridgeCaller     // Read-only binding to the contract
	IBridgeTransactor // Write-only binding to the contract
	IBridgeFilterer   // Log filterer for contract events
}

// IBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeSession struct {
	Contract     *IBridge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeCallerSession struct {
	Contract *IBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeTransactorSession struct {
	Contract     *IBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeRaw struct {
	Contract *IBridge // Generic contract binding to access the raw methods on
}

// IBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeCallerRaw struct {
	Contract *IBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeTransactorRaw struct {
	Contract *IBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridge creates a new instance of IBridge, bound to a specific deployed contract.
func NewIBridge(address common.Address, backend bind.ContractBackend) (*IBridge, error) {
	contract, err := bindIBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridge{IBridgeCaller: IBridgeCaller{contract: contract}, IBridgeTransactor: IBridgeTransactor{contract: contract}, IBridgeFilterer: IBridgeFilterer{contract: contract}}, nil
}

// NewIBridgeCaller creates a new read-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeCaller(address common.Address, caller bind.ContractCaller) (*IBridgeCaller, error) {
	contract, err := bindIBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeCaller{contract: contract}, nil
}

// NewIBridgeTransactor creates a new write-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeTransactor, error) {
	contract, err := bindIBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeTransactor{contract: contract}, nil
}

// NewIBridgeFilterer creates a new log filterer instance of IBridge, bound to a specific deployed contract.
func NewIBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeFilterer, error) {
	contract, err := bindIBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeFilterer{contract: contract}, nil
}

// bindIBridge binds a generic wrapper to an already deployed contract.
func bindIBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.IBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_IBridge *IBridgeCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_IBridge *IBridgeSession) ActiveOutbox() (common.Address, error) {
	return _IBridge.Contract.ActiveOutbox(&_IBridge.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_IBridge *IBridgeCallerSession) ActiveOutbox() (common.Address, error) {
	return _IBridge.Contract.ActiveOutbox(&_IBridge.CallOpts)
}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_IBridge *IBridgeCaller) AllowedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "allowedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_IBridge *IBridgeSession) AllowedInboxes(inbox common.Address) (bool, error) {
	return _IBridge.Contract.AllowedInboxes(&_IBridge.CallOpts, inbox)
}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_IBridge *IBridgeCallerSession) AllowedInboxes(inbox common.Address) (bool, error) {
	return _IBridge.Contract.AllowedInboxes(&_IBridge.CallOpts, inbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_IBridge *IBridgeCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_IBridge *IBridgeSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _IBridge.Contract.AllowedOutboxes(&_IBridge.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_IBridge *IBridgeCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _IBridge.Contract.AllowedOutboxes(&_IBridge.CallOpts, outbox)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeCaller) InboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "inboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _IBridge.Contract.InboxAccs(&_IBridge.CallOpts, index)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeCallerSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _IBridge.Contract.InboxAccs(&_IBridge.CallOpts, index)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_IBridge *IBridgeCaller) MessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "messageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_IBridge *IBridgeSession) MessageCount() (*big.Int, error) {
	return _IBridge.Contract.MessageCount(&_IBridge.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_IBridge *IBridgeCallerSession) MessageCount() (*big.Int, error) {
	return _IBridge.Contract.MessageCount(&_IBridge.CallOpts)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_IBridge *IBridgeTransactor) DeliverMessageToInbox(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "deliverMessageToInbox", kind, sender, messageDataHash)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_IBridge *IBridgeSession) DeliverMessageToInbox(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.Contract.DeliverMessageToInbox(&_IBridge.TransactOpts, kind, sender, messageDataHash)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_IBridge *IBridgeTransactorSession) DeliverMessageToInbox(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.Contract.DeliverMessageToInbox(&_IBridge.TransactOpts, kind, sender, messageDataHash)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_IBridge *IBridgeTransactor) ExecuteCall(opts *bind.TransactOpts, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "executeCall", destAddr, amount, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_IBridge *IBridgeSession) ExecuteCall(destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IBridge.Contract.ExecuteCall(&_IBridge.TransactOpts, destAddr, amount, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_IBridge *IBridgeTransactorSession) ExecuteCall(destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IBridge.Contract.ExecuteCall(&_IBridge.TransactOpts, destAddr, amount, data)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeTransactor) SetInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "setInbox", inbox, enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeSession) SetInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.Contract.SetInbox(&_IBridge.TransactOpts, inbox, enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeTransactorSession) SetInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.Contract.SetInbox(&_IBridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeTransactor) SetOutbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "setOutbox", inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeSession) SetOutbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.Contract.SetOutbox(&_IBridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeTransactorSession) SetOutbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.Contract.SetOutbox(&_IBridge.TransactOpts, inbox, enabled)
}

// IBridgeMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the IBridge contract.
type IBridgeMessageDeliveredIterator struct {
	Event *IBridgeMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IBridgeMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeMessageDelivered)
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
		it.Event = new(IBridgeMessageDelivered)
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
func (it *IBridgeMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeMessageDelivered represents a MessageDelivered event raised by the IBridge contract.
type IBridgeMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_IBridge *IBridgeFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*IBridgeMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &IBridgeMessageDeliveredIterator{contract: _IBridge.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_IBridge *IBridgeFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *IBridgeMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeMessageDelivered)
				if err := _IBridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_IBridge *IBridgeFilterer) ParseMessageDelivered(log types.Log) (*IBridgeMessageDelivered, error) {
	event := new(IBridgeMessageDelivered)
	if err := _IBridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c70f3658137bd01684c46dc413dff99dba840ff34270b124fea8da96fc159e9164736f6c634300060c0033"

// DeployMessages deploys a new Ethereum contract, binding an instance of Messages to it.
func DeployMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Messages, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// Messages is an auto generated Go binding around an Ethereum contract.
type Messages struct {
	MessagesCaller     // Read-only binding to the contract
	MessagesTransactor // Write-only binding to the contract
	MessagesFilterer   // Log filterer for contract events
}

// MessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagesSession struct {
	Contract     *Messages         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagesCallerSession struct {
	Contract *MessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagesTransactorSession struct {
	Contract     *MessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagesRaw struct {
	Contract *Messages // Generic contract binding to access the raw methods on
}

// MessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagesCallerRaw struct {
	Contract *MessagesCaller // Generic read-only contract binding to access the raw methods on
}

// MessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagesTransactorRaw struct {
	Contract *MessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessages creates a new instance of Messages, bound to a specific deployed contract.
func NewMessages(address common.Address, backend bind.ContractBackend) (*Messages, error) {
	contract, err := bindMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// NewMessagesCaller creates a new read-only instance of Messages, bound to a specific deployed contract.
func NewMessagesCaller(address common.Address, caller bind.ContractCaller) (*MessagesCaller, error) {
	contract, err := bindMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesCaller{contract: contract}, nil
}

// NewMessagesTransactor creates a new write-only instance of Messages, bound to a specific deployed contract.
func NewMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagesTransactor, error) {
	contract, err := bindMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesTransactor{contract: contract}, nil
}

// NewMessagesFilterer creates a new log filterer instance of Messages, bound to a specific deployed contract.
func NewMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagesFilterer, error) {
	contract, err := bindMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagesFilterer{contract: contract}, nil
}

// bindMessages binds a generic wrapper to an already deployed contract.
func bindMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.MessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transact(opts, method, params...)
}

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"totalMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[4]\",\"name\":\"fields\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStepDebug\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"startMachine\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"afterMachine\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"9d16dd04": "executeStep(address,uint256,bytes32[2],bytes,bytes)",
	"2ccebb7a": "executeStepDebug(address,uint256,bytes32[2],bytes,bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x608060405234801561001057600080fd5b50614d99806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632ccebb7a1461003b5780639d16dd04146101f2575b600080fd5b610114600480360360c081101561005157600080fd5b6001600160a01b038235169160208101359160408201919081019060a081016080820135600160201b81111561008657600080fd5b82018360208201111561009857600080fd5b803590602001918460018302840111600160201b831117156100b957600080fd5b919390929091602081019035600160201b8111156100d657600080fd5b8201836020820111156100e857600080fd5b803590602001918460018302840111600160201b8311171561010957600080fd5b50909250905061031d565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561015557818101518382015260200161013d565b50505050905090810190601f1680156101825780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156101b557818101518382015260200161019d565b50505050905090810190601f1680156101e25780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6102cb600480360360c081101561020857600080fd5b6001600160a01b038235169160208101359160408201919081019060a081016080820135600160201b81111561023d57600080fd5b82018360208201111561024f57600080fd5b803590602001918460018302840111600160201b8311171561027057600080fd5b919390929091602081019035600160201b81111561028d57600080fd5b82018360208201111561029f57600080fd5b803590602001918460018302840111600160201b831117156102c057600080fd5b5090925090506103e2565b60405180846001600160401b0316815260200183815260200182600460200280838360005b838110156103085781810151838201526020016102f0565b50505050905001935050505060405180910390f35b606080610328614acc565b6103ac898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8d018190048102820181019092528b815292508b91508a9081908401838280828437600081840152601f19601f820116905080830192505050505050508e6104a1565b90506103b781610955565b6103c48160200151610d63565b92506103d38160400151610d63565b91505097509795505050505050565b6000806103ed614b62565b6103f5614acc565b6104798a8a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b9081908401838280828437600081840152601f19601f820116905080830192505050505050508f6104a1565b905061048481610955565b61048d816110d7565b935093509350509750975097945050505050565b6104a9614acc565b6000846000815181106104b857fe5b602001015160f81c60f81b60f81c90506000856001815181106104d757fe5b602001015160f81c60f81b60f81c90506000866002815181106104f657fe5b016020015160f81c9050600360606004840160ff166001600160401b038111801561052057600080fd5b5060405190808252806020026020018201604052801561055a57816020015b610547614b80565b81526020019060019003908161053f5790505b50905060608360040160ff166001600160401b038111801561057b57600080fd5b506040519080825280602002602001820160405280156105b557816020015b6105a2614b80565b81526020019060019003908161059a5790505b50905060005b8560ff168110156105f1576105d08b85611196565b8483815181106105dc57fe5b602090810291909101015293506001016105bb565b5060005b8460ff1681101561062b5761060a8b85611196565b83838151811061061657fe5b602090810291909101015293506001016105f5565b50610634614bbd565b61063e8b85611358565b809250819550505060008b858151811061065457fe5b01602001516001959095019460f81c905061066d614acc565b6001600160a01b038b1681526020810183905261068983611409565b6040820152606081018f90528d6000602002013560808201528d60016020908102919091013560a0830152600060c0830181905260408051808201825260ff8c811682528185018a905260e086019190915281518083019092528a8116825292810187905261010084015283821660018114610120850152918b1661014084015261016083018f90526101a083018e90526101c08301526101808201879052158061073757508160ff166001145b6040518060400160405280600b81526020016a04241445f494d4d5f5459560ac1b815250906107e45760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156107a9578181015183820152602001610791565b50505050905090810190601f1680156107d65780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506107ed614b80565b60ff831661080e576108078a83602001516000015161147d565b90506108ae565b6000865111604051806040016040528060068152602001654e4f5f494d4d60d01b8152509061087e5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107a9578181015183820152602001610791565b506108ab8a8360200151600001518860018d0360ff168151811061089e57fe5b60200260200101516114e1565b90505b6108b781611565565b60208301515260005b838a0360ff168110156108ff576108f78782815181106108dc57fe5b602002602001015184602001516116d290919063ffffffff16565b6001016108c0565b5060005b8860ff168110156109405761093886828151811061091d57fe5b602002602001015184602001516116ec90919063ffffffff16565b600101610903565b50909f9e505050505050505050505050505050565b6000806000612d8c61096e85610140015160ff16611706565b935093509350935060008411806109885750846101200151155b8015610999575060e0850151518410155b806109c0575084610120015180156109af575083155b80156109c0575060e0850151516001145b6040518060400160405280600a815260200169535441434b5f4d414e5960b01b81525090610a2f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107a9578181015183820152602001610791565b50610100850151516040805180820190915260088152674155585f4d414e5960c01b602082015290841015610aa55760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107a9578181015183820152602001610791565b5060e085015151841115610b6257610ac3610abe611e85565b611565565b610ad4866040015160200151611565565b146040518060400160405280600d81526020016c535441434b5f4d495353494e4760981b81525090610b475760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107a9578181015183820152602001610791565b50610b53856005611ecc565b50610b5d85611f41565b610c22565b61010085015151831115610bfd57610b7b610abe611e85565b610b8c866040015160400151611565565b146040518060400160405280600b81526020016a4155585f4d495353494e4760a81b81525090610b475760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107a9578181015183820152602001610791565b610c078583611ecc565b15610c1557610b5d85611f41565b610c22858263ffffffff16565b846101c0015115610cc45760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201835281519101209086015160c001511415610c8757610c828560400151611f4c565b610cc4565b60006101c0860152604085015160c081015190526101208501518015610cab575083155b610cb95760e0850151600090525b610100850151600090525b60005b60e086015151811015610d0e57610d068660e00151602001518281518110610ceb57fe5b602002602001015187604001516116d290919063ffffffff16565b600101610cc7565b5060005b61010086015151811015610d5b57610d53866101000151602001518281518110610d3857fe5b602002602001015187604001516116ec90919063ffffffff16565b600101610d12565b505050505050565b6060610d728260000151611f57565b610d87610d828460200151611565565b611f57565b610d97610d828560400151611565565b610da7610d828660600151611565565b610db7610d828760800151611565565b610dc48760a00151612026565b610dd18860c00151611f57565b610de1610d828a60e00151611565565b60405160200180806709ac2c6d0d2dcca560c31b81525060080189805190602001908083835b60208310610e265780518252601f199092019160209182019101610e07565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528a516003909101928b0191508083835b60208310610e7d5780518252601f199092019160209182019101610e5e565b51815160209384036101000a60001901801990921691161790526216100560e91b9190930190815289516003909101928a0191508083835b60208310610ed45780518252601f199092019160209182019101610eb5565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528851600390910192890191508083835b60208310610f2b5780518252601f199092019160209182019101610f0c565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528751600390910192880191508083835b60208310610f825780518252601f199092019160209182019101610f63565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528651600390910192870191508083835b60208310610fd95780518252601f199092019160209182019101610fba565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528551600390910192860191508083835b602083106110305780518252601f199092019160209182019101611011565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528451600390910192850191508083835b602083106110875780518252601f199092019160209182019101611068565b6001836020036101000a0380198251168184511680821785525050505050509050018061148560f11b8152506002019850505050505050505060405160208183030381529060405290505b919050565b6000806110e2614b62565b60006110f18560200151610d63565b9061113d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107a9578181015183820152602001610791565b508360c00151846060015160405180608001604052806111608860200151612100565b81526020016111728860400151612100565b8152602001876080015181526020018760a001518152509250925092509193909250565b60006111a0614b80565b835183106111e6576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806111f386866121da565b915091506111ff612201565b60ff168160ff1614156112335760006112188784612206565b90935090508261122782612274565b94509450505050611351565b61123b612334565b60ff168160ff16141561125d576112528683612339565b935093505050611351565b6112656123db565b60ff168160ff16141561128d57600061127e8784612206565b909350905082611227826123e0565b6112956124cc565b60ff168160ff1614156112ac5761125286836124d1565b6112b4612566565b60ff168160ff16101580156112d557506112cc61256b565b60ff168160ff16105b156113115760006112e4612566565b8203905060606112f5828986612570565b90945090508361130482612618565b9550955050505050611351565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b6000611362614bbd565b61136a614bbd565b600061010082018190528061137f8787612206565b909650915061138e87876124d1565b6020850152955061139f87876124d1565b604085015295506113b08787611196565b606085015295506113c18787611196565b608085015295506113d28787612206565b60a085015295506113e38787612206565b90965090506113f28787611196565b60e085015291835260c08301529590945092505050565b611411614bbd565b60405180610120016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e0015181526020018361010001518152509050919050565b611485614b80565b6040805160608101825260ff8516815260208082018590528251600080825291810184526114d8938301916114d0565b6114bd614b80565b8152602001906001900390816114b55790505b509052612759565b90505b92915050565b6114e9614b80565b604080516001808252818301909252606091816020015b611508614b80565b815260200190600190039081611500579050509050828160008151811061152b57fe5b602002602001018190525061155c60405180606001604052808760ff16815260200186815260200183815250612759565b95945050505050565b600061156f612201565b60ff16826080015160ff16141561159257815161158b906127e9565b90506110d2565b61159a612334565b60ff16826080015160ff1614156115b85761158b826020015161280d565b6115c06124cc565b60ff16826080015160ff1614156115e257815160a083015161158b9190612902565b6115ea612566565b60ff16826080015160ff16141561162357611603614b80565b6116108360400151612950565b905061161b81611565565b9150506110d2565b61162b612ac5565b60ff16826080015160ff161415611644575080516110d2565b61164c6123db565b60ff16826080015160ff161415611691575060608082015160408051607b602080830191909152818301939093528151808203830181529301905281519101206110d2565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6116e0826020015182612aca565b82602001819052505050565b6116fa826040015182612aca565b82604001819052505050565b60008080612d8c600185148061171c5750600285145b806117275750600385145b156117415750600292506000915060039050612b48611e7e565b60048514806117505750600685145b1561176a5750600292506000915060049050612da7611e7e565b60058514806117795750600785145b156117935750600292506000915060079050612da7611e7e565b60088514806117a25750600985145b156117bc5750600392506000915060049050612e77611e7e565b600a8514156117da5750600292506000915060199050612b48611e7e565b600b8514156117f85750600292506000915060079050612b48611e7e565b60108514806118075750601185145b806118125750601285145b8061181d5750601385145b806118285750601685145b806118335750601785145b8061183e5750601885145b1561185757506002925060009150829050612b48611e7e565b601485141561187457506002925060009150829050612f60611e7e565b601585141561189157506001925060009150829050612fb6611e7e565b60198514156118ae57506001925060009150829050613006611e7e565b601a8514806118bd5750601b85145b806118c85750601c85145b806118d35750601d85145b156118ed5750600292506000915060049050612b48611e7e565b603085141561190a5750600192506000915082905061304a611e7e565b60318514156119275750600092508291506001905061305b611e7e565b603285141561194457506000925082915060019050613071611e7e565b60338514156119625750600192506000915060029050613087611e7e565b603485141561198057506001925060009150600490506130a0611e7e565b603585141561199e57506002925060009150600490506130e0611e7e565b60368514156119bb57506000925082915060029050613152611e7e565b60378514156119d857506000925082915060019050613194611e7e565b60388514156119f5575060019250600091508290506131af611e7e565b6039851415611a12575060009250600191508190506131c5611e7e565b603a851415611a2f575060009250829150600290506131db611e7e565b603b851415611a4c57506000925082915060019050612da4611e7e565b603c851415611a6957506000925082915060019050613209611e7e565b603d851415611a8657506001925060009150829050613224611e7e565b6040851415611aa357506001925060009150829050613267611e7e565b6041851415611ac1575060029250600091506001905061329a611e7e565b6042851415611adf57506003925060009150600190506132f2611e7e565b6043851415611afd575060029250600091506001905061336f611e7e565b6044851415611b1b57506003925060009150600190506133ab611e7e565b6050851415611b385750600292506000915082905061340c611e7e565b6051851415611b5657506003925060009150602890506134a6611e7e565b6052851415611b745750600192506000915060029050613566611e7e565b6053851415611b91575060019250829150600390506135af611e7e565b6054851415611baf5750600292506001915060299050613630611e7e565b6060851415611bcc57506000925082915060649050612da4611e7e565b6061851415611bea57506001925060009150606490506136eb611e7e565b6071851415611c085750600192506000915060289050613731611e7e565b6072851415611c25575060009250829150602890506137ac611e7e565b6073851415611c425750600092508291506005905061380a611e7e565b6074851415611c5f575060009250829150600a9050613813611e7e565b6075851415611c7c57506001925060009150829050613820611e7e565b6076851415611c995750600092508291506001905061385a611e7e565b6077851415611cb657506000925082915060199050613873611e7e565b6078851415611cd457506002925060009150601990506138c3611e7e565b6079851415611cf25750600392506000915060199050613938611e7e565b607b851415611d1057506001925060009150600a90506139c5611e7e565b6080851415611d2f57506004925060009150614e209050613a39611e7e565b6081851415611d4e57506004925060009150610dac9050613bb5611e7e565b6082851415611d6e57506003925060009150620140509050613cec611e7e565b6083851415611d8d575060019250600091506103e89050613deb611e7e565b6090851415611daa5750600192506000915082905061304a611e7e565b60a0851415611dc757506000925082915060019050614131611e7e565b60208510801590611dd9575060248511155b15611e155760405162461bcd60e51b815260040180806020018281038252602e815260200180614d09602e913960400191505060405180910390fd5b60a18510801590611e27575060a68511155b80611e325750607085145b15611e6e5760405162461bcd60e51b815260040180806020018281038252602d815260200180614d37602d913960400191505060405180910390fd5b50600092508291506005905061380a5b9193509193565b611e8d614b80565b60408051600080825260208201909252611ec791611ec1565b611eae614b80565b815260200190600190039081611ea65790505b50612618565b905090565b6000816001600160401b0316836040015160a001511015611f11575060c0820180516005016001600160401b03169052604082015160001960a09091015260016114db565b5060c0820180516001600160401b039083018116909152604083015160a0018051918316909103905260006114db565b60016101c090910152565b600161010090910152565b60408051818152606081810183529182919060208201818036833701905050905060005b602081101561201f576000848260208110611f9257fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b611fb88261416b565b858560020281518110611fc757fe5b60200101906001600160f81b031916908160001a905350611fe78161416b565b858560020260010181518110611ff957fe5b60200101906001600160f81b031916908160001a9053505060019092019150611f7b9050565b5092915050565b6060818061204d5750506040805180820190915260018152600360fc1b60208201526110d2565b8060005b811561206557600101600a82049150612051565b6060816001600160401b038111801561207d57600080fd5b506040519080825280601f01601f1916602001820160405280156120a8576020820181803683370190505b50905060001982015b84156120f657600a850660300160f81b828280600190039350815181106120d457fe5b60200101906001600160f81b031916908160001a905350600a850494506120b1565b5095945050505050565b600060028261010001511415612118575060006110d2565b6001826101000151141561212e575060016110d2565b8151602083015161213e90611565565b61214b8460400151611565565b6121588560600151611565565b6121658660800151611565565b8660a001518760c0015161217c8960e00151611565565b60405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012090506110d2565b600080826001018484815181106121ed57fe5b016020015190925060f81c90509250929050565b600090565b6000808284511015801561221e575060208385510310155b61225b576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301612269858561419c565b915091509250929050565b61227c614b80565b6040805160c08101825283815281516060810183526000808252602080830182905284518281528082018652939490850193908301916122d2565b6122bf614b80565b8152602001906001900390816122b75790505b5090528152602001600060405190808252806020026020018201604052801561231557816020015b612302614b80565b8152602001906001900390816122fa5790505b5081526000602082018190526040820152600160609091015292915050565b600190565b6000612343614b80565b8260008061234f614b80565b600061235b89866121da565b909550935061236a89866121da565b9095509250600160ff8516141561238b576123858986611196565b90955091505b61239589866141f5565b9095509050600160ff851614156123c057846123b28483856114e1565b965096505050505050611351565b846123cb848361147d565b9650965050505050509250929050565b600c90565b6123e8614b80565b6040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b815260200160006001600160401b038111801561242f57600080fd5b5060405190808252806020026020018201604052801561246957816020015b612456614b80565b81526020019060019003908161244e5790505b509052815260200160006040519080825280602002602001820160405280156124ac57816020015b612499614b80565b8152602001906001900390816124915790505b50815260208101849052600c604082015260016060909101529050919050565b600290565b60006124db614b80565b828451101580156124f0575060408385510310155b61252d576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60008061253a86866141f5565b90945091506125498685612206565b9094509050836125598383614206565b9350935050509250929050565b600390565b600d90565b60006060828160ff87166001600160401b038111801561258f57600080fd5b506040519080825280602002602001820160405280156125c957816020015b6125b6614b80565b8152602001906001900390816125ae5790505b50905060005b8760ff168160ff16101561260b576125e78784611196565b838360ff16815181106125f657fe5b602090810291909101015292506001016125cf565b5090969095509350505050565b612620614b80565b61262a82516142c5565b61267b576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156126b25783818151811061269557fe5b602002602001015160a00151820191508080600101915050612680565b506040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b815260200160006001600160401b03811180156126fa57600080fd5b5060405190808252806020026020018201604052801561273457816020015b612721614b80565b8152602001906001900390816127195790505b5090528152602081019490945260006040850152600360608501526080909301525090565b612761614b80565b6040518060c001604052806000815260200183815260200160006001600160401b038111801561279057600080fd5b506040519080825280602002602001820160405280156127ca57816020015b6127b7614b80565b8152602001906001900390816127af5790505b5081526000602082015260016040820181905260609091015292915050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061281e57fe5b60408201515161288157612830612334565b82600001518360200151604051602001808460ff1660f81b81526001018360ff1660f81b815260010182815260200193505050506040516020818303038152906040528051906020012090506110d2565b612889612334565b82600001516128af84604001516000815181106128a257fe5b6020026020010151611565565b8460200151604051602001808560ff1660f81b81526001018460ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b600061290c612566565b8383604051602001808460ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b612958614b80565b6008825111156129a6576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516001600160401b03811180156129bf57600080fd5b506040519080825280602002602001820160405280156129e9578160200160208202803683370190505b508051909150600160005b82811015612a4c57612a0b8682815181106128a257fe5b848281518110612a1757fe5b602002602001018181525050858181518110612a2f57fe5b602002602001015160a001518201915080806001019150506129f4565b506000835184604051602001808360ff1660f81b8152600101828051906020019060200280838360005b83811015612a8e578181015183820152602001612a76565b5050505090500192505050604051602081830303815290604052805190602001209050612abb8183614206565b9695505050505050565b606490565b612ad2614b80565b6040805160028082526060828101909352816020015b612af0614b80565b815260200190600190039081612ae85790505090508281600081518110612b1357fe5b60200260200101819052508381600181518110612b2c57fe5b6020026020010181905250612b4081612950565b949350505050565b612b50614b80565b612b5d8260e001516142cc565b9050612b67614b80565b612b748360e001516142cc565b9050612b7f8261430e565b1580612b915750612b8f8161430e565b155b15612ba657612b9f83614319565b5050612da4565b8151815161014085015160009060ff1660011415612bc75750818101612d8e565b61014086015160ff1660021415612be15750818102612d8e565b61014086015160ff1660031415612bfb5750808203612d8e565b61014086015160ff16600a1415612c15575080820a612d8e565b61014086015160ff16600b1415612c2f575080820b612d8e565b61014086015160ff1660101415612c495750808210612d8e565b61014086015160ff1660111415612c635750808211612d8e565b61014086015160ff1660121415612c7d5750808212612d8e565b61014086015160ff1660131415612c975750808213612d8e565b61014086015160ff1660161415612cb15750818116612d8e565b61014086015160ff1660171415612ccb5750818117612d8e565b61014086015160ff1660181415612ce55750818118612d8e565b61014086015160ff16601a1415612cff575080821a612d8e565b61014086015160ff16601b1415612d19575080821b612d8e565b61014086015160ff16601c1415612d33575080821c612d8e565b61014086015160ff16601d1415612d4d575080821d612d8e565b61014086015160ff1660221415612d8c575060408051602080820185905281830184905282518083038401815260609092019092528051910120612d8e565bfe5b610d5b8660e00151612d9f83612274565b614322565b50565b612daf614b80565b612dbc8260e001516142cc565b9050612dc6614b80565b612dd38360e001516142cc565b9050612dde8261430e565b1580612df05750612dee8161430e565b155b80612dfa57508051155b15612e0857612b9f83614319565b8151815161014085015160009060ff1660041415612e295750808204612d8e565b61014086015160ff1660051415612e435750808205612d8e565b61014086015160ff1660061415612e5d5750808206612d8e565b61014086015160ff1660071415612d8c5750808207612d8e565b612e7f614b80565b612e8c8260e001516142cc565b9050612e96614b80565b612ea38360e001516142cc565b9050612ead614b80565b612eba8460e001516142cc565b9050612ec58361430e565b1580612ed75750612ed58261430e565b155b80612ee85750612ee68161430e565b155b80612ef257508051155b15612f0857612f0084614319565b505050612da4565b82518251825161014087015160009060ff1660081415612f2d57818385089050612f45565b61014088015160ff1660091415612d8c578183850990505b612f568860e00151612d9f83612274565b5050505050505050565b612f68614b80565b612f758260e001516142cc565b9050612f7f614b80565b612f8c8360e001516142cc565b9050612fb18360e00151612d9f612fa284611565565b612fab86611565565b1461434c565b505050565b612fbe614b80565b612fcb8260e001516142cc565b9050612fd68161430e565b612fe957612fe382614319565b50612da4565b805160e083015181159061300090612d9f83612274565b50505050565b61300e614b80565b61301b8260e001516142cc565b90506130268161430e565b61303357612fe382614319565b805160e083015181199061300090612d9f83612274565b6130578160e001516142cc565b5050565b612da48160e00151826040015160800151614322565b612da48160e00151826040015160600151614322565b6130948160e001516142cc565b60409091015160600152565b6130a8614b80565b6130b58260e001516142cc565b90506130c08161436e565b6130cd57612fe382614319565b6130d681611565565b6040830151525050565b6130e8614b80565b6130f58260e001516142cc565b90506130ff614b80565b61310c8360e001516142cc565b90506131178261436e565b158061312957506131278161430e565b155b1561313757612b9f83614319565b805115612fb15761314782611565565b604084015152505050565b60e081015151600090158015613181575061316e610abe611e85565b61317f836040015160200151611565565b145b90506130578260e00151612d9f8361434c565b612da48160e00151612d9f836020015160000151600161437b565b612da4816101000151612d9f8360e001516142cc565b612da48160e00151612d9f8361010001516142cc565b6101008101515160009015801561318157506131f8610abe611e85565b61317f836040015160400151611565565b612da48160e00151612d9f836040015160c00151600161437b565b61322c614b80565b6132398260e001516142cc565b90506132448161436e565b61325157612fe382614319565b61325a81611565565b604083015160c001525050565b61326f614b80565b61327c8260e001516142cc565b905061328c8260e0015182614322565b6130578260e0015182614322565b6132a2614b80565b6132af8260e001516142cc565b90506132b9614b80565b6132c68360e001516142cc565b90506132d68360e0015182614322565b6132e48360e0015183614322565b612fb18360e0015182614322565b6132fa614b80565b6133078260e001516142cc565b9050613311614b80565b61331e8360e001516142cc565b9050613328614b80565b6133358460e001516142cc565b90506133458460e0015182614322565b6133538460e0015183614322565b6133618460e0015184614322565b6130008460e0015182614322565b613377614b80565b6133848260e001516142cc565b905061338e614b80565b61339b8360e001516142cc565b90506132e48360e0015183614322565b6133b3614b80565b6133c08260e001516142cc565b90506133ca614b80565b6133d78360e001516142cc565b90506133e1614b80565b6133ee8460e001516142cc565b90506133fe8460e0015184614322565b6133618460e0015183614322565b613414614b80565b6134218260e001516142cc565b905061342b614b80565b6134388360e001516142cc565b90506134438261430e565b158061345557506134538161443a565b155b8061346f575061346481614447565b60ff16826000015110155b1561347d57612b9f83614319565b612fb18360e00151826040015184600001518151811061349957fe5b6020026020010151614322565b6134ae614b80565b6134bb8260e001516142cc565b90506134c5614b80565b6134d28360e001516142cc565b90506134dc614b80565b6134e98460e001516142cc565b90506134f48361430e565b158061350657506135048261443a565b155b80613520575061351582614447565b60ff16836000015110155b1561352e57612f0084614319565b60408201518351815183918391811061354357fe5b602002602001018190525061355f8560e00151612d9f83612618565b5050505050565b61356e614b80565b61357b8260e001516142cc565b90506135868161443a565b61359357612fe382614319565b6130578260e00151612d9f6135a784614447565b60ff16612274565b6135b7614b80565b6135c48260e001516142cc565b90506135ce614b80565b6135dc8361010001516142cc565b90506135e78261430e565b15806135f957506135f78161443a565b155b80613613575061360881614447565b60ff16826000015110155b1561362157612b9f83614319565b61347d83610100015182614322565b613638614b80565b6136458260e001516142cc565b905061364f614b80565b61365c8360e001516142cc565b9050613666614b80565b6136748461010001516142cc565b905061367f8161443a565b1580613691575061368f8361430e565b155b806136ab57506136a081614447565b60ff16836000015110155b156136b957612f0084614319565b6040810151835181518491839181106136ce57fe5b602002602001018190525061355f856101000151612d9f83612618565b8060a00151613700610abe8360e001516142cc565b604080516020808201949094528082019290925280518083038201815260609092019052805191012060a090910152565b613739614b80565b6137468260e001516142cc565b9050613753610abe611e85565b613764836040015160e00151611565565b1461377c576137728261446e565b604083015160e001525b6130578260e00151612d9f61379084611565565b612fab866040015160e00151604001516001815181106128a257fe5b6137b7610abe611e85565b6137c8826040015160e00151611565565b146137f9576137e38160e00151826040015160e00151614322565b6137eb611e85565b604082015160e00152612da4565b612da48160e00151612d9f8361446e565b612da481614319565b612da4816040015161482f565b613828614b80565b6138358260e001516142cc565b90506138408161430e565b61384d57612fe382614319565b51604082015160a0015250565b612da48160e00151612d9f836040015160a00151612274565b60e081015160408051600160f81b6020808301919091526000602183018190526022808401919091528351808403909101815260429092019092528051910120612da49190612d9f90600161437b565b6138cb614b80565b6138d88260e001516142cc565b90506138e2614b80565b6138ef8360e001516142cc565b90506138fa8261430e565b158061390c575061390a8161436e565b155b1561391a57612b9f83614319565b612fb18360e00151612d9f846000015161393385611565565b61147d565b613940614b80565b61394d8260e001516142cc565b9050613957614b80565b6139648360e001516142cc565b905061396e614b80565b61397b8460e001516142cc565b90506139868361430e565b158061399857506139968161436e565b155b156139a657612f0084614319565b6130008460e00151612d9f85600001516139bf85611565565b866114e1565b6139cd614b80565b6139da8260e001516142cc565b90506139e58161430e565b6139f257612fe382614319565b60408051600080825260208201909252606091613a25565b613a12614b80565b815260200190600190039081613a0a5790505b509050612fb18360e00151612d9f83612618565b613a41614b80565b613a4e8260e001516142cc565b9050613a58614b80565b613a658360e001516142cc565b9050613a6f614b80565b613a7c8460e001516142cc565b9050613a86614b80565b613a938560e001516142cc565b9050613a9e8461430e565b1580613ab05750613aae8361430e565b155b80613ac15750613abf8261430e565b155b80613ad25750613ad08161430e565b155b15613ae957613ae085614319565b50505050612da4565b83518351835115801590613aff57508351600114155b15613b2157613b168760e00151612d9f6000612274565b505050505050612da4565b83518351604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015613b83573d6000803e3d6000fd5b505050602060405103519050613ba98a60e00151612d9f836001600160a01b0316612274565b50505050505050505050565b613bbd614b80565b613bca8260e001516142cc565b9050613bd4614b80565b613be18360e001516142cc565b9050613beb614b80565b613bf88460e001516142cc565b9050613c02614b80565b613c0f8560e001516142cc565b9050613c1a8461430e565b1580613c2c5750613c2a8361430e565b155b80613c3d5750613c3b8261430e565b155b80613c4e5750613c4c8161430e565b155b15613c5c57613ae085614319565b613c64614b62565b5060408051608081018252855181528451602082015283519181019190915281516060820152613c92614c28565b600060408260808560066107d05a03fa905080613cbe57613cb288614319565b50505050505050612da4565b60e0880151613cd890612d9f8460015b6020020151612274565b60e0880151612f5690612d9f846000613cce565b613cf4614b80565b613d018260e001516142cc565b9050613d0b614b80565b613d188360e001516142cc565b9050613d22614b80565b613d2f8460e001516142cc565b9050613d3a8361430e565b1580613d4c5750613d4a8261430e565b155b80613d5d5750613d5b8161430e565b155b15613d6b57612f0084614319565b613d73614c46565b50604080516060810182528451815283516020820152825191810191909152613d9a614c28565b600060408260808560076107d05a03fa905080613dba57613b1687614319565b60e0870151613dce90612d9f846001613cce565b60e0870151613de290612d9f846000613cce565b50505050505050565b613df3614b80565b613e008260e001516142cc565b9050613e0a614c64565b6000805b601e811015613e9c57613e208461443a565b613e2d5760019150613e9c565b60408401518051613e3e5750613e9c565b8051600214613e51576001925050613e9c565b80600081518110613e5e57fe5b60200260200101518483601e8110613e7257fe5b6020020152805181906001908110613e8657fe5b6020908102919091010151945050600101613e0e565b613eab856207a1208302611ecc565b15613ecf5760c0850180516103e719016001600160401b03169052613ae085611f41565b8180613ee15750613edf8461443a565b155b80613ef0575060408401515115155b15613efe57613ae085614319565b613f06614c92565b60005b828110156140e357613f19614b80565b8582601e8110613f2557fe5b60200201519050613f358161443a565b613f4257613cb288614319565b60408101518051600614613f6657613f5989614319565b5050505050505050612da4565b60005b6006811015613fb157613f8e828281518110613f8157fe5b602002602001015161430e565b613fa957613f9b8a614319565b505050505050505050612da4565b600101613f69565b5080600081518110613fbf57fe5b602002602001015160000151848460060260b48110613fda57fe5b6020020152805181906001908110613fee57fe5b602002602001015160000151848460060260010160b4811061400c57fe5b602002015280518190600390811061402057fe5b602002602001015160000151848460060260020160b4811061403e57fe5b602002015280518190600290811061405257fe5b602002602001015160000151848460060260030160b4811061407057fe5b602002015280518190600590811061408457fe5b602002602001015160000151848460060260040160b481106140a257fe5b60200201528051819060049081106140b657fe5b602002602001015160000151848460060260050160b481106140d457fe5b60200201525050600101613f09565b5060c082026140f0614cb1565b6000602082848660086107d05a03fa90508061410f57613f5989614319565b60e089015182516141269190612d9f90151561434c565b505050505050505050565b612da48160e00151612d9f6000801b60405160200180828152602001915050604051602081830303815290604052805190602001206123e0565b6000600a60f883901c101561418b578160f81c60300160f81b90506110d2565b8160f81c60570160f81b90506110d2565b600081602001835110156141ec576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60008060208301612269858561419c565b61420e614b80565b6040805160c0810182528481528151606081018352600080825260208083018290528451828152808201865293949085019390830191614264565b614251614b80565b8152602001906001900390816142495790505b509052815260200160006040519080825280602002602001820160405280156142a757816020015b614294614b80565b81526020019060019003908161428c5790505b50815260006020820152600260408201526060019290925250919050565b6008101590565b6142d4614b80565b6142dc614b80565b82602001516001846000015103815181106142f357fe5b60209081029190910101518351600019018452915050919050565b6080015160ff161590565b612da481611f41565b80826020015183600001518151811061433757fe5b60209081029190910101525080516001019052565b614354614b80565b81156143645761158b6001612274565b61158b6000612274565b6080015160ff1660011490565b614383614b80565b6040805160c08101825284815281516060810183526000808252602080830182905284518281528082018652939490850193908301916143d9565b6143c6614b80565b8152602001906001900390816143be5790505b5090528152602001600060405190808252806020026020018201604052801561441c57816020015b614409614b80565b8152602001906001900390816144015790505b50815260006020820152606460408201526060019290925250919050565b6080015160ff1660031490565b608081015160009060ff166003141561446657506040810151516110d2565b5060016110d2565b614476614b80565b610160820151604080516008808252610120820190925260009160609190816020015b6144a1614b80565b815260200190600190039081614499579050509050600083866101800151815181106144c957fe5b0160200151610180870180516001019081905260f89190911c915060009081908190819081906144fa908a9061483a565b6101808c01805160140190819052909150614516908a90612206565b6101808d01829052955061452b908a90612206565b6101808d018290529450614540908a90612206565b6101808d018290529350614555908a90612206565b6101808d01829052925060009061456d908b90612206565b6101808e018290529150600090614586908c908461489a565b6101808e01518c8101602001849020919250906145a88a868b8b8b8b876148cc565b9b506145b68a60ff16612274565b8b6000815181106145c357fe5b60200260200101819052506145d789612274565b8b6001815181106145e457fe5b60200260200101819052506145f888612274565b8b60028151811061460557fe5b6020026020010181905250614622856001600160a01b0316612274565b8b60038151811061462f57fe5b602002602001018190525061464387612274565b8b60048151811061465057fe5b602002602001018190525061466486612274565b8b60058151811061467157fe5b602002602001018190525061468584612274565b8b60068151811061469257fe5b60200260200101819052506146a883600161437b565b8b6007815181106146b557fe5b602090810291909101015250505060608c0151600097501595506147579450505050505785600001516001600160a01b031663d9dd67ab60018860600151036040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561472857600080fd5b505afa15801561473c573d6000803e3d6000fd5b505050506040513d602081101561475257600080fd5b505190505b85600001516001600160a01b031663d9dd67ab87606001516040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156147a357600080fd5b505afa1580156147b7573d6000803e3d6000fd5b505050506040513d60208110156147cd57600080fd5b50516147d98285614942565b1461481b576040805162461bcd60e51b815260206004820152600d60248201526c57524f4e475f4d45535341474560981b604482015290519081900360640190fd5b6060860180516001019052612abb82612618565b600261010090910152565b6000816014018351101561488a576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b500160200151600160601b900490565b6000806148b285856148ab8661496e565b6001614999565b50855190915061155c906148c7607b84614942565b614942565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6bffffffffffffffffffffffff19166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600060018211614980575060016110d2565b61498f6002600184010461496e565b60020290506110d2565b600080602084116149ed57855185106149c1576149b660006127e9565b600191509150614a64565b60006149d56149d08888614a6d565b6127e9565b9050806149e260006127e9565b909350149050614a64565b600080614a06886002880489016002895b046000614999565b91509150808015614a145750845b15614a3357614a2888886002890488614999565b935093505050614a64565b600080614a438a8a60028b6149fe565b91509150614a518285614942565b818015614a5b5750835b95509550505050505b94509492505050565b600080805b6020811015614ac457600882901b91506000818501865111614a95576000614ab3565b8582860181518110614aa357fe5b01602001516001600160f81b0319165b60f81c929092179150600101614a72565b509392505050565b604051806101e0016040528060006001600160a01b03168152602001614af0614bbd565b8152602001614afd614bbd565b81526000602082018190526040820181905260608201819052608082015260a001614b26614ccf565b8152602001614b33614ccf565b81526000602082018190526040820181905260608083018190526080830182905260a083015260c09091015290565b60405180608001604052806004906020820280368337509192915050565b6040518060c0016040528060008152602001614b9a614ce9565b815260606020820181905260006040830181905290820181905260809091015290565b6040805161012081019091526000815260208101614bd9614b80565b8152602001614be6614b80565b8152602001614bf3614b80565b8152602001614c00614b80565b81526000602082018190526040820152606001614c1b614b80565b8152602001600081525090565b60405180604001604052806002906020820280368337509192915050565b60405180606001604052806003906020820280368337509192915050565b604051806103c00160405280601e905b614c7c614b80565b815260200190600190039081614c745790505090565b60405180611680016040528060b4906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b604051806040016040528060008152602001606081525090565b604080516060808201835260008083526020830152918101919091529056fe75736520616e6f7468657220636f6e747261637420746f2068616e646c652068617368696e67206f70636f64657375736520616e6f7468657220636f6e747261637420746f2068616e646c6520627566666572206f70636f646573a264697066735822122058c67d800c30a3e44079032496dc68ceadfe32d5bd3aebce5dd2a29810ffa50f64736f6c634300060c0033"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// OneStepProof is an auto generated Go binding around an Ethereum contract.
type OneStepProof struct {
	OneStepProofCaller     // Read-only binding to the contract
	OneStepProofTransactor // Write-only binding to the contract
	OneStepProofFilterer   // Log filterer for contract events
}

// OneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofSession struct {
	Contract     *OneStepProof     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofCallerSession struct {
	Contract *OneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofTransactorSession struct {
	Contract     *OneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofRaw struct {
	Contract *OneStepProof // Generic contract binding to access the raw methods on
}

// OneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofCallerRaw struct {
	Contract *OneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofTransactorRaw struct {
	Contract *OneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof creates a new instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProof(address common.Address, backend bind.ContractBackend) (*OneStepProof, error) {
	contract, err := bindOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// NewOneStepProofCaller creates a new read-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofCaller, error) {
	contract, err := bindOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofCaller{contract: contract}, nil
}

// NewOneStepProofTransactor creates a new write-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofTransactor, error) {
	contract, err := bindOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTransactor{contract: contract}, nil
}

// NewOneStepProofFilterer creates a new log filterer instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofFilterer, error) {
	contract, err := bindOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofFilterer{contract: contract}, nil
}

// bindOneStepProof binds a generic wrapper to an already deployed contract.
func bindOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.OneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transact(opts, method, params...)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x9d16dd04.
//
// Solidity: function executeStep(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields)
func (_OneStepProof *OneStepProofCaller) ExecuteStep(opts *bind.CallOpts, bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	var out []interface{}
	err := _OneStepProof.contract.Call(opts, &out, "executeStep", bridge, initialMessagesRead, accs, proof, bproof)

	outstruct := new(struct {
		Gas               uint64
		TotalMessagesRead *big.Int
		Fields            [4][32]byte
	})

	outstruct.Gas = out[0].(uint64)
	outstruct.TotalMessagesRead = out[1].(*big.Int)
	outstruct.Fields = out[2].([4][32]byte)

	return *outstruct, err

}

// ExecuteStep is a free data retrieval call binding the contract method 0x9d16dd04.
//
// Solidity: function executeStep(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields)
func (_OneStepProof *OneStepProofSession) ExecuteStep(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _OneStepProof.Contract.ExecuteStep(&_OneStepProof.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x9d16dd04.
//
// Solidity: function executeStep(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields)
func (_OneStepProof *OneStepProofCallerSession) ExecuteStep(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _OneStepProof.Contract.ExecuteStep(&_OneStepProof.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x2ccebb7a.
//
// Solidity: function executeStepDebug(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_OneStepProof *OneStepProofCaller) ExecuteStepDebug(opts *bind.CallOpts, bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	var out []interface{}
	err := _OneStepProof.contract.Call(opts, &out, "executeStepDebug", bridge, initialMessagesRead, accs, proof, bproof)

	outstruct := new(struct {
		StartMachine string
		AfterMachine string
	})

	outstruct.StartMachine = out[0].(string)
	outstruct.AfterMachine = out[1].(string)

	return *outstruct, err

}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x2ccebb7a.
//
// Solidity: function executeStepDebug(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_OneStepProof *OneStepProofSession) ExecuteStepDebug(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _OneStepProof.Contract.ExecuteStepDebug(&_OneStepProof.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x2ccebb7a.
//
// Solidity: function executeStepDebug(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_OneStepProof *OneStepProofCallerSession) ExecuteStepDebug(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _OneStepProof.Contract.ExecuteStepDebug(&_OneStepProof.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}
