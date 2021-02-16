// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

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

// ClonesABI is the input ABI used to generate the binding from.
const ClonesABI = "[]"

// ClonesBin is the compiled bytecode used for deploying new contracts.
var ClonesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d30fcdeea53525f7e5c478e6c95608c9a572aacd8e5c00a80157e00e2b5ebdb564736f6c634300060c0033"

// DeployClones deploys a new Ethereum contract, binding an instance of Clones to it.
func DeployClones(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Clones, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClonesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// Clones is an auto generated Go binding around an Ethereum contract.
type Clones struct {
	ClonesCaller     // Read-only binding to the contract
	ClonesTransactor // Write-only binding to the contract
	ClonesFilterer   // Log filterer for contract events
}

// ClonesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClonesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClonesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClonesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClonesSession struct {
	Contract     *Clones           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClonesCallerSession struct {
	Contract *ClonesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClonesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClonesTransactorSession struct {
	Contract     *ClonesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClonesRaw struct {
	Contract *Clones // Generic contract binding to access the raw methods on
}

// ClonesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClonesCallerRaw struct {
	Contract *ClonesCaller // Generic read-only contract binding to access the raw methods on
}

// ClonesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClonesTransactorRaw struct {
	Contract *ClonesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClones creates a new instance of Clones, bound to a specific deployed contract.
func NewClones(address common.Address, backend bind.ContractBackend) (*Clones, error) {
	contract, err := bindClones(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// NewClonesCaller creates a new read-only instance of Clones, bound to a specific deployed contract.
func NewClonesCaller(address common.Address, caller bind.ContractCaller) (*ClonesCaller, error) {
	contract, err := bindClones(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesCaller{contract: contract}, nil
}

// NewClonesTransactor creates a new write-only instance of Clones, bound to a specific deployed contract.
func NewClonesTransactor(address common.Address, transactor bind.ContractTransactor) (*ClonesTransactor, error) {
	contract, err := bindClones(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesTransactor{contract: contract}, nil
}

// NewClonesFilterer creates a new log filterer instance of Clones, bound to a specific deployed contract.
func NewClonesFilterer(address common.Address, filterer bind.ContractFilterer) (*ClonesFilterer, error) {
	contract, err := bindClones(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClonesFilterer{contract: contract}, nil
}

// bindClones binds a generic wrapper to an already deployed contract.
func bindClones(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.ClonesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transact(opts, method, params...)
}

// INodeABI is the input ABI used to generate the binding from.
const INodeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"addStaker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"childCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstChildBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"newChildConfirmDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noChildConfirmedBeforeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"removeStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requirePastChildConfirmDeadline\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requirePastDeadline\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"latestConfirmed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"requireRejectExample\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// ChildCreated is a paid mutator transaction binding the contract method 0xe5269ed7.
//
// Solidity: function childCreated() returns()
func (_INode *INodeTransactor) ChildCreated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "childCreated")
}

// ChildCreated is a paid mutator transaction binding the contract method 0xe5269ed7.
//
// Solidity: function childCreated() returns()
func (_INode *INodeSession) ChildCreated() (*types.Transaction, error) {
	return _INode.Contract.ChildCreated(&_INode.TransactOpts)
}

// ChildCreated is a paid mutator transaction binding the contract method 0xe5269ed7.
//
// Solidity: function childCreated() returns()
func (_INode *INodeTransactorSession) ChildCreated() (*types.Transaction, error) {
	return _INode.Contract.ChildCreated(&_INode.TransactOpts)
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

// NodeFactoryABI is the input ABI used to generate the binding from.
const NodeFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"createNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"templateContract\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeFactoryBin is the compiled bytecode used for deploying new contracts.
var NodeFactoryBin = "0x608060405234801561001057600080fd5b5060405161001d9061005f565b604051809103906000f080158015610039573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b039290921691909117905561006c565b610831806103cb83390190565b6103508061007b6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806372be06d81461003b578063d45ab2b51461005f575b600080fd5b610043610094565b604080516001600160a01b039092168252519081900360200190f35b610043600480360360a081101561007557600080fd5b50803590602081013590604081013590606081013590608001356100a3565b6000546001600160a01b031681565b6000805481906100bb906001600160a01b031661014c565b60408051632901acdd60e21b8152336004820152602481018a905260448101899052606481018890526084810187905260a4810186905290519192506001600160a01b0383169163a406b3749160c48082019260009290919082900301818387803b15801561012957600080fd5b505af115801561013d573d6000803e3d6000fd5b50929998505050505050505050565b6000816001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b15801561018757600080fd5b505afa15801561019b573d6000803e3d6000fd5b505050506040513d60208110156101b157600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b60208201529061025f5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561022457818101518382015260200161020c565b50505050905090810190601f1680156102515780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50610272826001600160a01b0316610278565b92915050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116610315576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b91905056fea264697066735822122005f906eb4afd66af2590eb48035d736817d4df78c4b5bff6db0852a09b4c271d64736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff191660011790556108048061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100fc5760003560e01c80632466696e146101015780632edfb42a146101395780633aa1927414610141578063479c92541461014b5780635b8b2280146101535780636971dfe51461015b5780636f791d2914610178578063701da98e1461019457806383197ef01461019c57806388d221c6146101a45780639168ae72146101ac57806396a9fdc0146101d257806397bdc510146101f8578063a0369c1414610200578063a406b37414610208578063cb23bcb51461024c578063d7ff5e3514610270578063dff6978714610278578063e5269ed714610280578063feb508ab14610288575b600080fd5b6101276004803603602081101561011757600080fd5b50356001600160a01b03166102b4565b60408051918252519081900360200190f35b610127610397565b61014961039d565b005b6101276103e9565b6101276103ef565b6101496004803603602081101561017157600080fd5b50356103f5565b610180610447565b604080519115158252519081900360200190f35b610127610450565b610149610456565b6101496104a6565b610180600480360360208110156101c257600080fd5b50356001600160a01b03166104ef565b610149600480360360208110156101e857600080fd5b50356001600160a01b0316610504565b6101276105d6565b6101276105dc565b610149600480360360c081101561021e57600080fd5b506001600160a01b038135169060208101359060408101359060608101359060808101359060a001356105e2565b6102546106b5565b604080516001600160a01b039092168252519081900360200190f35b6101276106c4565b6101276106ca565b6101496106d0565b6101496004803603604081101561029e57600080fd5b50803590602001356001600160a01b031661072a565b6009546000906001600160a01b03163314610304576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b6001600160a01b03821660009081526008602052604090205460ff1615610363576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b506001600160a01b03166000908152600860205260409020805460ff19166001908117909155600780549091019081905590565b60055481565b6006544310156103e7576040805162461bcd60e51b815260206004820152601060248201526f10d212531117d513d3d7d49150d1539560821b604482015290519081900360640190fd5b565b60045481565b60025481565b6009546001600160a01b03163314610442576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b600655565b60005460ff1690565b60015481565b6009546001600160a01b031633146104a3576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b33ff5b6005544310156103e7576040805162461bcd60e51b815260206004820152600f60248201526e4245464f52455f444541444c494e4560881b604482015290519081900360640190fd5b60086020526000908152604090205460ff1681565b6009546001600160a01b03163314610551576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b6001600160a01b03811660009081526008602052604090205460ff166105ab576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6001600160a01b03166000908152600860205260409020805460ff1916905560078054600019019055565b60035481565b60065481565b6009546001600160a01b03161561062f576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b6001600160a01b038616610678576040805162461bcd60e51b815260206004820152600b60248201526a2927a6262aa82fa0a2222960a91b604482015290519081900360640190fd5b600980546001600160a01b0319166001600160a01b0397909716969096179095556001939093556002919091556003556004556005819055600655565b6009546001600160a01b031681565b600a5481565b60075481565b6009546001600160a01b0316331461071d576040805162461bcd60e51b815260206004820152600b60248201526a524f4c4c55505f4f4e4c5960a81b604482015290519081900360640190fd5b600a546103e75743600a55565b8160045414610770576040805162461bcd60e51b815260206004820152600d60248201526c2120a22fa9aaa1a1a2a9a9a7a960991b604482015290519081900360640190fd5b6001600160a01b03811660009081526008602052604090205460ff166107ca576040805162461bcd60e51b815260206004820152600a6024820152692120a22fa9aa20a5a2a960b11b604482015290519081900360640190fd5b505056fea264697066735822122074cce62f988aad9aa8eaf17300fe0a8e8cae0d4e9c6a93653ba68b4583c9cc2264736f6c634300060c0033"

// DeployNodeFactory deploys a new Ethereum contract, binding an instance of NodeFactory to it.
func DeployNodeFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeFactory{NodeFactoryCaller: NodeFactoryCaller{contract: contract}, NodeFactoryTransactor: NodeFactoryTransactor{contract: contract}, NodeFactoryFilterer: NodeFactoryFilterer{contract: contract}}, nil
}

// NodeFactory is an auto generated Go binding around an Ethereum contract.
type NodeFactory struct {
	NodeFactoryCaller     // Read-only binding to the contract
	NodeFactoryTransactor // Write-only binding to the contract
	NodeFactoryFilterer   // Log filterer for contract events
}

// NodeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeFactorySession struct {
	Contract     *NodeFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeFactoryCallerSession struct {
	Contract *NodeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeFactoryTransactorSession struct {
	Contract     *NodeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeFactoryRaw struct {
	Contract *NodeFactory // Generic contract binding to access the raw methods on
}

// NodeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeFactoryCallerRaw struct {
	Contract *NodeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// NodeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeFactoryTransactorRaw struct {
	Contract *NodeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeFactory creates a new instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactory(address common.Address, backend bind.ContractBackend) (*NodeFactory, error) {
	contract, err := bindNodeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeFactory{NodeFactoryCaller: NodeFactoryCaller{contract: contract}, NodeFactoryTransactor: NodeFactoryTransactor{contract: contract}, NodeFactoryFilterer: NodeFactoryFilterer{contract: contract}}, nil
}

// NewNodeFactoryCaller creates a new read-only instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryCaller(address common.Address, caller bind.ContractCaller) (*NodeFactoryCaller, error) {
	contract, err := bindNodeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryCaller{contract: contract}, nil
}

// NewNodeFactoryTransactor creates a new write-only instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeFactoryTransactor, error) {
	contract, err := bindNodeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryTransactor{contract: contract}, nil
}

// NewNodeFactoryFilterer creates a new log filterer instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFactoryFilterer, error) {
	contract, err := bindNodeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryFilterer{contract: contract}, nil
}

// bindNodeFactory binds a generic wrapper to an already deployed contract.
func bindNodeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeFactory *NodeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeFactory.Contract.NodeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeFactory *NodeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeFactory.Contract.NodeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeFactory *NodeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeFactory.Contract.NodeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeFactory *NodeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeFactory *NodeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeFactory *NodeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeFactory.Contract.contract.Transact(opts, method, params...)
}

// TemplateContract is a free data retrieval call binding the contract method 0x72be06d8.
//
// Solidity: function templateContract() view returns(address)
func (_NodeFactory *NodeFactoryCaller) TemplateContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeFactory.contract.Call(opts, &out, "templateContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TemplateContract is a free data retrieval call binding the contract method 0x72be06d8.
//
// Solidity: function templateContract() view returns(address)
func (_NodeFactory *NodeFactorySession) TemplateContract() (common.Address, error) {
	return _NodeFactory.Contract.TemplateContract(&_NodeFactory.CallOpts)
}

// TemplateContract is a free data retrieval call binding the contract method 0x72be06d8.
//
// Solidity: function templateContract() view returns(address)
func (_NodeFactory *NodeFactoryCallerSession) TemplateContract() (common.Address, error) {
	return _NodeFactory.Contract.TemplateContract(&_NodeFactory.CallOpts)
}

// CreateNode is a paid mutator transaction binding the contract method 0xd45ab2b5.
//
// Solidity: function createNode(bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns(address)
func (_NodeFactory *NodeFactoryTransactor) CreateNode(opts *bind.TransactOpts, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _NodeFactory.contract.Transact(opts, "createNode", _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// CreateNode is a paid mutator transaction binding the contract method 0xd45ab2b5.
//
// Solidity: function createNode(bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns(address)
func (_NodeFactory *NodeFactorySession) CreateNode(_stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _NodeFactory.Contract.CreateNode(&_NodeFactory.TransactOpts, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// CreateNode is a paid mutator transaction binding the contract method 0xd45ab2b5.
//
// Solidity: function createNode(bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns(address)
func (_NodeFactory *NodeFactoryTransactorSession) CreateNode(_stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _NodeFactory.Contract.CreateNode(&_NodeFactory.TransactOpts, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}
