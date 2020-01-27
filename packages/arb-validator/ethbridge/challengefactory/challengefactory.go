// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengefactory

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

// ChallengeFactoryABI is the input ABI used to generate the binding from.
const ChallengeFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messagesChallengeTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pendingTopChallengeTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executionChallengeTemplate\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"executionChallengeTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"name\":\"generateCloneAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"messagesChallengeTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingTopChallengeTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFactoryFuncSigs = map[string]string{
	"865da1c2": "createChallenge(address,address,uint256,bytes32,uint256)",
	"9b5f4dfc": "executionChallengeTemplate()",
	"c778f3f1": "generateCloneAddress(address,address,bytes32)",
	"e252f79a": "messagesChallengeTemplate()",
	"bef27054": "pendingTopChallengeTemplate()",
}

// ChallengeFactoryBin is the compiled bytecode used for deploying new contracts.
var ChallengeFactoryBin = "0x608060405234801561001057600080fd5b506040516104d93803806104d98339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b039384166001600160a01b031991821617909155600180549484169482169490941790935560028054929091169190921617905561044a8061008f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063865da1c21461005c5780639b5f4dfc146100ba578063bef27054146100c2578063c778f3f1146100ca578063e252f79a14610100575b600080fd5b61009e600480360360a081101561007257600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060800135610108565b604080516001600160a01b039092168252519081900360200190f35b61009e6101b9565b61009e6101c8565b61009e600480360360608110156100e057600080fd5b506001600160a01b038135811691602081013590911690604001356101d7565b61009e61025e565b6000806101148361026d565b9050600061012b826101268a8a610368565b6103b5565b604080516301568f2760e11b81523360048201526001600160a01b038b811660248301528a81166044830152606482018a9052608482018990529151929350908316916302ad1e4e9160a48082019260009290919082900301818387803b15801561019557600080fd5b505af11580156101a9573d6000803e3d6000fd5b50929a9950505050505050505050565b6002546001600160a01b031681565b6001546001600160a01b031681565b60006001600160f81b0319306101ed8686610368565b8460405160200180856001600160f81b0319166001600160f81b0319168152600101846001600160a01b03166001600160a01b031660601b81526014018381526020018281526020019450505050506040516020818303038152906040528051906020012060601c90509392505050565b6000546001600160a01b031681565b60008161028657506001546001600160a01b0316610363565b60018214156102a157506000546001600160a01b0316610363565b60028214156102bc57506002546001600160a01b0316610363565b60408051808201909152600c81526b494e56414c49445f5459504560a01b602082015260405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610328578181015183820152602001610310565b50505050905090810190601f1680156103555780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b919050565b604080516bffffffffffffffffffffffff19606094851b811660208084019190915293851b1660348201523390931b60488401528051603c818503018152605c9093019052815191012090565b60006103c0836103d2565b604051826037826000f5949350505050565b604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b815260609190911b60148201526e5af43d82803e903d91602b57fd5bf360881b60289091015256fea265627a7a72315820f56c5ad62b7ef039a74f1dc814cc2a94fb8c6d32c90fbb6abfd4804516fc5db464736f6c634300050f0032"

// DeployChallengeFactory deploys a new Ethereum contract, binding an instance of ChallengeFactory to it.
func DeployChallengeFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _messagesChallengeTemplate common.Address, _pendingTopChallengeTemplate common.Address, _executionChallengeTemplate common.Address) (common.Address, *types.Transaction, *ChallengeFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeFactoryBin), backend, _messagesChallengeTemplate, _pendingTopChallengeTemplate, _executionChallengeTemplate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeFactory{ChallengeFactoryCaller: ChallengeFactoryCaller{contract: contract}, ChallengeFactoryTransactor: ChallengeFactoryTransactor{contract: contract}, ChallengeFactoryFilterer: ChallengeFactoryFilterer{contract: contract}}, nil
}

// ChallengeFactory is an auto generated Go binding around an Ethereum contract.
type ChallengeFactory struct {
	ChallengeFactoryCaller     // Read-only binding to the contract
	ChallengeFactoryTransactor // Write-only binding to the contract
	ChallengeFactoryFilterer   // Log filterer for contract events
}

// ChallengeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeFactorySession struct {
	Contract     *ChallengeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeFactoryCallerSession struct {
	Contract *ChallengeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeFactoryTransactorSession struct {
	Contract     *ChallengeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeFactoryRaw struct {
	Contract *ChallengeFactory // Generic contract binding to access the raw methods on
}

// ChallengeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeFactoryCallerRaw struct {
	Contract *ChallengeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeFactoryTransactorRaw struct {
	Contract *ChallengeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeFactory creates a new instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactory(address common.Address, backend bind.ContractBackend) (*ChallengeFactory, error) {
	contract, err := bindChallengeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactory{ChallengeFactoryCaller: ChallengeFactoryCaller{contract: contract}, ChallengeFactoryTransactor: ChallengeFactoryTransactor{contract: contract}, ChallengeFactoryFilterer: ChallengeFactoryFilterer{contract: contract}}, nil
}

// NewChallengeFactoryCaller creates a new read-only instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryCaller(address common.Address, caller bind.ContractCaller) (*ChallengeFactoryCaller, error) {
	contract, err := bindChallengeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryCaller{contract: contract}, nil
}

// NewChallengeFactoryTransactor creates a new write-only instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeFactoryTransactor, error) {
	contract, err := bindChallengeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryTransactor{contract: contract}, nil
}

// NewChallengeFactoryFilterer creates a new log filterer instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFactoryFilterer, error) {
	contract, err := bindChallengeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryFilterer{contract: contract}, nil
}

// bindChallengeFactory binds a generic wrapper to an already deployed contract.
func bindChallengeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeFactory *ChallengeFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeFactory.Contract.ChallengeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeFactory *ChallengeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.ChallengeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeFactory *ChallengeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.ChallengeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeFactory *ChallengeFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeFactory *ChallengeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeFactory *ChallengeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.contract.Transact(opts, method, params...)
}

// ExecutionChallengeTemplate is a free data retrieval call binding the contract method 0x9b5f4dfc.
//
// Solidity: function executionChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) ExecutionChallengeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "executionChallengeTemplate")
	return *ret0, err
}

// ExecutionChallengeTemplate is a free data retrieval call binding the contract method 0x9b5f4dfc.
//
// Solidity: function executionChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) ExecutionChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.ExecutionChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// ExecutionChallengeTemplate is a free data retrieval call binding the contract method 0x9b5f4dfc.
//
// Solidity: function executionChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) ExecutionChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.ExecutionChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0xc778f3f1.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, bytes32 codeHash) constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) GenerateCloneAddress(opts *bind.CallOpts, asserter common.Address, challenger common.Address, codeHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "generateCloneAddress", asserter, challenger, codeHash)
	return *ret0, err
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0xc778f3f1.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, bytes32 codeHash) constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) GenerateCloneAddress(asserter common.Address, challenger common.Address, codeHash [32]byte) (common.Address, error) {
	return _ChallengeFactory.Contract.GenerateCloneAddress(&_ChallengeFactory.CallOpts, asserter, challenger, codeHash)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0xc778f3f1.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, bytes32 codeHash) constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) GenerateCloneAddress(asserter common.Address, challenger common.Address, codeHash [32]byte) (common.Address, error) {
	return _ChallengeFactory.Contract.GenerateCloneAddress(&_ChallengeFactory.CallOpts, asserter, challenger, codeHash)
}

// MessagesChallengeTemplate is a free data retrieval call binding the contract method 0xe252f79a.
//
// Solidity: function messagesChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) MessagesChallengeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "messagesChallengeTemplate")
	return *ret0, err
}

// MessagesChallengeTemplate is a free data retrieval call binding the contract method 0xe252f79a.
//
// Solidity: function messagesChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) MessagesChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.MessagesChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// MessagesChallengeTemplate is a free data retrieval call binding the contract method 0xe252f79a.
//
// Solidity: function messagesChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) MessagesChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.MessagesChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// PendingTopChallengeTemplate is a free data retrieval call binding the contract method 0xbef27054.
//
// Solidity: function pendingTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) PendingTopChallengeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "pendingTopChallengeTemplate")
	return *ret0, err
}

// PendingTopChallengeTemplate is a free data retrieval call binding the contract method 0xbef27054.
//
// Solidity: function pendingTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) PendingTopChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.PendingTopChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// PendingTopChallengeTemplate is a free data retrieval call binding the contract method 0xbef27054.
//
// Solidity: function pendingTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) PendingTopChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.PendingTopChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactor) CreateChallenge(opts *bind.TransactOpts, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeFactory.contract.Transact(opts, "createChallenge", _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_ChallengeFactory *ChallengeFactorySession) CreateChallenge(_asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x865da1c2.
//
// Solidity: function createChallenge(address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeHash, uint256 challengeType) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactorSession) CreateChallenge(_asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeHash [32]byte, challengeType *big.Int) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _asserter, _challenger, _challengePeriodTicks, _challengeHash, challengeType)
}

// ChallengeTypeABI is the input ABI used to generate the binding from.
const ChallengeTypeABI = "[]"

// ChallengeTypeBin is the compiled bytecode used for deploying new contracts.
var ChallengeTypeBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723158201bbf0efd3358d04cee283156a4da086f89c1f9888a7cabf5a3da50fac9a8d2e264736f6c634300050f0032"

// DeployChallengeType deploys a new Ethereum contract, binding an instance of ChallengeType to it.
func DeployChallengeType(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeType, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTypeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeTypeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeType{ChallengeTypeCaller: ChallengeTypeCaller{contract: contract}, ChallengeTypeTransactor: ChallengeTypeTransactor{contract: contract}, ChallengeTypeFilterer: ChallengeTypeFilterer{contract: contract}}, nil
}

// ChallengeType is an auto generated Go binding around an Ethereum contract.
type ChallengeType struct {
	ChallengeTypeCaller     // Read-only binding to the contract
	ChallengeTypeTransactor // Write-only binding to the contract
	ChallengeTypeFilterer   // Log filterer for contract events
}

// ChallengeTypeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeTypeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTypeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTypeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTypeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeTypeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTypeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeTypeSession struct {
	Contract     *ChallengeType    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeTypeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeTypeCallerSession struct {
	Contract *ChallengeTypeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ChallengeTypeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTypeTransactorSession struct {
	Contract     *ChallengeTypeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ChallengeTypeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeTypeRaw struct {
	Contract *ChallengeType // Generic contract binding to access the raw methods on
}

// ChallengeTypeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeTypeCallerRaw struct {
	Contract *ChallengeTypeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTypeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTypeTransactorRaw struct {
	Contract *ChallengeTypeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeType creates a new instance of ChallengeType, bound to a specific deployed contract.
func NewChallengeType(address common.Address, backend bind.ContractBackend) (*ChallengeType, error) {
	contract, err := bindChallengeType(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeType{ChallengeTypeCaller: ChallengeTypeCaller{contract: contract}, ChallengeTypeTransactor: ChallengeTypeTransactor{contract: contract}, ChallengeTypeFilterer: ChallengeTypeFilterer{contract: contract}}, nil
}

// NewChallengeTypeCaller creates a new read-only instance of ChallengeType, bound to a specific deployed contract.
func NewChallengeTypeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeTypeCaller, error) {
	contract, err := bindChallengeType(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTypeCaller{contract: contract}, nil
}

// NewChallengeTypeTransactor creates a new write-only instance of ChallengeType, bound to a specific deployed contract.
func NewChallengeTypeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTypeTransactor, error) {
	contract, err := bindChallengeType(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTypeTransactor{contract: contract}, nil
}

// NewChallengeTypeFilterer creates a new log filterer instance of ChallengeType, bound to a specific deployed contract.
func NewChallengeTypeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeTypeFilterer, error) {
	contract, err := bindChallengeType(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeTypeFilterer{contract: contract}, nil
}

// bindChallengeType binds a generic wrapper to an already deployed contract.
func bindChallengeType(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeTypeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeType *ChallengeTypeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeType.Contract.ChallengeTypeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeType *ChallengeTypeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeType.Contract.ChallengeTypeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeType *ChallengeTypeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeType.Contract.ChallengeTypeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeType *ChallengeTypeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeType.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeType *ChallengeTypeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeType.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeType *ChallengeTypeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeType.Contract.contract.Transact(opts, method, params...)
}

// CloneFactoryABI is the input ABI used to generate the binding from.
const CloneFactoryABI = "[]"

// CloneFactoryBin is the compiled bytecode used for deploying new contracts.
var CloneFactoryBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723158205904c63bcf1f3f83962d5ce7c3aaf5d7149a7ba75cb744d4d28720a3cb36dd3364736f6c634300050f0032"

// DeployCloneFactory deploys a new Ethereum contract, binding an instance of CloneFactory to it.
func DeployCloneFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CloneFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(CloneFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CloneFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CloneFactory{CloneFactoryCaller: CloneFactoryCaller{contract: contract}, CloneFactoryTransactor: CloneFactoryTransactor{contract: contract}, CloneFactoryFilterer: CloneFactoryFilterer{contract: contract}}, nil
}

// CloneFactory is an auto generated Go binding around an Ethereum contract.
type CloneFactory struct {
	CloneFactoryCaller     // Read-only binding to the contract
	CloneFactoryTransactor // Write-only binding to the contract
	CloneFactoryFilterer   // Log filterer for contract events
}

// CloneFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CloneFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CloneFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CloneFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CloneFactorySession struct {
	Contract     *CloneFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CloneFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CloneFactoryCallerSession struct {
	Contract *CloneFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CloneFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CloneFactoryTransactorSession struct {
	Contract     *CloneFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CloneFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CloneFactoryRaw struct {
	Contract *CloneFactory // Generic contract binding to access the raw methods on
}

// CloneFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CloneFactoryCallerRaw struct {
	Contract *CloneFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CloneFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CloneFactoryTransactorRaw struct {
	Contract *CloneFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCloneFactory creates a new instance of CloneFactory, bound to a specific deployed contract.
func NewCloneFactory(address common.Address, backend bind.ContractBackend) (*CloneFactory, error) {
	contract, err := bindCloneFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CloneFactory{CloneFactoryCaller: CloneFactoryCaller{contract: contract}, CloneFactoryTransactor: CloneFactoryTransactor{contract: contract}, CloneFactoryFilterer: CloneFactoryFilterer{contract: contract}}, nil
}

// NewCloneFactoryCaller creates a new read-only instance of CloneFactory, bound to a specific deployed contract.
func NewCloneFactoryCaller(address common.Address, caller bind.ContractCaller) (*CloneFactoryCaller, error) {
	contract, err := bindCloneFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CloneFactoryCaller{contract: contract}, nil
}

// NewCloneFactoryTransactor creates a new write-only instance of CloneFactory, bound to a specific deployed contract.
func NewCloneFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CloneFactoryTransactor, error) {
	contract, err := bindCloneFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CloneFactoryTransactor{contract: contract}, nil
}

// NewCloneFactoryFilterer creates a new log filterer instance of CloneFactory, bound to a specific deployed contract.
func NewCloneFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CloneFactoryFilterer, error) {
	contract, err := bindCloneFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CloneFactoryFilterer{contract: contract}, nil
}

// bindCloneFactory binds a generic wrapper to an already deployed contract.
func bindCloneFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CloneFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CloneFactory *CloneFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CloneFactory.Contract.CloneFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CloneFactory *CloneFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CloneFactory.Contract.CloneFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CloneFactory *CloneFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CloneFactory.Contract.CloneFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CloneFactory *CloneFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CloneFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CloneFactory *CloneFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CloneFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CloneFactory *CloneFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CloneFactory.Contract.contract.Transact(opts, method, params...)
}

// IBisectionChallengeABI is the input ABI used to generate the binding from.
const IBisectionChallengeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBisectionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var IBisectionChallengeFuncSigs = map[string]string{
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
}

// IBisectionChallenge is an auto generated Go binding around an Ethereum contract.
type IBisectionChallenge struct {
	IBisectionChallengeCaller     // Read-only binding to the contract
	IBisectionChallengeTransactor // Write-only binding to the contract
	IBisectionChallengeFilterer   // Log filterer for contract events
}

// IBisectionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBisectionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBisectionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBisectionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBisectionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBisectionChallengeSession struct {
	Contract     *IBisectionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IBisectionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBisectionChallengeCallerSession struct {
	Contract *IBisectionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IBisectionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBisectionChallengeTransactorSession struct {
	Contract     *IBisectionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IBisectionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBisectionChallengeRaw struct {
	Contract *IBisectionChallenge // Generic contract binding to access the raw methods on
}

// IBisectionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBisectionChallengeCallerRaw struct {
	Contract *IBisectionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// IBisectionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBisectionChallengeTransactorRaw struct {
	Contract *IBisectionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBisectionChallenge creates a new instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallenge(address common.Address, backend bind.ContractBackend) (*IBisectionChallenge, error) {
	contract, err := bindIBisectionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallenge{IBisectionChallengeCaller: IBisectionChallengeCaller{contract: contract}, IBisectionChallengeTransactor: IBisectionChallengeTransactor{contract: contract}, IBisectionChallengeFilterer: IBisectionChallengeFilterer{contract: contract}}, nil
}

// NewIBisectionChallengeCaller creates a new read-only instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeCaller(address common.Address, caller bind.ContractCaller) (*IBisectionChallengeCaller, error) {
	contract, err := bindIBisectionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeCaller{contract: contract}, nil
}

// NewIBisectionChallengeTransactor creates a new write-only instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBisectionChallengeTransactor, error) {
	contract, err := bindIBisectionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeTransactor{contract: contract}, nil
}

// NewIBisectionChallengeFilterer creates a new log filterer instance of IBisectionChallenge, bound to a specific deployed contract.
func NewIBisectionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBisectionChallengeFilterer, error) {
	contract, err := bindIBisectionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBisectionChallengeFilterer{contract: contract}, nil
}

// bindIBisectionChallenge binds a generic wrapper to an already deployed contract.
func bindIBisectionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBisectionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBisectionChallenge *IBisectionChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBisectionChallenge.Contract.IBisectionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBisectionChallenge *IBisectionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.IBisectionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBisectionChallenge *IBisectionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.IBisectionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBisectionChallenge *IBisectionChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBisectionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBisectionChallenge *IBisectionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBisectionChallenge *IBisectionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.contract.Transact(opts, method, params...)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.contract.Transact(opts, "initializeBisection", _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.InitializeBisection(&_IBisectionChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _vmAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_IBisectionChallenge *IBisectionChallengeTransactorSession) InitializeBisection(_vmAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _IBisectionChallenge.Contract.InitializeBisection(&_IBisectionChallenge.TransactOpts, _vmAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
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
