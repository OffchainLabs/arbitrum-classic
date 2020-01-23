// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengetester

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChallengeTesterABI is the input ABI used to generate the binding from.
const ChallengeTesterABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"challengeFactory\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"asserterAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"challengerAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"challengerPeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"challengerDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"startChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeTesterFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeTesterFuncSigs = map[string]string{
	"396f51cf": "resolveChallenge(address,address)",
	"3cda423d": "startChallenge(address,address,address,uint128,bytes32,uint256)",
}

// ChallengeTesterBin is the compiled bytecode used for deploying new contracts.
var ChallengeTesterBin = "0x608060405234801561001057600080fd5b5061019a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063396f51cf1461003b5780633cda423d1461006b575b600080fd5b6100696004803603604081101561005157600080fd5b506001600160a01b03813581169160200135166100bc565b005b610069600480360360c081101561008157600080fd5b506001600160a01b0381358116916020810135821691604082013516906001600160801b036060820135169060808101359060a001356100c0565b5050565b6040805163432ed0e160e11b81526001600160a01b03878116600483015286811660248301526001600160801b0386166044830152606482018590526084820184905291519188169163865da1c29160a4808201926020929091908290030181600087803b15801561013157600080fd5b505af1158015610145573d6000803e3d6000fd5b505050506040513d602081101561015b57600080fd5b505050505050505056fea265627a7a723158200d74090607810c506dfdf18c6298d75586c43641a3801e9a03772ae7da1ed74964736f6c634300050f0032"

// DeployChallengeTester deploys a new Ethereum contract, binding an instance of ChallengeTester to it.
func DeployChallengeTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeTester, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeTesterBin), backend)
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
func (_ChallengeTester *ChallengeTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ChallengeTester *ChallengeTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ChallengeTester *ChallengeTesterTransactor) ResolveChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ChallengeTester.contract.Transact(opts, "resolveChallenge", winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ChallengeTester *ChallengeTesterSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ChallengeTester.Contract.ResolveChallenge(&_ChallengeTester.TransactOpts, winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ChallengeTester *ChallengeTesterTransactorSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ChallengeTester.Contract.ResolveChallenge(&_ChallengeTester.TransactOpts, winner, loser)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x3cda423d.
//
// Solidity: function startChallenge(address challengeFactory, address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterTransactor) StartChallenge(opts *bind.TransactOpts, challengeFactory common.Address, asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.contract.Transact(opts, "startChallenge", challengeFactory, asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x3cda423d.
//
// Solidity: function startChallenge(address challengeFactory, address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterSession) StartChallenge(challengeFactory common.Address, asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.Contract.StartChallenge(&_ChallengeTester.TransactOpts, challengeFactory, asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}

// StartChallenge is a paid mutator transaction binding the contract method 0x3cda423d.
//
// Solidity: function startChallenge(address challengeFactory, address asserterAddress, address challengerAddress, uint128 challengerPeriodTicks, bytes32 challengerDataHash, uint256 challengeType) returns()
func (_ChallengeTester *ChallengeTesterTransactorSession) StartChallenge(challengeFactory common.Address, asserterAddress common.Address, challengerAddress common.Address, challengerPeriodTicks *big.Int, challengerDataHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeTester.Contract.StartChallenge(&_ChallengeTester.TransactOpts, challengeFactory, asserterAddress, challengerAddress, challengerPeriodTicks, challengerDataHash, challengeType)
}

// IChallengeFactoryABI is the input ABI used to generate the binding from.
const IChallengeFactoryABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"name\":\"generateCloneAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IChallengeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeFactoryFuncSigs = map[string]string{
	"865da1c2": "createChallenge(address,address,uint256,bytes32,uint256)",
	"c778f3f1": "generateCloneAddress(address,address,bytes32)",
}

// IChallengeFactory is an auto generated Go binding around an Ethereum contract.
type IChallengeFactory struct {
	IChallengeFactoryCaller     // Read-only binding to the contract
	IChallengeFactoryTransactor // Write-only binding to the contract
	IChallengeFactoryFilterer   // Log filterer for contract events
}

// IChallengeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeFactorySession struct {
	Contract     *IChallengeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeFactoryCallerSession struct {
	Contract *IChallengeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeFactoryTransactorSession struct {
	Contract     *IChallengeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeFactoryRaw struct {
	Contract *IChallengeFactory // Generic contract binding to access the raw methods on
}

// IChallengeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeFactoryCallerRaw struct {
	Contract *IChallengeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeFactoryTransactorRaw struct {
	Contract *IChallengeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeFactory creates a new instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactory(address common.Address, backend bind.ContractBackend) (*IChallengeFactory, error) {
	contract, err := bindIChallengeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactory{IChallengeFactoryCaller: IChallengeFactoryCaller{contract: contract}, IChallengeFactoryTransactor: IChallengeFactoryTransactor{contract: contract}, IChallengeFactoryFilterer: IChallengeFactoryFilterer{contract: contract}}, nil
}

// NewIChallengeFactoryCaller creates a new read-only instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryCaller(address common.Address, caller bind.ContractCaller) (*IChallengeFactoryCaller, error) {
	contract, err := bindIChallengeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryCaller{contract: contract}, nil
}

// NewIChallengeFactoryTransactor creates a new write-only instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeFactoryTransactor, error) {
	contract, err := bindIChallengeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryTransactor{contract: contract}, nil
}

// NewIChallengeFactoryFilterer creates a new log filterer instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeFactoryFilterer, error) {
	contract, err := bindIChallengeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryFilterer{contract: contract}, nil
}

// bindIChallengeFactory binds a generic wrapper to an already deployed contract.
func bindIChallengeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeFactory *IChallengeFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeFactory.Contract.IChallengeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeFactory *IChallengeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.IChallengeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeFactory *IChallengeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.IChallengeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeFactory *IChallengeFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeFactory *IChallengeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeFactory *IChallengeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.contract.Transact(opts, method, params...)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0xc778f3f1.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, bytes32 codeHash) constant returns(address)
func (_IChallengeFactory *IChallengeFactoryCaller) GenerateCloneAddress(opts *bind.CallOpts, asserter common.Address, challenger common.Address, codeHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IChallengeFactory.contract.Call(opts, out, "generateCloneAddress", asserter, challenger, codeHash)
	return *ret0, err
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0xc778f3f1.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, bytes32 codeHash) constant returns(address)
func (_IChallengeFactory *IChallengeFactorySession) GenerateCloneAddress(asserter common.Address, challenger common.Address, codeHash [32]byte) (common.Address, error) {
	return _IChallengeFactory.Contract.GenerateCloneAddress(&_IChallengeFactory.CallOpts, asserter, challenger, codeHash)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0xc778f3f1.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, bytes32 codeHash) constant returns(address)
func (_IChallengeFactory *IChallengeFactoryCallerSession) GenerateCloneAddress(asserter common.Address, challenger common.Address, codeHash [32]byte) (common.Address, error) {
	return _IChallengeFactory.Contract.GenerateCloneAddress(&_IChallengeFactory.CallOpts, asserter, challenger, codeHash)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_IChallengeFactory *IChallengeFactoryTransactor) CreateChallenge(opts *bind.TransactOpts, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _IChallengeFactory.contract.Transact(opts, "createChallenge", _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_IChallengeFactory *IChallengeFactorySession) CreateChallenge(_asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.CreateChallenge(&_IChallengeFactory.TransactOpts, _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_IChallengeFactory *IChallengeFactoryTransactorSession) CreateChallenge(_asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.CreateChallenge(&_IChallengeFactory.TransactOpts, _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}
