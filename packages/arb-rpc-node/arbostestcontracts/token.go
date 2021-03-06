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

// FailedERC20ABI is the input ABI used to generate the binding from.
const FailedERC20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"adminMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FailedERC20FuncSigs maps the 4-byte function signature to its string representation.
var FailedERC20FuncSigs = map[string]string{
	"e58306f9": "adminMint(address,uint256)",
}

// FailedERC20Bin is the compiled bytecode used for deploying new contracts.
var FailedERC20Bin = "0x6080604052348015600f57600080fd5b5060cf8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e58306f914602d575b600080fd5b605660048036036040811015604157600080fd5b506001600160a01b0381351690602001356058565b005b6040805162461bcd60e51b81526020600482015260116024820152706d696e7420616c77617973206661696c7360781b604482015290519081900360640190fdfea2646970667358221220ca64146f5426ebd333e3caf063afc082cee815732ed7b575cb1f3fc4220a97d564736f6c634300060c0033"

// DeployFailedERC20 deploys a new Ethereum contract, binding an instance of FailedERC20 to it.
func DeployFailedERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FailedERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(FailedERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FailedERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FailedERC20{FailedERC20Caller: FailedERC20Caller{contract: contract}, FailedERC20Transactor: FailedERC20Transactor{contract: contract}, FailedERC20Filterer: FailedERC20Filterer{contract: contract}}, nil
}

// FailedERC20 is an auto generated Go binding around an Ethereum contract.
type FailedERC20 struct {
	FailedERC20Caller     // Read-only binding to the contract
	FailedERC20Transactor // Write-only binding to the contract
	FailedERC20Filterer   // Log filterer for contract events
}

// FailedERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type FailedERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FailedERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FailedERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FailedERC20Session struct {
	Contract     *FailedERC20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FailedERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FailedERC20CallerSession struct {
	Contract *FailedERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FailedERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FailedERC20TransactorSession struct {
	Contract     *FailedERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FailedERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type FailedERC20Raw struct {
	Contract *FailedERC20 // Generic contract binding to access the raw methods on
}

// FailedERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FailedERC20CallerRaw struct {
	Contract *FailedERC20Caller // Generic read-only contract binding to access the raw methods on
}

// FailedERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FailedERC20TransactorRaw struct {
	Contract *FailedERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFailedERC20 creates a new instance of FailedERC20, bound to a specific deployed contract.
func NewFailedERC20(address common.Address, backend bind.ContractBackend) (*FailedERC20, error) {
	contract, err := bindFailedERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FailedERC20{FailedERC20Caller: FailedERC20Caller{contract: contract}, FailedERC20Transactor: FailedERC20Transactor{contract: contract}, FailedERC20Filterer: FailedERC20Filterer{contract: contract}}, nil
}

// NewFailedERC20Caller creates a new read-only instance of FailedERC20, bound to a specific deployed contract.
func NewFailedERC20Caller(address common.Address, caller bind.ContractCaller) (*FailedERC20Caller, error) {
	contract, err := bindFailedERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FailedERC20Caller{contract: contract}, nil
}

// NewFailedERC20Transactor creates a new write-only instance of FailedERC20, bound to a specific deployed contract.
func NewFailedERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*FailedERC20Transactor, error) {
	contract, err := bindFailedERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FailedERC20Transactor{contract: contract}, nil
}

// NewFailedERC20Filterer creates a new log filterer instance of FailedERC20, bound to a specific deployed contract.
func NewFailedERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*FailedERC20Filterer, error) {
	contract, err := bindFailedERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FailedERC20Filterer{contract: contract}, nil
}

// bindFailedERC20 binds a generic wrapper to an already deployed contract.
func bindFailedERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FailedERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FailedERC20 *FailedERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FailedERC20.Contract.FailedERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FailedERC20 *FailedERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FailedERC20.Contract.FailedERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FailedERC20 *FailedERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FailedERC20.Contract.FailedERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FailedERC20 *FailedERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FailedERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FailedERC20 *FailedERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FailedERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FailedERC20 *FailedERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FailedERC20.Contract.contract.Transact(opts, method, params...)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address account, uint256 amount) returns()
func (_FailedERC20 *FailedERC20Transactor) AdminMint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FailedERC20.contract.Transact(opts, "adminMint", account, amount)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address account, uint256 amount) returns()
func (_FailedERC20 *FailedERC20Session) AdminMint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FailedERC20.Contract.AdminMint(&_FailedERC20.TransactOpts, account, amount)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address account, uint256 amount) returns()
func (_FailedERC20 *FailedERC20TransactorSession) AdminMint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FailedERC20.Contract.AdminMint(&_FailedERC20.TransactOpts, account, amount)
}

// FailedERC721ABI is the input ABI used to generate the binding from.
const FailedERC721ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"adminMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FailedERC721FuncSigs maps the 4-byte function signature to its string representation.
var FailedERC721FuncSigs = map[string]string{
	"e58306f9": "adminMint(address,uint256)",
}

// FailedERC721Bin is the compiled bytecode used for deploying new contracts.
var FailedERC721Bin = "0x6080604052348015600f57600080fd5b5060cf8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e58306f914602d575b600080fd5b605660048036036040811015604157600080fd5b506001600160a01b0381351690602001356058565b005b6040805162461bcd60e51b81526020600482015260116024820152706d696e7420616c77617973206661696c7360781b604482015290519081900360640190fdfea26469706673582212200788e51560c6e4e15b8b7bfedf9586081b78b6989428ef7b1d4f06a380e3223464736f6c634300060c0033"

// DeployFailedERC721 deploys a new Ethereum contract, binding an instance of FailedERC721 to it.
func DeployFailedERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FailedERC721, error) {
	parsed, err := abi.JSON(strings.NewReader(FailedERC721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FailedERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FailedERC721{FailedERC721Caller: FailedERC721Caller{contract: contract}, FailedERC721Transactor: FailedERC721Transactor{contract: contract}, FailedERC721Filterer: FailedERC721Filterer{contract: contract}}, nil
}

// FailedERC721 is an auto generated Go binding around an Ethereum contract.
type FailedERC721 struct {
	FailedERC721Caller     // Read-only binding to the contract
	FailedERC721Transactor // Write-only binding to the contract
	FailedERC721Filterer   // Log filterer for contract events
}

// FailedERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type FailedERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FailedERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FailedERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FailedERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FailedERC721Session struct {
	Contract     *FailedERC721     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FailedERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FailedERC721CallerSession struct {
	Contract *FailedERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FailedERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FailedERC721TransactorSession struct {
	Contract     *FailedERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FailedERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type FailedERC721Raw struct {
	Contract *FailedERC721 // Generic contract binding to access the raw methods on
}

// FailedERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FailedERC721CallerRaw struct {
	Contract *FailedERC721Caller // Generic read-only contract binding to access the raw methods on
}

// FailedERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FailedERC721TransactorRaw struct {
	Contract *FailedERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFailedERC721 creates a new instance of FailedERC721, bound to a specific deployed contract.
func NewFailedERC721(address common.Address, backend bind.ContractBackend) (*FailedERC721, error) {
	contract, err := bindFailedERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FailedERC721{FailedERC721Caller: FailedERC721Caller{contract: contract}, FailedERC721Transactor: FailedERC721Transactor{contract: contract}, FailedERC721Filterer: FailedERC721Filterer{contract: contract}}, nil
}

// NewFailedERC721Caller creates a new read-only instance of FailedERC721, bound to a specific deployed contract.
func NewFailedERC721Caller(address common.Address, caller bind.ContractCaller) (*FailedERC721Caller, error) {
	contract, err := bindFailedERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FailedERC721Caller{contract: contract}, nil
}

// NewFailedERC721Transactor creates a new write-only instance of FailedERC721, bound to a specific deployed contract.
func NewFailedERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*FailedERC721Transactor, error) {
	contract, err := bindFailedERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FailedERC721Transactor{contract: contract}, nil
}

// NewFailedERC721Filterer creates a new log filterer instance of FailedERC721, bound to a specific deployed contract.
func NewFailedERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*FailedERC721Filterer, error) {
	contract, err := bindFailedERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FailedERC721Filterer{contract: contract}, nil
}

// bindFailedERC721 binds a generic wrapper to an already deployed contract.
func bindFailedERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FailedERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FailedERC721 *FailedERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FailedERC721.Contract.FailedERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FailedERC721 *FailedERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FailedERC721.Contract.FailedERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FailedERC721 *FailedERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FailedERC721.Contract.FailedERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FailedERC721 *FailedERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FailedERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FailedERC721 *FailedERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FailedERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FailedERC721 *FailedERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FailedERC721.Contract.contract.Transact(opts, method, params...)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address account, uint256 tokenId) returns()
func (_FailedERC721 *FailedERC721Transactor) AdminMint(opts *bind.TransactOpts, account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FailedERC721.contract.Transact(opts, "adminMint", account, tokenId)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address account, uint256 tokenId) returns()
func (_FailedERC721 *FailedERC721Session) AdminMint(account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FailedERC721.Contract.AdminMint(&_FailedERC721.TransactOpts, account, tokenId)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address account, uint256 tokenId) returns()
func (_FailedERC721 *FailedERC721TransactorSession) AdminMint(account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FailedERC721.Contract.AdminMint(&_FailedERC721.TransactOpts, account, tokenId)
}
