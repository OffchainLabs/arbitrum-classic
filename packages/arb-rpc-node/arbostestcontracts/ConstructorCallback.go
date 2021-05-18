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

// ConstructorCallbackABI is the input ABI used to generate the binding from.
const ConstructorCallbackABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConstructorCallbackBin is the compiled bytecode used for deploying new contracts.
var ConstructorCallbackBin = "0x608060405234801561001057600080fd5b506040805136815290517f1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d24149181900360200190a13661009c57604080516329e99f0760e01b815261021f6004820152905130916329e99f0791602480830192600092919082900301818387803b15801561008957600080fd5b505af192505050801561009a575060015b505b6082806100aa6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806329e99f0714602d575b600080fd5b604760048036036020811015604157600080fd5b50356049565b005b5056fea26469706673582212204cfeea3f0eabe97fd1613d9c1a99936ca17bebcade1afbace6003b18e1b8847464736f6c634300060c0033"

// DeployConstructorCallback deploys a new Ethereum contract, binding an instance of ConstructorCallback to it.
func DeployConstructorCallback(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConstructorCallback, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstructorCallbackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConstructorCallbackBin), backend)
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

// Test is a paid mutator transaction binding the contract method 0x29e99f07.
//
// Solidity: function test(uint256 data) returns()
func (_ConstructorCallback *ConstructorCallbackTransactor) Test(opts *bind.TransactOpts, data *big.Int) (*types.Transaction, error) {
	return _ConstructorCallback.contract.Transact(opts, "test", data)
}

// Test is a paid mutator transaction binding the contract method 0x29e99f07.
//
// Solidity: function test(uint256 data) returns()
func (_ConstructorCallback *ConstructorCallbackSession) Test(data *big.Int) (*types.Transaction, error) {
	return _ConstructorCallback.Contract.Test(&_ConstructorCallback.TransactOpts, data)
}

// Test is a paid mutator transaction binding the contract method 0x29e99f07.
//
// Solidity: function test(uint256 data) returns()
func (_ConstructorCallback *ConstructorCallbackTransactorSession) Test(data *big.Int) (*types.Transaction, error) {
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
