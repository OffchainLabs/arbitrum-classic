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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582030fe4bc4c02f5b857cc758b4af05238ac08d82c4209ae6d8e043a8a7ae26320c64736f6c634300050f0032"

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
var DebugPrintBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820e1b40ecc491bd6dc1565bd13fc83a5596d1048e47721f29a1f46ed30d588e35864736f6c634300050f0032"

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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582028ea2e2d8cea1a66da9699bb1f6b44ea51f3461266ce685e69f33d91f95d42a664736f6c634300050f0032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158204311e9cfec405cff4290a79e7b1ac99b9a2dbeaf75d5af585d66fa8841718da464736f6c634300050f0032"

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
const OneStepProofTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128[4]\",\"name\":\"timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxValueSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofTesterFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofTesterFuncSigs = map[string]string{
	"e987d887": "validateProof(bytes32,uint128[4],bytes32,uint256,bytes32,bool,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
}

// OneStepProofTesterBin is the compiled bytecode used for deploying new contracts.
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b506140eb806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e987d88714610030575b600080fd5b61015260048036036101e081101561004757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100dd57600080fd5b8201836020820111156100ef57600080fd5b8035906020019184600183028401116401000000008311171561011157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610164945050505050565b60408051918252519081900360200190f35b600061017a8d8d8d8d8d8d8d8d8d8d8d8d61018b565b9d9c50505050505050505050505050565b600061017a6040518061016001604052808f81526020018e81526020016101b28e8e6101f6565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff1681526020018481525061027c565b6101fe613ece565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610263565b610250613ece565b8152602001906001900390816102485790505b5081526002602082015260400183905290505b92915050565b600080600080606061028c613f02565b610294613f02565b61029d886113db565b93995092965090945092509050600160006102b78861175a565b67ffffffffffffffff168a610120015167ffffffffffffffff161461031a576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8960800151801561032e575060ff88166072145b8061034a5750896080015115801561034a575060ff8816607214155b61039b576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60ff8816600114156103e1576103da83866000815181106103b857fe5b6020026020010151876001815181106103cd57fe5b6020026020010151611bf6565b915061122f565b60ff881660021415610420576103da83866000815181106103fe57fe5b60200260200101518760018151811061041357fe5b6020026020010151611c46565b60ff88166003141561045f576103da838660008151811061043d57fe5b60200260200101518760018151811061045257fe5b6020026020010151611c87565b60ff88166004141561049e576103da838660008151811061047c57fe5b60200260200101518760018151811061049157fe5b6020026020010151611cc8565b60ff8816600514156104dd576103da83866000815181106104bb57fe5b6020026020010151876001815181106104d057fe5b6020026020010151611d19565b60ff88166006141561051c576103da83866000815181106104fa57fe5b60200260200101518760018151811061050f57fe5b6020026020010151611d6a565b60ff88166007141561055b576103da838660008151811061053957fe5b60200260200101518760018151811061054e57fe5b6020026020010151611dbb565b60ff8816600814156105af576103da838660008151811061057857fe5b60200260200101518760018151811061058d57fe5b6020026020010151886002815181106105a257fe5b6020026020010151611e0c565b60ff881660091415610603576103da83866000815181106105cc57fe5b6020026020010151876001815181106105e157fe5b6020026020010151886002815181106105f657fe5b6020026020010151611e76565b60ff8816600a1415610642576103da838660008151811061062057fe5b60200260200101518760018151811061063557fe5b6020026020010151611ecf565b60ff881660101415610681576103da838660008151811061065f57fe5b60200260200101518760018151811061067457fe5b6020026020010151611f10565b60ff8816601114156106c0576103da838660008151811061069e57fe5b6020026020010151876001815181106106b357fe5b6020026020010151611f51565b60ff8816601214156106ff576103da83866000815181106106dd57fe5b6020026020010151876001815181106106f257fe5b6020026020010151611f92565b60ff88166013141561073e576103da838660008151811061071c57fe5b60200260200101518760018151811061073157fe5b6020026020010151611fd3565b60ff88166014141561077d576103da838660008151811061075b57fe5b60200260200101518760018151811061077057fe5b6020026020010151612014565b60ff8816601514156107a7576103da838660008151811061079a57fe5b602002602001015161203e565b60ff8816601614156107e6576103da83866000815181106107c457fe5b6020026020010151876001815181106107d957fe5b6020026020010151612083565b60ff881660171415610825576103da838660008151811061080357fe5b60200260200101518760018151811061081857fe5b60200260200101516120c4565b60ff881660181415610864576103da838660008151811061084257fe5b60200260200101518760018151811061085757fe5b6020026020010151612105565b60ff88166019141561088e576103da838660008151811061088157fe5b6020026020010151612146565b60ff8816601a14156108cd576103da83866000815181106108ab57fe5b6020026020010151876001815181106108c057fe5b602002602001015161217c565b60ff8816601b141561090c576103da83866000815181106108ea57fe5b6020026020010151876001815181106108ff57fe5b60200260200101516121bd565b60ff881660201415610936576103da838660008151811061092957fe5b60200260200101516121fe565b60ff881660211415610960576103da838660008151811061095357fe5b6020026020010151612219565b60ff88166022141561099f576103da838660008151811061097d57fe5b60200260200101518760018151811061099257fe5b6020026020010151612234565b60ff8816603014156109c9576103da83866000815181106109bc57fe5b602002602001015161229a565b60ff8816603114156109de576103da836122a2565b60ff8816603214156109f3576103da836122c3565b60ff881660331415610a1d576103da8386600081518110610a1057fe5b60200260200101516122dc565b60ff881660341415610a47576103da8386600081518110610a3a57fe5b60200260200101516122e8565b60ff881660351415610a86576103da8386600081518110610a6457fe5b602002602001015187600181518110610a7957fe5b6020026020010151612313565b60ff881660361415610a9b576103da8361235b565b60ff881660371415610ab5576103da838560000151612385565b60ff881660381415610adf576103da8386600081518110610ad257fe5b6020026020010151612393565b60ff881660391415610b6b57610af3613ece565b610b028b6101400151886123a5565b9199509750905087610b455760405162461bcd60e51b81526004018080602001828103825260218152602001806140966021913960400191505060405180910390fd5b610b55858263ffffffff6124e316565b610b65848263ffffffff6124fd16565b5061122f565b60ff8816603a1415610b80576103da83612517565b60ff8816603b1415610b95576001915061122f565b60ff8816603c1415610baa576103da83612534565b60ff8816603d1415610bd4576103da8386600081518110610bc757fe5b6020026020010151612546565b60ff881660401415610bfe576103da8386600081518110610bf157fe5b6020026020010151612574565b60ff881660411415610c3d576103da8386600081518110610c1b57fe5b602002602001015187600181518110610c3057fe5b6020026020010151612596565b60ff881660421415610c91576103da8386600081518110610c5a57fe5b602002602001015187600181518110610c6f57fe5b602002602001015188600281518110610c8457fe5b60200260200101516125c8565b60ff881660431415610cd0576103da8386600081518110610cae57fe5b602002602001015187600181518110610cc357fe5b602002602001015161260a565b60ff881660441415610d24576103da8386600081518110610ced57fe5b602002602001015187600181518110610d0257fe5b602002602001015188600281518110610d1757fe5b602002602001015161261c565b60ff881660501415610d63576103da8386600081518110610d4157fe5b602002602001015187600181518110610d5657fe5b602002602001015161263e565b60ff881660511415610db7576103da8386600081518110610d8057fe5b602002602001015187600181518110610d9557fe5b602002602001015188600281518110610daa57fe5b60200260200101516126b4565b60ff881660521415610de1576103da8386600081518110610dd457fe5b6020026020010151612741565b60ff881660601415610df6576103da83612774565b60ff881660611415610ef457610e208386600081518110610e1357fe5b602002602001015161277a565b90925090508115610eeb578961010001518a60e001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610ea05760405162461bcd60e51b815260040180806020018281038252602581526020018061404a6025913960400191505060405180910390fd5b8960c001518a60a0015114610ee65760405162461bcd60e51b815260040180806020018281038252602781526020018061406f6027913960400191505060405180910390fd5b610eef565b5060005b61122f565b60ff88166070141561103457610f1e8386600081518110610f1157fe5b6020026020010151612794565b90925090508115610eeb5780610f79578960c001518a60a0015114610f745760405162461bcd60e51b81526004018080602001828103825260388152602001806140126038913960400191505060405180910390fd5b610ee6565b8960c001518a60a001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610fed5760405162461bcd60e51b8152600401808060200182810382526029815260200180613fa26029913960400191505060405180910390fd5b8961010001518a60e0015114610ee65760405162461bcd60e51b8152600401808060200182810382526026815260200180613fcb6026913960400191505060405180910390fd5b60ff8816607114156111495760408051600480825260a08201909252606091816020015b611060613ece565b81526020019060019003908161105857505060208c01519091506110949060005b60200201516001600160801b03166127d3565b816000815181106110a157fe5b60200260200101819052506110c08b6020015160016004811061108157fe5b816001815181106110cd57fe5b60200260200101819052506110ec8b6020015160026004811061108157fe5b816002815181106110f957fe5b60200260200101819052506111188b6020015160036004811061108157fe5b8160038151811061112557fe5b6020026020010181905250610b6561113c82612858565b859063ffffffff6124fd16565b60ff881660721415611197576103da838660008151811061116657fe5b60200260200101518c604001518d6020015160006004811061118457fe5b60200201516001600160801b0316612947565b60ff8816607314156111ac576000915061122f565b60ff8816607414156111c157610eef836129de565b60ff88166080141561122a576103da83866000815181106111de57fe5b6020026020010151876001815181106111f357fe5b60200260200101518860028151811061120857fe5b60200260200101518960038151811061121d57fe5b60200260200101516129e8565b600091505b806112c1578960c001518a60a001511461127a5760405162461bcd60e51b815260040180806020018281038252602781526020018061406f6027913960400191505060405180910390fd5b8961010001518a60e00151146112c15760405162461bcd60e51b8152600401808060200182810382526026815260200180613fcb6026913960400191505060405180910390fd5b816113225760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a0840151141561131a5761131583612b07565b611322565b60a083015183525b61132b84612b11565b8a51146113695760405162461bcd60e51b8152600401808060200182810382526022815260200180613f806022913960400191505060405180910390fd5b61137283612b11565b8a60600151146113c9576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b600060606113e7613f02565b6113ef613f02565b600080806113fb613f02565b61140481612bc9565b61141389610140015184612bd3565b90945090925090508161146d576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b611475613f02565b61147e82612ce4565b905060008a6101400151858151811061149357fe5b602001015160f81c60f81b60f81c905060008b610140015186600101815181106114b957fe5b016020015160f81c905060006114ce82612d42565b905060608160405190808252806020026020018201604052801561150c57816020015b6114f9613ece565b8152602001906001900390816114f15790505b5090506002880197508360ff166000148061152a57508360ff166001145b61157b576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff84166115a057611599611594848860000151612d5c565b612d96565b8652611663565b6115a8613ece565b6115b78f61014001518a6123a5565b909a50909850905087611611576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561163557808260008151811061162557fe5b6020026020010181905250611645565b611645868263ffffffff6124fd16565b61165f61159485896000015161165a85612d96565b612e8e565b8752505b60ff84165b828110156116f65761167f8f61014001518a6123a5565b845185908590811061168d57fe5b602090810291909101015299509750876116ee576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611668565b815115611743575060005b8460ff168251038110156117435761173b82826001855103038151811061172457fe5b6020026020010151886124fd90919063ffffffff16565b600101611701565b50919d919c50939a50919850939650945050505050565b600060ff821660011415611770575060036113d6565b60ff821660021415611784575060036113d6565b60ff821660031415611798575060036113d6565b60ff8216600414156117ac575060046113d6565b60ff8216600514156117c0575060076113d6565b60ff8216600614156117d4575060046113d6565b60ff8216600714156117e8575060076113d6565b60ff8216600814156117fc575060046113d6565b60ff821660091415611810575060046113d6565b60ff8216600a1415611824575060196113d6565b60ff821660101415611838575060026113d6565b60ff82166011141561184c575060026113d6565b60ff821660121415611860575060026113d6565b60ff821660131415611874575060026113d6565b60ff821660141415611888575060026113d6565b60ff82166015141561189c575060016113d6565b60ff8216601614156118b0575060026113d6565b60ff8216601714156118c4575060026113d6565b60ff8216601814156118d8575060026113d6565b60ff8216601914156118ec575060016113d6565b60ff8216601a1415611900575060046113d6565b60ff8216601b1415611914575060076113d6565b60ff821660201415611928575060076113d6565b60ff82166021141561193c575060036113d6565b60ff821660221415611950575060086113d6565b60ff821660301415611964575060016113d6565b60ff821660311415611978575060016113d6565b60ff82166032141561198c575060016113d6565b60ff8216603314156119a0575060026113d6565b60ff8216603414156119b4575060046113d6565b60ff8216603514156119c8575060046113d6565b60ff8216603614156119dc575060026113d6565b60ff8216603714156119f0575060016113d6565b60ff821660381415611a04575060016113d6565b60ff821660391415611a18575060016113d6565b60ff8216603a1415611a2c575060026113d6565b60ff8216603b1415611a40575060016113d6565b60ff8216603c1415611a54575060016113d6565b60ff8216603d1415611a68575060016113d6565b60ff821660401415611a7c575060016113d6565b60ff821660411415611a90575060016113d6565b60ff821660421415611aa4575060016113d6565b60ff821660431415611ab8575060016113d6565b60ff821660441415611acc575060016113d6565b60ff821660501415611ae0575060026113d6565b60ff821660511415611af4575060286113d6565b60ff821660521415611b08575060026113d6565b60ff821660601415611b1c575060646113d6565b60ff821660611415611b30575060646113d6565b60ff821660701415611b44575060646113d6565b60ff821660711415611b58575060286113d6565b60ff821660721415611b6c575060286113d6565b60ff821660731415611b80575060056113d6565b60ff821660741415611b945750600a6113d6565b60ff821660801415611ba95750614e206113d6565b6040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206f70636f64653a206f70476173436f737428290000000000604482015290519081900360640190fd5b6000611c0183612ec5565b1580611c135750611c1182612ec5565b155b15611c2057506000611c3f565b82518251808201611c37878263ffffffff612ed016565b600193505050505b9392505050565b6000611c5183612ec5565b1580611c635750611c6182612ec5565b155b15611c7057506000611c3f565b82518251808202611c37878263ffffffff612ed016565b6000611c9283612ec5565b1580611ca45750611ca282612ec5565b155b15611cb157506000611c3f565b82518251808203611c37878263ffffffff612ed016565b6000611cd383612ec5565b1580611ce55750611ce382612ec5565b155b15611cf257506000611c3f565b8251825180611d0657600092505050611c3f565b808204611c37878263ffffffff612ed016565b6000611d2483612ec5565b1580611d365750611d3482612ec5565b155b15611d4357506000611c3f565b8251825180611d5757600092505050611c3f565b808205611c37878263ffffffff612ed016565b6000611d7583612ec5565b1580611d875750611d8582612ec5565b155b15611d9457506000611c3f565b8251825180611da857600092505050611c3f565b808206611c37878263ffffffff612ed016565b6000611dc683612ec5565b1580611dd85750611dd682612ec5565b155b15611de557506000611c3f565b8251825180611df957600092505050611c3f565b808207611c37878263ffffffff612ed016565b6000611e1784612ec5565b1580611e295750611e2783612ec5565b155b15611e3657506000611e6e565b83518351835180611e4d5760009350505050611e6e565b6000818385089050611e65898263ffffffff612ed016565b60019450505050505b949350505050565b6000611e8184612ec5565b1580611e935750611e9183612ec5565b155b15611ea057506000611e6e565b83518351835180611eb75760009350505050611e6e565b6000818385099050611e65898263ffffffff612ed016565b6000611eda83612ec5565b1580611eec5750611eea82612ec5565b155b15611ef957506000611c3f565b8251825180820a611c37878263ffffffff612ed016565b6000611f1b83612ec5565b1580611f2d5750611f2b82612ec5565b155b15611f3a57506000611c3f565b82518251808210611c37878263ffffffff612ed016565b6000611f5c83612ec5565b1580611f6e5750611f6c82612ec5565b155b15611f7b57506000611c3f565b82518251808211611c37878263ffffffff612ed016565b6000611f9d83612ec5565b1580611faf5750611fad82612ec5565b155b15611fbc57506000611c3f565b82518251808212611c37878263ffffffff612ed016565b6000611fde83612ec5565b1580611ff05750611fee82612ec5565b155b15611ffd57506000611c3f565b82518251808213611c37878263ffffffff612ed016565b600061203461113c61202584612d96565b61202e86612d96565b14612ee6565b5060019392505050565b600061204982612ec5565b6120635761205e83600063ffffffff612ed016565b61207a565b81518015612077858263ffffffff612ed016565b50505b50600192915050565b600061208e83612ec5565b15806120a0575061209e82612ec5565b155b156120ad57506000611c3f565b82518251808216611c37878263ffffffff612ed016565b60006120cf83612ec5565b15806120e157506120df82612ec5565b155b156120ee57506000611c3f565b82518251808217611c37878263ffffffff612ed016565b600061211083612ec5565b1580612122575061212082612ec5565b155b1561212f57506000611c3f565b82518251808218611c37878263ffffffff612ed016565b600061215182612ec5565b61215d57506000610276565b81518019612171858263ffffffff612ed016565b506001949350505050565b600061218783612ec5565b1580612199575061219782612ec5565b155b156121a657506000611c3f565b8251825181811a611c37878263ffffffff612ed016565b60006121c883612ec5565b15806121da57506121d882612ec5565b155b156121e757506000611c3f565b8251825181810b611c37878263ffffffff612ed016565b600061207a61220c83612d96565b849063ffffffff612ed016565b600061207a61222783612f08565b849063ffffffff6124fd16565b600061223f83612ec5565b1580612251575061224f82612ec5565b155b1561225e57506000611c3f565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120611c37878263ffffffff612ed016565b600192915050565b60006122bb8260800151836124fd90919063ffffffff16565b506001919050565b60006122bb8260600151836124fd90919063ffffffff16565b60609190910152600190565b60006122f382612f91565b6122ff57506000610276565b61230882612d96565b835250600192915050565b600061231e83612f91565b61232a57506000611c3f565b61233382612ec5565b61233f57506000611c3f565b8151156120345761234f83612d96565b84525060019392505050565b60006122bb61237861236b612f9e565b61202e8560200151612d96565b839063ffffffff6124fd16565b600061207a61222783612fbf565b600061207a838363ffffffff6124e316565b6000806123b0613ece565b845184106123d0576000846123c560006127d3565b9250925092506124dc565b60008085905060008782815181106123e457fe5b016020015160019092019160f81c905060006123fe613f58565b60ff8316612432576124108a85613044565b919650945091508484612422846127d3565b97509750975050505050506124dc565b60ff83166001141561245a576124488a85613097565b919650945090508484612422836131f7565b60ff831660021415612470576124228a8561325e565b600360ff8416108015906124875750600c60ff8416105b156124c2576002198301606061249e828d88613303565b9198509650905086866124b083612858565b995099509950505050505050506124dc565b6000806124cf60006127d3565b9199509750955050505050505b9250925092565b6124f18260400151826133c1565b82604001819052505050565b61250b8260200151826133c1565b82602001819052505050565b60006122bb612378612527612f9e565b61202e8560400151612d96565b60006122bb6123788360a00151612fbf565b600061255182612f91565b61255d57506000610276565b61256682612d96565b60a084015250600192915050565b6000612586838363ffffffff6124fd16565b61207a838363ffffffff6124fd16565b60006125a8848363ffffffff6124fd16565b6125b8848463ffffffff6124fd16565b612034848363ffffffff6124fd16565b60006125da858363ffffffff6124fd16565b6125ea858463ffffffff6124fd16565b6125fa858563ffffffff6124fd16565b612171858363ffffffff6124fd16565b60006125b8848463ffffffff6124fd16565b600061262e858563ffffffff6124fd16565b6125fa858463ffffffff6124fd16565b600061264983612ec5565b158061265b57506126598261343f565b155b1561266857506000611c3f565b6126718261344e565b60ff1683600001511061268657506000611c3f565b612034826040015184600001518151811061269d57fe5b6020026020010151856124fd90919063ffffffff16565b60006126bf8361343f565b15806126d157506126cf84612ec5565b155b156126de57506000611e6e565b6126e78361344e565b60ff168460000151106126fc57506000611e6e565b60408301518451815184918391811061271157fe5b602002602001018190525061273561272882612858565b879063ffffffff6124fd16565b50600195945050505050565b600061274c8261343f565b61275857506000610276565b61207a6127648361344e565b849060ff1663ffffffff612ed016565b50600190565b600080600161278884612d96565b915091505b9250929050565b600080612710836080015111156127b05750600090508061278d565b6127b98361345d565b6127c85750600090508061278d565b600161278884612d96565b6127db613ece565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612840565b61282d613ece565b8152602001906001900390816128255790505b50815260006020820152600160409091015292915050565b612860613ece565b61286a8251613574565b6128bb576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156128f2578381815181106128d557fe5b6020026020010151608001518201915080806001019150506128c0565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b600061295284612ec5565b61295e57506000611e6e565b83518211158061297d5750612971612f9e565b61297a84612d96565b14155b6129ce576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b612171858463ffffffff6124fd16565b600260c090910152565b60006129f385612ec5565b1580612a055750612a0384612ec5565b155b80612a165750612a1483612ec5565b155b80612a275750612a2582612ec5565b155b15612a3457506000612afe565b84518451845115801590612a4a57508451600114155b15612a6b57612a6088600063ffffffff612ed016565b600192505050612afe565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612acd573d6000803e3d6000fd5b5050604051601f1901519150612af490508b6001600160a01b03831663ffffffff612ed016565b6001955050505050505b95945050505050565b600160c090910152565b600060028260c001511415612b28575060006113d6565b60018260c001511415612b3d575060016113d6565b81516020830151612b4d90612d96565b612b5a8460400151612d96565b612b678560600151612d96565b612b748660800151612d96565b8660a001516040516020018087815260200186815260200185815260200184815260200183815260200182815260200196505050505050506040516020818303038152906040528051906020012090506113d6565b600060c090910152565b600080612bde613f02565b612be6613f02565b600060c08201819052612bf9878761357b565b84529650905080612c1357506000935084925090506124dc565b612c1d878761325e565b60208501529650905080612c3a57506000935084925090506124dc565b612c44878761325e565b60408501529650905080612c6157506000935084925090506124dc565b612c6b87876123a5565b60608501529650905080612c8857506000935084925090506124dc565b612c9287876123a5565b60808501529650905080612caf57506000935084925090506124dc565b612cb9878761357b565b60a08501529650905080612cd657506000935084925090506124dc565b506001969495509392505050565b612cec613f02565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b6000806000612d538460ff166135cf565b50949350505050565b612d64613ece565b611c3f60405180608001604052808560ff1681526020018481526020016000151581526020016000801b8152506131f7565b606081015160009060ff16612db7578151612db090613a9b565b90506113d6565b606082015160ff1660011415612dea576020808301518051604082015160608301519290930151612db093919290613abf565b606082015160ff1660021415612e0357612db082613b67565b600360ff16826060015160ff1610158015612e2757506060820151600c60ff909116105b15612e3557612db082613bcd565b606082015160ff1660641415612e4d575080516113d6565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b612e96613ece565b611e6e60405180608001604052808660ff168152602001858152602001600115158152602001848152506131f7565b6060015160ff161590565b61250b8260200151612ee1836127d3565b6133c1565b612eee613ece565b8115612efe57612db060016127d3565b612db060006127d3565b612f10613ece565b816060015160ff1660021415612f575760405162461bcd60e51b8152600401808060200182810382526021815260200180613ff16021913960400191505060405180910390fd5b606082015160ff16612f6d57612db060006127d3565b816060015160ff1660011415612f8757612db060016127d3565b612db060036127d3565b6060015160ff1660011490565b60408051600080825260208201909252612fb9816001613beb565b91505090565b612fc7613ece565b6040805160a081018252838152815160808101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161302c565b613019613ece565b8152602001906001900390816130115790505b50815260646020820152600160409091015292915050565b600080600080855190508481108061305e57506020858203105b156130735750600092508391508290506124dc565b600160208601613089888863ffffffff613c0a16565b935093509350509250925092565b6000806130a2613f58565b600084905060008682815181106130b557fe5b602001015160f81c60f81b60f81c9050818060010192505060008783815181106130db57fe5b016020015160019384019360f89190911c915060009060ff84161415613161576000613105613ece565b61310f8b876123a5565b909750909250905081613153575050604080516080810182526000808252602082018190529181018290526060810182905290975088965094506124dc9350505050565b61315c81612d96565b925050505b60006131738a8663ffffffff613c0a16565b90506020850194508360ff16600114156131bf576040805160808101825260ff909416845260208401919091526001908301819052606083019190915295509193509091506124dc9050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b6131ff613ece565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190613246565b613233613ece565b81526020019060019003908161322b5790505b50815260016020820181905260409091015292915050565b600080613269613ece565b613271613ece565b855160009081908781108061328857506040888203105b156132a05760008885965096509650505050506124dc565b60006132b28a8a63ffffffff613c0a16565b90506020890198506132c48a8a613044565b909a509450925082156132ef576132db81856101f6565b6001985089975095506124dc945050505050565b6000898697509750975050505050506124dc565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561334e57816020015b61333b613ece565b8152602001906001900390816133335790505b50905060005b8960ff168160ff1610156133ab5761336c89856123a5565b8451859060ff861690811061337d57fe5b602090810291909101015294509250826133a357506000955086945092506133b8915050565b600101613354565b5060019550919350909150505b93509350939050565b6133c9613ece565b6040805160028082526060828101909352816020015b6133e7613ece565b8152602001906001900390816133df579050509050828160008151811061340a57fe5b6020026020010181905250838160018151811061342357fe5b6020026020010181905250611e6e61343a82612858565b613c26565b60006102768260600151613c9c565b60006102768260600151613cba565b606081015160009060ff16613474575060016113d6565b606082015160ff166001141561348c575060006113d6565b606082015160ff16600214156134e0576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff161015801561350457506060820151600c60ff909116105b1561355c5760408201515160005b818110156135515761353a8460400151828151811061352d57fe5b602002602001015161345d565b613549576000925050506113d6565b600101613512565b5060019150506113d6565b606082015160ff1660641415612e4d575060006113d6565b6008101590565b6000806000806000865190508581108061359757506020868203105b156135ab57506000935084925090506124dc565b6135bb878763ffffffff613c0a16565b6001955060208701945092506124dc915050565b60008060018314156135e75750600290506001613a96565b60028314156135fc5750600290506001613a96565b60038314156136115750600290506001613a96565b60048314156136265750600290506001613a96565b600583141561363b5750600290506001613a96565b60068314156136505750600290506001613a96565b60078314156136655750600290506001613a96565b600883141561367a5750600390506001613a96565b600983141561368f5750600390506001613a96565b600a8314156136a45750600290506001613a96565b60108314156136b95750600290506001613a96565b60118314156136ce5750600290506001613a96565b60128314156136e35750600290506001613a96565b60138314156136f85750600290506001613a96565b601483141561370d5750600290506001613a96565b601583141561372157506001905080613a96565b60168314156137365750600290506001613a96565b601783141561374b5750600290506001613a96565b60188314156137605750600290506001613a96565b601983141561377457506001905080613a96565b601a8314156137895750600290506001613a96565b601b83141561379e5750600290506001613a96565b60208314156137b257506001905080613a96565b60218314156137c657506001905080613a96565b60228314156137db5750600290506001613a96565b60308314156137f05750600190506000613a96565b60318314156138055750600090506001613a96565b603283141561381a5750600090506001613a96565b603383141561382f5750600190506000613a96565b60348314156138445750600190506000613a96565b60358314156138595750600290506000613a96565b603683141561386e5750600090506001613a96565b60378314156138835750600090506001613a96565b60388314156138985750600190506000613a96565b60398314156138ad5750600090506001613a96565b603a8314156138c25750600090506001613a96565b603b8314156138d657506000905080613a96565b603c8314156138eb5750600090506001613a96565b603d8314156139005750600190506000613a96565b60408314156139155750600190506002613a96565b604183141561392a5750600290506003613a96565b604283141561393f5750600390506004613a96565b604383141561395357506002905080613a96565b604483141561396757506003905080613a96565b605083141561397c5750600290506001613a96565b60518314156139915750600390506001613a96565b60528314156139a557506001905080613a96565b60608314156139b957506000905080613a96565b60618314156139ce5750600190506000613a96565b60708314156139e35750600190506000613a96565b60718314156139f85750600090506001613a96565b6072831415613a0c57506001905080613a96565b6073831415613a2057506000905080613a96565b6074831415613a3457506000905080613a96565b6080831415613a495750600490506001613a96565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613b19575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611e6e565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214613bbc576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b815160808301516102769190613cdd565b6000613bd7613ece565b613be083613c26565b9050611c3f81613b67565b6000613bf5613ece565b613bff8484613d17565b9050611e6e81613b67565b60008160200183511015613c1d57600080fd5b50016020015190565b613c2e613ece565b613c378261343f565b613c7d576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060613c8c8360400151613d36565b9050611c3f818460800151613d17565b6000600c60ff8316108015610276575050600360ff91909116101590565b6000613cc582613c9c565b15613cd5575060021981016113d6565b5060016113d6565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b613d1f613ece565b6000613d2a84613e0e565b9050611e6e81846101f6565b6060600882511115613d86576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613db3578160200160208202803883390190505b50805190915060005b81811015613e05576000613de2868381518110613dd557fe5b6020026020010151612d96565b905080848381518110613df157fe5b602090810291909101015250600101613dbc565b50909392505050565b6000600882511115613e5e576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613ea2578181015183820152602001613e8a565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001613ee8613f58565b815260606020820181905260006040830181905291015290565b6040805160e081019091526000815260208101613f1d613ece565b8152602001613f2a613ece565b8152602001613f37613ece565b8152602001613f44613ece565b815260006020820181905260409091015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820229335eaa1cf0eb527c5753c4d3a7e7b913eae735a1015a4b0dc04a3cd2b3cf664736f6c634300050f0032"

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

// ValidateProof is a free data retrieval call binding the contract method 0xe987d887.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProofTester *OneStepProofTesterCaller) ValidateProof(opts *bind.CallOpts, beforeHash [32]byte, timeBounds [4]*big.Int, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProofTester.contract.Call(opts, out, "validateProof", beforeHash, timeBounds, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0xe987d887.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProofTester *OneStepProofTesterSession) ValidateProof(beforeHash [32]byte, timeBounds [4]*big.Int, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProofTester.Contract.ValidateProof(&_OneStepProofTester.CallOpts, beforeHash, timeBounds, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0xe987d887.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[4] timeBounds, bytes32 beforeInbox, uint256 beforeInboxValueSize, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) pure returns(uint256)
func (_OneStepProofTester *OneStepProofTesterCallerSession) ValidateProof(beforeHash [32]byte, timeBounds [4]*big.Int, beforeInbox [32]byte, beforeInboxValueSize *big.Int, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProofTester.Contract.ValidateProof(&_OneStepProofTester.CallOpts, beforeHash, timeBounds, beforeInbox, beforeInboxValueSize, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a5d06abf84b56d18eed8febe66af8b55acb83311c57c3b654677213e0da0a90664736f6c634300050f0032"

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
