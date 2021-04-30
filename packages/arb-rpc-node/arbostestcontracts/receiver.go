// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbostestcontracts

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

// ReceiverABI is the input ABI used to generate the binding from.
const ReceiverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"otherReciver\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"mutate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"other\",\"outputs\":[{\"internalType\":\"contractReceiver2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ReceiverBin is the compiled bytecode used for deploying new contracts.
var ReceiverBin = "0x6080604052600560005534801561001557600080fd5b506040516101af3803806101af8339818101604052602081101561003857600080fd5b5051600180546001600160a01b0319166001600160a01b03909216919091179055610147806100686000396000f3fe6080604052600436106100345760003560e01c80637795b5fc146100395780638529587714610043578063f8a8fd6d14610074575b600080fd5b61004161009b565b005b34801561004f57600080fd5b506100586100fc565b604080516001600160a01b039092168252519081900360200190f35b34801561008057600080fd5b5061008961010b565b60408051918252519081900360200190f35b6006600090815560015460408051631de56d7f60e21b815290516001600160a01b0390921692637795b5fc9260048084019382900301818387803b1580156100e257600080fd5b505af11580156100f6573d6000803e3d6000fd5b50505050565b6001546001600160a01b031681565b6000548156fea264697066735822122055deec57309384dc79095e74b45d2bccbf3351a12cb3401564614356a141d8f064736f6c634300060c0033"

// DeployReceiver deploys a new Ethereum contract, binding an instance of Receiver to it.
func DeployReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, otherReciver common.Address) (common.Address, *types.Transaction, *Receiver, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReceiverBin), backend, otherReciver)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Receiver{ReceiverCaller: ReceiverCaller{contract: contract}, ReceiverTransactor: ReceiverTransactor{contract: contract}, ReceiverFilterer: ReceiverFilterer{contract: contract}}, nil
}

// Receiver is an auto generated Go binding around an Ethereum contract.
type Receiver struct {
	ReceiverCaller     // Read-only binding to the contract
	ReceiverTransactor // Write-only binding to the contract
	ReceiverFilterer   // Log filterer for contract events
}

// ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiverSession struct {
	Contract     *Receiver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiverCallerSession struct {
	Contract *ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiverTransactorSession struct {
	Contract     *ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiverRaw struct {
	Contract *Receiver // Generic contract binding to access the raw methods on
}

// ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiverCallerRaw struct {
	Contract *ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiverTransactorRaw struct {
	Contract *ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiver creates a new instance of Receiver, bound to a specific deployed contract.
func NewReceiver(address common.Address, backend bind.ContractBackend) (*Receiver, error) {
	contract, err := bindReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Receiver{ReceiverCaller: ReceiverCaller{contract: contract}, ReceiverTransactor: ReceiverTransactor{contract: contract}, ReceiverFilterer: ReceiverFilterer{contract: contract}}, nil
}

// NewReceiverCaller creates a new read-only instance of Receiver, bound to a specific deployed contract.
func NewReceiverCaller(address common.Address, caller bind.ContractCaller) (*ReceiverCaller, error) {
	contract, err := bindReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverCaller{contract: contract}, nil
}

// NewReceiverTransactor creates a new write-only instance of Receiver, bound to a specific deployed contract.
func NewReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiverTransactor, error) {
	contract, err := bindReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverTransactor{contract: contract}, nil
}

// NewReceiverFilterer creates a new log filterer instance of Receiver, bound to a specific deployed contract.
func NewReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiverFilterer, error) {
	contract, err := bindReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiverFilterer{contract: contract}, nil
}

// bindReceiver binds a generic wrapper to an already deployed contract.
func bindReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Receiver *ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Receiver.Contract.ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Receiver *ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Receiver *ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Receiver *ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Receiver *ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Receiver *ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Receiver.Contract.contract.Transact(opts, method, params...)
}

// Other is a free data retrieval call binding the contract method 0x85295877.
//
// Solidity: function other() view returns(address)
func (_Receiver *ReceiverCaller) Other(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Receiver.contract.Call(opts, &out, "other")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Other is a free data retrieval call binding the contract method 0x85295877.
//
// Solidity: function other() view returns(address)
func (_Receiver *ReceiverSession) Other() (common.Address, error) {
	return _Receiver.Contract.Other(&_Receiver.CallOpts)
}

// Other is a free data retrieval call binding the contract method 0x85295877.
//
// Solidity: function other() view returns(address)
func (_Receiver *ReceiverCallerSession) Other() (common.Address, error) {
	return _Receiver.Contract.Other(&_Receiver.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Receiver *ReceiverCaller) Test(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Receiver.contract.Call(opts, &out, "test")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Receiver *ReceiverSession) Test() (*big.Int, error) {
	return _Receiver.Contract.Test(&_Receiver.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Receiver *ReceiverCallerSession) Test() (*big.Int, error) {
	return _Receiver.Contract.Test(&_Receiver.CallOpts)
}

// Mutate is a paid mutator transaction binding the contract method 0x7795b5fc.
//
// Solidity: function mutate() payable returns()
func (_Receiver *ReceiverTransactor) Mutate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver.contract.Transact(opts, "mutate")
}

// Mutate is a paid mutator transaction binding the contract method 0x7795b5fc.
//
// Solidity: function mutate() payable returns()
func (_Receiver *ReceiverSession) Mutate() (*types.Transaction, error) {
	return _Receiver.Contract.Mutate(&_Receiver.TransactOpts)
}

// Mutate is a paid mutator transaction binding the contract method 0x7795b5fc.
//
// Solidity: function mutate() payable returns()
func (_Receiver *ReceiverTransactorSession) Mutate() (*types.Transaction, error) {
	return _Receiver.Contract.Mutate(&_Receiver.TransactOpts)
}

// Receiver2ABI is the input ABI used to generate the binding from.
const Receiver2ABI = "[{\"inputs\":[],\"name\":\"mutate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Receiver2Bin is the compiled bytecode used for deploying new contracts.
var Receiver2Bin = "0x60806040526007600055348015601457600080fd5b50609a806100236000396000f3fe60806040526004361060265760003560e01c80637795b5fc14602b578063f8a8fd6d146033575b600080fd5b60316057565b005b348015603e57600080fd5b506045605e565b60408051918252519081900360200190f35b6008600055565b6000548156fea2646970667358221220f67fdf6e265b1785e0f918553ab4382067717b815bf26f41c9c6ba335319586d64736f6c634300060c0033"

// DeployReceiver2 deploys a new Ethereum contract, binding an instance of Receiver2 to it.
func DeployReceiver2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Receiver2, error) {
	parsed, err := abi.JSON(strings.NewReader(Receiver2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Receiver2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Receiver2{Receiver2Caller: Receiver2Caller{contract: contract}, Receiver2Transactor: Receiver2Transactor{contract: contract}, Receiver2Filterer: Receiver2Filterer{contract: contract}}, nil
}

// Receiver2 is an auto generated Go binding around an Ethereum contract.
type Receiver2 struct {
	Receiver2Caller     // Read-only binding to the contract
	Receiver2Transactor // Write-only binding to the contract
	Receiver2Filterer   // Log filterer for contract events
}

// Receiver2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Receiver2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Receiver2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Receiver2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Receiver2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Receiver2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Receiver2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Receiver2Session struct {
	Contract     *Receiver2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Receiver2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Receiver2CallerSession struct {
	Contract *Receiver2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Receiver2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Receiver2TransactorSession struct {
	Contract     *Receiver2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Receiver2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Receiver2Raw struct {
	Contract *Receiver2 // Generic contract binding to access the raw methods on
}

// Receiver2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Receiver2CallerRaw struct {
	Contract *Receiver2Caller // Generic read-only contract binding to access the raw methods on
}

// Receiver2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Receiver2TransactorRaw struct {
	Contract *Receiver2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiver2 creates a new instance of Receiver2, bound to a specific deployed contract.
func NewReceiver2(address common.Address, backend bind.ContractBackend) (*Receiver2, error) {
	contract, err := bindReceiver2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Receiver2{Receiver2Caller: Receiver2Caller{contract: contract}, Receiver2Transactor: Receiver2Transactor{contract: contract}, Receiver2Filterer: Receiver2Filterer{contract: contract}}, nil
}

// NewReceiver2Caller creates a new read-only instance of Receiver2, bound to a specific deployed contract.
func NewReceiver2Caller(address common.Address, caller bind.ContractCaller) (*Receiver2Caller, error) {
	contract, err := bindReceiver2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Receiver2Caller{contract: contract}, nil
}

// NewReceiver2Transactor creates a new write-only instance of Receiver2, bound to a specific deployed contract.
func NewReceiver2Transactor(address common.Address, transactor bind.ContractTransactor) (*Receiver2Transactor, error) {
	contract, err := bindReceiver2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Receiver2Transactor{contract: contract}, nil
}

// NewReceiver2Filterer creates a new log filterer instance of Receiver2, bound to a specific deployed contract.
func NewReceiver2Filterer(address common.Address, filterer bind.ContractFilterer) (*Receiver2Filterer, error) {
	contract, err := bindReceiver2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Receiver2Filterer{contract: contract}, nil
}

// bindReceiver2 binds a generic wrapper to an already deployed contract.
func bindReceiver2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Receiver2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Receiver2 *Receiver2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Receiver2.Contract.Receiver2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Receiver2 *Receiver2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver2.Contract.Receiver2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Receiver2 *Receiver2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Receiver2.Contract.Receiver2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Receiver2 *Receiver2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Receiver2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Receiver2 *Receiver2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Receiver2 *Receiver2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Receiver2.Contract.contract.Transact(opts, method, params...)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Receiver2 *Receiver2Caller) Test(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Receiver2.contract.Call(opts, &out, "test")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Receiver2 *Receiver2Session) Test() (*big.Int, error) {
	return _Receiver2.Contract.Test(&_Receiver2.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Receiver2 *Receiver2CallerSession) Test() (*big.Int, error) {
	return _Receiver2.Contract.Test(&_Receiver2.CallOpts)
}

// Mutate is a paid mutator transaction binding the contract method 0x7795b5fc.
//
// Solidity: function mutate() payable returns()
func (_Receiver2 *Receiver2Transactor) Mutate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver2.contract.Transact(opts, "mutate")
}

// Mutate is a paid mutator transaction binding the contract method 0x7795b5fc.
//
// Solidity: function mutate() payable returns()
func (_Receiver2 *Receiver2Session) Mutate() (*types.Transaction, error) {
	return _Receiver2.Contract.Mutate(&_Receiver2.TransactOpts)
}

// Mutate is a paid mutator transaction binding the contract method 0x7795b5fc.
//
// Solidity: function mutate() payable returns()
func (_Receiver2 *Receiver2TransactorSession) Mutate() (*types.Transaction, error) {
	return _Receiver2.Contract.Mutate(&_Receiver2.TransactOpts)
}
