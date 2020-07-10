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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209c2730224f720457174c3cb5bdf89efcbcb3d762d610b29ebff0fe729129ea2364736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b50614199806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e987d88714610030575b600080fd5b61015260048036036101e081101561004757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100dd57600080fd5b8201836020820111156100ef57600080fd5b8035906020019184600183028401116401000000008311171561011157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610164945050505050565b60408051918252519081900360200190f35b600061017a8d8d8d8d8d8d8d8d8d8d8d8d61018b565b9d9c50505050505050505050505050565b600061017a6040518061016001604052808f81526020018e81526020016101b28e8e6101f6565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610283565b6101fe613f6d565b6040805160a08082018352858252825190810183526000808252602082810182905282850182905260608301829052608083018290528084019290925283518181529182018452919283019161026a565b610257613f6d565b81526020019060019003908161024f5790505b5081526002602082015260400183905290505b92915050565b60008060008060006060610295613fa1565b61029d613fa1565b6102a689611617565b6101208f0151959c50939a509297509095509350915060019060009067ffffffffffffffff168814610316576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8a60800151801561032a575060ff89166072145b8061034657508a60800151158015610346575060ff8916607214155b610397576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a080840180518a900390528401518811156103be5760001960a08401526000915061146a565b60ff891660011415610404576103fd83866000815181106103db57fe5b6020026020010151876001815181106103f057fe5b6020026020010151611971565b915061146a565b60ff891660021415610443576103fd838660008151811061042157fe5b60200260200101518760018151811061043657fe5b60200260200101516119c1565b60ff891660031415610482576103fd838660008151811061046057fe5b60200260200101518760018151811061047557fe5b6020026020010151611a02565b60ff8916600414156104c1576103fd838660008151811061049f57fe5b6020026020010151876001815181106104b457fe5b6020026020010151611a43565b60ff891660051415610500576103fd83866000815181106104de57fe5b6020026020010151876001815181106104f357fe5b6020026020010151611a94565b60ff89166006141561053f576103fd838660008151811061051d57fe5b60200260200101518760018151811061053257fe5b6020026020010151611ae5565b60ff89166007141561057e576103fd838660008151811061055c57fe5b60200260200101518760018151811061057157fe5b6020026020010151611b36565b60ff8916600814156105d2576103fd838660008151811061059b57fe5b6020026020010151876001815181106105b057fe5b6020026020010151886002815181106105c557fe5b6020026020010151611b87565b60ff891660091415610626576103fd83866000815181106105ef57fe5b60200260200101518760018151811061060457fe5b60200260200101518860028151811061061957fe5b6020026020010151611bf1565b60ff8916600a1415610665576103fd838660008151811061064357fe5b60200260200101518760018151811061065857fe5b6020026020010151611c4a565b60ff8916601014156106a4576103fd838660008151811061068257fe5b60200260200101518760018151811061069757fe5b6020026020010151611c8b565b60ff8916601114156106e3576103fd83866000815181106106c157fe5b6020026020010151876001815181106106d657fe5b6020026020010151611ccc565b60ff891660121415610722576103fd838660008151811061070057fe5b60200260200101518760018151811061071557fe5b6020026020010151611d0d565b60ff891660131415610761576103fd838660008151811061073f57fe5b60200260200101518760018151811061075457fe5b6020026020010151611d4e565b60ff8916601414156107a0576103fd838660008151811061077e57fe5b60200260200101518760018151811061079357fe5b6020026020010151611d8f565b60ff8916601514156107ca576103fd83866000815181106107bd57fe5b6020026020010151611db9565b60ff891660161415610809576103fd83866000815181106107e757fe5b6020026020010151876001815181106107fc57fe5b6020026020010151611dfe565b60ff891660171415610848576103fd838660008151811061082657fe5b60200260200101518760018151811061083b57fe5b6020026020010151611e3f565b60ff891660181415610887576103fd838660008151811061086557fe5b60200260200101518760018151811061087a57fe5b6020026020010151611e80565b60ff8916601914156108b1576103fd83866000815181106108a457fe5b6020026020010151611ec1565b60ff8916601a14156108f0576103fd83866000815181106108ce57fe5b6020026020010151876001815181106108e357fe5b6020026020010151611ef7565b60ff8916601b141561092f576103fd838660008151811061090d57fe5b60200260200101518760018151811061092257fe5b6020026020010151611f38565b60ff891660201415610959576103fd838660008151811061094c57fe5b6020026020010151611f79565b60ff891660211415610983576103fd838660008151811061097657fe5b6020026020010151611f94565b60ff8916602214156109c2576103fd83866000815181106109a057fe5b6020026020010151876001815181106109b557fe5b6020026020010151611faf565b60ff8916603014156109ec576103fd83866000815181106109df57fe5b6020026020010151612015565b60ff891660311415610a01576103fd8361201d565b60ff891660321415610a16576103fd8361203e565b60ff891660331415610a40576103fd8386600081518110610a3357fe5b6020026020010151612057565b60ff891660341415610a6a576103fd8386600081518110610a5d57fe5b6020026020010151612063565b60ff891660351415610aa9576103fd8386600081518110610a8757fe5b602002602001015187600181518110610a9c57fe5b602002602001015161208e565b60ff891660361415610abe576103fd836120d6565b60ff891660371415610ad8576103fd838560000151612100565b60ff891660381415610b02576103fd8386600081518110610af557fe5b6020026020010151612110565b60ff891660391415610b8e57610b16613f6d565b610b258c610140015188612122565b9199509750905087610b685760405162461bcd60e51b81526004018080602001828103825260218152602001806141446021913960400191505060405180910390fd5b610b78858263ffffffff61226016565b610b88848263ffffffff61227a16565b5061146a565b60ff8916603a1415610ba3576103fd83612294565b60ff8916603b1415610bb8576001915061146a565b60ff8916603c1415610bcd576103fd836122b1565b60ff8916603d1415610bf7576103fd8386600081518110610bea57fe5b60200260200101516122c5565b60ff891660401415610c21576103fd8386600081518110610c1457fe5b60200260200101516122f3565b60ff891660411415610c60576103fd8386600081518110610c3e57fe5b602002602001015187600181518110610c5357fe5b6020026020010151612315565b60ff891660421415610cb4576103fd8386600081518110610c7d57fe5b602002602001015187600181518110610c9257fe5b602002602001015188600281518110610ca757fe5b6020026020010151612347565b60ff891660431415610cf3576103fd8386600081518110610cd157fe5b602002602001015187600181518110610ce657fe5b6020026020010151612389565b60ff891660441415610d47576103fd8386600081518110610d1057fe5b602002602001015187600181518110610d2557fe5b602002602001015188600281518110610d3a57fe5b602002602001015161239b565b60ff891660501415610d86576103fd8386600081518110610d6457fe5b602002602001015187600181518110610d7957fe5b60200260200101516123bd565b60ff891660511415610dda576103fd8386600081518110610da357fe5b602002602001015187600181518110610db857fe5b602002602001015188600281518110610dcd57fe5b6020026020010151612433565b60ff891660521415610e04576103fd8386600081518110610df757fe5b60200260200101516124c0565b60ff891660531415610ea157610e18613f6d565b610e278c610140015188612122565b9199509750905087610e6a5760405162461bcd60e51b81526004018080602001828103825260218152602001806141446021913960400191505060405180910390fd5b610e7a858263ffffffff61226016565b610e998487600081518110610e8b57fe5b6020026020010151836124f3565b92505061146a565b60ff891660541415610f4b57610eb5613f6d565b610ec48c610140015188612122565b9199509750905087610f075760405162461bcd60e51b81526004018080602001828103825260218152602001806141446021913960400191505060405180910390fd5b610f17858263ffffffff61226016565b610e998487600081518110610f2857fe5b602002602001015188600181518110610f3d57fe5b60200260200101518461254b565b60ff891660601415610f60576103fd836125cc565b60ff89166061141561105e57610f8a8386600081518110610f7d57fe5b60200260200101516125d2565b90925090508115611055578a61010001518b60e00151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001201461100a5760405162461bcd60e51b81526004018080602001828103825260258152602001806140f86025913960400191505060405180910390fd5b8a60c001518b60a00151146110505760405162461bcd60e51b815260040180806020018281038252602781526020018061411d6027913960400191505060405180910390fd5b611059565b5060005b61146a565b60ff89166070141561119e57611088838660008151811061107b57fe5b60200260200101516125ec565b9092509050811561105557806110e3578a60c001518b60a00151146110de5760405162461bcd60e51b81526004018080602001828103825260388152602001806140c06038913960400191505060405180910390fd5b611050565b8a60c001518b60a0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146111575760405162461bcd60e51b81526004018080602001828103825260298152602001806140506029913960400191505060405180910390fd5b8a61010001518b60e00151146110505760405162461bcd60e51b81526004018080602001828103825260268152602001806140796026913960400191505060405180910390fd5b60ff8916607114156112bc5760408051600480825260a08201909252606091816020015b6111ca613f6d565b8152602001906001900390816111c257505060208d01519091506112079060005b60200201516fffffffffffffffffffffffffffffffff1661262b565b8160008151811061121457fe5b60200260200101819052506112338c602001516001600481106111eb57fe5b8160018151811061124057fe5b602002602001018190525061125f8c602001516002600481106111eb57fe5b8160028151811061126c57fe5b602002602001018190525061128b8c602001516003600481106111eb57fe5b8160038151811061129857fe5b6020026020010181905250610b886112af826126b7565b859063ffffffff61227a16565b60ff8916607214156112d6576103fd838c604001516127ae565b60ff8916607314156112eb576000915061146a565b60ff8916607414156113005761105983612814565b60ff89166075141561132a576103fd838660008151811061131d57fe5b602002602001015161281e565b60ff89166076141561133f576103fd83612843565b60ff891660771415611354576103fd8361285c565b60ff891660781415611393576103fd838660008151811061137157fe5b60200260200101518760018151811061138657fe5b60200260200101516128a5565b60ff8916607914156113e7576103fd83866000815181106113b057fe5b6020026020010151876001815181106113c557fe5b6020026020010151886002815181106113da57fe5b60200260200101516128ea565b60ff8916607b14156113fc576103fd8361293d565b60ff891660801415611465576103fd838660008151811061141957fe5b60200260200101518760018151811061142e57fe5b60200260200101518860028151811061144357fe5b60200260200101518960038151811061145857fe5b6020026020010151612980565b600091505b806114fc578a60c001518b60a00151146114b55760405162461bcd60e51b815260040180806020018281038252602781526020018061411d6027913960400191505060405180910390fd5b8a61010001518b60e00151146114fc5760405162461bcd60e51b81526004018080602001828103825260268152602001806140796026913960400191505060405180910390fd5b8161155d5760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156115555761155083612a9f565b61155d565b60c083015183525b61156684612aa9565b8b51146115a45760405162461bcd60e51b815260040180806020018281038252602281526020018061402e6022913960400191505060405180910390fd5b6115ad83612aa9565b8b6060015114611604576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b600099505050505050505050505b919050565b6000806060611624613fa1565b61162c613fa1565b60008061163884612b6d565b61164788610140015183612b77565b9550925090508061169f576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b6116a884612caf565b9250600088610140015183815181106116bd57fe5b602001015160f81c60f81b60f81c905088610140015183600101815181106116e157fe5b016020015160f81c975060006116f689612d18565b60408051838152602080850282010190915290995090915081801561173557816020015b611722613f6d565b81526020019060019003908161171a5790505b5096506002840193508160ff166000148061175357508160ff166001145b6117a4576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff82166117c9576117c26117bd8a886000015161328a565b6132cb565b8652611884565b6117d1613f6d565b6117e08b610140015186612122565b90965090945090508361183a576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b811561185e57808860008151811061184e57fe5b602002602001018190525061186e565b61186e868263ffffffff61227a16565b6118806117bd8b8960000151846133c3565b8752505b60ff82165b81811015611917576118a08b610140015186612122565b8a518b90859081106118ae57fe5b6020908102919091010152955093508361190f576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611889565b875115611964575060005b8260ff168851038110156119645761195c888260018b5103038151811061194557fe5b60200260200101518861227a90919063ffffffff16565b600101611922565b5050505091939550919395565b600061197c8361340c565b158061198e575061198c8261340c565b155b1561199b575060006119ba565b825182518082016119b2878263ffffffff61341716565b600193505050505b9392505050565b60006119cc8361340c565b15806119de57506119dc8261340c565b155b156119eb575060006119ba565b825182518082026119b2878263ffffffff61341716565b6000611a0d8361340c565b1580611a1f5750611a1d8261340c565b155b15611a2c575060006119ba565b825182518082036119b2878263ffffffff61341716565b6000611a4e8361340c565b1580611a605750611a5e8261340c565b155b15611a6d575060006119ba565b8251825180611a81576000925050506119ba565b8082046119b2878263ffffffff61341716565b6000611a9f8361340c565b1580611ab15750611aaf8261340c565b155b15611abe575060006119ba565b8251825180611ad2576000925050506119ba565b8082056119b2878263ffffffff61341716565b6000611af08361340c565b1580611b025750611b008261340c565b155b15611b0f575060006119ba565b8251825180611b23576000925050506119ba565b8082066119b2878263ffffffff61341716565b6000611b418361340c565b1580611b535750611b518261340c565b155b15611b60575060006119ba565b8251825180611b74576000925050506119ba565b8082076119b2878263ffffffff61341716565b6000611b928461340c565b1580611ba45750611ba28361340c565b155b15611bb157506000611be9565b83518351835180611bc85760009350505050611be9565b6000818385089050611be0898263ffffffff61341716565b60019450505050505b949350505050565b6000611bfc8461340c565b1580611c0e5750611c0c8361340c565b155b15611c1b57506000611be9565b83518351835180611c325760009350505050611be9565b6000818385099050611be0898263ffffffff61341716565b6000611c558361340c565b1580611c675750611c658261340c565b155b15611c74575060006119ba565b8251825180820a6119b2878263ffffffff61341716565b6000611c968361340c565b1580611ca85750611ca68261340c565b155b15611cb5575060006119ba565b825182518082106119b2878263ffffffff61341716565b6000611cd78361340c565b1580611ce95750611ce78261340c565b155b15611cf6575060006119ba565b825182518082116119b2878263ffffffff61341716565b6000611d188361340c565b1580611d2a5750611d288261340c565b155b15611d37575060006119ba565b825182518082126119b2878263ffffffff61341716565b6000611d598361340c565b1580611d6b5750611d698261340c565b155b15611d78575060006119ba565b825182518082136119b2878263ffffffff61341716565b6000611daf6112af611da0846132cb565b611da9866132cb565b1461342d565b5060019392505050565b6000611dc48261340c565b611dde57611dd983600063ffffffff61341716565b611df5565b81518015611df2858263ffffffff61341716565b50505b50600192915050565b6000611e098361340c565b1580611e1b5750611e198261340c565b155b15611e28575060006119ba565b825182518082166119b2878263ffffffff61341716565b6000611e4a8361340c565b1580611e5c5750611e5a8261340c565b155b15611e69575060006119ba565b825182518082176119b2878263ffffffff61341716565b6000611e8b8361340c565b1580611e9d5750611e9b8261340c565b155b15611eaa575060006119ba565b825182518082186119b2878263ffffffff61341716565b6000611ecc8261340c565b611ed85750600061027d565b81518019611eec858263ffffffff61341716565b506001949350505050565b6000611f028361340c565b1580611f145750611f128261340c565b155b15611f21575060006119ba565b8251825181811a6119b2878263ffffffff61341716565b6000611f438361340c565b1580611f555750611f538261340c565b155b15611f62575060006119ba565b8251825181810b6119b2878263ffffffff61341716565b6000611df5611f87836132cb565b849063ffffffff61341716565b6000611df5611fa28361344f565b849063ffffffff61227a16565b6000611fba8361340c565b1580611fcc5750611fca8261340c565b155b15611fd9575060006119ba565b82518251604080516020808201859052818301849052825180830384018152606090920190925280519101206119b2878263ffffffff61341716565b600192915050565b600061203682608001518361227a90919063ffffffff16565b506001919050565b600061203682606001518361227a90919063ffffffff16565b60609190910152600190565b600061206e826134d8565b61207a5750600061027d565b612083826132cb565b835250600192915050565b6000612099836134d8565b6120a5575060006119ba565b6120ae8261340c565b6120ba575060006119ba565b815115611daf576120ca836132cb565b84525060019392505050565b60006120366120f36120e66134e5565b611da985602001516132cb565b839063ffffffff61227a16565b6000611df5611fa2836001613506565b6000611df5838363ffffffff61226016565b60008061212d613f6d565b8451841061214d57600084612142600061262b565b925092509250612259565b600080859050600087828151811061216157fe5b016020015160019092019160f81c9050600061217b613fff565b60ff83166121af5761218d8a85613591565b91965094509150848461219f8461262b565b9750975097505050505050612259565b60ff8316600114156121d7576121c58a856135e4565b91965094509050848461219f83613762565b60ff8316600214156121ed5761219f8a856137c9565b600360ff8416108015906122045750600c60ff8416105b1561223f576002198301606061221b828d8861386e565b91985096509050868661222d836126b7565b99509950995050505050505050612259565b60008061224c600061262b565b9199509750955050505050505b9250925092565b61226e82604001518261392c565b82604001819052505050565b61228882602001518261392c565b82602001819052505050565b60006120366120f36122a46134e5565b611da985604001516132cb565b60006120366120f38360c001516001613506565b60006122d0826134d8565b6122dc5750600061027d565b6122e5826132cb565b60c084015250600192915050565b6000612305838363ffffffff61227a16565b611df5838363ffffffff61227a16565b6000612327848363ffffffff61227a16565b612337848463ffffffff61227a16565b611daf848363ffffffff61227a16565b6000612359858363ffffffff61227a16565b612369858463ffffffff61227a16565b612379858563ffffffff61227a16565b611eec858363ffffffff61227a16565b6000612337848463ffffffff61227a16565b60006123ad858563ffffffff61227a16565b612379858463ffffffff61227a16565b60006123c88361340c565b15806123da57506123d8826139aa565b155b156123e7575060006119ba565b6123f0826139b9565b60ff16836000015110612405575060006119ba565b611daf826040015184600001518151811061241c57fe5b60200260200101518561227a90919063ffffffff16565b600061243e836139aa565b1580612450575061244e8461340c565b155b1561245d57506000611be9565b612466836139b9565b60ff1684600001511061247b57506000611be9565b60408301518451815184918391811061249057fe5b60200260200101819052506124b46124a7826126b7565b879063ffffffff61227a16565b50600195945050505050565b60006124cb826139aa565b6124d75750600061027d565b611df56124e3836139b9565b849060ff1663ffffffff61341716565b60006124fe8361340c565b1580612510575061250e826139aa565b155b1561251d575060006119ba565b612526826139b9565b60ff1683600001511061253b575060006119ba565b612405848363ffffffff61226016565b6000612556826139aa565b158061256857506125668461340c565b155b1561257557506000611be9565b61257e826139b9565b60ff1684600001511061259357506000611be9565b6040820151845181518591839181106125a857fe5b60200260200101819052506124b46125bf826126b7565b879063ffffffff61226016565b50600190565b60008060016125e0846132cb565b915091505b9250929050565b60008061271083608001511115612608575060009050806125e5565b612611836139c8565b612620575060009050806125e5565b60016125e0846132cb565b612633613f6d565b6040805160a08082018352848252825190810183526000808252602082810182905282850182905260608301829052608083018290528084019290925283518181529182018452919283019161269f565b61268c613f6d565b8152602001906001900390816126845790505b50815260006020820152600160409091015292915050565b6126bf613f6d565b6126c98251613adf565b61271a576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156127515783818151811061273457fe5b60200260200101516080015182019150808060010191505061271f565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b60006127b86134e5565b6127c1836132cb565b1415612305576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b600260e090910152565b60006128298261340c565b6128355750600061027d565b505160a09190910152600190565b60006120368260a001518361341790919063ffffffff16565b60408051600160f81b60208083019190915260006021830181905260228084018290528451808503909101815260429093019093528151910120612036906120f3906001613506565b60006128b08361340c565b6128bc575060006119ba565b6128c5826134d8565b6128d1575060006119ba565b611daf6112af84600001516128e5856132cb565b61328a565b60006128f58461340c565b61290157506000611be9565b61290a826134d8565b61291657506000611be9565b611eec612930856000015161292a856132cb565b866133c3565b869063ffffffff61227a16565b6040805160008082526020820190925260609082612971565b61295e613f6d565b8152602001906001900390816129565790505b509050611df5611fa2826126b7565b600061298b8561340c565b158061299d575061299b8461340c565b155b806129ae57506129ac8361340c565b155b806129bf57506129bd8261340c565b155b156129cc57506000612a96565b845184518451158015906129e257508451600114155b15612a03576129f888600063ffffffff61341716565b600192505050612a96565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612a65573d6000803e3d6000fd5b5050604051601f1901519150612a8c90508b6001600160a01b03831663ffffffff61341716565b6001955050505050505b95945050505050565b600160e090910152565b600060028260e001511415612ac057506000611612565b60018260e001511415612ad557506001611612565b81516020830151612ae5906132cb565b612af284604001516132cb565b612aff85606001516132cb565b612b0c86608001516132cb565b8660a001518760c0015160405160200180888152602001878152602001868152602001858152602001848152602001838152602001828152602001975050505050505050604051602081830303815290604052805190602001209050611612565b600060e090910152565b600080612b82613fa1565b612b8a613fa1565b600060e08201819052612b9d8787613ae6565b84529650905080612bb75750600093508492509050612259565b612bc187876137c9565b60208501529650905080612bde5750600093508492509050612259565b612be887876137c9565b60408501529650905080612c055750600093508492509050612259565b612c0f8787612122565b60608501529650905080612c2c5750600093508492509050612259565b612c368787612122565b60808501529650905080612c535750600093508492509050612259565b612c5d8787613591565b60a08501529650905080612c7a5750600093508492509050612259565b612c848787613ae6565b60c08501529650905080612ca15750600093508492509050612259565b506001969495509392505050565b612cb7613fa1565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612d305750600290506003613285565b6002831415612d455750600290506003613285565b6003831415612d5a5750600290506003613285565b6004831415612d6f5750600290506004613285565b6005831415612d845750600290506007613285565b6006831415612d995750600290506004613285565b6007831415612dae5750600290506007613285565b6008831415612dc35750600390506004613285565b6009831415612dd85750600390506004613285565b600a831415612ded5750600290506019613285565b6010831415612e0157506002905080613285565b6011831415612e1557506002905080613285565b6012831415612e2957506002905080613285565b6013831415612e3d57506002905080613285565b6014831415612e5157506002905080613285565b6015831415612e6557506001905080613285565b6016831415612e7957506002905080613285565b6017831415612e8d57506002905080613285565b6018831415612ea157506002905080613285565b6019831415612eb557506001905080613285565b601a831415612eca5750600290506004613285565b601b831415612edf5750600290506007613285565b6020831415612ef45750600190506007613285565b6021831415612f095750600190506003613285565b6022831415612f1e5750600290506008613285565b6030831415612f3257506001905080613285565b6031831415612f475750600090506001613285565b6032831415612f5c5750600090506001613285565b6033831415612f715750600190506002613285565b6034831415612f865750600190506004613285565b6035831415612f9b5750600290506004613285565b6036831415612fb05750600090506002613285565b6037831415612fc55750600090506001613285565b6038831415612fd957506001905080613285565b6039831415612fee5750600090506001613285565b603a8314156130035750600090506002613285565b603b8314156130185750600090506001613285565b603c83141561302d5750600090506001613285565b603d83141561304157506001905080613285565b604083141561305557506001905080613285565b604183141561306a5750600290506001613285565b604283141561307f5750600390506001613285565b60438314156130945750600290506001613285565b60448314156130a95750600390506001613285565b60508314156130bd57506002905080613285565b60518314156130d25750600390506028613285565b60528314156130e75750600190506002613285565b60538314156130fc5750600190506003613285565b60548314156131115750600290506029613285565b60608314156131265750600090506064613285565b606183141561313b5750600190506064613285565b60708314156131505750600190506064613285565b60718314156131655750600090506028613285565b607283141561317a5750600090506028613285565b607383141561318f5750600090506005613285565b60748314156131a4575060009050600a613285565b60758314156131b95750600190506000613285565b60768314156131ce5750600090506001613285565b60778314156131e35750600090506019613285565b60788314156131f85750600290506019613285565b607983141561320d5750600390506019613285565b607b831415613222575060009050600a613285565b6080831415613238575060049050614e20613285565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b613292613f6d565b6119ba6040518060a001604052808560ff1681526020018481526020016000151581526020016000801b81526020016000815250613762565b606081015160009060ff166132ec5781516132e590613b3a565b9050611612565b606082015160ff166001141561331f5760208083015180516040820151606083015192909301516132e593919290613b5e565b606082015160ff1660021415613338576132e582613c06565b600360ff16826060015160ff161015801561335c57506060820151600c60ff909116105b1561336a576132e582613c6c565b606082015160ff166064141561338257508051611612565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6133cb613f6d565b611be96040518060a001604052808660ff1681526020018581526020016001151581526020016133fa856132cb565b81526020018460800151815250613762565b6060015160ff161590565b61228882602001516134288361262b565b61392c565b613435613f6d565b8115613445576132e5600161262b565b6132e5600061262b565b613457613f6d565b816060015160ff166002141561349e5760405162461bcd60e51b815260040180806020018281038252602181526020018061409f6021913960400191505060405180910390fd5b606082015160ff166134b4576132e5600061262b565b816060015160ff16600114156134ce576132e5600161262b565b6132e5600361262b565b6060015160ff1660011490565b60408051600080825260208201909252613500816001613c8a565b91505090565b61350e613f6d565b6040805160a08082018352858252825190810183526000808252602082810182905282850182905260608301829052608083018290528084019290925283518181529182018452919283019161357a565b613567613f6d565b81526020019060019003908161355f5790505b508152606460208201526040019290925250919050565b60008060008085519050848110806135ab57506020858203105b156135c0575060009250839150829050612259565b6001602086016135d6888863ffffffff613ca916565b935093509350509250925092565b6000806135ef613fff565b6000849050600086828151811061360257fe5b602001015160f81c60f81b60f81c90508180600101925050600087838151811061362857fe5b016020015160019384019360f89190911c9150600090819060ff851614156136bf576000613654613f6d565b61365e8c88612122565b9098509092509050816136aa5750506040805160a08101825260008082526020820181905291810182905260608101829052608081018290529098508997509550612259945050505050565b6136b3816132cb565b93508060800151925050505b60006136d18b8763ffffffff613ca916565b90506020860195508460ff1660011415613722576040805160a08101825260ff9095168552602085019190915260019084018190526060840192909252608083015295509193509091506122599050565b6040805160a08101825260ff9095168552602085019190915260009084018190526060840181905260808401525060019650929450925050509250925092565b61376a613f6d565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906137b1565b61379e613f6d565b8152602001906001900390816137965790505b50815260016020820181905260409091015292915050565b6000806137d4613f6d565b6137dc613f6d565b85516000908190878110806137f357506040888203105b1561380b576000888596509650965050505050612259565b600061381d8a8a63ffffffff613ca916565b905060208901985061382f8a8a613591565b909a5094509250821561385a5761384681856101f6565b600198508997509550612259945050505050565b600089869750975097505050505050612259565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156138b957816020015b6138a6613f6d565b81526020019060019003908161389e5790505b50905060005b8960ff168160ff161015613916576138d78985612122565b8451859060ff86169081106138e857fe5b6020908102919091010152945092508261390e5750600095508694509250613923915050565b6001016138bf565b5060019550919350909150505b93509350939050565b613934613f6d565b6040805160028082526060828101909352816020015b613952613f6d565b81526020019060019003908161394a579050509050828160008151811061397557fe5b6020026020010181905250838160018151811061398e57fe5b6020026020010181905250611be96139a5826126b7565b613cc5565b600061027d8260600151613d3b565b600061027d8260600151613d59565b606081015160009060ff166139df57506001611612565b606082015160ff16600114156139f757506000611612565b606082015160ff1660021415613a4b576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff1610158015613a6f57506060820151600c60ff909116105b15613ac75760408201515160005b81811015613abc57613aa584604001518281518110613a9857fe5b60200260200101516139c8565b613ab457600092505050611612565b600101613a7d565b506001915050611612565b606082015160ff166064141561338257506000611612565b6008101590565b60008060008060008651905085811080613b0257506020868203105b15613b165750600093508492509050612259565b613b26878763ffffffff613ca916565b600195506020870194509250612259915050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613bb8575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611be9565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214613c5b576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161027d9190613d7c565b6000613c76613f6d565b613c7f83613cc5565b90506119ba81613c06565b6000613c94613f6d565b613c9e8484613db6565b9050611be981613c06565b60008160200183511015613cbc57600080fd5b50016020015190565b613ccd613f6d565b613cd6826139aa565b613d1c576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060613d2b8360400151613dd5565b90506119ba818460800151613db6565b6000600c60ff831610801561027d575050600360ff91909116101590565b6000613d6482613d3b565b15613d7457506002198101611612565b506001611612565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b613dbe613f6d565b6000613dc984613ead565b9050611be981846101f6565b6060600882511115613e25576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613e52578160200160208202803883390190505b50805190915060005b81811015613ea4576000613e81868381518110613e7457fe5b60200260200101516132cb565b905080848381518110613e9057fe5b602090810291909101015250600101613e5b565b50909392505050565b6000600882511115613efd576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613f41578181015183820152602001613f29565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001613f87613fff565b815260606020820181905260006040830181905291015290565b6040805161010081019091526000815260208101613fbd613f6d565b8152602001613fca613f6d565b8152602001613fd7613f6d565b8152602001613fe4613f6d565b81526000602082018190526040820181905260609091015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820265eeb37c3f236026d06233f6b7d940e2da418e30186f441f85397a325a8f25964736f6c63430005110032"

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
