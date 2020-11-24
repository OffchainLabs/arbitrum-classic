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

// ChallengeTesterABI is the input ABI used to generate the binding from.
const ChallengeTesterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challengeFactory_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"asserterAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"challengerAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"challengerPeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"challengerDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"startChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeTesterFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeTesterFuncSigs = map[string]string{
	"396f51cf": "resolveChallenge(address,address)",
	"8f43ee32": "startChallenge(address,address,uint128,bytes32,uint256)",
}

// ChallengeTesterBin is the compiled bytecode used for deploying new contracts.
var ChallengeTesterBin = "0x608060405234801561001057600080fd5b506040516101fe3803806101fe8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610199806100656000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063396f51cf1461003b5780638f43ee321461006b575b600080fd5b6100696004803603604081101561005157600080fd5b506001600160a01b03813581169160200135166100b6565b005b610069600480360360a081101561008157600080fd5b506001600160a01b0381358116916020810135909116906001600160801b0360408201351690606081013590608001356100ba565b5050565b600080546040805163432ed0e160e11b81526001600160a01b03898116600483015288811660248301526001600160801b038816604483015260648201879052608482018690529151919092169263865da1c29260a480820193602093909283900390910190829087803b15801561013157600080fd5b505af1158015610145573d6000803e3d6000fd5b505050506040513d602081101561015b57600080fd5b5050505050505056fea265627a7a72315820a01815879b1caabcbb9040d6dcddee2189b506734b4a1d9b66a3bc7d20a025a164736f6c63430005110032"

// DeployChallengeTester deploys a new Ethereum contract, binding an instance of ChallengeTester to it.
func DeployChallengeTester(auth *bind.TransactOpts, backend bind.ContractBackend, challengeFactory_ common.Address) (common.Address, *types.Transaction, *ChallengeTester, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeTesterBin), backend, challengeFactory_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeTester{ChallengeTesterCaller: ChallengeTesterCaller{contract: contract}, ChallengeTesterTransactor: ChallengeTesterTransactor{contract: contract}, ChallengeTesterFilterer: ChallengeTesterFilterer{contract: contract}}, nil
}

// ChallengeTester is an auto generated Go binding around an Ethereum contract.
type ChallengeTester struct {
	ChallengeTesterCaller     // Read-only binding to the contract
	ChallengeTesterTransactor // Write-only binding to the contract
	ChallengeTesterFilterer   // Log filterer for contract events
}

// ChallengeTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeTesterSession struct {
	Contract     *ChallengeTester  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeTesterCallerSession struct {
	Contract *ChallengeTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ChallengeTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTesterTransactorSession struct {
	Contract     *ChallengeTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ChallengeTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeTesterRaw struct {
	Contract *ChallengeTester // Generic contract binding to access the raw methods on
}

// ChallengeTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeTesterCallerRaw struct {
	Contract *ChallengeTesterCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTesterTransactorRaw struct {
	Contract *ChallengeTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeTester creates a new instance of ChallengeTester, bound to a specific deployed contract.
func NewChallengeTester(address common.Address, backend bind.ContractBackend) (*ChallengeTester, error) {
	contract, err := bindChallengeTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeTester{ChallengeTesterCaller: ChallengeTesterCaller{contract: contract}, ChallengeTesterTransactor: ChallengeTesterTransactor{contract: contract}, ChallengeTesterFilterer: ChallengeTesterFilterer{contract: contract}}, nil
}

// NewChallengeTesterCaller creates a new read-only instance of ChallengeTester, bound to a specific deployed contract.
func NewChallengeTesterCaller(address common.Address, caller bind.ContractCaller) (*ChallengeTesterCaller, error) {
	contract, err := bindChallengeTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTesterCaller{contract: contract}, nil
}

// NewChallengeTesterTransactor creates a new write-only instance of ChallengeTester, bound to a specific deployed contract.
func NewChallengeTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTesterTransactor, error) {
	contract, err := bindChallengeTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTesterTransactor{contract: contract}, nil
}

// NewChallengeTesterFilterer creates a new log filterer instance of ChallengeTester, bound to a specific deployed contract.
func NewChallengeTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeTesterFilterer, error) {
	contract, err := bindChallengeTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeTesterFilterer{contract: contract}, nil
}

// bindChallengeTester binds a generic wrapper to an already deployed contract.
func bindChallengeTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeTester *ChallengeTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeTester.Contract.ChallengeTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeTester *ChallengeTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeTester.Contract.ChallengeTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeTester *ChallengeTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeTester.Contract.ChallengeTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeTester *ChallengeTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeTester *ChallengeTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeTester *ChallengeTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeTester.Contract.contract.Transact(opts, method, params...)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address , address ) returns()
func (_ChallengeTester *ChallengeTesterTransactor) ResolveChallenge(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _ChallengeTester.contract.Transact(opts, "resolveChallenge", arg0, arg1)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address , address ) returns()
func (_ChallengeTester *ChallengeTesterSession) ResolveChallenge(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _ChallengeTester.Contract.ResolveChallenge(&_ChallengeTester.TransactOpts, arg0, arg1)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address , address ) returns()
func (_ChallengeTester *ChallengeTesterTransactorSession) ResolveChallenge(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _ChallengeTester.Contract.ResolveChallenge(&_ChallengeTester.TransactOpts, arg0, arg1)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x8f43ee32.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterTransactor) StartChallenge(opts *bind.TransactOpts, asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.contract.Transact(opts, "startChallenge", asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x8f43ee32.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterSession) StartChallenge(asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.Contract.StartChallenge(&_ChallengeTester.TransactOpts, asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x8f43ee32.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterTransactorSession) StartChallenge(asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.Contract.StartChallenge(&_ChallengeTester.TransactOpts, asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}
