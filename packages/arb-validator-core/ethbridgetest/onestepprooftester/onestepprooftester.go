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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158207d8758ad86f665628dab702fc2bf77ffff692cbcea924b7642d8a631868fabc164736f6c634300050f0032"

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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820c7db00460d391763f5cfa9f3926f78aec8377d33c42f7c809b3542ed12d91f1464736f6c634300050f0032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582074923e07baddec034e6b56d8d65b267442837ae79dd17420b5bcae408b4ec25a64736f6c634300050f0032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b5061400f806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f0a516ba14610030575b600080fd5b610123600480360361016081101561004757600080fd5b813591602081013591604082013591606081013591608082013515159160a08101359160c08201359160e0810135916101008201359167ffffffffffffffff610120820135169181019061016081016101408201356401000000008111156100ae57600080fd5b8201836020820111156100c057600080fd5b803590602001918460018302840111640100000000831117156100e257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610135945050505050565b60408051918252519081900360200190f35b600061014a8c8c8c8c8c8c8c8c8c8c8c61015a565b9c9b505050505050505050505050565b600061014a6040518061014001604052808e815260200161017b8e8e6101bf565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff1681526020018481525061024c565b6101c7613de3565b6040805160a080820183528582528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191610233565b610220613de3565b8152602001906001900390816102185790505b5081526002602082015260400183905290505b92915050565b6000806000806000606061025e613e17565b610266613e17565b61026f896114bf565b6101008f0151959c50939a509297509095509350915060019060009067ffffffffffffffff1688146102df576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8a6060015180156102f3575060ff89166072145b8061030f57508a6060015115801561030f575060ff8916607214155b610360576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a080840180518a900390528401518811156103875760001960a084015260009150611313565b60ff8916600114156103cd576103c683866000815181106103a457fe5b6020026020010151876001815181106103b957fe5b6020026020010151611819565b9150611313565b60ff89166002141561040c576103c683866000815181106103ea57fe5b6020026020010151876001815181106103ff57fe5b6020026020010151611869565b60ff89166003141561044b576103c6838660008151811061042957fe5b60200260200101518760018151811061043e57fe5b60200260200101516118aa565b60ff89166004141561048a576103c6838660008151811061046857fe5b60200260200101518760018151811061047d57fe5b60200260200101516118eb565b60ff8916600514156104c9576103c683866000815181106104a757fe5b6020026020010151876001815181106104bc57fe5b602002602001015161193c565b60ff891660061415610508576103c683866000815181106104e657fe5b6020026020010151876001815181106104fb57fe5b602002602001015161198d565b60ff891660071415610547576103c6838660008151811061052557fe5b60200260200101518760018151811061053a57fe5b60200260200101516119de565b60ff89166008141561059b576103c6838660008151811061056457fe5b60200260200101518760018151811061057957fe5b60200260200101518860028151811061058e57fe5b6020026020010151611a2f565b60ff8916600914156105ef576103c683866000815181106105b857fe5b6020026020010151876001815181106105cd57fe5b6020026020010151886002815181106105e257fe5b6020026020010151611a99565b60ff8916600a141561062e576103c6838660008151811061060c57fe5b60200260200101518760018151811061062157fe5b6020026020010151611af2565b60ff89166010141561066d576103c6838660008151811061064b57fe5b60200260200101518760018151811061066057fe5b6020026020010151611b33565b60ff8916601114156106ac576103c6838660008151811061068a57fe5b60200260200101518760018151811061069f57fe5b6020026020010151611b74565b60ff8916601214156106eb576103c683866000815181106106c957fe5b6020026020010151876001815181106106de57fe5b6020026020010151611bb5565b60ff89166013141561072a576103c6838660008151811061070857fe5b60200260200101518760018151811061071d57fe5b6020026020010151611bf6565b60ff891660141415610769576103c6838660008151811061074757fe5b60200260200101518760018151811061075c57fe5b6020026020010151611c37565b60ff891660151415610793576103c6838660008151811061078657fe5b6020026020010151611c6e565b60ff8916601614156107d2576103c683866000815181106107b057fe5b6020026020010151876001815181106107c557fe5b6020026020010151611cb3565b60ff891660171415610811576103c683866000815181106107ef57fe5b60200260200101518760018151811061080457fe5b6020026020010151611cf4565b60ff891660181415610850576103c6838660008151811061082e57fe5b60200260200101518760018151811061084357fe5b6020026020010151611d35565b60ff89166019141561087a576103c6838660008151811061086d57fe5b6020026020010151611d76565b60ff8916601a14156108b9576103c6838660008151811061089757fe5b6020026020010151876001815181106108ac57fe5b6020026020010151611dac565b60ff8916601b14156108f8576103c683866000815181106108d657fe5b6020026020010151876001815181106108eb57fe5b6020026020010151611ded565b60ff891660201415610922576103c6838660008151811061091557fe5b6020026020010151611e2e565b60ff89166021141561094c576103c6838660008151811061093f57fe5b6020026020010151611e49565b60ff89166022141561098b576103c6838660008151811061096957fe5b60200260200101518760018151811061097e57fe5b6020026020010151611e64565b60ff8916603014156109b5576103c683866000815181106109a857fe5b6020026020010151611eca565b60ff8916603114156109ca576103c683611ed2565b60ff8916603214156109df576103c683611ef3565b60ff891660331415610a09576103c683866000815181106109fc57fe5b6020026020010151611f0c565b60ff891660341415610a33576103c68386600081518110610a2657fe5b6020026020010151611f18565b60ff891660351415610a72576103c68386600081518110610a5057fe5b602002602001015187600181518110610a6557fe5b6020026020010151611f43565b60ff891660361415610a87576103c683611f8b565b60ff891660371415610aa1576103c6838560000151611fb5565b60ff891660381415610acb576103c68386600081518110610abe57fe5b6020026020010151611fc5565b60ff891660391415610b5757610adf613de3565b610aee8c610120015188611fd7565b9199509750905087610b315760405162461bcd60e51b8152600401808060200182810382526021815260200180613fba6021913960400191505060405180910390fd5b610b41858263ffffffff6120f516565b610b51848263ffffffff61210f16565b50611313565b60ff8916603a1415610b6c576103c683612129565b60ff8916603b1415610b815760019150611313565b60ff8916603c1415610b96576103c683612146565b60ff8916603d1415610bc0576103c68386600081518110610bb357fe5b602002602001015161215a565b60ff891660401415610bea576103c68386600081518110610bdd57fe5b6020026020010151612188565b60ff891660411415610c29576103c68386600081518110610c0757fe5b602002602001015187600181518110610c1c57fe5b60200260200101516121aa565b60ff891660421415610c7d576103c68386600081518110610c4657fe5b602002602001015187600181518110610c5b57fe5b602002602001015188600281518110610c7057fe5b60200260200101516121dc565b60ff891660431415610cbc576103c68386600081518110610c9a57fe5b602002602001015187600181518110610caf57fe5b602002602001015161221e565b60ff891660441415610d10576103c68386600081518110610cd957fe5b602002602001015187600181518110610cee57fe5b602002602001015188600281518110610d0357fe5b6020026020010151612230565b60ff891660501415610d4f576103c68386600081518110610d2d57fe5b602002602001015187600181518110610d4257fe5b6020026020010151612252565b60ff891660511415610da3576103c68386600081518110610d6c57fe5b602002602001015187600181518110610d8157fe5b602002602001015188600281518110610d9657fe5b60200260200101516122c8565b60ff891660521415610dcd576103c68386600081518110610dc057fe5b6020026020010151612355565b60ff891660531415610e6a57610de1613de3565b610df08c610120015188611fd7565b9199509750905087610e335760405162461bcd60e51b8152600401808060200182810382526021815260200180613fba6021913960400191505060405180910390fd5b610e43858263ffffffff6120f516565b610e628487600081518110610e5457fe5b602002602001015183612388565b925050611313565b60ff891660541415610f1457610e7e613de3565b610e8d8c610120015188611fd7565b9199509750905087610ed05760405162461bcd60e51b8152600401808060200182810382526021815260200180613fba6021913960400191505060405180910390fd5b610ee0858263ffffffff6120f516565b610e628487600081518110610ef157fe5b602002602001015188600181518110610f0657fe5b6020026020010151846123e0565b60ff891660601415610f29576103c683612461565b60ff89166061141561102657610f538386600081518110610f4657fe5b6020026020010151612467565b9092509050811561101d578a60e001518b60c001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610fd25760405162461bcd60e51b8152600401808060200182810382526025815260200180613f6e6025913960400191505060405180910390fd5b8a60a001518b60800151146110185760405162461bcd60e51b8152600401808060200182810382526027815260200180613f936027913960400191505060405180910390fd5b611021565b5060005b611313565b60ff89166070141561116557611050838660008151811061104357fe5b6020026020010151612481565b9092509050811561101d57806110ab578a60a001518b60800151146110a65760405162461bcd60e51b8152600401808060200182810382526038815260200180613f366038913960400191505060405180910390fd5b611018565b8a60a001518b60800151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001201461111f5760405162461bcd60e51b8152600401808060200182810382526029815260200180613ec66029913960400191505060405180910390fd5b8a60e001518b60c00151146110185760405162461bcd60e51b8152600401808060200182810382526026815260200180613eef6026913960400191505060405180910390fd5b60ff89166072141561117f576103c6838c602001516124c0565b60ff8916607314156111945760009150611313565b60ff8916607414156111a95761102183612526565b60ff8916607514156111d3576103c683866000815181106111c657fe5b6020026020010151612530565b60ff8916607614156111e8576103c683612555565b60ff8916607714156111fd576103c68361256e565b60ff89166078141561123c576103c6838660008151811061121a57fe5b60200260200101518760018151811061122f57fe5b60200260200101516125b7565b60ff891660791415611290576103c6838660008151811061125957fe5b60200260200101518760018151811061126e57fe5b60200260200101518860028151811061128357fe5b60200260200101516125fc565b60ff8916607b14156112a5576103c68361264f565b60ff89166080141561130e576103c683866000815181106112c257fe5b6020026020010151876001815181106112d757fe5b6020026020010151886002815181106112ec57fe5b60200260200101518960038151811061130157fe5b6020026020010151612692565b600091505b806113a4578a60a001518b608001511461135e5760405162461bcd60e51b8152600401808060200182810382526027815260200180613f936027913960400191505060405180910390fd5b8a60e001518b60c00151146113a45760405162461bcd60e51b8152600401808060200182810382526026815260200180613eef6026913960400191505060405180910390fd5b816114055760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156113fd576113f8836127b1565b611405565b60c083015183525b61140e846127bb565b8b511461144c5760405162461bcd60e51b8152600401808060200182810382526022815260200180613ea46022913960400191505060405180910390fd5b611455836127bb565b8b60400151146114ac576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b600099505050505050505050505b919050565b60008060606114cc613e17565b6114d4613e17565b6000806114e08461287f565b6114ef88610120015183612889565b95509250905080611547576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b611550846129c1565b92506000886101200151838151811061156557fe5b602001015160f81c60f81b60f81c9050886101200151836001018151811061158957fe5b016020015160f81c9750600061159e89612a2a565b6040805183815260208085028201019091529099509091508180156115dd57816020015b6115ca613de3565b8152602001906001900390816115c25790505b5096506002840193508160ff16600014806115fb57508160ff166001145b61164c576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff82166116715761166a6116658a8860000151612f87565b612fc8565b865261172c565b611679613de3565b6116888b610120015186611fd7565b9096509094509050836116e2576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b81156117065780886000815181106116f657fe5b6020026020010181905250611716565b611716868263ffffffff61210f16565b6117286116658b8960000151846130c0565b8752505b60ff82165b818110156117bf576117488b610120015186611fd7565b8a518b908590811061175657fe5b602090810291909101015295509350836117b7576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611731565b87511561180c575060005b8260ff1688510381101561180c57611804888260018b510303815181106117ed57fe5b60200260200101518861210f90919063ffffffff16565b6001016117ca565b5050505091939550919395565b600061182483613109565b1580611836575061183482613109565b155b1561184357506000611862565b8251825180820161185a878263ffffffff61311416565b600193505050505b9392505050565b600061187483613109565b1580611886575061188482613109565b155b1561189357506000611862565b8251825180820261185a878263ffffffff61311416565b60006118b583613109565b15806118c757506118c582613109565b155b156118d457506000611862565b8251825180820361185a878263ffffffff61311416565b60006118f683613109565b1580611908575061190682613109565b155b1561191557506000611862565b825182518061192957600092505050611862565b80820461185a878263ffffffff61311416565b600061194783613109565b1580611959575061195782613109565b155b1561196657506000611862565b825182518061197a57600092505050611862565b80820561185a878263ffffffff61311416565b600061199883613109565b15806119aa57506119a882613109565b155b156119b757506000611862565b82518251806119cb57600092505050611862565b80820661185a878263ffffffff61311416565b60006119e983613109565b15806119fb57506119f982613109565b155b15611a0857506000611862565b8251825180611a1c57600092505050611862565b80820761185a878263ffffffff61311416565b6000611a3a84613109565b1580611a4c5750611a4a83613109565b155b15611a5957506000611a91565b83518351835180611a705760009350505050611a91565b6000818385089050611a88898263ffffffff61311416565b60019450505050505b949350505050565b6000611aa484613109565b1580611ab65750611ab483613109565b155b15611ac357506000611a91565b83518351835180611ada5760009350505050611a91565b6000818385099050611a88898263ffffffff61311416565b6000611afd83613109565b1580611b0f5750611b0d82613109565b155b15611b1c57506000611862565b8251825180820a61185a878263ffffffff61311416565b6000611b3e83613109565b1580611b505750611b4e82613109565b155b15611b5d57506000611862565b8251825180821061185a878263ffffffff61311416565b6000611b7f83613109565b1580611b915750611b8f82613109565b155b15611b9e57506000611862565b8251825180821161185a878263ffffffff61311416565b6000611bc083613109565b1580611bd25750611bd082613109565b155b15611bdf57506000611862565b8251825180821261185a878263ffffffff61311416565b6000611c0183613109565b1580611c135750611c1182613109565b155b15611c2057506000611862565b8251825180821361185a878263ffffffff61311416565b6000611c64611c57611c4884612fc8565b611c5186612fc8565b1461312a565b859063ffffffff61210f16565b5060019392505050565b6000611c7982613109565b611c9357611c8e83600063ffffffff61311416565b611caa565b81518015611ca7858263ffffffff61311416565b50505b50600192915050565b6000611cbe83613109565b1580611cd05750611cce82613109565b155b15611cdd57506000611862565b8251825180821661185a878263ffffffff61311416565b6000611cff83613109565b1580611d115750611d0f82613109565b155b15611d1e57506000611862565b8251825180821761185a878263ffffffff61311416565b6000611d4083613109565b1580611d525750611d5082613109565b155b15611d5f57506000611862565b8251825180821861185a878263ffffffff61311416565b6000611d8182613109565b611d8d57506000610246565b81518019611da1858263ffffffff61311416565b506001949350505050565b6000611db783613109565b1580611dc95750611dc782613109565b155b15611dd657506000611862565b8251825181811a61185a878263ffffffff61311416565b6000611df883613109565b1580611e0a5750611e0882613109565b155b15611e1757506000611862565b8251825181810b61185a878263ffffffff61311416565b6000611caa611e3c83612fc8565b849063ffffffff61311416565b6000611caa611e578361314c565b849063ffffffff61210f16565b6000611e6f83613109565b1580611e815750611e7f82613109565b155b15611e8e57506000611862565b825182516040805160208082018590528183018490528251808303840181526060909201909252805191012061185a878263ffffffff61311416565b600192915050565b6000611eeb82608001518361210f90919063ffffffff16565b506001919050565b6000611eeb82606001518361210f90919063ffffffff16565b60609190910152600190565b6000611f23826131d5565b611f2f57506000610246565b611f3882612fc8565b835250600192915050565b6000611f4e836131d5565b611f5a57506000611862565b611f6382613109565b611f6f57506000611862565b815115611c6457611f7f83612fc8565b84525060019392505050565b6000611eeb611fa8611f9b6131e2565b611c518560200151612fc8565b839063ffffffff61210f16565b6000611caa611e57836001613203565b6000611caa838363ffffffff6120f516565b600080611fe2613de3565b8451841061200257600084611ff7600061328e565b9250925092506120ee565b600080859050600087828151811061201657fe5b016020015160019092019160f81c905060008161205857612037898461331a565b9195509350905083836120498361328e565b965096509650505050506120ee565b60ff82166001141561206e57612049898461336d565b60ff82166002141561208457612049898461345d565b600360ff83161080159061209b5750600c60ff8316105b156120d557600219820160606120b2828c87613502565b9197509550905085856120c4836135c0565b9850985098505050505050506120ee565b6000806120e2600061328e565b91985096509450505050505b9250925092565b6121038260400151826136b7565b82604001819052505050565b61211d8260200151826136b7565b82602001819052505050565b6000611eeb611fa86121396131e2565b611c518560400151612fc8565b6000611eeb611fa88360c001516001613203565b6000612165826131d5565b61217157506000610246565b61217a82612fc8565b60c084015250600192915050565b600061219a838363ffffffff61210f16565b611caa838363ffffffff61210f16565b60006121bc848363ffffffff61210f16565b6121cc848463ffffffff61210f16565b611c64848363ffffffff61210f16565b60006121ee858363ffffffff61210f16565b6121fe858463ffffffff61210f16565b61220e858563ffffffff61210f16565b611da1858363ffffffff61210f16565b60006121cc848463ffffffff61210f16565b6000612242858563ffffffff61210f16565b61220e858463ffffffff61210f16565b600061225d83613109565b158061226f575061226d82613735565b155b1561227c57506000611862565b61228582613744565b60ff1683600001511061229a57506000611862565b611c6482604001518460000151815181106122b157fe5b60200260200101518561210f90919063ffffffff16565b60006122d383613735565b15806122e557506122e384613109565b155b156122f257506000611a91565b6122fb83613744565b60ff1684600001511061231057506000611a91565b60408301518451815184918391811061232557fe5b602002602001018190525061234961233c826135c0565b879063ffffffff61210f16565b50600195945050505050565b600061236082613735565b61236c57506000610246565b611caa61237883613744565b849060ff1663ffffffff61311416565b600061239383613109565b15806123a557506123a382613735565b155b156123b257506000611862565b6123bb82613744565b60ff168360000151106123d057506000611862565b61229a848363ffffffff6120f516565b60006123eb82613735565b15806123fd57506123fb84613109565b155b1561240a57506000611a91565b61241382613744565b60ff1684600001511061242857506000611a91565b60408201518451815185918391811061243d57fe5b6020026020010181905250612349612454826135c0565b879063ffffffff6120f516565b50600190565b600080600161247584612fc8565b915091505b9250929050565b6000806127108360800151111561249d5750600090508061247a565b6124a68361376f565b6124b55750600090508061247a565b600161247584612fc8565b60006124ca6131e2565b6124d383612fc8565b141561219a576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b600260e090910152565b600061253b82613109565b61254757506000610246565b505160a09190910152600190565b6000611eeb8260a001518361311490919063ffffffff16565b60408051600160f81b60208083019190915260006021830181905260228084018290528451808503909101815260429093019093528151910120611eeb90611fa8906001613203565b60006125c283613109565b6125ce57506000611862565b6125d7826131d5565b6125e357506000611862565b611c64611c5784600001516125f785612fc8565b612f87565b600061260784613109565b61261357506000611a91565b61261c826131d5565b61262857506000611a91565b611da1612642856000015161263c85612fc8565b866130c0565b869063ffffffff61210f16565b6040805160008082526020820190925260609082612683565b612670613de3565b8152602001906001900390816126685790505b509050611caa611e57826135c0565b600061269d85613109565b15806126af57506126ad84613109565b155b806126c057506126be83613109565b155b806126d157506126cf82613109565b155b156126de575060006127a8565b845184518451158015906126f457508451600114155b156127155761270a88600063ffffffff61311416565b6001925050506127a8565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612777573d6000803e3d6000fd5b5050604051601f190151915061279e90508b6001600160a01b03831663ffffffff61311416565b6001955050505050505b95945050505050565b600160e090910152565b600060028260e0015114156127d2575060006114ba565b60018260e0015114156127e7575060016114ba565b815160208301516127f790612fc8565b6128048460400151612fc8565b6128118560600151612fc8565b61281e8660800151612fc8565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090506114ba565b600060e090910152565b600080612894613e17565b61289c613e17565b600060e082018190526128af8787613886565b845296509050806128c957506000935084925090506120ee565b6128d3878761345d565b602085015296509050806128f057506000935084925090506120ee565b6128fa878761345d565b6040850152965090508061291757506000935084925090506120ee565b6129218787611fd7565b6060850152965090508061293e57506000935084925090506120ee565b6129488787611fd7565b6080850152965090508061296557506000935084925090506120ee565b61296f878761331a565b60a0850152965090508061298c57506000935084925090506120ee565b6129968787613886565b60c085015296509050806129b357506000935084925090506120ee565b506001969495509392505050565b6129c9613e17565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612a425750600290506003612f82565b6002831415612a575750600290506003612f82565b6003831415612a6c5750600290506003612f82565b6004831415612a815750600290506004612f82565b6005831415612a965750600290506007612f82565b6006831415612aab5750600290506004612f82565b6007831415612ac05750600290506007612f82565b6008831415612ad55750600390506004612f82565b6009831415612aea5750600390506004612f82565b600a831415612aff5750600290506019612f82565b6010831415612b1357506002905080612f82565b6011831415612b2757506002905080612f82565b6012831415612b3b57506002905080612f82565b6013831415612b4f57506002905080612f82565b6014831415612b6357506002905080612f82565b6015831415612b7757506001905080612f82565b6016831415612b8b57506002905080612f82565b6017831415612b9f57506002905080612f82565b6018831415612bb357506002905080612f82565b6019831415612bc757506001905080612f82565b601a831415612bdc5750600290506004612f82565b601b831415612bf15750600290506007612f82565b6020831415612c065750600190506007612f82565b6021831415612c1b5750600190506003612f82565b6022831415612c305750600290506008612f82565b6030831415612c4457506001905080612f82565b6031831415612c595750600090506001612f82565b6032831415612c6e5750600090506001612f82565b6033831415612c835750600190506002612f82565b6034831415612c985750600190506004612f82565b6035831415612cad5750600290506004612f82565b6036831415612cc25750600090506002612f82565b6037831415612cd75750600090506001612f82565b6038831415612ceb57506001905080612f82565b6039831415612d005750600090506001612f82565b603a831415612d155750600090506002612f82565b603b831415612d2a5750600090506001612f82565b603c831415612d3f5750600090506001612f82565b603d831415612d5357506001905080612f82565b6040831415612d6757506001905080612f82565b6041831415612d7c5750600290506001612f82565b6042831415612d915750600390506001612f82565b6043831415612da65750600290506001612f82565b6044831415612dbb5750600390506001612f82565b6050831415612dcf57506002905080612f82565b6051831415612de45750600390506028612f82565b6052831415612df95750600190506002612f82565b6053831415612e0e5750600190506003612f82565b6054831415612e235750600290506029612f82565b6060831415612e385750600090506064612f82565b6061831415612e4d5750600190506064612f82565b6070831415612e625750600190506064612f82565b6072831415612e775750600090506028612f82565b6073831415612e8c5750600090506005612f82565b6074831415612ea1575060009050600a612f82565b6075831415612eb65750600190506000612f82565b6076831415612ecb5750600090506001612f82565b6077831415612ee05750600090506019612f82565b6078831415612ef55750600290506019612f82565b6079831415612f0a5750600390506019612f82565b607b831415612f1f575060009050600a612f82565b6080831415612f35575060049050614e20612f82565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b612f8f613de3565b6118626040518060a001604052808560ff1681526020018481526020016000151581526020016000801b815260200160008152506138da565b606081015160009060ff16612fe9578151612fe290613941565b90506114ba565b606082015160ff166001141561301c576020808301518051604082015160608301519290930151612fe293919290613965565b606082015160ff166002141561303557612fe282613a0d565b600360ff16826060015160ff161015801561305957506060820151600c60ff909116105b1561306757612fe282613a73565b606082015160ff166064141561307f575080516114ba565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6130c8613de3565b611a916040518060a001604052808660ff1681526020018581526020016001151581526020016130f785612fc8565b815260200184608001518152506138da565b6060015160ff161590565b61211d82602001516131258361328e565b6136b7565b613132613de3565b811561314257612fe2600161328e565b612fe2600061328e565b613154613de3565b816060015160ff166002141561319b5760405162461bcd60e51b8152600401808060200182810382526021815260200180613f156021913960400191505060405180910390fd5b606082015160ff166131b157612fe2600061328e565b816060015160ff16600114156131cb57612fe2600161328e565b612fe2600361328e565b6060015160ff1660011490565b604080516000808252602082019092526131fd816001613a91565b91505090565b61320b613de3565b6040805160a080820183528582528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191613277565b613264613de3565b81526020019060019003908161325c5790505b508152606460208201526040019290925250919050565b613296613de3565b6040805160a080820183528482528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191613302565b6132ef613de3565b8152602001906001900390816132e75790505b50815260006020820152600160409091015292915050565b600080600080855190508481108061333457506020858203105b156133495750600092508391508290506120ee565b60016020860161335f888863ffffffff613ab016565b935093509350509250925092565b600080613378613de3565b6000849050600086828151811061338b57fe5b602001015160f81c60f81b60f81c9050818060010192505060008783815181106133b157fe5b016020015160019093019260f81c90506133c9613de3565b8260ff166001141561340a5760006133e18a86611fd7565b9096509250905080613408576000896133f8613acc565b97509750975050505050506120ee565b505b600061341c8a8663ffffffff613ab016565b90506020850194508360ff166001141561343e576001856133f88584866130c0565b60018561344b8584612f87565b97509750975050505050509250925092565b600080613468613de3565b613470613de3565b855160009081908781108061348757506040888203105b1561349f5760008885965096509650505050506120ee565b60006134b18a8a63ffffffff613ab016565b90506020890198506134c38a8a61331a565b909a509450925082156134ee576134da81856101bf565b6001985089975095506120ee945050505050565b6000898697509750975050505050506120ee565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561354d57816020015b61353a613de3565b8152602001906001900390816135325790505b50905060005b8960ff168160ff1610156135aa5761356b8985611fd7565b8451859060ff861690811061357c57fe5b602090810291909101015294509250826135a257506000955086945092506135b7915050565b600101613553565b5060019550919350909150505b93509350939050565b6135c8613de3565b6135d28251613b57565b613623576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b835181101561365a5783818151811061363d57fe5b602002602001015160800151820191508080600101915050613628565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b6136bf613de3565b6040805160028082526060828101909352816020015b6136dd613de3565b8152602001906001900390816136d5579050509050828160008151811061370057fe5b6020026020010181905250838160018151811061371957fe5b6020026020010181905250611a91613730826135c0565b613b5e565b60006102468260600151613bd4565b60006137538260600151613bd4565b1561376757506060810151600219016114ba565b5060016114ba565b606081015160009060ff16613786575060016114ba565b606082015160ff166001141561379e575060006114ba565b606082015160ff16600214156137f2576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff161015801561381657506060820151600c60ff909116105b1561386e5760408201515160005b818110156138635761384c8460400151828151811061383f57fe5b602002602001015161376f565b61385b576000925050506114ba565b600101613824565b5060019150506114ba565b606082015160ff166064141561307f575060006114ba565b600080600080600086519050858110806138a257506020868203105b156138b657506000935084925090506120ee565b6138c6878763ffffffff613ab016565b6001955060208701945092506120ee915050565b6138e2613de3565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190613929565b613916613de3565b81526020019060019003908161390e5790505b50815260016020820181905260409091015292915050565b60408051602080820193909352815180820384018152908201909152805191012090565b600083156139bf575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611a91565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214613a62576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b815160808301516102469190613bf2565b6000613a7d613de3565b613a8683613b5e565b905061186281613a0d565b6000613a9b613de3565b613aa58484613c2c565b9050611a9181613a0d565b60008160200183511015613ac357600080fd5b50016020015190565b613ad4613de3565b6040805160a080820183526000808352835191820184528082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191613b40565b613b2d613de3565b815260200190600190039081613b255790505b508152600360208201526001604090910152905090565b6008101590565b613b66613de3565b613b6f82613735565b613bb5576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060613bc48360400151613c4b565b9050611862818460800151613c2c565b6000600c60ff8316108015610246575050600360ff91909116101590565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b613c34613de3565b6000613c3f84613d23565b9050611a9181846101bf565b6060600882511115613c9b576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613cc8578160200160208202803883390190505b50805190915060005b81811015613d1a576000613cf7868381518110613cea57fe5b6020026020010151612fc8565b905080848381518110613d0657fe5b602090810291909101015250600101613cd1565b50909392505050565b6000600882511115613d73576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613db7578181015183820152602001613d9f565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001613dfd613e75565b815260606020820181905260006040830181905291015290565b6040805161010081019091526000815260208101613e33613de3565b8152602001613e40613de3565b8152602001613e4d613de3565b8152602001613e5a613de3565b81526000602082018190526040820181905260609091015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820709a95f00cba8c07d9ab579beb41c22251ba2e952bd1f8784ed7309d78877e1c64736f6c634300050f0032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820756e12f9a680c5aa2d821f4845604d9e478c3e9e61704ebc03fff3bdaf85e52964736f6c634300050f0032"

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
