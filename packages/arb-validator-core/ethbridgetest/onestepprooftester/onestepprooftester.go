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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205130ee8069d4d45b671687a9bd9471ae9b0f29d9a9560a567b2a039c58a80e1a64736f6c63430005110032"

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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200e6ac2c9230f023e04ee523bd0426d233795c75c34f0f3306f61b14825b2bf0764736f6c63430005110032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582089906d69af6ed6b1627d2bf256e3e1c9086270fcbffd8ccd1f4e3071f1939b8364736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b50613fa9806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633c41485d14610030575b600080fd5b61011d600480360361014081101561004757600080fd5b813591602081013591604082013591606081013515159160808201359160a08101359160c08201359160e08101359167ffffffffffffffff61010083013516919081019061014081016101208201356401000000008111156100a857600080fd5b8201836020820111156100ba57600080fd5b803590602001918460018302840111640100000000831117156100dc57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061012f945050505050565b60408051918252519081900360200190f35b600061014b6101468c8c8c8c8c8c8c8c8c8c61015a565b6101bf565b9b9a5050505050505050505050565b610162613d7d565b61014b6040518061012001604052808d81526020016101818d8d610284565b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610311565b600060028260e0015114156101d65750600061027f565b60018260e0015114156101eb5750600161027f565b815160208301516101fb9061151d565b610208846040015161151d565b610215856060015161151d565b610222866080015161151d565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090505b919050565b61028c613ddb565b6040805160a0808201835285825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916102f8565b6102e5613ddb565b8152602001906001900390816102dd5790505b5081526002602082015260400183905290505b92915050565b610319613d7d565b6000806000806060610329613d7d565b610331613d7d565b61033a89611615565b60e08f0151959c50939a509297509095509350915060019060009067ffffffffffffffff1688146103a9576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8a6040015180156103bd575060ff89166072145b806103d957508a604001511580156103d9575060ff8916607214155b61042a576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a080840180518a900390528401518811156104515760001960a0840152600091506113d5565b60ff89166001141561049757610490838660008151811061046e57fe5b60200260200101518760018151811061048357fe5b602002602001015161196f565b91506113d5565b60ff8916600214156104d65761049083866000815181106104b457fe5b6020026020010151876001815181106104c957fe5b60200260200101516119bf565b60ff8916600314156105155761049083866000815181106104f357fe5b60200260200101518760018151811061050857fe5b6020026020010151611a00565b60ff89166004141561055457610490838660008151811061053257fe5b60200260200101518760018151811061054757fe5b6020026020010151611a41565b60ff89166005141561059357610490838660008151811061057157fe5b60200260200101518760018151811061058657fe5b6020026020010151611a92565b60ff8916600614156105d25761049083866000815181106105b057fe5b6020026020010151876001815181106105c557fe5b6020026020010151611ae3565b60ff8916600714156106115761049083866000815181106105ef57fe5b60200260200101518760018151811061060457fe5b6020026020010151611b34565b60ff89166008141561066557610490838660008151811061062e57fe5b60200260200101518760018151811061064357fe5b60200260200101518860028151811061065857fe5b6020026020010151611b85565b60ff8916600914156106b957610490838660008151811061068257fe5b60200260200101518760018151811061069757fe5b6020026020010151886002815181106106ac57fe5b6020026020010151611bef565b60ff8916600a14156106f85761049083866000815181106106d657fe5b6020026020010151876001815181106106eb57fe5b6020026020010151611c48565b60ff89166010141561073757610490838660008151811061071557fe5b60200260200101518760018151811061072a57fe5b6020026020010151611c89565b60ff89166011141561077657610490838660008151811061075457fe5b60200260200101518760018151811061076957fe5b6020026020010151611cca565b60ff8916601214156107b557610490838660008151811061079357fe5b6020026020010151876001815181106107a857fe5b6020026020010151611d0b565b60ff8916601314156107f45761049083866000815181106107d257fe5b6020026020010151876001815181106107e757fe5b6020026020010151611d4c565b60ff89166014141561083357610490838660008151811061081157fe5b60200260200101518760018151811061082657fe5b6020026020010151611d8d565b60ff89166015141561085d57610490838660008151811061085057fe5b6020026020010151611dc4565b60ff89166016141561089c57610490838660008151811061087a57fe5b60200260200101518760018151811061088f57fe5b6020026020010151611e09565b60ff8916601714156108db5761049083866000815181106108b957fe5b6020026020010151876001815181106108ce57fe5b6020026020010151611e4a565b60ff89166018141561091a5761049083866000815181106108f857fe5b60200260200101518760018151811061090d57fe5b6020026020010151611e8b565b60ff89166019141561094457610490838660008151811061093757fe5b6020026020010151611ecc565b60ff8916601a141561098357610490838660008151811061096157fe5b60200260200101518760018151811061097657fe5b6020026020010151611f02565b60ff8916601b14156109c25761049083866000815181106109a057fe5b6020026020010151876001815181106109b557fe5b6020026020010151611f43565b60ff8916602014156109ec5761049083866000815181106109df57fe5b6020026020010151611f84565b60ff891660211415610a16576104908386600081518110610a0957fe5b6020026020010151611f9f565b60ff891660221415610a55576104908386600081518110610a3357fe5b602002602001015187600181518110610a4857fe5b6020026020010151611fba565b60ff891660301415610a7f576104908386600081518110610a7257fe5b6020026020010151612020565b60ff891660311415610a945761049083612028565b60ff891660321415610aa95761049083612049565b60ff891660331415610ad3576104908386600081518110610ac657fe5b6020026020010151612062565b60ff891660341415610afd576104908386600081518110610af057fe5b602002602001015161206e565b60ff891660351415610b3c576104908386600081518110610b1a57fe5b602002602001015187600181518110610b2f57fe5b6020026020010151612099565b60ff891660361415610b5157610490836120e1565b60ff891660371415610b6b5761049083856000015161210b565b60ff891660381415610b95576104908386600081518110610b8857fe5b602002602001015161211b565b60ff891660391415610c2157610ba9613ddb565b610bb88c61010001518861212d565b9199509750905087610bfb5760405162461bcd60e51b8152600401808060200182810382526021815260200180613f546021913960400191505060405180910390fd5b610c0b858263ffffffff61224b16565b610c1b848263ffffffff61226516565b506113d5565b60ff8916603a1415610c36576104908361227f565b60ff8916603b1415610c4b57600191506113d5565b60ff8916603c1415610c60576104908361229c565b60ff8916603d1415610c8a576104908386600081518110610c7d57fe5b60200260200101516122b0565b60ff891660401415610cb4576104908386600081518110610ca757fe5b60200260200101516122de565b60ff891660411415610cf3576104908386600081518110610cd157fe5b602002602001015187600181518110610ce657fe5b6020026020010151612300565b60ff891660421415610d47576104908386600081518110610d1057fe5b602002602001015187600181518110610d2557fe5b602002602001015188600281518110610d3a57fe5b6020026020010151612332565b60ff891660431415610d86576104908386600081518110610d6457fe5b602002602001015187600181518110610d7957fe5b6020026020010151612374565b60ff891660441415610dda576104908386600081518110610da357fe5b602002602001015187600181518110610db857fe5b602002602001015188600281518110610dcd57fe5b6020026020010151612386565b60ff891660501415610e19576104908386600081518110610df757fe5b602002602001015187600181518110610e0c57fe5b60200260200101516123a8565b60ff891660511415610e6d576104908386600081518110610e3657fe5b602002602001015187600181518110610e4b57fe5b602002602001015188600281518110610e6057fe5b602002602001015161241e565b60ff891660521415610e97576104908386600081518110610e8a57fe5b60200260200101516124ab565b60ff891660531415610f3457610eab613ddb565b610eba8c61010001518861212d565b9199509750905087610efd5760405162461bcd60e51b8152600401808060200182810382526021815260200180613f546021913960400191505060405180910390fd5b610f0d858263ffffffff61224b16565b610f2c8487600081518110610f1e57fe5b6020026020010151836124de565b9250506113d5565b60ff891660541415610fde57610f48613ddb565b610f578c61010001518861212d565b9199509750905087610f9a5760405162461bcd60e51b8152600401808060200182810382526021815260200180613f546021913960400191505060405180910390fd5b610faa858263ffffffff61224b16565b610f2c8487600081518110610fbb57fe5b602002602001015188600181518110610fd057fe5b602002602001015184612536565b60ff891660601415610ff357610490836125b7565b60ff8916606114156110f05761101d838660008151811061101057fe5b60200260200101516125bd565b909250905081156110e7578a60c001518b60a00151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001201461109c5760405162461bcd60e51b8152600401808060200182810382526025815260200180613f086025913960400191505060405180910390fd5b8a608001518b60600151146110e25760405162461bcd60e51b8152600401808060200182810382526027815260200180613f2d6027913960400191505060405180910390fd5b6110eb565b5060005b6113d5565b60ff8916607014156112275761111a838660008151811061110d57fe5b60200260200101516125d7565b909250905081156110e75780611175578a608001518b60600151146111705760405162461bcd60e51b8152600401808060200182810382526038815260200180613ed06038913960400191505060405180910390fd5b6110e2565b60808b01516060808d015160408051602080820193909352808201869052815180820383018152930190528151910120146111e15760405162461bcd60e51b8152600401808060200182810382526029815260200180613e606029913960400191505060405180910390fd5b8a60c001518b60a00151146110e25760405162461bcd60e51b8152600401808060200182810382526026815260200180613e896026913960400191505060405180910390fd5b60ff89166072141561124157610490838c60200151612616565b60ff89166073141561125657600091506113d5565b60ff89166074141561126b576110eb8361267c565b60ff89166075141561129557610490838660008151811061128857fe5b6020026020010151612686565b60ff8916607614156112aa57610490836126ab565b60ff8916607714156112bf57610490836126c4565b60ff8916607814156112fe5761049083866000815181106112dc57fe5b6020026020010151876001815181106112f157fe5b602002602001015161270d565b60ff89166079141561135257610490838660008151811061131b57fe5b60200260200101518760018151811061133057fe5b60200260200101518860028151811061134557fe5b6020026020010151612752565b60ff8916607b141561136757610490836127a5565b60ff8916608014156113d057610490838660008151811061138457fe5b60200260200101518760018151811061139957fe5b6020026020010151886002815181106113ae57fe5b6020026020010151896003815181106113c357fe5b60200260200101516127e8565b600091505b80611466578a608001518b60600151146114205760405162461bcd60e51b8152600401808060200182810382526027815260200180613f2d6027913960400191505060405180910390fd5b8a60c001518b60a00151146114665760405162461bcd60e51b8152600401808060200182810382526026815260200180613e896026913960400191505060405180910390fd5b816114c75760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156114bf576114ba83612907565b6114c7565b60c083015183525b6114d0846101bf565b8b511461150e5760405162461bcd60e51b8152600401808060200182810382526022815260200180613e3e6022913960400191505060405180910390fd5b50909998505050505050505050565b606081015160009060ff1661153e57815161153790612911565b905061027f565b606082015160ff166001141561157157602080830151805160408201516060830151929093015161153793919290612935565b606082015160ff166002141561158a57611537826129dd565b600360ff16826060015160ff16101580156115ae57506060820151600c60ff909116105b156115bc5761153782612a43565b606082015160ff16606414156115d45750805161027f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6000806060611622613d7d565b61162a613d7d565b60008061163684612a61565b61164588610100015183612a6b565b9550925090508061169d576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b6116a684612ba3565b9250600088610100015183815181106116bb57fe5b602001015160f81c60f81b60f81c905088610100015183600101815181106116df57fe5b016020015160f81c975060006116f489612c0c565b60408051838152602080850282010190915290995090915081801561173357816020015b611720613ddb565b8152602001906001900390816117185790505b5096506002840193508160ff166000148061175157508160ff166001145b6117a2576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff82166117c7576117c06117bb8a8860000151613169565b61151d565b8652611882565b6117cf613ddb565b6117de8b61010001518661212d565b909650909450905083611838576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b811561185c57808860008151811061184c57fe5b602002602001018190525061186c565b61186c868263ffffffff61226516565b61187e6117bb8b8960000151846131aa565b8752505b60ff82165b818110156119155761189e8b61010001518661212d565b8a518b90859081106118ac57fe5b6020908102919091010152955093508361190d576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611887565b875115611962575060005b8260ff168851038110156119625761195a888260018b5103038151811061194357fe5b60200260200101518861226590919063ffffffff16565b600101611920565b5050505091939550919395565b600061197a836131f3565b158061198c575061198a826131f3565b155b15611999575060006119b8565b825182518082016119b0878263ffffffff6131fe16565b600193505050505b9392505050565b60006119ca836131f3565b15806119dc57506119da826131f3565b155b156119e9575060006119b8565b825182518082026119b0878263ffffffff6131fe16565b6000611a0b836131f3565b1580611a1d5750611a1b826131f3565b155b15611a2a575060006119b8565b825182518082036119b0878263ffffffff6131fe16565b6000611a4c836131f3565b1580611a5e5750611a5c826131f3565b155b15611a6b575060006119b8565b8251825180611a7f576000925050506119b8565b8082046119b0878263ffffffff6131fe16565b6000611a9d836131f3565b1580611aaf5750611aad826131f3565b155b15611abc575060006119b8565b8251825180611ad0576000925050506119b8565b8082056119b0878263ffffffff6131fe16565b6000611aee836131f3565b1580611b005750611afe826131f3565b155b15611b0d575060006119b8565b8251825180611b21576000925050506119b8565b8082066119b0878263ffffffff6131fe16565b6000611b3f836131f3565b1580611b515750611b4f826131f3565b155b15611b5e575060006119b8565b8251825180611b72576000925050506119b8565b8082076119b0878263ffffffff6131fe16565b6000611b90846131f3565b1580611ba25750611ba0836131f3565b155b15611baf57506000611be7565b83518351835180611bc65760009350505050611be7565b6000818385089050611bde898263ffffffff6131fe16565b60019450505050505b949350505050565b6000611bfa846131f3565b1580611c0c5750611c0a836131f3565b155b15611c1957506000611be7565b83518351835180611c305760009350505050611be7565b6000818385099050611bde898263ffffffff6131fe16565b6000611c53836131f3565b1580611c655750611c63826131f3565b155b15611c72575060006119b8565b8251825180820a6119b0878263ffffffff6131fe16565b6000611c94836131f3565b1580611ca65750611ca4826131f3565b155b15611cb3575060006119b8565b825182518082106119b0878263ffffffff6131fe16565b6000611cd5836131f3565b1580611ce75750611ce5826131f3565b155b15611cf4575060006119b8565b825182518082116119b0878263ffffffff6131fe16565b6000611d16836131f3565b1580611d285750611d26826131f3565b155b15611d35575060006119b8565b825182518082126119b0878263ffffffff6131fe16565b6000611d57836131f3565b1580611d695750611d67826131f3565b155b15611d76575060006119b8565b825182518082136119b0878263ffffffff6131fe16565b6000611dba611dad611d9e8461151d565b611da78661151d565b14613214565b859063ffffffff61226516565b5060019392505050565b6000611dcf826131f3565b611de957611de483600063ffffffff6131fe16565b611e00565b81518015611dfd858263ffffffff6131fe16565b50505b50600192915050565b6000611e14836131f3565b1580611e265750611e24826131f3565b155b15611e33575060006119b8565b825182518082166119b0878263ffffffff6131fe16565b6000611e55836131f3565b1580611e675750611e65826131f3565b155b15611e74575060006119b8565b825182518082176119b0878263ffffffff6131fe16565b6000611e96836131f3565b1580611ea85750611ea6826131f3565b155b15611eb5575060006119b8565b825182518082186119b0878263ffffffff6131fe16565b6000611ed7826131f3565b611ee35750600061030b565b81518019611ef7858263ffffffff6131fe16565b506001949350505050565b6000611f0d836131f3565b1580611f1f5750611f1d826131f3565b155b15611f2c575060006119b8565b8251825181811a6119b0878263ffffffff6131fe16565b6000611f4e836131f3565b1580611f605750611f5e826131f3565b155b15611f6d575060006119b8565b8251825181810b6119b0878263ffffffff6131fe16565b6000611e00611f928361151d565b849063ffffffff6131fe16565b6000611e00611fad83613236565b849063ffffffff61226516565b6000611fc5836131f3565b1580611fd75750611fd5826131f3565b155b15611fe4575060006119b8565b82518251604080516020808201859052818301849052825180830384018152606090920190925280519101206119b0878263ffffffff6131fe16565b600192915050565b600061204182608001518361226590919063ffffffff16565b506001919050565b600061204182606001518361226590919063ffffffff16565b60609190910152600190565b6000612079826132bf565b6120855750600061030b565b61208e8261151d565b835250600192915050565b60006120a4836132bf565b6120b0575060006119b8565b6120b9826131f3565b6120c5575060006119b8565b815115611dba576120d58361151d565b84525060019392505050565b60006120416120fe6120f16132cc565b611da7856020015161151d565b839063ffffffff61226516565b6000611e00611fad8360016132ed565b6000611e00838363ffffffff61224b16565b600080612138613ddb565b845184106121585760008461214d6000613378565b925092509250612244565b600080859050600087828151811061216c57fe5b016020015160019092019160f81c90506000816121ae5761218d8984613404565b91955093509050838361219f83613378565b96509650965050505050612244565b60ff8216600114156121c45761219f8984613457565b60ff8216600214156121da5761219f8984613547565b600360ff8316108015906121f15750600c60ff8316105b1561222b5760021982016060612208828c876135ec565b91975095509050858561221a836136aa565b985098509850505050505050612244565b6000806122386000613378565b91985096509450505050505b9250925092565b6122598260400151826137a1565b82604001819052505050565b6122738260200151826137a1565b82602001819052505050565b60006120416120fe61228f6132cc565b611da7856040015161151d565b60006120416120fe8360c0015160016132ed565b60006122bb826132bf565b6122c75750600061030b565b6122d08261151d565b60c084015250600192915050565b60006122f0838363ffffffff61226516565b611e00838363ffffffff61226516565b6000612312848363ffffffff61226516565b612322848463ffffffff61226516565b611dba848363ffffffff61226516565b6000612344858363ffffffff61226516565b612354858463ffffffff61226516565b612364858563ffffffff61226516565b611ef7858363ffffffff61226516565b6000612322848463ffffffff61226516565b6000612398858563ffffffff61226516565b612364858463ffffffff61226516565b60006123b3836131f3565b15806123c557506123c38261381f565b155b156123d2575060006119b8565b6123db8261382e565b60ff168360000151106123f0575060006119b8565b611dba826040015184600001518151811061240757fe5b60200260200101518561226590919063ffffffff16565b60006124298361381f565b158061243b5750612439846131f3565b155b1561244857506000611be7565b6124518361382e565b60ff1684600001511061246657506000611be7565b60408301518451815184918391811061247b57fe5b602002602001018190525061249f612492826136aa565b879063ffffffff61226516565b50600195945050505050565b60006124b68261381f565b6124c25750600061030b565b611e006124ce8361382e565b849060ff1663ffffffff6131fe16565b60006124e9836131f3565b15806124fb57506124f98261381f565b155b15612508575060006119b8565b6125118261382e565b60ff16836000015110612526575060006119b8565b6123f0848363ffffffff61224b16565b60006125418261381f565b15806125535750612551846131f3565b155b1561256057506000611be7565b6125698261382e565b60ff1684600001511061257e57506000611be7565b60408201518451815185918391811061259357fe5b602002602001018190525061249f6125aa826136aa565b879063ffffffff61224b16565b50600190565b60008060016125cb8461151d565b915091505b9250929050565b600080612710836080015111156125f3575060009050806125d0565b6125fc83613859565b61260b575060009050806125d0565b60016125cb8461151d565b60006126206132cc565b6126298361151d565b14156122f0576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b600260e090910152565b6000612691826131f3565b61269d5750600061030b565b505160a09190910152600190565b60006120418260a00151836131fe90919063ffffffff16565b60408051600160f81b60208083019190915260006021830181905260228084018290528451808503909101815260429093019093528151910120612041906120fe9060016132ed565b6000612718836131f3565b612724575060006119b8565b61272d826132bf565b612739575060006119b8565b611dba611dad846000015161274d8561151d565b613169565b600061275d846131f3565b61276957506000611be7565b612772826132bf565b61277e57506000611be7565b611ef761279885600001516127928561151d565b866131aa565b869063ffffffff61226516565b60408051600080825260208201909252606090826127d9565b6127c6613ddb565b8152602001906001900390816127be5790505b509050611e00611fad826136aa565b60006127f3856131f3565b15806128055750612803846131f3565b155b806128165750612814836131f3565b155b806128275750612825826131f3565b155b15612834575060006128fe565b8451845184511580159061284a57508451600114155b1561286b5761286088600063ffffffff6131fe16565b6001925050506128fe565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa1580156128cd573d6000803e3d6000fd5b5050604051601f19015191506128f490508b6001600160a01b03831663ffffffff6131fe16565b6001955050505050505b95945050505050565b600160e090910152565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561298f575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611be7565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214612a32576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161030b9190613970565b6000612a4d613ddb565b612a56836139aa565b90506119b8816129dd565b600060e090910152565b600080612a76613d7d565b612a7e613d7d565b600060e08201819052612a918787613a20565b84529650905080612aab5750600093508492509050612244565b612ab58787613547565b60208501529650905080612ad25750600093508492509050612244565b612adc8787613547565b60408501529650905080612af95750600093508492509050612244565b612b03878761212d565b60608501529650905080612b205750600093508492509050612244565b612b2a878761212d565b60808501529650905080612b475750600093508492509050612244565b612b518787613404565b60a08501529650905080612b6e5750600093508492509050612244565b612b788787613a20565b60c08501529650905080612b955750600093508492509050612244565b506001969495509392505050565b612bab613d7d565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612c245750600290506003613164565b6002831415612c395750600290506003613164565b6003831415612c4e5750600290506003613164565b6004831415612c635750600290506004613164565b6005831415612c785750600290506007613164565b6006831415612c8d5750600290506004613164565b6007831415612ca25750600290506007613164565b6008831415612cb75750600390506004613164565b6009831415612ccc5750600390506004613164565b600a831415612ce15750600290506019613164565b6010831415612cf557506002905080613164565b6011831415612d0957506002905080613164565b6012831415612d1d57506002905080613164565b6013831415612d3157506002905080613164565b6014831415612d4557506002905080613164565b6015831415612d5957506001905080613164565b6016831415612d6d57506002905080613164565b6017831415612d8157506002905080613164565b6018831415612d9557506002905080613164565b6019831415612da957506001905080613164565b601a831415612dbe5750600290506004613164565b601b831415612dd35750600290506007613164565b6020831415612de85750600190506007613164565b6021831415612dfd5750600190506003613164565b6022831415612e125750600290506008613164565b6030831415612e2657506001905080613164565b6031831415612e3b5750600090506001613164565b6032831415612e505750600090506001613164565b6033831415612e655750600190506002613164565b6034831415612e7a5750600190506004613164565b6035831415612e8f5750600290506004613164565b6036831415612ea45750600090506002613164565b6037831415612eb95750600090506001613164565b6038831415612ecd57506001905080613164565b6039831415612ee25750600090506001613164565b603a831415612ef75750600090506002613164565b603b831415612f0c5750600090506001613164565b603c831415612f215750600090506001613164565b603d831415612f3557506001905080613164565b6040831415612f4957506001905080613164565b6041831415612f5e5750600290506001613164565b6042831415612f735750600390506001613164565b6043831415612f885750600290506001613164565b6044831415612f9d5750600390506001613164565b6050831415612fb157506002905080613164565b6051831415612fc65750600390506028613164565b6052831415612fdb5750600190506002613164565b6053831415612ff05750600190506003613164565b60548314156130055750600290506029613164565b606083141561301a5750600090506064613164565b606183141561302f5750600190506064613164565b60708314156130445750600190506064613164565b60728314156130595750600090506028613164565b607383141561306e5750600090506005613164565b6074831415613083575060009050600a613164565b60758314156130985750600190506000613164565b60768314156130ad5750600090506001613164565b60778314156130c25750600090506019613164565b60788314156130d75750600290506019613164565b60798314156130ec5750600390506019613164565b607b831415613101575060009050600a613164565b6080831415613117575060049050614e20613164565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b613171613ddb565b6119b86040518060a001604052808560ff1681526020018481526020016000151581526020016000801b81526020016000815250613a74565b6131b2613ddb565b611be76040518060a001604052808660ff1681526020018581526020016001151581526020016131e18561151d565b81526020018460800151815250613a74565b6060015160ff161590565b612273826020015161320f83613378565b6137a1565b61321c613ddb565b811561322c576115376001613378565b6115376000613378565b61323e613ddb565b816060015160ff16600214156132855760405162461bcd60e51b8152600401808060200182810382526021815260200180613eaf6021913960400191505060405180910390fd5b606082015160ff1661329b576115376000613378565b816060015160ff16600114156132b5576115376001613378565b6115376003613378565b6060015160ff1660011490565b604080516000808252602082019092526132e7816001613adb565b91505090565b6132f5613ddb565b6040805160a080820183528582528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191613361565b61334e613ddb565b8152602001906001900390816133465790505b508152606460208201526040019290925250919050565b613380613ddb565b6040805160a0808201835284825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916133ec565b6133d9613ddb565b8152602001906001900390816133d15790505b50815260006020820152600160409091015292915050565b600080600080855190508481108061341e57506020858203105b15613433575060009250839150829050612244565b600160208601613449888863ffffffff613afa16565b935093509350509250925092565b600080613462613ddb565b6000849050600086828151811061347557fe5b602001015160f81c60f81b60f81c90508180600101925050600087838151811061349b57fe5b016020015160019093019260f81c90506134b3613ddb565b8260ff16600114156134f45760006134cb8a8661212d565b90965092509050806134f2576000896134e2613b16565b9750975097505050505050612244565b505b60006135068a8663ffffffff613afa16565b90506020850194508360ff1660011415613528576001856134e28584866131aa565b6001856135358584613169565b97509750975050505050509250925092565b600080613552613ddb565b61355a613ddb565b855160009081908781108061357157506040888203105b15613589576000888596509650965050505050612244565b600061359b8a8a63ffffffff613afa16565b90506020890198506135ad8a8a613404565b909a509450925082156135d8576135c48185610284565b600198508997509550612244945050505050565b600089869750975097505050505050612244565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561363757816020015b613624613ddb565b81526020019060019003908161361c5790505b50905060005b8960ff168160ff16101561369457613655898561212d565b8451859060ff861690811061366657fe5b6020908102919091010152945092508261368c57506000955086945092506136a1915050565b60010161363d565b5060019550919350909150505b93509350939050565b6136b2613ddb565b6136bc8251613ba1565b61370d576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156137445783818151811061372757fe5b602002602001015160800151820191508080600101915050613712565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b6137a9613ddb565b6040805160028082526060828101909352816020015b6137c7613ddb565b8152602001906001900390816137bf57905050905082816000815181106137ea57fe5b6020026020010181905250838160018151811061380357fe5b6020026020010181905250611be761381a826136aa565b6139aa565b600061030b8260600151613ba8565b600061383d8260600151613ba8565b15613851575060608101516002190161027f565b50600161027f565b606081015160009060ff166138705750600161027f565b606082015160ff16600114156138885750600061027f565b606082015160ff16600214156138dc576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff161015801561390057506060820151600c60ff909116105b156139585760408201515160005b8181101561394d576139368460400151828151811061392957fe5b6020026020010151613859565b6139455760009250505061027f565b60010161390e565b50600191505061027f565b606082015160ff16606414156115d45750600061027f565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b6139b2613ddb565b6139bb8261381f565b613a01576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060613a108360400151613bc6565b90506119b8818460800151613c9e565b60008060008060008651905085811080613a3c57506020868203105b15613a505750600093508492509050612244565b613a60878763ffffffff613afa16565b600195506020870194509250612244915050565b613a7c613ddb565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190613ac3565b613ab0613ddb565b815260200190600190039081613aa85790505b50815260016020820181905260409091015292915050565b6000613ae5613ddb565b613aef8484613c9e565b9050611be7816129dd565b60008160200183511015613b0d57600080fd5b50016020015190565b613b1e613ddb565b6040805160a080820183526000808352835191820184528082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191613b8a565b613b77613ddb565b815260200190600190039081613b6f5790505b508152600360208201526001604090910152905090565b6008101590565b6000600c60ff831610801561030b575050600360ff91909116101590565b6060600882511115613c16576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613c43578160200160208202803883390190505b50805190915060005b81811015613c95576000613c72868381518110613c6557fe5b602002602001015161151d565b905080848381518110613c8157fe5b602090810291909101015250600101613c4c565b50909392505050565b613ca6613ddb565b6000613cb184613cbd565b9050611be78184610284565b6000600882511115613d0d576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613d51578181015183820152602001613d39565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040805161010081019091526000815260208101613d99613ddb565b8152602001613da6613ddb565b8152602001613db3613ddb565b8152602001613dc0613ddb565b81526000602082018190526040820181905260609091015290565b6040518060a0016040528060008152602001613df5613e0f565b815260606020820181905260006040830181905291015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a7231582001ac02fda2077c80b026e190533d4d38944b636c3c0d21a6df99bb87429eb02a64736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582087e25d1b3581edfbcc93fab1b50976d754d7ebc01e825ad79f9dee7c91ee0bf164736f6c63430005110032"

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
