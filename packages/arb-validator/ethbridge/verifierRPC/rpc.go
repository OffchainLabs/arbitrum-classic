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
var AddressBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820c3fc7d1e54cf0136b207b3a01cef207653566dfd21575adb15a2f70bc363a0c664736f6c634300050a0032"

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
const ArbBalanceTrackerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"bytes32\"}],\"name\":\"getNFTTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"bytes32\"},{\"name\":\"_to\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"bytes32\"}],\"name\":\"getTokenBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"bytes32\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasNFT\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"bytes32\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"ownerRemoveToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"hasFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"bytes32\"},{\"name\":\"_to\",\"type\":\"bytes32\"},{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferNFT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"bytes32\"}],\"name\":\"depositEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenContract\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ArbBalanceTrackerFuncSigs maps the 4-byte function signature to its string representation.
var ArbBalanceTrackerFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"97feb926": "depositERC20(address,uint256)",
	"d29a4bf6": "depositERC721(address,uint256)",
	"da63d7b6": "depositEth(bytes32)",
	"0aa114ae": "getNFTTokens(bytes32)",
	"a35be443": "getTokenBalance(address,bytes32)",
	"2a8a8e7f": "getTokenBalances(bytes32)",
	"c2465106": "hasFunds(bytes32,bytes21[],uint256[])",
	"82512757": "hasNFT(address,bytes32,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"b8569ccd": "ownerRemoveToken(bytes32,address,uint256)",
	"715018a6": "renounceOwnership()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"cdf25dc1": "transferNFT(bytes32,bytes32,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"1240117b": "transferToken(bytes32,bytes32,address,uint256)",
	"a1db9782": "withdrawERC20(address,uint256)",
	"f3e414f8": "withdrawERC721(address,uint256)",
	"c311d049": "withdrawEth(uint256)",
}

// ArbBalanceTrackerBin is the compiled bytecode used for deploying new contracts.
var ArbBalanceTrackerBin = "0x60806040819052600080546001600160a01b03191633178082556001600160a01b0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3612127806100576000396000f3fe6080604052600436106101815760003560e01c8063a1db9782116100d1578063c311d0491161008a578063da63d7b611610064578063da63d7b6146107b6578063dd62ed3e146107d3578063f2fde38b1461080e578063f3e414f81461084157610181565b8063c311d0491461070e578063cdf25dc114610738578063d29a4bf61461077d57610181565b8063a1db9782146104b0578063a35be443146104e9578063a457c2d714610522578063a9059cbb1461055b578063b8569ccd14610594578063c2465106146105d357610181565b8063395093511161013e578063825127571161011857806382512757146103f25780638da5cb5b146104315780638f32d59b1461046257806397feb9261461047757610181565b8063395093511461037157806370a08231146103aa578063715018a6146103dd57610181565b8063095ea7b3146101865780630aa114ae146101d35780631240117b1461029657806318160ddd146102dd57806323b872dd146103045780632a8a8e7f14610347575b600080fd5b34801561019257600080fd5b506101bf600480360360408110156101a957600080fd5b506001600160a01b03813516906020013561087a565b604080519115158252519081900360200190f35b3480156101df57600080fd5b506101fd600480360360208110156101f657600080fd5b5035610891565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610241578181015183820152602001610229565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610280578181015183820152602001610268565b5050505090500194505050505060405180910390f35b3480156102a257600080fd5b506102db600480360360808110156102b957600080fd5b508035906020810135906001600160a01b036040820135169060600135610a1f565b005b3480156102e957600080fd5b506102f2610abd565b60408051918252519081900360200190f35b34801561031057600080fd5b506101bf6004803603606081101561032757600080fd5b506001600160a01b03813581169160208101359091169060400135610ac4565b34801561035357600080fd5b506101fd6004803603602081101561036a57600080fd5b5035610b1c565b34801561037d57600080fd5b506101bf6004803603604081101561039457600080fd5b506001600160a01b038135169060200135610c48565b3480156103b657600080fd5b506102f2600480360360208110156103cd57600080fd5b50356001600160a01b0316610c84565b3480156103e957600080fd5b506102db610c9f565b3480156103fe57600080fd5b506101bf6004803603606081101561041557600080fd5b506001600160a01b038135169060208101359060400135610d30565b34801561043d57600080fd5b50610446610da5565b604080516001600160a01b039092168252519081900360200190f35b34801561046e57600080fd5b506101bf610db4565b34801561048357600080fd5b506102db6004803603604081101561049a57600080fd5b506001600160a01b038135169060200135610dc5565b3480156104bc57600080fd5b506102db600480360360408110156104d357600080fd5b506001600160a01b038135169060200135610e63565b3480156104f557600080fd5b506102f26004803603604081101561050c57600080fd5b506001600160a01b038135169060200135610f2f565b34801561052e57600080fd5b506101bf6004803603604081101561054557600080fd5b506001600160a01b038135169060200135610f91565b34801561056757600080fd5b506101bf6004803603604081101561057e57600080fd5b506001600160a01b038135169060200135610fcd565b3480156105a057600080fd5b506102db600480360360608110156105b757600080fd5b508035906001600160a01b036020820135169060400135610fda565b3480156105df57600080fd5b506101bf600480360360608110156105f657600080fd5b8135919081019060408101602082013564010000000081111561061857600080fd5b82018360208201111561062a57600080fd5b8035906020019184602083028401116401000000008311171561064c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561069c57600080fd5b8201836020820111156106ae57600080fd5b803590602001918460208302840111640100000000831117156106d057600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061106c945050505050565b34801561071a57600080fd5b506102db6004803603602081101561073157600080fd5b503561120d565b34801561074457600080fd5b506102db6004803603608081101561075b57600080fd5b508035906020810135906001600160a01b03604082013516906060013561128e565b34801561078957600080fd5b506102db600480360360408110156107a057600080fd5b506001600160a01b038135169060200135611337565b6102db600480360360208110156107cc57600080fd5b50356113c0565b3480156107df57600080fd5b506102f2600480360360408110156107f657600080fd5b506001600160a01b03813581169160200135166113cf565b34801561081a57600080fd5b506102db6004803603602081101561083157600080fd5b50356001600160a01b03166113fa565b34801561084d57600080fd5b506102db6004803603604081101561086457600080fd5b506001600160a01b03813516906020013561144a565b600061088733848461151e565b5060015b92915050565b6000818152600460205260408120606091829190805b60038301548110156108e3578260030181815481106108c257fe5b600091825260209091206002600390920201015491909101906001016108a7565b60608260405190808252806020026020018201604052801561090f578160200160208202803883390190505b50905060608360405190808252806020026020018201604052801561093e578160200160208202803883390190505b5060038601546000945090915083905b80851015610a1057600087600301868154811061096757fe5b600091825260208220600260039092020190810154909250905b81811015610a0257825487516001600160a01b03909116908890879081106109a557fe5b60200260200101906001600160a01b031690816001600160a01b0316815250508260020181815481106109d457fe5b90600052602060002001548686815181106109eb57fe5b602090810291909101015260019485019401610981565b50506001909501945061094e565b50919650945050505050915091565b610a27610db4565b610a66576040805162461bcd60e51b8152602060048201819052602482015260008051602061208a833981519152604482015290519081900360640190fd5b610a7184838361160a565b610aac5760405162461bcd60e51b815260040180806020018281038252602e81526020018061205c602e913960400191505060405180910390fd5b610ab78383836117a9565b50505050565b6003545b90565b6000610ad1848484611891565b6001600160a01b038416600090815260026020908152604080832033808552925290912054610b11918691610b0c908663ffffffff6119d516565b61151e565b5060015b9392505050565b600081815260046020908152604091829020600181015483518181528184028101909301909352606092839283918015610b60578160200160208202803883390190505b50905060608151604051908082528060200260200182016040528015610b90578160200160208202803883390190505b50825190915060005b81811015610c3b57846001018181548110610bb057fe5b600091825260209091206002909102015484516001600160a01b0390911690859083908110610bdb57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050846001018181548110610c0a57fe5b906000526020600020906002020160010154838281518110610c2857fe5b6020908102919091010152600101610b99565b5091945092505050915091565b3360008181526002602090815260408083206001600160a01b03871684529091528120549091610887918590610b0c908663ffffffff611a3216565b6001600160a01b031660009081526001602052604090205490565b610ca7610db4565b610ce6576040805162461bcd60e51b8152602060048201819052602482015260008051602061208a833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008281526004602090815260408083206001600160a01b03871684526002810190925282205480610d6757600092505050610b15565b816003016001820381548110610d7957fe5b600091825260208083209683526001600390920290960101909452505060409091205415159392505050565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038416916323b872dd9160648083019260209291908290030181600087803b158015610e1a57600080fd5b505af1158015610e2e573d6000803e3d6000fd5b505050506040513d6020811015610e4457600080fd5b50610e5f90506001600160601b03193360601b1683836117a9565b5050565b610e7b6001600160601b03193360601b16838361160a565b610eb65760405162461bcd60e51b815260040180806020018281038252602e81526020018061205c602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b158015610f0557600080fd5b505af1158015610f19573d6000803e3d6000fd5b505050506040513d6020811015610ab757600080fd5b60008181526004602090815260408083206001600160a01b03861684529182905282205480610f635760009250505061088b565b816001016001820381548110610f7557fe5b9060005260206000209060020201600101549250505092915050565b3360008181526002602090815260408083206001600160a01b03871684529091528120549091610887918590610b0c908663ffffffff6119d516565b6000610887338484611891565b610fe2610db4565b611021576040805162461bcd60e51b8152602060048201819052602482015260008051602061208a833981519152604482015290519081900360640190fd5b61102c83838361160a565b6110675760405162461bcd60e51b815260040180806020018281038252602e81526020018061205c602e913960400191505060405180910390fd5b505050565b8151600090815b818110156111475784818151811061108757fe5b602002602001015160146015811061109b57fe5b1a60f81b6001600160f81b031916600160f81b14156110fc576110e8868683815181106110c457fe5b602002602001015160601c8684815181106110db57fe5b6020026020010151611a8c565b6110f757600092505050610b15565b61113f565b6111308686838151811061110c57fe5b602002602001015160601c86848151811061112357fe5b602002602001015161160a565b61113f57600092505050610b15565b600101611073565b5060005b818110156112015784818151811061115f57fe5b602002602001015160146015811061117357fe5b1a60f81b6001600160f81b031916600160f81b14156111c5576111c08686838151811061119c57fe5b602002602001015160601c8684815181106111b357fe5b6020026020010151611cf6565b6111f9565b6111f9868683815181106111d557fe5b602002602001015160601c8684815181106111ec57fe5b60200260200101516117a9565b60010161114b565b50600195945050505050565b6112266001600160601b03193360601b1660008361160a565b6112615760405162461bcd60e51b815260040180806020018281038252602e81526020018061205c602e913960400191505060405180910390fd5b604051339082156108fc029083906000818181858888f19350505050158015610e5f573d6000803e3d6000fd5b611296610db4565b6112d5576040805162461bcd60e51b8152602060048201819052602482015260008051602061208a833981519152604482015290519081900360640190fd5b6112e0848383611a8c565b61132c576040805162461bcd60e51b81526020600482015260186024820152772bb0b63632ba103237b2b9b713ba1037bbb7103a37b5b2b760411b604482015290519081900360640190fd5b610ab7838383611cf6565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038416916323b872dd91606480830192600092919082900301818387803b15801561138b57600080fd5b505af115801561139f573d6000803e3d6000fd5b50505050610e5f3360601b6bffffffffffffffffffffffff19168383611cf6565b6113cc816000346117a9565b50565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b611402610db4565b611441576040805162461bcd60e51b8152602060048201819052602482015260008051602061208a833981519152604482015290519081900360640190fd5b6113cc81611e8d565b6114626001600160601b03193360601b168383611a8c565b6114ae576040805162461bcd60e51b81526020600482015260186024820152772bb0b63632ba103237b2b9b713ba1037bbb7103a37b5b2b760411b604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b15801561150257600080fd5b505af1158015611516573d6000803e3d6000fd5b505050505050565b6001600160a01b0383166115635760405162461bcd60e51b81526004018080602001828103825260248152602001806120cf6024913960400191505060405180910390fd5b6001600160a01b0382166115a85760405162461bcd60e51b815260040180806020018281038252602281526020018061203a6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b60008161161957506001610b15565b60008481526004602090815260408083206001600160a01b038716845291829052909120548061164e57600092505050610b15565b600082600101600183038154811061166257fe5b90600052602060002090600202019050806001015485111561168a5760009350505050610b15565b600181015461169f908663ffffffff6119d516565b6001820181905561179c57600183018054839185916000919060001981019081106116c657fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061170457fe5b906000526020600020906002020183600101600184038154811061172457fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061177257fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b5060019695505050505050565b806117b357611067565b60008381526004602090815260408083206001600160a01b038616845291829052909120548061184957506040805180820182526001600160a01b0385811680835260006020808501828152600188810180548083018083559186528486209851600290910290980180546001600160a01b03191698909716979097178655905194019390935590815290849052919091208190555b600082600101600183038154811061185d57fe5b90600052602060002090600202019050611884848260010154611a3290919063ffffffff16565b6001909101555050505050565b6001600160a01b0383166118d65760405162461bcd60e51b81526004018080602001828103825260258152602001806120aa6025913960400191505060405180910390fd5b6001600160a01b03821661191b5760405162461bcd60e51b8152600401808060200182810382526023815260200180611ff16023913960400191505060405180910390fd5b6001600160a01b038316600090815260016020526040902054611944908263ffffffff6119d516565b6001600160a01b038085166000908152600160205260408082209390935590841681522054611979908263ffffffff611a3216565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082821115611a2c576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015610b15576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60008381526004602090815260408083206001600160a01b03861684526002810190925282205480611ac357600092505050610b15565b6000826003016001830381548110611ad757fe5b600091825260208083208884526001600390930201918201905260409091205490915080611b0c576000945050505050610b15565b60028201805482916001850191600091906000198101908110611b2b57fe5b600091825260208083209091015483528201929092526040019020556002820180546000198101908110611b5b57fe5b9060005260206000200154826002016001830381548110611b7857fe5b600091825260208083209091019290925587815260018401909152604081205560028201805480611ba557fe5b6000828152602081208201600019908101919091550190556002820154611ce85760038401805484916002870191600091906000198101908110611be557fe5b60009182526020808320600392830201546001600160a01b031684528301939093526040909101902091909155840180546000198101908110611c2457fe5b9060005260206000209060030201846003016001850381548110611c4457fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b0390921691909117815560028083018054611c879284019190611f2d565b5050506001600160a01b038716600090815260028501602052604081205560038401805480611cb257fe5b60008281526020812060036000199093019283020180546001600160a01b031916815590611ce36002830182611f7d565b505090555b506001979650505050505050565b60008381526004602090815260408083206001600160a01b0386168452600281019092529091205480611dc9576040805180820182526001600160a01b03861681528151600080825260208281019094526003860193830191905090528154600181018084556000938452602093849020835160039093020180546001600160a01b0319166001600160a01b039093169290921782558284015180519194611da692600285019290910190611f9b565b5050506001600160a01b0385166000908152600284016020526040902081905590505b6000826003016001830381548110611ddd57fe5b9060005260206000209060030201905080600101600085815260200190815260200160002054600014611e57576040805162461bcd60e51b815260206004820152601d60248201527f63616e27742061646420616c7265616479206f776e656420746f6b656e000000604482015290519081900360640190fd5b60028101805460018181018355600083815260208082209093018890559254968352909201909152604090209290925550505050565b6001600160a01b038116611ed25760405162461bcd60e51b81526004018080602001828103825260268152602001806120146026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054828255906000526020600020908101928215611f6d5760005260206000209182015b82811115611f6d578254825591600101919060010190611f52565b50611f79929150611fd6565b5090565b50805460008255906000526020600020908101906113cc9190611fd6565b828054828255906000526020600020908101928215611f6d579160200282015b82811115611f6d578251825591602001919060010190611fbb565b610ac191905b80821115611f795760008155600101611fdc56fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737357616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656e4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657245524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a265627a7a723058208e4f55278c10187d530c1533ae5abe33be090d661186edb7390afc01fa047eb364736f6c634300050a0032"

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
	var (
		ret0 = new(*big.Int)
	)
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
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbBalanceTracker.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ArbBalanceTracker.Contract.BalanceOf(&_ArbBalanceTracker.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ArbBalanceTracker *ArbBalanceTrackerCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ArbBalanceTracker.Contract.BalanceOf(&_ArbBalanceTracker.CallOpts, account)
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
	var (
		ret0 = new(*big.Int)
	)
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
	var (
		ret0 = new(bool)
	)
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
	var (
		ret0 = new(bool)
	)
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
	var (
		ret0 = new(common.Address)
	)
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
	var (
		ret0 = new(*big.Int)
	)
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
// Solidity: function hasFunds(bytes32 _vmId, bytes21[] _tokenTypes, uint256[] _amounts) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) HasFunds(opts *bind.TransactOpts, _vmId [32]byte, _tokenTypes [][21]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "hasFunds", _vmId, _tokenTypes, _amounts)
}

// HasFunds is a paid mutator transaction binding the contract method 0xc2465106.
//
// Solidity: function hasFunds(bytes32 _vmId, bytes21[] _tokenTypes, uint256[] _amounts) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) HasFunds(_vmId [32]byte, _tokenTypes [][21]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.HasFunds(&_ArbBalanceTracker.TransactOpts, _vmId, _tokenTypes, _amounts)
}

// HasFunds is a paid mutator transaction binding the contract method 0xc2465106.
//
// Solidity: function hasFunds(bytes32 _vmId, bytes21[] _tokenTypes, uint256[] _amounts) returns(bool)
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

// OwnerRemoveToken is a paid mutator transaction binding the contract method 0xb8569ccd.
//
// Solidity: function ownerRemoveToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) OwnerRemoveToken(opts *bind.TransactOpts, _user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "ownerRemoveToken", _user, _tokenContract, _value)
}

// OwnerRemoveToken is a paid mutator transaction binding the contract method 0xb8569ccd.
//
// Solidity: function ownerRemoveToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerSession) OwnerRemoveToken(_user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.OwnerRemoveToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _value)
}

// OwnerRemoveToken is a paid mutator transaction binding the contract method 0xb8569ccd.
//
// Solidity: function ownerRemoveToken(bytes32 _user, address _tokenContract, uint256 _value) returns()
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) OwnerRemoveToken(_user [32]byte, _tokenContract common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.OwnerRemoveToken(&_ArbBalanceTracker.TransactOpts, _user, _tokenContract, _value)
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
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.Transfer(&_ArbBalanceTracker.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.Transfer(&_ArbBalanceTracker.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferFrom(&_ArbBalanceTracker.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ArbBalanceTracker *ArbBalanceTrackerTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBalanceTracker.Contract.TransferFrom(&_ArbBalanceTracker.TransactOpts, sender, recipient, amount)
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) ParseApproval(log types.Log) (*ArbBalanceTrackerApproval, error) {
	event := new(ArbBalanceTrackerApproval)
	if err := _ArbBalanceTracker.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) ParseOwnershipTransferred(log types.Log) (*ArbBalanceTrackerOwnershipTransferred, error) {
	event := new(ArbBalanceTrackerOwnershipTransferred)
	if err := _ArbBalanceTracker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ArbBalanceTracker *ArbBalanceTrackerFilterer) ParseTransfer(log types.Log) (*ArbBalanceTrackerTransfer, error) {
	event := new(ArbBalanceTrackerTransfer)
	if err := _ArbBalanceTracker.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbProtocolABI is the input ABI used to generate the binding from.
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[5]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"unanimousAssertHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"name\":\"_lastLogHash\",\"type\":\"bytes32\"},{\"name\":\"_totalMessageValueAmounts\",\"type\":\"uint256[]\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_tokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_destinations\",\"type\":\"bytes32[]\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_dest\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"bytes32\"}],\"name\":\"generateSentMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"countSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_signatures\",\"type\":\"bytes\"},{\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"parseSignature\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_dataHashes\",\"type\":\"bytes32[]\"},{\"name\":\"_tokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_destinations\",\"type\":\"bytes32[]\"}],\"name\":\"generateLastMessageHashStub\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"bytes32\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"},{\"name\":\"_newMessage\",\"type\":\"bytes32\"}],\"name\":\"appendInboxPendingMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"recoverAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_inboxHash\",\"type\":\"bytes32\"},{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"}],\"name\":\"appendInboxMessages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"f11fcc26": "appendInboxMessages(bytes32,bytes32)",
	"d78d18ea": "appendInboxPendingMessage(bytes32,bytes32)",
	"0f89fbff": "calculateBeforeValues(bytes21[],uint16[],uint256[])",
	"33ae3ad0": "countSignatures(bytes)",
	"20903721": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32,uint256[])",
	"25200160": "generateLastMessageHash(bytes21[],bytes,uint16[],uint256[],bytes32[])",
	"b3277495": "generateLastMessageHashStub(bytes21[],bytes32[],uint16[],uint256[],bytes32[])",
	"ccf69dd7": "generateMessageStubHash(bytes32,bytes21,uint256,bytes32)",
	"3e285598": "generatePreconditionHash(bytes32,uint64[2],bytes32,bytes21[],uint256[])",
	"2a0500d8": "generateSentMessageHash(bytes32,bytes32,bytes21,uint256,bytes32)",
	"b31d63cc": "parseSignature(bytes,uint256)",
	"c655d7aa": "recoverAddress(bytes32,bytes)",
	"f0c8e969": "recoverAddresses(bytes32,bytes)",
	"014bba5b": "unanimousAssertHash(bytes32[5],uint64[2],bytes21[],bytes,uint16[],uint256[],bytes32[])",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x61254d610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100f45760003560e01c8063b31d63cc11610096578063ccf69dd711610070578063ccf69dd714610ff9578063d78d18ea14611032578063f0c8e96914611055578063f11fcc2614611100576100f4565b8063b31d63cc14610bc1578063b327749514610c89578063c655d7aa14610f32576100f4565b806325200160116100d257806325200160146106d45780632a0500d81461098057806333ae3ad0146109bf5780633e28559814610a63576100f4565b8063014bba5b146100f95780630f89fbff146104145780632090372114610609575b600080fd5b610402600480360361018081101561011057600080fd5b810190808060a001906005806020026040519081016040528092919082600560200280828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561018657600080fd5b82018360208201111561019857600080fd5b803590602001918460208302840111600160201b831117156101b957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561020857600080fd5b82018360208201111561021a57600080fd5b803590602001918460018302840111600160201b8311171561023b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561028d57600080fd5b82018360208201111561029f57600080fd5b803590602001918460208302840111600160201b831117156102c057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561030f57600080fd5b82018360208201111561032157600080fd5b803590602001918460208302840111600160201b8311171561034257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561039157600080fd5b8201836020820111156103a357600080fd5b803590602001918460208302840111600160201b831117156103c457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611123945050505050565b60408051918252519081900360200190f35b6105b96004803603606081101561042a57600080fd5b810190602081018135600160201b81111561044457600080fd5b82018360208201111561045657600080fd5b803590602001918460208302840111600160201b8311171561047757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104c657600080fd5b8201836020820111156104d857600080fd5b803590602001918460208302840111600160201b831117156104f957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561054857600080fd5b82018360208201111561055a57600080fd5b803590602001918460208302840111600160201b8311171561057b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506112bc945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156105f55781810151838201526020016105dd565b505050509050019250505060405180910390f35b610402600480360360e081101561061f57600080fd5b81359163ffffffff6020820135169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561066357600080fd5b82018360208201111561067557600080fd5b803590602001918460208302840111600160201b8311171561069657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506114d1945050505050565b610402600480360360a08110156106ea57600080fd5b810190602081018135600160201b81111561070457600080fd5b82018360208201111561071657600080fd5b803590602001918460208302840111600160201b8311171561073757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561078657600080fd5b82018360208201111561079857600080fd5b803590602001918460018302840111600160201b831117156107b957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561080b57600080fd5b82018360208201111561081d57600080fd5b803590602001918460208302840111600160201b8311171561083e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561088d57600080fd5b82018360208201111561089f57600080fd5b803590602001918460208302840111600160201b831117156108c057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561090f57600080fd5b82018360208201111561092157600080fd5b803590602001918460208302840111600160201b8311171561094257600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061153c945050505050565b610402600480360360a081101561099657600080fd5b508035906020810135906001600160581b0319604082013516906060810135906080013561177f565b610402600480360360208110156109d557600080fd5b810190602081018135600160201b8111156109ef57600080fd5b820183602082011115610a0157600080fd5b803590602001918460018302840111600160201b83111715610a2257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061195d945050505050565b610402600480360360c0811015610a7957600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594909390925060408101915060200135600160201b811115610ace57600080fd5b820183602082011115610ae057600080fd5b803590602001918460208302840111600160201b83111715610b0157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b5057600080fd5b820183602082011115610b6257600080fd5b803590602001918460208302840111600160201b83111715610b8357600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061198c945050505050565b610c6760048036036040811015610bd757600080fd5b810190602081018135600160201b811115610bf157600080fd5b820183602082011115610c0357600080fd5b803590602001918460018302840111600160201b83111715610c2457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250611a77915050565b6040805160ff9094168452602084019290925282820152519081900360600190f35b610402600480360360a0811015610c9f57600080fd5b810190602081018135600160201b811115610cb957600080fd5b820183602082011115610ccb57600080fd5b803590602001918460208302840111600160201b83111715610cec57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d3b57600080fd5b820183602082011115610d4d57600080fd5b803590602001918460208302840111600160201b83111715610d6e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610dbd57600080fd5b820183602082011115610dcf57600080fd5b803590602001918460208302840111600160201b83111715610df057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e3f57600080fd5b820183602082011115610e5157600080fd5b803590602001918460208302840111600160201b83111715610e7257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610ec157600080fd5b820183602082011115610ed357600080fd5b803590602001918460208302840111600160201b83111715610ef457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611b05945050505050565b610fdd60048036036040811015610f4857600080fd5b81359190810190604081016020820135600160201b811115610f6957600080fd5b820183602082011115610f7b57600080fd5b803590602001918460018302840111600160201b83111715610f9c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611ca3945050505050565b604080516001600160a01b039092168252519081900360200190f35b6104026004803603608081101561100f57600080fd5b508035906001600160581b03196020820135169060408101359060600135611dd6565b6104026004803603604081101561104857600080fd5b5080359060200135611eb4565b6105b96004803603604081101561106b57600080fd5b81359190810190604081016020820135600160201b81111561108c57600080fd5b82018360208201111561109e57600080fd5b803590602001918460018302840111600160201b831117156110bf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611ef8945050505050565b6104026004803603604081101561111657600080fd5b508035906020013561209e565b6000878787878787876040516020018088600560200280838360005b8381101561115757818101518382015260200161113f565b5050505090500187600260200280838360005b8381101561118257818101518382015260200161116a565b50505050905001868051906020019060200280838360005b838110156111b257818101518382015260200161119a565b5050505090500185805190602001908083835b602083106111e45780518252601f1990920191602091820191016111c5565b51815160209384036101000a60001901801990921691161790528751919093019287810192500280838360005b83811015611229578181015183820152602001611211565b50505050905001838051906020019060200280838360005b83811015611259578181015183820152602001611241565b50505050905001828051906020019060200280838360005b83811015611289578181015183820152602001611271565b50505050905001975050505050505050604051602081830303815290604052805190602001209050979650505050505050565b60608084516040519080825280602002602001820160405280156112ea578160200160208202803883390190505b50845190915060005b818110156114c6578686828151811061130857fe5b602002602001015161ffff168151811061131e57fe5b602002602001015160146015811061133257fe5b1a60f81b6001600160f81b0319166113925784818151811061135057fe5b60200260200101518387838151811061136557fe5b602002602001015161ffff168151811061137b57fe5b6020026020010181815101915081815250506114be565b8286828151811061139f57fe5b602002602001015161ffff16815181106113b557fe5b6020026020010151600014611411576040805162461bcd60e51b815260206004820152601d60248201527f43616e277420696e636c756465204e465420746f6b656e207477696365000000604482015290519081900360640190fd5b84818151811061141d57fe5b60200260200101516000141561147a576040805162461bcd60e51b815260206004820152601f60248201527f4e465420746f6b656e206d7573742068617665206e6f6e2d7a65726f20696400604482015290519081900360640190fd5b84818151811061148657fe5b60200260200101518387838151811061149b57fe5b602002602001015161ffff16815181106114b157fe5b6020026020010181815250505b6001016112f3565b509095945050505050565b600087878787878787604051602001808881526020018763ffffffff1663ffffffff1660e01b81526004018681526020018581526020018481526020018381526020018280519060200190602002808383600083811015611289578181015183820152602001611271565b6000815183511461158a576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b83518351146115d6576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b825160009081908190815b818110156117715773__$0d86abb4a722a612872fb80f4c7e7e95bd$__6389df40da8b866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561165757818101518382015260200161163f565b50505050905090810190601f1680156116845780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b1580156116a157600080fd5b505af41580156116b5573d6000803e3d6000fd5b505050506040513d60408110156116cb57600080fd5b5080516020909101518a51919550935061173a9084908d908c90859081106116ef57fe5b602002602001015161ffff168151811061170557fe5b60200260200101518a848151811061171957fe5b60200260200101518a858151811061172d57fe5b6020026020010151611dd6565b60408051602080820198909852808201839052815180820383018152606090910190915280519601959095209492506001016115e1565b505050505095945050505050565b60408051602080820188905281830187905260608083018690526001600160581b03198716608084015283518084036075018152609584018086528151919093012060048084526101358501909552600094909391929160b5015b6117e26124e2565b8152602001906001900390816117da579050509050611800876120b8565b8160008151811061180d57fe5b602002602001018190525061182142612112565b8160018151811061182e57fe5b602002602001018190525061184243612112565b8160028151811061184f57fe5b602090810291909101015261186382612112565b8160038151811061187057fe5b602090810291909101015260408051600480825260a08201909252606091816020015b61189b6124e2565b8152602001906001900390816118935790505090506118b98261216c565b816000815181106118c657fe5b60209081029190910101526118da85612112565b816001815181106118e757fe5b60200260200101819052506118fb86612112565b8160028151811061190857fe5b60209081029190910101526119266001600160581b03198816612112565b8160038151811061193357fe5b602002602001018190525061194f61194a8261216c565b6121f4565b519998505050505050505050565b6000604182518161196a57fe5b0615611977576000611984565b604182518161198257fe5b045b90505b919050565b600085858260200201518660016020020151868686604051602001808781526020018667ffffffffffffffff1667ffffffffffffffff1660c01b81526008018567ffffffffffffffff1667ffffffffffffffff1660c01b8152600801848152602001838051906020019060200280838360005b83811015611a175781810151838201526020016119ff565b50505050905001828051906020019060200280838360005b83811015611a47578181015183820152602001611a2f565b50505050905001965050505050505060405160208183030381529060405280519060200120905095945050505050565b604180820283810160208101516040820151919093015160ff169291601b841015611aa357601b840193505b8360ff16601b1480611ab857508360ff16601c145b611afd576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b60008351855114611b53576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8251855114611b9f576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8151855114611beb576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b84516000908190815b81811015611c9557611c5e898281518110611c0b57fe5b60200260200101518b8a8481518110611c2057fe5b602002602001015161ffff1681518110611c3657fe5b6020026020010151898481518110611c4a57fe5b602002602001015189858151811061172d57fe5b6040805160208082019790975280820183905281518082038301815260609091019091528051950194909420939250600101611bf4565b509198975050505050505050565b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081886040516020018083805190602001908083835b60208310611d195780518252601f199092019160209182019101611cfa565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201905282519201919091209250611d6091508890506000611a77565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa158015611dbf573d6000803e3d6000fd5b5050604051601f1901519998505050505050505050565b60408051600480825260a0820190925260009160609190816020015b611dfa6124e2565b815260200190600190039081611df2579050509050611e18866120b8565b81600081518110611e2557fe5b6020908102919091010152611e3983612112565b81600181518110611e4657fe5b6020026020010181905250611e5a84612112565b81600281518110611e6757fe5b6020908102919091010152611e856001600160581b03198616612112565b81600381518110611e9257fe5b6020026020010181905250611ea961194a8261216c565b519695505050505050565b6000611ef16040518060600160405280611ece6000612112565b8152602001611edc866120b8565b8152602001611eea856120b8565b90526122e3565b9392505050565b6060600080600080611f098661195d565b9050606081604051908082528060200260200182016040528015611f37578160200160208202803883390190505b50905060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818a6040516020018083805190602001908083835b60208310611faa5780518252601f199092019160209182019101611f8b565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190528251920191909120925060009150505b8481101561208f57611ffa8a82611a77565b6040805160008152602080820180845288905260ff86168284015260608201859052608082018490529151949c50929a5090985060019260a080840193601f198301929081900390910190855afa158015612059573d6000803e3d6000fd5b5050506020604051035184828151811061206f57fe5b6001600160a01b0390921660209283029190910190910152600101611fe8565b50919998505050505050505050565b6000611ef16040518060600160405280611ece6001612112565b6120c06124e2565b604080516060810182528381528151600080825260208281019094529192830191612101565b6120ee6124e2565b8152602001906001900390816120e65790505b508152600260209091015292915050565b61211a6124e2565b60408051606081018252838152815160008082526020828101909452919283019161215b565b6121486124e2565b8152602001906001900390816121405790505b508152600060209091015292915050565b6121746124e2565b61217e825161236b565b6121cf576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b6121fc612506565b6040820151600c60ff9091161061224e576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff1661227b5760405180602001604052806122728460000151612372565b90529050611987565b604082015160ff16600214156122a05750604080516020810190915281518152611987565b600360ff16826040015160ff16101580156122c457506040820151600c60ff909116105b156122e15760405180602001604052806122728460200151612396565bfe5b6040805160038082526080820190925260009160609190816020015b6123076124e2565b8152602001906001900390816122ff575050805190915060005b818110156123595784816003811061233557fe5b602002015183828151811061234657fe5b6020908102919091010152600101612321565b5061236382612396565b949350505050565b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b60006008825111156123e6576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015612413578160200160208202803883390190505b50805190915060005b8181101561246f5761242c612506565b61244886838151811061243b57fe5b60200260200101516121f4565b9050806000015184838151811061245b57fe5b60209081029190910101525060010161241c565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156124b85781810151838201526020016124a0565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820c074231ea266da7a9e28881bd4ba549da90858a6918f9f96e58cba335741c0bc64736f6c634300050a0032"

// DeployArbProtocol deploys a new Ethereum contract, binding an instance of ArbProtocol to it.
func DeployArbProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbProtocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbProtocolBin = strings.Replace(ArbProtocolBin, "__$0d86abb4a722a612872fb80f4c7e7e95bd$__", arbValueAddr.String()[2:], -1)

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
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new([]*big.Int)
	)
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
	var (
		ret0 = new(*big.Int)
	)
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
	var (
		ret0 = new([32]byte)
	)
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
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHash", _tokenTypes, _data, _tokenNums, _amounts, _destinations)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _data, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _data, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHashStub(opts *bind.CallOpts, _tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHashStub", _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
	return *ret0, err
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, bytes32[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new([32]byte)
	)
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

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_ArbProtocol *ArbProtocolCaller) RecoverAddress(opts *bind.CallOpts, _messageHash [32]byte, _signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "recoverAddress", _messageHash, _signature)
	return *ret0, err
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_ArbProtocol *ArbProtocolSession) RecoverAddress(_messageHash [32]byte, _signature []byte) (common.Address, error) {
	return _ArbProtocol.Contract.RecoverAddress(&_ArbProtocol.CallOpts, _messageHash, _signature)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_ArbProtocol *ArbProtocolCallerSession) RecoverAddress(_messageHash [32]byte, _signature []byte) (common.Address, error) {
	return _ArbProtocol.Contract.RecoverAddress(&_ArbProtocol.CallOpts, _messageHash, _signature)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_ArbProtocol *ArbProtocolCaller) RecoverAddresses(opts *bind.CallOpts, _messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
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
	var (
		ret0 = new([32]byte)
	)
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
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValidValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidValueHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeValueHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueFuncSigs maps the 4-byte function signature to its string representation.
var ArbValueFuncSigs = map[string]string{
	"89df40da": "deserializeValidValueHash(bytes,uint256)",
	"8f346036": "deserializeValueHash(bytes)",
	"1f3d4d4e": "getNextValidValue(bytes,uint256)",
	"53409fab": "hashCodePointBasicValue(uint8,bytes32)",
	"264f384b": "hashCodePointImmediateValue(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"1667b411": "hashIntValue(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ArbValueBin is the compiled bytecode used for deploying new contracts.
var ArbValueBin = "0x610d71610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c806353409fab1161006557806353409fab1461022157806389df40da146102475780638f34603614610308578063b2b9dc62146103ae57610092565b80631667b411146100975780631f3d4d4e146100c6578063264f384b146101ed578063364df27714610219575b600080fd5b6100b4600480360360208110156100ad57600080fd5b50356103df565b60408051918252519081900360200190f35b61016e600480360360408110156100dc57600080fd5b8101906020810181356401000000008111156100f757600080fd5b82018360208201111561010957600080fd5b8035906020019184600183028401116401000000008311171561012b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610405915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101b1578181015183820152602001610199565b50505050905090810190601f1680156101de5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100b46004803603606081101561020357600080fd5b5060ff813516906020810135906040013561049a565b6100b46104ec565b6100b46004803603604081101561023757600080fd5b5060ff813516906020013561055f565b6102ef6004803603604081101561025d57600080fd5b81019060208101813564010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506105a6915050565b6040805192835260208301919091528051918290030190f35b6100b46004803603602081101561031e57600080fd5b81019060208101813564010000000081111561033957600080fd5b82018360208201111561034b57600080fd5b8035906020019184600183028401116401000000008311171561036d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610631945050505050565b6103cb600480360360208110156103c457600080fd5b50356106b5565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b60006060600080610414610d06565b61041e87876106bc565b919450925090508215610478576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161048c888880840363ffffffff61081116565b945094505050509250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b83811015610538578181015183820152602001610520565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000806000806105b4610d06565b6105be87876106bc565b919450925090508215610618576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161062282610891565b51909890975095505050505050565b6000808061063d610d06565b6106488560006106bc565b9194509250905082156106a2576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b6106ab81610891565b5195945050505050565b6008101590565b6000806106c7610d06565b8451841061071c576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b6000849050600086828151811061072f57fe5b016020015160019092019160f81c9050600081610771576107508884610980565b9093509050600083610761836109a7565b9197509550935061080a92505050565b60ff821660021415610798576107878884610980565b909350905060008361076183610a01565b600360ff8316108015906107af5750600c60ff8316105b156107eb576002198201606060006107c8838c88610a5b565b9097509250905080866107da84610b16565b98509850985050505050505061080a565b8160ff166127100160006107ff60006109a7565b919750955093505050505b9250925092565b60608183018451101561082357600080fd5b60608215801561083e57604051915060208201604052610888565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561087757805183526020928301920161085f565b5050858452601f01601f1916604052505b50949350505050565b610899610d2a565b6040820151600c60ff909116106108eb576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff1661091857604051806020016040528061090f84600001516103df565b90529050610400565b604082015160ff166002141561093d5750604080516020810190915281518152610400565b600360ff16826040015160ff161015801561096157506040820151600c60ff909116105b1561097e57604051806020016040528061090f8460200151610b9e565bfe5b6000808281610995868363ffffffff610cea16565b60209290920196919550909350505050565b6109af610d06565b6040805160608101825283815281516000808252602082810190945291928301916109f0565b6109dd610d06565b8152602001906001900390816109d55790505b508152600060209091015292915050565b610a09610d06565b604080516060810182528381528151600080825260208281019094529192830191610a4a565b610a37610d06565b815260200190600190039081610a2f5790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610aa657816020015b610a93610d06565b815260200190600190039081610a8b5790505b50905060005b8960ff168160ff161015610b0057610ac489856106bc565b8451859060ff8616908110610ad557fe5b6020908102919091010152945092508215610af857509094509092509050610b0d565b600101610aac565b5060009550919350909150505b93509350939050565b610b1e610d06565b610b2882516106b5565b610b79576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b6000600882511115610bee576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610c1b578160200160208202803883390190505b50805190915060005b81811015610c7757610c34610d2a565b610c50868381518110610c4357fe5b6020026020010151610891565b90508060000151848381518110610c6357fe5b602090810291909101015250600101610c24565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610cc0578181015183820152602001610ca8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60008160200183511015610cfd57600080fd5b50016020015190565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820e580d18e866df844b2ca541adbb02d9e0adb42ede410d1908ac1a96124e1d94b64736f6c634300050a0032"

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

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValidValueHash(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "deserializeValidValueHash", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValueHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "deserializeValueHash", data)
	return *ret0, err
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCaller) GetNextValidValue(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "getNextValidValue", data, offset)
	return *ret0, *ret1, err
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCallerSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointBasicValue(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new(bool)
	)
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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058202a4b33a7382bd41fc01a7b408d0a3491e17cc0d680148cb3a692e8a4446bc59d64736f6c634300050a0032"

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

// CountersABI is the input ABI used to generate the binding from.
const CountersABI = "[]"

// CountersBin is the compiled bytecode used for deploying new contracts.
var CountersBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820c959554fe0eae3c3a7f3befe11eaf67bb1ed6d3669a06ace5705566c0dd04fd164736f6c634300050a0032"

// DeployCounters deploys a new Ethereum contract, binding an instance of Counters to it.
func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counters, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CountersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// Counters is an auto generated Go binding around an Ethereum contract.
type Counters struct {
	CountersCaller     // Read-only binding to the contract
	CountersTransactor // Write-only binding to the contract
	CountersFilterer   // Log filterer for contract events
}

// CountersCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountersSession struct {
	Contract     *Counters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountersCallerSession struct {
	Contract *CountersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CountersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountersTransactorSession struct {
	Contract     *CountersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CountersRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountersRaw struct {
	Contract *Counters // Generic contract binding to access the raw methods on
}

// CountersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountersCallerRaw struct {
	Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
}

// CountersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountersTransactorRaw struct {
	Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounters creates a new instance of Counters, bound to a specific deployed contract.
func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
	contract, err := bindCounters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
	contract, err := bindCounters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountersCaller{contract: contract}, nil
}

// NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
	contract, err := bindCounters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountersTransactor{contract: contract}, nil
}

// NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
	contract, err := bindCounters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountersFilterer{contract: contract}, nil
}

// bindCounters binds a generic wrapper to an already deployed contract.
func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transact(opts, method, params...)
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ERC165FuncSigs maps the 4-byte function signature to its string representation.
var ERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
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
	var (
		ret0 = new(bool)
	)
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
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x608060405234801561001057600080fd5b506106e2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a457c2d71461016f578063a9059cbb1461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063395093511461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d561020b565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610211565b6100b96004803603604081101561013357600080fd5b506001600160a01b038135169060200135610268565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166102a4565b6100b96004803603604081101561018557600080fd5b506001600160a01b0381351690602001356102bf565b6100b9600480360360408110156101b157600080fd5b506001600160a01b0381351690602001356102fb565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610308565b6000610202338484610333565b50600192915050565b60025490565b600061021e84848461041f565b6001600160a01b03841660009081526001602090815260408083203380855292529091205461025e918691610259908663ffffffff61056116565b610333565b5060019392505050565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610202918590610259908663ffffffff6105be16565b6001600160a01b031660009081526020819052604090205490565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610202918590610259908663ffffffff61056116565b600061020233848461041f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103785760405162461bcd60e51b815260040180806020018281038252602481526020018061068a6024913960400191505060405180910390fd5b6001600160a01b0382166103bd5760405162461bcd60e51b81526004018080602001828103825260228152602001806106436022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166104645760405162461bcd60e51b81526004018080602001828103825260258152602001806106656025913960400191505060405180910390fd5b6001600160a01b0382166104a95760405162461bcd60e51b81526004018080602001828103825260238152602001806106206023913960400191505060405180910390fd5b6001600160a01b0383166000908152602081905260409020546104d2908263ffffffff61056116565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610507908263ffffffff6105be16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156105b8576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015610618576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a265627a7a7230582047b0c477b8631fd21167438956a7e560d19f088b2cf24f1d155042bb2f24189c64736f6c634300050a0032"

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
	var (
		ret0 = new(*big.Int)
	)
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
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
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
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721FuncSigs maps the 4-byte function signature to its string representation.
var ERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC721Bin is the compiled bytecode used for deploying new contracts.
var ERC721Bin = "0x608060405234801561001057600080fd5b506100437f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0361007a16565b6100757f80ac58cd000000000000000000000000000000000000000000000000000000006001600160e01b0361007a16565b610148565b7fffffffff00000000000000000000000000000000000000000000000000000000808216141561010b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610d23806101576000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636352211e116100665780636352211e146101b157806370a08231146101ce578063a22cb46514610206578063b88d4fde14610234578063e985e9c5146102fa5761009e565b806301ffc9a7146100a3578063081812fc146100de578063095ea7b31461011757806323b872dd1461014557806342842e0e1461017b575b600080fd5b6100ca600480360360208110156100b957600080fd5b50356001600160e01b031916610328565b604080519115158252519081900360200190f35b6100fb600480360360208110156100f457600080fd5b5035610347565b604080516001600160a01b039092168252519081900360200190f35b6101436004803603604081101561012d57600080fd5b506001600160a01b0381351690602001356103a9565b005b6101436004803603606081101561015b57600080fd5b506001600160a01b038135811691602081013590911690604001356104ba565b6101436004803603606081101561019157600080fd5b506001600160a01b0381358116916020810135909116906040013561050f565b6100fb600480360360208110156101c757600080fd5b503561052a565b6101f4600480360360208110156101e457600080fd5b50356001600160a01b0316610584565b60408051918252519081900360200190f35b6101436004803603604081101561021c57600080fd5b506001600160a01b03813516906020013515156105ec565b6101436004803603608081101561024a57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561028557600080fd5b82018360208201111561029757600080fd5b803590602001918460018302840111640100000000831117156102b957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106b8945050505050565b6100ca6004803603604081101561031057600080fd5b506001600160a01b0381358116916020013516610710565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103528261073e565b61038d5760405162461bcd60e51b815260040180806020018281038252602c815260200180610c48602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006103b48261052a565b9050806001600160a01b0316836001600160a01b031614156104075760405162461bcd60e51b8152600401808060200182810382526021815260200180610c9d6021913960400191505060405180910390fd5b336001600160a01b038216148061042357506104238133610710565b61045e5760405162461bcd60e51b8152600401808060200182810382526038815260200180610bbd6038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6104c4338261075b565b6104ff5760405162461bcd60e51b8152600401808060200182810382526031815260200180610cbe6031913960400191505060405180910390fd5b61050a8383836107ff565b505050565b61050a838383604051806020016040528060008152506106b8565b6000818152600160205260408120546001600160a01b03168061057e5760405162461bcd60e51b8152600401808060200182810382526029815260200180610c1f6029913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166105cb5760405162461bcd60e51b815260040180806020018281038252602a815260200180610bf5602a913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061057e90610943565b6001600160a01b03821633141561064a576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6106c38484846104ba565b6106cf84848484610947565b61070a5760405162461bcd60e51b8152600401808060200182810382526032815260200180610b3b6032913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b60006107668261073e565b6107a15760405162461bcd60e51b815260040180806020018281038252602c815260200180610b91602c913960400191505060405180910390fd5b60006107ac8361052a565b9050806001600160a01b0316846001600160a01b031614806107e75750836001600160a01b03166107dc84610347565b6001600160a01b0316145b806107f757506107f78185610710565b949350505050565b826001600160a01b03166108128261052a565b6001600160a01b0316146108575760405162461bcd60e51b8152600401808060200182810382526029815260200180610c746029913960400191505060405180910390fd5b6001600160a01b03821661089c5760405162461bcd60e51b8152600401808060200182810382526024815260200180610b6d6024913960400191505060405180910390fd5b6108a581610a7a565b6001600160a01b03831660009081526003602052604090206108c690610ab7565b6001600160a01b03821660009081526003602052604090206108e790610ace565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b5490565b600061095b846001600160a01b0316610ad7565b610967575060016107f7565b604051630a85bd0160e11b815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a169463150b7a029490938c938b938b939260a4019060208501908083838e5b838110156109e15781810151838201526020016109c9565b50505050905090810190601f168015610a0e5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610a3057600080fd5b505af1158015610a44573d6000803e3d6000fd5b505050506040513d6020811015610a5a57600080fd5b50516001600160e01b031916630a85bd0160e11b14915050949350505050565b6000818152600260205260409020546001600160a01b031615610ab457600081815260026020526040902080546001600160a01b03191690555b50565b8054610aca90600163ffffffff610add16565b9055565b80546001019055565b3b151590565b600082821115610b34576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a265627a7a7230582011643b83e464bd4aee7c4e039c3d289b01bf031588921bd56dd528f23df5a26164736f6c634300050a0032"

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
	var (
		ret0 = new(*big.Int)
	)
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
	var (
		ret0 = new(common.Address)
	)
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
	var (
		ret0 = new(bool)
	)
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
	var (
		ret0 = new(common.Address)
	)
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
	var (
		ret0 = new(bool)
	)
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

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xbbda45ac.
//
// Solidity: function safeTransferFrom0(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xbbda45ac.
//
// Solidity: function safeTransferFrom0(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xbbda45ac.
//
// Solidity: function safeTransferFrom0(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
	event := new(ERC721ApprovalForAll)
	if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
const IChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"vmId\",\"type\":\"bytes32\"},{\"name\":\"players\",\"type\":\"address[2]\"},{\"name\":\"escrows\",\"type\":\"uint128[2]\"},{\"name\":\"challengePeriod\",\"type\":\"uint32\"},{\"name\":\"challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeManagerFuncSigs = map[string]string{
	"2b50d42b": "initiateChallenge(bytes32,address[2],uint128[2],uint32,bytes32)",
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
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
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
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
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
// Solidity: function balanceOf(address account) constant returns(uint256)
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
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
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
const IERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

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
	"01ffc9a7": "supportsInterface(bytes4)",
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
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
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
	var (
		ret0 = new(common.Address)
	)
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
	var (
		ret0 = new(bool)
	)
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
	var (
		ret0 = new(common.Address)
	)
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
	var (
		ret0 = new(bool)
	)
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

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xbbda45ac.
//
// Solidity: function safeTransferFrom0(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xbbda45ac.
//
// Solidity: function safeTransferFrom0(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xbbda45ac.
//
// Solidity: function safeTransferFrom0(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ReceiverABI is the input ABI used to generate the binding from.
const IERC721ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721ReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IERC721ReceiverFuncSigs = map[string]string{
	"150b7a02": "onERC721Received(address,address,uint256,bytes)",
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

// MerkleLibFuncSigs maps the 4-byte function signature to its string representation.
var MerkleLibFuncSigs = map[string]string{
	"6a2dda67": "generateAddressRoot(address[])",
	"9898dc10": "generateRoot(bytes32[])",
	"b792d767": "verifyProof(bytes,bytes32,bytes32,uint256)",
}

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
var MerkleLibBin = "0x610575610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043f565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b6000815b600181511115610422576060600282516001018161031f57fe5b04604051908082528060200260200182016040528015610349578160200160208202803883390190505b50905060005b815181101561041a5782518160020260010110156103e25782816002028151811061037657fe5b602002602001015183826002026001018151811061039057fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d157fe5b602002602001018181525050610412565b8281600202815181106103f157fe5b602002602001015182828151811061040557fe5b6020026020010181815250505b60010161034f565b509050610305565b8060008151811061042f57fe5b6020026020010151915050919050565b600080838160205b88518111610532578089015193506020818a51036020018161046557fe5b0491505b60008211801561047c5750600286066001145b801561048a57508160020a86115b1561049d57600286046001019550610469565b600286066104e85783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104e057fe5b04955061052a565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052357fe5b0460010195505b602001610447565b50509094149594505050505056fea265627a7a723058209d7af0b10f14fb3c4510a87a685097ce94041ed1d974962866d79395034b110064736f6c634300050a0032"

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
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new([32]byte)
	)
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
	var (
		ret0 = new(bool)
	)
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

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
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
	var (
		ret0 = new(bool)
	)
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
	var (
		ret0 = new(common.Address)
	)
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058203ff4034c6ea4209ece9da4562199bebb6f4a5e4818ff91583cbf2d23590e03bd64736f6c634300050a0032"

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
const VMTrackerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_newInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"},{\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"finalizedUnanimousAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_challengeManager\",\"type\":\"address\"}],\"name\":\"addChallengeManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_unanRest\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_sequenceNum\",\"type\":\"uint64\"},{\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"pendingUnanimousAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"withinTimeBounds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"},{\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_newInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"confirmUnanimousAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[3]\"},{\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"name\":\"_challengeManagerNum\",\"type\":\"uint16\"},{\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"name\":\"_escrowCurrency\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"createVm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[5]\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageDataHash\",\"type\":\"bytes32[]\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_msgAmount\",\"type\":\"uint256[]\"},{\"name\":\"_msgDestination\",\"type\":\"bytes32[]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"}],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_balanceTrackerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"tokenType\",\"type\":\"bytes21\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"_escrowCurrency\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_challengeManagerNum\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"VMCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"unanHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"PendingUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"ConfirmedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"unanHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"fields\",\"type\":\"bytes32[3]\"},{\"indexed\":false,\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"indexed\":false,\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"lastMessageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"logsAccHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// VMTrackerFuncSigs maps the 4-byte function signature to its string representation.
var VMTrackerFuncSigs = map[string]string{
	"2e6bf2e6": "addChallengeManager(address)",
	"63d84637": "completeChallenge(bytes32,address[2],uint128[2])",
	"6f0edcd5": "confirmDisputableAsserted(bytes32,bytes32,bytes32,uint32,bytes21[],bytes,uint16[],uint256[],bytes32[],bytes32)",
	"84481bab": "confirmUnanimousAsserted(bytes32,bytes32,bytes32,bytes21[],bytes,uint16[],uint256[],bytes32[])",
	"8dbc44f6": "createVm(bytes32[3],uint32,uint32,uint16,uint128,address,address,bytes)",
	"0725397e": "finalizedUnanimousAssert(bytes32,bytes32,bytes32,bytes21[],bytes,uint16[],uint256[],bytes32[],bytes32,bytes)",
	"d37a1a95": "forwardMessage(bytes32,bytes21,uint256,bytes,bytes)",
	"ffd134c9": "initiateChallenge(bytes32,bytes32)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"fd0961b6": "ownerShutdown(bytes32)",
	"da9f547a": "pendingDisputableAssert(bytes32[5],uint32,uint64[2],bytes21[],bytes32[],uint16[],uint256[],bytes32[])",
	"3f2832cd": "pendingUnanimousAssert(bytes32,bytes32,bytes21[],uint16[],uint256[],uint64,bytes32,bytes)",
	"715018a6": "renounceOwnership()",
	"fa3eebe6": "sendEthMessage(bytes32,bytes)",
	"3cd8a322": "sendMessage(bytes32,bytes21,uint256,bytes)",
	"f2fde38b": "transferOwnership(address)",
	"42c0787e": "withinTimeBounds(uint64[2])",
}

// VMTrackerBin is the compiled bytecode used for deploying new contracts.
var VMTrackerBin = "0x60806040523480156200001157600080fd5b5060405162005e7838038062005e78833981810160405260208110156200003757600080fd5b5051600080546001600160a01b03191633178082556040516001600160a01b039190911691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3600280546001600160a01b039092166001600160a01b03199092169190911790556000805260046020527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec805460ff19166001179055615d9080620000e86000396000f3fe6080604052600436106101095760003560e01c80638da5cb5b11610095578063da9f547a11610064578063da9f547a146110e5578063f2fde38b14611409578063fa3eebe61461143c578063fd0961b6146114e7578063ffd134c91461151157610109565b80638da5cb5b14610e205780638dbc44f614610e515780638f32d59b14610f7d578063d37a1a9514610f9257610109565b806342c0787e116100dc57806342c0787e146107c457806363d84637146108305780636f0edcd514610864578063715018a614610b3e57806384481bab14610b5357610109565b80630725397e1461010e5780632e6bf2e61461046a5780633cd8a3221461049d5780633f2832cd1461056b575b600080fd5b34801561011a57600080fd5b50610468600480360361014081101561013257600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561015f57600080fd5b82018360208201111561017157600080fd5b803590602001918460208302840111600160201b8311171561019257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101e157600080fd5b8201836020820111156101f357600080fd5b803590602001918460018302840111600160201b8311171561021457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561026657600080fd5b82018360208201111561027857600080fd5b803590602001918460208302840111600160201b8311171561029957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102e857600080fd5b8201836020820111156102fa57600080fd5b803590602001918460208302840111600160201b8311171561031b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561036a57600080fd5b82018360208201111561037c57600080fd5b803590602001918460208302840111600160201b8311171561039d57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092958435959094909350604081019250602001359050600160201b8111156103f457600080fd5b82018360208201111561040657600080fd5b803590602001918460018302840111600160201b8311171561042757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611541945050505050565b005b34801561047657600080fd5b506104686004803603602081101561048d57600080fd5b50356001600160a01b031661159b565b3480156104a957600080fd5b50610468600480360360808110156104c057600080fd5b8135916001600160581b03196020820135169160408201359190810190608081016060820135600160201b8111156104f757600080fd5b82018360208201111561050957600080fd5b803590602001918460018302840111600160201b8311171561052a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611645945050505050565b34801561057757600080fd5b50610468600480360361010081101561058f57600080fd5b813591602081013591810190606081016040820135600160201b8111156105b557600080fd5b8201836020820111156105c757600080fd5b803590602001918460208302840111600160201b831117156105e857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561063757600080fd5b82018360208201111561064957600080fd5b803590602001918460208302840111600160201b8311171561066a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106b957600080fd5b8201836020820111156106cb57600080fd5b803590602001918460208302840111600160201b831117156106ec57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092956001600160401b0385351695602086013595919450925060608101915060400135600160201b81111561075057600080fd5b82018360208201111561076257600080fd5b803590602001918460018302840111600160201b8311171561078357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611658945050505050565b3480156107d057600080fd5b5061081c600480360360408110156107e757600080fd5b60408051808201825291830192918183019183906002908390839080828437600092019190915250919450611f159350505050565b604080519115158252519081900360200190f35b34801561083c57600080fd5b50610468600480360360a081101561085357600080fd5b508035906020810190606001611f45565b34801561087057600080fd5b50610468600480360361014081101561088857600080fd5b81359160208101359160408201359163ffffffff6060820135169181019060a081016080820135600160201b8111156108c057600080fd5b8201836020820111156108d257600080fd5b803590602001918460208302840111600160201b831117156108f357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561094257600080fd5b82018360208201111561095457600080fd5b803590602001918460018302840111600160201b8311171561097557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156109c757600080fd5b8201836020820111156109d957600080fd5b803590602001918460208302840111600160201b831117156109fa57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a4957600080fd5b820183602082011115610a5b57600080fd5b803590602001918460208302840111600160201b83111715610a7c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610acb57600080fd5b820183602082011115610add57600080fd5b803590602001918460208302840111600160201b83111715610afe57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506120d4915050565b348015610b4a57600080fd5b50610468612128565b348015610b5f57600080fd5b506104686004803603610100811015610b7757600080fd5b81359160208101359160408201359190810190608081016060820135600160201b811115610ba457600080fd5b820183602082011115610bb657600080fd5b803590602001918460208302840111600160201b83111715610bd757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610c2657600080fd5b820183602082011115610c3857600080fd5b803590602001918460018302840111600160201b83111715610c5957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610cab57600080fd5b820183602082011115610cbd57600080fd5b803590602001918460208302840111600160201b83111715610cde57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d2d57600080fd5b820183602082011115610d3f57600080fd5b803590602001918460208302840111600160201b83111715610d6057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610daf57600080fd5b820183602082011115610dc157600080fd5b803590602001918460208302840111600160201b83111715610de257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506121cb945050505050565b348015610e2c57600080fd5b50610e35612498565b604080516001600160a01b039092168252519081900360200190f35b348015610e5d57600080fd5b506104686004803603610140811015610e7557600080fd5b8101908080606001906003806020026040519081016040528092919082600360200280828437600092019190915250919463ffffffff843581169560208601359091169461ffff60408201351694506001600160801b0360608201351693506001600160a01b03608082013581169360a083013590911692909160e081019060c00135600160201b811115610f0957600080fd5b820183602082011115610f1b57600080fd5b803590602001918460018302840111600160201b83111715610f3c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506124a7945050505050565b348015610f8957600080fd5b5061081c612df5565b348015610f9e57600080fd5b50610468600480360360a0811015610fb557600080fd5b8135916001600160581b03196020820135169160408201359190810190608081016060820135600160201b811115610fec57600080fd5b820183602082011115610ffe57600080fd5b803590602001918460018302840111600160201b8311171561101f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561107157600080fd5b82018360208201111561108357600080fd5b803590602001918460018302840111600160201b831117156110a457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612e06945050505050565b3480156110f157600080fd5b5061046860048036036101a081101561110957600080fd5b810190808060a00190600580602002604051908101604052809291908260056020028082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561119057600080fd5b8201836020820111156111a257600080fd5b803590602001918460208302840111600160201b831117156111c357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561121257600080fd5b82018360208201111561122457600080fd5b803590602001918460208302840111600160201b8311171561124557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561129457600080fd5b8201836020820111156112a657600080fd5b803590602001918460208302840111600160201b831117156112c757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561131657600080fd5b82018360208201111561132857600080fd5b803590602001918460208302840111600160201b8311171561134957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561139857600080fd5b8201836020820111156113aa57600080fd5b803590602001918460208302840111600160201b831117156113cb57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613034945050505050565b34801561141557600080fd5b506104686004803603602081101561142c57600080fd5b50356001600160a01b03166130f3565b6104686004803603604081101561145257600080fd5b81359190810190604081016020820135600160201b81111561147357600080fd5b82018360208201111561148557600080fd5b803590602001918460018302840111600160201b831117156114a657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550613158945050505050565b3480156114f357600080fd5b506104686004803603602081101561150a57600080fd5b50356131d5565b34801561151d57600080fd5b506104686004803603604081101561153457600080fd5b508035906020013561324d565b61158f6040518061014001604052808c81526020018b81526020018a8152602001898152602001888152602001878152602001868152602001858152602001848152602001838152506135d7565b50505050505050505050565b6115a3612df5565b6115f4576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0392909216919091179055565b6116528484843385613afc565b50505050565b600088815260036020526040902080546116b9576040805162461bcd60e51b815260206004820152601b60248201527f43616e2774206173736572742068616c746564206d616368696e650000000000604482015290519081900360640190fd5b60008989836000015484600201548b8b8b8b60405160200180888152602001878152602001868152602001858051906020019060200280838360005b8381101561170d5781810151838201526020016116f5565b50505050905001848051906020019060200280838360005b8381101561173d578181015183820152602001611725565b50505050905001838051906020019060200280838360005b8381101561176d578181015183820152602001611755565b50505050905001826001600160401b03166001600160401b031660c01b81526008019750505050505050506040516020818303038152906040528051906020012085604051602001808481526020018381526020018281526020019350505050604051602081830303815290604052805190602001209050816004015473__$a50780cb42d41d2927e39f529dc62d6697$__636a2dda6773__$6b4cc75dad3e0abd6ad83b3d907747c608$__63f0c8e96985886040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561187257818101518382015260200161185a565b50505050905090810190601f16801561189f5780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b1580156118bd57600080fd5b505af41580156118d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156118fa57600080fd5b810190808051600160201b81111561191157600080fd5b8201602081018481111561192457600080fd5b81518560208202830111600160201b8211171561194057600080fd5b50506040516001600160e01b031960e087901b1681526020600482018181528351602484015283519396509450849350604490910191818601910280838360005b83811015611999578181015183820152602001611981565b505050509050019250505060206040518083038186803b1580156119bc57600080fd5b505af41580156119d0573d6000803e3d6000fd5b505050506040513d60208110156119e657600080fd5b505114611a3a576040805162461bcd60e51b815260206004820181905260248201527f56616c696461746f72207369676e61747572657320646f6e2774206d61746368604482015290519081900360640190fd5b6002600b830154600160601b900460ff166002811115611a5657fe5b1415611ab257600a8201546001600160401b03600160c01b909104811690861611611ab25760405162461bcd60e51b8152600401808060200182810382526042815260200180615b4d6042913960600191505060405180910390fd5b600260009054906101000a90046001600160a01b03166001600160a01b031663c24651068b8a73__$6b4cc75dad3e0abd6ad83b3d907747c608$__630f89fbff8d8d8d6040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015611b4c578181015183820152602001611b34565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015611b8b578181015183820152602001611b73565b50505050905001848103825285818151815260200191508051906020019060200280838360005b83811015611bca578181015183820152602001611bb2565b50505050905001965050505050505060006040518083038186803b158015611bf157600080fd5b505af4158015611c05573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015611c2e57600080fd5b810190808051600160201b811115611c4557600080fd5b82016020810184811115611c5857600080fd5b81518560208202830111600160201b82111715611c7457600080fd5b50509291905050506040518463ffffffff1660e01b8152600401808481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015611cd5578181015183820152602001611cbd565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611d14578181015183820152602001611cfc565b5050505090500195505050505050602060405180830381600087803b158015611d3c57600080fd5b505af1158015611d50573d6000803e3d6000fd5b505050506040513d6020811015611d6657600080fd5b5051611db9576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b611dc282613c2d565b611dcb82613d37565b600b8201805460ff60601b1916600160611b179055600a820180546001600160c01b0316600160c01b6001600160401b0388160217905560405188518991899189918d916020918201918291818801910280838360005b83811015611e3a578181015183820152602001611e22565b50505050905001848051906020019060200280838360005b83811015611e6a578181015183820152602001611e52565b50505050905001838051906020019060200280838360005b83811015611e9a578181015183820152602001611e82565b5050505091909101928352505060408051808303815260208084018084528251929091019190912060018901558690526001600160401b038a1681830152518e94507f2147b50aa60f106b4862803c213873094d4a6495020f15f436027600a58b1fff9350908190036060019150a250505050505050505050565b80516000906001600160401b03164310801590611f3f575060208201516001600160401b03164311155b92915050565b6000838152600360205260409020600b810154600180549091600160501b900461ffff16908110611f7257fe5b6000918252602090912001546001600160a01b03163314611fc45760405162461bcd60e51b815260040180806020018281038252602d815260200180615d2f602d913960400191505060405180910390fd5b600b810154600160681b900460ff1661200e5760405162461bcd60e51b8152600401808060200182810382526026815260200180615d096026913960400191505060405180910390fd5b600b8101805460ff60681b191690556120716001600160801b03833516600c8301600086815b60200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002054613d7590919063ffffffff16565b83356001600160a01b03166000908152600c8301602081815260408320939093556120ab928501356001600160801b031691866001612034565b6001600160a01b03602094850135166000908152600c9290920190935260409020919091555050565b61158f6040518061014001604052808c81526020018b81526020018a81526020018963ffffffff16815260200188815260200187815260200186815260200185815260200184815260200183815250613dd6565b612130612df5565b612181576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008881526003602052604090206002600b820154600160601b900460ff1660028111156121f557fe5b146122315760405162461bcd60e51b8152600401808060200182810382526030815260200180615a516030913960400191505060405180910390fd5b600a810154600160801b90046001600160401b031643116122835760405162461bcd60e51b815260040180806020018281038252603e815260200180615a13603e913960400191505060405180910390fd5b80600101548685858a8c8a886040516020018085815260200184815260200183805190602001908083835b602083106122cd5780518252601f1990920191602091820191016122ae565b51815160209384036101000a60001901801990921691161790528551919093019285810192500280838360005b838110156123125781810151838201526020016122fa565b505050509050019450505050506040516020818303038152906040528051906020012060405160200180858051906020019060200280838360005b8381101561236557818101518382015260200161234d565b50505050905001848051906020019060200280838360005b8381101561239557818101518382015260200161237d565b50505050905001838051906020019060200280838360005b838110156123c55781810151838201526020016123ad565b50505050905001828152602001945050505050604051602081830303815290604052805190602001201461242a5760405162461bcd60e51b8152600401808060200182810382526034815260200180615cd56034913960400191505060405180910390fd5b600281018790556124408989888888888861449d565b600a810154604080516001600160401b03600160c01b909304929092168252518a917fd23ef3a25058f2acee4ca7aee5cc78a587b9fc76d2875098af1ffdfc2bcc74c9919081900360200190a2505050505050505050565b6000546001600160a01b031690565b6201000060418251816124b657fe5b04106124ff576040805162461bcd60e51b8152602060048201526013602482015272546f6f206d616e792076616c696461746f727360681b604482015290519081900360640190fd5b87516bffffffffffffffffffffffff1981161415612553576040805162461bcd60e51b815260206004820152600c60248201526b125b9d985b1a59081d9b525960a21b604482015290519081900360640190fd5b60015461ffff8616106125ad576040805162461bcd60e51b815260206004820152601d60248201527f496e76616c6964206368616c6c656e6765206d616e61676572206e756d000000604482015290519081900360640190fd5b6000846001600160801b03161161260b576040805162461bcd60e51b815260206004820181905260248201527f564d206d7573742072657175697265206e6f6e2d7a65726f206465706f736974604482015290519081900360640190fd5b6001600160a01b03831660009081526004602052604090205460ff166126625760405162461bcd60e51b8152600401808060200182810382526029815260200180615ade6029913960400191505060405180910390fd5b604080890151815163f0c8e96960e01b8152600481018281526024820193845284516044830152845160609473__$6b4cc75dad3e0abd6ad83b3d907747c608$__9463f0c8e969949093889390929160640190602085019080838360005b838110156126d85781810151838201526020016126c0565b50505050905090810190601f1680156127055780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b15801561272357600080fd5b505af4158015612737573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561276057600080fd5b810190808051600160201b81111561277757600080fd5b8201602081018481111561278a57600080fd5b81518560208202830111600160201b821117156127a657600080fd5b509094508c9350600292506127b9915050565b60200201518886868a8d600160200201518b8988604051602001808963ffffffff1663ffffffff1660e01b8152600401886001600160801b03166001600160801b031660801b8152601001876001600160a01b03166001600160a01b031660601b81526014018663ffffffff1663ffffffff1660e01b81526004018581526020018461ffff1661ffff1660f01b8152600201836001600160a01b03166001600160a01b031660601b8152601401828051906020019060200280838360005b8381101561288f578181015183820152602001612877565b50505050905001985050505050505050506040516020818303038152906040528051906020012014612900576040805162461bcd60e51b815260206004820152601560248201527410dc99585d194819185d18481a5b98dbdc9c9958dd605a1b604482015290519081900360640190fd5b60005b81518110156129ce5760025482516001600160a01b039091169063b8569ccd9084908490811061292f57fe5b602002602001015160601b6bffffffffffffffffffffffff191687896040518463ffffffff1660e01b815260040180848152602001836001600160a01b03166001600160a01b03168152602001826001600160801b031681526020019350505050600060405180830381600087803b1580156129aa57600080fd5b505af11580156129be573d6000803e3d6000fd5b5050600190920191506129039050565b50885160009081526003602052604090208960016020020151816000018190555073__$0d86abb4a722a612872fb80f4c7e7e95bd$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b158015612a3357600080fd5b505af4158015612a47573d6000803e3d6000fd5b505050506040513d6020811015612a5d57600080fd5b505160028201556040805163364df27760e01b8152905173__$0d86abb4a722a612872fb80f4c7e7e95bd$__9163364df277916004808301926020929190829003018186803b158015612aaf57600080fd5b505af4158015612ac3573d6000803e3d6000fd5b505050506040513d6020811015612ad957600080fd5b50516003820155600b8101805461ffff60501b1916600160501b61ffff8a16021760ff60601b19169055600060018201819055604051636a2dda6760e01b815260206004820181815285516024840152855173__$a50780cb42d41d2927e39f529dc62d6697$__94636a2dda67948894849360449092019286820192909102908190849084905b83811015612b78578181015183820152602001612b60565b505050509050019250505060206040518083038186803b158015612b9b57600080fd5b505af4158015612baf573d6000803e3d6000fd5b505050506040513d6020811015612bc557600080fd5b505160048201558151600b82018054600a840180546fffffffffffffffffffffffffffffffff19166001600160801b038b161790556009840180546001600160a01b03199081166001600160a01b038b81169190911790925560078601805490911691891691909117905569ffff000000000000000019166801000000000000000061ffff909316929092029190911763ffffffff191663ffffffff8b81169190911767ffffffff000000001916600160201b918b169190910217905560005b8251811015612cd657866001600160801b031682600c016000858481518110612caa57fe5b6020908102919091018101516001600160a01b0316825281019190915260400160002055600101612c85565b5089600060200201517f4b2401fcd345e80785cf05556c9dc50bc985a521241d1e3037b4ad86c77f62418a88888c8f600160200201518d8b8a604051808963ffffffff1663ffffffff168152602001886001600160801b03166001600160801b03168152602001876001600160a01b03166001600160a01b031681526020018663ffffffff1663ffffffff1681526020018581526020018461ffff1661ffff168152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015612dcf578181015183820152602001612db7565b50505050905001995050505050505050505060405180910390a250505050505050505050565b6000546001600160a01b0316331490565b600073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63c655d7aa8773__$0d86abb4a722a612872fb80f4c7e7e95bd$__638f346036876040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612e8a578181015183820152602001612e72565b50505050905090810190601f168015612eb75780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015612ed457600080fd5b505af4158015612ee8573d6000803e3d6000fd5b505050506040513d6020811015612efe57600080fd5b50516040805160208181019490945280820192909252606082018990526001600160581b03198a16608083015280516075818403018152609583018083528151918501919091206001600160e01b031960e087901b169091526099830181815260b98401928352885160d985015288519194899491939260f9909201919085019080838360005b83811015612f9d578181015183820152602001612f85565b50505050905090810190601f168015612fca5780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b158015612fe857600080fd5b505af4158015612ffc573d6000803e3d6000fd5b505050506040513d602081101561301257600080fd5b5051905061302c8686866001600160a01b03851687613afc565b505050505050565b6130e96040518061018001604052808a60006005811061305057fe5b602002015181526020018a60016005811061306757fe5b602002015181526020018a60026005811061307e57fe5b602002015181526020018a60036005811061309557fe5b602002015181526020018a6004600581106130ac57fe5b602002015181526020018963ffffffff1681526020018881526020018781526020018681526020018581526020018481526020018381525061482c565b5050505050505050565b6130fb612df5565b61314c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b613155816154e9565b50565b60025460408051636d31ebdb60e11b81526004810185905290516001600160a01b039092169163da63d7b6913491602480830192600092919082900301818588803b1580156131a657600080fd5b505af11580156131ba573d6000803e3d6000fd5b506131d19350859250600091503490503385615589565b5050565b600081815260036020526040902060078101546001600160a01b03163314613244576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6c79206f776e65722063616e2073687574646f776e2074686520564d0000604482015290519081900360640190fd5b6131d1826158a7565b600082815260036020526040902060088101546001600160a01b03163314156132a75760405162461bcd60e51b8152600401808060200182810382526021815260200180615bb36021913960400191505060405180910390fd5b600a810154600160801b90046001600160401b03164311156132fa5760405162461bcd60e51b8152600401808060200182810382526026815260200180615c826026913960400191505060405180910390fd5b6001600b820154600160601b900460ff16600281111561331657fe5b146133525760405162461bcd60e51b815260040180806020018281038252602f815260200180615a81602f913960400191505060405180910390fd5b336000908152600c82016020526040902054600a8201546001600160801b031611156133af5760405162461bcd60e51b815260040180806020018281038252602781526020018061598d6027913960400191505060405180910390fd5b806001015482146133f15760405162461bcd60e51b81526004018080602001828103825260398152602001806159b46039913960400191505060405180910390fd5b600a810154336000908152600c83016020526040902054613420916001600160801b031663ffffffff61592f16565b336000908152600c83016020526040812091909155600180830191909155600b82018054600160681b61ffff60601b1990911617908190558154600160501b90910461ffff1690811061346f57fe5b600091825260208083209091015460408051808201825260088601546001600160a01b039081168252338286015282518084018452600a8801546001600160801b031680825295810195909552600b8701548351632b50d42b60e01b8152600481018b81529290951696632b50d42b968b969495909463ffffffff909316938b93909260249091019187918190849084905b83811015613519578181015183820152602001613501565b5050505090500184600260200280838360005b8381101561354457818101518382015260200161352c565b505050509050018363ffffffff1663ffffffff16815260200182815260200195505050505050600060405180830381600087803b15801561358457600080fd5b505af1158015613598573d6000803e3d6000fd5b50506040805133815290518693507fb8d6a904566965067cb7e3e946d2a494a4b1f955ffec3e72e1e6ebe4e8e52c7292509081900360200190a2505050565b80516000908152600360205260409020805461363a576040805162461bcd60e51b815260206004820152601b60248201527f43616e2774206173736572742068616c746564206d616368696e650000000000604482015290519081900360640190fd5b600082600001518360400151846020015185608001518660e001516040516020018085815260200184815260200183805190602001908083835b602083106136935780518252601f199092019160209182019101613674565b51815160209384036101000a60001901801990921691161790528551919093019285810192500280838360005b838110156136d85781810151838201526020016136c0565b50505050905001945050505050604051602081830303815290604052805190602001208360000154846002015486606001518760a001518860c0015160405160200180878152602001868152602001858152602001848051906020019060200280838360005b8381101561375657818101518382015260200161373e565b50505050905001838051906020019060200280838360005b8381101561378657818101518382015260200161376e565b50505050905001828051906020019060200280838360005b838110156137b657818101518382015260200161379e565b50505050905001965050505050505060405160208183030381529060405280519060200120846101000151604051602001808481526020018381526020018281526020019350505050604051602081830303815290604052805190602001209050816004015473__$a50780cb42d41d2927e39f529dc62d6697$__636a2dda6773__$6b4cc75dad3e0abd6ad83b3d907747c608$__63f0c8e969858861012001516040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156138a9578181015183820152602001613891565b50505050905090810190601f1680156138d65780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b1580156138f457600080fd5b505af4158015613908573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561393157600080fd5b810190808051600160201b81111561394857600080fd5b8201602081018481111561395b57600080fd5b81518560208202830111600160201b8211171561397757600080fd5b50506040516001600160e01b031960e087901b1681526020600482018181528351602484015283519396509450849350604490910191818601910280838360005b838110156139d05781810151838201526020016139b8565b505050509050019250505060206040518083038186803b1580156139f357600080fd5b505af4158015613a07573d6000803e3d6000fd5b505050506040513d6020811015613a1d57600080fd5b505114613a71576040805162461bcd60e51b815260206004820181905260248201527f56616c696461746f72207369676e61747572657320646f6e2774206d61746368604482015290519081900360640190fd5b613a7a82613c2d565b600b8201805460ff60601b1916905560408301516002830155825160208401516060850151608086015160a087015160c088015160e0890151613ac29695949392919061449d565b82516040805183815290517f7349e5fdd7efd858059f6107d3f1309d322a1a495dc102e112296badf7742a869181900360200190a2505050565b600160f81b6001600160f81b0319601486901a60f81b161415613b9b576002546040805163cdf25dc160e01b81526004810185905260248101889052606087901c60448201526064810186905290516001600160a01b039092169163cdf25dc19160848082019260009290919082900301818387803b158015613b7e57600080fd5b505af1158015613b92573d6000803e3d6000fd5b50505050613c19565b60025460408051631240117b60e01b81526004810185905260248101889052606087901c60448201526064810186905290516001600160a01b0390921691631240117b9160848082019260009290919082900301818387803b158015613c0057600080fd5b505af1158015613c14573d6000803e3d6000fd5b505050505b613c268585858585615589565b5050505050565b6000600b820154600160601b900460ff166002811115613c4957fe5b14613cb757600a810154600160801b90046001600160401b0316431115613cb7576040805162461bcd60e51b815260206004820152601c60248201527f43616e27742063616e63656c2066696e616c697a656420737461746500000000604482015290519081900360640190fd5b6001600b820154600160601b900460ff166002811115613cd357fe5b141561315557600a81015460088201546001600160a01b03166000908152600c83016020526040902054613d15916001600160801b031663ffffffff613d7516565b60088201546001600160a01b03166000908152600c8301602052604090205550565b600b810154600a909101805467ffffffffffffffff60801b1916600160801b63ffffffff90931643016001600160401b031692909202919091179055565b600082820183811015613dcf576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b805160009081526003602052604090206001600b820154600160601b900460ff166002811115613e0257fe5b14613e3e5760405162461bcd60e51b8152600401808060200182810382526022815260200180615b2b6022913960400191505060405180910390fd5b600a810154600160801b90046001600160401b03164311613e905760405162461bcd60e51b8152600401808060200182810382526024815260200180615b076024913960400191505060405180910390fd5b8060010154826020015173__$6b4cc75dad3e0abd6ad83b3d907747c608$__632090372185604001518660600151600073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63252001608a608001518b60a001518c60c001518d60e001518e61010001516040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015613f53578181015183820152602001613f3b565b5050505090500186810385528a818151815260200191508051906020019080838360005b83811015613f8f578181015183820152602001613f77565b50505050905090810190601f168015613fbc5780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015613ff1578181015183820152602001613fd9565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015614030578181015183820152602001614018565b50505050905001868103825287818151815260200191508051906020019060200280838360005b8381101561406f578181015183820152602001614057565b505050509050019a505050505050505050505060206040518083038186803b15801561409a57600080fd5b505af41580156140ae573d6000803e3d6000fd5b505050506040513d60208110156140c457600080fd5b810190808051906020019092919050505060008a610120015173__$6b4cc75dad3e0abd6ad83b3d907747c608$__630f89fbff8d608001518e60c001518f60e001516040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b8381101561415d578181015183820152602001614145565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561419c578181015183820152602001614184565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156141db5781810151838201526020016141c3565b50505050905001965050505050505060006040518083038186803b15801561420257600080fd5b505af4158015614216573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561423f57600080fd5b810190808051600160201b81111561425657600080fd5b8201602081018481111561426957600080fd5b81518560208202830111600160201b8211171561428557600080fd5b505060405160e08c811b6001600160e01b0319168252600482018c815263ffffffff8c166024840152604483018b9052606483018a90526084830189905260a4830188905260c48301918252835160e484015283519396509450925061010401906020808601910280838360005b8381101561430b5781810151838201526020016142f3565b505050509050019850505050505050505060206040518083038186803b15801561433457600080fd5b505af4158015614348573d6000803e3d6000fd5b505050506040513d602081101561435e57600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920190528051910120146143c55760405162461bcd60e51b81526004018080602001828103825260398152602001806159b46039913960400191505060405180910390fd5b600a81015460088201546001600160a01b03166000908152600c83016020526040902054614401916001600160801b031663ffffffff613d7516565b60088201546001600160a01b03166000908152600c830160205260409081902091909155825190830151608084015160a085015160c086015160e08701516101008801516144549695949392919061449d565b81516040808401516101208501518251918252602082015281517f9116ee6e53cceb1cd0f4ba7dafd4076656846670da45f42fc6dde02c9ebb8cc9929181900390910190a25050565b6000878152600360205260408120835160609190835b818110156146705773__$0d86abb4a722a612872fb80f4c7e7e95bd$__631f3d4d4e8a876040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015614529578181015183820152602001614511565b50505050905090810190601f1680156145565780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b15801561457457600080fd5b505af4158015614588573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160409081528110156145b157600080fd5b815160208301805191939283019291600160201b8111156145d157600080fd5b820160208101848111156145e457600080fd5b8151600160201b8111828201871017156145fd57600080fd5b50508a519499509750614668938a93508592508210905061461a57fe5b60200260200101518b8a848151811061462f57fe5b602002602001015161ffff168151811061464557fe5b602002602001015189848151811061465957fe5b60200260200101518f88613afc565b6001016144b3565b50898255600b8201805460ff60601b191690556040805163364df27760e01b8152905173__$0d86abb4a722a612872fb80f4c7e7e95bd$__9163364df277916004808301926020929190829003018186803b1580156146ce57600080fd5b505af41580156146e2573d6000803e3d6000fd5b505050506040513d60208110156146f857600080fd5b50516003830154146148115773__$6b4cc75dad3e0abd6ad83b3d907747c608$__63f11fcc26836002015484600301546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561476357600080fd5b505af4158015614777573d6000803e3d6000fd5b505050506040513d602081101561478d57600080fd5b505160028301556040805163364df27760e01b8152905173__$0d86abb4a722a612872fb80f4c7e7e95bd$__9163364df277916004808301926020929190829003018186803b1580156147df57600080fd5b505af41580156147f3573d6000803e3d6000fd5b505050506040513d602081101561480957600080fd5b505160038301555b8961481f5761481f8b6158a7565b5050505050505050505050565b8051600090815260036020526040812090600b820154600160601b900460ff16600281111561485757fe5b146148935760405162461bcd60e51b815260040180806020018281038252602d815260200180615ca8602d913960400191505060405180910390fd5b8054158015906148a557508054600114155b6148e05760405162461bcd60e51b815260040180806020018281038252603e815260200180615bd4603e913960400191505060405180910390fd5b600b810154600160681b900460ff161561492b5760405162461bcd60e51b815260040180806020018281038252602e815260200180615ab0602e913960400191505060405180910390fd5b336000908152600c82016020526040902054600a8201546001600160801b031611156149885760405162461bcd60e51b8152600401808060200182810382526027815260200180615c346027913960400191505060405180910390fd5b600b81015460a083015163ffffffff600160201b9092048216911611156149f6576040805162461bcd60e51b815260206004820152601f60248201527f547269656420746f206578656375746520746f6f206d616e7920737465707300604482015290519081900360640190fd5b614a038260c00151611f15565b614a3e5760405162461bcd60e51b8152600401808060200182810382526024815260200180615b8f6024913960400191505060405180910390fd5b8054602083015114614a815760405162461bcd60e51b8152600401808060200182810382526027815260200180615c5b6027913960400191505060405180910390fd5b806002015482604001511480614b24575073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63f11fcc26826002015483600301546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b158015614af157600080fd5b505af4158015614b05573d6000803e3d6000fd5b505050506040513d6020811015614b1b57600080fd5b50516040830151145b614b5f5760405162461bcd60e51b8152600401808060200182810382526022815260200180615c126022913960400191505060405180910390fd5b606073__$6b4cc75dad3e0abd6ad83b3d907747c608$__630f89fbff8460e001518561012001518661014001516040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015614be3578181015183820152602001614bcb565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015614c22578181015183820152602001614c0a565b50505050905001848103825285818151815260200191508051906020019060200280838360005b83811015614c61578181015183820152602001614c49565b50505050905001965050505050505060006040518083038186803b158015614c8857600080fd5b505af4158015614c9c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015614cc557600080fd5b810190808051600160201b811115614cdc57600080fd5b82016020810184811115614cef57600080fd5b81518560208202830111600160201b82111715614d0b57600080fd5b5050600254875160e0890151604051636123288360e11b815260048101838152606060248301908152835160648401528351969a506001600160a01b03909516985063c24651069750929550909388939160448101916084909101906020808801910280838360005b83811015614d8c578181015183820152602001614d74565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015614dcb578181015183820152602001614db3565b5050505090500195505050505050602060405180830381600087803b158015614df357600080fd5b505af1158015614e07573d6000803e3d6000fd5b505050506040513d6020811015614e1d57600080fd5b5051614e70576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b614e7982613d37565b600073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63b32774958560e001518661010001518761012001518861014001518961016001516040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015614f11578181015183820152602001614ef9565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b83811015614f50578181015183820152602001614f38565b50505050905001868103845289818151815260200191508051906020019060200280838360005b83811015614f8f578181015183820152602001614f77565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015614fce578181015183820152602001614fb6565b50505050905001868103825287818151815260200191508051906020019060200280838360005b8381101561500d578181015183820152602001614ff5565b505050509050019a505050505050505050505060206040518083038186803b15801561503857600080fd5b505af415801561504c573d6000803e3d6000fd5b505050506040513d602081101561506257600080fd5b5051602085015160c086015160408088015160e089015182516307c50ab360e31b81526004810186815296975073__$6b4cc75dad3e0abd6ad83b3d907747c608$__96633e2855989695948a9260240190869080838360005b838110156150d35781810151838201526020016150bb565b505050509050018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015615120578181015183820152602001615108565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561515f578181015183820152602001615147565b5050505090500197505050505050505060206040518083038186803b15801561518757600080fd5b505af415801561519b573d6000803e3d6000fd5b505050506040513d60208110156151b157600080fd5b5051606085015160a08601516080870151604051632090372160e01b81526004810184815263ffffffff84166024830152600060448301819052606483018890526084830181905260a4830184905260e060c48401908152895160e4850152895173__$6b4cc75dad3e0abd6ad83b3d907747c608$__976320903721979096909593948b94869492938e93916101040190602085810191028083838a5b8381101561526657818101518382015260200161524e565b505050509050019850505050505050505060206040518083038186803b15801561528f57600080fd5b505af41580156152a3573d6000803e3d6000fd5b505050506040513d60208110156152b957600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920181528151918301919091206001860155600a850154336000908152600c870190935291205461531a916001600160801b031663ffffffff61592f16565b336000818152600c860160205260409020919091556008840180546001600160a01b0319169091179055600b830180546001919060ff60601b1916600160601b83021790555083600001517f8466b2589a76ab0ff96670e87c3100a0a418647ac17a93ecbde4d036590aa344604051806060016040528087602001518152602001876040015181526020018760600151815250338760c001518860e001518960a00151878b608001518a6040518089600360200280838360005b838110156153ec5781810151838201526020016153d4565b505050506001600160a01b038b1692019182525060200187604080838360005b8381101561542457818101518382015260200161540c565b50505050905001806020018663ffffffff1663ffffffff16815260200185815260200184815260200180602001838103835288818151815260200191508051906020019060200280838360005b83811015615489578181015183820152602001615471565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156154c85781810151838201526020016154b0565b505050509050019a505050505050505050505060405180910390a250505050565b6001600160a01b03811661552e5760405162461bcd60e51b81526004018080602001828103825260268152602001806159ed6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6bffffffffffffffffffffffff19851685146157d2576000858152600360209081526040808320905163479a301b60e11b81526004810183815285516024830152855192949373__$6b4cc75dad3e0abd6ad83b3d907747c608$__93632a0500d8938c9373__$0d86abb4a722a612872fb80f4c7e7e95bd$__93638f346036938b938392604490910191908501908083838e5b8381101561563457818101518382015260200161561c565b50505050905090810190601f1680156156615780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b15801561567e57600080fd5b505af4158015615692573d6000803e3d6000fd5b505050506040513d60208110156156a857600080fd5b5051604080516001600160e01b031960e086901b168152600481019390935260248301919091526001600160581b03198a16604483015260648201899052608482018890525160a4808301926020929190829003018186803b15801561570d57600080fd5b505af4158015615721573d6000803e3d6000fd5b505050506040513d602081101561573757600080fd5b5051600383015460408051636bc68c7560e11b81526004810192909252602482018390525191925073__$6b4cc75dad3e0abd6ad83b3d907747c608$__9163d78d18ea91604480820192602092909190829003018186803b15801561579b57600080fd5b505af41580156157af573d6000803e3d6000fd5b505050506040513d60208110156157c557600080fd5b5051600390920191909155505b847f983eec3b16cf85cbe329d288a7863b78dc7d1c9a0d690b21434d4cc1f73539e58386868560405180858152602001846affffffffffffffffffffff19166affffffffffffffffffffff1916815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561586357818101518382015260200161584b565b50505050905090810190601f1680156158905780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a25050505050565b6000908152600360208190526040822082815560018101839055600281018390559081018290556004810182905560058101829055600681018290556007810180546001600160a01b0319908116909155600882018054821690556009820180549091169055600a810191909155600b0180546dffffffffffffffffffffffffffff19169055565b600082821115615986576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6e4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737343616e206f6e6c7920636f6e6669726d20617373657274696f6e2077686f7365206368616c6c656e676520646561646c696e65206861732070617373656443616e206f6e6c7920636f6e6669726d20696620746865726520697320612070656e64696e6720617373657274696f6e417373657274696f6e206d7573742062652070656e64696e6720746f20696e697469617465206368616c6c656e676543616e206f6e6c792064697370757461626c6520617373657274206966206e6f7420696e206368616c6c656e676553656c65637465642063757272656e6379206973206e6f7420616e2061636365707465642074797065417373657274696f6e206973207374696c6c2070656e64696e67206368616c6c656e6765564d20646f6573206e6f74206861766520617373657274696f6e2070656e64696e6743616e206f6e6c79207375706572736564652070726576696f757320617373657274696f6e207769746820677265617465722073657175656e6365206e756d626572507265636f6e646974696f6e3a206e6f742077697468696e2074696d6520626f756e64734368616c6c656e676520776173206372656174656420627920617373657274657243616e206f6e6c792064697370757461626c6520617373657274206966206d616368696e65206973206e6f74206572726f726564206f722068616c746564507265636f6e646974696f6e3a20696e626f7820646f6573206e6f74206d6174636856616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f77507265636f6e646974696f6e3a207374617465206861736820646f6573206e6f74206d617463684368616c6c656e676520646964206e6f7420636f6d65206265666f726520646561646c696e6543616e206f6e6c792064697370757461626c65206173736572742066726f6d2077616974696e6720737461746543616e206f6e6c7920636f6e6669726d20617373657274696f6e20746861742069732063757272656e746c792070656e64696e67564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e6765a265627a7a72305820437c666dba1f20e89e3699872258f85324c87241606fca8551b1a7a6645a25b264736f6c634300050a0032"

// DeployVMTracker deploys a new Ethereum contract, binding an instance of VMTracker to it.
func DeployVMTracker(auth *bind.TransactOpts, backend bind.ContractBackend, _balanceTrackerAddress common.Address) (common.Address, *types.Transaction, *VMTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(VMTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	VMTrackerBin = strings.Replace(VMTrackerBin, "__$0d86abb4a722a612872fb80f4c7e7e95bd$__", arbValueAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	VMTrackerBin = strings.Replace(VMTrackerBin, "__$6b4cc75dad3e0abd6ad83b3d907747c608$__", arbProtocolAddr.String()[2:], -1)

	merkleLibAddr, _, _, _ := DeployMerkleLib(auth, backend)
	VMTrackerBin = strings.Replace(VMTrackerBin, "__$a50780cb42d41d2927e39f529dc62d6697$__", merkleLibAddr.String()[2:], -1)

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
	var (
		ret0 = new(bool)
	)
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
	var (
		ret0 = new(common.Address)
	)
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
	var (
		ret0 = new(bool)
	)
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

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x6f0edcd5.
//
// Solidity: function confirmDisputableAsserted(bytes32 _vmId, bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, bytes32[] _messageDestination, bytes32 _logsAccHash) returns()
func (_VMTracker *VMTrackerTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _vmId [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestination [][32]byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "confirmDisputableAsserted", _vmId, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestination, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x6f0edcd5.
//
// Solidity: function confirmDisputableAsserted(bytes32 _vmId, bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, bytes32[] _messageDestination, bytes32 _logsAccHash) returns()
func (_VMTracker *VMTrackerSession) ConfirmDisputableAsserted(_vmId [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestination [][32]byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmDisputableAsserted(&_VMTracker.TransactOpts, _vmId, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestination, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x6f0edcd5.
//
// Solidity: function confirmDisputableAsserted(bytes32 _vmId, bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, bytes32[] _messageDestination, bytes32 _logsAccHash) returns()
func (_VMTracker *VMTrackerTransactorSession) ConfirmDisputableAsserted(_vmId [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestination [][32]byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmDisputableAsserted(&_VMTracker.TransactOpts, _vmId, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestination, _logsAccHash)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0x84481bab.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerTransactor) ConfirmUnanimousAsserted(opts *bind.TransactOpts, _vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "confirmUnanimousAsserted", _vmId, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0x84481bab.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerSession) ConfirmUnanimousAsserted(_vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmUnanimousAsserted(&_VMTracker.TransactOpts, _vmId, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0x84481bab.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) returns()
func (_VMTracker *VMTrackerTransactorSession) ConfirmUnanimousAsserted(_vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmUnanimousAsserted(&_VMTracker.TransactOpts, _vmId, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
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

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0x0725397e.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactor) FinalizedUnanimousAssert(opts *bind.TransactOpts, _vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "finalizedUnanimousAssert", _vmId, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination, _logsAccHash, _signatures)
}

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0x0725397e.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerSession) FinalizedUnanimousAssert(_vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.FinalizedUnanimousAssert(&_VMTracker.TransactOpts, _vmId, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination, _logsAccHash, _signatures)
}

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0x0725397e.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _vmId, bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactorSession) FinalizedUnanimousAssert(_vmId [32]byte, _afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.FinalizedUnanimousAssert(&_VMTracker.TransactOpts, _vmId, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination, _logsAccHash, _signatures)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0xd37a1a95.
//
// Solidity: function forwardMessage(bytes32 _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_VMTracker *VMTrackerTransactor) ForwardMessage(opts *bind.TransactOpts, _destination [32]byte, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "forwardMessage", _destination, _tokenType, _amount, _data, _signature)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0xd37a1a95.
//
// Solidity: function forwardMessage(bytes32 _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_VMTracker *VMTrackerSession) ForwardMessage(_destination [32]byte, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ForwardMessage(&_VMTracker.TransactOpts, _destination, _tokenType, _amount, _data, _signature)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0xd37a1a95.
//
// Solidity: function forwardMessage(bytes32 _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_VMTracker *VMTrackerTransactorSession) ForwardMessage(_destination [32]byte, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ForwardMessage(&_VMTracker.TransactOpts, _destination, _tokenType, _amount, _data, _signature)
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

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xda9f547a.
//
// Solidity: function pendingDisputableAssert(bytes32[5] _fields, uint32 _numSteps, uint64[2] timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNum, uint256[] _msgAmount, bytes32[] _msgDestination) returns()
func (_VMTracker *VMTrackerTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _fields [5][32]byte, _numSteps uint32, timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNum []uint16, _msgAmount []*big.Int, _msgDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "pendingDisputableAssert", _fields, _numSteps, timeBounds, _tokenTypes, _messageDataHash, _messageTokenNum, _msgAmount, _msgDestination)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xda9f547a.
//
// Solidity: function pendingDisputableAssert(bytes32[5] _fields, uint32 _numSteps, uint64[2] timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNum, uint256[] _msgAmount, bytes32[] _msgDestination) returns()
func (_VMTracker *VMTrackerSession) PendingDisputableAssert(_fields [5][32]byte, _numSteps uint32, timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNum []uint16, _msgAmount []*big.Int, _msgDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.PendingDisputableAssert(&_VMTracker.TransactOpts, _fields, _numSteps, timeBounds, _tokenTypes, _messageDataHash, _messageTokenNum, _msgAmount, _msgDestination)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xda9f547a.
//
// Solidity: function pendingDisputableAssert(bytes32[5] _fields, uint32 _numSteps, uint64[2] timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNum, uint256[] _msgAmount, bytes32[] _msgDestination) returns()
func (_VMTracker *VMTrackerTransactorSession) PendingDisputableAssert(_fields [5][32]byte, _numSteps uint32, timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNum []uint16, _msgAmount []*big.Int, _msgDestination [][32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.PendingDisputableAssert(&_VMTracker.TransactOpts, _fields, _numSteps, timeBounds, _tokenTypes, _messageDataHash, _messageTokenNum, _msgAmount, _msgDestination)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0x3f2832cd.
//
// Solidity: function pendingUnanimousAssert(bytes32 _vmId, bytes32 _unanRest, bytes21[] _tokenTypes, uint16[] _messageTokenNum, uint256[] _messageAmount, uint64 _sequenceNum, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactor) PendingUnanimousAssert(opts *bind.TransactOpts, _vmId [32]byte, _unanRest [32]byte, _tokenTypes [][21]byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _sequenceNum uint64, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "pendingUnanimousAssert", _vmId, _unanRest, _tokenTypes, _messageTokenNum, _messageAmount, _sequenceNum, _logsAccHash, _signatures)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0x3f2832cd.
//
// Solidity: function pendingUnanimousAssert(bytes32 _vmId, bytes32 _unanRest, bytes21[] _tokenTypes, uint16[] _messageTokenNum, uint256[] _messageAmount, uint64 _sequenceNum, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerSession) PendingUnanimousAssert(_vmId [32]byte, _unanRest [32]byte, _tokenTypes [][21]byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _sequenceNum uint64, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.PendingUnanimousAssert(&_VMTracker.TransactOpts, _vmId, _unanRest, _tokenTypes, _messageTokenNum, _messageAmount, _sequenceNum, _logsAccHash, _signatures)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0x3f2832cd.
//
// Solidity: function pendingUnanimousAssert(bytes32 _vmId, bytes32 _unanRest, bytes21[] _tokenTypes, uint16[] _messageTokenNum, uint256[] _messageAmount, uint64 _sequenceNum, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactorSession) PendingUnanimousAssert(_vmId [32]byte, _unanRest [32]byte, _tokenTypes [][21]byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _sequenceNum uint64, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.PendingUnanimousAssert(&_VMTracker.TransactOpts, _vmId, _unanRest, _tokenTypes, _messageTokenNum, _messageAmount, _sequenceNum, _logsAccHash, _signatures)
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

// VMTrackerConfirmedDisputableAssertionIterator is returned from FilterConfirmedDisputableAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedDisputableAssertion events raised by the VMTracker contract.
type VMTrackerConfirmedDisputableAssertionIterator struct {
	Event *VMTrackerConfirmedDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerConfirmedDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerConfirmedDisputableAssertion)
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
		it.Event = new(VMTrackerConfirmedDisputableAssertion)
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
func (it *VMTrackerConfirmedDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerConfirmedDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerConfirmedDisputableAssertion represents a ConfirmedDisputableAssertion event raised by the VMTracker contract.
type VMTrackerConfirmedDisputableAssertion struct {
	VmId        [32]byte
	NewState    [32]byte
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x9116ee6e53cceb1cd0f4ba7dafd4076656846670da45f42fc6dde02c9ebb8cc9.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 indexed vmId, bytes32 newState, bytes32 logsAccHash)
func (_VMTracker *VMTrackerFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerConfirmedDisputableAssertionIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "ConfirmedDisputableAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerConfirmedDisputableAssertionIterator{contract: _VMTracker.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x9116ee6e53cceb1cd0f4ba7dafd4076656846670da45f42fc6dde02c9ebb8cc9.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 indexed vmId, bytes32 newState, bytes32 logsAccHash)
func (_VMTracker *VMTrackerFilterer) WatchConfirmedDisputableAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerConfirmedDisputableAssertion, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "ConfirmedDisputableAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerConfirmedDisputableAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
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

// ParseConfirmedDisputableAssertion is a log parse operation binding the contract event 0x9116ee6e53cceb1cd0f4ba7dafd4076656846670da45f42fc6dde02c9ebb8cc9.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 indexed vmId, bytes32 newState, bytes32 logsAccHash)
func (_VMTracker *VMTrackerFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*VMTrackerConfirmedDisputableAssertion, error) {
	event := new(VMTrackerConfirmedDisputableAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseConfirmedUnanimousAssertion is a log parse operation binding the contract event 0xd23ef3a25058f2acee4ca7aee5cc78a587b9fc76d2875098af1ffdfc2bcc74c9.
//
// Solidity: event ConfirmedUnanimousAssertion(bytes32 indexed vmId, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) ParseConfirmedUnanimousAssertion(log types.Log) (*VMTrackerConfirmedUnanimousAssertion, error) {
	event := new(VMTrackerConfirmedUnanimousAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerFinalizedUnanimousAssertionIterator is returned from FilterFinalizedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for FinalizedUnanimousAssertion events raised by the VMTracker contract.
type VMTrackerFinalizedUnanimousAssertionIterator struct {
	Event *VMTrackerFinalizedUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerFinalizedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerFinalizedUnanimousAssertion)
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
		it.Event = new(VMTrackerFinalizedUnanimousAssertion)
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
func (it *VMTrackerFinalizedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerFinalizedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerFinalizedUnanimousAssertion represents a FinalizedUnanimousAssertion event raised by the VMTracker contract.
type VMTrackerFinalizedUnanimousAssertion struct {
	VmId     [32]byte
	UnanHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizedUnanimousAssertion is a free log retrieval operation binding the contract event 0x7349e5fdd7efd858059f6107d3f1309d322a1a495dc102e112296badf7742a86.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash)
func (_VMTracker *VMTrackerFilterer) FilterFinalizedUnanimousAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerFinalizedUnanimousAssertionIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "FinalizedUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerFinalizedUnanimousAssertionIterator{contract: _VMTracker.contract, event: "FinalizedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchFinalizedUnanimousAssertion is a free log subscription operation binding the contract event 0x7349e5fdd7efd858059f6107d3f1309d322a1a495dc102e112296badf7742a86.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash)
func (_VMTracker *VMTrackerFilterer) WatchFinalizedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerFinalizedUnanimousAssertion, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "FinalizedUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerFinalizedUnanimousAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
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

// ParseFinalizedUnanimousAssertion is a log parse operation binding the contract event 0x7349e5fdd7efd858059f6107d3f1309d322a1a495dc102e112296badf7742a86.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash)
func (_VMTracker *VMTrackerFilterer) ParseFinalizedUnanimousAssertion(log types.Log) (*VMTrackerFinalizedUnanimousAssertion, error) {
	event := new(VMTrackerFinalizedUnanimousAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0xb8d6a904566965067cb7e3e946d2a494a4b1f955ffec3e72e1e6ebe4e8e52c72.
//
// Solidity: event InitiatedChallenge(bytes32 indexed vmId, address challenger)
func (_VMTracker *VMTrackerFilterer) ParseInitiatedChallenge(log types.Log) (*VMTrackerInitiatedChallenge, error) {
	event := new(VMTrackerInitiatedChallenge)
	if err := _VMTracker.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x983eec3b16cf85cbe329d288a7863b78dc7d1c9a0d690b21434d4cc1f73539e5.
//
// Solidity: event MessageDelivered(bytes32 indexed vmId, bytes32 destination, bytes21 tokenType, uint256 value, bytes data)
func (_VMTracker *VMTrackerFilterer) ParseMessageDelivered(log types.Log) (*VMTrackerMessageDelivered, error) {
	event := new(VMTrackerMessageDelivered)
	if err := _VMTracker.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VMTracker *VMTrackerFilterer) ParseOwnershipTransferred(log types.Log) (*VMTrackerOwnershipTransferred, error) {
	event := new(VMTrackerOwnershipTransferred)
	if err := _VMTracker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerPendingDisputableAssertionIterator is returned from FilterPendingDisputableAssertion and is used to iterate over the raw logs and unpacked data for PendingDisputableAssertion events raised by the VMTracker contract.
type VMTrackerPendingDisputableAssertionIterator struct {
	Event *VMTrackerPendingDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerPendingDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerPendingDisputableAssertion)
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
		it.Event = new(VMTrackerPendingDisputableAssertion)
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
func (it *VMTrackerPendingDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerPendingDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerPendingDisputableAssertion represents a PendingDisputableAssertion event raised by the VMTracker contract.
type VMTrackerPendingDisputableAssertion struct {
	VmId            [32]byte
	Fields          [3][32]byte
	Asserter        common.Address
	TimeBounds      [2]uint64
	TokenTypes      [][21]byte
	NumSteps        uint32
	LastMessageHash [32]byte
	LogsAccHash     [32]byte
	Amounts         []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x8466b2589a76ab0ff96670e87c3100a0a418647ac17a93ecbde4d036590aa344.
//
// Solidity: event PendingDisputableAssertion(bytes32 indexed vmId, bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_VMTracker *VMTrackerFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerPendingDisputableAssertionIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "PendingDisputableAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerPendingDisputableAssertionIterator{contract: _VMTracker.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x8466b2589a76ab0ff96670e87c3100a0a418647ac17a93ecbde4d036590aa344.
//
// Solidity: event PendingDisputableAssertion(bytes32 indexed vmId, bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_VMTracker *VMTrackerFilterer) WatchPendingDisputableAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerPendingDisputableAssertion, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "PendingDisputableAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerPendingDisputableAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x8466b2589a76ab0ff96670e87c3100a0a418647ac17a93ecbde4d036590aa344.
//
// Solidity: event PendingDisputableAssertion(bytes32 indexed vmId, bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_VMTracker *VMTrackerFilterer) ParsePendingDisputableAssertion(log types.Log) (*VMTrackerPendingDisputableAssertion, error) {
	event := new(VMTrackerPendingDisputableAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerPendingUnanimousAssertionIterator is returned from FilterPendingUnanimousAssertion and is used to iterate over the raw logs and unpacked data for PendingUnanimousAssertion events raised by the VMTracker contract.
type VMTrackerPendingUnanimousAssertionIterator struct {
	Event *VMTrackerPendingUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *VMTrackerPendingUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerPendingUnanimousAssertion)
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
		it.Event = new(VMTrackerPendingUnanimousAssertion)
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
func (it *VMTrackerPendingUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerPendingUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerPendingUnanimousAssertion represents a PendingUnanimousAssertion event raised by the VMTracker contract.
type VMTrackerPendingUnanimousAssertion struct {
	VmId        [32]byte
	UnanHash    [32]byte
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingUnanimousAssertion is a free log retrieval operation binding the contract event 0x2147b50aa60f106b4862803c213873094d4a6495020f15f436027600a58b1fff.
//
// Solidity: event PendingUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) FilterPendingUnanimousAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerPendingUnanimousAssertionIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "PendingUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerPendingUnanimousAssertionIterator{contract: _VMTracker.contract, event: "PendingUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingUnanimousAssertion is a free log subscription operation binding the contract event 0x2147b50aa60f106b4862803c213873094d4a6495020f15f436027600a58b1fff.
//
// Solidity: event PendingUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) WatchPendingUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerPendingUnanimousAssertion, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "PendingUnanimousAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerPendingUnanimousAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
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

// ParsePendingUnanimousAssertion is a log parse operation binding the contract event 0x2147b50aa60f106b4862803c213873094d4a6495020f15f436027600a58b1fff.
//
// Solidity: event PendingUnanimousAssertion(bytes32 indexed vmId, bytes32 unanHash, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) ParsePendingUnanimousAssertion(log types.Log) (*VMTrackerPendingUnanimousAssertion, error) {
	event := new(VMTrackerPendingUnanimousAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
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
	VmId                [32]byte
	GracePeriod         uint32
	EscrowRequired      *big.Int
	EscrowCurrency      common.Address
	MaxExecutionSteps   uint32
	VmState             [32]byte
	ChallengeManagerNum uint16
	Owner               common.Address
	Validators          []common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVMCreated is a free log retrieval operation binding the contract event 0x4b2401fcd345e80785cf05556c9dc50bc985a521241d1e3037b4ad86c77f6241.
//
// Solidity: event VMCreated(bytes32 indexed vmId, uint32 _gracePeriod, uint128 _escrowRequired, address _escrowCurrency, uint32 _maxExecutionSteps, bytes32 _vmState, uint16 _challengeManagerNum, address _owner, address[] validators)
func (_VMTracker *VMTrackerFilterer) FilterVMCreated(opts *bind.FilterOpts, vmId [][32]byte) (*VMTrackerVMCreatedIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "VMCreated", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &VMTrackerVMCreatedIterator{contract: _VMTracker.contract, event: "VMCreated", logs: logs, sub: sub}, nil
}

// WatchVMCreated is a free log subscription operation binding the contract event 0x4b2401fcd345e80785cf05556c9dc50bc985a521241d1e3037b4ad86c77f6241.
//
// Solidity: event VMCreated(bytes32 indexed vmId, uint32 _gracePeriod, uint128 _escrowRequired, address _escrowCurrency, uint32 _maxExecutionSteps, bytes32 _vmState, uint16 _challengeManagerNum, address _owner, address[] validators)
func (_VMTracker *VMTrackerFilterer) WatchVMCreated(opts *bind.WatchOpts, sink chan<- *VMTrackerVMCreated, vmId [][32]byte) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "VMCreated", vmIdRule)
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

// ParseVMCreated is a log parse operation binding the contract event 0x4b2401fcd345e80785cf05556c9dc50bc985a521241d1e3037b4ad86c77f6241.
//
// Solidity: event VMCreated(bytes32 indexed vmId, uint32 _gracePeriod, uint128 _escrowRequired, address _escrowCurrency, uint32 _maxExecutionSteps, bytes32 _vmState, uint16 _challengeManagerNum, address _owner, address[] validators)
func (_VMTracker *VMTrackerFilterer) ParseVMCreated(log types.Log) (*VMTrackerVMCreated, error) {
	event := new(VMTrackerVMCreated)
	if err := _VMTracker.contract.UnpackLog(event, "VMCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
