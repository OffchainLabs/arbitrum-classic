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

// GasRefunderMetaData contains all meta data concerning the GasRefunder contract.
var GasRefunderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumGasRefunder.CommonParameterKey\",\"name\":\"parameter\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CommonParameterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"ContractAllowedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DisallowerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumGasRefunder.RefundDenyReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"RefundGasCostsDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"}],\"name\":\"RefundedGasCosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"RefundeeAllowedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"allowContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"allowRefundees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedRefundees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commonParams\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"maxRefundeeBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"extraGasMargin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"calldataCost\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"maxGasTip\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasCost\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxSingleGasUsage\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"disallowContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"disallowRefundees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disallower\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"refundee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"onGasSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newValue\",\"type\":\"uint8\"}],\"name\":\"setCalldataCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDisallower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setExtraGasMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValue\",\"type\":\"uint64\"}],\"name\":\"setMaxGasCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValue\",\"type\":\"uint64\"}],\"name\":\"setMaxGasTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newValue\",\"type\":\"uint128\"}],\"name\":\"setMaxRefundeeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setMaxSingleGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a336100a5565b6040805160c08101825260008152610fa06020820152600c9181019190915263773594006060820152641bf08eb0006080820152621e848060a090910152600480546001600160e81b03191678773594000c00000fa000000000000000000000000000000000179055600580546001600160601b0319166a1e84800000001bf08eb0001790556100f5565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611229806101046000396000f3fe6080604052600436106101235760003560e01c8063ca101295116100a0578063efe12b0111610064578063efe12b011461040d578063f1e845ca1461042d578063f2fde38b1461044d578063f3fef3a31461046d578063f52128eb1461048d57600080fd5b8063ca101295146102e0578063cd499da314610300578063d513894814610320578063e3db8a49146103cd578063e5207453146103ed57600080fd5b80637edddf45116100e75780637edddf451461022357806386b98895146102435780638da5cb5b14610263578063a89d217314610290578063bffe1780146102c057600080fd5b806325416bc9146101675780632ccb03f214610189578063500de431146101a957806351e0e26b146101c9578063715018a61461020e57600080fd5b3661016257604080513381523460208201527f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4910160405180910390a1005b600080fd5b34801561017357600080fd5b50610187610182366004610fbd565b6104ad565b005b34801561019557600080fd5b506101876101a4366004611080565b6104f5565b3480156101b557600080fd5b506101876101c436600461105a565b610571565b3480156101d557600080fd5b506101f96101e4366004610f3f565b60016020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561021a57600080fd5b506101876105e3565b34801561022f57600080fd5b5061018761023e36600461105a565b61061e565b34801561024f57600080fd5b5061018761025e366004611080565b610672565b34801561026f57600080fd5b506102786106cd565b6040516001600160a01b039091168152602001610205565b34801561029c57600080fd5b506101f96102ab366004610f3f565b60026020526000908152604090205460ff1681565b3480156102cc57600080fd5b506101876102db3660046110a9565b6106dc565b3480156102ec57600080fd5b506101876102fb366004610fbd565b610746565b34801561030c57600080fd5b5061018761031b366004610fbd565b610781565b34801561032c57600080fd5b5060045460055461037f916001600160801b0381169163ffffffff600160801b830481169260ff600160a01b820416926001600160401b03600160a81b90920482169291811691600160401b9091041686565b604080516001600160801b03909716875263ffffffff958616602088015260ff909416938601939093526001600160401b0391821660608601521660808401521660a082015260c001610205565b3480156103d957600080fd5b506101f96103e8366004610f88565b6107da565b3480156103f957600080fd5b50610187610408366004610fbd565b610ac0565b34801561041957600080fd5b50600354610278906001600160a01b031681565b34801561043957600080fd5b50610187610448366004610f3f565b610b19565b34801561045957600080fd5b50610187610468366004610f3f565b610b92565b34801561047957600080fd5b50610187610488366004610f5c565b610c32565b34801561049957600080fd5b506101876104a8366004611031565b610d44565b336104b66106cd565b6001600160a01b0316146104e55760405162461bcd60e51b81526004016104dc906110cc565b60405180910390fd5b6104f182826001610db7565b5050565b336104fe6106cd565b6001600160a01b0316146105245760405162461bcd60e51b81526004016104dc906110cc565b6005805467ffffffffffffffff19166001600160401b03831617905560045b6040516001600160401b03831681526000805160206111d4833981519152906020015b60405180910390a250565b3361057a6106cd565b6001600160a01b0316146105a05760405162461bcd60e51b81526004016104dc906110cc565b6005805463ffffffff60401b1916600160401b63ffffffff8416021781555b60405163ffffffff831681526000805160206111d483398151915290602001610566565b336105ec6106cd565b6001600160a01b0316146106125760405162461bcd60e51b81526004016104dc906110cc565b61061c6000610e56565b565b336106276106cd565b6001600160a01b03161461064d5760405162461bcd60e51b81526004016104dc906110cc565b6004805463ffffffff60801b1916600160801b63ffffffff84160217905560016105bf565b3361067b6106cd565b6001600160a01b0316146106a15760405162461bcd60e51b81526004016104dc906110cc565b6004805467ffffffffffffffff60a81b1916600160a81b6001600160401b038416021790556003610543565b6000546001600160a01b031690565b336106e56106cd565b6001600160a01b03161461070b5760405162461bcd60e51b81526004016104dc906110cc565b6004805460ff60a01b1916600160a01b60ff841602179055600260405160ff831681526000805160206111d483398151915290602001610566565b3361074f6106cd565b6001600160a01b0316146107755760405162461bcd60e51b81526004016104dc906110cc565b6104f182826001610ea6565b6107896106cd565b6001600160a01b0316336001600160a01b031614806107b257506003546001600160a01b031633145b6107ce5760405162461bcd60e51b81526004016104dc90611101565b6104f182826000610db7565b6000805a905047806108345760035b60405186815233906001600160a01b038916907f2b8ae00e22d9eaf5a92820a22b947c007aee773fa36502ad7a1c9a464ab4932b9060200160405180910390a4600092505050610ab9565b3360009081526001602052604090205460ff166108525760006107e9565b6001600160a01b03861660009081526002602052604090205460ff166108795760016107e9565b60045460009061089990600160a81b90046001600160401b031648611129565b9050803a10156108a657503a5b6005546001600160401b0316158015906108ca57506005546001600160401b031681115b156108dd57506005546001600160401b03165b6004546005546001600160a01b03891631916001600160801b03811691600160401b900463ffffffff169061091c90600160a01b900460ff1689611141565b60045461093690600160801b900463ffffffff1688611129565b6109409190611129565b61094a908a611129565b98505a610957908a611160565b9850801580159061096757508089115b15610970578098505b600061097c8a86611141565b905082158015906109955750826109938286611129565b115b15610a0157828411156109f45760026040518b815233906001600160a01b038e16907f2b8ae00e22d9eaf5a92820a22b947c007aee773fa36502ad7a1c9a464ab4932b9060200160405180910390a46000975050505050505050610ab9565b6109fe8484611160565b90505b85811115610a0c5750845b6040516001600160a01b038c16908290600081818185875af1925050503d8060008114610a55576040519150601f19603f3d011682016040523d82523d6000602084013e610a5a565b606091505b5050604080518c8152602081018890529081018390529098508815159033906001600160a01b038e16907fd0224505f828ccfcbc56ca0590d97442e239a7aa770f712948fd6388356b20de9060600160405180910390a4505050505050505b9392505050565b610ac86106cd565b6001600160a01b0316336001600160a01b03161480610af157506003546001600160a01b031633145b610b0d5760405162461bcd60e51b81526004016104dc90611101565b6104f182826000610ea6565b33610b226106cd565b6001600160a01b031614610b485760405162461bcd60e51b81526004016104dc906110cc565b600380546001600160a01b0319166001600160a01b0383169081179091556040517fc388cec0895ad7ee4635898ec92207ca48d42256d4355f7042efef62c368a97990600090a250565b33610b9b6106cd565b6001600160a01b031614610bc15760405162461bcd60e51b81526004016104dc906110cc565b6001600160a01b038116610c265760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104dc565b610c2f81610e56565b50565b33610c3b6106cd565b6001600160a01b031614610c615760405162461bcd60e51b81526004016104dc906110cc565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610cae576040519150601f19603f3d011682016040523d82523d6000602084013e610cb3565b606091505b5050905080610cf65760405162461bcd60e51b815260206004820152600f60248201526e15d2551211149055d7d19052531151608a1b60448201526064016104dc565b604080513381526001600160a01b03851660208201529081018390527fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060600160405180910390a1505050565b33610d4d6106cd565b6001600160a01b031614610d735760405162461bcd60e51b81526004016104dc906110cc565b600480546001600160801b0319166001600160801b03831617905560006040516001600160801b03831681526000805160206111d483398151915290602001610566565b60005b82811015610e50576000848483818110610dd657610dd66111a8565b9050602002016020810190610deb9190610f3f565b6001600160a01b038116600081815260016020526040808220805460ff19168815159081179091559051939450927fb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c019190a35080610e4881611177565b915050610dba565b50505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005b82811015610e50576000848483818110610ec557610ec56111a8565b9050602002016020810190610eda9190610f3f565b6001600160a01b038116600081815260026020526040808220805460ff19168815159081179091559051939450927ff544cca9d5484bfd447775bd759d12d53f1aa7c5f770be82c55070798ff9c63e9190a35080610f3781611177565b915050610ea9565b600060208284031215610f5157600080fd5b8135610ab9816111be565b60008060408385031215610f6f57600080fd5b8235610f7a816111be565b946020939093013593505050565b600080600060608486031215610f9d57600080fd5b8335610fa8816111be565b95602085013595506040909401359392505050565b60008060208385031215610fd057600080fd5b82356001600160401b0380821115610fe757600080fd5b818501915085601f830112610ffb57600080fd5b81358181111561100a57600080fd5b8660208260051b850101111561101f57600080fd5b60209290920196919550909350505050565b60006020828403121561104357600080fd5b81356001600160801b0381168114610ab957600080fd5b60006020828403121561106c57600080fd5b813563ffffffff81168114610ab957600080fd5b60006020828403121561109257600080fd5b81356001600160401b0381168114610ab957600080fd5b6000602082840312156110bb57600080fd5b813560ff81168114610ab957600080fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252600e908201526d1393d517d055551213d49256915160921b604082015260600190565b6000821982111561113c5761113c611192565b500190565b600081600019048311821515161561115b5761115b611192565b500290565b60008282101561117257611172611192565b500390565b600060001982141561118b5761118b611192565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6001600160a01b0381168114610c2f57600080fdfeda79b6b81f905f788560507c685a42d5a8ab209ee26538cbcf3ce3caed601f9ba2646970667358221220b0592293f3edc1793a3acbe5f8f747e41220c935aa067f91c0d67e30827114cd64736f6c63430008070033",
}

// GasRefunderABI is the input ABI used to generate the binding from.
// Deprecated: Use GasRefunderMetaData.ABI instead.
var GasRefunderABI = GasRefunderMetaData.ABI

// GasRefunderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasRefunderMetaData.Bin instead.
var GasRefunderBin = GasRefunderMetaData.Bin

// DeployGasRefunder deploys a new Ethereum contract, binding an instance of GasRefunder to it.
func DeployGasRefunder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasRefunder, error) {
	parsed, err := GasRefunderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasRefunderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasRefunder{GasRefunderCaller: GasRefunderCaller{contract: contract}, GasRefunderTransactor: GasRefunderTransactor{contract: contract}, GasRefunderFilterer: GasRefunderFilterer{contract: contract}}, nil
}

// GasRefunder is an auto generated Go binding around an Ethereum contract.
type GasRefunder struct {
	GasRefunderCaller     // Read-only binding to the contract
	GasRefunderTransactor // Write-only binding to the contract
	GasRefunderFilterer   // Log filterer for contract events
}

// GasRefunderCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasRefunderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasRefunderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasRefunderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasRefunderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasRefunderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasRefunderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasRefunderSession struct {
	Contract     *GasRefunder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasRefunderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasRefunderCallerSession struct {
	Contract *GasRefunderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GasRefunderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasRefunderTransactorSession struct {
	Contract     *GasRefunderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GasRefunderRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasRefunderRaw struct {
	Contract *GasRefunder // Generic contract binding to access the raw methods on
}

// GasRefunderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasRefunderCallerRaw struct {
	Contract *GasRefunderCaller // Generic read-only contract binding to access the raw methods on
}

// GasRefunderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasRefunderTransactorRaw struct {
	Contract *GasRefunderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasRefunder creates a new instance of GasRefunder, bound to a specific deployed contract.
func NewGasRefunder(address common.Address, backend bind.ContractBackend) (*GasRefunder, error) {
	contract, err := bindGasRefunder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasRefunder{GasRefunderCaller: GasRefunderCaller{contract: contract}, GasRefunderTransactor: GasRefunderTransactor{contract: contract}, GasRefunderFilterer: GasRefunderFilterer{contract: contract}}, nil
}

// NewGasRefunderCaller creates a new read-only instance of GasRefunder, bound to a specific deployed contract.
func NewGasRefunderCaller(address common.Address, caller bind.ContractCaller) (*GasRefunderCaller, error) {
	contract, err := bindGasRefunder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasRefunderCaller{contract: contract}, nil
}

// NewGasRefunderTransactor creates a new write-only instance of GasRefunder, bound to a specific deployed contract.
func NewGasRefunderTransactor(address common.Address, transactor bind.ContractTransactor) (*GasRefunderTransactor, error) {
	contract, err := bindGasRefunder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasRefunderTransactor{contract: contract}, nil
}

// NewGasRefunderFilterer creates a new log filterer instance of GasRefunder, bound to a specific deployed contract.
func NewGasRefunderFilterer(address common.Address, filterer bind.ContractFilterer) (*GasRefunderFilterer, error) {
	contract, err := bindGasRefunder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasRefunderFilterer{contract: contract}, nil
}

// bindGasRefunder binds a generic wrapper to an already deployed contract.
func bindGasRefunder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasRefunderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasRefunder *GasRefunderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasRefunder.Contract.GasRefunderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasRefunder *GasRefunderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasRefunder.Contract.GasRefunderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasRefunder *GasRefunderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasRefunder.Contract.GasRefunderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasRefunder *GasRefunderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasRefunder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasRefunder *GasRefunderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasRefunder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasRefunder *GasRefunderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasRefunder.Contract.contract.Transact(opts, method, params...)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_GasRefunder *GasRefunderCaller) AllowedContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "allowedContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_GasRefunder *GasRefunderSession) AllowedContracts(arg0 common.Address) (bool, error) {
	return _GasRefunder.Contract.AllowedContracts(&_GasRefunder.CallOpts, arg0)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_GasRefunder *GasRefunderCallerSession) AllowedContracts(arg0 common.Address) (bool, error) {
	return _GasRefunder.Contract.AllowedContracts(&_GasRefunder.CallOpts, arg0)
}

// AllowedRefundees is a free data retrieval call binding the contract method 0xa89d2173.
//
// Solidity: function allowedRefundees(address ) view returns(bool)
func (_GasRefunder *GasRefunderCaller) AllowedRefundees(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "allowedRefundees", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedRefundees is a free data retrieval call binding the contract method 0xa89d2173.
//
// Solidity: function allowedRefundees(address ) view returns(bool)
func (_GasRefunder *GasRefunderSession) AllowedRefundees(arg0 common.Address) (bool, error) {
	return _GasRefunder.Contract.AllowedRefundees(&_GasRefunder.CallOpts, arg0)
}

// AllowedRefundees is a free data retrieval call binding the contract method 0xa89d2173.
//
// Solidity: function allowedRefundees(address ) view returns(bool)
func (_GasRefunder *GasRefunderCallerSession) AllowedRefundees(arg0 common.Address) (bool, error) {
	return _GasRefunder.Contract.AllowedRefundees(&_GasRefunder.CallOpts, arg0)
}

// CommonParams is a free data retrieval call binding the contract method 0xd5138948.
//
// Solidity: function commonParams() view returns(uint128 maxRefundeeBalance, uint32 extraGasMargin, uint8 calldataCost, uint64 maxGasTip, uint64 maxGasCost, uint32 maxSingleGasUsage)
func (_GasRefunder *GasRefunderCaller) CommonParams(opts *bind.CallOpts) (struct {
	MaxRefundeeBalance *big.Int
	ExtraGasMargin     uint32
	CalldataCost       uint8
	MaxGasTip          uint64
	MaxGasCost         uint64
	MaxSingleGasUsage  uint32
}, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "commonParams")

	outstruct := new(struct {
		MaxRefundeeBalance *big.Int
		ExtraGasMargin     uint32
		CalldataCost       uint8
		MaxGasTip          uint64
		MaxGasCost         uint64
		MaxSingleGasUsage  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxRefundeeBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExtraGasMargin = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.CalldataCost = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.MaxGasTip = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.MaxGasCost = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.MaxSingleGasUsage = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// CommonParams is a free data retrieval call binding the contract method 0xd5138948.
//
// Solidity: function commonParams() view returns(uint128 maxRefundeeBalance, uint32 extraGasMargin, uint8 calldataCost, uint64 maxGasTip, uint64 maxGasCost, uint32 maxSingleGasUsage)
func (_GasRefunder *GasRefunderSession) CommonParams() (struct {
	MaxRefundeeBalance *big.Int
	ExtraGasMargin     uint32
	CalldataCost       uint8
	MaxGasTip          uint64
	MaxGasCost         uint64
	MaxSingleGasUsage  uint32
}, error) {
	return _GasRefunder.Contract.CommonParams(&_GasRefunder.CallOpts)
}

// CommonParams is a free data retrieval call binding the contract method 0xd5138948.
//
// Solidity: function commonParams() view returns(uint128 maxRefundeeBalance, uint32 extraGasMargin, uint8 calldataCost, uint64 maxGasTip, uint64 maxGasCost, uint32 maxSingleGasUsage)
func (_GasRefunder *GasRefunderCallerSession) CommonParams() (struct {
	MaxRefundeeBalance *big.Int
	ExtraGasMargin     uint32
	CalldataCost       uint8
	MaxGasTip          uint64
	MaxGasCost         uint64
	MaxSingleGasUsage  uint32
}, error) {
	return _GasRefunder.Contract.CommonParams(&_GasRefunder.CallOpts)
}

// Disallower is a free data retrieval call binding the contract method 0xefe12b01.
//
// Solidity: function disallower() view returns(address)
func (_GasRefunder *GasRefunderCaller) Disallower(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "disallower")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Disallower is a free data retrieval call binding the contract method 0xefe12b01.
//
// Solidity: function disallower() view returns(address)
func (_GasRefunder *GasRefunderSession) Disallower() (common.Address, error) {
	return _GasRefunder.Contract.Disallower(&_GasRefunder.CallOpts)
}

// Disallower is a free data retrieval call binding the contract method 0xefe12b01.
//
// Solidity: function disallower() view returns(address)
func (_GasRefunder *GasRefunderCallerSession) Disallower() (common.Address, error) {
	return _GasRefunder.Contract.Disallower(&_GasRefunder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasRefunder *GasRefunderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasRefunder *GasRefunderSession) Owner() (common.Address, error) {
	return _GasRefunder.Contract.Owner(&_GasRefunder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasRefunder *GasRefunderCallerSession) Owner() (common.Address, error) {
	return _GasRefunder.Contract.Owner(&_GasRefunder.CallOpts)
}

// AllowContracts is a paid mutator transaction binding the contract method 0x25416bc9.
//
// Solidity: function allowContracts(address[] addresses) returns()
func (_GasRefunder *GasRefunderTransactor) AllowContracts(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "allowContracts", addresses)
}

// AllowContracts is a paid mutator transaction binding the contract method 0x25416bc9.
//
// Solidity: function allowContracts(address[] addresses) returns()
func (_GasRefunder *GasRefunderSession) AllowContracts(addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.AllowContracts(&_GasRefunder.TransactOpts, addresses)
}

// AllowContracts is a paid mutator transaction binding the contract method 0x25416bc9.
//
// Solidity: function allowContracts(address[] addresses) returns()
func (_GasRefunder *GasRefunderTransactorSession) AllowContracts(addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.AllowContracts(&_GasRefunder.TransactOpts, addresses)
}

// AllowRefundees is a paid mutator transaction binding the contract method 0xca101295.
//
// Solidity: function allowRefundees(address[] addresses) returns()
func (_GasRefunder *GasRefunderTransactor) AllowRefundees(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "allowRefundees", addresses)
}

// AllowRefundees is a paid mutator transaction binding the contract method 0xca101295.
//
// Solidity: function allowRefundees(address[] addresses) returns()
func (_GasRefunder *GasRefunderSession) AllowRefundees(addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.AllowRefundees(&_GasRefunder.TransactOpts, addresses)
}

// AllowRefundees is a paid mutator transaction binding the contract method 0xca101295.
//
// Solidity: function allowRefundees(address[] addresses) returns()
func (_GasRefunder *GasRefunderTransactorSession) AllowRefundees(addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.AllowRefundees(&_GasRefunder.TransactOpts, addresses)
}

// DisallowContracts is a paid mutator transaction binding the contract method 0xcd499da3.
//
// Solidity: function disallowContracts(address[] addresses) returns()
func (_GasRefunder *GasRefunderTransactor) DisallowContracts(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "disallowContracts", addresses)
}

// DisallowContracts is a paid mutator transaction binding the contract method 0xcd499da3.
//
// Solidity: function disallowContracts(address[] addresses) returns()
func (_GasRefunder *GasRefunderSession) DisallowContracts(addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.DisallowContracts(&_GasRefunder.TransactOpts, addresses)
}

// DisallowContracts is a paid mutator transaction binding the contract method 0xcd499da3.
//
// Solidity: function disallowContracts(address[] addresses) returns()
func (_GasRefunder *GasRefunderTransactorSession) DisallowContracts(addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.DisallowContracts(&_GasRefunder.TransactOpts, addresses)
}

// DisallowRefundees is a paid mutator transaction binding the contract method 0xe5207453.
//
// Solidity: function disallowRefundees(address[] addresses) returns()
func (_GasRefunder *GasRefunderTransactor) DisallowRefundees(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "disallowRefundees", addresses)
}

// DisallowRefundees is a paid mutator transaction binding the contract method 0xe5207453.
//
// Solidity: function disallowRefundees(address[] addresses) returns()
func (_GasRefunder *GasRefunderSession) DisallowRefundees(addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.DisallowRefundees(&_GasRefunder.TransactOpts, addresses)
}

// DisallowRefundees is a paid mutator transaction binding the contract method 0xe5207453.
//
// Solidity: function disallowRefundees(address[] addresses) returns()
func (_GasRefunder *GasRefunderTransactorSession) DisallowRefundees(addresses []common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.DisallowRefundees(&_GasRefunder.TransactOpts, addresses)
}

// OnGasSpent is a paid mutator transaction binding the contract method 0xe3db8a49.
//
// Solidity: function onGasSpent(address refundee, uint256 gasUsed, uint256 calldataSize) returns(bool success)
func (_GasRefunder *GasRefunderTransactor) OnGasSpent(opts *bind.TransactOpts, refundee common.Address, gasUsed *big.Int, calldataSize *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "onGasSpent", refundee, gasUsed, calldataSize)
}

// OnGasSpent is a paid mutator transaction binding the contract method 0xe3db8a49.
//
// Solidity: function onGasSpent(address refundee, uint256 gasUsed, uint256 calldataSize) returns(bool success)
func (_GasRefunder *GasRefunderSession) OnGasSpent(refundee common.Address, gasUsed *big.Int, calldataSize *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.OnGasSpent(&_GasRefunder.TransactOpts, refundee, gasUsed, calldataSize)
}

// OnGasSpent is a paid mutator transaction binding the contract method 0xe3db8a49.
//
// Solidity: function onGasSpent(address refundee, uint256 gasUsed, uint256 calldataSize) returns(bool success)
func (_GasRefunder *GasRefunderTransactorSession) OnGasSpent(refundee common.Address, gasUsed *big.Int, calldataSize *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.OnGasSpent(&_GasRefunder.TransactOpts, refundee, gasUsed, calldataSize)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasRefunder *GasRefunderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasRefunder *GasRefunderSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasRefunder.Contract.RenounceOwnership(&_GasRefunder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasRefunder *GasRefunderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasRefunder.Contract.RenounceOwnership(&_GasRefunder.TransactOpts)
}

// SetCalldataCost is a paid mutator transaction binding the contract method 0xbffe1780.
//
// Solidity: function setCalldataCost(uint8 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetCalldataCost(opts *bind.TransactOpts, newValue uint8) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setCalldataCost", newValue)
}

// SetCalldataCost is a paid mutator transaction binding the contract method 0xbffe1780.
//
// Solidity: function setCalldataCost(uint8 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetCalldataCost(newValue uint8) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetCalldataCost(&_GasRefunder.TransactOpts, newValue)
}

// SetCalldataCost is a paid mutator transaction binding the contract method 0xbffe1780.
//
// Solidity: function setCalldataCost(uint8 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetCalldataCost(newValue uint8) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetCalldataCost(&_GasRefunder.TransactOpts, newValue)
}

// SetDisallower is a paid mutator transaction binding the contract method 0xf1e845ca.
//
// Solidity: function setDisallower(address addr) returns()
func (_GasRefunder *GasRefunderTransactor) SetDisallower(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setDisallower", addr)
}

// SetDisallower is a paid mutator transaction binding the contract method 0xf1e845ca.
//
// Solidity: function setDisallower(address addr) returns()
func (_GasRefunder *GasRefunderSession) SetDisallower(addr common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetDisallower(&_GasRefunder.TransactOpts, addr)
}

// SetDisallower is a paid mutator transaction binding the contract method 0xf1e845ca.
//
// Solidity: function setDisallower(address addr) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetDisallower(addr common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetDisallower(&_GasRefunder.TransactOpts, addr)
}

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0x7edddf45.
//
// Solidity: function setExtraGasMargin(uint32 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetExtraGasMargin(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setExtraGasMargin", newValue)
}

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0x7edddf45.
//
// Solidity: function setExtraGasMargin(uint32 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetExtraGasMargin(newValue uint32) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetExtraGasMargin(&_GasRefunder.TransactOpts, newValue)
}

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0x7edddf45.
//
// Solidity: function setExtraGasMargin(uint32 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetExtraGasMargin(newValue uint32) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetExtraGasMargin(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x2ccb03f2.
//
// Solidity: function setMaxGasCost(uint64 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxGasCost(opts *bind.TransactOpts, newValue uint64) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxGasCost", newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x2ccb03f2.
//
// Solidity: function setMaxGasCost(uint64 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxGasCost(newValue uint64) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasCost(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x2ccb03f2.
//
// Solidity: function setMaxGasCost(uint64 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxGasCost(newValue uint64) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasCost(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0x86b98895.
//
// Solidity: function setMaxGasTip(uint64 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxGasTip(opts *bind.TransactOpts, newValue uint64) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxGasTip", newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0x86b98895.
//
// Solidity: function setMaxGasTip(uint64 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxGasTip(newValue uint64) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasTip(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0x86b98895.
//
// Solidity: function setMaxGasTip(uint64 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxGasTip(newValue uint64) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasTip(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0xf52128eb.
//
// Solidity: function setMaxRefundeeBalance(uint128 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxRefundeeBalance(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxRefundeeBalance", newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0xf52128eb.
//
// Solidity: function setMaxRefundeeBalance(uint128 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxRefundeeBalance(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxRefundeeBalance(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0xf52128eb.
//
// Solidity: function setMaxRefundeeBalance(uint128 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxRefundeeBalance(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxRefundeeBalance(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x500de431.
//
// Solidity: function setMaxSingleGasUsage(uint32 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxSingleGasUsage(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxSingleGasUsage", newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x500de431.
//
// Solidity: function setMaxSingleGasUsage(uint32 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxSingleGasUsage(newValue uint32) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxSingleGasUsage(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x500de431.
//
// Solidity: function setMaxSingleGasUsage(uint32 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxSingleGasUsage(newValue uint32) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxSingleGasUsage(&_GasRefunder.TransactOpts, newValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasRefunder *GasRefunderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasRefunder *GasRefunderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.TransferOwnership(&_GasRefunder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasRefunder *GasRefunderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasRefunder.Contract.TransferOwnership(&_GasRefunder.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address destination, uint256 amount) returns()
func (_GasRefunder *GasRefunderTransactor) Withdraw(opts *bind.TransactOpts, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "withdraw", destination, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address destination, uint256 amount) returns()
func (_GasRefunder *GasRefunderSession) Withdraw(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.Withdraw(&_GasRefunder.TransactOpts, destination, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address destination, uint256 amount) returns()
func (_GasRefunder *GasRefunderTransactorSession) Withdraw(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.Withdraw(&_GasRefunder.TransactOpts, destination, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasRefunder *GasRefunderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasRefunder.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasRefunder *GasRefunderSession) Receive() (*types.Transaction, error) {
	return _GasRefunder.Contract.Receive(&_GasRefunder.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasRefunder *GasRefunderTransactorSession) Receive() (*types.Transaction, error) {
	return _GasRefunder.Contract.Receive(&_GasRefunder.TransactOpts)
}

// GasRefunderCommonParameterSetIterator is returned from FilterCommonParameterSet and is used to iterate over the raw logs and unpacked data for CommonParameterSet events raised by the GasRefunder contract.
type GasRefunderCommonParameterSetIterator struct {
	Event *GasRefunderCommonParameterSet // Event containing the contract specifics and raw log

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
func (it *GasRefunderCommonParameterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderCommonParameterSet)
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
		it.Event = new(GasRefunderCommonParameterSet)
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
func (it *GasRefunderCommonParameterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderCommonParameterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderCommonParameterSet represents a CommonParameterSet event raised by the GasRefunder contract.
type GasRefunderCommonParameterSet struct {
	Parameter uint8
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommonParameterSet is a free log retrieval operation binding the contract event 0xda79b6b81f905f788560507c685a42d5a8ab209ee26538cbcf3ce3caed601f9b.
//
// Solidity: event CommonParameterSet(uint8 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) FilterCommonParameterSet(opts *bind.FilterOpts, parameter []uint8) (*GasRefunderCommonParameterSetIterator, error) {

	var parameterRule []interface{}
	for _, parameterItem := range parameter {
		parameterRule = append(parameterRule, parameterItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "CommonParameterSet", parameterRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderCommonParameterSetIterator{contract: _GasRefunder.contract, event: "CommonParameterSet", logs: logs, sub: sub}, nil
}

// WatchCommonParameterSet is a free log subscription operation binding the contract event 0xda79b6b81f905f788560507c685a42d5a8ab209ee26538cbcf3ce3caed601f9b.
//
// Solidity: event CommonParameterSet(uint8 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) WatchCommonParameterSet(opts *bind.WatchOpts, sink chan<- *GasRefunderCommonParameterSet, parameter []uint8) (event.Subscription, error) {

	var parameterRule []interface{}
	for _, parameterItem := range parameter {
		parameterRule = append(parameterRule, parameterItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "CommonParameterSet", parameterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderCommonParameterSet)
				if err := _GasRefunder.contract.UnpackLog(event, "CommonParameterSet", log); err != nil {
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

// ParseCommonParameterSet is a log parse operation binding the contract event 0xda79b6b81f905f788560507c685a42d5a8ab209ee26538cbcf3ce3caed601f9b.
//
// Solidity: event CommonParameterSet(uint8 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) ParseCommonParameterSet(log types.Log) (*GasRefunderCommonParameterSet, error) {
	event := new(GasRefunderCommonParameterSet)
	if err := _GasRefunder.contract.UnpackLog(event, "CommonParameterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderContractAllowedSetIterator is returned from FilterContractAllowedSet and is used to iterate over the raw logs and unpacked data for ContractAllowedSet events raised by the GasRefunder contract.
type GasRefunderContractAllowedSetIterator struct {
	Event *GasRefunderContractAllowedSet // Event containing the contract specifics and raw log

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
func (it *GasRefunderContractAllowedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderContractAllowedSet)
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
		it.Event = new(GasRefunderContractAllowedSet)
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
func (it *GasRefunderContractAllowedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderContractAllowedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderContractAllowedSet represents a ContractAllowedSet event raised by the GasRefunder contract.
type GasRefunderContractAllowedSet struct {
	Addr    common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractAllowedSet is a free log retrieval operation binding the contract event 0xb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c01.
//
// Solidity: event ContractAllowedSet(address indexed addr, bool indexed allowed)
func (_GasRefunder *GasRefunderFilterer) FilterContractAllowedSet(opts *bind.FilterOpts, addr []common.Address, allowed []bool) (*GasRefunderContractAllowedSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "ContractAllowedSet", addrRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderContractAllowedSetIterator{contract: _GasRefunder.contract, event: "ContractAllowedSet", logs: logs, sub: sub}, nil
}

// WatchContractAllowedSet is a free log subscription operation binding the contract event 0xb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c01.
//
// Solidity: event ContractAllowedSet(address indexed addr, bool indexed allowed)
func (_GasRefunder *GasRefunderFilterer) WatchContractAllowedSet(opts *bind.WatchOpts, sink chan<- *GasRefunderContractAllowedSet, addr []common.Address, allowed []bool) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "ContractAllowedSet", addrRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderContractAllowedSet)
				if err := _GasRefunder.contract.UnpackLog(event, "ContractAllowedSet", log); err != nil {
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

// ParseContractAllowedSet is a log parse operation binding the contract event 0xb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c01.
//
// Solidity: event ContractAllowedSet(address indexed addr, bool indexed allowed)
func (_GasRefunder *GasRefunderFilterer) ParseContractAllowedSet(log types.Log) (*GasRefunderContractAllowedSet, error) {
	event := new(GasRefunderContractAllowedSet)
	if err := _GasRefunder.contract.UnpackLog(event, "ContractAllowedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GasRefunder contract.
type GasRefunderDepositedIterator struct {
	Event *GasRefunderDeposited // Event containing the contract specifics and raw log

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
func (it *GasRefunderDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderDeposited)
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
		it.Event = new(GasRefunderDeposited)
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
func (it *GasRefunderDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderDeposited represents a Deposited event raised by the GasRefunder contract.
type GasRefunderDeposited struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) FilterDeposited(opts *bind.FilterOpts) (*GasRefunderDepositedIterator, error) {

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &GasRefunderDepositedIterator{contract: _GasRefunder.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GasRefunderDeposited) (event.Subscription, error) {

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderDeposited)
				if err := _GasRefunder.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address sender, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) ParseDeposited(log types.Log) (*GasRefunderDeposited, error) {
	event := new(GasRefunderDeposited)
	if err := _GasRefunder.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderDisallowerSetIterator is returned from FilterDisallowerSet and is used to iterate over the raw logs and unpacked data for DisallowerSet events raised by the GasRefunder contract.
type GasRefunderDisallowerSetIterator struct {
	Event *GasRefunderDisallowerSet // Event containing the contract specifics and raw log

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
func (it *GasRefunderDisallowerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderDisallowerSet)
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
		it.Event = new(GasRefunderDisallowerSet)
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
func (it *GasRefunderDisallowerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderDisallowerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderDisallowerSet represents a DisallowerSet event raised by the GasRefunder contract.
type GasRefunderDisallowerSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDisallowerSet is a free log retrieval operation binding the contract event 0xc388cec0895ad7ee4635898ec92207ca48d42256d4355f7042efef62c368a979.
//
// Solidity: event DisallowerSet(address indexed addr)
func (_GasRefunder *GasRefunderFilterer) FilterDisallowerSet(opts *bind.FilterOpts, addr []common.Address) (*GasRefunderDisallowerSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "DisallowerSet", addrRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderDisallowerSetIterator{contract: _GasRefunder.contract, event: "DisallowerSet", logs: logs, sub: sub}, nil
}

// WatchDisallowerSet is a free log subscription operation binding the contract event 0xc388cec0895ad7ee4635898ec92207ca48d42256d4355f7042efef62c368a979.
//
// Solidity: event DisallowerSet(address indexed addr)
func (_GasRefunder *GasRefunderFilterer) WatchDisallowerSet(opts *bind.WatchOpts, sink chan<- *GasRefunderDisallowerSet, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "DisallowerSet", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderDisallowerSet)
				if err := _GasRefunder.contract.UnpackLog(event, "DisallowerSet", log); err != nil {
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

// ParseDisallowerSet is a log parse operation binding the contract event 0xc388cec0895ad7ee4635898ec92207ca48d42256d4355f7042efef62c368a979.
//
// Solidity: event DisallowerSet(address indexed addr)
func (_GasRefunder *GasRefunderFilterer) ParseDisallowerSet(log types.Log) (*GasRefunderDisallowerSet, error) {
	event := new(GasRefunderDisallowerSet)
	if err := _GasRefunder.contract.UnpackLog(event, "DisallowerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GasRefunder contract.
type GasRefunderOwnershipTransferredIterator struct {
	Event *GasRefunderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GasRefunderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderOwnershipTransferred)
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
		it.Event = new(GasRefunderOwnershipTransferred)
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
func (it *GasRefunderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderOwnershipTransferred represents a OwnershipTransferred event raised by the GasRefunder contract.
type GasRefunderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasRefunder *GasRefunderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasRefunderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderOwnershipTransferredIterator{contract: _GasRefunder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasRefunder *GasRefunderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasRefunderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderOwnershipTransferred)
				if err := _GasRefunder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GasRefunder *GasRefunderFilterer) ParseOwnershipTransferred(log types.Log) (*GasRefunderOwnershipTransferred, error) {
	event := new(GasRefunderOwnershipTransferred)
	if err := _GasRefunder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderRefundGasCostsDeniedIterator is returned from FilterRefundGasCostsDenied and is used to iterate over the raw logs and unpacked data for RefundGasCostsDenied events raised by the GasRefunder contract.
type GasRefunderRefundGasCostsDeniedIterator struct {
	Event *GasRefunderRefundGasCostsDenied // Event containing the contract specifics and raw log

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
func (it *GasRefunderRefundGasCostsDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderRefundGasCostsDenied)
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
		it.Event = new(GasRefunderRefundGasCostsDenied)
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
func (it *GasRefunderRefundGasCostsDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderRefundGasCostsDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderRefundGasCostsDenied represents a RefundGasCostsDenied event raised by the GasRefunder contract.
type GasRefunderRefundGasCostsDenied struct {
	Refundee        common.Address
	ContractAddress common.Address
	Reason          uint8
	Gas             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefundGasCostsDenied is a free log retrieval operation binding the contract event 0x2b8ae00e22d9eaf5a92820a22b947c007aee773fa36502ad7a1c9a464ab4932b.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint8 indexed reason, uint256 gas)
func (_GasRefunder *GasRefunderFilterer) FilterRefundGasCostsDenied(opts *bind.FilterOpts, refundee []common.Address, contractAddress []common.Address, reason []uint8) (*GasRefunderRefundGasCostsDeniedIterator, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "RefundGasCostsDenied", refundeeRule, contractAddressRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderRefundGasCostsDeniedIterator{contract: _GasRefunder.contract, event: "RefundGasCostsDenied", logs: logs, sub: sub}, nil
}

// WatchRefundGasCostsDenied is a free log subscription operation binding the contract event 0x2b8ae00e22d9eaf5a92820a22b947c007aee773fa36502ad7a1c9a464ab4932b.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint8 indexed reason, uint256 gas)
func (_GasRefunder *GasRefunderFilterer) WatchRefundGasCostsDenied(opts *bind.WatchOpts, sink chan<- *GasRefunderRefundGasCostsDenied, refundee []common.Address, contractAddress []common.Address, reason []uint8) (event.Subscription, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "RefundGasCostsDenied", refundeeRule, contractAddressRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderRefundGasCostsDenied)
				if err := _GasRefunder.contract.UnpackLog(event, "RefundGasCostsDenied", log); err != nil {
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

// ParseRefundGasCostsDenied is a log parse operation binding the contract event 0x2b8ae00e22d9eaf5a92820a22b947c007aee773fa36502ad7a1c9a464ab4932b.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint8 indexed reason, uint256 gas)
func (_GasRefunder *GasRefunderFilterer) ParseRefundGasCostsDenied(log types.Log) (*GasRefunderRefundGasCostsDenied, error) {
	event := new(GasRefunderRefundGasCostsDenied)
	if err := _GasRefunder.contract.UnpackLog(event, "RefundGasCostsDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderRefundedGasCostsIterator is returned from FilterRefundedGasCosts and is used to iterate over the raw logs and unpacked data for RefundedGasCosts events raised by the GasRefunder contract.
type GasRefunderRefundedGasCostsIterator struct {
	Event *GasRefunderRefundedGasCosts // Event containing the contract specifics and raw log

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
func (it *GasRefunderRefundedGasCostsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderRefundedGasCosts)
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
		it.Event = new(GasRefunderRefundedGasCosts)
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
func (it *GasRefunderRefundedGasCostsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderRefundedGasCostsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderRefundedGasCosts represents a RefundedGasCosts event raised by the GasRefunder contract.
type GasRefunderRefundedGasCosts struct {
	Refundee        common.Address
	ContractAddress common.Address
	Success         bool
	Gas             *big.Int
	GasPrice        *big.Int
	AmountPaid      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefundedGasCosts is a free log retrieval operation binding the contract event 0xd0224505f828ccfcbc56ca0590d97442e239a7aa770f712948fd6388356b20de.
//
// Solidity: event RefundedGasCosts(address indexed refundee, address indexed contractAddress, bool indexed success, uint256 gas, uint256 gasPrice, uint256 amountPaid)
func (_GasRefunder *GasRefunderFilterer) FilterRefundedGasCosts(opts *bind.FilterOpts, refundee []common.Address, contractAddress []common.Address, success []bool) (*GasRefunderRefundedGasCostsIterator, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "RefundedGasCosts", refundeeRule, contractAddressRule, successRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderRefundedGasCostsIterator{contract: _GasRefunder.contract, event: "RefundedGasCosts", logs: logs, sub: sub}, nil
}

// WatchRefundedGasCosts is a free log subscription operation binding the contract event 0xd0224505f828ccfcbc56ca0590d97442e239a7aa770f712948fd6388356b20de.
//
// Solidity: event RefundedGasCosts(address indexed refundee, address indexed contractAddress, bool indexed success, uint256 gas, uint256 gasPrice, uint256 amountPaid)
func (_GasRefunder *GasRefunderFilterer) WatchRefundedGasCosts(opts *bind.WatchOpts, sink chan<- *GasRefunderRefundedGasCosts, refundee []common.Address, contractAddress []common.Address, success []bool) (event.Subscription, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "RefundedGasCosts", refundeeRule, contractAddressRule, successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderRefundedGasCosts)
				if err := _GasRefunder.contract.UnpackLog(event, "RefundedGasCosts", log); err != nil {
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

// ParseRefundedGasCosts is a log parse operation binding the contract event 0xd0224505f828ccfcbc56ca0590d97442e239a7aa770f712948fd6388356b20de.
//
// Solidity: event RefundedGasCosts(address indexed refundee, address indexed contractAddress, bool indexed success, uint256 gas, uint256 gasPrice, uint256 amountPaid)
func (_GasRefunder *GasRefunderFilterer) ParseRefundedGasCosts(log types.Log) (*GasRefunderRefundedGasCosts, error) {
	event := new(GasRefunderRefundedGasCosts)
	if err := _GasRefunder.contract.UnpackLog(event, "RefundedGasCosts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderRefundeeAllowedSetIterator is returned from FilterRefundeeAllowedSet and is used to iterate over the raw logs and unpacked data for RefundeeAllowedSet events raised by the GasRefunder contract.
type GasRefunderRefundeeAllowedSetIterator struct {
	Event *GasRefunderRefundeeAllowedSet // Event containing the contract specifics and raw log

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
func (it *GasRefunderRefundeeAllowedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderRefundeeAllowedSet)
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
		it.Event = new(GasRefunderRefundeeAllowedSet)
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
func (it *GasRefunderRefundeeAllowedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderRefundeeAllowedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderRefundeeAllowedSet represents a RefundeeAllowedSet event raised by the GasRefunder contract.
type GasRefunderRefundeeAllowedSet struct {
	Addr    common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefundeeAllowedSet is a free log retrieval operation binding the contract event 0xf544cca9d5484bfd447775bd759d12d53f1aa7c5f770be82c55070798ff9c63e.
//
// Solidity: event RefundeeAllowedSet(address indexed addr, bool indexed allowed)
func (_GasRefunder *GasRefunderFilterer) FilterRefundeeAllowedSet(opts *bind.FilterOpts, addr []common.Address, allowed []bool) (*GasRefunderRefundeeAllowedSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "RefundeeAllowedSet", addrRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderRefundeeAllowedSetIterator{contract: _GasRefunder.contract, event: "RefundeeAllowedSet", logs: logs, sub: sub}, nil
}

// WatchRefundeeAllowedSet is a free log subscription operation binding the contract event 0xf544cca9d5484bfd447775bd759d12d53f1aa7c5f770be82c55070798ff9c63e.
//
// Solidity: event RefundeeAllowedSet(address indexed addr, bool indexed allowed)
func (_GasRefunder *GasRefunderFilterer) WatchRefundeeAllowedSet(opts *bind.WatchOpts, sink chan<- *GasRefunderRefundeeAllowedSet, addr []common.Address, allowed []bool) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "RefundeeAllowedSet", addrRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderRefundeeAllowedSet)
				if err := _GasRefunder.contract.UnpackLog(event, "RefundeeAllowedSet", log); err != nil {
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

// ParseRefundeeAllowedSet is a log parse operation binding the contract event 0xf544cca9d5484bfd447775bd759d12d53f1aa7c5f770be82c55070798ff9c63e.
//
// Solidity: event RefundeeAllowedSet(address indexed addr, bool indexed allowed)
func (_GasRefunder *GasRefunderFilterer) ParseRefundeeAllowedSet(log types.Log) (*GasRefunderRefundeeAllowedSet, error) {
	event := new(GasRefunderRefundeeAllowedSet)
	if err := _GasRefunder.contract.UnpackLog(event, "RefundeeAllowedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasRefunderWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GasRefunder contract.
type GasRefunderWithdrawnIterator struct {
	Event *GasRefunderWithdrawn // Event containing the contract specifics and raw log

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
func (it *GasRefunderWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderWithdrawn)
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
		it.Event = new(GasRefunderWithdrawn)
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
func (it *GasRefunderWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderWithdrawn represents a Withdrawn event raised by the GasRefunder contract.
type GasRefunderWithdrawn struct {
	Initiator   common.Address
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address initiator, address destination, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*GasRefunderWithdrawnIterator, error) {

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &GasRefunderWithdrawnIterator{contract: _GasRefunder.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address initiator, address destination, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GasRefunderWithdrawn) (event.Subscription, error) {

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderWithdrawn)
				if err := _GasRefunder.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address initiator, address destination, uint256 amount)
func (_GasRefunder *GasRefunderFilterer) ParseWithdrawn(log types.Log) (*GasRefunderWithdrawn, error) {
	event := new(GasRefunderWithdrawn)
	if err := _GasRefunder.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
