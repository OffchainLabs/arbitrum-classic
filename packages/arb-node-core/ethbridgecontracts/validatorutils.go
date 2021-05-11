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

// AbsRollupABI is the input ABI used to generate the binding from.
const AbsRollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterSendAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"NodeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentNodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterInboxBatchEndCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterInboxBatchAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"indexed\":false,\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"NodeRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startNode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endNode\",\"type\":\"uint256\"}],\"name\":\"NodesDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNode\",\"type\":\"uint256\"}],\"name\":\"StakerReassigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeSendAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"executionHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proposedTimes\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"maxMessageCounts\",\"type\":\"uint256[2]\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedBridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraChallengeTimeBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[7]\",\"name\":\"connectedContracts\",\"type\":\"address[7]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"requireUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireUnresolvedExists\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstUnresolvedNodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestNodeCreated\",\"type\":\"uint256\"}],\"name\":\"requiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerBridge\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expectedNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"},{\"internalType\":\"uint256\",\"name\":\"beforeProposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxMaxCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sequencerBatchProof\",\"type\":\"bytes\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"}],\"name\":\"upgradeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeImplementationAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbsRollup is an auto generated Go binding around an Ethereum contract.
type AbsRollup struct {
	AbsRollupCaller     // Read-only binding to the contract
	AbsRollupTransactor // Write-only binding to the contract
	AbsRollupFilterer   // Log filterer for contract events
}

// AbsRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsRollupSession struct {
	Contract     *AbsRollup        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsRollupCallerSession struct {
	Contract *AbsRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AbsRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsRollupTransactorSession struct {
	Contract     *AbsRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AbsRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsRollupRaw struct {
	Contract *AbsRollup // Generic contract binding to access the raw methods on
}

// AbsRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsRollupCallerRaw struct {
	Contract *AbsRollupCaller // Generic read-only contract binding to access the raw methods on
}

// AbsRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsRollupTransactorRaw struct {
	Contract *AbsRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsRollup creates a new instance of AbsRollup, bound to a specific deployed contract.
func NewAbsRollup(address common.Address, backend bind.ContractBackend) (*AbsRollup, error) {
	contract, err := bindAbsRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbsRollup{AbsRollupCaller: AbsRollupCaller{contract: contract}, AbsRollupTransactor: AbsRollupTransactor{contract: contract}, AbsRollupFilterer: AbsRollupFilterer{contract: contract}}, nil
}

// NewAbsRollupCaller creates a new read-only instance of AbsRollup, bound to a specific deployed contract.
func NewAbsRollupCaller(address common.Address, caller bind.ContractCaller) (*AbsRollupCaller, error) {
	contract, err := bindAbsRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsRollupCaller{contract: contract}, nil
}

// NewAbsRollupTransactor creates a new write-only instance of AbsRollup, bound to a specific deployed contract.
func NewAbsRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsRollupTransactor, error) {
	contract, err := bindAbsRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsRollupTransactor{contract: contract}, nil
}

// NewAbsRollupFilterer creates a new log filterer instance of AbsRollup, bound to a specific deployed contract.
func NewAbsRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsRollupFilterer, error) {
	contract, err := bindAbsRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsRollupFilterer{contract: contract}, nil
}

// bindAbsRollup binds a generic wrapper to an already deployed contract.
func bindAbsRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsRollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbsRollup *AbsRollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbsRollup.Contract.AbsRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbsRollup *AbsRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsRollup.Contract.AbsRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbsRollup *AbsRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbsRollup.Contract.AbsRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbsRollup *AbsRollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbsRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbsRollup *AbsRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbsRollup *AbsRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbsRollup.Contract.contract.Transact(opts, method, params...)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_AbsRollup *AbsRollupCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "_stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_AbsRollup *AbsRollupSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _AbsRollup.Contract.StakerMap(&_AbsRollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_AbsRollup *AbsRollupCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _AbsRollup.Contract.StakerMap(&_AbsRollup.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AbsRollup *AbsRollupCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AbsRollup *AbsRollupSession) Admin() (common.Address, error) {
	return _AbsRollup.Contract.Admin(&_AbsRollup.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) Admin() (common.Address, error) {
	return _AbsRollup.Contract.Admin(&_AbsRollup.CallOpts)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.AmountStaked(&_AbsRollup.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.AmountStaked(&_AbsRollup.CallOpts, staker)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ArbGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "arbGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_AbsRollup *AbsRollupSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _AbsRollup.Contract.ArbGasSpeedLimitPerBlock(&_AbsRollup.CallOpts)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _AbsRollup.Contract.ArbGasSpeedLimitPerBlock(&_AbsRollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) BaseStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "baseStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_AbsRollup *AbsRollupSession) BaseStake() (*big.Int, error) {
	return _AbsRollup.Contract.BaseStake(&_AbsRollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) BaseStake() (*big.Int, error) {
	return _AbsRollup.Contract.BaseStake(&_AbsRollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_AbsRollup *AbsRollupCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_AbsRollup *AbsRollupSession) ChallengeFactory() (common.Address, error) {
	return _AbsRollup.Contract.ChallengeFactory(&_AbsRollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) ChallengeFactory() (common.Address, error) {
	return _AbsRollup.Contract.ChallengeFactory(&_AbsRollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ConfirmPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "confirmPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _AbsRollup.Contract.ConfirmPeriodBlocks(&_AbsRollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _AbsRollup.Contract.ConfirmPeriodBlocks(&_AbsRollup.CallOpts)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) CountStakedZombies(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "countStakedZombies", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_AbsRollup *AbsRollupSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.CountStakedZombies(&_AbsRollup.CallOpts, node)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.CountStakedZombies(&_AbsRollup.CallOpts, node)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_AbsRollup *AbsRollupCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_AbsRollup *AbsRollupSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _AbsRollup.Contract.CurrentChallenge(&_AbsRollup.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_AbsRollup *AbsRollupCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _AbsRollup.Contract.CurrentChallenge(&_AbsRollup.CallOpts, staker)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_AbsRollup *AbsRollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _AbsRollup.Contract.CurrentRequiredStake(&_AbsRollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _AbsRollup.Contract.CurrentRequiredStake(&_AbsRollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_AbsRollup *AbsRollupCaller) DelayedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "delayedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_AbsRollup *AbsRollupSession) DelayedBridge() (common.Address, error) {
	return _AbsRollup.Contract.DelayedBridge(&_AbsRollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) DelayedBridge() (common.Address, error) {
	return _AbsRollup.Contract.DelayedBridge(&_AbsRollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ExtraChallengeTimeBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "extraChallengeTimeBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _AbsRollup.Contract.ExtraChallengeTimeBlocks(&_AbsRollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _AbsRollup.Contract.ExtraChallengeTimeBlocks(&_AbsRollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_AbsRollup *AbsRollupSession) FirstUnresolvedNode() (*big.Int, error) {
	return _AbsRollup.Contract.FirstUnresolvedNode(&_AbsRollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _AbsRollup.Contract.FirstUnresolvedNode(&_AbsRollup.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_AbsRollup *AbsRollupCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_AbsRollup *AbsRollupSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.GetNode(&_AbsRollup.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_AbsRollup *AbsRollupCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.GetNode(&_AbsRollup.CallOpts, nodeNum)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_AbsRollup *AbsRollupCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_AbsRollup *AbsRollupSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _AbsRollup.Contract.GetNodeHash(&_AbsRollup.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_AbsRollup *AbsRollupCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _AbsRollup.Contract.GetNodeHash(&_AbsRollup.CallOpts, index)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_AbsRollup *AbsRollupCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_AbsRollup *AbsRollupSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.GetStakerAddress(&_AbsRollup.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_AbsRollup *AbsRollupCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.GetStakerAddress(&_AbsRollup.CallOpts, stakerNum)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_AbsRollup *AbsRollupCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_AbsRollup *AbsRollupSession) IsMaster() (bool, error) {
	return _AbsRollup.Contract.IsMaster(&_AbsRollup.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_AbsRollup *AbsRollupCallerSession) IsMaster() (bool, error) {
	return _AbsRollup.Contract.IsMaster(&_AbsRollup.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_AbsRollup *AbsRollupCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_AbsRollup *AbsRollupSession) IsStaked(staker common.Address) (bool, error) {
	return _AbsRollup.Contract.IsStaked(&_AbsRollup.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_AbsRollup *AbsRollupCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _AbsRollup.Contract.IsStaked(&_AbsRollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_AbsRollup *AbsRollupCaller) IsZombie(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "isZombie", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_AbsRollup *AbsRollupSession) IsZombie(staker common.Address) (bool, error) {
	return _AbsRollup.Contract.IsZombie(&_AbsRollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_AbsRollup *AbsRollupCallerSession) IsZombie(staker common.Address) (bool, error) {
	return _AbsRollup.Contract.IsZombie(&_AbsRollup.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_AbsRollup *AbsRollupSession) LastStakeBlock() (*big.Int, error) {
	return _AbsRollup.Contract.LastStakeBlock(&_AbsRollup.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) LastStakeBlock() (*big.Int, error) {
	return _AbsRollup.Contract.LastStakeBlock(&_AbsRollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_AbsRollup *AbsRollupSession) LatestConfirmed() (*big.Int, error) {
	return _AbsRollup.Contract.LatestConfirmed(&_AbsRollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) LatestConfirmed() (*big.Int, error) {
	return _AbsRollup.Contract.LatestConfirmed(&_AbsRollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_AbsRollup *AbsRollupSession) LatestNodeCreated() (*big.Int, error) {
	return _AbsRollup.Contract.LatestNodeCreated(&_AbsRollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _AbsRollup.Contract.LatestNodeCreated(&_AbsRollup.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.LatestStakedNode(&_AbsRollup.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.LatestStakedNode(&_AbsRollup.CallOpts, staker)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_AbsRollup *AbsRollupSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _AbsRollup.Contract.MinimumAssertionPeriod(&_AbsRollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _AbsRollup.Contract.MinimumAssertionPeriod(&_AbsRollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_AbsRollup *AbsRollupCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_AbsRollup *AbsRollupSession) NodeFactory() (common.Address, error) {
	return _AbsRollup.Contract.NodeFactory(&_AbsRollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) NodeFactory() (common.Address, error) {
	return _AbsRollup.Contract.NodeFactory(&_AbsRollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_AbsRollup *AbsRollupCaller) Outbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "outbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_AbsRollup *AbsRollupSession) Outbox() (common.Address, error) {
	return _AbsRollup.Contract.Outbox(&_AbsRollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) Outbox() (common.Address, error) {
	return _AbsRollup.Contract.Outbox(&_AbsRollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbsRollup *AbsRollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbsRollup *AbsRollupSession) Owner() (common.Address, error) {
	return _AbsRollup.Contract.Owner(&_AbsRollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) Owner() (common.Address, error) {
	return _AbsRollup.Contract.Owner(&_AbsRollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AbsRollup *AbsRollupCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AbsRollup *AbsRollupSession) Paused() (bool, error) {
	return _AbsRollup.Contract.Paused(&_AbsRollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AbsRollup *AbsRollupCallerSession) Paused() (bool, error) {
	return _AbsRollup.Contract.Paused(&_AbsRollup.CallOpts)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_AbsRollup *AbsRollupCaller) RequireUnresolved(opts *bind.CallOpts, nodeNum *big.Int) error {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "requireUnresolved", nodeNum)

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_AbsRollup *AbsRollupSession) RequireUnresolved(nodeNum *big.Int) error {
	return _AbsRollup.Contract.RequireUnresolved(&_AbsRollup.CallOpts, nodeNum)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_AbsRollup *AbsRollupCallerSession) RequireUnresolved(nodeNum *big.Int) error {
	return _AbsRollup.Contract.RequireUnresolved(&_AbsRollup.CallOpts, nodeNum)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_AbsRollup *AbsRollupCaller) RequireUnresolvedExists(opts *bind.CallOpts) error {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "requireUnresolvedExists")

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_AbsRollup *AbsRollupSession) RequireUnresolvedExists() error {
	return _AbsRollup.Contract.RequireUnresolvedExists(&_AbsRollup.CallOpts)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_AbsRollup *AbsRollupCallerSession) RequireUnresolvedExists() error {
	return _AbsRollup.Contract.RequireUnresolvedExists(&_AbsRollup.CallOpts)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) RequiredStake(opts *bind.CallOpts, blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "requiredStake", blockNumber, firstUnresolvedNodeNum, latestNodeCreated)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_AbsRollup *AbsRollupSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _AbsRollup.Contract.RequiredStake(&_AbsRollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _AbsRollup.Contract.RequiredStake(&_AbsRollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_AbsRollup *AbsRollupCaller) RollupEventBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "rollupEventBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_AbsRollup *AbsRollupSession) RollupEventBridge() (common.Address, error) {
	return _AbsRollup.Contract.RollupEventBridge(&_AbsRollup.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) RollupEventBridge() (common.Address, error) {
	return _AbsRollup.Contract.RollupEventBridge(&_AbsRollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_AbsRollup *AbsRollupCaller) SequencerBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "sequencerBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_AbsRollup *AbsRollupSession) SequencerBridge() (common.Address, error) {
	return _AbsRollup.Contract.SequencerBridge(&_AbsRollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_AbsRollup *AbsRollupCallerSession) SequencerBridge() (common.Address, error) {
	return _AbsRollup.Contract.SequencerBridge(&_AbsRollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_AbsRollup *AbsRollupSession) StakerCount() (*big.Int, error) {
	return _AbsRollup.Contract.StakerCount(&_AbsRollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) StakerCount() (*big.Int, error) {
	return _AbsRollup.Contract.StakerCount(&_AbsRollup.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_AbsRollup *AbsRollupSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.WithdrawableFunds(&_AbsRollup.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _AbsRollup.Contract.WithdrawableFunds(&_AbsRollup.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_AbsRollup *AbsRollupCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_AbsRollup *AbsRollupSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.ZombieAddress(&_AbsRollup.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_AbsRollup *AbsRollupCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _AbsRollup.Contract.ZombieAddress(&_AbsRollup.CallOpts, zombieNum)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_AbsRollup *AbsRollupSession) ZombieCount() (*big.Int, error) {
	return _AbsRollup.Contract.ZombieCount(&_AbsRollup.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ZombieCount() (*big.Int, error) {
	return _AbsRollup.Contract.ZombieCount(&_AbsRollup.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_AbsRollup *AbsRollupCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AbsRollup.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_AbsRollup *AbsRollupSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _AbsRollup.Contract.ZombieLatestStakedNode(&_AbsRollup.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_AbsRollup *AbsRollupCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _AbsRollup.Contract.ZombieLatestStakedNode(&_AbsRollup.CallOpts, zombieNum)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_AbsRollup *AbsRollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "completeChallenge", winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_AbsRollup *AbsRollupSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.CompleteChallenge(&_AbsRollup.TransactOpts, winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_AbsRollup *AbsRollupTransactorSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.CompleteChallenge(&_AbsRollup.TransactOpts, winningStaker, losingStaker)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_AbsRollup *AbsRollupTransactor) ConfirmNextNode(opts *bind.TransactOpts, beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "confirmNextNode", beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_AbsRollup *AbsRollupSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.ConfirmNextNode(&_AbsRollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_AbsRollup *AbsRollupTransactorSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.ConfirmNextNode(&_AbsRollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_AbsRollup *AbsRollupTransactor) CreateChallenge(opts *bind.TransactOpts, stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "createChallenge", stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_AbsRollup *AbsRollupSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.CreateChallenge(&_AbsRollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_AbsRollup *AbsRollupTransactorSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.CreateChallenge(&_AbsRollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_AbsRollup *AbsRollupTransactor) Initialize(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "initialize", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_AbsRollup *AbsRollupSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.Initialize(&_AbsRollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_AbsRollup *AbsRollupTransactorSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.Initialize(&_AbsRollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AbsRollup *AbsRollupTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AbsRollup *AbsRollupSession) Pause() (*types.Transaction, error) {
	return _AbsRollup.Contract.Pause(&_AbsRollup.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AbsRollup *AbsRollupTransactorSession) Pause() (*types.Transaction, error) {
	return _AbsRollup.Contract.Pause(&_AbsRollup.TransactOpts)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_AbsRollup *AbsRollupTransactor) ReduceDeposit(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "reduceDeposit", target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_AbsRollup *AbsRollupSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.ReduceDeposit(&_AbsRollup.TransactOpts, target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_AbsRollup *AbsRollupTransactorSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.ReduceDeposit(&_AbsRollup.TransactOpts, target)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_AbsRollup *AbsRollupTransactor) RejectNextNode(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "rejectNextNode", stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_AbsRollup *AbsRollupSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.RejectNextNode(&_AbsRollup.TransactOpts, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_AbsRollup *AbsRollupTransactorSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.RejectNextNode(&_AbsRollup.TransactOpts, stakerAddress)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupTransactor) RemoveOldOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "removeOldOutbox", _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveOldOutbox(&_AbsRollup.TransactOpts, _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupTransactorSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveOldOutbox(&_AbsRollup.TransactOpts, _outbox)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_AbsRollup *AbsRollupTransactor) RemoveOldZombies(opts *bind.TransactOpts, startIndex *big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "removeOldZombies", startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_AbsRollup *AbsRollupSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveOldZombies(&_AbsRollup.TransactOpts, startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_AbsRollup *AbsRollupTransactorSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveOldZombies(&_AbsRollup.TransactOpts, startIndex)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_AbsRollup *AbsRollupTransactor) RemoveZombie(opts *bind.TransactOpts, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "removeZombie", zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_AbsRollup *AbsRollupSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveZombie(&_AbsRollup.TransactOpts, zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_AbsRollup *AbsRollupTransactorSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _AbsRollup.Contract.RemoveZombie(&_AbsRollup.TransactOpts, zombieNum, maxNodes)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AbsRollup *AbsRollupTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AbsRollup *AbsRollupSession) Resume() (*types.Transaction, error) {
	return _AbsRollup.Contract.Resume(&_AbsRollup.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AbsRollup *AbsRollupTransactorSession) Resume() (*types.Transaction, error) {
	return _AbsRollup.Contract.Resume(&_AbsRollup.TransactOpts)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_AbsRollup *AbsRollupTransactor) ReturnOldDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "returnOldDeposit", stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_AbsRollup *AbsRollupSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.ReturnOldDeposit(&_AbsRollup.TransactOpts, stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_AbsRollup *AbsRollupTransactorSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.ReturnOldDeposit(&_AbsRollup.TransactOpts, stakerAddress)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_AbsRollup *AbsRollupTransactor) SetInbox(opts *bind.TransactOpts, _inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "setInbox", _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_AbsRollup *AbsRollupSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _AbsRollup.Contract.SetInbox(&_AbsRollup.TransactOpts, _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_AbsRollup *AbsRollupTransactorSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _AbsRollup.Contract.SetInbox(&_AbsRollup.TransactOpts, _inbox, _enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupTransactor) SetOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "setOutbox", _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.SetOutbox(&_AbsRollup.TransactOpts, _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_AbsRollup *AbsRollupTransactorSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.SetOutbox(&_AbsRollup.TransactOpts, _outbox)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_AbsRollup *AbsRollupTransactor) StakeOnExistingNode(opts *bind.TransactOpts, nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "stakeOnExistingNode", nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_AbsRollup *AbsRollupSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.StakeOnExistingNode(&_AbsRollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_AbsRollup *AbsRollupTransactorSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.StakeOnExistingNode(&_AbsRollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_AbsRollup *AbsRollupTransactor) StakeOnNewNode(opts *bind.TransactOpts, expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "stakeOnNewNode", expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_AbsRollup *AbsRollupSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.StakeOnNewNode(&_AbsRollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_AbsRollup *AbsRollupTransactorSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.StakeOnNewNode(&_AbsRollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_AbsRollup *AbsRollupTransactor) UpgradeImplementation(opts *bind.TransactOpts, _newRollup common.Address) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "upgradeImplementation", _newRollup)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_AbsRollup *AbsRollupSession) UpgradeImplementation(_newRollup common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.UpgradeImplementation(&_AbsRollup.TransactOpts, _newRollup)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_AbsRollup *AbsRollupTransactorSession) UpgradeImplementation(_newRollup common.Address) (*types.Transaction, error) {
	return _AbsRollup.Contract.UpgradeImplementation(&_AbsRollup.TransactOpts, _newRollup)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_AbsRollup *AbsRollupTransactor) UpgradeImplementationAndCall(opts *bind.TransactOpts, _newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _AbsRollup.contract.Transact(opts, "upgradeImplementationAndCall", _newRollup, _data)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_AbsRollup *AbsRollupSession) UpgradeImplementationAndCall(_newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.UpgradeImplementationAndCall(&_AbsRollup.TransactOpts, _newRollup, _data)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_AbsRollup *AbsRollupTransactorSession) UpgradeImplementationAndCall(_newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _AbsRollup.Contract.UpgradeImplementationAndCall(&_AbsRollup.TransactOpts, _newRollup, _data)
}

// AbsRollupNodeConfirmedIterator is returned from FilterNodeConfirmed and is used to iterate over the raw logs and unpacked data for NodeConfirmed events raised by the AbsRollup contract.
type AbsRollupNodeConfirmedIterator struct {
	Event *AbsRollupNodeConfirmed // Event containing the contract specifics and raw log

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
func (it *AbsRollupNodeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupNodeConfirmed)
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
		it.Event = new(AbsRollupNodeConfirmed)
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
func (it *AbsRollupNodeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupNodeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupNodeConfirmed represents a NodeConfirmed event raised by the AbsRollup contract.
type AbsRollupNodeConfirmed struct {
	NodeNum        *big.Int
	AfterSendAcc   [32]byte
	AfterSendCount *big.Int
	AfterLogAcc    [32]byte
	AfterLogCount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeConfirmed is a free log retrieval operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_AbsRollup *AbsRollupFilterer) FilterNodeConfirmed(opts *bind.FilterOpts, nodeNum []*big.Int) (*AbsRollupNodeConfirmedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupNodeConfirmedIterator{contract: _AbsRollup.contract, event: "NodeConfirmed", logs: logs, sub: sub}, nil
}

// WatchNodeConfirmed is a free log subscription operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_AbsRollup *AbsRollupFilterer) WatchNodeConfirmed(opts *bind.WatchOpts, sink chan<- *AbsRollupNodeConfirmed, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupNodeConfirmed)
				if err := _AbsRollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
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

// ParseNodeConfirmed is a log parse operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_AbsRollup *AbsRollupFilterer) ParseNodeConfirmed(log types.Log) (*AbsRollupNodeConfirmed, error) {
	event := new(AbsRollupNodeConfirmed)
	if err := _AbsRollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the AbsRollup contract.
type AbsRollupNodeCreatedIterator struct {
	Event *AbsRollupNodeCreated // Event containing the contract specifics and raw log

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
func (it *AbsRollupNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupNodeCreated)
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
		it.Event = new(AbsRollupNodeCreated)
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
func (it *AbsRollupNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupNodeCreated represents a NodeCreated event raised by the AbsRollup contract.
type AbsRollupNodeCreated struct {
	NodeNum                 *big.Int
	ParentNodeHash          [32]byte
	NodeHash                [32]byte
	ExecutionHash           [32]byte
	InboxMaxCount           *big.Int
	AfterInboxBatchEndCount *big.Int
	AfterInboxBatchAcc      [32]byte
	AssertionBytes32Fields  [2][3][32]byte
	AssertionIntFields      [2][4]*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_AbsRollup *AbsRollupFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int, parentNodeHash [][32]byte) (*AbsRollupNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupNodeCreatedIterator{contract: _AbsRollup.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_AbsRollup *AbsRollupFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *AbsRollupNodeCreated, nodeNum []*big.Int, parentNodeHash [][32]byte) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupNodeCreated)
				if err := _AbsRollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
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

// ParseNodeCreated is a log parse operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_AbsRollup *AbsRollupFilterer) ParseNodeCreated(log types.Log) (*AbsRollupNodeCreated, error) {
	event := new(AbsRollupNodeCreated)
	if err := _AbsRollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupNodeRejectedIterator is returned from FilterNodeRejected and is used to iterate over the raw logs and unpacked data for NodeRejected events raised by the AbsRollup contract.
type AbsRollupNodeRejectedIterator struct {
	Event *AbsRollupNodeRejected // Event containing the contract specifics and raw log

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
func (it *AbsRollupNodeRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupNodeRejected)
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
		it.Event = new(AbsRollupNodeRejected)
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
func (it *AbsRollupNodeRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupNodeRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupNodeRejected represents a NodeRejected event raised by the AbsRollup contract.
type AbsRollupNodeRejected struct {
	NodeNum *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeRejected is a free log retrieval operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_AbsRollup *AbsRollupFilterer) FilterNodeRejected(opts *bind.FilterOpts, nodeNum []*big.Int) (*AbsRollupNodeRejectedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupNodeRejectedIterator{contract: _AbsRollup.contract, event: "NodeRejected", logs: logs, sub: sub}, nil
}

// WatchNodeRejected is a free log subscription operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_AbsRollup *AbsRollupFilterer) WatchNodeRejected(opts *bind.WatchOpts, sink chan<- *AbsRollupNodeRejected, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupNodeRejected)
				if err := _AbsRollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
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

// ParseNodeRejected is a log parse operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_AbsRollup *AbsRollupFilterer) ParseNodeRejected(log types.Log) (*AbsRollupNodeRejected, error) {
	event := new(AbsRollupNodeRejected)
	if err := _AbsRollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupNodesDestroyedIterator is returned from FilterNodesDestroyed and is used to iterate over the raw logs and unpacked data for NodesDestroyed events raised by the AbsRollup contract.
type AbsRollupNodesDestroyedIterator struct {
	Event *AbsRollupNodesDestroyed // Event containing the contract specifics and raw log

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
func (it *AbsRollupNodesDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupNodesDestroyed)
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
		it.Event = new(AbsRollupNodesDestroyed)
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
func (it *AbsRollupNodesDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupNodesDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupNodesDestroyed represents a NodesDestroyed event raised by the AbsRollup contract.
type AbsRollupNodesDestroyed struct {
	StartNode *big.Int
	EndNode   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodesDestroyed is a free log retrieval operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_AbsRollup *AbsRollupFilterer) FilterNodesDestroyed(opts *bind.FilterOpts, startNode []*big.Int, endNode []*big.Int) (*AbsRollupNodesDestroyedIterator, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupNodesDestroyedIterator{contract: _AbsRollup.contract, event: "NodesDestroyed", logs: logs, sub: sub}, nil
}

// WatchNodesDestroyed is a free log subscription operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_AbsRollup *AbsRollupFilterer) WatchNodesDestroyed(opts *bind.WatchOpts, sink chan<- *AbsRollupNodesDestroyed, startNode []*big.Int, endNode []*big.Int) (event.Subscription, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupNodesDestroyed)
				if err := _AbsRollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
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

// ParseNodesDestroyed is a log parse operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_AbsRollup *AbsRollupFilterer) ParseNodesDestroyed(log types.Log) (*AbsRollupNodesDestroyed, error) {
	event := new(AbsRollupNodesDestroyed)
	if err := _AbsRollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AbsRollup contract.
type AbsRollupPausedIterator struct {
	Event *AbsRollupPaused // Event containing the contract specifics and raw log

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
func (it *AbsRollupPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupPaused)
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
		it.Event = new(AbsRollupPaused)
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
func (it *AbsRollupPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupPaused represents a Paused event raised by the AbsRollup contract.
type AbsRollupPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AbsRollup *AbsRollupFilterer) FilterPaused(opts *bind.FilterOpts) (*AbsRollupPausedIterator, error) {

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AbsRollupPausedIterator{contract: _AbsRollup.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AbsRollup *AbsRollupFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AbsRollupPaused) (event.Subscription, error) {

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupPaused)
				if err := _AbsRollup.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AbsRollup *AbsRollupFilterer) ParsePaused(log types.Log) (*AbsRollupPaused, error) {
	event := new(AbsRollupPaused)
	if err := _AbsRollup.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the AbsRollup contract.
type AbsRollupRollupChallengeStartedIterator struct {
	Event *AbsRollupRollupChallengeStarted // Event containing the contract specifics and raw log

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
func (it *AbsRollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupRollupChallengeStarted)
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
		it.Event = new(AbsRollupRollupChallengeStarted)
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
func (it *AbsRollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the AbsRollup contract.
type AbsRollupRollupChallengeStarted struct {
	ChallengeContract common.Address
	Asserter          common.Address
	Challenger        common.Address
	ChallengedNode    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_AbsRollup *AbsRollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts, challengeContract []common.Address) (*AbsRollupRollupChallengeStartedIterator, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupRollupChallengeStartedIterator{contract: _AbsRollup.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_AbsRollup *AbsRollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *AbsRollupRollupChallengeStarted, challengeContract []common.Address) (event.Subscription, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupRollupChallengeStarted)
				if err := _AbsRollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
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

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_AbsRollup *AbsRollupFilterer) ParseRollupChallengeStarted(log types.Log) (*AbsRollupRollupChallengeStarted, error) {
	event := new(AbsRollupRollupChallengeStarted)
	if err := _AbsRollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the AbsRollup contract.
type AbsRollupRollupCreatedIterator struct {
	Event *AbsRollupRollupCreated // Event containing the contract specifics and raw log

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
func (it *AbsRollupRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupRollupCreated)
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
		it.Event = new(AbsRollupRollupCreated)
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
func (it *AbsRollupRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupRollupCreated represents a RollupCreated event raised by the AbsRollup contract.
type AbsRollupRollupCreated struct {
	MachineHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_AbsRollup *AbsRollupFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*AbsRollupRollupCreatedIterator, error) {

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &AbsRollupRollupCreatedIterator{contract: _AbsRollup.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_AbsRollup *AbsRollupFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *AbsRollupRollupCreated) (event.Subscription, error) {

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupRollupCreated)
				if err := _AbsRollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_AbsRollup *AbsRollupFilterer) ParseRollupCreated(log types.Log) (*AbsRollupRollupCreated, error) {
	event := new(AbsRollupRollupCreated)
	if err := _AbsRollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupStakerReassignedIterator is returned from FilterStakerReassigned and is used to iterate over the raw logs and unpacked data for StakerReassigned events raised by the AbsRollup contract.
type AbsRollupStakerReassignedIterator struct {
	Event *AbsRollupStakerReassigned // Event containing the contract specifics and raw log

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
func (it *AbsRollupStakerReassignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupStakerReassigned)
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
		it.Event = new(AbsRollupStakerReassigned)
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
func (it *AbsRollupStakerReassignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupStakerReassignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupStakerReassigned represents a StakerReassigned event raised by the AbsRollup contract.
type AbsRollupStakerReassigned struct {
	Staker  common.Address
	NewNode *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakerReassigned is a free log retrieval operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_AbsRollup *AbsRollupFilterer) FilterStakerReassigned(opts *bind.FilterOpts, staker []common.Address) (*AbsRollupStakerReassignedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return &AbsRollupStakerReassignedIterator{contract: _AbsRollup.contract, event: "StakerReassigned", logs: logs, sub: sub}, nil
}

// WatchStakerReassigned is a free log subscription operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_AbsRollup *AbsRollupFilterer) WatchStakerReassigned(opts *bind.WatchOpts, sink chan<- *AbsRollupStakerReassigned, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupStakerReassigned)
				if err := _AbsRollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
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

// ParseStakerReassigned is a log parse operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_AbsRollup *AbsRollupFilterer) ParseStakerReassigned(log types.Log) (*AbsRollupStakerReassigned, error) {
	event := new(AbsRollupStakerReassigned)
	if err := _AbsRollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsRollupUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AbsRollup contract.
type AbsRollupUnpausedIterator struct {
	Event *AbsRollupUnpaused // Event containing the contract specifics and raw log

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
func (it *AbsRollupUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsRollupUnpaused)
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
		it.Event = new(AbsRollupUnpaused)
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
func (it *AbsRollupUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsRollupUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsRollupUnpaused represents a Unpaused event raised by the AbsRollup contract.
type AbsRollupUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AbsRollup *AbsRollupFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AbsRollupUnpausedIterator, error) {

	logs, sub, err := _AbsRollup.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AbsRollupUnpausedIterator{contract: _AbsRollup.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AbsRollup *AbsRollupFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AbsRollupUnpaused) (event.Subscription, error) {

	logs, sub, err := _AbsRollup.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsRollupUnpaused)
				if err := _AbsRollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AbsRollup *AbsRollupFilterer) ParseUnpaused(log types.Log) (*AbsRollupUnpaused, error) {
	event := new(AbsRollupUnpaused)
	if err := _AbsRollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupABI is the input ABI used to generate the binding from.
const ERC20RollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterSendAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"NodeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentNodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterInboxBatchEndCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterInboxBatchAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"indexed\":false,\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"NodeRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startNode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endNode\",\"type\":\"uint256\"}],\"name\":\"NodesDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNode\",\"type\":\"uint256\"}],\"name\":\"StakerReassigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"addToDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeSendAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"executionHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proposedTimes\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"maxMessageCounts\",\"type\":\"uint256[2]\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedBridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraChallengeTimeBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[7]\",\"name\":\"connectedContracts\",\"type\":\"address[7]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"newStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"requireUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireUnresolvedExists\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstUnresolvedNodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestNodeCreated\",\"type\":\"uint256\"}],\"name\":\"requiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerBridge\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expectedNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"},{\"internalType\":\"uint256\",\"name\":\"beforeProposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxMaxCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sequencerBatchProof\",\"type\":\"bytes\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"}],\"name\":\"upgradeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeImplementationAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC20RollupBin is the compiled bytecode used for deploying new contracts.
var ERC20RollupBin = "0x608060405234801561001057600080fd5b506000805460ff19908116600117909155600b8054909116905561573e806200003a6000396000f3fe608060405234801561001057600080fd5b50600436106102b45760003560e01c8063046f7da2146102b957806304a28064146102c35780631e83d30f146102fb5780632b2af0ab146103185780632e7acfa6146103355780632f30cabd1461033d5780633e55c0c7146103635780633e96576e146103875780633fe38627146103ad578063414f23fe1461043d57806345e38b641461046057806345e593ca14610468578063488ed1a91461050f5780634d26732d1461053d5780634f0f4aa91461054557806351ed6a3014610562578063567ca41b1461056a5780635c975abb146105905780635dbaf68b146105ac5780635e8ef106146105b45780636177fd18146105bc57806362a82d7d146105e257806363721d6b146105ff57806365f7f80d1461060757806367425daf1461060f57806369fd251c146106175780636b94c33b1461063d5780636f791d29146106635780636f7d00261461066b5780637427be511461069457806376e7e23b146106ba578063771b2f97146106c25780637ba9534a146106ca5780637e2d2155146106d257806381fbc98a146106f557806383f94db71461071b5780638456cb59146107415780638640ce5f146107495780638da5cb5b1461075157806391c657e8146107595780639e8a713f1461077f578063bf5ddcb114610787578063ce11e6ab14610805578063d01e66021461080d578063d735e21d1461082a578063d93fe9c414610832578063dff697871461083a578063e45b7ce614610842578063e8bd492214610870578063edfd03ed146108cc578063ef40a670146108e9578063f31d863f1461090f578063f33e1fac146109e0578063f3f0a03e146109fd578063f51de41b14610a29578063f851a44014610a31578063f8d1f19414610a39578063fa7803e614610a56578063fb64884e14610a84578063ff204f3b14610aa1575b600080fd5b6102c1610ac7565b005b6102e9600480360360208110156102d957600080fd5b50356001600160a01b0316610b1d565b60408051918252519081900360200190f35b6102c16004803603602081101561031157600080fd5b5035610bd5565b6102c16004803603602081101561032e57600080fd5b5035610c4d565b6102e9610ce9565b6102e96004803603602081101561035357600080fd5b50356001600160a01b0316610cef565b61036b610d0a565b604080516001600160a01b039092168252519081900360200190f35b6102e96004803603602081101561039d57600080fd5b50356001600160a01b0316610d19565b6102c160048036036102408110156103c457600080fd5b813591602081019160e08201916101e08101359161020082013591908101906102408101610220820135600160201b8111156103ff57600080fd5b82018360208201111561041157600080fd5b803590602001918460018302840111600160201b8311171561043257600080fd5b509092509050610d37565b6102c16004803603604081101561045357600080fd5b50803590602001356115ef565b6102e96117ba565b6102c160048036036101e081101561047f57600080fd5b8135916020810135916040820135916060810135916080820135916001600160a01b0360a082013581169260c08301359091169190810190610100810160e0820135600160201b8111156104d257600080fd5b8201836020820111156104e457600080fd5b803590602001918460018302840111600160201b8311171561050557600080fd5b91935091506117bf565b6102c1600480360361014081101561052657600080fd5b50604081016080820160c08301610100840161184d565b6102e96120fa565b61036b6004803603602081101561055b57600080fd5b503561211f565b61036b61213a565b6102c16004803603602081101561058057600080fd5b50356001600160a01b0316612149565b61059861224c565b604080519115158252519081900360200190f35b61036b612255565b6102e9612264565b610598600480360360208110156105d257600080fd5b50356001600160a01b031661226a565b61036b600480360360208110156105f857600080fd5b5035612292565b6102e96122bc565b6102e96122c2565b6102c16122c8565b61036b6004803603602081101561062d57600080fd5b50356001600160a01b0316612332565b6102c16004803603602081101561065357600080fd5b50356001600160a01b0316612353565b61059861274a565b6102e96004803603606081101561068157600080fd5b5080359060208101359060400135612753565b6102c1600480360360208110156106aa57600080fd5b50356001600160a01b031661276a565b6102e9612815565b6102e961281b565b6102e9612821565b6102c1600480360360408110156106e857600080fd5b5080359060200135612827565b6102e96004803603602081101561070b57600080fd5b50356001600160a01b0316612a07565b6102c16004803603602081101561073157600080fd5b50356001600160a01b0316612b29565b6102c1612be2565b6102e9612c36565b61036b612c3c565b6105986004803603602081101561076f57600080fd5b50356001600160a01b0316612c4b565b61036b612ca5565b6102c16004803603604081101561079d57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156107c757600080fd5b8201836020820111156107d957600080fd5b803590602001918460018302840111600160201b831117156107fa57600080fd5b509092509050612cb4565b61036b612daa565b61036b6004803603602081101561082357600080fd5b5035612db9565b6102e9612de8565b61036b612dee565b6102e9612dfd565b6102c16004803603604081101561085857600080fd5b506001600160a01b0381351690602001351515612e03565b6108966004803603602081101561088657600080fd5b50356001600160a01b0316612ea5565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b6102c1600480360360208110156108e257600080fd5b5035612ee1565b6102e9600480360360208110156108ff57600080fd5b50356001600160a01b0316612f46565b6102c1600480360360c081101561092557600080fd5b81359190810190604081016020820135600160201b81111561094657600080fd5b82018360208201111561095857600080fd5b803590602001918460018302840111600160201b8311171561097957600080fd5b919390929091602081019035600160201b81111561099657600080fd5b8201836020820111156109a857600080fd5b803590602001918460208302840111600160201b831117156109c957600080fd5b919350915080359060208101359060400135612f64565b6102e9600480360360208110156109f657600080fd5b50356134f7565b6102c160048036036040811015610a1357600080fd5b506001600160a01b03813516906020013561351f565b61036b61363c565b61036b61364b565b6102e960048036036020811015610a4f57600080fd5b503561365a565b6102c160048036036040811015610a6c57600080fd5b506001600160a01b038135811691602001351661366c565b6102c160048036036020811015610a9a57600080fd5b50356136d4565b6102c160048036036020811015610ab757600080fd5b50356001600160a01b03166137ec565b6016546001600160a01b03163314610b13576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610b1b6138a2565b565b600080610b286122bc565b90506000805b82811015610bcb57846001600160a01b0316639168ae72610b4e83612db9565b6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610b8b57600080fd5b505afa158015610b9f573d6000803e3d6000fd5b505050506040513d6020811015610bb557600080fd5b505115610bc3576001909101905b600101610b2e565b509150505b919050565b610bdd61224c565b15610c1d576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b610c2633613942565b6000610c306120fa565b905080821015610c3e578091505b610c4833836139d9565b505050565b610c55612de8565b811015610c9b576040805162461bcd60e51b815260206004820152600f60248201526e1053149150511657d11150d2511151608a1b604482015290519081900360640190fd5b610ca3612821565b811115610ce6576040805162461bcd60e51b815260206004820152600c60248201526b1113d154d39517d1561254d560a21b604482015290519081900360640190fd5b50565b600c5481565b6001600160a01b03166000908152600a602052604090205490565b6011546001600160a01b031681565b6001600160a01b031660009081526008602052604090206001015490565b610d3f61224c565b15610d7f576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b610d883361226a565b610dc6576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6000610dd133610d19565b9050610ddb6155f9565b6000610de68361211f565b90506000601160009054906101000a90046001600160a01b03166001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b158015610e3857600080fd5b505afa158015610e4c573d6000803e3d6000fd5b505050506040513d6020811015610e6257600080fd5b50519050610e6e61562e565b60408051808201909152610f28908c60026000835b82821015610ec45760408051606081810190925290808402860190600390839083908082843760009201919091525050508152600190910190602001610e83565b50506040805180820190915291508d905060026000835b82821015610f1c5760408051608081810190925290808402860190600490839083908082843760009201919091525050508152600190910190602001610edb565b505050508b8b86613a9a565b90506000806000856001600160a01b031663701da98e6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f6857600080fd5b505afa158015610f7c573d6000803e3d6000fd5b505050506040513d6020811015610f9257600080fd5b50518451610f9f90613ae8565b14610fe3576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b6011546020850151604090810151815163036648d960e11b81526024810182905260048101928352604481018d90526001600160a01b03909316926306cc91b2928e928e929091908190606401858580828437600081840152601f19601f820116905080830192505050945050505050604080518083038186803b15801561106a57600080fd5b505afa15801561107e573d6000803e3d6000fd5b505050506040513d604081101561109457600080fd5b508051602090910151855160e0015191945092506000906110b6904390613b7d565b9050604b8110156110fb576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b8451516020860151516000916111119190613b7d565b9050856000015161010001518660200151604001511015806111405750600e5461113c908390613bda565b8110155b80611166575085516060908101516020880151909101516064916111649190613b7d565b145b6111a3576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b6111c360046111bd600e5485613bda90919063ffffffff16565b90613bda565b811115611203576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b600e546000906112289061122261121b826001613b7d565b8590613c33565b90613c8b565b90506112b6816112b0611246600c5443613c3390919063ffffffff16565b8c6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561127f57600080fd5b505afa158015611293573d6000803e3d6000fd5b505050506040513d60208110156112a957600080fd5b5051613cef565b90613c33565b93506000896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b1580156112f357600080fd5b505afa158015611307573d6000803e3d6000fd5b505050506040513d602081101561131d57600080fd5b50519050801561136f5761136c856113348361211f565b6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561127f57600080fd5b94505b50508686602001516040015111156113bf576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b5050600080600080896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b15801561140057600080fd5b505afa158015611414573d6000803e3d6000fd5b505050506040513d602081101561142a57600080fd5b505111905080156114a9576114a2896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b15801561147157600080fd5b505afa158015611485573d6000803e3d6000fd5b505050506040513d602081101561149b57600080fd5b505161365a565b91506114b5565b6114b28b61365a565b91505b6114c4878588888f8787613d05565b9a50925050508f8114611515576040805162461bcd60e51b81526020600482015260146024820152730aa9c8ab0a08a86a88a88be9c9e888abe9082a6960631b604482015290519081900360640190fd5b61152933611521612821565b600c54613fbb565b505050505050506115398261365a565b611541612821565b7f8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a548b84608001518560400151866000015187602001518f8f6040518088815260200187815260200186815260200185815260200184815260200183600260600280828437600083820152601f01601f191690910190508261010080828437600083820152604051601f909101601f1916909201829003995090975050505050505050a3505050505050505050565b6115f761224c565b15611637576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b6116403361226a565b61167e576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b806116888361365a565b146116c7576040805162461bcd60e51b815260206004820152600a6024820152694e4f44455f52454f524760b01b604482015290519081900360640190fd5b6116cf612de8565b82101580156116e557506116e1612821565b8211155b6116ee57600080fd5b60006116f98361211f565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561173457600080fd5b505afa158015611748573d6000803e3d6000fd5b505050506040513d602081101561175e57600080fd5b505161176933610d19565b146117ad576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b610c483384600c54613fbb565b604b81565b6018546001600160a01b031661180f576040805162461bcd60e51b815260206004820152601060248201526f2722a2a22fa9aa20a5a2afaa27a5a2a760811b604482015290519081900360640190fd5b6118218a8a8a8a8a8a8a8a8a8a61414e565b5050601880546001600160a01b0319166001600160a01b03949094169390931790925550505050505050565b61185561224c565b15611895576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b60208401358435106118dc576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b6118e4612821565b6020850135111561192b576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b83356119356122c2565b1061197b576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b600061198d85825b602002013561211f565b9050600061199c866001611983565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156119d757600080fd5b505afa1580156119eb573d6000803e3d6000fd5b505050506040513d6020811015611a0157600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b158015611a4357600080fd5b505afa158015611a57573d6000803e3d6000fd5b505050506040513d6020811015611a6d57600080fd5b505114611aad576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b611ac78760005b60200201356001600160a01b0316613942565b611ad2876001611ab4565b604080516348b4573960e11b81526001600160a01b03893581166004830152915191841691639168ae7291602480820192602092909190829003018186803b158015611b1d57600080fd5b505afa158015611b31573d6000803e3d6000fd5b505050506040513d6020811015611b4757600080fd5b5051611b8f576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b604080516348b4573960e11b81526001600160a01b0360208a81013582166004840152925190841692639168ae729260248082019391829003018186803b158015611bd957600080fd5b505afa158015611bed573d6000803e3d6000fd5b505050506040513d6020811015611c0357600080fd5b5051611c4b576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b611c60853585358560005b60200201356144d1565b826001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b158015611c9957600080fd5b505afa158015611cad573d6000803e3d6000fd5b505050506040513d6020811015611cc357600080fd5b505114611d04576040805162461bcd60e51b815260206004820152600a6024820152694348414c5f484153483160b01b604482015290519081900360640190fd5b611d1960208087013590860135856001611c56565b816001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b158015611d5257600080fd5b505afa158015611d66573d6000803e3d6000fd5b505050506040513d6020811015611d7c57600080fd5b505114611dbd576040805162461bcd60e51b815260206004820152600a60248201526921a420a62fa420a9a41960b11b604482015290519081900360640190fd5b6000611f0e611e2f846001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611dfe57600080fd5b505afa158015611e12573d6000803e3d6000fd5b505050506040513d6020811015611e2857600080fd5b505161211f565b6001600160a01b031663d7ff5e356040518163ffffffff1660e01b815260040160206040518083038186803b158015611e6757600080fd5b505afa158015611e7b573d6000803e3d6000fd5b505050506040513d6020811015611e9157600080fd5b5051600d546112b090818960006020020135886001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611edc57600080fd5b505afa158015611ef0573d6000803e3d6000fd5b505050506040513d6020811015611f0657600080fd5b505190613b7d565b90506020850135811015611f4857611f406001600160a01b0389351689600160200201356001600160a01b0316614508565b5050506120f3565b6014546000906001600160a01b0390811690638ecaab119030908a35908935908e35168e600160200201356001600160a01b0316611fa08d600060028110611f8c57fe5b60200201358a613b7d90919063ffffffff16565b611fba8e600160200201358b613b7d90919063ffffffff16565b601154601054604080516001600160e01b031960e08d901b1681526001600160a01b039a8b166004820152602481019990995260448901979097529488166064880152928716608487015260a486019190915260c4850152841660e484015290921661010482015290516101248083019260209291908290030181600087803b15801561204657600080fd5b505af115801561205a573d6000803e3d6000fd5b505050506040513d602081101561207057600080fd5b505190506120996001600160a01b038a35168a600160200201356001600160a01b031683614583565b604080518a356001600160a01b0390811682526020808d01358216908301528a35828401529151918316917fa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda758799181900360600190a2505050505b5050505050565b600080612105612de8565b90506121194382612114612821565b6145cd565b91505090565b6000908152600560205260409020546001600160a01b031690565b6018546001600160a01b031681565b6016546001600160a01b03163314612195576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6012546001600160a01b03828116911614156121e5576040805162461bcd60e51b815260206004820152600a602482015269086aaa4be9eaaa8849eb60b31b604482015290519081900360640190fd5b601054604080516319dc7ae560e31b81526001600160a01b038481166004830152600060248301819052925193169263cee3d7289260448084019391929182900301818387803b15801561223857600080fd5b505af11580156120f3573d6000803e3d6000fd5b600b5460ff1690565b6014546001600160a01b031681565b600e5481565b6001600160a01b0316600090815260086020526040902060030154600160a01b900460ff1690565b6000600782815481106122a157fe5b6000918252602090912001546001600160a01b031692915050565b60095490565b60015490565b60006122d2612de8565b90506122dc6122c2565b811180156122f157506122ed612821565b8111155b610ce6576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b6001600160a01b039081166000908152600860205260409020600301541690565b61235b61224c565b1561239b576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b6123a36122c8565b60006123ad6122c2565b905060006123b9612de8565b905060006123c68261211f565b905082816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561240257600080fd5b505afa158015612416573d6000803e3d6000fd5b505050506040513d602081101561242c57600080fd5b505114156126ac5761243d8461226a565b61247b576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b61248c61248785610d19565b610c4d565b806001600160a01b0316639168ae72856040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156124d957600080fd5b505afa1580156124ed573d6000803e3d6000fd5b505050506040513d602081101561250357600080fd5b50511561254a576040805162461bcd60e51b815260206004820152601060248201526f14d51052d15117d3d397d5105491d15560821b604482015290519081900360640190fd5b806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b15801561258357600080fd5b505afa158015612597573d6000803e3d6000fd5b505050506125a48361211f565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b1580156125dc57600080fd5b505afa1580156125f0573d6000803e3d6000fd5b505050506125fe6000612ee1565b61260781610b1d565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561264057600080fd5b505afa158015612654573d6000803e3d6000fd5b505050506040513d602081101561266a57600080fd5b5051146126ac576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b6126b4614860565b60135460408051630c2a09ad60e21b81526004810185905290516001600160a01b03909216916330a826b49160248082019260009290919082900301818387803b15801561270157600080fd5b505af1158015612715573d6000803e3d6000fd5b50506040518492507f9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a39150600090a250505050565b60005460ff1690565b60006127608484846145cd565b90505b9392505050565b61277261224c565b156127b2576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b6127ba6122c2565b6127c382610d19565b1115612803576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b61280c81613942565b610ce681614876565b600f5481565b600d5481565b60035490565b61282f61224c565b1561286f576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b6128776122bc565b8211156128bc576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b60006128c783612db9565b905060006128d4846134f7565b90506000806128e1612de8565b90505b8083101580156128f357508482105b156129df5760006129038461211f565b9050806001600160a01b03166396a9fdc0866040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561295457600080fd5b505af1158015612968573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156129a557600080fd5b505afa1580156129b9573d6000803e3d6000fd5b505050506040513d60208110156129cf57600080fd5b50519350506001909101906128e4565b808310156129f5576129f0866148c9565b6129ff565b6129ff8684614965565b505050505050565b6000612a1161224c565b15612a51576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b6000612a5c3361498c565b6018546040805163a9059cbb60e01b81526001600160a01b03878116600483015260248201859052915193945091169163a9059cbb916044808201926020929091908290030181600087803b158015612ab457600080fd5b505af1158015612ac8573d6000803e3d6000fd5b505050506040513d6020811015612ade57600080fd5b5051612b23576040805162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b604482015290519081900360640190fd5b92915050565b6016546001600160a01b03163314612b75576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6017546040805163266a23b160e21b815230600482018190526001600160a01b0385811660248401529251909392909216916399a88ec49160448082019260009290919082900301818387803b158015612bce57600080fd5b505af11580156129ff573d6000803e3d6000fd5b6016546001600160a01b03163314612c2e576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610b1b6149ab565b60045490565b6016546001600160a01b031681565b6000805b600954811015612c9c5760098181548110612c6657fe5b60009182526020909120600290910201546001600160a01b0384811691161415612c94576001915050610bd0565b600101612c4f565b50600092915050565b6013546001600160a01b031681565b6016546001600160a01b03163314612d00576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601754604051639623609d60e01b815230600482018181526001600160a01b0387811660248501526060604485019081526064850187905292941692639623609d92859289928992899291608401848480828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b158015612d8c57600080fd5b505af1158015612da0573d6000803e3d6000fd5b5050505050505050565b6012546001600160a01b031681565b600060098281548110612dc857fe5b60009182526020909120600290910201546001600160a01b031692915050565b60025490565b6015546001600160a01b031681565b60075490565b6016546001600160a01b03163314612e4f576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6010546040805163722dbe7360e11b81526001600160a01b03858116600483015284151560248301529151919092169163e45b7ce691604480830192600092919082900301818387803b158015612bce57600080fd5b6008602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6000612eeb6122bc565b90506000612ef7612de8565b9050825b82811015612f40575b81612f0e826134f7565b1015612f3857612f1d816148c9565b60001990920191828110612f3357505050610ce6565b612f04565b600101612efb565b50505050565b6001600160a01b031660009081526008602052604090206002015490565b612f6c61224c565b15612fac576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b612fb46122c8565b6000612fbe612dfd565b11612ffd576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f5354414b45525360b01b604482015290519081900360640190fd5b6000613007612de8565b905060006130148261211f565b9050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b15801561304f57600080fd5b505afa158015613063573d6000803e3d6000fd5b5050505061306f6122c2565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156130a857600080fd5b505afa1580156130bc573d6000803e3d6000fd5b505050506040513d60208110156130d257600080fd5b505114613115576040805162461bcd60e51b815260206004820152600c60248201526b24a72b20a624a22fa82922ab60a11b604482015290519081900360640190fd5b6131256131206122c2565b61211f565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b15801561315d57600080fd5b505afa158015613171573d6000803e3d6000fd5b5050505061317f6000612ee1565b61319361318b82610b1d565b6112b0612dfd565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b1580156131cc57600080fd5b505afa1580156131e0573d6000803e3d6000fd5b505050506040513d60208110156131f657600080fd5b50511461323b576040805162461bcd60e51b815260206004820152600e60248201526d1393d517d0531317d4d51052d15160921b604482015290519081900360640190fd5b60006132bc8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c918291850190849080828437600081840152601f19601f820116905080830192505050505050508d614a29565b90506132cb8b82878988614b2a565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b15801561330457600080fd5b505afa158015613318573d6000803e3d6000fd5b505050506040513d602081101561332e57600080fd5b505114613371576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b60125460408051630c72684760e01b815260048101918252604481018c90526001600160a01b0390921691630c726847918d918d918d918d919081906024810190606401878780828437600083820152601f01601f19169091018481038352858152602090810191508690860280828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b15801561341b57600080fd5b505af115801561342f573d6000803e3d6000fd5b5050505061343b614b71565b601354604080516316b9109b60e01b81526004810186905290516001600160a01b03909216916316b9109b9160248082019260009290919082900301818387803b15801561348857600080fd5b505af115801561349c573d6000803e3d6000fd5b505060408051848152602081018a90528082018990526060810188905290518693507f2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a292509081900360800190a25050505050505050505050565b60006009828154811061350657fe5b9060005260206000209060020201600101549050919050565b61352761224c565b15613567576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b6135718282614b8a565b601854604080516323b872dd60e01b81523360048201523060248201526044810184905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156135cb57600080fd5b505af11580156135df573d6000803e3d6000fd5b505050506040513d60208110156135f557600080fd5b5051613638576040805162461bcd60e51b815260206004820152600d60248201526c1514905394d1915497d1905253609a1b604482015290519081900360640190fd5b5050565b6010546001600160a01b031681565b6017546001600160a01b031681565b60009081526006602052604090205490565b6136768282614be5565b6001600160a01b0316336001600160a01b0316146136ca576040805162461bcd60e51b815260206004820152600c60248201526b2ba927a723afa9a2a72222a960a11b604482015290519081900360640190fd5b6136388282614508565b6136dc61224c565b1561371c576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b61372581614ca7565b601854604080516323b872dd60e01b81523360048201523060248201526044810184905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b15801561377f57600080fd5b505af1158015613793573d6000803e3d6000fd5b505050506040513d60208110156137a957600080fd5b5051610ce6576040805162461bcd60e51b815260206004820152600d60248201526c1514905394d1915497d1905253609a1b604482015290519081900360640190fd5b6016546001600160a01b03163314613838576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601280546001600160a01b0319166001600160a01b03838116918217909255601054604080516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b15801561223857600080fd5b6138aa61224c565b6138f2576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b600b805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613925614e42565b604080516001600160a01b039092168252519081900360200190a1565b61394b8161226a565b613989576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b600061399482612332565b6001600160a01b031614610ce6576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600860205260408120600281015480841115613a3e576040805162461bcd60e51b815260206004820152601060248201526f544f4f5f4c4954544c455f5354414b4560801b604482015290519081900360640190fd5b6000613a4a8286613b7d565b600284018690556001600160a01b0387166000908152600a6020526040902054909150613a779082613c33565b6001600160a01b0387166000908152600a60205260409020559250505092915050565b613aa261562e565b60408051808201909152865186518291613abd918888614e46565b8152602001613adc886001602002015188600160200201514387614e46565b90529695505050505050565b6000816000015182602001518360400151846060015185608001518660a001518760c001518860e00151896101000151604051602001808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019950505050505050505050604051602081830303815290604052805190602001209050919050565b600082821115613bd4576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082613be957506000612b23565b82820282848281613bf657fe5b04146127635760405162461bcd60e51b81526004018080602001828103825260218152602001806156c86021913960400191505060405180910390fd5b600082820183811015612763576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b6000808211613cde576040805162461bcd60e51b815260206004820152601a602482015279536166654d6174683a206469766973696f6e206279207a65726f60301b604482015290519081900360640190fd5b818381613ce757fe5b049392505050565b6000818311613cfe5781612763565b5090919050565b6000613d0f6155f9565b613d176155f9565b601160009054906101000a90046001600160a01b03166001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b158015613d6557600080fd5b505afa158015613d79573d6000803e3d6000fd5b505050506040513d6020811015613d8f57600080fd5b50516040820152613d9f8661211f565b6001600160a01b031660a08201526000613db7612821565b6001019050613dc58b614ee4565b60808301528882526020820188905260135460408051638b8ca19960e01b815260048101849052602481018a9052604481018d905233606482015290516001600160a01b0390921691638b8ca1999160848082019260009290919082900301818387803b158015613e3557600080fd5b505af1158015613e49573d6000803e3d6000fd5b505060155460208e01516001600160a01b03909116925063d45ab2b59150613e7090613ae8565b613e7f8e866080015143614f15565b613e888f614f2a565b8b8f6040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b158015613edc57600080fd5b505af1158015613ef0573d6000803e3d6000fd5b505050506040513d6020811015613f0657600080fd5b50516001600160a01b031660608301525060808101516020820151600091613f319187918991614f5a565b9050613f41826060015182614fbe565b8160a001516001600160a01b0316631bc09d0a613f5c612821565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015613f9257600080fd5b505af1158015613fa6573d6000803e3d6000fd5b50929d939c50929a5050505050505050505050565b6001600160a01b0380841660008181526008602090815260408083208784526005835281842054825163123334b760e11b815260048101969096529151909591909116938492632466696e9260248084019382900301818787803b15801561402257600080fd5b505af1158015614036573d6000803e3d6000fd5b505050506040513d602081101561404c57600080fd5b505160018085018790559091508114156129ff57600060056000846001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561409f57600080fd5b505afa1580156140b3573d6000803e3d6000fd5b505050506040513d60208110156140c957600080fd5b505181526020810191909152604001600020546001600160a01b0316905080636971dfe56140f74388613c33565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561412d57600080fd5b505af1158015614141573d6000803e3d6000fd5b5050505050505050505050565b600c5415614192576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b886141d6576040805162461bcd60e51b815260206004820152600f60248201526e10905117d0d3d39197d411549253d1608a1b604482015290519081900360640190fd5b601080546001600160a01b03199081166001600160a01b03602085013581169190911792839055601180548316604080870135841691909117909155601280546060870135841694168417905580516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b15801561426957600080fd5b505af115801561427d573d6000803e3d6000fd5b505050508060046007811061428e57fe5b601380546001600160a01b0319166001600160a01b0360209390930293909301358216929092179091556010546040805163722dbe7360e11b8152608085013584166004820152600160248201529051919092169163e45b7ce691604480830192600092919082900301818387803b15801561430957600080fd5b505af115801561431d573d6000803e3d6000fd5b505060135460405163b0f2af2960e01b8152600481018d8152602482018d9052604482018c9052606482018b90526001600160a01b038a8116608484015289811660a484015260e060c4840190815260e484018990529316945063b0f2af2993508d928d928d928d928d928d928d928d9261010401848480828437600081840152601f19601f8201169050808301925050509950505050505050505050600060405180830381600087803b1580156143d457600080fd5b505af11580156143e8573d6000803e3d6000fd5b50505050806005600781106143f957fe5b601480546001600160a01b03199081166001600160a01b03602094909402949094013583169390931790556015805490921660c084013590911617905560006144418b615008565b905061444c816150ff565b600c8a9055600d899055600e889055600f879055601680546001600160a01b038781166001600160a01b031992831617909255601780549091168435909216919091179055604080518c815290517f4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d9181900360200190a15050505050505050505050565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b600061451382612f46565b9050600061452084612f46565b9050808211156145415761453e61453784836139d9565b8390613b7d565b91505b6002820461454f858261514e565b6145598382613b7d565b92506145648561517f565b60165461457a906001600160a01b0316846151a9565b6120f3846151ec565b6001600160a01b03928316600090815260086020526040808220600390810180549487166001600160a01b0319958616811790915594909516825290209092018054909216179055565b6000816001840314156145e35750600f54612763565b60006145ee8461211f565b6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561462657600080fd5b505afa15801561463a573d6000803e3d6000fd5b505050506040513d602081101561465057600080fd5b5051905080851015614666575050600f54612763565b61466e615653565b506040805161014081018252600181526201e05b60208201526201f7d191810191909152620138916060820152620329e160808201526201be4360a08201526204cb8c60c08201526201fbc460e082015262036d32610100820152620279736101208201526146db615653565b506040805161014081018252600181526201c03060208201526201b6999181019190915261fde26060820152620265c6608082015262013b8e60a0820152620329e160c08201526201389160e08201526201f7d161010082015262015375610120820152600061474b8885613b7d565b90506000614769600c54611222600a85613bda90919063ffffffff16565b905060ff61477882600a613c8b565b1061478c5760001995505050505050612763565b600061479982600a613c8b565b60020a9050600085600a8406600a81106147af57fe5b602002015162ffffff168202905085600a8406600a81106147cc57fe5b602002015162ffffff168282816147df57fe5b04146147f657600019975050505050505050612763565b600061481b86600a8606600a811061480a57fe5b6020020151839062ffffff16613c8b565b905080614826575060015b600f54808202908290828161483757fe5b0414614850576000199950505050505050505050612763565b9c9b505050505050505050505050565b61486b60025461529c565b600280546001019055565b6001600160a01b03811660009081526008602090815260408083206002810154600a9093529220546148a791613c33565b6001600160a01b0383166000908152600a60205260409020556136388261531e565b6009805460001981019081106148db57fe5b9060005260206000209060020201600982815481106148f657fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600191820154910155600980548061493957fe5b60008281526020812060026000199093019283020180546001600160a01b031916815560010155905550565b806009838154811061497357fe5b9060005260206000209060020201600101819055505050565b6001600160a01b03166000908152600a60205260408120805491905590565b6149b361224c565b156149f3576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b600b805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613925614e42565b81518351600091829184835b83811015614adc576000888281518110614a4b57fe5b60200260200101519050838187011115614a9b576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868b0181018290206040805180840196909652858101919091528051808603820181526060909501905283519301929092209190940193600101614a35565b50818414614b1f576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b979650505050505050565b60408051602080820197909752808201959095526060850192909252608084019290925260a0808401929092528051808403909201825260c0909201909152805191012090565b614b7c60015461529c565b600280546001818155019055565b614b9261224c565b15614bd2576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b614bdb82613942565b613638828261514e565b6001600160a01b0380831660009081526008602052604080822084841683529082206003808301549082015493949293919290811691168114614c59576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b038116614c9e576040805162461bcd60e51b81526020600482015260076024820152661393d7d0d2105360ca1b604482015290519081900360640190fd5b95945050505050565b614caf61224c565b15614cef576040805162461bcd60e51b815260206004820152601060248201526000805160206156e9833981519152604482015290519081900360640190fd5b614cf83361226a565b15614d3b576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b614d4433612c4b565b15614d89576040805162461bcd60e51b815260206004820152601060248201526f5354414b45525f49535f5a4f4d42494560801b604482015290519081900360640190fd5b614d916120fa565b811015614dd8576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b614de23382615444565b6013546001600160a01b031663f03c04a533614dfc6122c2565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561223857600080fd5b3390565b614e4e615672565b604080516101208101825285518152865160208201529081018560016020020151815260200185600260048110614e8157fe5b6020020151815260200185600360048110614e9857fe5b6020020151815260200186600160038110614eaf57fe5b6020020151815260200186600260038110614ec657fe5b60200201518152602001848152602001838152509050949350505050565b80518051602083015151600092612b23929182900390614f039061550f565b614f10866020015161550f565b615544565b600061276083838660200151604001516144d1565b805160a09081015160208301519182015160c08301516060840151608090940151600094612b2394939291614b2a565b60008085614f69576000614f6c565b60015b905080858585604051602001808560ff1660f81b815260010184815260200183815260200182815260200194505050505060405160208183030381529060405280519060200120915050949350505050565b60038054600101808255600090815260056020908152604080832080546001600160a01b0319166001600160a01b0397909716969096179095559154815260069091529190912055565b60008061505e604051806101200160405280600081526020018581526020016000815260200160008152602001600081526020016000801b81526020016000801b81526020014381526020016001815250613ae8565b6015546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905243608483015291519394506001600160a01b039092169263d45ab2b59260a4808201936020939283900390910190829087803b1580156150cc57600080fd5b505af11580156150e0573d6000803e3d6000fd5b505050506040513d60208110156150f657600080fd5b50519392505050565b6000805260056020527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc80546001600160a01b0319166001600160a01b03929092169190911790556001600255565b6001600160a01b038216600090815260086020526040902060028101546151759083613c33565b6002909101555050565b6001600160a01b0316600090815260086020526040902060030180546001600160a01b0319169055565b6001600160a01b0382166000908152600a60205260409020546151cc9082613c33565b6001600160a01b039092166000908152600a602052604090209190915550565b6001600160a01b0381811660008181526008602090815260408083208151808301909252938152600180850154928201928352600980549182018155909352517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af600290930292830180546001600160a01b031916919095161790935591517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7b0909201919091556136388261531e565b60008181526005602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b1580156152e857600080fd5b505af11580156152fc573d6000803e3d6000fd5b50505060009182525060056020526040902080546001600160a01b0319169055565b6001600160a01b0381166000908152600860205260409020805460078054600019810190811061534a57fe5b600091825260209091200154600780546001600160a01b03909216918390811061537057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508060086000600784815481106153b057fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205560078054806153e057fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03949094168152600890935250506040812081815560018101829055600281019190915560030180546001600160a81b0319169055565b6007805460018082019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688810180546001600160a01b039586166001600160a01b031991821681179092556040805160a08101825293845284546020858101918252858301978852600060608701818152608088018981529682526008909252929092209451855551948401949094559351600283015591516003909101805492511515600160a01b0260ff60a01b199290951692909316919091171691909117905543600455565b6000612b23826000015161553f846040015185602001518660a0015187606001518860c001518960800151615582565b6155cd565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b6040518060400160405280615641615672565b815260200161564e615672565b905290565b604051806101400160405280600a906020820280368337509192915050565b604051806101200160405280600081526020016000801916815260200160008152602001600081526020016000815260200160008019168152602001600080191681526020016000815260200160008152509056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775061757361626c653a2070617573656400000000000000000000000000000000a26469706673582212209a13ce46cdd7078a60aa34749bb624e38e4504717fea17e2ab6d3f2ff644ed2a64736f6c634300060c0033"

// DeployERC20Rollup deploys a new Ethereum contract, binding an instance of ERC20Rollup to it.
func DeployERC20Rollup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Rollup, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20RollupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20RollupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Rollup{ERC20RollupCaller: ERC20RollupCaller{contract: contract}, ERC20RollupTransactor: ERC20RollupTransactor{contract: contract}, ERC20RollupFilterer: ERC20RollupFilterer{contract: contract}}, nil
}

// ERC20Rollup is an auto generated Go binding around an Ethereum contract.
type ERC20Rollup struct {
	ERC20RollupCaller     // Read-only binding to the contract
	ERC20RollupTransactor // Write-only binding to the contract
	ERC20RollupFilterer   // Log filterer for contract events
}

// ERC20RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20RollupSession struct {
	Contract     *ERC20Rollup      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20RollupCallerSession struct {
	Contract *ERC20RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20RollupTransactorSession struct {
	Contract     *ERC20RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20RollupRaw struct {
	Contract *ERC20Rollup // Generic contract binding to access the raw methods on
}

// ERC20RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20RollupCallerRaw struct {
	Contract *ERC20RollupCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20RollupTransactorRaw struct {
	Contract *ERC20RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Rollup creates a new instance of ERC20Rollup, bound to a specific deployed contract.
func NewERC20Rollup(address common.Address, backend bind.ContractBackend) (*ERC20Rollup, error) {
	contract, err := bindERC20Rollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Rollup{ERC20RollupCaller: ERC20RollupCaller{contract: contract}, ERC20RollupTransactor: ERC20RollupTransactor{contract: contract}, ERC20RollupFilterer: ERC20RollupFilterer{contract: contract}}, nil
}

// NewERC20RollupCaller creates a new read-only instance of ERC20Rollup, bound to a specific deployed contract.
func NewERC20RollupCaller(address common.Address, caller bind.ContractCaller) (*ERC20RollupCaller, error) {
	contract, err := bindERC20Rollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupCaller{contract: contract}, nil
}

// NewERC20RollupTransactor creates a new write-only instance of ERC20Rollup, bound to a specific deployed contract.
func NewERC20RollupTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20RollupTransactor, error) {
	contract, err := bindERC20Rollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupTransactor{contract: contract}, nil
}

// NewERC20RollupFilterer creates a new log filterer instance of ERC20Rollup, bound to a specific deployed contract.
func NewERC20RollupFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20RollupFilterer, error) {
	contract, err := bindERC20Rollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupFilterer{contract: contract}, nil
}

// bindERC20Rollup binds a generic wrapper to an already deployed contract.
func bindERC20Rollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Rollup *ERC20RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Rollup.Contract.ERC20RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Rollup *ERC20RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ERC20RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Rollup *ERC20RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ERC20RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Rollup *ERC20RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Rollup *ERC20RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Rollup *ERC20RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.contract.Transact(opts, method, params...)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_ERC20Rollup *ERC20RollupCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "_stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_ERC20Rollup *ERC20RollupSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _ERC20Rollup.Contract.StakerMap(&_ERC20Rollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_ERC20Rollup *ERC20RollupCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _ERC20Rollup.Contract.StakerMap(&_ERC20Rollup.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) Admin() (common.Address, error) {
	return _ERC20Rollup.Contract.Admin(&_ERC20Rollup.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) Admin() (common.Address, error) {
	return _ERC20Rollup.Contract.Admin(&_ERC20Rollup.CallOpts)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.AmountStaked(&_ERC20Rollup.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.AmountStaked(&_ERC20Rollup.CallOpts, staker)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ArbGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "arbGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _ERC20Rollup.Contract.ArbGasSpeedLimitPerBlock(&_ERC20Rollup.CallOpts)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _ERC20Rollup.Contract.ArbGasSpeedLimitPerBlock(&_ERC20Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) BaseStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "baseStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) BaseStake() (*big.Int, error) {
	return _ERC20Rollup.Contract.BaseStake(&_ERC20Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) BaseStake() (*big.Int, error) {
	return _ERC20Rollup.Contract.BaseStake(&_ERC20Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) ChallengeFactory() (common.Address, error) {
	return _ERC20Rollup.Contract.ChallengeFactory(&_ERC20Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) ChallengeFactory() (common.Address, error) {
	return _ERC20Rollup.Contract.ChallengeFactory(&_ERC20Rollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ConfirmPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "confirmPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _ERC20Rollup.Contract.ConfirmPeriodBlocks(&_ERC20Rollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _ERC20Rollup.Contract.ConfirmPeriodBlocks(&_ERC20Rollup.CallOpts)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) CountStakedZombies(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "countStakedZombies", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.CountStakedZombies(&_ERC20Rollup.CallOpts, node)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.CountStakedZombies(&_ERC20Rollup.CallOpts, node)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_ERC20Rollup *ERC20RollupSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _ERC20Rollup.Contract.CurrentChallenge(&_ERC20Rollup.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _ERC20Rollup.Contract.CurrentChallenge(&_ERC20Rollup.CallOpts, staker)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _ERC20Rollup.Contract.CurrentRequiredStake(&_ERC20Rollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _ERC20Rollup.Contract.CurrentRequiredStake(&_ERC20Rollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) DelayedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "delayedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) DelayedBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.DelayedBridge(&_ERC20Rollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) DelayedBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.DelayedBridge(&_ERC20Rollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ExtraChallengeTimeBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "extraChallengeTimeBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _ERC20Rollup.Contract.ExtraChallengeTimeBlocks(&_ERC20Rollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _ERC20Rollup.Contract.ExtraChallengeTimeBlocks(&_ERC20Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) FirstUnresolvedNode() (*big.Int, error) {
	return _ERC20Rollup.Contract.FirstUnresolvedNode(&_ERC20Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _ERC20Rollup.Contract.FirstUnresolvedNode(&_ERC20Rollup.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_ERC20Rollup *ERC20RollupSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.GetNode(&_ERC20Rollup.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.GetNode(&_ERC20Rollup.CallOpts, nodeNum)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_ERC20Rollup *ERC20RollupCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_ERC20Rollup *ERC20RollupSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _ERC20Rollup.Contract.GetNodeHash(&_ERC20Rollup.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_ERC20Rollup *ERC20RollupCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _ERC20Rollup.Contract.GetNodeHash(&_ERC20Rollup.CallOpts, index)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_ERC20Rollup *ERC20RollupSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.GetStakerAddress(&_ERC20Rollup.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.GetStakerAddress(&_ERC20Rollup.CallOpts, stakerNum)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ERC20Rollup *ERC20RollupCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ERC20Rollup *ERC20RollupSession) IsMaster() (bool, error) {
	return _ERC20Rollup.Contract.IsMaster(&_ERC20Rollup.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ERC20Rollup *ERC20RollupCallerSession) IsMaster() (bool, error) {
	return _ERC20Rollup.Contract.IsMaster(&_ERC20Rollup.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupSession) IsStaked(staker common.Address) (bool, error) {
	return _ERC20Rollup.Contract.IsStaked(&_ERC20Rollup.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _ERC20Rollup.Contract.IsStaked(&_ERC20Rollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupCaller) IsZombie(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "isZombie", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupSession) IsZombie(staker common.Address) (bool, error) {
	return _ERC20Rollup.Contract.IsZombie(&_ERC20Rollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_ERC20Rollup *ERC20RollupCallerSession) IsZombie(staker common.Address) (bool, error) {
	return _ERC20Rollup.Contract.IsZombie(&_ERC20Rollup.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) LastStakeBlock() (*big.Int, error) {
	return _ERC20Rollup.Contract.LastStakeBlock(&_ERC20Rollup.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) LastStakeBlock() (*big.Int, error) {
	return _ERC20Rollup.Contract.LastStakeBlock(&_ERC20Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) LatestConfirmed() (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestConfirmed(&_ERC20Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) LatestConfirmed() (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestConfirmed(&_ERC20Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) LatestNodeCreated() (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestNodeCreated(&_ERC20Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestNodeCreated(&_ERC20Rollup.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestStakedNode(&_ERC20Rollup.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.LatestStakedNode(&_ERC20Rollup.CallOpts, staker)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _ERC20Rollup.Contract.MinimumAssertionPeriod(&_ERC20Rollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _ERC20Rollup.Contract.MinimumAssertionPeriod(&_ERC20Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) NodeFactory() (common.Address, error) {
	return _ERC20Rollup.Contract.NodeFactory(&_ERC20Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) NodeFactory() (common.Address, error) {
	return _ERC20Rollup.Contract.NodeFactory(&_ERC20Rollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) Outbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "outbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) Outbox() (common.Address, error) {
	return _ERC20Rollup.Contract.Outbox(&_ERC20Rollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) Outbox() (common.Address, error) {
	return _ERC20Rollup.Contract.Outbox(&_ERC20Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) Owner() (common.Address, error) {
	return _ERC20Rollup.Contract.Owner(&_ERC20Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) Owner() (common.Address, error) {
	return _ERC20Rollup.Contract.Owner(&_ERC20Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Rollup *ERC20RollupCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Rollup *ERC20RollupSession) Paused() (bool, error) {
	return _ERC20Rollup.Contract.Paused(&_ERC20Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Rollup *ERC20RollupCallerSession) Paused() (bool, error) {
	return _ERC20Rollup.Contract.Paused(&_ERC20Rollup.CallOpts)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_ERC20Rollup *ERC20RollupCaller) RequireUnresolved(opts *bind.CallOpts, nodeNum *big.Int) error {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "requireUnresolved", nodeNum)

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_ERC20Rollup *ERC20RollupSession) RequireUnresolved(nodeNum *big.Int) error {
	return _ERC20Rollup.Contract.RequireUnresolved(&_ERC20Rollup.CallOpts, nodeNum)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_ERC20Rollup *ERC20RollupCallerSession) RequireUnresolved(nodeNum *big.Int) error {
	return _ERC20Rollup.Contract.RequireUnresolved(&_ERC20Rollup.CallOpts, nodeNum)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_ERC20Rollup *ERC20RollupCaller) RequireUnresolvedExists(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "requireUnresolvedExists")

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_ERC20Rollup *ERC20RollupSession) RequireUnresolvedExists() error {
	return _ERC20Rollup.Contract.RequireUnresolvedExists(&_ERC20Rollup.CallOpts)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_ERC20Rollup *ERC20RollupCallerSession) RequireUnresolvedExists() error {
	return _ERC20Rollup.Contract.RequireUnresolvedExists(&_ERC20Rollup.CallOpts)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) RequiredStake(opts *bind.CallOpts, blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "requiredStake", blockNumber, firstUnresolvedNodeNum, latestNodeCreated)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _ERC20Rollup.Contract.RequiredStake(&_ERC20Rollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _ERC20Rollup.Contract.RequiredStake(&_ERC20Rollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) RollupEventBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "rollupEventBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) RollupEventBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.RollupEventBridge(&_ERC20Rollup.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) RollupEventBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.RollupEventBridge(&_ERC20Rollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) SequencerBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "sequencerBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) SequencerBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.SequencerBridge(&_ERC20Rollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) SequencerBridge() (common.Address, error) {
	return _ERC20Rollup.Contract.SequencerBridge(&_ERC20Rollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ERC20Rollup *ERC20RollupSession) StakeToken() (common.Address, error) {
	return _ERC20Rollup.Contract.StakeToken(&_ERC20Rollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) StakeToken() (common.Address, error) {
	return _ERC20Rollup.Contract.StakeToken(&_ERC20Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) StakerCount() (*big.Int, error) {
	return _ERC20Rollup.Contract.StakerCount(&_ERC20Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) StakerCount() (*big.Int, error) {
	return _ERC20Rollup.Contract.StakerCount(&_ERC20Rollup.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.WithdrawableFunds(&_ERC20Rollup.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _ERC20Rollup.Contract.WithdrawableFunds(&_ERC20Rollup.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_ERC20Rollup *ERC20RollupSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.ZombieAddress(&_ERC20Rollup.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_ERC20Rollup *ERC20RollupCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _ERC20Rollup.Contract.ZombieAddress(&_ERC20Rollup.CallOpts, zombieNum)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ZombieCount() (*big.Int, error) {
	return _ERC20Rollup.Contract.ZombieCount(&_ERC20Rollup.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ZombieCount() (*big.Int, error) {
	return _ERC20Rollup.Contract.ZombieCount(&_ERC20Rollup.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Rollup.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _ERC20Rollup.Contract.ZombieLatestStakedNode(&_ERC20Rollup.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_ERC20Rollup *ERC20RollupCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _ERC20Rollup.Contract.ZombieLatestStakedNode(&_ERC20Rollup.CallOpts, zombieNum)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupTransactor) AddToDeposit(opts *bind.TransactOpts, stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "addToDeposit", stakerAddress, tokenAmount)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupSession) AddToDeposit(stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.AddToDeposit(&_ERC20Rollup.TransactOpts, stakerAddress, tokenAmount)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0xf3f0a03e.
//
// Solidity: function addToDeposit(address stakerAddress, uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) AddToDeposit(stakerAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.AddToDeposit(&_ERC20Rollup.TransactOpts, stakerAddress, tokenAmount)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_ERC20Rollup *ERC20RollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "completeChallenge", winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_ERC20Rollup *ERC20RollupSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.CompleteChallenge(&_ERC20Rollup.TransactOpts, winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.CompleteChallenge(&_ERC20Rollup.TransactOpts, winningStaker, losingStaker)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_ERC20Rollup *ERC20RollupTransactor) ConfirmNextNode(opts *bind.TransactOpts, beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "confirmNextNode", beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_ERC20Rollup *ERC20RollupSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ConfirmNextNode(&_ERC20Rollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ConfirmNextNode(&_ERC20Rollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_ERC20Rollup *ERC20RollupTransactor) CreateChallenge(opts *bind.TransactOpts, stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "createChallenge", stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_ERC20Rollup *ERC20RollupSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.CreateChallenge(&_ERC20Rollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.CreateChallenge(&_ERC20Rollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_ERC20Rollup *ERC20RollupTransactor) Initialize(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "initialize", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_ERC20Rollup *ERC20RollupSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Initialize(&_ERC20Rollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Initialize(&_ERC20Rollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupTransactor) NewStake(opts *bind.TransactOpts, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "newStake", tokenAmount)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupSession) NewStake(tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.NewStake(&_ERC20Rollup.TransactOpts, tokenAmount)
}

// NewStake is a paid mutator transaction binding the contract method 0xfb64884e.
//
// Solidity: function newStake(uint256 tokenAmount) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) NewStake(tokenAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.NewStake(&_ERC20Rollup.TransactOpts, tokenAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Rollup *ERC20RollupTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Rollup *ERC20RollupSession) Pause() (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Pause(&_ERC20Rollup.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Pause(&_ERC20Rollup.TransactOpts)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_ERC20Rollup *ERC20RollupTransactor) ReduceDeposit(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "reduceDeposit", target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_ERC20Rollup *ERC20RollupSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ReduceDeposit(&_ERC20Rollup.TransactOpts, target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ReduceDeposit(&_ERC20Rollup.TransactOpts, target)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupTransactor) RejectNextNode(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "rejectNextNode", stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RejectNextNode(&_ERC20Rollup.TransactOpts, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RejectNextNode(&_ERC20Rollup.TransactOpts, stakerAddress)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupTransactor) RemoveOldOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "removeOldOutbox", _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveOldOutbox(&_ERC20Rollup.TransactOpts, _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveOldOutbox(&_ERC20Rollup.TransactOpts, _outbox)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_ERC20Rollup *ERC20RollupTransactor) RemoveOldZombies(opts *bind.TransactOpts, startIndex *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "removeOldZombies", startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_ERC20Rollup *ERC20RollupSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveOldZombies(&_ERC20Rollup.TransactOpts, startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveOldZombies(&_ERC20Rollup.TransactOpts, startIndex)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_ERC20Rollup *ERC20RollupTransactor) RemoveZombie(opts *bind.TransactOpts, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "removeZombie", zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_ERC20Rollup *ERC20RollupSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveZombie(&_ERC20Rollup.TransactOpts, zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.RemoveZombie(&_ERC20Rollup.TransactOpts, zombieNum, maxNodes)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ERC20Rollup *ERC20RollupTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ERC20Rollup *ERC20RollupSession) Resume() (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Resume(&_ERC20Rollup.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) Resume() (*types.Transaction, error) {
	return _ERC20Rollup.Contract.Resume(&_ERC20Rollup.TransactOpts)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupTransactor) ReturnOldDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "returnOldDeposit", stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ReturnOldDeposit(&_ERC20Rollup.TransactOpts, stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.ReturnOldDeposit(&_ERC20Rollup.TransactOpts, stakerAddress)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_ERC20Rollup *ERC20RollupTransactor) SetInbox(opts *bind.TransactOpts, _inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "setInbox", _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_ERC20Rollup *ERC20RollupSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.SetInbox(&_ERC20Rollup.TransactOpts, _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.SetInbox(&_ERC20Rollup.TransactOpts, _inbox, _enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupTransactor) SetOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "setOutbox", _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.SetOutbox(&_ERC20Rollup.TransactOpts, _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.SetOutbox(&_ERC20Rollup.TransactOpts, _outbox)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_ERC20Rollup *ERC20RollupTransactor) StakeOnExistingNode(opts *bind.TransactOpts, nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "stakeOnExistingNode", nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_ERC20Rollup *ERC20RollupSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.StakeOnExistingNode(&_ERC20Rollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.StakeOnExistingNode(&_ERC20Rollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_ERC20Rollup *ERC20RollupTransactor) StakeOnNewNode(opts *bind.TransactOpts, expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "stakeOnNewNode", expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_ERC20Rollup *ERC20RollupSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.StakeOnNewNode(&_ERC20Rollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.StakeOnNewNode(&_ERC20Rollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_ERC20Rollup *ERC20RollupTransactor) UpgradeImplementation(opts *bind.TransactOpts, _newRollup common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "upgradeImplementation", _newRollup)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_ERC20Rollup *ERC20RollupSession) UpgradeImplementation(_newRollup common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.UpgradeImplementation(&_ERC20Rollup.TransactOpts, _newRollup)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) UpgradeImplementation(_newRollup common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.UpgradeImplementation(&_ERC20Rollup.TransactOpts, _newRollup)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_ERC20Rollup *ERC20RollupTransactor) UpgradeImplementationAndCall(opts *bind.TransactOpts, _newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "upgradeImplementationAndCall", _newRollup, _data)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_ERC20Rollup *ERC20RollupSession) UpgradeImplementationAndCall(_newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.UpgradeImplementationAndCall(&_ERC20Rollup.TransactOpts, _newRollup, _data)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_ERC20Rollup *ERC20RollupTransactorSession) UpgradeImplementationAndCall(_newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.UpgradeImplementationAndCall(&_ERC20Rollup.TransactOpts, _newRollup, _data)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_ERC20Rollup *ERC20RollupTransactor) WithdrawStakerFunds(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.contract.Transact(opts, "withdrawStakerFunds", destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_ERC20Rollup *ERC20RollupSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.WithdrawStakerFunds(&_ERC20Rollup.TransactOpts, destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_ERC20Rollup *ERC20RollupTransactorSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _ERC20Rollup.Contract.WithdrawStakerFunds(&_ERC20Rollup.TransactOpts, destination)
}

// ERC20RollupNodeConfirmedIterator is returned from FilterNodeConfirmed and is used to iterate over the raw logs and unpacked data for NodeConfirmed events raised by the ERC20Rollup contract.
type ERC20RollupNodeConfirmedIterator struct {
	Event *ERC20RollupNodeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20RollupNodeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupNodeConfirmed)
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
		it.Event = new(ERC20RollupNodeConfirmed)
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
func (it *ERC20RollupNodeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupNodeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupNodeConfirmed represents a NodeConfirmed event raised by the ERC20Rollup contract.
type ERC20RollupNodeConfirmed struct {
	NodeNum        *big.Int
	AfterSendAcc   [32]byte
	AfterSendCount *big.Int
	AfterLogAcc    [32]byte
	AfterLogCount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeConfirmed is a free log retrieval operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_ERC20Rollup *ERC20RollupFilterer) FilterNodeConfirmed(opts *bind.FilterOpts, nodeNum []*big.Int) (*ERC20RollupNodeConfirmedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupNodeConfirmedIterator{contract: _ERC20Rollup.contract, event: "NodeConfirmed", logs: logs, sub: sub}, nil
}

// WatchNodeConfirmed is a free log subscription operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_ERC20Rollup *ERC20RollupFilterer) WatchNodeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20RollupNodeConfirmed, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupNodeConfirmed)
				if err := _ERC20Rollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
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

// ParseNodeConfirmed is a log parse operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_ERC20Rollup *ERC20RollupFilterer) ParseNodeConfirmed(log types.Log) (*ERC20RollupNodeConfirmed, error) {
	event := new(ERC20RollupNodeConfirmed)
	if err := _ERC20Rollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the ERC20Rollup contract.
type ERC20RollupNodeCreatedIterator struct {
	Event *ERC20RollupNodeCreated // Event containing the contract specifics and raw log

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
func (it *ERC20RollupNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupNodeCreated)
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
		it.Event = new(ERC20RollupNodeCreated)
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
func (it *ERC20RollupNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupNodeCreated represents a NodeCreated event raised by the ERC20Rollup contract.
type ERC20RollupNodeCreated struct {
	NodeNum                 *big.Int
	ParentNodeHash          [32]byte
	NodeHash                [32]byte
	ExecutionHash           [32]byte
	InboxMaxCount           *big.Int
	AfterInboxBatchEndCount *big.Int
	AfterInboxBatchAcc      [32]byte
	AssertionBytes32Fields  [2][3][32]byte
	AssertionIntFields      [2][4]*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_ERC20Rollup *ERC20RollupFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int, parentNodeHash [][32]byte) (*ERC20RollupNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupNodeCreatedIterator{contract: _ERC20Rollup.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_ERC20Rollup *ERC20RollupFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *ERC20RollupNodeCreated, nodeNum []*big.Int, parentNodeHash [][32]byte) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupNodeCreated)
				if err := _ERC20Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
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

// ParseNodeCreated is a log parse operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_ERC20Rollup *ERC20RollupFilterer) ParseNodeCreated(log types.Log) (*ERC20RollupNodeCreated, error) {
	event := new(ERC20RollupNodeCreated)
	if err := _ERC20Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupNodeRejectedIterator is returned from FilterNodeRejected and is used to iterate over the raw logs and unpacked data for NodeRejected events raised by the ERC20Rollup contract.
type ERC20RollupNodeRejectedIterator struct {
	Event *ERC20RollupNodeRejected // Event containing the contract specifics and raw log

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
func (it *ERC20RollupNodeRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupNodeRejected)
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
		it.Event = new(ERC20RollupNodeRejected)
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
func (it *ERC20RollupNodeRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupNodeRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupNodeRejected represents a NodeRejected event raised by the ERC20Rollup contract.
type ERC20RollupNodeRejected struct {
	NodeNum *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeRejected is a free log retrieval operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_ERC20Rollup *ERC20RollupFilterer) FilterNodeRejected(opts *bind.FilterOpts, nodeNum []*big.Int) (*ERC20RollupNodeRejectedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupNodeRejectedIterator{contract: _ERC20Rollup.contract, event: "NodeRejected", logs: logs, sub: sub}, nil
}

// WatchNodeRejected is a free log subscription operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_ERC20Rollup *ERC20RollupFilterer) WatchNodeRejected(opts *bind.WatchOpts, sink chan<- *ERC20RollupNodeRejected, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupNodeRejected)
				if err := _ERC20Rollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
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

// ParseNodeRejected is a log parse operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_ERC20Rollup *ERC20RollupFilterer) ParseNodeRejected(log types.Log) (*ERC20RollupNodeRejected, error) {
	event := new(ERC20RollupNodeRejected)
	if err := _ERC20Rollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupNodesDestroyedIterator is returned from FilterNodesDestroyed and is used to iterate over the raw logs and unpacked data for NodesDestroyed events raised by the ERC20Rollup contract.
type ERC20RollupNodesDestroyedIterator struct {
	Event *ERC20RollupNodesDestroyed // Event containing the contract specifics and raw log

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
func (it *ERC20RollupNodesDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupNodesDestroyed)
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
		it.Event = new(ERC20RollupNodesDestroyed)
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
func (it *ERC20RollupNodesDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupNodesDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupNodesDestroyed represents a NodesDestroyed event raised by the ERC20Rollup contract.
type ERC20RollupNodesDestroyed struct {
	StartNode *big.Int
	EndNode   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodesDestroyed is a free log retrieval operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_ERC20Rollup *ERC20RollupFilterer) FilterNodesDestroyed(opts *bind.FilterOpts, startNode []*big.Int, endNode []*big.Int) (*ERC20RollupNodesDestroyedIterator, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupNodesDestroyedIterator{contract: _ERC20Rollup.contract, event: "NodesDestroyed", logs: logs, sub: sub}, nil
}

// WatchNodesDestroyed is a free log subscription operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_ERC20Rollup *ERC20RollupFilterer) WatchNodesDestroyed(opts *bind.WatchOpts, sink chan<- *ERC20RollupNodesDestroyed, startNode []*big.Int, endNode []*big.Int) (event.Subscription, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupNodesDestroyed)
				if err := _ERC20Rollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
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

// ParseNodesDestroyed is a log parse operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_ERC20Rollup *ERC20RollupFilterer) ParseNodesDestroyed(log types.Log) (*ERC20RollupNodesDestroyed, error) {
	event := new(ERC20RollupNodesDestroyed)
	if err := _ERC20Rollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20Rollup contract.
type ERC20RollupPausedIterator struct {
	Event *ERC20RollupPaused // Event containing the contract specifics and raw log

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
func (it *ERC20RollupPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupPaused)
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
		it.Event = new(ERC20RollupPaused)
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
func (it *ERC20RollupPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupPaused represents a Paused event raised by the ERC20Rollup contract.
type ERC20RollupPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20RollupPausedIterator, error) {

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20RollupPausedIterator{contract: _ERC20Rollup.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20RollupPaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupPaused)
				if err := _ERC20Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) ParsePaused(log types.Log) (*ERC20RollupPaused, error) {
	event := new(ERC20RollupPaused)
	if err := _ERC20Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the ERC20Rollup contract.
type ERC20RollupRollupChallengeStartedIterator struct {
	Event *ERC20RollupRollupChallengeStarted // Event containing the contract specifics and raw log

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
func (it *ERC20RollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupRollupChallengeStarted)
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
		it.Event = new(ERC20RollupRollupChallengeStarted)
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
func (it *ERC20RollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the ERC20Rollup contract.
type ERC20RollupRollupChallengeStarted struct {
	ChallengeContract common.Address
	Asserter          common.Address
	Challenger        common.Address
	ChallengedNode    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_ERC20Rollup *ERC20RollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts, challengeContract []common.Address) (*ERC20RollupRollupChallengeStartedIterator, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupRollupChallengeStartedIterator{contract: _ERC20Rollup.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_ERC20Rollup *ERC20RollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *ERC20RollupRollupChallengeStarted, challengeContract []common.Address) (event.Subscription, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupRollupChallengeStarted)
				if err := _ERC20Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
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

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_ERC20Rollup *ERC20RollupFilterer) ParseRollupChallengeStarted(log types.Log) (*ERC20RollupRollupChallengeStarted, error) {
	event := new(ERC20RollupRollupChallengeStarted)
	if err := _ERC20Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the ERC20Rollup contract.
type ERC20RollupRollupCreatedIterator struct {
	Event *ERC20RollupRollupCreated // Event containing the contract specifics and raw log

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
func (it *ERC20RollupRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupRollupCreated)
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
		it.Event = new(ERC20RollupRollupCreated)
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
func (it *ERC20RollupRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupRollupCreated represents a RollupCreated event raised by the ERC20Rollup contract.
type ERC20RollupRollupCreated struct {
	MachineHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_ERC20Rollup *ERC20RollupFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*ERC20RollupRollupCreatedIterator, error) {

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &ERC20RollupRollupCreatedIterator{contract: _ERC20Rollup.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_ERC20Rollup *ERC20RollupFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *ERC20RollupRollupCreated) (event.Subscription, error) {

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupRollupCreated)
				if err := _ERC20Rollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_ERC20Rollup *ERC20RollupFilterer) ParseRollupCreated(log types.Log) (*ERC20RollupRollupCreated, error) {
	event := new(ERC20RollupRollupCreated)
	if err := _ERC20Rollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupStakerReassignedIterator is returned from FilterStakerReassigned and is used to iterate over the raw logs and unpacked data for StakerReassigned events raised by the ERC20Rollup contract.
type ERC20RollupStakerReassignedIterator struct {
	Event *ERC20RollupStakerReassigned // Event containing the contract specifics and raw log

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
func (it *ERC20RollupStakerReassignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupStakerReassigned)
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
		it.Event = new(ERC20RollupStakerReassigned)
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
func (it *ERC20RollupStakerReassignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupStakerReassignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupStakerReassigned represents a StakerReassigned event raised by the ERC20Rollup contract.
type ERC20RollupStakerReassigned struct {
	Staker  common.Address
	NewNode *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakerReassigned is a free log retrieval operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_ERC20Rollup *ERC20RollupFilterer) FilterStakerReassigned(opts *bind.FilterOpts, staker []common.Address) (*ERC20RollupStakerReassignedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RollupStakerReassignedIterator{contract: _ERC20Rollup.contract, event: "StakerReassigned", logs: logs, sub: sub}, nil
}

// WatchStakerReassigned is a free log subscription operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_ERC20Rollup *ERC20RollupFilterer) WatchStakerReassigned(opts *bind.WatchOpts, sink chan<- *ERC20RollupStakerReassigned, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupStakerReassigned)
				if err := _ERC20Rollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
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

// ParseStakerReassigned is a log parse operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_ERC20Rollup *ERC20RollupFilterer) ParseStakerReassigned(log types.Log) (*ERC20RollupStakerReassigned, error) {
	event := new(ERC20RollupStakerReassigned)
	if err := _ERC20Rollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RollupUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20Rollup contract.
type ERC20RollupUnpausedIterator struct {
	Event *ERC20RollupUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20RollupUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RollupUnpaused)
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
		it.Event = new(ERC20RollupUnpaused)
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
func (it *ERC20RollupUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RollupUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RollupUnpaused represents a Unpaused event raised by the ERC20Rollup contract.
type ERC20RollupUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20RollupUnpausedIterator, error) {

	logs, sub, err := _ERC20Rollup.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20RollupUnpausedIterator{contract: _ERC20Rollup.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20RollupUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Rollup.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RollupUnpaused)
				if err := _ERC20Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Rollup *ERC20RollupFilterer) ParseUnpaused(log types.Log) (*ERC20RollupUnpaused, error) {
	event := new(ERC20RollupUnpaused)
	if err := _ERC20Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INodeABI is the input ABI used to generate the binding from.
const INodeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"addStaker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstChildBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_confirmData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadlineBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestChildNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"newChildConfirmDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noChildConfirmedBeforeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"removeStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requirePastChildConfirmDeadline\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requirePastDeadline\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetChildren\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// INode is an auto generated Go binding around an Ethereum contract.
type INode struct {
	INodeCaller     // Read-only binding to the contract
	INodeTransactor // Write-only binding to the contract
	INodeFilterer   // Log filterer for contract events
}

// INodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type INodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INodeSession struct {
	Contract     *INode            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INodeCallerSession struct {
	Contract *INodeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// INodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INodeTransactorSession struct {
	Contract     *INodeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type INodeRaw struct {
	Contract *INode // Generic contract binding to access the raw methods on
}

// INodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INodeCallerRaw struct {
	Contract *INodeCaller // Generic read-only contract binding to access the raw methods on
}

// INodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INodeTransactorRaw struct {
	Contract *INodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINode creates a new instance of INode, bound to a specific deployed contract.
func NewINode(address common.Address, backend bind.ContractBackend) (*INode, error) {
	contract, err := bindINode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INode{INodeCaller: INodeCaller{contract: contract}, INodeTransactor: INodeTransactor{contract: contract}, INodeFilterer: INodeFilterer{contract: contract}}, nil
}

// NewINodeCaller creates a new read-only instance of INode, bound to a specific deployed contract.
func NewINodeCaller(address common.Address, caller bind.ContractCaller) (*INodeCaller, error) {
	contract, err := bindINode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INodeCaller{contract: contract}, nil
}

// NewINodeTransactor creates a new write-only instance of INode, bound to a specific deployed contract.
func NewINodeTransactor(address common.Address, transactor bind.ContractTransactor) (*INodeTransactor, error) {
	contract, err := bindINode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INodeTransactor{contract: contract}, nil
}

// NewINodeFilterer creates a new log filterer instance of INode, bound to a specific deployed contract.
func NewINodeFilterer(address common.Address, filterer bind.ContractFilterer) (*INodeFilterer, error) {
	contract, err := bindINode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INodeFilterer{contract: contract}, nil
}

// bindINode binds a generic wrapper to an already deployed contract.
func bindINode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INode.Contract.INodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.contract.Transact(opts, method, params...)
}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_INode *INodeCaller) ChallengeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "challengeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_INode *INodeSession) ChallengeHash() ([32]byte, error) {
	return _INode.Contract.ChallengeHash(&_INode.CallOpts)
}

// ChallengeHash is a free data retrieval call binding the contract method 0x5b8b2280.
//
// Solidity: function challengeHash() view returns(bytes32)
func (_INode *INodeCallerSession) ChallengeHash() ([32]byte, error) {
	return _INode.Contract.ChallengeHash(&_INode.CallOpts)
}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_INode *INodeCaller) ConfirmData(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "confirmData")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_INode *INodeSession) ConfirmData() ([32]byte, error) {
	return _INode.Contract.ConfirmData(&_INode.CallOpts)
}

// ConfirmData is a free data retrieval call binding the contract method 0x97bdc510.
//
// Solidity: function confirmData() view returns(bytes32)
func (_INode *INodeCallerSession) ConfirmData() ([32]byte, error) {
	return _INode.Contract.ConfirmData(&_INode.CallOpts)
}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_INode *INodeCaller) DeadlineBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "deadlineBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_INode *INodeSession) DeadlineBlock() (*big.Int, error) {
	return _INode.Contract.DeadlineBlock(&_INode.CallOpts)
}

// DeadlineBlock is a free data retrieval call binding the contract method 0x2edfb42a.
//
// Solidity: function deadlineBlock() view returns(uint256)
func (_INode *INodeCallerSession) DeadlineBlock() (*big.Int, error) {
	return _INode.Contract.DeadlineBlock(&_INode.CallOpts)
}

// FirstChildBlock is a free data retrieval call binding the contract method 0xd7ff5e35.
//
// Solidity: function firstChildBlock() view returns(uint256)
func (_INode *INodeCaller) FirstChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "firstChildBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstChildBlock is a free data retrieval call binding the contract method 0xd7ff5e35.
//
// Solidity: function firstChildBlock() view returns(uint256)
func (_INode *INodeSession) FirstChildBlock() (*big.Int, error) {
	return _INode.Contract.FirstChildBlock(&_INode.CallOpts)
}

// FirstChildBlock is a free data retrieval call binding the contract method 0xd7ff5e35.
//
// Solidity: function firstChildBlock() view returns(uint256)
func (_INode *INodeCallerSession) FirstChildBlock() (*big.Int, error) {
	return _INode.Contract.FirstChildBlock(&_INode.CallOpts)
}

// LatestChildNumber is a free data retrieval call binding the contract method 0xf0dd77ff.
//
// Solidity: function latestChildNumber() view returns(uint256)
func (_INode *INodeCaller) LatestChildNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "latestChildNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestChildNumber is a free data retrieval call binding the contract method 0xf0dd77ff.
//
// Solidity: function latestChildNumber() view returns(uint256)
func (_INode *INodeSession) LatestChildNumber() (*big.Int, error) {
	return _INode.Contract.LatestChildNumber(&_INode.CallOpts)
}

// LatestChildNumber is a free data retrieval call binding the contract method 0xf0dd77ff.
//
// Solidity: function latestChildNumber() view returns(uint256)
func (_INode *INodeCallerSession) LatestChildNumber() (*big.Int, error) {
	return _INode.Contract.LatestChildNumber(&_INode.CallOpts)
}

// NoChildConfirmedBeforeBlock is a free data retrieval call binding the contract method 0xa0369c14.
//
// Solidity: function noChildConfirmedBeforeBlock() view returns(uint256)
func (_INode *INodeCaller) NoChildConfirmedBeforeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "noChildConfirmedBeforeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoChildConfirmedBeforeBlock is a free data retrieval call binding the contract method 0xa0369c14.
//
// Solidity: function noChildConfirmedBeforeBlock() view returns(uint256)
func (_INode *INodeSession) NoChildConfirmedBeforeBlock() (*big.Int, error) {
	return _INode.Contract.NoChildConfirmedBeforeBlock(&_INode.CallOpts)
}

// NoChildConfirmedBeforeBlock is a free data retrieval call binding the contract method 0xa0369c14.
//
// Solidity: function noChildConfirmedBeforeBlock() view returns(uint256)
func (_INode *INodeCallerSession) NoChildConfirmedBeforeBlock() (*big.Int, error) {
	return _INode.Contract.NoChildConfirmedBeforeBlock(&_INode.CallOpts)
}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_INode *INodeCaller) Prev(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "prev")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_INode *INodeSession) Prev() (*big.Int, error) {
	return _INode.Contract.Prev(&_INode.CallOpts)
}

// Prev is a free data retrieval call binding the contract method 0x479c9254.
//
// Solidity: function prev() view returns(uint256)
func (_INode *INodeCallerSession) Prev() (*big.Int, error) {
	return _INode.Contract.Prev(&_INode.CallOpts)
}

// RequirePastChildConfirmDeadline is a free data retrieval call binding the contract method 0x3aa19274.
//
// Solidity: function requirePastChildConfirmDeadline() view returns()
func (_INode *INodeCaller) RequirePastChildConfirmDeadline(opts *bind.CallOpts) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "requirePastChildConfirmDeadline")

	if err != nil {
		return err
	}

	return err

}

// RequirePastChildConfirmDeadline is a free data retrieval call binding the contract method 0x3aa19274.
//
// Solidity: function requirePastChildConfirmDeadline() view returns()
func (_INode *INodeSession) RequirePastChildConfirmDeadline() error {
	return _INode.Contract.RequirePastChildConfirmDeadline(&_INode.CallOpts)
}

// RequirePastChildConfirmDeadline is a free data retrieval call binding the contract method 0x3aa19274.
//
// Solidity: function requirePastChildConfirmDeadline() view returns()
func (_INode *INodeCallerSession) RequirePastChildConfirmDeadline() error {
	return _INode.Contract.RequirePastChildConfirmDeadline(&_INode.CallOpts)
}

// RequirePastDeadline is a free data retrieval call binding the contract method 0x88d221c6.
//
// Solidity: function requirePastDeadline() view returns()
func (_INode *INodeCaller) RequirePastDeadline(opts *bind.CallOpts) error {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "requirePastDeadline")

	if err != nil {
		return err
	}

	return err

}

// RequirePastDeadline is a free data retrieval call binding the contract method 0x88d221c6.
//
// Solidity: function requirePastDeadline() view returns()
func (_INode *INodeSession) RequirePastDeadline() error {
	return _INode.Contract.RequirePastDeadline(&_INode.CallOpts)
}

// RequirePastDeadline is a free data retrieval call binding the contract method 0x88d221c6.
//
// Solidity: function requirePastDeadline() view returns()
func (_INode *INodeCallerSession) RequirePastDeadline() error {
	return _INode.Contract.RequirePastDeadline(&_INode.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_INode *INodeCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_INode *INodeSession) StakerCount() (*big.Int, error) {
	return _INode.Contract.StakerCount(&_INode.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_INode *INodeCallerSession) StakerCount() (*big.Int, error) {
	return _INode.Contract.StakerCount(&_INode.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(bool)
func (_INode *INodeCaller) Stakers(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "stakers", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(bool)
func (_INode *INodeSession) Stakers(staker common.Address) (bool, error) {
	return _INode.Contract.Stakers(&_INode.CallOpts, staker)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(bool)
func (_INode *INodeCallerSession) Stakers(staker common.Address) (bool, error) {
	return _INode.Contract.Stakers(&_INode.CallOpts, staker)
}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_INode *INodeCaller) StateHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "stateHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_INode *INodeSession) StateHash() ([32]byte, error) {
	return _INode.Contract.StateHash(&_INode.CallOpts)
}

// StateHash is a free data retrieval call binding the contract method 0x701da98e.
//
// Solidity: function stateHash() view returns(bytes32)
func (_INode *INodeCallerSession) StateHash() ([32]byte, error) {
	return _INode.Contract.StateHash(&_INode.CallOpts)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns(uint256)
func (_INode *INodeTransactor) AddStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "addStaker", staker)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns(uint256)
func (_INode *INodeSession) AddStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.AddStaker(&_INode.TransactOpts, staker)
}

// AddStaker is a paid mutator transaction binding the contract method 0x2466696e.
//
// Solidity: function addStaker(address staker) returns(uint256)
func (_INode *INodeTransactorSession) AddStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.AddStaker(&_INode.TransactOpts, staker)
}

// ChildCreated is a paid mutator transaction binding the contract method 0x1bc09d0a.
//
// Solidity: function childCreated(uint256 ) returns()
func (_INode *INodeTransactor) ChildCreated(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "childCreated", arg0)
}

// ChildCreated is a paid mutator transaction binding the contract method 0x1bc09d0a.
//
// Solidity: function childCreated(uint256 ) returns()
func (_INode *INodeSession) ChildCreated(arg0 *big.Int) (*types.Transaction, error) {
	return _INode.Contract.ChildCreated(&_INode.TransactOpts, arg0)
}

// ChildCreated is a paid mutator transaction binding the contract method 0x1bc09d0a.
//
// Solidity: function childCreated(uint256 ) returns()
func (_INode *INodeTransactorSession) ChildCreated(arg0 *big.Int) (*types.Transaction, error) {
	return _INode.Contract.ChildCreated(&_INode.TransactOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_INode *INodeTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_INode *INodeSession) Destroy() (*types.Transaction, error) {
	return _INode.Contract.Destroy(&_INode.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_INode *INodeTransactorSession) Destroy() (*types.Transaction, error) {
	return _INode.Contract.Destroy(&_INode.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_INode *INodeTransactor) Initialize(opts *bind.TransactOpts, _rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "initialize", _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_INode *INodeSession) Initialize(_rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Initialize(&_INode.TransactOpts, _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xa406b374.
//
// Solidity: function initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock) returns()
func (_INode *INodeTransactorSession) Initialize(_rollup common.Address, _stateHash [32]byte, _challengeHash [32]byte, _confirmData [32]byte, _prev *big.Int, _deadlineBlock *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Initialize(&_INode.TransactOpts, _rollup, _stateHash, _challengeHash, _confirmData, _prev, _deadlineBlock)
}

// NewChildConfirmDeadline is a paid mutator transaction binding the contract method 0x6971dfe5.
//
// Solidity: function newChildConfirmDeadline(uint256 deadline) returns()
func (_INode *INodeTransactor) NewChildConfirmDeadline(opts *bind.TransactOpts, deadline *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "newChildConfirmDeadline", deadline)
}

// NewChildConfirmDeadline is a paid mutator transaction binding the contract method 0x6971dfe5.
//
// Solidity: function newChildConfirmDeadline(uint256 deadline) returns()
func (_INode *INodeSession) NewChildConfirmDeadline(deadline *big.Int) (*types.Transaction, error) {
	return _INode.Contract.NewChildConfirmDeadline(&_INode.TransactOpts, deadline)
}

// NewChildConfirmDeadline is a paid mutator transaction binding the contract method 0x6971dfe5.
//
// Solidity: function newChildConfirmDeadline(uint256 deadline) returns()
func (_INode *INodeTransactorSession) NewChildConfirmDeadline(deadline *big.Int) (*types.Transaction, error) {
	return _INode.Contract.NewChildConfirmDeadline(&_INode.TransactOpts, deadline)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_INode *INodeTransactor) RemoveStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "removeStaker", staker)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_INode *INodeSession) RemoveStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.RemoveStaker(&_INode.TransactOpts, staker)
}

// RemoveStaker is a paid mutator transaction binding the contract method 0x96a9fdc0.
//
// Solidity: function removeStaker(address staker) returns()
func (_INode *INodeTransactorSession) RemoveStaker(staker common.Address) (*types.Transaction, error) {
	return _INode.Contract.RemoveStaker(&_INode.TransactOpts, staker)
}

// ResetChildren is a paid mutator transaction binding the contract method 0xda83b48e.
//
// Solidity: function resetChildren() returns()
func (_INode *INodeTransactor) ResetChildren(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "resetChildren")
}

// ResetChildren is a paid mutator transaction binding the contract method 0xda83b48e.
//
// Solidity: function resetChildren() returns()
func (_INode *INodeSession) ResetChildren() (*types.Transaction, error) {
	return _INode.Contract.ResetChildren(&_INode.TransactOpts)
}

// ResetChildren is a paid mutator transaction binding the contract method 0xda83b48e.
//
// Solidity: function resetChildren() returns()
func (_INode *INodeTransactorSession) ResetChildren() (*types.Transaction, error) {
	return _INode.Contract.ResetChildren(&_INode.TransactOpts)
}

// ProxyAdminABI is the input ABI used to generate the binding from.
const ProxyAdminABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ProxyAdminBin is the compiled bytecode used for deploying new contracts.
var ProxyAdminBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6108658061007d6000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100bf5780637eff275e146100d65780638da5cb5b146101115780639623609d1461012657806399a88ec4146101e3578063f2fde38b1461021e578063f3b7dead14610251575b600080fd5b34801561007c57600080fd5b506100a36004803603602081101561009357600080fd5b50356001600160a01b0316610284565b604080516001600160a01b039092168252519081900360200190f35b3480156100cb57600080fd5b506100d4610316565b005b3480156100e257600080fd5b506100d4600480360360408110156100f957600080fd5b506001600160a01b03813581169160200135166103b0565b34801561011d57600080fd5b506100a361047d565b6100d46004803603606081101561013c57600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460018302840111600160201b831117156101a257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061048c945050505050565b3480156101ef57600080fd5b506100d46004803603604081101561020657600080fd5b506001600160a01b03813581169160200135166105c5565b34801561022a57600080fd5b506100d46004803603602081101561024157600080fd5b50356001600160a01b0316610676565b34801561025d57600080fd5b506100a36004803603602081101561027457600080fd5b50356001600160a01b0316610766565b6000806060836001600160a01b03166040518080635c60da1b60e01b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b606091505b5091509150816102f757600080fd5b80806020019051602081101561030c57600080fd5b5051949350505050565b61031e6107c5565b6001600160a01b031661032f61047d565b6001600160a01b031614610378576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610810833981519152908390a3600080546001600160a01b0319169055565b6103b86107c5565b6001600160a01b03166103c961047d565b6001600160a01b031614610412576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316638f283970826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b505af1158015610475573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b6104946107c5565b6001600160a01b03166104a561047d565b6001600160a01b0316146104ee576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b826001600160a01b0316634f1ef2863484846040518463ffffffff1660e01b815260040180836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561055b578181015183820152602001610543565b50505050905090810190601f1680156105885780820380516001836020036101000a031916815260200191505b5093505050506000604051808303818588803b1580156105a757600080fd5b505af11580156105bb573d6000803e3d6000fd5b5050505050505050565b6105cd6107c5565b6001600160a01b03166105de61047d565b6001600160a01b031614610627576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316633659cfe6826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b61067e6107c5565b6001600160a01b031661068f61047d565b6001600160a01b0316146106d8576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b6001600160a01b03811661071d5760405162461bcd60e51b81526004018080602001828103825260268152602001806107ca6026913960400191505060405180910390fd5b600080546040516001600160a01b038085169392169160008051602061081083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806060836001600160a01b031660405180806303e1469160e61b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122036b5070566a1002642d9947563bea5cf538675280563c6ec32ab459f386727a264736f6c634300060c0033"

// DeployProxyAdmin deploys a new Ethereum contract, binding an instance of ProxyAdmin to it.
func DeployProxyAdmin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProxyAdmin, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAdminABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyAdminBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// ProxyAdmin is an auto generated Go binding around an Ethereum contract.
type ProxyAdmin struct {
	ProxyAdminCaller     // Read-only binding to the contract
	ProxyAdminTransactor // Write-only binding to the contract
	ProxyAdminFilterer   // Log filterer for contract events
}

// ProxyAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyAdminSession struct {
	Contract     *ProxyAdmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyAdminCallerSession struct {
	Contract *ProxyAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProxyAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyAdminTransactorSession struct {
	Contract     *ProxyAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProxyAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyAdminRaw struct {
	Contract *ProxyAdmin // Generic contract binding to access the raw methods on
}

// ProxyAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyAdminCallerRaw struct {
	Contract *ProxyAdminCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyAdminTransactorRaw struct {
	Contract *ProxyAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyAdmin creates a new instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdmin(address common.Address, backend bind.ContractBackend) (*ProxyAdmin, error) {
	contract, err := bindProxyAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// NewProxyAdminCaller creates a new read-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminCaller(address common.Address, caller bind.ContractCaller) (*ProxyAdminCaller, error) {
	contract, err := bindProxyAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminCaller{contract: contract}, nil
}

// NewProxyAdminTransactor creates a new write-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyAdminTransactor, error) {
	contract, err := bindProxyAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminTransactor{contract: contract}, nil
}

// NewProxyAdminFilterer creates a new log filterer instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyAdminFilterer, error) {
	contract, err := bindProxyAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminFilterer{contract: contract}, nil
}

// bindProxyAdmin binds a generic wrapper to an already deployed contract.
func bindProxyAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.ProxyAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transact(opts, method, params...)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyAdmin(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyAdmin", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyImplementation(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyImplementation", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, proxy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "changeProxyAdmin", proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, proxy, newAdmin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactor) Upgrade(opts *bind.TransactOpts, proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgrade", proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, proxy, implementation)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactor) UpgradeAndCall(opts *bind.TransactOpts, proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgradeAndCall", proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, proxy, implementation, data)
}

// ProxyAdminOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferredIterator struct {
	Event *ProxyAdminOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProxyAdminOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAdminOwnershipTransferred)
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
		it.Event = new(ProxyAdminOwnershipTransferred)
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
func (it *ProxyAdminOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAdminOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAdminOwnershipTransferred represents a OwnershipTransferred event raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProxyAdminOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminOwnershipTransferredIterator{contract: _ProxyAdmin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyAdminOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAdminOwnershipTransferred)
				if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyAdminOwnershipTransferred, error) {
	event := new(ProxyAdminOwnershipTransferred)
	if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterSendAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"NodeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentNodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterInboxBatchEndCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterInboxBatchAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"indexed\":false,\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"NodeRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startNode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endNode\",\"type\":\"uint256\"}],\"name\":\"NodesDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNode\",\"type\":\"uint256\"}],\"name\":\"StakerReassigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"addToDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningStaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"losingStaker\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeSendAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"confirmNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"countStakedZombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[2]\",\"name\":\"stakers\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"nodeNums\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"executionHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proposedTimes\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"maxMessageCounts\",\"type\":\"uint256[2]\"}],\"name\":\"createChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedBridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraChallengeTimeBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[7]\",\"name\":\"connectedContracts\",\"type\":\"address[7]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"reduceDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"rejectNextNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodes\",\"type\":\"uint256\"}],\"name\":\"removeZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"requireUnresolved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireUnresolvedExists\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstUnresolvedNodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestNodeCreated\",\"type\":\"uint256\"}],\"name\":\"requiredStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"returnOldDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerBridge\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"stakeOnExistingNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expectedNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"},{\"internalType\":\"uint256\",\"name\":\"beforeProposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxMaxCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sequencerBatchProof\",\"type\":\"bytes\"}],\"name\":\"stakeOnNewNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"}],\"name\":\"upgradeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRollup\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeImplementationAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupBin is the compiled bytecode used for deploying new contracts.
var RollupBin = "0x608060405234801561001057600080fd5b506000805460ff19908116600117909155600b80549091169055615731806200003a6000396000f3fe60806040526004361061029c5760003560e01c8063046f7da2146102a157806304a28064146102b85780631e83d30f146102fd5780632b2af0ab146103275780632e7acfa6146103515780632f30cabd146103665780633e55c0c7146103995780633e96576e146103ca5780633fe38627146103fd578063414f23fe1461049a57806345c5b2c7146104ca57806345e38b64146104f057806345e593ca14610505578063488ed1a9146105b95780634d26732d146105f45780634f0f4aa914610609578063567ca41b146106335780635c975abb146106665780635dbaf68b1461068f5780635e8ef106146106a45780635f576db6146106b95780636177fd18146106c157806362a82d7d146106f457806363721d6b1461071e57806365f7f80d1461073357806367425daf1461074857806369fd251c1461075d5780636b94c33b146107905780636f791d29146107c35780636f7d0026146107d85780637427be511461080e57806376e7e23b14610841578063771b2f97146108565780637ba9534a1461086b5780637e2d21551461088057806381fbc98a146108b057806383f94db7146108e35780638456cb59146109165780638640ce5f1461092b5780638da5cb5b1461094057806391c657e8146109555780639e8a713f14610988578063bf5ddcb11461099d578063ce11e6ab14610a28578063d01e660214610a3d578063d735e21d14610a67578063d93fe9c414610a7c578063dff6978714610a91578063e45b7ce614610aa6578063e8bd492214610ae1578063edfd03ed14610b4a578063ef40a67014610b74578063f31d863f14610ba7578063f33e1fac14610c85578063f51de41b14610caf578063f851a44014610cc4578063f8d1f19414610cd9578063fa7803e614610d03578063ff204f3b14610d3e575b600080fd5b3480156102ad57600080fd5b506102b6610d71565b005b3480156102c457600080fd5b506102eb600480360360208110156102db57600080fd5b50356001600160a01b0316610dc7565b60408051918252519081900360200190f35b34801561030957600080fd5b506102b66004803603602081101561032057600080fd5b5035610e7f565b34801561033357600080fd5b506102b66004803603602081101561034a57600080fd5b5035610ef7565b34801561035d57600080fd5b506102eb610f93565b34801561037257600080fd5b506102eb6004803603602081101561038957600080fd5b50356001600160a01b0316610f99565b3480156103a557600080fd5b506103ae610fb4565b604080516001600160a01b039092168252519081900360200190f35b3480156103d657600080fd5b506102eb600480360360208110156103ed57600080fd5b50356001600160a01b0316610fc3565b34801561040957600080fd5b506102b6600480360361024081101561042157600080fd5b813591602081019160e08201916101e08101359161020082013591908101906102408101610220820135600160201b81111561045c57600080fd5b82018360208201111561046e57600080fd5b803590602001918460018302840111600160201b8311171561048f57600080fd5b509092509050610fe1565b3480156104a657600080fd5b506102b6600480360360408110156104bd57600080fd5b5080359060200135611899565b6102b6600480360360208110156104e057600080fd5b50356001600160a01b0316611a64565b3480156104fc57600080fd5b506102eb611ab6565b34801561051157600080fd5b506102b660048036036101e081101561052957600080fd5b8135916020810135916040820135916060810135916080820135916001600160a01b0360a082013581169260c08301359091169190810190610100810160e0820135600160201b81111561057c57600080fd5b82018360208201111561058e57600080fd5b803590602001918460018302840111600160201b831117156105af57600080fd5b9193509150611abb565b3480156105c557600080fd5b506102b660048036036101408110156105dd57600080fd5b50604081016080820160c083016101008401611e3e565b34801561060057600080fd5b506102eb6126eb565b34801561061557600080fd5b506103ae6004803603602081101561062c57600080fd5b5035612710565b34801561063f57600080fd5b506102b66004803603602081101561065657600080fd5b50356001600160a01b031661272b565b34801561067257600080fd5b5061067b61282e565b604080519115158252519081900360200190f35b34801561069b57600080fd5b506103ae612837565b3480156106b057600080fd5b506102eb612846565b6102b661284c565b3480156106cd57600080fd5b5061067b600480360360208110156106e457600080fd5b50356001600160a01b031661289d565b34801561070057600080fd5b506103ae6004803603602081101561071757600080fd5b50356128c5565b34801561072a57600080fd5b506102eb6128ef565b34801561073f57600080fd5b506102eb6128f5565b34801561075457600080fd5b506102b66128fb565b34801561076957600080fd5b506103ae6004803603602081101561078057600080fd5b50356001600160a01b0316612965565b34801561079c57600080fd5b506102b6600480360360208110156107b357600080fd5b50356001600160a01b0316612986565b3480156107cf57600080fd5b5061067b612d7d565b3480156107e457600080fd5b506102eb600480360360608110156107fb57600080fd5b5080359060208101359060400135612d86565b34801561081a57600080fd5b506102b66004803603602081101561083157600080fd5b50356001600160a01b0316612d9d565b34801561084d57600080fd5b506102eb612e48565b34801561086257600080fd5b506102eb612e4e565b34801561087757600080fd5b506102eb612e54565b34801561088c57600080fd5b506102b6600480360360408110156108a357600080fd5b5080359060200135612e5a565b3480156108bc57600080fd5b506102eb600480360360208110156108d357600080fd5b50356001600160a01b031661303a565b3480156108ef57600080fd5b506102b66004803603602081101561090657600080fd5b50356001600160a01b03166130cf565b34801561092257600080fd5b506102b6613188565b34801561093757600080fd5b506102eb6131dc565b34801561094c57600080fd5b506103ae6131e2565b34801561096157600080fd5b5061067b6004803603602081101561097857600080fd5b50356001600160a01b03166131f1565b34801561099457600080fd5b506103ae61324b565b3480156109a957600080fd5b506102b6600480360360408110156109c057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156109ea57600080fd5b8201836020820111156109fc57600080fd5b803590602001918460018302840111600160201b83111715610a1d57600080fd5b50909250905061325a565b348015610a3457600080fd5b506103ae613350565b348015610a4957600080fd5b506103ae60048036036020811015610a6057600080fd5b503561335f565b348015610a7357600080fd5b506102eb61338e565b348015610a8857600080fd5b506103ae613394565b348015610a9d57600080fd5b506102eb6133a3565b348015610ab257600080fd5b506102b660048036036040811015610ac957600080fd5b506001600160a01b03813516906020013515156133a9565b348015610aed57600080fd5b50610b1460048036036020811015610b0457600080fd5b50356001600160a01b031661344b565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b348015610b5657600080fd5b506102b660048036036020811015610b6d57600080fd5b5035613487565b348015610b8057600080fd5b506102eb60048036036020811015610b9757600080fd5b50356001600160a01b03166134ec565b348015610bb357600080fd5b506102b6600480360360c0811015610bca57600080fd5b81359190810190604081016020820135600160201b811115610beb57600080fd5b820183602082011115610bfd57600080fd5b803590602001918460018302840111600160201b83111715610c1e57600080fd5b919390929091602081019035600160201b811115610c3b57600080fd5b820183602082011115610c4d57600080fd5b803590602001918460208302840111600160201b83111715610c6e57600080fd5b91935091508035906020810135906040013561350a565b348015610c9157600080fd5b506102eb60048036036020811015610ca857600080fd5b5035613a9d565b348015610cbb57600080fd5b506103ae613ac5565b348015610cd057600080fd5b506103ae613ad4565b348015610ce557600080fd5b506102eb60048036036020811015610cfc57600080fd5b5035613ae3565b348015610d0f57600080fd5b506102b660048036036040811015610d2657600080fd5b506001600160a01b0381358116916020013516613af5565b348015610d4a57600080fd5b506102b660048036036020811015610d6157600080fd5b50356001600160a01b0316613b61565b6016546001600160a01b03163314610dbd576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610dc5613c17565b565b600080610dd26128ef565b90506000805b82811015610e7557846001600160a01b0316639168ae72610df88361335f565b6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610e3557600080fd5b505afa158015610e49573d6000803e3d6000fd5b505050506040513d6020811015610e5f57600080fd5b505115610e6d576001909101905b600101610dd8565b509150505b919050565b610e8761282e565b15610ec7576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b610ed033613cb7565b6000610eda6126eb565b905080821015610ee8578091505b610ef23383613d4e565b505050565b610eff61338e565b811015610f45576040805162461bcd60e51b815260206004820152600f60248201526e1053149150511657d11150d2511151608a1b604482015290519081900360640190fd5b610f4d612e54565b811115610f90576040805162461bcd60e51b815260206004820152600c60248201526b1113d154d39517d1561254d560a21b604482015290519081900360640190fd5b50565b600c5481565b6001600160a01b03166000908152600a602052604090205490565b6011546001600160a01b031681565b6001600160a01b031660009081526008602052604090206001015490565b610fe961282e565b15611029576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b6110323361289d565b611070576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b600061107b33610fc3565b90506110856155ec565b600061109083612710565b90506000601160009054906101000a90046001600160a01b03166001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b1580156110e257600080fd5b505afa1580156110f6573d6000803e3d6000fd5b505050506040513d602081101561110c57600080fd5b50519050611118615621565b604080518082019091526111d2908c60026000835b8282101561116e576040805160608181019092529080840286019060039083908390808284376000920191909152505050815260019091019060200161112d565b50506040805180820190915291508d905060026000835b828210156111c65760408051608081810190925290808402860190600490839083908082843760009201919091525050508152600190910190602001611185565b505050508b8b86613e10565b90506000806000856001600160a01b031663701da98e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561121257600080fd5b505afa158015611226573d6000803e3d6000fd5b505050506040513d602081101561123c57600080fd5b5051845161124990613e5e565b1461128d576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b6011546020850151604090810151815163036648d960e11b81526024810182905260048101928352604481018d90526001600160a01b03909316926306cc91b2928e928e929091908190606401858580828437600081840152601f19601f820116905080830192505050945050505050604080518083038186803b15801561131457600080fd5b505afa158015611328573d6000803e3d6000fd5b505050506040513d604081101561133e57600080fd5b508051602090910151855160e001519194509250600090611360904390613ef3565b9050604b8110156113a5576040805162461bcd60e51b815260206004820152600a60248201526954494d455f44454c544160b01b604482015290519081900360640190fd5b8451516020860151516000916113bb9190613ef3565b9050856000015161010001518660200151604001511015806113ea5750600e546113e6908390613f50565b8110155b806114105750855160609081015160208801519091015160649161140e9190613ef3565b145b61144d576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d350531360ba1b604482015290519081900360640190fd5b61146d6004611467600e5485613f5090919063ffffffff16565b90613f50565b8111156114ad576040805162461bcd60e51b8152602060048201526009602482015268544f4f5f4c4152474560b81b604482015290519081900360640190fd5b600e546000906114d2906114cc6114c5826001613ef3565b8590613fa9565b90614001565b90506115608161155a6114f0600c5443613fa990919063ffffffff16565b8c6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561152957600080fd5b505afa15801561153d573d6000803e3d6000fd5b505050506040513d602081101561155357600080fd5b5051614065565b90613fa9565b93506000896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b15801561159d57600080fd5b505afa1580156115b1573d6000803e3d6000fd5b505050506040513d60208110156115c757600080fd5b50519050801561161957611616856115de83612710565b6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561152957600080fd5b94505b5050868660200151604001511115611669576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b5050600080600080896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b1580156116aa57600080fd5b505afa1580156116be573d6000803e3d6000fd5b505050506040513d60208110156116d457600080fd5b505111905080156117535761174c896001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b15801561171b57600080fd5b505afa15801561172f573d6000803e3d6000fd5b505050506040513d602081101561174557600080fd5b5051613ae3565b915061175f565b61175c8b613ae3565b91505b61176e878588888f878761407b565b9a50925050508f81146117bf576040805162461bcd60e51b81526020600482015260146024820152730aa9c8ab0a08a86a88a88be9c9e888abe9082a6960631b604482015290519081900360640190fd5b6117d3336117cb612e54565b600c54614331565b505050505050506117e382613ae3565b6117eb612e54565b7f8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a548b84608001518560400151866000015187602001518f8f6040518088815260200187815260200186815260200185815260200184815260200183600260600280828437600083820152601f01601f191690910190508261010080828437600083820152604051601f909101601f1916909201829003995090975050505050505050a3505050505050505050565b6118a161282e565b156118e1576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b6118ea3361289d565b611928576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b8061193283613ae3565b14611971576040805162461bcd60e51b815260206004820152600a6024820152694e4f44455f52454f524760b01b604482015290519081900360640190fd5b61197961338e565b821015801561198f575061198b612e54565b8211155b61199857600080fd5b60006119a383612710565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156119de57600080fd5b505afa1580156119f2573d6000803e3d6000fd5b505050506040513d6020811015611a0857600080fd5b5051611a1333610fc3565b14611a57576040805162461bcd60e51b815260206004820152600f60248201526e2727aa2fa9aa20a5a2a22fa82922ab60891b604482015290519081900360640190fd5b610ef23384600c54614331565b611a6c61282e565b15611aac576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b610f9081346144c4565b604b81565b600c5415611aff576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b88611b43576040805162461bcd60e51b815260206004820152600f60248201526e10905117d0d3d39197d411549253d1608a1b604482015290519081900360640190fd5b601080546001600160a01b03199081166001600160a01b03602085013581169190911792839055601180548316604080870135841691909117909155601280546060870135841694168417905580516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b158015611bd657600080fd5b505af1158015611bea573d6000803e3d6000fd5b5050505080600460078110611bfb57fe5b601380546001600160a01b0319166001600160a01b0360209390930293909301358216929092179091556010546040805163722dbe7360e11b8152608085013584166004820152600160248201529051919092169163e45b7ce691604480830192600092919082900301818387803b158015611c7657600080fd5b505af1158015611c8a573d6000803e3d6000fd5b505060135460405163b0f2af2960e01b8152600481018d8152602482018d9052604482018c9052606482018b90526001600160a01b038a8116608484015289811660a484015260e060c4840190815260e484018990529316945063b0f2af2993508d928d928d928d928d928d928d928d9261010401848480828437600081840152601f19601f8201169050808301925050509950505050505050505050600060405180830381600087803b158015611d4157600080fd5b505af1158015611d55573d6000803e3d6000fd5b5050505080600560078110611d6657fe5b601480546001600160a01b03199081166001600160a01b03602094909402949094013583169390931790556015805490921660c08401359091161790556000611dae8b61451f565b9050611db981614616565b600c8a9055600d899055600e889055600f879055601680546001600160a01b038781166001600160a01b031992831617909255601780549091168435909216919091179055604080518c815290517f4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d9181900360200190a15050505050505050505050565b611e4661282e565b15611e86576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b6020840135843510611ecd576040805162461bcd60e51b815260206004820152600b60248201526a2ba927a723afa7a92222a960a91b604482015290519081900360640190fd5b611ed5612e54565b60208501351115611f1c576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d41493d413d4d15160a21b604482015290519081900360640190fd5b8335611f266128f5565b10611f6c576040805162461bcd60e51b81526020600482015260116024820152701053149150511657d0d3d3919254935151607a1b604482015290519081900360640190fd5b6000611f7e85825b6020020135612710565b90506000611f8d866001611f74565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611fc857600080fd5b505afa158015611fdc573d6000803e3d6000fd5b505050506040513d6020811015611ff257600080fd5b5051604080516311e7249560e21b815290516001600160a01b0385169163479c9254916004808301926020929190829003018186803b15801561203457600080fd5b505afa158015612048573d6000803e3d6000fd5b505050506040513d602081101561205e57600080fd5b50511461209e576040805162461bcd60e51b81526020600482015260096024820152682224a3232fa82922ab60b91b604482015290519081900360640190fd5b6120b88760005b60200201356001600160a01b0316613cb7565b6120c38760016120a5565b604080516348b4573960e11b81526001600160a01b03893581166004830152915191841691639168ae7291602480820192602092909190829003018186803b15801561210e57600080fd5b505afa158015612122573d6000803e3d6000fd5b505050506040513d602081101561213857600080fd5b5051612180576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c57d393d517d4d51052d15160721b604482015290519081900360640190fd5b604080516348b4573960e11b81526001600160a01b0360208a81013582166004840152925190841692639168ae729260248082019391829003018186803b1580156121ca57600080fd5b505afa1580156121de573d6000803e3d6000fd5b505050506040513d60208110156121f457600080fd5b505161223c576040805162461bcd60e51b815260206004820152601260248201527114d51052d1548c97d393d517d4d51052d15160721b604482015290519081900360640190fd5b612251853585358560005b6020020135614665565b826001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b15801561228a57600080fd5b505afa15801561229e573d6000803e3d6000fd5b505050506040513d60208110156122b457600080fd5b5051146122f5576040805162461bcd60e51b815260206004820152600a6024820152694348414c5f484153483160b01b604482015290519081900360640190fd5b61230a60208087013590860135856001612247565b816001600160a01b0316635b8b22806040518163ffffffff1660e01b815260040160206040518083038186803b15801561234357600080fd5b505afa158015612357573d6000803e3d6000fd5b505050506040513d602081101561236d57600080fd5b5051146123ae576040805162461bcd60e51b815260206004820152600a60248201526921a420a62fa420a9a41960b11b604482015290519081900360640190fd5b60006124ff612420846001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156123ef57600080fd5b505afa158015612403573d6000803e3d6000fd5b505050506040513d602081101561241957600080fd5b5051612710565b6001600160a01b031663d7ff5e356040518163ffffffff1660e01b815260040160206040518083038186803b15801561245857600080fd5b505afa15801561246c573d6000803e3d6000fd5b505050506040513d602081101561248257600080fd5b5051600d5461155a90818960006020020135886001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156124cd57600080fd5b505afa1580156124e1573d6000803e3d6000fd5b505050506040513d60208110156124f757600080fd5b505190613ef3565b90506020850135811015612539576125316001600160a01b0389351689600160200201356001600160a01b031661469c565b5050506126e4565b6014546000906001600160a01b0390811690638ecaab119030908a35908935908e35168e600160200201356001600160a01b03166125918d60006002811061257d57fe5b60200201358a613ef390919063ffffffff16565b6125ab8e600160200201358b613ef390919063ffffffff16565b601154601054604080516001600160e01b031960e08d901b1681526001600160a01b039a8b166004820152602481019990995260448901979097529488166064880152928716608487015260a486019190915260c4850152841660e484015290921661010482015290516101248083019260209291908290030181600087803b15801561263757600080fd5b505af115801561264b573d6000803e3d6000fd5b505050506040513d602081101561266157600080fd5b5051905061268a6001600160a01b038a35168a600160200201356001600160a01b031683614717565b604080518a356001600160a01b0390811682526020808d01358216908301528a35828401529151918316917fa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda758799181900360600190a2505050505b5050505050565b6000806126f661338e565b905061270a4382612705612e54565b614761565b91505090565b6000908152600560205260409020546001600160a01b031690565b6016546001600160a01b03163314612777576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6012546001600160a01b03828116911614156127c7576040805162461bcd60e51b815260206004820152600a602482015269086aaa4be9eaaa8849eb60b31b604482015290519081900360640190fd5b601054604080516319dc7ae560e31b81526001600160a01b038481166004830152600060248301819052925193169263cee3d7289260448084019391929182900301818387803b15801561281a57600080fd5b505af11580156126e4573d6000803e3d6000fd5b600b5460ff1690565b6014546001600160a01b031681565b600e5481565b61285461282e565b15612894576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b610dc5346149f4565b6001600160a01b0316600090815260086020526040902060030154600160a01b900460ff1690565b6000600782815481106128d457fe5b6000918252602090912001546001600160a01b031692915050565b60095490565b60015490565b600061290561338e565b905061290f6128f5565b811180156129245750612920612e54565b8111155b610f90576040805162461bcd60e51b815260206004820152600d60248201526c1393d7d553949154d3d3159151609a1b604482015290519081900360640190fd5b6001600160a01b039081166000908152600860205260409020600301541690565b61298e61282e565b156129ce576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b6129d66128fb565b60006129e06128f5565b905060006129ec61338e565b905060006129f982612710565b905082816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015612a3557600080fd5b505afa158015612a49573d6000803e3d6000fd5b505050506040513d6020811015612a5f57600080fd5b50511415612cdf57612a708461289d565b612aae576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b612abf612aba85610fc3565b610ef7565b806001600160a01b0316639168ae72856040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015612b0c57600080fd5b505afa158015612b20573d6000803e3d6000fd5b505050506040513d6020811015612b3657600080fd5b505115612b7d576040805162461bcd60e51b815260206004820152601060248201526f14d51052d15117d3d397d5105491d15560821b604482015290519081900360640190fd5b806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b158015612bb657600080fd5b505afa158015612bca573d6000803e3d6000fd5b50505050612bd783612710565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b158015612c0f57600080fd5b505afa158015612c23573d6000803e3d6000fd5b50505050612c316000613487565b612c3a81610dc7565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015612c7357600080fd5b505afa158015612c87573d6000803e3d6000fd5b505050506040513d6020811015612c9d57600080fd5b505114612cdf576040805162461bcd60e51b815260206004820152600b60248201526a4841535f5354414b45525360a81b604482015290519081900360640190fd5b612ce7614b8f565b60135460408051630c2a09ad60e21b81526004810185905290516001600160a01b03909216916330a826b49160248082019260009290919082900301818387803b158015612d3457600080fd5b505af1158015612d48573d6000803e3d6000fd5b50506040518492507f9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a39150600090a250505050565b60005460ff1690565b6000612d93848484614761565b90505b9392505050565b612da561282e565b15612de5576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b612ded6128f5565b612df682610fc3565b1115612e36576040805162461bcd60e51b815260206004820152600a6024820152691513d3d7d49150d1539560b21b604482015290519081900360640190fd5b612e3f81613cb7565b610f9081614ba5565b600f5481565b600d5481565b60035490565b612e6261282e565b15612ea2576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b612eaa6128ef565b821115612eef576040805162461bcd60e51b815260206004820152600e60248201526d4e4f5f535543485f5a4f4d42494560901b604482015290519081900360640190fd5b6000612efa8361335f565b90506000612f0784613a9d565b9050600080612f1461338e565b90505b808310158015612f2657508482105b15613012576000612f3684612710565b9050806001600160a01b03166396a9fdc0866040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015612f8757600080fd5b505af1158015612f9b573d6000803e3d6000fd5b50505050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015612fd857600080fd5b505afa158015612fec573d6000803e3d6000fd5b505050506040513d602081101561300257600080fd5b5051935050600190910190612f17565b808310156130285761302386614bf8565b613032565b6130328684614c94565b505050505050565b600061304461282e565b15613084576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b600061308f33614cbb565b6040519091506001600160a01b0384169082156108fc029083906000818181858888f193505050501580156130c8573d6000803e3d6000fd5b5092915050565b6016546001600160a01b0316331461311b576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6017546040805163266a23b160e21b815230600482018190526001600160a01b0385811660248401529251909392909216916399a88ec49160448082019260009290919082900301818387803b15801561317457600080fd5b505af1158015613032573d6000803e3d6000fd5b6016546001600160a01b031633146131d4576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b610dc5614cda565b60045490565b6016546001600160a01b031681565b6000805b600954811015613242576009818154811061320c57fe5b60009182526020909120600290910201546001600160a01b038481169116141561323a576001915050610e7a565b6001016131f5565b50600092915050565b6013546001600160a01b031681565b6016546001600160a01b031633146132a6576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601754604051639623609d60e01b815230600482018181526001600160a01b0387811660248501526060604485019081526064850187905292941692639623609d92859289928992899291608401848480828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b15801561333257600080fd5b505af1158015613346573d6000803e3d6000fd5b5050505050505050565b6012546001600160a01b031681565b60006009828154811061336e57fe5b60009182526020909120600290910201546001600160a01b031692915050565b60025490565b6015546001600160a01b031681565b60075490565b6016546001600160a01b031633146133f5576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b6010546040805163722dbe7360e11b81526001600160a01b03858116600483015284151560248301529151919092169163e45b7ce691604480830192600092919082900301818387803b15801561317457600080fd5b6008602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b60006134916128ef565b9050600061349d61338e565b9050825b828110156134e6575b816134b482613a9d565b10156134de576134c381614bf8565b600019909201918281106134d957505050610f90565b6134aa565b6001016134a1565b50505050565b6001600160a01b031660009081526008602052604090206002015490565b61351261282e565b15613552576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b61355a6128fb565b60006135646133a3565b116135a3576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f5354414b45525360b01b604482015290519081900360640190fd5b60006135ad61338e565b905060006135ba82612710565b9050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b1580156135f557600080fd5b505afa158015613609573d6000803e3d6000fd5b505050506136156128f5565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561364e57600080fd5b505afa158015613662573d6000803e3d6000fd5b505050506040513d602081101561367857600080fd5b5051146136bb576040805162461bcd60e51b815260206004820152600c60248201526b24a72b20a624a22fa82922ab60a11b604482015290519081900360640190fd5b6136cb6136c66128f5565b612710565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b15801561370357600080fd5b505afa158015613717573d6000803e3d6000fd5b505050506137256000613487565b61373961373182610dc7565b61155a6133a3565b816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561377257600080fd5b505afa158015613786573d6000803e3d6000fd5b505050506040513d602081101561379c57600080fd5b5051146137e1576040805162461bcd60e51b815260206004820152600e60248201526d1393d517d0531317d4d51052d15160921b604482015290519081900360640190fd5b60006138628a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c918291850190849080828437600081840152601f19601f820116905080830192505050505050508d614d58565b90506138718b82878988614e59565b826001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b1580156138aa57600080fd5b505afa1580156138be573d6000803e3d6000fd5b505050506040513d60208110156138d457600080fd5b505114613917576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b60125460408051630c72684760e01b815260048101918252604481018c90526001600160a01b0390921691630c726847918d918d918d918d919081906024810190606401878780828437600083820152601f01601f19169091018481038352858152602090810191508690860280828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b1580156139c157600080fd5b505af11580156139d5573d6000803e3d6000fd5b505050506139e1614ea0565b601354604080516316b9109b60e01b81526004810186905290516001600160a01b03909216916316b9109b9160248082019260009290919082900301818387803b158015613a2e57600080fd5b505af1158015613a42573d6000803e3d6000fd5b505060408051848152602081018a90528082018990526060810188905290518693507f2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a292509081900360800190a25050505050505050505050565b600060098281548110613aac57fe5b9060005260206000209060020201600101549050919050565b6010546001600160a01b031681565b6017546001600160a01b031681565b60009081526006602052604090205490565b613aff8282614eb9565b6001600160a01b0316336001600160a01b031614613b53576040805162461bcd60e51b815260206004820152600c60248201526b2ba927a723afa9a2a72222a960a11b604482015290519081900360640190fd5b613b5d828261469c565b5050565b6016546001600160a01b03163314613bad576040805162461bcd60e51b815260206004820152600a60248201526927a7262cafa7aba722a960b11b604482015290519081900360640190fd5b601280546001600160a01b0319166001600160a01b03838116918217909255601054604080516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b15801561281a57600080fd5b613c1f61282e565b613c67576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b600b805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613c9a614f7b565b604080516001600160a01b039092168252519081900360200190a1565b613cc08161289d565b613cfe576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b6000613d0982612965565b6001600160a01b031614610f90576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600860205260408120600281015480841115613db3576040805162461bcd60e51b815260206004820152601060248201526f544f4f5f4c4954544c455f5354414b4560801b604482015290519081900360640190fd5b6000613dbf8286613ef3565b600284018690556001600160a01b0387166000908152600a6020526040902054909150613dec9082613fa9565b6001600160a01b0387166000908152600a6020526040902055925050505b92915050565b613e18615621565b60408051808201909152865186518291613e33918888614f7f565b8152602001613e52886001602002015188600160200201514387614f7f565b90529695505050505050565b6000816000015182602001518360400151846060015185608001518660a001518760c001518860e00151896101000151604051602001808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019950505050505050505050604051602081830303815290604052805190602001209050919050565b600082821115613f4a576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082613f5f57506000613e0a565b82820282848281613f6c57fe5b0414612d965760405162461bcd60e51b81526004018080602001828103825260218152602001806156bb6021913960400191505060405180910390fd5b600082820183811015612d96576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b6000808211614054576040805162461bcd60e51b815260206004820152601a602482015279536166654d6174683a206469766973696f6e206279207a65726f60301b604482015290519081900360640190fd5b81838161405d57fe5b049392505050565b60008183116140745781612d96565b5090919050565b60006140856155ec565b61408d6155ec565b601160009054906101000a90046001600160a01b03166001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b1580156140db57600080fd5b505afa1580156140ef573d6000803e3d6000fd5b505050506040513d602081101561410557600080fd5b5051604082015261411586612710565b6001600160a01b031660a0820152600061412d612e54565b600101905061413b8b61501d565b60808301528882526020820188905260135460408051638b8ca19960e01b815260048101849052602481018a9052604481018d905233606482015290516001600160a01b0390921691638b8ca1999160848082019260009290919082900301818387803b1580156141ab57600080fd5b505af11580156141bf573d6000803e3d6000fd5b505060155460208e01516001600160a01b03909116925063d45ab2b591506141e690613e5e565b6141f58e86608001514361504e565b6141fe8f615063565b8b8f6040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b15801561425257600080fd5b505af1158015614266573d6000803e3d6000fd5b505050506040513d602081101561427c57600080fd5b50516001600160a01b0316606083015250608081015160208201516000916142a79187918991615093565b90506142b78260600151826150f7565b8160a001516001600160a01b0316631bc09d0a6142d2612e54565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561430857600080fd5b505af115801561431c573d6000803e3d6000fd5b50929d939c50929a5050505050505050505050565b6001600160a01b0380841660008181526008602090815260408083208784526005835281842054825163123334b760e11b815260048101969096529151909591909116938492632466696e9260248084019382900301818787803b15801561439857600080fd5b505af11580156143ac573d6000803e3d6000fd5b505050506040513d60208110156143c257600080fd5b5051600180850187905590915081141561303257600060056000846001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561441557600080fd5b505afa158015614429573d6000803e3d6000fd5b505050506040513d602081101561443f57600080fd5b505181526020810191909152604001600020546001600160a01b0316905080636971dfe561446d4388613fa9565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156144a357600080fd5b505af11580156144b7573d6000803e3d6000fd5b5050505050505050505050565b6144cc61282e565b1561450c576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b61451582613cb7565b613b5d8282615141565b600080614575604051806101200160405280600081526020018581526020016000815260200160008152602001600081526020016000801b81526020016000801b81526020014381526020016001815250613e5e565b6015546040805163d45ab2b560e01b815260048101849052600060248201819052604482018190526064820181905243608483015291519394506001600160a01b039092169263d45ab2b59260a4808201936020939283900390910190829087803b1580156145e357600080fd5b505af11580156145f7573d6000803e3d6000fd5b505050506040513d602081101561460d57600080fd5b50519392505050565b6000805260056020527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc80546001600160a01b0319166001600160a01b03929092169190911790556001600255565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b60006146a7826134ec565b905060006146b4846134ec565b9050808211156146d5576146d26146cb8483613d4e565b8390613ef3565b91505b600282046146e38582615141565b6146ed8382613ef3565b92506146f885615172565b60165461470e906001600160a01b03168461519c565b6126e4846151df565b6001600160a01b03928316600090815260086020526040808220600390810180549487166001600160a01b0319958616811790915594909516825290209092018054909216179055565b6000816001840314156147775750600f54612d96565b600061478284612710565b6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156147ba57600080fd5b505afa1580156147ce573d6000803e3d6000fd5b505050506040513d60208110156147e457600080fd5b50519050808510156147fa575050600f54612d96565b614802615646565b506040805161014081018252600181526201e05b60208201526201f7d191810191909152620138916060820152620329e160808201526201be4360a08201526204cb8c60c08201526201fbc460e082015262036d326101008201526202797361012082015261486f615646565b506040805161014081018252600181526201c03060208201526201b6999181019190915261fde26060820152620265c6608082015262013b8e60a0820152620329e160c08201526201389160e08201526201f7d16101008201526201537561012082015260006148df8885613ef3565b905060006148fd600c546114cc600a85613f5090919063ffffffff16565b905060ff61490c82600a614001565b106149205760001995505050505050612d96565b600061492d82600a614001565b60020a9050600085600a8406600a811061494357fe5b602002015162ffffff168202905085600a8406600a811061496057fe5b602002015162ffffff1682828161497357fe5b041461498a57600019975050505050505050612d96565b60006149af86600a8606600a811061499e57fe5b6020020151839062ffffff16614001565b9050806149ba575060015b600f5480820290829082816149cb57fe5b04146149e4576000199950505050505050505050612d96565b9c9b505050505050505050505050565b6149fc61282e565b15614a3c576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b614a453361289d565b15614a88576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4d51052d15160921b604482015290519081900360640190fd5b614a91336131f1565b15614ad6576040805162461bcd60e51b815260206004820152601060248201526f5354414b45525f49535f5a4f4d42494560801b604482015290519081900360640190fd5b614ade6126eb565b811015614b25576040805162461bcd60e51b815260206004820152601060248201526f4e4f545f454e4f5547485f5354414b4560801b604482015290519081900360640190fd5b614b2f338261528f565b6013546001600160a01b031663f03c04a533614b496128f5565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561281a57600080fd5b614b9a60025461535a565b600280546001019055565b6001600160a01b03811660009081526008602090815260408083206002810154600a909352922054614bd691613fa9565b6001600160a01b0383166000908152600a6020526040902055613b5d826153dc565b600980546000198101908110614c0a57fe5b906000526020600020906002020160098281548110614c2557fe5b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b039092169190911781556001918201549101556009805480614c6857fe5b60008281526020812060026000199093019283020180546001600160a01b031916815560010155905550565b8060098381548110614ca257fe5b9060005260206000209060020201600101819055505050565b6001600160a01b03166000908152600a60205260408120805491905590565b614ce261282e565b15614d22576040805162461bcd60e51b815260206004820152601060248201526000805160206156dc833981519152604482015290519081900360640190fd5b600b805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613c9a614f7b565b81518351600091829184835b83811015614e0b576000888281518110614d7a57fe5b60200260200101519050838187011115614dca576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868b0181018290206040805180840196909652858101919091528051808603820181526060909501905283519301929092209190940193600101614d64565b50818414614e4e576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b979650505050505050565b60408051602080820197909752808201959095526060850192909252608084019290925260a0808401929092528051808403909201825260c0909201909152805191012090565b614eab60015461535a565b600280546001818155019055565b6001600160a01b0380831660009081526008602052604080822084841683529082206003808301549082015493949293919290811691168114614f2d576040805162461bcd60e51b8152602060048201526007602482015266125397d0d2105360ca1b604482015290519081900360640190fd5b6001600160a01b038116614f72576040805162461bcd60e51b81526020600482015260076024820152661393d7d0d2105360ca1b604482015290519081900360640190fd5b95945050505050565b3390565b614f87615665565b604080516101208101825285518152865160208201529081018560016020020151815260200185600260048110614fba57fe5b6020020151815260200185600360048110614fd157fe5b6020020151815260200186600160038110614fe857fe5b6020020151815260200186600260038110614fff57fe5b60200201518152602001848152602001838152509050949350505050565b80518051602083015151600092613e0a92918290039061503c90615502565b6150498660200151615502565b615537565b6000612d938383866020015160400151614665565b805160a09081015160208301519182015160c08301516060840151608090940151600094613e0a94939291614e59565b600080856150a25760006150a5565b60015b905080858585604051602001808560ff1660f81b815260010184815260200183815260200182815260200194505050505060405160208183030381529060405280519060200120915050949350505050565b60038054600101808255600090815260056020908152604080832080546001600160a01b0319166001600160a01b0397909716969096179095559154815260069091529190912055565b6001600160a01b038216600090815260086020526040902060028101546151689083613fa9565b6002909101555050565b6001600160a01b0316600090815260086020526040902060030180546001600160a01b0319169055565b6001600160a01b0382166000908152600a60205260409020546151bf9082613fa9565b6001600160a01b039092166000908152600a602052604090209190915550565b6001600160a01b0381811660008181526008602090815260408083208151808301909252938152600180850154928201928352600980549182018155909352517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af600290930292830180546001600160a01b031916919095161790935591517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7b090920191909155613b5d826153dc565b6007805460018082019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688810180546001600160a01b039586166001600160a01b031991821681179092556040805160a08101825293845284546020858101918252858301978852600060608701818152608088018981529682526008909252929092209451855551948401949094559351600283015591516003909101805492511515600160a01b0260ff60a01b199290951692909316919091171691909117905543600455565b60008181526005602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b1580156153a657600080fd5b505af11580156153ba573d6000803e3d6000fd5b50505060009182525060056020526040902080546001600160a01b0319169055565b6001600160a01b0381166000908152600860205260409020805460078054600019810190811061540857fe5b600091825260209091200154600780546001600160a01b03909216918390811061542e57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600860006007848154811061546e57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600780548061549e57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03949094168152600890935250506040812081815560018101829055600281019190915560030180546001600160a81b0319169055565b6000613e0a8260000151615532846040015185602001518660a0015187606001518860c001518960800151615575565b6155c0565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b6040518060400160405280615634615665565b8152602001615641615665565b905290565b604051806101400160405280600a906020820280368337509192915050565b604051806101200160405280600081526020016000801916815260200160008152602001600081526020016000815260200160008019168152602001600080191681526020016000815260200160008152509056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775061757361626c653a2070617573656400000000000000000000000000000000a2646970667358221220c2b58fe40f2f91145a959ed6042270bd055e9aad072676f684b3d11515fa03fd64736f6c634300060c0033"

// DeployRollup deploys a new Ethereum contract, binding an instance of Rollup to it.
func DeployRollup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Rollup, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "_stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _Rollup.Contract.StakerMap(&_Rollup.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_Rollup *RollupCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _Rollup.Contract.StakerMap(&_Rollup.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Rollup *RollupCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Rollup *RollupSession) Admin() (common.Address, error) {
	return _Rollup.Contract.Admin(&_Rollup.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Rollup *RollupCallerSession) Admin() (common.Address, error) {
	return _Rollup.Contract.Admin(&_Rollup.CallOpts)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_Rollup *RollupCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_Rollup *RollupSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _Rollup.Contract.AmountStaked(&_Rollup.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_Rollup *RollupCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _Rollup.Contract.AmountStaked(&_Rollup.CallOpts, staker)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupCaller) ArbGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "arbGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _Rollup.Contract.ArbGasSpeedLimitPerBlock(&_Rollup.CallOpts)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_Rollup *RollupCallerSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _Rollup.Contract.ArbGasSpeedLimitPerBlock(&_Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupCaller) BaseStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "baseStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupSession) BaseStake() (*big.Int, error) {
	return _Rollup.Contract.BaseStake(&_Rollup.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_Rollup *RollupCallerSession) BaseStake() (*big.Int, error) {
	return _Rollup.Contract.BaseStake(&_Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupSession) ChallengeFactory() (common.Address, error) {
	return _Rollup.Contract.ChallengeFactory(&_Rollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Rollup *RollupCallerSession) ChallengeFactory() (common.Address, error) {
	return _Rollup.Contract.ChallengeFactory(&_Rollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_Rollup *RollupCaller) ConfirmPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "confirmPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_Rollup *RollupSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _Rollup.Contract.ConfirmPeriodBlocks(&_Rollup.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_Rollup *RollupCallerSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _Rollup.Contract.ConfirmPeriodBlocks(&_Rollup.CallOpts)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupCaller) CountStakedZombies(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "countStakedZombies", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _Rollup.Contract.CountStakedZombies(&_Rollup.CallOpts, node)
}

// CountStakedZombies is a free data retrieval call binding the contract method 0x04a28064.
//
// Solidity: function countStakedZombies(address node) view returns(uint256)
func (_Rollup *RollupCallerSession) CountStakedZombies(node common.Address) (*big.Int, error) {
	return _Rollup.Contract.CountStakedZombies(&_Rollup.CallOpts, node)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_Rollup *RollupCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_Rollup *RollupSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _Rollup.Contract.CurrentChallenge(&_Rollup.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_Rollup *RollupCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _Rollup.Contract.CurrentChallenge(&_Rollup.CallOpts, staker)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupCaller) CurrentRequiredStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "currentRequiredStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupSession) CurrentRequiredStake() (*big.Int, error) {
	return _Rollup.Contract.CurrentRequiredStake(&_Rollup.CallOpts)
}

// CurrentRequiredStake is a free data retrieval call binding the contract method 0x4d26732d.
//
// Solidity: function currentRequiredStake() view returns(uint256)
func (_Rollup *RollupCallerSession) CurrentRequiredStake() (*big.Int, error) {
	return _Rollup.Contract.CurrentRequiredStake(&_Rollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_Rollup *RollupCaller) DelayedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "delayedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_Rollup *RollupSession) DelayedBridge() (common.Address, error) {
	return _Rollup.Contract.DelayedBridge(&_Rollup.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_Rollup *RollupCallerSession) DelayedBridge() (common.Address, error) {
	return _Rollup.Contract.DelayedBridge(&_Rollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_Rollup *RollupCaller) ExtraChallengeTimeBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "extraChallengeTimeBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_Rollup *RollupSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _Rollup.Contract.ExtraChallengeTimeBlocks(&_Rollup.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_Rollup *RollupCallerSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _Rollup.Contract.ExtraChallengeTimeBlocks(&_Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupSession) FirstUnresolvedNode() (*big.Int, error) {
	return _Rollup.Contract.FirstUnresolvedNode(&_Rollup.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_Rollup *RollupCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _Rollup.Contract.FirstUnresolvedNode(&_Rollup.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_Rollup *RollupCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_Rollup *RollupSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetNode(&_Rollup.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_Rollup *RollupCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetNode(&_Rollup.CallOpts, nodeNum)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_Rollup *RollupCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_Rollup *RollupSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _Rollup.Contract.GetNodeHash(&_Rollup.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_Rollup *RollupCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _Rollup.Contract.GetNodeHash(&_Rollup.CallOpts, index)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_Rollup *RollupCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_Rollup *RollupSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetStakerAddress(&_Rollup.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_Rollup *RollupCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetStakerAddress(&_Rollup.CallOpts, stakerNum)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Rollup *RollupCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Rollup *RollupSession) IsMaster() (bool, error) {
	return _Rollup.Contract.IsMaster(&_Rollup.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Rollup *RollupCallerSession) IsMaster() (bool, error) {
	return _Rollup.Contract.IsMaster(&_Rollup.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_Rollup *RollupCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_Rollup *RollupSession) IsStaked(staker common.Address) (bool, error) {
	return _Rollup.Contract.IsStaked(&_Rollup.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_Rollup *RollupCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _Rollup.Contract.IsStaked(&_Rollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_Rollup *RollupCaller) IsZombie(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isZombie", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_Rollup *RollupSession) IsZombie(staker common.Address) (bool, error) {
	return _Rollup.Contract.IsZombie(&_Rollup.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_Rollup *RollupCallerSession) IsZombie(staker common.Address) (bool, error) {
	return _Rollup.Contract.IsZombie(&_Rollup.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_Rollup *RollupCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_Rollup *RollupSession) LastStakeBlock() (*big.Int, error) {
	return _Rollup.Contract.LastStakeBlock(&_Rollup.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_Rollup *RollupCallerSession) LastStakeBlock() (*big.Int, error) {
	return _Rollup.Contract.LastStakeBlock(&_Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupSession) LatestConfirmed() (*big.Int, error) {
	return _Rollup.Contract.LatestConfirmed(&_Rollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_Rollup *RollupCallerSession) LatestConfirmed() (*big.Int, error) {
	return _Rollup.Contract.LatestConfirmed(&_Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupSession) LatestNodeCreated() (*big.Int, error) {
	return _Rollup.Contract.LatestNodeCreated(&_Rollup.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_Rollup *RollupCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _Rollup.Contract.LatestNodeCreated(&_Rollup.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_Rollup *RollupCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_Rollup *RollupSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _Rollup.Contract.LatestStakedNode(&_Rollup.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_Rollup *RollupCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _Rollup.Contract.LatestStakedNode(&_Rollup.CallOpts, staker)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _Rollup.Contract.MinimumAssertionPeriod(&_Rollup.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_Rollup *RollupCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _Rollup.Contract.MinimumAssertionPeriod(&_Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupSession) NodeFactory() (common.Address, error) {
	return _Rollup.Contract.NodeFactory(&_Rollup.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_Rollup *RollupCallerSession) NodeFactory() (common.Address, error) {
	return _Rollup.Contract.NodeFactory(&_Rollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_Rollup *RollupCaller) Outbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "outbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_Rollup *RollupSession) Outbox() (common.Address, error) {
	return _Rollup.Contract.Outbox(&_Rollup.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_Rollup *RollupCallerSession) Outbox() (common.Address, error) {
	return _Rollup.Contract.Outbox(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCallerSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupSession) Paused() (bool, error) {
	return _Rollup.Contract.Paused(&_Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupCallerSession) Paused() (bool, error) {
	return _Rollup.Contract.Paused(&_Rollup.CallOpts)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_Rollup *RollupCaller) RequireUnresolved(opts *bind.CallOpts, nodeNum *big.Int) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "requireUnresolved", nodeNum)

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_Rollup *RollupSession) RequireUnresolved(nodeNum *big.Int) error {
	return _Rollup.Contract.RequireUnresolved(&_Rollup.CallOpts, nodeNum)
}

// RequireUnresolved is a free data retrieval call binding the contract method 0x2b2af0ab.
//
// Solidity: function requireUnresolved(uint256 nodeNum) view returns()
func (_Rollup *RollupCallerSession) RequireUnresolved(nodeNum *big.Int) error {
	return _Rollup.Contract.RequireUnresolved(&_Rollup.CallOpts, nodeNum)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_Rollup *RollupCaller) RequireUnresolvedExists(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "requireUnresolvedExists")

	if err != nil {
		return err
	}

	return err

}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_Rollup *RollupSession) RequireUnresolvedExists() error {
	return _Rollup.Contract.RequireUnresolvedExists(&_Rollup.CallOpts)
}

// RequireUnresolvedExists is a free data retrieval call binding the contract method 0x67425daf.
//
// Solidity: function requireUnresolvedExists() view returns()
func (_Rollup *RollupCallerSession) RequireUnresolvedExists() error {
	return _Rollup.Contract.RequireUnresolvedExists(&_Rollup.CallOpts)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_Rollup *RollupCaller) RequiredStake(opts *bind.CallOpts, blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "requiredStake", blockNumber, firstUnresolvedNodeNum, latestNodeCreated)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_Rollup *RollupSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _Rollup.Contract.RequiredStake(&_Rollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RequiredStake is a free data retrieval call binding the contract method 0x6f7d0026.
//
// Solidity: function requiredStake(uint256 blockNumber, uint256 firstUnresolvedNodeNum, uint256 latestNodeCreated) view returns(uint256)
func (_Rollup *RollupCallerSession) RequiredStake(blockNumber *big.Int, firstUnresolvedNodeNum *big.Int, latestNodeCreated *big.Int) (*big.Int, error) {
	return _Rollup.Contract.RequiredStake(&_Rollup.CallOpts, blockNumber, firstUnresolvedNodeNum, latestNodeCreated)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_Rollup *RollupCaller) RollupEventBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "rollupEventBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_Rollup *RollupSession) RollupEventBridge() (common.Address, error) {
	return _Rollup.Contract.RollupEventBridge(&_Rollup.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_Rollup *RollupCallerSession) RollupEventBridge() (common.Address, error) {
	return _Rollup.Contract.RollupEventBridge(&_Rollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_Rollup *RollupCaller) SequencerBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "sequencerBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_Rollup *RollupSession) SequencerBridge() (common.Address, error) {
	return _Rollup.Contract.SequencerBridge(&_Rollup.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_Rollup *RollupCallerSession) SequencerBridge() (common.Address, error) {
	return _Rollup.Contract.SequencerBridge(&_Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupSession) StakerCount() (*big.Int, error) {
	return _Rollup.Contract.StakerCount(&_Rollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Rollup *RollupCallerSession) StakerCount() (*big.Int, error) {
	return _Rollup.Contract.StakerCount(&_Rollup.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_Rollup *RollupCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_Rollup *RollupSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _Rollup.Contract.WithdrawableFunds(&_Rollup.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_Rollup *RollupCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _Rollup.Contract.WithdrawableFunds(&_Rollup.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_Rollup *RollupCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_Rollup *RollupSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.ZombieAddress(&_Rollup.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_Rollup *RollupCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _Rollup.Contract.ZombieAddress(&_Rollup.CallOpts, zombieNum)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Rollup *RollupCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Rollup *RollupSession) ZombieCount() (*big.Int, error) {
	return _Rollup.Contract.ZombieCount(&_Rollup.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Rollup *RollupCallerSession) ZombieCount() (*big.Int, error) {
	return _Rollup.Contract.ZombieCount(&_Rollup.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_Rollup *RollupCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_Rollup *RollupSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _Rollup.Contract.ZombieLatestStakedNode(&_Rollup.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_Rollup *RollupCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _Rollup.Contract.ZombieLatestStakedNode(&_Rollup.CallOpts, zombieNum)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupTransactor) AddToDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addToDeposit", stakerAddress)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupSession) AddToDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddToDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// AddToDeposit is a paid mutator transaction binding the contract method 0x45c5b2c7.
//
// Solidity: function addToDeposit(address stakerAddress) payable returns()
func (_Rollup *RollupTransactorSession) AddToDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddToDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupTransactor) CompleteChallenge(opts *bind.TransactOpts, winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "completeChallenge", winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.CompleteChallenge(&_Rollup.TransactOpts, winningStaker, losingStaker)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0xfa7803e6.
//
// Solidity: function completeChallenge(address winningStaker, address losingStaker) returns()
func (_Rollup *RollupTransactorSession) CompleteChallenge(winningStaker common.Address, losingStaker common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.CompleteChallenge(&_Rollup.TransactOpts, winningStaker, losingStaker)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_Rollup *RollupTransactor) ConfirmNextNode(opts *bind.TransactOpts, beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "confirmNextNode", beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_Rollup *RollupSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ConfirmNextNode(&_Rollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ConfirmNextNode is a paid mutator transaction binding the contract method 0xf31d863f.
//
// Solidity: function confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_Rollup *RollupTransactorSession) ConfirmNextNode(beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ConfirmNextNode(&_Rollup.TransactOpts, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_Rollup *RollupTransactor) CreateChallenge(opts *bind.TransactOpts, stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "createChallenge", stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_Rollup *RollupSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x488ed1a9.
//
// Solidity: function createChallenge(address[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts) returns()
func (_Rollup *RollupTransactorSession) CreateChallenge(stakers [2]common.Address, nodeNums [2]*big.Int, executionHashes [2][32]byte, proposedTimes [2]*big.Int, maxMessageCounts [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.CreateChallenge(&_Rollup.TransactOpts, stakers, nodeNums, executionHashes, proposedTimes, maxMessageCounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_Rollup *RollupSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45e593ca.
//
// Solidity: function initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[7] connectedContracts) returns()
func (_Rollup *RollupTransactorSession) Initialize(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte, connectedContracts [7]common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig, connectedContracts)
}

// NewStake is a paid mutator transaction binding the contract method 0x5f576db6.
//
// Solidity: function newStake() payable returns()
func (_Rollup *RollupTransactor) NewStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "newStake")
}

// NewStake is a paid mutator transaction binding the contract method 0x5f576db6.
//
// Solidity: function newStake() payable returns()
func (_Rollup *RollupSession) NewStake() (*types.Transaction, error) {
	return _Rollup.Contract.NewStake(&_Rollup.TransactOpts)
}

// NewStake is a paid mutator transaction binding the contract method 0x5f576db6.
//
// Solidity: function newStake() payable returns()
func (_Rollup *RollupTransactorSession) NewStake() (*types.Transaction, error) {
	return _Rollup.Contract.NewStake(&_Rollup.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rollup *RollupTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rollup *RollupSession) Pause() (*types.Transaction, error) {
	return _Rollup.Contract.Pause(&_Rollup.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rollup *RollupTransactorSession) Pause() (*types.Transaction, error) {
	return _Rollup.Contract.Pause(&_Rollup.TransactOpts)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_Rollup *RollupTransactor) ReduceDeposit(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "reduceDeposit", target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_Rollup *RollupSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ReduceDeposit(&_Rollup.TransactOpts, target)
}

// ReduceDeposit is a paid mutator transaction binding the contract method 0x1e83d30f.
//
// Solidity: function reduceDeposit(uint256 target) returns()
func (_Rollup *RollupTransactorSession) ReduceDeposit(target *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.ReduceDeposit(&_Rollup.TransactOpts, target)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_Rollup *RollupTransactor) RejectNextNode(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "rejectNextNode", stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_Rollup *RollupSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RejectNextNode(&_Rollup.TransactOpts, stakerAddress)
}

// RejectNextNode is a paid mutator transaction binding the contract method 0x6b94c33b.
//
// Solidity: function rejectNextNode(address stakerAddress) returns()
func (_Rollup *RollupTransactorSession) RejectNextNode(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RejectNextNode(&_Rollup.TransactOpts, stakerAddress)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_Rollup *RollupTransactor) RemoveOldOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeOldOutbox", _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_Rollup *RollupSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldOutbox(&_Rollup.TransactOpts, _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_Rollup *RollupTransactorSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldOutbox(&_Rollup.TransactOpts, _outbox)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupTransactor) RemoveOldZombies(opts *bind.TransactOpts, startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeOldZombies", startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldZombies(&_Rollup.TransactOpts, startIndex)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 startIndex) returns()
func (_Rollup *RollupTransactorSession) RemoveOldZombies(startIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveOldZombies(&_Rollup.TransactOpts, startIndex)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupTransactor) RemoveZombie(opts *bind.TransactOpts, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeZombie", zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveZombie(&_Rollup.TransactOpts, zombieNum, maxNodes)
}

// RemoveZombie is a paid mutator transaction binding the contract method 0x7e2d2155.
//
// Solidity: function removeZombie(uint256 zombieNum, uint256 maxNodes) returns()
func (_Rollup *RollupTransactorSession) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveZombie(&_Rollup.TransactOpts, zombieNum, maxNodes)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Rollup *RollupTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Rollup *RollupSession) Resume() (*types.Transaction, error) {
	return _Rollup.Contract.Resume(&_Rollup.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Rollup *RollupTransactorSession) Resume() (*types.Transaction, error) {
	return _Rollup.Contract.Resume(&_Rollup.TransactOpts)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupTransactor) ReturnOldDeposit(opts *bind.TransactOpts, stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "returnOldDeposit", stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ReturnOldDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// ReturnOldDeposit is a paid mutator transaction binding the contract method 0x7427be51.
//
// Solidity: function returnOldDeposit(address stakerAddress) returns()
func (_Rollup *RollupTransactorSession) ReturnOldDeposit(stakerAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ReturnOldDeposit(&_Rollup.TransactOpts, stakerAddress)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_Rollup *RollupTransactor) SetInbox(opts *bind.TransactOpts, _inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "setInbox", _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_Rollup *RollupSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _Rollup.Contract.SetInbox(&_Rollup.TransactOpts, _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_Rollup *RollupTransactorSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _Rollup.Contract.SetInbox(&_Rollup.TransactOpts, _inbox, _enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_Rollup *RollupTransactor) SetOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "setOutbox", _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_Rollup *RollupSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.SetOutbox(&_Rollup.TransactOpts, _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_Rollup *RollupTransactorSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.SetOutbox(&_Rollup.TransactOpts, _outbox)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_Rollup *RollupTransactor) StakeOnExistingNode(opts *bind.TransactOpts, nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "stakeOnExistingNode", nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_Rollup *RollupSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnExistingNode(&_Rollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnExistingNode is a paid mutator transaction binding the contract method 0x414f23fe.
//
// Solidity: function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) returns()
func (_Rollup *RollupTransactorSession) StakeOnExistingNode(nodeNum *big.Int, nodeHash [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnExistingNode(&_Rollup.TransactOpts, nodeNum, nodeHash)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_Rollup *RollupTransactor) StakeOnNewNode(opts *bind.TransactOpts, expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "stakeOnNewNode", expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_Rollup *RollupSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnNewNode(&_Rollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// StakeOnNewNode is a paid mutator transaction binding the contract method 0x3fe38627.
//
// Solidity: function stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, bytes sequencerBatchProof) returns()
func (_Rollup *RollupTransactorSession) StakeOnNewNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, sequencerBatchProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.StakeOnNewNode(&_Rollup.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, beforeProposedBlock, beforeInboxMaxCount, sequencerBatchProof)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_Rollup *RollupTransactor) UpgradeImplementation(opts *bind.TransactOpts, _newRollup common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "upgradeImplementation", _newRollup)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_Rollup *RollupSession) UpgradeImplementation(_newRollup common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.UpgradeImplementation(&_Rollup.TransactOpts, _newRollup)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newRollup) returns()
func (_Rollup *RollupTransactorSession) UpgradeImplementation(_newRollup common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.UpgradeImplementation(&_Rollup.TransactOpts, _newRollup)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_Rollup *RollupTransactor) UpgradeImplementationAndCall(opts *bind.TransactOpts, _newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "upgradeImplementationAndCall", _newRollup, _data)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_Rollup *RollupSession) UpgradeImplementationAndCall(_newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _Rollup.Contract.UpgradeImplementationAndCall(&_Rollup.TransactOpts, _newRollup, _data)
}

// UpgradeImplementationAndCall is a paid mutator transaction binding the contract method 0xbf5ddcb1.
//
// Solidity: function upgradeImplementationAndCall(address _newRollup, bytes _data) returns()
func (_Rollup *RollupTransactorSession) UpgradeImplementationAndCall(_newRollup common.Address, _data []byte) (*types.Transaction, error) {
	return _Rollup.Contract.UpgradeImplementationAndCall(&_Rollup.TransactOpts, _newRollup, _data)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_Rollup *RollupTransactor) WithdrawStakerFunds(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "withdrawStakerFunds", destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_Rollup *RollupSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStakerFunds(&_Rollup.TransactOpts, destination)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x81fbc98a.
//
// Solidity: function withdrawStakerFunds(address destination) returns(uint256)
func (_Rollup *RollupTransactorSession) WithdrawStakerFunds(destination common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStakerFunds(&_Rollup.TransactOpts, destination)
}

// RollupNodeConfirmedIterator is returned from FilterNodeConfirmed and is used to iterate over the raw logs and unpacked data for NodeConfirmed events raised by the Rollup contract.
type RollupNodeConfirmedIterator struct {
	Event *RollupNodeConfirmed // Event containing the contract specifics and raw log

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
func (it *RollupNodeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNodeConfirmed)
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
		it.Event = new(RollupNodeConfirmed)
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
func (it *RollupNodeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNodeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNodeConfirmed represents a NodeConfirmed event raised by the Rollup contract.
type RollupNodeConfirmed struct {
	NodeNum        *big.Int
	AfterSendAcc   [32]byte
	AfterSendCount *big.Int
	AfterLogAcc    [32]byte
	AfterLogCount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeConfirmed is a free log retrieval operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_Rollup *RollupFilterer) FilterNodeConfirmed(opts *bind.FilterOpts, nodeNum []*big.Int) (*RollupNodeConfirmedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupNodeConfirmedIterator{contract: _Rollup.contract, event: "NodeConfirmed", logs: logs, sub: sub}, nil
}

// WatchNodeConfirmed is a free log subscription operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_Rollup *RollupFilterer) WatchNodeConfirmed(opts *bind.WatchOpts, sink chan<- *RollupNodeConfirmed, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNodeConfirmed)
				if err := _Rollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
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

// ParseNodeConfirmed is a log parse operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_Rollup *RollupFilterer) ParseNodeConfirmed(log types.Log) (*RollupNodeConfirmed, error) {
	event := new(RollupNodeConfirmed)
	if err := _Rollup.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the Rollup contract.
type RollupNodeCreatedIterator struct {
	Event *RollupNodeCreated // Event containing the contract specifics and raw log

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
func (it *RollupNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNodeCreated)
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
		it.Event = new(RollupNodeCreated)
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
func (it *RollupNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNodeCreated represents a NodeCreated event raised by the Rollup contract.
type RollupNodeCreated struct {
	NodeNum                 *big.Int
	ParentNodeHash          [32]byte
	NodeHash                [32]byte
	ExecutionHash           [32]byte
	InboxMaxCount           *big.Int
	AfterInboxBatchEndCount *big.Int
	AfterInboxBatchAcc      [32]byte
	AssertionBytes32Fields  [2][3][32]byte
	AssertionIntFields      [2][4]*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_Rollup *RollupFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int, parentNodeHash [][32]byte) (*RollupNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupNodeCreatedIterator{contract: _Rollup.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_Rollup *RollupFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *RollupNodeCreated, nodeNum []*big.Int, parentNodeHash [][32]byte) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNodeCreated)
				if err := _Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
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

// ParseNodeCreated is a log parse operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_Rollup *RollupFilterer) ParseNodeCreated(log types.Log) (*RollupNodeCreated, error) {
	event := new(RollupNodeCreated)
	if err := _Rollup.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupNodeRejectedIterator is returned from FilterNodeRejected and is used to iterate over the raw logs and unpacked data for NodeRejected events raised by the Rollup contract.
type RollupNodeRejectedIterator struct {
	Event *RollupNodeRejected // Event containing the contract specifics and raw log

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
func (it *RollupNodeRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNodeRejected)
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
		it.Event = new(RollupNodeRejected)
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
func (it *RollupNodeRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNodeRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNodeRejected represents a NodeRejected event raised by the Rollup contract.
type RollupNodeRejected struct {
	NodeNum *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeRejected is a free log retrieval operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_Rollup *RollupFilterer) FilterNodeRejected(opts *bind.FilterOpts, nodeNum []*big.Int) (*RollupNodeRejectedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupNodeRejectedIterator{contract: _Rollup.contract, event: "NodeRejected", logs: logs, sub: sub}, nil
}

// WatchNodeRejected is a free log subscription operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_Rollup *RollupFilterer) WatchNodeRejected(opts *bind.WatchOpts, sink chan<- *RollupNodeRejected, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNodeRejected)
				if err := _Rollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
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

// ParseNodeRejected is a log parse operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_Rollup *RollupFilterer) ParseNodeRejected(log types.Log) (*RollupNodeRejected, error) {
	event := new(RollupNodeRejected)
	if err := _Rollup.contract.UnpackLog(event, "NodeRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupNodesDestroyedIterator is returned from FilterNodesDestroyed and is used to iterate over the raw logs and unpacked data for NodesDestroyed events raised by the Rollup contract.
type RollupNodesDestroyedIterator struct {
	Event *RollupNodesDestroyed // Event containing the contract specifics and raw log

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
func (it *RollupNodesDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNodesDestroyed)
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
		it.Event = new(RollupNodesDestroyed)
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
func (it *RollupNodesDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNodesDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNodesDestroyed represents a NodesDestroyed event raised by the Rollup contract.
type RollupNodesDestroyed struct {
	StartNode *big.Int
	EndNode   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodesDestroyed is a free log retrieval operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_Rollup *RollupFilterer) FilterNodesDestroyed(opts *bind.FilterOpts, startNode []*big.Int, endNode []*big.Int) (*RollupNodesDestroyedIterator, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return &RollupNodesDestroyedIterator{contract: _Rollup.contract, event: "NodesDestroyed", logs: logs, sub: sub}, nil
}

// WatchNodesDestroyed is a free log subscription operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_Rollup *RollupFilterer) WatchNodesDestroyed(opts *bind.WatchOpts, sink chan<- *RollupNodesDestroyed, startNode []*big.Int, endNode []*big.Int) (event.Subscription, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNodesDestroyed)
				if err := _Rollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
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

// ParseNodesDestroyed is a log parse operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_Rollup *RollupFilterer) ParseNodesDestroyed(log types.Log) (*RollupNodesDestroyed, error) {
	event := new(RollupNodesDestroyed)
	if err := _Rollup.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Rollup contract.
type RollupPausedIterator struct {
	Event *RollupPaused // Event containing the contract specifics and raw log

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
func (it *RollupPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupPaused)
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
		it.Event = new(RollupPaused)
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
func (it *RollupPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupPaused represents a Paused event raised by the Rollup contract.
type RollupPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) FilterPaused(opts *bind.FilterOpts) (*RollupPausedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RollupPausedIterator{contract: _Rollup.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RollupPaused) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupPaused)
				if err := _Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) ParsePaused(log types.Log) (*RollupPaused, error) {
	event := new(RollupPaused)
	if err := _Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the Rollup contract.
type RollupRollupChallengeStartedIterator struct {
	Event *RollupRollupChallengeStarted // Event containing the contract specifics and raw log

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
func (it *RollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRollupChallengeStarted)
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
		it.Event = new(RollupRollupChallengeStarted)
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
func (it *RollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the Rollup contract.
type RollupRollupChallengeStarted struct {
	ChallengeContract common.Address
	Asserter          common.Address
	Challenger        common.Address
	ChallengedNode    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_Rollup *RollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts, challengeContract []common.Address) (*RollupRollupChallengeStartedIterator, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return &RollupRollupChallengeStartedIterator{contract: _Rollup.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_Rollup *RollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *RollupRollupChallengeStarted, challengeContract []common.Address) (event.Subscription, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRollupChallengeStarted)
				if err := _Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
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

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_Rollup *RollupFilterer) ParseRollupChallengeStarted(log types.Log) (*RollupRollupChallengeStarted, error) {
	event := new(RollupRollupChallengeStarted)
	if err := _Rollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the Rollup contract.
type RollupRollupCreatedIterator struct {
	Event *RollupRollupCreated // Event containing the contract specifics and raw log

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
func (it *RollupRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRollupCreated)
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
		it.Event = new(RollupRollupCreated)
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
func (it *RollupRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRollupCreated represents a RollupCreated event raised by the Rollup contract.
type RollupRollupCreated struct {
	MachineHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_Rollup *RollupFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*RollupRollupCreatedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &RollupRollupCreatedIterator{contract: _Rollup.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_Rollup *RollupFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupRollupCreated) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRollupCreated)
				if err := _Rollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_Rollup *RollupFilterer) ParseRollupCreated(log types.Log) (*RollupRollupCreated, error) {
	event := new(RollupRollupCreated)
	if err := _Rollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupStakerReassignedIterator is returned from FilterStakerReassigned and is used to iterate over the raw logs and unpacked data for StakerReassigned events raised by the Rollup contract.
type RollupStakerReassignedIterator struct {
	Event *RollupStakerReassigned // Event containing the contract specifics and raw log

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
func (it *RollupStakerReassignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupStakerReassigned)
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
		it.Event = new(RollupStakerReassigned)
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
func (it *RollupStakerReassignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupStakerReassignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupStakerReassigned represents a StakerReassigned event raised by the Rollup contract.
type RollupStakerReassigned struct {
	Staker  common.Address
	NewNode *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakerReassigned is a free log retrieval operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_Rollup *RollupFilterer) FilterStakerReassigned(opts *bind.FilterOpts, staker []common.Address) (*RollupStakerReassignedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return &RollupStakerReassignedIterator{contract: _Rollup.contract, event: "StakerReassigned", logs: logs, sub: sub}, nil
}

// WatchStakerReassigned is a free log subscription operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_Rollup *RollupFilterer) WatchStakerReassigned(opts *bind.WatchOpts, sink chan<- *RollupStakerReassigned, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupStakerReassigned)
				if err := _Rollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
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

// ParseStakerReassigned is a log parse operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_Rollup *RollupFilterer) ParseStakerReassigned(log types.Log) (*RollupStakerReassigned, error) {
	event := new(RollupStakerReassigned)
	if err := _Rollup.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Rollup contract.
type RollupUnpausedIterator struct {
	Event *RollupUnpaused // Event containing the contract specifics and raw log

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
func (it *RollupUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUnpaused)
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
		it.Event = new(RollupUnpaused)
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
func (it *RollupUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUnpaused represents a Unpaused event raised by the Rollup contract.
type RollupUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RollupUnpausedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RollupUnpausedIterator{contract: _Rollup.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RollupUnpaused) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUnpaused)
				if err := _Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) ParseUnpaused(log types.Log) (*RollupUnpaused, error) {
	event := new(RollupUnpaused)
	if err := _Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCoreABI is the input ABI used to generate the binding from.
const RollupCoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupCoreBin is the compiled bytecode used for deploying new contracts.
var RollupCoreBin = "0x608060405234801561001057600080fd5b50610570806100206000396000f3fe608060405234801561001057600080fd5b50600436106100e65760003560e01c80632f30cabd146100eb5780633e96576e146101235780634f0f4aa9146101495780636177fd181461018257806362a82d7d146101bc57806363721d6b146101d957806365f7f80d146101e157806369fd251c146101e95780637ba9534a1461020f5780638640ce5f1461021757806391c657e81461021f578063d01e660214610245578063d735e21d14610262578063dff697871461026a578063e8bd492214610272578063ef40a670146102ce578063f33e1fac146102f4578063f8d1f19414610311575b600080fd5b6101116004803603602081101561010157600080fd5b50356001600160a01b031661032e565b60408051918252519081900360200190f35b6101116004803603602081101561013957600080fd5b50356001600160a01b031661034d565b6101666004803603602081101561015f57600080fd5b503561036b565b604080516001600160a01b039092168252519081900360200190f35b6101a86004803603602081101561019857600080fd5b50356001600160a01b0316610386565b604080519115158252519081900360200190f35b610166600480360360208110156101d257600080fd5b50356103ae565b6101116103d8565b6101116103de565b610166600480360360208110156101ff57600080fd5b50356001600160a01b03166103e4565b610111610405565b61011161040b565b6101a86004803603602081101561023557600080fd5b50356001600160a01b0316610411565b6101666004803603602081101561025b57600080fd5b503561046b565b61011161049a565b6101116104a0565b6102986004803603602081101561028857600080fd5b50356001600160a01b03166104a6565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b610111600480360360208110156102e457600080fd5b50356001600160a01b03166104e2565b6101116004803603602081101561030a57600080fd5b5035610500565b6101116004803603602081101561032757600080fd5b5035610528565b6001600160a01b0381166000908152600960205260409020545b919050565b6001600160a01b031660009081526007602052604090206001015490565b6000908152600460205260409020546001600160a01b031690565b6001600160a01b0316600090815260076020526040902060030154600160a01b900460ff1690565b6000600682815481106103bd57fe5b6000918252602090912001546001600160a01b031692915050565b60085490565b60005490565b6001600160a01b039081166000908152600760205260409020600301541690565b60025490565b60035490565b6000805b600854811015610462576008818154811061042c57fe5b60009182526020909120600290910201546001600160a01b038481169116141561045a576001915050610348565b600101610415565b50600092915050565b60006008828154811061047a57fe5b60009182526020909120600290910201546001600160a01b031692915050565b60015490565b60065490565b6007602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6001600160a01b031660009081526007602052604090206002015490565b60006008828154811061050f57fe5b9060005260206000209060020201600101549050919050565b6000908152600560205260409020549056fea264697066735822122097938160b753385047689b820b8399c4620753e0cf7c5c5e4c9d3555f748a7b564736f6c634300060c0033"

// DeployRollupCore deploys a new Ethereum contract, binding an instance of RollupCore to it.
func DeployRollupCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupCore, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCore{RollupCoreCaller: RollupCoreCaller{contract: contract}, RollupCoreTransactor: RollupCoreTransactor{contract: contract}, RollupCoreFilterer: RollupCoreFilterer{contract: contract}}, nil
}

// RollupCore is an auto generated Go binding around an Ethereum contract.
type RollupCore struct {
	RollupCoreCaller     // Read-only binding to the contract
	RollupCoreTransactor // Write-only binding to the contract
	RollupCoreFilterer   // Log filterer for contract events
}

// RollupCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCoreSession struct {
	Contract     *RollupCore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCoreCallerSession struct {
	Contract *RollupCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCoreTransactorSession struct {
	Contract     *RollupCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCoreRaw struct {
	Contract *RollupCore // Generic contract binding to access the raw methods on
}

// RollupCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCoreCallerRaw struct {
	Contract *RollupCoreCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCoreTransactorRaw struct {
	Contract *RollupCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCore creates a new instance of RollupCore, bound to a specific deployed contract.
func NewRollupCore(address common.Address, backend bind.ContractBackend) (*RollupCore, error) {
	contract, err := bindRollupCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCore{RollupCoreCaller: RollupCoreCaller{contract: contract}, RollupCoreTransactor: RollupCoreTransactor{contract: contract}, RollupCoreFilterer: RollupCoreFilterer{contract: contract}}, nil
}

// NewRollupCoreCaller creates a new read-only instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreCaller(address common.Address, caller bind.ContractCaller) (*RollupCoreCaller, error) {
	contract, err := bindRollupCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCoreCaller{contract: contract}, nil
}

// NewRollupCoreTransactor creates a new write-only instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCoreTransactor, error) {
	contract, err := bindRollupCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCoreTransactor{contract: contract}, nil
}

// NewRollupCoreFilterer creates a new log filterer instance of RollupCore, bound to a specific deployed contract.
func NewRollupCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCoreFilterer, error) {
	contract, err := bindRollupCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCoreFilterer{contract: contract}, nil
}

// bindRollupCore binds a generic wrapper to an already deployed contract.
func bindRollupCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCore *RollupCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCore.Contract.RollupCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCore *RollupCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCore.Contract.RollupCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCore *RollupCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCore.Contract.RollupCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCore *RollupCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCore *RollupCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCore *RollupCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCore.Contract.contract.Transact(opts, method, params...)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "_stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _RollupCore.Contract.StakerMap(&_RollupCore.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupCore *RollupCoreCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _RollupCore.Contract.StakerMap(&_RollupCore.CallOpts, arg0)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.AmountStaked(&_RollupCore.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.AmountStaked(&_RollupCore.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _RollupCore.Contract.CurrentChallenge(&_RollupCore.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupCore *RollupCoreCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _RollupCore.Contract.CurrentChallenge(&_RollupCore.CallOpts, staker)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreSession) FirstUnresolvedNode() (*big.Int, error) {
	return _RollupCore.Contract.FirstUnresolvedNode(&_RollupCore.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _RollupCore.Contract.FirstUnresolvedNode(&_RollupCore.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetNode(&_RollupCore.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetNode(&_RollupCore.CallOpts, nodeNum)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _RollupCore.Contract.GetNodeHash(&_RollupCore.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupCore *RollupCoreCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _RollupCore.Contract.GetNodeHash(&_RollupCore.CallOpts, index)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetStakerAddress(&_RollupCore.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.GetStakerAddress(&_RollupCore.CallOpts, stakerNum)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreSession) IsStaked(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsStaked(&_RollupCore.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupCore *RollupCoreCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsStaked(&_RollupCore.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupCore *RollupCoreCaller) IsZombie(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "isZombie", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupCore *RollupCoreSession) IsZombie(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsZombie(&_RollupCore.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupCore *RollupCoreCallerSession) IsZombie(staker common.Address) (bool, error) {
	return _RollupCore.Contract.IsZombie(&_RollupCore.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreSession) LastStakeBlock() (*big.Int, error) {
	return _RollupCore.Contract.LastStakeBlock(&_RollupCore.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LastStakeBlock() (*big.Int, error) {
	return _RollupCore.Contract.LastStakeBlock(&_RollupCore.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestConfirmed() (*big.Int, error) {
	return _RollupCore.Contract.LatestConfirmed(&_RollupCore.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestConfirmed() (*big.Int, error) {
	return _RollupCore.Contract.LatestConfirmed(&_RollupCore.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestNodeCreated() (*big.Int, error) {
	return _RollupCore.Contract.LatestNodeCreated(&_RollupCore.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _RollupCore.Contract.LatestNodeCreated(&_RollupCore.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.LatestStakedNode(&_RollupCore.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _RollupCore.Contract.LatestStakedNode(&_RollupCore.CallOpts, staker)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreSession) StakerCount() (*big.Int, error) {
	return _RollupCore.Contract.StakerCount(&_RollupCore.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) StakerCount() (*big.Int, error) {
	return _RollupCore.Contract.StakerCount(&_RollupCore.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _RollupCore.Contract.WithdrawableFunds(&_RollupCore.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _RollupCore.Contract.WithdrawableFunds(&_RollupCore.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.ZombieAddress(&_RollupCore.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupCore *RollupCoreCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _RollupCore.Contract.ZombieAddress(&_RollupCore.CallOpts, zombieNum)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreSession) ZombieCount() (*big.Int, error) {
	return _RollupCore.Contract.ZombieCount(&_RollupCore.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) ZombieCount() (*big.Int, error) {
	return _RollupCore.Contract.ZombieCount(&_RollupCore.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RollupCore.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _RollupCore.Contract.ZombieLatestStakedNode(&_RollupCore.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupCore *RollupCoreCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _RollupCore.Contract.ZombieLatestStakedNode(&_RollupCore.CallOpts, zombieNum)
}

// RollupEventBridgeABI is the input ABI used to generate the binding from.
const RollupEventBridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"claimNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"nodeConfirmed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"}],\"name\":\"nodeCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"nodeRejected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraConfig\",\"type\":\"bytes\"}],\"name\":\"rollupInitialized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"stakeCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupEventBridgeBin is the compiled bytecode used for deploying new contracts.
var RollupEventBridgeBin = "0x608060405234801561001057600080fd5b50604051610a0a380380610a0a8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556109908061007a6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806316b9109b1461006757806330a826b41461008657806364126c7c146100a35780638b8ca199146100cf578063b0f2af2914610107578063f03c04a5146101a6575b600080fd5b6100846004803603602081101561007d57600080fd5b50356101d2565b005b6100846004803603602081101561009c57600080fd5b5035610253565b610084600480360360408110156100b957600080fd5b50803590602001356001600160a01b03166102d1565b610084600480360360808110156100e557600080fd5b50803590602081013590604081013590606001356001600160a01b03166104f8565b610084600480360360e081101561011d57600080fd5b8135916020810135916040820135916060810135916001600160a01b03608083013581169260a08101359091169181019060e0810160c0820135600160201b81111561016857600080fd5b82018360208201111561017a57600080fd5b803590602001918460018302840111600160201b8311171561019b57600080fd5b509092509050610599565b610084600480360360408110156101bc57600080fd5b506001600160a01b038135169060200135610786565b6001546001600160a01b0316331461021f576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f81b6020820152602180820184905282518083039091018152604190910190915261025090610820565b50565b6001546001600160a01b031633146102a0576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f91b6020820152602180820184905282518083039091018152604190910190915261025090610820565b6001546001600160a01b0316331461031e576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60015460408051634f0f4aa960e01b81526004810185905290516001600160a01b03909216916000918391634f0f4aa991602480820192602092909190829003018186803b15801561036f57600080fd5b505afa158015610383573d6000803e3d6000fd5b505050506040513d602081101561039957600080fd5b5051604080516348b4573960e11b81526001600160a01b038681166004830152915192935090831691639168ae7291602480820192602092909190829003018186803b1580156103e857600080fd5b505afa1580156103fc573d6000803e3d6000fd5b505050506040513d602081101561041257600080fd5b5051610452576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b816001600160a01b0316632b2af0ab856040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561049657600080fd5b505afa1580156104aa573d6000803e3d6000fd5b505060408051600160fa1b6020820152602181018890526001600160a01b0387166041808301919091528251808303909101815260619091019091526104f292509050610820565b50505050565b6001546001600160a01b03163314610545576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600060208201526021810186905260418101859052436061820152608181018490526001600160a01b03831660a1808301919091528251808303909101815260c19091019091526104f290610820565b6001546001600160a01b031633146105e6576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b6060888888888860601b60601c6001600160a01b03168860601b60601c6001600160a01b03168888604051602001808981526020018881526020018781526020018681526020018581526020018481526020018383808284376040805191909301818103601f190182528084526000805483516020808601919091206302bbfad160e01b855260048086015233602486015260448501529551939f50909d506001600160a01b03169b506302bbfad19a5060648082019a509398509096508690039091019350849250899150889050803b1580156106c357600080fd5b505af11580156106d7573d6000803e3d6000fd5b505050506040513d60208110156106ed57600080fd5b50516040805160208082528551828201528551939450849360008051602061093b833981519152938793928392918301919085019080838360005b83811015610740578181015183820152602001610728565b50505050905090810190601f16801561076d5780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050505050505050565b6001546001600160a01b031633146107d3576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600360f81b60208201526001600160a01b0384166021820152604181018390524360618083019190915282518083039091018152608190910190915261081c90610820565b5050565b600080548251602080850191909120604080516302bbfad160e01b8152600860048201523360248201526044810192909252516001600160a01b03909316936302bbfad193606480840194939192918390030190829087803b15801561088557600080fd5b505af1158015610899573d6000803e3d6000fd5b505050506040513d60208110156108af57600080fd5b5051604080516020808252845182820152845160008051602061093b833981519152938693928392918301919085019080838360005b838110156108fd5781810151838201526020016108e5565b50505050905090810190601f16801561092a5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25056feff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60ba2646970667358221220aac8822573f2c56c1652269096a9948db510154eb161f4f64661f12fddb67c9864736f6c634300060c0033"

// DeployRollupEventBridge deploys a new Ethereum contract, binding an instance of RollupEventBridge to it.
func DeployRollupEventBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _rollup common.Address) (common.Address, *types.Transaction, *RollupEventBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupEventBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupEventBridgeBin), backend, _bridge, _rollup)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupEventBridge{RollupEventBridgeCaller: RollupEventBridgeCaller{contract: contract}, RollupEventBridgeTransactor: RollupEventBridgeTransactor{contract: contract}, RollupEventBridgeFilterer: RollupEventBridgeFilterer{contract: contract}}, nil
}

// RollupEventBridge is an auto generated Go binding around an Ethereum contract.
type RollupEventBridge struct {
	RollupEventBridgeCaller     // Read-only binding to the contract
	RollupEventBridgeTransactor // Write-only binding to the contract
	RollupEventBridgeFilterer   // Log filterer for contract events
}

// RollupEventBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupEventBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupEventBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupEventBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupEventBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupEventBridgeSession struct {
	Contract     *RollupEventBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RollupEventBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupEventBridgeCallerSession struct {
	Contract *RollupEventBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RollupEventBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupEventBridgeTransactorSession struct {
	Contract     *RollupEventBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RollupEventBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupEventBridgeRaw struct {
	Contract *RollupEventBridge // Generic contract binding to access the raw methods on
}

// RollupEventBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupEventBridgeCallerRaw struct {
	Contract *RollupEventBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// RollupEventBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupEventBridgeTransactorRaw struct {
	Contract *RollupEventBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupEventBridge creates a new instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridge(address common.Address, backend bind.ContractBackend) (*RollupEventBridge, error) {
	contract, err := bindRollupEventBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridge{RollupEventBridgeCaller: RollupEventBridgeCaller{contract: contract}, RollupEventBridgeTransactor: RollupEventBridgeTransactor{contract: contract}, RollupEventBridgeFilterer: RollupEventBridgeFilterer{contract: contract}}, nil
}

// NewRollupEventBridgeCaller creates a new read-only instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeCaller(address common.Address, caller bind.ContractCaller) (*RollupEventBridgeCaller, error) {
	contract, err := bindRollupEventBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeCaller{contract: contract}, nil
}

// NewRollupEventBridgeTransactor creates a new write-only instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupEventBridgeTransactor, error) {
	contract, err := bindRollupEventBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeTransactor{contract: contract}, nil
}

// NewRollupEventBridgeFilterer creates a new log filterer instance of RollupEventBridge, bound to a specific deployed contract.
func NewRollupEventBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupEventBridgeFilterer, error) {
	contract, err := bindRollupEventBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeFilterer{contract: contract}, nil
}

// bindRollupEventBridge binds a generic wrapper to an already deployed contract.
func bindRollupEventBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupEventBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupEventBridge *RollupEventBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupEventBridge.Contract.RollupEventBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupEventBridge *RollupEventBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupEventBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupEventBridge *RollupEventBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupEventBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupEventBridge *RollupEventBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupEventBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupEventBridge *RollupEventBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupEventBridge *RollupEventBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.contract.Transact(opts, method, params...)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) ClaimNode(opts *bind.TransactOpts, nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "claimNode", nodeNum, staker)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeSession) ClaimNode(nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.ClaimNode(&_RollupEventBridge.TransactOpts, nodeNum, staker)
}

// ClaimNode is a paid mutator transaction binding the contract method 0x64126c7c.
//
// Solidity: function claimNode(uint256 nodeNum, address staker) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) ClaimNode(nodeNum *big.Int, staker common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.ClaimNode(&_RollupEventBridge.TransactOpts, nodeNum, staker)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeConfirmed(opts *bind.TransactOpts, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeConfirmed", nodeNum)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeConfirmed(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeConfirmed(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeConfirmed is a paid mutator transaction binding the contract method 0x16b9109b.
//
// Solidity: function nodeConfirmed(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeConfirmed(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeConfirmed(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeCreated(opts *bind.TransactOpts, nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeCreated", nodeNum, prev, deadline, asserter)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeCreated(nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeCreated(&_RollupEventBridge.TransactOpts, nodeNum, prev, deadline, asserter)
}

// NodeCreated is a paid mutator transaction binding the contract method 0x8b8ca199.
//
// Solidity: function nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeCreated(nodeNum *big.Int, prev *big.Int, deadline *big.Int, asserter common.Address) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeCreated(&_RollupEventBridge.TransactOpts, nodeNum, prev, deadline, asserter)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) NodeRejected(opts *bind.TransactOpts, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "nodeRejected", nodeNum)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) NodeRejected(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeRejected(&_RollupEventBridge.TransactOpts, nodeNum)
}

// NodeRejected is a paid mutator transaction binding the contract method 0x30a826b4.
//
// Solidity: function nodeRejected(uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) NodeRejected(nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.NodeRejected(&_RollupEventBridge.TransactOpts, nodeNum)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) RollupInitialized(opts *bind.TransactOpts, confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "rollupInitialized", confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeSession) RollupInitialized(confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupInitialized(&_RollupEventBridge.TransactOpts, confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// RollupInitialized is a paid mutator transaction binding the contract method 0xb0f2af29.
//
// Solidity: function rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) RollupInitialized(confirmPeriodBlocks *big.Int, extraChallengeTimeBlocks *big.Int, arbGasSpeedLimitPerBlock *big.Int, baseStake *big.Int, stakeToken common.Address, owner common.Address, extraConfig []byte) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.RollupInitialized(&_RollupEventBridge.TransactOpts, confirmPeriodBlocks, extraChallengeTimeBlocks, arbGasSpeedLimitPerBlock, baseStake, stakeToken, owner, extraConfig)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactor) StakeCreated(opts *bind.TransactOpts, staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.contract.Transact(opts, "stakeCreated", staker, nodeNum)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeSession) StakeCreated(staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.StakeCreated(&_RollupEventBridge.TransactOpts, staker, nodeNum)
}

// StakeCreated is a paid mutator transaction binding the contract method 0xf03c04a5.
//
// Solidity: function stakeCreated(address staker, uint256 nodeNum) returns()
func (_RollupEventBridge *RollupEventBridgeTransactorSession) StakeCreated(staker common.Address, nodeNum *big.Int) (*types.Transaction, error) {
	return _RollupEventBridge.Contract.StakeCreated(&_RollupEventBridge.TransactOpts, staker, nodeNum)
}

// RollupEventBridgeInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredIterator struct {
	Event *RollupEventBridgeInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupEventBridgeInboxMessageDelivered)
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
		it.Event = new(RollupEventBridgeInboxMessageDelivered)
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
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupEventBridgeInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupEventBridgeInboxMessageDelivered represents a InboxMessageDelivered event raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*RollupEventBridgeInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeInboxMessageDeliveredIterator{contract: _RollupEventBridge.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *RollupEventBridgeInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupEventBridgeInboxMessageDelivered)
				if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_RollupEventBridge *RollupEventBridgeFilterer) ParseInboxMessageDelivered(log types.Log) (*RollupEventBridgeInboxMessageDelivered, error) {
	event := new(RollupEventBridgeInboxMessageDelivered)
	if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupEventBridgeInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredFromOriginIterator struct {
	Event *RollupEventBridgeInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
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
		it.Event = new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
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
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupEventBridgeInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupEventBridgeInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the RollupEventBridge contract.
type RollupEventBridgeInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*RollupEventBridgeInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupEventBridgeInboxMessageDeliveredFromOriginIterator{contract: _RollupEventBridge.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *RollupEventBridgeInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _RollupEventBridge.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
				if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_RollupEventBridge *RollupEventBridgeFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*RollupEventBridgeInboxMessageDeliveredFromOrigin, error) {
	event := new(RollupEventBridgeInboxMessageDeliveredFromOrigin)
	if err := _RollupEventBridge.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupLibABI is the input ABI used to generate the binding from.
const RollupLibABI = "[]"

// RollupLibBin is the compiled bytecode used for deploying new contracts.
var RollupLibBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d290838e51e3f6509fce4b58c83fc98085e08427722ae934c62fcac821431dc664736f6c634300060c0033"

// DeployRollupLib deploys a new Ethereum contract, binding an instance of RollupLib to it.
func DeployRollupLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupLib, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupLib{RollupLibCaller: RollupLibCaller{contract: contract}, RollupLibTransactor: RollupLibTransactor{contract: contract}, RollupLibFilterer: RollupLibFilterer{contract: contract}}, nil
}

// RollupLib is an auto generated Go binding around an Ethereum contract.
type RollupLib struct {
	RollupLibCaller     // Read-only binding to the contract
	RollupLibTransactor // Write-only binding to the contract
	RollupLibFilterer   // Log filterer for contract events
}

// RollupLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupLibSession struct {
	Contract     *RollupLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupLibCallerSession struct {
	Contract *RollupLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RollupLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupLibTransactorSession struct {
	Contract     *RollupLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RollupLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupLibRaw struct {
	Contract *RollupLib // Generic contract binding to access the raw methods on
}

// RollupLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupLibCallerRaw struct {
	Contract *RollupLibCaller // Generic read-only contract binding to access the raw methods on
}

// RollupLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupLibTransactorRaw struct {
	Contract *RollupLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupLib creates a new instance of RollupLib, bound to a specific deployed contract.
func NewRollupLib(address common.Address, backend bind.ContractBackend) (*RollupLib, error) {
	contract, err := bindRollupLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupLib{RollupLibCaller: RollupLibCaller{contract: contract}, RollupLibTransactor: RollupLibTransactor{contract: contract}, RollupLibFilterer: RollupLibFilterer{contract: contract}}, nil
}

// NewRollupLibCaller creates a new read-only instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibCaller(address common.Address, caller bind.ContractCaller) (*RollupLibCaller, error) {
	contract, err := bindRollupLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupLibCaller{contract: contract}, nil
}

// NewRollupLibTransactor creates a new write-only instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupLibTransactor, error) {
	contract, err := bindRollupLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupLibTransactor{contract: contract}, nil
}

// NewRollupLibFilterer creates a new log filterer instance of RollupLib, bound to a specific deployed contract.
func NewRollupLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupLibFilterer, error) {
	contract, err := bindRollupLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupLibFilterer{contract: contract}, nil
}

// bindRollupLib binds a generic wrapper to an already deployed contract.
func bindRollupLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupLib *RollupLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupLib.Contract.RollupLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupLib *RollupLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupLib.Contract.RollupLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupLib *RollupLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupLib.Contract.RollupLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupLib *RollupLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupLib *RollupLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupLib *RollupLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupLib.Contract.contract.Transact(opts, method, params...)
}

// ValidatorUtilsABI is the input ABI used to generate the binding from.
const ValidatorUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"areUnresolvedNodesLinear\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkDecidableNextNode\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.ConfirmType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"node1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findNodeConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findStakerConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"hasMore\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"refundableStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"requireConfirmable\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"requireRejectable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakedNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"stakerInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"timedOutChallenges\",\"outputs\":[{\"internalType\":\"contractIChallenge[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"hasMore\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorUtilsBin is the compiled bytecode used for deploying new contracts.
var ValidatorUtilsBin = "0x608060405234801561001057600080fd5b50612929806100206000396000f3fe608060405234801561001057600080fd5b50600436106100af5760003560e01c806301d9717d146100b45780630a46c1b5146100de5780631fc43bb6146100fe5780633082d0291461011357806371229340146101355780637464ae06146101555780637988ad37146101755780638f67e6bb14610188578063a8ac9cf3146101ab578063abeba4f7146101cc578063aea2f06e146101ed578063c308eaaf14610200578063e48a5f7b14610220575b600080fd5b6100c76100c2366004612566565b610243565b6040516100d59291906128b2565b60405180910390f35b6100f16100ec36600461254a565b6103cb565b6040516100d591906127b3565b61011161010c36600461254a565b6104ad565b005b610126610121366004612622565b610a2f565b6040516100d5939291906127c7565b61014861014336600461254a565b610f01565b6040516100d59190612782565b61016861016336600461254a565b61140e565b6040516100d591906126b3565b61012661018336600461259e565b611727565b61019b610196366004612566565b611849565b6040516100d5949392919061278d565b6101be6101b93660046125ee565b611a4f565b6040516100d59291906126ea565b6101df6101da3660046125ee565b611d29565b6040516100d59291906126c6565b6101486101fb36600461254a565b611eb9565b61021361020e366004612566565b6120ce565b6040516100d5919061273e565b61023361022e36600461254a565b61231d565b6040516100d594939291906128c0565b6000806000846001600160a01b0316633e96576e856040518263ffffffff1660e01b8152600401610274919061269f565b60206040518083038186803b15801561028c57600080fd5b505afa1580156102a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c49190612532565b90508061033f57846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561030457600080fd5b505afa158015610318573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061033c9190612532565b90505b604051633e347c6560e21b81526000906001600160a01b0387169063f8d1f1949061036e9085906004016128a9565b60206040518083038186803b15801561038657600080fd5b505afa15801561039a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103be9190612532565b9196919550909350505050565b604051630fe21ddb60e11b81526000903090631fc43bb6906103f190859060040161269f565b60006040518083038186803b15801561040957600080fd5b505afa92505050801561041a575060015b6104235761042b565b5060016104a8565b6040516301c48a4d60e61b8152309063712293409061044e90859060040161269f565b60206040518083038186803b15801561046657600080fd5b505afa925050508015610496575060408051601f3d908101601f1916820190925261049391810190612512565b60015b6104a2575060006104a8565b50600290505b919050565b806001600160a01b03166367425daf6040518163ffffffff1660e01b815260040160006040518083038186803b1580156104e657600080fd5b505afa1580156104fa573d6000803e3d6000fd5b505050506000816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561053957600080fd5b505afa15801561054d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105719190612532565b90506000811161059c5760405162461bcd60e51b81526004016105939061280e565b60405180910390fd5b6000826001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156105d757600080fd5b505afa1580156105eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060f9190612532565b90506000836001600160a01b0316634f0f4aa9836040518263ffffffff1660e01b815260040161063f91906128a9565b60206040518083038186803b15801561065757600080fd5b505afa15801561066b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068f91906124f6565b9050806001600160a01b03166388d221c66040518163ffffffff1660e01b815260040160006040518083038186803b1580156106ca57600080fd5b505afa1580156106de573d6000803e3d6000fd5b50505050836001600160a01b0316634f0f4aa9826001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561072a57600080fd5b505afa15801561073e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107629190612532565b6040518263ffffffff1660e01b815260040161077e91906128a9565b60206040518083038186803b15801561079657600080fd5b505afa1580156107aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ce91906124f6565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b15801561080657600080fd5b505afa15801561081a573d6000803e3d6000fd5b50505050836001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561085757600080fd5b505afa15801561086b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088f9190612532565b816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156108c857600080fd5b505afa1580156108dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109009190612532565b1461091d5760405162461bcd60e51b81526004016105939061285a565b604051630128a01960e21b81526001600160a01b038516906304a280649061094990849060040161269f565b60206040518083038186803b15801561096157600080fd5b505afa158015610975573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109999190612532565b8301816001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b1580156109d457600080fd5b505afa1580156109e8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0c9190612532565b14610a295760405162461bcd60e51b815260040161059390612832565b50505050565b600080600080876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a6e57600080fd5b505afa158015610a82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa69190612532565b90506000886001600160a01b0316634f0f4aa9896040518263ffffffff1660e01b8152600401610ad691906128a9565b60206040518083038186803b158015610aee57600080fd5b505afa158015610b02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2691906124f6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610b5e57600080fd5b505afa158015610b72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b969190612532565b90506000896001600160a01b0316634f0f4aa9896040518263ffffffff1660e01b8152600401610bc691906128a9565b60206040518083038186803b158015610bde57600080fd5b505afa158015610bf2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1691906124f6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610c4e57600080fd5b505afa158015610c62573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c869190612532565b905060005b87811015610ee857888a1415610cae5760008a8a96509650965050505050610ef7565b81831415610cc95760018a8a96509650965050505050610ef7565b8383108015610cd757508382105b15610cf057600260008096509650965050505050610ef7565b81831015610dee578198508a6001600160a01b0316634f0f4aa98a6040518263ffffffff1660e01b8152600401610d2791906128a9565b60206040518083038186803b158015610d3f57600080fd5b505afa158015610d53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7791906124f6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610daf57600080fd5b505afa158015610dc3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de79190612532565b9150610ee0565b8299508a6001600160a01b0316634f0f4aa98b6040518263ffffffff1660e01b8152600401610e1d91906128a9565b60206040518083038186803b158015610e3557600080fd5b505afa158015610e49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6d91906124f6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610ea557600080fd5b505afa158015610eb9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610edd9190612532565b92505b600101610c8b565b50600389899550955095505050505b9450945094915050565b6000816001600160a01b03166367425daf6040518163ffffffff1660e01b815260040160006040518083038186803b158015610f3c57600080fd5b505afa158015610f50573d6000803e3d6000fd5b505050506000826001600160a01b0316634f0f4aa9846001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f9e57600080fd5b505afa158015610fb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd69190612532565b6040518263ffffffff1660e01b8152600401610ff291906128a9565b60206040518083038186803b15801561100a57600080fd5b505afa15801561101e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061104291906124f6565b90506000836001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561107f57600080fd5b505afa158015611093573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b79190612532565b826001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156110f057600080fd5b505afa158015611104573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111289190612532565b149050801561140757816001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561116a57600080fd5b505afa15801561117e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111a29190612532565b4310156111c15760405162461bcd60e51b815260040161059390612880565b836001600160a01b0316634f0f4aa9836001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561120957600080fd5b505afa15801561121d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112419190612532565b6040518263ffffffff1660e01b815260040161125d91906128a9565b60206040518083038186803b15801561127557600080fd5b505afa158015611289573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ad91906124f6565b6001600160a01b0316633aa192746040518163ffffffff1660e01b815260040160006040518083038186803b1580156112e557600080fd5b505afa1580156112f9573d6000803e3d6000fd5b5050604051630128a01960e21b81526001600160a01b03871692506304a28064915061132990859060040161269f565b60206040518083038186803b15801561134157600080fd5b505afa158015611355573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113799190612532565b826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b1580156113b257600080fd5b505afa1580156113c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ea9190612532565b146114075760405162461bcd60e51b8152600401610593906127e9565b9392505050565b60606000826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561144b57600080fd5b505afa15801561145f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114839190612532565b90506060816001600160401b038111801561149d57600080fd5b506040519080825280602002602001820160405280156114c7578160200160208202803683370190505b5090506000846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561150557600080fd5b505afa158015611519573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153d9190612532565b90506000805b8481101561171c576040516362a82d7d60e01b81526000906001600160a01b038916906362a82d7d9061157a9085906004016128a9565b60206040518083038186803b15801561159257600080fd5b505afa1580156115a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ca91906124f6565b90506000886001600160a01b0316633e96576e836040518263ffffffff1660e01b81526004016115fa919061269f565b60206040518083038186803b15801561161257600080fd5b505afa158015611626573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164a9190612532565b90508481111580156116e15750604051631a7f494760e21b81526000906001600160a01b038b16906369fd251c9061168690869060040161269f565b60206040518083038186803b15801561169e57600080fd5b505afa1580156116b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d691906124f6565b6001600160a01b0316145b1561171257818685815181106116f357fe5b6001600160a01b03909216602092830291909101909101526001909301925b5050600101611543565b508252509392505050565b600080600080876001600160a01b0316633e96576e886040518263ffffffff1660e01b8152600401611759919061269f565b60206040518083038186803b15801561177157600080fd5b505afa158015611785573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117a99190612532565b90506000886001600160a01b0316633e96576e886040518263ffffffff1660e01b81526004016117d9919061269f565b60206040518083038186803b1580156117f157600080fd5b505afa158015611805573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118299190612532565b905061183789838389610a2f565b94509450945050509450945094915050565b600080600080856001600160a01b0316636177fd18866040518263ffffffff1660e01b815260040161187b919061269f565b60206040518083038186803b15801561189357600080fd5b505afa1580156118a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118cb9190612512565b604051631f4b2bb760e11b81526001600160a01b03881690633e96576e906118f790899060040161269f565b60206040518083038186803b15801561190f57600080fd5b505afa158015611923573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119479190612532565b604051630ef40a6760e41b81526001600160a01b0389169063ef40a67090611973908a9060040161269f565b60206040518083038186803b15801561198b57600080fd5b505afa15801561199f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c39190612532565b604051631a7f494760e21b81526001600160a01b038a16906369fd251c906119ef908b9060040161269f565b60206040518083038186803b158015611a0757600080fd5b505afa158015611a1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a3f91906124f6565b9299919850965090945092505050565b6060600060606000611a62878787611d29565b91509150606082516001600160401b0381118015611a7f57600080fd5b50604051908082528060200260200182016040528015611aa9578160200160208202803683370190505b5090506000805b8451811015611d1a576000858281518110611ac757fe5b6020026020010151905060008b6001600160a01b03166369fd251c836040518263ffffffff1660e01b8152600401611aff919061269f565b60206040518083038186803b158015611b1757600080fd5b505afa158015611b2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4f91906124f6565b90506001600160a01b03811615611d105760008190506000816001600160a01b031663925f9a966040518163ffffffff1660e01b815260040160206040518083038186803b158015611ba057600080fd5b505afa158015611bb4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd89190612532565b43039050816001600160a01b031663e87e35896040518163ffffffff1660e01b815260040160206040518083038186803b158015611c1557600080fd5b505afa158015611c29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c4d9190612532565b81118015611cdc5750836001600160a01b0316826001600160a01b031663bb4af0b16040518163ffffffff1660e01b815260040160206040518083038186803b158015611c9957600080fd5b505afa158015611cad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cd191906124f6565b6001600160a01b0316145b15611d0d5781878781518110611cee57fe5b6001600160a01b03909216602092830291909101909101526001909501945b50505b5050600101611ab0565b50815297909650945050505050565b6060600080856001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015611d6757600080fd5b505afa158015611d7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d9f9190612532565b90508084860111611db35750600190508383015b6060816001600160401b0381118015611dcb57600080fd5b50604051908082528060200260200182016040528015611df5578160200160208202803683370190505b50905060005b82811015611ead576040516362a82d7d60e01b81526001600160a01b038916906362a82d7d90611e31908a8501906004016128a9565b60206040518083038186803b158015611e4957600080fd5b505afa158015611e5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e8191906124f6565b828281518110611e8d57fe5b6001600160a01b0390921660209283029190910190910152600101611dfb565b50925050935093915050565b600080826001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611ef557600080fd5b505afa158015611f09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f2d9190612532565b90506000836001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611f6a57600080fd5b505afa158015611f7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fa29190612532565b90505b8181116120c4576000811180156120ac5750604051634f0f4aa960e01b81526000198201906001600160a01b03861690634f0f4aa990611fe99085906004016128a9565b60206040518083038186803b15801561200157600080fd5b505afa158015612015573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203991906124f6565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561207157600080fd5b505afa158015612085573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120a99190612532565b14155b156120bc576000925050506104a8565b600101611fa5565b5060019392505050565b60408051620186a08082526230d4208201909252606091829190602082016230d40080368337019050509050600080856001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561213657600080fd5b505afa15801561214a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061216e9190612532565b90505b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156121aa57600080fd5b505afa1580156121be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121e29190612532565b811161231357604051634f0f4aa960e01b81526000906001600160a01b03881690634f0f4aa9906122179085906004016128a9565b60206040518083038186803b15801561222f57600080fd5b505afa158015612243573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061226791906124f6565b6040516348b4573960e11b81529091506001600160a01b03821690639168ae729061229690899060040161269f565b60206040518083038186803b1580156122ae57600080fd5b505afa1580156122c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122e69190612512565b1561230a57818484815181106122f857fe5b60209081029190910101526001909201915b50600101612171565b5081529392505050565b600080600080846001600160a01b0316632e7acfa66040518163ffffffff1660e01b815260040160206040518083038186803b15801561235c57600080fd5b505afa158015612370573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123949190612532565b9350846001600160a01b031663771b2f976040518163ffffffff1660e01b815260040160206040518083038186803b1580156123cf57600080fd5b505afa1580156123e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124079190612532565b9250846001600160a01b0316635e8ef1066040518163ffffffff1660e01b815260040160206040518083038186803b15801561244257600080fd5b505afa158015612456573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061247a9190612532565b9150846001600160a01b03166376e7e23b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156124b557600080fd5b505afa1580156124c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124ed9190612532565b90509193509193565b600060208284031215612507578081fd5b8151611407816128db565b600060208284031215612523578081fd5b81518015158114611407578182fd5b600060208284031215612543578081fd5b5051919050565b60006020828403121561255b578081fd5b8135611407816128db565b60008060408385031215612578578081fd5b8235612583816128db565b91506020830135612593816128db565b809150509250929050565b600080600080608085870312156125b3578182fd5b84356125be816128db565b935060208501356125ce816128db565b925060408501356125de816128db565b9396929550929360600135925050565b600080600060608486031215612602578283fd5b833561260d816128db565b95602085013595506040909401359392505050565b60008060008060808587031215612637578384fd5b8435612642816128db565b966020860135965060408601359560600135945092505050565b6000815180845260208085019450808401835b838110156126945781516001600160a01b03168752958201959082019060010161266f565b509495945050505050565b6001600160a01b0391909116815260200190565b600060208252611407602083018461265c565b6000604082526126d9604083018561265c565b905082151560208301529392505050565b604080825283519082018190526000906020906060840190828701845b8281101561272c5781516001600160a01b031684529284019290840190600101612707565b50505093151592019190915250919050565b6020808252825182820181905260009190848201906040850190845b818110156127765783518352928401929184019160010161275a565b50909695505050505050565b901515815260200190565b9315158452602084019290925260408301526001600160a01b0316606082015260800190565b60208101600383106127c157fe5b91905290565b60608101600485106127d557fe5b938152602081019290925260409091015290565b6020808252600b908201526a4841535f5354414b45525360a81b604082015260600190565b6020808252600a90820152694e4f5f5354414b45525360b01b604082015260600190565b6020808252600e908201526d1393d517d0531317d4d51052d15160921b604082015260600190565b6020808252600c908201526b24a72b20a624a22fa82922ab60a11b604082015260600190565b6020808252600f908201526e4245464f52455f444541444c494e4560881b604082015260600190565b90815260200190565b918252602082015260400190565b93845260208401929092526040830152606082015260800190565b6001600160a01b03811681146128f057600080fd5b5056fea2646970667358221220c0b20e6fe375b6119afe7e232e03cbb61979c13fb5c0a3941b8a9d3fbfc939fe64736f6c634300060c0033"

// DeployValidatorUtils deploys a new Ethereum contract, binding an instance of ValidatorUtils to it.
func DeployValidatorUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorUtils{ValidatorUtilsCaller: ValidatorUtilsCaller{contract: contract}, ValidatorUtilsTransactor: ValidatorUtilsTransactor{contract: contract}, ValidatorUtilsFilterer: ValidatorUtilsFilterer{contract: contract}}, nil
}

// ValidatorUtils is an auto generated Go binding around an Ethereum contract.
type ValidatorUtils struct {
	ValidatorUtilsCaller     // Read-only binding to the contract
	ValidatorUtilsTransactor // Write-only binding to the contract
	ValidatorUtilsFilterer   // Log filterer for contract events
}

// ValidatorUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorUtilsSession struct {
	Contract     *ValidatorUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorUtilsCallerSession struct {
	Contract *ValidatorUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ValidatorUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorUtilsTransactorSession struct {
	Contract     *ValidatorUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ValidatorUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorUtilsRaw struct {
	Contract *ValidatorUtils // Generic contract binding to access the raw methods on
}

// ValidatorUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorUtilsCallerRaw struct {
	Contract *ValidatorUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorUtilsTransactorRaw struct {
	Contract *ValidatorUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorUtils creates a new instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtils(address common.Address, backend bind.ContractBackend) (*ValidatorUtils, error) {
	contract, err := bindValidatorUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtils{ValidatorUtilsCaller: ValidatorUtilsCaller{contract: contract}, ValidatorUtilsTransactor: ValidatorUtilsTransactor{contract: contract}, ValidatorUtilsFilterer: ValidatorUtilsFilterer{contract: contract}}, nil
}

// NewValidatorUtilsCaller creates a new read-only instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorUtilsCaller, error) {
	contract, err := bindValidatorUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsCaller{contract: contract}, nil
}

// NewValidatorUtilsTransactor creates a new write-only instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorUtilsTransactor, error) {
	contract, err := bindValidatorUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsTransactor{contract: contract}, nil
}

// NewValidatorUtilsFilterer creates a new log filterer instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorUtilsFilterer, error) {
	contract, err := bindValidatorUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsFilterer{contract: contract}, nil
}

// bindValidatorUtils binds a generic wrapper to an already deployed contract.
func bindValidatorUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorUtils *ValidatorUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorUtils.Contract.ValidatorUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorUtils *ValidatorUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.ValidatorUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorUtils *ValidatorUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.ValidatorUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorUtils *ValidatorUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorUtils *ValidatorUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorUtils *ValidatorUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.contract.Transact(opts, method, params...)
}

// AreUnresolvedNodesLinear is a free data retrieval call binding the contract method 0xaea2f06e.
//
// Solidity: function areUnresolvedNodesLinear(address rollup) view returns(bool)
func (_ValidatorUtils *ValidatorUtilsCaller) AreUnresolvedNodesLinear(opts *bind.CallOpts, rollup common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "areUnresolvedNodesLinear", rollup)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreUnresolvedNodesLinear is a free data retrieval call binding the contract method 0xaea2f06e.
//
// Solidity: function areUnresolvedNodesLinear(address rollup) view returns(bool)
func (_ValidatorUtils *ValidatorUtilsSession) AreUnresolvedNodesLinear(rollup common.Address) (bool, error) {
	return _ValidatorUtils.Contract.AreUnresolvedNodesLinear(&_ValidatorUtils.CallOpts, rollup)
}

// AreUnresolvedNodesLinear is a free data retrieval call binding the contract method 0xaea2f06e.
//
// Solidity: function areUnresolvedNodesLinear(address rollup) view returns(bool)
func (_ValidatorUtils *ValidatorUtilsCallerSession) AreUnresolvedNodesLinear(rollup common.Address) (bool, error) {
	return _ValidatorUtils.Contract.AreUnresolvedNodesLinear(&_ValidatorUtils.CallOpts, rollup)
}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x0a46c1b5.
//
// Solidity: function checkDecidableNextNode(address rollup) view returns(uint8)
func (_ValidatorUtils *ValidatorUtilsCaller) CheckDecidableNextNode(opts *bind.CallOpts, rollup common.Address) (uint8, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "checkDecidableNextNode", rollup)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x0a46c1b5.
//
// Solidity: function checkDecidableNextNode(address rollup) view returns(uint8)
func (_ValidatorUtils *ValidatorUtilsSession) CheckDecidableNextNode(rollup common.Address) (uint8, error) {
	return _ValidatorUtils.Contract.CheckDecidableNextNode(&_ValidatorUtils.CallOpts, rollup)
}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x0a46c1b5.
//
// Solidity: function checkDecidableNextNode(address rollup) view returns(uint8)
func (_ValidatorUtils *ValidatorUtilsCallerSession) CheckDecidableNextNode(rollup common.Address) (uint8, error) {
	return _ValidatorUtils.Contract.CheckDecidableNextNode(&_ValidatorUtils.CallOpts, rollup)
}

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCaller) FindNodeConflict(opts *bind.CallOpts, rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "findNodeConflict", rollup, node1, node2, maxDepth)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsSession) FindNodeConflict(rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindNodeConflict(&_ValidatorUtils.CallOpts, rollup, node1, node2, maxDepth)
}

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCallerSession) FindNodeConflict(rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindNodeConflict(&_ValidatorUtils.CallOpts, rollup, node1, node2, maxDepth)
}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCaller) FindStakerConflict(opts *bind.CallOpts, rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "findStakerConflict", rollup, staker1, staker2, maxDepth)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsSession) FindStakerConflict(rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindStakerConflict(&_ValidatorUtils.CallOpts, rollup, staker1, staker2, maxDepth)
}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCallerSession) FindStakerConflict(rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindStakerConflict(&_ValidatorUtils.CallOpts, rollup, staker1, staker2, maxDepth)
}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake)
func (_ValidatorUtils *ValidatorUtilsCaller) GetConfig(opts *bind.CallOpts, rollup common.Address) (struct {
	ConfirmPeriodBlocks      *big.Int
	ExtraChallengeTimeBlocks *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
}, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "getConfig", rollup)

	outstruct := new(struct {
		ConfirmPeriodBlocks      *big.Int
		ExtraChallengeTimeBlocks *big.Int
		ArbGasSpeedLimitPerBlock *big.Int
		BaseStake                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfirmPeriodBlocks = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExtraChallengeTimeBlocks = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ArbGasSpeedLimitPerBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BaseStake = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake)
func (_ValidatorUtils *ValidatorUtilsSession) GetConfig(rollup common.Address) (struct {
	ConfirmPeriodBlocks      *big.Int
	ExtraChallengeTimeBlocks *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
}, error) {
	return _ValidatorUtils.Contract.GetConfig(&_ValidatorUtils.CallOpts, rollup)
}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake)
func (_ValidatorUtils *ValidatorUtilsCallerSession) GetConfig(rollup common.Address) (struct {
	ConfirmPeriodBlocks      *big.Int
	ExtraChallengeTimeBlocks *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
}, error) {
	return _ValidatorUtils.Contract.GetConfig(&_ValidatorUtils.CallOpts, rollup)
}

// GetStakers is a free data retrieval call binding the contract method 0xabeba4f7.
//
// Solidity: function getStakers(address rollup, uint256 startIndex, uint256 max) view returns(address[], bool hasMore)
func (_ValidatorUtils *ValidatorUtilsCaller) GetStakers(opts *bind.CallOpts, rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, bool, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "getStakers", rollup, startIndex, max)

	if err != nil {
		return *new([]common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetStakers is a free data retrieval call binding the contract method 0xabeba4f7.
//
// Solidity: function getStakers(address rollup, uint256 startIndex, uint256 max) view returns(address[], bool hasMore)
func (_ValidatorUtils *ValidatorUtilsSession) GetStakers(rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, bool, error) {
	return _ValidatorUtils.Contract.GetStakers(&_ValidatorUtils.CallOpts, rollup, startIndex, max)
}

// GetStakers is a free data retrieval call binding the contract method 0xabeba4f7.
//
// Solidity: function getStakers(address rollup, uint256 startIndex, uint256 max) view returns(address[], bool hasMore)
func (_ValidatorUtils *ValidatorUtilsCallerSession) GetStakers(rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, bool, error) {
	return _ValidatorUtils.Contract.GetStakers(&_ValidatorUtils.CallOpts, rollup, startIndex, max)
}

// LatestStaked is a free data retrieval call binding the contract method 0x01d9717d.
//
// Solidity: function latestStaked(address rollup, address staker) view returns(uint256, bytes32)
func (_ValidatorUtils *ValidatorUtilsCaller) LatestStaked(opts *bind.CallOpts, rollup common.Address, staker common.Address) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "latestStaked", rollup, staker)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// LatestStaked is a free data retrieval call binding the contract method 0x01d9717d.
//
// Solidity: function latestStaked(address rollup, address staker) view returns(uint256, bytes32)
func (_ValidatorUtils *ValidatorUtilsSession) LatestStaked(rollup common.Address, staker common.Address) (*big.Int, [32]byte, error) {
	return _ValidatorUtils.Contract.LatestStaked(&_ValidatorUtils.CallOpts, rollup, staker)
}

// LatestStaked is a free data retrieval call binding the contract method 0x01d9717d.
//
// Solidity: function latestStaked(address rollup, address staker) view returns(uint256, bytes32)
func (_ValidatorUtils *ValidatorUtilsCallerSession) LatestStaked(rollup common.Address, staker common.Address) (*big.Int, [32]byte, error) {
	return _ValidatorUtils.Contract.LatestStaked(&_ValidatorUtils.CallOpts, rollup, staker)
}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsCaller) RefundableStakers(opts *bind.CallOpts, rollup common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "refundableStakers", rollup)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsSession) RefundableStakers(rollup common.Address) ([]common.Address, error) {
	return _ValidatorUtils.Contract.RefundableStakers(&_ValidatorUtils.CallOpts, rollup)
}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) RefundableStakers(rollup common.Address) ([]common.Address, error) {
	return _ValidatorUtils.Contract.RefundableStakers(&_ValidatorUtils.CallOpts, rollup)
}

// RequireConfirmable is a free data retrieval call binding the contract method 0x1fc43bb6.
//
// Solidity: function requireConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCaller) RequireConfirmable(opts *bind.CallOpts, rollup common.Address) error {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "requireConfirmable", rollup)

	if err != nil {
		return err
	}

	return err

}

// RequireConfirmable is a free data retrieval call binding the contract method 0x1fc43bb6.
//
// Solidity: function requireConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsSession) RequireConfirmable(rollup common.Address) error {
	return _ValidatorUtils.Contract.RequireConfirmable(&_ValidatorUtils.CallOpts, rollup)
}

// RequireConfirmable is a free data retrieval call binding the contract method 0x1fc43bb6.
//
// Solidity: function requireConfirmable(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCallerSession) RequireConfirmable(rollup common.Address) error {
	return _ValidatorUtils.Contract.RequireConfirmable(&_ValidatorUtils.CallOpts, rollup)
}

// RequireRejectable is a free data retrieval call binding the contract method 0x71229340.
//
// Solidity: function requireRejectable(address rollup) view returns(bool)
func (_ValidatorUtils *ValidatorUtilsCaller) RequireRejectable(opts *bind.CallOpts, rollup common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "requireRejectable", rollup)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireRejectable is a free data retrieval call binding the contract method 0x71229340.
//
// Solidity: function requireRejectable(address rollup) view returns(bool)
func (_ValidatorUtils *ValidatorUtilsSession) RequireRejectable(rollup common.Address) (bool, error) {
	return _ValidatorUtils.Contract.RequireRejectable(&_ValidatorUtils.CallOpts, rollup)
}

// RequireRejectable is a free data retrieval call binding the contract method 0x71229340.
//
// Solidity: function requireRejectable(address rollup) view returns(bool)
func (_ValidatorUtils *ValidatorUtilsCallerSession) RequireRejectable(rollup common.Address) (bool, error) {
	return _ValidatorUtils.Contract.RequireRejectable(&_ValidatorUtils.CallOpts, rollup)
}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCaller) StakedNodes(opts *bind.CallOpts, rollup common.Address, staker common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "stakedNodes", rollup, staker)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsSession) StakedNodes(rollup common.Address, staker common.Address) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.StakedNodes(&_ValidatorUtils.CallOpts, rollup, staker)
}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) StakedNodes(rollup common.Address, staker common.Address) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.StakedNodes(&_ValidatorUtils.CallOpts, rollup, staker)
}

// StakerInfo is a free data retrieval call binding the contract method 0x8f67e6bb.
//
// Solidity: function stakerInfo(address rollup, address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_ValidatorUtils *ValidatorUtilsCaller) StakerInfo(opts *bind.CallOpts, rollup common.Address, stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "stakerInfo", rollup, stakerAddress)

	outstruct := new(struct {
		IsStaked         bool
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStaked = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StakerInfo is a free data retrieval call binding the contract method 0x8f67e6bb.
//
// Solidity: function stakerInfo(address rollup, address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_ValidatorUtils *ValidatorUtilsSession) StakerInfo(rollup common.Address, stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	return _ValidatorUtils.Contract.StakerInfo(&_ValidatorUtils.CallOpts, rollup, stakerAddress)
}

// StakerInfo is a free data retrieval call binding the contract method 0x8f67e6bb.
//
// Solidity: function stakerInfo(address rollup, address stakerAddress) view returns(bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge)
func (_ValidatorUtils *ValidatorUtilsCallerSession) StakerInfo(rollup common.Address, stakerAddress common.Address) (struct {
	IsStaked         bool
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
}, error) {
	return _ValidatorUtils.Contract.StakerInfo(&_ValidatorUtils.CallOpts, rollup, stakerAddress)
}

// TimedOutChallenges is a free data retrieval call binding the contract method 0xa8ac9cf3.
//
// Solidity: function timedOutChallenges(address rollup, uint256 startIndex, uint256 max) view returns(address[], bool hasMore)
func (_ValidatorUtils *ValidatorUtilsCaller) TimedOutChallenges(opts *bind.CallOpts, rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, bool, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "timedOutChallenges", rollup, startIndex, max)

	if err != nil {
		return *new([]common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// TimedOutChallenges is a free data retrieval call binding the contract method 0xa8ac9cf3.
//
// Solidity: function timedOutChallenges(address rollup, uint256 startIndex, uint256 max) view returns(address[], bool hasMore)
func (_ValidatorUtils *ValidatorUtilsSession) TimedOutChallenges(rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, bool, error) {
	return _ValidatorUtils.Contract.TimedOutChallenges(&_ValidatorUtils.CallOpts, rollup, startIndex, max)
}

// TimedOutChallenges is a free data retrieval call binding the contract method 0xa8ac9cf3.
//
// Solidity: function timedOutChallenges(address rollup, uint256 startIndex, uint256 max) view returns(address[], bool hasMore)
func (_ValidatorUtils *ValidatorUtilsCallerSession) TimedOutChallenges(rollup common.Address, startIndex *big.Int, max *big.Int) ([]common.Address, bool, error) {
	return _ValidatorUtils.Contract.TimedOutChallenges(&_ValidatorUtils.CallOpts, rollup, startIndex, max)
}
