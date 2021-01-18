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
const INodeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"addStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"removeStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requirePastDeadline\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"latestConfirmed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"requireRejectExample\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// RequirePastDeadline is a free data retrieval call binding the contract method 0x88d221c6.
//
// Solidity: function requirePastDeadline() view returns()
func (_INode *INodeCaller) RequirePastDeadline(opts *bind.CallOpts) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "requirePastDeadline")

	if err != nil {
		return err
	}

	return err

}

// RequirePastDeadline is a free data retrieval call binding the contract method 0x88d221c6.
//
// Solidity: function requirePastDeadline() view returns()
func (_INode *INodeSession) RequirePastDeadline() error {
	return _INode.Contract.RequirePastDeadline(&_INode.CallOpts)
}

// RequirePastDeadline is a free data retrieval call binding the contract method 0x88d221c6.
//
// Solidity: function requirePastDeadline() view returns()
func (_INode *INodeCallerSession) RequirePastDeadline() error {
	return _INode.Contract.RequirePastDeadline(&_INode.CallOpts)
}

// RequireRejectExample is a free data retrieval call binding the contract method 0xfeb508ab.
//
// Solidity: function requireRejectExample(uint256 latestConfirmed, address stakerAddress) view returns()
func (_INode *INodeCaller) RequireRejectExample(opts *bind.CallOpts, latestConfirmed *big.Int, stakerAddress common.Address) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "requireRejectExample", latestConfirmed, stakerAddress)

	if err != nil {
		return err
	}

	return err

}

// RequireRejectExample is a free data retrieval call binding the contract method 0xfeb508ab.
//
// Solidity: function requireRejectExample(uint256 latestConfirmed, address stakerAddress) view returns()
func (_INode *INodeSession) RequireRejectExample(latestConfirmed *big.Int, stakerAddress common.Address) error {
	return _INode.Contract.RequireRejectExample(&_INode.CallOpts, latestConfirmed, stakerAddress)
}

// RequireRejectExample is a free data retrieval call binding the contract method 0xfeb508ab.
//
// Solidity: function requireRejectExample(uint256 latestConfirmed, address stakerAddress) view returns()
func (_INode *INodeCallerSession) RequireRejectExample(latestConfirmed *big.Int, stakerAddress common.Address) error {
	return _INode.Contract.RequireRejectExample(&_INode.CallOpts, latestConfirmed, stakerAddress)
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
const RollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"indexed\":false,\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"inboxMaxHash\",\"type\":\"bytes32\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"SentLogs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"addToDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLatestNodeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"beginTruncatingNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"logAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"continueTruncatingNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[3]\",\"name\":\"nodeFields\",\"type\":\"bytes32[3]\"},{\"internalType\":\"uint256\",\"name\":\"gasClaimed\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollupEventBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"newStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"successorWithStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireNoRecentStake\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"requireUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireUnresolvedExists\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[7]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[7]\"},{\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"}],\"name\":\"upgradeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeImplementationAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupBin is the compiled bytecode used for deploying new contracts.
var RollupBin = "0x608060405234801561001057600080fd5b506009805460ff19169055614e248061002a6000396000f3fe6080604052600436106102865760003560e01c8063046f7da21461028b57806304a28064146102a25780630e1ef04c146102e75780630e21b5ba146103205780631e83d30f1461034a5780632b2af0ab146103745780632f30cabd1461039e578063396b8cbc146103d15780633b333ea8146104a35780633e96576e146104d357806345e38b641461050657806346c2781a1461051b5780634d26732d146105305780634f0f4aa91461054557806351ed6a301461058b578063567ca41b146105a05780635c975abb146105d35780635dbaf68b146105fc5780635e8ef106146106115780635f81e5fc146106265780636177fd181461074157806362a82d7d1461077457806363721d6b1461079e57806365f7f80d146107b357806367425daf146107c857806369fd251c146107dd5780637427be511461081057806376e7e23b146108435780637ba9534a146108585780637e2d21551461086d57806381fbc98a1461089d57806383f94db7146108d05780638456cb59146109035780638640ce5f146109185780638da5cb5b1461092d5780638fd18f04146109425780639e8a713f14610978578063b15a38df1461098d578063b1b4181c146109a2578063bf5ddcb1146109d9578063ce11e6ab14610a64578063d01e660214610a79578063d735e21d14610aa3578063d93fe9c414610ab8578063dff6978714610acd578063e45b7ce614610ae2578063e78cea9214610b1d578063e8bd492214610b32578063edfd03ed14610b9b578063ef40a67014610bc5578063f019a1c114610bf8578063f33e1fac14610c3a578063f3f0a03e14610c64578063f851a44014610c90578063fa7803e614610ca5578063fb64884e14610ce0578063ff204f3b14610cfd575b600080fd5b34801561029757600080fd5b506102a0610d30565b005b3480156102ae57600080fd5b506102d5600480360360208110156102c557600080fd5b50356001600160a01b0316610dd1565b60408051918252519081900360200190f35b3480156102f357600080fd5b506102a06004803603604081101561030a57600080fd5b50803590602001356001600160a01b0316610e89565b34801561032c57600080fd5b506102a06004803603602081101561034357600080fd5b50356111b2565b34801561035657600080fd5b506102a06004803603602081101561036d57600080fd5b503561145e565b34801561038057600080fd5b506102a06004803603602081101561039757600080fd5b50356114d4565b3480156103aa57600080fd5b506102d5600480360360208110156103c157600080fd5b50356001600160a01b031661156d565b3480156103dd57600080fd5b506102a0600480360360608110156103f457600080fd5b81359190810190604081016020820135600160201b81111561041557600080fd5b82018360208201111561042757600080fd5b803590602001918460018302840111600160201b8311171561044857600080fd5b919390929091602081019035600160201b81111561046557600080fd5b82018360208201111561047757600080fd5b803590602001918460208302840111600160201b8311171561049857600080fd5b509092509050611588565b3480156104af57600080fd5b506102a0600480360360408110156104c657600080fd5b5080359060200135611a94565b3480156104df57600080fd5b506102d5600480360360208110156104f657600080fd5b50356001600160a01b0316611c28565b34801561051257600080fd5b506102d5611c46565b34801561052757600080fd5b506102d5611c5a565b34801561053c57600080fd5b506102d5611c60565b34801561055157600080fd5b5061056f6004803603602081101561056857600080fd5b5035611d83565b604080516001600160a01b039092168252519081900360200190f35b34801561059757600080fd5b5061056f611d9e565b3480156105ac57600080fd5b506102a0600480360360208110156105c357600080fd5b50356001600160a01b0316611dad565b3480156105df57600080fd5b506105e8611eb0565b604080519115158252519081900360200190f35b34801561060857600080fd5b5061056f611eb9565b34801561061d57600080fd5b506102d5611ec8565b34801561063257600080fd5b506102a060048036036101a081101561064a57600080fd5b6001600160a01b038235811692602081013582169260408201359260608301359260808101359260a08201359260c083013581169260e0810135821692610100820135831692610120830135811692610140810135909116918101906101808101610160820135600160201b8111156106c257600080fd5b8201836020820111156106d457600080fd5b803590602001918460018302840111600160201b831117156106f557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150611ece9050565b34801561074d57600080fd5b506105e86004803603602081101561076457600080fd5b50356001600160a01b03166122da565b34801561078057600080fd5b5061056f6004803603602081101561079757600080fd5b5035612302565b3480156107aa57600080fd5b506102d561232c565b3480156107bf57600080fd5b506102d5612332565b3480156107d457600080fd5b506102a0612338565b3480156107e957600080fd5b5061056f6004803603602081101561080057600080fd5b50356001600160a01b03166123a2565b34801561081c57600080fd5b506102a06004803603602081101561083357600080fd5b50356001600160a01b03166123c3565b34801561084f57600080fd5b506102d561246c565b34801561086457600080fd5b506102d5612472565b34801561087957600080fd5b506102a06004803603604081101561089057600080fd5b5080359060200135612478565b3480156108a957600080fd5b506102d5600480360360208110156108c057600080fd5b50356001600160a01b0316612656565b3480156108dc57600080fd5b506102a0600480360360208110156108f357600080fd5b50356001600160a01b03166126bc565b34801561090f57600080fd5b506102a0612775565b34801561092457600080fd5b506102d56127c9565b34801561093957600080fd5b5061056f6127cf565b34801561094e57600080fd5b506102a06004803603606081101561096557600080fd5b50803590602081013590604001356127de565b34801561098457600080fd5b5061056f6129ac565b34801561099957600080fd5b506102a06129bb565b3480156109ae57600080fd5b506102a060048036036101008110156109c657600080fd5b50604081016080820160e0830135612a0a565b3480156109e557600080fd5b506102a0600480360360408110156109fc57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b811115610a2657600080fd5b820183602082011115610a3857600080fd5b803590602001918460018302840111600160201b83111715610a5957600080fd5b509092509050613007565b348015610a7057600080fd5b5061056f6130fd565b348015610a8557600080fd5b5061056f60048036036020811015610a9c57600080fd5b503561310c565b348015610aaf57600080fd5b506102d561313b565b348015610ac457600080fd5b5061056f613141565b348015610ad957600080fd5b506102d5613150565b348015610aee57600080fd5b506102a060048036036040811015610b0557600080fd5b506001600160a01b0381351690602001351515613156565b348015610b2957600080fd5b5061056f6131f8565b348015610b3e57600080fd5b50610b6560048036036020811015610b5557600080fd5b50356001600160a01b0316613207565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b348015610ba757600080fd5b506102a060048036036020811015610bbe57600080fd5b5035613243565b348015610bd157600080fd5b506102d560048036036020811015610be857600080fd5b50356001600160a01b03166132a2565b348015610c0457600080fd5b506102a06004803603610280811015610c1c57600080fd5b508035906020810135906040810135906060810190610140016132c0565b348015610c4657600080fd5b506102d560048036036020811015610c5d57600080fd5b5035613961565b6102a060048036036040811015610c7a57600080fd5b506001600160a01b038135169060200135613989565b348015610c9c57600080fd5b5061056f6139ea565b348015610cb157600080fd5b506102a060048036036040811015610cc857600080fd5b506001600160a01b03813581169160200135166139f9565b6102a060048036036020811015610cf657600080fd5b5035613ac3565b348015610d0957600080fd5b506102a060048036036020811015610d2057600080fd5b50356001600160a01b0316613c1b565b6013546001600160a01b03163314610d7c576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b60175460ff1615610dc7576040805162461bcd60e51b815260206004820152601060248201526f5354494c4c5f5452554e434154494e4760801b604482015290519081900360640190fd5b610dcf613cd1565b565b600080610ddc61232c565b90506000805b82811015610e7f57846001600160a01b0316639168ae72610e028361310c565b6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610e3f57600080fd5b505afa158015610e53573d6000803e3d6000fd5b505050506040513d6020811015610e6957600080fd5b505115610e77576001909101905b600101610de2565b509150505b919050565b60095460ff1615610ecf576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b610ed7612338565b6000610ee1612332565b90506000610eed61313b565b90506000610efa82611d83565b905082816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610f3657600080fd5b505afa158015610f4a573d6000803e3d6000fd5b505050506040513d6020811015610f6057600080fd5b5051141561113e57610f706129bb565b610f79856114d4565b610f82846122da565b610fc0576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b610fc985611d83565b6001600160a01b031663feb508ab84866040518363ffffffff1660e01b815260040180838152602001826001600160a01b031681526020019250505060006040518083038186803b15801561101d57600080fd5b505afa158015611031573d6000803e3d6000fd5b50505050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b15801561106e57600080fd5b505afa158015611082573d6000803e3d6000fd5b505050506110906000613243565b61109981610dd1565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b1580156110d257600080fd5b505afa1580156110e6573d6000803e3d6000fd5b505050506040513d60208110156110fc57600080fd5b50511461113e576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b611146613d6f565b60105460408051630c2a09ad60e21b81526004810185905290516001600160a01b03909216916330a826b49160248082019260009290919082900301818387803b15801561119357600080fd5b505af11580156111a7573d6000803e3d6000fd5b505050505050505050565b6013546001600160a01b031633146111fe576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b60095460ff1661124c576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b60175460ff16611294576040805162461bcd60e51b815260206004820152600e60248201526d4e4f545f5452554e434154494e4760901b604482015290519081900360640190fd5b60155460165460006112a4613150565b90505b6000841180156112b657508082105b156113975760006112c683612302565b905060006112d382611c28565b90505b6000861180156112e557508481115b1561136c5760006112f582611d83565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561133057600080fd5b505afa158015611344573d6000803e3d6000fd5b505050506040513d602081101561135a57600080fd5b50516000199097019691506112d69050565b6113768282613d84565b8481111561138a575050506016555061145b565b50506001909101906112a7565b601682905560006113a6612472565b90505b6000851180156113b857508381115b156114305760006113c882611d83565b9050806001600160a01b03166383197ef06040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561140557600080fd5b505af1158015611419573d6000803e3d6000fd5b5050600019978801979390930192506113a9915050565b61143981613da3565b8381141561145657600060158190556016556017805460ff191690555b505050505b50565b60095460ff16156114a4576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b6114ad33613da8565b60006114b7611c60565b9050808210156114c5578091505b6114cf3383613e3f565b505050565b6114dc61313b565b811015611522576040805162461bcd60e51b815260206004820152600f60248201526e1053149150511657d11150d2511151608a1b604482015290519081900360640190fd5b61152a612472565b81111561145b576040805162461bcd60e51b815260206004820152600c60248201526b1113d154d39517d1561254d560a21b604482015290519081900360640190fd5b6001600160a01b031660009081526008602052604090205490565b60095460ff16156115ce576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b6115d6612338565b6115de6129bb565b60006115e8613150565b11611627576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f5354414b45525360b01b604482015290519081900360640190fd5b600061163161313b565b9050600061163e82611d83565b9050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b15801561167957600080fd5b505afa15801561168d573d6000803e3d6000fd5b50505050611699612332565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156116d257600080fd5b505afa1580156116e6573d6000803e3d6000fd5b505050506040513d60208110156116fc57600080fd5b50511461173f576040805162461bcd60e51b815260206004820152600c60248201526b24a72b20a624a22fa82922ab60a11b604482015290519081900360640190fd5b6117496000613243565b61175281610dd1565b61175a613150565b01816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561179457600080fd5b505afa1580156117a8573d6000803e3d6000fd5b505050506040513d60208110156117be57600080fd5b505114611803576040805162461bcd60e51b815260206004820152600e60248201526d1393d517d0531317d4d51052d15160921b604482015290519081900360640190fd5b600061187587878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808b0282810182019093528a82529093508a925089918291850190849080828437600092019190915250613ed792505050565b90506118818189613fd7565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b1580156118ba57600080fd5b505afa1580156118ce573d6000803e3d6000fd5b505050506040513d60208110156118e457600080fd5b505114611927576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b600f5460408051630c72684760e01b815260048101918252604481018990526001600160a01b0390921691630c726847918a918a918a918a919081906024810190606401878780828437600083820152601f01601f19169091018481038352858152602090810191508690860280828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b1580156119d157600080fd5b505af11580156119e5573d6000803e3d6000fd5b505050506119f1614003565b601054604080516316b9109b60e01b81526004810186905290516001600160a01b03909216916316b9109b9160248082019260009290919082900301818387803b158015611a3e57600080fd5b505af1158015611a52573d6000803e3d6000fd5b5050604080518b815290517fe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e849350908190036020019150a15050505050505050565b6013546001600160a01b03163314611ae0576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b60095460ff16611b2e576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b60175460ff1615611b7b576040805162461bcd60e51b8152602060048201526012602482015271414c52454144595f5452554e434154494e4760701b604482015290519081900360640190fd5b611b83612472565b8210611bc0576040805162461bcd60e51b8152602060048201526007602482015266544f4f5f4e455760c81b604482015290519081900360640190fd5b6001611bca61313b565b03821015611c09576040805162461bcd60e51b81526020600482015260076024820152661513d3d7d3d31160ca1b604482015290519081900360640190fd5b60158290556017805460ff19166001179055611c24816111b2565b5050565b6001600160a01b031660009081526006602052604090206001015490565b6000600a805481611c5357fe5b0490505b90565b600a5481565b600080611c6b61313b565b9050611c75612472565b600182031415611c89575050600c54611c57565b6000611c9482611d83565b9050600060001990506000826001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611cd857600080fd5b505afa158015611cec573d6000803e3d6000fd5b505050506040513d6020811015611d0257600080fd5b5051905043811115611d1c57600c54945050505050611c57565b600081430390506000600a548281611d3057fe5b04905060ff811115611d40575060ff5b600019600282900a0180611d52575060015b600c548581611d5d57fe5b04811115611d745784975050505050505050611c57565b600c5402965050505050505090565b6000908152600460205260409020546001600160a01b031690565b600d546001600160a01b031681565b6013546001600160a01b03163314611df9576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b600f546001600160a01b0382811691161415611e49576040805162461bcd60e51b815260206004820152600a602482015269086aaa4be9eaaa8849eb60b31b604482015290519081900360640190fd5b600e54604080516319dc7ae560e31b81526001600160a01b038481166004830152600060248301819052925193169263cee3d7289260448084019391929182900301818387803b158015611e9c57600080fd5b505af1158015611456573d6000803e3d6000fd5b60095460ff1690565b6011546001600160a01b031681565b600b5481565b84600e60006101000a8154816001600160a01b0302191690836001600160a01b031602179055508b601060006101000a8154816001600160a01b0302191690836001600160a01b031602179055508c600f60006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600e60009054906101000a90046001600160a01b03166001600160a01b031663cee3d7288e60016040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015611fb257600080fd5b505af1158015611fc6573d6000803e3d6000fd5b50505050600e60009054906101000a90046001600160a01b03166001600160a01b031663e45b7ce68d60016040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b15801561203957600080fd5b505af115801561204d573d6000803e3d6000fd5b50505050601060009054906101000a90046001600160a01b03166001600160a01b031662fc32f78b8b8b8b8b886040518763ffffffff1660e01b815260040180878152602001868152602001858152602001846001600160a01b03168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156120f65781810151838201526020016120de565b50505050905090810190601f1680156121235780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b15801561214757600080fd5b505af115801561215b573d6000803e3d6000fd5b5050601180546001600160a01b038089166001600160a01b031992831617909255601280549288169290911691909117905550600090506121a343828e81808080600161401d565b6012546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905260848201819052915193945090926001600160a01b039092169163d45ab2b59160a48082019260209290919082900301818787803b15801561221357600080fd5b505af1158015612227573d6000803e3d6000fd5b505050506040513d602081101561223d57600080fd5b5051905061224a81614077565b600a8c9055600b8b9055600c8a9055600d80546001600160a01b03808c166001600160a01b031992831617909255601380548b84169083161790556014805492861692909116919091179055604080518e815290517f4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d9181900360200190a1505050505050505050505050505050565b6001600160a01b0316600090815260066020526040902060030154600160a01b900460ff1690565b60006005828154811061231157fe5b6000918252602090912001546001600160a01b031692915050565b60075490565b60005490565b600061234261313b565b905061234c612332565b81118015612361575061235d612472565b8111155b61145b576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b6001600160a01b039081166000908152600660205260409020600301541690565b60095460ff1615612409576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b612411612332565b61241a82611c28565b111561245a576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b61246381613da8565b61145b816140c5565b600c5481565b60025490565b60095460ff16156124be576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b6124c661232c565b82111561250b576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b60006125168361310c565b9050600061252384613961565b905060008061253061313b565b90505b80831015801561254257508482105b1561262e57600061255284611d83565b9050806001600160a01b03166396a9fdc0866040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156125a357600080fd5b505af11580156125b7573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156125f457600080fd5b505afa158015612608573d6000803e3d6000fd5b505050506040513d602081101561261e57600080fd5b5051935050600190910190612533565b808310156126445761263f866140fc565b61264e565b61264e8684614198565b505050505050565b60095460009060ff161561269f576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b60006126aa336141bf565b90506126b683826141de565b92915050565b6013546001600160a01b03163314612708576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6014546040805163266a23b160e21b815230600482018190526001600160a01b0385811660248401529251909392909216916399a88ec49160448082019260009290919082900301818387803b15801561276157600080fd5b505af115801561264e573d6000803e3d6000fd5b6013546001600160a01b031633146127c1576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610dcf6142f9565b60035490565b6013546001600160a01b031681565b60095460ff1615612824576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b61282d336122da565b61286b576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b828240146128b6576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b6128be61313b565b81101580156128d457506128d0612472565b8111155b6128dd57600080fd5b60006128e882611d83565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561292357600080fd5b505afa158015612937573d6000803e3d6000fd5b505050506040513d602081101561294d57600080fd5b505161295833611c28565b1461299c576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b6129a63383614375565b50505050565b6010546001600160a01b031681565b600a546129c66127c9565b43031015610dcf576040805162461bcd60e51b815260206004820152600c60248201526b524543454e545f5354414b4560a01b604482015290519081900360640190fd5b60095460ff1615612a50576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b6020830135833510612a97576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b612a9f612472565b60208401351115612ae6576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b8235612af0612332565b10612b36576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b6000612b4884825b6020020135611d83565b90506000612b57856001612b3e565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015612b9257600080fd5b505afa158015612ba6573d6000803e3d6000fd5b505050506040513d6020811015612bbc57600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b158015612bfe57600080fd5b505afa158015612c12573d6000803e3d6000fd5b505050506040513d6020811015612c2857600080fd5b505114612c68576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b612c828660005b60200201356001600160a01b0316613da8565b612c8d866001612c6f565b604080516348b4573960e11b81526001600160a01b03883581166004830152915191841691639168ae7291602480820192602092909190829003018186803b158015612cd857600080fd5b505afa158015612cec573d6000803e3d6000fd5b505050506040513d6020811015612d0257600080fd5b5051612d4a576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b604080516348b4573960e11b81526001600160a01b0360208981013582166004840152925190841692639168ae729260248082019391829003018186803b158015612d9457600080fd5b505afa158015612da8573d6000803e3d6000fd5b505050506040513d6020811015612dbe57600080fd5b5051612e06576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b612e1b843560208601356040870135866143fe565b826001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b158015612e5457600080fd5b505afa158015612e68573d6000803e3d6000fd5b505050506040513d6020811015612e7e57600080fd5b505114612ebe576040805162461bcd60e51b8152602060048201526009602482015268086908298be9082a6960bb1b604482015290519081900360640190fd5b601154600b54600a546040805163471cfcbd60e01b8152306004820152883560248201526020808a013560448301528983013560648301526084820189905260a48201949094526001600160a01b038b35811660c48301528b850135811660e48301526101048201939093529051600094929092169263471cfcbd926101248084019382900301818787803b158015612f5657600080fd5b505af1158015612f6a573d6000803e3d6000fd5b505050506040513d6020811015612f8057600080fd5b50519050612fa96001600160a01b0388351688600160200201356001600160a01b03168361443c565b6040805188356001600160a01b0390811682526020808b01358216908301528835828401529151918316917fa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda758799181900360600190a250505050505050565b6013546001600160a01b03163314613053576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601454604051639623609d60e01b815230600482018181526001600160a01b0387811660248501526060604485019081526064850187905292941692639623609d92859289928992899291608401848480828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b1580156130df57600080fd5b505af11580156130f3573d6000803e3d6000fd5b5050505050505050565b600f546001600160a01b031681565b60006007828154811061311b57fe5b60009182526020909120600290910201546001600160a01b031692915050565b60015490565b6012546001600160a01b031681565b60055490565b6013546001600160a01b031633146131a2576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b600e546040805163722dbe7360e11b81526001600160a01b03858116600483015284151560248301529151919092169163e45b7ce691604480830192600092919082900301818387803b15801561276157600080fd5b600e546001600160a01b031681565b6006602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b600061324d61232c565b9050600061325961313b565b9050825b828110156129a6575b8161327082613961565b101561329a5761327f816140fc565b600019909201918281106132955750505061145b565b613266565b60010161325d565b6001600160a01b031660009081526006602052604090206002015490565b60095460ff1615613306576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b61330f336122da565b61334d576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b84844014613398576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b6133a0612472565b60010183146133e1576040805162461bcd60e51b81526020600482015260086024820152674e4f44455f4e554d60c01b604482015290519081900360640190fd5b6133e9614d42565b6040805160e081810190925261343d9185906007908390839080828437600092019190915250506040805161014081810190925291508590600a908390839080828437600092019190915250614486915050565b9050600061345261344d33611c28565b611d83565b9050806001600160a01b031663701da98e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561348d57600080fd5b505afa1580156134a1573d6000803e3d6000fd5b505050506040513d60208110156134b757600080fd5b50516134c2836145a5565b14613506576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b81514303613512611c46565b811015613553576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b82608001518360e001510383610120015110158061357a5750600b54810283610140015110155b8061358a57506101808301516064145b6135c7576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b600b5481026004028361014001511115613614576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b6000600a54430190506000836001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561365857600080fd5b505afa15801561366c573d6000803e3d6000fd5b505050506040513d602081101561368257600080fd5b5051905080821015613692578091505b600b54856101400151816136a257fe5b60105491900492909201916001600160a01b0316638b8ca199896136c533611c28565b85336040518563ffffffff1660e01b815260040180858152602001848152602001838152602001826001600160a01b03168152602001945050505050600060405180830381600087803b15801561371b57600080fd5b505af115801561372f573d6000803e3d6000fd5b50505050600080600e60009054906101000a90046001600160a01b03166001600160a01b0316630c301fb26040518163ffffffff1660e01b8152600401604080518083038186803b15801561378357600080fd5b505afa158015613797573d6000803e3d6000fd5b505050506040513d60408110156137ad57600080fd5b50805160209091015160808901516101208a01519294509092508303101561380d576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b6012546000906001600160a01b031663d45ab2b561382b8a866145d7565b6138368b8787614626565b61383f8c61469c565b61384833611c28565b8a6040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b15801561389b57600080fd5b505af11580156138af573d6000803e3d6000fd5b505050506040513d60208110156138c557600080fd5b505190506138d2816146b2565b6138dc338c614375565b8a7f4807480f255627d9b9350200277cb372949a76ccb4263935665257e2b3a1582a8b8b86866040518085600760200280828437600083820152601f01601f191690910190508461014080828437600083820152601f01601f19169091019384525050602082015260408051918290030192509050a250505050505050505050505050565b60006007828154811061397057fe5b9060005260206000209060020201600101549050919050565b60095460ff16156139cf576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b6139d882613da8565b611c24826139e5836146eb565b614859565b6014546001600160a01b031681565b613a03828261487e565b6001600160a01b0316336001600160a01b031614613a57576040805162461bcd60e51b815260206004820152600c60248201526b2ba927a723afa9a2a72222a960a11b604482015290519081900360640190fd5b6000613a62826132a2565b90506000613a6f846132a2565b905080821115613a8857613a838382613e3f565b820391505b60028204613a968582614859565b8083039250613aa485614940565b601354613aba906001600160a01b031684614859565b6114568461496a565b60095460ff1615613b09576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b613b12336122da565b15613b55576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b6000613b60826146eb565b9050613b6a611c60565b811015613bb1576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b613bbb3382614a1a565b6010546001600160a01b031663f03c04a533613bd5612332565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561276157600080fd5b6013546001600160a01b03163314613c67576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b600f80546001600160a01b0319166001600160a01b03838116918217909255600e54604080516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b158015611e9c57600080fd5b60095460ff16613d1f576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6009805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613d52614ae4565b604080516001600160a01b039092168252519081900360200190a1565b613d7a600154614ae8565b6001805481019055565b6001600160a01b03909116600090815260066020526040902060010155565b600255565b613db1816122da565b613def576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6000613dfa826123a2565b6001600160a01b03161461145b576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600660205260408120600281015480841115613ea4576040805162461bcd60e51b815260206004820152601060248201526f544f4f5f4c4954544c455f5354414b4560801b604482015290519081900360640190fd5b60029091018390556001600160a01b03841660009081526008602052604090208054918490039182019055905092915050565b80518251600091829182805b83811015613f8a576000878281518110613ef957fe5b60200260200101519050838187011115613f49576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868a0181018290206040805180840196909652858101919091528051808603820181526060909501905283519301929092209190940193600101613ee3565b50818414613fcd576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b9695505050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b61400e600054614ae8565b60018054600081905581019055565b6040805160208082019a909a52808201989098526060880196909652608087019490945260a086019290925260c085015260e084015261010080840191909152815180840390910181526101209092019052805191012090565b6000805260046020527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec80546001600160a01b0319166001600160a01b039290921691909117905560018055565b6001600160a01b038116600090815260066020908152604080832060028101546008909352922080549091019055611c2482614b6c565b60078054600019810190811061410e57fe5b90600052602060002090600202016007828154811061412957fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600780548061416c57fe5b60008281526020812060026000199093019283020180546001600160a01b031916815560010155905550565b80600783815481106141a657fe5b9060005260206000209060020201600101819055505050565b6001600160a01b03166000908152600860205260408120805491905590565b806141e857611c24565b600d546001600160a01b0316614234576040516001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561422e573d6000803e3d6000fd5b50611c24565b600d546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561428a57600080fd5b505af115801561429e573d6000803e3d6000fd5b505050506040513d60208110156142b457600080fd5b5051611c24576040805162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b604482015290519081900360640190fd5b60095460ff161561433f576040805162461bcd60e51b81526020600482015260106024820152600080516020614dcf833981519152604482015290519081900360640190fd5b6009805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613d52614ae4565b6001600160a01b03808316600081815260066020908152604080832086845260049283905281842054825163123334b760e11b8152938401959095529051909493909316928392632466696e926024808201939182900301818387803b1580156143de57600080fd5b505af11580156143f2573d6000803e3d6000fd5b50505050506001015550565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b6001600160a01b03928316600090815260066020526040808220600390810180549487166001600160a01b0319958616811790915594909516825290209092018054909216179055565b61448e614d42565b60408051610220810182528351815260208085015181830152855182840152850151606080830191909152848301516080808401919091529085015160a0808401919091529085015160c083015284015160e082015290840151610100820152610120810183600660200201518152602001836007600a811061450d57fe5b602002015181526020018460036007811061452457fe5b60200201518152602001836008600a811061453b57fe5b602002015181526020018460046007811061455257fe5b60200201518152602001836009600a811061456957fe5b602002015181526020018460056007811061458057fe5b602002015181526020018460066007811061459757fe5b602002015190529392505050565b60006126b6826000015183602001518460400151856060015186608001518760a001518860c001518960e0015161401d565b600061461f43846101400151856020015101856102000151866101e001518761012001518860800151018861018001518960a0015101896101c001518a60c00151018961401d565b9392505050565b6000806146486000866101200151876080015187030385886101e001516143fe565b905060006146816000876101200151614669896101e001516000801b613fd7565b61467c8a606001518b6101000151613fd7565b6143fe565b9050613fcd828261469189614c92565b8961014001516143fe565b60006126b6826101600151836101a00151613fd7565b6002805460010190819055600090815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b600d546000906001600160a01b0316614747578115614740576040805162461bcd60e51b815260206004820152600c60248201526b4241445f53544b5f5459504560a01b604482015290519081900360640190fd5b5034610e84565b3415614789576040805162461bcd60e51b815260206004820152600c60248201526b4241445f53544b5f5459504560a01b604482015290519081900360640190fd5b600d54604080516323b872dd60e01b81523360048201523060248201526044810185905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156147e357600080fd5b505af11580156147f7573d6000803e3d6000fd5b505050506040513d602081101561480d57600080fd5b5051614852576040805162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b604482015290519081900360640190fd5b5080610e84565b6001600160a01b03909116600090815260066020526040902060020180549091019055565b6001600160a01b03808316600090815260066020526040808220848416835290822060038083015490820154939492939192908116911681146148f2576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b038116614937576040805162461bcd60e51b81526020600482015260076024820152661393d7d0d2105360ca1b604482015290519081900360640190fd5b95945050505050565b6001600160a01b0316600090815260066020526040902060030180546001600160a01b0319169055565b6001600160a01b0381811660008181526006602090815260408083208151808301909252938152600180850154928201928352600780549182018155909352517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688600290930292830180546001600160a01b031916919095161790935591517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68990920191909155611c2482614b6c565b6005805460018082019092557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0810180546001600160a01b03199081166001600160a01b039687169081179092556040805160a08101825293845260008054602080870191825286840198895260608701838152608088018981529684526006909152929091209451855551948401949094559351600283015591516003918201805493519390941694169390931760ff60a01b1916600160a01b91151591909102179055439055565b3390565b600081815260046020819052604080832054815163083197ef60e41b815291516001600160a01b03909116936383197ef0938381019391929182900301818387803b158015614b3657600080fd5b505af1158015614b4a573d6000803e3d6000fd5b50505060009182525060046020526040902080546001600160a01b0319169055565b6001600160a01b03811660009081526006602052604090208054600580546000198101908110614b9857fe5b600091825260209091200154600580546001600160a01b039092169183908110614bbe57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806006600060058481548110614bfe57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556005805480614c2e57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03949094168152600690935250506040812081815560018101829055600281019190915560030180546001600160a81b0319169055565b60006126b66000836101400151614cc86000614cc387610100015188604001516000801b60008060001b6000614cf7565b613fd7565b61467c866101400151614cc36000801b8961020001518a61016001518b61018001518c6101a001518d6101c001515b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081018290526102008101919091529056fe5061757361626c653a2070617573656400000000000000000000000000000000a26469706673582212205cf78390c2189df3d9561a671bf52ec7ffc02b55ce56d74364322e823627fe8564736f6c634300060c0033"

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

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Rollup *RollupCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Rollup *RollupSession) Admin() (common.Address, error) {
	return _Rollup.Contract.Admin(&_Rollup.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Rollup *RollupCallerSession) Admin() (common.Address, error) {
	return _Rollup.Contract.Admin(&_Rollup.CallOpts)
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
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_Rollup *RollupCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_Rollup *RollupSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetNode(&_Rollup.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_Rollup *RollupCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetNode(&_Rollup.CallOpts, nodeNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_Rollup *RollupCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_Rollup *RollupSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetStakerAddress(&_Rollup.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_Rollup *RollupCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetStakerAddress(&_Rollup.CallOpts, stakerNum)
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

// RequireNoRecentStake is a free data retrieval call binding the contract method 0xb15a38df.
//
// Solidity: function requireNoRecentStake() view returns()
func (_Rollup *RollupCaller) RequireNoRecentStake(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "requireNoRecentStake")

	if err != nil {
		return err
	}

	return err

}

// RequireNoRecentStake is a free data retrieval call binding the contract method 0xb15a38df.
//
// Solidity: function requireNoRecentStake() view returns()
func (_Rollup *RollupSession) RequireNoRecentStake() error {
	return _Rollup.Contract.RequireNoRecentStake(&_Rollup.CallOpts)
}

// RequireNoRecentStake is a free data retrieval call binding the contract method 0xb15a38df.
//
// Solidity: function requireNoRecentStake() view returns()
func (_Rollup *RollupCallerSession) RequireNoRecentStake() error {
	return _Rollup.Contract.RequireNoRecentStake(&_Rollup.CallOpts)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_Rollup *RollupCaller) RequireUnresolved(opts *bind.CallOpts, nodeNum *big.Int) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "requireUnresolved", nodeNum)

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_Rollup *RollupSession) RequireUnresolved(nodeNum *big.Int) error {
	return _Rollup.Contract.RequireUnresolved(&_Rollup.CallOpts, nodeNum)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_Rollup *RollupCallerSession) RequireUnresolved(nodeNum *big.Int) error {
	return _Rollup.Contract.RequireUnresolved(&_Rollup.CallOpts, nodeNum)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_Rollup *RollupCaller) RequireUnresolvedExists(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "requireUnresolvedExists")

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_Rollup *RollupSession) RequireUnresolvedExists() error {
	return _Rollup.Contract.RequireUnresolvedExists(&_Rollup.CallOpts)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_Rollup *RollupCallerSession) RequireUnresolvedExists() error {
	return _Rollup.Contract.RequireUnresolvedExists(&_Rollup.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_Rollup *RollupCaller) RollupEventBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "rollupEventBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_Rollup *RollupSession) RollupEventBridge() (common.Address, error) {
	return _Rollup.Contract.RollupEventBridge(&_Rollup.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_Rollup *RollupCallerSession) RollupEventBridge() (common.Address, error) {
	return _Rollup.Contract.RollupEventBridge(&_Rollup.CallOpts)
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
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_Rollup *RollupCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_Rollup *RollupSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.ZombieAddress(&_Rollup.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_Rollup *RollupCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.ZombieAddress(&_Rollup.CallOpts, zombieNum)
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
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_Rollup *RollupCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_Rollup *RollupSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _Rollup.Contract.ZombieLatestStakedNode(&_Rollup.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_Rollup *RollupCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _Rollup.Contract.ZombieLatestStakedNode(&_Rollup.CallOpts, zombieNum)
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

// BeginTruncatingNodes is a paid mutator transaction binding the contract method 0x3b333ea8.
//
// Solidity: function beginTruncatingNodes(uint256 newLatestNodeCreated, uint256 maxItems) returns()
func (_Rollup *RollupTransactor) BeginTruncatingNodes(opts *bind.TransactOpts, newLatestNodeCreated *big.Int, maxItems *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "beginTruncatingNodes", newLatestNodeCreated, maxItems)
}

// BeginTruncatingNodes is a paid mutator transaction binding the contract method 0x3b333ea8.
//
// Solidity: function beginTruncatingNodes(uint256 newLatestNodeCreated, uint256 maxItems) returns()
func (_Rollup *RollupSession) BeginTruncatingNodes(newLatestNodeCreated *big.Int, maxItems *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.BeginTruncatingNodes(&_Rollup.TransactOpts, newLatestNodeCreated, maxItems)
}

// BeginTruncatingNodes is a paid mutator transaction binding the contract method 0x3b333ea8.
//
// Solidity: function beginTruncatingNodes(uint256 newLatestNodeCreated, uint256 maxItems) returns()
func (_Rollup *RollupTransactorSession) BeginTruncatingNodes(newLatestNodeCreated *big.Int, maxItems *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.BeginTruncatingNodes(&_Rollup.TransactOpts, newLatestNodeCreated, maxItems)
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

// ContinueTruncatingNodes is a paid mutator transaction binding the contract method 0x0e21b5ba.
//
// Solidity: function continueTruncatingNodes(uint256 maxItems) returns()
func (_Rollup *RollupTransactor) ContinueTruncatingNodes(opts *bind.TransactOpts, maxItems *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "continueTruncatingNodes", maxItems)
}

// ContinueTruncatingNodes is a paid mutator transaction binding the contract method 0x0e21b5ba.
//
// Solidity: function continueTruncatingNodes(uint256 maxItems) returns()
func (_Rollup *RollupSession) ContinueTruncatingNodes(maxItems *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ContinueTruncatingNodes(&_Rollup.TransactOpts, maxItems)
}

// ContinueTruncatingNodes is a paid mutator transaction binding the contract method 0x0e21b5ba.
//
// Solidity: function continueTruncatingNodes(uint256 maxItems) returns()
func (_Rollup *RollupTransactorSession) ContinueTruncatingNodes(maxItems *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ContinueTruncatingNodes(&_Rollup.TransactOpts, maxItems)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xb1b4181c.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[3] nodeFields, uint256 gasClaimed) returns()
func (_Rollup *RollupTransactor) CreateChallenge(opts *bind.TransactOpts, stakers [2]common.Address, nodeNums [2]*big.Int, nodeFields [3][32]byte, gasClaimed *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "createChallenge", stakers, nodeNums, nodeFields, gasClaimed)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xb1b4181c.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[3] nodeFields, uint256 gasClaimed) returns()
func (_Rollup *RollupSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, nodeFields [3][32]byte, gasClaimed *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, stakers, nodeNums, nodeFields, gasClaimed)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xb1b4181c.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[3] nodeFields, uint256 gasClaimed) returns()
func (_Rollup *RollupTransactorSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, nodeFields [3][32]byte, gasClaimed *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, stakers, nodeNums, nodeFields, gasClaimed)
}

// Initialize is a paid mutator transaction binding the contract method 0x5f81e5fc.
//
// Solidity: function initialize(address _outbox, address _rollupEventBridge, bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _bridge, address _challengeFactory, address _nodeFactory, bytes _extraConfig, address _admin) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _outbox common.Address, _rollupEventBridge common.Address, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _bridge common.Address, _challengeFactory common.Address, _nodeFactory common.Address, _extraConfig []byte, _admin common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _outbox, _rollupEventBridge, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _bridge, _challengeFactory, _nodeFactory, _extraConfig, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x5f81e5fc.
//
// Solidity: function initialize(address _outbox, address _rollupEventBridge, bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _bridge, address _challengeFactory, address _nodeFactory, bytes _extraConfig, address _admin) returns()
func (_Rollup *RollupSession) Initialize(_outbox common.Address, _rollupEventBridge common.Address, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _bridge common.Address, _challengeFactory common.Address, _nodeFactory common.Address, _extraConfig []byte, _admin common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _outbox, _rollupEventBridge, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _bridge, _challengeFactory, _nodeFactory, _extraConfig, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x5f81e5fc.
//
// Solidity: function initialize(address _outbox, address _rollupEventBridge, bytes32 _machineHash, uint256 _challengePeriodBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _bridge, address _challengeFactory, address _nodeFactory, bytes _extraConfig, address _admin) returns()
func (_Rollup *RollupTransactorSession) Initialize(_outbox common.Address, _rollupEventBridge common.Address, _machineHash [32]byte, _challengePeriodBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _bridge common.Address, _challengeFactory common.Address, _nodeFactory common.Address, _extraConfig []byte, _admin common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _outbox, _rollupEventBridge, _machineHash, _challengePeriodBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _bridge, _challengeFactory, _nodeFactory, _extraConfig, _admin)
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
const RollupCoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupCoreBin is the compiled bytecode used for deploying new contracts.
var RollupCoreBin = "0x608060405234801561001057600080fd5b506104a7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100d05760003560e01c80632f30cabd146100d55780633e96576e1461010d5780634f0f4aa9146101335780636177fd181461016c57806362a82d7d146101a657806363721d6b146101c357806365f7f80d146101cb57806369fd251c146101d35780637ba9534a146101f95780638640ce5f14610201578063d01e660214610209578063d735e21d14610226578063dff697871461022e578063e8bd492214610236578063ef40a67014610292578063f33e1fac146102b8575b600080fd5b6100fb600480360360208110156100eb57600080fd5b50356001600160a01b03166102d5565b60408051918252519081900360200190f35b6100fb6004803603602081101561012357600080fd5b50356001600160a01b03166102f0565b6101506004803603602081101561014957600080fd5b503561030e565b604080516001600160a01b039092168252519081900360200190f35b6101926004803603602081101561018257600080fd5b50356001600160a01b0316610329565b604080519115158252519081900360200190f35b610150600480360360208110156101bc57600080fd5b5035610351565b6100fb61037b565b6100fb610381565b610150600480360360208110156101e957600080fd5b50356001600160a01b0316610387565b6100fb6103a8565b6100fb6103ae565b6101506004803603602081101561021f57600080fd5b50356103b4565b6100fb6103e3565b6100fb6103e9565b61025c6004803603602081101561024c57600080fd5b50356001600160a01b03166103ef565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b6100fb600480360360208110156102a857600080fd5b50356001600160a01b031661042b565b6100fb600480360360208110156102ce57600080fd5b5035610449565b6001600160a01b031660009081526008602052604090205490565b6001600160a01b031660009081526006602052604090206001015490565b6000908152600460205260409020546001600160a01b031690565b6001600160a01b0316600090815260066020526040902060030154600160a01b900460ff1690565b60006005828154811061036057fe5b6000918252602090912001546001600160a01b031692915050565b60075490565b60005490565b6001600160a01b039081166000908152600660205260409020600301541690565b60025490565b60035490565b6000600782815481106103c357fe5b60009182526020909120600290910201546001600160a01b031692915050565b60015490565b60055490565b6006602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6001600160a01b031660009081526006602052604090206002015490565b60006007828154811061045857fe5b906000526020600020906002020160010154905091905056fea2646970667358221220f3ed7cb6ee3aaafdf75a49bdf1ddce1d13ec9b2457a03dc12c25a9442ada2d4664736f6c634300060c0033"

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
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetNode(&_RollupCore.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetNode(&_RollupCore.CallOpts, nodeNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetStakerAddress(&_RollupCore.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetStakerAddress(&_RollupCore.CallOpts, stakerNum)
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
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.ZombieAddress(&_RollupCore.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.ZombieAddress(&_RollupCore.CallOpts, zombieNum)
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
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _RollupCore.Contract.ZombieLatestStakedNode(&_RollupCore.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _RollupCore.Contract.ZombieLatestStakedNode(&_RollupCore.CallOpts, zombieNum)
}

// RollupEventBridgeABI is the input ABI used to generate the binding from.
const RollupEventBridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"claimNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"nodeConfirmed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"}],\"name\":\"nodeCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"nodeRejected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraConfig\",\"type\":\"bytes\"}],\"name\":\"rollupInitialized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"stakeCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupEventBridgeBin is the compiled bytecode used for deploying new contracts.
var RollupEventBridgeBin = "0x608060405234801561001057600080fd5b506040516108f23803806108f28339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556108788061007a6000396000f3fe608060405234801561001057600080fd5b50600436106100615760003560e01c8062fc32f71461006657806316b9109b1461010257806330a826b41461011f57806364126c7c1461013c5780638b8ca19914610168578063f03c04a5146101a0575b600080fd5b610100600480360360c081101561007c57600080fd5b8135916020810135916040820135916001600160a01b0360608201358116926080830135909116919081019060c0810160a0820135600160201b8111156100c257600080fd5b8201836020820111156100d457600080fd5b803590602001918460018302840111600160201b831117156100f557600080fd5b5090925090506101cc565b005b6101006004803603602081101561011857600080fd5b5035610320565b6101006004803603602081101561013557600080fd5b50356103ac565b6101006004803603604081101561015257600080fd5b50803590602001356001600160a01b0316610433565b6101006004803603608081101561017e57600080fd5b50803590602081013590604081013590606001356001600160a01b0316610666565b610100600480360360408110156101b657600080fd5b506001600160a01b038135169060200135610710565b6001546001600160a01b03163314610219576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60008787878760601b60601c6001600160a01b03168760601b60601c6001600160a01b03168787604051602001808881526020018781526020018681526020018581526020018481526020018383808284376040805191909301818103601f190182528084528151602092830120600080546302bbfad160e01b8452600480850152336024850152604484018390529551919e506001600160a01b039095169c506302bbfad19b5060648083019b509299509750879003019450859350909150889050803b1580156102ea57600080fd5b505af11580156102fe573d6000803e3d6000fd5b505050506040513d602081101561031457600080fd5b50505050505050505050565b6001546001600160a01b0316331461036d576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f81b602080830191909152602180830185905283518084039091018152604190920190925280519101206103a7816107b5565b505050565b6001546001600160a01b031633146103f9576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f91b602080830191909152602180830185905283518084039091018152604190920190925280519101206103a7816107b5565b6001546001600160a01b03163314610480576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60015460408051634f0f4aa960e01b81526004810185905290516001600160a01b03909216916000918391634f0f4aa991602480820192602092909190829003018186803b1580156104d157600080fd5b505afa1580156104e5573d6000803e3d6000fd5b505050506040513d60208110156104fb57600080fd5b5051604080516348b4573960e11b81526001600160a01b038681166004830152915192935090831691639168ae7291602480820192602092909190829003018186803b15801561054a57600080fd5b505afa15801561055e573d6000803e3d6000fd5b505050506040513d602081101561057457600080fd5b50516105b4576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b816001600160a01b0316632b2af0ab856040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156105f857600080fd5b505afa15801561060c573d6000803e3d6000fd5b505060408051600160fa1b602080830191909152602182018990526001600160a01b0388166041808401919091528351808403909101815260619092019092528051910120915061065e9050816107b5565b505050505050565b6001546001600160a01b031633146106b3576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b6040805160006020808301919091526021820187905260418201869052436061830152608182018590526001600160a01b03841660a1808401919091528351808403909101815260c1909201909252805191012061065e816107b5565b6001546001600160a01b0316331461075d576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600360f81b6020808301919091526001600160a01b0385166021830152604182018490524360618084019190915283518084039091018152608190920190925280519101206107af816107b5565b50505050565b60008054604080516302bbfad160e01b8152600860048201523360248201526044810185905290516001600160a01b03909216916302bbfad19160648082019260209290919082900301818787803b15801561081057600080fd5b505af1158015610824573d6000803e3d6000fd5b505050506040513d602081101561083a57600080fd5b50519291505056fea26469706673582212204ed0084a8b13fb2a411ef822056d5a557028e6f1c592b8ed4b0d8fe6a13d5ee164736f6c634300060c0033"

// DeployRollupEventBridge deploys a new Ethereum contract, binding an instance of RollupEventBridge to it.
func DeployRollupEventBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _rollup common.Address) (common.Address, *types.Transaction, *RollupEventBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupEventBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupEventBridgeBin), backend, _bridge, _rollup)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupEventBridge{RollupEventBridgeCaller: RollupEventBridgeCaller{contract: contract}, RollupEventBridgeTransactor: RollupEventBridgeTransactor{contract: contract}, RollupEventBridgeFilterer: RollupEventBridgeFilterer{contract: contract}}, nil
}

// RollupEventBridge is an auto generated Go binding around an Ethereum contract.
type RollupEventBridge struct {
	RollupEventBridgeCaller     // Read-only binding to the contract
	RollupEventBridgeTransactor // Write-only binding to the contract
	RollupEventBridgeFilterer   // Log filterer for contract events
}

// RollupEventBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupEventBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupEventBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupEventBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupEventBridgeSession struct {
	Contract     *RollupEventBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RollupEventBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupEventBridgeCallerSession struct {
	Contract *RollupEventBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RollupEventBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupEventBridgeTransactorSession struct {
	Contract     *RollupEventBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RollupEventBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupEventBridgeRaw struct {
	Contract *RollupEventBridge // Generic contract binding to access the raw methods on
}

// RollupEventBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupEventBridgeCallerRaw struct {
	Contract *RollupEventBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// RollupEventBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupEventBridgeTransactorRaw struct {
	Contract *RollupEventBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupEventBridge creates a new instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridge(address common.Address, backend bind.ContractBackend) (*RollupEventBridge, error) {
	contract, err := bindRollupEventBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridge{RollupEventBridgeCaller: RollupEventBridgeCaller{contract: contract}, RollupEventBridgeTransactor: RollupEventBridgeTransactor{contract: contract}, RollupEventBridgeFilterer: RollupEventBridgeFilterer{contract: contract}}, nil
}

// NewRollupEventBridgeCaller creates a new read-only instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeCaller(address common.Address, caller bind.ContractCaller) (*RollupEventBridgeCaller, error) {
	contract, err := bindRollupEventBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeCaller{contract: contract}, nil
}

// NewRollupEventBridgeTransactor creates a new write-only instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupEventBridgeTransactor, error) {
	contract, err := bindRollupEventBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeTransactor{contract: contract}, nil
}

// NewRollupEventBridgeFilterer creates a new log filterer instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupEventBridgeFilterer, error) {
	contract, err := bindRollupEventBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeFilterer{contract: contract}, nil
}

// bindRollupEventBridge binds a generic wrapper to an already deployed contract.
func bindRollupEventBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupEventBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupEventBridge *RollupEventBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupEventBridge.Contract.RollupEventBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupEventBridge *RollupEventBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupEventBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupEventBridge *RollupEventBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupEventBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupEventBridge *RollupEventBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupEventBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupEventBridge *RollupEventBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupEventBridge *RollupEventBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.contract.Transact(opts, method, params...)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) ClaimNode(opts *bind.TransactOpts, nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "claimNode", nodeNum, staker)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeSession) ClaimNode(nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.ClaimNode(&_RollupEventBridge.TransactOpts, nodeNum, staker)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) ClaimNode(nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.ClaimNode(&_RollupEventBridge.TransactOpts, nodeNum, staker)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeConfirmed(opts *bind.TransactOpts, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeConfirmed", nodeNum)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeConfirmed(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeConfirmed(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeConfirmed(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeConfirmed(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeCreated(opts *bind.TransactOpts, nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeCreated", nodeNum, prev, deadline, asserter)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeCreated(nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeCreated(&_RollupEventBridge.TransactOpts, nodeNum, prev, deadline, asserter)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeCreated(nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeCreated(&_RollupEventBridge.TransactOpts, nodeNum, prev, deadline, asserter)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeRejected(opts *bind.TransactOpts, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeRejected", nodeNum)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeRejected(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeRejected(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeRejected(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeRejected(&_RollupEventBridge.TransactOpts, nodeNum)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0x00fc32f7.
//
// Solidity: function rollupInitialized(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) RollupInitialized(opts *bind.TransactOpts, challengePeriodBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "rollupInitialized", challengePeriodBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0x00fc32f7.
//
// Solidity: function rollupInitialized(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeSession) RollupInitialized(challengePeriodBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupInitialized(&_RollupEventBridge.TransactOpts, challengePeriodBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0x00fc32f7.
//
// Solidity: function rollupInitialized(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) RollupInitialized(challengePeriodBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupInitialized(&_RollupEventBridge.TransactOpts, challengePeriodBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) StakeCreated(opts *bind.TransactOpts, staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "stakeCreated", staker, nodeNum)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) StakeCreated(staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.StakeCreated(&_RollupEventBridge.TransactOpts, staker, nodeNum)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) StakeCreated(staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.StakeCreated(&_RollupEventBridge.TransactOpts, staker, nodeNum)
}

// RollupLibABI is the input ABI used to generate the binding from.
const RollupLibABI = "[]"

// RollupLibBin is the compiled bytecode used for deploying new contracts.
var RollupLibBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202f92b88b92655590547eaab17ecbc2bf7e85b1d03c6d8f7711d587ed528e8a5864736f6c634300060c0033"

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
const ValidatorUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkDecidableNextNode\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.ConfirmType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"node1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findNodeConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findStakerConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"refundableStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"requireConfirmable\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"requireRejectableNextNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakedNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"stakerInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"successorNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorUtilsBin is the compiled bytecode used for deploying new contracts.
var ValidatorUtilsBin = "0x608060405234801561001057600080fd5b50612708806100206000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c80631fc43bb61461009e5780632b1062cf146100b35780633082d029146100de5780637464ae06146101005780637988ad37146101205780638730825e146101335780638f67e6bb14610153578063974dbf2a14610176578063abeba4f714610197578063c308eaaf146101aa578063e48a5f7b146101bd575b600080fd5b6100b16100ac36600461229c565b6101e0565b005b6100c66100c13660046123d9565b610677565b6040516100d593929190612545565b60405180910390f35b6100f16100ec36600461239f565b61077b565b6040516100d593929190612570565b61011361010e36600461229c565b610c4c565b6040516100d5919061246c565b6100f161012e3660046122f0565b610ed2565b610146610141366004612340565b610ff4565b6040516100d591906124b9565b6101666101613660046122b8565b6111c9565b6040516100d594939291906124f1565b6101896101843660046123d9565b6113cf565b6040516100d592919061267f565b6101136101a536600461236b565b611557565b6101466101b83660046122b8565b6116e1565b6101d06101cb36600461229c565b611926565b6040516100d59493929190612696565b806001600160a01b03166367425daf6040518163ffffffff1660e01b815260040160006040518083038186803b15801561021957600080fd5b505afa15801561022d573d6000803e3d6000fd5b50505050806001600160a01b031663b15a38df6040518163ffffffff1660e01b815260040160006040518083038186803b15801561026a57600080fd5b505afa15801561027e573d6000803e3d6000fd5b505050506000816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b1580156102bd57600080fd5b505afa1580156102d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f5919061241c565b9050600081116103205760405162461bcd60e51b8152600401610317906125b7565b60405180910390fd5b6000826001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561035b57600080fd5b505afa15801561036f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610393919061241c565b90506000836001600160a01b0316634f0f4aa9836040518263ffffffff1660e01b81526004016103c39190612676565b60206040518083038186803b1580156103db57600080fd5b505afa1580156103ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104139190612260565b9050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b15801561044e57600080fd5b505afa158015610462573d6000803e3d6000fd5b50505050836001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561049f57600080fd5b505afa1580156104b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d7919061241c565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561051057600080fd5b505afa158015610524573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610548919061241c565b146105655760405162461bcd60e51b815260040161031790612603565b604051630128a01960e21b81526001600160a01b038516906304a2806490610591908490600401612458565b60206040518083038186803b1580156105a957600080fd5b505afa1580156105bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e1919061241c565b8301816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561061c57600080fd5b505afa158015610630573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610654919061241c565b146106715760405162461bcd60e51b8152600401610317906125db565b50505050565b604051630fe21ddb60e11b8152600090819081903090631fc43bb6906106a1908b90600401612458565b60006040518083038186803b1580156106b957600080fd5b505afa9250505080156106ca575060015b6106d3576106e2565b50600191506000905080610770565b604051634ba6df9560e11b8152309063974dbf2a9061070d908b908b908b908b908b90600401612517565b604080518083038186803b15801561072457600080fd5b505afa925050508015610754575060408051601f3d908101601f1916820190925261075191810190612434565b60015b61076657506000915081905080610770565b6002945090925090505b955095509592505050565b600080600080876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107ba57600080fd5b505afa1580156107ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f2919061241c565b90506000886001600160a01b0316634f0f4aa9896040518263ffffffff1660e01b81526004016108229190612676565b60206040518083038186803b15801561083a57600080fd5b505afa15801561084e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108729190612260565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156108aa57600080fd5b505afa1580156108be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e2919061241c565b90506000896001600160a01b0316634f0f4aa9896040518263ffffffff1660e01b81526004016109129190612676565b60206040518083038186803b15801561092a57600080fd5b505afa15801561093e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109629190612260565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561099a57600080fd5b505afa1580156109ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d2919061241c565b905060005b87811015610c3357888a14156109fa5760008a8a96509650965050505050610c42565b81831415610a155760018a8a96509650965050505050610c42565b83831080610a2257508382105b15610a3b57600260008096509650965050505050610c42565b888a1015610b39578198508a6001600160a01b0316634f0f4aa98a6040518263ffffffff1660e01b8152600401610a729190612676565b60206040518083038186803b158015610a8a57600080fd5b505afa158015610a9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac29190612260565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610afa57600080fd5b505afa158015610b0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b32919061241c565b9150610c2b565b8299508a6001600160a01b0316634f0f4aa98b6040518263ffffffff1660e01b8152600401610b689190612676565b60206040518083038186803b158015610b8057600080fd5b505afa158015610b94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bb89190612260565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610bf057600080fd5b505afa158015610c04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c28919061241c565b92505b6001016109d7565b50600389899550955095505050505b9450945094915050565b60606000826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015610c8957600080fd5b505afa158015610c9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc1919061241c565b90506060816001600160401b0381118015610cdb57600080fd5b50604051908082528060200260200182016040528015610d05578160200160208202803683370190505b5090506000846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d4357600080fd5b505afa158015610d57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7b919061241c565b90506000805b84811015610ec7576040516362a82d7d60e01b81526000906001600160a01b038916906362a82d7d90610db8908590600401612676565b60206040518083038186803b158015610dd057600080fd5b505afa158015610de4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e089190612260565b90506000886001600160a01b0316633e96576e836040518263ffffffff1660e01b8152600401610e389190612458565b60206040518083038186803b158015610e5057600080fd5b505afa158015610e64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e88919061241c565b9050848111610ebd5781868581518110610e9e57fe5b6001600160a01b03909216602092830291909101909101526001909301925b5050600101610d81565b508252509392505050565b600080600080876001600160a01b0316633e96576e886040518263ffffffff1660e01b8152600401610f049190612458565b60206040518083038186803b158015610f1c57600080fd5b505afa158015610f30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f54919061241c565b90506000886001600160a01b0316633e96576e886040518263ffffffff1660e01b8152600401610f849190612458565b60206040518083038186803b158015610f9c57600080fd5b505afa158015610fb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd4919061241c565b9050610fe28983838961077b565b94509450945050509450945094915050565b60408051620186a08082526230d4208201909252606091829190602082016230d400803683370190505090506000600184015b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561106057600080fd5b505afa158015611074573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611098919061241c565b81116111bf57604051634f0f4aa960e01b81526000906001600160a01b03881690634f0f4aa9906110cd908590600401612676565b60206040518083038186803b1580156110e557600080fd5b505afa1580156110f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111d9190612260565b905085816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561115957600080fd5b505afa15801561116d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611191919061241c565b14156111b657818484815181106111a457fe5b60209081029190910101526001909201915b50600101611027565b5081529392505050565b600080600080856001600160a01b0316636177fd18866040518263ffffffff1660e01b81526004016111fb9190612458565b60206040518083038186803b15801561121357600080fd5b505afa158015611227573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124b919061227c565b604051631f4b2bb760e11b81526001600160a01b03881690633e96576e90611277908990600401612458565b60206040518083038186803b15801561128f57600080fd5b505afa1580156112a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c7919061241c565b604051630ef40a6760e41b81526001600160a01b0389169063ef40a670906112f3908a90600401612458565b60206040518083038186803b15801561130b57600080fd5b505afa15801561131f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611343919061241c565b604051631a7f494760e21b81526001600160a01b038a16906369fd251c9061136f908b90600401612458565b60206040518083038186803b15801561138757600080fd5b505afa15801561139b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113bf9190612260565b9299919850965090945092505050565b600080866001600160a01b03166367425daf6040518163ffffffff1660e01b815260040160006040518083038186803b15801561140b57600080fd5b505afa15801561141f573d6000803e3d6000fd5b50505050866001600160a01b031663b15a38df6040518163ffffffff1660e01b815260040160006040518083038186803b15801561145c57600080fd5b505afa158015611470573d6000803e3d6000fd5b505050506000876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156114af57600080fd5b505afa1580156114c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e7919061241c565b905060006114f489611aff565b9050801561150a5760008093509350505061154d565b60008060006115218c8c87600101018c8c8c611f21565b925092509250826115445760405162461bcd60e51b815260040161031790612629565b90955093505050505b9550959350505050565b60606000846001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561159457600080fd5b505afa1580156115a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115cc919061241c565b90508083850110156115dd57508282015b6060816001600160401b03811180156115f557600080fd5b5060405190808252806020026020018201604052801561161f578160200160208202803683370190505b50905060005b828110156116d7576040516362a82d7d60e01b81526001600160a01b038816906362a82d7d9061165b9089850190600401612676565b60206040518083038186803b15801561167357600080fd5b505afa158015611687573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116ab9190612260565b8282815181106116b757fe5b6001600160a01b0390921660209283029190910190910152600101611625565b5095945050505050565b60408051620186a08082526230d4208201909252606091829190602082016230d40080368337019050509050600080856001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561174957600080fd5b505afa15801561175d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611781919061241c565b90505b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117bd57600080fd5b505afa1580156117d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f5919061241c565b81116111bf57604051634f0f4aa960e01b81526000906001600160a01b03881690634f0f4aa99061182a908590600401612676565b60206040518083038186803b15801561184257600080fd5b505afa158015611856573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187a9190612260565b6040516348b4573960e11b81529091506001600160a01b03821690639168ae72906118a9908990600401612458565b60206040518083038186803b1580156118c157600080fd5b505afa1580156118d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f9919061227c565b1561191d578184848151811061190b57fe5b60209081029190910101526001909201915b50600101611784565b600080600080846001600160a01b03166346c2781a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561196557600080fd5b505afa158015611979573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061199d919061241c565b9350846001600160a01b0316635e8ef1066040518163ffffffff1660e01b815260040160206040518083038186803b1580156119d857600080fd5b505afa1580156119ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a10919061241c565b9250846001600160a01b03166376e7e23b6040518163ffffffff1660e01b815260040160206040518083038186803b158015611a4b57600080fd5b505afa158015611a5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a83919061241c565b9150846001600160a01b03166351ed6a306040518163ffffffff1660e01b815260040160206040518083038186803b158015611abe57600080fd5b505afa158015611ad2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611af69190612260565b90509193509193565b6000816001600160a01b03166367425daf6040518163ffffffff1660e01b815260040160006040518083038186803b158015611b3a57600080fd5b505afa158015611b4e573d6000803e3d6000fd5b505050506000826001600160a01b0316634f0f4aa9846001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611b9c57600080fd5b505afa158015611bb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd4919061241c565b6040518263ffffffff1660e01b8152600401611bf09190612676565b60206040518083038186803b158015611c0857600080fd5b505afa158015611c1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c409190612260565b90506000836001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c7d57600080fd5b505afa158015611c91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cb5919061241c565b826001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611cee57600080fd5b505afa158015611d02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d26919061241c565b1490508015611f1a57836001600160a01b031663b15a38df6040518163ffffffff1660e01b815260040160006040518083038186803b158015611d6857600080fd5b505afa158015611d7c573d6000803e3d6000fd5b50505050816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611db957600080fd5b505afa158015611dcd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611df1919061241c565b431015611e105760405162461bcd60e51b81526004016103179061264d565b604051630128a01960e21b81526001600160a01b038516906304a2806490611e3c908590600401612458565b60206040518083038186803b158015611e5457600080fd5b505afa158015611e68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e8c919061241c565b826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015611ec557600080fd5b505afa158015611ed9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611efd919061241c565b14611f1a5760405162461bcd60e51b815260040161031790612592565b9392505050565b600080600080886001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611f6057600080fd5b505afa158015611f74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f98919061241c565b905080881115611fb357600080600093509350935050610770565b87810387811115611fc15750865b6120488a8a8c6001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611fff57600080fd5b505afa158015612013573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612037919061241c565b846120438f8d8d611557565b61205b565b9450945094505050955095509592505050565b6000806000808451905060005b86811161224c57604051634f0f4aa960e01b8152898201906000906001600160a01b038d1690634f0f4aa9906120a2908590600401612676565b60206040518083038186803b1580156120ba57600080fd5b505afa1580156120ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120f29190612260565b905089816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561212e57600080fd5b505afa158015612142573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612166919061241c565b14612172575050612244565b60005b8481101561224057816001600160a01b0316639168ae728a838151811061219857fe5b60200260200101516040518263ffffffff1660e01b81526004016121bc9190612458565b60206040518083038186803b1580156121d457600080fd5b505afa1580156121e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061220c919061227c565b15612238576001838a838151811061222057fe5b60200260200101519750975097505050505050610770565b600101612175565b5050505b600101612068565b506000998a99508998509650505050505050565b600060208284031215612271578081fd5b8151611f1a816126ba565b60006020828403121561228d578081fd5b81518015158114611f1a578182fd5b6000602082840312156122ad578081fd5b8135611f1a816126ba565b600080604083850312156122ca578081fd5b82356122d5816126ba565b915060208301356122e5816126ba565b809150509250929050565b60008060008060808587031215612305578182fd5b8435612310816126ba565b93506020850135612320816126ba565b92506040850135612330816126ba565b9396929550929360600135925050565b60008060408385031215612352578182fd5b823561235d816126ba565b946020939093013593505050565b60008060006060848603121561237f578283fd5b833561238a816126ba565b95602085013595506040909401359392505050565b600080600080608085870312156123b4578384fd5b84356123bf816126ba565b966020860135965060408601359560600135945092505050565b600080600080600060a086880312156123f0578081fd5b85356123fb816126ba565b97602087013597506040870135966060810135965060800135945092505050565b60006020828403121561242d578081fd5b5051919050565b60008060408385031215612446578182fd5b8251915060208301516122e5816126ba565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b818110156124ad5783516001600160a01b031683529284019291840191600101612488565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156124ad578351835292840192918401916001016124d5565b9315158452602084019290925260408301526001600160a01b0316606082015260800190565b6001600160a01b03959095168552602085019390935260408401919091526060830152608082015260a00190565b606081016003851061255357fe5b93815260208101929092526001600160a01b031660409091015290565b606081016004851061257e57fe5b938152602081019290925260409091015290565b6020808252600b908201526a4841535f5354414b45525360a81b604082015260600190565b6020808252600a90820152694e4f5f5354414b45525360b01b604082015260600190565b6020808252600e908201526d1393d517d0531317d4d51052d15160921b604082015260600190565b6020808252600c908201526b24a72b20a624a22fa82922ab60a11b604082015260600190565b6020808252600a90820152694e4f5f4558414d504c4560b01b604082015260600190565b6020808252600f908201526e4245464f52455f444541444c494e4560881b604082015260600190565b90815260200190565b9182526001600160a01b0316602082015260400190565b938452602084019290925260408301526001600160a01b0316606082015260800190565b6001600160a01b03811681146126cf57600080fd5b5056fea2646970667358221220566443423c30e8ed19823ed2aab21d24d3d48d9bed5c0bd42bc947ef56683b8a64736f6c634300060c0033"

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

// RequireConfirmable is a free data retrieval call binding the contract method 0x1fc43bb6.
//
// Solidity: function requireConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCaller) RequireConfirmable(opts *bind.CallOpts, rollup common.Address) error {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "requireConfirmable", rollup)

	if err != nil {
		return err
	}

	return err

}

// RequireConfirmable is a free data retrieval call binding the contract method 0x1fc43bb6.
//
// Solidity: function requireConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsSession) RequireConfirmable(rollup common.Address) error {
	return _ValidatorUtils.Contract.RequireConfirmable(&_ValidatorUtils.CallOpts, rollup)
}

// RequireConfirmable is a free data retrieval call binding the contract method 0x1fc43bb6.
//
// Solidity: function requireConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCallerSession) RequireConfirmable(rollup common.Address) error {
	return _ValidatorUtils.Contract.RequireConfirmable(&_ValidatorUtils.CallOpts, rollup)
}

// RequireRejectableNextNode is a free data retrieval call binding the contract method 0x974dbf2a.
//
// Solidity: function requireRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsCaller) RequireRejectableNextNode(opts *bind.CallOpts, rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "requireRejectableNextNode", rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// RequireRejectableNextNode is a free data retrieval call binding the contract method 0x974dbf2a.
//
// Solidity: function requireRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsSession) RequireRejectableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.RequireRejectableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// RequireRejectableNextNode is a free data retrieval call binding the contract method 0x974dbf2a.
//
// Solidity: function requireRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsCallerSession) RequireRejectableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.RequireRejectableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
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
