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
const ArbFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"_gracePeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_arbGasSpeedLimitPerTick\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint64[2]\",\"name\":\"_maxTimeBoundsWidth\",\"type\":\"uint64[2]\"},{\"internalType\":\"uint128\",\"name\":\"_stakeRequirement\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createRollup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInboxAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rollupTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ArbFactoryFuncSigs = map[string]string{
	"62e3c0b1": "challengeFactoryAddress()",
	"b10c5f84": "createRollup(bytes32,uint128,uint128,uint64,uint64[2],uint128,address)",
	"582923c7": "globalInboxAddress()",
	"8689d996": "rollupTemplate()",
}

// ArbFactoryBin is the compiled bytecode used for deploying new contracts.
var ArbFactoryBin = "0x608060405234801561001057600080fd5b5060405161053f38038061053f8339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b039384166001600160a01b03199182161790915560018054948416948216949094179093556002805492909116919092161790556104b08061008f6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063582923c71461005157806362e3c0b1146100755780638689d9961461007d578063b10c5f8414610085575b600080fd5b6100596100e9565b604080516001600160a01b039092168252519081900360200190f35b6100596100f8565b610059610107565b6100e7600480360361010081101561009c57600080fd5b5080359060208101356001600160801b03908116916040810135821691606082013567ffffffffffffffff1691608081019160c0820135169060e001356001600160a01b0316610116565b005b6001546001600160a01b031681565b6002546001600160a01b031681565b6000546001600160a01b031681565b6000805461012c906001600160a01b03166102a1565b60025460015460408051638e0f716760e01b8152600481018d81526001600160801b03808e1660248401528c16604483015267ffffffffffffffff8b1660648301529495506001600160a01b0380871695638e0f7167958f958f958f958f958f958f958f95908316949216929091608490910190879080828437600081840152601f19601f820116905080830192505050856001600160801b03166001600160801b03168152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b03168152602001826001600160a01b03166001600160a01b031681526020019950505050505050505050600060405180830381600087803b15801561024257600080fd5b505af1158015610256573d6000803e3d6000fd5b5050604080516001600160a01b038516815290517f84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c9350908190036020019150a15050505050505050565b60006060604051806020016102b5906103be565b601f1982820381018352601f9091011660408181526001600160a01b038616602083810191909152815180840382018152828401909252835191926060019182918501908083835b6020831061031c5780518252601f1990920191602091820191016102fd565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106103645780518252601f199092019160209182019101610345565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050806020018151808234f09350836103b6573d6000803e3d6000fd5b505050919050565b60b1806103cb8339019056fe6080604052348015600f57600080fd5b5060405160b138038060b183398181016040526020811015602f57600080fd5b5051604080517f363d3d373d3d3d363d73000000000000000000000000000000000000000000006020828101919091526001600160601b0319606085901b16602a8301527f5af43d82803e903d91602b57fd5bf30000000000000000000000000000000000603e8301528251602d81840381018252604d9093019093528201f3fea265627a7a72315820383cfb78c66b49966c5aa0b2ee83fe6d6cd5485f06fe388ad088b70cd0d7969364736f6c634300050d0032"

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

// CreateRollup is a paid mutator transaction binding the contract method 0xb10c5f84.
//
// Solidity: function createRollup(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint64[2] _maxTimeBoundsWidth, uint128 _stakeRequirement, address _owner) returns()
func (_ArbFactory *ArbFactoryTransactor) CreateRollup(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _maxTimeBoundsWidth [2]uint64, _stakeRequirement *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ArbFactory.contract.Transact(opts, "createRollup", _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _maxTimeBoundsWidth, _stakeRequirement, _owner)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xb10c5f84.
//
// Solidity: function createRollup(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint64[2] _maxTimeBoundsWidth, uint128 _stakeRequirement, address _owner) returns()
func (_ArbFactory *ArbFactorySession) CreateRollup(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _maxTimeBoundsWidth [2]uint64, _stakeRequirement *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ArbFactory.Contract.CreateRollup(&_ArbFactory.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _maxTimeBoundsWidth, _stakeRequirement, _owner)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xb10c5f84.
//
// Solidity: function createRollup(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint64[2] _maxTimeBoundsWidth, uint128 _stakeRequirement, address _owner) returns()
func (_ArbFactory *ArbFactoryTransactorSession) CreateRollup(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _maxTimeBoundsWidth [2]uint64, _stakeRequirement *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ArbFactory.Contract.CreateRollup(&_ArbFactory.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _maxTimeBoundsWidth, _stakeRequirement, _owner)
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
var CloneFactoryBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723158209020939b4d87199d2051d82c5b9095e1e442b3fcff4b8e5ab550a15cd5b570e964736f6c634300050d0032"

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
const IArbRollupABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"forwardContractMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"_gracePeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_arbGasSpeedLimitPerTick\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint64[2]\",\"name\":\"_maxTimeBoundsWidth\",\"type\":\"uint64[2]\"},{\"internalType\":\"uint128\",\"name\":\"_stakeRequirement\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbRollupFuncSigs maps the 4-byte function signature to its string representation.
var IArbRollupFuncSigs = map[string]string{
	"c5818317": "forwardContractMessage(address,bytes)",
	"8e0f7167": "init(bytes32,uint128,uint128,uint64,uint64[2],uint128,address,address,address)",
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

// ForwardContractMessage is a paid mutator transaction binding the contract method 0xc5818317.
//
// Solidity: function forwardContractMessage(address _sender, bytes _data) payable returns()
func (_IArbRollup *IArbRollupTransactor) ForwardContractMessage(opts *bind.TransactOpts, _sender common.Address, _data []byte) (*types.Transaction, error) {
	return _IArbRollup.contract.Transact(opts, "forwardContractMessage", _sender, _data)
}

// ForwardContractMessage is a paid mutator transaction binding the contract method 0xc5818317.
//
// Solidity: function forwardContractMessage(address _sender, bytes _data) payable returns()
func (_IArbRollup *IArbRollupSession) ForwardContractMessage(_sender common.Address, _data []byte) (*types.Transaction, error) {
	return _IArbRollup.Contract.ForwardContractMessage(&_IArbRollup.TransactOpts, _sender, _data)
}

// ForwardContractMessage is a paid mutator transaction binding the contract method 0xc5818317.
//
// Solidity: function forwardContractMessage(address _sender, bytes _data) payable returns()
func (_IArbRollup *IArbRollupTransactorSession) ForwardContractMessage(_sender common.Address, _data []byte) (*types.Transaction, error) {
	return _IArbRollup.Contract.ForwardContractMessage(&_IArbRollup.TransactOpts, _sender, _data)
}

// Init is a paid mutator transaction binding the contract method 0x8e0f7167.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint64[2] _maxTimeBoundsWidth, uint128 _stakeRequirement, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_IArbRollup *IArbRollupTransactor) Init(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _maxTimeBoundsWidth [2]uint64, _stakeRequirement *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbRollup.contract.Transact(opts, "init", _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _maxTimeBoundsWidth, _stakeRequirement, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x8e0f7167.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint64[2] _maxTimeBoundsWidth, uint128 _stakeRequirement, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_IArbRollup *IArbRollupSession) Init(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _maxTimeBoundsWidth [2]uint64, _stakeRequirement *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbRollup.Contract.Init(&_IArbRollup.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _maxTimeBoundsWidth, _stakeRequirement, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x8e0f7167.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint64[2] _maxTimeBoundsWidth, uint128 _stakeRequirement, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_IArbRollup *IArbRollupTransactorSession) Init(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _maxTimeBoundsWidth [2]uint64, _stakeRequirement *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbRollup.Contract.Init(&_IArbRollup.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _maxTimeBoundsWidth, _stakeRequirement, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// SpawnABI is the input ABI used to generate the binding from.
const SpawnABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"logicContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SpawnBin is the compiled bytecode used for deploying new contracts.
var SpawnBin = "0x6080604052348015600f57600080fd5b5060405160b138038060b183398181016040526020811015602f57600080fd5b5051604080517f363d3d373d3d3d363d73000000000000000000000000000000000000000000006020828101919091526001600160601b0319606085901b16602a8301527f5af43d82803e903d91602b57fd5bf30000000000000000000000000000000000603e8301528251602d81840381018252604d9093019093528201f3fe"

// DeploySpawn deploys a new Ethereum contract, binding an instance of Spawn to it.
func DeploySpawn(auth *bind.TransactOpts, backend bind.ContractBackend, logicContract common.Address) (common.Address, *types.Transaction, *Spawn, error) {
	parsed, err := abi.JSON(strings.NewReader(SpawnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SpawnBin), backend, logicContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Spawn{SpawnCaller: SpawnCaller{contract: contract}, SpawnTransactor: SpawnTransactor{contract: contract}, SpawnFilterer: SpawnFilterer{contract: contract}}, nil
}

// Spawn is an auto generated Go binding around an Ethereum contract.
type Spawn struct {
	SpawnCaller     // Read-only binding to the contract
	SpawnTransactor // Write-only binding to the contract
	SpawnFilterer   // Log filterer for contract events
}

// SpawnCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpawnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpawnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpawnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpawnSession struct {
	Contract     *Spawn            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpawnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpawnCallerSession struct {
	Contract *SpawnCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SpawnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpawnTransactorSession struct {
	Contract     *SpawnTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpawnRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpawnRaw struct {
	Contract *Spawn // Generic contract binding to access the raw methods on
}

// SpawnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpawnCallerRaw struct {
	Contract *SpawnCaller // Generic read-only contract binding to access the raw methods on
}

// SpawnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpawnTransactorRaw struct {
	Contract *SpawnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpawn creates a new instance of Spawn, bound to a specific deployed contract.
func NewSpawn(address common.Address, backend bind.ContractBackend) (*Spawn, error) {
	contract, err := bindSpawn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spawn{SpawnCaller: SpawnCaller{contract: contract}, SpawnTransactor: SpawnTransactor{contract: contract}, SpawnFilterer: SpawnFilterer{contract: contract}}, nil
}

// NewSpawnCaller creates a new read-only instance of Spawn, bound to a specific deployed contract.
func NewSpawnCaller(address common.Address, caller bind.ContractCaller) (*SpawnCaller, error) {
	contract, err := bindSpawn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpawnCaller{contract: contract}, nil
}

// NewSpawnTransactor creates a new write-only instance of Spawn, bound to a specific deployed contract.
func NewSpawnTransactor(address common.Address, transactor bind.ContractTransactor) (*SpawnTransactor, error) {
	contract, err := bindSpawn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpawnTransactor{contract: contract}, nil
}

// NewSpawnFilterer creates a new log filterer instance of Spawn, bound to a specific deployed contract.
func NewSpawnFilterer(address common.Address, filterer bind.ContractFilterer) (*SpawnFilterer, error) {
	contract, err := bindSpawn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpawnFilterer{contract: contract}, nil
}

// bindSpawn binds a generic wrapper to an already deployed contract.
func bindSpawn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SpawnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spawn *SpawnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spawn.Contract.SpawnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spawn *SpawnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.Contract.SpawnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spawn *SpawnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spawn.Contract.SpawnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spawn *SpawnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spawn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spawn *SpawnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spawn *SpawnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spawn.Contract.contract.Transact(opts, method, params...)
}
