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

// FTRecordsABI is the input ABI used to generate the binding from.
const FTRecordsABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FTRecordsFuncSigs maps the 4-byte function signature to its string representation.
var FTRecordsFuncSigs = map[string]string{
	"c3a8962c": "getERC20Balance(address,address)",
}

// FTRecordsBin is the compiled bytecode used for deploying new contracts.
var FTRecordsBin = "0x608060405234801561001057600080fd5b50610104806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c3a8962c14602d575b600080fd5b605860048036036040811015604157600080fd5b506001600160a01b0381358116916020013516606a565b60408051918252519081900360200190f35b6001600160a01b038082166000908152602081815260408083209386168352908390528120549091908060a15760009250505060c9565b81600101600182038154811060b257fe5b906000526020600020906002020160010154925050505b9291505056fea265627a7a72315820707e222264b3323f5c48acb8baf592b5dcf0667b836a0df6d6d7678ac44b769e64736f6c634300050f0032"

// DeployFTRecords deploys a new Ethereum contract, binding an instance of FTRecords to it.
func DeployFTRecords(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FTRecords, error) {
	parsed, err := abi.JSON(strings.NewReader(FTRecordsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FTRecordsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FTRecords{FTRecordsCaller: FTRecordsCaller{contract: contract}, FTRecordsTransactor: FTRecordsTransactor{contract: contract}, FTRecordsFilterer: FTRecordsFilterer{contract: contract}}, nil
}

// FTRecords is an auto generated Go binding around an Ethereum contract.
type FTRecords struct {
	FTRecordsCaller     // Read-only binding to the contract
	FTRecordsTransactor // Write-only binding to the contract
	FTRecordsFilterer   // Log filterer for contract events
}

// FTRecordsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FTRecordsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FTRecordsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FTRecordsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FTRecordsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FTRecordsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FTRecordsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FTRecordsSession struct {
	Contract     *FTRecords        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FTRecordsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FTRecordsCallerSession struct {
	Contract *FTRecordsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FTRecordsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FTRecordsTransactorSession struct {
	Contract     *FTRecordsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FTRecordsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FTRecordsRaw struct {
	Contract *FTRecords // Generic contract binding to access the raw methods on
}

// FTRecordsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FTRecordsCallerRaw struct {
	Contract *FTRecordsCaller // Generic read-only contract binding to access the raw methods on
}

// FTRecordsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FTRecordsTransactorRaw struct {
	Contract *FTRecordsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFTRecords creates a new instance of FTRecords, bound to a specific deployed contract.
func NewFTRecords(address common.Address, backend bind.ContractBackend) (*FTRecords, error) {
	contract, err := bindFTRecords(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FTRecords{FTRecordsCaller: FTRecordsCaller{contract: contract}, FTRecordsTransactor: FTRecordsTransactor{contract: contract}, FTRecordsFilterer: FTRecordsFilterer{contract: contract}}, nil
}

// NewFTRecordsCaller creates a new read-only instance of FTRecords, bound to a specific deployed contract.
func NewFTRecordsCaller(address common.Address, caller bind.ContractCaller) (*FTRecordsCaller, error) {
	contract, err := bindFTRecords(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FTRecordsCaller{contract: contract}, nil
}

// NewFTRecordsTransactor creates a new write-only instance of FTRecords, bound to a specific deployed contract.
func NewFTRecordsTransactor(address common.Address, transactor bind.ContractTransactor) (*FTRecordsTransactor, error) {
	contract, err := bindFTRecords(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FTRecordsTransactor{contract: contract}, nil
}

// NewFTRecordsFilterer creates a new log filterer instance of FTRecords, bound to a specific deployed contract.
func NewFTRecordsFilterer(address common.Address, filterer bind.ContractFilterer) (*FTRecordsFilterer, error) {
	contract, err := bindFTRecords(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FTRecordsFilterer{contract: contract}, nil
}

// bindFTRecords binds a generic wrapper to an already deployed contract.
func bindFTRecords(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FTRecordsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FTRecords *FTRecordsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FTRecords.Contract.FTRecordsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FTRecords *FTRecordsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FTRecords.Contract.FTRecordsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FTRecords *FTRecordsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FTRecords.Contract.FTRecordsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FTRecords *FTRecordsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FTRecords.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FTRecords *FTRecordsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FTRecords.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FTRecords *FTRecordsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FTRecords.Contract.contract.Transact(opts, method, params...)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_FTRecords *FTRecordsCaller) GetERC20Balance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FTRecords.contract.Call(opts, out, "getERC20Balance", _tokenContract, _owner)
	return *ret0, err
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_FTRecords *FTRecordsSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _FTRecords.Contract.GetERC20Balance(&_FTRecords.CallOpts, _tokenContract, _owner)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_FTRecords *FTRecordsCallerSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _FTRecords.Contract.GetERC20Balance(&_FTRecords.CallOpts, _tokenContract, _owner)
}

// GlobalEthWalletABI is the input ABI used to generate the binding from.
const GlobalEthWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalEthWalletFuncSigs maps the 4-byte function signature to its string representation.
var GlobalEthWalletFuncSigs = map[string]string{
	"4d2301cc": "getEthBalance(address)",
	"a0ef91df": "withdrawEth()",
}

// GlobalEthWalletBin is the compiled bytecode used for deploying new contracts.
var GlobalEthWalletBin = "0x608060405234801561001057600080fd5b50610110806100206000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80634d2301cc146037578063a0ef91df14606c575b600080fd5b605a60048036036020811015604b57600080fd5b50356001600160a01b03166074565b60408051918252519081900360200190f35b6072608f565b005b6001600160a01b031660009081526020819052604090205490565b60006098336074565b3360008181526020819052604080822082905551929350909183156108fc0291849190818181858888f1935050505015801560d7573d6000803e3d6000fd5b505056fea265627a7a723158202365b7f05add3acdbdebe5b202ec070e77c3f5eacef2ea4cf445b7bf951f9aff64736f6c634300050f0032"

// DeployGlobalEthWallet deploys a new Ethereum contract, binding an instance of GlobalEthWallet to it.
func DeployGlobalEthWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalEthWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalEthWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalEthWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalEthWallet{GlobalEthWalletCaller: GlobalEthWalletCaller{contract: contract}, GlobalEthWalletTransactor: GlobalEthWalletTransactor{contract: contract}, GlobalEthWalletFilterer: GlobalEthWalletFilterer{contract: contract}}, nil
}

// GlobalEthWallet is an auto generated Go binding around an Ethereum contract.
type GlobalEthWallet struct {
	GlobalEthWalletCaller     // Read-only binding to the contract
	GlobalEthWalletTransactor // Write-only binding to the contract
	GlobalEthWalletFilterer   // Log filterer for contract events
}

// GlobalEthWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalEthWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalEthWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalEthWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalEthWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalEthWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalEthWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalEthWalletSession struct {
	Contract     *GlobalEthWallet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalEthWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalEthWalletCallerSession struct {
	Contract *GlobalEthWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GlobalEthWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalEthWalletTransactorSession struct {
	Contract     *GlobalEthWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GlobalEthWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalEthWalletRaw struct {
	Contract *GlobalEthWallet // Generic contract binding to access the raw methods on
}

// GlobalEthWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalEthWalletCallerRaw struct {
	Contract *GlobalEthWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalEthWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalEthWalletTransactorRaw struct {
	Contract *GlobalEthWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalEthWallet creates a new instance of GlobalEthWallet, bound to a specific deployed contract.
func NewGlobalEthWallet(address common.Address, backend bind.ContractBackend) (*GlobalEthWallet, error) {
	contract, err := bindGlobalEthWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalEthWallet{GlobalEthWalletCaller: GlobalEthWalletCaller{contract: contract}, GlobalEthWalletTransactor: GlobalEthWalletTransactor{contract: contract}, GlobalEthWalletFilterer: GlobalEthWalletFilterer{contract: contract}}, nil
}

// NewGlobalEthWalletCaller creates a new read-only instance of GlobalEthWallet, bound to a specific deployed contract.
func NewGlobalEthWalletCaller(address common.Address, caller bind.ContractCaller) (*GlobalEthWalletCaller, error) {
	contract, err := bindGlobalEthWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalEthWalletCaller{contract: contract}, nil
}

// NewGlobalEthWalletTransactor creates a new write-only instance of GlobalEthWallet, bound to a specific deployed contract.
func NewGlobalEthWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalEthWalletTransactor, error) {
	contract, err := bindGlobalEthWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalEthWalletTransactor{contract: contract}, nil
}

// NewGlobalEthWalletFilterer creates a new log filterer instance of GlobalEthWallet, bound to a specific deployed contract.
func NewGlobalEthWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalEthWalletFilterer, error) {
	contract, err := bindGlobalEthWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalEthWalletFilterer{contract: contract}, nil
}

// bindGlobalEthWallet binds a generic wrapper to an already deployed contract.
func bindGlobalEthWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalEthWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalEthWallet *GlobalEthWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalEthWallet.Contract.GlobalEthWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalEthWallet *GlobalEthWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.GlobalEthWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalEthWallet *GlobalEthWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.GlobalEthWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalEthWallet *GlobalEthWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalEthWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalEthWallet *GlobalEthWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalEthWallet *GlobalEthWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.contract.Transact(opts, method, params...)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) view returns(uint256)
func (_GlobalEthWallet *GlobalEthWalletCaller) GetEthBalance(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalEthWallet.contract.Call(opts, out, "getEthBalance", _owner)
	return *ret0, err
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) view returns(uint256)
func (_GlobalEthWallet *GlobalEthWalletSession) GetEthBalance(_owner common.Address) (*big.Int, error) {
	return _GlobalEthWallet.Contract.GetEthBalance(&_GlobalEthWallet.CallOpts, _owner)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) view returns(uint256)
func (_GlobalEthWallet *GlobalEthWalletCallerSession) GetEthBalance(_owner common.Address) (*big.Int, error) {
	return _GlobalEthWallet.Contract.GetEthBalance(&_GlobalEthWallet.CallOpts, _owner)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalEthWallet *GlobalEthWalletTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalEthWallet.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalEthWallet *GlobalEthWalletSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.WithdrawEth(&_GlobalEthWallet.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalEthWallet *GlobalEthWalletTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.WithdrawEth(&_GlobalEthWallet.TransactOpts)
}

// GlobalFTWalletABI is the input ABI used to generate the binding from.
const GlobalFTWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC20s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalFTWalletFuncSigs maps the 4-byte function signature to its string representation.
var GlobalFTWalletFuncSigs = map[string]string{
	"c3a8962c": "getERC20Balance(address,address)",
	"6e2b89c5": "ownedERC20s(address)",
	"f4f3b200": "withdrawERC20(address)",
}

// GlobalFTWalletBin is the compiled bytecode used for deploying new contracts.
var GlobalFTWalletBin = "0x608060405234801561001057600080fd5b50610516806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636e2b89c514610046578063c3a8962c146100bc578063f4f3b200146100fc575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b0316610124565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100a8578181015183820152602001610090565b505050509050019250505060405180910390f35b6100ea600480360360408110156100d257600080fd5b506001600160a01b03813581169160200135166101e5565b60408051918252519081900360200190f35b6101226004803603602081101561011257600080fd5b50356001600160a01b031661024d565b005b6001600160a01b03811660009081526020818152604091829020600181015483518181528184028101909301909352606092909183918015610170578160200160208202803883390190505b50805190915060005b818110156101db5783600101818154811061019057fe5b600091825260209091206002909102015483516001600160a01b03909116908490839081106101bb57fe5b6001600160a01b0390921660209283029190910190910152600101610179565b5090949350505050565b6001600160a01b038082166000908152602081815260408083209386168352908390528120549091908061021e57600092505050610247565b81600101600182038154811061023057fe5b906000526020600020906002020160010154925050505b92915050565b600061025982336101e5565b9050610266338383610320565b6102a15760405162461bcd60e51b815260040180806020018281038252602e8152602001806104b4602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b1580156102f057600080fd5b505af1158015610304573d6000803e3d6000fd5b505050506040513d602081101561031a57600080fd5b50505050565b60008161032f575060016104ac565b6001600160a01b0380851660009081526020818152604080832093871683529083905290205480610365576000925050506104ac565b600082600101600183038154811061037957fe5b9060005260206000209060020201905080600101548511156103a157600093505050506104ac565b600181018054869003908190556104a457600183018054839185916000919060001981019081106103ce57fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061040c57fe5b906000526020600020906002020183600101600184038154811061042c57fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061047a57fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b600193505050505b939250505056fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a72315820ac2830c558e731ea9055a68f80a1a89a4902de456c4d69afdded13362a5c48ba64736f6c634300050f0032"

// DeployGlobalFTWallet deploys a new Ethereum contract, binding an instance of GlobalFTWallet to it.
func DeployGlobalFTWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalFTWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalFTWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalFTWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalFTWallet{GlobalFTWalletCaller: GlobalFTWalletCaller{contract: contract}, GlobalFTWalletTransactor: GlobalFTWalletTransactor{contract: contract}, GlobalFTWalletFilterer: GlobalFTWalletFilterer{contract: contract}}, nil
}

// GlobalFTWallet is an auto generated Go binding around an Ethereum contract.
type GlobalFTWallet struct {
	GlobalFTWalletCaller     // Read-only binding to the contract
	GlobalFTWalletTransactor // Write-only binding to the contract
	GlobalFTWalletFilterer   // Log filterer for contract events
}

// GlobalFTWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalFTWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalFTWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalFTWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalFTWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalFTWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalFTWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalFTWalletSession struct {
	Contract     *GlobalFTWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalFTWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalFTWalletCallerSession struct {
	Contract *GlobalFTWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GlobalFTWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalFTWalletTransactorSession struct {
	Contract     *GlobalFTWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GlobalFTWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalFTWalletRaw struct {
	Contract *GlobalFTWallet // Generic contract binding to access the raw methods on
}

// GlobalFTWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalFTWalletCallerRaw struct {
	Contract *GlobalFTWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalFTWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalFTWalletTransactorRaw struct {
	Contract *GlobalFTWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalFTWallet creates a new instance of GlobalFTWallet, bound to a specific deployed contract.
func NewGlobalFTWallet(address common.Address, backend bind.ContractBackend) (*GlobalFTWallet, error) {
	contract, err := bindGlobalFTWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalFTWallet{GlobalFTWalletCaller: GlobalFTWalletCaller{contract: contract}, GlobalFTWalletTransactor: GlobalFTWalletTransactor{contract: contract}, GlobalFTWalletFilterer: GlobalFTWalletFilterer{contract: contract}}, nil
}

// NewGlobalFTWalletCaller creates a new read-only instance of GlobalFTWallet, bound to a specific deployed contract.
func NewGlobalFTWalletCaller(address common.Address, caller bind.ContractCaller) (*GlobalFTWalletCaller, error) {
	contract, err := bindGlobalFTWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalFTWalletCaller{contract: contract}, nil
}

// NewGlobalFTWalletTransactor creates a new write-only instance of GlobalFTWallet, bound to a specific deployed contract.
func NewGlobalFTWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalFTWalletTransactor, error) {
	contract, err := bindGlobalFTWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalFTWalletTransactor{contract: contract}, nil
}

// NewGlobalFTWalletFilterer creates a new log filterer instance of GlobalFTWallet, bound to a specific deployed contract.
func NewGlobalFTWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalFTWalletFilterer, error) {
	contract, err := bindGlobalFTWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalFTWalletFilterer{contract: contract}, nil
}

// bindGlobalFTWallet binds a generic wrapper to an already deployed contract.
func bindGlobalFTWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalFTWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalFTWallet *GlobalFTWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalFTWallet.Contract.GlobalFTWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalFTWallet *GlobalFTWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.GlobalFTWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalFTWallet *GlobalFTWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.GlobalFTWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalFTWallet *GlobalFTWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalFTWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalFTWallet *GlobalFTWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalFTWallet *GlobalFTWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.contract.Transact(opts, method, params...)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalFTWallet *GlobalFTWalletCaller) GetERC20Balance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalFTWallet.contract.Call(opts, out, "getERC20Balance", _tokenContract, _owner)
	return *ret0, err
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalFTWallet *GlobalFTWalletSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalFTWallet.Contract.GetERC20Balance(&_GlobalFTWallet.CallOpts, _tokenContract, _owner)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalFTWallet *GlobalFTWalletCallerSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalFTWallet.Contract.GetERC20Balance(&_GlobalFTWallet.CallOpts, _tokenContract, _owner)
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) view returns(address[])
func (_GlobalFTWallet *GlobalFTWalletCaller) OwnedERC20s(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GlobalFTWallet.contract.Call(opts, out, "ownedERC20s", _owner)
	return *ret0, err
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) view returns(address[])
func (_GlobalFTWallet *GlobalFTWalletSession) OwnedERC20s(_owner common.Address) ([]common.Address, error) {
	return _GlobalFTWallet.Contract.OwnedERC20s(&_GlobalFTWallet.CallOpts, _owner)
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) view returns(address[])
func (_GlobalFTWallet *GlobalFTWalletCallerSession) OwnedERC20s(_owner common.Address) ([]common.Address, error) {
	return _GlobalFTWallet.Contract.OwnedERC20s(&_GlobalFTWallet.CallOpts, _owner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalFTWallet *GlobalFTWalletTransactor) WithdrawERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalFTWallet.contract.Transact(opts, "withdrawERC20", _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalFTWallet *GlobalFTWalletSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.WithdrawERC20(&_GlobalFTWallet.TransactOpts, _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalFTWallet *GlobalFTWalletTransactorSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.WithdrawERC20(&_GlobalFTWallet.TransactOpts, _tokenContract)
}

// GlobalInboxABI is the input ABI used to generate the binding from.
const GlobalInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BuddyContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"}],\"name\":\"MessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"PaymentTransfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"depositERC721Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositPairedERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC721Tokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getInbox\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"getPaymentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasERC721\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"}],\"name\":\"isPairedContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC20s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC721s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendInitializationMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeHashes\",\"type\":\"bytes32[]\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"transferPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawPairedERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalInboxFuncSigs maps the 4-byte function signature to its string representation.
var GlobalInboxFuncSigs = map[string]string{
	"44c75b7f": "deployL2ContractPair(address,bytes)",
	"bca22b76": "depositERC20Message(address,address,address,uint256)",
	"8b7010aa": "depositERC721Message(address,address,address,uint256)",
	"5bd21290": "depositEthMessage(address,address)",
	"cb184d0e": "depositPairedERC20Message(address,address,address,uint256)",
	"c3a8962c": "getERC20Balance(address,address)",
	"0758fb0a": "getERC721Tokens(address,address)",
	"4d2301cc": "getEthBalance(address)",
	"02201681": "getInbox(address)",
	"bd4fbb36": "getPaymentOwner(address,bytes32,uint256)",
	"45a53f09": "hasERC721(address,address,uint256)",
	"659e42cd": "isPairedContract(address,address)",
	"6e2b89c5": "ownedERC20s(address)",
	"33f2ac42": "ownedERC721s(address)",
	"5cc96efa": "sendInitializationMessage(bytes)",
	"74c6eccc": "sendL2Message(address,bytes)",
	"fbef861b": "sendL2MessageFromOrigin(address,bytes)",
	"072fd2bb": "sendMessages(bytes,uint256[],bytes32[])",
	"d2256c66": "transferPayment(address,address,bytes32,uint256)",
	"f4f3b200": "withdrawERC20(address)",
	"f3e414f8": "withdrawERC721(address,uint256)",
	"a0ef91df": "withdrawEth()",
	"21e2f2fa": "withdrawPairedERC20(address)",
}

// GlobalInboxBin is the compiled bytecode used for deploying new contracts.
var GlobalInboxBin = "0x608060405234801561001057600080fd5b50612969806100206000396000f3fe6080604052600436106101405760003560e01c80636e2b89c5116100b6578063c3a8962c1161006f578063c3a8962c1461076a578063cb184d0e146107a5578063d2256c66146107ee578063f3e414f814610837578063f4f3b20014610870578063fbef861b146108a357610140565b80636e2b89c5146105aa57806374c6eccc146105dd5780638b7010aa14610668578063a0ef91df146106b1578063bca22b76146106c6578063bd4fbb361461070f57610140565b806344c75b7f1161010857806344c75b7f1461039f57806345a53f091461042a5780634d2301cc146104815780635bd21290146104c65780635cc96efa146104f4578063659e42cd1461056f57610140565b80630220168114610145578063072fd2bb146101915780630758fb0a146102ae57806321e2f2fa1461033957806333f2ac421461036c575b600080fd5b34801561015157600080fd5b506101786004803603602081101561016857600080fd5b50356001600160a01b031661092e565b6040805192835260208301919091528051918290030190f35b34801561019d57600080fd5b506102ac600480360360608110156101b457600080fd5b810190602081018135600160201b8111156101ce57600080fd5b8201836020820111156101e057600080fd5b803590602001918460018302840111600160201b8311171561020157600080fd5b919390929091602081019035600160201b81111561021e57600080fd5b82018360208201111561023057600080fd5b803590602001918460208302840111600160201b8311171561025157600080fd5b919390929091602081019035600160201b81111561026e57600080fd5b82018360208201111561028057600080fd5b803590602001918460208302840111600160201b831117156102a157600080fd5b509092509050610954565b005b3480156102ba57600080fd5b506102e9600480360360408110156102d157600080fd5b506001600160a01b0381358116916020013516610a1c565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561032557818101518382015260200161030d565b505050509050019250505060405180910390f35b34801561034557600080fd5b506102ac6004803603602081101561035c57600080fd5b50356001600160a01b0316610ae2565b34801561037857600080fd5b506102e96004803603602081101561038f57600080fd5b50356001600160a01b0316610b98565b3480156103ab57600080fd5b506102ac600480360360408110156103c257600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156103ec57600080fd5b8201836020820111156103fe57600080fd5b803590602001918460018302840111600160201b8311171561041f57600080fd5b509092509050610c5b565b34801561043657600080fd5b5061046d6004803603606081101561044d57600080fd5b506001600160a01b03813581169160208101359091169060400135610cad565b604080519115158252519081900360200190f35b34801561048d57600080fd5b506104b4600480360360208110156104a457600080fd5b50356001600160a01b0316610d2d565b60408051918252519081900360200190f35b6102ac600480360360408110156104dc57600080fd5b506001600160a01b0381358116916020013516610d48565b34801561050057600080fd5b506102ac6004803603602081101561051757600080fd5b810190602081018135600160201b81111561053157600080fd5b82018360208201111561054357600080fd5b803590602001918460018302840111600160201b8311171561056457600080fd5b509092509050610d8e565b34801561057b57600080fd5b5061046d6004803603604081101561059257600080fd5b506001600160a01b0381358116916020013516610dd1565b3480156105b657600080fd5b506102e9600480360360208110156105cd57600080fd5b50356001600160a01b0316610dff565b3480156105e957600080fd5b506102ac6004803603604081101561060057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561062a57600080fd5b82018360208201111561063c57600080fd5b803590602001918460018302840111600160201b8311171561065d57600080fd5b509092509050610eb6565b34801561067457600080fd5b506102ac6004803603608081101561068b57600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610ef9565b3480156106bd57600080fd5b506102ac610f60565b3480156106d257600080fd5b506102ac600480360360808110156106e957600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610fab565b34801561071b57600080fd5b5061074e6004803603606081101561073257600080fd5b506001600160a01b03813516906020810135906040013561100c565b604080516001600160a01b039092168252519081900360200190f35b34801561077657600080fd5b506104b46004803603604081101561078d57600080fd5b506001600160a01b0381358116916020013516611077565b3480156107b157600080fd5b506102ac600480360360808110156107c857600080fd5b506001600160a01b038135811691602081013582169160408201351690606001356110de565b3480156107fa57600080fd5b506102ac6004803603608081101561081157600080fd5b506001600160a01b038135811691602081013590911690604081013590606001356110e9565b34801561084357600080fd5b506102ac6004803603604081101561085a57600080fd5b506001600160a01b03813516906020013561120c565b34801561087c57600080fd5b506102ac6004803603602081101561089357600080fd5b50356001600160a01b03166112bc565b3480156108af57600080fd5b506102ac600480360360408110156108c657600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156108f057600080fd5b82018360208201111561090257600080fd5b803590602001918460018302840111600160201b8311171561092357600080fd5b509092509050611389565b6001600160a01b038116600090815260056020526040902080546001909101545b915091565b60008061095f6127ea565b8360005b81811015610a0e5760005b89898381811061097a57fe5b90506020020135811015610a05576109c98c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250899250611442915050565b91975095509350856109e057505050505050610a14565b6109fd8888848181106109ef57fe5b90506020020135828661152e565b60010161096e565b50600101610963565b50505050505b505050505050565b6001600160a01b0380821660009081526003602090815260408083209386168352908390529020546060919080610a655750506040805160008152602081019091529050610adc565b816001016001820381548110610a7757fe5b9060005260206000209060030201600201805480602002602001604051908101604052809291908181526020018280548015610ad257602002820191906000526020600020905b815481526020019060010190808311610abe575b5050505050925050505b92915050565b6000610aee8233611077565b9050610afb33838361175c565b610b365760405162461bcd60e51b815260040180806020018281038252602e815260200180612907602e913960400191505060405180910390fd5b604080516340c10f1960e01b81523360048201526024810183905290516001600160a01b038416916340c10f1991604480830192600092919082900301818387803b158015610b8457600080fd5b505af1158015610a14573d6000803e3d6000fd5b6001600160a01b038116600090815260036020908152604091829020600181015483518181528184028101909301909352606092909183918015610be6578160200160208202803883390190505b50805190915060005b81811015610c5157836001018181548110610c0657fe5b600091825260209091206003909102015483516001600160a01b0390911690849083908110610c3157fe5b6001600160a01b0390921660209283029190910190910152600101610bef565b5090949350505050565b610c6533846118ed565b610ca88360053385858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061192492505050565b505050565b6001600160a01b03808316600090815260036020908152604080832093871683529083905281205490919080610ce857600092505050610d26565b816001016001820381548110610cfa57fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b6001600160a01b031660009081526002602052604090205490565b610d51826119fb565b604080516001600160a01b038316602082015234818301528151808203830181526060909101909152610d8a9083906000903390611924565b5050565b610d8a3360043385858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061192492505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6001600160a01b03811660009081526020818152604091829020600181015483518181528184028101909301909352606092909183918015610e4b578160200160208202803883390190505b50805190915060005b81811015610c5157836001018181548110610e6b57fe5b600091825260209091206002909102015483516001600160a01b0390911690849083908110610e9657fe5b6001600160a01b0390921660209283029190910190910152600101610e54565b610ca88360033385858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061192492505050565b610f04838583611a1a565b610f5a846002338660601b6001600160601b0319168660601b6001600160601b03191686604051602001808481526020018381526020018281526020019350505050604051602081830303815290604052611924565b50505050565b6000610f6b33610d2d565b3360008181526002602052604080822082905551929350909183156108fc0291849190818181858888f19350505050158015610d8a573d6000803e3d6000fd5b610fb6838583611a91565b610f5a846001338660601b6001600160601b0319168660601b6001600160601b03191686604051602001808481526020018381526020018281526020019350505050604051602081830303815290604052611924565b604080516020808201859052818301849052606086811b6001600160601b03191690830152825180830360540181526074909201835281519181019190912060009081526004909152908120546001600160a01b0316806110705784915050610d26565b9050610d26565b6001600160a01b03808216600090815260208181526040808320938616835290839052812054909190806110b057600092505050610adc565b8160010160018203815481106110c257fe5b9060005260206000209060020201600101549250505092915050565b610fb6838583611b1e565b60006110f685848461100c565b9050336001600160a01b0382161461114e576040805162461bcd60e51b815260206004820152601660248201527526bab9ba103132903830bcb6b2b73a1037bbb732b91760511b604482015290519081900360640190fd5b604080516020808201869052818301859052606088811b6001600160601b031916908301528251808303605401815260748301808552815191830191909120600090815260049092529083902080546001600160a01b0319166001600160a01b03898116918217909255918790526094830186905280891660b4840152841660d483015260f482015290517fb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c6438391610114908290030190a15050505050565b611217338383611c6b565b611268576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b158015610b8457600080fd5b60006112c88233611077565b90506112d533838361175c565b6113105760405162461bcd60e51b815260040180806020018281038252602e815260200180612907602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b15801561135f57600080fd5b505af1158015611373573d6000803e3d6000fd5b505050506040513d6020811015610f5a57600080fd5b3332146113cb576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b60006113f68460033386866040518083838082843760405192018290039091209350611ed392505050565b60408051828152905191925033916003916001600160a01b038816917fe923069519faf69b0726ed766a213f61b6f07f2ecf11d55582cc440d8806b0bc9181900360200190a450505050565b60008061144d6127ea565b839150600085838151811061145e57fe5b016020015160019093019260f81c9050611476611f27565b60030160ff168160ff1614611492575060009250839150611527565b600061149e8785611f2d565b91965094509050846114b95750600093508492506115279050565b60ff8116835260006114cb8886611f2d565b91975095509050856114e7575060009450859350611527915050565b6001600160a01b03811660208501526115008886611faa565b604087015290965094508561151f575060009450859350611527915050565b506001945050505b9250925092565b805160ff16611598576000611541612809565b61154e8360400151612201565b915091508161155e575050610ca8565b600061156f8260000151878761100c565b90506115808260000151878761225d565b61158f338284602001516122b9565b50505050610ca8565b805160ff16600114156116355760006115af612820565b6115bc8360400151612317565b91509150816115cc575050610ca8565b60006115dd8260200151878761100c565b90506115ee8260200151878761225d565b81516115fa9033610dd1565b1561161957611613338284600001518560400151612394565b5061162d565b61158f3382846000015185604001516123e0565b505050610ca8565b805160ff166002141561169f57600061164c612820565b6116598360400151612317565b9150915081611669575050610ca8565b600061167a8260200151878761100c565b905061168b8260200151878761225d565b61158f33828460000151856040015161240f565b805160ff1660051415610ca85780602001516001600160a01b03167fa98915d9854858ea787b0abcd4e8e3a96802bc19a25474a8b7017a303628e44482604001516040518080602001828103825283818151815260200191508051906020019080838360005b8381101561171d578181015183820152602001611705565b50505050905090810190601f16801561174a5780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b60008161176b57506001610d26565b6001600160a01b03808516600090815260208181526040808320938716835290839052902054806117a157600092505050610d26565b60008260010160018303815481106117b557fe5b9060005260206000209060020201905080600101548511156117dd5760009350505050610d26565b600181018054869003908190556118e0576001830180548391859160009190600019810190811061180a57fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061184857fe5b906000526020600020906002020183600101600184038154811061186857fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b03938416178155600194850154908501559089168252859052604081205583018054806118b657fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b5060019695505050505050565b6001600160a01b0391821660009081526001602081815260408084209490951683529290925291909120805460ff19169091179055565b60006119398585858580519060200120611ed3565b9050826001600160a01b03168460ff16866001600160a01b03167f35e48d636f39df5c5ca2278452d6d89bf9f07c2ff15f46d08aa402c46638b88284866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156119b95781810151838201526020016119a1565b50505050905090810190601f1680156119e65780820380516001836020036101000a031916815260200191505b50935050505060405180910390a45050505050565b6001600160a01b03166000908152600260205260409020805434019055565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd91606480830192600092919082900301818387803b158015611a6e57600080fd5b505af1158015611a82573d6000803e3d6000fd5b50505050610ca882848361242f565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd9160648083019260209291908290030181600087803b158015611ae657600080fd5b505af1158015611afa573d6000803e3d6000fd5b505050506040513d6020811015611b1057600080fd5b50610ca890508284836125b3565b611b288383610dd1565b611b79576040805162461bcd60e51b815260206004820152601760248201527f6d7573742062652070616972656420636f6e7472616374000000000000000000604482015290519081900360640190fd5b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd9160648083019260209291908290030181600087803b158015611bce57600080fd5b505af1158015611be2573d6000803e3d6000fd5b505050506040513d6020811015611bf857600080fd5b505060408051632770a7eb60e21b81523060048201526024810183905290516001600160a01b03851691639dc29fac91604480830192600092919082900301818387803b158015611c4857600080fd5b505af1158015611c5c573d6000803e3d6000fd5b50505050610ca88284836125b3565b6001600160a01b03808416600090815260036020908152604080832093861683529083905281205490919080611ca657600092505050610d26565b6000826001016001830381548110611cba57fe5b600091825260208083208884526001600390930201918201905260409091205490915080611cef576000945050505050610d26565b60028201805482916001850191600091906000198101908110611d0e57fe5b600091825260208083209091015483528201929092526040019020556002820180546000198101908110611d3e57fe5b9060005260206000200154826002016001830381548110611d5b57fe5b600091825260208083209091019290925587815260018401909152604081205560028201805480611d8857fe5b6000828152602081208201600019908101919091550190556002820154611ec55760018401805484918691600091906000198101908110611dc557fe5b600091825260208083206003909202909101546001600160a01b031683528201929092526040019020556001840180546000198101908110611e0357fe5b9060005260206000209060030201846001016001850381548110611e2357fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b0390921691909117815560028083018054611e669284019190612840565b5050506001600160a01b03871660009081526020859052604081205560018401805480611e8f57fe5b60008281526020812060036000199093019283020180546001600160a01b031916815590611ec06002830182612890565b505090555b506001979650505050505050565b6001600160a01b03841660009081526005602052604081206001808201540182611f0187874342868a612688565b9050611f118360000154826126f1565b835550600190910181905590505b949350505050565b60035b90565b6000806000808551905084811080611f4757506021858203105b80611f6f5750611f5561271d565b60ff16868681518110611f6457fe5b016020015160f81c14155b15611f84575060009250839150829050611527565b600160218601611f9c8888840163ffffffff61272216565b935093509350509250925092565b60008060606000611fbb868661273e565b9195509350905083611fd1575060009250611527565b60208104601f8216600081611fe7576000611fea565b60015b60ff168301905060608360405190808252806020026020018201604052801561201d578160200160208202803883390190505b5090506060836040519080825280601f01601f19166020018201604052801561204d576020820181803883390190505b5090506000805b848110156121185760006120688e8c61273e565b919d509b5090508b612087575060009a50611527975050505050505050565b811580156120955750600087115b156120e8578060005b888110156120e1578181602081106120b257fe5b1a60f81b8682815181106120c257fe5b60200101906001600160f81b031916908160001a90535060010161209e565b505061210f565b8060001b858460018b0303815181106120fd57fe5b60209081029190910101526001909201915b50600101612054565b5060006121258d8b6127a0565b909a509050612132611f27565b60ff168160ff16146121505750600099506115279650505050505050565b60018a858560405160200180838051906020019060200280838360005b8381101561218557818101518382015260200161216d565b5050505090500182805190602001908083835b602083106121b75780518252601f199092019160209182019101612198565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529a509a509a5050505050505050509250925092565b600061220b612809565b60348351101561221e576000915061094f565b600c612230848263ffffffff6127c716565b6001600160a01b0316825260140161224e848263ffffffff61272216565b60208301525060019150915091565b60408051602080820194909452808201929092526001600160601b0319606094851b16938201939093528251605481830301815260749091018352805190820120600090815260049091522080546001600160a01b0319169055565b6001600160a01b0383166000908152600260205260408120548211156122e157506000610d26565b506001600160a01b0392831660009081526002602052604080822080548490039055929093168352912080549091019055600190565b6000612321612820565b604883511015612334576000915061094f565b600c612346848263ffffffff6127c716565b6001600160a01b03168252602001612364848263ffffffff6127c716565b6001600160a01b03166020830152601401612385848263ffffffff61272216565b60408301525060019150915091565b6000806123a18487611077565b9050828110156123bc576123b686858361175c565b506123c9565b6123c786858561175c565b505b6123d48585856125b3565b50600195945050505050565b60006123ed85848461175c565b6123f957506000611f1f565b6124048484846125b3565b506001949350505050565b600061241c858484611c6b565b61242857506000611f1f565b6124048484845b6001600160a01b038084166000908152600360209081526040808320938616835290839052902054806124ef576040805180820182526001600160a01b0386811682528251600080825260208083019095528484019182526001878101805491820180825590835291869020855160039092020180546001600160a01b031916919094161783559051805191946124ce926002850192909101906128b1565b5050506001600160a01b038516600090815260208490526040902081905590505b600082600101600183038154811061250357fe5b906000526020600020906003020190508060010160008581526020019081526020016000205460001461257d576040805162461bcd60e51b815260206004820152601d60248201527f63616e27742061646420616c7265616479206f776e656420746f6b656e000000604482015290519081900360640190fd5b60028101805460018181018355600083815260208082209093018890559254968352909201909152604090209290925550505050565b806125bd57610ca8565b6001600160a01b038084166000908152602081815260408083209386168352908390529020548061265457506040805180820182526001600160a01b0385811680835260006020808501828152600188810180548083018083559186528486209851600290910290980180546001600160a01b03191698909716979097178655905194019390935590815290849052919091208190555b8282600101600183038154811061266757fe5b60009182526020909120600160029092020101805490910190555050505050565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600090565b6000816020018351101561273557600080fd5b50016020015190565b60008060008061274e86866127a0565b909350905061275b611f27565b60020160ff168160ff1614612774575060009250611527565b61277e8684611f2d565b9195509350915083612794575060009250611527565b50600192509250925092565b600080826001018484815181106127b357fe5b016020015190925060f81c90509250929050565b600081601401835110156127da57600080fd5b500160200151600160601b900490565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b604080516060810182526000808252602082018190529181019190915290565b8280548282559060005260206000209081019282156128805760005260206000209182015b82811115612880578254825591600101919060010190612865565b5061288c9291506128ec565b5090565b50805460008255906000526020600020908101906128ae91906128ec565b50565b828054828255906000526020600020908101928215612880579160200282015b828111156128805782518255916020019190600101906128d1565b611f2a91905b8082111561288c57600081556001016128f256fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a723158200bf919b610855e73dde3a987f0ae160ef33e98fe6bbd6feaa271927f4a87a35264736f6c634300050f0032"

// DeployGlobalInbox deploys a new Ethereum contract, binding an instance of GlobalInbox to it.
func DeployGlobalInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalInbox, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalInboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalInbox{GlobalInboxCaller: GlobalInboxCaller{contract: contract}, GlobalInboxTransactor: GlobalInboxTransactor{contract: contract}, GlobalInboxFilterer: GlobalInboxFilterer{contract: contract}}, nil
}

// GlobalInbox is an auto generated Go binding around an Ethereum contract.
type GlobalInbox struct {
	GlobalInboxCaller     // Read-only binding to the contract
	GlobalInboxTransactor // Write-only binding to the contract
	GlobalInboxFilterer   // Log filterer for contract events
}

// GlobalInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalInboxSession struct {
	Contract     *GlobalInbox      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalInboxCallerSession struct {
	Contract *GlobalInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GlobalInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalInboxTransactorSession struct {
	Contract     *GlobalInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GlobalInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalInboxRaw struct {
	Contract *GlobalInbox // Generic contract binding to access the raw methods on
}

// GlobalInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalInboxCallerRaw struct {
	Contract *GlobalInboxCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalInboxTransactorRaw struct {
	Contract *GlobalInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalInbox creates a new instance of GlobalInbox, bound to a specific deployed contract.
func NewGlobalInbox(address common.Address, backend bind.ContractBackend) (*GlobalInbox, error) {
	contract, err := bindGlobalInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalInbox{GlobalInboxCaller: GlobalInboxCaller{contract: contract}, GlobalInboxTransactor: GlobalInboxTransactor{contract: contract}, GlobalInboxFilterer: GlobalInboxFilterer{contract: contract}}, nil
}

// NewGlobalInboxCaller creates a new read-only instance of GlobalInbox, bound to a specific deployed contract.
func NewGlobalInboxCaller(address common.Address, caller bind.ContractCaller) (*GlobalInboxCaller, error) {
	contract, err := bindGlobalInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalInboxCaller{contract: contract}, nil
}

// NewGlobalInboxTransactor creates a new write-only instance of GlobalInbox, bound to a specific deployed contract.
func NewGlobalInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalInboxTransactor, error) {
	contract, err := bindGlobalInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalInboxTransactor{contract: contract}, nil
}

// NewGlobalInboxFilterer creates a new log filterer instance of GlobalInbox, bound to a specific deployed contract.
func NewGlobalInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalInboxFilterer, error) {
	contract, err := bindGlobalInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalInboxFilterer{contract: contract}, nil
}

// bindGlobalInbox binds a generic wrapper to an already deployed contract.
func bindGlobalInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalInbox *GlobalInboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalInbox.Contract.GlobalInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalInbox *GlobalInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalInbox.Contract.GlobalInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalInbox *GlobalInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalInbox.Contract.GlobalInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalInbox *GlobalInboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalInbox *GlobalInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalInbox *GlobalInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalInbox.Contract.contract.Transact(opts, method, params...)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxCaller) GetERC20Balance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "getERC20Balance", _tokenContract, _owner)
	return *ret0, err
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalInbox.Contract.GetERC20Balance(&_GlobalInbox.CallOpts, _tokenContract, _owner)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxCallerSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalInbox.Contract.GetERC20Balance(&_GlobalInbox.CallOpts, _tokenContract, _owner)
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) view returns(uint256[])
func (_GlobalInbox *GlobalInboxCaller) GetERC721Tokens(opts *bind.CallOpts, _erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "getERC721Tokens", _erc721, _owner)
	return *ret0, err
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) view returns(uint256[])
func (_GlobalInbox *GlobalInboxSession) GetERC721Tokens(_erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	return _GlobalInbox.Contract.GetERC721Tokens(&_GlobalInbox.CallOpts, _erc721, _owner)
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) view returns(uint256[])
func (_GlobalInbox *GlobalInboxCallerSession) GetERC721Tokens(_erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	return _GlobalInbox.Contract.GetERC721Tokens(&_GlobalInbox.CallOpts, _erc721, _owner)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxCaller) GetEthBalance(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "getEthBalance", _owner)
	return *ret0, err
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxSession) GetEthBalance(_owner common.Address) (*big.Int, error) {
	return _GlobalInbox.Contract.GetEthBalance(&_GlobalInbox.CallOpts, _owner)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxCallerSession) GetEthBalance(_owner common.Address) (*big.Int, error) {
	return _GlobalInbox.Contract.GetEthBalance(&_GlobalInbox.CallOpts, _owner)
}

// GetInbox is a free data retrieval call binding the contract method 0x02201681.
//
// Solidity: function getInbox(address account) view returns(bytes32, uint256)
func (_GlobalInbox *GlobalInboxCaller) GetInbox(opts *bind.CallOpts, account common.Address) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _GlobalInbox.contract.Call(opts, out, "getInbox", account)
	return *ret0, *ret1, err
}

// GetInbox is a free data retrieval call binding the contract method 0x02201681.
//
// Solidity: function getInbox(address account) view returns(bytes32, uint256)
func (_GlobalInbox *GlobalInboxSession) GetInbox(account common.Address) ([32]byte, *big.Int, error) {
	return _GlobalInbox.Contract.GetInbox(&_GlobalInbox.CallOpts, account)
}

// GetInbox is a free data retrieval call binding the contract method 0x02201681.
//
// Solidity: function getInbox(address account) view returns(bytes32, uint256)
func (_GlobalInbox *GlobalInboxCallerSession) GetInbox(account common.Address) ([32]byte, *big.Int, error) {
	return _GlobalInbox.Contract.GetInbox(&_GlobalInbox.CallOpts, account)
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0xbd4fbb36.
//
// Solidity: function getPaymentOwner(address originalOwner, bytes32 nodeHash, uint256 messageIndex) view returns(address)
func (_GlobalInbox *GlobalInboxCaller) GetPaymentOwner(opts *bind.CallOpts, originalOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "getPaymentOwner", originalOwner, nodeHash, messageIndex)
	return *ret0, err
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0xbd4fbb36.
//
// Solidity: function getPaymentOwner(address originalOwner, bytes32 nodeHash, uint256 messageIndex) view returns(address)
func (_GlobalInbox *GlobalInboxSession) GetPaymentOwner(originalOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (common.Address, error) {
	return _GlobalInbox.Contract.GetPaymentOwner(&_GlobalInbox.CallOpts, originalOwner, nodeHash, messageIndex)
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0xbd4fbb36.
//
// Solidity: function getPaymentOwner(address originalOwner, bytes32 nodeHash, uint256 messageIndex) view returns(address)
func (_GlobalInbox *GlobalInboxCallerSession) GetPaymentOwner(originalOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (common.Address, error) {
	return _GlobalInbox.Contract.GetPaymentOwner(&_GlobalInbox.CallOpts, originalOwner, nodeHash, messageIndex)
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) view returns(bool)
func (_GlobalInbox *GlobalInboxCaller) HasERC721(opts *bind.CallOpts, _erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "hasERC721", _erc721, _owner, _tokenId)
	return *ret0, err
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) view returns(bool)
func (_GlobalInbox *GlobalInboxSession) HasERC721(_erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalInbox.Contract.HasERC721(&_GlobalInbox.CallOpts, _erc721, _owner, _tokenId)
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) view returns(bool)
func (_GlobalInbox *GlobalInboxCallerSession) HasERC721(_erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalInbox.Contract.HasERC721(&_GlobalInbox.CallOpts, _erc721, _owner, _tokenId)
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(bool)
func (_GlobalInbox *GlobalInboxCaller) IsPairedContract(opts *bind.CallOpts, _tokenContract common.Address, _chain common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "isPairedContract", _tokenContract, _chain)
	return *ret0, err
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(bool)
func (_GlobalInbox *GlobalInboxSession) IsPairedContract(_tokenContract common.Address, _chain common.Address) (bool, error) {
	return _GlobalInbox.Contract.IsPairedContract(&_GlobalInbox.CallOpts, _tokenContract, _chain)
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(bool)
func (_GlobalInbox *GlobalInboxCallerSession) IsPairedContract(_tokenContract common.Address, _chain common.Address) (bool, error) {
	return _GlobalInbox.Contract.IsPairedContract(&_GlobalInbox.CallOpts, _tokenContract, _chain)
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) view returns(address[])
func (_GlobalInbox *GlobalInboxCaller) OwnedERC20s(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "ownedERC20s", _owner)
	return *ret0, err
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) view returns(address[])
func (_GlobalInbox *GlobalInboxSession) OwnedERC20s(_owner common.Address) ([]common.Address, error) {
	return _GlobalInbox.Contract.OwnedERC20s(&_GlobalInbox.CallOpts, _owner)
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) view returns(address[])
func (_GlobalInbox *GlobalInboxCallerSession) OwnedERC20s(_owner common.Address) ([]common.Address, error) {
	return _GlobalInbox.Contract.OwnedERC20s(&_GlobalInbox.CallOpts, _owner)
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) view returns(address[])
func (_GlobalInbox *GlobalInboxCaller) OwnedERC721s(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "ownedERC721s", _owner)
	return *ret0, err
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) view returns(address[])
func (_GlobalInbox *GlobalInboxSession) OwnedERC721s(_owner common.Address) ([]common.Address, error) {
	return _GlobalInbox.Contract.OwnedERC721s(&_GlobalInbox.CallOpts, _owner)
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) view returns(address[])
func (_GlobalInbox *GlobalInboxCallerSession) OwnedERC721s(_owner common.Address) ([]common.Address, error) {
	return _GlobalInbox.Contract.OwnedERC721s(&_GlobalInbox.CallOpts, _owner)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x44c75b7f.
//
// Solidity: function deployL2ContractPair(address chain, bytes contractData) returns()
func (_GlobalInbox *GlobalInboxTransactor) DeployL2ContractPair(opts *bind.TransactOpts, chain common.Address, contractData []byte) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "deployL2ContractPair", chain, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x44c75b7f.
//
// Solidity: function deployL2ContractPair(address chain, bytes contractData) returns()
func (_GlobalInbox *GlobalInboxSession) DeployL2ContractPair(chain common.Address, contractData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DeployL2ContractPair(&_GlobalInbox.TransactOpts, chain, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x44c75b7f.
//
// Solidity: function deployL2ContractPair(address chain, bytes contractData) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) DeployL2ContractPair(chain common.Address, contractData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DeployL2ContractPair(&_GlobalInbox.TransactOpts, chain, contractData)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address chain, address erc20, address to, uint256 value) returns()
func (_GlobalInbox *GlobalInboxTransactor) DepositERC20Message(opts *bind.TransactOpts, chain common.Address, erc20 common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "depositERC20Message", chain, erc20, to, value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address chain, address erc20, address to, uint256 value) returns()
func (_GlobalInbox *GlobalInboxSession) DepositERC20Message(chain common.Address, erc20 common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DepositERC20Message(&_GlobalInbox.TransactOpts, chain, erc20, to, value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address chain, address erc20, address to, uint256 value) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) DepositERC20Message(chain common.Address, erc20 common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DepositERC20Message(&_GlobalInbox.TransactOpts, chain, erc20, to, value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address chain, address erc721, address to, uint256 id) returns()
func (_GlobalInbox *GlobalInboxTransactor) DepositERC721Message(opts *bind.TransactOpts, chain common.Address, erc721 common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "depositERC721Message", chain, erc721, to, id)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address chain, address erc721, address to, uint256 id) returns()
func (_GlobalInbox *GlobalInboxSession) DepositERC721Message(chain common.Address, erc721 common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DepositERC721Message(&_GlobalInbox.TransactOpts, chain, erc721, to, id)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address chain, address erc721, address to, uint256 id) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) DepositERC721Message(chain common.Address, erc721 common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DepositERC721Message(&_GlobalInbox.TransactOpts, chain, erc721, to, id)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address chain, address to) payable returns()
func (_GlobalInbox *GlobalInboxTransactor) DepositEthMessage(opts *bind.TransactOpts, chain common.Address, to common.Address) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "depositEthMessage", chain, to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address chain, address to) payable returns()
func (_GlobalInbox *GlobalInboxSession) DepositEthMessage(chain common.Address, to common.Address) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DepositEthMessage(&_GlobalInbox.TransactOpts, chain, to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address chain, address to) payable returns()
func (_GlobalInbox *GlobalInboxTransactorSession) DepositEthMessage(chain common.Address, to common.Address) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DepositEthMessage(&_GlobalInbox.TransactOpts, chain, to)
}

// DepositPairedERC20Message is a paid mutator transaction binding the contract method 0xcb184d0e.
//
// Solidity: function depositPairedERC20Message(address chain, address erc20, address to, uint256 value) returns()
func (_GlobalInbox *GlobalInboxTransactor) DepositPairedERC20Message(opts *bind.TransactOpts, chain common.Address, erc20 common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "depositPairedERC20Message", chain, erc20, to, value)
}

// DepositPairedERC20Message is a paid mutator transaction binding the contract method 0xcb184d0e.
//
// Solidity: function depositPairedERC20Message(address chain, address erc20, address to, uint256 value) returns()
func (_GlobalInbox *GlobalInboxSession) DepositPairedERC20Message(chain common.Address, erc20 common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DepositPairedERC20Message(&_GlobalInbox.TransactOpts, chain, erc20, to, value)
}

// DepositPairedERC20Message is a paid mutator transaction binding the contract method 0xcb184d0e.
//
// Solidity: function depositPairedERC20Message(address chain, address erc20, address to, uint256 value) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) DepositPairedERC20Message(chain common.Address, erc20 common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DepositPairedERC20Message(&_GlobalInbox.TransactOpts, chain, erc20, to, value)
}

// SendInitializationMessage is a paid mutator transaction binding the contract method 0x5cc96efa.
//
// Solidity: function sendInitializationMessage(bytes messageData) returns()
func (_GlobalInbox *GlobalInboxTransactor) SendInitializationMessage(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "sendInitializationMessage", messageData)
}

// SendInitializationMessage is a paid mutator transaction binding the contract method 0x5cc96efa.
//
// Solidity: function sendInitializationMessage(bytes messageData) returns()
func (_GlobalInbox *GlobalInboxSession) SendInitializationMessage(messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendInitializationMessage(&_GlobalInbox.TransactOpts, messageData)
}

// SendInitializationMessage is a paid mutator transaction binding the contract method 0x5cc96efa.
//
// Solidity: function sendInitializationMessage(bytes messageData) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) SendInitializationMessage(messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendInitializationMessage(&_GlobalInbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0x74c6eccc.
//
// Solidity: function sendL2Message(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxTransactor) SendL2Message(opts *bind.TransactOpts, chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "sendL2Message", chain, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0x74c6eccc.
//
// Solidity: function sendL2Message(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxSession) SendL2Message(chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendL2Message(&_GlobalInbox.TransactOpts, chain, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0x74c6eccc.
//
// Solidity: function sendL2Message(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) SendL2Message(chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendL2Message(&_GlobalInbox.TransactOpts, chain, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0xfbef861b.
//
// Solidity: function sendL2MessageFromOrigin(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "sendL2MessageFromOrigin", chain, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0xfbef861b.
//
// Solidity: function sendL2MessageFromOrigin(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxSession) SendL2MessageFromOrigin(chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendL2MessageFromOrigin(&_GlobalInbox.TransactOpts, chain, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0xfbef861b.
//
// Solidity: function sendL2MessageFromOrigin(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) SendL2MessageFromOrigin(chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendL2MessageFromOrigin(&_GlobalInbox.TransactOpts, chain, messageData)
}

// SendMessages is a paid mutator transaction binding the contract method 0x072fd2bb.
//
// Solidity: function sendMessages(bytes messages, uint256[] messageCounts, bytes32[] nodeHashes) returns()
func (_GlobalInbox *GlobalInboxTransactor) SendMessages(opts *bind.TransactOpts, messages []byte, messageCounts []*big.Int, nodeHashes [][32]byte) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "sendMessages", messages, messageCounts, nodeHashes)
}

// SendMessages is a paid mutator transaction binding the contract method 0x072fd2bb.
//
// Solidity: function sendMessages(bytes messages, uint256[] messageCounts, bytes32[] nodeHashes) returns()
func (_GlobalInbox *GlobalInboxSession) SendMessages(messages []byte, messageCounts []*big.Int, nodeHashes [][32]byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendMessages(&_GlobalInbox.TransactOpts, messages, messageCounts, nodeHashes)
}

// SendMessages is a paid mutator transaction binding the contract method 0x072fd2bb.
//
// Solidity: function sendMessages(bytes messages, uint256[] messageCounts, bytes32[] nodeHashes) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) SendMessages(messages []byte, messageCounts []*big.Int, nodeHashes [][32]byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendMessages(&_GlobalInbox.TransactOpts, messages, messageCounts, nodeHashes)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xd2256c66.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, bytes32 nodeHash, uint256 messageIndex) returns()
func (_GlobalInbox *GlobalInboxTransactor) TransferPayment(opts *bind.TransactOpts, originalOwner common.Address, newOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "transferPayment", originalOwner, newOwner, nodeHash, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xd2256c66.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, bytes32 nodeHash, uint256 messageIndex) returns()
func (_GlobalInbox *GlobalInboxSession) TransferPayment(originalOwner common.Address, newOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.TransferPayment(&_GlobalInbox.TransactOpts, originalOwner, newOwner, nodeHash, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xd2256c66.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, bytes32 nodeHash, uint256 messageIndex) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) TransferPayment(originalOwner common.Address, newOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.TransferPayment(&_GlobalInbox.TransactOpts, originalOwner, newOwner, nodeHash, messageIndex)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalInbox *GlobalInboxTransactor) WithdrawERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "withdrawERC20", _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalInbox *GlobalInboxSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalInbox.Contract.WithdrawERC20(&_GlobalInbox.TransactOpts, _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalInbox.Contract.WithdrawERC20(&_GlobalInbox.TransactOpts, _tokenContract)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalInbox *GlobalInboxTransactor) WithdrawERC721(opts *bind.TransactOpts, _erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "withdrawERC721", _erc721, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalInbox *GlobalInboxSession) WithdrawERC721(_erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.WithdrawERC721(&_GlobalInbox.TransactOpts, _erc721, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) WithdrawERC721(_erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.WithdrawERC721(&_GlobalInbox.TransactOpts, _erc721, _tokenId)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalInbox *GlobalInboxTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalInbox *GlobalInboxSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalInbox.Contract.WithdrawEth(&_GlobalInbox.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalInbox *GlobalInboxTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalInbox.Contract.WithdrawEth(&_GlobalInbox.TransactOpts)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_GlobalInbox *GlobalInboxTransactor) WithdrawPairedERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "withdrawPairedERC20", _tokenContract)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_GlobalInbox *GlobalInboxSession) WithdrawPairedERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalInbox.Contract.WithdrawPairedERC20(&_GlobalInbox.TransactOpts, _tokenContract)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) WithdrawPairedERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalInbox.Contract.WithdrawPairedERC20(&_GlobalInbox.TransactOpts, _tokenContract)
}

// GlobalInboxBuddyContractDeployedIterator is returned from FilterBuddyContractDeployed and is used to iterate over the raw logs and unpacked data for BuddyContractDeployed events raised by the GlobalInbox contract.
type GlobalInboxBuddyContractDeployedIterator struct {
	Event *GlobalInboxBuddyContractDeployed // Event containing the contract specifics and raw log

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
func (it *GlobalInboxBuddyContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalInboxBuddyContractDeployed)
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
		it.Event = new(GlobalInboxBuddyContractDeployed)
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
func (it *GlobalInboxBuddyContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalInboxBuddyContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalInboxBuddyContractDeployed represents a BuddyContractDeployed event raised by the GlobalInbox contract.
type GlobalInboxBuddyContractDeployed struct {
	Sender common.Address
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuddyContractDeployed is a free log retrieval operation binding the contract event 0xa98915d9854858ea787b0abcd4e8e3a96802bc19a25474a8b7017a303628e444.
//
// Solidity: event BuddyContractDeployed(address indexed sender, bytes data)
func (_GlobalInbox *GlobalInboxFilterer) FilterBuddyContractDeployed(opts *bind.FilterOpts, sender []common.Address) (*GlobalInboxBuddyContractDeployedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GlobalInbox.contract.FilterLogs(opts, "BuddyContractDeployed", senderRule)
	if err != nil {
		return nil, err
	}
	return &GlobalInboxBuddyContractDeployedIterator{contract: _GlobalInbox.contract, event: "BuddyContractDeployed", logs: logs, sub: sub}, nil
}

// WatchBuddyContractDeployed is a free log subscription operation binding the contract event 0xa98915d9854858ea787b0abcd4e8e3a96802bc19a25474a8b7017a303628e444.
//
// Solidity: event BuddyContractDeployed(address indexed sender, bytes data)
func (_GlobalInbox *GlobalInboxFilterer) WatchBuddyContractDeployed(opts *bind.WatchOpts, sink chan<- *GlobalInboxBuddyContractDeployed, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GlobalInbox.contract.WatchLogs(opts, "BuddyContractDeployed", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalInboxBuddyContractDeployed)
				if err := _GlobalInbox.contract.UnpackLog(event, "BuddyContractDeployed", log); err != nil {
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

// ParseBuddyContractDeployed is a log parse operation binding the contract event 0xa98915d9854858ea787b0abcd4e8e3a96802bc19a25474a8b7017a303628e444.
//
// Solidity: event BuddyContractDeployed(address indexed sender, bytes data)
func (_GlobalInbox *GlobalInboxFilterer) ParseBuddyContractDeployed(log types.Log) (*GlobalInboxBuddyContractDeployed, error) {
	event := new(GlobalInboxBuddyContractDeployed)
	if err := _GlobalInbox.contract.UnpackLog(event, "BuddyContractDeployed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalInboxMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the GlobalInbox contract.
type GlobalInboxMessageDeliveredIterator struct {
	Event *GlobalInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalInboxMessageDelivered)
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
		it.Event = new(GlobalInboxMessageDelivered)
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
func (it *GlobalInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalInboxMessageDelivered represents a MessageDelivered event raised by the GlobalInbox contract.
type GlobalInboxMessageDelivered struct {
	Chain       common.Address
	Kind        uint8
	Sender      common.Address
	InboxSeqNum *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x35e48d636f39df5c5ca2278452d6d89bf9f07c2ff15f46d08aa402c46638b882.
//
// Solidity: event MessageDelivered(address indexed chain, uint8 indexed kind, address indexed sender, uint256 inboxSeqNum, bytes data)
func (_GlobalInbox *GlobalInboxFilterer) FilterMessageDelivered(opts *bind.FilterOpts, chain []common.Address, kind []uint8, sender []common.Address) (*GlobalInboxMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GlobalInbox.contract.FilterLogs(opts, "MessageDelivered", chainRule, kindRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GlobalInboxMessageDeliveredIterator{contract: _GlobalInbox.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x35e48d636f39df5c5ca2278452d6d89bf9f07c2ff15f46d08aa402c46638b882.
//
// Solidity: event MessageDelivered(address indexed chain, uint8 indexed kind, address indexed sender, uint256 inboxSeqNum, bytes data)
func (_GlobalInbox *GlobalInboxFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalInboxMessageDelivered, chain []common.Address, kind []uint8, sender []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GlobalInbox.contract.WatchLogs(opts, "MessageDelivered", chainRule, kindRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalInboxMessageDelivered)
				if err := _GlobalInbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x35e48d636f39df5c5ca2278452d6d89bf9f07c2ff15f46d08aa402c46638b882.
//
// Solidity: event MessageDelivered(address indexed chain, uint8 indexed kind, address indexed sender, uint256 inboxSeqNum, bytes data)
func (_GlobalInbox *GlobalInboxFilterer) ParseMessageDelivered(log types.Log) (*GlobalInboxMessageDelivered, error) {
	event := new(GlobalInboxMessageDelivered)
	if err := _GlobalInbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalInboxMessageDeliveredFromOriginIterator is returned from FilterMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for MessageDeliveredFromOrigin events raised by the GlobalInbox contract.
type GlobalInboxMessageDeliveredFromOriginIterator struct {
	Event *GlobalInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *GlobalInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalInboxMessageDeliveredFromOrigin)
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
		it.Event = new(GlobalInboxMessageDeliveredFromOrigin)
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
func (it *GlobalInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalInboxMessageDeliveredFromOrigin represents a MessageDeliveredFromOrigin event raised by the GlobalInbox contract.
type GlobalInboxMessageDeliveredFromOrigin struct {
	Chain       common.Address
	Kind        uint8
	Sender      common.Address
	InboxSeqNum *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xe923069519faf69b0726ed766a213f61b6f07f2ecf11d55582cc440d8806b0bc.
//
// Solidity: event MessageDeliveredFromOrigin(address indexed chain, uint8 indexed kind, address indexed sender, uint256 inboxSeqNum)
func (_GlobalInbox *GlobalInboxFilterer) FilterMessageDeliveredFromOrigin(opts *bind.FilterOpts, chain []common.Address, kind []uint8, sender []common.Address) (*GlobalInboxMessageDeliveredFromOriginIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GlobalInbox.contract.FilterLogs(opts, "MessageDeliveredFromOrigin", chainRule, kindRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GlobalInboxMessageDeliveredFromOriginIterator{contract: _GlobalInbox.contract, event: "MessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xe923069519faf69b0726ed766a213f61b6f07f2ecf11d55582cc440d8806b0bc.
//
// Solidity: event MessageDeliveredFromOrigin(address indexed chain, uint8 indexed kind, address indexed sender, uint256 inboxSeqNum)
func (_GlobalInbox *GlobalInboxFilterer) WatchMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *GlobalInboxMessageDeliveredFromOrigin, chain []common.Address, kind []uint8, sender []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GlobalInbox.contract.WatchLogs(opts, "MessageDeliveredFromOrigin", chainRule, kindRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalInboxMessageDeliveredFromOrigin)
				if err := _GlobalInbox.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
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

// ParseMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xe923069519faf69b0726ed766a213f61b6f07f2ecf11d55582cc440d8806b0bc.
//
// Solidity: event MessageDeliveredFromOrigin(address indexed chain, uint8 indexed kind, address indexed sender, uint256 inboxSeqNum)
func (_GlobalInbox *GlobalInboxFilterer) ParseMessageDeliveredFromOrigin(log types.Log) (*GlobalInboxMessageDeliveredFromOrigin, error) {
	event := new(GlobalInboxMessageDeliveredFromOrigin)
	if err := _GlobalInbox.contract.UnpackLog(event, "MessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalInboxPaymentTransferIterator is returned from FilterPaymentTransfer and is used to iterate over the raw logs and unpacked data for PaymentTransfer events raised by the GlobalInbox contract.
type GlobalInboxPaymentTransferIterator struct {
	Event *GlobalInboxPaymentTransfer // Event containing the contract specifics and raw log

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
func (it *GlobalInboxPaymentTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalInboxPaymentTransfer)
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
		it.Event = new(GlobalInboxPaymentTransfer)
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
func (it *GlobalInboxPaymentTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalInboxPaymentTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalInboxPaymentTransfer represents a PaymentTransfer event raised by the GlobalInbox contract.
type GlobalInboxPaymentTransfer struct {
	NodeHash      [32]byte
	MessageIndex  *big.Int
	OriginalOwner common.Address
	PrevOwner     common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentTransfer is a free log retrieval operation binding the contract event 0xb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c64383.
//
// Solidity: event PaymentTransfer(bytes32 nodeHash, uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_GlobalInbox *GlobalInboxFilterer) FilterPaymentTransfer(opts *bind.FilterOpts) (*GlobalInboxPaymentTransferIterator, error) {

	logs, sub, err := _GlobalInbox.contract.FilterLogs(opts, "PaymentTransfer")
	if err != nil {
		return nil, err
	}
	return &GlobalInboxPaymentTransferIterator{contract: _GlobalInbox.contract, event: "PaymentTransfer", logs: logs, sub: sub}, nil
}

// WatchPaymentTransfer is a free log subscription operation binding the contract event 0xb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c64383.
//
// Solidity: event PaymentTransfer(bytes32 nodeHash, uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_GlobalInbox *GlobalInboxFilterer) WatchPaymentTransfer(opts *bind.WatchOpts, sink chan<- *GlobalInboxPaymentTransfer) (event.Subscription, error) {

	logs, sub, err := _GlobalInbox.contract.WatchLogs(opts, "PaymentTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalInboxPaymentTransfer)
				if err := _GlobalInbox.contract.UnpackLog(event, "PaymentTransfer", log); err != nil {
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

// ParsePaymentTransfer is a log parse operation binding the contract event 0xb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c64383.
//
// Solidity: event PaymentTransfer(bytes32 nodeHash, uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_GlobalInbox *GlobalInboxFilterer) ParsePaymentTransfer(log types.Log) (*GlobalInboxPaymentTransfer, error) {
	event := new(GlobalInboxPaymentTransfer)
	if err := _GlobalInbox.contract.UnpackLog(event, "PaymentTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalNFTWalletABI is the input ABI used to generate the binding from.
const GlobalNFTWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC721Tokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasERC721\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC721s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalNFTWalletFuncSigs maps the 4-byte function signature to its string representation.
var GlobalNFTWalletFuncSigs = map[string]string{
	"0758fb0a": "getERC721Tokens(address,address)",
	"45a53f09": "hasERC721(address,address,uint256)",
	"33f2ac42": "ownedERC721s(address)",
	"f3e414f8": "withdrawERC721(address,uint256)",
}

// GlobalNFTWalletBin is the compiled bytecode used for deploying new contracts.
var GlobalNFTWalletBin = "0x608060405234801561001057600080fd5b50610765806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630758fb0a1461005157806333f2ac42146100cf57806345a53f09146100f5578063f3e414f81461013f575b600080fd5b61007f6004803603604081101561006757600080fd5b506001600160a01b038135811691602001351661016d565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100bb5781810151838201526020016100a3565b505050509050019250505060405180910390f35b61007f600480360360208110156100e557600080fd5b50356001600160a01b0316610231565b61012b6004803603606081101561010b57600080fd5b506001600160a01b038135811691602081013590911690604001356102f2565b604080519115158252519081900360200190f35b61016b6004803603604081101561015557600080fd5b506001600160a01b038135169060200135610370565b005b6001600160a01b0380821660009081526020818152604080832093861683529083905290205460609190806101b4575050604080516000815260208101909152905061022b565b8160010160018203815481106101c657fe5b906000526020600020906003020160020180548060200260200160405190810160405280929190818152602001828054801561022157602002820191906000526020600020905b81548152602001906001019080831161020d575b5050505050925050505b92915050565b6001600160a01b0381166000908152602081815260409182902060018101548351818152818402810190930190935260609290918391801561027d578160200160208202803883390190505b50805190915060005b818110156102e85783600101818154811061029d57fe5b600091825260209091206003909102015483516001600160a01b03909116908490839081106102c857fe5b6001600160a01b0390921660209283029190910190910152600101610286565b5090949350505050565b6001600160a01b038083166000908152602081815260408083209387168352908390528120549091908061032b57600092505050610369565b81600101600182038154811061033d57fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b61037b33838361043c565b6103cc576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b15801561042057600080fd5b505af1158015610434573d6000803e3d6000fd5b505050505050565b6001600160a01b038084166000908152602081815260408083209386168352908390528120549091908061047557600092505050610369565b600082600101600183038154811061048957fe5b6000918252602080832088845260016003909302019182019052604090912054909150806104be576000945050505050610369565b600282018054829160018501916000919060001981019081106104dd57fe5b60009182526020808320909101548352820192909252604001902055600282018054600019810190811061050d57fe5b906000526020600020015482600201600183038154811061052a57fe5b60009182526020808320909101929092558781526001840190915260408120556002820180548061055757fe5b6000828152602081208201600019908101919091550190556002820154610694576001840180548491869160009190600019810190811061059457fe5b600091825260208083206003909202909101546001600160a01b0316835282019290925260400190205560018401805460001981019081106105d257fe5b90600052602060002090600302018460010160018503815481106105f257fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b039092169190911781556002808301805461063592840191906106a2565b5050506001600160a01b0387166000908152602085905260408120556001840180548061065e57fe5b60008281526020812060036000199093019283020180546001600160a01b03191681559061068f60028301826106f2565b505090555b506001979650505050505050565b8280548282559060005260206000209081019282156106e25760005260206000209182015b828111156106e25782548255916001019190600101906106c7565b506106ee929150610713565b5090565b50805460008255906000526020600020908101906107109190610713565b50565b61072d91905b808211156106ee5760008155600101610719565b9056fea265627a7a723158201788fcb1ef70eebd0853855ddf16c64cbe7143a9173bc6e7ac1cce537f4dc96b64736f6c634300050f0032"

// DeployGlobalNFTWallet deploys a new Ethereum contract, binding an instance of GlobalNFTWallet to it.
func DeployGlobalNFTWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalNFTWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalNFTWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalNFTWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalNFTWallet{GlobalNFTWalletCaller: GlobalNFTWalletCaller{contract: contract}, GlobalNFTWalletTransactor: GlobalNFTWalletTransactor{contract: contract}, GlobalNFTWalletFilterer: GlobalNFTWalletFilterer{contract: contract}}, nil
}

// GlobalNFTWallet is an auto generated Go binding around an Ethereum contract.
type GlobalNFTWallet struct {
	GlobalNFTWalletCaller     // Read-only binding to the contract
	GlobalNFTWalletTransactor // Write-only binding to the contract
	GlobalNFTWalletFilterer   // Log filterer for contract events
}

// GlobalNFTWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalNFTWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNFTWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalNFTWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNFTWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalNFTWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNFTWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalNFTWalletSession struct {
	Contract     *GlobalNFTWallet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalNFTWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalNFTWalletCallerSession struct {
	Contract *GlobalNFTWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GlobalNFTWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalNFTWalletTransactorSession struct {
	Contract     *GlobalNFTWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GlobalNFTWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalNFTWalletRaw struct {
	Contract *GlobalNFTWallet // Generic contract binding to access the raw methods on
}

// GlobalNFTWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalNFTWalletCallerRaw struct {
	Contract *GlobalNFTWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalNFTWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalNFTWalletTransactorRaw struct {
	Contract *GlobalNFTWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalNFTWallet creates a new instance of GlobalNFTWallet, bound to a specific deployed contract.
func NewGlobalNFTWallet(address common.Address, backend bind.ContractBackend) (*GlobalNFTWallet, error) {
	contract, err := bindGlobalNFTWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalNFTWallet{GlobalNFTWalletCaller: GlobalNFTWalletCaller{contract: contract}, GlobalNFTWalletTransactor: GlobalNFTWalletTransactor{contract: contract}, GlobalNFTWalletFilterer: GlobalNFTWalletFilterer{contract: contract}}, nil
}

// NewGlobalNFTWalletCaller creates a new read-only instance of GlobalNFTWallet, bound to a specific deployed contract.
func NewGlobalNFTWalletCaller(address common.Address, caller bind.ContractCaller) (*GlobalNFTWalletCaller, error) {
	contract, err := bindGlobalNFTWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalNFTWalletCaller{contract: contract}, nil
}

// NewGlobalNFTWalletTransactor creates a new write-only instance of GlobalNFTWallet, bound to a specific deployed contract.
func NewGlobalNFTWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalNFTWalletTransactor, error) {
	contract, err := bindGlobalNFTWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalNFTWalletTransactor{contract: contract}, nil
}

// NewGlobalNFTWalletFilterer creates a new log filterer instance of GlobalNFTWallet, bound to a specific deployed contract.
func NewGlobalNFTWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalNFTWalletFilterer, error) {
	contract, err := bindGlobalNFTWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalNFTWalletFilterer{contract: contract}, nil
}

// bindGlobalNFTWallet binds a generic wrapper to an already deployed contract.
func bindGlobalNFTWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalNFTWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalNFTWallet *GlobalNFTWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalNFTWallet.Contract.GlobalNFTWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalNFTWallet *GlobalNFTWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.GlobalNFTWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalNFTWallet *GlobalNFTWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.GlobalNFTWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalNFTWallet *GlobalNFTWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalNFTWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalNFTWallet *GlobalNFTWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalNFTWallet *GlobalNFTWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.contract.Transact(opts, method, params...)
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) view returns(uint256[])
func (_GlobalNFTWallet *GlobalNFTWalletCaller) GetERC721Tokens(opts *bind.CallOpts, _erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _GlobalNFTWallet.contract.Call(opts, out, "getERC721Tokens", _erc721, _owner)
	return *ret0, err
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) view returns(uint256[])
func (_GlobalNFTWallet *GlobalNFTWalletSession) GetERC721Tokens(_erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	return _GlobalNFTWallet.Contract.GetERC721Tokens(&_GlobalNFTWallet.CallOpts, _erc721, _owner)
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) view returns(uint256[])
func (_GlobalNFTWallet *GlobalNFTWalletCallerSession) GetERC721Tokens(_erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	return _GlobalNFTWallet.Contract.GetERC721Tokens(&_GlobalNFTWallet.CallOpts, _erc721, _owner)
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) view returns(bool)
func (_GlobalNFTWallet *GlobalNFTWalletCaller) HasERC721(opts *bind.CallOpts, _erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalNFTWallet.contract.Call(opts, out, "hasERC721", _erc721, _owner, _tokenId)
	return *ret0, err
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) view returns(bool)
func (_GlobalNFTWallet *GlobalNFTWalletSession) HasERC721(_erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalNFTWallet.Contract.HasERC721(&_GlobalNFTWallet.CallOpts, _erc721, _owner, _tokenId)
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) view returns(bool)
func (_GlobalNFTWallet *GlobalNFTWalletCallerSession) HasERC721(_erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalNFTWallet.Contract.HasERC721(&_GlobalNFTWallet.CallOpts, _erc721, _owner, _tokenId)
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) view returns(address[])
func (_GlobalNFTWallet *GlobalNFTWalletCaller) OwnedERC721s(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GlobalNFTWallet.contract.Call(opts, out, "ownedERC721s", _owner)
	return *ret0, err
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) view returns(address[])
func (_GlobalNFTWallet *GlobalNFTWalletSession) OwnedERC721s(_owner common.Address) ([]common.Address, error) {
	return _GlobalNFTWallet.Contract.OwnedERC721s(&_GlobalNFTWallet.CallOpts, _owner)
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) view returns(address[])
func (_GlobalNFTWallet *GlobalNFTWalletCallerSession) OwnedERC721s(_owner common.Address) ([]common.Address, error) {
	return _GlobalNFTWallet.Contract.OwnedERC721s(&_GlobalNFTWallet.CallOpts, _owner)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalNFTWallet *GlobalNFTWalletTransactor) WithdrawERC721(opts *bind.TransactOpts, _erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalNFTWallet.contract.Transact(opts, "withdrawERC721", _erc721, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalNFTWallet *GlobalNFTWalletSession) WithdrawERC721(_erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.WithdrawERC721(&_GlobalNFTWallet.TransactOpts, _erc721, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalNFTWallet *GlobalNFTWalletTransactorSession) WithdrawERC721(_erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.WithdrawERC721(&_GlobalNFTWallet.TransactOpts, _erc721, _tokenId)
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC165.contract.Call(opts, out, "supportsInterface", interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceID)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ABI is the input ABI used to generate the binding from.
const IERC721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IERC721FuncSigs maps the 4-byte function signature to its string representation.
var IERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_IERC721 *IERC721Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_IERC721 *IERC721CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_IERC721 *IERC721Session) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_IERC721 *IERC721CallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_IERC721 *IERC721Session) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_IERC721 *IERC721CallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, _approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721Session) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721TransactorSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, _approved, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data) payable returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data) payable returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, _from, _to, _tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data) payable returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, _from, _to, _tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721Session) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, _from, _to, _tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*IERC721TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PairedErc20ABI is the input ABI used to generate the binding from.
const PairedErc20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PairedErc20FuncSigs maps the 4-byte function signature to its string representation.
var PairedErc20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"9dc29fac": "burn(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// PairedErc20 is an auto generated Go binding around an Ethereum contract.
type PairedErc20 struct {
	PairedErc20Caller     // Read-only binding to the contract
	PairedErc20Transactor // Write-only binding to the contract
	PairedErc20Filterer   // Log filterer for contract events
}

// PairedErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type PairedErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PairedErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairedErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairedErc20Session struct {
	Contract     *PairedErc20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairedErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairedErc20CallerSession struct {
	Contract *PairedErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PairedErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairedErc20TransactorSession struct {
	Contract     *PairedErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PairedErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type PairedErc20Raw struct {
	Contract *PairedErc20 // Generic contract binding to access the raw methods on
}

// PairedErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairedErc20CallerRaw struct {
	Contract *PairedErc20Caller // Generic read-only contract binding to access the raw methods on
}

// PairedErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairedErc20TransactorRaw struct {
	Contract *PairedErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPairedErc20 creates a new instance of PairedErc20, bound to a specific deployed contract.
func NewPairedErc20(address common.Address, backend bind.ContractBackend) (*PairedErc20, error) {
	contract, err := bindPairedErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairedErc20{PairedErc20Caller: PairedErc20Caller{contract: contract}, PairedErc20Transactor: PairedErc20Transactor{contract: contract}, PairedErc20Filterer: PairedErc20Filterer{contract: contract}}, nil
}

// NewPairedErc20Caller creates a new read-only instance of PairedErc20, bound to a specific deployed contract.
func NewPairedErc20Caller(address common.Address, caller bind.ContractCaller) (*PairedErc20Caller, error) {
	contract, err := bindPairedErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairedErc20Caller{contract: contract}, nil
}

// NewPairedErc20Transactor creates a new write-only instance of PairedErc20, bound to a specific deployed contract.
func NewPairedErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*PairedErc20Transactor, error) {
	contract, err := bindPairedErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairedErc20Transactor{contract: contract}, nil
}

// NewPairedErc20Filterer creates a new log filterer instance of PairedErc20, bound to a specific deployed contract.
func NewPairedErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*PairedErc20Filterer, error) {
	contract, err := bindPairedErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairedErc20Filterer{contract: contract}, nil
}

// bindPairedErc20 binds a generic wrapper to an already deployed contract.
func bindPairedErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairedErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairedErc20 *PairedErc20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairedErc20.Contract.PairedErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairedErc20 *PairedErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairedErc20.Contract.PairedErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairedErc20 *PairedErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairedErc20.Contract.PairedErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairedErc20 *PairedErc20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairedErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairedErc20 *PairedErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairedErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairedErc20 *PairedErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairedErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PairedErc20 *PairedErc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PairedErc20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PairedErc20 *PairedErc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PairedErc20.Contract.Allowance(&_PairedErc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PairedErc20 *PairedErc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PairedErc20.Contract.Allowance(&_PairedErc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PairedErc20 *PairedErc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PairedErc20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PairedErc20 *PairedErc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _PairedErc20.Contract.BalanceOf(&_PairedErc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PairedErc20 *PairedErc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PairedErc20.Contract.BalanceOf(&_PairedErc20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PairedErc20 *PairedErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PairedErc20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PairedErc20 *PairedErc20Session) TotalSupply() (*big.Int, error) {
	return _PairedErc20.Contract.TotalSupply(&_PairedErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PairedErc20 *PairedErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _PairedErc20.Contract.TotalSupply(&_PairedErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.Approve(&_PairedErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.Approve(&_PairedErc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_PairedErc20 *PairedErc20Transactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_PairedErc20 *PairedErc20Session) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.Burn(&_PairedErc20.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_PairedErc20 *PairedErc20TransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.Burn(&_PairedErc20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_PairedErc20 *PairedErc20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_PairedErc20 *PairedErc20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.Mint(&_PairedErc20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_PairedErc20 *PairedErc20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.Mint(&_PairedErc20.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.Transfer(&_PairedErc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.Transfer(&_PairedErc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.TransferFrom(&_PairedErc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PairedErc20 *PairedErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PairedErc20.Contract.TransferFrom(&_PairedErc20.TransactOpts, sender, recipient, amount)
}

// PairedErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PairedErc20 contract.
type PairedErc20ApprovalIterator struct {
	Event *PairedErc20Approval // Event containing the contract specifics and raw log

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
func (it *PairedErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairedErc20Approval)
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
		it.Event = new(PairedErc20Approval)
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
func (it *PairedErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairedErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairedErc20Approval represents a Approval event raised by the PairedErc20 contract.
type PairedErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PairedErc20 *PairedErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PairedErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PairedErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PairedErc20ApprovalIterator{contract: _PairedErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PairedErc20 *PairedErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PairedErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PairedErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairedErc20Approval)
				if err := _PairedErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PairedErc20 *PairedErc20Filterer) ParseApproval(log types.Log) (*PairedErc20Approval, error) {
	event := new(PairedErc20Approval)
	if err := _PairedErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PairedErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PairedErc20 contract.
type PairedErc20TransferIterator struct {
	Event *PairedErc20Transfer // Event containing the contract specifics and raw log

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
func (it *PairedErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairedErc20Transfer)
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
		it.Event = new(PairedErc20Transfer)
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
func (it *PairedErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairedErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairedErc20Transfer represents a Transfer event raised by the PairedErc20 contract.
type PairedErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PairedErc20 *PairedErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PairedErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PairedErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairedErc20TransferIterator{contract: _PairedErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PairedErc20 *PairedErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PairedErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PairedErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairedErc20Transfer)
				if err := _PairedErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PairedErc20 *PairedErc20Filterer) ParseTransfer(log types.Log) (*PairedErc20Transfer, error) {
	event := new(PairedErc20Transfer)
	if err := _PairedErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PairedFTWalletABI is the input ABI used to generate the binding from.
const PairedFTWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"}],\"name\":\"isPairedContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawPairedERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PairedFTWalletFuncSigs maps the 4-byte function signature to its string representation.
var PairedFTWalletFuncSigs = map[string]string{
	"c3a8962c": "getERC20Balance(address,address)",
	"659e42cd": "isPairedContract(address,address)",
	"21e2f2fa": "withdrawPairedERC20(address)",
}

// PairedFTWalletBin is the compiled bytecode used for deploying new contracts.
var PairedFTWalletBin = "0x608060405234801561001057600080fd5b5061043b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806321e2f2fa14610046578063659e42cd1461006e578063c3a8962c146100b0575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b03166100f0565b005b61009c6004803603604081101561008457600080fd5b506001600160a01b03813581169160200135166101ae565b604080519115158252519081900360200190f35b6100de600480360360408110156100c657600080fd5b506001600160a01b03813581169160200135166101de565b60408051918252519081900360200190f35b60006100fc82336101de565b9050610109338383610245565b6101445760405162461bcd60e51b815260040180806020018281038252602e8152602001806103d9602e913960400191505060405180910390fd5b604080516340c10f1960e01b81523360048201526024810183905290516001600160a01b038416916340c10f1991604480830192600092919082900301818387803b15801561019257600080fd5b505af11580156101a6573d6000803e3d6000fd5b505050505050565b6001600160a01b0380831660009081526001602090815260408083209385168352929052205460ff165b92915050565b6001600160a01b0380821660009081526020818152604080832093861683529083905281205490919080610217576000925050506101d8565b81600101600182038154811061022957fe5b9060005260206000209060020201600101549250505092915050565b600081610254575060016103d1565b6001600160a01b038085166000908152602081815260408083209387168352908390529020548061028a576000925050506103d1565b600082600101600183038154811061029e57fe5b9060005260206000209060020201905080600101548511156102c657600093505050506103d1565b600181018054869003908190556103c957600183018054839185916000919060001981019081106102f357fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061033157fe5b906000526020600020906002020183600101600184038154811061035157fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061039f57fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b600193505050505b939250505056fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a723158206f94ea93f77e3ac8bb387b999dda92b37f521ade7bc8691586b57017f226cc4164736f6c634300050f0032"

// DeployPairedFTWallet deploys a new Ethereum contract, binding an instance of PairedFTWallet to it.
func DeployPairedFTWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PairedFTWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(PairedFTWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairedFTWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PairedFTWallet{PairedFTWalletCaller: PairedFTWalletCaller{contract: contract}, PairedFTWalletTransactor: PairedFTWalletTransactor{contract: contract}, PairedFTWalletFilterer: PairedFTWalletFilterer{contract: contract}}, nil
}

// PairedFTWallet is an auto generated Go binding around an Ethereum contract.
type PairedFTWallet struct {
	PairedFTWalletCaller     // Read-only binding to the contract
	PairedFTWalletTransactor // Write-only binding to the contract
	PairedFTWalletFilterer   // Log filterer for contract events
}

// PairedFTWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairedFTWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedFTWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairedFTWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedFTWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairedFTWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedFTWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairedFTWalletSession struct {
	Contract     *PairedFTWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairedFTWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairedFTWalletCallerSession struct {
	Contract *PairedFTWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PairedFTWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairedFTWalletTransactorSession struct {
	Contract     *PairedFTWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PairedFTWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairedFTWalletRaw struct {
	Contract *PairedFTWallet // Generic contract binding to access the raw methods on
}

// PairedFTWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairedFTWalletCallerRaw struct {
	Contract *PairedFTWalletCaller // Generic read-only contract binding to access the raw methods on
}

// PairedFTWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairedFTWalletTransactorRaw struct {
	Contract *PairedFTWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairedFTWallet creates a new instance of PairedFTWallet, bound to a specific deployed contract.
func NewPairedFTWallet(address common.Address, backend bind.ContractBackend) (*PairedFTWallet, error) {
	contract, err := bindPairedFTWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairedFTWallet{PairedFTWalletCaller: PairedFTWalletCaller{contract: contract}, PairedFTWalletTransactor: PairedFTWalletTransactor{contract: contract}, PairedFTWalletFilterer: PairedFTWalletFilterer{contract: contract}}, nil
}

// NewPairedFTWalletCaller creates a new read-only instance of PairedFTWallet, bound to a specific deployed contract.
func NewPairedFTWalletCaller(address common.Address, caller bind.ContractCaller) (*PairedFTWalletCaller, error) {
	contract, err := bindPairedFTWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairedFTWalletCaller{contract: contract}, nil
}

// NewPairedFTWalletTransactor creates a new write-only instance of PairedFTWallet, bound to a specific deployed contract.
func NewPairedFTWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*PairedFTWalletTransactor, error) {
	contract, err := bindPairedFTWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairedFTWalletTransactor{contract: contract}, nil
}

// NewPairedFTWalletFilterer creates a new log filterer instance of PairedFTWallet, bound to a specific deployed contract.
func NewPairedFTWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*PairedFTWalletFilterer, error) {
	contract, err := bindPairedFTWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairedFTWalletFilterer{contract: contract}, nil
}

// bindPairedFTWallet binds a generic wrapper to an already deployed contract.
func bindPairedFTWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairedFTWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairedFTWallet *PairedFTWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairedFTWallet.Contract.PairedFTWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairedFTWallet *PairedFTWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairedFTWallet.Contract.PairedFTWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairedFTWallet *PairedFTWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairedFTWallet.Contract.PairedFTWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairedFTWallet *PairedFTWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairedFTWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairedFTWallet *PairedFTWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairedFTWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairedFTWallet *PairedFTWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairedFTWallet.Contract.contract.Transact(opts, method, params...)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_PairedFTWallet *PairedFTWalletCaller) GetERC20Balance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PairedFTWallet.contract.Call(opts, out, "getERC20Balance", _tokenContract, _owner)
	return *ret0, err
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_PairedFTWallet *PairedFTWalletSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _PairedFTWallet.Contract.GetERC20Balance(&_PairedFTWallet.CallOpts, _tokenContract, _owner)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_PairedFTWallet *PairedFTWalletCallerSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _PairedFTWallet.Contract.GetERC20Balance(&_PairedFTWallet.CallOpts, _tokenContract, _owner)
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(bool)
func (_PairedFTWallet *PairedFTWalletCaller) IsPairedContract(opts *bind.CallOpts, _tokenContract common.Address, _chain common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairedFTWallet.contract.Call(opts, out, "isPairedContract", _tokenContract, _chain)
	return *ret0, err
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(bool)
func (_PairedFTWallet *PairedFTWalletSession) IsPairedContract(_tokenContract common.Address, _chain common.Address) (bool, error) {
	return _PairedFTWallet.Contract.IsPairedContract(&_PairedFTWallet.CallOpts, _tokenContract, _chain)
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(bool)
func (_PairedFTWallet *PairedFTWalletCallerSession) IsPairedContract(_tokenContract common.Address, _chain common.Address) (bool, error) {
	return _PairedFTWallet.Contract.IsPairedContract(&_PairedFTWallet.CallOpts, _tokenContract, _chain)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_PairedFTWallet *PairedFTWalletTransactor) WithdrawPairedERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _PairedFTWallet.contract.Transact(opts, "withdrawPairedERC20", _tokenContract)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_PairedFTWallet *PairedFTWalletSession) WithdrawPairedERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _PairedFTWallet.Contract.WithdrawPairedERC20(&_PairedFTWallet.TransactOpts, _tokenContract)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_PairedFTWallet *PairedFTWalletTransactorSession) WithdrawPairedERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _PairedFTWallet.Contract.WithdrawPairedERC20(&_PairedFTWallet.TransactOpts, _tokenContract)
}

// PaymentRecordsABI is the input ABI used to generate the binding from.
const PaymentRecordsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"PaymentTransfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"getPaymentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"transferPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PaymentRecordsFuncSigs maps the 4-byte function signature to its string representation.
var PaymentRecordsFuncSigs = map[string]string{
	"bd4fbb36": "getPaymentOwner(address,bytes32,uint256)",
	"d2256c66": "transferPayment(address,address,bytes32,uint256)",
}

// PaymentRecordsBin is the compiled bytecode used for deploying new contracts.
var PaymentRecordsBin = "0x608060405234801561001057600080fd5b50610295806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063bd4fbb361461003b578063d2256c6614610089575b600080fd5b61006d6004803603606081101561005157600080fd5b506001600160a01b0381351690602081013590604001356100c7565b604080516001600160a01b039092168252519081900360200190f35b6100c56004803603608081101561009f57600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135610139565b005b604080516020808201859052818301849052606086811b6bffffffffffffffffffffffff1916908301528251808303605401815260749092018352815191810191909120600090815290819052908120546001600160a01b03168061012f5784915050610132565b90505b9392505050565b60006101468584846100c7565b9050336001600160a01b0382161461019e576040805162461bcd60e51b815260206004820152601660248201527526bab9ba103132903830bcb6b2b73a1037bbb732b91760511b604482015290519081900360640190fd5b604080516020808201869052818301859052606088811b6bffffffffffffffffffffffff19169083015282518083036054018152607483018085528151918301919091206000908152918290529083902080546001600160a01b0319166001600160a01b03898116918217909255918790526094830186905280891660b4840152841660d483015260f482015290517fb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c6438391610114908290030190a1505050505056fea265627a7a72315820f7511d54fbfe75c5afeec71473690318b4c4a25fa894df895ea725726535837164736f6c634300050f0032"

// DeployPaymentRecords deploys a new Ethereum contract, binding an instance of PaymentRecords to it.
func DeployPaymentRecords(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PaymentRecords, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentRecordsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaymentRecordsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaymentRecords{PaymentRecordsCaller: PaymentRecordsCaller{contract: contract}, PaymentRecordsTransactor: PaymentRecordsTransactor{contract: contract}, PaymentRecordsFilterer: PaymentRecordsFilterer{contract: contract}}, nil
}

// PaymentRecords is an auto generated Go binding around an Ethereum contract.
type PaymentRecords struct {
	PaymentRecordsCaller     // Read-only binding to the contract
	PaymentRecordsTransactor // Write-only binding to the contract
	PaymentRecordsFilterer   // Log filterer for contract events
}

// PaymentRecordsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentRecordsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentRecordsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentRecordsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentRecordsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentRecordsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentRecordsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentRecordsSession struct {
	Contract     *PaymentRecords   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentRecordsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentRecordsCallerSession struct {
	Contract *PaymentRecordsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PaymentRecordsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentRecordsTransactorSession struct {
	Contract     *PaymentRecordsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PaymentRecordsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentRecordsRaw struct {
	Contract *PaymentRecords // Generic contract binding to access the raw methods on
}

// PaymentRecordsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentRecordsCallerRaw struct {
	Contract *PaymentRecordsCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentRecordsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentRecordsTransactorRaw struct {
	Contract *PaymentRecordsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentRecords creates a new instance of PaymentRecords, bound to a specific deployed contract.
func NewPaymentRecords(address common.Address, backend bind.ContractBackend) (*PaymentRecords, error) {
	contract, err := bindPaymentRecords(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentRecords{PaymentRecordsCaller: PaymentRecordsCaller{contract: contract}, PaymentRecordsTransactor: PaymentRecordsTransactor{contract: contract}, PaymentRecordsFilterer: PaymentRecordsFilterer{contract: contract}}, nil
}

// NewPaymentRecordsCaller creates a new read-only instance of PaymentRecords, bound to a specific deployed contract.
func NewPaymentRecordsCaller(address common.Address, caller bind.ContractCaller) (*PaymentRecordsCaller, error) {
	contract, err := bindPaymentRecords(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentRecordsCaller{contract: contract}, nil
}

// NewPaymentRecordsTransactor creates a new write-only instance of PaymentRecords, bound to a specific deployed contract.
func NewPaymentRecordsTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentRecordsTransactor, error) {
	contract, err := bindPaymentRecords(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentRecordsTransactor{contract: contract}, nil
}

// NewPaymentRecordsFilterer creates a new log filterer instance of PaymentRecords, bound to a specific deployed contract.
func NewPaymentRecordsFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentRecordsFilterer, error) {
	contract, err := bindPaymentRecords(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentRecordsFilterer{contract: contract}, nil
}

// bindPaymentRecords binds a generic wrapper to an already deployed contract.
func bindPaymentRecords(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentRecordsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentRecords *PaymentRecordsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentRecords.Contract.PaymentRecordsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentRecords *PaymentRecordsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentRecords.Contract.PaymentRecordsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentRecords *PaymentRecordsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentRecords.Contract.PaymentRecordsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentRecords *PaymentRecordsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentRecords.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentRecords *PaymentRecordsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentRecords.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentRecords *PaymentRecordsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentRecords.Contract.contract.Transact(opts, method, params...)
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0xbd4fbb36.
//
// Solidity: function getPaymentOwner(address originalOwner, bytes32 nodeHash, uint256 messageIndex) view returns(address)
func (_PaymentRecords *PaymentRecordsCaller) GetPaymentOwner(opts *bind.CallOpts, originalOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentRecords.contract.Call(opts, out, "getPaymentOwner", originalOwner, nodeHash, messageIndex)
	return *ret0, err
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0xbd4fbb36.
//
// Solidity: function getPaymentOwner(address originalOwner, bytes32 nodeHash, uint256 messageIndex) view returns(address)
func (_PaymentRecords *PaymentRecordsSession) GetPaymentOwner(originalOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (common.Address, error) {
	return _PaymentRecords.Contract.GetPaymentOwner(&_PaymentRecords.CallOpts, originalOwner, nodeHash, messageIndex)
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0xbd4fbb36.
//
// Solidity: function getPaymentOwner(address originalOwner, bytes32 nodeHash, uint256 messageIndex) view returns(address)
func (_PaymentRecords *PaymentRecordsCallerSession) GetPaymentOwner(originalOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (common.Address, error) {
	return _PaymentRecords.Contract.GetPaymentOwner(&_PaymentRecords.CallOpts, originalOwner, nodeHash, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xd2256c66.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, bytes32 nodeHash, uint256 messageIndex) returns()
func (_PaymentRecords *PaymentRecordsTransactor) TransferPayment(opts *bind.TransactOpts, originalOwner common.Address, newOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (*types.Transaction, error) {
	return _PaymentRecords.contract.Transact(opts, "transferPayment", originalOwner, newOwner, nodeHash, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xd2256c66.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, bytes32 nodeHash, uint256 messageIndex) returns()
func (_PaymentRecords *PaymentRecordsSession) TransferPayment(originalOwner common.Address, newOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (*types.Transaction, error) {
	return _PaymentRecords.Contract.TransferPayment(&_PaymentRecords.TransactOpts, originalOwner, newOwner, nodeHash, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xd2256c66.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, bytes32 nodeHash, uint256 messageIndex) returns()
func (_PaymentRecords *PaymentRecordsTransactorSession) TransferPayment(originalOwner common.Address, newOwner common.Address, nodeHash [32]byte, messageIndex *big.Int) (*types.Transaction, error) {
	return _PaymentRecords.Contract.TransferPayment(&_PaymentRecords.TransactOpts, originalOwner, newOwner, nodeHash, messageIndex)
}

// PaymentRecordsPaymentTransferIterator is returned from FilterPaymentTransfer and is used to iterate over the raw logs and unpacked data for PaymentTransfer events raised by the PaymentRecords contract.
type PaymentRecordsPaymentTransferIterator struct {
	Event *PaymentRecordsPaymentTransfer // Event containing the contract specifics and raw log

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
func (it *PaymentRecordsPaymentTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentRecordsPaymentTransfer)
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
		it.Event = new(PaymentRecordsPaymentTransfer)
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
func (it *PaymentRecordsPaymentTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentRecordsPaymentTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentRecordsPaymentTransfer represents a PaymentTransfer event raised by the PaymentRecords contract.
type PaymentRecordsPaymentTransfer struct {
	NodeHash      [32]byte
	MessageIndex  *big.Int
	OriginalOwner common.Address
	PrevOwner     common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentTransfer is a free log retrieval operation binding the contract event 0xb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c64383.
//
// Solidity: event PaymentTransfer(bytes32 nodeHash, uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_PaymentRecords *PaymentRecordsFilterer) FilterPaymentTransfer(opts *bind.FilterOpts) (*PaymentRecordsPaymentTransferIterator, error) {

	logs, sub, err := _PaymentRecords.contract.FilterLogs(opts, "PaymentTransfer")
	if err != nil {
		return nil, err
	}
	return &PaymentRecordsPaymentTransferIterator{contract: _PaymentRecords.contract, event: "PaymentTransfer", logs: logs, sub: sub}, nil
}

// WatchPaymentTransfer is a free log subscription operation binding the contract event 0xb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c64383.
//
// Solidity: event PaymentTransfer(bytes32 nodeHash, uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_PaymentRecords *PaymentRecordsFilterer) WatchPaymentTransfer(opts *bind.WatchOpts, sink chan<- *PaymentRecordsPaymentTransfer) (event.Subscription, error) {

	logs, sub, err := _PaymentRecords.contract.WatchLogs(opts, "PaymentTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentRecordsPaymentTransfer)
				if err := _PaymentRecords.contract.UnpackLog(event, "PaymentTransfer", log); err != nil {
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

// ParsePaymentTransfer is a log parse operation binding the contract event 0xb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c64383.
//
// Solidity: event PaymentTransfer(bytes32 nodeHash, uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_PaymentRecords *PaymentRecordsFilterer) ParsePaymentTransfer(log types.Log) (*PaymentRecordsPaymentTransfer, error) {
	event := new(PaymentRecordsPaymentTransfer)
	if err := _PaymentRecords.contract.UnpackLog(event, "PaymentTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
