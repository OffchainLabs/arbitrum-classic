// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

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

// RollupTesterABI is the input ABI used to generate the binding from.
const RollupTesterABI = "[]"

// RollupTesterBin is the compiled bytecode used for deploying new contracts.
var RollupTesterBin = "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220fa6e1617c92ac3b6d5b1f8d1adc4fb7d15298716635b0b7ec1231a8b6e7ed46764736f6c634300060c0033"

// DeployRollupTester deploys a new Ethereum contract, binding an instance of RollupTester to it.
func DeployRollupTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupTester, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupTester{RollupTesterCaller: RollupTesterCaller{contract: contract}, RollupTesterTransactor: RollupTesterTransactor{contract: contract}, RollupTesterFilterer: RollupTesterFilterer{contract: contract}}, nil
}

// RollupTester is an auto generated Go binding around an Ethereum contract.
type RollupTester struct {
	RollupTesterCaller     // Read-only binding to the contract
	RollupTesterTransactor // Write-only binding to the contract
	RollupTesterFilterer   // Log filterer for contract events
}

// RollupTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupTesterSession struct {
	Contract     *RollupTester     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupTesterCallerSession struct {
	Contract *RollupTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RollupTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTesterTransactorSession struct {
	Contract     *RollupTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RollupTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupTesterRaw struct {
	Contract *RollupTester // Generic contract binding to access the raw methods on
}

// RollupTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupTesterCallerRaw struct {
	Contract *RollupTesterCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTesterTransactorRaw struct {
	Contract *RollupTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupTester creates a new instance of RollupTester, bound to a specific deployed contract.
func NewRollupTester(address common.Address, backend bind.ContractBackend) (*RollupTester, error) {
	contract, err := bindRollupTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupTester{RollupTesterCaller: RollupTesterCaller{contract: contract}, RollupTesterTransactor: RollupTesterTransactor{contract: contract}, RollupTesterFilterer: RollupTesterFilterer{contract: contract}}, nil
}

// NewRollupTesterCaller creates a new read-only instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterCaller(address common.Address, caller bind.ContractCaller) (*RollupTesterCaller, error) {
	contract, err := bindRollupTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTesterCaller{contract: contract}, nil
}

// NewRollupTesterTransactor creates a new write-only instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTesterTransactor, error) {
	contract, err := bindRollupTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTesterTransactor{contract: contract}, nil
}

// NewRollupTesterFilterer creates a new log filterer instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupTesterFilterer, error) {
	contract, err := bindRollupTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupTesterFilterer{contract: contract}, nil
}

// bindRollupTester binds a generic wrapper to an already deployed contract.
func bindRollupTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTester *RollupTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupTester.Contract.RollupTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTester *RollupTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTester.Contract.RollupTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTester *RollupTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTester.Contract.RollupTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTester *RollupTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTester *RollupTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTester *RollupTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTester.Contract.contract.Transact(opts, method, params...)
}
