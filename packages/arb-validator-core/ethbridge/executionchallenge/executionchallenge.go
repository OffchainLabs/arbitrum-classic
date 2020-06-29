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
	_ = abi.U256
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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158204a4406891b86a3dacfb6ea735b92a605972908aec5c506cc187bb6f01ef6a24064736f6c634300050f0032"

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

// ChallengeUtilsABI is the input ABI used to generate the binding from.
const ChallengeUtilsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_EXECUTION_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_INBOX_TOP_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_MESSAGES_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALID_CHILD_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeUtilsFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeUtilsFuncSigs = map[string]string{
	"95312727": "INVALID_EXECUTION_TYPE()",
	"a697bcac": "INVALID_INBOX_TOP_TYPE()",
	"d7519b46": "INVALID_MESSAGES_TYPE()",
	"2e179be5": "VALID_CHILD_TYPE()",
}

// ChallengeUtilsBin is the compiled bytecode used for deploying new contracts.
var ChallengeUtilsBin = "0x60c9610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060515760003560e01c80632e179be51460565780639531272714606e578063a697bcac146074578063d7519b4614607a575b600080fd5b605c6080565b60408051918252519081900360200190f35b605c6085565b605c608a565b605c608f565b600381565b600281565b600081565b60018156fea265627a7a72315820700623dcb300d07a89dd66c2f312e517c83be08b285278c67140cfaaad930fb064736f6c634300050f0032"

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

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDEXECUTIONTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_EXECUTION_TYPE")
	return *ret0, err
}

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDEXECUTIONTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDEXECUTIONTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDEXECUTIONTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDEXECUTIONTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDINBOXTOPTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_INBOX_TOP_TYPE")
	return *ret0, err
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDINBOXTOPTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDINBOXTOPTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDINBOXTOPTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDINBOXTOPTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDMESSAGESTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_MESSAGES_TYPE")
	return *ret0, err
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDMESSAGESTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDMESSAGESTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDMESSAGESTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDMESSAGESTYPE(&_ChallengeUtils.CallOpts)
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) VALIDCHILDTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "VALID_CHILD_TYPE")
	return *ret0, err
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) VALIDCHILDTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.VALIDCHILDTYPE(&_ChallengeUtils.CallOpts)
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() constant returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) VALIDCHILDTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.VALIDCHILDTYPE(&_ChallengeUtils.CallOpts)
}

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[]"

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820627f6a459ae30c55dc148957baf3d8cfd79ce9663e6fd568623cac07d08f02b564736f6c634300050f0032"

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
var ExecutionChallengeBin = "0x608060405234801561001057600080fd5b50611ff3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806302ad1e4e1461005c57806379a9ad85146100a0578063a72c3ef414610150578063ced5c1bf1461043d578063f430e01114610445575b600080fd5b61009e600480360360a081101561007257600080fd5b506001600160a01b0381358116916020810135821691604082013516906060810135906080013561055f565b005b61009e600480360360808110156100b657600080fd5b81359190810190604081016020820135600160201b8111156100d757600080fd5b8201836020820111156100e957600080fd5b803590602001918460018302840111600160201b8311171561010a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610574565b61009e600480360361016081101561016757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091949392602081019250359050600160201b8111156101b857600080fd5b8201836020820111156101ca57600080fd5b803590602001918460208302840111600160201b831117156101eb57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561023a57600080fd5b82018360208201111561024c57600080fd5b803590602001918460208302840111600160201b8311171561026d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102bc57600080fd5b8201836020820111156102ce57600080fd5b803590602001918460208302840111600160201b831117156102ef57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561033e57600080fd5b82018360208201111561035057600080fd5b803590602001918460208302840111600160201b8311171561037157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103c057600080fd5b8201836020820111156103d257600080fd5b803590602001918460208302840111600160201b831117156103f357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505050903567ffffffffffffffff1691506108729050565b61009e610a47565b61009e60048036036101e081101561045c57600080fd5b60408051608081810183528435946020810135949381013593810192909160e0830191906060840190600490839083908082843760009201919091525091948335946020850135151594604081013594506060810135935060808101359260a08201359267ffffffffffffffff60c0840135169261010081019060e00135600160201b8111156104eb57600080fd5b8201836020820111156104fd57600080fd5b803590602001918460018302840111600160201b8311171561051e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b27945050505050565b61056b85858585610ede565b60065550505050565b60055460ff16600281111561058557fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b815250906106335760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156105f85781810151838201526020016105e0565b50505050905090810190601f1680156106255780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060035461064043610ff9565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b815250906106b35760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b0316331461072f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b815250906107a25760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b506107b283838387600101611000565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b815250906108205760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50600681905561082e611101565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60055460ff16600281111561088357fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906108f45760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060035461090143610ff9565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906109745760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146109f05760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50610a3d6040518061010001604052808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018367ffffffffffffffff1681525061111d565b5050505050505050565b600354610a5343610ff9565b11610aa5576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff166002811115610ab857fe5b1415610af4576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a1610aef6118b1565b610b25565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a1610b256118bc565b565b60055460ff166002811115610b3857fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090610ba95760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50600354610bb643610ff9565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610c295760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610ca55760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50610cb98c8a8d8d8c8c888d8d8d8d6118c4565b600073__$efcad257db8183701794ea8506d55e247c$__63e987d8878e8c8f8f8e8e8e8e8e8e8e8e6040518d63ffffffff1660e01b8152600401808d81526020018c600460200280838360005b83811015610d1e578181015183820152602001610d06565b505050509050018b81526020018a8152602001898152602001881515151581526020018781526020018681526020018581526020018481526020018367ffffffffffffffff1667ffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610dac578181015183820152602001610d94565b50505050905090810190601f168015610dd95780820380516001836020036101000a031916815260200191505b509d505050505050505050505050505060206040518083038186803b158015610e0157600080fd5b505af4158015610e15573d6000803e3d6000fd5b505050506040513d6020811015610e2b57600080fd5b505160408051808201909152600981526827a9a82fa82927a7a360b91b60208201529091508115610e9d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b506040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1610ecf6118bc565b50505050505050505050505050565b600060055460ff166002811115610ef157fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b81525090610f665760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50600080546001600160a01b038681166001600160a01b03199283161790925560018054868416908316178155600280549386169390921692909217905560048290556005805460ff19169091179055610fbe6118fe565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e80290565b600080838160205b885181116110f3578089015193506020818a51036020018161102657fe5b0491505b60008211801561103d5750600286066001145b801561104b57508160020a86115b1561105e5760028604600101955061102a565b600286066110a95783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816110a157fe5b0495506110eb565b82846040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816110e457fe5b0460010195505b602001611008565b505090941495945050505050565b600580546001919060ff191682805b0217905550610b256118fe565b6000600182604001515103905081606001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906111a15760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5081608001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b8152509061121c5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b508160a001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906112975760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b508160c001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b8152509061130f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b50600080805b83811015611363578460c00151818151811061132d57fe5b602002602001015183019250818061135957508460600151818151811061135057fe5b60200260200101515b9150600101611315565b506000611391856040015160008151811061137a57fe5b602002602001015186602001518760000151611910565b90506000611419866040015186815181106113a857fe5b6020026020010151848689608001516000815181106113c357fe5b60200260200101518a608001518a815181106113db57fe5b60200260200101518b60a001516000815181106113f457fe5b60200260200101518c60a001518c8151811061140c57fe5b602002602001015161198a565b905061143261142d8760e0015184846119ed565b611a35565b60608560405190808252806020026020018201604052801561145e578160200160208202803883390190505b50905061150b876040015160018151811061147557fe5b6020026020010151886060015160008151811061148e57fe5b60200260200101518960c001516000815181106114a757fe5b60200260200101518a608001516000815181106114c057fe5b60200260200101518b608001516001815181106114d957fe5b60200260200101518c60a001516000815181106114f257fe5b60200260200101518d60a0015160018151811061140c57fe5b915061155f6115288860e0015167ffffffffffffffff1688611aab565b63ffffffff16611559896040015160008151811061154257fe5b60200260200101518a602001518b60000151611910565b846119ed565b8160008151811061156c57fe5b602090810291909101015260015b868110156116cb578760600151600182038151811061159557fe5b6020026020010151156115ad576115aa611ac9565b88525b611659886040015182600101815181106115c357fe5b6020026020010151896060015183815181106115db57fe5b60200260200101518a60c0015184815181106115f357fe5b60200260200101518b60800151858151811061160b57fe5b60200260200101518c60800151866001018151811061162657fe5b60200260200101518d60a00151878151811061163e57fe5b60200260200101518e60a00151886001018151811061140c57fe5b92506116ac6116768960e0015167ffffffffffffffff1689611aea565b63ffffffff166116a68a60400151848151811061168f57fe5b60200260200101518b602001518c60000151611910565b856119ed565b8282815181106116b857fe5b602090810291909101015260010161157a565b506116d581611afd565b6116dd611b0c565b7f99a5d600237469c778df4f1eb561cbc9eada42285a37b35e762776647ab1ee478760400151886060015189608001518a60a001518b60c001518c60e001516003546040518080602001806020018060200180602001806020018867ffffffffffffffff1667ffffffffffffffff16815260200187815260200186810386528d818151815260200191508051906020019060200280838360005b8381101561178f578181015183820152602001611777565b5050505090500186810385528c818151815260200191508051906020019060200280838360005b838110156117ce5781810151838201526020016117b6565b5050505090500186810384528b818151815260200191508051906020019060200280838360005b8381101561180d5781810151838201526020016117f5565b5050505090500186810383528a818151815260200191508051906020019060200280838360005b8381101561184c578181015183820152602001611834565b50505050905001868103825289818151815260200191508051906020019060200280838360005b8381101561188b578181015183820152602001611873565b505050509050019c5050505050505050505050505060405180910390a150505050505050565b6118b9611b20565b33ff5b6118b9611bc9565b60006118d08a8a611bf0565b905060006118df8d8d84611910565b9050610ecf61142d6001836118f98d8d8d8d8d8d8d61198a565b6119ed565b60045461190a43610ff9565b01600355565b81516020808401516040808601516060968701518251808601999099526fffffffffffffffffffffffffffffffff19608096871b81168a85015293861b841660508a015290851b83169688019690965294831b166070860152818501929092528251808503909101815260a0909301909152815191012090565b6040805160208082019990995296151560f81b8782015260c09590951b6001600160c01b031916604187015260498601939093526069850191909152608984015260a9808401919091528151808403909101815260c99092019052805191012090565b6040805160c09490941b6001600160c01b0319166020808601919091526028850193909352604880850192909252805180850390920182526068909301909252815191012090565b6006548114604051806040016040528060088152602001672124a9afa82922ab60c11b81525090611aa75760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156105f85781810151838201526020016105e0565b5050565b6000818381611ab657fe5b06828481611ac057fe5b04019392505050565b60408051600080825260208201909252611ae4816001611c2a565b91505090565b6000818381611af557fe5b049392505050565b611b0681611c51565b60065550565b600580546002919060ff1916600183611110565b6000546002546001546001600160a01b0392831692636bc3cd229281169116611b47611d8f565b6040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b031681526020018281526020019350505050600060405180830381600087803b158015611baf57600080fd5b505af1158015611bc3573d6000803e3d6000fd5b50505050565b6000546001546002546001600160a01b0392831692636bc3cd229281169116611b47611d8f565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b6000611c34611f63565b611c3e8484611d94565b9050611c4981611db3565b949350505050565b6000815b600181511115611d725760606002825160010181611c6f57fe5b04604051908082528060200260200182016040528015611c99578160200160208202803883390190505b50905060005b8151811015611d6a578251816002026001011015611d3257828160020281518110611cc657fe5b6020026020010151838260020260010181518110611ce057fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611d2157fe5b602002602001018181525050611d62565b828160020281518110611d4157fe5b6020026020010151828281518110611d5557fe5b6020026020010181815250505b600101611c9f565b509050611c55565b80600081518110611d7f57fe5b6020026020010151915050919050565b600290565b611d9c611f63565b6000611da784611e1f565b9050611c498184611edf565b606081015160009060ff16600214611e08576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b81516080830151611e199190611bf0565b92915050565b6000600882511115611e6f576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611eb3578181015183820152602001611e9b565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b611ee7611f63565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611f4c565b611f39611f63565b815260200190600190039081611f315790505b508152600260208201526040019290925250919050565b6040518060a0016040528060008152602001611f7d611f97565b815260606020820181905260006040830181905291015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a723158200627c20d01ee11a5ef43618c4a55ca4ddb6ed6b2f4ce9dbcf54e0ae6470487a864736f6c634300050f0032"

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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582015bd005767d0e6404daebb546d63d966639bacc5d90eb633c96d9c743a320ff464736f6c634300050f0032"

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
var MerkleLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820bc6515a2a572f7a548a9ce0e25f851fe51fa2978807a6f991c940b238731ab6764736f6c634300050f0032"

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
var OneStepProofBin = "0x613e03610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063e987d8871461003a575b600080fd5b61015c60048036036101e081101561005157600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100e757600080fd5b8201836020820111156100f957600080fd5b8035906020019184600183028401116401000000008311171561011b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061016e945050505050565b60408051918252519081900360200190f35b60006101d96040518061016001604052808f81526020018e81526020016101958e8e6101ea565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610270565b9d9c50505050505050505050505050565b6101f2613be6565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610257565b610244613be6565b81526020019060019003908161023c5790505b5081526002602082015260400183905290505b92915050565b6000806000806060610280613c1a565b610288613c1a565b61029188611362565b93995092965090945092509050600160006102ab886116e1565b67ffffffffffffffff168a610120015167ffffffffffffffff161461030e576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b89608001518015610322575060ff88166072145b8061033e5750896080015115801561033e575060ff8816607214155b61038f576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60ff8816600114156103d5576103ce83866000815181106103ac57fe5b6020026020010151876001815181106103c157fe5b6020026020010151611b68565b91506111b6565b60ff881660021415610414576103ce83866000815181106103f257fe5b60200260200101518760018151811061040757fe5b6020026020010151611bb8565b60ff881660031415610453576103ce838660008151811061043157fe5b60200260200101518760018151811061044657fe5b6020026020010151611bf9565b60ff881660041415610492576103ce838660008151811061047057fe5b60200260200101518760018151811061048557fe5b6020026020010151611c3a565b60ff8816600514156104d1576103ce83866000815181106104af57fe5b6020026020010151876001815181106104c457fe5b6020026020010151611c8b565b60ff881660061415610510576103ce83866000815181106104ee57fe5b60200260200101518760018151811061050357fe5b6020026020010151611cdc565b60ff88166007141561054f576103ce838660008151811061052d57fe5b60200260200101518760018151811061054257fe5b6020026020010151611d2d565b60ff8816600814156105a3576103ce838660008151811061056c57fe5b60200260200101518760018151811061058157fe5b60200260200101518860028151811061059657fe5b6020026020010151611d7e565b60ff8816600914156105f7576103ce83866000815181106105c057fe5b6020026020010151876001815181106105d557fe5b6020026020010151886002815181106105ea57fe5b6020026020010151611de8565b60ff8816600a1415610636576103ce838660008151811061061457fe5b60200260200101518760018151811061062957fe5b6020026020010151611e41565b60ff881660101415610675576103ce838660008151811061065357fe5b60200260200101518760018151811061066857fe5b6020026020010151611e82565b60ff8816601114156106b4576103ce838660008151811061069257fe5b6020026020010151876001815181106106a757fe5b6020026020010151611ec3565b60ff8816601214156106f3576103ce83866000815181106106d157fe5b6020026020010151876001815181106106e657fe5b6020026020010151611f04565b60ff881660131415610732576103ce838660008151811061071057fe5b60200260200101518760018151811061072557fe5b6020026020010151611f45565b60ff881660141415610771576103ce838660008151811061074f57fe5b60200260200101518760018151811061076457fe5b6020026020010151611f86565b60ff88166015141561079b576103ce838660008151811061078e57fe5b6020026020010151611fb0565b60ff8816601614156107da576103ce83866000815181106107b857fe5b6020026020010151876001815181106107cd57fe5b6020026020010151611ff5565b60ff881660171415610819576103ce83866000815181106107f757fe5b60200260200101518760018151811061080c57fe5b6020026020010151612036565b60ff881660181415610858576103ce838660008151811061083657fe5b60200260200101518760018151811061084b57fe5b6020026020010151612077565b60ff881660191415610882576103ce838660008151811061087557fe5b60200260200101516120b8565b60ff8816601a14156108c1576103ce838660008151811061089f57fe5b6020026020010151876001815181106108b457fe5b60200260200101516120ee565b60ff8816601b1415610900576103ce83866000815181106108de57fe5b6020026020010151876001815181106108f357fe5b602002602001015161212f565b60ff88166020141561092a576103ce838660008151811061091d57fe5b6020026020010151612170565b60ff881660211415610954576103ce838660008151811061094757fe5b602002602001015161218b565b60ff881660221415610993576103ce838660008151811061097157fe5b60200260200101518760018151811061098657fe5b60200260200101516121a6565b60ff8816603014156109bd576103ce83866000815181106109b057fe5b602002602001015161220c565b60ff8816603114156109d2576103ce83612214565b60ff8816603214156109e7576103ce83612235565b60ff881660331415610a11576103ce8386600081518110610a0457fe5b602002602001015161224e565b60ff881660341415610a3b576103ce8386600081518110610a2e57fe5b602002602001015161225a565b60ff881660351415610a7a576103ce8386600081518110610a5857fe5b602002602001015187600181518110610a6d57fe5b6020026020010151612285565b60ff881660361415610a8f576103ce836122cd565b60ff881660371415610aa9576103ce8385600001516122f7565b60ff881660381415610ad3576103ce8386600081518110610ac657fe5b6020026020010151612305565b60ff881660391415610b5f57610ae7613be6565b610af68b610140015188612317565b9199509750905087610b395760405162461bcd60e51b8152600401808060200182810382526021815260200180613dae6021913960400191505060405180910390fd5b610b49858263ffffffff61245516565b610b59848263ffffffff61246f16565b506111b6565b60ff8816603a1415610b74576103ce83612489565b60ff8816603b1415610b85576111b6565b60ff8816603c1415610b9a576103ce836124a6565b60ff8816603d1415610bc4576103ce8386600081518110610bb757fe5b60200260200101516124b8565b60ff881660401415610bee576103ce8386600081518110610be157fe5b60200260200101516124e6565b60ff881660411415610c2d576103ce8386600081518110610c0b57fe5b602002602001015187600181518110610c2057fe5b6020026020010151612508565b60ff881660421415610c81576103ce8386600081518110610c4a57fe5b602002602001015187600181518110610c5f57fe5b602002602001015188600281518110610c7457fe5b602002602001015161253a565b60ff881660431415610cc0576103ce8386600081518110610c9e57fe5b602002602001015187600181518110610cb357fe5b602002602001015161257c565b60ff881660441415610d14576103ce8386600081518110610cdd57fe5b602002602001015187600181518110610cf257fe5b602002602001015188600281518110610d0757fe5b602002602001015161258e565b60ff881660501415610d53576103ce8386600081518110610d3157fe5b602002602001015187600181518110610d4657fe5b60200260200101516125b0565b60ff881660511415610da7576103ce8386600081518110610d7057fe5b602002602001015187600181518110610d8557fe5b602002602001015188600281518110610d9a57fe5b6020026020010151612626565b60ff881660521415610dd1576103ce8386600081518110610dc457fe5b60200260200101516126b3565b60ff881660601415610de6576103ce836126e6565b60ff881660611415610ee457610e108386600081518110610e0357fe5b60200260200101516126ec565b90925090508115610edb578961010001518a60e001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610e905760405162461bcd60e51b8152600401808060200182810382526025815260200180613d626025913960400191505060405180910390fd5b8960c001518a60a0015114610ed65760405162461bcd60e51b8152600401808060200182810382526027815260200180613d876027913960400191505060405180910390fd5b610edf565b5060005b6111b6565b60ff88166070141561102457610f0e8386600081518110610f0157fe5b6020026020010151612706565b90925090508115610edb5780610f69578960c001518a60a0015114610f645760405162461bcd60e51b8152600401808060200182810382526038815260200180613d2a6038913960400191505060405180910390fd5b610ed6565b8960c001518a60a001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610fdd5760405162461bcd60e51b8152600401808060200182810382526029815260200180613cba6029913960400191505060405180910390fd5b8961010001518a60e0015114610ed65760405162461bcd60e51b8152600401808060200182810382526026815260200180613ce36026913960400191505060405180910390fd5b60ff8816607114156111395760408051600480825260a08201909252606091816020015b611050613be6565b81526020019060019003908161104857505060208c01519091506110849060005b60200201516001600160801b0316612736565b8160008151811061109157fe5b60200260200101819052506110b08b6020015160016004811061107157fe5b816001815181106110bd57fe5b60200260200101819052506110dc8b6020015160026004811061107157fe5b816002815181106110e957fe5b60200260200101819052506111088b6020015160036004811061107157fe5b8160038151811061111557fe5b6020026020010181905250610b5961112c826127bb565b859063ffffffff61246f16565b60ff881660721415611187576103ce838660008151811061115657fe5b60200260200101518c604001518d6020015160006004811061117457fe5b60200201516001600160801b03166128aa565b60ff88166073141561119c57600091506111b6565b60ff8816607414156111b157610edf83612941565b600091505b80611248578960c001518a60a00151146112015760405162461bcd60e51b8152600401808060200182810382526027815260200180613d876027913960400191505060405180910390fd5b8961010001518a60e00151146112485760405162461bcd60e51b8152600401808060200182810382526026815260200180613ce36026913960400191505060405180910390fd5b816112a95760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a084015114156112a15761129c8361294b565b6112a9565b60a083015183525b6112b284612955565b8a51146112f05760405162461bcd60e51b8152600401808060200182810382526022815260200180613c986022913960400191505060405180910390fd5b6112f983612955565b8a6060015114611350576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b6000606061136e613c1a565b611376613c1a565b60008080611382613c1a565b61138b81612a0d565b61139a89610140015184612a17565b9094509092509050816113f4576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b6113fc613c1a565b61140582612b28565b905060008a6101400151858151811061141a57fe5b602001015160f81c60f81b60f81c905060008b6101400151866001018151811061144057fe5b016020015160f81c9050600061145582612b86565b905060608160405190808252806020026020018201604052801561149357816020015b611480613be6565b8152602001906001900390816114785790505b5090506002880197508360ff16600014806114b157508360ff166001145b611502576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff84166115275761152061151b848860000151612ba0565b612bda565b86526115ea565b61152f613be6565b61153e8f61014001518a612317565b909a50909850905087611598576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b82156115bc5780826000815181106115ac57fe5b60200260200101819052506115cc565b6115cc868263ffffffff61246f16565b6115e661151b8589600001516115e185612bda565b612cd2565b8752505b60ff84165b8281101561167d576116068f61014001518a612317565b845185908590811061161457fe5b60209081029190910101529950975087611675576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b6001016115ef565b8151156116ca575060005b8460ff168251038110156116ca576116c28282600185510303815181106116ab57fe5b60200260200101518861246f90919063ffffffff16565b600101611688565b50919d919c50939a50919850939650945050505050565b600060ff8216600114156116f75750600361135d565b60ff82166002141561170b5750600361135d565b60ff82166003141561171f5750600361135d565b60ff8216600414156117335750600461135d565b60ff8216600514156117475750600761135d565b60ff82166006141561175b5750600461135d565b60ff82166007141561176f5750600761135d565b60ff8216600814156117835750600461135d565b60ff8216600914156117975750600461135d565b60ff8216600a14156117ab5750601961135d565b60ff8216601014156117bf5750600261135d565b60ff8216601114156117d35750600261135d565b60ff8216601214156117e75750600261135d565b60ff8216601314156117fb5750600261135d565b60ff82166014141561180f5750600261135d565b60ff8216601514156118235750600161135d565b60ff8216601614156118375750600261135d565b60ff82166017141561184b5750600261135d565b60ff82166018141561185f5750600261135d565b60ff8216601914156118735750600161135d565b60ff8216601a14156118875750600461135d565b60ff8216601b141561189b5750600761135d565b60ff8216602014156118af5750600761135d565b60ff8216602114156118c35750600361135d565b60ff8216602214156118d75750600861135d565b60ff8216603014156118eb5750600161135d565b60ff8216603114156118ff5750600161135d565b60ff8216603214156119135750600161135d565b60ff8216603314156119275750600261135d565b60ff82166034141561193b5750600461135d565b60ff82166035141561194f5750600461135d565b60ff8216603614156119635750600261135d565b60ff8216603714156119775750600161135d565b60ff82166038141561198b5750600161135d565b60ff82166039141561199f5750600161135d565b60ff8216603a14156119b35750600261135d565b60ff8216603b14156119c75750600161135d565b60ff8216603c14156119db5750600161135d565b60ff8216603d14156119ef5750600161135d565b60ff821660401415611a035750600161135d565b60ff821660411415611a175750600161135d565b60ff821660421415611a2b5750600161135d565b60ff821660431415611a3f5750600161135d565b60ff821660441415611a535750600161135d565b60ff821660501415611a675750600261135d565b60ff821660511415611a7b5750602861135d565b60ff821660521415611a8f5750600261135d565b60ff821660601415611aa35750606461135d565b60ff821660611415611ab75750606461135d565b60ff821660701415611acb5750606461135d565b60ff821660711415611adf5750602861135d565b60ff821660721415611af35750602861135d565b60ff821660731415611b075750600561135d565b60ff821660741415611b1b5750600a61135d565b6040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206f70636f64653a206f70476173436f737428290000000000604482015290519081900360640190fd5b6000611b7383612d09565b1580611b855750611b8382612d09565b155b15611b9257506000611bb1565b82518251808201611ba9878263ffffffff612d1416565b600193505050505b9392505050565b6000611bc383612d09565b1580611bd55750611bd382612d09565b155b15611be257506000611bb1565b82518251808202611ba9878263ffffffff612d1416565b6000611c0483612d09565b1580611c165750611c1482612d09565b155b15611c2357506000611bb1565b82518251808203611ba9878263ffffffff612d1416565b6000611c4583612d09565b1580611c575750611c5582612d09565b155b15611c6457506000611bb1565b8251825180611c7857600092505050611bb1565b808204611ba9878263ffffffff612d1416565b6000611c9683612d09565b1580611ca85750611ca682612d09565b155b15611cb557506000611bb1565b8251825180611cc957600092505050611bb1565b808205611ba9878263ffffffff612d1416565b6000611ce783612d09565b1580611cf95750611cf782612d09565b155b15611d0657506000611bb1565b8251825180611d1a57600092505050611bb1565b808206611ba9878263ffffffff612d1416565b6000611d3883612d09565b1580611d4a5750611d4882612d09565b155b15611d5757506000611bb1565b8251825180611d6b57600092505050611bb1565b808207611ba9878263ffffffff612d1416565b6000611d8984612d09565b1580611d9b5750611d9983612d09565b155b15611da857506000611de0565b83518351835180611dbf5760009350505050611de0565b6000818385089050611dd7898263ffffffff612d1416565b60019450505050505b949350505050565b6000611df384612d09565b1580611e055750611e0383612d09565b155b15611e1257506000611de0565b83518351835180611e295760009350505050611de0565b6000818385099050611dd7898263ffffffff612d1416565b6000611e4c83612d09565b1580611e5e5750611e5c82612d09565b155b15611e6b57506000611bb1565b8251825180820a611ba9878263ffffffff612d1416565b6000611e8d83612d09565b1580611e9f5750611e9d82612d09565b155b15611eac57506000611bb1565b82518251808210611ba9878263ffffffff612d1416565b6000611ece83612d09565b1580611ee05750611ede82612d09565b155b15611eed57506000611bb1565b82518251808211611ba9878263ffffffff612d1416565b6000611f0f83612d09565b1580611f215750611f1f82612d09565b155b15611f2e57506000611bb1565b82518251808212611ba9878263ffffffff612d1416565b6000611f5083612d09565b1580611f625750611f6082612d09565b155b15611f6f57506000611bb1565b82518251808213611ba9878263ffffffff612d1416565b6000611fa661112c611f9784612bda565b611fa086612bda565b14612d2a565b5060019392505050565b6000611fbb82612d09565b611fd557611fd083600063ffffffff612d1416565b611fec565b81518015611fe9858263ffffffff612d1416565b50505b50600192915050565b600061200083612d09565b1580612012575061201082612d09565b155b1561201f57506000611bb1565b82518251808216611ba9878263ffffffff612d1416565b600061204183612d09565b1580612053575061205182612d09565b155b1561206057506000611bb1565b82518251808217611ba9878263ffffffff612d1416565b600061208283612d09565b1580612094575061209282612d09565b155b156120a157506000611bb1565b82518251808218611ba9878263ffffffff612d1416565b60006120c382612d09565b6120cf5750600061026a565b815180196120e3858263ffffffff612d1416565b506001949350505050565b60006120f983612d09565b158061210b575061210982612d09565b155b1561211857506000611bb1565b8251825181811a611ba9878263ffffffff612d1416565b600061213a83612d09565b158061214c575061214a82612d09565b155b1561215957506000611bb1565b8251825181810b611ba9878263ffffffff612d1416565b6000611fec61217e83612bda565b849063ffffffff612d1416565b6000611fec61219983612d4c565b849063ffffffff61246f16565b60006121b183612d09565b15806121c357506121c182612d09565b155b156121d057506000611bb1565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120611ba9878263ffffffff612d1416565b600192915050565b600061222d82608001518361246f90919063ffffffff16565b506001919050565b600061222d82606001518361246f90919063ffffffff16565b60609190910152600190565b600061226582612dd5565b6122715750600061026a565b61227a82612bda565b835250600192915050565b600061229083612dd5565b61229c57506000611bb1565b6122a582612d09565b6122b157506000611bb1565b815115611fa6576122c183612bda565b84525060019392505050565b600061222d6122ea6122dd612de2565b611fa08560200151612bda565b839063ffffffff61246f16565b6000611fec61219983612e03565b6000611fec838363ffffffff61245516565b600080612322613be6565b84518410612342576000846123376000612736565b92509250925061244e565b600080859050600087828151811061235657fe5b016020015160019092019160f81c90506000612370613c70565b60ff83166123a4576123828a85612e88565b91965094509150848461239484612736565b975097509750505050505061244e565b60ff8316600114156123cc576123ba8a85612edb565b9196509450905084846123948361303b565b60ff8316600214156123e2576123948a856130a2565b600360ff8416108015906123f95750600c60ff8416105b156124345760021983016060612410828d88613147565b919850965090508686612422836127bb565b9950995099505050505050505061244e565b6000806124416000612736565b9199509750955050505050505b9250925092565b612463826040015182613205565b82604001819052505050565b61247d826020015182613205565b82602001819052505050565b600061222d6122ea612499612de2565b611fa08560400151612bda565b600061222d6122ea8360a00151612e03565b60006124c382612dd5565b6124cf5750600061026a565b6124d882612bda565b60a084015250600192915050565b60006124f8838363ffffffff61246f16565b611fec838363ffffffff61246f16565b600061251a848363ffffffff61246f16565b61252a848463ffffffff61246f16565b611fa6848363ffffffff61246f16565b600061254c858363ffffffff61246f16565b61255c858463ffffffff61246f16565b61256c858563ffffffff61246f16565b6120e3858363ffffffff61246f16565b600061252a848463ffffffff61246f16565b60006125a0858563ffffffff61246f16565b61256c858463ffffffff61246f16565b60006125bb83612d09565b15806125cd57506125cb82613283565b155b156125da57506000611bb1565b6125e382613292565b60ff168360000151106125f857506000611bb1565b611fa6826040015184600001518151811061260f57fe5b60200260200101518561246f90919063ffffffff16565b600061263183613283565b1580612643575061264184612d09565b155b1561265057506000611de0565b61265983613292565b60ff1684600001511061266e57506000611de0565b60408301518451815184918391811061268357fe5b60200260200101819052506126a761269a826127bb565b879063ffffffff61246f16565b50600195945050505050565b60006126be82613283565b6126ca5750600061026a565b611fec6126d683613292565b849060ff1663ffffffff612d1416565b50600190565b60008060016126fa84612bda565b915091505b9250929050565b60008061271083608001511161272a57600161272184612bda565b915091506126ff565b506001905060006126ff565b61273e613be6565b6040805160a08101825283815281516080810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916127a3565b612790613be6565b8152602001906001900390816127885790505b50815260006020820152600160409091015292915050565b6127c3613be6565b6127cd82516132a1565b61281e576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156128555783818151811061283857fe5b602002602001015160800151820191508080600101915050612823565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b60006128b584612d09565b6128c157506000611de0565b8351821115806128e057506128d4612de2565b6128dd84612bda565b14155b612931576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b6120e3858463ffffffff61246f16565b600260c090910152565b600160c090910152565b600060028260c00151141561296c5750600061135d565b60018260c0015114156129815750600161135d565b8151602083015161299190612bda565b61299e8460400151612bda565b6129ab8560600151612bda565b6129b88660800151612bda565b8660a0015160405160200180878152602001868152602001858152602001848152602001838152602001828152602001965050505050505060405160208183030381529060405280519060200120905061135d565b600060c090910152565b600080612a22613c1a565b612a2a613c1a565b600060c08201819052612a3d87876132a8565b84529650905080612a57575060009350849250905061244e565b612a6187876130a2565b60208501529650905080612a7e575060009350849250905061244e565b612a8887876130a2565b60408501529650905080612aa5575060009350849250905061244e565b612aaf8787612317565b60608501529650905080612acc575060009350849250905061244e565b612ad68787612317565b60808501529650905080612af3575060009350849250905061244e565b612afd87876132a8565b60a08501529650905080612b1a575060009350849250905061244e565b506001969495509392505050565b612b30613c1a565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b6000806000612b978460ff166132fc565b50949350505050565b612ba8613be6565b611bb160405180608001604052808560ff1681526020018481526020016000151581526020016000801b81525061303b565b606081015160009060ff16612bfb578151612bf4906137b3565b905061135d565b606082015160ff1660011415612c2e576020808301518051604082015160608301519290930151612bf4939192906137d7565b606082015160ff1660021415612c4757612bf48261387f565b600360ff16826060015160ff1610158015612c6b57506060820151600c60ff909116105b15612c7957612bf4826138e5565b606082015160ff1660641415612c915750805161135d565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b612cda613be6565b611de060405180608001604052808660ff1681526020018581526020016001151581526020018481525061303b565b6060015160ff161590565b61247d8260200151612d2583612736565b613205565b612d32613be6565b8115612d4257612bf46001612736565b612bf46000612736565b612d54613be6565b816060015160ff1660021415612d9b5760405162461bcd60e51b8152600401808060200182810382526021815260200180613d096021913960400191505060405180910390fd5b606082015160ff16612db157612bf46000612736565b816060015160ff1660011415612dcb57612bf46001612736565b612bf46003612736565b6060015160ff1660011490565b60408051600080825260208201909252612dfd816001613903565b91505090565b612e0b613be6565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612e70565b612e5d613be6565b815260200190600190039081612e555790505b50815260646020820152600160409091015292915050565b6000806000808551905084811080612ea257506020858203105b15612eb757506000925083915082905061244e565b600160208601612ecd888863ffffffff61392216565b935093509350509250925092565b600080612ee6613c70565b60008490506000868281518110612ef957fe5b602001015160f81c60f81b60f81c905081806001019250506000878381518110612f1f57fe5b016020015160019384019360f89190911c915060009060ff84161415612fa5576000612f49613be6565b612f538b87612317565b909750909250905081612f975750506040805160808101825260008082526020820181905291810182905260608101829052909750889650945061244e9350505050565b612fa081612bda565b925050505b6000612fb78a8663ffffffff61392216565b90506020850194508360ff1660011415613003576040805160808101825260ff9094168452602084019190915260019083018190526060830191909152955091935090915061244e9050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b613043613be6565b6040805160a081018252600080825260208083018690528351828152908101845291928301919061308a565b613077613be6565b81526020019060019003908161306f5790505b50815260016020820181905260409091015292915050565b6000806130ad613be6565b6130b5613be6565b85516000908190878110806130cc57506040888203105b156130e457600088859650965096505050505061244e565b60006130f68a8a63ffffffff61392216565b90506020890198506131088a8a612e88565b909a509450925082156131335761311f81856101ea565b60019850899750955061244e945050505050565b60008986975097509750505050505061244e565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561319257816020015b61317f613be6565b8152602001906001900390816131775790505b50905060005b8960ff168160ff1610156131ef576131b08985612317565b8451859060ff86169081106131c157fe5b602090810291909101015294509250826131e757506000955086945092506131fc915050565b600101613198565b5060019550919350909150505b93509350939050565b61320d613be6565b6040805160028082526060828101909352816020015b61322b613be6565b815260200190600190039081613223579050509050828160008151811061324e57fe5b6020026020010181905250838160018151811061326757fe5b6020026020010181905250611de061327e826127bb565b61393e565b600061026a82606001516139b4565b600061026a82606001516139d2565b6008101590565b600080600080600086519050858110806132c457506020868203105b156132d8575060009350849250905061244e565b6132e8878763ffffffff61392216565b60019550602087019450925061244e915050565b600080600183141561331457506002905060016137ae565b600283141561332957506002905060016137ae565b600383141561333e57506002905060016137ae565b600483141561335357506002905060016137ae565b600583141561336857506002905060016137ae565b600683141561337d57506002905060016137ae565b600783141561339257506002905060016137ae565b60088314156133a757506003905060016137ae565b60098314156133bc57506003905060016137ae565b600a8314156133d157506002905060016137ae565b60108314156133e657506002905060016137ae565b60118314156133fb57506002905060016137ae565b601283141561341057506002905060016137ae565b601383141561342557506002905060016137ae565b601483141561343a57506002905060016137ae565b601583141561344e575060019050806137ae565b601683141561346357506002905060016137ae565b601783141561347857506002905060016137ae565b601883141561348d57506002905060016137ae565b60198314156134a1575060019050806137ae565b601a8314156134b657506002905060016137ae565b601b8314156134cb57506002905060016137ae565b60208314156134df575060019050806137ae565b60218314156134f3575060019050806137ae565b602283141561350857506002905060016137ae565b603083141561351d57506001905060006137ae565b603183141561353257506000905060016137ae565b603283141561354757506000905060016137ae565b603383141561355c57506001905060006137ae565b603483141561357157506001905060006137ae565b603583141561358657506002905060006137ae565b603683141561359b57506000905060016137ae565b60378314156135b057506000905060016137ae565b60388314156135c557506001905060006137ae565b60398314156135da57506000905060016137ae565b603a8314156135ef57506000905060016137ae565b603b831415613603575060009050806137ae565b603c83141561361857506000905060016137ae565b603d83141561362d57506001905060006137ae565b604083141561364257506001905060026137ae565b604183141561365757506002905060036137ae565b604283141561366c57506003905060046137ae565b6043831415613680575060029050806137ae565b6044831415613694575060039050806137ae565b60508314156136a957506002905060016137ae565b60518314156136be57506003905060016137ae565b60528314156136d2575060019050806137ae565b60608314156136e6575060009050806137ae565b60618314156136fb57506001905060006137ae565b607083141561371057506001905060006137ae565b607183141561372557506000905060016137ae565b6072831415613739575060019050806137ae565b607383141561374d575060009050806137ae565b6074831415613761575060009050806137ae565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613831575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611de0565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff166002146138d4576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161026a91906139f5565b60006138ef613be6565b6138f88361393e565b9050611bb18161387f565b600061390d613be6565b6139178484613a2f565b9050611de08161387f565b6000816020018351101561393557600080fd5b50016020015190565b613946613be6565b61394f82613283565b613995576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b60606139a48360400151613a4e565b9050611bb1818460800151613a2f565b6000600c60ff831610801561026a575050600360ff91909116101590565b60006139dd826139b4565b156139ed5750600219810161135d565b50600161135d565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b613a37613be6565b6000613a4284613b26565b9050611de081846101ea565b6060600882511115613a9e576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613acb578160200160208202803883390190505b50805190915060005b81811015613b1d576000613afa868381518110613aed57fe5b6020026020010151612bda565b905080848381518110613b0957fe5b602090810291909101015250600101613ad4565b50909392505050565b6000600882511115613b76576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613bba578181015183820152602001613ba2565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001613c00613c70565b815260606020820181905260006040830181905291015290565b6040805160e081019091526000815260208101613c35613be6565b8152602001613c42613be6565b8152602001613c4f613be6565b8152602001613c5c613be6565b815260006020820181905260409091015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820b75a7490e99c2b03ef94918e047e6d8cb6bc7c78ed27ecf11299d4a1b1feb56064736f6c634300050f0032"

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
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
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
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(beforeHash [32]byte, timeBounds [4]*big.Int, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBounds, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0xe987d887.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(beforeHash [32]byte, timeBounds [4]*big.Int, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBounds, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b9dfda0e8967cce081d0a445c1a4924f7d5665789e802649e31b46830564633564736f6c634300050f0032"

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
var RollupTimeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820c104b2abeaa99278b05e5ee999cf1befae1e76998c6df565b4b0a6270c5d15e964736f6c634300050f0032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208ee905886c683022f063711a65435844571ece73993149dfb05f636dca5a99c464736f6c634300050f0032"

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
