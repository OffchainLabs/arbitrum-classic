// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machinetester

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

// MachineTesterABI is the input ABI used to generate the binding from.
const MachineTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data2\",\"type\":\"bytes\"}],\"name\":\"addStackVal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeMachine\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MachineTesterFuncSigs maps the 4-byte function signature to its string representation.
var MachineTesterFuncSigs = map[string]string{
	"5f098d7f": "addStackVal(bytes,bytes)",
	"5270f3e9": "deserializeMachine(bytes)",
}

// MachineTesterBin is the compiled bytecode used for deploying new contracts.
var MachineTesterBin = "0x608060405234801561001057600080fd5b506113e7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635270f3e91461003b5780635f098d7f146100f3575b600080fd5b6100e16004803603602081101561005157600080fd5b81019060208101813564010000000081111561006c57600080fd5b82018360208201111561007e57600080fd5b803590602001918460018302840111640100000000831117156100a057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610220945050505050565b60408051918252519081900360200190f35b6100e16004803603604081101561010957600080fd5b81019060208101813564010000000081111561012457600080fd5b82018360208201111561013657600080fd5b8035906020019184600183028401116401000000008311171561015857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101ab57600080fd5b8201836020820111156101bd57600080fd5b803590602001918460018302840111640100000000831117156101df57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610253945050505050565b600080600061022d611301565b610238856000610333565b919450925090506102488161046f565b93505050505b919050565b600080600061026061135f565b61026861135f565b610273876000610533565b91955093509150836102bf576040805162461bcd60e51b815260206004820152601060248201526f1d985b1d594c481a5b98dbdc9c9958dd60821b604482015290519081900360640190fd5b6102ca866000610533565b9195509350905083610316576040805162461bcd60e51b815260206004820152601060248201526f1d985b1d594c881a5b98dbdc9c9958dd60821b604482015290519081900360640190fd5b6103286103238383610650565b6106d6565b979650505050505050565b60008061033e611301565b610346611301565b600060e0820181905261035987876107db565b845296509050806103735750600093508492509050610468565b61037d878761082f565b6020850152965090508061039a5750600093508492509050610468565b6103a4878761082f565b604085015296509050806103c15750600093508492509050610468565b6103cb8787610533565b606085015296509050806103e85750600093508492509050610468565b6103f28787610533565b6080850152965090508061040f5750600093508492509050610468565b61041987876108d4565b60a085015296509050806104365750600093508492509050610468565b61044087876107db565b60c0850152965090508061045d5750600093508492509050610468565b506001935084925090505b9250925092565b600060028260e0015114156104865750600061024e565b60018260e00151141561049b5750600161024e565b815160208301516104ab906106d6565b6104b884604001516106d6565b6104c585606001516106d6565b6104d286608001516106d6565b8660a001518760c001516040516020018088815260200187815260200186815260200185815260200184815260200183815260200182815260200197505050505050505060405160208183030381529060405280519060200120905061024e565b60008061053e61135f565b8451841061055e576000846105536000610927565b925092509250610468565b600080859050600087828151811061057257fe5b016020015160019092019160f81c90506000816105b45761059389846108d4565b9195509350905083836105a583610927565b96509650965050505050610468565b60ff8216600114156105ca576105a589846109d9565b60ff8216600214156105e0576105a5898461082f565b600360ff8316108015906105f75750600c60ff8316105b15610631576002198201606061060e828c87610ac9565b91975095509050858561062083610b87565b985098509850505050505050610468565b60008061063e6000610927565b91985096509450505050509250925092565b61065861135f565b6040805160028082526060828101909352816020015b61067661135f565b81526020019060019003908161066e579050509050828160008151811061069957fe5b602002602001018190525083816001815181106106b257fe5b60200260200101819052506106ce6106c982610b87565b610c9e565b949350505050565b606081015160009060ff166106f75781516106f090610ddd565b905061024e565b606082015160ff1660011415610714576106f08260200151610e01565b606082015160ff166002141561073557815160808301516106f09190610ee0565b600360ff16826060015160ff161015801561075957506060820151600c60ff909116105b156107825761076661135f565b61076f83610c9e565b905061077a816106d6565b91505061024e565b606082015160ff166064141561079a5750805161024e565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b600080600080600086519050858110806107f757506020868203105b1561080b5750600093508492509050610468565b61081b878763ffffffff610f1a16565b600195506020870194509250610468915050565b60008061083a61135f565b61084261135f565b855160009081908781108061085957506040888203105b15610871576000888596509650965050505050610468565b60006108838a8a63ffffffff610f1a16565b90506020890198506108958a8a6108d4565b909a509450925082156108c0576108ac8185610f36565b600198508997509550610468945050505050565b600089869750975097505050505050610468565b60008060008085519050848110806108ee57506020858203105b15610903575060009250839150829050610468565b600160208601610919888863ffffffff610f1a16565b935093509350509250925092565b61092f61135f565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610985565b61097261135f565b81526020019060019003908161096a5790505b509052815260408051600080825260208281019093529190920191906109c1565b6109ae61135f565b8152602001906001900390816109a65790505b50815260006020820152600160409091015292915050565b6000806109e461135f565b600084905060008682815181106109f757fe5b602001015160f81c60f81b60f81c905081806001019250506000878381518110610a1d57fe5b016020015160019093019260f81c9050610a3561135f565b8260ff1660011415610a76576000610a4d8a86610533565b9096509250905080610a7457600089610a64610fe7565b9750975097505050505050610468565b505b6000610a888a8663ffffffff610f1a16565b90506020850194508360ff1660011415610aaa57600185610a6485848661109b565b600185610ab7858461111f565b97509750975050505050509250925092565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610b1457816020015b610b0161135f565b815260200190600190039081610af95790505b50905060005b8960ff168160ff161015610b7157610b328985610533565b8451859060ff8616908110610b4357fe5b60209081029190910101529450925082610b695750600095508694509250610b7e915050565b600101610b1a565b5060019550919350909150505b93509350939050565b610b8f61135f565b610b998251611181565b610bea576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610c2157838181518110610c0457fe5b602002602001015160800151820191508080600101915050610bef565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610c7b565b610c6861135f565b815260200190600190039081610c605790505b509052815260208101859052935160030160ff1660408501526060909301525090565b610ca661135f565b610caf82611188565b610cf5576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b60088260400151511115610d47576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060826040015151604051908082528060200260200182016040528015610d78578160200160208202803883390190505b50805190915060005b81811015610dce576000610dab86604001518381518110610d9e57fe5b60200260200101516106d6565b905080848381518110610dba57fe5b602090810291909101015250600101610d81565b506106ce82856080015161119d565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600282604001515110610e1257fe5b604082015151610e6c5750805160208083015160408051600160f81b8185015260f89490941b6001600160f81b0319166021850152602280850192909252805180850390920182526042909301909252815191012061024e565b60018260000151610e878460400151600081518110610d9e57fe5b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b60008160200183511015610f2d57600080fd5b50016020015190565b610f3e61135f565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191610f94565b610f8161135f565b815260200190600190039081610f795790505b50905281526040805160008082526020828101909352919092019190610fd0565b610fbd61135f565b815260200190600190039081610fb55790505b508152600260208201526040019290925250919050565b610fef61135f565b6040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190611048565b61103561135f565b81526020019060019003908161102d5790505b50905281526040805160008082526020828101909352919092019190611084565b61107161135f565b8152602001906001900390816110695790505b508152600360208201526001604090910152905090565b6110a361135f565b604080516001808252818301909252606091816020015b6110c261135f565b8152602001906001900390816110ba57905050905082816000815181106110e557fe5b602002602001018190525061111660405180606001604052808760ff168152602001868152602001838152506111bc565b95945050505050565b61112761135f565b6040805160608101825260ff85168152602080820185905282516000808252918101845261117a93830191611172565b61115f61135f565b8152602001906001900390816111575790505b5090526111bc565b9392505050565b6008101590565b60006111978260600151611223565b92915050565b6111a561135f565b60006111b084611241565b90506106ce8184610f36565b6111c461135f565b6040805160a081018252600080825260208083018690528351828152908101845291928301919061120b565b6111f861135f565b8152602001906001900390816111f05790505b50815260016020820181905260409091015292915050565b6000600c60ff8316108015611197575050600360ff91909116101590565b6000600882511115611291576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156112d55781810151838201526020016112bd565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b604080516101008101909152600081526020810161131d61135f565b815260200161132a61135f565b815260200161133761135f565b815260200161134461135f565b81526000602082018190526040820181905260609091015290565b6040518060a0016040528060008152602001611379611393565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fea265627a7a72315820504c116967b35bb9d701d561e1185f65ae331c944d14bc51e8e435d19612ebaa64736f6c63430005110032"

// DeployMachineTester deploys a new Ethereum contract, binding an instance of MachineTester to it.
func DeployMachineTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MachineTester, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MachineTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MachineTester{MachineTesterCaller: MachineTesterCaller{contract: contract}, MachineTesterTransactor: MachineTesterTransactor{contract: contract}, MachineTesterFilterer: MachineTesterFilterer{contract: contract}}, nil
}

// MachineTester is an auto generated Go binding around an Ethereum contract.
type MachineTester struct {
	MachineTesterCaller     // Read-only binding to the contract
	MachineTesterTransactor // Write-only binding to the contract
	MachineTesterFilterer   // Log filterer for contract events
}

// MachineTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineTesterSession struct {
	Contract     *MachineTester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineTesterCallerSession struct {
	Contract *MachineTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MachineTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineTesterTransactorSession struct {
	Contract     *MachineTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MachineTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineTesterRaw struct {
	Contract *MachineTester // Generic contract binding to access the raw methods on
}

// MachineTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineTesterCallerRaw struct {
	Contract *MachineTesterCaller // Generic read-only contract binding to access the raw methods on
}

// MachineTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineTesterTransactorRaw struct {
	Contract *MachineTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachineTester creates a new instance of MachineTester, bound to a specific deployed contract.
func NewMachineTester(address common.Address, backend bind.ContractBackend) (*MachineTester, error) {
	contract, err := bindMachineTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MachineTester{MachineTesterCaller: MachineTesterCaller{contract: contract}, MachineTesterTransactor: MachineTesterTransactor{contract: contract}, MachineTesterFilterer: MachineTesterFilterer{contract: contract}}, nil
}

// NewMachineTesterCaller creates a new read-only instance of MachineTester, bound to a specific deployed contract.
func NewMachineTesterCaller(address common.Address, caller bind.ContractCaller) (*MachineTesterCaller, error) {
	contract, err := bindMachineTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineTesterCaller{contract: contract}, nil
}

// NewMachineTesterTransactor creates a new write-only instance of MachineTester, bound to a specific deployed contract.
func NewMachineTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineTesterTransactor, error) {
	contract, err := bindMachineTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineTesterTransactor{contract: contract}, nil
}

// NewMachineTesterFilterer creates a new log filterer instance of MachineTester, bound to a specific deployed contract.
func NewMachineTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineTesterFilterer, error) {
	contract, err := bindMachineTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineTesterFilterer{contract: contract}, nil
}

// bindMachineTester binds a generic wrapper to an already deployed contract.
func bindMachineTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineTester *MachineTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MachineTester.Contract.MachineTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineTester *MachineTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineTester.Contract.MachineTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineTester *MachineTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineTester.Contract.MachineTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineTester *MachineTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MachineTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineTester *MachineTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineTester *MachineTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineTester.Contract.contract.Transact(opts, method, params...)
}

// AddStackVal is a free data retrieval call binding the contract method 0x5f098d7f.
//
// Solidity: function addStackVal(bytes data1, bytes data2) pure returns(bytes32)
func (_MachineTester *MachineTesterCaller) AddStackVal(opts *bind.CallOpts, data1 []byte, data2 []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MachineTester.contract.Call(opts, out, "addStackVal", data1, data2)
	return *ret0, err
}

// AddStackVal is a free data retrieval call binding the contract method 0x5f098d7f.
//
// Solidity: function addStackVal(bytes data1, bytes data2) pure returns(bytes32)
func (_MachineTester *MachineTesterSession) AddStackVal(data1 []byte, data2 []byte) ([32]byte, error) {
	return _MachineTester.Contract.AddStackVal(&_MachineTester.CallOpts, data1, data2)
}

// AddStackVal is a free data retrieval call binding the contract method 0x5f098d7f.
//
// Solidity: function addStackVal(bytes data1, bytes data2) pure returns(bytes32)
func (_MachineTester *MachineTesterCallerSession) AddStackVal(data1 []byte, data2 []byte) ([32]byte, error) {
	return _MachineTester.Contract.AddStackVal(&_MachineTester.CallOpts, data1, data2)
}

// DeserializeMachine is a free data retrieval call binding the contract method 0x5270f3e9.
//
// Solidity: function deserializeMachine(bytes data) pure returns(bytes32)
func (_MachineTester *MachineTesterCaller) DeserializeMachine(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MachineTester.contract.Call(opts, out, "deserializeMachine", data)
	return *ret0, err
}

// DeserializeMachine is a free data retrieval call binding the contract method 0x5270f3e9.
//
// Solidity: function deserializeMachine(bytes data) pure returns(bytes32)
func (_MachineTester *MachineTesterSession) DeserializeMachine(data []byte) ([32]byte, error) {
	return _MachineTester.Contract.DeserializeMachine(&_MachineTester.CallOpts, data)
}

// DeserializeMachine is a free data retrieval call binding the contract method 0x5270f3e9.
//
// Solidity: function deserializeMachine(bytes data) pure returns(bytes32)
func (_MachineTester *MachineTesterCallerSession) DeserializeMachine(data []byte) ([32]byte, error) {
	return _MachineTester.Contract.DeserializeMachine(&_MachineTester.CallOpts, data)
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
