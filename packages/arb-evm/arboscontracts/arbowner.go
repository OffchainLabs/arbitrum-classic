// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arboscontracts

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

// ArbOwnerABI is the input ABI used to generate the binding from.
const ArbOwnerABI = "[{\"inputs\":[],\"name\":\"addToReserveFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pluggableId\",\"type\":\"uint256\"}],\"name\":\"bindAddressToPluggable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"changeSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"marshalledCode\",\"type\":\"bytes\"}],\"name\":\"continueCodeUpload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finishCodeUploadAsArbosUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"keepState\",\"type\":\"bool\"}],\"name\":\"finishCodeUploadAsPluggable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeMaxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwnerAddr\",\"type\":\"address\"}],\"name\":\"giveOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksPerSend\",\"type\":\"uint256\"}],\"name\":\"setBlocksPerSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom2\",\"type\":\"uint256\"}],\"name\":\"setFeeMaxes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom2\",\"type\":\"uint256\"}],\"name\":\"setFeeRates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startCodeUpload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbOwnerFuncSigs maps the 4-byte function signature to its string representation.
var ArbOwnerFuncSigs = map[string]string{
	"c3bf429d": "addToReserveFunds()",
	"f589445c": "bindAddressToPluggable(address,uint256)",
	"ea180a38": "changeSequencer(address,uint256,uint256)",
	"56331f75": "continueCodeUpload(bytes)",
	"fbb53a17": "finishCodeUploadAsArbosUpgrade()",
	"f4f4e136": "finishCodeUploadAsPluggable(uint256,bool)",
	"23e089dd": "getFeeMaxes()",
	"d6e7a55e": "getFeeRates()",
	"4ccb20c0": "getFeeRecipient()",
	"e3a0a148": "giveOwnership(address)",
	"340e4fc2": "setBlocksPerSend(uint256)",
	"72861aa4": "setFeeMaxes(uint256,uint256,uint256,uint256)",
	"30d2361e": "setFeeRates(uint256,uint256,uint256,uint256)",
	"e74b981b": "setFeeRecipient(address)",
	"31acdf5e": "startCodeUpload()",
}

// ArbOwner is an auto generated Go binding around an Ethereum contract.
type ArbOwner struct {
	ArbOwnerCaller     // Read-only binding to the contract
	ArbOwnerTransactor // Write-only binding to the contract
	ArbOwnerFilterer   // Log filterer for contract events
}

// ArbOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbOwnerSession struct {
	Contract     *ArbOwner         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbOwnerCallerSession struct {
	Contract *ArbOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbOwnerTransactorSession struct {
	Contract     *ArbOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbOwnerRaw struct {
	Contract *ArbOwner // Generic contract binding to access the raw methods on
}

// ArbOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbOwnerCallerRaw struct {
	Contract *ArbOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// ArbOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbOwnerTransactorRaw struct {
	Contract *ArbOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbOwner creates a new instance of ArbOwner, bound to a specific deployed contract.
func NewArbOwner(address common.Address, backend bind.ContractBackend) (*ArbOwner, error) {
	contract, err := bindArbOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbOwner{ArbOwnerCaller: ArbOwnerCaller{contract: contract}, ArbOwnerTransactor: ArbOwnerTransactor{contract: contract}, ArbOwnerFilterer: ArbOwnerFilterer{contract: contract}}, nil
}

// NewArbOwnerCaller creates a new read-only instance of ArbOwner, bound to a specific deployed contract.
func NewArbOwnerCaller(address common.Address, caller bind.ContractCaller) (*ArbOwnerCaller, error) {
	contract, err := bindArbOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbOwnerCaller{contract: contract}, nil
}

// NewArbOwnerTransactor creates a new write-only instance of ArbOwner, bound to a specific deployed contract.
func NewArbOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbOwnerTransactor, error) {
	contract, err := bindArbOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbOwnerTransactor{contract: contract}, nil
}

// NewArbOwnerFilterer creates a new log filterer instance of ArbOwner, bound to a specific deployed contract.
func NewArbOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbOwnerFilterer, error) {
	contract, err := bindArbOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbOwnerFilterer{contract: contract}, nil
}

// bindArbOwner binds a generic wrapper to an already deployed contract.
func bindArbOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbOwner *ArbOwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbOwner.Contract.ArbOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbOwner *ArbOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.Contract.ArbOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbOwner *ArbOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbOwner.Contract.ArbOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbOwner *ArbOwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbOwner *ArbOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbOwner *ArbOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbOwner.Contract.contract.Transact(opts, method, params...)
}

// GetFeeMaxes is a free data retrieval call binding the contract method 0x23e089dd.
//
// Solidity: function getFeeMaxes() view returns(uint256, uint256, uint256, uint256)
func (_ArbOwner *ArbOwnerCaller) GetFeeMaxes(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getFeeMaxes")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetFeeMaxes is a free data retrieval call binding the contract method 0x23e089dd.
//
// Solidity: function getFeeMaxes() view returns(uint256, uint256, uint256, uint256)
func (_ArbOwner *ArbOwnerSession) GetFeeMaxes() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbOwner.Contract.GetFeeMaxes(&_ArbOwner.CallOpts)
}

// GetFeeMaxes is a free data retrieval call binding the contract method 0x23e089dd.
//
// Solidity: function getFeeMaxes() view returns(uint256, uint256, uint256, uint256)
func (_ArbOwner *ArbOwnerCallerSession) GetFeeMaxes() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbOwner.Contract.GetFeeMaxes(&_ArbOwner.CallOpts)
}

// GetFeeRates is a free data retrieval call binding the contract method 0xd6e7a55e.
//
// Solidity: function getFeeRates() view returns(uint256, uint256, uint256, uint256)
func (_ArbOwner *ArbOwnerCaller) GetFeeRates(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getFeeRates")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetFeeRates is a free data retrieval call binding the contract method 0xd6e7a55e.
//
// Solidity: function getFeeRates() view returns(uint256, uint256, uint256, uint256)
func (_ArbOwner *ArbOwnerSession) GetFeeRates() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbOwner.Contract.GetFeeRates(&_ArbOwner.CallOpts)
}

// GetFeeRates is a free data retrieval call binding the contract method 0xd6e7a55e.
//
// Solidity: function getFeeRates() view returns(uint256, uint256, uint256, uint256)
func (_ArbOwner *ArbOwnerCallerSession) GetFeeRates() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ArbOwner.Contract.GetFeeRates(&_ArbOwner.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_ArbOwner *ArbOwnerCaller) GetFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_ArbOwner *ArbOwnerSession) GetFeeRecipient() (common.Address, error) {
	return _ArbOwner.Contract.GetFeeRecipient(&_ArbOwner.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_ArbOwner *ArbOwnerCallerSession) GetFeeRecipient() (common.Address, error) {
	return _ArbOwner.Contract.GetFeeRecipient(&_ArbOwner.CallOpts)
}

// AddToReserveFunds is a paid mutator transaction binding the contract method 0xc3bf429d.
//
// Solidity: function addToReserveFunds() payable returns()
func (_ArbOwner *ArbOwnerTransactor) AddToReserveFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "addToReserveFunds")
}

// AddToReserveFunds is a paid mutator transaction binding the contract method 0xc3bf429d.
//
// Solidity: function addToReserveFunds() payable returns()
func (_ArbOwner *ArbOwnerSession) AddToReserveFunds() (*types.Transaction, error) {
	return _ArbOwner.Contract.AddToReserveFunds(&_ArbOwner.TransactOpts)
}

// AddToReserveFunds is a paid mutator transaction binding the contract method 0xc3bf429d.
//
// Solidity: function addToReserveFunds() payable returns()
func (_ArbOwner *ArbOwnerTransactorSession) AddToReserveFunds() (*types.Transaction, error) {
	return _ArbOwner.Contract.AddToReserveFunds(&_ArbOwner.TransactOpts)
}

// BindAddressToPluggable is a paid mutator transaction binding the contract method 0xf589445c.
//
// Solidity: function bindAddressToPluggable(address addr, uint256 pluggableId) returns()
func (_ArbOwner *ArbOwnerTransactor) BindAddressToPluggable(opts *bind.TransactOpts, addr common.Address, pluggableId *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "bindAddressToPluggable", addr, pluggableId)
}

// BindAddressToPluggable is a paid mutator transaction binding the contract method 0xf589445c.
//
// Solidity: function bindAddressToPluggable(address addr, uint256 pluggableId) returns()
func (_ArbOwner *ArbOwnerSession) BindAddressToPluggable(addr common.Address, pluggableId *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.BindAddressToPluggable(&_ArbOwner.TransactOpts, addr, pluggableId)
}

// BindAddressToPluggable is a paid mutator transaction binding the contract method 0xf589445c.
//
// Solidity: function bindAddressToPluggable(address addr, uint256 pluggableId) returns()
func (_ArbOwner *ArbOwnerTransactorSession) BindAddressToPluggable(addr common.Address, pluggableId *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.BindAddressToPluggable(&_ArbOwner.TransactOpts, addr, pluggableId)
}

// ChangeSequencer is a paid mutator transaction binding the contract method 0xea180a38.
//
// Solidity: function changeSequencer(address sequencerAddr, uint256 maxDelayBlocks, uint256 maxDelaySeconds) returns()
func (_ArbOwner *ArbOwnerTransactor) ChangeSequencer(opts *bind.TransactOpts, sequencerAddr common.Address, maxDelayBlocks *big.Int, maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "changeSequencer", sequencerAddr, maxDelayBlocks, maxDelaySeconds)
}

// ChangeSequencer is a paid mutator transaction binding the contract method 0xea180a38.
//
// Solidity: function changeSequencer(address sequencerAddr, uint256 maxDelayBlocks, uint256 maxDelaySeconds) returns()
func (_ArbOwner *ArbOwnerSession) ChangeSequencer(sequencerAddr common.Address, maxDelayBlocks *big.Int, maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.ChangeSequencer(&_ArbOwner.TransactOpts, sequencerAddr, maxDelayBlocks, maxDelaySeconds)
}

// ChangeSequencer is a paid mutator transaction binding the contract method 0xea180a38.
//
// Solidity: function changeSequencer(address sequencerAddr, uint256 maxDelayBlocks, uint256 maxDelaySeconds) returns()
func (_ArbOwner *ArbOwnerTransactorSession) ChangeSequencer(sequencerAddr common.Address, maxDelayBlocks *big.Int, maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.ChangeSequencer(&_ArbOwner.TransactOpts, sequencerAddr, maxDelayBlocks, maxDelaySeconds)
}

// ContinueCodeUpload is a paid mutator transaction binding the contract method 0x56331f75.
//
// Solidity: function continueCodeUpload(bytes marshalledCode) returns()
func (_ArbOwner *ArbOwnerTransactor) ContinueCodeUpload(opts *bind.TransactOpts, marshalledCode []byte) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "continueCodeUpload", marshalledCode)
}

// ContinueCodeUpload is a paid mutator transaction binding the contract method 0x56331f75.
//
// Solidity: function continueCodeUpload(bytes marshalledCode) returns()
func (_ArbOwner *ArbOwnerSession) ContinueCodeUpload(marshalledCode []byte) (*types.Transaction, error) {
	return _ArbOwner.Contract.ContinueCodeUpload(&_ArbOwner.TransactOpts, marshalledCode)
}

// ContinueCodeUpload is a paid mutator transaction binding the contract method 0x56331f75.
//
// Solidity: function continueCodeUpload(bytes marshalledCode) returns()
func (_ArbOwner *ArbOwnerTransactorSession) ContinueCodeUpload(marshalledCode []byte) (*types.Transaction, error) {
	return _ArbOwner.Contract.ContinueCodeUpload(&_ArbOwner.TransactOpts, marshalledCode)
}

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0xfbb53a17.
//
// Solidity: function finishCodeUploadAsArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerTransactor) FinishCodeUploadAsArbosUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "finishCodeUploadAsArbosUpgrade")
}

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0xfbb53a17.
//
// Solidity: function finishCodeUploadAsArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerSession) FinishCodeUploadAsArbosUpgrade() (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishCodeUploadAsArbosUpgrade(&_ArbOwner.TransactOpts)
}

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0xfbb53a17.
//
// Solidity: function finishCodeUploadAsArbosUpgrade() returns()
func (_ArbOwner *ArbOwnerTransactorSession) FinishCodeUploadAsArbosUpgrade() (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishCodeUploadAsArbosUpgrade(&_ArbOwner.TransactOpts)
}

// FinishCodeUploadAsPluggable is a paid mutator transaction binding the contract method 0xf4f4e136.
//
// Solidity: function finishCodeUploadAsPluggable(uint256 id, bool keepState) returns()
func (_ArbOwner *ArbOwnerTransactor) FinishCodeUploadAsPluggable(opts *bind.TransactOpts, id *big.Int, keepState bool) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "finishCodeUploadAsPluggable", id, keepState)
}

// FinishCodeUploadAsPluggable is a paid mutator transaction binding the contract method 0xf4f4e136.
//
// Solidity: function finishCodeUploadAsPluggable(uint256 id, bool keepState) returns()
func (_ArbOwner *ArbOwnerSession) FinishCodeUploadAsPluggable(id *big.Int, keepState bool) (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishCodeUploadAsPluggable(&_ArbOwner.TransactOpts, id, keepState)
}

// FinishCodeUploadAsPluggable is a paid mutator transaction binding the contract method 0xf4f4e136.
//
// Solidity: function finishCodeUploadAsPluggable(uint256 id, bool keepState) returns()
func (_ArbOwner *ArbOwnerTransactorSession) FinishCodeUploadAsPluggable(id *big.Int, keepState bool) (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishCodeUploadAsPluggable(&_ArbOwner.TransactOpts, id, keepState)
}

// GiveOwnership is a paid mutator transaction binding the contract method 0xe3a0a148.
//
// Solidity: function giveOwnership(address newOwnerAddr) returns()
func (_ArbOwner *ArbOwnerTransactor) GiveOwnership(opts *bind.TransactOpts, newOwnerAddr common.Address) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "giveOwnership", newOwnerAddr)
}

// GiveOwnership is a paid mutator transaction binding the contract method 0xe3a0a148.
//
// Solidity: function giveOwnership(address newOwnerAddr) returns()
func (_ArbOwner *ArbOwnerSession) GiveOwnership(newOwnerAddr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.GiveOwnership(&_ArbOwner.TransactOpts, newOwnerAddr)
}

// GiveOwnership is a paid mutator transaction binding the contract method 0xe3a0a148.
//
// Solidity: function giveOwnership(address newOwnerAddr) returns()
func (_ArbOwner *ArbOwnerTransactorSession) GiveOwnership(newOwnerAddr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.GiveOwnership(&_ArbOwner.TransactOpts, newOwnerAddr)
}

// SetBlocksPerSend is a paid mutator transaction binding the contract method 0x340e4fc2.
//
// Solidity: function setBlocksPerSend(uint256 blocksPerSend) returns()
func (_ArbOwner *ArbOwnerTransactor) SetBlocksPerSend(opts *bind.TransactOpts, blocksPerSend *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setBlocksPerSend", blocksPerSend)
}

// SetBlocksPerSend is a paid mutator transaction binding the contract method 0x340e4fc2.
//
// Solidity: function setBlocksPerSend(uint256 blocksPerSend) returns()
func (_ArbOwner *ArbOwnerSession) SetBlocksPerSend(blocksPerSend *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetBlocksPerSend(&_ArbOwner.TransactOpts, blocksPerSend)
}

// SetBlocksPerSend is a paid mutator transaction binding the contract method 0x340e4fc2.
//
// Solidity: function setBlocksPerSend(uint256 blocksPerSend) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetBlocksPerSend(blocksPerSend *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetBlocksPerSend(&_ArbOwner.TransactOpts, blocksPerSend)
}

// SetFeeMaxes is a paid mutator transaction binding the contract method 0x72861aa4.
//
// Solidity: function setFeeMaxes(uint256 num1, uint256 denom1, uint256 num2, uint256 denom2) returns()
func (_ArbOwner *ArbOwnerTransactor) SetFeeMaxes(opts *bind.TransactOpts, num1 *big.Int, denom1 *big.Int, num2 *big.Int, denom2 *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setFeeMaxes", num1, denom1, num2, denom2)
}

// SetFeeMaxes is a paid mutator transaction binding the contract method 0x72861aa4.
//
// Solidity: function setFeeMaxes(uint256 num1, uint256 denom1, uint256 num2, uint256 denom2) returns()
func (_ArbOwner *ArbOwnerSession) SetFeeMaxes(num1 *big.Int, denom1 *big.Int, num2 *big.Int, denom2 *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeeMaxes(&_ArbOwner.TransactOpts, num1, denom1, num2, denom2)
}

// SetFeeMaxes is a paid mutator transaction binding the contract method 0x72861aa4.
//
// Solidity: function setFeeMaxes(uint256 num1, uint256 denom1, uint256 num2, uint256 denom2) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetFeeMaxes(num1 *big.Int, denom1 *big.Int, num2 *big.Int, denom2 *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeeMaxes(&_ArbOwner.TransactOpts, num1, denom1, num2, denom2)
}

// SetFeeRates is a paid mutator transaction binding the contract method 0x30d2361e.
//
// Solidity: function setFeeRates(uint256 num1, uint256 denom1, uint256 num2, uint256 denom2) returns()
func (_ArbOwner *ArbOwnerTransactor) SetFeeRates(opts *bind.TransactOpts, num1 *big.Int, denom1 *big.Int, num2 *big.Int, denom2 *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setFeeRates", num1, denom1, num2, denom2)
}

// SetFeeRates is a paid mutator transaction binding the contract method 0x30d2361e.
//
// Solidity: function setFeeRates(uint256 num1, uint256 denom1, uint256 num2, uint256 denom2) returns()
func (_ArbOwner *ArbOwnerSession) SetFeeRates(num1 *big.Int, denom1 *big.Int, num2 *big.Int, denom2 *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeeRates(&_ArbOwner.TransactOpts, num1, denom1, num2, denom2)
}

// SetFeeRates is a paid mutator transaction binding the contract method 0x30d2361e.
//
// Solidity: function setFeeRates(uint256 num1, uint256 denom1, uint256 num2, uint256 denom2) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetFeeRates(num1 *big.Int, denom1 *big.Int, num2 *big.Int, denom2 *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeeRates(&_ArbOwner.TransactOpts, num1, denom1, num2, denom2)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_ArbOwner *ArbOwnerTransactor) SetFeeRecipient(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setFeeRecipient", recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_ArbOwner *ArbOwnerSession) SetFeeRecipient(recipient common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeeRecipient(&_ArbOwner.TransactOpts, recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetFeeRecipient(recipient common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeeRecipient(&_ArbOwner.TransactOpts, recipient)
}

// StartCodeUpload is a paid mutator transaction binding the contract method 0x31acdf5e.
//
// Solidity: function startCodeUpload() returns()
func (_ArbOwner *ArbOwnerTransactor) StartCodeUpload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "startCodeUpload")
}

// StartCodeUpload is a paid mutator transaction binding the contract method 0x31acdf5e.
//
// Solidity: function startCodeUpload() returns()
func (_ArbOwner *ArbOwnerSession) StartCodeUpload() (*types.Transaction, error) {
	return _ArbOwner.Contract.StartCodeUpload(&_ArbOwner.TransactOpts)
}

// StartCodeUpload is a paid mutator transaction binding the contract method 0x31acdf5e.
//
// Solidity: function startCodeUpload() returns()
func (_ArbOwner *ArbOwnerTransactorSession) StartCodeUpload() (*types.Transaction, error) {
	return _ArbOwner.Contract.StartCodeUpload(&_ArbOwner.TransactOpts)
}
