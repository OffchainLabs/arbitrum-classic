// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifierRPC

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
const AddressBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058205834f0d15aea77a3b9dbab8f1a54c4da46f7f5197ef831c04d46e149c8d88ab564736f6c63430005090032`

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ArbBalanceTrackerABI is the input ABI used to generate the binding from.
const ArbBalanceTrackerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"bytes32\"}],\"name\":\"getNFTTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"bytes32\"},{\"name\":\"_to\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"removeNFTToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"bytes32\"}],\"name\":\"getTokenBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"bytes32\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasNFT\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"bytes32\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"hasFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"bytes32\"},{\"name\":\"_to\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"addNFTToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"bytes32\"}],\"name\":\"depositEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ArbBalanceTrackerBin is the compiled bytecode used for deploying new contracts.
const ArbBalanceTrackerBin = `0x60806040819052600080546001600160a01b03191633178082556001600160a01b0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3611f75806100576000396000f3fe6080604052600436106101c25760003560e01c8063974765a1116100f7578063c311d04911610095578063da63d7b611610064578063da63d7b6146108b4578063dd62ed3e146108d1578063f2fde38b1461090c578063f3e414f81461093f576101c2565b8063c311d049146107cd578063cdf25dc1146107f7578063d29a4bf61461083c578063d339bf3c14610875576101c2565b8063a35be443116100d1578063a35be443146105e7578063a457c2d714610620578063a9059cbb14610659578063c246510614610692576101c2565b8063974765a11461053657806397feb92614610575578063a1db9782146105ae576101c2565b80632a8a8e7f11610164578063715018a61161013e578063715018a61461049c57806382512757146104b15780638da5cb5b146104f05780638f32d59b14610521576101c2565b80632a8a8e7f14610406578063395093511461043057806370a0823114610469576101c2565b80631240117b116101a05780631240117b1461031857806318160ddd1461035d57806323b872dd1461038457806326a6939a146103c7576101c2565b8063095ea7b3146101c75780630aa114ae146102145780630e227be3146102d7575b600080fd5b3480156101d357600080fd5b50610200600480360360408110156101ea57600080fd5b506001600160a01b038135169060200135610978565b604080519115158252519081900360200190f35b34801561022057600080fd5b5061023e6004803603602081101561023757600080fd5b50356109f5565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561028257818101518382015260200161026a565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156102c15781810151838201526020016102a9565b5050505090500194505050505060405180910390f35b3480156102e357600080fd5b50610316600480360360608110156102fa57600080fd5b508035906001600160a01b036020820135169060400135610b7a565b005b34801561032457600080fd5b506103166004803603608081101561033b57600080fd5b508035906020810135906001600160a01b036040820135169060600135610d58565b34801561036957600080fd5b50610372610d85565b60408051918252519081900360200190f35b34801561039057600080fd5b50610200600480360360608110156103a757600080fd5b506001600160a01b03813581169160208101359091169060400135610d8c565b3480156103d357600080fd5b50610316600480360360608110156103ea57600080fd5b508035906001600160a01b036020820135169060400135610e56565b34801561041257600080fd5b5061023e6004803603602081101561042957600080fd5b50356110fd565b34801561043c57600080fd5b506102006004803603604081101561045357600080fd5b506001600160a01b038135169060200135611226565b34801561047557600080fd5b506103726004803603602081101561048c57600080fd5b50356001600160a01b03166112d4565b3480156104a857600080fd5b506103166112ef565b3480156104bd57600080fd5b50610200600480360360608110156104d457600080fd5b506001600160a01b03813516906020810135906040013561134a565b3480156104fc57600080fd5b506105056113bf565b604080516001600160a01b039092168252519081900360200190f35b34801561052d57600080fd5b506102006113ce565b34801561054257600080fd5b506103166004803603606081101561055957600080fd5b508035906001600160a01b0360208201351690604001356113df565b34801561058157600080fd5b506103166004803603604081101561059857600080fd5b506001600160a01b0381351690602001356114d8565b3480156105ba57600080fd5b50610316600480360360408110156105d157600080fd5b506001600160a01b03813516906020013561158c565b3480156105f357600080fd5b506103726004803603604081101561060a57600080fd5b506001600160a01b038135169060200135611633565b34801561062c57600080fd5b506102006004803603604081101561064357600080fd5b506001600160a01b038135169060200135611695565b34801561066557600080fd5b506102006004803603604081101561067c57600080fd5b506001600160a01b0381351690602001356116de565b34801561069e57600080fd5b50610316600480360360608110156106b557600080fd5b813591908101906040810160208201356401000000008111156106d757600080fd5b8201836020820111156106e957600080fd5b8035906020019184602083028401116401000000008311171561070b57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561075b57600080fd5b82018360208201111561076d57600080fd5b8035906020019184602083028401116401000000008311171561078f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506116f4945050505050565b3480156107d957600080fd5b50610316600480360360208110156107f057600080fd5b5035611869565b34801561080357600080fd5b506103166004803603608081101561081a57600080fd5b508035906020810135906001600160a01b0360408201351690606001356118af565b34801561084857600080fd5b506103166004803603604081101561085f57600080fd5b506001600160a01b0381351690602001356118d6565b34801561088157600080fd5b506103166004803603606081101561089857600080fd5b508035906001600160a01b036020820135169060400135611971565b610316600480360360208110156108ca57600080fd5b5035611ad7565b3480156108dd57600080fd5b50610372600480360360408110156108f457600080fd5b506001600160a01b0381358116916020013516611ae6565b34801561091857600080fd5b506103166004803603602081101561092f57600080fd5b50356001600160a01b0316611b11565b34801561094b57600080fd5b506103166004803603604081101561096257600080fd5b506001600160a01b038135169060200135611b2b565b60006001600160a01b03831661098d57600080fd5b3360008181526002602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b6000818152600560205260408120606091829190805b6003830154811015610a4757826003018181548110610a2657fe5b60009182526020909120600260039092020101549190910190600101610a0b565b606082604051908082528060200260200182016040528015610a73578160200160208202803883390190505b509050606083604051908082528060200260200182016040528015610aa2578160200160208202803883390190505b50600093509050825b6003860154841015610b6c576000866003018581548110610ac857fe5b60009182526020822060039091020191505b6002820154811015610b5f57815485516001600160a01b0390911690869085908110610b0257fe5b60200260200101906001600160a01b031690816001600160a01b031681525050816002018181548110610b3157fe5b9060005260206000200154848481518110610b4857fe5b602090810291909101015260019283019201610ada565b5050600190930192610aab565b509095509350505050915091565b610b826113ce565b610b8b57600080fd5b80610b9557610d53565b60008381526005602090815260408083206001600160a01b0386168452918290529091205480610bf65760405162461bcd60e51b815260040180806020018281038252602d815260200180611ec2602d913960400191505060405180910390fd5b6000826001016001830381548110610c0a57fe5b906000526020600020906002020190508060010154841115610c5d5760405162461bcd60e51b8152600401808060200182810382526027815260200180611eef6027913960400191505060405180910390fd5b6001810154610c72908563ffffffff611bc116565b60018201819055610d4f5760018301805483918591600091906000198101908110610c9957fe5b600091825260208083206002909202909101546001600160a01b031683528201929092526040019020556001830180546000198101908110610cd757fe5b9060005260206000209060020201836001016001840381548110610cf757fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590881682528590526040812055830180546000190190610d4d9082611d23565b505b5050505b505050565b610d606113ce565b610d6957600080fd5b610d74848383610b7a565b610d7f8383836113df565b50505050565b6003545b90565b6001600160a01b0383166000908152600260209081526040808320338452909152812054610dc0908363ffffffff611bc116565b6001600160a01b0385166000908152600260209081526040808320338452909152902055610def848484611bd6565b6001600160a01b0384166000818152600260209081526040808320338085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b9392505050565b610e5e6113ce565b610e6757600080fd5b60008381526005602090815260408083206001600160a01b0386168452600281019092529091205480610ecb5760405162461bcd60e51b815260040180806020018281038252602b815260200180611f16602b913960400191505060405180910390fd5b6000826003016001830381548110610edf57fe5b600091825260208083208784526001600390930201918201905260409091205490915080610f54576040805162461bcd60e51b815260206004820181905260248201527f57616c6c657420646f6573206e6f74206f776e207370656369666963204e4654604482015290519081900360640190fd5b60028201805482916001850191600091906000198101908110610f7357fe5b600091825260208083209091015483528201929092526040019020556002820180546000198101908110610fa357fe5b9060005260206000200154826002016001830381548110610fc057fe5b60009182526020808320909101929092558681526001840190915260408120556002820180546000190190610ff59082611d4f565b506002820154610d4d576003840180548491600287019160009190600019810190811061101e57fe5b60009182526020808320600392830201546001600160a01b03168452830193909352604090910190209190915584018054600019810190811061105d57fe5b906000526020600020906003020184600301600185038154811061107d57fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b03909216919091178155600280830180546110c09284019190611d73565b5050506001600160a01b038616600090815260028501602052604081205560038401805460001901906110f39082611dc3565b5050505050505050565b600081815260056020908152604091829020600181015483518181528184028101909301909352606092839283918015611141578160200160208202803883390190505b50905060608151604051908082528060200260200182016040528015611171578160200160208202803883390190505b50905060005b825181101561121a5783600101818154811061118f57fe5b600091825260209091206002909102015483516001600160a01b03909116908490839081106111ba57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250508360010181815481106111e957fe5b90600052602060002090600202016001015482828151811061120757fe5b6020908102919091010152600101611177565b50909350915050915091565b60006001600160a01b03831661123b57600080fd5b3360009081526002602090815260408083206001600160a01b038716845290915290205461126f908363ffffffff611ca316565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b031660009081526001602052604090205490565b6112f76113ce565b61130057600080fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008281526005602090815260408083206001600160a01b0387168452600281019092528220548061138157600092505050610e4f565b81600301600182038154811061139357fe5b600091825260208083209683526001600390920290960101909452505060409091205415159392505050565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b6113e76113ce565b6113f057600080fd5b806113fa57610d53565b60008381526005602090815260408083206001600160a01b038616845291829052909120548061149057506040805180820182526001600160a01b0385811680835260006020808501828152600188810180548083018083559186528486209851600290910290980180546001600160a01b03191698909716979097178655905194019390935590815290849052919091208190555b60008260010160018303815481106114a457fe5b906000526020600020906002020190506114cb848260010154611ca390919063ffffffff16565b6001909101555050505050565b6001600160a01b0382163014156114ee57600080fd5b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038416916323b872dd9160648083019260209291908290030181600087803b15801561154357600080fd5b505af1158015611557573d6000803e3d6000fd5b505050506040513d602081101561156d57600080fd5b5061158890506001600160601b03193360601b1683836113df565b5050565b6001600160a01b0382163014156115a257600080fd5b6115ba6001600160601b03193360601b168383610b7a565b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b15801561160957600080fd5b505af115801561161d573d6000803e3d6000fd5b505050506040513d6020811015610d7f57600080fd5b60008181526005602090815260408083206001600160a01b03861684529182905282205480611667576000925050506109ef565b81600101600182038154811061167957fe5b9060005260206000209060020201600101549250505092915050565b60006001600160a01b0383166116aa57600080fd5b3360009081526002602090815260408083206001600160a01b038716845290915290205461126f908363ffffffff611bc116565b60006116eb338484611bd6565b50600192915050565b60005b82518110156117ae5782818151811061170c57fe5b602002602001015160146015811061172057fe5b1a60f81b6001600160f81b031916600160f81b14156117725761176d8484838151811061174957fe5b602002602001015160601c84848151811061176057fe5b6020026020010151610e56565b6117a6565b6117a68484838151811061178257fe5b602002602001015160601c84848151811061179957fe5b6020026020010151610b7a565b6001016116f7565b5060005b8251811015610d7f578281815181106117c757fe5b60200260200101516014601581106117db57fe5b1a60f81b6001600160f81b031916600160f81b141561182d576118288484838151811061180457fe5b602002602001015160601c84848151811061181b57fe5b6020026020010151611971565b611861565b6118618484838151811061183d57fe5b602002602001015160601c84848151811061185457fe5b60200260200101516113df565b6001016117b2565b6118826001600160601b03193360601b16600083610b7a565b604051339082156108fc029083906000818181858888f19350505050158015611588573d6000803e3d6000fd5b6118b76113ce565b6118c057600080fd5b6118cb848383610e56565b610d7f838383611971565b6001600160a01b0382163014156118ec57600080fd5b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038416916323b872dd91606480830192600092919082900301818387803b15801561194057600080fd5b505af1158015611954573d6000803e3d6000fd5b505050506115883360601b6bffffffffffffffffffffffff191683835b6119796113ce565b61198257600080fd5b60008381526005602090815260408083206001600160a01b0386168452600281019092529091205480611a5b576040805180820182526001600160a01b03861681528151600080825260208281019094526003860193830191905090528154600181018084556000938452602093849020835160039093020180546001600160a01b0319166001600160a01b039093169290921782558284015180519194611a3292600285019290910190611def565b50505060038301546001600160a01b038616600090815260028501602052604090208190559150505b6000826003016001830381548110611a6f57fe5b9060005260206000209060030201905080600101600085815260200190815260200160002054600014611aa157600080fd5b60028101805460018181018355600083815260208082209093018890559254968352909201909152604090209290925550505050565b611ae3816000346113df565b50565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b611b196113ce565b611b2257600080fd5b611ae381611cb5565b6001600160a01b038216301415611b4157600080fd5b611b596001600160601b03193360601b168383610e56565b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b158015611bad57600080fd5b505af1158015610d4f573d6000803e3d6000fd5b600082821115611bd057600080fd5b50900390565b6001600160a01b038216611be957600080fd5b6001600160a01b038316600090815260016020526040902054611c12908263ffffffff611bc116565b6001600160a01b038085166000908152600160205260408082209390935590841681522054611c47908263ffffffff611ca316565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082820183811015610e4f57600080fd5b6001600160a01b038116611cc857600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b815481835581811115610d5357600202816002028360005260206000209182019101610d539190611e2a565b815481835581811115610d5357600083815260209020610d53918101908301611e55565b828054828255906000526020600020908101928215611db35760005260206000209182015b82811115611db3578254825591600101919060010190611d98565b50611dbf929150611e55565b5090565b815481835581811115610d5357600302816003028360005260206000209182019101610d539190611e6f565b828054828255906000526020600020908101928215611db3579160200282015b82811115611db3578251825591602001919060010190611e0f565b610d8991905b80821115611dbf5780546001600160a01b031916815560006001820155600201611e30565b610d8991905b80821115611dbf5760008155600101611e5b565b610d8991905b80821115611dbf5780546001600160a01b03191681556000611e9a6002830182611ea3565b50600301611e75565b5080546000825590600052602060002090810190611ae39190611e5556fe57616c6c657420686173206e6f20636f696e732066726f6d20676976656e20455243323020636f6e747261637457616c6c657420646f6573206e6f74206f776e20656e6f75676820455243323020746f6b656e7357616c6c657420686173206e6f20636f696e732066726f6d20676976656e204e465420636f6e7472616374a265627a7a723058207250d69eaa9f0e1a34e06f77b1650405e1c18536f6564cd067244a3ec4562ef264736f6c63430005090032`

// DeployArbBalanceTracker deploys a new Ethereum contract, binding an instance of ArbBalanceTracker to it.
func DeployArbBalanceTracker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbBalanceTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbBalanceTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbBalanceTrackerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbBalanceTracker{ArbBalanceTrackerCaller: ArbBalanceTrackerCaller{contract: contract}, ArbBalanceTrackerTransactor: ArbBalanceTrackerTransactor{contract: contract}, ArbBalanceTrackerFilterer: ArbBalanceTrackerFilterer{contract: contract}}, nil
}

// ArbBalanceTracker is an auto generated Go binding around an Ethereum contract.
type ArbBalanceTracker struct {
	ArbBalanceTrackerCaller     // Read-only binding to the contract
	ArbBalanceTrackerTransactor // Write-only binding to the contract
	ArbBalanceTrackerFilterer   // Log filterer for contract events
}

// ArbBalanceTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbBalanceTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBalanceTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbBalanceTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBalanceTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbBalanceTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBalanceTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbBalanceTrackerSession struct {
	Contract     *ArbBalanceTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ArbBalanceTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbBalanceTrackerCallerSession struct {
	Contract *ArbBalanceTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ArbBalanceTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbBalanceTrackerTransactorSession struct {
	Contract     *ArbBalanceTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ArbBalanceTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbBalanceTrackerRaw struct {
	Contract *ArbBalanceTracker // Generic contract binding to access the raw methods on
}

// ArbBalanceTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbBalanceTrackerCallerRaw struct {
	Contract *ArbBalanceTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// ArbBalanceTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbBalanceTrackerTransactorRaw struct {
	Contract *ArbBalanceTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbBalanceTracker creates a new instance of ArbBalanceTracker, bound to a specific deployed contract.
func NewArbBalanceTracker(address common.Address, backend bind.ContractBackend) (*ArbBalanceTracker, error) {
	contract, err := bindArbBalanceTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbBalanceTracker{ArbBalanceTrackerCaller: ArbBalanceTrackerCaller{contract: contract}, ArbBalanceTrackerTransactor: ArbBalanceTrackerTransactor{contract: contract}, ArbBalanceTrackerFilterer: ArbBalanceTrackerFilterer{contract: contract}}, nil
}

// NewArbBalanceTrackerCaller creates a new read-only instance of ArbBalanceTracker, bound to a specific deployed contract.
func NewArbBalanceTrackerCaller(address common.Address, caller bind.ContractCaller) (*ArbBalanceTrackerCaller, error) {
	contract, err := bindArbBalanceTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbBalanceTrackerCaller{contract: contract}, nil
}

// NewArbBalanceTrackerTransactor creates a new write-only instance of ArbBalanceTracker, bound to a specific deployed contract.
func NewArbBalanceTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbBalanceTrackerTransactor, error) {
	contract, err := bindArbBalanceTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbBalanceTrackerTransactor{contract: contract}, nil
}

// NewArbBalanceTrackerFilterer creates a new log filterer instance of ArbBalanceTracker, bound to a specific deployed contract.
func NewArbBalanceTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbBalanceTrackerFilterer, error) {
	contract, err := bindArbBalanceTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbBalanceTrackerFilterer{contract: contract}, nil
}

// bindArbBalanceTracker binds a generic wrapper to an already deployed contract.
func bindArbBalanceTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbBalanceTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbBalanceTracker *ArbBalanceTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbBalanceTracker.Contract.ArbBalanceTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbBalanceTracker *ArbBalanceTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.ArbBalanceTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbBalanceTracker *ArbBalanceTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.ArbBalanceTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbBalanceTracker *ArbBalanceTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbBalanceTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ArbBalanceTracker.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ArbBalanceTracker.Contract.Allowance(&_ArbBalanceTracker.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ArbBalanceTracker.Contract.Allowance(&_ArbBalanceTracker.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ArbBalanceTracker.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ArbBalanceTracker.Contract.BalanceOf(&_ArbBalanceTracker.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ArbBalanceTracker.Contract.BalanceOf(&_ArbBalanceTracker.CallOpts, owner)
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x0aa114ae.
//
// Solidity: function getNFTTokens(bytes32 _owner) constant returns(address[], uint256[])
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) GetNFTTokens(opts *bind.CallOpts, _owner [32]byte) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbBalanceTracker.contract.Call(opts, out, "getNFTTokens", _owner)
	return *ret0, *ret1, err
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x0aa114ae.
//
// Solidity: function getNFTTokens(bytes32 _owner) constant returns(address[], uint256[])
func (_ArbBalanceTracker *ArbBalanceTrackerSession) GetNFTTokens(_owner [32]byte) ([]common.Address, []*big.Int, error) {
	return _ArbBalanceTracker.Contract.GetNFTTokens(&_ArbBalanceTracker.CallOpts, _owner)
}

// GetNFTTokens is a free data retrieval call binding the contract method 0x0aa114ae.
//
// Solidity: function getNFTTokens(bytes32 _owner) constant returns(address[], uint256[])
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) GetNFTTokens(_owner [32]byte) ([]common.Address, []*big.Int, error) {
	return _ArbBalanceTracker.Contract.GetNFTTokens(&_ArbBalanceTracker.CallOpts, _owner)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xa35be443.
//
// Solidity: function getTokenBalance(address _tokenContract, bytes32 _owner) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) GetTokenBalance(opts *bind.CallOpts, _tokenContract common.Address, _owner [32]byte) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ArbBalanceTracker.contract.Call(opts, out, "getTokenBalance", _tokenContract, _owner)
	return *ret0, err
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xa35be443.
//
// Solidity: function getTokenBalance(address _tokenContract, bytes32 _owner) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) GetTokenBalance(_tokenContract common.Address, _owner [32]byte) (*big.Int, error) {
	return _ArbBalanceTracker.Contract.GetTokenBalance(&_ArbBalanceTracker.CallOpts, _tokenContract, _owner)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xa35be443.
//
// Solidity: function getTokenBalance(address _tokenContract, bytes32 _owner) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) GetTokenBalance(_tokenContract common.Address, _owner [32]byte) (*big.Int, error) {
	return _ArbBalanceTracker.Contract.GetTokenBalance(&_ArbBalanceTracker.CallOpts, _tokenContract, _owner)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x2a8a8e7f.
//
// Solidity: function getTokenBalances(bytes32 _owner) constant returns(address[], uint256[])
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) GetTokenBalances(opts *bind.CallOpts, _owner [32]byte) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbBalanceTracker.contract.Call(opts, out, "getTokenBalances", _owner)
	return *ret0, *ret1, err
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x2a8a8e7f.
//
// Solidity: function getTokenBalances(bytes32 _owner) constant returns(address[], uint256[])
func (_ArbBalanceTracker *ArbBalanceTrackerSession) GetTokenBalances(_owner [32]byte) ([]common.Address, []*big.Int, error) {
	return _ArbBalanceTracker.Contract.GetTokenBalances(&_ArbBalanceTracker.CallOpts, _owner)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0x2a8a8e7f.
//
// Solidity: function getTokenBalances(bytes32 _owner) constant returns(address[], uint256[])
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) GetTokenBalances(_owner [32]byte) ([]common.Address, []*big.Int, error) {
	return _ArbBalanceTracker.Contract.GetTokenBalances(&_ArbBalanceTracker.CallOpts, _owner)
}

// HasNFT is a free data retrieval call binding the contract method 0x82512757.
//
// Solidity: function hasNFT(address _tokenContract, bytes32 _owner, uint256 _tokenId) constant returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) HasNFT(opts *bind.CallOpts, _tokenContract common.Address, _owner [32]byte, _tokenId *big.Int) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _ArbBalanceTracker.contract.Call(opts, out, "hasNFT", _tokenContract, _owner, _tokenId)
	return *ret0, err
}

// HasNFT is a free data retrieval call binding the contract method 0x82512757.
//
// Solidity: function hasNFT(address _tokenContract, bytes32 _owner, uint256 _tokenId) constant returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) HasNFT(_tokenContract common.Address, _owner [32]byte, _tokenId *big.Int) (bool, error) {
	return _ArbBalanceTracker.Contract.HasNFT(&_ArbBalanceTracker.CallOpts, _tokenContract, _owner, _tokenId)
}

// HasNFT is a free data retrieval call binding the contract method 0x82512757.
//
// Solidity: function hasNFT(address _tokenContract, bytes32 _owner, uint256 _tokenId) constant returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) HasNFT(_tokenContract common.Address, _owner [32]byte, _tokenId *big.Int) (bool, error) {
	return _ArbBalanceTracker.Contract.HasNFT(&_ArbBalanceTracker.CallOpts, _tokenContract, _owner, _tokenId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _ArbBalanceTracker.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) IsOwner() (bool, error) {
	return _ArbBalanceTracker.Contract.IsOwner(&_ArbBalanceTracker.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) IsOwner() (bool, error) {
	return _ArbBalanceTracker.Contract.IsOwner(&_ArbBalanceTracker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)

	out := ret0
	err := _ArbBalanceTracker.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) Owner() (common.Address, error) {
	return _ArbBalanceTracker.Contract.Owner(&_ArbBalanceTracker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) Owner() (common.Address, error) {
	return _ArbBalanceTracker.Contract.Owner(&_ArbBalanceTracker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ArbBalanceTracker.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) TotalSupply() (*big.Int, error) {
	return _ArbBalanceTracker.Contract.TotalSupply(&_ArbBalanceTracker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) TotalSupply() (*big.Int, error) {
	return _ArbBalanceTracker.Contract.TotalSupply(&_ArbBalanceTracker.CallOpts)
}

// AddNFTToken is a paid mutator transaction binding the contract method 0xd339bf3c.
//
// Solidity: function addNFTToken(bytes32 _user, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) AddNFTToken(opts *bind.TransactOpts, _user [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "addNFTToken", _user, _tokenContract, _tokenId)
}

// AddNFTToken is a paid mutator transaction binding the contract method 0xd339bf3c.
//
// Solidity: function addNFTToken(bytes32 _user, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) AddNFTToken(_user [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.AddNFTToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _tokenId)
}

// AddNFTToken is a paid mutator transaction binding the contract method 0xd339bf3c.
//
// Solidity: function addNFTToken(bytes32 _user, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) AddNFTToken(_user [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.AddNFTToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _tokenId)
}

// AddToken is a paid mutator transaction binding the contract method 0x974765a1.
//
// Solidity: function addToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) AddToken(opts *bind.TransactOpts, _user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "addToken", _user, _tokenContract, _value)
}

// AddToken is a paid mutator transaction binding the contract method 0x974765a1.
//
// Solidity: function addToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) AddToken(_user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.AddToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _value)
}

// AddToken is a paid mutator transaction binding the contract method 0x974765a1.
//
// Solidity: function addToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) AddToken(_user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.AddToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.Approve(&_ArbBalanceTracker.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.Approve(&_ArbBalanceTracker.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.DecreaseAllowance(&_ArbBalanceTracker.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.DecreaseAllowance(&_ArbBalanceTracker.TransactOpts, spender, subtractedValue)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) DepositERC20(opts *bind.TransactOpts, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "depositERC20", _tokenContract, _value)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) DepositERC20(_tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.DepositERC20(&_ArbBalanceTracker.TransactOpts, _tokenContract, _value)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) DepositERC20(_tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.DepositERC20(&_ArbBalanceTracker.TransactOpts, _tokenContract, _value)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) DepositERC721(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "depositERC721", _tokenContract, _tokenId)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) DepositERC721(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.DepositERC721(&_ArbBalanceTracker.TransactOpts, _tokenContract, _tokenId)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) DepositERC721(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.DepositERC721(&_ArbBalanceTracker.TransactOpts, _tokenContract, _tokenId)
}

// DepositEth is a paid mutator transaction binding the contract method 0xda63d7b6.
//
// Solidity: function depositEth(bytes32 _destination) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) DepositEth(opts *bind.TransactOpts, _destination [32]byte) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "depositEth", _destination)
}

// DepositEth is a paid mutator transaction binding the contract method 0xda63d7b6.
//
// Solidity: function depositEth(bytes32 _destination) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) DepositEth(_destination [32]byte) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.DepositEth(&_ArbBalanceTracker.TransactOpts, _destination)
}

// DepositEth is a paid mutator transaction binding the contract method 0xda63d7b6.
//
// Solidity: function depositEth(bytes32 _destination) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) DepositEth(_destination [32]byte) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.DepositEth(&_ArbBalanceTracker.TransactOpts, _destination)
}

// HasFunds is a paid mutator transaction binding the contract method 0xc2465106.
//
// Solidity: function hasFunds(bytes32 _vmId, bytes21[] _tokenTypes, uint256[] _amounts) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) HasFunds(opts *bind.TransactOpts, _vmId [32]byte, _tokenTypes [][21]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "hasFunds", _vmId, _tokenTypes, _amounts)
}

// HasFunds is a paid mutator transaction binding the contract method 0xc2465106.
//
// Solidity: function hasFunds(bytes32 _vmId, bytes21[] _tokenTypes, uint256[] _amounts) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) HasFunds(_vmId [32]byte, _tokenTypes [][21]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.HasFunds(&_ArbBalanceTracker.TransactOpts, _vmId, _tokenTypes, _amounts)
}

// HasFunds is a paid mutator transaction binding the contract method 0xc2465106.
//
// Solidity: function hasFunds(bytes32 _vmId, bytes21[] _tokenTypes, uint256[] _amounts) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) HasFunds(_vmId [32]byte, _tokenTypes [][21]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.HasFunds(&_ArbBalanceTracker.TransactOpts, _vmId, _tokenTypes, _amounts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.IncreaseAllowance(&_ArbBalanceTracker.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.IncreaseAllowance(&_ArbBalanceTracker.TransactOpts, spender, addedValue)
}

// RemoveNFTToken is a paid mutator transaction binding the contract method 0x26a6939a.
//
// Solidity: function removeNFTToken(bytes32 _user, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) RemoveNFTToken(opts *bind.TransactOpts, _user [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "removeNFTToken", _user, _tokenContract, _tokenId)
}

// RemoveNFTToken is a paid mutator transaction binding the contract method 0x26a6939a.
//
// Solidity: function removeNFTToken(bytes32 _user, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) RemoveNFTToken(_user [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.RemoveNFTToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _tokenId)
}

// RemoveNFTToken is a paid mutator transaction binding the contract method 0x26a6939a.
//
// Solidity: function removeNFTToken(bytes32 _user, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) RemoveNFTToken(_user [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.RemoveNFTToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _tokenId)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x0e227be3.
//
// Solidity: function removeToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) RemoveToken(opts *bind.TransactOpts, _user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "removeToken", _user, _tokenContract, _value)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x0e227be3.
//
// Solidity: function removeToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) RemoveToken(_user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.RemoveToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _value)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x0e227be3.
//
// Solidity: function removeToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) RemoveToken(_user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.RemoveToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.RenounceOwnership(&_ArbBalanceTracker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.RenounceOwnership(&_ArbBalanceTracker.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.Transfer(&_ArbBalanceTracker.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.Transfer(&_ArbBalanceTracker.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferFrom(&_ArbBalanceTracker.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferFrom(&_ArbBalanceTracker.TransactOpts, from, to, value)
}

// TransferNFT is a paid mutator transaction binding the contract method 0xcdf25dc1.
//
// Solidity: function transferNFT(bytes32 _from, bytes32 _to, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) TransferNFT(opts *bind.TransactOpts, _from [32]byte, _to [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "transferNFT", _from, _to, _tokenContract, _tokenId)
}

// TransferNFT is a paid mutator transaction binding the contract method 0xcdf25dc1.
//
// Solidity: function transferNFT(bytes32 _from, bytes32 _to, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) TransferNFT(_from [32]byte, _to [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferNFT(&_ArbBalanceTracker.TransactOpts, _from, _to, _tokenContract, _tokenId)
}

// TransferNFT is a paid mutator transaction binding the contract method 0xcdf25dc1.
//
// Solidity: function transferNFT(bytes32 _from, bytes32 _to, address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) TransferNFT(_from [32]byte, _to [32]byte, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferNFT(&_ArbBalanceTracker.TransactOpts, _from, _to, _tokenContract, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferOwnership(&_ArbBalanceTracker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferOwnership(&_ArbBalanceTracker.TransactOpts, newOwner)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1240117b.
//
// Solidity: function transferToken(bytes32 _from, bytes32 _to, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) TransferToken(opts *bind.TransactOpts, _from [32]byte, _to [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "transferToken", _from, _to, _tokenContract, _value)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1240117b.
//
// Solidity: function transferToken(bytes32 _from, bytes32 _to, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) TransferToken(_from [32]byte, _to [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferToken(&_ArbBalanceTracker.TransactOpts, _from, _to, _tokenContract, _value)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1240117b.
//
// Solidity: function transferToken(bytes32 _from, bytes32 _to, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) TransferToken(_from [32]byte, _to [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferToken(&_ArbBalanceTracker.TransactOpts, _from, _to, _tokenContract, _value)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) WithdrawERC20(opts *bind.TransactOpts, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "withdrawERC20", _tokenContract, _value)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) WithdrawERC20(_tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.WithdrawERC20(&_ArbBalanceTracker.TransactOpts, _tokenContract, _value)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) WithdrawERC20(_tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.WithdrawERC20(&_ArbBalanceTracker.TransactOpts, _tokenContract, _value)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) WithdrawERC721(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "withdrawERC721", _tokenContract, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) WithdrawERC721(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.WithdrawERC721(&_ArbBalanceTracker.TransactOpts, _tokenContract, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _tokenContract, uint256 _tokenId) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) WithdrawERC721(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.WithdrawERC721(&_ArbBalanceTracker.TransactOpts, _tokenContract, _tokenId)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) WithdrawEth(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "withdrawEth", _value)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) WithdrawEth(_value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.WithdrawEth(&_ArbBalanceTracker.TransactOpts, _value)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) WithdrawEth(_value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.WithdrawEth(&_ArbBalanceTracker.TransactOpts, _value)
}

// ArbBalanceTrackerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ArbBalanceTracker contract.
type ArbBalanceTrackerApprovalIterator struct {
	Event *ArbBalanceTrackerApproval // Event containing the contract specifics and raw log

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
func (it *ArbBalanceTrackerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbBalanceTrackerApproval)
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
		it.Event = new(ArbBalanceTrackerApproval)
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
func (it *ArbBalanceTrackerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbBalanceTrackerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbBalanceTrackerApproval represents a Approval event raised by the ArbBalanceTracker contract.
type ArbBalanceTrackerApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ArbBalanceTrackerApprovalIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ArbBalanceTracker.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ArbBalanceTrackerApprovalIterator{contract: _ArbBalanceTracker.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ArbBalanceTrackerApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ArbBalanceTracker.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbBalanceTrackerApproval)
				if err := _ArbBalanceTracker.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ArbBalanceTrackerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ArbBalanceTracker contract.
type ArbBalanceTrackerOwnershipTransferredIterator struct {
	Event *ArbBalanceTrackerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArbBalanceTrackerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbBalanceTrackerOwnershipTransferred)
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
		it.Event = new(ArbBalanceTrackerOwnershipTransferred)
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
func (it *ArbBalanceTrackerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbBalanceTrackerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbBalanceTrackerOwnershipTransferred represents a OwnershipTransferred event raised by the ArbBalanceTracker contract.
type ArbBalanceTrackerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArbBalanceTrackerOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArbBalanceTracker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArbBalanceTrackerOwnershipTransferredIterator{contract: _ArbBalanceTracker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbBalanceTrackerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArbBalanceTracker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbBalanceTrackerOwnershipTransferred)
				if err := _ArbBalanceTracker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ArbBalanceTrackerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ArbBalanceTracker contract.
type ArbBalanceTrackerTransferIterator struct {
	Event *ArbBalanceTrackerTransfer // Event containing the contract specifics and raw log

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
func (it *ArbBalanceTrackerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbBalanceTrackerTransfer)
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
		it.Event = new(ArbBalanceTrackerTransfer)
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
func (it *ArbBalanceTrackerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbBalanceTrackerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbBalanceTrackerTransfer represents a Transfer event raised by the ArbBalanceTracker contract.
type ArbBalanceTrackerTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbBalanceTrackerTransferIterator, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbBalanceTracker.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbBalanceTrackerTransferIterator{contract: _ArbBalanceTracker.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ArbBalanceTrackerTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbBalanceTracker.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbBalanceTrackerTransfer)
				if err := _ArbBalanceTracker.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ArbProtocolABI is the input ABI used to generate the binding from.
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[5]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"unanimousAssertHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"name\":\"_lastLogHash\",\"type\":\"bytes32\"},{\"name\":\"_totalMessageValueAmounts\",\"type\":\"uint256[]\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_dest\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"bytes32\"}],\"name\":\"generateSentMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"countSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_signatures\",\"type\":\"bytes\"},{\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"parseSignature\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageDataHashes\",\"type\":\"bytes32[]\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageValueAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"generateLastMessageHashStub\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"bytes32\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"},{\"name\":\"_newMessage\",\"type\":\"bytes32\"}],\"name\":\"appendInboxPendingMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"recoverAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_inboxHash\",\"type\":\"bytes32\"},{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"}],\"name\":\"appendInboxMessages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
const ArbProtocolBin = `0x61205a610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100d95760003560e01c80633e28559811610096578063ccf69dd711610070578063ccf69dd714610f17578063d78d18ea14610f50578063f0c8e96914610f73578063f11fcc261461101e576100d9565b80633e28559814610a48578063b31d63cc14610ba6578063b327749514610c6e576100d9565b8063014bba5b146100de5780630f89fbff146103f957806320903721146105ee57806325200160146106b95780632a0500d81461096557806333ae3ad0146109a4575b600080fd5b6103e760048036036101808110156100f557600080fd5b810190808060a001906005806020026040519081016040528092919082600560200280828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561016b57600080fd5b82018360208201111561017d57600080fd5b803590602001918460208302840111600160201b8311171561019e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101ed57600080fd5b8201836020820111156101ff57600080fd5b803590602001918460018302840111600160201b8311171561022057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561027257600080fd5b82018360208201111561028457600080fd5b803590602001918460208302840111600160201b831117156102a557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102f457600080fd5b82018360208201111561030657600080fd5b803590602001918460208302840111600160201b8311171561032757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561037657600080fd5b82018360208201111561038857600080fd5b803590602001918460208302840111600160201b831117156103a957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611041945050505050565b60408051918252519081900360200190f35b61059e6004803603606081101561040f57600080fd5b810190602081018135600160201b81111561042957600080fd5b82018360208201111561043b57600080fd5b803590602001918460208302840111600160201b8311171561045c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104ab57600080fd5b8201836020820111156104bd57600080fd5b803590602001918460208302840111600160201b831117156104de57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561052d57600080fd5b82018360208201111561053f57600080fd5b803590602001918460208302840111600160201b8311171561056057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506111da945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156105da5781810151838201526020016105c2565b505050509050019250505060405180910390f35b6103e7600480360360e081101561060457600080fd5b81359163ffffffff6020820135169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561064857600080fd5b82018360208201111561065a57600080fd5b803590602001918460208302840111600160201b8311171561067b57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061135b945050505050565b6103e7600480360360a08110156106cf57600080fd5b810190602081018135600160201b8111156106e957600080fd5b8201836020820111156106fb57600080fd5b803590602001918460208302840111600160201b8311171561071c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561076b57600080fd5b82018360208201111561077d57600080fd5b803590602001918460018302840111600160201b8311171561079e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156107f057600080fd5b82018360208201111561080257600080fd5b803590602001918460208302840111600160201b8311171561082357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561087257600080fd5b82018360208201111561088457600080fd5b803590602001918460208302840111600160201b831117156108a557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156108f457600080fd5b82018360208201111561090657600080fd5b803590602001918460208302840111600160201b8311171561092757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506113c6945050505050565b6103e7600480360360a081101561097b57600080fd5b508035906020810135906001600160581b03196040820135169060608101359060800135611588565b6103e7600480360360208110156109ba57600080fd5b810190602081018135600160201b8111156109d457600080fd5b8201836020820111156109e657600080fd5b803590602001918460018302840111600160201b83111715610a0757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611766945050505050565b6103e7600480360360c0811015610a5e57600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594909390925060408101915060200135600160201b811115610ab357600080fd5b820183602082011115610ac557600080fd5b803590602001918460208302840111600160201b83111715610ae657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b3557600080fd5b820183602082011115610b4757600080fd5b803590602001918460208302840111600160201b83111715610b6857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611795945050505050565b610c4c60048036036040811015610bbc57600080fd5b810190602081018135600160201b811115610bd657600080fd5b820183602082011115610be857600080fd5b803590602001918460018302840111600160201b83111715610c0957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250611880915050565b6040805160ff9094168452602084019290925282820152519081900360600190f35b6103e7600480360360a0811015610c8457600080fd5b810190602081018135600160201b811115610c9e57600080fd5b820183602082011115610cb057600080fd5b803590602001918460208302840111600160201b83111715610cd157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d2057600080fd5b820183602082011115610d3257600080fd5b803590602001918460208302840111600160201b83111715610d5357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610da257600080fd5b820183602082011115610db457600080fd5b803590602001918460208302840111600160201b83111715610dd557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e2457600080fd5b820183602082011115610e3657600080fd5b803590602001918460208302840111600160201b83111715610e5757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610ea657600080fd5b820183602082011115610eb857600080fd5b803590602001918460208302840111600160201b83111715610ed957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506118d2945050505050565b6103e760048036036080811015610f2d57600080fd5b508035906001600160581b031960208201351690604081013590606001356119b2565b6103e760048036036040811015610f6657600080fd5b5080359060200135611a90565b61059e60048036036040811015610f8957600080fd5b81359190810190604081016020820135600160201b811115610faa57600080fd5b820183602082011115610fbc57600080fd5b803590602001918460018302840111600160201b83111715610fdd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611ad4945050505050565b6103e76004803603604081101561103457600080fd5b5080359060200135611c7a565b6000878787878787876040516020018088600560200280838360005b8381101561107557818101518382015260200161105d565b5050505090500187600260200280838360005b838110156110a0578181015183820152602001611088565b50505050905001868051906020019060200280838360005b838110156110d05781810151838201526020016110b8565b5050505090500185805190602001908083835b602083106111025780518252601f1990920191602091820191016110e3565b51815160209384036101000a60001901801990921691161790528751919093019287810192500280838360005b8381101561114757818101518382015260200161112f565b50505050905001838051906020019060200280838360005b8381101561117757818101518382015260200161115f565b50505050905001828051906020019060200280838360005b838110156111a757818101518382015260200161118f565b50505050905001975050505050505050604051602081830303815290604052805190602001209050979650505050505050565b6060808451604051908082528060200260200182016040528015611208578160200160208202803883390190505b50905060005b8451811015611352578585828151811061122457fe5b602002602001015161ffff168151811061123a57fe5b602002602001015160146015811061124e57fe5b1a60f81b6001600160f81b0319166112ae5783818151811061126c57fe5b60200260200101518286838151811061128157fe5b602002602001015161ffff168151811061129757fe5b60200260200101818151019150818152505061134a565b818582815181106112bb57fe5b602002602001015161ffff16815181106112d157fe5b60200260200101516000146112e557600080fd5b8381815181106112f157fe5b60200260200101516000141561130657600080fd5b83818151811061131257fe5b60200260200101518286838151811061132757fe5b602002602001015161ffff168151811061133d57fe5b6020026020010181815250505b60010161120e565b50949350505050565b600087878787878787604051602001808881526020018763ffffffff1663ffffffff1660e01b815260040186815260200185815260200184815260200183815260200182805190602001906020028083836000838110156111a757818101518382015260200161118f565b600081518351146113d657600080fd5b83518351146113e457600080fd5b60008080805b865181101561157b5773__$0d86abb4a722a612872fb80f4c7e7e95bd$__63615c39b08a856040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015611461578181015183820152602001611449565b50505050905090810190601f16801561148e5780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b1580156114ab57600080fd5b505af41580156114bf573d6000803e3d6000fd5b505050506040513d60408110156114d557600080fd5b508051602090910151895191945092506115449083908c908b90859081106114f957fe5b602002602001015161ffff168151811061150f57fe5b602002602001015189848151811061152357fe5b602002602001015189858151811061153757fe5b60200260200101516119b2565b60408051602080820197909752808201839052815180820383018152606090910190915280519501949094209391506001016113ea565b5050505095945050505050565b60408051602080820188905281830187905260608083018690526001600160581b03198716608084015283518084036075018152609584018086528151919093012060048084526101358501909552600094909391929160b5015b6115eb611fef565b8152602001906001900390816115e357905050905061160987611c94565b8160008151811061161657fe5b602002602001018190525061162a42611cee565b8160018151811061163757fe5b602002602001018190525061164b43611cee565b8160028151811061165857fe5b602090810291909101015261166c82611cee565b8160038151811061167957fe5b602090810291909101015260408051600480825260a08201909252606091816020015b6116a4611fef565b81526020019060019003908161169c5790505090506116c282611d48565b816000815181106116cf57fe5b60209081029190910101526116e385611cee565b816001815181106116f057fe5b602002602001018190525061170486611cee565b8160028151811061171157fe5b602090810291909101015261172f6001600160581b03198816611cee565b8160038151811061173c57fe5b602002602001018190525061175861175382611d48565b611d88565b519998505050505050505050565b6000604182518161177357fe5b061561178057600061178d565b604182518161178b57fe5b045b90505b919050565b600085858260200201518660016020020151868686604051602001808781526020018667ffffffffffffffff1667ffffffffffffffff1660c01b81526008018567ffffffffffffffff1667ffffffffffffffff1660c01b8152600801848152602001838051906020019060200280838360005b83811015611820578181015183820152602001611808565b50505050905001828051906020019060200280838360005b83811015611850578181015183820152602001611838565b50505050905001965050505050505060405160208183030381529060405280519060200120905095945050505050565b604180820283810160208101516040820151919093015160ff169291601b8410156118ac57601b840193505b8360ff16601b14806118c157508360ff16601c145b6118ca57600080fd5b509250925092565b600083518551146118e257600080fd5b82518551146118f057600080fd5b81518551146118fe57600080fd5b600080805b87518110156119a55761196e88828151811061191b57fe5b60200260200101518a89848151811061193057fe5b602002602001015161ffff168151811061194657fe5b602002602001015188848151811061195a57fe5b602002602001015188858151811061153757fe5b6040805160208082019690965280820183905281518082038301815260609091019091528051940193909320929150600101611903565b5090979650505050505050565b60408051600480825260a0820190925260009160609190816020015b6119d6611fef565b8152602001906001900390816119ce5790505090506119f486611c94565b81600081518110611a0157fe5b6020908102919091010152611a1583611cee565b81600181518110611a2257fe5b6020026020010181905250611a3684611cee565b81600281518110611a4357fe5b6020908102919091010152611a616001600160581b03198616611cee565b81600381518110611a6e57fe5b6020026020010181905250611a8561175382611d48565b519695505050505050565b6000611acd6040518060600160405280611aaa6000611cee565b8152602001611ab886611c94565b8152602001611ac685611c94565b9052611e3b565b9392505050565b6060600080600080611ae586611766565b9050606081604051908082528060200260200182016040528015611b13578160200160208202803883390190505b50905060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818a6040516020018083805190602001908083835b60208310611b865780518252601f199092019160209182019101611b67565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190528251920191909120925060009150505b84811015611c6b57611bd68a82611880565b6040805160008152602080820180845288905260ff86168284015260608201859052608082018490529151949c50929a5090985060019260a080840193601f198301929081900390910190855afa158015611c35573d6000803e3d6000fd5b50505060206040510351848281518110611c4b57fe5b6001600160a01b0390921660209283029190910190910152600101611bc4565b50919998505050505050505050565b6000611acd6040518060600160405280611aaa6001611cee565b611c9c611fef565b604080516060810182528381528151600080825260208281019094529192830191611cdd565b611cca611fef565b815260200190600190039081611cc25790505b508152600260209091015292915050565b611cf6611fef565b604080516060810182528381528151600080825260208281019094529192830191611d37565b611d24611fef565b815260200190600190039081611d1c5790505b508152600060209091015292915050565b611d50611fef565b611d5a8251611eba565b611d6357600080fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b611d90612013565b6040820151600c60ff90911610611da657600080fd5b604082015160ff16611dd3576040518060200160405280611dca8460000151611ec1565b90529050611790565b604082015160ff1660021415611df85750604080516020810190915281518152611790565b600360ff16826040015160ff1610158015611e1c57506040820151600c60ff909116105b15611e39576040518060200160405280611dca8460200151611ee5565bfe5b6040805160038082526080820190925260009160609190816020015b611e5f611fef565b815260200190600190039081611e5757905050905060005b8151811015611eb057838160038110611e8c57fe5b6020020151828281518110611e9d57fe5b6020908102919091010152600101611e77565b50611acd81611ee5565b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600882511115611ef657600080fd5b60608251604051908082528060200260200182016040528015611f23578160200160208202803883390190505b50905060005b8151811015611f7d57611f3a612013565b611f56858381518110611f4957fe5b6020026020010151611d88565b90508060000151838381518110611f6957fe5b602090810291909101015250600101611f29565b508251600360ff160181604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611fc6578181015183820152602001611fae565b505050509050019250505060405160208183030381529060405280519060200120915050919050565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820ff2289cb1e97494199d4938bf2d1ce0af4c2b5efce9fb90eba94c766a5e8b8dc64736f6c63430005090032`

// DeployArbProtocol deploys a new Ethereum contract, binding an instance of ArbProtocol to it.
func DeployArbProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbProtocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// ArbProtocol is an auto generated Go binding around an Ethereum contract.
type ArbProtocol struct {
	ArbProtocolCaller     // Read-only binding to the contract
	ArbProtocolTransactor // Write-only binding to the contract
	ArbProtocolFilterer   // Log filterer for contract events
}

// ArbProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbProtocolSession struct {
	Contract     *ArbProtocol      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbProtocolCallerSession struct {
	Contract *ArbProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ArbProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbProtocolTransactorSession struct {
	Contract     *ArbProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ArbProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbProtocolRaw struct {
	Contract *ArbProtocol // Generic contract binding to access the raw methods on
}

// ArbProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbProtocolCallerRaw struct {
	Contract *ArbProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ArbProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbProtocolTransactorRaw struct {
	Contract *ArbProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbProtocol creates a new instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocol(address common.Address, backend bind.ContractBackend) (*ArbProtocol, error) {
	contract, err := bindArbProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// NewArbProtocolCaller creates a new read-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolCaller(address common.Address, caller bind.ContractCaller) (*ArbProtocolCaller, error) {
	contract, err := bindArbProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolCaller{contract: contract}, nil
}

// NewArbProtocolTransactor creates a new write-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbProtocolTransactor, error) {
	contract, err := bindArbProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolTransactor{contract: contract}, nil
}

// NewArbProtocolFilterer creates a new log filterer instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbProtocolFilterer, error) {
	contract, err := bindArbProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolFilterer{contract: contract}, nil
}

// bindArbProtocol binds a generic wrapper to an already deployed contract.
func bindArbProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.ArbProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transact(opts, method, params...)
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) AppendInboxMessages(opts *bind.CallOpts, _inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "appendInboxMessages", _inboxHash, _pendingMessages)
	return *ret0, err
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) AppendInboxMessages(_inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxMessages(&_ArbProtocol.CallOpts, _inboxHash, _pendingMessages)
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) AppendInboxMessages(_inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxMessages(&_ArbProtocol.CallOpts, _inboxHash, _pendingMessages)
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) AppendInboxPendingMessage(opts *bind.CallOpts, _pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "appendInboxPendingMessage", _pendingMessages, _newMessage)
	return *ret0, err
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) AppendInboxPendingMessage(_pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxPendingMessage(&_ArbProtocol.CallOpts, _pendingMessages, _newMessage)
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) AppendInboxPendingMessage(_pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxPendingMessage(&_ArbProtocol.CallOpts, _pendingMessages, _newMessage)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCaller) CalculateBeforeValues(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	ret0 := new([]*big.Int)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "calculateBeforeValues", _tokenTypes, _messageTokenNums, _messageAmounts)
	return *ret0, err
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCallerSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_ArbProtocol *ArbProtocolCaller) CountSignatures(opts *bind.CallOpts, _signatures []byte) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "countSignatures", _signatures)
	return *ret0, err
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_ArbProtocol *ArbProtocolSession) CountSignatures(_signatures []byte) (*big.Int, error) {
	return _ArbProtocol.Contract.CountSignatures(&_ArbProtocol.CallOpts, _signatures)
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_ArbProtocol *ArbProtocolCallerSession) CountSignatures(_signatures []byte) (*big.Int, error) {
	return _ArbProtocol.Contract.CountSignatures(&_ArbProtocol.CallOpts, _signatures)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHash", _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _messageDataHashes, uint16[] _messageTokenNum, uint256[] _messageValueAmounts, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHashStub(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageDataHashes [][32]byte, _messageTokenNum []uint16, _messageValueAmounts []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHashStub", _tokenTypes, _messageDataHashes, _messageTokenNum, _messageValueAmounts, _messageDestination)
	return *ret0, err
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _messageDataHashes, uint16[] _messageTokenNum, uint256[] _messageValueAmounts, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _messageDataHashes [][32]byte, _messageTokenNum []uint16, _messageValueAmounts []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _messageDataHashes, _messageTokenNum, _messageValueAmounts, _messageDestination)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _messageDataHashes, uint16[] _messageTokenNum, uint256[] _messageValueAmounts, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _messageDataHashes [][32]byte, _messageTokenNum []uint16, _messageValueAmounts []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _messageDataHashes, _messageTokenNum, _messageValueAmounts, _messageDestination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateSentMessageHash(opts *bind.CallOpts, _dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateSentMessageHash", _dest, _data, _tokenType, _value, _sender)
	return *ret0, err
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateSentMessageHash(_dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateSentMessageHash(&_ArbProtocol.CallOpts, _dest, _data, _tokenType, _value, _sender)
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateSentMessageHash(_dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateSentMessageHash(&_ArbProtocol.CallOpts, _dest, _data, _tokenType, _value, _sender)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_ArbProtocol *ArbProtocolCaller) ParseSignature(opts *bind.CallOpts, _signatures []byte, _pos *big.Int) (struct {
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
	err := _ArbProtocol.contract.Call(opts, out, "parseSignature", _signatures, _pos)
	return *ret, err
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_ArbProtocol *ArbProtocolSession) ParseSignature(_signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _ArbProtocol.Contract.ParseSignature(&_ArbProtocol.CallOpts, _signatures, _pos)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_ArbProtocol *ArbProtocolCallerSession) ParseSignature(_signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _ArbProtocol.Contract.ParseSignature(&_ArbProtocol.CallOpts, _signatures, _pos)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_ArbProtocol *ArbProtocolCaller) RecoverAddresses(opts *bind.CallOpts, _messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	ret0 := new([]common.Address)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "recoverAddresses", _messageHash, _signatures)
	return *ret0, err
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_ArbProtocol *ArbProtocolSession) RecoverAddresses(_messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	return _ArbProtocol.Contract.RecoverAddresses(&_ArbProtocol.CallOpts, _messageHash, _signatures)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_ArbProtocol *ArbProtocolCallerSession) RecoverAddresses(_messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	return _ArbProtocol.Contract.RecoverAddresses(&_ArbProtocol.CallOpts, _messageHash, _signatures)
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) UnanimousAssertHash(opts *bind.CallOpts, _fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "unanimousAssertHash", _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
	return *ret0, err
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) UnanimousAssertHash(_fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.UnanimousAssertHash(&_ArbProtocol.CallOpts, _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) UnanimousAssertHash(_fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.UnanimousAssertHash(&_ArbProtocol.CallOpts, _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// ArbValueABI is the input ABI used to generate the binding from.
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"get_next_valid_value\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserialize_valid_value_hash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserialize_value_hash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueBin is the compiled bytecode used for deploying new contracts.
const ArbValueBin = `0x610b6c610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c806353409fab1161006557806353409fab14610221578063615c39b01461024757806392516ac714610308578063b2b9dc62146103ae57610092565b80631667b41114610097578063264f384b146100c6578063364df277146100f25780634d00ef7a146100fa575b600080fd5b6100b4600480360360208110156100ad57600080fd5b50356103df565b60408051918252519081900360200190f35b6100b4600480360360608110156100dc57600080fd5b5060ff8135169060208101359060400135610405565b6100b4610457565b6101a26004803603604081101561011057600080fd5b81019060208101813564010000000081111561012b57600080fd5b82018360208201111561013d57600080fd5b8035906020019184600183028401116401000000008311171561015f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506104ca915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101e55781810151838201526020016101cd565b50505050905090810190601f1680156102125780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100b46004803603604081101561023757600080fd5b5060ff8135169060200135610517565b6102ef6004803603604081101561025d57600080fd5b81019060208101813564010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061055e915050565b6040805192835260208301919091528051918290030190f35b6100b46004803603602081101561031e57600080fd5b81019060208101813564010000000081111561033957600080fd5b82018360208201111561034b57600080fd5b8035906020019184600183028401116401000000008311171561036d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105a1945050505050565b6103cb600480360360208110156103c457600080fd5b50356105dd565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156104a357818101518382015260200161048b565b50505050905001925050506040516020818303038152906040528051906020012091505090565b600060606000806104d9610b01565b6104e387876105e4565b9194509250905082156104f557600080fd5b81610509888880840363ffffffff6106dc16565b945094505050509250929050565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b60008060008061056c610b01565b61057687876105e4565b91945092509050821561058857600080fd5b816105928261075c565b51909890975095505050505050565b600080806105ad610b01565b6105b88560006105e4565b9194509250905082156105ca57600080fd5b6105d38161075c565b5195945050505050565b6008101590565b6000806105ef610b01565b60008585815181106105fd57fe5b016020015160019095019460f81c905060008161063e5761061e878761080f565b909650905060008661062f83610832565b919650945092506106d5915050565b60ff82166002141561066557610654878761080f565b909650905060008661062f8361088c565b600360ff83161080159061067c5750600c60ff8316105b156106b757600219820160606000610695838b8b6108e6565b909a509250905080896106a78461099b565b97509750975050505050506106d5565b8160ff166127100160006106cb6000610832565b9196509450925050505b9250925092565b6060818301845110156106ee57600080fd5b60608215801561070957604051915060208201604052610753565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561074257805183526020928301920161072a565b5050858452601f01601f1916604052505b50949350505050565b610764610b25565b6040820151600c60ff9091161061077a57600080fd5b604082015160ff166107a757604051806020016040528061079e84600001516103df565b90529050610400565b604082015160ff16600214156107cc5750604080516020810190915281518152610400565b600360ff16826040015160ff16101580156107f057506040820151600c60ff909116105b1561080d57604051806020016040528061079e84602001516109db565bfe5b60008080610823858563ffffffff610ae516565b60209490940195939450505050565b61083a610b01565b60408051606081018252838152815160008082526020828101909452919283019161087b565b610868610b01565b8152602001906001900390816108605790505b508152600060209091015292915050565b610894610b01565b6040805160608101825283815281516000808252602082810190945291928301916108d5565b6108c2610b01565b8152602001906001900390816108ba5790505b508152600260209091015292915050565b6000806060600060608760ff1660405190808252806020026020018201604052801561092c57816020015b610919610b01565b8152602001906001900390816109115790505b50905060005b8860ff168160ff1610156109865761094a88886105e4565b8451859060ff861690811061095b57fe5b602090810291909101015297509250821561097e57509093508492509050610992565b600101610932565b50600094508593509150505b93509350939050565b6109a3610b01565b6109ad82516105dd565b6109b657600080fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b60006008825111156109ec57600080fd5b60608251604051908082528060200260200182016040528015610a19578160200160208202803883390190505b50905060005b8151811015610a7357610a30610b25565b610a4c858381518110610a3f57fe5b602002602001015161075c565b90508060000151838381518110610a5f57fe5b602090810291909101015250600101610a1f565b508251600360ff160181604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610abc578181015183820152602001610aa4565b505050509050019250505060405160208183030381529060405280519060200120915050919050565b60008160200183511015610af857600080fd5b50016020015190565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a723058205f519cde41045de65b3c250d163984e9b8cab11f031f580ef558872415345c3c64736f6c63430005090032`

// DeployArbValue deploys a new Ethereum contract, binding an instance of ArbValue to it.
func DeployArbValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbValue, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// ArbValue is an auto generated Go binding around an Ethereum contract.
type ArbValue struct {
	ArbValueCaller     // Read-only binding to the contract
	ArbValueTransactor // Write-only binding to the contract
	ArbValueFilterer   // Log filterer for contract events
}

// ArbValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbValueSession struct {
	Contract     *ArbValue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbValueCallerSession struct {
	Contract *ArbValueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbValueTransactorSession struct {
	Contract     *ArbValueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbValueRaw struct {
	Contract *ArbValue // Generic contract binding to access the raw methods on
}

// ArbValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbValueCallerRaw struct {
	Contract *ArbValueCaller // Generic read-only contract binding to access the raw methods on
}

// ArbValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbValueTransactorRaw struct {
	Contract *ArbValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbValue creates a new instance of ArbValue, bound to a specific deployed contract.
func NewArbValue(address common.Address, backend bind.ContractBackend) (*ArbValue, error) {
	contract, err := bindArbValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// NewArbValueCaller creates a new read-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueCaller(address common.Address, caller bind.ContractCaller) (*ArbValueCaller, error) {
	contract, err := bindArbValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueCaller{contract: contract}, nil
}

// NewArbValueTransactor creates a new write-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbValueTransactor, error) {
	contract, err := bindArbValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueTransactor{contract: contract}, nil
}

// NewArbValueFilterer creates a new log filterer instance of ArbValue, bound to a specific deployed contract.
func NewArbValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbValueFilterer, error) {
	contract, err := bindArbValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbValueFilterer{contract: contract}, nil
}

// bindArbValue binds a generic wrapper to an already deployed contract.
func bindArbValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.ArbValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transact(opts, method, params...)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x615c39b0.
//
// Solidity: function deserialize_valid_value_hash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValidValueHash(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "deserialize_valid_value_hash", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x615c39b0.
//
// Solidity: function deserialize_valid_value_hash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x615c39b0.
//
// Solidity: function deserialize_valid_value_hash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x92516ac7.
//
// Solidity: function deserialize_value_hash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValueHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbValue.contract.Call(opts, out, "deserialize_value_hash", data)
	return *ret0, err
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x92516ac7.
//
// Solidity: function deserialize_value_hash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x92516ac7.
//
// Solidity: function deserialize_value_hash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x4d00ef7a.
//
// Solidity: function get_next_valid_value(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCaller) GetNextValidValue(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "get_next_valid_value", data, offset)
	return *ret0, *ret1, err
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x4d00ef7a.
//
// Solidity: function get_next_valid_value(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x4d00ef7a.
//
// Solidity: function get_next_valid_value(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCallerSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointBasicValue(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointBasicValue", opcode, nextCodePoint)
	return *ret0, err
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointImmediateValue(opts *bind.CallOpts, opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointImmediateValue", opcode, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashIntValue(opts *bind.CallOpts, val *big.Int) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashIntValue", val)
	return *ret0, err
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCaller) IsValidTupleSize(opts *bind.CallOpts, size *big.Int) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _ArbValue.contract.Call(opts, out, "isValidTupleSize", size)
	return *ret0, err
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCallerSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
const BytesLibBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820c6694409c11fc2dfd409c9a86b8959cc4bc5705ed732a15d129c94993240c58b64736f6c63430005090032`

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ERC165Bin is the compiled bytecode used for deploying new contracts.
const ERC165Bin = `0x`

// DeployERC165 deploys a new Ethereum contract, binding an instance of ERC165 to it.
func DeployERC165(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC165, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC165Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _ERC165.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x608060405234801561001057600080fd5b506105c1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a457c2d71461016f578063a9059cbb1461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063395093511461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d5610271565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610277565b6100b96004803603604081101561013357600080fd5b506001600160a01b038135169060200135610340565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166103ee565b6100b96004803603604081101561018557600080fd5b506001600160a01b038135169060200135610409565b6100b9600480360360408110156101b157600080fd5b506001600160a01b038135169060200135610452565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610468565b60006001600160a01b03831661020a57600080fd5b3360008181526001602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60025490565b6001600160a01b03831660009081526001602090815260408083203384529091528120546102ab908363ffffffff61049316565b6001600160a01b03851660009081526001602090815260408083203384529091529020556102da8484846104a8565b6001600160a01b0384166000818152600160209081526040808320338085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60006001600160a01b03831661035557600080fd5b3360009081526001602090815260408083206001600160a01b0387168452909152902054610389908363ffffffff61057316565b3360008181526001602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b031660009081526020819052604090205490565b60006001600160a01b03831661041e57600080fd5b3360009081526001602090815260408083206001600160a01b0387168452909152902054610389908363ffffffff61049316565b600061045f3384846104a8565b50600192915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6000828211156104a257600080fd5b50900390565b6001600160a01b0382166104bb57600080fd5b6001600160a01b0383166000908152602081905260409020546104e4908263ffffffff61049316565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610519908263ffffffff61057316565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008282018381101561058557600080fd5b939250505056fea265627a7a72305820e75fdbaf51818524eadaf8fd8ad0052782bd9e60f8636e1ca7807a008b1b0c1464736f6c63430005090032`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721Bin is the compiled bytecode used for deploying new contracts.
const ERC721Bin = `0x608060405234801561001057600080fd5b506100437f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0361007a16565b6100757f80ac58cd000000000000000000000000000000000000000000000000000000006001600160e01b0361007a16565b6100e6565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100a957600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b6108f5806100f56000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636352211e116100665780636352211e146101b157806370a08231146101ce578063a22cb46514610206578063b88d4fde14610234578063e985e9c5146102fa5761009e565b806301ffc9a7146100a3578063081812fc146100de578063095ea7b31461011757806323b872dd1461014557806342842e0e1461017b575b600080fd5b6100ca600480360360208110156100b957600080fd5b50356001600160e01b031916610328565b604080519115158252519081900360200190f35b6100fb600480360360208110156100f457600080fd5b5035610347565b604080516001600160a01b039092168252519081900360200190f35b6101436004803603604081101561012d57600080fd5b506001600160a01b038135169060200135610377565b005b6101436004803603606081101561015b57600080fd5b506001600160a01b03813581169160208101359091169060400135610424565b6101436004803603606081101561019157600080fd5b506001600160a01b03813581169160208101359091169060400135610447565b6100fb600480360360208110156101c757600080fd5b5035610462565b6101f4600480360360208110156101e457600080fd5b50356001600160a01b031661048a565b60408051918252519081900360200190f35b6101436004803603604081101561021c57600080fd5b506001600160a01b03813516906020013515156104bb565b6101436004803603608081101561024a57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561028557600080fd5b82018360208201111561029757600080fd5b803590602001918460018302840111640100000000831117156102b957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061053f945050505050565b6100ca6004803603604081101561031057600080fd5b506001600160a01b0381358116916020013516610565565b6001600160e01b03191660009081526020819052604090205460ff1690565b600061035282610593565b61035b57600080fd5b506000908152600260205260409020546001600160a01b031690565b600061038282610462565b9050806001600160a01b0316836001600160a01b031614156103a357600080fd5b336001600160a01b03821614806103bf57506103bf8133610565565b6103c857600080fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61042e33826105b0565b61043757600080fd5b61044283838361060f565b505050565b6104428383836040518060200160405280600081525061053f565b6000818152600160205260408120546001600160a01b03168061048457600080fd5b92915050565b60006001600160a01b03821661049f57600080fd5b506001600160a01b031660009081526003602052604090205490565b6001600160a01b0382163314156104d157600080fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61054a848484610424565b6105568484848461071c565b61055f57600080fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6000806105bc83610462565b9050806001600160a01b0316846001600160a01b031614806105f75750836001600160a01b03166105ec84610347565b6001600160a01b0316145b8061060757506106078185610565565b949350505050565b826001600160a01b031661062282610462565b6001600160a01b03161461063557600080fd5b6001600160a01b03821661064857600080fd5b6106518161084f565b6001600160a01b03831660009081526003602052604090205461067b90600163ffffffff61088c16565b6001600160a01b0380851660009081526003602052604080822093909355908416815220546106b190600163ffffffff6108a116565b6001600160a01b038084166000818152600360209081526040808320959095558582526001905283812080546001600160a01b031916831790559251849391928716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6000610730846001600160a01b03166108ba565b61073c57506001610607565b604051630a85bd0160e11b815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a169463150b7a029490938c938b938b939260a4019060208501908083838e5b838110156107b657818101518382015260200161079e565b50505050905090810190601f1680156107e35780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561080557600080fd5b505af1158015610819573d6000803e3d6000fd5b505050506040513d602081101561082f57600080fd5b50516001600160e01b031916630a85bd0160e11b14915050949350505050565b6000818152600260205260409020546001600160a01b03161561088957600081815260026020526040902080546001600160a01b03191690555b50565b60008282111561089b57600080fd5b50900390565b6000828201838110156108b357600080fd5b9392505050565b3b15159056fea265627a7a72305820f0e92d3229cb645872e9762ad3e0abb716d2e891f806295440dd6c14885fe85b64736f6c63430005090032`

// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ERC721.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721 *ERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	ret0 := new(common.Address)

	out := ret0
	err := _ERC721.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _ERC721.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	ret0 := new(common.Address)

	out := ret0
	err := _ERC721.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _ERC721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// IChallengeManagerABI is the input ABI used to generate the binding from.
const IChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"vmId\",\"type\":\"bytes32\"},{\"name\":\"players\",\"type\":\"address[2]\"},{\"name\":\"escrows\",\"type\":\"uint128[2]\"},{\"name\":\"challengePeriod\",\"type\":\"uint32\"},{\"name\":\"challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeManagerBin is the compiled bytecode used for deploying new contracts.
const IChallengeManagerBin = `0x`

// DeployIChallengeManager deploys a new Ethereum contract, binding an instance of IChallengeManager to it.
func DeployIChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IChallengeManager, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IChallengeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initiateChallenge", vmId, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerSession) InitiateChallenge(vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, vmId, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) InitiateChallenge(vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, vmId, players, escrows, challengePeriod, challengeRoot)
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165Bin is the compiled bytecode used for deploying new contracts.
const IERC165Bin = `0x`

// DeployIERC165 deploys a new Ethereum contract, binding an instance of IERC165 to it.
func DeployIERC165(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC165, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC165Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
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
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _IERC165.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20Bin is the compiled bytecode used for deploying new contracts.
const IERC20Bin = `0x`

// DeployIERC20 deploys a new Ethereum contract, binding an instance of IERC20 to it.
func DeployIERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
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

// IERC721ABI is the input ABI used to generate the binding from.
const IERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// IERC721Bin is the compiled bytecode used for deploying new contracts.
const IERC721Bin = `0x`

// DeployIERC721 deploys a new Ethereum contract, binding an instance of IERC721 to it.
func DeployIERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC721, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
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
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _IERC721.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	ret0 := new(common.Address)

	out := ret0
	err := _IERC721.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _IERC721.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	ret0 := new(common.Address)

	out := ret0
	err := _IERC721.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _IERC721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
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

// IERC721ReceiverABI is the input ABI used to generate the binding from.
const IERC721ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721ReceiverBin is the compiled bytecode used for deploying new contracts.
const IERC721ReceiverBin = `0x`

// DeployIERC721Receiver deploys a new Ethereum contract, binding an instance of IERC721Receiver to it.
func DeployIERC721Receiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC721Receiver, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC721ReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC721Receiver{IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract}}, nil
}

// IERC721Receiver is an auto generated Go binding around an Ethereum contract.
type IERC721Receiver struct {
	IERC721ReceiverCaller     // Read-only binding to the contract
	IERC721ReceiverTransactor // Write-only binding to the contract
	IERC721ReceiverFilterer   // Log filterer for contract events
}

// IERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721ReceiverSession struct {
	Contract     *IERC721Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721ReceiverCallerSession struct {
	Contract *IERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721ReceiverTransactorSession struct {
	Contract     *IERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721ReceiverRaw struct {
	Contract *IERC721Receiver // Generic contract binding to access the raw methods on
}

// IERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721ReceiverCallerRaw struct {
	Contract *IERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactorRaw struct {
	Contract *IERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Receiver creates a new instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721Receiver(address common.Address, backend bind.ContractBackend) (*IERC721Receiver, error) {
	contract, err := bindIERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Receiver{IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract}}, nil
}

// NewIERC721ReceiverCaller creates a new read-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721ReceiverCaller, error) {
	contract, err := bindIERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverCaller{contract: contract}, nil
}

// NewIERC721ReceiverTransactor creates a new write-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ReceiverTransactor, error) {
	contract, err := bindIERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverTransactor{contract: contract}, nil
}

// NewIERC721ReceiverFilterer creates a new log filterer instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ReceiverFilterer, error) {
	contract, err := bindIERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverFilterer{contract: contract}, nil
}

// bindIERC721Receiver binds a generic wrapper to an already deployed contract.
func bindIERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.IERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// MerkleLibABI is the input ABI used to generate the binding from.
const MerkleLibABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"generateAddressRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"generateRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
const MerkleLibBin = `0x610573610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043d565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b60005b600182511115610421576060600283516001018161031e57fe5b04604051908082528060200260200182016040528015610348578160200160208202803883390190505b50905060005b81518110156104195783518160020260010110156103e15783816002028151811061037557fe5b602002602001015184826002026001018151811061038f57fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d057fe5b602002602001018181525050610411565b8381600202815181106103f057fe5b602002602001015182828151811061040457fe5b6020026020010181815250505b60010161034e565b509150610304565b8160008151811061042e57fe5b60200260200101519050919050565b600080838160205b88518111610530578089015193506020818a51036020018161046357fe5b0491505b60008211801561047a5750600286066001145b801561048857508160020a86115b1561049b57600286046001019550610467565b600286066104e65783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104de57fe5b049550610528565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052157fe5b0460010195505b602001610445565b50509094149594505050505056fea265627a7a7230582060bd4256cc8dc39e31eb5f7eed0ad08ab723bf0337043846f63b965afb197ae664736f6c63430005090032`

// DeployMerkleLib deploys a new Ethereum contract, binding an instance of MerkleLib to it.
func DeployMerkleLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleLib, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// MerkleLib is an auto generated Go binding around an Ethereum contract.
type MerkleLib struct {
	MerkleLibCaller     // Read-only binding to the contract
	MerkleLibTransactor // Write-only binding to the contract
	MerkleLibFilterer   // Log filterer for contract events
}

// MerkleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleLibSession struct {
	Contract     *MerkleLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleLibCallerSession struct {
	Contract *MerkleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MerkleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleLibTransactorSession struct {
	Contract     *MerkleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MerkleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleLibRaw struct {
	Contract *MerkleLib // Generic contract binding to access the raw methods on
}

// MerkleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleLibCallerRaw struct {
	Contract *MerkleLibCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleLibTransactorRaw struct {
	Contract *MerkleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleLib creates a new instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLib(address common.Address, backend bind.ContractBackend) (*MerkleLib, error) {
	contract, err := bindMerkleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// NewMerkleLibCaller creates a new read-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibCaller(address common.Address, caller bind.ContractCaller) (*MerkleLibCaller, error) {
	contract, err := bindMerkleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibCaller{contract: contract}, nil
}

// NewMerkleLibTransactor creates a new write-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleLibTransactor, error) {
	contract, err := bindMerkleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibTransactor{contract: contract}, nil
}

// NewMerkleLibFilterer creates a new log filterer instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleLibFilterer, error) {
	contract, err := bindMerkleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleLibFilterer{contract: contract}, nil
}

// bindMerkleLib binds a generic wrapper to an already deployed contract.
func bindMerkleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.MerkleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transact(opts, method, params...)
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibCaller) GenerateAddressRoot(opts *bind.CallOpts, _addresses []common.Address) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "generateAddressRoot", _addresses)
	return *ret0, err
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibSession) GenerateAddressRoot(_addresses []common.Address) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateAddressRoot(&_MerkleLib.CallOpts, _addresses)
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibCallerSession) GenerateAddressRoot(_addresses []common.Address) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateAddressRoot(&_MerkleLib.CallOpts, _addresses)
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibCaller) GenerateRoot(opts *bind.CallOpts, _hashes [][32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "generateRoot", _hashes)
	return *ret0, err
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibSession) GenerateRoot(_hashes [][32]byte) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateRoot(&_MerkleLib.CallOpts, _hashes)
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibCallerSession) GenerateRoot(_hashes [][32]byte) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateRoot(&_MerkleLib.CallOpts, _hashes)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibCaller) VerifyProof(opts *bind.CallOpts, proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "verifyProof", proof, root, hash, index)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibSession) VerifyProof(proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	return _MerkleLib.Contract.VerifyProof(&_MerkleLib.CallOpts, proof, root, hash, index)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibCallerSession) VerifyProof(proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	return _MerkleLib.Contract.VerifyProof(&_MerkleLib.CallOpts, proof, root, hash, index)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)

	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7230582061d33a32aaf0206e9cb2c3252951d51904539b6b2ad671f2da507142f8c56fcd64736f6c63430005090032`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// VMTrackerABI is the input ABI used to generate the binding from.
const VMTrackerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[5]\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageDataHash\",\"type\":\"bytes32[]\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_msgAmount\",\"type\":\"uint256[]\"},{\"name\":\"_msgDestination\",\"type\":\"bytes32[]\"}],\"name\":\"disputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_challengeManager\",\"type\":\"address\"}],\"name\":\"addChallengeManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"withinTimeBounds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_newInbox\",\"type\":\"bytes32\"},{\"name\":\"_logsHash\",\"type\":\"bytes32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"unanimousAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_unanRest\",\"type\":\"bytes32\"},{\"name\":\"_sequenceNum\",\"type\":\"uint64\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"proposeUnanimousAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_logsHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"confirmAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[3]\"},{\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"name\":\"_challengeManagerNum\",\"type\":\"uint16\"},{\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"name\":\"_escrowCurrency\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"createVm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_logsHash\",\"type\":\"bytes32\"},{\"name\":\"_newInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"ConfirmUnanimousAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"}],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_balanceTrackerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"tokenType\",\"type\":\"bytes21\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"_escrowCurrency\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_challengeManagerNum\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"VMCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"unanHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"ProposedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"ConfirmedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"unanHash\",\"type\":\"bytes32\"}],\"name\":\"FinalUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"fields\",\"type\":\"bytes32[3]\"},{\"indexed\":false,\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"indexed\":false,\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"lastMessageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"logsHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"DisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"newState\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// VMTrackerBin is the compiled bytecode used for deploying new contracts.
const VMTrackerBin = `0x608060405234801561001057600080fd5b50604051620054fa380380620054fa8339818101604052602081101561003557600080fd5b5051600080546001600160a01b03191633178082556040516001600160a01b039190911691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3600280546001600160a01b039092166001600160a01b03199092169190911790556000805260046020527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec805460ff1916600117905561541480620000e66000396000f3fe6080604052600436106100fe5760003560e01c806389d9957e11610095578063d0b41b6c11610064578063d0b41b6c14611035578063f2fde38b14611307578063fa3eebe61461133a578063fd0961b6146113e5578063ffd134c91461140f576100fe565b806389d9957e14610be45780638da5cb5b14610ec35780638dbc44f614610ef45780638f32d59b14611020576100fe565b806363d84637116100d157806363d846371461059a5780636c94603a146105ce578063715018a6146109515780637d5d7b7714610966576100fe565b80631c0fc42c146101035780632e6bf2e6146104295780633cd8a3221461045c57806342c0787e1461052e575b600080fd5b34801561010f57600080fd5b5061042760048036036101a081101561012757600080fd5b810190808060a00190600580602002604051908101604052809291908260056020028082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b8111156101ae57600080fd5b8201836020820111156101c057600080fd5b803590602001918460208302840111600160201b831117156101e157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561023057600080fd5b82018360208201111561024257600080fd5b803590602001918460208302840111600160201b8311171561026357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102b257600080fd5b8201836020820111156102c457600080fd5b803590602001918460208302840111600160201b831117156102e557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561033457600080fd5b82018360208201111561034657600080fd5b803590602001918460208302840111600160201b8311171561036757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103b657600080fd5b8201836020820111156103c857600080fd5b803590602001918460208302840111600160201b831117156103e957600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061143f945050505050565b005b34801561043557600080fd5b506104276004803603602081101561044c57600080fd5b50356001600160a01b03166114fe565b34801561046857600080fd5b506104276004803603608081101561047f57600080fd5b8135916affffffffffffffffffffff196020820135169160408201359190810190608081016060820135600160201b8111156104ba57600080fd5b8201836020820111156104cc57600080fd5b803590602001918460018302840111600160201b831117156104ed57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611560945050505050565b34801561053a57600080fd5b506105866004803603604081101561055157600080fd5b604080518082018252918301929181830191839060029083908390808284376000920191909152509194506115809350505050565b604080519115158252519081900360200190f35b3480156105a657600080fd5b50610427600480360360a08110156105bd57600080fd5b5080359060208101906060016115b0565b3480156105da57600080fd5b5061042760048036036101808110156105f257600080fd5b6040805180820182528335936020810135938382013593606083013593918301929160c08301916080840190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561065057600080fd5b82018360208201111561066257600080fd5b803590602001918460208302840111600160201b8311171561068357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106d257600080fd5b8201836020820111156106e457600080fd5b803590602001918460018302840111600160201b8311171561070557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561075757600080fd5b82018360208201111561076957600080fd5b803590602001918460208302840111600160201b8311171561078a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107d957600080fd5b8201836020820111156107eb57600080fd5b803590602001918460208302840111600160201b8311171561080c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561085b57600080fd5b82018360208201111561086d57600080fd5b803590602001918460208302840111600160201b8311171561088e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156108dd57600080fd5b8201836020820111156108ef57600080fd5b803590602001918460018302840111600160201b8311171561091057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506116db945050505050565b34801561095d57600080fd5b5061042761173c565b34801561097257600080fd5b50610427600480360361012081101561098a57600080fd5b6040805180820182528335936020810135936001600160401b038483013516939082019260a08301916060840190600290839083908082843760009201919091525091949392602081019250359050600160201b8111156109ea57600080fd5b8201836020820111156109fc57600080fd5b803590602001918460208302840111600160201b83111715610a1d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a6c57600080fd5b820183602082011115610a7e57600080fd5b803590602001918460208302840111600160201b83111715610a9f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610aee57600080fd5b820183602082011115610b0057600080fd5b803590602001918460208302840111600160201b83111715610b2157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b7057600080fd5b820183602082011115610b8257600080fd5b803590602001918460018302840111600160201b83111715610ba357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611797945050505050565b348015610bf057600080fd5b506104276004803603610140811015610c0857600080fd5b81359160208101359160408201359160608101359163ffffffff608083013516919081019060c0810160a0820135600160201b811115610c4757600080fd5b820183602082011115610c5957600080fd5b803590602001918460208302840111600160201b83111715610c7a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610cc957600080fd5b820183602082011115610cdb57600080fd5b803590602001918460018302840111600160201b83111715610cfc57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610d4e57600080fd5b820183602082011115610d6057600080fd5b803590602001918460208302840111600160201b83111715610d8157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610dd057600080fd5b820183602082011115610de257600080fd5b803590602001918460208302840111600160201b83111715610e0357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e5257600080fd5b820183602082011115610e6457600080fd5b803590602001918460208302840111600160201b83111715610e8557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611fae945050505050565b348015610ecf57600080fd5b50610ed861200e565b604080516001600160a01b039092168252519081900360200190f35b348015610f0057600080fd5b506104276004803603610140811015610f1857600080fd5b8101908080606001906003806020026040519081016040528092919082600360200280828437600092019190915250919463ffffffff843581169560208601359091169461ffff60408201351694506001600160801b0360608201351693506001600160a01b03608082013581169360a083013590911692909160e081019060c00135600160201b811115610fac57600080fd5b820183602082011115610fbe57600080fd5b803590602001918460018302840111600160201b83111715610fdf57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061201d945050505050565b34801561102c57600080fd5b5061058661292b565b34801561104157600080fd5b50610427600480360361012081101561105957600080fd5b81359160208101359160408201359160608101359181019060a081016080820135600160201b81111561108b57600080fd5b82018360208201111561109d57600080fd5b803590602001918460208302840111600160201b831117156110be57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561110d57600080fd5b82018360208201111561111f57600080fd5b803590602001918460018302840111600160201b8311171561114057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561119257600080fd5b8201836020820111156111a457600080fd5b803590602001918460208302840111600160201b831117156111c557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561121457600080fd5b82018360208201111561122657600080fd5b803590602001918460208302840111600160201b8311171561124757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561129657600080fd5b8201836020820111156112a857600080fd5b803590602001918460208302840111600160201b831117156112c957600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061293c945050505050565b34801561131357600080fd5b506104276004803603602081101561132a57600080fd5b50356001600160a01b0316612b7c565b6104276004803603604081101561135057600080fd5b81359190810190604081016020820135600160201b81111561137157600080fd5b82018360208201111561138357600080fd5b803590602001918460018302840111600160201b831117156113a457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612b99945050505050565b3480156113f157600080fd5b506104276004803603602081101561140857600080fd5b5035612c27565b34801561141b57600080fd5b506104276004803603604081101561143257600080fd5b5080359060200135612c9f565b6114f46040518061018001604052808a60006005811061145b57fe5b602002015181526020018a60016005811061147257fe5b602002015181526020018a60026005811061148957fe5b602002015181526020018a6003600581106114a057fe5b602002015181526020018a6004600581106114b757fe5b602002015181526020018963ffffffff16815260200188815260200187815260200186815260200185815260200184815260200183815250613029565b5050505050505050565b61150661292b565b61150f57600080fd5b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0392909216919091179055565b61157a8484846001600160601b03193360601b1685613c1d565b50505050565b80516000906001600160401b031643108015906115aa575060208201516001600160401b03164311155b92915050565b6000838152600360205260409020600b810154600180549091600160501b900461ffff169081106115dd57fe5b6000918252602090912001546001600160a01b031633146115fd57600080fd5b600b810154600160681b900460ff1661161557600080fd5b600b8101805460ff60681b191690556116786001600160801b03833516600c8301600086815b60200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002054613d4e90919063ffffffff16565b83356001600160a01b03166000908152600c8301602081815260408320939093556116b2928501356001600160801b03169186600161163b565b6001600160a01b03602094850135166000908152600c9290920190935260409020919091555050565b61172f6040518061016001604052808d81526020018c81526020018a81526020018b815260200189815260200188815260200187815260200186815260200185815260200184815260200183815250613d67565b5050505050505050505050565b61174461292b565b61174d57600080fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008881526003602052604090206117ae86611580565b6117e95760405162461bcd60e51b81526004018080602001828103825260248152602001806152d86024913960400191505060405180910390fd5b80546117f457600080fd5b600089898989856000015486600201548b8b8b604051602001808a8152602001898152602001886001600160401b03166001600160401b031660c01b815260080187600260200280838360005b83811015611859578181015183820152602001611841565b50505050905001868152602001858152602001848051906020019060200280838360005b8381101561189557818101518382015260200161187d565b50505050905001838051906020019060200280838360005b838110156118c55781810151838201526020016118ad565b50505050905001828051906020019060200280838360005b838110156118f55781810151838201526020016118dd565b505050509050019950505050505050505050604051602081830303815290604052805190602001209050816004015473__$a50780cb42d41d2927e39f529dc62d6697$__636a2dda6773__$6b4cc75dad3e0abd6ad83b3d907747c608$__63f0c8e96985886040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156119ac578181015183820152602001611994565b50505050905090810190601f1680156119d95780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b1580156119f757600080fd5b505af4158015611a0b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015611a3457600080fd5b810190808051600160201b811115611a4b57600080fd5b82016020810184811115611a5e57600080fd5b81518560208202830111600160201b82111715611a7a57600080fd5b50506040516001600160e01b031960e087901b1681526020600482018181528351602484015283519396509450849350604490910191818601910280838360005b83811015611ad3578181015183820152602001611abb565b505050509050019250505060206040518083038186803b158015611af657600080fd5b505af4158015611b0a573d6000803e3d6000fd5b505050506040513d6020811015611b2057600080fd5b505114611b74576040805162461bcd60e51b815260206004820181905260248201527f56616c696461746f72207369676e61747572657320646f6e2774206d61746368604482015290519081900360640190fd5b6002600b830154600160601b900460ff166002811115611b9057fe5b1415611bba57600a8201546001600160401b03600160c01b909104811690891611611bba57600080fd5b600254604051630f89fbff60e01b81526060600482019081528851606483015288516001600160a01b039093169263c2465106928e928b9273__$6b4cc75dad3e0abd6ad83b3d907747c608$__92630f89fbff9285928e928e9291829160248101916044820191608401906020808a01910280838360005b83811015611c4a578181015183820152602001611c32565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015611c89578181015183820152602001611c71565b50505050905001848103825285818151815260200191508051906020019060200280838360005b83811015611cc8578181015183820152602001611cb0565b50505050905001965050505050505060006040518083038186803b158015611cef57600080fd5b505af4158015611d03573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015611d2c57600080fd5b810190808051600160201b811115611d4357600080fd5b82016020810184811115611d5657600080fd5b81518560208202830111600160201b82111715611d7257600080fd5b50509291905050506040518463ffffffff1660e01b8152600401808481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015611dd3578181015183820152602001611dbb565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611e12578181015183820152602001611dfa565b5050505090500195505050505050600060405180830381600087803b158015611e3a57600080fd5b505af1158015611e4e573d6000803e3d6000fd5b50505050611e5b82614274565b611e6482614336565b600b8201805460ff60601b1916600160611b179055600a820180546001600160c01b0316600160c01b6001600160401b038b160217905560405186518791879187918d916020918201918291818801910280838360005b83811015611ed3578181015183820152602001611ebb565b50505050905001848051906020019060200280838360005b83811015611f03578181015183820152602001611eeb565b50505050905001838051906020019060200280838360005b83811015611f33578181015183820152602001611f1b565b5050505091909101928352505060408051808303815260208084018084528251929091019190912060018901558690526001600160401b038d1681830152518e94507fa1d1249870ac5b939cf041991350480912abf0af99cf0337b13e440e1e150b969350908190036060019150a250505050505050505050565b6120026040518061014001604052808c81526020018b81526020018a81526020018981526020018863ffffffff16815260200187815260200186815260200185815260200184815260200183815250614374565b50505050505050505050565b6000546001600160a01b031690565b62010000604182518161202c57fe5b0410612075576040805162461bcd60e51b8152602060048201526013602482015272546f6f206d616e792076616c696461746f727360681b604482015290519081900360640190fd5b87516001600160601b0319811614156120c4576040805162461bcd60e51b815260206004820152600c60248201526b125b9d985b1a59081d9b525960a21b604482015290519081900360640190fd5b60015461ffff86161061211e576040805162461bcd60e51b815260206004820152601d60248201527f496e76616c6964206368616c6c656e6765206d616e61676572206e756d000000604482015290519081900360640190fd5b6000846001600160801b03161161213457600080fd5b6001600160a01b03831660009081526004602052604090205460ff1661218b5760405162461bcd60e51b81526004018080602001828103825260298152602001806152696029913960400191505060405180910390fd5b604080890151815163f0c8e96960e01b8152600481018281526024820193845284516044830152845160609473__$6b4cc75dad3e0abd6ad83b3d907747c608$__9463f0c8e969949093889390929160640190602085019080838360005b838110156122015781810151838201526020016121e9565b50505050905090810190601f16801561222e5780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b15801561224c57600080fd5b505af4158015612260573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561228957600080fd5b810190808051600160201b8111156122a057600080fd5b820160208101848111156122b357600080fd5b81518560208202830111600160201b821117156122cf57600080fd5b509094508b93508892508791508a90508c600160200201518a8887604051602001808963ffffffff1663ffffffff1660e01b8152600401886001600160801b03166001600160801b031660801b8152601001876001600160a01b03166001600160a01b031660601b81526014018663ffffffff1663ffffffff1660e01b81526004018581526020018461ffff1661ffff1660f01b8152600201836001600160a01b03166001600160a01b031660601b8152601401828051906020019060200280838360005b838110156123ac578181015183820152602001612394565b505050509050019850505050505050505060405160208183030381529060405280519060200120896002600381106123e057fe5b60200201511461242f576040805162461bcd60e51b815260206004820152601560248201527410dc99585d194819185d18481a5b98dbdc9c9958dd605a1b604482015290519081900360640190fd5b60005b81518110156124fd5760025482516001600160a01b0390911690630e227be39084908490811061245e57fe5b602002602001015160601b6bffffffffffffffffffffffff191687896040518463ffffffff1660e01b815260040180848152602001836001600160a01b03166001600160a01b03168152602001826001600160801b031681526020019350505050600060405180830381600087803b1580156124d957600080fd5b505af11580156124ed573d6000803e3d6000fd5b5050600190920191506124329050565b50885160009081526003602052604090208960016020020151816000018190555073__$0d86abb4a722a612872fb80f4c7e7e95bd$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b15801561256257600080fd5b505af4158015612576573d6000803e3d6000fd5b505050506040513d602081101561258c57600080fd5b505160028201556040805163364df27760e01b8152905173__$0d86abb4a722a612872fb80f4c7e7e95bd$__9163364df277916004808301926020929190829003018186803b1580156125de57600080fd5b505af41580156125f2573d6000803e3d6000fd5b505050506040513d602081101561260857600080fd5b50516003820155600b8101805461ffff60501b1916600160501b61ffff8a16021760ff60601b19169055600060018201819055604051636a2dda6760e01b815260206004820181815285516024840152855173__$a50780cb42d41d2927e39f529dc62d6697$__94636a2dda67948894849360449092019286820192909102908190849084905b838110156126a757818101518382015260200161268f565b505050509050019250505060206040518083038186803b1580156126ca57600080fd5b505af41580156126de573d6000803e3d6000fd5b505050506040513d60208110156126f457600080fd5b505160048201558151600b82018054600a840180546fffffffffffffffffffffffffffffffff19166001600160801b038b161790556009840180546001600160a01b03199081166001600160a01b038b81169190911790925560078601805490911691891691909117905569ffff000000000000000019166801000000000000000061ffff909316929092029190911763ffffffff191663ffffffff8b81169190911767ffffffff000000001916600160201b918b169190910217905560005b825181101561280557866001600160801b031682600c0160008584815181106127d957fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020556001016127b4565b507fe251d26fa5921a42fd8dfa2415cfa1f48aae04e9c147d0edb7fd63d5bdc9e6f48987878b8e600060200201518f600160200201518d8b8a604051808a63ffffffff1663ffffffff168152602001896001600160801b03166001600160801b03168152602001886001600160a01b03166001600160a01b031681526020018763ffffffff1663ffffffff1681526020018681526020018581526020018461ffff1661ffff168152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156129045781810151838201526020016128ec565b505050509050019a505050505050505050505060405180910390a150505050505050505050565b6000546001600160a01b0316331490565b60008981526003602052604090206002600b820154600160601b900460ff16600281111561296657fe5b1461297057600080fd5b600a810154600160801b90046001600160401b0316431161299057600080fd5b858484898c8c8a886040516020018086815260200185815260200184815260200183805190602001908083835b602083106129dc5780518252601f1990920191602091820191016129bd565b51815160209384036101000a60001901801990921691161790528551919093019285810192500280838360005b83811015612a21578181015183820152602001612a09565b50505050905001955050505050506040516020818303038152906040528051906020012060405160200180858051906020019060200280838360005b83811015612a75578181015183820152602001612a5d565b50505050905001848051906020019060200280838360005b83811015612aa5578181015183820152602001612a8d565b50505050905001838051906020019060200280838360005b83811015612ad5578181015183820152602001612abd565b5050505090500182815260200194505050505060405160208183030381529060405280519060200120816001015414612b0d57600080fd5b60028101879055612b238a8a8888888888614a31565b600a810154604080516001600160401b03600160c01b909304929092168252518b917fd23ef3a25058f2acee4ca7aee5cc78a587b9fc76d2875098af1ffdfc2bcc74c9919081900360200190a250505050505050505050565b612b8461292b565b612b8d57600080fd5b612b9681614db1565b50565b60025460408051636d31ebdb60e11b81526004810185905290516001600160a01b039092169163da63d7b6913491602480830192600092919082900301818588803b158015612be757600080fd5b505af1158015612bfb573d6000803e3d6000fd5b5050505050612c2382600060581b343360601b6bffffffffffffffffffffffff191685614e1f565b5050565b600081815260036020526040902060078101546001600160a01b03163314612c96576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6c79206f776e65722063616e2073687574646f776e2074686520564d0000604482015290519081900360640190fd5b612c238261513c565b600082815260036020526040902060088101546001600160a01b0316331415612cf95760405162461bcd60e51b81526004018080602001828103825260218152602001806152fc6021913960400191505060405180910390fd5b600a810154600160801b90046001600160401b0316431115612d4c5760405162461bcd60e51b815260040180806020018281038252602681526020018061538d6026913960400191505060405180910390fd5b6001600b820154600160601b900460ff166002811115612d6857fe5b14612da45760405162461bcd60e51b815260040180806020018281038252602f81526020018061523a602f913960400191505060405180910390fd5b336000908152600c82016020526040902054600a8201546001600160801b03161115612e015760405162461bcd60e51b81526004018080602001828103825260278152602001806151da6027913960400191505060405180910390fd5b80600101548214612e435760405162461bcd60e51b81526004018080602001828103825260398152602001806152016039913960400191505060405180910390fd5b600a810154336000908152600c83016020526040902054612e72916001600160801b031663ffffffff6151c416565b336000908152600c83016020526040812091909155600180830191909155600b82018054600160681b61ffff60601b1990911617908190558154600160501b90910461ffff16908110612ec157fe5b600091825260208083209091015460408051808201825260088601546001600160a01b039081168252338286015282518084018452600a8801546001600160801b031680825295810195909552600b8701548351632b50d42b60e01b8152600481018b81529290951696632b50d42b968b969495909463ffffffff909316938b93909260249091019187918190849084905b83811015612f6b578181015183820152602001612f53565b5050505090500184600260200280838360005b83811015612f96578181015183820152602001612f7e565b505050509050018363ffffffff1663ffffffff16815260200182815260200195505050505050600060405180830381600087803b158015612fd657600080fd5b505af1158015612fea573d6000803e3d6000fd5b50506040805133815290518693507fb8d6a904566965067cb7e3e946d2a494a4b1f955ffec3e72e1e6ebe4e8e52c7292509081900360200190a2505050565b8051600090815260036020526040812090600b820154600160601b900460ff16600281111561305457fe5b146130905760405162461bcd60e51b815260040180806020018281038252602d8152602001806153b3602d913960400191505060405180910390fd5b8054158015906130a257508054600114155b6130ab57600080fd5b600b810154600160681b900460ff16156130c457600080fd5b336000908152600c82016020526040902054600a8201546001600160801b031611156131215760405162461bcd60e51b815260040180806020018281038252602781526020018061533f6027913960400191505060405180910390fd5b600b81015460a083015163ffffffff600160201b90920482169116111561318f576040805162461bcd60e51b815260206004820152601f60248201527f547269656420746f206578656375746520746f6f206d616e7920737465707300604482015290519081900360640190fd5b61319c8260c00151611580565b6131d75760405162461bcd60e51b81526004018080602001828103825260248152602001806152d86024913960400191505060405180910390fd5b805460208301511461321a5760405162461bcd60e51b81526004018080602001828103825260278152602001806153666027913960400191505060405180910390fd5b8060020154826040015114806132bd575073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63f11fcc26826002015483600301546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561328a57600080fd5b505af415801561329e573d6000803e3d6000fd5b505050506040513d60208110156132b457600080fd5b50516040830151145b6132f85760405162461bcd60e51b815260040180806020018281038252602281526020018061531d6022913960400191505060405180910390fd5b606073__$6b4cc75dad3e0abd6ad83b3d907747c608$__630f89fbff8460e001518561012001518661014001516040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b8381101561337c578181015183820152602001613364565b50505050905001848103835286818151815260200191508051906020019060200280838360005b838110156133bb5781810151838201526020016133a3565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156133fa5781810151838201526020016133e2565b50505050905001965050505050505060006040518083038186803b15801561342157600080fd5b505af4158015613435573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561345e57600080fd5b810190808051600160201b81111561347557600080fd5b8201602081018481111561348857600080fd5b81518560208202830111600160201b821117156134a457600080fd5b5050600254875160e0890151604051636123288360e11b815260048101838152606060248301908152835160648401528351969a506001600160a01b03909516985063c24651069750929550909388939160448101916084909101906020808801910280838360005b8381101561352557818101518382015260200161350d565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561356457818101518382015260200161354c565b5050505090500195505050505050600060405180830381600087803b15801561358c57600080fd5b505af11580156135a0573d6000803e3d6000fd5b505050506135ad82614336565b600073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63b32774958560e001518661010001518761012001518861014001518961016001516040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b8381101561364557818101518382015260200161362d565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b8381101561368457818101518382015260200161366c565b50505050905001868103845289818151815260200191508051906020019060200280838360005b838110156136c35781810151838201526020016136ab565b50505050905001868103835288818151815260200191508051906020019060200280838360005b838110156137025781810151838201526020016136ea565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015613741578181015183820152602001613729565b505050509050019a505050505050505050505060206040518083038186803b15801561376c57600080fd5b505af4158015613780573d6000803e3d6000fd5b505050506040513d602081101561379657600080fd5b5051602085015160c086015160408088015160e089015182516307c50ab360e31b81526004810186815296975073__$6b4cc75dad3e0abd6ad83b3d907747c608$__96633e2855989695948a9260240190869080838360005b838110156138075781810151838201526020016137ef565b505050509050018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561385457818101518382015260200161383c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561389357818101518382015260200161387b565b5050505090500197505050505050505060206040518083038186803b1580156138bb57600080fd5b505af41580156138cf573d6000803e3d6000fd5b505050506040513d60208110156138e557600080fd5b5051606085015160a08601516080870151604051632090372160e01b81526004810184815263ffffffff84166024830152600060448301819052606483018890526084830181905260a4830184905260e060c48401908152895160e4850152895173__$6b4cc75dad3e0abd6ad83b3d907747c608$__976320903721979096909593948b94869492938e93916101040190602085810191028083838a5b8381101561399a578181015183820152602001613982565b505050509050019850505050505050505060206040518083038186803b1580156139c357600080fd5b505af41580156139d7573d6000803e3d6000fd5b505050506040513d60208110156139ed57600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920181528151918301919091206001860155600a850154336000908152600c8701909352912054613a4e916001600160801b031663ffffffff6151c416565b336000818152600c860160205260409020919091556008840180546001600160a01b0319169091179055600b830180546001919060ff60601b1916600160601b83021790555083600001517fe3457847513bb99aa86c769d9800b4a4c1f0758755234d78d40c406de93935d1604051806060016040528087602001518152602001876040015181526020018760600151815250338760c001518860e001518960a00151878b608001518a6040518089600360200280838360005b83811015613b20578181015183820152602001613b08565b505050506001600160a01b038b1692019182525060200187604080838360005b83811015613b58578181015183820152602001613b40565b50505050905001806020018663ffffffff1663ffffffff16815260200185815260200184815260200180602001838103835288818151815260200191508051906020019060200280838360005b83811015613bbd578181015183820152602001613ba5565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015613bfc578181015183820152602001613be4565b505050509050019a505050505050505050505060405180910390a250505050565b600160f81b6001600160f81b0319601486901a60f81b161415613cbc576002546040805163cdf25dc160e01b81526004810185905260248101889052606087901c60448201526064810186905290516001600160a01b039092169163cdf25dc19160848082019260009290919082900301818387803b158015613c9f57600080fd5b505af1158015613cb3573d6000803e3d6000fd5b50505050613d3a565b60025460408051631240117b60e01b81526004810185905260248101889052606087901c60448201526064810186905290516001600160a01b0390921691631240117b9160848082019260009290919082900301818387803b158015613d2157600080fd5b505af1158015613d35573d6000803e3d6000fd5b505050505b613d478585858585614e1f565b5050505050565b600082820183811015613d6057600080fd5b9392505050565b805160009081526003602052604090208054613d8257600080fd5b613d8f8260800151611580565b613dca5760405162461bcd60e51b81526004018080602001828103825260248152602001806152d86024913960400191505060405180910390fd5b600082600001518260000154836002015485602001518660400151876060015188608001518960a001518a60c001518b60e001518c61010001518d6101200151604051602001808d81526020018c81526020018b81526020018a815260200189815260200188815260200187600260200280838360005b83811015613e59578181015183820152602001613e41565b50505050905001868051906020019060200280838360005b83811015613e89578181015183820152602001613e71565b5050505090500185805190602001908083835b60208310613ebb5780518252601f199092019160209182019101613e9c565b51815160209384036101000a60001901801990921691161790528751919093019287810192500280838360005b83811015613f00578181015183820152602001613ee8565b50505050905001838051906020019060200280838360005b83811015613f30578181015183820152602001613f18565b50505050905001828051906020019060200280838360005b83811015613f60578181015183820152602001613f48565b505050509050019c50505050505050505050505050604051602081830303815290604052805190602001209050816004015473__$a50780cb42d41d2927e39f529dc62d6697$__636a2dda6773__$6b4cc75dad3e0abd6ad83b3d907747c608$__63f0c8e969858861014001516040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561401f578181015183820152602001614007565b50505050905090810190601f16801561404c5780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b15801561406a57600080fd5b505af415801561407e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156140a757600080fd5b810190808051600160201b8111156140be57600080fd5b820160208101848111156140d157600080fd5b81518560208202830111600160201b821117156140ed57600080fd5b50506040516001600160e01b031960e087901b1681526020600482018181528351602484015283519396509450849350604490910191818601910280838360005b8381101561414657818101518382015260200161412e565b505050509050019250505060206040518083038186803b15801561416957600080fd5b505af415801561417d573d6000803e3d6000fd5b505050506040513d602081101561419357600080fd5b5051146141e7576040805162461bcd60e51b815260206004820181905260248201527f56616c696461746f72207369676e61747572657320646f6e2774206d61746368604482015290519081900360640190fd5b6141f082614274565b600b8201805460ff60601b19169055606083015160028301558251602084015160a085015160c086015160e087015161010088015161012089015161423a96959493929190614a31565b82516040805183815290517fd13341838a6e5e2972e2b8638f8d2652f77945f3755d3a2829ec5076b51470679181900360200190a2505050565b6000600b820154600160601b900460ff16600281111561429057fe5b146142b657600a810154600160801b90046001600160401b03164311156142b657600080fd5b6001600b820154600160601b900460ff1660028111156142d257fe5b1415612b9657600a81015460088201546001600160a01b03166000908152600c83016020526040902054614314916001600160801b031663ffffffff613d4e16565b60088201546001600160a01b03166000908152600c8301602052604090205550565b600b810154600a909101805467ffffffffffffffff60801b1916600160801b63ffffffff90931643016001600160401b031692909202919091179055565b805160009081526003602052604090206001600b820154600160601b900460ff1660028111156143a057fe5b146143dc5760405162461bcd60e51b81526004018080602001828103825260228152602001806152b66022913960400191505060405180910390fd5b600a810154600160801b90046001600160401b0316431161442e5760405162461bcd60e51b81526004018080602001828103825260248152602001806152926024913960400191505060405180910390fd5b8060010154826020015173__$6b4cc75dad3e0abd6ad83b3d907747c608$__632090372185604001518660800151600073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63252001608a60a001518b60c001518c60e001518d61010001518e61012001516040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b838110156144f25781810151838201526020016144da565b5050505090500186810385528a818151815260200191508051906020019080838360005b8381101561452e578181015183820152602001614516565b50505050905090810190601f16801561455b5780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015614590578181015183820152602001614578565b50505050905001868103835288818151815260200191508051906020019060200280838360005b838110156145cf5781810151838201526020016145b7565b50505050905001868103825287818151815260200191508051906020019060200280838360005b8381101561460e5781810151838201526020016145f6565b505050509050019a505050505050505050505060206040518083038186803b15801561463957600080fd5b505af415801561464d573d6000803e3d6000fd5b505050506040513d602081101561466357600080fd5b810190808051906020019092919050505060008a6060015173__$6b4cc75dad3e0abd6ad83b3d907747c608$__630f89fbff8d60a001518e60e001518f61010001516040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b838110156146fc5781810151838201526020016146e4565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561473b578181015183820152602001614723565b50505050905001848103825285818151815260200191508051906020019060200280838360005b8381101561477a578181015183820152602001614762565b50505050905001965050505050505060006040518083038186803b1580156147a157600080fd5b505af41580156147b5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156147de57600080fd5b810190808051600160201b8111156147f557600080fd5b8201602081018481111561480857600080fd5b81518560208202830111600160201b8211171561482457600080fd5b505060405160e08c811b6001600160e01b0319168252600482018c815263ffffffff8c166024840152604483018b9052606483018a90526084830189905260a4830188905260c48301918252835160e484015283519396509450925061010401906020808601910280838360005b838110156148aa578181015183820152602001614892565b505050509050019850505050505050505060206040518083038186803b1580156148d357600080fd5b505af41580156148e7573d6000803e3d6000fd5b505050506040513d60208110156148fd57600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920190528051910120146149645760405162461bcd60e51b81526004018080602001828103825260398152602001806152016039913960400191505060405180910390fd5b600a81015460088201546001600160a01b03166000908152600c830160205260409020546149a0916001600160801b031663ffffffff613d4e16565b60088201546001600160a01b03166000908152600c83016020526040908190209190915582519083015160a084015160c085015160e08601516101008701516101208801516149f496959493929190614a31565b8151604080840151815190815290517f8bf82ba359336cefa1c284ed9d3864f99aa1cd41dddf9330afd656bcb9dd7a7b9181900360200190a25050565b6000878152600360205260408120606090825b8551811015614c025773__$0d86abb4a722a612872fb80f4c7e7e95bd$__634d00ef7a89866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015614abb578181015183820152602001614aa3565b50505050905090810190601f168015614ae85780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b158015614b0657600080fd5b505af4158015614b1a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015614b4357600080fd5b815160208301805191939283019291600160201b811115614b6357600080fd5b82016020810184811115614b7657600080fd5b8151600160201b811182820187101715614b8f57600080fd5b505089519498509650614bfa9389935085925082109050614bac57fe5b60200260200101518a898481518110614bc157fe5b602002602001015161ffff1681518110614bd757fe5b6020026020010151888481518110614beb57fe5b60200260200101518e87613c1d565b600101614a44565b50888155600b8101805460ff60601b191690556040805163364df27760e01b8152905173__$0d86abb4a722a612872fb80f4c7e7e95bd$__9163364df277916004808301926020929190829003018186803b158015614c6057600080fd5b505af4158015614c74573d6000803e3d6000fd5b505050506040513d6020811015614c8a57600080fd5b5051600382015414614da35773__$6b4cc75dad3e0abd6ad83b3d907747c608$__63f11fcc26826002015483600301546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b158015614cf557600080fd5b505af4158015614d09573d6000803e3d6000fd5b505050506040513d6020811015614d1f57600080fd5b505160028201556040805163364df27760e01b8152905173__$0d86abb4a722a612872fb80f4c7e7e95bd$__9163364df277916004808301926020929190829003018186803b158015614d7157600080fd5b505af4158015614d85573d6000803e3d6000fd5b505050506040513d6020811015614d9b57600080fd5b505160038201555b88612002576120028a61513c565b6001600160a01b038116614dc457600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160601b03198516851461506757600085815260036020908152604080832090516392516ac760e01b81526004810183815285516024830152855192949373__$6b4cc75dad3e0abd6ad83b3d907747c608$__93632a0500d8938c9373__$0d86abb4a722a612872fb80f4c7e7e95bd$__936392516ac7938b938392604490910191908501908083838e5b83811015614ec5578181015183820152602001614ead565b50505050905090810190601f168015614ef25780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015614f0f57600080fd5b505af4158015614f23573d6000803e3d6000fd5b505050506040513d6020811015614f3957600080fd5b5051604080516001600160e01b031960e086901b168152600481019390935260248301919091526affffffffffffffffffffff198a16604483015260648201899052608482018890525160a4808301926020929190829003018186803b158015614fa257600080fd5b505af4158015614fb6573d6000803e3d6000fd5b505050506040513d6020811015614fcc57600080fd5b5051600383015460408051636bc68c7560e11b81526004810192909252602482018390525191925073__$6b4cc75dad3e0abd6ad83b3d907747c608$__9163d78d18ea91604480820192602092909190829003018186803b15801561503057600080fd5b505af4158015615044573d6000803e3d6000fd5b505050506040513d602081101561505a57600080fd5b5051600390920191909155505b847f983eec3b16cf85cbe329d288a7863b78dc7d1c9a0d690b21434d4cc1f73539e58386868560405180858152602001846affffffffffffffffffffff19166affffffffffffffffffffff1916815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156150f85781810151838201526020016150e0565b50505050905090810190601f1680156151255780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a25050505050565b6000908152600360208190526040822082815560018101839055600281018390559081018290556004810182905560058101829055600681018290556007810180546001600160a01b0319908116909155600882018054821690556009820180549091169055600a810191909155600b0180546dffffffffffffffffffffffffffff19169055565b6000828211156151d357600080fd5b5090039056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6e417373657274696f6e206d7573742062652070656e64696e6720746f20696e697469617465206368616c6c656e676553656c65637465642063757272656e6379206973206e6f7420616e2061636365707465642074797065417373657274696f6e206973207374696c6c2070656e64696e67206368616c6c656e6765564d20646f6573206e6f74206861766520617373657274696f6e2070656e64696e67507265636f6e646974696f6e3a206e6f742077697468696e2074696d6520626f756e64734368616c6c656e6765207761732063726561746564206279206173736572746572507265636f6e646974696f6e3a20696e626f7820646f6573206e6f74206d6174636856616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f77507265636f6e646974696f6e3a207374617465206861736820646f6573206e6f74206d617463684368616c6c656e676520646964206e6f7420636f6d65206265666f726520646561646c696e6543616e206f6e6c792064697370757461626c65206173736572742066726f6d2077616974696e67207374617465a265627a7a7230582013bd0a0c1a00da6b81789521f37c030dee2d426e75aea9d817d77c2bad3f19c964736f6c63430005090032`

// DeployVMTracker deploys a new Ethereum contract, binding an instance of VMTracker to it.
func DeployVMTracker(auth *bind.TransactOpts, backend bind.ContractBackend, _balanceTrackerAddress common.Address) (common.Address, *types.Transaction, *VMTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(VMTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VMTrackerBin), backend, _balanceTrackerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VMTracker{VMTrackerCaller: VMTrackerCaller{contract: contract}, VMTrackerTransactor: VMTrackerTransactor{contract: contract}, VMTrackerFilterer: VMTrackerFilterer{contract: contract}}, nil
}

// VMTracker is an auto generated Go binding around an Ethereum contract.
type VMTracker struct {
	VMTrackerCaller     // Read-only binding to the contract
	VMTrackerTransactor // Write-only binding to the contract
	VMTrackerFilterer   // Log filterer for contract events
}

// VMTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VMTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VMTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VMTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VMTrackerSession struct {
	Contract     *VMTracker        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VMTrackerCallerSession struct {
	Contract *VMTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VMTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VMTrackerTransactorSession struct {
	Contract     *VMTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VMTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VMTrackerRaw struct {
	Contract *VMTracker // Generic contract binding to access the raw methods on
}

// VMTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VMTrackerCallerRaw struct {
	Contract *VMTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// VMTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VMTrackerTransactorRaw struct {
	Contract *VMTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVMTracker creates a new instance of VMTracker, bound to a specific deployed contract.
func NewVMTracker(address common.Address, backend bind.ContractBackend) (*VMTracker, error) {
	contract, err := bindVMTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VMTracker{VMTrackerCaller: VMTrackerCaller{contract: contract}, VMTrackerTransactor: VMTrackerTransactor{contract: contract}, VMTrackerFilterer: VMTrackerFilterer{contract: contract}}, nil
}

// NewVMTrackerCaller creates a new read-only instance of VMTracker, bound to a specific deployed contract.
func NewVMTrackerCaller(address common.Address, caller bind.ContractCaller) (*VMTrackerCaller, error) {
	contract, err := bindVMTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VMTrackerCaller{contract: contract}, nil
}

// NewVMTrackerTransactor creates a new write-only instance of VMTracker, bound to a specific deployed contract.
func NewVMTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*VMTrackerTransactor, error) {
	contract, err := bindVMTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VMTrackerTransactor{contract: contract}, nil
}

// NewVMTrackerFilterer creates a new log filterer instance of VMTracker, bound to a specific deployed contract.
func NewVMTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*VMTrackerFilterer, error) {
	contract, err := bindVMTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VMTrackerFilterer{contract: contract}, nil
}

// bindVMTracker binds a generic wrapper to an already deployed contract.
func bindVMTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VMTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VMTracker *VMTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VMTracker.Contract.VMTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VMTracker *VMTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMTracker.Contract.VMTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VMTracker *VMTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VMTracker.Contract.VMTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VMTracker *VMTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VMTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VMTracker *VMTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VMTracker *VMTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VMTracker.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VMTracker *VMTrackerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _VMTracker.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VMTracker *VMTrackerSession) IsOwner() (bool, error) {
	return _VMTracker.Contract.IsOwner(&_VMTracker.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VMTracker *VMTrackerCallerSession) IsOwner() (bool, error) {
	return _VMTracker.Contract.IsOwner(&_VMTracker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VMTracker *VMTrackerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)

	out := ret0
	err := _VMTracker.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VMTracker *VMTrackerSession) Owner() (common.Address, error) {
	return _VMTracker.Contract.Owner(&_VMTracker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VMTracker *VMTrackerCallerSession) Owner() (common.Address, error) {
	return _VMTracker.Contract.Owner(&_VMTracker.CallOpts)
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_VMTracker *VMTrackerCaller) WithinTimeBounds(opts *bind.CallOpts, _timeBounds [2]uint64) (bool, error) {
	ret0 := new(bool)

	out := ret0
	err := _VMTracker.contract.Call(opts, out, "withinTimeBounds", _timeBounds)
	return *ret0, err
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_VMTracker *VMTrackerSession) WithinTimeBounds(_timeBounds [2]uint64) (bool, error) {
	return _VMTracker.Contract.WithinTimeBounds(&_VMTracker.CallOpts, _timeBounds)
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_VMTracker *VMTrackerCallerSession) WithinTimeBounds(_timeBounds [2]uint64) (bool, error) {
	return _VMTracker.Contract.WithinTimeBounds(&_VMTracker.CallOpts, _timeBounds)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0xd0b41b6c.
//
// Solidity: function ConfirmUnanimousAsserted(bytes32 _vmId, bytes32 _afterHash, bytes32 _logsHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerTransactor) ConfirmUnanimousAsserted(opts *bind.TransactOpts, _vmId [32]byte, _afterHash [32]byte, _logsHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "ConfirmUnanimousAsserted", _vmId, _afterHash, _logsHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0xd0b41b6c.
//
// Solidity: function ConfirmUnanimousAsserted(bytes32 _vmId, bytes32 _afterHash, bytes32 _logsHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerSession) ConfirmUnanimousAsserted(_vmId [32]byte, _afterHash [32]byte, _logsHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmUnanimousAsserted(&_VMTracker.TransactOpts, _vmId, _afterHash, _logsHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0xd0b41b6c.
//
// Solidity: function ConfirmUnanimousAsserted(bytes32 _vmId, bytes32 _afterHash, bytes32 _logsHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerTransactorSession) ConfirmUnanimousAsserted(_vmId [32]byte, _afterHash [32]byte, _logsHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmUnanimousAsserted(&_VMTracker.TransactOpts, _vmId, _afterHash, _logsHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// AddChallengeManager is a paid mutator transaction binding the contract method 0x2e6bf2e6.
//
// Solidity: function addChallengeManager(address _challengeManager) returns()
func (_VMTracker *VMTrackerTransactor) AddChallengeManager(opts *bind.TransactOpts, _challengeManager common.Address) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "addChallengeManager", _challengeManager)
}

// AddChallengeManager is a paid mutator transaction binding the contract method 0x2e6bf2e6.
//
// Solidity: function addChallengeManager(address _challengeManager) returns()
func (_VMTracker *VMTrackerSession) AddChallengeManager(_challengeManager common.Address) (*types.Transaction, error) {
	return _VMTracker.Contract.AddChallengeManager(&_VMTracker.TransactOpts, _challengeManager)
}

// AddChallengeManager is a paid mutator transaction binding the contract method 0x2e6bf2e6.
//
// Solidity: function addChallengeManager(address _challengeManager) returns()
func (_VMTracker *VMTrackerTransactorSession) AddChallengeManager(_challengeManager common.Address) (*types.Transaction, error) {
	return _VMTracker.Contract.AddChallengeManager(&_VMTracker.TransactOpts, _challengeManager)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_VMTracker *VMTrackerTransactor) CompleteChallenge(opts *bind.TransactOpts, _vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "completeChallenge", _vmId, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_VMTracker *VMTrackerSession) CompleteChallenge(_vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _VMTracker.Contract.CompleteChallenge(&_VMTracker.TransactOpts, _vmId, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_VMTracker *VMTrackerTransactorSession) CompleteChallenge(_vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _VMTracker.Contract.CompleteChallenge(&_VMTracker.TransactOpts, _vmId, _players, _rewards)
}

// ConfirmAsserted is a paid mutator transaction binding the contract method 0x89d9957e.
//
// Solidity: function confirmAsserted(bytes32 _vmId, bytes32 _preconditionHash, bytes32 _afterHash, bytes32 _logsHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerTransactor) ConfirmAsserted(opts *bind.TransactOpts, _vmId [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _logsHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "confirmAsserted", _vmId, _preconditionHash, _afterHash, _logsHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestination)
}

// ConfirmAsserted is a paid mutator transaction binding the contract method 0x89d9957e.
//
// Solidity: function confirmAsserted(bytes32 _vmId, bytes32 _preconditionHash, bytes32 _afterHash, bytes32 _logsHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerSession) ConfirmAsserted(_vmId [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _logsHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmAsserted(&_VMTracker.TransactOpts, _vmId, _preconditionHash, _afterHash, _logsHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestination)
}

// ConfirmAsserted is a paid mutator transaction binding the contract method 0x89d9957e.
//
// Solidity: function confirmAsserted(bytes32 _vmId, bytes32 _preconditionHash, bytes32 _afterHash, bytes32 _logsHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerTransactorSession) ConfirmAsserted(_vmId [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _logsHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmAsserted(&_VMTracker.TransactOpts, _vmId, _preconditionHash, _afterHash, _logsHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestination)
}

// CreateVm is a paid mutator transaction binding the contract method 0x8dbc44f6.
//
// Solidity: function createVm(bytes32[3] _fields, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint16 _challengeManagerNum, uint128 _escrowRequired, address _escrowCurrency, address _owner, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactor) CreateVm(opts *bind.TransactOpts, _fields [3][32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _challengeManagerNum uint16, _escrowRequired *big.Int, _escrowCurrency common.Address, _owner common.Address, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "createVm", _fields, _gracePeriod, _maxExecutionSteps, _challengeManagerNum, _escrowRequired, _escrowCurrency, _owner, _signatures)
}

// CreateVm is a paid mutator transaction binding the contract method 0x8dbc44f6.
//
// Solidity: function createVm(bytes32[3] _fields, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint16 _challengeManagerNum, uint128 _escrowRequired, address _escrowCurrency, address _owner, bytes _signatures) returns()
func (_VMTracker *VMTrackerSession) CreateVm(_fields [3][32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _challengeManagerNum uint16, _escrowRequired *big.Int, _escrowCurrency common.Address, _owner common.Address, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.CreateVm(&_VMTracker.TransactOpts, _fields, _gracePeriod, _maxExecutionSteps, _challengeManagerNum, _escrowRequired, _escrowCurrency, _owner, _signatures)
}

// CreateVm is a paid mutator transaction binding the contract method 0x8dbc44f6.
//
// Solidity: function createVm(bytes32[3] _fields, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint16 _challengeManagerNum, uint128 _escrowRequired, address _escrowCurrency, address _owner, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactorSession) CreateVm(_fields [3][32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _challengeManagerNum uint16, _escrowRequired *big.Int, _escrowCurrency common.Address, _owner common.Address, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.CreateVm(&_VMTracker.TransactOpts, _fields, _gracePeriod, _maxExecutionSteps, _challengeManagerNum, _escrowRequired, _escrowCurrency, _owner, _signatures)
}

// DisputableAssert is a paid mutator transaction binding the contract method 0x1c0fc42c.
//
// Solidity: function disputableAssert(bytes32[5] _fields, uint32 _numSteps, uint64[2] timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNum, uint256[] _msgAmount, bytes32[] _msgDestination) returns()
func (_VMTracker *VMTrackerTransactor) DisputableAssert(opts *bind.TransactOpts, _fields [5][32]byte, _numSteps uint32, timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNum []uint16, _msgAmount []*big.Int, _msgDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "disputableAssert", _fields, _numSteps, timeBounds, _tokenTypes, _messageDataHash, _messageTokenNum, _msgAmount, _msgDestination)
}

// DisputableAssert is a paid mutator transaction binding the contract method 0x1c0fc42c.
//
// Solidity: function disputableAssert(bytes32[5] _fields, uint32 _numSteps, uint64[2] timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNum, uint256[] _msgAmount, bytes32[] _msgDestination) returns()
func (_VMTracker *VMTrackerSession) DisputableAssert(_fields [5][32]byte, _numSteps uint32, timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNum []uint16, _msgAmount []*big.Int, _msgDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.DisputableAssert(&_VMTracker.TransactOpts, _fields, _numSteps, timeBounds, _tokenTypes, _messageDataHash, _messageTokenNum, _msgAmount, _msgDestination)
}

// DisputableAssert is a paid mutator transaction binding the contract method 0x1c0fc42c.
//
// Solidity: function disputableAssert(bytes32[5] _fields, uint32 _numSteps, uint64[2] timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNum, uint256[] _msgAmount, bytes32[] _msgDestination) returns()
func (_VMTracker *VMTrackerTransactorSession) DisputableAssert(_fields [5][32]byte, _numSteps uint32, timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNum []uint16, _msgAmount []*big.Int, _msgDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.DisputableAssert(&_VMTracker.TransactOpts, _fields, _numSteps, timeBounds, _tokenTypes, _messageDataHash, _messageTokenNum, _msgAmount, _msgDestination)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0xffd134c9.
//
// Solidity: function initiateChallenge(bytes32 _vmId, bytes32 _assertPreHash) returns()
func (_VMTracker *VMTrackerTransactor) InitiateChallenge(opts *bind.TransactOpts, _vmId [32]byte, _assertPreHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "initiateChallenge", _vmId, _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0xffd134c9.
//
// Solidity: function initiateChallenge(bytes32 _vmId, bytes32 _assertPreHash) returns()
func (_VMTracker *VMTrackerSession) InitiateChallenge(_vmId [32]byte, _assertPreHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.InitiateChallenge(&_VMTracker.TransactOpts, _vmId, _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0xffd134c9.
//
// Solidity: function initiateChallenge(bytes32 _vmId, bytes32 _assertPreHash) returns()
func (_VMTracker *VMTrackerTransactorSession) InitiateChallenge(_vmId [32]byte, _assertPreHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.InitiateChallenge(&_VMTracker.TransactOpts, _vmId, _assertPreHash)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xfd0961b6.
//
// Solidity: function ownerShutdown(bytes32 _vmId) returns()
func (_VMTracker *VMTrackerTransactor) OwnerShutdown(opts *bind.TransactOpts, _vmId [32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "ownerShutdown", _vmId)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xfd0961b6.
//
// Solidity: function ownerShutdown(bytes32 _vmId) returns()
func (_VMTracker *VMTrackerSession) OwnerShutdown(_vmId [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.OwnerShutdown(&_VMTracker.TransactOpts, _vmId)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xfd0961b6.
//
// Solidity: function ownerShutdown(bytes32 _vmId) returns()
func (_VMTracker *VMTrackerTransactorSession) OwnerShutdown(_vmId [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.OwnerShutdown(&_VMTracker.TransactOpts, _vmId)
}

// ProposeUnanimousAssert is a paid mutator transaction binding the contract method 0x7d5d7b77.
//
// Solidity: function proposeUnanimousAssert(bytes32 _vmId, bytes32 _unanRest, uint64 _sequenceNum, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactor) ProposeUnanimousAssert(opts *bind.TransactOpts, _vmId [32]byte, _unanRest [32]byte, _sequenceNum uint64, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "proposeUnanimousAssert", _vmId, _unanRest, _sequenceNum, _timeBounds, _tokenTypes, _messageTokenNum, _messageAmount, _signatures)
}

// ProposeUnanimousAssert is a paid mutator transaction binding the contract method 0x7d5d7b77.
//
// Solidity: function proposeUnanimousAssert(bytes32 _vmId, bytes32 _unanRest, uint64 _sequenceNum, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes _signatures) returns()
func (_VMTracker *VMTrackerSession) ProposeUnanimousAssert(_vmId [32]byte, _unanRest [32]byte, _sequenceNum uint64, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ProposeUnanimousAssert(&_VMTracker.TransactOpts, _vmId, _unanRest, _sequenceNum, _timeBounds, _tokenTypes, _messageTokenNum, _messageAmount, _signatures)
}

// ProposeUnanimousAssert is a paid mutator transaction binding the contract method 0x7d5d7b77.
//
// Solidity: function proposeUnanimousAssert(bytes32 _vmId, bytes32 _unanRest, uint64 _sequenceNum, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactorSession) ProposeUnanimousAssert(_vmId [32]byte, _unanRest [32]byte, _sequenceNum uint64, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ProposeUnanimousAssert(&_VMTracker.TransactOpts, _vmId, _unanRest, _sequenceNum, _timeBounds, _tokenTypes, _messageTokenNum, _messageAmount, _signatures)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VMTracker *VMTrackerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VMTracker *VMTrackerSession) RenounceOwnership() (*types.Transaction, error) {
	return _VMTracker.Contract.RenounceOwnership(&_VMTracker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VMTracker *VMTrackerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VMTracker.Contract.RenounceOwnership(&_VMTracker.TransactOpts)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0xfa3eebe6.
//
// Solidity: function sendEthMessage(bytes32 _destination, bytes _data) returns()
func (_VMTracker *VMTrackerTransactor) SendEthMessage(opts *bind.TransactOpts, _destination [32]byte, _data []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "sendEthMessage", _destination, _data)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0xfa3eebe6.
//
// Solidity: function sendEthMessage(bytes32 _destination, bytes _data) returns()
func (_VMTracker *VMTrackerSession) SendEthMessage(_destination [32]byte, _data []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.SendEthMessage(&_VMTracker.TransactOpts, _destination, _data)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0xfa3eebe6.
//
// Solidity: function sendEthMessage(bytes32 _destination, bytes _data) returns()
func (_VMTracker *VMTrackerTransactorSession) SendEthMessage(_destination [32]byte, _data []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.SendEthMessage(&_VMTracker.TransactOpts, _destination, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3cd8a322.
//
// Solidity: function sendMessage(bytes32 _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_VMTracker *VMTrackerTransactor) SendMessage(opts *bind.TransactOpts, _destination [32]byte, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "sendMessage", _destination, _tokenType, _amount, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3cd8a322.
//
// Solidity: function sendMessage(bytes32 _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_VMTracker *VMTrackerSession) SendMessage(_destination [32]byte, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.SendMessage(&_VMTracker.TransactOpts, _destination, _tokenType, _amount, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3cd8a322.
//
// Solidity: function sendMessage(bytes32 _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_VMTracker *VMTrackerTransactorSession) SendMessage(_destination [32]byte, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.SendMessage(&_VMTracker.TransactOpts, _destination, _tokenType, _amount, _data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VMTracker *VMTrackerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VMTracker *VMTrackerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VMTracker.Contract.TransferOwnership(&_VMTracker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VMTracker *VMTrackerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VMTracker.Contract.TransferOwnership(&_VMTracker.TransactOpts, newOwner)
}

// UnanimousAssert is a paid mutator transaction binding the contract method 0x6c94603a.
//
// Solidity: function unanimousAssert(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes32 _logsHash, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactor) UnanimousAssert(opts *bind.TransactOpts, _vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _logsHash [32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "unanimousAssert", _vmId, _afterHash, _newInbox, _logsHash, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination, _signatures)
}

// UnanimousAssert is a paid mutator transaction binding the contract method 0x6c94603a.
//
// Solidity: function unanimousAssert(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes32 _logsHash, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination, bytes _signatures) returns()
func (_VMTracker *VMTrackerSession) UnanimousAssert(_vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _logsHash [32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.UnanimousAssert(&_VMTracker.TransactOpts, _vmId, _afterHash, _newInbox, _logsHash, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination, _signatures)
}

// UnanimousAssert is a paid mutator transaction binding the contract method 0x6c94603a.
//
// Solidity: function unanimousAssert(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes32 _logsHash, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactorSession) UnanimousAssert(_vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _logsHash [32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.UnanimousAssert(&_VMTracker.TransactOpts, _vmId, _afterHash, _newInbox, _logsHash, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination, _signatures)
}

// VMTrackerConfirmedAssertionIterator is returned from FilterConfirmedAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedAssertion events raised by the VMTracker contract.
type VMTrackerConfirmedAssertionIterator struct {
	Event *VMTrackerConfirmedAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerConfirmedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerConfirmedAssertion)
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
		it.Event = new(VMTrackerConfirmedAssertion)
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
func (it *VMTrackerConfirmedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerConfirmedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerConfirmedAssertion represents a ConfirmedAssertion event raised by the VMTracker contract.
type VMTrackerConfirmedAssertion struct {
	VmId     [32]byte
	NewState [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfirmedAssertion is a free log retrieval operation binding the contract event 0x8bf82ba359336cefa1c284ed9d3864f99aa1cd41dddf9330afd656bcb9dd7a7b.
//
// Solidity: event ConfirmedAssertion(bytes32 indexed vmId, bytes32 newState)
func (_VMTracker *VMTrackerFilterer) FilterConfirmedAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerConfirmedAssertionIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "ConfirmedAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerConfirmedAssertionIterator{contract: _VMTracker.contract, event: "ConfirmedAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedAssertion is a free log subscription operation binding the contract event 0x8bf82ba359336cefa1c284ed9d3864f99aa1cd41dddf9330afd656bcb9dd7a7b.
//
// Solidity: event ConfirmedAssertion(bytes32 indexed vmId, bytes32 newState)
func (_VMTracker *VMTrackerFilterer) WatchConfirmedAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerConfirmedAssertion, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "ConfirmedAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerConfirmedAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "ConfirmedAssertion", log); err != nil {
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

// VMTrackerConfirmedUnanimousAssertionIterator is returned from FilterConfirmedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedUnanimousAssertion events raised by the VMTracker contract.
type VMTrackerConfirmedUnanimousAssertionIterator struct {
	Event *VMTrackerConfirmedUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerConfirmedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerConfirmedUnanimousAssertion)
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
		it.Event = new(VMTrackerConfirmedUnanimousAssertion)
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
func (it *VMTrackerConfirmedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerConfirmedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerConfirmedUnanimousAssertion represents a ConfirmedUnanimousAssertion event raised by the VMTracker contract.
type VMTrackerConfirmedUnanimousAssertion struct {
	VmId        [32]byte
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedUnanimousAssertion is a free log retrieval operation binding the contract event 0xd23ef3a25058f2acee4ca7aee5cc78a587b9fc76d2875098af1ffdfc2bcc74c9.
//
// Solidity: event ConfirmedUnanimousAssertion(bytes32 indexed vmId, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) FilterConfirmedUnanimousAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerConfirmedUnanimousAssertionIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "ConfirmedUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerConfirmedUnanimousAssertionIterator{contract: _VMTracker.contract, event: "ConfirmedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedUnanimousAssertion is a free log subscription operation binding the contract event 0xd23ef3a25058f2acee4ca7aee5cc78a587b9fc76d2875098af1ffdfc2bcc74c9.
//
// Solidity: event ConfirmedUnanimousAssertion(bytes32 indexed vmId, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) WatchConfirmedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerConfirmedUnanimousAssertion, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "ConfirmedUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerConfirmedUnanimousAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
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

// VMTrackerDisputableAssertionIterator is returned from FilterDisputableAssertion and is used to iterate over the raw logs and unpacked data for DisputableAssertion events raised by the VMTracker contract.
type VMTrackerDisputableAssertionIterator struct {
	Event *VMTrackerDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerDisputableAssertion)
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
		it.Event = new(VMTrackerDisputableAssertion)
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
func (it *VMTrackerDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerDisputableAssertion represents a DisputableAssertion event raised by the VMTracker contract.
type VMTrackerDisputableAssertion struct {
	VmId            [32]byte
	Fields          [3][32]byte
	Asserter        common.Address
	TimeBounds      [2]uint64
	TokenTypes      [][21]byte
	NumSteps        uint32
	LastMessageHash [32]byte
	LogsHash        [32]byte
	Amounts         []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDisputableAssertion is a free log retrieval operation binding the contract event 0xe3457847513bb99aa86c769d9800b4a4c1f0758755234d78d40c406de93935d1.
//
// Solidity: event DisputableAssertion(bytes32 indexed vmId, bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsHash, uint256[] amounts)
func (_VMTracker *VMTrackerFilterer) FilterDisputableAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerDisputableAssertionIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "DisputableAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerDisputableAssertionIterator{contract: _VMTracker.contract, event: "DisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchDisputableAssertion is a free log subscription operation binding the contract event 0xe3457847513bb99aa86c769d9800b4a4c1f0758755234d78d40c406de93935d1.
//
// Solidity: event DisputableAssertion(bytes32 indexed vmId, bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsHash, uint256[] amounts)
func (_VMTracker *VMTrackerFilterer) WatchDisputableAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerDisputableAssertion, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "DisputableAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerDisputableAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "DisputableAssertion", log); err != nil {
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

// VMTrackerFinalUnanimousAssertionIterator is returned from FilterFinalUnanimousAssertion and is used to iterate over the raw logs and unpacked data for FinalUnanimousAssertion events raised by the VMTracker contract.
type VMTrackerFinalUnanimousAssertionIterator struct {
	Event *VMTrackerFinalUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerFinalUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerFinalUnanimousAssertion)
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
		it.Event = new(VMTrackerFinalUnanimousAssertion)
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
func (it *VMTrackerFinalUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerFinalUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerFinalUnanimousAssertion represents a FinalUnanimousAssertion event raised by the VMTracker contract.
type VMTrackerFinalUnanimousAssertion struct {
	VmId     [32]byte
	UnanHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalUnanimousAssertion is a free log retrieval operation binding the contract event 0xd13341838a6e5e2972e2b8638f8d2652f77945f3755d3a2829ec5076b5147067.
//
// Solidity: event FinalUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash)
func (_VMTracker *VMTrackerFilterer) FilterFinalUnanimousAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerFinalUnanimousAssertionIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "FinalUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerFinalUnanimousAssertionIterator{contract: _VMTracker.contract, event: "FinalUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchFinalUnanimousAssertion is a free log subscription operation binding the contract event 0xd13341838a6e5e2972e2b8638f8d2652f77945f3755d3a2829ec5076b5147067.
//
// Solidity: event FinalUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash)
func (_VMTracker *VMTrackerFilterer) WatchFinalUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerFinalUnanimousAssertion, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "FinalUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerFinalUnanimousAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "FinalUnanimousAssertion", log); err != nil {
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

// VMTrackerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the VMTracker contract.
type VMTrackerInitiatedChallengeIterator struct {
	Event *VMTrackerInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *VMTrackerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerInitiatedChallenge)
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
		it.Event = new(VMTrackerInitiatedChallenge)
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
func (it *VMTrackerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerInitiatedChallenge represents a InitiatedChallenge event raised by the VMTracker contract.
type VMTrackerInitiatedChallenge struct {
	VmId       [32]byte
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xb8d6a904566965067cb7e3e946d2a494a4b1f955ffec3e72e1e6ebe4e8e52c72.
//
// Solidity: event InitiatedChallenge(bytes32 indexed vmId, address challenger)
func (_VMTracker *VMTrackerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerInitiatedChallengeIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "InitiatedChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerInitiatedChallengeIterator{contract: _VMTracker.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xb8d6a904566965067cb7e3e946d2a494a4b1f955ffec3e72e1e6ebe4e8e52c72.
//
// Solidity: event InitiatedChallenge(bytes32 indexed vmId, address challenger)
func (_VMTracker *VMTrackerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *VMTrackerInitiatedChallenge, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "InitiatedChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerInitiatedChallenge)
				if err := _VMTracker.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// VMTrackerMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the VMTracker contract.
type VMTrackerMessageDeliveredIterator struct {
	Event *VMTrackerMessageDelivered // Event containing the contract specifics and raw log

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
func (it *VMTrackerMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerMessageDelivered)
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
		it.Event = new(VMTrackerMessageDelivered)
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
func (it *VMTrackerMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerMessageDelivered represents a MessageDelivered event raised by the VMTracker contract.
type VMTrackerMessageDelivered struct {
	VmId        [32]byte
	Destination [32]byte
	TokenType   [21]byte
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x983eec3b16cf85cbe329d288a7863b78dc7d1c9a0d690b21434d4cc1f73539e5.
//
// Solidity: event MessageDelivered(bytes32 indexed vmId, bytes32 destination, bytes21 tokenType, uint256 value, bytes data)
func (_VMTracker *VMTrackerFilterer) FilterMessageDelivered(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerMessageDeliveredIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "MessageDelivered", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerMessageDeliveredIterator{contract: _VMTracker.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x983eec3b16cf85cbe329d288a7863b78dc7d1c9a0d690b21434d4cc1f73539e5.
//
// Solidity: event MessageDelivered(bytes32 indexed vmId, bytes32 destination, bytes21 tokenType, uint256 value, bytes data)
func (_VMTracker *VMTrackerFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *VMTrackerMessageDelivered, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "MessageDelivered", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerMessageDelivered)
				if err := _VMTracker.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// VMTrackerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VMTracker contract.
type VMTrackerOwnershipTransferredIterator struct {
	Event *VMTrackerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VMTrackerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerOwnershipTransferred)
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
		it.Event = new(VMTrackerOwnershipTransferred)
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
func (it *VMTrackerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerOwnershipTransferred represents a OwnershipTransferred event raised by the VMTracker contract.
type VMTrackerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VMTracker *VMTrackerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VMTrackerOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerOwnershipTransferredIterator{contract: _VMTracker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VMTracker *VMTrackerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VMTrackerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerOwnershipTransferred)
				if err := _VMTracker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// VMTrackerProposedUnanimousAssertionIterator is returned from FilterProposedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for ProposedUnanimousAssertion events raised by the VMTracker contract.
type VMTrackerProposedUnanimousAssertionIterator struct {
	Event *VMTrackerProposedUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerProposedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerProposedUnanimousAssertion)
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
		it.Event = new(VMTrackerProposedUnanimousAssertion)
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
func (it *VMTrackerProposedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerProposedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerProposedUnanimousAssertion represents a ProposedUnanimousAssertion event raised by the VMTracker contract.
type VMTrackerProposedUnanimousAssertion struct {
	VmId        [32]byte
	UnanHash    [32]byte
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposedUnanimousAssertion is a free log retrieval operation binding the contract event 0xa1d1249870ac5b939cf041991350480912abf0af99cf0337b13e440e1e150b96.
//
// Solidity: event ProposedUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) FilterProposedUnanimousAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerProposedUnanimousAssertionIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "ProposedUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerProposedUnanimousAssertionIterator{contract: _VMTracker.contract, event: "ProposedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchProposedUnanimousAssertion is a free log subscription operation binding the contract event 0xa1d1249870ac5b939cf041991350480912abf0af99cf0337b13e440e1e150b96.
//
// Solidity: event ProposedUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) WatchProposedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerProposedUnanimousAssertion, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "ProposedUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerProposedUnanimousAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "ProposedUnanimousAssertion", log); err != nil {
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

// VMTrackerVMCreatedIterator is returned from FilterVMCreated and is used to iterate over the raw logs and unpacked data for VMCreated events raised by the VMTracker contract.
type VMTrackerVMCreatedIterator struct {
	Event *VMTrackerVMCreated // Event containing the contract specifics and raw log

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
func (it *VMTrackerVMCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerVMCreated)
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
		it.Event = new(VMTrackerVMCreated)
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
func (it *VMTrackerVMCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerVMCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerVMCreated represents a VMCreated event raised by the VMTracker contract.
type VMTrackerVMCreated struct {
	GracePeriod         uint32
	EscrowRequired      *big.Int
	EscrowCurrency      common.Address
	MaxExecutionSteps   uint32
	VmId                [32]byte
	VmState             [32]byte
	ChallengeManagerNum uint16
	Owner               common.Address
	Validators          []common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVMCreated is a free log retrieval operation binding the contract event 0xe251d26fa5921a42fd8dfa2415cfa1f48aae04e9c147d0edb7fd63d5bdc9e6f4.
//
// Solidity: event VMCreated(uint32 _gracePeriod, uint128 _escrowRequired, address _escrowCurrency, uint32 _maxExecutionSteps, bytes32 _vmId, bytes32 _vmState, uint16 _challengeManagerNum, address _owner, address[] validators)
func (_VMTracker *VMTrackerFilterer) FilterVMCreated(opts *bind.FilterOpts) (*VMTrackerVMCreatedIterator, error) {
	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "VMCreated")
	if err != nil {
		return nil, err
	}
	return &VMTrackerVMCreatedIterator{contract: _VMTracker.contract, event: "VMCreated", logs: logs, sub: sub}, nil
}

// WatchVMCreated is a free log subscription operation binding the contract event 0xe251d26fa5921a42fd8dfa2415cfa1f48aae04e9c147d0edb7fd63d5bdc9e6f4.
//
// Solidity: event VMCreated(uint32 _gracePeriod, uint128 _escrowRequired, address _escrowCurrency, uint32 _maxExecutionSteps, bytes32 _vmId, bytes32 _vmState, uint16 _challengeManagerNum, address _owner, address[] validators)
func (_VMTracker *VMTrackerFilterer) WatchVMCreated(opts *bind.WatchOpts, sink chan<- *VMTrackerVMCreated) (event.Subscription, error) {
	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "VMCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerVMCreated)
				if err := _VMTracker.contract.UnpackLog(event, "VMCreated", log); err != nil {
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
