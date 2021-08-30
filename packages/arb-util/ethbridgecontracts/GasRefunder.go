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

// GasRefunderABI is the input ABI used to generate the binding from.
const GasRefunderABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumGasRefunder.CommonParameterKey\",\"name\":\"parameter\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CommonParameterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"ContractAllowedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DisallowerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumGasRefunder.RefundDenyReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"RefundGasCostsDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"}],\"name\":\"RefundedGasCosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"RefundeeAllowedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"allowContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"allowRefundees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedRefundees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commonParams\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"maxRefundeeBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"extraGasMargin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"calldataCost\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"maxGasTip\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasCost\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxSingleGasUsage\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"disallowContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"disallowRefundees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disallower\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastContractRefund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"refundee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"onGasSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newValue\",\"type\":\"uint8\"}],\"name\":\"setCalldataCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDisallower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setExtraGasMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValue\",\"type\":\"uint64\"}],\"name\":\"setMaxGasCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValue\",\"type\":\"uint64\"}],\"name\":\"setMaxGasTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newValue\",\"type\":\"uint128\"}],\"name\":\"setMaxRefundeeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setMaxSingleGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// GasRefunderBin is the compiled bytecode used for deploying new contracts.
var GasRefunderBin = "0x608060405234801561001057600080fd5b5061001a336100a5565b6040805160c08101825260008152610fa06020820152600c9181019190915263773594006060820152641bf08eb0006080820152621e848060a090910152600580546001600160e81b03191678773594000c00000fa000000000000000000000000000000000179055600680546001600160601b0319166a1e84800000001bf08eb0001790556100f5565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61129e806101046000396000f3fe60806040526004361061012e5760003560e01c8063bffe1780116100ab578063e52074531161006f578063e520745314610433578063efe12b0114610453578063f1e845ca14610473578063f2fde38b14610493578063f3fef3a3146104b3578063f52128eb146104d357600080fd5b8063bffe178014610306578063ca10129514610326578063cd499da314610346578063d513894814610366578063e3db8a491461041357600080fd5b80637edddf45116100f25780637edddf451461022e57806386b988951461024e5780638da5cb5b1461026e578063a89d21731461029b578063bddaf01d146102cb57600080fd5b806325416bc9146101725780632ccb03f214610194578063500de431146101b457806351e0e26b146101d4578063715018a61461021957600080fd5b3661016d57604080513381523460208201527f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4910160405180910390a1005b600080fd5b34801561017e57600080fd5b5061019261018d366004611032565b6104f3565b005b3480156101a057600080fd5b506101926101af3660046110f5565b61053b565b3480156101c057600080fd5b506101926101cf3660046110cf565b6105b7565b3480156101e057600080fd5b506102046101ef366004610fb4565b60016020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561022557600080fd5b5061019261062b565b34801561023a57600080fd5b506101926102493660046110cf565b610666565b34801561025a57600080fd5b506101926102693660046110f5565b6106ba565b34801561027a57600080fd5b50610283610715565b6040516001600160a01b039091168152602001610210565b3480156102a757600080fd5b506102046102b6366004610fb4565b60026020526000908152604090205460ff1681565b3480156102d757600080fd5b506102f86102e6366004610fb4565b60036020526000908152604090205481565b604051908152602001610210565b34801561031257600080fd5b5061019261032136600461111e565b610724565b34801561033257600080fd5b50610192610341366004611032565b61078e565b34801561035257600080fd5b50610192610361366004611032565b6107c9565b34801561037257600080fd5b506005546006546103c5916001600160801b0381169163ffffffff600160801b830481169260ff600160a01b820416926001600160401b03600160a81b90920482169291811691600160401b9091041686565b604080516001600160801b03909716875263ffffffff958616602088015260ff909416938601939093526001600160401b0391821660608601521660808401521660a082015260c001610210565b34801561041f57600080fd5b5061020461042e366004610ffd565b610822565b34801561043f57600080fd5b5061019261044e366004611032565b610b35565b34801561045f57600080fd5b50600454610283906001600160a01b031681565b34801561047f57600080fd5b5061019261048e366004610fb4565b610b8e565b34801561049f57600080fd5b506101926104ae366004610fb4565b610c07565b3480156104bf57600080fd5b506101926104ce366004610fd1565b610ca7565b3480156104df57600080fd5b506101926104ee3660046110a6565b610db9565b336104fc610715565b6001600160a01b03161461052b5760405162461bcd60e51b815260040161052290611141565b60405180910390fd5b61053782826001610e2c565b5050565b33610544610715565b6001600160a01b03161461056a5760405162461bcd60e51b815260040161052290611141565b6006805467ffffffffffffffff19166001600160401b03831617905560045b6040516001600160401b0383168152600080516020611249833981519152906020015b60405180910390a250565b336105c0610715565b6001600160a01b0316146105e65760405162461bcd60e51b815260040161052290611141565b6006805463ffffffff60401b1916600160401b63ffffffff84160217905560055b60405163ffffffff83168152600080516020611249833981519152906020016105ac565b33610634610715565b6001600160a01b03161461065a5760405162461bcd60e51b815260040161052290611141565b6106646000610ecb565b565b3361066f610715565b6001600160a01b0316146106955760405162461bcd60e51b815260040161052290611141565b6005805463ffffffff60801b1916600160801b63ffffffff8416021790556001610607565b336106c3610715565b6001600160a01b0316146106e95760405162461bcd60e51b815260040161052290611141565b6005805467ffffffffffffffff60a81b1916600160a81b6001600160401b038416021790556003610589565b6000546001600160a01b031690565b3361072d610715565b6001600160a01b0316146107535760405162461bcd60e51b815260040161052290611141565b6005805460ff60a01b1916600160a01b60ff841602179055600260405160ff83168152600080516020611249833981519152906020016105ac565b33610797610715565b6001600160a01b0316146107bd5760405162461bcd60e51b815260040161052290611141565b61053782826001610f1b565b6107d1610715565b6001600160a01b0316336001600160a01b031614806107fa57506004546001600160a01b031633145b6108165760405162461bcd60e51b815260040161052290611176565b61053782826000610e2c565b6000805a9050478061087c5760045b60405186815233906001600160a01b038916907f2b8ae00e22d9eaf5a92820a22b947c007aee773fa36502ad7a1c9a464ab4932b9060200160405180910390a4600092505050610b2e565b3360009081526001602052604090205460ff1661089a576000610831565b6001600160a01b03861660009081526002602052604090205460ff166108c1576001610831565b336000908152600360205260409020544314156108df576002610831565b33600090815260036020526040812043905560055461090e90600160a81b90046001600160401b03164861119e565b9050803a101561091b57503a5b6006546001600160401b03161580159061093f57506006546001600160401b031681115b1561095257506006546001600160401b03165b6005546006546001600160a01b03891631916001600160801b03811691600160401b900463ffffffff169061099190600160a01b900460ff16896111b6565b6005546109ab90600160801b900463ffffffff168861119e565b6109b5919061119e565b6109bf908a61119e565b98505a6109cc908a6111d5565b985080158015906109dc57508089115b156109e5578098505b60006109f18a866111b6565b90508215801590610a0a575082610a08828661119e565b115b15610a765782841115610a695760036040518b815233906001600160a01b038e16907f2b8ae00e22d9eaf5a92820a22b947c007aee773fa36502ad7a1c9a464ab4932b9060200160405180910390a46000975050505050505050610b2e565b610a7384846111d5565b90505b85811115610a815750845b6040516001600160a01b038c16908290600081818185875af1925050503d8060008114610aca576040519150601f19603f3d011682016040523d82523d6000602084013e610acf565b606091505b5050604080518c8152602081018890529081018390529098508815159033906001600160a01b038e16907fd0224505f828ccfcbc56ca0590d97442e239a7aa770f712948fd6388356b20de9060600160405180910390a4505050505050505b9392505050565b610b3d610715565b6001600160a01b0316336001600160a01b03161480610b6657506004546001600160a01b031633145b610b825760405162461bcd60e51b815260040161052290611176565b61053782826000610f1b565b33610b97610715565b6001600160a01b031614610bbd5760405162461bcd60e51b815260040161052290611141565b600480546001600160a01b0319166001600160a01b0383169081179091556040517fc388cec0895ad7ee4635898ec92207ca48d42256d4355f7042efef62c368a97990600090a250565b33610c10610715565b6001600160a01b031614610c365760405162461bcd60e51b815260040161052290611141565b6001600160a01b038116610c9b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610522565b610ca481610ecb565b50565b33610cb0610715565b6001600160a01b031614610cd65760405162461bcd60e51b815260040161052290611141565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610d23576040519150601f19603f3d011682016040523d82523d6000602084013e610d28565b606091505b5050905080610d6b5760405162461bcd60e51b815260206004820152600f60248201526e15d2551211149055d7d19052531151608a1b6044820152606401610522565b604080513381526001600160a01b03851660208201529081018390527fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060600160405180910390a1505050565b33610dc2610715565b6001600160a01b031614610de85760405162461bcd60e51b815260040161052290611141565b600580546001600160801b0319166001600160801b03831617905560006040516001600160801b0383168152600080516020611249833981519152906020016105ac565b60005b82811015610ec5576000848483818110610e4b57610e4b61121d565b9050602002016020810190610e609190610fb4565b6001600160a01b038116600081815260016020526040808220805460ff19168815159081179091559051939450927fb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c019190a35080610ebd816111ec565b915050610e2f565b50505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005b82811015610ec5576000848483818110610f3a57610f3a61121d565b9050602002016020810190610f4f9190610fb4565b6001600160a01b038116600081815260026020526040808220805460ff19168815159081179091559051939450927ff544cca9d5484bfd447775bd759d12d53f1aa7c5f770be82c55070798ff9c63e9190a35080610fac816111ec565b915050610f1e565b600060208284031215610fc657600080fd5b8135610b2e81611233565b60008060408385031215610fe457600080fd5b8235610fef81611233565b946020939093013593505050565b60008060006060848603121561101257600080fd5b833561101d81611233565b95602085013595506040909401359392505050565b6000806020838503121561104557600080fd5b82356001600160401b038082111561105c57600080fd5b818501915085601f83011261107057600080fd5b81358181111561107f57600080fd5b8660208260051b850101111561109457600080fd5b60209290920196919550909350505050565b6000602082840312156110b857600080fd5b81356001600160801b0381168114610b2e57600080fd5b6000602082840312156110e157600080fd5b813563ffffffff81168114610b2e57600080fd5b60006020828403121561110757600080fd5b81356001600160401b0381168114610b2e57600080fd5b60006020828403121561113057600080fd5b813560ff81168114610b2e57600080fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252600e908201526d1393d517d055551213d49256915160921b604082015260600190565b600082198211156111b1576111b1611207565b500190565b60008160001904831182151516156111d0576111d0611207565b500290565b6000828210156111e7576111e7611207565b500390565b600060001982141561120057611200611207565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6001600160a01b0381168114610ca457600080fdfeda79b6b81f905f788560507c685a42d5a8ab209ee26538cbcf3ce3caed601f9ba264697066735822122057fd98a022d90ecf9d947d4257a574db310566830ec84b24efc7c83e4f6c2f1a64736f6c63430008070033"

// DeployGasRefunder deploys a new Ethereum contract, binding an instance of GasRefunder to it.
func DeployGasRefunder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasRefunder, error) {
	parsed, err := abi.JSON(strings.NewReader(GasRefunderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GasRefunderBin), backend)
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

// LastContractRefund is a free data retrieval call binding the contract method 0xbddaf01d.
//
// Solidity: function lastContractRefund(address ) view returns(uint256)
func (_GasRefunder *GasRefunderCaller) LastContractRefund(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "lastContractRefund", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastContractRefund is a free data retrieval call binding the contract method 0xbddaf01d.
//
// Solidity: function lastContractRefund(address ) view returns(uint256)
func (_GasRefunder *GasRefunderSession) LastContractRefund(arg0 common.Address) (*big.Int, error) {
	return _GasRefunder.Contract.LastContractRefund(&_GasRefunder.CallOpts, arg0)
}

// LastContractRefund is a free data retrieval call binding the contract method 0xbddaf01d.
//
// Solidity: function lastContractRefund(address ) view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) LastContractRefund(arg0 common.Address) (*big.Int, error) {
	return _GasRefunder.Contract.LastContractRefund(&_GasRefunder.CallOpts, arg0)
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
