// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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

// ValidatorMetaData contains all meta data concerning the Validator contract.
var ValidatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"refunder\",\"type\":\"address\"}],\"name\":\"executeTransactionWithGasRefunder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"executeTransactions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"refunder\",\"type\":\"address\"}],\"name\":\"executeTransactionsWithGasRefunder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollupUser\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"returnOldDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollupUser\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"refunder\",\"type\":\"address\"}],\"name\":\"returnOldDepositsWithGasRefunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallenge[]\",\"name\":\"challenges\",\"type\":\"address[]\"}],\"name\":\"timeoutChallenges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallenge[]\",\"name\":\"challenges\",\"type\":\"address[]\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"refunder\",\"type\":\"address\"}],\"name\":\"timeoutChallengesWithGasRefunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506065805460ff191660011790556114888061002d6000396000f3fe6080604052600436106100b25760003560e01c806381aac2d91161006f57806381aac2d9146101545780638da5cb5b14610174578063944f44951461019657806397da185a146101b6578063c083bc5b146101c9578063ce1d571f146101e9578063f2fde38b146101fc576100b2565b806347336bf0146100b757806348815c09146100d95780636f791d29146100ec578063715018a61461011757806372f458661461012c5780638129fc1c1461013f575b600080fd5b3480156100c357600080fd5b506100d76100d2366004611082565b61021c565b005b6100d76100e7366004611150565b6103de565b3480156100f857600080fd5b5061010161056c565b60405161010e91906112bd565b60405180910390f35b34801561012357600080fd5b506100d7610575565b6100d761013a366004610f05565b6105ec565b34801561014b57600080fd5b506100d7610643565b34801561016057600080fd5b506100d761016f366004611043565b6106ce565b34801561018057600080fd5b5061018961071d565b60405161010e9190611288565b3480156101a257600080fd5b506100d76101b13660046111c1565b61072c565b6100d76101c4366004610f9a565b61077d565b3480156101d557600080fd5b506100d76101e4366004611213565b6109b3565b6100d76101f73660046110f6565b610b77565b34801561020857600080fd5b506100d7610217366004610ee2565b610bca565b610224610c79565b6001600160a01b031661023561071d565b6001600160a01b0316146102645760405162461bcd60e51b815260040161025b90611385565b60405180910390fd5b8060005a90508360005b8181101561033e5786868281811061028257fe5b90506020020160208101906102979190610ee2565b6001600160a01b03166370dea79a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156102d157600080fd5b505af19250505080156102e2575060015b610336573d808015610310576040519150601f19603f3d011682016040523d82523d6000602084013e610315565b606091505b5080516103345760405162461bcd60e51b815260040161025b906113ba565b505b60010161026e565b50506001600160a01b038216156103d757366001600160a01b03831663e3db8a49335a8503846040518463ffffffff1660e01b81526004016103829392919061129c565b602060405180830381600087803b15801561039c57600080fd5b505af11580156103b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d491906110d6565b50505b5050505050565b6103e6610c79565b6001600160a01b03166103f761071d565b6001600160a01b03161461041d5760405162461bcd60e51b815260040161025b90611385565b8060005a905085156104575761043b856001600160a01b0316610c7d565b6104575760405162461bcd60e51b815260040161025b906112c8565b6000856001600160a01b0316858989604051610474929190611278565b60006040518083038185875af1925050503d80600081146104b1576040519150601f19603f3d011682016040523d82523d6000602084013e6104b6565b606091505b50509050806104cc576040513d806000833e8082fd5b506001600160a01b038216156103d457366001600160a01b03831663e3db8a49335a8503846040518463ffffffff1660e01b815260040161050f9392919061129c565b602060405180830381600087803b15801561052957600080fd5b505af115801561053d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056191906110d6565b505050505050505050565b60655460ff1690565b61057d610c79565b6001600160a01b031661058e61071d565b6001600160a01b0316146105b45760405162461bcd60e51b815260040161025b90611385565b6033546040516000916001600160a01b031690600080516020611433833981519152908390a3603380546001600160a01b0319169055565b6105f4610c79565b6001600160a01b031661060561071d565b6001600160a01b03161461062b5760405162461bcd60e51b815260040161025b90611385565b61063b868686868686600061077d565b505050505050565b600054610100900460ff168061065c575061065c610c83565b8061066a575060005460ff16155b6106865760405162461bcd60e51b815260040161025b90611337565b600054610100900460ff161580156106b1576000805460ff1961ff0019909116610100171660011790555b6106b9610c94565b80156106cb576000805461ff00191690555b50565b6106d6610c79565b6001600160a01b03166106e761071d565b6001600160a01b03161461070d5760405162461bcd60e51b815260040161025b90611385565b6107198282600061021c565b5050565b6033546001600160a01b031690565b610734610c79565b6001600160a01b031661074561071d565b6001600160a01b03161461076b5760405162461bcd60e51b815260040161025b90611385565b61077883838360006109b3565b505050565b610785610c79565b6001600160a01b031661079661071d565b6001600160a01b0316146107bc5760405162461bcd60e51b815260040161025b90611385565b8060005a90508760005b818110156109105760008b8b838181106107dc57fe5b90506020028101906107ee91906113d7565b905011156108445761082889898381811061080557fe5b905060200201602081019061081a9190610ee2565b6001600160a01b0316610c7d565b6108445760405162461bcd60e51b815260040161025b906112c8565b600089898381811061085257fe5b90506020020160208101906108679190610ee2565b6001600160a01b031688888481811061087c57fe5b905060200201358d8d8581811061088f57fe5b90506020028101906108a191906113d7565b6040516108af929190611278565b60006040518083038185875af1925050503d80600081146108ec576040519150601f19603f3d011682016040523d82523d6000602084013e6108f1565b606091505b5050905080610907576040513d806000833e8082fd5b506001016107c6565b50506001600160a01b0382161561056157366001600160a01b03831663e3db8a49335a8503846040518463ffffffff1660e01b81526004016109549392919061129c565b602060405180830381600087803b15801561096e57600080fd5b505af1158015610982573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a691906110d6565b5050505050505050505050565b6109bb610c79565b6001600160a01b03166109cc61071d565b6001600160a01b0316146109f25760405162461bcd60e51b815260040161025b90611385565b8060005a90508360005b81811015610ad757876001600160a01b0316637427be51888884818110610a1f57fe5b9050602002016020810190610a349190610ee2565b6040518263ffffffff1660e01b8152600401610a509190611288565b600060405180830381600087803b158015610a6a57600080fd5b505af1925050508015610a7b575060015b610acf573d808015610aa9576040519150601f19603f3d011682016040523d82523d6000602084013e610aae565b606091505b508051610acd5760405162461bcd60e51b815260040161025b906113ba565b505b6001016109fc565b50506001600160a01b0382161561063b57366001600160a01b03831663e3db8a49335a8503846040518463ffffffff1660e01b8152600401610b1b9392919061129c565b602060405180830381600087803b158015610b3557600080fd5b505af1158015610b49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6d91906110d6565b5050505050505050565b610b7f610c79565b6001600160a01b0316610b9061071d565b6001600160a01b031614610bb65760405162461bcd60e51b815260040161025b90611385565b610bc48484848460006103de565b50505050565b610bd2610c79565b6001600160a01b0316610be361071d565b6001600160a01b031614610c095760405162461bcd60e51b815260040161025b90611385565b6001600160a01b038116610c2f5760405162461bcd60e51b815260040161025b906112f1565b6033546040516001600160a01b0380841692169060008051602061143383398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b3b151590565b6000610c8e30610c7d565b15905090565b600054610100900460ff1680610cad5750610cad610c83565b80610cbb575060005460ff16155b610cd75760405162461bcd60e51b815260040161025b90611337565b600054610100900460ff16158015610d02576000805460ff1961ff0019909116610100171660011790555b610d0a610d12565b6106b9610d93565b600054610100900460ff1680610d2b5750610d2b610c83565b80610d39575060005460ff16155b610d555760405162461bcd60e51b815260040161025b90611337565b600054610100900460ff161580156106b9576000805460ff1961ff00199091166101001716600117905580156106cb576000805461ff001916905550565b600054610100900460ff1680610dac5750610dac610c83565b80610dba575060005460ff16155b610dd65760405162461bcd60e51b815260040161025b90611337565b600054610100900460ff16158015610e01576000805460ff1961ff0019909116610100171660011790555b6000610e0b610c79565b603380546001600160a01b0319166001600160a01b03831690811790915560405191925090600090600080516020611433833981519152908290a35080156106cb576000805461ff001916905550565b60008083601f840112610e6c578182fd5b5081356001600160401b03811115610e82578182fd5b6020830191508360208083028501011115610e9c57600080fd5b9250929050565b60008083601f840112610eb4578182fd5b5081356001600160401b03811115610eca578182fd5b602083019150836020828501011115610e9c57600080fd5b600060208284031215610ef3578081fd5b8135610efe8161141d565b9392505050565b60008060008060008060608789031215610f1d578182fd5b86356001600160401b0380821115610f33578384fd5b610f3f8a838b01610e5b565b90985096506020890135915080821115610f57578384fd5b610f638a838b01610e5b565b90965094506040890135915080821115610f7b578384fd5b50610f8889828a01610e5b565b979a9699509497509295939492505050565b60008060008060008060006080888a031215610fb4578081fd5b87356001600160401b0380821115610fca578283fd5b610fd68b838c01610e5b565b909950975060208a0135915080821115610fee578283fd5b610ffa8b838c01610e5b565b909750955060408a0135915080821115611012578283fd5b5061101f8a828b01610e5b565b90945092505060608801356110338161141d565b8091505092959891949750929550565b60008060208385031215611055578182fd5b82356001600160401b0381111561106a578283fd5b61107685828601610e5b565b90969095509350505050565b600080600060408486031215611096578283fd5b83356001600160401b038111156110ab578384fd5b6110b786828701610e5b565b90945092505060208401356110cb8161141d565b809150509250925092565b6000602082840312156110e7578081fd5b81518015158114610efe578182fd5b6000806000806060858703121561110b578384fd5b84356001600160401b03811115611120578485fd5b61112c87828801610ea3565b90955093505060208501356111408161141d565b9396929550929360400135925050565b600080600080600060808688031215611167578081fd5b85356001600160401b0381111561117c578182fd5b61118888828901610ea3565b909650945050602086013561119c8161141d565b92506040860135915060608601356111b38161141d565b809150509295509295909350565b6000806000604084860312156111d5578081fd5b83356111e08161141d565b925060208401356001600160401b038111156111fa578182fd5b61120686828701610e5b565b9497909650939450505050565b60008060008060608587031215611228578182fd5b84356112338161141d565b935060208501356001600160401b0381111561124d578283fd5b61125987828801610e5b565b909450925050604085013561126d8161141d565b939692955090935050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b6001600160a01b039390931683526020830191909152604082015260600190565b901515815260200190565b6020808252600f908201526e2727afa1a7a222afa0aa2fa0a2222960891b604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526003908201526247415360e81b604082015260600190565b6000808335601e198436030181126113ed578283fd5b808401803592506001600160401b03831115611407578384fd5b60200192505036819003821315610e9c57600080fd5b6001600160a01b03811681146106cb57600080fdfe8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122006ec104bed52d31b2588e4bd351e85ebd4c4dd0906ddadd47e550881adea9d6664736f6c634300060b0033",
}

// ValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMetaData.ABI instead.
var ValidatorABI = ValidatorMetaData.ABI

// ValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorMetaData.Bin instead.
var ValidatorBin = ValidatorMetaData.Bin

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := ValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorBin), backend)
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

// ExecuteTransactionWithGasRefunder is a paid mutator transaction binding the contract method 0x48815c09.
//
// Solidity: function executeTransactionWithGasRefunder(bytes data, address destination, uint256 amount, address refunder) payable returns()
func (_Validator *ValidatorTransactor) ExecuteTransactionWithGasRefunder(opts *bind.TransactOpts, data []byte, destination common.Address, amount *big.Int, refunder common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "executeTransactionWithGasRefunder", data, destination, amount, refunder)
}

// ExecuteTransactionWithGasRefunder is a paid mutator transaction binding the contract method 0x48815c09.
//
// Solidity: function executeTransactionWithGasRefunder(bytes data, address destination, uint256 amount, address refunder) payable returns()
func (_Validator *ValidatorSession) ExecuteTransactionWithGasRefunder(data []byte, destination common.Address, amount *big.Int, refunder common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactionWithGasRefunder(&_Validator.TransactOpts, data, destination, amount, refunder)
}

// ExecuteTransactionWithGasRefunder is a paid mutator transaction binding the contract method 0x48815c09.
//
// Solidity: function executeTransactionWithGasRefunder(bytes data, address destination, uint256 amount, address refunder) payable returns()
func (_Validator *ValidatorTransactorSession) ExecuteTransactionWithGasRefunder(data []byte, destination common.Address, amount *big.Int, refunder common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactionWithGasRefunder(&_Validator.TransactOpts, data, destination, amount, refunder)
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

// ExecuteTransactionsWithGasRefunder is a paid mutator transaction binding the contract method 0x97da185a.
//
// Solidity: function executeTransactionsWithGasRefunder(bytes[] data, address[] destination, uint256[] amount, address refunder) payable returns()
func (_Validator *ValidatorTransactor) ExecuteTransactionsWithGasRefunder(opts *bind.TransactOpts, data [][]byte, destination []common.Address, amount []*big.Int, refunder common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "executeTransactionsWithGasRefunder", data, destination, amount, refunder)
}

// ExecuteTransactionsWithGasRefunder is a paid mutator transaction binding the contract method 0x97da185a.
//
// Solidity: function executeTransactionsWithGasRefunder(bytes[] data, address[] destination, uint256[] amount, address refunder) payable returns()
func (_Validator *ValidatorSession) ExecuteTransactionsWithGasRefunder(data [][]byte, destination []common.Address, amount []*big.Int, refunder common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactionsWithGasRefunder(&_Validator.TransactOpts, data, destination, amount, refunder)
}

// ExecuteTransactionsWithGasRefunder is a paid mutator transaction binding the contract method 0x97da185a.
//
// Solidity: function executeTransactionsWithGasRefunder(bytes[] data, address[] destination, uint256[] amount, address refunder) payable returns()
func (_Validator *ValidatorTransactorSession) ExecuteTransactionsWithGasRefunder(data [][]byte, destination []common.Address, amount []*big.Int, refunder common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ExecuteTransactionsWithGasRefunder(&_Validator.TransactOpts, data, destination, amount, refunder)
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

// ReturnOldDepositsWithGasRefunder is a paid mutator transaction binding the contract method 0xc083bc5b.
//
// Solidity: function returnOldDepositsWithGasRefunder(address rollup, address[] stakers, address refunder) returns()
func (_Validator *ValidatorTransactor) ReturnOldDepositsWithGasRefunder(opts *bind.TransactOpts, rollup common.Address, stakers []common.Address, refunder common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "returnOldDepositsWithGasRefunder", rollup, stakers, refunder)
}

// ReturnOldDepositsWithGasRefunder is a paid mutator transaction binding the contract method 0xc083bc5b.
//
// Solidity: function returnOldDepositsWithGasRefunder(address rollup, address[] stakers, address refunder) returns()
func (_Validator *ValidatorSession) ReturnOldDepositsWithGasRefunder(rollup common.Address, stakers []common.Address, refunder common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ReturnOldDepositsWithGasRefunder(&_Validator.TransactOpts, rollup, stakers, refunder)
}

// ReturnOldDepositsWithGasRefunder is a paid mutator transaction binding the contract method 0xc083bc5b.
//
// Solidity: function returnOldDepositsWithGasRefunder(address rollup, address[] stakers, address refunder) returns()
func (_Validator *ValidatorTransactorSession) ReturnOldDepositsWithGasRefunder(rollup common.Address, stakers []common.Address, refunder common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ReturnOldDepositsWithGasRefunder(&_Validator.TransactOpts, rollup, stakers, refunder)
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

// TimeoutChallengesWithGasRefunder is a paid mutator transaction binding the contract method 0x47336bf0.
//
// Solidity: function timeoutChallengesWithGasRefunder(address[] challenges, address refunder) returns()
func (_Validator *ValidatorTransactor) TimeoutChallengesWithGasRefunder(opts *bind.TransactOpts, challenges []common.Address, refunder common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "timeoutChallengesWithGasRefunder", challenges, refunder)
}

// TimeoutChallengesWithGasRefunder is a paid mutator transaction binding the contract method 0x47336bf0.
//
// Solidity: function timeoutChallengesWithGasRefunder(address[] challenges, address refunder) returns()
func (_Validator *ValidatorSession) TimeoutChallengesWithGasRefunder(challenges []common.Address, refunder common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TimeoutChallengesWithGasRefunder(&_Validator.TransactOpts, challenges, refunder)
}

// TimeoutChallengesWithGasRefunder is a paid mutator transaction binding the contract method 0x47336bf0.
//
// Solidity: function timeoutChallengesWithGasRefunder(address[] challenges, address refunder) returns()
func (_Validator *ValidatorTransactorSession) TimeoutChallengesWithGasRefunder(challenges []common.Address, refunder common.Address) (*types.Transaction, error) {
	return _Validator.Contract.TimeoutChallengesWithGasRefunder(&_Validator.TransactOpts, challenges, refunder)
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
