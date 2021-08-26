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
const ArbOwnerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAllowedSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addToReserveFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowAllSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowOnlyOwnerToSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pluggableId\",\"type\":\"uint256\"}],\"name\":\"bindAddressToPluggable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"marshalledCode\",\"type\":\"bytes\"}],\"name\":\"continueCodeUpload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"createChainParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"constructorData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"deemedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deemedNonce\",\"type\":\"uint256\"}],\"name\":\"deployContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldCodeHash\",\"type\":\"bytes32\"}],\"name\":\"finishCodeUploadAsArbosUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"keepState\",\"type\":\"bool\"}],\"name\":\"finishCodeUploadAsPluggable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAllowedSenders\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllFairGasPriceSenders\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"}],\"name\":\"getChainParameter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalOfEthBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUploadedCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAllowedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isFairGasPriceSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAllowedSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serializeAllParameters\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"which\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setChainParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isFairGasPriceSender\",\"type\":\"bool\"}],\"name\":\"setFairGasPriceSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInGwei\",\"type\":\"uint256\"}],\"name\":\"setL1GasPriceEstimate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startCodeUpload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// GetAllAllowedSenders is a free data retrieval call binding the contract method 0x817ef62e.
//
// Solidity: function getAllAllowedSenders() view returns(bytes)
func (_ArbOwner *ArbOwnerCaller) GetAllAllowedSenders(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getAllAllowedSenders")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAllAllowedSenders is a free data retrieval call binding the contract method 0x817ef62e.
//
// Solidity: function getAllAllowedSenders() view returns(bytes)
func (_ArbOwner *ArbOwnerSession) GetAllAllowedSenders() ([]byte, error) {
	return _ArbOwner.Contract.GetAllAllowedSenders(&_ArbOwner.CallOpts)
}

// GetAllAllowedSenders is a free data retrieval call binding the contract method 0x817ef62e.
//
// Solidity: function getAllAllowedSenders() view returns(bytes)
func (_ArbOwner *ArbOwnerCallerSession) GetAllAllowedSenders() ([]byte, error) {
	return _ArbOwner.Contract.GetAllAllowedSenders(&_ArbOwner.CallOpts)
}

// GetAllFairGasPriceSenders is a free data retrieval call binding the contract method 0xa88bae30.
//
// Solidity: function getAllFairGasPriceSenders() view returns(bytes)
func (_ArbOwner *ArbOwnerCaller) GetAllFairGasPriceSenders(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getAllFairGasPriceSenders")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAllFairGasPriceSenders is a free data retrieval call binding the contract method 0xa88bae30.
//
// Solidity: function getAllFairGasPriceSenders() view returns(bytes)
func (_ArbOwner *ArbOwnerSession) GetAllFairGasPriceSenders() ([]byte, error) {
	return _ArbOwner.Contract.GetAllFairGasPriceSenders(&_ArbOwner.CallOpts)
}

// GetAllFairGasPriceSenders is a free data retrieval call binding the contract method 0xa88bae30.
//
// Solidity: function getAllFairGasPriceSenders() view returns(bytes)
func (_ArbOwner *ArbOwnerCallerSession) GetAllFairGasPriceSenders() ([]byte, error) {
	return _ArbOwner.Contract.GetAllFairGasPriceSenders(&_ArbOwner.CallOpts)
}

// GetChainParameter is a free data retrieval call binding the contract method 0x84352b19.
//
// Solidity: function getChainParameter(bytes32 which) view returns(uint256)
func (_ArbOwner *ArbOwnerCaller) GetChainParameter(opts *bind.CallOpts, which [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getChainParameter", which)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainParameter is a free data retrieval call binding the contract method 0x84352b19.
//
// Solidity: function getChainParameter(bytes32 which) view returns(uint256)
func (_ArbOwner *ArbOwnerSession) GetChainParameter(which [32]byte) (*big.Int, error) {
	return _ArbOwner.Contract.GetChainParameter(&_ArbOwner.CallOpts, which)
}

// GetChainParameter is a free data retrieval call binding the contract method 0x84352b19.
//
// Solidity: function getChainParameter(bytes32 which) view returns(uint256)
func (_ArbOwner *ArbOwnerCallerSession) GetChainParameter(which [32]byte) (*big.Int, error) {
	return _ArbOwner.Contract.GetChainParameter(&_ArbOwner.CallOpts, which)
}

// GetTotalOfEthBalances is a free data retrieval call binding the contract method 0x2816aba5.
//
// Solidity: function getTotalOfEthBalances() view returns(uint256)
func (_ArbOwner *ArbOwnerCaller) GetTotalOfEthBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "getTotalOfEthBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalOfEthBalances is a free data retrieval call binding the contract method 0x2816aba5.
//
// Solidity: function getTotalOfEthBalances() view returns(uint256)
func (_ArbOwner *ArbOwnerSession) GetTotalOfEthBalances() (*big.Int, error) {
	return _ArbOwner.Contract.GetTotalOfEthBalances(&_ArbOwner.CallOpts)
}

// GetTotalOfEthBalances is a free data retrieval call binding the contract method 0x2816aba5.
//
// Solidity: function getTotalOfEthBalances() view returns(uint256)
func (_ArbOwner *ArbOwnerCallerSession) GetTotalOfEthBalances() (*big.Int, error) {
	return _ArbOwner.Contract.GetTotalOfEthBalances(&_ArbOwner.CallOpts)
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

// IsAllowedSender is a free data retrieval call binding the contract method 0xbe8c97b0.
//
// Solidity: function isAllowedSender(address addr) view returns(bool)
func (_ArbOwner *ArbOwnerCaller) IsAllowedSender(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "isAllowedSender", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedSender is a free data retrieval call binding the contract method 0xbe8c97b0.
//
// Solidity: function isAllowedSender(address addr) view returns(bool)
func (_ArbOwner *ArbOwnerSession) IsAllowedSender(addr common.Address) (bool, error) {
	return _ArbOwner.Contract.IsAllowedSender(&_ArbOwner.CallOpts, addr)
}

// IsAllowedSender is a free data retrieval call binding the contract method 0xbe8c97b0.
//
// Solidity: function isAllowedSender(address addr) view returns(bool)
func (_ArbOwner *ArbOwnerCallerSession) IsAllowedSender(addr common.Address) (bool, error) {
	return _ArbOwner.Contract.IsAllowedSender(&_ArbOwner.CallOpts, addr)
}

// IsFairGasPriceSender is a free data retrieval call binding the contract method 0x973f9730.
//
// Solidity: function isFairGasPriceSender(address addr) view returns(bool)
func (_ArbOwner *ArbOwnerCaller) IsFairGasPriceSender(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "isFairGasPriceSender", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFairGasPriceSender is a free data retrieval call binding the contract method 0x973f9730.
//
// Solidity: function isFairGasPriceSender(address addr) view returns(bool)
func (_ArbOwner *ArbOwnerSession) IsFairGasPriceSender(addr common.Address) (bool, error) {
	return _ArbOwner.Contract.IsFairGasPriceSender(&_ArbOwner.CallOpts, addr)
}

// IsFairGasPriceSender is a free data retrieval call binding the contract method 0x973f9730.
//
// Solidity: function isFairGasPriceSender(address addr) view returns(bool)
func (_ArbOwner *ArbOwnerCallerSession) IsFairGasPriceSender(addr common.Address) (bool, error) {
	return _ArbOwner.Contract.IsFairGasPriceSender(&_ArbOwner.CallOpts, addr)
}

// SerializeAllParameters is a free data retrieval call binding the contract method 0xd12ac0e1.
//
// Solidity: function serializeAllParameters() view returns(bytes)
func (_ArbOwner *ArbOwnerCaller) SerializeAllParameters(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ArbOwner.contract.Call(opts, &out, "serializeAllParameters")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SerializeAllParameters is a free data retrieval call binding the contract method 0xd12ac0e1.
//
// Solidity: function serializeAllParameters() view returns(bytes)
func (_ArbOwner *ArbOwnerSession) SerializeAllParameters() ([]byte, error) {
	return _ArbOwner.Contract.SerializeAllParameters(&_ArbOwner.CallOpts)
}

// SerializeAllParameters is a free data retrieval call binding the contract method 0xd12ac0e1.
//
// Solidity: function serializeAllParameters() view returns(bytes)
func (_ArbOwner *ArbOwnerCallerSession) SerializeAllParameters() ([]byte, error) {
	return _ArbOwner.Contract.SerializeAllParameters(&_ArbOwner.CallOpts)
}

// AddAllowedSender is a paid mutator transaction binding the contract method 0xc746c8f4.
//
// Solidity: function addAllowedSender(address addr) returns()
func (_ArbOwner *ArbOwnerTransactor) AddAllowedSender(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "addAllowedSender", addr)
}

// AddAllowedSender is a paid mutator transaction binding the contract method 0xc746c8f4.
//
// Solidity: function addAllowedSender(address addr) returns()
func (_ArbOwner *ArbOwnerSession) AddAllowedSender(addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.AddAllowedSender(&_ArbOwner.TransactOpts, addr)
}

// AddAllowedSender is a paid mutator transaction binding the contract method 0xc746c8f4.
//
// Solidity: function addAllowedSender(address addr) returns()
func (_ArbOwner *ArbOwnerTransactorSession) AddAllowedSender(addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.AddAllowedSender(&_ArbOwner.TransactOpts, addr)
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

// AllowAllSenders is a paid mutator transaction binding the contract method 0xdebb08f5.
//
// Solidity: function allowAllSenders() returns()
func (_ArbOwner *ArbOwnerTransactor) AllowAllSenders(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "allowAllSenders")
}

// AllowAllSenders is a paid mutator transaction binding the contract method 0xdebb08f5.
//
// Solidity: function allowAllSenders() returns()
func (_ArbOwner *ArbOwnerSession) AllowAllSenders() (*types.Transaction, error) {
	return _ArbOwner.Contract.AllowAllSenders(&_ArbOwner.TransactOpts)
}

// AllowAllSenders is a paid mutator transaction binding the contract method 0xdebb08f5.
//
// Solidity: function allowAllSenders() returns()
func (_ArbOwner *ArbOwnerTransactorSession) AllowAllSenders() (*types.Transaction, error) {
	return _ArbOwner.Contract.AllowAllSenders(&_ArbOwner.TransactOpts)
}

// AllowOnlyOwnerToSend is a paid mutator transaction binding the contract method 0xca4ba78c.
//
// Solidity: function allowOnlyOwnerToSend() returns()
func (_ArbOwner *ArbOwnerTransactor) AllowOnlyOwnerToSend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "allowOnlyOwnerToSend")
}

// AllowOnlyOwnerToSend is a paid mutator transaction binding the contract method 0xca4ba78c.
//
// Solidity: function allowOnlyOwnerToSend() returns()
func (_ArbOwner *ArbOwnerSession) AllowOnlyOwnerToSend() (*types.Transaction, error) {
	return _ArbOwner.Contract.AllowOnlyOwnerToSend(&_ArbOwner.TransactOpts)
}

// AllowOnlyOwnerToSend is a paid mutator transaction binding the contract method 0xca4ba78c.
//
// Solidity: function allowOnlyOwnerToSend() returns()
func (_ArbOwner *ArbOwnerTransactorSession) AllowOnlyOwnerToSend() (*types.Transaction, error) {
	return _ArbOwner.Contract.AllowOnlyOwnerToSend(&_ArbOwner.TransactOpts)
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

// CreateChainParameter is a paid mutator transaction binding the contract method 0xe1e639cf.
//
// Solidity: function createChainParameter(bytes32 which, uint256 value) returns()
func (_ArbOwner *ArbOwnerTransactor) CreateChainParameter(opts *bind.TransactOpts, which [32]byte, value *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "createChainParameter", which, value)
}

// CreateChainParameter is a paid mutator transaction binding the contract method 0xe1e639cf.
//
// Solidity: function createChainParameter(bytes32 which, uint256 value) returns()
func (_ArbOwner *ArbOwnerSession) CreateChainParameter(which [32]byte, value *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.CreateChainParameter(&_ArbOwner.TransactOpts, which, value)
}

// CreateChainParameter is a paid mutator transaction binding the contract method 0xe1e639cf.
//
// Solidity: function createChainParameter(bytes32 which, uint256 value) returns()
func (_ArbOwner *ArbOwnerTransactorSession) CreateChainParameter(which [32]byte, value *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.CreateChainParameter(&_ArbOwner.TransactOpts, which, value)
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

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0xb3def0d7.
//
// Solidity: function finishCodeUploadAsArbosUpgrade(bytes32 newCodeHash, bytes32 oldCodeHash) returns()
func (_ArbOwner *ArbOwnerTransactor) FinishCodeUploadAsArbosUpgrade(opts *bind.TransactOpts, newCodeHash [32]byte, oldCodeHash [32]byte) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "finishCodeUploadAsArbosUpgrade", newCodeHash, oldCodeHash)
}

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0xb3def0d7.
//
// Solidity: function finishCodeUploadAsArbosUpgrade(bytes32 newCodeHash, bytes32 oldCodeHash) returns()
func (_ArbOwner *ArbOwnerSession) FinishCodeUploadAsArbosUpgrade(newCodeHash [32]byte, oldCodeHash [32]byte) (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishCodeUploadAsArbosUpgrade(&_ArbOwner.TransactOpts, newCodeHash, oldCodeHash)
}

// FinishCodeUploadAsArbosUpgrade is a paid mutator transaction binding the contract method 0xb3def0d7.
//
// Solidity: function finishCodeUploadAsArbosUpgrade(bytes32 newCodeHash, bytes32 oldCodeHash) returns()
func (_ArbOwner *ArbOwnerTransactorSession) FinishCodeUploadAsArbosUpgrade(newCodeHash [32]byte, oldCodeHash [32]byte) (*types.Transaction, error) {
	return _ArbOwner.Contract.FinishCodeUploadAsArbosUpgrade(&_ArbOwner.TransactOpts, newCodeHash, oldCodeHash)
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

// RemoveAllowedSender is a paid mutator transaction binding the contract method 0x471eab5c.
//
// Solidity: function removeAllowedSender(address addr) returns()
func (_ArbOwner *ArbOwnerTransactor) RemoveAllowedSender(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "removeAllowedSender", addr)
}

// RemoveAllowedSender is a paid mutator transaction binding the contract method 0x471eab5c.
//
// Solidity: function removeAllowedSender(address addr) returns()
func (_ArbOwner *ArbOwnerSession) RemoveAllowedSender(addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.RemoveAllowedSender(&_ArbOwner.TransactOpts, addr)
}

// RemoveAllowedSender is a paid mutator transaction binding the contract method 0x471eab5c.
//
// Solidity: function removeAllowedSender(address addr) returns()
func (_ArbOwner *ArbOwnerTransactorSession) RemoveAllowedSender(addr common.Address) (*types.Transaction, error) {
	return _ArbOwner.Contract.RemoveAllowedSender(&_ArbOwner.TransactOpts, addr)
}

// SetChainParameter is a paid mutator transaction binding the contract method 0x966e2505.
//
// Solidity: function setChainParameter(bytes32 which, uint256 value) returns()
func (_ArbOwner *ArbOwnerTransactor) SetChainParameter(opts *bind.TransactOpts, which [32]byte, value *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setChainParameter", which, value)
}

// SetChainParameter is a paid mutator transaction binding the contract method 0x966e2505.
//
// Solidity: function setChainParameter(bytes32 which, uint256 value) returns()
func (_ArbOwner *ArbOwnerSession) SetChainParameter(which [32]byte, value *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetChainParameter(&_ArbOwner.TransactOpts, which, value)
}

// SetChainParameter is a paid mutator transaction binding the contract method 0x966e2505.
//
// Solidity: function setChainParameter(bytes32 which, uint256 value) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetChainParameter(which [32]byte, value *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetChainParameter(&_ArbOwner.TransactOpts, which, value)
}

// SetFairGasPriceSender is a paid mutator transaction binding the contract method 0xcad462d4.
//
// Solidity: function setFairGasPriceSender(address addr, bool isFairGasPriceSender) returns()
func (_ArbOwner *ArbOwnerTransactor) SetFairGasPriceSender(opts *bind.TransactOpts, addr common.Address, isFairGasPriceSender bool) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setFairGasPriceSender", addr, isFairGasPriceSender)
}

// SetFairGasPriceSender is a paid mutator transaction binding the contract method 0xcad462d4.
//
// Solidity: function setFairGasPriceSender(address addr, bool isFairGasPriceSender) returns()
func (_ArbOwner *ArbOwnerSession) SetFairGasPriceSender(addr common.Address, isFairGasPriceSender bool) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFairGasPriceSender(&_ArbOwner.TransactOpts, addr, isFairGasPriceSender)
}

// SetFairGasPriceSender is a paid mutator transaction binding the contract method 0xcad462d4.
//
// Solidity: function setFairGasPriceSender(address addr, bool isFairGasPriceSender) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetFairGasPriceSender(addr common.Address, isFairGasPriceSender bool) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetFairGasPriceSender(&_ArbOwner.TransactOpts, addr, isFairGasPriceSender)
}

// SetL1GasPriceEstimate is a paid mutator transaction binding the contract method 0x4290549e.
//
// Solidity: function setL1GasPriceEstimate(uint256 priceInGwei) returns()
func (_ArbOwner *ArbOwnerTransactor) SetL1GasPriceEstimate(opts *bind.TransactOpts, priceInGwei *big.Int) (*types.Transaction, error) {
	return _ArbOwner.contract.Transact(opts, "setL1GasPriceEstimate", priceInGwei)
}

// SetL1GasPriceEstimate is a paid mutator transaction binding the contract method 0x4290549e.
//
// Solidity: function setL1GasPriceEstimate(uint256 priceInGwei) returns()
func (_ArbOwner *ArbOwnerSession) SetL1GasPriceEstimate(priceInGwei *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetL1GasPriceEstimate(&_ArbOwner.TransactOpts, priceInGwei)
}

// SetL1GasPriceEstimate is a paid mutator transaction binding the contract method 0x4290549e.
//
// Solidity: function setL1GasPriceEstimate(uint256 priceInGwei) returns()
func (_ArbOwner *ArbOwnerTransactorSession) SetL1GasPriceEstimate(priceInGwei *big.Int) (*types.Transaction, error) {
	return _ArbOwner.Contract.SetL1GasPriceEstimate(&_ArbOwner.TransactOpts, priceInGwei)
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
