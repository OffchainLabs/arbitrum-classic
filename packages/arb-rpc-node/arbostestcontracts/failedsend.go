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

// FailedSendABI is the input ABI used to generate the binding from.
const FailedSendABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FailedSendBin is the compiled bytecode used for deploying new contracts.
var FailedSendBin = "0x608060405234801561001057600080fd5b506101c7806100206000396000f3fe6080604052600436106100295760003560e01c80633e58c58c1461002e57806368742da614610056575b600080fd5b6100546004803603602081101561004457600080fd5b50356001600160a01b0316610089565b005b34801561006257600080fd5b506100546004803603602081101561007957600080fd5b50356001600160a01b0316610125565b6040805163343a16d360e11b81526001600160a01b0383166004820152905130916368742da691602480830192600092919082900301818387803b1580156100d057600080fd5b505af11580156100e4573d6000803e3d6000fd5b505050506040805162461bcd60e51b815260206004820152600d60248201526c666f726365206661696c75726560981b604482015290519081900360640190fd5b60646001600160a01b03166325e1606347836040518363ffffffff1660e01b815260040180826001600160a01b031681526020019150506000604051808303818588803b15801561017557600080fd5b505af1158015610189573d6000803e3d6000fd5b50505050505056fea264697066735822122044178fdeee557751148d0f3a7d757bd6219ecb5e0043bcee1ef78336d14b2e6d64736f6c634300060c0033"

// DeployFailedSend deploys a new Ethereum contract, binding an instance of FailedSend to it.
func DeployFailedSend(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FailedSend, error) {
	parsed, err := abi.JSON(strings.NewReader(FailedSendABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FailedSendBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FailedSend{FailedSendCaller: FailedSendCaller{contract: contract}, FailedSendTransactor: FailedSendTransactor{contract: contract}, FailedSendFilterer: FailedSendFilterer{contract: contract}}, nil
}

// FailedSend is an auto generated Go binding around an Ethereum contract.
type FailedSend struct {
	FailedSendCaller     // Read-only binding to the contract
	FailedSendTransactor // Write-only binding to the contract
	FailedSendFilterer   // Log filterer for contract events
}

// FailedSendCaller is an auto generated read-only Go binding around an Ethereum contract.
type FailedSendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedSendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FailedSendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedSendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FailedSendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedSendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FailedSendSession struct {
	Contract     *FailedSend       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FailedSendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FailedSendCallerSession struct {
	Contract *FailedSendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FailedSendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FailedSendTransactorSession struct {
	Contract     *FailedSendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FailedSendRaw is an auto generated low-level Go binding around an Ethereum contract.
type FailedSendRaw struct {
	Contract *FailedSend // Generic contract binding to access the raw methods on
}

// FailedSendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FailedSendCallerRaw struct {
	Contract *FailedSendCaller // Generic read-only contract binding to access the raw methods on
}

// FailedSendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FailedSendTransactorRaw struct {
	Contract *FailedSendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFailedSend creates a new instance of FailedSend, bound to a specific deployed contract.
func NewFailedSend(address common.Address, backend bind.ContractBackend) (*FailedSend, error) {
	contract, err := bindFailedSend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FailedSend{FailedSendCaller: FailedSendCaller{contract: contract}, FailedSendTransactor: FailedSendTransactor{contract: contract}, FailedSendFilterer: FailedSendFilterer{contract: contract}}, nil
}

// NewFailedSendCaller creates a new read-only instance of FailedSend, bound to a specific deployed contract.
func NewFailedSendCaller(address common.Address, caller bind.ContractCaller) (*FailedSendCaller, error) {
	contract, err := bindFailedSend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FailedSendCaller{contract: contract}, nil
}

// NewFailedSendTransactor creates a new write-only instance of FailedSend, bound to a specific deployed contract.
func NewFailedSendTransactor(address common.Address, transactor bind.ContractTransactor) (*FailedSendTransactor, error) {
	contract, err := bindFailedSend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FailedSendTransactor{contract: contract}, nil
}

// NewFailedSendFilterer creates a new log filterer instance of FailedSend, bound to a specific deployed contract.
func NewFailedSendFilterer(address common.Address, filterer bind.ContractFilterer) (*FailedSendFilterer, error) {
	contract, err := bindFailedSend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FailedSendFilterer{contract: contract}, nil
}

// bindFailedSend binds a generic wrapper to an already deployed contract.
func bindFailedSend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FailedSendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FailedSend *FailedSendRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FailedSend.Contract.FailedSendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FailedSend *FailedSendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FailedSend.Contract.FailedSendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FailedSend *FailedSendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FailedSend.Contract.FailedSendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FailedSend *FailedSendCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FailedSend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FailedSend *FailedSendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FailedSend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FailedSend *FailedSendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FailedSend.Contract.contract.Transact(opts, method, params...)
}

// Send is a paid mutator transaction binding the contract method 0x3e58c58c.
//
// Solidity: function send(address dest) payable returns()
func (_FailedSend *FailedSendTransactor) Send(opts *bind.TransactOpts, dest common.Address) (*types.Transaction, error) {
	return _FailedSend.contract.Transact(opts, "send", dest)
}

// Send is a paid mutator transaction binding the contract method 0x3e58c58c.
//
// Solidity: function send(address dest) payable returns()
func (_FailedSend *FailedSendSession) Send(dest common.Address) (*types.Transaction, error) {
	return _FailedSend.Contract.Send(&_FailedSend.TransactOpts, dest)
}

// Send is a paid mutator transaction binding the contract method 0x3e58c58c.
//
// Solidity: function send(address dest) payable returns()
func (_FailedSend *FailedSendTransactorSession) Send(dest common.Address) (*types.Transaction, error) {
	return _FailedSend.Contract.Send(&_FailedSend.TransactOpts, dest)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address dest) returns()
func (_FailedSend *FailedSendTransactor) WithdrawFunds(opts *bind.TransactOpts, dest common.Address) (*types.Transaction, error) {
	return _FailedSend.contract.Transact(opts, "withdrawFunds", dest)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address dest) returns()
func (_FailedSend *FailedSendSession) WithdrawFunds(dest common.Address) (*types.Transaction, error) {
	return _FailedSend.Contract.WithdrawFunds(&_FailedSend.TransactOpts, dest)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address dest) returns()
func (_FailedSend *FailedSendTransactorSession) WithdrawFunds(dest common.Address) (*types.Transaction, error) {
	return _FailedSend.Contract.WithdrawFunds(&_FailedSend.TransactOpts, dest)
}

// IFailedSendABI is the input ABI used to generate the binding from.
const IFailedSendABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IFailedSend is an auto generated Go binding around an Ethereum contract.
type IFailedSend struct {
	IFailedSendCaller     // Read-only binding to the contract
	IFailedSendTransactor // Write-only binding to the contract
	IFailedSendFilterer   // Log filterer for contract events
}

// IFailedSendCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFailedSendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFailedSendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFailedSendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFailedSendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFailedSendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFailedSendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFailedSendSession struct {
	Contract     *IFailedSend      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFailedSendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFailedSendCallerSession struct {
	Contract *IFailedSendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IFailedSendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFailedSendTransactorSession struct {
	Contract     *IFailedSendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IFailedSendRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFailedSendRaw struct {
	Contract *IFailedSend // Generic contract binding to access the raw methods on
}

// IFailedSendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFailedSendCallerRaw struct {
	Contract *IFailedSendCaller // Generic read-only contract binding to access the raw methods on
}

// IFailedSendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFailedSendTransactorRaw struct {
	Contract *IFailedSendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFailedSend creates a new instance of IFailedSend, bound to a specific deployed contract.
func NewIFailedSend(address common.Address, backend bind.ContractBackend) (*IFailedSend, error) {
	contract, err := bindIFailedSend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFailedSend{IFailedSendCaller: IFailedSendCaller{contract: contract}, IFailedSendTransactor: IFailedSendTransactor{contract: contract}, IFailedSendFilterer: IFailedSendFilterer{contract: contract}}, nil
}

// NewIFailedSendCaller creates a new read-only instance of IFailedSend, bound to a specific deployed contract.
func NewIFailedSendCaller(address common.Address, caller bind.ContractCaller) (*IFailedSendCaller, error) {
	contract, err := bindIFailedSend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFailedSendCaller{contract: contract}, nil
}

// NewIFailedSendTransactor creates a new write-only instance of IFailedSend, bound to a specific deployed contract.
func NewIFailedSendTransactor(address common.Address, transactor bind.ContractTransactor) (*IFailedSendTransactor, error) {
	contract, err := bindIFailedSend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFailedSendTransactor{contract: contract}, nil
}

// NewIFailedSendFilterer creates a new log filterer instance of IFailedSend, bound to a specific deployed contract.
func NewIFailedSendFilterer(address common.Address, filterer bind.ContractFilterer) (*IFailedSendFilterer, error) {
	contract, err := bindIFailedSend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFailedSendFilterer{contract: contract}, nil
}

// bindIFailedSend binds a generic wrapper to an already deployed contract.
func bindIFailedSend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFailedSendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFailedSend *IFailedSendRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFailedSend.Contract.IFailedSendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFailedSend *IFailedSendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFailedSend.Contract.IFailedSendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFailedSend *IFailedSendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFailedSend.Contract.IFailedSendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFailedSend *IFailedSendCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFailedSend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFailedSend *IFailedSendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFailedSend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFailedSend *IFailedSendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFailedSend.Contract.contract.Transact(opts, method, params...)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address dest) returns()
func (_IFailedSend *IFailedSendTransactor) WithdrawFunds(opts *bind.TransactOpts, dest common.Address) (*types.Transaction, error) {
	return _IFailedSend.contract.Transact(opts, "withdrawFunds", dest)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address dest) returns()
func (_IFailedSend *IFailedSendSession) WithdrawFunds(dest common.Address) (*types.Transaction, error) {
	return _IFailedSend.Contract.WithdrawFunds(&_IFailedSend.TransactOpts, dest)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address dest) returns()
func (_IFailedSend *IFailedSendTransactorSession) WithdrawFunds(dest common.Address) (*types.Transaction, error) {
	return _IFailedSend.Contract.WithdrawFunds(&_IFailedSend.TransactOpts, dest)
}

// SysABI is the input ABI used to generate the binding from.
const SysABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Sys is an auto generated Go binding around an Ethereum contract.
type Sys struct {
	SysCaller     // Read-only binding to the contract
	SysTransactor // Write-only binding to the contract
	SysFilterer   // Log filterer for contract events
}

// SysCaller is an auto generated read-only Go binding around an Ethereum contract.
type SysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SysSession struct {
	Contract     *Sys              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SysCallerSession struct {
	Contract *SysCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SysTransactorSession struct {
	Contract     *SysTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SysRaw is an auto generated low-level Go binding around an Ethereum contract.
type SysRaw struct {
	Contract *Sys // Generic contract binding to access the raw methods on
}

// SysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SysCallerRaw struct {
	Contract *SysCaller // Generic read-only contract binding to access the raw methods on
}

// SysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SysTransactorRaw struct {
	Contract *SysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSys creates a new instance of Sys, bound to a specific deployed contract.
func NewSys(address common.Address, backend bind.ContractBackend) (*Sys, error) {
	contract, err := bindSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sys{SysCaller: SysCaller{contract: contract}, SysTransactor: SysTransactor{contract: contract}, SysFilterer: SysFilterer{contract: contract}}, nil
}

// NewSysCaller creates a new read-only instance of Sys, bound to a specific deployed contract.
func NewSysCaller(address common.Address, caller bind.ContractCaller) (*SysCaller, error) {
	contract, err := bindSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SysCaller{contract: contract}, nil
}

// NewSysTransactor creates a new write-only instance of Sys, bound to a specific deployed contract.
func NewSysTransactor(address common.Address, transactor bind.ContractTransactor) (*SysTransactor, error) {
	contract, err := bindSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SysTransactor{contract: contract}, nil
}

// NewSysFilterer creates a new log filterer instance of Sys, bound to a specific deployed contract.
func NewSysFilterer(address common.Address, filterer bind.ContractFilterer) (*SysFilterer, error) {
	contract, err := bindSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SysFilterer{contract: contract}, nil
}

// bindSys binds a generic wrapper to an already deployed contract.
func bindSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sys *SysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sys.Contract.SysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sys *SysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sys.Contract.SysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sys *SysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sys.Contract.SysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sys *SysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sys *SysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sys *SysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sys.Contract.contract.Transact(opts, method, params...)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address dest) payable returns()
func (_Sys *SysTransactor) WithdrawEth(opts *bind.TransactOpts, dest common.Address) (*types.Transaction, error) {
	return _Sys.contract.Transact(opts, "withdrawEth", dest)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address dest) payable returns()
func (_Sys *SysSession) WithdrawEth(dest common.Address) (*types.Transaction, error) {
	return _Sys.Contract.WithdrawEth(&_Sys.TransactOpts, dest)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address dest) payable returns()
func (_Sys *SysTransactorSession) WithdrawEth(dest common.Address) (*types.Transaction, error) {
	return _Sys.Contract.WithdrawEth(&_Sys.TransactOpts, dest)
}
