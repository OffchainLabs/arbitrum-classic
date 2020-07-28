// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbostest

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

// SubredditPointsV0ABI is the input ABI used to generate the binding from.
const SubredditPointsV0ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"subreddit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SubredditPointsV0FuncSigs maps the 4-byte function signature to its string representation.
var SubredditPointsV0FuncSigs = map[string]string{
	"bdc330cb": "subreddit()",
}

// SubredditPointsV0Bin is the compiled bytecode used for deploying new contracts.
var SubredditPointsV0Bin = "0x608060405234801561001057600080fd5b50604080518082019091526007808252667465737431323360c81b602090920191825261003f91600091610045565b506100e0565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008657805160ff19168380011785556100b3565b828001600101855582156100b3579182015b828111156100b3578251825591602001919060010190610098565b506100bf9291506100c3565b5090565b6100dd91905b808211156100bf57600081556001016100c9565b90565b610178806100ef6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063bdc330cb14610030575b600080fd5b6100386100ad565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561007257818101518382015260200161005a565b50505050905090810190601f16801561009f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156101395780601f1061010e57610100808354040283529160200191610139565b820191906000526020600020905b81548152906001019060200180831161011c57829003601f168201915b505050505090509056fea265627a7a7231582037496c97748b186e37550b8d1edcf6c9a1cd548ed74ca0f9f62b2148e963d84f64736f6c63430005110032"

// DeploySubredditPointsV0 deploys a new Ethereum contract, binding an instance of SubredditPointsV0 to it.
func DeploySubredditPointsV0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SubredditPointsV0, error) {
	parsed, err := abi.JSON(strings.NewReader(SubredditPointsV0ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SubredditPointsV0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubredditPointsV0{SubredditPointsV0Caller: SubredditPointsV0Caller{contract: contract}, SubredditPointsV0Transactor: SubredditPointsV0Transactor{contract: contract}, SubredditPointsV0Filterer: SubredditPointsV0Filterer{contract: contract}}, nil
}

// SubredditPointsV0 is an auto generated Go binding around an Ethereum contract.
type SubredditPointsV0 struct {
	SubredditPointsV0Caller     // Read-only binding to the contract
	SubredditPointsV0Transactor // Write-only binding to the contract
	SubredditPointsV0Filterer   // Log filterer for contract events
}

// SubredditPointsV0Caller is an auto generated read-only Go binding around an Ethereum contract.
type SubredditPointsV0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubredditPointsV0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SubredditPointsV0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubredditPointsV0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubredditPointsV0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubredditPointsV0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubredditPointsV0Session struct {
	Contract     *SubredditPointsV0 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SubredditPointsV0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubredditPointsV0CallerSession struct {
	Contract *SubredditPointsV0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SubredditPointsV0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubredditPointsV0TransactorSession struct {
	Contract     *SubredditPointsV0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SubredditPointsV0Raw is an auto generated low-level Go binding around an Ethereum contract.
type SubredditPointsV0Raw struct {
	Contract *SubredditPointsV0 // Generic contract binding to access the raw methods on
}

// SubredditPointsV0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubredditPointsV0CallerRaw struct {
	Contract *SubredditPointsV0Caller // Generic read-only contract binding to access the raw methods on
}

// SubredditPointsV0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubredditPointsV0TransactorRaw struct {
	Contract *SubredditPointsV0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSubredditPointsV0 creates a new instance of SubredditPointsV0, bound to a specific deployed contract.
func NewSubredditPointsV0(address common.Address, backend bind.ContractBackend) (*SubredditPointsV0, error) {
	contract, err := bindSubredditPointsV0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubredditPointsV0{SubredditPointsV0Caller: SubredditPointsV0Caller{contract: contract}, SubredditPointsV0Transactor: SubredditPointsV0Transactor{contract: contract}, SubredditPointsV0Filterer: SubredditPointsV0Filterer{contract: contract}}, nil
}

// NewSubredditPointsV0Caller creates a new read-only instance of SubredditPointsV0, bound to a specific deployed contract.
func NewSubredditPointsV0Caller(address common.Address, caller bind.ContractCaller) (*SubredditPointsV0Caller, error) {
	contract, err := bindSubredditPointsV0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubredditPointsV0Caller{contract: contract}, nil
}

// NewSubredditPointsV0Transactor creates a new write-only instance of SubredditPointsV0, bound to a specific deployed contract.
func NewSubredditPointsV0Transactor(address common.Address, transactor bind.ContractTransactor) (*SubredditPointsV0Transactor, error) {
	contract, err := bindSubredditPointsV0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubredditPointsV0Transactor{contract: contract}, nil
}

// NewSubredditPointsV0Filterer creates a new log filterer instance of SubredditPointsV0, bound to a specific deployed contract.
func NewSubredditPointsV0Filterer(address common.Address, filterer bind.ContractFilterer) (*SubredditPointsV0Filterer, error) {
	contract, err := bindSubredditPointsV0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubredditPointsV0Filterer{contract: contract}, nil
}

// bindSubredditPointsV0 binds a generic wrapper to an already deployed contract.
func bindSubredditPointsV0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubredditPointsV0ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubredditPointsV0 *SubredditPointsV0Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubredditPointsV0.Contract.SubredditPointsV0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubredditPointsV0 *SubredditPointsV0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubredditPointsV0.Contract.SubredditPointsV0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubredditPointsV0 *SubredditPointsV0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubredditPointsV0.Contract.SubredditPointsV0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubredditPointsV0 *SubredditPointsV0CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubredditPointsV0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubredditPointsV0 *SubredditPointsV0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubredditPointsV0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubredditPointsV0 *SubredditPointsV0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubredditPointsV0.Contract.contract.Transact(opts, method, params...)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_SubredditPointsV0 *SubredditPointsV0Caller) Subreddit(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SubredditPointsV0.contract.Call(opts, out, "subreddit")
	return *ret0, err
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_SubredditPointsV0 *SubredditPointsV0Session) Subreddit() (string, error) {
	return _SubredditPointsV0.Contract.Subreddit(&_SubredditPointsV0.CallOpts)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_SubredditPointsV0 *SubredditPointsV0CallerSession) Subreddit() (string, error) {
	return _SubredditPointsV0.Contract.Subreddit(&_SubredditPointsV0.CallOpts)
}
