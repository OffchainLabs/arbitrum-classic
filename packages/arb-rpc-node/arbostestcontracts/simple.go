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
	Bin: "0x60806040526040516102b13803806102b18339818101604052602081101561002657600080fd5b50516040805163099f12b960e21b81529051339163267c4ae49160048083019260209291908290030181600087803b15801561006157600080fd5b505af1158015610075573d6000803e3d6000fd5b505050506040513d602081101561008b57600080fd5b508190506002340461028e6040516100a290610127565b90815260405183918190036020019083f5915050801580156100c8573d6000803e3d6000fd5b505060408051639b7c9da360e01b81526036600482015290513391639b7c9da391602480830192600092919082900301818387803b15801561010957600080fd5b505af115801561011d573d6000803e3d6000fd5b5050505050610133565b60e7806101ca83390190565b6089806101416000396000f3fe608060405260043610601f5760003560e01c8063e1cb0e5214602a576025565b36602557005b600080fd5b348015603557600080fd5b50603c604e565b60408051918252519081900360200190f35b60149056fea2646970667358221220884dc0aad81e4fbdfb03eea198fb27577f71d8f82244e896729d8603bf285abf64736f6c6343000706003360806040526040516100e73803806100e78339818101604052602081101561002657600080fd5b505160405133906002340480156108fc02916000818181858888f19350505050158015610057573d6000803e3d6000fd5b50506080806100676000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b60336045565b60408051918252519081900360200190f35b60149056fea2646970667358221220b58936794eb5834de57f1e8a0a4e598f105dfe7dbf0092c89c1732bab85aa57464736f6c63430007060033",
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
	Bin: "0x60806040526040516100e73803806100e78339818101604052602081101561002657600080fd5b505160405133906002340480156108fc02916000818181858888f19350505050158015610057573d6000803e3d6000fd5b50506080806100676000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b60336045565b60408051918252519081900360200190f35b60149056fea2646970667358221220b58936794eb5834de57f1e8a0a4e598f105dfe7dbf0092c89c1732bab85aa57464736f6c63430007060033",
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
	Bin: "0x608060405234801561001057600080fd5b506102f3806100206000396000f3fe60806040526004361061004a5760003560e01c8063658fd1041461004f5780636b59084d1461008757806383197ef01461008f578063b11bead4146100a6578063c5da4876146100cc575b600080fd5b6100756004803603602081101561006557600080fd5b50356001600160a01b03166100f2565b60408051918252519081900360200190f35b610075610161565b34801561009b57600080fd5b506100a4610166565b005b610075600480360360208110156100bc57600080fd5b50356001600160a01b0316610169565b610075600480360360208110156100e257600080fd5b50356001600160a01b0316610267565b6000816001600160a01b0316636b59084d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561012f57600080fd5b505af1158015610143573d6000803e3d6000fd5b505050506040513d602081101561015957600080fd5b505192915050565b600a90565b33ff5b60408051600481526024810182526020810180516001600160e01b0316636b59084d60e01b17815291518151600093849384936001600160a01b03881693919290918291908083835b602083106101d15780518252601f1990920191602091820191016101b2565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610231576040519150601f19603f3d011682016040523d82523d6000602084013e610236565b606091505b50915091508161024557600080fd5b600081806020019051602081101561025c57600080fd5b505195945050505050565b6040805160048152602481018252602080820180516001600160e01b0316636b59084d60e01b178152825193516000949285929190818584868b5af29250805195505050806102b557600080fd5b50505091905056fea2646970667358221220b79a409fb525793c6a8cc2d5ce9cb15ffc47b826e110cabff084a66dbab52cb864736f6c63430007060033",
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
	Bin: "0x6080604052348015600f57600080fd5b506040805162461bcd60e51b8152602060048201526012602482015271125b9d195b9d1a5bdb985b081c995d995c9d60721b604482015290519081900360640190fdfe",
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
	Bin: "0x6080604081815234600181905582523360a0527f9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc91a16110cd806100446000396000f3fe6080604052600436106100a05760003560e01c80639436bc1f116100645780639436bc1f146101595780639b7c9da31461016e578063a4ab282d14610198578063a56dfe4a14610246578063a68a4fed1461025b578063ae0aba8c14610278576100e2565b80630324332e146100e7578063267c4ae41461012c5780633bccbbc91461013457806347f6ac111461013c578063588ee29b14610144576100e2565b366100e2576040805162461bcd60e51b815260206004820152600b60248201526a6e6f206465706f7369747360a81b604482015290519081900360640190fd5b005b600080fd5b3480156100f357600080fd5b5061011a6004803603602081101561010a57600080fd5b50356001600160a01b0316610280565b60408051918252519081900360200190f35b61011a6102f2565b6100e0610339565b61011a610379565b34801561015057600080fd5b506100e06103f2565b34801561016557600080fd5b506100e0610377565b34801561017a57600080fd5b506100e06004803603602081101561019157600080fd5b5035610459565b3480156101a457600080fd5b506101d1600480360360408110156101bb57600080fd5b50803590602001356001600160a01b03166104a4565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561020b5781810151838201526020016101f3565b50505050905090810190601f1680156102385780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561025257600080fd5b5061011a610567565b61011a6004803603602081101561027157600080fd5b503561056d565b6100e0610377565b6000816001600160a01b031663267c4ae46040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156102bd57600080fd5b505af11580156102d1573d6000803e3d6000fd5b505050506040513d60208110156102e757600080fd5b505160010192915050565b600560009081556040805134815233602082015281517f9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc929181900390910190a150600a90565b6040805162461bcd60e51b815260206004820152600e60248201526d1d1a1a5cc81a5cc818481d195cdd60921b604482015290519081900360640190fd5b565b60008054600280546001818101835591845291017f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace909101556040805134815233602082015281517f9457b0abc6a87108b750271d78f46ad30369fbeb6a7454888743813252fca3fc929181900390910190a150600a90565b6103fa610a1b565b7fd09ca446fa2fa15b4f798d40263d8fb2d571205ba080eba1839db8b1efd5f676816040518082600a60200280838360005b8381101561044457818101518382015260200161042c565b5050505090500191505060405180910390a150565b60405130908290600081818185875af1925050503d8060008114610499576040519150601f19603f3d011682016040523d82523d6000602084013e61049e565b606091505b50505050565b6060306001600160a01b031663267c4ae46040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156104e157600080fd5b505af11580156104f5573d6000803e3d6000fd5b505050506040513d602081101561050b57600080fd5b50506040516000906001600160a01b0384169085908381818185875af1925050503d8060008114610558576040519150601f19603f3d011682016040523d82523d6000602084013e61055d565b606091505b5095945050505050565b60015481565b6000806002340460405161058090610a3a565b660c1e0d0ccc8d4d60ca1b8152604051908190036020019082f09050801580156105ae573d6000803e3d6000fd5b5090506040516105bd90610a47565b604051809103906000f080156105d05760015b6105d9576105db565b505b6040516105e790610a53565b604051809103906000f080158015610603573d6000803e3d6000fd5b5050600060405161061390610a5f565b604051809103906000f08015801561062f573d6000803e3d6000fd5b509050600060405161064090610a5f565b604051809103906000f08015801561065c573d6000803e3d6000fd5b50604080516001600160a01b038316602480830182905283518084038201815260449384018552602081810180516001600160e01b03908116632c46fab560e21b1782528751808601969096528751808703909501855294909501865282810180519094166362ed243b60e11b17909352805194519596509490939287929091908185846000875af2506044016040819052602085810192508185846000875af250604481016040525050806001600160a01b0316846040518082805190602001908083835b602083106107415780518252601f199092019160209182019101610722565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146107a1576040519150601f19603f3d011682016040523d82523d6000602084013e6107a6565b606091505b505050806001600160a01b0316836040518082805190602001908083835b602083106107e35780518252601f1990920191602091820191016107c4565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610843576040519150601f19603f3d011682016040523d82523d6000602084013e610848565b606091505b5050604080516001600160a01b0388811660248084019190915283518084039091018152604490920183526020820180516001600160e01b0316631963f44160e21b178152925182519186169450919282918083835b602083106108bd5780518252601f19909201916020918201910161089e565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461091d576040519150601f19603f3d011682016040523d82523d6000602084013e610922565b606091505b505050856001600160a01b031663b11bead4866040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050602060405180830381600087803b15801561097457600080fd5b505af1158015610988573d6000803e3d6000fd5b505050506040513d602081101561099e57600080fd5b5050604080516370e5872960e11b815290516001600160a01b0389169163e1cb0e529160048083019260209291908290030181600087803b1580156109e257600080fd5b505af11580156109f6573d6000803e3d6000fd5b505050506040513d6020811015610a0c57600080fd5b50519998505050505050505050565b604051806101400160405280600a906020820280368337509192915050565b6102b180610a6d83390190565b605380610d1e83390190565b601480610d7183390190565b61031380610d858339019056fe60806040526040516102b13803806102b18339818101604052602081101561002657600080fd5b50516040805163099f12b960e21b81529051339163267c4ae49160048083019260209291908290030181600087803b15801561006157600080fd5b505af1158015610075573d6000803e3d6000fd5b505050506040513d602081101561008b57600080fd5b508190506002340461028e6040516100a290610127565b90815260405183918190036020019083f5915050801580156100c8573d6000803e3d6000fd5b505060408051639b7c9da360e01b81526036600482015290513391639b7c9da391602480830192600092919082900301818387803b15801561010957600080fd5b505af115801561011d573d6000803e3d6000fd5b5050505050610133565b60e7806101ca83390190565b6089806101416000396000f3fe608060405260043610601f5760003560e01c8063e1cb0e5214602a576025565b36602557005b600080fd5b348015603557600080fd5b50603c604e565b60408051918252519081900360200190f35b60149056fea2646970667358221220884dc0aad81e4fbdfb03eea198fb27577f71d8f82244e896729d8603bf285abf64736f6c6343000706003360806040526040516100e73803806100e78339818101604052602081101561002657600080fd5b505160405133906002340480156108fc02916000818181858888f19350505050158015610057573d6000803e3d6000fd5b50506080806100676000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e1cb0e5214602d575b600080fd5b60336045565b60408051918252519081900360200190f35b60149056fea2646970667358221220b58936794eb5834de57f1e8a0a4e598f105dfe7dbf0092c89c1732bab85aa57464736f6c634300070600336080604052348015600f57600080fd5b506040805162461bcd60e51b8152602060048201526012602482015271125b9d195b9d1a5bdb985b081c995d995c9d60721b604482015290519081900360640190fdfe6080604052348015600f57600080fd5b5033fffe608060405234801561001057600080fd5b506102f3806100206000396000f3fe60806040526004361061004a5760003560e01c8063658fd1041461004f5780636b59084d1461008757806383197ef01461008f578063b11bead4146100a6578063c5da4876146100cc575b600080fd5b6100756004803603602081101561006557600080fd5b50356001600160a01b03166100f2565b60408051918252519081900360200190f35b610075610161565b34801561009b57600080fd5b506100a4610166565b005b610075600480360360208110156100bc57600080fd5b50356001600160a01b0316610169565b610075600480360360208110156100e257600080fd5b50356001600160a01b0316610267565b6000816001600160a01b0316636b59084d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561012f57600080fd5b505af1158015610143573d6000803e3d6000fd5b505050506040513d602081101561015957600080fd5b505192915050565b600a90565b33ff5b60408051600481526024810182526020810180516001600160e01b0316636b59084d60e01b17815291518151600093849384936001600160a01b03881693919290918291908083835b602083106101d15780518252601f1990920191602091820191016101b2565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610231576040519150601f19603f3d011682016040523d82523d6000602084013e610236565b606091505b50915091508161024557600080fd5b600081806020019051602081101561025c57600080fd5b505195945050505050565b6040805160048152602481018252602080820180516001600160e01b0316636b59084d60e01b178152825193516000949285929190818584868b5af29250805195505050806102b557600080fd5b50505091905056fea2646970667358221220b79a409fb525793c6a8cc2d5ce9cb15ffc47b826e110cabff084a66dbab52cb864736f6c63430007060033a26469706673582212209c00d5cb25c01282c7c9a91013f33398c6aab554b65f7a62e1a9a4ea17d7c48864736f6c63430007060033",
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
