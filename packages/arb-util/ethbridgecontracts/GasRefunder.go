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
const GasRefunderABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"ContractAllowedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DisallowerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"parameter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ParameterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reason\",\"type\":\"uint256\"}],\"name\":\"RefundGasCostsDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"}],\"name\":\"RefundedGasCosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"RefundeeAllowedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"allowContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"allowRefundees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedRefundees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calldataCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"disallowContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"disallowRefundees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disallower\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraGasMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastContractRefund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasTip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxRefundeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSingleGasUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"refundee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"onGasSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setCalldataCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDisallower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setExtraGasMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxGasCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxGasTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxRefundeeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxSingleGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// GasRefunderBin is the compiled bytecode used for deploying new contracts.
var GasRefunderBin = "0x608060405234801561001057600080fd5b5061001a33610045565b610fa0600655600c6007556377359400600855641bf08eb0006009556501d1a94a2000600a55610095565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611113806100a46000396000f3fe6080604052600436106101655760003560e01c8063b029df96116100c1578063cd499da31161007a578063cd499da314610421578063e3db8a4914610441578063e520745314610461578063efe12b0114610481578063f1e845ca146104a1578063f2fde38b146104c1578063f3fef3a3146104e157600080fd5b8063b029df9614610368578063bddaf01d1461037e578063c53d26a0146103ab578063c88bee17146103cb578063ca101295146103e1578063cd2c4e841461040157600080fd5b8063715018a61161011e578063715018a61461028a57806380599c3d1461029f5780638da5cb5b146102bf5780638e0bc1b6146102ec57806395554f7114610302578063a715dd5914610322578063a89d21731461033857600080fd5b80630119b029146101a957806325416bc9146101d25780633a16db73146101f457806351e0e26b14610214578063549e89bb14610254578063696e930c1461027457600080fd5b366101a457604080513381523460208201527f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4910160405180910390a1005b600080fd5b3480156101b557600080fd5b506101bf60085481565b6040519081526020015b60405180910390f35b3480156101de57600080fd5b506101f26101ed366004610f08565b610501565b005b34801561020057600080fd5b506101f261020f366004610f7d565b610549565b34801561022057600080fd5b5061024461022f366004610e8a565b60016020526000908152604090205460ff1681565b60405190151581526020016101c9565b34801561026057600080fd5b506101f261026f366004610f7d565b6105a5565b34801561028057600080fd5b506101bf600a5481565b34801561029657600080fd5b506101f26105fa565b3480156102ab57600080fd5b506101f26102ba366004610f7d565b610635565b3480156102cb57600080fd5b506102d461068a565b6040516001600160a01b0390911681526020016101c9565b3480156102f857600080fd5b506101bf60065481565b34801561030e57600080fd5b506101f261031d366004610f7d565b610699565b34801561032e57600080fd5b506101bf60095481565b34801561034457600080fd5b50610244610353366004610e8a565b60026020526000908152604090205460ff1681565b34801561037457600080fd5b506101bf60075481565b34801561038a57600080fd5b506101bf610399366004610e8a565b60036020526000908152604090205481565b3480156103b757600080fd5b506101f26103c6366004610f7d565b6106ee565b3480156103d757600080fd5b506101bf60055481565b3480156103ed57600080fd5b506101f26103fc366004610f08565b610743565b34801561040d57600080fd5b506101f261041c366004610f7d565b61077e565b34801561042d57600080fd5b506101f261043c366004610f08565b6107d3565b34801561044d57600080fd5b5061024461045c366004610ed3565b61082c565b34801561046d57600080fd5b506101f261047c366004610f08565b610add565b34801561048d57600080fd5b506004546102d4906001600160a01b031681565b3480156104ad57600080fd5b506101f26104bc366004610e8a565b610b36565b3480156104cd57600080fd5b506101f26104dc366004610e8a565b610baf565b3480156104ed57600080fd5b506101f26104fc366004610ea7565b610c4f565b3361050a61068a565b6001600160a01b0316146105395760405162461bcd60e51b815260040161053090610f96565b60405180910390fd5b61054582826001610d02565b5050565b3361055261068a565b6001600160a01b0316146105785760405162461bcd60e51b815260040161053090610f96565b60058190556040518181526000906000805160206110be833981519152906020015b60405180910390a250565b336105ae61068a565b6001600160a01b0316146105d45760405162461bcd60e51b815260040161053090610f96565b60078190556040518181526002906000805160206110be8339815191529060200161059a565b3361060361068a565b6001600160a01b0316146106295760405162461bcd60e51b815260040161053090610f96565b6106336000610da1565b565b3361063e61068a565b6001600160a01b0316146106645760405162461bcd60e51b815260040161053090610f96565b600a8190556040518181526005906000805160206110be8339815191529060200161059a565b6000546001600160a01b031690565b336106a261068a565b6001600160a01b0316146106c85760405162461bcd60e51b815260040161053090610f96565b60098190556040518181526004906000805160206110be8339815191529060200161059a565b336106f761068a565b6001600160a01b03161461071d5760405162461bcd60e51b815260040161053090610f96565b60088190556040518181526003906000805160206110be8339815191529060200161059a565b3361074c61068a565b6001600160a01b0316146107725760405162461bcd60e51b815260040161053090610f96565b61054582826001610df1565b3361078761068a565b6001600160a01b0316146107ad5760405162461bcd60e51b815260040161053090610f96565b60068190556040518181526001906000805160206110be8339815191529060200161059a565b6107db61068a565b6001600160a01b0316336001600160a01b0316148061080457506004546001600160a01b031633145b6108205760405162461bcd60e51b815260040161053090610fcb565b61054582826000610d02565b6000805a3360009081526001602052604090205490915060ff1661088957604080518581526000602082015233916001600160a01b0388169160008051602061109e83398151915291015b60405180910390a36000915050610ad6565b6001600160a01b03851660009081526002602052604090205460ff166108da57604080518581526001602082015233916001600160a01b0388169160008051602061109e8339815191529101610877565b3360009081526003602052604090205443141561092257604080518581526002602082015233916001600160a01b0388169160008051602061109e8339815191529101610877565b3360009081526003602052604081204390556008546109419048610ff3565b9050803a101561094e57503a5b60095415801590610960575060095481115b1561096a57506009545b600554600a546007546001600160a01b0389163192919061098b908861100b565b6006546109989087610ff3565b6109a29190610ff3565b6109ac9089610ff3565b97505a6109b9908961102a565b975080158015906109cb5750600a5488115b156109d4578097505b60006109e0898661100b565b905082158015906109f95750826109f78286610ff3565b115b15610a565782841115610a4957604080518a81526003602082015233916001600160a01b038d169160008051602061109e833981519152910160405180910390a360009650505050505050610ad6565b610a53848461102a565b90505b6040516001600160a01b038b169082156108fc029083906000818181858888f1604080518f8152602081018c9052908101879052909b508b151594503393506001600160a01b038f1692507fd0224505f828ccfcbc56ca0590d97442e239a7aa770f712948fd6388356b20de915060600160405180910390a45050505050505b9392505050565b610ae561068a565b6001600160a01b0316336001600160a01b03161480610b0e57506004546001600160a01b031633145b610b2a5760405162461bcd60e51b815260040161053090610fcb565b61054582826000610df1565b33610b3f61068a565b6001600160a01b031614610b655760405162461bcd60e51b815260040161053090610f96565b600480546001600160a01b0319166001600160a01b0383169081179091556040517fc388cec0895ad7ee4635898ec92207ca48d42256d4355f7042efef62c368a97990600090a250565b33610bb861068a565b6001600160a01b031614610bde5760405162461bcd60e51b815260040161053090610f96565b6001600160a01b038116610c435760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610530565b610c4c81610da1565b50565b33610c5861068a565b6001600160a01b031614610c7e5760405162461bcd60e51b815260040161053090610f96565b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610cb4573d6000803e3d6000fd5b50604080513381526001600160a01b03841660208201529081018290527fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060600160405180910390a15050565b60005b82811015610d9b576000848483818110610d2157610d21611072565b9050602002016020810190610d369190610e8a565b6001600160a01b038116600081815260016020526040808220805460ff19168815159081179091559051939450927fb0918cd965657b8d231f8adba328fa810b6d61d800de9c795d40eb3623498c019190a35080610d9381611041565b915050610d05565b50505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005b82811015610d9b576000848483818110610e1057610e10611072565b9050602002016020810190610e259190610e8a565b6001600160a01b038116600081815260026020526040808220805460ff19168815159081179091559051939450927ff544cca9d5484bfd447775bd759d12d53f1aa7c5f770be82c55070798ff9c63e9190a35080610e8281611041565b915050610df4565b600060208284031215610e9c57600080fd5b8135610ad681611088565b60008060408385031215610eba57600080fd5b8235610ec581611088565b946020939093013593505050565b600080600060608486031215610ee857600080fd5b8335610ef381611088565b95602085013595506040909401359392505050565b60008060208385031215610f1b57600080fd5b823567ffffffffffffffff80821115610f3357600080fd5b818501915085601f830112610f4757600080fd5b813581811115610f5657600080fd5b8660208260051b8501011115610f6b57600080fd5b60209290920196919550909350505050565b600060208284031215610f8f57600080fd5b5035919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252600e908201526d1393d517d055551213d49256915160921b604082015260600190565b600082198211156110065761100661105c565b500190565b60008160001904831182151516156110255761102561105c565b500290565b60008282101561103c5761103c61105c565b500390565b60006000198214156110555761105561105c565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6001600160a01b0381168114610c4c57600080fdfe5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc0fda05912fc0926602ebe745a349eb343a59dc2248f8ca238f896f385712170ca2646970667358221220cdf6ffcd8464e5a25743c94b42f889666c0f78f10405c86cf59a7431b9459a7664736f6c63430008070033"

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

// CalldataCost is a free data retrieval call binding the contract method 0xb029df96.
//
// Solidity: function calldataCost() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) CalldataCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "calldataCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalldataCost is a free data retrieval call binding the contract method 0xb029df96.
//
// Solidity: function calldataCost() view returns(uint256)
func (_GasRefunder *GasRefunderSession) CalldataCost() (*big.Int, error) {
	return _GasRefunder.Contract.CalldataCost(&_GasRefunder.CallOpts)
}

// CalldataCost is a free data retrieval call binding the contract method 0xb029df96.
//
// Solidity: function calldataCost() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) CalldataCost() (*big.Int, error) {
	return _GasRefunder.Contract.CalldataCost(&_GasRefunder.CallOpts)
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

// ExtraGasMargin is a free data retrieval call binding the contract method 0x8e0bc1b6.
//
// Solidity: function extraGasMargin() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) ExtraGasMargin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "extraGasMargin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraGasMargin is a free data retrieval call binding the contract method 0x8e0bc1b6.
//
// Solidity: function extraGasMargin() view returns(uint256)
func (_GasRefunder *GasRefunderSession) ExtraGasMargin() (*big.Int, error) {
	return _GasRefunder.Contract.ExtraGasMargin(&_GasRefunder.CallOpts)
}

// ExtraGasMargin is a free data retrieval call binding the contract method 0x8e0bc1b6.
//
// Solidity: function extraGasMargin() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) ExtraGasMargin() (*big.Int, error) {
	return _GasRefunder.Contract.ExtraGasMargin(&_GasRefunder.CallOpts)
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

// MaxGasCost is a free data retrieval call binding the contract method 0xa715dd59.
//
// Solidity: function maxGasCost() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) MaxGasCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "maxGasCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasCost is a free data retrieval call binding the contract method 0xa715dd59.
//
// Solidity: function maxGasCost() view returns(uint256)
func (_GasRefunder *GasRefunderSession) MaxGasCost() (*big.Int, error) {
	return _GasRefunder.Contract.MaxGasCost(&_GasRefunder.CallOpts)
}

// MaxGasCost is a free data retrieval call binding the contract method 0xa715dd59.
//
// Solidity: function maxGasCost() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) MaxGasCost() (*big.Int, error) {
	return _GasRefunder.Contract.MaxGasCost(&_GasRefunder.CallOpts)
}

// MaxGasTip is a free data retrieval call binding the contract method 0x0119b029.
//
// Solidity: function maxGasTip() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) MaxGasTip(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "maxGasTip")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasTip is a free data retrieval call binding the contract method 0x0119b029.
//
// Solidity: function maxGasTip() view returns(uint256)
func (_GasRefunder *GasRefunderSession) MaxGasTip() (*big.Int, error) {
	return _GasRefunder.Contract.MaxGasTip(&_GasRefunder.CallOpts)
}

// MaxGasTip is a free data retrieval call binding the contract method 0x0119b029.
//
// Solidity: function maxGasTip() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) MaxGasTip() (*big.Int, error) {
	return _GasRefunder.Contract.MaxGasTip(&_GasRefunder.CallOpts)
}

// MaxRefundeeBalance is a free data retrieval call binding the contract method 0xc88bee17.
//
// Solidity: function maxRefundeeBalance() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) MaxRefundeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "maxRefundeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRefundeeBalance is a free data retrieval call binding the contract method 0xc88bee17.
//
// Solidity: function maxRefundeeBalance() view returns(uint256)
func (_GasRefunder *GasRefunderSession) MaxRefundeeBalance() (*big.Int, error) {
	return _GasRefunder.Contract.MaxRefundeeBalance(&_GasRefunder.CallOpts)
}

// MaxRefundeeBalance is a free data retrieval call binding the contract method 0xc88bee17.
//
// Solidity: function maxRefundeeBalance() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) MaxRefundeeBalance() (*big.Int, error) {
	return _GasRefunder.Contract.MaxRefundeeBalance(&_GasRefunder.CallOpts)
}

// MaxSingleGasUsage is a free data retrieval call binding the contract method 0x696e930c.
//
// Solidity: function maxSingleGasUsage() view returns(uint256)
func (_GasRefunder *GasRefunderCaller) MaxSingleGasUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasRefunder.contract.Call(opts, &out, "maxSingleGasUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSingleGasUsage is a free data retrieval call binding the contract method 0x696e930c.
//
// Solidity: function maxSingleGasUsage() view returns(uint256)
func (_GasRefunder *GasRefunderSession) MaxSingleGasUsage() (*big.Int, error) {
	return _GasRefunder.Contract.MaxSingleGasUsage(&_GasRefunder.CallOpts)
}

// MaxSingleGasUsage is a free data retrieval call binding the contract method 0x696e930c.
//
// Solidity: function maxSingleGasUsage() view returns(uint256)
func (_GasRefunder *GasRefunderCallerSession) MaxSingleGasUsage() (*big.Int, error) {
	return _GasRefunder.Contract.MaxSingleGasUsage(&_GasRefunder.CallOpts)
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

// SetCalldataCost is a paid mutator transaction binding the contract method 0x549e89bb.
//
// Solidity: function setCalldataCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetCalldataCost(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setCalldataCost", newValue)
}

// SetCalldataCost is a paid mutator transaction binding the contract method 0x549e89bb.
//
// Solidity: function setCalldataCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetCalldataCost(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetCalldataCost(&_GasRefunder.TransactOpts, newValue)
}

// SetCalldataCost is a paid mutator transaction binding the contract method 0x549e89bb.
//
// Solidity: function setCalldataCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetCalldataCost(newValue *big.Int) (*types.Transaction, error) {
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

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0xcd2c4e84.
//
// Solidity: function setExtraGasMargin(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetExtraGasMargin(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setExtraGasMargin", newValue)
}

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0xcd2c4e84.
//
// Solidity: function setExtraGasMargin(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetExtraGasMargin(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetExtraGasMargin(&_GasRefunder.TransactOpts, newValue)
}

// SetExtraGasMargin is a paid mutator transaction binding the contract method 0xcd2c4e84.
//
// Solidity: function setExtraGasMargin(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetExtraGasMargin(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetExtraGasMargin(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x95554f71.
//
// Solidity: function setMaxGasCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxGasCost(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxGasCost", newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x95554f71.
//
// Solidity: function setMaxGasCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxGasCost(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasCost(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasCost is a paid mutator transaction binding the contract method 0x95554f71.
//
// Solidity: function setMaxGasCost(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxGasCost(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasCost(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0xc53d26a0.
//
// Solidity: function setMaxGasTip(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxGasTip(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxGasTip", newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0xc53d26a0.
//
// Solidity: function setMaxGasTip(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxGasTip(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasTip(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxGasTip is a paid mutator transaction binding the contract method 0xc53d26a0.
//
// Solidity: function setMaxGasTip(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxGasTip(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxGasTip(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0x3a16db73.
//
// Solidity: function setMaxRefundeeBalance(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxRefundeeBalance(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxRefundeeBalance", newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0x3a16db73.
//
// Solidity: function setMaxRefundeeBalance(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxRefundeeBalance(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxRefundeeBalance(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxRefundeeBalance is a paid mutator transaction binding the contract method 0x3a16db73.
//
// Solidity: function setMaxRefundeeBalance(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxRefundeeBalance(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxRefundeeBalance(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x80599c3d.
//
// Solidity: function setMaxSingleGasUsage(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactor) SetMaxSingleGasUsage(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.contract.Transact(opts, "setMaxSingleGasUsage", newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x80599c3d.
//
// Solidity: function setMaxSingleGasUsage(uint256 newValue) returns()
func (_GasRefunder *GasRefunderSession) SetMaxSingleGasUsage(newValue *big.Int) (*types.Transaction, error) {
	return _GasRefunder.Contract.SetMaxSingleGasUsage(&_GasRefunder.TransactOpts, newValue)
}

// SetMaxSingleGasUsage is a paid mutator transaction binding the contract method 0x80599c3d.
//
// Solidity: function setMaxSingleGasUsage(uint256 newValue) returns()
func (_GasRefunder *GasRefunderTransactorSession) SetMaxSingleGasUsage(newValue *big.Int) (*types.Transaction, error) {
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

// GasRefunderParameterSetIterator is returned from FilterParameterSet and is used to iterate over the raw logs and unpacked data for ParameterSet events raised by the GasRefunder contract.
type GasRefunderParameterSetIterator struct {
	Event *GasRefunderParameterSet // Event containing the contract specifics and raw log

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
func (it *GasRefunderParameterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasRefunderParameterSet)
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
		it.Event = new(GasRefunderParameterSet)
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
func (it *GasRefunderParameterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasRefunderParameterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasRefunderParameterSet represents a ParameterSet event raised by the GasRefunder contract.
type GasRefunderParameterSet struct {
	Parameter *big.Int
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterParameterSet is a free log retrieval operation binding the contract event 0x0fda05912fc0926602ebe745a349eb343a59dc2248f8ca238f896f385712170c.
//
// Solidity: event ParameterSet(uint256 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) FilterParameterSet(opts *bind.FilterOpts, parameter []*big.Int) (*GasRefunderParameterSetIterator, error) {

	var parameterRule []interface{}
	for _, parameterItem := range parameter {
		parameterRule = append(parameterRule, parameterItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "ParameterSet", parameterRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderParameterSetIterator{contract: _GasRefunder.contract, event: "ParameterSet", logs: logs, sub: sub}, nil
}

// WatchParameterSet is a free log subscription operation binding the contract event 0x0fda05912fc0926602ebe745a349eb343a59dc2248f8ca238f896f385712170c.
//
// Solidity: event ParameterSet(uint256 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) WatchParameterSet(opts *bind.WatchOpts, sink chan<- *GasRefunderParameterSet, parameter []*big.Int) (event.Subscription, error) {

	var parameterRule []interface{}
	for _, parameterItem := range parameter {
		parameterRule = append(parameterRule, parameterItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "ParameterSet", parameterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasRefunderParameterSet)
				if err := _GasRefunder.contract.UnpackLog(event, "ParameterSet", log); err != nil {
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

// ParseParameterSet is a log parse operation binding the contract event 0x0fda05912fc0926602ebe745a349eb343a59dc2248f8ca238f896f385712170c.
//
// Solidity: event ParameterSet(uint256 indexed parameter, uint256 value)
func (_GasRefunder *GasRefunderFilterer) ParseParameterSet(log types.Log) (*GasRefunderParameterSet, error) {
	event := new(GasRefunderParameterSet)
	if err := _GasRefunder.contract.UnpackLog(event, "ParameterSet", log); err != nil {
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
	Gas             *big.Int
	Reason          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefundGasCostsDenied is a free log retrieval operation binding the contract event 0x5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 reason)
func (_GasRefunder *GasRefunderFilterer) FilterRefundGasCostsDenied(opts *bind.FilterOpts, refundee []common.Address, contractAddress []common.Address) (*GasRefunderRefundGasCostsDeniedIterator, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _GasRefunder.contract.FilterLogs(opts, "RefundGasCostsDenied", refundeeRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &GasRefunderRefundGasCostsDeniedIterator{contract: _GasRefunder.contract, event: "RefundGasCostsDenied", logs: logs, sub: sub}, nil
}

// WatchRefundGasCostsDenied is a free log subscription operation binding the contract event 0x5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 reason)
func (_GasRefunder *GasRefunderFilterer) WatchRefundGasCostsDenied(opts *bind.WatchOpts, sink chan<- *GasRefunderRefundGasCostsDenied, refundee []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var refundeeRule []interface{}
	for _, refundeeItem := range refundee {
		refundeeRule = append(refundeeRule, refundeeItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _GasRefunder.contract.WatchLogs(opts, "RefundGasCostsDenied", refundeeRule, contractAddressRule)
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

// ParseRefundGasCostsDenied is a log parse operation binding the contract event 0x5ea5eb8c916c8ba2696eb02ba87979a3ba62781166df7c7e8770ca8e4d22f3fc.
//
// Solidity: event RefundGasCostsDenied(address indexed refundee, address indexed contractAddress, uint256 gas, uint256 reason)
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
