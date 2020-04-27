// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package messageschallenge

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

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158203e9c9fccd846d5ea361aeb66070d924a8678b51a168545dc7c1f102219dbb3f264736f6c634300050d0032"

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
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// MessagesChallengeABI is the input ABI used to generate the binding from.
const MessagesChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"segmentHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"segmentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Continued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_chainHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_segmentHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_chainLength\",\"type\":\"uint256\"}],\"name\":\"bisect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_segmentToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"chooseSegment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_lowerHashA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lowerHashB\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_messageNum\",\"type\":\"uint256\"}],\"name\":\"oneStepProofContractTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_lowerHashA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lowerHashB\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_messageNum\",\"type\":\"uint256\"}],\"name\":\"oneStepProofERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_lowerHashA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lowerHashB\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_messageNum\",\"type\":\"uint256\"}],\"name\":\"oneStepProofERC721Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_lowerHashA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lowerHashB\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_messageNum\",\"type\":\"uint256\"}],\"name\":\"oneStepProofEthMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"lowerHashA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lowerHashB\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"seqNumbers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"dataLengths\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"blockAndTimestamp\",\"type\":\"uint256[2]\"}],\"name\":\"oneStepProofTransactionBatchMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_lowerHashA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lowerHashB\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"oneStepProofTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MessagesChallengeFuncSigs maps the 4-byte function signature to its string representation.
var MessagesChallengeFuncSigs = map[string]string{
	"500439e5": "bisect(bytes32[],bytes32[],uint256)",
	"79a9ad85": "chooseSegment(uint256,bytes,bytes32,bytes32)",
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
	"225b3d09": "oneStepProofContractTransactionMessage(bytes32,bytes32,address,address,uint256,bytes,uint256,uint256,uint256)",
	"0c31dd2d": "oneStepProofERC20Message(bytes32,bytes32,address,address,address,uint256,uint256,uint256,uint256)",
	"3f08600e": "oneStepProofERC721Message(bytes32,bytes32,address,address,address,uint256,uint256,uint256,uint256)",
	"3261ee97": "oneStepProofEthMessage(bytes32,bytes32,address,address,uint256,uint256,uint256,uint256)",
	"4b7bea76": "oneStepProofTransactionBatchMessage(bytes32,bytes32,address,address[],uint256[],uint256[],uint32[],bytes,bytes,uint256[2])",
	"d2c56f27": "oneStepProofTransactionMessage(bytes32,bytes32,address,address,address,uint256,uint256,bytes,uint256,uint256)",
	"ced5c1bf": "timeoutChallenge()",
}

// MessagesChallengeBin is the compiled bytecode used for deploying new contracts.
var MessagesChallengeBin = "0x608060405234801561001057600080fd5b50613bbd806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80634b7bea76116100665780634b7bea76146102d0578063500439e51461064757806379a9ad851461076c578063ced5c1bf1461081c578063d2c56f27146108245761009e565b806302ad1e4e146100a35780630c31dd2d146100e7578063225b3d09146101435780633261ee971461021f5780633f08600e14610274575b600080fd5b6100e5600480360360a08110156100b957600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060800135610908565b005b6100e560048036036101208110156100fe57600080fd5b508035906020810135906001600160a01b03604082013581169160608101358216916080820135169060a08101359060c08101359060e081013590610100013561091d565b6100e5600480360361012081101561015a57600080fd5b8135916020810135916001600160a01b0360408301358116926060810135909116916080820135919081019060c0810160a0820135600160201b8111156101a057600080fd5b8201836020820111156101b257600080fd5b803590602001918460018302840111600160201b831117156101d357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060208101359060400135610b18565b6100e5600480360361010081101561023657600080fd5b508035906020810135906001600160a01b03604082013581169160608101359091169060808101359060a08101359060c08101359060e00135610cbb565b6100e5600480360361012081101561028b57600080fd5b508035906020810135906001600160a01b03604082013581169160608101358216916080820135169060a08101359060c08101359060e0810135906101000135610e76565b6100e560048036036101608110156102e757600080fd5b8135916020810135916001600160a01b036040830135169190810190608081016060820135600160201b81111561031d57600080fd5b82018360208201111561032f57600080fd5b803590602001918460208302840111600160201b8311171561035057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561039f57600080fd5b8201836020820111156103b157600080fd5b803590602001918460208302840111600160201b831117156103d257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561042157600080fd5b82018360208201111561043357600080fd5b803590602001918460208302840111600160201b8311171561045457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104a357600080fd5b8201836020820111156104b557600080fd5b803590602001918460208302840111600160201b831117156104d657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561052557600080fd5b82018360208201111561053757600080fd5b803590602001918460018302840111600160201b8311171561055857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156105aa57600080fd5b8201836020820111156105bc57600080fd5b803590602001918460018302840111600160201b831117156105dd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051808201825293969594818101949350915060029083908390808284376000920191909152509194506110199350505050565b6100e56004803603606081101561065d57600080fd5b810190602081018135600160201b81111561067757600080fd5b82018360208201111561068957600080fd5b803590602001918460208302840111600160201b831117156106aa57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106f957600080fd5b82018360208201111561070b57600080fd5b803590602001918460208302840111600160201b8311171561072c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250611204915050565b6100e56004803603608081101561078257600080fd5b81359190810190604081016020820135600160201b8111156107a357600080fd5b8201836020820111156107b557600080fd5b803590602001918460018302840111600160201b831117156107d657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561167e565b6100e561193f565b6100e5600480360361014081101561083b57600080fd5b8135916020810135916001600160a01b036040830135811692606081013582169260808201359092169160a08201359160c081013591810190610100810160e0820135600160201b81111561088f57600080fd5b8201836020820111156108a157600080fd5b803590602001918460018302840111600160201b831117156108c257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135611a1f565b61091485858585611bd9565b60065550505050565b60055460ff16600281111561092e57fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906109dc5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109a1578181015183820152602001610989565b50505050905090810190601f1680156109ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506003546109e943611cf4565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610a5c5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610ad85760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b506000610aea88888888888888611cff565b90506000610afd89898989898989611d1e565b9050610b0b8b8b8484611d31565b5050505050505050505050565b60055460ff166002811115610b2957fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090610b9a5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b50600354610ba743611cf4565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610c1a5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610c965760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b506000610ca888888888888888611d62565b90506000610afd89898989898989611e41565b60055460ff166002811115610ccc57fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090610d3d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b50600354610d4a43611cf4565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610dbd5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610e395760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b506000610e4a878787878787611fef565b90506000610e5c888888888888612059565b9050610e6a8a8a8484611d31565b50505050505050505050565b60055460ff166002811115610e8757fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090610ef85760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b50600354610f0543611cf4565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610f785760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610ff45760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b50600061100688888888888888612257565b90506000610afd8989898989898961226a565b60055460ff16600281111561102a57fe5b600114604051806040016040528060098152602001684249535f535441544560b81b8152509061109b5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b506003546110a843611cf4565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b8152509061111b5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146111975760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060006111b889898989898989898960200201518a6001602002015161227d565b905060006111cd8b8b8b8b8b8b8b8b8b6125bb565b90506111ee6111e98d6111e08f86612770565b8e85600161279c565b6127e2565b6111f6612858565b505050505050505050505050565b60055460ff16600281111561121557fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906112865760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060035461129343611cf4565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906113065760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146113825760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b508251825160408051808201909152600d81526c2429afa124a9afa4a7282622a760991b60208201526000198301929091146113ff5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5061145e6111e98560008151811061141357fe5b602002602001015186848151811061142757fe5b60200260200101518660008151811061143c57fe5b602002602001015187868151811061145057fe5b60200260200101518761279c565b60608160405190808252806020026020018201604052801561148a578160200160208202803883390190505b5090506114f38560008151811061149d57fe5b6020026020010151866001815181106114b257fe5b6020026020010151866000815181106114c757fe5b6020026020010151876001815181106114dc57fe5b60200260200101516114ee8888612889565b61279c565b8160008151811061150057fe5b602090810291909101015260015b828110156115985761157986828151811061152557fe5b602002602001015187836001018151811061153c57fe5b602002602001015187848151811061155057fe5b602002602001015188856001018151811061156757fe5b60200260200101516114ee89896128a7565b82828151811061158557fe5b602090810291909101015260010161150e565b506115a2816128ba565b6115aa6128c9565b7f500c4a1bbd12a65d684bde95626a41abdd6a8c5d30f84c5c9b81e5bdb0cc0bd3858585600354604051808060200180602001858152602001848152602001838103835287818151815260200191508051906020019060200280838360005b83811015611621578181015183820152602001611609565b50505050905001838103825286818151815260200191508051906020019060200280838360005b83811015611660578181015183820152602001611648565b50505050905001965050505050505060405180910390a15050505050565b60055460ff16600281111561168f57fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b815250906117005760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060035461170d43611cf4565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b815250906117805760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b031633146117fc5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b8152509061186f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5061187f838383876001016128e6565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b815250906118ed5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060068190556118fb6129e9565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60035461194b43611cf4565b1161199d576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156119b057fe5b14156119ec576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a16119e76129fc565b611a1d565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a1611a1d612a07565b565b60055460ff166002811115611a3057fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090611aa15760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b50600354611aae43611cf4565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090611b215760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314611b9d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b506000611bb7898989898989805190602001208989612a0f565b90506000611bcb8a8a8a8a8a8a8a8a612a86565b90506111f68c8c8484611d31565b600060055460ff166002811115611bec57fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b81525090611c615760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b50600080546001600160a01b038681166001600160a01b03199283161790925560018054868416908316178155600280549386169390921692909217905560048290556005805460ff19169091179055611cb9612ab9565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e881025b919050565b6000611d12600289898989898989612acb565b98975050505050505050565b6000611d12600289898989898989612b49565b611d546111e985611d428786612770565b86611d4d8887612d71565b600161279c565b611d5c612858565b50505050565b6000600488888888888888604051602001808960ff1660ff1660f81b8152600101886001600160a01b03166001600160a01b031660601b8152601401876001600160a01b03166001600160a01b031660601b815260140186815260200185805190602001908083835b60208310611dea5780518252601f199092019160209182019101611dcb565b51815160209384036101000a600019018019909216911617905292019586525084810193909352506040808401919091528051808403820181526060909301905281519101209d9c50505050505050505050505050565b600080611e518660008851612da6565b6040805160038082526080820190925291925060609190816020015b611e75613b01565b815260200190600190039081611e6d579050509050611e9c8a6001600160a01b0316612e9c565b81600081518110611ea957fe5b6020026020010181905250611ebd88612e9c565b81600281518110611eca57fe5b6020026020010181905250611ede82612f1a565b81600381518110611eeb57fe5b602090810291909101015260408051600380825260808201909252606091816020015b611f16613b01565b815260200190600190039081611f0e579050509050611f356004612e9c565b81600081518110611f4257fe5b6020026020010181905250611f5f8a6001600160a01b0316612e9c565b81600181518110611f6c57fe5b6020026020010181905250611f8082612f98565b81600281518110611f8d57fe5b6020026020010181905250611fe06040518060800160405280611faf8a612e9c565b8152602001611fbd89612e9c565b8152602001611fcb88612e9c565b8152602001611fd984612f98565b9052613048565b9b9a5050505050505050505050565b60408051600160f81b6020808301919091526001600160601b03196060998a1b811660218401529790981b909616603587015260498601949094526069850192909252608984015260a9808401919091528151808403909101815260c99092019052805191012090565b60408051600280825260608281019093526000929190816020015b61207c613b01565b8152602001906001900390816120745790505090506120a3886001600160a01b0316612e9c565b816000815181106120b057fe5b60200260200101819052506120c486612e9c565b816001815181106120d157fe5b602090810291909101015260408051600380825260808201909252606091816020015b6120fc613b01565b8152602001906001900390816120f457905050905061211b6001612e9c565b8160008151811061212857fe5b6020026020010181905250612145886001600160a01b0316612e9c565b8160018151811061215257fe5b602002602001018190525061216682612f98565b8160028151811061217357fe5b602090810291909101015260408051600480825260a08201909252606091816020015b61219e613b01565b8152602001906001900390816121965790505090506121bc87612e9c565b816000815181106121c957fe5b60200260200101819052506121dd86612e9c565b816001815181106121ea57fe5b60200260200101819052506121fe85612e9c565b8160028151811061220b57fe5b602002602001018190525061221f82612f98565b8160038151811061222c57fe5b602002602001018190525061224861224382612f98565b6130c8565b519a9950505050505050505050565b6000611d12600389898989898989612acb565b6000611d12600389898989898989612b49565b875187516000919081146122cd576040805162461bcd60e51b81526020600482015260126024820152710eee4dedcce40d2dce0eae840d8cadccee8d60731b604482015290519081900360640190fd5b80885114612317576040805162461bcd60e51b81526020600482015260126024820152710eee4dedcce40d2dce0eae840d8cadccee8d60731b604482015290519081900360640190fd5b80875114612361576040805162461bcd60e51b81526020600482015260126024820152710eee4dedcce40d2dce0eae840d8cadccee8d60731b604482015290519081900360640190fd5b60608b8b8b8b8b8b8b60405160200180886001600160a01b03166001600160a01b0316815260200180602001806020018060200180602001806020018060200187810387528d818151815260200191508051906020019060200280838360005b838110156123d95781810151838201526020016123c1565b5050505090500187810386528c818151815260200191508051906020019060200280838360005b83811015612418578181015183820152602001612400565b5050505090500187810385528b818151815260200191508051906020019060200280838360005b8381101561245757818101518382015260200161243f565b5050505090500187810384528a818151815260200191508051906020019060200280838360005b8381101561249657818101518382015260200161247e565b50505050905001878103835289818151815260200191508051906020019080838360005b838110156124d25781810151838201526020016124ba565b50505050905090810190601f1680156124ff5780820380516001836020036101000a031916815260200191505b5087810382528851815288516020918201918a019080838360005b8381101561253257818101518382015260200161251a565b50505050905090810190601f16801561255f5780820380516001836020036101000a031916815260200191505b509d505050505050505050505050505060405160208183030381529060405290506000815190506000601f830160068153505081810160208101879052604001859052602101601f909101209150509998505050505050505050565b60006125c5613b2f565b600080805b8b5181101561275e5760008982815181106125e157fe5b602002602001015163ffffffff169050808460208b01012092506126af8e8e848151811061260b57fe5b60200260200101518e858151811061261f57fe5b60200260200101518e868151811061263357fe5b60200260200101518760405160200180866001600160a01b03166001600160a01b031660601b8152601401856001600160a01b03166001600160a01b031660601b8152601401848152602001838152602001828152602001955050505050506040516020818303038152906040528051906020012089846131fe565b6001600160a01b031685526126c5898583612da6565b8560400181815250506127388e8e84815181106126de57fe5b602002602001015187600001518f86815181106126f757fe5b60200260200101518f878151811061270b57fe5b6020026020010151888b604001518e60006002811061272657fe5b60200201518f60016020020151613331565b85602001818152505061274f8f8660200151612d71565b9e5092909201916001016125ca565b509b9c9b505050505050505050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60408051602080820197909752808201959095526060850193909352608084019190915260a0808401919091528151808403909101815260c09092019052805191012090565b6006548114604051806040016040528060088152602001672124a9afa82922ab60c11b815250906128545760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156109a1578181015183820152602001610989565b5050565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1611a1d612a07565b600081838161289457fe5b0682848161289e57fe5b04019392505050565b60008183816128b257fe5b049392505050565b6128c38161353d565b60065550565b600580546002919060ff19166001835b0217905550611a1d612ab9565b600080838160205b885181116129d9578089015193506020818a51036020018161290c57fe5b0491505b6000821180156129235750600286066001145b801561293157508160020a86115b1561294457600286046001019550612910565b6002860661298f57838360405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161298757fe5b0495506129d1565b82846040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816129ca57fe5b0460010195505b6020016128ee565b505085149150505b949350505050565b600580546001919060ff191682806128d9565b612a0461367b565b33ff5b612a046136f1565b6040805160006020808301919091526001600160601b031960609b8c1b81166021840152998b1b8a1660358301529790991b9097166049890152605d880194909452607d870192909252609d86015260bd85015260dd808501919091528251808503909101815260fd909301909152815191012090565b6000612aac89898989898980519060200120612aa58b60008d51612da6565b8a8a613331565b9998505050505050505050565b600454612ac543611cf4565b01600355565b6040805160f89990991b6001600160f81b0319166020808b0191909152606098891b6001600160601b031990811660218c015297891b881660358b01529590971b9095166049880152605d870192909252607d860152609d85015260bd808501929092528251808503909201825260dd909301909152805191012090565b6040805160038082526080820190925260009160609190816020015b612b6d613b01565b815260200190600190039081612b65579050509050612b94876001600160a01b0316612e9c565b81600081518110612ba157fe5b6020026020010181905250612bbe896001600160a01b0316612e9c565b81600181518110612bcb57fe5b6020026020010181905250612bdf86612e9c565b81600281518110612bec57fe5b602090810291909101015260408051600380825260808201909252606091816020015b612c17613b01565b815260200190600190039081612c0f579050509050612c388b60ff16612e9c565b81600081518110612c4557fe5b6020026020010181905250612c62896001600160a01b0316612e9c565b81600181518110612c6f57fe5b6020026020010181905250612c8382612f98565b81600281518110612c9057fe5b602090810291909101015260408051600480825260a08201909252606091816020015b612cbb613b01565b815260200190600190039081612cb3579050509050612cd987612e9c565b81600081518110612ce657fe5b6020026020010181905250612cfa86612e9c565b81600181518110612d0757fe5b6020026020010181905250612d1b85612e9c565b81600281518110612d2857fe5b6020026020010181905250612d3c82612f98565b81600381518110612d4957fe5b6020026020010181905250612d6061224382612f98565b519c9b505050505050505050505050565b6000612d9f6040518060400160405280612d8a86612f1a565b8152602001612d9885612f1a565b9052613750565b9392505050565b6000602080830490601f84010482612dbc6137c5565b905060005b83811015612e0f57612e056040518060400160405280612de085612f1a565b8152602001612d98612e00856020028c018d61383890919063ffffffff16565b612e9c565b9150600101612dc1565b5081831015612e6c576000612e3088601f198989010163ffffffff61383816565b905083602002866020030360080281901b9050612e686040518060400160405280612e5a85612f1a565b8152602001612d9884612e9c565b9150505b612e916040518060400160405280612e8388612e9c565b8152602001612d9884612f1a565b979650505050505050565b612ea4613b01565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612f09565b612ef6613b01565b815260200190600190039081612eee5790505b508152600060209091015292915050565b612f22613b01565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612f87565b612f74613b01565b815260200190600190039081612f6c5790505b508152600260209091015292915050565b612fa0613b01565b612faa8251613854565b612ffb576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b60408051600480825260a0820190925260009160609190816020015b61306c613b01565b815260200190600190039081613064575050805190915060005b818110156130be5784816004811061309a57fe5b60200201518382815181106130ab57fe5b6020908102919091010152600101613086565b506129e18261385b565b6130d0613b4f565b6060820151600c60ff90911610613122576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661314f57604051806020016040528061314684600001516139a7565b90529050611cfa565b606082015160ff16600114156131965760405180602001604052806131468460200151600001518560200151604001518660200151606001518760200151602001516139cb565b606082015160ff16600214156131bb5750604080516020810190915281518152611cfa565b600360ff16826060015160ff16101580156131df57506060820151600c60ff909116105b156131fc576040518060200160405280613146846040015161385b565bfe5b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081896040516020018083805190602001908083835b602083106132745780518252601f199092019160209182019101613255565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381529382019052825192019190912092506132ba915089905088613a73565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa158015613319573d6000803e3d6000fd5b5050604051601f1901519a9950505050505050505050565b60408051600060208083018290526001600160601b031960608e811b821660218601528d811b821660358601528c811b9091166049850152605d84018b9052607d84018a9052609d8085018a90528551808603909101815260bd850180875281519190930120600480845261015d909501909552919392816020015b6133b5613b01565b8152602001906001900390816133ad5790505090506133dc8b6001600160a01b0316612e9c565b816000815181106133e957fe5b60200260200101819052506133fd89612e9c565b8160018151811061340a57fe5b602002602001018190525061341e88612e9c565b8160028151811061342b57fe5b602002602001018190525061343f86612f1a565b8160038151811061344c57fe5b602090810291909101015260408051600380825260808201909252606091816020015b613477613b01565b81526020019060019003908161346f5790505090506134966000612e9c565b816000815181106134a357fe5b60200260200101819052506134c08b6001600160a01b0316612e9c565b816001815181106134cd57fe5b60200260200101819052506134e182612f98565b816002815181106134ee57fe5b602002602001018190525061352c604051806080016040528061351089612e9c565b815260200161351e88612e9c565b8152602001611fcb86612e9c565b9d9c50505050505050505050505050565b6000815b60018151111561365e576060600282516001018161355b57fe5b04604051908082528060200260200182016040528015613585578160200160208202803883390190505b50905060005b815181101561365657825181600202600101101561361e578281600202815181106135b257fe5b60200260200101518382600202600101815181106135cc57fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012082828151811061360d57fe5b60200260200101818152505061364e565b82816002028151811061362d57fe5b602002602001015182828151811061364157fe5b6020026020010181815250505b60010161358b565b509050613541565b8060008151811061366b57fe5b6020026020010151915050919050565b6000805460025460018054604080516335e1e69160e11b81526001600160a01b0394851660048201529184166024830152604482019290925290519190921692636bc3cd22926064808201939182900301818387803b1580156136dd57600080fd5b505af1158015611d5c573d6000803e3d6000fd5b6000805460018054600254604080516335e1e69160e11b81526001600160a01b039384166004820152918316602483015260448201939093529151921692636bc3cd229260648084019382900301818387803b1580156136dd57600080fd5b60408051600280825260608281019093526000929190816020015b613773613b01565b81526020019060019003908161376b575050805190915060005b818110156130be578481600281106137a157fe5b60200201518382815181106137b257fe5b602090810291909101015260010161378d565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156138115781810151838201526020016137f9565b50505050905001925050506040516020818303038152906040528051906020012091505090565b6000816020018351101561384b57600080fd5b50016020015190565b6008101590565b60006008825111156138ab576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156138d8578160200160208202803883390190505b50805190915060005b81811015613934576138f1613b4f565b61390d86838151811061390057fe5b60200260200101516130c8565b9050806000015184838151811061392057fe5b6020908102919091010152506001016138e1565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561397d578181015183820152602001613965565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613a25575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206129e1565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b604180820283810160208101516040820151919093015160ff169291601b841015613a9f57601b840193505b8360ff16601b1480613ab457508360ff16601c145b613af9576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b604051806080016040528060008152602001613b1b613b61565b815260606020820152600060409091015290565b604080516060810182526000808252602082018190529181019190915290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a72315820bf51b338193c8004fb2d101262fdd58ce006f62c008c178ac3a64a9a97b1446064736f6c634300050d0032"

// DeployMessagesChallenge deploys a new Ethereum contract, binding an instance of MessagesChallenge to it.
func DeployMessagesChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessagesChallenge, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessagesChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessagesChallenge{MessagesChallengeCaller: MessagesChallengeCaller{contract: contract}, MessagesChallengeTransactor: MessagesChallengeTransactor{contract: contract}, MessagesChallengeFilterer: MessagesChallengeFilterer{contract: contract}}, nil
}

// MessagesChallenge is an auto generated Go binding around an Ethereum contract.
type MessagesChallenge struct {
	MessagesChallengeCaller     // Read-only binding to the contract
	MessagesChallengeTransactor // Write-only binding to the contract
	MessagesChallengeFilterer   // Log filterer for contract events
}

// MessagesChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagesChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagesChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagesChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagesChallengeSession struct {
	Contract     *MessagesChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MessagesChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagesChallengeCallerSession struct {
	Contract *MessagesChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MessagesChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagesChallengeTransactorSession struct {
	Contract     *MessagesChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MessagesChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagesChallengeRaw struct {
	Contract *MessagesChallenge // Generic contract binding to access the raw methods on
}

// MessagesChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagesChallengeCallerRaw struct {
	Contract *MessagesChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// MessagesChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagesChallengeTransactorRaw struct {
	Contract *MessagesChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessagesChallenge creates a new instance of MessagesChallenge, bound to a specific deployed contract.
func NewMessagesChallenge(address common.Address, backend bind.ContractBackend) (*MessagesChallenge, error) {
	contract, err := bindMessagesChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessagesChallenge{MessagesChallengeCaller: MessagesChallengeCaller{contract: contract}, MessagesChallengeTransactor: MessagesChallengeTransactor{contract: contract}, MessagesChallengeFilterer: MessagesChallengeFilterer{contract: contract}}, nil
}

// NewMessagesChallengeCaller creates a new read-only instance of MessagesChallenge, bound to a specific deployed contract.
func NewMessagesChallengeCaller(address common.Address, caller bind.ContractCaller) (*MessagesChallengeCaller, error) {
	contract, err := bindMessagesChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeCaller{contract: contract}, nil
}

// NewMessagesChallengeTransactor creates a new write-only instance of MessagesChallenge, bound to a specific deployed contract.
func NewMessagesChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagesChallengeTransactor, error) {
	contract, err := bindMessagesChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeTransactor{contract: contract}, nil
}

// NewMessagesChallengeFilterer creates a new log filterer instance of MessagesChallenge, bound to a specific deployed contract.
func NewMessagesChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagesChallengeFilterer, error) {
	contract, err := bindMessagesChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeFilterer{contract: contract}, nil
}

// bindMessagesChallenge binds a generic wrapper to an already deployed contract.
func bindMessagesChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessagesChallenge *MessagesChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MessagesChallenge.Contract.MessagesChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessagesChallenge *MessagesChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.MessagesChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessagesChallenge *MessagesChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.MessagesChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessagesChallenge *MessagesChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MessagesChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessagesChallenge *MessagesChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessagesChallenge *MessagesChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.contract.Transact(opts, method, params...)
}

// Bisect is a paid mutator transaction binding the contract method 0x500439e5.
//
// Solidity: function bisect(bytes32[] _chainHashes, bytes32[] _segmentHashes, uint256 _chainLength) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) Bisect(opts *bind.TransactOpts, _chainHashes [][32]byte, _segmentHashes [][32]byte, _chainLength *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "bisect", _chainHashes, _segmentHashes, _chainLength)
}

// Bisect is a paid mutator transaction binding the contract method 0x500439e5.
//
// Solidity: function bisect(bytes32[] _chainHashes, bytes32[] _segmentHashes, uint256 _chainLength) returns()
func (_MessagesChallenge *MessagesChallengeSession) Bisect(_chainHashes [][32]byte, _segmentHashes [][32]byte, _chainLength *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.Bisect(&_MessagesChallenge.TransactOpts, _chainHashes, _segmentHashes, _chainLength)
}

// Bisect is a paid mutator transaction binding the contract method 0x500439e5.
//
// Solidity: function bisect(bytes32[] _chainHashes, bytes32[] _segmentHashes, uint256 _chainLength) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) Bisect(_chainHashes [][32]byte, _segmentHashes [][32]byte, _chainLength *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.Bisect(&_MessagesChallenge.TransactOpts, _chainHashes, _segmentHashes, _chainLength)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) ChooseSegment(opts *bind.TransactOpts, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_MessagesChallenge *MessagesChallengeSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.ChooseSegment(&_MessagesChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.ChooseSegment(&_MessagesChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "initializeBisection", _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_MessagesChallenge *MessagesChallengeSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.InitializeBisection(&_MessagesChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.InitializeBisection(&_MessagesChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// OneStepProofContractTransactionMessage is a paid mutator transaction binding the contract method 0x225b3d09.
//
// Solidity: function oneStepProofContractTransactionMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, uint256 _value, bytes _data, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofContractTransactionMessage(opts *bind.TransactOpts, _lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _value *big.Int, _data []byte, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "oneStepProofContractTransactionMessage", _lowerHashA, _lowerHashB, _to, _from, _value, _data, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofContractTransactionMessage is a paid mutator transaction binding the contract method 0x225b3d09.
//
// Solidity: function oneStepProofContractTransactionMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, uint256 _value, bytes _data, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeSession) OneStepProofContractTransactionMessage(_lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _value *big.Int, _data []byte, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofContractTransactionMessage(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _to, _from, _value, _data, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofContractTransactionMessage is a paid mutator transaction binding the contract method 0x225b3d09.
//
// Solidity: function oneStepProofContractTransactionMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, uint256 _value, bytes _data, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) OneStepProofContractTransactionMessage(_lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _value *big.Int, _data []byte, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofContractTransactionMessage(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _to, _from, _value, _data, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofERC20Message is a paid mutator transaction binding the contract method 0x0c31dd2d.
//
// Solidity: function oneStepProofERC20Message(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, address _erc20, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofERC20Message(opts *bind.TransactOpts, _lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _erc20 common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "oneStepProofERC20Message", _lowerHashA, _lowerHashB, _to, _from, _erc20, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofERC20Message is a paid mutator transaction binding the contract method 0x0c31dd2d.
//
// Solidity: function oneStepProofERC20Message(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, address _erc20, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeSession) OneStepProofERC20Message(_lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _erc20 common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofERC20Message(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _to, _from, _erc20, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofERC20Message is a paid mutator transaction binding the contract method 0x0c31dd2d.
//
// Solidity: function oneStepProofERC20Message(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, address _erc20, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) OneStepProofERC20Message(_lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _erc20 common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofERC20Message(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _to, _from, _erc20, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofERC721Message is a paid mutator transaction binding the contract method 0x3f08600e.
//
// Solidity: function oneStepProofERC721Message(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, address _erc721, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofERC721Message(opts *bind.TransactOpts, _lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _erc721 common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "oneStepProofERC721Message", _lowerHashA, _lowerHashB, _to, _from, _erc721, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofERC721Message is a paid mutator transaction binding the contract method 0x3f08600e.
//
// Solidity: function oneStepProofERC721Message(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, address _erc721, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeSession) OneStepProofERC721Message(_lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _erc721 common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofERC721Message(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _to, _from, _erc721, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofERC721Message is a paid mutator transaction binding the contract method 0x3f08600e.
//
// Solidity: function oneStepProofERC721Message(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, address _erc721, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) OneStepProofERC721Message(_lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _erc721 common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofERC721Message(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _to, _from, _erc721, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofEthMessage is a paid mutator transaction binding the contract method 0x3261ee97.
//
// Solidity: function oneStepProofEthMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofEthMessage(opts *bind.TransactOpts, _lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "oneStepProofEthMessage", _lowerHashA, _lowerHashB, _to, _from, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofEthMessage is a paid mutator transaction binding the contract method 0x3261ee97.
//
// Solidity: function oneStepProofEthMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeSession) OneStepProofEthMessage(_lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofEthMessage(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _to, _from, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofEthMessage is a paid mutator transaction binding the contract method 0x3261ee97.
//
// Solidity: function oneStepProofEthMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _to, address _from, uint256 _value, uint256 _blockNumber, uint256 _timestamp, uint256 _messageNum) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) OneStepProofEthMessage(_lowerHashA [32]byte, _lowerHashB [32]byte, _to common.Address, _from common.Address, _value *big.Int, _blockNumber *big.Int, _timestamp *big.Int, _messageNum *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofEthMessage(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _to, _from, _value, _blockNumber, _timestamp, _messageNum)
}

// OneStepProofTransactionBatchMessage is a paid mutator transaction binding the contract method 0x4b7bea76.
//
// Solidity: function oneStepProofTransactionBatchMessage(bytes32 lowerHashA, bytes32 lowerHashB, address chain, address[] tos, uint256[] seqNumbers, uint256[] values, uint32[] dataLengths, bytes data, bytes signatures, uint256[2] blockAndTimestamp) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofTransactionBatchMessage(opts *bind.TransactOpts, lowerHashA [32]byte, lowerHashB [32]byte, chain common.Address, tos []common.Address, seqNumbers []*big.Int, values []*big.Int, dataLengths []uint32, data []byte, signatures []byte, blockAndTimestamp [2]*big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "oneStepProofTransactionBatchMessage", lowerHashA, lowerHashB, chain, tos, seqNumbers, values, dataLengths, data, signatures, blockAndTimestamp)
}

// OneStepProofTransactionBatchMessage is a paid mutator transaction binding the contract method 0x4b7bea76.
//
// Solidity: function oneStepProofTransactionBatchMessage(bytes32 lowerHashA, bytes32 lowerHashB, address chain, address[] tos, uint256[] seqNumbers, uint256[] values, uint32[] dataLengths, bytes data, bytes signatures, uint256[2] blockAndTimestamp) returns()
func (_MessagesChallenge *MessagesChallengeSession) OneStepProofTransactionBatchMessage(lowerHashA [32]byte, lowerHashB [32]byte, chain common.Address, tos []common.Address, seqNumbers []*big.Int, values []*big.Int, dataLengths []uint32, data []byte, signatures []byte, blockAndTimestamp [2]*big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofTransactionBatchMessage(&_MessagesChallenge.TransactOpts, lowerHashA, lowerHashB, chain, tos, seqNumbers, values, dataLengths, data, signatures, blockAndTimestamp)
}

// OneStepProofTransactionBatchMessage is a paid mutator transaction binding the contract method 0x4b7bea76.
//
// Solidity: function oneStepProofTransactionBatchMessage(bytes32 lowerHashA, bytes32 lowerHashB, address chain, address[] tos, uint256[] seqNumbers, uint256[] values, uint32[] dataLengths, bytes data, bytes signatures, uint256[2] blockAndTimestamp) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) OneStepProofTransactionBatchMessage(lowerHashA [32]byte, lowerHashB [32]byte, chain common.Address, tos []common.Address, seqNumbers []*big.Int, values []*big.Int, dataLengths []uint32, data []byte, signatures []byte, blockAndTimestamp [2]*big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofTransactionBatchMessage(&_MessagesChallenge.TransactOpts, lowerHashA, lowerHashB, chain, tos, seqNumbers, values, dataLengths, data, signatures, blockAndTimestamp)
}

// OneStepProofTransactionMessage is a paid mutator transaction binding the contract method 0xd2c56f27.
//
// Solidity: function oneStepProofTransactionMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _chain, address _to, address _from, uint256 _seqNumber, uint256 _value, bytes _data, uint256 _blockNumber, uint256 _timestamp) returns()
func (_MessagesChallenge *MessagesChallengeTransactor) OneStepProofTransactionMessage(opts *bind.TransactOpts, _lowerHashA [32]byte, _lowerHashB [32]byte, _chain common.Address, _to common.Address, _from common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _blockNumber *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "oneStepProofTransactionMessage", _lowerHashA, _lowerHashB, _chain, _to, _from, _seqNumber, _value, _data, _blockNumber, _timestamp)
}

// OneStepProofTransactionMessage is a paid mutator transaction binding the contract method 0xd2c56f27.
//
// Solidity: function oneStepProofTransactionMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _chain, address _to, address _from, uint256 _seqNumber, uint256 _value, bytes _data, uint256 _blockNumber, uint256 _timestamp) returns()
func (_MessagesChallenge *MessagesChallengeSession) OneStepProofTransactionMessage(_lowerHashA [32]byte, _lowerHashB [32]byte, _chain common.Address, _to common.Address, _from common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _blockNumber *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofTransactionMessage(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _chain, _to, _from, _seqNumber, _value, _data, _blockNumber, _timestamp)
}

// OneStepProofTransactionMessage is a paid mutator transaction binding the contract method 0xd2c56f27.
//
// Solidity: function oneStepProofTransactionMessage(bytes32 _lowerHashA, bytes32 _lowerHashB, address _chain, address _to, address _from, uint256 _seqNumber, uint256 _value, bytes _data, uint256 _blockNumber, uint256 _timestamp) returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) OneStepProofTransactionMessage(_lowerHashA [32]byte, _lowerHashB [32]byte, _chain common.Address, _to common.Address, _from common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _blockNumber *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _MessagesChallenge.Contract.OneStepProofTransactionMessage(&_MessagesChallenge.TransactOpts, _lowerHashA, _lowerHashB, _chain, _to, _from, _seqNumber, _value, _data, _blockNumber, _timestamp)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_MessagesChallenge *MessagesChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagesChallenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_MessagesChallenge *MessagesChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _MessagesChallenge.Contract.TimeoutChallenge(&_MessagesChallenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_MessagesChallenge *MessagesChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _MessagesChallenge.Contract.TimeoutChallenge(&_MessagesChallenge.TransactOpts)
}

// MessagesChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the MessagesChallenge contract.
type MessagesChallengeAsserterTimedOutIterator struct {
	Event *MessagesChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *MessagesChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagesChallengeAsserterTimedOut)
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
		it.Event = new(MessagesChallengeAsserterTimedOut)
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
func (it *MessagesChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagesChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagesChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the MessagesChallenge contract.
type MessagesChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_MessagesChallenge *MessagesChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*MessagesChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _MessagesChallenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeAsserterTimedOutIterator{contract: _MessagesChallenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_MessagesChallenge *MessagesChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *MessagesChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _MessagesChallenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagesChallengeAsserterTimedOut)
				if err := _MessagesChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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
func (_MessagesChallenge *MessagesChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*MessagesChallengeAsserterTimedOut, error) {
	event := new(MessagesChallengeAsserterTimedOut)
	if err := _MessagesChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MessagesChallengeBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the MessagesChallenge contract.
type MessagesChallengeBisectedIterator struct {
	Event *MessagesChallengeBisected // Event containing the contract specifics and raw log

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
func (it *MessagesChallengeBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagesChallengeBisected)
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
		it.Event = new(MessagesChallengeBisected)
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
func (it *MessagesChallengeBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagesChallengeBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagesChallengeBisected represents a Bisected event raised by the MessagesChallenge contract.
type MessagesChallengeBisected struct {
	ChainHashes   [][32]byte
	SegmentHashes [][32]byte
	TotalLength   *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x500c4a1bbd12a65d684bde95626a41abdd6a8c5d30f84c5c9b81e5bdb0cc0bd3.
//
// Solidity: event Bisected(bytes32[] chainHashes, bytes32[] segmentHashes, uint256 totalLength, uint256 deadlineTicks)
func (_MessagesChallenge *MessagesChallengeFilterer) FilterBisected(opts *bind.FilterOpts) (*MessagesChallengeBisectedIterator, error) {

	logs, sub, err := _MessagesChallenge.contract.FilterLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeBisectedIterator{contract: _MessagesChallenge.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x500c4a1bbd12a65d684bde95626a41abdd6a8c5d30f84c5c9b81e5bdb0cc0bd3.
//
// Solidity: event Bisected(bytes32[] chainHashes, bytes32[] segmentHashes, uint256 totalLength, uint256 deadlineTicks)
func (_MessagesChallenge *MessagesChallengeFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *MessagesChallengeBisected) (event.Subscription, error) {

	logs, sub, err := _MessagesChallenge.contract.WatchLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagesChallengeBisected)
				if err := _MessagesChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x500c4a1bbd12a65d684bde95626a41abdd6a8c5d30f84c5c9b81e5bdb0cc0bd3.
//
// Solidity: event Bisected(bytes32[] chainHashes, bytes32[] segmentHashes, uint256 totalLength, uint256 deadlineTicks)
func (_MessagesChallenge *MessagesChallengeFilterer) ParseBisected(log types.Log) (*MessagesChallengeBisected, error) {
	event := new(MessagesChallengeBisected)
	if err := _MessagesChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MessagesChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the MessagesChallenge contract.
type MessagesChallengeChallengerTimedOutIterator struct {
	Event *MessagesChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *MessagesChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagesChallengeChallengerTimedOut)
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
		it.Event = new(MessagesChallengeChallengerTimedOut)
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
func (it *MessagesChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagesChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagesChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the MessagesChallenge contract.
type MessagesChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_MessagesChallenge *MessagesChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*MessagesChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _MessagesChallenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeChallengerTimedOutIterator{contract: _MessagesChallenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_MessagesChallenge *MessagesChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *MessagesChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _MessagesChallenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagesChallengeChallengerTimedOut)
				if err := _MessagesChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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
func (_MessagesChallenge *MessagesChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*MessagesChallengeChallengerTimedOut, error) {
	event := new(MessagesChallengeChallengerTimedOut)
	if err := _MessagesChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MessagesChallengeContinuedIterator is returned from FilterContinued and is used to iterate over the raw logs and unpacked data for Continued events raised by the MessagesChallenge contract.
type MessagesChallengeContinuedIterator struct {
	Event *MessagesChallengeContinued // Event containing the contract specifics and raw log

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
func (it *MessagesChallengeContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagesChallengeContinued)
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
		it.Event = new(MessagesChallengeContinued)
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
func (it *MessagesChallengeContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagesChallengeContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagesChallengeContinued represents a Continued event raised by the MessagesChallenge contract.
type MessagesChallengeContinued struct {
	SegmentIndex  *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContinued is a free log retrieval operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_MessagesChallenge *MessagesChallengeFilterer) FilterContinued(opts *bind.FilterOpts) (*MessagesChallengeContinuedIterator, error) {

	logs, sub, err := _MessagesChallenge.contract.FilterLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeContinuedIterator{contract: _MessagesChallenge.contract, event: "Continued", logs: logs, sub: sub}, nil
}

// WatchContinued is a free log subscription operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_MessagesChallenge *MessagesChallengeFilterer) WatchContinued(opts *bind.WatchOpts, sink chan<- *MessagesChallengeContinued) (event.Subscription, error) {

	logs, sub, err := _MessagesChallenge.contract.WatchLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagesChallengeContinued)
				if err := _MessagesChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
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
func (_MessagesChallenge *MessagesChallengeFilterer) ParseContinued(log types.Log) (*MessagesChallengeContinued, error) {
	event := new(MessagesChallengeContinued)
	if err := _MessagesChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MessagesChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the MessagesChallenge contract.
type MessagesChallengeInitiatedChallengeIterator struct {
	Event *MessagesChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *MessagesChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagesChallengeInitiatedChallenge)
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
		it.Event = new(MessagesChallengeInitiatedChallenge)
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
func (it *MessagesChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagesChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagesChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the MessagesChallenge contract.
type MessagesChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_MessagesChallenge *MessagesChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*MessagesChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _MessagesChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeInitiatedChallengeIterator{contract: _MessagesChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_MessagesChallenge *MessagesChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *MessagesChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _MessagesChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagesChallengeInitiatedChallenge)
				if err := _MessagesChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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
func (_MessagesChallenge *MessagesChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*MessagesChallengeInitiatedChallenge, error) {
	event := new(MessagesChallengeInitiatedChallenge)
	if err := _MessagesChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MessagesChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the MessagesChallenge contract.
type MessagesChallengeOneStepProofCompletedIterator struct {
	Event *MessagesChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *MessagesChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagesChallengeOneStepProofCompleted)
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
		it.Event = new(MessagesChallengeOneStepProofCompleted)
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
func (it *MessagesChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagesChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagesChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the MessagesChallenge contract.
type MessagesChallengeOneStepProofCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_MessagesChallenge *MessagesChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*MessagesChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _MessagesChallenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &MessagesChallengeOneStepProofCompletedIterator{contract: _MessagesChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_MessagesChallenge *MessagesChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *MessagesChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _MessagesChallenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagesChallengeOneStepProofCompleted)
				if err := _MessagesChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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
func (_MessagesChallenge *MessagesChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*MessagesChallengeOneStepProofCompleted, error) {
	event := new(MessagesChallengeOneStepProofCompleted)
	if err := _MessagesChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582091cf918fd045609f547f9f78a6bb01fa06cf9bd44e9b70bddc32fa7ba1db60dd64736f6c634300050d0032"

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

// SigUtilsABI is the input ABI used to generate the binding from.
const SigUtilsABI = "[]"

// SigUtilsBin is the compiled bytecode used for deploying new contracts.
var SigUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158203009e1da8a87034267e27897ece1c70ebe00db684d53e6c89c2c669aee9bffa564736f6c634300050d0032"

// DeploySigUtils deploys a new Ethereum contract, binding an instance of SigUtils to it.
func DeploySigUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// SigUtils is an auto generated Go binding around an Ethereum contract.
type SigUtils struct {
	SigUtilsCaller     // Read-only binding to the contract
	SigUtilsTransactor // Write-only binding to the contract
	SigUtilsFilterer   // Log filterer for contract events
}

// SigUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsSession struct {
	Contract     *SigUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsCallerSession struct {
	Contract *SigUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTransactorSession struct {
	Contract     *SigUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsRaw struct {
	Contract *SigUtils // Generic contract binding to access the raw methods on
}

// SigUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsCallerRaw struct {
	Contract *SigUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTransactorRaw struct {
	Contract *SigUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtils creates a new instance of SigUtils, bound to a specific deployed contract.
func NewSigUtils(address common.Address, backend bind.ContractBackend) (*SigUtils, error) {
	contract, err := bindSigUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// NewSigUtilsCaller creates a new read-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsCaller, error) {
	contract, err := bindSigUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsCaller{contract: contract}, nil
}

// NewSigUtilsTransactor creates a new write-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTransactor, error) {
	contract, err := bindSigUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTransactor{contract: contract}, nil
}

// NewSigUtilsFilterer creates a new log filterer instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsFilterer, error) {
	contract, err := bindSigUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsFilterer{contract: contract}, nil
}

// bindSigUtils binds a generic wrapper to an already deployed contract.
func bindSigUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.SigUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208fdfb18a73ebb88d436490cd627d84adf2cd3acb35d8d407415ccc0ec0f186c364736f6c634300050d0032"

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
