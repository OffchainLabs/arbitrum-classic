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

// FibonacciABI is the input ABI used to generate the binding from.
const FibonacciABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"generateFib\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"getFib\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FibonacciBin is the compiled bytecode used for deploying new contracts.
var FibonacciBin = "0x608060405234801561001057600080fd5b506101b8806100206000396000f3fe6080604052600436106100295760003560e01c80632ddec39b1461002e57806390a3e3de1461004d575b600080fd5b61004b6004803603602081101561004457600080fd5b5035610089565b005b34801561005957600080fd5b506100776004803603602081101561007057600080fd5b5035610162565b60408051918252519081900360200190f35b60008054600181810183558280527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563918201819055825480820190935591015560025b8181101561012b5760008060028303815481106100e557fe5b90600052602060002001546000600184038154811061010057fe5b60009182526020808320909101548454600181810187559584529190922092909101910155016100cc565b506040805182815290517f1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d24149181900360200190a150565b600080828154811061017057fe5b9060005260206000200154905091905056fea2646970667358221220969aba9ca2a55008daac1dce8266cb0759b2e3a14b5b08f7c634b4bf5192918964736f6c634300060c0033"

// DeployFibonacci deploys a new Ethereum contract, binding an instance of Fibonacci to it.
func DeployFibonacci(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Fibonacci, error) {
	parsed, err := abi.JSON(strings.NewReader(FibonacciABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FibonacciBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fibonacci{FibonacciCaller: FibonacciCaller{contract: contract}, FibonacciTransactor: FibonacciTransactor{contract: contract}, FibonacciFilterer: FibonacciFilterer{contract: contract}}, nil
}

// Fibonacci is an auto generated Go binding around an Ethereum contract.
type Fibonacci struct {
	FibonacciCaller     // Read-only binding to the contract
	FibonacciTransactor // Write-only binding to the contract
	FibonacciFilterer   // Log filterer for contract events
}

// FibonacciCaller is an auto generated read-only Go binding around an Ethereum contract.
type FibonacciCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FibonacciTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FibonacciTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FibonacciFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FibonacciFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FibonacciSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FibonacciSession struct {
	Contract     *Fibonacci        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FibonacciCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FibonacciCallerSession struct {
	Contract *FibonacciCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FibonacciTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FibonacciTransactorSession struct {
	Contract     *FibonacciTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FibonacciRaw is an auto generated low-level Go binding around an Ethereum contract.
type FibonacciRaw struct {
	Contract *Fibonacci // Generic contract binding to access the raw methods on
}

// FibonacciCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FibonacciCallerRaw struct {
	Contract *FibonacciCaller // Generic read-only contract binding to access the raw methods on
}

// FibonacciTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FibonacciTransactorRaw struct {
	Contract *FibonacciTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFibonacci creates a new instance of Fibonacci, bound to a specific deployed contract.
func NewFibonacci(address common.Address, backend bind.ContractBackend) (*Fibonacci, error) {
	contract, err := bindFibonacci(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fibonacci{FibonacciCaller: FibonacciCaller{contract: contract}, FibonacciTransactor: FibonacciTransactor{contract: contract}, FibonacciFilterer: FibonacciFilterer{contract: contract}}, nil
}

// NewFibonacciCaller creates a new read-only instance of Fibonacci, bound to a specific deployed contract.
func NewFibonacciCaller(address common.Address, caller bind.ContractCaller) (*FibonacciCaller, error) {
	contract, err := bindFibonacci(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FibonacciCaller{contract: contract}, nil
}

// NewFibonacciTransactor creates a new write-only instance of Fibonacci, bound to a specific deployed contract.
func NewFibonacciTransactor(address common.Address, transactor bind.ContractTransactor) (*FibonacciTransactor, error) {
	contract, err := bindFibonacci(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FibonacciTransactor{contract: contract}, nil
}

// NewFibonacciFilterer creates a new log filterer instance of Fibonacci, bound to a specific deployed contract.
func NewFibonacciFilterer(address common.Address, filterer bind.ContractFilterer) (*FibonacciFilterer, error) {
	contract, err := bindFibonacci(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FibonacciFilterer{contract: contract}, nil
}

// bindFibonacci binds a generic wrapper to an already deployed contract.
func bindFibonacci(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FibonacciABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fibonacci *FibonacciRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fibonacci.Contract.FibonacciCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fibonacci *FibonacciRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fibonacci.Contract.FibonacciTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fibonacci *FibonacciRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fibonacci.Contract.FibonacciTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fibonacci *FibonacciCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fibonacci.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fibonacci *FibonacciTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fibonacci.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fibonacci *FibonacciTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fibonacci.Contract.contract.Transact(opts, method, params...)
}

// GetFib is a free data retrieval call binding the contract method 0x90a3e3de.
//
// Solidity: function getFib(uint256 n) view returns(uint256)
func (_Fibonacci *FibonacciCaller) GetFib(opts *bind.CallOpts, n *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fibonacci.contract.Call(opts, &out, "getFib", n)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFib is a free data retrieval call binding the contract method 0x90a3e3de.
//
// Solidity: function getFib(uint256 n) view returns(uint256)
func (_Fibonacci *FibonacciSession) GetFib(n *big.Int) (*big.Int, error) {
	return _Fibonacci.Contract.GetFib(&_Fibonacci.CallOpts, n)
}

// GetFib is a free data retrieval call binding the contract method 0x90a3e3de.
//
// Solidity: function getFib(uint256 n) view returns(uint256)
func (_Fibonacci *FibonacciCallerSession) GetFib(n *big.Int) (*big.Int, error) {
	return _Fibonacci.Contract.GetFib(&_Fibonacci.CallOpts, n)
}

// GenerateFib is a paid mutator transaction binding the contract method 0x2ddec39b.
//
// Solidity: function generateFib(uint256 n) payable returns()
func (_Fibonacci *FibonacciTransactor) GenerateFib(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Fibonacci.contract.Transact(opts, "generateFib", n)
}

// GenerateFib is a paid mutator transaction binding the contract method 0x2ddec39b.
//
// Solidity: function generateFib(uint256 n) payable returns()
func (_Fibonacci *FibonacciSession) GenerateFib(n *big.Int) (*types.Transaction, error) {
	return _Fibonacci.Contract.GenerateFib(&_Fibonacci.TransactOpts, n)
}

// GenerateFib is a paid mutator transaction binding the contract method 0x2ddec39b.
//
// Solidity: function generateFib(uint256 n) payable returns()
func (_Fibonacci *FibonacciTransactorSession) GenerateFib(n *big.Int) (*types.Transaction, error) {
	return _Fibonacci.Contract.GenerateFib(&_Fibonacci.TransactOpts, n)
}

// FibonacciTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the Fibonacci contract.
type FibonacciTestEventIterator struct {
	Event *FibonacciTestEvent // Event containing the contract specifics and raw log

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
func (it *FibonacciTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FibonacciTestEvent)
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
		it.Event = new(FibonacciTestEvent)
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
func (it *FibonacciTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FibonacciTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FibonacciTestEvent represents a TestEvent event raised by the Fibonacci contract.
type FibonacciTestEvent struct {
	Number *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 number)
func (_Fibonacci *FibonacciFilterer) FilterTestEvent(opts *bind.FilterOpts) (*FibonacciTestEventIterator, error) {

	logs, sub, err := _Fibonacci.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &FibonacciTestEventIterator{contract: _Fibonacci.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 number)
func (_Fibonacci *FibonacciFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *FibonacciTestEvent) (event.Subscription, error) {

	logs, sub, err := _Fibonacci.contract.WatchLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FibonacciTestEvent)
				if err := _Fibonacci.contract.UnpackLog(event, "TestEvent", log); err != nil {
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
// Solidity: event TestEvent(uint256 number)
func (_Fibonacci *FibonacciFilterer) ParseTestEvent(log types.Log) (*FibonacciTestEvent, error) {
	event := new(FibonacciTestEvent)
	if err := _Fibonacci.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
