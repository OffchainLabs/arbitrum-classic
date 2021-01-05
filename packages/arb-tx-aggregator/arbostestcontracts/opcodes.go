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
const OpCodesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"conn\",\"type\":\"address\"}],\"name\":\"getNestedOrigin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"conn\",\"type\":\"address\"}],\"name\":\"getNestedSend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrigin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OpCodesFuncSigs maps the 4-byte function signature to its string representation.
var OpCodesFuncSigs = map[string]string{
	"9663f88f": "getBlockHash()",
	"03dd3df4": "getNestedOrigin(address)",
	"84e58c86": "getNestedSend(address)",
	"df1f29ee": "getOrigin()",
	"5e01eb5a": "getSender()",
}

// OpCodesBin is the compiled bytecode used for deploying new contracts.
var OpCodesBin = "0x608060405234801561001057600080fd5b506101e0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806303dd3df41461005c5780635e01eb5a1461009e57806384e58c86146100a65780639663f88f146100cc578063df1f29ee146100e6575b600080fd5b6100826004803603602081101561007257600080fd5b50356001600160a01b03166100ee565b604080516001600160a01b039092168252519081900360200190f35b61008261015d565b610082600480360360208110156100bc57600080fd5b50356001600160a01b0316610161565b6100d461019e565b60408051918252519081900360200190f35b6100826101a7565b6000816001600160a01b031663df1f29ee6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561012b57600080fd5b505af115801561013f573d6000803e3d6000fd5b505050506040513d602081101561015557600080fd5b505192915050565b3390565b6000816001600160a01b0316635e01eb5a6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561012b57600080fd5b60001943014090565b329056fea265627a7a723158201a270483605a2f27a05812c416c647d1a8ea35cb5dccaf2a12f4443d6b7d5c6364736f6c63430005110032"

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
