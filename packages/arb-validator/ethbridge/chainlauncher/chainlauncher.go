// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlauncher

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

// ArbChainABI is the input ABI used to generate the binding from.
const ArbChainABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[3]\",\"name\":\"fields\",\"type\":\"bytes32[3]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"bytes21[]\",\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastMessageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeManager\",\"outputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"bytes\",\"name\":\"_messageData\",\"type\":\"bytes\"},{\"internalType\":\"uint16[]\",\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_messageDestinations\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumVM.State\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"_fields\",\"type\":\"bytes32[4]\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_messageDataHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_messageDestinations\",\"type\":\"address[]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"enumVM.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"inChallenge\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbChainFuncSigs maps the 4-byte function signature to its string representation.
var ArbChainFuncSigs = map[string]string{
	"94af716b": "activateVM()",
	"023a96fe": "challengeManager()",
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"4526c5d9": "confirmDisputableAsserted(bytes32,bytes32,uint32,bytes21[],bytes,uint16[],uint256[],address[],bytes32)",
	"08dc89d7": "currentDeposit(address)",
	"aca0f372": "escrowRequired()",
	"6be00229": "exitAddress()",
	"1865c57d": "getState()",
	"d489113a": "globalInbox()",
	"05b050de": "increaseDeposit()",
	"2782e87e": "initiateChallenge(bytes32)",
	"8da5cb5b": "owner()",
	"cfa80707": "ownerShutdown()",
	"97e2e256": "pendingDisputableAssert(bytes32[4],uint32,uint64[2],bytes21[],bytes32[],uint16[],uint256[],address[])",
	"60675a87": "terminateAddress()",
	"3a768463": "vm()",
}

// ArbChainBin is the compiled bytecode used for deploying new contracts.
var ArbChainBin = "0x608060405234801561001057600080fd5b50604051620022f4380380620022f4833981810160405260e081101561003557600080fd5b50805160208201516040808401516060850151608086015160a087015160c090970151600180546001600160a01b038084166001600160a01b0319928316179283905560008054828d16931692909217825587517ff39723830000000000000000000000000000000000000000000000000000000081529751999a989996989597949693948b948b948b948b948b948b948b949092169263f39723839260048084019382900301818387803b1580156100ed57600080fd5b505af1158015610101573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b15801561018857600080fd5b505af415801561019c573d6000803e3d6000fd5b505050506040513d60208110156101b257600080fd5b50516004555050600680546001600160801b0319166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161763ffffffff60201b19166401000000009290931691909102919091178082556001925060ff60401b191668010000000000000000830217905550505050505050506120b480620002406000396000f3fe6080604052600436106100f35760003560e01c806360675a871161008a57806397e2e2561161005957806397e2e256146105db578063aca0f372146108ff578063cfa8070714610914578063d489113a14610929576100f3565b806360675a87146105875780636be002291461059c5780638da5cb5b146105b157806394af716b146105c6576100f3565b806322c091bc116100c657806322c091bc146101b15780632782e87e146101de5780633a768463146102085780634526c5d9146102b2576100f3565b8063023a96fe146100f857806305b050de1461012957806308dc89d7146101335780631865c57d14610178575b600080fd5b34801561010457600080fd5b5061010d61093e565b604080516001600160a01b039092168252519081900360200190f35b61013161094d565b005b34801561013f57600080fd5b506101666004803603602081101561015657600080fd5b50356001600160a01b0316610964565b60408051918252519081900360200190f35b34801561018457600080fd5b5061018d610983565b6040518082600381111561019d57fe5b60ff16815260200191505060405180910390f35b3480156101bd57600080fd5b50610131600480360360808110156101d457600080fd5b5060408101610993565b3480156101ea57600080fd5b506101316004803603602081101561020157600080fd5b5035610ae6565b34801561021457600080fd5b5061021d610cea565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561028d57fe5b60ff1681529115156020830152506040805191829003019a5098505050505050505050f35b3480156102be57600080fd5b5061013160048036036101208110156102d657600080fd5b81359160208101359163ffffffff6040830135169190810190608081016060820135600160201b81111561030957600080fd5b82018360208201111561031b57600080fd5b803590602001918460208302840111600160201b8311171561033c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561038b57600080fd5b82018360208201111561039d57600080fd5b803590602001918460018302840111600160201b831117156103be57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561041057600080fd5b82018360208201111561042257600080fd5b803590602001918460208302840111600160201b8311171561044357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561049257600080fd5b8201836020820111156104a457600080fd5b803590602001918460208302840111600160201b831117156104c557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561051457600080fd5b82018360208201111561052657600080fd5b803590602001918460208302840111600160201b8311171561054757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250610d5b915050565b34801561059357600080fd5b5061010d610fd7565b3480156105a857600080fd5b5061010d610fe6565b3480156105bd57600080fd5b5061010d610ff5565b3480156105d257600080fd5b50610131611004565b3480156105e757600080fd5b5061013160048036036101808110156105ff57600080fd5b810190808060800190600480602002604051908101604052809291908260046020028082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561068657600080fd5b82018360208201111561069857600080fd5b803590602001918460208302840111600160201b831117156106b957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561070857600080fd5b82018360208201111561071a57600080fd5b803590602001918460208302840111600160201b8311171561073b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561078a57600080fd5b82018360208201111561079c57600080fd5b803590602001918460208302840111600160201b831117156107bd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561080c57600080fd5b82018360208201111561081e57600080fd5b803590602001918460208302840111600160201b8311171561083f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561088e57600080fd5b8201836020820111156108a057600080fd5b803590602001918460208302840111600160201b831117156108c157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611097945050505050565b34801561090b57600080fd5b506101666117cb565b34801561092057600080fd5b506101316117da565b34801561093557600080fd5b5061010d61183a565b6000546001600160a01b031681565b336000908152600860205260409020805434019055565b6001600160a01b0381166000908152600860205260409020545b919050565b600754600160401b900460ff1690565b6000546001600160a01b031633146109dc5760405162461bcd60e51b815260040180806020018281038252602d815260200180611ffe602d913960400191505060405180910390fd5b600754600160481b900460ff16610a245760405162461bcd60e51b8152600401808060200182810382526026815260200180611fd86026913960400191505060405180910390fd5b6007805469ff00000000000000000019169055610a896001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000205461184990919063ffffffff16565b82356001600160a01b03166000908152600860208181526040832093909355610ac1928401356001600160801b031691856001610a4c565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b03161115610b3f5760405162461bcd60e51b8152600401808060200182810382526027815260200180611fb16027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b158015610bb957600080fd5b505af4158015610bcd573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b0316808252918101919091526007548451630823813560e21b815292909516975063208e04d496509194919363ffffffff16928892600490920191829187918190849084905b83811015610c65578181015183820152602001610c4d565b5050505090500184600260200280838360005b83811015610c90578181015183820152602001610c78565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050600060405180830381600087803b158015610ccf57600080fd5b505af1158015610ce3573d6000803e3d6000fd5b5050505050565b6002546003546004546005546006546007546001600160a01b03909216916001600160801b0382169167ffffffffffffffff600160801b8204811692600160c01b909204169063ffffffff80821691600160201b81049091169060ff600160401b8204811691600160481b9004168b565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63924e7b3760028b8b8b8b8b8b8b8b8b6040518b63ffffffff1660e01b8152600401808b81526020018a81526020018981526020018863ffffffff1663ffffffff168152602001806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b83811015610e09578181015183820152602001610df1565b5050505090500186810385528b818151815260200191508051906020019080838360005b83811015610e45578181015183820152602001610e2d565b50505050905090810190601f168015610e725780820380516001836020036101000a031916815260200191505b5086810384528a5181528a51602091820191808d01910280838360005b83811015610ea7578181015183820152602001610e8f565b50505050905001868103835289818151815260200191508051906020019060200280838360005b83811015610ee6578181015183820152602001610ece565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015610f25578181015183820152602001610f0d565b505050509050019f5050505050505050505050505050505060006040518083038186803b158015610f5557600080fd5b505af4158015610f69573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610fa4935091506001600160801b031663ffffffff61184916565b6005546001600160a01b0316600090815260086020526040902055610fcc86868686866118aa565b505050505050505050565b600a546001600160a01b031681565b6009546001600160a01b031681565b600b546001600160a01b031681565b600b546001600160a01b0316331461105c576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff16600381111561107657fe5b1415611095576007805468ff00000000000000001916600160401b1790555b565b336000908152600860205260409020546006546001600160801b031611156110f05760405162461bcd60e51b815260040180806020018281038252603181526020018061202b6031913960400191505060405180910390fd5b60065433600090815260086020908152604080832080546001600160801b039095169094039093559151630f89fbff60e01b8152606060048201818152895160648401528951919473__$9836fa7140e5a33041d4b827682e675a30$__94630f89fbff948c948b948b9490938493602481019360448201936084909201928a83019291909102908190849084905b8381101561119657818101518382015260200161117e565b50505050905001848103835286818151815260200191508051906020019060200280838360005b838110156111d55781810151838201526020016111bd565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156112145781810151838201526020016111fc565b50505050905001965050505050505060006040518083038186803b15801561123b57600080fd5b505af415801561124f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561127857600080fd5b8101908080516040519392919084600160201b82111561129757600080fd5b9083019060208201858111156112ac57600080fd5b82518660208202830111600160201b821117156112c857600080fd5b82525081516020918201928201910280838360005b838110156112f55781810151838201526020016112dd565b50505050905001604052505050905073__$9836fa7140e5a33041d4b827682e675a30$__63af17d92287836040518363ffffffff1660e01b8152600401808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561137357818101518382015260200161135b565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156113b257818101518382015260200161139a565b5050505090500194505050505060206040518083038186803b1580156113d757600080fd5b505af41580156113eb573d6000803e3d6000fd5b505050506040513d602081101561140157600080fd5b505161143e5760405162461bcd60e51b815260040180806020018281038252602481526020018061205c6024913960400191505060405180910390fd5b60015460405163565b19db60e11b815230600482018181526060602484019081528a5160648501528a516001600160a01b039095169463acb633b6948c9388939092909160448101916084909101906020808801910280838360005b838110156114b257818101518382015260200161149a565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156114f15781810151838201526020016114d9565b505050509050019550505050505060206040518083038186803b15801561151757600080fd5b505afa15801561152b573d6000803e3d6000fd5b505050506040513d602081101561154157600080fd5b5051611594576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63c97c8eec60028b8b8b8b8b8b8b8b6040518a63ffffffff1660e01b8152600401808a815260200189600460200280838360005b838110156115f55781810151838201526020016115dd565b5050505063ffffffff8b1692019182525060200187604080838360005b8381101561162a578181015183820152602001611612565b50505050905001806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b8381101561167d578181015183820152602001611665565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b838110156116bc5781810151838201526020016116a4565b50505050905001868103845289818151815260200191508051906020019060200280838360005b838110156116fb5781810151838201526020016116e3565b50505050905001868103835288818151815260200191508051906020019060200280838360005b8381101561173a578181015183820152602001611722565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611779578181015183820152602001611761565b505050509050019e50505050505050505050505050505060006040518083038186803b1580156117a857600080fd5b505af41580156117bc573d6000803e3d6000fd5b50505050505050505050505050565b6006546001600160801b031690565b600b546001600160a01b03163314611832576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b611095611bd1565b6001546001600160a01b031681565b6000828201838110156118a3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b1580156118f057600080fd5b505af1158015611904573d6000803e3d6000fd5b505050506040513d602081101561191a57600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b15801561196b57600080fd5b505af415801561197f573d6000803e3d6000fd5b505050506040513d602081101561199557600080fd5b505181146119e0576119dc60405180606001604052806119b56001611bdf565b81526020016119c76002800154611c39565b81526020016119d584611c39565b9052611c93565b6004555b600160009054906101000a90046001600160a01b03166001600160a01b031663ec22a76787878787876040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015611a68578181015183820152602001611a50565b5050505090500186810385528a818151815260200191508051906020019080838360005b83811015611aa4578181015183820152602001611a8c565b50505050905090810190601f168015611ad15780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015611b06578181015183820152602001611aee565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015611b45578181015183820152602001611b2d565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611b84578181015183820152602001611b6c565b505050509050019a5050505050505050505050600060405180830381600087803b158015611bb157600080fd5b505af1158015611bc5573d6000803e3d6000fd5b50505050505050505050565b600b546001600160a01b0316ff5b611be7611f7a565b604080516060810182528381528151600080825260208281019094529192830191611c28565b611c15611f7a565b815260200190600190039081611c0d5790505b508152600060209091015292915050565b611c41611f7a565b604080516060810182528381528151600080825260208281019094529192830191611c82565b611c6f611f7a565b815260200190600190039081611c675790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b611cb7611f7a565b815260200190600190039081611caf575050805190915060005b81811015611d0957848160038110611ce557fe5b6020020151838281518110611cf657fe5b6020908102919091010152600101611cd1565b50611d1382611d1b565b949350505050565b6000600882511115611d6b576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611d98578160200160208202803883390190505b50805190915060005b81811015611df457611db1611f9e565b611dcd868381518110611dc057fe5b6020026020010151611e67565b90508060000151848381518110611de057fe5b602090810291909101015250600101611da1565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611e3d578181015183820152602001611e25565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b611e6f611f9e565b6040820151600c60ff90911610611ec1576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff16611eee576040518060200160405280611ee58460000151611f56565b9052905061097e565b604082015160ff1660021415611f13575060408051602081019091528151815261097e565b600360ff16826040015160ff1610158015611f3757506040820151600c60ff909116105b15611f54576040518060200160405280611ee58460200151611d1b565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274546f6b656e207479706573206d7573742062652076616c696420616e6420736f72746564a265627a7a7231582023f0161eb695a83cbe5a70f160e46fba3851d099415def78a48000d8a2d77ff064736f6c634300050c0032"

// DeployArbChain deploys a new Ethereum contract, binding an instance of ArbChain to it.
func DeployArbChain(auth *bind.TransactOpts, backend bind.ContractBackend, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeManagerAddress common.Address, _globalInboxAddress common.Address) (common.Address, *types.Transaction, *ArbChain, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ArbChainBin = strings.Replace(ArbChainBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	ArbChainBin = strings.Replace(ArbChainBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbChainBin = strings.Replace(ArbChainBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbChainBin), backend, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeManagerAddress, _globalInboxAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbChain{ArbChainCaller: ArbChainCaller{contract: contract}, ArbChainTransactor: ArbChainTransactor{contract: contract}, ArbChainFilterer: ArbChainFilterer{contract: contract}}, nil
}

// ArbChain is an auto generated Go binding around an Ethereum contract.
type ArbChain struct {
	ArbChainCaller     // Read-only binding to the contract
	ArbChainTransactor // Write-only binding to the contract
	ArbChainFilterer   // Log filterer for contract events
}

// ArbChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbChainSession struct {
	Contract     *ArbChain         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbChainCallerSession struct {
	Contract *ArbChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbChainTransactorSession struct {
	Contract     *ArbChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbChainRaw struct {
	Contract *ArbChain // Generic contract binding to access the raw methods on
}

// ArbChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbChainCallerRaw struct {
	Contract *ArbChainCaller // Generic read-only contract binding to access the raw methods on
}

// ArbChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbChainTransactorRaw struct {
	Contract *ArbChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbChain creates a new instance of ArbChain, bound to a specific deployed contract.
func NewArbChain(address common.Address, backend bind.ContractBackend) (*ArbChain, error) {
	contract, err := bindArbChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbChain{ArbChainCaller: ArbChainCaller{contract: contract}, ArbChainTransactor: ArbChainTransactor{contract: contract}, ArbChainFilterer: ArbChainFilterer{contract: contract}}, nil
}

// NewArbChainCaller creates a new read-only instance of ArbChain, bound to a specific deployed contract.
func NewArbChainCaller(address common.Address, caller bind.ContractCaller) (*ArbChainCaller, error) {
	contract, err := bindArbChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbChainCaller{contract: contract}, nil
}

// NewArbChainTransactor creates a new write-only instance of ArbChain, bound to a specific deployed contract.
func NewArbChainTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbChainTransactor, error) {
	contract, err := bindArbChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbChainTransactor{contract: contract}, nil
}

// NewArbChainFilterer creates a new log filterer instance of ArbChain, bound to a specific deployed contract.
func NewArbChainFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbChainFilterer, error) {
	contract, err := bindArbChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbChainFilterer{contract: contract}, nil
}

// bindArbChain binds a generic wrapper to an already deployed contract.
func bindArbChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbChain *ArbChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbChain.Contract.ArbChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbChain *ArbChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChain.Contract.ArbChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbChain *ArbChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbChain.Contract.ArbChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbChain *ArbChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbChain *ArbChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbChain *ArbChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbChain.Contract.contract.Transact(opts, method, params...)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbChain *ArbChainCaller) ChallengeManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "challengeManager")
	return *ret0, err
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbChain *ArbChainSession) ChallengeManager() (common.Address, error) {
	return _ArbChain.Contract.ChallengeManager(&_ArbChain.CallOpts)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbChain *ArbChainCallerSession) ChallengeManager() (common.Address, error) {
	return _ArbChain.Contract.ChallengeManager(&_ArbChain.CallOpts)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbChain *ArbChainCaller) CurrentDeposit(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "currentDeposit", validator)
	return *ret0, err
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbChain *ArbChainSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _ArbChain.Contract.CurrentDeposit(&_ArbChain.CallOpts, validator)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbChain *ArbChainCallerSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _ArbChain.Contract.CurrentDeposit(&_ArbChain.CallOpts, validator)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbChain *ArbChainCaller) EscrowRequired(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "escrowRequired")
	return *ret0, err
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbChain *ArbChainSession) EscrowRequired() (*big.Int, error) {
	return _ArbChain.Contract.EscrowRequired(&_ArbChain.CallOpts)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbChain *ArbChainCallerSession) EscrowRequired() (*big.Int, error) {
	return _ArbChain.Contract.EscrowRequired(&_ArbChain.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbChain *ArbChainCaller) ExitAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "exitAddress")
	return *ret0, err
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbChain *ArbChainSession) ExitAddress() (common.Address, error) {
	return _ArbChain.Contract.ExitAddress(&_ArbChain.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbChain *ArbChainCallerSession) ExitAddress() (common.Address, error) {
	return _ArbChain.Contract.ExitAddress(&_ArbChain.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbChain *ArbChainCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "getState")
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbChain *ArbChainSession) GetState() (uint8, error) {
	return _ArbChain.Contract.GetState(&_ArbChain.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbChain *ArbChainCallerSession) GetState() (uint8, error) {
	return _ArbChain.Contract.GetState(&_ArbChain.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbChain *ArbChainCaller) GlobalInbox(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "globalInbox")
	return *ret0, err
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbChain *ArbChainSession) GlobalInbox() (common.Address, error) {
	return _ArbChain.Contract.GlobalInbox(&_ArbChain.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbChain *ArbChainCallerSession) GlobalInbox() (common.Address, error) {
	return _ArbChain.Contract.GlobalInbox(&_ArbChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbChain *ArbChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbChain *ArbChainSession) Owner() (common.Address, error) {
	return _ArbChain.Contract.Owner(&_ArbChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbChain *ArbChainCallerSession) Owner() (common.Address, error) {
	return _ArbChain.Contract.Owner(&_ArbChain.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbChain *ArbChainCaller) TerminateAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "terminateAddress")
	return *ret0, err
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbChain *ArbChainSession) TerminateAddress() (common.Address, error) {
	return _ArbChain.Contract.TerminateAddress(&_ArbChain.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbChain *ArbChainCallerSession) TerminateAddress() (common.Address, error) {
	return _ArbChain.Contract.TerminateAddress(&_ArbChain.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbChain *ArbChainCaller) Vm(opts *bind.CallOpts) (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	State             uint8
	InChallenge       bool
}, error) {
	ret := new(struct {
		MachineHash       [32]byte
		PendingHash       [32]byte
		Inbox             [32]byte
		Asserter          common.Address
		EscrowRequired    *big.Int
		Deadline          uint64
		SequenceNum       uint64
		GracePeriod       uint32
		MaxExecutionSteps uint32
		State             uint8
		InChallenge       bool
	})
	out := ret
	err := _ArbChain.contract.Call(opts, out, "vm")
	return *ret, err
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbChain *ArbChainSession) Vm() (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	State             uint8
	InChallenge       bool
}, error) {
	return _ArbChain.Contract.Vm(&_ArbChain.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbChain *ArbChainCallerSession) Vm() (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	State             uint8
	InChallenge       bool
}, error) {
	return _ArbChain.Contract.Vm(&_ArbChain.CallOpts)
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbChain *ArbChainTransactor) ActivateVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "activateVM")
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbChain *ArbChainSession) ActivateVM() (*types.Transaction, error) {
	return _ArbChain.Contract.ActivateVM(&_ArbChain.TransactOpts)
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbChain *ArbChainTransactorSession) ActivateVM() (*types.Transaction, error) {
	return _ArbChain.Contract.ActivateVM(&_ArbChain.TransactOpts)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbChain *ArbChainTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbChain *ArbChainSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbChain.Contract.CompleteChallenge(&_ArbChain.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbChain *ArbChainTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbChain.Contract.CompleteChallenge(&_ArbChain.TransactOpts, _players, _rewards)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "confirmDisputableAsserted", _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.ConfirmDisputableAsserted(&_ArbChain.TransactOpts, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainTransactorSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.ConfirmDisputableAsserted(&_ArbChain.TransactOpts, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_ArbChain *ArbChainTransactor) IncreaseDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "increaseDeposit")
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_ArbChain *ArbChainSession) IncreaseDeposit() (*types.Transaction, error) {
	return _ArbChain.Contract.IncreaseDeposit(&_ArbChain.TransactOpts)
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_ArbChain *ArbChainTransactorSession) IncreaseDeposit() (*types.Transaction, error) {
	return _ArbChain.Contract.IncreaseDeposit(&_ArbChain.TransactOpts)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbChain *ArbChainTransactor) InitiateChallenge(opts *bind.TransactOpts, _assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "initiateChallenge", _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbChain *ArbChainSession) InitiateChallenge(_assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.InitiateChallenge(&_ArbChain.TransactOpts, _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbChain *ArbChainTransactorSession) InitiateChallenge(_assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.InitiateChallenge(&_ArbChain.TransactOpts, _assertPreHash)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbChain *ArbChainTransactor) OwnerShutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "ownerShutdown")
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbChain *ArbChainSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbChain.Contract.OwnerShutdown(&_ArbChain.TransactOpts)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbChain *ArbChainTransactorSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbChain.Contract.OwnerShutdown(&_ArbChain.TransactOpts)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_ArbChain *ArbChainTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "pendingDisputableAssert", _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_ArbChain *ArbChainSession) PendingDisputableAssert(_fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _ArbChain.Contract.PendingDisputableAssert(&_ArbChain.TransactOpts, _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_ArbChain *ArbChainTransactorSession) PendingDisputableAssert(_fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _ArbChain.Contract.PendingDisputableAssert(&_ArbChain.TransactOpts, _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// ArbChainConfirmedDisputableAssertionIterator is returned from FilterConfirmedDisputableAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedDisputableAssertion events raised by the ArbChain contract.
type ArbChainConfirmedDisputableAssertionIterator struct {
	Event *ArbChainConfirmedDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChainConfirmedDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChainConfirmedDisputableAssertion)
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
		it.Event = new(ArbChainConfirmedDisputableAssertion)
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
func (it *ArbChainConfirmedDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChainConfirmedDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChainConfirmedDisputableAssertion represents a ConfirmedDisputableAssertion event raised by the ArbChain contract.
type ArbChainConfirmedDisputableAssertion struct {
	NewState    [32]byte
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbChain *ArbChainFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts) (*ArbChainConfirmedDisputableAssertionIterator, error) {

	logs, sub, err := _ArbChain.contract.FilterLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChainConfirmedDisputableAssertionIterator{contract: _ArbChain.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbChain *ArbChainFilterer) WatchConfirmedDisputableAssertion(opts *bind.WatchOpts, sink chan<- *ArbChainConfirmedDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChain.contract.WatchLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChainConfirmedDisputableAssertion)
				if err := _ArbChain.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
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

// ParseConfirmedDisputableAssertion is a log parse operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbChain *ArbChainFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*ArbChainConfirmedDisputableAssertion, error) {
	event := new(ArbChainConfirmedDisputableAssertion)
	if err := _ArbChain.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChainInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ArbChain contract.
type ArbChainInitiatedChallengeIterator struct {
	Event *ArbChainInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ArbChainInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChainInitiatedChallenge)
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
		it.Event = new(ArbChainInitiatedChallenge)
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
func (it *ArbChainInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChainInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChainInitiatedChallenge represents a InitiatedChallenge event raised by the ArbChain contract.
type ArbChainInitiatedChallenge struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_ArbChain *ArbChainFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ArbChainInitiatedChallengeIterator, error) {

	logs, sub, err := _ArbChain.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ArbChainInitiatedChallengeIterator{contract: _ArbChain.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_ArbChain *ArbChainFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ArbChainInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _ArbChain.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChainInitiatedChallenge)
				if err := _ArbChain.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_ArbChain *ArbChainFilterer) ParseInitiatedChallenge(log types.Log) (*ArbChainInitiatedChallenge, error) {
	event := new(ArbChainInitiatedChallenge)
	if err := _ArbChain.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChainPendingDisputableAssertionIterator is returned from FilterPendingDisputableAssertion and is used to iterate over the raw logs and unpacked data for PendingDisputableAssertion events raised by the ArbChain contract.
type ArbChainPendingDisputableAssertionIterator struct {
	Event *ArbChainPendingDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChainPendingDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChainPendingDisputableAssertion)
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
		it.Event = new(ArbChainPendingDisputableAssertion)
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
func (it *ArbChainPendingDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChainPendingDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChainPendingDisputableAssertion represents a PendingDisputableAssertion event raised by the ArbChain contract.
type ArbChainPendingDisputableAssertion struct {
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

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_ArbChain *ArbChainFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*ArbChainPendingDisputableAssertionIterator, error) {

	logs, sub, err := _ArbChain.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChainPendingDisputableAssertionIterator{contract: _ArbChain.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_ArbChain *ArbChainFilterer) WatchPendingDisputableAssertion(opts *bind.WatchOpts, sink chan<- *ArbChainPendingDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChain.contract.WatchLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChainPendingDisputableAssertion)
				if err := _ArbChain.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_ArbChain *ArbChainFilterer) ParsePendingDisputableAssertion(log types.Log) (*ArbChainPendingDisputableAssertion, error) {
	event := new(ArbChainPendingDisputableAssertion)
	if err := _ArbChain.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbProtocolABI is the input ABI used to generate the binding from.
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"beforeBalancesValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalMessageValueAmounts\",\"type\":\"uint256[]\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"af17d922": "beforeBalancesValid(bytes21[],uint256[])",
	"0f89fbff": "calculateBeforeValues(bytes21[],uint16[],uint256[])",
	"20903721": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32,uint256[])",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"3e285598": "generatePreconditionHash(bytes32,uint64[2],bytes32,bytes21[],uint256[])",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x611085610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100605760003560e01c80624c28f6146100655780630f89fbff146100bd57806320903721146102b25780633e2855981461037d578063af17d922146104db575b600080fd5b6100ab6004803603608081101561007b57600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b0316610612565b60408051918252519081900360200190f35b610262600480360360608110156100d357600080fd5b810190602081018135600160201b8111156100ed57600080fd5b8201836020820111156100ff57600080fd5b803590602001918460208302840111600160201b8311171561012057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460208302840111600160201b831117156101a257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101f157600080fd5b82018360208201111561020357600080fd5b803590602001918460208302840111600160201b8311171561022457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610702945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561029e578181015183820152602001610286565b505050509050019250505060405180910390f35b6100ab600480360360e08110156102c857600080fd5b81359163ffffffff6020820135169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561030c57600080fd5b82018360208201111561031e57600080fd5b803590602001918460208302840111600160201b8311171561033f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108e7945050505050565b6100ab600480360360c081101561039357600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594909390925060408101915060200135600160201b8111156103e857600080fd5b8201836020820111156103fa57600080fd5b803590602001918460208302840111600160201b8311171561041b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561046a57600080fd5b82018360208201111561047c57600080fd5b803590602001918460208302840111600160201b8311171561049d57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610986945050505050565b6105fe600480360360408110156104f157600080fd5b810190602081018135600160201b81111561050b57600080fd5b82018360208201111561051d57600080fd5b803590602001918460208302840111600160201b8311171561053e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561058d57600080fd5b82018360208201111561059f57600080fd5b803590602001918460208302840111600160201b831117156105c057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610a71945050505050565b604080519115158252519081900360200190f35b60408051600480825260a0820190925260009160609190816020015b61063661101a565b81526020019060019003908161062e57905050905061065486610c76565b8160008151811061066157fe5b602002602001018190525061067e836001600160a01b0316610cd2565b8160018151811061068b57fe5b602002602001018190525061069f84610cd2565b816002815181106106ac57fe5b60209081029190910101526106ce6affffffffffffffffffffff198616610cd2565b816003815181106106db57fe5b60200260200101819052506106f76106f282610d2c565b610db4565b519695505050505050565b606060008351905060608551604051908082528060200260200182016040528015610737578160200160208202803883390190505b50905060005b828110156108dd57600086828151811061075357fe5b60200260200101519050878161ffff168151811061076d57fe5b602002602001015160146015811061078157fe5b1a60f81b6001600160f81b0319166107ce5785828151811061079f57fe5b6020026020010151838261ffff16815181106107b757fe5b6020026020010181815101915081815250506108d4565b828161ffff16815181106107de57fe5b602002602001015160001461083a576040805162461bcd60e51b815260206004820152601d60248201527f43616e277420696e636c756465204e465420746f6b656e207477696365000000604482015290519081900360640190fd5b85828151811061084657fe5b6020026020010151600014156108a3576040805162461bcd60e51b815260206004820152601f60248201527f4e465420746f6b656e206d7573742068617665206e6f6e2d7a65726f20696400604482015290519081900360640190fd5b8582815181106108af57fe5b6020026020010151838261ffff16815181106108c757fe5b6020026020010181815250505b5060010161073d565b5095945050505050565b600087878787878787604051602001808881526020018763ffffffff1663ffffffff1660e01b8152600401868152602001858152602001848152602001838152602001828051906020019060200280838360005b8381101561095357818101518382015260200161093b565b50505050905001975050505050505050604051602081830303815290604052805190602001209050979650505050505050565b600085858260200201518660016020020151868686604051602001808781526020018667ffffffffffffffff1667ffffffffffffffff1660c01b81526008018567ffffffffffffffff1667ffffffffffffffff1660c01b8152600801848152602001838051906020019060200280838360005b83811015610a115781810151838201526020016109f9565b50505050905001828051906020019060200280838360005b83811015610a41578181015183820152602001610a29565b50505050905001965050505050505060405160208183030381529060405280519060200120905095945050505050565b8151600090801580610a835750806001145b15610a92576001915050610c70565b60005b60018203811015610c23576000858281518110610aae57fe5b6020026020010151601460158110610ac257fe5b1a60f81b90506001600160f81b03198116610b2c57858281518110610ae357fe5b60200260200101516001600160581b031916868360010181518110610b0457fe5b60200260200101516001600160581b03191611610b275760009350505050610c70565b610c1a565b600160f81b6001600160f81b031982161415610c0e57858281518110610b4e57fe5b60200260200101516001600160581b031916868360010181518110610b6f57fe5b60200260200101516001600160581b0319161080610bfd5750858281518110610b9457fe5b60200260200101516001600160581b031916868360010181518110610bb557fe5b60200260200101516001600160581b031916148015610bfd5750848281518110610bdb57fe5b6020026020010151858360010181518110610bf257fe5b602002602001015111155b15610b275760009350505050610c70565b60009350505050610c70565b50600101610a95565b50600160f81b846001830381518110610c3857fe5b6020026020010151601460158110610c4c57fe5b1a60f81b6001600160f81b0319161115610c6a576000915050610c70565b60019150505b92915050565b610c7e61101a565b604080516060810182528381528151600080825260208281019094529192830191610cbf565b610cac61101a565b815260200190600190039081610ca45790505b508152600260209091015290505b919050565b610cda61101a565b604080516060810182528381528151600080825260208281019094529192830191610d1b565b610d0861101a565b815260200190600190039081610d005790505b508152600060209091015292915050565b610d3461101a565b610d3e8251610ea3565b610d8f576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b610dbc61103e565b6040820151600c60ff90911610610e0e576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff16610e3b576040518060200160405280610e328460000151610eaa565b90529050610ccd565b604082015160ff1660021415610e605750604080516020810190915281518152610ccd565b600360ff16826040015160ff1610158015610e8457506040820151600c60ff909116105b15610ea1576040518060200160405280610e328460200151610ece565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600882511115610f1e576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610f4b578160200160208202803883390190505b50805190915060005b81811015610fa757610f6461103e565b610f80868381518110610f7357fe5b6020026020010151610db4565b90508060000151848381518110610f9357fe5b602090810291909101015250600101610f54565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610ff0578181015183820152602001610fd8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72315820ecf8fa30c51cb3aa3c6396efee3048b7f8418b043bcd0a81f4220f2e7d56697f64736f6c634300050c0032"

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

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCaller) BeforeBalancesValid(opts *bind.CallOpts, _tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "beforeBalancesValid", _tokenTypes, _beforeBalances)
	return *ret0, err
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCallerSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
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

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
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

// ArbValueABI is the input ABI used to generate the binding from.
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidValueHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeValueHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValidValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

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
var ArbValueBin = "0x610d71610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c806353409fab1161006557806353409fab1461022157806389df40da146102475780638f34603614610308578063b2b9dc62146103ae57610092565b80631667b411146100975780631f3d4d4e146100c6578063264f384b146101ed578063364df27714610219575b600080fd5b6100b4600480360360208110156100ad57600080fd5b50356103df565b60408051918252519081900360200190f35b61016e600480360360408110156100dc57600080fd5b8101906020810181356401000000008111156100f757600080fd5b82018360208201111561010957600080fd5b8035906020019184600183028401116401000000008311171561012b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610405915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101b1578181015183820152602001610199565b50505050905090810190601f1680156101de5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100b46004803603606081101561020357600080fd5b5060ff813516906020810135906040013561049a565b6100b46104ec565b6100b46004803603604081101561023757600080fd5b5060ff813516906020013561055f565b6102ef6004803603604081101561025d57600080fd5b81019060208101813564010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506105a6915050565b6040805192835260208301919091528051918290030190f35b6100b46004803603602081101561031e57600080fd5b81019060208101813564010000000081111561033957600080fd5b82018360208201111561034b57600080fd5b8035906020019184600183028401116401000000008311171561036d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610631945050505050565b6103cb600480360360208110156103c457600080fd5b50356106b5565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b60006060600080610414610d06565b61041e87876106bc565b919450925090508215610478576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161048c888880840363ffffffff61081116565b945094505050509250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b83811015610538578181015183820152602001610520565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000806000806105b4610d06565b6105be87876106bc565b919450925090508215610618576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161062282610891565b51909890975095505050505050565b6000808061063d610d06565b6106488560006106bc565b9194509250905082156106a2576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b6106ab81610891565b5195945050505050565b6008101590565b6000806106c7610d06565b8451841061071c576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b6000849050600086828151811061072f57fe5b016020015160019092019160f81c9050600081610771576107508884610980565b9093509050600083610761836109a7565b9197509550935061080a92505050565b60ff821660021415610798576107878884610980565b909350905060008361076183610a01565b600360ff8316108015906107af5750600c60ff8316105b156107eb576002198201606060006107c8838c88610a5b565b9097509250905080866107da84610b16565b98509850985050505050505061080a565b8160ff166127100160006107ff60006109a7565b919750955093505050505b9250925092565b60608183018451101561082357600080fd5b60608215801561083e57604051915060208201604052610888565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561087757805183526020928301920161085f565b5050858452601f01601f1916604052505b50949350505050565b610899610d2a565b6040820151600c60ff909116106108eb576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff1661091857604051806020016040528061090f84600001516103df565b90529050610400565b604082015160ff166002141561093d5750604080516020810190915281518152610400565b600360ff16826040015160ff161015801561096157506040820151600c60ff909116105b1561097e57604051806020016040528061090f8460200151610b9e565bfe5b6000808281610995868363ffffffff610cea16565b60209290920196919550909350505050565b6109af610d06565b6040805160608101825283815281516000808252602082810190945291928301916109f0565b6109dd610d06565b8152602001906001900390816109d55790505b508152600060209091015292915050565b610a09610d06565b604080516060810182528381528151600080825260208281019094529192830191610a4a565b610a37610d06565b815260200190600190039081610a2f5790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610aa657816020015b610a93610d06565b815260200190600190039081610a8b5790505b50905060005b8960ff168160ff161015610b0057610ac489856106bc565b8451859060ff8616908110610ad557fe5b6020908102919091010152945092508215610af857509094509092509050610b0d565b600101610aac565b5060009550919350909150505b93509350939050565b610b1e610d06565b610b2882516106b5565b610b79576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b6000600882511115610bee576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610c1b578160200160208202803883390190505b50805190915060005b81811015610c7757610c34610d2a565b610c50868381518110610c4357fe5b6020026020010151610891565b90508060000151848381518110610c6357fe5b602090810291909101015250600101610c24565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610cc0578181015183820152602001610ca8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60008160200183511015610cfd57600080fd5b50016020015190565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a723158206a404b91a39afaa6b7b4b2ed1dc47b9a7da59b3ac11fb6c6bd29a2df197a166764736f6c634300050c0032"

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

// ArbitrumVMABI is the input ABI used to generate the binding from.
const ArbitrumVMABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[3]\",\"name\":\"fields\",\"type\":\"bytes32[3]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"bytes21[]\",\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastMessageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeManager\",\"outputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"bytes\",\"name\":\"_messageData\",\"type\":\"bytes\"},{\"internalType\":\"uint16[]\",\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_messageDestinations\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumVM.State\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"_fields\",\"type\":\"bytes32[4]\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_messageDataHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_messageDestinations\",\"type\":\"address[]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"enumVM.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"inChallenge\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbitrumVMFuncSigs maps the 4-byte function signature to its string representation.
var ArbitrumVMFuncSigs = map[string]string{
	"94af716b": "activateVM()",
	"023a96fe": "challengeManager()",
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"4526c5d9": "confirmDisputableAsserted(bytes32,bytes32,uint32,bytes21[],bytes,uint16[],uint256[],address[],bytes32)",
	"08dc89d7": "currentDeposit(address)",
	"aca0f372": "escrowRequired()",
	"6be00229": "exitAddress()",
	"1865c57d": "getState()",
	"d489113a": "globalInbox()",
	"2782e87e": "initiateChallenge(bytes32)",
	"8da5cb5b": "owner()",
	"cfa80707": "ownerShutdown()",
	"97e2e256": "pendingDisputableAssert(bytes32[4],uint32,uint64[2],bytes21[],bytes32[],uint16[],uint256[],address[])",
	"60675a87": "terminateAddress()",
	"3a768463": "vm()",
}

// ArbitrumVMBin is the compiled bytecode used for deploying new contracts.
var ArbitrumVMBin = "0x608060405234801561001057600080fd5b50604051620021e5380380620021e5833981810160405260e081101561003557600080fd5b50805160208201516040808401516060850151608086015160a087015160c090970151600180546001600160a01b038084166001600160a01b0319928316179283905560008054828d16931692909217825587517ff39723830000000000000000000000000000000000000000000000000000000081529751999a9899969895979496939492169263f397238392600480820193929182900301818387803b1580156100e057600080fd5b505af11580156100f4573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b15801561017b57600080fd5b505af415801561018f573d6000803e3d6000fd5b505050506040513d60208110156101a557600080fd5b50516004555050600680546001600160801b0319166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161763ffffffff60201b1916640100000000929093169190910291909117905550611fd480620002116000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806360675a871161009757806397e2e2561161006657806397e2e25614610546578063aca0f3721461085d578063cfa8070714610865578063d489113a1461086d576100f5565b806360675a87146105265780636be002291461052e5780638da5cb5b1461053657806394af716b1461053e576100f5565b806322c091bc116100d357806322c091bc146101825780632782e87e146101a45780633a768463146101c15780634526c5d91461025e576100f5565b8063023a96fe146100fa57806308dc89d71461011e5780631865c57d14610156575b600080fd5b610102610875565b604080516001600160a01b039092168252519081900360200190f35b6101446004803603602081101561013457600080fd5b50356001600160a01b0316610884565b60408051918252519081900360200190f35b61015e6108a3565b6040518082600381111561016e57fe5b60ff16815260200191505060405180910390f35b6101a26004803603608081101561019857600080fd5b50604081016108b3565b005b6101a2600480360360208110156101ba57600080fd5b5035610a06565b6101c9610c0a565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561023957fe5b60ff1681529115156020830152506040805191829003019a5098505050505050505050f35b6101a2600480360361012081101561027557600080fd5b81359160208101359163ffffffff6040830135169190810190608081016060820135600160201b8111156102a857600080fd5b8201836020820111156102ba57600080fd5b803590602001918460208302840111600160201b831117156102db57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561032a57600080fd5b82018360208201111561033c57600080fd5b803590602001918460018302840111600160201b8311171561035d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103af57600080fd5b8201836020820111156103c157600080fd5b803590602001918460208302840111600160201b831117156103e257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561043157600080fd5b82018360208201111561044357600080fd5b803590602001918460208302840111600160201b8311171561046457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104b357600080fd5b8201836020820111156104c557600080fd5b803590602001918460208302840111600160201b831117156104e657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250610c7b915050565b610102610ef7565b610102610f06565b610102610f15565b6101a2610f24565b6101a2600480360361018081101561055d57600080fd5b810190808060800190600480602002604051908101604052809291908260046020028082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b8111156105e457600080fd5b8201836020820111156105f657600080fd5b803590602001918460208302840111600160201b8311171561061757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561066657600080fd5b82018360208201111561067857600080fd5b803590602001918460208302840111600160201b8311171561069957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106e857600080fd5b8201836020820111156106fa57600080fd5b803590602001918460208302840111600160201b8311171561071b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561076a57600080fd5b82018360208201111561077c57600080fd5b803590602001918460208302840111600160201b8311171561079d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107ec57600080fd5b8201836020820111156107fe57600080fd5b803590602001918460208302840111600160201b8311171561081f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610fb7945050505050565b6101446116eb565b6101a26116fa565b61010261175a565b6000546001600160a01b031681565b6001600160a01b0381166000908152600860205260409020545b919050565b600754600160401b900460ff1690565b6000546001600160a01b031633146108fc5760405162461bcd60e51b815260040180806020018281038252602d815260200180611f1e602d913960400191505060405180910390fd5b600754600160481b900460ff166109445760405162461bcd60e51b8152600401808060200182810382526026815260200180611ef86026913960400191505060405180910390fd5b6007805469ff000000000000000000191690556109a96001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000205461176990919063ffffffff16565b82356001600160a01b031660009081526008602081815260408320939093556109e1928401356001600160801b03169185600161096c565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b03161115610a5f5760405162461bcd60e51b8152600401808060200182810382526027815260200180611ed16027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b158015610ad957600080fd5b505af4158015610aed573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b0316808252918101919091526007548451630823813560e21b815292909516975063208e04d496509194919363ffffffff16928892600490920191829187918190849084905b83811015610b85578181015183820152602001610b6d565b5050505090500184600260200280838360005b83811015610bb0578181015183820152602001610b98565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050600060405180830381600087803b158015610bef57600080fd5b505af1158015610c03573d6000803e3d6000fd5b5050505050565b6002546003546004546005546006546007546001600160a01b03909216916001600160801b0382169167ffffffffffffffff600160801b8204811692600160c01b909204169063ffffffff80821691600160201b81049091169060ff600160401b8204811691600160481b9004168b565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63924e7b3760028b8b8b8b8b8b8b8b8b6040518b63ffffffff1660e01b8152600401808b81526020018a81526020018981526020018863ffffffff1663ffffffff168152602001806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b83811015610d29578181015183820152602001610d11565b5050505090500186810385528b818151815260200191508051906020019080838360005b83811015610d65578181015183820152602001610d4d565b50505050905090810190601f168015610d925780820380516001836020036101000a031916815260200191505b5086810384528a5181528a51602091820191808d01910280838360005b83811015610dc7578181015183820152602001610daf565b50505050905001868103835289818151815260200191508051906020019060200280838360005b83811015610e06578181015183820152602001610dee565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015610e45578181015183820152602001610e2d565b505050509050019f5050505050505050505050505050505060006040518083038186803b158015610e7557600080fd5b505af4158015610e89573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610ec4935091506001600160801b031663ffffffff61176916565b6005546001600160a01b0316600090815260086020526040902055610eec86868686866117ca565b505050505050505050565b600a546001600160a01b031681565b6009546001600160a01b031681565b600b546001600160a01b031681565b600b546001600160a01b03163314610f7c576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff166003811115610f9657fe5b1415610fb5576007805468ff00000000000000001916600160401b1790555b565b336000908152600860205260409020546006546001600160801b031611156110105760405162461bcd60e51b8152600401808060200182810382526031815260200180611f4b6031913960400191505060405180910390fd5b60065433600090815260086020908152604080832080546001600160801b039095169094039093559151630f89fbff60e01b8152606060048201818152895160648401528951919473__$9836fa7140e5a33041d4b827682e675a30$__94630f89fbff948c948b948b9490938493602481019360448201936084909201928a83019291909102908190849084905b838110156110b657818101518382015260200161109e565b50505050905001848103835286818151815260200191508051906020019060200280838360005b838110156110f55781810151838201526020016110dd565b50505050905001848103825285818151815260200191508051906020019060200280838360005b8381101561113457818101518382015260200161111c565b50505050905001965050505050505060006040518083038186803b15801561115b57600080fd5b505af415801561116f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561119857600080fd5b8101908080516040519392919084600160201b8211156111b757600080fd5b9083019060208201858111156111cc57600080fd5b82518660208202830111600160201b821117156111e857600080fd5b82525081516020918201928201910280838360005b838110156112155781810151838201526020016111fd565b50505050905001604052505050905073__$9836fa7140e5a33041d4b827682e675a30$__63af17d92287836040518363ffffffff1660e01b8152600401808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561129357818101518382015260200161127b565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156112d25781810151838201526020016112ba565b5050505090500194505050505060206040518083038186803b1580156112f757600080fd5b505af415801561130b573d6000803e3d6000fd5b505050506040513d602081101561132157600080fd5b505161135e5760405162461bcd60e51b8152600401808060200182810382526024815260200180611f7c6024913960400191505060405180910390fd5b60015460405163565b19db60e11b815230600482018181526060602484019081528a5160648501528a516001600160a01b039095169463acb633b6948c9388939092909160448101916084909101906020808801910280838360005b838110156113d25781810151838201526020016113ba565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156114115781810151838201526020016113f9565b505050509050019550505050505060206040518083038186803b15801561143757600080fd5b505afa15801561144b573d6000803e3d6000fd5b505050506040513d602081101561146157600080fd5b50516114b4576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63c97c8eec60028b8b8b8b8b8b8b8b6040518a63ffffffff1660e01b8152600401808a815260200189600460200280838360005b838110156115155781810151838201526020016114fd565b5050505063ffffffff8b1692019182525060200187604080838360005b8381101561154a578181015183820152602001611532565b50505050905001806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b8381101561159d578181015183820152602001611585565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b838110156115dc5781810151838201526020016115c4565b50505050905001868103845289818151815260200191508051906020019060200280838360005b8381101561161b578181015183820152602001611603565b50505050905001868103835288818151815260200191508051906020019060200280838360005b8381101561165a578181015183820152602001611642565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611699578181015183820152602001611681565b505050509050019e50505050505050505050505050505060006040518083038186803b1580156116c857600080fd5b505af41580156116dc573d6000803e3d6000fd5b50505050505050505050505050565b6006546001600160801b031690565b600b546001600160a01b03163314611752576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b610fb5611af1565b6001546001600160a01b031681565b6000828201838110156117c3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b15801561181057600080fd5b505af1158015611824573d6000803e3d6000fd5b505050506040513d602081101561183a57600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b15801561188b57600080fd5b505af415801561189f573d6000803e3d6000fd5b505050506040513d60208110156118b557600080fd5b50518114611900576118fc60405180606001604052806118d56001611aff565b81526020016118e76002800154611b59565b81526020016118f584611b59565b9052611bb3565b6004555b600160009054906101000a90046001600160a01b03166001600160a01b031663ec22a76787878787876040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015611988578181015183820152602001611970565b5050505090500186810385528a818151815260200191508051906020019080838360005b838110156119c45781810151838201526020016119ac565b50505050905090810190601f1680156119f15780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015611a26578181015183820152602001611a0e565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015611a65578181015183820152602001611a4d565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611aa4578181015183820152602001611a8c565b505050509050019a5050505050505050505050600060405180830381600087803b158015611ad157600080fd5b505af1158015611ae5573d6000803e3d6000fd5b50505050505050505050565b600b546001600160a01b0316ff5b611b07611e9a565b604080516060810182528381528151600080825260208281019094529192830191611b48565b611b35611e9a565b815260200190600190039081611b2d5790505b508152600060209091015292915050565b611b61611e9a565b604080516060810182528381528151600080825260208281019094529192830191611ba2565b611b8f611e9a565b815260200190600190039081611b875790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b611bd7611e9a565b815260200190600190039081611bcf575050805190915060005b81811015611c2957848160038110611c0557fe5b6020020151838281518110611c1657fe5b6020908102919091010152600101611bf1565b50611c3382611c3b565b949350505050565b6000600882511115611c8b576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611cb8578160200160208202803883390190505b50805190915060005b81811015611d1457611cd1611ebe565b611ced868381518110611ce057fe5b6020026020010151611d87565b90508060000151848381518110611d0057fe5b602090810291909101015250600101611cc1565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611d5d578181015183820152602001611d45565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b611d8f611ebe565b6040820151600c60ff90911610611de1576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff16611e0e576040518060200160405280611e058460000151611e76565b9052905061089e565b604082015160ff1660021415611e33575060408051602081019091528151815261089e565b600360ff16826040015160ff1610158015611e5757506040820151600c60ff909116105b15611e74576040518060200160405280611e058460200151611c3b565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274546f6b656e207479706573206d7573742062652076616c696420616e6420736f72746564a265627a7a72315820ee588d9159887df6a24bcffc665874487e0f27fb78c0b9a47290ad778cffcc5664736f6c634300050c0032"

// DeployArbitrumVM deploys a new Ethereum contract, binding an instance of ArbitrumVM to it.
func DeployArbitrumVM(auth *bind.TransactOpts, backend bind.ContractBackend, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeManagerAddress common.Address, _globalInboxAddress common.Address) (common.Address, *types.Transaction, *ArbitrumVM, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrumVMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ArbitrumVMBin = strings.Replace(ArbitrumVMBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	ArbitrumVMBin = strings.Replace(ArbitrumVMBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbitrumVMBin = strings.Replace(ArbitrumVMBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbitrumVMBin), backend, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeManagerAddress, _globalInboxAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumVM{ArbitrumVMCaller: ArbitrumVMCaller{contract: contract}, ArbitrumVMTransactor: ArbitrumVMTransactor{contract: contract}, ArbitrumVMFilterer: ArbitrumVMFilterer{contract: contract}}, nil
}

// ArbitrumVM is an auto generated Go binding around an Ethereum contract.
type ArbitrumVM struct {
	ArbitrumVMCaller     // Read-only binding to the contract
	ArbitrumVMTransactor // Write-only binding to the contract
	ArbitrumVMFilterer   // Log filterer for contract events
}

// ArbitrumVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbitrumVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrumVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbitrumVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrumVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbitrumVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrumVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbitrumVMSession struct {
	Contract     *ArbitrumVM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbitrumVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbitrumVMCallerSession struct {
	Contract *ArbitrumVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbitrumVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbitrumVMTransactorSession struct {
	Contract     *ArbitrumVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbitrumVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbitrumVMRaw struct {
	Contract *ArbitrumVM // Generic contract binding to access the raw methods on
}

// ArbitrumVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbitrumVMCallerRaw struct {
	Contract *ArbitrumVMCaller // Generic read-only contract binding to access the raw methods on
}

// ArbitrumVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbitrumVMTransactorRaw struct {
	Contract *ArbitrumVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbitrumVM creates a new instance of ArbitrumVM, bound to a specific deployed contract.
func NewArbitrumVM(address common.Address, backend bind.ContractBackend) (*ArbitrumVM, error) {
	contract, err := bindArbitrumVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumVM{ArbitrumVMCaller: ArbitrumVMCaller{contract: contract}, ArbitrumVMTransactor: ArbitrumVMTransactor{contract: contract}, ArbitrumVMFilterer: ArbitrumVMFilterer{contract: contract}}, nil
}

// NewArbitrumVMCaller creates a new read-only instance of ArbitrumVM, bound to a specific deployed contract.
func NewArbitrumVMCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumVMCaller, error) {
	contract, err := bindArbitrumVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMCaller{contract: contract}, nil
}

// NewArbitrumVMTransactor creates a new write-only instance of ArbitrumVM, bound to a specific deployed contract.
func NewArbitrumVMTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumVMTransactor, error) {
	contract, err := bindArbitrumVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMTransactor{contract: contract}, nil
}

// NewArbitrumVMFilterer creates a new log filterer instance of ArbitrumVM, bound to a specific deployed contract.
func NewArbitrumVMFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumVMFilterer, error) {
	contract, err := bindArbitrumVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMFilterer{contract: contract}, nil
}

// bindArbitrumVM binds a generic wrapper to an already deployed contract.
func bindArbitrumVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrumVMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbitrumVM *ArbitrumVMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbitrumVM.Contract.ArbitrumVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbitrumVM *ArbitrumVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.ArbitrumVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbitrumVM *ArbitrumVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.ArbitrumVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbitrumVM *ArbitrumVMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbitrumVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbitrumVM *ArbitrumVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbitrumVM *ArbitrumVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.contract.Transact(opts, method, params...)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCaller) ChallengeManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "challengeManager")
	return *ret0, err
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbitrumVM *ArbitrumVMSession) ChallengeManager() (common.Address, error) {
	return _ArbitrumVM.Contract.ChallengeManager(&_ArbitrumVM.CallOpts)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCallerSession) ChallengeManager() (common.Address, error) {
	return _ArbitrumVM.Contract.ChallengeManager(&_ArbitrumVM.CallOpts)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbitrumVM *ArbitrumVMCaller) CurrentDeposit(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "currentDeposit", validator)
	return *ret0, err
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbitrumVM *ArbitrumVMSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _ArbitrumVM.Contract.CurrentDeposit(&_ArbitrumVM.CallOpts, validator)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbitrumVM *ArbitrumVMCallerSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _ArbitrumVM.Contract.CurrentDeposit(&_ArbitrumVM.CallOpts, validator)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbitrumVM *ArbitrumVMCaller) EscrowRequired(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "escrowRequired")
	return *ret0, err
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbitrumVM *ArbitrumVMSession) EscrowRequired() (*big.Int, error) {
	return _ArbitrumVM.Contract.EscrowRequired(&_ArbitrumVM.CallOpts)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbitrumVM *ArbitrumVMCallerSession) EscrowRequired() (*big.Int, error) {
	return _ArbitrumVM.Contract.EscrowRequired(&_ArbitrumVM.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCaller) ExitAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "exitAddress")
	return *ret0, err
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbitrumVM *ArbitrumVMSession) ExitAddress() (common.Address, error) {
	return _ArbitrumVM.Contract.ExitAddress(&_ArbitrumVM.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCallerSession) ExitAddress() (common.Address, error) {
	return _ArbitrumVM.Contract.ExitAddress(&_ArbitrumVM.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbitrumVM *ArbitrumVMCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "getState")
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbitrumVM *ArbitrumVMSession) GetState() (uint8, error) {
	return _ArbitrumVM.Contract.GetState(&_ArbitrumVM.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbitrumVM *ArbitrumVMCallerSession) GetState() (uint8, error) {
	return _ArbitrumVM.Contract.GetState(&_ArbitrumVM.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCaller) GlobalInbox(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "globalInbox")
	return *ret0, err
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbitrumVM *ArbitrumVMSession) GlobalInbox() (common.Address, error) {
	return _ArbitrumVM.Contract.GlobalInbox(&_ArbitrumVM.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCallerSession) GlobalInbox() (common.Address, error) {
	return _ArbitrumVM.Contract.GlobalInbox(&_ArbitrumVM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbitrumVM *ArbitrumVMSession) Owner() (common.Address, error) {
	return _ArbitrumVM.Contract.Owner(&_ArbitrumVM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCallerSession) Owner() (common.Address, error) {
	return _ArbitrumVM.Contract.Owner(&_ArbitrumVM.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCaller) TerminateAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "terminateAddress")
	return *ret0, err
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbitrumVM *ArbitrumVMSession) TerminateAddress() (common.Address, error) {
	return _ArbitrumVM.Contract.TerminateAddress(&_ArbitrumVM.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCallerSession) TerminateAddress() (common.Address, error) {
	return _ArbitrumVM.Contract.TerminateAddress(&_ArbitrumVM.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbitrumVM *ArbitrumVMCaller) Vm(opts *bind.CallOpts) (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	State             uint8
	InChallenge       bool
}, error) {
	ret := new(struct {
		MachineHash       [32]byte
		PendingHash       [32]byte
		Inbox             [32]byte
		Asserter          common.Address
		EscrowRequired    *big.Int
		Deadline          uint64
		SequenceNum       uint64
		GracePeriod       uint32
		MaxExecutionSteps uint32
		State             uint8
		InChallenge       bool
	})
	out := ret
	err := _ArbitrumVM.contract.Call(opts, out, "vm")
	return *ret, err
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbitrumVM *ArbitrumVMSession) Vm() (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	State             uint8
	InChallenge       bool
}, error) {
	return _ArbitrumVM.Contract.Vm(&_ArbitrumVM.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbitrumVM *ArbitrumVMCallerSession) Vm() (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	State             uint8
	InChallenge       bool
}, error) {
	return _ArbitrumVM.Contract.Vm(&_ArbitrumVM.CallOpts)
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbitrumVM *ArbitrumVMTransactor) ActivateVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumVM.contract.Transact(opts, "activateVM")
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbitrumVM *ArbitrumVMSession) ActivateVM() (*types.Transaction, error) {
	return _ArbitrumVM.Contract.ActivateVM(&_ArbitrumVM.TransactOpts)
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbitrumVM *ArbitrumVMTransactorSession) ActivateVM() (*types.Transaction, error) {
	return _ArbitrumVM.Contract.ActivateVM(&_ArbitrumVM.TransactOpts)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbitrumVM *ArbitrumVMTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbitrumVM.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbitrumVM *ArbitrumVMSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.CompleteChallenge(&_ArbitrumVM.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbitrumVM *ArbitrumVMTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.CompleteChallenge(&_ArbitrumVM.TransactOpts, _players, _rewards)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_ArbitrumVM *ArbitrumVMTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.contract.Transact(opts, "confirmDisputableAsserted", _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_ArbitrumVM *ArbitrumVMSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.ConfirmDisputableAsserted(&_ArbitrumVM.TransactOpts, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_ArbitrumVM *ArbitrumVMTransactorSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.ConfirmDisputableAsserted(&_ArbitrumVM.TransactOpts, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbitrumVM *ArbitrumVMTransactor) InitiateChallenge(opts *bind.TransactOpts, _assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.contract.Transact(opts, "initiateChallenge", _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbitrumVM *ArbitrumVMSession) InitiateChallenge(_assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.InitiateChallenge(&_ArbitrumVM.TransactOpts, _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbitrumVM *ArbitrumVMTransactorSession) InitiateChallenge(_assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.InitiateChallenge(&_ArbitrumVM.TransactOpts, _assertPreHash)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbitrumVM *ArbitrumVMTransactor) OwnerShutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumVM.contract.Transact(opts, "ownerShutdown")
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbitrumVM *ArbitrumVMSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbitrumVM.Contract.OwnerShutdown(&_ArbitrumVM.TransactOpts)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbitrumVM *ArbitrumVMTransactorSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbitrumVM.Contract.OwnerShutdown(&_ArbitrumVM.TransactOpts)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_ArbitrumVM *ArbitrumVMTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _ArbitrumVM.contract.Transact(opts, "pendingDisputableAssert", _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_ArbitrumVM *ArbitrumVMSession) PendingDisputableAssert(_fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.PendingDisputableAssert(&_ArbitrumVM.TransactOpts, _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_ArbitrumVM *ArbitrumVMTransactorSession) PendingDisputableAssert(_fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.PendingDisputableAssert(&_ArbitrumVM.TransactOpts, _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// ArbitrumVMConfirmedDisputableAssertionIterator is returned from FilterConfirmedDisputableAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedDisputableAssertion events raised by the ArbitrumVM contract.
type ArbitrumVMConfirmedDisputableAssertionIterator struct {
	Event *ArbitrumVMConfirmedDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *ArbitrumVMConfirmedDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumVMConfirmedDisputableAssertion)
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
		it.Event = new(ArbitrumVMConfirmedDisputableAssertion)
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
func (it *ArbitrumVMConfirmedDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrumVMConfirmedDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrumVMConfirmedDisputableAssertion represents a ConfirmedDisputableAssertion event raised by the ArbitrumVM contract.
type ArbitrumVMConfirmedDisputableAssertion struct {
	NewState    [32]byte
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbitrumVM *ArbitrumVMFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts) (*ArbitrumVMConfirmedDisputableAssertionIterator, error) {

	logs, sub, err := _ArbitrumVM.contract.FilterLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMConfirmedDisputableAssertionIterator{contract: _ArbitrumVM.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbitrumVM *ArbitrumVMFilterer) WatchConfirmedDisputableAssertion(opts *bind.WatchOpts, sink chan<- *ArbitrumVMConfirmedDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbitrumVM.contract.WatchLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrumVMConfirmedDisputableAssertion)
				if err := _ArbitrumVM.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
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

// ParseConfirmedDisputableAssertion is a log parse operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbitrumVM *ArbitrumVMFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*ArbitrumVMConfirmedDisputableAssertion, error) {
	event := new(ArbitrumVMConfirmedDisputableAssertion)
	if err := _ArbitrumVM.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbitrumVMInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ArbitrumVM contract.
type ArbitrumVMInitiatedChallengeIterator struct {
	Event *ArbitrumVMInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ArbitrumVMInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumVMInitiatedChallenge)
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
		it.Event = new(ArbitrumVMInitiatedChallenge)
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
func (it *ArbitrumVMInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrumVMInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrumVMInitiatedChallenge represents a InitiatedChallenge event raised by the ArbitrumVM contract.
type ArbitrumVMInitiatedChallenge struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_ArbitrumVM *ArbitrumVMFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ArbitrumVMInitiatedChallengeIterator, error) {

	logs, sub, err := _ArbitrumVM.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMInitiatedChallengeIterator{contract: _ArbitrumVM.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_ArbitrumVM *ArbitrumVMFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ArbitrumVMInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _ArbitrumVM.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrumVMInitiatedChallenge)
				if err := _ArbitrumVM.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_ArbitrumVM *ArbitrumVMFilterer) ParseInitiatedChallenge(log types.Log) (*ArbitrumVMInitiatedChallenge, error) {
	event := new(ArbitrumVMInitiatedChallenge)
	if err := _ArbitrumVM.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbitrumVMPendingDisputableAssertionIterator is returned from FilterPendingDisputableAssertion and is used to iterate over the raw logs and unpacked data for PendingDisputableAssertion events raised by the ArbitrumVM contract.
type ArbitrumVMPendingDisputableAssertionIterator struct {
	Event *ArbitrumVMPendingDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *ArbitrumVMPendingDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumVMPendingDisputableAssertion)
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
		it.Event = new(ArbitrumVMPendingDisputableAssertion)
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
func (it *ArbitrumVMPendingDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrumVMPendingDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrumVMPendingDisputableAssertion represents a PendingDisputableAssertion event raised by the ArbitrumVM contract.
type ArbitrumVMPendingDisputableAssertion struct {
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

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_ArbitrumVM *ArbitrumVMFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*ArbitrumVMPendingDisputableAssertionIterator, error) {

	logs, sub, err := _ArbitrumVM.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMPendingDisputableAssertionIterator{contract: _ArbitrumVM.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_ArbitrumVM *ArbitrumVMFilterer) WatchPendingDisputableAssertion(opts *bind.WatchOpts, sink chan<- *ArbitrumVMPendingDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbitrumVM.contract.WatchLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrumVMPendingDisputableAssertion)
				if err := _ArbitrumVM.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_ArbitrumVM *ArbitrumVMFilterer) ParsePendingDisputableAssertion(log types.Log) (*ArbitrumVMPendingDisputableAssertion, error) {
	event := new(ArbitrumVMPendingDisputableAssertion)
	if err := _ArbitrumVM.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a9fb49964eecf6140d019ab385582e49921deabebb21f44e37efb5be3dbaaf2b64736f6c634300050c0032"

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

// ChainLauncherABI is the input ABI used to generate the binding from.
const ChainLauncherABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeManagerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"name\":\"ChainCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"launchChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChainLauncherFuncSigs maps the 4-byte function signature to its string representation.
var ChainLauncherFuncSigs = map[string]string{
	"e2b491e3": "launchChain(bytes32,uint32,uint32,uint128,address)",
}

// ChainLauncherBin is the compiled bytecode used for deploying new contracts.
var ChainLauncherBin = "0x608060405234801561001057600080fd5b506040516125143803806125148339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b0319918216179091556001805493909216921691909117905561249a8061007a6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e2b491e314610030575b600080fd5b610082600480360360a081101561004657600080fd5b50803590602081013563ffffffff9081169160408101359091169060608101356001600160801b031690608001356001600160a01b0316610084565b005b600154600080546040519192889288928892889288926001600160a01b039283169216906100b190610164565b96875263ffffffff9586166020880152939094166040808701919091526001600160801b0390921660608601526001600160a01b03908116608086015292831660a0850152911660c0830152519081900360e001906000f08015801561011b573d6000803e3d6000fd5b50604080516001600160a01b038316815290519192507fa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f919081900360200190a1505050505050565b6122f4806101728339019056fe608060405234801561001057600080fd5b50604051620022f4380380620022f4833981810160405260e081101561003557600080fd5b50805160208201516040808401516060850151608086015160a087015160c090970151600180546001600160a01b038084166001600160a01b0319928316179283905560008054828d16931692909217825587517ff39723830000000000000000000000000000000000000000000000000000000081529751999a989996989597949693948b948b948b948b948b948b948b949092169263f39723839260048084019382900301818387803b1580156100ed57600080fd5b505af1158015610101573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b15801561018857600080fd5b505af415801561019c573d6000803e3d6000fd5b505050506040513d60208110156101b257600080fd5b50516004555050600680546001600160801b0319166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161763ffffffff60201b19166401000000009290931691909102919091178082556001925060ff60401b191668010000000000000000830217905550505050505050506120b480620002406000396000f3fe6080604052600436106100f35760003560e01c806360675a871161008a57806397e2e2561161005957806397e2e256146105db578063aca0f372146108ff578063cfa8070714610914578063d489113a14610929576100f3565b806360675a87146105875780636be002291461059c5780638da5cb5b146105b157806394af716b146105c6576100f3565b806322c091bc116100c657806322c091bc146101b15780632782e87e146101de5780633a768463146102085780634526c5d9146102b2576100f3565b8063023a96fe146100f857806305b050de1461012957806308dc89d7146101335780631865c57d14610178575b600080fd5b34801561010457600080fd5b5061010d61093e565b604080516001600160a01b039092168252519081900360200190f35b61013161094d565b005b34801561013f57600080fd5b506101666004803603602081101561015657600080fd5b50356001600160a01b0316610964565b60408051918252519081900360200190f35b34801561018457600080fd5b5061018d610983565b6040518082600381111561019d57fe5b60ff16815260200191505060405180910390f35b3480156101bd57600080fd5b50610131600480360360808110156101d457600080fd5b5060408101610993565b3480156101ea57600080fd5b506101316004803603602081101561020157600080fd5b5035610ae6565b34801561021457600080fd5b5061021d610cea565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561028d57fe5b60ff1681529115156020830152506040805191829003019a5098505050505050505050f35b3480156102be57600080fd5b5061013160048036036101208110156102d657600080fd5b81359160208101359163ffffffff6040830135169190810190608081016060820135600160201b81111561030957600080fd5b82018360208201111561031b57600080fd5b803590602001918460208302840111600160201b8311171561033c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561038b57600080fd5b82018360208201111561039d57600080fd5b803590602001918460018302840111600160201b831117156103be57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561041057600080fd5b82018360208201111561042257600080fd5b803590602001918460208302840111600160201b8311171561044357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561049257600080fd5b8201836020820111156104a457600080fd5b803590602001918460208302840111600160201b831117156104c557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561051457600080fd5b82018360208201111561052657600080fd5b803590602001918460208302840111600160201b8311171561054757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250610d5b915050565b34801561059357600080fd5b5061010d610fd7565b3480156105a857600080fd5b5061010d610fe6565b3480156105bd57600080fd5b5061010d610ff5565b3480156105d257600080fd5b50610131611004565b3480156105e757600080fd5b5061013160048036036101808110156105ff57600080fd5b810190808060800190600480602002604051908101604052809291908260046020028082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561068657600080fd5b82018360208201111561069857600080fd5b803590602001918460208302840111600160201b831117156106b957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561070857600080fd5b82018360208201111561071a57600080fd5b803590602001918460208302840111600160201b8311171561073b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561078a57600080fd5b82018360208201111561079c57600080fd5b803590602001918460208302840111600160201b831117156107bd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561080c57600080fd5b82018360208201111561081e57600080fd5b803590602001918460208302840111600160201b8311171561083f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561088e57600080fd5b8201836020820111156108a057600080fd5b803590602001918460208302840111600160201b831117156108c157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611097945050505050565b34801561090b57600080fd5b506101666117cb565b34801561092057600080fd5b506101316117da565b34801561093557600080fd5b5061010d61183a565b6000546001600160a01b031681565b336000908152600860205260409020805434019055565b6001600160a01b0381166000908152600860205260409020545b919050565b600754600160401b900460ff1690565b6000546001600160a01b031633146109dc5760405162461bcd60e51b815260040180806020018281038252602d815260200180611ffe602d913960400191505060405180910390fd5b600754600160481b900460ff16610a245760405162461bcd60e51b8152600401808060200182810382526026815260200180611fd86026913960400191505060405180910390fd5b6007805469ff00000000000000000019169055610a896001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000205461184990919063ffffffff16565b82356001600160a01b03166000908152600860208181526040832093909355610ac1928401356001600160801b031691856001610a4c565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b03161115610b3f5760405162461bcd60e51b8152600401808060200182810382526027815260200180611fb16027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b158015610bb957600080fd5b505af4158015610bcd573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b0316808252918101919091526007548451630823813560e21b815292909516975063208e04d496509194919363ffffffff16928892600490920191829187918190849084905b83811015610c65578181015183820152602001610c4d565b5050505090500184600260200280838360005b83811015610c90578181015183820152602001610c78565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050600060405180830381600087803b158015610ccf57600080fd5b505af1158015610ce3573d6000803e3d6000fd5b5050505050565b6002546003546004546005546006546007546001600160a01b03909216916001600160801b0382169167ffffffffffffffff600160801b8204811692600160c01b909204169063ffffffff80821691600160201b81049091169060ff600160401b8204811691600160481b9004168b565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63924e7b3760028b8b8b8b8b8b8b8b8b6040518b63ffffffff1660e01b8152600401808b81526020018a81526020018981526020018863ffffffff1663ffffffff168152602001806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b83811015610e09578181015183820152602001610df1565b5050505090500186810385528b818151815260200191508051906020019080838360005b83811015610e45578181015183820152602001610e2d565b50505050905090810190601f168015610e725780820380516001836020036101000a031916815260200191505b5086810384528a5181528a51602091820191808d01910280838360005b83811015610ea7578181015183820152602001610e8f565b50505050905001868103835289818151815260200191508051906020019060200280838360005b83811015610ee6578181015183820152602001610ece565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015610f25578181015183820152602001610f0d565b505050509050019f5050505050505050505050505050505060006040518083038186803b158015610f5557600080fd5b505af4158015610f69573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610fa4935091506001600160801b031663ffffffff61184916565b6005546001600160a01b0316600090815260086020526040902055610fcc86868686866118aa565b505050505050505050565b600a546001600160a01b031681565b6009546001600160a01b031681565b600b546001600160a01b031681565b600b546001600160a01b0316331461105c576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff16600381111561107657fe5b1415611095576007805468ff00000000000000001916600160401b1790555b565b336000908152600860205260409020546006546001600160801b031611156110f05760405162461bcd60e51b815260040180806020018281038252603181526020018061202b6031913960400191505060405180910390fd5b60065433600090815260086020908152604080832080546001600160801b039095169094039093559151630f89fbff60e01b8152606060048201818152895160648401528951919473__$9836fa7140e5a33041d4b827682e675a30$__94630f89fbff948c948b948b9490938493602481019360448201936084909201928a83019291909102908190849084905b8381101561119657818101518382015260200161117e565b50505050905001848103835286818151815260200191508051906020019060200280838360005b838110156111d55781810151838201526020016111bd565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156112145781810151838201526020016111fc565b50505050905001965050505050505060006040518083038186803b15801561123b57600080fd5b505af415801561124f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561127857600080fd5b8101908080516040519392919084600160201b82111561129757600080fd5b9083019060208201858111156112ac57600080fd5b82518660208202830111600160201b821117156112c857600080fd5b82525081516020918201928201910280838360005b838110156112f55781810151838201526020016112dd565b50505050905001604052505050905073__$9836fa7140e5a33041d4b827682e675a30$__63af17d92287836040518363ffffffff1660e01b8152600401808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561137357818101518382015260200161135b565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156113b257818101518382015260200161139a565b5050505090500194505050505060206040518083038186803b1580156113d757600080fd5b505af41580156113eb573d6000803e3d6000fd5b505050506040513d602081101561140157600080fd5b505161143e5760405162461bcd60e51b815260040180806020018281038252602481526020018061205c6024913960400191505060405180910390fd5b60015460405163565b19db60e11b815230600482018181526060602484019081528a5160648501528a516001600160a01b039095169463acb633b6948c9388939092909160448101916084909101906020808801910280838360005b838110156114b257818101518382015260200161149a565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156114f15781810151838201526020016114d9565b505050509050019550505050505060206040518083038186803b15801561151757600080fd5b505afa15801561152b573d6000803e3d6000fd5b505050506040513d602081101561154157600080fd5b5051611594576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63c97c8eec60028b8b8b8b8b8b8b8b6040518a63ffffffff1660e01b8152600401808a815260200189600460200280838360005b838110156115f55781810151838201526020016115dd565b5050505063ffffffff8b1692019182525060200187604080838360005b8381101561162a578181015183820152602001611612565b50505050905001806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b8381101561167d578181015183820152602001611665565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b838110156116bc5781810151838201526020016116a4565b50505050905001868103845289818151815260200191508051906020019060200280838360005b838110156116fb5781810151838201526020016116e3565b50505050905001868103835288818151815260200191508051906020019060200280838360005b8381101561173a578181015183820152602001611722565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611779578181015183820152602001611761565b505050509050019e50505050505050505050505050505060006040518083038186803b1580156117a857600080fd5b505af41580156117bc573d6000803e3d6000fd5b50505050505050505050505050565b6006546001600160801b031690565b600b546001600160a01b03163314611832576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b611095611bd1565b6001546001600160a01b031681565b6000828201838110156118a3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b1580156118f057600080fd5b505af1158015611904573d6000803e3d6000fd5b505050506040513d602081101561191a57600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b15801561196b57600080fd5b505af415801561197f573d6000803e3d6000fd5b505050506040513d602081101561199557600080fd5b505181146119e0576119dc60405180606001604052806119b56001611bdf565b81526020016119c76002800154611c39565b81526020016119d584611c39565b9052611c93565b6004555b600160009054906101000a90046001600160a01b03166001600160a01b031663ec22a76787878787876040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015611a68578181015183820152602001611a50565b5050505090500186810385528a818151815260200191508051906020019080838360005b83811015611aa4578181015183820152602001611a8c565b50505050905090810190601f168015611ad15780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015611b06578181015183820152602001611aee565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015611b45578181015183820152602001611b2d565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611b84578181015183820152602001611b6c565b505050509050019a5050505050505050505050600060405180830381600087803b158015611bb157600080fd5b505af1158015611bc5573d6000803e3d6000fd5b50505050505050505050565b600b546001600160a01b0316ff5b611be7611f7a565b604080516060810182528381528151600080825260208281019094529192830191611c28565b611c15611f7a565b815260200190600190039081611c0d5790505b508152600060209091015292915050565b611c41611f7a565b604080516060810182528381528151600080825260208281019094529192830191611c82565b611c6f611f7a565b815260200190600190039081611c675790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b611cb7611f7a565b815260200190600190039081611caf575050805190915060005b81811015611d0957848160038110611ce557fe5b6020020151838281518110611cf657fe5b6020908102919091010152600101611cd1565b50611d1382611d1b565b949350505050565b6000600882511115611d6b576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611d98578160200160208202803883390190505b50805190915060005b81811015611df457611db1611f9e565b611dcd868381518110611dc057fe5b6020026020010151611e67565b90508060000151848381518110611de057fe5b602090810291909101015250600101611da1565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611e3d578181015183820152602001611e25565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b611e6f611f9e565b6040820151600c60ff90911610611ec1576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff16611eee576040518060200160405280611ee58460000151611f56565b9052905061097e565b604082015160ff1660021415611f13575060408051602081019091528151815261097e565b600360ff16826040015160ff1610158015611f3757506040820151600c60ff909116105b15611f54576040518060200160405280611ee58460200151611d1b565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274546f6b656e207479706573206d7573742062652076616c696420616e6420736f72746564a265627a7a7231582023f0161eb695a83cbe5a70f160e46fba3851d099415def78a48000d8a2d77ff064736f6c634300050c0032a265627a7a72315820b502a8a513cb57ab5d6c7a35893582f1e8f5a4389988608da9ec73809325b1c164736f6c634300050c0032"

// DeployChainLauncher deploys a new Ethereum contract, binding an instance of ChainLauncher to it.
func DeployChainLauncher(auth *bind.TransactOpts, backend bind.ContractBackend, _globalInboxAddress common.Address, _challengeManagerAddress common.Address) (common.Address, *types.Transaction, *ChainLauncher, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainLauncherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ChainLauncherBin = strings.Replace(ChainLauncherBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	ChainLauncherBin = strings.Replace(ChainLauncherBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ChainLauncherBin = strings.Replace(ChainLauncherBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChainLauncherBin), backend, _globalInboxAddress, _challengeManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChainLauncher{ChainLauncherCaller: ChainLauncherCaller{contract: contract}, ChainLauncherTransactor: ChainLauncherTransactor{contract: contract}, ChainLauncherFilterer: ChainLauncherFilterer{contract: contract}}, nil
}

// ChainLauncher is an auto generated Go binding around an Ethereum contract.
type ChainLauncher struct {
	ChainLauncherCaller     // Read-only binding to the contract
	ChainLauncherTransactor // Write-only binding to the contract
	ChainLauncherFilterer   // Log filterer for contract events
}

// ChainLauncherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainLauncherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainLauncherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainLauncherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainLauncherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainLauncherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainLauncherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainLauncherSession struct {
	Contract     *ChainLauncher    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainLauncherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainLauncherCallerSession struct {
	Contract *ChainLauncherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ChainLauncherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainLauncherTransactorSession struct {
	Contract     *ChainLauncherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ChainLauncherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainLauncherRaw struct {
	Contract *ChainLauncher // Generic contract binding to access the raw methods on
}

// ChainLauncherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainLauncherCallerRaw struct {
	Contract *ChainLauncherCaller // Generic read-only contract binding to access the raw methods on
}

// ChainLauncherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainLauncherTransactorRaw struct {
	Contract *ChainLauncherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainLauncher creates a new instance of ChainLauncher, bound to a specific deployed contract.
func NewChainLauncher(address common.Address, backend bind.ContractBackend) (*ChainLauncher, error) {
	contract, err := bindChainLauncher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainLauncher{ChainLauncherCaller: ChainLauncherCaller{contract: contract}, ChainLauncherTransactor: ChainLauncherTransactor{contract: contract}, ChainLauncherFilterer: ChainLauncherFilterer{contract: contract}}, nil
}

// NewChainLauncherCaller creates a new read-only instance of ChainLauncher, bound to a specific deployed contract.
func NewChainLauncherCaller(address common.Address, caller bind.ContractCaller) (*ChainLauncherCaller, error) {
	contract, err := bindChainLauncher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainLauncherCaller{contract: contract}, nil
}

// NewChainLauncherTransactor creates a new write-only instance of ChainLauncher, bound to a specific deployed contract.
func NewChainLauncherTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainLauncherTransactor, error) {
	contract, err := bindChainLauncher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainLauncherTransactor{contract: contract}, nil
}

// NewChainLauncherFilterer creates a new log filterer instance of ChainLauncher, bound to a specific deployed contract.
func NewChainLauncherFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainLauncherFilterer, error) {
	contract, err := bindChainLauncher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainLauncherFilterer{contract: contract}, nil
}

// bindChainLauncher binds a generic wrapper to an already deployed contract.
func bindChainLauncher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainLauncherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainLauncher *ChainLauncherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainLauncher.Contract.ChainLauncherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainLauncher *ChainLauncherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainLauncher.Contract.ChainLauncherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainLauncher *ChainLauncherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainLauncher.Contract.ChainLauncherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainLauncher *ChainLauncherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainLauncher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainLauncher *ChainLauncherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainLauncher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainLauncher *ChainLauncherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainLauncher.Contract.contract.Transact(opts, method, params...)
}

// LaunchChain is a paid mutator transaction binding the contract method 0xe2b491e3.
//
// Solidity: function launchChain(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner) returns()
func (_ChainLauncher *ChainLauncherTransactor) LaunchChain(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ChainLauncher.contract.Transact(opts, "launchChain", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner)
}

// LaunchChain is a paid mutator transaction binding the contract method 0xe2b491e3.
//
// Solidity: function launchChain(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner) returns()
func (_ChainLauncher *ChainLauncherSession) LaunchChain(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ChainLauncher.Contract.LaunchChain(&_ChainLauncher.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner)
}

// LaunchChain is a paid mutator transaction binding the contract method 0xe2b491e3.
//
// Solidity: function launchChain(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner) returns()
func (_ChainLauncher *ChainLauncherTransactorSession) LaunchChain(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ChainLauncher.Contract.LaunchChain(&_ChainLauncher.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner)
}

// ChainLauncherChainCreatedIterator is returned from FilterChainCreated and is used to iterate over the raw logs and unpacked data for ChainCreated events raised by the ChainLauncher contract.
type ChainLauncherChainCreatedIterator struct {
	Event *ChainLauncherChainCreated // Event containing the contract specifics and raw log

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
func (it *ChainLauncherChainCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainLauncherChainCreated)
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
		it.Event = new(ChainLauncherChainCreated)
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
func (it *ChainLauncherChainCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainLauncherChainCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainLauncherChainCreated represents a ChainCreated event raised by the ChainLauncher contract.
type ChainLauncherChainCreated struct {
	VmAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChainCreated is a free log retrieval operation binding the contract event 0xa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f.
//
// Solidity: event ChainCreated(address vmAddress)
func (_ChainLauncher *ChainLauncherFilterer) FilterChainCreated(opts *bind.FilterOpts) (*ChainLauncherChainCreatedIterator, error) {

	logs, sub, err := _ChainLauncher.contract.FilterLogs(opts, "ChainCreated")
	if err != nil {
		return nil, err
	}
	return &ChainLauncherChainCreatedIterator{contract: _ChainLauncher.contract, event: "ChainCreated", logs: logs, sub: sub}, nil
}

// WatchChainCreated is a free log subscription operation binding the contract event 0xa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f.
//
// Solidity: event ChainCreated(address vmAddress)
func (_ChainLauncher *ChainLauncherFilterer) WatchChainCreated(opts *bind.WatchOpts, sink chan<- *ChainLauncherChainCreated) (event.Subscription, error) {

	logs, sub, err := _ChainLauncher.contract.WatchLogs(opts, "ChainCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainLauncherChainCreated)
				if err := _ChainLauncher.contract.UnpackLog(event, "ChainCreated", log); err != nil {
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

// ParseChainCreated is a log parse operation binding the contract event 0xa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f.
//
// Solidity: event ChainCreated(address vmAddress)
func (_ChainLauncher *ChainLauncherFilterer) ParseChainCreated(log types.Log) (*ChainLauncherChainCreated, error) {
	event := new(ChainLauncherChainCreated)
	if err := _ChainLauncher.contract.UnpackLog(event, "ChainCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DisputableABI is the input ABI used to generate the binding from.
const DisputableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[3]\",\"name\":\"fields\",\"type\":\"bytes32[3]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"bytes21[]\",\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastMessageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_dataHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_tokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"}],\"name\":\"generateLastMessageHashStub\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"withinTimeBounds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DisputableFuncSigs maps the 4-byte function signature to its string representation.
var DisputableFuncSigs = map[string]string{
	"924e7b37": "confirmDisputableAsserted(VM.Data storage,bytes32,bytes32,uint32,bytes21[],bytes,uint16[],uint256[],address[],bytes32)",
	"d14cf098": "generateLastMessageHashStub(bytes21[],bytes32[],uint16[],uint256[],address[])",
	"5af93c7a": "initiateChallenge(VM.Data storage,bytes32)",
	"c97c8eec": "pendingDisputableAssert(VM.Data storage,bytes32[4],uint32,uint64[2],bytes21[],bytes32[],uint16[],uint256[],address[])",
	"42c0787e": "withinTimeBounds(uint64[2])",
}

// DisputableBin is the compiled bytecode used for deploying new contracts.
var DisputableBin = "0x612400610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c806342c0787e146100665780635af93c7a146100c5578063924e7b37146100f7578063c97c8eec146103d1578063d14cf098146106f9575b600080fd5b6100b16004803603604081101561007c57600080fd5b604080518082018252918301929181830191839060029083908390808284376000920191909152509194506109b49350505050565b604080519115158252519081900360200190f35b8180156100d157600080fd5b506100f5600480360360408110156100e857600080fd5b50803590602001356109e6565b005b81801561010357600080fd5b506100f5600480360361014081101561011b57600080fd5b81359160208101359160408201359163ffffffff6060820135169181019060a081016080820135600160201b81111561015357600080fd5b82018360208201111561016557600080fd5b803590602001918460208302840111600160201b8311171561018657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101d557600080fd5b8201836020820111156101e757600080fd5b803590602001918460018302840111600160201b8311171561020857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561025a57600080fd5b82018360208201111561026c57600080fd5b803590602001918460208302840111600160201b8311171561028d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102dc57600080fd5b8201836020820111156102ee57600080fd5b803590602001918460208302840111600160201b8311171561030f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561035e57600080fd5b82018360208201111561037057600080fd5b803590602001918460208302840111600160201b8311171561039157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250610be8915050565b8180156103dd57600080fd5b506100f560048036036101a08110156103f557600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561048057600080fd5b82018360208201111561049257600080fd5b803590602001918460208302840111600160201b831117156104b357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561050257600080fd5b82018360208201111561051457600080fd5b803590602001918460208302840111600160201b8311171561053557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561058457600080fd5b82018360208201111561059657600080fd5b803590602001918460208302840111600160201b831117156105b757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561060657600080fd5b82018360208201111561061857600080fd5b803590602001918460208302840111600160201b8311171561063957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561068857600080fd5b82018360208201111561069a57600080fd5b803590602001918460208302840111600160201b831117156106bb57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610c50945050505050565b6109a2600480360360a081101561070f57600080fd5b810190602081018135600160201b81111561072957600080fd5b82018360208201111561073b57600080fd5b803590602001918460208302840111600160201b8311171561075c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107ab57600080fd5b8201836020820111156107bd57600080fd5b803590602001918460208302840111600160201b831117156107de57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561082d57600080fd5b82018360208201111561083f57600080fd5b803590602001918460208302840111600160201b8311171561086057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156108af57600080fd5b8201836020820111156108c157600080fd5b803590602001918460208302840111600160201b831117156108e257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561093157600080fd5b82018360208201111561094357600080fd5b803590602001918460208302840111600160201b8311171561096457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610cf9945050505050565b60408051918252519081900360200190f35b805160009067ffffffffffffffff1643108015906109e05750602082015167ffffffffffffffff164311155b92915050565b60038201546001600160a01b0316331415610a325760405162461bcd60e51b81526004018080602001828103825260218152602001806122376021913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610a8157600080fd5b505af4158015610a95573d6000803e3d6000fd5b505050506040513d6020811015610aab57600080fd5b5051610ae85760405162461bcd60e51b81526004018080602001828103825260268152602001806122df6026913960400191505060405180910390fd5b60026005830154600160401b900460ff166003811115610b0457fe5b14610b405760405162461bcd60e51b815260040180806020018281038252602f815260200180612170602f913960400191505060405180910390fd5b81600101548114610b825760405162461bcd60e51b815260040180806020018281038252604d815260200180612332604d913960600191505060405180910390fd5b6000600183015560058201805460ff60401b1916600160401b1769ff0000000000000000001916600160481b1790556040805133815290517f255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b79181900360200190a15050565b610c448a6040518060a001604052808c81526020018b81526020018a63ffffffff1681526020018981526020016040518060a001604052808a815260200189815260200188815260200187815260200186815250815250610f50565b50505050505050505050565b610cee896040518061016001604052808b600060048110610c6d57fe5b602002015181526020018b600160048110610c8457fe5b602002015181526020018b600260048110610c9b57fe5b602002015181526020018a63ffffffff1681526020018981526020018881526020018781526020018681526020018581526020018481526020018b600360048110610ce257fe5b602002015190526114af565b505050505050505050565b60008351855114610d47576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8251855114610d93576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8151855114610ddf576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b84516000908190815b81811015610f425773__$9836fa7140e5a33041d4b827682e675a30$__624c28f68a8381518110610e1557fe5b60200260200101518c8b8581518110610e2a57fe5b602002602001015161ffff1681518110610e4057fe5b60200260200101518a8581518110610e5457fe5b60200260200101518a8681518110610e6857fe5b60200260200101516040518563ffffffff1660e01b815260040180858152602001846001600160581b0319166001600160581b0319168152602001838152602001826001600160a01b03166001600160a01b0316815260200194505050505060206040518083038186803b158015610edf57600080fd5b505af4158015610ef3573d6000803e3d6000fd5b505050506040513d6020811015610f0957600080fd5b50516040805160208181019790975280820183905281518082038301815260609091019091528051950194909420939250600101610de8565b509198975050505050505050565b60026005830154600160401b900460ff166003811115610f6c57fe5b14610fa85760405162461bcd60e51b81526004018080602001828103825260228152602001806121f16022913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610ff757600080fd5b505af415801561100b573d6000803e3d6000fd5b505050506040513d602081101561102157600080fd5b50511561105f5760405162461bcd60e51b81526004018080602001828103825260248152602001806121cd6024913960400191505060405180910390fd5b6001820154815160208084015160408086015160608088015160808901518051968101519481015192015173__$9836fa7140e5a33041d4b827682e675a30$__96632090372196956000946110b49493611e7e565b600089608001516080015173__$9836fa7140e5a33041d4b827682e675a30$__630f89fbff8c606001518d60800151602001518e60800151604001516040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b8381101561114757818101518382015260200161112f565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561118657818101518382015260200161116e565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156111c55781810151838201526020016111ad565b50505050905001965050505050505060006040518083038186803b1580156111ec57600080fd5b505af4158015611200573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561122957600080fd5b8101908080516040519392919084600160201b82111561124857600080fd5b90830190602082018581111561125d57600080fd5b82518660208202830111600160201b8211171561127957600080fd5b82525081516020918201928201910280838360005b838110156112a657818101518382015260200161128e565b505050509050016040525050506040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff1681526020018660001b81526020018581526020018460001b815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015611338578181015183820152602001611320565b505050509050019850505050505050505060206040518083038186803b15801561136157600080fd5b505af4158015611375573d6000803e3d6000fd5b505050506040513d602081101561138b57600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920190528051910120146113f25760405162461bcd60e51b815260040180806020018281038252604d81526020018061237f604d913960600191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63eb49982c8383602001516040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561144d57600080fd5b505af4158015611461573d6000803e3d6000fd5b5050506020808301516080808501510151604080519283529282015281517f4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb93509081900390910190a15050565b60016005830154600160401b900460ff1660038111156114cb57fe5b146115075760405162461bcd60e51b815260040180806020018281038252602d815260200180612305602d913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__632a3e0a97836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561155657600080fd5b505af415801561156a573d6000803e3d6000fd5b505050506040513d602081101561158057600080fd5b5051158015611607575073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156115d957600080fd5b505af41580156115ed573d6000803e3d6000fd5b505050506040513d602081101561160357600080fd5b5051155b6116425760405162461bcd60e51b815260040180806020018281038252603e815260200180612258603e913960400191505060405180910390fd5b6005820154600160481b900460ff161561168d5760405162461bcd60e51b815260040180806020018281038252602e81526020018061219f602e913960400191505060405180910390fd5b6005820154606082015163ffffffff600160201b9092048216911611156116fb576040805162461bcd60e51b815260206004820152601f60248201527f547269656420746f206578656375746520746f6f206d616e7920737465707300604482015290519081900360640190fd5b61170881608001516109b4565b6117435760405162461bcd60e51b81526004018080602001828103825260248152602001806122136024913960400191505060405180910390fd5b81548151146117835760405162461bcd60e51b81526004018080602001828103825260278152602001806122b86027913960400191505060405180910390fd5b80602001518260020154146117c95760405162461bcd60e51b81526004018080602001828103825260228152602001806122966022913960400191505060405180910390fd5b606073__$9836fa7140e5a33041d4b827682e675a30$__630f89fbff8360a001518460e001518561010001516040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b8381101561184c578181015183820152602001611834565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561188b578181015183820152602001611873565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156118ca5781810151838201526020016118b2565b50505050905001965050505050505060006040518083038186803b1580156118f157600080fd5b505af4158015611905573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561192e57600080fd5b8101908080516040519392919084600160201b82111561194d57600080fd5b90830190602082018581111561196257600080fd5b82518660208202830111600160201b8211171561197e57600080fd5b82525081516020918201928201910280838360005b838110156119ab578181015183820152602001611993565b50505050905001604052505050905073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63a3a162cb846040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015611a0957600080fd5b505af4158015611a1d573d6000803e3d6000fd5b505050506000611a468360a001518460c001518560e00151866101000151876101200151610cf9565b905073__$9836fa7140e5a33041d4b827682e675a30$__633e2855988460000151856080015186602001518760a00151876040518663ffffffff1660e01b81526004018086815260200185600260200280838360005b83811015611ab4578181015183820152602001611a9c565b505050509050018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015611b01578181015183820152602001611ae9565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611b40578181015183820152602001611b28565b5050505090500197505050505050505060206040518083038186803b158015611b6857600080fd5b505af4158015611b7c573d6000803e3d6000fd5b505050506040513d6020811015611b9257600080fd5b505160408085015160608601516101408701519251632090372160e01b81526004810183815263ffffffff83166024830152600060448301819052606483018890526084830181905260a4830186905260e060c48401908152895160e4850152895173__$9836fa7140e5a33041d4b827682e675a30$__97632090372197969593948b94869492938e93916101040190602085810191028083838a5b83811015611c46578181015183820152602001611c2e565b505050509050019850505050505050505060206040518083038186803b158015611c6f57600080fd5b505af4158015611c83573d6000803e3d6000fd5b505050506040513d6020811015611c9957600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019052805191012060018501556003840180546001600160a01b031916331790556005840180546002919060ff60401b1916600160401b8302179055507f5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f6040518060600160405280856000015181526020018560200151815260200185604001518152503385608001518660a00151876060015186896101400151896040518089600360200280838360005b83811015611d81578181015183820152602001611d69565b505050506001600160a01b038b1692019182525060200187604080838360005b83811015611db9578181015183820152602001611da1565b50505050905001806020018663ffffffff1663ffffffff16815260200185815260200184815260200180602001838103835288818151815260200191508051906020019060200280838360005b83811015611e1e578181015183820152602001611e06565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611e5d578181015183820152602001611e45565b505050509050019a505050505050505050505060405180910390a150505050565b60008151835114611ecc576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8351835114611f18576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b825160009081908190815b818110156121605773__$d969135829891f807aa9c34494da4ecd99$__6389df40da8b866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015611f99578181015183820152602001611f81565b50505050905090810190601f168015611fc65780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b158015611fe357600080fd5b505af4158015611ff7573d6000803e3d6000fd5b505050506040513d604081101561200d57600080fd5b5080516020909101518a51919550935073__$9836fa7140e5a33041d4b827682e675a30$__90624c28f69085908e908d908690811061204857fe5b602002602001015161ffff168151811061205e57fe5b60200260200101518b858151811061207257fe5b60200260200101518b868151811061208657fe5b60200260200101516040518563ffffffff1660e01b815260040180858152602001846001600160581b0319166001600160581b0319168152602001838152602001826001600160a01b03166001600160a01b0316815260200194505050505060206040518083038186803b1580156120fd57600080fd5b505af4158015612111573d6000803e3d6000fd5b505050506040513d602081101561212757600080fd5b50516040805160208181019890985280820183905281518082038301815260609091019091528051960195909520949250600101611f23565b5092999850505050505050505056fe417373657274696f6e206d7573742062652070656e64696e6720746f20696e697469617465206368616c6c656e676543616e206f6e6c792064697370757461626c6520617373657274206966206e6f7420696e206368616c6c656e6765417373657274696f6e206973207374696c6c2070656e64696e67206368616c6c656e6765564d20646f6573206e6f74206861766520617373657274696f6e2070656e64696e67507265636f6e646974696f6e3a206e6f742077697468696e2074696d6520626f756e64734368616c6c656e676520776173206372656174656420627920617373657274657243616e206f6e6c792064697370757461626c6520617373657274206966206d616368696e65206973206e6f74206572726f726564206f722068616c746564507265636f6e646974696f6e3a20696e626f7820646f6573206e6f74206d61746368507265636f6e646974696f6e3a207374617465206861736820646f6573206e6f74206d617463684368616c6c656e676520646964206e6f7420636f6d65206265666f726520646561646c696e6543616e206f6e6c792064697370757461626c65206173736572742066726f6d2077616974696e67207374617465496e697469617465204368616c6c656e67653a20507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6e436f6e6669726d2044697370757461626c653a20507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6ea265627a7a723158205ca5c3c0b967a27f5201957744dbab5bb1a036908fed9bfb42e9611dcc8eaa9364736f6c634300050c0032"

// DeployDisputable deploys a new Ethereum contract, binding an instance of Disputable to it.
func DeployDisputable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Disputable, error) {
	parsed, err := abi.JSON(strings.NewReader(DisputableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	vMAddr, _, _, _ := DeployVM(auth, backend)
	DisputableBin = strings.Replace(DisputableBin, "__$8e266570c8a7fb2aaac83b3e040afaf9e1$__", vMAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	DisputableBin = strings.Replace(DisputableBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	DisputableBin = strings.Replace(DisputableBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DisputableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Disputable{DisputableCaller: DisputableCaller{contract: contract}, DisputableTransactor: DisputableTransactor{contract: contract}, DisputableFilterer: DisputableFilterer{contract: contract}}, nil
}

// Disputable is an auto generated Go binding around an Ethereum contract.
type Disputable struct {
	DisputableCaller     // Read-only binding to the contract
	DisputableTransactor // Write-only binding to the contract
	DisputableFilterer   // Log filterer for contract events
}

// DisputableCaller is an auto generated read-only Go binding around an Ethereum contract.
type DisputableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DisputableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DisputableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DisputableSession struct {
	Contract     *Disputable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DisputableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DisputableCallerSession struct {
	Contract *DisputableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DisputableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DisputableTransactorSession struct {
	Contract     *DisputableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DisputableRaw is an auto generated low-level Go binding around an Ethereum contract.
type DisputableRaw struct {
	Contract *Disputable // Generic contract binding to access the raw methods on
}

// DisputableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DisputableCallerRaw struct {
	Contract *DisputableCaller // Generic read-only contract binding to access the raw methods on
}

// DisputableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DisputableTransactorRaw struct {
	Contract *DisputableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDisputable creates a new instance of Disputable, bound to a specific deployed contract.
func NewDisputable(address common.Address, backend bind.ContractBackend) (*Disputable, error) {
	contract, err := bindDisputable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Disputable{DisputableCaller: DisputableCaller{contract: contract}, DisputableTransactor: DisputableTransactor{contract: contract}, DisputableFilterer: DisputableFilterer{contract: contract}}, nil
}

// NewDisputableCaller creates a new read-only instance of Disputable, bound to a specific deployed contract.
func NewDisputableCaller(address common.Address, caller bind.ContractCaller) (*DisputableCaller, error) {
	contract, err := bindDisputable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DisputableCaller{contract: contract}, nil
}

// NewDisputableTransactor creates a new write-only instance of Disputable, bound to a specific deployed contract.
func NewDisputableTransactor(address common.Address, transactor bind.ContractTransactor) (*DisputableTransactor, error) {
	contract, err := bindDisputable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DisputableTransactor{contract: contract}, nil
}

// NewDisputableFilterer creates a new log filterer instance of Disputable, bound to a specific deployed contract.
func NewDisputableFilterer(address common.Address, filterer bind.ContractFilterer) (*DisputableFilterer, error) {
	contract, err := bindDisputable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DisputableFilterer{contract: contract}, nil
}

// bindDisputable binds a generic wrapper to an already deployed contract.
func bindDisputable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DisputableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Disputable *DisputableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Disputable.Contract.DisputableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Disputable *DisputableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Disputable.Contract.DisputableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Disputable *DisputableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Disputable.Contract.DisputableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Disputable *DisputableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Disputable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Disputable *DisputableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Disputable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Disputable *DisputableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Disputable.Contract.contract.Transact(opts, method, params...)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xd14cf098.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_Disputable *DisputableCaller) GenerateLastMessageHashStub(opts *bind.CallOpts, _tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Disputable.contract.Call(opts, out, "generateLastMessageHashStub", _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
	return *ret0, err
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xd14cf098.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_Disputable *DisputableSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	return _Disputable.Contract.GenerateLastMessageHashStub(&_Disputable.CallOpts, _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xd14cf098.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_Disputable *DisputableCallerSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	return _Disputable.Contract.GenerateLastMessageHashStub(&_Disputable.CallOpts, _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_Disputable *DisputableCaller) WithinTimeBounds(opts *bind.CallOpts, _timeBounds [2]uint64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Disputable.contract.Call(opts, out, "withinTimeBounds", _timeBounds)
	return *ret0, err
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_Disputable *DisputableSession) WithinTimeBounds(_timeBounds [2]uint64) (bool, error) {
	return _Disputable.Contract.WithinTimeBounds(&_Disputable.CallOpts, _timeBounds)
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_Disputable *DisputableCallerSession) WithinTimeBounds(_timeBounds [2]uint64) (bool, error) {
	return _Disputable.Contract.WithinTimeBounds(&_Disputable.CallOpts, _timeBounds)
}

// DisputableConfirmedDisputableAssertionIterator is returned from FilterConfirmedDisputableAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedDisputableAssertion events raised by the Disputable contract.
type DisputableConfirmedDisputableAssertionIterator struct {
	Event *DisputableConfirmedDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *DisputableConfirmedDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputableConfirmedDisputableAssertion)
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
		it.Event = new(DisputableConfirmedDisputableAssertion)
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
func (it *DisputableConfirmedDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputableConfirmedDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputableConfirmedDisputableAssertion represents a ConfirmedDisputableAssertion event raised by the Disputable contract.
type DisputableConfirmedDisputableAssertion struct {
	NewState    [32]byte
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_Disputable *DisputableFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts) (*DisputableConfirmedDisputableAssertionIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &DisputableConfirmedDisputableAssertionIterator{contract: _Disputable.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_Disputable *DisputableFilterer) WatchConfirmedDisputableAssertion(opts *bind.WatchOpts, sink chan<- *DisputableConfirmedDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _Disputable.contract.WatchLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputableConfirmedDisputableAssertion)
				if err := _Disputable.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
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

// ParseConfirmedDisputableAssertion is a log parse operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_Disputable *DisputableFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*DisputableConfirmedDisputableAssertion, error) {
	event := new(DisputableConfirmedDisputableAssertion)
	if err := _Disputable.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DisputableInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the Disputable contract.
type DisputableInitiatedChallengeIterator struct {
	Event *DisputableInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *DisputableInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputableInitiatedChallenge)
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
		it.Event = new(DisputableInitiatedChallenge)
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
func (it *DisputableInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputableInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputableInitiatedChallenge represents a InitiatedChallenge event raised by the Disputable contract.
type DisputableInitiatedChallenge struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_Disputable *DisputableFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*DisputableInitiatedChallengeIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &DisputableInitiatedChallengeIterator{contract: _Disputable.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_Disputable *DisputableFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *DisputableInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _Disputable.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputableInitiatedChallenge)
				if err := _Disputable.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_Disputable *DisputableFilterer) ParseInitiatedChallenge(log types.Log) (*DisputableInitiatedChallenge, error) {
	event := new(DisputableInitiatedChallenge)
	if err := _Disputable.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DisputablePendingDisputableAssertionIterator is returned from FilterPendingDisputableAssertion and is used to iterate over the raw logs and unpacked data for PendingDisputableAssertion events raised by the Disputable contract.
type DisputablePendingDisputableAssertionIterator struct {
	Event *DisputablePendingDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *DisputablePendingDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputablePendingDisputableAssertion)
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
		it.Event = new(DisputablePendingDisputableAssertion)
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
func (it *DisputablePendingDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputablePendingDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputablePendingDisputableAssertion represents a PendingDisputableAssertion event raised by the Disputable contract.
type DisputablePendingDisputableAssertion struct {
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

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_Disputable *DisputableFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*DisputablePendingDisputableAssertionIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &DisputablePendingDisputableAssertionIterator{contract: _Disputable.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_Disputable *DisputableFilterer) WatchPendingDisputableAssertion(opts *bind.WatchOpts, sink chan<- *DisputablePendingDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _Disputable.contract.WatchLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputablePendingDisputableAssertion)
				if err := _Disputable.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_Disputable *DisputableFilterer) ParsePendingDisputableAssertion(log types.Log) (*DisputablePendingDisputableAssertion, error) {
	event := new(DisputablePendingDisputableAssertion)
	if err := _Disputable.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
const IChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeManagerFuncSigs = map[string]string{
	"208e04d4": "initiateChallenge(address[2],uint128[2],uint32,bytes32)",
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

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initiateChallenge", players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerSession) InitiateChallenge(players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) InitiateChallenge(players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, players, escrows, challengePeriod, challengeRoot)
}

// IGlobalPendingInboxABI is the input ABI used to generate the binding from.
const IGlobalPendingInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes21\",\"name\":\"tokenType\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"hasFunds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pullPendingMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"bytes\",\"name\":\"_messageData\",\"type\":\"bytes\"},{\"internalType\":\"uint16[]\",\"name\":\"_tokenTypeNum\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var IGlobalPendingInboxFuncSigs = map[string]string{
	"3bbc3c32": "forwardMessage(address,bytes21,uint256,bytes,bytes)",
	"acb633b6": "hasFunds(address,bytes21[],uint256[])",
	"d106ec19": "pullPendingMessages()",
	"f3972383": "registerForInbox()",
	"3fc6eb80": "sendEthMessage(address,bytes)",
	"626cef85": "sendMessage(address,bytes21,uint256,bytes)",
	"ec22a767": "sendMessages(bytes21[],bytes,uint16[],uint256[],address[])",
}

// IGlobalPendingInbox is an auto generated Go binding around an Ethereum contract.
type IGlobalPendingInbox struct {
	IGlobalPendingInboxCaller     // Read-only binding to the contract
	IGlobalPendingInboxTransactor // Write-only binding to the contract
	IGlobalPendingInboxFilterer   // Log filterer for contract events
}

// IGlobalPendingInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGlobalPendingInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGlobalPendingInboxSession struct {
	Contract     *IGlobalPendingInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGlobalPendingInboxCallerSession struct {
	Contract *IGlobalPendingInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IGlobalPendingInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGlobalPendingInboxTransactorSession struct {
	Contract     *IGlobalPendingInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGlobalPendingInboxRaw struct {
	Contract *IGlobalPendingInbox // Generic contract binding to access the raw methods on
}

// IGlobalPendingInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCallerRaw struct {
	Contract *IGlobalPendingInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IGlobalPendingInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactorRaw struct {
	Contract *IGlobalPendingInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGlobalPendingInbox creates a new instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInbox(address common.Address, backend bind.ContractBackend) (*IGlobalPendingInbox, error) {
	contract, err := bindIGlobalPendingInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInbox{IGlobalPendingInboxCaller: IGlobalPendingInboxCaller{contract: contract}, IGlobalPendingInboxTransactor: IGlobalPendingInboxTransactor{contract: contract}, IGlobalPendingInboxFilterer: IGlobalPendingInboxFilterer{contract: contract}}, nil
}

// NewIGlobalPendingInboxCaller creates a new read-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxCaller(address common.Address, caller bind.ContractCaller) (*IGlobalPendingInboxCaller, error) {
	contract, err := bindIGlobalPendingInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxCaller{contract: contract}, nil
}

// NewIGlobalPendingInboxTransactor creates a new write-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IGlobalPendingInboxTransactor, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxTransactor{contract: contract}, nil
}

// NewIGlobalPendingInboxFilterer creates a new log filterer instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IGlobalPendingInboxFilterer, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxFilterer{contract: contract}, nil
}

// bindIGlobalPendingInbox binds a generic wrapper to an already deployed contract.
func bindIGlobalPendingInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGlobalPendingInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transact(opts, method, params...)
}

// HasFunds is a free data retrieval call binding the contract method 0xacb633b6.
//
// Solidity: function hasFunds(address _owner, bytes21[] _tokenTypes, uint256[] _amounts) constant returns(bool)
func (_IGlobalPendingInbox *IGlobalPendingInboxCaller) HasFunds(opts *bind.CallOpts, _owner common.Address, _tokenTypes [][21]byte, _amounts []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IGlobalPendingInbox.contract.Call(opts, out, "hasFunds", _owner, _tokenTypes, _amounts)
	return *ret0, err
}

// HasFunds is a free data retrieval call binding the contract method 0xacb633b6.
//
// Solidity: function hasFunds(address _owner, bytes21[] _tokenTypes, uint256[] _amounts) constant returns(bool)
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) HasFunds(_owner common.Address, _tokenTypes [][21]byte, _amounts []*big.Int) (bool, error) {
	return _IGlobalPendingInbox.Contract.HasFunds(&_IGlobalPendingInbox.CallOpts, _owner, _tokenTypes, _amounts)
}

// HasFunds is a free data retrieval call binding the contract method 0xacb633b6.
//
// Solidity: function hasFunds(address _owner, bytes21[] _tokenTypes, uint256[] _amounts) constant returns(bool)
func (_IGlobalPendingInbox *IGlobalPendingInboxCallerSession) HasFunds(_owner common.Address, _tokenTypes [][21]byte, _amounts []*big.Int) (bool, error) {
	return _IGlobalPendingInbox.Contract.HasFunds(&_IGlobalPendingInbox.CallOpts, _owner, _tokenTypes, _amounts)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) ForwardMessage(opts *bind.TransactOpts, _destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "forwardMessage", _destination, _tokenType, _amount, _data, _signature)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) ForwardMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data, _signature)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) ForwardMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data, _signature)
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0xd106ec19.
//
// Solidity: function pullPendingMessages() returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) PullPendingMessages(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "pullPendingMessages")
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0xd106ec19.
//
// Solidity: function pullPendingMessages() returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) PullPendingMessages() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.PullPendingMessages(&_IGlobalPendingInbox.TransactOpts)
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0xd106ec19.
//
// Solidity: function pullPendingMessages() returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) PullPendingMessages() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.PullPendingMessages(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) RegisterForInbox(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "registerForInbox")
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendEthMessage(opts *bind.TransactOpts, _destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendEthMessage", _destination, _data)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendEthMessage(_destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendEthMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _data)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendEthMessage(_destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendEthMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessage(opts *bind.TransactOpts, _destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessage", _destination, _tokenType, _amount, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data)
}

// SendMessages is a paid mutator transaction binding the contract method 0xec22a767.
//
// Solidity: function sendMessages(bytes21[] _tokenTypes, bytes _messageData, uint16[] _tokenTypeNum, uint256[] _amounts, address[] _destinations) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessages(opts *bind.TransactOpts, _tokenTypes [][21]byte, _messageData []byte, _tokenTypeNum []uint16, _amounts []*big.Int, _destinations []common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessages", _tokenTypes, _messageData, _tokenTypeNum, _amounts, _destinations)
}

// SendMessages is a paid mutator transaction binding the contract method 0xec22a767.
//
// Solidity: function sendMessages(bytes21[] _tokenTypes, bytes _messageData, uint16[] _tokenTypeNum, uint256[] _amounts, address[] _destinations) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessages(_tokenTypes [][21]byte, _messageData []byte, _tokenTypeNum []uint16, _amounts []*big.Int, _destinations []common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _tokenTypes, _messageData, _tokenTypeNum, _amounts, _destinations)
}

// SendMessages is a paid mutator transaction binding the contract method 0xec22a767.
//
// Solidity: function sendMessages(bytes21[] _tokenTypes, bytes _messageData, uint16[] _tokenTypeNum, uint256[] _amounts, address[] _destinations) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessages(_tokenTypes [][21]byte, _messageData []byte, _tokenTypeNum []uint16, _amounts []*big.Int, _destinations []common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _tokenTypes, _messageData, _tokenTypeNum, _amounts, _destinations)
}

// IGlobalPendingInboxMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxMessageDelivered)
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
func (it *IGlobalPendingInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxMessageDelivered represents a MessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxMessageDelivered struct {
	VmId      common.Address
	Sender    common.Address
	TokenType [21]byte
	Value     *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterMessageDelivered(opts *bind.FilterOpts, vmId []common.Address) (*IGlobalPendingInboxMessageDeliveredIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "MessageDelivered", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxMessageDelivered, vmId []common.Address) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "MessageDelivered", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseMessageDelivered(log types.Log) (*IGlobalPendingInboxMessageDelivered, error) {
	event := new(IGlobalPendingInboxMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820c718d6b10f268fec78a101782cfc4dd183b6cfc3b616aba4f2bc1d3eb1f83dd164736f6c634300050c0032"

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

// VMABI is the input ABI used to generate the binding from.
const VMABI = "[]"

// VMFuncSigs maps the 4-byte function signature to its string representation.
var VMFuncSigs = map[string]string{
	"eb49982c": "acceptAssertion(VM.Data storage,bytes32)",
	"2a3e0a97": "isErrored(VM.Data storage)",
	"e2fe93ca": "isHalted(VM.Data storage)",
	"a3a162cb": "resetDeadline(VM.Data storage)",
	"8ab48be5": "withinDeadline(VM.Data storage)",
}

// VMBin is the compiled bytecode used for deploying new contracts.
var VMBin = "0x6101ea610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80632a3e0a97146100665780638ab48be514610097578063a3a162cb146100b4578063e2fe93ca146100e0578063eb49982c146100fd575b600080fd5b6100836004803603602081101561007c57600080fd5b503561012d565b604080519115158252519081900360200190f35b610083600480360360208110156100ad57600080fd5b5035610134565b8180156100c057600080fd5b506100de600480360360208110156100d757600080fd5b503561014f565b005b610083600480360360208110156100f657600080fd5b503561018e565b81801561010957600080fd5b506100de6004803603604081101561012057600080fd5b5080359060200135610193565b5460011490565b60040154600160801b900467ffffffffffffffff1643111590565b60058101546004909101805467ffffffffffffffff60801b1916600160801b63ffffffff909316430167ffffffffffffffff1692909202919091179055565b541590565b8155600501805468ff000000000000000019166801000000000000000017905556fea265627a7a723158205595b42683b610c72e205d7b576262103534d458dd1544f803e09fb250d894c564736f6c634300050c0032"

// DeployVM deploys a new Ethereum contract, binding an instance of VM to it.
func DeployVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VM, error) {
	parsed, err := abi.JSON(strings.NewReader(VMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VM{VMCaller: VMCaller{contract: contract}, VMTransactor: VMTransactor{contract: contract}, VMFilterer: VMFilterer{contract: contract}}, nil
}

// VM is an auto generated Go binding around an Ethereum contract.
type VM struct {
	VMCaller     // Read-only binding to the contract
	VMTransactor // Write-only binding to the contract
	VMFilterer   // Log filterer for contract events
}

// VMCaller is an auto generated read-only Go binding around an Ethereum contract.
type VMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VMSession struct {
	Contract     *VM               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VMCallerSession struct {
	Contract *VMCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VMTransactorSession struct {
	Contract     *VMTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMRaw is an auto generated low-level Go binding around an Ethereum contract.
type VMRaw struct {
	Contract *VM // Generic contract binding to access the raw methods on
}

// VMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VMCallerRaw struct {
	Contract *VMCaller // Generic read-only contract binding to access the raw methods on
}

// VMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VMTransactorRaw struct {
	Contract *VMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVM creates a new instance of VM, bound to a specific deployed contract.
func NewVM(address common.Address, backend bind.ContractBackend) (*VM, error) {
	contract, err := bindVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VM{VMCaller: VMCaller{contract: contract}, VMTransactor: VMTransactor{contract: contract}, VMFilterer: VMFilterer{contract: contract}}, nil
}

// NewVMCaller creates a new read-only instance of VM, bound to a specific deployed contract.
func NewVMCaller(address common.Address, caller bind.ContractCaller) (*VMCaller, error) {
	contract, err := bindVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VMCaller{contract: contract}, nil
}

// NewVMTransactor creates a new write-only instance of VM, bound to a specific deployed contract.
func NewVMTransactor(address common.Address, transactor bind.ContractTransactor) (*VMTransactor, error) {
	contract, err := bindVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VMTransactor{contract: contract}, nil
}

// NewVMFilterer creates a new log filterer instance of VM, bound to a specific deployed contract.
func NewVMFilterer(address common.Address, filterer bind.ContractFilterer) (*VMFilterer, error) {
	contract, err := bindVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VMFilterer{contract: contract}, nil
}

// bindVM binds a generic wrapper to an already deployed contract.
func bindVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VM *VMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VM.Contract.VMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VM *VMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VM.Contract.VMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VM *VMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VM.Contract.VMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VM *VMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VM *VMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VM *VMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VM.Contract.contract.Transact(opts, method, params...)
}
