// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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

// INodeABI is the input ABI used to generate the binding from.
const INodeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"addStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkConfirmInvalid\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"latestConfirmed\",\"type\":\"uint256\"}],\"name\":\"checkConfirmOutOfOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStakerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestConfirmed\",\"type\":\"uint256\"}],\"name\":\"checkConfirmValid\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"removeStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// INodeFuncSigs maps the 4-byte function signature to its string representation.
var INodeFuncSigs = map[string]string{
	"2466696e": "addStaker(address)",
	"5b8b2280": "challengeHash()",
	"1a8a092b": "checkConfirmInvalid(uint256)",
	"284426b2": "checkConfirmOutOfOrder(uint256)",
	"6cf00e7e": "checkConfirmValid(uint256,uint256)",
	"97bdc510": "confirmData()",
	"2edfb42a": "deadlineBlock()",
	"83197ef0": "destroy()",
	"a406b374": "initialize(address,bytes32,bytes32,bytes32,uint256,uint256)",
	"479c9254": "prev()",
	"96a9fdc0": "removeStaker(address)",
	"dff69787": "stakerCount()",
	"9168ae72": "stakers(address)",
	"701da98e": "stateHash()",
}

// INode is an auto generated Go binding around an Ethereum contract.
type INode struct {
	INodeCaller     // Read-only binding to the contract
	INodeTransactor // Write-only binding to the contract
	INodeFilterer   // Log filterer for contract events
}

// INodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type INodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INodeSession struct {
	Contract     *INode            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INodeCallerSession struct {
	Contract *INodeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// INodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INodeTransactorSession struct {
	Contract     *INodeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type INodeRaw struct {
	Contract *INode // Generic contract binding to access the raw methods on
}

// INodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INodeCallerRaw struct {
	Contract *INodeCaller // Generic read-only contract binding to access the raw methods on
}

// INodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INodeTransactorRaw struct {
	Contract *INodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINode creates a new instance of INode, bound to a specific deployed contract.
func NewINode(address common.Address, backend bind.ContractBackend) (*INode, error) {
	contract, err := bindINode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INode{INodeCaller: INodeCaller{contract: contract}, INodeTransactor: INodeTransactor{contract: contract}, INodeFilterer: INodeFilterer{contract: contract}}, nil
}

// NewINodeCaller creates a new read-only instance of INode, bound to a specific deployed contract.
func NewINodeCaller(address common.Address, caller bind.ContractCaller) (*INodeCaller, error) {
	contract, err := bindINode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INodeCaller{contract: contract}, nil
}

// NewINodeTransactor creates a new write-only instance of INode, bound to a specific deployed contract.
func NewINodeTransactor(address common.Address, transactor bind.ContractTransactor) (*INodeTransactor, error) {
	contract, err := bindINode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INodeTransactor{contract: contract}, nil
}

// NewINodeFilterer creates a new log filterer instance of INode, bound to a specific deployed contract.
func NewINodeFilterer(address common.Address, filterer bind.ContractFilterer) (*INodeFilterer, error) {
	contract, err := bindINode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INodeFilterer{contract: contract}, nil
}

// bindINode binds a generic wrapper to an already deployed contract.
func bindINode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INode.Contract.INodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.contract.Transact(opts, method, params...)
}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_INode *INodeCaller) ChallengeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "challengeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_INode *INodeSession) ChallengeHash() ([32]byte, error) {
	return _INode.Contract.ChallengeHash(&_INode.CallOpts)
}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_INode *INodeCallerSession) ChallengeHash() ([32]byte, error) {
	return _INode.Contract.ChallengeHash(&_INode.CallOpts)
}

// CheckConfirmInvalid is a free data retrieval call binding the contract method 0x1a8a092b.
//
// Solidity: function checkConfirmInvalid(uint256 zombieStakerCount) view returns()
func (_INode *INodeCaller) CheckConfirmInvalid(opts *bind.CallOpts, zombieStakerCount *big.Int) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "checkConfirmInvalid", zombieStakerCount)

	if err != nil {
		return err
	}

	return err

}

// CheckConfirmInvalid is a free data retrieval call binding the contract method 0x1a8a092b.
//
// Solidity: function checkConfirmInvalid(uint256 zombieStakerCount) view returns()
func (_INode *INodeSession) CheckConfirmInvalid(zombieStakerCount *big.Int) error {
	return _INode.Contract.CheckConfirmInvalid(&_INode.CallOpts, zombieStakerCount)
}

// CheckConfirmInvalid is a free data retrieval call binding the contract method 0x1a8a092b.
//
// Solidity: function checkConfirmInvalid(uint256 zombieStakerCount) view returns()
func (_INode *INodeCallerSession) CheckConfirmInvalid(zombieStakerCount *big.Int) error {
	return _INode.Contract.CheckConfirmInvalid(&_INode.CallOpts, zombieStakerCount)
}

// CheckConfirmOutOfOrder is a free data retrieval call binding the contract method 0x284426b2.
//
// Solidity: function checkConfirmOutOfOrder(uint256 latestConfirmed) view returns()
func (_INode *INodeCaller) CheckConfirmOutOfOrder(opts *bind.CallOpts, latestConfirmed *big.Int) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "checkConfirmOutOfOrder", latestConfirmed)

	if err != nil {
		return err
	}

	return err

}

// CheckConfirmOutOfOrder is a free data retrieval call binding the contract method 0x284426b2.
//
// Solidity: function checkConfirmOutOfOrder(uint256 latestConfirmed) view returns()
func (_INode *INodeSession) CheckConfirmOutOfOrder(latestConfirmed *big.Int) error {
	return _INode.Contract.CheckConfirmOutOfOrder(&_INode.CallOpts, latestConfirmed)
}

// CheckConfirmOutOfOrder is a free data retrieval call binding the contract method 0x284426b2.
//
// Solidity: function checkConfirmOutOfOrder(uint256 latestConfirmed) view returns()
func (_INode *INodeCallerSession) CheckConfirmOutOfOrder(latestConfirmed *big.Int) error {
	return _INode.Contract.CheckConfirmOutOfOrder(&_INode.CallOpts, latestConfirmed)
}

// CheckConfirmValid is a free data retrieval call binding the contract method 0x6cf00e7e.
//
// Solidity: function checkConfirmValid(uint256 totalStakerCount, uint256 latestConfirmed) view returns()
func (_INode *INodeCaller) CheckConfirmValid(opts *bind.CallOpts, totalStakerCount *big.Int, latestConfirmed *big.Int) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "checkConfirmValid", totalStakerCount, latestConfirmed)

	if err != nil {
		return err
	}

	return err

}

// CheckConfirmValid is a free data retrieval call binding the contract method 0x6cf00e7e.
//
// Solidity: function checkConfirmValid(uint256 totalStakerCount, uint256 latestConfirmed) view returns()
func (_INode *INodeSession) CheckConfirmValid(totalStakerCount *big.Int, latestConfirmed *big.Int) error {
	return _INode.Contract.CheckConfirmValid(&_INode.CallOpts, totalStakerCount, latestConfirmed)
}

// CheckConfirmValid is a free data retrieval call binding the contract method 0x6cf00e7e.
//
// Solidity: function checkConfirmValid(uint256 totalStakerCount, uint256 latestConfirmed) view returns()
func (_INode *INodeCallerSession) CheckConfirmValid(totalStakerCount *big.Int, latestConfirmed *big.Int) error {
	return _INode.Contract.CheckConfirmValid(&_INode.CallOpts, totalStakerCount, latestConfirmed)
}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_INode *INodeCaller) ConfirmData(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "confirmData")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_INode *INodeSession) ConfirmData() ([32]byte, error) {
	return _INode.Contract.ConfirmData(&_INode.CallOpts)
}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_INode *INodeCallerSession) ConfirmData() ([32]byte, error) {
	return _INode.Contract.ConfirmData(&_INode.CallOpts)
}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_INode *INodeCaller) DeadlineBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "deadlineBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_INode *INodeSession) DeadlineBlock() (*big.Int, error) {
	return _INode.Contract.DeadlineBlock(&_INode.CallOpts)
}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_INode *INodeCallerSession) DeadlineBlock() (*big.Int, error) {
	return _INode.Contract.DeadlineBlock(&_INode.CallOpts)
}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_INode *INodeCaller) Prev(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "prev")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_INode *INodeSession) Prev() (*big.Int, error) {
	return _INode.Contract.Prev(&_INode.CallOpts)
}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_INode *INodeCallerSession) Prev() (*big.Int, error) {
	return _INode.Contract.Prev(&_INode.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_INode *INodeCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_INode *INodeSession) StakerCount() (*big.Int, error) {
	return _INode.Contract.StakerCount(&_INode.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_INode *INodeCallerSession) StakerCount() (*big.Int, error) {
	return _INode.Contract.StakerCount(&_INode.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(bool)
func (_INode *INodeCaller) Stakers(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "stakers", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(bool)
func (_INode *INodeSession) Stakers(staker common.Address) (bool, error) {
	return _INode.Contract.Stakers(&_INode.CallOpts, staker)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(bool)
func (_INode *INodeCallerSession) Stakers(staker common.Address) (bool, error) {
	return _INode.Contract.Stakers(&_INode.CallOpts, staker)
}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_INode *INodeCaller) StateHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "stateHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_INode *INodeSession) StateHash() ([32]byte, error) {
	return _INode.Contract.StateHash(&_INode.CallOpts)
}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_INode *INodeCallerSession) StateHash() ([32]byte, error) {
	return _INode.Contract.StateHash(&_INode.CallOpts)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns()
func (_INode *INodeTransactor) AddStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "addStaker", staker)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns()
func (_INode *INodeSession) AddStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.AddStaker(&_INode.TransactOpts, staker)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns()
func (_INode *INodeTransactorSession) AddStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.AddStaker(&_INode.TransactOpts, staker)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_INode *INodeTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_INode *INodeSession) Destroy() (*types.Transaction, error) {
	return _INode.Contract.Destroy(&_INode.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_INode *INodeTransactorSession) Destroy() (*types.Transaction, error) {
	return _INode.Contract.Destroy(&_INode.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_INode *INodeTransactor) Initialize(opts *bind.TransactOpts, _rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "initialize", _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_INode *INodeSession) Initialize(_rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Initialize(&_INode.TransactOpts, _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_INode *INodeTransactorSession) Initialize(_rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Initialize(&_INode.TransactOpts, _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_INode *INodeTransactor) RemoveStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "removeStaker", staker)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_INode *INodeSession) RemoveStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.RemoveStaker(&_INode.TransactOpts, staker)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_INode *INodeTransactorSession) RemoveStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.RemoveStaker(&_INode.TransactOpts, staker)
}

// InboxABI is the input ABI used to generate the binding from.
const InboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"BuddyContractPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxMaxAcc\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxMaxCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InboxFuncSigs maps the 4-byte function signature to its string representation.
var InboxFuncSigs = map[string]string{
	"6f5dfdca": "deployL2ContractPair(uint256,uint256,uint256,bytes)",
	"afcc220b": "depositEthMessage(address)",
	"f0a79973": "inboxMaxAcc()",
	"917cae02": "inboxMaxCount()",
	"b75436bb": "sendL2Message(bytes)",
	"1fe927cf": "sendL2MessageFromOrigin(bytes)",
}

// InboxBin is the compiled bytecode used for deploying new contracts.
var InboxBin = "0x608060405234801561001057600080fd5b50610682806100206000396000f3fe6080604052600436106100555760003560e01c80631fe927cf1461005a5780636f5dfdca146100d9578063917cae0214610169578063afcc220b14610190578063b75436bb146101b6578063f0a7997314610233575b600080fd5b34801561006657600080fd5b506100d76004803603602081101561007d57600080fd5b81019060208101813564010000000081111561009857600080fd5b8201836020820111156100aa57600080fd5b803590602001918460018302840111640100000000831117156100cc57600080fd5b509092509050610248565b005b3480156100e557600080fd5b506100d7600480360360808110156100fc57600080fd5b8135916020810135916040820135919081019060808101606082013564010000000081111561012a57600080fd5b82018360208201111561013c57600080fd5b8035906020019184600183028401116401000000008311171561015e57600080fd5b5090925090506102fe565b34801561017557600080fd5b5061017e6103d6565b60408051918252519081900360200190f35b6100d7600480360360208110156101a657600080fd5b50356001600160a01b03166103dc565b3480156101c257600080fd5b506100d7600480360360208110156101d957600080fd5b8101906020810181356401000000008111156101f457600080fd5b82018360208201111561020657600080fd5b8035906020019184600183028401116401000000008311171561022857600080fd5b509092509050610416565b34801561023f57600080fd5b5061017e61045c565b33321461028a576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b6000806102b56003338686604051808383808284376040519201829003909120935061046292505050565b60408051600381523360208201528151939550919350839285927f852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab92908290030190a350505050565b6103073361049f565b610358576040805162461bcd60e51b815260206004820152601a60248201527f6d7573742062652063616c6c656420627920636f6e7472616374000000000000604482015290519081900360640190fd5b6103a460053387878787876040516020018086815260200185815260200184815260200183838082843780830192505050955050505050506040516020818303038152906040526104db565b60405133907f49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f90600090a25050505050565b60015481565b604080516001600160a01b0383166020820152348183015281518082038301815260609091019091526104139060009033906104db565b50565b61045860033384848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104db92505050565b5050565b60005481565b60015460008054909182918261047c88884342878b6105b2565b90506104888282610620565b600055506001828101905590969095509350505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906104d357508115155b949350505050565b6000806104f085858580519060200120610462565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561056f578181015183820152602001610557565b50505050905090810190601f16801561059c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6bffffffffffffffffffffffff191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b60408051602080820194909452808201929092528051808303820181526060909201905280519101209056fea2646970667358221220d2016bdc427974123ee55d414df76f344d9fbf7a8c7575789681ecbef9417bde64736f6c634300060c0033"

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// InboxMaxAcc is a free data retrieval call binding the contract method 0xf0a79973.
//
// Solidity: function inboxMaxAcc() view returns(bytes32)
func (_Inbox *InboxCaller) InboxMaxAcc(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "inboxMaxAcc")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxMaxAcc is a free data retrieval call binding the contract method 0xf0a79973.
//
// Solidity: function inboxMaxAcc() view returns(bytes32)
func (_Inbox *InboxSession) InboxMaxAcc() ([32]byte, error) {
	return _Inbox.Contract.InboxMaxAcc(&_Inbox.CallOpts)
}

// InboxMaxAcc is a free data retrieval call binding the contract method 0xf0a79973.
//
// Solidity: function inboxMaxAcc() view returns(bytes32)
func (_Inbox *InboxCallerSession) InboxMaxAcc() ([32]byte, error) {
	return _Inbox.Contract.InboxMaxAcc(&_Inbox.CallOpts)
}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Inbox *InboxCaller) InboxMaxCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "inboxMaxCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Inbox *InboxSession) InboxMaxCount() (*big.Int, error) {
	return _Inbox.Contract.InboxMaxCount(&_Inbox.CallOpts)
}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Inbox *InboxCallerSession) InboxMaxCount() (*big.Int, error) {
	return _Inbox.Contract.InboxMaxCount(&_Inbox.CallOpts)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxTransactor) DeployL2ContractPair(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "deployL2ContractPair", maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.DeployL2ContractPair(&_Inbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxTransactorSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.DeployL2ContractPair(&_Inbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxTransactor) DepositEthMessage(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEthMessage", to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEthMessage(&_Inbox.TransactOpts, to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxTransactorSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEthMessage(&_Inbox.TransactOpts, to)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// InboxBuddyContractPairIterator is returned from FilterBuddyContractPair and is used to iterate over the raw logs and unpacked data for BuddyContractPair events raised by the Inbox contract.
type InboxBuddyContractPairIterator struct {
	Event *InboxBuddyContractPair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InboxBuddyContractPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxBuddyContractPair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InboxBuddyContractPair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InboxBuddyContractPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxBuddyContractPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxBuddyContractPair represents a BuddyContractPair event raised by the Inbox contract.
type InboxBuddyContractPair struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuddyContractPair is a free log retrieval operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) FilterBuddyContractPair(opts *bind.FilterOpts, sender []common.Address) (*InboxBuddyContractPairIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxBuddyContractPairIterator{contract: _Inbox.contract, event: "BuddyContractPair", logs: logs, sub: sub}, nil
}

// WatchBuddyContractPair is a free log subscription operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) WatchBuddyContractPair(opts *bind.WatchOpts, sink chan<- *InboxBuddyContractPair, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxBuddyContractPair)
				if err := _Inbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuddyContractPair is a log parse operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) ParseBuddyContractPair(log types.Log) (*InboxBuddyContractPair, error) {
	event := new(InboxBuddyContractPair)
	if err := _Inbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the Inbox contract.
type InboxMessageDeliveredIterator struct {
	Event *InboxMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InboxMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxMessageDelivered represents a MessageDelivered event raised by the Inbox contract.
type InboxMessageDelivered struct {
	MessageNum     *big.Int
	BeforeInboxAcc [32]byte
	Kind           uint8
	Sender         common.Address
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int, beforeInboxAcc [][32]byte) (*InboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "MessageDelivered", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &InboxMessageDeliveredIterator{contract: _Inbox.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxMessageDelivered, messageNum []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "MessageDelivered", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxMessageDelivered)
				if err := _Inbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) ParseMessageDelivered(log types.Log) (*InboxMessageDelivered, error) {
	event := new(InboxMessageDelivered)
	if err := _Inbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxMessageDeliveredFromOriginIterator is returned from FilterMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for MessageDeliveredFromOrigin events raised by the Inbox contract.
type InboxMessageDeliveredFromOriginIterator struct {
	Event *InboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxMessageDeliveredFromOrigin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InboxMessageDeliveredFromOrigin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxMessageDeliveredFromOrigin represents a MessageDeliveredFromOrigin event raised by the Inbox contract.
type InboxMessageDeliveredFromOrigin struct {
	MessageNum     *big.Int
	BeforeInboxAcc [32]byte
	Kind           uint8
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Inbox *InboxFilterer) FilterMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int, beforeInboxAcc [][32]byte) (*InboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "MessageDeliveredFromOrigin", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &InboxMessageDeliveredFromOriginIterator{contract: _Inbox.contract, event: "MessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Inbox *InboxFilterer) WatchMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxMessageDeliveredFromOrigin, messageNum []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "MessageDeliveredFromOrigin", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxMessageDeliveredFromOrigin)
				if err := _Inbox.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDeliveredFromOrigin is a log parse operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Inbox *InboxFilterer) ParseMessageDeliveredFromOrigin(log types.Log) (*InboxMessageDeliveredFromOrigin, error) {
	event := new(InboxMessageDeliveredFromOrigin)
	if err := _Inbox.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e8f4ee94ad90a2d67cd585167b36a7f428a7bc59a03e09c07b327d449e7bdee364736f6c634300060c0033"

// DeployMessages deploys a new Ethereum contract, binding an instance of Messages to it.
func DeployMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Messages, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// Messages is an auto generated Go binding around an Ethereum contract.
type Messages struct {
	MessagesCaller     // Read-only binding to the contract
	MessagesTransactor // Write-only binding to the contract
	MessagesFilterer   // Log filterer for contract events
}

// MessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagesSession struct {
	Contract     *Messages         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagesCallerSession struct {
	Contract *MessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagesTransactorSession struct {
	Contract     *MessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagesRaw struct {
	Contract *Messages // Generic contract binding to access the raw methods on
}

// MessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagesCallerRaw struct {
	Contract *MessagesCaller // Generic read-only contract binding to access the raw methods on
}

// MessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagesTransactorRaw struct {
	Contract *MessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessages creates a new instance of Messages, bound to a specific deployed contract.
func NewMessages(address common.Address, backend bind.ContractBackend) (*Messages, error) {
	contract, err := bindMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// NewMessagesCaller creates a new read-only instance of Messages, bound to a specific deployed contract.
func NewMessagesCaller(address common.Address, caller bind.ContractCaller) (*MessagesCaller, error) {
	contract, err := bindMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesCaller{contract: contract}, nil
}

// NewMessagesTransactor creates a new write-only instance of Messages, bound to a specific deployed contract.
func NewMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagesTransactor, error) {
	contract, err := bindMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesTransactor{contract: contract}, nil
}

// NewMessagesFilterer creates a new log filterer instance of Messages, bound to a specific deployed contract.
func NewMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagesFilterer, error) {
	contract, err := bindMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagesFilterer{contract: contract}, nil
}

// bindMessages binds a generic wrapper to an already deployed contract.
func bindMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.MessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transact(opts, method, params...)
}

// OutboxABI is the input ABI used to generate the binding from.
const OutboxABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OutboxFuncSigs maps the 4-byte function signature to its string representation.
var OutboxFuncSigs = map[string]string{
	"c4fb000c": "executeTransaction(uint256,bytes,uint256,address,uint256,bytes)",
}

// OutboxBin is the compiled bytecode used for deploying new contracts.
var OutboxBin = "0x608060405234801561001057600080fd5b5061040d806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c4fb000c14610030575b600080fd5b610114600480360360c081101561004657600080fd5b8135919081019060408101602082013564010000000081111561006857600080fd5b82018360208201111561007a57600080fd5b8035906020019184600183028401116401000000008311171561009c57600080fd5b919390928235926001600160a01b036020820135169260408201359290916080810190606001356401000000008111156100d557600080fd5b8201836020820111156100e757600080fd5b8035906020019184600183028401116401000000008311171561010957600080fd5b509092509050610116565b005b60008460601b60601c6001600160a01b031684848460405160200180858152602001848152602001838380828437808301925050509450505050506040516020818303038152906040528051906020012090506101ad8989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b925086915061022c9050565b6000856001600160a01b0316858585604051808383808284376040519201945060009350909150508083038185875af1925050503d806000811461020d576040519150601f19603f3d011682016040523d82523d6000602084013e610212565b606091505b505090508061022057600080fd5b50505050505050505050565b600160001b8118905060006102458483856001016102ca565b5090506000858154811061025557fe5b600091825260208220015460408051630ad0379b60e01b8152600481018590526024810187905290516001600160a01b0390921692630ad0379b9260448084019382900301818387803b1580156102ab57600080fd5b505af11580156102bf573d6000803e3d6000fd5b505050505050505050565b60008080848160205b885181116103c9578089015193506020818a5103602001816102f157fe5b0491505b6000821180156103085750600287066001145b801561031657508160020a87115b1561032e5760029096046001908101969401936102f5565b6002870661037957838360405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002878161037157fe5b0496506103bb565b82846040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600287816103b457fe5b0460010196505b6001909401936020016102d3565b50909350505093509391505056fea2646970667358221220c7ee93972a06fe138a786dbdfb015ffe51e340f00c81c31f4a16b2669e267b8f64736f6c634300060c0033"

// DeployOutbox deploys a new Ethereum contract, binding an instance of Outbox to it.
func DeployOutbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Outbox, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// Outbox is an auto generated Go binding around an Ethereum contract.
type Outbox struct {
	OutboxCaller     // Read-only binding to the contract
	OutboxTransactor // Write-only binding to the contract
	OutboxFilterer   // Log filterer for contract events
}

// OutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxSession struct {
	Contract     *Outbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxCallerSession struct {
	Contract *OutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxTransactorSession struct {
	Contract     *OutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxRaw struct {
	Contract *Outbox // Generic contract binding to access the raw methods on
}

// OutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxCallerRaw struct {
	Contract *OutboxCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxTransactorRaw struct {
	Contract *OutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutbox creates a new instance of Outbox, bound to a specific deployed contract.
func NewOutbox(address common.Address, backend bind.ContractBackend) (*Outbox, error) {
	contract, err := bindOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// NewOutboxCaller creates a new read-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxCaller(address common.Address, caller bind.ContractCaller) (*OutboxCaller, error) {
	contract, err := bindOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxCaller{contract: contract}, nil
}

// NewOutboxTransactor creates a new write-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxTransactor, error) {
	contract, err := bindOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxTransactor{contract: contract}, nil
}

// NewOutboxFilterer creates a new log filterer instance of Outbox, bound to a specific deployed contract.
func NewOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxFilterer, error) {
	contract, err := bindOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxFilterer{contract: contract}, nil
}

// bindOutbox binds a generic wrapper to an already deployed contract.
func bindOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.OutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transact(opts, method, params...)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactor) ExecuteTransaction(opts *bind.TransactOpts, outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeTransaction", outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxSession) ExecuteTransaction(outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactorSession) ExecuteTransaction(outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// OutboxEntryABI is the input ABI used to generate the binding from.
const OutboxEntryABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"calcRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"spendOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OutboxEntryFuncSigs maps the 4-byte function signature to its string representation.
var OutboxEntryFuncSigs = map[string]string{
	"0ad0379b": "spendOutput(bytes32,uint256)",
}

// OutboxEntryBin is the compiled bytecode used for deploying new contracts.
var OutboxEntryBin = "0x608060405234801561001057600080fd5b506040516101373803806101378339818101604052602081101561003357600080fd5b5051600080546001600160a01b0319163317905560015560df806100586000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630ad0379b14602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135604f565b005b6000546001600160a01b03163314606557600080fd5b60008181526002602052604090205460ff1615608057600080fd5b6001548214608d57600080fd5b6000908152600260205260409020805460ff191660011790555056fea264697066735822122045771b255caea9c72008e9382dc6c60996b94002366693a4d85d84c26f81042064736f6c634300060c0033"

// DeployOutboxEntry deploys a new Ethereum contract, binding an instance of OutboxEntry to it.
func DeployOutboxEntry(auth *bind.TransactOpts, backend bind.ContractBackend, root [32]byte) (common.Address, *types.Transaction, *OutboxEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxEntryBin), backend, root)
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

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 index) returns()
func (_OutboxEntry *OutboxEntryTransactor) SpendOutput(opts *bind.TransactOpts, calcRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "spendOutput", calcRoot, index)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 index) returns()
func (_OutboxEntry *OutboxEntrySession) SpendOutput(calcRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, calcRoot, index)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 index) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) SpendOutput(calcRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, calcRoot, index)
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"BuddyContractPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"indexed\":false,\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"inboxMaxHash\",\"type\":\"bytes32\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"SentLogs\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"addToDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkNoRecentStake\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"logAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[3]\",\"name\":\"nodeFields\",\"type\":\"bytes32[3]\"},{\"internalType\":\"uint256\",\"name\":\"executionCheckTime\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxMaxAcc\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxMaxCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxReduction\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"successorWithStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"stakerInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerList\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"zombieInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupFuncSigs maps the 4-byte function signature to its string representation.
var RollupFuncSigs = map[string]string{
	"45c5b2c7": "addToDeposit(address)",
	"5e8ef106": "arbGasSpeedLimitPerBlock()",
	"76e7e23b": "baseStake()",
	"5dbaf68b": "challengeFactory()",
	"46c2781a": "challengePeriodBlocks()",
	"be211c9a": "checkNoRecentStake()",
	"73f33b06": "checkUnresolved()",
	"fa7803e6": "completeChallenge(address,address)",
	"396b8cbc": "confirmNextNode(bytes32,bytes,uint256[])",
	"04a28064": "countStakedZombies(address)",
	"b1b4181c": "createChallenge(address[2],uint256[2],bytes32[3],uint256)",
	"4d26732d": "currentRequiredStake()",
	"6f5dfdca": "deployL2ContractPair(uint256,uint256,uint256,bytes)",
	"afcc220b": "depositEthMessage(address)",
	"c4fb000c": "executeTransaction(uint256,bytes,uint256,address,uint256,bytes)",
	"d735e21d": "firstUnresolvedNode()",
	"ad71bd36": "getStakers(uint256,uint256)",
	"f0a79973": "inboxMaxAcc()",
	"917cae02": "inboxMaxCount()",
	"8640ce5f": "lastStakeBlock()",
	"65f7f80d": "latestConfirmed()",
	"7ba9534a": "latestNodeCreated()",
	"45e38b64": "minimumAssertionPeriod()",
	"5f576db6": "newStake()",
	"d93fe9c4": "nodeFactory()",
	"1c53c280": "nodes(uint256)",
	"1e83d30f": "reduceDeposit(uint256)",
	"0e1ef04c": "rejectNextNode(uint256,address)",
	"edfd03ed": "removeOldZombies(uint256)",
	"7e2d2155": "removeZombie(uint256,uint256)",
	"7427be51": "returnOldDeposit(address)",
	"b75436bb": "sendL2Message(bytes)",
	"1fe927cf": "sendL2MessageFromOrigin(bytes)",
	"8fd18f04": "stakeOnExistingNode(bytes32,uint256,uint256)",
	"f019a1c1": "stakeOnNewNode(bytes32,uint256,uint256,bytes32[7],uint256[10])",
	"51ed6a30": "stakeToken()",
	"dff69787": "stakerCount()",
	"4e745f1f": "stakerInfo(address)",
	"348e50c6": "stakerList(uint256)",
	"729cfe3b": "stakerMap(address)",
	"63721d6b": "zombieCount()",
	"4a95e20e": "zombieInfo(uint256)",
}

// RollupBin is the compiled bytecode used for deploying new contracts.
var RollupBin = "0x60806040523480156200001157600080fd5b506040516200434b3803806200434b83398181016040526101208110156200003857600080fd5b815160208301516040808501516060860151608087015160a088015160c089015160e08a01516101008b0180519751999b989a969995989497939692959194919392820192846401000000008211156200009157600080fd5b908301906020820185811115620000a757600080fd5b8251640100000000811182820188101715620000c257600080fd5b82525081516020918201929091019080838360005b83811015620000f1578181015183820152602001620000d7565b50505050905090810190601f1680156200011f5780820380516001836020036101000a031916815260200191505b5060405250505082600f60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081601060006101000a8154816001600160a01b0302191690836001600160a01b031602179055506200022f8888888860601b6001600160601b0319168860601b6001600160601b031916866040516020018087815260200186815260200185815260200184815260200183815260200182805190602001908083835b60208310620001ea5780518252601f199092019160209182019101620001c9565b6001836020036101000a0380198251168184511680821785525050505050509050019650505050505050604051602081830303815290604052620003ab60201b60201c565b6000620002564360008c6000801b6000806000600154620003bc60201b62002fa81760201c565b6010546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905260848201819052915193945090926001600160a01b039092169163d45ab2b59160a48082019260209290919082900301818787803b158015620002c757600080fd5b505af1158015620002dc573d6000803e3d6000fd5b505050506040513d6020811015620002f357600080fd5b505160008052600660209081527f54cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f880546001600160a01b038085166001600160a01b031992831617909255600b8e9055600c8d9055600d8c9055600e8054928c16929091169190911790556001600455604080518e815290519293507f4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d929081900390910190a1505050505050505050505062000602565b620003b96004308362000416565b50565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b6000806200043385858580519060200120620004f860201b60201c565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015620004b45781810151838201526020016200049a565b50505050905090810190601f168015620004e25780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b6000806000600154905060008054905060006200052588884342878b6200055560201b620030021760201c565b90506200053e8282620005d660201b620030701760201c565b600055506001828101905590969095509350505050565b6040805160f89790971b7fff000000000000000000000000000000000000000000000000000000000000001660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b613d3980620006126000396000f3fe6080604052600436106102515760003560e01c806373f33b0611610139578063b1b4181c116100b6578063d93fe9c41161007a578063d93fe9c414610aac578063dff6978714610ac1578063edfd03ed14610ad6578063f019a1c114610b00578063f0a7997314610b42578063fa7803e614610b5757610251565b8063b1b4181c146108e3578063b75436bb1461091a578063be211c9a14610995578063c4fb000c146109aa578063d735e21d14610a9757610251565b80638640ce5f116100fd5780638640ce5f146107dd5780638fd18f04146107f2578063917cae0214610828578063ad71bd361461083d578063afcc220b146108bd57610251565b806373f33b061461073b5780637427be511461075057806376e7e23b146107835780637ba9534a146107985780637e2d2155146107ad57610251565b80634a95e20e116101d25780635e8ef106116101965780635e8ef106146105fd5780635f576db61461061257806363721d6b1461061a57806365f7f80d1461062f5780636f5dfdca14610644578063729cfe3b146106d257610251565b80634a95e20e1461050d5780634d26732d1461055a5780634e745f1f1461056f57806351ed6a30146105d35780635dbaf68b146105e857610251565b8063348e50c611610219578063348e50c6146103c1578063396b8cbc146103eb57806345c5b2c7146104bd57806345e38b64146104e357806346c2781a146104f857610251565b806304a28064146102565780630e1ef04c1461029b5780631c53c280146102d65780631e83d30f1461031c5780631fe927cf14610346575b600080fd5b34801561026257600080fd5b506102896004803603602081101561027957600080fd5b50356001600160a01b0316610b92565b60408051918252519081900360200190f35b3480156102a757600080fd5b506102d4600480360360408110156102be57600080fd5b50803590602001356001600160a01b0316610c55565b005b3480156102e257600080fd5b50610300600480360360208110156102f957600080fd5b50356110f1565b604080516001600160a01b039092168252519081900360200190f35b34801561032857600080fd5b506102d46004803603602081101561033f57600080fd5b503561110c565b34801561035257600080fd5b506102d46004803603602081101561036957600080fd5b810190602081018135600160201b81111561038357600080fd5b82018360208201111561039557600080fd5b803590602001918460018302840111600160201b831117156103b657600080fd5b509092509050611187565b3480156103cd57600080fd5b50610300600480360360208110156103e457600080fd5b503561123d565b3480156103f757600080fd5b506102d46004803603606081101561040e57600080fd5b81359190810190604081016020820135600160201b81111561042f57600080fd5b82018360208201111561044157600080fd5b803590602001918460018302840111600160201b8311171561046257600080fd5b919390929091602081019035600160201b81111561047f57600080fd5b82018360208201111561049157600080fd5b803590602001918460208302840111600160201b831117156104b257600080fd5b509092509050611264565b6102d4600480360360208110156104d357600080fd5b50356001600160a01b03166114f7565b3480156104ef57600080fd5b50610289611524565b34801561050457600080fd5b50610289611539565b34801561051957600080fd5b506105376004803603602081101561053057600080fd5b503561153f565b604080516001600160a01b03909316835260208301919091528051918290030190f35b34801561056657600080fd5b5061028961157e565b34801561057b57600080fd5b506105a26004803603602081101561059257600080fd5b50356001600160a01b0316611679565b6040805194151585526020850193909352838301919091526001600160a01b03166060830152519081900360800190f35b3480156105df57600080fd5b506103006116b6565b3480156105f457600080fd5b506103006116c5565b34801561060957600080fd5b506102896116d4565b6102d46116da565b34801561062657600080fd5b50610289611855565b34801561063b57600080fd5b5061028961185b565b34801561065057600080fd5b506102d46004803603608081101561066757600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561069457600080fd5b8201836020820111156106a657600080fd5b803590602001918460018302840111600160201b831117156106c757600080fd5b509092509050611861565b3480156106de57600080fd5b50610705600480360360208110156106f557600080fd5b50356001600160a01b0316611939565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b34801561074757600080fd5b506102d4611975565b34801561075c57600080fd5b506102d46004803603602081101561077357600080fd5b50356001600160a01b03166119cf565b34801561078f57600080fd5b50610289611a82565b3480156107a457600080fd5b50610289611a88565b3480156107b957600080fd5b506102d4600480360360408110156107d057600080fd5b5080359060200135611a8e565b3480156107e957600080fd5b50610289611ca9565b3480156107fe57600080fd5b506102d46004803603606081101561081557600080fd5b5080359060208101359060400135611caf565b34801561083457600080fd5b50610289611e9f565b34801561084957600080fd5b5061086d6004803603604081101561086057600080fd5b5080359060200135611ea5565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156108a9578181015183820152602001610891565b505050509050019250505060405180910390f35b6102d4600480360360208110156108d357600080fd5b50356001600160a01b0316611f6d565b3480156108ef57600080fd5b506102d4600480360361010081101561090757600080fd5b50604081016080820160e0830135611fa7565b34801561092657600080fd5b506102d46004803603602081101561093d57600080fd5b810190602081018135600160201b81111561095757600080fd5b82018360208201111561096957600080fd5b803590602001918460018302840111600160201b8311171561098a57600080fd5b509092509050612563565b3480156109a157600080fd5b506102d46125a9565b3480156109b657600080fd5b506102d4600480360360c08110156109cd57600080fd5b81359190810190604081016020820135600160201b8111156109ee57600080fd5b820183602082011115610a0057600080fd5b803590602001918460018302840111600160201b83111715610a2157600080fd5b919390928235926001600160a01b03602082013516926040820135929091608081019060600135600160201b811115610a5957600080fd5b820183602082011115610a6b57600080fd5b803590602001918460018302840111600160201b83111715610a8c57600080fd5b5090925090506125f3565b348015610aa357600080fd5b50610289612709565b348015610ab857600080fd5b5061030061270f565b348015610acd57600080fd5b5061028961271e565b348015610ae257600080fd5b506102d460048036036020811015610af957600080fd5b5035612724565b348015610b0c57600080fd5b506102d46004803603610280811015610b2457600080fd5b5080359060208101359060408101359060608101906101400161283f565b348015610b4e57600080fd5b50610289612e30565b348015610b6357600080fd5b506102d460048036036040811015610b7a57600080fd5b506001600160a01b0381358116916020013516612e36565b600a5460009081805b82811015610c4d576000600a8281548110610bb257fe5b60009182526020918290206002909102018054604080516348b4573960e11b81526001600160a01b039283166004820152905192945090891692639168ae7292602480840193829003018186803b158015610c0c57600080fd5b505afa158015610c20573d6000803e3d6000fd5b505050506040513d6020811015610c3657600080fd5b505115610c44576001909201915b50600101610b9b565b509392505050565b610c5d611975565b600060066000600454815260200190815260200160002060009054906101000a90046001600160a01b03169050600354816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610cc657600080fd5b505afa158015610cda573d6000803e3d6000fd5b505050506040513d6020811015610cf057600080fd5b505114156110d857610d006125a9565b6004548311610d49576040805162461bcd60e51b815260206004820152601060248201526f535543434553534f525f544f5f4c4f5760801b604482015290519081900360640190fd5b600554831115610d94576040805162461bcd60e51b81526020600482015260116024820152700a6aa86868aa6a69ea4bea89ebe90928e9607b1b604482015290519081900360640190fd5b6001600160a01b038216600090815260096020526040902060030154600160a01b900460ff16610df8576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6000838152600660209081526040918290205460035483516311e7249560e21b815293516001600160a01b03909216939092849263479c9254926004808201939291829003018186803b158015610e4e57600080fd5b505afa158015610e62573d6000803e3d6000fd5b505050506040513d6020811015610e7857600080fd5b505114610ebc576040805162461bcd60e51b815260206004820152600d60248201526c2120a22fa9aaa1a1a2a9a9a7a960991b604482015290519081900360640190fd5b806001600160a01b0316639168ae72846040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610f0957600080fd5b505afa158015610f1d573d6000803e3d6000fd5b505050506040513d6020811015610f3357600080fd5b5051610f73576040805162461bcd60e51b815260206004820152600a6024820152692120a22fa9aa20a5a2a960b11b604482015290519081900360640190fd5b610f7d6000612724565b816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610fb657600080fd5b505afa158015610fca573d6000803e3d6000fd5b505050506040513d6020811015610fe057600080fd5b5051431015611028576040805162461bcd60e51b815260206004820152600f60248201526e4245464f52455f444541444c494e4560881b604482015290519081900360640190fd5b61103182610b92565b826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561106a57600080fd5b505afa15801561107e573d6000803e3d6000fd5b505050506040513d602081101561109457600080fd5b5051146110d6576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b505b6110e360045461309c565b505060048054600101905550565b6006602052600090815260409020546001600160a01b031681565b3360009081526009602052604090206111248161311e565b600061112e61157e565b90508082600201541161114057600080fd5b6002820154819003838111156111535750825b604051339082156108fc029083906000818181858888f19350505050158015611180573d6000803e3d6000fd5b5050505050565b3332146111c9576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b6000806111f4600333868660405180838380828437604051920182900390912093506131b592505050565b60408051600381523360208201528151939550919350839285927f852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab92908290030190a350505050565b6008818154811061124a57fe5b6000918252602090912001546001600160a01b0316905081565b61126c611975565b6112746125a9565b6004546000908152600660205260408120546001600160a01b03169061129990612724565b806001600160a01b0316636cf00e7e6112b183610b92565b600880549050016003546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156112f657600080fd5b505afa15801561130a573d6000803e3d6000fd5b50505050600061138086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808a028281018201909352898252909350899250889182918501908490808284376000920191909152506131f292505050565b905061138c8188613070565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b1580156113c557600080fd5b505afa1580156113d9573d6000803e3d6000fd5b505050506040513d60208110156113ef57600080fd5b505114611432576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b6114a286868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808a028281018201909352898252909350899250889182918501908490808284376000920191909152506132f292505050565b6114ad60035461309c565b60048054600381905560010190556040805188815290517fe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e849181900360200190a150505050505050565b6001600160a01b03811660009081526009602052604090206115188161311e565b60020180543401905550565b6000600a600b548161153257fe5b0490505b90565b600b5481565b6000806000600a848154811061155157fe5b6000918252602090912060029091020180546001909101546001600160a01b039091169350915050915091565b600354600090815260066020908152604080832054815163176fda1560e11b815291516000199385936001600160a01b0390931692632edfb42a9260048083019392829003018186803b1580156115d457600080fd5b505afa1580156115e8573d6000803e3d6000fd5b505050506040513d60208110156115fe57600080fd5b505190504381111561161657600d5492505050611536565b600081430390506000600b54828161162a57fe5b04905060ff81111561163a575060ff5b600019600282900a018061164c575060015b600d54858161165757fe5b0481111561166c578495505050505050611536565b600d540294505050505090565b6001600160a01b03908116600090815260096020526040902060038101546001820154600290920154600160a01b820460ff169492939092911690565b600e546001600160a01b031681565b600f546001600160a01b031681565b600c5481565b33600090815260096020526040902060030154600160a01b900460ff161561173a576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b61174261157e565b341015611789576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b6008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee381018054336001600160a01b031991821681179092556040805160a0810182529384526003805460208087019182523487850190815260006060890181815260808a018b8152988252600990935294909420965187559051968601969096559051600285015593519290930180549151919093166001600160a01b039092169190911760ff60a01b1916600160a01b9115159190910217905543600755565b600a5490565b60035481565b61186a336133bc565b6118bb576040805162461bcd60e51b815260206004820152601a60248201527f6d7573742062652063616c6c656420627920636f6e7472616374000000000000604482015290519081900360640190fd5b61190760053387878787876040516020018086815260200185815260200184815260200183838082843780830192505050955050505050506040516020818303038152906040526133f8565b60405133907f49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f90600090a25050505050565b6009602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b60035460045411801561198c575060055460045411155b6119cd576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b565b6001600160a01b038116600090815260096020526040902060035460018201541115611a2f576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b611a388161311e565b6002810154611a46826134cf565b6040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015611a7c573d6000803e3d6000fd5b50505050565b600d5481565b60055481565b600a54821115611ad6576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b6000600a8381548110611ae557fe5b9060005260206000209060020201905060008160010154905060005b60045482118015611b1157508381105b15611bf557600082815260066020526040808220548554825163025aa7f760e61b81526001600160a01b039182166004820152925191169283926396a9fdc0926024808301939282900301818387803b158015611b6d57600080fd5b505af1158015611b81573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611bbe57600080fd5b505afa158015611bd2573d6000803e3d6000fd5b505050506040513d6020811015611be857600080fd5b5051925050600101611b01565b600454821015611c9d57600a80546000198101908110611c1157fe5b9060005260206000209060020201600a8681548110611c2c57fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600a805480611c6f57fe5b60008281526020812060026000199093019283020180546001600160a01b0319168155600101559055611180565b50600191909101555050565b60075481565b3360009081526009602052604090206003810154600160a01b900460ff16611d0b576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b83834014611d56576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b6004548210158015611d6a57506005548211155b611d7357600080fd5b6000828152600660209081526040918290205482516311e7249560e21b815292516001600160a01b0390911692839263479c925492600480840193829003018186803b158015611dc257600080fd5b505afa158015611dd6573d6000803e3d6000fd5b505050506040513d6020811015611dec57600080fd5b5051600183015414611e37576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b6040805163123334b760e11b815233600482015290516001600160a01b03831691632466696e91602480830192600092919082900301818387803b158015611e7e57600080fd5b505af1158015611e92573d6000803e3d6000fd5b5050505050600101555050565b60015481565b600854606090838301811115611eba57508282015b60608167ffffffffffffffff81118015611ed357600080fd5b50604051908082528060200260200182016040528015611efd578160200160208202803683370190505b50905060005b82811015611f6457600881870181548110611f1a57fe5b9060005260206000200160009054906101000a90046001600160a01b0316828281518110611f4457fe5b6001600160a01b0390921660209283029190910190910152600101611f03565b50949350505050565b604080516001600160a01b038316602082015234818301528151808203830181526060909101909152611fa49060009033906133f8565b50565b6020830135833510611fee576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b60055460208401351115612038576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b600354833511612083576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b8235600090815260066020908152604080832054828701358452928190205481516311e7249560e21b815291516001600160a01b039485169490911692839263479c92549260048083019392829003018186803b1580156120e357600080fd5b505afa1580156120f7573d6000803e3d6000fd5b505050506040513d602081101561210d57600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b15801561214f57600080fd5b505afa158015612163573d6000803e3d6000fd5b505050506040513d602081101561217957600080fd5b5051146121b9576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b6001600160a01b03863581166000908152600960209081526040808320918a013590931682529190206121eb8261311e565b6121f48161311e565b604080516348b4573960e11b81526001600160a01b038a3581166004830152915191861691639168ae7291602480820192602092909190829003018186803b15801561223f57600080fd5b505afa158015612253573d6000803e3d6000fd5b505050506040513d602081101561226957600080fd5b50516122b1576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b604080516348b4573960e11b81526001600160a01b0360208b81013582166004840152925190861692639168ae729260248082019391829003018186803b1580156122fb57600080fd5b505afa15801561230f573d6000803e3d6000fd5b505050506040513d602081101561232557600080fd5b505161236d576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b61238286356020880135604089013588613609565b846001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b1580156123bb57600080fd5b505afa1580156123cf573d6000803e3d6000fd5b505050506040513d60208110156123e557600080fd5b505114612425576040805162461bcd60e51b8152602060048201526009602482015268086908298be9082a6960bb1b604482015290519081900360640190fd5b600f54600b546040805163877c8c2b60e01b8152893560048201526020808b013560248301528a8301356044830152606482018a90526001600160a01b038d35811660848401528d820135811660a484015260c48301949094529151600094939093169263877c8c2b9260e48084019391929182900301818787803b1580156124ad57600080fd5b505af11580156124c1573d6000803e3d6000fd5b505050506040513d60208110156124d757600080fd5b5051600384810180546001600160a01b038085166001600160a01b03199283168117909355928601805490911682179055604080518d35841681526020808f0135909416938101939093528b358382015251929350917fa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879916060908290030190a2505050505050505050565b6125a560033384848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506133f892505050565b5050565b600b54600754430310156119cd576040805162461bcd60e51b815260206004820152600c60248201526b524543454e545f5354414b4560a01b604482015290519081900360640190fd5b60008460601b60601c6001600160a01b0316848484604051602001808581526020018481526020018383808284378083019250505094505050505060405160208183030381529060405280519060200120905061268a8989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508691506136479050565b6000856001600160a01b0316858585604051808383808284376040519201945060009350909150508083038185875af1925050503d80600081146126ea576040519150601f19603f3d011682016040523d82523d6000602084013e6126ef565b606091505b50509050806126fd57600080fd5b50505050505050505050565b60045481565b6010546001600160a01b031681565b60085490565b600a54815b8181101561283a576000600a828154811061274057fe5b906000526020600020906002020190505b6004548160010154101561283157600a600184038154811061276f57fe5b9060005260206000209060020201600a838154811061278a57fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600a8054806127cd57fe5b6000828152602081206000199283016002810290910180546001600160a01b031916815560010191909155909155929092019182821061280f57505050611fa4565b600a828154811061281c57fe5b90600052602060002090600202019050612751565b50600101612729565b505050565b3360009081526009602052604090206003810154600160a01b900460ff1661289b576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b858540146128e6576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b600554600101841461292a576040805162461bcd60e51b81526020600482015260086024820152674e4f44455f4e554d60c01b604482015290519081900360640190fd5b612932613b33565b6040805160e08181019092526129869186906007908390839080828437600092019190915250506040805161014081810190925291508690600a9083908390808284376000920191909152506136e5915050565b600183015460008181526006602090815260409182902054825163380ed4c760e11b8152925194955092936001600160a01b0390931692839263701da98e926004808301939192829003018186803b1580156129e157600080fd5b505afa1580156129f5573d6000803e3d6000fd5b505050506040513d6020811015612a0b57600080fd5b5051612a1684613804565b14612a5a576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b8260800151600154038361012001511115612aad576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b6000816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015612ae857600080fd5b505afa158015612afc573d6000803e3d6000fd5b505050506040513d6020811015612b1257600080fd5b5051845190915043036000612b25611524565b600c54909150820281831015612b6f576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b86608001518760e0015103876101200151101580612b9257508087610140015110155b612bcf576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b806004028761014001511115612c18576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b600b54430184811015612c285750835b6000600c5489610140015181612c3a57fe5b04905080820191506000601060009054906101000a90046001600160a01b03166001600160a01b031663d45ab2b5612c748c60015461383c565b612c848d6001546000548861388b565b612c8d8e613907565b8d886040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b158015612ce157600080fd5b505af1158015612cf5573d6000803e3d6000fd5b505050506040513d6020811015612d0b57600080fd5b5051600580546001019081905560009081526006602052604080822080546001600160a01b0319166001600160a01b038516908117909155815163123334b760e11b8152336004820152915193945092632466696e9260248084019391929182900301818387803b158015612d7f57600080fd5b505af1158015612d93573d6000803e3d6000fd5b505050506005548b600101819055506005547f4807480f255627d9b9350200277cb372949a76ccb4263935665257e2b3a1582a8e8e6001546000546040518085600760200280828437600083820152601f01601f191690910190508461014080828437600083820152601f01601f19169091019384525050602082015260408051918290030192509050a250505050505050505050505050505050565b60005481565b6001600160a01b0380831660009081526009602052604080822084841683529120600382015491929091163314612e6c57600080fd5b60038101546001600160a01b03163314612e8557600080fd5b816002015481600201541115612ee75760028083015490820154604051919003906001600160a01b0385169082156108fc029083906000818181858888f19350505050158015612ed9573d6000803e3d6000fd5b506002820180549190910390555b60028181015483820180549183900490910190556003830180546001600160a01b0319908116909155604080518082019091526001600160a01b03868116825260018086015460208401908152600a80549283018155600052925194027fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a88101805495909216949093169390931790925590517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a990910155611a7c816134cf565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6bffffffffffffffffffffffff191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008181526006602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b1580156130e857600080fd5b505af11580156130fc573d6000803e3d6000fd5b50505060009182525060066020526040902080546001600160a01b0319169055565b6003810154600160a01b900460ff1661316b576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b60038101546001600160a01b031615611fa4576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001546000805490918291826131cf88884342878b613002565b90506131db8282613070565b600055506001828101905590969095509350505050565b80518251600091829182805b838110156132a557600087828151811061321457fe5b60200260200101519050838187011115613264576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868a01810182902060408051808401969096528581019190915280518086038201815260609095019052835193019290922091909401936001016131fe565b508184146132e8576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b9695505050505050565b80516000805b8281101561118057600060ff1685838151811061331157fe5b016020015160f81c141561339757600061332e866001850161391d565b905060028160405161333f90613bbf565b90815260405190819003602001906000f080158015613362573d6000803e3d6000fd5b5081546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b03909216919091179055505b8381815181106133a357fe5b60200260200101518201915080806001019150506132f8565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906133f057508115155b949350505050565b60008061340d858585805190602001206131b5565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561348c578181015183820152602001613474565b50505050905090810190601f1680156134b95780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b80546008805460009190839081106134e357fe5b600091825260209091200154600880546001600160a01b03909216925090600019810190811061350f57fe5b600091825260209091200154600880546001600160a01b03909216918490811061353557fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600960006008858154811061357557fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205560088054806135a557fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0392909216815260099091526040812081815560018101829055600281019190915560030180546001600160a81b03191690555050565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b600160001b811890506000613660848385600101613976565b5090506002858154811061367057fe5b600091825260208220015460408051630ad0379b60e01b8152600481018590526024810187905290516001600160a01b0390921692630ad0379b9260448084019382900301818387803b1580156136c657600080fd5b505af11580156136da573d6000803e3d6000fd5b505050505050505050565b6136ed613b33565b60408051610220810182528351815260208085015181830152855182840152850151606080830191909152848301516080808401919091529085015160a0808401919091529085015160c083015284015160e082015290840151610100820152610120810183600660200201518152602001836007600a811061376c57fe5b602002015181526020018460036007811061378357fe5b60200201518152602001836008600a811061379a57fe5b60200201518152602001846004600781106137b157fe5b60200201518152602001836009600a81106137c857fe5b60200201518152602001846005600781106137df57fe5b60200201518152602001846006600781106137f657fe5b602002015190529392505050565b6000613836826000015183602001518460400151856060015186608001518760a001518860c001518960e00151612fa8565b92915050565b600061388443846101400151856020015101856102000151866101e001518761012001518860800151018861018001518960a0015101896101c001518a60c001510189612fa8565b9392505050565b6000806138ad6000876101200151886080015188030386896101e00151613609565b905060006138e660008861012001516138ce8a6101e001516000801b613070565b6138e18b606001518c6101000151613070565b613609565b90506138fc82826138f68a613a83565b87613609565b979650505050505050565b6000613836826101600151836101a00151613070565b6000816020018351101561396d576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60008080848160205b88518111613a75578089015193506020818a51036020018161399d57fe5b0491505b6000821180156139b45750600287066001145b80156139c257508160020a87115b156139da5760029096046001908101969401936139a1565b60028706613a25578383604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120925060028781613a1d57fe5b049650613a67565b8284604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120925060028781613a6057fe5b0460010196505b60019094019360200161397f565b509093505050935093915050565b60006138366000836101400151613ab96000613ab487610100015188604001516000801b60008060001b6000613ae8565b613070565b6138e1866101400151613ab46000801b8961020001518a61016001518b61018001518c6101a001518d6101c001515b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e0810182905261020081019190915290565b61013780613bcd8339019056fe608060405234801561001057600080fd5b506040516101373803806101378339818101604052602081101561003357600080fd5b5051600080546001600160a01b0319163317905560015560df806100586000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630ad0379b14602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135604f565b005b6000546001600160a01b03163314606557600080fd5b60008181526002602052604090205460ff1615608057600080fd5b6001548214608d57600080fd5b6000908152600260205260409020805460ff191660011790555056fea264697066735822122045771b255caea9c72008e9382dc6c60996b94002366693a4d85d84c26f81042064736f6c634300060c0033a2646970667358221220bc9429f0e21e7ed0d8801cdd78484890208a8cbf3850b33364561d2acf4436f664736f6c634300060c0033"

// DeployRollup deploys a new Ethereum contract, binding an instance of Rollup to it.
func DeployRollup(auth *bind.TransactOpts, backend bind.ContractBackend, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _challengeFactory common.Address, _nodeFactory common.Address, _extraConfig []byte) (common.Address, *types.Transaction, *Rollup, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupBin), backend, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _challengeFactory, _nodeFactory, _extraConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupCaller) ArbGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "arbGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _Rollup.Contract.ArbGasSpeedLimitPerBlock(&_Rollup.CallOpts)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupCallerSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _Rollup.Contract.ArbGasSpeedLimitPerBlock(&_Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupCaller) BaseStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "baseStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupSession) BaseStake() (*big.Int, error) {
	return _Rollup.Contract.BaseStake(&_Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupCallerSession) BaseStake() (*big.Int, error) {
	return _Rollup.Contract.BaseStake(&_Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupSession) ChallengeFactory() (common.Address, error) {
	return _Rollup.Contract.ChallengeFactory(&_Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupCallerSession) ChallengeFactory() (common.Address, error) {
	return _Rollup.Contract.ChallengeFactory(&_Rollup.CallOpts)
}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint256)
func (_Rollup *RollupCaller) ChallengePeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challengePeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint256)
func (_Rollup *RollupSession) ChallengePeriodBlocks() (*big.Int, error) {
	return _Rollup.Contract.ChallengePeriodBlocks(&_Rollup.CallOpts)
}

// ChallengePeriodBlocks is a free data retrieval call binding the contract method 0x46c2781a.
//
// Solidity: function challengePeriodBlocks() view returns(uint256)
func (_Rollup *RollupCallerSession) ChallengePeriodBlocks() (*big.Int, error) {
	return _Rollup.Contract.ChallengePeriodBlocks(&_Rollup.CallOpts)
}

// CheckNoRecentStake is a free data retrieval call binding the contract method 0xbe211c9a.
//
// Solidity: function checkNoRecentStake() view returns()
func (_Rollup *RollupCaller) CheckNoRecentStake(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "checkNoRecentStake")

	if err != nil {
		return err
	}

	return err

}

// CheckNoRecentStake is a free data retrieval call binding the contract method 0xbe211c9a.
//
// Solidity: function checkNoRecentStake() view returns()
func (_Rollup *RollupSession) CheckNoRecentStake() error {
	return _Rollup.Contract.CheckNoRecentStake(&_Rollup.CallOpts)
}

// CheckNoRecentStake is a free data retrieval call binding the contract method 0xbe211c9a.
//
// Solidity: function checkNoRecentStake() view returns()
func (_Rollup *RollupCallerSession) CheckNoRecentStake() error {
	return _Rollup.Contract.CheckNoRecentStake(&_Rollup.CallOpts)
}

// CheckUnresolved is a free data retrieval call binding the contract method 0x73f33b06.
//
// Solidity: function checkUnresolved() view returns()
func (_Rollup *RollupCaller) CheckUnresolved(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "checkUnresolved")

	if err != nil {
		return err
	}

	return err

}

// CheckUnresolved is a free data retrieval call binding the contract method 0x73f33b06.
//
// Solidity: function checkUnresolved() view returns()
func (_Rollup *RollupSession) CheckUnresolved() error {
	return _Rollup.Contract.CheckUnresolved(&_Rollup.CallOpts)
}

// CheckUnresolved is a free data retrieval call binding the contract method 0x73f33b06.
//
// Solidity: function checkUnresolved() view returns()
func (_Rollup *RollupCallerSession) CheckUnresolved() error {
	return _Rollup.Contract.CheckUnresolved(&_Rollup.CallOpts)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupCaller) CountStakedZombies(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "countStakedZombies", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _Rollup.Contract.CountStakedZombies(&_Rollup.CallOpts, node)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupCallerSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _Rollup.Contract.CountStakedZombies(&_Rollup.CallOpts, node)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _Rollup.Contract.CurrentRequiredStake(&_Rollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _Rollup.Contract.CurrentRequiredStake(&_Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupSession) FirstUnresolvedNode() (*big.Int, error) {
	return _Rollup.Contract.FirstUnresolvedNode(&_Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _Rollup.Contract.FirstUnresolvedNode(&_Rollup.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 startIndex, uint256 max) view returns(address[])
func (_Rollup *RollupCaller) GetStakers(opts *bind.CallOpts, startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getStakers", startIndex, max)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 startIndex, uint256 max) view returns(address[])
func (_Rollup *RollupSession) GetStakers(startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	return _Rollup.Contract.GetStakers(&_Rollup.CallOpts, startIndex, max)
}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 startIndex, uint256 max) view returns(address[])
func (_Rollup *RollupCallerSession) GetStakers(startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	return _Rollup.Contract.GetStakers(&_Rollup.CallOpts, startIndex, max)
}

// InboxMaxAcc is a free data retrieval call binding the contract method 0xf0a79973.
//
// Solidity: function inboxMaxAcc() view returns(bytes32)
func (_Rollup *RollupCaller) InboxMaxAcc(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "inboxMaxAcc")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxMaxAcc is a free data retrieval call binding the contract method 0xf0a79973.
//
// Solidity: function inboxMaxAcc() view returns(bytes32)
func (_Rollup *RollupSession) InboxMaxAcc() ([32]byte, error) {
	return _Rollup.Contract.InboxMaxAcc(&_Rollup.CallOpts)
}

// InboxMaxAcc is a free data retrieval call binding the contract method 0xf0a79973.
//
// Solidity: function inboxMaxAcc() view returns(bytes32)
func (_Rollup *RollupCallerSession) InboxMaxAcc() ([32]byte, error) {
	return _Rollup.Contract.InboxMaxAcc(&_Rollup.CallOpts)
}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Rollup *RollupCaller) InboxMaxCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "inboxMaxCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Rollup *RollupSession) InboxMaxCount() (*big.Int, error) {
	return _Rollup.Contract.InboxMaxCount(&_Rollup.CallOpts)
}

// InboxMaxCount is a free data retrieval call binding the contract method 0x917cae02.
//
// Solidity: function inboxMaxCount() view returns(uint256)
func (_Rollup *RollupCallerSession) InboxMaxCount() (*big.Int, error) {
	return _Rollup.Contract.InboxMaxCount(&_Rollup.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_Rollup *RollupCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_Rollup *RollupSession) LastStakeBlock() (*big.Int, error) {
	return _Rollup.Contract.LastStakeBlock(&_Rollup.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_Rollup *RollupCallerSession) LastStakeBlock() (*big.Int, error) {
	return _Rollup.Contract.LastStakeBlock(&_Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupSession) LatestConfirmed() (*big.Int, error) {
	return _Rollup.Contract.LatestConfirmed(&_Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupCallerSession) LatestConfirmed() (*big.Int, error) {
	return _Rollup.Contract.LatestConfirmed(&_Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupSession) LatestNodeCreated() (*big.Int, error) {
	return _Rollup.Contract.LatestNodeCreated(&_Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _Rollup.Contract.LatestNodeCreated(&_Rollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _Rollup.Contract.MinimumAssertionPeriod(&_Rollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _Rollup.Contract.MinimumAssertionPeriod(&_Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupSession) NodeFactory() (common.Address, error) {
	return _Rollup.Contract.NodeFactory(&_Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupCallerSession) NodeFactory() (common.Address, error) {
	return _Rollup.Contract.NodeFactory(&_Rollup.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_Rollup *RollupCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "nodes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_Rollup *RollupSession) Nodes(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.Nodes(&_Rollup.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_Rollup *RollupCallerSession) Nodes(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.Nodes(&_Rollup.CallOpts, arg0)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupSession) StakeToken() (common.Address, error) {
	return _Rollup.Contract.StakeToken(&_Rollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Rollup *RollupCallerSession) StakeToken() (common.Address, error) {
	return _Rollup.Contract.StakeToken(&_Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupSession) StakerCount() (*big.Int, error) {
	return _Rollup.Contract.StakerCount(&_Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupCallerSession) StakerCount() (*big.Int, error) {
	return _Rollup.Contract.StakerCount(&_Rollup.CallOpts)
}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_Rollup *RollupCaller) StakerInfo(opts *bind.CallOpts, stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakerInfo", stakerAddress)

	outstruct := new(struct {
		IsStaked         bool
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
	})

	outstruct.IsStaked = out[0].(bool)
	outstruct.LatestStakedNode = out[1].(*big.Int)
	outstruct.AmountStaked = out[2].(*big.Int)
	outstruct.CurrentChallenge = out[3].(common.Address)

	return *outstruct, err

}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_Rollup *RollupSession) StakerInfo(stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	return _Rollup.Contract.StakerInfo(&_Rollup.CallOpts, stakerAddress)
}

// StakerInfo is a free data retrieval call binding the contract method 0x4e745f1f.
//
// Solidity: function stakerInfo(address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_Rollup *RollupCallerSession) StakerInfo(stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	return _Rollup.Contract.StakerInfo(&_Rollup.CallOpts, stakerAddress)
}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_Rollup *RollupCaller) StakerList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakerList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_Rollup *RollupSession) StakerList(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.StakerList(&_Rollup.CallOpts, arg0)
}

// StakerList is a free data retrieval call binding the contract method 0x348e50c6.
//
// Solidity: function stakerList(uint256 ) view returns(address)
func (_Rollup *RollupCallerSession) StakerList(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.StakerList(&_Rollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0x729cfe3b.
//
// Solidity: function stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})

	outstruct.Index = out[0].(*big.Int)
	outstruct.LatestStakedNode = out[1].(*big.Int)
	outstruct.AmountStaked = out[2].(*big.Int)
	outstruct.CurrentChallenge = out[3].(common.Address)
	outstruct.IsStaked = out[4].(bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0x729cfe3b.
//
// Solidity: function stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _Rollup.Contract.StakerMap(&_Rollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0x729cfe3b.
//
// Solidity: function stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _Rollup.Contract.StakerMap(&_Rollup.CallOpts, arg0)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Rollup *RollupCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Rollup *RollupSession) ZombieCount() (*big.Int, error) {
	return _Rollup.Contract.ZombieCount(&_Rollup.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Rollup *RollupCallerSession) ZombieCount() (*big.Int, error) {
	return _Rollup.Contract.ZombieCount(&_Rollup.CallOpts)
}

// ZombieInfo is a free data retrieval call binding the contract method 0x4a95e20e.
//
// Solidity: function zombieInfo(uint256 index) view returns(address stakerAddress, uint256 latestStakedNode)
func (_Rollup *RollupCaller) ZombieInfo(opts *bind.CallOpts, index *big.Int) (struct {
	StakerAddress    common.Address
	LatestStakedNode *big.Int
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieInfo", index)

	outstruct := new(struct {
		StakerAddress    common.Address
		LatestStakedNode *big.Int
	})

	outstruct.StakerAddress = out[0].(common.Address)
	outstruct.LatestStakedNode = out[1].(*big.Int)

	return *outstruct, err

}

// ZombieInfo is a free data retrieval call binding the contract method 0x4a95e20e.
//
// Solidity: function zombieInfo(uint256 index) view returns(address stakerAddress, uint256 latestStakedNode)
func (_Rollup *RollupSession) ZombieInfo(index *big.Int) (struct {
	StakerAddress    common.Address
	LatestStakedNode *big.Int
}, error) {
	return _Rollup.Contract.ZombieInfo(&_Rollup.CallOpts, index)
}

// ZombieInfo is a free data retrieval call binding the contract method 0x4a95e20e.
//
// Solidity: function zombieInfo(uint256 index) view returns(address stakerAddress, uint256 latestStakedNode)
func (_Rollup *RollupCallerSession) ZombieInfo(index *big.Int) (struct {
	StakerAddress    common.Address
	LatestStakedNode *big.Int
}, error) {
	return _Rollup.Contract.ZombieInfo(&_Rollup.CallOpts, index)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupTransactor) AddToDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addToDeposit", stakerAddress)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupSession) AddToDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddToDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupTransactorSession) AddToDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddToDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "completeChallenge", winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.CompleteChallenge(&_Rollup.TransactOpts, winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupTransactorSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.CompleteChallenge(&_Rollup.TransactOpts, winningStaker, losingStaker)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0x396b8cbc.
//
// Solidity: function confirmNextNode(bytes32 logAcc, bytes sendsData, uint256[] sendLengths) returns()
func (_Rollup *RollupTransactor) ConfirmNextNode(opts *bind.TransactOpts, logAcc [32]byte, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "confirmNextNode", logAcc, sendsData, sendLengths)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0x396b8cbc.
//
// Solidity: function confirmNextNode(bytes32 logAcc, bytes sendsData, uint256[] sendLengths) returns()
func (_Rollup *RollupSession) ConfirmNextNode(logAcc [32]byte, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ConfirmNextNode(&_Rollup.TransactOpts, logAcc, sendsData, sendLengths)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0x396b8cbc.
//
// Solidity: function confirmNextNode(bytes32 logAcc, bytes sendsData, uint256[] sendLengths) returns()
func (_Rollup *RollupTransactorSession) ConfirmNextNode(logAcc [32]byte, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ConfirmNextNode(&_Rollup.TransactOpts, logAcc, sendsData, sendLengths)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xb1b4181c.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[3] nodeFields, uint256 executionCheckTime) returns()
func (_Rollup *RollupTransactor) CreateChallenge(opts *bind.TransactOpts, stakers [2]common.Address, nodeNums [2]*big.Int, nodeFields [3][32]byte, executionCheckTime *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "createChallenge", stakers, nodeNums, nodeFields, executionCheckTime)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xb1b4181c.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[3] nodeFields, uint256 executionCheckTime) returns()
func (_Rollup *RollupSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, nodeFields [3][32]byte, executionCheckTime *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, stakers, nodeNums, nodeFields, executionCheckTime)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xb1b4181c.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[3] nodeFields, uint256 executionCheckTime) returns()
func (_Rollup *RollupTransactorSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, nodeFields [3][32]byte, executionCheckTime *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, stakers, nodeNums, nodeFields, executionCheckTime)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Rollup *RollupTransactor) DeployL2ContractPair(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "deployL2ContractPair", maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Rollup *RollupSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.DeployL2ContractPair(&_Rollup.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Rollup *RollupTransactorSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.DeployL2ContractPair(&_Rollup.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Rollup *RollupTransactor) DepositEthMessage(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "depositEthMessage", to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Rollup *RollupSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.DepositEthMessage(&_Rollup.TransactOpts, to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Rollup *RollupTransactorSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.DepositEthMessage(&_Rollup.TransactOpts, to)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Rollup *RollupTransactor) ExecuteTransaction(opts *bind.TransactOpts, outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "executeTransaction", outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Rollup *RollupSession) ExecuteTransaction(outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ExecuteTransaction(&_Rollup.TransactOpts, outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc4fb000c.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes _proof, uint256 _index, address destAddr, uint256 amount, bytes calldataForL1) returns()
func (_Rollup *RollupTransactorSession) ExecuteTransaction(outboxIndex *big.Int, _proof []byte, _index *big.Int, destAddr common.Address, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ExecuteTransaction(&_Rollup.TransactOpts, outboxIndex, _proof, _index, destAddr, amount, calldataForL1)
}

// NewStake is a paid mutator transaction binding the contract method 0x5f576db6.
//
// Solidity: function newStake() payable returns()
func (_Rollup *RollupTransactor) NewStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "newStake")
}

// NewStake is a paid mutator transaction binding the contract method 0x5f576db6.
//
// Solidity: function newStake() payable returns()
func (_Rollup *RollupSession) NewStake() (*types.Transaction, error) {
	return _Rollup.Contract.NewStake(&_Rollup.TransactOpts)
}

// NewStake is a paid mutator transaction binding the contract method 0x5f576db6.
//
// Solidity: function newStake() payable returns()
func (_Rollup *RollupTransactorSession) NewStake() (*types.Transaction, error) {
	return _Rollup.Contract.NewStake(&_Rollup.TransactOpts)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 maxReduction) returns()
func (_Rollup *RollupTransactor) ReduceDeposit(opts *bind.TransactOpts, maxReduction *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "reduceDeposit", maxReduction)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 maxReduction) returns()
func (_Rollup *RollupSession) ReduceDeposit(maxReduction *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ReduceDeposit(&_Rollup.TransactOpts, maxReduction)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 maxReduction) returns()
func (_Rollup *RollupTransactorSession) ReduceDeposit(maxReduction *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ReduceDeposit(&_Rollup.TransactOpts, maxReduction)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x0e1ef04c.
//
// Solidity: function rejectNextNode(uint256 successorWithStake, address stakerAddress) returns()
func (_Rollup *RollupTransactor) RejectNextNode(opts *bind.TransactOpts, successorWithStake *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "rejectNextNode", successorWithStake, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x0e1ef04c.
//
// Solidity: function rejectNextNode(uint256 successorWithStake, address stakerAddress) returns()
func (_Rollup *RollupSession) RejectNextNode(successorWithStake *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RejectNextNode(&_Rollup.TransactOpts, successorWithStake, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x0e1ef04c.
//
// Solidity: function rejectNextNode(uint256 successorWithStake, address stakerAddress) returns()
func (_Rollup *RollupTransactorSession) RejectNextNode(successorWithStake *big.Int, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RejectNextNode(&_Rollup.TransactOpts, successorWithStake, stakerAddress)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupTransactor) RemoveOldZombies(opts *bind.TransactOpts, startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeOldZombies", startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldZombies(&_Rollup.TransactOpts, startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupTransactorSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldZombies(&_Rollup.TransactOpts, startIndex)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupTransactor) RemoveZombie(opts *bind.TransactOpts, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeZombie", zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveZombie(&_Rollup.TransactOpts, zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupTransactorSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveZombie(&_Rollup.TransactOpts, zombieNum, maxNodes)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupTransactor) ReturnOldDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "returnOldDeposit", stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ReturnOldDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupTransactorSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ReturnOldDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Rollup *RollupTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Rollup *RollupSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.SendL2Message(&_Rollup.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Rollup *RollupTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.SendL2Message(&_Rollup.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Rollup *RollupTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Rollup *RollupSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.SendL2MessageFromOrigin(&_Rollup.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Rollup *RollupTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Rollup.Contract.SendL2MessageFromOrigin(&_Rollup.TransactOpts, messageData)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x8fd18f04.
//
// Solidity: function stakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) returns()
func (_Rollup *RollupTransactor) StakeOnExistingNode(opts *bind.TransactOpts, blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "stakeOnExistingNode", blockHash, blockNumber, nodeNum)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x8fd18f04.
//
// Solidity: function stakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) returns()
func (_Rollup *RollupSession) StakeOnExistingNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnExistingNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x8fd18f04.
//
// Solidity: function stakeOnExistingNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum) returns()
func (_Rollup *RollupTransactorSession) StakeOnExistingNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnExistingNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0xf019a1c1.
//
// Solidity: function stakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupTransactor) StakeOnNewNode(opts *bind.TransactOpts, blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "stakeOnNewNode", blockHash, blockNumber, nodeNum, assertionBytes32Fields, assertionIntFields)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0xf019a1c1.
//
// Solidity: function stakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupSession) StakeOnNewNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnNewNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum, assertionBytes32Fields, assertionIntFields)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0xf019a1c1.
//
// Solidity: function stakeOnNewNode(bytes32 blockHash, uint256 blockNumber, uint256 nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupTransactorSession) StakeOnNewNode(blockHash [32]byte, blockNumber *big.Int, nodeNum *big.Int, assertionBytes32Fields [7][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnNewNode(&_Rollup.TransactOpts, blockHash, blockNumber, nodeNum, assertionBytes32Fields, assertionIntFields)
}

// RollupBuddyContractPairIterator is returned from FilterBuddyContractPair and is used to iterate over the raw logs and unpacked data for BuddyContractPair events raised by the Rollup contract.
type RollupBuddyContractPairIterator struct {
	Event *RollupBuddyContractPair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupBuddyContractPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBuddyContractPair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupBuddyContractPair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupBuddyContractPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBuddyContractPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBuddyContractPair represents a BuddyContractPair event raised by the Rollup contract.
type RollupBuddyContractPair struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuddyContractPair is a free log retrieval operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Rollup *RollupFilterer) FilterBuddyContractPair(opts *bind.FilterOpts, sender []common.Address) (*RollupBuddyContractPairIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return &RollupBuddyContractPairIterator{contract: _Rollup.contract, event: "BuddyContractPair", logs: logs, sub: sub}, nil
}

// WatchBuddyContractPair is a free log subscription operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Rollup *RollupFilterer) WatchBuddyContractPair(opts *bind.WatchOpts, sink chan<- *RollupBuddyContractPair, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBuddyContractPair)
				if err := _Rollup.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuddyContractPair is a log parse operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Rollup *RollupFilterer) ParseBuddyContractPair(log types.Log) (*RollupBuddyContractPair, error) {
	event := new(RollupBuddyContractPair)
	if err := _Rollup.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the Rollup contract.
type RollupMessageDeliveredIterator struct {
	Event *RollupMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMessageDelivered represents a MessageDelivered event raised by the Rollup contract.
type RollupMessageDelivered struct {
	MessageNum     *big.Int
	BeforeInboxAcc [32]byte
	Kind           uint8
	Sender         common.Address
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Rollup *RollupFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int, beforeInboxAcc [][32]byte) (*RollupMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "MessageDelivered", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &RollupMessageDeliveredIterator{contract: _Rollup.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Rollup *RollupFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *RollupMessageDelivered, messageNum []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "MessageDelivered", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMessageDelivered)
				if err := _Rollup.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0xfc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003.
//
// Solidity: event MessageDelivered(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender, bytes data)
func (_Rollup *RollupFilterer) ParseMessageDelivered(log types.Log) (*RollupMessageDelivered, error) {
	event := new(RollupMessageDelivered)
	if err := _Rollup.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMessageDeliveredFromOriginIterator is returned from FilterMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for MessageDeliveredFromOrigin events raised by the Rollup contract.
type RollupMessageDeliveredFromOriginIterator struct {
	Event *RollupMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMessageDeliveredFromOrigin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupMessageDeliveredFromOrigin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMessageDeliveredFromOrigin represents a MessageDeliveredFromOrigin event raised by the Rollup contract.
type RollupMessageDeliveredFromOrigin struct {
	MessageNum     *big.Int
	BeforeInboxAcc [32]byte
	Kind           uint8
	Sender         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Rollup *RollupFilterer) FilterMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int, beforeInboxAcc [][32]byte) (*RollupMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "MessageDeliveredFromOrigin", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &RollupMessageDeliveredFromOriginIterator{contract: _Rollup.contract, event: "MessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Rollup *RollupFilterer) WatchMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *RollupMessageDeliveredFromOrigin, messageNum []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "MessageDeliveredFromOrigin", messageNumRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMessageDeliveredFromOrigin)
				if err := _Rollup.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDeliveredFromOrigin is a log parse operation binding the contract event 0x852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab.
//
// Solidity: event MessageDeliveredFromOrigin(uint256 indexed messageNum, bytes32 indexed beforeInboxAcc, uint8 kind, address sender)
func (_Rollup *RollupFilterer) ParseMessageDeliveredFromOrigin(log types.Log) (*RollupMessageDeliveredFromOrigin, error) {
	event := new(RollupMessageDeliveredFromOrigin)
	if err := _Rollup.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the Rollup contract.
type RollupNodeCreatedIterator struct {
	Event *RollupNodeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNodeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupNodeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNodeCreated represents a NodeCreated event raised by the Rollup contract.
type RollupNodeCreated struct {
	NodeNum                *big.Int
	AssertionBytes32Fields [7][32]byte
	AssertionIntFields     [10]*big.Int
	InboxMaxCount          *big.Int
	InboxMaxHash           [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x4807480f255627d9b9350200277cb372949a76ccb4263935665257e2b3a1582a.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields, uint256 inboxMaxCount, bytes32 inboxMaxHash)
func (_Rollup *RollupFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int) (*RollupNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NodeCreated", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupNodeCreatedIterator{contract: _Rollup.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x4807480f255627d9b9350200277cb372949a76ccb4263935665257e2b3a1582a.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields, uint256 inboxMaxCount, bytes32 inboxMaxHash)
func (_Rollup *RollupFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *RollupNodeCreated, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NodeCreated", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNodeCreated)
				if err := _Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeCreated is a log parse operation binding the contract event 0x4807480f255627d9b9350200277cb372949a76ccb4263935665257e2b3a1582a.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32[7] assertionBytes32Fields, uint256[10] assertionIntFields, uint256 inboxMaxCount, bytes32 inboxMaxHash)
func (_Rollup *RollupFilterer) ParseNodeCreated(log types.Log) (*RollupNodeCreated, error) {
	event := new(RollupNodeCreated)
	if err := _Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the Rollup contract.
type RollupRollupChallengeStartedIterator struct {
	Event *RollupRollupChallengeStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRollupChallengeStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupRollupChallengeStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the Rollup contract.
type RollupRollupChallengeStarted struct {
	ChallengeContract common.Address
	Asserter          common.Address
	Challenger        common.Address
	ChallengedNode    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_Rollup *RollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts, challengeContract []common.Address) (*RollupRollupChallengeStartedIterator, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return &RollupRollupChallengeStartedIterator{contract: _Rollup.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_Rollup *RollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *RollupRollupChallengeStarted, challengeContract []common.Address) (event.Subscription, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRollupChallengeStarted)
				if err := _Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_Rollup *RollupFilterer) ParseRollupChallengeStarted(log types.Log) (*RollupRollupChallengeStarted, error) {
	event := new(RollupRollupChallengeStarted)
	if err := _Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the Rollup contract.
type RollupRollupCreatedIterator struct {
	Event *RollupRollupCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRollupCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupRollupCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRollupCreated represents a RollupCreated event raised by the Rollup contract.
type RollupRollupCreated struct {
	MachineHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_Rollup *RollupFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*RollupRollupCreatedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &RollupRollupCreatedIterator{contract: _Rollup.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_Rollup *RollupFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupRollupCreated) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRollupCreated)
				if err := _Rollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupCreated is a log parse operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_Rollup *RollupFilterer) ParseRollupCreated(log types.Log) (*RollupRollupCreated, error) {
	event := new(RollupRollupCreated)
	if err := _Rollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupSentLogsIterator is returned from FilterSentLogs and is used to iterate over the raw logs and unpacked data for SentLogs events raised by the Rollup contract.
type RollupSentLogsIterator struct {
	Event *RollupSentLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupSentLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupSentLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupSentLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupSentLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupSentLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupSentLogs represents a SentLogs event raised by the Rollup contract.
type RollupSentLogs struct {
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSentLogs is a free log retrieval operation binding the contract event 0xe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e84.
//
// Solidity: event SentLogs(bytes32 logsAccHash)
func (_Rollup *RollupFilterer) FilterSentLogs(opts *bind.FilterOpts) (*RollupSentLogsIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "SentLogs")
	if err != nil {
		return nil, err
	}
	return &RollupSentLogsIterator{contract: _Rollup.contract, event: "SentLogs", logs: logs, sub: sub}, nil
}

// WatchSentLogs is a free log subscription operation binding the contract event 0xe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e84.
//
// Solidity: event SentLogs(bytes32 logsAccHash)
func (_Rollup *RollupFilterer) WatchSentLogs(opts *bind.WatchOpts, sink chan<- *RollupSentLogs) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "SentLogs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupSentLogs)
				if err := _Rollup.contract.UnpackLog(event, "SentLogs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSentLogs is a log parse operation binding the contract event 0xe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e84.
//
// Solidity: event SentLogs(bytes32 logsAccHash)
func (_Rollup *RollupFilterer) ParseSentLogs(log types.Log) (*RollupSentLogs, error) {
	event := new(RollupSentLogs)
	if err := _Rollup.contract.UnpackLog(event, "SentLogs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorABI is the input ABI used to generate the binding from.
const RollupCreatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"name\":\"createRollup\",\"outputs\":[{\"internalType\":\"contractRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupCreatorFuncSigs maps the 4-byte function signature to its string representation.
var RollupCreatorFuncSigs = map[string]string{
	"d798f5be": "createRollup(bytes32,uint256,uint256,uint256,address,address,bytes)",
}

// RollupCreatorBin is the compiled bytecode used for deploying new contracts.
var RollupCreatorBin = "0x608060405234801561001057600080fd5b506040516146213803806146218339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556145a78061007a6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063d798f5be14610030575b600080fd5b6100d1600480360360e081101561004657600080fd5b8135916020810135916040820135916060810135916001600160a01b03608083013581169260a08101359091169181019060e0810160c082013564010000000081111561009257600080fd5b8201836020820111156100a457600080fd5b803590602001918460018302840111640100000000831117156100c657600080fd5b5090925090506100ed565b604080516001600160a01b039092168252519081900360200190f35b6000805460015460405183928c928c928c928c928c928c926001600160a01b039081169216908c908c9061012090610219565b808b81526020018a8152602001898152602001888152602001876001600160a01b03168152602001866001600160a01b03168152602001856001600160a01b03168152602001846001600160a01b031681526020018060200182810382528484828181526020019250808284376000838201819052604051601f909201601f19169093018190039e509c50909a5050505050505050505050f0801580156101cb573d6000803e3d6000fd5b50604080516001600160a01b038316815290519192507f84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c919081900360200190a19998505050505050505050565b61434b806102278339019056fe60806040523480156200001157600080fd5b506040516200434b3803806200434b83398181016040526101208110156200003857600080fd5b815160208301516040808501516060860151608087015160a088015160c089015160e08a01516101008b0180519751999b989a969995989497939692959194919392820192846401000000008211156200009157600080fd5b908301906020820185811115620000a757600080fd5b8251640100000000811182820188101715620000c257600080fd5b82525081516020918201929091019080838360005b83811015620000f1578181015183820152602001620000d7565b50505050905090810190601f1680156200011f5780820380516001836020036101000a031916815260200191505b5060405250505082600f60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081601060006101000a8154816001600160a01b0302191690836001600160a01b031602179055506200022f8888888860601b6001600160601b0319168860601b6001600160601b031916866040516020018087815260200186815260200185815260200184815260200183815260200182805190602001908083835b60208310620001ea5780518252601f199092019160209182019101620001c9565b6001836020036101000a0380198251168184511680821785525050505050509050019650505050505050604051602081830303815290604052620003ab60201b60201c565b6000620002564360008c6000801b6000806000600154620003bc60201b62002fa81760201c565b6010546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905260848201819052915193945090926001600160a01b039092169163d45ab2b59160a48082019260209290919082900301818787803b158015620002c757600080fd5b505af1158015620002dc573d6000803e3d6000fd5b505050506040513d6020811015620002f357600080fd5b505160008052600660209081527f54cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f880546001600160a01b038085166001600160a01b031992831617909255600b8e9055600c8d9055600d8c9055600e8054928c16929091169190911790556001600455604080518e815290519293507f4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d929081900390910190a1505050505050505050505062000602565b620003b96004308362000416565b50565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b6000806200043385858580519060200120620004f860201b60201c565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015620004b45781810151838201526020016200049a565b50505050905090810190601f168015620004e25780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b6000806000600154905060008054905060006200052588884342878b6200055560201b620030021760201c565b90506200053e8282620005d660201b620030701760201c565b600055506001828101905590969095509350505050565b6040805160f89790971b7fff000000000000000000000000000000000000000000000000000000000000001660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b613d3980620006126000396000f3fe6080604052600436106102515760003560e01c806373f33b0611610139578063b1b4181c116100b6578063d93fe9c41161007a578063d93fe9c414610aac578063dff6978714610ac1578063edfd03ed14610ad6578063f019a1c114610b00578063f0a7997314610b42578063fa7803e614610b5757610251565b8063b1b4181c146108e3578063b75436bb1461091a578063be211c9a14610995578063c4fb000c146109aa578063d735e21d14610a9757610251565b80638640ce5f116100fd5780638640ce5f146107dd5780638fd18f04146107f2578063917cae0214610828578063ad71bd361461083d578063afcc220b146108bd57610251565b806373f33b061461073b5780637427be511461075057806376e7e23b146107835780637ba9534a146107985780637e2d2155146107ad57610251565b80634a95e20e116101d25780635e8ef106116101965780635e8ef106146105fd5780635f576db61461061257806363721d6b1461061a57806365f7f80d1461062f5780636f5dfdca14610644578063729cfe3b146106d257610251565b80634a95e20e1461050d5780634d26732d1461055a5780634e745f1f1461056f57806351ed6a30146105d35780635dbaf68b146105e857610251565b8063348e50c611610219578063348e50c6146103c1578063396b8cbc146103eb57806345c5b2c7146104bd57806345e38b64146104e357806346c2781a146104f857610251565b806304a28064146102565780630e1ef04c1461029b5780631c53c280146102d65780631e83d30f1461031c5780631fe927cf14610346575b600080fd5b34801561026257600080fd5b506102896004803603602081101561027957600080fd5b50356001600160a01b0316610b92565b60408051918252519081900360200190f35b3480156102a757600080fd5b506102d4600480360360408110156102be57600080fd5b50803590602001356001600160a01b0316610c55565b005b3480156102e257600080fd5b50610300600480360360208110156102f957600080fd5b50356110f1565b604080516001600160a01b039092168252519081900360200190f35b34801561032857600080fd5b506102d46004803603602081101561033f57600080fd5b503561110c565b34801561035257600080fd5b506102d46004803603602081101561036957600080fd5b810190602081018135600160201b81111561038357600080fd5b82018360208201111561039557600080fd5b803590602001918460018302840111600160201b831117156103b657600080fd5b509092509050611187565b3480156103cd57600080fd5b50610300600480360360208110156103e457600080fd5b503561123d565b3480156103f757600080fd5b506102d46004803603606081101561040e57600080fd5b81359190810190604081016020820135600160201b81111561042f57600080fd5b82018360208201111561044157600080fd5b803590602001918460018302840111600160201b8311171561046257600080fd5b919390929091602081019035600160201b81111561047f57600080fd5b82018360208201111561049157600080fd5b803590602001918460208302840111600160201b831117156104b257600080fd5b509092509050611264565b6102d4600480360360208110156104d357600080fd5b50356001600160a01b03166114f7565b3480156104ef57600080fd5b50610289611524565b34801561050457600080fd5b50610289611539565b34801561051957600080fd5b506105376004803603602081101561053057600080fd5b503561153f565b604080516001600160a01b03909316835260208301919091528051918290030190f35b34801561056657600080fd5b5061028961157e565b34801561057b57600080fd5b506105a26004803603602081101561059257600080fd5b50356001600160a01b0316611679565b6040805194151585526020850193909352838301919091526001600160a01b03166060830152519081900360800190f35b3480156105df57600080fd5b506103006116b6565b3480156105f457600080fd5b506103006116c5565b34801561060957600080fd5b506102896116d4565b6102d46116da565b34801561062657600080fd5b50610289611855565b34801561063b57600080fd5b5061028961185b565b34801561065057600080fd5b506102d46004803603608081101561066757600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561069457600080fd5b8201836020820111156106a657600080fd5b803590602001918460018302840111600160201b831117156106c757600080fd5b509092509050611861565b3480156106de57600080fd5b50610705600480360360208110156106f557600080fd5b50356001600160a01b0316611939565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b34801561074757600080fd5b506102d4611975565b34801561075c57600080fd5b506102d46004803603602081101561077357600080fd5b50356001600160a01b03166119cf565b34801561078f57600080fd5b50610289611a82565b3480156107a457600080fd5b50610289611a88565b3480156107b957600080fd5b506102d4600480360360408110156107d057600080fd5b5080359060200135611a8e565b3480156107e957600080fd5b50610289611ca9565b3480156107fe57600080fd5b506102d46004803603606081101561081557600080fd5b5080359060208101359060400135611caf565b34801561083457600080fd5b50610289611e9f565b34801561084957600080fd5b5061086d6004803603604081101561086057600080fd5b5080359060200135611ea5565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156108a9578181015183820152602001610891565b505050509050019250505060405180910390f35b6102d4600480360360208110156108d357600080fd5b50356001600160a01b0316611f6d565b3480156108ef57600080fd5b506102d4600480360361010081101561090757600080fd5b50604081016080820160e0830135611fa7565b34801561092657600080fd5b506102d46004803603602081101561093d57600080fd5b810190602081018135600160201b81111561095757600080fd5b82018360208201111561096957600080fd5b803590602001918460018302840111600160201b8311171561098a57600080fd5b509092509050612563565b3480156109a157600080fd5b506102d46125a9565b3480156109b657600080fd5b506102d4600480360360c08110156109cd57600080fd5b81359190810190604081016020820135600160201b8111156109ee57600080fd5b820183602082011115610a0057600080fd5b803590602001918460018302840111600160201b83111715610a2157600080fd5b919390928235926001600160a01b03602082013516926040820135929091608081019060600135600160201b811115610a5957600080fd5b820183602082011115610a6b57600080fd5b803590602001918460018302840111600160201b83111715610a8c57600080fd5b5090925090506125f3565b348015610aa357600080fd5b50610289612709565b348015610ab857600080fd5b5061030061270f565b348015610acd57600080fd5b5061028961271e565b348015610ae257600080fd5b506102d460048036036020811015610af957600080fd5b5035612724565b348015610b0c57600080fd5b506102d46004803603610280811015610b2457600080fd5b5080359060208101359060408101359060608101906101400161283f565b348015610b4e57600080fd5b50610289612e30565b348015610b6357600080fd5b506102d460048036036040811015610b7a57600080fd5b506001600160a01b0381358116916020013516612e36565b600a5460009081805b82811015610c4d576000600a8281548110610bb257fe5b60009182526020918290206002909102018054604080516348b4573960e11b81526001600160a01b039283166004820152905192945090891692639168ae7292602480840193829003018186803b158015610c0c57600080fd5b505afa158015610c20573d6000803e3d6000fd5b505050506040513d6020811015610c3657600080fd5b505115610c44576001909201915b50600101610b9b565b509392505050565b610c5d611975565b600060066000600454815260200190815260200160002060009054906101000a90046001600160a01b03169050600354816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610cc657600080fd5b505afa158015610cda573d6000803e3d6000fd5b505050506040513d6020811015610cf057600080fd5b505114156110d857610d006125a9565b6004548311610d49576040805162461bcd60e51b815260206004820152601060248201526f535543434553534f525f544f5f4c4f5760801b604482015290519081900360640190fd5b600554831115610d94576040805162461bcd60e51b81526020600482015260116024820152700a6aa86868aa6a69ea4bea89ebe90928e9607b1b604482015290519081900360640190fd5b6001600160a01b038216600090815260096020526040902060030154600160a01b900460ff16610df8576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6000838152600660209081526040918290205460035483516311e7249560e21b815293516001600160a01b03909216939092849263479c9254926004808201939291829003018186803b158015610e4e57600080fd5b505afa158015610e62573d6000803e3d6000fd5b505050506040513d6020811015610e7857600080fd5b505114610ebc576040805162461bcd60e51b815260206004820152600d60248201526c2120a22fa9aaa1a1a2a9a9a7a960991b604482015290519081900360640190fd5b806001600160a01b0316639168ae72846040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610f0957600080fd5b505afa158015610f1d573d6000803e3d6000fd5b505050506040513d6020811015610f3357600080fd5b5051610f73576040805162461bcd60e51b815260206004820152600a6024820152692120a22fa9aa20a5a2a960b11b604482015290519081900360640190fd5b610f7d6000612724565b816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610fb657600080fd5b505afa158015610fca573d6000803e3d6000fd5b505050506040513d6020811015610fe057600080fd5b5051431015611028576040805162461bcd60e51b815260206004820152600f60248201526e4245464f52455f444541444c494e4560881b604482015290519081900360640190fd5b61103182610b92565b826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561106a57600080fd5b505afa15801561107e573d6000803e3d6000fd5b505050506040513d602081101561109457600080fd5b5051146110d6576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b505b6110e360045461309c565b505060048054600101905550565b6006602052600090815260409020546001600160a01b031681565b3360009081526009602052604090206111248161311e565b600061112e61157e565b90508082600201541161114057600080fd5b6002820154819003838111156111535750825b604051339082156108fc029083906000818181858888f19350505050158015611180573d6000803e3d6000fd5b5050505050565b3332146111c9576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b6000806111f4600333868660405180838380828437604051920182900390912093506131b592505050565b60408051600381523360208201528151939550919350839285927f852c244ccfbd16d5d60ea0e5f658494f8cab6024d58590c2b9f3bae95639b9ab92908290030190a350505050565b6008818154811061124a57fe5b6000918252602090912001546001600160a01b0316905081565b61126c611975565b6112746125a9565b6004546000908152600660205260408120546001600160a01b03169061129990612724565b806001600160a01b0316636cf00e7e6112b183610b92565b600880549050016003546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156112f657600080fd5b505afa15801561130a573d6000803e3d6000fd5b50505050600061138086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808a028281018201909352898252909350899250889182918501908490808284376000920191909152506131f292505050565b905061138c8188613070565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b1580156113c557600080fd5b505afa1580156113d9573d6000803e3d6000fd5b505050506040513d60208110156113ef57600080fd5b505114611432576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b6114a286868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808a028281018201909352898252909350899250889182918501908490808284376000920191909152506132f292505050565b6114ad60035461309c565b60048054600381905560010190556040805188815290517fe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e849181900360200190a150505050505050565b6001600160a01b03811660009081526009602052604090206115188161311e565b60020180543401905550565b6000600a600b548161153257fe5b0490505b90565b600b5481565b6000806000600a848154811061155157fe5b6000918252602090912060029091020180546001909101546001600160a01b039091169350915050915091565b600354600090815260066020908152604080832054815163176fda1560e11b815291516000199385936001600160a01b0390931692632edfb42a9260048083019392829003018186803b1580156115d457600080fd5b505afa1580156115e8573d6000803e3d6000fd5b505050506040513d60208110156115fe57600080fd5b505190504381111561161657600d5492505050611536565b600081430390506000600b54828161162a57fe5b04905060ff81111561163a575060ff5b600019600282900a018061164c575060015b600d54858161165757fe5b0481111561166c578495505050505050611536565b600d540294505050505090565b6001600160a01b03908116600090815260096020526040902060038101546001820154600290920154600160a01b820460ff169492939092911690565b600e546001600160a01b031681565b600f546001600160a01b031681565b600c5481565b33600090815260096020526040902060030154600160a01b900460ff161561173a576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b61174261157e565b341015611789576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b6008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee381018054336001600160a01b031991821681179092556040805160a0810182529384526003805460208087019182523487850190815260006060890181815260808a018b8152988252600990935294909420965187559051968601969096559051600285015593519290930180549151919093166001600160a01b039092169190911760ff60a01b1916600160a01b9115159190910217905543600755565b600a5490565b60035481565b61186a336133bc565b6118bb576040805162461bcd60e51b815260206004820152601a60248201527f6d7573742062652063616c6c656420627920636f6e7472616374000000000000604482015290519081900360640190fd5b61190760053387878787876040516020018086815260200185815260200184815260200183838082843780830192505050955050505050506040516020818303038152906040526133f8565b60405133907f49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f90600090a25050505050565b6009602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b60035460045411801561198c575060055460045411155b6119cd576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b565b6001600160a01b038116600090815260096020526040902060035460018201541115611a2f576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b611a388161311e565b6002810154611a46826134cf565b6040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015611a7c573d6000803e3d6000fd5b50505050565b600d5481565b60055481565b600a54821115611ad6576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b6000600a8381548110611ae557fe5b9060005260206000209060020201905060008160010154905060005b60045482118015611b1157508381105b15611bf557600082815260066020526040808220548554825163025aa7f760e61b81526001600160a01b039182166004820152925191169283926396a9fdc0926024808301939282900301818387803b158015611b6d57600080fd5b505af1158015611b81573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611bbe57600080fd5b505afa158015611bd2573d6000803e3d6000fd5b505050506040513d6020811015611be857600080fd5b5051925050600101611b01565b600454821015611c9d57600a80546000198101908110611c1157fe5b9060005260206000209060020201600a8681548110611c2c57fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600a805480611c6f57fe5b60008281526020812060026000199093019283020180546001600160a01b0319168155600101559055611180565b50600191909101555050565b60075481565b3360009081526009602052604090206003810154600160a01b900460ff16611d0b576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b83834014611d56576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b6004548210158015611d6a57506005548211155b611d7357600080fd5b6000828152600660209081526040918290205482516311e7249560e21b815292516001600160a01b0390911692839263479c925492600480840193829003018186803b158015611dc257600080fd5b505afa158015611dd6573d6000803e3d6000fd5b505050506040513d6020811015611dec57600080fd5b5051600183015414611e37576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b6040805163123334b760e11b815233600482015290516001600160a01b03831691632466696e91602480830192600092919082900301818387803b158015611e7e57600080fd5b505af1158015611e92573d6000803e3d6000fd5b5050505050600101555050565b60015481565b600854606090838301811115611eba57508282015b60608167ffffffffffffffff81118015611ed357600080fd5b50604051908082528060200260200182016040528015611efd578160200160208202803683370190505b50905060005b82811015611f6457600881870181548110611f1a57fe5b9060005260206000200160009054906101000a90046001600160a01b0316828281518110611f4457fe5b6001600160a01b0390921660209283029190910190910152600101611f03565b50949350505050565b604080516001600160a01b038316602082015234818301528151808203830181526060909101909152611fa49060009033906133f8565b50565b6020830135833510611fee576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b60055460208401351115612038576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b600354833511612083576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b8235600090815260066020908152604080832054828701358452928190205481516311e7249560e21b815291516001600160a01b039485169490911692839263479c92549260048083019392829003018186803b1580156120e357600080fd5b505afa1580156120f7573d6000803e3d6000fd5b505050506040513d602081101561210d57600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b15801561214f57600080fd5b505afa158015612163573d6000803e3d6000fd5b505050506040513d602081101561217957600080fd5b5051146121b9576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b6001600160a01b03863581166000908152600960209081526040808320918a013590931682529190206121eb8261311e565b6121f48161311e565b604080516348b4573960e11b81526001600160a01b038a3581166004830152915191861691639168ae7291602480820192602092909190829003018186803b15801561223f57600080fd5b505afa158015612253573d6000803e3d6000fd5b505050506040513d602081101561226957600080fd5b50516122b1576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b604080516348b4573960e11b81526001600160a01b0360208b81013582166004840152925190861692639168ae729260248082019391829003018186803b1580156122fb57600080fd5b505afa15801561230f573d6000803e3d6000fd5b505050506040513d602081101561232557600080fd5b505161236d576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b61238286356020880135604089013588613609565b846001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b1580156123bb57600080fd5b505afa1580156123cf573d6000803e3d6000fd5b505050506040513d60208110156123e557600080fd5b505114612425576040805162461bcd60e51b8152602060048201526009602482015268086908298be9082a6960bb1b604482015290519081900360640190fd5b600f54600b546040805163877c8c2b60e01b8152893560048201526020808b013560248301528a8301356044830152606482018a90526001600160a01b038d35811660848401528d820135811660a484015260c48301949094529151600094939093169263877c8c2b9260e48084019391929182900301818787803b1580156124ad57600080fd5b505af11580156124c1573d6000803e3d6000fd5b505050506040513d60208110156124d757600080fd5b5051600384810180546001600160a01b038085166001600160a01b03199283168117909355928601805490911682179055604080518d35841681526020808f0135909416938101939093528b358382015251929350917fa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879916060908290030190a2505050505050505050565b6125a560033384848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506133f892505050565b5050565b600b54600754430310156119cd576040805162461bcd60e51b815260206004820152600c60248201526b524543454e545f5354414b4560a01b604482015290519081900360640190fd5b60008460601b60601c6001600160a01b0316848484604051602001808581526020018481526020018383808284378083019250505094505050505060405160208183030381529060405280519060200120905061268a8989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508691506136479050565b6000856001600160a01b0316858585604051808383808284376040519201945060009350909150508083038185875af1925050503d80600081146126ea576040519150601f19603f3d011682016040523d82523d6000602084013e6126ef565b606091505b50509050806126fd57600080fd5b50505050505050505050565b60045481565b6010546001600160a01b031681565b60085490565b600a54815b8181101561283a576000600a828154811061274057fe5b906000526020600020906002020190505b6004548160010154101561283157600a600184038154811061276f57fe5b9060005260206000209060020201600a838154811061278a57fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600a8054806127cd57fe5b6000828152602081206000199283016002810290910180546001600160a01b031916815560010191909155909155929092019182821061280f57505050611fa4565b600a828154811061281c57fe5b90600052602060002090600202019050612751565b50600101612729565b505050565b3360009081526009602052604090206003810154600160a01b900460ff1661289b576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b858540146128e6576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b600554600101841461292a576040805162461bcd60e51b81526020600482015260086024820152674e4f44455f4e554d60c01b604482015290519081900360640190fd5b612932613b33565b6040805160e08181019092526129869186906007908390839080828437600092019190915250506040805161014081810190925291508690600a9083908390808284376000920191909152506136e5915050565b600183015460008181526006602090815260409182902054825163380ed4c760e11b8152925194955092936001600160a01b0390931692839263701da98e926004808301939192829003018186803b1580156129e157600080fd5b505afa1580156129f5573d6000803e3d6000fd5b505050506040513d6020811015612a0b57600080fd5b5051612a1684613804565b14612a5a576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b8260800151600154038361012001511115612aad576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b6000816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015612ae857600080fd5b505afa158015612afc573d6000803e3d6000fd5b505050506040513d6020811015612b1257600080fd5b5051845190915043036000612b25611524565b600c54909150820281831015612b6f576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b86608001518760e0015103876101200151101580612b9257508087610140015110155b612bcf576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b806004028761014001511115612c18576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b600b54430184811015612c285750835b6000600c5489610140015181612c3a57fe5b04905080820191506000601060009054906101000a90046001600160a01b03166001600160a01b031663d45ab2b5612c748c60015461383c565b612c848d6001546000548861388b565b612c8d8e613907565b8d886040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b158015612ce157600080fd5b505af1158015612cf5573d6000803e3d6000fd5b505050506040513d6020811015612d0b57600080fd5b5051600580546001019081905560009081526006602052604080822080546001600160a01b0319166001600160a01b038516908117909155815163123334b760e11b8152336004820152915193945092632466696e9260248084019391929182900301818387803b158015612d7f57600080fd5b505af1158015612d93573d6000803e3d6000fd5b505050506005548b600101819055506005547f4807480f255627d9b9350200277cb372949a76ccb4263935665257e2b3a1582a8e8e6001546000546040518085600760200280828437600083820152601f01601f191690910190508461014080828437600083820152601f01601f19169091019384525050602082015260408051918290030192509050a250505050505050505050505050505050565b60005481565b6001600160a01b0380831660009081526009602052604080822084841683529120600382015491929091163314612e6c57600080fd5b60038101546001600160a01b03163314612e8557600080fd5b816002015481600201541115612ee75760028083015490820154604051919003906001600160a01b0385169082156108fc029083906000818181858888f19350505050158015612ed9573d6000803e3d6000fd5b506002820180549190910390555b60028181015483820180549183900490910190556003830180546001600160a01b0319908116909155604080518082019091526001600160a01b03868116825260018086015460208401908152600a80549283018155600052925194027fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a88101805495909216949093169390931790925590517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a990910155611a7c816134cf565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6bffffffffffffffffffffffff191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008181526006602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b1580156130e857600080fd5b505af11580156130fc573d6000803e3d6000fd5b50505060009182525060066020526040902080546001600160a01b0319169055565b6003810154600160a01b900460ff1661316b576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b60038101546001600160a01b031615611fa4576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001546000805490918291826131cf88884342878b613002565b90506131db8282613070565b600055506001828101905590969095509350505050565b80518251600091829182805b838110156132a557600087828151811061321457fe5b60200260200101519050838187011115613264576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868a01810182902060408051808401969096528581019190915280518086038201815260609095019052835193019290922091909401936001016131fe565b508184146132e8576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b9695505050505050565b80516000805b8281101561118057600060ff1685838151811061331157fe5b016020015160f81c141561339757600061332e866001850161391d565b905060028160405161333f90613bbf565b90815260405190819003602001906000f080158015613362573d6000803e3d6000fd5b5081546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b03909216919091179055505b8381815181106133a357fe5b60200260200101518201915080806001019150506132f8565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906133f057508115155b949350505050565b60008061340d858585805190602001206131b5565b9150915080827ffc06a498d72efc51848331933699060ef69722b3ffaae0a25c549c461d48c003878787604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561348c578181015183820152602001613474565b50505050905090810190601f1680156134b95780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050565b80546008805460009190839081106134e357fe5b600091825260209091200154600880546001600160a01b03909216925090600019810190811061350f57fe5b600091825260209091200154600880546001600160a01b03909216918490811061353557fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600960006008858154811061357557fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205560088054806135a557fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0392909216815260099091526040812081815560018101829055600281019190915560030180546001600160a81b03191690555050565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b600160001b811890506000613660848385600101613976565b5090506002858154811061367057fe5b600091825260208220015460408051630ad0379b60e01b8152600481018590526024810187905290516001600160a01b0390921692630ad0379b9260448084019382900301818387803b1580156136c657600080fd5b505af11580156136da573d6000803e3d6000fd5b505050505050505050565b6136ed613b33565b60408051610220810182528351815260208085015181830152855182840152850151606080830191909152848301516080808401919091529085015160a0808401919091529085015160c083015284015160e082015290840151610100820152610120810183600660200201518152602001836007600a811061376c57fe5b602002015181526020018460036007811061378357fe5b60200201518152602001836008600a811061379a57fe5b60200201518152602001846004600781106137b157fe5b60200201518152602001836009600a81106137c857fe5b60200201518152602001846005600781106137df57fe5b60200201518152602001846006600781106137f657fe5b602002015190529392505050565b6000613836826000015183602001518460400151856060015186608001518760a001518860c001518960e00151612fa8565b92915050565b600061388443846101400151856020015101856102000151866101e001518761012001518860800151018861018001518960a0015101896101c001518a60c001510189612fa8565b9392505050565b6000806138ad6000876101200151886080015188030386896101e00151613609565b905060006138e660008861012001516138ce8a6101e001516000801b613070565b6138e18b606001518c6101000151613070565b613609565b90506138fc82826138f68a613a83565b87613609565b979650505050505050565b6000613836826101600151836101a00151613070565b6000816020018351101561396d576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60008080848160205b88518111613a75578089015193506020818a51036020018161399d57fe5b0491505b6000821180156139b45750600287066001145b80156139c257508160020a87115b156139da5760029096046001908101969401936139a1565b60028706613a25578383604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120925060028781613a1d57fe5b049650613a67565b8284604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120925060028781613a6057fe5b0460010196505b60019094019360200161397f565b509093505050935093915050565b60006138366000836101400151613ab96000613ab487610100015188604001516000801b60008060001b6000613ae8565b613070565b6138e1866101400151613ab46000801b8961020001518a61016001518b61018001518c6101a001518d6101c001515b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e0810182905261020081019190915290565b61013780613bcd8339019056fe608060405234801561001057600080fd5b506040516101373803806101378339818101604052602081101561003357600080fd5b5051600080546001600160a01b0319163317905560015560df806100586000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630ad0379b14602d575b600080fd5b604d60048036036040811015604157600080fd5b5080359060200135604f565b005b6000546001600160a01b03163314606557600080fd5b60008181526002602052604090205460ff1615608057600080fd5b6001548214608d57600080fd5b6000908152600260205260409020805460ff191660011790555056fea264697066735822122045771b255caea9c72008e9382dc6c60996b94002366693a4d85d84c26f81042064736f6c634300060c0033a2646970667358221220bc9429f0e21e7ed0d8801cdd78484890208a8cbf3850b33364561d2acf4436f664736f6c634300060c0033a26469706673582212202aad699ba2b877594d7906766732177f843e401c46956e80511d6ba75b3589c664736f6c634300060c0033"

// DeployRollupCreator deploys a new Ethereum contract, binding an instance of RollupCreator to it.
func DeployRollupCreator(auth *bind.TransactOpts, backend bind.ContractBackend, _challengeFactory common.Address, _nodeFactory common.Address) (common.Address, *types.Transaction, *RollupCreator, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCreatorBin), backend, _challengeFactory, _nodeFactory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// RollupCreator is an auto generated Go binding around an Ethereum contract.
type RollupCreator struct {
	RollupCreatorCaller     // Read-only binding to the contract
	RollupCreatorTransactor // Write-only binding to the contract
	RollupCreatorFilterer   // Log filterer for contract events
}

// RollupCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCreatorSession struct {
	Contract     *RollupCreator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCreatorCallerSession struct {
	Contract *RollupCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RollupCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCreatorTransactorSession struct {
	Contract     *RollupCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RollupCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCreatorRaw struct {
	Contract *RollupCreator // Generic contract binding to access the raw methods on
}

// RollupCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCreatorCallerRaw struct {
	Contract *RollupCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCreatorTransactorRaw struct {
	Contract *RollupCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCreator creates a new instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreator(address common.Address, backend bind.ContractBackend) (*RollupCreator, error) {
	contract, err := bindRollupCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// NewRollupCreatorCaller creates a new read-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorCaller(address common.Address, caller bind.ContractCaller) (*RollupCreatorCaller, error) {
	contract, err := bindRollupCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorCaller{contract: contract}, nil
}

// NewRollupCreatorTransactor creates a new write-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCreatorTransactor, error) {
	contract, err := bindRollupCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorTransactor{contract: contract}, nil
}

// NewRollupCreatorFilterer creates a new log filterer instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCreatorFilterer, error) {
	contract, err := bindRollupCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorFilterer{contract: contract}, nil
}

// bindRollupCreator binds a generic wrapper to an already deployed contract.
func bindRollupCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.RollupCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transact(opts, method, params...)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xd798f5be.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorTransactor) CreateRollup(opts *bind.TransactOpts, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "createRollup", _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xd798f5be.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorSession) CreateRollup(_machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xd798f5be.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorTransactorSession) CreateRollup(_machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// RollupCreatorRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the RollupCreator contract.
type RollupCreatorRollupCreatedIterator struct {
	Event *RollupCreatorRollupCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupCreatorRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorRollupCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupCreatorRollupCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupCreatorRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorRollupCreated represents a RollupCreated event raised by the RollupCreator contract.
type RollupCreatorRollupCreated struct {
	RollupAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*RollupCreatorRollupCreatedIterator, error) {

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &RollupCreatorRollupCreatedIterator{contract: _RollupCreator.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupCreatorRollupCreated) (event.Subscription, error) {

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorRollupCreated)
				if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupCreated is a log parse operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) ParseRollupCreated(log types.Log) (*RollupCreatorRollupCreated, error) {
	event := new(RollupCreatorRollupCreated)
	if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupLibABI is the input ABI used to generate the binding from.
const RollupLibABI = "[]"

// RollupLibBin is the compiled bytecode used for deploying new contracts.
var RollupLibBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cc9f23d5349197405d1242d19e74630c294791596f8f0f4596866cd0d48f416e64736f6c634300060c0033"

// DeployRollupLib deploys a new Ethereum contract, binding an instance of RollupLib to it.
func DeployRollupLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupLib, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupLib{RollupLibCaller: RollupLibCaller{contract: contract}, RollupLibTransactor: RollupLibTransactor{contract: contract}, RollupLibFilterer: RollupLibFilterer{contract: contract}}, nil
}

// RollupLib is an auto generated Go binding around an Ethereum contract.
type RollupLib struct {
	RollupLibCaller     // Read-only binding to the contract
	RollupLibTransactor // Write-only binding to the contract
	RollupLibFilterer   // Log filterer for contract events
}

// RollupLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupLibSession struct {
	Contract     *RollupLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupLibCallerSession struct {
	Contract *RollupLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RollupLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupLibTransactorSession struct {
	Contract     *RollupLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RollupLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupLibRaw struct {
	Contract *RollupLib // Generic contract binding to access the raw methods on
}

// RollupLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupLibCallerRaw struct {
	Contract *RollupLibCaller // Generic read-only contract binding to access the raw methods on
}

// RollupLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupLibTransactorRaw struct {
	Contract *RollupLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupLib creates a new instance of RollupLib, bound to a specific deployed contract.
func NewRollupLib(address common.Address, backend bind.ContractBackend) (*RollupLib, error) {
	contract, err := bindRollupLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupLib{RollupLibCaller: RollupLibCaller{contract: contract}, RollupLibTransactor: RollupLibTransactor{contract: contract}, RollupLibFilterer: RollupLibFilterer{contract: contract}}, nil
}

// NewRollupLibCaller creates a new read-only instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibCaller(address common.Address, caller bind.ContractCaller) (*RollupLibCaller, error) {
	contract, err := bindRollupLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupLibCaller{contract: contract}, nil
}

// NewRollupLibTransactor creates a new write-only instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupLibTransactor, error) {
	contract, err := bindRollupLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupLibTransactor{contract: contract}, nil
}

// NewRollupLibFilterer creates a new log filterer instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupLibFilterer, error) {
	contract, err := bindRollupLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupLibFilterer{contract: contract}, nil
}

// bindRollupLib binds a generic wrapper to an already deployed contract.
func bindRollupLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupLib *RollupLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupLib.Contract.RollupLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupLib *RollupLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupLib.Contract.RollupLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupLib *RollupLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupLib.Contract.RollupLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupLib *RollupLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupLib *RollupLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupLib *RollupLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupLib.Contract.contract.Transact(opts, method, params...)
}
