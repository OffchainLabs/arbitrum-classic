// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// IOneStepProofABI is the input ABI used to generate the binding from.
const IOneStepProofABI = "[{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"bridges\",\"type\":\"address[2]\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"afterMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[4]\",\"name\":\"fields\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"bridges\",\"type\":\"address[2]\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStepDebug\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"startMachine\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"afterMachine\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IOneStepProof is an auto generated Go binding around an Ethereum contract.
type IOneStepProof struct {
	IOneStepProofCaller     // Read-only binding to the contract
	IOneStepProofTransactor // Write-only binding to the contract
	IOneStepProofFilterer   // Log filterer for contract events
}

// IOneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProofSession struct {
	Contract     *IOneStepProof    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProofCallerSession struct {
	Contract *IOneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IOneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProofTransactorSession struct {
	Contract     *IOneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IOneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProofRaw struct {
	Contract *IOneStepProof // Generic contract binding to access the raw methods on
}

// IOneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProofCallerRaw struct {
	Contract *IOneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProofTransactorRaw struct {
	Contract *IOneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProof creates a new instance of IOneStepProof, bound to a specific deployed contract.
func NewIOneStepProof(address common.Address, backend bind.ContractBackend) (*IOneStepProof, error) {
	contract, err := bindIOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProof{IOneStepProofCaller: IOneStepProofCaller{contract: contract}, IOneStepProofTransactor: IOneStepProofTransactor{contract: contract}, IOneStepProofFilterer: IOneStepProofFilterer{contract: contract}}, nil
}

// NewIOneStepProofCaller creates a new read-only instance of IOneStepProof, bound to a specific deployed contract.
func NewIOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProofCaller, error) {
	contract, err := bindIOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofCaller{contract: contract}, nil
}

// NewIOneStepProofTransactor creates a new write-only instance of IOneStepProof, bound to a specific deployed contract.
func NewIOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProofTransactor, error) {
	contract, err := bindIOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofTransactor{contract: contract}, nil
}

// NewIOneStepProofFilterer creates a new log filterer instance of IOneStepProof, bound to a specific deployed contract.
func NewIOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProofFilterer, error) {
	contract, err := bindIOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofFilterer{contract: contract}, nil
}

// bindIOneStepProof binds a generic wrapper to an already deployed contract.
func bindIOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProof *IOneStepProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProof.Contract.IOneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProof *IOneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProof.Contract.IOneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProof *IOneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProof.Contract.IOneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProof *IOneStepProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProof *IOneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProof *IOneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProof.Contract.contract.Transact(opts, method, params...)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x47dda1d6.
//
// Solidity: function executeStep(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 afterMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofCaller) ExecuteStep(opts *bind.CallOpts, bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	AfterMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	var out []interface{}
	err := _IOneStepProof.contract.Call(opts, &out, "executeStep", bridges, initialMessagesRead, accs, proof, bproof)

	outstruct := new(struct {
		Gas               uint64
		AfterMessagesRead *big.Int
		Fields            [4][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gas = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AfterMessagesRead = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fields = *abi.ConvertType(out[2], new([4][32]byte)).(*[4][32]byte)

	return *outstruct, err

}

// ExecuteStep is a free data retrieval call binding the contract method 0x47dda1d6.
//
// Solidity: function executeStep(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 afterMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofSession) ExecuteStep(bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	AfterMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _IOneStepProof.Contract.ExecuteStep(&_IOneStepProof.CallOpts, bridges, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x47dda1d6.
//
// Solidity: function executeStep(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 afterMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofCallerSession) ExecuteStep(bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	AfterMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _IOneStepProof.Contract.ExecuteStep(&_IOneStepProof.CallOpts, bridges, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0xeba67f6e.
//
// Solidity: function executeStepDebug(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofCaller) ExecuteStepDebug(opts *bind.CallOpts, bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	var out []interface{}
	err := _IOneStepProof.contract.Call(opts, &out, "executeStepDebug", bridges, initialMessagesRead, accs, proof, bproof)

	outstruct := new(struct {
		StartMachine string
		AfterMachine string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartMachine = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AfterMachine = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0xeba67f6e.
//
// Solidity: function executeStepDebug(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofSession) ExecuteStepDebug(bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _IOneStepProof.Contract.ExecuteStepDebug(&_IOneStepProof.CallOpts, bridges, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0xeba67f6e.
//
// Solidity: function executeStepDebug(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofCallerSession) ExecuteStepDebug(bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _IOneStepProof.Contract.ExecuteStepDebug(&_IOneStepProof.CallOpts, bridges, initialMessagesRead, accs, proof, bproof)
}
