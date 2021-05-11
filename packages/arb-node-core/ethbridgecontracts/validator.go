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

// ValidatorABI is the input ABI used to generate the binding from.
const ValidatorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"executeTransactions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"returnOldDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallenge[]\",\"name\":\"challenges\",\"type\":\"address[]\"}],\"name\":\"timeoutChallenges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610804806100326000396000f3fe60806040526004361061003f5760003560e01c806372f458661461004457806381aac2d914610059578063944f449514610079578063ce1d571f14610099575b600080fd5b61005761005236600461053b565b6100ac565b005b34801561006557600080fd5b506100576100743660046105d0565b610237565b34801561008557600080fd5b50610057610094366004610692565b610311565b6100576100a736600461060f565b6103f7565b6000546001600160a01b031633146100df5760405162461bcd60e51b81526004016100d69061074e565b60405180910390fd5b8460005b8181101561022d5760008888838181106100f957fe5b905060200281019061010b9190610772565b905011156101615761014586868381811061012257fe5b90506020020160208101906101379190610518565b6001600160a01b03166104ca565b6101615760405162461bcd60e51b81526004016100d690610708565b600086868381811061016f57fe5b90506020020160208101906101849190610518565b6001600160a01b031685858481811061019957fe5b905060200201358a8a858181106101ac57fe5b90506020028101906101be9190610772565b6040516101cc9291906106e4565b60006040518083038185875af1925050503d8060008114610209576040519150601f19603f3d011682016040523d82523d6000602084013e61020e565b606091505b5050905080610224576040513d806000833e8082fd5b506001016100e3565b5050505050505050565b8060005b8181101561030b5783838281811061024f57fe5b90506020020160208101906102649190610518565b6001600160a01b03166370dea79a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561029e57600080fd5b505af19250505080156102af575060015b610303573d8080156102dd576040519150601f19603f3d011682016040523d82523d6000602084013e6102e2565b606091505b5080516103015760405162461bcd60e51b81526004016100d690610731565b505b60010161023b565b50505050565b8060005b818110156103f057846001600160a01b0316637427be5185858481811061033857fe5b905060200201602081019061034d9190610518565b6040518263ffffffff1660e01b815260040161036991906106f4565b600060405180830381600087803b15801561038357600080fd5b505af1925050508015610394575060015b6103e8573d8080156103c2576040519150601f19603f3d011682016040523d82523d6000602084013e6103c7565b606091505b5080516103e65760405162461bcd60e51b81526004016100d690610731565b505b600101610315565b5050505050565b6000546001600160a01b031633146104215760405162461bcd60e51b81526004016100d69061074e565b821561045557610439826001600160a01b03166104ca565b6104555760405162461bcd60e51b81526004016100d690610708565b6000826001600160a01b03168286866040516104729291906106e4565b60006040518083038185875af1925050503d80600081146104af576040519150601f19603f3d011682016040523d82523d6000602084013e6104b4565b606091505b50509050806103f0576040513d806000833e8082fd5b3b151590565b60008083601f8401126104e1578182fd5b5081356001600160401b038111156104f7578182fd5b602083019150836020808302850101111561051157600080fd5b9250929050565b600060208284031215610529578081fd5b8135610534816107b6565b9392505050565b60008060008060008060608789031215610553578182fd5b86356001600160401b0380821115610569578384fd5b6105758a838b016104d0565b9098509650602089013591508082111561058d578384fd5b6105998a838b016104d0565b909650945060408901359150808211156105b1578384fd5b506105be89828a016104d0565b979a9699509497509295939492505050565b600080602083850312156105e2578182fd5b82356001600160401b038111156105f7578283fd5b610603858286016104d0565b90969095509350505050565b60008060008060608587031215610624578384fd5b84356001600160401b038082111561063a578586fd5b818701915087601f83011261064d578586fd5b81358181111561065b578687fd5b88602082850101111561066c578687fd5b60209283019650945050850135610682816107b6565b9396929550929360400135925050565b6000806000604084860312156106a6578283fd5b83356106b1816107b6565b925060208401356001600160401b038111156106cb578283fd5b6106d7868287016104d0565b9497909650939450505050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b6020808252600f908201526e2727afa1a7a222afa0aa2fa0a2222960891b604082015260600190565b60208082526003908201526247415360e81b604082015260600190565b6020808252600a908201526927a7262cafa7aba722a960b11b604082015260600190565b6000808335601e19843603018112610788578283fd5b8301803591506001600160401b038211156107a1578283fd5b60200191503681900382131561051157600080fd5b6001600160a01b03811681146107cb57600080fd5b5056fea2646970667358221220d4441b34053c76e9e4dde6be0c6bdc4cb19f250533981fa0003ca86098f4fa8664736f6c634300060c0033"

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorTransactor) ExecuteTransaction(opts *bind.TransactOpts, data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "executeTransaction", data, destination, amount)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorSession) ExecuteTransaction(data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransaction(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorTransactorSession) ExecuteTransaction(data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransaction(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorTransactor) ExecuteTransactions(opts *bind.TransactOpts, data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "executeTransactions", data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorSession) ExecuteTransactions(data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactions(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorTransactorSession) ExecuteTransactions(data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactions(&_Validator.TransactOpts, data, destination, amount)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorTransactor) ReturnOldDeposits(opts *bind.TransactOpts, rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "returnOldDeposits", rollup, stakers)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorSession) ReturnOldDeposits(rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ReturnOldDeposits(&_Validator.TransactOpts, rollup, stakers)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorTransactorSession) ReturnOldDeposits(rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ReturnOldDeposits(&_Validator.TransactOpts, rollup, stakers)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x81aac2d9.
//
// Solidity: function timeoutChallenges(address[] challenges) returns()
func (_Validator *ValidatorTransactor) TimeoutChallenges(opts *bind.TransactOpts, challenges []common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "timeoutChallenges", challenges)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x81aac2d9.
//
// Solidity: function timeoutChallenges(address[] challenges) returns()
func (_Validator *ValidatorSession) TimeoutChallenges(challenges []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TimeoutChallenges(&_Validator.TransactOpts, challenges)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x81aac2d9.
//
// Solidity: function timeoutChallenges(address[] challenges) returns()
func (_Validator *ValidatorTransactorSession) TimeoutChallenges(challenges []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TimeoutChallenges(&_Validator.TransactOpts, challenges)
}
