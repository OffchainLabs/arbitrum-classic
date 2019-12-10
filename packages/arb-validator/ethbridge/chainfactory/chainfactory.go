// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainfactory

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

// ChainFactoryABI is the input ABI used to generate the binding from.
const ChainFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"name\":\"ChainCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChainFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ChainFactoryFuncSigs = map[string]string{
	"b7399786": "createChain(bytes32,uint32,uint32,uint128,address)",
}

// ChainFactoryBin is the compiled bytecode used for deploying new contracts.
var ChainFactoryBin = "0x608060405234801561001057600080fd5b5060405161029a38038061029a8339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b039384166001600160a01b031991821617909155600180549484169482169490941790935560028054929091169190921617905561020b8061008f6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063b739978614610030575b600080fd5b610082600480360360a081101561004657600080fd5b50803590602081013563ffffffff9081169160408101359091169060608101356001600160801b031690608001356001600160a01b0316610084565b005b6000805461009a906001600160a01b0316610184565b60025460015460408051638364fe4760e01b8152600481018b905263ffffffff808b166024830152891660448201526001600160801b03881660648201526001600160a01b03878116608483015293841660a482015291831660c48301525192935090831691638364fe479160e48082019260009290919082900301818387803b15801561012757600080fd5b505af115801561013b573d6000803e3d6000fd5b5050604080516001600160a01b038516815290517fa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f9350908190036020019150a1505050505050565b6000808260601b9050604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528160148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f094935050505056fea265627a7a7231582019faebc319adec13b73c918a28a3d787b5bc273fb70990cf60e6755ed8c3741864736f6c634300050d0032"

// DeployChainFactory deploys a new Ethereum contract, binding an instance of ChainFactory to it.
func DeployChainFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _chainTemplate common.Address, _globalInboxAddress common.Address, _challengeFactoryAddress common.Address) (common.Address, *types.Transaction, *ChainFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChainFactoryBin), backend, _chainTemplate, _globalInboxAddress, _challengeFactoryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChainFactory{ChainFactoryCaller: ChainFactoryCaller{contract: contract}, ChainFactoryTransactor: ChainFactoryTransactor{contract: contract}, ChainFactoryFilterer: ChainFactoryFilterer{contract: contract}}, nil
}

// ChainFactory is an auto generated Go binding around an Ethereum contract.
type ChainFactory struct {
	ChainFactoryCaller     // Read-only binding to the contract
	ChainFactoryTransactor // Write-only binding to the contract
	ChainFactoryFilterer   // Log filterer for contract events
}

// ChainFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainFactorySession struct {
	Contract     *ChainFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainFactoryCallerSession struct {
	Contract *ChainFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ChainFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainFactoryTransactorSession struct {
	Contract     *ChainFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChainFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainFactoryRaw struct {
	Contract *ChainFactory // Generic contract binding to access the raw methods on
}

// ChainFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainFactoryCallerRaw struct {
	Contract *ChainFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ChainFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainFactoryTransactorRaw struct {
	Contract *ChainFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainFactory creates a new instance of ChainFactory, bound to a specific deployed contract.
func NewChainFactory(address common.Address, backend bind.ContractBackend) (*ChainFactory, error) {
	contract, err := bindChainFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainFactory{ChainFactoryCaller: ChainFactoryCaller{contract: contract}, ChainFactoryTransactor: ChainFactoryTransactor{contract: contract}, ChainFactoryFilterer: ChainFactoryFilterer{contract: contract}}, nil
}

// NewChainFactoryCaller creates a new read-only instance of ChainFactory, bound to a specific deployed contract.
func NewChainFactoryCaller(address common.Address, caller bind.ContractCaller) (*ChainFactoryCaller, error) {
	contract, err := bindChainFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainFactoryCaller{contract: contract}, nil
}

// NewChainFactoryTransactor creates a new write-only instance of ChainFactory, bound to a specific deployed contract.
func NewChainFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainFactoryTransactor, error) {
	contract, err := bindChainFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainFactoryTransactor{contract: contract}, nil
}

// NewChainFactoryFilterer creates a new log filterer instance of ChainFactory, bound to a specific deployed contract.
func NewChainFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainFactoryFilterer, error) {
	contract, err := bindChainFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainFactoryFilterer{contract: contract}, nil
}

// bindChainFactory binds a generic wrapper to an already deployed contract.
func bindChainFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainFactory *ChainFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainFactory.Contract.ChainFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainFactory *ChainFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainFactory.Contract.ChainFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainFactory *ChainFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainFactory.Contract.ChainFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainFactory *ChainFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainFactory *ChainFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainFactory *ChainFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateChain is a paid mutator transaction binding the contract method 0xb7399786.
//
// Solidity: function createChain(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner) returns()
func (_ChainFactory *ChainFactoryTransactor) CreateChain(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ChainFactory.contract.Transact(opts, "createChain", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner)
}

// CreateChain is a paid mutator transaction binding the contract method 0xb7399786.
//
// Solidity: function createChain(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner) returns()
func (_ChainFactory *ChainFactorySession) CreateChain(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ChainFactory.Contract.CreateChain(&_ChainFactory.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner)
}

// CreateChain is a paid mutator transaction binding the contract method 0xb7399786.
//
// Solidity: function createChain(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner) returns()
func (_ChainFactory *ChainFactoryTransactorSession) CreateChain(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ChainFactory.Contract.CreateChain(&_ChainFactory.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner)
}

// ChainFactoryChainCreatedIterator is returned from FilterChainCreated and is used to iterate over the raw logs and unpacked data for ChainCreated events raised by the ChainFactory contract.
type ChainFactoryChainCreatedIterator struct {
	Event *ChainFactoryChainCreated // Event containing the contract specifics and raw log

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
func (it *ChainFactoryChainCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainFactoryChainCreated)
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
		it.Event = new(ChainFactoryChainCreated)
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
func (it *ChainFactoryChainCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainFactoryChainCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainFactoryChainCreated represents a ChainCreated event raised by the ChainFactory contract.
type ChainFactoryChainCreated struct {
	VmAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChainCreated is a free log retrieval operation binding the contract event 0xa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f.
//
// Solidity: event ChainCreated(address vmAddress)
func (_ChainFactory *ChainFactoryFilterer) FilterChainCreated(opts *bind.FilterOpts) (*ChainFactoryChainCreatedIterator, error) {

	logs, sub, err := _ChainFactory.contract.FilterLogs(opts, "ChainCreated")
	if err != nil {
		return nil, err
	}
	return &ChainFactoryChainCreatedIterator{contract: _ChainFactory.contract, event: "ChainCreated", logs: logs, sub: sub}, nil
}

// WatchChainCreated is a free log subscription operation binding the contract event 0xa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f.
//
// Solidity: event ChainCreated(address vmAddress)
func (_ChainFactory *ChainFactoryFilterer) WatchChainCreated(opts *bind.WatchOpts, sink chan<- *ChainFactoryChainCreated) (event.Subscription, error) {

	logs, sub, err := _ChainFactory.contract.WatchLogs(opts, "ChainCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainFactoryChainCreated)
				if err := _ChainFactory.contract.UnpackLog(event, "ChainCreated", log); err != nil {
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

// ParseChainCreated is a log parse operation binding the contract event 0xa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f.
//
// Solidity: event ChainCreated(address vmAddress)
func (_ChainFactory *ChainFactoryFilterer) ParseChainCreated(log types.Log) (*ChainFactoryChainCreated, error) {
	event := new(ChainFactoryChainCreated)
	if err := _ChainFactory.contract.UnpackLog(event, "ChainCreated", log); err != nil {
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

// IArbChainABI is the input ABI used to generate the binding from.
const IArbChainABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeLauncherAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbChainFuncSigs maps the 4-byte function signature to its string representation.
var IArbChainFuncSigs = map[string]string{
	"8364fe47": "init(bytes32,uint32,uint32,uint128,address,address,address)",
}

// IArbChain is an auto generated Go binding around an Ethereum contract.
type IArbChain struct {
	IArbChainCaller     // Read-only binding to the contract
	IArbChainTransactor // Write-only binding to the contract
	IArbChainFilterer   // Log filterer for contract events
}

// IArbChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbChainSession struct {
	Contract     *IArbChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbChainCallerSession struct {
	Contract *IArbChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IArbChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbChainTransactorSession struct {
	Contract     *IArbChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IArbChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbChainRaw struct {
	Contract *IArbChain // Generic contract binding to access the raw methods on
}

// IArbChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbChainCallerRaw struct {
	Contract *IArbChainCaller // Generic read-only contract binding to access the raw methods on
}

// IArbChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbChainTransactorRaw struct {
	Contract *IArbChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbChain creates a new instance of IArbChain, bound to a specific deployed contract.
func NewIArbChain(address common.Address, backend bind.ContractBackend) (*IArbChain, error) {
	contract, err := bindIArbChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbChain{IArbChainCaller: IArbChainCaller{contract: contract}, IArbChainTransactor: IArbChainTransactor{contract: contract}, IArbChainFilterer: IArbChainFilterer{contract: contract}}, nil
}

// NewIArbChainCaller creates a new read-only instance of IArbChain, bound to a specific deployed contract.
func NewIArbChainCaller(address common.Address, caller bind.ContractCaller) (*IArbChainCaller, error) {
	contract, err := bindIArbChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChainCaller{contract: contract}, nil
}

// NewIArbChainTransactor creates a new write-only instance of IArbChain, bound to a specific deployed contract.
func NewIArbChainTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbChainTransactor, error) {
	contract, err := bindIArbChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChainTransactor{contract: contract}, nil
}

// NewIArbChainFilterer creates a new log filterer instance of IArbChain, bound to a specific deployed contract.
func NewIArbChainFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbChainFilterer, error) {
	contract, err := bindIArbChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbChainFilterer{contract: contract}, nil
}

// bindIArbChain binds a generic wrapper to an already deployed contract.
func bindIArbChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChain *IArbChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChain.Contract.IArbChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChain *IArbChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChain.Contract.IArbChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChain *IArbChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChain.Contract.IArbChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChain *IArbChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChain *IArbChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChain *IArbChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChain.Contract.contract.Transact(opts, method, params...)
}

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_IArbChain *IArbChainTransactor) Init(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbChain.contract.Transact(opts, "init", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_IArbChain *IArbChainSession) Init(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbChain.Contract.Init(&_IArbChain.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_IArbChain *IArbChainTransactorSession) Init(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbChain.Contract.Init(&_IArbChain.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}
