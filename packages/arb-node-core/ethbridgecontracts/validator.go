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

// ValidatorFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorFuncSigs = map[string]string{
	"ce1d571f": "executeTransaction(bytes,address,uint256)",
	"72f45866": "executeTransactions(bytes[],address[],uint256[])",
	"944f4495": "returnOldDeposits(address,address[])",
	"81aac2d9": "timeoutChallenges(address[])",
}

// ValidatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317905561066e806100326000396000f3fe60806040526004361061003f5760003560e01c806372f458661461004457806381aac2d914610059578063944f449514610079578063ce1d571f14610099575b600080fd5b6100576100523660046103e6565b6100ac565b005b34801561006557600080fd5b5061005761007436600461047c565b6101c1565b34801561008557600080fd5b50610057610094366004610540565b610248565b6100576100a73660046104bc565b6102db565b6000546001600160a01b031633146100df5760405162461bcd60e51b81526004016100d6906105b7565b60405180910390fd5b8460005b818110156101b75760008686838181106100f957fe5b905060200201602081019061010e91906103c3565b6001600160a01b031685858481811061012357fe5b905060200201358a8a8581811061013657fe5b905060200281019061014891906105db565b604051610156929190610593565b60006040518083038185875af1925050503d8060008114610193576040519150601f19603f3d011682016040523d82523d6000602084013e610198565b606091505b50509050806101ae576040513d806000833e8082fd5b506001016100e3565b5050505050505050565b8060005b81811015610242578383828181106101d957fe5b90506020020160208101906101ee91906103c3565b6001600160a01b03166370dea79a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561022857600080fd5b505af1925050508015610239575060015b506001016101c5565b50505050565b8060005b818110156102d457846001600160a01b0316637427be5185858481811061026f57fe5b905060200201602081019061028491906103c3565b6040518263ffffffff1660e01b81526004016102a091906105a3565b600060405180830381600087803b1580156102ba57600080fd5b505af19250505080156102cb575060015b5060010161024c565b5050505050565b6000546001600160a01b031633146103055760405162461bcd60e51b81526004016100d6906105b7565b6000826001600160a01b0316828686604051610322929190610593565b60006040518083038185875af1925050503d806000811461035f576040519150601f19603f3d011682016040523d82523d6000602084013e610364565b606091505b50509050806102d4576040513d806000833e8082fd5b60008083601f84011261038b578182fd5b50813567ffffffffffffffff8111156103a2578182fd5b60208301915083602080830285010111156103bc57600080fd5b9250929050565b6000602082840312156103d4578081fd5b81356103df81610620565b9392505050565b600080600080600080606087890312156103fe578182fd5b863567ffffffffffffffff80821115610415578384fd5b6104218a838b0161037a565b90985096506020890135915080821115610439578384fd5b6104458a838b0161037a565b9096509450604089013591508082111561045d578384fd5b5061046a89828a0161037a565b979a9699509497509295939492505050565b6000806020838503121561048e578182fd5b823567ffffffffffffffff8111156104a4578283fd5b6104b08582860161037a565b90969095509350505050565b600080600080606085870312156104d1578384fd5b843567ffffffffffffffff808211156104e8578586fd5b818701915087601f8301126104fb578586fd5b813581811115610509578687fd5b88602082850101111561051a578687fd5b6020928301965094505085013561053081610620565b9396929550929360400135925050565b600080600060408486031215610554578283fd5b833561055f81610620565b9250602084013567ffffffffffffffff81111561057a578283fd5b6105868682870161037a565b9497909650939450505050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b6020808252600a908201526927a7262cafa7aba722a960b11b604082015260600190565b6000808335601e198436030181126105f1578283fd5b83018035915067ffffffffffffffff82111561060b578283fd5b6020019150368190038213156103bc57600080fd5b6001600160a01b038116811461063557600080fd5b5056fea26469706673582212206f4668317b2b5c15f61dab6b87191431209e568581edb4f9141e9b7c75be1ef464736f6c634300060c0033"

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
