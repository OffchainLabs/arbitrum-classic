// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbostestcontracts

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

// OpCodesABI is the input ABI used to generate the binding from.
const OpCodesABI = "[{\"inputs\":[],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conn\",\"type\":\"address\"}],\"name\":\"getNestedOrigin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conn\",\"type\":\"address\"}],\"name\":\"getNestedSend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrigin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OpCodesBin is the compiled bytecode used for deploying new contracts.
var OpCodesBin = "0x608060405234801561001057600080fd5b506101fc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806303dd3df41461006757806351be4eaa146100a95780635e01eb5a146100c357806384e58c86146100cb5780639663f88f146100f1578063df1f29ee146100f9575b600080fd5b61008d6004803603602081101561007d57600080fd5b50356001600160a01b0316610101565b604080516001600160a01b039092168252519081900360200190f35b6100b1610170565b60408051918252519081900360200190f35b61008d610178565b61008d600480360360208110156100e157600080fd5b50356001600160a01b031661017c565b6100b16101b9565b61008d6101c2565b6000816001600160a01b031663df1f29ee6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561013e57600080fd5b505af1158015610152573d6000803e3d6000fd5b505050506040513d602081101561016857600080fd5b505192915050565b60005a905090565b3390565b6000816001600160a01b0316635e01eb5a6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561013e57600080fd5b60001943014090565b329056fea26469706673582212204166b6173dacc8365955ede6f319adebd7ea285941b302c0273b199fc378aaa464736f6c634300060c0033"

// DeployOpCodes deploys a new Ethereum contract, binding an instance of OpCodes to it.
func DeployOpCodes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OpCodes, error) {
	parsed, err := abi.JSON(strings.NewReader(OpCodesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OpCodesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OpCodes{OpCodesCaller: OpCodesCaller{contract: contract}, OpCodesTransactor: OpCodesTransactor{contract: contract}, OpCodesFilterer: OpCodesFilterer{contract: contract}}, nil
}

// OpCodes is an auto generated Go binding around an Ethereum contract.
type OpCodes struct {
	OpCodesCaller     // Read-only binding to the contract
	OpCodesTransactor // Write-only binding to the contract
	OpCodesFilterer   // Log filterer for contract events
}

// OpCodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpCodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpCodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpCodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpCodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpCodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpCodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpCodesSession struct {
	Contract     *OpCodes          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpCodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpCodesCallerSession struct {
	Contract *OpCodesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OpCodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpCodesTransactorSession struct {
	Contract     *OpCodesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OpCodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpCodesRaw struct {
	Contract *OpCodes // Generic contract binding to access the raw methods on
}

// OpCodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpCodesCallerRaw struct {
	Contract *OpCodesCaller // Generic read-only contract binding to access the raw methods on
}

// OpCodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpCodesTransactorRaw struct {
	Contract *OpCodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpCodes creates a new instance of OpCodes, bound to a specific deployed contract.
func NewOpCodes(address common.Address, backend bind.ContractBackend) (*OpCodes, error) {
	contract, err := bindOpCodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpCodes{OpCodesCaller: OpCodesCaller{contract: contract}, OpCodesTransactor: OpCodesTransactor{contract: contract}, OpCodesFilterer: OpCodesFilterer{contract: contract}}, nil
}

// NewOpCodesCaller creates a new read-only instance of OpCodes, bound to a specific deployed contract.
func NewOpCodesCaller(address common.Address, caller bind.ContractCaller) (*OpCodesCaller, error) {
	contract, err := bindOpCodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpCodesCaller{contract: contract}, nil
}

// NewOpCodesTransactor creates a new write-only instance of OpCodes, bound to a specific deployed contract.
func NewOpCodesTransactor(address common.Address, transactor bind.ContractTransactor) (*OpCodesTransactor, error) {
	contract, err := bindOpCodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpCodesTransactor{contract: contract}, nil
}

// NewOpCodesFilterer creates a new log filterer instance of OpCodes, bound to a specific deployed contract.
func NewOpCodesFilterer(address common.Address, filterer bind.ContractFilterer) (*OpCodesFilterer, error) {
	contract, err := bindOpCodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpCodesFilterer{contract: contract}, nil
}

// bindOpCodes binds a generic wrapper to an already deployed contract.
func bindOpCodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpCodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpCodes *OpCodesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpCodes.Contract.OpCodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpCodes *OpCodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpCodes.Contract.OpCodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpCodes *OpCodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpCodes.Contract.OpCodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpCodes *OpCodesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpCodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpCodes *OpCodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpCodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpCodes *OpCodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpCodes.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a paid mutator transaction binding the contract method 0x9663f88f.
//
// Solidity: function getBlockHash() returns(bytes32)
func (_OpCodes *OpCodesTransactor) GetBlockHash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpCodes.contract.Transact(opts, "getBlockHash")
}

// GetBlockHash is a paid mutator transaction binding the contract method 0x9663f88f.
//
// Solidity: function getBlockHash() returns(bytes32)
func (_OpCodes *OpCodesSession) GetBlockHash() (*types.Transaction, error) {
	return _OpCodes.Contract.GetBlockHash(&_OpCodes.TransactOpts)
}

// GetBlockHash is a paid mutator transaction binding the contract method 0x9663f88f.
//
// Solidity: function getBlockHash() returns(bytes32)
func (_OpCodes *OpCodesTransactorSession) GetBlockHash() (*types.Transaction, error) {
	return _OpCodes.Contract.GetBlockHash(&_OpCodes.TransactOpts)
}

// GetGasLeft is a paid mutator transaction binding the contract method 0x51be4eaa.
//
// Solidity: function getGasLeft() returns(uint256)
func (_OpCodes *OpCodesTransactor) GetGasLeft(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpCodes.contract.Transact(opts, "getGasLeft")
}

// GetGasLeft is a paid mutator transaction binding the contract method 0x51be4eaa.
//
// Solidity: function getGasLeft() returns(uint256)
func (_OpCodes *OpCodesSession) GetGasLeft() (*types.Transaction, error) {
	return _OpCodes.Contract.GetGasLeft(&_OpCodes.TransactOpts)
}

// GetGasLeft is a paid mutator transaction binding the contract method 0x51be4eaa.
//
// Solidity: function getGasLeft() returns(uint256)
func (_OpCodes *OpCodesTransactorSession) GetGasLeft() (*types.Transaction, error) {
	return _OpCodes.Contract.GetGasLeft(&_OpCodes.TransactOpts)
}

// GetNestedOrigin is a paid mutator transaction binding the contract method 0x03dd3df4.
//
// Solidity: function getNestedOrigin(address conn) returns(address)
func (_OpCodes *OpCodesTransactor) GetNestedOrigin(opts *bind.TransactOpts, conn common.Address) (*types.Transaction, error) {
	return _OpCodes.contract.Transact(opts, "getNestedOrigin", conn)
}

// GetNestedOrigin is a paid mutator transaction binding the contract method 0x03dd3df4.
//
// Solidity: function getNestedOrigin(address conn) returns(address)
func (_OpCodes *OpCodesSession) GetNestedOrigin(conn common.Address) (*types.Transaction, error) {
	return _OpCodes.Contract.GetNestedOrigin(&_OpCodes.TransactOpts, conn)
}

// GetNestedOrigin is a paid mutator transaction binding the contract method 0x03dd3df4.
//
// Solidity: function getNestedOrigin(address conn) returns(address)
func (_OpCodes *OpCodesTransactorSession) GetNestedOrigin(conn common.Address) (*types.Transaction, error) {
	return _OpCodes.Contract.GetNestedOrigin(&_OpCodes.TransactOpts, conn)
}

// GetNestedSend is a paid mutator transaction binding the contract method 0x84e58c86.
//
// Solidity: function getNestedSend(address conn) returns(address)
func (_OpCodes *OpCodesTransactor) GetNestedSend(opts *bind.TransactOpts, conn common.Address) (*types.Transaction, error) {
	return _OpCodes.contract.Transact(opts, "getNestedSend", conn)
}

// GetNestedSend is a paid mutator transaction binding the contract method 0x84e58c86.
//
// Solidity: function getNestedSend(address conn) returns(address)
func (_OpCodes *OpCodesSession) GetNestedSend(conn common.Address) (*types.Transaction, error) {
	return _OpCodes.Contract.GetNestedSend(&_OpCodes.TransactOpts, conn)
}

// GetNestedSend is a paid mutator transaction binding the contract method 0x84e58c86.
//
// Solidity: function getNestedSend(address conn) returns(address)
func (_OpCodes *OpCodesTransactorSession) GetNestedSend(conn common.Address) (*types.Transaction, error) {
	return _OpCodes.Contract.GetNestedSend(&_OpCodes.TransactOpts, conn)
}

// GetOrigin is a paid mutator transaction binding the contract method 0xdf1f29ee.
//
// Solidity: function getOrigin() returns(address)
func (_OpCodes *OpCodesTransactor) GetOrigin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpCodes.contract.Transact(opts, "getOrigin")
}

// GetOrigin is a paid mutator transaction binding the contract method 0xdf1f29ee.
//
// Solidity: function getOrigin() returns(address)
func (_OpCodes *OpCodesSession) GetOrigin() (*types.Transaction, error) {
	return _OpCodes.Contract.GetOrigin(&_OpCodes.TransactOpts)
}

// GetOrigin is a paid mutator transaction binding the contract method 0xdf1f29ee.
//
// Solidity: function getOrigin() returns(address)
func (_OpCodes *OpCodesTransactorSession) GetOrigin() (*types.Transaction, error) {
	return _OpCodes.Contract.GetOrigin(&_OpCodes.TransactOpts)
}

// GetSender is a paid mutator transaction binding the contract method 0x5e01eb5a.
//
// Solidity: function getSender() returns(address)
func (_OpCodes *OpCodesTransactor) GetSender(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpCodes.contract.Transact(opts, "getSender")
}

// GetSender is a paid mutator transaction binding the contract method 0x5e01eb5a.
//
// Solidity: function getSender() returns(address)
func (_OpCodes *OpCodesSession) GetSender() (*types.Transaction, error) {
	return _OpCodes.Contract.GetSender(&_OpCodes.TransactOpts)
}

// GetSender is a paid mutator transaction binding the contract method 0x5e01eb5a.
//
// Solidity: function getSender() returns(address)
func (_OpCodes *OpCodesTransactorSession) GetSender() (*types.Transaction, error) {
	return _OpCodes.Contract.GetSender(&_OpCodes.TransactOpts)
}
