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

// OneStepProof2ABI is the input ABI used to generate the binding from.
const OneStepProof2ABI = "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"totalMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[4]\",\"name\":\"fields\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStepDebug\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"startMachine\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"afterMachine\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"parseProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProof2FuncSigs maps the 4-byte function signature to its string representation.
var OneStepProof2FuncSigs = map[string]string{
	"9d16dd04": "executeStep(address,uint256,bytes32[2],bytes,bytes)",
	"2ccebb7a": "executeStepDebug(address,uint256,bytes32[2],bytes,bytes)",
	"793deea3": "parseProof(bytes)",
}

// OneStepProof2Bin is the compiled bytecode used for deploying new contracts.
var OneStepProof2Bin = "0x608060405234801561001057600080fd5b506141e8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632ccebb7a14610046578063793deea3146101fd5780639d16dd04146103c4575b600080fd5b61011f600480360360c081101561005c57600080fd5b6001600160a01b038235169160208101359160408201919081019060a081016080820135600160201b81111561009157600080fd5b8201836020820111156100a357600080fd5b803590602001918460018302840111600160201b831117156100c457600080fd5b919390929091602081019035600160201b8111156100e157600080fd5b8201836020820111156100f357600080fd5b803590602001918460018302840111600160201b8311171561011457600080fd5b5090925090506104ef565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610160578181015183820152602001610148565b50505050905090810190601f16801561018d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156101c05781810151838201526020016101a8565b50505050905090810190601f1680156101ed5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6102a16004803603602081101561021357600080fd5b810190602081018135600160201b81111561022d57600080fd5b82018360208201111561023f57600080fd5b803590602001918460018302840111600160201b8311171561026057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105b4945050505050565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b838110156102ed5781810151838201526020016102d5565b50505050905001858103845288818151815260200191508051906020019060200280838360005b8381101561032c578181015183820152602001610314565b50505050905001858103835287818151815260200191508051906020019060200280838360005b8381101561036b578181015183820152602001610353565b50505050905001858103825286818151815260200191508051906020019060200280838360005b838110156103aa578181015183820152602001610392565b505050509050019850505050505050505060405180910390f35b61049d600480360360c08110156103da57600080fd5b6001600160a01b038235169160208101359160408201919081019060a081016080820135600160201b81111561040f57600080fd5b82018360208201111561042157600080fd5b803590602001918460018302840111600160201b8311171561044257600080fd5b919390929091602081019035600160201b81111561045f57600080fd5b82018360208201111561047157600080fd5b803590602001918460018302840111600160201b8311171561049257600080fd5b5090925090506105f0565b60405180846001600160401b0316815260200183815260200182600460200280838360005b838110156104da5781810151838201526020016104c2565b50505050905001935050505060405180910390f35b6060806104fa613fc7565b61057e898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8d018190048102820181019092528b815292508b91508a9081908401838280828437600081840152601f19601f820116905080830192505050505050508e6106af565b905061058981610b63565b6105968160200151610f71565b92506105a58160400151610f71565b91505097509795505050505050565b6060806060806105c261405d565b6105cb866112e5565b80516020820151604083015160609093015191975095509093509150505b9193509193565b6000806105fb614085565b610603613fc7565b6106878a8a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b9081908401838280828437600081840152601f19601f820116905080830192505050505050508f6106af565b905061069281610b63565b61069b816113db565b935093509350509750975097945050505050565b6106b7613fc7565b6000846000815181106106c657fe5b602001015160f81c60f81b60f81c90506000856001815181106106e557fe5b602001015160f81c60f81b60f81c905060008660028151811061070457fe5b016020015160f81c9050600360606004840160ff166001600160401b038111801561072e57600080fd5b5060405190808252806020026020018201604052801561076857816020015b6107556140a3565b81526020019060019003908161074d5790505b50905060608360040160ff166001600160401b038111801561078957600080fd5b506040519080825280602002602001820160405280156107c357816020015b6107b06140a3565b8152602001906001900390816107a85790505b50905060005b8560ff168110156107ff576107de8b8561149a565b8483815181106107ea57fe5b602090810291909101015293506001016107c9565b5060005b8460ff16811015610839576108188b8561149a565b83838151811061082457fe5b60209081029190910101529350600101610803565b506108426140e0565b61084c8b8561165c565b809250819550505060008b858151811061086257fe5b01602001516001959095019460f81c905061087b613fc7565b6001600160a01b038b168152602081018390526108978361170d565b6040820152606081018f90528d6000602002013560808201528d60016020908102919091013560a0830152600060c0830181905260408051808201825260ff8c811682528185018a905260e086019190915281518083019092528a8116825292810187905261010084015283821660018114610120850152918b1661014084015261016083018f90526101a083018e90526101c08301526101808201879052158061094557508160ff166001145b6040518060400160405280600b81526020016a04241445f494d4d5f5459560ac1b815250906109f25760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109b757818101518382015260200161099f565b50505050905090810190601f1680156109e45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506109fb6140a3565b60ff8316610a1c57610a158a836020015160000151611781565b9050610abc565b6000865111604051806040016040528060068152602001654e4f5f494d4d60d01b81525090610a8c5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109b757818101518382015260200161099f565b50610ab98a8360200151600001518860018d0360ff1681518110610aac57fe5b60200260200101516117e5565b90505b610ac58161186b565b60208301515260005b838a0360ff16811015610b0d57610b05878281518110610aea57fe5b602002602001015184602001516119d890919063ffffffff16565b600101610ace565b5060005b8860ff16811015610b4e57610b46868281518110610b2b57fe5b602002602001015184602001516119f290919063ffffffff16565b600101610b11565b50909f9e505050505050505050505050505050565b600080600061414b610b7c85610140015160ff16611a0c565b93509350935093506000841180610b965750846101200151155b8015610ba7575060e0850151518410155b80610bce57508461012001518015610bbd575083155b8015610bce575060e0850151516001145b6040518060400160405280600a815260200169535441434b5f4d414e5960b01b81525090610c3d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109b757818101518382015260200161099f565b50610100850151516040805180820190915260088152674155585f4d414e5960c01b602082015290841015610cb35760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109b757818101518382015260200161099f565b5060e085015151841115610d7057610cd1610ccc611b1c565b61186b565b610ce286604001516020015161186b565b146040518060400160405280600d81526020016c535441434b5f4d495353494e4760981b81525090610d555760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109b757818101518382015260200161099f565b50610d61856005611b63565b50610d6b85611bd8565b610e30565b61010085015151831115610e0b57610d89610ccc611b1c565b610d9a86604001516040015161186b565b146040518060400160405280600b81526020016a4155585f4d495353494e4760a81b81525090610d555760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109b757818101518382015260200161099f565b610e158583611b63565b15610e2357610d6b85611bd8565b610e30858263ffffffff16565b846101c0015115610ed25760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201835281519101209086015160c001511415610e9557610e908560400151611be3565b610ed2565b60006101c0860152604085015160c081015190526101208501518015610eb9575083155b610ec75760e0850151600090525b610100850151600090525b60005b60e086015151811015610f1c57610f148660e00151602001518281518110610ef957fe5b602002602001015187604001516119d890919063ffffffff16565b600101610ed5565b5060005b61010086015151811015610f6957610f61866101000151602001518281518110610f4657fe5b602002602001015187604001516119f290919063ffffffff16565b600101610f20565b505050505050565b6060610f808260000151611bee565b610f95610f90846020015161186b565b611bee565b610fa5610f90856040015161186b565b610fb5610f90866060015161186b565b610fc5610f90876080015161186b565b610fd28760a00151611cbd565b610fdf8860c00151611bee565b610fef610f908a60e0015161186b565b60405160200180806709ac2c6d0d2dcca560c31b81525060080189805190602001908083835b602083106110345780518252601f199092019160209182019101611015565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528a516003909101928b0191508083835b6020831061108b5780518252601f19909201916020918201910161106c565b51815160209384036101000a60001901801990921691161790526216100560e91b9190930190815289516003909101928a0191508083835b602083106110e25780518252601f1990920191602091820191016110c3565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528851600390910192890191508083835b602083106111395780518252601f19909201916020918201910161111a565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528751600390910192880191508083835b602083106111905780518252601f199092019160209182019101611171565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528651600390910192870191508083835b602083106111e75780518252601f1990920191602091820191016111c8565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528551600390910192860191508083835b6020831061123e5780518252601f19909201916020918201910161121f565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528451600390910192850191508083835b602083106112955780518252601f199092019160209182019101611276565b6001836020036101000a0380198251168184511680821785525050505050509050018061148560f11b8152506002019850505050505050505060405160208183030381529060405290505b919050565b6112ed61405d565b606061132c838460008151811061130057fe5b602001015160f81c60f81b8560018151811061131857fe5b01602001516001600160f81b031916611d97565b90506060611359848560018151811061134157fe5b602001015160f81c60f81b8660028151811061131857fe5b90506060611386858660028151811061136e57fe5b602001015160f81c60f81b8760038151811061131857fe5b905060606113b3868760038151811061139b57fe5b602001015160f81c60f81b8860048151811061131857fe5b6040805160808101825295865260208601949094529284019190915250606082015292915050565b6000806113e6614085565b60006113f58560200151610f71565b906114415760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109b757818101518382015260200161099f565b508360c00151846060015160405180608001604052806114648860200151611e34565b81526020016114768860400151611e34565b8152602001876080015181526020018760a001518152509250925092509193909250565b60006114a46140a3565b835183106114ea576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806114f78686611f0e565b91509150611503611f35565b60ff168160ff16141561153757600061151c8784611f3a565b90935090508261152b82611fa8565b94509450505050611655565b61153f612068565b60ff168160ff16141561156157611556868361206d565b935093505050611655565b61156961210f565b60ff168160ff1614156115915760006115828784611f3a565b90935090508261152b82612114565b611599612200565b60ff168160ff1614156115b0576115568683612205565b6115b861229a565b60ff168160ff16101580156115d957506115d061229f565b60ff168160ff16105b156116155760006115e861229a565b8203905060606115f98289866122a4565b9094509050836116088261234c565b9550955050505050611655565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b60006116666140e0565b61166e6140e0565b60006101008201819052806116838787611f3a565b90965091506116928787612205565b602085015295506116a38787612205565b604085015295506116b4878761149a565b606085015295506116c5878761149a565b608085015295506116d68787611f3a565b60a085015295506116e78787611f3a565b90965090506116f6878761149a565b60e085015291835260c08301529590945092505050565b6117156140e0565b60405180610120016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e0015181526020018361010001518152509050919050565b6117896140a3565b6040805160608101825260ff8516815260208082018590528251600080825291810184526117dc938301916117d4565b6117c16140a3565b8152602001906001900390816117b95790505b50905261248d565b90505b92915050565b6117ed6140a3565b604080516001808252818301909252606091816020015b61180c6140a3565b815260200190600190039081611804579050509050828160008151811061182f57fe5b602002602001018190525061186060405180606001604052808760ff1681526020018681526020018381525061248d565b9150505b9392505050565b6000611875611f35565b60ff16826080015160ff1614156118985781516118919061251d565b90506112e0565b6118a0612068565b60ff16826080015160ff1614156118be576118918260200151612541565b6118c6612200565b60ff16826080015160ff1614156118e857815160a08301516118919190612636565b6118f061229a565b60ff16826080015160ff161415611929576119096140a3565b6119168360400151612684565b90506119218161186b565b9150506112e0565b6119316127f9565b60ff16826080015160ff16141561194a575080516112e0565b61195261210f565b60ff16826080015160ff161415611997575060608082015160408051607b602080830191909152818301939093528151808203830181529301905281519101206112e0565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6119e68260200151826127fe565b82602001819052505050565b611a008260400151826127fe565b82604001819052505050565b6000808061414b60a1851415611a3157506002925060009150600a905061287c6105e9565b60a2851415611a4f57506002925060009150600a90506129306105e9565b60a3851415611a6d57506002925060009150600a90506129c26105e9565b60a4851415611a8b5750600392506000915060649050612a546105e9565b60a5851415611aa95750600392506000915060649050612b416105e9565b60a6851415611ac75750600392506000915060649050612c126105e9565b6070851415611ae55750600292506000915060649050612cd16105e9565b60405162461bcd60e51b815260040180806020018281038252602c815260200180614187602c913960400191505060405180910390fd5b611b246140a3565b60408051600080825260208201909252611b5e91611b58565b611b456140a3565b815260200190600190039081611b3d5790505b5061234c565b905090565b6000816001600160401b0316836040015160a001511015611ba8575060c0820180516005016001600160401b03169052604082015160001960a09091015260016117df565b5060c0820180516001600160401b039083018116909152604083015160a0018051918316909103905260006117df565b60016101c090910152565b600161010090910152565b60408051818152606081810183529182919060208201818036833701905050905060005b6020811015611cb6576000848260208110611c2957fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b611c4f82612e6b565b858560020281518110611c5e57fe5b60200101906001600160f81b031916908160001a905350611c7e81612e6b565b858560020260010181518110611c9057fe5b60200101906001600160f81b031916908160001a9053505060019092019150611c129050565b5092915050565b60608180611ce45750506040805180820190915260018152600360fc1b60208201526112e0565b8060005b8115611cfc57600101600a82049150611ce8565b6060816001600160401b0381118015611d1457600080fd5b506040519080825280601f01601f191660200182016040528015611d3f576020820181803683370190505b50905060001982015b8415611d8d57600a850660300160f81b82828060019003935081518110611d6b57fe5b60200101906001600160f81b031916908160001a905350600a85049450611d48565b5095945050505050565b606060f883811c9083901c81900360ff169082826001600160401b0381118015611dc057600080fd5b50604051908082528060200260200182016040528015611dea578160200160208202803683370190505b50905060005b83811015611e2957611e0788828501602002612e9c565b60001b828281518110611e1657fe5b6020908102919091010152600101611df0565b509695505050505050565b600060028261010001511415611e4c575060006112e0565b60018261010001511415611e62575060016112e0565b81516020830151611e729061186b565b611e7f846040015161186b565b611e8c856060015161186b565b611e99866080015161186b565b8660a001518760c00151611eb08960e0015161186b565b60405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012090506112e0565b60008082600101848481518110611f2157fe5b016020015190925060f81c90509250929050565b600090565b60008082845110158015611f52575060208385510310155b611f8f576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301611f9d8585612edc565b915091509250929050565b611fb06140a3565b6040805160c0810182528381528151606081018352600080825260208083018290528451828152808201865293949085019390830191612006565b611ff36140a3565b815260200190600190039081611feb5790505b5090528152602001600060405190808252806020026020018201604052801561204957816020015b6120366140a3565b81526020019060019003908161202e5790505b5081526000602082018190526040820152600160609091015292915050565b600190565b60006120776140a3565b826000806120836140a3565b600061208f8986611f0e565b909550935061209e8986611f0e565b9095509250600160ff851614156120bf576120b9898661149a565b90955091505b6120c98986612f35565b9095509050600160ff851614156120f457846120e68483856117e5565b965096505050505050611655565b846120ff8483611781565b9650965050505050509250929050565b600c90565b61211c6140a3565b6040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b815260200160006001600160401b038111801561216357600080fd5b5060405190808252806020026020018201604052801561219d57816020015b61218a6140a3565b8152602001906001900390816121825790505b509052815260200160006040519080825280602002602001820160405280156121e057816020015b6121cd6140a3565b8152602001906001900390816121c55790505b50815260208101849052600c604082015260016060909101529050919050565b600290565b600061220f6140a3565b82845110158015612224575060408385510310155b612261576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60008061226e8686612f35565b909450915061227d8685611f3a565b90945090508361228d8383612f46565b9350935050509250929050565b600390565b600d90565b60006060828160ff87166001600160401b03811180156122c357600080fd5b506040519080825280602002602001820160405280156122fd57816020015b6122ea6140a3565b8152602001906001900390816122e25790505b50905060005b8760ff168160ff16101561233f5761231b878461149a565b838360ff168151811061232a57fe5b60209081029190910101529250600101612303565b5090969095509350505050565b6123546140a3565b61235e8251613005565b6123af576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156123e6578381815181106123c957fe5b602002602001015160a001518201915080806001019150506123b4565b506040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b815260200160006001600160401b038111801561242e57600080fd5b5060405190808252806020026020018201604052801561246857816020015b6124556140a3565b81526020019060019003908161244d5790505b5090528152602081019490945260006040850152600360608501526080909301525090565b6124956140a3565b6040518060c001604052806000815260200183815260200160006001600160401b03811180156124c457600080fd5b506040519080825280602002602001820160405280156124fe57816020015b6124eb6140a3565b8152602001906001900390816124e35790505b5081526000602082015260016040820181905260609091015292915050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061255257fe5b6040820151516125b557612564612068565b82600001518360200151604051602001808460ff1660f81b81526001018360ff1660f81b815260010182815260200193505050506040516020818303038152906040528051906020012090506112e0565b6125bd612068565b82600001516125e384604001516000815181106125d657fe5b602002602001015161186b565b8460200151604051602001808560ff1660f81b81526001018460ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b600061264061229a565b8383604051602001808460ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b61268c6140a3565b6008825111156126da576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516001600160401b03811180156126f357600080fd5b5060405190808252806020026020018201604052801561271d578160200160208202803683370190505b508051909150600160005b828110156127805761273f8682815181106125d657fe5b84828151811061274b57fe5b60200260200101818152505085818151811061276357fe5b602002602001015160a00151820191508080600101915050612728565b506000835184604051602001808360ff1660f81b8152600101828051906020019060200280838360005b838110156127c25781810151838201526020016127aa565b50505050905001925050506040516020818303038152906040528051906020012090506127ef8183612f46565b9695505050505050565b606490565b6128066140a3565b6040805160028082526060828101909352816020015b6128246140a3565b81526020019060019003908161281c579050509050828160008151811061284757fe5b6020026020010181905250838160018151811061286057fe5b602002602001018190525061287481612684565b949350505050565b6128846140a3565b6128918260e0015161300c565b905061289b6140a3565b6128a88360e0015161300c565b90506128b38261304e565b15806128c557506128c38161306c565b155b156128da576128d383613079565b505061292d565b8151600160401b116128ef576128d383613079565b60006129118260600151846000015161290c876101a001516112e5565b613082565b90506129298460e0015161292483611fa8565b6130a4565b5050505b50565b6129386140a3565b6129458260e0015161300c565b905061294f6140a3565b61295c8360e0015161300c565b90506129678261304e565b158061297957506129778161306c565b155b15612987576128d383613079565b815167fffffffffffffff9116129a0576128d383613079565b6000612911826060015184600001516129bd876101a001516112e5565b6130ce565b6129ca6140a3565b6129d78260e0015161300c565b90506129e16140a3565b6129ee8360e0015161300c565b90506129f98261304e565b1580612a0b5750612a098161306c565b155b15612a19576128d383613079565b815167ffffffffffffffe111612a32576128d383613079565b600061291182606001518460000151612a4f876101a001516112e5565b61322d565b612a5c6140a3565b612a698260e0015161300c565b9050612a736140a3565b612a808360e0015161300c565b9050612a8a6140a3565b612a978460e0015161300c565b9050612aa28361304e565b1580612ab45750612ab282613360565b155b80612ac55750612ac38161306c565b155b15612adb57612ad384613079565b50505061292d565b8251600160401b111580612af25750815161010011155b15612b0057612ad384613079565b6000612b27826060015185600001518560000151612b22896101a001516112e5565b61336b565b9050612b3a8560e0015161292483612114565b5050505050565b612b496140a3565b612b568260e0015161300c565b9050612b606140a3565b612b6d8360e0015161300c565b9050612b776140a3565b612b848460e0015161300c565b9050612b8f8361304e565b1580612ba15750612b9f82613360565b155b80612bb25750612bb08161306c565b155b15612bc057612ad384613079565b825167fffffffffffffff9111580612bdd57508151600160401b11155b15612beb57612ad384613079565b6000612b27826060015185600001518560000151612c0d896101a001516112e5565b6133b4565b612c1a6140a3565b612c278260e0015161300c565b9050612c316140a3565b612c3e8360e0015161300c565b9050612c486140a3565b612c558460e0015161300c565b9050612c608361304e565b1580612c725750612c7082613360565b155b80612c835750612c818161306c565b155b15612c9157612ad384613079565b825167ffffffffffffffe111612caa57612ad384613079565b6000612b27826060015185600001518560000151612ccc896101a001516112e5565b6134fd565b612cd96140a3565b612ce68260e0015161300c565b9050612cf06140a3565b612cfd8360e0015161300c565b9050612d088261304e565b1580612d1a5750612d188161306c565b155b15612d28576128d383613079565b81516127101080612d3857508151155b15612d46576128d383613079565b826101600151518361018001511415612dc157612d7981606001518360000151612d74866101a001516112e5565b6135cf565b15612db8576040805162461bcd60e51b815260206004820152600a602482015269084aa8cbe988a9c8ea8960b31b604482015290519081900360640190fd5b6128d383613079565b61018083015182516101608501516000612ddc828585613623565b905080612de88661186b565b14612e27576040805162461bcd60e51b815260206004820152600a60248201526915d493d391d7d4d1539160b21b604482015290519081900360640190fd5b509091016020908101919091206080850180516040805180860192909252818101939093528251808203840181526060909101909252815191909201209052505050565b6000600a60f883901c1015612e8b578160f81c60300160f81b90506112e0565b8160f81c60570160f81b90506112e0565b600080805b6020811015612ed457600882901b91508481850181518110612ebf57fe5b016020015160f81c9190911790600101612ea1565b509392505050565b60008160200183511015612f2c576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60008060208301611f9d8585612edc565b612f4e6140a3565b6040805160c0810182528481528151606081018352600080825260208083018290528451828152808201865293949085019390830191612fa4565b612f916140a3565b815260200190600190039081612f895790505b50905281526020016000604051908082528060200260200182016040528015612fe757816020015b612fd46140a3565b815260200190600190039081612fcc5790505b50815260006020820152600260408201526060019290925250919050565b6008101590565b6130146140a3565b61301c6140a3565b826020015160018460000151038151811061303357fe5b60209081029190910101518351600019018452915050919050565b608081015160009060ff161580156117df57505051600160401b1190565b6080015160ff16600c1490565b61292d81611bd8565b600061287461309a856020865b048560000151613655565b6020855b066137c3565b8082602001518360000151815181106130b957fe5b60209081029190910101525080516001019052565b604080516008808252818301909252600091606091906020820181803683370190505090506000613108866020875b048660000151613655565b90506020808606600801106131db57600061312f876020885b046001018760400151613655565b905060005b6018601f88166008030181101561318357613155838260208a5b06016137c3565b60f81b84828151811061316457fe5b60200101906001600160f81b031916908160001a905350600101613134565b506018601f8716600803015b60088110156131d4576131a682602089840161309e565b60f81b8482815181106131b557fe5b60200101906001600160f81b031916908160001a90535060010161318f565b5050613224565b60005b6008811015613222576131f4828260208961314e565b60f81b83828151811061320357fe5b60200101906001600160f81b031916908160001a9053506001016131de565b505b6127ef826137d0565b604080516020808252818301909252600091606091906020820181803683370190505090506000613260866020876130fd565b905060208086066020011061331957600061327d87602088613121565b905060005b601f87166020038110156132cb5761329d838260208a61314e565b60f81b8482815181106132ac57fe5b60200101906001600160f81b031916908160001a905350600101613282565b50601f86166008035b60208110156131d4576132eb82602089840161309e565b60f81b8482815181106132fa57fe5b60200101906001600160f81b031916908160001a9053506001016132d4565b60005b602081101561322257613332828260208961314e565b60f81b83828151811061334157fe5b60200101906001600160f81b031916908160001a90535060010161331c565b6080015160ff161590565b60008061337a8660208761308f565b9050600061338c826020880687613806565b905060006133a888602089048488600001518960200151613845565b98975050505050505050565b600060606133c1846138eb565b905060006133d1876020886130fd565b9050602080870660080111156134b35760005b6018601f88166008030181101561342c576134228260208984010685846018018151811061340e57fe5b01602001516001600160f81b031916613955565b91506001016133e4565b50613446876020885b048387600001518860200151613845565b9650600061345688602089613121565b90506018601f8816600803015b600881101561348f576134858260208a84010686846018018151811061340e57fe5b9150600101613463565b506134ab88602089046001018388604001518960600151613845565b9750506134f2565b60005b60088110156134e2576134d8828260208a060185846018018151811061340e57fe5b91506001016134b6565b506134ef87602088613435565b96505b509495945050505050565b6000606061350a846138eb565b9050600061351a876020886130fd565b9050602080870660200111156135ac5760005b601f871660200381101561355c57613552828260208a5b060185848151811061340e57fe5b915060010161352d565b5061356987602088613435565b9650600061357988602089613121565b9050601f87166020035b602081101561348f576135a28260208a84010686848151811061340e57fe5b9150600101613583565b60005b60208110156134e2576135c5828260208a613544565b91506001016135af565b6000806135de8560208661308f565b9050601f84165b6020811015613610576135f882826137c3565b1561360857600092505050611864565b6001016135e5565b5061186085602086048560000151613971565b60008061363b858561363486613aed565b6001613b18565b50855190915061186090613650607b84613bec565b613bec565b60008151600014156136be5761366b600061251d565b84146136b6576040805162461bcd60e51b815260206004820152601560248201527432bc3832b1ba32b21032b6b83a3c90313ab33332b960591b604482015290519081900360640190fd5b506000611864565b60006136dd836000815181106136d057fe5b602002602001015161251d565b905060015b835181101561374757846001166001141561371b5761371484828151811061370657fe5b602002602001015183613bec565b915061373b565b6137388285838151811061372b57fe5b6020026020010151613bec565b91505b600194851c94016136e2565b50848114613794576040805162461bcd60e51b8152602060048201526015602482015274195e1c1958dd19590818dbdc9c9958dd081c9bdbdd605a1b604482015290519081900360640190fd5b83156137a4575060009050611864565b826000815181106137b157fe5b60200260200101519150509392505050565b601f036008021c60ff1690565b600080805b8351811015611cb657600882901b91508381815181106137f157fe5b016020015160f81c91909117906001016137d5565b60006060613813856138eb565b90508260f81b81858151811061382557fe5b60200101906001600160f81b031916908160001a905350611860816137d0565b6000815160031461389d576040805162461bcd60e51b815260206004820152601760248201527f4241445f4e4f524d414c495a4154494f4e5f50524f4f46000000000000000000604482015290519081900360640190fd5b6127ef86868686866000815181106138b157fe5b602002602001015160001c876001815181106138c957fe5b6020026020010151886002815181106138de57fe5b6020026020010151613c18565b6040805160208082528183019092526060918391839160208201818036833701905050905060005b6020811015612ed4578260f81b8282601f038151811061392f57fe5b60200101906001600160f81b031916908160001a90535060089290921c91600101613913565b60006060613962856138eb565b90508281858151811061382557fe5b60008151600014156139da57613987600061251d565b84146139d2576040805162461bcd60e51b815260206004820152601560248201527432bc3832b1ba32b21032b6b83a3c90313ab33332b960591b604482015290519081900360640190fd5b506001611864565b60006139ec836000815181106136d057fe5b9050600160606139fa613ea9565b905060015b8551811015613a8e578660011660011415613a3857613a31868281518110613a2357fe5b602002602001015185613bec565b9350613a82565b613a488487838151811061372b57fe5b9350828015613a7f5750816001820381518110613a6157fe5b6020026020010151868281518110613a7557fe5b6020026020010151145b92505b600196871c96016139ff565b50868314613adb576040805162461bcd60e51b8152602060048201526015602482015274195e1c1958dd19590818dbdc9c9958dd081c9bdbdd605a1b604482015290519081900360640190fd5b8515611d8d5760019350505050611864565b600060018211613aff575060016112e0565b613b0e60026001840104613aed565b60020290506112e0565b60008060208411613b6c5785518510613b4057613b35600061251d565b600191509150613be3565b6000613b54613b4f8888613f4a565b61251d565b905080613b61600061251d565b909350149050613be3565b600080613b85886002880489016002895b046000613b18565b91509150808015613b935750845b15613bb257613ba788886002890488613b18565b935093505050613be3565b600080613bc28a8a60028b613b7d565b91509150613bd08285613bec565b818015613bda5750835b95509550505050505b94509492505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080613c248761251d565b9050613c31898988613655565b506060613c3c613ea9565b905060018751036001901b8910613d005787613c5c578992505050613e9e565b6000613c678a613fa1565b88519091505b60018203811015613c9557613c8b8c84600184038151811061372b57fe5b9b50600101613c6d565b5060015b60018203811015613ceb578a60011660011415613cc957613cc2836001830381518110613a2357fe5b9350613cdf565b613cdc8484600184038151811061372b57fe5b93505b60019a8b1c9a01613c99565b50613cf68b84613bec565b9350505050613e9e565b60015b8751811015613d805760008a600116600114613d1f5783613d34565b888281518110613d2b57fe5b60200260200101515b905060008b600116600114613d5c57898381518110613d4f57fe5b6020026020010151613d5e565b845b9050613d6a8282613bec565b60019c8d1c9c909550929092019150613d039050565b508715613d8f57509050613e9e565b808681518110613d9b57fe5b602002602001015184141580613daf575085155b613e00576040805162461bcd60e51b815260206004820152601c60248201527f726967687420737562747265652063616e6e6f74206265207a65726f00000000604482015290519081900360640190fd5b60008615613e1757613e128686613bec565b613e19565b855b90508615613e2957600019909601955b80875b60018a5103811015613e5257613e488285838151811061372b57fe5b9150600101613e2c565b50838114613e98576040805162461bcd60e51b815260206004820152600e60248201526d0caf0e0cac6e8cac840dac2e8c6d60931b604482015290519081900360640190fd5b50925050505b979650505050505050565b60408051818152610820810182526060918291906020820161080080368337019050509050613ed8600061251d565b81600081518110613ee557fe5b602090810291909101015260015b6040811015613f4457613f25826001830381518110613f0e57fe5b602002602001015183600184038151811061372b57fe5b828281518110613f3157fe5b6020908102919091010152600101613ef3565b50905090565b600080805b6020811015612ed457600882901b91506000818501865111613f72576000613f90565b8582860181518110613f8057fe5b01602001516001600160f81b0319165b60f81c929092179150600101613f4f565b600081613fb0575060016112e0565b613fbd600183901c613fa1565b60010190506112e0565b604051806101e0016040528060006001600160a01b03168152602001613feb6140e0565b8152602001613ff86140e0565b81526000602082018190526040820181905260608201819052608082015260a00161402161414d565b815260200161402e61414d565b81526000602082018190526040820181905260608083018190526080830182905260a083015260c09091015290565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b60405180608001604052806004906020820280368337509192915050565b6040518060c00160405280600081526020016140bd614167565b815260606020820181905260006040830181905290820181905260809091015290565b60408051610120810190915260008152602081016140fc6140a3565b81526020016141096140a3565b81526020016141166140a3565b81526020016141236140a3565b8152600060208201819052604082015260600161413e6140a3565b8152602001600081525090565bfe5b604051806040016040528060008152602001606081525090565b604080516060808201835260008083526020830152918101919091529056fe75736520616e6f7468657220636f6e747261637420746f2068616e646c65206f74686572206f70636f646573a26469706673582212209ae17101c926c9c5afc9fbf7d83991a708fdd5d8dcceaeeaaf9db2dcddb8cb3664736f6c634300060c0033"

// DeployOneStepProof2 deploys a new Ethereum contract, binding an instance of OneStepProof2 to it.
func DeployOneStepProof2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof2, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProof2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProof2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof2{OneStepProof2Caller: OneStepProof2Caller{contract: contract}, OneStepProof2Transactor: OneStepProof2Transactor{contract: contract}, OneStepProof2Filterer: OneStepProof2Filterer{contract: contract}}, nil
}

// OneStepProof2 is an auto generated Go binding around an Ethereum contract.
type OneStepProof2 struct {
	OneStepProof2Caller     // Read-only binding to the contract
	OneStepProof2Transactor // Write-only binding to the contract
	OneStepProof2Filterer   // Log filterer for contract events
}

// OneStepProof2Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProof2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProof2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProof2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProof2Session struct {
	Contract     *OneStepProof2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProof2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProof2CallerSession struct {
	Contract *OneStepProof2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OneStepProof2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProof2TransactorSession struct {
	Contract     *OneStepProof2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OneStepProof2Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProof2Raw struct {
	Contract *OneStepProof2 // Generic contract binding to access the raw methods on
}

// OneStepProof2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProof2CallerRaw struct {
	Contract *OneStepProof2Caller // Generic read-only contract binding to access the raw methods on
}

// OneStepProof2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProof2TransactorRaw struct {
	Contract *OneStepProof2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof2 creates a new instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2(address common.Address, backend bind.ContractBackend) (*OneStepProof2, error) {
	contract, err := bindOneStepProof2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2{OneStepProof2Caller: OneStepProof2Caller{contract: contract}, OneStepProof2Transactor: OneStepProof2Transactor{contract: contract}, OneStepProof2Filterer: OneStepProof2Filterer{contract: contract}}, nil
}

// NewOneStepProof2Caller creates a new read-only instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Caller(address common.Address, caller bind.ContractCaller) (*OneStepProof2Caller, error) {
	contract, err := bindOneStepProof2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Caller{contract: contract}, nil
}

// NewOneStepProof2Transactor creates a new write-only instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Transactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProof2Transactor, error) {
	contract, err := bindOneStepProof2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Transactor{contract: contract}, nil
}

// NewOneStepProof2Filterer creates a new log filterer instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Filterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProof2Filterer, error) {
	contract, err := bindOneStepProof2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Filterer{contract: contract}, nil
}

// bindOneStepProof2 binds a generic wrapper to an already deployed contract.
func bindOneStepProof2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProof2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof2 *OneStepProof2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProof2.Contract.OneStepProof2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof2 *OneStepProof2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof2.Contract.OneStepProof2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof2 *OneStepProof2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof2.Contract.OneStepProof2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof2 *OneStepProof2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProof2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof2 *OneStepProof2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof2 *OneStepProof2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof2.Contract.contract.Transact(opts, method, params...)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x9d16dd04.
//
// Solidity: function executeStep(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields)
func (_OneStepProof2 *OneStepProof2Caller) ExecuteStep(opts *bind.CallOpts, bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	var out []interface{}
	err := _OneStepProof2.contract.Call(opts, &out, "executeStep", bridge, initialMessagesRead, accs, proof, bproof)

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
func (_OneStepProof2 *OneStepProof2Session) ExecuteStep(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _OneStepProof2.Contract.ExecuteStep(&_OneStepProof2.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x9d16dd04.
//
// Solidity: function executeStep(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields)
func (_OneStepProof2 *OneStepProof2CallerSession) ExecuteStep(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _OneStepProof2.Contract.ExecuteStep(&_OneStepProof2.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x2ccebb7a.
//
// Solidity: function executeStepDebug(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_OneStepProof2 *OneStepProof2Caller) ExecuteStepDebug(opts *bind.CallOpts, bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	var out []interface{}
	err := _OneStepProof2.contract.Call(opts, &out, "executeStepDebug", bridge, initialMessagesRead, accs, proof, bproof)

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
func (_OneStepProof2 *OneStepProof2Session) ExecuteStepDebug(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _OneStepProof2.Contract.ExecuteStepDebug(&_OneStepProof2.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x2ccebb7a.
//
// Solidity: function executeStepDebug(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_OneStepProof2 *OneStepProof2CallerSession) ExecuteStepDebug(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _OneStepProof2.Contract.ExecuteStepDebug(&_OneStepProof2.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ParseProof is a free data retrieval call binding the contract method 0x793deea3.
//
// Solidity: function parseProof(bytes proof) pure returns(bytes32[], bytes32[], bytes32[], bytes32[])
func (_OneStepProof2 *OneStepProof2Caller) ParseProof(opts *bind.CallOpts, proof []byte) ([][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	var out []interface{}
	err := _OneStepProof2.contract.Call(opts, &out, "parseProof", proof)

	if err != nil {
		return *new([][32]byte), *new([][32]byte), *new([][32]byte), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	out1 := *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)
	out3 := *abi.ConvertType(out[3], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, out3, err

}

// ParseProof is a free data retrieval call binding the contract method 0x793deea3.
//
// Solidity: function parseProof(bytes proof) pure returns(bytes32[], bytes32[], bytes32[], bytes32[])
func (_OneStepProof2 *OneStepProof2Session) ParseProof(proof []byte) ([][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _OneStepProof2.Contract.ParseProof(&_OneStepProof2.CallOpts, proof)
}

// ParseProof is a free data retrieval call binding the contract method 0x793deea3.
//
// Solidity: function parseProof(bytes proof) pure returns(bytes32[], bytes32[], bytes32[], bytes32[])
func (_OneStepProof2 *OneStepProof2CallerSession) ParseProof(proof []byte) ([][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _OneStepProof2.Contract.ParseProof(&_OneStepProof2.CallOpts, proof)
}
