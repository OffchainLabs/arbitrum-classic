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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205130ee8069d4d45b671687a9bd9471ae9b0f29d9a9560a567b2a039c58a80e1a64736f6c63430005110032"

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
var ChallengeUtilsBin = "0x60c9610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060515760003560e01c80632e179be51460565780639531272714606e578063a697bcac146074578063d7519b4614607a575b600080fd5b605c6080565b60408051918252519081900360200190f35b605c6085565b605c608a565b605c608f565b600381565b600281565b600081565b60018156fea265627a7a72315820c280990b95818baaf8ee3caded0e9484dbb4af3ca80c8d1a575334ffdc8d07fa64736f6c63430005110032"

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
var NodeGraphUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209293c110c6efbe3eef20bb1d868e154111e3c06e07ec112880ff77c015aafab764736f6c63430005110032"

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
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820bf32b46652fc871de80b14dd5f2895b7770836c862bad0579325746f07565b8964736f6c63430005110032"

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
const RollupTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"calculateLeafFromPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"prevNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nodeDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"childType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"vmProtoStateHash\",\"type\":\"bytes32\"}],\"name\":\"childNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"computePrevLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inboxTop\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxCount\",\"type\":\"uint256\"}],\"name\":\"computeProtoHashBefore\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"confNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initalProtoStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"branches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deadlineTicks\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengeNodeData\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"logsAcc\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"vmProtoStateHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"}],\"name\":\"confirm\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"validNodeHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"lastNode\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gracePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkTimeTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"generateInvalidExecutionLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"invalidInboxData\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"generateInvalidInboxTopLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gracePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"generateInvalidMessagesLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[9]\",\"name\":\"_fields\",\"type\":\"bytes32[9]\"},{\"internalType\":\"uint256\",\"name\":\"_beforeInboxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevDeadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_importedMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"}],\"name\":\"generateValidLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"initalProtoStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"branches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deadlineTicks\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengeNodeData\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"logsAcc\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"vmProtoStateHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"processValidNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// RollupTesterFuncSigs maps the 4-byte function signature to its string representation.
var RollupTesterFuncSigs = map[string]string{
	"9584b946": "calculateLeafFromPath(bytes32,bytes32[])",
	"df8f77ed": "childNodeHash(bytes32,uint256,bytes32,uint256,bytes32)",
	"4eaecd2b": "computePrevLeaf(bytes32[9],uint256,uint256,uint32,uint64,uint256,bool,uint64)",
	"8ea546c7": "computeProtoHashBefore(bytes32,bytes32,uint256)",
	"bd912e8e": "confirm(bytes32,bytes32,uint256[],uint256[],bytes32[],bytes32[],bytes32[],uint256[],bytes)",
	"9e462421": "generateInvalidExecutionLeaf(uint256,uint256,uint256,bytes32[9],uint256,uint256,uint32,uint64,uint256,bool,uint64)",
	"c7d89635": "generateInvalidInboxTopLeaf(uint256[4],bytes32[9],uint256,uint256,uint32,uint64,uint256,bool,uint64)",
	"ce9ed78c": "generateInvalidMessagesLeaf(uint256,uint256,bytes32[9],uint256,uint256,uint32,uint64,uint256,bool,uint64)",
	"02be0bd0": "generateLastMessageHash(bytes,uint256,uint256)",
	"5a852d5b": "generateValidLeaf(uint256,bytes32[9],uint256,uint256,uint32,uint64,uint256,bool,uint64)",
	"caf32e44": "processValidNode(bytes32,uint256[],uint256[],bytes32[],bytes32[],bytes32[],uint256[],bytes,uint256,uint256)",
}

// RollupTesterBin is the compiled bytecode used for deploying new contracts.
var RollupTesterBin = "0x608060405234801561001057600080fd5b506127f8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639e462421116100715780639e4624211461036e578063bd912e8e14610408578063c7d8963514610820578063caf32e44146108d7578063ce9ed78c14610cb2578063df8f77ed14610d46576100a9565b806302be0bd0146100ae5780634eaecd2b146101705780635a852d5b146101fb5780638ea546c71461029d5780639584b946146102c6575b600080fd5b610157600480360360608110156100c457600080fd5b810190602081018135600160201b8111156100de57600080fd5b8201836020820111156100f057600080fd5b803590602001918460018302840111600160201b8311171561011157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610d7b565b6040805192835260208301919091528051918290030190f35b610157600480360361020081101561018757600080fd5b81019080806101200190600980602002604051908101604052809291908260096020028082843760009201919091525091945050813592505060208101359063ffffffff6040820135169067ffffffffffffffff606082013581169160808101359160a082013515159160c0013516610d95565b61028b600480360361022081101561021257600080fd5b604080516101208181019092528335939283019291610140830191906020840190600990839083908082843760009201919091525091945050813592505060208101359063ffffffff6040820135169067ffffffffffffffff606082013581169160808101359160a082013515159160c0013516610e7b565b60408051918252519081900360200190f35b61028b600480360360608110156102b357600080fd5b5080359060208101359060400135610f4f565b61028b600480360360408110156102dc57600080fd5b81359190810190604081016020820135600160201b8111156102fd57600080fd5b82018360208201111561030f57600080fd5b803590602001918460208302840111600160201b8311171561033057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610f64945050505050565b61028b600480360361026081101561038557600080fd5b60408051610120818101835284359460208101359493810135938101929091610180830191906060840190600990839083908082843760009201919091525091945050813592505060208101359063ffffffff6040820135169067ffffffffffffffff606082013581169160808101359160a082013515159160c0013516610f77565b6107c5600480360361012081101561041f57600080fd5b813591602081013591810190606081016040820135600160201b81111561044557600080fd5b82018360208201111561045757600080fd5b803590602001918460208302840111600160201b8311171561047857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104c757600080fd5b8201836020820111156104d957600080fd5b803590602001918460208302840111600160201b831117156104fa57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561054957600080fd5b82018360208201111561055b57600080fd5b803590602001918460208302840111600160201b8311171561057c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156105cb57600080fd5b8201836020820111156105dd57600080fd5b803590602001918460208302840111600160201b831117156105fe57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561064d57600080fd5b82018360208201111561065f57600080fd5b803590602001918460208302840111600160201b8311171561068057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106cf57600080fd5b8201836020820111156106e157600080fd5b803590602001918460208302840111600160201b8311171561070257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561075157600080fd5b82018360208201111561076357600080fd5b803590602001918460018302840111600160201b8311171561078457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061104f945050505050565b6040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b8381101561080b5781810151838201526020016107f3565b50505050905001935050505060405180910390f35b61028b600480360361028081101561083757600080fd5b810190808060800190600480602002604051908101604052809291908260046020028082843760009201919091525050604080516101208181019092529295949381810193925090600990839083908082843760009201919091525091945050813592505060208101359063ffffffff6040820135169067ffffffffffffffff606082013581169160808101359160a082013515159160c00135166110a8565b610c9460048036036101408110156108ee57600080fd5b81359190810190604081016020820135600160201b81111561090f57600080fd5b82018360208201111561092157600080fd5b803590602001918460208302840111600160201b8311171561094257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561099157600080fd5b8201836020820111156109a357600080fd5b803590602001918460208302840111600160201b831117156109c457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a1357600080fd5b820183602082011115610a2557600080fd5b803590602001918460208302840111600160201b83111715610a4657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a9557600080fd5b820183602082011115610aa757600080fd5b803590602001918460208302840111600160201b83111715610ac857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b1757600080fd5b820183602082011115610b2957600080fd5b803590602001918460208302840111600160201b83111715610b4a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b9957600080fd5b820183602082011115610bab57600080fd5b803590602001918460208302840111600160201b83111715610bcc57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610c1b57600080fd5b820183602082011115610c2d57600080fd5b803590602001918460018302840111600160201b83111715610c4e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561116d565b60408051938452602084019290925282820152519081900360600190f35b61028b6004803603610240811015610cc957600080fd5b60408051610120818101835284359460208101359481019390926101608401929091840190600990839083908082843760009201919091525091945050813592505060208101359063ffffffff6040820135169067ffffffffffffffff606082013581169160808101359160a082013515159160c00135166111cc565b61028b600480360360a0811015610d5c57600080fd5b50803590602081013590604081013590606081013590608001356112a2565b600080610d898585856112bb565b91509150935093915050565b600080610da06126be565b60408051610200810182528c5181526020808e0151908201528082018c9052908c015160608083019190915260808083018c9052908d015160a08084019190915263ffffffff8b1660c08085019190915267ffffffffffffffff8b811660e086015261010085018b9052928f0151610120850152908e01516101408401528d015161016083015286151561018083015285166101a08201526101c081018c6007602002015181526020018c600860098110610e5757fe5b602002015190529050610e698161132d565b92509250509850989650505050505050565b6000610e856126be565b5060408051610200810182528a5181526020808c0151908201528082018a9052908a015160608083019190915260808083018a9052908b015160a08084019190915263ffffffff891660c08085019190915267ffffffffffffffff89811660e0808701919091526101008087018b9052948f0151610120870152928e0151610140860152908d015161016085015286151561018085015285166101a08401528b01516101c08301528a01516101e0820152610f40818c61137b565b9b9a5050505050505050505050565b6000610f5c848484611395565b949350505050565b6000610f7083836113cc565b9392505050565b6000610f816126be565b5060408051610200810182528a5181526020808c0151908201528082018a9052908a015160608083019190915260808083018a9052908b015160a08084019190915263ffffffff891660c08085019190915267ffffffffffffffff89811660e0808701919091526101008087018b9052948f0151610120870152928e0151610140860152908d015161016085015286151561018085015285166101a08401528b01516101c08301528a01516101e082015261103e818e8e8e6113dc565b9d9c50505050505050505050505050565b606060006110966040518061010001604052808c81526020018b81526020018a8152602001898152602001888152602001878152602001868152602001858152508c611407565b91509150995099975050505050505050565b60006110b26126be565b5060408051610200810182528a5181526020808c0151908201528082018a9052908a015160608083019190915260808083018a9052908b015160a08084019190915263ffffffff891660c08085019190915267ffffffffffffffff89811660e0808701919091526101008087018b9052948f0151610120870152928e0151610140860152908d015161016085015286151561018085015285166101a08401528b01516101c08301528a01516101e0820152610f40818c6114c5565b60008060006111b66040518061010001604052808f81526020018e81526020018d81526020018c81526020018b81526020018a8152602001898152602001888152508686611504565b9250925092509a509a509a975050505050505050565b60006111d66126be565b5060408051610200810182528a5181526020808c0151908201528082018a9052908a015160608083019190915260808083018a9052908b015160a08084019190915263ffffffff891660c08085019190915267ffffffffffffffff89811660e0808701919091526101008087018b9052948f0151610120870152928e0151610140860152908d015161016085015286151561018085015285166101a08401528b01516101c08301528a01516101e0820152611292818d8d611589565b9c9b505050505050505050505050565b60006112b186868686866115a8565b9695505050505050565b600080806112c7612742565b8560005b8681101561131e576112dd8983611610565b93509150836112eb84611766565b604080516020808201949094528082019290925280518083038201815260609092019052805191012093506001016112cb565b50919791965090945050505050565b6000806000611349846000015185602001518660400151611395565b90506000611370856060015186608001518760a001518860c0015163ffffffff16866115a8565b935090915050915091565b6000806113878461132d565b509050610f5c848285611874565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b6000610f708383600085516118bc565b60008060006113ea8761132d565b915091506113fc878386848a8a611924565b979650505050505050565b60606000611414846119ae565b6020808501515160c0860151516040805182815282850281019094019052909181801561144b578160200160208202803883390190505b509350611456612776565b86516114629087611a05565b90506000805b848110156114b35761147b898483611a3f565b909350915081156114ab5782608001518760018560000151038151811061149e57fe5b6020026020010181815250505b600101611468565b505060800151925050505b9250929050565b60008060006114d38561132d565b60608601518651602088015160408901519496509294506114fb938993879392918790611b0e565b95945050505050565b60008060008060006115328860e00151878a60c001518a8151811061152557fe5b60200260200101516112bb565b915091506000611559838a608001518a8151811061154c57fe5b6020026020010151611b65565b905060008960a00151898151811061156d57fe5b6020908102919091010151929a91995091975095505050505050565b60008060006115978661132d565b915091506112b18683868489611b91565b6040805160208082018490528183018790526060820186905260808083018690528351808403909101815260a08301845280519082012060c0830189905260e08084019190915283518084039091018152610100909201909252805191012095945050505050565b600061161a612742565b83518310611660576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000839050600085828151811061167357fe5b016020015160019092019160f81c90506000816116af576116948784611bf6565b9093509050826116a382611c6a565b945094505050506114be565b60ff8216600114156116c5576116a38784611d1c565b60ff8216600214156116db576116a38784611de0565b600360ff8316108015906116f25750600c60ff8316105b156117265760021982016060611709828a87611e84565b90955090508461171882611f1d565b9650965050505050506114be565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b606081015160009060ff1661178757815161178090612034565b905061186f565b606082015160ff16600114156117a4576117808260200151612058565b606082015160ff16600214156117c557815160808301516117809190612144565b600360ff16826060015160ff16101580156117e957506060820151600c60ff909116105b15611816576117f6612742565b611803836040015161217e565b905061180e81611766565b91505061186f565b606082015160ff166064141561182e5750805161186f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b919050565b6000610f5c838361188f876101c00151886101e00151611b65565b6118976122d6565b6118b78961016001518a61012001518b61010001518c6040015101611395565b6115a8565b600084835b8381101561191a57818682815181106118d657fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120915080806001019150506118c1565b5095945050505050565b60008061193a8860000151896101400151611b65565b9050600061196c8961016001518a61018001518b6101a001516000801b8d6101c001516000801b8f6101e001516122db565b9050600061197f8a60e00151848461233e565b90506119a0898961199284898b01611b65565b61199a612386565b8b6115a8565b9a9950505050505050505050565b60208101515160c08201515160a08301515181146119cb57600080fd5b80836080015151146119dc57600080fd5b81836040015151146119ed57600080fd5b80820383606001515114611a0057600080fd5b505050565b611a0d612776565b6040518060a0016040528060008152602001600081526020016000815260200184815260200183815250905092915050565b611a47612776565b60008085602001518481518110611a5a57fe5b60200260200101519050600060038214905060008115611aa257611a878888600001518960400151611504565b60608a01526040890191909152875160010188529050611ad1565b8760600151876020015181518110611ab657fe5b60200260200101519050866020018051809190600101815250505b611afd876080015189604001518881518110611ae957fe5b602002602001015183868b606001516115a8565b608088015250949694955050505050565b600080611b2e896101200151878b61010001518c60400151018803611395565b9050611b588888611b4a84611b43600161238b565b8801611b65565b611b52612392565b886115a8565b9998505050505050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080611bbb8760200151886101200151611baa612397565b8a61014001518b61010001516123ae565b90506000611bd482611bcd600161238b565b8601611b65565b9050611bea878783611be46123f4565b896115a8565b98975050505050505050565b60008082845110158015611c0e575060208385510310155b611c4b576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301611c5f858563ffffffff6123f916565b915091509250929050565b611c72612742565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191611cc8565b611cb5612742565b815260200190600190039081611cad5790505b50905281526040805160008082526020828101909352919092019190611d04565b611cf1612742565b815260200190600190039081611ce95790505b50815260006020820152600160409091015292915050565b6000611d26612742565b60008390506000858281518110611d3957fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110611d5f57fe5b016020015160019093019260f81c9050611d77612742565b8260ff1660011415611d9357611d8d8885611610565b90945090505b6000611da5898663ffffffff6123f916565b90506020850194508360ff1660011415611dc55784611718848385612415565b84611dd08483612490565b9650965050505050509250929050565b6000611dea612742565b600083855110158015611e01575060408486510310155b611e3d576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b6000611e4f868663ffffffff6123f916565b9050602085019450611e618686611bf6565b9095509150611e6e612742565b611e7882846124eb565b95979596505050505050565b60006060600083905060608660ff16604051908082528060200260200182016040528015611ecc57816020015b611eb9612742565b815260200190600190039081611eb15790505b50905060005b8760ff168160ff161015611f1057611eea8784611610565b8351849060ff8516908110611efb57fe5b60209081029190910101529250600101611ed2565b5090969095509350505050565b611f25612742565b611f2f825161259c565b611f80576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015611fb757838181518110611f9a57fe5b602002602001015160800151820191508080600101915050611f85565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190612011565b611ffe612742565b815260200190600190039081611ff65790505b509052815260208101859052935160030160ff1660408501526060909301525090565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061206957fe5b6040820151516120c35750805160208083015160408051600160f81b8185015260f89490941b6001600160f81b0319166021850152602280850192909252805180850390920182526042909301909252815191012061186f565b600182600001516120eb84604001516000815181106120de57fe5b6020026020010151611766565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b612186612742565b6008825111156121d4576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015612201578160200160208202803883390190505b508051909150600160005b82811015612264576122238682815181106120de57fe5b84828151811061222f57fe5b60200260200101818152505085818151811061224757fe5b60200260200101516080015182019150808060010191505061220c565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156122a9578181015183820152602001612291565b50505050905001925050506040516020818303038152906040528051906020012090506112b181836124eb565b600390565b6040805160208082019990995296151560f81b8782015260c09590951b6001600160c01b031916604187015260498601939093526069850191909152608984015260a9808401919091528151808403909101815260c99092019052805191012090565b6040805160c09490941b6001600160c01b0319166020808601919091526028850193909352604880850192909252805180850390920182526068909301909252815191012090565b600290565b6103e80290565b600090565b60006123a96123a46125a3565b611766565b905090565b60408051602080820197909752808201959095526060850193909352608084019190915260a0808401919091528151808403909101815260c09092019052805191012090565b600190565b6000816020018351101561240c57600080fd5b50016020015190565b61241d612742565b604080516001808252818301909252606091816020015b61243c612742565b815260200190600190039081612434579050509050828160008151811061245f57fe5b60200260200101819052506114fb60405180606001604052808760ff16815260200186815260200183815250612657565b612498612742565b6040805160608101825260ff851681526020808201859052825160008082529181018452610f70938301916124e3565b6124d0612742565b8152602001906001900390816124c85790505b509052612657565b6124f3612742565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191612549565b612536612742565b81526020019060019003908161252e5790505b50905281526040805160008082526020828101909352919092019190612585565b612572612742565b81526020019060019003908161256a5790505b508152600260208201526040019290925250919050565b6008101590565b6125ab612742565b6040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190612604565b6125f1612742565b8152602001906001900390816125e95790505b50905281526040805160008082526020828101909352919092019190612640565b61262d612742565b8152602001906001900390816126255790505b508152600360208201526001604090910152905090565b61265f612742565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906126a6565b612693612742565b81526020019060019003908161268b5790505b50815260016020820181905260409091015292915050565b6040805161020081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081019190915290565b6040518060a001604052806000815260200161275c6127a4565b815260606020820181905260006040830181905291015290565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b604080516060808201835260008083526020830152918101919091529056fea265627a7a723158200af51512f158a764aec2c7422f50b707c9cbf098cebd6d9095a164a779dc809b64736f6c63430005110032"

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

// ComputePrevLeaf is a free data retrieval call binding the contract method 0x4eaecd2b.
//
// Solidity: function computePrevLeaf(bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32, bytes32)
func (_RollupTester *RollupTesterCaller) ComputePrevLeaf(opts *bind.CallOpts, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RollupTester.contract.Call(opts, out, "computePrevLeaf", _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, *ret1, err
}

// ComputePrevLeaf is a free data retrieval call binding the contract method 0x4eaecd2b.
//
// Solidity: function computePrevLeaf(bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32, bytes32)
func (_RollupTester *RollupTesterSession) ComputePrevLeaf(_fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, [32]byte, error) {
	return _RollupTester.Contract.ComputePrevLeaf(&_RollupTester.CallOpts, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// ComputePrevLeaf is a free data retrieval call binding the contract method 0x4eaecd2b.
//
// Solidity: function computePrevLeaf(bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32, bytes32)
func (_RollupTester *RollupTesterCallerSession) ComputePrevLeaf(_fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, [32]byte, error) {
	return _RollupTester.Contract.ComputePrevLeaf(&_RollupTester.CallOpts, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
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

// GenerateInvalidExecutionLeaf is a free data retrieval call binding the contract method 0x9e462421.
//
// Solidity: function generateInvalidExecutionLeaf(uint256 gracePeriodTicks, uint256 checkTimeTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) GenerateInvalidExecutionLeaf(opts *bind.CallOpts, gracePeriodTicks *big.Int, checkTimeTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "generateInvalidExecutionLeaf", gracePeriodTicks, checkTimeTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, err
}

// GenerateInvalidExecutionLeaf is a free data retrieval call binding the contract method 0x9e462421.
//
// Solidity: function generateInvalidExecutionLeaf(uint256 gracePeriodTicks, uint256 checkTimeTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) GenerateInvalidExecutionLeaf(gracePeriodTicks *big.Int, checkTimeTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidExecutionLeaf(&_RollupTester.CallOpts, gracePeriodTicks, checkTimeTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidExecutionLeaf is a free data retrieval call binding the contract method 0x9e462421.
//
// Solidity: function generateInvalidExecutionLeaf(uint256 gracePeriodTicks, uint256 checkTimeTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) GenerateInvalidExecutionLeaf(gracePeriodTicks *big.Int, checkTimeTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidExecutionLeaf(&_RollupTester.CallOpts, gracePeriodTicks, checkTimeTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidInboxTopLeaf is a free data retrieval call binding the contract method 0xc7d89635.
//
// Solidity: function generateInvalidInboxTopLeaf(uint256[4] invalidInboxData, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) GenerateInvalidInboxTopLeaf(opts *bind.CallOpts, invalidInboxData [4]*big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "generateInvalidInboxTopLeaf", invalidInboxData, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, err
}

// GenerateInvalidInboxTopLeaf is a free data retrieval call binding the contract method 0xc7d89635.
//
// Solidity: function generateInvalidInboxTopLeaf(uint256[4] invalidInboxData, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) GenerateInvalidInboxTopLeaf(invalidInboxData [4]*big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidInboxTopLeaf(&_RollupTester.CallOpts, invalidInboxData, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidInboxTopLeaf is a free data retrieval call binding the contract method 0xc7d89635.
//
// Solidity: function generateInvalidInboxTopLeaf(uint256[4] invalidInboxData, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) GenerateInvalidInboxTopLeaf(invalidInboxData [4]*big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidInboxTopLeaf(&_RollupTester.CallOpts, invalidInboxData, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidMessagesLeaf is a free data retrieval call binding the contract method 0xce9ed78c.
//
// Solidity: function generateInvalidMessagesLeaf(uint256 gracePeriodTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) GenerateInvalidMessagesLeaf(opts *bind.CallOpts, gracePeriodTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "generateInvalidMessagesLeaf", gracePeriodTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, err
}

// GenerateInvalidMessagesLeaf is a free data retrieval call binding the contract method 0xce9ed78c.
//
// Solidity: function generateInvalidMessagesLeaf(uint256 gracePeriodTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) GenerateInvalidMessagesLeaf(gracePeriodTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidMessagesLeaf(&_RollupTester.CallOpts, gracePeriodTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateInvalidMessagesLeaf is a free data retrieval call binding the contract method 0xce9ed78c.
//
// Solidity: function generateInvalidMessagesLeaf(uint256 gracePeriodTicks, uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) GenerateInvalidMessagesLeaf(gracePeriodTicks *big.Int, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateInvalidMessagesLeaf(&_RollupTester.CallOpts, gracePeriodTicks, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
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

// GenerateValidLeaf is a free data retrieval call binding the contract method 0x5a852d5b.
//
// Solidity: function generateValidLeaf(uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCaller) GenerateValidLeaf(opts *bind.CallOpts, deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RollupTester.contract.Call(opts, out, "generateValidLeaf", deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
	return *ret0, err
}

// GenerateValidLeaf is a free data retrieval call binding the contract method 0x5a852d5b.
//
// Solidity: function generateValidLeaf(uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterSession) GenerateValidLeaf(deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateValidLeaf(&_RollupTester.CallOpts, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
}

// GenerateValidLeaf is a free data retrieval call binding the contract method 0x5a852d5b.
//
// Solidity: function generateValidLeaf(uint256 deadlineTicks, bytes32[9] _fields, uint256 _beforeInboxCount, uint256 _prevDeadlineTicks, uint32 _prevChildType, uint64 _numSteps, uint256 _importedMessageCount, bool _didInboxInsn, uint64 _numArbGas) pure returns(bytes32)
func (_RollupTester *RollupTesterCallerSession) GenerateValidLeaf(deadlineTicks *big.Int, _fields [9][32]byte, _beforeInboxCount *big.Int, _prevDeadlineTicks *big.Int, _prevChildType uint32, _numSteps uint64, _importedMessageCount *big.Int, _didInboxInsn bool, _numArbGas uint64) ([32]byte, error) {
	return _RollupTester.Contract.GenerateValidLeaf(&_RollupTester.CallOpts, deadlineTicks, _fields, _beforeInboxCount, _prevDeadlineTicks, _prevChildType, _numSteps, _importedMessageCount, _didInboxInsn, _numArbGas)
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
var RollupTimeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820bc5e1c1239ee257a8886402d63fd0e30faaddb124d4cfd39eb5105366f8e112b64736f6c63430005110032"

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
var RollupUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820282bc926ff34a15a2cb97267d1aff8c54c95033849099c209eae2038002e918664736f6c63430005110032"

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
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b6e334180b77cc68114f188c0aa47be28528be9d593d39d3e17df219ffce7d1b64736f6c63430005110032"

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
var VMBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582037c08b19ae0e4b19b1d10ce346658aafd39a52067692d3644af2c73df2e9612c64736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820783df50db562880f31b02767d5965e7d3e1dfa00c7f3fec9ee824ef6915cf4ab64736f6c63430005110032"

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
