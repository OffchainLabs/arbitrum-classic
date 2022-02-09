// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

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

// RollupMockMetaData contains all meta data concerning the RollupMock contract.
var RollupMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"sequencerInboxMaxDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInboxMaxDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sequencerInboxMaxDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerInboxMaxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"setMock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060ce8061001f6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c806314828f9214604157806327ce1def146059578063addd678414607b575b600080fd5b60476081565b60408051918252519081900360200190f35b607960048036036040811015606d57600080fd5b50803590602001356087565b005b60476092565b60005481565b600091909155600155565b6001548156fea2646970667358221220c0af72a94be3f5845409886d45c545eab74fd73dbf89c874a5c9fd3f250f27e164736f6c634300060b0033",
}

// RollupMockABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMockMetaData.ABI instead.
var RollupMockABI = RollupMockMetaData.ABI

// RollupMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupMockMetaData.Bin instead.
var RollupMockBin = RollupMockMetaData.Bin

// DeployRollupMock deploys a new Ethereum contract, binding an instance of RollupMock to it.
func DeployRollupMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupMock, error) {
	parsed, err := RollupMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupMock{RollupMockCaller: RollupMockCaller{contract: contract}, RollupMockTransactor: RollupMockTransactor{contract: contract}, RollupMockFilterer: RollupMockFilterer{contract: contract}}, nil
}

// RollupMock is an auto generated Go binding around an Ethereum contract.
type RollupMock struct {
	RollupMockCaller     // Read-only binding to the contract
	RollupMockTransactor // Write-only binding to the contract
	RollupMockFilterer   // Log filterer for contract events
}

// RollupMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupMockSession struct {
	Contract     *RollupMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupMockCallerSession struct {
	Contract *RollupMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupMockTransactorSession struct {
	Contract     *RollupMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupMockRaw struct {
	Contract *RollupMock // Generic contract binding to access the raw methods on
}

// RollupMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupMockCallerRaw struct {
	Contract *RollupMockCaller // Generic read-only contract binding to access the raw methods on
}

// RollupMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupMockTransactorRaw struct {
	Contract *RollupMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupMock creates a new instance of RollupMock, bound to a specific deployed contract.
func NewRollupMock(address common.Address, backend bind.ContractBackend) (*RollupMock, error) {
	contract, err := bindRollupMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupMock{RollupMockCaller: RollupMockCaller{contract: contract}, RollupMockTransactor: RollupMockTransactor{contract: contract}, RollupMockFilterer: RollupMockFilterer{contract: contract}}, nil
}

// NewRollupMockCaller creates a new read-only instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockCaller(address common.Address, caller bind.ContractCaller) (*RollupMockCaller, error) {
	contract, err := bindRollupMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMockCaller{contract: contract}, nil
}

// NewRollupMockTransactor creates a new write-only instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupMockTransactor, error) {
	contract, err := bindRollupMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMockTransactor{contract: contract}, nil
}

// NewRollupMockFilterer creates a new log filterer instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupMockFilterer, error) {
	contract, err := bindRollupMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupMockFilterer{contract: contract}, nil
}

// bindRollupMock binds a generic wrapper to an already deployed contract.
func bindRollupMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMock *RollupMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupMock.Contract.RollupMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMock *RollupMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.Contract.RollupMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMock *RollupMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMock.Contract.RollupMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMock *RollupMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMock *RollupMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMock *RollupMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMock.Contract.contract.Transact(opts, method, params...)
}

// SequencerInboxMaxDelayBlocks is a free data retrieval call binding the contract method 0x14828f92.
//
// Solidity: function sequencerInboxMaxDelayBlocks() view returns(uint256)
func (_RollupMock *RollupMockCaller) SequencerInboxMaxDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupMock.contract.Call(opts, &out, "sequencerInboxMaxDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerInboxMaxDelayBlocks is a free data retrieval call binding the contract method 0x14828f92.
//
// Solidity: function sequencerInboxMaxDelayBlocks() view returns(uint256)
func (_RollupMock *RollupMockSession) SequencerInboxMaxDelayBlocks() (*big.Int, error) {
	return _RollupMock.Contract.SequencerInboxMaxDelayBlocks(&_RollupMock.CallOpts)
}

// SequencerInboxMaxDelayBlocks is a free data retrieval call binding the contract method 0x14828f92.
//
// Solidity: function sequencerInboxMaxDelayBlocks() view returns(uint256)
func (_RollupMock *RollupMockCallerSession) SequencerInboxMaxDelayBlocks() (*big.Int, error) {
	return _RollupMock.Contract.SequencerInboxMaxDelayBlocks(&_RollupMock.CallOpts)
}

// SequencerInboxMaxDelaySeconds is a free data retrieval call binding the contract method 0xaddd6784.
//
// Solidity: function sequencerInboxMaxDelaySeconds() view returns(uint256)
func (_RollupMock *RollupMockCaller) SequencerInboxMaxDelaySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupMock.contract.Call(opts, &out, "sequencerInboxMaxDelaySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerInboxMaxDelaySeconds is a free data retrieval call binding the contract method 0xaddd6784.
//
// Solidity: function sequencerInboxMaxDelaySeconds() view returns(uint256)
func (_RollupMock *RollupMockSession) SequencerInboxMaxDelaySeconds() (*big.Int, error) {
	return _RollupMock.Contract.SequencerInboxMaxDelaySeconds(&_RollupMock.CallOpts)
}

// SequencerInboxMaxDelaySeconds is a free data retrieval call binding the contract method 0xaddd6784.
//
// Solidity: function sequencerInboxMaxDelaySeconds() view returns(uint256)
func (_RollupMock *RollupMockCallerSession) SequencerInboxMaxDelaySeconds() (*big.Int, error) {
	return _RollupMock.Contract.SequencerInboxMaxDelaySeconds(&_RollupMock.CallOpts)
}

// SetMock is a paid mutator transaction binding the contract method 0x27ce1def.
//
// Solidity: function setMock(uint256 _sequencerInboxMaxDelayBlocks, uint256 _sequencerInboxMaxDelaySeconds) returns()
func (_RollupMock *RollupMockTransactor) SetMock(opts *bind.TransactOpts, _sequencerInboxMaxDelayBlocks *big.Int, _sequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupMock.contract.Transact(opts, "setMock", _sequencerInboxMaxDelayBlocks, _sequencerInboxMaxDelaySeconds)
}

// SetMock is a paid mutator transaction binding the contract method 0x27ce1def.
//
// Solidity: function setMock(uint256 _sequencerInboxMaxDelayBlocks, uint256 _sequencerInboxMaxDelaySeconds) returns()
func (_RollupMock *RollupMockSession) SetMock(_sequencerInboxMaxDelayBlocks *big.Int, _sequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupMock.Contract.SetMock(&_RollupMock.TransactOpts, _sequencerInboxMaxDelayBlocks, _sequencerInboxMaxDelaySeconds)
}

// SetMock is a paid mutator transaction binding the contract method 0x27ce1def.
//
// Solidity: function setMock(uint256 _sequencerInboxMaxDelayBlocks, uint256 _sequencerInboxMaxDelaySeconds) returns()
func (_RollupMock *RollupMockTransactorSession) SetMock(_sequencerInboxMaxDelayBlocks *big.Int, _sequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupMock.Contract.SetMock(&_RollupMock.TransactOpts, _sequencerInboxMaxDelayBlocks, _sequencerInboxMaxDelaySeconds)
}
