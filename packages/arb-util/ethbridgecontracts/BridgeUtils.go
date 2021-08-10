// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// BridgeUtilsABI is the input ABI used to generate the binding from.
const BridgeUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"delayedBridge\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox\",\"type\":\"address\"}],\"name\":\"getCountsAndAccumulators\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"counts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BridgeUtilsBin is the compiled bytecode used for deploying new contracts.
var BridgeUtilsBin = "0x608060405234801561001057600080fd5b50610372806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806301ccad0214610030575b600080fd5b61005e6004803603604081101561004657600080fd5b506001600160a01b03813581169160200135166100c5565b6040518083600260200280838360005b8381101561008657818101518382015260200161006e565b5050505090500182600260200280838360005b838110156100b1578181015183820152602001610099565b505050509050019250505060405180910390f35b6100cd61031e565b6100d561031e565b6000846001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b15801561011057600080fd5b505afa158015610124573d6000803e3d6000fd5b505050506040513d602081101561013a57600080fd5b5051905080156101c0578083526040805163d9dd67ab60e01b81526000198301600482015290516001600160a01b0387169163d9dd67ab916024808301926020929190829003018186803b15801561019157600080fd5b505afa1580156101a5573d6000803e3d6000fd5b505050506040513d60208110156101bb57600080fd5b505182525b6000846001600160a01b031663d9b141ff6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101fb57600080fd5b505afa15801561020f573d6000803e3d6000fd5b505050506040513d602081101561022557600080fd5b50519050801561031557846001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b15801561026857600080fd5b505afa15801561027c573d6000803e3d6000fd5b505050506040513d602081101561029257600080fd5b50516020808601919091526040805163d9dd67ab60e01b81526000198401600482015290516001600160a01b0388169263d9dd67ab9260248082019391829003018186803b1580156102e357600080fd5b505afa1580156102f7573d6000803e3d6000fd5b505050506040513d602081101561030d57600080fd5b505160208401525b50509250929050565b6040518060400160405280600290602082028036833750919291505056fea2646970667358221220c5ffb5175bd411731a58af6859380d03c0d76b265c23705e7695a064d85c99b964736f6c634300060b0033"

// DeployBridgeUtils deploys a new Ethereum contract, binding an instance of BridgeUtils to it.
func DeployBridgeUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeUtils{BridgeUtilsCaller: BridgeUtilsCaller{contract: contract}, BridgeUtilsTransactor: BridgeUtilsTransactor{contract: contract}, BridgeUtilsFilterer: BridgeUtilsFilterer{contract: contract}}, nil
}

// BridgeUtils is an auto generated Go binding around an Ethereum contract.
type BridgeUtils struct {
	BridgeUtilsCaller     // Read-only binding to the contract
	BridgeUtilsTransactor // Write-only binding to the contract
	BridgeUtilsFilterer   // Log filterer for contract events
}

// BridgeUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeUtilsSession struct {
	Contract     *BridgeUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeUtilsCallerSession struct {
	Contract *BridgeUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeUtilsTransactorSession struct {
	Contract     *BridgeUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeUtilsRaw struct {
	Contract *BridgeUtils // Generic contract binding to access the raw methods on
}

// BridgeUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeUtilsCallerRaw struct {
	Contract *BridgeUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeUtilsTransactorRaw struct {
	Contract *BridgeUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeUtils creates a new instance of BridgeUtils, bound to a specific deployed contract.
func NewBridgeUtils(address common.Address, backend bind.ContractBackend) (*BridgeUtils, error) {
	contract, err := bindBridgeUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeUtils{BridgeUtilsCaller: BridgeUtilsCaller{contract: contract}, BridgeUtilsTransactor: BridgeUtilsTransactor{contract: contract}, BridgeUtilsFilterer: BridgeUtilsFilterer{contract: contract}}, nil
}

// NewBridgeUtilsCaller creates a new read-only instance of BridgeUtils, bound to a specific deployed contract.
func NewBridgeUtilsCaller(address common.Address, caller bind.ContractCaller) (*BridgeUtilsCaller, error) {
	contract, err := bindBridgeUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeUtilsCaller{contract: contract}, nil
}

// NewBridgeUtilsTransactor creates a new write-only instance of BridgeUtils, bound to a specific deployed contract.
func NewBridgeUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeUtilsTransactor, error) {
	contract, err := bindBridgeUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeUtilsTransactor{contract: contract}, nil
}

// NewBridgeUtilsFilterer creates a new log filterer instance of BridgeUtils, bound to a specific deployed contract.
func NewBridgeUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeUtilsFilterer, error) {
	contract, err := bindBridgeUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeUtilsFilterer{contract: contract}, nil
}

// bindBridgeUtils binds a generic wrapper to an already deployed contract.
func bindBridgeUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeUtils *BridgeUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeUtils.Contract.BridgeUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeUtils *BridgeUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeUtils.Contract.BridgeUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeUtils *BridgeUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeUtils.Contract.BridgeUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeUtils *BridgeUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeUtils *BridgeUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeUtils *BridgeUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeUtils.Contract.contract.Transact(opts, method, params...)
}

// GetCountsAndAccumulators is a free data retrieval call binding the contract method 0x01ccad02.
//
// Solidity: function getCountsAndAccumulators(address delayedBridge, address sequencerInbox) view returns(uint256[2] counts, bytes32[2] accs)
func (_BridgeUtils *BridgeUtilsCaller) GetCountsAndAccumulators(opts *bind.CallOpts, delayedBridge common.Address, sequencerInbox common.Address) (struct {
	Counts [2]*big.Int
	Accs   [2][32]byte
}, error) {
	var out []interface{}
	err := _BridgeUtils.contract.Call(opts, &out, "getCountsAndAccumulators", delayedBridge, sequencerInbox)

	outstruct := new(struct {
		Counts [2]*big.Int
		Accs   [2][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Counts = *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)
	outstruct.Accs = *abi.ConvertType(out[1], new([2][32]byte)).(*[2][32]byte)

	return *outstruct, err

}

// GetCountsAndAccumulators is a free data retrieval call binding the contract method 0x01ccad02.
//
// Solidity: function getCountsAndAccumulators(address delayedBridge, address sequencerInbox) view returns(uint256[2] counts, bytes32[2] accs)
func (_BridgeUtils *BridgeUtilsSession) GetCountsAndAccumulators(delayedBridge common.Address, sequencerInbox common.Address) (struct {
	Counts [2]*big.Int
	Accs   [2][32]byte
}, error) {
	return _BridgeUtils.Contract.GetCountsAndAccumulators(&_BridgeUtils.CallOpts, delayedBridge, sequencerInbox)
}

// GetCountsAndAccumulators is a free data retrieval call binding the contract method 0x01ccad02.
//
// Solidity: function getCountsAndAccumulators(address delayedBridge, address sequencerInbox) view returns(uint256[2] counts, bytes32[2] accs)
func (_BridgeUtils *BridgeUtilsCallerSession) GetCountsAndAccumulators(delayedBridge common.Address, sequencerInbox common.Address) (struct {
	Counts [2]*big.Int
	Accs   [2][32]byte
}, error) {
	return _BridgeUtils.Contract.GetCountsAndAccumulators(&_BridgeUtils.CallOpts, delayedBridge, sequencerInbox)
}
