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
var GlobalEthWalletBin = "0x608060405234801561001057600080fd5b50610110806100206000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80634d2301cc146037578063a0ef91df14606c575b600080fd5b605a60048036036020811015604b57600080fd5b50356001600160a01b03166074565b60408051918252519081900360200190f35b6072608f565b005b6001600160a01b031660009081526020819052604090205490565b60006098336074565b3360008181526020819052604080822082905551929350909183156108fc0291849190818181858888f1935050505015801560d7573d6000803e3d6000fd5b505056fea265627a7a723158207f7c22f12286eac205117065b03e1b4b3eee79a5ac0cefe05583191116e7701164736f6c63430005110032"

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
const GlobalFTWalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"FAILED_TRANSFER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"}],\"name\":\"isPairedContract\",\"outputs\":[{\"internalType\":\"enumGlobalFTWallet.PairingStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC20s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalFTWalletFuncSigs maps the 4-byte function signature to its string representation.
var GlobalFTWalletFuncSigs = map[string]string{
	"56434fe9": "FAILED_TRANSFER()",
	"c3a8962c": "getERC20Balance(address,address)",
	"659e42cd": "isPairedContract(address,address)",
	"6e2b89c5": "ownedERC20s(address)",
	"f4f3b200": "withdrawERC20(address)",
}

// GlobalFTWalletBin is the compiled bytecode used for deploying new contracts.
var GlobalFTWalletBin = "0x608060405234801561001057600080fd5b50610774806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806356434fe91461005c578063659e42cd146100d95780636e2b89c51461012b578063c3a8962c146101a1578063f4f3b200146101e1575b600080fd5b610064610209565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561009e578181015183820152602001610086565b50505050905090810190601f1680156100cb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610107600480360360408110156100ef57600080fd5b506001600160a01b0381358116916020013516610234565b6040518082600281111561011757fe5b60ff16815260200191505060405180910390f35b6101516004803603602081101561014157600080fd5b50356001600160a01b0316610267565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561018d578181015183820152602001610175565b505050509050019250505060405180910390f35b6101cf600480360360408110156101b757600080fd5b506001600160a01b0381358116916020013516610328565b60408051918252519081900360200190f35b610207600480360360208110156101f757600080fd5b50356001600160a01b031661038f565b005b6040518060400160405280600f81526020016e2320a4a622a22faa2920a729a322a960891b81525081565b6001600160a01b0380831660009081526001602081815260408084209486168452939091019052205460ff165b92915050565b6001600160a01b038116600090815260208181526040918290206001810154835181815281840281019093019093526060929091839180156102b3578160200160208202803883390190505b50805190915060005b8181101561031e578360010181815481106102d357fe5b600091825260209091206002909102015483516001600160a01b03909116908490839081106102fe57fe5b6001600160a01b03909216602092830291909101909101526001016102bc565b5090949350505050565b6001600160a01b038082166000908152602081815260408083209386168352908390528120549091908061036157600092505050610261565b81600101600182038154811061037357fe5b9060005260206000209060020201600101549250505092915050565b600061039b8233610328565b90506103a83383836105ac565b6103f0576040805162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e742062616c616e636560601b604482015290519081900360640190fd5b6001600160a01b03821660009081526001602052604090205460ff161561047c57604080516340c10f1960e01b81523360048201526024810183905290516001600160a01b038416916340c10f1991604480830192600092919082900301818387803b15801561045f57600080fd5b505af1158015610473573d6000803e3d6000fd5b505050506105a8565b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b1580156104cb57600080fd5b505af11580156104df573d6000803e3d6000fd5b505050506040513d60208110156104f557600080fd5b505160408051808201909152600f81526e2320a4a622a22faa2920a729a322a960891b6020820152906105a65760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561056b578181015183820152602001610553565b50505050905090810190601f1680156105985780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505b5050565b6000816105bb57506001610738565b6001600160a01b03808516600090815260208181526040808320938716835290839052902054806105f157600092505050610738565b600082600101600183038154811061060557fe5b90600052602060002090600202019050806001015485111561062d5760009350505050610738565b60018101805486900390819055610730576001830180548391859160009190600019810190811061065a57fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061069857fe5b90600052602060002090600202018360010160018403815481106106b857fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061070657fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b600193505050505b939250505056fea265627a7a7231582046222a7dac8d1ac17efcbadae22b42e6fd4f0cb01051b5b1d6265e97dc77688664736f6c63430005110032"

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

// FAILEDTRANSFER is a free data retrieval call binding the contract method 0x56434fe9.
//
// Solidity: function FAILED_TRANSFER() view returns(string)
func (_GlobalFTWallet *GlobalFTWalletCaller) FAILEDTRANSFER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GlobalFTWallet.contract.Call(opts, out, "FAILED_TRANSFER")
	return *ret0, err
}

// FAILEDTRANSFER is a free data retrieval call binding the contract method 0x56434fe9.
//
// Solidity: function FAILED_TRANSFER() view returns(string)
func (_GlobalFTWallet *GlobalFTWalletSession) FAILEDTRANSFER() (string, error) {
	return _GlobalFTWallet.Contract.FAILEDTRANSFER(&_GlobalFTWallet.CallOpts)
}

// FAILEDTRANSFER is a free data retrieval call binding the contract method 0x56434fe9.
//
// Solidity: function FAILED_TRANSFER() view returns(string)
func (_GlobalFTWallet *GlobalFTWalletCallerSession) FAILEDTRANSFER() (string, error) {
	return _GlobalFTWallet.Contract.FAILEDTRANSFER(&_GlobalFTWallet.CallOpts)
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

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(uint8)
func (_GlobalFTWallet *GlobalFTWalletCaller) IsPairedContract(opts *bind.CallOpts, _tokenContract common.Address, _chain common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GlobalFTWallet.contract.Call(opts, out, "isPairedContract", _tokenContract, _chain)
	return *ret0, err
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(uint8)
func (_GlobalFTWallet *GlobalFTWalletSession) IsPairedContract(_tokenContract common.Address, _chain common.Address) (uint8, error) {
	return _GlobalFTWallet.Contract.IsPairedContract(&_GlobalFTWallet.CallOpts, _tokenContract, _chain)
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(uint8)
func (_GlobalFTWallet *GlobalFTWalletCallerSession) IsPairedContract(_tokenContract common.Address, _chain common.Address) (uint8, error) {
	return _GlobalFTWallet.Contract.IsPairedContract(&_GlobalFTWallet.CallOpts, _tokenContract, _chain)
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
const GlobalInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BuddyContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"BuddyContractPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"}],\"name\":\"MessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"PaymentTransfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"FAILED_TRANSFER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"depositERC721Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC721Tokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getInbox\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"getPaymentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasERC721\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"}],\"name\":\"isPairedContract\",\"outputs\":[{\"internalType\":\"enumGlobalFTWallet.PairingStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC20s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC721s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendInitializationMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"initialMaxSendCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalMaxSendCount\",\"type\":\"uint256\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"transferPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalInboxFuncSigs maps the 4-byte function signature to its string representation.
var GlobalInboxFuncSigs = map[string]string{
	"56434fe9": "FAILED_TRANSFER()",
	"41acf614": "deployL2ContractPair(address,uint256,uint256,uint256,bytes)",
	"bca22b76": "depositERC20Message(address,address,address,uint256)",
	"8b7010aa": "depositERC721Message(address,address,address,uint256)",
	"5bd21290": "depositEthMessage(address,address)",
	"c3a8962c": "getERC20Balance(address,address)",
	"0758fb0a": "getERC721Tokens(address,address)",
	"4d2301cc": "getEthBalance(address)",
	"02201681": "getInbox(address)",
	"0547e1b4": "getPaymentOwner(address,uint256)",
	"45a53f09": "hasERC721(address,address,uint256)",
	"659e42cd": "isPairedContract(address,address)",
	"6e2b89c5": "ownedERC20s(address)",
	"33f2ac42": "ownedERC721s(address)",
	"5cc96efa": "sendInitializationMessage(bytes)",
	"74c6eccc": "sendL2Message(address,bytes)",
	"fbef861b": "sendL2MessageFromOrigin(address,bytes)",
	"caba3af0": "sendMessages(bytes,uint256,uint256)",
	"e628c91c": "transferPayment(address,address,uint256)",
	"f4f3b200": "withdrawERC20(address)",
	"f3e414f8": "withdrawERC721(address,uint256)",
	"a0ef91df": "withdrawEth()",
}

// GlobalInboxBin is the compiled bytecode used for deploying new contracts.
var GlobalInboxBin = "0x608060405234801561001057600080fd5b50612cee806100206000396000f3fe6080604052600436106101355760003560e01c80636e2b89c5116100ab578063c3a8962c1161006f578063c3a8962c146106ca578063caba3af014610705578063e628c91c14610786578063f3e414f8146107c9578063f4f3b20014610802578063fbef861b1461083557610135565b80636e2b89c51461056557806374c6eccc146105985780638b7010aa14610623578063a0ef91df1461066c578063bca22b761461068157610135565b806345a53f09116100fd57806345a53f09146103375780634d2301cc1461038e57806356434fe9146103d35780635bd212901461045d5780635cc96efa1461048b578063659e42cd1461050657610135565b8063022016811461013a5780630547e1b4146101865780630758fb0a146101db57806333f2ac421461026657806341acf61414610299575b600080fd5b34801561014657600080fd5b5061016d6004803603602081101561015d57600080fd5b50356001600160a01b03166108c0565b6040805192835260208301919091528051918290030190f35b34801561019257600080fd5b506101bf600480360360408110156101a957600080fd5b506001600160a01b0381351690602001356108e6565b604080516001600160a01b039092168252519081900360200190f35b3480156101e757600080fd5b50610216600480360360408110156101fe57600080fd5b506001600160a01b038135811691602001351661094d565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561025257818101518382015260200161023a565b505050509050019250505060405180910390f35b34801561027257600080fd5b506102166004803603602081101561028957600080fd5b50356001600160a01b0316610a12565b3480156102a557600080fd5b50610335600480360360a08110156102bc57600080fd5b6001600160a01b038235169160208101359160408201359160608101359181019060a081016080820135600160201b8111156102f757600080fd5b82018360208201111561030957600080fd5b803590602001918460018302840111600160201b8311171561032a57600080fd5b509092509050610ad5565b005b34801561034357600080fd5b5061037a6004803603606081101561035a57600080fd5b506001600160a01b03813581169160208101359091169060400135610bcd565b604080519115158252519081900360200190f35b34801561039a57600080fd5b506103c1600480360360208110156103b157600080fd5b50356001600160a01b0316610c4d565b60408051918252519081900360200190f35b3480156103df57600080fd5b506103e8610c68565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561042257818101518382015260200161040a565b50505050905090810190601f16801561044f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103356004803603604081101561047357600080fd5b506001600160a01b0381358116916020013516610c93565b34801561049757600080fd5b50610335600480360360208110156104ae57600080fd5b810190602081018135600160201b8111156104c857600080fd5b8201836020820111156104da57600080fd5b803590602001918460018302840111600160201b831117156104fb57600080fd5b509092509050610cd9565b34801561051257600080fd5b506105416004803603604081101561052957600080fd5b506001600160a01b0381358116916020013516610d1c565b6040518082600281111561055157fe5b60ff16815260200191505060405180910390f35b34801561057157600080fd5b506102166004803603602081101561058857600080fd5b50356001600160a01b0316610d4e565b3480156105a457600080fd5b50610335600480360360408110156105bb57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105e557600080fd5b8201836020820111156105f757600080fd5b803590602001918460018302840111600160201b8311171561061857600080fd5b509092509050610e05565b34801561062f57600080fd5b506103356004803603608081101561064657600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610e4d565b34801561067857600080fd5b50610335610ea2565b34801561068d57600080fd5b50610335600480360360808110156106a457600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610eed565b3480156106d657600080fd5b506103c1600480360360408110156106ed57600080fd5b506001600160a01b0381358116916020013516610f3c565b34801561071157600080fd5b506103356004803603606081101561072857600080fd5b810190602081018135600160201b81111561074257600080fd5b82018360208201111561075457600080fd5b803590602001918460018302840111600160201b8311171561077557600080fd5b919350915080359060200135610fa5565b34801561079257600080fd5b50610335600480360360608110156107a957600080fd5b506001600160a01b0381358116916020810135909116906040013561102c565b3480156107d557600080fd5b50610335600480360360408110156107ec57600080fd5b506001600160a01b03813516906020013561113f565b34801561080e57600080fd5b506103356004803603602081101561082557600080fd5b50356001600160a01b031661120b565b34801561084157600080fd5b506103356004803603604081101561085857600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561088257600080fd5b82018360208201111561089457600080fd5b803590602001918460018302840111600160201b831117156108b557600080fd5b509092509050611422565b6001600160a01b038116600090815260056020526040902080546001909101545b915091565b6040805160208082018490526001600160601b0319606086901b1682840152825160348184030181526054909201835281519181019190912060009081526004909152908120546001600160a01b0316806109445783915050610947565b90505b92915050565b6001600160a01b03808216600090815260036020908152604080832093861683529083905290205460609190806109965750506040805160008152602081019091529050610947565b8160010160018203815481106109a857fe5b9060005260206000209060030201600201805480602002602001604051908101604052809291908181526020018280548015610a0357602002820191906000526020600020905b8154815260200190600101908083116109ef575b50505050509250505092915050565b6001600160a01b038116600090815260036020908152604091829020600181015483518181528184028101909301909352606092909183918015610a60578160200160208202803883390190505b50805190915060005b81811015610acb57836001018181548110610a8057fe5b600091825260209091206003909102015483516001600160a01b0390911690849083908110610aab57fe5b6001600160a01b0390921660209283029190910190910152600101610a69565b5090949350505050565b610ade336114db565b610b2f576040805162461bcd60e51b815260206004820152601a60248201527f6d7573742062652063616c6c656420627920636f6e7472616374000000000000604482015290519081900360640190fd5b610b393387611517565b610b868660053388888888886040516020018086815260200185815260200184815260200183838082843780830192505050955050505050506040516020818303038152906040526116c3565b604080516001600160a01b0388168152905133917feaa7eb17fe081a8c502cff47a2a944377a71c63065a02cd44b16a06d1a0d4dc7919081900360200190a2505050505050565b6001600160a01b03808316600090815260036020908152604080832093871683529083905281205490919080610c0857600092505050610c46565b816001016001820381548110610c1a57fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b6001600160a01b031660009081526020819052604090205490565b6040518060400160405280600f81526020016e2320a4a622a22faa2920a729a322a960891b81525081565b610c9c8261179a565b604080516001600160a01b038316602082015234818301528151808203830181526060909101909152610cd590839060009033906116c3565b5050565b610cd53360043385858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116c392505050565b6001600160a01b0391821660009081526002602090815260408083209390941682526001909201909152205460ff1690565b6001600160a01b03811660009081526001602081815260409283902091820154835181815281830281019092019093526060928391908015610d9a578160200160208202803883390190505b50805190915060005b81811015610acb57836001018181548110610dba57fe5b600091825260209091206002909102015483516001600160a01b0390911690849083908110610de557fe5b6001600160a01b0390921660209283029190910190910152600101610da3565b610e488360033385858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116c392505050565b505050565b610e588385836117b9565b604080516001600160a01b0380861660208301528416818301526060818101849052825180830390910181526080909101909152610e9c90859060029033906116c3565b50505050565b6000610ead33610c4d565b3360008181526020819052604080822082905551929350909183156108fc0291849190818181858888f19350505050158015610cd5573d6000803e3d6000fd5b610ef8838583611830565b604080516001600160a01b0380861660208301528416818301526060818101849052825180830390910181526080909101909152610e9c90859060019033906116c3565b6001600160a01b03808216600090815260016020908152604080832093861683529083905281205490919080610f7757600092505050610947565b816001016001820381548110610f8957fe5b9060005260206000209060020201600101549250505092915050565b600080610fb0612b9d565b845b8481101561102257610ffb88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250611a06915050565b91955093509150836110105750505050610e9c565b61101a8183611af2565b600101610fb2565b5050505050505050565b600061103884836108e6565b9050336001600160a01b03821614611090576040805162461bcd60e51b815260206004820152601660248201527526bab9ba103132903830bcb6b2b73a1037bbb732b91760511b604482015290519081900360640190fd5b6040805160208082018590526001600160601b0319606088901b16828401528251603481840301815260548301808552815191830191909120600090815260049092529083902080546001600160a01b038089166001600160a01b031990921682179092559186905280881660748401528416609483015260b482015290517f56e2b14cbe3f09b8a7337ab17e6ce5da744339e435db498d5ba53655695c03e59181900360d40190a150505050565b61114a338383611d08565b61119b576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b1580156111ef57600080fd5b505af1158015611203573d6000803e3d6000fd5b505050505050565b60006112178233610f3c565b9050611224338383611f70565b61126c576040805162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e742062616c616e636560601b604482015290519081900360640190fd5b6001600160a01b03821660009081526002602052604090205460ff16156112f857604080516340c10f1960e01b81523360048201526024810183905290516001600160a01b038416916340c10f1991604480830192600092919082900301818387803b1580156112db57600080fd5b505af11580156112ef573d6000803e3d6000fd5b50505050610cd5565b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b15801561134757600080fd5b505af115801561135b573d6000803e3d6000fd5b505050506040513d602081101561137157600080fd5b505160408051808201909152600f81526e2320a4a622a22faa2920a729a322a960891b602082015290610e485760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156113e75781810151838201526020016113cf565b50505050905090810190601f1680156114145780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b333214611464576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600061148f846003338686604051808383808284376040519201829003909120935061210392505050565b60408051828152905191925033916003916001600160a01b038816917fe923069519faf69b0726ed766a213f61b6f07f2ecf11d55582cc440d8806b0bc9181900360200190a450505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061150f57508115155b949350505050565b6001600160a01b0382166000908152600260205260408120906001600160a01b038316600090815260018301602052604090205460ff16600281111561155957fe5b1461159e576040805162461bcd60e51b815260206004820152601060248201526f1b5d5cdd081899481d5b9c185a5c995960821b604482015290519081900360640190fd5b805460ff1661169857805460ff19166001178155604080516370a0823160e01b81523060048201819052915185926001600160a01b03841692639dc29fac9284916370a08231916024808301926020929190829003018186803b15801561160457600080fd5b505afa158015611618573d6000803e3d6000fd5b505050506040513d602081101561162e57600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152602483019190915251604480830192600092919082900301818387803b15801561167e57600080fd5b505af1158015611692573d6000803e3d6000fd5b50505050505b6001600160a01b03919091166000908152600191820160205260409020805460ff1916909117905550565b60006116d88585858580519060200120612103565b9050826001600160a01b03168460ff16866001600160a01b03167f35e48d636f39df5c5ca2278452d6d89bf9f07c2ff15f46d08aa402c46638b88284866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611758578181015183820152602001611740565b50505050905090810190601f1680156117855780820380516001836020036101000a031916815260200191505b50935050505060405180910390a45050505050565b6001600160a01b03166000908152602081905260409020805434019055565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd91606480830192600092919082900301818387803b15801561180d57600080fd5b505af1158015611821573d6000803e3d6000fd5b50505050610e48828483612155565b6001600160a01b03831660009081526002602052604081208054909160ff90911690818015611887575060026001600160a01b038616600090815260018501602052604090205460ff16600281111561188557fe5b145b905080611899576118998587866122d9565b811561190a5760408051632770a7eb60e21b81523360048201526024810186905290516001600160a01b03881691639dc29fac91604480830192600092919082900301818387803b1580156118ed57600080fd5b505af1158015611901573d6000803e3d6000fd5b50505050611203565b604080516323b872dd60e01b81523360048201523060248201526044810186905290516001600160a01b038816916323b872dd9160648083019260209291908290030181600087803b15801561195f57600080fd5b505af1158015611973573d6000803e3d6000fd5b505050506040513d602081101561198957600080fd5b505160408051808201909152600f81526e2320a4a622a22faa2920a729a322a960891b6020820152906119fd5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156113e75781810151838201526020016113cf565b50505050505050565b600080611a11612b9d565b8391506000858381518110611a2257fe5b016020015160019093019260f81c9050611a3a6123b0565b60030160ff168160ff1614611a56575060009250839150611aeb565b6000611a6287856123b6565b9196509450905084611a7d575060009350849250611aeb9050565b60ff811683526000611a8f88866123b6565b9197509550905085611aab575060009450859350611aeb915050565b6001600160a01b0381166020850152611ac48886612433565b6040870152909650945085611ae3575060009450859350611aeb915050565b506001945050505b9250925092565b805160ff16611b51576000611b05612bbc565b611b12836040015161268a565b9150915081611b22575050610cd5565b6000611b328260000151866108e6565b9050611b428260000151866126e6565b6112ef3382846020015161273c565b805160ff1660011415611bc2576000611b68612bd3565b611b75836040015161279a565b9150915081611b85575050610cd5565b6000611b958260200151866108e6565b9050611bab338284600001518560400151612817565b50611bba8260200151866126e6565b505050610cd5565b805160ff1660021415611c1c576000611bd9612bd3565b611be6836040015161279a565b9150915081611bf6575050610cd5565b6000611c068260200151866108e6565b9050611bab3382846000015185604001516128eb565b805160ff1660051415610cd557611c598160200151338360400151600081518110611c4357fe5b01602001516001600160f81b031916151561291a565b80602001516001600160a01b03167fa98915d9854858ea787b0abcd4e8e3a96802bc19a25474a8b7017a303628e44482604001516040518080602001828103825283818151815260200191508051906020019080838360005b83811015611cca578181015183820152602001611cb2565b50505050905090810190601f168015611cf75780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b6001600160a01b03808416600090815260036020908152604080832093861683529083905281205490919080611d4357600092505050610c46565b6000826001016001830381548110611d5757fe5b600091825260208083208884526001600390930201918201905260409091205490915080611d8c576000945050505050610c46565b60028201805482916001850191600091906000198101908110611dab57fe5b600091825260208083209091015483528201929092526040019020556002820180546000198101908110611ddb57fe5b9060005260206000200154826002016001830381548110611df857fe5b600091825260208083209091019290925587815260018401909152604081205560028201805480611e2557fe5b6000828152602081208201600019908101919091550190556002820154611f625760018401805484918691600091906000198101908110611e6257fe5b600091825260208083206003909202909101546001600160a01b031683528201929092526040019020556001840180546000198101908110611ea057fe5b9060005260206000209060030201846001016001850381548110611ec057fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b0390921691909117815560028083018054611f039284019190612bf3565b5050506001600160a01b03871660009081526020859052604081205560018401805480611f2c57fe5b60008281526020812060036000199093019283020180546001600160a01b031916815590611f5d6002830182612c43565b505090555b506001979650505050505050565b600081611f7f57506001610c46565b6001600160a01b03808516600090815260016020908152604080832093871683529083905290205480611fb757600092505050610c46565b6000826001016001830381548110611fcb57fe5b906000526020600020906002020190508060010154851115611ff35760009350505050610c46565b600181018054869003908190556120f6576001830180548391859160009190600019810190811061202057fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061205e57fe5b906000526020600020906002020183600101600184038154811061207e57fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b03938416178155600194850154908501559089168252859052604081205583018054806120cc57fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b5060019695505050505050565b6001600160a01b0384166000908152600560205260408120600180820154018261213187874342868a6129c1565b9050612141836000015482612a2a565b835550600190910181905595945050505050565b6001600160a01b03808416600090815260036020908152604080832093861683529083905290205480612215576040805180820182526001600160a01b0386811682528251600080825260208083019095528484019182526001878101805491820180825590835291869020855160039092020180546001600160a01b031916919094161783559051805191946121f492600285019290910190612c64565b5050506001600160a01b038516600090815260208490526040902081905590505b600082600101600183038154811061222957fe5b90600052602060002090600302019050806001016000858152602001908152602001600020546000146122a3576040805162461bcd60e51b815260206004820152601d60248201527f63616e27742061646420616c7265616479206f776e656420746f6b656e000000604482015290519081900360640190fd5b60028101805460018181018355600083815260208082209093018890559254968352909201909152604090209290925550505050565b806122e357610e48565b6001600160a01b0380841660009081526001602090815260408083209386168352908390529020548061237c57506040805180820182526001600160a01b0385811680835260006020808501828152600188810180548083018083559186528486209851600290910290980180546001600160a01b03191698909716979097178655905194019390935590815290849052919091208190555b8282600101600183038154811061238f57fe5b60009182526020909120600160029092020101805490910190555050505050565b60035b90565b60008060008085519050848110806123d057506021858203105b806123f857506123de612a56565b60ff168686815181106123ed57fe5b016020015160f81c14155b1561240d575060009250839150829050611aeb565b6001602186016124258888840163ffffffff612a5b16565b935093509350509250925092565b600080606060006124448686612ab4565b919550935090508361245a575060009250611aeb565b60208104601f8216600081612470576000612473565b60015b60ff16830190506060836040519080825280602002602001820160405280156124a6578160200160208202803883390190505b5090506060836040519080825280601f01601f1916602001820160405280156124d6576020820181803883390190505b5090506000805b848110156125a15760006124f18e8c612ab4565b919d509b5090508b612510575060009a50611aeb975050505050505050565b8115801561251e5750600087115b15612571578060005b8881101561256a5781816020811061253b57fe5b1a60f81b86828151811061254b57fe5b60200101906001600160f81b031916908160001a905350600101612527565b5050612598565b8060001b858460018b03038151811061258657fe5b60209081029190910101526001909201915b506001016124dd565b5060006125ae8d8b612b16565b909a5090506125bb6123b0565b60ff168160ff16146125d9575060009950611aeb9650505050505050565b60018a858560405160200180838051906020019060200280838360005b8381101561260e5781810151838201526020016125f6565b5050505090500182805190602001908083835b602083106126405780518252601f199092019160209182019101612621565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529a509a509a5050505050505050509250925092565b6000612694612bbc565b6034835110156126a757600091506108e1565b600c6126b9848263ffffffff612b3d16565b6001600160a01b031682526014016126d7848263ffffffff612a5b16565b60208301525060019150915091565b6040805160208082019390935260609390931b6001600160601b031916838201528051808403603401815260549093018152825192820192909220600090815260049091522080546001600160a01b0319169055565b6001600160a01b03831660009081526020819052604081205482111561276457506000610c46565b506001600160a01b0392831660009081526020819052604080822080548490039055929093168352912080549091019055600190565b60006127a4612bd3565b6048835110156127b757600091506108e1565b600c6127c9848263ffffffff612b3d16565b6001600160a01b031682526020016127e7848263ffffffff612b3d16565b6001600160a01b03166020830152601401612808848263ffffffff612a5b16565b60408301525060019150915091565b6001600160a01b0382166000908152600260205260408120805460ff168281801561286a575060026001600160a01b038916600090815260018501602052604090205460ff16600281111561286857fe5b145b9050801580156128825750612880888787611f70565b155b15612893576000935050505061150f565b60008280156128ca575060026001600160a01b038916600090815260018601602052604090205460ff1660028111156128c857fe5b145b9050806128dc576128dc8888886122d9565b50600198975050505050505050565b60006128f8858484611d08565b6129045750600061150f565b61290f848484612155565b506001949350505050565b6001600160a01b038316600090815260026020526040902060016001600160a01b038416600090815260018301602052604090205460ff16600281111561295d57fe5b146129685750610e48565b8115612998576001600160a01b03831660009081526001820160205260409020805460ff19166002179055610e9c565b6001600160a01b0392909216600090815260019092016020525060409020805460ff1916905550565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6001600160601b03191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600090565b60008160200183511015612aab576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b600080600080612ac48686612b16565b9093509050612ad16123b0565b60020160ff168160ff1614612aea575060009250611aeb565b612af486846123b6565b9195509350915083612b0a575060009250611aeb565b50600192509250925092565b60008082600101848481518110612b2957fe5b016020015190925060f81c90509250929050565b60008160140183511015612b8d576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b500160200151600160601b900490565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b604080516060810182526000808252602082018190529181019190915290565b828054828255906000526020600020908101928215612c335760005260206000209182015b82811115612c33578254825591600101919060010190612c18565b50612c3f929150612c9f565b5090565b5080546000825590600052602060002090810190612c619190612c9f565b50565b828054828255906000526020600020908101928215612c33579160200282015b82811115612c33578251825591602001919060010190612c84565b6123b391905b80821115612c3f5760008155600101612ca556fea265627a7a723158200b990cc22ea9a2aa7c89474bb4deeb2f3684beb7be5e6cb500292664e3748ad264736f6c63430005110032"

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

// FAILEDTRANSFER is a free data retrieval call binding the contract method 0x56434fe9.
//
// Solidity: function FAILED_TRANSFER() view returns(string)
func (_GlobalInbox *GlobalInboxCaller) FAILEDTRANSFER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "FAILED_TRANSFER")
	return *ret0, err
}

// FAILEDTRANSFER is a free data retrieval call binding the contract method 0x56434fe9.
//
// Solidity: function FAILED_TRANSFER() view returns(string)
func (_GlobalInbox *GlobalInboxSession) FAILEDTRANSFER() (string, error) {
	return _GlobalInbox.Contract.FAILEDTRANSFER(&_GlobalInbox.CallOpts)
}

// FAILEDTRANSFER is a free data retrieval call binding the contract method 0x56434fe9.
//
// Solidity: function FAILED_TRANSFER() view returns(string)
func (_GlobalInbox *GlobalInboxCallerSession) FAILEDTRANSFER() (string, error) {
	return _GlobalInbox.Contract.FAILEDTRANSFER(&_GlobalInbox.CallOpts)
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

// GetPaymentOwner is a free data retrieval call binding the contract method 0x0547e1b4.
//
// Solidity: function getPaymentOwner(address originalOwner, uint256 messageIndex) view returns(address)
func (_GlobalInbox *GlobalInboxCaller) GetPaymentOwner(opts *bind.CallOpts, originalOwner common.Address, messageIndex *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "getPaymentOwner", originalOwner, messageIndex)
	return *ret0, err
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0x0547e1b4.
//
// Solidity: function getPaymentOwner(address originalOwner, uint256 messageIndex) view returns(address)
func (_GlobalInbox *GlobalInboxSession) GetPaymentOwner(originalOwner common.Address, messageIndex *big.Int) (common.Address, error) {
	return _GlobalInbox.Contract.GetPaymentOwner(&_GlobalInbox.CallOpts, originalOwner, messageIndex)
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0x0547e1b4.
//
// Solidity: function getPaymentOwner(address originalOwner, uint256 messageIndex) view returns(address)
func (_GlobalInbox *GlobalInboxCallerSession) GetPaymentOwner(originalOwner common.Address, messageIndex *big.Int) (common.Address, error) {
	return _GlobalInbox.Contract.GetPaymentOwner(&_GlobalInbox.CallOpts, originalOwner, messageIndex)
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
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(uint8)
func (_GlobalInbox *GlobalInboxCaller) IsPairedContract(opts *bind.CallOpts, _tokenContract common.Address, _chain common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GlobalInbox.contract.Call(opts, out, "isPairedContract", _tokenContract, _chain)
	return *ret0, err
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(uint8)
func (_GlobalInbox *GlobalInboxSession) IsPairedContract(_tokenContract common.Address, _chain common.Address) (uint8, error) {
	return _GlobalInbox.Contract.IsPairedContract(&_GlobalInbox.CallOpts, _tokenContract, _chain)
}

// IsPairedContract is a free data retrieval call binding the contract method 0x659e42cd.
//
// Solidity: function isPairedContract(address _tokenContract, address _chain) view returns(uint8)
func (_GlobalInbox *GlobalInboxCallerSession) IsPairedContract(_tokenContract common.Address, _chain common.Address) (uint8, error) {
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

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x41acf614.
//
// Solidity: function deployL2ContractPair(address chain, uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_GlobalInbox *GlobalInboxTransactor) DeployL2ContractPair(opts *bind.TransactOpts, chain common.Address, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "deployL2ContractPair", chain, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x41acf614.
//
// Solidity: function deployL2ContractPair(address chain, uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_GlobalInbox *GlobalInboxSession) DeployL2ContractPair(chain common.Address, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DeployL2ContractPair(&_GlobalInbox.TransactOpts, chain, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x41acf614.
//
// Solidity: function deployL2ContractPair(address chain, uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) DeployL2ContractPair(chain common.Address, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _GlobalInbox.Contract.DeployL2ContractPair(&_GlobalInbox.TransactOpts, chain, maxGas, gasPriceBid, payment, contractData)
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

// SendMessages is a paid mutator transaction binding the contract method 0xcaba3af0.
//
// Solidity: function sendMessages(bytes messages, uint256 initialMaxSendCount, uint256 finalMaxSendCount) returns()
func (_GlobalInbox *GlobalInboxTransactor) SendMessages(opts *bind.TransactOpts, messages []byte, initialMaxSendCount *big.Int, finalMaxSendCount *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "sendMessages", messages, initialMaxSendCount, finalMaxSendCount)
}

// SendMessages is a paid mutator transaction binding the contract method 0xcaba3af0.
//
// Solidity: function sendMessages(bytes messages, uint256 initialMaxSendCount, uint256 finalMaxSendCount) returns()
func (_GlobalInbox *GlobalInboxSession) SendMessages(messages []byte, initialMaxSendCount *big.Int, finalMaxSendCount *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendMessages(&_GlobalInbox.TransactOpts, messages, initialMaxSendCount, finalMaxSendCount)
}

// SendMessages is a paid mutator transaction binding the contract method 0xcaba3af0.
//
// Solidity: function sendMessages(bytes messages, uint256 initialMaxSendCount, uint256 finalMaxSendCount) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) SendMessages(messages []byte, initialMaxSendCount *big.Int, finalMaxSendCount *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.SendMessages(&_GlobalInbox.TransactOpts, messages, initialMaxSendCount, finalMaxSendCount)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xe628c91c.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, uint256 messageIndex) returns()
func (_GlobalInbox *GlobalInboxTransactor) TransferPayment(opts *bind.TransactOpts, originalOwner common.Address, newOwner common.Address, messageIndex *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.contract.Transact(opts, "transferPayment", originalOwner, newOwner, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xe628c91c.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, uint256 messageIndex) returns()
func (_GlobalInbox *GlobalInboxSession) TransferPayment(originalOwner common.Address, newOwner common.Address, messageIndex *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.TransferPayment(&_GlobalInbox.TransactOpts, originalOwner, newOwner, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xe628c91c.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, uint256 messageIndex) returns()
func (_GlobalInbox *GlobalInboxTransactorSession) TransferPayment(originalOwner common.Address, newOwner common.Address, messageIndex *big.Int) (*types.Transaction, error) {
	return _GlobalInbox.Contract.TransferPayment(&_GlobalInbox.TransactOpts, originalOwner, newOwner, messageIndex)
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

// GlobalInboxBuddyContractPairIterator is returned from FilterBuddyContractPair and is used to iterate over the raw logs and unpacked data for BuddyContractPair events raised by the GlobalInbox contract.
type GlobalInboxBuddyContractPairIterator struct {
	Event *GlobalInboxBuddyContractPair // Event containing the contract specifics and raw log

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
func (it *GlobalInboxBuddyContractPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalInboxBuddyContractPair)
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
		it.Event = new(GlobalInboxBuddyContractPair)
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
func (it *GlobalInboxBuddyContractPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalInboxBuddyContractPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalInboxBuddyContractPair represents a BuddyContractPair event raised by the GlobalInbox contract.
type GlobalInboxBuddyContractPair struct {
	Sender common.Address
	Data   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuddyContractPair is a free log retrieval operation binding the contract event 0xeaa7eb17fe081a8c502cff47a2a944377a71c63065a02cd44b16a06d1a0d4dc7.
//
// Solidity: event BuddyContractPair(address indexed sender, address data)
func (_GlobalInbox *GlobalInboxFilterer) FilterBuddyContractPair(opts *bind.FilterOpts, sender []common.Address) (*GlobalInboxBuddyContractPairIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GlobalInbox.contract.FilterLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return &GlobalInboxBuddyContractPairIterator{contract: _GlobalInbox.contract, event: "BuddyContractPair", logs: logs, sub: sub}, nil
}

// WatchBuddyContractPair is a free log subscription operation binding the contract event 0xeaa7eb17fe081a8c502cff47a2a944377a71c63065a02cd44b16a06d1a0d4dc7.
//
// Solidity: event BuddyContractPair(address indexed sender, address data)
func (_GlobalInbox *GlobalInboxFilterer) WatchBuddyContractPair(opts *bind.WatchOpts, sink chan<- *GlobalInboxBuddyContractPair, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GlobalInbox.contract.WatchLogs(opts, "BuddyContractPair", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalInboxBuddyContractPair)
				if err := _GlobalInbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
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

// ParseBuddyContractPair is a log parse operation binding the contract event 0xeaa7eb17fe081a8c502cff47a2a944377a71c63065a02cd44b16a06d1a0d4dc7.
//
// Solidity: event BuddyContractPair(address indexed sender, address data)
func (_GlobalInbox *GlobalInboxFilterer) ParseBuddyContractPair(log types.Log) (*GlobalInboxBuddyContractPair, error) {
	event := new(GlobalInboxBuddyContractPair)
	if err := _GlobalInbox.contract.UnpackLog(event, "BuddyContractPair", log); err != nil {
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
	MessageIndex  *big.Int
	OriginalOwner common.Address
	PrevOwner     common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentTransfer is a free log retrieval operation binding the contract event 0x56e2b14cbe3f09b8a7337ab17e6ce5da744339e435db498d5ba53655695c03e5.
//
// Solidity: event PaymentTransfer(uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_GlobalInbox *GlobalInboxFilterer) FilterPaymentTransfer(opts *bind.FilterOpts) (*GlobalInboxPaymentTransferIterator, error) {

	logs, sub, err := _GlobalInbox.contract.FilterLogs(opts, "PaymentTransfer")
	if err != nil {
		return nil, err
	}
	return &GlobalInboxPaymentTransferIterator{contract: _GlobalInbox.contract, event: "PaymentTransfer", logs: logs, sub: sub}, nil
}

// WatchPaymentTransfer is a free log subscription operation binding the contract event 0x56e2b14cbe3f09b8a7337ab17e6ce5da744339e435db498d5ba53655695c03e5.
//
// Solidity: event PaymentTransfer(uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
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

// ParsePaymentTransfer is a log parse operation binding the contract event 0x56e2b14cbe3f09b8a7337ab17e6ce5da744339e435db498d5ba53655695c03e5.
//
// Solidity: event PaymentTransfer(uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
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
var GlobalNFTWalletBin = "0x608060405234801561001057600080fd5b50610765806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630758fb0a1461005157806333f2ac42146100cf57806345a53f09146100f5578063f3e414f81461013f575b600080fd5b61007f6004803603604081101561006757600080fd5b506001600160a01b038135811691602001351661016d565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100bb5781810151838201526020016100a3565b505050509050019250505060405180910390f35b61007f600480360360208110156100e557600080fd5b50356001600160a01b0316610231565b61012b6004803603606081101561010b57600080fd5b506001600160a01b038135811691602081013590911690604001356102f2565b604080519115158252519081900360200190f35b61016b6004803603604081101561015557600080fd5b506001600160a01b038135169060200135610370565b005b6001600160a01b0380821660009081526020818152604080832093861683529083905290205460609190806101b4575050604080516000815260208101909152905061022b565b8160010160018203815481106101c657fe5b906000526020600020906003020160020180548060200260200160405190810160405280929190818152602001828054801561022157602002820191906000526020600020905b81548152602001906001019080831161020d575b5050505050925050505b92915050565b6001600160a01b0381166000908152602081815260409182902060018101548351818152818402810190930190935260609290918391801561027d578160200160208202803883390190505b50805190915060005b818110156102e85783600101818154811061029d57fe5b600091825260209091206003909102015483516001600160a01b03909116908490839081106102c857fe5b6001600160a01b0390921660209283029190910190910152600101610286565b5090949350505050565b6001600160a01b038083166000908152602081815260408083209387168352908390528120549091908061032b57600092505050610369565b81600101600182038154811061033d57fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b61037b33838361043c565b6103cc576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b15801561042057600080fd5b505af1158015610434573d6000803e3d6000fd5b505050505050565b6001600160a01b038084166000908152602081815260408083209386168352908390528120549091908061047557600092505050610369565b600082600101600183038154811061048957fe5b6000918252602080832088845260016003909302019182019052604090912054909150806104be576000945050505050610369565b600282018054829160018501916000919060001981019081106104dd57fe5b60009182526020808320909101548352820192909252604001902055600282018054600019810190811061050d57fe5b906000526020600020015482600201600183038154811061052a57fe5b60009182526020808320909101929092558781526001840190915260408120556002820180548061055757fe5b6000828152602081208201600019908101919091550190556002820154610694576001840180548491869160009190600019810190811061059457fe5b600091825260208083206003909202909101546001600160a01b0316835282019290925260400190205560018401805460001981019081106105d257fe5b90600052602060002090600302018460010160018503815481106105f257fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b039092169190911781556002808301805461063592840191906106a2565b5050506001600160a01b0387166000908152602085905260408120556001840180548061065e57fe5b60008281526020812060036000199093019283020180546001600160a01b03191681559061068f60028301826106f2565b505090555b506001979650505050505050565b8280548282559060005260206000209081019282156106e25760005260206000209182015b828111156106e25782548255916001019190600101906106c7565b506106ee929150610713565b5090565b50805460008255906000526020600020908101906107109190610713565b50565b61072d91905b808211156106ee5760008155600101610719565b9056fea265627a7a7231582035d61780e777c116b6157cdf3594088be1dfeb070fd14862cf6975a80deb4e4a64736f6c63430005110032"

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

// IPairedErc20ABI is the input ABI used to generate the binding from.
const IPairedErc20ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPairedErc20FuncSigs maps the 4-byte function signature to its string representation.
var IPairedErc20FuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"9dc29fac": "burn(address,uint256)",
	"40c10f19": "mint(address,uint256)",
}

// IPairedErc20 is an auto generated Go binding around an Ethereum contract.
type IPairedErc20 struct {
	IPairedErc20Caller     // Read-only binding to the contract
	IPairedErc20Transactor // Write-only binding to the contract
	IPairedErc20Filterer   // Log filterer for contract events
}

// IPairedErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IPairedErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPairedErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IPairedErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPairedErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPairedErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPairedErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPairedErc20Session struct {
	Contract     *IPairedErc20     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPairedErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPairedErc20CallerSession struct {
	Contract *IPairedErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPairedErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPairedErc20TransactorSession struct {
	Contract     *IPairedErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPairedErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IPairedErc20Raw struct {
	Contract *IPairedErc20 // Generic contract binding to access the raw methods on
}

// IPairedErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPairedErc20CallerRaw struct {
	Contract *IPairedErc20Caller // Generic read-only contract binding to access the raw methods on
}

// IPairedErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPairedErc20TransactorRaw struct {
	Contract *IPairedErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIPairedErc20 creates a new instance of IPairedErc20, bound to a specific deployed contract.
func NewIPairedErc20(address common.Address, backend bind.ContractBackend) (*IPairedErc20, error) {
	contract, err := bindIPairedErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPairedErc20{IPairedErc20Caller: IPairedErc20Caller{contract: contract}, IPairedErc20Transactor: IPairedErc20Transactor{contract: contract}, IPairedErc20Filterer: IPairedErc20Filterer{contract: contract}}, nil
}

// NewIPairedErc20Caller creates a new read-only instance of IPairedErc20, bound to a specific deployed contract.
func NewIPairedErc20Caller(address common.Address, caller bind.ContractCaller) (*IPairedErc20Caller, error) {
	contract, err := bindIPairedErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPairedErc20Caller{contract: contract}, nil
}

// NewIPairedErc20Transactor creates a new write-only instance of IPairedErc20, bound to a specific deployed contract.
func NewIPairedErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*IPairedErc20Transactor, error) {
	contract, err := bindIPairedErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPairedErc20Transactor{contract: contract}, nil
}

// NewIPairedErc20Filterer creates a new log filterer instance of IPairedErc20, bound to a specific deployed contract.
func NewIPairedErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*IPairedErc20Filterer, error) {
	contract, err := bindIPairedErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPairedErc20Filterer{contract: contract}, nil
}

// bindIPairedErc20 binds a generic wrapper to an already deployed contract.
func bindIPairedErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPairedErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPairedErc20 *IPairedErc20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPairedErc20.Contract.IPairedErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPairedErc20 *IPairedErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPairedErc20.Contract.IPairedErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPairedErc20 *IPairedErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPairedErc20.Contract.IPairedErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPairedErc20 *IPairedErc20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPairedErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPairedErc20 *IPairedErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPairedErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPairedErc20 *IPairedErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPairedErc20.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IPairedErc20 *IPairedErc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IPairedErc20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IPairedErc20 *IPairedErc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IPairedErc20.Contract.BalanceOf(&_IPairedErc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IPairedErc20 *IPairedErc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IPairedErc20.Contract.BalanceOf(&_IPairedErc20.CallOpts, account)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_IPairedErc20 *IPairedErc20Transactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPairedErc20.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_IPairedErc20 *IPairedErc20Session) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPairedErc20.Contract.Burn(&_IPairedErc20.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_IPairedErc20 *IPairedErc20TransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPairedErc20.Contract.Burn(&_IPairedErc20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IPairedErc20 *IPairedErc20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPairedErc20.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IPairedErc20 *IPairedErc20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPairedErc20.Contract.Mint(&_IPairedErc20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IPairedErc20 *IPairedErc20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPairedErc20.Contract.Mint(&_IPairedErc20.TransactOpts, account, amount)
}

// PaymentRecordsABI is the input ABI used to generate the binding from.
const PaymentRecordsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"PaymentTransfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"getPaymentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"transferPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PaymentRecordsFuncSigs maps the 4-byte function signature to its string representation.
var PaymentRecordsFuncSigs = map[string]string{
	"0547e1b4": "getPaymentOwner(address,uint256)",
	"e628c91c": "transferPayment(address,address,uint256)",
}

// PaymentRecordsBin is the compiled bytecode used for deploying new contracts.
var PaymentRecordsBin = "0x608060405234801561001057600080fd5b50610272806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630547e1b41461003b578063e628c91c14610083575b600080fd5b6100676004803603604081101561005157600080fd5b506001600160a01b0381351690602001356100bb565b604080516001600160a01b039092168252519081900360200190f35b6100b96004803603606081101561009957600080fd5b506001600160a01b03813581169160208101359091169060400135610126565b005b6040805160208082018490526bffffffffffffffffffffffff19606086901b16828401528251603481840301815260549092018352815191810191909120600090815290819052908120546001600160a01b03168061011d5783915050610120565b90505b92915050565b600061013284836100bb565b9050336001600160a01b0382161461018a576040805162461bcd60e51b815260206004820152601660248201527526bab9ba103132903830bcb6b2b73a1037bbb732b91760511b604482015290519081900360640190fd5b6040805160208082018590526bffffffffffffffffffffffff19606088901b168284015282516034818403018152605483018085528151918301919091206000908152918290529083902080546001600160a01b038089166001600160a01b031990921682179092559186905280881660748401528416609483015260b482015290517f56e2b14cbe3f09b8a7337ab17e6ce5da744339e435db498d5ba53655695c03e59181900360d40190a15050505056fea265627a7a72315820a2beb2d79241b65fa56bafe45686c5214b1a1245461af8289d58a73de082435664736f6c63430005110032"

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

// GetPaymentOwner is a free data retrieval call binding the contract method 0x0547e1b4.
//
// Solidity: function getPaymentOwner(address originalOwner, uint256 messageIndex) view returns(address)
func (_PaymentRecords *PaymentRecordsCaller) GetPaymentOwner(opts *bind.CallOpts, originalOwner common.Address, messageIndex *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaymentRecords.contract.Call(opts, out, "getPaymentOwner", originalOwner, messageIndex)
	return *ret0, err
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0x0547e1b4.
//
// Solidity: function getPaymentOwner(address originalOwner, uint256 messageIndex) view returns(address)
func (_PaymentRecords *PaymentRecordsSession) GetPaymentOwner(originalOwner common.Address, messageIndex *big.Int) (common.Address, error) {
	return _PaymentRecords.Contract.GetPaymentOwner(&_PaymentRecords.CallOpts, originalOwner, messageIndex)
}

// GetPaymentOwner is a free data retrieval call binding the contract method 0x0547e1b4.
//
// Solidity: function getPaymentOwner(address originalOwner, uint256 messageIndex) view returns(address)
func (_PaymentRecords *PaymentRecordsCallerSession) GetPaymentOwner(originalOwner common.Address, messageIndex *big.Int) (common.Address, error) {
	return _PaymentRecords.Contract.GetPaymentOwner(&_PaymentRecords.CallOpts, originalOwner, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xe628c91c.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, uint256 messageIndex) returns()
func (_PaymentRecords *PaymentRecordsTransactor) TransferPayment(opts *bind.TransactOpts, originalOwner common.Address, newOwner common.Address, messageIndex *big.Int) (*types.Transaction, error) {
	return _PaymentRecords.contract.Transact(opts, "transferPayment", originalOwner, newOwner, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xe628c91c.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, uint256 messageIndex) returns()
func (_PaymentRecords *PaymentRecordsSession) TransferPayment(originalOwner common.Address, newOwner common.Address, messageIndex *big.Int) (*types.Transaction, error) {
	return _PaymentRecords.Contract.TransferPayment(&_PaymentRecords.TransactOpts, originalOwner, newOwner, messageIndex)
}

// TransferPayment is a paid mutator transaction binding the contract method 0xe628c91c.
//
// Solidity: function transferPayment(address originalOwner, address newOwner, uint256 messageIndex) returns()
func (_PaymentRecords *PaymentRecordsTransactorSession) TransferPayment(originalOwner common.Address, newOwner common.Address, messageIndex *big.Int) (*types.Transaction, error) {
	return _PaymentRecords.Contract.TransferPayment(&_PaymentRecords.TransactOpts, originalOwner, newOwner, messageIndex)
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
	MessageIndex  *big.Int
	OriginalOwner common.Address
	PrevOwner     common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentTransfer is a free log retrieval operation binding the contract event 0x56e2b14cbe3f09b8a7337ab17e6ce5da744339e435db498d5ba53655695c03e5.
//
// Solidity: event PaymentTransfer(uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_PaymentRecords *PaymentRecordsFilterer) FilterPaymentTransfer(opts *bind.FilterOpts) (*PaymentRecordsPaymentTransferIterator, error) {

	logs, sub, err := _PaymentRecords.contract.FilterLogs(opts, "PaymentTransfer")
	if err != nil {
		return nil, err
	}
	return &PaymentRecordsPaymentTransferIterator{contract: _PaymentRecords.contract, event: "PaymentTransfer", logs: logs, sub: sub}, nil
}

// WatchPaymentTransfer is a free log subscription operation binding the contract event 0x56e2b14cbe3f09b8a7337ab17e6ce5da744339e435db498d5ba53655695c03e5.
//
// Solidity: event PaymentTransfer(uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
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

// ParsePaymentTransfer is a log parse operation binding the contract event 0x56e2b14cbe3f09b8a7337ab17e6ce5da744339e435db498d5ba53655695c03e5.
//
// Solidity: event PaymentTransfer(uint256 messageIndex, address originalOwner, address prevOwner, address newOwner)
func (_PaymentRecords *PaymentRecordsFilterer) ParsePaymentTransfer(log types.Log) (*PaymentRecordsPaymentTransfer, error) {
	event := new(PaymentRecordsPaymentTransfer)
	if err := _PaymentRecords.contract.UnpackLog(event, "PaymentTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
