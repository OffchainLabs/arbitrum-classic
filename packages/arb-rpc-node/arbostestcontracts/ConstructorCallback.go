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

// ConstructorCallbackMetaData contains all meta data concerning the ConstructorCallback contract.
var ConstructorCallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dataLength\",\"type\":\"address\"}],\"name\":\"TestEvent2\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060408190523681527f1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d241490602090a1336001600160a01b03166366e41cb76040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561006a57600080fd5b505af115801561007e573d6000803e3d6000fd5b5050505060c7806100906000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063bb29998e14602d575b600080fd5b605060048036036020811015604157600080fd5b50356001600160a01b03166052565b005b604080516001600160a01b038316815290517fba829c4567200650d8324f5576706bb44be221bc498741a8ddaa9a2739407b7d9181900360200190a15056fea2646970667358221220bd86493538c1d9a00ae7c9a3a334b1dd6ae55f8122422647d819a53637f4e3db64736f6c634300060c0033",
}

// ConstructorCallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstructorCallbackMetaData.ABI instead.
var ConstructorCallbackABI = ConstructorCallbackMetaData.ABI

// ConstructorCallbackBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConstructorCallbackMetaData.Bin instead.
var ConstructorCallbackBin = ConstructorCallbackMetaData.Bin

// DeployConstructorCallback deploys a new Ethereum contract, binding an instance of ConstructorCallback to it.
func DeployConstructorCallback(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConstructorCallback, error) {
	parsed, err := ConstructorCallbackMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConstructorCallbackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConstructorCallback{ConstructorCallbackCaller: ConstructorCallbackCaller{contract: contract}, ConstructorCallbackTransactor: ConstructorCallbackTransactor{contract: contract}, ConstructorCallbackFilterer: ConstructorCallbackFilterer{contract: contract}}, nil
}

// ConstructorCallback is an auto generated Go binding around an Ethereum contract.
type ConstructorCallback struct {
	ConstructorCallbackCaller     // Read-only binding to the contract
	ConstructorCallbackTransactor // Write-only binding to the contract
	ConstructorCallbackFilterer   // Log filterer for contract events
}

// ConstructorCallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstructorCallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstructorCallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstructorCallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstructorCallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstructorCallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstructorCallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstructorCallbackSession struct {
	Contract     *ConstructorCallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConstructorCallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstructorCallbackCallerSession struct {
	Contract *ConstructorCallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ConstructorCallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstructorCallbackTransactorSession struct {
	Contract     *ConstructorCallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ConstructorCallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstructorCallbackRaw struct {
	Contract *ConstructorCallback // Generic contract binding to access the raw methods on
}

// ConstructorCallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstructorCallbackCallerRaw struct {
	Contract *ConstructorCallbackCaller // Generic read-only contract binding to access the raw methods on
}

// ConstructorCallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstructorCallbackTransactorRaw struct {
	Contract *ConstructorCallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstructorCallback creates a new instance of ConstructorCallback, bound to a specific deployed contract.
func NewConstructorCallback(address common.Address, backend bind.ContractBackend) (*ConstructorCallback, error) {
	contract, err := bindConstructorCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallback{ConstructorCallbackCaller: ConstructorCallbackCaller{contract: contract}, ConstructorCallbackTransactor: ConstructorCallbackTransactor{contract: contract}, ConstructorCallbackFilterer: ConstructorCallbackFilterer{contract: contract}}, nil
}

// NewConstructorCallbackCaller creates a new read-only instance of ConstructorCallback, bound to a specific deployed contract.
func NewConstructorCallbackCaller(address common.Address, caller bind.ContractCaller) (*ConstructorCallbackCaller, error) {
	contract, err := bindConstructorCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallbackCaller{contract: contract}, nil
}

// NewConstructorCallbackTransactor creates a new write-only instance of ConstructorCallback, bound to a specific deployed contract.
func NewConstructorCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstructorCallbackTransactor, error) {
	contract, err := bindConstructorCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallbackTransactor{contract: contract}, nil
}

// NewConstructorCallbackFilterer creates a new log filterer instance of ConstructorCallback, bound to a specific deployed contract.
func NewConstructorCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstructorCallbackFilterer, error) {
	contract, err := bindConstructorCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallbackFilterer{contract: contract}, nil
}

// bindConstructorCallback binds a generic wrapper to an already deployed contract.
func bindConstructorCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstructorCallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstructorCallback *ConstructorCallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstructorCallback.Contract.ConstructorCallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstructorCallback *ConstructorCallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstructorCallback.Contract.ConstructorCallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstructorCallback *ConstructorCallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstructorCallback.Contract.ConstructorCallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstructorCallback *ConstructorCallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstructorCallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstructorCallback *ConstructorCallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstructorCallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstructorCallback *ConstructorCallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstructorCallback.Contract.contract.Transact(opts, method, params...)
}

// Test is a paid mutator transaction binding the contract method 0xbb29998e.
//
// Solidity: function test(address data) returns()
func (_ConstructorCallback *ConstructorCallbackTransactor) Test(opts *bind.TransactOpts, data common.Address) (*types.Transaction, error) {
	return _ConstructorCallback.contract.Transact(opts, "test", data)
}

// Test is a paid mutator transaction binding the contract method 0xbb29998e.
//
// Solidity: function test(address data) returns()
func (_ConstructorCallback *ConstructorCallbackSession) Test(data common.Address) (*types.Transaction, error) {
	return _ConstructorCallback.Contract.Test(&_ConstructorCallback.TransactOpts, data)
}

// Test is a paid mutator transaction binding the contract method 0xbb29998e.
//
// Solidity: function test(address data) returns()
func (_ConstructorCallback *ConstructorCallbackTransactorSession) Test(data common.Address) (*types.Transaction, error) {
	return _ConstructorCallback.Contract.Test(&_ConstructorCallback.TransactOpts, data)
}

// ConstructorCallbackTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the ConstructorCallback contract.
type ConstructorCallbackTestEventIterator struct {
	Event *ConstructorCallbackTestEvent // Event containing the contract specifics and raw log

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
func (it *ConstructorCallbackTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstructorCallbackTestEvent)
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
		it.Event = new(ConstructorCallbackTestEvent)
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
func (it *ConstructorCallbackTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstructorCallbackTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstructorCallbackTestEvent represents a TestEvent event raised by the ConstructorCallback contract.
type ConstructorCallbackTestEvent struct {
	DataLength *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 dataLength)
func (_ConstructorCallback *ConstructorCallbackFilterer) FilterTestEvent(opts *bind.FilterOpts) (*ConstructorCallbackTestEventIterator, error) {

	logs, sub, err := _ConstructorCallback.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &ConstructorCallbackTestEventIterator{contract: _ConstructorCallback.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 dataLength)
func (_ConstructorCallback *ConstructorCallbackFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *ConstructorCallbackTestEvent) (event.Subscription, error) {

	logs, sub, err := _ConstructorCallback.contract.WatchLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstructorCallbackTestEvent)
				if err := _ConstructorCallback.contract.UnpackLog(event, "TestEvent", log); err != nil {
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

// ParseTestEvent is a log parse operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 dataLength)
func (_ConstructorCallback *ConstructorCallbackFilterer) ParseTestEvent(log types.Log) (*ConstructorCallbackTestEvent, error) {
	event := new(ConstructorCallbackTestEvent)
	if err := _ConstructorCallback.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstructorCallbackTestEvent2Iterator is returned from FilterTestEvent2 and is used to iterate over the raw logs and unpacked data for TestEvent2 events raised by the ConstructorCallback contract.
type ConstructorCallbackTestEvent2Iterator struct {
	Event *ConstructorCallbackTestEvent2 // Event containing the contract specifics and raw log

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
func (it *ConstructorCallbackTestEvent2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstructorCallbackTestEvent2)
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
		it.Event = new(ConstructorCallbackTestEvent2)
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
func (it *ConstructorCallbackTestEvent2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstructorCallbackTestEvent2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstructorCallbackTestEvent2 represents a TestEvent2 event raised by the ConstructorCallback contract.
type ConstructorCallbackTestEvent2 struct {
	DataLength common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTestEvent2 is a free log retrieval operation binding the contract event 0xba829c4567200650d8324f5576706bb44be221bc498741a8ddaa9a2739407b7d.
//
// Solidity: event TestEvent2(address dataLength)
func (_ConstructorCallback *ConstructorCallbackFilterer) FilterTestEvent2(opts *bind.FilterOpts) (*ConstructorCallbackTestEvent2Iterator, error) {

	logs, sub, err := _ConstructorCallback.contract.FilterLogs(opts, "TestEvent2")
	if err != nil {
		return nil, err
	}
	return &ConstructorCallbackTestEvent2Iterator{contract: _ConstructorCallback.contract, event: "TestEvent2", logs: logs, sub: sub}, nil
}

// WatchTestEvent2 is a free log subscription operation binding the contract event 0xba829c4567200650d8324f5576706bb44be221bc498741a8ddaa9a2739407b7d.
//
// Solidity: event TestEvent2(address dataLength)
func (_ConstructorCallback *ConstructorCallbackFilterer) WatchTestEvent2(opts *bind.WatchOpts, sink chan<- *ConstructorCallbackTestEvent2) (event.Subscription, error) {

	logs, sub, err := _ConstructorCallback.contract.WatchLogs(opts, "TestEvent2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstructorCallbackTestEvent2)
				if err := _ConstructorCallback.contract.UnpackLog(event, "TestEvent2", log); err != nil {
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

// ParseTestEvent2 is a log parse operation binding the contract event 0xba829c4567200650d8324f5576706bb44be221bc498741a8ddaa9a2739407b7d.
//
// Solidity: event TestEvent2(address dataLength)
func (_ConstructorCallback *ConstructorCallbackFilterer) ParseTestEvent2(log types.Log) (*ConstructorCallbackTestEvent2, error) {
	event := new(ConstructorCallbackTestEvent2)
	if err := _ConstructorCallback.contract.UnpackLog(event, "TestEvent2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstructorCallback2MetaData contains all meta data concerning the ConstructorCallback2 contract.
var ConstructorCallback2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"TestEvent3\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061037c806100206000396000f3fe6080604052600436106100295760003560e01c806366e41cb71461002e578063f8a8fd6d14610038575b600080fd5b610036610040565b005b6100366101b6565b6040805133602480830182905283518084039091018152604490920183526020820180516001600160e01b0316635d94ccc760e11b1781529251825160009460609492918291908083835b602083106100aa5780518252601f19909201916020918201910161008b565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461010c576040519150601f19603f3d011682016040523d82523d6000602084013e610111565b606091505b50915091508115157fe7713ed83c9f3ef742bc9aec2c297f6bc4c7be68042d4aa69be6ba74848d1882826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610178578181015183820152602001610160565b50505050905090810190601f1680156101a55780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b6040516101c2906101e2565b604051809103906000f0801580156101de573d6000803e3d6000fd5b5050565b610157806101f08339019056fe608060408190523681527f1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d241490602090a1336001600160a01b03166366e41cb76040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561006a57600080fd5b505af115801561007e573d6000803e3d6000fd5b5050505060c7806100906000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063bb29998e14602d575b600080fd5b605060048036036020811015604157600080fd5b50356001600160a01b03166052565b005b604080516001600160a01b038316815290517fba829c4567200650d8324f5576706bb44be221bc498741a8ddaa9a2739407b7d9181900360200190a15056fea2646970667358221220bd86493538c1d9a00ae7c9a3a334b1dd6ae55f8122422647d819a53637f4e3db64736f6c634300060c0033a26469706673582212204f85473267682fe0de9ce65ceec2d8ed6b6cbeba40b9baa15b090627566612cd64736f6c634300060c0033",
}

// ConstructorCallback2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstructorCallback2MetaData.ABI instead.
var ConstructorCallback2ABI = ConstructorCallback2MetaData.ABI

// ConstructorCallback2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConstructorCallback2MetaData.Bin instead.
var ConstructorCallback2Bin = ConstructorCallback2MetaData.Bin

// DeployConstructorCallback2 deploys a new Ethereum contract, binding an instance of ConstructorCallback2 to it.
func DeployConstructorCallback2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConstructorCallback2, error) {
	parsed, err := ConstructorCallback2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConstructorCallback2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConstructorCallback2{ConstructorCallback2Caller: ConstructorCallback2Caller{contract: contract}, ConstructorCallback2Transactor: ConstructorCallback2Transactor{contract: contract}, ConstructorCallback2Filterer: ConstructorCallback2Filterer{contract: contract}}, nil
}

// ConstructorCallback2 is an auto generated Go binding around an Ethereum contract.
type ConstructorCallback2 struct {
	ConstructorCallback2Caller     // Read-only binding to the contract
	ConstructorCallback2Transactor // Write-only binding to the contract
	ConstructorCallback2Filterer   // Log filterer for contract events
}

// ConstructorCallback2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ConstructorCallback2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstructorCallback2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstructorCallback2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstructorCallback2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstructorCallback2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstructorCallback2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstructorCallback2Session struct {
	Contract     *ConstructorCallback2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ConstructorCallback2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstructorCallback2CallerSession struct {
	Contract *ConstructorCallback2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ConstructorCallback2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstructorCallback2TransactorSession struct {
	Contract     *ConstructorCallback2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ConstructorCallback2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ConstructorCallback2Raw struct {
	Contract *ConstructorCallback2 // Generic contract binding to access the raw methods on
}

// ConstructorCallback2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstructorCallback2CallerRaw struct {
	Contract *ConstructorCallback2Caller // Generic read-only contract binding to access the raw methods on
}

// ConstructorCallback2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstructorCallback2TransactorRaw struct {
	Contract *ConstructorCallback2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewConstructorCallback2 creates a new instance of ConstructorCallback2, bound to a specific deployed contract.
func NewConstructorCallback2(address common.Address, backend bind.ContractBackend) (*ConstructorCallback2, error) {
	contract, err := bindConstructorCallback2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallback2{ConstructorCallback2Caller: ConstructorCallback2Caller{contract: contract}, ConstructorCallback2Transactor: ConstructorCallback2Transactor{contract: contract}, ConstructorCallback2Filterer: ConstructorCallback2Filterer{contract: contract}}, nil
}

// NewConstructorCallback2Caller creates a new read-only instance of ConstructorCallback2, bound to a specific deployed contract.
func NewConstructorCallback2Caller(address common.Address, caller bind.ContractCaller) (*ConstructorCallback2Caller, error) {
	contract, err := bindConstructorCallback2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallback2Caller{contract: contract}, nil
}

// NewConstructorCallback2Transactor creates a new write-only instance of ConstructorCallback2, bound to a specific deployed contract.
func NewConstructorCallback2Transactor(address common.Address, transactor bind.ContractTransactor) (*ConstructorCallback2Transactor, error) {
	contract, err := bindConstructorCallback2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallback2Transactor{contract: contract}, nil
}

// NewConstructorCallback2Filterer creates a new log filterer instance of ConstructorCallback2, bound to a specific deployed contract.
func NewConstructorCallback2Filterer(address common.Address, filterer bind.ContractFilterer) (*ConstructorCallback2Filterer, error) {
	contract, err := bindConstructorCallback2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallback2Filterer{contract: contract}, nil
}

// bindConstructorCallback2 binds a generic wrapper to an already deployed contract.
func bindConstructorCallback2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstructorCallback2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstructorCallback2 *ConstructorCallback2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstructorCallback2.Contract.ConstructorCallback2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstructorCallback2 *ConstructorCallback2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstructorCallback2.Contract.ConstructorCallback2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstructorCallback2 *ConstructorCallback2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstructorCallback2.Contract.ConstructorCallback2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstructorCallback2 *ConstructorCallback2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstructorCallback2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstructorCallback2 *ConstructorCallback2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstructorCallback2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstructorCallback2 *ConstructorCallback2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstructorCallback2.Contract.contract.Transact(opts, method, params...)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() payable returns()
func (_ConstructorCallback2 *ConstructorCallback2Transactor) Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstructorCallback2.contract.Transact(opts, "test")
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() payable returns()
func (_ConstructorCallback2 *ConstructorCallback2Session) Test() (*types.Transaction, error) {
	return _ConstructorCallback2.Contract.Test(&_ConstructorCallback2.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() payable returns()
func (_ConstructorCallback2 *ConstructorCallback2TransactorSession) Test() (*types.Transaction, error) {
	return _ConstructorCallback2.Contract.Test(&_ConstructorCallback2.TransactOpts)
}

// Test2 is a paid mutator transaction binding the contract method 0x66e41cb7.
//
// Solidity: function test2() payable returns()
func (_ConstructorCallback2 *ConstructorCallback2Transactor) Test2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstructorCallback2.contract.Transact(opts, "test2")
}

// Test2 is a paid mutator transaction binding the contract method 0x66e41cb7.
//
// Solidity: function test2() payable returns()
func (_ConstructorCallback2 *ConstructorCallback2Session) Test2() (*types.Transaction, error) {
	return _ConstructorCallback2.Contract.Test2(&_ConstructorCallback2.TransactOpts)
}

// Test2 is a paid mutator transaction binding the contract method 0x66e41cb7.
//
// Solidity: function test2() payable returns()
func (_ConstructorCallback2 *ConstructorCallback2TransactorSession) Test2() (*types.Transaction, error) {
	return _ConstructorCallback2.Contract.Test2(&_ConstructorCallback2.TransactOpts)
}

// ConstructorCallback2TestEvent3Iterator is returned from FilterTestEvent3 and is used to iterate over the raw logs and unpacked data for TestEvent3 events raised by the ConstructorCallback2 contract.
type ConstructorCallback2TestEvent3Iterator struct {
	Event *ConstructorCallback2TestEvent3 // Event containing the contract specifics and raw log

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
func (it *ConstructorCallback2TestEvent3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstructorCallback2TestEvent3)
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
		it.Event = new(ConstructorCallback2TestEvent3)
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
func (it *ConstructorCallback2TestEvent3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstructorCallback2TestEvent3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstructorCallback2TestEvent3 represents a TestEvent3 event raised by the ConstructorCallback2 contract.
type ConstructorCallback2TestEvent3 struct {
	Success    bool
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTestEvent3 is a free log retrieval operation binding the contract event 0xe7713ed83c9f3ef742bc9aec2c297f6bc4c7be68042d4aa69be6ba74848d1882.
//
// Solidity: event TestEvent3(bool indexed success, bytes returnData)
func (_ConstructorCallback2 *ConstructorCallback2Filterer) FilterTestEvent3(opts *bind.FilterOpts, success []bool) (*ConstructorCallback2TestEvent3Iterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ConstructorCallback2.contract.FilterLogs(opts, "TestEvent3", successRule)
	if err != nil {
		return nil, err
	}
	return &ConstructorCallback2TestEvent3Iterator{contract: _ConstructorCallback2.contract, event: "TestEvent3", logs: logs, sub: sub}, nil
}

// WatchTestEvent3 is a free log subscription operation binding the contract event 0xe7713ed83c9f3ef742bc9aec2c297f6bc4c7be68042d4aa69be6ba74848d1882.
//
// Solidity: event TestEvent3(bool indexed success, bytes returnData)
func (_ConstructorCallback2 *ConstructorCallback2Filterer) WatchTestEvent3(opts *bind.WatchOpts, sink chan<- *ConstructorCallback2TestEvent3, success []bool) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ConstructorCallback2.contract.WatchLogs(opts, "TestEvent3", successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstructorCallback2TestEvent3)
				if err := _ConstructorCallback2.contract.UnpackLog(event, "TestEvent3", log); err != nil {
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

// ParseTestEvent3 is a log parse operation binding the contract event 0xe7713ed83c9f3ef742bc9aec2c297f6bc4c7be68042d4aa69be6ba74848d1882.
//
// Solidity: event TestEvent3(bool indexed success, bytes returnData)
func (_ConstructorCallback2 *ConstructorCallback2Filterer) ParseTestEvent3(log types.Log) (*ConstructorCallback2TestEvent3, error) {
	event := new(ConstructorCallback2TestEvent3)
	if err := _ConstructorCallback2.contract.UnpackLog(event, "TestEvent3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
