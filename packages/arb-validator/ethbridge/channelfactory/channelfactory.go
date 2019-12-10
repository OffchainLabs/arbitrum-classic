// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package channelfactory

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

// ChannelFactoryABI is the input ABI used to generate the binding from.
const ChannelFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_channelTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"name\":\"ChannelCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_validatorKeys\",\"type\":\"address[]\"}],\"name\":\"createChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChannelFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ChannelFactoryFuncSigs = map[string]string{
	"7d35dde4": "createChannel(bytes32,uint32,uint32,uint128,address,address[])",
}

// ChannelFactoryBin is the compiled bytecode used for deploying new contracts.
var ChannelFactoryBin = "0x608060405234801561001057600080fd5b506040516103823803806103828339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b039384166001600160a01b03199182161790915560018054948416948216949094179093556002805492909116919092161790556102f38061008f6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80637d35dde414610030575b600080fd5b61010e600480360360c081101561004657600080fd5b81359163ffffffff60208201358116926040830135909116916001600160801b03606082013516916001600160a01b03608083013516919081019060c0810160a082013564010000000081111561009c57600080fd5b8201836020820111156100ae57600080fd5b803590602001918460208302840111640100000000831117156100d057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610110945050505050565b005b60008054610126906001600160a01b031661026c565b600254600154604051637588110b60e01b8152600481018b815263ffffffff808c1660248401528a1660448301526001600160801b03891660648301526001600160a01b03888116608484015293841660a4830181905292841660c4830181905261010060e484019081528851610104850152885196975094871695637588110b958e958e958e958e958e95939490938e9361012401906020808601910280838360005b838110156101e25781810151838201526020016101ca565b505050509050019950505050505050505050600060405180830381600087803b15801561020e57600080fd5b505af1158015610222573d6000803e3d6000fd5b5050604080516001600160a01b038516815290517fc625d37dd8b556110d70984e62f74ba35c77422c83c5f548fbd21b697a67ef5c9350908190036020019150a150505050505050565b6000808260601b9050604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528160148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f094935050505056fea265627a7a723158202732dd7bf4cdaec4766b897d80fd729f224d687ebeb1b056c8d75f03a462536564736f6c634300050d0032"

// DeployChannelFactory deploys a new Ethereum contract, binding an instance of ChannelFactory to it.
func DeployChannelFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _channelTemplate common.Address, _globalInboxAddress common.Address, _challengeFactoryAddress common.Address) (common.Address, *types.Transaction, *ChannelFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelFactoryBin), backend, _channelTemplate, _globalInboxAddress, _challengeFactoryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChannelFactory{ChannelFactoryCaller: ChannelFactoryCaller{contract: contract}, ChannelFactoryTransactor: ChannelFactoryTransactor{contract: contract}, ChannelFactoryFilterer: ChannelFactoryFilterer{contract: contract}}, nil
}

// ChannelFactory is an auto generated Go binding around an Ethereum contract.
type ChannelFactory struct {
	ChannelFactoryCaller     // Read-only binding to the contract
	ChannelFactoryTransactor // Write-only binding to the contract
	ChannelFactoryFilterer   // Log filterer for contract events
}

// ChannelFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelFactorySession struct {
	Contract     *ChannelFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChannelFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelFactoryCallerSession struct {
	Contract *ChannelFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ChannelFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelFactoryTransactorSession struct {
	Contract     *ChannelFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChannelFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelFactoryRaw struct {
	Contract *ChannelFactory // Generic contract binding to access the raw methods on
}

// ChannelFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelFactoryCallerRaw struct {
	Contract *ChannelFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelFactoryTransactorRaw struct {
	Contract *ChannelFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannelFactory creates a new instance of ChannelFactory, bound to a specific deployed contract.
func NewChannelFactory(address common.Address, backend bind.ContractBackend) (*ChannelFactory, error) {
	contract, err := bindChannelFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChannelFactory{ChannelFactoryCaller: ChannelFactoryCaller{contract: contract}, ChannelFactoryTransactor: ChannelFactoryTransactor{contract: contract}, ChannelFactoryFilterer: ChannelFactoryFilterer{contract: contract}}, nil
}

// NewChannelFactoryCaller creates a new read-only instance of ChannelFactory, bound to a specific deployed contract.
func NewChannelFactoryCaller(address common.Address, caller bind.ContractCaller) (*ChannelFactoryCaller, error) {
	contract, err := bindChannelFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelFactoryCaller{contract: contract}, nil
}

// NewChannelFactoryTransactor creates a new write-only instance of ChannelFactory, bound to a specific deployed contract.
func NewChannelFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelFactoryTransactor, error) {
	contract, err := bindChannelFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelFactoryTransactor{contract: contract}, nil
}

// NewChannelFactoryFilterer creates a new log filterer instance of ChannelFactory, bound to a specific deployed contract.
func NewChannelFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelFactoryFilterer, error) {
	contract, err := bindChannelFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelFactoryFilterer{contract: contract}, nil
}

// bindChannelFactory binds a generic wrapper to an already deployed contract.
func bindChannelFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelFactory *ChannelFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelFactory.Contract.ChannelFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelFactory *ChannelFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelFactory.Contract.ChannelFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelFactory *ChannelFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelFactory.Contract.ChannelFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelFactory *ChannelFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelFactory *ChannelFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelFactory *ChannelFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x7d35dde4.
//
// Solidity: function createChannel(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address[] _validatorKeys) returns()
func (_ChannelFactory *ChannelFactoryTransactor) CreateChannel(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _ChannelFactory.contract.Transact(opts, "createChannel", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _validatorKeys)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x7d35dde4.
//
// Solidity: function createChannel(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address[] _validatorKeys) returns()
func (_ChannelFactory *ChannelFactorySession) CreateChannel(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _ChannelFactory.Contract.CreateChannel(&_ChannelFactory.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _validatorKeys)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x7d35dde4.
//
// Solidity: function createChannel(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address[] _validatorKeys) returns()
func (_ChannelFactory *ChannelFactoryTransactorSession) CreateChannel(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _ChannelFactory.Contract.CreateChannel(&_ChannelFactory.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _validatorKeys)
}

// ChannelFactoryChannelCreatedIterator is returned from FilterChannelCreated and is used to iterate over the raw logs and unpacked data for ChannelCreated events raised by the ChannelFactory contract.
type ChannelFactoryChannelCreatedIterator struct {
	Event *ChannelFactoryChannelCreated // Event containing the contract specifics and raw log

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
func (it *ChannelFactoryChannelCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelFactoryChannelCreated)
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
		it.Event = new(ChannelFactoryChannelCreated)
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
func (it *ChannelFactoryChannelCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelFactoryChannelCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelFactoryChannelCreated represents a ChannelCreated event raised by the ChannelFactory contract.
type ChannelFactoryChannelCreated struct {
	VmAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelCreated is a free log retrieval operation binding the contract event 0xc625d37dd8b556110d70984e62f74ba35c77422c83c5f548fbd21b697a67ef5c.
//
// Solidity: event ChannelCreated(address vmAddress)
func (_ChannelFactory *ChannelFactoryFilterer) FilterChannelCreated(opts *bind.FilterOpts) (*ChannelFactoryChannelCreatedIterator, error) {

	logs, sub, err := _ChannelFactory.contract.FilterLogs(opts, "ChannelCreated")
	if err != nil {
		return nil, err
	}
	return &ChannelFactoryChannelCreatedIterator{contract: _ChannelFactory.contract, event: "ChannelCreated", logs: logs, sub: sub}, nil
}

// WatchChannelCreated is a free log subscription operation binding the contract event 0xc625d37dd8b556110d70984e62f74ba35c77422c83c5f548fbd21b697a67ef5c.
//
// Solidity: event ChannelCreated(address vmAddress)
func (_ChannelFactory *ChannelFactoryFilterer) WatchChannelCreated(opts *bind.WatchOpts, sink chan<- *ChannelFactoryChannelCreated) (event.Subscription, error) {

	logs, sub, err := _ChannelFactory.contract.WatchLogs(opts, "ChannelCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelFactoryChannelCreated)
				if err := _ChannelFactory.contract.UnpackLog(event, "ChannelCreated", log); err != nil {
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

// ParseChannelCreated is a log parse operation binding the contract event 0xc625d37dd8b556110d70984e62f74ba35c77422c83c5f548fbd21b697a67ef5c.
//
// Solidity: event ChannelCreated(address vmAddress)
func (_ChannelFactory *ChannelFactoryFilterer) ParseChannelCreated(log types.Log) (*ChannelFactoryChannelCreated, error) {
	event := new(ChannelFactoryChannelCreated)
	if err := _ChannelFactory.contract.UnpackLog(event, "ChannelCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CloneFactoryABI is the input ABI used to generate the binding from.
const CloneFactoryABI = "[]"

// CloneFactoryBin is the compiled bytecode used for deploying new contracts.
var CloneFactoryBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820878796a809b90734a93e84a3759d73ac4b0e17a8e5f99d67f46df0f3aad1545a64736f6c634300050d0032"

// DeployCloneFactory deploys a new Ethereum contract, binding an instance of CloneFactory to it.
func DeployCloneFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CloneFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(CloneFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CloneFactoryBin), backend)
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
func (_CloneFactory *CloneFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_CloneFactory *CloneFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// IArbChannelABI is the input ABI used to generate the binding from.
const IArbChannelABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeLauncherAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_validatorKeys\",\"type\":\"address[]\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"isValidatorList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IArbChannelFuncSigs maps the 4-byte function signature to its string representation.
var IArbChannelFuncSigs = map[string]string{
	"7588110b": "init(bytes32,uint32,uint32,uint128,address,address,address,address[])",
	"513164fe": "isValidatorList(address[])",
}

// IArbChannel is an auto generated Go binding around an Ethereum contract.
type IArbChannel struct {
	IArbChannelCaller     // Read-only binding to the contract
	IArbChannelTransactor // Write-only binding to the contract
	IArbChannelFilterer   // Log filterer for contract events
}

// IArbChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbChannelSession struct {
	Contract     *IArbChannel      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbChannelCallerSession struct {
	Contract *IArbChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IArbChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbChannelTransactorSession struct {
	Contract     *IArbChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IArbChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbChannelRaw struct {
	Contract *IArbChannel // Generic contract binding to access the raw methods on
}

// IArbChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbChannelCallerRaw struct {
	Contract *IArbChannelCaller // Generic read-only contract binding to access the raw methods on
}

// IArbChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbChannelTransactorRaw struct {
	Contract *IArbChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbChannel creates a new instance of IArbChannel, bound to a specific deployed contract.
func NewIArbChannel(address common.Address, backend bind.ContractBackend) (*IArbChannel, error) {
	contract, err := bindIArbChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbChannel{IArbChannelCaller: IArbChannelCaller{contract: contract}, IArbChannelTransactor: IArbChannelTransactor{contract: contract}, IArbChannelFilterer: IArbChannelFilterer{contract: contract}}, nil
}

// NewIArbChannelCaller creates a new read-only instance of IArbChannel, bound to a specific deployed contract.
func NewIArbChannelCaller(address common.Address, caller bind.ContractCaller) (*IArbChannelCaller, error) {
	contract, err := bindIArbChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChannelCaller{contract: contract}, nil
}

// NewIArbChannelTransactor creates a new write-only instance of IArbChannel, bound to a specific deployed contract.
func NewIArbChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbChannelTransactor, error) {
	contract, err := bindIArbChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChannelTransactor{contract: contract}, nil
}

// NewIArbChannelFilterer creates a new log filterer instance of IArbChannel, bound to a specific deployed contract.
func NewIArbChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbChannelFilterer, error) {
	contract, err := bindIArbChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbChannelFilterer{contract: contract}, nil
}

// bindIArbChannel binds a generic wrapper to an already deployed contract.
func bindIArbChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChannel *IArbChannelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChannel.Contract.IArbChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChannel *IArbChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChannel.Contract.IArbChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChannel *IArbChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChannel.Contract.IArbChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChannel *IArbChannelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChannel *IArbChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChannel *IArbChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChannel.Contract.contract.Transact(opts, method, params...)
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_IArbChannel *IArbChannelCaller) IsValidatorList(opts *bind.CallOpts, _validators []common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IArbChannel.contract.Call(opts, out, "isValidatorList", _validators)
	return *ret0, err
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_IArbChannel *IArbChannelSession) IsValidatorList(_validators []common.Address) (bool, error) {
	return _IArbChannel.Contract.IsValidatorList(&_IArbChannel.CallOpts, _validators)
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_IArbChannel *IArbChannelCallerSession) IsValidatorList(_validators []common.Address) (bool, error) {
	return _IArbChannel.Contract.IsValidatorList(&_IArbChannel.CallOpts, _validators)
}

// Init is a paid mutator transaction binding the contract method 0x7588110b.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress, address[] _validatorKeys) returns()
func (_IArbChannel *IArbChannelTransactor) Init(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _IArbChannel.contract.Transact(opts, "init", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress, _validatorKeys)
}

// Init is a paid mutator transaction binding the contract method 0x7588110b.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress, address[] _validatorKeys) returns()
func (_IArbChannel *IArbChannelSession) Init(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _IArbChannel.Contract.Init(&_IArbChannel.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress, _validatorKeys)
}

// Init is a paid mutator transaction binding the contract method 0x7588110b.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress, address[] _validatorKeys) returns()
func (_IArbChannel *IArbChannelTransactorSession) Init(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _IArbChannel.Contract.Init(&_IArbChannel.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress, _validatorKeys)
}
