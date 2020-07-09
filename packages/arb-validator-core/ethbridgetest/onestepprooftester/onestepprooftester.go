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
var OneStepProofBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820bfc6478b806bc891087e8a926d72345230d88c28546cdffa55d74a2b873451c864736f6c63430005110032"

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
var OneStepProofTesterBin = "0x608060405234801561001057600080fd5b506141f5806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e987d88714610030575b600080fd5b61015260048036036101e081101561004757600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525091948335946020850135946040810135945060608101351515935060808101359260a08201359260c08301359260e08101359267ffffffffffffffff6101008301351692909161014081019061012001356401000000008111156100dd57600080fd5b8201836020820111156100ef57600080fd5b8035906020019184600183028401116401000000008311171561011157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610164945050505050565b60408051918252519081900360200190f35b600061017a8d8d8d8d8d8d8d8d8d8d8d8d61018b565b9d9c50505050505050505050505050565b600061017a6040518061016001604052808f81526020018e81526020016101b28e8e6101f6565b81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff16815260200184815250610283565b6101fe613fc9565b6040805160a08082018352858252825190810183526000808252602082810182905282850182905260608301829052608083018290528084019290925283518181529182018452919283019161026a565b610257613fc9565b81526020019060019003908161024f5790505b5081526002602082015260400183905290505b92915050565b60008060008060006060610295613ffd565b61029d613ffd565b6102a689611642565b6101208f0151959c50939a509297509095509350915060019060009067ffffffffffffffff168814610316576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b8a60800151801561032a575060ff89166072145b8061034657508a60800151158015610346575060ff8916607214155b610397576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60a080840180518a900390528401518811156103be5760001960a084015260009150611495565b60ff891660011415610404576103fd83866000815181106103db57fe5b6020026020010151876001815181106103f057fe5b602002602001015161199c565b9150611495565b60ff891660021415610443576103fd838660008151811061042157fe5b60200260200101518760018151811061043657fe5b60200260200101516119ec565b60ff891660031415610482576103fd838660008151811061046057fe5b60200260200101518760018151811061047557fe5b6020026020010151611a2d565b60ff8916600414156104c1576103fd838660008151811061049f57fe5b6020026020010151876001815181106104b457fe5b6020026020010151611a6e565b60ff891660051415610500576103fd83866000815181106104de57fe5b6020026020010151876001815181106104f357fe5b6020026020010151611abf565b60ff89166006141561053f576103fd838660008151811061051d57fe5b60200260200101518760018151811061053257fe5b6020026020010151611b10565b60ff89166007141561057e576103fd838660008151811061055c57fe5b60200260200101518760018151811061057157fe5b6020026020010151611b61565b60ff8916600814156105d2576103fd838660008151811061059b57fe5b6020026020010151876001815181106105b057fe5b6020026020010151886002815181106105c557fe5b6020026020010151611bb2565b60ff891660091415610626576103fd83866000815181106105ef57fe5b60200260200101518760018151811061060457fe5b60200260200101518860028151811061061957fe5b6020026020010151611c1c565b60ff8916600a1415610665576103fd838660008151811061064357fe5b60200260200101518760018151811061065857fe5b6020026020010151611c75565b60ff8916601014156106a4576103fd838660008151811061068257fe5b60200260200101518760018151811061069757fe5b6020026020010151611cb6565b60ff8916601114156106e3576103fd83866000815181106106c157fe5b6020026020010151876001815181106106d657fe5b6020026020010151611cf7565b60ff891660121415610722576103fd838660008151811061070057fe5b60200260200101518760018151811061071557fe5b6020026020010151611d38565b60ff891660131415610761576103fd838660008151811061073f57fe5b60200260200101518760018151811061075457fe5b6020026020010151611d79565b60ff8916601414156107a0576103fd838660008151811061077e57fe5b60200260200101518760018151811061079357fe5b6020026020010151611dba565b60ff8916601514156107ca576103fd83866000815181106107bd57fe5b6020026020010151611de4565b60ff891660161415610809576103fd83866000815181106107e757fe5b6020026020010151876001815181106107fc57fe5b6020026020010151611e29565b60ff891660171415610848576103fd838660008151811061082657fe5b60200260200101518760018151811061083b57fe5b6020026020010151611e6a565b60ff891660181415610887576103fd838660008151811061086557fe5b60200260200101518760018151811061087a57fe5b6020026020010151611eab565b60ff8916601914156108b1576103fd83866000815181106108a457fe5b6020026020010151611eec565b60ff8916601a14156108f0576103fd83866000815181106108ce57fe5b6020026020010151876001815181106108e357fe5b6020026020010151611f22565b60ff8916601b141561092f576103fd838660008151811061090d57fe5b60200260200101518760018151811061092257fe5b6020026020010151611f63565b60ff891660201415610959576103fd838660008151811061094c57fe5b6020026020010151611fa4565b60ff891660211415610983576103fd838660008151811061097657fe5b6020026020010151611fbf565b60ff8916602214156109c2576103fd83866000815181106109a057fe5b6020026020010151876001815181106109b557fe5b6020026020010151611fda565b60ff8916603014156109ec576103fd83866000815181106109df57fe5b6020026020010151612040565b60ff891660311415610a01576103fd83612048565b60ff891660321415610a16576103fd83612069565b60ff891660331415610a40576103fd8386600081518110610a3357fe5b6020026020010151612082565b60ff891660341415610a6a576103fd8386600081518110610a5d57fe5b602002602001015161208e565b60ff891660351415610aa9576103fd8386600081518110610a8757fe5b602002602001015187600181518110610a9c57fe5b60200260200101516120b9565b60ff891660361415610abe576103fd83612101565b60ff891660371415610ad8576103fd83856000015161212b565b60ff891660381415610b02576103fd8386600081518110610af557fe5b602002602001015161213b565b60ff891660391415610b8e57610b16613fc9565b610b258c61014001518861214d565b9199509750905087610b685760405162461bcd60e51b81526004018080602001828103825260218152602001806141a06021913960400191505060405180910390fd5b610b78858263ffffffff61228b16565b610b88848263ffffffff6122a516565b50611495565b60ff8916603a1415610ba3576103fd836122bf565b60ff8916603b1415610bb85760019150611495565b60ff8916603c1415610bcd576103fd836122dc565b60ff8916603d1415610bf7576103fd8386600081518110610bea57fe5b60200260200101516122f0565b60ff891660401415610c21576103fd8386600081518110610c1457fe5b602002602001015161231e565b60ff891660411415610c60576103fd8386600081518110610c3e57fe5b602002602001015187600181518110610c5357fe5b6020026020010151612340565b60ff891660421415610cb4576103fd8386600081518110610c7d57fe5b602002602001015187600181518110610c9257fe5b602002602001015188600281518110610ca757fe5b6020026020010151612372565b60ff891660431415610cf3576103fd8386600081518110610cd157fe5b602002602001015187600181518110610ce657fe5b60200260200101516123b4565b60ff891660441415610d47576103fd8386600081518110610d1057fe5b602002602001015187600181518110610d2557fe5b602002602001015188600281518110610d3a57fe5b60200260200101516123c6565b60ff891660501415610d86576103fd8386600081518110610d6457fe5b602002602001015187600181518110610d7957fe5b60200260200101516123e8565b60ff891660511415610dda576103fd8386600081518110610da357fe5b602002602001015187600181518110610db857fe5b602002602001015188600281518110610dcd57fe5b602002602001015161245e565b60ff891660521415610e04576103fd8386600081518110610df757fe5b60200260200101516124eb565b60ff891660531415610ea157610e18613fc9565b610e278c61014001518861214d565b9199509750905087610e6a5760405162461bcd60e51b81526004018080602001828103825260218152602001806141a06021913960400191505060405180910390fd5b610e7a858263ffffffff61228b16565b610e998487600081518110610e8b57fe5b60200260200101518361251e565b925050611495565b60ff891660541415610f4b57610eb5613fc9565b610ec48c61014001518861214d565b9199509750905087610f075760405162461bcd60e51b81526004018080602001828103825260218152602001806141a06021913960400191505060405180910390fd5b610f17858263ffffffff61228b16565b610e998487600081518110610f2857fe5b602002602001015188600181518110610f3d57fe5b602002602001015184612576565b60ff891660601415610f60576103fd836125f7565b60ff89166061141561105e57610f8a8386600081518110610f7d57fe5b60200260200101516125fd565b90925090508115611055578a61010001518b60e00151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001201461100a5760405162461bcd60e51b81526004018080602001828103825260258152602001806141546025913960400191505060405180910390fd5b8a60c001518b60a00151146110505760405162461bcd60e51b81526004018080602001828103825260278152602001806141796027913960400191505060405180910390fd5b611059565b5060005b611495565b60ff89166070141561119e57611088838660008151811061107b57fe5b6020026020010151612617565b9092509050811561105557806110e3578a60c001518b60a00151146110de5760405162461bcd60e51b815260040180806020018281038252603881526020018061411c6038913960400191505060405180910390fd5b611050565b8a60c001518b60a0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146111575760405162461bcd60e51b81526004018080602001828103825260298152602001806140ac6029913960400191505060405180910390fd5b8a61010001518b60e00151146110505760405162461bcd60e51b81526004018080602001828103825260268152602001806140d56026913960400191505060405180910390fd5b60ff8916607114156112b35760408051600480825260a08201909252606091816020015b6111ca613fc9565b8152602001906001900390816111c257505060208d01519091506111fe9060005b60200201516001600160801b0316612656565b8160008151811061120b57fe5b602002602001018190525061122a8c602001516001600481106111eb57fe5b8160018151811061123757fe5b60200260200101819052506112568c602001516002600481106111eb57fe5b8160028151811061126357fe5b60200260200101819052506112828c602001516003600481106111eb57fe5b8160038151811061128f57fe5b6020026020010181905250610b886112a6826126e2565b859063ffffffff6122a516565b60ff891660721415611301576103fd83866000815181106112d057fe5b60200260200101518d604001518e602001516000600481106112ee57fe5b60200201516001600160801b03166127d9565b60ff8916607314156113165760009150611495565b60ff89166074141561132b5761105983612870565b60ff891660751415611355576103fd838660008151811061134857fe5b602002602001015161287a565b60ff89166076141561136a576103fd8361289f565b60ff89166077141561137f576103fd836128b8565b60ff8916607814156113be576103fd838660008151811061139c57fe5b6020026020010151876001815181106113b157fe5b6020026020010151612901565b60ff891660791415611412576103fd83866000815181106113db57fe5b6020026020010151876001815181106113f057fe5b60200260200101518860028151811061140557fe5b6020026020010151612946565b60ff8916607b1415611427576103fd83612999565b60ff891660801415611490576103fd838660008151811061144457fe5b60200260200101518760018151811061145957fe5b60200260200101518860028151811061146e57fe5b60200260200101518960038151811061148357fe5b60200260200101516129dc565b600091505b80611527578a60c001518b60a00151146114e05760405162461bcd60e51b81526004018080602001828103825260278152602001806141796027913960400191505060405180910390fd5b8a61010001518b60e00151146115275760405162461bcd60e51b81526004018080602001828103825260268152602001806140d56026913960400191505060405180910390fd5b816115885760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060c084015114156115805761157b83612afb565b611588565b60c083015183525b61159184612b05565b8b51146115cf5760405162461bcd60e51b815260040180806020018281038252602281526020018061408a6022913960400191505060405180910390fd5b6115d883612b05565b8b606001511461162f576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b600099505050505050505050505b919050565b600080606061164f613ffd565b611657613ffd565b60008061166384612bc9565b61167288610140015183612bd3565b955092509050806116ca576040805162461bcd60e51b815260206004820152601e60248201527f6c6f61644d616368696e6528293a20696e76616c6964206d616368696e650000604482015290519081900360640190fd5b6116d384612d0b565b9250600088610140015183815181106116e857fe5b602001015160f81c60f81b60f81c9050886101400151836001018151811061170c57fe5b016020015160f81c9750600061172189612d74565b60408051838152602080850282010190915290995090915081801561176057816020015b61174d613fc9565b8152602001906001900390816117455790505b5096506002840193508160ff166000148061177e57508160ff166001145b6117cf576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff82166117f4576117ed6117e88a88600001516132e6565b613327565b86526118af565b6117fc613fc9565b61180b8b61014001518661214d565b909650909450905083611865576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b811561188957808860008151811061187957fe5b6020026020010181905250611899565b611899868263ffffffff6122a516565b6118ab6117e88b89600001518461341f565b8752505b60ff82165b81811015611942576118cb8b61014001518661214d565b8a518b90859081106118d957fe5b6020908102919091010152955093508361193a576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b6001016118b4565b87511561198f575060005b8260ff1688510381101561198f57611987888260018b5103038151811061197057fe5b6020026020010151886122a590919063ffffffff16565b60010161194d565b5050505091939550919395565b60006119a783613468565b15806119b957506119b782613468565b155b156119c6575060006119e5565b825182518082016119dd878263ffffffff61347316565b600193505050505b9392505050565b60006119f783613468565b1580611a095750611a0782613468565b155b15611a16575060006119e5565b825182518082026119dd878263ffffffff61347316565b6000611a3883613468565b1580611a4a5750611a4882613468565b155b15611a57575060006119e5565b825182518082036119dd878263ffffffff61347316565b6000611a7983613468565b1580611a8b5750611a8982613468565b155b15611a98575060006119e5565b8251825180611aac576000925050506119e5565b8082046119dd878263ffffffff61347316565b6000611aca83613468565b1580611adc5750611ada82613468565b155b15611ae9575060006119e5565b8251825180611afd576000925050506119e5565b8082056119dd878263ffffffff61347316565b6000611b1b83613468565b1580611b2d5750611b2b82613468565b155b15611b3a575060006119e5565b8251825180611b4e576000925050506119e5565b8082066119dd878263ffffffff61347316565b6000611b6c83613468565b1580611b7e5750611b7c82613468565b155b15611b8b575060006119e5565b8251825180611b9f576000925050506119e5565b8082076119dd878263ffffffff61347316565b6000611bbd84613468565b1580611bcf5750611bcd83613468565b155b15611bdc57506000611c14565b83518351835180611bf35760009350505050611c14565b6000818385089050611c0b898263ffffffff61347316565b60019450505050505b949350505050565b6000611c2784613468565b1580611c395750611c3783613468565b155b15611c4657506000611c14565b83518351835180611c5d5760009350505050611c14565b6000818385099050611c0b898263ffffffff61347316565b6000611c8083613468565b1580611c925750611c9082613468565b155b15611c9f575060006119e5565b8251825180820a6119dd878263ffffffff61347316565b6000611cc183613468565b1580611cd35750611cd182613468565b155b15611ce0575060006119e5565b825182518082106119dd878263ffffffff61347316565b6000611d0283613468565b1580611d145750611d1282613468565b155b15611d21575060006119e5565b825182518082116119dd878263ffffffff61347316565b6000611d4383613468565b1580611d555750611d5382613468565b155b15611d62575060006119e5565b825182518082126119dd878263ffffffff61347316565b6000611d8483613468565b1580611d965750611d9482613468565b155b15611da3575060006119e5565b825182518082136119dd878263ffffffff61347316565b6000611dda6112a6611dcb84613327565b611dd486613327565b14613489565b5060019392505050565b6000611def82613468565b611e0957611e0483600063ffffffff61347316565b611e20565b81518015611e1d858263ffffffff61347316565b50505b50600192915050565b6000611e3483613468565b1580611e465750611e4482613468565b155b15611e53575060006119e5565b825182518082166119dd878263ffffffff61347316565b6000611e7583613468565b1580611e875750611e8582613468565b155b15611e94575060006119e5565b825182518082176119dd878263ffffffff61347316565b6000611eb683613468565b1580611ec85750611ec682613468565b155b15611ed5575060006119e5565b825182518082186119dd878263ffffffff61347316565b6000611ef782613468565b611f035750600061027d565b81518019611f17858263ffffffff61347316565b506001949350505050565b6000611f2d83613468565b1580611f3f5750611f3d82613468565b155b15611f4c575060006119e5565b8251825181811a6119dd878263ffffffff61347316565b6000611f6e83613468565b1580611f805750611f7e82613468565b155b15611f8d575060006119e5565b8251825181810b6119dd878263ffffffff61347316565b6000611e20611fb283613327565b849063ffffffff61347316565b6000611e20611fcd836134ab565b849063ffffffff6122a516565b6000611fe583613468565b1580611ff75750611ff582613468565b155b15612004575060006119e5565b82518251604080516020808201859052818301849052825180830384018152606090920190925280519101206119dd878263ffffffff61347316565b600192915050565b60006120618260800151836122a590919063ffffffff16565b506001919050565b60006120618260600151836122a590919063ffffffff16565b60609190910152600190565b600061209982613534565b6120a55750600061027d565b6120ae82613327565b835250600192915050565b60006120c483613534565b6120d0575060006119e5565b6120d982613468565b6120e5575060006119e5565b815115611dda576120f583613327565b84525060019392505050565b600061206161211e612111613541565b611dd48560200151613327565b839063ffffffff6122a516565b6000611e20611fcd836001613562565b6000611e20838363ffffffff61228b16565b600080612158613fc9565b845184106121785760008461216d6000612656565b925092509250612284565b600080859050600087828151811061218c57fe5b016020015160019092019160f81c905060006121a661405b565b60ff83166121da576121b88a856135ed565b9196509450915084846121ca84612656565b9750975097505050505050612284565b60ff831660011415612202576121f08a85613640565b9196509450905084846121ca836137be565b60ff831660021415612218576121ca8a85613825565b600360ff84161080159061222f5750600c60ff8416105b1561226a5760021983016060612246828d886138ca565b919850965090508686612258836126e2565b99509950995050505050505050612284565b6000806122776000612656565b9199509750955050505050505b9250925092565b612299826040015182613988565b82604001819052505050565b6122b3826020015182613988565b82602001819052505050565b600061206161211e6122cf613541565b611dd48560400151613327565b600061206161211e8360c001516001613562565b60006122fb82613534565b6123075750600061027d565b61231082613327565b60c084015250600192915050565b6000612330838363ffffffff6122a516565b611e20838363ffffffff6122a516565b6000612352848363ffffffff6122a516565b612362848463ffffffff6122a516565b611dda848363ffffffff6122a516565b6000612384858363ffffffff6122a516565b612394858463ffffffff6122a516565b6123a4858563ffffffff6122a516565b611f17858363ffffffff6122a516565b6000612362848463ffffffff6122a516565b60006123d8858563ffffffff6122a516565b6123a4858463ffffffff6122a516565b60006123f383613468565b1580612405575061240382613a06565b155b15612412575060006119e5565b61241b82613a15565b60ff16836000015110612430575060006119e5565b611dda826040015184600001518151811061244757fe5b6020026020010151856122a590919063ffffffff16565b600061246983613a06565b158061247b575061247984613468565b155b1561248857506000611c14565b61249183613a15565b60ff168460000151106124a657506000611c14565b6040830151845181518491839181106124bb57fe5b60200260200101819052506124df6124d2826126e2565b879063ffffffff6122a516565b50600195945050505050565b60006124f682613a06565b6125025750600061027d565b611e2061250e83613a15565b849060ff1663ffffffff61347316565b600061252983613468565b158061253b575061253982613a06565b155b15612548575060006119e5565b61255182613a15565b60ff16836000015110612566575060006119e5565b612430848363ffffffff61228b16565b600061258182613a06565b1580612593575061259184613468565b155b156125a057506000611c14565b6125a982613a15565b60ff168460000151106125be57506000611c14565b6040820151845181518591839181106125d357fe5b60200260200101819052506124df6125ea826126e2565b879063ffffffff61228b16565b50600190565b600080600161260b84613327565b915091505b9250929050565b6000806127108360800151111561263357506000905080612610565b61263c83613a24565b61264b57506000905080612610565b600161260b84613327565b61265e613fc9565b6040805160a0808201835284825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916126ca565b6126b7613fc9565b8152602001906001900390816126af5790505b50815260006020820152600160409091015292915050565b6126ea613fc9565b6126f48251613b3b565b612745576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b835181101561277c5783818151811061275f57fe5b60200260200101516080015182019150808060010191505061274a565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b60006127e484613468565b6127f057506000611c14565b83518211158061280f5750612803613541565b61280c84613327565b14155b612860576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b611f17858463ffffffff6122a516565b600260e090910152565b600061288582613468565b6128915750600061027d565b505160a09190910152600190565b60006120618260a001518361347390919063ffffffff16565b60408051600160f81b602080830191909152600060218301819052602280840182905284518085039091018152604290930190935281519101206120619061211e906001613562565b600061290c83613468565b612918575060006119e5565b61292182613534565b61292d575060006119e5565b611dda6112a6846000015161294185613327565b6132e6565b600061295184613468565b61295d57506000611c14565b61296682613534565b61297257506000611c14565b611f1761298c856000015161298685613327565b8661341f565b869063ffffffff6122a516565b60408051600080825260208201909252606090826129cd565b6129ba613fc9565b8152602001906001900390816129b25790505b509050611e20611fcd826126e2565b60006129e785613468565b15806129f957506129f784613468565b155b80612a0a5750612a0883613468565b155b80612a1b5750612a1982613468565b155b15612a2857506000612af2565b84518451845115801590612a3e57508451600114155b15612a5f57612a5488600063ffffffff61347316565b600192505050612af2565b84518451604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa158015612ac1573d6000803e3d6000fd5b5050604051601f1901519150612ae890508b6001600160a01b03831663ffffffff61347316565b6001955050505050505b95945050505050565b600160e090910152565b600060028260e001511415612b1c5750600061163d565b60018260e001511415612b315750600161163d565b81516020830151612b4190613327565b612b4e8460400151613327565b612b5b8560600151613327565b612b688660800151613327565b8660a001518760c001516040516020018088815260200187815260200186815260200185815260200184815260200183815260200182815260200197505050505050505060405160208183030381529060405280519060200120905061163d565b600060e090910152565b600080612bde613ffd565b612be6613ffd565b600060e08201819052612bf98787613b42565b84529650905080612c135750600093508492509050612284565b612c1d8787613825565b60208501529650905080612c3a5750600093508492509050612284565b612c448787613825565b60408501529650905080612c615750600093508492509050612284565b612c6b878761214d565b60608501529650905080612c885750600093508492509050612284565b612c92878761214d565b60808501529650905080612caf5750600093508492509050612284565b612cb987876135ed565b60a08501529650905080612cd65750600093508492509050612284565b612ce08787613b42565b60c08501529650905080612cfd5750600093508492509050612284565b506001969495509392505050565b612d13613ffd565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b6000806001831415612d8c57506002905060036132e1565b6002831415612da157506002905060036132e1565b6003831415612db657506002905060036132e1565b6004831415612dcb57506002905060046132e1565b6005831415612de057506002905060076132e1565b6006831415612df557506002905060046132e1565b6007831415612e0a57506002905060076132e1565b6008831415612e1f57506003905060046132e1565b6009831415612e3457506003905060046132e1565b600a831415612e4957506002905060196132e1565b6010831415612e5d575060029050806132e1565b6011831415612e71575060029050806132e1565b6012831415612e85575060029050806132e1565b6013831415612e99575060029050806132e1565b6014831415612ead575060029050806132e1565b6015831415612ec1575060019050806132e1565b6016831415612ed5575060029050806132e1565b6017831415612ee9575060029050806132e1565b6018831415612efd575060029050806132e1565b6019831415612f11575060019050806132e1565b601a831415612f2657506002905060046132e1565b601b831415612f3b57506002905060076132e1565b6020831415612f5057506001905060076132e1565b6021831415612f6557506001905060036132e1565b6022831415612f7a57506002905060086132e1565b6030831415612f8e575060019050806132e1565b6031831415612fa357506000905060016132e1565b6032831415612fb857506000905060016132e1565b6033831415612fcd57506001905060026132e1565b6034831415612fe257506001905060046132e1565b6035831415612ff757506002905060046132e1565b603683141561300c57506000905060026132e1565b603783141561302157506000905060016132e1565b6038831415613035575060019050806132e1565b603983141561304a57506000905060016132e1565b603a83141561305f57506000905060026132e1565b603b83141561307457506000905060016132e1565b603c83141561308957506000905060016132e1565b603d83141561309d575060019050806132e1565b60408314156130b1575060019050806132e1565b60418314156130c657506002905060016132e1565b60428314156130db57506003905060016132e1565b60438314156130f057506002905060016132e1565b604483141561310557506003905060016132e1565b6050831415613119575060029050806132e1565b605183141561312e57506003905060286132e1565b605283141561314357506001905060026132e1565b605383141561315857506001905060036132e1565b605483141561316d57506002905060296132e1565b606083141561318257506000905060646132e1565b606183141561319757506001905060646132e1565b60708314156131ac57506001905060646132e1565b60718314156131c157506000905060286132e1565b60728314156131d657506001905060286132e1565b60738314156131eb57506000905060056132e1565b6074831415613200575060009050600a6132e1565b607583141561321557506001905060006132e1565b607683141561322a57506000905060016132e1565b607783141561323f57506000905060196132e1565b607883141561325457506002905060196132e1565b607983141561326957506003905060196132e1565b607b83141561327e575060009050600a6132e1565b6080831415613294575060049050614e206132e1565b6040805162461bcd60e51b815260206004820152601860248201527f496e76616c6964206f70636f64653a206f70496e666f28290000000000000000604482015290519081900360640190fd5b915091565b6132ee613fc9565b6119e56040518060a001604052808560ff1681526020018481526020016000151581526020016000801b815260200160008152506137be565b606081015160009060ff1661334857815161334190613b96565b905061163d565b606082015160ff166001141561337b57602080830151805160408201516060830151929093015161334193919290613bba565b606082015160ff16600214156133945761334182613c62565b600360ff16826060015160ff16101580156133b857506060820151600c60ff909116105b156133c65761334182613cc8565b606082015160ff16606414156133de5750805161163d565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b613427613fc9565b611c146040518060a001604052808660ff16815260200185815260200160011515815260200161345685613327565b815260200184608001518152506137be565b6060015160ff161590565b6122b3826020015161348483612656565b613988565b613491613fc9565b81156134a1576133416001612656565b6133416000612656565b6134b3613fc9565b816060015160ff16600214156134fa5760405162461bcd60e51b81526004018080602001828103825260218152602001806140fb6021913960400191505060405180910390fd5b606082015160ff16613510576133416000612656565b816060015160ff166001141561352a576133416001612656565b6133416003612656565b6060015160ff1660011490565b6040805160008082526020820190925261355c816001613ce6565b91505090565b61356a613fc9565b6040805160a0808201835285825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916135d6565b6135c3613fc9565b8152602001906001900390816135bb5790505b508152606460208201526040019290925250919050565b600080600080855190508481108061360757506020858203105b1561361c575060009250839150829050612284565b600160208601613632888863ffffffff613d0516565b935093509350509250925092565b60008061364b61405b565b6000849050600086828151811061365e57fe5b602001015160f81c60f81b60f81c90508180600101925050600087838151811061368457fe5b016020015160019384019360f89190911c9150600090819060ff8516141561371b5760006136b0613fc9565b6136ba8c8861214d565b9098509092509050816137065750506040805160a08101825260008082526020820181905291810182905260608101829052608081018290529098508997509550612284945050505050565b61370f81613327565b93508060800151925050505b600061372d8b8763ffffffff613d0516565b90506020860195508460ff166001141561377e576040805160a08101825260ff9095168552602085019190915260019084018190526060840192909252608083015295509193509091506122849050565b6040805160a08101825260ff9095168552602085019190915260009084018190526060840181905260808401525060019650929450925050509250925092565b6137c6613fc9565b6040805160a081018252600080825260208083018690528351828152908101845291928301919061380d565b6137fa613fc9565b8152602001906001900390816137f25790505b50815260016020820181905260409091015292915050565b600080613830613fc9565b613838613fc9565b855160009081908781108061384f57506040888203105b15613867576000888596509650965050505050612284565b60006138798a8a63ffffffff613d0516565b905060208901985061388b8a8a6135ed565b909a509450925082156138b6576138a281856101f6565b600198508997509550612284945050505050565b600089869750975097505050505050612284565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561391557816020015b613902613fc9565b8152602001906001900390816138fa5790505b50905060005b8960ff168160ff16101561397257613933898561214d565b8451859060ff861690811061394457fe5b6020908102919091010152945092508261396a575060009550869450925061397f915050565b60010161391b565b5060019550919350909150505b93509350939050565b613990613fc9565b6040805160028082526060828101909352816020015b6139ae613fc9565b8152602001906001900390816139a657905050905082816000815181106139d157fe5b602002602001018190525083816001815181106139ea57fe5b6020026020010181905250611c14613a01826126e2565b613d21565b600061027d8260600151613d97565b600061027d8260600151613db5565b606081015160009060ff16613a3b5750600161163d565b606082015160ff1660011415613a535750600061163d565b606082015160ff1660021415613aa7576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b600360ff16826060015160ff1610158015613acb57506060820151600c60ff909116105b15613b235760408201515160005b81811015613b1857613b0184604001518281518110613af457fe5b6020026020010151613a24565b613b105760009250505061163d565b600101613ad9565b50600191505061163d565b606082015160ff16606414156133de5750600061163d565b6008101590565b60008060008060008651905085811080613b5e57506020868203105b15613b725750600093508492509050612284565b613b82878763ffffffff613d0516565b600195506020870194509250612284915050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315613c14575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611c14565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214613cb7576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161027d9190613dd8565b6000613cd2613fc9565b613cdb83613d21565b90506119e581613c62565b6000613cf0613fc9565b613cfa8484613e12565b9050611c1481613c62565b60008160200183511015613d1857600080fd5b50016020015190565b613d29613fc9565b613d3282613a06565b613d78576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060613d878360400151613e31565b90506119e5818460800151613e12565b6000600c60ff831610801561027d575050600360ff91909116101590565b6000613dc082613d97565b15613dd05750600219810161163d565b50600161163d565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b613e1a613fc9565b6000613e2584613f09565b9050611c1481846101f6565b6060600882511115613e81576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613eae578160200160208202803883390190505b50805190915060005b81811015613f00576000613edd868381518110613ed057fe5b6020026020010151613327565b905080848381518110613eec57fe5b602090810291909101015250600101613eb7565b50909392505050565b6000600882511115613f59576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613f9d578181015183820152602001613f85565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001613fe361405b565b815260606020820181905260006040830181905291015290565b6040805161010081019091526000815260208101614019613fc9565b8152602001614026613fc9565b8152602001614033613fc9565b8152602001614040613fc9565b81526000602082018190526040820181905260609091015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f646553656e642076616c756520657863656564732073697a65206c696d69742c206e6f206d6573736167652073686f756c642062652073656e744c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a723158200d16f4e9d86d307f2fcc910ef7d3bcf6c3f137adc34e9e44354f3679474361f164736f6c63430005110032"

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
