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

// ValidatorABI is the input ABI used to generate the binding from.
const ValidatorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"executeTransactions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"returnOldDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallenge[]\",\"name\":\"challenges\",\"type\":\"address[]\"}],\"name\":\"timeoutChallenges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorBin = "0x608060405234801561001057600080fd5b506065805460ff19166001179055610dec8061002d6000396000f3fe6080604052600436106100765760003560e01c80636f791d291461007b578063715018a6146100a657806372f45866146100bd5780638129fc1c146100d057806381aac2d9146100e55780638da5cb5b14610105578063944f449514610127578063ce1d571f14610147578063f2fde38b1461015a575b600080fd5b34801561008757600080fd5b5061009061017a565b60405161009d9190610c23565b60405180910390f35b3480156100b257600080fd5b506100bb610183565b005b6100bb6100cb366004610a56565b610203565b3480156100dc57600080fd5b506100bb61039a565b3480156100f157600080fd5b506100bb610100366004610aeb565b610425565b34801561011157600080fd5b5061011a61053e565b60405161009d9190610c0f565b34801561013357600080fd5b506100bb610142366004610bad565b61054d565b6100bb610155366004610b2a565b610672565b34801561016657600080fd5b506100bb610175366004610a33565b61075a565b60655460ff1690565b61018b610809565b6001600160a01b031661019c61053e565b6001600160a01b0316146101cb5760405162461bcd60e51b81526004016101c290610ceb565b60405180910390fd5b6033546040516000916001600160a01b031690600080516020610d97833981519152908390a3603380546001600160a01b0319169055565b61020b610809565b6001600160a01b031661021c61053e565b6001600160a01b0316146102425760405162461bcd60e51b81526004016101c290610ceb565b8460005b8181101561039057600088888381811061025c57fe5b905060200281019061026e9190610d3d565b905011156102c4576102a886868381811061028557fe5b905060200201602081019061029a9190610a33565b6001600160a01b031661080d565b6102c45760405162461bcd60e51b81526004016101c290610c2e565b60008686838181106102d257fe5b90506020020160208101906102e79190610a33565b6001600160a01b03168585848181106102fc57fe5b905060200201358a8a8581811061030f57fe5b90506020028101906103219190610d3d565b60405161032f929190610bff565b60006040518083038185875af1925050503d806000811461036c576040519150601f19603f3d011682016040523d82523d6000602084013e610371565b606091505b5050905080610387576040513d806000833e8082fd5b50600101610246565b5050505050505050565b600054610100900460ff16806103b357506103b3610813565b806103c1575060005460ff16155b6103dd5760405162461bcd60e51b81526004016101c290610c9d565b600054610100900460ff16158015610408576000805460ff1961ff0019909116610100171660011790555b610410610824565b8015610422576000805461ff00191690555b50565b61042d610809565b6001600160a01b031661043e61053e565b6001600160a01b0316146104645760405162461bcd60e51b81526004016101c290610ceb565b8060005b818110156105385783838281811061047c57fe5b90506020020160208101906104919190610a33565b6001600160a01b03166370dea79a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156104cb57600080fd5b505af19250505080156104dc575060015b610530573d80801561050a576040519150601f19603f3d011682016040523d82523d6000602084013e61050f565b606091505b50805161052e5760405162461bcd60e51b81526004016101c290610d20565b505b600101610468565b50505050565b6033546001600160a01b031690565b610555610809565b6001600160a01b031661056661053e565b6001600160a01b03161461058c5760405162461bcd60e51b81526004016101c290610ceb565b8060005b8181101561066b57846001600160a01b0316637427be518585848181106105b357fe5b90506020020160208101906105c89190610a33565b6040518263ffffffff1660e01b81526004016105e49190610c0f565b600060405180830381600087803b1580156105fe57600080fd5b505af192505050801561060f575060015b610663573d80801561063d576040519150601f19603f3d011682016040523d82523d6000602084013e610642565b606091505b5080516106615760405162461bcd60e51b81526004016101c290610d20565b505b600101610590565b5050505050565b61067a610809565b6001600160a01b031661068b61053e565b6001600160a01b0316146106b15760405162461bcd60e51b81526004016101c290610ceb565b82156106e5576106c9826001600160a01b031661080d565b6106e55760405162461bcd60e51b81526004016101c290610c2e565b6000826001600160a01b0316828686604051610702929190610bff565b60006040518083038185875af1925050503d806000811461073f576040519150601f19603f3d011682016040523d82523d6000602084013e610744565b606091505b505090508061066b576040513d806000833e8082fd5b610762610809565b6001600160a01b031661077361053e565b6001600160a01b0316146107995760405162461bcd60e51b81526004016101c290610ceb565b6001600160a01b0381166107bf5760405162461bcd60e51b81526004016101c290610c57565b6033546040516001600160a01b03808416921690600080516020610d9783398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b3b151590565b600061081e3061080d565b15905090565b600054610100900460ff168061083d575061083d610813565b8061084b575060005460ff16155b6108675760405162461bcd60e51b81526004016101c290610c9d565b600054610100900460ff16158015610892576000805460ff1961ff0019909116610100171660011790555b61089a6108a2565b610410610923565b600054610100900460ff16806108bb57506108bb610813565b806108c9575060005460ff16155b6108e55760405162461bcd60e51b81526004016101c290610c9d565b600054610100900460ff16158015610410576000805460ff1961ff0019909116610100171660011790558015610422576000805461ff001916905550565b600054610100900460ff168061093c575061093c610813565b8061094a575060005460ff16155b6109665760405162461bcd60e51b81526004016101c290610c9d565b600054610100900460ff16158015610991576000805460ff1961ff0019909116610100171660011790555b600061099b610809565b603380546001600160a01b0319166001600160a01b03831690811790915560405191925090600090600080516020610d97833981519152908290a3508015610422576000805461ff001916905550565b60008083601f8401126109fc578182fd5b5081356001600160401b03811115610a12578182fd5b6020830191508360208083028501011115610a2c57600080fd5b9250929050565b600060208284031215610a44578081fd5b8135610a4f81610d81565b9392505050565b60008060008060008060608789031215610a6e578182fd5b86356001600160401b0380821115610a84578384fd5b610a908a838b016109eb565b90985096506020890135915080821115610aa8578384fd5b610ab48a838b016109eb565b90965094506040890135915080821115610acc578384fd5b50610ad989828a016109eb565b979a9699509497509295939492505050565b60008060208385031215610afd578182fd5b82356001600160401b03811115610b12578283fd5b610b1e858286016109eb565b90969095509350505050565b60008060008060608587031215610b3f578384fd5b84356001600160401b0380821115610b55578586fd5b818701915087601f830112610b68578586fd5b813581811115610b76578687fd5b886020828501011115610b87578687fd5b60209283019650945050850135610b9d81610d81565b9396929550929360400135925050565b600080600060408486031215610bc1578283fd5b8335610bcc81610d81565b925060208401356001600160401b03811115610be6578283fd5b610bf2868287016109eb565b9497909650939450505050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b901515815260200190565b6020808252600f908201526e2727afa1a7a222afa0aa2fa0a2222960891b604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526003908201526247415360e81b604082015260600190565b6000808335601e19843603018112610d53578283fd5b8301803591506001600160401b03821115610d6c578283fd5b602001915036819003821315610a2c57600080fd5b6001600160a01b038116811461042257600080fdfe8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220c32f3e207e1ba826a701c5a23fab40d10471e63cefa1871f63062a427f95a7e364736f6c634300060c0033"

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Validator *ValidatorCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Validator *ValidatorSession) IsMaster() (bool, error) {
	return _Validator.Contract.IsMaster(&_Validator.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Validator *ValidatorCallerSession) IsMaster() (bool, error) {
	return _Validator.Contract.IsMaster(&_Validator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validator *ValidatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validator *ValidatorSession) Owner() (common.Address, error) {
	return _Validator.Contract.Owner(&_Validator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validator *ValidatorCallerSession) Owner() (common.Address, error) {
	return _Validator.Contract.Owner(&_Validator.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorTransactor) ExecuteTransaction(opts *bind.TransactOpts, data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "executeTransaction", data, destination, amount)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorSession) ExecuteTransaction(data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransaction(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xce1d571f.
//
// Solidity: function executeTransaction(bytes data, address destination, uint256 amount) payable returns()
func (_Validator *ValidatorTransactorSession) ExecuteTransaction(data []byte, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransaction(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorTransactor) ExecuteTransactions(opts *bind.TransactOpts, data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "executeTransactions", data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorSession) ExecuteTransactions(data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactions(&_Validator.TransactOpts, data, destination, amount)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0x72f45866.
//
// Solidity: function executeTransactions(bytes[] data, address[] destination, uint256[] amount) payable returns()
func (_Validator *ValidatorTransactorSession) ExecuteTransactions(data [][]byte, destination []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactions(&_Validator.TransactOpts, data, destination, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Validator *ValidatorTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Validator *ValidatorSession) Initialize() (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Validator *ValidatorTransactorSession) Initialize() (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validator *ValidatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validator *ValidatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validator.Contract.RenounceOwnership(&_Validator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validator *ValidatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validator.Contract.RenounceOwnership(&_Validator.TransactOpts)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorTransactor) ReturnOldDeposits(opts *bind.TransactOpts, rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "returnOldDeposits", rollup, stakers)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorSession) ReturnOldDeposits(rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ReturnOldDeposits(&_Validator.TransactOpts, rollup, stakers)
}

// ReturnOldDeposits is a paid mutator transaction binding the contract method 0x944f4495.
//
// Solidity: function returnOldDeposits(address rollup, address[] stakers) returns()
func (_Validator *ValidatorTransactorSession) ReturnOldDeposits(rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ReturnOldDeposits(&_Validator.TransactOpts, rollup, stakers)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x81aac2d9.
//
// Solidity: function timeoutChallenges(address[] challenges) returns()
func (_Validator *ValidatorTransactor) TimeoutChallenges(opts *bind.TransactOpts, challenges []common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "timeoutChallenges", challenges)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x81aac2d9.
//
// Solidity: function timeoutChallenges(address[] challenges) returns()
func (_Validator *ValidatorSession) TimeoutChallenges(challenges []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TimeoutChallenges(&_Validator.TransactOpts, challenges)
}

// TimeoutChallenges is a paid mutator transaction binding the contract method 0x81aac2d9.
//
// Solidity: function timeoutChallenges(address[] challenges) returns()
func (_Validator *ValidatorTransactorSession) TimeoutChallenges(challenges []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TimeoutChallenges(&_Validator.TransactOpts, challenges)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validator *ValidatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validator *ValidatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TransferOwnership(&_Validator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validator *ValidatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TransferOwnership(&_Validator.TransactOpts, newOwner)
}

// ValidatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Validator contract.
type ValidatorOwnershipTransferredIterator struct {
	Event *ValidatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorOwnershipTransferred)
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
		it.Event = new(ValidatorOwnershipTransferred)
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
func (it *ValidatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorOwnershipTransferred represents a OwnershipTransferred event raised by the Validator contract.
type ValidatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validator *ValidatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorOwnershipTransferredIterator{contract: _Validator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validator *ValidatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorOwnershipTransferred)
				if err := _Validator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validator *ValidatorFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorOwnershipTransferred, error) {
	event := new(ValidatorOwnershipTransferred)
	if err := _Validator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWalletCreatorABI is the input ABI used to generate the binding from.
const ValidatorWalletCreatorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TemplateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"adminProxy\",\"type\":\"address\"}],\"name\":\"WalletCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_template\",\"type\":\"address\"}],\"name\":\"setTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"template\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorWalletCreatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorWalletCreatorBin = "0x608060405234801561001057600080fd5b50600061001b6100b3565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350604051610071906100b7565b604051809103906000f08015801561008d573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b03929092169190911790556100c4565b3390565b610e1980611c1d83390190565b611b4a806100d36000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806311ebbf24146100675780636f2ddd931461008b578063715018a61461009357806389c716d11461009d5780638da5cb5b146100c3578063f2fde38b146100cb575b600080fd5b61006f6100f1565b604080516001600160a01b039092168252519081900360200190f35b61006f6102e9565b61009b6102f8565b005b61009b600480360360208110156100b357600080fd5b50356001600160a01b0316610392565b61006f61043b565b61009b600480360360208110156100e157600080fd5b50356001600160a01b031661044a565b6000806040516101009061053e565b604051809103906000f08015801561011c573d6000803e3d6000fd5b506001546040519192506000916001600160a01b039091169083906101409061054b565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f080158015610183573d6000803e3d6000fd5b509050816001600160a01b031663f2fde38b336040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156101d557600080fd5b505af11580156101e9573d6000803e3d6000fd5b50505050806001600160a01b0316638129fc1c6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561022857600080fd5b505af115801561023c573d6000803e3d6000fd5b50506040805163f2fde38b60e01b815233600482015290516001600160a01b038516935063f2fde38b9250602480830192600092919082900301818387803b15801561028757600080fd5b505af115801561029b573d6000803e3d6000fd5b5050604080516001600160a01b038681168252915133945091851692507fca0b7dde26052d34217ef1a0cee48085a07ca32da0a918609937a307d496bbf5919081900360200190a391505090565b6001546001600160a01b031681565b61030061053a565b6001600160a01b031661031161043b565b6001600160a01b03161461035a576040805162461bcd60e51b81526020600482018190526024820152600080516020611ad5833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020611af5833981519152908390a3600080546001600160a01b0319169055565b61039a61053a565b6001600160a01b03166103ab61043b565b6001600160a01b0316146103f4576040805162461bcd60e51b81526020600482018190526024820152600080516020611ad5833981519152604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0383161790556040517f6eb26f176dd9180849dd4874d3530de0e5c1f62a6e6798d34e3abfc11f1db2cc90600090a150565b6000546001600160a01b031690565b61045261053a565b6001600160a01b031661046361043b565b6001600160a01b0316146104ac576040805162461bcd60e51b81526020600482018190526024820152600080516020611ad5833981519152604482015290519081900360640190fd5b6001600160a01b0381166104f15760405162461bcd60e51b8152600401808060200182810382526026815260200180611aaf6026913960400191505060405180910390fd5b600080546040516001600160a01b0380851693921691600080516020611af583398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b6108e28061055983390190565b610c7480610e3b8339019056fe608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6108658061007d6000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100bf5780637eff275e146100d65780638da5cb5b146101115780639623609d1461012657806399a88ec4146101e3578063f2fde38b1461021e578063f3b7dead14610251575b600080fd5b34801561007c57600080fd5b506100a36004803603602081101561009357600080fd5b50356001600160a01b0316610284565b604080516001600160a01b039092168252519081900360200190f35b3480156100cb57600080fd5b506100d4610316565b005b3480156100e257600080fd5b506100d4600480360360408110156100f957600080fd5b506001600160a01b03813581169160200135166103b0565b34801561011d57600080fd5b506100a361047d565b6100d46004803603606081101561013c57600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460018302840111600160201b831117156101a257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061048c945050505050565b3480156101ef57600080fd5b506100d46004803603604081101561020657600080fd5b506001600160a01b03813581169160200135166105c5565b34801561022a57600080fd5b506100d46004803603602081101561024157600080fd5b50356001600160a01b0316610676565b34801561025d57600080fd5b506100a36004803603602081101561027457600080fd5b50356001600160a01b0316610766565b6000806060836001600160a01b03166040518080635c60da1b60e01b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b606091505b5091509150816102f757600080fd5b80806020019051602081101561030c57600080fd5b5051949350505050565b61031e6107c5565b6001600160a01b031661032f61047d565b6001600160a01b031614610378576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610810833981519152908390a3600080546001600160a01b0319169055565b6103b86107c5565b6001600160a01b03166103c961047d565b6001600160a01b031614610412576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316638f283970826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b505af1158015610475573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b6104946107c5565b6001600160a01b03166104a561047d565b6001600160a01b0316146104ee576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b826001600160a01b0316634f1ef2863484846040518463ffffffff1660e01b815260040180836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561055b578181015183820152602001610543565b50505050905090810190601f1680156105885780820380516001836020036101000a031916815260200191505b5093505050506000604051808303818588803b1580156105a757600080fd5b505af11580156105bb573d6000803e3d6000fd5b5050505050505050565b6105cd6107c5565b6001600160a01b03166105de61047d565b6001600160a01b031614610627576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316633659cfe6826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b61067e6107c5565b6001600160a01b031661068f61047d565b6001600160a01b0316146106d8576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b6001600160a01b03811661071d5760405162461bcd60e51b81526004018080602001828103825260268152602001806107ca6026913960400191505060405180910390fd5b600080546040516001600160a01b038085169392169160008051602061081083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806060836001600160a01b031660405180806303e1469160e61b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122036b5070566a1002642d9947563bea5cf538675280563c6ec32ab459f386727a264736f6c634300060c0033608060405260405162000c7438038062000c74833981810160405260608110156200002957600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200005557600080fd5b9083019060208201858111156200006b57600080fd5b82516401000000008111828201881017156200008657600080fd5b82525081516020918201929091019080838360005b83811015620000b55781810151838201526020016200009b565b50505050905090810190601f168015620000e35780820380516001836020036101000a031916815260200191505b5060405250849150829050620000f98262000137565b8051156200011a57620001188282620001ae60201b620003821760201c565b505b50620001239050565b6200012e82620001dd565b505050620003bf565b6200014d816200020160201b620003ae1760201c565b6200018a5760405162461bcd60e51b815260040180806020018281038252603681526020018062000c186036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b6060620001d6838360405180606001604052806027815260200162000bf16027913962000207565b9392505050565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b3b151590565b6060620002148462000201565b620002515760405162461bcd60e51b815260040180806020018281038252602681526020018062000c4e6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b60208310620002915780518252601f19909201916020918201910162000270565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114620002f3576040519150601f19603f3d011682016040523d82523d6000602084013e620002f8565b606091505b5090925090506200030b82828662000315565b9695505050505050565b6060831562000326575081620001d6565b825115620003375782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200038357818101518382015260200162000369565b50505050905090810190601f168015620003b15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b61082280620003cf6000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100985780635c60da1b146101165780638f28397014610147578063f851a4401461017a5761005d565b3661005d5761005b61018f565b005b61005b61018f565b34801561007157600080fd5b5061005b6004803603602081101561008857600080fd5b50356001600160a01b03166101a9565b61005b600480360360408110156100ae57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156100d857600080fd5b8201836020820111156100ea57600080fd5b803590602001918460018302840111600160201b8311171561010b57600080fd5b5090925090506101e3565b34801561012257600080fd5b5061012b610260565b604080516001600160a01b039092168252519081900360200190f35b34801561015357600080fd5b5061005b6004803603602081101561016a57600080fd5b50356001600160a01b031661029d565b34801561018657600080fd5b5061012b610357565b6101976103b4565b6101a76101a2610414565b610427565b565b6101b161044b565b6001600160a01b0316336001600160a01b031614156101d8576101d38161045e565b6101e0565b6101e061018f565b50565b6101eb61044b565b6001600160a01b0316336001600160a01b031614156102535761020d8361045e565b61024d8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038292505050565b5061025b565b61025b61018f565b505050565b600061026a61044b565b6001600160a01b0316336001600160a01b031614156102925761028b610414565b905061029a565b61029a61018f565b90565b6102a561044b565b6001600160a01b0316336001600160a01b031614156101d8576001600160a01b0381166103035760405162461bcd60e51b815260040180806020018281038252603a8152602001806106ce603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61032c61044b565b604080516001600160a01b03928316815291841660208301528051918290030190a16101d38161049e565b600061036161044b565b6001600160a01b0316336001600160a01b031614156102925761028b61044b565b60606103a78383604051806060016040528060278152602001610728602791396104b0565b9392505050565b3b151590565b6103bc61044b565b6001600160a01b0316336001600160a01b0316141561040c5760405162461bcd60e51b81526004018080602001828103825260428152602001806107ab6042913960600191505060405180910390fd5b6101a76101a7565b6000805160206107088339815191525490565b3660008037600080366000845af43d6000803e808015610446573d6000f35b3d6000fd5b6000805160206106ae8339815191525490565b610467816105b3565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6000805160206106ae83398151915255565b60606104bb846103ae565b6104f65760405162461bcd60e51b81526004018080602001828103825260268152602001806107856026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106105345780518252601f199092019160209182019101610515565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610594576040519150601f19603f3d011682016040523d82523d6000602084013e610599565b606091505b50915091506105a9828286610609565b9695505050505050565b6105bc816103ae565b6105f75760405162461bcd60e51b815260040180806020018281038252603681526020018061074f6036913960400191505060405180910390fd5b60008051602061070883398151915255565b606083156106185750816103a7565b8251156106285782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561067257818101518382015260200161065a565b50505050905090810190601f16801561069f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfeb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f2061646472657373360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a264697066735822122081df10ca698e92e05c31196844f9e9c309ede61eb7e45463098c415a459a179564736f6c634300060c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163744f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a26469706673582212204d75b380f18b73dd568c15a6b119c20a0ca06f14d44e5dae4a7e3f3f0619db4264736f6c634300060c0033608060405234801561001057600080fd5b506065805460ff19166001179055610dec8061002d6000396000f3fe6080604052600436106100765760003560e01c80636f791d291461007b578063715018a6146100a657806372f45866146100bd5780638129fc1c146100d057806381aac2d9146100e55780638da5cb5b14610105578063944f449514610127578063ce1d571f14610147578063f2fde38b1461015a575b600080fd5b34801561008757600080fd5b5061009061017a565b60405161009d9190610c23565b60405180910390f35b3480156100b257600080fd5b506100bb610183565b005b6100bb6100cb366004610a56565b610203565b3480156100dc57600080fd5b506100bb61039a565b3480156100f157600080fd5b506100bb610100366004610aeb565b610425565b34801561011157600080fd5b5061011a61053e565b60405161009d9190610c0f565b34801561013357600080fd5b506100bb610142366004610bad565b61054d565b6100bb610155366004610b2a565b610672565b34801561016657600080fd5b506100bb610175366004610a33565b61075a565b60655460ff1690565b61018b610809565b6001600160a01b031661019c61053e565b6001600160a01b0316146101cb5760405162461bcd60e51b81526004016101c290610ceb565b60405180910390fd5b6033546040516000916001600160a01b031690600080516020610d97833981519152908390a3603380546001600160a01b0319169055565b61020b610809565b6001600160a01b031661021c61053e565b6001600160a01b0316146102425760405162461bcd60e51b81526004016101c290610ceb565b8460005b8181101561039057600088888381811061025c57fe5b905060200281019061026e9190610d3d565b905011156102c4576102a886868381811061028557fe5b905060200201602081019061029a9190610a33565b6001600160a01b031661080d565b6102c45760405162461bcd60e51b81526004016101c290610c2e565b60008686838181106102d257fe5b90506020020160208101906102e79190610a33565b6001600160a01b03168585848181106102fc57fe5b905060200201358a8a8581811061030f57fe5b90506020028101906103219190610d3d565b60405161032f929190610bff565b60006040518083038185875af1925050503d806000811461036c576040519150601f19603f3d011682016040523d82523d6000602084013e610371565b606091505b5050905080610387576040513d806000833e8082fd5b50600101610246565b5050505050505050565b600054610100900460ff16806103b357506103b3610813565b806103c1575060005460ff16155b6103dd5760405162461bcd60e51b81526004016101c290610c9d565b600054610100900460ff16158015610408576000805460ff1961ff0019909116610100171660011790555b610410610824565b8015610422576000805461ff00191690555b50565b61042d610809565b6001600160a01b031661043e61053e565b6001600160a01b0316146104645760405162461bcd60e51b81526004016101c290610ceb565b8060005b818110156105385783838281811061047c57fe5b90506020020160208101906104919190610a33565b6001600160a01b03166370dea79a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156104cb57600080fd5b505af19250505080156104dc575060015b610530573d80801561050a576040519150601f19603f3d011682016040523d82523d6000602084013e61050f565b606091505b50805161052e5760405162461bcd60e51b81526004016101c290610d20565b505b600101610468565b50505050565b6033546001600160a01b031690565b610555610809565b6001600160a01b031661056661053e565b6001600160a01b03161461058c5760405162461bcd60e51b81526004016101c290610ceb565b8060005b8181101561066b57846001600160a01b0316637427be518585848181106105b357fe5b90506020020160208101906105c89190610a33565b6040518263ffffffff1660e01b81526004016105e49190610c0f565b600060405180830381600087803b1580156105fe57600080fd5b505af192505050801561060f575060015b610663573d80801561063d576040519150601f19603f3d011682016040523d82523d6000602084013e610642565b606091505b5080516106615760405162461bcd60e51b81526004016101c290610d20565b505b600101610590565b5050505050565b61067a610809565b6001600160a01b031661068b61053e565b6001600160a01b0316146106b15760405162461bcd60e51b81526004016101c290610ceb565b82156106e5576106c9826001600160a01b031661080d565b6106e55760405162461bcd60e51b81526004016101c290610c2e565b6000826001600160a01b0316828686604051610702929190610bff565b60006040518083038185875af1925050503d806000811461073f576040519150601f19603f3d011682016040523d82523d6000602084013e610744565b606091505b505090508061066b576040513d806000833e8082fd5b610762610809565b6001600160a01b031661077361053e565b6001600160a01b0316146107995760405162461bcd60e51b81526004016101c290610ceb565b6001600160a01b0381166107bf5760405162461bcd60e51b81526004016101c290610c57565b6033546040516001600160a01b03808416921690600080516020610d9783398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b3b151590565b600061081e3061080d565b15905090565b600054610100900460ff168061083d575061083d610813565b8061084b575060005460ff16155b6108675760405162461bcd60e51b81526004016101c290610c9d565b600054610100900460ff16158015610892576000805460ff1961ff0019909116610100171660011790555b61089a6108a2565b610410610923565b600054610100900460ff16806108bb57506108bb610813565b806108c9575060005460ff16155b6108e55760405162461bcd60e51b81526004016101c290610c9d565b600054610100900460ff16158015610410576000805460ff1961ff0019909116610100171660011790558015610422576000805461ff001916905550565b600054610100900460ff168061093c575061093c610813565b8061094a575060005460ff16155b6109665760405162461bcd60e51b81526004016101c290610c9d565b600054610100900460ff16158015610991576000805460ff1961ff0019909116610100171660011790555b600061099b610809565b603380546001600160a01b0319166001600160a01b03831690811790915560405191925090600090600080516020610d97833981519152908290a3508015610422576000805461ff001916905550565b60008083601f8401126109fc578182fd5b5081356001600160401b03811115610a12578182fd5b6020830191508360208083028501011115610a2c57600080fd5b9250929050565b600060208284031215610a44578081fd5b8135610a4f81610d81565b9392505050565b60008060008060008060608789031215610a6e578182fd5b86356001600160401b0380821115610a84578384fd5b610a908a838b016109eb565b90985096506020890135915080821115610aa8578384fd5b610ab48a838b016109eb565b90965094506040890135915080821115610acc578384fd5b50610ad989828a016109eb565b979a9699509497509295939492505050565b60008060208385031215610afd578182fd5b82356001600160401b03811115610b12578283fd5b610b1e858286016109eb565b90969095509350505050565b60008060008060608587031215610b3f578384fd5b84356001600160401b0380821115610b55578586fd5b818701915087601f830112610b68578586fd5b813581811115610b76578687fd5b886020828501011115610b87578687fd5b60209283019650945050850135610b9d81610d81565b9396929550929360400135925050565b600080600060408486031215610bc1578283fd5b8335610bcc81610d81565b925060208401356001600160401b03811115610be6578283fd5b610bf2868287016109eb565b9497909650939450505050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b901515815260200190565b6020808252600f908201526e2727afa1a7a222afa0aa2fa0a2222960891b604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526003908201526247415360e81b604082015260600190565b6000808335601e19843603018112610d53578283fd5b8301803591506001600160401b03821115610d6c578283fd5b602001915036819003821315610a2c57600080fd5b6001600160a01b038116811461042257600080fdfe8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a2646970667358221220c32f3e207e1ba826a701c5a23fab40d10471e63cefa1871f63062a427f95a7e364736f6c634300060c0033"

// DeployValidatorWalletCreator deploys a new Ethereum contract, binding an instance of ValidatorWalletCreator to it.
func DeployValidatorWalletCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorWalletCreator, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorWalletCreatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorWalletCreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorWalletCreator{ValidatorWalletCreatorCaller: ValidatorWalletCreatorCaller{contract: contract}, ValidatorWalletCreatorTransactor: ValidatorWalletCreatorTransactor{contract: contract}, ValidatorWalletCreatorFilterer: ValidatorWalletCreatorFilterer{contract: contract}}, nil
}

// ValidatorWalletCreator is an auto generated Go binding around an Ethereum contract.
type ValidatorWalletCreator struct {
	ValidatorWalletCreatorCaller     // Read-only binding to the contract
	ValidatorWalletCreatorTransactor // Write-only binding to the contract
	ValidatorWalletCreatorFilterer   // Log filterer for contract events
}

// ValidatorWalletCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorWalletCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorWalletCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorWalletCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorWalletCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorWalletCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorWalletCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorWalletCreatorSession struct {
	Contract     *ValidatorWalletCreator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorWalletCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorWalletCreatorCallerSession struct {
	Contract *ValidatorWalletCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ValidatorWalletCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorWalletCreatorTransactorSession struct {
	Contract     *ValidatorWalletCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ValidatorWalletCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorWalletCreatorRaw struct {
	Contract *ValidatorWalletCreator // Generic contract binding to access the raw methods on
}

// ValidatorWalletCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorWalletCreatorCallerRaw struct {
	Contract *ValidatorWalletCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorWalletCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorWalletCreatorTransactorRaw struct {
	Contract *ValidatorWalletCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorWalletCreator creates a new instance of ValidatorWalletCreator, bound to a specific deployed contract.
func NewValidatorWalletCreator(address common.Address, backend bind.ContractBackend) (*ValidatorWalletCreator, error) {
	contract, err := bindValidatorWalletCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorWalletCreator{ValidatorWalletCreatorCaller: ValidatorWalletCreatorCaller{contract: contract}, ValidatorWalletCreatorTransactor: ValidatorWalletCreatorTransactor{contract: contract}, ValidatorWalletCreatorFilterer: ValidatorWalletCreatorFilterer{contract: contract}}, nil
}

// NewValidatorWalletCreatorCaller creates a new read-only instance of ValidatorWalletCreator, bound to a specific deployed contract.
func NewValidatorWalletCreatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorWalletCreatorCaller, error) {
	contract, err := bindValidatorWalletCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorWalletCreatorCaller{contract: contract}, nil
}

// NewValidatorWalletCreatorTransactor creates a new write-only instance of ValidatorWalletCreator, bound to a specific deployed contract.
func NewValidatorWalletCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorWalletCreatorTransactor, error) {
	contract, err := bindValidatorWalletCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorWalletCreatorTransactor{contract: contract}, nil
}

// NewValidatorWalletCreatorFilterer creates a new log filterer instance of ValidatorWalletCreator, bound to a specific deployed contract.
func NewValidatorWalletCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorWalletCreatorFilterer, error) {
	contract, err := bindValidatorWalletCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorWalletCreatorFilterer{contract: contract}, nil
}

// bindValidatorWalletCreator binds a generic wrapper to an already deployed contract.
func bindValidatorWalletCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorWalletCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorWalletCreator *ValidatorWalletCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorWalletCreator.Contract.ValidatorWalletCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorWalletCreator *ValidatorWalletCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.ValidatorWalletCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorWalletCreator *ValidatorWalletCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.ValidatorWalletCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorWalletCreator *ValidatorWalletCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorWalletCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorWalletCreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorSession) Owner() (common.Address, error) {
	return _ValidatorWalletCreator.Contract.Owner(&_ValidatorWalletCreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorCallerSession) Owner() (common.Address, error) {
	return _ValidatorWalletCreator.Contract.Owner(&_ValidatorWalletCreator.CallOpts)
}

// Template is a free data retrieval call binding the contract method 0x6f2ddd93.
//
// Solidity: function template() view returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorCaller) Template(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorWalletCreator.contract.Call(opts, &out, "template")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Template is a free data retrieval call binding the contract method 0x6f2ddd93.
//
// Solidity: function template() view returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorSession) Template() (common.Address, error) {
	return _ValidatorWalletCreator.Contract.Template(&_ValidatorWalletCreator.CallOpts)
}

// Template is a free data retrieval call binding the contract method 0x6f2ddd93.
//
// Solidity: function template() view returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorCallerSession) Template() (common.Address, error) {
	return _ValidatorWalletCreator.Contract.Template(&_ValidatorWalletCreator.CallOpts)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x11ebbf24.
//
// Solidity: function createWallet() returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactor) CreateWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWalletCreator.contract.Transact(opts, "createWallet")
}

// CreateWallet is a paid mutator transaction binding the contract method 0x11ebbf24.
//
// Solidity: function createWallet() returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorSession) CreateWallet() (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.CreateWallet(&_ValidatorWalletCreator.TransactOpts)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x11ebbf24.
//
// Solidity: function createWallet() returns(address)
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactorSession) CreateWallet() (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.CreateWallet(&_ValidatorWalletCreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWalletCreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.RenounceOwnership(&_ValidatorWalletCreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.RenounceOwnership(&_ValidatorWalletCreator.TransactOpts)
}

// SetTemplate is a paid mutator transaction binding the contract method 0x89c716d1.
//
// Solidity: function setTemplate(address _template) returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactor) SetTemplate(opts *bind.TransactOpts, _template common.Address) (*types.Transaction, error) {
	return _ValidatorWalletCreator.contract.Transact(opts, "setTemplate", _template)
}

// SetTemplate is a paid mutator transaction binding the contract method 0x89c716d1.
//
// Solidity: function setTemplate(address _template) returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorSession) SetTemplate(_template common.Address) (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.SetTemplate(&_ValidatorWalletCreator.TransactOpts, _template)
}

// SetTemplate is a paid mutator transaction binding the contract method 0x89c716d1.
//
// Solidity: function setTemplate(address _template) returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactorSession) SetTemplate(_template common.Address) (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.SetTemplate(&_ValidatorWalletCreator.TransactOpts, _template)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorWalletCreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.TransferOwnership(&_ValidatorWalletCreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorWalletCreator *ValidatorWalletCreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorWalletCreator.Contract.TransferOwnership(&_ValidatorWalletCreator.TransactOpts, newOwner)
}

// ValidatorWalletCreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValidatorWalletCreator contract.
type ValidatorWalletCreatorOwnershipTransferredIterator struct {
	Event *ValidatorWalletCreatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorWalletCreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWalletCreatorOwnershipTransferred)
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
		it.Event = new(ValidatorWalletCreatorOwnershipTransferred)
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
func (it *ValidatorWalletCreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWalletCreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWalletCreatorOwnershipTransferred represents a OwnershipTransferred event raised by the ValidatorWalletCreator contract.
type ValidatorWalletCreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorWalletCreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorWalletCreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorWalletCreatorOwnershipTransferredIterator{contract: _ValidatorWalletCreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorWalletCreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorWalletCreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWalletCreatorOwnershipTransferred)
				if err := _ValidatorWalletCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorWalletCreatorOwnershipTransferred, error) {
	event := new(ValidatorWalletCreatorOwnershipTransferred)
	if err := _ValidatorWalletCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWalletCreatorTemplateUpdatedIterator is returned from FilterTemplateUpdated and is used to iterate over the raw logs and unpacked data for TemplateUpdated events raised by the ValidatorWalletCreator contract.
type ValidatorWalletCreatorTemplateUpdatedIterator struct {
	Event *ValidatorWalletCreatorTemplateUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorWalletCreatorTemplateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWalletCreatorTemplateUpdated)
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
		it.Event = new(ValidatorWalletCreatorTemplateUpdated)
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
func (it *ValidatorWalletCreatorTemplateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWalletCreatorTemplateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWalletCreatorTemplateUpdated represents a TemplateUpdated event raised by the ValidatorWalletCreator contract.
type ValidatorWalletCreatorTemplateUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTemplateUpdated is a free log retrieval operation binding the contract event 0x6eb26f176dd9180849dd4874d3530de0e5c1f62a6e6798d34e3abfc11f1db2cc.
//
// Solidity: event TemplateUpdated()
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) FilterTemplateUpdated(opts *bind.FilterOpts) (*ValidatorWalletCreatorTemplateUpdatedIterator, error) {

	logs, sub, err := _ValidatorWalletCreator.contract.FilterLogs(opts, "TemplateUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorWalletCreatorTemplateUpdatedIterator{contract: _ValidatorWalletCreator.contract, event: "TemplateUpdated", logs: logs, sub: sub}, nil
}

// WatchTemplateUpdated is a free log subscription operation binding the contract event 0x6eb26f176dd9180849dd4874d3530de0e5c1f62a6e6798d34e3abfc11f1db2cc.
//
// Solidity: event TemplateUpdated()
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) WatchTemplateUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorWalletCreatorTemplateUpdated) (event.Subscription, error) {

	logs, sub, err := _ValidatorWalletCreator.contract.WatchLogs(opts, "TemplateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWalletCreatorTemplateUpdated)
				if err := _ValidatorWalletCreator.contract.UnpackLog(event, "TemplateUpdated", log); err != nil {
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

// ParseTemplateUpdated is a log parse operation binding the contract event 0x6eb26f176dd9180849dd4874d3530de0e5c1f62a6e6798d34e3abfc11f1db2cc.
//
// Solidity: event TemplateUpdated()
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) ParseTemplateUpdated(log types.Log) (*ValidatorWalletCreatorTemplateUpdated, error) {
	event := new(ValidatorWalletCreatorTemplateUpdated)
	if err := _ValidatorWalletCreator.contract.UnpackLog(event, "TemplateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWalletCreatorWalletCreatedIterator is returned from FilterWalletCreated and is used to iterate over the raw logs and unpacked data for WalletCreated events raised by the ValidatorWalletCreator contract.
type ValidatorWalletCreatorWalletCreatedIterator struct {
	Event *ValidatorWalletCreatorWalletCreated // Event containing the contract specifics and raw log

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
func (it *ValidatorWalletCreatorWalletCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWalletCreatorWalletCreated)
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
		it.Event = new(ValidatorWalletCreatorWalletCreated)
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
func (it *ValidatorWalletCreatorWalletCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWalletCreatorWalletCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWalletCreatorWalletCreated represents a WalletCreated event raised by the ValidatorWalletCreator contract.
type ValidatorWalletCreatorWalletCreated struct {
	WalletAddress common.Address
	UserAddress   common.Address
	AdminProxy    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWalletCreated is a free log retrieval operation binding the contract event 0xca0b7dde26052d34217ef1a0cee48085a07ca32da0a918609937a307d496bbf5.
//
// Solidity: event WalletCreated(address indexed walletAddress, address indexed userAddress, address adminProxy)
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) FilterWalletCreated(opts *bind.FilterOpts, walletAddress []common.Address, userAddress []common.Address) (*ValidatorWalletCreatorWalletCreatedIterator, error) {

	var walletAddressRule []interface{}
	for _, walletAddressItem := range walletAddress {
		walletAddressRule = append(walletAddressRule, walletAddressItem)
	}
	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _ValidatorWalletCreator.contract.FilterLogs(opts, "WalletCreated", walletAddressRule, userAddressRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorWalletCreatorWalletCreatedIterator{contract: _ValidatorWalletCreator.contract, event: "WalletCreated", logs: logs, sub: sub}, nil
}

// WatchWalletCreated is a free log subscription operation binding the contract event 0xca0b7dde26052d34217ef1a0cee48085a07ca32da0a918609937a307d496bbf5.
//
// Solidity: event WalletCreated(address indexed walletAddress, address indexed userAddress, address adminProxy)
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) WatchWalletCreated(opts *bind.WatchOpts, sink chan<- *ValidatorWalletCreatorWalletCreated, walletAddress []common.Address, userAddress []common.Address) (event.Subscription, error) {

	var walletAddressRule []interface{}
	for _, walletAddressItem := range walletAddress {
		walletAddressRule = append(walletAddressRule, walletAddressItem)
	}
	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _ValidatorWalletCreator.contract.WatchLogs(opts, "WalletCreated", walletAddressRule, userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWalletCreatorWalletCreated)
				if err := _ValidatorWalletCreator.contract.UnpackLog(event, "WalletCreated", log); err != nil {
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

// ParseWalletCreated is a log parse operation binding the contract event 0xca0b7dde26052d34217ef1a0cee48085a07ca32da0a918609937a307d496bbf5.
//
// Solidity: event WalletCreated(address indexed walletAddress, address indexed userAddress, address adminProxy)
func (_ValidatorWalletCreator *ValidatorWalletCreatorFilterer) ParseWalletCreated(log types.Log) (*ValidatorWalletCreatorWalletCreated, error) {
	event := new(ValidatorWalletCreatorWalletCreated)
	if err := _ValidatorWalletCreator.contract.UnpackLog(event, "WalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
