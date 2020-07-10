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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820df1cab9f5706f831964c6cfc36e8d53dfb187147670225cce9914aa402e7580164736f6c63430005110032"

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

// MachineABI is the input ABI used to generate the binding from.
const MachineABI = "[]"

// MachineBin is the compiled bytecode used for deploying new contracts.
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201d9917b93a82c10d949c28f2aaa3de85a8f1130cfd8d28c74daf93dfd95a0ea464736f6c63430005110032"

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

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[]"

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b82516f2155f7b046fff1f9b3b19c99d6ab5f8f95c732f5038c7087cea3b7f7664736f6c63430005110032"

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
const OneStepProofTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxValueSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofTesterFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofTesterFuncSigs = map[string]string{
	"f0a516ba": "validateProof(bytes32,bytes32,uint256,bytes32,bool,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
}

// OneStepProofTesterBin is the compiled bytecode used for deploying new contracts.
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b50614039806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f0a516ba14610030575b600080fd5b610123600480360361016081101561004757600080fd5b813591602081013591604082013591606081013591608082013515159160a08101359160c08201359160e0810135916101008201359167ffffffffffffffff610120820135169181019061016081016101408201356401000000008111156100ae57600080fd5b8201836020820111156100c057600080fd5b803590602001918460018302840111640100000000831117156100e257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610135945050505050565b60408051918252519081900360200190f35b600061014a8c8c8c8c8c8c8c8c8c8c8c61015a565b9c9b505050505050505050505050565b600061014a6040518061014001604052808e815260200161017b8e8e6101bf565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff1681526020018481525061024c565b6101c7613e0d565b6040805160a080820183528582528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191610233565b610220613e0d565b8152602001906001900390816102185790505b5081526002602082015260400183905290505b92915050565b6000806000806000606061025e613e41565b610266613e41565b61026f896114bf565b6101008f0151959c50939a509297509095509350915060019060009067ffffffffffffffff1688146102df576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8a6060015180156102f3575060ff89166072145b8061030f57508a6060015115801561030f575060ff8916607214155b610360576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a080840180518a900390528401518811156103875760001960a084015260009150611313565b60ff8916600114156103cd576103c683866000815181106103a457fe5b6020026020010151876001815181106103b957fe5b6020026020010151611819565b9150611313565b60ff89166002141561040c576103c683866000815181106103ea57fe5b6020026020010151876001815181106103ff57fe5b6020026020010151611869565b60ff89166003141561044b576103c6838660008151811061042957fe5b60200260200101518760018151811061043e57fe5b60200260200101516118aa565b60ff89166004141561048a576103c6838660008151811061046857fe5b60200260200101518760018151811061047d57fe5b60200260200101516118eb565b60ff8916600514156104c9576103c683866000815181106104a757fe5b6020026020010151876001815181106104bc57fe5b602002602001015161193c565b60ff891660061415610508576103c683866000815181106104e657fe5b6020026020010151876001815181106104fb57fe5b602002602001015161198d565b60ff891660071415610547576103c6838660008151811061052557fe5b60200260200101518760018151811061053a57fe5b60200260200101516119de565b60ff89166008141561059b576103c6838660008151811061056457fe5b60200260200101518760018151811061057957fe5b60200260200101518860028151811061058e57fe5b6020026020010151611a2f565b60ff8916600914156105ef576103c683866000815181106105b857fe5b6020026020010151876001815181106105cd57fe5b6020026020010151886002815181106105e257fe5b6020026020010151611a99565b60ff8916600a141561062e576103c6838660008151811061060c57fe5b60200260200101518760018151811061062157fe5b6020026020010151611af2565b60ff89166010141561066d576103c6838660008151811061064b57fe5b60200260200101518760018151811061066057fe5b6020026020010151611b33565b60ff8916601114156106ac576103c6838660008151811061068a57fe5b60200260200101518760018151811061069f57fe5b6020026020010151611b74565b60ff8916601214156106eb576103c683866000815181106106c957fe5b6020026020010151876001815181106106de57fe5b6020026020010151611bb5565b60ff89166013141561072a576103c6838660008151811061070857fe5b60200260200101518760018151811061071d57fe5b6020026020010151611bf6565b60ff891660141415610769576103c6838660008151811061074757fe5b60200260200101518760018151811061075c57fe5b6020026020010151611c37565b60ff891660151415610793576103c6838660008151811061078657fe5b6020026020010151611c6e565b60ff8916601614156107d2576103c683866000815181106107b057fe5b6020026020010151876001815181106107c557fe5b6020026020010151611cb3565b60ff891660171415610811576103c683866000815181106107ef57fe5b60200260200101518760018151811061080457fe5b6020026020010151611cf4565b60ff891660181415610850576103c6838660008151811061082e57fe5b60200260200101518760018151811061084357fe5b6020026020010151611d35565b60ff89166019141561087a576103c6838660008151811061086d57fe5b6020026020010151611d76565b60ff8916601a14156108b9576103c6838660008151811061089757fe5b6020026020010151876001815181106108ac57fe5b6020026020010151611dac565b60ff8916601b14156108f8576103c683866000815181106108d657fe5b6020026020010151876001815181106108eb57fe5b6020026020010151611ded565b60ff891660201415610922576103c6838660008151811061091557fe5b6020026020010151611e2e565b60ff89166021141561094c576103c6838660008151811061093f57fe5b6020026020010151611e49565b60ff89166022141561098b576103c6838660008151811061096957fe5b60200260200101518760018151811061097e57fe5b6020026020010151611e64565b60ff8916603014156109b5576103c683866000815181106109a857fe5b6020026020010151611eca565b60ff8916603114156109ca576103c683611ed2565b60ff8916603214156109df576103c683611ef3565b60ff891660331415610a09576103c683866000815181106109fc57fe5b6020026020010151611f0c565b60ff891660341415610a33576103c68386600081518110610a2657fe5b6020026020010151611f18565b60ff891660351415610a72576103c68386600081518110610a5057fe5b602002602001015187600181518110610a6557fe5b6020026020010151611f43565b60ff891660361415610a87576103c683611f8b565b60ff891660371415610aa1576103c6838560000151611fb5565b60ff891660381415610acb576103c68386600081518110610abe57fe5b6020026020010151611fc5565b60ff891660391415610b5757610adf613e0d565b610aee8c610120015188611fd7565b9199509750905087610b315760405162461bcd60e51b8152600401808060200182810382526021815260200180613fe46021913960400191505060405180910390fd5b610b41858263ffffffff61211516565b610b51848263ffffffff61212f16565b50611313565b60ff8916603a1415610b6c576103c683612149565b60ff8916603b1415610b815760019150611313565b60ff8916603c1415610b96576103c683612166565b60ff8916603d1415610bc0576103c68386600081518110610bb357fe5b602002602001015161217a565b60ff891660401415610bea576103c68386600081518110610bdd57fe5b60200260200101516121a8565b60ff891660411415610c29576103c68386600081518110610c0757fe5b602002602001015187600181518110610c1c57fe5b60200260200101516121ca565b60ff891660421415610c7d576103c68386600081518110610c4657fe5b602002602001015187600181518110610c5b57fe5b602002602001015188600281518110610c7057fe5b60200260200101516121fc565b60ff891660431415610cbc576103c68386600081518110610c9a57fe5b602002602001015187600181518110610caf57fe5b602002602001015161223e565b60ff891660441415610d10576103c68386600081518110610cd957fe5b602002602001015187600181518110610cee57fe5b602002602001015188600281518110610d0357fe5b6020026020010151612250565b60ff891660501415610d4f576103c68386600081518110610d2d57fe5b602002602001015187600181518110610d4257fe5b6020026020010151612272565b60ff891660511415610da3576103c68386600081518110610d6c57fe5b602002602001015187600181518110610d8157fe5b602002602001015188600281518110610d9657fe5b60200260200101516122e8565b60ff891660521415610dcd576103c68386600081518110610dc057fe5b6020026020010151612375565b60ff891660531415610e6a57610de1613e0d565b610df08c610120015188611fd7565b9199509750905087610e335760405162461bcd60e51b8152600401808060200182810382526021815260200180613fe46021913960400191505060405180910390fd5b610e43858263ffffffff61211516565b610e628487600081518110610e5457fe5b6020026020010151836123a8565b925050611313565b60ff891660541415610f1457610e7e613e0d565b610e8d8c610120015188611fd7565b9199509750905087610ed05760405162461bcd60e51b8152600401808060200182810382526021815260200180613fe46021913960400191505060405180910390fd5b610ee0858263ffffffff61211516565b610e628487600081518110610ef157fe5b602002602001015188600181518110610f0657fe5b602002602001015184612400565b60ff891660601415610f29576103c683612481565b60ff89166061141561102657610f538386600081518110610f4657fe5b6020026020010151612487565b9092509050811561101d578a60e001518b60c001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610fd25760405162461bcd60e51b8152600401808060200182810382526025815260200180613f986025913960400191505060405180910390fd5b8a60a001518b60800151146110185760405162461bcd60e51b8152600401808060200182810382526027815260200180613fbd6027913960400191505060405180910390fd5b611021565b5060005b611313565b60ff89166070141561116557611050838660008151811061104357fe5b60200260200101516124a1565b9092509050811561101d57806110ab578a60a001518b60800151146110a65760405162461bcd60e51b8152600401808060200182810382526038815260200180613f606038913960400191505060405180910390fd5b611018565b8a60a001518b60800151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001201461111f5760405162461bcd60e51b8152600401808060200182810382526029815260200180613ef06029913960400191505060405180910390fd5b8a60e001518b60c00151146110185760405162461bcd60e51b8152600401808060200182810382526026815260200180613f196026913960400191505060405180910390fd5b60ff89166072141561117f576103c6838c602001516124e0565b60ff8916607314156111945760009150611313565b60ff8916607414156111a95761102183612546565b60ff8916607514156111d3576103c683866000815181106111c657fe5b6020026020010151612550565b60ff8916607614156111e8576103c683612575565b60ff8916607714156111fd576103c68361258e565b60ff89166078141561123c576103c6838660008151811061121a57fe5b60200260200101518760018151811061122f57fe5b60200260200101516125d7565b60ff891660791415611290576103c6838660008151811061125957fe5b60200260200101518760018151811061126e57fe5b60200260200101518860028151811061128357fe5b602002602001015161261c565b60ff8916607b14156112a5576103c68361266f565b60ff89166080141561130e576103c683866000815181106112c257fe5b6020026020010151876001815181106112d757fe5b6020026020010151886002815181106112ec57fe5b60200260200101518960038151811061130157fe5b60200260200101516126b2565b600091505b806113a4578a60a001518b608001511461135e5760405162461bcd60e51b8152600401808060200182810382526027815260200180613fbd6027913960400191505060405180910390fd5b8a60e001518b60c00151146113a45760405162461bcd60e51b8152600401808060200182810382526026815260200180613f196026913960400191505060405180910390fd5b816114055760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156113fd576113f8836127d1565b611405565b60c083015183525b61140e846127db565b8b511461144c5760405162461bcd60e51b8152600401808060200182810382526022815260200180613ece6022913960400191505060405180910390fd5b611455836127db565b8b60400151146114ac576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b600099505050505050505050505b919050565b60008060606114cc613e41565b6114d4613e41565b6000806114e08461289f565b6114ef886101200151836128a9565b95509250905080611547576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b611550846129e1565b92506000886101200151838151811061156557fe5b602001015160f81c60f81b60f81c9050886101200151836001018151811061158957fe5b016020015160f81c9750600061159e89612a4a565b6040805183815260208085028201019091529099509091508180156115dd57816020015b6115ca613e0d565b8152602001906001900390816115c25790505b5096506002840193508160ff16600014806115fb57508160ff166001145b61164c576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff82166116715761166a6116658a8860000151612fa7565b612fe8565b865261172c565b611679613e0d565b6116888b610120015186611fd7565b9096509094509050836116e2576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b81156117065780886000815181106116f657fe5b6020026020010181905250611716565b611716868263ffffffff61212f16565b6117286116658b8960000151846130e0565b8752505b60ff82165b818110156117bf576117488b610120015186611fd7565b8a518b908590811061175657fe5b602090810291909101015295509350836117b7576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611731565b87511561180c575060005b8260ff1688510381101561180c57611804888260018b510303815181106117ed57fe5b60200260200101518861212f90919063ffffffff16565b6001016117ca565b5050505091939550919395565b600061182483613129565b1580611836575061183482613129565b155b1561184357506000611862565b8251825180820161185a878263ffffffff61313416565b600193505050505b9392505050565b600061187483613129565b1580611886575061188482613129565b155b1561189357506000611862565b8251825180820261185a878263ffffffff61313416565b60006118b583613129565b15806118c757506118c582613129565b155b156118d457506000611862565b8251825180820361185a878263ffffffff61313416565b60006118f683613129565b1580611908575061190682613129565b155b1561191557506000611862565b825182518061192957600092505050611862565b80820461185a878263ffffffff61313416565b600061194783613129565b1580611959575061195782613129565b155b1561196657506000611862565b825182518061197a57600092505050611862565b80820561185a878263ffffffff61313416565b600061199883613129565b15806119aa57506119a882613129565b155b156119b757506000611862565b82518251806119cb57600092505050611862565b80820661185a878263ffffffff61313416565b60006119e983613129565b15806119fb57506119f982613129565b155b15611a0857506000611862565b8251825180611a1c57600092505050611862565b80820761185a878263ffffffff61313416565b6000611a3a84613129565b1580611a4c5750611a4a83613129565b155b15611a5957506000611a91565b83518351835180611a705760009350505050611a91565b6000818385089050611a88898263ffffffff61313416565b60019450505050505b949350505050565b6000611aa484613129565b1580611ab65750611ab483613129565b155b15611ac357506000611a91565b83518351835180611ada5760009350505050611a91565b6000818385099050611a88898263ffffffff61313416565b6000611afd83613129565b1580611b0f5750611b0d82613129565b155b15611b1c57506000611862565b8251825180820a61185a878263ffffffff61313416565b6000611b3e83613129565b1580611b505750611b4e82613129565b155b15611b5d57506000611862565b8251825180821061185a878263ffffffff61313416565b6000611b7f83613129565b1580611b915750611b8f82613129565b155b15611b9e57506000611862565b8251825180821161185a878263ffffffff61313416565b6000611bc083613129565b1580611bd25750611bd082613129565b155b15611bdf57506000611862565b8251825180821261185a878263ffffffff61313416565b6000611c0183613129565b1580611c135750611c1182613129565b155b15611c2057506000611862565b8251825180821361185a878263ffffffff61313416565b6000611c64611c57611c4884612fe8565b611c5186612fe8565b1461314a565b859063ffffffff61212f16565b5060019392505050565b6000611c7982613129565b611c9357611c8e83600063ffffffff61313416565b611caa565b81518015611ca7858263ffffffff61313416565b50505b50600192915050565b6000611cbe83613129565b1580611cd05750611cce82613129565b155b15611cdd57506000611862565b8251825180821661185a878263ffffffff61313416565b6000611cff83613129565b1580611d115750611d0f82613129565b155b15611d1e57506000611862565b8251825180821761185a878263ffffffff61313416565b6000611d4083613129565b1580611d525750611d5082613129565b155b15611d5f57506000611862565b8251825180821861185a878263ffffffff61313416565b6000611d8182613129565b611d8d57506000610246565b81518019611da1858263ffffffff61313416565b506001949350505050565b6000611db783613129565b1580611dc95750611dc782613129565b155b15611dd657506000611862565b8251825181811a61185a878263ffffffff61313416565b6000611df883613129565b1580611e0a5750611e0882613129565b155b15611e1757506000611862565b8251825181810b61185a878263ffffffff61313416565b6000611caa611e3c83612fe8565b849063ffffffff61313416565b6000611caa611e578361316c565b849063ffffffff61212f16565b6000611e6f83613129565b1580611e815750611e7f82613129565b155b15611e8e57506000611862565b825182516040805160208082018590528183018490528251808303840181526060909201909252805191012061185a878263ffffffff61313416565b600192915050565b6000611eeb82608001518361212f90919063ffffffff16565b506001919050565b6000611eeb82606001518361212f90919063ffffffff16565b60609190910152600190565b6000611f23826131f5565b611f2f57506000610246565b611f3882612fe8565b835250600192915050565b6000611f4e836131f5565b611f5a57506000611862565b611f6382613129565b611f6f57506000611862565b815115611c6457611f7f83612fe8565b84525060019392505050565b6000611eeb611fa8611f9b613202565b611c518560200151612fe8565b839063ffffffff61212f16565b6000611caa611e57836001613223565b6000611caa838363ffffffff61211516565b600080611fe2613e0d565b8451841061200257600084611ff760006132ae565b92509250925061210e565b600080859050600087828151811061201657fe5b016020015160019092019160f81c90506000612030613e9f565b60ff8316612064576120428a8561333a565b919650945091508484612054846132ae565b975097509750505050505061210e565b60ff83166001141561208c5761207a8a8561338d565b9196509450905084846120548361350b565b60ff8316600214156120a2576120548a85613572565b600360ff8416108015906120b95750600c60ff8416105b156120f457600219830160606120d0828d88613617565b9198509650905086866120e2836136d5565b9950995099505050505050505061210e565b60008061210160006132ae565b9199509750955050505050505b9250925092565b6121238260400151826137cc565b82604001819052505050565b61213d8260200151826137cc565b82602001819052505050565b6000611eeb611fa8612159613202565b611c518560400151612fe8565b6000611eeb611fa88360c001516001613223565b6000612185826131f5565b61219157506000610246565b61219a82612fe8565b60c084015250600192915050565b60006121ba838363ffffffff61212f16565b611caa838363ffffffff61212f16565b60006121dc848363ffffffff61212f16565b6121ec848463ffffffff61212f16565b611c64848363ffffffff61212f16565b600061220e858363ffffffff61212f16565b61221e858463ffffffff61212f16565b61222e858563ffffffff61212f16565b611da1858363ffffffff61212f16565b60006121ec848463ffffffff61212f16565b6000612262858563ffffffff61212f16565b61222e858463ffffffff61212f16565b600061227d83613129565b158061228f575061228d8261384a565b155b1561229c57506000611862565b6122a582613859565b60ff168360000151106122ba57506000611862565b611c6482604001518460000151815181106122d157fe5b60200260200101518561212f90919063ffffffff16565b60006122f38361384a565b1580612305575061230384613129565b155b1561231257506000611a91565b61231b83613859565b60ff1684600001511061233057506000611a91565b60408301518451815184918391811061234557fe5b602002602001018190525061236961235c826136d5565b879063ffffffff61212f16565b50600195945050505050565b60006123808261384a565b61238c57506000610246565b611caa61239883613859565b849060ff1663ffffffff61313416565b60006123b383613129565b15806123c557506123c38261384a565b155b156123d257506000611862565b6123db82613859565b60ff168360000151106123f057506000611862565b6122ba848363ffffffff61211516565b600061240b8261384a565b158061241d575061241b84613129565b155b1561242a57506000611a91565b61243382613859565b60ff1684600001511061244857506000611a91565b60408201518451815185918391811061245d57fe5b6020026020010181905250612369612474826136d5565b879063ffffffff61211516565b50600190565b600080600161249584612fe8565b915091505b9250929050565b600080612710836080015111156124bd5750600090508061249a565b6124c683613868565b6124d55750600090508061249a565b600161249584612fe8565b60006124ea613202565b6124f383612fe8565b14156121ba576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b600260e090910152565b600061255b82613129565b61256757506000610246565b505160a09190910152600190565b6000611eeb8260a001518361313490919063ffffffff16565b60408051600160f81b60208083019190915260006021830181905260228084018290528451808503909101815260429093019093528151910120611eeb90611fa8906001613223565b60006125e283613129565b6125ee57506000611862565b6125f7826131f5565b61260357506000611862565b611c64611c57846000015161261785612fe8565b612fa7565b600061262784613129565b61263357506000611a91565b61263c826131f5565b61264857506000611a91565b611da1612662856000015161265c85612fe8565b866130e0565b869063ffffffff61212f16565b60408051600080825260208201909252606090826126a3565b612690613e0d565b8152602001906001900390816126885790505b509050611caa611e57826136d5565b60006126bd85613129565b15806126cf57506126cd84613129565b155b806126e057506126de83613129565b155b806126f157506126ef82613129565b155b156126fe575060006127c8565b8451845184511580159061271457508451600114155b156127355761272a88600063ffffffff61313416565b6001925050506127c8565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612797573d6000803e3d6000fd5b5050604051601f19015191506127be90508b6001600160a01b03831663ffffffff61313416565b6001955050505050505b95945050505050565b600160e090910152565b600060028260e0015114156127f2575060006114ba565b60018260e001511415612807575060016114ba565b8151602083015161281790612fe8565b6128248460400151612fe8565b6128318560600151612fe8565b61283e8660800151612fe8565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090506114ba565b600060e090910152565b6000806128b4613e41565b6128bc613e41565b600060e082018190526128cf878761397f565b845296509050806128e9575060009350849250905061210e565b6128f38787613572565b60208501529650905080612910575060009350849250905061210e565b61291a8787613572565b60408501529650905080612937575060009350849250905061210e565b6129418787611fd7565b6060850152965090508061295e575060009350849250905061210e565b6129688787611fd7565b60808501529650905080612985575060009350849250905061210e565b61298f878761333a565b60a085015296509050806129ac575060009350849250905061210e565b6129b6878761397f565b60c085015296509050806129d3575060009350849250905061210e565b506001969495509392505050565b6129e9613e41565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612a625750600290506003612fa2565b6002831415612a775750600290506003612fa2565b6003831415612a8c5750600290506003612fa2565b6004831415612aa15750600290506004612fa2565b6005831415612ab65750600290506007612fa2565b6006831415612acb5750600290506004612fa2565b6007831415612ae05750600290506007612fa2565b6008831415612af55750600390506004612fa2565b6009831415612b0a5750600390506004612fa2565b600a831415612b1f5750600290506019612fa2565b6010831415612b3357506002905080612fa2565b6011831415612b4757506002905080612fa2565b6012831415612b5b57506002905080612fa2565b6013831415612b6f57506002905080612fa2565b6014831415612b8357506002905080612fa2565b6015831415612b9757506001905080612fa2565b6016831415612bab57506002905080612fa2565b6017831415612bbf57506002905080612fa2565b6018831415612bd357506002905080612fa2565b6019831415612be757506001905080612fa2565b601a831415612bfc5750600290506004612fa2565b601b831415612c115750600290506007612fa2565b6020831415612c265750600190506007612fa2565b6021831415612c3b5750600190506003612fa2565b6022831415612c505750600290506008612fa2565b6030831415612c6457506001905080612fa2565b6031831415612c795750600090506001612fa2565b6032831415612c8e5750600090506001612fa2565b6033831415612ca35750600190506002612fa2565b6034831415612cb85750600190506004612fa2565b6035831415612ccd5750600290506004612fa2565b6036831415612ce25750600090506002612fa2565b6037831415612cf75750600090506001612fa2565b6038831415612d0b57506001905080612fa2565b6039831415612d205750600090506001612fa2565b603a831415612d355750600090506002612fa2565b603b831415612d4a5750600090506001612fa2565b603c831415612d5f5750600090506001612fa2565b603d831415612d7357506001905080612fa2565b6040831415612d8757506001905080612fa2565b6041831415612d9c5750600290506001612fa2565b6042831415612db15750600390506001612fa2565b6043831415612dc65750600290506001612fa2565b6044831415612ddb5750600390506001612fa2565b6050831415612def57506002905080612fa2565b6051831415612e045750600390506028612fa2565b6052831415612e195750600190506002612fa2565b6053831415612e2e5750600190506003612fa2565b6054831415612e435750600290506029612fa2565b6060831415612e585750600090506064612fa2565b6061831415612e6d5750600190506064612fa2565b6070831415612e825750600190506064612fa2565b6072831415612e975750600090506028612fa2565b6073831415612eac5750600090506005612fa2565b6074831415612ec1575060009050600a612fa2565b6075831415612ed65750600190506000612fa2565b6076831415612eeb5750600090506001612fa2565b6077831415612f005750600090506019612fa2565b6078831415612f155750600290506019612fa2565b6079831415612f2a5750600390506019612fa2565b607b831415612f3f575060009050600a612fa2565b6080831415612f55575060049050614e20612fa2565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b612faf613e0d565b6118626040518060a001604052808560ff1681526020018481526020016000151581526020016000801b8152602001600081525061350b565b606081015160009060ff16613009578151613002906139d3565b90506114ba565b606082015160ff166001141561303c576020808301518051604082015160608301519290930151613002939192906139f7565b606082015160ff16600214156130555761300282613a9f565b600360ff16826060015160ff161015801561307957506060820151600c60ff909116105b156130875761300282613b05565b606082015160ff166064141561309f575080516114ba565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6130e8613e0d565b611a916040518060a001604052808660ff16815260200185815260200160011515815260200161311785612fe8565b8152602001846080015181525061350b565b6060015160ff161590565b61213d8260200151613145836132ae565b6137cc565b613152613e0d565b81156131625761300260016132ae565b61300260006132ae565b613174613e0d565b816060015160ff16600214156131bb5760405162461bcd60e51b8152600401808060200182810382526021815260200180613f3f6021913960400191505060405180910390fd5b606082015160ff166131d15761300260006132ae565b816060015160ff16600114156131eb5761300260016132ae565b61300260036132ae565b6060015160ff1660011490565b6040805160008082526020820190925261321d816001613b23565b91505090565b61322b613e0d565b6040805160a080820183528582528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191613297565b613284613e0d565b81526020019060019003908161327c5790505b508152606460208201526040019290925250919050565b6132b6613e0d565b6040805160a080820183528482528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191613322565b61330f613e0d565b8152602001906001900390816133075790505b50815260006020820152600160409091015292915050565b600080600080855190508481108061335457506020858203105b1561336957506000925083915082905061210e565b60016020860161337f888863ffffffff613b4216565b935093509350509250925092565b600080613398613e9f565b600084905060008682815181106133ab57fe5b602001015160f81c60f81b60f81c9050818060010192505060008783815181106133d157fe5b016020015160019384019360f89190911c9150600090819060ff851614156134685760006133fd613e0d565b6134078c88611fd7565b9098509092509050816134535750506040805160a0810182526000808252602082018190529181018290526060810182905260808101829052909850899750955061210e945050505050565b61345c81612fe8565b93508060800151925050505b600061347a8b8763ffffffff613b4216565b90506020860195508460ff16600114156134cb576040805160a08101825260ff90951685526020850191909152600190840181905260608401929092526080830152955091935090915061210e9050565b6040805160a08101825260ff9095168552602085019190915260009084018190526060840181905260808401525060019650929450925050509250925092565b613513613e0d565b6040805160a081018252600080825260208083018690528351828152908101845291928301919061355a565b613547613e0d565b81526020019060019003908161353f5790505b50815260016020820181905260409091015292915050565b60008061357d613e0d565b613585613e0d565b855160009081908781108061359c57506040888203105b156135b457600088859650965096505050505061210e565b60006135c68a8a63ffffffff613b4216565b90506020890198506135d88a8a61333a565b909a50945092508215613603576135ef81856101bf565b60019850899750955061210e945050505050565b60008986975097509750505050505061210e565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561366257816020015b61364f613e0d565b8152602001906001900390816136475790505b50905060005b8960ff168160ff1610156136bf576136808985611fd7565b8451859060ff861690811061369157fe5b602090810291909101015294509250826136b757506000955086945092506136cc915050565b600101613668565b5060019550919350909150505b93509350939050565b6136dd613e0d565b6136e78251613b5e565b613738576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b835181101561376f5783818151811061375257fe5b60200260200101516080015182019150808060010191505061373d565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b6137d4613e0d565b6040805160028082526060828101909352816020015b6137f2613e0d565b8152602001906001900390816137ea579050509050828160008151811061381557fe5b6020026020010181905250838160018151811061382e57fe5b6020026020010181905250611a91613845826136d5565b613b65565b60006102468260600151613bdb565b60006102468260600151613bf9565b606081015160009060ff1661387f575060016114ba565b606082015160ff1660011415613897575060006114ba565b606082015160ff16600214156138eb576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff161015801561390f57506060820151600c60ff909116105b156139675760408201515160005b8181101561395c576139458460400151828151811061393857fe5b6020026020010151613868565b613954576000925050506114ba565b60010161391d565b5060019150506114ba565b606082015160ff166064141561309f575060006114ba565b6000806000806000865190508581108061399b57506020868203105b156139af575060009350849250905061210e565b6139bf878763ffffffff613b4216565b60019550602087019450925061210e915050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613a51575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611a91565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214613af4576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b815160808301516102469190613c1c565b6000613b0f613e0d565b613b1883613b65565b905061186281613a9f565b6000613b2d613e0d565b613b378484613c56565b9050611a9181613a9f565b60008160200183511015613b5557600080fd5b50016020015190565b6008101590565b613b6d613e0d565b613b768261384a565b613bbc576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060613bcb8360400151613c75565b9050611862818460800151613c56565b6000600c60ff8316108015610246575050600360ff91909116101590565b6000613c0482613bdb565b15613c14575060021981016114ba565b5060016114ba565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b613c5e613e0d565b6000613c6984613d4d565b9050611a9181846101bf565b6060600882511115613cc5576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613cf2578160200160208202803883390190505b50805190915060005b81811015613d44576000613d21868381518110613d1457fe5b6020026020010151612fe8565b905080848381518110613d3057fe5b602090810291909101015250600101613cfb565b50909392505050565b6000600882511115613d9d576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613de1578181015183820152602001613dc9565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001613e27613e9f565b815260606020820181905260006040830181905291015290565b6040805161010081019091526000815260208101613e5d613e0d565b8152602001613e6a613e0d565b8152602001613e77613e0d565b8152602001613e84613e0d565b81526000602082018190526040820181905260609091015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a723158202ad3e28b9cc624c4dc34d8401028bcfbc9374b27ae85c36eb067b530e53b145464736f6c63430005110032"

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

// ValidateProof is a free data retrieval call binding the contract method 0xf0a516ba.
//
// Solidity: function validateProof(bytes32 beforeHash, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProofTester *OneStepProofTesterCaller) ValidateProof(opts *bind.CallOpts, beforeHash [32]byte, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProofTester.contract.Call(opts, out, "validateProof", beforeHash, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0xf0a516ba.
//
// Solidity: function validateProof(bytes32 beforeHash, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProofTester *OneStepProofTesterSession) ValidateProof(beforeHash [32]byte, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProofTester.Contract.ValidateProof(&_OneStepProofTester.CallOpts, beforeHash, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0xf0a516ba.
//
// Solidity: function validateProof(bytes32 beforeHash, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProofTester *OneStepProofTesterCallerSession) ValidateProof(beforeHash [32]byte, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProofTester.Contract.ValidateProof(&_OneStepProofTester.CallOpts, beforeHash, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820ef656722934892a7fe491023d83e2264b4f0f5cb994ff208179614208b4740e564736f6c63430005110032"

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
