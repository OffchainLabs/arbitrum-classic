// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rolluptester

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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208665b32688d4c96486cc28d46b2357f88d5452e05a309ca0f9e30def208efbea64736f6c63430005110032"

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

// ChallengeUtilsABI is the input ABI used to generate the binding from.
const ChallengeUtilsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_EXECUTION_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_INBOX_TOP_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVALID_MESSAGES_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALID_CHILD_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeUtilsFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeUtilsFuncSigs = map[string]string{
	"95312727": "INVALID_EXECUTION_TYPE()",
	"a697bcac": "INVALID_INBOX_TOP_TYPE()",
	"d7519b46": "INVALID_MESSAGES_TYPE()",
	"2e179be5": "VALID_CHILD_TYPE()",
}

// ChallengeUtilsBin is the compiled bytecode used for deploying new contracts.
var ChallengeUtilsBin = "0x60c9610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060515760003560e01c80632e179be51460565780639531272714606e578063a697bcac146074578063d7519b4614607a575b600080fd5b605c6080565b60408051918252519081900360200190f35b605c6085565b605c608a565b605c608f565b600381565b600281565b600081565b60018156fea265627a7a72315820ba9f38222d69e81a75675f06da9325b78df15816f56ed13980746298a2cb39cf64736f6c63430005110032"

// DeployChallengeUtils deploys a new Ethereum contract, binding an instance of ChallengeUtils to it.
func DeployChallengeUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeUtils{ChallengeUtilsCaller: ChallengeUtilsCaller{contract: contract}, ChallengeUtilsTransactor: ChallengeUtilsTransactor{contract: contract}, ChallengeUtilsFilterer: ChallengeUtilsFilterer{contract: contract}}, nil
}

// ChallengeUtils is an auto generated Go binding around an Ethereum contract.
type ChallengeUtils struct {
	ChallengeUtilsCaller     // Read-only binding to the contract
	ChallengeUtilsTransactor // Write-only binding to the contract
	ChallengeUtilsFilterer   // Log filterer for contract events
}

// ChallengeUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeUtilsSession struct {
	Contract     *ChallengeUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeUtilsCallerSession struct {
	Contract *ChallengeUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ChallengeUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeUtilsTransactorSession struct {
	Contract     *ChallengeUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChallengeUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeUtilsRaw struct {
	Contract *ChallengeUtils // Generic contract binding to access the raw methods on
}

// ChallengeUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeUtilsCallerRaw struct {
	Contract *ChallengeUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeUtilsTransactorRaw struct {
	Contract *ChallengeUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeUtils creates a new instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtils(address common.Address, backend bind.ContractBackend) (*ChallengeUtils, error) {
	contract, err := bindChallengeUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtils{ChallengeUtilsCaller: ChallengeUtilsCaller{contract: contract}, ChallengeUtilsTransactor: ChallengeUtilsTransactor{contract: contract}, ChallengeUtilsFilterer: ChallengeUtilsFilterer{contract: contract}}, nil
}

// NewChallengeUtilsCaller creates a new read-only instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsCaller(address common.Address, caller bind.ContractCaller) (*ChallengeUtilsCaller, error) {
	contract, err := bindChallengeUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsCaller{contract: contract}, nil
}

// NewChallengeUtilsTransactor creates a new write-only instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeUtilsTransactor, error) {
	contract, err := bindChallengeUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsTransactor{contract: contract}, nil
}

// NewChallengeUtilsFilterer creates a new log filterer instance of ChallengeUtils, bound to a specific deployed contract.
func NewChallengeUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeUtilsFilterer, error) {
	contract, err := bindChallengeUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeUtilsFilterer{contract: contract}, nil
}

// bindChallengeUtils binds a generic wrapper to an already deployed contract.
func bindChallengeUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeUtils *ChallengeUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeUtils.Contract.ChallengeUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeUtils *ChallengeUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.ChallengeUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeUtils *ChallengeUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.ChallengeUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeUtils *ChallengeUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeUtils *ChallengeUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeUtils *ChallengeUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeUtils.Contract.contract.Transact(opts, method, params...)
}

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDEXECUTIONTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_EXECUTION_TYPE")
	return *ret0, err
}

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDEXECUTIONTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDEXECUTIONTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDEXECUTIONTYPE is a free data retrieval call binding the contract method 0x95312727.
//
// Solidity: function INVALID_EXECUTION_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDEXECUTIONTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDEXECUTIONTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDINBOXTOPTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_INBOX_TOP_TYPE")
	return *ret0, err
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDINBOXTOPTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDINBOXTOPTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDINBOXTOPTYPE is a free data retrieval call binding the contract method 0xa697bcac.
//
// Solidity: function INVALID_INBOX_TOP_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDINBOXTOPTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDINBOXTOPTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) INVALIDMESSAGESTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "INVALID_MESSAGES_TYPE")
	return *ret0, err
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) INVALIDMESSAGESTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDMESSAGESTYPE(&_ChallengeUtils.CallOpts)
}

// INVALIDMESSAGESTYPE is a free data retrieval call binding the contract method 0xd7519b46.
//
// Solidity: function INVALID_MESSAGES_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) INVALIDMESSAGESTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.INVALIDMESSAGESTYPE(&_ChallengeUtils.CallOpts)
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCaller) VALIDCHILDTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChallengeUtils.contract.Call(opts, out, "VALID_CHILD_TYPE")
	return *ret0, err
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsSession) VALIDCHILDTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.VALIDCHILDTYPE(&_ChallengeUtils.CallOpts)
}

// VALIDCHILDTYPE is a free data retrieval call binding the contract method 0x2e179be5.
//
// Solidity: function VALID_CHILD_TYPE() view returns(uint256)
func (_ChallengeUtils *ChallengeUtilsCallerSession) VALIDCHILDTYPE() (*big.Int, error) {
	return _ChallengeUtils.Contract.VALIDCHILDTYPE(&_ChallengeUtils.CallOpts)
}

// NodeGraphUtilsABI is the input ABI used to generate the binding from.
const NodeGraphUtilsABI = "[]"

// NodeGraphUtilsBin is the compiled bytecode used for deploying new contracts.
var NodeGraphUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208ec34a33c87a5afa35ca927148ee380555a42a452d3c635ceb2fd708bfb485e964736f6c63430005110032"

// DeployNodeGraphUtils deploys a new Ethereum contract, binding an instance of NodeGraphUtils to it.
func DeployNodeGraphUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeGraphUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeGraphUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeGraphUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeGraphUtils{NodeGraphUtilsCaller: NodeGraphUtilsCaller{contract: contract}, NodeGraphUtilsTransactor: NodeGraphUtilsTransactor{contract: contract}, NodeGraphUtilsFilterer: NodeGraphUtilsFilterer{contract: contract}}, nil
}

// NodeGraphUtils is an auto generated Go binding around an Ethereum contract.
type NodeGraphUtils struct {
	NodeGraphUtilsCaller     // Read-only binding to the contract
	NodeGraphUtilsTransactor // Write-only binding to the contract
	NodeGraphUtilsFilterer   // Log filterer for contract events
}

// NodeGraphUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeGraphUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeGraphUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeGraphUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeGraphUtilsSession struct {
	Contract     *NodeGraphUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeGraphUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeGraphUtilsCallerSession struct {
	Contract *NodeGraphUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NodeGraphUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeGraphUtilsTransactorSession struct {
	Contract     *NodeGraphUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NodeGraphUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeGraphUtilsRaw struct {
	Contract *NodeGraphUtils // Generic contract binding to access the raw methods on
}

// NodeGraphUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeGraphUtilsCallerRaw struct {
	Contract *NodeGraphUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// NodeGraphUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeGraphUtilsTransactorRaw struct {
	Contract *NodeGraphUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeGraphUtils creates a new instance of NodeGraphUtils, bound to a specific deployed contract.
func NewNodeGraphUtils(address common.Address, backend bind.ContractBackend) (*NodeGraphUtils, error) {
	contract, err := bindNodeGraphUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeGraphUtils{NodeGraphUtilsCaller: NodeGraphUtilsCaller{contract: contract}, NodeGraphUtilsTransactor: NodeGraphUtilsTransactor{contract: contract}, NodeGraphUtilsFilterer: NodeGraphUtilsFilterer{contract: contract}}, nil
}

// NewNodeGraphUtilsCaller creates a new read-only instance of NodeGraphUtils, bound to a specific deployed contract.
func NewNodeGraphUtilsCaller(address common.Address, caller bind.ContractCaller) (*NodeGraphUtilsCaller, error) {
	contract, err := bindNodeGraphUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeGraphUtilsCaller{contract: contract}, nil
}

// NewNodeGraphUtilsTransactor creates a new write-only instance of NodeGraphUtils, bound to a specific deployed contract.
func NewNodeGraphUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeGraphUtilsTransactor, error) {
	contract, err := bindNodeGraphUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeGraphUtilsTransactor{contract: contract}, nil
}

// NewNodeGraphUtilsFilterer creates a new log filterer instance of NodeGraphUtils, bound to a specific deployed contract.
func NewNodeGraphUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeGraphUtilsFilterer, error) {
	contract, err := bindNodeGraphUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeGraphUtilsFilterer{contract: contract}, nil
}

// bindNodeGraphUtils binds a generic wrapper to an already deployed contract.
func bindNodeGraphUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeGraphUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeGraphUtils *NodeGraphUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeGraphUtils.Contract.NodeGraphUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeGraphUtils *NodeGraphUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeGraphUtils.Contract.NodeGraphUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeGraphUtils *NodeGraphUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeGraphUtils.Contract.NodeGraphUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeGraphUtils *NodeGraphUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeGraphUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeGraphUtils *NodeGraphUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeGraphUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeGraphUtils *NodeGraphUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeGraphUtils.Contract.contract.Transact(opts, method, params...)
}

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582034d0286ee5fa4342467387fc32b1bf963092a5b3173541fa9f7fa152a5c38feb64736f6c63430005110032"

// DeployProtocol deploys a new Ethereum contract, binding an instance of Protocol to it.
func DeployProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Protocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// Protocol is an auto generated Go binding around an Ethereum contract.
type Protocol struct {
	ProtocolCaller     // Read-only binding to the contract
	ProtocolTransactor // Write-only binding to the contract
	ProtocolFilterer   // Log filterer for contract events
}

// ProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolSession struct {
	Contract     *Protocol         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolCallerSession struct {
	Contract *ProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolTransactorSession struct {
	Contract     *ProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRaw struct {
	Contract *Protocol // Generic contract binding to access the raw methods on
}

// ProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolCallerRaw struct {
	Contract *ProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolTransactorRaw struct {
	Contract *ProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocol creates a new instance of Protocol, bound to a specific deployed contract.
func NewProtocol(address common.Address, backend bind.ContractBackend) (*Protocol, error) {
	contract, err := bindProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// NewProtocolCaller creates a new read-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolCaller(address common.Address, caller bind.ContractCaller) (*ProtocolCaller, error) {
	contract, err := bindProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolCaller{contract: contract}, nil
}

// NewProtocolTransactor creates a new write-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolTransactor, error) {
	contract, err := bindProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTransactor{contract: contract}, nil
}

// NewProtocolFilterer creates a new log filterer instance of Protocol, bound to a specific deployed contract.
func NewProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolFilterer, error) {
	contract, err := bindProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolFilterer{contract: contract}, nil
}

// bindProtocol binds a generic wrapper to an already deployed contract.
func bindProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.ProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transact(opts, method, params...)
}

// RollupTesterABI is the input ABI used to generate the binding from.
const RollupTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"calculateLeafFromPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"prevNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nodeDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"childType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"vmProtoStateHash\",\"type\":\"bytes32\"}],\"name\":\"childNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint128[4]\",\"name\":\"_timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"computePrevLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inboxTop\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxCount\",\"type\":\"uint256\"}],\"name\":\"computeProtoHashBefore\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"confNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initalProtoStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"branches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deadlineTicks\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengeNodeData\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"logsAcc\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"vmProtoStateHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"}],\"name\":\"confirm\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"validNodeHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"lastNode\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gracePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkTimeTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint128[4]\",\"name\":\"_timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"generateInvalidExecutionLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"invalidInboxData\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint128[4]\",\"name\":\"_timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"generateInvalidInboxTopLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gracePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint128[4]\",\"name\":\"_timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"generateInvalidMessagesLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint128[4]\",\"name\":\"_timeBounds\",\"type\":\"uint128[4]\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"generateValidLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"initalProtoStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"branches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deadlineTicks\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengeNodeData\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"logsAcc\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"vmProtoStateHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"processValidNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// RollupTesterFuncSigs maps the 4-byte function signature to its string representation.
var RollupTesterFuncSigs = map[string]string{
	"9584b946": "calculateLeafFromPath(bytes32,bytes32[])",
	"df8f77ed": "childNodeHash(bytes32,uint256,bytes32,uint256,bytes32)",
	"992e2329": "computePrevLeaf(bytes32[9],uint256,uint256,uint32,uint64,uint128[4],uint256,bool,uint64)",
	"8ea546c7": "computeProtoHashBefore(bytes32,bytes32,uint256)",
	"bd912e8e": "confirm(bytes32,bytes32,uint256[],uint256[],bytes32[],bytes32[],bytes32[],uint256[],bytes)",
	"8fa33b1b": "generateInvalidExecutionLeaf(uint256,uint256,uint256,bytes32[9],uint256,uint256,uint32,uint64,uint128[4],uint256,bool,uint64)",
	"d7ae9406": "generateInvalidInboxTopLeaf(uint256[4],bytes32[9],uint256,uint256,uint32,uint64,uint128[4],uint256,bool,uint64)",
	"f9b25c0f": "generateInvalidMessagesLeaf(uint256,uint256,bytes32[9],uint256,uint256,uint32,uint64,uint128[4],uint256,bool,uint64)",
	"02be0bd0": "generateLastMessageHash(bytes,uint256,uint256)",
	"b7bf41cb": "generateValidLeaf(uint256,bytes32[9],uint256,uint256,uint32,uint64,uint128[4],uint256,bool,uint64)",
	"caf32e44": "processValidNode(bytes32,uint256[],uint256[],bytes32[],bytes32[],bytes32[],uint256[],bytes,uint256,uint256)",
}

// RollupTesterBin is the compiled bytecode used for deploying new contracts.
var RollupTesterBin = "0x608060405234801561001057600080fd5b50612b07806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063b7bf41cb11610071578063b7bf41cb146103e0578063bd912e8e146104a4578063caf32e44146108bc578063d7ae940614610c97578063df8f77ed14610d82578063f9b25c0f14610db7576100a9565b806302be0bd0146100ae5780638ea546c7146101705780638fa33b1b146101ab5780639584b94614610279578063992e232914610321575b600080fd5b610157600480360360608110156100c457600080fd5b810190602081018135600160201b8111156100de57600080fd5b8201836020820111156100f057600080fd5b803590602001918460018302840111600160201b8311171561011157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610e7f565b6040805192835260208301919091528051918290030190f35b6101996004803603606081101561018657600080fd5b5080359060208101359060400135610e99565b60408051918252519081900360200190f35b61019960048036036102e08110156101c257600080fd5b60408051610120818101835284359460208101359493810135938101929091610180830191906060840190600990839083908082843760009201919091525050604080516080808201835293968535966020870135969384013563ffffffff169560608501356001600160401b0316955091939092610100820192909181019060049083908390808284376000920191909152509194505081359250506020810135151590604001356001600160401b0316610eae565b6101996004803603604081101561028f57600080fd5b81359190810190604081016020820135600160201b8111156102b057600080fd5b8201836020820111156102c257600080fd5b803590602001918460208302840111600160201b831117156102e357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610fc6945050505050565b610157600480360361028081101561033857600080fd5b81019080806101200190600980602002604051908101604052809291908260096020028082843760009201919091525050604080516080808201835293968535966020870135969384013563ffffffff169560608501356001600160401b0316955091939092610100820192909181019060049083908390808284376000920191909152509194505081359250506020810135151590604001356001600160401b0316610fd9565b61019960048036036102a08110156103f757600080fd5b604080516101208181019092528335939283019291610140830191906020840190600990839083908082843760009201919091525050604080516080808201835293968535966020870135969384013563ffffffff169560608501356001600160401b0316955091939092610100820192909181019060049083908390808284376000920191909152509194505081359250506020810135151590604001356001600160401b03166110f0565b61086160048036036101208110156104bb57600080fd5b813591602081013591810190606081016040820135600160201b8111156104e157600080fd5b8201836020820111156104f357600080fd5b803590602001918460208302840111600160201b8311171561051457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561056357600080fd5b82018360208201111561057557600080fd5b803590602001918460208302840111600160201b8311171561059657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156105e557600080fd5b8201836020820111156105f757600080fd5b803590602001918460208302840111600160201b8311171561061857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561066757600080fd5b82018360208201111561067957600080fd5b803590602001918460208302840111600160201b8311171561069a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106e957600080fd5b8201836020820111156106fb57600080fd5b803590602001918460208302840111600160201b8311171561071c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561076b57600080fd5b82018360208201111561077d57600080fd5b803590602001918460208302840111600160201b8311171561079e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107ed57600080fd5b8201836020820111156107ff57600080fd5b803590602001918460018302840111600160201b8311171561082057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611204945050505050565b6040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b838110156108a757818101518382015260200161088f565b50505050905001935050505060405180910390f35b610c7960048036036101408110156108d357600080fd5b81359190810190604081016020820135600160201b8111156108f457600080fd5b82018360208201111561090657600080fd5b803590602001918460208302840111600160201b8311171561092757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561097657600080fd5b82018360208201111561098857600080fd5b803590602001918460208302840111600160201b831117156109a957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156109f857600080fd5b820183602082011115610a0a57600080fd5b803590602001918460208302840111600160201b83111715610a2b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a7a57600080fd5b820183602082011115610a8c57600080fd5b803590602001918460208302840111600160201b83111715610aad57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610afc57600080fd5b820183602082011115610b0e57600080fd5b803590602001918460208302840111600160201b83111715610b2f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b7e57600080fd5b820183602082011115610b9057600080fd5b803590602001918460208302840111600160201b83111715610bb157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610c0057600080fd5b820183602082011115610c1257600080fd5b803590602001918460018302840111600160201b83111715610c3357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561125d565b60408051938452602084019290925282820152519081900360600190f35b6101996004803603610300811015610cae57600080fd5b810190808060800190600480602002604051908101604052809291908260046020028082843760009201919091525050604080516101208181019092529295949381810193925090600990839083908082843760009201919091525050604080516080808201835293968535966020870135969384013563ffffffff169560608501356001600160401b0316955091939092610100820192909181019060049083908390808284376000920191909152509194505081359250506020810135151590604001356001600160401b03166112bc565b610199600480360360a0811015610d9857600080fd5b50803590602081013590604081013590606081013590608001356113c0565b61019960048036036102c0811015610dce57600080fd5b60408051610120818101835284359460208101359481019390926101608401929091840190600990839083908082843760009201919091525050604080516080808201835293968535966020870135969384013563ffffffff169560608501356001600160401b0316955091939092610100820192909181019060049083908390808284376000920191909152509194505081359250506020810135151590604001356001600160401b03166113d9565b600080610e8d8585856114ef565b91509150935093915050565b6000610ea68484846115b2565b949350505050565b6000610eb861299e565b60408051610220810182528c5181526020808e0151908201528082018c9052908c01516060820152608081018a905260a081018c6003602002015181526020018963ffffffff168152602001886001600160401b031681526020018781526020018681526020018c600460098110610f2c57fe5b602002015181526020018c600560098110610f4357fe5b602002015181526020018c600660098110610f5a57fe5b602002015181526020018515158152602001846001600160401b031681526020018c600760098110610f8857fe5b602002015181526020018c600860098110610f9f57fe5b602002015190529050610fb4818f8f8f6115e9565b9e9d5050505050505050505050505050565b6000610fd28383611614565b9392505050565b600080610fe461299e565b60408051610220810182528d5181526020808f0151908201528082018d9052908d01516060820152608081018b905260a081018d6003602002015181526020018a63ffffffff168152602001896001600160401b031681526020018881526020018781526020018d60046009811061105857fe5b602002015181526020018d60056009811061106f57fe5b602002015181526020018d60066009811061108657fe5b602002015181526020018615158152602001856001600160401b031681526020018d6007600981106110b457fe5b602002015181526020018d6008600981106110cb57fe5b6020020151905290506110dd81611624565b9250925050995099975050505050505050565b60006110fa61299e565b60408051610220810182528c5181526020808e0151908201528082018c9052908c01516060820152608081018a905260a081018c6003602002015181526020018963ffffffff168152602001886001600160401b031681526020018781526020018681526020018c60046009811061116e57fe5b602002015181526020018c60056009811061118557fe5b602002015181526020018c60066009811061119c57fe5b602002015181526020018515158152602001846001600160401b031681526020018c6007600981106111ca57fe5b602002015181526020018c6008600981106111e157fe5b6020020151905290506111f4818d611672565b9c9b505050505050505050505050565b6060600061124b6040518061010001604052808c81526020018b81526020018a8152602001898152602001888152602001878152602001868152602001858152508c61168c565b91509150995099975050505050505050565b60008060006112a66040518061010001604052808f81526020018e81526020018d81526020018c81526020018b81526020018a8152602001898152602001888152508686611749565b9250925092509a509a509a975050505050505050565b60006112c661299e565b60408051610220810182528c5181526020808e0151908201528082018c9052908c01516060820152608081018a905260a081018c6003602002015181526020018963ffffffff168152602001886001600160401b031681526020018781526020018681526020018c60046009811061133a57fe5b602002015181526020018c60056009811061135157fe5b602002015181526020018c60066009811061136857fe5b602002015181526020018515158152602001846001600160401b031681526020018c60076009811061139657fe5b602002015181526020018c6008600981106113ad57fe5b6020020151905290506111f4818d6117d3565b60006113cf8686868686611812565b9695505050505050565b60006113e361299e565b60408051610220810182528c5181526020808e0151908201528082018c9052908c01516060820152608081018a905260a081018c6003602002015181526020018963ffffffff168152602001886001600160401b031681526020018781526020018681526020018c60046009811061145757fe5b602002015181526020018c60056009811061146e57fe5b602002015181526020018c60066009811061148557fe5b602002015181526020018515158152602001846001600160401b031681526020018c6007600981106114b357fe5b602002015181526020018c6008600981106114ca57fe5b6020020151905290506114de818e8e61187a565b9d9c50505050505050505050505050565b60008080806114fc612a2b565b8660005b878110156115a2576115128a83611899565b919650909350915084611565576040805162461bcd60e51b8152602060048201526016602482015275496e76616c6964206f7574707574206d65737361676560501b604482015290519081900360640190fd5b8361156f846119d7565b60408051602080820194909452808201929092528051808303820181526060909201905280519101209350600101611500565b5091989197509095505050505050565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b60008060006115f787611624565b91509150611609878386848a8a611ad2565b979650505050505050565b6000610fd2838360008551611b62565b60008060006116408460000151856020015186604001516115b2565b90506000611667856060015186608001518760a001518860c0015163ffffffff1686611812565b935090915050915091565b60008061167e84611624565b509050610ea6848285611bca565b6060600061169984611c12565b6020808501515160c086015151604080518281528285028101909401905290918180156116d0578160200160208202803883390190505b5093506116db612a5f565b86516116e79087611c69565b90506000805b8481101561173857611700898483611ca3565b909350915081156117305782608001518760018560000151038151811061172357fe5b6020026020010181815250505b6001016116ed565b505060800151925050509250929050565b60008060008060006117778860e00151878a60c001518a8151811061176a57fe5b60200260200101516114ef565b91509150600061179e838a608001518a8151811061179157fe5b6020026020010151611d72565b905060008960a0015189815181106117b257fe5b60200260200101519050828282965096509650505050505b93509350939050565b60008060006117e185611624565b6060860151865160208801516040890151949650929450611809938993879392918790611d9e565b95945050505050565b6040805160208082018490528183018790526060820186905260808083018690528351808403909101815260a08301845280519082012060c0830189905260e08084019190915283518084039091018152610100909201909252805191012095945050505050565b600080600061188886611624565b915091506113cf8683868489611df5565b6000806118a4612a2b565b845184106118c4576000846118b96000611e5a565b9250925092506119d0565b60008085905060008782815181106118d857fe5b016020015160019092019160f81c905060006118f2612a8d565b60ff8316611926576119048a85611edf565b91965094509150848461191684611e5a565b97509750975050505050506119d0565b60ff83166001141561194e5761193c8a85611f32565b91965094509050848461191683612092565b60ff831660021415611964576119168a856120f9565b600360ff84161080159061197b5750600c60ff8416105b156119b65760021983016060611992828d8861219e565b9198509650905086866119a483612258565b995099509950505050505050506119d0565b6000806119c36000611e5a565b9199509750955050505050505b9250925092565b6000600360090160ff16826060015160ff1610611a2f576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16611a4d578151611a4690612347565b9050611acd565b606082015160ff1660011415611a80576020808301518051604082015160608301519290930151611a469391929061236b565b606082015160ff1660021415611a9957611a4682612413565b600360ff16826060015160ff1610158015611abd57506060820151600c60ff909116105b15611acb57611a468261247f565bfe5b919050565b600080611aee88600001518961010001518a610160015161249d565b90506000611b208961018001518a6101a001518b6101c001516000801b8d6101e001516000801b8f6102000151612517565b90506000611b338a60e00151848461257a565b9050611b548989611b4684898b01611d72565b611b4e6125c2565b8b611812565b9a9950505050505050505050565b600084835b83811015611bc05781868281518110611b7c57fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091508080600101915050611b67565b5095945050505050565b6000610ea68383611be5876101e00151886102000151611d72565b611bed6125c7565b611c0d8961018001518a61014001518b61012001518c60400151016115b2565b611812565b60208101515160c08201515160a0830151518114611c2f57600080fd5b8083608001515114611c4057600080fd5b8183604001515114611c5157600080fd5b80820383606001515114611c6457600080fd5b505050565b611c71612a5f565b6040518060a0016040528060008152602001600081526020016000815260200184815260200183815250905092915050565b611cab612a5f565b60008085602001518481518110611cbe57fe5b60200260200101519050600060038214905060008115611d0657611ceb8888600001518960400151611749565b60608a01526040890191909152875160010188529050611d35565b8760600151876020015181518110611d1a57fe5b60200260200101519050866020018051809190600101815250505b611d61876080015189604001518881518110611d4d57fe5b602002602001015183868b60600151611812565b608088015250949694955050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080611dbe896101400151878b61012001518c604001510188036115b2565b9050611de88888611dda84611dd360016125cc565b8801611d72565b611de26125d3565b88611812565b9998505050505050505050565b600080611e1f8760200151886101400151611e0e6125d8565b8a61016001518b61012001516125f9565b90506000611e3882611e3160016125cc565b8601611d72565b9050611e4e878783611e4861263f565b89611812565b98975050505050505050565b611e62612a2b565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611ec7565b611eb4612a2b565b815260200190600190039081611eac5790505b50815260006020820152600160409091015292915050565b6000806000808551905084811080611ef957506020858203105b15611f0e5750600092508391508290506119d0565b600160208601611f24888863ffffffff61264416565b935093509350509250925092565b600080611f3d612a8d565b60008490506000868281518110611f5057fe5b602001015160f81c60f81b60f81c905081806001019250506000878381518110611f7657fe5b016020015160019384019360f89190911c915060009060ff84161415611ffc576000611fa0612a2b565b611faa8b87611899565b909750909250905081611fee575050604080516080810182526000808252602082018190529181018290526060810182905290975088965094506119d09350505050565b611ff7816119d7565b925050505b600061200e8a8663ffffffff61264416565b90506020850194508360ff166001141561205a576040805160808101825260ff909416845260208401919091526001908301819052606083019190915295509193509091506119d09050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b61209a612a2b565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906120e1565b6120ce612a2b565b8152602001906001900390816120c65790505b50815260016020820181905260409091015292915050565b600080612104612a2b565b61210c612a2b565b855160009081908781108061212357506040888203105b1561213b5760008885965096509650505050506119d0565b600061214d8a8a63ffffffff61264416565b905060208901985061215f8a8a611edf565b909a5094509250821561218a576121768185612660565b6001985089975095506119d0945050505050565b6000898697509750975050505050506119d0565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156121e957816020015b6121d6612a2b565b8152602001906001900390816121ce5790505b50905060005b8960ff168160ff161015612246576122078985611899565b8451859060ff861690811061221857fe5b6020908102919091010152945092508261223e57506000955086945092506117ca915050565b6001016121ef565b50600199929850965090945050505050565b612260612a2b565b61226a82516126e4565b6122bb576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156122f2578381815181106122d557fe5b6020026020010151608001518201915080806001019150506122c0565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b60408051602080820193909352815180820384018152908201909152805191012090565b600083156123c5575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120610ea6565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214612468576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161247991906126eb565b92915050565b6000612489612a2b565b61249283612725565b9050610fd281612413565b81516020808401516040808601516060968701518251808601999099526fffffffffffffffffffffffffffffffff19608096871b81168a85015293861b841660508a015290851b83169688019690965294831b166070860152818501929092528251808503909101815260a0909301909152815191012090565b6040805160208082019990995296151560f81b8782015260c09590951b6001600160c01b031916604187015260498601939093526069850191909152608984015260a9808401919091528151808403909101815260c99092019052805191012090565b6040805160c09490941b6001600160c01b0319166020808601919091526028850193909352604880850192909252805180850390920182526068909301909252815191012090565b600290565b600390565b6103e80290565b600090565b604080516000808252602082019092526125f381600161279b565b91505090565b60408051602080820197909752808201959095526060850193909352608084019190915260a0808401919091528151808403909101815260c09092019052805191012090565b600190565b6000816020018351101561265757600080fd5b50016020015190565b612668612a2b565b6040805160a08101825284815281516080810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916126cd565b6126ba612a2b565b8152602001906001900390816126b25790505b508152600260208201526040019290925250919050565b6008101590565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b61272d612a2b565b612736826127ba565b61277c576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b606061278b83604001516127c9565b9050610fd28184608001516128a1565b60006127a5612a2b565b6127af84846128a1565b9050610ea681612413565b600061247982606001516128c0565b6060600882511115612819576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015612846578160200160208202803883390190505b50805190915060005b8181101561289857600061287586838151811061286857fe5b60200260200101516119d7565b90508084838151811061288457fe5b60209081029190910101525060010161284f565b50909392505050565b6128a9612a2b565b60006128b4846128de565b9050610ea68184612660565b6000600c60ff8316108015612479575050600360ff91909116101590565b600060088251111561292e576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561297257818101518382015260200161295a565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915261010081016129ec612ab4565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c0820181905260e082018190526101009091015290565b6040518060a0016040528060008152602001612a45612a8d565b815260606020820181905260006040830181905291015290565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040518060800160405280600490602082028038833950919291505056fea265627a7a72315820fac37ead54e0087a97ef9393272118d8cf9e23a0b8587937a11396b81bffaeb364736f6c63430005110032"

// DeployRollupTester deploys a new Ethereum contract, binding an instance of RollupTester to it.
func DeployRollupTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupTester, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupTester{RollupTesterCaller: RollupTesterCaller{contract: contract}, RollupTesterTransactor: RollupTesterTransactor{contract: contract}, RollupTesterFilterer: RollupTesterFilterer{contract: contract}}, nil
}

// RollupTester is an auto generated Go binding around an Ethereum contract.
type RollupTester struct {
	RollupTesterCaller     // Read-only binding to the contract
	RollupTesterTransactor // Write-only binding to the contract
	RollupTesterFilterer   // Log filterer for contract events
}

// RollupTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupTesterSession struct {
	Contract     *RollupTester     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupTesterCallerSession struct {
	Contract *RollupTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RollupTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTesterTransactorSession struct {
	Contract     *RollupTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RollupTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupTesterRaw struct {
	Contract *RollupTester // Generic contract binding to access the raw methods on
}

// RollupTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupTesterCallerRaw struct {
	Contract *RollupTesterCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTesterTransactorRaw struct {
	Contract *RollupTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupTester creates a new instance of RollupTester, bound to a specific deployed contract.
func NewRollupTester(address common.Address, backend bind.ContractBackend) (*RollupTester, error) {
	contract, err := bindRollupTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupTester{RollupTesterCaller: RollupTesterCaller{contract: contract}, RollupTesterTransactor: RollupTesterTransactor{contract: contract}, RollupTesterFilterer: RollupTesterFilterer{contract: contract}}, nil
}

// NewRollupTesterCaller creates a new read-only instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterCaller(address common.Address, caller bind.ContractCaller) (*RollupTesterCaller, error) {
	contract, err := bindRollupTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTesterCaller{contract: contract}, nil
}

// NewRollupTesterTransactor creates a new write-only instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTesterTransactor, error) {
	contract, err := bindRollupTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTesterTransactor{contract: contract}, nil
}

// NewRollupTesterFilterer creates a new log filterer instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupTesterFilterer, error) {
	contract, err := bindRollupTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupTesterFilterer{contract: contract}, nil
}

// bindRollupTester binds a generic wrapper to an already deployed contract.
func bindRollupTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTester *RollupTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTester.Contract.RollupTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTester *RollupTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTester.Contract.RollupTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTester *RollupTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTester.Contract.RollupTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTester *RollupTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTester *RollupTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTester *RollupTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTester.Contract.contract.Transact(opts, method, params...)
}

// CalculateLeafFromPath is a free data retrieval call binding the contract method 0x9584b946.
//
// Solidity: function calculateLeafFromPath(bytes32 from, bytes32[] proof) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) CalculateLeafFromPath(opts *bind.CallOpts, from [32]byte, proof [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "calculateLeafFromPath", from, proof)
	return *ret0, err
}

// CalculateLeafFromPath is a free data retrieval call binding the contract method 0x9584b946.
//
// Solidity: function calculateLeafFromPath(bytes32 from, bytes32[] proof) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) CalculateLeafFromPath(from [32]byte, proof [][32]byte) ([32]byte, error) {
	return _RollupTester.Contract.CalculateLeafFromPath(&_RollupTester.CallOpts, from, proof)
}

// CalculateLeafFromPath is a free data retrieval call binding the contract method 0x9584b946.
//
// Solidity: function calculateLeafFromPath(bytes32 from, bytes32[] proof) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) CalculateLeafFromPath(from [32]byte, proof [][32]byte) ([32]byte, error) {
	return _RollupTester.Contract.CalculateLeafFromPath(&_RollupTester.CallOpts, from, proof)
}

// ChildNodeHash is a free data retrieval call binding the contract method 0xdf8f77ed.
//
// Solidity: function childNodeHash(bytes32 prevNodeHash, uint256 deadlineTicks, bytes32 nodeDataHash, uint256 childType, bytes32 vmProtoStateHash) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) ChildNodeHash(opts *bind.CallOpts, prevNodeHash [32]byte, deadlineTicks *big.Int, nodeDataHash [32]byte, childType *big.Int, vmProtoStateHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "childNodeHash", prevNodeHash, deadlineTicks, nodeDataHash, childType, vmProtoStateHash)
	return *ret0, err
}

// ChildNodeHash is a free data retrieval call binding the contract method 0xdf8f77ed.
//
// Solidity: function childNodeHash(bytes32 prevNodeHash, uint256 deadlineTicks, bytes32 nodeDataHash, uint256 childType, bytes32 vmProtoStateHash) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) ChildNodeHash(prevNodeHash [32]byte, deadlineTicks *big.Int, nodeDataHash [32]byte, childType *big.Int, vmProtoStateHash [32]byte) ([32]byte, error) {
	return _RollupTester.Contract.ChildNodeHash(&_RollupTester.CallOpts, prevNodeHash, deadlineTicks, nodeDataHash, childType, vmProtoStateHash)
}

// ChildNodeHash is a free data retrieval call binding the contract method 0xdf8f77ed.
//
// Solidity: function childNodeHash(bytes32 prevNodeHash, uint256 deadlineTicks, bytes32 nodeDataHash, uint256 childType, bytes32 vmProtoStateHash) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) ChildNodeHash(prevNodeHash [32]byte, deadlineTicks *big.Int, nodeDataHash [32]byte, childType *big.Int, vmProtoStateHash [32]byte) ([32]byte, error) {
	return _RollupTester.Contract.ChildNodeHash(&_RollupTester.CallOpts, prevNodeHash, deadlineTicks, nodeDataHash, childType, vmProtoStateHash)
}

// ComputePrevLeaf is a free data retrieval call binding the contract method 0x992e2329.
//
// Solidity: function computePrevLeaf(bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32, bytes32)
func (_RollupTester *RollupTesterCaller) ComputePrevLeaf(opts *bind.CallOpts, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RollupTester.contract.Call(opts, out, "computePrevLeaf", _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, *ret1, err
}

// ComputePrevLeaf is a free data retrieval call binding the contract method 0x992e2329.
//
// Solidity: function computePrevLeaf(bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32, bytes32)
func (_RollupTester *RollupTesterSession) ComputePrevLeaf(_fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, [32]byte, error) {
	return _RollupTester.Contract.ComputePrevLeaf(&_RollupTester.CallOpts, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// ComputePrevLeaf is a free data retrieval call binding the contract method 0x992e2329.
//
// Solidity: function computePrevLeaf(bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32, bytes32)
func (_RollupTester *RollupTesterCallerSession) ComputePrevLeaf(_fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, [32]byte, error) {
	return _RollupTester.Contract.ComputePrevLeaf(&_RollupTester.CallOpts, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// ComputeProtoHashBefore is a free data retrieval call binding the contract method 0x8ea546c7.
//
// Solidity: function computeProtoHashBefore(bytes32 machineHash, bytes32 inboxTop, uint256 inboxCount) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) ComputeProtoHashBefore(opts *bind.CallOpts, machineHash [32]byte, inboxTop [32]byte, inboxCount *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "computeProtoHashBefore", machineHash, inboxTop, inboxCount)
	return *ret0, err
}

// ComputeProtoHashBefore is a free data retrieval call binding the contract method 0x8ea546c7.
//
// Solidity: function computeProtoHashBefore(bytes32 machineHash, bytes32 inboxTop, uint256 inboxCount) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) ComputeProtoHashBefore(machineHash [32]byte, inboxTop [32]byte, inboxCount *big.Int) ([32]byte, error) {
	return _RollupTester.Contract.ComputeProtoHashBefore(&_RollupTester.CallOpts, machineHash, inboxTop, inboxCount)
}

// ComputeProtoHashBefore is a free data retrieval call binding the contract method 0x8ea546c7.
//
// Solidity: function computeProtoHashBefore(bytes32 machineHash, bytes32 inboxTop, uint256 inboxCount) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) ComputeProtoHashBefore(machineHash [32]byte, inboxTop [32]byte, inboxCount *big.Int) ([32]byte, error) {
	return _RollupTester.Contract.ComputeProtoHashBefore(&_RollupTester.CallOpts, machineHash, inboxTop, inboxCount)
}

// Confirm is a free data retrieval call binding the contract method 0xbd912e8e.
//
// Solidity: function confirm(bytes32 confNode, bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages) pure returns(bytes32[] validNodeHashes, bytes32 lastNode)
func (_RollupTester *RollupTesterCaller) Confirm(opts *bind.CallOpts, confNode [32]byte, initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte) (struct {
	ValidNodeHashes [][32]byte
	LastNode        [32]byte
}, error) {
	ret := new(struct {
		ValidNodeHashes [][32]byte
		LastNode        [32]byte
	})
	out := ret
	err := _RollupTester.contract.Call(opts, out, "confirm", confNode, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages)
	return *ret, err
}

// Confirm is a free data retrieval call binding the contract method 0xbd912e8e.
//
// Solidity: function confirm(bytes32 confNode, bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages) pure returns(bytes32[] validNodeHashes, bytes32 lastNode)
func (_RollupTester *RollupTesterSession) Confirm(confNode [32]byte, initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte) (struct {
	ValidNodeHashes [][32]byte
	LastNode        [32]byte
}, error) {
	return _RollupTester.Contract.Confirm(&_RollupTester.CallOpts, confNode, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages)
}

// Confirm is a free data retrieval call binding the contract method 0xbd912e8e.
//
// Solidity: function confirm(bytes32 confNode, bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages) pure returns(bytes32[] validNodeHashes, bytes32 lastNode)
func (_RollupTester *RollupTesterCallerSession) Confirm(confNode [32]byte, initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte) (struct {
	ValidNodeHashes [][32]byte
	LastNode        [32]byte
}, error) {
	return _RollupTester.Contract.Confirm(&_RollupTester.CallOpts, confNode, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages)
}

// GenerateInvalidExecutionLeaf is a free data retrieval call binding the contract method 0x8fa33b1b.
//
// Solidity: function generateInvalidExecutionLeaf(uint256 gracePeriodTicks, uint256 checkTimeTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) GenerateInvalidExecutionLeaf(opts *bind.CallOpts, gracePeriodTicks *big.Int, checkTimeTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "generateInvalidExecutionLeaf", gracePeriodTicks, checkTimeTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, err
}

// GenerateInvalidExecutionLeaf is a free data retrieval call binding the contract method 0x8fa33b1b.
//
// Solidity: function generateInvalidExecutionLeaf(uint256 gracePeriodTicks, uint256 checkTimeTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) GenerateInvalidExecutionLeaf(gracePeriodTicks *big.Int, checkTimeTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidExecutionLeaf(&_RollupTester.CallOpts, gracePeriodTicks, checkTimeTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidExecutionLeaf is a free data retrieval call binding the contract method 0x8fa33b1b.
//
// Solidity: function generateInvalidExecutionLeaf(uint256 gracePeriodTicks, uint256 checkTimeTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) GenerateInvalidExecutionLeaf(gracePeriodTicks *big.Int, checkTimeTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidExecutionLeaf(&_RollupTester.CallOpts, gracePeriodTicks, checkTimeTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidInboxTopLeaf is a free data retrieval call binding the contract method 0xd7ae9406.
//
// Solidity: function generateInvalidInboxTopLeaf(uint256[4] invalidInboxData, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) GenerateInvalidInboxTopLeaf(opts *bind.CallOpts, invalidInboxData [4]*big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "generateInvalidInboxTopLeaf", invalidInboxData, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, err
}

// GenerateInvalidInboxTopLeaf is a free data retrieval call binding the contract method 0xd7ae9406.
//
// Solidity: function generateInvalidInboxTopLeaf(uint256[4] invalidInboxData, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) GenerateInvalidInboxTopLeaf(invalidInboxData [4]*big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidInboxTopLeaf(&_RollupTester.CallOpts, invalidInboxData, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidInboxTopLeaf is a free data retrieval call binding the contract method 0xd7ae9406.
//
// Solidity: function generateInvalidInboxTopLeaf(uint256[4] invalidInboxData, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) GenerateInvalidInboxTopLeaf(invalidInboxData [4]*big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidInboxTopLeaf(&_RollupTester.CallOpts, invalidInboxData, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidMessagesLeaf is a free data retrieval call binding the contract method 0xf9b25c0f.
//
// Solidity: function generateInvalidMessagesLeaf(uint256 gracePeriodTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) GenerateInvalidMessagesLeaf(opts *bind.CallOpts, gracePeriodTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "generateInvalidMessagesLeaf", gracePeriodTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, err
}

// GenerateInvalidMessagesLeaf is a free data retrieval call binding the contract method 0xf9b25c0f.
//
// Solidity: function generateInvalidMessagesLeaf(uint256 gracePeriodTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) GenerateInvalidMessagesLeaf(gracePeriodTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidMessagesLeaf(&_RollupTester.CallOpts, gracePeriodTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidMessagesLeaf is a free data retrieval call binding the contract method 0xf9b25c0f.
//
// Solidity: function generateInvalidMessagesLeaf(uint256 gracePeriodTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) GenerateInvalidMessagesLeaf(gracePeriodTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidMessagesLeaf(&_RollupTester.CallOpts, gracePeriodTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) pure returns(bytes32, uint256)
func (_RollupTester *RollupTesterCaller) GenerateLastMessageHash(opts *bind.CallOpts, messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RollupTester.contract.Call(opts, out, "generateLastMessageHash", messages, startOffset, length)
	return *ret0, *ret1, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) pure returns(bytes32, uint256)
func (_RollupTester *RollupTesterSession) GenerateLastMessageHash(messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	return _RollupTester.Contract.GenerateLastMessageHash(&_RollupTester.CallOpts, messages, startOffset, length)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) pure returns(bytes32, uint256)
func (_RollupTester *RollupTesterCallerSession) GenerateLastMessageHash(messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	return _RollupTester.Contract.GenerateLastMessageHash(&_RollupTester.CallOpts, messages, startOffset, length)
}

// GenerateValidLeaf is a free data retrieval call binding the contract method 0xb7bf41cb.
//
// Solidity: function generateValidLeaf(uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) GenerateValidLeaf(opts *bind.CallOpts, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "generateValidLeaf", deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, err
}

// GenerateValidLeaf is a free data retrieval call binding the contract method 0xb7bf41cb.
//
// Solidity: function generateValidLeaf(uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) GenerateValidLeaf(deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateValidLeaf(&_RollupTester.CallOpts, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateValidLeaf is a free data retrieval call binding the contract method 0xb7bf41cb.
//
// Solidity: function generateValidLeaf(uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint128[4] _timeBounds, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) GenerateValidLeaf(deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _timeBounds [4]*big.Int, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateValidLeaf(&_RollupTester.CallOpts, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _timeBounds, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// ProcessValidNode is a free data retrieval call binding the contract method 0xcaf32e44.
//
// Solidity: function processValidNode(bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, uint256 validNum, uint256 startOffset) pure returns(uint256, bytes32, bytes32)
func (_RollupTester *RollupTesterCaller) ProcessValidNode(opts *bind.CallOpts, initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, validNum *big.Int, startOffset *big.Int) (*big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _RollupTester.contract.Call(opts, out, "processValidNode", initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, validNum, startOffset)
	return *ret0, *ret1, *ret2, err
}

// ProcessValidNode is a free data retrieval call binding the contract method 0xcaf32e44.
//
// Solidity: function processValidNode(bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, uint256 validNum, uint256 startOffset) pure returns(uint256, bytes32, bytes32)
func (_RollupTester *RollupTesterSession) ProcessValidNode(initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, validNum *big.Int, startOffset *big.Int) (*big.Int, [32]byte, [32]byte, error) {
	return _RollupTester.Contract.ProcessValidNode(&_RollupTester.CallOpts, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, validNum, startOffset)
}

// ProcessValidNode is a free data retrieval call binding the contract method 0xcaf32e44.
//
// Solidity: function processValidNode(bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, uint256 validNum, uint256 startOffset) pure returns(uint256, bytes32, bytes32)
func (_RollupTester *RollupTesterCallerSession) ProcessValidNode(initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, validNum *big.Int, startOffset *big.Int) (*big.Int, [32]byte, [32]byte, error) {
	return _RollupTester.Contract.ProcessValidNode(&_RollupTester.CallOpts, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, validNum, startOffset)
}

// RollupTimeABI is the input ABI used to generate the binding from.
const RollupTimeABI = "[]"

// RollupTimeBin is the compiled bytecode used for deploying new contracts.
var RollupTimeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582088854a2b1da179961b6aeb78e3b4d67b1336777fbebcf88159b4f07cf3407e6264736f6c63430005110032"

// DeployRollupTime deploys a new Ethereum contract, binding an instance of RollupTime to it.
func DeployRollupTime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupTime, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupTimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupTime{RollupTimeCaller: RollupTimeCaller{contract: contract}, RollupTimeTransactor: RollupTimeTransactor{contract: contract}, RollupTimeFilterer: RollupTimeFilterer{contract: contract}}, nil
}

// RollupTime is an auto generated Go binding around an Ethereum contract.
type RollupTime struct {
	RollupTimeCaller     // Read-only binding to the contract
	RollupTimeTransactor // Write-only binding to the contract
	RollupTimeFilterer   // Log filterer for contract events
}

// RollupTimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupTimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupTimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupTimeSession struct {
	Contract     *RollupTime       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupTimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupTimeCallerSession struct {
	Contract *RollupTimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupTimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTimeTransactorSession struct {
	Contract     *RollupTimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupTimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupTimeRaw struct {
	Contract *RollupTime // Generic contract binding to access the raw methods on
}

// RollupTimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupTimeCallerRaw struct {
	Contract *RollupTimeCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTimeTransactorRaw struct {
	Contract *RollupTimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupTime creates a new instance of RollupTime, bound to a specific deployed contract.
func NewRollupTime(address common.Address, backend bind.ContractBackend) (*RollupTime, error) {
	contract, err := bindRollupTime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupTime{RollupTimeCaller: RollupTimeCaller{contract: contract}, RollupTimeTransactor: RollupTimeTransactor{contract: contract}, RollupTimeFilterer: RollupTimeFilterer{contract: contract}}, nil
}

// NewRollupTimeCaller creates a new read-only instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeCaller(address common.Address, caller bind.ContractCaller) (*RollupTimeCaller, error) {
	contract, err := bindRollupTime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTimeCaller{contract: contract}, nil
}

// NewRollupTimeTransactor creates a new write-only instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTimeTransactor, error) {
	contract, err := bindRollupTime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTimeTransactor{contract: contract}, nil
}

// NewRollupTimeFilterer creates a new log filterer instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupTimeFilterer, error) {
	contract, err := bindRollupTime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupTimeFilterer{contract: contract}, nil
}

// bindRollupTime binds a generic wrapper to an already deployed contract.
func bindRollupTime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTime *RollupTimeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTime.Contract.RollupTimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTime *RollupTimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTime.Contract.RollupTimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTime *RollupTimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTime.Contract.RollupTimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTime *RollupTimeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTime *RollupTimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTime *RollupTimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTime.Contract.contract.Transact(opts, method, params...)
}

// RollupUtilsABI is the input ABI used to generate the binding from.
const RollupUtilsABI = "[]"

// RollupUtilsBin is the compiled bytecode used for deploying new contracts.
var RollupUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158204f7cf3e9752d6a8f42dde678e9068c97afd0da89b34e9040ed27bf56c760e16f64736f6c63430005110032"

// DeployRollupUtils deploys a new Ethereum contract, binding an instance of RollupUtils to it.
func DeployRollupUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupUtils{RollupUtilsCaller: RollupUtilsCaller{contract: contract}, RollupUtilsTransactor: RollupUtilsTransactor{contract: contract}, RollupUtilsFilterer: RollupUtilsFilterer{contract: contract}}, nil
}

// RollupUtils is an auto generated Go binding around an Ethereum contract.
type RollupUtils struct {
	RollupUtilsCaller     // Read-only binding to the contract
	RollupUtilsTransactor // Write-only binding to the contract
	RollupUtilsFilterer   // Log filterer for contract events
}

// RollupUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupUtilsSession struct {
	Contract     *RollupUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupUtilsCallerSession struct {
	Contract *RollupUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RollupUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupUtilsTransactorSession struct {
	Contract     *RollupUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RollupUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupUtilsRaw struct {
	Contract *RollupUtils // Generic contract binding to access the raw methods on
}

// RollupUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupUtilsCallerRaw struct {
	Contract *RollupUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// RollupUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupUtilsTransactorRaw struct {
	Contract *RollupUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupUtils creates a new instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtils(address common.Address, backend bind.ContractBackend) (*RollupUtils, error) {
	contract, err := bindRollupUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupUtils{RollupUtilsCaller: RollupUtilsCaller{contract: contract}, RollupUtilsTransactor: RollupUtilsTransactor{contract: contract}, RollupUtilsFilterer: RollupUtilsFilterer{contract: contract}}, nil
}

// NewRollupUtilsCaller creates a new read-only instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsCaller(address common.Address, caller bind.ContractCaller) (*RollupUtilsCaller, error) {
	contract, err := bindRollupUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsCaller{contract: contract}, nil
}

// NewRollupUtilsTransactor creates a new write-only instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupUtilsTransactor, error) {
	contract, err := bindRollupUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsTransactor{contract: contract}, nil
}

// NewRollupUtilsFilterer creates a new log filterer instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupUtilsFilterer, error) {
	contract, err := bindRollupUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsFilterer{contract: contract}, nil
}

// bindRollupUtils binds a generic wrapper to an already deployed contract.
func bindRollupUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupUtils *RollupUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupUtils.Contract.RollupUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupUtils *RollupUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupUtils.Contract.RollupUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupUtils *RollupUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupUtils.Contract.RollupUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupUtils *RollupUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupUtils *RollupUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupUtils *RollupUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupUtils.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820227ba37fef003ab44dba94eda774b038aafd5f48da618ac76acada099cedc69d64736f6c63430005110032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// VMABI is the input ABI used to generate the binding from.
const VMABI = "[]"

// VMBin is the compiled bytecode used for deploying new contracts.
var VMBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b46eeb36b8eb669c6dc1f1b9d12780eb06451b7b41894407cb56ec4906d20bde64736f6c63430005110032"

// DeployVM deploys a new Ethereum contract, binding an instance of VM to it.
func DeployVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VM, error) {
	parsed, err := abi.JSON(strings.NewReader(VMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VM{VMCaller: VMCaller{contract: contract}, VMTransactor: VMTransactor{contract: contract}, VMFilterer: VMFilterer{contract: contract}}, nil
}

// VM is an auto generated Go binding around an Ethereum contract.
type VM struct {
	VMCaller     // Read-only binding to the contract
	VMTransactor // Write-only binding to the contract
	VMFilterer   // Log filterer for contract events
}

// VMCaller is an auto generated read-only Go binding around an Ethereum contract.
type VMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VMSession struct {
	Contract     *VM               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VMCallerSession struct {
	Contract *VMCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VMTransactorSession struct {
	Contract     *VMTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMRaw is an auto generated low-level Go binding around an Ethereum contract.
type VMRaw struct {
	Contract *VM // Generic contract binding to access the raw methods on
}

// VMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VMCallerRaw struct {
	Contract *VMCaller // Generic read-only contract binding to access the raw methods on
}

// VMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VMTransactorRaw struct {
	Contract *VMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVM creates a new instance of VM, bound to a specific deployed contract.
func NewVM(address common.Address, backend bind.ContractBackend) (*VM, error) {
	contract, err := bindVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VM{VMCaller: VMCaller{contract: contract}, VMTransactor: VMTransactor{contract: contract}, VMFilterer: VMFilterer{contract: contract}}, nil
}

// NewVMCaller creates a new read-only instance of VM, bound to a specific deployed contract.
func NewVMCaller(address common.Address, caller bind.ContractCaller) (*VMCaller, error) {
	contract, err := bindVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VMCaller{contract: contract}, nil
}

// NewVMTransactor creates a new write-only instance of VM, bound to a specific deployed contract.
func NewVMTransactor(address common.Address, transactor bind.ContractTransactor) (*VMTransactor, error) {
	contract, err := bindVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VMTransactor{contract: contract}, nil
}

// NewVMFilterer creates a new log filterer instance of VM, bound to a specific deployed contract.
func NewVMFilterer(address common.Address, filterer bind.ContractFilterer) (*VMFilterer, error) {
	contract, err := bindVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VMFilterer{contract: contract}, nil
}

// bindVM binds a generic wrapper to an already deployed contract.
func bindVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VM *VMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VM.Contract.VMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VM *VMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VM.Contract.VMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VM *VMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VM.Contract.VMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VM *VMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VM *VMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VM *VMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VM.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200d59d38da03f51c5f0efc425ab0f2cd7d701399e9fbdce2cf7027514a993ea7564736f6c63430005110032"

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
