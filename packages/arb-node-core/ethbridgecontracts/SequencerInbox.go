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
const SequencerInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"afterAccAndDelayed\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"DelayedInboxForced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"firstMessageNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqBatchIndex\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"lengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedInbox\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"l1BlockAndTimestamp\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_delayedInbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"inboxCount\",\"type\":\"uint256\"}],\"name\":\"proveBatchContainsSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SequencerInboxBin is the compiled bytecode used for deploying new contracts.
var SequencerInboxBin = "0x608060405234801561001057600080fd5b506000805460ff191660011790556115cf8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638af00545116100715780638af00545146102745780639afc500d146102c1578063b71939b114610391578063d9dd67ab14610399578063e367a2c1146103b6578063eb990c59146103be576100b4565b806306cc91b2146100b95780630a17a46414610140578063342025fa146102125780633dbcc8d11461022c5780635c1bba38146102345780636f791d2914610258575b600080fd5b610127600480360360408110156100cf57600080fd5b810190602081018135600160201b8111156100e957600080fd5b8201836020820111156100fb57600080fd5b803590602001918460018302840111600160201b8311171561011c57600080fd5b9193509150356103fa565b6040805192835260208301919091528051918290030190f35b610210600480360360c081101561015657600080fd5b810190602081018135600160201b81111561017057600080fd5b82018360208201111561018257600080fd5b803590602001918460018302840111600160201b831117156101a357600080fd5b919390929091602081019035600160201b8111156101c057600080fd5b8201836020820111156101d257600080fd5b803590602001918460208302840111600160201b831117156101f357600080fd5b91935091508035906020810135906040810135906060013561054c565b005b61021a610646565b60408051918252519081900360200190f35b61021a61064c565b61023c610652565b604080516001600160a01b039092168252519081900360200190f35b610260610661565b604080519115158252519081900360200190f35b610210600480360361010081101561028b57600080fd5b5080359060ff60208201351690604081019060808101359060a0810135906001600160a01b0360c0820135169060e0013561066a565b610210600480360360c08110156102d757600080fd5b810190602081018135600160201b8111156102f157600080fd5b82018360208201111561030357600080fd5b803590602001918460018302840111600160201b8311171561032457600080fd5b919390929091602081019035600160201b81111561034157600080fd5b82018360208201111561035357600080fd5b803590602001918460208302840111600160201b8311171561037457600080fd5b9193509150803590602081013590604081013590606001356109d9565b61023c610a99565b61021a600480360360208110156103af57600080fd5b5035610aa8565b61021a610ac6565b610210600480360360808110156103d457600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135610acc565b6000808261040d57506000905080610544565b60008061044f87878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509250610b54915050565b90925090506000811561048a5761048188888560018087038154811061047157fe5b9060005260206000200154610bc8565b90935060010190505b60006001838154811061049957fe5b9060005260206000200154905060006104b48a8a8785610bc8565b90955090508288116104fb576040805162461bcd60e51b815260206004820152600b60248201526a10905510d217d4d510549560aa1b604482015290519081900360640190fd5b8088111561053c576040805162461bcd60e51b815260206004820152600960248201526810905510d217d1539160ba1b604482015290519081900360640190fd5b955093505050505b935093915050565b6002546000806105628b8b8b8b8b8b8b8b610d88565b9150915081837f43ca2bb3f5bb808f726cc6c9ebb2c1c26f8bb96a92e4ada823f15ff47138e063600254878f8f8f8f8f8f8f8c6001808054905003604051808c81526020018b8152602001806020018060200188815260200187815260200186815260200185815260200184815260200183810383528c8c82818152602001925080828437600083820152601f01601f191690910184810383528a8152602090810191508b908b0280828437600083820152604051601f909101601f19169092018290039f50909d5050505050505050505050505050a35050505050505050505050565b60075481565b60025481565b6005546001600160a01b031681565b60005460ff1690565b60035487116106b4576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b60006106ca8784883560208a013589898861119c565b6006549091504387359091011061071b576040805162461bcd60e51b815260206004820152601060248201526f4d41585f44454c41595f424c4f434b5360801b604482015290519081900360640190fd5b6007544260208801359091011061076a576040805162461bcd60e51b815260206004820152600e60248201526d4d41585f44454c41595f54494d4560901b604482015290519081900360640190fd5b600060018911156107f457600480546040805163d9dd67ab60e01b81526001198d0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b1580156107c557600080fd5b505afa1580156107d9573d6000803e3d6000fd5b505050506040513d60208110156107ef57600080fd5b505190505b6107fe8183611212565b600480546040805163d9dd67ab60e01b81526000198e0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561084e57600080fd5b505afa158015610862573d6000803e3d6000fd5b505050506040513d602081101561087857600080fd5b5051146108c2576040805162461bcd60e51b81526020600482015260136024820152722222a620aca2a22fa0a1a1aaa6aaa620aa27a960691b604482015290519081900360640190fd5b5050600254600154600090156108f2576001805460001981019081106108e457fe5b906000526020600020015490505b600080600061090484868e434261123e565b92509250925060018290806001815401808255809150506001900390600052602060002001600090919091909150558060028190555083857f85b6a949bf20bfd6bc6e20f98fb490c7944ab61dcfa5a30b5dae543412c9a8a0838f60405180604001604052808881526020018981525060018080549050036040518085815260200184815260200183600260200280838360005b838110156109b0578181015183820152602001610998565b5050505090500182815260200194505050505060405180910390a3505050505050505050505050565b333214610a1b576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600254600080610a318b8b8b8b8b8b8b8b610d88565b6002546001546040805192835260208301899052828101849052600019909101606083015251929450909250839185917f90d3659be0edf0014931e9f8a1c145ec8dbc792776c08a028a148a67700a5812919081900360800190a35050505050505050505050565b6004546001600160a01b031681565b60018181548110610ab557fe5b600091825260209091200154905081565b60065481565b6004546001600160a01b031615610b19576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600480546001600160a01b039586166001600160a01b0319918216179091556005805494909516931692909217909255600691909155600755565b60008082845110158015610b6c575060208385510310155b610ba9576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301610bbd858563ffffffff6114b016565b915091509250929050565b6000806000806000806000610c148b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250610b54915050565b809550819a505050610c5d8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250610b54915050565b809450819a505050610ca68b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250610b54915050565b809350819a505050610cef8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250610b54915050565b604080516020808201989098528082018790526060810186905260808082018490528251808303909101815260a09091019091528051960195909520909950600184019550939050878414610d77576040805162461bcd60e51b815260206004820152600960248201526842415443485f41434360b81b604482015290519081900360640190fd5b509699929850919650505050505050565b600554600090819087906001600160a01b03163314610ddf576040805162461bcd60e51b815260206004820152600e60248201526d27a7262cafa9a2a8aaa2a721a2a960911b604482015290519081900360640190fd5b4360065488011015610e28576040805162461bcd60e51b815260206004820152600d60248201526c109313d0d2d7d513d3d7d3d311609a1b604482015290519081900360640190fd5b43871115610e6d576040805162461bcd60e51b815260206004820152600d60248201526c424c4f434b5f544f4f5f4e455760981b604482015290519081900360640190fd5b4260075487011015610eb5576040805162461bcd60e51b815260206004820152600c60248201526b1512535157d513d3d7d3d31160a21b604482015290519081900360640190fd5b42861115610ef9576040805162461bcd60e51b815260206004820152600c60248201526b54494d455f544f4f5f4e455760a01b604482015290519081900360640190fd5b600354851015610f44576040805162461bcd60e51b815260206004820152601160248201527044454c415945445f4241434b574152445360781b604482015290519081900360640190fd5b6001851015610f8e576040805162461bcd60e51b8152602060048201526011602482015270135554d517d111531056515117d2539255607a1b604482015290519081900360640190fd5b6001600354101580610f9e575080155b610fe9576040805162461bcd60e51b8152602060048201526017602482015276135554d517d111531056515117d253925517d4d5105495604a1b604482015290519081900360640190fd5b600154156110115760018054600019810190811061100357fe5b906000526020600020015492505b600033888860405160200180846001600160a01b03166001600160a01b031660601b815260140183815260200182815260200193505050506040516020818303038152906040528051906020012090506000806110b58e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d8d868a611509565b915091506110c682828a8d8d61123e565b60025492975090935091508111611112576040805162461bcd60e51b815260206004820152600b60248201526a08a9aa0a8b2be8482a886960ab1b604482015290519081900360640190fd5b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601829055600281905586821461118b576040805162461bcd60e51b815260206004820152600960248201526841465445525f41434360b81b604482015290519081900360640190fd5b505050509850989650505050505050565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6bffffffffffffffffffffffff19166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6000806000806003548711156114a6576004805460408051633dbcc8d160e01b815290516001600160a01b0390921692633dbcc8d1928282019260209290829003018186803b15801561129057600080fd5b505afa1580156112a4573d6000803e3d6000fd5b505050506040513d60208110156112ba57600080fd5b5051871115611302576040805162461bcd60e51b815260206004820152600f60248201526e2222a620aca2a22faa27a7afa320a960891b604482015290519081900360640190fd5b600480546040805163d9dd67ab60e01b81526000198b0193810193909352516001600160a01b039091169163d9dd67ab916024808301926020929190829003018186803b15801561135257600080fd5b505afa158015611366573d6000803e3d6000fd5b505050506040513d602081101561137c57600080fd5b81019080805190602001909291905050509050888860035489846040516020018080702232b630bcb2b21036b2b9b9b0b3b2b99d60791b81525060110186815260200185815260200184815260200183815260200182815260200195505050505050604051602081830303815290604052805190602001209850600354870388019750606089896000898960405160200180846001600160a01b03166001600160a01b031660601b8152601401838152602001828152602001935050505060405160208183030381529060405280519060200120838051906020012060405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001209950888060010199505087600381905550505b9895505050505050565b60008160200183511015611500576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6002548190846020880160005b8281101561158c57600089898381811061152c57fe5b60209081029290920135808620604080518086019b909b528a81018a905260608b018d90526080808c01929092528051808c03909201825260a0909a01909952885198909201979097209650600195860195930192919091019050611516565b505050955095935050505056fea2646970667358221220e26dcb05fb6490bdd4f4d8f264b01fc5f45ff0d5f005c8741d5efe80c54be5ef64736f6c634300060b0033"

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

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, uint256 _maxDelayBlocks, uint256 _maxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxTransactor) Initialize(opts *bind.TransactOpts, _delayedInbox common.Address, _sequencer common.Address, _maxDelayBlocks *big.Int, _maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "initialize", _delayedInbox, _sequencer, _maxDelayBlocks, _maxDelaySeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, uint256 _maxDelayBlocks, uint256 _maxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxSession) Initialize(_delayedInbox common.Address, _sequencer common.Address, _maxDelayBlocks *big.Int, _maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, _delayedInbox, _sequencer, _maxDelayBlocks, _maxDelaySeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _delayedInbox, address _sequencer, uint256 _maxDelayBlocks, uint256 _maxDelaySeconds) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) Initialize(_delayedInbox common.Address, _sequencer common.Address, _maxDelayBlocks *big.Int, _maxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, _delayedInbox, _sequencer, _maxDelayBlocks, _maxDelaySeconds)
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
