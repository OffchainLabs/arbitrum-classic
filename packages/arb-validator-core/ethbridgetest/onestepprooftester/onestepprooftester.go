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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820179721f1f2ac7ea116acd38d6debc7e337b65879c4b8ebdaa57878b2c275a6f964736f6c63430005110032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582029aaf06121529eac710057f2c4a1e46dccc3dc81de1c17f1a3a39338004527d064736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b50613d32806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e987d88714610030575b600080fd5b61015260048036036101e081101561004757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100dd57600080fd5b8201836020820111156100ef57600080fd5b8035906020019184600183028401116401000000008311171561011157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610164945050505050565b60408051918252519081900360200190f35b600061017a8d8d8d8d8d8d8d8d8d8d8d8d61018b565b9d9c50505050505050505050505050565b600061017a6040518061016001604052808f81526020018e81526020016101b28e8e6101f6565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff1681526020018481525061027c565b6101fe613b0a565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610263565b610250613b0a565b8152602001906001900390816102485790505b5081526002602082015260400183905290505b92915050565b600080600080606061028c613b3e565b610294613b3e565b61029d8861137b565b93995092965090945092509050600160006102b7886116ff565b67ffffffffffffffff168a610120015167ffffffffffffffff161461031a576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8960800151801561032e575060ff88166072145b8061034a5750896080015115801561034a575060ff8816607214155b61039b576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60ff8816600114156103e1576103da83866000815181106103b857fe5b6020026020010151876001815181106103cd57fe5b6020026020010151611b86565b91506111c6565b60ff881660021415610420576103da83866000815181106103fe57fe5b60200260200101518760018151811061041357fe5b6020026020010151611bd6565b60ff88166003141561045f576103da838660008151811061043d57fe5b60200260200101518760018151811061045257fe5b6020026020010151611c17565b60ff88166004141561049e576103da838660008151811061047c57fe5b60200260200101518760018151811061049157fe5b6020026020010151611c58565b60ff8816600514156104dd576103da83866000815181106104bb57fe5b6020026020010151876001815181106104d057fe5b6020026020010151611ca9565b60ff88166006141561051c576103da83866000815181106104fa57fe5b60200260200101518760018151811061050f57fe5b6020026020010151611cfa565b60ff88166007141561055b576103da838660008151811061053957fe5b60200260200101518760018151811061054e57fe5b6020026020010151611d4b565b60ff8816600814156105af576103da838660008151811061057857fe5b60200260200101518760018151811061058d57fe5b6020026020010151886002815181106105a257fe5b6020026020010151611d9c565b60ff881660091415610603576103da83866000815181106105cc57fe5b6020026020010151876001815181106105e157fe5b6020026020010151886002815181106105f657fe5b6020026020010151611e06565b60ff8816600a1415610642576103da838660008151811061062057fe5b60200260200101518760018151811061063557fe5b6020026020010151611e5f565b60ff881660101415610681576103da838660008151811061065f57fe5b60200260200101518760018151811061067457fe5b6020026020010151611ea0565b60ff8816601114156106c0576103da838660008151811061069e57fe5b6020026020010151876001815181106106b357fe5b6020026020010151611ee1565b60ff8816601214156106ff576103da83866000815181106106dd57fe5b6020026020010151876001815181106106f257fe5b6020026020010151611f22565b60ff88166013141561073e576103da838660008151811061071c57fe5b60200260200101518760018151811061073157fe5b6020026020010151611f63565b60ff88166014141561077d576103da838660008151811061075b57fe5b60200260200101518760018151811061077057fe5b6020026020010151611fa4565b60ff8816601514156107a7576103da838660008151811061079a57fe5b6020026020010151611fce565b60ff8816601614156107e6576103da83866000815181106107c457fe5b6020026020010151876001815181106107d957fe5b6020026020010151612013565b60ff881660171415610825576103da838660008151811061080357fe5b60200260200101518760018151811061081857fe5b6020026020010151612054565b60ff881660181415610864576103da838660008151811061084257fe5b60200260200101518760018151811061085757fe5b6020026020010151612095565b60ff88166019141561088e576103da838660008151811061088157fe5b60200260200101516120d6565b60ff8816601a14156108cd576103da83866000815181106108ab57fe5b6020026020010151876001815181106108c057fe5b602002602001015161210c565b60ff8816601b141561090c576103da83866000815181106108ea57fe5b6020026020010151876001815181106108ff57fe5b602002602001015161214d565b60ff881660201415610936576103da838660008151811061092957fe5b602002602001015161218e565b60ff881660211415610960576103da838660008151811061095357fe5b60200260200101516121a9565b60ff88166022141561099f576103da838660008151811061097d57fe5b60200260200101518760018151811061099257fe5b60200260200101516121c4565b60ff8816603014156109c9576103da83866000815181106109bc57fe5b602002602001015161222a565b60ff8816603114156109de576103da83612232565b60ff8816603214156109f3576103da83612253565b60ff881660331415610a1d576103da8386600081518110610a1057fe5b602002602001015161226c565b60ff881660341415610a47576103da8386600081518110610a3a57fe5b6020026020010151612278565b60ff881660351415610a86576103da8386600081518110610a6457fe5b602002602001015187600181518110610a7957fe5b602002602001015161227f565b60ff881660361415610a9b576103da836122bb565b60ff881660371415610ab5576103da8385600001516122e5565b60ff881660381415610adf576103da8386600081518110610ad257fe5b60200260200101516122f7565b60ff881660391415610b6b57610af3613b0a565b610b028b610140015188612309565b9199509750905087610b455760405162461bcd60e51b8152600401808060200182810382526021815260200180613cdd6021913960400191505060405180910390fd5b610b55858263ffffffff61244716565b610b65848263ffffffff61246116565b506111c6565b60ff8816603a1415610b80576103da8361247b565b60ff8816603b1415610b9557600191506111c6565b60ff8816603c1415610baa576103da83612498565b60ff8816603d1415610bd4576103da8386600081518110610bc757fe5b60200260200101516124b1565b60ff881660401415610bfe576103da8386600081518110610bf157fe5b60200260200101516124d5565b60ff881660411415610c3d576103da8386600081518110610c1b57fe5b602002602001015187600181518110610c3057fe5b60200260200101516124f7565b60ff881660421415610c91576103da8386600081518110610c5a57fe5b602002602001015187600181518110610c6f57fe5b602002602001015188600281518110610c8457fe5b6020026020010151612529565b60ff881660431415610cd0576103da8386600081518110610cae57fe5b602002602001015187600181518110610cc357fe5b602002602001015161256b565b60ff881660441415610d24576103da8386600081518110610ced57fe5b602002602001015187600181518110610d0257fe5b602002602001015188600281518110610d1757fe5b602002602001015161257d565b60ff881660501415610d63576103da8386600081518110610d4157fe5b602002602001015187600181518110610d5657fe5b602002602001015161259f565b60ff881660511415610db7576103da8386600081518110610d8057fe5b602002602001015187600181518110610d9557fe5b602002602001015188600281518110610daa57fe5b6020026020010151612615565b60ff881660521415610de1576103da8386600081518110610dd457fe5b60200260200101516126a2565b60ff881660601415610df6576103da836126d5565b60ff881660611415610ef457610e208386600081518110610e1357fe5b60200260200101516126db565b90925090508115610eeb578961010001518a60e001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610ea05760405162461bcd60e51b8152600401808060200182810382526025815260200180613c916025913960400191505060405180910390fd5b8960c001518a60a0015114610ee65760405162461bcd60e51b8152600401808060200182810382526027815260200180613cb66027913960400191505060405180910390fd5b610eef565b5060005b6111c6565b60ff88166070141561103457610f1e8386600081518110610f1157fe5b60200260200101516126f5565b90925090508115610eeb5780610f79578960c001518a60a0015114610f745760405162461bcd60e51b8152600401808060200182810382526038815260200180613c596038913960400191505060405180910390fd5b610ee6565b8960c001518a60a001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610fed5760405162461bcd60e51b8152600401808060200182810382526029815260200180613be96029913960400191505060405180910390fd5b8961010001518a60e0015114610ee65760405162461bcd60e51b8152600401808060200182810382526026815260200180613c126026913960400191505060405180910390fd5b60ff8816607114156111495760408051600480825260a08201909252606091816020015b611060613b0a565b81526020019060019003908161105857505060208c01519091506110949060005b60200201516001600160801b0316612725565b816000815181106110a157fe5b60200260200101819052506110c08b6020015160016004811061108157fe5b816001815181106110cd57fe5b60200260200101819052506110ec8b6020015160026004811061108157fe5b816002815181106110f957fe5b60200260200101819052506111188b6020015160036004811061108157fe5b8160038151811061112557fe5b6020026020010181905250610b6561113c826127aa565b859063ffffffff61246116565b60ff881660721415611197576103da838660008151811061116657fe5b60200260200101518c604001518d6020015160006004811061118457fe5b60200201516001600160801b0316612899565b60ff8816607314156111ac57600091506111c6565b60ff8816607414156111c157610eef83612930565b600091505b80611258578960c001518a60a00151146112115760405162461bcd60e51b8152600401808060200182810382526027815260200180613cb66027913960400191505060405180910390fd5b8961010001518a60e00151146112585760405162461bcd60e51b8152600401808060200182810382526026815260200180613c126026913960400191505060405180910390fd5b816112c25760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a08401516112a69061293a565b14156112ba576112b583612a30565b6112c2565b60a083015183525b6112cb84612a3a565b8a51146113095760405162461bcd60e51b8152600401808060200182810382526022815260200180613bc76022913960400191505060405180910390fd5b61131283612a3a565b8a6060015114611369576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b60006060611387613b3e565b61138f613b3e565b6000808061139b613b3e565b6113a481612b02565b6113b389610140015184612b0c565b90945090925090508161140d576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b611415613b3e565b61141e82612c1d565b905060008a6101400151858151811061143357fe5b602001015160f81c60f81b60f81c905060008b6101400151866001018151811061145957fe5b016020015160f81c9050600061146e82612c7b565b90506060816040519080825280602002602001820160405280156114ac57816020015b611499613b0a565b8152602001906001900390816114915790505b5090506002880197508360ff16600014806114ca57508360ff166001145b61151b576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff84166115405761153983611534886000015161293a565b612c95565b8652611608565b611548613b0a565b6115578f61014001518a612309565b909a509098509050876115b1576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b82156115d55780826000815181106115c557fe5b60200260200101819052506115e5565b6115e5868263ffffffff61246116565b611604846115f6896000015161293a565b6115ff8461293a565b612ccf565b8752505b60ff84165b8281101561169b576116248f61014001518a612309565b845185908590811061163257fe5b60209081029190910101529950975087611693576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b60010161160d565b8151156116e8575060005b8460ff168251038110156116e8576116e08282600185510303815181106116c957fe5b60200260200101518861246190919063ffffffff16565b6001016116a6565b50919d919c50939a50919850939650945050505050565b600060ff82166001141561171557506003611376565b60ff82166002141561172957506003611376565b60ff82166003141561173d57506003611376565b60ff82166004141561175157506004611376565b60ff82166005141561176557506007611376565b60ff82166006141561177957506004611376565b60ff82166007141561178d57506007611376565b60ff8216600814156117a157506004611376565b60ff8216600914156117b557506004611376565b60ff8216600a14156117c957506019611376565b60ff8216601014156117dd57506002611376565b60ff8216601114156117f157506002611376565b60ff82166012141561180557506002611376565b60ff82166013141561181957506002611376565b60ff82166014141561182d57506002611376565b60ff82166015141561184157506001611376565b60ff82166016141561185557506002611376565b60ff82166017141561186957506002611376565b60ff82166018141561187d57506002611376565b60ff82166019141561189157506001611376565b60ff8216601a14156118a557506004611376565b60ff8216601b14156118b957506007611376565b60ff8216602014156118cd57506007611376565b60ff8216602114156118e157506003611376565b60ff8216602214156118f557506008611376565b60ff82166030141561190957506001611376565b60ff82166031141561191d57506001611376565b60ff82166032141561193157506001611376565b60ff82166033141561194557506002611376565b60ff82166034141561195957506004611376565b60ff82166035141561196d57506004611376565b60ff82166036141561198157506002611376565b60ff82166037141561199557506001611376565b60ff8216603814156119a957506001611376565b60ff8216603914156119bd57506001611376565b60ff8216603a14156119d157506002611376565b60ff8216603b14156119e557506001611376565b60ff8216603c14156119f957506001611376565b60ff8216603d1415611a0d57506001611376565b60ff821660401415611a2157506001611376565b60ff821660411415611a3557506001611376565b60ff821660421415611a4957506001611376565b60ff821660431415611a5d57506001611376565b60ff821660441415611a7157506001611376565b60ff821660501415611a8557506002611376565b60ff821660511415611a9957506028611376565b60ff821660521415611aad57506002611376565b60ff821660601415611ac157506064611376565b60ff821660611415611ad557506064611376565b60ff821660701415611ae957506064611376565b60ff821660711415611afd57506028611376565b60ff821660721415611b1157506028611376565b60ff821660731415611b2557506005611376565b60ff821660741415611b395750600a611376565b6040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206f70636f64653a206f70476173436f737428290000000000604482015290519081900360640190fd5b6000611b9183612d06565b1580611ba35750611ba182612d06565b155b15611bb057506000611bcf565b82518251808201611bc7878263ffffffff612d1116565b600193505050505b9392505050565b6000611be183612d06565b1580611bf35750611bf182612d06565b155b15611c0057506000611bcf565b82518251808202611bc7878263ffffffff612d1116565b6000611c2283612d06565b1580611c345750611c3282612d06565b155b15611c4157506000611bcf565b82518251808203611bc7878263ffffffff612d1116565b6000611c6383612d06565b1580611c755750611c7382612d06565b155b15611c8257506000611bcf565b8251825180611c9657600092505050611bcf565b808204611bc7878263ffffffff612d1116565b6000611cb483612d06565b1580611cc65750611cc482612d06565b155b15611cd357506000611bcf565b8251825180611ce757600092505050611bcf565b808205611bc7878263ffffffff612d1116565b6000611d0583612d06565b1580611d175750611d1582612d06565b155b15611d2457506000611bcf565b8251825180611d3857600092505050611bcf565b808206611bc7878263ffffffff612d1116565b6000611d5683612d06565b1580611d685750611d6682612d06565b155b15611d7557506000611bcf565b8251825180611d8957600092505050611bcf565b808207611bc7878263ffffffff612d1116565b6000611da784612d06565b1580611db95750611db783612d06565b155b15611dc657506000611dfe565b83518351835180611ddd5760009350505050611dfe565b6000818385089050611df5898263ffffffff612d1116565b60019450505050505b949350505050565b6000611e1184612d06565b1580611e235750611e2183612d06565b155b15611e3057506000611dfe565b83518351835180611e475760009350505050611dfe565b6000818385099050611df5898263ffffffff612d1116565b6000611e6a83612d06565b1580611e7c5750611e7a82612d06565b155b15611e8957506000611bcf565b8251825180820a611bc7878263ffffffff612d1116565b6000611eab83612d06565b1580611ebd5750611ebb82612d06565b155b15611eca57506000611bcf565b82518251808210611bc7878263ffffffff612d1116565b6000611eec83612d06565b1580611efe5750611efc82612d06565b155b15611f0b57506000611bcf565b82518251808211611bc7878263ffffffff612d1116565b6000611f2d83612d06565b1580611f3f5750611f3d82612d06565b155b15611f4c57506000611bcf565b82518251808212611bc7878263ffffffff612d1116565b6000611f6e83612d06565b1580611f805750611f7e82612d06565b155b15611f8d57506000611bcf565b82518251808213611bc7878263ffffffff612d1116565b6000611fc461113c611fb58461293a565b611fbe8661293a565b14612d27565b5060019392505050565b6000611fd982612d06565b611ff357611fee83600063ffffffff612d1116565b61200a565b81518015612007858263ffffffff612d1116565b50505b50600192915050565b600061201e83612d06565b1580612030575061202e82612d06565b155b1561203d57506000611bcf565b82518251808216611bc7878263ffffffff612d1116565b600061205f83612d06565b1580612071575061206f82612d06565b155b1561207e57506000611bcf565b82518251808217611bc7878263ffffffff612d1116565b60006120a083612d06565b15806120b257506120b082612d06565b155b156120bf57506000611bcf565b82518251808218611bc7878263ffffffff612d1116565b60006120e182612d06565b6120ed57506000610276565b81518019612101858263ffffffff612d1116565b506001949350505050565b600061211783612d06565b1580612129575061212782612d06565b155b1561213657506000611bcf565b8251825181811a611bc7878263ffffffff612d1116565b600061215883612d06565b158061216a575061216882612d06565b155b1561217757506000611bcf565b8251825181810b611bc7878263ffffffff612d1116565b600061200a61219c8361293a565b849063ffffffff612d1116565b600061200a6121b783612d49565b849063ffffffff61246116565b60006121cf83612d06565b15806121e157506121df82612d06565b155b156121ee57506000611bcf565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120611bc7878263ffffffff612d1116565b600192915050565b600061224b82608001518361246190919063ffffffff16565b506001919050565b600061224b82606001518361246190919063ffffffff16565b60609190910152600190565b9052600190565b600061228a83612dd2565b61229657506000611bcf565b61229f82612d06565b6122ab57506000611bcf565b815115611fc45750509052600190565b600061224b6122d86122cb612ddf565b611fbe856020015161293a565b839063ffffffff61246116565b600061200a838363ffffffff61246116565b600061200a838363ffffffff61244716565b600080612314613b0a565b84518410612334576000846123296000612725565b925092509250612440565b600080859050600087828151811061234857fe5b016020015160019092019160f81c90506000612362613b9f565b60ff8316612396576123748a85612e00565b91965094509150848461238684612725565b9750975097505050505050612440565b60ff8316600114156123be576123ac8a85612e53565b91965094509050848461238683612fb3565b60ff8316600214156123d4576123868a8561301a565b600360ff8416108015906123eb5750600c60ff8416105b156124265760021983016060612402828d886130bf565b919850965090508686612414836127aa565b99509950995050505050505050612440565b6000806124336000612725565b9199509750955050505050505b9250925092565b61245582604001518261317d565b82604001819052505050565b61246f82602001518261317d565b82602001819052505050565b600061224b6122d861248b612ddf565b611fbe856040015161293a565b600061224b8260a001518361246190919063ffffffff16565b60006124bc82612dd2565b6124c857506000610276565b5060a09190910152600190565b60006124e7838363ffffffff61246116565b61200a838363ffffffff61246116565b6000612509848363ffffffff61246116565b612519848463ffffffff61246116565b611fc4848363ffffffff61246116565b600061253b858363ffffffff61246116565b61254b858463ffffffff61246116565b61255b858563ffffffff61246116565b612101858363ffffffff61246116565b6000612519848463ffffffff61246116565b600061258f858563ffffffff61246116565b61255b858463ffffffff61246116565b60006125aa83612d06565b15806125bc57506125ba826131fb565b155b156125c957506000611bcf565b6125d28261320a565b60ff168360000151106125e757506000611bcf565b611fc482604001518460000151815181106125fe57fe5b60200260200101518561246190919063ffffffff16565b6000612620836131fb565b1580612632575061263084612d06565b155b1561263f57506000611dfe565b6126488361320a565b60ff1684600001511061265d57506000611dfe565b60408301518451815184918391811061267257fe5b6020026020010181905250612696612689826127aa565b879063ffffffff61246116565b50600195945050505050565b60006126ad826131fb565b6126b957506000610276565b61200a6126c58361320a565b849060ff1663ffffffff612d1116565b50600190565b60008060016126e98461293a565b915091505b9250929050565b6000806127108360800151116127195760016127108461293a565b915091506126ee565b506001905060006126ee565b61272d613b0a565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612792565b61277f613b0a565b8152602001906001900390816127775790505b50815260006020820152600160409091015292915050565b6127b2613b0a565b6127bc8251613219565b61280d576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156128445783818151811061282757fe5b602002602001015160800151820191508080600101915050612812565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b60006128a484612d06565b6128b057506000611dfe565b8351821115806128cf57506128c3612ddf565b6128cc8461293a565b14155b612920576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b612101858463ffffffff61246116565b600260c090910152565b6000600360090160ff16826060015160ff1610612992576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166129b05781516129a990613220565b9050611376565b606082015160ff16600114156129e35760208083015180516040820151606083015192909301516129a993919290613244565b606082015160ff16600214156129fc576129a9826132ec565b600360ff16826060015160ff1610158015612a2057506060820151600c60ff909116105b15612a2e576129a982613352565bfe5b600160c090910152565b600060028260c001511415612a5157506000611376565b60018260c001511415612a6657506001611376565b8151612a719061293a565b612a7e836020015161293a565b612a8b846040015161293a565b612a98856060015161293a565b612aa5866080015161293a565b612ab28760a0015161293a565b604051602001808781526020018681526020018581526020018481526020018381526020018281526020019650505050505050604051602081830303815290604052805190602001209050611376565b600060c090910152565b600080612b17613b3e565b612b1f613b3e565b600060c08201819052612b328787612309565b84529650905080612b4c5750600093508492509050612440565b612b56878761301a565b60208501529650905080612b735750600093508492509050612440565b612b7d878761301a565b60408501529650905080612b9a5750600093508492509050612440565b612ba48787612309565b60608501529650905080612bc15750600093508492509050612440565b612bcb8787612309565b60808501529650905080612be85750600093508492509050612440565b612bf28787612309565b60a08501529650905080612c0f5750600093508492509050612440565b506001969495509392505050565b612c25613b3e565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b6000806000612c8c8460ff16613370565b50949350505050565b612c9d613b0a565b611bcf60405180608001604052808560ff1681526020018481526020016000151581526020016000801b815250612fb3565b612cd7613b0a565b611dfe60405180608001604052808660ff16815260200185815260200160011515815260200184815250612fb3565b6060015160ff161590565b61246f8260200151612d2283612725565b61317d565b612d2f613b0a565b8115612d3f576129a96001612725565b6129a96000612725565b612d51613b0a565b816060015160ff1660021415612d985760405162461bcd60e51b8152600401808060200182810382526021815260200180613c386021913960400191505060405180910390fd5b606082015160ff16612dae576129a96000612725565b816060015160ff1660011415612dc8576129a96001612725565b6129a96003612725565b6060015160ff1660011490565b60408051600080825260208201909252612dfa816001613827565b91505090565b6000806000808551905084811080612e1a57506020858203105b15612e2f575060009250839150829050612440565b600160208601612e45888863ffffffff61384616565b935093509350509250925092565b600080612e5e613b9f565b60008490506000868281518110612e7157fe5b602001015160f81c60f81b60f81c905081806001019250506000878381518110612e9757fe5b016020015160019384019360f89190911c915060009060ff84161415612f1d576000612ec1613b0a565b612ecb8b87612309565b909750909250905081612f0f575050604080516080810182526000808252602082018190529181018290526060810182905290975088965094506124409350505050565b612f188161293a565b925050505b6000612f2f8a8663ffffffff61384616565b90506020850194508360ff1660011415612f7b576040805160808101825260ff909416845260208401919091526001908301819052606083019190915295509193509091506124409050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b612fbb613b0a565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190613002565b612fef613b0a565b815260200190600190039081612fe75790505b50815260016020820181905260409091015292915050565b600080613025613b0a565b61302d613b0a565b855160009081908781108061304457506040888203105b1561305c576000888596509650965050505050612440565b600061306e8a8a63ffffffff61384616565b90506020890198506130808a8a612e00565b909a509450925082156130ab5761309781856101f6565b600198508997509550612440945050505050565b600089869750975097505050505050612440565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561310a57816020015b6130f7613b0a565b8152602001906001900390816130ef5790505b50905060005b8960ff168160ff161015613167576131288985612309565b8451859060ff861690811061313957fe5b6020908102919091010152945092508261315f5750600095508694509250613174915050565b600101613110565b5060019550919350909150505b93509350939050565b613185613b0a565b6040805160028082526060828101909352816020015b6131a3613b0a565b81526020019060019003908161319b57905050905082816000815181106131c657fe5b602002602001018190525083816001815181106131df57fe5b6020026020010181905250611dfe6131f6826127aa565b613862565b600061027682606001516138d8565b600061027682606001516138f6565b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561329e575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611dfe565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214613341576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b815160808301516102769190613919565b600061335c613b0a565b61336583613862565b9050611bcf816132ec565b60008060018314156133885750600290506001613822565b600283141561339d5750600290506001613822565b60038314156133b25750600290506001613822565b60048314156133c75750600290506001613822565b60058314156133dc5750600290506001613822565b60068314156133f15750600290506001613822565b60078314156134065750600290506001613822565b600883141561341b5750600390506001613822565b60098314156134305750600390506001613822565b600a8314156134455750600290506001613822565b601083141561345a5750600290506001613822565b601183141561346f5750600290506001613822565b60128314156134845750600290506001613822565b60138314156134995750600290506001613822565b60148314156134ae5750600290506001613822565b60158314156134c257506001905080613822565b60168314156134d75750600290506001613822565b60178314156134ec5750600290506001613822565b60188314156135015750600290506001613822565b601983141561351557506001905080613822565b601a83141561352a5750600290506001613822565b601b83141561353f5750600290506001613822565b602083141561355357506001905080613822565b602183141561356757506001905080613822565b602283141561357c5750600290506001613822565b60308314156135915750600190506000613822565b60318314156135a65750600090506001613822565b60328314156135bb5750600090506001613822565b60338314156135d05750600190506000613822565b60348314156135e55750600190506000613822565b60358314156135fa5750600290506000613822565b603683141561360f5750600090506001613822565b60378314156136245750600090506001613822565b60388314156136395750600190506000613822565b603983141561364e5750600090506001613822565b603a8314156136635750600090506001613822565b603b83141561367757506000905080613822565b603c83141561368c5750600090506001613822565b603d8314156136a15750600190506000613822565b60408314156136b65750600190506002613822565b60418314156136cb5750600290506003613822565b60428314156136e05750600390506004613822565b60438314156136f457506002905080613822565b604483141561370857506003905080613822565b605083141561371d5750600290506001613822565b60518314156137325750600390506001613822565b605283141561374657506001905080613822565b606083141561375a57506000905080613822565b606183141561376f5750600190506000613822565b60708314156137845750600190506000613822565b60718314156137995750600090506001613822565b60728314156137ad57506001905080613822565b60738314156137c157506000905080613822565b60748314156137d557506000905080613822565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b6000613831613b0a565b61383b8484613953565b9050611dfe816132ec565b6000816020018351101561385957600080fd5b50016020015190565b61386a613b0a565b613873826131fb565b6138b9576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b60606138c88360400151613972565b9050611bcf818460800151613953565b6000600c60ff8316108015610276575050600360ff91909116101590565b6000613901826138d8565b1561391157506002198101611376565b506001611376565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b61395b613b0a565b600061396684613a4a565b9050611dfe81846101f6565b60606008825111156139c2576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156139ef578160200160208202803883390190505b50805190915060005b81811015613a41576000613a1e868381518110613a1157fe5b602002602001015161293a565b905080848381518110613a2d57fe5b6020908102919091010152506001016139f8565b50909392505050565b6000600882511115613a9a576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613ade578181015183820152602001613ac6565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001613b24613b9f565b815260606020820181905260006040830181905291015290565b6040518060e00160405280613b51613b0a565b8152602001613b5e613b0a565b8152602001613b6b613b0a565b8152602001613b78613b0a565b8152602001613b85613b0a565b8152602001613b92613b0a565b8152602001600081525090565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820764115e9947f224761293f71fc6399763b1af6bbda804eb0330503fb85cf140c64736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820c775845ddc314c7c4421bfc9bf144c44e6fa0fd28b736f3046bcfdf0e698694d64736f6c63430005110032"

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
