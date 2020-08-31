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

// DistributionsV0ABI is the input ABI used to generate the binding from.
const DistributionsV0ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subredditPoints_\",\"type\":\"address\"}],\"name\":\"instantiateContractTest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"subreddit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testVar\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DistributionsV0FuncSigs maps the 4-byte function signature to its string representation.
var DistributionsV0FuncSigs = map[string]string{
	"378dc3dc": "initialSupply()",
	"1c2a2551": "instantiateContractTest(address)",
	"bdc330cb": "subreddit()",
	"9c328fb3": "testVar()",
}

// DistributionsV0Bin is the compiled bytecode used for deploying new contracts.
var DistributionsV0Bin = "0x60c0604052600e60808190526d6f726967696e616c2076616c756560901b60a090815261002f9160029190610042565b5034801561003c57600080fd5b506100dd565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008357805160ff19168380011785556100b0565b828001600101855582156100b0579182015b828111156100b0578251825591602001919060010190610095565b506100bc9291506100c0565b5090565b6100da91905b808211156100bc57600081556001016100c6565b90565b610445806100ec6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631c2a255114610051578063378dc3dc146100795780639c328fb314610093578063bdc330cb14610110575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b0316610118565b005b610081610289565b60408051918252519081900360200190f35b61009b61028f565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100d55781810151838201526020016100bd565b50505050905090810190601f1680156101025780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61009b61031a565b806001600160a01b031663bdc330cb6040518163ffffffff1660e01b815260040160006040518083038186803b15801561015157600080fd5b505afa158015610165573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561018e57600080fd5b81019080805160405193929190846401000000008211156101ae57600080fd5b9083019060208201858111156101c357600080fd5b82516401000000008111828201881017156101dd57600080fd5b82525081516020918201929091019080838360005b8381101561020a5781810151838201526020016101f2565b50505050905090810190601f1680156102375780820380516001836020036101000a031916815260200191505b5060405250508151610250926000925060200190610375565b5060408051808201909152600d8082526c757064617465642076616c756560981b602090920191825261028591600291610375565b5050565b60015481565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156103125780601f106102e757610100808354040283529160200191610312565b820191906000526020600020905b8154815290600101906020018083116102f557829003601f168201915b505050505081565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103125780601f106102e757610100808354040283529160200191610312565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106103b657805160ff19168380011785556103e3565b828001600101855582156103e3579182015b828111156103e35782518255916020019190600101906103c8565b506103ef9291506103f3565b5090565b61040d91905b808211156103ef57600081556001016103f9565b9056fea265627a7a7231582042acc0357bfb7ebdaf71e3186d1ac5553962f3b80d260b728f02e7d74afca76f64736f6c63430005110032"

// DeployDistributionsV0 deploys a new Ethereum contract, binding an instance of DistributionsV0 to it.
func DeployDistributionsV0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DistributionsV0, error) {
	parsed, err := abi.JSON(strings.NewReader(DistributionsV0ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DistributionsV0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DistributionsV0{DistributionsV0Caller: DistributionsV0Caller{contract: contract}, DistributionsV0Transactor: DistributionsV0Transactor{contract: contract}, DistributionsV0Filterer: DistributionsV0Filterer{contract: contract}}, nil
}

// DistributionsV0 is an auto generated Go binding around an Ethereum contract.
type DistributionsV0 struct {
	DistributionsV0Caller     // Read-only binding to the contract
	DistributionsV0Transactor // Write-only binding to the contract
	DistributionsV0Filterer   // Log filterer for contract events
}

// DistributionsV0Caller is an auto generated read-only Go binding around an Ethereum contract.
type DistributionsV0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionsV0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributionsV0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionsV0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributionsV0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionsV0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributionsV0Session struct {
	Contract     *DistributionsV0  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributionsV0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributionsV0CallerSession struct {
	Contract *DistributionsV0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DistributionsV0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributionsV0TransactorSession struct {
	Contract     *DistributionsV0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DistributionsV0Raw is an auto generated low-level Go binding around an Ethereum contract.
type DistributionsV0Raw struct {
	Contract *DistributionsV0 // Generic contract binding to access the raw methods on
}

// DistributionsV0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributionsV0CallerRaw struct {
	Contract *DistributionsV0Caller // Generic read-only contract binding to access the raw methods on
}

// DistributionsV0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributionsV0TransactorRaw struct {
	Contract *DistributionsV0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributionsV0 creates a new instance of DistributionsV0, bound to a specific deployed contract.
func NewDistributionsV0(address common.Address, backend bind.ContractBackend) (*DistributionsV0, error) {
	contract, err := bindDistributionsV0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DistributionsV0{DistributionsV0Caller: DistributionsV0Caller{contract: contract}, DistributionsV0Transactor: DistributionsV0Transactor{contract: contract}, DistributionsV0Filterer: DistributionsV0Filterer{contract: contract}}, nil
}

// NewDistributionsV0Caller creates a new read-only instance of DistributionsV0, bound to a specific deployed contract.
func NewDistributionsV0Caller(address common.Address, caller bind.ContractCaller) (*DistributionsV0Caller, error) {
	contract, err := bindDistributionsV0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionsV0Caller{contract: contract}, nil
}

// NewDistributionsV0Transactor creates a new write-only instance of DistributionsV0, bound to a specific deployed contract.
func NewDistributionsV0Transactor(address common.Address, transactor bind.ContractTransactor) (*DistributionsV0Transactor, error) {
	contract, err := bindDistributionsV0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionsV0Transactor{contract: contract}, nil
}

// NewDistributionsV0Filterer creates a new log filterer instance of DistributionsV0, bound to a specific deployed contract.
func NewDistributionsV0Filterer(address common.Address, filterer bind.ContractFilterer) (*DistributionsV0Filterer, error) {
	contract, err := bindDistributionsV0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributionsV0Filterer{contract: contract}, nil
}

// bindDistributionsV0 binds a generic wrapper to an already deployed contract.
func bindDistributionsV0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistributionsV0ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistributionsV0 *DistributionsV0Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DistributionsV0.Contract.DistributionsV0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistributionsV0 *DistributionsV0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistributionsV0.Contract.DistributionsV0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistributionsV0 *DistributionsV0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistributionsV0.Contract.DistributionsV0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistributionsV0 *DistributionsV0CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DistributionsV0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistributionsV0 *DistributionsV0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistributionsV0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistributionsV0 *DistributionsV0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistributionsV0.Contract.contract.Transact(opts, method, params...)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_DistributionsV0 *DistributionsV0Caller) InitialSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DistributionsV0.contract.Call(opts, out, "initialSupply")
	return *ret0, err
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_DistributionsV0 *DistributionsV0Session) InitialSupply() (*big.Int, error) {
	return _DistributionsV0.Contract.InitialSupply(&_DistributionsV0.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_DistributionsV0 *DistributionsV0CallerSession) InitialSupply() (*big.Int, error) {
	return _DistributionsV0.Contract.InitialSupply(&_DistributionsV0.CallOpts)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_DistributionsV0 *DistributionsV0Caller) Subreddit(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DistributionsV0.contract.Call(opts, out, "subreddit")
	return *ret0, err
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_DistributionsV0 *DistributionsV0Session) Subreddit() (string, error) {
	return _DistributionsV0.Contract.Subreddit(&_DistributionsV0.CallOpts)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_DistributionsV0 *DistributionsV0CallerSession) Subreddit() (string, error) {
	return _DistributionsV0.Contract.Subreddit(&_DistributionsV0.CallOpts)
}

// TestVar is a free data retrieval call binding the contract method 0x9c328fb3.
//
// Solidity: function testVar() view returns(string)
func (_DistributionsV0 *DistributionsV0Caller) TestVar(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DistributionsV0.contract.Call(opts, out, "testVar")
	return *ret0, err
}

// TestVar is a free data retrieval call binding the contract method 0x9c328fb3.
//
// Solidity: function testVar() view returns(string)
func (_DistributionsV0 *DistributionsV0Session) TestVar() (string, error) {
	return _DistributionsV0.Contract.TestVar(&_DistributionsV0.CallOpts)
}

// TestVar is a free data retrieval call binding the contract method 0x9c328fb3.
//
// Solidity: function testVar() view returns(string)
func (_DistributionsV0 *DistributionsV0CallerSession) TestVar() (string, error) {
	return _DistributionsV0.Contract.TestVar(&_DistributionsV0.CallOpts)
}

// InstantiateContractTest is a paid mutator transaction binding the contract method 0x1c2a2551.
//
// Solidity: function instantiateContractTest(address subredditPoints_) returns()
func (_DistributionsV0 *DistributionsV0Transactor) InstantiateContractTest(opts *bind.TransactOpts, subredditPoints_ common.Address) (*types.Transaction, error) {
	return _DistributionsV0.contract.Transact(opts, "instantiateContractTest", subredditPoints_)
}

// InstantiateContractTest is a paid mutator transaction binding the contract method 0x1c2a2551.
//
// Solidity: function instantiateContractTest(address subredditPoints_) returns()
func (_DistributionsV0 *DistributionsV0Session) InstantiateContractTest(subredditPoints_ common.Address) (*types.Transaction, error) {
	return _DistributionsV0.Contract.InstantiateContractTest(&_DistributionsV0.TransactOpts, subredditPoints_)
}

// InstantiateContractTest is a paid mutator transaction binding the contract method 0x1c2a2551.
//
// Solidity: function instantiateContractTest(address subredditPoints_) returns()
func (_DistributionsV0 *DistributionsV0TransactorSession) InstantiateContractTest(subredditPoints_ common.Address) (*types.Transaction, error) {
	return _DistributionsV0.Contract.InstantiateContractTest(&_DistributionsV0.TransactOpts, subredditPoints_)
}

// ISubredditPointsABI is the input ABI used to generate the binding from.
const ISubredditPointsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"subreddit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ISubredditPointsFuncSigs maps the 4-byte function signature to its string representation.
var ISubredditPointsFuncSigs = map[string]string{
	"bdc330cb": "subreddit()",
}

// ISubredditPoints is an auto generated Go binding around an Ethereum contract.
type ISubredditPoints struct {
	ISubredditPointsCaller     // Read-only binding to the contract
	ISubredditPointsTransactor // Write-only binding to the contract
	ISubredditPointsFilterer   // Log filterer for contract events
}

// ISubredditPointsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISubredditPointsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubredditPointsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISubredditPointsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubredditPointsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISubredditPointsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubredditPointsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISubredditPointsSession struct {
	Contract     *ISubredditPoints // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISubredditPointsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISubredditPointsCallerSession struct {
	Contract *ISubredditPointsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ISubredditPointsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISubredditPointsTransactorSession struct {
	Contract     *ISubredditPointsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ISubredditPointsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISubredditPointsRaw struct {
	Contract *ISubredditPoints // Generic contract binding to access the raw methods on
}

// ISubredditPointsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISubredditPointsCallerRaw struct {
	Contract *ISubredditPointsCaller // Generic read-only contract binding to access the raw methods on
}

// ISubredditPointsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISubredditPointsTransactorRaw struct {
	Contract *ISubredditPointsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISubredditPoints creates a new instance of ISubredditPoints, bound to a specific deployed contract.
func NewISubredditPoints(address common.Address, backend bind.ContractBackend) (*ISubredditPoints, error) {
	contract, err := bindISubredditPoints(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISubredditPoints{ISubredditPointsCaller: ISubredditPointsCaller{contract: contract}, ISubredditPointsTransactor: ISubredditPointsTransactor{contract: contract}, ISubredditPointsFilterer: ISubredditPointsFilterer{contract: contract}}, nil
}

// NewISubredditPointsCaller creates a new read-only instance of ISubredditPoints, bound to a specific deployed contract.
func NewISubredditPointsCaller(address common.Address, caller bind.ContractCaller) (*ISubredditPointsCaller, error) {
	contract, err := bindISubredditPoints(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISubredditPointsCaller{contract: contract}, nil
}

// NewISubredditPointsTransactor creates a new write-only instance of ISubredditPoints, bound to a specific deployed contract.
func NewISubredditPointsTransactor(address common.Address, transactor bind.ContractTransactor) (*ISubredditPointsTransactor, error) {
	contract, err := bindISubredditPoints(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISubredditPointsTransactor{contract: contract}, nil
}

// NewISubredditPointsFilterer creates a new log filterer instance of ISubredditPoints, bound to a specific deployed contract.
func NewISubredditPointsFilterer(address common.Address, filterer bind.ContractFilterer) (*ISubredditPointsFilterer, error) {
	contract, err := bindISubredditPoints(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISubredditPointsFilterer{contract: contract}, nil
}

// bindISubredditPoints binds a generic wrapper to an already deployed contract.
func bindISubredditPoints(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISubredditPointsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubredditPoints *ISubredditPointsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubredditPoints.Contract.ISubredditPointsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubredditPoints *ISubredditPointsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubredditPoints.Contract.ISubredditPointsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubredditPoints *ISubredditPointsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubredditPoints.Contract.ISubredditPointsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubredditPoints *ISubredditPointsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubredditPoints.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubredditPoints *ISubredditPointsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubredditPoints.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubredditPoints *ISubredditPointsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubredditPoints.Contract.contract.Transact(opts, method, params...)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_ISubredditPoints *ISubredditPointsCaller) Subreddit(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ISubredditPoints.contract.Call(opts, out, "subreddit")
	return *ret0, err
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_ISubredditPoints *ISubredditPointsSession) Subreddit() (string, error) {
	return _ISubredditPoints.Contract.Subreddit(&_ISubredditPoints.CallOpts)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() view returns(string)
func (_ISubredditPoints *ISubredditPointsCallerSession) Subreddit() (string, error) {
	return _ISubredditPoints.Contract.Subreddit(&_ISubredditPoints.CallOpts)
}
