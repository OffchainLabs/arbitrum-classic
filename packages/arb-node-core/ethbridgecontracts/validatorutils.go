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
const INodeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"addStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkPastDeadline\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"latestConfirmed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"checkRejectExample\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"removeStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// CheckPastDeadline is a free data retrieval call binding the contract method 0x0cad16a5.
//
// Solidity: function checkPastDeadline() view returns()
func (_INode *INodeCaller) CheckPastDeadline(opts *bind.CallOpts) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "checkPastDeadline")

	if err != nil {
		return err
	}

	return err

}

// CheckPastDeadline is a free data retrieval call binding the contract method 0x0cad16a5.
//
// Solidity: function checkPastDeadline() view returns()
func (_INode *INodeSession) CheckPastDeadline() error {
	return _INode.Contract.CheckPastDeadline(&_INode.CallOpts)
}

// CheckPastDeadline is a free data retrieval call binding the contract method 0x0cad16a5.
//
// Solidity: function checkPastDeadline() view returns()
func (_INode *INodeCallerSession) CheckPastDeadline() error {
	return _INode.Contract.CheckPastDeadline(&_INode.CallOpts)
}

// CheckRejectExample is a free data retrieval call binding the contract method 0x0564921f.
//
// Solidity: function checkRejectExample(uint256 latestConfirmed, address stakerAddress) view returns()
func (_INode *INodeCaller) CheckRejectExample(opts *bind.CallOpts, latestConfirmed *big.Int, stakerAddress common.Address) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "checkRejectExample", latestConfirmed, stakerAddress)

	if err != nil {
		return err
	}

	return err

}

// CheckRejectExample is a free data retrieval call binding the contract method 0x0564921f.
//
// Solidity: function checkRejectExample(uint256 latestConfirmed, address stakerAddress) view returns()
func (_INode *INodeSession) CheckRejectExample(latestConfirmed *big.Int, stakerAddress common.Address) error {
	return _INode.Contract.CheckRejectExample(&_INode.CallOpts, latestConfirmed, stakerAddress)
}

// CheckRejectExample is a free data retrieval call binding the contract method 0x0564921f.
//
// Solidity: function checkRejectExample(uint256 latestConfirmed, address stakerAddress) view returns()
func (_INode *INodeCallerSession) CheckRejectExample(latestConfirmed *big.Int, stakerAddress common.Address) error {
	return _INode.Contract.CheckRejectExample(&_INode.CallOpts, latestConfirmed, stakerAddress)
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

// ProxyAdminABI is the input ABI used to generate the binding from.
const ProxyAdminABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ProxyAdminBin is the compiled bytecode used for deploying new contracts.
var ProxyAdminBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6108338061007d6000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100bf5780637eff275e146100d65780638da5cb5b146101115780639623609d1461012657806399a88ec4146101e3578063f2fde38b1461021e578063f3b7dead14610251575b600080fd5b34801561007c57600080fd5b506100a36004803603602081101561009357600080fd5b50356001600160a01b0316610284565b604080516001600160a01b039092168252519081900360200190f35b3480156100cb57600080fd5b506100d4610316565b005b3480156100e257600080fd5b506100d4600480360360408110156100f957600080fd5b506001600160a01b03813581169160200135166103a6565b34801561011d57600080fd5b506100a3610469565b6100d46004803603606081101561013c57600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460018302840111600160201b831117156101a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610478945050505050565b3480156101ef57600080fd5b506100d46004803603604081101561020657600080fd5b506001600160a01b03813581169160200135166105a7565b34801561022a57600080fd5b506100d46004803603602081101561024157600080fd5b50356001600160a01b031661064e565b34801561025d57600080fd5b506100a36004803603602081101561027457600080fd5b50356001600160a01b0316610734565b6000806060836001600160a01b03166040518080635c60da1b60e01b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b606091505b5091509150816102f757600080fd5b80806020019051602081101561030c57600080fd5b5051949350505050565b61031e610793565b6000546001600160a01b0390811691161461036e576040805162461bcd60e51b815260206004820181905260248201526000805160206107be833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116906000805160206107de833981519152908390a3600080546001600160a01b0319169055565b6103ae610793565b6000546001600160a01b039081169116146103fe576040805162461bcd60e51b815260206004820181905260248201526000805160206107be833981519152604482015290519081900360640190fd5b816001600160a01b0316638f283970826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561044d57600080fd5b505af1158015610461573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b610480610793565b6000546001600160a01b039081169116146104d0576040805162461bcd60e51b815260206004820181905260248201526000805160206107be833981519152604482015290519081900360640190fd5b826001600160a01b0316634f1ef2863484846040518463ffffffff1660e01b815260040180836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561053d578181015183820152602001610525565b50505050905090810190601f16801561056a5780820380516001836020036101000a031916815260200191505b5093505050506000604051808303818588803b15801561058957600080fd5b505af115801561059d573d6000803e3d6000fd5b5050505050505050565b6105af610793565b6000546001600160a01b039081169116146105ff576040805162461bcd60e51b815260206004820181905260248201526000805160206107be833981519152604482015290519081900360640190fd5b816001600160a01b0316633659cfe6826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561044d57600080fd5b610656610793565b6000546001600160a01b039081169116146106a6576040805162461bcd60e51b815260206004820181905260248201526000805160206107be833981519152604482015290519081900360640190fd5b6001600160a01b0381166106eb5760405162461bcd60e51b81526004018080602001828103825260268152602001806107986026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216916000805160206107de83398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806060836001600160a01b031660405180806303e1469160e61b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a26469706673582212207065f73260f6991242aacef0846fca2c74e0bd794508d746ea8bcbf1a6a4cb7064736f6c634300060c0033"

// DeployProxyAdmin deploys a new Ethereum contract, binding an instance of ProxyAdmin to it.
func DeployProxyAdmin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProxyAdmin, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAdminABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyAdminBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// ProxyAdmin is an auto generated Go binding around an Ethereum contract.
type ProxyAdmin struct {
	ProxyAdminCaller     // Read-only binding to the contract
	ProxyAdminTransactor // Write-only binding to the contract
	ProxyAdminFilterer   // Log filterer for contract events
}

// ProxyAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyAdminSession struct {
	Contract     *ProxyAdmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyAdminCallerSession struct {
	Contract *ProxyAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProxyAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyAdminTransactorSession struct {
	Contract     *ProxyAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProxyAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyAdminRaw struct {
	Contract *ProxyAdmin // Generic contract binding to access the raw methods on
}

// ProxyAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyAdminCallerRaw struct {
	Contract *ProxyAdminCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyAdminTransactorRaw struct {
	Contract *ProxyAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyAdmin creates a new instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdmin(address common.Address, backend bind.ContractBackend) (*ProxyAdmin, error) {
	contract, err := bindProxyAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// NewProxyAdminCaller creates a new read-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminCaller(address common.Address, caller bind.ContractCaller) (*ProxyAdminCaller, error) {
	contract, err := bindProxyAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminCaller{contract: contract}, nil
}

// NewProxyAdminTransactor creates a new write-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyAdminTransactor, error) {
	contract, err := bindProxyAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminTransactor{contract: contract}, nil
}

// NewProxyAdminFilterer creates a new log filterer instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyAdminFilterer, error) {
	contract, err := bindProxyAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminFilterer{contract: contract}, nil
}

// bindProxyAdmin binds a generic wrapper to an already deployed contract.
func bindProxyAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.ProxyAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transact(opts, method, params...)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyAdmin(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyAdmin", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyImplementation(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyImplementation", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, proxy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "changeProxyAdmin", proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, proxy, newAdmin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactor) Upgrade(opts *bind.TransactOpts, proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgrade", proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, proxy, implementation)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactor) UpgradeAndCall(opts *bind.TransactOpts, proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgradeAndCall", proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, proxy, implementation, data)
}

// ProxyAdminOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferredIterator struct {
	Event *ProxyAdminOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProxyAdminOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAdminOwnershipTransferred)
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
		it.Event = new(ProxyAdminOwnershipTransferred)
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
func (it *ProxyAdminOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAdminOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAdminOwnershipTransferred represents a OwnershipTransferred event raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProxyAdminOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminOwnershipTransferredIterator{contract: _ProxyAdmin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyAdminOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAdminOwnershipTransferred)
				if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyAdminOwnershipTransferred, error) {
	event := new(ProxyAdminOwnershipTransferred)
	if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"indexed\":false,\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"inboxMaxHash\",\"type\":\"bytes32\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"SentLogs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"addToDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkNoRecentStake\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"logAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[3]\",\"name\":\"nodeFields\",\"type\":\"bytes32[3]\"},{\"internalType\":\"uint256\",\"name\":\"executionCheckTime\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"newStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"successorWithStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_latestNodeCreated\",\"type\":\"uint256\"}],\"name\":\"truncateNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"}],\"name\":\"upgradeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeImplementationAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupBin is the compiled bytecode used for deploying new contracts.
var RollupBin = "0x608060405234801561001057600080fd5b506009805460ff1916905561474d8061002a6000396000f3fe60806040526004361061025a5760003560e01c8063046f7da21461025f57806304a28064146102765780630e1ef04c146102bb57806315edeb3a146102f45780631e83d30f1461031e5780632f30cabd14610348578063396b8cbc1461037b5780633e03cf431461044d5780633e96576e1461055e57806345e38b641461059157806346c2781a146105a65780634d26732d146105bb5780634f0f4aa9146105d057806351ed6a3014610616578063567ca41b1461062b5780635c975abb1461065e5780635dbaf68b146106875780635e8ef1061461069c5780636177fd18146106b157806362a82d7d146106e457806363721d6b1461070e57806365f7f80d1461072357806369fd251c1461073857806373f33b061461076b5780637427be511461078057806376e7e23b146107b35780637ba9534a146107c85780637e2d2155146107dd57806381fbc98a1461080d57806383f94db7146108405780638456cb59146108735780638640ce5f146108885780638da5cb5b1461089d5780638fd18f04146108b2578063b1b4181c146108e8578063be211c9a1461091f578063bf5ddcb114610934578063ce11e6ab146109bf578063d01e6602146109d4578063d735e21d146109fe578063d93fe9c414610a13578063dff6978714610a28578063e45b7ce614610a3d578063e78cea9214610a78578063e8bd492214610a8d578063edfd03ed14610af6578063ef40a67014610b20578063f019a1c114610b53578063f33e1fac14610b95578063f3f0a03e14610bbf578063fa7803e614610beb578063fb64884e14610c26578063ff204f3b14610c43575b600080fd5b34801561026b57600080fd5b50610274610c76565b005b34801561028257600080fd5b506102a96004803603602081101561029957600080fd5b50356001600160a01b0316610ccc565b60408051918252519081900360200190f35b3480156102c757600080fd5b50610274600480360360408110156102de57600080fd5b50803590602001356001600160a01b0316610d84565b34801561030057600080fd5b506102746004803603602081101561031757600080fd5b50356110d6565b34801561032a57600080fd5b506102746004803603602081101561034157600080fd5b503561123a565b34801561035457600080fd5b506102a96004803603602081101561036b57600080fd5b50356001600160a01b03166112b0565b34801561038757600080fd5b506102746004803603606081101561039e57600080fd5b81359190810190604081016020820135600160201b8111156103bf57600080fd5b8201836020820111156103d157600080fd5b803590602001918460018302840111600160201b831117156103f257600080fd5b919390929091602081019035600160201b81111561040f57600080fd5b82018360208201111561042157600080fd5b803590602001918460208302840111600160201b8311171561044257600080fd5b5090925090506112cb565b34801561045957600080fd5b50610274600480360361018081101561047157600080fd5b6001600160a01b03823581169260208101359260408201359260608301359260808101359260a082013583169260c083013581169260e08101358216926101008201358316926101208301351691908101906101608101610140820135600160201b8111156104df57600080fd5b8201836020820111156104f157600080fd5b803590602001918460018302840111600160201b8311171561051257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506117719050565b34801561056a57600080fd5b506102a96004803603602081101561058157600080fd5b50356001600160a01b0316611acb565b34801561059d57600080fd5b506102a9611ae9565b3480156105b257600080fd5b506102a9611afd565b3480156105c757600080fd5b506102a9611b03565b3480156105dc57600080fd5b506105fa600480360360208110156105f357600080fd5b5035611bf6565b604080516001600160a01b039092168252519081900360200190f35b34801561062257600080fd5b506105fa611c11565b34801561063757600080fd5b506102746004803603602081101561064e57600080fd5b50356001600160a01b0316611c20565b34801561066a57600080fd5b50610673611d23565b604080519115158252519081900360200190f35b34801561069357600080fd5b506105fa611d2c565b3480156106a857600080fd5b506102a9611d3b565b3480156106bd57600080fd5b50610673600480360360208110156106d457600080fd5b50356001600160a01b0316611d41565b3480156106f057600080fd5b506105fa6004803603602081101561070757600080fd5b5035611d69565b34801561071a57600080fd5b506102a9611d93565b34801561072f57600080fd5b506102a9611d99565b34801561074457600080fd5b506105fa6004803603602081101561075b57600080fd5b50356001600160a01b0316611d9f565b34801561077757600080fd5b50610274611dc0565b34801561078c57600080fd5b50610274600480360360208110156107a357600080fd5b50356001600160a01b0316611e2d565b3480156107bf57600080fd5b506102a9611ed6565b3480156107d457600080fd5b506102a9611edc565b3480156107e957600080fd5b506102746004803603604081101561080057600080fd5b5080359060200135611ee2565b34801561081957600080fd5b506102a96004803603602081101561083057600080fd5b50356001600160a01b03166120bf565b34801561084c57600080fd5b506102746004803603602081101561086357600080fd5b50356001600160a01b0316612125565b34801561087f57600080fd5b506102746121de565b34801561089457600080fd5b506102a9612232565b3480156108a957600080fd5b506105fa612238565b3480156108be57600080fd5b50610274600480360360608110156108d557600080fd5b5080359060208101359060400135612247565b3480156108f457600080fd5b50610274600480360361010081101561090c57600080fd5b50604081016080820160e0830135612415565b34801561092b57600080fd5b50610274612a09565b34801561094057600080fd5b506102746004803603604081101561095757600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561098157600080fd5b82018360208201111561099357600080fd5b803590602001918460018302840111600160201b831117156109b457600080fd5b509092509050612a58565b3480156109cb57600080fd5b506105fa612b4e565b3480156109e057600080fd5b506105fa600480360360208110156109f757600080fd5b5035612b5d565b348015610a0a57600080fd5b506102a9612b8c565b348015610a1f57600080fd5b506105fa612b92565b348015610a3457600080fd5b506102a9612ba1565b348015610a4957600080fd5b5061027460048036036040811015610a6057600080fd5b506001600160a01b0381351690602001351515612ba7565b348015610a8457600080fd5b506105fa612c49565b348015610a9957600080fd5b50610ac060048036036020811015610ab057600080fd5b50356001600160a01b0316612c58565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b348015610b0257600080fd5b5061027460048036036020811015610b1957600080fd5b5035612c94565b348015610b2c57600080fd5b506102a960048036036020811015610b4357600080fd5b50356001600160a01b0316612cf3565b348015610b5f57600080fd5b506102746004803603610280811015610b7757600080fd5b50803590602081013590604081013590606081019061014001612d11565b348015610ba157600080fd5b506102a960048036036020811015610bb857600080fd5b503561331a565b61027460048036036040811015610bd557600080fd5b506001600160a01b038135169060200135613342565b348015610bf757600080fd5b5061027460048036036040811015610c0e57600080fd5b506001600160a01b03813581169160200135166133a3565b61027460048036036020811015610c3c57600080fd5b5035613457565b348015610c4f57600080fd5b5061027460048036036020811015610c6657600080fd5b50356001600160a01b031661354f565b6012546001600160a01b03163314610cc2576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610cca613605565b565b600080610cd7611d93565b90506000805b82811015610d7a57846001600160a01b0316639168ae72610cfd83612b5d565b6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610d3a57600080fd5b505afa158015610d4e573d6000803e3d6000fd5b505050506040513d6020811015610d6457600080fd5b505115610d72576001909101905b600101610cdd565b509150505b919050565b60095460ff1615610dca576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b610dd2611dc0565b6000610ddc611d99565b90506000610de8612b8c565b90506000610df582611bf6565b905082816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610e3157600080fd5b505afa158015610e45573d6000803e3d6000fd5b505050506040513d6020811015610e5b57600080fd5b505114156110c757610e6b612a09565b818511610eb2576040805162461bcd60e51b815260206004820152601060248201526f535543434553534f525f544f5f4c4f5760801b604482015290519081900360640190fd5b610eba611edc565b851115610f02576040805162461bcd60e51b81526020600482015260116024820152700a6aa86868aa6a69ea4bea89ebe90928e9607b1b604482015290519081900360640190fd5b610f0b84611d41565b610f49576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b610f5285611bf6565b6001600160a01b0316630564921f84866040518363ffffffff1660e01b815260040180838152602001826001600160a01b031681526020019250505060006040518083038186803b158015610fa657600080fd5b505afa158015610fba573d6000803e3d6000fd5b50505050806001600160a01b0316630cad16a56040518163ffffffff1660e01b815260040160006040518083038186803b158015610ff757600080fd5b505afa15801561100b573d6000803e3d6000fd5b505050506110196000612c94565b61102281610ccc565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561105b57600080fd5b505afa15801561106f573d6000803e3d6000fd5b505050506040513d602081101561108557600080fd5b5051146110c7576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b6110cf6136a3565b5050505050565b6012546001600160a01b03163314611122576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b600061112c611edc565b905080821061116c576040805162461bcd60e51b8152602060048201526007602482015266544f4f5f4e455760c81b604482015290519081900360640190fd5b6001611176612b8c565b038210156111b5576040805162461bcd60e51b81526020600482015260076024820152661513d3d7d3d31160ca1b604482015290519081900360640190fd5b805b82811061122c5760006111c982611bf6565b9050806001600160a01b03166383197ef06040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561120657600080fd5b505af115801561121a573d6000803e3d6000fd5b505060001990930192506111b7915050565b50611236826136b8565b5050565b60095460ff1615611280576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b611289336136bd565b6000611293611b03565b9050808211156112a1578091505b6112ab3383613754565b505050565b6001600160a01b031660009081526008602052604090205490565b60095460ff1615611311576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b611319611dc0565b611321612a09565b600061132b612ba1565b1161136a576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f5354414b45525360b01b604482015290519081900360640190fd5b6000611374612b8c565b9050600061138182611bf6565b9050806001600160a01b0316630cad16a56040518163ffffffff1660e01b815260040160006040518083038186803b1580156113bc57600080fd5b505afa1580156113d0573d6000803e3d6000fd5b505050506113dc611d99565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561141557600080fd5b505afa158015611429573d6000803e3d6000fd5b505050506040513d602081101561143f57600080fd5b505114611482576040805162461bcd60e51b815260206004820152600c60248201526b24a72b20a624a22fa82922ab60a11b604482015290519081900360640190fd5b61148c6000612c94565b61149581610ccc565b61149d612ba1565b01816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b1580156114d757600080fd5b505afa1580156114eb573d6000803e3d6000fd5b505050506040513d602081101561150157600080fd5b505114611546576040805162461bcd60e51b815260206004820152600e60248201526d1393d517d0531317d4d51052d15160921b604482015290519081900360640190fd5b60006115b887878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808b0282810182019093528a82529093508a9250899182918501908490808284376000920191909152506137ec92505050565b90506115c481896138ec565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b1580156115fd57600080fd5b505afa158015611611573d6000803e3d6000fd5b505050506040513d602081101561162757600080fd5b50511461166a576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b600f5460408051630c72684760e01b815260048101918252604481018990526001600160a01b0390921691630c726847918a918a918a918a919081906024810190606401878780828437600083820152601f01601f19169091018481038352858152602090810191508690860280828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b15801561171457600080fd5b505af1158015611728573d6000803e3d6000fd5b50505050611734613918565b6040805189815290517fe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e849181900360200190a15050505050505050565b600e80546001600160a01b0319166001600160a01b0387811691909117918290556040805163722dbe7360e11b8152306004820152600160248201529051929091169163e45b7ce69160448082019260009290919082900301818387803b1580156117db57600080fd5b505af11580156117ef573d6000803e3d6000fd5b505050508b600f60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060008a8a8a8a60601b6001600160601b0319168a60601b6001600160601b031916876040516020018087815260200186815260200185815260200184815260200183815260200182805190602001908083835b6020831061188c5780518252601f19909201916020918201910161186d565b6001836020036101000a0380198251168184511680821785525050505050509050019650505050505050604051602081830303815290604052805190602001209050600e60009054906101000a90046001600160a01b03166001600160a01b031663636f807d6119026004304342600088613932565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561193857600080fd5b505af115801561194c573d6000803e3d6000fd5b5050601080546001600160a01b03808a166001600160a01b0319928316179092556011805492891692909116919091179055506000905061199443828f81808080600161399b565b6011546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905260848201819052915193945090926001600160a01b039092169163d45ab2b59160a48082019260209290919082900301818787803b158015611a0457600080fd5b505af1158015611a18573d6000803e3d6000fd5b505050506040513d6020811015611a2e57600080fd5b50519050611a3b816139f5565b600a8d9055600b8c9055600c8b9055600d80546001600160a01b03808d166001600160a01b031992831617909255601280548c84169083161790556013805492871692909116919091179055604080518f815290517f4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d9181900360200190a1505050505050505050505050505050565b6001600160a01b031660009081526006602052604090206001015490565b6000600a805481611af657fe5b0490505b90565b600a5481565b600060001981611b19611b14611d99565b611bf6565b6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611b5157600080fd5b505afa158015611b65573d6000803e3d6000fd5b505050506040513d6020811015611b7b57600080fd5b5051905043811115611b9357600c5492505050611afa565b600081430390506000600a548281611ba757fe5b04905060ff811115611bb7575060ff5b600019600282900a0180611bc9575060015b600c548581611bd457fe5b04811115611be9578495505050505050611afa565b600c540294505050505090565b6000908152600460205260409020546001600160a01b031690565b600d546001600160a01b031681565b6012546001600160a01b03163314611c6c576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b600f546001600160a01b0382811691161415611cbc576040805162461bcd60e51b815260206004820152600a602482015269086aaa4be9eaaa8849eb60b31b604482015290519081900360640190fd5b600e54604080516319dc7ae560e31b81526001600160a01b038481166004830152600060248301819052925193169263cee3d7289260448084019391929182900301818387803b158015611d0f57600080fd5b505af11580156110cf573d6000803e3d6000fd5b60095460ff1690565b6010546001600160a01b031681565b600b5481565b6001600160a01b0316600090815260066020526040902060030154600160a01b900460ff1690565b600060058281548110611d7857fe5b6000918252602090912001546001600160a01b031692915050565b60075490565b60005490565b6001600160a01b039081166000908152600660205260409020600301541690565b6000611dca612b8c565b9050611dd4611d99565b81118015611de95750611de5611edc565b8111155b611e2a576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b50565b60095460ff1615611e73576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b611e7b611d99565b611e8482611acb565b1115611ec4576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b611ecd816136bd565b611e2a81613a43565b600c5481565b60025490565b60095460ff1615611f28576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b611f30611d93565b821115611f75576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b6000611f8083612b5d565b90506000611f8d8461331a565b9050600080611f9a612b8c565b90505b8083118015611fab57508482105b15612097576000611fbb84611bf6565b9050806001600160a01b03166396a9fdc0866040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561200c57600080fd5b505af1158015612020573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561205d57600080fd5b505afa158015612071573d6000803e3d6000fd5b505050506040513d602081101561208757600080fd5b5051935050600190910190611f9d565b808310156120ad576120a886613a7b565b6120b7565b6120b78684613b17565b505050505050565b60095460009060ff1615612108576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b600061211333613b3e565b905061211f8382613b5d565b92915050565b6012546001600160a01b03163314612171576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6013546040805163266a23b160e21b815230600482018190526001600160a01b0385811660248401529251909392909216916399a88ec49160448082019260009290919082900301818387803b1580156121ca57600080fd5b505af11580156120b7573d6000803e3d6000fd5b6012546001600160a01b0316331461222a576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610cca613c78565b60035490565b6012546001600160a01b031681565b60095460ff161561228d576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b61229633611d41565b6122d4576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b8282401461231f576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b612327612b8c565b811015801561233d5750612339611edc565b8111155b61234657600080fd5b600061235182611bf6565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561238c57600080fd5b505afa1580156123a0573d6000803e3d6000fd5b505050506040513d60208110156123b657600080fd5b50516123c133611acb565b14612405576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b61240f3383613cf4565b50505050565b60095460ff161561245b576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b60208301358335106124a2576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b6124aa611edc565b602084013511156124f1576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b82356124fb611d99565b10612541576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b600061255384825b6020020135611bf6565b90506000612562856001612549565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561259d57600080fd5b505afa1580156125b1573d6000803e3d6000fd5b505050506040513d60208110156125c757600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b15801561260957600080fd5b505afa15801561261d573d6000803e3d6000fd5b505050506040513d602081101561263357600080fd5b505114612673576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b61268d8660005b60200201356001600160a01b03166136bd565b61269886600161267a565b604080516348b4573960e11b81526001600160a01b03883581166004830152915191841691639168ae7291602480820192602092909190829003018186803b1580156126e357600080fd5b505afa1580156126f7573d6000803e3d6000fd5b505050506040513d602081101561270d57600080fd5b5051612755576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b604080516348b4573960e11b81526001600160a01b0360208981013582166004840152925190841692639168ae729260248082019391829003018186803b15801561279f57600080fd5b505afa1580156127b3573d6000803e3d6000fd5b505050506040513d60208110156127c957600080fd5b5051612811576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b61282684356020860135604087013586613d7d565b826001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b15801561285f57600080fd5b505afa158015612873573d6000803e3d6000fd5b505050506040513d602081101561288957600080fd5b5051146128c9576040805162461bcd60e51b8152602060048201526009602482015268086908298be9082a6960bb1b604482015290519081900360640190fd5b601054600a5460408051636fcbc15160e11b8152306004820152873560248201526020808901356044830152888301356064830152608482018890526001600160a01b038b35811660a48401528b820135811660c484015260e48301949094529151600094939093169263df9782a2926101048084019391929182900301818787803b15801561295857600080fd5b505af115801561296c573d6000803e3d6000fd5b505050506040513d602081101561298257600080fd5b505190506129ab6001600160a01b0388351688600160200201356001600160a01b031683613dbb565b6040805188356001600160a01b0390811682526020808b01358216908301528835828401529151918316917fa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda758799181900360600190a250505050505050565b600a54612a14612232565b43031015610cca576040805162461bcd60e51b815260206004820152600c60248201526b524543454e545f5354414b4560a01b604482015290519081900360640190fd5b6012546001600160a01b03163314612aa4576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601354604051639623609d60e01b815230600482018181526001600160a01b0387811660248501526060604485019081526064850187905292941692639623609d92859289928992899291608401848480828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b158015612b3057600080fd5b505af1158015612b44573d6000803e3d6000fd5b5050505050505050565b600f546001600160a01b031681565b600060078281548110612b6c57fe5b60009182526020909120600290910201546001600160a01b031692915050565b60015490565b6011546001600160a01b031681565b60055490565b6012546001600160a01b03163314612bf3576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b600e546040805163722dbe7360e11b81526001600160a01b03858116600483015284151560248301529151919092169163e45b7ce691604480830192600092919082900301818387803b1580156121ca57600080fd5b600e546001600160a01b031681565b6006602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6000612c9e611d93565b90506000612caa612b8c565b9050825b8281101561240f575b81612cc18261331a565b1015612ceb57612cd081613a7b565b60001990920191828110612ce657505050611e2a565b612cb7565b600101612cae565b6001600160a01b031660009081526006602052604090206002015490565b60095460ff1615612d57576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b612d6033611d41565b612d9e576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b84844014612de9576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b612df1611edc565b6001018314612e32576040805162461bcd60e51b81526020600482015260086024820152674e4f44455f4e554d60c01b604482015290519081900360640190fd5b612e3a61466b565b6040805160e0818101909252612e8e9185906007908390839080828437600092019190915250506040805161014081810190925291508590600a908390839080828437600092019190915250613e05915050565b90506000612e9e611b1433611acb565b9050806001600160a01b031663701da98e6040518163ffffffff1660e01b815260040160206040518083038186803b158015612ed957600080fd5b505afa158015612eed573d6000803e3d6000fd5b505050506040513d6020811015612f0357600080fd5b5051612f0e83613f24565b14612f52576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b600e54604080516306180fd960e11b8152815160009384936001600160a01b0390911692630c301fb29260048083019392829003018186803b158015612f9757600080fd5b505afa158015612fab573d6000803e3d6000fd5b505050506040513d6040811015612fc157600080fd5b508051602090910151608086015161012087015192945090925083031015613021576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b8351430361302d611ae9565b81101561306e576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b84608001518560e00151038561012001511015806130955750600b54810285610140015110155b6130d2576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b600b548102600402856101400151111561311f576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b6000600a54430190506000856001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561316357600080fd5b505afa158015613177573d6000803e3d6000fd5b505050506040513d602081101561318d57600080fd5b505190508082101561319d578091505b600b54876101400151816131ad57fe5b60115491900492909201916000906001600160a01b031663d45ab2b56131d38a89613f56565b6131ef8b8a8a600b548f6101400151816131e957fe5b04613fa5565b6131f88c614021565b61320133611acb565b886040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b15801561325457600080fd5b505af1158015613268573d6000803e3d6000fd5b505050506040513d602081101561327e57600080fd5b5051905061328b81614037565b613295338c613cf4565b8a7f4807480f255627d9b9350200277cb372949a76ccb4263935665257e2b3a1582a8b8b89896040518085600760200280828437600083820152601f01601f191690910190508461014080828437600083820152601f01601f19169091019384525050602082015260408051918290030192509050a250505050505050505050505050565b60006007828154811061332957fe5b9060005260206000209060020201600101549050919050565b60095460ff1615613388576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b613391826136bd565b6112368261339e83614070565b6141de565b6133ad8282614203565b6001600160a01b0316336001600160a01b031614613401576040805162461bcd60e51b815260206004820152600c60248201526b2ba927a723afa9a2a72222a960a11b604482015290519081900360640190fd5b600061340c82612cf3565b9050600061341984612cf3565b9050808211156134325761342d8382613754565b820391505b6002820461344085826141de565b808303925061344e85614280565b6110cf846142aa565b60095460ff161561349d576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b6134a633611d41565b156134e9576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b60006134f482614070565b90506134fe611b03565b811015613545576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b611236338261435b565b6012546001600160a01b0316331461359b576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b600f80546001600160a01b0319166001600160a01b03838116918217909255600e54604080516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b158015611d0f57600080fd5b60095460ff16613653576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6009805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613686614425565b604080516001600160a01b039092168252519081900360200190a1565b6136ae600154614429565b6001805481019055565b600255565b6136c681611d41565b613704576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b600061370f82611d9f565b6001600160a01b031614611e2a576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b03821660009081526006602052604081206002810154808411156137b9576040805162461bcd60e51b815260206004820152601060248201526f544f4f5f4c4954544c455f5354414b4560801b604482015290519081900360640190fd5b60029091018390556001600160a01b03841660009081526008602052604090208054918490039182019055905092915050565b80518251600091829182805b8381101561389f57600087828151811061380e57fe5b6020026020010151905083818701111561385e576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868a01810182902060408051808401969096528581019190915280518086038201815260609095019052835193019290922091909401936001016137f8565b508184146138e2576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b9695505050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b613923600054614429565b60018054600081905581019055565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b6000805260046020527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec80546001600160a01b0319166001600160a01b039290921691909117905560018055565b6001600160a01b03811660009081526006602090815260408083206002810154600890935292208054909101905561123681836144ad565b600780546000198101908110613a8d57fe5b906000526020600020906002020160078281548110613aa857fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b039092169190911781556001918201549101556007805480613aeb57fe5b60008281526020812060026000199093019283020180546001600160a01b031916815560010155905550565b8060078381548110613b2557fe5b9060005260206000209060020201600101819055505050565b6001600160a01b03166000908152600860205260408120805491905590565b80613b6757611236565b600d546001600160a01b0316613bb3576040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015613bad573d6000803e3d6000fd5b50611236565b600d546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015613c0957600080fd5b505af1158015613c1d573d6000803e3d6000fd5b505050506040513d6020811015613c3357600080fd5b5051611236576040805162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b604482015290519081900360640190fd5b60095460ff1615613cbe576040805162461bcd60e51b815260206004820152601060248201526000805160206146f8833981519152604482015290519081900360640190fd5b6009805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613686614425565b6001600160a01b03808316600081815260066020908152604080832086845260049283905281842054825163123334b760e11b8152938401959095529051909493909316928392632466696e926024808201939182900301818387803b158015613d5d57600080fd5b505af1158015613d71573d6000803e3d6000fd5b50505050506001015550565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b6001600160a01b03928316600090815260066020526040808220600390810180549487166001600160a01b0319958616811790915594909516825290209092018054909216179055565b613e0d61466b565b60408051610220810182528351815260208085015181830152855182840152850151606080830191909152848301516080808401919091529085015160a0808401919091529085015160c083015284015160e082015290840151610100820152610120810183600660200201518152602001836007600a8110613e8c57fe5b6020020151815260200184600360078110613ea357fe5b60200201518152602001836008600a8110613eba57fe5b6020020151815260200184600460078110613ed157fe5b60200201518152602001836009600a8110613ee857fe5b6020020151815260200184600560078110613eff57fe5b6020020151815260200184600660078110613f1657fe5b602002015190529392505050565b600061211f826000015183602001518460400151856060015186608001518760a001518860c001518960e0015161399b565b6000613f9e43846101400151856020015101856102000151866101e001518761012001518860800151018861018001518960a0015101896101c001518a60c00151018961399b565b9392505050565b600080613fc76000876101200151886080015188030386896101e00151613d7d565b905060006140006000886101200151613fe88a6101e001516000801b6138ec565b613ffb8b606001518c61010001516138ec565b613d7d565b905061401682826140108a6145bb565b87613d7d565b979650505050505050565b600061211f826101600151836101a001516138ec565b6002805460010190819055600090815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b600d546000906001600160a01b03166140cc5781156140c5576040805162461bcd60e51b815260206004820152600c60248201526b4241445f53544b5f5459504560a01b604482015290519081900360640190fd5b5034610d7f565b341561410e576040805162461bcd60e51b815260206004820152600c60248201526b4241445f53544b5f5459504560a01b604482015290519081900360640190fd5b600d54604080516323b872dd60e01b81523360048201523060248201526044810185905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b15801561416857600080fd5b505af115801561417c573d6000803e3d6000fd5b505050506040513d602081101561419257600080fd5b50516141d7576040805162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b604482015290519081900360640190fd5b5080610d7f565b6001600160a01b03909116600090815260066020526040902060020180549091019055565b6001600160a01b0380831660009081526006602052604080822084841683529082206003808301549082015493949293919290811691168114614277576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b95945050505050565b6001600160a01b0316600090815260066020526040902060030180546001600160a01b0319169055565b6001600160a01b0381811660008181526006602090815260408083208151808301909252938152600180850154928201928352600780549182018155909352517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688600290930292830180546001600160a01b031916919095161790935591517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6899092019190915561123681836144ad565b6005805460018082019092557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0810180546001600160a01b03199081166001600160a01b039687169081179092556040805160a08101825293845260008054602080870191825286840198895260608701838152608088018981529684526006909152929091209451855551948401949094559351600283015591516003918201805493519390941694169390931760ff60a01b1916600160a01b91151591909102179055439055565b3390565b600081815260046020819052604080832054815163083197ef60e41b815291516001600160a01b03909116936383197ef0938381019391929182900301818387803b15801561447757600080fd5b505af115801561448b573d6000803e3d6000fd5b50505060009182525060046020526040902080546001600160a01b0319169055565b81546005805460001981019081106144c157fe5b600091825260209091200154600580546001600160a01b0390921691839081106144e757fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600660006005848154811061452757fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600580548061455757fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b039390931681526006909252506040812081815560018101829055600281019190915560030180546001600160a81b031916905550565b600061211f60008361014001516145f160006145ec87610100015188604001516000801b60008060001b6000614620565b6138ec565b613ffb8661014001516145ec6000801b8961020001518a61016001518b61018001518c6101a001518d6101c001515b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081018290526102008101919091529056fe5061757361626c653a2070617573656400000000000000000000000000000000a2646970667358221220bc781b805da6321b5a2a2ecbfc007bf638ac7caa10e6092c37cf34bb2ab427b064736f6c634300060c0033"

// DeployRollup deploys a new Ethereum contract, binding an instance of Rollup to it.
func DeployRollup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Rollup, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupBin), backend)
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

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "_stakerMap", arg0)

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

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _Rollup.Contract.StakerMap(&_Rollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _Rollup.Contract.StakerMap(&_Rollup.CallOpts, arg0)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_Rollup *RollupCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_Rollup *RollupSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _Rollup.Contract.AmountStaked(&_Rollup.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_Rollup *RollupCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _Rollup.Contract.AmountStaked(&_Rollup.CallOpts, staker)
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

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Rollup *RollupCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Rollup *RollupSession) Bridge() (common.Address, error) {
	return _Rollup.Contract.Bridge(&_Rollup.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Rollup *RollupCallerSession) Bridge() (common.Address, error) {
	return _Rollup.Contract.Bridge(&_Rollup.CallOpts)
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

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_Rollup *RollupCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_Rollup *RollupSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _Rollup.Contract.CurrentChallenge(&_Rollup.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_Rollup *RollupCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _Rollup.Contract.CurrentChallenge(&_Rollup.CallOpts, staker)
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

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 i) view returns(address)
func (_Rollup *RollupCaller) GetNode(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getNode", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 i) view returns(address)
func (_Rollup *RollupSession) GetNode(i *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetNode(&_Rollup.CallOpts, i)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 i) view returns(address)
func (_Rollup *RollupCallerSession) GetNode(i *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetNode(&_Rollup.CallOpts, i)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 i) view returns(address)
func (_Rollup *RollupCaller) GetStakerAddress(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getStakerAddress", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 i) view returns(address)
func (_Rollup *RollupSession) GetStakerAddress(i *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetStakerAddress(&_Rollup.CallOpts, i)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 i) view returns(address)
func (_Rollup *RollupCallerSession) GetStakerAddress(i *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetStakerAddress(&_Rollup.CallOpts, i)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_Rollup *RollupCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_Rollup *RollupSession) IsStaked(staker common.Address) (bool, error) {
	return _Rollup.Contract.IsStaked(&_Rollup.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_Rollup *RollupCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _Rollup.Contract.IsStaked(&_Rollup.CallOpts, staker)
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

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_Rollup *RollupCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_Rollup *RollupSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _Rollup.Contract.LatestStakedNode(&_Rollup.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_Rollup *RollupCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _Rollup.Contract.LatestStakedNode(&_Rollup.CallOpts, staker)
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

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_Rollup *RollupCaller) Outbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "outbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_Rollup *RollupSession) Outbox() (common.Address, error) {
	return _Rollup.Contract.Outbox(&_Rollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_Rollup *RollupCallerSession) Outbox() (common.Address, error) {
	return _Rollup.Contract.Outbox(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCallerSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupSession) Paused() (bool, error) {
	return _Rollup.Contract.Paused(&_Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupCallerSession) Paused() (bool, error) {
	return _Rollup.Contract.Paused(&_Rollup.CallOpts)
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

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_Rollup *RollupCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_Rollup *RollupSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _Rollup.Contract.WithdrawableFunds(&_Rollup.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_Rollup *RollupCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _Rollup.Contract.WithdrawableFunds(&_Rollup.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 i) view returns(address)
func (_Rollup *RollupCaller) ZombieAddress(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieAddress", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 i) view returns(address)
func (_Rollup *RollupSession) ZombieAddress(i *big.Int) (common.Address, error) {
	return _Rollup.Contract.ZombieAddress(&_Rollup.CallOpts, i)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 i) view returns(address)
func (_Rollup *RollupCallerSession) ZombieAddress(i *big.Int) (common.Address, error) {
	return _Rollup.Contract.ZombieAddress(&_Rollup.CallOpts, i)
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

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 i) view returns(uint256)
func (_Rollup *RollupCaller) ZombieLatestStakedNode(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieLatestStakedNode", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 i) view returns(uint256)
func (_Rollup *RollupSession) ZombieLatestStakedNode(i *big.Int) (*big.Int, error) {
	return _Rollup.Contract.ZombieLatestStakedNode(&_Rollup.CallOpts, i)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 i) view returns(uint256)
func (_Rollup *RollupCallerSession) ZombieLatestStakedNode(i *big.Int) (*big.Int, error) {
	return _Rollup.Contract.ZombieLatestStakedNode(&_Rollup.CallOpts, i)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) payable returns()
func (_Rollup *RollupTransactor) AddToDeposit(opts *bind.TransactOpts, stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addToDeposit", stakerAddress, tokenAmount)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) payable returns()
func (_Rollup *RollupSession) AddToDeposit(stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.AddToDeposit(&_Rollup.TransactOpts, stakerAddress, tokenAmount)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) payable returns()
func (_Rollup *RollupTransactorSession) AddToDeposit(stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.AddToDeposit(&_Rollup.TransactOpts, stakerAddress, tokenAmount)
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

// Initialize is a paid mutator transaction binding the contract method 0x3e03cf43.
//
// Solidity: function initialize(address _outbox, bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _bridge, address _challengeFactory, address _nodeFactory, bytes _extraConfig, address _admin) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _outbox common.Address, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _bridge common.Address, _challengeFactory common.Address, _nodeFactory common.Address, _extraConfig []byte, _admin common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _outbox, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _bridge, _challengeFactory, _nodeFactory, _extraConfig, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e03cf43.
//
// Solidity: function initialize(address _outbox, bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _bridge, address _challengeFactory, address _nodeFactory, bytes _extraConfig, address _admin) returns()
func (_Rollup *RollupSession) Initialize(_outbox common.Address, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _bridge common.Address, _challengeFactory common.Address, _nodeFactory common.Address, _extraConfig []byte, _admin common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _outbox, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _bridge, _challengeFactory, _nodeFactory, _extraConfig, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e03cf43.
//
// Solidity: function initialize(address _outbox, bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _bridge, address _challengeFactory, address _nodeFactory, bytes _extraConfig, address _admin) returns()
func (_Rollup *RollupTransactorSession) Initialize(_outbox common.Address, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _bridge common.Address, _challengeFactory common.Address, _nodeFactory common.Address, _extraConfig []byte, _admin common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _outbox, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _bridge, _challengeFactory, _nodeFactory, _extraConfig, _admin)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) payable returns()
func (_Rollup *RollupTransactor) NewStake(opts *bind.TransactOpts, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "newStake", tokenAmount)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) payable returns()
func (_Rollup *RollupSession) NewStake(tokenAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.NewStake(&_Rollup.TransactOpts, tokenAmount)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) payable returns()
func (_Rollup *RollupTransactorSession) NewStake(tokenAmount *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.NewStake(&_Rollup.TransactOpts, tokenAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rollup *RollupTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rollup *RollupSession) Pause() (*types.Transaction, error) {
	return _Rollup.Contract.Pause(&_Rollup.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rollup *RollupTransactorSession) Pause() (*types.Transaction, error) {
	return _Rollup.Contract.Pause(&_Rollup.TransactOpts)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_Rollup *RollupTransactor) ReduceDeposit(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "reduceDeposit", target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_Rollup *RollupSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ReduceDeposit(&_Rollup.TransactOpts, target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_Rollup *RollupTransactorSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ReduceDeposit(&_Rollup.TransactOpts, target)
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

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_Rollup *RollupTransactor) RemoveOldOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeOldOutbox", _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_Rollup *RollupSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldOutbox(&_Rollup.TransactOpts, _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_Rollup *RollupTransactorSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldOutbox(&_Rollup.TransactOpts, _outbox)
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

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Rollup *RollupTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Rollup *RollupSession) Resume() (*types.Transaction, error) {
	return _Rollup.Contract.Resume(&_Rollup.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Rollup *RollupTransactorSession) Resume() (*types.Transaction, error) {
	return _Rollup.Contract.Resume(&_Rollup.TransactOpts)
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

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_Rollup *RollupTransactor) SetInbox(opts *bind.TransactOpts, _inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "setInbox", _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_Rollup *RollupSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _Rollup.Contract.SetInbox(&_Rollup.TransactOpts, _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_Rollup *RollupTransactorSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _Rollup.Contract.SetInbox(&_Rollup.TransactOpts, _inbox, _enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_Rollup *RollupTransactor) SetOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "setOutbox", _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_Rollup *RollupSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.SetOutbox(&_Rollup.TransactOpts, _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_Rollup *RollupTransactorSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.SetOutbox(&_Rollup.TransactOpts, _outbox)
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

// TruncateNodes is a paid mutator transaction binding the contract method 0x15edeb3a.
//
// Solidity: function truncateNodes(uint256 _latestNodeCreated) returns()
func (_Rollup *RollupTransactor) TruncateNodes(opts *bind.TransactOpts, _latestNodeCreated *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "truncateNodes", _latestNodeCreated)
}

// TruncateNodes is a paid mutator transaction binding the contract method 0x15edeb3a.
//
// Solidity: function truncateNodes(uint256 _latestNodeCreated) returns()
func (_Rollup *RollupSession) TruncateNodes(_latestNodeCreated *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.TruncateNodes(&_Rollup.TransactOpts, _latestNodeCreated)
}

// TruncateNodes is a paid mutator transaction binding the contract method 0x15edeb3a.
//
// Solidity: function truncateNodes(uint256 _latestNodeCreated) returns()
func (_Rollup *RollupTransactorSession) TruncateNodes(_latestNodeCreated *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.TruncateNodes(&_Rollup.TransactOpts, _latestNodeCreated)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_Rollup *RollupTransactor) UpgradeImplementation(opts *bind.TransactOpts, _newRollup common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "upgradeImplementation", _newRollup)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_Rollup *RollupSession) UpgradeImplementation(_newRollup common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.UpgradeImplementation(&_Rollup.TransactOpts, _newRollup)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_Rollup *RollupTransactorSession) UpgradeImplementation(_newRollup common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.UpgradeImplementation(&_Rollup.TransactOpts, _newRollup)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_Rollup *RollupTransactor) UpgradeImplementationAndCall(opts *bind.TransactOpts, _newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "upgradeImplementationAndCall", _newRollup, _data)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_Rollup *RollupSession) UpgradeImplementationAndCall(_newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _Rollup.Contract.UpgradeImplementationAndCall(&_Rollup.TransactOpts, _newRollup, _data)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_Rollup *RollupTransactorSession) UpgradeImplementationAndCall(_newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _Rollup.Contract.UpgradeImplementationAndCall(&_Rollup.TransactOpts, _newRollup, _data)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_Rollup *RollupTransactor) WithdrawStakerFunds(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "withdrawStakerFunds", destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_Rollup *RollupSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStakerFunds(&_Rollup.TransactOpts, destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_Rollup *RollupTransactorSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStakerFunds(&_Rollup.TransactOpts, destination)
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

// RollupPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Rollup contract.
type RollupPausedIterator struct {
	Event *RollupPaused // Event containing the contract specifics and raw log

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
func (it *RollupPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupPaused)
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
		it.Event = new(RollupPaused)
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
func (it *RollupPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupPaused represents a Paused event raised by the Rollup contract.
type RollupPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) FilterPaused(opts *bind.FilterOpts) (*RollupPausedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RollupPausedIterator{contract: _Rollup.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RollupPaused) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupPaused)
				if err := _Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) ParsePaused(log types.Log) (*RollupPaused, error) {
	event := new(RollupPaused)
	if err := _Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
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

// RollupUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Rollup contract.
type RollupUnpausedIterator struct {
	Event *RollupUnpaused // Event containing the contract specifics and raw log

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
func (it *RollupUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUnpaused)
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
		it.Event = new(RollupUnpaused)
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
func (it *RollupUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUnpaused represents a Unpaused event raised by the Rollup contract.
type RollupUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RollupUnpausedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RollupUnpausedIterator{contract: _Rollup.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RollupUnpaused) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUnpaused)
				if err := _Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) ParseUnpaused(log types.Log) (*RollupUnpaused, error) {
	event := new(RollupUnpaused)
	if err := _Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCoreABI is the input ABI used to generate the binding from.
const RollupCoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupCoreBin is the compiled bytecode used for deploying new contracts.
var RollupCoreBin = "0x608060405234801561001057600080fd5b506104a7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100d05760003560e01c80632f30cabd146100d55780633e96576e1461010d5780634f0f4aa9146101335780636177fd181461016c57806362a82d7d146101a657806363721d6b146101c357806365f7f80d146101cb57806369fd251c146101d35780637ba9534a146101f95780638640ce5f14610201578063d01e660214610209578063d735e21d14610226578063dff697871461022e578063e8bd492214610236578063ef40a67014610292578063f33e1fac146102b8575b600080fd5b6100fb600480360360208110156100eb57600080fd5b50356001600160a01b03166102d5565b60408051918252519081900360200190f35b6100fb6004803603602081101561012357600080fd5b50356001600160a01b03166102f0565b6101506004803603602081101561014957600080fd5b503561030e565b604080516001600160a01b039092168252519081900360200190f35b6101926004803603602081101561018257600080fd5b50356001600160a01b0316610329565b604080519115158252519081900360200190f35b610150600480360360208110156101bc57600080fd5b5035610351565b6100fb61037b565b6100fb610381565b610150600480360360208110156101e957600080fd5b50356001600160a01b0316610387565b6100fb6103a8565b6100fb6103ae565b6101506004803603602081101561021f57600080fd5b50356103b4565b6100fb6103e3565b6100fb6103e9565b61025c6004803603602081101561024c57600080fd5b50356001600160a01b03166103ef565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b6100fb600480360360208110156102a857600080fd5b50356001600160a01b031661042b565b6100fb600480360360208110156102ce57600080fd5b5035610449565b6001600160a01b031660009081526008602052604090205490565b6001600160a01b031660009081526006602052604090206001015490565b6000908152600460205260409020546001600160a01b031690565b6001600160a01b0316600090815260066020526040902060030154600160a01b900460ff1690565b60006005828154811061036057fe5b6000918252602090912001546001600160a01b031692915050565b60075490565b60005490565b6001600160a01b039081166000908152600660205260409020600301541690565b60025490565b60035490565b6000600782815481106103c357fe5b60009182526020909120600290910201546001600160a01b031692915050565b60015490565b60055490565b6006602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6001600160a01b031660009081526006602052604090206002015490565b60006007828154811061045857fe5b906000526020600020906002020160010154905091905056fea2646970667358221220beb436aa93f28e7a0f641658ae3bce116352f23689fcb9f13ffb9d37285fc69664736f6c634300060c0033"

// DeployRollupCore deploys a new Ethereum contract, binding an instance of RollupCore to it.
func DeployRollupCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupCore, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCore{RollupCoreCaller: RollupCoreCaller{contract: contract}, RollupCoreTransactor: RollupCoreTransactor{contract: contract}, RollupCoreFilterer: RollupCoreFilterer{contract: contract}}, nil
}

// RollupCore is an auto generated Go binding around an Ethereum contract.
type RollupCore struct {
	RollupCoreCaller     // Read-only binding to the contract
	RollupCoreTransactor // Write-only binding to the contract
	RollupCoreFilterer   // Log filterer for contract events
}

// RollupCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCoreSession struct {
	Contract     *RollupCore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCoreCallerSession struct {
	Contract *RollupCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCoreTransactorSession struct {
	Contract     *RollupCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCoreRaw struct {
	Contract *RollupCore // Generic contract binding to access the raw methods on
}

// RollupCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCoreCallerRaw struct {
	Contract *RollupCoreCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCoreTransactorRaw struct {
	Contract *RollupCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCore creates a new instance of RollupCore, bound to a specific deployed contract.
func NewRollupCore(address common.Address, backend bind.ContractBackend) (*RollupCore, error) {
	contract, err := bindRollupCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCore{RollupCoreCaller: RollupCoreCaller{contract: contract}, RollupCoreTransactor: RollupCoreTransactor{contract: contract}, RollupCoreFilterer: RollupCoreFilterer{contract: contract}}, nil
}

// NewRollupCoreCaller creates a new read-only instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreCaller(address common.Address, caller bind.ContractCaller) (*RollupCoreCaller, error) {
	contract, err := bindRollupCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCoreCaller{contract: contract}, nil
}

// NewRollupCoreTransactor creates a new write-only instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCoreTransactor, error) {
	contract, err := bindRollupCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCoreTransactor{contract: contract}, nil
}

// NewRollupCoreFilterer creates a new log filterer instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCoreFilterer, error) {
	contract, err := bindRollupCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCoreFilterer{contract: contract}, nil
}

// bindRollupCore binds a generic wrapper to an already deployed contract.
func bindRollupCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCore *RollupCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCore.Contract.RollupCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCore *RollupCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCore.Contract.RollupCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCore *RollupCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCore.Contract.RollupCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCore *RollupCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCore *RollupCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCore *RollupCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCore.Contract.contract.Transact(opts, method, params...)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "_stakerMap", arg0)

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

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _RollupCore.Contract.StakerMap(&_RollupCore.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _RollupCore.Contract.StakerMap(&_RollupCore.CallOpts, arg0)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.AmountStaked(&_RollupCore.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.AmountStaked(&_RollupCore.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _RollupCore.Contract.CurrentChallenge(&_RollupCore.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _RollupCore.Contract.CurrentChallenge(&_RollupCore.CallOpts, staker)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreSession) FirstUnresolvedNode() (*big.Int, error) {
	return _RollupCore.Contract.FirstUnresolvedNode(&_RollupCore.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _RollupCore.Contract.FirstUnresolvedNode(&_RollupCore.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 i) view returns(address)
func (_RollupCore *RollupCoreCaller) GetNode(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getNode", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 i) view returns(address)
func (_RollupCore *RollupCoreSession) GetNode(i *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetNode(&_RollupCore.CallOpts, i)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 i) view returns(address)
func (_RollupCore *RollupCoreCallerSession) GetNode(i *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetNode(&_RollupCore.CallOpts, i)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 i) view returns(address)
func (_RollupCore *RollupCoreCaller) GetStakerAddress(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getStakerAddress", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 i) view returns(address)
func (_RollupCore *RollupCoreSession) GetStakerAddress(i *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetStakerAddress(&_RollupCore.CallOpts, i)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 i) view returns(address)
func (_RollupCore *RollupCoreCallerSession) GetStakerAddress(i *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetStakerAddress(&_RollupCore.CallOpts, i)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreSession) IsStaked(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsStaked(&_RollupCore.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsStaked(&_RollupCore.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreSession) LastStakeBlock() (*big.Int, error) {
	return _RollupCore.Contract.LastStakeBlock(&_RollupCore.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LastStakeBlock() (*big.Int, error) {
	return _RollupCore.Contract.LastStakeBlock(&_RollupCore.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestConfirmed() (*big.Int, error) {
	return _RollupCore.Contract.LatestConfirmed(&_RollupCore.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestConfirmed() (*big.Int, error) {
	return _RollupCore.Contract.LatestConfirmed(&_RollupCore.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestNodeCreated() (*big.Int, error) {
	return _RollupCore.Contract.LatestNodeCreated(&_RollupCore.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _RollupCore.Contract.LatestNodeCreated(&_RollupCore.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.LatestStakedNode(&_RollupCore.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.LatestStakedNode(&_RollupCore.CallOpts, staker)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreSession) StakerCount() (*big.Int, error) {
	return _RollupCore.Contract.StakerCount(&_RollupCore.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) StakerCount() (*big.Int, error) {
	return _RollupCore.Contract.StakerCount(&_RollupCore.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _RollupCore.Contract.WithdrawableFunds(&_RollupCore.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _RollupCore.Contract.WithdrawableFunds(&_RollupCore.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 i) view returns(address)
func (_RollupCore *RollupCoreCaller) ZombieAddress(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieAddress", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 i) view returns(address)
func (_RollupCore *RollupCoreSession) ZombieAddress(i *big.Int) (common.Address, error) {
	return _RollupCore.Contract.ZombieAddress(&_RollupCore.CallOpts, i)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 i) view returns(address)
func (_RollupCore *RollupCoreCallerSession) ZombieAddress(i *big.Int) (common.Address, error) {
	return _RollupCore.Contract.ZombieAddress(&_RollupCore.CallOpts, i)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreSession) ZombieCount() (*big.Int, error) {
	return _RollupCore.Contract.ZombieCount(&_RollupCore.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) ZombieCount() (*big.Int, error) {
	return _RollupCore.Contract.ZombieCount(&_RollupCore.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 i) view returns(uint256)
func (_RollupCore *RollupCoreCaller) ZombieLatestStakedNode(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieLatestStakedNode", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 i) view returns(uint256)
func (_RollupCore *RollupCoreSession) ZombieLatestStakedNode(i *big.Int) (*big.Int, error) {
	return _RollupCore.Contract.ZombieLatestStakedNode(&_RollupCore.CallOpts, i)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 i) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) ZombieLatestStakedNode(i *big.Int) (*big.Int, error) {
	return _RollupCore.Contract.ZombieLatestStakedNode(&_RollupCore.CallOpts, i)
}

// RollupLibABI is the input ABI used to generate the binding from.
const RollupLibABI = "[]"

// RollupLibBin is the compiled bytecode used for deploying new contracts.
var RollupLibBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e09d6de0a828720abe58dea03e199cccc66fa99c5d7364236d6c1b637f43ebcd64736f6c634300060c0033"

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

// ValidatorUtilsABI is the input ABI used to generate the binding from.
const ValidatorUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkConfirmable\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkDecidableNextNode\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.ConfirmType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkRejectableNextNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"node1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findNodeConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findStakerConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"refundableStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakedNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"stakerInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"successorNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorUtilsBin is the compiled bytecode used for deploying new contracts.
var ValidatorUtilsBin = "0x608060405234801561001057600080fd5b50612708806100206000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c80632b1062cf1461009e5780633082d029146100c9578063422e3550146100eb5780637464ae061461010c5780637988ad371461012c5780638730825e1461013f5780638f67e6bb1461015f578063abeba4f714610182578063c2b553ec14610195578063c308eaaf146101aa578063e48a5f7b146101bd575b600080fd5b6100b16100ac3660046123d9565b6101e0565b6040516100c093929190612545565b60405180910390f35b6100dc6100d736600461239f565b6102e4565b6040516100c093929190612570565b6100fe6100f93660046123d9565b6107b5565b6040516100c092919061267f565b61011f61011a36600461229c565b610946565b6040516100c0919061246c565b6100dc61013a3660046122f0565b610bcc565b61015261014d366004612340565b610cee565b6040516100c091906124b9565b61017261016d3660046122b8565b610ec3565b6040516100c094939291906124f1565b61011f61019036600461236b565b6110c9565b6101a86101a336600461229c565b611253565b005b6101526101b83660046122b8565b6116e1565b6101d06101cb36600461229c565b611926565b6040516100c09493929190612696565b6040516330ad54fb60e21b815260009081908190309063c2b553ec9061020a908b90600401612458565b60006040518083038186803b15801561022257600080fd5b505afa925050508015610233575060015b61023c5761024b565b506001915060009050806102d9565b604051630422e35560e41b8152309063422e355090610276908b908b908b908b908b90600401612517565b604080518083038186803b15801561028d57600080fd5b505afa9250505080156102bd575060408051601f3d908101601f191682019092526102ba91810190612434565b60015b6102cf575060009150819050806102d9565b6002945090925090505b955095509592505050565b600080600080876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561032357600080fd5b505afa158015610337573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035b919061241c565b90506000886001600160a01b0316634f0f4aa9896040518263ffffffff1660e01b815260040161038b9190612676565b60206040518083038186803b1580156103a357600080fd5b505afa1580156103b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103db9190612260565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561041357600080fd5b505afa158015610427573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044b919061241c565b90506000896001600160a01b0316634f0f4aa9896040518263ffffffff1660e01b815260040161047b9190612676565b60206040518083038186803b15801561049357600080fd5b505afa1580156104a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cb9190612260565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561050357600080fd5b505afa158015610517573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053b919061241c565b905060005b8781101561079c57888a14156105635760008a8a965096509650505050506107ab565b8183141561057e5760018a8a965096509650505050506107ab565b8383108061058b57508382105b156105a4576002600080965096509650505050506107ab565b888a10156106a2578198508a6001600160a01b0316634f0f4aa98a6040518263ffffffff1660e01b81526004016105db9190612676565b60206040518083038186803b1580156105f357600080fd5b505afa158015610607573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062b9190612260565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561066357600080fd5b505afa158015610677573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069b919061241c565b9150610794565b8299508a6001600160a01b0316634f0f4aa98b6040518263ffffffff1660e01b81526004016106d19190612676565b60206040518083038186803b1580156106e957600080fd5b505afa1580156106fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107219190612260565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561075957600080fd5b505afa15801561076d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610791919061241c565b92505b600101610540565b50600389899550955095505050505b9450945094915050565b600080866001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b1580156107f157600080fd5b505afa158015610805573d6000803e3d6000fd5b50505050866001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561084257600080fd5b505afa158015610856573d6000803e3d6000fd5b505050506000876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561089557600080fd5b505afa1580156108a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108cd919061241c565b905060006108da89611aff565b905080156108f05760008093509350505061093c565b60008060006109078c8c87600101018c8c8c611f21565b925092509250826109335760405162461bcd60e51b815260040161092a90612629565b60405180910390fd5b90955093505050505b9550959350505050565b60606000826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561098357600080fd5b505afa158015610997573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bb919061241c565b90506060816001600160401b03811180156109d557600080fd5b506040519080825280602002602001820160405280156109ff578160200160208202803683370190505b5090506000846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a3d57600080fd5b505afa158015610a51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a75919061241c565b90506000805b84811015610bc1576040516362a82d7d60e01b81526000906001600160a01b038916906362a82d7d90610ab2908590600401612676565b60206040518083038186803b158015610aca57600080fd5b505afa158015610ade573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b029190612260565b90506000886001600160a01b0316633e96576e836040518263ffffffff1660e01b8152600401610b329190612458565b60206040518083038186803b158015610b4a57600080fd5b505afa158015610b5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b82919061241c565b9050848111610bb75781868581518110610b9857fe5b6001600160a01b03909216602092830291909101909101526001909301925b5050600101610a7b565b508252509392505050565b600080600080876001600160a01b0316633e96576e886040518263ffffffff1660e01b8152600401610bfe9190612458565b60206040518083038186803b158015610c1657600080fd5b505afa158015610c2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4e919061241c565b90506000886001600160a01b0316633e96576e886040518263ffffffff1660e01b8152600401610c7e9190612458565b60206040518083038186803b158015610c9657600080fd5b505afa158015610caa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cce919061241c565b9050610cdc898383896102e4565b94509450945050509450945094915050565b60408051620186a08082526230d4208201909252606091829190602082016230d400803683370190505090506000600184015b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d5a57600080fd5b505afa158015610d6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d92919061241c565b8111610eb957604051634f0f4aa960e01b81526000906001600160a01b03881690634f0f4aa990610dc7908590600401612676565b60206040518083038186803b158015610ddf57600080fd5b505afa158015610df3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e179190612260565b905085816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610e5357600080fd5b505afa158015610e67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8b919061241c565b1415610eb05781848481518110610e9e57fe5b60209081029190910101526001909201915b50600101610d21565b5081529392505050565b600080600080856001600160a01b0316636177fd18866040518263ffffffff1660e01b8152600401610ef59190612458565b60206040518083038186803b158015610f0d57600080fd5b505afa158015610f21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f45919061227c565b604051631f4b2bb760e11b81526001600160a01b03881690633e96576e90610f71908990600401612458565b60206040518083038186803b158015610f8957600080fd5b505afa158015610f9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc1919061241c565b604051630ef40a6760e41b81526001600160a01b0389169063ef40a67090610fed908a90600401612458565b60206040518083038186803b15801561100557600080fd5b505afa158015611019573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103d919061241c565b604051631a7f494760e21b81526001600160a01b038a16906369fd251c90611069908b90600401612458565b60206040518083038186803b15801561108157600080fd5b505afa158015611095573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b99190612260565b9299919850965090945092505050565b60606000846001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561110657600080fd5b505afa15801561111a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061113e919061241c565b905080838501101561114f57508282015b6060816001600160401b038111801561116757600080fd5b50604051908082528060200260200182016040528015611191578160200160208202803683370190505b50905060005b82811015611249576040516362a82d7d60e01b81526001600160a01b038816906362a82d7d906111cd9089850190600401612676565b60206040518083038186803b1580156111e557600080fd5b505afa1580156111f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121d9190612260565b82828151811061122957fe5b6001600160a01b0390921660209283029190910190910152600101611197565b5095945050505050565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b15801561128c57600080fd5b505afa1580156112a0573d6000803e3d6000fd5b50505050806001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b1580156112dd57600080fd5b505afa1580156112f1573d6000803e3d6000fd5b505050506000816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561133057600080fd5b505afa158015611344573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611368919061241c565b90506000811161138a5760405162461bcd60e51b815260040161092a906125b7565b6000826001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156113c557600080fd5b505afa1580156113d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113fd919061241c565b90506000836001600160a01b0316634f0f4aa9836040518263ffffffff1660e01b815260040161142d9190612676565b60206040518083038186803b15801561144557600080fd5b505afa158015611459573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147d9190612260565b9050806001600160a01b0316630cad16a56040518163ffffffff1660e01b815260040160006040518083038186803b1580156114b857600080fd5b505afa1580156114cc573d6000803e3d6000fd5b50505050836001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561150957600080fd5b505afa15801561151d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611541919061241c565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561157a57600080fd5b505afa15801561158e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b2919061241c565b146115cf5760405162461bcd60e51b815260040161092a90612603565b604051630128a01960e21b81526001600160a01b038516906304a28064906115fb908490600401612458565b60206040518083038186803b15801561161357600080fd5b505afa158015611627573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164b919061241c565b8301816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561168657600080fd5b505afa15801561169a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116be919061241c565b146116db5760405162461bcd60e51b815260040161092a906125db565b50505050565b60408051620186a08082526230d4208201909252606091829190602082016230d40080368337019050509050600080856001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561174957600080fd5b505afa15801561175d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611781919061241c565b90505b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117bd57600080fd5b505afa1580156117d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f5919061241c565b8111610eb957604051634f0f4aa960e01b81526000906001600160a01b03881690634f0f4aa99061182a908590600401612676565b60206040518083038186803b15801561184257600080fd5b505afa158015611856573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187a9190612260565b6040516348b4573960e11b81529091506001600160a01b03821690639168ae72906118a9908990600401612458565b60206040518083038186803b1580156118c157600080fd5b505afa1580156118d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f9919061227c565b1561191d578184848151811061190b57fe5b60209081029190910101526001909201915b50600101611784565b600080600080846001600160a01b03166346c2781a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561196557600080fd5b505afa158015611979573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061199d919061241c565b9350846001600160a01b0316635e8ef1066040518163ffffffff1660e01b815260040160206040518083038186803b1580156119d857600080fd5b505afa1580156119ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a10919061241c565b9250846001600160a01b03166376e7e23b6040518163ffffffff1660e01b815260040160206040518083038186803b158015611a4b57600080fd5b505afa158015611a5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a83919061241c565b9150846001600160a01b03166351ed6a306040518163ffffffff1660e01b815260040160206040518083038186803b158015611abe57600080fd5b505afa158015611ad2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611af69190612260565b90509193509193565b6000816001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b158015611b3a57600080fd5b505afa158015611b4e573d6000803e3d6000fd5b505050506000826001600160a01b0316634f0f4aa9846001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611b9c57600080fd5b505afa158015611bb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd4919061241c565b6040518263ffffffff1660e01b8152600401611bf09190612676565b60206040518083038186803b158015611c0857600080fd5b505afa158015611c1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c409190612260565b90506000836001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c7d57600080fd5b505afa158015611c91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cb5919061241c565b826001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611cee57600080fd5b505afa158015611d02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d26919061241c565b1490508015611f1a57836001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b158015611d6857600080fd5b505afa158015611d7c573d6000803e3d6000fd5b50505050816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611db957600080fd5b505afa158015611dcd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611df1919061241c565b431015611e105760405162461bcd60e51b815260040161092a9061264d565b604051630128a01960e21b81526001600160a01b038516906304a2806490611e3c908590600401612458565b60206040518083038186803b158015611e5457600080fd5b505afa158015611e68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e8c919061241c565b826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015611ec557600080fd5b505afa158015611ed9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611efd919061241c565b14611f1a5760405162461bcd60e51b815260040161092a90612592565b9392505050565b600080600080886001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611f6057600080fd5b505afa158015611f74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f98919061241c565b905080881115611fb3576000806000935093509350506102d9565b87810387811115611fc15750865b6120488a8a8c6001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611fff57600080fd5b505afa158015612013573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612037919061241c565b846120438f8d8d6110c9565b61205b565b9450945094505050955095509592505050565b6000806000808451905060005b86811161224c57604051634f0f4aa960e01b8152898201906000906001600160a01b038d1690634f0f4aa9906120a2908590600401612676565b60206040518083038186803b1580156120ba57600080fd5b505afa1580156120ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120f29190612260565b905089816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561212e57600080fd5b505afa158015612142573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612166919061241c565b14612172575050612244565b60005b8481101561224057816001600160a01b0316639168ae728a838151811061219857fe5b60200260200101516040518263ffffffff1660e01b81526004016121bc9190612458565b60206040518083038186803b1580156121d457600080fd5b505afa1580156121e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061220c919061227c565b15612238576001838a838151811061222057fe5b602002602001015197509750975050505050506102d9565b600101612175565b5050505b600101612068565b506000998a99508998509650505050505050565b600060208284031215612271578081fd5b8151611f1a816126ba565b60006020828403121561228d578081fd5b81518015158114611f1a578182fd5b6000602082840312156122ad578081fd5b8135611f1a816126ba565b600080604083850312156122ca578081fd5b82356122d5816126ba565b915060208301356122e5816126ba565b809150509250929050565b60008060008060808587031215612305578182fd5b8435612310816126ba565b93506020850135612320816126ba565b92506040850135612330816126ba565b9396929550929360600135925050565b60008060408385031215612352578182fd5b823561235d816126ba565b946020939093013593505050565b60008060006060848603121561237f578283fd5b833561238a816126ba565b95602085013595506040909401359392505050565b600080600080608085870312156123b4578384fd5b84356123bf816126ba565b966020860135965060408601359560600135945092505050565b600080600080600060a086880312156123f0578081fd5b85356123fb816126ba565b97602087013597506040870135966060810135965060800135945092505050565b60006020828403121561242d578081fd5b5051919050565b60008060408385031215612446578182fd5b8251915060208301516122e5816126ba565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b818110156124ad5783516001600160a01b031683529284019291840191600101612488565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156124ad578351835292840192918401916001016124d5565b9315158452602084019290925260408301526001600160a01b0316606082015260800190565b6001600160a01b03959095168552602085019390935260408401919091526060830152608082015260a00190565b606081016003851061255357fe5b93815260208101929092526001600160a01b031660409091015290565b606081016004851061257e57fe5b938152602081019290925260409091015290565b6020808252600b908201526a4841535f5354414b45525360a81b604082015260600190565b6020808252600a90820152694e4f5f5354414b45525360b01b604082015260600190565b6020808252600e908201526d1393d517d0531317d4d51052d15160921b604082015260600190565b6020808252600c908201526b24a72b20a624a22fa82922ab60a11b604082015260600190565b6020808252600a90820152694e4f5f4558414d504c4560b01b604082015260600190565b6020808252600f908201526e4245464f52455f444541444c494e4560881b604082015260600190565b90815260200190565b9182526001600160a01b0316602082015260400190565b938452602084019290925260408301526001600160a01b0316606082015260800190565b6001600160a01b03811681146126cf57600080fd5b5056fea2646970667358221220c251e93909aab2d129638c6b21042f1ba8350edb1db137b9f86998172718414d64736f6c634300060c0033"

// DeployValidatorUtils deploys a new Ethereum contract, binding an instance of ValidatorUtils to it.
func DeployValidatorUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorUtils{ValidatorUtilsCaller: ValidatorUtilsCaller{contract: contract}, ValidatorUtilsTransactor: ValidatorUtilsTransactor{contract: contract}, ValidatorUtilsFilterer: ValidatorUtilsFilterer{contract: contract}}, nil
}

// ValidatorUtils is an auto generated Go binding around an Ethereum contract.
type ValidatorUtils struct {
	ValidatorUtilsCaller     // Read-only binding to the contract
	ValidatorUtilsTransactor // Write-only binding to the contract
	ValidatorUtilsFilterer   // Log filterer for contract events
}

// ValidatorUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorUtilsSession struct {
	Contract     *ValidatorUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorUtilsCallerSession struct {
	Contract *ValidatorUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ValidatorUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorUtilsTransactorSession struct {
	Contract     *ValidatorUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ValidatorUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorUtilsRaw struct {
	Contract *ValidatorUtils // Generic contract binding to access the raw methods on
}

// ValidatorUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorUtilsCallerRaw struct {
	Contract *ValidatorUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorUtilsTransactorRaw struct {
	Contract *ValidatorUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorUtils creates a new instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtils(address common.Address, backend bind.ContractBackend) (*ValidatorUtils, error) {
	contract, err := bindValidatorUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtils{ValidatorUtilsCaller: ValidatorUtilsCaller{contract: contract}, ValidatorUtilsTransactor: ValidatorUtilsTransactor{contract: contract}, ValidatorUtilsFilterer: ValidatorUtilsFilterer{contract: contract}}, nil
}

// NewValidatorUtilsCaller creates a new read-only instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorUtilsCaller, error) {
	contract, err := bindValidatorUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsCaller{contract: contract}, nil
}

// NewValidatorUtilsTransactor creates a new write-only instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorUtilsTransactor, error) {
	contract, err := bindValidatorUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsTransactor{contract: contract}, nil
}

// NewValidatorUtilsFilterer creates a new log filterer instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorUtilsFilterer, error) {
	contract, err := bindValidatorUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsFilterer{contract: contract}, nil
}

// bindValidatorUtils binds a generic wrapper to an already deployed contract.
func bindValidatorUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorUtils *ValidatorUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorUtils.Contract.ValidatorUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorUtils *ValidatorUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.ValidatorUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorUtils *ValidatorUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.ValidatorUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorUtils *ValidatorUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorUtils *ValidatorUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorUtils *ValidatorUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.contract.Transact(opts, method, params...)
}

// CheckConfirmable is a free data retrieval call binding the contract method 0xc2b553ec.
//
// Solidity: function checkConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCaller) CheckConfirmable(opts *bind.CallOpts, rollup common.Address) error {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "checkConfirmable", rollup)

	if err != nil {
		return err
	}

	return err

}

// CheckConfirmable is a free data retrieval call binding the contract method 0xc2b553ec.
//
// Solidity: function checkConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsSession) CheckConfirmable(rollup common.Address) error {
	return _ValidatorUtils.Contract.CheckConfirmable(&_ValidatorUtils.CallOpts, rollup)
}

// CheckConfirmable is a free data retrieval call binding the contract method 0xc2b553ec.
//
// Solidity: function checkConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCallerSession) CheckConfirmable(rollup common.Address) error {
	return _ValidatorUtils.Contract.CheckConfirmable(&_ValidatorUtils.CallOpts, rollup)
}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x2b1062cf.
//
// Solidity: function checkDecidableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint8, uint256, address)
func (_ValidatorUtils *ValidatorUtilsCaller) CheckDecidableNextNode(opts *bind.CallOpts, rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (uint8, *big.Int, common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "checkDecidableNextNode", rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x2b1062cf.
//
// Solidity: function checkDecidableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint8, uint256, address)
func (_ValidatorUtils *ValidatorUtilsSession) CheckDecidableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (uint8, *big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.CheckDecidableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x2b1062cf.
//
// Solidity: function checkDecidableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint8, uint256, address)
func (_ValidatorUtils *ValidatorUtilsCallerSession) CheckDecidableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (uint8, *big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.CheckDecidableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// CheckRejectableNextNode is a free data retrieval call binding the contract method 0x422e3550.
//
// Solidity: function checkRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsCaller) CheckRejectableNextNode(opts *bind.CallOpts, rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "checkRejectableNextNode", rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// CheckRejectableNextNode is a free data retrieval call binding the contract method 0x422e3550.
//
// Solidity: function checkRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsSession) CheckRejectableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.CheckRejectableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// CheckRejectableNextNode is a free data retrieval call binding the contract method 0x422e3550.
//
// Solidity: function checkRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsCallerSession) CheckRejectableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.CheckRejectableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCaller) FindNodeConflict(opts *bind.CallOpts, rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "findNodeConflict", rollup, node1, node2, maxDepth)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsSession) FindNodeConflict(rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindNodeConflict(&_ValidatorUtils.CallOpts, rollup, node1, node2, maxDepth)
}

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCallerSession) FindNodeConflict(rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindNodeConflict(&_ValidatorUtils.CallOpts, rollup, node1, node2, maxDepth)
}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCaller) FindStakerConflict(opts *bind.CallOpts, rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "findStakerConflict", rollup, staker1, staker2, maxDepth)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsSession) FindStakerConflict(rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindStakerConflict(&_ValidatorUtils.CallOpts, rollup, staker1, staker2, maxDepth)
}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCallerSession) FindStakerConflict(rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindStakerConflict(&_ValidatorUtils.CallOpts, rollup, staker1, staker2, maxDepth)
}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsCaller) GetConfig(opts *bind.CallOpts, rollup common.Address) (struct {
	ChallengePeriodBlocks    *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "getConfig", rollup)

	outstruct := new(struct {
		ChallengePeriodBlocks    *big.Int
		ArbGasSpeedLimitPerBlock *big.Int
		BaseStake                *big.Int
		StakeToken               common.Address
	})

	outstruct.ChallengePeriodBlocks = out[0].(*big.Int)
	outstruct.ArbGasSpeedLimitPerBlock = out[1].(*big.Int)
	outstruct.BaseStake = out[2].(*big.Int)
	outstruct.StakeToken = out[3].(common.Address)

	return *outstruct, err

}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsSession) GetConfig(rollup common.Address) (struct {
	ChallengePeriodBlocks    *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}, error) {
	return _ValidatorUtils.Contract.GetConfig(&_ValidatorUtils.CallOpts, rollup)
}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsCallerSession) GetConfig(rollup common.Address) (struct {
	ChallengePeriodBlocks    *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}, error) {
	return _ValidatorUtils.Contract.GetConfig(&_ValidatorUtils.CallOpts, rollup)
}

// GetStakers is a free data retrieval call binding the contract method 0xabeba4f7.
//
// Solidity: function getStakers(address rollup, uint256 startIndex, uint256 max) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsCaller) GetStakers(opts *bind.CallOpts, rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "getStakers", rollup, startIndex, max)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakers is a free data retrieval call binding the contract method 0xabeba4f7.
//
// Solidity: function getStakers(address rollup, uint256 startIndex, uint256 max) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsSession) GetStakers(rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	return _ValidatorUtils.Contract.GetStakers(&_ValidatorUtils.CallOpts, rollup, startIndex, max)
}

// GetStakers is a free data retrieval call binding the contract method 0xabeba4f7.
//
// Solidity: function getStakers(address rollup, uint256 startIndex, uint256 max) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) GetStakers(rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, error) {
	return _ValidatorUtils.Contract.GetStakers(&_ValidatorUtils.CallOpts, rollup, startIndex, max)
}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsCaller) RefundableStakers(opts *bind.CallOpts, rollup common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "refundableStakers", rollup)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsSession) RefundableStakers(rollup common.Address) ([]common.Address, error) {
	return _ValidatorUtils.Contract.RefundableStakers(&_ValidatorUtils.CallOpts, rollup)
}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) RefundableStakers(rollup common.Address) ([]common.Address, error) {
	return _ValidatorUtils.Contract.RefundableStakers(&_ValidatorUtils.CallOpts, rollup)
}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCaller) StakedNodes(opts *bind.CallOpts, rollup common.Address, staker common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "stakedNodes", rollup, staker)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsSession) StakedNodes(rollup common.Address, staker common.Address) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.StakedNodes(&_ValidatorUtils.CallOpts, rollup, staker)
}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) StakedNodes(rollup common.Address, staker common.Address) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.StakedNodes(&_ValidatorUtils.CallOpts, rollup, staker)
}

// StakerInfo is a free data retrieval call binding the contract method 0x8f67e6bb.
//
// Solidity: function stakerInfo(address rollup, address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_ValidatorUtils *ValidatorUtilsCaller) StakerInfo(opts *bind.CallOpts, rollup common.Address, stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "stakerInfo", rollup, stakerAddress)

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

// StakerInfo is a free data retrieval call binding the contract method 0x8f67e6bb.
//
// Solidity: function stakerInfo(address rollup, address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_ValidatorUtils *ValidatorUtilsSession) StakerInfo(rollup common.Address, stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	return _ValidatorUtils.Contract.StakerInfo(&_ValidatorUtils.CallOpts, rollup, stakerAddress)
}

// StakerInfo is a free data retrieval call binding the contract method 0x8f67e6bb.
//
// Solidity: function stakerInfo(address rollup, address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_ValidatorUtils *ValidatorUtilsCallerSession) StakerInfo(rollup common.Address, stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	return _ValidatorUtils.Contract.StakerInfo(&_ValidatorUtils.CallOpts, rollup, stakerAddress)
}

// SuccessorNodes is a free data retrieval call binding the contract method 0x8730825e.
//
// Solidity: function successorNodes(address rollup, uint256 nodeNum) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCaller) SuccessorNodes(opts *bind.CallOpts, rollup common.Address, nodeNum *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "successorNodes", rollup, nodeNum)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SuccessorNodes is a free data retrieval call binding the contract method 0x8730825e.
//
// Solidity: function successorNodes(address rollup, uint256 nodeNum) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsSession) SuccessorNodes(rollup common.Address, nodeNum *big.Int) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.SuccessorNodes(&_ValidatorUtils.CallOpts, rollup, nodeNum)
}

// SuccessorNodes is a free data retrieval call binding the contract method 0x8730825e.
//
// Solidity: function successorNodes(address rollup, uint256 nodeNum) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) SuccessorNodes(rollup common.Address, nodeNum *big.Int) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.SuccessorNodes(&_ValidatorUtils.CallOpts, rollup, nodeNum)
}
