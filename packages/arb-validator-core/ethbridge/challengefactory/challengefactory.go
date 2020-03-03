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
const ChallengeFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messagesChallengeTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inboxTopChallengeTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executionChallengeTemplate\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"executionChallengeTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"generateCloneAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inboxTopChallengeTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"messagesChallengeTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFactoryFuncSigs = map[string]string{
	"865da1c2": "createChallenge(address,address,uint256,bytes32,uint256)",
	"9b5f4dfc": "executionChallengeTemplate()",
	"729406c8": "generateCloneAddress(address,address,uint256)",
	"f089f1c2": "inboxTopChallengeTemplate()",
	"e252f79a": "messagesChallengeTemplate()",
}

// ChallengeFactoryBin is the compiled bytecode used for deploying new contracts.
var ChallengeFactoryBin = "0x608060405234801561001057600080fd5b506040516107423803806107428339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b039384166001600160a01b03199182161790915560018054948416948216949094179093556002805492909116919092161790556106b38061008f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063729406c81461005c578063865da1c2146100ae5780639b5f4dfc146100f0578063e252f79a146100f8578063f089f1c214610100575b600080fd5b6100926004803603606081101561007257600080fd5b506001600160a01b03813581169160208101359091169060400135610108565b604080516001600160a01b039092168252519081900360200190f35b610092600480360360a08110156100c457600080fd5b506001600160a01b0381358116916020810135909116906040810135906060810135906080013561019f565b610092610246565b610092610255565b610092610264565b60006001600160f81b03193061011e8686610273565b61012f61012a866102c0565b6103bb565b60405160200180856001600160f81b0319166001600160f81b0319168152600101846001600160a01b03166001600160a01b031660601b81526014018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c90509392505050565b6000806101ab836102c0565b905060006101b8826104c5565b604080516301568f2760e11b81523360048201526001600160a01b038b811660248301528a81166044830152606482018a9052608482018990529151929350908316916302ad1e4e9160a48082019260009290919082900301818387803b15801561022257600080fd5b505af1158015610236573d6000803e3d6000fd5b50929a9950505050505050505050565b6002546001600160a01b031681565b6000546001600160a01b031681565b6001546001600160a01b031681565b604080516bffffffffffffffffffffffff19606094851b811660208084019190915293851b1660348201523390931b60488401528051603c818503018152605c9093019052815191012090565b6000816102d957506001546001600160a01b03166103b6565b60018214156102f457506000546001600160a01b03166103b6565b600282141561030f57506002546001600160a01b03166103b6565b60408051808201909152600c81526b494e56414c49445f5459504560a01b602082015260405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561037b578181015183820152602001610363565b50505050905090810190601f1680156103a85780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b919050565b60006060604051806020016103cf906105e2565b601f1982820381018352601f9091011660408181526001600160a01b038616602083810191909152815180840382018152828401909252835191926060019182918501908083835b602083106104365780518252601f199092019160209182019101610417565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061047e5780518252601f19909201916020918201910161045f565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290508080519060200120915050919050565b60006060604051806020016104d9906105e2565b601f1982820381018352601f9091011660408181526001600160a01b038616602083810191909152815180840382018152828401909252835191926060019182918501908083835b602083106105405780518252601f199092019160209182019101610521565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106105885780518252601f199092019160209182019101610569565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050806020018151808234f09350836105da573d6000803e3d6000fd5b505050919050565b6090806105ef8339019056fe6080604052348015600f57600080fd5b506040516090380380609083398181016040526020811015602f57600080fd5b50516040805169363d3d373d3d3d363d7360b01b6020828101919091526001600160601b0319606085901b16602a8301526e5af43d82803e903d91602b57fd5bf360881b603e8301528251602d81840381018252604d9093019093528201f3fea265627a7a7231582027cd24d3389191b68829b14f2587980146348c0aad2b7c8a6d5af0cfdb788a7f64736f6c63430005100032"

// DeployChallengeFactory deploys a new Ethereum contract, binding an instance of ChallengeFactory to it.
func DeployChallengeFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _messagesChallengeTemplate common.Address, _inboxTopChallengeTemplate common.Address, _executionChallengeTemplate common.Address) (common.Address, *types.Transaction, *ChallengeFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeFactoryBin), backend, _messagesChallengeTemplate, _inboxTopChallengeTemplate, _executionChallengeTemplate)
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

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) GenerateCloneAddress(opts *bind.CallOpts, asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "generateCloneAddress", asserter, challenger, challengeType)
	return *ret0, err
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) GenerateCloneAddress(asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	return _ChallengeFactory.Contract.GenerateCloneAddress(&_ChallengeFactory.CallOpts, asserter, challenger, challengeType)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) GenerateCloneAddress(asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	return _ChallengeFactory.Contract.GenerateCloneAddress(&_ChallengeFactory.CallOpts, asserter, challenger, challengeType)
}

// InboxTopChallengeTemplate is a free data retrieval call binding the contract method 0xf089f1c2.
//
// Solidity: function inboxTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) InboxTopChallengeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChallengeFactory.contract.Call(opts, out, "inboxTopChallengeTemplate")
	return *ret0, err
}

// InboxTopChallengeTemplate is a free data retrieval call binding the contract method 0xf089f1c2.
//
// Solidity: function inboxTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactorySession) InboxTopChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.InboxTopChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// InboxTopChallengeTemplate is a free data retrieval call binding the contract method 0xf089f1c2.
//
// Solidity: function inboxTopChallengeTemplate() constant returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) InboxTopChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.InboxTopChallengeTemplate(&_ChallengeFactory.CallOpts)
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
var ChallengeTypeBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a7231582040a322686e88dcf2eafa0d1379fc4e9d838048207f22426fd72b47241e53688c64736f6c63430005100032"

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
var CloneFactoryBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820d1d0c871538b9a28c627c2f3a94d6605b6ef1c7c751f818513babac79009f84664736f6c63430005100032"

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
const IChallengeFactoryABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"}],\"name\":\"generateCloneAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IChallengeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeFactoryFuncSigs = map[string]string{
	"865da1c2": "createChallenge(address,address,uint256,bytes32,uint256)",
	"729406c8": "generateCloneAddress(address,address,uint256)",
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

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_IChallengeFactory *IChallengeFactoryCaller) GenerateCloneAddress(opts *bind.CallOpts, asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IChallengeFactory.contract.Call(opts, out, "generateCloneAddress", asserter, challenger, challengeType)
	return *ret0, err
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_IChallengeFactory *IChallengeFactorySession) GenerateCloneAddress(asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	return _IChallengeFactory.Contract.GenerateCloneAddress(&_IChallengeFactory.CallOpts, asserter, challenger, challengeType)
}

// GenerateCloneAddress is a free data retrieval call binding the contract method 0x729406c8.
//
// Solidity: function generateCloneAddress(address asserter, address challenger, uint256 challengeType) constant returns(address)
func (_IChallengeFactory *IChallengeFactoryCallerSession) GenerateCloneAddress(asserter common.Address, challenger common.Address, challengeType *big.Int) (common.Address, error) {
	return _IChallengeFactory.Contract.GenerateCloneAddress(&_IChallengeFactory.CallOpts, asserter, challenger, challengeType)
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

// SpawnABI is the input ABI used to generate the binding from.
const SpawnABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"logicContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SpawnBin is the compiled bytecode used for deploying new contracts.
var SpawnBin = "0x6080604052348015600f57600080fd5b506040516090380380609083398181016040526020811015602f57600080fd5b50516040805169363d3d373d3d3d363d7360b01b6020828101919091526001600160601b0319606085901b16602a8301526e5af43d82803e903d91602b57fd5bf360881b603e8301528251602d81840381018252604d9093019093528201f3fe"

// DeploySpawn deploys a new Ethereum contract, binding an instance of Spawn to it.
func DeploySpawn(auth *bind.TransactOpts, backend bind.ContractBackend, logicContract common.Address) (common.Address, *types.Transaction, *Spawn, error) {
	parsed, err := abi.JSON(strings.NewReader(SpawnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SpawnBin), backend, logicContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Spawn{SpawnCaller: SpawnCaller{contract: contract}, SpawnTransactor: SpawnTransactor{contract: contract}, SpawnFilterer: SpawnFilterer{contract: contract}}, nil
}

// Spawn is an auto generated Go binding around an Ethereum contract.
type Spawn struct {
	SpawnCaller     // Read-only binding to the contract
	SpawnTransactor // Write-only binding to the contract
	SpawnFilterer   // Log filterer for contract events
}

// SpawnCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpawnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpawnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpawnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpawnSession struct {
	Contract     *Spawn            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpawnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpawnCallerSession struct {
	Contract *SpawnCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SpawnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpawnTransactorSession struct {
	Contract     *SpawnTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpawnRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpawnRaw struct {
	Contract *Spawn // Generic contract binding to access the raw methods on
}

// SpawnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpawnCallerRaw struct {
	Contract *SpawnCaller // Generic read-only contract binding to access the raw methods on
}

// SpawnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpawnTransactorRaw struct {
	Contract *SpawnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpawn creates a new instance of Spawn, bound to a specific deployed contract.
func NewSpawn(address common.Address, backend bind.ContractBackend) (*Spawn, error) {
	contract, err := bindSpawn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spawn{SpawnCaller: SpawnCaller{contract: contract}, SpawnTransactor: SpawnTransactor{contract: contract}, SpawnFilterer: SpawnFilterer{contract: contract}}, nil
}

// NewSpawnCaller creates a new read-only instance of Spawn, bound to a specific deployed contract.
func NewSpawnCaller(address common.Address, caller bind.ContractCaller) (*SpawnCaller, error) {
	contract, err := bindSpawn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpawnCaller{contract: contract}, nil
}

// NewSpawnTransactor creates a new write-only instance of Spawn, bound to a specific deployed contract.
func NewSpawnTransactor(address common.Address, transactor bind.ContractTransactor) (*SpawnTransactor, error) {
	contract, err := bindSpawn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpawnTransactor{contract: contract}, nil
}

// NewSpawnFilterer creates a new log filterer instance of Spawn, bound to a specific deployed contract.
func NewSpawnFilterer(address common.Address, filterer bind.ContractFilterer) (*SpawnFilterer, error) {
	contract, err := bindSpawn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpawnFilterer{contract: contract}, nil
}

// bindSpawn binds a generic wrapper to an already deployed contract.
func bindSpawn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SpawnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spawn *SpawnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spawn.Contract.SpawnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spawn *SpawnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.Contract.SpawnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spawn *SpawnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spawn.Contract.SpawnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spawn *SpawnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spawn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spawn *SpawnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spawn *SpawnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spawn.Contract.contract.Transact(opts, method, params...)
}
