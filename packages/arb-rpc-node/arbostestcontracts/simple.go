// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbostestcontracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ComplexConstructorConMetaData contains all meta data concerning the ComplexConstructorCon contract.
var ComplexConstructorConMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getVal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526040516103273803806103278339810160408190526100229161013d565b336001600160a01b031663267c4ae46040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561005d57600080fd5b505af1158015610071573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610095919061013d565b50806100a2600234610156565b61028e6040516100b190610130565b90815260200182906040518091039083f5915050801580156100d7573d6000803e3d6000fd5b5050604051639b7c9da360e01b8152603660048201523390639b7c9da390602401600060405180830381600087803b15801561011257600080fd5b505af1158015610126573d6000803e3d6000fd5b5050505050610178565b6101208061020783390190565b60006020828403121561014f57600080fd5b5051919050565b60008261017357634e487b7160e01b600052601260045260246000fd5b500490565b6081806101866000396000f3fe60806040526004361060205760003560e01c8063e1cb0e5214602b57600080fd5b36602657005b600080fd5b348015603657600080fd5b50601460405190815260200160405180910390f3fea264697066735822122052367f231416078848403787be9cfad07982e47c5dfe9f5209c7dc1b9d63011764736f6c63430008070033608060405260405161012038038061012083398101604081905261002291610060565b336108fc610031600234610079565b6040518115909202916000818181858888f19350505050158015610059573d6000803e3d6000fd5b505061009b565b60006020828403121561007257600080fd5b5051919050565b60008261009657634e487b7160e01b600052601260045260246000fd5b500490565b6077806100a96000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b601460405190815260200160405180910390f3fea2646970667358221220024516a865eb97a54cf216467a7132a0e365122ef447c7170573528168d5b44b64736f6c63430008070033",
}

// ComplexConstructorConABI is the input ABI used to generate the binding from.
// Deprecated: Use ComplexConstructorConMetaData.ABI instead.
var ComplexConstructorConABI = ComplexConstructorConMetaData.ABI

// ComplexConstructorConBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ComplexConstructorConMetaData.Bin instead.
var ComplexConstructorConBin = ComplexConstructorConMetaData.Bin

// DeployComplexConstructorCon deploys a new Ethereum contract, binding an instance of ComplexConstructorCon to it.
func DeployComplexConstructorCon(auth *bind.TransactOpts, backend bind.ContractBackend, salt [32]byte) (common.Address, *types.Transaction, *ComplexConstructorCon, error) {
	parsed, err := ComplexConstructorConMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ComplexConstructorConBin), backend, salt)
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

// ComplexConstructorCon2MetaData contains all meta data concerning the ComplexConstructorCon2 contract.
var ComplexConstructorCon2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getVal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260405161012038038061012083398101604081905261002291610060565b336108fc610031600234610079565b6040518115909202916000818181858888f19350505050158015610059573d6000803e3d6000fd5b505061009b565b60006020828403121561007257600080fd5b5051919050565b60008261009657634e487b7160e01b600052601260045260246000fd5b500490565b6077806100a96000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b601460405190815260200160405180910390f3fea2646970667358221220024516a865eb97a54cf216467a7132a0e365122ef447c7170573528168d5b44b64736f6c63430008070033",
}

// ComplexConstructorCon2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ComplexConstructorCon2MetaData.ABI instead.
var ComplexConstructorCon2ABI = ComplexConstructorCon2MetaData.ABI

// ComplexConstructorCon2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ComplexConstructorCon2MetaData.Bin instead.
var ComplexConstructorCon2Bin = ComplexConstructorCon2MetaData.Bin

// DeployComplexConstructorCon2 deploys a new Ethereum contract, binding an instance of ComplexConstructorCon2 to it.
func DeployComplexConstructorCon2(auth *bind.TransactOpts, backend bind.ContractBackend, val *big.Int) (common.Address, *types.Transaction, *ComplexConstructorCon2, error) {
	parsed, err := ComplexConstructorCon2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ComplexConstructorCon2Bin), backend, val)
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

// Destroyer1MetaData contains all meta data concerning the Destroyer1 contract.
var Destroyer1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5033fffe",
}

// Destroyer1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Destroyer1MetaData.ABI instead.
var Destroyer1ABI = Destroyer1MetaData.ABI

// Destroyer1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Destroyer1MetaData.Bin instead.
var Destroyer1Bin = Destroyer1MetaData.Bin

// DeployDestroyer1 deploys a new Ethereum contract, binding an instance of Destroyer1 to it.
func DeployDestroyer1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Destroyer1, error) {
	parsed, err := Destroyer1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Destroyer1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Destroyer1{Destroyer1Caller: Destroyer1Caller{contract: contract}, Destroyer1Transactor: Destroyer1Transactor{contract: contract}, Destroyer1Filterer: Destroyer1Filterer{contract: contract}}, nil
}

// Destroyer1 is an auto generated Go binding around an Ethereum contract.
type Destroyer1 struct {
	Destroyer1Caller     // Read-only binding to the contract
	Destroyer1Transactor // Write-only binding to the contract
	Destroyer1Filterer   // Log filterer for contract events
}

// Destroyer1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Destroyer1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Destroyer1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Destroyer1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Destroyer1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Destroyer1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Destroyer1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Destroyer1Session struct {
	Contract     *Destroyer1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Destroyer1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Destroyer1CallerSession struct {
	Contract *Destroyer1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Destroyer1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Destroyer1TransactorSession struct {
	Contract     *Destroyer1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Destroyer1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Destroyer1Raw struct {
	Contract *Destroyer1 // Generic contract binding to access the raw methods on
}

// Destroyer1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Destroyer1CallerRaw struct {
	Contract *Destroyer1Caller // Generic read-only contract binding to access the raw methods on
}

// Destroyer1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Destroyer1TransactorRaw struct {
	Contract *Destroyer1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDestroyer1 creates a new instance of Destroyer1, bound to a specific deployed contract.
func NewDestroyer1(address common.Address, backend bind.ContractBackend) (*Destroyer1, error) {
	contract, err := bindDestroyer1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destroyer1{Destroyer1Caller: Destroyer1Caller{contract: contract}, Destroyer1Transactor: Destroyer1Transactor{contract: contract}, Destroyer1Filterer: Destroyer1Filterer{contract: contract}}, nil
}

// NewDestroyer1Caller creates a new read-only instance of Destroyer1, bound to a specific deployed contract.
func NewDestroyer1Caller(address common.Address, caller bind.ContractCaller) (*Destroyer1Caller, error) {
	contract, err := bindDestroyer1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Destroyer1Caller{contract: contract}, nil
}

// NewDestroyer1Transactor creates a new write-only instance of Destroyer1, bound to a specific deployed contract.
func NewDestroyer1Transactor(address common.Address, transactor bind.ContractTransactor) (*Destroyer1Transactor, error) {
	contract, err := bindDestroyer1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Destroyer1Transactor{contract: contract}, nil
}

// NewDestroyer1Filterer creates a new log filterer instance of Destroyer1, bound to a specific deployed contract.
func NewDestroyer1Filterer(address common.Address, filterer bind.ContractFilterer) (*Destroyer1Filterer, error) {
	contract, err := bindDestroyer1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Destroyer1Filterer{contract: contract}, nil
}

// bindDestroyer1 binds a generic wrapper to an already deployed contract.
func bindDestroyer1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Destroyer1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destroyer1 *Destroyer1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destroyer1.Contract.Destroyer1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destroyer1 *Destroyer1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyer1.Contract.Destroyer1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destroyer1 *Destroyer1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destroyer1.Contract.Destroyer1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destroyer1 *Destroyer1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destroyer1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destroyer1 *Destroyer1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyer1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destroyer1 *Destroyer1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destroyer1.Contract.contract.Transact(opts, method, params...)
}

// Destroyer2MetaData contains all meta data concerning the Destroyer2 contract.
var Destroyer2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"test2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"test3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"test4\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102f5806100206000396000f3fe60806040526004361061004a5760003560e01c8063658fd1041461004f5780636b59084d1461007457806383197ef01461007b578063b11bead414610090578063c5da4876146100a3575b600080fd5b61006261005d36600461023b565b6100b6565b60405190815260200160405180910390f35b600a610062565b34801561008757600080fd5b5061008e33ff5b005b61006261009e36600461023b565b610131565b6100626100b136600461023b565b6101e5565b6000816001600160a01b0316636b59084d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156100f357600080fd5b505af1158015610107573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061012b919061026b565b92915050565b60408051600481526024810182526020810180516001600160e01b0316636b59084d60e01b1790529051600091829182916001600160a01b038616916101779190610284565b600060405180830381855af49150503d80600081146101b2576040519150601f19603f3d011682016040523d82523d6000602084013e6101b7565b606091505b5091509150816101c657600080fd5b6000818060200190518101906101dc919061026b565b95945050505050565b6040805160048152602481018252602080820180516001600160e01b0316636b59084d60e01b178152825193516000949285929190818584868b5af292508051955050508061023357600080fd5b505050919050565b60006020828403121561024d57600080fd5b81356001600160a01b038116811461026457600080fd5b9392505050565b60006020828403121561027d57600080fd5b5051919050565b6000825160005b818110156102a5576020818601810151858301520161028b565b818111156102b4576000828501525b50919091019291505056fea2646970667358221220f710687d3c62d36f16d60229dc62f04e60ffb7ca7202ae528d79037d77642d2764736f6c63430008070033",
}

// Destroyer2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Destroyer2MetaData.ABI instead.
var Destroyer2ABI = Destroyer2MetaData.ABI

// Destroyer2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Destroyer2MetaData.Bin instead.
var Destroyer2Bin = Destroyer2MetaData.Bin

// DeployDestroyer2 deploys a new Ethereum contract, binding an instance of Destroyer2 to it.
func DeployDestroyer2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Destroyer2, error) {
	parsed, err := Destroyer2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Destroyer2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Destroyer2{Destroyer2Caller: Destroyer2Caller{contract: contract}, Destroyer2Transactor: Destroyer2Transactor{contract: contract}, Destroyer2Filterer: Destroyer2Filterer{contract: contract}}, nil
}

// Destroyer2 is an auto generated Go binding around an Ethereum contract.
type Destroyer2 struct {
	Destroyer2Caller     // Read-only binding to the contract
	Destroyer2Transactor // Write-only binding to the contract
	Destroyer2Filterer   // Log filterer for contract events
}

// Destroyer2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Destroyer2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Destroyer2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Destroyer2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Destroyer2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Destroyer2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Destroyer2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Destroyer2Session struct {
	Contract     *Destroyer2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Destroyer2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Destroyer2CallerSession struct {
	Contract *Destroyer2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Destroyer2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Destroyer2TransactorSession struct {
	Contract     *Destroyer2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Destroyer2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Destroyer2Raw struct {
	Contract *Destroyer2 // Generic contract binding to access the raw methods on
}

// Destroyer2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Destroyer2CallerRaw struct {
	Contract *Destroyer2Caller // Generic read-only contract binding to access the raw methods on
}

// Destroyer2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Destroyer2TransactorRaw struct {
	Contract *Destroyer2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDestroyer2 creates a new instance of Destroyer2, bound to a specific deployed contract.
func NewDestroyer2(address common.Address, backend bind.ContractBackend) (*Destroyer2, error) {
	contract, err := bindDestroyer2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destroyer2{Destroyer2Caller: Destroyer2Caller{contract: contract}, Destroyer2Transactor: Destroyer2Transactor{contract: contract}, Destroyer2Filterer: Destroyer2Filterer{contract: contract}}, nil
}

// NewDestroyer2Caller creates a new read-only instance of Destroyer2, bound to a specific deployed contract.
func NewDestroyer2Caller(address common.Address, caller bind.ContractCaller) (*Destroyer2Caller, error) {
	contract, err := bindDestroyer2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Destroyer2Caller{contract: contract}, nil
}

// NewDestroyer2Transactor creates a new write-only instance of Destroyer2, bound to a specific deployed contract.
func NewDestroyer2Transactor(address common.Address, transactor bind.ContractTransactor) (*Destroyer2Transactor, error) {
	contract, err := bindDestroyer2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Destroyer2Transactor{contract: contract}, nil
}

// NewDestroyer2Filterer creates a new log filterer instance of Destroyer2, bound to a specific deployed contract.
func NewDestroyer2Filterer(address common.Address, filterer bind.ContractFilterer) (*Destroyer2Filterer, error) {
	contract, err := bindDestroyer2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Destroyer2Filterer{contract: contract}, nil
}

// bindDestroyer2 binds a generic wrapper to an already deployed contract.
func bindDestroyer2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Destroyer2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destroyer2 *Destroyer2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destroyer2.Contract.Destroyer2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destroyer2 *Destroyer2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyer2.Contract.Destroyer2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destroyer2 *Destroyer2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destroyer2.Contract.Destroyer2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destroyer2 *Destroyer2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destroyer2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destroyer2 *Destroyer2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyer2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destroyer2 *Destroyer2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destroyer2.Contract.contract.Transact(opts, method, params...)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destroyer2 *Destroyer2Transactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyer2.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destroyer2 *Destroyer2Session) Destroy() (*types.Transaction, error) {
	return _Destroyer2.Contract.Destroy(&_Destroyer2.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destroyer2 *Destroyer2TransactorSession) Destroy() (*types.Transaction, error) {
	return _Destroyer2.Contract.Destroy(&_Destroyer2.TransactOpts)
}

// Test1 is a paid mutator transaction binding the contract method 0x6b59084d.
//
// Solidity: function test1() payable returns(uint256)
func (_Destroyer2 *Destroyer2Transactor) Test1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyer2.contract.Transact(opts, "test1")
}

// Test1 is a paid mutator transaction binding the contract method 0x6b59084d.
//
// Solidity: function test1() payable returns(uint256)
func (_Destroyer2 *Destroyer2Session) Test1() (*types.Transaction, error) {
	return _Destroyer2.Contract.Test1(&_Destroyer2.TransactOpts)
}

// Test1 is a paid mutator transaction binding the contract method 0x6b59084d.
//
// Solidity: function test1() payable returns(uint256)
func (_Destroyer2 *Destroyer2TransactorSession) Test1() (*types.Transaction, error) {
	return _Destroyer2.Contract.Test1(&_Destroyer2.TransactOpts)
}

// Test2 is a paid mutator transaction binding the contract method 0xb11bead4.
//
// Solidity: function test2(address to) payable returns(uint256)
func (_Destroyer2 *Destroyer2Transactor) Test2(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Destroyer2.contract.Transact(opts, "test2", to)
}

// Test2 is a paid mutator transaction binding the contract method 0xb11bead4.
//
// Solidity: function test2(address to) payable returns(uint256)
func (_Destroyer2 *Destroyer2Session) Test2(to common.Address) (*types.Transaction, error) {
	return _Destroyer2.Contract.Test2(&_Destroyer2.TransactOpts, to)
}

// Test2 is a paid mutator transaction binding the contract method 0xb11bead4.
//
// Solidity: function test2(address to) payable returns(uint256)
func (_Destroyer2 *Destroyer2TransactorSession) Test2(to common.Address) (*types.Transaction, error) {
	return _Destroyer2.Contract.Test2(&_Destroyer2.TransactOpts, to)
}

// Test3 is a paid mutator transaction binding the contract method 0xc5da4876.
//
// Solidity: function test3(address to) payable returns(uint256 c)
func (_Destroyer2 *Destroyer2Transactor) Test3(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Destroyer2.contract.Transact(opts, "test3", to)
}

// Test3 is a paid mutator transaction binding the contract method 0xc5da4876.
//
// Solidity: function test3(address to) payable returns(uint256 c)
func (_Destroyer2 *Destroyer2Session) Test3(to common.Address) (*types.Transaction, error) {
	return _Destroyer2.Contract.Test3(&_Destroyer2.TransactOpts, to)
}

// Test3 is a paid mutator transaction binding the contract method 0xc5da4876.
//
// Solidity: function test3(address to) payable returns(uint256 c)
func (_Destroyer2 *Destroyer2TransactorSession) Test3(to common.Address) (*types.Transaction, error) {
	return _Destroyer2.Contract.Test3(&_Destroyer2.TransactOpts, to)
}

// Test4 is a paid mutator transaction binding the contract method 0x658fd104.
//
// Solidity: function test4(address to) payable returns(uint256)
func (_Destroyer2 *Destroyer2Transactor) Test4(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Destroyer2.contract.Transact(opts, "test4", to)
}

// Test4 is a paid mutator transaction binding the contract method 0x658fd104.
//
// Solidity: function test4(address to) payable returns(uint256)
func (_Destroyer2 *Destroyer2Session) Test4(to common.Address) (*types.Transaction, error) {
	return _Destroyer2.Contract.Test4(&_Destroyer2.TransactOpts, to)
}

// Test4 is a paid mutator transaction binding the contract method 0x658fd104.
//
// Solidity: function test4(address to) payable returns(uint256)
func (_Destroyer2 *Destroyer2TransactorSession) Test4(to common.Address) (*types.Transaction, error) {
	return _Destroyer2.Contract.Test4(&_Destroyer2.TransactOpts, to)
}

// ReverterMetaData contains all meta data concerning the Reverter contract.
var ReverterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5060405162461bcd60e51b8152602060048201526012602482015271125b9d195b9d1a5bdb985b081c995d995c9d60721b604482015260640160405180910390fdfe",
}

// ReverterABI is the input ABI used to generate the binding from.
// Deprecated: Use ReverterMetaData.ABI instead.
var ReverterABI = ReverterMetaData.ABI

// ReverterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReverterMetaData.Bin instead.
var ReverterBin = ReverterMetaData.Bin

// DeployReverter deploys a new Ethereum contract, binding an instance of Reverter to it.
func DeployReverter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reverter, error) {
	parsed, err := ReverterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReverterBin), backend)
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

// SimpleMetaData contains all meta data concerning the Simple contract.
var SimpleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool[10]\",\"name\":\"_variable\",\"type\":\"bool[10]\"}],\"name\":\"Variable\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arrayPush\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractSimple\",\"name\":\"con\",\"type\":\"address\"}],\"name\":\"crossCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debug\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"nestedCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"nestedCall2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rejectPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reverts\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg\",\"type\":\"uint256\"}],\"name\":\"trace\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604081815234600181905582523360a0527f9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc91a1611132806100446000396000f3fe6080604052600436106100a05760003560e01c80639436bc1f116100645780639436bc1f146101475780639b7c9da314610156578063a4ab282d14610176578063a56dfe4a146101a3578063a68a4fed146101b9578063ae0aba8c146100e057600080fd5b80630324332e146100e7578063267c4ae41461011a5780633bccbbc91461012257806347f6ac111461012a578063588ee29b1461013257600080fd5b366100e25760405162461bcd60e51b815260206004820152600b60248201526a6e6f206465706f7369747360a81b60448201526064015b60405180910390fd5b005b600080fd5b3480156100f357600080fd5b506101076101023660046108c5565b6101cc565b6040519081526020015b60405180910390f35b610107610252565b6100e0610298565b6101076102d1565b34801561013e57600080fd5b506100e0610330565b34801561015357600080fd5b50005b34801561016257600080fd5b506100e06101713660046108e9565b610372565b34801561018257600080fd5b5061019661019136600461091b565b6103bd565b604051610111919061099b565b3480156101af57600080fd5b5061010760015481565b6101076101c73660046108e9565b61048f565b6000816001600160a01b031663267c4ae46040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561020957600080fd5b505af115801561021d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102419190610902565b61024c9060016109ce565b92915050565b60056000908155604080513481523360208201527f9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc91015b60405180910390a150600a90565b60405162461bcd60e51b815260206004820152600e60248201526d1d1a1a5cc81a5cc818481d195cdd60921b60448201526064016100d7565b6000600260005460016102e491906109ce565b8154600181018355600092835260209283902001556040805134815233928101929092527f9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc910161028a565b610338610874565b7fd09ca446fa2fa15b4f798d40263d8fb2d571205ba080eba1839db8b1efd5f676816040516103679190610967565b60405180910390a150565b60405130908290600081818185875af1925050503d80600081146103b2576040519150601f19603f3d011682016040523d82523d6000602084013e6103b7565b606091505b50505050565b6060306001600160a01b031663267c4ae46040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156103fa57600080fd5b505af115801561040e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104329190610902565b506000826001600160a01b03168460405160006040518083038185875af1925050503d8060008114610480576040519150601f19603f3d011682016040523d82523d6000602084013e610485565b606091505b5095945050505050565b60008061049d6002346109f4565b6040516104a990610893565b660c1e0d0ccc8d4d60ca1b81526020016040518091039082f09050801580156104d6573d6000803e3d6000fd5b5090506040516104e5906108a0565b604051809103906000f080156104f85760015b61050157610503565b505b60405161050f906108ac565b604051809103906000f08015801561052b573d6000803e3d6000fd5b5050600060405161053b906108b8565b604051809103906000f080158015610557573d6000803e3d6000fd5b5090506000604051610568906108b8565b604051809103906000f080158015610584573d6000803e3d6000fd5b50604080516001600160a01b038316602480830182905283518084038201815260449384018552602080820180516001600160e01b03908116632c46fab560e21b1782528751808601969096528751808703909501855294909501865282810180519094166362ed243b60e11b17909352805194519596509490939287929091908185846000875af2506044016040819052602085810192508185846000875af25060440160408190526001600160a01b038316915061064590869061094b565b600060405180830381855af49150503d8060008114610680576040519150601f19603f3d011682016040523d82523d6000602084013e610685565b606091505b505050806001600160a01b0316836040516106a0919061094b565b600060405180830381855af49150503d80600081146106db576040519150601f19603f3d011682016040523d82523d6000602084013e6106e0565b606091505b5050604080516001600160a01b0388811660248084019190915283518084039091018152604490920183526020820180516001600160e01b0316631963f44160e21b179052915191841692506107359161094b565b600060405180830381855af49150503d8060008114610770576040519150601f19603f3d011682016040523d82523d6000602084013e610775565b606091505b5050604051632c46fab560e21b81526001600160a01b0387811660048301528816915063b11bead490602401602060405180830381600087803b1580156107bb57600080fd5b505af11580156107cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f39190610902565b50866001600160a01b031663e1cb0e526040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561082f57600080fd5b505af1158015610843573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108679190610902565b9998505050505050505050565b604051806101400160405280600a906020820280368337509192915050565b61032780610a5b83390190565b605280610d8283390190565b601480610dd483390190565b61031580610de883390190565b6000602082840312156108d757600080fd5b81356108e281610a42565b9392505050565b6000602082840312156108fb57600080fd5b5035919050565b60006020828403121561091457600080fd5b5051919050565b6000806040838503121561092e57600080fd5b82359150602083013561094081610a42565b809150509250929050565b6000825161095d818460208701610a16565b9190910192915050565b6101408101818360005b600a8110156109925781511515835260209283019290910190600101610971565b50505092915050565b60208152600082518060208401526109ba816040850160208701610a16565b601f01601f19169190910160400192915050565b600082198211156109ef57634e487b7160e01b600052601160045260246000fd5b500190565b600082610a1157634e487b7160e01b600052601260045260246000fd5b500490565b60005b83811015610a31578181015183820152602001610a19565b838111156103b75750506000910152565b6001600160a01b0381168114610a5757600080fd5b5056fe60806040526040516103273803806103278339810160408190526100229161013d565b336001600160a01b031663267c4ae46040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561005d57600080fd5b505af1158015610071573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610095919061013d565b50806100a2600234610156565b61028e6040516100b190610130565b90815260200182906040518091039083f5915050801580156100d7573d6000803e3d6000fd5b5050604051639b7c9da360e01b8152603660048201523390639b7c9da390602401600060405180830381600087803b15801561011257600080fd5b505af1158015610126573d6000803e3d6000fd5b5050505050610178565b6101208061020783390190565b60006020828403121561014f57600080fd5b5051919050565b60008261017357634e487b7160e01b600052601260045260246000fd5b500490565b6081806101866000396000f3fe60806040526004361060205760003560e01c8063e1cb0e5214602b57600080fd5b36602657005b600080fd5b348015603657600080fd5b50601460405190815260200160405180910390f3fea264697066735822122052367f231416078848403787be9cfad07982e47c5dfe9f5209c7dc1b9d63011764736f6c63430008070033608060405260405161012038038061012083398101604081905261002291610060565b336108fc610031600234610079565b6040518115909202916000818181858888f19350505050158015610059573d6000803e3d6000fd5b505061009b565b60006020828403121561007257600080fd5b5051919050565b60008261009657634e487b7160e01b600052601260045260246000fd5b500490565b6077806100a96000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b601460405190815260200160405180910390f3fea2646970667358221220024516a865eb97a54cf216467a7132a0e365122ef447c7170573528168d5b44b64736f6c634300080700336080604052348015600f57600080fd5b5060405162461bcd60e51b8152602060048201526012602482015271125b9d195b9d1a5bdb985b081c995d995c9d60721b604482015260640160405180910390fdfe6080604052348015600f57600080fd5b5033fffe608060405234801561001057600080fd5b506102f5806100206000396000f3fe60806040526004361061004a5760003560e01c8063658fd1041461004f5780636b59084d1461007457806383197ef01461007b578063b11bead414610090578063c5da4876146100a3575b600080fd5b61006261005d36600461023b565b6100b6565b60405190815260200160405180910390f35b600a610062565b34801561008757600080fd5b5061008e33ff5b005b61006261009e36600461023b565b610131565b6100626100b136600461023b565b6101e5565b6000816001600160a01b0316636b59084d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156100f357600080fd5b505af1158015610107573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061012b919061026b565b92915050565b60408051600481526024810182526020810180516001600160e01b0316636b59084d60e01b1790529051600091829182916001600160a01b038616916101779190610284565b600060405180830381855af49150503d80600081146101b2576040519150601f19603f3d011682016040523d82523d6000602084013e6101b7565b606091505b5091509150816101c657600080fd5b6000818060200190518101906101dc919061026b565b95945050505050565b6040805160048152602481018252602080820180516001600160e01b0316636b59084d60e01b178152825193516000949285929190818584868b5af292508051955050508061023357600080fd5b505050919050565b60006020828403121561024d57600080fd5b81356001600160a01b038116811461026457600080fd5b9392505050565b60006020828403121561027d57600080fd5b5051919050565b6000825160005b818110156102a5576020818601810151858301520161028b565b818111156102b4576000828501525b50919091019291505056fea2646970667358221220f710687d3c62d36f16d60229dc62f04e60ffb7ca7202ae528d79037d77642d2764736f6c63430008070033a2646970667358221220360464dfaa2777ea5bcb298dfca3bf4bdfd2d6f1fdba63fa1fe62ca6be5fcaeb64736f6c63430008070033",
}

// SimpleABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleMetaData.ABI instead.
var SimpleABI = SimpleMetaData.ABI

// SimpleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleMetaData.Bin instead.
var SimpleBin = SimpleMetaData.Bin

// DeploySimple deploys a new Ethereum contract, binding an instance of Simple to it.
func DeploySimple(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simple, error) {
	parsed, err := SimpleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleBin), backend)
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

// ArrayPush is a paid mutator transaction binding the contract method 0x47f6ac11.
//
// Solidity: function arrayPush() payable returns(uint256)
func (_Simple *SimpleTransactor) ArrayPush(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "arrayPush")
}

// ArrayPush is a paid mutator transaction binding the contract method 0x47f6ac11.
//
// Solidity: function arrayPush() payable returns(uint256)
func (_Simple *SimpleSession) ArrayPush() (*types.Transaction, error) {
	return _Simple.Contract.ArrayPush(&_Simple.TransactOpts)
}

// ArrayPush is a paid mutator transaction binding the contract method 0x47f6ac11.
//
// Solidity: function arrayPush() payable returns(uint256)
func (_Simple *SimpleTransactorSession) ArrayPush() (*types.Transaction, error) {
	return _Simple.Contract.ArrayPush(&_Simple.TransactOpts)
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

// Debug is a paid mutator transaction binding the contract method 0x588ee29b.
//
// Solidity: function debug() returns()
func (_Simple *SimpleTransactor) Debug(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "debug")
}

// Debug is a paid mutator transaction binding the contract method 0x588ee29b.
//
// Solidity: function debug() returns()
func (_Simple *SimpleSession) Debug() (*types.Transaction, error) {
	return _Simple.Contract.Debug(&_Simple.TransactOpts)
}

// Debug is a paid mutator transaction binding the contract method 0x588ee29b.
//
// Solidity: function debug() returns()
func (_Simple *SimpleTransactorSession) Debug() (*types.Transaction, error) {
	return _Simple.Contract.Debug(&_Simple.TransactOpts)
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
	Value  *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0x9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc.
//
// Solidity: event TestEvent(uint256 value, address sender)
func (_Simple *SimpleFilterer) FilterTestEvent(opts *bind.FilterOpts) (*SimpleTestEventIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &SimpleTestEventIterator{contract: _Simple.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0x9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc.
//
// Solidity: event TestEvent(uint256 value, address sender)
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

// ParseTestEvent is a log parse operation binding the contract event 0x9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc.
//
// Solidity: event TestEvent(uint256 value, address sender)
func (_Simple *SimpleFilterer) ParseTestEvent(log types.Log) (*SimpleTestEvent, error) {
	event := new(SimpleTestEvent)
	if err := _Simple.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleVariableIterator is returned from FilterVariable and is used to iterate over the raw logs and unpacked data for Variable events raised by the Simple contract.
type SimpleVariableIterator struct {
	Event *SimpleVariable // Event containing the contract specifics and raw log

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
func (it *SimpleVariableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleVariable)
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
		it.Event = new(SimpleVariable)
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
func (it *SimpleVariableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleVariableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleVariable represents a Variable event raised by the Simple contract.
type SimpleVariable struct {
	Variable [10]bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVariable is a free log retrieval operation binding the contract event 0xd09ca446fa2fa15b4f798d40263d8fb2d571205ba080eba1839db8b1efd5f676.
//
// Solidity: event Variable(bool[10] _variable)
func (_Simple *SimpleFilterer) FilterVariable(opts *bind.FilterOpts) (*SimpleVariableIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "Variable")
	if err != nil {
		return nil, err
	}
	return &SimpleVariableIterator{contract: _Simple.contract, event: "Variable", logs: logs, sub: sub}, nil
}

// WatchVariable is a free log subscription operation binding the contract event 0xd09ca446fa2fa15b4f798d40263d8fb2d571205ba080eba1839db8b1efd5f676.
//
// Solidity: event Variable(bool[10] _variable)
func (_Simple *SimpleFilterer) WatchVariable(opts *bind.WatchOpts, sink chan<- *SimpleVariable) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "Variable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleVariable)
				if err := _Simple.contract.UnpackLog(event, "Variable", log); err != nil {
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

// ParseVariable is a log parse operation binding the contract event 0xd09ca446fa2fa15b4f798d40263d8fb2d571205ba080eba1839db8b1efd5f676.
//
// Solidity: event Variable(bool[10] _variable)
func (_Simple *SimpleFilterer) ParseVariable(log types.Log) (*SimpleVariable, error) {
	event := new(SimpleVariable)
	if err := _Simple.contract.UnpackLog(event, "Variable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
