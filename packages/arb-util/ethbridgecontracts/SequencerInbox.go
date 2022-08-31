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

// SequencerInboxMetaData contains all meta data concerning the SequencerInbox contract.
var SequencerInboxMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"afterAccAndDelayed\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"DelayedInboxForced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSequencer\",\"type\":\"bool\"}],\"name\":\"IsSequencerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxDelayBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"MaxDelayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"sectionsMetadata\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"shutdown\",\"type\":\"bool\"}],\"name\":\"ShutdownForNitroSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sectionsMetadata\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sectionsMetadata\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sectionsMetadata\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"}],\"name\":\"addSequencerL2BatchFromOriginWithGasRefunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedInbox\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"l1BlockAndTimestamp\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInboxAccsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_delayedInbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isNitroReady\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdownForNitro\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_messageCount\",\"type\":\"uint256\"}],\"name\":\"proveBatchContainsSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_messageCount\",\"type\":\"uint256\"}],\"name\":\"proveInboxContainsMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"newIsSequencer\",\"type\":\"bool\"}],\"name\":\"setIsSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMaxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"setMaxDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"}],\"name\":\"shutdownForNitro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelayedMessagesRead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undoShutdownForNitro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506000805460ff191660011790556123bd8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106101635760003560e01c80636f791d29116100ce578063b71939b111610087578063b71939b1146106a2578063c0c53b8b146106aa578063cb23bcb5146106e2578063d9b141ff146106ea578063d9dd67ab146106f2578063dc1b7b1f14610168578063e367a2c11461070f57610163565b80636f791d291461055d5780637e6c255f146105655780637fa3a40e1461056d5780638a2df18d1461057557806395fcea7814610692578063a8929e0b1461069a57610163565b8063381b003b11610120578063381b003b146103be5780633dbcc8d1146103da57806344c7cc30146103e25780634d480faa146104f05780635c1bba38146105135780636d46e9871461053757610163565b806306cc91b2146101685780630c4a1e59146101ef5780631a734229146102455780631f956632146103535780632aa8690914610381578063342025fa146103a4575b600080fd5b6101d66004803603604081101561017e57600080fd5b810190602081018135600160201b81111561019857600080fd5b8201836020820111156101aa57600080fd5b803590602001918460018302840111600160201b831117156101cb57600080fd5b919350915035610717565b6040805192835260208301919091528051918290030190f35b610243600480360361012081101561020657600080fd5b5080359060ff60208201351690604081019060808101359060a0810135906001600160a01b0360c0820135169060e0810135906101000135610732565b005b6102436004803603608081101561025b57600080fd5b810190602081018135600160201b81111561027557600080fd5b82018360208201111561028757600080fd5b803590602001918460018302840111600160201b831117156102a857600080fd5b919390929091602081019035600160201b8111156102c557600080fd5b8201836020820111156102d757600080fd5b803590602001918460208302840111600160201b831117156102f857600080fd5b919390929091602081019035600160201b81111561031557600080fd5b82018360208201111561032757600080fd5b803590602001918460208302840111600160201b8311171561034857600080fd5b919350915035610a10565b6102436004803603604081101561036957600080fd5b506001600160a01b0381351690602001351515610be6565b6102436004803603604081101561039757600080fd5b5080359060200135610c97565b6103ac610e2c565b60408051918252519081900360200190f35b6103c6610e32565b604080519115158252519081900360200190f35b6103ac610e3b565b610243600480360360808110156103f857600080fd5b810190602081018135600160201b81111561041257600080fd5b82018360208201111561042457600080fd5b803590602001918460018302840111600160201b8311171561044557600080fd5b919390929091602081019035600160201b81111561046257600080fd5b82018360208201111561047457600080fd5b803590602001918460208302840111600160201b8311171561049557600080fd5b919390929091602081019035600160201b8111156104b257600080fd5b8201836020820111156104c457600080fd5b803590602001918460208302840111600160201b831117156104e557600080fd5b919350915035610e41565b6102436004803603604081101561050657600080fd5b5080359060200135610fab565b61051b611041565b604080516001600160a01b039092168252519081900360200190f35b6103c66004803603602081101561054d57600080fd5b50356001600160a01b0316611050565b6103c6611065565b61024361106e565b6103ac6111b3565b610243600480360360a081101561058b57600080fd5b810190602081018135600160201b8111156105a557600080fd5b8201836020820111156105b757600080fd5b803590602001918460018302840111600160201b831117156105d857600080fd5b919390929091602081019035600160201b8111156105f557600080fd5b82018360208201111561060757600080fd5b803590602001918460208302840111600160201b8311171561062857600080fd5b919390929091602081019035600160201b81111561064557600080fd5b82018360208201111561065757600080fd5b803590602001918460208302840111600160201b8311171561067857600080fd5b9193509150803590602001356001600160a01b03166111b9565b6102436113d0565b6103ac61142d565b61051b611433565b610243600480360360608110156106c057600080fd5b506001600160a01b038135811691602081013582169160409091013516611442565b61051b6114d9565b6103ac6114e8565b6103ac6004803603602081101561070857600080fd5b50356114ee565b6103ac61150c565b600080610725858585611512565b915091505b935093915050565b600a5460408051808201909152601281527153485554444f574e5f464f525f4e4954524f60701b60208201529060ff16156107eb5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156107b0578181015183820152602001610798565b50505050905090810190601f1680156107dd5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060006108028885893560208b01358a8a8961165f565b60085490915043883590910110610853576040805162461bcd60e51b815260206004820152601060248201526f4d41585f44454c41595f424c4f434b5360801b604482015290519081900360640190fd5b600954426020890135909101106108a2576040805162461bcd60e51b815260206004820152600e60248201526d4d41585f44454c41595f54494d4560901b604482015290519081900360640190fd5b600060018a111561092c57600480546040805163d9dd67ab60e01b81526001198e0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b1580156108fd57600080fd5b505afa158015610911573d6000803e3d6000fd5b505050506040513d602081101561092757600080fd5b505190505b61093681836116d5565b600480546040805163d9dd67ab60e01b81526000198f0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561098657600080fd5b505afa15801561099a573d6000803e3d6000fd5b505050506040513d60208110156109b057600080fd5b5051146109fa576040805162461bcd60e51b81526020600482015260136024820152722222a620aca2a22fa0a1a1aaa6aaa620aa27a960691b604482015290519081900360640190fd5b5050610a068882611701565b5050505050505050565b600a5460408051808201909152601281527153485554444f574e5f464f525f4e4954524f60701b60208201529060ff1615610a8c5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b0578181015183820152602001610798565b50600060025490506000610adc89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91508990508888611865565b905080827f3bf85aebd2a1dc6c510ffc4795a3785e786b5817ab30144f88501d4c6456c986600254868d8d8d8d8d8d600180805490500333604051808b81526020018a8152602001806020018060200180602001868152602001856001600160a01b03166001600160a01b0316815260200184810384528c8c82818152602001925080828437600083820152601f01601f191690910185810384528a8152602090810191508b908b0280828437600083820152601f01601f19169091018581038352888152602090810191508990890280828437600083820152604051601f909101601f19169092018290039f50909d5050505050505050505050505050a3505050505050505050565b6006546001600160a01b03163314610c33576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b6001600160a01b038216600081815260076020908152604091829020805460ff191685151590811790915582519384529083015280517fce86e570206e55533301cb66529b33afbd75e991c575b85adeaca10146be8cb49281900390910190a15050565b600a5460408051808201909152601281527153485554444f574e5f464f525f4e4954524f60701b60208201529060ff1615610d135760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b0578181015183820152602001610798565b5060065460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b158015610d5857600080fd5b505afa158015610d6c573d6000803e3d6000fd5b505050506040513d6020811015610d8257600080fd5b50516001600160a01b031614610dd3576040805162461bcd60e51b815260206004820152601160248201527027a7262cafa927a6262aa82fa7aba722a960791b604482015290519081900360640190fd5b6003548214610de657610de68282611701565b600a805460ff1916600190811790915560408051918252517fe6d1c315c736941d015418a8728891eae6aea39817b9c134486054bf34cc336b9181900360200190a15050565b60095481565b600a5460ff1681565b60025481565b600a5460408051808201909152601281527153485554444f574e5f464f525f4e4954524f60701b60208201529060ff1615610ebd5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b0578181015183820152602001610798565b50333214610f00576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600060025490506000610f4f89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91508990508888611865565b60025460015460408051928352602083018790526000199091018282015251919250829184917f10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682919081900360600190a3505050505050505050565b6006546001600160a01b03163314610ff8576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60088290556009819055604080518381526020810183905281517f3bcd3c6d4304309e4b36d94f90517baf304582bb1ac828906808577e067e6b6e929181900390910190a15050565b6005546001600160a01b031690565b60076020526000908152604090205460ff1681565b60005460ff1690565b60065460408051638da5cb5b60e01b8152905133926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b1580156110b257600080fd5b505afa1580156110c6573d6000803e3d6000fd5b505050506040513d60208110156110dc57600080fd5b50516001600160a01b03161461112d576040805162461bcd60e51b815260206004820152601160248201527027a7262cafa927a6262aa82fa7aba722a960791b604482015290519081900360640190fd5b600a5460ff16611173576040805162461bcd60e51b815260206004820152600c60248201526b2727aa2fa9a42aaa2227aba760a11b604482015290519081900360640190fd5b600a805460ff19169055604080516000815290517fe6d1c315c736941d015418a8728891eae6aea39817b9c134486054bf34cc336b9181900360200190a1565b60035481565b600a5460408051808201909152601281527153485554444f574e5f464f525f4e4954524f60701b60208201529060ff16156112355760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b0578181015183820152602001610798565b50333214611278576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b60005a600254604080516020601f8d018190048102820181019092528b815292935036926000916112cb91908e908e90819084018382808284376000920191909152508e92508d91508c90508b8b611865565b60025460015460408051928352602083018a90526000199091018282015251919250829184917f10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682919081900360600190a36001600160a01b038516156113c257846001600160a01b031663e3db8a49335a8703866040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b031681526020018381526020018281526020019350505050602060405180830381600087803b15801561139557600080fd5b505af11580156113a9573d6000803e3d6000fd5b505050506040513d60208110156113bf57600080fd5b50505b505050505050505050505050565b60006113da611de3565b9050336001600160a01b0382161461142a576040805162461bcd60e51b815260206004820152600e60248201526d2727aa2fa32927a6afa0a226a4a760911b604482015290519081900360640190fd5b50565b61a4b790565b6004546001600160a01b031681565b6004546001600160a01b03161561148f576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600480546001600160a01b039485166001600160a01b0319918216179091559183166000908152600760205260409020805460ff1916600117905560068054919093169116179055565b6006546001600160a01b031681565b60015490565b600181815481106114fb57fe5b600091825260209091200154905081565b60085481565b600080826115255750600090508061072a565b60008061156787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509250611e08915050565b90925090506000811561159f5761159988888560018087038154811061158957fe5b9060005260206000200154611e7c565b90935090505b6000600183815481106115ae57fe5b9060005260206000200154905060006115c98a8a8785611e7c565b9095509050828811611610576040805162461bcd60e51b815260206004820152600b60248201526a10905510d217d4d510549560aa1b604482015290519081900360640190fd5b80881115611651576040805162461bcd60e51b815260206004820152600960248201526810905510d217d1539160ba1b604482015290519081900360640190fd5b999098509650505050505050565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6bffffffffffffffffffffffff19166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600354821161174b576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b600254600154600090156117795760018054600019810190811061176b57fe5b906000526020600020015490505b60008061178a83858843428a61203c565b60018054808201825560008281527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6909101849055600283905560408051808201825285815260208181018c9052935482518681529485018d9052959750939550879489947f85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a09488948e9492936000190192828101918591908190849084905b8381101561184257818101518382015260200161182a565b5050505090500182815260200194505050505060405180910390a3505050505050565b3360009081526007602052604081205460ff166118ba576040805162461bcd60e51b815260206004820152600e60248201526d27a7262cafa9a2a8aaa2a721a2a960911b604482015290519081900360640190fd5b600154156118e2576001805460001981019081106118d457fe5b906000526020600020015490505b60025481600060208a01815b600581018810611c7c57600089898360010181811061190957fe5b905060200201359050436008548201101561195b576040805162461bcd60e51b815260206004820152600d60248201526c109313d0d2d7d513d3d7d3d311609a1b604482015290519081900360640190fd5b438111156119a0576040805162461bcd60e51b815260206004820152600d60248201526c424c4f434b5f544f4f5f4e455760981b604482015290519081900360640190fd5b5060008989836002018181106119b257fe5b9050602002013590504260095482011015611a03576040805162461bcd60e51b815260206004820152600c60248201526b1512535157d513d3d7d3d31160a21b604482015290519081900360640190fd5b42811115611a47576040805162461bcd60e51b815260206004820152600c60248201526b54494d455f544f4f5f4e455760a01b604482015290519081900360640190fd5b506000338a8a84600101818110611a5a57fe5b905060200201358b8b85600201818110611a7057fe5b9050602002013560405160200180846001600160a01b03166001600160a01b031660601b8152601401838152602001828152602001935050505060405160208183030381529060405280519060200120905060008a8a84818110611ad057fe5b905060200201359050611ae9848e8e8885878d8d6122a0565b909850909650940193925060009050898960038401818110611b0757fe5b905060200201359050600354811015611b5b576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b6001811015611ba5576040805162461bcd60e51b8152602060048201526011602482015270135554d517d111531056515117d2539255607a1b604482015290519081900360640190fd5b6001600354101580611bc95750898983818110611bbe57fe5b905060200201356000145b611c14576040805162461bcd60e51b8152602060048201526017602482015276135554d517d111531056515117d253925517d4d5105495604a1b604482015290519081900360640190fd5b600354811115611c7357611c6e8587838d8d87600101818110611c3357fe5b905060200201358e8e88600201818110611c4957fe5b905060200201358f8f89600401818110611c5f57fe5b9050602002013560001b61203c565b965094505b506005016118ee565b5060208b0180821015611cc8576040805162461bcd60e51b815260206004820152600f60248201526e4f46465345545f4f564552464c4f5760881b604482015290519081900360640190fd5b8b518101821115611d17576040805162461bcd60e51b81526020600482015260146024820152732a2920a729a0a1aa24a7a729afa7ab22a9292aa760611b604482015290519081900360640190fd5b6002548511611d5b576040805162461bcd60e51b815260206004820152600b60248201526a08a9aa0a8b2be8482a886960ab1b604482015290519081900360640190fd5b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018490556002859055868414611dd4576040805162461bcd60e51b815260206004820152600960248201526841465445525f41434360b81b604482015290519081900360640190fd5b50505050509695505050505050565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b60008082845110158015611e20575060208385510310155b611e5d576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301611e71858563ffffffff61232e16565b915091509250929050565b6000806000806000806000611ec88b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250611e08915050565b809550819a505050611f118b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250611e08915050565b809450819a505050611f5a8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250611e08915050565b809350819a505050611fa38b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250611e08915050565b604080516020808201989098528082018790526060810186905260808082018490528251808303909101815260a0909101909152805196019590952090995060018401955093905087841461202b576040805162461bcd60e51b815260206004820152600960248201526842415443485f41434360b81b604482015290519081900360640190fd5b509699929850919650505050505050565b6004805460408051633dbcc8d160e01b8152905160009384936001600160a01b031692633dbcc8d19281830192602092829003018186803b15801561208057600080fd5b505afa158015612094573d6000803e3d6000fd5b505050506040513d60208110156120aa57600080fd5b50518611156120f2576040805162461bcd60e51b815260206004820152600f60248201526e2222a620aca2a22faa27a7afa320a960891b604482015290519081900360640190fd5b600480546040805163d9dd67ab60e01b81526000198a0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561214257600080fd5b505afa158015612156573d6000803e3d6000fd5b505050506040513d602081101561216c57600080fd5b505183146121af576040805162461bcd60e51b815260206004820152600b60248201526a44454c415945445f41434360a81b604482015290519081900360640190fd5b50506003805460408051702232b630bcb2b21036b2b9b9b0b3b2b99d60791b602080830191909152603182019a909a5260518101899052607181018390526091810188905260b1808201959095528151808203909501855260d1810182528451948a0194909420600060f186015261010585019690965261012580850195909552805180850390950185526101458401815284519489019490942060605160802061016585019690965290860390960161018583018190526101a58301969096526101c580830194909452825180830390940184526101e59091019091528151919094012092559091600190910190565b92840192808289875b8781101561231f5760008b8b838181106122bf57fe5b60209081029290920135808620604080518086019a909a5289810189905260608a018d90526080808b01929092528051808b03909201825260a09099019098528751979092019690962095506001948501949301929190910190506122a9565b50985098509895505050505050565b6000816020018351101561237e576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b5001602001519056fea2646970667358221220b5ba73c644de2ce471db488ffce4cfbd774d5ee918b7d60282d056ff815c684e64736f6c634300060b0033",
}

// SequencerInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerInboxMetaData.ABI instead.
var SequencerInboxABI = SequencerInboxMetaData.ABI

// SequencerInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerInboxMetaData.Bin instead.
var SequencerInboxBin = SequencerInboxMetaData.Bin

// DeploySequencerInbox deploys a new Ethereum contract, binding an instance of SequencerInbox to it.
func DeploySequencerInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SequencerInbox, error) {
	parsed, err := SequencerInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SequencerInbox{SequencerInboxCaller: SequencerInboxCaller{contract: contract}, SequencerInboxTransactor: SequencerInboxTransactor{contract: contract}, SequencerInboxFilterer: SequencerInboxFilterer{contract: contract}}, nil
}

// SequencerInbox is an auto generated Go binding around an Ethereum contract.
type SequencerInbox struct {
	SequencerInboxCaller     // Read-only binding to the contract
	SequencerInboxTransactor // Write-only binding to the contract
	SequencerInboxFilterer   // Log filterer for contract events
}

// SequencerInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerInboxSession struct {
	Contract     *SequencerInbox   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SequencerInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerInboxCallerSession struct {
	Contract *SequencerInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SequencerInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerInboxTransactorSession struct {
	Contract     *SequencerInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SequencerInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerInboxRaw struct {
	Contract *SequencerInbox // Generic contract binding to access the raw methods on
}

// SequencerInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerInboxCallerRaw struct {
	Contract *SequencerInboxCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerInboxTransactorRaw struct {
	Contract *SequencerInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerInbox creates a new instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInbox(address common.Address, backend bind.ContractBackend) (*SequencerInbox, error) {
	contract, err := bindSequencerInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerInbox{SequencerInboxCaller: SequencerInboxCaller{contract: contract}, SequencerInboxTransactor: SequencerInboxTransactor{contract: contract}, SequencerInboxFilterer: SequencerInboxFilterer{contract: contract}}, nil
}

// NewSequencerInboxCaller creates a new read-only instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxCaller(address common.Address, caller bind.ContractCaller) (*SequencerInboxCaller, error) {
	contract, err := bindSequencerInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxCaller{contract: contract}, nil
}

// NewSequencerInboxTransactor creates a new write-only instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerInboxTransactor, error) {
	contract, err := bindSequencerInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxTransactor{contract: contract}, nil
}

// NewSequencerInboxFilterer creates a new log filterer instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerInboxFilterer, error) {
	contract, err := bindSequencerInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxFilterer{contract: contract}, nil
}

// bindSequencerInbox binds a generic wrapper to an already deployed contract.
func bindSequencerInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInbox *SequencerInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInbox.Contract.SequencerInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInbox *SequencerInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SequencerInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInbox *SequencerInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SequencerInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInbox *SequencerInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInbox *SequencerInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInbox *SequencerInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInbox.Contract.contract.Transact(opts, method, params...)
}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_SequencerInbox *SequencerInboxCaller) DelayedInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "delayedInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_SequencerInbox *SequencerInboxSession) DelayedInbox() (common.Address, error) {
	return _SequencerInbox.Contract.DelayedInbox(&_SequencerInbox.CallOpts)
}

// DelayedInbox is a free data retrieval call binding the contract method 0xb71939b1.
//
// Solidity: function delayedInbox() view returns(address)
func (_SequencerInbox *SequencerInboxCallerSession) DelayedInbox() (common.Address, error) {
	return _SequencerInbox.Contract.DelayedInbox(&_SequencerInbox.CallOpts)
}

// GetInboxAccsLength is a free data retrieval call binding the contract method 0xd9b141ff.
//
// Solidity: function getInboxAccsLength() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) GetInboxAccsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "getInboxAccsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInboxAccsLength is a free data retrieval call binding the contract method 0xd9b141ff.
//
// Solidity: function getInboxAccsLength() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) GetInboxAccsLength() (*big.Int, error) {
	return _SequencerInbox.Contract.GetInboxAccsLength(&_SequencerInbox.CallOpts)
}

// GetInboxAccsLength is a free data retrieval call binding the contract method 0xd9b141ff.
//
// Solidity: function getInboxAccsLength() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) GetInboxAccsLength() (*big.Int, error) {
	return _SequencerInbox.Contract.GetInboxAccsLength(&_SequencerInbox.CallOpts)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_SequencerInbox *SequencerInboxCaller) InboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "inboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_SequencerInbox *SequencerInboxSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _SequencerInbox.Contract.InboxAccs(&_SequencerInbox.CallOpts, arg0)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _SequencerInbox.Contract.InboxAccs(&_SequencerInbox.CallOpts, arg0)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_SequencerInbox *SequencerInboxCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_SequencerInbox *SequencerInboxSession) IsMaster() (bool, error) {
	return _SequencerInbox.Contract.IsMaster(&_SequencerInbox.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_SequencerInbox *SequencerInboxCallerSession) IsMaster() (bool, error) {
	return _SequencerInbox.Contract.IsMaster(&_SequencerInbox.CallOpts)
}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) IsNitroReady(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "isNitroReady")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_SequencerInbox *SequencerInboxSession) IsNitroReady() (*big.Int, error) {
	return _SequencerInbox.Contract.IsNitroReady(&_SequencerInbox.CallOpts)
}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) IsNitroReady() (*big.Int, error) {
	return _SequencerInbox.Contract.IsNitroReady(&_SequencerInbox.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInbox *SequencerInboxCaller) IsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "isSequencer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInbox *SequencerInboxSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _SequencerInbox.Contract.IsSequencer(&_SequencerInbox.CallOpts, arg0)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInbox *SequencerInboxCallerSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _SequencerInbox.Contract.IsSequencer(&_SequencerInbox.CallOpts, arg0)
}

// IsShutdownForNitro is a free data retrieval call binding the contract method 0x381b003b.
//
// Solidity: function isShutdownForNitro() view returns(bool)
func (_SequencerInbox *SequencerInboxCaller) IsShutdownForNitro(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "isShutdownForNitro")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdownForNitro is a free data retrieval call binding the contract method 0x381b003b.
//
// Solidity: function isShutdownForNitro() view returns(bool)
func (_SequencerInbox *SequencerInboxSession) IsShutdownForNitro() (bool, error) {
	return _SequencerInbox.Contract.IsShutdownForNitro(&_SequencerInbox.CallOpts)
}

// IsShutdownForNitro is a free data retrieval call binding the contract method 0x381b003b.
//
// Solidity: function isShutdownForNitro() view returns(bool)
func (_SequencerInbox *SequencerInboxCallerSession) IsShutdownForNitro() (bool, error) {
	return _SequencerInbox.Contract.IsShutdownForNitro(&_SequencerInbox.CallOpts)
}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) MaxDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "maxDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) MaxDelayBlocks() (*big.Int, error) {
	return _SequencerInbox.Contract.MaxDelayBlocks(&_SequencerInbox.CallOpts)
}

// MaxDelayBlocks is a free data retrieval call binding the contract method 0xe367a2c1.
//
// Solidity: function maxDelayBlocks() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) MaxDelayBlocks() (*big.Int, error) {
	return _SequencerInbox.Contract.MaxDelayBlocks(&_SequencerInbox.CallOpts)
}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) MaxDelaySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "maxDelaySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) MaxDelaySeconds() (*big.Int, error) {
	return _SequencerInbox.Contract.MaxDelaySeconds(&_SequencerInbox.CallOpts)
}

// MaxDelaySeconds is a free data retrieval call binding the contract method 0x342025fa.
//
// Solidity: function maxDelaySeconds() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) MaxDelaySeconds() (*big.Int, error) {
	return _SequencerInbox.Contract.MaxDelaySeconds(&_SequencerInbox.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) MessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "messageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) MessageCount() (*big.Int, error) {
	return _SequencerInbox.Contract.MessageCount(&_SequencerInbox.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) MessageCount() (*big.Int, error) {
	return _SequencerInbox.Contract.MessageCount(&_SequencerInbox.CallOpts)
}

// PostUpgradeInit is a free data retrieval call binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() view returns()
func (_SequencerInbox *SequencerInboxCaller) PostUpgradeInit(opts *bind.CallOpts) error {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "postUpgradeInit")

	if err != nil {
		return err
	}

	return err

}

// PostUpgradeInit is a free data retrieval call binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() view returns()
func (_SequencerInbox *SequencerInboxSession) PostUpgradeInit() error {
	return _SequencerInbox.Contract.PostUpgradeInit(&_SequencerInbox.CallOpts)
}

// PostUpgradeInit is a free data retrieval call binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() view returns()
func (_SequencerInbox *SequencerInboxCallerSession) PostUpgradeInit() error {
	return _SequencerInbox.Contract.PostUpgradeInit(&_SequencerInbox.CallOpts)
}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxCaller) ProveBatchContainsSequenceNumber(opts *bind.CallOpts, proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "proveBatchContainsSequenceNumber", proof, _messageCount)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxSession) ProveBatchContainsSequenceNumber(proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	return _SequencerInbox.Contract.ProveBatchContainsSequenceNumber(&_SequencerInbox.CallOpts, proof, _messageCount)
}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) ProveBatchContainsSequenceNumber(proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	return _SequencerInbox.Contract.ProveBatchContainsSequenceNumber(&_SequencerInbox.CallOpts, proof, _messageCount)
}

// ProveInboxContainsMessage is a free data retrieval call binding the contract method 0xdc1b7b1f.
//
// Solidity: function proveInboxContainsMessage(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxCaller) ProveInboxContainsMessage(opts *bind.CallOpts, proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "proveInboxContainsMessage", proof, _messageCount)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// ProveInboxContainsMessage is a free data retrieval call binding the contract method 0xdc1b7b1f.
//
// Solidity: function proveInboxContainsMessage(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxSession) ProveInboxContainsMessage(proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	return _SequencerInbox.Contract.ProveInboxContainsMessage(&_SequencerInbox.CallOpts, proof, _messageCount)
}

// ProveInboxContainsMessage is a free data retrieval call binding the contract method 0xdc1b7b1f.
//
// Solidity: function proveInboxContainsMessage(bytes proof, uint256 _messageCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) ProveInboxContainsMessage(proof []byte, _messageCount *big.Int) (*big.Int, [32]byte, error) {
	return _SequencerInbox.Contract.ProveInboxContainsMessage(&_SequencerInbox.CallOpts, proof, _messageCount)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInbox *SequencerInboxCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInbox *SequencerInboxSession) Rollup() (common.Address, error) {
	return _SequencerInbox.Contract.Rollup(&_SequencerInbox.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInbox *SequencerInboxCallerSession) Rollup() (common.Address, error) {
	return _SequencerInbox.Contract.Rollup(&_SequencerInbox.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_SequencerInbox *SequencerInboxCaller) Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_SequencerInbox *SequencerInboxSession) Sequencer() (common.Address, error) {
	return _SequencerInbox.Contract.Sequencer(&_SequencerInbox.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_SequencerInbox *SequencerInboxCallerSession) Sequencer() (common.Address, error) {
	return _SequencerInbox.Contract.Sequencer(&_SequencerInbox.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) TotalDelayedMessagesRead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "totalDelayedMessagesRead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInbox.Contract.TotalDelayedMessagesRead(&_SequencerInbox.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInbox.Contract.TotalDelayedMessagesRead(&_SequencerInbox.CallOpts)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x1a734229.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2Batch", transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x1a734229.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2Batch(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2Batch(&_SequencerInbox.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x1a734229.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2Batch(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2Batch(&_SequencerInbox.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x44c7cc30.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2BatchFromOrigin(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2BatchFromOrigin", transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x44c7cc30.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2BatchFromOrigin(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInbox.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x44c7cc30.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2BatchFromOrigin(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInbox.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc)
}

// AddSequencerL2BatchFromOriginWithGasRefunder is a paid mutator transaction binding the contract method 0x8a2df18d.
//
// Solidity: function addSequencerL2BatchFromOriginWithGasRefunder(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2BatchFromOriginWithGasRefunder(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2BatchFromOriginWithGasRefunder", transactions, lengths, sectionsMetadata, afterAcc, gasRefunder)
}

// AddSequencerL2BatchFromOriginWithGasRefunder is a paid mutator transaction binding the contract method 0x8a2df18d.
//
// Solidity: function addSequencerL2BatchFromOriginWithGasRefunder(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2BatchFromOriginWithGasRefunder(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOriginWithGasRefunder(&_SequencerInbox.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc, gasRefunder)
}

// AddSequencerL2BatchFromOriginWithGasRefunder is a paid mutator transaction binding the contract method 0x8a2df18d.
//
// Solidity: function addSequencerL2BatchFromOriginWithGasRefunder(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2BatchFromOriginWithGasRefunder(transactions []byte, lengths []*big.Int, sectionsMetadata []*big.Int, afterAcc [32]byte, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOriginWithGasRefunder(&_SequencerInbox.TransactOpts, transactions, lengths, sectionsMetadata, afterAcc, gasRefunder)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x0c4a1e59.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash, bytes32 delayedAcc) returns()
func (_SequencerInbox *SequencerInboxTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte, delayedAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash, delayedAcc)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x0c4a1e59.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash, bytes32 delayedAcc) returns()
func (_SequencerInbox *SequencerInboxSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte, delayedAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ForceInclusion(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash, delayedAcc)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x0c4a1e59.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash, bytes32 delayedAcc) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte, delayedAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ForceInclusion(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash, delayedAcc)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, address _rollup) returns()
func (_SequencerInbox *SequencerInboxTransactor) Initialize(opts *bind.TransactOpts, _delayedInbox common.Address, _sequencer common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "initialize", _delayedInbox, _sequencer, _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, address _rollup) returns()
func (_SequencerInbox *SequencerInboxSession) Initialize(_delayedInbox common.Address, _sequencer common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, _delayedInbox, _sequencer, _rollup)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, address _rollup) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) Initialize(_delayedInbox common.Address, _sequencer common.Address, _rollup common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, _delayedInbox, _sequencer, _rollup)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool newIsSequencer) returns()
func (_SequencerInbox *SequencerInboxTransactor) SetIsSequencer(opts *bind.TransactOpts, addr common.Address, newIsSequencer bool) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "setIsSequencer", addr, newIsSequencer)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool newIsSequencer) returns()
func (_SequencerInbox *SequencerInboxSession) SetIsSequencer(addr common.Address, newIsSequencer bool) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetIsSequencer(&_SequencerInbox.TransactOpts, addr, newIsSequencer)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool newIsSequencer) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) SetIsSequencer(addr common.Address, newIsSequencer bool) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetIsSequencer(&_SequencerInbox.TransactOpts, addr, newIsSequencer)
}

// SetMaxDelay is a paid mutator transaction binding the contract method 0x4d480faa.
//
// Solidity: function setMaxDelay(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxTransactor) SetMaxDelay(opts *bind.TransactOpts, newMaxDelayBlocks *big.Int, newMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "setMaxDelay", newMaxDelayBlocks, newMaxDelaySeconds)
}

// SetMaxDelay is a paid mutator transaction binding the contract method 0x4d480faa.
//
// Solidity: function setMaxDelay(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxSession) SetMaxDelay(newMaxDelayBlocks *big.Int, newMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetMaxDelay(&_SequencerInbox.TransactOpts, newMaxDelayBlocks, newMaxDelaySeconds)
}

// SetMaxDelay is a paid mutator transaction binding the contract method 0x4d480faa.
//
// Solidity: function setMaxDelay(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) SetMaxDelay(newMaxDelayBlocks *big.Int, newMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetMaxDelay(&_SequencerInbox.TransactOpts, newMaxDelayBlocks, newMaxDelaySeconds)
}

// ShutdownForNitro is a paid mutator transaction binding the contract method 0x2aa86909.
//
// Solidity: function shutdownForNitro(uint256 _totalDelayedMessagesRead, bytes32 delayedAcc) returns()
func (_SequencerInbox *SequencerInboxTransactor) ShutdownForNitro(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, delayedAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "shutdownForNitro", _totalDelayedMessagesRead, delayedAcc)
}

// ShutdownForNitro is a paid mutator transaction binding the contract method 0x2aa86909.
//
// Solidity: function shutdownForNitro(uint256 _totalDelayedMessagesRead, bytes32 delayedAcc) returns()
func (_SequencerInbox *SequencerInboxSession) ShutdownForNitro(_totalDelayedMessagesRead *big.Int, delayedAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ShutdownForNitro(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, delayedAcc)
}

// ShutdownForNitro is a paid mutator transaction binding the contract method 0x2aa86909.
//
// Solidity: function shutdownForNitro(uint256 _totalDelayedMessagesRead, bytes32 delayedAcc) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) ShutdownForNitro(_totalDelayedMessagesRead *big.Int, delayedAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ShutdownForNitro(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, delayedAcc)
}

// UndoShutdownForNitro is a paid mutator transaction binding the contract method 0x7e6c255f.
//
// Solidity: function undoShutdownForNitro() returns()
func (_SequencerInbox *SequencerInboxTransactor) UndoShutdownForNitro(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "undoShutdownForNitro")
}

// UndoShutdownForNitro is a paid mutator transaction binding the contract method 0x7e6c255f.
//
// Solidity: function undoShutdownForNitro() returns()
func (_SequencerInbox *SequencerInboxSession) UndoShutdownForNitro() (*types.Transaction, error) {
	return _SequencerInbox.Contract.UndoShutdownForNitro(&_SequencerInbox.TransactOpts)
}

// UndoShutdownForNitro is a paid mutator transaction binding the contract method 0x7e6c255f.
//
// Solidity: function undoShutdownForNitro() returns()
func (_SequencerInbox *SequencerInboxTransactorSession) UndoShutdownForNitro() (*types.Transaction, error) {
	return _SequencerInbox.Contract.UndoShutdownForNitro(&_SequencerInbox.TransactOpts)
}

// SequencerInboxDelayedInboxForcedIterator is returned from FilterDelayedInboxForced and is used to iterate over the raw logs and unpacked data for DelayedInboxForced events raised by the SequencerInbox contract.
type SequencerInboxDelayedInboxForcedIterator struct {
	Event *SequencerInboxDelayedInboxForced // Event containing the contract specifics and raw log

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
func (it *SequencerInboxDelayedInboxForcedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxDelayedInboxForced)
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
		it.Event = new(SequencerInboxDelayedInboxForced)
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
func (it *SequencerInboxDelayedInboxForcedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxDelayedInboxForcedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxDelayedInboxForced represents a DelayedInboxForced event raised by the SequencerInbox contract.
type SequencerInboxDelayedInboxForced struct {
	FirstMessageNum          *big.Int
	BeforeAcc                [32]byte
	NewMessageCount          *big.Int
	TotalDelayedMessagesRead *big.Int
	AfterAccAndDelayed       [2][32]byte
	SeqBatchIndex            *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDelayedInboxForced is a free log retrieval operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) FilterDelayedInboxForced(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*SequencerInboxDelayedInboxForcedIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "DelayedInboxForced", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxDelayedInboxForcedIterator{contract: _SequencerInbox.contract, event: "DelayedInboxForced", logs: logs, sub: sub}, nil
}

// WatchDelayedInboxForced is a free log subscription operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) WatchDelayedInboxForced(opts *bind.WatchOpts, sink chan<- *SequencerInboxDelayedInboxForced, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "DelayedInboxForced", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxDelayedInboxForced)
				if err := _SequencerInbox.contract.UnpackLog(event, "DelayedInboxForced", log); err != nil {
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

// ParseDelayedInboxForced is a log parse operation binding the contract event 0x85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0.
//
// Solidity: event DelayedInboxForced(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) ParseDelayedInboxForced(log types.Log) (*SequencerInboxDelayedInboxForced, error) {
	event := new(SequencerInboxDelayedInboxForced)
	if err := _SequencerInbox.contract.UnpackLog(event, "DelayedInboxForced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxIsSequencerUpdatedIterator is returned from FilterIsSequencerUpdated and is used to iterate over the raw logs and unpacked data for IsSequencerUpdated events raised by the SequencerInbox contract.
type SequencerInboxIsSequencerUpdatedIterator struct {
	Event *SequencerInboxIsSequencerUpdated // Event containing the contract specifics and raw log

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
func (it *SequencerInboxIsSequencerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxIsSequencerUpdated)
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
		it.Event = new(SequencerInboxIsSequencerUpdated)
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
func (it *SequencerInboxIsSequencerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxIsSequencerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxIsSequencerUpdated represents a IsSequencerUpdated event raised by the SequencerInbox contract.
type SequencerInboxIsSequencerUpdated struct {
	Addr        common.Address
	IsSequencer bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIsSequencerUpdated is a free log retrieval operation binding the contract event 0xce86e570206e55533301cb66529b33afbd75e991c575b85adeaca10146be8cb4.
//
// Solidity: event IsSequencerUpdated(address addr, bool isSequencer)
func (_SequencerInbox *SequencerInboxFilterer) FilterIsSequencerUpdated(opts *bind.FilterOpts) (*SequencerInboxIsSequencerUpdatedIterator, error) {

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "IsSequencerUpdated")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxIsSequencerUpdatedIterator{contract: _SequencerInbox.contract, event: "IsSequencerUpdated", logs: logs, sub: sub}, nil
}

// WatchIsSequencerUpdated is a free log subscription operation binding the contract event 0xce86e570206e55533301cb66529b33afbd75e991c575b85adeaca10146be8cb4.
//
// Solidity: event IsSequencerUpdated(address addr, bool isSequencer)
func (_SequencerInbox *SequencerInboxFilterer) WatchIsSequencerUpdated(opts *bind.WatchOpts, sink chan<- *SequencerInboxIsSequencerUpdated) (event.Subscription, error) {

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "IsSequencerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxIsSequencerUpdated)
				if err := _SequencerInbox.contract.UnpackLog(event, "IsSequencerUpdated", log); err != nil {
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

// ParseIsSequencerUpdated is a log parse operation binding the contract event 0xce86e570206e55533301cb66529b33afbd75e991c575b85adeaca10146be8cb4.
//
// Solidity: event IsSequencerUpdated(address addr, bool isSequencer)
func (_SequencerInbox *SequencerInboxFilterer) ParseIsSequencerUpdated(log types.Log) (*SequencerInboxIsSequencerUpdated, error) {
	event := new(SequencerInboxIsSequencerUpdated)
	if err := _SequencerInbox.contract.UnpackLog(event, "IsSequencerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxMaxDelayUpdatedIterator is returned from FilterMaxDelayUpdated and is used to iterate over the raw logs and unpacked data for MaxDelayUpdated events raised by the SequencerInbox contract.
type SequencerInboxMaxDelayUpdatedIterator struct {
	Event *SequencerInboxMaxDelayUpdated // Event containing the contract specifics and raw log

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
func (it *SequencerInboxMaxDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxMaxDelayUpdated)
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
		it.Event = new(SequencerInboxMaxDelayUpdated)
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
func (it *SequencerInboxMaxDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxMaxDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxMaxDelayUpdated represents a MaxDelayUpdated event raised by the SequencerInbox contract.
type SequencerInboxMaxDelayUpdated struct {
	NewMaxDelayBlocks  *big.Int
	NewMaxDelaySeconds *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxDelayUpdated is a free log retrieval operation binding the contract event 0x3bcd3c6d4304309e4b36d94f90517baf304582bb1ac828906808577e067e6b6e.
//
// Solidity: event MaxDelayUpdated(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds)
func (_SequencerInbox *SequencerInboxFilterer) FilterMaxDelayUpdated(opts *bind.FilterOpts) (*SequencerInboxMaxDelayUpdatedIterator, error) {

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "MaxDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxMaxDelayUpdatedIterator{contract: _SequencerInbox.contract, event: "MaxDelayUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxDelayUpdated is a free log subscription operation binding the contract event 0x3bcd3c6d4304309e4b36d94f90517baf304582bb1ac828906808577e067e6b6e.
//
// Solidity: event MaxDelayUpdated(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds)
func (_SequencerInbox *SequencerInboxFilterer) WatchMaxDelayUpdated(opts *bind.WatchOpts, sink chan<- *SequencerInboxMaxDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "MaxDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxMaxDelayUpdated)
				if err := _SequencerInbox.contract.UnpackLog(event, "MaxDelayUpdated", log); err != nil {
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

// ParseMaxDelayUpdated is a log parse operation binding the contract event 0x3bcd3c6d4304309e4b36d94f90517baf304582bb1ac828906808577e067e6b6e.
//
// Solidity: event MaxDelayUpdated(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds)
func (_SequencerInbox *SequencerInboxFilterer) ParseMaxDelayUpdated(log types.Log) (*SequencerInboxMaxDelayUpdated, error) {
	event := new(SequencerInboxMaxDelayUpdated)
	if err := _SequencerInbox.contract.UnpackLog(event, "MaxDelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDeliveredIterator struct {
	Event *SequencerInboxSequencerBatchDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxSequencerBatchDelivered)
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
		it.Event = new(SequencerInboxSequencerBatchDelivered)
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
func (it *SequencerInboxSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDelivered struct {
	FirstMessageNum  *big.Int
	BeforeAcc        [32]byte
	NewMessageCount  *big.Int
	AfterAcc         [32]byte
	Transactions     []byte
	Lengths          []*big.Int
	SectionsMetadata []*big.Int
	SeqBatchIndex    *big.Int
	Sequencer        common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDelivered is a free log retrieval operation binding the contract event 0x3bf85aebd2a1dc6c510ffc4795a3785e786b5817ab30144f88501d4c6456c986.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, uint256 seqBatchIndex, address sequencer)
func (_SequencerInbox *SequencerInboxFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*SequencerInboxSequencerBatchDeliveredIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "SequencerBatchDelivered", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxSequencerBatchDeliveredIterator{contract: _SequencerInbox.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x3bf85aebd2a1dc6c510ffc4795a3785e786b5817ab30144f88501d4c6456c986.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, uint256 seqBatchIndex, address sequencer)
func (_SequencerInbox *SequencerInboxFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxSequencerBatchDelivered, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "SequencerBatchDelivered", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxSequencerBatchDelivered)
				if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
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

// ParseSequencerBatchDelivered is a log parse operation binding the contract event 0x3bf85aebd2a1dc6c510ffc4795a3785e786b5817ab30144f88501d4c6456c986.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, uint256 seqBatchIndex, address sequencer)
func (_SequencerInbox *SequencerInboxFilterer) ParseSequencerBatchDelivered(log types.Log) (*SequencerInboxSequencerBatchDelivered, error) {
	event := new(SequencerInboxSequencerBatchDelivered)
	if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxSequencerBatchDeliveredFromOriginIterator is returned from FilterSequencerBatchDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for SequencerBatchDeliveredFromOrigin events raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDeliveredFromOriginIterator struct {
	Event *SequencerInboxSequencerBatchDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *SequencerInboxSequencerBatchDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxSequencerBatchDeliveredFromOrigin)
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
		it.Event = new(SequencerInboxSequencerBatchDeliveredFromOrigin)
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
func (it *SequencerInboxSequencerBatchDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxSequencerBatchDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxSequencerBatchDeliveredFromOrigin represents a SequencerBatchDeliveredFromOrigin event raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDeliveredFromOrigin struct {
	FirstMessageNum *big.Int
	BeforeAcc       [32]byte
	NewMessageCount *big.Int
	AfterAcc        [32]byte
	SeqBatchIndex   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) FilterSequencerBatchDeliveredFromOrigin(opts *bind.FilterOpts, firstMessageNum []*big.Int, beforeAcc [][32]byte) (*SequencerInboxSequencerBatchDeliveredFromOriginIterator, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "SequencerBatchDeliveredFromOrigin", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxSequencerBatchDeliveredFromOriginIterator{contract: _SequencerInbox.contract, event: "SequencerBatchDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDeliveredFromOrigin is a free log subscription operation binding the contract event 0x10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) WatchSequencerBatchDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *SequencerInboxSequencerBatchDeliveredFromOrigin, firstMessageNum []*big.Int, beforeAcc [][32]byte) (event.Subscription, error) {

	var firstMessageNumRule []interface{}
	for _, firstMessageNumItem := range firstMessageNum {
		firstMessageNumRule = append(firstMessageNumRule, firstMessageNumItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "SequencerBatchDeliveredFromOrigin", firstMessageNumRule, beforeAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxSequencerBatchDeliveredFromOrigin)
				if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDeliveredFromOrigin", log); err != nil {
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

// ParseSequencerBatchDeliveredFromOrigin is a log parse operation binding the contract event 0x10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) ParseSequencerBatchDeliveredFromOrigin(log types.Log) (*SequencerInboxSequencerBatchDeliveredFromOrigin, error) {
	event := new(SequencerInboxSequencerBatchDeliveredFromOrigin)
	if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxShutdownForNitroSetIterator is returned from FilterShutdownForNitroSet and is used to iterate over the raw logs and unpacked data for ShutdownForNitroSet events raised by the SequencerInbox contract.
type SequencerInboxShutdownForNitroSetIterator struct {
	Event *SequencerInboxShutdownForNitroSet // Event containing the contract specifics and raw log

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
func (it *SequencerInboxShutdownForNitroSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxShutdownForNitroSet)
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
		it.Event = new(SequencerInboxShutdownForNitroSet)
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
func (it *SequencerInboxShutdownForNitroSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxShutdownForNitroSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxShutdownForNitroSet represents a ShutdownForNitroSet event raised by the SequencerInbox contract.
type SequencerInboxShutdownForNitroSet struct {
	Shutdown bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterShutdownForNitroSet is a free log retrieval operation binding the contract event 0xe6d1c315c736941d015418a8728891eae6aea39817b9c134486054bf34cc336b.
//
// Solidity: event ShutdownForNitroSet(bool shutdown)
func (_SequencerInbox *SequencerInboxFilterer) FilterShutdownForNitroSet(opts *bind.FilterOpts) (*SequencerInboxShutdownForNitroSetIterator, error) {

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "ShutdownForNitroSet")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxShutdownForNitroSetIterator{contract: _SequencerInbox.contract, event: "ShutdownForNitroSet", logs: logs, sub: sub}, nil
}

// WatchShutdownForNitroSet is a free log subscription operation binding the contract event 0xe6d1c315c736941d015418a8728891eae6aea39817b9c134486054bf34cc336b.
//
// Solidity: event ShutdownForNitroSet(bool shutdown)
func (_SequencerInbox *SequencerInboxFilterer) WatchShutdownForNitroSet(opts *bind.WatchOpts, sink chan<- *SequencerInboxShutdownForNitroSet) (event.Subscription, error) {

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "ShutdownForNitroSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxShutdownForNitroSet)
				if err := _SequencerInbox.contract.UnpackLog(event, "ShutdownForNitroSet", log); err != nil {
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

// ParseShutdownForNitroSet is a log parse operation binding the contract event 0xe6d1c315c736941d015418a8728891eae6aea39817b9c134486054bf34cc336b.
//
// Solidity: event ShutdownForNitroSet(bool shutdown)
func (_SequencerInbox *SequencerInboxFilterer) ParseShutdownForNitroSet(log types.Log) (*SequencerInboxShutdownForNitroSet, error) {
	event := new(SequencerInboxShutdownForNitroSet)
	if err := _SequencerInbox.contract.UnpackLog(event, "ShutdownForNitroSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
