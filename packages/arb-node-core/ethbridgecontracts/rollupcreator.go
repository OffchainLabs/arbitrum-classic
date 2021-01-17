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

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"deliverMessageToInbox\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inboxInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6108a28061007d6000396000f3fe60806040526004361061008c5760003560e01c80630c301fb214610091578063413b35bd146100bf578063636f807d14610106578063715018a6146101255780638da5cb5b1461013a5780639e5d4c491461016b578063ab5d89431461027c578063c29372de14610291578063cee3d728146102c4578063e45b7ce6146102ff578063f2fde38b1461033a575b600080fd5b34801561009d57600080fd5b506100a661036d565b6040805192835260208301919091528051918290030190f35b3480156100cb57600080fd5b506100f2600480360360208110156100e257600080fd5b50356001600160a01b0316610377565b604080519115158252519081900360200190f35b6101236004803603602081101561011c57600080fd5b503561038c565b005b34801561013157600080fd5b50610123610437565b34801561014657600080fd5b5061014f6104c7565b604080516001600160a01b039092168252519081900360200190f35b34801561017757600080fd5b506101fb6004803603606081101561018e57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156101bd57600080fd5b8201836020820111156101cf57600080fd5b803590602001918460018302840111600160201b831117156101f057600080fd5b5090925090506104d6565b60405180831515815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610240578181015183820152602001610228565b50505050905090810190601f16801561026d5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561028857600080fd5b5061014f6105c6565b34801561029d57600080fd5b506100f2600480360360208110156102b457600080fd5b50356001600160a01b03166105d5565b3480156102d057600080fd5b50610123600480360360408110156102e757600080fd5b506001600160a01b03813516906020013515156105ea565b34801561030b57600080fd5b506101236004803603604081101561032257600080fd5b506001600160a01b038135169060200135151561066d565b34801561034657600080fd5b506101236004803603602081101561035d57600080fd5b50356001600160a01b03166106f0565b6005546004549091565b60026020526000908152604090205460ff1681565b3360009081526001602052604090205460ff166103e1576040805162461bcd60e51b815260206004820152600e60248201526d09c9ea8be8ca49e9abe929c849eb60931b604482015290519081900360640190fd5b6005546004546103f181846107d6565b60045560018201600555604080513381529051829184917fb9c63117c789d95c5dca3f3c80aceb20aeaf5e413cc5e5831f77b62e5489e9fe9181900360200190a3505050565b61043f610802565b6000546001600160a01b0390811691161461048f576040805162461bcd60e51b8152602060048201819052602482015260008051602061082d833981519152604482015290519081900360640190fd5b600080546040516001600160a01b039091169060008051602061084d833981519152908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b3360009081526002602052604081205460609060ff1661052f576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b600380546001600160a01b031916331790556040516001600160a01b03871690869086908690808383808284376040519201945060009350909150508083038185875af1925050503d80600081146105a3576040519150601f19603f3d011682016040523d82523d6000602084013e6105a8565b606091505b50600380546001600160a01b03191690559097909650945050505050565b6003546001600160a01b031681565b60016020526000908152604090205460ff1681565b6105f2610802565b6000546001600160a01b03908116911614610642576040805162461bcd60e51b8152602060048201819052602482015260008051602061082d833981519152604482015290519081900360640190fd5b6001600160a01b03919091166000908152600260205260409020805460ff1916911515919091179055565b610675610802565b6000546001600160a01b039081169116146106c5576040805162461bcd60e51b8152602060048201819052602482015260008051602061082d833981519152604482015290519081900360640190fd5b6001600160a01b03919091166000908152600160205260409020805460ff1916911515919091179055565b6106f8610802565b6000546001600160a01b03908116911614610748576040805162461bcd60e51b8152602060048201819052602482015260008051602061082d833981519152604482015290519081900360640190fd5b6001600160a01b03811661078d5760405162461bcd60e51b81526004018080602001828103825260268152602001806108076026913960400191505060405180910390fd5b600080546040516001600160a01b038085169392169160008051602061084d83398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220c38005c18303dc743881af9a2bdd19796d7cd7378536d0b69ee06cdbb8a3ca3264736f6c634300060c0033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeSession) ActiveOutbox() (common.Address, error) {
	return _Bridge.Contract.ActiveOutbox(&_Bridge.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeCallerSession) ActiveOutbox() (common.Address, error) {
	return _Bridge.Contract.ActiveOutbox(&_Bridge.CallOpts)
}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address ) view returns(bool)
func (_Bridge *BridgeCaller) AllowedInboxes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedInboxes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address ) view returns(bool)
func (_Bridge *BridgeSession) AllowedInboxes(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.AllowedInboxes(&_Bridge.CallOpts, arg0)
}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) AllowedInboxes(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.AllowedInboxes(&_Bridge.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) view returns(bool)
func (_Bridge *BridgeCaller) AllowedOutboxes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedOutboxes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) view returns(bool)
func (_Bridge *BridgeSession) AllowedOutboxes(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.AllowedOutboxes(&_Bridge.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) AllowedOutboxes(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.AllowedOutboxes(&_Bridge.CallOpts, arg0)
}

// InboxInfo is a free data retrieval call binding the contract method 0x0c301fb2.
//
// Solidity: function inboxInfo() view returns(uint256, bytes32)
func (_Bridge *BridgeCaller) InboxInfo(opts *bind.CallOpts) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "inboxInfo")

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// InboxInfo is a free data retrieval call binding the contract method 0x0c301fb2.
//
// Solidity: function inboxInfo() view returns(uint256, bytes32)
func (_Bridge *BridgeSession) InboxInfo() (*big.Int, [32]byte, error) {
	return _Bridge.Contract.InboxInfo(&_Bridge.CallOpts)
}

// InboxInfo is a free data retrieval call binding the contract method 0x0c301fb2.
//
// Solidity: function inboxInfo() view returns(uint256, bytes32)
func (_Bridge *BridgeCallerSession) InboxInfo() (*big.Int, [32]byte, error) {
	return _Bridge.Contract.InboxInfo(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x636f807d.
//
// Solidity: function deliverMessageToInbox(bytes32 messageHash) payable returns()
func (_Bridge *BridgeTransactor) DeliverMessageToInbox(opts *bind.TransactOpts, messageHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deliverMessageToInbox", messageHash)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x636f807d.
//
// Solidity: function deliverMessageToInbox(bytes32 messageHash) payable returns()
func (_Bridge *BridgeSession) DeliverMessageToInbox(messageHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DeliverMessageToInbox(&_Bridge.TransactOpts, messageHash)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x636f807d.
//
// Solidity: function deliverMessageToInbox(bytes32 messageHash) payable returns()
func (_Bridge *BridgeTransactorSession) DeliverMessageToInbox(messageHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DeliverMessageToInbox(&_Bridge.TransactOpts, messageHash)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeTransactor) ExecuteCall(opts *bind.TransactOpts, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeCall", destAddr, amount, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeSession) ExecuteCall(destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteCall(&_Bridge.TransactOpts, destAddr, amount, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeTransactorSession) ExecuteCall(destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteCall(&_Bridge.TransactOpts, destAddr, amount, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactor) SetInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setInbox", inbox, enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeSession) SetInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetInbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactorSession) SetInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetInbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactor) SetOutbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setOutbox", inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeSession) SetOutbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOutbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactorSession) SetOutbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOutbox(&_Bridge.TransactOpts, inbox, enabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// BridgeMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the Bridge contract.
type BridgeMessageDeliveredIterator struct {
	Event *BridgeMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMessageDelivered)
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
		it.Event = new(BridgeMessageDelivered)
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
func (it *BridgeMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMessageDelivered represents a MessageDelivered event raised by the Bridge contract.
type BridgeMessageDelivered struct {
	MessageIndex   *big.Int
	BeforeInboxAcc [32]byte
	Inbox          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0xb9c63117c789d95c5dca3f3c80aceb20aeaf5e413cc5e5831f77b62e5489e9fe.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox)
func (_Bridge *BridgeFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeMessageDeliveredIterator{contract: _Bridge.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0xb9c63117c789d95c5dca3f3c80aceb20aeaf5e413cc5e5831f77b62e5489e9fe.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox)
func (_Bridge *BridgeFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMessageDelivered)
				if err := _Bridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0xb9c63117c789d95c5dca3f3c80aceb20aeaf5e413cc5e5831f77b62e5489e9fe.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox)
func (_Bridge *BridgeFilterer) ParseMessageDelivered(log types.Log) (*BridgeMessageDelivered, error) {
	event := new(BridgeMessageDelivered)
	if err := _Bridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInboxABI is the input ABI used to generate the binding from.
const IInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"BuddyContractPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IInbox is an auto generated Go binding around an Ethereum contract.
type IInbox struct {
	IInboxCaller     // Read-only binding to the contract
	IInboxTransactor // Write-only binding to the contract
	IInboxFilterer   // Log filterer for contract events
}

// IInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInboxSession struct {
	Contract     *IInbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInboxCallerSession struct {
	Contract *IInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInboxTransactorSession struct {
	Contract     *IInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInboxRaw struct {
	Contract *IInbox // Generic contract binding to access the raw methods on
}

// IInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInboxCallerRaw struct {
	Contract *IInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInboxTransactorRaw struct {
	Contract *IInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInbox creates a new instance of IInbox, bound to a specific deployed contract.
func NewIInbox(address common.Address, backend bind.ContractBackend) (*IInbox, error) {
	contract, err := bindIInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInbox{IInboxCaller: IInboxCaller{contract: contract}, IInboxTransactor: IInboxTransactor{contract: contract}, IInboxFilterer: IInboxFilterer{contract: contract}}, nil
}

// NewIInboxCaller creates a new read-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxCaller(address common.Address, caller bind.ContractCaller) (*IInboxCaller, error) {
	contract, err := bindIInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxCaller{contract: contract}, nil
}

// NewIInboxTransactor creates a new write-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IInboxTransactor, error) {
	contract, err := bindIInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxTransactor{contract: contract}, nil
}

// NewIInboxFilterer creates a new log filterer instance of IInbox, bound to a specific deployed contract.
func NewIInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IInboxFilterer, error) {
	contract, err := bindIInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInboxFilterer{contract: contract}, nil
}

// bindIInbox binds a generic wrapper to an already deployed contract.
func bindIInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.IInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transact(opts, method, params...)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_IInbox *IInboxTransactor) DeployL2ContractPair(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "deployL2ContractPair", maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_IInbox *IInboxSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.DeployL2ContractPair(&_IInbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_IInbox *IInboxTransactorSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.DeployL2ContractPair(&_IInbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_IInbox *IInboxTransactor) DepositEthMessage(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "depositEthMessage", to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_IInbox *IInboxSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEthMessage(&_IInbox.TransactOpts, to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_IInbox *IInboxTransactorSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEthMessage(&_IInbox.TransactOpts, to)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_IInbox *IInboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_IInbox *IInboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL2Message(&_IInbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_IInbox *IInboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL2Message(&_IInbox.TransactOpts, messageData)
}

// IInboxBuddyContractPairIterator is returned from FilterBuddyContractPair and is used to iterate over the raw logs and unpacked data for BuddyContractPair events raised by the IInbox contract.
type IInboxBuddyContractPairIterator struct {
	Event *IInboxBuddyContractPair // Event containing the contract specifics and raw log

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
func (it *IInboxBuddyContractPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxBuddyContractPair)
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
		it.Event = new(IInboxBuddyContractPair)
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
func (it *IInboxBuddyContractPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxBuddyContractPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxBuddyContractPair represents a BuddyContractPair event raised by the IInbox contract.
type IInboxBuddyContractPair struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuddyContractPair is a free log retrieval operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_IInbox *IInboxFilterer) FilterBuddyContractPair(opts *bind.FilterOpts, sender []common.Address) (*IInboxBuddyContractPairIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return &IInboxBuddyContractPairIterator{contract: _IInbox.contract, event: "BuddyContractPair", logs: logs, sub: sub}, nil
}

// WatchBuddyContractPair is a free log subscription operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_IInbox *IInboxFilterer) WatchBuddyContractPair(opts *bind.WatchOpts, sink chan<- *IInboxBuddyContractPair, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxBuddyContractPair)
				if err := _IInbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
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

// ParseBuddyContractPair is a log parse operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_IInbox *IInboxFilterer) ParseBuddyContractPair(log types.Log) (*IInboxBuddyContractPair, error) {
	event := new(IInboxBuddyContractPair)
	if err := _IInbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the IInbox contract.
type IInboxInboxMessageDeliveredIterator struct {
	Event *IInboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IInboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxInboxMessageDelivered)
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
		it.Event = new(IInboxInboxMessageDelivered)
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
func (it *IInboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the IInbox contract.
type IInboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Kind       uint8
	Sender     common.Address
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xe5855b2ddc828cd927ab9e242f891de16135ed173f3ae45ad42179be197f8c8b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, uint8 kind, address sender, bytes data)
func (_IInbox *IInboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*IInboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IInboxInboxMessageDeliveredIterator{contract: _IInbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xe5855b2ddc828cd927ab9e242f891de16135ed173f3ae45ad42179be197f8c8b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, uint8 kind, address sender, bytes data)
func (_IInbox *IInboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *IInboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxInboxMessageDelivered)
				if err := _IInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xe5855b2ddc828cd927ab9e242f891de16135ed173f3ae45ad42179be197f8c8b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, uint8 kind, address sender, bytes data)
func (_IInbox *IInboxFilterer) ParseInboxMessageDelivered(log types.Log) (*IInboxInboxMessageDelivered, error) {
	event := new(IInboxInboxMessageDelivered)
	if err := _IInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the IInbox contract.
type IInboxInboxMessageDeliveredFromOriginIterator struct {
	Event *IInboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(IInboxInboxMessageDeliveredFromOrigin)
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
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the IInbox contract.
type IInboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Kind       uint8
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x491fab296ae66721d1d8dfc5f3457bec8c5f2586ce6b355adf0feeb4a4a24bdb.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum, uint8 kind, address sender)
func (_IInbox *IInboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*IInboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IInboxInboxMessageDeliveredFromOriginIterator{contract: _IInbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0x491fab296ae66721d1d8dfc5f3457bec8c5f2586ce6b355adf0feeb4a4a24bdb.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum, uint8 kind, address sender)
func (_IInbox *IInboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *IInboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxInboxMessageDeliveredFromOrigin)
				if err := _IInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0x491fab296ae66721d1d8dfc5f3457bec8c5f2586ce6b355adf0feeb4a4a24bdb.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum, uint8 kind, address sender)
func (_IInbox *IInboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*IInboxInboxMessageDeliveredFromOrigin, error) {
	event := new(IInboxInboxMessageDeliveredFromOrigin)
	if err := _IInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOutboxEntryABI is the input ABI used to generate the binding from.
const IOutboxEntryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"calcRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"}],\"name\":\"spendOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IOutboxEntry is an auto generated Go binding around an Ethereum contract.
type IOutboxEntry struct {
	IOutboxEntryCaller     // Read-only binding to the contract
	IOutboxEntryTransactor // Write-only binding to the contract
	IOutboxEntryFilterer   // Log filterer for contract events
}

// IOutboxEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOutboxEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOutboxEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOutboxEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOutboxEntrySession struct {
	Contract     *IOutboxEntry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOutboxEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOutboxEntryCallerSession struct {
	Contract *IOutboxEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IOutboxEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOutboxEntryTransactorSession struct {
	Contract     *IOutboxEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IOutboxEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOutboxEntryRaw struct {
	Contract *IOutboxEntry // Generic contract binding to access the raw methods on
}

// IOutboxEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOutboxEntryCallerRaw struct {
	Contract *IOutboxEntryCaller // Generic read-only contract binding to access the raw methods on
}

// IOutboxEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOutboxEntryTransactorRaw struct {
	Contract *IOutboxEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOutboxEntry creates a new instance of IOutboxEntry, bound to a specific deployed contract.
func NewIOutboxEntry(address common.Address, backend bind.ContractBackend) (*IOutboxEntry, error) {
	contract, err := bindIOutboxEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOutboxEntry{IOutboxEntryCaller: IOutboxEntryCaller{contract: contract}, IOutboxEntryTransactor: IOutboxEntryTransactor{contract: contract}, IOutboxEntryFilterer: IOutboxEntryFilterer{contract: contract}}, nil
}

// NewIOutboxEntryCaller creates a new read-only instance of IOutboxEntry, bound to a specific deployed contract.
func NewIOutboxEntryCaller(address common.Address, caller bind.ContractCaller) (*IOutboxEntryCaller, error) {
	contract, err := bindIOutboxEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOutboxEntryCaller{contract: contract}, nil
}

// NewIOutboxEntryTransactor creates a new write-only instance of IOutboxEntry, bound to a specific deployed contract.
func NewIOutboxEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*IOutboxEntryTransactor, error) {
	contract, err := bindIOutboxEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOutboxEntryTransactor{contract: contract}, nil
}

// NewIOutboxEntryFilterer creates a new log filterer instance of IOutboxEntry, bound to a specific deployed contract.
func NewIOutboxEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*IOutboxEntryFilterer, error) {
	contract, err := bindIOutboxEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOutboxEntryFilterer{contract: contract}, nil
}

// bindIOutboxEntry binds a generic wrapper to an already deployed contract.
func bindIOutboxEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOutboxEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOutboxEntry *IOutboxEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOutboxEntry.Contract.IOutboxEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOutboxEntry *IOutboxEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOutboxEntry.Contract.IOutboxEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOutboxEntry *IOutboxEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOutboxEntry.Contract.IOutboxEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOutboxEntry *IOutboxEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOutboxEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOutboxEntry *IOutboxEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOutboxEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOutboxEntry *IOutboxEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOutboxEntry.Contract.contract.Transact(opts, method, params...)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_IOutboxEntry *IOutboxEntryCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IOutboxEntry.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_IOutboxEntry *IOutboxEntrySession) IsMaster() (bool, error) {
	return _IOutboxEntry.Contract.IsMaster(&_IOutboxEntry.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_IOutboxEntry *IOutboxEntryCallerSession) IsMaster() (bool, error) {
	return _IOutboxEntry.Contract.IsMaster(&_IOutboxEntry.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(address owner, bytes32 root) returns()
func (_IOutboxEntry *IOutboxEntryTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, root [32]byte) (*types.Transaction, error) {
	return _IOutboxEntry.contract.Transact(opts, "initialize", owner, root)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(address owner, bytes32 root) returns()
func (_IOutboxEntry *IOutboxEntrySession) Initialize(owner common.Address, root [32]byte) (*types.Transaction, error) {
	return _IOutboxEntry.Contract.Initialize(&_IOutboxEntry.TransactOpts, owner, root)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(address owner, bytes32 root) returns()
func (_IOutboxEntry *IOutboxEntryTransactorSession) Initialize(owner common.Address, root [32]byte) (*types.Transaction, error) {
	return _IOutboxEntry.Contract.Initialize(&_IOutboxEntry.TransactOpts, owner, root)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 path) returns()
func (_IOutboxEntry *IOutboxEntryTransactor) SpendOutput(opts *bind.TransactOpts, calcRoot [32]byte, path *big.Int) (*types.Transaction, error) {
	return _IOutboxEntry.contract.Transact(opts, "spendOutput", calcRoot, path)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 path) returns()
func (_IOutboxEntry *IOutboxEntrySession) SpendOutput(calcRoot [32]byte, path *big.Int) (*types.Transaction, error) {
	return _IOutboxEntry.Contract.SpendOutput(&_IOutboxEntry.TransactOpts, calcRoot, path)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 calcRoot, uint256 path) returns()
func (_IOutboxEntry *IOutboxEntryTransactorSession) SpendOutput(calcRoot [32]byte, path *big.Int) (*types.Transaction, error) {
	return _IOutboxEntry.Contract.SpendOutput(&_IOutboxEntry.TransactOpts, calcRoot, path)
}

// InboxABI is the input ABI used to generate the binding from.
const InboxABI = "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"BuddyContractPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InboxBin is the compiled bytecode used for deploying new contracts.
var InboxBin = "0x608060405234801561001057600080fd5b506040516107683803806107688339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610703806100656000396000f3fe60806040526004361061003f5760003560e01c80631fe927cf146100445780636f5dfdca146100c1578063afcc220b1461014f578063b75436bb14610175575b600080fd5b34801561005057600080fd5b506100bf6004803603602081101561006757600080fd5b810190602081018135600160201b81111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460018302840111600160201b831117156100b457600080fd5b5090925090506101f0565b005b3480156100cd57600080fd5b506100bf600480360360808110156100e457600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561011157600080fd5b82018360208201111561012357600080fd5b803590602001918460018302840111600160201b8311171561014457600080fd5b50909250905061031f565b6100bf6004803603602081101561016557600080fd5b50356001600160a01b03166103f4565b34801561018157600080fd5b506100bf6004803603602081101561019857600080fd5b810190602081018135600160201b8111156101b257600080fd5b8201836020820111156101c457600080fd5b803590602001918460018302840111600160201b831117156101e557600080fd5b50909250905061042e565b333214610232576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b60008054604080516306180fd960e11b815281516001600160a01b0390931692630c301fb292600480840193919291829003018186803b15801561027557600080fd5b505afa158015610289573d6000803e3d6000fd5b505050506040513d604081101561029f57600080fd5b50516040519091506102e0906102db9060039033904390429087908a908a90808383808284376040519201829003909120935061047492505050565b6104dd565b6040805160038152336020820152815183927f491fab296ae66721d1d8dfc5f3457bec8c5f2586ce6b355adf0feeb4a4a24bdb928290030190a2505050565b61032833610545565b610376576040805162461bcd60e51b815260206004820152601a6024820152791b5d5cdd0818994818d85b1b195908189e4818dbdb9d1c9858dd60321b604482015290519081900360640190fd5b6103c26005338787878787604051602001808681526020018581526020018481526020018383808284378083019250505095505050505050604051602081830303815290604052610581565b60405133907f49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f90600090a25050505050565b604080516001600160a01b03831660208201523481830152815180820383018152606090910190915261042b906000903390610581565b50565b61047060033384848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061058192505050565b5050565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b600080546040805163636f807d60e01b81526004810185905290516001600160a01b039092169263636f807d923492602480820193929182900301818588803b15801561052957600080fd5b505af115801561053d573d6000803e3d6000fd5b505050505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061057957508115155b949350505050565b60008054604080516306180fd960e11b815281516001600160a01b0390931692630c301fb292600480840193919291829003018186803b1580156105c457600080fd5b505afa1580156105d8573d6000803e3d6000fd5b505050506040513d60408110156105ee57600080fd5b505182516020840120909150610611906102db9086908690439042908790610474565b807fe5855b2ddc828cd927ab9e242f891de16135ed173f3ae45ad42179be197f8c8b858585604051808460ff168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561068b578181015183820152602001610673565b50505050905090810190601f1680156106b85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a25050505056fea264697066735822122031a99336f5cced8178124a34b3103c1ec0d12e4e541bc04b2f85ac3ab1e4986664736f6c634300060c0033"

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxTransactor) DeployL2ContractPair(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "deployL2ContractPair", maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.DeployL2ContractPair(&_Inbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxTransactorSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.DeployL2ContractPair(&_Inbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxTransactor) DepositEthMessage(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEthMessage", to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEthMessage(&_Inbox.TransactOpts, to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0xafcc220b.
//
// Solidity: function depositEthMessage(address to) payable returns()
func (_Inbox *InboxTransactorSession) DepositEthMessage(to common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEthMessage(&_Inbox.TransactOpts, to)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// InboxBuddyContractPairIterator is returned from FilterBuddyContractPair and is used to iterate over the raw logs and unpacked data for BuddyContractPair events raised by the Inbox contract.
type InboxBuddyContractPairIterator struct {
	Event *InboxBuddyContractPair // Event containing the contract specifics and raw log

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
func (it *InboxBuddyContractPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxBuddyContractPair)
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
		it.Event = new(InboxBuddyContractPair)
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
func (it *InboxBuddyContractPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxBuddyContractPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxBuddyContractPair represents a BuddyContractPair event raised by the Inbox contract.
type InboxBuddyContractPair struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuddyContractPair is a free log retrieval operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) FilterBuddyContractPair(opts *bind.FilterOpts, sender []common.Address) (*InboxBuddyContractPairIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxBuddyContractPairIterator{contract: _Inbox.contract, event: "BuddyContractPair", logs: logs, sub: sub}, nil
}

// WatchBuddyContractPair is a free log subscription operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) WatchBuddyContractPair(opts *bind.WatchOpts, sink chan<- *InboxBuddyContractPair, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxBuddyContractPair)
				if err := _Inbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
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

// ParseBuddyContractPair is a log parse operation binding the contract event 0x49a9f3e01a6efd03cb8dde057ae548630fe394281202dfc3722eb0b109ccd94f.
//
// Solidity: event BuddyContractPair(address indexed sender)
func (_Inbox *InboxFilterer) ParseBuddyContractPair(log types.Log) (*InboxBuddyContractPair, error) {
	event := new(InboxBuddyContractPair)
	if err := _Inbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the Inbox contract.
type InboxInboxMessageDeliveredIterator struct {
	Event *InboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDelivered)
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
		it.Event = new(InboxInboxMessageDelivered)
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
func (it *InboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the Inbox contract.
type InboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Kind       uint8
	Sender     common.Address
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xe5855b2ddc828cd927ab9e242f891de16135ed173f3ae45ad42179be197f8c8b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredIterator{contract: _Inbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xe5855b2ddc828cd927ab9e242f891de16135ed173f3ae45ad42179be197f8c8b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDelivered)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xe5855b2ddc828cd927ab9e242f891de16135ed173f3ae45ad42179be197f8c8b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, uint8 kind, address sender, bytes data)
func (_Inbox *InboxFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxInboxMessageDelivered, error) {
	event := new(InboxInboxMessageDelivered)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Kind       uint8
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x491fab296ae66721d1d8dfc5f3457bec8c5f2586ce6b355adf0feeb4a4a24bdb.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum, uint8 kind, address sender)
func (_Inbox *InboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredFromOriginIterator{contract: _Inbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0x491fab296ae66721d1d8dfc5f3457bec8c5f2586ce6b355adf0feeb4a4a24bdb.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum, uint8 kind, address sender)
func (_Inbox *InboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDeliveredFromOrigin)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0x491fab296ae66721d1d8dfc5f3457bec8c5f2586ce6b355adf0feeb4a4a24bdb.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum, uint8 kind, address sender)
func (_Inbox *InboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxInboxMessageDeliveredFromOrigin)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxABI is the input ABI used to generate the binding from.
const OutboxABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"processOutgoingMessages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OutboxBin is the compiled bytecode used for deploying new contracts.
var OutboxBin = "0x608060405234801561001057600080fd5b5060405161134e38038061134e8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b038085166001600160a01b0319928316179092556001805492841692909116919091179055604051610078906100bd565b604051809103906000f080158015610094573d6000803e3d6000fd5b50600280546001600160a01b0319166001600160a01b0392909216919091179055506100ca9050565b6105ae80610da083390190565b610cc7806100d96000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630c7268471461005c578063465477901461011c57806372fea01f1461013657806380648b021461022d578063b0f305371461011c575b600080fd5b61011a6004803603604081101561007257600080fd5b810190602081018135600160201b81111561008c57600080fd5b82018360208201111561009e57600080fd5b803590602001918460018302840111600160201b831117156100bf57600080fd5b919390929091602081019035600160201b8111156100dc57600080fd5b8201836020820111156100ee57600080fd5b803590602001918460208302840111600160201b8311171561010f57600080fd5b509092509050610251565b005b610124610336565b60408051918252519081900360200190f35b61011a600480360361012081101561014d57600080fd5b81359190810190604081016020820135600160201b81111561016e57600080fd5b82018360208201111561018057600080fd5b803590602001918460208302840111600160201b831117156101a157600080fd5b919390928235926001600160a01b03602082013581169360408301359091169260608301359260808101359260a082013592909160e081019060c00135600160201b8111156101ef57600080fd5b82018360208201111561020157600080fd5b803590602001918460018302840111600160201b8311171561022257600080fd5b509092509050610345565b610235610336565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b0316331461029e576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b806000805b8281101561032d5761030b8783888888868181106102bd57fe5b90506020020135926102d193929190610c69565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061063092505050565b84848281811061031757fe5b60200291909101359290920191506001016102a3565b50505050505050565b6004546001600160a01b031690565b60008760601b60601c6001600160a01b03168760601b60601c6001600160a01b0316878787878760405160200180888152602001878152602001868152602001858152602001848152602001838380828437808301925050509750505050505050506040516020818303038152906040528051906020012090506104008c8c8c808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508691506107429050565b600480546001600160a01b03808b166001600160a01b0319909216919091178255600580546001600160801b03898116600160801b02818c166001600160801b03199093169290921716179055600154604051639e5d4c4960e01b81528a831693810193845260248101889052606060448201908152606482018790526000949290931692639e5d4c49928c928a928a928a92909190608401848480828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b1580156104d757600080fd5b505af11580156104eb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561051457600080fd5b815160208301805160405192949293830192919084600160201b82111561053a57600080fd5b90830190602082018581111561054f57600080fd5b8251600160201b81118282018810171561056857600080fd5b82525081516020918201929091019080838360005b8381101561059557818101518382015260200161057d565b50505050905090810190601f1680156105c25780820380516001836020036101000a031916815260200191505b506040525050505090508061060c576040805162461bcd60e51b815260206004820152600b60248201526a10d0531317d1905253115160aa1b604482015290519081900360640190fd5b5050600480546001600160a01b031916905550506000600555505050505050505050565b80516000908290829061063f57fe5b01602001516001600160f81b031916141561073f5760006106618260016109dd565b60025490915060009061067c906001600160a01b0316610a36565b60015460408051632f84fd1f60e21b81526001600160a01b0392831660048201526024810186905290519293509083169163be13f47c9160448082019260009290919082900301818387803b1580156106d457600080fd5b505af11580156106e8573d6000803e3d6000fd5b5050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b0394909416939093179092555050505b50565b60006107758484846040516020018082815260200191505060405160208183030381529060405280519060200120610b9b565b600154600380549293506000926001600160a01b0390921691639e5d4c4991908990811061079f57fe5b60009182526020808320909101546040805160248082018a905260448083018d90528351808403820181526064938401855295860180516001600160e01b0316630ad0379b60e01b178152935160e089901b6001600160e01b03191681526001600160a01b0390951660048601818152928601889052606091860191825286519386019390935285519296959491939092608490920191808383895b8381101561085357818101518382015260200161083b565b50505050905090810190601f1680156108805780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b1580156108a157600080fd5b505af11580156108b5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160409081528110156108de57600080fd5b815160208301805160405192949293830192919084600160201b82111561090457600080fd5b90830190602082018581111561091957600080fd5b8251600160201b81118282018810171561093257600080fd5b82525081516020918201929091019080838360005b8381101561095f578181015183820152602001610947565b50505050905090810190601f16801561098c5780820380516001836020036101000a031916815260200191505b50604052505050509050806109d5576040805162461bcd60e51b815260206004820152600a60248201526910d0539517d4d411539160b21b604482015290519081900360640190fd5b505050505050565b60008160200183511015610a2d576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000816001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b158015610a7157600080fd5b505afa158015610a85573d6000803e3d6000fd5b505050506040513d6020811015610a9b57600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b602082015290610b495760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b0e578181015183820152602001610af6565b50505050905090810190601f168015610b3b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008260601b9050604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528160148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f0949350505050565b8251600090610100811115610baf57600080fd5b8260005b82811015610c5f5760028606610c0c57868181518110610bcf57fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150610c51565b81878281518110610c1957fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101610bb3565b5095945050505050565b60008085851115610c78578182fd5b83861115610c84578182fd5b505082019391909203915056fea2646970667358221220b182b8c0473d9dea119d8396143d86f15c4ed5399221af235a477bcb48c6ccbc64736f6c634300060c0033608060405234801561001057600080fd5b50600061001b61007d565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b1916600160a01b179055610081565b3390565b61051e806100906000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c80630ad0379b1461007d5780631c1724fa146100a25780636f791d29146100d3578063715018a6146100db5780638da5cb5b146100e3578063be13f47c14610107578063ebf0c71714610133578063f2fde38b1461014d575b600080fd5b6100a06004803603604081101561009357600080fd5b5080359060200135610173565b005b6100bf600480360360208110156100b857600080fd5b503561027c565b604080519115158252519081900360200190f35b6100bf610291565b6100a06102a1565b6100eb610331565b604080516001600160a01b039092168252519081900360200190f35b6100a06004803603604081101561011d57600080fd5b506001600160a01b038135169060200135610340565b61013b610392565b60408051918252519081900360200190f35b6100a06004803603602081101561016357600080fd5b50356001600160a01b0316610398565b61017b61047e565b6000546001600160a01b039081169116146101cb576040805162461bcd60e51b815260206004820181905260248201526000805160206104a9833981519152604482015290519081900360640190fd5b60008181526002602052604090205460ff161561021f576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b6001548214610260576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b6000908152600260205260409020805460ff1916600117905550565b60026020526000908152604090205460ff1681565b600054600160a01b900460ff1690565b6102a961047e565b6000546001600160a01b039081169116146102f9576040805162461bcd60e51b815260206004820181905260248201526000805160206104a9833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116906000805160206104c9833981519152908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600154610383576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b61038c82610398565b60015550565b60015481565b6103a061047e565b6000546001600160a01b039081169116146103f0576040805162461bcd60e51b815260206004820181905260248201526000805160206104a9833981519152604482015290519081900360640190fd5b6001600160a01b0381166104355760405162461bcd60e51b81526004018080602001828103825260268152602001806104836026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216916000805160206104c983398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220aa10f8d93a58f7094b1091547232e767fa400bfaf95d0de51f2aee435ecf6ba764736f6c634300060c0033"

// DeployOutbox deploys a new Ethereum contract, binding an instance of Outbox to it.
func DeployOutbox(auth *bind.TransactOpts, backend bind.ContractBackend, _rollup common.Address, _bridge common.Address) (common.Address, *types.Transaction, *Outbox, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxBin), backend, _rollup, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// Outbox is an auto generated Go binding around an Ethereum contract.
type Outbox struct {
	OutboxCaller     // Read-only binding to the contract
	OutboxTransactor // Write-only binding to the contract
	OutboxFilterer   // Log filterer for contract events
}

// OutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxSession struct {
	Contract     *Outbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxCallerSession struct {
	Contract *OutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxTransactorSession struct {
	Contract     *OutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxRaw struct {
	Contract *Outbox // Generic contract binding to access the raw methods on
}

// OutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxCallerRaw struct {
	Contract *OutboxCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxTransactorRaw struct {
	Contract *OutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutbox creates a new instance of Outbox, bound to a specific deployed contract.
func NewOutbox(address common.Address, backend bind.ContractBackend) (*Outbox, error) {
	contract, err := bindOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// NewOutboxCaller creates a new read-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxCaller(address common.Address, caller bind.ContractCaller) (*OutboxCaller, error) {
	contract, err := bindOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxCaller{contract: contract}, nil
}

// NewOutboxTransactor creates a new write-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxTransactor, error) {
	contract, err := bindOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxTransactor{contract: contract}, nil
}

// NewOutboxFilterer creates a new log filterer instance of Outbox, bound to a specific deployed contract.
func NewOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxFilterer, error) {
	contract, err := bindOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxFilterer{contract: contract}, nil
}

// bindOutbox binds a generic wrapper to an already deployed contract.
func bindOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.OutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transact(opts, method, params...)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCallerSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x72fea01f.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactor) ExecuteTransaction(opts *bind.TransactOpts, outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeTransaction", outboxIndex, proof, index, l2Sender, destAddr, l2Block, l2Timestamp, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x72fea01f.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxSession) ExecuteTransaction(outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, proof, index, l2Sender, destAddr, l2Block, l2Timestamp, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x72fea01f.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactorSession) ExecuteTransaction(outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, proof, index, l2Sender, destAddr, l2Block, l2Timestamp, amount, calldataForL1)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxTransactor) ProcessOutgoingMessages(opts *bind.TransactOpts, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "processOutgoingMessages", sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.Contract.ProcessOutgoingMessages(&_Outbox.TransactOpts, sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxTransactorSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.Contract.ProcessOutgoingMessages(&_Outbox.TransactOpts, sendsData, sendLengths)
}

// OutboxEntryABI is the input ABI used to generate the binding from.
const OutboxEntryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"spendOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spentOutput\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OutboxEntryBin is the compiled bytecode used for deploying new contracts.
var OutboxEntryBin = "0x608060405234801561001057600080fd5b50600061001b61007d565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b1916600160a01b179055610081565b3390565b61051e806100906000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c80630ad0379b1461007d5780631c1724fa146100a25780636f791d29146100d3578063715018a6146100db5780638da5cb5b146100e3578063be13f47c14610107578063ebf0c71714610133578063f2fde38b1461014d575b600080fd5b6100a06004803603604081101561009357600080fd5b5080359060200135610173565b005b6100bf600480360360208110156100b857600080fd5b503561027c565b604080519115158252519081900360200190f35b6100bf610291565b6100a06102a1565b6100eb610331565b604080516001600160a01b039092168252519081900360200190f35b6100a06004803603604081101561011d57600080fd5b506001600160a01b038135169060200135610340565b61013b610392565b60408051918252519081900360200190f35b6100a06004803603602081101561016357600080fd5b50356001600160a01b0316610398565b61017b61047e565b6000546001600160a01b039081169116146101cb576040805162461bcd60e51b815260206004820181905260248201526000805160206104a9833981519152604482015290519081900360640190fd5b60008181526002602052604090205460ff161561021f576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b6001548214610260576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b6000908152600260205260409020805460ff1916600117905550565b60026020526000908152604090205460ff1681565b600054600160a01b900460ff1690565b6102a961047e565b6000546001600160a01b039081169116146102f9576040805162461bcd60e51b815260206004820181905260248201526000805160206104a9833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116906000805160206104c9833981519152908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600154610383576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b61038c82610398565b60015550565b60015481565b6103a061047e565b6000546001600160a01b039081169116146103f0576040805162461bcd60e51b815260206004820181905260248201526000805160206104a9833981519152604482015290519081900360640190fd5b6001600160a01b0381166104355760405162461bcd60e51b81526004018080602001828103825260268152602001806104836026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216916000805160206104c983398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220aa10f8d93a58f7094b1091547232e767fa400bfaf95d0de51f2aee435ecf6ba764736f6c634300060c0033"

// DeployOutboxEntry deploys a new Ethereum contract, binding an instance of OutboxEntry to it.
func DeployOutboxEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OutboxEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// OutboxEntry is an auto generated Go binding around an Ethereum contract.
type OutboxEntry struct {
	OutboxEntryCaller     // Read-only binding to the contract
	OutboxEntryTransactor // Write-only binding to the contract
	OutboxEntryFilterer   // Log filterer for contract events
}

// OutboxEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxEntrySession struct {
	Contract     *OutboxEntry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxEntryCallerSession struct {
	Contract *OutboxEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OutboxEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxEntryTransactorSession struct {
	Contract     *OutboxEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OutboxEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxEntryRaw struct {
	Contract *OutboxEntry // Generic contract binding to access the raw methods on
}

// OutboxEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxEntryCallerRaw struct {
	Contract *OutboxEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxEntryTransactorRaw struct {
	Contract *OutboxEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutboxEntry creates a new instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntry(address common.Address, backend bind.ContractBackend) (*OutboxEntry, error) {
	contract, err := bindOutboxEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// NewOutboxEntryCaller creates a new read-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryCaller(address common.Address, caller bind.ContractCaller) (*OutboxEntryCaller, error) {
	contract, err := bindOutboxEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryCaller{contract: contract}, nil
}

// NewOutboxEntryTransactor creates a new write-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxEntryTransactor, error) {
	contract, err := bindOutboxEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryTransactor{contract: contract}, nil
}

// NewOutboxEntryFilterer creates a new log filterer instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxEntryFilterer, error) {
	contract, err := bindOutboxEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryFilterer{contract: contract}, nil
}

// bindOutboxEntry binds a generic wrapper to an already deployed contract.
func bindOutboxEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.OutboxEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transact(opts, method, params...)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntryCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntrySession) IsMaster() (bool, error) {
	return _OutboxEntry.Contract.IsMaster(&_OutboxEntry.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntryCallerSession) IsMaster() (bool, error) {
	return _OutboxEntry.Contract.IsMaster(&_OutboxEntry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OutboxEntry *OutboxEntryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OutboxEntry *OutboxEntrySession) Owner() (common.Address, error) {
	return _OutboxEntry.Contract.Owner(&_OutboxEntry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OutboxEntry *OutboxEntryCallerSession) Owner() (common.Address, error) {
	return _OutboxEntry.Contract.Owner(&_OutboxEntry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntryCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntrySession) Root() ([32]byte, error) {
	return _OutboxEntry.Contract.Root(&_OutboxEntry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntryCallerSession) Root() ([32]byte, error) {
	return _OutboxEntry.Contract.Root(&_OutboxEntry.CallOpts)
}

// SpentOutput is a free data retrieval call binding the contract method 0x1c1724fa.
//
// Solidity: function spentOutput(uint256 ) view returns(bool)
func (_OutboxEntry *OutboxEntryCaller) SpentOutput(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "spentOutput", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SpentOutput is a free data retrieval call binding the contract method 0x1c1724fa.
//
// Solidity: function spentOutput(uint256 ) view returns(bool)
func (_OutboxEntry *OutboxEntrySession) SpentOutput(arg0 *big.Int) (bool, error) {
	return _OutboxEntry.Contract.SpentOutput(&_OutboxEntry.CallOpts, arg0)
}

// SpentOutput is a free data retrieval call binding the contract method 0x1c1724fa.
//
// Solidity: function spentOutput(uint256 ) view returns(bool)
func (_OutboxEntry *OutboxEntryCallerSession) SpentOutput(arg0 *big.Int) (bool, error) {
	return _OutboxEntry.Contract.SpentOutput(&_OutboxEntry.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(address _owner, bytes32 _root) returns()
func (_OutboxEntry *OutboxEntryTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _root [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "initialize", _owner, _root)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(address _owner, bytes32 _root) returns()
func (_OutboxEntry *OutboxEntrySession) Initialize(_owner common.Address, _root [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.Contract.Initialize(&_OutboxEntry.TransactOpts, _owner, _root)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(address _owner, bytes32 _root) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) Initialize(_owner common.Address, _root [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.Contract.Initialize(&_OutboxEntry.TransactOpts, _owner, _root)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OutboxEntry *OutboxEntryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OutboxEntry *OutboxEntrySession) RenounceOwnership() (*types.Transaction, error) {
	return _OutboxEntry.Contract.RenounceOwnership(&_OutboxEntry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OutboxEntry *OutboxEntryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OutboxEntry.Contract.RenounceOwnership(&_OutboxEntry.TransactOpts)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 _root, uint256 _id) returns()
func (_OutboxEntry *OutboxEntryTransactor) SpendOutput(opts *bind.TransactOpts, _root [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "spendOutput", _root, _id)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 _root, uint256 _id) returns()
func (_OutboxEntry *OutboxEntrySession) SpendOutput(_root [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, _root, _id)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x0ad0379b.
//
// Solidity: function spendOutput(bytes32 _root, uint256 _id) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) SpendOutput(_root [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, _root, _id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OutboxEntry *OutboxEntryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OutboxEntry *OutboxEntrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OutboxEntry.Contract.TransferOwnership(&_OutboxEntry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OutboxEntry.Contract.TransferOwnership(&_OutboxEntry.TransactOpts, newOwner)
}

// OutboxEntryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OutboxEntry contract.
type OutboxEntryOwnershipTransferredIterator struct {
	Event *OutboxEntryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OutboxEntryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxEntryOwnershipTransferred)
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
		it.Event = new(OutboxEntryOwnershipTransferred)
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
func (it *OutboxEntryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxEntryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxEntryOwnershipTransferred represents a OwnershipTransferred event raised by the OutboxEntry contract.
type OutboxEntryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OutboxEntry *OutboxEntryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OutboxEntryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OutboxEntry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryOwnershipTransferredIterator{contract: _OutboxEntry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OutboxEntry *OutboxEntryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OutboxEntryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OutboxEntry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxEntryOwnershipTransferred)
				if err := _OutboxEntry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OutboxEntry *OutboxEntryFilterer) ParseOwnershipTransferred(log types.Log) (*OutboxEntryOwnershipTransferred, error) {
	event := new(OutboxEntryOwnershipTransferred)
	if err := _OutboxEntry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
