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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582022116fffad8c531f1b75c6f53e271efcc426b7d76e13cb0acd2c182bd92366a764736f6c63430005110032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200eaa7d1c7ea5b6de61795d05939865abcf8c7764445171ea8c21ef4a4a5a8a1264736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b506144c2806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e987d88714610030575b600080fd5b61015260048036036101e081101561004757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100dd57600080fd5b8201836020820111156100ef57600080fd5b8035906020019184600183028401116401000000008311171561011157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610164945050505050565b60408051918252519081900360200190f35b600061017a8d8d8d8d8d8d8d8d8d8d8d8d61018b565b9d9c50505050505050505050505050565b600061017a6040518061016001604052808f81526020018e81526020016101b28e8e6101f6565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff1681526020018481525061027c565b6101fe61429d565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610263565b61025061429d565b8152602001906001900390816102485790505b5081526002602082015260400183905290505b92915050565b600080600080606061028c6142d1565b6102946142d1565b61029d886115b2565b93995092965090945092509050600160006102b788611931565b67ffffffffffffffff168a610120015167ffffffffffffffff161461031a576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8960800151801561032e575060ff88166072145b8061034a5750896080015115801561034a575060ff8816607214155b61039b576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b6103a488611931565b67ffffffffffffffff168460a0015110156103ca5760001960a0840152600091506113e4565b60ff8816600114156104105761040983866000815181106103e757fe5b6020026020010151876001815181106103fc57fe5b6020026020010151611e1d565b91506113e4565b60ff88166002141561044f57610409838660008151811061042d57fe5b60200260200101518760018151811061044257fe5b6020026020010151611e6d565b60ff88166003141561048e57610409838660008151811061046c57fe5b60200260200101518760018151811061048157fe5b6020026020010151611eae565b60ff8816600414156104cd5761040983866000815181106104ab57fe5b6020026020010151876001815181106104c057fe5b6020026020010151611eef565b60ff88166005141561050c5761040983866000815181106104ea57fe5b6020026020010151876001815181106104ff57fe5b6020026020010151611f40565b60ff88166006141561054b57610409838660008151811061052957fe5b60200260200101518760018151811061053e57fe5b6020026020010151611f91565b60ff88166007141561058a57610409838660008151811061056857fe5b60200260200101518760018151811061057d57fe5b6020026020010151611fe2565b60ff8816600814156105de5761040983866000815181106105a757fe5b6020026020010151876001815181106105bc57fe5b6020026020010151886002815181106105d157fe5b6020026020010151612033565b60ff8816600914156106325761040983866000815181106105fb57fe5b60200260200101518760018151811061061057fe5b60200260200101518860028151811061062557fe5b602002602001015161209d565b60ff8816600a141561067157610409838660008151811061064f57fe5b60200260200101518760018151811061066457fe5b60200260200101516120f6565b60ff8816601014156106b057610409838660008151811061068e57fe5b6020026020010151876001815181106106a357fe5b6020026020010151612137565b60ff8816601114156106ef5761040983866000815181106106cd57fe5b6020026020010151876001815181106106e257fe5b6020026020010151612178565b60ff88166012141561072e57610409838660008151811061070c57fe5b60200260200101518760018151811061072157fe5b60200260200101516121b9565b60ff88166013141561076d57610409838660008151811061074b57fe5b60200260200101518760018151811061076057fe5b60200260200101516121fa565b60ff8816601414156107ac57610409838660008151811061078a57fe5b60200260200101518760018151811061079f57fe5b602002602001015161223b565b60ff8816601514156107d65761040983866000815181106107c957fe5b6020026020010151612265565b60ff8816601614156108155761040983866000815181106107f357fe5b60200260200101518760018151811061080857fe5b60200260200101516122aa565b60ff88166017141561085457610409838660008151811061083257fe5b60200260200101518760018151811061084757fe5b60200260200101516122eb565b60ff88166018141561089357610409838660008151811061087157fe5b60200260200101518760018151811061088657fe5b602002602001015161232c565b60ff8816601914156108bd5761040983866000815181106108b057fe5b602002602001015161236d565b60ff8816601a14156108fc5761040983866000815181106108da57fe5b6020026020010151876001815181106108ef57fe5b60200260200101516123a3565b60ff8816601b141561093b57610409838660008151811061091957fe5b60200260200101518760018151811061092e57fe5b60200260200101516123e4565b60ff88166020141561096557610409838660008151811061095857fe5b6020026020010151612425565b60ff88166021141561098f57610409838660008151811061098257fe5b6020026020010151612440565b60ff8816602214156109ce5761040983866000815181106109ac57fe5b6020026020010151876001815181106109c157fe5b602002602001015161245b565b60ff8816603014156109f85761040983866000815181106109eb57fe5b60200260200101516124c1565b60ff881660311415610a0d57610409836124c9565b60ff881660321415610a2257610409836124ea565b60ff881660331415610a4c576104098386600081518110610a3f57fe5b6020026020010151612503565b60ff881660341415610a76576104098386600081518110610a6957fe5b602002602001015161250f565b60ff881660351415610ab5576104098386600081518110610a9357fe5b602002602001015187600181518110610aa857fe5b602002602001015161253a565b60ff881660361415610aca5761040983612582565b60ff881660371415610ae4576104098385600001516125ac565b60ff881660381415610b0e576104098386600081518110610b0157fe5b60200260200101516125ba565b60ff881660391415610b9a57610b2261429d565b610b318b6101400151886125cc565b9199509750905087610b745760405162461bcd60e51b815260040180806020018281038252602181526020018061446d6021913960400191505060405180910390fd5b610b84858263ffffffff61270a16565b610b94848263ffffffff61272416565b506113e4565b60ff8816603a1415610baf576104098361273e565b60ff8816603b1415610bc457600191506113e4565b60ff8816603c1415610bd9576104098361275b565b60ff8816603d1415610c03576104098386600081518110610bf657fe5b602002602001015161276d565b60ff881660401415610c2d576104098386600081518110610c2057fe5b602002602001015161279b565b60ff881660411415610c6c576104098386600081518110610c4a57fe5b602002602001015187600181518110610c5f57fe5b60200260200101516127bd565b60ff881660421415610cc0576104098386600081518110610c8957fe5b602002602001015187600181518110610c9e57fe5b602002602001015188600281518110610cb357fe5b60200260200101516127ef565b60ff881660431415610cff576104098386600081518110610cdd57fe5b602002602001015187600181518110610cf257fe5b6020026020010151612831565b60ff881660441415610d53576104098386600081518110610d1c57fe5b602002602001015187600181518110610d3157fe5b602002602001015188600281518110610d4657fe5b6020026020010151612843565b60ff881660501415610d92576104098386600081518110610d7057fe5b602002602001015187600181518110610d8557fe5b6020026020010151612865565b60ff881660511415610de6576104098386600081518110610daf57fe5b602002602001015187600181518110610dc457fe5b602002602001015188600281518110610dd957fe5b60200260200101516128db565b60ff881660521415610e10576104098386600081518110610e0357fe5b6020026020010151612968565b60ff881660531415610ead57610e2461429d565b610e338b6101400151886125cc565b9199509750905087610e765760405162461bcd60e51b815260040180806020018281038252602181526020018061446d6021913960400191505060405180910390fd5b610e86858263ffffffff61270a16565b610ea58487600081518110610e9757fe5b60200260200101518361299b565b9250506113e4565b60ff881660541415610f5757610ec161429d565b610ed08b6101400151886125cc565b9199509750905087610f135760405162461bcd60e51b815260040180806020018281038252602181526020018061446d6021913960400191505060405180910390fd5b610f23858263ffffffff61270a16565b610ea58487600081518110610f3457fe5b602002602001015188600181518110610f4957fe5b6020026020010151846129f3565b60ff881660601415610f6c5761040983612a74565b60ff88166061141561106a57610f968386600081518110610f8957fe5b6020026020010151612a7a565b90925090508115611061578961010001518a60e0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146110165760405162461bcd60e51b81526004018080602001828103825260258152602001806144216025913960400191505060405180910390fd5b8960c001518a60a001511461105c5760405162461bcd60e51b81526004018080602001828103825260278152602001806144466027913960400191505060405180910390fd5b611065565b5060005b6113e4565b60ff8816607014156111aa57611094838660008151811061108757fe5b6020026020010151612a94565b9092509050811561106157806110ef578960c001518a60a00151146110ea5760405162461bcd60e51b81526004018080602001828103825260388152602001806143e96038913960400191505060405180910390fd5b61105c565b8960c001518a60a0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146111635760405162461bcd60e51b81526004018080602001828103825260298152602001806143796029913960400191505060405180910390fd5b8961010001518a60e001511461105c5760405162461bcd60e51b81526004018080602001828103825260268152602001806143a26026913960400191505060405180910390fd5b60ff8816607114156112bf5760408051600480825260a08201909252606091816020015b6111d661429d565b8152602001906001900390816111ce57505060208c015190915061120a9060005b60200201516001600160801b0316612ad3565b8160008151811061121757fe5b60200260200101819052506112368b602001516001600481106111f757fe5b8160018151811061124357fe5b60200260200101819052506112628b602001516002600481106111f757fe5b8160028151811061126f57fe5b602002602001018190525061128e8b602001516003600481106111f757fe5b8160038151811061129b57fe5b6020026020010181905250610b946112b282612b58565b859063ffffffff61272416565b60ff88166072141561130d5761040983866000815181106112dc57fe5b60200260200101518c604001518d602001516000600481106112fa57fe5b60200201516001600160801b0316612c47565b60ff88166073141561132257600091506113e4565b60ff8816607414156113375761106583612cde565b60ff88166075141561136157610409838660008151811061135457fe5b6020026020010151612ce8565b60ff8816607614156113765761040983612d0d565b60ff8816608014156113df57610409838660008151811061139357fe5b6020026020010151876001815181106113a857fe5b6020026020010151886002815181106113bd57fe5b6020026020010151896003815181106113d257fe5b6020026020010151612d26565b600091505b80611476578960c001518a60a001511461142f5760405162461bcd60e51b81526004018080602001828103825260278152602001806144466027913960400191505060405180910390fd5b8961010001518a60e00151146114765760405162461bcd60e51b81526004018080602001828103825260268152602001806143a26026913960400191505060405180910390fd5b61147f88611931565b67ffffffffffffffff168360a00151038360a0018181525050816114f95760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156114f1576114ec83612e45565b6114f9565b60c083015183525b61150284612e4f565b8a51146115405760405162461bcd60e51b81526004018080602001828103825260228152602001806143576022913960400191505060405180910390fd5b61154983612e4f565b8a60600151146115a0576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b600060606115be6142d1565b6115c66142d1565b600080806115d26142d1565b6115db81612f13565b6115ea89610140015184612f1d565b909450909250905081611644576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b61164c6142d1565b61165582613055565b905060008a6101400151858151811061166a57fe5b602001015160f81c60f81b60f81c905060008b6101400151866001018151811061169057fe5b016020015160f81c905060006116a5826130be565b90506060816040519080825280602002602001820160405280156116e357816020015b6116d061429d565b8152602001906001900390816116c85790505b5090506002880197508360ff166000148061170157508360ff166001145b611752576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff84166117775761177061176b8488600001516130d8565b613112565b865261183a565b61177f61429d565b61178e8f61014001518a6125cc565b909a509098509050876117e8576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561180c5780826000815181106117fc57fe5b602002602001018190525061181c565b61181c868263ffffffff61272416565b61183661176b85896000015161183185613112565b61320a565b8752505b60ff84165b828110156118cd576118568f61014001518a6125cc565b845185908590811061186457fe5b602090810291909101015299509750876118c5576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b60010161183f565b81511561191a575060005b8460ff1682510381101561191a576119128282600185510303815181106118fb57fe5b60200260200101518861272490919063ffffffff16565b6001016118d8565b50919d919c50939a50919850939650945050505050565b600060ff821660011415611947575060036115ad565b60ff82166002141561195b575060036115ad565b60ff82166003141561196f575060036115ad565b60ff821660041415611983575060046115ad565b60ff821660051415611997575060076115ad565b60ff8216600614156119ab575060046115ad565b60ff8216600714156119bf575060076115ad565b60ff8216600814156119d3575060046115ad565b60ff8216600914156119e7575060046115ad565b60ff8216600a14156119fb575060196115ad565b60ff821660101415611a0f575060026115ad565b60ff821660111415611a23575060026115ad565b60ff821660121415611a37575060026115ad565b60ff821660131415611a4b575060026115ad565b60ff821660141415611a5f575060026115ad565b60ff821660151415611a73575060016115ad565b60ff821660161415611a87575060026115ad565b60ff821660171415611a9b575060026115ad565b60ff821660181415611aaf575060026115ad565b60ff821660191415611ac3575060016115ad565b60ff8216601a1415611ad7575060046115ad565b60ff8216601b1415611aeb575060076115ad565b60ff821660201415611aff575060076115ad565b60ff821660211415611b13575060036115ad565b60ff821660221415611b27575060086115ad565b60ff821660301415611b3b575060016115ad565b60ff821660311415611b4f575060016115ad565b60ff821660321415611b63575060016115ad565b60ff821660331415611b77575060026115ad565b60ff821660341415611b8b575060046115ad565b60ff821660351415611b9f575060046115ad565b60ff821660361415611bb3575060026115ad565b60ff821660371415611bc7575060016115ad565b60ff821660381415611bdb575060016115ad565b60ff821660391415611bef575060016115ad565b60ff8216603a1415611c03575060026115ad565b60ff8216603b1415611c17575060016115ad565b60ff8216603c1415611c2b575060016115ad565b60ff8216603d1415611c3f575060016115ad565b60ff821660401415611c53575060016115ad565b60ff821660411415611c67575060016115ad565b60ff821660421415611c7b575060016115ad565b60ff821660431415611c8f575060016115ad565b60ff821660441415611ca3575060016115ad565b60ff821660501415611cb7575060026115ad565b60ff821660511415611ccb575060286115ad565b60ff821660521415611cdf575060026115ad565b60ff821660531415611cf3575060036115ad565b60ff821660541415611d07575060296115ad565b60ff821660601415611d1b575060646115ad565b60ff821660611415611d2f575060646115ad565b60ff821660701415611d43575060646115ad565b60ff821660711415611d57575060286115ad565b60ff821660721415611d6b575060286115ad565b60ff821660731415611d7f575060056115ad565b60ff821660741415611d935750600a6115ad565b60ff821660751415611da7575060006115ad565b60ff821660761415611dbb575060016115ad565b60ff821660801415611dd05750614e206115ad565b6040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206f70636f64653a206f70476173436f737428290000000000604482015290519081900360640190fd5b6000611e2883613241565b1580611e3a5750611e3882613241565b155b15611e4757506000611e66565b82518251808201611e5e878263ffffffff61324c16565b600193505050505b9392505050565b6000611e7883613241565b1580611e8a5750611e8882613241565b155b15611e9757506000611e66565b82518251808202611e5e878263ffffffff61324c16565b6000611eb983613241565b1580611ecb5750611ec982613241565b155b15611ed857506000611e66565b82518251808203611e5e878263ffffffff61324c16565b6000611efa83613241565b1580611f0c5750611f0a82613241565b155b15611f1957506000611e66565b8251825180611f2d57600092505050611e66565b808204611e5e878263ffffffff61324c16565b6000611f4b83613241565b1580611f5d5750611f5b82613241565b155b15611f6a57506000611e66565b8251825180611f7e57600092505050611e66565b808205611e5e878263ffffffff61324c16565b6000611f9c83613241565b1580611fae5750611fac82613241565b155b15611fbb57506000611e66565b8251825180611fcf57600092505050611e66565b808206611e5e878263ffffffff61324c16565b6000611fed83613241565b1580611fff5750611ffd82613241565b155b1561200c57506000611e66565b825182518061202057600092505050611e66565b808207611e5e878263ffffffff61324c16565b600061203e84613241565b1580612050575061204e83613241565b155b1561205d57506000612095565b835183518351806120745760009350505050612095565b600081838508905061208c898263ffffffff61324c16565b60019450505050505b949350505050565b60006120a884613241565b15806120ba57506120b883613241565b155b156120c757506000612095565b835183518351806120de5760009350505050612095565b600081838509905061208c898263ffffffff61324c16565b600061210183613241565b1580612113575061211182613241565b155b1561212057506000611e66565b8251825180820a611e5e878263ffffffff61324c16565b600061214283613241565b1580612154575061215282613241565b155b1561216157506000611e66565b82518251808210611e5e878263ffffffff61324c16565b600061218383613241565b1580612195575061219382613241565b155b156121a257506000611e66565b82518251808211611e5e878263ffffffff61324c16565b60006121c483613241565b15806121d657506121d482613241565b155b156121e357506000611e66565b82518251808212611e5e878263ffffffff61324c16565b600061220583613241565b1580612217575061221582613241565b155b1561222457506000611e66565b82518251808213611e5e878263ffffffff61324c16565b600061225b6112b261224c84613112565b61225586613112565b14613262565b5060019392505050565b600061227082613241565b61228a5761228583600063ffffffff61324c16565b6122a1565b8151801561229e858263ffffffff61324c16565b50505b50600192915050565b60006122b583613241565b15806122c757506122c582613241565b155b156122d457506000611e66565b82518251808216611e5e878263ffffffff61324c16565b60006122f683613241565b1580612308575061230682613241565b155b1561231557506000611e66565b82518251808217611e5e878263ffffffff61324c16565b600061233783613241565b1580612349575061234782613241565b155b1561235657506000611e66565b82518251808218611e5e878263ffffffff61324c16565b600061237882613241565b61238457506000610276565b81518019612398858263ffffffff61324c16565b506001949350505050565b60006123ae83613241565b15806123c057506123be82613241565b155b156123cd57506000611e66565b8251825181811a611e5e878263ffffffff61324c16565b60006123ef83613241565b158061240157506123ff82613241565b155b1561240e57506000611e66565b8251825181810b611e5e878263ffffffff61324c16565b60006122a161243383613112565b849063ffffffff61324c16565b60006122a161244e83613284565b849063ffffffff61272416565b600061246683613241565b1580612478575061247682613241565b155b1561248557506000611e66565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120611e5e878263ffffffff61324c16565b600192915050565b60006124e282608001518361272490919063ffffffff16565b506001919050565b60006124e282606001518361272490919063ffffffff16565b60609190910152600190565b600061251a8261330d565b61252657506000610276565b61252f82613112565b835250600192915050565b60006125458361330d565b61255157506000611e66565b61255a82613241565b61256657506000611e66565b81511561225b5761257683613112565b84525060019392505050565b60006124e261259f61259261331a565b6122558560200151613112565b839063ffffffff61272416565b60006122a161244e8361333b565b60006122a1838363ffffffff61270a16565b6000806125d761429d565b845184106125f7576000846125ec6000612ad3565b925092509250612703565b600080859050600087828151811061260b57fe5b016020015160019092019160f81c9050600061262561432f565b60ff8316612659576126378a856133c0565b91965094509150848461264984612ad3565b9750975097505050505050612703565b60ff8316600114156126815761266f8a85613413565b91965094509050848461264983613573565b60ff831660021415612697576126498a856135da565b600360ff8416108015906126ae5750600c60ff8416105b156126e957600219830160606126c5828d8861367f565b9198509650905086866126d783612b58565b99509950995050505050505050612703565b6000806126f66000612ad3565b9199509750955050505050505b9250925092565b61271882604001518261373d565b82604001819052505050565b61273282602001518261373d565b82602001819052505050565b60006124e261259f61274e61331a565b6122558560400151613112565b60006124e261259f8360c0015161333b565b60006127788261330d565b61278457506000610276565b61278d82613112565b60c084015250600192915050565b60006127ad838363ffffffff61272416565b6122a1838363ffffffff61272416565b60006127cf848363ffffffff61272416565b6127df848463ffffffff61272416565b61225b848363ffffffff61272416565b6000612801858363ffffffff61272416565b612811858463ffffffff61272416565b612821858563ffffffff61272416565b612398858363ffffffff61272416565b60006127df848463ffffffff61272416565b6000612855858563ffffffff61272416565b612821858463ffffffff61272416565b600061287083613241565b15806128825750612880826137bb565b155b1561288f57506000611e66565b612898826137ca565b60ff168360000151106128ad57506000611e66565b61225b82604001518460000151815181106128c457fe5b60200260200101518561272490919063ffffffff16565b60006128e6836137bb565b15806128f857506128f684613241565b155b1561290557506000612095565b61290e836137ca565b60ff1684600001511061292357506000612095565b60408301518451815184918391811061293857fe5b602002602001018190525061295c61294f82612b58565b879063ffffffff61272416565b50600195945050505050565b6000612973826137bb565b61297f57506000610276565b6122a161298b836137ca565b849060ff1663ffffffff61324c16565b60006129a683613241565b15806129b857506129b6826137bb565b155b156129c557506000611e66565b6129ce826137ca565b60ff168360000151106129e357506000611e66565b6128ad848363ffffffff61270a16565b60006129fe826137bb565b1580612a105750612a0e84613241565b155b15612a1d57506000612095565b612a26826137ca565b60ff16846000015110612a3b57506000612095565b604082015184518151859183918110612a5057fe5b602002602001018190525061295c612a6782612b58565b879063ffffffff61270a16565b50600190565b6000806001612a8884613112565b915091505b9250929050565b60008061271083608001511115612ab057506000905080612a8d565b612ab9836137d9565b612ac857506000905080612a8d565b6001612a8884613112565b612adb61429d565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612b40565b612b2d61429d565b815260200190600190039081612b255790505b50815260006020820152600160409091015292915050565b612b6061429d565b612b6a82516138f0565b612bbb576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015612bf257838181518110612bd557fe5b602002602001015160800151820191508080600101915050612bc0565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b6000612c5284613241565b612c5e57506000612095565b835182111580612c7d5750612c7161331a565b612c7a84613112565b14155b612cce576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b612398858463ffffffff61272416565b600260e090910152565b6000612cf382613241565b612cff57506000610276565b505160a09190910152600190565b60006124e28260a001518361324c90919063ffffffff16565b6000612d3185613241565b1580612d435750612d4184613241565b155b80612d545750612d5283613241565b155b80612d655750612d6382613241565b155b15612d7257506000612e3c565b84518451845115801590612d8857508451600114155b15612da957612d9e88600063ffffffff61324c16565b600192505050612e3c565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612e0b573d6000803e3d6000fd5b5050604051601f1901519150612e3290508b6001600160a01b03831663ffffffff61324c16565b6001955050505050505b95945050505050565b600160e090910152565b600060028260e001511415612e66575060006115ad565b60018260e001511415612e7b575060016115ad565b81516020830151612e8b90613112565b612e988460400151613112565b612ea58560600151613112565b612eb28660800151613112565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090506115ad565b600060e090910152565b600080612f286142d1565b612f306142d1565b600060e08201819052612f4387876138f7565b84529650905080612f5d5750600093508492509050612703565b612f6787876135da565b60208501529650905080612f845750600093508492509050612703565b612f8e87876135da565b60408501529650905080612fab5750600093508492509050612703565b612fb587876125cc565b60608501529650905080612fd25750600093508492509050612703565b612fdc87876125cc565b60808501529650905080612ff95750600093508492509050612703565b61300387876133c0565b60a085015296509050806130205750600093508492509050612703565b61302a87876138f7565b60c085015296509050806130475750600093508492509050612703565b506001969495509392505050565b61305d6142d1565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b60008060006130cf8460ff1661394b565b50949350505050565b6130e061429d565b611e6660405180608001604052808560ff1681526020018481526020016000151581526020016000801b815250613573565b606081015160009060ff1661313357815161312c90613e6a565b90506115ad565b606082015160ff166001141561316657602080830151805160408201516060830151929093015161312c93919290613e8e565b606082015160ff166002141561317f5761312c82613f36565b600360ff16826060015160ff16101580156131a357506060820151600c60ff909116105b156131b15761312c82613f9c565b606082015160ff16606414156131c9575080516115ad565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b61321261429d565b61209560405180608001604052808660ff16815260200185815260200160011515815260200184815250613573565b6060015160ff161590565b612732826020015161325d83612ad3565b61373d565b61326a61429d565b811561327a5761312c6001612ad3565b61312c6000612ad3565b61328c61429d565b816060015160ff16600214156132d35760405162461bcd60e51b81526004018080602001828103825260218152602001806143c86021913960400191505060405180910390fd5b606082015160ff166132e95761312c6000612ad3565b816060015160ff16600114156133035761312c6001612ad3565b61312c6003612ad3565b6060015160ff1660011490565b60408051600080825260208201909252613335816001613fba565b91505090565b61334361429d565b6040805160a08101825283815281516080810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916133a8565b61339561429d565b81526020019060019003908161338d5790505b50815260646020820152600160409091015292915050565b60008060008085519050848110806133da57506020858203105b156133ef575060009250839150829050612703565b600160208601613405888863ffffffff613fd916565b935093509350509250925092565b60008061341e61432f565b6000849050600086828151811061343157fe5b602001015160f81c60f81b60f81c90508180600101925050600087838151811061345757fe5b016020015160019384019360f89190911c915060009060ff841614156134dd57600061348161429d565b61348b8b876125cc565b9097509092509050816134cf575050604080516080810182526000808252602082018190529181018290526060810182905290975088965094506127039350505050565b6134d881613112565b925050505b60006134ef8a8663ffffffff613fd916565b90506020850194508360ff166001141561353b576040805160808101825260ff909416845260208401919091526001908301819052606083019190915295509193509091506127039050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b61357b61429d565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906135c2565b6135af61429d565b8152602001906001900390816135a75790505b50815260016020820181905260409091015292915050565b6000806135e561429d565b6135ed61429d565b855160009081908781108061360457506040888203105b1561361c576000888596509650965050505050612703565b600061362e8a8a63ffffffff613fd916565b90506020890198506136408a8a6133c0565b909a5094509250821561366b5761365781856101f6565b600198508997509550612703945050505050565b600089869750975097505050505050612703565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156136ca57816020015b6136b761429d565b8152602001906001900390816136af5790505b50905060005b8960ff168160ff161015613727576136e889856125cc565b8451859060ff86169081106136f957fe5b6020908102919091010152945092508261371f5750600095508694509250613734915050565b6001016136d0565b5060019550919350909150505b93509350939050565b61374561429d565b6040805160028082526060828101909352816020015b61376361429d565b81526020019060019003908161375b579050509050828160008151811061378657fe5b6020026020010181905250838160018151811061379f57fe5b60200260200101819052506120956137b682612b58565b613ff5565b6000610276826060015161406b565b60006102768260600151614089565b606081015160009060ff166137f0575060016115ad565b606082015160ff1660011415613808575060006115ad565b606082015160ff166002141561385c576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff161015801561388057506060820151600c60ff909116105b156138d85760408201515160005b818110156138cd576138b6846040015182815181106138a957fe5b60200260200101516137d9565b6138c5576000925050506115ad565b60010161388e565b5060019150506115ad565b606082015160ff16606414156131c9575060006115ad565b6008101590565b6000806000806000865190508581108061391357506020868203105b156139275750600093508492509050612703565b613937878763ffffffff613fd916565b600195506020870194509250612703915050565b60008060018314156139635750600290506001613e65565b60028314156139785750600290506001613e65565b600383141561398d5750600290506001613e65565b60048314156139a25750600290506001613e65565b60058314156139b75750600290506001613e65565b60068314156139cc5750600290506001613e65565b60078314156139e15750600290506001613e65565b60088314156139f65750600390506001613e65565b6009831415613a0b5750600390506001613e65565b600a831415613a205750600290506001613e65565b6010831415613a355750600290506001613e65565b6011831415613a4a5750600290506001613e65565b6012831415613a5f5750600290506001613e65565b6013831415613a745750600290506001613e65565b6014831415613a895750600290506001613e65565b6015831415613a9d57506001905080613e65565b6016831415613ab25750600290506001613e65565b6017831415613ac75750600290506001613e65565b6018831415613adc5750600290506001613e65565b6019831415613af057506001905080613e65565b601a831415613b055750600290506001613e65565b601b831415613b1a5750600290506001613e65565b6020831415613b2e57506001905080613e65565b6021831415613b4257506001905080613e65565b6022831415613b575750600290506001613e65565b6030831415613b6c5750600190506000613e65565b6031831415613b815750600090506001613e65565b6032831415613b965750600090506001613e65565b6033831415613bab5750600190506000613e65565b6034831415613bc05750600190506000613e65565b6035831415613bd55750600290506000613e65565b6036831415613bea5750600090506001613e65565b6037831415613bff5750600090506001613e65565b6038831415613c145750600190506000613e65565b6039831415613c295750600090506001613e65565b603a831415613c3e5750600090506001613e65565b603b831415613c5257506000905080613e65565b603c831415613c675750600090506001613e65565b603d831415613c7c5750600190506000613e65565b6040831415613c915750600190506002613e65565b6041831415613ca65750600290506003613e65565b6042831415613cbb5750600390506004613e65565b6043831415613ccf57506002905080613e65565b6044831415613ce357506003905080613e65565b6050831415613cf85750600290506001613e65565b6051831415613d0d5750600390506001613e65565b6052831415613d2157506001905080613e65565b6053831415613d3557506001905080613e65565b6054831415613d4a5750600290506001613e65565b6060831415613d5e57506000905080613e65565b6061831415613d735750600190506000613e65565b6070831415613d885750600190506000613e65565b6071831415613d9d5750600090506001613e65565b6072831415613db157506001905080613e65565b6073831415613dc557506000905080613e65565b6074831415613dd957506000905080613e65565b6075831415613dee5750600190506000613e65565b6076831415613e035750600090506001613e65565b6080831415613e185750600490506001613e65565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613ee8575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120612095565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214613f8b576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161027691906140ac565b6000613fa661429d565b613faf83613ff5565b9050611e6681613f36565b6000613fc461429d565b613fce84846140e6565b905061209581613f36565b60008160200183511015613fec57600080fd5b50016020015190565b613ffd61429d565b614006826137bb565b61404c576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b606061405b8360400151614105565b9050611e668184608001516140e6565b6000600c60ff8316108015610276575050600360ff91909116101590565b60006140948261406b565b156140a4575060021981016115ad565b5060016115ad565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b6140ee61429d565b60006140f9846141dd565b905061209581846101f6565b6060600882511115614155576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015614182578160200160208202803883390190505b50805190915060005b818110156141d45760006141b18683815181106141a457fe5b6020026020010151613112565b9050808483815181106141c057fe5b60209081029190910101525060010161418b565b50909392505050565b600060088251111561422d576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015614271578181015183820152602001614259565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a00160405280600081526020016142b761432f565b815260606020820181905260006040830181905291015290565b60408051610100810190915260008152602081016142ed61429d565b81526020016142fa61429d565b815260200161430761429d565b815260200161431461429d565b81526000602082018190526040820181905260609091015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820059846e375fc4a32121f72a17f8b854669083db27d0e2fffd59f42de5d6ecb0264736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205ab71856940dee726040a5dde5c092cb823b6db5d358aae393772fe1fcfefa2764736f6c63430005110032"

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
