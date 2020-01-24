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
const ChallengeTesterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challengeFactory_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"asserterAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"challengerAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"challengerPeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"challengerDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"startChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeTesterFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeTesterFuncSigs = map[string]string{
	"396f51cf": "resolveChallenge(address,address)",
	"8f43ee32": "startChallenge(address,address,uint128,bytes32,uint256)",
}

// ChallengeTesterBin is the compiled bytecode used for deploying new contracts.
var ChallengeTesterBin = "0x608060405234801561001057600080fd5b506040516103a23803806103a28339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b031990921691909117905561033d806100656000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063396f51cf1461003b5780638f43ee321461006b575b600080fd5b6100696004803603604081101561005157600080fd5b506001600160a01b03813581169160200135166100b6565b005b610069600480360360a081101561008157600080fd5b506001600160a01b0381358116916020810135909116906001600160801b036040820135169060608101359060800135610236565b600080546040805163c778f3f160e01b81526001600160a01b038681166004830152858116602483015233803f60448401819052935190959394919091169163c778f3f1916064808301926020929190829003018186803b15801561011a57600080fd5b505afa15801561012e573d6000803e3d6000fd5b505050506040513d602081101561014457600080fd5b5051600080546040805163c778f3f160e01b81526001600160a01b038a811660048301528981166024830152604482018890529151949550929391169163c778f3f1916064808301926020929190829003018186803b1580156101a657600080fd5b505afa1580156101ba573d6000803e3d6000fd5b505050506040513d60208110156101d057600080fd5b505190506001600160a01b0382163314806101f357506001600160a01b03811633145b61022e5760405162461bcd60e51b81526004018080602001828103825260288152602001806102e16028913960400191505060405180910390fd5b505050505050565b600080546040805163432ed0e160e11b81526001600160a01b03898116600483015288811660248301526001600160801b038816604483015260648201879052608482018690529151919092169263865da1c29260a480820193602093909283900390910190829087803b1580156102ad57600080fd5b505af11580156102c1573d6000803e3d6000fd5b505050506040513d60208110156102d757600080fd5b5050505050505056fe4368616c6c656e6765206e6f7420636f6d706c657465642062792070726f7065722073656e646572a265627a7a72315820da4a3646496c2db6d409813e21528bda474e1352d156e343296dd8c84e8fe8f864736f6c634300050d0032"

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

// ResolveChallenge is a free data retrieval call binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) constant returns()
func (_ChallengeTester *ChallengeTesterCaller) ResolveChallenge(opts *bind.CallOpts, winner common.Address, loser common.Address) error {
	var ()
	out := &[]interface{}{}
	err := _ChallengeTester.contract.Call(opts, out, "resolveChallenge", winner, loser)
	return err
}

// ResolveChallenge is a free data retrieval call binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) constant returns()
func (_ChallengeTester *ChallengeTesterSession) ResolveChallenge(winner common.Address, loser common.Address) error {
	return _ChallengeTester.Contract.ResolveChallenge(&_ChallengeTester.CallOpts, winner, loser)
}

// ResolveChallenge is a free data retrieval call binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) constant returns()
func (_ChallengeTester *ChallengeTesterCallerSession) ResolveChallenge(winner common.Address, loser common.Address) error {
	return _ChallengeTester.Contract.ResolveChallenge(&_ChallengeTester.CallOpts, winner, loser)
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
