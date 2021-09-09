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

// ArbSysMetaData contains all meta data concerning the ArbSys contract.
var ArbSysMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"isTopLevelCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ArbSysABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbSysMetaData.ABI instead.
var ArbSysABI = ArbSysMetaData.ABI

// ArbSys is an auto generated Go binding around an Ethereum contract.
type ArbSys struct {
	ArbSysCaller     // Read-only binding to the contract
	ArbSysTransactor // Write-only binding to the contract
	ArbSysFilterer   // Log filterer for contract events
}

// ArbSysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbSysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbSysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbSysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbSysSession struct {
	Contract     *ArbSys           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbSysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbSysCallerSession struct {
	Contract *ArbSysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbSysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbSysTransactorSession struct {
	Contract     *ArbSysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbSysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbSysRaw struct {
	Contract *ArbSys // Generic contract binding to access the raw methods on
}

// ArbSysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbSysCallerRaw struct {
	Contract *ArbSysCaller // Generic read-only contract binding to access the raw methods on
}

// ArbSysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbSysTransactorRaw struct {
	Contract *ArbSysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbSys creates a new instance of ArbSys, bound to a specific deployed contract.
func NewArbSys(address common.Address, backend bind.ContractBackend) (*ArbSys, error) {
	contract, err := bindArbSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbSys{ArbSysCaller: ArbSysCaller{contract: contract}, ArbSysTransactor: ArbSysTransactor{contract: contract}, ArbSysFilterer: ArbSysFilterer{contract: contract}}, nil
}

// NewArbSysCaller creates a new read-only instance of ArbSys, bound to a specific deployed contract.
func NewArbSysCaller(address common.Address, caller bind.ContractCaller) (*ArbSysCaller, error) {
	contract, err := bindArbSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbSysCaller{contract: contract}, nil
}

// NewArbSysTransactor creates a new write-only instance of ArbSys, bound to a specific deployed contract.
func NewArbSysTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbSysTransactor, error) {
	contract, err := bindArbSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbSysTransactor{contract: contract}, nil
}

// NewArbSysFilterer creates a new log filterer instance of ArbSys, bound to a specific deployed contract.
func NewArbSysFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbSysFilterer, error) {
	contract, err := bindArbSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbSysFilterer{contract: contract}, nil
}

// bindArbSys binds a generic wrapper to an already deployed contract.
func bindArbSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbSysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbSys *ArbSysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbSys.Contract.ArbSysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbSys *ArbSysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbSys.Contract.ArbSysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbSys *ArbSysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbSys.Contract.ArbSysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbSys *ArbSysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbSys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbSys *ArbSysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbSys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbSys *ArbSysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbSys.Contract.contract.Transact(opts, method, params...)
}

// IsTopLevelCall is a free data retrieval call binding the contract method 0x08bd624c.
//
// Solidity: function isTopLevelCall() view returns(bool)
func (_ArbSys *ArbSysCaller) IsTopLevelCall(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "isTopLevelCall")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTopLevelCall is a free data retrieval call binding the contract method 0x08bd624c.
//
// Solidity: function isTopLevelCall() view returns(bool)
func (_ArbSys *ArbSysSession) IsTopLevelCall() (bool, error) {
	return _ArbSys.Contract.IsTopLevelCall(&_ArbSys.CallOpts)
}

// IsTopLevelCall is a free data retrieval call binding the contract method 0x08bd624c.
//
// Solidity: function isTopLevelCall() view returns(bool)
func (_ArbSys *ArbSysCallerSession) IsTopLevelCall() (bool, error) {
	return _ArbSys.Contract.IsTopLevelCall(&_ArbSys.CallOpts)
}

// TopLevelMetaData contains all meta data concerning the TopLevel contract.
var TopLevelMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"top\",\"type\":\"bool\"}],\"name\":\"TopLevelEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"isTopLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nestedNotTop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610173806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806339f484b71461003b578063813c13b114610045575b600080fd5b61004361004d565b005b6100436100e8565b600060646001600160a01b03166308bd624c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561008957600080fd5b505afa15801561009d573d6000803e3d6000fd5b505050506040513d60208110156100b357600080fd5b5051604051909150811515907f9cd8e0cef591b8295292293a053ccf65910a134f855c5fdb104fe56fa1d0722d90600090a250565b306001600160a01b03166339f484b76040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561012357600080fd5b505af1158015610137573d6000803e3d6000fd5b5050505056fea2646970667358221220451e18e5c6fb7a8f3facbd6fbecfbb9134c3b40eb7fcc8904b58b0ed97b375f364736f6c634300060c0033",
}

// TopLevelABI is the input ABI used to generate the binding from.
// Deprecated: Use TopLevelMetaData.ABI instead.
var TopLevelABI = TopLevelMetaData.ABI

// TopLevelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TopLevelMetaData.Bin instead.
var TopLevelBin = TopLevelMetaData.Bin

// DeployTopLevel deploys a new Ethereum contract, binding an instance of TopLevel to it.
func DeployTopLevel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TopLevel, error) {
	parsed, err := TopLevelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TopLevelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TopLevel{TopLevelCaller: TopLevelCaller{contract: contract}, TopLevelTransactor: TopLevelTransactor{contract: contract}, TopLevelFilterer: TopLevelFilterer{contract: contract}}, nil
}

// TopLevel is an auto generated Go binding around an Ethereum contract.
type TopLevel struct {
	TopLevelCaller     // Read-only binding to the contract
	TopLevelTransactor // Write-only binding to the contract
	TopLevelFilterer   // Log filterer for contract events
}

// TopLevelCaller is an auto generated read-only Go binding around an Ethereum contract.
type TopLevelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TopLevelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TopLevelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TopLevelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TopLevelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TopLevelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TopLevelSession struct {
	Contract     *TopLevel         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TopLevelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TopLevelCallerSession struct {
	Contract *TopLevelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TopLevelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TopLevelTransactorSession struct {
	Contract     *TopLevelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TopLevelRaw is an auto generated low-level Go binding around an Ethereum contract.
type TopLevelRaw struct {
	Contract *TopLevel // Generic contract binding to access the raw methods on
}

// TopLevelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TopLevelCallerRaw struct {
	Contract *TopLevelCaller // Generic read-only contract binding to access the raw methods on
}

// TopLevelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TopLevelTransactorRaw struct {
	Contract *TopLevelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTopLevel creates a new instance of TopLevel, bound to a specific deployed contract.
func NewTopLevel(address common.Address, backend bind.ContractBackend) (*TopLevel, error) {
	contract, err := bindTopLevel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TopLevel{TopLevelCaller: TopLevelCaller{contract: contract}, TopLevelTransactor: TopLevelTransactor{contract: contract}, TopLevelFilterer: TopLevelFilterer{contract: contract}}, nil
}

// NewTopLevelCaller creates a new read-only instance of TopLevel, bound to a specific deployed contract.
func NewTopLevelCaller(address common.Address, caller bind.ContractCaller) (*TopLevelCaller, error) {
	contract, err := bindTopLevel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TopLevelCaller{contract: contract}, nil
}

// NewTopLevelTransactor creates a new write-only instance of TopLevel, bound to a specific deployed contract.
func NewTopLevelTransactor(address common.Address, transactor bind.ContractTransactor) (*TopLevelTransactor, error) {
	contract, err := bindTopLevel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TopLevelTransactor{contract: contract}, nil
}

// NewTopLevelFilterer creates a new log filterer instance of TopLevel, bound to a specific deployed contract.
func NewTopLevelFilterer(address common.Address, filterer bind.ContractFilterer) (*TopLevelFilterer, error) {
	contract, err := bindTopLevel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TopLevelFilterer{contract: contract}, nil
}

// bindTopLevel binds a generic wrapper to an already deployed contract.
func bindTopLevel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TopLevelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TopLevel *TopLevelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TopLevel.Contract.TopLevelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TopLevel *TopLevelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopLevel.Contract.TopLevelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TopLevel *TopLevelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TopLevel.Contract.TopLevelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TopLevel *TopLevelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TopLevel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TopLevel *TopLevelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopLevel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TopLevel *TopLevelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TopLevel.Contract.contract.Transact(opts, method, params...)
}

// IsTopLevel is a paid mutator transaction binding the contract method 0x39f484b7.
//
// Solidity: function isTopLevel() returns()
func (_TopLevel *TopLevelTransactor) IsTopLevel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopLevel.contract.Transact(opts, "isTopLevel")
}

// IsTopLevel is a paid mutator transaction binding the contract method 0x39f484b7.
//
// Solidity: function isTopLevel() returns()
func (_TopLevel *TopLevelSession) IsTopLevel() (*types.Transaction, error) {
	return _TopLevel.Contract.IsTopLevel(&_TopLevel.TransactOpts)
}

// IsTopLevel is a paid mutator transaction binding the contract method 0x39f484b7.
//
// Solidity: function isTopLevel() returns()
func (_TopLevel *TopLevelTransactorSession) IsTopLevel() (*types.Transaction, error) {
	return _TopLevel.Contract.IsTopLevel(&_TopLevel.TransactOpts)
}

// NestedNotTop is a paid mutator transaction binding the contract method 0x813c13b1.
//
// Solidity: function nestedNotTop() returns()
func (_TopLevel *TopLevelTransactor) NestedNotTop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopLevel.contract.Transact(opts, "nestedNotTop")
}

// NestedNotTop is a paid mutator transaction binding the contract method 0x813c13b1.
//
// Solidity: function nestedNotTop() returns()
func (_TopLevel *TopLevelSession) NestedNotTop() (*types.Transaction, error) {
	return _TopLevel.Contract.NestedNotTop(&_TopLevel.TransactOpts)
}

// NestedNotTop is a paid mutator transaction binding the contract method 0x813c13b1.
//
// Solidity: function nestedNotTop() returns()
func (_TopLevel *TopLevelTransactorSession) NestedNotTop() (*types.Transaction, error) {
	return _TopLevel.Contract.NestedNotTop(&_TopLevel.TransactOpts)
}

// TopLevelTopLevelEventIterator is returned from FilterTopLevelEvent and is used to iterate over the raw logs and unpacked data for TopLevelEvent events raised by the TopLevel contract.
type TopLevelTopLevelEventIterator struct {
	Event *TopLevelTopLevelEvent // Event containing the contract specifics and raw log

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
func (it *TopLevelTopLevelEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TopLevelTopLevelEvent)
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
		it.Event = new(TopLevelTopLevelEvent)
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
func (it *TopLevelTopLevelEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TopLevelTopLevelEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TopLevelTopLevelEvent represents a TopLevelEvent event raised by the TopLevel contract.
type TopLevelTopLevelEvent struct {
	Top bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTopLevelEvent is a free log retrieval operation binding the contract event 0x9cd8e0cef591b8295292293a053ccf65910a134f855c5fdb104fe56fa1d0722d.
//
// Solidity: event TopLevelEvent(bool indexed top)
func (_TopLevel *TopLevelFilterer) FilterTopLevelEvent(opts *bind.FilterOpts, top []bool) (*TopLevelTopLevelEventIterator, error) {

	var topRule []interface{}
	for _, topItem := range top {
		topRule = append(topRule, topItem)
	}

	logs, sub, err := _TopLevel.contract.FilterLogs(opts, "TopLevelEvent", topRule)
	if err != nil {
		return nil, err
	}
	return &TopLevelTopLevelEventIterator{contract: _TopLevel.contract, event: "TopLevelEvent", logs: logs, sub: sub}, nil
}

// WatchTopLevelEvent is a free log subscription operation binding the contract event 0x9cd8e0cef591b8295292293a053ccf65910a134f855c5fdb104fe56fa1d0722d.
//
// Solidity: event TopLevelEvent(bool indexed top)
func (_TopLevel *TopLevelFilterer) WatchTopLevelEvent(opts *bind.WatchOpts, sink chan<- *TopLevelTopLevelEvent, top []bool) (event.Subscription, error) {

	var topRule []interface{}
	for _, topItem := range top {
		topRule = append(topRule, topItem)
	}

	logs, sub, err := _TopLevel.contract.WatchLogs(opts, "TopLevelEvent", topRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TopLevelTopLevelEvent)
				if err := _TopLevel.contract.UnpackLog(event, "TopLevelEvent", log); err != nil {
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

// ParseTopLevelEvent is a log parse operation binding the contract event 0x9cd8e0cef591b8295292293a053ccf65910a134f855c5fdb104fe56fa1d0722d.
//
// Solidity: event TopLevelEvent(bool indexed top)
func (_TopLevel *TopLevelFilterer) ParseTopLevelEvent(log types.Log) (*TopLevelTopLevelEvent, error) {
	event := new(TopLevelTopLevelEvent)
	if err := _TopLevel.contract.UnpackLog(event, "TopLevelEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
