// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onestepproof

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820557675814df39add0fa761bb6fc1ca2754618e3660ac97035c79af694b750a9264736f6c634300050d0032"

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
var DebugPrintBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201f1c2865a2b2af809ffe6caa0b1b6fbf9445248c5d42ae39aff757df53e40eb564736f6c634300050d0032"

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
var MachineBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209a127206c59692e0f0c1f1a3f9dad17d8b47066166a7f0eb8c245518a0e0643b64736f6c634300050d0032"

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
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128[2]\",\"name\":\"timeBoundsBlocks\",\"type\":\"uint128[2]\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"0b1cb787": "validateProof(bytes32,uint128[2],bytes32,bytes32,bool,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x61363b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c80630b1cb7871461003a575b600080fd5b610152600480360361018081101561005157600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091948335946020850135946040810135151594506060810135935060808101359260a08201359260c08301359267ffffffffffffffff60e08201351692919061012081019061010001356401000000008111156100dd57600080fd5b8201836020820111156100ef57600080fd5b8035906020019184600183028401116401000000008311171561011157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610164945050505050565b60408051918252519081900360200190f35b60006101c66040518061016001604052808e81526020018d81526020018c81526020018b81526020018a151581526020018981526020018881526020018781526020018681526020018567ffffffffffffffff168152602001848152506101d6565b9c9b505050505050505050505050565b600080808060606101e561343f565b6101ed61343f565b6101f6886111e9565b93995092965090945092509050600160006102108861152a565b67ffffffffffffffff168a610120015167ffffffffffffffff1614610273576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b89608001518015610287575060ff88166072145b806102a3575089608001511580156102a3575060ff8816607214155b6102f4576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420646964496e626f78496e736e20636c61696d000000000000604482015290519081900360640190fd5b60ff88166001141561033a57610333838660008151811061031157fe5b60200260200101518760018151811061032657fe5b6020026020010151611530565b915061103c565b60ff88166002141561037957610333838660008151811061035757fe5b60200260200101518760018151811061036c57fe5b6020026020010151611580565b60ff8816600314156103b857610333838660008151811061039657fe5b6020026020010151876001815181106103ab57fe5b60200260200101516115c1565b60ff8816600414156103f75761033383866000815181106103d557fe5b6020026020010151876001815181106103ea57fe5b6020026020010151611602565b60ff88166005141561043657610333838660008151811061041457fe5b60200260200101518760018151811061042957fe5b6020026020010151611653565b60ff88166006141561047557610333838660008151811061045357fe5b60200260200101518760018151811061046857fe5b60200260200101516116a4565b60ff8816600714156104b457610333838660008151811061049257fe5b6020026020010151876001815181106104a757fe5b60200260200101516116f5565b60ff8816600814156105085761033383866000815181106104d157fe5b6020026020010151876001815181106104e657fe5b6020026020010151886002815181106104fb57fe5b6020026020010151611746565b60ff88166009141561055c57610333838660008151811061052557fe5b60200260200101518760018151811061053a57fe5b60200260200101518860028151811061054f57fe5b60200260200101516117b0565b60ff8816600a141561059b57610333838660008151811061057957fe5b60200260200101518760018151811061058e57fe5b6020026020010151611809565b60ff8816601014156105da5761033383866000815181106105b857fe5b6020026020010151876001815181106105cd57fe5b602002602001015161184a565b60ff8816601114156106195761033383866000815181106105f757fe5b60200260200101518760018151811061060c57fe5b602002602001015161188b565b60ff88166012141561065857610333838660008151811061063657fe5b60200260200101518760018151811061064b57fe5b60200260200101516118cc565b60ff88166013141561069757610333838660008151811061067557fe5b60200260200101518760018151811061068a57fe5b602002602001015161190d565b60ff8816601414156106d65761033383866000815181106106b457fe5b6020026020010151876001815181106106c957fe5b602002602001015161194e565b60ff8816601514156107005761033383866000815181106106f357fe5b602002602001015161197a565b60ff88166016141561073f57610333838660008151811061071d57fe5b60200260200101518760018151811061073257fe5b60200260200101516119c0565b60ff88166017141561077e57610333838660008151811061075c57fe5b60200260200101518760018151811061077157fe5b6020026020010151611a01565b60ff8816601814156107bd57610333838660008151811061079b57fe5b6020026020010151876001815181106107b057fe5b6020026020010151611a42565b60ff8816601914156107e75761033383866000815181106107da57fe5b6020026020010151611a83565b60ff8816601a141561082657610333838660008151811061080457fe5b60200260200101518760018151811061081957fe5b6020026020010151611ab9565b60ff8816601b141561086557610333838660008151811061084357fe5b60200260200101518760018151811061085857fe5b6020026020010151611afa565b60ff88166020141561088f57610333838660008151811061088257fe5b6020026020010151611b3b565b60ff8816602114156108b95761033383866000815181106108ac57fe5b6020026020010151611b57565b60ff8816603014156108e35761033383866000815181106108d657fe5b6020026020010151611b72565b60ff8816603114156108f85761033383611b7a565b60ff88166032141561090d5761033383611b9b565b60ff88166033141561093757610333838660008151811061092a57fe5b6020026020010151611bb4565b60ff88166034141561096157610333838660008151811061095457fe5b6020026020010151611bcd565b60ff8816603514156109a057610333838660008151811061097e57fe5b60200260200101518760018151811061099357fe5b6020026020010151611be3565b60ff8816603614156109b55761033383611c2b565b60ff8816603714156109cf57610333838560000151611c5d565b60ff8816603814156109f95761033383866000815181106109ec57fe5b6020026020010151611c6f565b60ff881660391415610a8657610a0d6134a0565b610a1c8b610140015188611c81565b919950975090508715610a605760405162461bcd60e51b81526004018080602001828103825260218152602001806135e66021913960400191505060405180910390fd5b610a70858263ffffffff611e0b16565b610a80848263ffffffff611e2d16565b5061103c565b60ff8816603a1415610a9b5761033383611e4a565b60ff8816603b1415610aac5761103c565b60ff8816603c1415610ac15761033383611e6a565b60ff8816603d1415610aeb576103338386600081518110610ade57fe5b6020026020010151611e83565b60ff881660401415610b15576103338386600081518110610b0857fe5b6020026020010151611eb1565b60ff881660411415610b54576103338386600081518110610b3257fe5b602002602001015187600181518110610b4757fe5b6020026020010151611ed3565b60ff881660421415610ba8576103338386600081518110610b7157fe5b602002602001015187600181518110610b8657fe5b602002602001015188600281518110610b9b57fe5b6020026020010151611f05565b60ff881660431415610be7576103338386600081518110610bc557fe5b602002602001015187600181518110610bda57fe5b6020026020010151611f47565b60ff881660441415610c3b576103338386600081518110610c0457fe5b602002602001015187600181518110610c1957fe5b602002602001015188600281518110610c2e57fe5b6020026020010151611f59565b60ff881660501415610c7a576103338386600081518110610c5857fe5b602002602001015187600181518110610c6d57fe5b6020026020010151611f7b565b60ff881660511415610cce576103338386600081518110610c9757fe5b602002602001015187600181518110610cac57fe5b602002602001015188600281518110610cc157fe5b6020026020010151611ff1565b60ff881660521415610cf8576103338386600081518110610ceb57fe5b6020026020010151612069565b60ff881660601415610d0d576103338361152a565b60ff881660611415610e0b57610d378386600081518110610d2a57fe5b602002602001015161209c565b90925090508115610e02578961010001518a60e001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610db75760405162461bcd60e51b815260040180806020018281038252602581526020018061359a6025913960400191505060405180910390fd5b8960c001518a60a0015114610dfd5760405162461bcd60e51b81526004018080602001828103825260278152602001806135bf6027913960400191505060405180910390fd5b610e06565b5060005b61103c565b60ff881660701415610efb57610e358386600081518110610e2857fe5b60200260200101516120c0565b90925090508115610e02578960c001518a60a001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610eb45760405162461bcd60e51b815260040180806020018281038252602981526020018061352a6029913960400191505060405180910390fd5b8961010001518a60e0015114610dfd5760405162461bcd60e51b81526004018080602001828103825260268152602001806135536026913960400191505060405180910390fd5b60ff881660711415610fb6576040805160028082526060828101909352816020015b610f256134a0565b815260200190600190039081610f1d57505060208c0151909150610f599060005b60200201516001600160801b03166120da565b81600081518110610f6657fe5b6020026020010181905250610f858b60200151600160028110610f4657fe5b81600181518110610f9257fe5b6020026020010181905250610a80610fa982612158565b859063ffffffff611e2d16565b60ff881660721415611012576103338386600081518110610fd357fe5b602002602001015160405180602001604052808e604001518152508d60200151600060028110610fff57fe5b60200201516001600160801b0316612208565b60ff881660731415611027576000915061103c565b60ff88166074141561103c5761103c83612297565b806110ce578960c001518a60a00151146110875760405162461bcd60e51b81526004018080602001828103825260278152602001806135bf6027913960400191505060405180910390fd5b8961010001518a60e00151146110ce5760405162461bcd60e51b81526004018080602001828103825260268152602001806135536026913960400191505060405180910390fd5b816111305760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a084015151141561112857611123836122a1565b611130565b60a083015183525b611139846122ab565b8a51146111775760405162461bcd60e51b81526004018080602001828103825260228152602001806135086022913960400191505060405180910390fd5b611180836122ab565b8a60600151146111d7576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b600060606111f561343f565b6111fd61343f565b6000808061120961343f565b61121281612340565b6112218961014001518461234a565b909450909250905061123161343f565b61123a8261244f565b905060008a6101400151858151811061124f57fe5b602001015160f81c60f81b60f81c905060008b6101400151866001018151811061127557fe5b016020015160f81c9050600061128a826124ad565b90506060816040519080825280602002602001820160405280156112c857816020015b6112b56134a0565b8152602001906001900390816112ad5790505b5090506002880197508360ff16600014806112e657508360ff166001145b611337576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff841661136557604051806020016040528061135c858960000151600001516124c7565b90528652611432565b61136d6134a0565b61137c8f61014001518a611c81565b909a50909850905087156113d7576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b82156113fb5780826000815181106113eb57fe5b602002602001018190525061140b565b61140b868263ffffffff611e2d16565b604051806020016040528061142c866114238561250e565b518b5151612644565b90528752505b60ff84165b828110156114c65761144e8f61014001518a611c81565b845185908590811061145c57fe5b60209081029190910101529950975087156114be576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611437565b815115611513575060005b8460ff168251038110156115135761150b8282600185510303815181106114f457fe5b602002602001015188611e2d90919063ffffffff16565b6001016114d1565b50919d919c50939a50919850939650945050505050565b50600190565b600061153b83612696565b158061154d575061154b82612696565b155b1561155a57506000611579565b82518251808201611571878263ffffffff6126a116565b600193505050505b9392505050565b600061158b83612696565b158061159d575061159b82612696565b155b156115aa57506000611579565b82518251808202611571878263ffffffff6126a116565b60006115cc83612696565b15806115de57506115dc82612696565b155b156115eb57506000611579565b82518251808203611571878263ffffffff6126a116565b600061160d83612696565b158061161f575061161d82612696565b155b1561162c57506000611579565b825182518061164057600092505050611579565b808204611571878263ffffffff6126a116565b600061165e83612696565b1580611670575061166e82612696565b155b1561167d57506000611579565b825182518061169157600092505050611579565b808205611571878263ffffffff6126a116565b60006116af83612696565b15806116c157506116bf82612696565b155b156116ce57506000611579565b82518251806116e257600092505050611579565b808206611571878263ffffffff6126a116565b600061170083612696565b1580611712575061171082612696565b155b1561171f57506000611579565b825182518061173357600092505050611579565b808207611571878263ffffffff6126a116565b600061175184612696565b1580611763575061176183612696565b155b15611770575060006117a8565b8351835183518061178757600093505050506117a8565b600081838508905061179f898263ffffffff6126a116565b60019450505050505b949350505050565b60006117bb84612696565b15806117cd57506117cb83612696565b155b156117da575060006117a8565b835183518351806117f157600093505050506117a8565b600081838509905061179f898263ffffffff6126a116565b600061181483612696565b1580611826575061182482612696565b155b1561183357506000611579565b8251825180820a611571878263ffffffff6126a116565b600061185583612696565b1580611867575061186582612696565b155b1561187457506000611579565b82518251808210611571878263ffffffff6126a116565b600061189683612696565b15806118a857506118a682612696565b155b156118b557506000611579565b82518251808211611571878263ffffffff6126a116565b60006118d783612696565b15806118e957506118e782612696565b155b156118f657506000611579565b82518251808212611571878263ffffffff6126a116565b600061191883612696565b158061192a575061192882612696565b155b1561193757506000611579565b82518251808213611571878263ffffffff6126a116565b6000611970610fa961195f8461250e565b516119698661250e565b51146126b5565b5060019392505050565b600061198582612696565b61199f5761199a83600063ffffffff6126a116565b6119b6565b815180156119b3858263ffffffff6126a116565b50505b5060015b92915050565b60006119cb83612696565b15806119dd57506119db82612696565b155b156119ea57506000611579565b82518251808216611571878263ffffffff6126a116565b6000611a0c83612696565b1580611a1e5750611a1c82612696565b155b15611a2b57506000611579565b82518251808217611571878263ffffffff6126a116565b6000611a4d83612696565b1580611a5f5750611a5d82612696565b155b15611a6c57506000611579565b82518251808218611571878263ffffffff6126a116565b6000611a8e82612696565b611a9a575060006119ba565b81518019611aae858263ffffffff6126a116565b506001949350505050565b6000611ac483612696565b1580611ad65750611ad482612696565b155b15611ae357506000611579565b8251825181811a611571878263ffffffff6126a116565b6000611b0583612696565b1580611b175750611b1582612696565b155b15611b2457506000611579565b8251825181810b611571878263ffffffff6126a116565b60006119b6611b498361250e565b51849063ffffffff6126a116565b60006119b6611b65836126de565b849063ffffffff611e2d16565b600192915050565b6000611b9382608001518361276790919063ffffffff16565b506001919050565b6000611b9382606001518361276790919063ffffffff16565b6000611bbf8261250e565b606084015250600192915050565b6000611bd88261250e565b835250600192915050565b6000611bee83612775565b611bfa57506000611579565b611c0382612696565b611c0f57506000611579565b81511561197057611c1f8361250e565b84525060019392505050565b6000611b93611c50611c43611c3e612782565b61250e565b51602085015151146126b5565b839063ffffffff611e2d16565b60006119b6838363ffffffff61276716565b60006119b6838363ffffffff611e0b16565b600080611c8c6134a0565b84518410611ce1576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110611cf457fe5b016020015160019092019160f81c90506000611d0e6134ce565b60ff8316611d4257611d2089856127ff565b9094509150600084611d31846120da565b91985096509450611e049350505050565b60ff831660011415611d6957611d588985612826565b9094509050600084611d3183612993565b60ff831660021415611d9057611d7f89856127ff565b9094509150600084611d31846129f3565b600360ff841610801590611da75750600c60ff8416105b15611de457600219830160606000611dc0838d89612a71565b909850925090508087611dd284612158565b99509950995050505050505050611e04565b8260ff16612710016000611df860006120da565b91985096509450505050505b9250925092565b611e218260400151611e1c8361250e565b612b2c565b82604001819052505050565b611e3e8260200151611e1c8361250e565b82602001819052505050565b6000611b93611c50611e5d611c3e612782565b51604085015151146126b5565b6000611b938260a001518361276790919063ffffffff16565b6000611e8e82612775565b611e9a575060006119ba565b611ea38261250e565b60a084015250600192915050565b6000611ec3838363ffffffff611e2d16565b6119b6838363ffffffff611e2d16565b6000611ee5848363ffffffff611e2d16565b611ef5848463ffffffff611e2d16565b611970848363ffffffff611e2d16565b6000611f17858363ffffffff611e2d16565b611f27858463ffffffff611e2d16565b611f37858563ffffffff611e2d16565b611aae858363ffffffff611e2d16565b6000611ef5848463ffffffff611e2d16565b6000611f6b858563ffffffff611e2d16565b611f37858463ffffffff611e2d16565b6000611f8683612696565b1580611f985750611f9682612be2565b155b15611fa557506000611579565b611fae82612bf1565b60ff16836000015110611fc357506000611579565b6119708260400151846000015181518110611fda57fe5b602002602001015185611e2d90919063ffffffff16565b6000611ffc83612be2565b158061200e575061200c84612696565b155b1561201b575060006117a8565b61202483612bf1565b60ff16846000015110612039575060006117a8565b81836040015185600001518151811061204e57fe5b6020908102919091010152611aae858463ffffffff611e2d16565b600061207482612be2565b612080575060006119ba565b6119b661208c83612bf1565b849060ff1663ffffffff6126a116565b6000806120a76134f5565b6120b08461250e565b51600193509150505b9250929050565b60008060016120ce8461250e565b51909590945092505050565b6120e26134a0565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612147565b6121346134a0565b81526020019060019003908161212c5790505b508152600060209091015292915050565b6121606134a0565b61216a8251612c00565b6121bb576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b600061221384612696565b61221f575060006117a8565b8351821080156122365750612232612c07565b8351145b612287576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b611aae858463ffffffff61276716565b600260c090910152565b600160c090910152565b600060028260c0015114156122c2575060006111e4565b60018260c0015114156122d7575060016111e4565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101206111e4565b600060c090910152565b60008061235561343f565b61235d61343f565b600060c082018190526123708787612c7a565b84529650905080156123885793508492509050611e04565b6123928787612c7a565b60208501529650905080156123ad5793508492509050611e04565b6123b78787612c7a565b60408501529650905080156123d25793508492509050611e04565b6123dc8787612c7a565b60608501529650905080156123f75793508492509050611e04565b6124018787612c7a565b608085015296509050801561241c5793508492509050611e04565b6124268787612c7a565b60a08501529650905080156124415793508492509050611e04565b506000969495509392505050565b61245761343f565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b60008060006124be8460ff16612cb8565b50949350505050565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6125166134f5565b6060820151600c60ff90911610612568576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661259557604051806020016040528061258c846000015161314b565b905290506111e4565b606082015160ff16600114156125dc57604051806020016040528061258c84602001516000015185602001516040015186602001516060015187602001516020015161316f565b606082015160ff166002141561260157506040805160208101909152815181526111e4565b600360ff16826060015160ff161015801561262557506060820151600c60ff909116105b1561264257604051806020016040528061258c8460400151613217565bfe5b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6060015160ff161590565b611e3e8260200151611e1c611c3e846120da565b6126bd6134a0565b81156126d4576126cd60016120da565b90506111e4565b6126cd60006120da565b6126e66134a0565b816060015160ff166002141561272d5760405162461bcd60e51b81526004018080602001828103825260218152602001806135796021913960400191505060405180910390fd5b606082015160ff16612743576126cd60006120da565b816060015160ff166001141561275d576126cd60016120da565b6126cd60036120da565b611e3e826020015182612b2c565b6060015160ff1660011490565b61278a6134a0565b6040805160808082018352600080835283519182018452808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916127ef565b6127dc6134a0565b8152602001906001900390816127d45790505b5081526003602090910152905090565b6000808281612814868363ffffffff61336316565b60209290920196919550909350505050565b60006128306134ce565b6000839050600085828151811061284357fe5b602001015160f81c60f81b60f81c90508180600101925050600086838151811061286957fe5b016020015160019384019360f89190911c915060009060ff841614156129075760006128936134a0565b61289d8a87611c81565b909750909250905081156128f8576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b6129018161250e565b51925050505b6000612919898663ffffffff61336316565b90506020850194508360ff166001141561295e576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506120b99050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b61299b6134a0565b6040805160808101825260008082526020808301869052835182815290810184529192830191906129e2565b6129cf6134a0565b8152602001906001900390816129c75790505b508152600160209091015292915050565b6129fb6134a0565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612a60565b612a4d6134a0565b815260200190600190039081612a455790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015612abc57816020015b612aa96134a0565b815260200190600190039081612aa15790505b50905060005b8960ff168160ff161015612b1657612ada8985611c81565b8451859060ff8616908110612aeb57fe5b6020908102919091010152945092508215612b0e57509094509092509050612b23565b600101612ac2565b5060009550919350909150505b93509350939050565b612b346134f5565b6040805160028082526060828101909352816020015b612b526134f5565b815260200190600190039081612b4a5790505090508281600081518110612b7557fe5b60200260200101819052508381600181518110612b8e57fe5b60200260200101819052506040518060200160405280612bd86040518060400160405280612bbf88600001516129f3565b8152602001612bd189600001516129f3565b905261337f565b9052949350505050565b60006119ba82606001516133fe565b60006119ba826060015161341c565b6008101590565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b83811015612c53578181015183820152602001612c3b565b50505050905001925050506040516020818303038152906040528051906020012091505090565b600080612c856134f5565b836000612c98878363ffffffff61336316565b604080516020808201909252918252600099930197509550909350505050565b6000806001831415612cd05750600290506001613146565b6002831415612ce55750600290506001613146565b6003831415612cfa5750600290506001613146565b6004831415612d0f5750600290506001613146565b6005831415612d245750600290506001613146565b6006831415612d395750600290506001613146565b6007831415612d4e5750600290506001613146565b6008831415612d635750600390506001613146565b6009831415612d785750600390506001613146565b600a831415612d8d5750600290506001613146565b6010831415612da25750600290506001613146565b6011831415612db75750600290506001613146565b6012831415612dcc5750600290506001613146565b6013831415612de15750600290506001613146565b6014831415612df65750600290506001613146565b6015831415612e0a57506001905080613146565b6016831415612e1f5750600290506001613146565b6017831415612e345750600290506001613146565b6018831415612e495750600290506001613146565b6019831415612e5d57506001905080613146565b601a831415612e725750600290506001613146565b601b831415612e875750600290506001613146565b6020831415612e9b57506001905080613146565b6021831415612eaf57506001905080613146565b6030831415612ec45750600190506000613146565b6031831415612ed95750600090506001613146565b6032831415612eee5750600090506001613146565b6033831415612f035750600190506000613146565b6034831415612f185750600190506000613146565b6035831415612f2d5750600290506000613146565b6036831415612f425750600090506001613146565b6037831415612f575750600090506001613146565b6038831415612f6c5750600190506000613146565b6039831415612f815750600090506001613146565b603a831415612f965750600090506001613146565b603b831415612faa57506000905080613146565b603c831415612fbf5750600090506001613146565b603d831415612fd45750600190506000613146565b6040831415612fe95750600190506002613146565b6041831415612ffe5750600290506003613146565b60428314156130135750600390506004613146565b604383141561302757506002905080613146565b604483141561303b57506003905080613146565b60508314156130505750600290506001613146565b60518314156130655750600390506001613146565b605283141561307957506001905080613146565b606083141561308d57506000905080613146565b60618314156130a25750600190506000613146565b60708314156130b75750600190506000613146565b60718314156130cc5750600090506001613146565b60728314156130e057506001905080613146565b60738314156130f457506000905080613146565b607483141561310857506000905080613146565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b600083156131c9575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206117a8565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6000600882511115613267576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613294578160200160208202803883390190505b50805190915060005b818110156132f0576132ad6134f5565b6132c98683815181106132bc57fe5b602002602001015161250e565b905080600001518483815181106132dc57fe5b60209081029190910101525060010161329d565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613339578181015183820152602001613321565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6000816020018351101561337657600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b6133a26134a0565b81526020019060019003908161339a575050805190915060005b818110156133f4578481600281106133d057fe5b60200201518382815181106133e157fe5b60209081029190910101526001016133bc565b506117a882613217565b6000600c60ff83161080156119ba575050600360ff91909116101590565b6000613427826133fe565b15613437575060021981016111e4565b5060016111e4565b6040518060e001604052806134526134f5565b815260200161345f6134f5565b815260200161346c6134f5565b81526020016134796134f5565b81526020016134866134f5565b81526020016134936134f5565b8152602001600081525090565b6040518060800160405280600081526020016134ba6134ce565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f64654c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72315820be339cddae029ce5d5555e4f053d0783d61443c1c5bb5ff3086cb1f22461857a64736f6c634300050d0032"

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

// ValidateProof is a free data retrieval call binding the contract method 0x0b1cb787.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[2] timeBoundsBlocks, bytes32 beforeInbox, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, beforeHash [32]byte, timeBoundsBlocks [2]*big.Int, beforeInbox [32]byte, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", beforeHash, timeBoundsBlocks, beforeInbox, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0x0b1cb787.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[2] timeBoundsBlocks, bytes32 beforeInbox, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(beforeHash [32]byte, timeBoundsBlocks [2]*big.Int, beforeInbox [32]byte, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBoundsBlocks, beforeInbox, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0x0b1cb787.
//
// Solidity: function validateProof(bytes32 beforeHash, uint128[2] timeBoundsBlocks, bytes32 beforeInbox, bytes32 afterHash, bool didInboxInsn, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(beforeHash [32]byte, timeBoundsBlocks [2]*big.Int, beforeInbox [32]byte, afterHash [32]byte, didInboxInsn bool, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBoundsBlocks, beforeInbox, afterHash, didInboxInsn, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582049ce534ba6c84dba84c2fb5f505617f971c8a248e7ac44e49250f60527d06f3364736f6c634300050d0032"

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
