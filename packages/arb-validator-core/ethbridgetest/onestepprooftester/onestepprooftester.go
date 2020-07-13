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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582078448232f00d68ffdf23b3768b45536938fe03b143ca6c3b609de5424513d9f964736f6c63430005110032"

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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d16b51af712f29ca1d2b8113ee0dada23c4cfd51c804f1734a090e0a1aabf79464736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b50613cab806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633c41485d14610030575b600080fd5b61011d600480360361014081101561004757600080fd5b813591602081013591604082013591606081013515159160808201359160a08101359160c08201359160e08101359167ffffffffffffffff61010083013516919081019061014081016101208201356401000000008111156100a857600080fd5b8201836020820111156100ba57600080fd5b803590602001918460018302840111640100000000831117156100dc57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061012f945050505050565b60408051918252519081900360200190f35b600061014b6101468c8c8c8c8c8c8c8c8c8c61015a565b6101bf565b9b9a5050505050505050505050565b610162613aaf565b61014b6040518061012001604052808d81526020016101818d8d610284565b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610337565b600060028260e0015114156101d65750600061027f565b60018260e0015114156101eb5750600161027f565b815160208301516101fb90611487565b6102088460400151611487565b6102158560600151611487565b6102228660800151611487565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090505b919050565b61028c613b0d565b6040805160a08101825284815281516060810183526000808252602082810182905284518281528082018652939490850193908301916102e2565b6102cf613b0d565b8152602001906001900390816102c75790505b5090528152604080516000808252602082810190935291909201919061031e565b61030b613b0d565b8152602001906001900390816103035790505b5081526002602082015260400183905290505b92915050565b61033f613aaf565b6000806000606061034e613aaf565b610356613aaf565b61035f88611590565b60e08e0151959b509399509297509095509350915060019060009067ffffffffffffffff1687146103ce576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b896040015180156103e2575060ff88166072145b806103fe575089604001511580156103fe575060ff8816607214155b61044f576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a0808401805189900390528401518711156104765760001960a084015260009150611340565b60ff8816600114156104bc576104b5838660008151811061049357fe5b6020026020010151876001815181106104a857fe5b60200260200101516117eb565b9150611340565b60ff8816600214156104fb576104b583866000815181106104d957fe5b6020026020010151876001815181106104ee57fe5b602002602001015161183b565b60ff88166003141561053a576104b5838660008151811061051857fe5b60200260200101518760018151811061052d57fe5b602002602001015161187c565b60ff881660041415610579576104b5838660008151811061055757fe5b60200260200101518760018151811061056c57fe5b60200260200101516118bd565b60ff8816600514156105b8576104b5838660008151811061059657fe5b6020026020010151876001815181106105ab57fe5b602002602001015161190e565b60ff8816600614156105f7576104b583866000815181106105d557fe5b6020026020010151876001815181106105ea57fe5b602002602001015161195f565b60ff881660071415610636576104b5838660008151811061061457fe5b60200260200101518760018151811061062957fe5b60200260200101516119b0565b60ff88166008141561068a576104b5838660008151811061065357fe5b60200260200101518760018151811061066857fe5b60200260200101518860028151811061067d57fe5b6020026020010151611a01565b60ff8816600914156106de576104b583866000815181106106a757fe5b6020026020010151876001815181106106bc57fe5b6020026020010151886002815181106106d157fe5b6020026020010151611a6b565b60ff8816600a141561071d576104b583866000815181106106fb57fe5b60200260200101518760018151811061071057fe5b6020026020010151611ac4565b60ff88166010141561075c576104b5838660008151811061073a57fe5b60200260200101518760018151811061074f57fe5b6020026020010151611b05565b60ff88166011141561079b576104b5838660008151811061077957fe5b60200260200101518760018151811061078e57fe5b6020026020010151611b46565b60ff8816601214156107da576104b583866000815181106107b857fe5b6020026020010151876001815181106107cd57fe5b6020026020010151611b87565b60ff881660131415610819576104b583866000815181106107f757fe5b60200260200101518760018151811061080c57fe5b6020026020010151611bc8565b60ff881660141415610858576104b5838660008151811061083657fe5b60200260200101518760018151811061084b57fe5b6020026020010151611c09565b60ff881660151415610882576104b5838660008151811061087557fe5b6020026020010151611c40565b60ff8816601614156108c1576104b5838660008151811061089f57fe5b6020026020010151876001815181106108b457fe5b6020026020010151611c85565b60ff881660171415610900576104b583866000815181106108de57fe5b6020026020010151876001815181106108f357fe5b6020026020010151611cc6565b60ff88166018141561093f576104b5838660008151811061091d57fe5b60200260200101518760018151811061093257fe5b6020026020010151611d07565b60ff881660191415610969576104b5838660008151811061095c57fe5b6020026020010151611d48565b60ff8816601a14156109a8576104b5838660008151811061098657fe5b60200260200101518760018151811061099b57fe5b6020026020010151611d7e565b60ff8816601b14156109e7576104b583866000815181106109c557fe5b6020026020010151876001815181106109da57fe5b6020026020010151611dbf565b60ff881660201415610a11576104b58386600081518110610a0457fe5b6020026020010151611e00565b60ff881660211415610a3b576104b58386600081518110610a2e57fe5b6020026020010151611e1b565b60ff881660221415610a7a576104b58386600081518110610a5857fe5b602002602001015187600181518110610a6d57fe5b6020026020010151611e36565b60ff881660301415610aa4576104b58386600081518110610a9757fe5b6020026020010151611e9c565b60ff881660311415610ab9576104b583611ea4565b60ff881660321415610ace576104b583611ec5565b60ff881660331415610af8576104b58386600081518110610aeb57fe5b6020026020010151611ede565b60ff881660341415610b22576104b58386600081518110610b1557fe5b6020026020010151611eea565b60ff881660351415610b61576104b58386600081518110610b3f57fe5b602002602001015187600181518110610b5457fe5b6020026020010151611f15565b60ff881660361415610b76576104b583611f5d565b60ff881660371415610b90576104b5838560000151611f87565b60ff881660381415610bba576104b58386600081518110610bad57fe5b6020026020010151611f97565b60ff881660391415610c0857610bce613b0d565b610bdd8b610100015188611fa9565b9097509050610bf2858263ffffffff61210616565b610c02848263ffffffff61212016565b50611340565b60ff8816603a1415610c1d576104b58361213a565b60ff8816603b1415610c325760019150611340565b60ff8816603c1415610c47576104b583612157565b60ff8816603d1415610c71576104b58386600081518110610c6457fe5b602002602001015161216b565b60ff881660401415610c9b576104b58386600081518110610c8e57fe5b6020026020010151612199565b60ff881660411415610cda576104b58386600081518110610cb857fe5b602002602001015187600181518110610ccd57fe5b60200260200101516121bb565b60ff881660421415610d2e576104b58386600081518110610cf757fe5b602002602001015187600181518110610d0c57fe5b602002602001015188600281518110610d2157fe5b60200260200101516121ed565b60ff881660431415610d6d576104b58386600081518110610d4b57fe5b602002602001015187600181518110610d6057fe5b602002602001015161222f565b60ff881660441415610dc1576104b58386600081518110610d8a57fe5b602002602001015187600181518110610d9f57fe5b602002602001015188600281518110610db457fe5b6020026020010151612241565b60ff881660501415610e00576104b58386600081518110610dde57fe5b602002602001015187600181518110610df357fe5b6020026020010151612263565b60ff881660511415610e54576104b58386600081518110610e1d57fe5b602002602001015187600181518110610e3257fe5b602002602001015188600281518110610e4757fe5b60200260200101516122d9565b60ff881660521415610e7e576104b58386600081518110610e7157fe5b6020026020010151612366565b60ff881660531415610edd57610e92613b0d565b610ea18b610100015188611fa9565b9097509050610eb6858263ffffffff61210616565b610ed58487600081518110610ec757fe5b602002602001015183612399565b925050611340565b60ff881660541415610f4957610ef1613b0d565b610f008b610100015188611fa9565b9097509050610f15858263ffffffff61210616565b610ed58487600081518110610f2657fe5b602002602001015188600181518110610f3b57fe5b6020026020010151846123f1565b60ff881660601415610f5e576104b583612472565b60ff88166061141561105b57610f888386600081518110610f7b57fe5b6020026020010151612478565b90925090508115611052578960c001518a60a0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146110075760405162461bcd60e51b8152600401808060200182810382526025815260200180613c2b6025913960400191505060405180910390fd5b89608001518a606001511461104d5760405162461bcd60e51b8152600401808060200182810382526027815260200180613c506027913960400191505060405180910390fd5b611056565b5060005b611340565b60ff88166070141561119257611085838660008151811061107857fe5b6020026020010151612491565b9092509050811561105257806110e05789608001518a60600151146110db5760405162461bcd60e51b8152600401808060200182810382526038815260200180613bf36038913960400191505060405180910390fd5b61104d565b60808a01516060808c0151604080516020808201939093528082018690528151808203830181529301905281519101201461114c5760405162461bcd60e51b8152600401808060200182810382526029815260200180613b836029913960400191505060405180910390fd5b8960c001518a60a001511461104d5760405162461bcd60e51b8152600401808060200182810382526026815260200180613bac6026913960400191505060405180910390fd5b60ff8816607214156111ac576104b5838b602001516124d0565b60ff8816607314156111c15760009150611340565b60ff8816607414156111d65761105683612536565b60ff881660751415611200576104b583866000815181106111f357fe5b6020026020010151612540565b60ff881660761415611215576104b583612565565b60ff88166077141561122a576104b58361257e565b60ff881660781415611269576104b5838660008151811061124757fe5b60200260200101518760018151811061125c57fe5b60200260200101516125c7565b60ff8816607914156112bd576104b5838660008151811061128657fe5b60200260200101518760018151811061129b57fe5b6020026020010151886002815181106112b057fe5b602002602001015161260c565b60ff8816607b14156112d2576104b58361265f565b60ff88166080141561133b576104b583866000815181106112ef57fe5b60200260200101518760018151811061130457fe5b60200260200101518860028151811061131957fe5b60200260200101518960038151811061132e57fe5b60200260200101516126a2565b600091505b806113d15789608001518a606001511461138b5760405162461bcd60e51b8152600401808060200182810382526027815260200180613c506027913960400191505060405180910390fd5b8960c001518a60a00151146113d15760405162461bcd60e51b8152600401808060200182810382526026815260200180613bac6026913960400191505060405180910390fd5b816114325760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c0840151141561142a57611425836127c1565b611432565b60c083015183525b61143b846101bf565b8a51146114795760405162461bcd60e51b8152600401808060200182810382526022815260200180613b616022913960400191505060405180910390fd5b509098975050505050505050565b606081015160009060ff166114a85781516114a1906127cb565b905061027f565b606082015160ff16600114156114c5576114a182602001516127ef565b606082015160ff16600214156114e657815160808301516114a191906128db565b600360ff16826060015160ff161015801561150a57506060820151600c60ff909116105b1561153757611517613b0d565b6115248360400151612915565b905061152f81611487565b91505061027f565b606082015160ff166064141561154f5750805161027f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b600080606061159d613aaf565b6115a5613aaf565b60006115b083612a77565b6115bf87610100015182612a81565b935090506115cc83612b21565b9150600087610100015182815181106115e157fe5b602001015160f81c60f81b60f81c9050876101000151826001018151811061160557fe5b016020015160f81c9650600061161a88612b8a565b60408051838152602080850282010190915290985090915081801561165957816020015b611646613b0d565b81526020019060019003908161163e5790505b5095506002830192508160ff166000148061167757508160ff166001145b6116c8576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff82166116ed576116e66116e18987600001516130e7565b611487565b8552611753565b6116f5613b0d565b6117048a610100015185611fa9565b9094509050811561172d57808760008151811061171d57fe5b602002602001018190525061173d565b61173d858263ffffffff61212016565b61174f6116e18a886000015184613142565b8652505b60ff82165b818110156117925761176f8a610100015185611fa9565b885189908490811061177d57fe5b60209081029190910101529350600101611758565b8651156117df575060005b8260ff168751038110156117df576117d7878260018a510303815181106117c057fe5b60200260200101518761212090919063ffffffff16565b60010161179d565b50505091939550919395565b60006117f6836131bd565b15806118085750611806826131bd565b155b1561181557506000611834565b8251825180820161182c878263ffffffff6131c816565b600193505050505b9392505050565b6000611846836131bd565b15806118585750611856826131bd565b155b1561186557506000611834565b8251825180820261182c878263ffffffff6131c816565b6000611887836131bd565b15806118995750611897826131bd565b155b156118a657506000611834565b8251825180820361182c878263ffffffff6131c816565b60006118c8836131bd565b15806118da57506118d8826131bd565b155b156118e757506000611834565b82518251806118fb57600092505050611834565b80820461182c878263ffffffff6131c816565b6000611919836131bd565b158061192b5750611929826131bd565b155b1561193857506000611834565b825182518061194c57600092505050611834565b80820561182c878263ffffffff6131c816565b600061196a836131bd565b158061197c575061197a826131bd565b155b1561198957506000611834565b825182518061199d57600092505050611834565b80820661182c878263ffffffff6131c816565b60006119bb836131bd565b15806119cd57506119cb826131bd565b155b156119da57506000611834565b82518251806119ee57600092505050611834565b80820761182c878263ffffffff6131c816565b6000611a0c846131bd565b1580611a1e5750611a1c836131bd565b155b15611a2b57506000611a63565b83518351835180611a425760009350505050611a63565b6000818385089050611a5a898263ffffffff6131c816565b60019450505050505b949350505050565b6000611a76846131bd565b1580611a885750611a86836131bd565b155b15611a9557506000611a63565b83518351835180611aac5760009350505050611a63565b6000818385099050611a5a898263ffffffff6131c816565b6000611acf836131bd565b1580611ae15750611adf826131bd565b155b15611aee57506000611834565b8251825180820a61182c878263ffffffff6131c816565b6000611b10836131bd565b1580611b225750611b20826131bd565b155b15611b2f57506000611834565b8251825180821061182c878263ffffffff6131c816565b6000611b51836131bd565b1580611b635750611b61826131bd565b155b15611b7057506000611834565b8251825180821161182c878263ffffffff6131c816565b6000611b92836131bd565b1580611ba45750611ba2826131bd565b155b15611bb157506000611834565b8251825180821261182c878263ffffffff6131c816565b6000611bd3836131bd565b1580611be55750611be3826131bd565b155b15611bf257506000611834565b8251825180821361182c878263ffffffff6131c816565b6000611c36611c29611c1a84611487565b611c2386611487565b146131de565b859063ffffffff61212016565b5060019392505050565b6000611c4b826131bd565b611c6557611c6083600063ffffffff6131c816565b611c7c565b81518015611c79858263ffffffff6131c816565b50505b50600192915050565b6000611c90836131bd565b1580611ca25750611ca0826131bd565b155b15611caf57506000611834565b8251825180821661182c878263ffffffff6131c816565b6000611cd1836131bd565b1580611ce35750611ce1826131bd565b155b15611cf057506000611834565b8251825180821761182c878263ffffffff6131c816565b6000611d12836131bd565b1580611d245750611d22826131bd565b155b15611d3157506000611834565b8251825180821861182c878263ffffffff6131c816565b6000611d53826131bd565b611d5f57506000610331565b81518019611d73858263ffffffff6131c816565b506001949350505050565b6000611d89836131bd565b1580611d9b5750611d99826131bd565b155b15611da857506000611834565b8251825181811a61182c878263ffffffff6131c816565b6000611dca836131bd565b1580611ddc5750611dda826131bd565b155b15611de957506000611834565b8251825181810b61182c878263ffffffff6131c816565b6000611c7c611e0e83611487565b849063ffffffff6131c816565b6000611c7c611e2983613200565b849063ffffffff61212016565b6000611e41836131bd565b1580611e535750611e51826131bd565b155b15611e6057506000611834565b825182516040805160208082018590528183018490528251808303840181526060909201909252805191012061182c878263ffffffff6131c816565b600192915050565b6000611ebd82608001518361212090919063ffffffff16565b506001919050565b6000611ebd82606001518361212090919063ffffffff16565b60609190910152600190565b6000611ef582613289565b611f0157506000610331565b611f0a82611487565b835250600192915050565b6000611f2083613289565b611f2c57506000611834565b611f35826131bd565b611f4157506000611834565b815115611c3657611f5183611487565b84525060019392505050565b6000611ebd611f7a611f6d613296565b611c238560200151611487565b839063ffffffff61212016565b6000611c7c611e298360016132a8565b6000611c7c838363ffffffff61210616565b6000611fb3613b0d565b83518310611ff9576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000839050600085828151811061200c57fe5b016020015160019092019160f81c90506000816120485761202d8784613359565b90935090508261203c826133c2565b945094505050506120ff565b60ff82166001141561205e5761203c8784613474565b60ff8216600214156120745761203c8784613538565b600360ff83161080159061208b5750600c60ff8316105b156120bf57600219820160606120a2828a876135dc565b9095509050846120b182613675565b9650965050505050506120ff565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b61211482604001518261378c565b82604001819052505050565b61212e82602001518261378c565b82602001819052505050565b6000611ebd611f7a61214a613296565b611c238560400151611487565b6000611ebd611f7a8360c0015160016132a8565b600061217682613289565b61218257506000610331565b61218b82611487565b60c084015250600192915050565b60006121ab838363ffffffff61212016565b611c7c838363ffffffff61212016565b60006121cd848363ffffffff61212016565b6121dd848463ffffffff61212016565b611c36848363ffffffff61212016565b60006121ff858363ffffffff61212016565b61220f858463ffffffff61212016565b61221f858563ffffffff61212016565b611d73858363ffffffff61212016565b60006121dd848463ffffffff61212016565b6000612253858563ffffffff61212016565b61221f858463ffffffff61212016565b600061226e836131bd565b1580612280575061227e82613802565b155b1561228d57506000611834565b61229682613811565b60ff168360000151106122ab57506000611834565b611c3682604001518460000151815181106122c257fe5b60200260200101518561212090919063ffffffff16565b60006122e483613802565b15806122f657506122f4846131bd565b155b1561230357506000611a63565b61230c83613811565b60ff1684600001511061232157506000611a63565b60408301518451815184918391811061233657fe5b602002602001018190525061235a61234d82613675565b879063ffffffff61212016565b50600195945050505050565b600061237182613802565b61237d57506000610331565b611c7c61238983613811565b849060ff1663ffffffff6131c816565b60006123a4836131bd565b15806123b657506123b482613802565b155b156123c357506000611834565b6123cc82613811565b60ff168360000151106123e157506000611834565b6122ab848363ffffffff61210616565b60006123fc82613802565b158061240e575061240c846131bd565b155b1561241b57506000611a63565b61242482613811565b60ff1684600001511061243957506000611a63565b60408201518451815185918391811061244e57fe5b602002602001018190525061235a61246582613675565b879063ffffffff61210616565b50600190565b600080600161248684611487565b915091509250929050565b600080612710836080015111156124ad575060009050806120ff565b6124b68361383c565b6124c5575060009050806120ff565b600161248684611487565b60006124da613296565b6124e383611487565b14156121ab576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b600260e090910152565b600061254b826131bd565b61255757506000610331565b505160a09190910152600190565b6000611ebd8260a00151836131c890919063ffffffff16565b60408051600160f81b60208083019190915260006021830181905260228084018290528451808503909101815260429093019093528151910120611ebd90611f7a9060016132a8565b60006125d2836131bd565b6125de57506000611834565b6125e782613289565b6125f357506000611834565b611c36611c29846000015161260785611487565b6130e7565b6000612617846131bd565b61262357506000611a63565b61262c82613289565b61263857506000611a63565b611d73612652856000015161264c85611487565b86613142565b869063ffffffff61212016565b6040805160008082526020820190925260609082612693565b612680613b0d565b8152602001906001900390816126785790505b509050611c7c611e2982613675565b60006126ad856131bd565b15806126bf57506126bd846131bd565b155b806126d057506126ce836131bd565b155b806126e157506126df826131bd565b155b156126ee575060006127b8565b8451845184511580159061270457508451600114155b156127255761271a88600063ffffffff6131c816565b6001925050506127b8565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612787573d6000803e3d6000fd5b5050604051601f19015191506127ae90508b6001600160a01b03831663ffffffff6131c816565b6001955050505050505b95945050505050565b600160e090910152565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061280057fe5b60408201515161285a5750805160208083015160408051600160f81b8185015260f89490941b6001600160f81b0319166021850152602280850192909252805180850390920182526042909301909252815191012061027f565b60018260000151612882846040015160008151811061287557fe5b6020026020010151611487565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b61291d613b0d565b60088251111561296b576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015612998578160200160208202803883390190505b508051909150600160005b828110156129fb576129ba86828151811061287557fe5b8482815181106129c657fe5b6020026020010181815250508581815181106129de57fe5b6020026020010151608001518201915080806001019150506129a3565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015612a40578181015183820152602001612a28565b5050505090500192505050604051602081830303815290604052805190602001209050612a6d8183610284565b9695505050505050565b600060e090910152565b6000612a8b613aaf565b612a93613aaf565b600060e0820181905280612aa78787613359565b9096509150612ab68787613538565b60208501529550612ac78787613538565b60408501529550612ad88787611fa9565b60608501529550612ae98787611fa9565b60808501529550612afa8787613359565b60a08501529550612b0b8787613359565b92845260c0840192909252509590945092505050565b612b29613aaf565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612ba257506002905060036130e2565b6002831415612bb757506002905060036130e2565b6003831415612bcc57506002905060036130e2565b6004831415612be157506002905060046130e2565b6005831415612bf657506002905060076130e2565b6006831415612c0b57506002905060046130e2565b6007831415612c2057506002905060076130e2565b6008831415612c3557506003905060046130e2565b6009831415612c4a57506003905060046130e2565b600a831415612c5f57506002905060196130e2565b6010831415612c73575060029050806130e2565b6011831415612c87575060029050806130e2565b6012831415612c9b575060029050806130e2565b6013831415612caf575060029050806130e2565b6014831415612cc3575060029050806130e2565b6015831415612cd7575060019050806130e2565b6016831415612ceb575060029050806130e2565b6017831415612cff575060029050806130e2565b6018831415612d13575060029050806130e2565b6019831415612d27575060019050806130e2565b601a831415612d3c57506002905060046130e2565b601b831415612d5157506002905060076130e2565b6020831415612d6657506001905060076130e2565b6021831415612d7b57506001905060036130e2565b6022831415612d9057506002905060086130e2565b6030831415612da4575060019050806130e2565b6031831415612db957506000905060016130e2565b6032831415612dce57506000905060016130e2565b6033831415612de357506001905060026130e2565b6034831415612df857506001905060046130e2565b6035831415612e0d57506002905060046130e2565b6036831415612e2257506000905060026130e2565b6037831415612e3757506000905060016130e2565b6038831415612e4b575060019050806130e2565b6039831415612e6057506000905060016130e2565b603a831415612e7557506000905060026130e2565b603b831415612e8a57506000905060016130e2565b603c831415612e9f57506000905060016130e2565b603d831415612eb3575060019050806130e2565b6040831415612ec7575060019050806130e2565b6041831415612edc57506002905060016130e2565b6042831415612ef157506003905060016130e2565b6043831415612f0657506002905060016130e2565b6044831415612f1b57506003905060016130e2565b6050831415612f2f575060029050806130e2565b6051831415612f4457506003905060286130e2565b6052831415612f5957506001905060026130e2565b6053831415612f6e57506001905060036130e2565b6054831415612f8357506002905060296130e2565b6060831415612f9857506000905060646130e2565b6061831415612fad57506001905060646130e2565b6070831415612fc257506001905060646130e2565b6072831415612fd757506000905060286130e2565b6073831415612fec57506000905060056130e2565b6074831415613001575060009050600a6130e2565b607583141561301657506001905060006130e2565b607683141561302b57506000905060016130e2565b607783141561304057506000905060196130e2565b607883141561305557506002905060196130e2565b607983141561306a57506003905060196130e2565b607b83141561307f575060009050600a6130e2565b6080831415613095575060049050614e206130e2565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b6130ef613b0d565b6040805160608101825260ff8516815260208082018590528251600080825291810184526118349383019161313a565b613127613b0d565b81526020019060019003908161311f5790505b509052613953565b61314a613b0d565b604080516001808252818301909252606091816020015b613169613b0d565b815260200190600190039081613161579050509050828160008151811061318c57fe5b60200260200101819052506127b860405180606001604052808760ff16815260200186815260200183815250613953565b6060015160ff161590565b61212e82602001516131d9836133c2565b61378c565b6131e6613b0d565b81156131f6576114a160016133c2565b6114a160006133c2565b613208613b0d565b816060015160ff166002141561324f5760405162461bcd60e51b8152600401808060200182810382526021815260200180613bd26021913960400191505060405180910390fd5b606082015160ff16613265576114a160006133c2565b816060015160ff166001141561327f576114a160016133c2565b6114a160036133c2565b6060015160ff1660011490565b60006132a36116e16139ba565b905090565b6132b0613b0d565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191613306565b6132f3613b0d565b8152602001906001900390816132eb5790505b50905281526040805160008082526020828101909352919092019190613342565b61332f613b0d565b8152602001906001900390816133275790505b508152606460208201526040019290925250919050565b60008082845110158015613371575060208385510310155b6133ae576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301612486858563ffffffff613a6e16565b6133ca613b0d565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191613420565b61340d613b0d565b8152602001906001900390816134055790505b5090528152604080516000808252602082810190935291909201919061345c565b613449613b0d565b8152602001906001900390816134415790505b50815260006020820152600160409091015292915050565b600061347e613b0d565b6000839050600085828151811061349157fe5b602001015160f81c60f81b60f81c9050818060010192505060008683815181106134b757fe5b016020015160019093019260f81c90506134cf613b0d565b8260ff16600114156134eb576134e58885611fa9565b90945090505b60006134fd898663ffffffff613a6e16565b90506020850194508360ff166001141561351d57846120b1848385613142565b8461352884836130e7565b9650965050505050509250929050565b6000613542613b0d565b600083855110158015613559575060408486510310155b613595576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b60006135a7868663ffffffff613a6e16565b90506020850194506135b98686613359565b90955091506135c6613b0d565b6135d08284610284565b95979596505050505050565b60006060600083905060608660ff1660405190808252806020026020018201604052801561362457816020015b613611613b0d565b8152602001906001900390816136095790505b50905060005b8760ff168160ff161015613668576136428784611fa9565b8351849060ff851690811061365357fe5b6020908102919091010152925060010161362a565b5090969095509350505050565b61367d613b0d565b6136878251613a8a565b6136d8576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b835181101561370f578381815181106136f257fe5b6020026020010151608001518201915080806001019150506136dd565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190613769565b613756613b0d565b81526020019060019003908161374e5790505b509052815260208101859052935160030160ff1660408501526060909301525090565b613794613b0d565b6040805160028082526060828101909352816020015b6137b2613b0d565b8152602001906001900390816137aa57905050905082816000815181106137d557fe5b602002602001018190525083816001815181106137ee57fe5b6020026020010181905250611a6381612915565b60006103318260600151613a91565b60006138208260600151613a91565b15613834575060608101516002190161027f565b50600161027f565b606081015160009060ff166138535750600161027f565b606082015160ff166001141561386b5750600061027f565b606082015160ff16600214156138bf576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff16101580156138e357506060820151600c60ff909116105b1561393b5760408201515160005b81811015613930576139198460400151828151811061390c57fe5b602002602001015161383c565b6139285760009250505061027f565b6001016138f1565b50600191505061027f565b606082015160ff166064141561154f5750600061027f565b61395b613b0d565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906139a2565b61398f613b0d565b8152602001906001900390816139875790505b50815260016020820181905260409091015292915050565b6139c2613b0d565b6040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190613a1b565b613a08613b0d565b815260200190600190039081613a005790505b50905281526040805160008082526020828101909352919092019190613a57565b613a44613b0d565b815260200190600190039081613a3c5790505b508152600360208201526001604090910152905090565b60008160200183511015613a8157600080fd5b50016020015190565b6008101590565b6000600c60ff8316108015610331575050600360ff91909116101590565b6040805161010081019091526000815260208101613acb613b0d565b8152602001613ad8613b0d565b8152602001613ae5613b0d565b8152602001613af2613b0d565b81526000602082018190526040820181905260609091015290565b6040518060a0016040528060008152602001613b27613b41565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726fa265627a7a723158204985304c84114f5a0b4d0ab35fec0ff207f57ffd7314c481dba4cbbb8a68e24964736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820783df50db562880f31b02767d5965e7d3e1dfa00c7f3fec9ee824ef6915cf4ab64736f6c63430005110032"

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
