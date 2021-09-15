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

// CloneFactoryMetaData contains all meta data concerning the CloneFactory contract.
var CloneFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"}],\"name\":\"CreatedClone\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"create2Clone\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610143806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c91091c314610030575b600080fd5b61005c6004803603604081101561004657600080fd5b506001600160a01b038135169060200135610078565b604080516001600160a01b039092168252519081900360200190f35b6000808360601b90506000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260148201526e5af43d82803e903d91602b57fd5bf360881b6028820152846037826000f5604080516001600160a01b038316815290519193507f8bbdbba0e10077e3bdd81d5076242c5eca7c410250c1bf0ff4a0d8e40a6a8b31925081900360200190a194935050505056fea2646970667358221220983e5b820a0e32dac3b93a774092e3c2450a0e396c89998e51e07d808766643364736f6c634300060c0033",
}

// CloneFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CloneFactoryMetaData.ABI instead.
var CloneFactoryABI = CloneFactoryMetaData.ABI

// CloneFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CloneFactoryMetaData.Bin instead.
var CloneFactoryBin = CloneFactoryMetaData.Bin

// DeployCloneFactory deploys a new Ethereum contract, binding an instance of CloneFactory to it.
func DeployCloneFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CloneFactory, error) {
	parsed, err := CloneFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CloneFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CloneFactory{CloneFactoryCaller: CloneFactoryCaller{contract: contract}, CloneFactoryTransactor: CloneFactoryTransactor{contract: contract}, CloneFactoryFilterer: CloneFactoryFilterer{contract: contract}}, nil
}

// CloneFactory is an auto generated Go binding around an Ethereum contract.
type CloneFactory struct {
	CloneFactoryCaller     // Read-only binding to the contract
	CloneFactoryTransactor // Write-only binding to the contract
	CloneFactoryFilterer   // Log filterer for contract events
}

// CloneFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CloneFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CloneFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CloneFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CloneFactorySession struct {
	Contract     *CloneFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CloneFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CloneFactoryCallerSession struct {
	Contract *CloneFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CloneFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CloneFactoryTransactorSession struct {
	Contract     *CloneFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CloneFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CloneFactoryRaw struct {
	Contract *CloneFactory // Generic contract binding to access the raw methods on
}

// CloneFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CloneFactoryCallerRaw struct {
	Contract *CloneFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CloneFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CloneFactoryTransactorRaw struct {
	Contract *CloneFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCloneFactory creates a new instance of CloneFactory, bound to a specific deployed contract.
func NewCloneFactory(address common.Address, backend bind.ContractBackend) (*CloneFactory, error) {
	contract, err := bindCloneFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CloneFactory{CloneFactoryCaller: CloneFactoryCaller{contract: contract}, CloneFactoryTransactor: CloneFactoryTransactor{contract: contract}, CloneFactoryFilterer: CloneFactoryFilterer{contract: contract}}, nil
}

// NewCloneFactoryCaller creates a new read-only instance of CloneFactory, bound to a specific deployed contract.
func NewCloneFactoryCaller(address common.Address, caller bind.ContractCaller) (*CloneFactoryCaller, error) {
	contract, err := bindCloneFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CloneFactoryCaller{contract: contract}, nil
}

// NewCloneFactoryTransactor creates a new write-only instance of CloneFactory, bound to a specific deployed contract.
func NewCloneFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CloneFactoryTransactor, error) {
	contract, err := bindCloneFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CloneFactoryTransactor{contract: contract}, nil
}

// NewCloneFactoryFilterer creates a new log filterer instance of CloneFactory, bound to a specific deployed contract.
func NewCloneFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CloneFactoryFilterer, error) {
	contract, err := bindCloneFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CloneFactoryFilterer{contract: contract}, nil
}

// bindCloneFactory binds a generic wrapper to an already deployed contract.
func bindCloneFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CloneFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CloneFactory *CloneFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CloneFactory.Contract.CloneFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CloneFactory *CloneFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CloneFactory.Contract.CloneFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CloneFactory *CloneFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CloneFactory.Contract.CloneFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CloneFactory *CloneFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CloneFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CloneFactory *CloneFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CloneFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CloneFactory *CloneFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CloneFactory.Contract.contract.Transact(opts, method, params...)
}

// Create2Clone is a paid mutator transaction binding the contract method 0xc91091c3.
//
// Solidity: function create2Clone(address target, uint256 salt) returns(address)
func (_CloneFactory *CloneFactoryTransactor) Create2Clone(opts *bind.TransactOpts, target common.Address, salt *big.Int) (*types.Transaction, error) {
	return _CloneFactory.contract.Transact(opts, "create2Clone", target, salt)
}

// Create2Clone is a paid mutator transaction binding the contract method 0xc91091c3.
//
// Solidity: function create2Clone(address target, uint256 salt) returns(address)
func (_CloneFactory *CloneFactorySession) Create2Clone(target common.Address, salt *big.Int) (*types.Transaction, error) {
	return _CloneFactory.Contract.Create2Clone(&_CloneFactory.TransactOpts, target, salt)
}

// Create2Clone is a paid mutator transaction binding the contract method 0xc91091c3.
//
// Solidity: function create2Clone(address target, uint256 salt) returns(address)
func (_CloneFactory *CloneFactoryTransactorSession) Create2Clone(target common.Address, salt *big.Int) (*types.Transaction, error) {
	return _CloneFactory.Contract.Create2Clone(&_CloneFactory.TransactOpts, target, salt)
}

// CloneFactoryCreatedCloneIterator is returned from FilterCreatedClone and is used to iterate over the raw logs and unpacked data for CreatedClone events raised by the CloneFactory contract.
type CloneFactoryCreatedCloneIterator struct {
	Event *CloneFactoryCreatedClone // Event containing the contract specifics and raw log

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
func (it *CloneFactoryCreatedCloneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloneFactoryCreatedClone)
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
		it.Event = new(CloneFactoryCreatedClone)
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
func (it *CloneFactoryCreatedCloneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloneFactoryCreatedCloneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloneFactoryCreatedClone represents a CreatedClone event raised by the CloneFactory contract.
type CloneFactoryCreatedClone struct {
	Clone common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreatedClone is a free log retrieval operation binding the contract event 0x8bbdbba0e10077e3bdd81d5076242c5eca7c410250c1bf0ff4a0d8e40a6a8b31.
//
// Solidity: event CreatedClone(address clone)
func (_CloneFactory *CloneFactoryFilterer) FilterCreatedClone(opts *bind.FilterOpts) (*CloneFactoryCreatedCloneIterator, error) {

	logs, sub, err := _CloneFactory.contract.FilterLogs(opts, "CreatedClone")
	if err != nil {
		return nil, err
	}
	return &CloneFactoryCreatedCloneIterator{contract: _CloneFactory.contract, event: "CreatedClone", logs: logs, sub: sub}, nil
}

// WatchCreatedClone is a free log subscription operation binding the contract event 0x8bbdbba0e10077e3bdd81d5076242c5eca7c410250c1bf0ff4a0d8e40a6a8b31.
//
// Solidity: event CreatedClone(address clone)
func (_CloneFactory *CloneFactoryFilterer) WatchCreatedClone(opts *bind.WatchOpts, sink chan<- *CloneFactoryCreatedClone) (event.Subscription, error) {

	logs, sub, err := _CloneFactory.contract.WatchLogs(opts, "CreatedClone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloneFactoryCreatedClone)
				if err := _CloneFactory.contract.UnpackLog(event, "CreatedClone", log); err != nil {
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

// ParseCreatedClone is a log parse operation binding the contract event 0x8bbdbba0e10077e3bdd81d5076242c5eca7c410250c1bf0ff4a0d8e40a6a8b31.
//
// Solidity: event CreatedClone(address clone)
func (_CloneFactory *CloneFactoryFilterer) ParseCreatedClone(log types.Log) (*CloneFactoryCreatedClone, error) {
	event := new(CloneFactoryCreatedClone)
	if err := _CloneFactory.contract.UnpackLog(event, "CreatedClone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
