// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executionchallenge

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

// BisectionChallengeABI is the input ABI used to generate the binding from.
const BisectionChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"segmentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Continued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_segmentToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"chooseSegment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BisectionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var BisectionChallengeFuncSigs = map[string]string{
	"79a9ad85": "chooseSegment(uint256,bytes,bytes32,bytes32)",
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
	"ced5c1bf": "timeoutChallenge()",
}

// BisectionChallenge is an auto generated Go binding around an Ethereum contract.
type BisectionChallenge struct {
	BisectionChallengeCaller     // Read-only binding to the contract
	BisectionChallengeTransactor // Write-only binding to the contract
	BisectionChallengeFilterer   // Log filterer for contract events
}

// BisectionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BisectionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BisectionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BisectionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BisectionChallengeSession struct {
	Contract     *BisectionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BisectionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BisectionChallengeCallerSession struct {
	Contract *BisectionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BisectionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BisectionChallengeTransactorSession struct {
	Contract     *BisectionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BisectionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BisectionChallengeRaw struct {
	Contract *BisectionChallenge // Generic contract binding to access the raw methods on
}

// BisectionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BisectionChallengeCallerRaw struct {
	Contract *BisectionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// BisectionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BisectionChallengeTransactorRaw struct {
	Contract *BisectionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBisectionChallenge creates a new instance of BisectionChallenge, bound to a specific deployed contract.
func NewBisectionChallenge(address common.Address, backend bind.ContractBackend) (*BisectionChallenge, error) {
	contract, err := bindBisectionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BisectionChallenge{BisectionChallengeCaller: BisectionChallengeCaller{contract: contract}, BisectionChallengeTransactor: BisectionChallengeTransactor{contract: contract}, BisectionChallengeFilterer: BisectionChallengeFilterer{contract: contract}}, nil
}

// NewBisectionChallengeCaller creates a new read-only instance of BisectionChallenge, bound to a specific deployed contract.
func NewBisectionChallengeCaller(address common.Address, caller bind.ContractCaller) (*BisectionChallengeCaller, error) {
	contract, err := bindBisectionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeCaller{contract: contract}, nil
}

// NewBisectionChallengeTransactor creates a new write-only instance of BisectionChallenge, bound to a specific deployed contract.
func NewBisectionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*BisectionChallengeTransactor, error) {
	contract, err := bindBisectionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeTransactor{contract: contract}, nil
}

// NewBisectionChallengeFilterer creates a new log filterer instance of BisectionChallenge, bound to a specific deployed contract.
func NewBisectionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*BisectionChallengeFilterer, error) {
	contract, err := bindBisectionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeFilterer{contract: contract}, nil
}

// bindBisectionChallenge binds a generic wrapper to an already deployed contract.
func bindBisectionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BisectionChallenge *BisectionChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BisectionChallenge.Contract.BisectionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BisectionChallenge *BisectionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.BisectionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BisectionChallenge *BisectionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.BisectionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BisectionChallenge *BisectionChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BisectionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BisectionChallenge *BisectionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BisectionChallenge *BisectionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.contract.Transact(opts, method, params...)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_BisectionChallenge *BisectionChallengeTransactor) ChooseSegment(opts *bind.TransactOpts, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.contract.Transact(opts, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_BisectionChallenge *BisectionChallengeSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.ChooseSegment(&_BisectionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_BisectionChallenge *BisectionChallengeTransactorSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.ChooseSegment(&_BisectionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_BisectionChallenge *BisectionChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.contract.Transact(opts, "initializeBisection", _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_BisectionChallenge *BisectionChallengeSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.InitializeBisection(&_BisectionChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_BisectionChallenge *BisectionChallengeTransactorSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _BisectionChallenge.Contract.InitializeBisection(&_BisectionChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_BisectionChallenge *BisectionChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BisectionChallenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_BisectionChallenge *BisectionChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _BisectionChallenge.Contract.TimeoutChallenge(&_BisectionChallenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_BisectionChallenge *BisectionChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _BisectionChallenge.Contract.TimeoutChallenge(&_BisectionChallenge.TransactOpts)
}

// BisectionChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the BisectionChallenge contract.
type BisectionChallengeAsserterTimedOutIterator struct {
	Event *BisectionChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *BisectionChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionChallengeAsserterTimedOut)
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
		it.Event = new(BisectionChallengeAsserterTimedOut)
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
func (it *BisectionChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the BisectionChallenge contract.
type BisectionChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*BisectionChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _BisectionChallenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeAsserterTimedOutIterator{contract: _BisectionChallenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *BisectionChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _BisectionChallenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionChallengeAsserterTimedOut)
				if err := _BisectionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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

// ParseAsserterTimedOut is a log parse operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*BisectionChallengeAsserterTimedOut, error) {
	event := new(BisectionChallengeAsserterTimedOut)
	if err := _BisectionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the BisectionChallenge contract.
type BisectionChallengeChallengerTimedOutIterator struct {
	Event *BisectionChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *BisectionChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionChallengeChallengerTimedOut)
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
		it.Event = new(BisectionChallengeChallengerTimedOut)
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
func (it *BisectionChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the BisectionChallenge contract.
type BisectionChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*BisectionChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _BisectionChallenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeChallengerTimedOutIterator{contract: _BisectionChallenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *BisectionChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _BisectionChallenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionChallengeChallengerTimedOut)
				if err := _BisectionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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

// ParseChallengerTimedOut is a log parse operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_BisectionChallenge *BisectionChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*BisectionChallengeChallengerTimedOut, error) {
	event := new(BisectionChallengeChallengerTimedOut)
	if err := _BisectionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionChallengeContinuedIterator is returned from FilterContinued and is used to iterate over the raw logs and unpacked data for Continued events raised by the BisectionChallenge contract.
type BisectionChallengeContinuedIterator struct {
	Event *BisectionChallengeContinued // Event containing the contract specifics and raw log

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
func (it *BisectionChallengeContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionChallengeContinued)
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
		it.Event = new(BisectionChallengeContinued)
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
func (it *BisectionChallengeContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionChallengeContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionChallengeContinued represents a Continued event raised by the BisectionChallenge contract.
type BisectionChallengeContinued struct {
	SegmentIndex  *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContinued is a free log retrieval operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) FilterContinued(opts *bind.FilterOpts) (*BisectionChallengeContinuedIterator, error) {

	logs, sub, err := _BisectionChallenge.contract.FilterLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeContinuedIterator{contract: _BisectionChallenge.contract, event: "Continued", logs: logs, sub: sub}, nil
}

// WatchContinued is a free log subscription operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) WatchContinued(opts *bind.WatchOpts, sink chan<- *BisectionChallengeContinued) (event.Subscription, error) {

	logs, sub, err := _BisectionChallenge.contract.WatchLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionChallengeContinued)
				if err := _BisectionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
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

// ParseContinued is a log parse operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) ParseContinued(log types.Log) (*BisectionChallengeContinued, error) {
	event := new(BisectionChallengeContinued)
	if err := _BisectionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the BisectionChallenge contract.
type BisectionChallengeInitiatedChallengeIterator struct {
	Event *BisectionChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *BisectionChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionChallengeInitiatedChallenge)
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
		it.Event = new(BisectionChallengeInitiatedChallenge)
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
func (it *BisectionChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the BisectionChallenge contract.
type BisectionChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*BisectionChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _BisectionChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &BisectionChallengeInitiatedChallengeIterator{contract: _BisectionChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *BisectionChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _BisectionChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionChallengeInitiatedChallenge)
				if err := _BisectionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_BisectionChallenge *BisectionChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*BisectionChallengeInitiatedChallenge, error) {
	event := new(BisectionChallengeInitiatedChallenge)
	if err := _BisectionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820557675814df39add0fa761bb6fc1ca2754618e3660ac97035c79af694b750a9264736f6c634300050d0032"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// ChallengeABI is the input ABI used to generate the binding from.
const ChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFuncSigs = map[string]string{
	"ced5c1bf": "timeoutChallenge()",
}

// Challenge is an auto generated Go binding around an Ethereum contract.
type Challenge struct {
	ChallengeCaller     // Read-only binding to the contract
	ChallengeTransactor // Write-only binding to the contract
	ChallengeFilterer   // Log filterer for contract events
}

// ChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeSession struct {
	Contract     *Challenge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeCallerSession struct {
	Contract *ChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTransactorSession struct {
	Contract     *ChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeRaw struct {
	Contract *Challenge // Generic contract binding to access the raw methods on
}

// ChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeCallerRaw struct {
	Contract *ChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTransactorRaw struct {
	Contract *ChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallenge creates a new instance of Challenge, bound to a specific deployed contract.
func NewChallenge(address common.Address, backend bind.ContractBackend) (*Challenge, error) {
	contract, err := bindChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// NewChallengeCaller creates a new read-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeCaller, error) {
	contract, err := bindChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeCaller{contract: contract}, nil
}

// NewChallengeTransactor creates a new write-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTransactor, error) {
	contract, err := bindChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTransactor{contract: contract}, nil
}

// NewChallengeFilterer creates a new log filterer instance of Challenge, bound to a specific deployed contract.
func NewChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFilterer, error) {
	contract, err := bindChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFilterer{contract: contract}, nil
}

// bindChallenge binds a generic wrapper to an already deployed contract.
func bindChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.ChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transact(opts, method, params...)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_Challenge *ChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_Challenge *ChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _Challenge.Contract.TimeoutChallenge(&_Challenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_Challenge *ChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _Challenge.Contract.TimeoutChallenge(&_Challenge.TransactOpts)
}

// ChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the Challenge contract.
type ChallengeAsserterTimedOutIterator struct {
	Event *ChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *ChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeAsserterTimedOut)
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
		it.Event = new(ChallengeAsserterTimedOut)
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
func (it *ChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the Challenge contract.
type ChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_Challenge *ChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*ChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &ChallengeAsserterTimedOutIterator{contract: _Challenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_Challenge *ChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *ChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeAsserterTimedOut)
				if err := _Challenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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

// ParseAsserterTimedOut is a log parse operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_Challenge *ChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*ChallengeAsserterTimedOut, error) {
	event := new(ChallengeAsserterTimedOut)
	if err := _Challenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the Challenge contract.
type ChallengeChallengerTimedOutIterator struct {
	Event *ChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *ChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeChallengerTimedOut)
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
		it.Event = new(ChallengeChallengerTimedOut)
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
func (it *ChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the Challenge contract.
type ChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_Challenge *ChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*ChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &ChallengeChallengerTimedOutIterator{contract: _Challenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_Challenge *ChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *ChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeChallengerTimedOut)
				if err := _Challenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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

// ParseChallengerTimedOut is a log parse operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_Challenge *ChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*ChallengeChallengerTimedOut, error) {
	event := new(ChallengeChallengerTimedOut)
	if err := _Challenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the Challenge contract.
type ChallengeInitiatedChallengeIterator struct {
	Event *ChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeInitiatedChallenge)
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
		it.Event = new(ChallengeInitiatedChallenge)
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
func (it *ChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the Challenge contract.
type ChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_Challenge *ChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ChallengeInitiatedChallengeIterator{contract: _Challenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_Challenge *ChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeInitiatedChallenge)
				if err := _Challenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_Challenge *ChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ChallengeInitiatedChallenge, error) {
	event := new(ChallengeInitiatedChallenge)
	if err := _Challenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeTypeABI is the input ABI used to generate the binding from.
const ChallengeTypeABI = "[]"

// ChallengeTypeBin is the compiled bytecode used for deploying new contracts.
var ChallengeTypeBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a7231582022b77c85b5fc06a9af3a4a0d6d6ee856ab2c07ba44bd5245ee6e95680a5c1c8c64736f6c634300050d0032"

// DeployChallengeType deploys a new Ethereum contract, binding an instance of ChallengeType to it.
func DeployChallengeType(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeType, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTypeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeTypeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeType{ChallengeTypeCaller: ChallengeTypeCaller{contract: contract}, ChallengeTypeTransactor: ChallengeTypeTransactor{contract: contract}, ChallengeTypeFilterer: ChallengeTypeFilterer{contract: contract}}, nil
}

// ChallengeType is an auto generated Go binding around an Ethereum contract.
type ChallengeType struct {
	ChallengeTypeCaller     // Read-only binding to the contract
	ChallengeTypeTransactor // Write-only binding to the contract
	ChallengeTypeFilterer   // Log filterer for contract events
}

// ChallengeTypeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeTypeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTypeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTypeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTypeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeTypeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTypeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeTypeSession struct {
	Contract     *ChallengeType    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeTypeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeTypeCallerSession struct {
	Contract *ChallengeTypeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ChallengeTypeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTypeTransactorSession struct {
	Contract     *ChallengeTypeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ChallengeTypeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeTypeRaw struct {
	Contract *ChallengeType // Generic contract binding to access the raw methods on
}

// ChallengeTypeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeTypeCallerRaw struct {
	Contract *ChallengeTypeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTypeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTypeTransactorRaw struct {
	Contract *ChallengeTypeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeType creates a new instance of ChallengeType, bound to a specific deployed contract.
func NewChallengeType(address common.Address, backend bind.ContractBackend) (*ChallengeType, error) {
	contract, err := bindChallengeType(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeType{ChallengeTypeCaller: ChallengeTypeCaller{contract: contract}, ChallengeTypeTransactor: ChallengeTypeTransactor{contract: contract}, ChallengeTypeFilterer: ChallengeTypeFilterer{contract: contract}}, nil
}

// NewChallengeTypeCaller creates a new read-only instance of ChallengeType, bound to a specific deployed contract.
func NewChallengeTypeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeTypeCaller, error) {
	contract, err := bindChallengeType(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTypeCaller{contract: contract}, nil
}

// NewChallengeTypeTransactor creates a new write-only instance of ChallengeType, bound to a specific deployed contract.
func NewChallengeTypeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTypeTransactor, error) {
	contract, err := bindChallengeType(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTypeTransactor{contract: contract}, nil
}

// NewChallengeTypeFilterer creates a new log filterer instance of ChallengeType, bound to a specific deployed contract.
func NewChallengeTypeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeTypeFilterer, error) {
	contract, err := bindChallengeType(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeTypeFilterer{contract: contract}, nil
}

// bindChallengeType binds a generic wrapper to an already deployed contract.
func bindChallengeType(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTypeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeType *ChallengeTypeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeType.Contract.ChallengeTypeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeType *ChallengeTypeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeType.Contract.ChallengeTypeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeType *ChallengeTypeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeType.Contract.ChallengeTypeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeType *ChallengeTypeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeType.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeType *ChallengeTypeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeType.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeType *ChallengeTypeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeType.Contract.contract.Transact(opts, method, params...)
}

// ChallengeUtilsABI is the input ABI used to generate the binding from.
const ChallengeUtilsABI = "[]"

// ChallengeUtilsBin is the compiled bytecode used for deploying new contracts.
var ChallengeUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820607aae7be6ed60ed7096d0341561f4183aaa07bcafa77d66f753e60c7600bc2264736f6c634300050d0032"

// DeployChallengeUtils deploys a new Ethereum contract, binding an instance of ChallengeUtils to it.
func DeployChallengeUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeUtils{ChallengeUtilsCaller: ChallengeUtilsCaller{contract: contract}, ChallengeUtilsTransactor: ChallengeUtilsTransactor{contract: contract}, ChallengeUtilsFilterer: ChallengeUtilsFilterer{contract: contract}}, nil
}

// ChallengeUtils is an auto generated Go binding around an Ethereum contract.
type ChallengeUtils struct {
	ChallengeUtilsCaller     // Read-only binding to the contract
	ChallengeUtilsTransactor // Write-only binding to the contract
	ChallengeUtilsFilterer   // Log filterer for contract events
}

// ChallengeUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeUtilsSession struct {
	Contract     *ChallengeUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeUtilsCallerSession struct {
	Contract *ChallengeUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ChallengeUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeUtilsTransactorSession struct {
	Contract     *ChallengeUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChallengeUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeUtilsRaw struct {
	Contract *ChallengeUtils // Generic contract binding to access the raw methods on
}

// ChallengeUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeUtilsCallerRaw struct {
	Contract *ChallengeUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeUtilsTransactorRaw struct {
	Contract *ChallengeUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeUtils creates a new instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtils(address common.Address, backend bind.ContractBackend) (*ChallengeUtils, error) {
	contract, err := bindChallengeUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtils{ChallengeUtilsCaller: ChallengeUtilsCaller{contract: contract}, ChallengeUtilsTransactor: ChallengeUtilsTransactor{contract: contract}, ChallengeUtilsFilterer: ChallengeUtilsFilterer{contract: contract}}, nil
}

// NewChallengeUtilsCaller creates a new read-only instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsCaller(address common.Address, caller bind.ContractCaller) (*ChallengeUtilsCaller, error) {
	contract, err := bindChallengeUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsCaller{contract: contract}, nil
}

// NewChallengeUtilsTransactor creates a new write-only instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeUtilsTransactor, error) {
	contract, err := bindChallengeUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsTransactor{contract: contract}, nil
}

// NewChallengeUtilsFilterer creates a new log filterer instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeUtilsFilterer, error) {
	contract, err := bindChallengeUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsFilterer{contract: contract}, nil
}

// bindChallengeUtils binds a generic wrapper to an already deployed contract.
func bindChallengeUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeUtils *ChallengeUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeUtils.Contract.ChallengeUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeUtils *ChallengeUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.ChallengeUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeUtils *ChallengeUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.ChallengeUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeUtils *ChallengeUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeUtils *ChallengeUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeUtils *ChallengeUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.contract.Transact(opts, method, params...)
}

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[]"

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201f1c2865a2b2af809ffe6caa0b1b6fbf9445248c5d42ae39aff757df53e40eb564736f6c634300050d0032"

// DeployDebugPrint deploys a new Ethereum contract, binding an instance of DebugPrint to it.
func DeployDebugPrint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebugPrint, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DebugPrintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// DebugPrint is an auto generated Go binding around an Ethereum contract.
type DebugPrint struct {
	DebugPrintCaller     // Read-only binding to the contract
	DebugPrintTransactor // Write-only binding to the contract
	DebugPrintFilterer   // Log filterer for contract events
}

// DebugPrintCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebugPrintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebugPrintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebugPrintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebugPrintSession struct {
	Contract     *DebugPrint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebugPrintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebugPrintCallerSession struct {
	Contract *DebugPrintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DebugPrintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebugPrintTransactorSession struct {
	Contract     *DebugPrintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DebugPrintRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebugPrintRaw struct {
	Contract *DebugPrint // Generic contract binding to access the raw methods on
}

// DebugPrintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebugPrintCallerRaw struct {
	Contract *DebugPrintCaller // Generic read-only contract binding to access the raw methods on
}

// DebugPrintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebugPrintTransactorRaw struct {
	Contract *DebugPrintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebugPrint creates a new instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrint(address common.Address, backend bind.ContractBackend) (*DebugPrint, error) {
	contract, err := bindDebugPrint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// NewDebugPrintCaller creates a new read-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintCaller(address common.Address, caller bind.ContractCaller) (*DebugPrintCaller, error) {
	contract, err := bindDebugPrint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintCaller{contract: contract}, nil
}

// NewDebugPrintTransactor creates a new write-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintTransactor(address common.Address, transactor bind.ContractTransactor) (*DebugPrintTransactor, error) {
	contract, err := bindDebugPrint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintTransactor{contract: contract}, nil
}

// NewDebugPrintFilterer creates a new log filterer instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintFilterer(address common.Address, filterer bind.ContractFilterer) (*DebugPrintFilterer, error) {
	contract, err := bindDebugPrint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebugPrintFilterer{contract: contract}, nil
}

// bindDebugPrint binds a generic wrapper to an already deployed contract.
func bindDebugPrint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.DebugPrintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transact(opts, method, params...)
}

// ExecutionChallengeABI is the input ABI used to generate the binding from.
const ExecutionChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"machineHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"didInboxInsns\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"logAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"gases\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalSteps\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"segmentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Continued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint128[4]\",\"name\":\"_timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_machineHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"_didInboxInsns\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_messageAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_logAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_gases\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_totalSteps\",\"type\":\"uint64\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_segmentToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"chooseSegment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxValueSize\",\"type\":\"uint256\"},{\"internalType\":\"uint128[4]\",\"name\":\"_timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsns\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExecutionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionChallengeFuncSigs = map[string]string{
	"a72c3ef4": "bisectAssertion(bytes32,uint128[4],bytes32[],bool[],bytes32[],bytes32[],uint64[],uint64)",
	"79a9ad85": "chooseSegment(uint256,bytes,bytes32,bytes32)",
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
	"f430e011": "oneStepProof(bytes32,bytes32,uint256,uint128[4],bytes32,bool,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
	"ced5c1bf": "timeoutChallenge()",
}

// ExecutionChallengeBin is the compiled bytecode used for deploying new contracts.
var ExecutionChallengeBin = "0x608060405234801561001057600080fd5b50611ff9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806302ad1e4e1461005c57806379a9ad85146100a0578063a72c3ef414610150578063ced5c1bf1461043d578063f430e01114610445575b600080fd5b61009e600480360360a081101561007257600080fd5b506001600160a01b0381358116916020810135821691604082013516906060810135906080013561055f565b005b61009e600480360360808110156100b657600080fd5b81359190810190604081016020820135600160201b8111156100d757600080fd5b8201836020820111156100e957600080fd5b803590602001918460018302840111600160201b8311171561010a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610574565b61009e600480360361016081101561016757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091949392602081019250359050600160201b8111156101b857600080fd5b8201836020820111156101ca57600080fd5b803590602001918460208302840111600160201b831117156101eb57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561023a57600080fd5b82018360208201111561024c57600080fd5b803590602001918460208302840111600160201b8311171561026d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102bc57600080fd5b8201836020820111156102ce57600080fd5b803590602001918460208302840111600160201b831117156102ef57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561033e57600080fd5b82018360208201111561035057600080fd5b803590602001918460208302840111600160201b8311171561037157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103c057600080fd5b8201836020820111156103d257600080fd5b803590602001918460208302840111600160201b831117156103f357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505050903567ffffffffffffffff1691506108729050565b61009e610a47565b61009e60048036036101e081101561045c57600080fd5b60408051608081810183528435946020810135949381013593810192909160e0830191906060840190600490839083908082843760009201919091525091948335946020850135151594604081013594506060810135935060808101359260a08201359267ffffffffffffffff60c0840135169261010081019060e00135600160201b8111156104eb57600080fd5b8201836020820111156104fd57600080fd5b803590602001918460018302840111600160201b8311171561051e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b27945050505050565b61056b85858585610ede565b60065550505050565b60055460ff16600281111561058557fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b815250906106335760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156105f85781810151838201526020016105e0565b50505050905090810190601f1680156106255780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060035461064043610ff9565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b815250906106b35760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b0316331461072f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b815250906107a25760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b506107b283838387600101611000565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b815250906108205760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50600681905561082e611101565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60055460ff16600281111561088357fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906108f45760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060035461090143610ff9565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906109745760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146109f05760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50610a3d6040518061010001604052808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018367ffffffffffffffff1681525061111d565b5050505050505050565b600354610a5343610ff9565b11610aa5576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff166002811115610ab857fe5b1415610af4576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a1610aef6118b1565b610b25565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a1610b256118bc565b565b60055460ff166002811115610b3857fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090610ba95760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50600354610bb643610ff9565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610c295760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610ca55760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50610cb98c8a8d8d8c8c888d8d8d8d6118c4565b600073__$efcad257db8183701794ea8506d55e247c$__63e987d8878e8c8f8f8e8e8e8e8e8e8e8e6040518d63ffffffff1660e01b8152600401808d81526020018c600460200280838360005b83811015610d1e578181015183820152602001610d06565b505050509050018b81526020018a8152602001898152602001881515151581526020018781526020018681526020018581526020018481526020018367ffffffffffffffff1667ffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610dac578181015183820152602001610d94565b50505050905090810190601f168015610dd95780820380516001836020036101000a031916815260200191505b509d505050505050505050505050505060206040518083038186803b158015610e0157600080fd5b505af4158015610e15573d6000803e3d6000fd5b505050506040513d6020811015610e2b57600080fd5b505160408051808201909152600981526827a9a82fa82927a7a360b91b60208201529091508115610e9d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b506040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1610ecf6118bc565b50505050505050505050505050565b600060055460ff166002811115610ef157fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b81525090610f665760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50600080546001600160a01b038681166001600160a01b03199283161790925560018054868416908316178155600280549386169390921692909217905560048290556005805460ff19169091179055610fbe6118fe565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e80290565b600080838160205b885181116110f3578089015193506020818a51036020018161102657fe5b0491505b60008211801561103d5750600286066001145b801561104b57508160020a86115b1561105e5760028604600101955061102a565b600286066110a95783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816110a157fe5b0495506110eb565b82846040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816110e457fe5b0460010195505b602001611008565b505090941495945050505050565b600580546001919060ff191682805b0217905550610b256118fe565b6000600182604001515103905081606001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906111a15760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5081608001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b8152509061121c5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b508160a001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906112975760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b508160c001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b8152509061130f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50600080805b83811015611363578460c00151818151811061132d57fe5b602002602001015183019250818061135957508460600151818151811061135057fe5b60200260200101515b9150600101611315565b506000611391856040015160008151811061137a57fe5b602002602001015186602001518760000151611910565b90506000611419866040015186815181106113a857fe5b6020026020010151848689608001516000815181106113c357fe5b60200260200101518a608001518a815181106113db57fe5b60200260200101518b60a001516000815181106113f457fe5b60200260200101518c60a001518c8151811061140c57fe5b602002602001015161198a565b905061143261142d8760e0015184846119ed565b611a35565b60608560405190808252806020026020018201604052801561145e578160200160208202803883390190505b50905061150b876040015160018151811061147557fe5b6020026020010151886060015160008151811061148e57fe5b60200260200101518960c001516000815181106114a757fe5b60200260200101518a608001516000815181106114c057fe5b60200260200101518b608001516001815181106114d957fe5b60200260200101518c60a001516000815181106114f257fe5b60200260200101518d60a0015160018151811061140c57fe5b915061155f6115288860e0015167ffffffffffffffff1688611aab565b63ffffffff16611559896040015160008151811061154257fe5b60200260200101518a602001518b60000151611910565b846119ed565b8160008151811061156c57fe5b602090810291909101015260015b868110156116cb578760600151600182038151811061159557fe5b6020026020010151156115ad576115aa611ac9565b88525b611659886040015182600101815181106115c357fe5b6020026020010151896060015183815181106115db57fe5b60200260200101518a60c0015184815181106115f357fe5b60200260200101518b60800151858151811061160b57fe5b60200260200101518c60800151866001018151811061162657fe5b60200260200101518d60a00151878151811061163e57fe5b60200260200101518e60a00151886001018151811061140c57fe5b92506116ac6116768960e0015167ffffffffffffffff1689611aea565b63ffffffff166116a68a60400151848151811061168f57fe5b60200260200101518b602001518c60000151611910565b856119ed565b8282815181106116b857fe5b602090810291909101015260010161157a565b506116d581611afd565b6116dd611b0c565b7f99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee478760400151886060015189608001518a60a001518b60c001518c60e001516003546040518080602001806020018060200180602001806020018867ffffffffffffffff1667ffffffffffffffff16815260200187815260200186810386528d818151815260200191508051906020019060200280838360005b8381101561178f578181015183820152602001611777565b5050505090500186810385528c818151815260200191508051906020019060200280838360005b838110156117ce5781810151838201526020016117b6565b5050505090500186810384528b818151815260200191508051906020019060200280838360005b8381101561180d5781810151838201526020016117f5565b5050505090500186810383528a818151815260200191508051906020019060200280838360005b8381101561184c578181015183820152602001611834565b50505050905001868103825289818151815260200191508051906020019060200280838360005b8381101561188b578181015183820152602001611873565b505050509050019c5050505050505050505050505060405180910390a150505050505050565b6118b9611b20565b33ff5b6118b9611b99565b60006118d08a8a611bfb565b905060006118df8d8d84611910565b9050610ecf61142d6001836118f98d8d8d8d8d8d8d61198a565b6119ed565b60045461190a43610ff9565b01600355565b81516020808401516040808601516060968701518251808601999099526fffffffffffffffffffffffffffffffff19608096871b81168a85015293861b841660508a015290851b83169688019690965294831b166070860152818501929092528251808503909101815260a0909301909152815191012090565b6040805160208082019990995296151560f81b8782015260c09590951b6001600160c01b031916604187015260498601939093526069850191909152608984015260a9808401919091528151808403909101815260c99092019052805191012090565b6040805160c09490941b6001600160c01b0319166020808601919091526028850193909352604880850192909252805180850390920182526068909301909252815191012090565b6006548114604051806040016040528060088152602001672124a9afa82922ab60c11b81525090611aa75760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5050565b6000818381611ab657fe5b06828481611ac057fe5b04019392505050565b60408051600080825260208201909252611ae4816001611c35565b91505090565b6000818381611af557fe5b049392505050565b611b0681611c5c565b60065550565b600580546002919060ff1916600183611110565b6000805460028054600154604080516335e1e69160e11b81526001600160a01b039384166004820152918316602483015260448201939093529151921692636bc3cd229260648084019382900301818387803b158015611b7f57600080fd5b505af1158015611b93573d6000803e3d6000fd5b50505050565b6000805460015460028054604080516335e1e69160e11b81526001600160a01b0394851660048201529184166024830152604482019290925290519190921692636bc3cd22926064808201939182900301818387803b158015611b7f57600080fd5b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b6000611c3f611f69565b611c498484611d9a565b9050611c5481611db9565b949350505050565b6000815b600181511115611d7d5760606002825160010181611c7a57fe5b04604051908082528060200260200182016040528015611ca4578160200160208202803883390190505b50905060005b8151811015611d75578251816002026001011015611d3d57828160020281518110611cd157fe5b6020026020010151838260020260010181518110611ceb57fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611d2c57fe5b602002602001018181525050611d6d565b828160020281518110611d4c57fe5b6020026020010151828281518110611d6057fe5b6020026020010181815250505b600101611caa565b509050611c60565b80600081518110611d8a57fe5b6020026020010151915050919050565b611da2611f69565b6000611dad84611e25565b9050611c548184611ee5565b606081015160009060ff16600214611e0e576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b81516080830151611e1f9190611bfb565b92915050565b6000600882511115611e75576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611eb9578181015183820152602001611ea1565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b611eed611f69565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611f52565b611f3f611f69565b815260200190600190039081611f375790505b508152600260208201526040019290925250919050565b6040518060a0016040528060008152602001611f83611f9d565b815260606020820181905260006040830181905291015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a7231582032c5e62b58fffbe80f05cf9d8efeb58db701097d78c2dffbc1bd79b05bed972164736f6c634300050d0032"

// DeployExecutionChallenge deploys a new Ethereum contract, binding an instance of ExecutionChallenge to it.
func DeployExecutionChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExecutionChallenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	oneStepProofAddr, _, _, _ := DeployOneStepProof(auth, backend)
	ExecutionChallengeBin = strings.Replace(ExecutionChallengeBin, "__$efcad257db8183701794ea8506d55e247c$__", oneStepProofAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExecutionChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutionChallenge{ExecutionChallengeCaller: ExecutionChallengeCaller{contract: contract}, ExecutionChallengeTransactor: ExecutionChallengeTransactor{contract: contract}, ExecutionChallengeFilterer: ExecutionChallengeFilterer{contract: contract}}, nil
}

// ExecutionChallenge is an auto generated Go binding around an Ethereum contract.
type ExecutionChallenge struct {
	ExecutionChallengeCaller     // Read-only binding to the contract
	ExecutionChallengeTransactor // Write-only binding to the contract
	ExecutionChallengeFilterer   // Log filterer for contract events
}

// ExecutionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionChallengeSession struct {
	Contract     *ExecutionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionChallengeCallerSession struct {
	Contract *ExecutionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ExecutionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionChallengeTransactorSession struct {
	Contract     *ExecutionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ExecutionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionChallengeRaw struct {
	Contract *ExecutionChallenge // Generic contract binding to access the raw methods on
}

// ExecutionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionChallengeCallerRaw struct {
	Contract *ExecutionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionChallengeTransactorRaw struct {
	Contract *ExecutionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionChallenge creates a new instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallenge(address common.Address, backend bind.ContractBackend) (*ExecutionChallenge, error) {
	contract, err := bindExecutionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallenge{ExecutionChallengeCaller: ExecutionChallengeCaller{contract: contract}, ExecutionChallengeTransactor: ExecutionChallengeTransactor{contract: contract}, ExecutionChallengeFilterer: ExecutionChallengeFilterer{contract: contract}}, nil
}

// NewExecutionChallengeCaller creates a new read-only instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeCaller(address common.Address, caller bind.ContractCaller) (*ExecutionChallengeCaller, error) {
	contract, err := bindExecutionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeCaller{contract: contract}, nil
}

// NewExecutionChallengeTransactor creates a new write-only instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionChallengeTransactor, error) {
	contract, err := bindExecutionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeTransactor{contract: contract}, nil
}

// NewExecutionChallengeFilterer creates a new log filterer instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionChallengeFilterer, error) {
	contract, err := bindExecutionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeFilterer{contract: contract}, nil
}

// bindExecutionChallenge binds a generic wrapper to an already deployed contract.
func bindExecutionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionChallenge *ExecutionChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExecutionChallenge.Contract.ExecutionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionChallenge *ExecutionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ExecutionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionChallenge *ExecutionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ExecutionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionChallenge *ExecutionChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExecutionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionChallenge *ExecutionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionChallenge *ExecutionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.contract.Transact(opts, method, params...)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xa72c3ef4.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, uint128[4] _timeBounds, bytes32[] _machineHashes, bool[] _didInboxInsns, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) BisectAssertion(opts *bind.TransactOpts, _beforeInbox [32]byte, _timeBounds [4]*big.Int, _machineHashes [][32]byte, _didInboxInsns []bool, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "bisectAssertion", _beforeInbox, _timeBounds, _machineHashes, _didInboxInsns, _messageAccs, _logAccs, _gases, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xa72c3ef4.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, uint128[4] _timeBounds, bytes32[] _machineHashes, bool[] _didInboxInsns, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) BisectAssertion(_beforeInbox [32]byte, _timeBounds [4]*big.Int, _machineHashes [][32]byte, _didInboxInsns []bool, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.BisectAssertion(&_ExecutionChallenge.TransactOpts, _beforeInbox, _timeBounds, _machineHashes, _didInboxInsns, _messageAccs, _logAccs, _gases, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xa72c3ef4.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, uint128[4] _timeBounds, bytes32[] _machineHashes, bool[] _didInboxInsns, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) BisectAssertion(_beforeInbox [32]byte, _timeBounds [4]*big.Int, _machineHashes [][32]byte, _didInboxInsns []bool, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.BisectAssertion(&_ExecutionChallenge.TransactOpts, _beforeInbox, _timeBounds, _machineHashes, _didInboxInsns, _messageAccs, _logAccs, _gases, _totalSteps)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) ChooseSegment(opts *bind.TransactOpts, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ChooseSegment(&_ExecutionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ChooseSegment(&_ExecutionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "initializeBisection", _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.InitializeBisection(&_ExecutionChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.InitializeBisection(&_ExecutionChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xf430e011.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint256 _beforeInboxValueSize, uint128[4] _timeBounds, bytes32 _afterHash, bool _didInboxInsns, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProof(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _beforeInboxValueSize *big.Int, _timeBounds [4]*big.Int, _afterHash [32]byte, _didInboxInsns bool, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "oneStepProof", _beforeHash, _beforeInbox, _beforeInboxValueSize, _timeBounds, _afterHash, _didInboxInsns, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xf430e011.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint256 _beforeInboxValueSize, uint128[4] _timeBounds, bytes32 _afterHash, bool _didInboxInsns, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) OneStepProof(_beforeHash [32]byte, _beforeInbox [32]byte, _beforeInboxValueSize *big.Int, _timeBounds [4]*big.Int, _afterHash [32]byte, _didInboxInsns bool, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.OneStepProof(&_ExecutionChallenge.TransactOpts, _beforeHash, _beforeInbox, _beforeInboxValueSize, _timeBounds, _afterHash, _didInboxInsns, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xf430e011.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint256 _beforeInboxValueSize, uint128[4] _timeBounds, bytes32 _afterHash, bool _didInboxInsns, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) OneStepProof(_beforeHash [32]byte, _beforeInbox [32]byte, _beforeInboxValueSize *big.Int, _timeBounds [4]*big.Int, _afterHash [32]byte, _didInboxInsns bool, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.OneStepProof(&_ExecutionChallenge.TransactOpts, _beforeHash, _beforeInbox, _beforeInboxValueSize, _timeBounds, _afterHash, _didInboxInsns, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.TimeoutChallenge(&_ExecutionChallenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.TimeoutChallenge(&_ExecutionChallenge.TransactOpts)
}

// ExecutionChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the ExecutionChallenge contract.
type ExecutionChallengeAsserterTimedOutIterator struct {
	Event *ExecutionChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeAsserterTimedOut)
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
		it.Event = new(ExecutionChallengeAsserterTimedOut)
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
func (it *ExecutionChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the ExecutionChallenge contract.
type ExecutionChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*ExecutionChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeAsserterTimedOutIterator{contract: _ExecutionChallenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeAsserterTimedOut)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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

// ParseAsserterTimedOut is a log parse operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*ExecutionChallengeAsserterTimedOut, error) {
	event := new(ExecutionChallengeAsserterTimedOut)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ExecutionChallenge contract.
type ExecutionChallengeBisectedAssertionIterator struct {
	Event *ExecutionChallengeBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeBisectedAssertion)
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
		it.Event = new(ExecutionChallengeBisectedAssertion)
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
func (it *ExecutionChallengeBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeBisectedAssertion represents a BisectedAssertion event raised by the ExecutionChallenge contract.
type ExecutionChallengeBisectedAssertion struct {
	MachineHashes [][32]byte
	DidInboxInsns []bool
	MessageAccs   [][32]byte
	LogAccs       [][32]byte
	Gases         []uint64
	TotalSteps    uint64
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0x99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee47.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bool[] didInboxInsns, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint64 totalSteps, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterBisectedAssertion(opts *bind.FilterOpts) (*ExecutionChallengeBisectedAssertionIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeBisectedAssertionIterator{contract: _ExecutionChallenge.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0x99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee47.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bool[] didInboxInsns, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint64 totalSteps, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeBisectedAssertion) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeBisectedAssertion)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0x99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee47.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bool[] didInboxInsns, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint64 totalSteps, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseBisectedAssertion(log types.Log) (*ExecutionChallengeBisectedAssertion, error) {
	event := new(ExecutionChallengeBisectedAssertion)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the ExecutionChallenge contract.
type ExecutionChallengeChallengerTimedOutIterator struct {
	Event *ExecutionChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeChallengerTimedOut)
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
		it.Event = new(ExecutionChallengeChallengerTimedOut)
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
func (it *ExecutionChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the ExecutionChallenge contract.
type ExecutionChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*ExecutionChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeChallengerTimedOutIterator{contract: _ExecutionChallenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeChallengerTimedOut)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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

// ParseChallengerTimedOut is a log parse operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*ExecutionChallengeChallengerTimedOut, error) {
	event := new(ExecutionChallengeChallengerTimedOut)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeContinuedIterator is returned from FilterContinued and is used to iterate over the raw logs and unpacked data for Continued events raised by the ExecutionChallenge contract.
type ExecutionChallengeContinuedIterator struct {
	Event *ExecutionChallengeContinued // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeContinued)
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
		it.Event = new(ExecutionChallengeContinued)
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
func (it *ExecutionChallengeContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeContinued represents a Continued event raised by the ExecutionChallenge contract.
type ExecutionChallengeContinued struct {
	SegmentIndex  *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContinued is a free log retrieval operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterContinued(opts *bind.FilterOpts) (*ExecutionChallengeContinuedIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeContinuedIterator{contract: _ExecutionChallenge.contract, event: "Continued", logs: logs, sub: sub}, nil
}

// WatchContinued is a free log subscription operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchContinued(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeContinued) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeContinued)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
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

// ParseContinued is a log parse operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseContinued(log types.Log) (*ExecutionChallengeContinued, error) {
	event := new(ExecutionChallengeContinued)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ExecutionChallenge contract.
type ExecutionChallengeInitiatedChallengeIterator struct {
	Event *ExecutionChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeInitiatedChallenge)
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
		it.Event = new(ExecutionChallengeInitiatedChallenge)
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
func (it *ExecutionChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the ExecutionChallenge contract.
type ExecutionChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ExecutionChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeInitiatedChallengeIterator{contract: _ExecutionChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeInitiatedChallenge)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ExecutionChallengeInitiatedChallenge, error) {
	event := new(ExecutionChallengeInitiatedChallenge)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ExecutionChallenge contract.
type ExecutionChallengeOneStepProofCompletedIterator struct {
	Event *ExecutionChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeOneStepProofCompleted)
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
		it.Event = new(ExecutionChallengeOneStepProofCompleted)
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
func (it *ExecutionChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the ExecutionChallenge contract.
type ExecutionChallengeOneStepProofCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*ExecutionChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeOneStepProofCompletedIterator{contract: _ExecutionChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeOneStepProofCompleted)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*ExecutionChallengeOneStepProofCompleted, error) {
	event := new(ExecutionChallengeOneStepProofCompleted)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IBisectionChallengeABI is the input ABI used to generate the binding from.
const IBisectionChallengeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBisectionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var IBisectionChallengeFuncSigs = map[string]string{
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
}

// IBisectionChallenge is an auto generated Go binding around an Ethereum contract.
type IBisectionChallenge struct {
	IBisectionChallengeCaller     // Read-only binding to the contract
	IBisectionChallengeTransactor // Write-only binding to the contract
	IBisectionChallengeFilterer   // Log filterer for contract events
}

// IBisectionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBisectionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBisectionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBisectionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBisectionChallengeSession struct {
	Contract     *IBisectionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IBisectionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBisectionChallengeCallerSession struct {
	Contract *IBisectionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IBisectionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBisectionChallengeTransactorSession struct {
	Contract     *IBisectionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IBisectionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBisectionChallengeRaw struct {
	Contract *IBisectionChallenge // Generic contract binding to access the raw methods on
}

// IBisectionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBisectionChallengeCallerRaw struct {
	Contract *IBisectionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// IBisectionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBisectionChallengeTransactorRaw struct {
	Contract *IBisectionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBisectionChallenge creates a new instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallenge(address common.Address, backend bind.ContractBackend) (*IBisectionChallenge, error) {
	contract, err := bindIBisectionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallenge{IBisectionChallengeCaller: IBisectionChallengeCaller{contract: contract}, IBisectionChallengeTransactor: IBisectionChallengeTransactor{contract: contract}, IBisectionChallengeFilterer: IBisectionChallengeFilterer{contract: contract}}, nil
}

// NewIBisectionChallengeCaller creates a new read-only instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeCaller(address common.Address, caller bind.ContractCaller) (*IBisectionChallengeCaller, error) {
	contract, err := bindIBisectionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeCaller{contract: contract}, nil
}

// NewIBisectionChallengeTransactor creates a new write-only instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBisectionChallengeTransactor, error) {
	contract, err := bindIBisectionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeTransactor{contract: contract}, nil
}

// NewIBisectionChallengeFilterer creates a new log filterer instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBisectionChallengeFilterer, error) {
	contract, err := bindIBisectionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeFilterer{contract: contract}, nil
}

// bindIBisectionChallenge binds a generic wrapper to an already deployed contract.
func bindIBisectionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBisectionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBisectionChallenge *IBisectionChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBisectionChallenge.Contract.IBisectionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBisectionChallenge *IBisectionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.IBisectionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBisectionChallenge *IBisectionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.IBisectionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBisectionChallenge *IBisectionChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBisectionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBisectionChallenge *IBisectionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBisectionChallenge *IBisectionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.contract.Transact(opts, method, params...)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.contract.Transact(opts, "initializeBisection", _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.InitializeBisection(&_IBisectionChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeTransactorSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.InitializeBisection(&_IBisectionChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// IStakingABI is the input ABI used to generate the binding from.
const IStakingABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IStakingFuncSigs maps the 4-byte function signature to its string representation.
var IStakingFuncSigs = map[string]string{
	"6bc3cd22": "resolveChallenge(address,address,uint256)",
}

// IStaking is an auto generated Go binding around an Ethereum contract.
type IStaking struct {
	IStakingCaller     // Read-only binding to the contract
	IStakingTransactor // Write-only binding to the contract
	IStakingFilterer   // Log filterer for contract events
}

// IStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingSession struct {
	Contract     *IStaking         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingCallerSession struct {
	Contract *IStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingTransactorSession struct {
	Contract     *IStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingRaw struct {
	Contract *IStaking // Generic contract binding to access the raw methods on
}

// IStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingCallerRaw struct {
	Contract *IStakingCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingTransactorRaw struct {
	Contract *IStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStaking creates a new instance of IStaking, bound to a specific deployed contract.
func NewIStaking(address common.Address, backend bind.ContractBackend) (*IStaking, error) {
	contract, err := bindIStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStaking{IStakingCaller: IStakingCaller{contract: contract}, IStakingTransactor: IStakingTransactor{contract: contract}, IStakingFilterer: IStakingFilterer{contract: contract}}, nil
}

// NewIStakingCaller creates a new read-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingCaller(address common.Address, caller bind.ContractCaller) (*IStakingCaller, error) {
	contract, err := bindIStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingCaller{contract: contract}, nil
}

// NewIStakingTransactor creates a new write-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingTransactor, error) {
	contract, err := bindIStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingTransactor{contract: contract}, nil
}

// NewIStakingFilterer creates a new log filterer instance of IStaking, bound to a specific deployed contract.
func NewIStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingFilterer, error) {
	contract, err := bindIStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingFilterer{contract: contract}, nil
}

// bindIStaking binds a generic wrapper to an already deployed contract.
func bindIStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.IStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transact(opts, method, params...)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x6bc3cd22.
//
// Solidity: function resolveChallenge(address winner, address loser, uint256 challengeType) returns()
func (_IStaking *IStakingTransactor) ResolveChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address, challengeType *big.Int) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "resolveChallenge", winner, loser, challengeType)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x6bc3cd22.
//
// Solidity: function resolveChallenge(address winner, address loser, uint256 challengeType) returns()
func (_IStaking *IStakingSession) ResolveChallenge(winner common.Address, loser common.Address, challengeType *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.ResolveChallenge(&_IStaking.TransactOpts, winner, loser, challengeType)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x6bc3cd22.
//
// Solidity: function resolveChallenge(address winner, address loser, uint256 challengeType) returns()
func (_IStaking *IStakingTransactorSession) ResolveChallenge(winner common.Address, loser common.Address, challengeType *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.ResolveChallenge(&_IStaking.TransactOpts, winner, loser, challengeType)
}

// MachineABI is the input ABI used to generate the binding from.
const MachineABI = "[]"

// MachineBin is the compiled bytecode used for deploying new contracts.
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200ef2d9def1af0771a70131496bac711a04c0e907d9256396f76870c5dc4594de64736f6c634300050d0032"

// DeployMachine deploys a new Ethereum contract, binding an instance of Machine to it.
func DeployMachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Machine, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Machine{MachineCaller: MachineCaller{contract: contract}, MachineTransactor: MachineTransactor{contract: contract}, MachineFilterer: MachineFilterer{contract: contract}}, nil
}

// Machine is an auto generated Go binding around an Ethereum contract.
type Machine struct {
	MachineCaller     // Read-only binding to the contract
	MachineTransactor // Write-only binding to the contract
	MachineFilterer   // Log filterer for contract events
}

// MachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineSession struct {
	Contract     *Machine          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineCallerSession struct {
	Contract *MachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineTransactorSession struct {
	Contract     *MachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineRaw struct {
	Contract *Machine // Generic contract binding to access the raw methods on
}

// MachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineCallerRaw struct {
	Contract *MachineCaller // Generic read-only contract binding to access the raw methods on
}

// MachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineTransactorRaw struct {
	Contract *MachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachine creates a new instance of Machine, bound to a specific deployed contract.
func NewMachine(address common.Address, backend bind.ContractBackend) (*Machine, error) {
	contract, err := bindMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Machine{MachineCaller: MachineCaller{contract: contract}, MachineTransactor: MachineTransactor{contract: contract}, MachineFilterer: MachineFilterer{contract: contract}}, nil
}

// NewMachineCaller creates a new read-only instance of Machine, bound to a specific deployed contract.
func NewMachineCaller(address common.Address, caller bind.ContractCaller) (*MachineCaller, error) {
	contract, err := bindMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineCaller{contract: contract}, nil
}

// NewMachineTransactor creates a new write-only instance of Machine, bound to a specific deployed contract.
func NewMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineTransactor, error) {
	contract, err := bindMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineTransactor{contract: contract}, nil
}

// NewMachineFilterer creates a new log filterer instance of Machine, bound to a specific deployed contract.
func NewMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineFilterer, error) {
	contract, err := bindMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineFilterer{contract: contract}, nil
}

// bindMachine binds a generic wrapper to an already deployed contract.
func bindMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machine *MachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Machine.Contract.MachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machine *MachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machine.Contract.MachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machine *MachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machine.Contract.MachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machine *MachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Machine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machine *MachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machine *MachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machine.Contract.contract.Transact(opts, method, params...)
}

// MerkleLibABI is the input ABI used to generate the binding from.
const MerkleLibABI = "[]"

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
var MerkleLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820837e4ef820a26b927f6d21de3452972fb4640e402d128b58d6eec0d200feb37564736f6c634300050d0032"

// DeployMerkleLib deploys a new Ethereum contract, binding an instance of MerkleLib to it.
func DeployMerkleLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleLib, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// MerkleLib is an auto generated Go binding around an Ethereum contract.
type MerkleLib struct {
	MerkleLibCaller     // Read-only binding to the contract
	MerkleLibTransactor // Write-only binding to the contract
	MerkleLibFilterer   // Log filterer for contract events
}

// MerkleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleLibSession struct {
	Contract     *MerkleLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleLibCallerSession struct {
	Contract *MerkleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MerkleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleLibTransactorSession struct {
	Contract     *MerkleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MerkleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleLibRaw struct {
	Contract *MerkleLib // Generic contract binding to access the raw methods on
}

// MerkleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleLibCallerRaw struct {
	Contract *MerkleLibCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleLibTransactorRaw struct {
	Contract *MerkleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleLib creates a new instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLib(address common.Address, backend bind.ContractBackend) (*MerkleLib, error) {
	contract, err := bindMerkleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// NewMerkleLibCaller creates a new read-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibCaller(address common.Address, caller bind.ContractCaller) (*MerkleLibCaller, error) {
	contract, err := bindMerkleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibCaller{contract: contract}, nil
}

// NewMerkleLibTransactor creates a new write-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleLibTransactor, error) {
	contract, err := bindMerkleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibTransactor{contract: contract}, nil
}

// NewMerkleLibFilterer creates a new log filterer instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleLibFilterer, error) {
	contract, err := bindMerkleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleLibFilterer{contract: contract}, nil
}

// bindMerkleLib binds a generic wrapper to an already deployed contract.
func bindMerkleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.MerkleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transact(opts, method, params...)
}

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128[4]\",\"name\":\"timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxValueSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"e987d887": "validateProof(bytes32,uint128[4],bytes32,uint256,bytes32,bool,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x613eec610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063e987d8871461003a575b600080fd5b61015c60048036036101e081101561005157600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100e757600080fd5b8201836020820111156100f957600080fd5b8035906020019184600183028401116401000000008311171561011b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061016e945050505050565b60408051918252519081900360200190f35b60006101d96040518061016001604052808f81526020018e81526020016101958e8e6101ea565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610270565b9d9c50505050505050505050505050565b6101f2613cc4565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610257565b610244613cc4565b81526020019060019003908161023c5790505b5081526002602082015260400183905290505b92915050565b6000806000806060610280613cf8565b610288613cf8565b610291886113d4565b93995092965090945092509050600160006102ab88611758565b67ffffffffffffffff168a610120015167ffffffffffffffff161461030e576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b89608001518015610322575060ff88166072145b8061033e5750896080015115801561033e575060ff8816607214155b61038f576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60ff8816600114156103d5576103ce83866000815181106103ac57fe5b6020026020010151876001815181106103c157fe5b6020026020010151611bf4565b915061121f565b60ff881660021415610414576103ce83866000815181106103f257fe5b60200260200101518760018151811061040757fe5b6020026020010151611c44565b60ff881660031415610453576103ce838660008151811061043157fe5b60200260200101518760018151811061044657fe5b6020026020010151611c85565b60ff881660041415610492576103ce838660008151811061047057fe5b60200260200101518760018151811061048557fe5b6020026020010151611cc6565b60ff8816600514156104d1576103ce83866000815181106104af57fe5b6020026020010151876001815181106104c457fe5b6020026020010151611d17565b60ff881660061415610510576103ce83866000815181106104ee57fe5b60200260200101518760018151811061050357fe5b6020026020010151611d68565b60ff88166007141561054f576103ce838660008151811061052d57fe5b60200260200101518760018151811061054257fe5b6020026020010151611db9565b60ff8816600814156105a3576103ce838660008151811061056c57fe5b60200260200101518760018151811061058157fe5b60200260200101518860028151811061059657fe5b6020026020010151611e0a565b60ff8816600914156105f7576103ce83866000815181106105c057fe5b6020026020010151876001815181106105d557fe5b6020026020010151886002815181106105ea57fe5b6020026020010151611e74565b60ff8816600a1415610636576103ce838660008151811061061457fe5b60200260200101518760018151811061062957fe5b6020026020010151611ecd565b60ff881660101415610675576103ce838660008151811061065357fe5b60200260200101518760018151811061066857fe5b6020026020010151611f0e565b60ff8816601114156106b4576103ce838660008151811061069257fe5b6020026020010151876001815181106106a757fe5b6020026020010151611f4f565b60ff8816601214156106f3576103ce83866000815181106106d157fe5b6020026020010151876001815181106106e657fe5b6020026020010151611f90565b60ff881660131415610732576103ce838660008151811061071057fe5b60200260200101518760018151811061072557fe5b6020026020010151611fd1565b60ff881660141415610771576103ce838660008151811061074f57fe5b60200260200101518760018151811061076457fe5b6020026020010151612012565b60ff88166015141561079b576103ce838660008151811061078e57fe5b602002602001015161203c565b60ff8816601614156107da576103ce83866000815181106107b857fe5b6020026020010151876001815181106107cd57fe5b6020026020010151612081565b60ff881660171415610819576103ce83866000815181106107f757fe5b60200260200101518760018151811061080c57fe5b60200260200101516120c2565b60ff881660181415610858576103ce838660008151811061083657fe5b60200260200101518760018151811061084b57fe5b6020026020010151612103565b60ff881660191415610882576103ce838660008151811061087557fe5b6020026020010151612144565b60ff8816601a14156108c1576103ce838660008151811061089f57fe5b6020026020010151876001815181106108b457fe5b602002602001015161217a565b60ff8816601b1415610900576103ce83866000815181106108de57fe5b6020026020010151876001815181106108f357fe5b60200260200101516121bb565b60ff88166020141561092a576103ce838660008151811061091d57fe5b60200260200101516121fc565b60ff881660211415610954576103ce838660008151811061094757fe5b6020026020010151612217565b60ff881660221415610993576103ce838660008151811061097157fe5b60200260200101518760018151811061098657fe5b6020026020010151612232565b60ff8816603014156109bd576103ce83866000815181106109b057fe5b6020026020010151612298565b60ff8816603114156109d2576103ce836122a0565b60ff8816603214156109e7576103ce836122c1565b60ff881660331415610a11576103ce8386600081518110610a0457fe5b60200260200101516122da565b60ff881660341415610a3b576103ce8386600081518110610a2e57fe5b60200260200101516122e6565b60ff881660351415610a7a576103ce8386600081518110610a5857fe5b602002602001015187600181518110610a6d57fe5b60200260200101516122ed565b60ff881660361415610a8f576103ce83612329565b60ff881660371415610aa9576103ce838560000151612353565b60ff881660381415610ad3576103ce8386600081518110610ac657fe5b6020026020010151612365565b60ff881660391415610b5f57610ae7613cc4565b610af68b610140015188612377565b9199509750905087610b395760405162461bcd60e51b8152600401808060200182810382526021815260200180613e976021913960400191505060405180910390fd5b610b49858263ffffffff6124b516565b610b59848263ffffffff6124cf16565b5061121f565b60ff8816603a1415610b74576103ce836124e9565b60ff8816603b1415610b855761121f565b60ff8816603c1415610b9a576103ce83612506565b60ff8816603d1415610bc4576103ce8386600081518110610bb757fe5b602002602001015161251f565b60ff881660401415610bee576103ce8386600081518110610be157fe5b6020026020010151612543565b60ff881660411415610c2d576103ce8386600081518110610c0b57fe5b602002602001015187600181518110610c2057fe5b6020026020010151612565565b60ff881660421415610c81576103ce8386600081518110610c4a57fe5b602002602001015187600181518110610c5f57fe5b602002602001015188600281518110610c7457fe5b6020026020010151612597565b60ff881660431415610cc0576103ce8386600081518110610c9e57fe5b602002602001015187600181518110610cb357fe5b60200260200101516125d9565b60ff881660441415610d14576103ce8386600081518110610cdd57fe5b602002602001015187600181518110610cf257fe5b602002602001015188600281518110610d0757fe5b60200260200101516125eb565b60ff881660501415610d53576103ce8386600081518110610d3157fe5b602002602001015187600181518110610d4657fe5b602002602001015161260d565b60ff881660511415610da7576103ce8386600081518110610d7057fe5b602002602001015187600181518110610d8557fe5b602002602001015188600281518110610d9a57fe5b6020026020010151612683565b60ff881660521415610dd1576103ce8386600081518110610dc457fe5b6020026020010151612710565b60ff881660601415610de6576103ce83612743565b60ff881660611415610ee457610e108386600081518110610e0357fe5b6020026020010151612749565b90925090508115610edb578961010001518a60e001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610e905760405162461bcd60e51b8152600401808060200182810382526025815260200180613e4b6025913960400191505060405180910390fd5b8960c001518a60a0015114610ed65760405162461bcd60e51b8152600401808060200182810382526027815260200180613e706027913960400191505060405180910390fd5b610edf565b5060005b61121f565b60ff88166070141561102457610f0e8386600081518110610f0157fe5b6020026020010151612763565b90925090508115610edb5780610f69578960c001518a60a0015114610f645760405162461bcd60e51b8152600401808060200182810382526038815260200180613e136038913960400191505060405180910390fd5b610ed6565b8960c001518a60a001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610fdd5760405162461bcd60e51b8152600401808060200182810382526029815260200180613da36029913960400191505060405180910390fd5b8961010001518a60e0015114610ed65760405162461bcd60e51b8152600401808060200182810382526026815260200180613dcc6026913960400191505060405180910390fd5b60ff8816607114156111395760408051600480825260a08201909252606091816020015b611050613cc4565b81526020019060019003908161104857505060208c01519091506110849060005b60200201516001600160801b0316612793565b8160008151811061109157fe5b60200260200101819052506110b08b6020015160016004811061107157fe5b816001815181106110bd57fe5b60200260200101819052506110dc8b6020015160026004811061107157fe5b816002815181106110e957fe5b60200260200101819052506111088b6020015160036004811061107157fe5b8160038151811061111557fe5b6020026020010181905250610b5961112c82612818565b859063ffffffff6124cf16565b60ff881660721415611187576103ce838660008151811061115657fe5b60200260200101518c604001518d6020015160006004811061117457fe5b60200201516001600160801b0316612907565b60ff88166073141561119c576000915061121f565b60ff8816607414156111b157610edf8361299e565b60ff88166080141561121a576103ce83866000815181106111ce57fe5b6020026020010151876001815181106111e357fe5b6020026020010151886002815181106111f857fe5b60200260200101518960038151811061120d57fe5b60200260200101516129a8565b600091505b806112b1578960c001518a60a001511461126a5760405162461bcd60e51b8152600401808060200182810382526027815260200180613e706027913960400191505060405180910390fd5b8961010001518a60e00151146112b15760405162461bcd60e51b8152600401808060200182810382526026815260200180613dcc6026913960400191505060405180910390fd5b8161131b5760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a08401516112ff90612ac7565b14156113135761130e83612bd5565b61131b565b60a083015183525b61132484612bdf565b8a51146113625760405162461bcd60e51b8152600401808060200182810382526022815260200180613d816022913960400191505060405180910390fd5b61136b83612bdf565b8a60600151146113c2576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b600060606113e0613cf8565b6113e8613cf8565b600080806113f4613cf8565b6113fd81612ca7565b61140c89610140015184612cb1565b909450909250905081611466576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b61146e613cf8565b61147782612dc2565b905060008a6101400151858151811061148c57fe5b602001015160f81c60f81b60f81c905060008b610140015186600101815181106114b257fe5b016020015160f81c905060006114c782612e20565b905060608160405190808252806020026020018201604052801561150557816020015b6114f2613cc4565b8152602001906001900390816114ea5790505b5090506002880197508360ff166000148061152357508360ff166001145b611574576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff8416611599576115928361158d8860000151612ac7565b612e3a565b8652611661565b6115a1613cc4565b6115b08f61014001518a612377565b909a5090985090508761160a576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561162e57808260008151811061161e57fe5b602002602001018190525061163e565b61163e868263ffffffff6124cf16565b61165d8461164f8960000151612ac7565b61165884612ac7565b612e74565b8752505b60ff84165b828110156116f45761167d8f61014001518a612377565b845185908590811061168b57fe5b602090810291909101015299509750876116ec576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611666565b815115611741575060005b8460ff168251038110156117415761173982826001855103038151811061172257fe5b6020026020010151886124cf90919063ffffffff16565b6001016116ff565b50919d919c50939a50919850939650945050505050565b600060ff82166001141561176e575060036113cf565b60ff821660021415611782575060036113cf565b60ff821660031415611796575060036113cf565b60ff8216600414156117aa575060046113cf565b60ff8216600514156117be575060076113cf565b60ff8216600614156117d2575060046113cf565b60ff8216600714156117e6575060076113cf565b60ff8216600814156117fa575060046113cf565b60ff82166009141561180e575060046113cf565b60ff8216600a1415611822575060196113cf565b60ff821660101415611836575060026113cf565b60ff82166011141561184a575060026113cf565b60ff82166012141561185e575060026113cf565b60ff821660131415611872575060026113cf565b60ff821660141415611886575060026113cf565b60ff82166015141561189a575060016113cf565b60ff8216601614156118ae575060026113cf565b60ff8216601714156118c2575060026113cf565b60ff8216601814156118d6575060026113cf565b60ff8216601914156118ea575060016113cf565b60ff8216601a14156118fe575060046113cf565b60ff8216601b1415611912575060076113cf565b60ff821660201415611926575060076113cf565b60ff82166021141561193a575060036113cf565b60ff82166022141561194e575060086113cf565b60ff821660301415611962575060016113cf565b60ff821660311415611976575060016113cf565b60ff82166032141561198a575060016113cf565b60ff82166033141561199e575060026113cf565b60ff8216603414156119b2575060046113cf565b60ff8216603514156119c6575060046113cf565b60ff8216603614156119da575060026113cf565b60ff8216603714156119ee575060016113cf565b60ff821660381415611a02575060016113cf565b60ff821660391415611a16575060016113cf565b60ff8216603a1415611a2a575060026113cf565b60ff8216603b1415611a3e575060016113cf565b60ff8216603c1415611a52575060016113cf565b60ff8216603d1415611a66575060016113cf565b60ff821660401415611a7a575060016113cf565b60ff821660411415611a8e575060016113cf565b60ff821660421415611aa2575060016113cf565b60ff821660431415611ab6575060016113cf565b60ff821660441415611aca575060016113cf565b60ff821660501415611ade575060026113cf565b60ff821660511415611af2575060286113cf565b60ff821660521415611b06575060026113cf565b60ff821660601415611b1a575060646113cf565b60ff821660611415611b2e575060646113cf565b60ff821660701415611b42575060646113cf565b60ff821660711415611b56575060286113cf565b60ff821660721415611b6a575060286113cf565b60ff821660731415611b7e575060056113cf565b60ff821660741415611b925750600a6113cf565b60ff821660801415611ba75750614e206113cf565b6040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206f70636f64653a206f70476173436f737428290000000000604482015290519081900360640190fd5b6000611bff83612eab565b1580611c115750611c0f82612eab565b155b15611c1e57506000611c3d565b82518251808201611c35878263ffffffff612eb616565b600193505050505b9392505050565b6000611c4f83612eab565b1580611c615750611c5f82612eab565b155b15611c6e57506000611c3d565b82518251808202611c35878263ffffffff612eb616565b6000611c9083612eab565b1580611ca25750611ca082612eab565b155b15611caf57506000611c3d565b82518251808203611c35878263ffffffff612eb616565b6000611cd183612eab565b1580611ce35750611ce182612eab565b155b15611cf057506000611c3d565b8251825180611d0457600092505050611c3d565b808204611c35878263ffffffff612eb616565b6000611d2283612eab565b1580611d345750611d3282612eab565b155b15611d4157506000611c3d565b8251825180611d5557600092505050611c3d565b808205611c35878263ffffffff612eb616565b6000611d7383612eab565b1580611d855750611d8382612eab565b155b15611d9257506000611c3d565b8251825180611da657600092505050611c3d565b808206611c35878263ffffffff612eb616565b6000611dc483612eab565b1580611dd65750611dd482612eab565b155b15611de357506000611c3d565b8251825180611df757600092505050611c3d565b808207611c35878263ffffffff612eb616565b6000611e1584612eab565b1580611e275750611e2583612eab565b155b15611e3457506000611e6c565b83518351835180611e4b5760009350505050611e6c565b6000818385089050611e63898263ffffffff612eb616565b60019450505050505b949350505050565b6000611e7f84612eab565b1580611e915750611e8f83612eab565b155b15611e9e57506000611e6c565b83518351835180611eb55760009350505050611e6c565b6000818385099050611e63898263ffffffff612eb616565b6000611ed883612eab565b1580611eea5750611ee882612eab565b155b15611ef757506000611c3d565b8251825180820a611c35878263ffffffff612eb616565b6000611f1983612eab565b1580611f2b5750611f2982612eab565b155b15611f3857506000611c3d565b82518251808210611c35878263ffffffff612eb616565b6000611f5a83612eab565b1580611f6c5750611f6a82612eab565b155b15611f7957506000611c3d565b82518251808211611c35878263ffffffff612eb616565b6000611f9b83612eab565b1580611fad5750611fab82612eab565b155b15611fba57506000611c3d565b82518251808212611c35878263ffffffff612eb616565b6000611fdc83612eab565b1580611fee5750611fec82612eab565b155b15611ffb57506000611c3d565b82518251808213611c35878263ffffffff612eb616565b600061203261112c61202384612ac7565b61202c86612ac7565b14612ecc565b5060019392505050565b600061204782612eab565b6120615761205c83600063ffffffff612eb616565b612078565b81518015612075858263ffffffff612eb616565b50505b50600192915050565b600061208c83612eab565b158061209e575061209c82612eab565b155b156120ab57506000611c3d565b82518251808216611c35878263ffffffff612eb616565b60006120cd83612eab565b15806120df57506120dd82612eab565b155b156120ec57506000611c3d565b82518251808217611c35878263ffffffff612eb616565b600061210e83612eab565b1580612120575061211e82612eab565b155b1561212d57506000611c3d565b82518251808218611c35878263ffffffff612eb616565b600061214f82612eab565b61215b5750600061026a565b8151801961216f858263ffffffff612eb616565b506001949350505050565b600061218583612eab565b1580612197575061219582612eab565b155b156121a457506000611c3d565b8251825181811a611c35878263ffffffff612eb616565b60006121c683612eab565b15806121d857506121d682612eab565b155b156121e557506000611c3d565b8251825181810b611c35878263ffffffff612eb616565b600061207861220a83612ac7565b849063ffffffff612eb616565b600061207861222583612eee565b849063ffffffff6124cf16565b600061223d83612eab565b158061224f575061224d82612eab565b155b1561225c57506000611c3d565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120611c35878263ffffffff612eb616565b600192915050565b60006122b98260800151836124cf90919063ffffffff16565b506001919050565b60006122b98260600151836124cf90919063ffffffff16565b60609190910152600190565b9052600190565b60006122f883612f77565b61230457506000611c3d565b61230d82612eab565b61231957506000611c3d565b8151156120325750509052600190565b60006122b9612346612339612f84565b61202c8560200151612ac7565b839063ffffffff6124cf16565b6000612078838363ffffffff6124cf16565b6000612078838363ffffffff6124b516565b600080612382613cc4565b845184106123a2576000846123976000612793565b9250925092506124ae565b60008085905060008782815181106123b657fe5b016020015160019092019160f81c905060006123d0613d59565b60ff8316612404576123e28a85612fa5565b9196509450915084846123f484612793565b97509750975050505050506124ae565b60ff83166001141561242c5761241a8a85612ff8565b9196509450905084846123f483613158565b60ff831660021415612442576123f48a856131bf565b600360ff8416108015906124595750600c60ff8416105b156124945760021983016060612470828d88613264565b91985096509050868661248283612818565b995099509950505050505050506124ae565b6000806124a16000612793565b9199509750955050505050505b9250925092565b6124c3826040015182613322565b82604001819052505050565b6124dd826020015182613322565b82602001819052505050565b60006122b96123466124f9612f84565b61202c8560400151612ac7565b60006122b98260a00151836124cf90919063ffffffff16565b600061252a82612f77565b6125365750600061026a565b5060a09190910152600190565b6000612555838363ffffffff6124cf16565b612078838363ffffffff6124cf16565b6000612577848363ffffffff6124cf16565b612587848463ffffffff6124cf16565b612032848363ffffffff6124cf16565b60006125a9858363ffffffff6124cf16565b6125b9858463ffffffff6124cf16565b6125c9858563ffffffff6124cf16565b61216f858363ffffffff6124cf16565b6000612587848463ffffffff6124cf16565b60006125fd858563ffffffff6124cf16565b6125c9858463ffffffff6124cf16565b600061261883612eab565b158061262a5750612628826133a0565b155b1561263757506000611c3d565b612640826133af565b60ff1683600001511061265557506000611c3d565b612032826040015184600001518151811061266c57fe5b6020026020010151856124cf90919063ffffffff16565b600061268e836133a0565b15806126a0575061269e84612eab565b155b156126ad57506000611e6c565b6126b6836133af565b60ff168460000151106126cb57506000611e6c565b6040830151845181518491839181106126e057fe5b60200260200101819052506127046126f782612818565b879063ffffffff6124cf16565b50600195945050505050565b600061271b826133a0565b6127275750600061026a565b612078612733836133af565b849060ff1663ffffffff612eb616565b50600190565b600080600161275784612ac7565b915091505b9250929050565b60008061271083608001511161278757600161277e84612ac7565b9150915061275c565b5060019050600061275c565b61279b613cc4565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612800565b6127ed613cc4565b8152602001906001900390816127e55790505b50815260006020820152600160409091015292915050565b612820613cc4565b61282a82516133be565b61287b576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156128b25783818151811061289557fe5b602002602001015160800151820191508080600101915050612880565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b600061291284612eab565b61291e57506000611e6c565b83518211158061293d5750612931612f84565b61293a84612ac7565b14155b61298e576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b61216f858463ffffffff6124cf16565b600260c090910152565b60006129b385612eab565b15806129c557506129c384612eab565b155b806129d657506129d483612eab565b155b806129e757506129e582612eab565b155b156129f457506000612abe565b84518451845115801590612a0a57508451600114155b15612a2b57612a2088600063ffffffff612eb616565b600192505050612abe565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612a8d573d6000803e3d6000fd5b5050604051601f1901519150612ab490508b6001600160a01b03831663ffffffff612eb616565b6001955050505050505b95945050505050565b6000600360090160ff16826060015160ff1610612b1f576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16612b3d578151612b36906133c5565b90506113cf565b606082015160ff1660011415612b70576020808301518051604082015160608301519290930151612b36939192906133e9565b606082015160ff1660021415612ba157602080830151015160011415612b98575080516113cf565b612b3682613491565b600360ff16826060015160ff1610158015612bc557506060820151600c60ff909116105b15612bd357612b36826134f7565bfe5b600160c090910152565b600060028260c001511415612bf6575060006113cf565b60018260c001511415612c0b575060016113cf565b8151612c1690612ac7565b612c238360200151612ac7565b612c308460400151612ac7565b612c3d8560600151612ac7565b612c4a8660800151612ac7565b612c578760a00151612ac7565b6040516020018087815260200186815260200185815260200184815260200183815260200182815260200196505050505050506040516020818303038152906040528051906020012090506113cf565b600060c090910152565b600080612cbc613cf8565b612cc4613cf8565b600060c08201819052612cd78787612377565b84529650905080612cf157506000935084925090506124ae565b612cfb87876131bf565b60208501529650905080612d1857506000935084925090506124ae565b612d2287876131bf565b60408501529650905080612d3f57506000935084925090506124ae565b612d498787612377565b60608501529650905080612d6657506000935084925090506124ae565b612d708787612377565b60808501529650905080612d8d57506000935084925090506124ae565b612d978787612377565b60a08501529650905080612db457506000935084925090506124ae565b506001969495509392505050565b612dca613cf8565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b6000806000612e318460ff16613515565b50949350505050565b612e42613cc4565b611c3d60405180608001604052808560ff1681526020018481526020016000151581526020016000801b815250613158565b612e7c613cc4565b611e6c60405180608001604052808660ff16815260200185815260200160011515815260200184815250613158565b6060015160ff161590565b6124dd8260200151612ec783612793565b613322565b612ed4613cc4565b8115612ee457612b366001612793565b612b366000612793565b612ef6613cc4565b816060015160ff1660021415612f3d5760405162461bcd60e51b8152600401808060200182810382526021815260200180613df26021913960400191505060405180910390fd5b606082015160ff16612f5357612b366000612793565b816060015160ff1660011415612f6d57612b366001612793565b612b366003612793565b6060015160ff1660011490565b60408051600080825260208201909252612f9f8160016139e1565b91505090565b6000806000808551905084811080612fbf57506020858203105b15612fd45750600092508391508290506124ae565b600160208601612fea888863ffffffff613a0016565b935093509350509250925092565b600080613003613d59565b6000849050600086828151811061301657fe5b602001015160f81c60f81b60f81c90508180600101925050600087838151811061303c57fe5b016020015160019384019360f89190911c915060009060ff841614156130c2576000613066613cc4565b6130708b87612377565b9097509092509050816130b4575050604080516080810182526000808252602082018190529181018290526060810182905290975088965094506124ae9350505050565b6130bd81612ac7565b925050505b60006130d48a8663ffffffff613a0016565b90506020850194508360ff1660011415613120576040805160808101825260ff909416845260208401919091526001908301819052606083019190915295509193509091506124ae9050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b613160613cc4565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906131a7565b613194613cc4565b81526020019060019003908161318c5790505b50815260016020820181905260409091015292915050565b6000806131ca613cc4565b6131d2613cc4565b85516000908190878110806131e957506040888203105b156132015760008885965096509650505050506124ae565b60006132138a8a63ffffffff613a0016565b90506020890198506132258a8a612fa5565b909a509450925082156132505761323c81856101ea565b6001985089975095506124ae945050505050565b6000898697509750975050505050506124ae565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156132af57816020015b61329c613cc4565b8152602001906001900390816132945790505b50905060005b8960ff168160ff16101561330c576132cd8985612377565b8451859060ff86169081106132de57fe5b602090810291909101015294509250826133045750600095508694509250613319915050565b6001016132b5565b5060019550919350909150505b93509350939050565b61332a613cc4565b6040805160028082526060828101909352816020015b613348613cc4565b815260200190600190039081613340579050509050828160008151811061336b57fe5b6020026020010181905250838160018151811061338457fe5b6020026020010181905250611e6c61339b82612818565b613a1c565b600061026a8260600151613a92565b600061026a8260600151613ab0565b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613443575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611e6c565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff166002146134e6576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161026a9190613ad3565b6000613501613cc4565b61350a83613a1c565b9050611c3d81613491565b600080600183141561352d57506002905060016139dc565b600283141561354257506002905060016139dc565b600383141561355757506002905060016139dc565b600483141561356c57506002905060016139dc565b600583141561358157506002905060016139dc565b600683141561359657506002905060016139dc565b60078314156135ab57506002905060016139dc565b60088314156135c057506003905060016139dc565b60098314156135d557506003905060016139dc565b600a8314156135ea57506002905060016139dc565b60108314156135ff57506002905060016139dc565b601183141561361457506002905060016139dc565b601283141561362957506002905060016139dc565b601383141561363e57506002905060016139dc565b601483141561365357506002905060016139dc565b6015831415613667575060019050806139dc565b601683141561367c57506002905060016139dc565b601783141561369157506002905060016139dc565b60188314156136a657506002905060016139dc565b60198314156136ba575060019050806139dc565b601a8314156136cf57506002905060016139dc565b601b8314156136e457506002905060016139dc565b60208314156136f8575060019050806139dc565b602183141561370c575060019050806139dc565b602283141561372157506002905060016139dc565b603083141561373657506001905060006139dc565b603183141561374b57506000905060016139dc565b603283141561376057506000905060016139dc565b603383141561377557506001905060006139dc565b603483141561378a57506001905060006139dc565b603583141561379f57506002905060006139dc565b60368314156137b457506000905060016139dc565b60378314156137c957506000905060016139dc565b60388314156137de57506001905060006139dc565b60398314156137f357506000905060016139dc565b603a83141561380857506000905060016139dc565b603b83141561381c575060009050806139dc565b603c83141561383157506000905060016139dc565b603d83141561384657506001905060006139dc565b604083141561385b57506001905060026139dc565b604183141561387057506002905060036139dc565b604283141561388557506003905060046139dc565b6043831415613899575060029050806139dc565b60448314156138ad575060039050806139dc565b60508314156138c257506002905060016139dc565b60518314156138d757506003905060016139dc565b60528314156138eb575060019050806139dc565b60608314156138ff575060009050806139dc565b606183141561391457506001905060006139dc565b607083141561392957506001905060006139dc565b607183141561393e57506000905060016139dc565b6072831415613952575060019050806139dc565b6073831415613966575060009050806139dc565b607483141561397a575060009050806139dc565b608083141561398f57506004905060016139dc565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b60006139eb613cc4565b6139f58484613b0d565b9050611e6c81613491565b60008160200183511015613a1357600080fd5b50016020015190565b613a24613cc4565b613a2d826133a0565b613a73576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060613a828360400151613b2c565b9050611c3d818460800151613b0d565b6000600c60ff831610801561026a575050600360ff91909116101590565b6000613abb82613a92565b15613acb575060021981016113cf565b5060016113cf565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b613b15613cc4565b6000613b2084613c04565b9050611e6c81846101ea565b6060600882511115613b7c576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613ba9578160200160208202803883390190505b50805190915060005b81811015613bfb576000613bd8868381518110613bcb57fe5b6020026020010151612ac7565b905080848381518110613be757fe5b602090810291909101015250600101613bb2565b50909392505050565b6000600882511115613c54576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613c98578181015183820152602001613c80565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001613cde613d59565b815260606020820181905260006040830181905291015290565b6040518060e00160405280613d0b613cc4565b8152602001613d18613cc4565b8152602001613d25613cc4565b8152602001613d32613cc4565b8152602001613d3f613cc4565b8152602001613d4c613cc4565b8152602001600081525090565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820797294b5be74afeebaff3e76a1c255ed64f6db4833e72cdf6e8c3b9fda2d0efc64736f6c634300050d0032"

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
func (_OneStepProof *OneStepProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_OneStepProof *OneStepProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// ValidateProof is a free data retrieval call binding the contract method 0xe987d887.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, beforeHash [32]byte, timeBounds [4]*big.Int, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", beforeHash, timeBounds, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0xe987d887.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(beforeHash [32]byte, timeBounds [4]*big.Int, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBounds, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0xe987d887.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(beforeHash [32]byte, timeBounds [4]*big.Int, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBounds, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820da8d11bf1947a4571c85676ea8b846317a3049a34118d9ab59be573dd308ff8e64736f6c634300050d0032"

// DeployProtocol deploys a new Ethereum contract, binding an instance of Protocol to it.
func DeployProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Protocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// Protocol is an auto generated Go binding around an Ethereum contract.
type Protocol struct {
	ProtocolCaller     // Read-only binding to the contract
	ProtocolTransactor // Write-only binding to the contract
	ProtocolFilterer   // Log filterer for contract events
}

// ProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolSession struct {
	Contract     *Protocol         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolCallerSession struct {
	Contract *ProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolTransactorSession struct {
	Contract     *ProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRaw struct {
	Contract *Protocol // Generic contract binding to access the raw methods on
}

// ProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolCallerRaw struct {
	Contract *ProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolTransactorRaw struct {
	Contract *ProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocol creates a new instance of Protocol, bound to a specific deployed contract.
func NewProtocol(address common.Address, backend bind.ContractBackend) (*Protocol, error) {
	contract, err := bindProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// NewProtocolCaller creates a new read-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolCaller(address common.Address, caller bind.ContractCaller) (*ProtocolCaller, error) {
	contract, err := bindProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolCaller{contract: contract}, nil
}

// NewProtocolTransactor creates a new write-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolTransactor, error) {
	contract, err := bindProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTransactor{contract: contract}, nil
}

// NewProtocolFilterer creates a new log filterer instance of Protocol, bound to a specific deployed contract.
func NewProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolFilterer, error) {
	contract, err := bindProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolFilterer{contract: contract}, nil
}

// bindProtocol binds a generic wrapper to an already deployed contract.
func bindProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.ProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transact(opts, method, params...)
}

// RollupTimeABI is the input ABI used to generate the binding from.
const RollupTimeABI = "[]"

// RollupTimeBin is the compiled bytecode used for deploying new contracts.
var RollupTimeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582037f06363d238a70c34530f76a5f2c329256ab1414b55b90256aa470c78073d0964736f6c634300050d0032"

// DeployRollupTime deploys a new Ethereum contract, binding an instance of RollupTime to it.
func DeployRollupTime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupTime, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupTimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupTime{RollupTimeCaller: RollupTimeCaller{contract: contract}, RollupTimeTransactor: RollupTimeTransactor{contract: contract}, RollupTimeFilterer: RollupTimeFilterer{contract: contract}}, nil
}

// RollupTime is an auto generated Go binding around an Ethereum contract.
type RollupTime struct {
	RollupTimeCaller     // Read-only binding to the contract
	RollupTimeTransactor // Write-only binding to the contract
	RollupTimeFilterer   // Log filterer for contract events
}

// RollupTimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupTimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupTimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupTimeSession struct {
	Contract     *RollupTime       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupTimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupTimeCallerSession struct {
	Contract *RollupTimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupTimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTimeTransactorSession struct {
	Contract     *RollupTimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupTimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupTimeRaw struct {
	Contract *RollupTime // Generic contract binding to access the raw methods on
}

// RollupTimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupTimeCallerRaw struct {
	Contract *RollupTimeCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTimeTransactorRaw struct {
	Contract *RollupTimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupTime creates a new instance of RollupTime, bound to a specific deployed contract.
func NewRollupTime(address common.Address, backend bind.ContractBackend) (*RollupTime, error) {
	contract, err := bindRollupTime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupTime{RollupTimeCaller: RollupTimeCaller{contract: contract}, RollupTimeTransactor: RollupTimeTransactor{contract: contract}, RollupTimeFilterer: RollupTimeFilterer{contract: contract}}, nil
}

// NewRollupTimeCaller creates a new read-only instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeCaller(address common.Address, caller bind.ContractCaller) (*RollupTimeCaller, error) {
	contract, err := bindRollupTime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTimeCaller{contract: contract}, nil
}

// NewRollupTimeTransactor creates a new write-only instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTimeTransactor, error) {
	contract, err := bindRollupTime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTimeTransactor{contract: contract}, nil
}

// NewRollupTimeFilterer creates a new log filterer instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupTimeFilterer, error) {
	contract, err := bindRollupTime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupTimeFilterer{contract: contract}, nil
}

// bindRollupTime binds a generic wrapper to an already deployed contract.
func bindRollupTime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTime *RollupTimeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTime.Contract.RollupTimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTime *RollupTimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTime.Contract.RollupTimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTime *RollupTimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTime.Contract.RollupTimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTime *RollupTimeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTime *RollupTimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTime *RollupTimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTime.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d6218032e387455079786fec723dac80ece2fd829cb6194ef40576341993733164736f6c634300050d0032"

// DeployValue deploys a new Ethereum contract, binding an instance of Value to it.
func DeployValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Value, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// Value is an auto generated Go binding around an Ethereum contract.
type Value struct {
	ValueCaller     // Read-only binding to the contract
	ValueTransactor // Write-only binding to the contract
	ValueFilterer   // Log filterer for contract events
}

// ValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueSession struct {
	Contract     *Value            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueCallerSession struct {
	Contract *ValueCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueTransactorSession struct {
	Contract     *ValueTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueRaw struct {
	Contract *Value // Generic contract binding to access the raw methods on
}

// ValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueCallerRaw struct {
	Contract *ValueCaller // Generic read-only contract binding to access the raw methods on
}

// ValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueTransactorRaw struct {
	Contract *ValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValue creates a new instance of Value, bound to a specific deployed contract.
func NewValue(address common.Address, backend bind.ContractBackend) (*Value, error) {
	contract, err := bindValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// NewValueCaller creates a new read-only instance of Value, bound to a specific deployed contract.
func NewValueCaller(address common.Address, caller bind.ContractCaller) (*ValueCaller, error) {
	contract, err := bindValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueCaller{contract: contract}, nil
}

// NewValueTransactor creates a new write-only instance of Value, bound to a specific deployed contract.
func NewValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueTransactor, error) {
	contract, err := bindValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTransactor{contract: contract}, nil
}

// NewValueFilterer creates a new log filterer instance of Value, bound to a specific deployed contract.
func NewValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueFilterer, error) {
	contract, err := bindValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueFilterer{contract: contract}, nil
}

// bindValue binds a generic wrapper to an already deployed contract.
func bindValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.ValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.contract.Transact(opts, method, params...)
}
