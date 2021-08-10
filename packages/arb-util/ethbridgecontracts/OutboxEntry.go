// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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

// OutboxEntryABI is the input ABI used to generate the binding from.
const OutboxEntryABI = "[{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_numInBatch\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"spendOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"spentOutput\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OutboxEntryBin is the compiled bytecode used for deploying new contracts.
var OutboxEntryBin = "0x608060405234801561001057600080fd5b506000805460ff191660011790556104a88061002d6000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80635780e4e71461007257806357d61c0b1461008c5780635b36c66b146100af5780636f791d29146100d457806383197ef0146100f05780639db9af81146100f8578063ebf0c71714610115575b600080fd5b61007a61011d565b60408051918252519081900360200190f35b61007a600480360360408110156100a257600080fd5b5080359060200135610123565b6100d2600480360360408110156100c557600080fd5b508035906020013561023b565b005b6100dc610330565b604080519115158252519081900360200190f35b6100d2610339565b6100dc6004803603602081101561010e57600080fd5b503561039a565b61007a6103af565b60025481565b6000805461010090046001600160a01b0316331461017a576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b60008281526003602052604090205460ff16156101ce576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b600154831461020f576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b506000818152600360205260409020805460ff1916600117905560028054600019019081905592915050565b60005461010090046001600160a01b03161561028d576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600154156102d1576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b8161030e576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b60008054610100600160a81b0319163361010002179055600191909155600255565b60005460ff1690565b60005461010090046001600160a01b0316331461038f576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b610398336103b5565b565b60036020526000908152604090205460ff1681565b60015481565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff16156104655760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561042a578181015183820152602001610412565b50505050905090810190601f1680156104575780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316fffea264697066735822122060ded5f2eed3062d72d682d9c9161508dc45514a0b9865f849c11a588daae5e864736f6c634300060b0033"

// DeployOutboxEntry deploys a new Ethereum contract, binding an instance of OutboxEntry to it.
func DeployOutboxEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OutboxEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// OutboxEntry is an auto generated Go binding around an Ethereum contract.
type OutboxEntry struct {
	OutboxEntryCaller     // Read-only binding to the contract
	OutboxEntryTransactor // Write-only binding to the contract
	OutboxEntryFilterer   // Log filterer for contract events
}

// OutboxEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxEntrySession struct {
	Contract     *OutboxEntry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxEntryCallerSession struct {
	Contract *OutboxEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OutboxEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxEntryTransactorSession struct {
	Contract     *OutboxEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OutboxEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxEntryRaw struct {
	Contract *OutboxEntry // Generic contract binding to access the raw methods on
}

// OutboxEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxEntryCallerRaw struct {
	Contract *OutboxEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxEntryTransactorRaw struct {
	Contract *OutboxEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutboxEntry creates a new instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntry(address common.Address, backend bind.ContractBackend) (*OutboxEntry, error) {
	contract, err := bindOutboxEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// NewOutboxEntryCaller creates a new read-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryCaller(address common.Address, caller bind.ContractCaller) (*OutboxEntryCaller, error) {
	contract, err := bindOutboxEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryCaller{contract: contract}, nil
}

// NewOutboxEntryTransactor creates a new write-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxEntryTransactor, error) {
	contract, err := bindOutboxEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryTransactor{contract: contract}, nil
}

// NewOutboxEntryFilterer creates a new log filterer instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxEntryFilterer, error) {
	contract, err := bindOutboxEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryFilterer{contract: contract}, nil
}

// bindOutboxEntry binds a generic wrapper to an already deployed contract.
func bindOutboxEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.OutboxEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transact(opts, method, params...)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntryCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntrySession) IsMaster() (bool, error) {
	return _OutboxEntry.Contract.IsMaster(&_OutboxEntry.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntryCallerSession) IsMaster() (bool, error) {
	return _OutboxEntry.Contract.IsMaster(&_OutboxEntry.CallOpts)
}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntryCaller) NumRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "numRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntrySession) NumRemaining() (*big.Int, error) {
	return _OutboxEntry.Contract.NumRemaining(&_OutboxEntry.CallOpts)
}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntryCallerSession) NumRemaining() (*big.Int, error) {
	return _OutboxEntry.Contract.NumRemaining(&_OutboxEntry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntryCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntrySession) Root() ([32]byte, error) {
	return _OutboxEntry.Contract.Root(&_OutboxEntry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntryCallerSession) Root() ([32]byte, error) {
	return _OutboxEntry.Contract.Root(&_OutboxEntry.CallOpts)
}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntryCaller) SpentOutput(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "spentOutput", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntrySession) SpentOutput(arg0 [32]byte) (bool, error) {
	return _OutboxEntry.Contract.SpentOutput(&_OutboxEntry.CallOpts, arg0)
}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntryCallerSession) SpentOutput(arg0 [32]byte) (bool, error) {
	return _OutboxEntry.Contract.SpentOutput(&_OutboxEntry.CallOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntryTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntrySession) Destroy() (*types.Transaction, error) {
	return _OutboxEntry.Contract.Destroy(&_OutboxEntry.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntryTransactorSession) Destroy() (*types.Transaction, error) {
	return _OutboxEntry.Contract.Destroy(&_OutboxEntry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b36c66b.
//
// Solidity: function initialize(bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntryTransactor) Initialize(opts *bind.TransactOpts, _root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "initialize", _root, _numInBatch)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b36c66b.
//
// Solidity: function initialize(bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntrySession) Initialize(_root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.Initialize(&_OutboxEntry.TransactOpts, _root, _numInBatch)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b36c66b.
//
// Solidity: function initialize(bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) Initialize(_root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.Initialize(&_OutboxEntry.TransactOpts, _root, _numInBatch)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns(uint256)
func (_OutboxEntry *OutboxEntryTransactor) SpendOutput(opts *bind.TransactOpts, _root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "spendOutput", _root, _id)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns(uint256)
func (_OutboxEntry *OutboxEntrySession) SpendOutput(_root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, _root, _id)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns(uint256)
func (_OutboxEntry *OutboxEntryTransactorSession) SpendOutput(_root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, _root, _id)
}
