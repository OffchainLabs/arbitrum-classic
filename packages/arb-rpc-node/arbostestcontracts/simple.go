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

// ComplexConstructorConABI is the input ABI used to generate the binding from.
const ComplexConstructorConABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getVal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ComplexConstructorConBin is the compiled bytecode used for deploying new contracts.
var ComplexConstructorConBin = "0x60806040526040516102b13803806102b18339818101604052602081101561002657600080fd5b50516040805163099f12b960e21b81529051339163267c4ae49160048083019260209291908290030181600087803b15801561006157600080fd5b505af1158015610075573d6000803e3d6000fd5b505050506040513d602081101561008b57600080fd5b508190506002340461028e6040516100a290610127565b90815260405183918190036020019083f5915050801580156100c8573d6000803e3d6000fd5b505060408051639b7c9da360e01b81526036600482015290513391639b7c9da391602480830192600092919082900301818387803b15801561010957600080fd5b505af115801561011d573d6000803e3d6000fd5b5050505050610133565b60e7806101ca83390190565b6089806101416000396000f3fe608060405260043610601f5760003560e01c8063e1cb0e5214602a576025565b36602557005b600080fd5b348015603557600080fd5b50603c604e565b60408051918252519081900360200190f35b60149056fea2646970667358221220a4313dc2605287a7a0062ed358ddbd15728ce272ca8a5e4f9dd085d9bc90a43264736f6c634300060c003360806040526040516100e73803806100e78339818101604052602081101561002657600080fd5b505160405133906002340480156108fc02916000818181858888f19350505050158015610057573d6000803e3d6000fd5b50506080806100676000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b60336045565b60408051918252519081900360200190f35b60149056fea2646970667358221220987822a2268d74f17c5d03c1c82a985b32037cf6e762369f81bd42f1f4093a7664736f6c634300060c0033"

// DeployComplexConstructorCon deploys a new Ethereum contract, binding an instance of ComplexConstructorCon to it.
func DeployComplexConstructorCon(auth *bind.TransactOpts, backend bind.ContractBackend, salt [32]byte) (common.Address, *types.Transaction, *ComplexConstructorCon, error) {
	parsed, err := abi.JSON(strings.NewReader(ComplexConstructorConABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ComplexConstructorConBin), backend, salt)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ComplexConstructorCon{ComplexConstructorConCaller: ComplexConstructorConCaller{contract: contract}, ComplexConstructorConTransactor: ComplexConstructorConTransactor{contract: contract}, ComplexConstructorConFilterer: ComplexConstructorConFilterer{contract: contract}}, nil
}

// ComplexConstructorCon is an auto generated Go binding around an Ethereum contract.
type ComplexConstructorCon struct {
	ComplexConstructorConCaller     // Read-only binding to the contract
	ComplexConstructorConTransactor // Write-only binding to the contract
	ComplexConstructorConFilterer   // Log filterer for contract events
}

// ComplexConstructorConCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComplexConstructorConCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplexConstructorConTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComplexConstructorConTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplexConstructorConFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComplexConstructorConFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplexConstructorConSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComplexConstructorConSession struct {
	Contract     *ComplexConstructorCon // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ComplexConstructorConCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComplexConstructorConCallerSession struct {
	Contract *ComplexConstructorConCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ComplexConstructorConTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComplexConstructorConTransactorSession struct {
	Contract     *ComplexConstructorConTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ComplexConstructorConRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComplexConstructorConRaw struct {
	Contract *ComplexConstructorCon // Generic contract binding to access the raw methods on
}

// ComplexConstructorConCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComplexConstructorConCallerRaw struct {
	Contract *ComplexConstructorConCaller // Generic read-only contract binding to access the raw methods on
}

// ComplexConstructorConTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComplexConstructorConTransactorRaw struct {
	Contract *ComplexConstructorConTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComplexConstructorCon creates a new instance of ComplexConstructorCon, bound to a specific deployed contract.
func NewComplexConstructorCon(address common.Address, backend bind.ContractBackend) (*ComplexConstructorCon, error) {
	contract, err := bindComplexConstructorCon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComplexConstructorCon{ComplexConstructorConCaller: ComplexConstructorConCaller{contract: contract}, ComplexConstructorConTransactor: ComplexConstructorConTransactor{contract: contract}, ComplexConstructorConFilterer: ComplexConstructorConFilterer{contract: contract}}, nil
}

// NewComplexConstructorConCaller creates a new read-only instance of ComplexConstructorCon, bound to a specific deployed contract.
func NewComplexConstructorConCaller(address common.Address, caller bind.ContractCaller) (*ComplexConstructorConCaller, error) {
	contract, err := bindComplexConstructorCon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComplexConstructorConCaller{contract: contract}, nil
}

// NewComplexConstructorConTransactor creates a new write-only instance of ComplexConstructorCon, bound to a specific deployed contract.
func NewComplexConstructorConTransactor(address common.Address, transactor bind.ContractTransactor) (*ComplexConstructorConTransactor, error) {
	contract, err := bindComplexConstructorCon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComplexConstructorConTransactor{contract: contract}, nil
}

// NewComplexConstructorConFilterer creates a new log filterer instance of ComplexConstructorCon, bound to a specific deployed contract.
func NewComplexConstructorConFilterer(address common.Address, filterer bind.ContractFilterer) (*ComplexConstructorConFilterer, error) {
	contract, err := bindComplexConstructorCon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComplexConstructorConFilterer{contract: contract}, nil
}

// bindComplexConstructorCon binds a generic wrapper to an already deployed contract.
func bindComplexConstructorCon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComplexConstructorConABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplexConstructorCon *ComplexConstructorConRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComplexConstructorCon.Contract.ComplexConstructorConCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplexConstructorCon *ComplexConstructorConRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplexConstructorCon.Contract.ComplexConstructorConTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplexConstructorCon *ComplexConstructorConRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplexConstructorCon.Contract.ComplexConstructorConTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplexConstructorCon *ComplexConstructorConCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComplexConstructorCon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplexConstructorCon *ComplexConstructorConTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplexConstructorCon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplexConstructorCon *ComplexConstructorConTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplexConstructorCon.Contract.contract.Transact(opts, method, params...)
}

// GetVal is a paid mutator transaction binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() returns(uint256)
func (_ComplexConstructorCon *ComplexConstructorConTransactor) GetVal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplexConstructorCon.contract.Transact(opts, "getVal")
}

// GetVal is a paid mutator transaction binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() returns(uint256)
func (_ComplexConstructorCon *ComplexConstructorConSession) GetVal() (*types.Transaction, error) {
	return _ComplexConstructorCon.Contract.GetVal(&_ComplexConstructorCon.TransactOpts)
}

// GetVal is a paid mutator transaction binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() returns(uint256)
func (_ComplexConstructorCon *ComplexConstructorConTransactorSession) GetVal() (*types.Transaction, error) {
	return _ComplexConstructorCon.Contract.GetVal(&_ComplexConstructorCon.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ComplexConstructorCon *ComplexConstructorConTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplexConstructorCon.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ComplexConstructorCon *ComplexConstructorConSession) Receive() (*types.Transaction, error) {
	return _ComplexConstructorCon.Contract.Receive(&_ComplexConstructorCon.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ComplexConstructorCon *ComplexConstructorConTransactorSession) Receive() (*types.Transaction, error) {
	return _ComplexConstructorCon.Contract.Receive(&_ComplexConstructorCon.TransactOpts)
}

// ComplexConstructorCon2ABI is the input ABI used to generate the binding from.
const ComplexConstructorCon2ABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getVal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ComplexConstructorCon2Bin is the compiled bytecode used for deploying new contracts.
var ComplexConstructorCon2Bin = "0x60806040526040516100e73803806100e78339818101604052602081101561002657600080fd5b505160405133906002340480156108fc02916000818181858888f19350505050158015610057573d6000803e3d6000fd5b50506080806100676000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b60336045565b60408051918252519081900360200190f35b60149056fea2646970667358221220987822a2268d74f17c5d03c1c82a985b32037cf6e762369f81bd42f1f4093a7664736f6c634300060c0033"

// DeployComplexConstructorCon2 deploys a new Ethereum contract, binding an instance of ComplexConstructorCon2 to it.
func DeployComplexConstructorCon2(auth *bind.TransactOpts, backend bind.ContractBackend, val *big.Int) (common.Address, *types.Transaction, *ComplexConstructorCon2, error) {
	parsed, err := abi.JSON(strings.NewReader(ComplexConstructorCon2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ComplexConstructorCon2Bin), backend, val)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ComplexConstructorCon2{ComplexConstructorCon2Caller: ComplexConstructorCon2Caller{contract: contract}, ComplexConstructorCon2Transactor: ComplexConstructorCon2Transactor{contract: contract}, ComplexConstructorCon2Filterer: ComplexConstructorCon2Filterer{contract: contract}}, nil
}

// ComplexConstructorCon2 is an auto generated Go binding around an Ethereum contract.
type ComplexConstructorCon2 struct {
	ComplexConstructorCon2Caller     // Read-only binding to the contract
	ComplexConstructorCon2Transactor // Write-only binding to the contract
	ComplexConstructorCon2Filterer   // Log filterer for contract events
}

// ComplexConstructorCon2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ComplexConstructorCon2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplexConstructorCon2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ComplexConstructorCon2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplexConstructorCon2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComplexConstructorCon2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplexConstructorCon2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComplexConstructorCon2Session struct {
	Contract     *ComplexConstructorCon2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ComplexConstructorCon2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComplexConstructorCon2CallerSession struct {
	Contract *ComplexConstructorCon2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ComplexConstructorCon2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComplexConstructorCon2TransactorSession struct {
	Contract     *ComplexConstructorCon2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ComplexConstructorCon2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ComplexConstructorCon2Raw struct {
	Contract *ComplexConstructorCon2 // Generic contract binding to access the raw methods on
}

// ComplexConstructorCon2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComplexConstructorCon2CallerRaw struct {
	Contract *ComplexConstructorCon2Caller // Generic read-only contract binding to access the raw methods on
}

// ComplexConstructorCon2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComplexConstructorCon2TransactorRaw struct {
	Contract *ComplexConstructorCon2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewComplexConstructorCon2 creates a new instance of ComplexConstructorCon2, bound to a specific deployed contract.
func NewComplexConstructorCon2(address common.Address, backend bind.ContractBackend) (*ComplexConstructorCon2, error) {
	contract, err := bindComplexConstructorCon2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComplexConstructorCon2{ComplexConstructorCon2Caller: ComplexConstructorCon2Caller{contract: contract}, ComplexConstructorCon2Transactor: ComplexConstructorCon2Transactor{contract: contract}, ComplexConstructorCon2Filterer: ComplexConstructorCon2Filterer{contract: contract}}, nil
}

// NewComplexConstructorCon2Caller creates a new read-only instance of ComplexConstructorCon2, bound to a specific deployed contract.
func NewComplexConstructorCon2Caller(address common.Address, caller bind.ContractCaller) (*ComplexConstructorCon2Caller, error) {
	contract, err := bindComplexConstructorCon2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComplexConstructorCon2Caller{contract: contract}, nil
}

// NewComplexConstructorCon2Transactor creates a new write-only instance of ComplexConstructorCon2, bound to a specific deployed contract.
func NewComplexConstructorCon2Transactor(address common.Address, transactor bind.ContractTransactor) (*ComplexConstructorCon2Transactor, error) {
	contract, err := bindComplexConstructorCon2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComplexConstructorCon2Transactor{contract: contract}, nil
}

// NewComplexConstructorCon2Filterer creates a new log filterer instance of ComplexConstructorCon2, bound to a specific deployed contract.
func NewComplexConstructorCon2Filterer(address common.Address, filterer bind.ContractFilterer) (*ComplexConstructorCon2Filterer, error) {
	contract, err := bindComplexConstructorCon2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComplexConstructorCon2Filterer{contract: contract}, nil
}

// bindComplexConstructorCon2 binds a generic wrapper to an already deployed contract.
func bindComplexConstructorCon2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComplexConstructorCon2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplexConstructorCon2 *ComplexConstructorCon2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComplexConstructorCon2.Contract.ComplexConstructorCon2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplexConstructorCon2 *ComplexConstructorCon2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplexConstructorCon2.Contract.ComplexConstructorCon2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplexConstructorCon2 *ComplexConstructorCon2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplexConstructorCon2.Contract.ComplexConstructorCon2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplexConstructorCon2 *ComplexConstructorCon2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComplexConstructorCon2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplexConstructorCon2 *ComplexConstructorCon2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplexConstructorCon2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplexConstructorCon2 *ComplexConstructorCon2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplexConstructorCon2.Contract.contract.Transact(opts, method, params...)
}

// GetVal is a paid mutator transaction binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() returns(uint256)
func (_ComplexConstructorCon2 *ComplexConstructorCon2Transactor) GetVal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplexConstructorCon2.contract.Transact(opts, "getVal")
}

// GetVal is a paid mutator transaction binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() returns(uint256)
func (_ComplexConstructorCon2 *ComplexConstructorCon2Session) GetVal() (*types.Transaction, error) {
	return _ComplexConstructorCon2.Contract.GetVal(&_ComplexConstructorCon2.TransactOpts)
}

// GetVal is a paid mutator transaction binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() returns(uint256)
func (_ComplexConstructorCon2 *ComplexConstructorCon2TransactorSession) GetVal() (*types.Transaction, error) {
	return _ComplexConstructorCon2.Contract.GetVal(&_ComplexConstructorCon2.TransactOpts)
}

// ReverterABI is the input ABI used to generate the binding from.
const ReverterABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReverterBin is the compiled bytecode used for deploying new contracts.
var ReverterBin = "0x6080604052348015600f57600080fd5b506040805162461bcd60e51b8152602060048201526012602482015271125b9d195b9d1a5bdb985b081c995d995c9d60721b604482015290519081900360640190fdfe"

// DeployReverter deploys a new Ethereum contract, binding an instance of Reverter to it.
func DeployReverter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reverter, error) {
	parsed, err := abi.JSON(strings.NewReader(ReverterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReverterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reverter{ReverterCaller: ReverterCaller{contract: contract}, ReverterTransactor: ReverterTransactor{contract: contract}, ReverterFilterer: ReverterFilterer{contract: contract}}, nil
}

// Reverter is an auto generated Go binding around an Ethereum contract.
type Reverter struct {
	ReverterCaller     // Read-only binding to the contract
	ReverterTransactor // Write-only binding to the contract
	ReverterFilterer   // Log filterer for contract events
}

// ReverterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReverterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReverterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReverterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReverterSession struct {
	Contract     *Reverter         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReverterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReverterCallerSession struct {
	Contract *ReverterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ReverterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReverterTransactorSession struct {
	Contract     *ReverterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReverterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReverterRaw struct {
	Contract *Reverter // Generic contract binding to access the raw methods on
}

// ReverterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReverterCallerRaw struct {
	Contract *ReverterCaller // Generic read-only contract binding to access the raw methods on
}

// ReverterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReverterTransactorRaw struct {
	Contract *ReverterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReverter creates a new instance of Reverter, bound to a specific deployed contract.
func NewReverter(address common.Address, backend bind.ContractBackend) (*Reverter, error) {
	contract, err := bindReverter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reverter{ReverterCaller: ReverterCaller{contract: contract}, ReverterTransactor: ReverterTransactor{contract: contract}, ReverterFilterer: ReverterFilterer{contract: contract}}, nil
}

// NewReverterCaller creates a new read-only instance of Reverter, bound to a specific deployed contract.
func NewReverterCaller(address common.Address, caller bind.ContractCaller) (*ReverterCaller, error) {
	contract, err := bindReverter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReverterCaller{contract: contract}, nil
}

// NewReverterTransactor creates a new write-only instance of Reverter, bound to a specific deployed contract.
func NewReverterTransactor(address common.Address, transactor bind.ContractTransactor) (*ReverterTransactor, error) {
	contract, err := bindReverter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReverterTransactor{contract: contract}, nil
}

// NewReverterFilterer creates a new log filterer instance of Reverter, bound to a specific deployed contract.
func NewReverterFilterer(address common.Address, filterer bind.ContractFilterer) (*ReverterFilterer, error) {
	contract, err := bindReverter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReverterFilterer{contract: contract}, nil
}

// bindReverter binds a generic wrapper to an already deployed contract.
func bindReverter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReverterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reverter *ReverterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reverter.Contract.ReverterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reverter *ReverterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reverter.Contract.ReverterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reverter *ReverterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reverter.Contract.ReverterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reverter *ReverterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reverter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reverter *ReverterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reverter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reverter *ReverterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reverter.Contract.contract.Transact(opts, method, params...)
}

// SimpleABI is the input ABI used to generate the binding from.
const SimpleABI = "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractSimple\",\"name\":\"con\",\"type\":\"address\"}],\"name\":\"crossCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"nestedCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"nestedCall2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rejectPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reverts\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg\",\"type\":\"uint256\"}],\"name\":\"trace\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SimpleBin is the compiled bytecode used for deploying new contracts.
var SimpleBin = "0x6080604081905234600181905581527f1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d241490602090a1610883806100436000396000f3fe60806040526004361061008a5760003560e01c80639b7c9da3116100595780639b7c9da31461013b578063a4ab282d14610165578063a56dfe4a14610213578063a68a4fed14610228578063ae0aba8c14610245576100cc565b80630324332e146100d1578063267c4ae4146101165780633bccbbc91461011e5780639436bc1f14610126576100cc565b366100cc576040805162461bcd60e51b815260206004820152600b60248201526a6e6f206465706f7369747360a81b604482015290519081900360640190fd5b005b600080fd5b3480156100dd57600080fd5b50610104600480360360208110156100f457600080fd5b50356001600160a01b031661024d565b60408051918252519081900360200190f35b6101046102bf565b6100ca6102ff565b34801561013257600080fd5b506100ca61033d565b34801561014757600080fd5b506100ca6004803603602081101561015e57600080fd5b503561033f565b34801561017157600080fd5b5061019e6004803603604081101561018857600080fd5b50803590602001356001600160a01b031661038a565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101d85781810151838201526020016101c0565b50505050905090810190601f1680156102055780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561021f57600080fd5b5061010461044e565b6101046004803603602081101561023e57600080fd5b5035610454565b6100ca61033d565b6000816001600160a01b031663267c4ae46040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561028a57600080fd5b505af115801561029e573d6000803e3d6000fd5b505050506040513d60208110156102b457600080fd5b505160010192915050565b600560009081556040805134815290517f1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d24149181900360200190a150600a90565b6040805162461bcd60e51b815260206004820152600e60248201526d1d1a1a5cc81a5cc818481d195cdd60921b604482015290519081900360640190fd5b565b60405130908290600081818185875af1925050503d806000811461037f576040519150601f19603f3d011682016040523d82523d6000602084013e610384565b606091505b50505050565b6060306001600160a01b031663267c4ae46040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156103c757600080fd5b505af11580156103db573d6000803e3d6000fd5b505050506040513d60208110156103f157600080fd5b50506040516060906001600160a01b038416908590600081818185875af1925050503d806000811461043f576040519150601f19603f3d011682016040523d82523d6000602084013e610444565b606091505b5095945050505050565b60015481565b6000806002340460405161046790610530565b660c1e0d0ccc8d4d60ca1b8152604051908190036020019082f0905080158015610495573d6000803e3d6000fd5b5090506040516104a49061053d565b604051809103906000f080156104b75760015b6104c0576104c2565b505b806001600160a01b031663e1cb0e526040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156104fd57600080fd5b505af1158015610511573d6000803e3d6000fd5b505050506040513d602081101561052757600080fd5b50519392505050565b6102b18061054a83390190565b6053806107fb8339019056fe60806040526040516102b13803806102b18339818101604052602081101561002657600080fd5b50516040805163099f12b960e21b81529051339163267c4ae49160048083019260209291908290030181600087803b15801561006157600080fd5b505af1158015610075573d6000803e3d6000fd5b505050506040513d602081101561008b57600080fd5b508190506002340461028e6040516100a290610127565b90815260405183918190036020019083f5915050801580156100c8573d6000803e3d6000fd5b505060408051639b7c9da360e01b81526036600482015290513391639b7c9da391602480830192600092919082900301818387803b15801561010957600080fd5b505af115801561011d573d6000803e3d6000fd5b5050505050610133565b60e7806101ca83390190565b6089806101416000396000f3fe608060405260043610601f5760003560e01c8063e1cb0e5214602a576025565b36602557005b600080fd5b348015603557600080fd5b50603c604e565b60408051918252519081900360200190f35b60149056fea2646970667358221220a4313dc2605287a7a0062ed358ddbd15728ce272ca8a5e4f9dd085d9bc90a43264736f6c634300060c003360806040526040516100e73803806100e78339818101604052602081101561002657600080fd5b505160405133906002340480156108fc02916000818181858888f19350505050158015610057573d6000803e3d6000fd5b50506080806100676000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b60336045565b60408051918252519081900360200190f35b60149056fea2646970667358221220987822a2268d74f17c5d03c1c82a985b32037cf6e762369f81bd42f1f4093a7664736f6c634300060c00336080604052348015600f57600080fd5b506040805162461bcd60e51b8152602060048201526012602482015271125b9d195b9d1a5bdb985b081c995d995c9d60721b604482015290519081900360640190fdfea26469706673582212200ae55ec82222f368bdd1f13b659437a2040175901725345fb4359ecc327f143364736f6c634300060c0033"

// DeploySimple deploys a new Ethereum contract, binding an instance of Simple to it.
func DeploySimple(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simple, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// Simple is an auto generated Go binding around an Ethereum contract.
type Simple struct {
	SimpleCaller     // Read-only binding to the contract
	SimpleTransactor // Write-only binding to the contract
	SimpleFilterer   // Log filterer for contract events
}

// SimpleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleSession struct {
	Contract     *Simple           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCallerSession struct {
	Contract *SimpleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SimpleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleTransactorSession struct {
	Contract     *SimpleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRaw struct {
	Contract *Simple // Generic contract binding to access the raw methods on
}

// SimpleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCallerRaw struct {
	Contract *SimpleCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleTransactorRaw struct {
	Contract *SimpleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimple creates a new instance of Simple, bound to a specific deployed contract.
func NewSimple(address common.Address, backend bind.ContractBackend) (*Simple, error) {
	contract, err := bindSimple(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// NewSimpleCaller creates a new read-only instance of Simple, bound to a specific deployed contract.
func NewSimpleCaller(address common.Address, caller bind.ContractCaller) (*SimpleCaller, error) {
	contract, err := bindSimple(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCaller{contract: contract}, nil
}

// NewSimpleTransactor creates a new write-only instance of Simple, bound to a specific deployed contract.
func NewSimpleTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTransactor, error) {
	contract, err := bindSimple(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransactor{contract: contract}, nil
}

// NewSimpleFilterer creates a new log filterer instance of Simple, bound to a specific deployed contract.
func NewSimpleFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleFilterer, error) {
	contract, err := bindSimple(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleFilterer{contract: contract}, nil
}

// bindSimple binds a generic wrapper to an already deployed contract.
func bindSimple(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.SimpleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transact(opts, method, params...)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_Simple *SimpleCaller) Y(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "y")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_Simple *SimpleSession) Y() (*big.Int, error) {
	return _Simple.Contract.Y(&_Simple.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_Simple *SimpleCallerSession) Y() (*big.Int, error) {
	return _Simple.Contract.Y(&_Simple.CallOpts)
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xae0aba8c.
//
// Solidity: function acceptPayment() payable returns()
func (_Simple *SimpleTransactor) AcceptPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "acceptPayment")
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xae0aba8c.
//
// Solidity: function acceptPayment() payable returns()
func (_Simple *SimpleSession) AcceptPayment() (*types.Transaction, error) {
	return _Simple.Contract.AcceptPayment(&_Simple.TransactOpts)
}

// AcceptPayment is a paid mutator transaction binding the contract method 0xae0aba8c.
//
// Solidity: function acceptPayment() payable returns()
func (_Simple *SimpleTransactorSession) AcceptPayment() (*types.Transaction, error) {
	return _Simple.Contract.AcceptPayment(&_Simple.TransactOpts)
}

// CrossCall is a paid mutator transaction binding the contract method 0x0324332e.
//
// Solidity: function crossCall(address con) returns(uint256)
func (_Simple *SimpleTransactor) CrossCall(opts *bind.TransactOpts, con common.Address) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "crossCall", con)
}

// CrossCall is a paid mutator transaction binding the contract method 0x0324332e.
//
// Solidity: function crossCall(address con) returns(uint256)
func (_Simple *SimpleSession) CrossCall(con common.Address) (*types.Transaction, error) {
	return _Simple.Contract.CrossCall(&_Simple.TransactOpts, con)
}

// CrossCall is a paid mutator transaction binding the contract method 0x0324332e.
//
// Solidity: function crossCall(address con) returns(uint256)
func (_Simple *SimpleTransactorSession) CrossCall(con common.Address) (*types.Transaction, error) {
	return _Simple.Contract.CrossCall(&_Simple.TransactOpts, con)
}

// Exists is a paid mutator transaction binding the contract method 0x267c4ae4.
//
// Solidity: function exists() payable returns(uint256)
func (_Simple *SimpleTransactor) Exists(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "exists")
}

// Exists is a paid mutator transaction binding the contract method 0x267c4ae4.
//
// Solidity: function exists() payable returns(uint256)
func (_Simple *SimpleSession) Exists() (*types.Transaction, error) {
	return _Simple.Contract.Exists(&_Simple.TransactOpts)
}

// Exists is a paid mutator transaction binding the contract method 0x267c4ae4.
//
// Solidity: function exists() payable returns(uint256)
func (_Simple *SimpleTransactorSession) Exists() (*types.Transaction, error) {
	return _Simple.Contract.Exists(&_Simple.TransactOpts)
}

// NestedCall is a paid mutator transaction binding the contract method 0x9b7c9da3.
//
// Solidity: function nestedCall(uint256 value) returns()
func (_Simple *SimpleTransactor) NestedCall(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "nestedCall", value)
}

// NestedCall is a paid mutator transaction binding the contract method 0x9b7c9da3.
//
// Solidity: function nestedCall(uint256 value) returns()
func (_Simple *SimpleSession) NestedCall(value *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.NestedCall(&_Simple.TransactOpts, value)
}

// NestedCall is a paid mutator transaction binding the contract method 0x9b7c9da3.
//
// Solidity: function nestedCall(uint256 value) returns()
func (_Simple *SimpleTransactorSession) NestedCall(value *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.NestedCall(&_Simple.TransactOpts, value)
}

// NestedCall2 is a paid mutator transaction binding the contract method 0xa4ab282d.
//
// Solidity: function nestedCall2(uint256 value, address dest) returns(bytes)
func (_Simple *SimpleTransactor) NestedCall2(opts *bind.TransactOpts, value *big.Int, dest common.Address) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "nestedCall2", value, dest)
}

// NestedCall2 is a paid mutator transaction binding the contract method 0xa4ab282d.
//
// Solidity: function nestedCall2(uint256 value, address dest) returns(bytes)
func (_Simple *SimpleSession) NestedCall2(value *big.Int, dest common.Address) (*types.Transaction, error) {
	return _Simple.Contract.NestedCall2(&_Simple.TransactOpts, value, dest)
}

// NestedCall2 is a paid mutator transaction binding the contract method 0xa4ab282d.
//
// Solidity: function nestedCall2(uint256 value, address dest) returns(bytes)
func (_Simple *SimpleTransactorSession) NestedCall2(value *big.Int, dest common.Address) (*types.Transaction, error) {
	return _Simple.Contract.NestedCall2(&_Simple.TransactOpts, value, dest)
}

// RejectPayment is a paid mutator transaction binding the contract method 0x9436bc1f.
//
// Solidity: function rejectPayment() returns()
func (_Simple *SimpleTransactor) RejectPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "rejectPayment")
}

// RejectPayment is a paid mutator transaction binding the contract method 0x9436bc1f.
//
// Solidity: function rejectPayment() returns()
func (_Simple *SimpleSession) RejectPayment() (*types.Transaction, error) {
	return _Simple.Contract.RejectPayment(&_Simple.TransactOpts)
}

// RejectPayment is a paid mutator transaction binding the contract method 0x9436bc1f.
//
// Solidity: function rejectPayment() returns()
func (_Simple *SimpleTransactorSession) RejectPayment() (*types.Transaction, error) {
	return _Simple.Contract.RejectPayment(&_Simple.TransactOpts)
}

// Reverts is a paid mutator transaction binding the contract method 0x3bccbbc9.
//
// Solidity: function reverts() payable returns()
func (_Simple *SimpleTransactor) Reverts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "reverts")
}

// Reverts is a paid mutator transaction binding the contract method 0x3bccbbc9.
//
// Solidity: function reverts() payable returns()
func (_Simple *SimpleSession) Reverts() (*types.Transaction, error) {
	return _Simple.Contract.Reverts(&_Simple.TransactOpts)
}

// Reverts is a paid mutator transaction binding the contract method 0x3bccbbc9.
//
// Solidity: function reverts() payable returns()
func (_Simple *SimpleTransactorSession) Reverts() (*types.Transaction, error) {
	return _Simple.Contract.Reverts(&_Simple.TransactOpts)
}

// Trace is a paid mutator transaction binding the contract method 0xa68a4fed.
//
// Solidity: function trace(uint256 arg) payable returns(uint256)
func (_Simple *SimpleTransactor) Trace(opts *bind.TransactOpts, arg *big.Int) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "trace", arg)
}

// Trace is a paid mutator transaction binding the contract method 0xa68a4fed.
//
// Solidity: function trace(uint256 arg) payable returns(uint256)
func (_Simple *SimpleSession) Trace(arg *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.Trace(&_Simple.TransactOpts, arg)
}

// Trace is a paid mutator transaction binding the contract method 0xa68a4fed.
//
// Solidity: function trace(uint256 arg) payable returns(uint256)
func (_Simple *SimpleTransactorSession) Trace(arg *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.Trace(&_Simple.TransactOpts, arg)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simple *SimpleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simple *SimpleSession) Receive() (*types.Transaction, error) {
	return _Simple.Contract.Receive(&_Simple.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simple *SimpleTransactorSession) Receive() (*types.Transaction, error) {
	return _Simple.Contract.Receive(&_Simple.TransactOpts)
}

// SimpleTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the Simple contract.
type SimpleTestEventIterator struct {
	Event *SimpleTestEvent // Event containing the contract specifics and raw log

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
func (it *SimpleTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTestEvent)
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
		it.Event = new(SimpleTestEvent)
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
func (it *SimpleTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTestEvent represents a TestEvent event raised by the Simple contract.
type SimpleTestEvent struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 value)
func (_Simple *SimpleFilterer) FilterTestEvent(opts *bind.FilterOpts) (*SimpleTestEventIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &SimpleTestEventIterator{contract: _Simple.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 value)
func (_Simple *SimpleFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *SimpleTestEvent) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTestEvent)
				if err := _Simple.contract.UnpackLog(event, "TestEvent", log); err != nil {
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

// ParseTestEvent is a log parse operation binding the contract event 0x1440c4dd67b4344ea1905ec0318995133b550f168b4ee959a0da6b503d7d2414.
//
// Solidity: event TestEvent(uint256 value)
func (_Simple *SimpleFilterer) ParseTestEvent(log types.Log) (*SimpleTestEvent, error) {
	event := new(SimpleTestEvent)
	if err := _Simple.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
