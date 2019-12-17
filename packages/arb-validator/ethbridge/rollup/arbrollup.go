// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbrollup

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

// ArbRollupABI is the input ABI used to generate the binding from.
const ArbRollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevLeafHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"beforeVMHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"beforeInboxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_afterVMHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"rollupAsserted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"rollupChallengeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"}],\"name\":\"rollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"rollupConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"rollupPruned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"rollupStakeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"toNodeHash\",\"type\":\"bytes32\"}],\"name\":\"rollupStakeMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"rollupStakeRefunded\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"stakerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakerProofOffsets\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"prev\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"branch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"location\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"createStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_stakeRequirement\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeVMHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInboxHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_prevPrevLeafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_prevDisputableHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prevChildType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_prevVMprotoHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prevLeafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_prevLeafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_stakerProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_afterVMHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterInboxHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numArbGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"makeAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newLocation\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof2\",\"type\":\"bytes32[]\"}],\"name\":\"moveStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"latestConfirmedProof\",\"type\":\"bytes32[]\"}],\"name\":\"pruneLeaf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"recoverStakeConfirmed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"disputableHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"latestConfirmedProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeProof\",\"type\":\"bytes32[]\"}],\"name\":\"recoverStakeMooted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker1Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2Address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"disputableDeadline\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"disputableHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"staker1position\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staker2position\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"vmProtoHash1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vmProtoHash2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof2\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"startChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"withinTimeBounds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbRollupFuncSigs maps the 4-byte function signature to its string representation.
var ArbRollupFuncSigs = map[string]string{
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"f6a7a291": "confirm(bytes32,uint256,bytes32[],bytes32[],uint256[],bytes32,uint256,uint256,bytes32,bytes32)",
	"d6903fdd": "createStake(bytes32,bytes32[])",
	"d489113a": "globalInbox()",
	"703320a5": "initialize(bytes32,uint32,uint32,uint128,address,address)",
	"367e96ca": "makeAssertion(bytes32,bytes32,bytes32,bytes32,uint256,bytes32,uint256,bytes32[],bytes32[],bytes32,bytes32,bytes32,bytes32,uint32,uint64,uint64[2])",
	"2d0bdcb4": "moveStake(bytes32,uint256,bytes32[],bytes32[])",
	"75702b8b": "pruneLeaf(uint256,bytes32,bytes32[],bytes32[])",
	"7cfaaf67": "recoverStakeConfirmed(bytes32[])",
	"8d28cd81": "recoverStakeMooted(bytes32,bytes32[],bytes32[])",
	"396f51cf": "resolveChallenge(address,address)",
	"e3c6fbbe": "startChallenge(address,address,bytes32,uint64,bytes32,uint256,uint256,bytes32,bytes32,bytes32[],bytes32[],bytes32,bytes32,uint64[2],bytes32)",
	"42c0787e": "withinTimeBounds(uint64[2])",
}

// ArbRollup is an auto generated Go binding around an Ethereum contract.
type ArbRollup struct {
	ArbRollupCaller     // Read-only binding to the contract
	ArbRollupTransactor // Write-only binding to the contract
	ArbRollupFilterer   // Log filterer for contract events
}

// ArbRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbRollupSession struct {
	Contract     *ArbRollup        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbRollupCallerSession struct {
	Contract *ArbRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArbRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbRollupTransactorSession struct {
	Contract     *ArbRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArbRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbRollupRaw struct {
	Contract *ArbRollup // Generic contract binding to access the raw methods on
}

// ArbRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbRollupCallerRaw struct {
	Contract *ArbRollupCaller // Generic read-only contract binding to access the raw methods on
}

// ArbRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbRollupTransactorRaw struct {
	Contract *ArbRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbRollup creates a new instance of ArbRollup, bound to a specific deployed contract.
func NewArbRollup(address common.Address, backend bind.ContractBackend) (*ArbRollup, error) {
	contract, err := bindArbRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbRollup{ArbRollupCaller: ArbRollupCaller{contract: contract}, ArbRollupTransactor: ArbRollupTransactor{contract: contract}, ArbRollupFilterer: ArbRollupFilterer{contract: contract}}, nil
}

// NewArbRollupCaller creates a new read-only instance of ArbRollup, bound to a specific deployed contract.
func NewArbRollupCaller(address common.Address, caller bind.ContractCaller) (*ArbRollupCaller, error) {
	contract, err := bindArbRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbRollupCaller{contract: contract}, nil
}

// NewArbRollupTransactor creates a new write-only instance of ArbRollup, bound to a specific deployed contract.
func NewArbRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbRollupTransactor, error) {
	contract, err := bindArbRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbRollupTransactor{contract: contract}, nil
}

// NewArbRollupFilterer creates a new log filterer instance of ArbRollup, bound to a specific deployed contract.
func NewArbRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbRollupFilterer, error) {
	contract, err := bindArbRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbRollupFilterer{contract: contract}, nil
}

// bindArbRollup binds a generic wrapper to an already deployed contract.
func bindArbRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbRollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbRollup *ArbRollupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbRollup.Contract.ArbRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbRollup *ArbRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbRollup.Contract.ArbRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbRollup *ArbRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbRollup.Contract.ArbRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbRollup *ArbRollupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbRollup *ArbRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbRollup *ArbRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbRollup.Contract.contract.Transact(opts, method, params...)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbRollup *ArbRollupCaller) GlobalInbox(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbRollup.contract.Call(opts, out, "globalInbox")
	return *ret0, err
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbRollup *ArbRollupSession) GlobalInbox() (common.Address, error) {
	return _ArbRollup.Contract.GlobalInbox(&_ArbRollup.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbRollup *ArbRollupCallerSession) GlobalInbox() (common.Address, error) {
	return _ArbRollup.Contract.GlobalInbox(&_ArbRollup.CallOpts)
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_ArbRollup *ArbRollupCaller) WithinTimeBounds(opts *bind.CallOpts, _timeBounds [2]uint64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbRollup.contract.Call(opts, out, "withinTimeBounds", _timeBounds)
	return *ret0, err
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_ArbRollup *ArbRollupSession) WithinTimeBounds(_timeBounds [2]uint64) (bool, error) {
	return _ArbRollup.Contract.WithinTimeBounds(&_ArbRollup.CallOpts, _timeBounds)
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_ArbRollup *ArbRollupCallerSession) WithinTimeBounds(_timeBounds [2]uint64) (bool, error) {
	return _ArbRollup.Contract.WithinTimeBounds(&_ArbRollup.CallOpts, _timeBounds)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbRollup *ArbRollupTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbRollup *ArbRollupSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.CompleteChallenge(&_ArbRollup.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbRollup *ArbRollupTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.CompleteChallenge(&_ArbRollup.TransactOpts, _players, _rewards)
}

// Confirm is a paid mutator transaction binding the contract method 0xf6a7a291.
//
// Solidity: function confirm(bytes32 to, uint256 _leafIndex, bytes32[] proof1, bytes32[] stakerProofs, uint256[] stakerProofOffsets, bytes32 prev, uint256 branch, uint256 deadline, bytes32 _preconditionHash, bytes32 _assertionHash) returns()
func (_ArbRollup *ArbRollupTransactor) Confirm(opts *bind.TransactOpts, to [32]byte, _leafIndex *big.Int, proof1 [][32]byte, stakerProofs [][32]byte, stakerProofOffsets []*big.Int, prev [32]byte, branch *big.Int, deadline *big.Int, _preconditionHash [32]byte, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "confirm", to, _leafIndex, proof1, stakerProofs, stakerProofOffsets, prev, branch, deadline, _preconditionHash, _assertionHash)
}

// Confirm is a paid mutator transaction binding the contract method 0xf6a7a291.
//
// Solidity: function confirm(bytes32 to, uint256 _leafIndex, bytes32[] proof1, bytes32[] stakerProofs, uint256[] stakerProofOffsets, bytes32 prev, uint256 branch, uint256 deadline, bytes32 _preconditionHash, bytes32 _assertionHash) returns()
func (_ArbRollup *ArbRollupSession) Confirm(to [32]byte, _leafIndex *big.Int, proof1 [][32]byte, stakerProofs [][32]byte, stakerProofOffsets []*big.Int, prev [32]byte, branch *big.Int, deadline *big.Int, _preconditionHash [32]byte, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.Confirm(&_ArbRollup.TransactOpts, to, _leafIndex, proof1, stakerProofs, stakerProofOffsets, prev, branch, deadline, _preconditionHash, _assertionHash)
}

// Confirm is a paid mutator transaction binding the contract method 0xf6a7a291.
//
// Solidity: function confirm(bytes32 to, uint256 _leafIndex, bytes32[] proof1, bytes32[] stakerProofs, uint256[] stakerProofOffsets, bytes32 prev, uint256 branch, uint256 deadline, bytes32 _preconditionHash, bytes32 _assertionHash) returns()
func (_ArbRollup *ArbRollupTransactorSession) Confirm(to [32]byte, _leafIndex *big.Int, proof1 [][32]byte, stakerProofs [][32]byte, stakerProofOffsets []*big.Int, prev [32]byte, branch *big.Int, deadline *big.Int, _preconditionHash [32]byte, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.Confirm(&_ArbRollup.TransactOpts, to, _leafIndex, proof1, stakerProofs, stakerProofOffsets, prev, branch, deadline, _preconditionHash, _assertionHash)
}

// CreateStake is a paid mutator transaction binding the contract method 0xd6903fdd.
//
// Solidity: function createStake(bytes32 location, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactor) CreateStake(opts *bind.TransactOpts, location [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "createStake", location, proof)
}

// CreateStake is a paid mutator transaction binding the contract method 0xd6903fdd.
//
// Solidity: function createStake(bytes32 location, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupSession) CreateStake(location [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.CreateStake(&_ArbRollup.TransactOpts, location, proof)
}

// CreateStake is a paid mutator transaction binding the contract method 0xd6903fdd.
//
// Solidity: function createStake(bytes32 location, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactorSession) CreateStake(location [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.CreateStake(&_ArbRollup.TransactOpts, location, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0x703320a5.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _stakeRequirement, address _owner, address _globalInboxAddress) returns()
func (_ArbRollup *ArbRollupTransactor) Initialize(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _stakeRequirement *big.Int, _owner common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "initialize", _vmState, _gracePeriod, _maxExecutionSteps, _stakeRequirement, _owner, _globalInboxAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x703320a5.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _stakeRequirement, address _owner, address _globalInboxAddress) returns()
func (_ArbRollup *ArbRollupSession) Initialize(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _stakeRequirement *big.Int, _owner common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.Initialize(&_ArbRollup.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _stakeRequirement, _owner, _globalInboxAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x703320a5.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _stakeRequirement, address _owner, address _globalInboxAddress) returns()
func (_ArbRollup *ArbRollupTransactorSession) Initialize(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _stakeRequirement *big.Int, _owner common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.Initialize(&_ArbRollup.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _stakeRequirement, _owner, _globalInboxAddress)
}

// MakeAssertion is a paid mutator transaction binding the contract method 0x367e96ca.
//
// Solidity: function makeAssertion(bytes32 beforeVMHash, bytes32 beforeInboxHash, bytes32 _prevPrevLeafHash, bytes32 _prevDisputableHash, uint256 _prevChildType, bytes32 _prevVMprotoHash, uint256 _prevLeafIndex, bytes32[] _prevLeafProof, bytes32[] _stakerProof, bytes32 _afterVMHash, bytes32 _afterInboxHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numArbGas, uint64[2] _timeBounds) returns()
func (_ArbRollup *ArbRollupTransactor) MakeAssertion(opts *bind.TransactOpts, beforeVMHash [32]byte, beforeInboxHash [32]byte, _prevPrevLeafHash [32]byte, _prevDisputableHash [32]byte, _prevChildType *big.Int, _prevVMprotoHash [32]byte, _prevLeafIndex *big.Int, _prevLeafProof [][32]byte, _stakerProof [][32]byte, _afterVMHash [32]byte, _afterInboxHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numArbGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "makeAssertion", beforeVMHash, beforeInboxHash, _prevPrevLeafHash, _prevDisputableHash, _prevChildType, _prevVMprotoHash, _prevLeafIndex, _prevLeafProof, _stakerProof, _afterVMHash, _afterInboxHash, _messagesAccHash, _logsAccHash, _numSteps, _numArbGas, _timeBounds)
}

// MakeAssertion is a paid mutator transaction binding the contract method 0x367e96ca.
//
// Solidity: function makeAssertion(bytes32 beforeVMHash, bytes32 beforeInboxHash, bytes32 _prevPrevLeafHash, bytes32 _prevDisputableHash, uint256 _prevChildType, bytes32 _prevVMprotoHash, uint256 _prevLeafIndex, bytes32[] _prevLeafProof, bytes32[] _stakerProof, bytes32 _afterVMHash, bytes32 _afterInboxHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numArbGas, uint64[2] _timeBounds) returns()
func (_ArbRollup *ArbRollupSession) MakeAssertion(beforeVMHash [32]byte, beforeInboxHash [32]byte, _prevPrevLeafHash [32]byte, _prevDisputableHash [32]byte, _prevChildType *big.Int, _prevVMprotoHash [32]byte, _prevLeafIndex *big.Int, _prevLeafProof [][32]byte, _stakerProof [][32]byte, _afterVMHash [32]byte, _afterInboxHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numArbGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbRollup.Contract.MakeAssertion(&_ArbRollup.TransactOpts, beforeVMHash, beforeInboxHash, _prevPrevLeafHash, _prevDisputableHash, _prevChildType, _prevVMprotoHash, _prevLeafIndex, _prevLeafProof, _stakerProof, _afterVMHash, _afterInboxHash, _messagesAccHash, _logsAccHash, _numSteps, _numArbGas, _timeBounds)
}

// MakeAssertion is a paid mutator transaction binding the contract method 0x367e96ca.
//
// Solidity: function makeAssertion(bytes32 beforeVMHash, bytes32 beforeInboxHash, bytes32 _prevPrevLeafHash, bytes32 _prevDisputableHash, uint256 _prevChildType, bytes32 _prevVMprotoHash, uint256 _prevLeafIndex, bytes32[] _prevLeafProof, bytes32[] _stakerProof, bytes32 _afterVMHash, bytes32 _afterInboxHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numArbGas, uint64[2] _timeBounds) returns()
func (_ArbRollup *ArbRollupTransactorSession) MakeAssertion(beforeVMHash [32]byte, beforeInboxHash [32]byte, _prevPrevLeafHash [32]byte, _prevDisputableHash [32]byte, _prevChildType *big.Int, _prevVMprotoHash [32]byte, _prevLeafIndex *big.Int, _prevLeafProof [][32]byte, _stakerProof [][32]byte, _afterVMHash [32]byte, _afterInboxHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numArbGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbRollup.Contract.MakeAssertion(&_ArbRollup.TransactOpts, beforeVMHash, beforeInboxHash, _prevPrevLeafHash, _prevDisputableHash, _prevChildType, _prevVMprotoHash, _prevLeafIndex, _prevLeafProof, _stakerProof, _afterVMHash, _afterInboxHash, _messagesAccHash, _logsAccHash, _numSteps, _numArbGas, _timeBounds)
}

// MoveStake is a paid mutator transaction binding the contract method 0x2d0bdcb4.
//
// Solidity: function moveStake(bytes32 newLocation, uint256 _leafIndex, bytes32[] proof1, bytes32[] proof2) returns()
func (_ArbRollup *ArbRollupTransactor) MoveStake(opts *bind.TransactOpts, newLocation [32]byte, _leafIndex *big.Int, proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "moveStake", newLocation, _leafIndex, proof1, proof2)
}

// MoveStake is a paid mutator transaction binding the contract method 0x2d0bdcb4.
//
// Solidity: function moveStake(bytes32 newLocation, uint256 _leafIndex, bytes32[] proof1, bytes32[] proof2) returns()
func (_ArbRollup *ArbRollupSession) MoveStake(newLocation [32]byte, _leafIndex *big.Int, proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.MoveStake(&_ArbRollup.TransactOpts, newLocation, _leafIndex, proof1, proof2)
}

// MoveStake is a paid mutator transaction binding the contract method 0x2d0bdcb4.
//
// Solidity: function moveStake(bytes32 newLocation, uint256 _leafIndex, bytes32[] proof1, bytes32[] proof2) returns()
func (_ArbRollup *ArbRollupTransactorSession) MoveStake(newLocation [32]byte, _leafIndex *big.Int, proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.MoveStake(&_ArbRollup.TransactOpts, newLocation, _leafIndex, proof1, proof2)
}

// PruneLeaf is a paid mutator transaction binding the contract method 0x75702b8b.
//
// Solidity: function pruneLeaf(uint256 _leafIndex, bytes32 from, bytes32[] leafProof, bytes32[] latestConfirmedProof) returns()
func (_ArbRollup *ArbRollupTransactor) PruneLeaf(opts *bind.TransactOpts, _leafIndex *big.Int, from [32]byte, leafProof [][32]byte, latestConfirmedProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "pruneLeaf", _leafIndex, from, leafProof, latestConfirmedProof)
}

// PruneLeaf is a paid mutator transaction binding the contract method 0x75702b8b.
//
// Solidity: function pruneLeaf(uint256 _leafIndex, bytes32 from, bytes32[] leafProof, bytes32[] latestConfirmedProof) returns()
func (_ArbRollup *ArbRollupSession) PruneLeaf(_leafIndex *big.Int, from [32]byte, leafProof [][32]byte, latestConfirmedProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.PruneLeaf(&_ArbRollup.TransactOpts, _leafIndex, from, leafProof, latestConfirmedProof)
}

// PruneLeaf is a paid mutator transaction binding the contract method 0x75702b8b.
//
// Solidity: function pruneLeaf(uint256 _leafIndex, bytes32 from, bytes32[] leafProof, bytes32[] latestConfirmedProof) returns()
func (_ArbRollup *ArbRollupTransactorSession) PruneLeaf(_leafIndex *big.Int, from [32]byte, leafProof [][32]byte, latestConfirmedProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.PruneLeaf(&_ArbRollup.TransactOpts, _leafIndex, from, leafProof, latestConfirmedProof)
}

// RecoverStakeConfirmed is a paid mutator transaction binding the contract method 0x7cfaaf67.
//
// Solidity: function recoverStakeConfirmed(bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactor) RecoverStakeConfirmed(opts *bind.TransactOpts, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "recoverStakeConfirmed", proof)
}

// RecoverStakeConfirmed is a paid mutator transaction binding the contract method 0x7cfaaf67.
//
// Solidity: function recoverStakeConfirmed(bytes32[] proof) returns()
func (_ArbRollup *ArbRollupSession) RecoverStakeConfirmed(proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeConfirmed(&_ArbRollup.TransactOpts, proof)
}

// RecoverStakeConfirmed is a paid mutator transaction binding the contract method 0x7cfaaf67.
//
// Solidity: function recoverStakeConfirmed(bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactorSession) RecoverStakeConfirmed(proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeConfirmed(&_ArbRollup.TransactOpts, proof)
}

// RecoverStakeMooted is a paid mutator transaction binding the contract method 0x8d28cd81.
//
// Solidity: function recoverStakeMooted(bytes32 disputableHash, bytes32[] latestConfirmedProof, bytes32[] nodeProof) returns()
func (_ArbRollup *ArbRollupTransactor) RecoverStakeMooted(opts *bind.TransactOpts, disputableHash [32]byte, latestConfirmedProof [][32]byte, nodeProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "recoverStakeMooted", disputableHash, latestConfirmedProof, nodeProof)
}

// RecoverStakeMooted is a paid mutator transaction binding the contract method 0x8d28cd81.
//
// Solidity: function recoverStakeMooted(bytes32 disputableHash, bytes32[] latestConfirmedProof, bytes32[] nodeProof) returns()
func (_ArbRollup *ArbRollupSession) RecoverStakeMooted(disputableHash [32]byte, latestConfirmedProof [][32]byte, nodeProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeMooted(&_ArbRollup.TransactOpts, disputableHash, latestConfirmedProof, nodeProof)
}

// RecoverStakeMooted is a paid mutator transaction binding the contract method 0x8d28cd81.
//
// Solidity: function recoverStakeMooted(bytes32 disputableHash, bytes32[] latestConfirmedProof, bytes32[] nodeProof) returns()
func (_ArbRollup *ArbRollupTransactorSession) RecoverStakeMooted(disputableHash [32]byte, latestConfirmedProof [][32]byte, nodeProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeMooted(&_ArbRollup.TransactOpts, disputableHash, latestConfirmedProof, nodeProof)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ArbRollup *ArbRollupTransactor) ResolveChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "resolveChallenge", winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ArbRollup *ArbRollupSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.ResolveChallenge(&_ArbRollup.TransactOpts, winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ArbRollup *ArbRollupTransactorSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.ResolveChallenge(&_ArbRollup.TransactOpts, winner, loser)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xe3c6fbbe.
//
// Solidity: function startChallenge(address staker1Address, address staker2Address, bytes32 node, uint64 disputableDeadline, bytes32 disputableHash, uint256 staker1position, uint256 staker2position, bytes32 vmProtoHash1, bytes32 vmProtoHash2, bytes32[] proof1, bytes32[] proof2, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbRollup *ArbRollupTransactor) StartChallenge(opts *bind.TransactOpts, staker1Address common.Address, staker2Address common.Address, node [32]byte, disputableDeadline uint64, disputableHash [32]byte, staker1position *big.Int, staker2position *big.Int, vmProtoHash1 [32]byte, vmProtoHash2 [32]byte, proof1 [][32]byte, proof2 [][32]byte, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "startChallenge", staker1Address, staker2Address, node, disputableDeadline, disputableHash, staker1position, staker2position, vmProtoHash1, vmProtoHash2, proof1, proof2, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xe3c6fbbe.
//
// Solidity: function startChallenge(address staker1Address, address staker2Address, bytes32 node, uint64 disputableDeadline, bytes32 disputableHash, uint256 staker1position, uint256 staker2position, bytes32 vmProtoHash1, bytes32 vmProtoHash2, bytes32[] proof1, bytes32[] proof2, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbRollup *ArbRollupSession) StartChallenge(staker1Address common.Address, staker2Address common.Address, node [32]byte, disputableDeadline uint64, disputableHash [32]byte, staker1position *big.Int, staker2position *big.Int, vmProtoHash1 [32]byte, vmProtoHash2 [32]byte, proof1 [][32]byte, proof2 [][32]byte, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.StartChallenge(&_ArbRollup.TransactOpts, staker1Address, staker2Address, node, disputableDeadline, disputableHash, staker1position, staker2position, vmProtoHash1, vmProtoHash2, proof1, proof2, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xe3c6fbbe.
//
// Solidity: function startChallenge(address staker1Address, address staker2Address, bytes32 node, uint64 disputableDeadline, bytes32 disputableHash, uint256 staker1position, uint256 staker2position, bytes32 vmProtoHash1, bytes32 vmProtoHash2, bytes32[] proof1, bytes32[] proof2, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbRollup *ArbRollupTransactorSession) StartChallenge(staker1Address common.Address, staker2Address common.Address, node [32]byte, disputableDeadline uint64, disputableHash [32]byte, staker1position *big.Int, staker2position *big.Int, vmProtoHash1 [32]byte, vmProtoHash2 [32]byte, proof1 [][32]byte, proof2 [][32]byte, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.StartChallenge(&_ArbRollup.TransactOpts, staker1Address, staker2Address, node, disputableDeadline, disputableHash, staker1position, staker2position, vmProtoHash1, vmProtoHash2, proof1, proof2, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// ArbRollupRollupAssertedIterator is returned from FilterRollupAsserted and is used to iterate over the raw logs and unpacked data for RollupAsserted events raised by the ArbRollup contract.
type ArbRollupRollupAssertedIterator struct {
	Event *ArbRollupRollupAsserted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupAssertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupAsserted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupAsserted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupAssertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupAssertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupAsserted represents a RollupAsserted event raised by the ArbRollup contract.
type ArbRollupRollupAsserted struct {
	PrevLeafHash    [32]byte
	BeforeVMHash    [32]byte
	TimeBounds      [2]uint64
	BeforeInboxHash [32]byte
	AfterVMHash     [32]byte
	NumSteps        uint32
	NumArbGas       uint64
	MessagesAccHash [32]byte
	LogsAccHash     [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRollupAsserted is a free log retrieval operation binding the contract event 0x7d534d82e41d1e3e4e8346219de6c8b579b8d348396f7de001d0e442fa3eccea.
//
// Solidity: event rollupAsserted(bytes32 prevLeafHash, bytes32 beforeVMHash, uint64[2] _timeBounds, bytes32 beforeInboxHash, bytes32 _afterVMHash, uint32 _numSteps, uint64 _numArbGas, bytes32 _messagesAccHash, bytes32 _logsAccHash)
func (_ArbRollup *ArbRollupFilterer) FilterRollupAsserted(opts *bind.FilterOpts) (*ArbRollupRollupAssertedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "rollupAsserted")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupAssertedIterator{contract: _ArbRollup.contract, event: "rollupAsserted", logs: logs, sub: sub}, nil
}

// WatchRollupAsserted is a free log subscription operation binding the contract event 0x7d534d82e41d1e3e4e8346219de6c8b579b8d348396f7de001d0e442fa3eccea.
//
// Solidity: event rollupAsserted(bytes32 prevLeafHash, bytes32 beforeVMHash, uint64[2] _timeBounds, bytes32 beforeInboxHash, bytes32 _afterVMHash, uint32 _numSteps, uint64 _numArbGas, bytes32 _messagesAccHash, bytes32 _logsAccHash)
func (_ArbRollup *ArbRollupFilterer) WatchRollupAsserted(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupAsserted) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "rollupAsserted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupAsserted)
				if err := _ArbRollup.contract.UnpackLog(event, "rollupAsserted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupAsserted is a log parse operation binding the contract event 0x7d534d82e41d1e3e4e8346219de6c8b579b8d348396f7de001d0e442fa3eccea.
//
// Solidity: event rollupAsserted(bytes32 prevLeafHash, bytes32 beforeVMHash, uint64[2] _timeBounds, bytes32 beforeInboxHash, bytes32 _afterVMHash, uint32 _numSteps, uint64 _numArbGas, bytes32 _messagesAccHash, bytes32 _logsAccHash)
func (_ArbRollup *ArbRollupFilterer) ParseRollupAsserted(log types.Log) (*ArbRollupRollupAsserted, error) {
	event := new(ArbRollupRollupAsserted)
	if err := _ArbRollup.contract.UnpackLog(event, "rollupAsserted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbRollupRollupChallengeCompletedIterator is returned from FilterRollupChallengeCompleted and is used to iterate over the raw logs and unpacked data for RollupChallengeCompleted events raised by the ArbRollup contract.
type ArbRollupRollupChallengeCompletedIterator struct {
	Event *ArbRollupRollupChallengeCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupChallengeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupChallengeCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupChallengeCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupChallengeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupChallengeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupChallengeCompleted represents a RollupChallengeCompleted event raised by the ArbRollup contract.
type ArbRollupRollupChallengeCompleted struct {
	ChallengeContract common.Address
	Winner            common.Address
	Loser             common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeCompleted is a free log retrieval operation binding the contract event 0xcf28c40762468afb8d98ba08e82252198424a5b4db2ee0434a85016027faa60b.
//
// Solidity: event rollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_ArbRollup *ArbRollupFilterer) FilterRollupChallengeCompleted(opts *bind.FilterOpts) (*ArbRollupRollupChallengeCompletedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "rollupChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupChallengeCompletedIterator{contract: _ArbRollup.contract, event: "rollupChallengeCompleted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeCompleted is a free log subscription operation binding the contract event 0xcf28c40762468afb8d98ba08e82252198424a5b4db2ee0434a85016027faa60b.
//
// Solidity: event rollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_ArbRollup *ArbRollupFilterer) WatchRollupChallengeCompleted(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupChallengeCompleted) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "rollupChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupChallengeCompleted)
				if err := _ArbRollup.contract.UnpackLog(event, "rollupChallengeCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupChallengeCompleted is a log parse operation binding the contract event 0xcf28c40762468afb8d98ba08e82252198424a5b4db2ee0434a85016027faa60b.
//
// Solidity: event rollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_ArbRollup *ArbRollupFilterer) ParseRollupChallengeCompleted(log types.Log) (*ArbRollupRollupChallengeCompleted, error) {
	event := new(ArbRollupRollupChallengeCompleted)
	if err := _ArbRollup.contract.UnpackLog(event, "rollupChallengeCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbRollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the ArbRollup contract.
type ArbRollupRollupChallengeStartedIterator struct {
	Event *ArbRollupRollupChallengeStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupChallengeStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupChallengeStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the ArbRollup contract.
type ArbRollupRollupChallengeStarted struct {
	Asserter          common.Address
	Challenger        common.Address
	ChallengeType     *big.Int
	ChallengeContract common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0xf5f442014ccab1bef67bc52b11a6e9e4ee5b1c70a0d306ee795420b1d599a0c4.
//
// Solidity: event rollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_ArbRollup *ArbRollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts) (*ArbRollupRollupChallengeStartedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "rollupChallengeStarted")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupChallengeStartedIterator{contract: _ArbRollup.contract, event: "rollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0xf5f442014ccab1bef67bc52b11a6e9e4ee5b1c70a0d306ee795420b1d599a0c4.
//
// Solidity: event rollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_ArbRollup *ArbRollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupChallengeStarted) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "rollupChallengeStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupChallengeStarted)
				if err := _ArbRollup.contract.UnpackLog(event, "rollupChallengeStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0xf5f442014ccab1bef67bc52b11a6e9e4ee5b1c70a0d306ee795420b1d599a0c4.
//
// Solidity: event rollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_ArbRollup *ArbRollupFilterer) ParseRollupChallengeStarted(log types.Log) (*ArbRollupRollupChallengeStarted, error) {
	event := new(ArbRollupRollupChallengeStarted)
	if err := _ArbRollup.contract.UnpackLog(event, "rollupChallengeStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbRollupRollupConfirmedIterator is returned from FilterRollupConfirmed and is used to iterate over the raw logs and unpacked data for RollupConfirmed events raised by the ArbRollup contract.
type ArbRollupRollupConfirmedIterator struct {
	Event *ArbRollupRollupConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupConfirmed represents a RollupConfirmed event raised by the ArbRollup contract.
type ArbRollupRollupConfirmed struct {
	NodeHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollupConfirmed is a free log retrieval operation binding the contract event 0xde207829ddeaef46a01a54fab518baafbacdcf7720d040a6586f40d7831b6d90.
//
// Solidity: event rollupConfirmed(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) FilterRollupConfirmed(opts *bind.FilterOpts) (*ArbRollupRollupConfirmedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "rollupConfirmed")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupConfirmedIterator{contract: _ArbRollup.contract, event: "rollupConfirmed", logs: logs, sub: sub}, nil
}

// WatchRollupConfirmed is a free log subscription operation binding the contract event 0xde207829ddeaef46a01a54fab518baafbacdcf7720d040a6586f40d7831b6d90.
//
// Solidity: event rollupConfirmed(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) WatchRollupConfirmed(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupConfirmed) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "rollupConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupConfirmed)
				if err := _ArbRollup.contract.UnpackLog(event, "rollupConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupConfirmed is a log parse operation binding the contract event 0xde207829ddeaef46a01a54fab518baafbacdcf7720d040a6586f40d7831b6d90.
//
// Solidity: event rollupConfirmed(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) ParseRollupConfirmed(log types.Log) (*ArbRollupRollupConfirmed, error) {
	event := new(ArbRollupRollupConfirmed)
	if err := _ArbRollup.contract.UnpackLog(event, "rollupConfirmed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbRollupRollupPrunedIterator is returned from FilterRollupPruned and is used to iterate over the raw logs and unpacked data for RollupPruned events raised by the ArbRollup contract.
type ArbRollupRollupPrunedIterator struct {
	Event *ArbRollupRollupPruned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupPrunedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupPruned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupPruned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupPrunedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupPrunedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupPruned represents a RollupPruned event raised by the ArbRollup contract.
type ArbRollupRollupPruned struct {
	NodeHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollupPruned is a free log retrieval operation binding the contract event 0xdecdc5d807444a59805ca7f59f6deff8010a82b725e43d8e5f437f51f0150930.
//
// Solidity: event rollupPruned(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) FilterRollupPruned(opts *bind.FilterOpts) (*ArbRollupRollupPrunedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "rollupPruned")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupPrunedIterator{contract: _ArbRollup.contract, event: "rollupPruned", logs: logs, sub: sub}, nil
}

// WatchRollupPruned is a free log subscription operation binding the contract event 0xdecdc5d807444a59805ca7f59f6deff8010a82b725e43d8e5f437f51f0150930.
//
// Solidity: event rollupPruned(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) WatchRollupPruned(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupPruned) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "rollupPruned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupPruned)
				if err := _ArbRollup.contract.UnpackLog(event, "rollupPruned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupPruned is a log parse operation binding the contract event 0xdecdc5d807444a59805ca7f59f6deff8010a82b725e43d8e5f437f51f0150930.
//
// Solidity: event rollupPruned(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) ParseRollupPruned(log types.Log) (*ArbRollupRollupPruned, error) {
	event := new(ArbRollupRollupPruned)
	if err := _ArbRollup.contract.UnpackLog(event, "rollupPruned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbRollupRollupStakeCreatedIterator is returned from FilterRollupStakeCreated and is used to iterate over the raw logs and unpacked data for RollupStakeCreated events raised by the ArbRollup contract.
type ArbRollupRollupStakeCreatedIterator struct {
	Event *ArbRollupRollupStakeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupStakeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupStakeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupStakeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupStakeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupStakeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupStakeCreated represents a RollupStakeCreated event raised by the ArbRollup contract.
type ArbRollupRollupStakeCreated struct {
	Staker      common.Address
	NodeHash    [32]byte
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeCreated is a free log retrieval operation binding the contract event 0x7b3632958ee552a498e60e560ce8ba7141711bc878664ebb8035b9b443d1afe0.
//
// Solidity: event rollupStakeCreated(address staker, bytes32 nodeHash, uint256 blockNumber)
func (_ArbRollup *ArbRollupFilterer) FilterRollupStakeCreated(opts *bind.FilterOpts) (*ArbRollupRollupStakeCreatedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "rollupStakeCreated")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupStakeCreatedIterator{contract: _ArbRollup.contract, event: "rollupStakeCreated", logs: logs, sub: sub}, nil
}

// WatchRollupStakeCreated is a free log subscription operation binding the contract event 0x7b3632958ee552a498e60e560ce8ba7141711bc878664ebb8035b9b443d1afe0.
//
// Solidity: event rollupStakeCreated(address staker, bytes32 nodeHash, uint256 blockNumber)
func (_ArbRollup *ArbRollupFilterer) WatchRollupStakeCreated(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupStakeCreated) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "rollupStakeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupStakeCreated)
				if err := _ArbRollup.contract.UnpackLog(event, "rollupStakeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeCreated is a log parse operation binding the contract event 0x7b3632958ee552a498e60e560ce8ba7141711bc878664ebb8035b9b443d1afe0.
//
// Solidity: event rollupStakeCreated(address staker, bytes32 nodeHash, uint256 blockNumber)
func (_ArbRollup *ArbRollupFilterer) ParseRollupStakeCreated(log types.Log) (*ArbRollupRollupStakeCreated, error) {
	event := new(ArbRollupRollupStakeCreated)
	if err := _ArbRollup.contract.UnpackLog(event, "rollupStakeCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbRollupRollupStakeMovedIterator is returned from FilterRollupStakeMoved and is used to iterate over the raw logs and unpacked data for RollupStakeMoved events raised by the ArbRollup contract.
type ArbRollupRollupStakeMovedIterator struct {
	Event *ArbRollupRollupStakeMoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupStakeMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupStakeMoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupStakeMoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupStakeMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupStakeMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupStakeMoved represents a RollupStakeMoved event raised by the ArbRollup contract.
type ArbRollupRollupStakeMoved struct {
	Staker     common.Address
	ToNodeHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeMoved is a free log retrieval operation binding the contract event 0xe470c1e2f2bfb6e768c29c891291b7ee035ef0fd2105a7c41d5449a11cc98721.
//
// Solidity: event rollupStakeMoved(address staker, bytes32 toNodeHash)
func (_ArbRollup *ArbRollupFilterer) FilterRollupStakeMoved(opts *bind.FilterOpts) (*ArbRollupRollupStakeMovedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "rollupStakeMoved")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupStakeMovedIterator{contract: _ArbRollup.contract, event: "rollupStakeMoved", logs: logs, sub: sub}, nil
}

// WatchRollupStakeMoved is a free log subscription operation binding the contract event 0xe470c1e2f2bfb6e768c29c891291b7ee035ef0fd2105a7c41d5449a11cc98721.
//
// Solidity: event rollupStakeMoved(address staker, bytes32 toNodeHash)
func (_ArbRollup *ArbRollupFilterer) WatchRollupStakeMoved(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupStakeMoved) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "rollupStakeMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupStakeMoved)
				if err := _ArbRollup.contract.UnpackLog(event, "rollupStakeMoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeMoved is a log parse operation binding the contract event 0xe470c1e2f2bfb6e768c29c891291b7ee035ef0fd2105a7c41d5449a11cc98721.
//
// Solidity: event rollupStakeMoved(address staker, bytes32 toNodeHash)
func (_ArbRollup *ArbRollupFilterer) ParseRollupStakeMoved(log types.Log) (*ArbRollupRollupStakeMoved, error) {
	event := new(ArbRollupRollupStakeMoved)
	if err := _ArbRollup.contract.UnpackLog(event, "rollupStakeMoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbRollupRollupStakeRefundedIterator is returned from FilterRollupStakeRefunded and is used to iterate over the raw logs and unpacked data for RollupStakeRefunded events raised by the ArbRollup contract.
type ArbRollupRollupStakeRefundedIterator struct {
	Event *ArbRollupRollupStakeRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupStakeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupStakeRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupStakeRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupStakeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupStakeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupStakeRefunded represents a RollupStakeRefunded event raised by the ArbRollup contract.
type ArbRollupRollupStakeRefunded struct {
	Staker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeRefunded is a free log retrieval operation binding the contract event 0x2a0295967334441364f99dcd4fe5939bc2d0c940802569bc452435d5660956c1.
//
// Solidity: event rollupStakeRefunded(address staker)
func (_ArbRollup *ArbRollupFilterer) FilterRollupStakeRefunded(opts *bind.FilterOpts) (*ArbRollupRollupStakeRefundedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "rollupStakeRefunded")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupStakeRefundedIterator{contract: _ArbRollup.contract, event: "rollupStakeRefunded", logs: logs, sub: sub}, nil
}

// WatchRollupStakeRefunded is a free log subscription operation binding the contract event 0x2a0295967334441364f99dcd4fe5939bc2d0c940802569bc452435d5660956c1.
//
// Solidity: event rollupStakeRefunded(address staker)
func (_ArbRollup *ArbRollupFilterer) WatchRollupStakeRefunded(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupStakeRefunded) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "rollupStakeRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupStakeRefunded)
				if err := _ArbRollup.contract.UnpackLog(event, "rollupStakeRefunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeRefunded is a log parse operation binding the contract event 0x2a0295967334441364f99dcd4fe5939bc2d0c940802569bc452435d5660956c1.
//
// Solidity: event rollupStakeRefunded(address staker)
func (_ArbRollup *ArbRollupFilterer) ParseRollupStakeRefunded(log types.Log) (*ArbRollupRollupStakeRefunded, error) {
	event := new(ArbRollupRollupStakeRefunded)
	if err := _ArbRollup.contract.UnpackLog(event, "rollupStakeRefunded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209209a93bce785e91cc5cbff2f58b8586736753177e41f1b5a490a5098a55b3ae64736f6c634300050e0032"

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

// ChallengeLauncherABI is the input ABI used to generate the binding from.
const ChallengeLauncherABI = "[]"

// ChallengeLauncherFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeLauncherFuncSigs = map[string]string{
	"eda77a6f": "startExecutionChallenge(address,address,bytes32)",
	"b69a1367": "startInvalidMessagesChallenge(address,address,bytes32)",
	"e1e73f8a": "startInvalidPendingTopChallenge(address,address,bytes32)",
}

// ChallengeLauncherBin is the compiled bytecode used for deploying new contracts.
var ChallengeLauncherBin = "0x60e5610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060475760003560e01c8063b69a136714604c578063e1e73f8a14604c578063eda77a6f14604c575b600080fd5b818015605757600080fd5b50608b60048036036060811015606c57600080fd5b506001600160a01b0381358116916020810135909116906040013560a7565b604080516001600160a01b039092168252519081900360200190f35b6000939250505056fea265627a7a723158201130380ca466c14def00a0522472b14dcb06e4792ac2bb5607c7380db0f17b9864736f6c634300050e0032"

// DeployChallengeLauncher deploys a new Ethereum contract, binding an instance of ChallengeLauncher to it.
func DeployChallengeLauncher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeLauncher, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeLauncherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeLauncherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeLauncher{ChallengeLauncherCaller: ChallengeLauncherCaller{contract: contract}, ChallengeLauncherTransactor: ChallengeLauncherTransactor{contract: contract}, ChallengeLauncherFilterer: ChallengeLauncherFilterer{contract: contract}}, nil
}

// ChallengeLauncher is an auto generated Go binding around an Ethereum contract.
type ChallengeLauncher struct {
	ChallengeLauncherCaller     // Read-only binding to the contract
	ChallengeLauncherTransactor // Write-only binding to the contract
	ChallengeLauncherFilterer   // Log filterer for contract events
}

// ChallengeLauncherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeLauncherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLauncherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeLauncherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLauncherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeLauncherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLauncherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeLauncherSession struct {
	Contract     *ChallengeLauncher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChallengeLauncherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeLauncherCallerSession struct {
	Contract *ChallengeLauncherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ChallengeLauncherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeLauncherTransactorSession struct {
	Contract     *ChallengeLauncherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ChallengeLauncherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeLauncherRaw struct {
	Contract *ChallengeLauncher // Generic contract binding to access the raw methods on
}

// ChallengeLauncherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeLauncherCallerRaw struct {
	Contract *ChallengeLauncherCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeLauncherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeLauncherTransactorRaw struct {
	Contract *ChallengeLauncherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeLauncher creates a new instance of ChallengeLauncher, bound to a specific deployed contract.
func NewChallengeLauncher(address common.Address, backend bind.ContractBackend) (*ChallengeLauncher, error) {
	contract, err := bindChallengeLauncher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeLauncher{ChallengeLauncherCaller: ChallengeLauncherCaller{contract: contract}, ChallengeLauncherTransactor: ChallengeLauncherTransactor{contract: contract}, ChallengeLauncherFilterer: ChallengeLauncherFilterer{contract: contract}}, nil
}

// NewChallengeLauncherCaller creates a new read-only instance of ChallengeLauncher, bound to a specific deployed contract.
func NewChallengeLauncherCaller(address common.Address, caller bind.ContractCaller) (*ChallengeLauncherCaller, error) {
	contract, err := bindChallengeLauncher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeLauncherCaller{contract: contract}, nil
}

// NewChallengeLauncherTransactor creates a new write-only instance of ChallengeLauncher, bound to a specific deployed contract.
func NewChallengeLauncherTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeLauncherTransactor, error) {
	contract, err := bindChallengeLauncher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeLauncherTransactor{contract: contract}, nil
}

// NewChallengeLauncherFilterer creates a new log filterer instance of ChallengeLauncher, bound to a specific deployed contract.
func NewChallengeLauncherFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeLauncherFilterer, error) {
	contract, err := bindChallengeLauncher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeLauncherFilterer{contract: contract}, nil
}

// bindChallengeLauncher binds a generic wrapper to an already deployed contract.
func bindChallengeLauncher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeLauncherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeLauncher *ChallengeLauncherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeLauncher.Contract.ChallengeLauncherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeLauncher *ChallengeLauncherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeLauncher.Contract.ChallengeLauncherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeLauncher *ChallengeLauncherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeLauncher.Contract.ChallengeLauncherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeLauncher *ChallengeLauncherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeLauncher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeLauncher *ChallengeLauncherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeLauncher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeLauncher *ChallengeLauncherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeLauncher.Contract.contract.Transact(opts, method, params...)
}

// IArbBaseABI is the input ABI used to generate the binding from.
const IArbBaseABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbBaseFuncSigs maps the 4-byte function signature to its string representation.
var IArbBaseFuncSigs = map[string]string{
	"22c091bc": "completeChallenge(address[2],uint128[2])",
}

// IArbBase is an auto generated Go binding around an Ethereum contract.
type IArbBase struct {
	IArbBaseCaller     // Read-only binding to the contract
	IArbBaseTransactor // Write-only binding to the contract
	IArbBaseFilterer   // Log filterer for contract events
}

// IArbBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbBaseSession struct {
	Contract     *IArbBase         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbBaseCallerSession struct {
	Contract *IArbBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IArbBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbBaseTransactorSession struct {
	Contract     *IArbBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IArbBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbBaseRaw struct {
	Contract *IArbBase // Generic contract binding to access the raw methods on
}

// IArbBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbBaseCallerRaw struct {
	Contract *IArbBaseCaller // Generic read-only contract binding to access the raw methods on
}

// IArbBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbBaseTransactorRaw struct {
	Contract *IArbBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbBase creates a new instance of IArbBase, bound to a specific deployed contract.
func NewIArbBase(address common.Address, backend bind.ContractBackend) (*IArbBase, error) {
	contract, err := bindIArbBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbBase{IArbBaseCaller: IArbBaseCaller{contract: contract}, IArbBaseTransactor: IArbBaseTransactor{contract: contract}, IArbBaseFilterer: IArbBaseFilterer{contract: contract}}, nil
}

// NewIArbBaseCaller creates a new read-only instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseCaller(address common.Address, caller bind.ContractCaller) (*IArbBaseCaller, error) {
	contract, err := bindIArbBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbBaseCaller{contract: contract}, nil
}

// NewIArbBaseTransactor creates a new write-only instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbBaseTransactor, error) {
	contract, err := bindIArbBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbBaseTransactor{contract: contract}, nil
}

// NewIArbBaseFilterer creates a new log filterer instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbBaseFilterer, error) {
	contract, err := bindIArbBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbBaseFilterer{contract: contract}, nil
}

// bindIArbBase binds a generic wrapper to an already deployed contract.
func bindIArbBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbBase *IArbBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbBase.Contract.IArbBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbBase *IArbBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbBase.Contract.IArbBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbBase *IArbBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbBase.Contract.IArbBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbBase *IArbBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbBase *IArbBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbBase *IArbBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbBase.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.Contract.CompleteChallenge(&_IArbBase.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.Contract.CompleteChallenge(&_IArbBase.TransactOpts, _players, _rewards)
}

// IGlobalPendingInboxABI is the input ABI used to generate the binding from.
const IGlobalPendingInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes21\",\"name\":\"tokenType\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pullPendingMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var IGlobalPendingInboxFuncSigs = map[string]string{
	"3bbc3c32": "forwardMessage(address,bytes21,uint256,bytes,bytes)",
	"d106ec19": "pullPendingMessages()",
	"f3972383": "registerForInbox()",
	"3fc6eb80": "sendEthMessage(address,bytes)",
	"626cef85": "sendMessage(address,bytes21,uint256,bytes)",
	"e4eb8c63": "sendMessages(bytes)",
}

// IGlobalPendingInbox is an auto generated Go binding around an Ethereum contract.
type IGlobalPendingInbox struct {
	IGlobalPendingInboxCaller     // Read-only binding to the contract
	IGlobalPendingInboxTransactor // Write-only binding to the contract
	IGlobalPendingInboxFilterer   // Log filterer for contract events
}

// IGlobalPendingInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGlobalPendingInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGlobalPendingInboxSession struct {
	Contract     *IGlobalPendingInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGlobalPendingInboxCallerSession struct {
	Contract *IGlobalPendingInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IGlobalPendingInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGlobalPendingInboxTransactorSession struct {
	Contract     *IGlobalPendingInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGlobalPendingInboxRaw struct {
	Contract *IGlobalPendingInbox // Generic contract binding to access the raw methods on
}

// IGlobalPendingInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCallerRaw struct {
	Contract *IGlobalPendingInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IGlobalPendingInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactorRaw struct {
	Contract *IGlobalPendingInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGlobalPendingInbox creates a new instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInbox(address common.Address, backend bind.ContractBackend) (*IGlobalPendingInbox, error) {
	contract, err := bindIGlobalPendingInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInbox{IGlobalPendingInboxCaller: IGlobalPendingInboxCaller{contract: contract}, IGlobalPendingInboxTransactor: IGlobalPendingInboxTransactor{contract: contract}, IGlobalPendingInboxFilterer: IGlobalPendingInboxFilterer{contract: contract}}, nil
}

// NewIGlobalPendingInboxCaller creates a new read-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxCaller(address common.Address, caller bind.ContractCaller) (*IGlobalPendingInboxCaller, error) {
	contract, err := bindIGlobalPendingInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxCaller{contract: contract}, nil
}

// NewIGlobalPendingInboxTransactor creates a new write-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IGlobalPendingInboxTransactor, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxTransactor{contract: contract}, nil
}

// NewIGlobalPendingInboxFilterer creates a new log filterer instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IGlobalPendingInboxFilterer, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxFilterer{contract: contract}, nil
}

// bindIGlobalPendingInbox binds a generic wrapper to an already deployed contract.
func bindIGlobalPendingInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGlobalPendingInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transact(opts, method, params...)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) ForwardMessage(opts *bind.TransactOpts, _destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "forwardMessage", _destination, _tokenType, _amount, _data, _signature)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) ForwardMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data, _signature)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) ForwardMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data, _signature)
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0xd106ec19.
//
// Solidity: function pullPendingMessages() returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) PullPendingMessages(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "pullPendingMessages")
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0xd106ec19.
//
// Solidity: function pullPendingMessages() returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) PullPendingMessages() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.PullPendingMessages(&_IGlobalPendingInbox.TransactOpts)
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0xd106ec19.
//
// Solidity: function pullPendingMessages() returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) PullPendingMessages() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.PullPendingMessages(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) RegisterForInbox(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "registerForInbox")
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendEthMessage(opts *bind.TransactOpts, _destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendEthMessage", _destination, _data)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendEthMessage(_destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendEthMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _data)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendEthMessage(_destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendEthMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessage(opts *bind.TransactOpts, _destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessage", _destination, _tokenType, _amount, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessages(opts *bind.TransactOpts, _messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessages", _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _messages)
}

// IGlobalPendingInboxMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGlobalPendingInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGlobalPendingInboxMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGlobalPendingInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxMessageDelivered represents a MessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxMessageDelivered struct {
	VmId      common.Address
	Sender    common.Address
	TokenType [21]byte
	Value     *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterMessageDelivered(opts *bind.FilterOpts, vmId []common.Address) (*IGlobalPendingInboxMessageDeliveredIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "MessageDelivered", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxMessageDelivered, vmId []common.Address) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "MessageDelivered", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseMessageDelivered(log types.Log) (*IGlobalPendingInboxMessageDelivered, error) {
	event := new(IGlobalPendingInboxMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ProtocolFuncSigs = map[string]string{
	"1cd765fa": "generateAssertionHash(bytes32,uint32,uint64,bytes32,bytes32,bytes32,bytes32)",
	"e83f4bfe": "generateLastMessageHash(bytes)",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"85ecb92a": "generatePreconditionHash(bytes32,uint64[2],bytes32)",
}

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x610a93610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100555760003560e01c80624c28f61461005a5780631cd765fa146100b257806385ecb92a14610103578063e83f4bfe14610158575b600080fd5b6100a06004803603608081101561007057600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b03166101fe565b60408051918252519081900360200190f35b6100a0600480360360e08110156100c857600080fd5b5080359063ffffffff6020820135169067ffffffffffffffff6040820135169060608101359060808101359060a08101359060c001356102f0565b6100a06004803603608081101561011957600080fd5b604080518082018252833593928301929160608301919060208401906002908390839080828437600092019190915250919450509035915061035d9050565b6100a06004803603602081101561016e57600080fd5b81019060208101813564010000000081111561018957600080fd5b82018360208201111561019b57600080fd5b803590602001918460018302840111640100000000831117156101bd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103b1945050505050565b60408051600480825260a0820190925260009160609190816020015b6102226109f7565b81526020019060019003908161021a579050509050610240866104f6565b8160008151811061024d57fe5b602002602001018190525061026a836001600160a01b0316610574565b8160018151811061027757fe5b602002602001018190525061028b84610574565b8160028151811061029857fe5b60209081029190910101526102ba6affffffffffffffffffffff198616610574565b816003815181106102c757fe5b60200260200101819052506102e36102de826105f2565b6106a2565b519150505b949350505050565b6040805160208082019990995260e09790971b6001600160e01b0319168782015260c09590951b6001600160c01b0319166044870152604c860193909352606c850191909152608c84015260ac808401919091528151808403909101815260cc9092019052805191012090565b815160209283015160408051808601969096526001600160c01b031960c093841b8116878301529190921b166048850152605080850192909252805180850390920182526070909301909252815191012090565b8051600090819081908190815b818110156104e95773__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63d36cfac288866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561043457818101518382015260200161041c565b50505050905090810190601f1680156104615780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561047e57600080fd5b505af4158015610492573d6000803e3d6000fd5b505050506040513d60408110156104a857600080fd5b50805160209182015160408051808501999099528881018290528051808a0382018152606090990190528751979092019690962095945092506001016103be565b509293505050505b919050565b6104fe6109f7565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610563565b6105506109f7565b8152602001906001900390816105485790505b508152600260209091015292915050565b61057c6109f7565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916105e1565b6105ce6109f7565b8152602001906001900390816105c65790505b508152600060209091015292915050565b6105fa6109f7565b61060482516107d8565b610655576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6106aa610a25565b6060820151600c60ff909116106106fc576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661072957604051806020016040528061072084600001516107df565b905290506104f1565b606082015160ff1660011415610770576040518060200160405280610720846020015160000151856020015160400151866020015160600151876020015160200151610803565b606082015160ff166002141561079557506040805160208101909152815181526104f1565b600360ff16826060015160ff16101580156107b957506060820151600c60ff909116105b156107d657604051806020016040528061072084604001516108ab565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561085d575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206102e8565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b60006008825111156108fb576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610928578160200160208202803883390190505b50805190915060005b8181101561098457610941610a25565b61095d86838151811061095057fe5b60200260200101516106a2565b9050806000015184838151811061097057fe5b602090810291909101015250600101610931565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156109cd5781810151838201526020016109b5565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b604051806080016040528060008152602001610a11610a37565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a7231582055bab63760b108ada8f06a691f624211fa2775d8f302f8f5ee3e6c4924b153d664736f6c634300050e0032"

// DeployProtocol deploys a new Ethereum contract, binding an instance of Protocol to it.
func DeployProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Protocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	valueAddr, _, _, _ := DeployValue(auth, backend)
	ProtocolBin = strings.Replace(ProtocolBin, "__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__", valueAddr.String()[2:], -1)

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

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1cd765fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _numGas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _numGas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1cd765fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _numGas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateAssertionHash(&_Protocol.CallOpts, _afterHash, _numSteps, _numGas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1cd765fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, uint64 _numGas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _numGas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateAssertionHash(&_Protocol.CallOpts, _afterHash, _numSteps, _numGas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _messages []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateLastMessageHash", _messages)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateLastMessageHash(&_Protocol.CallOpts, _messages)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateLastMessageHash(&_Protocol.CallOpts, _messages)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _Protocol.Contract.GenerateMessageStubHash(&_Protocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _Protocol.Contract.GenerateMessageStubHash(&_Protocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GeneratePreconditionHash(&_Protocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GeneratePreconditionHash(&_Protocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158206e634b19b85c6f44f6988b30e2f365b9cee8b2cacb101081a262451b4b05483764736f6c634300050e0032"

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
const VMABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmStateHash\",\"type\":\"bytes32\"}],\"name\":\"isErrored\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmStateHash\",\"type\":\"bytes32\"}],\"name\":\"isHalted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inboxHash\",\"type\":\"bytes32\"}],\"name\":\"protoStateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// VMFuncSigs maps the 4-byte function signature to its string representation.
var VMFuncSigs = map[string]string{
	"6aee3ecc": "isErrored(bytes32)",
	"f9e2b912": "isHalted(bytes32)",
	"8348abc9": "protoStateHash(bytes32,bytes32)",
}

// VMBin is the compiled bytecode used for deploying new contracts.
var VMBin = "0x610131610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060475760003560e01c80636aee3ecc14604c5780638348abc914607a578063f9e2b9121460ac575b600080fd5b606660048036036020811015606057600080fd5b503560c6565b604080519115158252519081900360200190f35b609a60048036036040811015608e57600080fd5b508035906020013560cc565b60408051918252519081900360200190f35b60666004803603602081101560c057600080fd5b503560f8565b60011490565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b159056fea265627a7a723158208a05ec03547d1432c13322ac70d4b0b2b854126d0065eceac23cbb199d42e35164736f6c634300050e0032"

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

// IsErrored is a free data retrieval call binding the contract method 0x6aee3ecc.
//
// Solidity: function isErrored(bytes32 vmStateHash) constant returns(bool)
func (_VM *VMCaller) IsErrored(opts *bind.CallOpts, vmStateHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VM.contract.Call(opts, out, "isErrored", vmStateHash)
	return *ret0, err
}

// IsErrored is a free data retrieval call binding the contract method 0x6aee3ecc.
//
// Solidity: function isErrored(bytes32 vmStateHash) constant returns(bool)
func (_VM *VMSession) IsErrored(vmStateHash [32]byte) (bool, error) {
	return _VM.Contract.IsErrored(&_VM.CallOpts, vmStateHash)
}

// IsErrored is a free data retrieval call binding the contract method 0x6aee3ecc.
//
// Solidity: function isErrored(bytes32 vmStateHash) constant returns(bool)
func (_VM *VMCallerSession) IsErrored(vmStateHash [32]byte) (bool, error) {
	return _VM.Contract.IsErrored(&_VM.CallOpts, vmStateHash)
}

// IsHalted is a free data retrieval call binding the contract method 0xf9e2b912.
//
// Solidity: function isHalted(bytes32 vmStateHash) constant returns(bool)
func (_VM *VMCaller) IsHalted(opts *bind.CallOpts, vmStateHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VM.contract.Call(opts, out, "isHalted", vmStateHash)
	return *ret0, err
}

// IsHalted is a free data retrieval call binding the contract method 0xf9e2b912.
//
// Solidity: function isHalted(bytes32 vmStateHash) constant returns(bool)
func (_VM *VMSession) IsHalted(vmStateHash [32]byte) (bool, error) {
	return _VM.Contract.IsHalted(&_VM.CallOpts, vmStateHash)
}

// IsHalted is a free data retrieval call binding the contract method 0xf9e2b912.
//
// Solidity: function isHalted(bytes32 vmStateHash) constant returns(bool)
func (_VM *VMCallerSession) IsHalted(vmStateHash [32]byte) (bool, error) {
	return _VM.Contract.IsHalted(&_VM.CallOpts, vmStateHash)
}

// ProtoStateHash is a free data retrieval call binding the contract method 0x8348abc9.
//
// Solidity: function protoStateHash(bytes32 machineHash, bytes32 inboxHash) constant returns(bytes32)
func (_VM *VMCaller) ProtoStateHash(opts *bind.CallOpts, machineHash [32]byte, inboxHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _VM.contract.Call(opts, out, "protoStateHash", machineHash, inboxHash)
	return *ret0, err
}

// ProtoStateHash is a free data retrieval call binding the contract method 0x8348abc9.
//
// Solidity: function protoStateHash(bytes32 machineHash, bytes32 inboxHash) constant returns(bytes32)
func (_VM *VMSession) ProtoStateHash(machineHash [32]byte, inboxHash [32]byte) ([32]byte, error) {
	return _VM.Contract.ProtoStateHash(&_VM.CallOpts, machineHash, inboxHash)
}

// ProtoStateHash is a free data retrieval call binding the contract method 0x8348abc9.
//
// Solidity: function protoStateHash(bytes32 machineHash, bytes32 inboxHash) constant returns(bytes32)
func (_VM *VMCallerSession) ProtoStateHash(machineHash [32]byte, inboxHash [32]byte) ([32]byte, error) {
	return _VM.Contract.ProtoStateHash(&_VM.CallOpts, machineHash, inboxHash)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeHashed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidHashed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"immediate\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasic\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashInt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueFuncSigs maps the 4-byte function signature to its string representation.
var ValueFuncSigs = map[string]string{
	"3d730ed2": "deserializeHashed(bytes)",
	"32e6cc21": "deserializeMessage(bytes,uint256)",
	"d36cfac2": "deserializeValidHashed(bytes,uint256)",
	"72403aa0": "getNextValid(bytes,uint256)",
	"826513e0": "hashCodePoint(uint8,bool,bytes32,bytes32)",
	"b697e085": "hashCodePointBasic(uint8,bytes32)",
	"3c786053": "hashCodePointImmediate(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"5043dff1": "hashInt(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x61152b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c806372403aa01161007057806372403aa014610300578063826513e014610425578063b2b9dc6214610459578063b697e0851461048a578063d36cfac2146104b0576100a8565b806332e6cc21146100ad578063364df277146101f95780633c786053146102135780633d730ed21461023f5780635043dff1146102e3575b600080fd5b610153600480360360408110156100c357600080fd5b810190602081018135600160201b8111156100dd57600080fd5b8201836020820111156100ef57600080fd5b803590602001918460018302840111600160201b8311171561011057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061056f915050565b604051808815151515815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101b85781810151838201526020016101a0565b50505050905090810190601f1680156101e55780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b61020161076d565b60408051918252519081900360200190f35b6102016004803603606081101561022957600080fd5b5060ff81351690602081013590604001356107e0565b6102016004803603602081101561025557600080fd5b810190602081018135600160201b81111561026f57600080fd5b82018360208201111561028157600080fd5b803590602001918460018302840111600160201b831117156102a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610832945050505050565b610201600480360360208110156102f957600080fd5b50356108a6565b6103a66004803603604081101561031657600080fd5b810190602081018135600160201b81111561033057600080fd5b82018360208201111561034257600080fd5b803590602001918460018302840111600160201b8311171561036357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506108ca915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103e95781810151838201526020016103d1565b50505050905090810190601f1680156104165780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6102016004803603608081101561043b57600080fd5b5060ff8135169060208101351515906040810135906060013561094e565b6104766004803603602081101561046f57600080fd5b50356109f7565b604080519115158252519081900360200190f35b610201600480360360408110156104a057600080fd5b5060ff81351690602001356109fe565b610556600480360360408110156104c657600080fd5b810190602081018135600160201b8111156104e057600080fd5b8201836020820111156104f257600080fd5b803590602001918460018302840111600160201b8311171561051357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610a45915050565b6040805192835260208301919091528051918290030190f35b6000806000806000806060600088965060008a888151811061058d57fe5b016020015160019098019760f81c9050600781146105bf576105b28b60018a03610a45565b9098509650610761915050565b6105c98b89610a45565b90985091506105e88b60018c016000198d8c030163ffffffff610abe16565b92508a88815181106105f657fe5b016020015160019098019760f81c90508015610619576105b28b60018a03610a45565b6106238b89610b3e565b80995081975050508a888151811061063757fe5b016020015160019098019760f81c9050801561065a576105b28b60018a03610a45565b6106648b89610b3e565b80995081965050508a888151811061067857fe5b016020015160019098019760f81c9050801561069b576105b28b60018a03610a45565b6106a58b89610b3e565b60408051600480825260a0820190925260019c50919a509195506060916020820160808038833901905050905082816000815181106106e057fe5b6020026020010181815250506106f5876108a6565b8160018151811061070257fe5b602002602001018181525050610717866108a6565b8160028151811061072457fe5b602002602001018181525050610739856108a6565b8160038151811061074657fe5b60200260200101818152505061075b81610b65565b97505050505b92959891949750929550565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156107b95781810151838201526020016107a1565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6000808061083e61146f565b610849856000610c25565b919450925090508215610891576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61089a81610daf565b5193505050505b919050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060606000806108d961146f565b6108e38787610c25565b91945092509050821561092b576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b8161093f888880840363ffffffff610abe16565b945094505050505b9250929050565b600083156109a8575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206109ef565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b6008101590565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b600080600080610a5361146f565b610a5d8787610c25565b919450925090508215610aa5576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b81610aaf82610daf565b51909890975095505050505050565b606081830184511015610ad057600080fd5b606082158015610aeb57604051915060208201604052610b35565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610b24578051835260209283019201610b0c565b5050858452601f01601f1916604052505b50949350505050565b6000808281610b53868363ffffffff610ee516565b60209290920196919550909350505050565b6000600882511115610bb5576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b8151600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610bfd578181015183820152602001610be5565b5050505090500192505050604051602081830303815290604052805190602001209050919050565b600080610c3061146f565b84518410610c85576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110610c9857fe5b016020015160019092019160f81c90506000610cb261149d565b60ff8316610ce657610cc48985610b3e565b9094509150600084610cd584610f01565b91985096509450610da89350505050565b60ff831660011415610d0d57610cfc8985610f7f565b9094509050600084610cd5836110da565b60ff831660021415610d3457610d238985610b3e565b9094509150600084610cd58461113a565b600360ff841610801590610d4b5750600c60ff8416105b15610d8857600219830160606000610d64838d896111b8565b909850925090508087610d7684611273565b99509950995050505050505050610da8565b8260ff16612710016000610d9c6000610f01565b91985096509450505050505b9250925092565b610db76114c4565b6060820151600c60ff90911610610e09576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610e36576040518060200160405280610e2d84600001516108a6565b905290506108a1565b606082015160ff1660011415610e7d576040518060200160405280610e2d84602001516000015185602001516040015186602001516060015187602001516020015161094e565b606082015160ff1660021415610ea257506040805160208101909152815181526108a1565b600360ff16826060015160ff1610158015610ec657506060820151600c60ff909116105b15610ee3576040518060200160405280610e2d8460400151611323565bfe5b60008160200183511015610ef857600080fd5b50016020015190565b610f0961146f565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610f6e565b610f5b61146f565b815260200190600190039081610f535790505b508152600060209091015292915050565b6000610f8961149d565b60008390506000858281518110610f9c57fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110610fc257fe5b016020015160019384019360f89190911c915060009060ff8416141561104e576000610fec61146f565b610ff68a87610c25565b9097509092509050811561103f576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61104881610daf565b51925050505b6000611060898663ffffffff610ee516565b90506020850194508360ff16600114156110a5576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506109479050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6110e261146f565b604080516080810182526000808252602080830186905283518281529081018452919283019190611129565b61111661146f565b81526020019060019003908161110e5790505b508152600160209091015292915050565b61114261146f565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916111a7565b61119461146f565b81526020019060019003908161118c5790505b508152600260209091015292915050565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561120357816020015b6111f061146f565b8152602001906001900390816111e85790505b50905060005b8960ff168160ff16101561125d576112218985610c25565b8451859060ff861690811061123257fe5b60209081029190910101529450925082156112555750909450909250905061126a565b600101611209565b5060009550919350909150505b93509350939050565b61127b61146f565b61128582516109f7565b6112d6576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6000600882511115611373576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156113a0578160200160208202803883390190505b50805190915060005b818110156113fc576113b96114c4565b6113d58683815181106113c857fe5b6020026020010151610daf565b905080600001518483815181106113e857fe5b6020908102919091010152506001016113a9565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561144557818101518382015260200161142d565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180608001604052806000815260200161148961149d565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4d61727368616c6c65642076616c7565206d7573742062652076616c69640000a265627a7a7231582002c6ecc6d323338594373e09189f176b7c79457073cee7bd7cd2ed946f380fa664736f6c634300050e0032"

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

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueCaller) DeserializeHashed(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "deserializeHashed", data)
	return *ret0, err
}

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueSession) DeserializeHashed(data []byte) ([32]byte, error) {
	return _Value.Contract.DeserializeHashed(&_Value.CallOpts, data)
}

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueCallerSession) DeserializeHashed(data []byte) ([32]byte, error) {
	return _Value.Contract.DeserializeHashed(&_Value.CallOpts, data)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueCaller) DeserializeMessage(opts *bind.CallOpts, data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	ret := new(struct {
		Valid       bool
		Offset      *big.Int
		MessageHash [32]byte
		Destination *big.Int
		Value       *big.Int
		TokenType   *big.Int
		MessageData []byte
	})
	out := ret
	err := _Value.contract.Call(opts, out, "deserializeMessage", data, startOffset)
	return *ret, err
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _Value.Contract.DeserializeMessage(&_Value.CallOpts, data, startOffset)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueCallerSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _Value.Contract.DeserializeMessage(&_Value.CallOpts, data, startOffset)
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueCaller) DeserializeValidHashed(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Value.contract.Call(opts, out, "deserializeValidHashed", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueSession) DeserializeValidHashed(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _Value.Contract.DeserializeValidHashed(&_Value.CallOpts, data, offset)
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueCallerSession) DeserializeValidHashed(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _Value.Contract.DeserializeValidHashed(&_Value.CallOpts, data, offset)
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueCaller) GetNextValid(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Value.contract.Call(opts, out, "getNextValid", data, offset)
	return *ret0, *ret1, err
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueSession) GetNextValid(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _Value.Contract.GetNextValid(&_Value.CallOpts, data, offset)
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueCallerSession) GetNextValid(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _Value.Contract.GetNextValid(&_Value.CallOpts, data, offset)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePoint(opts *bind.CallOpts, opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePoint", opcode, immediate, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePoint(&_Value.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePoint(&_Value.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePointBasic(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePointBasic", opcode, nextCodePoint)
	return *ret0, err
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePointBasic(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointBasic(&_Value.CallOpts, opcode, nextCodePoint)
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePointBasic(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointBasic(&_Value.CallOpts, opcode, nextCodePoint)
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePointImmediate(opts *bind.CallOpts, opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePointImmediate", opcode, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePointImmediate(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointImmediate(&_Value.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePointImmediate(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointImmediate(&_Value.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueSession) HashEmptyTuple() ([32]byte, error) {
	return _Value.Contract.HashEmptyTuple(&_Value.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _Value.Contract.HashEmptyTuple(&_Value.CallOpts)
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueCaller) HashInt(opts *bind.CallOpts, val *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashInt", val)
	return *ret0, err
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueSession) HashInt(val *big.Int) ([32]byte, error) {
	return _Value.Contract.HashInt(&_Value.CallOpts, val)
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueCallerSession) HashInt(val *big.Int) ([32]byte, error) {
	return _Value.Contract.HashInt(&_Value.CallOpts, val)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueCaller) IsValidTupleSize(opts *bind.CallOpts, size *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "isValidTupleSize", size)
	return *ret0, err
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _Value.Contract.IsValidTupleSize(&_Value.CallOpts, size)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueCallerSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _Value.Contract.IsValidTupleSize(&_Value.CallOpts, size)
}
