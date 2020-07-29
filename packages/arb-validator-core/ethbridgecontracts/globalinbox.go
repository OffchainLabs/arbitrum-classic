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
var GlobalFTWalletBin = "0x608060405234801561001057600080fd5b50610516806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636e2b89c514610046578063c3a8962c146100bc578063f4f3b200146100fc575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b0316610124565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100a8578181015183820152602001610090565b505050509050019250505060405180910390f35b6100ea600480360360408110156100d257600080fd5b506001600160a01b03813581169160200135166101e5565b60408051918252519081900360200190f35b6101226004803603602081101561011257600080fd5b50356001600160a01b031661024d565b005b6001600160a01b03811660009081526020818152604091829020600181015483518181528184028101909301909352606092909183918015610170578160200160208202803883390190505b50805190915060005b818110156101db5783600101818154811061019057fe5b600091825260209091206002909102015483516001600160a01b03909116908490839081106101bb57fe5b6001600160a01b0390921660209283029190910190910152600101610179565b5090949350505050565b6001600160a01b038082166000908152602081815260408083209386168352908390528120549091908061021e57600092505050610247565b81600101600182038154811061023057fe5b906000526020600020906002020160010154925050505b92915050565b600061025982336101e5565b9050610266338383610320565b6102a15760405162461bcd60e51b815260040180806020018281038252602e8152602001806104b4602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b1580156102f057600080fd5b505af1158015610304573d6000803e3d6000fd5b505050506040513d602081101561031a57600080fd5b50505050565b60008161032f575060016104ac565b6001600160a01b0380851660009081526020818152604080832093871683529083905290205480610365576000925050506104ac565b600082600101600183038154811061037957fe5b9060005260206000209060020201905080600101548511156103a157600093505050506104ac565b600181018054869003908190556104a457600183018054839185916000919060001981019081106103ce57fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061040c57fe5b906000526020600020906002020183600101600184038154811061042c57fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061047a57fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b600193505050505b939250505056fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a72315820ec76a4faf6e0f0c2e8fa97e9e80500c6843ea76f90e9e395913bd09d2c6975ad64736f6c634300050f0032"

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
const GlobalInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BuddyContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"}],\"name\":\"MessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"PaymentTransfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"depositERC721Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositPairedERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC721Tokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getInbox\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getPairedERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"getPaymentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasERC721\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBuddyContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC20s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC721s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendInitializationMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2BuddyDeploy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeHashes\",\"type\":\"bytes32[]\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"transferPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawPairedERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalInboxFuncSigs maps the 4-byte function signature to its string representation.
var GlobalInboxFuncSigs = map[string]string{
	"bca22b76": "depositERC20Message(address,address,address,uint256)",
	"8b7010aa": "depositERC721Message(address,address,address,uint256)",
	"5bd21290": "depositEthMessage(address,address)",
	"cb184d0e": "depositPairedERC20Message(address,address,address,uint256)",
	"c3a8962c": "getERC20Balance(address,address)",
	"0758fb0a": "getERC721Tokens(address,address)",
	"4d2301cc": "getEthBalance(address)",
	"02201681": "getInbox(address)",
	"2f5bf211": "getPairedERC20Balance(address,address)",
	"bd4fbb36": "getPaymentOwner(address,bytes32,uint256)",
	"45a53f09": "hasERC721(address,address,uint256)",
	"b4346fec": "isBuddyContract(address)",
	"6e2b89c5": "ownedERC20s(address)",
	"33f2ac42": "ownedERC721s(address)",
	"5cc96efa": "sendInitializationMessage(bytes)",
	"0a1a4af6": "sendL2BuddyDeploy(address,bytes)",
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
var GlobalInboxBin = "0x608060405234801561001057600080fd5b506129a0806100206000396000f3fe60806040526004361061014b5760003560e01c806374c6eccc116100b6578063c3a8962c1161006f578063c3a8962c146107a8578063cb184d0e146107e3578063d2256c661461082c578063f3e414f814610875578063f4f3b200146108ae578063fbef861b146108e15761014b565b806374c6eccc146105e85780638b7010aa14610673578063a0ef91df146106bc578063b4346fec146106d1578063bca22b7614610704578063bd4fbb361461074d5761014b565b806333f2ac421161010857806333f2ac421461044f57806345a53f09146104825780634d2301cc146104d95780635bd212901461050c5780635cc96efa1461053a5780636e2b89c5146105b55761014b565b80630220168114610150578063072fd2bb1461019c5780630758fb0a146102b95780630a1a4af61461034457806321e2f2fa146103cf5780632f5bf21114610402575b600080fd5b34801561015c57600080fd5b506101836004803603602081101561017357600080fd5b50356001600160a01b031661096c565b6040805192835260208301919091528051918290030190f35b3480156101a857600080fd5b506102b7600480360360608110156101bf57600080fd5b810190602081018135600160201b8111156101d957600080fd5b8201836020820111156101eb57600080fd5b803590602001918460018302840111600160201b8311171561020c57600080fd5b919390929091602081019035600160201b81111561022957600080fd5b82018360208201111561023b57600080fd5b803590602001918460208302840111600160201b8311171561025c57600080fd5b919390929091602081019035600160201b81111561027957600080fd5b82018360208201111561028b57600080fd5b803590602001918460208302840111600160201b831117156102ac57600080fd5b509092509050610992565b005b3480156102c557600080fd5b506102f4600480360360408110156102dc57600080fd5b506001600160a01b0381358116916020013516610a5a565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610330578181015183820152602001610318565b505050509050019250505060405180910390f35b34801561035057600080fd5b506102b76004803603604081101561036757600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561039157600080fd5b8201836020820111156103a357600080fd5b803590602001918460018302840111600160201b831117156103c457600080fd5b509092509050610b20565b3480156103db57600080fd5b506102b7600480360360208110156103f257600080fd5b50356001600160a01b0316610b68565b34801561040e57600080fd5b5061043d6004803603604081101561042557600080fd5b506001600160a01b0381358116916020013516610c1e565b60408051918252519081900360200190f35b34801561045b57600080fd5b506102f46004803603602081101561047257600080fd5b50356001600160a01b0316610c87565b34801561048e57600080fd5b506104c5600480360360608110156104a557600080fd5b506001600160a01b03813581169160208101359091169060400135610d4a565b604080519115158252519081900360200190f35b3480156104e557600080fd5b5061043d600480360360208110156104fc57600080fd5b50356001600160a01b0316610dca565b6102b76004803603604081101561052257600080fd5b506001600160a01b0381358116916020013516610de5565b34801561054657600080fd5b506102b76004803603602081101561055d57600080fd5b810190602081018135600160201b81111561057757600080fd5b82018360208201111561058957600080fd5b803590602001918460018302840111600160201b831117156105aa57600080fd5b509092509050610e2b565b3480156105c157600080fd5b506102f4600480360360208110156105d857600080fd5b50356001600160a01b0316610e6e565b3480156105f457600080fd5b506102b76004803603604081101561060b57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561063557600080fd5b82018360208201111561064757600080fd5b803590602001918460018302840111600160201b8311171561066857600080fd5b509092509050610f27565b34801561067f57600080fd5b506102b76004803603608081101561069657600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610f6a565b3480156106c857600080fd5b506102b7610fd1565b3480156106dd57600080fd5b506104c5600480360360208110156106f457600080fd5b50356001600160a01b031661101c565b34801561071057600080fd5b506102b76004803603608081101561072757600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135611031565b34801561075957600080fd5b5061078c6004803603606081101561077057600080fd5b506001600160a01b038135169060208101359060400135611092565b604080516001600160a01b039092168252519081900360200190f35b3480156107b457600080fd5b5061043d600480360360408110156107cb57600080fd5b506001600160a01b03813581169160200135166110fd565b3480156107ef57600080fd5b506102b76004803603608081101561080657600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135611138565b34801561083857600080fd5b506102b76004803603608081101561084f57600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135611143565b34801561088157600080fd5b506102b76004803603604081101561089857600080fd5b506001600160a01b038135169060200135611266565b3480156108ba57600080fd5b506102b7600480360360208110156108d157600080fd5b50356001600160a01b0316611316565b3480156108ed57600080fd5b506102b76004803603604081101561090457600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561092e57600080fd5b82018360208201111561094057600080fd5b803590602001918460018302840111600160201b8311171561096157600080fd5b5090925090506113e3565b6001600160a01b038116600090815260076020526040902080546001909101545b915091565b60008061099d612821565b8360005b81811015610a4c5760005b8989838181106109b857fe5b90506020020135811015610a4357610a078c8c8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525089925061149c915050565b9197509550935085610a1e57505050505050610a52565b610a3b888884818110610a2d57fe5b905060200201358286611588565b6001016109ac565b506001016109a1565b50505050505b505050505050565b6001600160a01b0380821660009081526005602090815260408083209386168352908390529020546060919080610aa35750506040805160008152602081019091529050610b1a565b816001016001820381548110610ab557fe5b9060005260206000209060030201600201805480602002602001604051908101604052809291908181526020018280548015610b1057602002820191906000526020600020905b815481526020019060010190808311610afc575b5050505050925050505b92915050565b610b638360053385858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117ab92505050565b505050565b6000610b748233610c1e565b9050610b81338383611882565b610bbc5760405162461bcd60e51b815260040180806020018281038252602e81526020018061293e602e913960400191505060405180910390fd5b604080516340c10f1960e01b81523360048201526024810183905290516001600160a01b038416916340c10f1991604480830192600092919082900301818387803b158015610c0a57600080fd5b505af1158015610a52573d6000803e3d6000fd5b6001600160a01b03808216600090815260026020908152604080832093861683529083905281205490919080610c5957600092505050610b1a565b816001016001820381548110610c6b57fe5b9060005260206000209060020201600101549250505092915050565b6001600160a01b038116600090815260056020908152604091829020600181015483518181528184028101909301909352606092909183918015610cd5578160200160208202803883390190505b50805190915060005b81811015610d4057836001018181548110610cf557fe5b600091825260209091206003909102015483516001600160a01b0390911690849083908110610d2057fe5b6001600160a01b0390921660209283029190910190910152600101610cde565b5090949350505050565b6001600160a01b03808316600090815260056020908152604080832093871683529083905281205490919080610d8557600092505050610dc3565b816001016001820381548110610d9757fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b6001600160a01b031660009081526003602052604090205490565b610dee82611a15565b604080516001600160a01b038316602082015234818301528151808203830181526060909101909152610e2790839060009033906117ab565b5050565b610e273360043385858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117ab92505050565b6001600160a01b038116600090815260046020908152604091829020600181015483518181528184028101909301909352606092909183918015610ebc578160200160208202803883390190505b50805190915060005b81811015610d4057836001018181548110610edc57fe5b600091825260209091206002909102015483516001600160a01b0390911690849083908110610f0757fe5b6001600160a01b0390921660209283029190910190910152600101610ec5565b610b638360033385858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117ab92505050565b610f75838583611a34565b610fcb846002338660601b6001600160601b0319168660601b6001600160601b031916866040516020018084815260200183815260200182815260200193505050506040516020818303038152906040526117ab565b50505050565b6000610fdc33610dca565b3360008181526003602052604080822082905551929350909183156108fc0291849190818181858888f19350505050158015610e27573d6000803e3d6000fd5b60006020819052908152604090205460ff1681565b61103c838583611aab565b610fcb846001338660601b6001600160601b0319168660601b6001600160601b031916866040516020018084815260200183815260200182815260200193505050506040516020818303038152906040526117ab565b604080516020808201859052818301849052606086811b6001600160601b03191690830152825180830360540181526074909201835281519181019190912060009081526006909152908120546001600160a01b0316806110f65784915050610dc3565b9050610dc3565b6001600160a01b03808216600090815260046020908152604080832093861683529083905281205490919080610c5957600092505050610b1a565b61103c838583611b38565b6000611150858484611092565b9050336001600160a01b038216146111a8576040805162461bcd60e51b815260206004820152601660248201527526bab9ba103132903830bcb6b2b73a1037bbb732b91760511b604482015290519081900360640190fd5b604080516020808201869052818301859052606088811b6001600160601b031916908301528251808303605401815260748301808552815191830191909120600090815260069092529083902080546001600160a01b0319166001600160a01b03898116918217909255918790526094830186905280891660b4840152841660d483015260f482015290517fb6cb19e71486466b0282ce82fd31a015b7c00f7d67cddc0da09cccfa58c6438391610114908290030190a15050505050565b611271338383611ba9565b6112c2576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b158015610c0a57600080fd5b600061132282336110fd565b905061132f338383611e11565b61136a5760405162461bcd60e51b815260040180806020018281038252602e81526020018061293e602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b1580156113b957600080fd5b505af11580156113cd573d6000803e3d6000fd5b505050506040513d6020811015610fcb57600080fd5b333214611425576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b60006114508460033386866040518083838082843760405192018290039091209350611e5892505050565b60408051828152905191925033916003916001600160a01b038816917fe923069519faf69b0726ed766a213f61b6f07f2ecf11d55582cc440d8806b0bc9181900360200190a450505050565b6000806114a7612821565b83915060008583815181106114b857fe5b016020015160019093019260f81c90506114d0611eac565b60030160ff168160ff16146114ec575060009250839150611581565b60006114f88785611eb2565b91965094509050846115135750600093508492506115819050565b60ff8116835260006115258886611eb2565b9197509550905085611541575060009450859350611581915050565b6001600160a01b038116602085015261155a8886611f2f565b6040870152909650945085611579575060009450859350611581915050565b506001945050505b9250925092565b805160ff166115f257600061159b612840565b6115a88360400151612186565b91509150816115b8575050610b63565b60006115c982600001518787611092565b90506115da826000015187876121e2565b6115e93382846020015161223e565b50505050610b63565b805160ff16600114156116a0576000611609612857565b611616836040015161229c565b9150915081611626575050610b63565b600061163782602001518787611092565b9050611648826020015187876121e2565b81516001600160a01b031660009081526020819052604090205460ff16156116845761167e338284600001518560400151612319565b50611698565b6115e9338284600001518560400151612365565b505050610b63565b805160ff166002141561170a5760006116b7612857565b6116c4836040015161229c565b91509150816116d4575050610b63565b60006116e582602001518787611092565b90506116f6826020015187876121e2565b6115e9338284600001518560400151612394565b805160ff1660051415610b63576040805162461bcd60e51b815260206004820152600d60248201526c6a757374206661696c206e6f7760981b604482015290519081900360640190fd5b8381101561176c578181015183820152602001611754565b50505050905090810190601f1680156117995780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b60006117c08585858580519060200120611e58565b9050826001600160a01b03168460ff16866001600160a01b03167f35e48d636f39df5c5ca2278452d6d89bf9f07c2ff15f46d08aa402c46638b88284866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611840578181015183820152602001611828565b50505050905090810190601f16801561186d5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a45050505050565b60008161189157506001610dc3565b6001600160a01b038085166000908152600260209081526040808320938716835290839052902054806118c957600092505050610dc3565b60008260010160018303815481106118dd57fe5b9060005260206000209060020201905080600101548511156119055760009350505050610dc3565b60018101805486900390819055611a08576001830180548391859160009190600019810190811061193257fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061197057fe5b906000526020600020906002020183600101600184038154811061199057fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b03938416178155600194850154908501559089168252859052604081205583018054806119de57fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b5060019695505050505050565b6001600160a01b03166000908152600360205260409020805434019055565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd91606480830192600092919082900301818387803b158015611a8857600080fd5b505af1158015611a9c573d6000803e3d6000fd5b50505050610b638284836123b4565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd9160648083019260209291908290030181600087803b158015611b0057600080fd5b505af1158015611b14573d6000803e3d6000fd5b505050506040513d6020811015611b2a57600080fd5b50610b639050828483612538565b60408051632770a7eb60e21b81523360048201526024810183905290516001600160a01b03851691639dc29fac91604480830192600092919082900301818387803b158015611b8657600080fd5b505af1158015611b9a573d6000803e3d6000fd5b50505050610b6382848361260f565b6001600160a01b03808416600090815260056020908152604080832093861683529083905281205490919080611be457600092505050610dc3565b6000826001016001830381548110611bf857fe5b600091825260208083208884526001600390930201918201905260409091205490915080611c2d576000945050505050610dc3565b60028201805482916001850191600091906000198101908110611c4c57fe5b600091825260208083209091015483528201929092526040019020556002820180546000198101908110611c7c57fe5b9060005260206000200154826002016001830381548110611c9957fe5b600091825260208083209091019290925587815260018401909152604081205560028201805480611cc657fe5b6000828152602081208201600019908101919091550190556002820154611e035760018401805484918691600091906000198101908110611d0357fe5b600091825260208083206003909202909101546001600160a01b031683528201929092526040019020556001840180546000198101908110611d4157fe5b9060005260206000209060030201846001016001850381548110611d6157fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b0390921691909117815560028083018054611da49284019190612877565b5050506001600160a01b03871660009081526020859052604081205560018401805480611dcd57fe5b60008281526020812060036000199093019283020180546001600160a01b031916815590611dfe60028301826128c7565b505090555b506001979650505050505050565b600081611e2057506001610dc3565b6001600160a01b038085166000908152600460209081526040808320938716835290839052902054806118c957600092505050610dc3565b6001600160a01b03841660009081526007602052604081206001808201540182611e8687874342868a6126bf565b9050611e96836000015482612728565b835550600190910181905590505b949350505050565b60035b90565b6000806000808551905084811080611ecc57506021858203105b80611ef45750611eda612754565b60ff16868681518110611ee957fe5b016020015160f81c14155b15611f09575060009250839150829050611581565b600160218601611f218888840163ffffffff61275916565b935093509350509250925092565b60008060606000611f408686612775565b9195509350905083611f56575060009250611581565b60208104601f8216600081611f6c576000611f6f565b60015b60ff1683019050606083604051908082528060200260200182016040528015611fa2578160200160208202803883390190505b5090506060836040519080825280601f01601f191660200182016040528015611fd2576020820181803883390190505b5090506000805b8481101561209d576000611fed8e8c612775565b919d509b5090508b61200c575060009a50611581975050505050505050565b8115801561201a5750600087115b1561206d578060005b888110156120665781816020811061203757fe5b1a60f81b86828151811061204757fe5b60200101906001600160f81b031916908160001a905350600101612023565b5050612094565b8060001b858460018b03038151811061208257fe5b60209081029190910101526001909201915b50600101611fd9565b5060006120aa8d8b6127d7565b909a5090506120b7611eac565b60ff168160ff16146120d55750600099506115819650505050505050565b60018a858560405160200180838051906020019060200280838360005b8381101561210a5781810151838201526020016120f2565b5050505090500182805190602001908083835b6020831061213c5780518252601f19909201916020918201910161211d565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529a509a509a5050505050505050509250925092565b6000612190612840565b6034835110156121a3576000915061098d565b600c6121b5848263ffffffff6127fe16565b6001600160a01b031682526014016121d3848263ffffffff61275916565b60208301525060019150915091565b60408051602080820194909452808201929092526001600160601b0319606094851b16938201939093528251605481830301815260749091018352805190820120600090815260069091522080546001600160a01b0319169055565b6001600160a01b03831660009081526003602052604081205482111561226657506000610dc3565b506001600160a01b0392831660009081526003602052604080822080548490039055929093168352912080549091019055600190565b60006122a6612857565b6048835110156122b9576000915061098d565b600c6122cb848263ffffffff6127fe16565b6001600160a01b031682526020016122e9848263ffffffff6127fe16565b6001600160a01b0316602083015260140161230a848263ffffffff61275916565b60408301525060019150915091565b6000806123268487610c1e565b9050828110156123415761233b868583611882565b5061234e565b61234c868585611882565b505b61235985858561260f565b50600195945050505050565b6000612372858484611e11565b61237e57506000611ea4565b612389848484612538565b506001949350505050565b60006123a1858484611ba9565b6123ad57506000611ea4565b6123898484845b6001600160a01b03808416600090815260056020908152604080832093861683529083905290205480612474576040805180820182526001600160a01b0386811682528251600080825260208083019095528484019182526001878101805491820180825590835291869020855160039092020180546001600160a01b03191691909416178355905180519194612453926002850192909101906128e8565b5050506001600160a01b038516600090815260208490526040902081905590505b600082600101600183038154811061248857fe5b9060005260206000209060030201905080600101600085815260200190815260200160002054600014612502576040805162461bcd60e51b815260206004820152601d60248201527f63616e27742061646420616c7265616479206f776e656420746f6b656e000000604482015290519081900360640190fd5b60028101805460018181018355600083815260208082209093018890559254968352909201909152604090209290925550505050565b8061254257610b63565b6001600160a01b038084166000908152600460209081526040808320938616835290839052902054806125db57506040805180820182526001600160a01b0385811680835260006020808501828152600188810180548083018083559186528486209851600290910290980180546001600160a01b03191698909716979097178655905194019390935590815290849052919091208190555b828260010160018303815481106125ee57fe5b60009182526020909120600160029092020101805490910190555050505050565b8061261957610b63565b6001600160a01b038084166000908152600260209081526040808320938616835290839052902054806125db57506040805180820182526001600160a01b038581168083526000602080850182815260018881018054808301808355828752858720995160028302909a0180546001600160a01b0319169a909916999099178855925196909101959095559282528690529390932082905580549192859281106125ee57fe5b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600090565b6000816020018351101561276c57600080fd5b50016020015190565b60008060008061278586866127d7565b9093509050612792611eac565b60020160ff168160ff16146127ab575060009250611581565b6127b58684611eb2565b91955093509150836127cb575060009250611581565b50600192509250925092565b600080826001018484815181106127ea57fe5b016020015190925060f81c90509250929050565b6000816014018351101561281157600080fd5b500160200151600160601b900490565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b604080516060810182526000808252602082018190529181019190915290565b8280548282559060005260206000209081019282156128b75760005260206000209182015b828111156128b757825482559160010191906001019061289c565b506128c3929150612923565b5090565b50805460008255906000526020600020908101906128e59190612923565b50565b8280548282559060005260206000209081019282156128b7579160200282015b828111156128b7578251825591602001919060010190612908565b611eaf91905b808211156128c3576000815560010161292956fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a72315820412e9f0510df14552e2adfefc4e19893dfc4cb1b25951b0c0cb3e4080683363864736f6c634300050f0032"

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

// GetPairedERC20Balance is a free data retrieval call binding the contract method 0x2f5bf211.
//
// Solidity: function getPairedERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxCaller) GetPairedERC20Balance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "getPairedERC20Balance", _tokenContract, _owner)
	return *ret0, err
}

// GetPairedERC20Balance is a free data retrieval call binding the contract method 0x2f5bf211.
//
// Solidity: function getPairedERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxSession) GetPairedERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalInbox.Contract.GetPairedERC20Balance(&_GlobalInbox.CallOpts, _tokenContract, _owner)
}

// GetPairedERC20Balance is a free data retrieval call binding the contract method 0x2f5bf211.
//
// Solidity: function getPairedERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_GlobalInbox *GlobalInboxCallerSession) GetPairedERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalInbox.Contract.GetPairedERC20Balance(&_GlobalInbox.CallOpts, _tokenContract, _owner)
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

// IsBuddyContract is a free data retrieval call binding the contract method 0xb4346fec.
//
// Solidity: function isBuddyContract(address ) view returns(bool)
func (_GlobalInbox *GlobalInboxCaller) IsBuddyContract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "isBuddyContract", arg0)
	return *ret0, err
}

// IsBuddyContract is a free data retrieval call binding the contract method 0xb4346fec.
//
// Solidity: function isBuddyContract(address ) view returns(bool)
func (_GlobalInbox *GlobalInboxSession) IsBuddyContract(arg0 common.Address) (bool, error) {
	return _GlobalInbox.Contract.IsBuddyContract(&_GlobalInbox.CallOpts, arg0)
}

// IsBuddyContract is a free data retrieval call binding the contract method 0xb4346fec.
//
// Solidity: function isBuddyContract(address ) view returns(bool)
func (_GlobalInbox *GlobalInboxCallerSession) IsBuddyContract(arg0 common.Address) (bool, error) {
	return _GlobalInbox.Contract.IsBuddyContract(&_GlobalInbox.CallOpts, arg0)
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

// SendL2BuddyDeploy is a paid mutator transaction binding the contract method 0x0a1a4af6.
//
// Solidity: function sendL2BuddyDeploy(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxTransactor) SendL2BuddyDeploy(opts *bind.TransactOpts, chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "sendL2BuddyDeploy", chain, messageData)
}

// SendL2BuddyDeploy is a paid mutator transaction binding the contract method 0x0a1a4af6.
//
// Solidity: function sendL2BuddyDeploy(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxSession) SendL2BuddyDeploy(chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendL2BuddyDeploy(&_GlobalInbox.TransactOpts, chain, messageData)
}

// SendL2BuddyDeploy is a paid mutator transaction binding the contract method 0x0a1a4af6.
//
// Solidity: function sendL2BuddyDeploy(address chain, bytes messageData) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) SendL2BuddyDeploy(chain common.Address, messageData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendL2BuddyDeploy(&_GlobalInbox.TransactOpts, chain, messageData)
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

// PairedFTContractsABI is the input ABI used to generate the binding from.
const PairedFTContractsABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getPairedERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBuddyContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawPairedERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PairedFTContractsFuncSigs maps the 4-byte function signature to its string representation.
var PairedFTContractsFuncSigs = map[string]string{
	"2f5bf211": "getPairedERC20Balance(address,address)",
	"b4346fec": "isBuddyContract(address)",
	"21e2f2fa": "withdrawPairedERC20(address)",
}

// PairedFTContractsBin is the compiled bytecode used for deploying new contracts.
var PairedFTContractsBin = "0x608060405234801561001057600080fd5b5061041d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806321e2f2fa146100465780632f5bf2111461006e578063b4346fec146100ae575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b03166100e8565b005b61009c6004803603604081101561008457600080fd5b506001600160a01b03813581169160200135166101a6565b60408051918252519081900360200190f35b6100d4600480360360208110156100c457600080fd5b50356001600160a01b0316610210565b604080519115158252519081900360200190f35b60006100f482336101a6565b9050610101338383610225565b61013c5760405162461bcd60e51b815260040180806020018281038252602e8152602001806103bb602e913960400191505060405180910390fd5b604080516340c10f1960e01b81523360048201526024810183905290516001600160a01b038416916340c10f1991604480830192600092919082900301818387803b15801561018a57600080fd5b505af115801561019e573d6000803e3d6000fd5b505050505050565b6001600160a01b038082166000908152600260209081526040808320938616835290839052812054909190806101e15760009250505061020a565b8160010160018203815481106101f357fe5b906000526020600020906002020160010154925050505b92915050565b60006020819052908152604090205460ff1681565b600081610234575060016103b3565b6001600160a01b0380851660009081526002602090815260408083209387168352908390529020548061026c576000925050506103b3565b600082600101600183038154811061028057fe5b9060005260206000209060020201905080600101548511156102a857600093505050506103b3565b600181018054869003908190556103ab57600183018054839185916000919060001981019081106102d557fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061031357fe5b906000526020600020906002020183600101600184038154811061033357fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061038157fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b600193505050505b939250505056fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a7231582002aa87205d03ec009d4b998688db863331f93fe8ec7c55e5b8b441ecfdb453c364736f6c634300050f0032"

// DeployPairedFTContracts deploys a new Ethereum contract, binding an instance of PairedFTContracts to it.
func DeployPairedFTContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PairedFTContracts, error) {
	parsed, err := abi.JSON(strings.NewReader(PairedFTContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairedFTContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PairedFTContracts{PairedFTContractsCaller: PairedFTContractsCaller{contract: contract}, PairedFTContractsTransactor: PairedFTContractsTransactor{contract: contract}, PairedFTContractsFilterer: PairedFTContractsFilterer{contract: contract}}, nil
}

// PairedFTContracts is an auto generated Go binding around an Ethereum contract.
type PairedFTContracts struct {
	PairedFTContractsCaller     // Read-only binding to the contract
	PairedFTContractsTransactor // Write-only binding to the contract
	PairedFTContractsFilterer   // Log filterer for contract events
}

// PairedFTContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairedFTContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedFTContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairedFTContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedFTContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairedFTContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairedFTContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairedFTContractsSession struct {
	Contract     *PairedFTContracts // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairedFTContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairedFTContractsCallerSession struct {
	Contract *PairedFTContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PairedFTContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairedFTContractsTransactorSession struct {
	Contract     *PairedFTContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PairedFTContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairedFTContractsRaw struct {
	Contract *PairedFTContracts // Generic contract binding to access the raw methods on
}

// PairedFTContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairedFTContractsCallerRaw struct {
	Contract *PairedFTContractsCaller // Generic read-only contract binding to access the raw methods on
}

// PairedFTContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairedFTContractsTransactorRaw struct {
	Contract *PairedFTContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairedFTContracts creates a new instance of PairedFTContracts, bound to a specific deployed contract.
func NewPairedFTContracts(address common.Address, backend bind.ContractBackend) (*PairedFTContracts, error) {
	contract, err := bindPairedFTContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairedFTContracts{PairedFTContractsCaller: PairedFTContractsCaller{contract: contract}, PairedFTContractsTransactor: PairedFTContractsTransactor{contract: contract}, PairedFTContractsFilterer: PairedFTContractsFilterer{contract: contract}}, nil
}

// NewPairedFTContractsCaller creates a new read-only instance of PairedFTContracts, bound to a specific deployed contract.
func NewPairedFTContractsCaller(address common.Address, caller bind.ContractCaller) (*PairedFTContractsCaller, error) {
	contract, err := bindPairedFTContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairedFTContractsCaller{contract: contract}, nil
}

// NewPairedFTContractsTransactor creates a new write-only instance of PairedFTContracts, bound to a specific deployed contract.
func NewPairedFTContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*PairedFTContractsTransactor, error) {
	contract, err := bindPairedFTContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairedFTContractsTransactor{contract: contract}, nil
}

// NewPairedFTContractsFilterer creates a new log filterer instance of PairedFTContracts, bound to a specific deployed contract.
func NewPairedFTContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*PairedFTContractsFilterer, error) {
	contract, err := bindPairedFTContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairedFTContractsFilterer{contract: contract}, nil
}

// bindPairedFTContracts binds a generic wrapper to an already deployed contract.
func bindPairedFTContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairedFTContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairedFTContracts *PairedFTContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairedFTContracts.Contract.PairedFTContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairedFTContracts *PairedFTContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairedFTContracts.Contract.PairedFTContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairedFTContracts *PairedFTContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairedFTContracts.Contract.PairedFTContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairedFTContracts *PairedFTContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairedFTContracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairedFTContracts *PairedFTContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairedFTContracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairedFTContracts *PairedFTContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairedFTContracts.Contract.contract.Transact(opts, method, params...)
}

// GetPairedERC20Balance is a free data retrieval call binding the contract method 0x2f5bf211.
//
// Solidity: function getPairedERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_PairedFTContracts *PairedFTContractsCaller) GetPairedERC20Balance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PairedFTContracts.contract.Call(opts, out, "getPairedERC20Balance", _tokenContract, _owner)
	return *ret0, err
}

// GetPairedERC20Balance is a free data retrieval call binding the contract method 0x2f5bf211.
//
// Solidity: function getPairedERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_PairedFTContracts *PairedFTContractsSession) GetPairedERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _PairedFTContracts.Contract.GetPairedERC20Balance(&_PairedFTContracts.CallOpts, _tokenContract, _owner)
}

// GetPairedERC20Balance is a free data retrieval call binding the contract method 0x2f5bf211.
//
// Solidity: function getPairedERC20Balance(address _tokenContract, address _owner) view returns(uint256)
func (_PairedFTContracts *PairedFTContractsCallerSession) GetPairedERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _PairedFTContracts.Contract.GetPairedERC20Balance(&_PairedFTContracts.CallOpts, _tokenContract, _owner)
}

// IsBuddyContract is a free data retrieval call binding the contract method 0xb4346fec.
//
// Solidity: function isBuddyContract(address ) view returns(bool)
func (_PairedFTContracts *PairedFTContractsCaller) IsBuddyContract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PairedFTContracts.contract.Call(opts, out, "isBuddyContract", arg0)
	return *ret0, err
}

// IsBuddyContract is a free data retrieval call binding the contract method 0xb4346fec.
//
// Solidity: function isBuddyContract(address ) view returns(bool)
func (_PairedFTContracts *PairedFTContractsSession) IsBuddyContract(arg0 common.Address) (bool, error) {
	return _PairedFTContracts.Contract.IsBuddyContract(&_PairedFTContracts.CallOpts, arg0)
}

// IsBuddyContract is a free data retrieval call binding the contract method 0xb4346fec.
//
// Solidity: function isBuddyContract(address ) view returns(bool)
func (_PairedFTContracts *PairedFTContractsCallerSession) IsBuddyContract(arg0 common.Address) (bool, error) {
	return _PairedFTContracts.Contract.IsBuddyContract(&_PairedFTContracts.CallOpts, arg0)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_PairedFTContracts *PairedFTContractsTransactor) WithdrawPairedERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _PairedFTContracts.contract.Transact(opts, "withdrawPairedERC20", _tokenContract)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_PairedFTContracts *PairedFTContractsSession) WithdrawPairedERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _PairedFTContracts.Contract.WithdrawPairedERC20(&_PairedFTContracts.TransactOpts, _tokenContract)
}

// WithdrawPairedERC20 is a paid mutator transaction binding the contract method 0x21e2f2fa.
//
// Solidity: function withdrawPairedERC20(address _tokenContract) returns()
func (_PairedFTContracts *PairedFTContractsTransactorSession) WithdrawPairedERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _PairedFTContracts.Contract.WithdrawPairedERC20(&_PairedFTContracts.TransactOpts, _tokenContract)
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
