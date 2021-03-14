// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

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

// InboxHelperABI is the input ABI used to generate the binding from.
const InboxHelperABI = "[]"

// InboxHelperBin is the compiled bytecode used for deploying new contracts.
var InboxHelperBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209e460ac9f8760b35998f5f3dbece65268c6e6c1d44cc1fb15db695b8b7e95e4164736f6c634300060c0033"

// DeployInboxHelper deploys a new Ethereum contract, binding an instance of InboxHelper to it.
func DeployInboxHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InboxHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InboxHelper{InboxHelperCaller: InboxHelperCaller{contract: contract}, InboxHelperTransactor: InboxHelperTransactor{contract: contract}, InboxHelperFilterer: InboxHelperFilterer{contract: contract}}, nil
}

// InboxHelper is an auto generated Go binding around an Ethereum contract.
type InboxHelper struct {
	InboxHelperCaller     // Read-only binding to the contract
	InboxHelperTransactor // Write-only binding to the contract
	InboxHelperFilterer   // Log filterer for contract events
}

// InboxHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxHelperSession struct {
	Contract     *InboxHelper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxHelperCallerSession struct {
	Contract *InboxHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InboxHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxHelperTransactorSession struct {
	Contract     *InboxHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InboxHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxHelperRaw struct {
	Contract *InboxHelper // Generic contract binding to access the raw methods on
}

// InboxHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxHelperCallerRaw struct {
	Contract *InboxHelperCaller // Generic read-only contract binding to access the raw methods on
}

// InboxHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxHelperTransactorRaw struct {
	Contract *InboxHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInboxHelper creates a new instance of InboxHelper, bound to a specific deployed contract.
func NewInboxHelper(address common.Address, backend bind.ContractBackend) (*InboxHelper, error) {
	contract, err := bindInboxHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InboxHelper{InboxHelperCaller: InboxHelperCaller{contract: contract}, InboxHelperTransactor: InboxHelperTransactor{contract: contract}, InboxHelperFilterer: InboxHelperFilterer{contract: contract}}, nil
}

// NewInboxHelperCaller creates a new read-only instance of InboxHelper, bound to a specific deployed contract.
func NewInboxHelperCaller(address common.Address, caller bind.ContractCaller) (*InboxHelperCaller, error) {
	contract, err := bindInboxHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxHelperCaller{contract: contract}, nil
}

// NewInboxHelperTransactor creates a new write-only instance of InboxHelper, bound to a specific deployed contract.
func NewInboxHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxHelperTransactor, error) {
	contract, err := bindInboxHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxHelperTransactor{contract: contract}, nil
}

// NewInboxHelperFilterer creates a new log filterer instance of InboxHelper, bound to a specific deployed contract.
func NewInboxHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxHelperFilterer, error) {
	contract, err := bindInboxHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxHelperFilterer{contract: contract}, nil
}

// bindInboxHelper binds a generic wrapper to an already deployed contract.
func bindInboxHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxHelper *InboxHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxHelper.Contract.InboxHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxHelper *InboxHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxHelper.Contract.InboxHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxHelper *InboxHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxHelper.Contract.InboxHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxHelper *InboxHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxHelper *InboxHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxHelper *InboxHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxHelper.Contract.contract.Transact(opts, method, params...)
}

// InboxHelperTesterABI is the input ABI used to generate the binding from.
const InboxHelperTesterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"requestID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"retryableTicketID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// InboxHelperTesterFuncSigs maps the 4-byte function signature to its string representation.
var InboxHelperTesterFuncSigs = map[string]string{
	"a64371ed": "chainId(address)",
	"9c829800": "requestID(uint256,address)",
	"d96e2802": "retryableTicketID(uint256,address)",
}

// InboxHelperTesterBin is the compiled bytecode used for deploying new contracts.
var InboxHelperTesterBin = "0x608060405234801561001057600080fd5b506101c4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80639c82980014610046578063a64371ed14610084578063d96e2802146100aa575b600080fd5b6100726004803603604081101561005c57600080fd5b50803590602001356001600160a01b03166100d6565b60408051918252519081900360200190f35b6100726004803603602081101561009a57600080fd5b50356001600160a01b03166100e9565b610072600480360360408110156100c057600080fd5b50803590602001356001600160a01b03166100fa565b60006100e28383610106565b9392505050565b60006100f482610147565b92915050565b60006100e28383610152565b600061011182610147565b83604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120905092915050565b65ffffffffffff1690565b600061015e8383610106565b6040805160208082019390935260008183015281518082038301815260609091019091528051910120939250505056fea26469706673582212201c7d3cd46ef14ced44c16d993cbb68c29ca85137f10721c86457636865eb684164736f6c634300060c0033"

// DeployInboxHelperTester deploys a new Ethereum contract, binding an instance of InboxHelperTester to it.
func DeployInboxHelperTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InboxHelperTester, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxHelperTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxHelperTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InboxHelperTester{InboxHelperTesterCaller: InboxHelperTesterCaller{contract: contract}, InboxHelperTesterTransactor: InboxHelperTesterTransactor{contract: contract}, InboxHelperTesterFilterer: InboxHelperTesterFilterer{contract: contract}}, nil
}

// InboxHelperTester is an auto generated Go binding around an Ethereum contract.
type InboxHelperTester struct {
	InboxHelperTesterCaller     // Read-only binding to the contract
	InboxHelperTesterTransactor // Write-only binding to the contract
	InboxHelperTesterFilterer   // Log filterer for contract events
}

// InboxHelperTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxHelperTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxHelperTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxHelperTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxHelperTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxHelperTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxHelperTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxHelperTesterSession struct {
	Contract     *InboxHelperTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InboxHelperTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxHelperTesterCallerSession struct {
	Contract *InboxHelperTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InboxHelperTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxHelperTesterTransactorSession struct {
	Contract     *InboxHelperTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InboxHelperTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxHelperTesterRaw struct {
	Contract *InboxHelperTester // Generic contract binding to access the raw methods on
}

// InboxHelperTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxHelperTesterCallerRaw struct {
	Contract *InboxHelperTesterCaller // Generic read-only contract binding to access the raw methods on
}

// InboxHelperTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxHelperTesterTransactorRaw struct {
	Contract *InboxHelperTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInboxHelperTester creates a new instance of InboxHelperTester, bound to a specific deployed contract.
func NewInboxHelperTester(address common.Address, backend bind.ContractBackend) (*InboxHelperTester, error) {
	contract, err := bindInboxHelperTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InboxHelperTester{InboxHelperTesterCaller: InboxHelperTesterCaller{contract: contract}, InboxHelperTesterTransactor: InboxHelperTesterTransactor{contract: contract}, InboxHelperTesterFilterer: InboxHelperTesterFilterer{contract: contract}}, nil
}

// NewInboxHelperTesterCaller creates a new read-only instance of InboxHelperTester, bound to a specific deployed contract.
func NewInboxHelperTesterCaller(address common.Address, caller bind.ContractCaller) (*InboxHelperTesterCaller, error) {
	contract, err := bindInboxHelperTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxHelperTesterCaller{contract: contract}, nil
}

// NewInboxHelperTesterTransactor creates a new write-only instance of InboxHelperTester, bound to a specific deployed contract.
func NewInboxHelperTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxHelperTesterTransactor, error) {
	contract, err := bindInboxHelperTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxHelperTesterTransactor{contract: contract}, nil
}

// NewInboxHelperTesterFilterer creates a new log filterer instance of InboxHelperTester, bound to a specific deployed contract.
func NewInboxHelperTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxHelperTesterFilterer, error) {
	contract, err := bindInboxHelperTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxHelperTesterFilterer{contract: contract}, nil
}

// bindInboxHelperTester binds a generic wrapper to an already deployed contract.
func bindInboxHelperTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxHelperTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxHelperTester *InboxHelperTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxHelperTester.Contract.InboxHelperTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxHelperTester *InboxHelperTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxHelperTester.Contract.InboxHelperTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxHelperTester *InboxHelperTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxHelperTester.Contract.InboxHelperTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxHelperTester *InboxHelperTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxHelperTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxHelperTester *InboxHelperTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxHelperTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxHelperTester *InboxHelperTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxHelperTester.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0xa64371ed.
//
// Solidity: function chainId(address rollup) pure returns(uint256)
func (_InboxHelperTester *InboxHelperTesterCaller) ChainId(opts *bind.CallOpts, rollup common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InboxHelperTester.contract.Call(opts, &out, "chainId", rollup)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0xa64371ed.
//
// Solidity: function chainId(address rollup) pure returns(uint256)
func (_InboxHelperTester *InboxHelperTesterSession) ChainId(rollup common.Address) (*big.Int, error) {
	return _InboxHelperTester.Contract.ChainId(&_InboxHelperTester.CallOpts, rollup)
}

// ChainId is a free data retrieval call binding the contract method 0xa64371ed.
//
// Solidity: function chainId(address rollup) pure returns(uint256)
func (_InboxHelperTester *InboxHelperTesterCallerSession) ChainId(rollup common.Address) (*big.Int, error) {
	return _InboxHelperTester.Contract.ChainId(&_InboxHelperTester.CallOpts, rollup)
}

// RequestID is a free data retrieval call binding the contract method 0x9c829800.
//
// Solidity: function requestID(uint256 messageNum, address rollup) pure returns(bytes32)
func (_InboxHelperTester *InboxHelperTesterCaller) RequestID(opts *bind.CallOpts, messageNum *big.Int, rollup common.Address) ([32]byte, error) {
	var out []interface{}
	err := _InboxHelperTester.contract.Call(opts, &out, "requestID", messageNum, rollup)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestID is a free data retrieval call binding the contract method 0x9c829800.
//
// Solidity: function requestID(uint256 messageNum, address rollup) pure returns(bytes32)
func (_InboxHelperTester *InboxHelperTesterSession) RequestID(messageNum *big.Int, rollup common.Address) ([32]byte, error) {
	return _InboxHelperTester.Contract.RequestID(&_InboxHelperTester.CallOpts, messageNum, rollup)
}

// RequestID is a free data retrieval call binding the contract method 0x9c829800.
//
// Solidity: function requestID(uint256 messageNum, address rollup) pure returns(bytes32)
func (_InboxHelperTester *InboxHelperTesterCallerSession) RequestID(messageNum *big.Int, rollup common.Address) ([32]byte, error) {
	return _InboxHelperTester.Contract.RequestID(&_InboxHelperTester.CallOpts, messageNum, rollup)
}

// RetryableTicketID is a free data retrieval call binding the contract method 0xd96e2802.
//
// Solidity: function retryableTicketID(uint256 messageNum, address rollup) pure returns(bytes32)
func (_InboxHelperTester *InboxHelperTesterCaller) RetryableTicketID(opts *bind.CallOpts, messageNum *big.Int, rollup common.Address) ([32]byte, error) {
	var out []interface{}
	err := _InboxHelperTester.contract.Call(opts, &out, "retryableTicketID", messageNum, rollup)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RetryableTicketID is a free data retrieval call binding the contract method 0xd96e2802.
//
// Solidity: function retryableTicketID(uint256 messageNum, address rollup) pure returns(bytes32)
func (_InboxHelperTester *InboxHelperTesterSession) RetryableTicketID(messageNum *big.Int, rollup common.Address) ([32]byte, error) {
	return _InboxHelperTester.Contract.RetryableTicketID(&_InboxHelperTester.CallOpts, messageNum, rollup)
}

// RetryableTicketID is a free data retrieval call binding the contract method 0xd96e2802.
//
// Solidity: function retryableTicketID(uint256 messageNum, address rollup) pure returns(bytes32)
func (_InboxHelperTester *InboxHelperTesterCallerSession) RetryableTicketID(messageNum *big.Int, rollup common.Address) ([32]byte, error) {
	return _InboxHelperTester.Contract.RetryableTicketID(&_InboxHelperTester.CallOpts, messageNum, rollup)
}
