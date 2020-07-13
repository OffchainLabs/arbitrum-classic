// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onestepprooftester

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

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209cc6ea8cfb5d0f6e66ccce67c7494628093cb8581f492f996c0110163c56d51b64736f6c63430005110032"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[]"

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820f1f511f3dfbb775ccca8f03416243a8f73f5f620d61b74b943ba2554e9fdc65a64736f6c63430005110032"

// DeployDebugPrint deploys a new Ethereum contract, binding an instance of DebugPrint to it.
func DeployDebugPrint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebugPrint, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DebugPrintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// DebugPrint is an auto generated Go binding around an Ethereum contract.
type DebugPrint struct {
	DebugPrintCaller     // Read-only binding to the contract
	DebugPrintTransactor // Write-only binding to the contract
	DebugPrintFilterer   // Log filterer for contract events
}

// DebugPrintCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebugPrintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebugPrintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebugPrintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebugPrintSession struct {
	Contract     *DebugPrint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebugPrintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebugPrintCallerSession struct {
	Contract *DebugPrintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DebugPrintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebugPrintTransactorSession struct {
	Contract     *DebugPrintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DebugPrintRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebugPrintRaw struct {
	Contract *DebugPrint // Generic contract binding to access the raw methods on
}

// DebugPrintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebugPrintCallerRaw struct {
	Contract *DebugPrintCaller // Generic read-only contract binding to access the raw methods on
}

// DebugPrintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebugPrintTransactorRaw struct {
	Contract *DebugPrintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebugPrint creates a new instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrint(address common.Address, backend bind.ContractBackend) (*DebugPrint, error) {
	contract, err := bindDebugPrint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// NewDebugPrintCaller creates a new read-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintCaller(address common.Address, caller bind.ContractCaller) (*DebugPrintCaller, error) {
	contract, err := bindDebugPrint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintCaller{contract: contract}, nil
}

// NewDebugPrintTransactor creates a new write-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintTransactor(address common.Address, transactor bind.ContractTransactor) (*DebugPrintTransactor, error) {
	contract, err := bindDebugPrint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintTransactor{contract: contract}, nil
}

// NewDebugPrintFilterer creates a new log filterer instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintFilterer(address common.Address, filterer bind.ContractFilterer) (*DebugPrintFilterer, error) {
	contract, err := bindDebugPrint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebugPrintFilterer{contract: contract}, nil
}

// bindDebugPrint binds a generic wrapper to an already deployed contract.
func bindDebugPrint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.DebugPrintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transact(opts, method, params...)
}

// HashingABI is the input ABI used to generate the binding from.
const HashingABI = "[]"

// HashingBin is the compiled bytecode used for deploying new contracts.
var HashingBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820533a398cab9ecff892064e05f3998adba897326c113e2afc9a01e5783f6d05c864736f6c63430005110032"

// DeployHashing deploys a new Ethereum contract, binding an instance of Hashing to it.
func DeployHashing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hashing, error) {
	parsed, err := abi.JSON(strings.NewReader(HashingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HashingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hashing{HashingCaller: HashingCaller{contract: contract}, HashingTransactor: HashingTransactor{contract: contract}, HashingFilterer: HashingFilterer{contract: contract}}, nil
}

// Hashing is an auto generated Go binding around an Ethereum contract.
type Hashing struct {
	HashingCaller     // Read-only binding to the contract
	HashingTransactor // Write-only binding to the contract
	HashingFilterer   // Log filterer for contract events
}

// HashingCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashingSession struct {
	Contract     *Hashing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashingCallerSession struct {
	Contract *HashingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HashingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashingTransactorSession struct {
	Contract     *HashingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HashingRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashingRaw struct {
	Contract *Hashing // Generic contract binding to access the raw methods on
}

// HashingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashingCallerRaw struct {
	Contract *HashingCaller // Generic read-only contract binding to access the raw methods on
}

// HashingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashingTransactorRaw struct {
	Contract *HashingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashing creates a new instance of Hashing, bound to a specific deployed contract.
func NewHashing(address common.Address, backend bind.ContractBackend) (*Hashing, error) {
	contract, err := bindHashing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hashing{HashingCaller: HashingCaller{contract: contract}, HashingTransactor: HashingTransactor{contract: contract}, HashingFilterer: HashingFilterer{contract: contract}}, nil
}

// NewHashingCaller creates a new read-only instance of Hashing, bound to a specific deployed contract.
func NewHashingCaller(address common.Address, caller bind.ContractCaller) (*HashingCaller, error) {
	contract, err := bindHashing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashingCaller{contract: contract}, nil
}

// NewHashingTransactor creates a new write-only instance of Hashing, bound to a specific deployed contract.
func NewHashingTransactor(address common.Address, transactor bind.ContractTransactor) (*HashingTransactor, error) {
	contract, err := bindHashing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashingTransactor{contract: contract}, nil
}

// NewHashingFilterer creates a new log filterer instance of Hashing, bound to a specific deployed contract.
func NewHashingFilterer(address common.Address, filterer bind.ContractFilterer) (*HashingFilterer, error) {
	contract, err := bindHashing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashingFilterer{contract: contract}, nil
}

// bindHashing binds a generic wrapper to an already deployed contract.
func bindHashing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashing *HashingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hashing.Contract.HashingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashing *HashingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashing.Contract.HashingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashing *HashingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashing.Contract.HashingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashing *HashingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hashing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashing *HashingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashing *HashingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashing.Contract.contract.Transact(opts, method, params...)
}

// MachineABI is the input ABI used to generate the binding from.
const MachineABI = "[]"

// MachineBin is the compiled bytecode used for deploying new contracts.
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820e26a47154592b98a0b2092aaadfa56172623e07699faa9e53369e06767bd8d8a64736f6c63430005110032"

// DeployMachine deploys a new Ethereum contract, binding an instance of Machine to it.
func DeployMachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Machine, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Machine{MachineCaller: MachineCaller{contract: contract}, MachineTransactor: MachineTransactor{contract: contract}, MachineFilterer: MachineFilterer{contract: contract}}, nil
}

// Machine is an auto generated Go binding around an Ethereum contract.
type Machine struct {
	MachineCaller     // Read-only binding to the contract
	MachineTransactor // Write-only binding to the contract
	MachineFilterer   // Log filterer for contract events
}

// MachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineSession struct {
	Contract     *Machine          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineCallerSession struct {
	Contract *MachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineTransactorSession struct {
	Contract     *MachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineRaw struct {
	Contract *Machine // Generic contract binding to access the raw methods on
}

// MachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineCallerRaw struct {
	Contract *MachineCaller // Generic read-only contract binding to access the raw methods on
}

// MachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineTransactorRaw struct {
	Contract *MachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachine creates a new instance of Machine, bound to a specific deployed contract.
func NewMachine(address common.Address, backend bind.ContractBackend) (*Machine, error) {
	contract, err := bindMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Machine{MachineCaller: MachineCaller{contract: contract}, MachineTransactor: MachineTransactor{contract: contract}, MachineFilterer: MachineFilterer{contract: contract}}, nil
}

// NewMachineCaller creates a new read-only instance of Machine, bound to a specific deployed contract.
func NewMachineCaller(address common.Address, caller bind.ContractCaller) (*MachineCaller, error) {
	contract, err := bindMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineCaller{contract: contract}, nil
}

// NewMachineTransactor creates a new write-only instance of Machine, bound to a specific deployed contract.
func NewMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineTransactor, error) {
	contract, err := bindMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineTransactor{contract: contract}, nil
}

// NewMachineFilterer creates a new log filterer instance of Machine, bound to a specific deployed contract.
func NewMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineFilterer, error) {
	contract, err := bindMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineFilterer{contract: contract}, nil
}

// bindMachine binds a generic wrapper to an already deployed contract.
func bindMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machine *MachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Machine.Contract.MachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machine *MachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machine.Contract.MachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machine *MachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machine.Contract.MachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machine *MachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Machine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machine *MachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machine *MachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machine.Contract.contract.Transact(opts, method, params...)
}

// MarshalingABI is the input ABI used to generate the binding from.
const MarshalingABI = "[]"

// MarshalingBin is the compiled bytecode used for deploying new contracts.
var MarshalingBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582070fef28954c64e536133f92ce88cf591d83cfd4bf8678a07d7f000cf47a5c5ca64736f6c63430005110032"

// DeployMarshaling deploys a new Ethereum contract, binding an instance of Marshaling to it.
func DeployMarshaling(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Marshaling, error) {
	parsed, err := abi.JSON(strings.NewReader(MarshalingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarshalingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Marshaling{MarshalingCaller: MarshalingCaller{contract: contract}, MarshalingTransactor: MarshalingTransactor{contract: contract}, MarshalingFilterer: MarshalingFilterer{contract: contract}}, nil
}

// Marshaling is an auto generated Go binding around an Ethereum contract.
type Marshaling struct {
	MarshalingCaller     // Read-only binding to the contract
	MarshalingTransactor // Write-only binding to the contract
	MarshalingFilterer   // Log filterer for contract events
}

// MarshalingCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarshalingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarshalingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarshalingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarshalingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarshalingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarshalingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarshalingSession struct {
	Contract     *Marshaling       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarshalingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarshalingCallerSession struct {
	Contract *MarshalingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MarshalingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarshalingTransactorSession struct {
	Contract     *MarshalingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MarshalingRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarshalingRaw struct {
	Contract *Marshaling // Generic contract binding to access the raw methods on
}

// MarshalingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarshalingCallerRaw struct {
	Contract *MarshalingCaller // Generic read-only contract binding to access the raw methods on
}

// MarshalingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarshalingTransactorRaw struct {
	Contract *MarshalingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarshaling creates a new instance of Marshaling, bound to a specific deployed contract.
func NewMarshaling(address common.Address, backend bind.ContractBackend) (*Marshaling, error) {
	contract, err := bindMarshaling(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marshaling{MarshalingCaller: MarshalingCaller{contract: contract}, MarshalingTransactor: MarshalingTransactor{contract: contract}, MarshalingFilterer: MarshalingFilterer{contract: contract}}, nil
}

// NewMarshalingCaller creates a new read-only instance of Marshaling, bound to a specific deployed contract.
func NewMarshalingCaller(address common.Address, caller bind.ContractCaller) (*MarshalingCaller, error) {
	contract, err := bindMarshaling(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarshalingCaller{contract: contract}, nil
}

// NewMarshalingTransactor creates a new write-only instance of Marshaling, bound to a specific deployed contract.
func NewMarshalingTransactor(address common.Address, transactor bind.ContractTransactor) (*MarshalingTransactor, error) {
	contract, err := bindMarshaling(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarshalingTransactor{contract: contract}, nil
}

// NewMarshalingFilterer creates a new log filterer instance of Marshaling, bound to a specific deployed contract.
func NewMarshalingFilterer(address common.Address, filterer bind.ContractFilterer) (*MarshalingFilterer, error) {
	contract, err := bindMarshaling(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarshalingFilterer{contract: contract}, nil
}

// bindMarshaling binds a generic wrapper to an already deployed contract.
func bindMarshaling(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarshalingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marshaling *MarshalingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Marshaling.Contract.MarshalingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marshaling *MarshalingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marshaling.Contract.MarshalingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marshaling *MarshalingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marshaling.Contract.MarshalingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marshaling *MarshalingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Marshaling.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marshaling *MarshalingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marshaling.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marshaling *MarshalingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marshaling.Contract.contract.Transact(opts, method, params...)
}

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[]"

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201e6806e3bd9c1c2b0a9f47eac003fc9ad4476831be3cc6d5c604f1121c999c8564736f6c63430005110032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// OneStepProof is an auto generated Go binding around an Ethereum contract.
type OneStepProof struct {
	OneStepProofCaller     // Read-only binding to the contract
	OneStepProofTransactor // Write-only binding to the contract
	OneStepProofFilterer   // Log filterer for contract events
}

// OneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofSession struct {
	Contract     *OneStepProof     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofCallerSession struct {
	Contract *OneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofTransactorSession struct {
	Contract     *OneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofRaw struct {
	Contract *OneStepProof // Generic contract binding to access the raw methods on
}

// OneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofCallerRaw struct {
	Contract *OneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofTransactorRaw struct {
	Contract *OneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof creates a new instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProof(address common.Address, backend bind.ContractBackend) (*OneStepProof, error) {
	contract, err := bindOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// NewOneStepProofCaller creates a new read-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofCaller, error) {
	contract, err := bindOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofCaller{contract: contract}, nil
}

// NewOneStepProofTransactor creates a new write-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofTransactor, error) {
	contract, err := bindOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTransactor{contract: contract}, nil
}

// NewOneStepProofFilterer creates a new log filterer instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofFilterer, error) {
	contract, err := bindOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofFilterer{contract: contract}, nil
}

// bindOneStepProof binds a generic wrapper to an already deployed contract.
func bindOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.OneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transact(opts, method, params...)
}

// OneStepProofTesterABI is the input ABI used to generate the binding from.
const OneStepProofTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxValueSize\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofTesterFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofTesterFuncSigs = map[string]string{
	"3c41485d": "validateProof(bytes32,bytes32,uint256,bool,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
}

// OneStepProofTesterBin is the compiled bytecode used for deploying new contracts.
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b50613c5d806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633c41485d14610030575b600080fd5b61011d600480360361014081101561004757600080fd5b813591602081013591604082013591606081013515159160808201359160a08101359160c08201359160e08101359167ffffffffffffffff61010083013516919081019061014081016101208201356401000000008111156100a857600080fd5b8201836020820111156100ba57600080fd5b803590602001918460018302840111640100000000831117156100dc57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061012f945050505050565b60408051918252519081900360200190f35b600061014b6101468c8c8c8c8c8c8c8c8c8c61015a565b6101bf565b9b9a5050505050505050505050565b610162613a61565b61014b6040518061012001604052808d81526020016101818d8d610284565b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610337565b600060028260e0015114156101d65750600061027f565b60018260e0015114156101eb5750600161027f565b815160208301516101fb90611487565b6102088460400151611487565b6102158560600151611487565b6102228660800151611487565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090505b919050565b61028c613abf565b6040805160a08101825284815281516060810183526000808252602082810182905284518281528082018652939490850193908301916102e2565b6102cf613abf565b8152602001906001900390816102c75790505b5090528152604080516000808252602082810190935291909201919061031e565b61030b613abf565b8152602001906001900390816103035790505b5081526002602082015260400183905290505b92915050565b61033f613a61565b6000806000606061034e613a61565b610356613a61565b61035f886115a7565b60e08e0151959b509399509297509095509350915060019060009067ffffffffffffffff1687146103ce576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b896040015180156103e2575060ff88166072145b806103fe575089604001511580156103fe575060ff8816607214155b61044f576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a0808401805189900390528401518711156104765760001960a084015260009150611340565b60ff8816600114156104bc576104b5838660008151811061049357fe5b6020026020010151876001815181106104a857fe5b6020026020010151611802565b9150611340565b60ff8816600214156104fb576104b583866000815181106104d957fe5b6020026020010151876001815181106104ee57fe5b6020026020010151611852565b60ff88166003141561053a576104b5838660008151811061051857fe5b60200260200101518760018151811061052d57fe5b6020026020010151611893565b60ff881660041415610579576104b5838660008151811061055757fe5b60200260200101518760018151811061056c57fe5b60200260200101516118d4565b60ff8816600514156105b8576104b5838660008151811061059657fe5b6020026020010151876001815181106105ab57fe5b6020026020010151611925565b60ff8816600614156105f7576104b583866000815181106105d557fe5b6020026020010151876001815181106105ea57fe5b6020026020010151611976565b60ff881660071415610636576104b5838660008151811061061457fe5b60200260200101518760018151811061062957fe5b60200260200101516119c7565b60ff88166008141561068a576104b5838660008151811061065357fe5b60200260200101518760018151811061066857fe5b60200260200101518860028151811061067d57fe5b6020026020010151611a18565b60ff8816600914156106de576104b583866000815181106106a757fe5b6020026020010151876001815181106106bc57fe5b6020026020010151886002815181106106d157fe5b6020026020010151611a82565b60ff8816600a141561071d576104b583866000815181106106fb57fe5b60200260200101518760018151811061071057fe5b6020026020010151611adb565b60ff88166010141561075c576104b5838660008151811061073a57fe5b60200260200101518760018151811061074f57fe5b6020026020010151611b1c565b60ff88166011141561079b576104b5838660008151811061077957fe5b60200260200101518760018151811061078e57fe5b6020026020010151611b5d565b60ff8816601214156107da576104b583866000815181106107b857fe5b6020026020010151876001815181106107cd57fe5b6020026020010151611b9e565b60ff881660131415610819576104b583866000815181106107f757fe5b60200260200101518760018151811061080c57fe5b6020026020010151611bdf565b60ff881660141415610858576104b5838660008151811061083657fe5b60200260200101518760018151811061084b57fe5b6020026020010151611c20565b60ff881660151415610882576104b5838660008151811061087557fe5b6020026020010151611c57565b60ff8816601614156108c1576104b5838660008151811061089f57fe5b6020026020010151876001815181106108b457fe5b6020026020010151611c9c565b60ff881660171415610900576104b583866000815181106108de57fe5b6020026020010151876001815181106108f357fe5b6020026020010151611cdd565b60ff88166018141561093f576104b5838660008151811061091d57fe5b60200260200101518760018151811061093257fe5b6020026020010151611d1e565b60ff881660191415610969576104b5838660008151811061095c57fe5b6020026020010151611d5f565b60ff8816601a14156109a8576104b5838660008151811061098657fe5b60200260200101518760018151811061099b57fe5b6020026020010151611d95565b60ff8816601b14156109e7576104b583866000815181106109c557fe5b6020026020010151876001815181106109da57fe5b6020026020010151611dd6565b60ff881660201415610a11576104b58386600081518110610a0457fe5b6020026020010151611e17565b60ff881660211415610a3b576104b58386600081518110610a2e57fe5b6020026020010151611e32565b60ff881660221415610a7a576104b58386600081518110610a5857fe5b602002602001015187600181518110610a6d57fe5b6020026020010151611e4d565b60ff881660301415610aa4576104b58386600081518110610a9757fe5b6020026020010151611eb3565b60ff881660311415610ab9576104b583611ebb565b60ff881660321415610ace576104b583611edc565b60ff881660331415610af8576104b58386600081518110610aeb57fe5b6020026020010151611ef5565b60ff881660341415610b22576104b58386600081518110610b1557fe5b6020026020010151611f01565b60ff881660351415610b61576104b58386600081518110610b3f57fe5b602002602001015187600181518110610b5457fe5b6020026020010151611f2c565b60ff881660361415610b76576104b583611f74565b60ff881660371415610b90576104b5838560000151611fa1565b60ff881660381415610bba576104b58386600081518110610bad57fe5b6020026020010151611fb1565b60ff881660391415610c0857610bce613abf565b610bdd8b610100015188611fc3565b9097509050610bf2858263ffffffff61215d16565b610c02848263ffffffff61217716565b50611340565b60ff8816603a1415610c1d576104b583612191565b60ff8816603b1415610c325760019150611340565b60ff8816603c1415610c47576104b5836121b1565b60ff8816603d1415610c71576104b58386600081518110610c6457fe5b60200260200101516121c5565b60ff881660401415610c9b576104b58386600081518110610c8e57fe5b60200260200101516121f3565b60ff881660411415610cda576104b58386600081518110610cb857fe5b602002602001015187600181518110610ccd57fe5b6020026020010151612215565b60ff881660421415610d2e576104b58386600081518110610cf757fe5b602002602001015187600181518110610d0c57fe5b602002602001015188600281518110610d2157fe5b6020026020010151612247565b60ff881660431415610d6d576104b58386600081518110610d4b57fe5b602002602001015187600181518110610d6057fe5b6020026020010151612289565b60ff881660441415610dc1576104b58386600081518110610d8a57fe5b602002602001015187600181518110610d9f57fe5b602002602001015188600281518110610db457fe5b602002602001015161229b565b60ff881660501415610e00576104b58386600081518110610dde57fe5b602002602001015187600181518110610df357fe5b60200260200101516122bd565b60ff881660511415610e54576104b58386600081518110610e1d57fe5b602002602001015187600181518110610e3257fe5b602002602001015188600281518110610e4757fe5b6020026020010151612333565b60ff881660521415610e7e576104b58386600081518110610e7157fe5b60200260200101516123c0565b60ff881660531415610edd57610e92613abf565b610ea18b610100015188611fc3565b9097509050610eb6858263ffffffff61215d16565b610ed58487600081518110610ec757fe5b6020026020010151836123f3565b925050611340565b60ff881660541415610f4957610ef1613abf565b610f008b610100015188611fc3565b9097509050610f15858263ffffffff61215d16565b610ed58487600081518110610f2657fe5b602002602001015188600181518110610f3b57fe5b60200260200101518461244b565b60ff881660601415610f5e576104b5836124cc565b60ff88166061141561105b57610f888386600081518110610f7b57fe5b60200260200101516124d2565b90925090508115611052578960c001518a60a0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146110075760405162461bcd60e51b8152600401808060200182810382526025815260200180613bdd6025913960400191505060405180910390fd5b89608001518a606001511461104d5760405162461bcd60e51b8152600401808060200182810382526027815260200180613c026027913960400191505060405180910390fd5b611056565b5060005b611340565b60ff88166070141561119257611085838660008151811061107857fe5b60200260200101516124eb565b9092509050811561105257806110e05789608001518a60600151146110db5760405162461bcd60e51b8152600401808060200182810382526038815260200180613ba56038913960400191505060405180910390fd5b61104d565b60808a01516060808c0151604080516020808201939093528082018690528151808203830181529301905281519101201461114c5760405162461bcd60e51b8152600401808060200182810382526029815260200180613b356029913960400191505060405180910390fd5b8960c001518a60a001511461104d5760405162461bcd60e51b8152600401808060200182810382526026815260200180613b5e6026913960400191505060405180910390fd5b60ff8816607214156111ac576104b5838b6020015161252a565b60ff8816607314156111c15760009150611340565b60ff8816607414156111d65761105683612593565b60ff881660751415611200576104b583866000815181106111f357fe5b602002602001015161259d565b60ff881660761415611215576104b5836125c2565b60ff88166077141561122a576104b5836125db565b60ff881660781415611269576104b5838660008151811061124757fe5b60200260200101518760018151811061125c57fe5b6020026020010151612624565b60ff8816607914156112bd576104b5838660008151811061128657fe5b60200260200101518760018151811061129b57fe5b6020026020010151886002815181106112b057fe5b6020026020010151612669565b60ff8816607b14156112d2576104b5836126bc565b60ff88166080141561133b576104b583866000815181106112ef57fe5b60200260200101518760018151811061130457fe5b60200260200101518860028151811061131957fe5b60200260200101518960038151811061132e57fe5b60200260200101516126ff565b600091505b806113d15789608001518a606001511461138b5760405162461bcd60e51b8152600401808060200182810382526027815260200180613c026027913960400191505060405180910390fd5b8960c001518a60a00151146113d15760405162461bcd60e51b8152600401808060200182810382526026815260200180613b5e6026913960400191505060405180910390fd5b816114325760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c0840151141561142a576114258361281e565b611432565b60c083015183525b61143b846101bf565b8a51146114795760405162461bcd60e51b8152600401808060200182810382526022815260200180613b136022913960400191505060405180910390fd5b509098975050505050505050565b6000611491612828565b60ff16826060015160ff1614156114b45781516114ad9061282d565b905061027f565b6114bc612851565b60ff16826060015160ff1614156114da576114ad8260200151612856565b6114e2612953565b60ff16826060015160ff16141561150457815160808301516114ad9190612958565b61150c6129a9565b60ff16826060015160ff16141561154557611525613abf565b61153283604001516129ae565b905061153d81611487565b91505061027f565b61154d612b10565b60ff16826060015160ff1614156115665750805161027f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b60008060606115b4613a61565b6115bc613a61565b60006115c783612b15565b6115d687610100015182612b1f565b935090506115e383612bbf565b9150600087610100015182815181106115f857fe5b602001015160f81c60f81b60f81c9050876101000151826001018151811061161c57fe5b016020015160f81c9650600061163188612c28565b60408051838152602080850282010190915290985090915081801561167057816020015b61165d613abf565b8152602001906001900390816116555790505b5095506002830192508160ff166000148061168e57508160ff166001145b6116df576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff8216611704576116fd6116f8898760000151613185565b611487565b855261176a565b61170c613abf565b61171b8a610100015185611fc3565b9094509050811561174457808760008151811061173457fe5b6020026020010181905250611754565b611754858263ffffffff61217716565b6117666116f88a8860000151846131e0565b8652505b60ff82165b818110156117a9576117868a610100015185611fc3565b885189908490811061179457fe5b6020908102919091010152935060010161176f565b8651156117f6575060005b8260ff168751038110156117f6576117ee878260018a510303815181106117d757fe5b60200260200101518761217790919063ffffffff16565b6001016117b4565b50505091939550919395565b600061180d8361325b565b158061181f575061181d8261325b565b155b1561182c5750600061184b565b82518251808201611843878263ffffffff61326616565b600193505050505b9392505050565b600061185d8361325b565b158061186f575061186d8261325b565b155b1561187c5750600061184b565b82518251808202611843878263ffffffff61326616565b600061189e8361325b565b15806118b057506118ae8261325b565b155b156118bd5750600061184b565b82518251808203611843878263ffffffff61326616565b60006118df8361325b565b15806118f157506118ef8261325b565b155b156118fe5750600061184b565b82518251806119125760009250505061184b565b808204611843878263ffffffff61326616565b60006119308361325b565b158061194257506119408261325b565b155b1561194f5750600061184b565b82518251806119635760009250505061184b565b808205611843878263ffffffff61326616565b60006119818361325b565b158061199357506119918261325b565b155b156119a05750600061184b565b82518251806119b45760009250505061184b565b808206611843878263ffffffff61326616565b60006119d28361325b565b15806119e457506119e28261325b565b155b156119f15750600061184b565b8251825180611a055760009250505061184b565b808207611843878263ffffffff61326616565b6000611a238461325b565b1580611a355750611a338361325b565b155b15611a4257506000611a7a565b83518351835180611a595760009350505050611a7a565b6000818385089050611a71898263ffffffff61326616565b60019450505050505b949350505050565b6000611a8d8461325b565b1580611a9f5750611a9d8361325b565b155b15611aac57506000611a7a565b83518351835180611ac35760009350505050611a7a565b6000818385099050611a71898263ffffffff61326616565b6000611ae68361325b565b1580611af85750611af68261325b565b155b15611b055750600061184b565b8251825180820a611843878263ffffffff61326616565b6000611b278361325b565b1580611b395750611b378261325b565b155b15611b465750600061184b565b82518251808210611843878263ffffffff61326616565b6000611b688361325b565b1580611b7a5750611b788261325b565b155b15611b875750600061184b565b82518251808211611843878263ffffffff61326616565b6000611ba98361325b565b1580611bbb5750611bb98261325b565b155b15611bc85750600061184b565b82518251808212611843878263ffffffff61326616565b6000611bea8361325b565b1580611bfc5750611bfa8261325b565b155b15611c095750600061184b565b82518251808213611843878263ffffffff61326616565b6000611c4d611c40611c3184611487565b611c3a86611487565b1461327c565b859063ffffffff61217716565b5060019392505050565b6000611c628261325b565b611c7c57611c7783600063ffffffff61326616565b611c93565b81518015611c90858263ffffffff61326616565b50505b50600192915050565b6000611ca78361325b565b1580611cb95750611cb78261325b565b155b15611cc65750600061184b565b82518251808216611843878263ffffffff61326616565b6000611ce88361325b565b1580611cfa5750611cf88261325b565b155b15611d075750600061184b565b82518251808217611843878263ffffffff61326616565b6000611d298361325b565b1580611d3b5750611d398261325b565b155b15611d485750600061184b565b82518251808218611843878263ffffffff61326616565b6000611d6a8261325b565b611d7657506000610331565b81518019611d8a858263ffffffff61326616565b506001949350505050565b6000611da08361325b565b1580611db25750611db08261325b565b155b15611dbf5750600061184b565b8251825181811a611843878263ffffffff61326616565b6000611de18361325b565b1580611df35750611df18261325b565b155b15611e005750600061184b565b8251825181810b611843878263ffffffff61326616565b6000611c93611e2583611487565b849063ffffffff61326616565b6000611c93611e408361329e565b849063ffffffff61217716565b6000611e588361325b565b1580611e6a5750611e688261325b565b155b15611e775750600061184b565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120611843878263ffffffff61326616565b600192915050565b6000611ed482608001518361217790919063ffffffff16565b506001919050565b6000611ed482606001518361217790919063ffffffff16565b60609190910152600190565b6000611f0c826132fd565b611f1857506000610331565b611f2182611487565b835250600192915050565b6000611f37836132fd565b611f435750600061184b565b611f4c8261325b565b611f585750600061184b565b815115611c4d57611f6883611487565b84525060019392505050565b6000611ed4611f94611f876116f861330a565b611c3a8560200151611487565b839063ffffffff61217716565b6000611c93611e40836001613351565b6000611c93838363ffffffff61215d16565b6000611fcd613abf565b83518310612013576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000839050600085828151811061202657fe5b016020015160019092019160f81c90506000612040612828565b60ff168260ff161415612072576120578784613402565b9093509050826120668261346b565b94509450505050612156565b61207a612851565b60ff168260ff16141561209157612066878461351d565b612099612953565b60ff168260ff1614156120b05761206687846135e1565b6120b86129a9565b60ff168260ff16101580156120d957506120d061367f565b60ff168260ff16105b156121165760006120e86129a9565b8303905060606120f9828a87613684565b9095509050846121088261371d565b965096505050505050612156565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b61216b82604001518261382f565b82604001819052505050565b61218582602001518261382f565b82602001819052505050565b6000611ed4611f946121a46116f861330a565b611c3a8560400151611487565b6000611ed4611f948360c001516001613351565b60006121d0826132fd565b6121dc57506000610331565b6121e582611487565b60c084015250600192915050565b6000612205838363ffffffff61217716565b611c93838363ffffffff61217716565b6000612227848363ffffffff61217716565b612237848463ffffffff61217716565b611c4d848363ffffffff61217716565b6000612259858363ffffffff61217716565b612269858463ffffffff61217716565b612279858563ffffffff61217716565b611d8a858363ffffffff61217716565b6000612237848463ffffffff61217716565b60006122ad858563ffffffff61217716565b612279858463ffffffff61217716565b60006122c88361325b565b15806122da57506122d8826138a5565b155b156122e75750600061184b565b6122f0826138b2565b60ff168360000151106123055750600061184b565b611c4d826040015184600001518151811061231c57fe5b60200260200101518561217790919063ffffffff16565b600061233e836138a5565b1580612350575061234e8461325b565b155b1561235d57506000611a7a565b612366836138b2565b60ff1684600001511061237b57506000611a7a565b60408301518451815184918391811061239057fe5b60200260200101819052506123b46123a78261371d565b879063ffffffff61217716565b50600195945050505050565b60006123cb826138a5565b6123d757506000610331565b611c936123e3836138b2565b849060ff1663ffffffff61326616565b60006123fe8361325b565b1580612410575061240e826138a5565b155b1561241d5750600061184b565b612426826138b2565b60ff1683600001511061243b5750600061184b565b612305848363ffffffff61215d16565b6000612456826138a5565b158061246857506124668461325b565b155b1561247557506000611a7a565b61247e826138b2565b60ff1684600001511061249357506000611a7a565b6040820151845181518591839181106124a857fe5b60200260200101819052506123b46124bf8261371d565b879063ffffffff61215d16565b50600190565b60008060016124e084611487565b915091509250929050565b6000806127108360800151111561250757506000905080612156565b612510836138d9565b61251f57506000905080612156565b60016124e084611487565b60006125376116f861330a565b61254083611487565b1415612205576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b600260e090910152565b60006125a88261325b565b6125b457506000610331565b505160a09190910152600190565b6000611ed48260a001518361326690919063ffffffff16565b60408051600160f81b60208083019190915260006021830181905260228084018290528451808503909101815260429093019093528151910120611ed490611f94906001613351565b600061262f8361325b565b61263b5750600061184b565b612644826132fd565b6126505750600061184b565b611c4d611c40846000015161266485611487565b613185565b60006126748461325b565b61268057506000611a7a565b612689826132fd565b61269557506000611a7a565b611d8a6126af85600001516126a985611487565b866131e0565b869063ffffffff61217716565b60408051600080825260208201909252606090826126f0565b6126dd613abf565b8152602001906001900390816126d55790505b509050611c93611e408261371d565b600061270a8561325b565b158061271c575061271a8461325b565b155b8061272d575061272b8361325b565b155b8061273e575061273c8261325b565b155b1561274b57506000612815565b8451845184511580159061276157508451600114155b156127825761277788600063ffffffff61326616565b600192505050612815565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa1580156127e4573d6000803e3d6000fd5b5050604051601f190151915061280b90508b6001600160a01b03831663ffffffff61326616565b6001955050505050505b95945050505050565b600160e090910152565b600090565b60408051602080820193909352815180820384018152908201909152805191012090565b600190565b600060028260400151511061286757fe5b6040820151516128cc57612879612851565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b90931660218501526022808501919091528251808503909101815260429093019091528151910120905061027f565b6128d4612851565b82600001516128fa84604001516000815181106128ed57fe5b6020026020010151611487565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b600290565b60006129626129a9565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b600390565b6129b6613abf565b600882511115612a04576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015612a31578160200160208202803883390190505b508051909150600160005b82811015612a9457612a538682815181106128ed57fe5b848281518110612a5f57fe5b602002602001018181525050858181518110612a7757fe5b602002602001015160800151820191508080600101915050612a3c565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015612ad9578181015183820152602001612ac1565b5050505090500192505050604051602081830303815290604052805190602001209050612b068183610284565b9695505050505050565b606490565b600060e090910152565b6000612b29613a61565b612b31613a61565b600060e0820181905280612b458787613402565b9096509150612b5487876135e1565b60208501529550612b6587876135e1565b60408501529550612b768787611fc3565b60608501529550612b878787611fc3565b60808501529550612b988787613402565b60a08501529550612ba98787613402565b92845260c0840192909252509590945092505050565b612bc7613a61565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612c405750600290506003613180565b6002831415612c555750600290506003613180565b6003831415612c6a5750600290506003613180565b6004831415612c7f5750600290506004613180565b6005831415612c945750600290506007613180565b6006831415612ca95750600290506004613180565b6007831415612cbe5750600290506007613180565b6008831415612cd35750600390506004613180565b6009831415612ce85750600390506004613180565b600a831415612cfd5750600290506019613180565b6010831415612d1157506002905080613180565b6011831415612d2557506002905080613180565b6012831415612d3957506002905080613180565b6013831415612d4d57506002905080613180565b6014831415612d6157506002905080613180565b6015831415612d7557506001905080613180565b6016831415612d8957506002905080613180565b6017831415612d9d57506002905080613180565b6018831415612db157506002905080613180565b6019831415612dc557506001905080613180565b601a831415612dda5750600290506004613180565b601b831415612def5750600290506007613180565b6020831415612e045750600190506007613180565b6021831415612e195750600190506003613180565b6022831415612e2e5750600290506008613180565b6030831415612e4257506001905080613180565b6031831415612e575750600090506001613180565b6032831415612e6c5750600090506001613180565b6033831415612e815750600190506002613180565b6034831415612e965750600190506004613180565b6035831415612eab5750600290506004613180565b6036831415612ec05750600090506002613180565b6037831415612ed55750600090506001613180565b6038831415612ee957506001905080613180565b6039831415612efe5750600090506001613180565b603a831415612f135750600090506002613180565b603b831415612f285750600090506001613180565b603c831415612f3d5750600090506001613180565b603d831415612f5157506001905080613180565b6040831415612f6557506001905080613180565b6041831415612f7a5750600290506001613180565b6042831415612f8f5750600390506001613180565b6043831415612fa45750600290506001613180565b6044831415612fb95750600390506001613180565b6050831415612fcd57506002905080613180565b6051831415612fe25750600390506028613180565b6052831415612ff75750600190506002613180565b605383141561300c5750600190506003613180565b60548314156130215750600290506029613180565b60608314156130365750600090506064613180565b606183141561304b5750600190506064613180565b60708314156130605750600190506064613180565b60728314156130755750600090506028613180565b607383141561308a5750600090506005613180565b607483141561309f575060009050600a613180565b60758314156130b45750600190506000613180565b60768314156130c95750600090506001613180565b60778314156130de5750600090506019613180565b60788314156130f35750600290506019613180565b60798314156131085750600390506019613180565b607b83141561311d575060009050600a613180565b6080831415613133575060049050614e20613180565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b61318d613abf565b6040805160608101825260ff85168152602080820185905282516000808252918101845261184b938301916131d8565b6131c5613abf565b8152602001906001900390816131bd5790505b5090526139d7565b6131e8613abf565b604080516001808252818301909252606091816020015b613207613abf565b8152602001906001900390816131ff579050509050828160008151811061322a57fe5b602002602001018190525061281560405180606001604052808760ff168152602001868152602001838152506139d7565b6060015160ff161590565b61218582602001516132778361346b565b61382f565b613284613abf565b8115613294576114ad600161346b565b6114ad600061346b565b6132a6613abf565b816060015160ff16600214156132ed5760405162461bcd60e51b8152600401808060200182810382526021815260200180613b846021913960400191505060405180910390fd5b610331826060015160ff1661346b565b6060015160ff1660011490565b613312613abf565b6040805160008082526020820190925261334c91613346565b613333613abf565b81526020019060019003908161332b5790505b5061371d565b905090565b613359613abf565b6040805160a08101825284815281516060810183526000808252602082810182905284518281528082018652939490850193908301916133af565b61339c613abf565b8152602001906001900390816133945790505b509052815260408051600080825260208281019093529190920191906133eb565b6133d8613abf565b8152602001906001900390816133d05790505b508152606460208201526040019290925250919050565b6000808284511015801561341a575060208385510310155b613457576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b602083016124e0858563ffffffff613a3e16565b613473613abf565b6040805160a08101825283815281516060810183526000808252602082810182905284518281528082018652939490850193908301916134c9565b6134b6613abf565b8152602001906001900390816134ae5790505b50905281526040805160008082526020828101909352919092019190613505565b6134f2613abf565b8152602001906001900390816134ea5790505b50815260006020820152600160409091015292915050565b6000613527613abf565b6000839050600085828151811061353a57fe5b602001015160f81c60f81b60f81c90508180600101925050600086838151811061356057fe5b016020015160019093019260f81c9050613578613abf565b8260ff16600114156135945761358e8885611fc3565b90945090505b60006135a6898663ffffffff613a3e16565b90506020850194508360ff16600114156135c657846121088483856131e0565b846135d18483613185565b9650965050505050509250929050565b60006135eb613abf565b600083855110158015613602575060408486510310155b61363e576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b6000613650868663ffffffff613a3e16565b90506020850194506136628686613402565b9095509150846136728284610284565b9350935050509250929050565b600c90565b60006060600083905060608660ff166040519080825280602002602001820160405280156136cc57816020015b6136b9613abf565b8152602001906001900390816136b15790505b50905060005b8760ff168160ff161015613710576136ea8784611fc3565b8351849060ff85169081106136fb57fe5b602090810291909101015292506001016136d2565b5090969095509350505050565b613725613abf565b61372f8251613a5a565b613780576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156137b75783818151811061379a57fe5b602002602001015160800151820191508080600101915050613785565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190613811565b6137fe613abf565b8152602001906001900390816137f65790505b50905281526020810194909452600360408501526060909301525090565b613837613abf565b6040805160028082526060828101909352816020015b613855613abf565b81526020019060019003908161384d579050509050828160008151811061387857fe5b6020026020010181905250838160018151811061389157fe5b6020026020010181905250611a7a816129ae565b6060015160ff1660031490565b606081015160009060ff16600314156138d1575060408101515161027f565b50600161027f565b606081015160009060ff166138f05750600161027f565b606082015160ff16600114156139085750600061027f565b606082015160ff166002141561395c576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b606082015160ff16600314156139bf5760408201515160005b818110156139b45761399d8460400151828151811061399057fe5b60200260200101516138d9565b6139ac5760009250505061027f565b600101613975565b50600191505061027f565b606082015160ff16606414156115665750600061027f565b6139df613abf565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190613a26565b613a13613abf565b815260200190600190039081613a0b5790505b50815260016020820181905260409091015292915050565b60008160200183511015613a5157600080fd5b50016020015190565b6008101590565b6040805161010081019091526000815260208101613a7d613abf565b8152602001613a8a613abf565b8152602001613a97613abf565b8152602001613aa4613abf565b81526000602082018190526040820181905260609091015290565b6040518060a0016040528060008152602001613ad9613af3565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726fa265627a7a723158202078ed62be2930ac56b1cda329f0f705d7721c551eb075c92493baad11f5827264736f6c63430005110032"

// DeployOneStepProofTester deploys a new Ethereum contract, binding an instance of OneStepProofTester to it.
func DeployOneStepProofTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProofTester, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProofTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofTester{OneStepProofTesterCaller: OneStepProofTesterCaller{contract: contract}, OneStepProofTesterTransactor: OneStepProofTesterTransactor{contract: contract}, OneStepProofTesterFilterer: OneStepProofTesterFilterer{contract: contract}}, nil
}

// OneStepProofTester is an auto generated Go binding around an Ethereum contract.
type OneStepProofTester struct {
	OneStepProofTesterCaller     // Read-only binding to the contract
	OneStepProofTesterTransactor // Write-only binding to the contract
	OneStepProofTesterFilterer   // Log filterer for contract events
}

// OneStepProofTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofTesterSession struct {
	Contract     *OneStepProofTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OneStepProofTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofTesterCallerSession struct {
	Contract *OneStepProofTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OneStepProofTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofTesterTransactorSession struct {
	Contract     *OneStepProofTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OneStepProofTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofTesterRaw struct {
	Contract *OneStepProofTester // Generic contract binding to access the raw methods on
}

// OneStepProofTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofTesterCallerRaw struct {
	Contract *OneStepProofTesterCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofTesterTransactorRaw struct {
	Contract *OneStepProofTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofTester creates a new instance of OneStepProofTester, bound to a specific deployed contract.
func NewOneStepProofTester(address common.Address, backend bind.ContractBackend) (*OneStepProofTester, error) {
	contract, err := bindOneStepProofTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTester{OneStepProofTesterCaller: OneStepProofTesterCaller{contract: contract}, OneStepProofTesterTransactor: OneStepProofTesterTransactor{contract: contract}, OneStepProofTesterFilterer: OneStepProofTesterFilterer{contract: contract}}, nil
}

// NewOneStepProofTesterCaller creates a new read-only instance of OneStepProofTester, bound to a specific deployed contract.
func NewOneStepProofTesterCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofTesterCaller, error) {
	contract, err := bindOneStepProofTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTesterCaller{contract: contract}, nil
}

// NewOneStepProofTesterTransactor creates a new write-only instance of OneStepProofTester, bound to a specific deployed contract.
func NewOneStepProofTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofTesterTransactor, error) {
	contract, err := bindOneStepProofTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTesterTransactor{contract: contract}, nil
}

// NewOneStepProofTesterFilterer creates a new log filterer instance of OneStepProofTester, bound to a specific deployed contract.
func NewOneStepProofTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofTesterFilterer, error) {
	contract, err := bindOneStepProofTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTesterFilterer{contract: contract}, nil
}

// bindOneStepProofTester binds a generic wrapper to an already deployed contract.
func bindOneStepProofTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofTester *OneStepProofTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProofTester.Contract.OneStepProofTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofTester *OneStepProofTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofTester.Contract.OneStepProofTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofTester *OneStepProofTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofTester.Contract.OneStepProofTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofTester *OneStepProofTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProofTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofTester *OneStepProofTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofTester *OneStepProofTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofTester.Contract.contract.Transact(opts, method, params...)
}

// ValidateProof is a free data retrieval call binding the contract method 0x3c41485d.
//
// Solidity: function validateProof(bytes32 beforeHash, bytes32 beforeInbox, uint256 beforeInboxValueSize, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(bytes32)
func (_OneStepProofTester *OneStepProofTesterCaller) ValidateProof(opts *bind.CallOpts, beforeHash [32]byte, beforeInbox [32]byte, beforeInboxValueSize *big.Int, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OneStepProofTester.contract.Call(opts, out, "validateProof", beforeHash, beforeInbox, beforeInboxValueSize, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0x3c41485d.
//
// Solidity: function validateProof(bytes32 beforeHash, bytes32 beforeInbox, uint256 beforeInboxValueSize, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(bytes32)
func (_OneStepProofTester *OneStepProofTesterSession) ValidateProof(beforeHash [32]byte, beforeInbox [32]byte, beforeInboxValueSize *big.Int, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) ([32]byte, error) {
	return _OneStepProofTester.Contract.ValidateProof(&_OneStepProofTester.CallOpts, beforeHash, beforeInbox, beforeInboxValueSize, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0x3c41485d.
//
// Solidity: function validateProof(bytes32 beforeHash, bytes32 beforeInbox, uint256 beforeInboxValueSize, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(bytes32)
func (_OneStepProofTester *OneStepProofTesterCallerSession) ValidateProof(beforeHash [32]byte, beforeInbox [32]byte, beforeInboxValueSize *big.Int, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) ([32]byte, error) {
	return _OneStepProofTester.Contract.ValidateProof(&_OneStepProofTester.CallOpts, beforeHash, beforeInbox, beforeInboxValueSize, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820aeb39c6e1b0f43c610e24b542939c7260ded0c5e95ea777aaaf10a3f4227661d64736f6c63430005110032"

// DeployValue deploys a new Ethereum contract, binding an instance of Value to it.
func DeployValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Value, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// Value is an auto generated Go binding around an Ethereum contract.
type Value struct {
	ValueCaller     // Read-only binding to the contract
	ValueTransactor // Write-only binding to the contract
	ValueFilterer   // Log filterer for contract events
}

// ValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueSession struct {
	Contract     *Value            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueCallerSession struct {
	Contract *ValueCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueTransactorSession struct {
	Contract     *ValueTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueRaw struct {
	Contract *Value // Generic contract binding to access the raw methods on
}

// ValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueCallerRaw struct {
	Contract *ValueCaller // Generic read-only contract binding to access the raw methods on
}

// ValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueTransactorRaw struct {
	Contract *ValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValue creates a new instance of Value, bound to a specific deployed contract.
func NewValue(address common.Address, backend bind.ContractBackend) (*Value, error) {
	contract, err := bindValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// NewValueCaller creates a new read-only instance of Value, bound to a specific deployed contract.
func NewValueCaller(address common.Address, caller bind.ContractCaller) (*ValueCaller, error) {
	contract, err := bindValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueCaller{contract: contract}, nil
}

// NewValueTransactor creates a new write-only instance of Value, bound to a specific deployed contract.
func NewValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueTransactor, error) {
	contract, err := bindValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTransactor{contract: contract}, nil
}

// NewValueFilterer creates a new log filterer instance of Value, bound to a specific deployed contract.
func NewValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueFilterer, error) {
	contract, err := bindValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueFilterer{contract: contract}, nil
}

// bindValue binds a generic wrapper to an already deployed contract.
func bindValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.ValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.contract.Transact(opts, method, params...)
}
