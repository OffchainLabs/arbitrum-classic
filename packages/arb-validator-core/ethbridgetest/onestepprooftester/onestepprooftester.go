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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158207c82241594b68ba85e193257ef73fe80d60621fa438add08dffe3e08258c8b7b64736f6c63430005110032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b82ceb7f91e648389b8f820edef0ec179e35160c2ca2f29f171532674156cf4764736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b50614042806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633c41485d14610030575b600080fd5b61011d600480360361014081101561004757600080fd5b813591602081013591604082013591606081013515159160808201359160a08101359160c08201359160e08101359167ffffffffffffffff61010083013516919081019061014081016101208201356401000000008111156100a857600080fd5b8201836020820111156100ba57600080fd5b803590602001918460018302840111640100000000831117156100dc57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061012f945050505050565b60408051918252519081900360200190f35b600061014b6101468c8c8c8c8c8c8c8c8c8c61015a565b6101bf565b9b9a5050505050505050505050565b610162613e25565b61014b6040518061012001604052808d81526020016101818d8d610284565b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610337565b600060028260e0015114156101d65750600061027f565b60018260e0015114156101eb5750600161027f565b815160208301516101fb90611543565b6102088460400151611543565b6102158560600151611543565b6102228660800151611543565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090505b919050565b61028c613e83565b6040805160a08101825284815281516060810183526000808252602082810182905284518281528082018652939490850193908301916102e2565b6102cf613e83565b8152602001906001900390816102c75790505b5090528152604080516000808252602082810190935291909201919061031e565b61030b613e83565b8152602001906001900390816103035790505b5081526002602082015260400183905290505b92915050565b61033f613e25565b600080600080606061034f613e25565b610357613e25565b61036089611648565b60e08f0151959c50939a509297509095509350915060019060009067ffffffffffffffff1688146103cf576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8a6040015180156103e3575060ff89166072145b806103ff57508a604001511580156103ff575060ff8916607214155b610450576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a080840180518a900390528401518811156104775760001960a0840152600091506113fb565b60ff8916600114156104bd576104b6838660008151811061049457fe5b6020026020010151876001815181106104a957fe5b60200260200101516119a2565b91506113fb565b60ff8916600214156104fc576104b683866000815181106104da57fe5b6020026020010151876001815181106104ef57fe5b60200260200101516119f2565b60ff89166003141561053b576104b6838660008151811061051957fe5b60200260200101518760018151811061052e57fe5b6020026020010151611a33565b60ff89166004141561057a576104b6838660008151811061055857fe5b60200260200101518760018151811061056d57fe5b6020026020010151611a74565b60ff8916600514156105b9576104b6838660008151811061059757fe5b6020026020010151876001815181106105ac57fe5b6020026020010151611ac5565b60ff8916600614156105f8576104b683866000815181106105d657fe5b6020026020010151876001815181106105eb57fe5b6020026020010151611b16565b60ff891660071415610637576104b6838660008151811061061557fe5b60200260200101518760018151811061062a57fe5b6020026020010151611b67565b60ff89166008141561068b576104b6838660008151811061065457fe5b60200260200101518760018151811061066957fe5b60200260200101518860028151811061067e57fe5b6020026020010151611bb8565b60ff8916600914156106df576104b683866000815181106106a857fe5b6020026020010151876001815181106106bd57fe5b6020026020010151886002815181106106d257fe5b6020026020010151611c22565b60ff8916600a141561071e576104b683866000815181106106fc57fe5b60200260200101518760018151811061071157fe5b6020026020010151611c7b565b60ff89166010141561075d576104b6838660008151811061073b57fe5b60200260200101518760018151811061075057fe5b6020026020010151611cbc565b60ff89166011141561079c576104b6838660008151811061077a57fe5b60200260200101518760018151811061078f57fe5b6020026020010151611cfd565b60ff8916601214156107db576104b683866000815181106107b957fe5b6020026020010151876001815181106107ce57fe5b6020026020010151611d3e565b60ff89166013141561081a576104b683866000815181106107f857fe5b60200260200101518760018151811061080d57fe5b6020026020010151611d7f565b60ff891660141415610859576104b6838660008151811061083757fe5b60200260200101518760018151811061084c57fe5b6020026020010151611dc0565b60ff891660151415610883576104b6838660008151811061087657fe5b6020026020010151611df7565b60ff8916601614156108c2576104b683866000815181106108a057fe5b6020026020010151876001815181106108b557fe5b6020026020010151611e3c565b60ff891660171415610901576104b683866000815181106108df57fe5b6020026020010151876001815181106108f457fe5b6020026020010151611e7d565b60ff891660181415610940576104b6838660008151811061091e57fe5b60200260200101518760018151811061093357fe5b6020026020010151611ebe565b60ff89166019141561096a576104b6838660008151811061095d57fe5b6020026020010151611eff565b60ff8916601a14156109a9576104b6838660008151811061098757fe5b60200260200101518760018151811061099c57fe5b6020026020010151611f35565b60ff8916601b14156109e8576104b683866000815181106109c657fe5b6020026020010151876001815181106109db57fe5b6020026020010151611f76565b60ff891660201415610a12576104b68386600081518110610a0557fe5b6020026020010151611fb7565b60ff891660211415610a3c576104b68386600081518110610a2f57fe5b6020026020010151611fd2565b60ff891660221415610a7b576104b68386600081518110610a5957fe5b602002602001015187600181518110610a6e57fe5b6020026020010151611fed565b60ff891660301415610aa5576104b68386600081518110610a9857fe5b6020026020010151612053565b60ff891660311415610aba576104b68361205b565b60ff891660321415610acf576104b68361207c565b60ff891660331415610af9576104b68386600081518110610aec57fe5b6020026020010151612095565b60ff891660341415610b23576104b68386600081518110610b1657fe5b60200260200101516120a1565b60ff891660351415610b62576104b68386600081518110610b4057fe5b602002602001015187600181518110610b5557fe5b60200260200101516120cc565b60ff891660361415610b77576104b683612114565b60ff891660371415610b91576104b683856000015161213e565b60ff891660381415610bbb576104b68386600081518110610bae57fe5b602002602001015161214e565b60ff891660391415610c4757610bcf613e83565b610bde8c610100015188612160565b9199509750905087610c215760405162461bcd60e51b8152600401808060200182810382526021815260200180613fed6021913960400191505060405180910390fd5b610c31858263ffffffff61227e16565b610c41848263ffffffff61229816565b506113fb565b60ff8916603a1415610c5c576104b6836122b2565b60ff8916603b1415610c7157600191506113fb565b60ff8916603c1415610c86576104b6836122cf565b60ff8916603d1415610cb0576104b68386600081518110610ca357fe5b60200260200101516122e3565b60ff891660401415610cda576104b68386600081518110610ccd57fe5b6020026020010151612311565b60ff891660411415610d19576104b68386600081518110610cf757fe5b602002602001015187600181518110610d0c57fe5b6020026020010151612333565b60ff891660421415610d6d576104b68386600081518110610d3657fe5b602002602001015187600181518110610d4b57fe5b602002602001015188600281518110610d6057fe5b6020026020010151612365565b60ff891660431415610dac576104b68386600081518110610d8a57fe5b602002602001015187600181518110610d9f57fe5b60200260200101516123a7565b60ff891660441415610e00576104b68386600081518110610dc957fe5b602002602001015187600181518110610dde57fe5b602002602001015188600281518110610df357fe5b60200260200101516123b9565b60ff891660501415610e3f576104b68386600081518110610e1d57fe5b602002602001015187600181518110610e3257fe5b60200260200101516123db565b60ff891660511415610e93576104b68386600081518110610e5c57fe5b602002602001015187600181518110610e7157fe5b602002602001015188600281518110610e8657fe5b6020026020010151612451565b60ff891660521415610ebd576104b68386600081518110610eb057fe5b60200260200101516124de565b60ff891660531415610f5a57610ed1613e83565b610ee08c610100015188612160565b9199509750905087610f235760405162461bcd60e51b8152600401808060200182810382526021815260200180613fed6021913960400191505060405180910390fd5b610f33858263ffffffff61227e16565b610f528487600081518110610f4457fe5b602002602001015183612511565b9250506113fb565b60ff89166054141561100457610f6e613e83565b610f7d8c610100015188612160565b9199509750905087610fc05760405162461bcd60e51b8152600401808060200182810382526021815260200180613fed6021913960400191505060405180910390fd5b610fd0858263ffffffff61227e16565b610f528487600081518110610fe157fe5b602002602001015188600181518110610ff657fe5b602002602001015184612569565b60ff891660601415611019576104b6836125ea565b60ff89166061141561111657611043838660008151811061103657fe5b60200260200101516125f0565b9092509050811561110d578a60c001518b60a0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146110c25760405162461bcd60e51b8152600401808060200182810382526025815260200180613fa16025913960400191505060405180910390fd5b8a608001518b60600151146111085760405162461bcd60e51b8152600401808060200182810382526027815260200180613fc66027913960400191505060405180910390fd5b611111565b5060005b6113fb565b60ff89166070141561124d57611140838660008151811061113357fe5b602002602001015161260a565b9092509050811561110d578061119b578a608001518b60600151146111965760405162461bcd60e51b8152600401808060200182810382526038815260200180613f696038913960400191505060405180910390fd5b611108565b60808b01516060808d015160408051602080820193909352808201869052815180820383018152930190528151910120146112075760405162461bcd60e51b8152600401808060200182810382526029815260200180613ef96029913960400191505060405180910390fd5b8a60c001518b60a00151146111085760405162461bcd60e51b8152600401808060200182810382526026815260200180613f226026913960400191505060405180910390fd5b60ff891660721415611267576104b6838c60200151612649565b60ff89166073141561127c57600091506113fb565b60ff89166074141561129157611111836126af565b60ff8916607514156112bb576104b683866000815181106112ae57fe5b60200260200101516126b9565b60ff8916607614156112d0576104b6836126de565b60ff8916607714156112e5576104b6836126f7565b60ff891660781415611324576104b6838660008151811061130257fe5b60200260200101518760018151811061131757fe5b6020026020010151612740565b60ff891660791415611378576104b6838660008151811061134157fe5b60200260200101518760018151811061135657fe5b60200260200101518860028151811061136b57fe5b6020026020010151612785565b60ff8916607b141561138d576104b6836127d8565b60ff8916608014156113f6576104b683866000815181106113aa57fe5b6020026020010151876001815181106113bf57fe5b6020026020010151886002815181106113d457fe5b6020026020010151896003815181106113e957fe5b602002602001015161281b565b600091505b8061148c578a608001518b60600151146114465760405162461bcd60e51b8152600401808060200182810382526027815260200180613fc66027913960400191505060405180910390fd5b8a60c001518b60a001511461148c5760405162461bcd60e51b8152600401808060200182810382526026815260200180613f226026913960400191505060405180910390fd5b816114ed5760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156114e5576114e08361293a565b6114ed565b60c083015183525b6114f6846101bf565b8b51146115345760405162461bcd60e51b8152600401808060200182810382526022815260200180613ed76022913960400191505060405180910390fd5b50909998505050505050505050565b606081015160009060ff1661156457815161155d90612944565b905061027f565b606082015160ff16600114156115815761155d8260200151612968565b606082015160ff16600214156115a2578151608083015161155d9190612a54565b600360ff16826060015160ff16101580156115c657506060820151600c60ff909116105b156115ef576115d3613e83565b6115dc83612a8e565b90506115e781611543565b91505061027f565b606082015160ff16606414156116075750805161027f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6000806060611655613e25565b61165d613e25565b60008061166984612bc0565b61167888610100015183612bca565b955092509050806116d0576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b6116d984612d02565b9250600088610100015183815181106116ee57fe5b602001015160f81c60f81b60f81c9050886101000151836001018151811061171257fe5b016020015160f81c9750600061172789612d6b565b60408051838152602080850282010190915290995090915081801561176657816020015b611753613e83565b81526020019060019003908161174b5790505b5096506002840193508160ff166000148061178457508160ff166001145b6117d5576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff82166117fa576117f36117ee8a88600001516132c8565b611543565b86526118b5565b611802613e83565b6118118b610100015186612160565b90965090945090508361186b576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b811561188f57808860008151811061187f57fe5b602002602001018190525061189f565b61189f868263ffffffff61229816565b6118b16117ee8b896000015184613323565b8752505b60ff82165b81811015611948576118d18b610100015186612160565b8a518b90859081106118df57fe5b60209081029190910101529550935083611940576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b6001016118ba565b875115611995575060005b8260ff168851038110156119955761198d888260018b5103038151811061197657fe5b60200260200101518861229890919063ffffffff16565b600101611953565b5050505091939550919395565b60006119ad8361339e565b15806119bf57506119bd8261339e565b155b156119cc575060006119eb565b825182518082016119e3878263ffffffff6133a916565b600193505050505b9392505050565b60006119fd8361339e565b1580611a0f5750611a0d8261339e565b155b15611a1c575060006119eb565b825182518082026119e3878263ffffffff6133a916565b6000611a3e8361339e565b1580611a505750611a4e8261339e565b155b15611a5d575060006119eb565b825182518082036119e3878263ffffffff6133a916565b6000611a7f8361339e565b1580611a915750611a8f8261339e565b155b15611a9e575060006119eb565b8251825180611ab2576000925050506119eb565b8082046119e3878263ffffffff6133a916565b6000611ad08361339e565b1580611ae25750611ae08261339e565b155b15611aef575060006119eb565b8251825180611b03576000925050506119eb565b8082056119e3878263ffffffff6133a916565b6000611b218361339e565b1580611b335750611b318261339e565b155b15611b40575060006119eb565b8251825180611b54576000925050506119eb565b8082066119e3878263ffffffff6133a916565b6000611b728361339e565b1580611b845750611b828261339e565b155b15611b91575060006119eb565b8251825180611ba5576000925050506119eb565b8082076119e3878263ffffffff6133a916565b6000611bc38461339e565b1580611bd55750611bd38361339e565b155b15611be257506000611c1a565b83518351835180611bf95760009350505050611c1a565b6000818385089050611c11898263ffffffff6133a916565b60019450505050505b949350505050565b6000611c2d8461339e565b1580611c3f5750611c3d8361339e565b155b15611c4c57506000611c1a565b83518351835180611c635760009350505050611c1a565b6000818385099050611c11898263ffffffff6133a916565b6000611c868361339e565b1580611c985750611c968261339e565b155b15611ca5575060006119eb565b8251825180820a6119e3878263ffffffff6133a916565b6000611cc78361339e565b1580611cd95750611cd78261339e565b155b15611ce6575060006119eb565b825182518082106119e3878263ffffffff6133a916565b6000611d088361339e565b1580611d1a5750611d188261339e565b155b15611d27575060006119eb565b825182518082116119e3878263ffffffff6133a916565b6000611d498361339e565b1580611d5b5750611d598261339e565b155b15611d68575060006119eb565b825182518082126119e3878263ffffffff6133a916565b6000611d8a8361339e565b1580611d9c5750611d9a8261339e565b155b15611da9575060006119eb565b825182518082136119e3878263ffffffff6133a916565b6000611ded611de0611dd184611543565b611dda86611543565b146133bf565b859063ffffffff61229816565b5060019392505050565b6000611e028261339e565b611e1c57611e1783600063ffffffff6133a916565b611e33565b81518015611e30858263ffffffff6133a916565b50505b50600192915050565b6000611e478361339e565b1580611e595750611e578261339e565b155b15611e66575060006119eb565b825182518082166119e3878263ffffffff6133a916565b6000611e888361339e565b1580611e9a5750611e988261339e565b155b15611ea7575060006119eb565b825182518082176119e3878263ffffffff6133a916565b6000611ec98361339e565b1580611edb5750611ed98261339e565b155b15611ee8575060006119eb565b825182518082186119e3878263ffffffff6133a916565b6000611f0a8261339e565b611f1657506000610331565b81518019611f2a858263ffffffff6133a916565b506001949350505050565b6000611f408361339e565b1580611f525750611f508261339e565b155b15611f5f575060006119eb565b8251825181811a6119e3878263ffffffff6133a916565b6000611f818361339e565b1580611f935750611f918261339e565b155b15611fa0575060006119eb565b8251825181810b6119e3878263ffffffff6133a916565b6000611e33611fc583611543565b849063ffffffff6133a916565b6000611e33611fe0836133e1565b849063ffffffff61229816565b6000611ff88361339e565b158061200a57506120088261339e565b155b15612017575060006119eb565b82518251604080516020808201859052818301849052825180830384018152606090920190925280519101206119e3878263ffffffff6133a916565b600192915050565b600061207482608001518361229890919063ffffffff16565b506001919050565b600061207482606001518361229890919063ffffffff16565b60609190910152600190565b60006120ac8261346a565b6120b857506000610331565b6120c182611543565b835250600192915050565b60006120d78361346a565b6120e3575060006119eb565b6120ec8261339e565b6120f8575060006119eb565b815115611ded5761210883611543565b84525060019392505050565b6000612074612131612124613477565b611dda8560200151611543565b839063ffffffff61229816565b6000611e33611fe0836001613498565b6000611e33838363ffffffff61227e16565b60008061216b613e83565b8451841061218b576000846121806000613549565b925092509250612277565b600080859050600087828151811061219f57fe5b016020015160019092019160f81c90506000816121e1576121c089846135fb565b9195509350905083836121d283613549565b96509650965050505050612277565b60ff8216600114156121f7576121d2898461364e565b60ff82166002141561220d576121d2898461373e565b600360ff8316108015906122245750600c60ff8316105b1561225e576002198201606061223b828c876137e3565b91975095509050858561224d836138a1565b985098509850505050505050612277565b60008061226b6000613549565b91985096509450505050505b9250925092565b61228c8260400151826139b8565b82604001819052505050565b6122a68260200151826139b8565b82602001819052505050565b60006120746121316122c2613477565b611dda8560400151611543565b60006120746121318360c001516001613498565b60006122ee8261346a565b6122fa57506000610331565b61230382611543565b60c084015250600192915050565b6000612323838363ffffffff61229816565b611e33838363ffffffff61229816565b6000612345848363ffffffff61229816565b612355848463ffffffff61229816565b611ded848363ffffffff61229816565b6000612377858363ffffffff61229816565b612387858463ffffffff61229816565b612397858563ffffffff61229816565b611f2a858363ffffffff61229816565b6000612355848463ffffffff61229816565b60006123cb858563ffffffff61229816565b612397858463ffffffff61229816565b60006123e68361339e565b15806123f857506123f682613a36565b155b15612405575060006119eb565b61240e82613a45565b60ff16836000015110612423575060006119eb565b611ded826040015184600001518151811061243a57fe5b60200260200101518561229890919063ffffffff16565b600061245c83613a36565b158061246e575061246c8461339e565b155b1561247b57506000611c1a565b61248483613a45565b60ff1684600001511061249957506000611c1a565b6040830151845181518491839181106124ae57fe5b60200260200101819052506124d26124c5826138a1565b879063ffffffff61229816565b50600195945050505050565b60006124e982613a36565b6124f557506000610331565b611e3361250183613a45565b849060ff1663ffffffff6133a916565b600061251c8361339e565b158061252e575061252c82613a36565b155b1561253b575060006119eb565b61254482613a45565b60ff16836000015110612559575060006119eb565b612423848363ffffffff61227e16565b600061257482613a36565b158061258657506125848461339e565b155b1561259357506000611c1a565b61259c82613a45565b60ff168460000151106125b157506000611c1a565b6040820151845181518591839181106125c657fe5b60200260200101819052506124d26125dd826138a1565b879063ffffffff61227e16565b50600190565b60008060016125fe84611543565b915091505b9250929050565b6000806127108360800151111561262657506000905080612603565b61262f83613a70565b61263e57506000905080612603565b60016125fe84611543565b6000612653613477565b61265c83611543565b1415612323576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b600260e090910152565b60006126c48261339e565b6126d057506000610331565b505160a09190910152600190565b60006120748260a00151836133a990919063ffffffff16565b60408051600160f81b6020808301919091526000602183018190526022808401829052845180850390910181526042909301909352815191012061207490612131906001613498565b600061274b8361339e565b612757575060006119eb565b6127608261346a565b61276c575060006119eb565b611ded611de0846000015161278085611543565b6132c8565b60006127908461339e565b61279c57506000611c1a565b6127a58261346a565b6127b157506000611c1a565b611f2a6127cb85600001516127c585611543565b86613323565b869063ffffffff61229816565b604080516000808252602082019092526060908261280c565b6127f9613e83565b8152602001906001900390816127f15790505b509050611e33611fe0826138a1565b60006128268561339e565b158061283857506128368461339e565b155b8061284957506128478361339e565b155b8061285a57506128588261339e565b155b1561286757506000612931565b8451845184511580159061287d57508451600114155b1561289e5761289388600063ffffffff6133a916565b600192505050612931565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612900573d6000803e3d6000fd5b5050604051601f190151915061292790508b6001600160a01b03831663ffffffff6133a916565b6001955050505050505b95945050505050565b600160e090910152565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061297957fe5b6040820151516129d35750805160208083015160408051600160f81b8185015260f89490941b6001600160f81b0319166021850152602280850192909252805180850390920182526042909301909252815191012061027f565b600182600001516129fb84604001516000815181106129ee57fe5b6020026020010151611543565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b612a96613e83565b612a9f82613a36565b612ae5576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b60088260400151511115612b37576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060826040015151604051908082528060200260200182016040528015612b68578160200160208202803883390190505b50805190915060005b81811015612bb1576000612b8e866040015183815181106129ee57fe5b905080848381518110612b9d57fe5b602090810291909101015250600101612b71565b50611c1a828560800151613b87565b600060e090910152565b600080612bd5613e25565b612bdd613e25565b600060e08201819052612bf08787613ba6565b84529650905080612c0a5750600093508492509050612277565b612c14878761373e565b60208501529650905080612c315750600093508492509050612277565b612c3b878761373e565b60408501529650905080612c585750600093508492509050612277565b612c628787612160565b60608501529650905080612c7f5750600093508492509050612277565b612c898787612160565b60808501529650905080612ca65750600093508492509050612277565b612cb087876135fb565b60a08501529650905080612ccd5750600093508492509050612277565b612cd78787613ba6565b60c08501529650905080612cf45750600093508492509050612277565b506001969495509392505050565b612d0a613e25565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612d8357506002905060036132c3565b6002831415612d9857506002905060036132c3565b6003831415612dad57506002905060036132c3565b6004831415612dc257506002905060046132c3565b6005831415612dd757506002905060076132c3565b6006831415612dec57506002905060046132c3565b6007831415612e0157506002905060076132c3565b6008831415612e1657506003905060046132c3565b6009831415612e2b57506003905060046132c3565b600a831415612e4057506002905060196132c3565b6010831415612e54575060029050806132c3565b6011831415612e68575060029050806132c3565b6012831415612e7c575060029050806132c3565b6013831415612e90575060029050806132c3565b6014831415612ea4575060029050806132c3565b6015831415612eb8575060019050806132c3565b6016831415612ecc575060029050806132c3565b6017831415612ee0575060029050806132c3565b6018831415612ef4575060029050806132c3565b6019831415612f08575060019050806132c3565b601a831415612f1d57506002905060046132c3565b601b831415612f3257506002905060076132c3565b6020831415612f4757506001905060076132c3565b6021831415612f5c57506001905060036132c3565b6022831415612f7157506002905060086132c3565b6030831415612f85575060019050806132c3565b6031831415612f9a57506000905060016132c3565b6032831415612faf57506000905060016132c3565b6033831415612fc457506001905060026132c3565b6034831415612fd957506001905060046132c3565b6035831415612fee57506002905060046132c3565b603683141561300357506000905060026132c3565b603783141561301857506000905060016132c3565b603883141561302c575060019050806132c3565b603983141561304157506000905060016132c3565b603a83141561305657506000905060026132c3565b603b83141561306b57506000905060016132c3565b603c83141561308057506000905060016132c3565b603d831415613094575060019050806132c3565b60408314156130a8575060019050806132c3565b60418314156130bd57506002905060016132c3565b60428314156130d257506003905060016132c3565b60438314156130e757506002905060016132c3565b60448314156130fc57506003905060016132c3565b6050831415613110575060029050806132c3565b605183141561312557506003905060286132c3565b605283141561313a57506001905060026132c3565b605383141561314f57506001905060036132c3565b605483141561316457506002905060296132c3565b606083141561317957506000905060646132c3565b606183141561318e57506001905060646132c3565b60708314156131a357506001905060646132c3565b60728314156131b857506000905060286132c3565b60738314156131cd57506000905060056132c3565b60748314156131e2575060009050600a6132c3565b60758314156131f757506001905060006132c3565b607683141561320c57506000905060016132c3565b607783141561322157506000905060196132c3565b607883141561323657506002905060196132c3565b607983141561324b57506003905060196132c3565b607b831415613260575060009050600a6132c3565b6080831415613276575060049050614e206132c3565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b6132d0613e83565b6040805160608101825260ff8516815260208082018590528251600080825291810184526119eb9383019161331b565b613308613e83565b8152602001906001900390816133005790505b509052613bfa565b61332b613e83565b604080516001808252818301909252606091816020015b61334a613e83565b815260200190600190039081613342579050509050828160008151811061336d57fe5b602002602001018190525061293160405180606001604052808760ff16815260200186815260200183815250613bfa565b6060015160ff161590565b6122a682602001516133ba83613549565b6139b8565b6133c7613e83565b81156133d75761155d6001613549565b61155d6000613549565b6133e9613e83565b816060015160ff16600214156134305760405162461bcd60e51b8152600401808060200182810382526021815260200180613f486021913960400191505060405180910390fd5b606082015160ff166134465761155d6000613549565b816060015160ff16600114156134605761155d6001613549565b61155d6003613549565b6060015160ff1660011490565b60408051600080825260208201909252613492816001613c61565b91505090565b6134a0613e83565b6040805160a08101825284815281516060810183526000808252602082810182905284518281528082018652939490850193908301916134f6565b6134e3613e83565b8152602001906001900390816134db5790505b50905281526040805160008082526020828101909352919092019190613532565b61351f613e83565b8152602001906001900390816135175790505b508152606460208201526040019290925250919050565b613551613e83565b6040805160a08101825283815281516060810183526000808252602082810182905284518281528082018652939490850193908301916135a7565b613594613e83565b81526020019060019003908161358c5790505b509052815260408051600080825260208281019093529190920191906135e3565b6135d0613e83565b8152602001906001900390816135c85790505b50815260006020820152600160409091015292915050565b600080600080855190508481108061361557506020858203105b1561362a575060009250839150829050612277565b600160208601613640888863ffffffff613c7016565b935093509350509250925092565b600080613659613e83565b6000849050600086828151811061366c57fe5b602001015160f81c60f81b60f81c90508180600101925050600087838151811061369257fe5b016020015160019093019260f81c90506136aa613e83565b8260ff16600114156136eb5760006136c28a86612160565b90965092509050806136e9576000896136d9613c8c565b9750975097505050505050612277565b505b60006136fd8a8663ffffffff613c7016565b90506020850194508360ff166001141561371f576001856136d9858486613323565b60018561372c85846132c8565b97509750975050505050509250925092565b600080613749613e83565b613751613e83565b855160009081908781108061376857506040888203105b15613780576000888596509650965050505050612277565b60006137928a8a63ffffffff613c7016565b90506020890198506137a48a8a6135fb565b909a509450925082156137cf576137bb8185610284565b600198508997509550612277945050505050565b600089869750975097505050505050612277565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561382e57816020015b61381b613e83565b8152602001906001900390816138135790505b50905060005b8960ff168160ff16101561388b5761384c8985612160565b8451859060ff861690811061385d57fe5b602090810291909101015294509250826138835750600095508694509250613898915050565b600101613834565b5060019550919350909150505b93509350939050565b6138a9613e83565b6138b38251613d40565b613904576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b835181101561393b5783818151811061391e57fe5b602002602001015160800151820191508080600101915050613909565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190613995565b613982613e83565b81526020019060019003908161397a5790505b509052815260208101859052935160030160ff1660408501526060909301525090565b6139c0613e83565b6040805160028082526060828101909352816020015b6139de613e83565b8152602001906001900390816139d65790505090508281600081518110613a0157fe5b60200260200101819052508381600181518110613a1a57fe5b6020026020010181905250611c1a613a31826138a1565b612a8e565b60006103318260600151613d47565b6000613a548260600151613d47565b15613a68575060608101516002190161027f565b50600161027f565b606081015160009060ff16613a875750600161027f565b606082015160ff1660011415613a9f5750600061027f565b606082015160ff1660021415613af3576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff1610158015613b1757506060820151600c60ff909116105b15613b6f5760408201515160005b81811015613b6457613b4d84604001518281518110613b4057fe5b6020026020010151613a70565b613b5c5760009250505061027f565b600101613b25565b50600191505061027f565b606082015160ff16606414156116075750600061027f565b613b8f613e83565b6000613b9a84613d65565b9050611c1a8184610284565b60008060008060008651905085811080613bc257506020868203105b15613bd65750600093508492509050612277565b613be6878763ffffffff613c7016565b600195506020870194509250612277915050565b613c02613e83565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190613c49565b613c36613e83565b815260200190600190039081613c2e5790505b50815260016020820181905260409091015292915050565b60006119eb6117ee8484613b87565b60008160200183511015613c8357600080fd5b50016020015190565b613c94613e83565b6040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190613ced565b613cda613e83565b815260200190600190039081613cd25790505b50905281526040805160008082526020828101909352919092019190613d29565b613d16613e83565b815260200190600190039081613d0e5790505b508152600360208201526001604090910152905090565b6008101590565b6000600c60ff8316108015610331575050600360ff91909116101590565b6000600882511115613db5576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613df9578181015183820152602001613de1565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040805161010081019091526000815260208101613e41613e83565b8152602001613e4e613e83565b8152602001613e5b613e83565b8152602001613e68613e83565b81526000602082018190526040820181905260609091015290565b6040518060a0016040528060008152602001613e9d613eb7565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820d361c7afc7d5f11d342e1444c7fb7649b514a4c28c539177b7cf362d22dbfec264736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820142891281d8f03ec873930437a8ca5e0c61de65eedd6218afc3087c76cc8fbbd64736f6c63430005110032"

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
