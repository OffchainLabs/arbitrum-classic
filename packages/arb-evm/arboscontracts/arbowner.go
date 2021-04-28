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
const ArbOwnerABI = "[{\"inputs\":[],\"name\":\"addToReserveFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pluggableId\",\"type\":\"uint256\"}],\"name\":\"bindAddressToPluggable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"marshalledCode\",\"type\":\"bytes\"}],\"name\":\"continueCodeUpload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"constructorData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"deemedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deemedNonce\",\"type\":\"uint256\"}],\"name\":\"deployContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requiredCodeHash\",\"type\":\"bytes32\"}],\"name\":\"finishCodeUploadAsArbosUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"keepState\",\"type\":\"bool\"}],\"name\":\"finishCodeUploadAsPluggable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipients\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUploadedCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwnerAddr\",\"type\":\"address\"}],\"name\":\"giveOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setFairGasPriceSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"netFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"congestionFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipients\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setFeesEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"speedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPoolMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTxGasLimit\",\"type\":\"uint256\"}],\"name\":\"setGasAccountingParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksPerSend\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startCodeUpload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbOwnerFuncSigs maps the 4-byte function signature to its string representation.
var ArbOwnerFuncSigs = map[string]string{
	"c3bf429d": "addToReserveFunds()",
	"f589445c": "bindAddressToPluggable(address,uint256)",
	"56331f75": "continueCodeUpload(bytes)",
	"5b0b7f1e": "deployContract(bytes,address,uint256)",
	"10e7af17": "finishCodeUploadAsArbosUpgrade(bytes32)",
	"f4f4e136": "finishCodeUploadAsPluggable(uint256,bool)",
	"08df6923": "getFeeRecipients()",
	"c060180d": "getUploadedCodeHash()",
	"e3a0a148": "giveOwnership(address)",
	"ba7f4cc6": "setFairGasPriceSender(address)",
	"c6cabb40": "setFeeRecipients(address,address)",
	"a901dd92": "setFeesEnabled(bool)",
	"8ea93643": "setGasAccountingParams(uint256,uint256,uint256)",
	"29854f47": "setSecondsPerSend(uint256)",
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

// GetFeeRecipients is a free data retrieval call binding the contract method 0x08df6923.
//
// Solidity: function getFeeRecipients() view returns(address, address)
func (_ArbOwner *ArbOwnerCaller) GetFeeRecipients(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getFeeRecipients")

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetFeeRecipients is a free data retrieval call binding the contract method 0x08df6923.
//
// Solidity: function getFeeRecipients() view returns(address, address)
func (_ArbOwner *ArbOwnerSession) GetFeeRecipients() (common.Address, common.Address, error) {
	return _ArbOwner.Contract.GetFeeRecipients(&_ArbOwner.CallOpts)
}

// GetFeeRecipients is a free data retrieval call binding the contract method 0x08df6923.
//
// Solidity: function getFeeRecipients() view returns(address, address)
func (_ArbOwner *ArbOwnerCallerSession) GetFeeRecipients() (common.Address, common.Address, error) {
	return _ArbOwner.Contract.GetFeeRecipients(&_ArbOwner.CallOpts)
}

// GetUploadedCodeHash is a free data retrieval call binding the contract method 0xc060180d.
//
// Solidity: function getUploadedCodeHash() view returns(bytes32)
func (_ArbOwner *ArbOwnerCaller) GetUploadedCodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getUploadedCodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUploadedCodeHash is a free data retrieval call binding the contract method 0xc060180d.
//
// Solidity: function getUploadedCodeHash() view returns(bytes32)
func (_ArbOwner *ArbOwnerSession) GetUploadedCodeHash() ([32]byte, error) {
	return _ArbOwner.Contract.GetUploadedCodeHash(&_ArbOwner.CallOpts)
}

// GetUploadedCodeHash is a free data retrieval call binding the contract method 0xc060180d.
//
// Solidity: function getUploadedCodeHash() view returns(bytes32)
func (_ArbOwner *ArbOwnerCallerSession) GetUploadedCodeHash() ([32]byte, error) {
	return _ArbOwner.Contract.GetUploadedCodeHash(&_ArbOwner.CallOpts)
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

// DeployContract is a paid mutator transaction binding the contract method 0x5b0b7f1e.
//
// Solidity: function deployContract(bytes constructorData, address deemedSender, uint256 deemedNonce) payable returns(address)
func (_ArbOwner *ArbOwnerTransactor) DeployContract(opts *bind.TransactOpts, constructorData []byte, deemedSender common.Address, deemedNonce *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "deployContract", constructorData, deemedSender, deemedNonce)
}

// DeployContract is a paid mutator transaction binding the contract method 0x5b0b7f1e.
//
// Solidity: function deployContract(bytes constructorData, address deemedSender, uint256 deemedNonce) payable returns(address)
func (_ArbOwner *ArbOwnerSession) DeployContract(constructorData []byte, deemedSender common.Address, deemedNonce *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.DeployContract(&_ArbOwner.TransactOpts, constructorData, deemedSender, deemedNonce)
}

// DeployContract is a paid mutator transaction binding the contract method 0x5b0b7f1e.
//
// Solidity: function deployContract(bytes constructorData, address deemedSender, uint256 deemedNonce) payable returns(address)
func (_ArbOwner *ArbOwnerTransactorSession) DeployContract(constructorData []byte, deemedSender common.Address, deemedNonce *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.DeployContract(&_ArbOwner.TransactOpts, constructorData, deemedSender, deemedNonce)
}

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0x10e7af17.
//
// Solidity: function finishCodeUploadAsArbosUpgrade(bytes32 requiredCodeHash) returns()
func (_ArbOwner *ArbOwnerTransactor) FinishCodeUploadAsArbosUpgrade(opts *bind.TransactOpts, requiredCodeHash [32]byte) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "finishCodeUploadAsArbosUpgrade", requiredCodeHash)
}

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0x10e7af17.
//
// Solidity: function finishCodeUploadAsArbosUpgrade(bytes32 requiredCodeHash) returns()
func (_ArbOwner *ArbOwnerSession) FinishCodeUploadAsArbosUpgrade(requiredCodeHash [32]byte) (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishCodeUploadAsArbosUpgrade(&_ArbOwner.TransactOpts, requiredCodeHash)
}

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0x10e7af17.
//
// Solidity: function finishCodeUploadAsArbosUpgrade(bytes32 requiredCodeHash) returns()
func (_ArbOwner *ArbOwnerTransactorSession) FinishCodeUploadAsArbosUpgrade(requiredCodeHash [32]byte) (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishCodeUploadAsArbosUpgrade(&_ArbOwner.TransactOpts, requiredCodeHash)
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

// SetFairGasPriceSender is a paid mutator transaction binding the contract method 0xba7f4cc6.
//
// Solidity: function setFairGasPriceSender(address addr) returns()
func (_ArbOwner *ArbOwnerTransactor) SetFairGasPriceSender(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setFairGasPriceSender", addr)
}

// SetFairGasPriceSender is a paid mutator transaction binding the contract method 0xba7f4cc6.
//
// Solidity: function setFairGasPriceSender(address addr) returns()
func (_ArbOwner *ArbOwnerSession) SetFairGasPriceSender(addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFairGasPriceSender(&_ArbOwner.TransactOpts, addr)
}

// SetFairGasPriceSender is a paid mutator transaction binding the contract method 0xba7f4cc6.
//
// Solidity: function setFairGasPriceSender(address addr) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetFairGasPriceSender(addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFairGasPriceSender(&_ArbOwner.TransactOpts, addr)
}

// SetFeeRecipients is a paid mutator transaction binding the contract method 0xc6cabb40.
//
// Solidity: function setFeeRecipients(address netFeeRecipient, address congestionFeeRecipient) returns()
func (_ArbOwner *ArbOwnerTransactor) SetFeeRecipients(opts *bind.TransactOpts, netFeeRecipient common.Address, congestionFeeRecipient common.Address) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setFeeRecipients", netFeeRecipient, congestionFeeRecipient)
}

// SetFeeRecipients is a paid mutator transaction binding the contract method 0xc6cabb40.
//
// Solidity: function setFeeRecipients(address netFeeRecipient, address congestionFeeRecipient) returns()
func (_ArbOwner *ArbOwnerSession) SetFeeRecipients(netFeeRecipient common.Address, congestionFeeRecipient common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeeRecipients(&_ArbOwner.TransactOpts, netFeeRecipient, congestionFeeRecipient)
}

// SetFeeRecipients is a paid mutator transaction binding the contract method 0xc6cabb40.
//
// Solidity: function setFeeRecipients(address netFeeRecipient, address congestionFeeRecipient) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetFeeRecipients(netFeeRecipient common.Address, congestionFeeRecipient common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeeRecipients(&_ArbOwner.TransactOpts, netFeeRecipient, congestionFeeRecipient)
}

// SetFeesEnabled is a paid mutator transaction binding the contract method 0xa901dd92.
//
// Solidity: function setFeesEnabled(bool enabled) returns()
func (_ArbOwner *ArbOwnerTransactor) SetFeesEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setFeesEnabled", enabled)
}

// SetFeesEnabled is a paid mutator transaction binding the contract method 0xa901dd92.
//
// Solidity: function setFeesEnabled(bool enabled) returns()
func (_ArbOwner *ArbOwnerSession) SetFeesEnabled(enabled bool) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeesEnabled(&_ArbOwner.TransactOpts, enabled)
}

// SetFeesEnabled is a paid mutator transaction binding the contract method 0xa901dd92.
//
// Solidity: function setFeesEnabled(bool enabled) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetFeesEnabled(enabled bool) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFeesEnabled(&_ArbOwner.TransactOpts, enabled)
}

// SetGasAccountingParams is a paid mutator transaction binding the contract method 0x8ea93643.
//
// Solidity: function setGasAccountingParams(uint256 speedLimitPerBlock, uint256 gasPoolMax, uint256 maxTxGasLimit) returns()
func (_ArbOwner *ArbOwnerTransactor) SetGasAccountingParams(opts *bind.TransactOpts, speedLimitPerBlock *big.Int, gasPoolMax *big.Int, maxTxGasLimit *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setGasAccountingParams", speedLimitPerBlock, gasPoolMax, maxTxGasLimit)
}

// SetGasAccountingParams is a paid mutator transaction binding the contract method 0x8ea93643.
//
// Solidity: function setGasAccountingParams(uint256 speedLimitPerBlock, uint256 gasPoolMax, uint256 maxTxGasLimit) returns()
func (_ArbOwner *ArbOwnerSession) SetGasAccountingParams(speedLimitPerBlock *big.Int, gasPoolMax *big.Int, maxTxGasLimit *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetGasAccountingParams(&_ArbOwner.TransactOpts, speedLimitPerBlock, gasPoolMax, maxTxGasLimit)
}

// SetGasAccountingParams is a paid mutator transaction binding the contract method 0x8ea93643.
//
// Solidity: function setGasAccountingParams(uint256 speedLimitPerBlock, uint256 gasPoolMax, uint256 maxTxGasLimit) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetGasAccountingParams(speedLimitPerBlock *big.Int, gasPoolMax *big.Int, maxTxGasLimit *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetGasAccountingParams(&_ArbOwner.TransactOpts, speedLimitPerBlock, gasPoolMax, maxTxGasLimit)
}

// SetSecondsPerSend is a paid mutator transaction binding the contract method 0x29854f47.
//
// Solidity: function setSecondsPerSend(uint256 blocksPerSend) returns()
func (_ArbOwner *ArbOwnerTransactor) SetSecondsPerSend(opts *bind.TransactOpts, blocksPerSend *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setSecondsPerSend", blocksPerSend)
}

// SetSecondsPerSend is a paid mutator transaction binding the contract method 0x29854f47.
//
// Solidity: function setSecondsPerSend(uint256 blocksPerSend) returns()
func (_ArbOwner *ArbOwnerSession) SetSecondsPerSend(blocksPerSend *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetSecondsPerSend(&_ArbOwner.TransactOpts, blocksPerSend)
}

// SetSecondsPerSend is a paid mutator transaction binding the contract method 0x29854f47.
//
// Solidity: function setSecondsPerSend(uint256 blocksPerSend) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetSecondsPerSend(blocksPerSend *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetSecondsPerSend(&_ArbOwner.TransactOpts, blocksPerSend)
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
