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

// TransferMetaData contains all meta data concerning the Transfer contract.
var TransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"wrapped\",\"type\":\"address\"}],\"name\":\"send2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"send3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"send4\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405261023b806100136000396000f3fe6080604052600436106100435760003560e01c80633386b1a214610082578063540d7a2f146100b75780636b2e1f1a146100e1578063b46300ec146100f65761007d565b3661007d576040805134815290517f1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d24149181900360200190a1005b600080fd5b34801561008e57600080fd5b506100b5600480360360208110156100a557600080fd5b50356001600160a01b031661010b565b005b3480156100c357600080fd5b506100b5600480360360208110156100da57600080fd5b5035610161565b3480156100ed57600080fd5b506100b56101b0565b34801561010257600080fd5b506100b56101dc565b806001600160a01b031663b46300ec6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561014657600080fd5b505af115801561015a573d6000803e3d6000fd5b5050505050565b604051309082906001906000818181858888f193505050503d80600081146101a5576040519150601f19603f3d011682016040523d82523d6000602084013e6101aa565b606091505b50505050565b604051309060009060019082818181858883f193505050501580156101d9573d6000803e3d6000fd5b50565b604051339060009060019082818181858883f193505050501580156101d9573d6000803e3d6000fdfea2646970667358221220c9157f14331e6d3610def92788c663aa59baee48c850d26b9645e23b8bef6c4164736f6c634300060c0033",
}

// TransferABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferMetaData.ABI instead.
var TransferABI = TransferMetaData.ABI

// TransferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferMetaData.Bin instead.
var TransferBin = TransferMetaData.Bin

// DeployTransfer deploys a new Ethereum contract, binding an instance of Transfer to it.
func DeployTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Transfer, error) {
	parsed, err := TransferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

// Transfer is an auto generated Go binding around an Ethereum contract.
type Transfer struct {
	TransferCaller     // Read-only binding to the contract
	TransferTransactor // Write-only binding to the contract
	TransferFilterer   // Log filterer for contract events
}

// TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferSession struct {
	Contract     *Transfer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferCallerSession struct {
	Contract *TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferTransactorSession struct {
	Contract     *TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferRaw struct {
	Contract *Transfer // Generic contract binding to access the raw methods on
}

// TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferCallerRaw struct {
	Contract *TransferCaller // Generic read-only contract binding to access the raw methods on
}

// TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferTransactorRaw struct {
	Contract *TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer creates a new instance of Transfer, bound to a specific deployed contract.
func NewTransfer(address common.Address, backend bind.ContractBackend) (*Transfer, error) {
	contract, err := bindTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

// NewTransferCaller creates a new read-only instance of Transfer, bound to a specific deployed contract.
func NewTransferCaller(address common.Address, caller bind.ContractCaller) (*TransferCaller, error) {
	contract, err := bindTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferCaller{contract: contract}, nil
}

// NewTransferTransactor creates a new write-only instance of Transfer, bound to a specific deployed contract.
func NewTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferTransactor, error) {
	contract, err := bindTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferTransactor{contract: contract}, nil
}

// NewTransferFilterer creates a new log filterer instance of Transfer, bound to a specific deployed contract.
func NewTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferFilterer, error) {
	contract, err := bindTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferFilterer{contract: contract}, nil
}

// bindTransfer binds a generic wrapper to an already deployed contract.
func bindTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transact(opts, method, params...)
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_Transfer *TransferTransactor) Send(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "send")
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_Transfer *TransferSession) Send() (*types.Transaction, error) {
	return _Transfer.Contract.Send(&_Transfer.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_Transfer *TransferTransactorSession) Send() (*types.Transaction, error) {
	return _Transfer.Contract.Send(&_Transfer.TransactOpts)
}

// Send2 is a paid mutator transaction binding the contract method 0x3386b1a2.
//
// Solidity: function send2(address wrapped) returns()
func (_Transfer *TransferTransactor) Send2(opts *bind.TransactOpts, wrapped common.Address) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "send2", wrapped)
}

// Send2 is a paid mutator transaction binding the contract method 0x3386b1a2.
//
// Solidity: function send2(address wrapped) returns()
func (_Transfer *TransferSession) Send2(wrapped common.Address) (*types.Transaction, error) {
	return _Transfer.Contract.Send2(&_Transfer.TransactOpts, wrapped)
}

// Send2 is a paid mutator transaction binding the contract method 0x3386b1a2.
//
// Solidity: function send2(address wrapped) returns()
func (_Transfer *TransferTransactorSession) Send2(wrapped common.Address) (*types.Transaction, error) {
	return _Transfer.Contract.Send2(&_Transfer.TransactOpts, wrapped)
}

// Send3 is a paid mutator transaction binding the contract method 0x6b2e1f1a.
//
// Solidity: function send3() returns()
func (_Transfer *TransferTransactor) Send3(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "send3")
}

// Send3 is a paid mutator transaction binding the contract method 0x6b2e1f1a.
//
// Solidity: function send3() returns()
func (_Transfer *TransferSession) Send3() (*types.Transaction, error) {
	return _Transfer.Contract.Send3(&_Transfer.TransactOpts)
}

// Send3 is a paid mutator transaction binding the contract method 0x6b2e1f1a.
//
// Solidity: function send3() returns()
func (_Transfer *TransferTransactorSession) Send3() (*types.Transaction, error) {
	return _Transfer.Contract.Send3(&_Transfer.TransactOpts)
}

// Send4 is a paid mutator transaction binding the contract method 0x540d7a2f.
//
// Solidity: function send4(uint256 gas) returns()
func (_Transfer *TransferTransactor) Send4(opts *bind.TransactOpts, gas *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "send4", gas)
}

// Send4 is a paid mutator transaction binding the contract method 0x540d7a2f.
//
// Solidity: function send4(uint256 gas) returns()
func (_Transfer *TransferSession) Send4(gas *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Send4(&_Transfer.TransactOpts, gas)
}

// Send4 is a paid mutator transaction binding the contract method 0x540d7a2f.
//
// Solidity: function send4(uint256 gas) returns()
func (_Transfer *TransferTransactorSession) Send4(gas *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Send4(&_Transfer.TransactOpts, gas)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transfer *TransferTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transfer *TransferSession) Receive() (*types.Transaction, error) {
	return _Transfer.Contract.Receive(&_Transfer.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transfer *TransferTransactorSession) Receive() (*types.Transaction, error) {
	return _Transfer.Contract.Receive(&_Transfer.TransactOpts)
}

// TransferTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the Transfer contract.
type TransferTestEventIterator struct {
	Event *TransferTestEvent // Event containing the contract specifics and raw log

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
func (it *TransferTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferTestEvent)
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
		it.Event = new(TransferTestEvent)
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
func (it *TransferTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferTestEvent represents a TestEvent event raised by the Transfer contract.
type TransferTestEvent struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 value)
func (_Transfer *TransferFilterer) FilterTestEvent(opts *bind.FilterOpts) (*TransferTestEventIterator, error) {

	logs, sub, err := _Transfer.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &TransferTestEventIterator{contract: _Transfer.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 value)
func (_Transfer *TransferFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *TransferTestEvent) (event.Subscription, error) {

	logs, sub, err := _Transfer.contract.WatchLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferTestEvent)
				if err := _Transfer.contract.UnpackLog(event, "TestEvent", log); err != nil {
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
// Solidity: event TestEvent(uint256 value)
func (_Transfer *TransferFilterer) ParseTestEvent(log types.Log) (*TransferTestEvent, error) {
	event := new(TransferTestEvent)
	if err := _Transfer.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
