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
const ArbChainABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeLauncherAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"ChallengeLaunched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeLauncher\",\"outputs\":[{\"internalType\":\"contractIChallengeLauncher\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumVM.State\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"enumVM.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"activeChallengeManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbChainFuncSigs maps the 4-byte function signature to its string representation.
var ArbChainFuncSigs = map[string]string{
	"94af716b": "activateVM()",
	"8aee8c79": "challengeLauncher()",
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"f2204f74": "confirmDisputableAsserted(bytes32,bytes32,uint32,bytes,bytes32)",
	"08dc89d7": "currentDeposit(address)",
	"aca0f372": "escrowRequired()",
	"6be00229": "exitAddress()",
	"1865c57d": "getState()",
	"d489113a": "globalInbox()",
	"05b050de": "increaseDeposit()",
	"2782e87e": "initiateChallenge(bytes32)",
	"8da5cb5b": "owner()",
	"cfa80707": "ownerShutdown()",
	"fec5a286": "pendingDisputableAssert(bytes32,bytes32,bytes32,bytes32,bytes32,uint32,uint64[2])",
	"60675a87": "terminateAddress()",
	"3a768463": "vm()",
}

// ArbChainBin is the compiled bytecode used for deploying new contracts.
var ArbChainBin = "0x608060405234801561001057600080fd5b50604051611795380380611795833981810160405260e081101561003357600080fd5b50805160208201516040808401516060850151608086015160a087015160c090970151600180546001600160a01b038084166001600160a01b0319928316179283905560008054828d16931692909217825587517ff39723830000000000000000000000000000000000000000000000000000000081529751999a989996989597949693948b948b948b948b948b948b948b949092169263f39723839260048084019382900301818387803b1580156100eb57600080fd5b505af11580156100ff573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b15801561018657600080fd5b505af415801561019a573d6000803e3d6000fd5b505050506040513d60208110156101b057600080fd5b50516004555050600680546001600160801b0319166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161763ffffffff60201b19166401000000009290931691909102919091178082556001925060ff60401b191668010000000000000000830217905550505050505050506115588061023d6000396000f3fe6080604052600436106100f35760003560e01c80638aee8c791161008a578063cfa8070711610059578063cfa8070714610323578063d489113a14610338578063f2204f741461034d578063fec5a2861461041b576100f3565b80638aee8c79146102cf5780638da5cb5b146102e457806394af716b146102f9578063aca0f3721461030e576100f3565b80632782e87e116100c65780632782e87e146101ad5780633a768463146101d757806360675a87146102895780636be00229146102ba576100f3565b806305b050de146100f857806308dc89d7146101025780631865c57d1461014757806322c091bc14610180575b600080fd5b61010061049f565b005b34801561010e57600080fd5b506101356004803603602081101561012557600080fd5b50356001600160a01b03166104b6565b60408051918252519081900360200190f35b34801561015357600080fd5b5061015c6104d5565b6040518082600381111561016c57fe5b60ff16815260200191505060405180910390f35b34801561018c57600080fd5b50610100600480360360808110156101a357600080fd5b50604081016104e5565b3480156101b957600080fd5b50610100600480360360208110156101d057600080fd5b50356105f7565b3480156101e357600080fd5b506101ec610879565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561025c57fe5b60ff1681526001600160a01b039092166020830152506040805191829003019a5098505050505050505050f35b34801561029557600080fd5b5061029e6108ea565b604080516001600160a01b039092168252519081900360200190f35b3480156102c657600080fd5b5061029e6108f9565b3480156102db57600080fd5b5061029e610908565b3480156102f057600080fd5b5061029e610917565b34801561030557600080fd5b50610100610926565b34801561031a57600080fd5b506101356109b9565b34801561032f57600080fd5b506101006109c8565b34801561034457600080fd5b5061029e610a28565b34801561035957600080fd5b50610100600480360360a081101561037057600080fd5b81359160208101359163ffffffff60408301351691908101906080810160608201356401000000008111156103a457600080fd5b8201836020820111156103b657600080fd5b803590602001918460018302840111640100000000831117156103d857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610a37915050565b34801561042757600080fd5b50610100600480360361010081101561043f57600080fd5b604080518082018252833593602081013593838201359360608301359360808401359363ffffffff60a08201351693810192909161010083019160c08401906002908390839080828437600092019190915250919450610b979350505050565b336000908152600860205260409020805434019055565b6001600160a01b0381166000908152600860205260409020545b919050565b600754600160401b900460ff1690565b600754600160481b90046001600160a01b031633146105355760405162461bcd60e51b815260040180806020018281038252602d8152602001806114c6602d913960400191505060405180910390fd5b60078054600160481b600160e81b031916905561059a6001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002054610cf890919063ffffffff16565b82356001600160a01b031660009081526008602081815260408320939093556105d2928401356001600160801b03169185600161055d565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b031611156106505760405162461bcd60e51b815260040180806020018281038252602781526020018061149f6027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b1580156106ca57600080fd5b505af41580156106de573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b03168082529181019190915260075484516328d23fe960e01b81529290951697506328d23fe996509194919363ffffffff16928892600490920191829187918190849084905b8381101561077657818101518382015260200161075e565b5050505090500184600260200280838360005b838110156107a1578181015183820152602001610789565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050602060405180830381600087803b1580156107e057600080fd5b505af11580156107f4573d6000803e3d6000fd5b505050506040513d602081101561080a57600080fd5b505160078054600160481b600160e81b031916600160481b6001600160a01b03938416810291909117918290556040805191909204909216825233602083015280517f65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f9281900390910190a150565b6002546003546004546005546006546007546001600160a01b03928316926001600160801b0383169267ffffffffffffffff600160801b8204811693600160c01b909204169163ffffffff8083169264010000000081049091169160ff600160401b83041691600160481b9004168b565b600a546001600160a01b031681565b6009546001600160a01b031681565b6000546001600160a01b031681565b600b546001600160a01b031681565b600b546001600160a01b0316331461097e576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff16600381111561099857fe5b14156109b7576007805468ff00000000000000001916600160401b1790555b565b6006546001600160801b031690565b600b546001600160a01b03163314610a20576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6109b7610d59565b6001546001600160a01b031681565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63ce9d5122600287878787876040518763ffffffff1660e01b8152600401808781526020018681526020018581526020018463ffffffff1663ffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610ace578181015183820152602001610ab6565b50505050905090810190601f168015610afb5780820380516001836020036101000a031916815260200191505b5097505050505050505060006040518083038186803b158015610b1d57600080fd5b505af4158015610b31573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610b6c935091506001600160801b031663ffffffff610cf816565b6005546001600160a01b0316600090815260086020526040902055610b9082610d67565b5050505050565b336000908152600860205260409020546006546001600160801b03161115610bf05760405162461bcd60e51b81526004018080602001828103825260318152602001806114f36031913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151636c36f28f60e11b8152600260048201818152602483018c9052604483018b9052606483018a90526084830189905260a4830188905263ffffffff871660c484015273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__9463d86de51e9492938d938d938d938d938d938d938d93909260e401918491908190849084905b83811015610cae578181015183820152602001610c96565b505050509050019850505050505050505060006040518083038186803b158015610cd757600080fd5b505af4158015610ceb573d6000803e3d6000fd5b5050505050505050505050565b600082820183811015610d52576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600b546001600160a01b0316ff5b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b158015610dad57600080fd5b505af1158015610dc1573d6000803e3d6000fd5b505050506040513d6020811015610dd757600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b158015610e2857600080fd5b505af4158015610e3c573d6000803e3d6000fd5b505050506040513d6020811015610e5257600080fd5b50518114610e9d57610e996040518060600160405280610e726001610f65565b8152602001610e846002800154610fe3565b8152602001610e9284610fe3565b9052611061565b6004555b60015460405163e4eb8c6360e01b81526020600482018181528551602484015285516001600160a01b039094169363e4eb8c6393879383926044909201919085019080838360005b83811015610efd578181015183820152602001610ee5565b50505050905090810190601f168015610f2a5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610f4957600080fd5b505af1158015610f5d573d6000803e3d6000fd5b505050505050565b610f6d611437565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610fd2565b610fbf611437565b815260200190600190039081610fb75790505b508152600060209091015292915050565b610feb611437565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611050565b61103d611437565b8152602001906001900390816110355790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b611085611437565b81526020019060019003908161107d575050805190915060005b818110156110d7578481600381106110b357fe5b60200201518382815181106110c457fe5b602090810291909101015260010161109f565b506110e1826110e9565b949350505050565b6000600882511115611139576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611166578160200160208202803883390190505b50805190915060005b818110156111c25761117f611465565b61119b86838151811061118e57fe5b6020026020010151611235565b905080600001518483815181106111ae57fe5b60209081029190910101525060010161116f565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561120b5781810151838201526020016111f3565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b61123d611465565b6060820151600c60ff9091161061128f576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166112bc5760405180602001604052806112b3846000015161136b565b905290506104d0565b606082015160ff16600114156113035760405180602001604052806112b384602001516000015185602001516040015186602001516060015187602001516020015161138f565b606082015160ff166002141561132857506040805160208101909152815181526104d0565b600360ff16826060015160ff161015801561134c57506060820151600c60ff909116105b156113695760405180602001604052806112b384604001516110e9565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b600083156113e9575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206110e1565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b604051806080016040528060008152602001611451611477565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f7765644f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274a265627a7a7231582007c6388d1687268196b33f5d08d891b5c2d7231f6d67bece174dcd9a5e3a08f764736f6c634300050d0032"

// DeployArbChain deploys a new Ethereum contract, binding an instance of ArbChain to it.
func DeployArbChain(auth *bind.TransactOpts, backend bind.ContractBackend, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (common.Address, *types.Transaction, *ArbChain, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ArbChainBin = strings.Replace(ArbChainBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbChainBin = strings.Replace(ArbChainBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbChainBin), backend, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
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

// ChallengeLauncher is a free data retrieval call binding the contract method 0x8aee8c79.
//
// Solidity: function challengeLauncher() constant returns(address)
func (_ArbChain *ArbChainCaller) ChallengeLauncher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChain.contract.Call(opts, out, "challengeLauncher")
	return *ret0, err
}

// ChallengeLauncher is a free data retrieval call binding the contract method 0x8aee8c79.
//
// Solidity: function challengeLauncher() constant returns(address)
func (_ArbChain *ArbChainSession) ChallengeLauncher() (common.Address, error) {
	return _ArbChain.Contract.ChallengeLauncher(&_ArbChain.CallOpts)
}

// ChallengeLauncher is a free data retrieval call binding the contract method 0x8aee8c79.
//
// Solidity: function challengeLauncher() constant returns(address)
func (_ArbChain *ArbChainCallerSession) ChallengeLauncher() (common.Address, error) {
	return _ArbChain.Contract.ChallengeLauncher(&_ArbChain.CallOpts)
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

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "confirmDisputableAsserted", _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.ConfirmDisputableAsserted(&_ArbChain.TransactOpts, _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChain *ArbChainTransactorSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChain.Contract.ConfirmDisputableAsserted(&_ArbChain.TransactOpts, _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
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

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbChain *ArbChainTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChain.contract.Transact(opts, "pendingDisputableAssert", _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbChain *ArbChainSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChain.Contract.PendingDisputableAssert(&_ArbChain.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbChain *ArbChainTransactorSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChain.Contract.PendingDisputableAssert(&_ArbChain.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
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
	Deadline   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
func (_ArbChain *ArbChainFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*ArbChainPendingDisputableAssertionIterator, error) {

	logs, sub, err := _ArbChain.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChainPendingDisputableAssertionIterator{contract: _ArbChain.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
func (_ArbChain *ArbChainFilterer) ParsePendingDisputableAssertion(log types.Log) (*ArbChainPendingDisputableAssertion, error) {
	event := new(ArbChainPendingDisputableAssertion)
	if err := _ArbChain.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeABI is the input ABI used to generate the binding from.
const ArbChallengeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"challengerWrong\",\"type\":\"bool\"}],\"name\":\"TimedOutChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"asserterTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint32\",\"name\":\"_totalSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"challengerTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assertionToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"continueChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"_beforeHashAndInbox\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32[5]\",\"name\":\"_afterHashAndMessages\",\"type\":\"bytes32[5]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ArbChallengeFuncSigs = map[string]string{
	"d5345e07": "asserterTimedOut()",
	"71f45bbf": "bisectAssertion(bytes32,bytes32[],uint32,uint64[2])",
	"635e28a7": "challengerTimedOut()",
	"79d84776": "continueChallenge(uint256,bytes,bytes32,bytes32)",
	"1d7aaea9": "oneStepProof(bytes32[2],uint64[2],bytes32[5],bytes)",
}

// ArbChallengeBin is the compiled bytecode used for deploying new contracts.
var ArbChallengeBin = "0x60806040523480156200001157600080fd5b5060405162000e1538038062000e15833981810160405260e08110156200003757600080fd5b50805160a08083015160c0808501516040805160e0810182526001600160a01b03871680825260208083018590526060808b01948401859052990198820189905263ffffffff86164381016001600160401b0381166080850152978301526001948201859052600080546001600160a01b03191690911781559383905595969590949192620000c960028781620001c3565b506060820151620000e190600383019060026200026f565b50608082015160058201805460a085015163ffffffff16680100000000000000000263ffffffff60401b196001600160401b039094166001600160401b031990921691909117929092169190911780825560c0840151919060ff60601b19166c010000000000000000000000008360028111156200015b57fe5b02179055505050602080860151604080516001600160a01b0390921682526001600160401b0384169282019290925281517f6dc74e1677661f103d909b4e12489baf51a49c7baf11ba1d4bf1a9fc899e9a0e929181900390910190a150505050505062000319565b6001830191839082156200025d5791602002820160005b838211156200022657835183826101000a8154816001600160801b0302191690836001600160801b031602179055509260200192601001602081600f01049283019260010302620001da565b80156200025b5782816101000a8154906001600160801b030219169055601001602081600f0104928301926001030262000226565b505b506200026b929150620002c8565b5090565b8260028101928215620002ba579160200282015b82811115620002ba57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000283565b506200026b929150620002f2565b620002ef91905b808211156200026b5780546001600160801b0319168155600101620002cf565b90565b620002ef91905b808211156200026b5780546001600160a01b0319168155600101620002f9565b610aec80620003296000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631d7aaea91461005c578063635e28a71461018657806371f45bbf1461018e57806379d8477614610270578063d5345e0714610322575b600080fd5b610184600480360361014081101561007357600080fd5b604080518082018252918301929181830191839060029083908390808284376000920191909152505060408051808201825292959493818101939250906002908390839080828437600092019190915250506040805160a0818101909252929594938181019392509060059083908390808284376000920191909152509194939260208101925035905064010000000081111561010f57600080fd5b82018360208201111561012157600080fd5b8035906020019184600183028401116401000000008311171561014357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061032a945050505050565b005b610184610544565b610184600480360360a08110156101a457600080fd5b813591908101906040810160208201356401000000008111156101c657600080fd5b8201836020820111156101d857600080fd5b803590602001918460208302840111640100000000831117156101fa57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080518082018252939663ffffffff863516969095909460608201945092506020019060029083908390808284376000920191909152509194506106329350505050565b6101846004803603608081101561028657600080fd5b813591908101906040810160208201356401000000008111156102a857600080fd5b8201836020820111156102ba57600080fd5b803590602001918460018302840111640100000000831117156102dc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561072f565b610184610801565b73__$f55f7f918072b72dcc999cdc8e581605f5$__63a49c33086000868686866040518663ffffffff1660e01b81526004018086815260200185600260200280838360005b8381101561038757818101518382015260200161036f565b5050505090500184600260200280838360005b838110156103b257818101518382015260200161039a565b5050505090500183600560200280838360005b838110156103dd5781810151838201526020016103c5565b5050505090500180602001828103825283818151815260200191508051906020019080838360005b8381101561041d578181015183820152602001610405565b50505050905090810190601f16801561044a5780820380516001836020036101000a031916815260200191505b50965050505050505060006040518083038186803b15801561046b57600080fd5b505af415801561047f573d6000803e3d6000fd5b5050505061048b6108ef565b7f1a96858c84fb221338517840d882a8fd1434f2cbbea7738d6e70333a801231a8338260405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156105035781810151838201526020016104eb565b50505050905090810190601f1680156105305780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150505050565b6002600554600160601b900460ff16600281111561055e57fe5b1461059a5760405162461bcd60e51b8152600401808060200182810382526030815260200180610a886030913960400191505060405180910390fd5b60055467ffffffffffffffff1643116105f4576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b6105fc6108ef565b604080516000815290517fd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c9181900360200190a1565b73__$f5eea941ded5358daea4da7ea13a2128fd$__6392dbcf206000868686866040518663ffffffff1660e01b815260040180868152602001858152602001806020018463ffffffff1663ffffffff16815260200183600260200280838360005b838110156106ab578181015183820152602001610693565b50505050905001828103825285818151815260200191508051906020019060200280838360005b838110156106ea5781810151838201526020016106d2565b50505050905001965050505050505060006040518083038186803b15801561071157600080fd5b505af4158015610725573d6000803e3d6000fd5b5050505050505050565b73__$f5eea941ded5358daea4da7ea13a2128fd$__63110112ae6000868686866040518663ffffffff1660e01b81526004018086815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156107b357818101518382015260200161079b565b50505050905090810190601f1680156107e05780820380516001836020036101000a031916815260200191505b50965050505050505060006040518083038186803b15801561071157600080fd5b6001600554600160601b900460ff16600281111561081b57fe5b146108575760405162461bcd60e51b815260040180806020018281038252602e815260200180610a5a602e913960400191505060405180910390fd5b60055467ffffffffffffffff1643116108b1576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b6108b96109d1565b604080516001815290517fd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c9181900360200190a1565b60008054604080518082018252600280546001600160801b03808216600160801b909204811692909204011681526020810193909352516308b0246f60e21b81526001600160a01b03909116916322c091bc9160039190600481019060440183825b81546001600160a01b031681526001909101906020018083116109515750839050604080838360005b8381101561099257818101518382015260200161097a565b5050505090500192505050600060405180830381600087803b1580156109b757600080fd5b505af11580156109cb573d6000803e3d6000fd5b50505050565b6000805460408051808201825292835260028054600160801b81046001600160801b039081169181169290920401166020840152516308b0246f60e21b8152600380546001600160a01b03908116600480850191825291909416946322c091bc949293909291604482019160240180831161095157505082518152826040808383602061097a56fe43616e206f6e6c792074696d65206f7574206173736572746572206966206974206973207468656972207475726e43616e206f6e6c792074696d65206f7574206368616c6c656e676572206966206974206973207468656972207475726ea265627a7a7231582085efa582387c2a2fc66bdbb33cfb6cb6f3bdcd51c8b1ae045c1e404264c05fdb64736f6c634300050d0032"

// DeployArbChallenge deploys a new Ethereum contract, binding an instance of ArbChallenge to it.
func DeployArbChallenge(auth *bind.TransactOpts, backend bind.ContractBackend, vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (common.Address, *types.Transaction, *ArbChallenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	oneStepProofAddr, _, _, _ := DeployOneStepProof(auth, backend)
	ArbChallengeBin = strings.Replace(ArbChallengeBin, "__$f55f7f918072b72dcc999cdc8e581605f5$__", oneStepProofAddr.String()[2:], -1)

	bisectionAddr, _, _, _ := DeployBisection(auth, backend)
	ArbChallengeBin = strings.Replace(ArbChallengeBin, "__$f5eea941ded5358daea4da7ea13a2128fd$__", bisectionAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbChallengeBin), backend, vmAddress, _players, _escrows, _challengePeriod, _challengeRoot)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbChallenge{ArbChallengeCaller: ArbChallengeCaller{contract: contract}, ArbChallengeTransactor: ArbChallengeTransactor{contract: contract}, ArbChallengeFilterer: ArbChallengeFilterer{contract: contract}}, nil
}

// ArbChallenge is an auto generated Go binding around an Ethereum contract.
type ArbChallenge struct {
	ArbChallengeCaller     // Read-only binding to the contract
	ArbChallengeTransactor // Write-only binding to the contract
	ArbChallengeFilterer   // Log filterer for contract events
}

// ArbChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbChallengeSession struct {
	Contract     *ArbChallenge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbChallengeCallerSession struct {
	Contract *ArbChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ArbChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbChallengeTransactorSession struct {
	Contract     *ArbChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ArbChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbChallengeRaw struct {
	Contract *ArbChallenge // Generic contract binding to access the raw methods on
}

// ArbChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbChallengeCallerRaw struct {
	Contract *ArbChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ArbChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbChallengeTransactorRaw struct {
	Contract *ArbChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbChallenge creates a new instance of ArbChallenge, bound to a specific deployed contract.
func NewArbChallenge(address common.Address, backend bind.ContractBackend) (*ArbChallenge, error) {
	contract, err := bindArbChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbChallenge{ArbChallengeCaller: ArbChallengeCaller{contract: contract}, ArbChallengeTransactor: ArbChallengeTransactor{contract: contract}, ArbChallengeFilterer: ArbChallengeFilterer{contract: contract}}, nil
}

// NewArbChallengeCaller creates a new read-only instance of ArbChallenge, bound to a specific deployed contract.
func NewArbChallengeCaller(address common.Address, caller bind.ContractCaller) (*ArbChallengeCaller, error) {
	contract, err := bindArbChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbChallengeCaller{contract: contract}, nil
}

// NewArbChallengeTransactor creates a new write-only instance of ArbChallenge, bound to a specific deployed contract.
func NewArbChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbChallengeTransactor, error) {
	contract, err := bindArbChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbChallengeTransactor{contract: contract}, nil
}

// NewArbChallengeFilterer creates a new log filterer instance of ArbChallenge, bound to a specific deployed contract.
func NewArbChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbChallengeFilterer, error) {
	contract, err := bindArbChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbChallengeFilterer{contract: contract}, nil
}

// bindArbChallenge binds a generic wrapper to an already deployed contract.
func bindArbChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbChallenge *ArbChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbChallenge.Contract.ArbChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbChallenge *ArbChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChallenge.Contract.ArbChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbChallenge *ArbChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbChallenge.Contract.ArbChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbChallenge *ArbChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbChallenge *ArbChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbChallenge *ArbChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbChallenge.Contract.contract.Transact(opts, method, params...)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_ArbChallenge *ArbChallengeTransactor) AsserterTimedOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "asserterTimedOut")
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_ArbChallenge *ArbChallengeSession) AsserterTimedOut() (*types.Transaction, error) {
	return _ArbChallenge.Contract.AsserterTimedOut(&_ArbChallenge.TransactOpts)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_ArbChallenge *ArbChallengeTransactorSession) AsserterTimedOut() (*types.Transaction, error) {
	return _ArbChallenge.Contract.AsserterTimedOut(&_ArbChallenge.TransactOpts)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x71f45bbf.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, bytes32[] _afterHashAndMessageAndLogsBisections, uint32 _totalSteps, uint64[2] _timeBounds) returns()
func (_ArbChallenge *ArbChallengeTransactor) BisectAssertion(opts *bind.TransactOpts, _beforeInbox [32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "bisectAssertion", _beforeInbox, _afterHashAndMessageAndLogsBisections, _totalSteps, _timeBounds)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x71f45bbf.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, bytes32[] _afterHashAndMessageAndLogsBisections, uint32 _totalSteps, uint64[2] _timeBounds) returns()
func (_ArbChallenge *ArbChallengeSession) BisectAssertion(_beforeInbox [32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChallenge.Contract.BisectAssertion(&_ArbChallenge.TransactOpts, _beforeInbox, _afterHashAndMessageAndLogsBisections, _totalSteps, _timeBounds)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x71f45bbf.
//
// Solidity: function bisectAssertion(bytes32 _beforeInbox, bytes32[] _afterHashAndMessageAndLogsBisections, uint32 _totalSteps, uint64[2] _timeBounds) returns()
func (_ArbChallenge *ArbChallengeTransactorSession) BisectAssertion(_beforeInbox [32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChallenge.Contract.BisectAssertion(&_ArbChallenge.TransactOpts, _beforeInbox, _afterHashAndMessageAndLogsBisections, _totalSteps, _timeBounds)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_ArbChallenge *ArbChallengeTransactor) ChallengerTimedOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "challengerTimedOut")
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_ArbChallenge *ArbChallengeSession) ChallengerTimedOut() (*types.Transaction, error) {
	return _ArbChallenge.Contract.ChallengerTimedOut(&_ArbChallenge.TransactOpts)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_ArbChallenge *ArbChallengeTransactorSession) ChallengerTimedOut() (*types.Transaction, error) {
	return _ArbChallenge.Contract.ChallengerTimedOut(&_ArbChallenge.TransactOpts)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ArbChallenge *ArbChallengeTransactor) ContinueChallenge(opts *bind.TransactOpts, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "continueChallenge", _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ArbChallenge *ArbChallengeSession) ContinueChallenge(_assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.ContinueChallenge(&_ArbChallenge.TransactOpts, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ArbChallenge *ArbChallengeTransactorSession) ContinueChallenge(_assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.ContinueChallenge(&_ArbChallenge.TransactOpts, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x1d7aaea9.
//
// Solidity: function oneStepProof(bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes32[5] _afterHashAndMessages, bytes _proof) returns()
func (_ArbChallenge *ArbChallengeTransactor) OneStepProof(opts *bind.TransactOpts, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _afterHashAndMessages [5][32]byte, _proof []byte) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "oneStepProof", _beforeHashAndInbox, _timeBounds, _afterHashAndMessages, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x1d7aaea9.
//
// Solidity: function oneStepProof(bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes32[5] _afterHashAndMessages, bytes _proof) returns()
func (_ArbChallenge *ArbChallengeSession) OneStepProof(_beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _afterHashAndMessages [5][32]byte, _proof []byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.OneStepProof(&_ArbChallenge.TransactOpts, _beforeHashAndInbox, _timeBounds, _afterHashAndMessages, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x1d7aaea9.
//
// Solidity: function oneStepProof(bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes32[5] _afterHashAndMessages, bytes _proof) returns()
func (_ArbChallenge *ArbChallengeTransactorSession) OneStepProof(_beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _afterHashAndMessages [5][32]byte, _proof []byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.OneStepProof(&_ArbChallenge.TransactOpts, _beforeHashAndInbox, _timeBounds, _afterHashAndMessages, _proof)
}

// ArbChallengeBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ArbChallenge contract.
type ArbChallengeBisectedAssertionIterator struct {
	Event *ArbChallengeBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChallengeBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeBisectedAssertion)
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
		it.Event = new(ArbChallengeBisectedAssertion)
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
func (it *ArbChallengeBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeBisectedAssertion represents a BisectedAssertion event raised by the ArbChallenge contract.
type ArbChallengeBisectedAssertion struct {
	Bisecter                             common.Address
	AfterHashAndMessageAndLogsBisections [][32]byte
	TotalSteps                           uint32
	Deadline                             uint64
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0xd8fd4c0d938111394281c7239c621322f6397e9cbcf45e6fc552f17994966863.
//
// Solidity: event BisectedAssertion(address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) FilterBisectedAssertion(opts *bind.FilterOpts) (*ArbChallengeBisectedAssertionIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeBisectedAssertionIterator{contract: _ArbChallenge.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0xd8fd4c0d938111394281c7239c621322f6397e9cbcf45e6fc552f17994966863.
//
// Solidity: event BisectedAssertion(address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ArbChallengeBisectedAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeBisectedAssertion)
				if err := _ArbChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0xd8fd4c0d938111394281c7239c621322f6397e9cbcf45e6fc552f17994966863.
//
// Solidity: event BisectedAssertion(address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) ParseBisectedAssertion(log types.Log) (*ArbChallengeBisectedAssertion, error) {
	event := new(ArbChallengeBisectedAssertion)
	if err := _ArbChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the ArbChallenge contract.
type ArbChallengeContinuedChallengeIterator struct {
	Event *ArbChallengeContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *ArbChallengeContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeContinuedChallenge)
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
		it.Event = new(ArbChallengeContinuedChallenge)
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
func (it *ArbChallengeContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeContinuedChallenge represents a ContinuedChallenge event raised by the ArbChallenge contract.
type ArbChallengeContinuedChallenge struct {
	Challenger     common.Address
	AssertionIndex *big.Int
	Deadline       uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0x9e8153d5e3460213b94c7ddab1ab9aef35bf5a5bbf29f198ae8c142e155c46f1.
//
// Solidity: event ContinuedChallenge(address challenger, uint256 assertionIndex, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) FilterContinuedChallenge(opts *bind.FilterOpts) (*ArbChallengeContinuedChallengeIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeContinuedChallengeIterator{contract: _ArbChallenge.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0x9e8153d5e3460213b94c7ddab1ab9aef35bf5a5bbf29f198ae8c142e155c46f1.
//
// Solidity: event ContinuedChallenge(address challenger, uint256 assertionIndex, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *ArbChallengeContinuedChallenge) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeContinuedChallenge)
				if err := _ArbChallenge.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0x9e8153d5e3460213b94c7ddab1ab9aef35bf5a5bbf29f198ae8c142e155c46f1.
//
// Solidity: event ContinuedChallenge(address challenger, uint256 assertionIndex, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) ParseContinuedChallenge(log types.Log) (*ArbChallengeContinuedChallenge, error) {
	event := new(ArbChallengeContinuedChallenge)
	if err := _ArbChallenge.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ArbChallenge contract.
type ArbChallengeInitiatedChallengeIterator struct {
	Event *ArbChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ArbChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeInitiatedChallenge)
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
		it.Event = new(ArbChallengeInitiatedChallenge)
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
func (it *ArbChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the ArbChallenge contract.
type ArbChallengeInitiatedChallenge struct {
	Challenger common.Address
	Deadline   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x6dc74e1677661f103d909b4e12489baf51a49c7baf11ba1d4bf1a9fc899e9a0e.
//
// Solidity: event InitiatedChallenge(address challenger, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ArbChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeInitiatedChallengeIterator{contract: _ArbChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x6dc74e1677661f103d909b4e12489baf51a49c7baf11ba1d4bf1a9fc899e9a0e.
//
// Solidity: event InitiatedChallenge(address challenger, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ArbChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeInitiatedChallenge)
				if err := _ArbChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x6dc74e1677661f103d909b4e12489baf51a49c7baf11ba1d4bf1a9fc899e9a0e.
//
// Solidity: event InitiatedChallenge(address challenger, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ArbChallengeInitiatedChallenge, error) {
	event := new(ArbChallengeInitiatedChallenge)
	if err := _ArbChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ArbChallenge contract.
type ArbChallengeOneStepProofCompletedIterator struct {
	Event *ArbChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ArbChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeOneStepProofCompleted)
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
		it.Event = new(ArbChallengeOneStepProofCompleted)
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
func (it *ArbChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the ArbChallenge contract.
type ArbChallengeOneStepProofCompleted struct {
	Asserter common.Address
	Proof    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x1a96858c84fb221338517840d882a8fd1434f2cbbea7738d6e70333a801231a8.
//
// Solidity: event OneStepProofCompleted(address asserter, bytes proof)
func (_ArbChallenge *ArbChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*ArbChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeOneStepProofCompletedIterator{contract: _ArbChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x1a96858c84fb221338517840d882a8fd1434f2cbbea7738d6e70333a801231a8.
//
// Solidity: event OneStepProofCompleted(address asserter, bytes proof)
func (_ArbChallenge *ArbChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ArbChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeOneStepProofCompleted)
				if err := _ArbChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0x1a96858c84fb221338517840d882a8fd1434f2cbbea7738d6e70333a801231a8.
//
// Solidity: event OneStepProofCompleted(address asserter, bytes proof)
func (_ArbChallenge *ArbChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*ArbChallengeOneStepProofCompleted, error) {
	event := new(ArbChallengeOneStepProofCompleted)
	if err := _ArbChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeTimedOutChallengeIterator is returned from FilterTimedOutChallenge and is used to iterate over the raw logs and unpacked data for TimedOutChallenge events raised by the ArbChallenge contract.
type ArbChallengeTimedOutChallengeIterator struct {
	Event *ArbChallengeTimedOutChallenge // Event containing the contract specifics and raw log

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
func (it *ArbChallengeTimedOutChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeTimedOutChallenge)
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
		it.Event = new(ArbChallengeTimedOutChallenge)
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
func (it *ArbChallengeTimedOutChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeTimedOutChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeTimedOutChallenge represents a TimedOutChallenge event raised by the ArbChallenge contract.
type ArbChallengeTimedOutChallenge struct {
	ChallengerWrong bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedOutChallenge is a free log retrieval operation binding the contract event 0xd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c.
//
// Solidity: event TimedOutChallenge(bool challengerWrong)
func (_ArbChallenge *ArbChallengeFilterer) FilterTimedOutChallenge(opts *bind.FilterOpts) (*ArbChallengeTimedOutChallengeIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "TimedOutChallenge")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeTimedOutChallengeIterator{contract: _ArbChallenge.contract, event: "TimedOutChallenge", logs: logs, sub: sub}, nil
}

// WatchTimedOutChallenge is a free log subscription operation binding the contract event 0xd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c.
//
// Solidity: event TimedOutChallenge(bool challengerWrong)
func (_ArbChallenge *ArbChallengeFilterer) WatchTimedOutChallenge(opts *bind.WatchOpts, sink chan<- *ArbChallengeTimedOutChallenge) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "TimedOutChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeTimedOutChallenge)
				if err := _ArbChallenge.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
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

// ParseTimedOutChallenge is a log parse operation binding the contract event 0xd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c.
//
// Solidity: event TimedOutChallenge(bool challengerWrong)
func (_ArbChallenge *ArbChallengeFilterer) ParseTimedOutChallenge(log types.Log) (*ArbChallengeTimedOutChallenge, error) {
	event := new(ArbChallengeTimedOutChallenge)
	if err := _ArbChallenge.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbMachineABI is the input ABI used to generate the binding from.
const ArbMachineABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"instructionStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"errHandlerHash\",\"type\":\"bytes32\"}],\"name\":\"machineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbMachineFuncSigs maps the 4-byte function signature to its string representation.
var ArbMachineFuncSigs = map[string]string{
	"c1355b59": "machineHash(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32)",
}

// ArbMachineBin is the compiled bytecode used for deploying new contracts.
var ArbMachineBin = "0x6101d6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063c1355b591461003a575b600080fd5b610075600480360360c081101561005057600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610087565b60408051918252519081900360200190f35b604080516101008101825260e081018881528152815160208181018452888252808301919091528251808201845287815282840152825180820184528681526060830152825180820184528581526080830152825190810190925282825260a0810191909152600060c08201819052906101009061010b565b979650505050505050565b600060028260c0015114156101225750600061019c565b60018260c0015114156101375750600161019c565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101205b91905056fea265627a7a7231582043036cd0ca19bbb2e4fe99d7d83ada2376007667bb60d5411626bcbcc6b4cd9e64736f6c634300050d0032"

// DeployArbMachine deploys a new Ethereum contract, binding an instance of ArbMachine to it.
func DeployArbMachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbMachine, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbMachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbMachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbMachine{ArbMachineCaller: ArbMachineCaller{contract: contract}, ArbMachineTransactor: ArbMachineTransactor{contract: contract}, ArbMachineFilterer: ArbMachineFilterer{contract: contract}}, nil
}

// ArbMachine is an auto generated Go binding around an Ethereum contract.
type ArbMachine struct {
	ArbMachineCaller     // Read-only binding to the contract
	ArbMachineTransactor // Write-only binding to the contract
	ArbMachineFilterer   // Log filterer for contract events
}

// ArbMachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbMachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbMachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbMachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbMachineSession struct {
	Contract     *ArbMachine       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbMachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbMachineCallerSession struct {
	Contract *ArbMachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbMachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbMachineTransactorSession struct {
	Contract     *ArbMachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbMachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbMachineRaw struct {
	Contract *ArbMachine // Generic contract binding to access the raw methods on
}

// ArbMachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbMachineCallerRaw struct {
	Contract *ArbMachineCaller // Generic read-only contract binding to access the raw methods on
}

// ArbMachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbMachineTransactorRaw struct {
	Contract *ArbMachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbMachine creates a new instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachine(address common.Address, backend bind.ContractBackend) (*ArbMachine, error) {
	contract, err := bindArbMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbMachine{ArbMachineCaller: ArbMachineCaller{contract: contract}, ArbMachineTransactor: ArbMachineTransactor{contract: contract}, ArbMachineFilterer: ArbMachineFilterer{contract: contract}}, nil
}

// NewArbMachineCaller creates a new read-only instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineCaller(address common.Address, caller bind.ContractCaller) (*ArbMachineCaller, error) {
	contract, err := bindArbMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbMachineCaller{contract: contract}, nil
}

// NewArbMachineTransactor creates a new write-only instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbMachineTransactor, error) {
	contract, err := bindArbMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbMachineTransactor{contract: contract}, nil
}

// NewArbMachineFilterer creates a new log filterer instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbMachineFilterer, error) {
	contract, err := bindArbMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbMachineFilterer{contract: contract}, nil
}

// bindArbMachine binds a generic wrapper to an already deployed contract.
func bindArbMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbMachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbMachine *ArbMachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbMachine.Contract.ArbMachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbMachine *ArbMachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbMachine.Contract.ArbMachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbMachine *ArbMachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbMachine.Contract.ArbMachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbMachine *ArbMachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbMachine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbMachine *ArbMachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbMachine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbMachine *ArbMachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbMachine.Contract.contract.Transact(opts, method, params...)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineCaller) MachineHash(opts *bind.CallOpts, instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbMachine.contract.Call(opts, out, "machineHash", instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
	return *ret0, err
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _ArbMachine.Contract.MachineHash(&_ArbMachine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineCallerSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _ArbMachine.Contract.MachineHash(&_ArbMachine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// ArbProtocolABI is the input ABI used to generate the binding from.
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"7ddf59d6": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32)",
	"e83f4bfe": "generateLastMessageHash(bytes)",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"85ecb92a": "generatePreconditionHash(bytes32,uint64[2],bytes32)",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x610a6e610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100555760003560e01c80624c28f61461005a5780637ddf59d6146100b257806385ecb92a146100f3578063e83f4bfe14610148575b600080fd5b6100a06004803603608081101561007057600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b03166101ee565b60408051918252519081900360200190f35b6100a0600480360360c08110156100c857600080fd5b5080359063ffffffff6020820135169060408101359060608101359060808101359060a001356102e0565b6100a06004803603608081101561010957600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091945050903591506103389050565b6100a06004803603602081101561015e57600080fd5b81019060208101813564010000000081111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460018302840111640100000000831117156101ad57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061038c945050505050565b60408051600480825260a0820190925260009160609190816020015b6102126109d2565b81526020019060019003908161020a579050509050610230866104d1565b8160008151811061023d57fe5b602002602001018190525061025a836001600160a01b031661054f565b8160018151811061026757fe5b602002602001018190525061027b8461054f565b8160028151811061028857fe5b60209081029190910101526102aa6affffffffffffffffffffff19861661054f565b816003815181106102b757fe5b60200260200101819052506102d36102ce826105cd565b61067d565b519150505b949350505050565b6040805160208082019890985260e09690961b6001600160e01b0319168682015260448601949094526064850192909252608484015260a4808401919091528151808403909101815260c49092019052805191012090565b815160209283015160408051808601969096526001600160c01b031960c093841b8116878301529190921b166048850152605080850192909252805180850390920182526070909301909252815191012090565b8051600090819081908190815b818110156104c45773__$d969135829891f807aa9c34494da4ecd99$__6389df40da88866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561040f5781810151838201526020016103f7565b50505050905090810190601f16801561043c5780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561045957600080fd5b505af415801561046d573d6000803e3d6000fd5b505050506040513d604081101561048357600080fd5b50805160209182015160408051808501999099528881018290528051808a038201815260609099019052875197909201969096209594509250600101610399565b509293505050505b919050565b6104d96109d2565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161053e565b61052b6109d2565b8152602001906001900390816105235790505b508152600260209091015292915050565b6105576109d2565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916105bc565b6105a96109d2565b8152602001906001900390816105a15790505b508152600060209091015292915050565b6105d56109d2565b6105df82516107b3565b610630576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b610685610a00565b6060820151600c60ff909116106106d7576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166107045760405180602001604052806106fb84600001516107ba565b905290506104cc565b606082015160ff166001141561074b5760405180602001604052806106fb8460200151600001518560200151604001518660200151606001518760200151602001516107de565b606082015160ff166002141561077057506040805160208101909152815181526104cc565b600360ff16826060015160ff161015801561079457506060820151600c60ff909116105b156107b15760405180602001604052806106fb8460400151610886565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315610838575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206102d8565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b60006008825111156108d6576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610903578160200160208202803883390190505b50805190915060005b8181101561095f5761091c610a00565b61093886838151811061092b57fe5b602002602001015161067d565b9050806000015184838151811061094b57fe5b60209081029190910101525060010161090c565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156109a8578181015183820152602001610990565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6040518060800160405280600081526020016109ec610a12565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a723158201b90d05c0f27da9305a94b86c30470b81dc58f707f5afd0bd5c7cb8b4160e70664736f6c634300050d0032"

// DeployArbProtocol deploys a new Ethereum contract, binding an instance of ArbProtocol to it.
func DeployArbProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbProtocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbProtocolBin = strings.Replace(ArbProtocolBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

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

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x7ddf59d6.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x7ddf59d6.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x7ddf59d6.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _messages []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHash", _messages)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _messages)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _messages)
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

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// ArbValueABI is the input ABI used to generate the binding from.
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidValueHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeValueHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValidValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"immediate\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueFuncSigs maps the 4-byte function signature to its string representation.
var ArbValueFuncSigs = map[string]string{
	"32e6cc21": "deserializeMessage(bytes,uint256)",
	"89df40da": "deserializeValidValueHash(bytes,uint256)",
	"8f346036": "deserializeValueHash(bytes)",
	"1f3d4d4e": "getNextValidValue(bytes,uint256)",
	"826513e0": "hashCodePoint(uint8,bool,bytes32,bytes32)",
	"53409fab": "hashCodePointBasicValue(uint8,bytes32)",
	"264f384b": "hashCodePointImmediateValue(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"1667b411": "hashIntValue(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ArbValueBin is the compiled bytecode used for deploying new contracts.
var ArbValueBin = "0x61152b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c806353409fab1161007057806353409fab14610381578063826513e0146103a757806389df40da146103db5780638f3460361461049a578063b2b9dc621461053e576100a8565b80631667b411146100ad5780631f3d4d4e146100dc578063264f384b1461020157806332e6cc211461022d578063364df27714610379575b600080fd5b6100ca600480360360208110156100c357600080fd5b503561056f565b60408051918252519081900360200190f35b610182600480360360408110156100f257600080fd5b810190602081018135600160201b81111561010c57600080fd5b82018360208201111561011e57600080fd5b803590602001918460018302840111600160201b8311171561013f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610595915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101c55781810151838201526020016101ad565b50505050905090810190601f1680156101f25780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100ca6004803603606081101561021757600080fd5b5060ff8135169060208101359060400135610619565b6102d36004803603604081101561024357600080fd5b810190602081018135600160201b81111561025d57600080fd5b82018360208201111561026f57600080fd5b803590602001918460018302840111600160201b8311171561029057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061066b915050565b604051808815151515815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610338578181015183820152602001610320565b50505050905090810190601f1680156103655780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6100ca610869565b6100ca6004803603604081101561039757600080fd5b5060ff81351690602001356108dc565b6100ca600480360360808110156103bd57600080fd5b5060ff81351690602081013515159060408101359060600135610923565b610481600480360360408110156103f157600080fd5b810190602081018135600160201b81111561040b57600080fd5b82018360208201111561041d57600080fd5b803590602001918460018302840111600160201b8311171561043e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506109cc915050565b6040805192835260208301919091528051918290030190f35b6100ca600480360360208110156104b057600080fd5b810190602081018135600160201b8111156104ca57600080fd5b8201836020820111156104dc57600080fd5b803590602001918460018302840111600160201b831117156104fd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a45945050505050565b61055b6004803603602081101561055457600080fd5b5035610ab7565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b600060606000806105a461146f565b6105ae8787610abe565b9194509250905082156105f6576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b8161060a888880840363ffffffff610c4816565b945094505050505b9250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6000806000806000806060600088965060008a888151811061068957fe5b016020015160019098019760f81c9050600781146106bb576106ae8b60018a036109cc565b909850965061085d915050565b6106c58b896109cc565b90985091506106e48b60018c016000198d8c030163ffffffff610c4816565b92508a88815181106106f257fe5b016020015160019098019760f81c90508015610715576106ae8b60018a036109cc565b61071f8b89610cc8565b80995081975050508a888151811061073357fe5b016020015160019098019760f81c90508015610756576106ae8b60018a036109cc565b6107608b89610cc8565b80995081965050508a888151811061077457fe5b016020015160019098019760f81c90508015610797576106ae8b60018a036109cc565b6107a18b89610cc8565b60408051600480825260a0820190925260019c50919a509195506060916020820160808038833901905050905082816000815181106107dc57fe5b6020026020010181815250506107f18761056f565b816001815181106107fe57fe5b6020026020010181815250506108138661056f565b8160028151811061082057fe5b6020026020010181815250506108358561056f565b8160038151811061084257fe5b60200260200101818152505061085781610cef565b97505050505b92959891949750929550565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156108b557818101518382015260200161089d565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000831561097d575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206109c4565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b6000806000806109da61146f565b6109e48787610abe565b919450925090508215610a2c576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b81610a3682610daf565b51909890975095505050505050565b60008080610a5161146f565b610a5c856000610abe565b919450925090508215610aa4576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b610aad81610daf565b5195945050505050565b6008101590565b600080610ac961146f565b84518410610b1e576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110610b3157fe5b016020015160019092019160f81c90506000610b4b61149d565b60ff8316610b7f57610b5d8985610cc8565b9094509150600084610b6e84610ee5565b91985096509450610c419350505050565b60ff831660011415610ba657610b958985610f63565b9094509050600084610b6e836110be565b60ff831660021415610bcd57610bbc8985610cc8565b9094509150600084610b6e8461111e565b600360ff841610801590610be45750600c60ff8416105b15610c2157600219830160606000610bfd838d8961119c565b909850925090508087610c0f84611257565b99509950995050505050505050610c41565b8260ff16612710016000610c356000610ee5565b91985096509450505050505b9250925092565b606081830184511015610c5a57600080fd5b606082158015610c7557604051915060208201604052610cbf565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610cae578051835260209283019201610c96565b5050858452601f01601f1916604052505b50949350505050565b6000808281610cdd868363ffffffff61130716565b60209290920196919550909350505050565b6000600882511115610d3f576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b8151600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610d87578181015183820152602001610d6f565b5050505090500192505050604051602081830303815290604052805190602001209050919050565b610db76114c4565b6060820151600c60ff90911610610e09576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610e36576040518060200160405280610e2d846000015161056f565b90529050610590565b606082015160ff1660011415610e7d576040518060200160405280610e2d846020015160000151856020015160400151866020015160600151876020015160200151610923565b606082015160ff1660021415610ea25750604080516020810190915281518152610590565b600360ff16826060015160ff1610158015610ec657506060820151600c60ff909116105b15610ee3576040518060200160405280610e2d8460400151611323565bfe5b610eed61146f565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610f52565b610f3f61146f565b815260200190600190039081610f375790505b508152600060209091015292915050565b6000610f6d61149d565b60008390506000858281518110610f8057fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110610fa657fe5b016020015160019384019360f89190911c915060009060ff84161415611032576000610fd061146f565b610fda8a87610abe565b90975090925090508115611023576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61102c81610daf565b51925050505b6000611044898663ffffffff61130716565b90506020850194508360ff1660011415611089576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506106129050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6110c661146f565b60408051608081018252600080825260208083018690528351828152908101845291928301919061110d565b6110fa61146f565b8152602001906001900390816110f25790505b508152600160209091015292915050565b61112661146f565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161118b565b61117861146f565b8152602001906001900390816111705790505b508152600260209091015292915050565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156111e757816020015b6111d461146f565b8152602001906001900390816111cc5790505b50905060005b8960ff168160ff161015611241576112058985610abe565b8451859060ff861690811061121657fe5b60209081029190910101529450925082156112395750909450909250905061124e565b6001016111ed565b5060009550919350909150505b93509350939050565b61125f61146f565b6112698251610ab7565b6112ba576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6000816020018351101561131a57600080fd5b50016020015190565b6000600882511115611373576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156113a0578160200160208202803883390190505b50805190915060005b818110156113fc576113b96114c4565b6113d58683815181106113c857fe5b6020026020010151610daf565b905080600001518483815181106113e857fe5b6020908102919091010152506001016113a9565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561144557818101518382015260200161142d565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180608001604052806000815260200161148961149d565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4d61727368616c6c65642076616c7565206d7573742062652076616c69640000a265627a7a72315820184abb01cd16a80d27b9e3125b6e33c2f53720dceececcf7c7182d2f0bcee9ee64736f6c634300050d0032"

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

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_ArbValue *ArbValueCaller) DeserializeMessage(opts *bind.CallOpts, data []byte, startOffset *big.Int) (struct {
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
	err := _ArbValue.contract.Call(opts, out, "deserializeMessage", data, startOffset)
	return *ret, err
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_ArbValue *ArbValueSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _ArbValue.Contract.DeserializeMessage(&_ArbValue.CallOpts, data, startOffset)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_ArbValue *ArbValueCallerSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _ArbValue.Contract.DeserializeMessage(&_ArbValue.CallOpts, data, startOffset)
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

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePoint(opts *bind.CallOpts, opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePoint", opcode, immediate, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePoint(&_ArbValue.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePoint(&_ArbValue.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
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
const ArbitrumVMABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeLauncherAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"ChallengeLaunched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeLauncher\",\"outputs\":[{\"internalType\":\"contractIChallengeLauncher\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumVM.State\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"enumVM.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"activeChallengeManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbitrumVMFuncSigs maps the 4-byte function signature to its string representation.
var ArbitrumVMFuncSigs = map[string]string{
	"94af716b": "activateVM()",
	"8aee8c79": "challengeLauncher()",
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"f2204f74": "confirmDisputableAsserted(bytes32,bytes32,uint32,bytes,bytes32)",
	"08dc89d7": "currentDeposit(address)",
	"aca0f372": "escrowRequired()",
	"6be00229": "exitAddress()",
	"1865c57d": "getState()",
	"d489113a": "globalInbox()",
	"2782e87e": "initiateChallenge(bytes32)",
	"8da5cb5b": "owner()",
	"cfa80707": "ownerShutdown()",
	"fec5a286": "pendingDisputableAssert(bytes32,bytes32,bytes32,bytes32,bytes32,uint32,uint64[2])",
	"60675a87": "terminateAddress()",
	"3a768463": "vm()",
}

// ArbitrumVMBin is the compiled bytecode used for deploying new contracts.
var ArbitrumVMBin = "0x608060405234801561001057600080fd5b50604051611686380380611686833981810160405260e081101561003357600080fd5b50805160208201516040808401516060850151608086015160a087015160c090970151600180546001600160a01b038084166001600160a01b0319928316179283905560008054828d16931692909217825587517ff39723830000000000000000000000000000000000000000000000000000000081529751999a9899969895979496939492169263f397238392600480820193929182900301818387803b1580156100de57600080fd5b505af11580156100f2573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b15801561017957600080fd5b505af415801561018d573d6000803e3d6000fd5b505050506040513d60208110156101a357600080fd5b50516004555050600680546001600160801b0319166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161763ffffffff60201b19166401000000009290931691909102919091179055506114788061020e6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638aee8c7911610097578063cfa8070711610066578063cfa807071461028e578063d489113a14610296578063f2204f741461029e578063fec5a2861461035f576100f5565b80638aee8c791461026e5780638da5cb5b1461027657806394af716b1461027e578063aca0f37214610286576100f5565b80632782e87e116100d35780632782e87e146101805780633a7684631461019d57806360675a87146102425780636be0022914610266576100f5565b806308dc89d7146100fa5780631865c57d1461013257806322c091bc1461015e575b600080fd5b6101206004803603602081101561011057600080fd5b50356001600160a01b03166103d6565b60408051918252519081900360200190f35b61013a6103f5565b6040518082600381111561014a57fe5b60ff16815260200191505060405180910390f35b61017e6004803603608081101561017457600080fd5b5060408101610405565b005b61017e6004803603602081101561019657600080fd5b5035610517565b6101a5610799565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561021557fe5b60ff1681526001600160a01b039092166020830152506040805191829003019a5098505050505050505050f35b61024a61080a565b604080516001600160a01b039092168252519081900360200190f35b61024a610819565b61024a610828565b61024a610837565b61017e610846565b6101206108d9565b61017e6108e8565b61024a610948565b61017e600480360360a08110156102b457600080fd5b81359160208101359163ffffffff60408301351691908101906080810160608201356401000000008111156102e857600080fd5b8201836020820111156102fa57600080fd5b8035906020019184600183028401116401000000008311171561031c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610957915050565b61017e600480360361010081101561037657600080fd5b604080518082018252833593602081013593838201359360608301359360808401359363ffffffff60a08201351693810192909161010083019160c08401906002908390839080828437600092019190915250919450610ab79350505050565b6001600160a01b0381166000908152600860205260409020545b919050565b600754600160401b900460ff1690565b600754600160481b90046001600160a01b031633146104555760405162461bcd60e51b815260040180806020018281038252602d8152602001806113e6602d913960400191505060405180910390fd5b60078054600160481b600160e81b03191690556104ba6001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002054610c1890919063ffffffff16565b82356001600160a01b031660009081526008602081815260408320939093556104f2928401356001600160801b03169185600161047d565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b031611156105705760405162461bcd60e51b81526004018080602001828103825260278152602001806113bf6027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b1580156105ea57600080fd5b505af41580156105fe573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b03168082529181019190915260075484516328d23fe960e01b81529290951697506328d23fe996509194919363ffffffff16928892600490920191829187918190849084905b8381101561069657818101518382015260200161067e565b5050505090500184600260200280838360005b838110156106c15781810151838201526020016106a9565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050602060405180830381600087803b15801561070057600080fd5b505af1158015610714573d6000803e3d6000fd5b505050506040513d602081101561072a57600080fd5b505160078054600160481b600160e81b031916600160481b6001600160a01b03938416810291909117918290556040805191909204909216825233602083015280517f65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f9281900390910190a150565b6002546003546004546005546006546007546001600160a01b03928316926001600160801b0383169267ffffffffffffffff600160801b8204811693600160c01b909204169163ffffffff8083169264010000000081049091169160ff600160401b83041691600160481b9004168b565b600a546001600160a01b031681565b6009546001600160a01b031681565b6000546001600160a01b031681565b600b546001600160a01b031681565b600b546001600160a01b0316331461089e576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff1660038111156108b857fe5b14156108d7576007805468ff00000000000000001916600160401b1790555b565b6006546001600160801b031690565b600b546001600160a01b03163314610940576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6108d7610c79565b6001546001600160a01b031681565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63ce9d5122600287878787876040518763ffffffff1660e01b8152600401808781526020018681526020018581526020018463ffffffff1663ffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156109ee5781810151838201526020016109d6565b50505050905090810190601f168015610a1b5780820380516001836020036101000a031916815260200191505b5097505050505050505060006040518083038186803b158015610a3d57600080fd5b505af4158015610a51573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610a8c935091506001600160801b031663ffffffff610c1816565b6005546001600160a01b0316600090815260086020526040902055610ab082610c87565b5050505050565b336000908152600860205260409020546006546001600160801b03161115610b105760405162461bcd60e51b81526004018080602001828103825260318152602001806114136031913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151636c36f28f60e11b8152600260048201818152602483018c9052604483018b9052606483018a90526084830189905260a4830188905263ffffffff871660c484015273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__9463d86de51e9492938d938d938d938d938d938d938d93909260e401918491908190849084905b83811015610bce578181015183820152602001610bb6565b505050509050019850505050505050505060006040518083038186803b158015610bf757600080fd5b505af4158015610c0b573d6000803e3d6000fd5b5050505050505050505050565b600082820183811015610c72576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600b546001600160a01b0316ff5b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b158015610ccd57600080fd5b505af1158015610ce1573d6000803e3d6000fd5b505050506040513d6020811015610cf757600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b158015610d4857600080fd5b505af4158015610d5c573d6000803e3d6000fd5b505050506040513d6020811015610d7257600080fd5b50518114610dbd57610db96040518060600160405280610d926001610e85565b8152602001610da46002800154610f03565b8152602001610db284610f03565b9052610f81565b6004555b60015460405163e4eb8c6360e01b81526020600482018181528551602484015285516001600160a01b039094169363e4eb8c6393879383926044909201919085019080838360005b83811015610e1d578181015183820152602001610e05565b50505050905090810190601f168015610e4a5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610e6957600080fd5b505af1158015610e7d573d6000803e3d6000fd5b505050505050565b610e8d611357565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610ef2565b610edf611357565b815260200190600190039081610ed75790505b508152600060209091015292915050565b610f0b611357565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610f70565b610f5d611357565b815260200190600190039081610f555790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b610fa5611357565b815260200190600190039081610f9d575050805190915060005b81811015610ff757848160038110610fd357fe5b6020020151838281518110610fe457fe5b6020908102919091010152600101610fbf565b5061100182611009565b949350505050565b6000600882511115611059576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611086578160200160208202803883390190505b50805190915060005b818110156110e25761109f611385565b6110bb8683815181106110ae57fe5b6020026020010151611155565b905080600001518483815181106110ce57fe5b60209081029190910101525060010161108f565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561112b578181015183820152602001611113565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b61115d611385565b6060820151600c60ff909116106111af576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166111dc5760405180602001604052806111d3846000015161128b565b905290506103f0565b606082015160ff16600114156112235760405180602001604052806111d38460200151600001518560200151604001518660200151606001518760200151602001516112af565b606082015160ff166002141561124857506040805160208101909152815181526103f0565b600360ff16826060015160ff161015801561126c57506060820151600c60ff909116105b156112895760405180602001604052806111d38460400151611009565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b60008315611309575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611001565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b604051806080016040528060008152602001611371611397565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f7765644f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274a265627a7a723158200ef8330171e1c4206dd746e99e41b7192382b9180f5abbeb2ec773e5ec5100b164736f6c634300050d0032"

// DeployArbitrumVM deploys a new Ethereum contract, binding an instance of ArbitrumVM to it.
func DeployArbitrumVM(auth *bind.TransactOpts, backend bind.ContractBackend, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeLauncherAddress common.Address, _globalInboxAddress common.Address) (common.Address, *types.Transaction, *ArbitrumVM, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrumVMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ArbitrumVMBin = strings.Replace(ArbitrumVMBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbitrumVMBin = strings.Replace(ArbitrumVMBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbitrumVMBin), backend, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeLauncherAddress, _globalInboxAddress)
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

// ChallengeLauncher is a free data retrieval call binding the contract method 0x8aee8c79.
//
// Solidity: function challengeLauncher() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCaller) ChallengeLauncher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbitrumVM.contract.Call(opts, out, "challengeLauncher")
	return *ret0, err
}

// ChallengeLauncher is a free data retrieval call binding the contract method 0x8aee8c79.
//
// Solidity: function challengeLauncher() constant returns(address)
func (_ArbitrumVM *ArbitrumVMSession) ChallengeLauncher() (common.Address, error) {
	return _ArbitrumVM.Contract.ChallengeLauncher(&_ArbitrumVM.CallOpts)
}

// ChallengeLauncher is a free data retrieval call binding the contract method 0x8aee8c79.
//
// Solidity: function challengeLauncher() constant returns(address)
func (_ArbitrumVM *ArbitrumVMCallerSession) ChallengeLauncher() (common.Address, error) {
	return _ArbitrumVM.Contract.ChallengeLauncher(&_ArbitrumVM.CallOpts)
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
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbitrumVM *ArbitrumVMCaller) Vm(opts *bind.CallOpts) (struct {
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
	err := _ArbitrumVM.contract.Call(opts, out, "vm")
	return *ret, err
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbitrumVM *ArbitrumVMSession) Vm() (struct {
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
	return _ArbitrumVM.Contract.Vm(&_ArbitrumVM.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, address activeChallengeManager)
func (_ArbitrumVM *ArbitrumVMCallerSession) Vm() (struct {
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

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbitrumVM *ArbitrumVMTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.contract.Transact(opts, "confirmDisputableAsserted", _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbitrumVM *ArbitrumVMSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.ConfirmDisputableAsserted(&_ArbitrumVM.TransactOpts, _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbitrumVM *ArbitrumVMTransactorSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.ConfirmDisputableAsserted(&_ArbitrumVM.TransactOpts, _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
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

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbitrumVM *ArbitrumVMTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbitrumVM.contract.Transact(opts, "pendingDisputableAssert", _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbitrumVM *ArbitrumVMSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.PendingDisputableAssert(&_ArbitrumVM.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbitrumVM *ArbitrumVMTransactorSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbitrumVM.Contract.PendingDisputableAssert(&_ArbitrumVM.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
}

// ArbitrumVMChallengeLaunchedIterator is returned from FilterChallengeLaunched and is used to iterate over the raw logs and unpacked data for ChallengeLaunched events raised by the ArbitrumVM contract.
type ArbitrumVMChallengeLaunchedIterator struct {
	Event *ArbitrumVMChallengeLaunched // Event containing the contract specifics and raw log

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
func (it *ArbitrumVMChallengeLaunchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumVMChallengeLaunched)
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
		it.Event = new(ArbitrumVMChallengeLaunched)
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
func (it *ArbitrumVMChallengeLaunchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrumVMChallengeLaunchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrumVMChallengeLaunched represents a ChallengeLaunched event raised by the ArbitrumVM contract.
type ArbitrumVMChallengeLaunched struct {
	ChallengeContract common.Address
	Challenger        common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChallengeLaunched is a free log retrieval operation binding the contract event 0x65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f.
//
// Solidity: event ChallengeLaunched(address challengeContract, address challenger)
func (_ArbitrumVM *ArbitrumVMFilterer) FilterChallengeLaunched(opts *bind.FilterOpts) (*ArbitrumVMChallengeLaunchedIterator, error) {

	logs, sub, err := _ArbitrumVM.contract.FilterLogs(opts, "ChallengeLaunched")
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMChallengeLaunchedIterator{contract: _ArbitrumVM.contract, event: "ChallengeLaunched", logs: logs, sub: sub}, nil
}

// WatchChallengeLaunched is a free log subscription operation binding the contract event 0x65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f.
//
// Solidity: event ChallengeLaunched(address challengeContract, address challenger)
func (_ArbitrumVM *ArbitrumVMFilterer) WatchChallengeLaunched(opts *bind.WatchOpts, sink chan<- *ArbitrumVMChallengeLaunched) (event.Subscription, error) {

	logs, sub, err := _ArbitrumVM.contract.WatchLogs(opts, "ChallengeLaunched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrumVMChallengeLaunched)
				if err := _ArbitrumVM.contract.UnpackLog(event, "ChallengeLaunched", log); err != nil {
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
func (_ArbitrumVM *ArbitrumVMFilterer) ParseChallengeLaunched(log types.Log) (*ArbitrumVMChallengeLaunched, error) {
	event := new(ArbitrumVMChallengeLaunched)
	if err := _ArbitrumVM.contract.UnpackLog(event, "ChallengeLaunched", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Fields     [5][32]byte
	Asserter   common.Address
	TimeBounds [2]uint64
	NumSteps   uint32
	Deadline   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
func (_ArbitrumVM *ArbitrumVMFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*ArbitrumVMPendingDisputableAssertionIterator, error) {

	logs, sub, err := _ArbitrumVM.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMPendingDisputableAssertionIterator{contract: _ArbitrumVM.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
func (_ArbitrumVM *ArbitrumVMFilterer) ParsePendingDisputableAssertion(log types.Log) (*ArbitrumVMPendingDisputableAssertion, error) {
	event := new(ArbitrumVMPendingDisputableAssertion)
	if err := _ArbitrumVM.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionABI is the input ABI used to generate the binding from.
const BisectionABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"}]"

// BisectionFuncSigs maps the 4-byte function signature to its string representation.
var BisectionFuncSigs = map[string]string{
	"92dbcf20": "bisectAssertion(Challenge.Data storage,bytes32,bytes32[],uint32,uint64[2])",
	"110112ae": "continueChallenge(Challenge.Data storage,uint256,bytes,bytes32,bytes32)",
}

// BisectionBin is the compiled bytecode used for deploying new contracts.
var BisectionBin = "0x610dcc610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063110112ae1461004557806392dbcf201461010b575b600080fd5b81801561005157600080fd5b50610109600480360360a081101561006857600080fd5b81359160208101359181019060608101604082013564010000000081111561008f57600080fd5b8201836020820111156100a157600080fd5b803590602001918460018302840111640100000000831117156100c357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356101ff565b005b81801561011757600080fd5b50610109600480360360c081101561012e57600080fd5b81359160208101359181019060608101604082013564010000000081111561015557600080fd5b82018360208201111561016757600080fd5b8035906020019184602083028401116401000000008311171561018957600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080518082018252939663ffffffff863516969095909460608201945092506020019060029083908390808284376000920191909152509194506104e09350505050565b846001015482146102415760405162461bcd60e51b815260040180806020018281038252602b815260200180610d3e602b913960400191505060405180910390fd5b600585015467ffffffffffffffff164311156102a4576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60038501600101546001600160a01b031633146102f25760405162461bcd60e51b815260040180806020018281038252602f815260200180610d69602f913960400191505060405180910390fd5b73__$800fcb2f4a98daa165a5cdb21a355d7a15$__63b792d767848484886001016040518563ffffffff1660e01b81526004018080602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b83811015610371578181015183820152602001610359565b50505050905090810190601f16801561039e5780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b1580156103be57600080fd5b505af41580156103d2573d6000803e3d6000fd5b505050506040513d60208110156103e857600080fd5b505161043b576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420617373657274696f6e2073656c6563746564000000000000604482015290519081900360640190fd5b60058501805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160601b179384041601811667ffffffffffffffff19929092169190911791829055600187018390556004870154604080516001600160a01b03909216825260208201889052929091168183015290517f9e8153d5e3460213b94c7ddab1ab9aef35bf5a5bbf29f198ae8c142e155c46f19181900360600190a15050505050565b60016005860154600160601b900460ff1660028111156104fc57fe5b146105385760405162461bcd60e51b8152600401808060200182810382526034815260200180610d0a6034913960400191505060405180910390fd5b600585015467ffffffffffffffff1643111561059b576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60038501600001546001600160a01b031633146105ff576040805162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f7269676e616c2061737365727465722063616e20626973656374604482015290519081900360640190fd5b6000606061064c6040518060a001604052806001600389518161061e57fe5b040363ffffffff1681526020018781526020018663ffffffff1681526020018581526020018881525061086a565b6001890154919350915082146106a9576040805162461bcd60e51b815260206004820152601960248201527f446f6573206e6f74206d61746368207072657620737461746500000000000000604482015290519081900360640190fd5b60058701805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160611b1793840416011667ffffffffffffffff19919091161790556040516309898dc160e41b815260206004820181815283516024840152835173__$800fcb2f4a98daa165a5cdb21a355d7a15$__93639898dc1093869392839260440191858101910280838360005b83811015610750578181015183820152602001610738565b505050509050019250505060206040518083038186803b15801561077357600080fd5b505af4158015610787573d6000803e3d6000fd5b505050506040513d602081101561079d57600080fd5b505160018801557fd8fd4c0d938111394281c7239c621322f6397e9cbcf45e6fc552f1799496686360038801600001546005890154604080516001600160a01b0390931680845263ffffffff89169184019190915267ffffffffffffffff90911660608301819052608060208085018281528b51928601929092528a5193948b948b9493919260a0840191818801910280838360005b8381101561084b578181015183820152602001610833565b505050509050019550505050505060405180910390a150505050505050565b60006060600080846000015163ffffffff16856040015163ffffffff168161088e57fe5b046001019050846000015163ffffffff166040519080825280602002602001820160405280156108c8578160200160208202803883390190505b50925060005b855163ffffffff16811015610d0157856000015163ffffffff16866040015163ffffffff16816108fa57fe5b0663ffffffff1681141561091057600019909101905b73__$9836fa7140e5a33041d4b827682e675a30$__6385ecb92a8760200151836003028151811061093d57fe5b6020026020010151886060015189608001516040518463ffffffff1660e01b81526004018084815260200183600260200280838360005b8381101561098c578181015183820152602001610974565b50505050905001828152602001935050505060206040518083038186803b1580156109b657600080fd5b505af41580156109ca573d6000803e3d6000fd5b505050506040513d60208110156109e057600080fd5b505160208701518051919450849173__$9836fa7140e5a33041d4b827682e675a30$__91637ddf59d69160036001870102908110610a1a57fe5b6020026020010151858a602001518660030260010181518110610a3957fe5b60200260200101518b602001518760010160030260010181518110610a5a57fe5b60200260200101518c602001518860030260020181518110610a7857fe5b60200260200101518d602001518960010160030260020181518110610a9957fe5b60200260200101516040518763ffffffff1660e01b8152600401808781526020018663ffffffff1663ffffffff168152602001858152602001848152602001838152602001828152602001965050505050505060206040518083038186803b158015610b0457600080fd5b505af4158015610b18573d6000803e3d6000fd5b505050506040513d6020811015610b2e57600080fd5b505160408051602081810194909452808201929092528051808303820181526060909201905280519101208451859083908110610b6757fe5b602090810291909101015280610cf9578273__$9836fa7140e5a33041d4b827682e675a30$__637ddf59d68860200151896000015160030263ffffffff1681518110610baf57fe5b602002602001015189604001518a60200151600181518110610bcd57fe5b60200260200101518b602001518c6000015160030260010163ffffffff1681518110610bf557fe5b60200260200101518c60200151600281518110610c0e57fe5b60200260200101518d602001518e6000015160030260020163ffffffff1681518110610c3657fe5b60200260200101516040518763ffffffff1660e01b8152600401808781526020018663ffffffff1663ffffffff168152602001858152602001848152602001838152602001828152602001965050505050505060206040518083038186803b158015610ca157600080fd5b505af4158015610cb5573d6000803e3d6000fd5b505050506040513d6020811015610ccb57600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019052805191012094505b6001016108ce565b50505091509156fe43616e206f6e6c792062697365637420617373657274696f6e20696e20726573706f6e736520746f2061206368616c6c656e6765636f6e74696e75654368616c6c656e67653a20496e636f72726563742070726576696f75732073746174654f6e6c79206f726967696e616c206368616c6c656e6765722063616e20636f6e74696e7565206368616c6c656e6765a265627a7a7231582084289a0881c05434a1b1cec2c6f95a55840c0185293fd50b5ba83456eb16f0c564736f6c634300050d0032"

// DeployBisection deploys a new Ethereum contract, binding an instance of Bisection to it.
func DeployBisection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bisection, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	merkleLibAddr, _, _, _ := DeployMerkleLib(auth, backend)
	BisectionBin = strings.Replace(BisectionBin, "__$800fcb2f4a98daa165a5cdb21a355d7a15$__", merkleLibAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	BisectionBin = strings.Replace(BisectionBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BisectionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bisection{BisectionCaller: BisectionCaller{contract: contract}, BisectionTransactor: BisectionTransactor{contract: contract}, BisectionFilterer: BisectionFilterer{contract: contract}}, nil
}

// Bisection is an auto generated Go binding around an Ethereum contract.
type Bisection struct {
	BisectionCaller     // Read-only binding to the contract
	BisectionTransactor // Write-only binding to the contract
	BisectionFilterer   // Log filterer for contract events
}

// BisectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BisectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BisectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BisectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BisectionSession struct {
	Contract     *Bisection        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BisectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BisectionCallerSession struct {
	Contract *BisectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BisectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BisectionTransactorSession struct {
	Contract     *BisectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BisectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BisectionRaw struct {
	Contract *Bisection // Generic contract binding to access the raw methods on
}

// BisectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BisectionCallerRaw struct {
	Contract *BisectionCaller // Generic read-only contract binding to access the raw methods on
}

// BisectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BisectionTransactorRaw struct {
	Contract *BisectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBisection creates a new instance of Bisection, bound to a specific deployed contract.
func NewBisection(address common.Address, backend bind.ContractBackend) (*Bisection, error) {
	contract, err := bindBisection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bisection{BisectionCaller: BisectionCaller{contract: contract}, BisectionTransactor: BisectionTransactor{contract: contract}, BisectionFilterer: BisectionFilterer{contract: contract}}, nil
}

// NewBisectionCaller creates a new read-only instance of Bisection, bound to a specific deployed contract.
func NewBisectionCaller(address common.Address, caller bind.ContractCaller) (*BisectionCaller, error) {
	contract, err := bindBisection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionCaller{contract: contract}, nil
}

// NewBisectionTransactor creates a new write-only instance of Bisection, bound to a specific deployed contract.
func NewBisectionTransactor(address common.Address, transactor bind.ContractTransactor) (*BisectionTransactor, error) {
	contract, err := bindBisection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionTransactor{contract: contract}, nil
}

// NewBisectionFilterer creates a new log filterer instance of Bisection, bound to a specific deployed contract.
func NewBisectionFilterer(address common.Address, filterer bind.ContractFilterer) (*BisectionFilterer, error) {
	contract, err := bindBisection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BisectionFilterer{contract: contract}, nil
}

// bindBisection binds a generic wrapper to an already deployed contract.
func bindBisection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bisection *BisectionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bisection.Contract.BisectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bisection *BisectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bisection.Contract.BisectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bisection *BisectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bisection.Contract.BisectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bisection *BisectionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bisection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bisection *BisectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bisection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bisection *BisectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bisection.Contract.contract.Transact(opts, method, params...)
}

// BisectionBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the Bisection contract.
type BisectionBisectedAssertionIterator struct {
	Event *BisectionBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *BisectionBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionBisectedAssertion)
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
		it.Event = new(BisectionBisectedAssertion)
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
func (it *BisectionBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionBisectedAssertion represents a BisectedAssertion event raised by the Bisection contract.
type BisectionBisectedAssertion struct {
	Bisecter                             common.Address
	AfterHashAndMessageAndLogsBisections [][32]byte
	TotalSteps                           uint32
	Deadline                             uint64
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0xd8fd4c0d938111394281c7239c621322f6397e9cbcf45e6fc552f17994966863.
//
// Solidity: event BisectedAssertion(address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint64 deadline)
func (_Bisection *BisectionFilterer) FilterBisectedAssertion(opts *bind.FilterOpts) (*BisectionBisectedAssertionIterator, error) {

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return &BisectionBisectedAssertionIterator{contract: _Bisection.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0xd8fd4c0d938111394281c7239c621322f6397e9cbcf45e6fc552f17994966863.
//
// Solidity: event BisectedAssertion(address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint64 deadline)
func (_Bisection *BisectionFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *BisectionBisectedAssertion) (event.Subscription, error) {

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionBisectedAssertion)
				if err := _Bisection.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0xd8fd4c0d938111394281c7239c621322f6397e9cbcf45e6fc552f17994966863.
//
// Solidity: event BisectedAssertion(address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint64 deadline)
func (_Bisection *BisectionFilterer) ParseBisectedAssertion(log types.Log) (*BisectionBisectedAssertion, error) {
	event := new(BisectionBisectedAssertion)
	if err := _Bisection.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the Bisection contract.
type BisectionContinuedChallengeIterator struct {
	Event *BisectionContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *BisectionContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionContinuedChallenge)
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
		it.Event = new(BisectionContinuedChallenge)
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
func (it *BisectionContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionContinuedChallenge represents a ContinuedChallenge event raised by the Bisection contract.
type BisectionContinuedChallenge struct {
	Challenger     common.Address
	AssertionIndex *big.Int
	Deadline       uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0x9e8153d5e3460213b94c7ddab1ab9aef35bf5a5bbf29f198ae8c142e155c46f1.
//
// Solidity: event ContinuedChallenge(address challenger, uint256 assertionIndex, uint64 deadline)
func (_Bisection *BisectionFilterer) FilterContinuedChallenge(opts *bind.FilterOpts) (*BisectionContinuedChallengeIterator, error) {

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return &BisectionContinuedChallengeIterator{contract: _Bisection.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0x9e8153d5e3460213b94c7ddab1ab9aef35bf5a5bbf29f198ae8c142e155c46f1.
//
// Solidity: event ContinuedChallenge(address challenger, uint256 assertionIndex, uint64 deadline)
func (_Bisection *BisectionFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *BisectionContinuedChallenge) (event.Subscription, error) {

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionContinuedChallenge)
				if err := _Bisection.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0x9e8153d5e3460213b94c7ddab1ab9aef35bf5a5bbf29f198ae8c142e155c46f1.
//
// Solidity: event ContinuedChallenge(address challenger, uint256 assertionIndex, uint64 deadline)
func (_Bisection *BisectionFilterer) ParseContinuedChallenge(log types.Log) (*BisectionContinuedChallenge, error) {
	event := new(BisectionContinuedChallenge)
	if err := _Bisection.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820085bb168467a1b1b9144d77c586e271e631429488c107be4f03f7ca8146dd95064736f6c634300050d0032"

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
const ChainLauncherABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeLauncherAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"name\":\"ChainCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"launchChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChainLauncherFuncSigs maps the 4-byte function signature to its string representation.
var ChainLauncherFuncSigs = map[string]string{
	"e2b491e3": "launchChain(bytes32,uint32,uint32,uint128,address)",
}

// ChainLauncherBin is the compiled bytecode used for deploying new contracts.
var ChainLauncherBin = "0x608060405234801561001057600080fd5b506040516119b53803806119b58339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b0319918216179091556001805493909216921691909117905561193b8061007a6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e2b491e314610030575b600080fd5b610082600480360360a081101561004657600080fd5b50803590602081013563ffffffff9081169160408101359091169060608101356001600160801b031690608001356001600160a01b0316610084565b005b600154600080546040519192889288928892889288926001600160a01b039283169216906100b190610164565b96875263ffffffff9586166020880152939094166040808701919091526001600160801b0390921660608601526001600160a01b03908116608086015292831660a0850152911660c0830152519081900360e001906000f08015801561011b573d6000803e3d6000fd5b50604080516001600160a01b038316815290519192507fa8ee415251435dd34f5b8cc67a5659f0a26f5ed9f7a91f59c9016e799580457f919081900360200190a1505050505050565b611795806101728339019056fe608060405234801561001057600080fd5b50604051611795380380611795833981810160405260e081101561003357600080fd5b50805160208201516040808401516060850151608086015160a087015160c090970151600180546001600160a01b038084166001600160a01b0319928316179283905560008054828d16931692909217825587517ff39723830000000000000000000000000000000000000000000000000000000081529751999a989996989597949693948b948b948b948b948b948b948b949092169263f39723839260048084019382900301818387803b1580156100eb57600080fd5b505af11580156100ff573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b15801561018657600080fd5b505af415801561019a573d6000803e3d6000fd5b505050506040513d60208110156101b057600080fd5b50516004555050600680546001600160801b0319166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161763ffffffff60201b19166401000000009290931691909102919091178082556001925060ff60401b191668010000000000000000830217905550505050505050506115588061023d6000396000f3fe6080604052600436106100f35760003560e01c80638aee8c791161008a578063cfa8070711610059578063cfa8070714610323578063d489113a14610338578063f2204f741461034d578063fec5a2861461041b576100f3565b80638aee8c79146102cf5780638da5cb5b146102e457806394af716b146102f9578063aca0f3721461030e576100f3565b80632782e87e116100c65780632782e87e146101ad5780633a768463146101d757806360675a87146102895780636be00229146102ba576100f3565b806305b050de146100f857806308dc89d7146101025780631865c57d1461014757806322c091bc14610180575b600080fd5b61010061049f565b005b34801561010e57600080fd5b506101356004803603602081101561012557600080fd5b50356001600160a01b03166104b6565b60408051918252519081900360200190f35b34801561015357600080fd5b5061015c6104d5565b6040518082600381111561016c57fe5b60ff16815260200191505060405180910390f35b34801561018c57600080fd5b50610100600480360360808110156101a357600080fd5b50604081016104e5565b3480156101b957600080fd5b50610100600480360360208110156101d057600080fd5b50356105f7565b3480156101e357600080fd5b506101ec610879565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561025c57fe5b60ff1681526001600160a01b039092166020830152506040805191829003019a5098505050505050505050f35b34801561029557600080fd5b5061029e6108ea565b604080516001600160a01b039092168252519081900360200190f35b3480156102c657600080fd5b5061029e6108f9565b3480156102db57600080fd5b5061029e610908565b3480156102f057600080fd5b5061029e610917565b34801561030557600080fd5b50610100610926565b34801561031a57600080fd5b506101356109b9565b34801561032f57600080fd5b506101006109c8565b34801561034457600080fd5b5061029e610a28565b34801561035957600080fd5b50610100600480360360a081101561037057600080fd5b81359160208101359163ffffffff60408301351691908101906080810160608201356401000000008111156103a457600080fd5b8201836020820111156103b657600080fd5b803590602001918460018302840111640100000000831117156103d857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610a37915050565b34801561042757600080fd5b50610100600480360361010081101561043f57600080fd5b604080518082018252833593602081013593838201359360608301359360808401359363ffffffff60a08201351693810192909161010083019160c08401906002908390839080828437600092019190915250919450610b979350505050565b336000908152600860205260409020805434019055565b6001600160a01b0381166000908152600860205260409020545b919050565b600754600160401b900460ff1690565b600754600160481b90046001600160a01b031633146105355760405162461bcd60e51b815260040180806020018281038252602d8152602001806114c6602d913960400191505060405180910390fd5b60078054600160481b600160e81b031916905561059a6001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002054610cf890919063ffffffff16565b82356001600160a01b031660009081526008602081815260408320939093556105d2928401356001600160801b03169185600161055d565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b031611156106505760405162461bcd60e51b815260040180806020018281038252602781526020018061149f6027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b1580156106ca57600080fd5b505af41580156106de573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b03168082529181019190915260075484516328d23fe960e01b81529290951697506328d23fe996509194919363ffffffff16928892600490920191829187918190849084905b8381101561077657818101518382015260200161075e565b5050505090500184600260200280838360005b838110156107a1578181015183820152602001610789565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050602060405180830381600087803b1580156107e057600080fd5b505af11580156107f4573d6000803e3d6000fd5b505050506040513d602081101561080a57600080fd5b505160078054600160481b600160e81b031916600160481b6001600160a01b03938416810291909117918290556040805191909204909216825233602083015280517f65a25beed90da238c6d5cff94ce6d71c6d0b2ff27fdc5cd0ce1efba8074ed99f9281900390910190a150565b6002546003546004546005546006546007546001600160a01b03928316926001600160801b0383169267ffffffffffffffff600160801b8204811693600160c01b909204169163ffffffff8083169264010000000081049091169160ff600160401b83041691600160481b9004168b565b600a546001600160a01b031681565b6009546001600160a01b031681565b6000546001600160a01b031681565b600b546001600160a01b031681565b600b546001600160a01b0316331461097e576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff16600381111561099857fe5b14156109b7576007805468ff00000000000000001916600160401b1790555b565b6006546001600160801b031690565b600b546001600160a01b03163314610a20576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6109b7610d59565b6001546001600160a01b031681565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63ce9d5122600287878787876040518763ffffffff1660e01b8152600401808781526020018681526020018581526020018463ffffffff1663ffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610ace578181015183820152602001610ab6565b50505050905090810190601f168015610afb5780820380516001836020036101000a031916815260200191505b5097505050505050505060006040518083038186803b158015610b1d57600080fd5b505af4158015610b31573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610b6c935091506001600160801b031663ffffffff610cf816565b6005546001600160a01b0316600090815260086020526040902055610b9082610d67565b5050505050565b336000908152600860205260409020546006546001600160801b03161115610bf05760405162461bcd60e51b81526004018080602001828103825260318152602001806114f36031913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151636c36f28f60e11b8152600260048201818152602483018c9052604483018b9052606483018a90526084830189905260a4830188905263ffffffff871660c484015273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__9463d86de51e9492938d938d938d938d938d938d938d93909260e401918491908190849084905b83811015610cae578181015183820152602001610c96565b505050509050019850505050505050505060006040518083038186803b158015610cd757600080fd5b505af4158015610ceb573d6000803e3d6000fd5b5050505050505050505050565b600082820183811015610d52576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600b546001600160a01b0316ff5b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b158015610dad57600080fd5b505af1158015610dc1573d6000803e3d6000fd5b505050506040513d6020811015610dd757600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b158015610e2857600080fd5b505af4158015610e3c573d6000803e3d6000fd5b505050506040513d6020811015610e5257600080fd5b50518114610e9d57610e996040518060600160405280610e726001610f65565b8152602001610e846002800154610fe3565b8152602001610e9284610fe3565b9052611061565b6004555b60015460405163e4eb8c6360e01b81526020600482018181528551602484015285516001600160a01b039094169363e4eb8c6393879383926044909201919085019080838360005b83811015610efd578181015183820152602001610ee5565b50505050905090810190601f168015610f2a5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610f4957600080fd5b505af1158015610f5d573d6000803e3d6000fd5b505050505050565b610f6d611437565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610fd2565b610fbf611437565b815260200190600190039081610fb75790505b508152600060209091015292915050565b610feb611437565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611050565b61103d611437565b8152602001906001900390816110355790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b611085611437565b81526020019060019003908161107d575050805190915060005b818110156110d7578481600381106110b357fe5b60200201518382815181106110c457fe5b602090810291909101015260010161109f565b506110e1826110e9565b949350505050565b6000600882511115611139576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611166578160200160208202803883390190505b50805190915060005b818110156111c25761117f611465565b61119b86838151811061118e57fe5b6020026020010151611235565b905080600001518483815181106111ae57fe5b60209081029190910101525060010161116f565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561120b5781810151838201526020016111f3565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b61123d611465565b6060820151600c60ff9091161061128f576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166112bc5760405180602001604052806112b3846000015161136b565b905290506104d0565b606082015160ff16600114156113035760405180602001604052806112b384602001516000015185602001516040015186602001516060015187602001516020015161138f565b606082015160ff166002141561132857506040805160208101909152815181526104d0565b600360ff16826060015160ff161015801561134c57506060820151600c60ff909116105b156113695760405180602001604052806112b384604001516110e9565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b600083156113e9575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206110e1565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b604051806080016040528060008152602001611451611477565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f7765644f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274a265627a7a7231582007c6388d1687268196b33f5d08d891b5c2d7231f6d67bece174dcd9a5e3a08f764736f6c634300050d0032a265627a7a723158205ca4e0ff54de195bd4f377c6ad4c89dc011374ff70b22aed9c2b9722cebefaa464736f6c634300050d0032"

// DeployChainLauncher deploys a new Ethereum contract, binding an instance of ChainLauncher to it.
func DeployChainLauncher(auth *bind.TransactOpts, backend bind.ContractBackend, _globalInboxAddress common.Address, _challengeLauncherAddress common.Address) (common.Address, *types.Transaction, *ChainLauncher, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainLauncherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ChainLauncherBin = strings.Replace(ChainLauncherBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ChainLauncherBin = strings.Replace(ChainLauncherBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChainLauncherBin), backend, _globalInboxAddress, _challengeLauncherAddress)
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

// ChallengeABI is the input ABI used to generate the binding from.
const ChallengeABI = "[]"

// ChallengeBin is the compiled bytecode used for deploying new contracts.
var ChallengeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820ca2ff07d848317e18aa9377faa0dba77ac4bcdda8cc56590dca63cae921b850264736f6c634300050d0032"

// DeployChallenge deploys a new Ethereum contract, binding an instance of Challenge to it.
func DeployChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Challenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// Challenge is an auto generated Go binding around an Ethereum contract.
type Challenge struct {
	ChallengeCaller     // Read-only binding to the contract
	ChallengeTransactor // Write-only binding to the contract
	ChallengeFilterer   // Log filterer for contract events
}

// ChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeSession struct {
	Contract     *Challenge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeCallerSession struct {
	Contract *ChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTransactorSession struct {
	Contract     *ChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeRaw struct {
	Contract *Challenge // Generic contract binding to access the raw methods on
}

// ChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeCallerRaw struct {
	Contract *ChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTransactorRaw struct {
	Contract *ChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallenge creates a new instance of Challenge, bound to a specific deployed contract.
func NewChallenge(address common.Address, backend bind.ContractBackend) (*Challenge, error) {
	contract, err := bindChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// NewChallengeCaller creates a new read-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeCaller, error) {
	contract, err := bindChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeCaller{contract: contract}, nil
}

// NewChallengeTransactor creates a new write-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTransactor, error) {
	contract, err := bindChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTransactor{contract: contract}, nil
}

// NewChallengeFilterer creates a new log filterer instance of Challenge, bound to a specific deployed contract.
func NewChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFilterer, error) {
	contract, err := bindChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFilterer{contract: contract}, nil
}

// bindChallenge binds a generic wrapper to an already deployed contract.
func bindChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.ChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transact(opts, method, params...)
}

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"b32\",\"type\":\"bytes32\"}],\"name\":\"bytes32string\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"out\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// DebugPrintFuncSigs maps the 4-byte function signature to its string representation.
var DebugPrintFuncSigs = map[string]string{
	"252fb38d": "bytes32string(bytes32)",
}

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x610202610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063252fb38d1461003a575b600080fd5b6100576004803603602081101561005057600080fd5b50356100cc565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610091578181015183820152602001610079565b50505050905090810190601f1680156100be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051818152606081810183529182919060208201818038833901905050905060005b602081101561019457600084826020811061010757fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b61012d8261019b565b85856002028151811061013c57fe5b60200101906001600160f81b031916908160001a90535061015c8161019b565b85856002026001018151811061016e57fe5b60200101906001600160f81b031916908160001a90535050600190920191506100f09050565b5092915050565b6000600a60f883901c10156101bb578160f81c60300160f81b90506101c8565b8160f81c60570160f81b90505b91905056fea265627a7a7231582066b66e0573d8b50e4f34a432fb74922dc886db7d72dedc7491b12fc19e449f1664736f6c634300050d0032"

// DeployDebugPrint deploys a new Ethereum contract, binding an instance of DebugPrint to it.
func DeployDebugPrint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebugPrint, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DebugPrintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// DebugPrint is an auto generated Go binding around an Ethereum contract.
type DebugPrint struct {
	DebugPrintCaller     // Read-only binding to the contract
	DebugPrintTransactor // Write-only binding to the contract
	DebugPrintFilterer   // Log filterer for contract events
}

// DebugPrintCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebugPrintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebugPrintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebugPrintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebugPrintSession struct {
	Contract     *DebugPrint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebugPrintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebugPrintCallerSession struct {
	Contract *DebugPrintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DebugPrintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebugPrintTransactorSession struct {
	Contract     *DebugPrintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DebugPrintRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebugPrintRaw struct {
	Contract *DebugPrint // Generic contract binding to access the raw methods on
}

// DebugPrintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebugPrintCallerRaw struct {
	Contract *DebugPrintCaller // Generic read-only contract binding to access the raw methods on
}

// DebugPrintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebugPrintTransactorRaw struct {
	Contract *DebugPrintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebugPrint creates a new instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrint(address common.Address, backend bind.ContractBackend) (*DebugPrint, error) {
	contract, err := bindDebugPrint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// NewDebugPrintCaller creates a new read-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintCaller(address common.Address, caller bind.ContractCaller) (*DebugPrintCaller, error) {
	contract, err := bindDebugPrint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintCaller{contract: contract}, nil
}

// NewDebugPrintTransactor creates a new write-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintTransactor(address common.Address, transactor bind.ContractTransactor) (*DebugPrintTransactor, error) {
	contract, err := bindDebugPrint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintTransactor{contract: contract}, nil
}

// NewDebugPrintFilterer creates a new log filterer instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintFilterer(address common.Address, filterer bind.ContractFilterer) (*DebugPrintFilterer, error) {
	contract, err := bindDebugPrint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebugPrintFilterer{contract: contract}, nil
}

// bindDebugPrint binds a generic wrapper to an already deployed contract.
func bindDebugPrint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.DebugPrintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transact(opts, method, params...)
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintCaller) Bytes32string(opts *bind.CallOpts, b32 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DebugPrint.contract.Call(opts, out, "bytes32string", b32)
	return *ret0, err
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintSession) Bytes32string(b32 [32]byte) (string, error) {
	return _DebugPrint.Contract.Bytes32string(&_DebugPrint.CallOpts, b32)
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintCallerSession) Bytes32string(b32 [32]byte) (string, error) {
	return _DebugPrint.Contract.Bytes32string(&_DebugPrint.CallOpts, b32)
}

// DisputableABI is the input ABI used to generate the binding from.
const DisputableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"withinTimeBounds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DisputableFuncSigs maps the 4-byte function signature to its string representation.
var DisputableFuncSigs = map[string]string{
	"ce9d5122": "confirmDisputableAsserted(VM.Data storage,bytes32,bytes32,uint32,bytes,bytes32)",
	"5af93c7a": "initiateChallenge(VM.Data storage,bytes32)",
	"d86de51e": "pendingDisputableAssert(VM.Data storage,bytes32,bytes32,bytes32,bytes32,bytes32,uint32,uint64[2])",
	"42c0787e": "withinTimeBounds(uint64[2])",
}

// DisputableBin is the compiled bytecode used for deploying new contracts.
var DisputableBin = "0x6110b4610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806342c0787e1461005b5780635af93c7a146100ba578063ce9d5122146100ec578063d86de51e146101bf575b600080fd5b6100a66004803603604081101561007157600080fd5b604080518082018252918301929181830191839060029083908390808284376000920191909152509194506102489350505050565b604080519115158252519081900360200190f35b8180156100c657600080fd5b506100ea600480360360408110156100dd57600080fd5b508035906020013561027a565b005b8180156100f857600080fd5b506100ea600480360360c081101561010f57600080fd5b81359160208101359160408201359163ffffffff6060820135169181019060a08101608082013564010000000081111561014857600080fd5b82018360208201111561015a57600080fd5b8035906020019184600183028401116401000000008311171561017c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610434915050565b8180156101cb57600080fd5b506100ea60048036036101208110156101e357600080fd5b604080518082018252833593602081013593838201359360608301359360808401359360a08101359363ffffffff60c083013516939082019261012083019160e084019060029083908390808284376000920191909152509194506107ed9350505050565b805160009067ffffffffffffffff1643108015906102745750602082015167ffffffffffffffff164311155b92915050565b60038201546001600160a01b03163314156102c65760405162461bcd60e51b8152600401808060200182810382526021815260200180610eeb6021913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561031557600080fd5b505af4158015610329573d6000803e3d6000fd5b505050506040513d602081101561033f57600080fd5b505161037c5760405162461bcd60e51b8152600401808060200182810382526026815260200180610f936026913960400191505060405180910390fd5b60026005830154600160401b900460ff16600381111561039857fe5b146103d45760405162461bcd60e51b815260040180806020018281038252602f815260200180610e24602f913960400191505060405180910390fd5b816001015481146104165760405162461bcd60e51b815260040180806020018281038252604d815260200180610fe6604d913960600191505060405180910390fd5b5060006001820155600501805460ff60401b1916600160401b179055565b60026005870154600160401b900460ff16600381111561045057fe5b1461048c5760405162461bcd60e51b8152600401808060200182810382526022815260200180610ea56022913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5876040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156104db57600080fd5b505af41580156104ef573d6000803e3d6000fd5b505050506040513d602081101561050557600080fd5b5051156105435760405162461bcd60e51b8152600401808060200182810382526024815260200180610e816024913960400191505060405180910390fd5b85600101548573__$9836fa7140e5a33041d4b827682e675a30$__637ddf59d68787600073__$9836fa7140e5a33041d4b827682e675a30$__63e83f4bfe8a6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156105ce5781810151838201526020016105b6565b50505050905090810190601f1680156105fb5780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b15801561061857600080fd5b505af415801561062c573d6000803e3d6000fd5b505050506040513d602081101561064257600080fd5b5051604080516001600160e01b031960e088901b168152600481019590955263ffffffff939093166024850152604484019190915260648301526000608483015260a482018790525160c4808301926020929190829003018186803b1580156106aa57600080fd5b505af41580156106be573d6000803e3d6000fd5b505050506040513d60208110156106d457600080fd5b505160408051602081810194909452808201929092528051808303820181526060909201905280519101201461073b5760405162461bcd60e51b815260040180806020018281038252604d815260200180611033604d913960600191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63eb49982c87866040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561079257600080fd5b505af41580156107a6573d6000803e3d6000fd5b5050604080518781526020810185905281517f4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb9450908190039091019150a1505050505050565b60016005890154600160401b900460ff16600381111561080957fe5b146108455760405162461bcd60e51b815260040180806020018281038252602d815260200180610fb9602d913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__632a3e0a97896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561089457600080fd5b505af41580156108a8573d6000803e3d6000fd5b505050506040513d60208110156108be57600080fd5b5051158015610945575073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561091757600080fd5b505af415801561092b573d6000803e3d6000fd5b505050506040513d602081101561094157600080fd5b5051155b6109805760405162461bcd60e51b815260040180806020018281038252603e815260200180610f0c603e913960400191505060405180910390fd5b6005880154690100000000000000000090046001600160a01b0316156109d75760405162461bcd60e51b815260040180806020018281038252602e815260200180610e53602e913960400191505060405180910390fd5b600588015463ffffffff64010000000090910481169083161115610a42576040805162461bcd60e51b815260206004820152601f60248201527f547269656420746f206578656375746520746f6f206d616e7920737465707300604482015290519081900360640190fd5b610a4b81610248565b610a865760405162461bcd60e51b8152600401808060200182810382526024815260200180610ec76024913960400191505060405180910390fd5b87548714610ac55760405162461bcd60e51b8152600401808060200182810382526027815260200180610f6c6027913960400191505060405180910390fd5b87600201548614610b075760405162461bcd60e51b8152600401808060200182810382526022815260200180610f4a6022913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63a3a162cb896040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610b5657600080fd5b505af4158015610b6a573d6000803e3d6000fd5b5050505073__$9836fa7140e5a33041d4b827682e675a30$__6385ecb92a8883896040518463ffffffff1660e01b81526004018084815260200183600260200280838360005b83811015610bc8578181015183820152602001610bb0565b50505050905001828152602001935050505060206040518083038186803b158015610bf257600080fd5b505af4158015610c06573d6000803e3d6000fd5b505050506040513d6020811015610c1c57600080fd5b505160408051633eefaceb60e11b81526004810188905263ffffffff8516602482015260006044820181905260648201889052608482015260a48101869052905173__$9836fa7140e5a33041d4b827682e675a30$__91637ddf59d69160c4808301926020929190829003018186803b158015610c9857600080fd5b505af4158015610cac573d6000803e3d6000fd5b505050506040513d6020811015610cc257600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019052805191012060018901556003880180546001600160a01b031916331790556005880180546002919060ff60401b1916600160401b8302179055507f247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe244366040518060a00160405280898152602001888152602001878152602001868152602001858152503383858c60040160109054906101000a900467ffffffffffffffff166040518086600560200280838360005b83811015610dae578181015183820152602001610d96565b505050506001600160a01b03881692019182525060200184604080838360005b83811015610de6578181015183820152602001610dce565b50505063ffffffff90961692019182525067ffffffffffffffff909216602083015250604080519182900301945092505050a1505050505050505056fe417373657274696f6e206d7573742062652070656e64696e6720746f20696e697469617465206368616c6c656e676543616e206f6e6c792064697370757461626c6520617373657274206966206e6f7420696e206368616c6c656e6765417373657274696f6e206973207374696c6c2070656e64696e67206368616c6c656e6765564d20646f6573206e6f74206861766520617373657274696f6e2070656e64696e67507265636f6e646974696f6e3a206e6f742077697468696e2074696d6520626f756e64734368616c6c656e676520776173206372656174656420627920617373657274657243616e206f6e6c792064697370757461626c6520617373657274206966206d616368696e65206973206e6f74206572726f726564206f722068616c746564507265636f6e646974696f6e3a20696e626f7820646f6573206e6f74206d61746368507265636f6e646974696f6e3a207374617465206861736820646f6573206e6f74206d617463684368616c6c656e676520646964206e6f7420636f6d65206265666f726520646561646c696e6543616e206f6e6c792064697370757461626c65206173736572742066726f6d2077616974696e67207374617465496e697469617465204368616c6c656e67653a20507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6e436f6e6669726d2044697370757461626c653a20507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6ea265627a7a7231582012d0ea63c930431ee772e1f2ac8c1575ad090c63bf44fa3e162b11c9e2eab46d64736f6c634300050d0032"

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
	Fields     [5][32]byte
	Asserter   common.Address
	TimeBounds [2]uint64
	NumSteps   uint32
	Deadline   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
func (_Disputable *DisputableFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*DisputablePendingDisputableAssertionIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &DisputablePendingDisputableAssertionIterator{contract: _Disputable.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
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

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
func (_Disputable *DisputableFilterer) ParsePendingDisputableAssertion(log types.Log) (*DisputablePendingDisputableAssertion, error) {
	event := new(DisputablePendingDisputableAssertion)
	if err := _Disputable.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IChallengeLauncherABI is the input ABI used to generate the binding from.
const IChallengeLauncherABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"launchChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeLauncherFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeLauncherFuncSigs = map[string]string{
	"28d23fe9": "launchChallenge(address[2],uint128[2],uint32,bytes32)",
}

// IChallengeLauncher is an auto generated Go binding around an Ethereum contract.
type IChallengeLauncher struct {
	IChallengeLauncherCaller     // Read-only binding to the contract
	IChallengeLauncherTransactor // Write-only binding to the contract
	IChallengeLauncherFilterer   // Log filterer for contract events
}

// IChallengeLauncherCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeLauncherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeLauncherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeLauncherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeLauncherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeLauncherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeLauncherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeLauncherSession struct {
	Contract     *IChallengeLauncher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IChallengeLauncherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeLauncherCallerSession struct {
	Contract *IChallengeLauncherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IChallengeLauncherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeLauncherTransactorSession struct {
	Contract     *IChallengeLauncherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IChallengeLauncherRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeLauncherRaw struct {
	Contract *IChallengeLauncher // Generic contract binding to access the raw methods on
}

// IChallengeLauncherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeLauncherCallerRaw struct {
	Contract *IChallengeLauncherCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeLauncherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeLauncherTransactorRaw struct {
	Contract *IChallengeLauncherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeLauncher creates a new instance of IChallengeLauncher, bound to a specific deployed contract.
func NewIChallengeLauncher(address common.Address, backend bind.ContractBackend) (*IChallengeLauncher, error) {
	contract, err := bindIChallengeLauncher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeLauncher{IChallengeLauncherCaller: IChallengeLauncherCaller{contract: contract}, IChallengeLauncherTransactor: IChallengeLauncherTransactor{contract: contract}, IChallengeLauncherFilterer: IChallengeLauncherFilterer{contract: contract}}, nil
}

// NewIChallengeLauncherCaller creates a new read-only instance of IChallengeLauncher, bound to a specific deployed contract.
func NewIChallengeLauncherCaller(address common.Address, caller bind.ContractCaller) (*IChallengeLauncherCaller, error) {
	contract, err := bindIChallengeLauncher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeLauncherCaller{contract: contract}, nil
}

// NewIChallengeLauncherTransactor creates a new write-only instance of IChallengeLauncher, bound to a specific deployed contract.
func NewIChallengeLauncherTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeLauncherTransactor, error) {
	contract, err := bindIChallengeLauncher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeLauncherTransactor{contract: contract}, nil
}

// NewIChallengeLauncherFilterer creates a new log filterer instance of IChallengeLauncher, bound to a specific deployed contract.
func NewIChallengeLauncherFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeLauncherFilterer, error) {
	contract, err := bindIChallengeLauncher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeLauncherFilterer{contract: contract}, nil
}

// bindIChallengeLauncher binds a generic wrapper to an already deployed contract.
func bindIChallengeLauncher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeLauncherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeLauncher *IChallengeLauncherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeLauncher.Contract.IChallengeLauncherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeLauncher *IChallengeLauncherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeLauncher.Contract.IChallengeLauncherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeLauncher *IChallengeLauncherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeLauncher.Contract.IChallengeLauncherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeLauncher *IChallengeLauncherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeLauncher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeLauncher *IChallengeLauncherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeLauncher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeLauncher *IChallengeLauncherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeLauncher.Contract.contract.Transact(opts, method, params...)
}

// LaunchChallenge is a paid mutator transaction binding the contract method 0x28d23fe9.
//
// Solidity: function launchChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns(address)
func (_IChallengeLauncher *IChallengeLauncherTransactor) LaunchChallenge(opts *bind.TransactOpts, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeLauncher.contract.Transact(opts, "launchChallenge", _players, _escrows, _challengePeriod, _challengeRoot)
}

// LaunchChallenge is a paid mutator transaction binding the contract method 0x28d23fe9.
//
// Solidity: function launchChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns(address)
func (_IChallengeLauncher *IChallengeLauncherSession) LaunchChallenge(_players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeLauncher.Contract.LaunchChallenge(&_IChallengeLauncher.TransactOpts, _players, _escrows, _challengePeriod, _challengeRoot)
}

// LaunchChallenge is a paid mutator transaction binding the contract method 0x28d23fe9.
//
// Solidity: function launchChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns(address)
func (_IChallengeLauncher *IChallengeLauncherTransactorSession) LaunchChallenge(_players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeLauncher.Contract.LaunchChallenge(&_IChallengeLauncher.TransactOpts, _players, _escrows, _challengePeriod, _challengeRoot)
}

// IGlobalPendingInboxABI is the input ABI used to generate the binding from.
const IGlobalPendingInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes21\",\"name\":\"tokenType\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"hasFunds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pullPendingMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var IGlobalPendingInboxFuncSigs = map[string]string{
	"3bbc3c32": "forwardMessage(address,bytes21,uint256,bytes,bytes)",
	"acb633b6": "hasFunds(address,bytes21[],uint256[])",
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

// IVMTrackerABI is the input ABI used to generate the binding from.
const IVMTrackerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IVMTrackerFuncSigs maps the 4-byte function signature to its string representation.
var IVMTrackerFuncSigs = map[string]string{
	"22c091bc": "completeChallenge(address[2],uint128[2])",
}

// IVMTracker is an auto generated Go binding around an Ethereum contract.
type IVMTracker struct {
	IVMTrackerCaller     // Read-only binding to the contract
	IVMTrackerTransactor // Write-only binding to the contract
	IVMTrackerFilterer   // Log filterer for contract events
}

// IVMTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVMTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVMTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVMTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVMTrackerSession struct {
	Contract     *IVMTracker       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVMTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVMTrackerCallerSession struct {
	Contract *IVMTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IVMTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVMTrackerTransactorSession struct {
	Contract     *IVMTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IVMTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVMTrackerRaw struct {
	Contract *IVMTracker // Generic contract binding to access the raw methods on
}

// IVMTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVMTrackerCallerRaw struct {
	Contract *IVMTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// IVMTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVMTrackerTransactorRaw struct {
	Contract *IVMTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVMTracker creates a new instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTracker(address common.Address, backend bind.ContractBackend) (*IVMTracker, error) {
	contract, err := bindIVMTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVMTracker{IVMTrackerCaller: IVMTrackerCaller{contract: contract}, IVMTrackerTransactor: IVMTrackerTransactor{contract: contract}, IVMTrackerFilterer: IVMTrackerFilterer{contract: contract}}, nil
}

// NewIVMTrackerCaller creates a new read-only instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerCaller(address common.Address, caller bind.ContractCaller) (*IVMTrackerCaller, error) {
	contract, err := bindIVMTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerCaller{contract: contract}, nil
}

// NewIVMTrackerTransactor creates a new write-only instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVMTrackerTransactor, error) {
	contract, err := bindIVMTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerTransactor{contract: contract}, nil
}

// NewIVMTrackerFilterer creates a new log filterer instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVMTrackerFilterer, error) {
	contract, err := bindIVMTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerFilterer{contract: contract}, nil
}

// bindIVMTracker binds a generic wrapper to an already deployed contract.
func bindIVMTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVMTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVMTracker *IVMTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVMTracker.Contract.IVMTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVMTracker *IVMTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVMTracker.Contract.IVMTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVMTracker *IVMTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVMTracker.Contract.IVMTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVMTracker *IVMTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVMTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVMTracker *IVMTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVMTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVMTracker *IVMTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVMTracker.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _players, _rewards)
}

// MerkleLibABI is the input ABI used to generate the binding from.
const MerkleLibABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"generateAddressRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"generateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleLibFuncSigs maps the 4-byte function signature to its string representation.
var MerkleLibFuncSigs = map[string]string{
	"6a2dda67": "generateAddressRoot(address[])",
	"9898dc10": "generateRoot(bytes32[])",
	"b792d767": "verifyProof(bytes,bytes32,bytes32,uint256)",
}

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
var MerkleLibBin = "0x610575610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043f565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b6000815b600181511115610422576060600282516001018161031f57fe5b04604051908082528060200260200182016040528015610349578160200160208202803883390190505b50905060005b815181101561041a5782518160020260010110156103e25782816002028151811061037657fe5b602002602001015183826002026001018151811061039057fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d157fe5b602002602001018181525050610412565b8281600202815181106103f157fe5b602002602001015182828151811061040557fe5b6020026020010181815250505b60010161034f565b509050610305565b8060008151811061042f57fe5b6020026020010151915050919050565b600080838160205b88518111610532578089015193506020818a51036020018161046557fe5b0491505b60008211801561047c5750600286066001145b801561048a57508160020a86115b1561049d57600286046001019550610469565b600286066104e85783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104e057fe5b04955061052a565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052357fe5b0460010195505b602001610447565b50509094149594505050505056fea265627a7a72315820dab403e0cf8b5b3cea673ec093051f92ba4d6ae087244559622ae30d50f21ec664736f6c634300050d0032"

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

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[7]\",\"name\":\"fields\",\"type\":\"bytes32[7]\"},{\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"a49c3308": "oneStepProof(Challenge.Data storage,bytes32[2],uint64[2],bytes32[5],bytes)",
	"c0fee45d": "validateProof(bytes32[7],uint64[2],bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x613a7b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063a49c330814610045578063c0fee45d14610177575b600080fd5b610175600480360361016081101561005c57600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152505060408051808201825292959493818101939250906002908390839080828437600092019190915250506040805160a0818101909252929594938181019392509060059083908390808284376000920191909152509194939260208101925035905064010000000081111561010057600080fd5b82018360208201111561011257600080fd5b8035906020019184600183028401116401000000008311171561013457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061028c945050505050565b005b61027a600480360361014081101561018e57600080fd5b810190808060e00190600780602002604051908101604052809291908260076020028082843760009201919091525050604080518082018252929594938181019392509060029083908390808284376000920191909152509194939260208101925035905064010000000081111561020557600080fd5b82018360208201111561021757600080fd5b8035906020019184600183028401116401000000008311171561023957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610627945050505050565b60408051918252519081900360200190f35b60016005860154600160601b900460ff1660028111156102a857fe5b146102e45760405162461bcd60e51b8152600401808060200182810382526039815260200180613a0e6039913960400191505060405180910390fd5b600585015467ffffffffffffffff16431115610347576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6520737465702070726f6f66206d697373656420646561646c696e650000604482015290519081900360640190fd5b600185015484516020860151604080516342f65c9560e11b81526004810184815273__$9836fa7140e5a33041d4b827682e675a30$__946385ecb92a9490938a9391929160240190849080838360005b838110156103af578181015183820152602001610397565b50505050905001828152602001935050505060206040518083038186803b1580156103d957600080fd5b505af41580156103ed573d6000803e3d6000fd5b505050506040513d602081101561040357600080fd5b50518351602080860151604080880151606089015160808a01518351633eefaceb60e11b815260048101979097526001602488015260448701949094526064860191909152608485015260a48401919091525173__$9836fa7140e5a33041d4b827682e675a30$__92637ddf59d69260c4808301939192829003018186803b15801561048e57600080fd5b505af41580156104a2573d6000803e3d6000fd5b505050506040513d60208110156104b857600080fd5b505160408051602081810194909452808201929092528051808303820181526060909201905280519101201461051f5760405162461bcd60e51b81526004018080602001828103825260268152602001806138e96026913960400191505060405180910390fd5b60006105d46040518060e001604052808760006002811061053c57fe5b602002015181526020018760016002811061055357fe5b602002015181526020018560006005811061056a57fe5b602002015181526020018560016005811061058157fe5b602002015181526020018560026005811061059857fe5b60200201518152602001856003600581106105af57fe5b60200201518152602001856004600581106105c657fe5b602002015190528584610627565b9050801561061f576040805162461bcd60e51b8152602060048201526013602482015272141c9bdbd9881dd85cc81a5b98dbdc9c9958dd606a1b604482015290519081900360640190fd5b505050505050565b60006106e86040518061012001604052808660006007811061064557fe5b602002015181526020018581526020018660016007811061066257fe5b602002015181526020018660026007811061067957fe5b602002015181526020018660036007811061069057fe5b60200201518152602001866004600781106106a757fe5b60200201518152602001866005600781106106be57fe5b60200201518152602001866006600781106106d557fe5b60200201518152602001848152506106f2565b90505b9392505050565b60008080806060610701613820565b610709613820565b610712886115f6565b939950929650909450925090506001600060ff881682141561076857610761838660008151811061073f57fe5b60200260200101518760018151811061075457fe5b6020026020010151611a36565b915061144a565b60ff8816600214156107a757610761838660008151811061078557fe5b60200260200101518760018151811061079a57fe5b6020026020010151611a84565b60ff8816600314156107e65761076183866000815181106107c457fe5b6020026020010151876001815181106107d957fe5b6020026020010151611ac5565b60ff88166004141561082557610761838660008151811061080357fe5b60200260200101518760018151811061081857fe5b6020026020010151611b06565b60ff88166005141561086457610761838660008151811061084257fe5b60200260200101518760018151811061085757fe5b6020026020010151611b57565b60ff8816600614156108a357610761838660008151811061088157fe5b60200260200101518760018151811061089657fe5b6020026020010151611ba8565b60ff8816600714156108e25761076183866000815181106108c057fe5b6020026020010151876001815181106108d557fe5b6020026020010151611bf9565b60ff8816600814156109365761076183866000815181106108ff57fe5b60200260200101518760018151811061091457fe5b60200260200101518860028151811061092957fe5b6020026020010151611c4a565b60ff88166009141561098a57610761838660008151811061095357fe5b60200260200101518760018151811061096857fe5b60200260200101518860028151811061097d57fe5b6020026020010151611cb4565b60ff8816600a14156109c95761076183866000815181106109a757fe5b6020026020010151876001815181106109bc57fe5b6020026020010151611d0d565b60ff881660101415610a085761076183866000815181106109e657fe5b6020026020010151876001815181106109fb57fe5b6020026020010151611d4e565b60ff881660111415610a47576107618386600081518110610a2557fe5b602002602001015187600181518110610a3a57fe5b6020026020010151611d8f565b60ff881660121415610a86576107618386600081518110610a6457fe5b602002602001015187600181518110610a7957fe5b6020026020010151611dd0565b60ff881660131415610ac5576107618386600081518110610aa357fe5b602002602001015187600181518110610ab857fe5b6020026020010151611e11565b60ff881660141415610b04576107618386600081518110610ae257fe5b602002602001015187600181518110610af757fe5b6020026020010151611e52565b60ff881660151415610b2e576107618386600081518110610b2157fe5b6020026020010151611e7e565b60ff881660161415610b6d576107618386600081518110610b4b57fe5b602002602001015187600181518110610b6057fe5b6020026020010151611ec4565b60ff881660171415610bac576107618386600081518110610b8a57fe5b602002602001015187600181518110610b9f57fe5b6020026020010151611f05565b60ff881660181415610beb576107618386600081518110610bc957fe5b602002602001015187600181518110610bde57fe5b6020026020010151611f46565b60ff881660191415610c15576107618386600081518110610c0857fe5b6020026020010151611f87565b60ff8816601a1415610c54576107618386600081518110610c3257fe5b602002602001015187600181518110610c4757fe5b6020026020010151611fbd565b60ff8816601b1415610c93576107618386600081518110610c7157fe5b602002602001015187600181518110610c8657fe5b6020026020010151611ffe565b60ff881660201415610cbd576107618386600081518110610cb057fe5b602002602001015161203f565b60ff881660211415610ce7576107618386600081518110610cda57fe5b602002602001015161205b565b60ff881660301415610d11576107618386600081518110610d0457fe5b6020026020010151612076565b60ff881660311415610d26576107618361207e565b60ff881660321415610d3b576107618361209f565b60ff881660331415610d65576107618386600081518110610d5857fe5b60200260200101516120b8565b60ff881660341415610d8f576107618386600081518110610d8257fe5b60200260200101516120d1565b60ff881660351415610dce576107618386600081518110610dac57fe5b602002602001015187600181518110610dc157fe5b60200260200101516120e7565b60ff881660361415610de3576107618361212f565b60ff881660371415610dfd57610761838560000151612161565b60ff881660381415610e27576107618386600081518110610e1a57fe5b6020026020010151612173565b60ff881660391415610eb457610e3b613881565b610e4a8b610100015188612185565b919950975090508715610e8e5760405162461bcd60e51b81526004018080602001828103825260218152602001806139ed6021913960400191505060405180910390fd5b610e9e858263ffffffff61230f16565b610eae848263ffffffff61233116565b5061144a565b60ff8816603a1415610ec9576107618361234e565b60ff8816603b1415610eda5761144a565b60ff8816603c1415610eef576107618361236e565b60ff8816603d1415610f19576107618386600081518110610f0c57fe5b6020026020010151612387565b60ff881660401415610f43576107618386600081518110610f3657fe5b60200260200101516123b5565b60ff881660411415610f82576107618386600081518110610f6057fe5b602002602001015187600181518110610f7557fe5b60200260200101516123d7565b60ff881660421415610fd6576107618386600081518110610f9f57fe5b602002602001015187600181518110610fb457fe5b602002602001015188600281518110610fc957fe5b6020026020010151612409565b60ff881660431415611015576107618386600081518110610ff357fe5b60200260200101518760018151811061100857fe5b602002602001015161244b565b60ff88166044141561106957610761838660008151811061103257fe5b60200260200101518760018151811061104757fe5b60200260200101518860028151811061105c57fe5b602002602001015161245d565b60ff8816605014156110a857610761838660008151811061108657fe5b60200260200101518760018151811061109b57fe5b602002602001015161247f565b60ff8816605114156110fc5761076183866000815181106110c557fe5b6020026020010151876001815181106110da57fe5b6020026020010151886002815181106110ef57fe5b60200260200101516124f5565b60ff88166052141561112657610761838660008151811061111957fe5b602002602001015161256d565b60ff88166060141561113b57610761836125a0565b60ff88166061141561123857611165838660008151811061115857fe5b60200260200101516125a6565b9092509050811561122f578960e001518a60c0015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146111e45760405162461bcd60e51b81526004018080602001828103825260258152602001806139a16025913960400191505060405180910390fd5b8960a001518a608001511461122a5760405162461bcd60e51b81526004018080602001828103825260278152602001806139c66027913960400191505060405180910390fd5b611233565b5060005b61144a565b60ff88166070141561132757611262838660008151811061125557fe5b60200260200101516125ca565b9092509050811561122f578960a001518a6080015182604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146112e15760405162461bcd60e51b81526004018080602001828103825260298152602001806139316029913960400191505060405180910390fd5b8960e001518a60c001511461122a5760405162461bcd60e51b815260040180806020018281038252602681526020018061395a6026913960400191505060405180910390fd5b60ff8816607114156113e3576040805160028082526060828101909352816020015b611351613881565b81526020019060019003908161134957505060208c01519091506113869060005b602002015167ffffffffffffffff166125e4565b8160008151811061139357fe5b60200260200101819052506113b28b6020015160016002811061137257fe5b816001815181106113bf57fe5b6020026020010181905250610eae6113d682612662565b859063ffffffff61233116565b60ff88166072141561142057610761838660008151811061140057fe5b602002602001015160405180602001604052808e60400151815250612712565b60ff881660731415611435576000915061144a565b60ff88166074141561144a5761144a83612784565b806114db578960a001518a60800151146114955760405162461bcd60e51b81526004018080602001828103825260278152602001806139c66027913960400191505060405180910390fd5b8960e001518a60c00151146114db5760405162461bcd60e51b815260040180806020018281038252602681526020018061395a6026913960400191505060405180910390fd5b8161153d5760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a0840151511415611535576115308361278e565b61153d565b60a083015183525b61154684612798565b8a51146115845760405162461bcd60e51b815260040180806020018281038252602281526020018061390f6022913960400191505060405180910390fd5b61158d83612798565b8a60600151146115e4576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b60006060611602613820565b61160a613820565b60008080611616613820565b61161f8161282d565b61162e89610100015184612837565b909450909250905061163e613820565b6116478261293c565b905060008a6101000151858151811061165c57fe5b602001015160f81c60f81b60f81c905060008b6101000151866001018151811061168257fe5b016020015160f81c905060006116978261299a565b90506060816040519080825280602002602001820160405280156116d557816020015b6116c2613881565b8152602001906001900390816116ba5790505b5090506002880197508360ff16600014806116f357508360ff166001145b611744576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff84166117e7576040805160208082018084528951516353409fab60e01b90915260ff87166024840152604483015291519091829173__$d969135829891f807aa9c34494da4ecd99$__916353409fab916064808601929190818703018186803b1580156117b257600080fd5b505af41580156117c6573d6000803e3d6000fd5b505050506040513d60208110156117dc57600080fd5b50519052865261193e565b6117ef613881565b6117fe8f61010001518a612185565b909a5090985090508715611859576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561187d57808260008151811061186d57fe5b602002602001018190525061188d565b61188d868263ffffffff61233116565b604051806020016040528073__$d969135829891f807aa9c34494da4ecd99$__63264f384b876118bc866129b4565b518c5151604080516001600160e01b031960e087901b16815260ff909416600485015260248401929092526044830152516064808301926020929190829003018186803b15801561190c57600080fd5b505af4158015611920573d6000803e3d6000fd5b505050506040513d602081101561193657600080fd5b505190528752505b60ff84165b828110156119d25761195a8f61010001518a612185565b845185908590811061196857fe5b60209081029190910101529950975087156119ca576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611943565b815115611a1f575060005b8460ff16825103811015611a1f57611a17828260018551030381518110611a0057fe5b60200260200101518861233190919063ffffffff16565b6001016119dd565b50919d919c50939a50919850939650945050505050565b6000611a4183612aea565b1580611a535750611a5182612aea565b155b15611a60575060006106eb565b82518251808201611a77878263ffffffff612af516565b5060019695505050505050565b6000611a8f83612aea565b1580611aa15750611a9f82612aea565b155b15611aae575060006106eb565b82518251808202611a77878263ffffffff612af516565b6000611ad083612aea565b1580611ae25750611ae082612aea565b155b15611aef575060006106eb565b82518251808203611a77878263ffffffff612af516565b6000611b1183612aea565b1580611b235750611b2182612aea565b155b15611b30575060006106eb565b8251825180611b44576000925050506106eb565b808204611a77878263ffffffff612af516565b6000611b6283612aea565b1580611b745750611b7282612aea565b155b15611b81575060006106eb565b8251825180611b95576000925050506106eb565b808205611a77878263ffffffff612af516565b6000611bb383612aea565b1580611bc55750611bc382612aea565b155b15611bd2575060006106eb565b8251825180611be6576000925050506106eb565b808206611a77878263ffffffff612af516565b6000611c0483612aea565b1580611c165750611c1482612aea565b155b15611c23575060006106eb565b8251825180611c37576000925050506106eb565b808207611a77878263ffffffff612af516565b6000611c5584612aea565b1580611c675750611c6583612aea565b155b15611c7457506000611cac565b83518351835180611c8b5760009350505050611cac565b6000818385089050611ca3898263ffffffff612af516565b60019450505050505b949350505050565b6000611cbf84612aea565b1580611cd15750611ccf83612aea565b155b15611cde57506000611cac565b83518351835180611cf55760009350505050611cac565b6000818385099050611ca3898263ffffffff612af516565b6000611d1883612aea565b1580611d2a5750611d2882612aea565b155b15611d37575060006106eb565b8251825180820a611a77878263ffffffff612af516565b6000611d5983612aea565b1580611d6b5750611d6982612aea565b155b15611d78575060006106eb565b82518251808210611a77878263ffffffff612af516565b6000611d9a83612aea565b1580611dac5750611daa82612aea565b155b15611db9575060006106eb565b82518251808211611a77878263ffffffff612af516565b6000611ddb83612aea565b1580611ded5750611deb82612aea565b155b15611dfa575060006106eb565b82518251808212611a77878263ffffffff612af516565b6000611e1c83612aea565b1580611e2e5750611e2c82612aea565b155b15611e3b575060006106eb565b82518251808213611a77878263ffffffff612af516565b6000611e746113d6611e63846129b4565b51611e6d866129b4565b5114612b09565b5060019392505050565b6000611e8982612aea565b611ea357611e9e83600063ffffffff612af516565b611eba565b81518015611eb7858263ffffffff612af516565b50505b5060015b92915050565b6000611ecf83612aea565b1580611ee15750611edf82612aea565b155b15611eee575060006106eb565b82518251808216611a77878263ffffffff612af516565b6000611f1083612aea565b1580611f225750611f2082612aea565b155b15611f2f575060006106eb565b82518251808217611a77878263ffffffff612af516565b6000611f5183612aea565b1580611f635750611f6182612aea565b155b15611f70575060006106eb565b82518251808218611a77878263ffffffff612af516565b6000611f9282612aea565b611f9e57506000611ebe565b81518019611fb2858263ffffffff612af516565b506001949350505050565b6000611fc883612aea565b1580611fda5750611fd882612aea565b155b15611fe7575060006106eb565b8251825181811a611a77878263ffffffff612af516565b600061200983612aea565b158061201b575061201982612aea565b155b15612028575060006106eb565b8251825181810b611a77878263ffffffff612af516565b6000611eba61204d836129b4565b51849063ffffffff612af516565b6000611eba61206983612b32565b849063ffffffff61233116565b600192915050565b6000612097826080015183612bbb90919063ffffffff16565b506001919050565b6000612097826060015183612bbb90919063ffffffff16565b60006120c3826129b4565b606084015250600192915050565b60006120dc826129b4565b835250600192915050565b60006120f283612bc9565b6120fe575060006106eb565b61210782612aea565b612113575060006106eb565b815115611e7457612123836129b4565b84525060019392505050565b6000612097612154612147612142612bd6565b6129b4565b5160208501515114612b09565b839063ffffffff61233116565b6000611eba838363ffffffff612bbb16565b6000611eba838363ffffffff61230f16565b600080612190613881565b845184106121e5576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b600084905060008682815181106121f857fe5b016020015160019092019160f81c905060006122126138af565b60ff8316612246576122248985612c53565b9094509150600084612235846125e4565b919850965094506123089350505050565b60ff83166001141561226d5761225c8985612c7a565b909450905060008461223583612de7565b60ff831660021415612294576122838985612c53565b909450915060008461223584612e47565b600360ff8416108015906122ab5750600c60ff8416105b156122e8576002198301606060006122c4838d89612ec5565b9098509250905080876122d684612662565b99509950995050505050505050612308565b8260ff166127100160006122fc60006125e4565b91985096509450505050505b9250925092565b6123258260400151612320836129b4565b612f80565b82604001819052505050565b6123428260200151612320836129b4565b82602001819052505050565b6000612097612154612361612142612bd6565b5160408501515114612b09565b60006120978260a0015183612bbb90919063ffffffff16565b600061239282612bc9565b61239e57506000611ebe565b6123a7826129b4565b60a084015250600192915050565b60006123c7838363ffffffff61233116565b611eba838363ffffffff61233116565b60006123e9848363ffffffff61233116565b6123f9848463ffffffff61233116565b611e74848363ffffffff61233116565b600061241b858363ffffffff61233116565b61242b858463ffffffff61233116565b61243b858563ffffffff61233116565b611fb2858363ffffffff61233116565b60006123f9848463ffffffff61233116565b600061246f858563ffffffff61233116565b61243b858463ffffffff61233116565b600061248a83612aea565b158061249c575061249a82613036565b155b156124a9575060006106eb565b6124b282613045565b60ff168360000151106124c7575060006106eb565b611e7482604001518460000151815181106124de57fe5b60200260200101518561233190919063ffffffff16565b600061250083613036565b1580612512575061251084612aea565b155b1561251f57506000611cac565b61252883613045565b60ff1684600001511061253d57506000611cac565b81836040015185600001518151811061255257fe5b6020908102919091010152611fb2858463ffffffff61233116565b600061257882613036565b61258457506000611ebe565b611eba61259083613045565b849060ff1663ffffffff612af516565b50600190565b6000806125b16138d6565b6125ba846129b4565b51600193509150505b9250929050565b60008060016125d8846129b4565b51909590945092505050565b6125ec613881565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612651565b61263e613881565b8152602001906001900390816126365790505b508152600060209091015292915050565b61266a613881565b6126748251613054565b6126c5576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b8051600090612720846129b4565b511415612774576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b611e74848363ffffffff612bbb16565b600260c090910152565b600160c090910152565b600060028260c0015114156127af575060006115f1565b60018260c0015114156127c4575060016115f1565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101206115f1565b600060c090910152565b600080612842613820565b61284a613820565b600060c0820181905261285d878761305b565b84529650905080156128755793508492509050612308565b61287f878761305b565b602085015296509050801561289a5793508492509050612308565b6128a4878761305b565b60408501529650905080156128bf5793508492509050612308565b6128c9878761305b565b60608501529650905080156128e45793508492509050612308565b6128ee878761305b565b60808501529650905080156129095793508492509050612308565b612913878761305b565b60a085015296509050801561292e5793508492509050612308565b506000969495509392505050565b612944613820565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b60008060006129ab8460ff16613099565b50949350505050565b6129bc6138d6565b6060820151600c60ff90911610612a0e576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16612a3b576040518060200160405280612a32846000015161352c565b905290506115f1565b606082015160ff1660011415612a82576040518060200160405280612a32846020015160000151856020015160400151866020015160600151876020015160200151613550565b606082015160ff1660021415612aa757506040805160208101909152815181526115f1565b600360ff16826060015160ff1610158015612acb57506060820151600c60ff909116105b15612ae8576040518060200160405280612a3284604001516135f8565bfe5b6060015160ff161590565b6123428260200151612320612142846125e4565b612b11613881565b8115612b2857612b2160016125e4565b90506115f1565b612b2160006125e4565b612b3a613881565b816060015160ff1660021415612b815760405162461bcd60e51b81526004018080602001828103825260218152602001806139806021913960400191505060405180910390fd5b606082015160ff16612b9757612b2160006125e4565b816060015160ff1660011415612bb157612b2160016125e4565b612b2160036125e4565b612342826020015182612f80565b6060015160ff1660011490565b612bde613881565b604080516080808201835260008083528351918201845280825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612c43565b612c30613881565b815260200190600190039081612c285790505b5081526003602090910152905090565b6000808281612c68868363ffffffff61374416565b60209290920196919550909350505050565b6000612c846138af565b60008390506000858281518110612c9757fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110612cbd57fe5b016020015160019384019360f89190911c915060009060ff84161415612d5b576000612ce7613881565b612cf18a87612185565b90975090925090508115612d4c576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b612d55816129b4565b51925050505b6000612d6d898663ffffffff61374416565b90506020850194508360ff1660011415612db2576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506125c39050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b612def613881565b604080516080810182526000808252602080830186905283518281529081018452919283019190612e36565b612e23613881565b815260200190600190039081612e1b5790505b508152600160209091015292915050565b612e4f613881565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612eb4565b612ea1613881565b815260200190600190039081612e995790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015612f1057816020015b612efd613881565b815260200190600190039081612ef55790505b50905060005b8960ff168160ff161015612f6a57612f2e8985612185565b8451859060ff8616908110612f3f57fe5b6020908102919091010152945092508215612f6257509094509092509050612f77565b600101612f16565b5060009550919350909150505b93509350939050565b612f886138d6565b6040805160028082526060828101909352816020015b612fa66138d6565b815260200190600190039081612f9e5790505090508281600081518110612fc957fe5b60200260200101819052508381600181518110612fe257fe5b6020026020010181905250604051806020016040528061302c60405180604001604052806130138860000151612e47565b81526020016130258960000151612e47565b9052613760565b9052949350505050565b6000611ebe82606001516137df565b6000611ebe82606001516137fd565b6008101590565b6000806130666138d6565b836000613079878363ffffffff61374416565b604080516020808201909252918252600099930197509550909350505050565b60008060018314156130b15750600290506001613527565b60028314156130c65750600290506001613527565b60038314156130db5750600290506001613527565b60048314156130f05750600290506001613527565b60058314156131055750600290506001613527565b600683141561311a5750600290506001613527565b600783141561312f5750600290506001613527565b60088314156131445750600390506001613527565b60098314156131595750600390506001613527565b600a83141561316e5750600290506001613527565b60108314156131835750600290506001613527565b60118314156131985750600290506001613527565b60128314156131ad5750600290506001613527565b60138314156131c25750600290506001613527565b60148314156131d75750600290506001613527565b60158314156131eb57506001905080613527565b60168314156132005750600290506001613527565b60178314156132155750600290506001613527565b601883141561322a5750600290506001613527565b601983141561323e57506001905080613527565b601a8314156132535750600290506001613527565b601b8314156132685750600290506001613527565b602083141561327c57506001905080613527565b602183141561329057506001905080613527565b60308314156132a55750600190506000613527565b60318314156132ba5750600090506001613527565b60328314156132cf5750600090506001613527565b60338314156132e45750600190506000613527565b60348314156132f95750600190506000613527565b603583141561330e5750600290506000613527565b60368314156133235750600090506001613527565b60378314156133385750600090506001613527565b603883141561334d5750600190506000613527565b60398314156133625750600090506001613527565b603a8314156133775750600090506001613527565b603b83141561338b57506000905080613527565b603c8314156133a05750600090506001613527565b603d8314156133b55750600190506000613527565b60408314156133ca5750600190506002613527565b60418314156133df5750600290506003613527565b60428314156133f45750600390506004613527565b604383141561340857506002905080613527565b604483141561341c57506003905080613527565b60508314156134315750600290506001613527565b60518314156134465750600390506001613527565b605283141561345a57506001905080613527565b606083141561346e57506000905080613527565b60618314156134835750600190506000613527565b60708314156134985750600190506000613527565b60718314156134ad5750600090506001613527565b60728314156134c157506001905080613527565b60738314156134d557506000905080613527565b60748314156134e957506000905080613527565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b600083156135aa575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611cac565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6000600882511115613648576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613675578160200160208202803883390190505b50805190915060005b818110156136d15761368e6138d6565b6136aa86838151811061369d57fe5b60200260200101516129b4565b905080600001518483815181106136bd57fe5b60209081029190910101525060010161367e565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561371a578181015183820152602001613702565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6000816020018351101561375757600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b613783613881565b81526020019060019003908161377b575050805190915060005b818110156137d5578481600281106137b157fe5b60200201518382815181106137c257fe5b602090810291909101015260010161379d565b50611cac826135f8565b6000600c60ff8316108015611ebe575050600360ff91909116101590565b6000613808826137df565b15613818575060021981016115f1565b5060016115f1565b6040518060e001604052806138336138d6565b81526020016138406138d6565b815260200161384d6138d6565b815260200161385a6138d6565b81526020016138676138d6565b81526020016138746138d6565b8152602001600081525090565b60405180608001604052806000815260200161389b6138af565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4f6e6520737465702070726f6f66207769746820696e76616c6964207072657620737461746550726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f64654c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c756543616e206f6e6c79206f6e6520737465702070726f6f6620666f6c6c6f77696e6720612073696e676c652073746570206368616c6c656e6765a265627a7a723158205454f99bb5f3f12b0999e28a1e80e969c813b68ea193a7a88263713ed7626da764736f6c634300050d0032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// OneStepProof is an auto generated Go binding around an Ethereum contract.
type OneStepProof struct {
	OneStepProofCaller     // Read-only binding to the contract
	OneStepProofTransactor // Write-only binding to the contract
	OneStepProofFilterer   // Log filterer for contract events
}

// OneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofSession struct {
	Contract     *OneStepProof     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofCallerSession struct {
	Contract *OneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofTransactorSession struct {
	Contract     *OneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofRaw struct {
	Contract *OneStepProof // Generic contract binding to access the raw methods on
}

// OneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofCallerRaw struct {
	Contract *OneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofTransactorRaw struct {
	Contract *OneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof creates a new instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProof(address common.Address, backend bind.ContractBackend) (*OneStepProof, error) {
	contract, err := bindOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// NewOneStepProofCaller creates a new read-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofCaller, error) {
	contract, err := bindOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofCaller{contract: contract}, nil
}

// NewOneStepProofTransactor creates a new write-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofTransactor, error) {
	contract, err := bindOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTransactor{contract: contract}, nil
}

// NewOneStepProofFilterer creates a new log filterer instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofFilterer, error) {
	contract, err := bindOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofFilterer{contract: contract}, nil
}

// bindOneStepProof binds a generic wrapper to an already deployed contract.
func bindOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.OneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transact(opts, method, params...)
}

// ValidateProof is a free data retrieval call binding the contract method 0xc0fee45d.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, fields [7][32]byte, timeBounds [2]uint64, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", fields, timeBounds, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0xc0fee45d.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0xc0fee45d.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, proof)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d6e8e06479213c9152f9553b84c33b25d6e34e76d5ebd7701bb5d4993e7e97fb64736f6c634300050d0032"

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
var VMBin = "0x6101ea610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80632a3e0a97146100665780638ab48be514610097578063a3a162cb146100b4578063e2fe93ca146100e0578063eb49982c146100fd575b600080fd5b6100836004803603602081101561007c57600080fd5b503561012d565b604080519115158252519081900360200190f35b610083600480360360208110156100ad57600080fd5b5035610134565b8180156100c057600080fd5b506100de600480360360208110156100d757600080fd5b503561014f565b005b610083600480360360208110156100f657600080fd5b503561018e565b81801561010957600080fd5b506100de6004803603604081101561012057600080fd5b5080359060200135610193565b5460011490565b60040154600160801b900467ffffffffffffffff1643111590565b60058101546004909101805467ffffffffffffffff60801b1916600160801b63ffffffff909316430167ffffffffffffffff1692909202919091179055565b541590565b8155600501805468ff000000000000000019166801000000000000000017905556fea265627a7a72315820e69b7cf8a45b77664496fae79d6ba2c7caf8edd822b5925029ed5b054f1c2d3c64736f6c634300050d0032"

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
