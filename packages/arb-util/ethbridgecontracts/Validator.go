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
const ValidatorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"executeTransactions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollupUser\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"returnOldDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallenge[]\",\"name\":\"challenges\",\"type\":\"address[]\"}],\"name\":\"timeoutChallenges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorBin = "0x608060405234801561001057600080fd5b506065805460ff19166001179055610df58061002d6000396000f3fe6080604052600436106100765760003560e01c80636f791d291461007b578063715018a6146100a657806372f45866146100bd5780638129fc1c146100d057806381aac2d9146100e55780638da5cb5b14610105578063944f449514610127578063ce1d571f14610147578063f2fde38b1461015a575b600080fd5b34801561008757600080fd5b5061009061017a565b60405161009d9190610c29565b60405180910390f35b3480156100b257600080fd5b506100bb610183565b005b6100bb6100cb366004610a57565b610203565b3480156100dc57600080fd5b506100bb61039a565b3480156100f157600080fd5b506100bb610100366004610aed565b610425565b34801561011157600080fd5b5061011a61053e565b60405161009d9190610c15565b34801561013357600080fd5b506100bb610142366004610bb2565b61054d565b6100bb610155366004610b2d565b610672565b34801561016657600080fd5b506100bb610175366004610a34565b61075a565b60655460ff1690565b61018b610809565b6001600160a01b031661019c61053e565b6001600160a01b0316146101cb5760405162461bcd60e51b81526004016101c290610cf1565b60405180910390fd5b6033546040516000916001600160a01b031690600080516020610da0833981519152908390a3603380546001600160a01b0319169055565b61020b610809565b6001600160a01b031661021c61053e565b6001600160a01b0316146102425760405162461bcd60e51b81526004016101c290610cf1565b8460005b8181101561039057600088888381811061025c57fe5b905060200281019061026e9190610d43565b905011156102c4576102a886868381811061028557fe5b905060200201602081019061029a9190610a34565b6001600160a01b031661080d565b6102c45760405162461bcd60e51b81526004016101c290610c34565b60008686838181106102d257fe5b90506020020160208101906102e79190610a34565b6001600160a01b03168585848181106102fc57fe5b905060200201358a8a8581811061030f57fe5b90506020028101906103219190610d43565b60405161032f929190610c05565b60006040518083038185875af1925050503d806000811461036c576040519150601f19603f3d011682016040523d82523d6000602084013e610371565b606091505b5050905080610387576040513d806000833e8082fd5b50600101610246565b5050505050505050565b600054610100900460ff16806103b357506103b3610813565b806103c1575060005460ff16155b6103dd5760405162461bcd60e51b81526004016101c290610ca3565b600054610100900460ff16158015610408576000805460ff1961ff0019909116610100171660011790555b610410610824565b8015610422576000805461ff00191690555b50565b61042d610809565b6001600160a01b031661043e61053e565b6001600160a01b0316146104645760405162461bcd60e51b81526004016101c290610cf1565b8060005b818110156105385783838281811061047c57fe5b90506020020160208101906104919190610a34565b6001600160a01b03166370dea79a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156104cb57600080fd5b505af19250505080156104dc575060015b610530573d80801561050a576040519150601f19603f3d011682016040523d82523d6000602084013e61050f565b606091505b50805161052e5760405162461bcd60e51b81526004016101c290610d26565b505b600101610468565b50505050565b6033546001600160a01b031690565b610555610809565b6001600160a01b031661056661053e565b6001600160a01b03161461058c5760405162461bcd60e51b81526004016101c290610cf1565b8060005b8181101561066b57846001600160a01b0316637427be518585848181106105b357fe5b90506020020160208101906105c89190610a34565b6040518263ffffffff1660e01b81526004016105e49190610c15565b600060405180830381600087803b1580156105fe57600080fd5b505af192505050801561060f575060015b610663573d80801561063d576040519150601f19603f3d011682016040523d82523d6000602084013e610642565b606091505b5080516106615760405162461bcd60e51b81526004016101c290610d26565b505b600101610590565b5050505050565b61067a610809565b6001600160a01b031661068b61053e565b6001600160a01b0316146106b15760405162461bcd60e51b81526004016101c290610cf1565b82156106e5576106c9826001600160a01b031661080d565b6106e55760405162461bcd60e51b81526004016101c290610c34565b6000826001600160a01b0316828686604051610702929190610c05565b60006040518083038185875af1925050503d806000811461073f576040519150601f19603f3d011682016040523d82523d6000602084013e610744565b606091505b505090508061066b576040513d806000833e8082fd5b610762610809565b6001600160a01b031661077361053e565b6001600160a01b0316146107995760405162461bcd60e51b81526004016101c290610cf1565b6001600160a01b0381166107bf5760405162461bcd60e51b81526004016101c290610c5d565b6033546040516001600160a01b03808416921690600080516020610da083398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b3b151590565b600061081e3061080d565b15905090565b600054610100900460ff168061083d575061083d610813565b8061084b575060005460ff16155b6108675760405162461bcd60e51b81526004016101c290610ca3565b600054610100900460ff16158015610892576000805460ff1961ff0019909116610100171660011790555b61089a6108a2565b610410610923565b600054610100900460ff16806108bb57506108bb610813565b806108c9575060005460ff16155b6108e55760405162461bcd60e51b81526004016101c290610ca3565b600054610100900460ff16158015610410576000805460ff1961ff0019909116610100171660011790558015610422576000805461ff001916905550565b600054610100900460ff168061093c575061093c610813565b8061094a575060005460ff16155b6109665760405162461bcd60e51b81526004016101c290610ca3565b600054610100900460ff16158015610991576000805460ff1961ff0019909116610100171660011790555b600061099b610809565b603380546001600160a01b0319166001600160a01b03831690811790915560405191925090600090600080516020610da0833981519152908290a3508015610422576000805461ff001916905550565b60008083601f8401126109fc578182fd5b50813567ffffffffffffffff811115610a13578182fd5b6020830191508360208083028501011115610a2d57600080fd5b9250929050565b600060208284031215610a45578081fd5b8135610a5081610d8a565b9392505050565b60008060008060008060608789031215610a6f578182fd5b863567ffffffffffffffff80821115610a86578384fd5b610a928a838b016109eb565b90985096506020890135915080821115610aaa578384fd5b610ab68a838b016109eb565b90965094506040890135915080821115610ace578384fd5b50610adb89828a016109eb565b979a9699509497509295939492505050565b60008060208385031215610aff578182fd5b823567ffffffffffffffff811115610b15578283fd5b610b21858286016109eb565b90969095509350505050565b60008060008060608587031215610b42578384fd5b843567ffffffffffffffff80821115610b59578586fd5b81870188601f820112610b6a578687fd5b8035925081831115610b7a578687fd5b886020848301011115610b8b578687fd5b6020908101965091945050850135610ba281610d8a565b9396929550929360400135925050565b600080600060408486031215610bc6578283fd5b8335610bd181610d8a565b9250602084013567ffffffffffffffff811115610bec578283fd5b610bf8868287016109eb565b9497909650939450505050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b901515815260200190565b6020808252600f908201526e2727afa1a7a222afa0aa2fa0a2222960891b604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526003908201526247415360e81b604082015260600190565b6000808335601e19843603018112610d59578283fd5b8084018035925067ffffffffffffffff831115610d74578384fd5b60200192505036819003821315610a2d57600080fd5b6001600160a01b038116811461042257600080fdfe8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a26469706673582212200bb21cb23aca9a11a091fc0a357ce63ac1fbd675a6598b4baed85449804874a864736f6c634300060b0033"

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
