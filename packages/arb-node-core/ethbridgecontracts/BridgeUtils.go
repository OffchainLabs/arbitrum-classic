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

// BridgeUtilsABI is the input ABI used to generate the binding from.
const BridgeUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractIBridge[2]\",\"name\":\"bridges\",\"type\":\"address[2]\"}],\"name\":\"getCountsAndAccumulators\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"counts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BridgeUtilsBin is the compiled bytecode used for deploying new contracts.
var BridgeUtilsBin = "0x608060405234801561001057600080fd5b50610257806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063732c6d7c14610030575b600080fd5b61004c6004803603604081101561004657600080fd5b506100b3565b6040518083600260200280838360005b8381101561007457818101518382015260200161005c565b5050505090500182600260200280838360005b8381101561009f578181015183820152602001610087565b505050509050019250505060405180910390f35b6100bb610203565b6100c3610203565b60005b60028110156101fd5760008482600281106100dd57fe5b60200201356001600160a01b031690506000816001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b15801561012857600080fd5b505afa15801561013c573d6000803e3d6000fd5b505050506040513d602081101561015257600080fd5b505190508085846002811061016357fe5b602002015280156101f357816001600160a01b031663d9dd67ab600183036040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156101b557600080fd5b505afa1580156101c9573d6000803e3d6000fd5b505050506040513d60208110156101df57600080fd5b50518484600281106101ed57fe5b60200201525b50506001016100c6565b50915091565b6040518060400160405280600290602082028036833750919291505056fea26469706673582212200dadb7f7dbf4fd73e5a75ec06afe0b0fbca664e5ed018c2101449ebe5a6bafc164736f6c634300060b0033"

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

// GetCountsAndAccumulators is a free data retrieval call binding the contract method 0x732c6d7c.
//
// Solidity: function getCountsAndAccumulators(address[2] bridges) view returns(uint256[2] counts, bytes32[2] accs)
func (_BridgeUtils *BridgeUtilsCaller) GetCountsAndAccumulators(opts *bind.CallOpts, bridges [2]common.Address) (struct {
	Counts [2]*big.Int
	Accs   [2][32]byte
}, error) {
	var out []interface{}
	err := _BridgeUtils.contract.Call(opts, &out, "getCountsAndAccumulators", bridges)

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

// GetCountsAndAccumulators is a free data retrieval call binding the contract method 0x732c6d7c.
//
// Solidity: function getCountsAndAccumulators(address[2] bridges) view returns(uint256[2] counts, bytes32[2] accs)
func (_BridgeUtils *BridgeUtilsSession) GetCountsAndAccumulators(bridges [2]common.Address) (struct {
	Counts [2]*big.Int
	Accs   [2][32]byte
}, error) {
	return _BridgeUtils.Contract.GetCountsAndAccumulators(&_BridgeUtils.CallOpts, bridges)
}

// GetCountsAndAccumulators is a free data retrieval call binding the contract method 0x732c6d7c.
//
// Solidity: function getCountsAndAccumulators(address[2] bridges) view returns(uint256[2] counts, bytes32[2] accs)
func (_BridgeUtils *BridgeUtilsCallerSession) GetCountsAndAccumulators(bridges [2]common.Address) (struct {
	Counts [2]*big.Int
	Accs   [2][32]byte
}, error) {
	return _BridgeUtils.Contract.GetCountsAndAccumulators(&_BridgeUtils.CallOpts, bridges)
}
