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

// StorageABI is the input ABI used to generate the binding from.
const StorageABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"failGetStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StorageFuncSigs maps the 4-byte function signature to its string representation.
var StorageFuncSigs = map[string]string{
	"188f9139": "failGetStorage()",
}

// StorageBin is the compiled bytecode used for deploying new contracts.
var StorageBin = "0x608060405234801561001057600080fd5b5061303960015560f2806100256000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063188f913914602d575b600080fd5b60336045565b60408051918252519081900360200190f35b6040805163a169625f60e01b815230600482015260016024820152905160009160649163a169625f91604480820192602092909190829003018186803b158015608d57600080fd5b505afa15801560a0573d6000803e3d6000fd5b505050506040513d602081101560b557600080fd5b505190509056fea2646970667358221220a67a8209534ea581ed7713550e703740681f0a77f4d4d3c6592bb7376261492064736f6c634300060c0033"

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// FailGetStorage is a paid mutator transaction binding the contract method 0x188f9139.
//
// Solidity: function failGetStorage() returns(uint256)
func (_Storage *StorageTransactor) FailGetStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "failGetStorage")
}

// FailGetStorage is a paid mutator transaction binding the contract method 0x188f9139.
//
// Solidity: function failGetStorage() returns(uint256)
func (_Storage *StorageSession) FailGetStorage() (*types.Transaction, error) {
	return _Storage.Contract.FailGetStorage(&_Storage.TransactOpts)
}

// FailGetStorage is a paid mutator transaction binding the contract method 0x188f9139.
//
// Solidity: function failGetStorage() returns(uint256)
func (_Storage *StorageTransactorSession) FailGetStorage() (*types.Transaction, error) {
	return _Storage.Contract.FailGetStorage(&_Storage.TransactOpts)
}

// Sys2ABI is the input ABI used to generate the binding from.
const Sys2ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Sys2FuncSigs maps the 4-byte function signature to its string representation.
var Sys2FuncSigs = map[string]string{
	"a169625f": "getStorageAt(address,uint256)",
}

// Sys2 is an auto generated Go binding around an Ethereum contract.
type Sys2 struct {
	Sys2Caller     // Read-only binding to the contract
	Sys2Transactor // Write-only binding to the contract
	Sys2Filterer   // Log filterer for contract events
}

// Sys2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Sys2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sys2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Sys2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sys2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Sys2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sys2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Sys2Session struct {
	Contract     *Sys2             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Sys2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Sys2CallerSession struct {
	Contract *Sys2Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Sys2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Sys2TransactorSession struct {
	Contract     *Sys2Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Sys2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Sys2Raw struct {
	Contract *Sys2 // Generic contract binding to access the raw methods on
}

// Sys2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Sys2CallerRaw struct {
	Contract *Sys2Caller // Generic read-only contract binding to access the raw methods on
}

// Sys2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Sys2TransactorRaw struct {
	Contract *Sys2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSys2 creates a new instance of Sys2, bound to a specific deployed contract.
func NewSys2(address common.Address, backend bind.ContractBackend) (*Sys2, error) {
	contract, err := bindSys2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sys2{Sys2Caller: Sys2Caller{contract: contract}, Sys2Transactor: Sys2Transactor{contract: contract}, Sys2Filterer: Sys2Filterer{contract: contract}}, nil
}

// NewSys2Caller creates a new read-only instance of Sys2, bound to a specific deployed contract.
func NewSys2Caller(address common.Address, caller bind.ContractCaller) (*Sys2Caller, error) {
	contract, err := bindSys2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Sys2Caller{contract: contract}, nil
}

// NewSys2Transactor creates a new write-only instance of Sys2, bound to a specific deployed contract.
func NewSys2Transactor(address common.Address, transactor bind.ContractTransactor) (*Sys2Transactor, error) {
	contract, err := bindSys2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Sys2Transactor{contract: contract}, nil
}

// NewSys2Filterer creates a new log filterer instance of Sys2, bound to a specific deployed contract.
func NewSys2Filterer(address common.Address, filterer bind.ContractFilterer) (*Sys2Filterer, error) {
	contract, err := bindSys2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Sys2Filterer{contract: contract}, nil
}

// bindSys2 binds a generic wrapper to an already deployed contract.
func bindSys2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Sys2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sys2 *Sys2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sys2.Contract.Sys2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sys2 *Sys2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sys2.Contract.Sys2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sys2 *Sys2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sys2.Contract.Sys2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sys2 *Sys2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sys2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sys2 *Sys2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sys2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sys2 *Sys2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sys2.Contract.contract.Transact(opts, method, params...)
}

// GetStorageAt is a free data retrieval call binding the contract method 0xa169625f.
//
// Solidity: function getStorageAt(address account, uint256 index) view returns(uint256)
func (_Sys2 *Sys2Caller) GetStorageAt(opts *bind.CallOpts, account common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sys2.contract.Call(opts, &out, "getStorageAt", account, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0xa169625f.
//
// Solidity: function getStorageAt(address account, uint256 index) view returns(uint256)
func (_Sys2 *Sys2Session) GetStorageAt(account common.Address, index *big.Int) (*big.Int, error) {
	return _Sys2.Contract.GetStorageAt(&_Sys2.CallOpts, account, index)
}

// GetStorageAt is a free data retrieval call binding the contract method 0xa169625f.
//
// Solidity: function getStorageAt(address account, uint256 index) view returns(uint256)
func (_Sys2 *Sys2CallerSession) GetStorageAt(account common.Address, index *big.Int) (*big.Int, error) {
	return _Sys2.Contract.GetStorageAt(&_Sys2.CallOpts, account, index)
}
