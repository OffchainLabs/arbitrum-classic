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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158207cf27292a79ee68dd16e8b9a6d7c574887ac61b35185d3f108e5e04e8e6aceea64736f6c63430005110032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d0e4ceaa9f569d33509039a349ca88b315e18d21cf692acc8928ff4e8e34f6d664736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b50614a3a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e987d88714610030575b600080fd5b61015260048036036101e081101561004757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100dd57600080fd5b8201836020820111156100ef57600080fd5b8035906020019184600183028401116401000000008311171561011157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610164945050505050565b60408051918252519081900360200190f35b600061017a8d8d8d8d8d8d8d8d8d8d8d8d61018b565b9d9c50505050505050505050505050565b600061017a6040518061016001604052808f81526020018e81526020016101b28e8e6101f6565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff1681526020018481525061027c565b6101fe6147f1565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610263565b6102506147f1565b8152602001906001900390816102485790505b5081526002602082015260400183905290505b92915050565b600080600080606061028c614825565b610294614825565b61029d88611768565b93995092965090945092509050600160006102b788611ae7565b67ffffffffffffffff168a610120015167ffffffffffffffff161461031a576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8960800151801561032e575060ff88166072145b8061034a5750896080015115801561034a575060ff8816607214155b61039b576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b6103a488611ae7565b67ffffffffffffffff168460a0015110156103ca5760001960a08401526000915061129d565b60ff8816600114156104105761040983866000815181106103e757fe5b6020026020010151876001815181106103fc57fe5b6020026020010151611fab565b915061129d565b60ff88166002141561044f57610409838660008151811061042d57fe5b60200260200101518760018151811061044257fe5b6020026020010151611ffb565b60ff88166003141561048e57610409838660008151811061046c57fe5b60200260200101518760018151811061048157fe5b602002602001015161203c565b60ff8816600414156104cd5761040983866000815181106104ab57fe5b6020026020010151876001815181106104c057fe5b602002602001015161207d565b60ff88166005141561050c5761040983866000815181106104ea57fe5b6020026020010151876001815181106104ff57fe5b60200260200101516120ce565b60ff88166006141561054b57610409838660008151811061052957fe5b60200260200101518760018151811061053e57fe5b602002602001015161211f565b60ff88166007141561058a57610409838660008151811061056857fe5b60200260200101518760018151811061057d57fe5b6020026020010151612170565b60ff8816600814156105de5761040983866000815181106105a757fe5b6020026020010151876001815181106105bc57fe5b6020026020010151886002815181106105d157fe5b60200260200101516121c1565b60ff8816600914156106325761040983866000815181106105fb57fe5b60200260200101518760018151811061061057fe5b60200260200101518860028151811061062557fe5b602002602001015161222b565b60ff8816600a141561067157610409838660008151811061064f57fe5b60200260200101518760018151811061066457fe5b6020026020010151612284565b60ff8816601014156106b057610409838660008151811061068e57fe5b6020026020010151876001815181106106a357fe5b60200260200101516122c5565b60ff8816601114156106ef5761040983866000815181106106cd57fe5b6020026020010151876001815181106106e257fe5b6020026020010151612306565b60ff88166012141561072e57610409838660008151811061070c57fe5b60200260200101518760018151811061072157fe5b6020026020010151612347565b60ff88166013141561076d57610409838660008151811061074b57fe5b60200260200101518760018151811061076057fe5b6020026020010151612388565b60ff8816601414156107ac57610409838660008151811061078a57fe5b60200260200101518760018151811061079f57fe5b60200260200101516123c9565b60ff8816601514156107d65761040983866000815181106107c957fe5b60200260200101516123f3565b60ff8816601614156108155761040983866000815181106107f357fe5b60200260200101518760018151811061080857fe5b6020026020010151612438565b60ff88166017141561085457610409838660008151811061083257fe5b60200260200101518760018151811061084757fe5b6020026020010151612479565b60ff88166018141561089357610409838660008151811061087157fe5b60200260200101518760018151811061088657fe5b60200260200101516124ba565b60ff8816601914156108bd5761040983866000815181106108b057fe5b60200260200101516124fb565b60ff8816601a14156108fc5761040983866000815181106108da57fe5b6020026020010151876001815181106108ef57fe5b6020026020010151612531565b60ff8816601b141561093b57610409838660008151811061091957fe5b60200260200101518760018151811061092e57fe5b6020026020010151612572565b60ff88166020141561096557610409838660008151811061095857fe5b60200260200101516125b3565b60ff88166021141561098f57610409838660008151811061098257fe5b60200260200101516125ce565b60ff8816602214156109ce5761040983866000815181106109ac57fe5b6020026020010151876001815181106109c157fe5b60200260200101516125e9565b60ff8816603014156109f85761040983866000815181106109eb57fe5b602002602001015161264f565b60ff881660311415610a0d5761040983612657565b60ff881660321415610a225761040983612678565b60ff881660331415610a4c576104098386600081518110610a3f57fe5b6020026020010151612691565b60ff881660341415610a76576104098386600081518110610a6957fe5b602002602001015161269d565b60ff881660351415610ab5576104098386600081518110610a9357fe5b602002602001015187600181518110610aa857fe5b60200260200101516126c8565b60ff881660361415610aca5761040983612710565b60ff881660371415610ae45761040983856000015161273a565b60ff881660381415610b0e576104098386600081518110610b0157fe5b6020026020010151612748565b60ff881660391415610b9a57610b226147f1565b610b318b61014001518861275a565b9199509750905087610b745760405162461bcd60e51b81526004018080602001828103825260218152602001806149e56021913960400191505060405180910390fd5b610b84858263ffffffff61289816565b610b94848263ffffffff6128b216565b5061129d565b60ff8816603a1415610baf57610409836128cc565b60ff8816603b1415610bc4576001915061129d565b60ff8816603c1415610bd957610409836128e9565b60ff8816603d1415610c03576104098386600081518110610bf657fe5b60200260200101516128fb565b60ff881660401415610c2d576104098386600081518110610c2057fe5b6020026020010151612929565b60ff881660411415610c6c576104098386600081518110610c4a57fe5b602002602001015187600181518110610c5f57fe5b602002602001015161294b565b60ff881660421415610cc0576104098386600081518110610c8957fe5b602002602001015187600181518110610c9e57fe5b602002602001015188600281518110610cb357fe5b602002602001015161297d565b60ff881660431415610cff576104098386600081518110610cdd57fe5b602002602001015187600181518110610cf257fe5b60200260200101516129bf565b60ff881660441415610d53576104098386600081518110610d1c57fe5b602002602001015187600181518110610d3157fe5b602002602001015188600281518110610d4657fe5b60200260200101516129d1565b60ff881660501415610d92576104098386600081518110610d7057fe5b602002602001015187600181518110610d8557fe5b60200260200101516129f3565b60ff881660511415610de6576104098386600081518110610daf57fe5b602002602001015187600181518110610dc457fe5b602002602001015188600281518110610dd957fe5b6020026020010151612a69565b60ff881660521415610e10576104098386600081518110610e0357fe5b6020026020010151612af6565b60ff881660601415610e255761040983612b29565b60ff881660611415610f2357610e4f8386600081518110610e4257fe5b6020026020010151612b2f565b90925090508115610f1a578961010001518a60e001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610ecf5760405162461bcd60e51b81526004018080602001828103825260258152602001806149996025913960400191505060405180910390fd5b8960c001518a60a0015114610f155760405162461bcd60e51b81526004018080602001828103825260278152602001806149be6027913960400191505060405180910390fd5b610f1e565b5060005b61129d565b60ff88166070141561106357610f4d8386600081518110610f4057fe5b6020026020010151612b49565b90925090508115610f1a5780610fa8578960c001518a60a0015114610fa35760405162461bcd60e51b81526004018080602001828103825260388152602001806149616038913960400191505060405180910390fd5b610f15565b8960c001518a60a00151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001201461101c5760405162461bcd60e51b81526004018080602001828103825260298152602001806148ab6029913960400191505060405180910390fd5b8961010001518a60e0015114610f155760405162461bcd60e51b81526004018080602001828103825260268152602001806148d46026913960400191505060405180910390fd5b60ff8816607114156111785760408051600480825260a08201909252606091816020015b61108f6147f1565b81526020019060019003908161108757505060208c01519091506110c39060005b60200201516001600160801b0316612b88565b816000815181106110d057fe5b60200260200101819052506110ef8b602001516001600481106110b057fe5b816001815181106110fc57fe5b602002602001018190525061111b8b602001516002600481106110b057fe5b8160028151811061112857fe5b60200260200101819052506111478b602001516003600481106110b057fe5b8160038151811061115457fe5b6020026020010181905250610b9461116b82612c0d565b859063ffffffff6128b216565b60ff8816607214156111c657610409838660008151811061119557fe5b60200260200101518c604001518d602001516000600481106111b357fe5b60200201516001600160801b0316612cfc565b60ff8816607314156111db576000915061129d565b60ff8816607414156111f057610f1e83612d93565b60ff88166075141561121a57610409838660008151811061120d57fe5b6020026020010151612d9d565b60ff88166076141561122f5761040983612dc2565b60ff88166080141561129857610409838660008151811061124c57fe5b60200260200101518760018151811061126157fe5b60200260200101518860028151811061127657fe5b60200260200101518960038151811061128b57fe5b6020026020010151612ddb565b600091505b8061132f578960c001518a60a00151146112e85760405162461bcd60e51b81526004018080602001828103825260278152602001806149be6027913960400191505060405180910390fd5b8961010001518a60e001511461132f5760405162461bcd60e51b81526004018080602001828103825260268152602001806148d46026913960400191505060405180910390fd5b61133888611ae7565b67ffffffffffffffff168360a00151038360a0018181525050816113b25760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156113aa576113a583612efa565b6113b2565b60c083015183525b6113bb84612f04565b8a51146113c785612fc8565b8b516113d2906132cc565b6113e36113de88612f04565b6132cc565b6040516020018060246148fa823960240184805190602001908083835b6020831061141f5780518252601f199092019160209182019101611400565b51815160209384036101000a60001901801990921691161790526d0103132b337b932a430b9b4101e960951b919093019081528551600e90910192860191508083835b602083106114815780518252601f199092019160209182019101611462565b51815160209384036101000a60001901801990921691161790526e539ba30b93a26b0b1b434b732901e960851b919093019081528451601090910192850191508083835b602083106114e45780518252601f1990920191602091820191016114c5565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051602081830303815290604052906115a35760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611568578181015183820152602001611550565b50505050905090810190601f1680156115955780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506115ad83612f04565b8a60600151146115bc84612fc8565b6115c98c606001516132cc565b6115d56113de87612f04565b60405160200180602261493f823960220184805190602001908083835b602083106116115780518252601f1990920191602091820191016115f2565b51815160209384036101000a60001901801990921691161790526c01030b33a32b92430b9b4101e9609d1b919093019081528551600d90910192860191508083835b602083106116725780518252601f199092019160209182019101611653565b51815160209384036101000a60001901801990921691161790526c532b73226b0b1b434b732901e960951b919093019081528451600e90910192850191508083835b602083106116d35780518252601f1990920191602091820191016116b4565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051602081830303815290604052906117555760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611568578181015183820152602001611550565b506000985050505050505050505b919050565b60006060611774614825565b61177c614825565b60008080611788614825565b6117918161339b565b6117a0896101400151846133a5565b9094509092509050816117fa576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b611802614825565b61180b826134dd565b905060008a6101400151858151811061182057fe5b602001015160f81c60f81b60f81c905060008b6101400151866001018151811061184657fe5b016020015160f81c9050600061185b82613546565b905060608160405190808252806020026020018201604052801561189957816020015b6118866147f1565b81526020019060019003908161187e5790505b5090506002880197508360ff16600014806118b757508360ff166001145b611908576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff841661192d57611926611921848860000151613560565b61359a565b86526119f0565b6119356147f1565b6119448f61014001518a61275a565b909a5090985090508761199e576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b82156119c25780826000815181106119b257fe5b60200260200101819052506119d2565b6119d2868263ffffffff6128b216565b6119ec6119218589600001516119e78561359a565b613692565b8752505b60ff84165b82811015611a8357611a0c8f61014001518a61275a565b8451859085908110611a1a57fe5b60209081029190910101529950975087611a7b576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b6001016119f5565b815115611ad0575060005b8460ff16825103811015611ad057611ac8828260018551030381518110611ab157fe5b6020026020010151886128b290919063ffffffff16565b600101611a8e565b50919d919c50939a50919850939650945050505050565b600060ff821660011415611afd57506003611763565b60ff821660021415611b1157506003611763565b60ff821660031415611b2557506003611763565b60ff821660041415611b3957506004611763565b60ff821660051415611b4d57506007611763565b60ff821660061415611b6157506004611763565b60ff821660071415611b7557506007611763565b60ff821660081415611b8957506004611763565b60ff821660091415611b9d57506004611763565b60ff8216600a1415611bb157506019611763565b60ff821660101415611bc557506002611763565b60ff821660111415611bd957506002611763565b60ff821660121415611bed57506002611763565b60ff821660131415611c0157506002611763565b60ff821660141415611c1557506002611763565b60ff821660151415611c2957506001611763565b60ff821660161415611c3d57506002611763565b60ff821660171415611c5157506002611763565b60ff821660181415611c6557506002611763565b60ff821660191415611c7957506001611763565b60ff8216601a1415611c8d57506004611763565b60ff8216601b1415611ca157506007611763565b60ff821660201415611cb557506007611763565b60ff821660211415611cc957506003611763565b60ff821660221415611cdd57506008611763565b60ff821660301415611cf157506001611763565b60ff821660311415611d0557506001611763565b60ff821660321415611d1957506001611763565b60ff821660331415611d2d57506002611763565b60ff821660341415611d4157506004611763565b60ff821660351415611d5557506004611763565b60ff821660361415611d6957506002611763565b60ff821660371415611d7d57506001611763565b60ff821660381415611d9157506001611763565b60ff821660391415611da557506001611763565b60ff8216603a1415611db957506002611763565b60ff8216603b1415611dcd57506001611763565b60ff8216603c1415611de157506001611763565b60ff8216603d1415611df557506001611763565b60ff821660401415611e0957506001611763565b60ff821660411415611e1d57506001611763565b60ff821660421415611e3157506001611763565b60ff821660431415611e4557506001611763565b60ff821660441415611e5957506001611763565b60ff821660501415611e6d57506002611763565b60ff821660511415611e8157506028611763565b60ff821660521415611e9557506002611763565b60ff821660601415611ea957506064611763565b60ff821660611415611ebd57506064611763565b60ff821660701415611ed157506064611763565b60ff821660711415611ee557506028611763565b60ff821660721415611ef957506028611763565b60ff821660731415611f0d57506005611763565b60ff821660741415611f215750600a611763565b60ff821660751415611f3557506000611763565b60ff821660761415611f4957506001611763565b60ff821660801415611f5e5750614e20611763565b6040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206f70636f64653a206f70476173436f737428290000000000604482015290519081900360640190fd5b6000611fb6836136c9565b1580611fc85750611fc6826136c9565b155b15611fd557506000611ff4565b82518251808201611fec878263ffffffff6136d416565b600193505050505b9392505050565b6000612006836136c9565b15806120185750612016826136c9565b155b1561202557506000611ff4565b82518251808202611fec878263ffffffff6136d416565b6000612047836136c9565b15806120595750612057826136c9565b155b1561206657506000611ff4565b82518251808203611fec878263ffffffff6136d416565b6000612088836136c9565b158061209a5750612098826136c9565b155b156120a757506000611ff4565b82518251806120bb57600092505050611ff4565b808204611fec878263ffffffff6136d416565b60006120d9836136c9565b15806120eb57506120e9826136c9565b155b156120f857506000611ff4565b825182518061210c57600092505050611ff4565b808205611fec878263ffffffff6136d416565b600061212a836136c9565b158061213c575061213a826136c9565b155b1561214957506000611ff4565b825182518061215d57600092505050611ff4565b808206611fec878263ffffffff6136d416565b600061217b836136c9565b158061218d575061218b826136c9565b155b1561219a57506000611ff4565b82518251806121ae57600092505050611ff4565b808207611fec878263ffffffff6136d416565b60006121cc846136c9565b15806121de57506121dc836136c9565b155b156121eb57506000612223565b835183518351806122025760009350505050612223565b600081838508905061221a898263ffffffff6136d416565b60019450505050505b949350505050565b6000612236846136c9565b15806122485750612246836136c9565b155b1561225557506000612223565b8351835183518061226c5760009350505050612223565b600081838509905061221a898263ffffffff6136d416565b600061228f836136c9565b15806122a1575061229f826136c9565b155b156122ae57506000611ff4565b8251825180820a611fec878263ffffffff6136d416565b60006122d0836136c9565b15806122e257506122e0826136c9565b155b156122ef57506000611ff4565b82518251808210611fec878263ffffffff6136d416565b6000612311836136c9565b15806123235750612321826136c9565b155b1561233057506000611ff4565b82518251808211611fec878263ffffffff6136d416565b6000612352836136c9565b15806123645750612362826136c9565b155b1561237157506000611ff4565b82518251808212611fec878263ffffffff6136d416565b6000612393836136c9565b15806123a557506123a3826136c9565b155b156123b257506000611ff4565b82518251808213611fec878263ffffffff6136d416565b60006123e961116b6123da8461359a565b6123e38661359a565b146136ea565b5060019392505050565b60006123fe826136c9565b6124185761241383600063ffffffff6136d416565b61242f565b8151801561242c858263ffffffff6136d416565b50505b50600192915050565b6000612443836136c9565b15806124555750612453826136c9565b155b1561246257506000611ff4565b82518251808216611fec878263ffffffff6136d416565b6000612484836136c9565b15806124965750612494826136c9565b155b156124a357506000611ff4565b82518251808217611fec878263ffffffff6136d416565b60006124c5836136c9565b15806124d757506124d5826136c9565b155b156124e457506000611ff4565b82518251808218611fec878263ffffffff6136d416565b6000612506826136c9565b61251257506000610276565b81518019612526858263ffffffff6136d416565b506001949350505050565b600061253c836136c9565b158061254e575061254c826136c9565b155b1561255b57506000611ff4565b8251825181811a611fec878263ffffffff6136d416565b600061257d836136c9565b158061258f575061258d826136c9565b155b1561259c57506000611ff4565b8251825181810b611fec878263ffffffff6136d416565b600061242f6125c18361359a565b849063ffffffff6136d416565b600061242f6125dc8361370c565b849063ffffffff6128b216565b60006125f4836136c9565b15806126065750612604826136c9565b155b1561261357506000611ff4565b8251825160408051602080820185905281830184905282518083038401815260609092019092528051910120611fec878263ffffffff6136d416565b600192915050565b60006126708260800151836128b290919063ffffffff16565b506001919050565b60006126708260600151836128b290919063ffffffff16565b60609190910152600190565b60006126a882613795565b6126b457506000610276565b6126bd8261359a565b835250600192915050565b60006126d383613795565b6126df57506000611ff4565b6126e8826136c9565b6126f457506000611ff4565b8151156123e9576127048361359a565b84525060019392505050565b600061267061272d6127206137a2565b6123e3856020015161359a565b839063ffffffff6128b216565b600061242f6125dc836137c3565b600061242f838363ffffffff61289816565b6000806127656147f1565b845184106127855760008461277a6000612b88565b925092509250612891565b600080859050600087828151811061279957fe5b016020015160019092019160f81c905060006127b3614883565b60ff83166127e7576127c58a85613848565b9196509450915084846127d784612b88565b9750975097505050505050612891565b60ff83166001141561280f576127fd8a8561389b565b9196509450905084846127d7836139fb565b60ff831660021415612825576127d78a85613a62565b600360ff84161080159061283c5750600c60ff8416105b156128775760021983016060612853828d88613b07565b91985096509050868661286583612c0d565b99509950995050505050505050612891565b6000806128846000612b88565b9199509750955050505050505b9250925092565b6128a6826040015182613bc5565b82604001819052505050565b6128c0826020015182613bc5565b82602001819052505050565b600061267061272d6128dc6137a2565b6123e3856040015161359a565b600061267061272d8360c001516137c3565b600061290682613795565b61291257506000610276565b61291b8261359a565b60c084015250600192915050565b600061293b838363ffffffff6128b216565b61242f838363ffffffff6128b216565b600061295d848363ffffffff6128b216565b61296d848463ffffffff6128b216565b6123e9848363ffffffff6128b216565b600061298f858363ffffffff6128b216565b61299f858463ffffffff6128b216565b6129af858563ffffffff6128b216565b612526858363ffffffff6128b216565b600061296d848463ffffffff6128b216565b60006129e3858563ffffffff6128b216565b6129af858463ffffffff6128b216565b60006129fe836136c9565b1580612a105750612a0e82613c43565b155b15612a1d57506000611ff4565b612a2682613c52565b60ff16836000015110612a3b57506000611ff4565b6123e98260400151846000015181518110612a5257fe5b6020026020010151856128b290919063ffffffff16565b6000612a7483613c43565b1580612a865750612a84846136c9565b155b15612a9357506000612223565b612a9c83613c52565b60ff16846000015110612ab157506000612223565b604083015184518151849183918110612ac657fe5b6020026020010181905250612aea612add82612c0d565b879063ffffffff6128b216565b50600195945050505050565b6000612b0182613c43565b612b0d57506000610276565b61242f612b1983613c52565b849060ff1663ffffffff6136d416565b50600190565b6000806001612b3d8461359a565b915091505b9250929050565b60008061271083608001511115612b6557506000905080612b42565b612b6e83613c61565b612b7d57506000905080612b42565b6001612b3d8461359a565b612b906147f1565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612bf5565b612be26147f1565b815260200190600190039081612bda5790505b50815260006020820152600160409091015292915050565b612c156147f1565b612c1f8251613d78565b612c70576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015612ca757838181518110612c8a57fe5b602002602001015160800151820191508080600101915050612c75565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b6000612d07846136c9565b612d1357506000612223565b835182111580612d325750612d266137a2565b612d2f8461359a565b14155b612d83576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b612526858463ffffffff6128b216565b600260e090910152565b6000612da8826136c9565b612db457506000610276565b505160a09190910152600190565b60006126708260a00151836136d490919063ffffffff16565b6000612de6856136c9565b1580612df85750612df6846136c9565b155b80612e095750612e07836136c9565b155b80612e1a5750612e18826136c9565b155b15612e2757506000612ef1565b84518451845115801590612e3d57508451600114155b15612e5e57612e5388600063ffffffff6136d416565b600192505050612ef1565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612ec0573d6000803e3d6000fd5b5050604051601f1901519150612ee790508b6001600160a01b03831663ffffffff6136d416565b6001955050505050505b95945050505050565b600160e090910152565b600060028260e001511415612f1b57506000611763565b60018260e001511415612f3057506001611763565b81516020830151612f409061359a565b612f4d846040015161359a565b612f5a856060015161359a565b612f67866080015161359a565b8660a001518760c0015160405160200180888152602001878152602001868152602001858152602001848152602001838152602001828152602001975050505050505050604051602081830303815290604052805190602001209050611763565b6060612fd782600001516132cc565b612fe76113de846020015161359a565b612ff76113de856040015161359a565b6130076113de866060015161359a565b6130176113de876080015161359a565b6130248760a00151613d7f565b6130318860c001516132cc565b60405160200180806709ac2c6d0d2dcca560c31b81525060080188805190602001908083835b602083106130765780518252601f199092019160209182019101613057565b51815160209384036101000a60001901801990921691161790526216100560e91b9190930190815289516003909101928a0191508083835b602083106130cd5780518252601f1990920191602091820191016130ae565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528851600390910192890191508083835b602083106131245780518252601f199092019160209182019101613105565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528751600390910192880191508083835b6020831061317b5780518252601f19909201916020918201910161315c565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528651600390910192870191508083835b602083106131d25780518252601f1990920191602091820191016131b3565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528551600390910192860191508083835b602083106132295780518252601f19909201916020918201910161320a565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528451600390910192850191508083835b602083106132805780518252601f199092019160209182019101613261565b5181516020939093036101000a600019018019909116921691909117905261148560f11b92019182525060408051808303601d19018152600290920190529a9950505050505050505050565b60408051818152606081810183529182919060208201818038833901905050905060005b602081101561339457600084826020811061330757fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b61332d82613e43565b85856002028151811061333c57fe5b60200101906001600160f81b031916908160001a90535061335c81613e43565b85856002026001018151811061336e57fe5b60200101906001600160f81b031916908160001a90535050600190920191506132f09050565b5092915050565b600060e090910152565b6000806133b0614825565b6133b8614825565b600060e082018190526133cb8787613e74565b845296509050806133e55750600093508492509050612891565b6133ef8787613a62565b6020850152965090508061340c5750600093508492509050612891565b6134168787613a62565b604085015296509050806134335750600093508492509050612891565b61343d878761275a565b6060850152965090508061345a5750600093508492509050612891565b613464878761275a565b608085015296509050806134815750600093508492509050612891565b61348b8787613848565b60a085015296509050806134a85750600093508492509050612891565b6134b28787613e74565b60c085015296509050806134cf5750600093508492509050612891565b506001969495509392505050565b6134e5614825565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b60008060006135578460ff16613ec8565b50949350505050565b6135686147f1565b611ff460405180608001604052808560ff1681526020018481526020016000151581526020016000801b8152506139fb565b606081015160009060ff166135bb5781516135b4906143be565b9050611763565b606082015160ff16600114156135ee5760208083015180516040820151606083015192909301516135b4939192906143e2565b606082015160ff1660021415613607576135b48261448a565b600360ff16826060015160ff161015801561362b57506060820151600c60ff909116105b15613639576135b4826144f0565b606082015160ff166064141561365157508051611763565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b61369a6147f1565b61222360405180608001604052808660ff168152602001858152602001600115158152602001848152506139fb565b6060015160ff161590565b6128c082602001516136e583612b88565b613bc5565b6136f26147f1565b8115613702576135b46001612b88565b6135b46000612b88565b6137146147f1565b816060015160ff166002141561375b5760405162461bcd60e51b815260040180806020018281038252602181526020018061491e6021913960400191505060405180910390fd5b606082015160ff16613771576135b46000612b88565b816060015160ff166001141561378b576135b46001612b88565b6135b46003612b88565b6060015160ff1660011490565b604080516000808252602082019092526137bd81600161450e565b91505090565b6137cb6147f1565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191613830565b61381d6147f1565b8152602001906001900390816138155790505b50815260646020820152600160409091015292915050565b600080600080855190508481108061386257506020858203105b15613877575060009250839150829050612891565b60016020860161388d888863ffffffff61452d16565b935093509350509250925092565b6000806138a6614883565b600084905060008682815181106138b957fe5b602001015160f81c60f81b60f81c9050818060010192505060008783815181106138df57fe5b016020015160019384019360f89190911c915060009060ff841614156139655760006139096147f1565b6139138b8761275a565b909750909250905081613957575050604080516080810182526000808252602082018190529181018290526060810182905290975088965094506128919350505050565b6139608161359a565b925050505b60006139778a8663ffffffff61452d16565b90506020850194508360ff16600114156139c3576040805160808101825260ff909416845260208401919091526001908301819052606083019190915295509193509091506128919050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b613a036147f1565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190613a4a565b613a376147f1565b815260200190600190039081613a2f5790505b50815260016020820181905260409091015292915050565b600080613a6d6147f1565b613a756147f1565b8551600090819087811080613a8c57506040888203105b15613aa4576000888596509650965050505050612891565b6000613ab68a8a63ffffffff61452d16565b9050602089019850613ac88a8a613848565b909a50945092508215613af357613adf81856101f6565b600198508997509550612891945050505050565b600089869750975097505050505050612891565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015613b5257816020015b613b3f6147f1565b815260200190600190039081613b375790505b50905060005b8960ff168160ff161015613baf57613b70898561275a565b8451859060ff8616908110613b8157fe5b60209081029190910101529450925082613ba75750600095508694509250613bbc915050565b600101613b58565b5060019550919350909150505b93509350939050565b613bcd6147f1565b6040805160028082526060828101909352816020015b613beb6147f1565b815260200190600190039081613be35790505090508281600081518110613c0e57fe5b60200260200101819052508381600181518110613c2757fe5b6020026020010181905250612223613c3e82612c0d565b614549565b600061027682606001516145bf565b600061027682606001516145dd565b606081015160009060ff16613c7857506001611763565b606082015160ff1660011415613c9057506000611763565b606082015160ff1660021415613ce4576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff1610158015613d0857506060820151600c60ff909116105b15613d605760408201515160005b81811015613d5557613d3e84604001518281518110613d3157fe5b6020026020010151613c61565b613d4d57600092505050611763565b600101613d16565b506001915050611763565b606082015160ff166064141561365157506000611763565b6008101590565b60608180613da65750506040805180820190915260018152600360fc1b6020820152611763565b8060005b8115613dbe57600101600a82049150613daa565b6060816040519080825280601f01601f191660200182016040528015613deb576020820181803883390190505b50905060001982015b8415613e3957600a850660300160f81b82828060019003935081518110613e1757fe5b60200101906001600160f81b031916908160001a905350600a85049450613df4565b5095945050505050565b6000600a60f883901c1015613e63578160f81c60300160f81b9050611763565b8160f81c60570160f81b9050611763565b60008060008060008651905085811080613e9057506020868203105b15613ea45750600093508492509050612891565b613eb4878763ffffffff61452d16565b600195506020870194509250612891915050565b6000806001831415613ee057506002905060016143b9565b6002831415613ef557506002905060016143b9565b6003831415613f0a57506002905060016143b9565b6004831415613f1f57506002905060016143b9565b6005831415613f3457506002905060016143b9565b6006831415613f4957506002905060016143b9565b6007831415613f5e57506002905060016143b9565b6008831415613f7357506003905060016143b9565b6009831415613f8857506003905060016143b9565b600a831415613f9d57506002905060016143b9565b6010831415613fb257506002905060016143b9565b6011831415613fc757506002905060016143b9565b6012831415613fdc57506002905060016143b9565b6013831415613ff157506002905060016143b9565b601483141561400657506002905060016143b9565b601583141561401a575060019050806143b9565b601683141561402f57506002905060016143b9565b601783141561404457506002905060016143b9565b601883141561405957506002905060016143b9565b601983141561406d575060019050806143b9565b601a83141561408257506002905060016143b9565b601b83141561409757506002905060016143b9565b60208314156140ab575060019050806143b9565b60218314156140bf575060019050806143b9565b60228314156140d457506002905060016143b9565b60308314156140e957506001905060006143b9565b60318314156140fe57506000905060016143b9565b603283141561411357506000905060016143b9565b603383141561412857506001905060006143b9565b603483141561413d57506001905060006143b9565b603583141561415257506002905060006143b9565b603683141561416757506000905060016143b9565b603783141561417c57506000905060016143b9565b603883141561419157506001905060006143b9565b60398314156141a657506000905060016143b9565b603a8314156141bb57506000905060016143b9565b603b8314156141cf575060009050806143b9565b603c8314156141e457506000905060016143b9565b603d8314156141f957506001905060006143b9565b604083141561420e57506001905060026143b9565b604183141561422357506002905060036143b9565b604283141561423857506003905060046143b9565b604383141561424c575060029050806143b9565b6044831415614260575060039050806143b9565b605083141561427557506002905060016143b9565b605183141561428a57506003905060016143b9565b605283141561429e575060019050806143b9565b60608314156142b2575060009050806143b9565b60618314156142c757506001905060006143b9565b60708314156142dc57506001905060006143b9565b60718314156142f157506000905060016143b9565b6072831415614305575060019050806143b9565b6073831415614319575060009050806143b9565b607483141561432d575060009050806143b9565b607583141561434257506001905060006143b9565b607683141561435757506000905060016143b9565b608083141561436c57506004905060016143b9565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561443c575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120612223565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff166002146144df576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b815160808301516102769190614600565b60006144fa6147f1565b61450383614549565b9050611ff48161448a565b60006145186147f1565b614522848461463a565b90506122238161448a565b6000816020018351101561454057600080fd5b50016020015190565b6145516147f1565b61455a82613c43565b6145a0576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b60606145af8360400151614659565b9050611ff481846080015161463a565b6000600c60ff8316108015610276575050600360ff91909116101590565b60006145e8826145bf565b156145f857506002198101611763565b506001611763565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b6146426147f1565b600061464d84614731565b905061222381846101f6565b60606008825111156146a9576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156146d6578160200160208202803883390190505b50805190915060005b818110156147285760006147058683815181106146f857fe5b602002602001015161359a565b90508084838151811061471457fe5b6020908102919091010152506001016146df565b50909392505050565b6000600882511115614781576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156147c55781810151838201526020016147ad565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a001604052806000815260200161480b614883565b815260606020820181905260006040830181905291015290565b60408051610100810190915260008152602081016148416147f1565b815260200161484e6147f1565b815260200161485b6147f1565b81526020016148686147f1565b81526000602082018190526040820181905260609091015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe73656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f6620686164206e6f6e206d61746368696e672073746172742073746174653a2056616c7565206d757374206861766520612076616c6964207479706520636f646550726f6f6620686164206e6f6e206d61746368696e6720656e642073746174653a2053656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a723158206d7effd0771c3b5fbafd360fcd526acc1105ecf5f24757ecfaee6822efb6342364736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582074383e6f9199b72896a89a067a4eab3190a47e2836df7aaf013dbc168136775b64736f6c63430005110032"

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
