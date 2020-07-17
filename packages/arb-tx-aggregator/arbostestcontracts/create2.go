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

// CloneFactoryABI is the input ABI used to generate the binding from.
const CloneFactoryABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"create2Clone\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CloneFactoryFuncSigs maps the 4-byte function signature to its string representation.
var CloneFactoryFuncSigs = map[string]string{
	"c91091c3": "create2Clone(address,uint256)",
}

// CloneFactoryBin is the compiled bytecode used for deploying new contracts.
var CloneFactoryBin = "0x608060405234801561001057600080fd5b5060fb8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c91091c314602d575b600080fd5b605660048036036040811015604157600080fd5b506001600160a01b0381351690602001356072565b604080516001600160a01b039092168252519081900360200190f35b6000808360601b9050604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528160148201526e5af43d82803e903d91602b57fd5bf360881b6028820152836037826000f59594505050505056fea265627a7a723158207d794861e38a1bb610c376aeb351b67396537afd34b87cac965958a4b858f31c64736f6c63430005110032"

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

// Create2Clone is a paid mutator transaction binding the contract method 0xc91091c3.
//
// Solidity: function create2Clone(address target, uint256 salt) returns(address result)
func (_CloneFactory *CloneFactoryTransactor) Create2Clone(opts *bind.TransactOpts, target common.Address, salt *big.Int) (*types.Transaction, error) {
	return _CloneFactory.contract.Transact(opts, "create2Clone", target, salt)
}

// Create2Clone is a paid mutator transaction binding the contract method 0xc91091c3.
//
// Solidity: function create2Clone(address target, uint256 salt) returns(address result)
func (_CloneFactory *CloneFactorySession) Create2Clone(target common.Address, salt *big.Int) (*types.Transaction, error) {
	return _CloneFactory.Contract.Create2Clone(&_CloneFactory.TransactOpts, target, salt)
}

// Create2Clone is a paid mutator transaction binding the contract method 0xc91091c3.
//
// Solidity: function create2Clone(address target, uint256 salt) returns(address result)
func (_CloneFactory *CloneFactoryTransactorSession) Create2Clone(target common.Address, salt *big.Int) (*types.Transaction, error) {
	return _CloneFactory.Contract.Create2Clone(&_CloneFactory.TransactOpts, target, salt)
}
