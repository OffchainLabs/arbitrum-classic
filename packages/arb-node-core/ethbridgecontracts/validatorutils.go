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
const INodeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"addStaker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstChildBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestChildNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"newChildConfirmDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noChildConfirmedBeforeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"removeStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requirePastChildConfirmDeadline\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requirePastDeadline\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"latestConfirmed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"requireRejectExample\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// FirstChildBlock is a free data retrieval call binding the contract method 0xd7ff5e35.
//
// Solidity: function firstChildBlock() view returns(uint256)
func (_INode *INodeCaller) FirstChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "firstChildBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstChildBlock is a free data retrieval call binding the contract method 0xd7ff5e35.
//
// Solidity: function firstChildBlock() view returns(uint256)
func (_INode *INodeSession) FirstChildBlock() (*big.Int, error) {
	return _INode.Contract.FirstChildBlock(&_INode.CallOpts)
}

// FirstChildBlock is a free data retrieval call binding the contract method 0xd7ff5e35.
//
// Solidity: function firstChildBlock() view returns(uint256)
func (_INode *INodeCallerSession) FirstChildBlock() (*big.Int, error) {
	return _INode.Contract.FirstChildBlock(&_INode.CallOpts)
}

// LatestChildNumber is a free data retrieval call binding the contract method 0xf0dd77ff.
//
// Solidity: function latestChildNumber() view returns(uint256)
func (_INode *INodeCaller) LatestChildNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "latestChildNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestChildNumber is a free data retrieval call binding the contract method 0xf0dd77ff.
//
// Solidity: function latestChildNumber() view returns(uint256)
func (_INode *INodeSession) LatestChildNumber() (*big.Int, error) {
	return _INode.Contract.LatestChildNumber(&_INode.CallOpts)
}

// LatestChildNumber is a free data retrieval call binding the contract method 0xf0dd77ff.
//
// Solidity: function latestChildNumber() view returns(uint256)
func (_INode *INodeCallerSession) LatestChildNumber() (*big.Int, error) {
	return _INode.Contract.LatestChildNumber(&_INode.CallOpts)
}

// NoChildConfirmedBeforeBlock is a free data retrieval call binding the contract method 0xa0369c14.
//
// Solidity: function noChildConfirmedBeforeBlock() view returns(uint256)
func (_INode *INodeCaller) NoChildConfirmedBeforeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "noChildConfirmedBeforeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoChildConfirmedBeforeBlock is a free data retrieval call binding the contract method 0xa0369c14.
//
// Solidity: function noChildConfirmedBeforeBlock() view returns(uint256)
func (_INode *INodeSession) NoChildConfirmedBeforeBlock() (*big.Int, error) {
	return _INode.Contract.NoChildConfirmedBeforeBlock(&_INode.CallOpts)
}

// NoChildConfirmedBeforeBlock is a free data retrieval call binding the contract method 0xa0369c14.
//
// Solidity: function noChildConfirmedBeforeBlock() view returns(uint256)
func (_INode *INodeCallerSession) NoChildConfirmedBeforeBlock() (*big.Int, error) {
	return _INode.Contract.NoChildConfirmedBeforeBlock(&_INode.CallOpts)
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

// RequirePastChildConfirmDeadline is a free data retrieval call binding the contract method 0x3aa19274.
//
// Solidity: function requirePastChildConfirmDeadline() view returns()
func (_INode *INodeCaller) RequirePastChildConfirmDeadline(opts *bind.CallOpts) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "requirePastChildConfirmDeadline")

	if err != nil {
		return err
	}

	return err

}

// RequirePastChildConfirmDeadline is a free data retrieval call binding the contract method 0x3aa19274.
//
// Solidity: function requirePastChildConfirmDeadline() view returns()
func (_INode *INodeSession) RequirePastChildConfirmDeadline() error {
	return _INode.Contract.RequirePastChildConfirmDeadline(&_INode.CallOpts)
}

// RequirePastChildConfirmDeadline is a free data retrieval call binding the contract method 0x3aa19274.
//
// Solidity: function requirePastChildConfirmDeadline() view returns()
func (_INode *INodeCallerSession) RequirePastChildConfirmDeadline() error {
	return _INode.Contract.RequirePastChildConfirmDeadline(&_INode.CallOpts)
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
// Solidity: function addStaker(address staker) returns(uint256)
func (_INode *INodeTransactor) AddStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "addStaker", staker)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns(uint256)
func (_INode *INodeSession) AddStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.AddStaker(&_INode.TransactOpts, staker)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns(uint256)
func (_INode *INodeTransactorSession) AddStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.AddStaker(&_INode.TransactOpts, staker)
}

// ChildCreated is a paid mutator transaction binding the contract method 0x1bc09d0a.
//
// Solidity: function childCreated(uint256 ) returns()
func (_INode *INodeTransactor) ChildCreated(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "childCreated", arg0)
}

// ChildCreated is a paid mutator transaction binding the contract method 0x1bc09d0a.
//
// Solidity: function childCreated(uint256 ) returns()
func (_INode *INodeSession) ChildCreated(arg0 *big.Int) (*types.Transaction, error) {
	return _INode.Contract.ChildCreated(&_INode.TransactOpts, arg0)
}

// ChildCreated is a paid mutator transaction binding the contract method 0x1bc09d0a.
//
// Solidity: function childCreated(uint256 ) returns()
func (_INode *INodeTransactorSession) ChildCreated(arg0 *big.Int) (*types.Transaction, error) {
	return _INode.Contract.ChildCreated(&_INode.TransactOpts, arg0)
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

// NewChildConfirmDeadline is a paid mutator transaction binding the contract method 0x6971dfe5.
//
// Solidity: function newChildConfirmDeadline(uint256 deadline) returns()
func (_INode *INodeTransactor) NewChildConfirmDeadline(opts *bind.TransactOpts, deadline *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "newChildConfirmDeadline", deadline)
}

// NewChildConfirmDeadline is a paid mutator transaction binding the contract method 0x6971dfe5.
//
// Solidity: function newChildConfirmDeadline(uint256 deadline) returns()
func (_INode *INodeSession) NewChildConfirmDeadline(deadline *big.Int) (*types.Transaction, error) {
	return _INode.Contract.NewChildConfirmDeadline(&_INode.TransactOpts, deadline)
}

// NewChildConfirmDeadline is a paid mutator transaction binding the contract method 0x6971dfe5.
//
// Solidity: function newChildConfirmDeadline(uint256 deadline) returns()
func (_INode *INodeTransactorSession) NewChildConfirmDeadline(deadline *big.Int) (*types.Transaction, error) {
	return _INode.Contract.NewChildConfirmDeadline(&_INode.TransactOpts, deadline)
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
var ProxyAdminBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6108658061007d6000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100bf5780637eff275e146100d65780638da5cb5b146101115780639623609d1461012657806399a88ec4146101e3578063f2fde38b1461021e578063f3b7dead14610251575b600080fd5b34801561007c57600080fd5b506100a36004803603602081101561009357600080fd5b50356001600160a01b0316610284565b604080516001600160a01b039092168252519081900360200190f35b3480156100cb57600080fd5b506100d4610316565b005b3480156100e257600080fd5b506100d4600480360360408110156100f957600080fd5b506001600160a01b03813581169160200135166103b0565b34801561011d57600080fd5b506100a361047d565b6100d46004803603606081101561013c57600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460018302840111600160201b831117156101a257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061048c945050505050565b3480156101ef57600080fd5b506100d46004803603604081101561020657600080fd5b506001600160a01b03813581169160200135166105c5565b34801561022a57600080fd5b506100d46004803603602081101561024157600080fd5b50356001600160a01b0316610676565b34801561025d57600080fd5b506100a36004803603602081101561027457600080fd5b50356001600160a01b0316610766565b6000806060836001600160a01b03166040518080635c60da1b60e01b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b606091505b5091509150816102f757600080fd5b80806020019051602081101561030c57600080fd5b5051949350505050565b61031e6107c5565b6001600160a01b031661032f61047d565b6001600160a01b031614610378576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610810833981519152908390a3600080546001600160a01b0319169055565b6103b86107c5565b6001600160a01b03166103c961047d565b6001600160a01b031614610412576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316638f283970826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b505af1158015610475573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b6104946107c5565b6001600160a01b03166104a561047d565b6001600160a01b0316146104ee576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b826001600160a01b0316634f1ef2863484846040518463ffffffff1660e01b815260040180836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561055b578181015183820152602001610543565b50505050905090810190601f1680156105885780820380516001836020036101000a031916815260200191505b5093505050506000604051808303818588803b1580156105a757600080fd5b505af11580156105bb573d6000803e3d6000fd5b5050505050505050565b6105cd6107c5565b6001600160a01b03166105de61047d565b6001600160a01b031614610627576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316633659cfe6826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b61067e6107c5565b6001600160a01b031661068f61047d565b6001600160a01b0316146106d8576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b6001600160a01b03811661071d5760405162461bcd60e51b81526004018080602001828103825260268152602001806107ca6026913960400191505060405180910390fd5b600080546040516001600160a01b038085169392169160008051602061081083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806060836001600160a01b031660405180806303e1469160e61b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220aaa1273181c4f76a5bfaf7202f501d17216f4bc6fa228e2bed790fcf561d2dd664736f6c634300060c0033"

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
const RollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentNodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[4]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[4]\"},{\"indexed\":false,\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"SentLogs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"addToDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLatestNodeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"beginTruncatingNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"logAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"continueTruncatingNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"executionHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proposedTimes\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"maxMessageCounts\",\"type\":\"uint256[2]\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraChallengeTimeBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[6]\",\"name\":\"connectedContracts\",\"type\":\"address[6]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"newStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"successorWithStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"requireUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireUnresolvedExists\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expectedNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[4]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[4]\"},{\"internalType\":\"uint256[10]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[10]\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"}],\"name\":\"upgradeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeImplementationAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupBin is the compiled bytecode used for deploying new contracts.
var RollupBin = "0x608060405234801561001057600080fd5b506000805460ff19908116600117909155600b80549091169055615872806200003a6000396000f3fe60806040526004361061029c5760003560e01c8063046f7da2146102a157806304a28064146102b85780630e1ef04c146102fd5780630e21b5ba146103365780631e83d30f146103605780632b2af0ab1461038a5780632e7acfa6146103b45780632f30cabd146103c9578063396b8cbc146103fc5780633b333ea8146104ce5780633e96576e146104fe578063414f23fe1461053157806345e38b6414610561578063488ed1a9146105765780634d26732d146105b15780634f0f4aa9146105c657806351ed6a301461060c578063567ca41b146106215780635c975abb146106545780635dbaf68b1461067d5780635e8ef106146106925780636177fd18146106a757806362a82d7d146106da57806363721d6b1461070457806365f7f80d1461071957806367425daf1461072e57806369fd251c146107435780636f791d29146107765780637427be511461078b57806376e7e23b146107be578063771b2f97146107d35780637ba9534a146107e85780637e2d2155146107fd57806381fbc98a1461082d57806383f94db7146108605780638456cb59146108935780638640ce5f146108a85780638da5cb5b146108bd5780639e8a713f146108d2578063bf5ddcb1146108e7578063ce11e6ab14610972578063d01e660214610987578063d50b580b146109b1578063d735e21d146109e6578063d93fe9c4146109fb578063dff6978714610a10578063e45b7ce614610a25578063e78cea9214610a60578063e8bd492214610a75578063edfd03ed14610ade578063ef40a67014610b08578063f33e1fac14610b3b578063f3f0a03e14610b65578063f851a44014610b91578063f8d1f19414610ba6578063fa7803e614610bd0578063fb64884e14610c0b578063fdaf579714610c28578063ff204f3b14610cdc575b600080fd5b3480156102ad57600080fd5b506102b6610d0f565b005b3480156102c457600080fd5b506102eb600480360360208110156102db57600080fd5b50356001600160a01b0316610db0565b60408051918252519081900360200190f35b34801561030957600080fd5b506102b66004803603604081101561032057600080fd5b50803590602001356001600160a01b0316610e68565b34801561034257600080fd5b506102b66004803603602081101561035957600080fd5b50356111e4565b34801561036c57600080fd5b506102b66004803603602081101561038357600080fd5b5035611492565b34801561039657600080fd5b506102b6600480360360208110156103ad57600080fd5b503561150a565b3480156103c057600080fd5b506102eb6115a3565b3480156103d557600080fd5b506102eb600480360360208110156103ec57600080fd5b50356001600160a01b03166115a9565b34801561040857600080fd5b506102b66004803603606081101561041f57600080fd5b81359190810190604081016020820135600160201b81111561044057600080fd5b82018360208201111561045257600080fd5b803590602001918460018302840111600160201b8311171561047357600080fd5b919390929091602081019035600160201b81111561049057600080fd5b8201836020820111156104a257600080fd5b803590602001918460208302840111600160201b831117156104c357600080fd5b5090925090506115c4565b3480156104da57600080fd5b506102b6600480360360408110156104f157600080fd5b5080359060200135611b32565b34801561050a57600080fd5b506102eb6004803603602081101561052157600080fd5b50356001600160a01b0316611cc8565b34801561053d57600080fd5b506102b66004803603604081101561055457600080fd5b5080359060200135611ce6565b34801561056d57600080fd5b506102eb611eb7565b34801561058257600080fd5b506102b6600480360361014081101561059a57600080fd5b50604081016080820160c083016101008401611ebc565b3480156105bd57600080fd5b506102eb612722565b3480156105d257600080fd5b506105f0600480360360208110156105e957600080fd5b50356129dc565b604080516001600160a01b039092168252519081900360200190f35b34801561061857600080fd5b506105f06129f7565b34801561062d57600080fd5b506102b66004803603602081101561064457600080fd5b50356001600160a01b0316612a06565b34801561066057600080fd5b50610669612b09565b604080519115158252519081900360200190f35b34801561068957600080fd5b506105f0612b12565b34801561069e57600080fd5b506102eb612b21565b3480156106b357600080fd5b50610669600480360360208110156106ca57600080fd5b50356001600160a01b0316612b27565b3480156106e657600080fd5b506105f0600480360360208110156106fd57600080fd5b5035612b4f565b34801561071057600080fd5b506102eb612b79565b34801561072557600080fd5b506102eb612b7f565b34801561073a57600080fd5b506102b6612b85565b34801561074f57600080fd5b506105f06004803603602081101561076657600080fd5b50356001600160a01b0316612bef565b34801561078257600080fd5b50610669612c10565b34801561079757600080fd5b506102b6600480360360208110156107ae57600080fd5b50356001600160a01b0316612c19565b3480156107ca57600080fd5b506102eb612cc4565b3480156107df57600080fd5b506102eb612cca565b3480156107f457600080fd5b506102eb612cd0565b34801561080957600080fd5b506102b66004803603604081101561082057600080fd5b5080359060200135612cd6565b34801561083957600080fd5b506102eb6004803603602081101561085057600080fd5b50356001600160a01b0316612eb6565b34801561086c57600080fd5b506102b66004803603602081101561088357600080fd5b50356001600160a01b0316612f1d565b34801561089f57600080fd5b506102b6612fd6565b3480156108b457600080fd5b506102eb61302a565b3480156108c957600080fd5b506105f0613030565b3480156108de57600080fd5b506105f061303f565b3480156108f357600080fd5b506102b66004803603604081101561090a57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561093457600080fd5b82018360208201111561094657600080fd5b803590602001918460018302840111600160201b8311171561096757600080fd5b50909250905061304e565b34801561097e57600080fd5b506105f0613144565b34801561099357600080fd5b506105f0600480360360208110156109aa57600080fd5b5035613153565b3480156109bd57600080fd5b506102b660048036036101e08110156109d557600080fd5b50803590602081019060a001613182565b3480156109f257600080fd5b506102eb613ac4565b348015610a0757600080fd5b506105f0613aca565b348015610a1c57600080fd5b506102eb613ad9565b348015610a3157600080fd5b506102b660048036036040811015610a4857600080fd5b506001600160a01b0381351690602001351515613adf565b348015610a6c57600080fd5b506105f0613b81565b348015610a8157600080fd5b50610aa860048036036020811015610a9857600080fd5b50356001600160a01b0316613b90565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b348015610aea57600080fd5b506102b660048036036020811015610b0157600080fd5b5035613bcc565b348015610b1457600080fd5b506102eb60048036036020811015610b2b57600080fd5b50356001600160a01b0316613c2b565b348015610b4757600080fd5b506102eb60048036036020811015610b5e57600080fd5b5035613c49565b6102b660048036036040811015610b7b57600080fd5b506001600160a01b038135169060200135613c71565b348015610b9d57600080fd5b506105f0613cd4565b348015610bb257600080fd5b506102eb60048036036020811015610bc957600080fd5b5035613ce3565b348015610bdc57600080fd5b506102b660048036036040811015610bf357600080fd5b506001600160a01b0381358116916020013516613cf5565b6102b660048036036020811015610c2157600080fd5b5035613dce565b348015610c3457600080fd5b506102b660048036036101c0811015610c4c57600080fd5b8135916020810135916040820135916060810135916080820135916001600160a01b0360a082013581169260c08301359091169190810190610100810160e0820135600160201b811115610c9f57600080fd5b820183602082011115610cb157600080fd5b803590602001918460018302840111600160201b83111715610cd257600080fd5b9193509150613f28565b348015610ce857600080fd5b506102b660048036036020811015610cff57600080fd5b50356001600160a01b03166142a8565b6016546001600160a01b03163314610d5b576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601a5460ff1615610da6576040805162461bcd60e51b815260206004820152601060248201526f5354494c4c5f5452554e434154494e4760801b604482015290519081900360640190fd5b610dae61435e565b565b600080610dbb612b79565b90506000805b82811015610e5e57846001600160a01b0316639168ae72610de183613153565b6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610e1e57600080fd5b505afa158015610e32573d6000803e3d6000fd5b505050506040513d6020811015610e4857600080fd5b505115610e56576001909101905b600101610dc1565b509150505b919050565b610e70612b09565b15610eb0576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b610eb8612b85565b6000610ec2612b7f565b90506000610ece613ac4565b90506000610edb826129dc565b905082816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610f1757600080fd5b505afa158015610f2b573d6000803e3d6000fd5b505050506040513d6020811015610f4157600080fd5b5051141561117057610f528561150a565b610f5b84612b27565b610f99576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b610fa2856129dc565b6001600160a01b031663feb508ab84866040518363ffffffff1660e01b815260040180838152602001826001600160a01b031681526020019250505060006040518083038186803b158015610ff657600080fd5b505afa15801561100a573d6000803e3d6000fd5b50505050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b15801561104757600080fd5b505afa15801561105b573d6000803e3d6000fd5b50505050611068836129dc565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b1580156110a057600080fd5b505afa1580156110b4573d6000803e3d6000fd5b505050506110c26000613bcc565b6110cb81610db0565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561110457600080fd5b505afa158015611118573d6000803e3d6000fd5b505050506040513d602081101561112e57600080fd5b505114611170576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b6111786143fe565b60135460408051630c2a09ad60e21b81526004810185905290516001600160a01b03909216916330a826b49160248082019260009290919082900301818387803b1580156111c557600080fd5b505af11580156111d9573d6000803e3d6000fd5b505050505050505050565b6016546001600160a01b03163314611230576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b611238612b09565b611280576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b601a5460ff166112c8576040805162461bcd60e51b815260206004820152600e60248201526d4e4f545f5452554e434154494e4760901b604482015290519081900360640190fd5b60185460195460006112d8613ad9565b90505b6000841180156112ea57508082105b156113cb5760006112fa83612b4f565b9050600061130782611cc8565b90505b60008611801561131957508481115b156113a0576000611329826129dc565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561136457600080fd5b505afa158015611378573d6000803e3d6000fd5b505050506040513d602081101561138e57600080fd5b505160001990970196915061130a9050565b6113aa8282614414565b848111156113be575050506019555061148f565b50506001909101906112db565b601982905560006113da612cd0565b90505b6000851180156113ec57508381115b156114645760006113fc826129dc565b9050806001600160a01b03166383197ef06040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561143957600080fd5b505af115801561144d573d6000803e3d6000fd5b5050600019978801979390930192506113dd915050565b61146d81614433565b8381141561148a5760006018819055601955601a805460ff191690555b505050505b50565b61149a612b09565b156114da576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b6114e333614438565b60006114ed612722565b9050808210156114fb578091505b61150533836144cf565b505050565b611512613ac4565b811015611558576040805162461bcd60e51b815260206004820152600f60248201526e1053149150511657d11150d2511151608a1b604482015290519081900360640190fd5b611560612cd0565b81111561148f576040805162461bcd60e51b815260206004820152600c60248201526b1113d154d39517d1561254d560a21b604482015290519081900360640190fd5b600c5481565b6001600160a01b03166000908152600a602052604090205490565b6115cc612b09565b1561160c576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b611614612b85565b600061161e613ad9565b1161165d576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f5354414b45525360b01b604482015290519081900360640190fd5b6000611667613ac4565b90506000611674826129dc565b9050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b1580156116af57600080fd5b505afa1580156116c3573d6000803e3d6000fd5b505050506116cf612b7f565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561170857600080fd5b505afa15801561171c573d6000803e3d6000fd5b505050506040513d602081101561173257600080fd5b505114611775576040805162461bcd60e51b815260206004820152600c60248201526b24a72b20a624a22fa82922ab60a11b604482015290519081900360640190fd5b611785611780612b7f565b6129dc565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b1580156117bd57600080fd5b505afa1580156117d1573d6000803e3d6000fd5b505050506117df6000613bcc565b6117f96117eb82610db0565b6117f3613ad9565b90614590565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561183257600080fd5b505afa158015611846573d6000803e3d6000fd5b505050506040513d602081101561185c57600080fd5b5051146118a1576040805162461bcd60e51b815260206004820152600e60248201526d1393d517d0531317d4d51052d15160921b604482015290519081900360640190fd5b600061191387878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808b0282810182019093528a82529093508a9250899182918501908490808284376000920191909152506145ef92505050565b905061191f81896146ef565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b15801561195857600080fd5b505afa15801561196c573d6000803e3d6000fd5b505050506040513d602081101561198257600080fd5b5051146119c5576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b60125460408051630c72684760e01b815260048101918252604481018990526001600160a01b0390921691630c726847918a918a918a918a919081906024810190606401878780828437600083820152601f01601f19169091018481038352858152602090810191508690860280828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b158015611a6f57600080fd5b505af1158015611a83573d6000803e3d6000fd5b50505050611a8f61471b565b601354604080516316b9109b60e01b81526004810186905290516001600160a01b03909216916316b9109b9160248082019260009290919082900301818387803b158015611adc57600080fd5b505af1158015611af0573d6000803e3d6000fd5b5050604080518b815290517fe54a4159af1f53fd9d722f1d91a305ea3fed5271b8ba233f16692a5cc6f01e849350908190036020019150a15050505050505050565b6016546001600160a01b03163314611b7e576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b611b86612b09565b611bce576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b601a5460ff1615611c1b576040805162461bcd60e51b8152602060048201526012602482015271414c52454144595f5452554e434154494e4760701b604482015290519081900360640190fd5b611c23612cd0565b8210611c60576040805162461bcd60e51b8152602060048201526007602482015266544f4f5f4e455760c81b604482015290519081900360640190fd5b6001611c6a613ac4565b03821015611ca9576040805162461bcd60e51b81526020600482015260076024820152661513d3d7d3d31160ca1b604482015290519081900360640190fd5b6018829055601a805460ff19166001179055611cc4816111e4565b5050565b6001600160a01b031660009081526008602052604090206001015490565b611cee612b09565b15611d2e576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b611d3733612b27565b611d75576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b80611d7f83613ce3565b14611dbe576040805162461bcd60e51b815260206004820152600a6024820152694e4f44455f52454f524760b01b604482015290519081900360640190fd5b611dc6613ac4565b8210158015611ddc5750611dd8612cd0565b8211155b611de557600080fd5b6000611df0836129dc565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611e2b57600080fd5b505afa158015611e3f573d6000803e3d6000fd5b505050506040513d6020811015611e5557600080fd5b5051611e6033611cc8565b14611ea4576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b611eb13384600c54614734565b50505050565b604b81565b611ec4612b09565b15611f04576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b6020840135843510611f4b576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b611f53612cd0565b60208501351115611f9a576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b8335611fa4612b7f565b10611fea576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b6000611ffc85825b60200201356129dc565b9050600061200b866001611ff2565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561204657600080fd5b505afa15801561205a573d6000803e3d6000fd5b505050506040513d602081101561207057600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b1580156120b257600080fd5b505afa1580156120c6573d6000803e3d6000fd5b505050506040513d60208110156120dc57600080fd5b50511461211c576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b6121368760005b60200201356001600160a01b0316614438565b612141876001612123565b604080516348b4573960e11b81526001600160a01b03893581166004830152915191841691639168ae7291602480820192602092909190829003018186803b15801561218c57600080fd5b505afa1580156121a0573d6000803e3d6000fd5b505050506040513d60208110156121b657600080fd5b50516121fe576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b604080516348b4573960e11b81526001600160a01b0360208a81013582166004840152925190841692639168ae729260248082019391829003018186803b15801561224857600080fd5b505afa15801561225c573d6000803e3d6000fd5b505050506040513d602081101561227257600080fd5b50516122ba576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b6122cf853585358560005b60200201356148cd565b826001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b15801561230857600080fd5b505afa15801561231c573d6000803e3d6000fd5b505050506040513d602081101561233257600080fd5b505114612372576040805162461bcd60e51b8152602060048201526009602482015268086908298be9082a6960bb1b604482015290519081900360640190fd5b612387602080870135908601358560016122c5565b816001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b1580156123c057600080fd5b505afa1580156123d4573d6000803e3d6000fd5b505050506040513d60208110156123ea57600080fd5b50511461242a576040805162461bcd60e51b8152602060048201526009602482015268086908298be9082a6960bb1b604482015290519081900360640190fd5b600061257b61249c846001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561246b57600080fd5b505afa15801561247f573d6000803e3d6000fd5b505050506040513d602081101561249557600080fd5b50516129dc565b6001600160a01b031663d7ff5e356040518163ffffffff1660e01b815260040160206040518083038186803b1580156124d457600080fd5b505afa1580156124e8573d6000803e3d6000fd5b505050506040513d60208110156124fe57600080fd5b5051600d546117f390818960006020020135886001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561254957600080fd5b505afa15801561255d573d6000803e3d6000fd5b505050506040513d602081101561257357600080fd5b505190614904565b6014549091506000906001600160a01b03908116906356a44dbb9030908a35908935908e35168e600160200201356001600160a01b03166125d68d6000600281106125c257fe5b60200201358a61490490919063ffffffff16565b6125f08e600160200201358b61490490919063ffffffff16565b601154604080516001600160e01b031960e08c901b1681526001600160a01b03998a166004820152602481019890985260448801969096529387166064870152918616608486015260a485015260c48401529290921660e482015290516101048083019260209291908290030181600087803b15801561266f57600080fd5b505af1158015612683573d6000803e3d6000fd5b505050506040513d602081101561269957600080fd5b505190506126c26001600160a01b038a35168a600160200201356001600160a01b031683614961565b604080518a356001600160a01b0390811682526020808d01358216908301528a35828401529151918316917fa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda758799181900360600190a2505050505050505050565b60008061272d613ac4565b9050612737612cd0565b60018203141561274b575050600f546129d9565b6000612756826129dc565b90506000816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561279357600080fd5b505afa1580156127a7573d6000803e3d6000fd5b505050506040513d60208110156127bd57600080fd5b50519050438111156127d657600f5493505050506129d9565b6127de615761565b506040805161014081018252600181526201e05b60208201526201f7d191810191909152620138916060820152620329e160808201526201be4360a08201526204cb8c60c08201526201fbc460e082015262036d326101008201526202797361012082015261284b615761565b506040805161014081018252600181526201c03060208201526201b6999181019190915261fde26060820152620265c6608082015262013b8e60a0820152620329e160c08201526201389160e08201526201f7d16101008201526201537561012082015260006128bb4385614904565b905060006128df600c546128d9600a856149ab90919063ffffffff16565b90614a04565b905060ff6128ee82600a614a04565b10612904576000199750505050505050506129d9565b600061291182600a614a04565b60020a9050600085600a8406600a811061292757fe5b602002015162ffffff168202905085600a8406600a811061294457fe5b602002015162ffffff1682828161295757fe5b04146129705760001999505050505050505050506129d9565b600061299586600a8606600a811061298457fe5b6020020151839062ffffff16614a04565b9050806129a0575060015b600f5480820290829082816129b157fe5b04146129cc576000199b5050505050505050505050506129d9565b9a50505050505050505050505b90565b6000908152600560205260409020546001600160a01b031690565b6010546001600160a01b031681565b6016546001600160a01b03163314612a52576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6012546001600160a01b0382811691161415612aa2576040805162461bcd60e51b815260206004820152600a602482015269086aaa4be9eaaa8849eb60b31b604482015290519081900360640190fd5b601154604080516319dc7ae560e31b81526001600160a01b038481166004830152600060248301819052925193169263cee3d7289260448084019391929182900301818387803b158015612af557600080fd5b505af115801561148a573d6000803e3d6000fd5b600b5460ff1690565b6014546001600160a01b031681565b600e5481565b6001600160a01b0316600090815260086020526040902060030154600160a01b900460ff1690565b600060078281548110612b5e57fe5b6000918252602090912001546001600160a01b031692915050565b60095490565b60015490565b6000612b8f613ac4565b9050612b99612b7f565b81118015612bae5750612baa612cd0565b8111155b61148f576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b6001600160a01b039081166000908152600860205260409020600301541690565b60005460ff1690565b612c21612b09565b15612c61576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b612c69612b7f565b612c7282611cc8565b1115612cb2576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b612cbb81614438565b61148f81614a68565b600f5481565b600d5481565b60035490565b612cde612b09565b15612d1e576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b612d26612b79565b821115612d6b576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b6000612d7683613153565b90506000612d8384613c49565b9050600080612d90613ac4565b90505b808310158015612da257508482105b15612e8e576000612db2846129dc565b9050806001600160a01b03166396a9fdc0866040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015612e0357600080fd5b505af1158015612e17573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015612e5457600080fd5b505afa158015612e68573d6000803e3d6000fd5b505050506040513d6020811015612e7e57600080fd5b5051935050600190910190612d93565b80831015612ea457612e9f86614abb565b612eae565b612eae8684614b57565b505050505050565b6000612ec0612b09565b15612f00576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b6000612f0b33614b7e565b9050612f178382614b9d565b92915050565b6016546001600160a01b03163314612f69576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6017546040805163266a23b160e21b815230600482018190526001600160a01b0385811660248401529251909392909216916399a88ec49160448082019260009290919082900301818387803b158015612fc257600080fd5b505af1158015612eae573d6000803e3d6000fd5b6016546001600160a01b03163314613022576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610dae614cb8565b60045490565b6016546001600160a01b031681565b6013546001600160a01b031681565b6016546001600160a01b0316331461309a576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601754604051639623609d60e01b815230600482018181526001600160a01b0387811660248501526060604485019081526064850187905292941692639623609d92859289928992899291608401848480828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b15801561312657600080fd5b505af115801561313a573d6000803e3d6000fd5b5050505050505050565b6012546001600160a01b031681565b60006009828154811061316257fe5b60009182526020909120600290910201546001600160a01b031692915050565b61318a612b09565b156131ca576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b6131d333612b27565b613211576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b600061321b612cd0565b60010190506000808080808061323361178033611cc8565b905061323d615780565b604080516080818101909252613291918c906004908390839080828437600092019190915250506040805161014081810190925291508c90600a908390839080828437600092019190915250614d36915050565b905061329c81614df1565b9250816001600160a01b031663701da98e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156132d757600080fd5b505afa1580156132eb573d6000803e3d6000fd5b505050506040513d602081101561330157600080fd5b505161330c82614e65565b14613350576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b8051600090613360904390614904565b9050604b8110156133a5576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b606082015160c08301516133b891614904565b8260e001511015806133dc5750600e546133d39082906149ab565b82610100015110155b806133ec57506101408201516064145b613429576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b6134496004613443600e54846149ab90919063ffffffff16565b906149ab565b826101000151111561348e576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b600e546000906134b2906128d96134a6826001614904565b61010087015190614590565b905061353a816117f36134d0600c544361459090919063ffffffff16565b876001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561350957600080fd5b505afa15801561351d573d6000803e3d6000fd5b505050506040513d602081101561353357600080fd5b5051614e92565b6013549099506001600160a01b03169050638b8ca1998a61355a33611cc8565b8b336040518563ffffffff1660e01b815260040180858152602001848152602001838152602001826001600160a01b03168152602001945050505050600060405180830381600087803b1580156135b057600080fd5b505af11580156135c4573d6000803e3d6000fd5b50505050601160009054906101000a90046001600160a01b03166001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b15801561361657600080fd5b505afa15801561362a573d6000803e3d6000fd5b505050506040513d602081101561364057600080fd5b505160e0830151606084015191985060009161365b91614590565b9050878111156136a3576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b8015613727576011546040805163d9dd67ab60e01b81526000198401600482015290516001600160a01b039092169163d9dd67ab91602480820192602092909190829003018186803b1580156136f857600080fd5b505afa15801561370c573d6000803e3d6000fd5b505050506040513d602081101561372257600080fd5b505196505b6015546001600160a01b031663d45ab2b5613742858b614ea8565b61374d868943614ee9565b61375687614f08565b61375f33611cc8565b8e6040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b1580156137b257600080fd5b505af11580156137c6573d6000803e3d6000fd5b505050506040513d60208110156137dc57600080fd5b50516040805163f0dd77ff60e01b81529051919750600094508493506001600160a01b038616925063f0dd77ff916004808301926020929190829003018186803b15801561382957600080fd5b505afa15801561383d573d6000803e3d6000fd5b505050506040513d602081101561385357600080fd5b5051905080158015906138d4576138cd846001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b15801561389c57600080fd5b505afa1580156138b0573d6000803e3d6000fd5b505050506040513d60208110156138c657600080fd5b5051613ce3565b9250613913565b613910866001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561389c57600080fd5b92505b60006139218285888b614f1e565b90508d811461396e576040805162461bcd60e51b81526020600482015260146024820152730aa9c8ab0a08a86a88a88be9c9e888abe9082a6960631b604482015290519081900360640190fd5b6139788782614f82565b846001600160a01b0316631bc09d0a8c6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156139be57600080fd5b505af11580156139d2573d6000803e3d6000fd5b50505050505050506139e73388600c54614734565b50613a24836001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561389c57600080fd5b877fa47903165cab28e7f468037b6eb8201a4b1cf821da83143b7b6907e342f9dd2e613a4f8a613ce3565b8589898f8f6040518087815260200186815260200185815260200184815260200183600460200280828437600083820152601f01601f191690910190508261014080828437600083820152604051601f909101601f19169092018290039850909650505050505050a350505050505050505050565b60025490565b6015546001600160a01b031681565b60075490565b6016546001600160a01b03163314613b2b576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6011546040805163722dbe7360e11b81526001600160a01b03858116600483015284151560248301529151919092169163e45b7ce691604480830192600092919082900301818387803b158015612fc257600080fd5b6011546001600160a01b031681565b6008602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6000613bd6612b79565b90506000613be2613ac4565b9050825b82811015611eb1575b81613bf982613c49565b1015613c2357613c0881614abb565b60001990920191828110613c1e5750505061148f565b613bef565b600101613be6565b6001600160a01b031660009081526008602052604090206002015490565b600060098281548110613c5857fe5b9060005260206000209060020201600101549050919050565b613c79612b09565b15613cb9576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b613cc282614438565b611cc482613ccf83614fcc565b61513a565b6017546001600160a01b031681565b60009081526006602052604090205490565b613cff828261516b565b6001600160a01b0316336001600160a01b031614613d53576040805162461bcd60e51b815260206004820152600c60248201526b2ba927a723afa9a2a72222a960a11b604482015290519081900360640190fd5b6000613d5e82613c2b565b90506000613d6b84613c2b565b905080821115613d8c57613d89613d8284836144cf565b8390614904565b91505b60028204613d9a858261513a565b613da48382614904565b9250613daf8561522d565b601654613dc5906001600160a01b03168461513a565b61148a84615257565b613dd6612b09565b15613e16576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b613e1f33612b27565b15613e62576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b6000613e6d82614fcc565b9050613e77612722565b811015613ebe576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b613ec83382615307565b6013546001600160a01b031663f03c04a533613ee2612b7f565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015612fc257600080fd5b600c5415613f6c576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b88613fb0576040805162461bcd60e51b815260206004820152600f60248201526e10905117d0d3d39197d411549253d1608a1b604482015290519081900360640190fd5b6011805460208301356001600160a01b039081166001600160a01b0319928316179283905560128054909216604080860135831691821790935582516319dc7ae560e31b8152600481019190915260016024820152915192169163cee3d7289160448082019260009290919082900301818387803b15801561403157600080fd5b505af1158015614045573d6000803e3d6000fd5b505050508060036006811061405657fe5b601380546001600160a01b0319166001600160a01b0360209390930293909301358216929092179091556011546040805163722dbe7360e11b8152606085013584166004820152600160248201529051919092169163e45b7ce691604480830192600092919082900301818387803b1580156140d157600080fd5b505af11580156140e5573d6000803e3d6000fd5b505060135460405163b0f2af2960e01b8152600481018d8152602482018d9052604482018c9052606482018b90526001600160a01b038a8116608484015289811660a484015260e060c4840190815260e484018990529316945063b0f2af2993508d928d928d928d928d928d928d928d9261010401848480828437600081840152601f19601f8201169050808301925050509950505050505050505050600060405180830381600087803b15801561419c57600080fd5b505af11580156141b0573d6000803e3d6000fd5b50505050806004600681106141c157fe5b601480546001600160a01b03199081166001600160a01b03602094909402949094013583169390931790556015805490921660a084013590911617905560006142098b6153d2565b905061421481615489565b600c8a9055600d899055600e889055600f879055601080546001600160a01b03199081166001600160a01b0389811691909117909255601680548216888416179055601780549091168435909216919091179055604080518c815290517f4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d9181900360200190a15050505050505050505050565b6016546001600160a01b031633146142f4576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601280546001600160a01b0319166001600160a01b03838116918217909255601154604080516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b158015612af557600080fd5b614366612b09565b6143ae576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b600b805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6143e16154d8565b604080516001600160a01b039092168252519081900360200190a1565b6144096002546154dc565b600280546001019055565b6001600160a01b03909116600090815260086020526040902060010155565b600355565b61444181612b27565b61447f576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b600061448a82612bef565b6001600160a01b03161461148f576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600860205260408120600281015480841115614534576040805162461bcd60e51b815260206004820152601060248201526f544f4f5f4c4954544c455f5354414b4560801b604482015290519081900360640190fd5b60006145408286614904565b600284018690556001600160a01b0387166000908152600a602052604090205490915061456d9082614590565b6001600160a01b0387166000908152600a60205260409020559250505092915050565b6000828201838110156145e8576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b80518251600091829182805b838110156146a257600087828151811061461157fe5b60200260200101519050838187011115614661576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868a01810182902060408051808401969096528581019190915280518086038201815260609095019052835193019290922091909401936001016145fb565b508184146146e5576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b9695505050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6147266001546154dc565b600280546001818155019055565b6001600160a01b0380841660008181526008602090815260408083208784526005835281842054825163123334b760e11b8152600481019690965291519395909491169285928492632466696e926024808301939282900301818787803b15801561479e57600080fd5b505af11580156147b2573d6000803e3d6000fd5b505050506040513d60208110156147c857600080fd5b505160018085018890559091508114156148c357600060056000846001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561481b57600080fd5b505afa15801561482f573d6000803e3d6000fd5b505050506040513d602081101561484557600080fd5b505181526020810191909152604001600020546001600160a01b0316905080636971dfe56148734389614590565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156148a957600080fd5b505af11580156148bd573d6000803e3d6000fd5b50505050505b5050509392505050565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b60008282111561495b576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6001600160a01b03928316600090815260086020526040808220600390810180549487166001600160a01b0319958616811790915594909516825290209092018054909216179055565b6000826149ba57506000612f17565b828202828482816149c757fe5b04146145e85760405162461bcd60e51b81526004018080602001828103825260218152602001806157fc6021913960400191505060405180910390fd5b6000808211614a57576040805162461bcd60e51b815260206004820152601a602482015279536166654d6174683a206469766973696f6e206279207a65726f60301b604482015290519081900360640190fd5b818381614a6057fe5b049392505050565b6001600160a01b03811660009081526008602090815260408083206002810154600a909352922054614a9991614590565b6001600160a01b0383166000908152600a6020526040902055611cc48261555e565b600980546000198101908110614acd57fe5b906000526020600020906002020160098281548110614ae857fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b039092169190911781556001918201549101556009805480614b2b57fe5b60008281526020812060026000199093019283020180546001600160a01b031916815560010155905550565b8060098381548110614b6557fe5b9060005260206000209060020201600101819055505050565b6001600160a01b03166000908152600a60205260408120805491905590565b80614ba757611cc4565b6010546001600160a01b0316614bf3576040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015614bed573d6000803e3d6000fd5b50611cc4565b6010546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015614c4957600080fd5b505af1158015614c5d573d6000803e3d6000fd5b505050506040513d6020811015614c7357600080fd5b5051611cc4576040805162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b604482015290519081900360640190fd5b614cc0612b09565b15614d00576040805162461bcd60e51b8152602060048201526010602482015260008051602061581d833981519152604482015290519081900360640190fd5b600b805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586143e16154d8565b614d3e615780565b604080516101c08101825283518152602080850151818301528551828401528483015160608084019190915285015160808084019190915285015160a08084019190915285015160c08084019190915285015160e08084019190915285015161010080840191909152908601516101208301528401516101408201529084015161016082015261018081018360096020020151815260200184600360048110614de357fe5b602002015190529392505050565b6000612f176000836101000151614e266000614e21876060015188604001516000801b60008060001b6000615684565b6146ef565b614e60866101000151614e218860e00151896060015101896101a001518a61012001518b61014001518c61016001518d6101800151615684565b6156cf565b6000612f17826000015183602001518460400151856060015186608001518760a001518860c0015161570d565b6000818311614ea157816145e8565b5090919050565b60006145e843846101000151856020015101856101a001518660e001518760600151018761014001518860800151018861018001518960a00151018861570d565b6000614f0083838660e001518760600151016148cd565b949350505050565b6000612f178261012001518361016001516146ef565b60008085614f2d576000614f30565b60015b905080858585604051602001808560ff1660f81b815260010184815260200183815260200182815260200194505050505060405160208183030381529060405280519060200120915050949350505050565b60038054600101808255600090815260056020908152604080832080546001600160a01b0319166001600160a01b0397909716969096179095559154815260069091529190912055565b6010546000906001600160a01b0316615028578115615021576040805162461bcd60e51b815260206004820152600c60248201526b4241445f53544b5f5459504560a01b604482015290519081900360640190fd5b5034610e63565b341561506a576040805162461bcd60e51b815260206004820152600c60248201526b4241445f53544b5f5459504560a01b604482015290519081900360640190fd5b601054604080516323b872dd60e01b81523360048201523060248201526044810185905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156150c457600080fd5b505af11580156150d8573d6000803e3d6000fd5b505050506040513d60208110156150ee57600080fd5b5051615133576040805162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b604482015290519081900360640190fd5b5080610e63565b6001600160a01b038216600090815260086020526040902060028101546151619083614590565b6002909101555050565b6001600160a01b03808316600090815260086020526040808220848416835290822060038083015490820154939492939192908116911681146151df576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b038116615224576040805162461bcd60e51b81526020600482015260076024820152661393d7d0d2105360ca1b604482015290519081900360640190fd5b95945050505050565b6001600160a01b0316600090815260086020526040902060030180546001600160a01b0319169055565b6001600160a01b0381811660008181526008602090815260408083208151808301909252938152600180850154928201928352600980549182018155909352517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af600290930292830180546001600160a01b031916919095161790935591517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7b090920191909155611cc48261555e565b6007805460018082019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688810180546001600160a01b039586166001600160a01b031991821681179092556040805160a08101825293845284546020858101918252858301978852600060608701818152608088018981529682526008909252929092209451855551948401949094559351600283015591516003909101805492511515600160a01b0260ff60a01b199290951692909316919091171691909117905543600455565b6000806153e8436000856000806000600161570d565b6015546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905243608483015291519394506001600160a01b039092169263d45ab2b59260a4808201936020939283900390910190829087803b15801561545657600080fd5b505af115801561546a573d6000803e3d6000fd5b505050506040513d602081101561548057600080fd5b50519392505050565b6000805260056020527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc80546001600160a01b0319166001600160a01b03929092169190911790556001600255565b3390565b60008181526005602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b15801561552857600080fd5b505af115801561553c573d6000803e3d6000fd5b50505060009182525060056020526040902080546001600160a01b0319169055565b6001600160a01b0381166000908152600860205260409020805460078054600019810190811061558a57fe5b600091825260209091200154600780546001600160a01b0390921691839081106155b057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508060086000600784815481106155f057fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600780548061562057fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03949094168152600890935250506040812081815560018101829055600281019190915560030180546001600160a81b0319169055565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b60408051602080820199909952808201979097526060870195909552608086019390935260a085019190915260c084015260e080840191909152815180840390910181526101009092019052805191012090565b604051806101400160405280600a906020820280368337509192915050565b604051806101c001604052806000815260200160008152602001600080191681526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000801916815260200160008152602001600080191681526020016000815260200160008019168152509056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775061757361626c653a2070617573656400000000000000000000000000000000a2646970667358221220aad62d5fe483206ca558c966caee93c5221dcd517db3a9395d4865613749debe64736f6c634300060c0033"

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

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_Rollup *RollupCaller) ConfirmPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "confirmPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_Rollup *RollupSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _Rollup.Contract.ConfirmPeriodBlocks(&_Rollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_Rollup *RollupCallerSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _Rollup.Contract.ConfirmPeriodBlocks(&_Rollup.CallOpts)
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

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_Rollup *RollupCaller) ExtraChallengeTimeBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "extraChallengeTimeBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_Rollup *RollupSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _Rollup.Contract.ExtraChallengeTimeBlocks(&_Rollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_Rollup *RollupCallerSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _Rollup.Contract.ExtraChallengeTimeBlocks(&_Rollup.CallOpts)
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

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_Rollup *RollupCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_Rollup *RollupSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _Rollup.Contract.GetNodeHash(&_Rollup.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_Rollup *RollupCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _Rollup.Contract.GetNodeHash(&_Rollup.CallOpts, index)
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

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Rollup *RollupCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Rollup *RollupSession) IsMaster() (bool, error) {
	return _Rollup.Contract.IsMaster(&_Rollup.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Rollup *RollupCallerSession) IsMaster() (bool, error) {
	return _Rollup.Contract.IsMaster(&_Rollup.CallOpts)
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

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_Rollup *RollupTransactor) CreateChallenge(opts *bind.TransactOpts, stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "createChallenge", stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_Rollup *RollupSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_Rollup *RollupTransactorSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_Rollup *RollupSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdaf5797.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts) returns()
func (_Rollup *RollupTransactorSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [6]common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
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

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_Rollup *RollupTransactor) StakeOnExistingNode(opts *bind.TransactOpts, nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "stakeOnExistingNode", nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_Rollup *RollupSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnExistingNode(&_Rollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_Rollup *RollupTransactorSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnExistingNode(&_Rollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0xd50b580b.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[4] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupTransactor) StakeOnNewNode(opts *bind.TransactOpts, expectedNodeHash [32]byte, assertionBytes32Fields [4][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "stakeOnNewNode", expectedNodeHash, assertionBytes32Fields, assertionIntFields)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0xd50b580b.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[4] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [4][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnNewNode(&_Rollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0xd50b580b.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[4] assertionBytes32Fields, uint256[10] assertionIntFields) returns()
func (_Rollup *RollupTransactorSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [4][32]byte, assertionIntFields [10]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnNewNode(&_Rollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields)
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
	ParentNodeHash         [32]byte
	NodeHash               [32]byte
	ExecutionHash          [32]byte
	InboxMaxCount          *big.Int
	AfterInboxAcc          [32]byte
	AssertionBytes32Fields [4][32]byte
	AssertionIntFields     [10]*big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0xa47903165cab28e7f468037b6eb8201a4b1cf821da83143b7b6907e342f9dd2e.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, bytes32 afterInboxAcc, bytes32[4] assertionBytes32Fields, uint256[10] assertionIntFields)
func (_Rollup *RollupFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int, parentNodeHash [][32]byte) (*RollupNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupNodeCreatedIterator{contract: _Rollup.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0xa47903165cab28e7f468037b6eb8201a4b1cf821da83143b7b6907e342f9dd2e.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, bytes32 afterInboxAcc, bytes32[4] assertionBytes32Fields, uint256[10] assertionIntFields)
func (_Rollup *RollupFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *RollupNodeCreated, nodeNum []*big.Int, parentNodeHash [][32]byte) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
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

// ParseNodeCreated is a log parse operation binding the contract event 0xa47903165cab28e7f468037b6eb8201a4b1cf821da83143b7b6907e342f9dd2e.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, bytes32 afterInboxAcc, bytes32[4] assertionBytes32Fields, uint256[10] assertionIntFields)
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
const RollupCoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupCoreBin is the compiled bytecode used for deploying new contracts.
var RollupCoreBin = "0x608060405234801561001057600080fd5b506104e1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100db5760003560e01c80632f30cabd146100e05780633e96576e146101185780634f0f4aa91461013e5780636177fd181461017757806362a82d7d146101b157806363721d6b146101ce57806365f7f80d146101d657806369fd251c146101de5780637ba9534a146102045780638640ce5f1461020c578063d01e660214610214578063d735e21d14610231578063dff6978714610239578063e8bd492214610241578063ef40a6701461029d578063f33e1fac146102c3578063f8d1f194146102e0575b600080fd5b610106600480360360208110156100f657600080fd5b50356001600160a01b03166102fd565b60408051918252519081900360200190f35b6101066004803603602081101561012e57600080fd5b50356001600160a01b0316610318565b61015b6004803603602081101561015457600080fd5b5035610336565b604080516001600160a01b039092168252519081900360200190f35b61019d6004803603602081101561018d57600080fd5b50356001600160a01b0316610351565b604080519115158252519081900360200190f35b61015b600480360360208110156101c757600080fd5b5035610379565b6101066103a3565b6101066103a9565b61015b600480360360208110156101f457600080fd5b50356001600160a01b03166103af565b6101066103d0565b6101066103d6565b61015b6004803603602081101561022a57600080fd5b50356103dc565b61010661040b565b610106610411565b6102676004803603602081101561025757600080fd5b50356001600160a01b0316610417565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b610106600480360360208110156102b357600080fd5b50356001600160a01b0316610453565b610106600480360360208110156102d957600080fd5b5035610471565b610106600480360360208110156102f657600080fd5b5035610499565b6001600160a01b031660009081526009602052604090205490565b6001600160a01b031660009081526007602052604090206001015490565b6000908152600460205260409020546001600160a01b031690565b6001600160a01b0316600090815260076020526040902060030154600160a01b900460ff1690565b60006006828154811061038857fe5b6000918252602090912001546001600160a01b031692915050565b60085490565b60005490565b6001600160a01b039081166000908152600760205260409020600301541690565b60025490565b60035490565b6000600882815481106103eb57fe5b60009182526020909120600290910201546001600160a01b031692915050565b60015490565b60065490565b6007602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6001600160a01b031660009081526007602052604090206002015490565b60006008828154811061048057fe5b9060005260206000209060020201600101549050919050565b6000908152600560205260409020549056fea26469706673582212203d622e530eae60cf75a0379e072cec59eb48a1a77e33ae08e3c2512e90e130f564736f6c634300060c0033"

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

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _RollupCore.Contract.GetNodeHash(&_RollupCore.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _RollupCore.Contract.GetNodeHash(&_RollupCore.CallOpts, index)
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
const RollupEventBridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"claimNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"nodeConfirmed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"}],\"name\":\"nodeCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"nodeRejected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraConfig\",\"type\":\"bytes\"}],\"name\":\"rollupInitialized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"stakeCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupEventBridgeBin is the compiled bytecode used for deploying new contracts.
var RollupEventBridgeBin = "0x608060405234801561001057600080fd5b50604051610a0a380380610a0a8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556109908061007a6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806316b9109b1461006757806330a826b41461008657806364126c7c146100a35780638b8ca199146100cf578063b0f2af2914610107578063f03c04a5146101a6575b600080fd5b6100846004803603602081101561007d57600080fd5b50356101d2565b005b6100846004803603602081101561009c57600080fd5b5035610253565b610084600480360360408110156100b957600080fd5b50803590602001356001600160a01b03166102d1565b610084600480360360808110156100e557600080fd5b50803590602081013590604081013590606001356001600160a01b03166104f8565b610084600480360360e081101561011d57600080fd5b8135916020810135916040820135916060810135916001600160a01b03608083013581169260a08101359091169181019060e0810160c0820135600160201b81111561016857600080fd5b82018360208201111561017a57600080fd5b803590602001918460018302840111600160201b8311171561019b57600080fd5b509092509050610599565b610084600480360360408110156101bc57600080fd5b506001600160a01b038135169060200135610786565b6001546001600160a01b0316331461021f576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f81b6020820152602180820184905282518083039091018152604190910190915261025090610820565b50565b6001546001600160a01b031633146102a0576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f91b6020820152602180820184905282518083039091018152604190910190915261025090610820565b6001546001600160a01b0316331461031e576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60015460408051634f0f4aa960e01b81526004810185905290516001600160a01b03909216916000918391634f0f4aa991602480820192602092909190829003018186803b15801561036f57600080fd5b505afa158015610383573d6000803e3d6000fd5b505050506040513d602081101561039957600080fd5b5051604080516348b4573960e11b81526001600160a01b038681166004830152915192935090831691639168ae7291602480820192602092909190829003018186803b1580156103e857600080fd5b505afa1580156103fc573d6000803e3d6000fd5b505050506040513d602081101561041257600080fd5b5051610452576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b816001600160a01b0316632b2af0ab856040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561049657600080fd5b505afa1580156104aa573d6000803e3d6000fd5b505060408051600160fa1b6020820152602181018890526001600160a01b0387166041808301919091528251808303909101815260619091019091526104f292509050610820565b50505050565b6001546001600160a01b03163314610545576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600060208201526021810186905260418101859052436061820152608181018490526001600160a01b03831660a1808301919091528251808303909101815260c19091019091526104f290610820565b6001546001600160a01b031633146105e6576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b6060888888888860601b60601c6001600160a01b03168860601b60601c6001600160a01b03168888604051602001808981526020018881526020018781526020018681526020018581526020018481526020018383808284376040805191909301818103601f190182528084526000805483516020808601919091206302bbfad160e01b855260048086015233602486015260448501529551939f50909d506001600160a01b03169b506302bbfad19a5060648082019a509398509096508690039091019350849250899150889050803b1580156106c357600080fd5b505af11580156106d7573d6000803e3d6000fd5b505050506040513d60208110156106ed57600080fd5b50516040805160208082528551828201528551939450849360008051602061093b833981519152938793928392918301919085019080838360005b83811015610740578181015183820152602001610728565b50505050905090810190601f16801561076d5780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050505050505050565b6001546001600160a01b031633146107d3576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600360f81b60208201526001600160a01b0384166021820152604181018390524360618083019190915282518083039091018152608190910190915261081c90610820565b5050565b600080548251602080850191909120604080516302bbfad160e01b8152600860048201523360248201526044810192909252516001600160a01b03909316936302bbfad193606480840194939192918390030190829087803b15801561088557600080fd5b505af1158015610899573d6000803e3d6000fd5b505050506040513d60208110156108af57600080fd5b5051604080516020808252845182820152845160008051602061093b833981519152938693928392918301919085019080838360005b838110156108fd5781810151838201526020016108e5565b50505050905090810190601f16801561092a5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25056feff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60ba26469706673582212202b3e74b50b6e56e9b567846ff040a187a2a10f63b3281a0a40f1656215d489ae64736f6c634300060c0033"

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

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) RollupInitialized(opts *bind.TransactOpts, confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "rollupInitialized", confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeSession) RollupInitialized(confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupInitialized(&_RollupEventBridge.TransactOpts, confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) RollupInitialized(confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupInitialized(&_RollupEventBridge.TransactOpts, confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
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

// RollupEventBridgeInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredIterator struct {
	Event *RollupEventBridgeInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupEventBridgeInboxMessageDelivered)
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
		it.Event = new(RollupEventBridgeInboxMessageDelivered)
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
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupEventBridgeInboxMessageDelivered represents a InboxMessageDelivered event raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*RollupEventBridgeInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeInboxMessageDeliveredIterator{contract: _RollupEventBridge.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *RollupEventBridgeInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupEventBridgeInboxMessageDelivered)
				if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) ParseInboxMessageDelivered(log types.Log) (*RollupEventBridgeInboxMessageDelivered, error) {
	event := new(RollupEventBridgeInboxMessageDelivered)
	if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupEventBridgeInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredFromOriginIterator struct {
	Event *RollupEventBridgeInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
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
		it.Event = new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
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
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupEventBridgeInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*RollupEventBridgeInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeInboxMessageDeliveredFromOriginIterator{contract: _RollupEventBridge.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *RollupEventBridgeInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
				if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*RollupEventBridgeInboxMessageDeliveredFromOrigin, error) {
	event := new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
	if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupLibABI is the input ABI used to generate the binding from.
const RollupLibABI = "[]"

// RollupLibBin is the compiled bytecode used for deploying new contracts.
var RollupLibBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f557ef6ecd29ec1f530f2de54d6eaf392524ad1b2a782078f3ad6bea971a7c0464736f6c634300060c0033"

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
const ValidatorUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkDecidableNextNode\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.ConfirmType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"node1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findNodeConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findStakerConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"refundableStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"requireConfirmable\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"requireRejectableNextNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakedNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"stakerInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorUtilsBin is the compiled bytecode used for deploying new contracts.
var ValidatorUtilsBin = "0x608060405234801561001057600080fd5b50612869806100206000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c806301d9717d1461009e5780631fc43bb6146100c85780632b1062cf146100dd5780633082d029146100ff5780637464ae06146101215780637988ad37146101415780638f67e6bb14610154578063974dbf2a14610177578063abeba4f714610198578063c308eaaf146101ab578063e48a5f7b146101cb575b600080fd5b6100b16100ac366004612446565b6101ef565b6040516100bf9291906127e1565b60405180910390f35b6100db6100d636600461242a565b610377565b005b6100f06100eb36600461253c565b6108f9565b6040516100bf93929190612690565b61011261010d366004612502565b6109fd565b6040516100bf939291906126bb565b61013461012f36600461242a565b610ece565b6040516100bf91906125b7565b61011261014f36600461247e565b611154565b610167610162366004612446565b611276565b6040516100bf949392919061263c565b61018a61018536600461253c565b61147c565b6040516100bf9291906127ca565b6101346101a63660046124ce565b611561565b6101be6101b9366004612446565b6116eb565b6040516100bf9190612604565b6101de6101d936600461242a565b61193a565b6040516100bf9594939291906127ef565b6000806000846001600160a01b0316633e96576e856040518263ffffffff1660e01b815260040161022091906125a3565b60206040518083038186803b15801561023857600080fd5b505afa15801561024c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102709190612412565b9050806102eb57846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156102b057600080fd5b505afa1580156102c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e89190612412565b90505b604051633e347c6560e21b81526000906001600160a01b0387169063f8d1f1949061031a9085906004016127c1565b60206040518083038186803b15801561033257600080fd5b505afa158015610346573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036a9190612412565b9196919550909350505050565b806001600160a01b03166367425daf6040518163ffffffff1660e01b815260040160006040518083038186803b1580156103b057600080fd5b505afa1580156103c4573d6000803e3d6000fd5b505050506000816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561040357600080fd5b505afa158015610417573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043b9190612412565b9050600081116104665760405162461bcd60e51b815260040161045d90612702565b60405180910390fd5b6000826001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156104a157600080fd5b505afa1580156104b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d99190612412565b90506000836001600160a01b0316634f0f4aa9836040518263ffffffff1660e01b815260040161050991906127c1565b60206040518083038186803b15801561052157600080fd5b505afa158015610535573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061055991906123d6565b9050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b15801561059457600080fd5b505afa1580156105a8573d6000803e3d6000fd5b50505050836001600160a01b0316634f0f4aa9826001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156105f457600080fd5b505afa158015610608573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062c9190612412565b6040518263ffffffff1660e01b815260040161064891906127c1565b60206040518083038186803b15801561066057600080fd5b505afa158015610674573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069891906123d6565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b1580156106d057600080fd5b505afa1580156106e4573d6000803e3d6000fd5b50505050836001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561072157600080fd5b505afa158015610735573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107599190612412565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561079257600080fd5b505afa1580156107a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ca9190612412565b146107e75760405162461bcd60e51b815260040161045d9061274e565b604051630128a01960e21b81526001600160a01b038516906304a28064906108139084906004016125a3565b60206040518083038186803b15801561082b57600080fd5b505afa15801561083f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108639190612412565b8301816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561089e57600080fd5b505afa1580156108b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d69190612412565b146108f35760405162461bcd60e51b815260040161045d90612726565b50505050565b604051630fe21ddb60e11b8152600090819081903090631fc43bb690610923908b906004016125a3565b60006040518083038186803b15801561093b57600080fd5b505afa92505050801561094c575060015b61095557610964565b506001915060009050806109f2565b604051634ba6df9560e11b8152309063974dbf2a9061098f908b908b908b908b908b90600401612662565b604080518083038186803b1580156109a657600080fd5b505afa9250505080156109d6575060408051601f3d908101601f191682019092526109d39181019061257f565b60015b6109e8575060009150819050806109f2565b6002945090925090505b955095509592505050565b600080600080876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a3c57600080fd5b505afa158015610a50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a749190612412565b90506000886001600160a01b0316634f0f4aa9896040518263ffffffff1660e01b8152600401610aa491906127c1565b60206040518083038186803b158015610abc57600080fd5b505afa158015610ad0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af491906123d6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610b2c57600080fd5b505afa158015610b40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b649190612412565b90506000896001600160a01b0316634f0f4aa9896040518263ffffffff1660e01b8152600401610b9491906127c1565b60206040518083038186803b158015610bac57600080fd5b505afa158015610bc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be491906123d6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610c1c57600080fd5b505afa158015610c30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c549190612412565b905060005b87811015610eb557888a1415610c7c5760008a8a96509650965050505050610ec4565b81831415610c975760018a8a96509650965050505050610ec4565b83831080610ca457508382105b15610cbd57600260008096509650965050505050610ec4565b888a1015610dbb578198508a6001600160a01b0316634f0f4aa98a6040518263ffffffff1660e01b8152600401610cf491906127c1565b60206040518083038186803b158015610d0c57600080fd5b505afa158015610d20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4491906123d6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610d7c57600080fd5b505afa158015610d90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db49190612412565b9150610ead565b8299508a6001600160a01b0316634f0f4aa98b6040518263ffffffff1660e01b8152600401610dea91906127c1565b60206040518083038186803b158015610e0257600080fd5b505afa158015610e16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3a91906123d6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610e7257600080fd5b505afa158015610e86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eaa9190612412565b92505b600101610c59565b50600389899550955095505050505b9450945094915050565b60606000826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015610f0b57600080fd5b505afa158015610f1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f439190612412565b90506060816001600160401b0381118015610f5d57600080fd5b50604051908082528060200260200182016040528015610f87578160200160208202803683370190505b5090506000846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610fc557600080fd5b505afa158015610fd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffd9190612412565b90506000805b84811015611149576040516362a82d7d60e01b81526000906001600160a01b038916906362a82d7d9061103a9085906004016127c1565b60206040518083038186803b15801561105257600080fd5b505afa158015611066573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108a91906123d6565b90506000886001600160a01b0316633e96576e836040518263ffffffff1660e01b81526004016110ba91906125a3565b60206040518083038186803b1580156110d257600080fd5b505afa1580156110e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110a9190612412565b905084811161113f578186858151811061112057fe5b6001600160a01b03909216602092830291909101909101526001909301925b5050600101611003565b508252509392505050565b600080600080876001600160a01b0316633e96576e886040518263ffffffff1660e01b815260040161118691906125a3565b60206040518083038186803b15801561119e57600080fd5b505afa1580156111b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111d69190612412565b90506000886001600160a01b0316633e96576e886040518263ffffffff1660e01b815260040161120691906125a3565b60206040518083038186803b15801561121e57600080fd5b505afa158015611232573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112569190612412565b9050611264898383896109fd565b94509450945050509450945094915050565b600080600080856001600160a01b0316636177fd18866040518263ffffffff1660e01b81526004016112a891906125a3565b60206040518083038186803b1580156112c057600080fd5b505afa1580156112d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f891906123f2565b604051631f4b2bb760e11b81526001600160a01b03881690633e96576e906113249089906004016125a3565b60206040518083038186803b15801561133c57600080fd5b505afa158015611350573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113749190612412565b604051630ef40a6760e41b81526001600160a01b0389169063ef40a670906113a0908a906004016125a3565b60206040518083038186803b1580156113b857600080fd5b505afa1580156113cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f09190612412565b604051631a7f494760e21b81526001600160a01b038a16906369fd251c9061141c908b906004016125a3565b60206040518083038186803b15801561143457600080fd5b505afa158015611448573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061146c91906123d6565b9299919850965090945092505050565b600080600061148a88611b8a565b9050801561149f576000809250925050611557565b6000886001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156114da57600080fd5b505afa1580156114ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115129190612412565b9050600080600061152b8c8c86600101018c8c8c612097565b9250925092508261154e5760405162461bcd60e51b815260040161045d90612774565b90955093505050505b9550959350505050565b60606000846001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561159e57600080fd5b505afa1580156115b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d69190612412565b90508083850110156115e757508282015b6060816001600160401b03811180156115ff57600080fd5b50604051908082528060200260200182016040528015611629578160200160208202803683370190505b50905060005b828110156116e1576040516362a82d7d60e01b81526001600160a01b038816906362a82d7d9061166590898501906004016127c1565b60206040518083038186803b15801561167d57600080fd5b505afa158015611691573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116b591906123d6565b8282815181106116c157fe5b6001600160a01b039092166020928302919091019091015260010161162f565b5095945050505050565b60408051620186a08082526230d4208201909252606091829190602082016230d40080368337019050509050600080856001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561175357600080fd5b505afa158015611767573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061178b9190612412565b90505b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117c757600080fd5b505afa1580156117db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117ff9190612412565b811161193057604051634f0f4aa960e01b81526000906001600160a01b03881690634f0f4aa9906118349085906004016127c1565b60206040518083038186803b15801561184c57600080fd5b505afa158015611860573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061188491906123d6565b6040516348b4573960e11b81529091506001600160a01b03821690639168ae72906118b39089906004016125a3565b60206040518083038186803b1580156118cb57600080fd5b505afa1580156118df573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061190391906123f2565b15611927578184848151811061191557fe5b60209081029190910101526001909201915b5060010161178e565b5081529392505050565b6000806000806000856001600160a01b0316632e7acfa66040518163ffffffff1660e01b815260040160206040518083038186803b15801561197b57600080fd5b505afa15801561198f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119b39190612412565b9450856001600160a01b031663771b2f976040518163ffffffff1660e01b815260040160206040518083038186803b1580156119ee57600080fd5b505afa158015611a02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a269190612412565b9350856001600160a01b0316635e8ef1066040518163ffffffff1660e01b815260040160206040518083038186803b158015611a6157600080fd5b505afa158015611a75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a999190612412565b9250856001600160a01b03166376e7e23b6040518163ffffffff1660e01b815260040160206040518083038186803b158015611ad457600080fd5b505afa158015611ae8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b0c9190612412565b9150856001600160a01b03166351ed6a306040518163ffffffff1660e01b815260040160206040518083038186803b158015611b4757600080fd5b505afa158015611b5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b7f91906123d6565b905091939590929450565b6000816001600160a01b03166367425daf6040518163ffffffff1660e01b815260040160006040518083038186803b158015611bc557600080fd5b505afa158015611bd9573d6000803e3d6000fd5b505050506000826001600160a01b0316634f0f4aa9846001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c2757600080fd5b505afa158015611c3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c5f9190612412565b6040518263ffffffff1660e01b8152600401611c7b91906127c1565b60206040518083038186803b158015611c9357600080fd5b505afa158015611ca7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ccb91906123d6565b90506000836001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611d0857600080fd5b505afa158015611d1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d409190612412565b826001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611d7957600080fd5b505afa158015611d8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611db19190612412565b149050801561209057816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611df357600080fd5b505afa158015611e07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e2b9190612412565b431015611e4a5760405162461bcd60e51b815260040161045d90612798565b836001600160a01b0316634f0f4aa9836001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611e9257600080fd5b505afa158015611ea6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eca9190612412565b6040518263ffffffff1660e01b8152600401611ee691906127c1565b60206040518083038186803b158015611efe57600080fd5b505afa158015611f12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f3691906123d6565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b158015611f6e57600080fd5b505afa158015611f82573d6000803e3d6000fd5b5050604051630128a01960e21b81526001600160a01b03871692506304a280649150611fb29085906004016125a3565b60206040518083038186803b158015611fca57600080fd5b505afa158015611fde573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120029190612412565b826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561203b57600080fd5b505afa15801561204f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120739190612412565b146120905760405162461bcd60e51b815260040161045d906126dd565b9392505050565b600080600080886001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156120d657600080fd5b505afa1580156120ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061210e9190612412565b905080881115612129576000806000935093509350506109f2565b878103878111156121375750865b6121be8a8a8c6001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561217557600080fd5b505afa158015612189573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121ad9190612412565b846121b98f8d8d611561565b6121d1565b9450945094505050955095509592505050565b6000806000808451905060005b8681116123c257604051634f0f4aa960e01b8152898201906000906001600160a01b038d1690634f0f4aa9906122189085906004016127c1565b60206040518083038186803b15801561223057600080fd5b505afa158015612244573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061226891906123d6565b905089816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156122a457600080fd5b505afa1580156122b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122dc9190612412565b146122e85750506123ba565b60005b848110156123b657816001600160a01b0316639168ae728a838151811061230e57fe5b60200260200101516040518263ffffffff1660e01b815260040161233291906125a3565b60206040518083038186803b15801561234a57600080fd5b505afa15801561235e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061238291906123f2565b156123ae576001838a838151811061239657fe5b602002602001015197509750975050505050506109f2565b6001016122eb565b5050505b6001016121de565b506000998a99508998509650505050505050565b6000602082840312156123e7578081fd5b81516120908161281b565b600060208284031215612403578081fd5b81518015158114612090578182fd5b600060208284031215612423578081fd5b5051919050565b60006020828403121561243b578081fd5b81356120908161281b565b60008060408385031215612458578081fd5b82356124638161281b565b915060208301356124738161281b565b809150509250929050565b60008060008060808587031215612493578182fd5b843561249e8161281b565b935060208501356124ae8161281b565b925060408501356124be8161281b565b9396929550929360600135925050565b6000806000606084860312156124e2578283fd5b83356124ed8161281b565b95602085013595506040909401359392505050565b60008060008060808587031215612517578384fd5b84356125228161281b565b966020860135965060408601359560600135945092505050565b600080600080600060a08688031215612553578081fd5b853561255e8161281b565b97602087013597506040870135966060810135965060800135945092505050565b60008060408385031215612591578182fd5b8251915060208301516124738161281b565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b818110156125f85783516001600160a01b0316835292840192918401916001016125d3565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156125f857835183529284019291840191600101612620565b9315158452602084019290925260408301526001600160a01b0316606082015260800190565b6001600160a01b03959095168552602085019390935260408401919091526060830152608082015260a00190565b606081016003851061269e57fe5b93815260208101929092526001600160a01b031660409091015290565b60608101600485106126c957fe5b938152602081019290925260409091015290565b6020808252600b908201526a4841535f5354414b45525360a81b604082015260600190565b6020808252600a90820152694e4f5f5354414b45525360b01b604082015260600190565b6020808252600e908201526d1393d517d0531317d4d51052d15160921b604082015260600190565b6020808252600c908201526b24a72b20a624a22fa82922ab60a11b604082015260600190565b6020808252600a90820152694e4f5f4558414d504c4560b01b604082015260600190565b6020808252600f908201526e4245464f52455f444541444c494e4560881b604082015260600190565b90815260200190565b9182526001600160a01b0316602082015260400190565b918252602082015260400190565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b6001600160a01b038116811461283057600080fd5b5056fea2646970667358221220a6ffd737dc8999b977710551ba77dfc9d285843fdba311edac0344d790fd942164736f6c634300060c0033"

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
// Solidity: function getConfig(address rollup) view returns(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsCaller) GetConfig(opts *bind.CallOpts, rollup common.Address) (struct {
	ConfirmPeriodBlocks      *big.Int
	ExtraChallengeTimeBlocks *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "getConfig", rollup)

	outstruct := new(struct {
		ConfirmPeriodBlocks      *big.Int
		ExtraChallengeTimeBlocks *big.Int
		ArbGasSpeedLimitPerBlock *big.Int
		BaseStake                *big.Int
		StakeToken               common.Address
	})

	outstruct.ConfirmPeriodBlocks = out[0].(*big.Int)
	outstruct.ExtraChallengeTimeBlocks = out[1].(*big.Int)
	outstruct.ArbGasSpeedLimitPerBlock = out[2].(*big.Int)
	outstruct.BaseStake = out[3].(*big.Int)
	outstruct.StakeToken = out[4].(common.Address)

	return *outstruct, err

}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsSession) GetConfig(rollup common.Address) (struct {
	ConfirmPeriodBlocks      *big.Int
	ExtraChallengeTimeBlocks *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}, error) {
	return _ValidatorUtils.Contract.GetConfig(&_ValidatorUtils.CallOpts, rollup)
}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsCallerSession) GetConfig(rollup common.Address) (struct {
	ConfirmPeriodBlocks      *big.Int
	ExtraChallengeTimeBlocks *big.Int
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

// LatestStaked is a free data retrieval call binding the contract method 0x01d9717d.
//
// Solidity: function latestStaked(address rollup, address staker) view returns(uint256, bytes32)
func (_ValidatorUtils *ValidatorUtilsCaller) LatestStaked(opts *bind.CallOpts, rollup common.Address, staker common.Address) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "latestStaked", rollup, staker)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// LatestStaked is a free data retrieval call binding the contract method 0x01d9717d.
//
// Solidity: function latestStaked(address rollup, address staker) view returns(uint256, bytes32)
func (_ValidatorUtils *ValidatorUtilsSession) LatestStaked(rollup common.Address, staker common.Address) (*big.Int, [32]byte, error) {
	return _ValidatorUtils.Contract.LatestStaked(&_ValidatorUtils.CallOpts, rollup, staker)
}

// LatestStaked is a free data retrieval call binding the contract method 0x01d9717d.
//
// Solidity: function latestStaked(address rollup, address staker) view returns(uint256, bytes32)
func (_ValidatorUtils *ValidatorUtilsCallerSession) LatestStaked(rollup common.Address, staker common.Address) (*big.Int, [32]byte, error) {
	return _ValidatorUtils.Contract.LatestStaked(&_ValidatorUtils.CallOpts, rollup, staker)
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
