// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sigutilstester

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

// SigUtilsABI is the input ABI used to generate the binding from.
const SigUtilsABI = "[]"

// SigUtilsBin is the compiled bytecode used for deploying new contracts.
var SigUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820717b5c7ed64eb86680a9f6c54e47a6181806003e0d74370625e19e35e33548c964736f6c63430005110032"

// DeploySigUtils deploys a new Ethereum contract, binding an instance of SigUtils to it.
func DeploySigUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// SigUtils is an auto generated Go binding around an Ethereum contract.
type SigUtils struct {
	SigUtilsCaller     // Read-only binding to the contract
	SigUtilsTransactor // Write-only binding to the contract
	SigUtilsFilterer   // Log filterer for contract events
}

// SigUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsSession struct {
	Contract     *SigUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsCallerSession struct {
	Contract *SigUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTransactorSession struct {
	Contract     *SigUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsRaw struct {
	Contract *SigUtils // Generic contract binding to access the raw methods on
}

// SigUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsCallerRaw struct {
	Contract *SigUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTransactorRaw struct {
	Contract *SigUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtils creates a new instance of SigUtils, bound to a specific deployed contract.
func NewSigUtils(address common.Address, backend bind.ContractBackend) (*SigUtils, error) {
	contract, err := bindSigUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// NewSigUtilsCaller creates a new read-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsCaller, error) {
	contract, err := bindSigUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsCaller{contract: contract}, nil
}

// NewSigUtilsTransactor creates a new write-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTransactor, error) {
	contract, err := bindSigUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTransactor{contract: contract}, nil
}

// NewSigUtilsFilterer creates a new log filterer instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsFilterer, error) {
	contract, err := bindSigUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsFilterer{contract: contract}, nil
}

// bindSigUtils binds a generic wrapper to an already deployed contract.
func bindSigUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.SigUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transact(opts, method, params...)
}

// SigUtilsTesterABI is the input ABI used to generate the binding from.
const SigUtilsTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"}],\"name\":\"parseSignature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"}],\"name\":\"recoverAddressFromData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SigUtilsTesterFuncSigs maps the 4-byte function signature to its string representation.
var SigUtilsTesterFuncSigs = map[string]string{
	"b31d63cc": "parseSignature(bytes,uint256)",
	"d4916333": "recoverAddressFromData(bytes32,bytes,uint256)",
}

// SigUtilsTesterBin is the compiled bytecode used for deploying new contracts.
var SigUtilsTesterBin = "0x608060405234801561001057600080fd5b506103f7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063b31d63cc1461003b578063d491633314610105575b600080fd5b6100e36004803603604081101561005157600080fd5b81019060208101813564010000000081111561006c57600080fd5b82018360208201111561007e57600080fd5b803590602001918460018302840111640100000000831117156100a057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506101d0915050565b6040805160ff9094168452602084019290925282820152519081900360600190f35b6101b46004803603606081101561011b57600080fd5b8135919081019060408101602082013564010000000081111561013d57600080fd5b82018360208201111561014f57600080fd5b8035906020019184600183028401116401000000008311171561017157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506101ec915050565b604080516001600160a01b039092168252519081900360200190f35b60008060006101df8585610201565b9250925092509250925092565b60006101f984848461028f565b949350505050565b818101602081810151604083015160609093015160001a9290918401601b84101561022d57601b840193505b8360ff16601b148061024257508360ff16601c145b610287576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081896040516020018083805190602001908083835b602083106103055780518252601f1990920191602091820191016102e6565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190528251920191909120925061034b915089905088610201565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa1580156103aa573d6000803e3d6000fd5b5050604051601f1901519a995050505050505050505056fea265627a7a7231582005bb3a47757843b07399923074c6f7838f0e2c8db1106e8633bd2b0bab94154764736f6c63430005110032"

// DeploySigUtilsTester deploys a new Ethereum contract, binding an instance of SigUtilsTester to it.
func DeploySigUtilsTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtilsTester, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtilsTester{SigUtilsTesterCaller: SigUtilsTesterCaller{contract: contract}, SigUtilsTesterTransactor: SigUtilsTesterTransactor{contract: contract}, SigUtilsTesterFilterer: SigUtilsTesterFilterer{contract: contract}}, nil
}

// SigUtilsTester is an auto generated Go binding around an Ethereum contract.
type SigUtilsTester struct {
	SigUtilsTesterCaller     // Read-only binding to the contract
	SigUtilsTesterTransactor // Write-only binding to the contract
	SigUtilsTesterFilterer   // Log filterer for contract events
}

// SigUtilsTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsTesterSession struct {
	Contract     *SigUtilsTester   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsTesterCallerSession struct {
	Contract *SigUtilsTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SigUtilsTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTesterTransactorSession struct {
	Contract     *SigUtilsTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SigUtilsTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsTesterRaw struct {
	Contract *SigUtilsTester // Generic contract binding to access the raw methods on
}

// SigUtilsTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsTesterCallerRaw struct {
	Contract *SigUtilsTesterCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTesterTransactorRaw struct {
	Contract *SigUtilsTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtilsTester creates a new instance of SigUtilsTester, bound to a specific deployed contract.
func NewSigUtilsTester(address common.Address, backend bind.ContractBackend) (*SigUtilsTester, error) {
	contract, err := bindSigUtilsTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTester{SigUtilsTesterCaller: SigUtilsTesterCaller{contract: contract}, SigUtilsTesterTransactor: SigUtilsTesterTransactor{contract: contract}, SigUtilsTesterFilterer: SigUtilsTesterFilterer{contract: contract}}, nil
}

// NewSigUtilsTesterCaller creates a new read-only instance of SigUtilsTester, bound to a specific deployed contract.
func NewSigUtilsTesterCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsTesterCaller, error) {
	contract, err := bindSigUtilsTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTesterCaller{contract: contract}, nil
}

// NewSigUtilsTesterTransactor creates a new write-only instance of SigUtilsTester, bound to a specific deployed contract.
func NewSigUtilsTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTesterTransactor, error) {
	contract, err := bindSigUtilsTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTesterTransactor{contract: contract}, nil
}

// NewSigUtilsTesterFilterer creates a new log filterer instance of SigUtilsTester, bound to a specific deployed contract.
func NewSigUtilsTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsTesterFilterer, error) {
	contract, err := bindSigUtilsTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTesterFilterer{contract: contract}, nil
}

// bindSigUtilsTester binds a generic wrapper to an already deployed contract.
func bindSigUtilsTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtilsTester *SigUtilsTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtilsTester.Contract.SigUtilsTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtilsTester *SigUtilsTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtilsTester.Contract.SigUtilsTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtilsTester *SigUtilsTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtilsTester.Contract.SigUtilsTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtilsTester *SigUtilsTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtilsTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtilsTester *SigUtilsTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtilsTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtilsTester *SigUtilsTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtilsTester.Contract.contract.Transact(opts, method, params...)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _data, uint256 _start) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtilsTester *SigUtilsTesterCaller) ParseSignature(opts *bind.CallOpts, _data []byte, _start *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	ret := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	out := ret
	err := _SigUtilsTester.contract.Call(opts, out, "parseSignature", _data, _start)
	return *ret, err
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _data, uint256 _start) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtilsTester *SigUtilsTesterSession) ParseSignature(_data []byte, _start *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigUtilsTester.Contract.ParseSignature(&_SigUtilsTester.CallOpts, _data, _start)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _data, uint256 _start) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtilsTester *SigUtilsTesterCallerSession) ParseSignature(_data []byte, _start *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigUtilsTester.Contract.ParseSignature(&_SigUtilsTester.CallOpts, _data, _start)
}

// RecoverAddressFromData is a free data retrieval call binding the contract method 0xd4916333.
//
// Solidity: function recoverAddressFromData(bytes32 _messageHash, bytes _data, uint256 _offset) pure returns(address)
func (_SigUtilsTester *SigUtilsTesterCaller) RecoverAddressFromData(opts *bind.CallOpts, _messageHash [32]byte, _data []byte, _offset *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SigUtilsTester.contract.Call(opts, out, "recoverAddressFromData", _messageHash, _data, _offset)
	return *ret0, err
}

// RecoverAddressFromData is a free data retrieval call binding the contract method 0xd4916333.
//
// Solidity: function recoverAddressFromData(bytes32 _messageHash, bytes _data, uint256 _offset) pure returns(address)
func (_SigUtilsTester *SigUtilsTesterSession) RecoverAddressFromData(_messageHash [32]byte, _data []byte, _offset *big.Int) (common.Address, error) {
	return _SigUtilsTester.Contract.RecoverAddressFromData(&_SigUtilsTester.CallOpts, _messageHash, _data, _offset)
}

// RecoverAddressFromData is a free data retrieval call binding the contract method 0xd4916333.
//
// Solidity: function recoverAddressFromData(bytes32 _messageHash, bytes _data, uint256 _offset) pure returns(address)
func (_SigUtilsTester *SigUtilsTesterCallerSession) RecoverAddressFromData(_messageHash [32]byte, _data []byte, _offset *big.Int) (common.Address, error) {
	return _SigUtilsTester.Contract.RecoverAddressFromData(&_SigUtilsTester.CallOpts, _messageHash, _data, _offset)
}
