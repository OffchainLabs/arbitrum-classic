// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbchain

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

// ArbBaseABI is the input ABI used to generate the binding from.
const ArbBaseABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"ChallengeLaunched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numGas\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumVM.State\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"enumVM.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"activeChallengeManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbBaseFuncSigs maps the 4-byte function signature to its string representation.
var ArbBaseFuncSigs = map[string]string{
	"94af716b": "activateVM()",
	"5dbaf68b": "challengeFactory()",
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"4676a75e": "confirmDisputableAsserted(bytes32,uint64[2],bytes32,bytes32,bytes32,bool,uint32,uint64,bytes,bytes32)",
	"08dc89d7": "currentDeposit(address)",
	"aca0f372": "escrowRequired()",
	"6be00229": "exitAddress()",
	"1865c57d": "getState()",
	"d489113a": "globalInbox()",
	"4cbb9ff2": "initialize(bytes32,uint32,uint32,uint128,address,address,address)",
	"0badcbbf": "initiateChallenge(bytes32,bytes32,uint64[2],bytes32)",
	"8da5cb5b": "owner()",
	"cfa80707": "ownerShutdown()",
	"6d45809f": "pendingDisputableAssert(bytes32,bytes32,bytes32,bool,bytes32,bytes32,uint32,uint64,uint64[2])",
	"60675a87": "terminateAddress()",
	"3a768463": "vm()",
}

// ArbBaseBin is the compiled bytecode used for deploying new contracts.
var ArbBaseBin = "0x608060405234801561001057600080fd5b506118e5806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806360675a871161009757806394af716b1161006657806394af716b146104ca578063aca0f372146104d2578063cfa80707146104da578063d489113a146104e257610100565b806360675a87146104225780636be002291461042a5780636d45809f146104325780638da5cb5b146104c257610100565b80633a768463116100d35780633a768463146101e45780634676a75e146102895780634cbb9ff21461039b5780635dbaf68b146103fe57610100565b806308dc89d7146101055780630badcbbf1461013d5780631865c57d1461019857806322c091bc146101c4575b600080fd5b61012b6004803603602081101561011b57600080fd5b50356001600160a01b03166104ea565b60408051918252519081900360200190f35b610196600480360360a081101561015357600080fd5b6040805180820182528335936020810135938101929091608083019180840190600290839083908082843760009201919091525091945050903591506105099050565b005b6101a0610819565b604051808260038111156101b057fe5b60ff16815260200191505060405180910390f35b610196600480360360808110156101da57600080fd5b5060408101610829565b6101ec61093b565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561025c57fe5b60ff1681526001600160a01b039092166020830152506040805191829003019a5098505050505050505050f35b61019660048036036101608110156102a057600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091948335946020850135946040810135945060608101351515935063ffffffff6080820135169267ffffffffffffffff60a08301351692909160e081019060c0013564010000000081111561032457600080fd5b82018360208201111561033657600080fd5b8035906020019184600183028401116401000000008311171561035857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506109ac915050565b610196600480360360e08110156103b157600080fd5b5080359063ffffffff60208201358116916040810135909116906001600160801b03606082013516906001600160a01b03608082013581169160a081013582169160c09091013516610b6e565b610406610d9a565b604080516001600160a01b039092168252519081900360200190f35b610406610da9565b610406610db8565b610196600480360361014081101561044957600080fd5b6040805180820182528335936020810135938382013593606083013515159360808401359360a08101359363ffffffff60c0830135169367ffffffffffffffff60e0840135169391830192916101408301916101008401906002908390839080828437600092019190915250919450610dc79350505050565b610406610f49565b610196610f58565b61012b610fe6565b610196610ff5565b610406611055565b6001600160a01b0381166000908152600860205260409020545b919050565b336000908152600860205260409020546006546001600160801b031611156105625760405162461bcd60e51b815260040180806020018281038252602781526020018061180b6027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b0390941690930390925581516337d8913360e01b8152600260048201818152602483018990526044830188905273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__946337d891339492938a938a938a938a9391926064909201918591908190849084905b838110156105fb5781810151838201526020016105e3565b505050509050018281526020019550505050505060006040518083038186803b15801561062757600080fd5b505af415801561063b573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b0316808252918101919091526007548451631a6ef3c360e31b815292909516975063d3779e1896509194919363ffffffff16928b928b928b928b9260049092019182918a918190849084905b838110156106d95781810151838201526020016106c1565b5050505090500187600260200280838360005b838110156107045781810151838201526020016106ec565b505050509050018663ffffffff1663ffffffff16815260200185815260200184815260200183600260200280838360005b8381101561074d578181015183820152602001610735565b50505050905001828152602001975050505050505050602060405180830381600087803b15801561077d57600080fd5b505af1158015610791573d6000803e3d6000fd5b505050506040513d60208110156107a757600080fd5b505160078054600160481b600160e81b031916600160481b6001600160a01b03938416810291909117918290556040805191909204909216825233602083015280517f65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f9281900390910190a150505050565b600754600160401b900460ff1690565b600754600160481b90046001600160a01b031633146108795760405162461bcd60e51b815260040180806020018281038252602d815260200180611853602d913960400191505060405180910390fd5b60078054600160481b600160e81b03191690556108de6001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000205461106490919063ffffffff16565b82356001600160a01b03166000908152600860208181526040832093909355610916928401356001600160801b0316918560016108a1565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b6002546003546004546005546006546007546001600160a01b03928316926001600160801b0383169267ffffffffffffffff600160801b8204811693600160c01b909204169163ffffffff8083169264010000000081049091169160ff600160401b83041691600160481b9004168b565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__639774ec0960028c8c8c8b8b8b8b8b8b6040518b63ffffffff1660e01b8152600401808b81526020018a815260200189600260200280838360005b83811015610a145781810151838201526020016109fc565b50505050905001888152602001878152602001861515151581526020018563ffffffff1663ffffffff1681526020018467ffffffffffffffff1667ffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610a9c578181015183820152602001610a84565b50505050905090810190601f168015610ac95780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060006040518083038186803b158015610aef57600080fd5b505af4158015610b03573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610b3e935091506001600160801b031663ffffffff61106416565b6005546001600160a01b0316600090815260086020526040902055610b62826110c5565b50505050505050505050565b6000546001600160a01b031615610bc5576040805162461bcd60e51b8152602060048201526016602482015275159348185b1c9958591e481a5b9a5d1a585b1a5e995960521b604482015290519081900360640190fd5b6001600160a01b038216610c0a5760405162461bcd60e51b81526004018080602001828103825260218152602001806118326021913960400191505060405180910390fd5b600180546001600160a01b038084166001600160a01b031992831617928390556000805486831693169290921782556040805163f397238360e01b81529051939091169263f39723839260048084019391929182900301818387803b158015610c7257600080fd5b505af1158015610c86573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b1916600160401b83021790555073__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b158015610d0857600080fd5b505af4158015610d1c573d6000803e3d6000fd5b505050506040513d6020811015610d3257600080fd5b50516004555050600680546fffffffffffffffffffffffffffffffff19166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161767ffffffff000000001916640100000000929093169190910291909117905550565b6000546001600160a01b031681565b600a546001600160a01b031681565b6009546001600160a01b031681565b336000908152600860205260409020546006546001600160801b03161115610e205760405162461bcd60e51b81526004018080602001828103825260318152602001806118806031913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151633169230760e01b8152600260048201818152602483018e9052604483018d9052606483018c90528a1515608484015260a483018a905260c4830189905263ffffffff881660e484015267ffffffffffffffff871661010484015273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__9463316923079492938f938f938f938f938f938f938f938f938f936101249091019184918190849084905b83811015610efb578181015183820152602001610ee3565b505050509050019a505050505050505050505060006040518083038186803b158015610f2657600080fd5b505af4158015610f3a573d6000803e3d6000fd5b50505050505050505050505050565b600b546001600160a01b031681565b600b546001600160a01b03163314610fb0576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff166003811115610fca57fe5b1415610fe4576007805460ff60401b1916600160401b1790555b565b6006546001600160801b031690565b600b546001600160a01b0316331461104d576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b610fe46112c3565b6001546001600160a01b031681565b6000828201838110156110be576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b15801561110b57600080fd5b505af115801561111f573d6000803e3d6000fd5b505050506040513d602081101561113557600080fd5b50516040805163364df27760e01b8152905191925073__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__9163364df27791600480820192602092909190829003018186803b15801561118657600080fd5b505af415801561119a573d6000803e3d6000fd5b505050506040513d60208110156111b057600080fd5b505181146111fb576111f760405180606001604052806111d060016112d1565b81526020016111e2600280015461134f565b81526020016111f08461134f565b90526113cd565b6004555b60015460405163e4eb8c6360e01b81526020600482018181528551602484015285516001600160a01b039094169363e4eb8c6393879383926044909201919085019080838360005b8381101561125b578181015183820152602001611243565b50505050905090810190601f1680156112885780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156112a757600080fd5b505af11580156112bb573d6000803e3d6000fd5b505050505050565b600b546001600160a01b0316ff5b6112d96117a3565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161133e565b61132b6117a3565b8152602001906001900390816113235790505b508152600060209091015292915050565b6113576117a3565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916113bc565b6113a96117a3565b8152602001906001900390816113a15790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b6113f16117a3565b8152602001906001900390816113e9575050805190915060005b818110156114435784816003811061141f57fe5b602002015183828151811061143057fe5b602090810291909101015260010161140b565b5061144d82611455565b949350505050565b60006008825111156114a5576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156114d2578160200160208202803883390190505b50805190915060005b8181101561152e576114eb6117d1565b6115078683815181106114fa57fe5b60200260200101516115a1565b9050806000015184838151811061151a57fe5b6020908102919091010152506001016114db565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561157757818101518382015260200161155f565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6115a96117d1565b6060820151600c60ff909116106115fb576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661162857604051806020016040528061161f84600001516116d7565b90529050610504565b606082015160ff166001141561166f57604051806020016040528061161f8460200151600001518560200151604001518660200151606001518760200151602001516116fb565b606082015160ff16600214156116945750604080516020810190915281518152610504565b600360ff16826060015160ff16101580156116b857506060820151600c60ff909116105b156116d557604051806020016040528061161f8460400151611455565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b60008315611755575060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602282018590526042808301859052835180840390910181526062909201909252805191012061144d565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6040518060800160405280600081526020016117bd6117e3565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f7765644368616c6c656e676520666163746f72792061646472657373206e6f74207365744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274a265627a7a7231582075c5fde4f0019aaff8246f2f24dce5fc380c28f34a89f22eb42602382ef73e5f64736f6c634300050f0032"

// DeployArbBase deploys a new Ethereum contract, binding an instance of ArbBase to it.
func DeployArbBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbBase, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ArbBaseBin = strings.Replace(ArbBaseBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	valueAddr, _, _, _ := DeployValue(auth, backend)
	ArbBaseBin = strings.Replace(ArbBaseBin, "__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__", valueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbBase{ArbBaseCaller: ArbBaseCaller{contract: contract}, ArbBaseTransactor: ArbBaseTransactor{contract: contract}, ArbBaseFilterer: ArbBaseFilterer{contract: contract}}, nil
}

// ArbBase is an auto generated Go binding around an Ethereum contract.
type ArbBase struct {
	ArbBaseCaller     // Read-only binding to the contract
	ArbBaseTransactor // Write-only binding to the contract
	ArbBaseFilterer   // Log filterer for contract events
}

// ArbBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbBaseSession struct {
	Contract     *ArbBase          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbBaseCallerSession struct {
	Contract *ArbBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ArbBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbBaseTransactorSession struct {
	Contract     *ArbBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ArbBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbBaseRaw struct {
	Contract *ArbBase // Generic contract binding to access the raw methods on
}

// ArbBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbBaseCallerRaw struct {
	Contract *ArbBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ArbBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbBaseTransactorRaw struct {
	Contract *ArbBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbBase creates a new instance of ArbBase, bound to a specific deployed contract.
func NewArbBase(address common.Address, backend bind.ContractBackend) (*ArbBase, error) {
	contract, err := bindArbBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbBase{ArbBaseCaller: ArbBaseCaller{contract: contract}, ArbBaseTransactor: ArbBaseTransactor{contract: contract}, ArbBaseFilterer: ArbBaseFilterer{contract: contract}}, nil
}

// NewArbBaseCaller creates a new read-only instance of ArbBase, bound to a specific deployed contract.
func NewArbBaseCaller(address common.Address, caller bind.ContractCaller) (*ArbBaseCaller, error) {
	contract, err := bindArbBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbBaseCaller{contract: contract}, nil
}

// NewArbBaseTransactor creates a new write-only instance of ArbBase, bound to a specific deployed contract.
func NewArbBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbBaseTransactor, error) {
	contract, err := bindArbBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbBaseTransactor{contract: contract}, nil
}

// NewArbBaseFilterer creates a new log filterer instance of ArbBase, bound to a specific deployed contract.
func NewArbBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbBaseFilterer, error) {
	contract, err := bindArbBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbBaseFilterer{contract: contract}, nil
}

// bindArbBase binds a generic wrapper to an already deployed contract.
func bindArbBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbBase *ArbBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbBase.Contract.ArbBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbBase *ArbBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBase.Contract.ArbBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbBase *ArbBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbBase.Contract.ArbBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbBase *ArbBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbBase *ArbBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbBase *ArbBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbBase.Contract.contract.Transact(opts, method, params...)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() constant returns(address)
func (_ArbBase *ArbBaseCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbBase.contract.Call(opts, out, "challengeFactory")
	return *ret0, err
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() constant returns(address)
func (_ArbBase *ArbBaseSession) ChallengeFactory() (common.Address, error) {
	return _ArbBase.Contract.ChallengeFactory(&_ArbBase.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() constant returns(address)
func (_ArbBase *ArbBaseCallerSession) ChallengeFactory() (common.Address, error) {
	return _ArbBase.Contract.ChallengeFactory(&_ArbBase.CallOpts)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbBase *ArbBaseCaller) CurrentDeposit(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbBase.contract.Call(opts, out, "currentDeposit", validator)
	return *ret0, err
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbBase *ArbBaseSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _ArbBase.Contract.CurrentDeposit(&_ArbBase.CallOpts, validator)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbBase *ArbBaseCallerSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _ArbBase.Contract.CurrentDeposit(&_ArbBase.CallOpts, validator)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbBase *ArbBaseCaller) EscrowRequired(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbBase.contract.Call(opts, out, "escrowRequired")
	return *ret0, err
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbBase *ArbBaseSession) EscrowRequired() (*big.Int, error) {
	return _ArbBase.Contract.EscrowRequired(&_ArbBase.CallOpts)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbBase *ArbBaseCallerSession) EscrowRequired() (*big.Int, error) {
	return _ArbBase.Contract.EscrowRequired(&_ArbBase.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbBase *ArbBaseCaller) ExitAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbBase.contract.Call(opts, out, "exitAddress")
	return *ret0, err
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbBase *ArbBaseSession) ExitAddress() (common.Address, error) {
	return _ArbBase.Contract.ExitAddress(&_ArbBase.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbBase *ArbBaseCallerSession) ExitAddress() (common.Address, error) {
	return _ArbBase.Contract.ExitAddress(&_ArbBase.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbBase *ArbBaseCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ArbBase.contract.Call(opts, out, "getState")
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbBase *ArbBaseSession) GetState() (uint8, error) {
	return _ArbBase.Contract.GetState(&_ArbBase.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbBase *ArbBaseCallerSession) GetState() (uint8, error) {
	return _ArbBase.Contract.GetState(&_ArbBase.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbBase *ArbBaseCaller) GlobalInbox(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbBase.contract.Call(opts, out, "globalInbox")
	return *ret0, err
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbBase *ArbBaseSession) GlobalInbox() (common.Address, error) {
	return _ArbBase.Contract.GlobalInbox(&_ArbBase.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbBase *ArbBaseCallerSession) GlobalInbox() (common.Address, error) {
	return _ArbBase.Contract.GlobalInbox(&_ArbBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbBase *ArbBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbBase.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbBase *ArbBaseSession) Owner() (common.Address, error) {
	return _ArbBase.Contract.Owner(&_ArbBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbBase *ArbBaseCallerSession) Owner() (common.Address, error) {
	return _ArbBase.Contract.Owner(&_ArbBase.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbBase *ArbBaseCaller) TerminateAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbBase.contract.Call(opts, out, "terminateAddress")
	return *ret0, err
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbBase *ArbBaseSession) TerminateAddress() (common.Address, error) {
	return _ArbBase.Contract.TerminateAddress(&_ArbBase.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbBase *ArbBaseCallerSession) TerminateAddress() (common.Address, error) {
	return _ArbBase.Contract.TerminateAddress(&_ArbBase.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbBase *ArbBaseCaller) Vm(opts *bind.CallOpts) (struct {
	MachineHash            [32]byte
	PendingHash            [32]byte
	Inbox                  [32]byte
	Asserter               common.Address
	EscrowRequired         *big.Int
	Deadline               uint64
	SequenceNum            uint64
	GracePeriod            uint32
	MaxExecutionSteps      uint32
	State                  uint8
	ActiveChallengeManager common.Address
}, error) {
	ret := new(struct {
		MachineHash            [32]byte
		PendingHash            [32]byte
		Inbox                  [32]byte
		Asserter               common.Address
		EscrowRequired         *big.Int
		Deadline               uint64
		SequenceNum            uint64
		GracePeriod            uint32
		MaxExecutionSteps      uint32
		State                  uint8
		ActiveChallengeManager common.Address
	})
	out := ret
	err := _ArbBase.contract.Call(opts, out, "vm")
	return *ret, err
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbBase *ArbBaseSession) Vm() (struct {
	MachineHash            [32]byte
	PendingHash            [32]byte
	Inbox                  [32]byte
	Asserter               common.Address
	EscrowRequired         *big.Int
	Deadline               uint64
	SequenceNum            uint64
	GracePeriod            uint32
	MaxExecutionSteps      uint32
	State                  uint8
	ActiveChallengeManager common.Address
}, error) {
	return _ArbBase.Contract.Vm(&_ArbBase.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbBase *ArbBaseCallerSession) Vm() (struct {
	MachineHash            [32]byte
	PendingHash            [32]byte
	Inbox                  [32]byte
	Asserter               common.Address
	EscrowRequired         *big.Int
	Deadline               uint64
	SequenceNum            uint64
	GracePeriod            uint32
	MaxExecutionSteps      uint32
	State                  uint8
	ActiveChallengeManager common.Address
}, error) {
	return _ArbBase.Contract.Vm(&_ArbBase.CallOpts)
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbBase *ArbBaseTransactor) ActivateVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBase.contract.Transact(opts, "activateVM")
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbBase *ArbBaseSession) ActivateVM() (*types.Transaction, error) {
	return _ArbBase.Contract.ActivateVM(&_ArbBase.TransactOpts)
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbBase *ArbBaseTransactorSession) ActivateVM() (*types.Transaction, error) {
	return _ArbBase.Contract.ActivateVM(&_ArbBase.TransactOpts)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbBase *ArbBaseTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbBase.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbBase *ArbBaseSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbBase.Contract.CompleteChallenge(&_ArbBase.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbBase *ArbBaseTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbBase.Contract.CompleteChallenge(&_ArbBase.TransactOpts, _players, _rewards)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4676a75e.
//
// Solidity: function confirmDisputableAsserted(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes32 _preconditionHash, bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _numGas, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbBase *ArbBaseTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _numGas uint64, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbBase.contract.Transact(opts, "confirmDisputableAsserted", _beforeHash, _timeBounds, _beforeInbox, _preconditionHash, _afterHash, _didInboxInsn, _numSteps, _numGas, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4676a75e.
//
// Solidity: function confirmDisputableAsserted(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes32 _preconditionHash, bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _numGas, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbBase *ArbBaseSession) ConfirmDisputableAsserted(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _numGas uint64, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbBase.Contract.ConfirmDisputableAsserted(&_ArbBase.TransactOpts, _beforeHash, _timeBounds, _beforeInbox, _preconditionHash, _afterHash, _didInboxInsn, _numSteps, _numGas, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4676a75e.
//
// Solidity: function confirmDisputableAsserted(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes32 _preconditionHash, bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _numGas, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbBase *ArbBaseTransactorSession) ConfirmDisputableAsserted(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _numGas uint64, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbBase.Contract.ConfirmDisputableAsserted(&_ArbBase.TransactOpts, _beforeHash, _timeBounds, _beforeInbox, _preconditionHash, _afterHash, _didInboxInsn, _numSteps, _numGas, _messages, _logsAccHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cbb9ff2.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_ArbBase *ArbBaseTransactor) Initialize(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbBase.contract.Transact(opts, "initialize", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cbb9ff2.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_ArbBase *ArbBaseSession) Initialize(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbBase.Contract.Initialize(&_ArbBase.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cbb9ff2.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_ArbBase *ArbBaseTransactorSession) Initialize(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbBase.Contract.Initialize(&_ArbBase.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x0badcbbf.
//
// Solidity: function initiateChallenge(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbBase *ArbBaseTransactor) InitiateChallenge(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbBase.contract.Transact(opts, "initiateChallenge", _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x0badcbbf.
//
// Solidity: function initiateChallenge(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbBase *ArbBaseSession) InitiateChallenge(_beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbBase.Contract.InitiateChallenge(&_ArbBase.TransactOpts, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x0badcbbf.
//
// Solidity: function initiateChallenge(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbBase *ArbBaseTransactorSession) InitiateChallenge(_beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbBase.Contract.InitiateChallenge(&_ArbBase.TransactOpts, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbBase *ArbBaseTransactor) OwnerShutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBase.contract.Transact(opts, "ownerShutdown")
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbBase *ArbBaseSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbBase.Contract.OwnerShutdown(&_ArbBase.TransactOpts)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbBase *ArbBaseTransactorSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbBase.Contract.OwnerShutdown(&_ArbBase.TransactOpts)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x6d45809f.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bool _didInboxInsn, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numGas, uint64[2] _timeBounds) returns()
func (_ArbBase *ArbBaseTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _didInboxInsn bool, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbBase.contract.Transact(opts, "pendingDisputableAssert", _beforeHash, _beforeInbox, _afterHash, _didInboxInsn, _messagesAccHash, _logsAccHash, _numSteps, _numGas, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x6d45809f.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bool _didInboxInsn, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numGas, uint64[2] _timeBounds) returns()
func (_ArbBase *ArbBaseSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _didInboxInsn bool, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbBase.Contract.PendingDisputableAssert(&_ArbBase.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _didInboxInsn, _messagesAccHash, _logsAccHash, _numSteps, _numGas, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x6d45809f.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bool _didInboxInsn, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numGas, uint64[2] _timeBounds) returns()
func (_ArbBase *ArbBaseTransactorSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _didInboxInsn bool, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbBase.Contract.PendingDisputableAssert(&_ArbBase.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _didInboxInsn, _messagesAccHash, _logsAccHash, _numSteps, _numGas, _timeBounds)
}

// ArbBaseChallengeLaunchedIterator is returned from FilterChallengeLaunched and is used to iterate over the raw logs and unpacked data for ChallengeLaunched events raised by the ArbBase contract.
type ArbBaseChallengeLaunchedIterator struct {
	Event *ArbBaseChallengeLaunched // Event containing the contract specifics and raw log

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
func (it *ArbBaseChallengeLaunchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbBaseChallengeLaunched)
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
		it.Event = new(ArbBaseChallengeLaunched)
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
func (it *ArbBaseChallengeLaunchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbBaseChallengeLaunchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbBaseChallengeLaunched represents a ChallengeLaunched event raised by the ArbBase contract.
type ArbBaseChallengeLaunched struct {
	ChallengeContract common.Address
	Challenger        common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChallengeLaunched is a free log retrieval operation binding the contract event 0x65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f.
//
// Solidity: event ChallengeLaunched(address challengeContract, address challenger)
func (_ArbBase *ArbBaseFilterer) FilterChallengeLaunched(opts *bind.FilterOpts) (*ArbBaseChallengeLaunchedIterator, error) {

	logs, sub, err := _ArbBase.contract.FilterLogs(opts, "ChallengeLaunched")
	if err != nil {
		return nil, err
	}
	return &ArbBaseChallengeLaunchedIterator{contract: _ArbBase.contract, event: "ChallengeLaunched", logs: logs, sub: sub}, nil
}

// WatchChallengeLaunched is a free log subscription operation binding the contract event 0x65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f.
//
// Solidity: event ChallengeLaunched(address challengeContract, address challenger)
func (_ArbBase *ArbBaseFilterer) WatchChallengeLaunched(opts *bind.WatchOpts, sink chan<- *ArbBaseChallengeLaunched) (event.Subscription, error) {

	logs, sub, err := _ArbBase.contract.WatchLogs(opts, "ChallengeLaunched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbBaseChallengeLaunched)
				if err := _ArbBase.contract.UnpackLog(event, "ChallengeLaunched", log); err != nil {
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

// ParseChallengeLaunched is a log parse operation binding the contract event 0x65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f.
//
// Solidity: event ChallengeLaunched(address challengeContract, address challenger)
func (_ArbBase *ArbBaseFilterer) ParseChallengeLaunched(log types.Log) (*ArbBaseChallengeLaunched, error) {
	event := new(ArbBaseChallengeLaunched)
	if err := _ArbBase.contract.UnpackLog(event, "ChallengeLaunched", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbBaseConfirmedDisputableAssertionIterator is returned from FilterConfirmedDisputableAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedDisputableAssertion events raised by the ArbBase contract.
type ArbBaseConfirmedDisputableAssertionIterator struct {
	Event *ArbBaseConfirmedDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *ArbBaseConfirmedDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbBaseConfirmedDisputableAssertion)
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
		it.Event = new(ArbBaseConfirmedDisputableAssertion)
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
func (it *ArbBaseConfirmedDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbBaseConfirmedDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbBaseConfirmedDisputableAssertion represents a ConfirmedDisputableAssertion event raised by the ArbBase contract.
type ArbBaseConfirmedDisputableAssertion struct {
	NewState    [32]byte
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbBase *ArbBaseFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts) (*ArbBaseConfirmedDisputableAssertionIterator, error) {

	logs, sub, err := _ArbBase.contract.FilterLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbBaseConfirmedDisputableAssertionIterator{contract: _ArbBase.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbBase *ArbBaseFilterer) WatchConfirmedDisputableAssertion(opts *bind.WatchOpts, sink chan<- *ArbBaseConfirmedDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbBase.contract.WatchLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbBaseConfirmedDisputableAssertion)
				if err := _ArbBase.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
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
func (_ArbBase *ArbBaseFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*ArbBaseConfirmedDisputableAssertion, error) {
	event := new(ArbBaseConfirmedDisputableAssertion)
	if err := _ArbBase.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbBasePendingDisputableAssertionIterator is returned from FilterPendingDisputableAssertion and is used to iterate over the raw logs and unpacked data for PendingDisputableAssertion events raised by the ArbBase contract.
type ArbBasePendingDisputableAssertionIterator struct {
	Event *ArbBasePendingDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *ArbBasePendingDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbBasePendingDisputableAssertion)
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
		it.Event = new(ArbBasePendingDisputableAssertion)
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
func (it *ArbBasePendingDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbBasePendingDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbBasePendingDisputableAssertion represents a PendingDisputableAssertion event raised by the ArbBase contract.
type ArbBasePendingDisputableAssertion struct {
	Fields     [5][32]byte
	Asserter   common.Address
	TimeBounds [2]uint64
	NumSteps   uint32
	NumGas     uint64
	Deadline   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0xa2ec50303fe3da5e6c44070c162829210df70e10f193ecdcc7fff65776dfa539.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 numGas, uint64 deadline)
func (_ArbBase *ArbBaseFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*ArbBasePendingDisputableAssertionIterator, error) {

	logs, sub, err := _ArbBase.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbBasePendingDisputableAssertionIterator{contract: _ArbBase.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0xa2ec50303fe3da5e6c44070c162829210df70e10f193ecdcc7fff65776dfa539.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 numGas, uint64 deadline)
func (_ArbBase *ArbBaseFilterer) WatchPendingDisputableAssertion(opts *bind.WatchOpts, sink chan<- *ArbBasePendingDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbBase.contract.WatchLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbBasePendingDisputableAssertion)
				if err := _ArbBase.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0xa2ec50303fe3da5e6c44070c162829210df70e10f193ecdcc7fff65776dfa539.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 numGas, uint64 deadline)
func (_ArbBase *ArbBaseFilterer) ParsePendingDisputableAssertion(log types.Log) (*ArbBasePendingDisputableAssertion, error) {
	event := new(ArbBasePendingDisputableAssertion)
	if err := _ArbBase.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChainABI is the input ABI used to generate the binding from.
const ArbChainABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"ChallengeLaunched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numGas\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumVM.State\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeLauncherAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_numGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"enumVM.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"activeChallengeManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbChainFuncSigs maps the 4-byte function signature to its string representation.
var ArbChainFuncSigs = map[string]string{
	"94af716b": "activateVM()",
	"5dbaf68b": "challengeFactory()",
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"4676a75e": "confirmDisputableAsserted(bytes32,uint64[2],bytes32,bytes32,bytes32,bool,uint32,uint64,bytes,bytes32)",
	"08dc89d7": "currentDeposit(address)",
	"aca0f372": "escrowRequired()",
	"6be00229": "exitAddress()",
	"1865c57d": "getState()",
	"d489113a": "globalInbox()",
	"05b050de": "increaseDeposit()",
	"8364fe47": "init(bytes32,uint32,uint32,uint128,address,address,address)",
	"4cbb9ff2": "initialize(bytes32,uint32,uint32,uint128,address,address,address)",
	"0badcbbf": "initiateChallenge(bytes32,bytes32,uint64[2],bytes32)",
	"8da5cb5b": "owner()",
	"cfa80707": "ownerShutdown()",
	"6d45809f": "pendingDisputableAssert(bytes32,bytes32,bytes32,bool,bytes32,bytes32,uint32,uint64,uint64[2])",
	"60675a87": "terminateAddress()",
	"3a768463": "vm()",
}

// ArbChainBin is the compiled bytecode used for deploying new contracts.
var ArbChainBin = "0x608060405234801561001057600080fd5b50611a78806100206000396000f3fe6080604052600436106101095760003560e01c806360675a87116100955780638da5cb5b116100645780638da5cb5b146105d257806394af716b146105e7578063aca0f372146105fc578063cfa8070714610611578063d489113a1461062657610109565b806360675a871461049b5780636be00229146104b05780636d45809f146104c55780638364fe471461056257610109565b806322c091bc116100dc57806322c091bc146101fc5780633a768463146102295780634676a75e146102db5780634cbb9ff2146103fa5780635dbaf68b1461046a57610109565b806305b050de1461010e57806308dc89d7146101185780630badcbbf1461015d5780631865c57d146101c3575b600080fd5b61011661063b565b005b34801561012457600080fd5b5061014b6004803603602081101561013b57600080fd5b50356001600160a01b0316610652565b60408051918252519081900360200190f35b34801561016957600080fd5b50610116600480360360a081101561018057600080fd5b6040805180820182528335936020810135938101929091608083019180840190600290839083908082843760009201919091525091945050903591506106719050565b3480156101cf57600080fd5b506101d8610981565b604051808260038111156101e857fe5b60ff16815260200191505060405180910390f35b34801561020857600080fd5b506101166004803603608081101561021f57600080fd5b5060408101610991565b34801561023557600080fd5b5061023e610aa3565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e0830152841661010082015261012081018360038111156102ae57fe5b60ff1681526001600160a01b039092166020830152506040805191829003019a5098505050505050505050f35b3480156102e757600080fd5b5061011660048036036101608110156102ff57600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091948335946020850135946040810135945060608101351515935063ffffffff6080820135169267ffffffffffffffff60a08301351692909160e081019060c0013564010000000081111561038357600080fd5b82018360208201111561039557600080fd5b803590602001918460018302840111640100000000831117156103b757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610b14915050565b34801561040657600080fd5b50610116600480360360e081101561041d57600080fd5b5080359063ffffffff60208201358116916040810135909116906001600160801b03606082013516906001600160a01b03608082013581169160a081013582169160c09091013516610cd6565b34801561047657600080fd5b5061047f610f02565b604080516001600160a01b039092168252519081900360200190f35b3480156104a757600080fd5b5061047f610f11565b3480156104bc57600080fd5b5061047f610f20565b3480156104d157600080fd5b5061011660048036036101408110156104e957600080fd5b6040805180820182528335936020810135938382013593606083013515159360808401359360a08101359363ffffffff60c0830135169367ffffffffffffffff60e0840135169391830192916101408301916101008401906002908390839080828437600092019190915250919450610f2f9350505050565b34801561056e57600080fd5b50610116600480360360e081101561058557600080fd5b5080359063ffffffff60208201358116916040810135909116906001600160801b03606082013516906001600160a01b03608082013581169160a081013582169160c090910135166110b1565b3480156105de57600080fd5b5061047f6110dc565b3480156105f357600080fd5b506101166110eb565b34801561060857600080fd5b5061014b611179565b34801561061d57600080fd5b50610116611188565b34801561063257600080fd5b5061047f6111e8565b336000908152600860205260409020805434019055565b6001600160a01b0381166000908152600860205260409020545b919050565b336000908152600860205260409020546006546001600160801b031611156106ca5760405162461bcd60e51b815260040180806020018281038252602781526020018061199e6027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b0390941690930390925581516337d8913360e01b8152600260048201818152602483018990526044830188905273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__946337d891339492938a938a938a938a9391926064909201918591908190849084905b8381101561076357818101518382015260200161074b565b505050509050018281526020019550505050505060006040518083038186803b15801561078f57600080fd5b505af41580156107a3573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b0316808252918101919091526007548451631a6ef3c360e31b815292909516975063d3779e1896509194919363ffffffff16928b928b928b928b9260049092019182918a918190849084905b83811015610841578181015183820152602001610829565b5050505090500187600260200280838360005b8381101561086c578181015183820152602001610854565b505050509050018663ffffffff1663ffffffff16815260200185815260200184815260200183600260200280838360005b838110156108b557818101518382015260200161089d565b50505050905001828152602001975050505050505050602060405180830381600087803b1580156108e557600080fd5b505af11580156108f9573d6000803e3d6000fd5b505050506040513d602081101561090f57600080fd5b505160078054600160481b600160e81b031916600160481b6001600160a01b03938416810291909117918290556040805191909204909216825233602083015280517f65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f9281900390910190a150505050565b600754600160401b900460ff1690565b600754600160481b90046001600160a01b031633146109e15760405162461bcd60e51b815260040180806020018281038252602d8152602001806119e6602d913960400191505060405180910390fd5b60078054600160481b600160e81b0319169055610a466001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b03168152602001908152602001600020546111f790919063ffffffff16565b82356001600160a01b03166000908152600860208181526040832093909355610a7e928401356001600160801b031691856001610a09565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b6002546003546004546005546006546007546001600160a01b03928316926001600160801b0383169267ffffffffffffffff600160801b8204811693600160c01b909204169163ffffffff8083169264010000000081049091169160ff600160401b83041691600160481b9004168b565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__639774ec0960028c8c8c8b8b8b8b8b8b6040518b63ffffffff1660e01b8152600401808b81526020018a815260200189600260200280838360005b83811015610b7c578181015183820152602001610b64565b50505050905001888152602001878152602001861515151581526020018563ffffffff1663ffffffff1681526020018467ffffffffffffffff1667ffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610c04578181015183820152602001610bec565b50505050905090810190601f168015610c315780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060006040518083038186803b158015610c5757600080fd5b505af4158015610c6b573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610ca6935091506001600160801b031663ffffffff6111f716565b6005546001600160a01b0316600090815260086020526040902055610cca82611258565b50505050505050505050565b6000546001600160a01b031615610d2d576040805162461bcd60e51b8152602060048201526016602482015275159348185b1c9958591e481a5b9a5d1a585b1a5e995960521b604482015290519081900360640190fd5b6001600160a01b038216610d725760405162461bcd60e51b81526004018080602001828103825260218152602001806119c56021913960400191505060405180910390fd5b600180546001600160a01b038084166001600160a01b031992831617928390556000805486831693169290921782556040805163f397238360e01b81529051939091169263f39723839260048084019391929182900301818387803b158015610dda57600080fd5b505af1158015610dee573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b1916600160401b83021790555073__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b158015610e7057600080fd5b505af4158015610e84573d6000803e3d6000fd5b505050506040513d6020811015610e9a57600080fd5b50516004555050600680546fffffffffffffffffffffffffffffffff19166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161767ffffffff000000001916640100000000929093169190910291909117905550565b6000546001600160a01b031681565b600a546001600160a01b031681565b6009546001600160a01b031681565b336000908152600860205260409020546006546001600160801b03161115610f885760405162461bcd60e51b8152600401808060200182810382526031815260200180611a136031913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151633169230760e01b8152600260048201818152602483018e9052604483018d9052606483018c90528a1515608484015260a483018a905260c4830189905263ffffffff881660e484015267ffffffffffffffff871661010484015273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__9463316923079492938f938f938f938f938f938f938f938f938f936101249091019184918190849084905b8381101561106357818101518382015260200161104b565b505050509050019a505050505050505050505060006040518083038186803b15801561108e57600080fd5b505af41580156110a2573d6000803e3d6000fd5b50505050505050505050505050565b6110c087878787878787610cd6565b50506007805460ff60401b1916600160401b1790555050505050565b600b546001600160a01b031681565b600b546001600160a01b03163314611143576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff16600381111561115d57fe5b1415611177576007805460ff60401b1916600160401b1790555b565b6006546001600160801b031690565b600b546001600160a01b031633146111e0576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b611177611456565b6001546001600160a01b031681565b600082820183811015611251576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b15801561129e57600080fd5b505af11580156112b2573d6000803e3d6000fd5b505050506040513d60208110156112c857600080fd5b50516040805163364df27760e01b8152905191925073__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__9163364df27791600480820192602092909190829003018186803b15801561131957600080fd5b505af415801561132d573d6000803e3d6000fd5b505050506040513d602081101561134357600080fd5b5051811461138e5761138a60405180606001604052806113636001611464565b815260200161137560028001546114e2565b8152602001611383846114e2565b9052611560565b6004555b60015460405163e4eb8c6360e01b81526020600482018181528551602484015285516001600160a01b039094169363e4eb8c6393879383926044909201919085019080838360005b838110156113ee5781810151838201526020016113d6565b50505050905090810190601f16801561141b5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561143a57600080fd5b505af115801561144e573d6000803e3d6000fd5b505050505050565b600b546001600160a01b0316ff5b61146c611936565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916114d1565b6114be611936565b8152602001906001900390816114b65790505b508152600060209091015292915050565b6114ea611936565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161154f565b61153c611936565b8152602001906001900390816115345790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b611584611936565b81526020019060019003908161157c575050805190915060005b818110156115d6578481600381106115b257fe5b60200201518382815181106115c357fe5b602090810291909101015260010161159e565b506115e0826115e8565b949350505050565b6000600882511115611638576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611665578160200160208202803883390190505b50805190915060005b818110156116c15761167e611964565b61169a86838151811061168d57fe5b6020026020010151611734565b905080600001518483815181106116ad57fe5b60209081029190910101525060010161166e565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561170a5781810151838201526020016116f2565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b61173c611964565b6060820151600c60ff9091161061178e576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166117bb5760405180602001604052806117b2846000015161186a565b9052905061066c565b606082015160ff16600114156118025760405180602001604052806117b284602001516000015185602001516040015186602001516060015187602001516020015161188e565b606082015160ff1660021415611827575060408051602081019091528151815261066c565b600360ff16826060015160ff161015801561184b57506060820151600c60ff909116105b156118685760405180602001604052806117b284604001516115e8565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b600083156118e8575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206115e0565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b604051806080016040528060008152602001611950611976565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f7765644368616c6c656e676520666163746f72792061646472657373206e6f74207365744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274a265627a7a72315820eb2816d7fabfabdf76fab7edbe8dd1cec2353a7e61f1f9ec8591f82475eb2b7464736f6c634300050f0032"

// DeployArbChain deploys a new Ethereum contract, binding an instance of ArbChain to it.
func DeployArbChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbChain, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ArbChainBin = strings.Replace(ArbChainBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	valueAddr, _, _, _ := DeployValue(auth, backend)
	ArbChainBin = strings.Replace(ArbChainBin, "__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__", valueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbChainBin), backend)
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

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() constant returns(address)
func (_ArbChain *ArbChainCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "challengeFactory")
	return *ret0, err
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() constant returns(address)
func (_ArbChain *ArbChainSession) ChallengeFactory() (common.Address, error) {
	return _ArbChain.Contract.ChallengeFactory(&_ArbChain.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() constant returns(address)
func (_ArbChain *ArbChainCallerSession) ChallengeFactory() (common.Address, error) {
	return _ArbChain.Contract.ChallengeFactory(&_ArbChain.CallOpts)
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
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbChain *ArbChainCaller) Vm(opts *bind.CallOpts) (struct {
	MachineHash            [32]byte
	PendingHash            [32]byte
	Inbox                  [32]byte
	Asserter               common.Address
	EscrowRequired         *big.Int
	Deadline               uint64
	SequenceNum            uint64
	GracePeriod            uint32
	MaxExecutionSteps      uint32
	State                  uint8
	ActiveChallengeManager common.Address
}, error) {
	ret := new(struct {
		MachineHash            [32]byte
		PendingHash            [32]byte
		Inbox                  [32]byte
		Asserter               common.Address
		EscrowRequired         *big.Int
		Deadline               uint64
		SequenceNum            uint64
		GracePeriod            uint32
		MaxExecutionSteps      uint32
		State                  uint8
		ActiveChallengeManager common.Address
	})
	out := ret
	err := _ArbChain.contract.Call(opts, out, "vm")
	return *ret, err
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbChain *ArbChainSession) Vm() (struct {
	MachineHash            [32]byte
	PendingHash            [32]byte
	Inbox                  [32]byte
	Asserter               common.Address
	EscrowRequired         *big.Int
	Deadline               uint64
	SequenceNum            uint64
	GracePeriod            uint32
	MaxExecutionSteps      uint32
	State                  uint8
	ActiveChallengeManager common.Address
}, error) {
	return _ArbChain.Contract.Vm(&_ArbChain.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbChain *ArbChainCallerSession) Vm() (struct {
	MachineHash            [32]byte
	PendingHash            [32]byte
	Inbox                  [32]byte
	Asserter               common.Address
	EscrowRequired         *big.Int
	Deadline               uint64
	SequenceNum            uint64
	GracePeriod            uint32
	MaxExecutionSteps      uint32
	State                  uint8
	ActiveChallengeManager common.Address
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

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4676a75e.
//
// Solidity: function confirmDisputableAsserted(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes32 _preconditionHash, bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _numGas, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _numGas uint64, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "confirmDisputableAsserted", _beforeHash, _timeBounds, _beforeInbox, _preconditionHash, _afterHash, _didInboxInsn, _numSteps, _numGas, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4676a75e.
//
// Solidity: function confirmDisputableAsserted(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes32 _preconditionHash, bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _numGas, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainSession) ConfirmDisputableAsserted(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _numGas uint64, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.ConfirmDisputableAsserted(&_ArbChain.TransactOpts, _beforeHash, _timeBounds, _beforeInbox, _preconditionHash, _afterHash, _didInboxInsn, _numSteps, _numGas, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4676a75e.
//
// Solidity: function confirmDisputableAsserted(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes32 _preconditionHash, bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _numGas, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainTransactorSession) ConfirmDisputableAsserted(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _preconditionHash [32]byte, _afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _numGas uint64, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.ConfirmDisputableAsserted(&_ArbChain.TransactOpts, _beforeHash, _timeBounds, _beforeInbox, _preconditionHash, _afterHash, _didInboxInsn, _numSteps, _numGas, _messages, _logsAccHash)
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

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_ArbChain *ArbChainTransactor) Init(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "init", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_ArbChain *ArbChainSession) Init(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbChain.Contract.Init(&_ArbChain.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_ArbChain *ArbChainTransactorSession) Init(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbChain.Contract.Init(&_ArbChain.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cbb9ff2.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_ArbChain *ArbChainTransactor) Initialize(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "initialize", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cbb9ff2.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_ArbChain *ArbChainSession) Initialize(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbChain.Contract.Initialize(&_ArbChain.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cbb9ff2.
//
// Solidity: function initialize(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeFactoryAddress, address _globalInboxAddress) returns()
func (_ArbChain *ArbChainTransactorSession) Initialize(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _ArbChain.Contract.Initialize(&_ArbChain.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeFactoryAddress, _globalInboxAddress)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x0badcbbf.
//
// Solidity: function initiateChallenge(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbChain *ArbChainTransactor) InitiateChallenge(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "initiateChallenge", _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x0badcbbf.
//
// Solidity: function initiateChallenge(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbChain *ArbChainSession) InitiateChallenge(_beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.InitiateChallenge(&_ArbChain.TransactOpts, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x0badcbbf.
//
// Solidity: function initiateChallenge(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbChain *ArbChainTransactorSession) InitiateChallenge(_beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.InitiateChallenge(&_ArbChain.TransactOpts, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
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

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x6d45809f.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bool _didInboxInsn, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numGas, uint64[2] _timeBounds) returns()
func (_ArbChain *ArbChainTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _didInboxInsn bool, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "pendingDisputableAssert", _beforeHash, _beforeInbox, _afterHash, _didInboxInsn, _messagesAccHash, _logsAccHash, _numSteps, _numGas, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x6d45809f.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bool _didInboxInsn, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numGas, uint64[2] _timeBounds) returns()
func (_ArbChain *ArbChainSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _didInboxInsn bool, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChain.Contract.PendingDisputableAssert(&_ArbChain.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _didInboxInsn, _messagesAccHash, _logsAccHash, _numSteps, _numGas, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x6d45809f.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bool _didInboxInsn, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64 _numGas, uint64[2] _timeBounds) returns()
func (_ArbChain *ArbChainTransactorSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _didInboxInsn bool, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _numGas uint64, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChain.Contract.PendingDisputableAssert(&_ArbChain.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _didInboxInsn, _messagesAccHash, _logsAccHash, _numSteps, _numGas, _timeBounds)
}

// ArbChainChallengeLaunchedIterator is returned from FilterChallengeLaunched and is used to iterate over the raw logs and unpacked data for ChallengeLaunched events raised by the ArbChain contract.
type ArbChainChallengeLaunchedIterator struct {
	Event *ArbChainChallengeLaunched // Event containing the contract specifics and raw log

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
func (it *ArbChainChallengeLaunchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChainChallengeLaunched)
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
		it.Event = new(ArbChainChallengeLaunched)
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
func (it *ArbChainChallengeLaunchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChainChallengeLaunchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChainChallengeLaunched represents a ChallengeLaunched event raised by the ArbChain contract.
type ArbChainChallengeLaunched struct {
	ChallengeContract common.Address
	Challenger        common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChallengeLaunched is a free log retrieval operation binding the contract event 0x65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f.
//
// Solidity: event ChallengeLaunched(address challengeContract, address challenger)
func (_ArbChain *ArbChainFilterer) FilterChallengeLaunched(opts *bind.FilterOpts) (*ArbChainChallengeLaunchedIterator, error) {

	logs, sub, err := _ArbChain.contract.FilterLogs(opts, "ChallengeLaunched")
	if err != nil {
		return nil, err
	}
	return &ArbChainChallengeLaunchedIterator{contract: _ArbChain.contract, event: "ChallengeLaunched", logs: logs, sub: sub}, nil
}

// WatchChallengeLaunched is a free log subscription operation binding the contract event 0x65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f.
//
// Solidity: event ChallengeLaunched(address challengeContract, address challenger)
func (_ArbChain *ArbChainFilterer) WatchChallengeLaunched(opts *bind.WatchOpts, sink chan<- *ArbChainChallengeLaunched) (event.Subscription, error) {

	logs, sub, err := _ArbChain.contract.WatchLogs(opts, "ChallengeLaunched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChainChallengeLaunched)
				if err := _ArbChain.contract.UnpackLog(event, "ChallengeLaunched", log); err != nil {
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

// ParseChallengeLaunched is a log parse operation binding the contract event 0x65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f.
//
// Solidity: event ChallengeLaunched(address challengeContract, address challenger)
func (_ArbChain *ArbChainFilterer) ParseChallengeLaunched(log types.Log) (*ArbChainChallengeLaunched, error) {
	event := new(ArbChainChallengeLaunched)
	if err := _ArbChain.contract.UnpackLog(event, "ChallengeLaunched", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Fields     [5][32]byte
	Asserter   common.Address
	TimeBounds [2]uint64
	NumSteps   uint32
	NumGas     uint64
	Deadline   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0xa2ec50303fe3da5e6c44070c162829210df70e10f193ecdcc7fff65776dfa539.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 numGas, uint64 deadline)
func (_ArbChain *ArbChainFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*ArbChainPendingDisputableAssertionIterator, error) {

	logs, sub, err := _ArbChain.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChainPendingDisputableAssertionIterator{contract: _ArbChain.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0xa2ec50303fe3da5e6c44070c162829210df70e10f193ecdcc7fff65776dfa539.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 numGas, uint64 deadline)
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0xa2ec50303fe3da5e6c44070c162829210df70e10f193ecdcc7fff65776dfa539.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 numGas, uint64 deadline)
func (_ArbChain *ArbChainFilterer) ParsePendingDisputableAssertion(log types.Log) (*ArbChainPendingDisputableAssertion, error) {
	event := new(ArbChainPendingDisputableAssertion)
	if err := _ArbChain.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158204a4406891b86a3dacfb6ea735b92a605972908aec5c506cc187bb6f01ef6a24064736f6c634300050f0032"

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

// DisputableABI is the input ABI used to generate the binding from.
const DisputableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"didInboxInsn\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"didInboxInsn\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numGas\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"withinTimeBounds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DisputableFuncSigs maps the 4-byte function signature to its string representation.
var DisputableFuncSigs = map[string]string{
	"9774ec09": "confirmDisputableAsserted(VM.Data storage,bytes32,uint64[2],bytes32,bytes32,bool,uint32,uint64,bytes,bytes32)",
	"37d89133": "initiateChallenge(VM.Data storage,bytes32,bytes32,uint64[2],bytes32)",
	"31692307": "pendingDisputableAssert(VM.Data storage,bytes32,bytes32,bytes32,bool,bytes32,bytes32,uint32,uint64,uint64[2])",
	"42c0787e": "withinTimeBounds(uint64[2])",
}

// DisputableBin is the compiled bytecode used for deploying new contracts.
var DisputableBin = "0x6113dc610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c8063316923071461005b57806337d891331461010257806342c0787e1461016d5780639774ec09146101cc575b600080fd5b81801561006757600080fd5b50610100600480360361016081101561007f57600080fd5b6040805180820182528335936020810135938382013593606083013593608084013515159360a08101359360c08201359363ffffffff60e0840135169367ffffffffffffffff610100850135169392830192916101608301919061012084019060029083908390808284376000920191909152509194506102e99350505050565b005b81801561010e57600080fd5b50610100600480360360c081101561012557600080fd5b60408051808201825283359360208101359383820135939082019260a08301916060840190600290839083908082843760009201919091525091945050903591506108779050565b6101b86004803603604081101561018357600080fd5b60408051808201825291830192918183019183906002908390839080828437600092019190915250919450610b0c9350505050565b604080519115158252519081900360200190f35b8180156101d857600080fd5b5061010060048036036101608110156101f057600080fd5b60408051808201825283359360208101359381019290916080830191808401906002908390839080828437600092019190915250919483359460208501359460408101351515945063ffffffff606082013516935067ffffffffffffffff60808201351692919060c081019060a0013564010000000081111561027257600080fd5b82018360208201111561028457600080fd5b803590602001918460018302840111640100000000831117156102a657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610b3e915050565b600160058b0154600160401b900460ff16600381111561030557fe5b146103415760405162461bcd60e51b815260040180806020018281038252602d8152602001806112e1602d913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__632a3e0a978b6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561039057600080fd5b505af41580156103a4573d6000803e3d6000fd5b505050506040513d60208110156103ba57600080fd5b5051158015610441575073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca8b6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561041357600080fd5b505af4158015610427573d6000803e3d6000fd5b505050506040513d602081101561043d57600080fd5b5051155b61047c5760405162461bcd60e51b815260040180806020018281038252603e815260200180611234603e913960400191505060405180910390fd5b60058a0154690100000000000000000090046001600160a01b0316156104d35760405162461bcd60e51b815260040180806020018281038252602e81526020018061117b602e913960400191505060405180910390fd5b60058a015463ffffffff6401000000009091048116908416111561053e576040805162461bcd60e51b815260206004820152601f60248201527f547269656420746f206578656375746520746f6f206d616e7920737465707300604482015290519081900360640190fd5b61054781610b0c565b6105825760405162461bcd60e51b81526004018080602001828103825260248152602001806111ef6024913960400191505060405180910390fd5b895489146105c15760405162461bcd60e51b81526004018080602001828103825260278152602001806112946027913960400191505060405180910390fd5b896002015488146106035760405162461bcd60e51b81526004018080602001828103825260228152602001806112726022913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63a3a162cb8b6040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561065257600080fd5b505af4158015610666573d6000803e3d6000fd5b5050505073__$2556963077056ca10a6804584182250fbf$__6385ecb92a8a838b6040518463ffffffff1660e01b81526004018084815260200183600260200280838360005b838110156106c45781810151838201526020016106ac565b50505050905001828152602001935050505060206040518083038186803b1580156106ee57600080fd5b505af4158015610702573d6000803e3d6000fd5b505050506040513d602081101561071857600080fd5b505160408051631b0aa96b60e01b8152600481018a9052881515602482015263ffffffff8616604482015267ffffffffffffffff8516606482015260006084820181905260a4820189905260c482015260e48101879052905173__$2556963077056ca10a6804584182250fbf$__91631b0aa96b91610104808301926020929190829003018186803b1580156107ad57600080fd5b505af41580156107c1573d6000803e3d6000fd5b505050506040513d60208110156107d757600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019052805191012060018b015560038a0180546001600160a01b0319163317905560058a0180546002919060ff60401b1916600160401b83021790555061086b6040518060a001604052808b81526020018a815260200189815260200187815260200186815250828886868f61105c565b50505050505050505050565b60038501546001600160a01b03163314156108c35760405162461bcd60e51b81526004018080602001828103825260218152602001806112136021913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5866040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561091257600080fd5b505af4158015610926573d6000803e3d6000fd5b505050506040513d602081101561093c57600080fd5b50516109795760405162461bcd60e51b81526004018080602001828103825260268152602001806112bb6026913960400191505060405180910390fd5b60026005860154600160401b900460ff16600381111561099557fe5b146109d15760405162461bcd60e51b815260040180806020018281038252602f81526020018061114c602f913960400191505060405180910390fd5b846001015473__$2556963077056ca10a6804584182250fbf$__6385ecb92a8685876040518463ffffffff1660e01b81526004018084815260200183600260200280838360005b83811015610a30578181015183820152602001610a18565b50505050905001828152602001935050505060206040518083038186803b158015610a5a57600080fd5b505af4158015610a6e573d6000803e3d6000fd5b505050506040513d6020811015610a8457600080fd5b5051604080516020818101939093528082018590528151808203830181526060909101909152805191012014610aeb5760405162461bcd60e51b815260040180806020018281038252604d81526020018061130e604d913960600191505060405180910390fd5b5050600060018401555050600501805460ff60401b1916600160401b179055565b805160009067ffffffffffffffff164310801590610b385750602082015167ffffffffffffffff164311155b92915050565b600260058b0154600160401b900460ff166003811115610b5a57fe5b14610b965760405162461bcd60e51b81526004018080602001828103825260228152602001806111cd6022913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be58b6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610be557600080fd5b505af4158015610bf9573d6000803e3d6000fd5b505050506040513d6020811015610c0f57600080fd5b505115610c4d5760405162461bcd60e51b81526004018080602001828103825260248152602001806111a96024913960400191505060405180910390fd5b896001015473__$2556963077056ca10a6804584182250fbf$__6385ecb92a8b8b8b6040518463ffffffff1660e01b81526004018084815260200183600260200280838360005b83811015610cac578181015183820152602001610c94565b50505050905001828152602001935050505060206040518083038186803b158015610cd657600080fd5b505af4158015610cea573d6000803e3d6000fd5b505050506040513d6020811015610d0057600080fd5b810190808051906020019092919050505073__$2556963077056ca10a6804584182250fbf$__631b0aa96b89898989600073__$2556963077056ca10a6804584182250fbf$__63e83f4bfe8c6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d98578181015183820152602001610d80565b50505050905090810190601f168015610dc55780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610de257600080fd5b505af4158015610df6573d6000803e3d6000fd5b505050506040513d6020811015610e0c57600080fd5b5051604080516001600160e01b031960e08a901b1681526004810197909752941515602487015263ffffffff93909316604486015267ffffffffffffffff9091166064850152608484015260a4830152600060c483015260e4820187905251610104808301926020929190829003018186803b158015610e8b57600080fd5b505af4158015610e9f573d6000803e3d6000fd5b505050506040513d6020811015610eb557600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019052805191012014610f1c5760405162461bcd60e51b815260040180806020018281038252604d81526020018061135b604d913960600191505060405180910390fd5b868515610f965773__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b158015610f6757600080fd5b505af4158015610f7b573d6000803e3d6000fd5b505050506040513d6020811015610f9157600080fd5b505190505b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63c20e84b28c89846040518463ffffffff1660e01b815260040180848152602001838152602001828152602001935050505060006040518083038186803b158015610ff557600080fd5b505af4158015611009573d6000803e3d6000fd5b5050604080518a8152891515602082015280820186905290517f413022a2562906c82f015c48fe406b4dd469f6cd003f9e3ff01f1ad240ba02ee9350908190036060019150a15050505050505050505050565b7f88bbc776ed3a8967efd66392bbebd915ac4fd011f320c12cbadd7baefcf1630b8633878787878760040160109054906101000a900467ffffffffffffffff166040518088600560200280838360005b838110156110c45781810151838201526020016110ac565b505050506001600160a01b038a1692019182525060200186604080838360005b838110156110fc5781810151838201526020016110e4565b505050509615159190960190815263ffffffff90941660208501525067ffffffffffffffff918216604080850191909152911660608301525190819003608001945092505050a150505050505056fe417373657274696f6e206d7573742062652070656e64696e6720746f20696e697469617465206368616c6c656e676543616e206f6e6c792064697370757461626c6520617373657274206966206e6f7420696e206368616c6c656e6765417373657274696f6e206973207374696c6c2070656e64696e67206368616c6c656e6765564d20646f6573206e6f74206861766520617373657274696f6e2070656e64696e67507265636f6e646974696f6e3a206e6f742077697468696e2074696d6520626f756e64734368616c6c656e676520776173206372656174656420627920617373657274657243616e206f6e6c792064697370757461626c6520617373657274206966206d616368696e65206973206e6f74206572726f726564206f722068616c746564507265636f6e646974696f6e3a20696e626f7820646f6573206e6f74206d61746368507265636f6e646974696f6e3a207374617465206861736820646f6573206e6f74206d617463684368616c6c656e676520646964206e6f7420636f6d65206265666f726520646561646c696e6543616e206f6e6c792064697370757461626c65206173736572742066726f6d2077616974696e67207374617465496e697469617465204368616c6c656e67653a20507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6e436f6e6669726d2044697370757461626c653a20507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6ea265627a7a7231582003ea82e0011e368f78ba1173e8f63382a74521c08868c82771052a605893ce4d64736f6c634300050f0032"

// DeployDisputable deploys a new Ethereum contract, binding an instance of Disputable to it.
func DeployDisputable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Disputable, error) {
	parsed, err := abi.JSON(strings.NewReader(DisputableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	protocolAddr, _, _, _ := DeployProtocol(auth, backend)
	DisputableBin = strings.Replace(DisputableBin, "__$2556963077056ca10a6804584182250fbf$__", protocolAddr.String()[2:], -1)

	valueAddr, _, _, _ := DeployValue(auth, backend)
	DisputableBin = strings.Replace(DisputableBin, "__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__", valueAddr.String()[2:], -1)

	vMAddr, _, _, _ := DeployVM(auth, backend)
	DisputableBin = strings.Replace(DisputableBin, "__$8e266570c8a7fb2aaac83b3e040afaf9e1$__", vMAddr.String()[2:], -1)

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
	NewState     [32]byte
	DidInboxInsn bool
	LogsAccHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x413022a2562906c82f015c48fe406b4dd469f6cd003f9e3ff01f1ad240ba02ee.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bool didInboxInsn, bytes32 logsAccHash)
func (_Disputable *DisputableFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts) (*DisputableConfirmedDisputableAssertionIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &DisputableConfirmedDisputableAssertionIterator{contract: _Disputable.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x413022a2562906c82f015c48fe406b4dd469f6cd003f9e3ff01f1ad240ba02ee.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bool didInboxInsn, bytes32 logsAccHash)
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

// ParseConfirmedDisputableAssertion is a log parse operation binding the contract event 0x413022a2562906c82f015c48fe406b4dd469f6cd003f9e3ff01f1ad240ba02ee.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bool didInboxInsn, bytes32 logsAccHash)
func (_Disputable *DisputableFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*DisputableConfirmedDisputableAssertion, error) {
	event := new(DisputableConfirmedDisputableAssertion)
	if err := _Disputable.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
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
	Fields       [5][32]byte
	Asserter     common.Address
	TimeBounds   [2]uint64
	DidInboxInsn bool
	NumSteps     uint32
	NumGas       uint64
	Deadline     uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x88bbc776ed3a8967efd66392bbebd915ac4fd011f320c12cbadd7baefcf1630b.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, bool didInboxInsn, uint32 numSteps, uint64 numGas, uint64 deadline)
func (_Disputable *DisputableFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*DisputablePendingDisputableAssertionIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &DisputablePendingDisputableAssertionIterator{contract: _Disputable.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x88bbc776ed3a8967efd66392bbebd915ac4fd011f320c12cbadd7baefcf1630b.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, bool didInboxInsn, uint32 numSteps, uint64 numGas, uint64 deadline)
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x88bbc776ed3a8967efd66392bbebd915ac4fd011f320c12cbadd7baefcf1630b.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, bool didInboxInsn, uint32 numSteps, uint64 numGas, uint64 deadline)
func (_Disputable *DisputableFilterer) ParsePendingDisputableAssertion(log types.Log) (*DisputablePendingDisputableAssertion, error) {
	event := new(DisputablePendingDisputableAssertion)
	if err := _Disputable.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IArbBaseABI is the input ABI used to generate the binding from.
const IArbBaseABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbBaseFuncSigs maps the 4-byte function signature to its string representation.
var IArbBaseFuncSigs = map[string]string{
	"22c091bc": "completeChallenge(address[2],uint128[2])",
}

// IArbBase is an auto generated Go binding around an Ethereum contract.
type IArbBase struct {
	IArbBaseCaller     // Read-only binding to the contract
	IArbBaseTransactor // Write-only binding to the contract
	IArbBaseFilterer   // Log filterer for contract events
}

// IArbBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbBaseSession struct {
	Contract     *IArbBase         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbBaseCallerSession struct {
	Contract *IArbBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IArbBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbBaseTransactorSession struct {
	Contract     *IArbBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IArbBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbBaseRaw struct {
	Contract *IArbBase // Generic contract binding to access the raw methods on
}

// IArbBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbBaseCallerRaw struct {
	Contract *IArbBaseCaller // Generic read-only contract binding to access the raw methods on
}

// IArbBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbBaseTransactorRaw struct {
	Contract *IArbBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbBase creates a new instance of IArbBase, bound to a specific deployed contract.
func NewIArbBase(address common.Address, backend bind.ContractBackend) (*IArbBase, error) {
	contract, err := bindIArbBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbBase{IArbBaseCaller: IArbBaseCaller{contract: contract}, IArbBaseTransactor: IArbBaseTransactor{contract: contract}, IArbBaseFilterer: IArbBaseFilterer{contract: contract}}, nil
}

// NewIArbBaseCaller creates a new read-only instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseCaller(address common.Address, caller bind.ContractCaller) (*IArbBaseCaller, error) {
	contract, err := bindIArbBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbBaseCaller{contract: contract}, nil
}

// NewIArbBaseTransactor creates a new write-only instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbBaseTransactor, error) {
	contract, err := bindIArbBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbBaseTransactor{contract: contract}, nil
}

// NewIArbBaseFilterer creates a new log filterer instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbBaseFilterer, error) {
	contract, err := bindIArbBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbBaseFilterer{contract: contract}, nil
}

// bindIArbBase binds a generic wrapper to an already deployed contract.
func bindIArbBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbBase *IArbBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbBase.Contract.IArbBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbBase *IArbBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbBase.Contract.IArbBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbBase *IArbBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbBase.Contract.IArbBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbBase *IArbBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbBase *IArbBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbBase *IArbBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbBase.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.Contract.CompleteChallenge(&_IArbBase.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.Contract.CompleteChallenge(&_IArbBase.TransactOpts, _players, _rewards)
}

// IArbChainABI is the input ABI used to generate the binding from.
const IArbChainABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeLauncherAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbChainFuncSigs maps the 4-byte function signature to its string representation.
var IArbChainFuncSigs = map[string]string{
	"8364fe47": "init(bytes32,uint32,uint32,uint128,address,address,address)",
}

// IArbChain is an auto generated Go binding around an Ethereum contract.
type IArbChain struct {
	IArbChainCaller     // Read-only binding to the contract
	IArbChainTransactor // Write-only binding to the contract
	IArbChainFilterer   // Log filterer for contract events
}

// IArbChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbChainSession struct {
	Contract     *IArbChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbChainCallerSession struct {
	Contract *IArbChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IArbChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbChainTransactorSession struct {
	Contract     *IArbChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IArbChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbChainRaw struct {
	Contract *IArbChain // Generic contract binding to access the raw methods on
}

// IArbChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbChainCallerRaw struct {
	Contract *IArbChainCaller // Generic read-only contract binding to access the raw methods on
}

// IArbChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbChainTransactorRaw struct {
	Contract *IArbChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbChain creates a new instance of IArbChain, bound to a specific deployed contract.
func NewIArbChain(address common.Address, backend bind.ContractBackend) (*IArbChain, error) {
	contract, err := bindIArbChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbChain{IArbChainCaller: IArbChainCaller{contract: contract}, IArbChainTransactor: IArbChainTransactor{contract: contract}, IArbChainFilterer: IArbChainFilterer{contract: contract}}, nil
}

// NewIArbChainCaller creates a new read-only instance of IArbChain, bound to a specific deployed contract.
func NewIArbChainCaller(address common.Address, caller bind.ContractCaller) (*IArbChainCaller, error) {
	contract, err := bindIArbChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChainCaller{contract: contract}, nil
}

// NewIArbChainTransactor creates a new write-only instance of IArbChain, bound to a specific deployed contract.
func NewIArbChainTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbChainTransactor, error) {
	contract, err := bindIArbChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChainTransactor{contract: contract}, nil
}

// NewIArbChainFilterer creates a new log filterer instance of IArbChain, bound to a specific deployed contract.
func NewIArbChainFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbChainFilterer, error) {
	contract, err := bindIArbChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbChainFilterer{contract: contract}, nil
}

// bindIArbChain binds a generic wrapper to an already deployed contract.
func bindIArbChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChain *IArbChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChain.Contract.IArbChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChain *IArbChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChain.Contract.IArbChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChain *IArbChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChain.Contract.IArbChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChain *IArbChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChain *IArbChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChain *IArbChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChain.Contract.contract.Transact(opts, method, params...)
}

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_IArbChain *IArbChainTransactor) Init(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbChain.contract.Transact(opts, "init", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_IArbChain *IArbChainSession) Init(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbChain.Contract.Init(&_IArbChain.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}

// Init is a paid mutator transaction binding the contract method 0x8364fe47.
//
// Solidity: function init(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address _challengeLauncherAddress, address _globalInboxAddress) returns()
func (_IArbChain *IArbChainTransactorSession) Init(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (*types.Transaction, error) {
	return _IArbChain.Contract.Init(&_IArbChain.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
}

// IChallengeFactoryABI is the input ABI used to generate the binding from.
const IChallengeFactoryABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeFactoryFuncSigs = map[string]string{
	"d3779e18": "createChallenge(address[2],uint128[2],uint32,bytes32,bytes32,uint64[2],bytes32)",
}

// IChallengeFactory is an auto generated Go binding around an Ethereum contract.
type IChallengeFactory struct {
	IChallengeFactoryCaller     // Read-only binding to the contract
	IChallengeFactoryTransactor // Write-only binding to the contract
	IChallengeFactoryFilterer   // Log filterer for contract events
}

// IChallengeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeFactorySession struct {
	Contract     *IChallengeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeFactoryCallerSession struct {
	Contract *IChallengeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeFactoryTransactorSession struct {
	Contract     *IChallengeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeFactoryRaw struct {
	Contract *IChallengeFactory // Generic contract binding to access the raw methods on
}

// IChallengeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeFactoryCallerRaw struct {
	Contract *IChallengeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeFactoryTransactorRaw struct {
	Contract *IChallengeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeFactory creates a new instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactory(address common.Address, backend bind.ContractBackend) (*IChallengeFactory, error) {
	contract, err := bindIChallengeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactory{IChallengeFactoryCaller: IChallengeFactoryCaller{contract: contract}, IChallengeFactoryTransactor: IChallengeFactoryTransactor{contract: contract}, IChallengeFactoryFilterer: IChallengeFactoryFilterer{contract: contract}}, nil
}

// NewIChallengeFactoryCaller creates a new read-only instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryCaller(address common.Address, caller bind.ContractCaller) (*IChallengeFactoryCaller, error) {
	contract, err := bindIChallengeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryCaller{contract: contract}, nil
}

// NewIChallengeFactoryTransactor creates a new write-only instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeFactoryTransactor, error) {
	contract, err := bindIChallengeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryTransactor{contract: contract}, nil
}

// NewIChallengeFactoryFilterer creates a new log filterer instance of IChallengeFactory, bound to a specific deployed contract.
func NewIChallengeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeFactoryFilterer, error) {
	contract, err := bindIChallengeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeFactoryFilterer{contract: contract}, nil
}

// bindIChallengeFactory binds a generic wrapper to an already deployed contract.
func bindIChallengeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeFactory *IChallengeFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeFactory.Contract.IChallengeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeFactory *IChallengeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.IChallengeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeFactory *IChallengeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.IChallengeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeFactory *IChallengeFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeFactory *IChallengeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeFactory *IChallengeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xd3779e18.
//
// Solidity: function createChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns(address)
func (_IChallengeFactory *IChallengeFactoryTransactor) CreateChallenge(opts *bind.TransactOpts, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IChallengeFactory.contract.Transact(opts, "createChallenge", _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xd3779e18.
//
// Solidity: function createChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns(address)
func (_IChallengeFactory *IChallengeFactorySession) CreateChallenge(_players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.CreateChallenge(&_IChallengeFactory.TransactOpts, _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xd3779e18.
//
// Solidity: function createChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns(address)
func (_IChallengeFactory *IChallengeFactoryTransactorSession) CreateChallenge(_players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IChallengeFactory.Contract.CreateChallenge(&_IChallengeFactory.TransactOpts, _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// IGlobalPendingInboxABI is the input ABI used to generate the binding from.
const IGlobalPendingInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes21\",\"name\":\"tokenType\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pullPendingMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var IGlobalPendingInboxFuncSigs = map[string]string{
	"3bbc3c32": "forwardMessage(address,bytes21,uint256,bytes,bytes)",
	"d106ec19": "pullPendingMessages()",
	"f3972383": "registerForInbox()",
	"3fc6eb80": "sendEthMessage(address,bytes)",
	"626cef85": "sendMessage(address,bytes21,uint256,bytes)",
	"e4eb8c63": "sendMessages(bytes)",
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

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessages(opts *bind.TransactOpts, _messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessages", _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _messages)
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

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_didInboxInsn\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ProtocolFuncSigs = map[string]string{
	"1b0aa96b": "generateAssertionHash(bytes32,bool,uint32,uint64,bytes32,bytes32,bytes32,bytes32)",
	"e83f4bfe": "generateLastMessageHash(bytes)",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"85ecb92a": "generatePreconditionHash(bytes32,uint64[2],bytes32)",
}

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x610aa4610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100555760003560e01c80624c28f61461005a5780631b0aa96b146100b257806385ecb92a1461010c578063e83f4bfe14610161575b600080fd5b6100a06004803603608081101561007057600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b0316610207565b60408051918252519081900360200190f35b6100a060048036036101008110156100c957600080fd5b50803590602081013515159063ffffffff6040820135169067ffffffffffffffff6060820135169060808101359060a08101359060c08101359060e001356102f9565b6100a06004803603608081101561012257600080fd5b604080518082018252833593928301929160608301919060208401906002908390839080828437600092019190915250919450509035915061036e9050565b6100a06004803603602081101561017757600080fd5b81019060208101813564010000000081111561019257600080fd5b8201836020820111156101a457600080fd5b803590602001918460018302840111640100000000831117156101c657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103c2945050505050565b60408051600480825260a0820190925260009160609190816020015b61022b610a08565b81526020019060019003908161022357905050905061024986610507565b8160008151811061025657fe5b6020026020010181905250610273836001600160a01b0316610585565b8160018151811061028057fe5b602002602001018190525061029484610585565b816002815181106102a157fe5b60209081029190910101526102c36affffffffffffffffffffff198616610585565b816003815181106102d057fe5b60200260200101819052506102ec6102e782610603565b6106b3565b519150505b949350505050565b6040805160208082019a909a5297151560f81b8882015260e09690961b6001600160e01b031916604188015260c09490941b6001600160c01b0319166045870152604d860192909252606d850152608d84015260ad808401919091528151808403909101815260cd9092019052805191012090565b815160209283015160408051808601969096526001600160c01b031960c093841b8116878301529190921b166048850152605080850192909252805180850390920182526070909301909252815191012090565b8051600090819081908190815b818110156104fa5773__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63d36cfac288866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561044557818101518382015260200161042d565b50505050905090810190601f1680156104725780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561048f57600080fd5b505af41580156104a3573d6000803e3d6000fd5b505050506040513d60408110156104b957600080fd5b50805160209182015160408051808501999099528881018290528051808a0382018152606090990190528751979092019690962095945092506001016103cf565b509293505050505b919050565b61050f610a08565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610574565b610561610a08565b8152602001906001900390816105595790505b508152600260209091015292915050565b61058d610a08565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916105f2565b6105df610a08565b8152602001906001900390816105d75790505b508152600060209091015292915050565b61060b610a08565b61061582516107e9565b610666576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6106bb610a36565b6060820151600c60ff9091161061070d576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661073a57604051806020016040528061073184600001516107f0565b90529050610502565b606082015160ff1660011415610781576040518060200160405280610731846020015160000151856020015160400151866020015160600151876020015160200151610814565b606082015160ff16600214156107a65750604080516020810190915281518152610502565b600360ff16826060015160ff16101580156107ca57506060820151600c60ff909116105b156107e757604051806020016040528061073184604001516108bc565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561086e575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206102f1565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b600060088251111561090c576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610939578160200160208202803883390190505b50805190915060005b8181101561099557610952610a36565b61096e86838151811061096157fe5b60200260200101516106b3565b9050806000015184838151811061098157fe5b602090810291909101015250600101610942565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156109de5781810151838201526020016109c6565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b604051806080016040528060008152602001610a22610a48565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a723158205b3a97f76d858b1c74d36531849bba18daf859f117baaa2b514915594d47bfe864736f6c634300050f0032"

// DeployProtocol deploys a new Ethereum contract, binding an instance of Protocol to it.
func DeployProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Protocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	valueAddr, _, _, _ := DeployValue(auth, backend)
	ProtocolBin = strings.Replace(ProtocolBin, "__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__", valueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// Protocol is an auto generated Go binding around an Ethereum contract.
type Protocol struct {
	ProtocolCaller     // Read-only binding to the contract
	ProtocolTransactor // Write-only binding to the contract
	ProtocolFilterer   // Log filterer for contract events
}

// ProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolSession struct {
	Contract     *Protocol         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolCallerSession struct {
	Contract *ProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolTransactorSession struct {
	Contract     *ProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRaw struct {
	Contract *Protocol // Generic contract binding to access the raw methods on
}

// ProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolCallerRaw struct {
	Contract *ProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolTransactorRaw struct {
	Contract *ProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocol creates a new instance of Protocol, bound to a specific deployed contract.
func NewProtocol(address common.Address, backend bind.ContractBackend) (*Protocol, error) {
	contract, err := bindProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// NewProtocolCaller creates a new read-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolCaller(address common.Address, caller bind.ContractCaller) (*ProtocolCaller, error) {
	contract, err := bindProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolCaller{contract: contract}, nil
}

// NewProtocolTransactor creates a new write-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolTransactor, error) {
	contract, err := bindProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTransactor{contract: contract}, nil
}

// NewProtocolFilterer creates a new log filterer instance of Protocol, bound to a specific deployed contract.
func NewProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolFilterer, error) {
	contract, err := bindProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolFilterer{contract: contract}, nil
}

// bindProtocol binds a generic wrapper to an already deployed contract.
func bindProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.ProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transact(opts, method, params...)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1b0aa96b.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _gas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _gas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _didInboxInsn, _numSteps, _gas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1b0aa96b.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _gas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _gas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateAssertionHash(&_Protocol.CallOpts, _afterHash, _didInboxInsn, _numSteps, _gas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1b0aa96b.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, bool _didInboxInsn, uint32 _numSteps, uint64 _gas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _didInboxInsn bool, _numSteps uint32, _gas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateAssertionHash(&_Protocol.CallOpts, _afterHash, _didInboxInsn, _numSteps, _gas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _messages []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateLastMessageHash", _messages)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateLastMessageHash(&_Protocol.CallOpts, _messages)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateLastMessageHash(&_Protocol.CallOpts, _messages)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _Protocol.Contract.GenerateMessageStubHash(&_Protocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _Protocol.Contract.GenerateMessageStubHash(&_Protocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GeneratePreconditionHash(&_Protocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GeneratePreconditionHash(&_Protocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201ca52c8cfd94848edb89c2f757752003702d51eace6d9e6123cdcc328cbf2aca64736f6c634300050f0032"

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
	"c20e84b2": "acceptAssertion(VM.Data storage,bytes32,bytes32)",
	"2a3e0a97": "isErrored(VM.Data storage)",
	"e2fe93ca": "isHalted(VM.Data storage)",
	"a3a162cb": "resetDeadline(VM.Data storage)",
	"8ab48be5": "withinDeadline(VM.Data storage)",
}

// VMBin is the compiled bytecode used for deploying new contracts.
var VMBin = "0x6101f8610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80632a3e0a97146100665780638ab48be514610097578063a3a162cb146100b4578063c20e84b2146100e0578063e2fe93ca14610116575b600080fd5b6100836004803603602081101561007c57600080fd5b5035610133565b604080519115158252519081900360200190f35b610083600480360360208110156100ad57600080fd5b503561013a565b8180156100c057600080fd5b506100de600480360360208110156100d757600080fd5b5035610155565b005b8180156100ec57600080fd5b506100de6004803603606081101561010357600080fd5b5080359060208101359060400135610194565b6100836004803603602081101561012c57600080fd5b50356101be565b5460011490565b60040154600160801b900467ffffffffffffffff1643111590565b60058101546004909101805467ffffffffffffffff60801b1916600160801b63ffffffff909316430167ffffffffffffffff1692909202919091179055565b90825560058201805468ff0000000000000000191668010000000000000000179055600290910155565b54159056fea265627a7a723158203eff68375df4026b7479a5aff051781edec2335e66aa6581368c2c79d70987ae64736f6c634300050f0032"

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

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeHashed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidHashed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"immediate\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasic\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashInt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueFuncSigs maps the 4-byte function signature to its string representation.
var ValueFuncSigs = map[string]string{
	"3d730ed2": "deserializeHashed(bytes)",
	"32e6cc21": "deserializeMessage(bytes,uint256)",
	"d36cfac2": "deserializeValidHashed(bytes,uint256)",
	"72403aa0": "getNextValid(bytes,uint256)",
	"826513e0": "hashCodePoint(uint8,bool,bytes32,bytes32)",
	"b697e085": "hashCodePointBasic(uint8,bytes32)",
	"3c786053": "hashCodePointImmediate(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"5043dff1": "hashInt(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x61152b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c806372403aa01161007057806372403aa014610300578063826513e014610425578063b2b9dc6214610459578063b697e0851461048a578063d36cfac2146104b0576100a8565b806332e6cc21146100ad578063364df277146101f95780633c786053146102135780633d730ed21461023f5780635043dff1146102e3575b600080fd5b610153600480360360408110156100c357600080fd5b810190602081018135600160201b8111156100dd57600080fd5b8201836020820111156100ef57600080fd5b803590602001918460018302840111600160201b8311171561011057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061056f915050565b604051808815151515815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101b85781810151838201526020016101a0565b50505050905090810190601f1680156101e55780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b61020161076d565b60408051918252519081900360200190f35b6102016004803603606081101561022957600080fd5b5060ff81351690602081013590604001356107e0565b6102016004803603602081101561025557600080fd5b810190602081018135600160201b81111561026f57600080fd5b82018360208201111561028157600080fd5b803590602001918460018302840111600160201b831117156102a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610832945050505050565b610201600480360360208110156102f957600080fd5b50356108a6565b6103a66004803603604081101561031657600080fd5b810190602081018135600160201b81111561033057600080fd5b82018360208201111561034257600080fd5b803590602001918460018302840111600160201b8311171561036357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506108ca915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103e95781810151838201526020016103d1565b50505050905090810190601f1680156104165780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6102016004803603608081101561043b57600080fd5b5060ff8135169060208101351515906040810135906060013561094e565b6104766004803603602081101561046f57600080fd5b50356109f7565b604080519115158252519081900360200190f35b610201600480360360408110156104a057600080fd5b5060ff81351690602001356109fe565b610556600480360360408110156104c657600080fd5b810190602081018135600160201b8111156104e057600080fd5b8201836020820111156104f257600080fd5b803590602001918460018302840111600160201b8311171561051357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610a45915050565b6040805192835260208301919091528051918290030190f35b6000806000806000806060600088965060008a888151811061058d57fe5b016020015160019098019760f81c9050600781146105bf576105b28b60018a03610a45565b9098509650610761915050565b6105c98b89610a45565b90985091506105e88b60018c016000198d8c030163ffffffff610abe16565b92508a88815181106105f657fe5b016020015160019098019760f81c90508015610619576105b28b60018a03610a45565b6106238b89610b3e565b80995081975050508a888151811061063757fe5b016020015160019098019760f81c9050801561065a576105b28b60018a03610a45565b6106648b89610b3e565b80995081965050508a888151811061067857fe5b016020015160019098019760f81c9050801561069b576105b28b60018a03610a45565b6106a58b89610b3e565b60408051600480825260a0820190925260019c50919a509195506060916020820160808038833901905050905082816000815181106106e057fe5b6020026020010181815250506106f5876108a6565b8160018151811061070257fe5b602002602001018181525050610717866108a6565b8160028151811061072457fe5b602002602001018181525050610739856108a6565b8160038151811061074657fe5b60200260200101818152505061075b81610b65565b97505050505b92959891949750929550565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156107b95781810151838201526020016107a1565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6000808061083e61146f565b610849856000610c25565b919450925090508215610891576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61089a81610daf565b5193505050505b919050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060606000806108d961146f565b6108e38787610c25565b91945092509050821561092b576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b8161093f888880840363ffffffff610abe16565b945094505050505b9250929050565b600083156109a8575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206109ef565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b6008101590565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b600080600080610a5361146f565b610a5d8787610c25565b919450925090508215610aa5576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b81610aaf82610daf565b51909890975095505050505050565b606081830184511015610ad057600080fd5b606082158015610aeb57604051915060208201604052610b35565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610b24578051835260209283019201610b0c565b5050858452601f01601f1916604052505b50949350505050565b6000808281610b53868363ffffffff610ee516565b60209290920196919550909350505050565b6000600882511115610bb5576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b8151600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610bfd578181015183820152602001610be5565b5050505090500192505050604051602081830303815290604052805190602001209050919050565b600080610c3061146f565b84518410610c85576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110610c9857fe5b016020015160019092019160f81c90506000610cb261149d565b60ff8316610ce657610cc48985610b3e565b9094509150600084610cd584610f01565b91985096509450610da89350505050565b60ff831660011415610d0d57610cfc8985610f7f565b9094509050600084610cd5836110da565b60ff831660021415610d3457610d238985610b3e565b9094509150600084610cd58461113a565b600360ff841610801590610d4b5750600c60ff8416105b15610d8857600219830160606000610d64838d896111b8565b909850925090508087610d7684611273565b99509950995050505050505050610da8565b8260ff16612710016000610d9c6000610f01565b91985096509450505050505b9250925092565b610db76114c4565b6060820151600c60ff90911610610e09576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610e36576040518060200160405280610e2d84600001516108a6565b905290506108a1565b606082015160ff1660011415610e7d576040518060200160405280610e2d84602001516000015185602001516040015186602001516060015187602001516020015161094e565b606082015160ff1660021415610ea257506040805160208101909152815181526108a1565b600360ff16826060015160ff1610158015610ec657506060820151600c60ff909116105b15610ee3576040518060200160405280610e2d8460400151611323565bfe5b60008160200183511015610ef857600080fd5b50016020015190565b610f0961146f565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610f6e565b610f5b61146f565b815260200190600190039081610f535790505b508152600060209091015292915050565b6000610f8961149d565b60008390506000858281518110610f9c57fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110610fc257fe5b016020015160019384019360f89190911c915060009060ff8416141561104e576000610fec61146f565b610ff68a87610c25565b9097509092509050811561103f576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61104881610daf565b51925050505b6000611060898663ffffffff610ee516565b90506020850194508360ff16600114156110a5576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506109479050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6110e261146f565b604080516080810182526000808252602080830186905283518281529081018452919283019190611129565b61111661146f565b81526020019060019003908161110e5790505b508152600160209091015292915050565b61114261146f565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916111a7565b61119461146f565b81526020019060019003908161118c5790505b508152600260209091015292915050565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561120357816020015b6111f061146f565b8152602001906001900390816111e85790505b50905060005b8960ff168160ff16101561125d576112218985610c25565b8451859060ff861690811061123257fe5b60209081029190910101529450925082156112555750909450909250905061126a565b600101611209565b5060009550919350909150505b93509350939050565b61127b61146f565b61128582516109f7565b6112d6576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6000600882511115611373576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156113a0578160200160208202803883390190505b50805190915060005b818110156113fc576113b96114c4565b6113d58683815181106113c857fe5b6020026020010151610daf565b905080600001518483815181106113e857fe5b6020908102919091010152506001016113a9565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561144557818101518382015260200161142d565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180608001604052806000815260200161148961149d565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4d61727368616c6c65642076616c7565206d7573742062652076616c69640000a265627a7a72315820f789a39100f31f1c48384e7b3a19af130cf6e23c6edeba2981393c94e1ea928064736f6c634300050f0032"

// DeployValue deploys a new Ethereum contract, binding an instance of Value to it.
func DeployValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Value, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// Value is an auto generated Go binding around an Ethereum contract.
type Value struct {
	ValueCaller     // Read-only binding to the contract
	ValueTransactor // Write-only binding to the contract
	ValueFilterer   // Log filterer for contract events
}

// ValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueSession struct {
	Contract     *Value            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueCallerSession struct {
	Contract *ValueCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueTransactorSession struct {
	Contract     *ValueTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueRaw struct {
	Contract *Value // Generic contract binding to access the raw methods on
}

// ValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueCallerRaw struct {
	Contract *ValueCaller // Generic read-only contract binding to access the raw methods on
}

// ValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueTransactorRaw struct {
	Contract *ValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValue creates a new instance of Value, bound to a specific deployed contract.
func NewValue(address common.Address, backend bind.ContractBackend) (*Value, error) {
	contract, err := bindValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// NewValueCaller creates a new read-only instance of Value, bound to a specific deployed contract.
func NewValueCaller(address common.Address, caller bind.ContractCaller) (*ValueCaller, error) {
	contract, err := bindValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueCaller{contract: contract}, nil
}

// NewValueTransactor creates a new write-only instance of Value, bound to a specific deployed contract.
func NewValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueTransactor, error) {
	contract, err := bindValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTransactor{contract: contract}, nil
}

// NewValueFilterer creates a new log filterer instance of Value, bound to a specific deployed contract.
func NewValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueFilterer, error) {
	contract, err := bindValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueFilterer{contract: contract}, nil
}

// bindValue binds a generic wrapper to an already deployed contract.
func bindValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.ValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.contract.Transact(opts, method, params...)
}

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueCaller) DeserializeHashed(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "deserializeHashed", data)
	return *ret0, err
}

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueSession) DeserializeHashed(data []byte) ([32]byte, error) {
	return _Value.Contract.DeserializeHashed(&_Value.CallOpts, data)
}

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueCallerSession) DeserializeHashed(data []byte) ([32]byte, error) {
	return _Value.Contract.DeserializeHashed(&_Value.CallOpts, data)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueCaller) DeserializeMessage(opts *bind.CallOpts, data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	ret := new(struct {
		Valid       bool
		Offset      *big.Int
		MessageHash [32]byte
		Destination *big.Int
		Value       *big.Int
		TokenType   *big.Int
		MessageData []byte
	})
	out := ret
	err := _Value.contract.Call(opts, out, "deserializeMessage", data, startOffset)
	return *ret, err
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _Value.Contract.DeserializeMessage(&_Value.CallOpts, data, startOffset)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueCallerSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _Value.Contract.DeserializeMessage(&_Value.CallOpts, data, startOffset)
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueCaller) DeserializeValidHashed(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Value.contract.Call(opts, out, "deserializeValidHashed", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueSession) DeserializeValidHashed(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _Value.Contract.DeserializeValidHashed(&_Value.CallOpts, data, offset)
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueCallerSession) DeserializeValidHashed(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _Value.Contract.DeserializeValidHashed(&_Value.CallOpts, data, offset)
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueCaller) GetNextValid(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Value.contract.Call(opts, out, "getNextValid", data, offset)
	return *ret0, *ret1, err
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueSession) GetNextValid(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _Value.Contract.GetNextValid(&_Value.CallOpts, data, offset)
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueCallerSession) GetNextValid(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _Value.Contract.GetNextValid(&_Value.CallOpts, data, offset)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePoint(opts *bind.CallOpts, opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePoint", opcode, immediate, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePoint(&_Value.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePoint(&_Value.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePointBasic(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePointBasic", opcode, nextCodePoint)
	return *ret0, err
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePointBasic(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointBasic(&_Value.CallOpts, opcode, nextCodePoint)
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePointBasic(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointBasic(&_Value.CallOpts, opcode, nextCodePoint)
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePointImmediate(opts *bind.CallOpts, opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePointImmediate", opcode, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePointImmediate(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointImmediate(&_Value.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePointImmediate(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointImmediate(&_Value.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueSession) HashEmptyTuple() ([32]byte, error) {
	return _Value.Contract.HashEmptyTuple(&_Value.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _Value.Contract.HashEmptyTuple(&_Value.CallOpts)
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueCaller) HashInt(opts *bind.CallOpts, val *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashInt", val)
	return *ret0, err
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueSession) HashInt(val *big.Int) ([32]byte, error) {
	return _Value.Contract.HashInt(&_Value.CallOpts, val)
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueCallerSession) HashInt(val *big.Int) ([32]byte, error) {
	return _Value.Contract.HashInt(&_Value.CallOpts, val)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueCaller) IsValidTupleSize(opts *bind.CallOpts, size *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "isValidTupleSize", size)
	return *ret0, err
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _Value.Contract.IsValidTupleSize(&_Value.CallOpts, size)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueCallerSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _Value.Contract.IsValidTupleSize(&_Value.CallOpts, size)
}
