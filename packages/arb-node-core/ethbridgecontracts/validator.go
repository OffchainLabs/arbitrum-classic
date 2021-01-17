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
const ValidatorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"executeTransactions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"returnOldDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorFuncSigs = map[string]string{
	"ce1d571f": "executeTransaction(bytes,address,uint256)",
	"72f45866": "executeTransactions(bytes[],address[],uint256[])",
	"944f4495": "returnOldDeposits(address,address[])",
}

// ValidatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556105b5806100326000396000f3fe6080604052600436106100345760003560e01c806372f4586614610039578063944f44951461004e578063ce1d571f1461006e575b600080fd5b61004c61004736600461034a565b610081565b005b34801561005a57600080fd5b5061004c610069366004610464565b6101a1565b61004c61007c3660046103e0565b610234565b6000546001600160a01b031633146100b45760405162461bcd60e51b81526004016100ab906104fe565b60405180910390fd5b8460005b818110156101975760008686838181106100ce57fe5b90506020020160208101906100e39190610327565b6001600160a01b03168585848181106100f857fe5b905060200201358a8a8581811061010b57fe5b905060200281019061011d9190610522565b60405161012b9291906104b7565b60006040518083038185875af1925050503d8060008114610168576040519150601f19603f3d011682016040523d82523d6000602084013e61016d565b606091505b505090508061018e5760405162461bcd60e51b81526004016100ab906104db565b506001016100b8565b5050505050505050565b8060005b8181101561022d57846001600160a01b0316637427be518585848181106101c857fe5b90506020020160208101906101dd9190610327565b6040518263ffffffff1660e01b81526004016101f991906104c7565b600060405180830381600087803b15801561021357600080fd5b505af1925050508015610224575060015b506001016101a5565b5050505050565b6000546001600160a01b0316331461025e5760405162461bcd60e51b81526004016100ab906104fe565b6000826001600160a01b031682868660405161027b9291906104b7565b60006040518083038185875af1925050503d80600081146102b8576040519150601f19603f3d011682016040523d82523d6000602084013e6102bd565b606091505b505090508061022d5760405162461bcd60e51b81526004016100ab906104db565b60008083601f8401126102ef578182fd5b50813567ffffffffffffffff811115610306578182fd5b602083019150836020808302850101111561032057600080fd5b9250929050565b600060208284031215610338578081fd5b813561034381610567565b9392505050565b60008060008060008060608789031215610362578182fd5b863567ffffffffffffffff80821115610379578384fd5b6103858a838b016102de565b9098509650602089013591508082111561039d578384fd5b6103a98a838b016102de565b909650945060408901359150808211156103c1578384fd5b506103ce89828a016102de565b979a9699509497509295939492505050565b600080600080606085870312156103f5578384fd5b843567ffffffffffffffff8082111561040c578586fd5b818701915087601f83011261041f578586fd5b81358181111561042d578687fd5b88602082850101111561043e578687fd5b6020928301965094505085013561045481610567565b9396929550929360400135925050565b600080600060408486031215610478578283fd5b833561048381610567565b9250602084013567ffffffffffffffff81111561049e578283fd5b6104aa868287016102de565b9497909650939450505050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b602080825260099082015268151617d1905253115160ba1b604082015260600190565b6020808252600a908201526927a7262cafa7aba722a960b11b604082015260600190565b6000808335601e19843603018112610538578283fd5b83018035915067ffffffffffffffff821115610552578283fd5b60200191503681900382131561032057600080fd5b6001600160a01b038116811461057c57600080fd5b5056fea26469706673582212203ea446aa781834fb1393221aa13089722bd51b6d98962e1fb54cdacf0cde7ea764736f6c634300060c0033"

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
