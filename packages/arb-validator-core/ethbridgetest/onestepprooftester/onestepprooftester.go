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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820fcb304cfbea52487f276b22383dd4c75b7acf2a9445b6b822940114e5856cfd764736f6c63430005110032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b086ea420a75a6d4f5159b8803c9e7460c17bd9a432878dfbc211b21d85831d764736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b506142d8806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e987d88714610030575b600080fd5b61015260048036036101e081101561004757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100dd57600080fd5b8201836020820111156100ef57600080fd5b8035906020019184600183028401116401000000008311171561011157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610164945050505050565b60408051918252519081900360200190f35b600061017a8d8d8d8d8d8d8d8d8d8d8d8d61018b565b9d9c50505050505050505050505050565b600061017a6040518061016001604052808f81526020018e81526020016101b28e8e6101f6565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610283565b6101fe6140ac565b6040805160a08082018352858252825190810183526000808252602082810182905282850182905260608301829052608083018290528084019290925283518181529182018452919283019161026a565b6102576140ac565b81526020019060019003908161024f5790505b5081526002602082015260400183905290505b92915050565b600080600080600060606102956140e0565b61029d6140e0565b6102a68961166d565b6101208f0151959c50939a509297509095509350915060019060009067ffffffffffffffff168814610316576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8a60800151801561032a575060ff89166072145b8061034657508a60800151158015610346575060ff8916607214155b610397576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b878460a0015110156103b45760001960a0840152600091506114b5565b60ff8916600114156103fa576103f383866000815181106103d157fe5b6020026020010151876001815181106103e657fe5b60200260200101516119c7565b91506114b5565b60ff891660021415610439576103f3838660008151811061041757fe5b60200260200101518760018151811061042c57fe5b6020026020010151611a17565b60ff891660031415610478576103f3838660008151811061045657fe5b60200260200101518760018151811061046b57fe5b6020026020010151611a58565b60ff8916600414156104b7576103f3838660008151811061049557fe5b6020026020010151876001815181106104aa57fe5b6020026020010151611a99565b60ff8916600514156104f6576103f383866000815181106104d457fe5b6020026020010151876001815181106104e957fe5b6020026020010151611aea565b60ff891660061415610535576103f3838660008151811061051357fe5b60200260200101518760018151811061052857fe5b6020026020010151611b3b565b60ff891660071415610574576103f3838660008151811061055257fe5b60200260200101518760018151811061056757fe5b6020026020010151611b8c565b60ff8916600814156105c8576103f3838660008151811061059157fe5b6020026020010151876001815181106105a657fe5b6020026020010151886002815181106105bb57fe5b6020026020010151611bdd565b60ff89166009141561061c576103f383866000815181106105e557fe5b6020026020010151876001815181106105fa57fe5b60200260200101518860028151811061060f57fe5b6020026020010151611c47565b60ff8916600a141561065b576103f3838660008151811061063957fe5b60200260200101518760018151811061064e57fe5b6020026020010151611ca0565b60ff89166010141561069a576103f3838660008151811061067857fe5b60200260200101518760018151811061068d57fe5b6020026020010151611ce1565b60ff8916601114156106d9576103f383866000815181106106b757fe5b6020026020010151876001815181106106cc57fe5b6020026020010151611d22565b60ff891660121415610718576103f383866000815181106106f657fe5b60200260200101518760018151811061070b57fe5b6020026020010151611d63565b60ff891660131415610757576103f3838660008151811061073557fe5b60200260200101518760018151811061074a57fe5b6020026020010151611da4565b60ff891660141415610796576103f3838660008151811061077457fe5b60200260200101518760018151811061078957fe5b6020026020010151611de5565b60ff8916601514156107c0576103f383866000815181106107b357fe5b6020026020010151611e0f565b60ff8916601614156107ff576103f383866000815181106107dd57fe5b6020026020010151876001815181106107f257fe5b6020026020010151611e54565b60ff89166017141561083e576103f3838660008151811061081c57fe5b60200260200101518760018151811061083157fe5b6020026020010151611e95565b60ff89166018141561087d576103f3838660008151811061085b57fe5b60200260200101518760018151811061087057fe5b6020026020010151611ed6565b60ff8916601914156108a7576103f3838660008151811061089a57fe5b6020026020010151611f17565b60ff8916601a14156108e6576103f383866000815181106108c457fe5b6020026020010151876001815181106108d957fe5b6020026020010151611f4d565b60ff8916601b1415610925576103f3838660008151811061090357fe5b60200260200101518760018151811061091857fe5b6020026020010151611f8e565b60ff89166020141561094f576103f3838660008151811061094257fe5b6020026020010151611fcf565b60ff891660211415610979576103f3838660008151811061096c57fe5b6020026020010151611fea565b60ff8916602214156109b8576103f3838660008151811061099657fe5b6020026020010151876001815181106109ab57fe5b6020026020010151612005565b60ff8916603014156109e2576103f383866000815181106109d557fe5b602002602001015161206b565b60ff8916603114156109f7576103f383612073565b60ff891660321415610a0c576103f383612094565b60ff891660331415610a36576103f38386600081518110610a2957fe5b60200260200101516120ad565b60ff891660341415610a60576103f38386600081518110610a5357fe5b60200260200101516120b9565b60ff891660351415610a9f576103f38386600081518110610a7d57fe5b602002602001015187600181518110610a9257fe5b60200260200101516120e4565b60ff891660361415610ab4576103f38361212c565b60ff891660371415610ace576103f3838560000151612156565b60ff891660381415610af8576103f38386600081518110610aeb57fe5b6020026020010151612166565b60ff891660391415610b8457610b0c6140ac565b610b1b8c610140015188612178565b9199509750905087610b5e5760405162461bcd60e51b81526004018080602001828103825260218152602001806142836021913960400191505060405180910390fd5b610b6e858263ffffffff6122b616565b610b7e848263ffffffff6122d016565b506114b5565b60ff8916603a1415610b99576103f3836122ea565b60ff8916603b1415610bae57600191506114b5565b60ff8916603c1415610bc3576103f383612307565b60ff8916603d1415610bed576103f38386600081518110610be057fe5b602002602001015161231b565b60ff891660401415610c17576103f38386600081518110610c0a57fe5b6020026020010151612349565b60ff891660411415610c56576103f38386600081518110610c3457fe5b602002602001015187600181518110610c4957fe5b602002602001015161236b565b60ff891660421415610caa576103f38386600081518110610c7357fe5b602002602001015187600181518110610c8857fe5b602002602001015188600281518110610c9d57fe5b602002602001015161239d565b60ff891660431415610ce9576103f38386600081518110610cc757fe5b602002602001015187600181518110610cdc57fe5b60200260200101516123df565b60ff891660441415610d3d576103f38386600081518110610d0657fe5b602002602001015187600181518110610d1b57fe5b602002602001015188600281518110610d3057fe5b60200260200101516123f1565b60ff891660501415610d7c576103f38386600081518110610d5a57fe5b602002602001015187600181518110610d6f57fe5b6020026020010151612413565b60ff891660511415610dd0576103f38386600081518110610d9957fe5b602002602001015187600181518110610dae57fe5b602002602001015188600281518110610dc357fe5b6020026020010151612489565b60ff891660521415610dfa576103f38386600081518110610ded57fe5b6020026020010151612516565b60ff891660531415610e9757610e0e6140ac565b610e1d8c610140015188612178565b9199509750905087610e605760405162461bcd60e51b81526004018080602001828103825260218152602001806142836021913960400191505060405180910390fd5b610e70858263ffffffff6122b616565b610e8f8487600081518110610e8157fe5b602002602001015183612549565b9250506114b5565b60ff891660541415610f4157610eab6140ac565b610eba8c610140015188612178565b9199509750905087610efd5760405162461bcd60e51b81526004018080602001828103825260218152602001806142836021913960400191505060405180910390fd5b610f0d858263ffffffff6122b616565b610e8f8487600081518110610f1e57fe5b602002602001015188600181518110610f3357fe5b6020026020010151846125a1565b60ff891660601415610f56576103f383612622565b60ff89166061141561105457610f808386600081518110610f7357fe5b6020026020010151612628565b9092509050811561104b578a61010001518b60e0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146110005760405162461bcd60e51b81526004018080602001828103825260258152602001806142376025913960400191505060405180910390fd5b8a60c001518b60a00151146110465760405162461bcd60e51b815260040180806020018281038252602781526020018061425c6027913960400191505060405180910390fd5b61104f565b5060005b6114b5565b60ff8916607014156111945761107e838660008151811061107157fe5b6020026020010151612642565b9092509050811561104b57806110d9578a60c001518b60a00151146110d45760405162461bcd60e51b81526004018080602001828103825260388152602001806141ff6038913960400191505060405180910390fd5b611046565b8a60c001518b60a00151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001201461114d5760405162461bcd60e51b815260040180806020018281038252602981526020018061418f6029913960400191505060405180910390fd5b8a61010001518b60e00151146110465760405162461bcd60e51b81526004018080602001828103825260268152602001806141b86026913960400191505060405180910390fd5b60ff8916607114156112a95760408051600480825260a08201909252606091816020015b6111c06140ac565b8152602001906001900390816111b857505060208d01519091506111f49060005b60200201516001600160801b0316612681565b8160008151811061120157fe5b60200260200101819052506112208c602001516001600481106111e157fe5b8160018151811061122d57fe5b602002602001018190525061124c8c602001516002600481106111e157fe5b8160028151811061125957fe5b60200260200101819052506112788c602001516003600481106111e157fe5b8160038151811061128557fe5b6020026020010181905250610b7e61129c8261270d565b859063ffffffff6122d016565b60ff8916607214156112f7576103f383866000815181106112c657fe5b60200260200101518d604001518e602001516000600481106112e457fe5b60200201516001600160801b0316612804565b60ff89166073141561130c57600091506114b5565b60ff8916607414156113215761104f8361289b565b60ff89166075141561134b576103f3838660008151811061133e57fe5b60200260200101516128a5565b60ff891660761415611360576103f3836128ca565b60ff891660771415611375576103f3836128e3565b60ff8916607814156113b4576103f3838660008151811061139257fe5b6020026020010151876001815181106113a757fe5b602002602001015161292c565b60ff891660791415611408576103f383866000815181106113d157fe5b6020026020010151876001815181106113e657fe5b6020026020010151886002815181106113fb57fe5b6020026020010151612971565b60ff8916607a1415611432576103f3838660008151811061142557fe5b60200260200101516129c4565b60ff8916607b1415611447576103f383612a67565b60ff8916608014156114b0576103f3838660008151811061146457fe5b60200260200101518760018151811061147957fe5b60200260200101518860028151811061148e57fe5b6020026020010151896003815181106114a357fe5b6020026020010151612aaa565b600091505b80611547578a60c001518b60a00151146115005760405162461bcd60e51b815260040180806020018281038252602781526020018061425c6027913960400191505060405180910390fd5b8a61010001518b60e00151146115475760405162461bcd60e51b81526004018080602001828103825260268152602001806141b86026913960400191505060405180910390fd5b60a0830180518990039052816115b35760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156115ab576115a683612bc9565b6115b3565b60c083015183525b6115bc84612bd3565b8b51146115fa5760405162461bcd60e51b815260040180806020018281038252602281526020018061416d6022913960400191505060405180910390fd5b61160383612bd3565b8b606001511461165a576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b600099505050505050505050505b919050565b600080606061167a6140e0565b6116826140e0565b60008061168e84612c97565b61169d88610140015183612ca1565b955092509050806116f5576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b6116fe84612dd9565b92506000886101400151838151811061171357fe5b602001015160f81c60f81b60f81c9050886101400151836001018151811061173757fe5b016020015160f81c9750600061174c89612e42565b60408051838152602080850282010190915290995090915081801561178b57816020015b6117786140ac565b8152602001906001900390816117705790505b5096506002840193508160ff16600014806117a957508160ff166001145b6117fa576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff821661181f576118186118138a88600001516133c9565b61340a565b86526118da565b6118276140ac565b6118368b610140015186612178565b909650909450905083611890576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b81156118b45780886000815181106118a457fe5b60200260200101819052506118c4565b6118c4868263ffffffff6122d016565b6118d66118138b896000015184613502565b8752505b60ff82165b8181101561196d576118f68b610140015186612178565b8a518b908590811061190457fe5b60209081029190910101529550935083611965576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b6001016118df565b8751156119ba575060005b8260ff168851038110156119ba576119b2888260018b5103038151811061199b57fe5b6020026020010151886122d090919063ffffffff16565b600101611978565b5050505091939550919395565b60006119d28361354b565b15806119e457506119e28261354b565b155b156119f157506000611a10565b82518251808201611a08878263ffffffff61355616565b600193505050505b9392505050565b6000611a228361354b565b1580611a345750611a328261354b565b155b15611a4157506000611a10565b82518251808202611a08878263ffffffff61355616565b6000611a638361354b565b1580611a755750611a738261354b565b155b15611a8257506000611a10565b82518251808203611a08878263ffffffff61355616565b6000611aa48361354b565b1580611ab65750611ab48261354b565b155b15611ac357506000611a10565b8251825180611ad757600092505050611a10565b808204611a08878263ffffffff61355616565b6000611af58361354b565b1580611b075750611b058261354b565b155b15611b1457506000611a10565b8251825180611b2857600092505050611a10565b808205611a08878263ffffffff61355616565b6000611b468361354b565b1580611b585750611b568261354b565b155b15611b6557506000611a10565b8251825180611b7957600092505050611a10565b808206611a08878263ffffffff61355616565b6000611b978361354b565b1580611ba95750611ba78261354b565b155b15611bb657506000611a10565b8251825180611bca57600092505050611a10565b808207611a08878263ffffffff61355616565b6000611be88461354b565b1580611bfa5750611bf88361354b565b155b15611c0757506000611c3f565b83518351835180611c1e5760009350505050611c3f565b6000818385089050611c36898263ffffffff61355616565b60019450505050505b949350505050565b6000611c528461354b565b1580611c645750611c628361354b565b155b15611c7157506000611c3f565b83518351835180611c885760009350505050611c3f565b6000818385099050611c36898263ffffffff61355616565b6000611cab8361354b565b1580611cbd5750611cbb8261354b565b155b15611cca57506000611a10565b8251825180820a611a08878263ffffffff61355616565b6000611cec8361354b565b1580611cfe5750611cfc8261354b565b155b15611d0b57506000611a10565b82518251808210611a08878263ffffffff61355616565b6000611d2d8361354b565b1580611d3f5750611d3d8261354b565b155b15611d4c57506000611a10565b82518251808211611a08878263ffffffff61355616565b6000611d6e8361354b565b1580611d805750611d7e8261354b565b155b15611d8d57506000611a10565b82518251808212611a08878263ffffffff61355616565b6000611daf8361354b565b1580611dc15750611dbf8261354b565b155b15611dce57506000611a10565b82518251808213611a08878263ffffffff61355616565b6000611e0561129c611df68461340a565b611dff8661340a565b1461356c565b5060019392505050565b6000611e1a8261354b565b611e3457611e2f83600063ffffffff61355616565b611e4b565b81518015611e48858263ffffffff61355616565b50505b50600192915050565b6000611e5f8361354b565b1580611e715750611e6f8261354b565b155b15611e7e57506000611a10565b82518251808216611a08878263ffffffff61355616565b6000611ea08361354b565b1580611eb25750611eb08261354b565b155b15611ebf57506000611a10565b82518251808217611a08878263ffffffff61355616565b6000611ee18361354b565b1580611ef35750611ef18261354b565b155b15611f0057506000611a10565b82518251808218611a08878263ffffffff61355616565b6000611f228261354b565b611f2e5750600061027d565b81518019611f42858263ffffffff61355616565b506001949350505050565b6000611f588361354b565b1580611f6a5750611f688261354b565b155b15611f7757506000611a10565b8251825181811a611a08878263ffffffff61355616565b6000611f998361354b565b1580611fab5750611fa98261354b565b155b15611fb857506000611a10565b8251825181810b611a08878263ffffffff61355616565b6000611e4b611fdd8361340a565b849063ffffffff61355616565b6000611e4b611ff88361358e565b849063ffffffff6122d016565b60006120108361354b565b158061202257506120208261354b565b155b1561202f57506000611a10565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120611a08878263ffffffff61355616565b600192915050565b600061208c8260800151836122d090919063ffffffff16565b506001919050565b600061208c8260600151836122d090919063ffffffff16565b60609190910152600190565b60006120c482613617565b6120d05750600061027d565b6120d98261340a565b835250600192915050565b60006120ef83613617565b6120fb57506000611a10565b6121048261354b565b61211057506000611a10565b815115611e05576121208361340a565b84525060019392505050565b600061208c61214961213c613624565b611dff856020015161340a565b839063ffffffff6122d016565b6000611e4b611ff8836001613645565b6000611e4b838363ffffffff6122b616565b6000806121836140ac565b845184106121a3576000846121986000612681565b9250925092506122af565b60008085905060008782815181106121b757fe5b016020015160019092019160f81c905060006121d161413e565b60ff8316612205576121e38a856136d0565b9196509450915084846121f584612681565b97509750975050505050506122af565b60ff83166001141561222d5761221b8a85613723565b9196509450905084846121f5836138a1565b60ff831660021415612243576121f58a85613908565b600360ff84161080159061225a5750600c60ff8416105b156122955760021983016060612271828d886139ad565b9198509650905086866122838361270d565b995099509950505050505050506122af565b6000806122a26000612681565b9199509750955050505050505b9250925092565b6122c4826040015182613a6b565b82604001819052505050565b6122de826020015182613a6b565b82602001819052505050565b600061208c6121496122fa613624565b611dff856040015161340a565b600061208c6121498360c001516001613645565b600061232682613617565b6123325750600061027d565b61233b8261340a565b60c084015250600192915050565b600061235b838363ffffffff6122d016565b611e4b838363ffffffff6122d016565b600061237d848363ffffffff6122d016565b61238d848463ffffffff6122d016565b611e05848363ffffffff6122d016565b60006123af858363ffffffff6122d016565b6123bf858463ffffffff6122d016565b6123cf858563ffffffff6122d016565b611f42858363ffffffff6122d016565b600061238d848463ffffffff6122d016565b6000612403858563ffffffff6122d016565b6123cf858463ffffffff6122d016565b600061241e8361354b565b1580612430575061242e82613ae9565b155b1561243d57506000611a10565b61244682613af8565b60ff1683600001511061245b57506000611a10565b611e05826040015184600001518151811061247257fe5b6020026020010151856122d090919063ffffffff16565b600061249483613ae9565b15806124a657506124a48461354b565b155b156124b357506000611c3f565b6124bc83613af8565b60ff168460000151106124d157506000611c3f565b6040830151845181518491839181106124e657fe5b602002602001018190525061250a6124fd8261270d565b879063ffffffff6122d016565b50600195945050505050565b600061252182613ae9565b61252d5750600061027d565b611e4b61253983613af8565b849060ff1663ffffffff61355616565b60006125548361354b565b1580612566575061256482613ae9565b155b1561257357506000611a10565b61257c82613af8565b60ff1683600001511061259157506000611a10565b61245b848363ffffffff6122b616565b60006125ac82613ae9565b15806125be57506125bc8461354b565b155b156125cb57506000611c3f565b6125d482613af8565b60ff168460000151106125e957506000611c3f565b6040820151845181518591839181106125fe57fe5b602002602001018190525061250a6126158261270d565b879063ffffffff6122b616565b50600190565b60008060016126368461340a565b915091505b9250929050565b6000806127108360800151111561265e5750600090508061263b565b61266783613b07565b6126765750600090508061263b565b60016126368461340a565b6126896140ac565b6040805160a0808201835284825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916126f5565b6126e26140ac565b8152602001906001900390816126da5790505b50815260006020820152600160409091015292915050565b6127156140ac565b61271f8251613c1e565b612770576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156127a75783818151811061278a57fe5b602002602001015160800151820191508080600101915050612775565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b600061280f8461354b565b61281b57506000611c3f565b83518211158061283a575061282e613624565b6128378461340a565b14155b61288b576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b611f42858463ffffffff6122d016565b600260e090910152565b60006128b08261354b565b6128bc5750600061027d565b505160a09190910152600190565b600061208c8260a001518361355690919063ffffffff16565b60408051600160f81b6020808301919091526000602183018190526022808401829052845180850390910181526042909301909352815191012061208c90612149906001613645565b60006129378361354b565b61294357506000611a10565b61294c82613617565b61295857506000611a10565b611e0561129c846000015161296c8561340a565b6133c9565b600061297c8461354b565b61298857506000611c3f565b61299182613617565b61299d57506000611c3f565b611f426129b785600001516129b18561340a565b86613502565b869063ffffffff6122d016565b60006129cf82613617565b6129db5750600061027d565b6129e361413e565b506020820151604081015115612a0d57612a0861129c82606001518360800151613645565b612a51565b60408051600080825260208201909252606091612a40565b612a2d6140ac565b815260200190600190039081612a255790505b509050612a4f6129b78261270d565b505b8051611e0590859060ff1663ffffffff61355616565b6040805160008082526020820190925260609082612a9b565b612a886140ac565b815260200190600190039081612a805790505b509050611e4b611ff88261270d565b6000612ab58561354b565b1580612ac75750612ac58461354b565b155b80612ad85750612ad68361354b565b155b80612ae95750612ae78261354b565b155b15612af657506000612bc0565b84518451845115801590612b0c57508451600114155b15612b2d57612b2288600063ffffffff61355616565b600192505050612bc0565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612b8f573d6000803e3d6000fd5b5050604051601f1901519150612bb690508b6001600160a01b03831663ffffffff61355616565b6001955050505050505b95945050505050565b600160e090910152565b600060028260e001511415612bea57506000611668565b60018260e001511415612bff57506001611668565b81516020830151612c0f9061340a565b612c1c846040015161340a565b612c29856060015161340a565b612c36866080015161340a565b8660a001518760c0015160405160200180888152602001878152602001868152602001858152602001848152602001838152602001828152602001975050505050505050604051602081830303815290604052805190602001209050611668565b600060e090910152565b600080612cac6140e0565b612cb46140e0565b600060e08201819052612cc78787613c25565b84529650905080612ce157506000935084925090506122af565b612ceb8787613908565b60208501529650905080612d0857506000935084925090506122af565b612d128787613908565b60408501529650905080612d2f57506000935084925090506122af565b612d398787612178565b60608501529650905080612d5657506000935084925090506122af565b612d608787612178565b60808501529650905080612d7d57506000935084925090506122af565b612d8787876136d0565b60a08501529650905080612da457506000935084925090506122af565b612dae8787613c25565b60c08501529650905080612dcb57506000935084925090506122af565b506001969495509392505050565b612de16140e0565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612e5a57506002905060036133c4565b6002831415612e6f57506002905060036133c4565b6003831415612e8457506002905060036133c4565b6004831415612e9957506002905060046133c4565b6005831415612eae57506002905060076133c4565b6006831415612ec357506002905060046133c4565b6007831415612ed857506002905060076133c4565b6008831415612eed57506003905060046133c4565b6009831415612f0257506003905060046133c4565b600a831415612f1757506002905060196133c4565b6010831415612f2b575060029050806133c4565b6011831415612f3f575060029050806133c4565b6012831415612f53575060029050806133c4565b6013831415612f67575060029050806133c4565b6014831415612f7b575060029050806133c4565b6015831415612f8f575060019050806133c4565b6016831415612fa3575060029050806133c4565b6017831415612fb7575060029050806133c4565b6018831415612fcb575060029050806133c4565b6019831415612fdf575060019050806133c4565b601a831415612ff457506002905060046133c4565b601b83141561300957506002905060076133c4565b602083141561301e57506001905060076133c4565b602183141561303357506001905060036133c4565b602283141561304857506002905060086133c4565b603083141561305c575060019050806133c4565b603183141561307157506000905060016133c4565b603283141561308657506000905060016133c4565b603383141561309b57506001905060026133c4565b60348314156130b057506001905060046133c4565b60358314156130c557506002905060046133c4565b60368314156130da57506000905060026133c4565b60378314156130ef57506000905060016133c4565b6038831415613103575060019050806133c4565b603983141561311857506000905060016133c4565b603a83141561312d57506000905060026133c4565b603b83141561314257506000905060016133c4565b603c83141561315757506000905060016133c4565b603d83141561316b575060019050806133c4565b604083141561317f575060019050806133c4565b604183141561319457506002905060016133c4565b60428314156131a957506003905060016133c4565b60438314156131be57506002905060016133c4565b60448314156131d357506003905060016133c4565b60508314156131e7575060029050806133c4565b60518314156131fc57506003905060286133c4565b605283141561321157506001905060026133c4565b605383141561322657506001905060036133c4565b605483141561323b57506002905060296133c4565b606083141561325057506000905060646133c4565b606183141561326557506001905060646133c4565b607083141561327a57506001905060646133c4565b607183141561328f57506000905060286133c4565b60728314156132a457506001905060286133c4565b60738314156132b957506000905060056133c4565b60748314156132ce575060009050600a6133c4565b60758314156132e357506001905060006133c4565b60768314156132f857506000905060016133c4565b607783141561330d57506000905060196133c4565b607883141561332257506002905060196133c4565b607983141561333757506003905060196133c4565b607a83141561334c57506001905060196133c4565b607b831415613361575060009050600a6133c4565b6080831415613377575060049050614e206133c4565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b6133d16140ac565b611a106040518060a001604052808560ff1681526020018481526020016000151581526020016000801b815260200160008152506138a1565b606081015160009060ff1661342b57815161342490613c79565b9050611668565b606082015160ff166001141561345e57602080830151805160408201516060830151929093015161342493919290613c9d565b606082015160ff16600214156134775761342482613d45565b600360ff16826060015160ff161015801561349b57506060820151600c60ff909116105b156134a95761342482613dab565b606082015160ff16606414156134c157508051611668565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b61350a6140ac565b611c3f6040518060a001604052808660ff1681526020018581526020016001151581526020016135398561340a565b815260200184608001518152506138a1565b6060015160ff161590565b6122de826020015161356783612681565b613a6b565b6135746140ac565b8115613584576134246001612681565b6134246000612681565b6135966140ac565b816060015160ff16600214156135dd5760405162461bcd60e51b81526004018080602001828103825260218152602001806141de6021913960400191505060405180910390fd5b606082015160ff166135f3576134246000612681565b816060015160ff166001141561360d576134246001612681565b6134246003612681565b6060015160ff1660011490565b6040805160008082526020820190925261363f816001613dc9565b91505090565b61364d6140ac565b6040805160a0808201835285825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916136b9565b6136a66140ac565b81526020019060019003908161369e5790505b508152606460208201526040019290925250919050565b60008060008085519050848110806136ea57506020858203105b156136ff5750600092508391508290506122af565b600160208601613715888863ffffffff613de816565b935093509350509250925092565b60008061372e61413e565b6000849050600086828151811061374157fe5b602001015160f81c60f81b60f81c90508180600101925050600087838151811061376757fe5b016020015160019384019360f89190911c9150600090819060ff851614156137fe5760006137936140ac565b61379d8c88612178565b9098509092509050816137e95750506040805160a081018252600080825260208201819052918101829052606081018290526080810182905290985089975095506122af945050505050565b6137f28161340a565b93508060800151925050505b60006138108b8763ffffffff613de816565b90506020860195508460ff1660011415613861576040805160a08101825260ff9095168552602085019190915260019084018190526060840192909252608083015295509193509091506122af9050565b6040805160a08101825260ff9095168552602085019190915260009084018190526060840181905260808401525060019650929450925050509250925092565b6138a96140ac565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906138f0565b6138dd6140ac565b8152602001906001900390816138d55790505b50815260016020820181905260409091015292915050565b6000806139136140ac565b61391b6140ac565b855160009081908781108061393257506040888203105b1561394a5760008885965096509650505050506122af565b600061395c8a8a63ffffffff613de816565b905060208901985061396e8a8a6136d0565b909a509450925082156139995761398581856101f6565b6001985089975095506122af945050505050565b6000898697509750975050505050506122af565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156139f857816020015b6139e56140ac565b8152602001906001900390816139dd5790505b50905060005b8960ff168160ff161015613a5557613a168985612178565b8451859060ff8616908110613a2757fe5b60209081029190910101529450925082613a4d5750600095508694509250613a62915050565b6001016139fe565b5060019550919350909150505b93509350939050565b613a736140ac565b6040805160028082526060828101909352816020015b613a916140ac565b815260200190600190039081613a895790505090508281600081518110613ab457fe5b60200260200101819052508381600181518110613acd57fe5b6020026020010181905250611c3f613ae48261270d565b613e04565b600061027d8260600151613e7a565b600061027d8260600151613e98565b606081015160009060ff16613b1e57506001611668565b606082015160ff1660011415613b3657506000611668565b606082015160ff1660021415613b8a576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff1610158015613bae57506060820151600c60ff909116105b15613c065760408201515160005b81811015613bfb57613be484604001518281518110613bd757fe5b6020026020010151613b07565b613bf357600092505050611668565b600101613bbc565b506001915050611668565b606082015160ff16606414156134c157506000611668565b6008101590565b60008060008060008651905085811080613c4157506020868203105b15613c5557506000935084925090506122af565b613c65878763ffffffff613de816565b6001955060208701945092506122af915050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613cf7575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611c3f565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214613d9a576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161027d9190613ebb565b6000613db56140ac565b613dbe83613e04565b9050611a1081613d45565b6000613dd36140ac565b613ddd8484613ef5565b9050611c3f81613d45565b60008160200183511015613dfb57600080fd5b50016020015190565b613e0c6140ac565b613e1582613ae9565b613e5b576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060613e6a8360400151613f14565b9050611a10818460800151613ef5565b6000600c60ff831610801561027d575050600360ff91909116101590565b6000613ea382613e7a565b15613eb357506002198101611668565b506001611668565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b613efd6140ac565b6000613f0884613fec565b9050611c3f81846101f6565b6060600882511115613f64576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613f91578160200160208202803883390190505b50805190915060005b81811015613fe3576000613fc0868381518110613fb357fe5b602002602001015161340a565b905080848381518110613fcf57fe5b602090810291909101015250600101613f9a565b50909392505050565b600060088251111561403c576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015614080578181015183820152602001614068565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a00160405280600081526020016140c661413e565b815260606020820181905260006040830181905291015290565b60408051610100810190915260008152602081016140fc6140ac565b81526020016141096140ac565b81526020016141166140ac565b81526020016141236140ac565b81526000602082018190526040820181905260609091015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820213138c92ed5a3245ec7045eaff93844cf208aa33c32627877fd65d90532d5cc64736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209a81e3fc238267b7e44ca5549710241de669e8db2d08f84f90a6b17eb6424fc364736f6c63430005110032"

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
