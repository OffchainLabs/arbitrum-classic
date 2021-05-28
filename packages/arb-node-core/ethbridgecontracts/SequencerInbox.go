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

// SequencerInboxABI is the input ABI used to generate the binding from.
const SequencerInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"afterAccAndDelayed\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"DelayedInboxForced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedInbox\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"l1BlockAndTimestamp\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_delayedInbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inboxCount\",\"type\":\"uint256\"}],\"name\":\"proveBatchContainsSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SequencerInboxBin is the compiled bytecode used for deploying new contracts.
var SequencerInboxBin = "0x608060405234801561001057600080fd5b506000805460ff191660011790556116488061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100bf5760003560e01c80638af005451161007c5780638af005451461027f5780639afc500d146102cc578063b71939b11461039c578063c0c53b8b146103a4578063cb23bcb5146103dc578063d9dd67ab146103e4578063e367a2c114610401576100bf565b806306cc91b2146100c45780630a17a4641461014b578063342025fa1461021d5780633dbcc8d1146102375780635c1bba381461023f5780636f791d2914610263575b600080fd5b610132600480360360408110156100da57600080fd5b810190602081018135600160201b8111156100f457600080fd5b82018360208201111561010657600080fd5b803590602001918460018302840111600160201b8311171561012757600080fd5b919350915035610409565b6040805192835260208301919091528051918290030190f35b61021b600480360360c081101561016157600080fd5b810190602081018135600160201b81111561017b57600080fd5b82018360208201111561018d57600080fd5b803590602001918460018302840111600160201b831117156101ae57600080fd5b919390929091602081019035600160201b8111156101cb57600080fd5b8201836020820111156101dd57600080fd5b803590602001918460208302840111600160201b831117156101fe57600080fd5b91935091508035906020810135906040810135906060013561055b565b005b610225610655565b60408051918252519081900360200190f35b6102256106cb565b6102476106d1565b604080516001600160a01b039092168252519081900360200190f35b61026b6106e0565b604080519115158252519081900360200190f35b61021b600480360361010081101561029657600080fd5b5080359060ff60208201351690604081019060808101359060a0810135906001600160a01b0360c0820135169060e001356106e9565b61021b600480360360c08110156102e257600080fd5b810190602081018135600160201b8111156102fc57600080fd5b82018360208201111561030e57600080fd5b803590602001918460018302840111600160201b8311171561032f57600080fd5b919390929091602081019035600160201b81111561034c57600080fd5b82018360208201111561035e57600080fd5b803590602001918460208302840111600160201b8311171561037f57600080fd5b919350915080359060208101359060408101359060600135610a5d565b610247610b1d565b61021b600480360360608110156103ba57600080fd5b506001600160a01b038135811691602081013582169160409091013516610b2c565b610247610bb8565b610225600480360360208110156103fa57600080fd5b5035610bc7565b610225610be5565b6000808261041c57506000905080610553565b60008061045e87878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509250610c2a915050565b9092509050600081156104995761049088888560018087038154811061048057fe5b9060005260206000200154610c9e565b90935060010190505b6000600183815481106104a857fe5b9060005260206000200154905060006104c38a8a8785610c9e565b909550905082881161050a576040805162461bcd60e51b815260206004820152600b60248201526a10905510d217d4d510549560aa1b604482015290519081900360640190fd5b8088111561054b576040805162461bcd60e51b815260206004820152600960248201526810905510d217d1539160ba1b604482015290519081900360640190fd5b955093505050505b935093915050565b6002546000806105718b8b8b8b8b8b8b8b610e2e565b9150915081837f43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063600254878f8f8f8f8f8f8f8c6001808054905003604051808c81526020018b8152602001806020018060200188815260200187815260200186815260200185815260200184815260200183810383528c8c82818152602001925080828437600083820152601f01601f191690910184810383528a8152602090810191508b908b0280828437600083820152604051601f909101601f19169092018290039f50909d5050505050505050505050505050a35050505050505050505050565b60065460408051632b7759e160e21b815290516000926001600160a01b03169163addd6784916004808301926020929190829003018186803b15801561069a57600080fd5b505afa1580156106ae573d6000803e3d6000fd5b505050506040513d60208110156106c457600080fd5b5051905090565b60025481565b6005546001600160a01b031681565b60005460ff1690565b6003548711610733576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b60006107498784883560208a01358989886112b5565b905043610754610be5565b8735011061079c576040805162461bcd60e51b815260206004820152601060248201526f4d41585f44454c41595f424c4f434b5360801b604482015290519081900360640190fd5b426107a5610655565b602088013501106107ee576040805162461bcd60e51b815260206004820152600e60248201526d4d41585f44454c41595f54494d4560901b604482015290519081900360640190fd5b6000600189111561087857600480546040805163d9dd67ab60e01b81526001198d0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561084957600080fd5b505afa15801561085d573d6000803e3d6000fd5b505050506040513d602081101561087357600080fd5b505190505b610882818361132b565b600480546040805163d9dd67ab60e01b81526000198e0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b1580156108d257600080fd5b505afa1580156108e6573d6000803e3d6000fd5b505050506040513d60208110156108fc57600080fd5b505114610946576040805162461bcd60e51b81526020600482015260136024820152722222a620aca2a22fa0a1a1aaa6aaa620aa27a960691b604482015290519081900360640190fd5b5050600254600154600090156109765760018054600019810190811061096857fe5b906000526020600020015490505b600080600061098884868e4342611357565b92509250925060018290806001815401808255809150506001900390600052602060002001600090919091909150558060028190555083857f85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0838f60405180604001604052808881526020018981525060018080549050036040518085815260200184815260200183600260200280838360005b83811015610a34578181015183820152602001610a1c565b5050505090500182815260200194505050505060405180910390a3505050505050505050505050565b333214610a9f576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600254600080610ab58b8b8b8b8b8b8b8b610e2e565b6002546001546040805192835260208301899052828101849052600019909101606083015251929450909250839185917f90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812919081900360800190a35050505050505050505050565b6004546001600160a01b031681565b6004546001600160a01b031615610b79576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600480546001600160a01b039485166001600160a01b031991821617909155600580549385169382169390931790925560068054919093169116179055565b6006546001600160a01b031681565b60018181548110610bd457fe5b600091825260209091200154905081565b60065460408051630a4147c960e11b815290516000926001600160a01b0316916314828f92916004808301926020929190829003018186803b15801561069a57600080fd5b60008082845110158015610c42575060208385510310155b610c7f576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301610c93858563ffffffff61159116565b915091509250929050565b600080600080600080610ce88a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610c2a915050565b8094508199505050610d318a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610c2a915050565b8093508199505050610d7a8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610c2a915050565b809250819950505082828260405160200180807129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b8152506012018481526020018381526020018281526020019350505050604051602081830303815290604052805190602001209250816001019350868314610e1e576040805162461bcd60e51b815260206004820152600960248201526842415443485f41434360b81b604482015290519081900360640190fd5b5095989197509095505050505050565b60055460009081906001600160a01b03163314610e83576040805162461bcd60e51b815260206004820152600e60248201526d27a7262cafa9a2a8aaa2a721a2a960911b604482015290519081900360640190fd5b43610e8c610be5565b87011015610ed1576040805162461bcd60e51b815260206004820152600d60248201526c109313d0d2d7d513d3d7d3d311609a1b604482015290519081900360640190fd5b43861115610f16576040805162461bcd60e51b815260206004820152600d60248201526c424c4f434b5f544f4f5f4e455760981b604482015290519081900360640190fd5b42610f1f610655565b86011015610f63576040805162461bcd60e51b815260206004820152600c60248201526b1512535157d513d3d7d3d31160a21b604482015290519081900360640190fd5b42851115610fa7576040805162461bcd60e51b815260206004820152600c60248201526b54494d455f544f4f5f4e455760a01b604482015290519081900360640190fd5b600354841015610ff2576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b600184101561103c576040805162461bcd60e51b8152602060048201526011602482015270135554d517d111531056515117d2539255607a1b604482015290519081900360640190fd5b600160035410158061104c575086155b611097576040805162461bcd60e51b8152602060048201526017602482015276135554d517d111531056515117d253925517d4d5105495604a1b604482015290519081900360640190fd5b600154156110bf576001805460001981019081106110b157fe5b906000526020600020015491505b600254826000805b8a8110156111d35760008e838f8f8f868181106110e057fe5b905060200201358601926110f6939291906115ea565b604051808383808284378083019250505092505050604051809103902090506000600390508d8d8481811061112757fe5b905060200201356000141561113a575060065b600061114c82338f8f8b6000896112b5565b905085878260405160200180807129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b81525060120184815260200183815260200182815260200193505050506040516020818303038152906040528051906020012095508e8e858181106111b057fe5b9050602002013585019450868060010197505050505080806001019150506110c7565b506111e18284898c8c611357565b60025492965094509250831161122c576040805162461bcd60e51b815260206004820152600b60248201526a08a9aa0a8b2be8482a886960ab1b604482015290519081900360640190fd5b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60182905560028390558582146112a5576040805162461bcd60e51b815260206004820152600960248201526841465445525f41434360b81b604482015290519081900360640190fd5b5050509850989650505050505050565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6bffffffffffffffffffffffff19166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080600080600354871115611587576004805460408051633dbcc8d160e01b815290516001600160a01b0390921692633dbcc8d1928282019260209290829003018186803b1580156113a957600080fd5b505afa1580156113bd573d6000803e3d6000fd5b505050506040513d60208110156113d357600080fd5b505187111561141b576040805162461bcd60e51b815260206004820152600f60248201526e2222a620aca2a22faa27a7afa320a960891b604482015290519081900360640190fd5b600480546040805163d9dd67ab60e01b81526000198b0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561146b57600080fd5b505afa15801561147f573d6000803e3d6000fd5b505050506040513d602081101561149557600080fd5b505160035460408051702232b630bcb2b21036b2b9b9b0b3b2b99d60791b602082810191909152603182019d909d52605181018c905260718101839052609181018b905260b18082018590528251808303909101815260d190910190915280519b019a909a2060608051608020919b8a03909a01999192506000906115269060069083908b908b908f9084906112b5565b604080517129b2b8bab2b731b2b91036b2b9b9b0b3b29d60711b602080830191909152603282019e909e52605281018d905260728082019390935281518082039093018352609201905280519b019a909a2060038990559950506001909701965b9895505050505050565b600081602001835110156115e1576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b600080858511156115f9578182fd5b83861115611605578182fd5b505082019391909203915056fea2646970667358221220781458722eb17fca715f78af14fb8ac03bf3491e2cb06dd75478fa23dcca4c0f64736f6c634300060b0033"

// DeploySequencerInbox deploys a new Ethereum contract, binding an instance of SequencerInbox to it.
func DeploySequencerInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SequencerInbox, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerInboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SequencerInboxBin), backend)
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

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxCaller) ProveBatchContainsSequenceNumber(opts *bind.CallOpts, proof []byte, inboxCount *big.Int) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "proveBatchContainsSequenceNumber", proof, inboxCount)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxSession) ProveBatchContainsSequenceNumber(proof []byte, inboxCount *big.Int) (*big.Int, [32]byte, error) {
	return _SequencerInbox.Contract.ProveBatchContainsSequenceNumber(&_SequencerInbox.CallOpts, proof, inboxCount)
}

// ProveBatchContainsSequenceNumber is a free data retrieval call binding the contract method 0x06cc91b2.
//
// Solidity: function proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) view returns(uint256, bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) ProveBatchContainsSequenceNumber(proof []byte, inboxCount *big.Int) (*big.Int, [32]byte, error) {
	return _SequencerInbox.Contract.ProveBatchContainsSequenceNumber(&_SequencerInbox.CallOpts, proof, inboxCount)
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

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x0a17a464.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2Batch", transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x0a17a464.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2Batch(transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2Batch(&_SequencerInbox.TransactOpts, transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x0a17a464.
//
// Solidity: function addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2Batch(transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2Batch(&_SequencerInbox.TransactOpts, transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x9afc500d.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2BatchFromOrigin(opts *bind.TransactOpts, transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2BatchFromOrigin", transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x9afc500d.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2BatchFromOrigin(transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInbox.TransactOpts, transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x9afc500d.
//
// Solidity: function addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2BatchFromOrigin(transactions []byte, lengths []*big.Int, l1BlockNumber *big.Int, timestamp *big.Int, _totalDelayedMessagesRead *big.Int, afterAcc [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInbox.TransactOpts, transactions, lengths, l1BlockNumber, timestamp, _totalDelayedMessagesRead, afterAcc)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x8af00545.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x8af00545.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ForceInclusion(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0x8af00545.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTimestamp [2]*big.Int, inboxSeqNum *big.Int, gasPriceL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ForceInclusion(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTimestamp, inboxSeqNum, gasPriceL1, sender, messageDataHash)
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
	FirstMessageNum          *big.Int
	BeforeAcc                [32]byte
	NewMessageCount          *big.Int
	AfterAcc                 [32]byte
	Transactions             []byte
	Lengths                  []*big.Int
	L1BlockNumber            *big.Int
	Timestamp                *big.Int
	TotalDelayedMessagesRead *big.Int
	DelayedAcc               [32]byte
	SeqBatchIndex            *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDelivered is a free log retrieval operation binding the contract event 0x43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 totalDelayedMessagesRead, bytes32 delayedAcc, uint256 seqBatchIndex)
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

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 totalDelayedMessagesRead, bytes32 delayedAcc, uint256 seqBatchIndex)
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

// ParseSequencerBatchDelivered is a log parse operation binding the contract event 0x43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 totalDelayedMessagesRead, bytes32 delayedAcc, uint256 seqBatchIndex)
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
	DelayedAcc      [32]byte
	SeqBatchIndex   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDeliveredFromOrigin is a free log retrieval operation binding the contract event 0x90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes32 delayedAcc, uint256 seqBatchIndex)
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

// WatchSequencerBatchDeliveredFromOrigin is a free log subscription operation binding the contract event 0x90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes32 delayedAcc, uint256 seqBatchIndex)
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

// ParseSequencerBatchDeliveredFromOrigin is a log parse operation binding the contract event 0x90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812.
//
// Solidity: event SequencerBatchDeliveredFromOrigin(uint256 indexed firstMessageNum, bytes32 indexed beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes32 delayedAcc, uint256 seqBatchIndex)
func (_SequencerInbox *SequencerInboxFilterer) ParseSequencerBatchDeliveredFromOrigin(log types.Log) (*SequencerInboxSequencerBatchDeliveredFromOrigin, error) {
	event := new(SequencerInboxSequencerBatchDeliveredFromOrigin)
	if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
