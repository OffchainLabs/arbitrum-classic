// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arboscontracts

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

// NodeInterfaceABI is the input ABI used to generate the binding from.
const NodeInterfaceABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"lookupMessageBatchProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NodeInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var NodeInterfaceFuncSigs = map[string]string{
	"52d388b8": "lookupMessageBatchProof(uint256,uint64)",
}

// NodeInterface is an auto generated Go binding around an Ethereum contract.
type NodeInterface struct {
	NodeInterfaceCaller     // Read-only binding to the contract
	NodeInterfaceTransactor // Write-only binding to the contract
	NodeInterfaceFilterer   // Log filterer for contract events
}

// NodeInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeInterfaceSession struct {
	Contract     *NodeInterface    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeInterfaceCallerSession struct {
	Contract *NodeInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NodeInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeInterfaceTransactorSession struct {
	Contract     *NodeInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NodeInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeInterfaceRaw struct {
	Contract *NodeInterface // Generic contract binding to access the raw methods on
}

// NodeInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeInterfaceCallerRaw struct {
	Contract *NodeInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// NodeInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeInterfaceTransactorRaw struct {
	Contract *NodeInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeInterface creates a new instance of NodeInterface, bound to a specific deployed contract.
func NewNodeInterface(address common.Address, backend bind.ContractBackend) (*NodeInterface, error) {
	contract, err := bindNodeInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeInterface{NodeInterfaceCaller: NodeInterfaceCaller{contract: contract}, NodeInterfaceTransactor: NodeInterfaceTransactor{contract: contract}, NodeInterfaceFilterer: NodeInterfaceFilterer{contract: contract}}, nil
}

// NewNodeInterfaceCaller creates a new read-only instance of NodeInterface, bound to a specific deployed contract.
func NewNodeInterfaceCaller(address common.Address, caller bind.ContractCaller) (*NodeInterfaceCaller, error) {
	contract, err := bindNodeInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeInterfaceCaller{contract: contract}, nil
}

// NewNodeInterfaceTransactor creates a new write-only instance of NodeInterface, bound to a specific deployed contract.
func NewNodeInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeInterfaceTransactor, error) {
	contract, err := bindNodeInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeInterfaceTransactor{contract: contract}, nil
}

// NewNodeInterfaceFilterer creates a new log filterer instance of NodeInterface, bound to a specific deployed contract.
func NewNodeInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeInterfaceFilterer, error) {
	contract, err := bindNodeInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeInterfaceFilterer{contract: contract}, nil
}

// bindNodeInterface binds a generic wrapper to an already deployed contract.
func bindNodeInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeInterface *NodeInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeInterface.Contract.NodeInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeInterface *NodeInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeInterface.Contract.NodeInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeInterface *NodeInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeInterface.Contract.NodeInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeInterface *NodeInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeInterface *NodeInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeInterface *NodeInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeInterface.Contract.contract.Transact(opts, method, params...)
}

// LookupMessageBatchProof is a paid mutator transaction binding the contract method 0x52d388b8.
//
// Solidity: function lookupMessageBatchProof(uint256 batchNum, uint64 index) returns(bytes32[] proof, uint256 path, address l2Sender, address l1Dest, uint256 l2Block, uint256 l1Block, uint256 timestamp, uint256 amount, bytes calldataForL1)
func (_NodeInterface *NodeInterfaceTransactor) LookupMessageBatchProof(opts *bind.TransactOpts, batchNum *big.Int, index uint64) (*types.Transaction, error) {
	return _NodeInterface.contract.Transact(opts, "lookupMessageBatchProof", batchNum, index)
}

// LookupMessageBatchProof is a paid mutator transaction binding the contract method 0x52d388b8.
//
// Solidity: function lookupMessageBatchProof(uint256 batchNum, uint64 index) returns(bytes32[] proof, uint256 path, address l2Sender, address l1Dest, uint256 l2Block, uint256 l1Block, uint256 timestamp, uint256 amount, bytes calldataForL1)
func (_NodeInterface *NodeInterfaceSession) LookupMessageBatchProof(batchNum *big.Int, index uint64) (*types.Transaction, error) {
	return _NodeInterface.Contract.LookupMessageBatchProof(&_NodeInterface.TransactOpts, batchNum, index)
}

// LookupMessageBatchProof is a paid mutator transaction binding the contract method 0x52d388b8.
//
// Solidity: function lookupMessageBatchProof(uint256 batchNum, uint64 index) returns(bytes32[] proof, uint256 path, address l2Sender, address l1Dest, uint256 l2Block, uint256 l1Block, uint256 timestamp, uint256 amount, bytes calldataForL1)
func (_NodeInterface *NodeInterfaceTransactorSession) LookupMessageBatchProof(batchNum *big.Int, index uint64) (*types.Transaction, error) {
	return _NodeInterface.Contract.LookupMessageBatchProof(&_NodeInterface.TransactOpts, batchNum, index)
}
