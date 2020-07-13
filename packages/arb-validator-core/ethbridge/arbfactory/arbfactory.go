// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbfactory

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

// ArbFactoryABI is the input ABI used to generate the binding from.
const ArbFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"_gracePeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_arbGasSpeedLimitPerTick\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_stakeRequirement\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createRollup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInboxAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rollupTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ArbFactoryFuncSigs = map[string]string{
	"62e3c0b1": "challengeFactoryAddress()",
	"f67e0439": "createRollup(bytes32,uint128,uint128,uint64,uint128,address)",
	"582923c7": "globalInboxAddress()",
	"8689d996": "rollupTemplate()",
}

// ArbFactoryBin is the compiled bytecode used for deploying new contracts.
var ArbFactoryBin = "0x608060405234801561001057600080fd5b506040516103363803806103368339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b039384166001600160a01b03199182161790915560018054948416948216949094179093556002805492909116919092161790556102a78061008f6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063582923c71461005157806362e3c0b1146100755780638689d9961461007d578063f67e043914610085575b600080fd5b6100596100e5565b604080516001600160a01b039092168252519081900360200190f35b6100596100f4565b610059610103565b6100e3600480360360c081101561009b57600080fd5b5080359060208101356001600160801b03908116916040810135821691606082013567ffffffffffffffff169160808101359091169060a001356001600160a01b0316610112565b005b6001546001600160a01b031681565b6002546001600160a01b031681565b6000546001600160a01b031681565b60008054610128906001600160a01b0316610220565b600254600154604080516319f4c3a760e01b8152600481018c90526001600160801b03808c166024830152808b16604483015267ffffffffffffffff8a166064830152881660848201526001600160a01b0387811660a483015293841660c482015291831660e483015251929350908316916319f4c3a7916101048082019260009290919082900301818387803b1580156101c257600080fd5b505af11580156101d6573d6000803e3d6000fd5b5050604080516001600160a01b038516815290517f84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c9350908190036020019150a150505050505050565b6000808260601b9050604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528160148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f094935050505056fea265627a7a72315820758e90592396810e25ae1b05de59b8d3280babaa14824f54fb442e557fbe57af64736f6c63430005110032"

// DeployArbFactory deploys a new Ethereum contract, binding an instance of ArbFactory to it.
func DeployArbFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _rollupTemplate common.Address, _globalInboxAddress common.Address, _challengeFactoryAddress common.Address) (common.Address, *types.Transaction, *ArbFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbFactoryBin), backend, _rollupTemplate, _globalInboxAddress, _challengeFactoryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbFactory{ArbFactoryCaller: ArbFactoryCaller{contract: contract}, ArbFactoryTransactor: ArbFactoryTransactor{contract: contract}, ArbFactoryFilterer: ArbFactoryFilterer{contract: contract}}, nil
}

// ArbFactory is an auto generated Go binding around an Ethereum contract.
type ArbFactory struct {
	ArbFactoryCaller     // Read-only binding to the contract
	ArbFactoryTransactor // Write-only binding to the contract
	ArbFactoryFilterer   // Log filterer for contract events
}

// ArbFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbFactorySession struct {
	Contract     *ArbFactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbFactoryCallerSession struct {
	Contract *ArbFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbFactoryTransactorSession struct {
	Contract     *ArbFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbFactoryRaw struct {
	Contract *ArbFactory // Generic contract binding to access the raw methods on
}

// ArbFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbFactoryCallerRaw struct {
	Contract *ArbFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ArbFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbFactoryTransactorRaw struct {
	Contract *ArbFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbFactory creates a new instance of ArbFactory, bound to a specific deployed contract.
func NewArbFactory(address common.Address, backend bind.ContractBackend) (*ArbFactory, error) {
	contract, err := bindArbFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbFactory{ArbFactoryCaller: ArbFactoryCaller{contract: contract}, ArbFactoryTransactor: ArbFactoryTransactor{contract: contract}, ArbFactoryFilterer: ArbFactoryFilterer{contract: contract}}, nil
}

// NewArbFactoryCaller creates a new read-only instance of ArbFactory, bound to a specific deployed contract.
func NewArbFactoryCaller(address common.Address, caller bind.ContractCaller) (*ArbFactoryCaller, error) {
	contract, err := bindArbFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbFactoryCaller{contract: contract}, nil
}

// NewArbFactoryTransactor creates a new write-only instance of ArbFactory, bound to a specific deployed contract.
func NewArbFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbFactoryTransactor, error) {
	contract, err := bindArbFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbFactoryTransactor{contract: contract}, nil
}

// NewArbFactoryFilterer creates a new log filterer instance of ArbFactory, bound to a specific deployed contract.
func NewArbFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbFactoryFilterer, error) {
	contract, err := bindArbFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbFactoryFilterer{contract: contract}, nil
}

// bindArbFactory binds a generic wrapper to an already deployed contract.
func bindArbFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbFactory *ArbFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbFactory.Contract.ArbFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbFactory *ArbFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbFactory.Contract.ArbFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbFactory *ArbFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbFactory.Contract.ArbFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbFactory *ArbFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbFactory *ArbFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbFactory *ArbFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbFactory.Contract.contract.Transact(opts, method, params...)
}

// ChallengeFactoryAddress is a free data retrieval call binding the contract method 0x62e3c0b1.
//
// Solidity: function challengeFactoryAddress() view returns(address)
func (_ArbFactory *ArbFactoryCaller) ChallengeFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbFactory.contract.Call(opts, out, "challengeFactoryAddress")
	return *ret0, err
}

// ChallengeFactoryAddress is a free data retrieval call binding the contract method 0x62e3c0b1.
//
// Solidity: function challengeFactoryAddress() view returns(address)
func (_ArbFactory *ArbFactorySession) ChallengeFactoryAddress() (common.Address, error) {
	return _ArbFactory.Contract.ChallengeFactoryAddress(&_ArbFactory.CallOpts)
}

// ChallengeFactoryAddress is a free data retrieval call binding the contract method 0x62e3c0b1.
//
// Solidity: function challengeFactoryAddress() view returns(address)
func (_ArbFactory *ArbFactoryCallerSession) ChallengeFactoryAddress() (common.Address, error) {
	return _ArbFactory.Contract.ChallengeFactoryAddress(&_ArbFactory.CallOpts)
}

// GlobalInboxAddress is a free data retrieval call binding the contract method 0x582923c7.
//
// Solidity: function globalInboxAddress() view returns(address)
func (_ArbFactory *ArbFactoryCaller) GlobalInboxAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbFactory.contract.Call(opts, out, "globalInboxAddress")
	return *ret0, err
}

// GlobalInboxAddress is a free data retrieval call binding the contract method 0x582923c7.
//
// Solidity: function globalInboxAddress() view returns(address)
func (_ArbFactory *ArbFactorySession) GlobalInboxAddress() (common.Address, error) {
	return _ArbFactory.Contract.GlobalInboxAddress(&_ArbFactory.CallOpts)
}

// GlobalInboxAddress is a free data retrieval call binding the contract method 0x582923c7.
//
// Solidity: function globalInboxAddress() view returns(address)
func (_ArbFactory *ArbFactoryCallerSession) GlobalInboxAddress() (common.Address, error) {
	return _ArbFactory.Contract.GlobalInboxAddress(&_ArbFactory.CallOpts)
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_ArbFactory *ArbFactoryCaller) RollupTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbFactory.contract.Call(opts, out, "rollupTemplate")
	return *ret0, err
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_ArbFactory *ArbFactorySession) RollupTemplate() (common.Address, error) {
	return _ArbFactory.Contract.RollupTemplate(&_ArbFactory.CallOpts)
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_ArbFactory *ArbFactoryCallerSession) RollupTemplate() (common.Address, error) {
	return _ArbFactory.Contract.RollupTemplate(&_ArbFactory.CallOpts)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xf67e0439.
//
// Solidity: function createRollup(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _owner) returns()
func (_ArbFactory *ArbFactoryTransactor) CreateRollup(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ArbFactory.contract.Transact(opts, "createRollup", _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _owner)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xf67e0439.
//
// Solidity: function createRollup(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _owner) returns()
func (_ArbFactory *ArbFactorySession) CreateRollup(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ArbFactory.Contract.CreateRollup(&_ArbFactory.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _owner)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xf67e0439.
//
// Solidity: function createRollup(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _owner) returns()
func (_ArbFactory *ArbFactoryTransactorSession) CreateRollup(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ArbFactory.Contract.CreateRollup(&_ArbFactory.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _owner)
}

// ArbFactoryRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the ArbFactory contract.
type ArbFactoryRollupCreatedIterator struct {
	Event *ArbFactoryRollupCreated // Event containing the contract specifics and raw log

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
func (it *ArbFactoryRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbFactoryRollupCreated)
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
		it.Event = new(ArbFactoryRollupCreated)
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
func (it *ArbFactoryRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbFactoryRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbFactoryRollupCreated represents a RollupCreated event raised by the ArbFactory contract.
type ArbFactoryRollupCreated struct {
	VmAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address vmAddress)
func (_ArbFactory *ArbFactoryFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*ArbFactoryRollupCreatedIterator, error) {

	logs, sub, err := _ArbFactory.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &ArbFactoryRollupCreatedIterator{contract: _ArbFactory.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address vmAddress)
func (_ArbFactory *ArbFactoryFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *ArbFactoryRollupCreated) (event.Subscription, error) {

	logs, sub, err := _ArbFactory.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbFactoryRollupCreated)
				if err := _ArbFactory.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address vmAddress)
func (_ArbFactory *ArbFactoryFilterer) ParseRollupCreated(log types.Log) (*ArbFactoryRollupCreated, error) {
	event := new(ArbFactoryRollupCreated)
	if err := _ArbFactory.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CloneFactoryABI is the input ABI used to generate the binding from.
const CloneFactoryABI = "[]"

// CloneFactoryBin is the compiled bytecode used for deploying new contracts.
var CloneFactoryBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820796a132d4efb93bcf19bd98efef09e26e851946439e650891a87138b9b93c94f64736f6c63430005110032"

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

// IArbRollupABI is the input ABI used to generate the binding from.
const IArbRollupABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"forwardContractMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"_gracePeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_arbGasSpeedLimitPerTick\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_stakeRequirement\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbRollupFuncSigs maps the 4-byte function signature to its string representation.
var IArbRollupFuncSigs = map[string]string{
	"64ebfd03": "forwardContractMessage(address,uint256,uint256,bytes)",
	"19f4c3a7": "init(bytes32,uint128,uint128,uint64,uint128,address,address,address)",
}

// IArbRollup is an auto generated Go binding around an Ethereum contract.
type IArbRollup struct {
	IArbRollupCaller     // Read-only binding to the contract
	IArbRollupTransactor // Write-only binding to the contract
	IArbRollupFilterer   // Log filterer for contract events
}

// IArbRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbRollupSession struct {
	Contract     *IArbRollup       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbRollupCallerSession struct {
	Contract *IArbRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IArbRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbRollupTransactorSession struct {
	Contract     *IArbRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IArbRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbRollupRaw struct {
	Contract *IArbRollup // Generic contract binding to access the raw methods on
}

// IArbRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbRollupCallerRaw struct {
	Contract *IArbRollupCaller // Generic read-only contract binding to access the raw methods on
}

// IArbRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbRollupTransactorRaw struct {
	Contract *IArbRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbRollup creates a new instance of IArbRollup, bound to a specific deployed contract.
func NewIArbRollup(address common.Address, backend bind.ContractBackend) (*IArbRollup, error) {
	contract, err := bindIArbRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbRollup{IArbRollupCaller: IArbRollupCaller{contract: contract}, IArbRollupTransactor: IArbRollupTransactor{contract: contract}, IArbRollupFilterer: IArbRollupFilterer{contract: contract}}, nil
}

// NewIArbRollupCaller creates a new read-only instance of IArbRollup, bound to a specific deployed contract.
func NewIArbRollupCaller(address common.Address, caller bind.ContractCaller) (*IArbRollupCaller, error) {
	contract, err := bindIArbRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbRollupCaller{contract: contract}, nil
}

// NewIArbRollupTransactor creates a new write-only instance of IArbRollup, bound to a specific deployed contract.
func NewIArbRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbRollupTransactor, error) {
	contract, err := bindIArbRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbRollupTransactor{contract: contract}, nil
}

// NewIArbRollupFilterer creates a new log filterer instance of IArbRollup, bound to a specific deployed contract.
func NewIArbRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbRollupFilterer, error) {
	contract, err := bindIArbRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbRollupFilterer{contract: contract}, nil
}

// bindIArbRollup binds a generic wrapper to an already deployed contract.
func bindIArbRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbRollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbRollup *IArbRollupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbRollup.Contract.IArbRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbRollup *IArbRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbRollup.Contract.IArbRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbRollup *IArbRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbRollup.Contract.IArbRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbRollup *IArbRollupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbRollup *IArbRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbRollup *IArbRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbRollup.Contract.contract.Transact(opts, method, params...)
}

// ForwardContractMessage is a paid mutator transaction binding the contract method 0x64ebfd03.
//
// Solidity: function forwardContractMessage(address _sender, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns()
func (_IArbRollup *IArbRollupTransactor) ForwardContractMessage(opts *bind.TransactOpts, _sender common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _IArbRollup.contract.Transact(opts, "forwardContractMessage", _sender, _maxGas, _gasPriceBid, _data)
}

// ForwardContractMessage is a paid mutator transaction binding the contract method 0x64ebfd03.
//
// Solidity: function forwardContractMessage(address _sender, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns()
func (_IArbRollup *IArbRollupSession) ForwardContractMessage(_sender common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _IArbRollup.Contract.ForwardContractMessage(&_IArbRollup.TransactOpts, _sender, _maxGas, _gasPriceBid, _data)
}

// ForwardContractMessage is a paid mutator transaction binding the contract method 0x64ebfd03.
//
// Solidity: function forwardContractMessage(address _sender, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns()
func (_IArbRollup *IArbRollupTransactorSession) ForwardContractMessage(_sender common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _IArbRollup.Contract.ForwardContractMessage(&_IArbRollup.TransactOpts, _sender, _maxGas, _gasPriceBid, _data)
}

// Init is a paid mutator transaction binding the contract method 0x19f4c3a7.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_IArbRollup *IArbRollupTransactor) Init(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbRollup.contract.Transact(opts, "init", _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19f4c3a7.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_IArbRollup *IArbRollupSession) Init(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbRollup.Contract.Init(&_IArbRollup.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19f4c3a7.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_IArbRollup *IArbRollupTransactorSession) Init(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbRollup.Contract.Init(&_IArbRollup.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _owner, _challengeFactoryAddress, _globalInboxAddress)
}
