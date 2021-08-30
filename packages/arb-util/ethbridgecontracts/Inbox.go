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

// InboxABI is the input ABI used to generate the binding from.
const InboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"PauseToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"RewriteToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSource\",\"type\":\"address\"}],\"name\":\"WhitelistSourceUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicketNoRefundAliasRewrite\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCreateRetryablePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseCreateRetryables\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shouldRewriteSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startRewriteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopRewriteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseCreateRetryables\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSource\",\"type\":\"address\"}],\"name\":\"updateWhitelistSource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// InboxBin is the compiled bytecode used for deploying new contracts.
var InboxBin = "0x608060405234801561001057600080fd5b506000805460ff60a01b1916600160a01b179055611e82806100336000396000f3fe60806040526004361061011f5760003560e01c80636f791d29116100a05780639fe12da5116100645780639fe12da514610681578063b4d9ec4414610696578063b75436bb146106ab578063e78cea9214610726578063fdebb9b31461073b5761011f565b80636f791d2914610561578063794cfd511461058a5780637ae8d8b31461059f5780638a631aa6146105b457806393e59dc1146106505761011f565b8063485cc955116100e7578063485cc955146102c15780635075788b146102fc5780635e9167581461039f578063679b6ded1461042957806367ef3ab8146104d25761011f565b80630f4d14e9146101245780631b871c8d146101535780631fe927cf146101fc5780632b40609a1461027757806347466f981461028e575b600080fd5b6101416004803603602081101561013a57600080fd5b5035610750565b60408051918252519081900360200190f35b610141600480360361010081101561016a57600080fd5b6001600160a01b038235811692602081013592604082013592606083013581169260808101359091169160a08201359160c081013591810190610100810160e0820135600160201b8111156101be57600080fd5b8201836020820111156101d057600080fd5b803590602001918460018302840111600160201b831117156101f157600080fd5b50909250905061093d565b34801561020857600080fd5b506101416004803603602081101561021f57600080fd5b810190602081018135600160201b81111561023957600080fd5b82018360208201111561024b57600080fd5b803590602001918460018302840111600160201b8311171561026c57600080fd5b509092509050610b1d565b34801561028357600080fd5b5061028c610c8a565b005b34801561029a57600080fd5b5061028c600480360360208110156102b157600080fd5b50356001600160a01b0316610e52565b3480156102cd57600080fd5b5061028c600480360360408110156102e457600080fd5b506001600160a01b0381358116916020013516610ef5565b34801561030857600080fd5b50610141600480360360c081101561031f57600080fd5b8135916020810135916040820135916001600160a01b03606082013516916080820135919081019060c0810160a0820135600160201b81111561036157600080fd5b82018360208201111561037357600080fd5b803590602001918460018302840111600160201b8311171561039457600080fd5b509092509050610f70565b610141600480360360808110156103b557600080fd5b8135916020810135916001600160a01b036040830135169190810190608081016060820135600160201b8111156103eb57600080fd5b8201836020820111156103fd57600080fd5b803590602001918460018302840111600160201b8311171561041e57600080fd5b5090925090506110c5565b610141600480360361010081101561044057600080fd5b6001600160a01b038235811692602081013592604082013592606083013581169260808101359091169160a08201359160c081013591810190610100810160e0820135600160201b81111561049457600080fd5b8201836020820111156104a657600080fd5b803590602001918460018302840111600160201b831117156104c757600080fd5b509092509050611210565b610141600480360360a08110156104e857600080fd5b8135916020810135916040820135916001600160a01b036060820135169181019060a081016080820135600160201b81111561052357600080fd5b82018360208201111561053557600080fd5b803590602001918460018302840111600160201b8311171561055657600080fd5b509092509050611349565b34801561056d57600080fd5b5061057661149d565b604080519115158252519081900360200190f35b34801561059657600080fd5b5061028c6114ad565b3480156105ab57600080fd5b5061028c61166f565b3480156105c057600080fd5b50610141600480360360a08110156105d757600080fd5b8135916020810135916001600160a01b036040830135169160608101359181019060a081016080820135600160201b81111561061257600080fd5b82018360208201111561062457600080fd5b803590602001918460018302840111600160201b8311171561064557600080fd5b50909250905061183a565b34801561065c57600080fd5b5061066561197b565b604080516001600160a01b039092168252519081900360200190f35b34801561068d57600080fd5b5061028c61198a565b3480156106a257600080fd5b50610576611b49565b3480156106b757600080fd5b50610141600480360360208110156106ce57600080fd5b810190602081018135600160201b8111156106e857600080fd5b8201836020820111156106fa57600080fd5b803590602001918460018302840111600160201b8311171561071b57600080fd5b509092509050611b59565b34801561073257600080fd5b50610665611cb9565b34801561074757600080fd5b50610576611cc8565b600080546001600160a01b03161561081c576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b1580156107ad57600080fd5b505afa1580156107c1573d6000803e3d6000fd5b505050506040513d60208110156107d757600080fd5b505161081c576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b600154600160a01b900460ff1615610876576040805162461bcd60e51b815260206004820152601860248201527710d49150551157d4915514965050931154d7d4105554d15160421b604482015290519081900360640190fd5b60015433908190600160a81b900460ff16156108c25761089582611cd8565b1580156108a157503233145b156108b6576108af82611cde565b91506108c2565b6108bf81611ced565b90505b604080516001600160a01b0383166020820181905260008284018190523460608401526080830188905260a0830182905260c083019190915260e0820181905261010082018190526101208083019190915282518083039091018152610140909101909152610935906009908490611cfb565b949350505050565b600080546001600160a01b031615610a09576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b15801561099a57600080fd5b505afa1580156109ae573d6000803e3d6000fd5b505050506040513d60208110156109c457600080fd5b5051610a09576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b600154600160a01b900460ff1615610a63576040805162461bcd60e51b815260206004820152601860248201527710d49150551157d4915514965050931154d7d4105554d15160421b604482015290519081900360640190fd5b610b0f6009338c60601b60601c6001600160a01b03168c348d8d60601b60601c6001600160a01b03168d60601b60601c6001600160a01b03168d8d8d8d90508e8e604051602001808c81526020018b81526020018a8152602001898152602001888152602001878152602001868152602001858152602001848152602001838380828437808301925050509b505050505050505050505050604051602081830303815290604052611cfb565b9a9950505050505050505050565b600080546001600160a01b031615610be9576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b158015610b7a57600080fd5b505afa158015610b8e573d6000803e3d6000fd5b505050506040513d6020811015610ba457600080fd5b5051610be9576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b333214610c2b576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b6000610c5560033386866040518083838082843760405192018290039091209350611db492505050565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a29392505050565b60015460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b158015610ccf57600080fd5b505afa158015610ce3573d6000803e3d6000fd5b505050506040513d6020811015610cf957600080fd5b505160408051638da5cb5b60e01b815290519192506000916001600160a01b03841691638da5cb5b916004808301926020929190829003018186803b158015610d4157600080fd5b505afa158015610d55573d6000803e3d6000fd5b505050506040513d6020811015610d6b57600080fd5b50519050336001600160a01b03821614610db9576040805162461bcd60e51b815260206004820152600a60248201526904e4f545f524f4c4c55560b41b604482015290519081900360640190fd5b600154600160a01b900460ff1615610e09576040805162461bcd60e51b815260206004820152600e60248201526d1053149150511657d4105554d15160921b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b17815560408051918252517f9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b9181900360200190a15050565b6000546001600160a01b03163314610ea1576040805162461bcd60e51b815260206004820152600d60248201526c1393d517d19493d357d31254d5609a1b604482015290519081900360640190fd5b600080546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f37389c47920d5cc3229678a0205d0455002c07541a4139ebdce91ac2274657779181900360200190a150565b6001546001600160a01b031615610f42576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600180546001600160a01b039384166001600160a01b03199182161790915560008054929093169116179055565b600080546001600160a01b03161561103c576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b158015610fcd57600080fd5b505afa158015610fe1573d6000803e3d6000fd5b505050506040513d6020811015610ff757600080fd5b505161103c576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b6110b960033360008b8b8b8b60601b60601c6001600160a01b03168b8b8b604051602001808960ff1660ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052611cfb565b98975050505050505050565b600080546001600160a01b031615611191576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b15801561112257600080fd5b505afa158015611136573d6000803e3d6000fd5b505050506040513d602081101561114c57600080fd5b5051611191576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b611206600733600189898960601b60601c6001600160a01b0316348a8a604051602001808860ff1660ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052611cfb565b9695505050505050565b600080546001600160a01b0316156112dc576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b15801561126d57600080fd5b505afa158015611281573d6000803e3d6000fd5b505050506040513d602081101561129757600080fd5b50516112dc576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b600154600160a81b900460ff1680156112f957506112f987611cd8565b1561130a5761130787611ced565b96505b600154600160a81b900460ff168015611327575061132786611cd8565b156113385761133586611ced565b95505b610b0f8a8a8a8a8a8a8a8a8a61093d565b600080546001600160a01b031615611415576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b1580156113a657600080fd5b505afa1580156113ba573d6000803e3d6000fd5b505050506040513d60208110156113d057600080fd5b5051611415576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b61149260073360008a8a8a8a60601b60601c6001600160a01b0316348b8b604051602001808960ff1660ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052611cfb565b979650505050505050565b600054600160a01b900460ff1690565b60015460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b1580156114f257600080fd5b505afa158015611506573d6000803e3d6000fd5b505050506040513d602081101561151c57600080fd5b505160408051638da5cb5b60e01b815290519192506000916001600160a01b03841691638da5cb5b916004808301926020929190829003018186803b15801561156457600080fd5b505afa158015611578573d6000803e3d6000fd5b505050506040513d602081101561158e57600080fd5b50519050336001600160a01b038216146115dc576040805162461bcd60e51b815260206004820152600a60248201526904e4f545f524f4c4c55560b41b604482015290519081900360640190fd5b600154600160a81b900460ff1661162a576040805162461bcd60e51b815260206004820152600d60248201526c4e4f545f524557524954494e4760981b604482015290519081900360640190fd5b6001805460ff60a81b19169055604080516000815290517fab1ea65fd25ce96d303e895d1bd43edddb89841544a3705d3e61fc947a5fc25b9181900360200190a15050565b60015460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b1580156116b457600080fd5b505afa1580156116c8573d6000803e3d6000fd5b505050506040513d60208110156116de57600080fd5b505160408051638da5cb5b60e01b815290519192506000916001600160a01b03841691638da5cb5b916004808301926020929190829003018186803b15801561172657600080fd5b505afa15801561173a573d6000803e3d6000fd5b505050506040513d602081101561175057600080fd5b50519050336001600160a01b0382161461179e576040805162461bcd60e51b815260206004820152600a60248201526904e4f545f524f4c4c55560b41b604482015290519081900360640190fd5b600154600160a81b900460ff16156117f1576040805162461bcd60e51b8152602060048201526011602482015270414c52454144595f524557524954494e4760781b604482015290519081900360640190fd5b6001805460ff60a81b1916600160a81b17815560408051918252517fab1ea65fd25ce96d303e895d1bd43edddb89841544a3705d3e61fc947a5fc25b9181900360200190a15050565b600080546001600160a01b031615611906576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b15801561189757600080fd5b505afa1580156118ab573d6000803e3d6000fd5b505050506040513d60208110156118c157600080fd5b5051611906576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b61149260033360018a8a8a60601b60601c6001600160a01b03168a8a8a604051602001808860ff1660ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052611cfb565b6000546001600160a01b031681565b60015460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b1580156119cf57600080fd5b505afa1580156119e3573d6000803e3d6000fd5b505050506040513d60208110156119f957600080fd5b505160408051638da5cb5b60e01b815290519192506000916001600160a01b03841691638da5cb5b916004808301926020929190829003018186803b158015611a4157600080fd5b505afa158015611a55573d6000803e3d6000fd5b505050506040513d6020811015611a6b57600080fd5b50519050336001600160a01b03821614611ab9576040805162461bcd60e51b815260206004820152600a60248201526904e4f545f524f4c4c55560b41b604482015290519081900360640190fd5b600154600160a01b900460ff16611b04576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4105554d15160b21b604482015290519081900360640190fd5b6001805460ff60a01b19169055604080516000815290517f9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b9181900360200190a15050565b600154600160a01b900460ff1681565b600080546001600160a01b031615611c25576000546040805163babcc53960e01b815233600482015290516001600160a01b039092169163babcc53991602480820192602092909190829003018186803b158015611bb657600080fd5b505afa158015611bca573d6000803e3d6000fd5b505050506040513d6020811015611be057600080fd5b5051611c25576040805162461bcd60e51b815260206004820152600f60248201526e1393d517d5d2125511531254d51151608a1b604482015290519081900360640190fd5b6000611c4f60033386866040518083838082843760405192018290039091209350611db492505050565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b858560405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a29392505050565b6001546001600160a01b031681565b600154600160a81b900460ff1681565b3b151590565b61111061111160901b01190190565b61111161111160901b010190565b600080611d1085858580519060200120611db4565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b846040518080602001828103825283818151815260200191508051906020019080838360005b83811015611d72578181015183820152602001611d5a565b50505050905090810190601f168015611d9f5780820380516001836020036101000a031916815260200191505b509250505060405180910390a2949350505050565b600154604080516302bbfad160e01b815260ff861660048201526001600160a01b03858116602483015260448201859052915160009392909216916302bbfad1913491606480830192602092919082900301818588803b158015611e1757600080fd5b505af1158015611e2b573d6000803e3d6000fd5b50505050506040513d6020811015611e4257600080fd5b505194935050505056fea2646970667358221220dc6eb9d9af41f06db1644e58af1e76e013ba9018a7c7bc6b82973889c7f6720464736f6c634300060b0033"

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCallerSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// IsCreateRetryablePaused is a free data retrieval call binding the contract method 0xb4d9ec44.
//
// Solidity: function isCreateRetryablePaused() view returns(bool)
func (_Inbox *InboxCaller) IsCreateRetryablePaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "isCreateRetryablePaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCreateRetryablePaused is a free data retrieval call binding the contract method 0xb4d9ec44.
//
// Solidity: function isCreateRetryablePaused() view returns(bool)
func (_Inbox *InboxSession) IsCreateRetryablePaused() (bool, error) {
	return _Inbox.Contract.IsCreateRetryablePaused(&_Inbox.CallOpts)
}

// IsCreateRetryablePaused is a free data retrieval call binding the contract method 0xb4d9ec44.
//
// Solidity: function isCreateRetryablePaused() view returns(bool)
func (_Inbox *InboxCallerSession) IsCreateRetryablePaused() (bool, error) {
	return _Inbox.Contract.IsCreateRetryablePaused(&_Inbox.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Inbox *InboxCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Inbox *InboxSession) IsMaster() (bool, error) {
	return _Inbox.Contract.IsMaster(&_Inbox.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Inbox *InboxCallerSession) IsMaster() (bool, error) {
	return _Inbox.Contract.IsMaster(&_Inbox.CallOpts)
}

// ShouldRewriteSender is a free data retrieval call binding the contract method 0xfdebb9b3.
//
// Solidity: function shouldRewriteSender() view returns(bool)
func (_Inbox *InboxCaller) ShouldRewriteSender(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "shouldRewriteSender")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldRewriteSender is a free data retrieval call binding the contract method 0xfdebb9b3.
//
// Solidity: function shouldRewriteSender() view returns(bool)
func (_Inbox *InboxSession) ShouldRewriteSender() (bool, error) {
	return _Inbox.Contract.ShouldRewriteSender(&_Inbox.CallOpts)
}

// ShouldRewriteSender is a free data retrieval call binding the contract method 0xfdebb9b3.
//
// Solidity: function shouldRewriteSender() view returns(bool)
func (_Inbox *InboxCallerSession) ShouldRewriteSender() (bool, error) {
	return _Inbox.Contract.ShouldRewriteSender(&_Inbox.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_Inbox *InboxCaller) Whitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "whitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_Inbox *InboxSession) Whitelist() (common.Address, error) {
	return _Inbox.Contract.Whitelist(&_Inbox.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_Inbox *InboxCallerSession) Whitelist() (common.Address, error) {
	return _Inbox.Contract.Whitelist(&_Inbox.CallOpts)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) CreateRetryableTicket(opts *bind.TransactOpts, destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "createRetryableTicket", destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) CreateRetryableTicket(destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicket(&_Inbox.TransactOpts, destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) CreateRetryableTicket(destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicket(&_Inbox.TransactOpts, destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) CreateRetryableTicketNoRefundAliasRewrite(opts *bind.TransactOpts, destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "createRetryableTicketNoRefundAliasRewrite", destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) CreateRetryableTicketNoRefundAliasRewrite(destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicketNoRefundAliasRewrite(&_Inbox.TransactOpts, destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) CreateRetryableTicketNoRefundAliasRewrite(destAddr common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, maxGas *big.Int, gasPriceBid *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicketNoRefundAliasRewrite(&_Inbox.TransactOpts, destAddr, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, maxGas, gasPriceBid, data)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 maxSubmissionCost) payable returns(uint256)
func (_Inbox *InboxTransactor) DepositEth(opts *bind.TransactOpts, maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEth", maxSubmissionCost)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 maxSubmissionCost) payable returns(uint256)
func (_Inbox *InboxSession) DepositEth(maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, maxSubmissionCost)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 maxSubmissionCost) payable returns(uint256)
func (_Inbox *InboxTransactorSession) DepositEth(maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, maxSubmissionCost)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _whitelist) returns()
func (_Inbox *InboxTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address, _whitelist common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "initialize", _bridge, _whitelist)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _whitelist) returns()
func (_Inbox *InboxSession) Initialize(_bridge common.Address, _whitelist common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, _bridge, _whitelist)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _whitelist) returns()
func (_Inbox *InboxTransactorSession) Initialize(_bridge common.Address, _whitelist common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, _bridge, _whitelist)
}

// PauseCreateRetryables is a paid mutator transaction binding the contract method 0x2b40609a.
//
// Solidity: function pauseCreateRetryables() returns()
func (_Inbox *InboxTransactor) PauseCreateRetryables(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "pauseCreateRetryables")
}

// PauseCreateRetryables is a paid mutator transaction binding the contract method 0x2b40609a.
//
// Solidity: function pauseCreateRetryables() returns()
func (_Inbox *InboxSession) PauseCreateRetryables() (*types.Transaction, error) {
	return _Inbox.Contract.PauseCreateRetryables(&_Inbox.TransactOpts)
}

// PauseCreateRetryables is a paid mutator transaction binding the contract method 0x2b40609a.
//
// Solidity: function pauseCreateRetryables() returns()
func (_Inbox *InboxTransactorSession) PauseCreateRetryables() (*types.Transaction, error) {
	return _Inbox.Contract.PauseCreateRetryables(&_Inbox.TransactOpts)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendContractTransaction", maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedContractTransaction", maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// StartRewriteAddress is a paid mutator transaction binding the contract method 0x7ae8d8b3.
//
// Solidity: function startRewriteAddress() returns()
func (_Inbox *InboxTransactor) StartRewriteAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "startRewriteAddress")
}

// StartRewriteAddress is a paid mutator transaction binding the contract method 0x7ae8d8b3.
//
// Solidity: function startRewriteAddress() returns()
func (_Inbox *InboxSession) StartRewriteAddress() (*types.Transaction, error) {
	return _Inbox.Contract.StartRewriteAddress(&_Inbox.TransactOpts)
}

// StartRewriteAddress is a paid mutator transaction binding the contract method 0x7ae8d8b3.
//
// Solidity: function startRewriteAddress() returns()
func (_Inbox *InboxTransactorSession) StartRewriteAddress() (*types.Transaction, error) {
	return _Inbox.Contract.StartRewriteAddress(&_Inbox.TransactOpts)
}

// StopRewriteAddress is a paid mutator transaction binding the contract method 0x794cfd51.
//
// Solidity: function stopRewriteAddress() returns()
func (_Inbox *InboxTransactor) StopRewriteAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "stopRewriteAddress")
}

// StopRewriteAddress is a paid mutator transaction binding the contract method 0x794cfd51.
//
// Solidity: function stopRewriteAddress() returns()
func (_Inbox *InboxSession) StopRewriteAddress() (*types.Transaction, error) {
	return _Inbox.Contract.StopRewriteAddress(&_Inbox.TransactOpts)
}

// StopRewriteAddress is a paid mutator transaction binding the contract method 0x794cfd51.
//
// Solidity: function stopRewriteAddress() returns()
func (_Inbox *InboxTransactorSession) StopRewriteAddress() (*types.Transaction, error) {
	return _Inbox.Contract.StopRewriteAddress(&_Inbox.TransactOpts)
}

// UnpauseCreateRetryables is a paid mutator transaction binding the contract method 0x9fe12da5.
//
// Solidity: function unpauseCreateRetryables() returns()
func (_Inbox *InboxTransactor) UnpauseCreateRetryables(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "unpauseCreateRetryables")
}

// UnpauseCreateRetryables is a paid mutator transaction binding the contract method 0x9fe12da5.
//
// Solidity: function unpauseCreateRetryables() returns()
func (_Inbox *InboxSession) UnpauseCreateRetryables() (*types.Transaction, error) {
	return _Inbox.Contract.UnpauseCreateRetryables(&_Inbox.TransactOpts)
}

// UnpauseCreateRetryables is a paid mutator transaction binding the contract method 0x9fe12da5.
//
// Solidity: function unpauseCreateRetryables() returns()
func (_Inbox *InboxTransactorSession) UnpauseCreateRetryables() (*types.Transaction, error) {
	return _Inbox.Contract.UnpauseCreateRetryables(&_Inbox.TransactOpts)
}

// UpdateWhitelistSource is a paid mutator transaction binding the contract method 0x47466f98.
//
// Solidity: function updateWhitelistSource(address newSource) returns()
func (_Inbox *InboxTransactor) UpdateWhitelistSource(opts *bind.TransactOpts, newSource common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "updateWhitelistSource", newSource)
}

// UpdateWhitelistSource is a paid mutator transaction binding the contract method 0x47466f98.
//
// Solidity: function updateWhitelistSource(address newSource) returns()
func (_Inbox *InboxSession) UpdateWhitelistSource(newSource common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.UpdateWhitelistSource(&_Inbox.TransactOpts, newSource)
}

// UpdateWhitelistSource is a paid mutator transaction binding the contract method 0x47466f98.
//
// Solidity: function updateWhitelistSource(address newSource) returns()
func (_Inbox *InboxTransactorSession) UpdateWhitelistSource(newSource common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.UpdateWhitelistSource(&_Inbox.TransactOpts, newSource)
}

// InboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the Inbox contract.
type InboxInboxMessageDeliveredIterator struct {
	Event *InboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDelivered)
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
		it.Event = new(InboxInboxMessageDelivered)
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
func (it *InboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the Inbox contract.
type InboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredIterator{contract: _Inbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDelivered)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxInboxMessageDelivered, error) {
	event := new(InboxInboxMessageDelivered)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredFromOriginIterator{contract: _Inbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDeliveredFromOrigin)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxInboxMessageDeliveredFromOrigin)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxPauseToggledIterator is returned from FilterPauseToggled and is used to iterate over the raw logs and unpacked data for PauseToggled events raised by the Inbox contract.
type InboxPauseToggledIterator struct {
	Event *InboxPauseToggled // Event containing the contract specifics and raw log

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
func (it *InboxPauseToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxPauseToggled)
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
		it.Event = new(InboxPauseToggled)
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
func (it *InboxPauseToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxPauseToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxPauseToggled represents a PauseToggled event raised by the Inbox contract.
type InboxPauseToggled struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauseToggled is a free log retrieval operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool enabled)
func (_Inbox *InboxFilterer) FilterPauseToggled(opts *bind.FilterOpts) (*InboxPauseToggledIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return &InboxPauseToggledIterator{contract: _Inbox.contract, event: "PauseToggled", logs: logs, sub: sub}, nil
}

// WatchPauseToggled is a free log subscription operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool enabled)
func (_Inbox *InboxFilterer) WatchPauseToggled(opts *bind.WatchOpts, sink chan<- *InboxPauseToggled) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxPauseToggled)
				if err := _Inbox.contract.UnpackLog(event, "PauseToggled", log); err != nil {
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

// ParsePauseToggled is a log parse operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool enabled)
func (_Inbox *InboxFilterer) ParsePauseToggled(log types.Log) (*InboxPauseToggled, error) {
	event := new(InboxPauseToggled)
	if err := _Inbox.contract.UnpackLog(event, "PauseToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRewriteToggledIterator is returned from FilterRewriteToggled and is used to iterate over the raw logs and unpacked data for RewriteToggled events raised by the Inbox contract.
type InboxRewriteToggledIterator struct {
	Event *InboxRewriteToggled // Event containing the contract specifics and raw log

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
func (it *InboxRewriteToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRewriteToggled)
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
		it.Event = new(InboxRewriteToggled)
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
func (it *InboxRewriteToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRewriteToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRewriteToggled represents a RewriteToggled event raised by the Inbox contract.
type InboxRewriteToggled struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewriteToggled is a free log retrieval operation binding the contract event 0xab1ea65fd25ce96d303e895d1bd43edddb89841544a3705d3e61fc947a5fc25b.
//
// Solidity: event RewriteToggled(bool enabled)
func (_Inbox *InboxFilterer) FilterRewriteToggled(opts *bind.FilterOpts) (*InboxRewriteToggledIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RewriteToggled")
	if err != nil {
		return nil, err
	}
	return &InboxRewriteToggledIterator{contract: _Inbox.contract, event: "RewriteToggled", logs: logs, sub: sub}, nil
}

// WatchRewriteToggled is a free log subscription operation binding the contract event 0xab1ea65fd25ce96d303e895d1bd43edddb89841544a3705d3e61fc947a5fc25b.
//
// Solidity: event RewriteToggled(bool enabled)
func (_Inbox *InboxFilterer) WatchRewriteToggled(opts *bind.WatchOpts, sink chan<- *InboxRewriteToggled) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RewriteToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRewriteToggled)
				if err := _Inbox.contract.UnpackLog(event, "RewriteToggled", log); err != nil {
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

// ParseRewriteToggled is a log parse operation binding the contract event 0xab1ea65fd25ce96d303e895d1bd43edddb89841544a3705d3e61fc947a5fc25b.
//
// Solidity: event RewriteToggled(bool enabled)
func (_Inbox *InboxFilterer) ParseRewriteToggled(log types.Log) (*InboxRewriteToggled, error) {
	event := new(InboxRewriteToggled)
	if err := _Inbox.contract.UnpackLog(event, "RewriteToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxWhitelistSourceUpdatedIterator is returned from FilterWhitelistSourceUpdated and is used to iterate over the raw logs and unpacked data for WhitelistSourceUpdated events raised by the Inbox contract.
type InboxWhitelistSourceUpdatedIterator struct {
	Event *InboxWhitelistSourceUpdated // Event containing the contract specifics and raw log

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
func (it *InboxWhitelistSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxWhitelistSourceUpdated)
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
		it.Event = new(InboxWhitelistSourceUpdated)
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
func (it *InboxWhitelistSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxWhitelistSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxWhitelistSourceUpdated represents a WhitelistSourceUpdated event raised by the Inbox contract.
type InboxWhitelistSourceUpdated struct {
	NewSource common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWhitelistSourceUpdated is a free log retrieval operation binding the contract event 0x37389c47920d5cc3229678a0205d0455002c07541a4139ebdce91ac227465777.
//
// Solidity: event WhitelistSourceUpdated(address newSource)
func (_Inbox *InboxFilterer) FilterWhitelistSourceUpdated(opts *bind.FilterOpts) (*InboxWhitelistSourceUpdatedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "WhitelistSourceUpdated")
	if err != nil {
		return nil, err
	}
	return &InboxWhitelistSourceUpdatedIterator{contract: _Inbox.contract, event: "WhitelistSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistSourceUpdated is a free log subscription operation binding the contract event 0x37389c47920d5cc3229678a0205d0455002c07541a4139ebdce91ac227465777.
//
// Solidity: event WhitelistSourceUpdated(address newSource)
func (_Inbox *InboxFilterer) WatchWhitelistSourceUpdated(opts *bind.WatchOpts, sink chan<- *InboxWhitelistSourceUpdated) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "WhitelistSourceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxWhitelistSourceUpdated)
				if err := _Inbox.contract.UnpackLog(event, "WhitelistSourceUpdated", log); err != nil {
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

// ParseWhitelistSourceUpdated is a log parse operation binding the contract event 0x37389c47920d5cc3229678a0205d0455002c07541a4139ebdce91ac227465777.
//
// Solidity: event WhitelistSourceUpdated(address newSource)
func (_Inbox *InboxFilterer) ParseWhitelistSourceUpdated(log types.Log) (*InboxWhitelistSourceUpdated, error) {
	event := new(InboxWhitelistSourceUpdated)
	if err := _Inbox.contract.UnpackLog(event, "WhitelistSourceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
