// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package channellauncher

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

// ArbChannelABI is the input ABI used to generate the binding from.
const ArbChannelABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_validatorKeys\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"ConfirmedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unanHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PendingAssertionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unanHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingUnanimousAssertion\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activatedValidators\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeManager\",\"outputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"confirmUnanimousAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"finalizedUnanimousAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumVM.State\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isListedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"isValidatorList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_unanRest\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"pendingUnanimousAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"enumVM.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"inChallenge\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbChannelFuncSigs maps the 4-byte function signature to its string representation.
var ArbChannelFuncSigs = map[string]string{
	"94af716b": "activateVM()",
	"899b7c74": "activatedValidators()",
	"023a96fe": "challengeManager()",
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"f2204f74": "confirmDisputableAsserted(bytes32,bytes32,uint32,bytes,bytes32)",
	"e1e5d090": "confirmUnanimousAsserted(bytes32,bytes32,bytes)",
	"08dc89d7": "currentDeposit(address)",
	"aca0f372": "escrowRequired()",
	"6be00229": "exitAddress()",
	"812fa865": "finalizedUnanimousAssert(bytes32,bytes32,bytes,bytes32,bytes)",
	"1865c57d": "getState()",
	"d489113a": "globalInbox()",
	"05b050de": "increaseDeposit()",
	"2782e87e": "initiateChallenge(bytes32)",
	"b99738e0": "isListedValidator(address)",
	"513164fe": "isValidatorList(address[])",
	"8da5cb5b": "owner()",
	"cfa80707": "ownerShutdown()",
	"fec5a286": "pendingDisputableAssert(bytes32,bytes32,bytes32,bytes32,bytes32,uint32,uint64[2])",
	"df949904": "pendingUnanimousAssert(bytes32,uint64,bytes32,bytes32,bytes)",
	"60675a87": "terminateAddress()",
	"0f43a677": "validatorCount()",
	"3a768463": "vm()",
}

// ArbChannelBin is the compiled bytecode used for deploying new contracts.
var ArbChannelBin = "0x60806040523480156200001157600080fd5b50604051620023783803806200237883398181016040526101008110156200003857600080fd5b815160208301516040808501516060860151608087015160a088015160c089015160e08a0180519651989a9799959894979396929591949391820192846401000000008211156200008857600080fd5b9083019060208201858111156200009e57600080fd5b8251866020820283011164010000000082111715620000bc57600080fd5b82525081516020918201928201910280838360005b83811015620000eb578181015183820152602001620000d1565b505050509050016040525050508787878787878780600160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000806101000a8154816001600160a01b0302191690836001600160a01b03160217905550600160009054906101000a90046001600160a01b03166001600160a01b031663f39723836040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156200019d57600080fd5b505af1158015620001b2573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b1580156200023a57600080fd5b505af41580156200024f573d6000803e3d6000fd5b505050506040513d60208110156200026657600080fd5b50516004555050600680546001600160801b039093166001600160801b031990931692909217909155506007805463ffffffff9283166401000000000263ffffffff60201b199390941663ffffffff19918216179290921692909217909155600d8054845161ffff1692169190911790555060005b600d5461ffff908116908216101562000340576001600c6000848461ffff16815181106200030557fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055600101620002db565b50505050505050505061201f80620003596000396000f3fe6080604052600436106101405760003560e01c8063812fa865116100b6578063cfa807071161006f578063cfa80707146105b5578063d489113a146105ca578063df949904146105df578063e1e5d090146106b2578063f2204f741461076f578063fec5a2861461083b57610140565b8063812fa865146103e4578063899b7c741461052e5780638da5cb5b1461054357806394af716b14610558578063aca0f3721461056d578063b99738e01461058257610140565b806322c091bc1161010857806322c091bc1461022a5780632782e87e146102575780633a76846314610281578063513164fe1461032b57806360675a87146103ba5780636be00229146103cf57610140565b8063023a96fe1461014557806305b050de1461017657806308dc89d7146101805780630f43a677146101c55780631865c57d146101f1575b600080fd5b34801561015157600080fd5b5061015a6108bf565b604080516001600160a01b039092168252519081900360200190f35b61017e6108ce565b005b34801561018c57600080fd5b506101b3600480360360208110156101a357600080fd5b50356001600160a01b03166109eb565b60408051918252519081900360200190f35b3480156101d157600080fd5b506101da610a0a565b6040805161ffff9092168252519081900360200190f35b3480156101fd57600080fd5b50610206610a14565b6040518082600381111561021657fe5b60ff16815260200191505060405180910390f35b34801561023657600080fd5b5061017e6004803603608081101561024d57600080fd5b5060408101610a24565b34801561026357600080fd5b5061017e6004803603602081101561027a57600080fd5b5035610b77565b34801561028d57600080fd5b50610296610d7b565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561030657fe5b60ff1681529115156020830152506040805191829003019a5098505050505050505050f35b34801561033757600080fd5b506103a66004803603602081101561034e57600080fd5b810190602081018135600160201b81111561036857600080fd5b82018360208201111561037a57600080fd5b803590602001918460208302840111600160201b8311171561039b57600080fd5b509092509050610dec565b604080519115158252519081900360200190f35b3480156103c657600080fd5b5061015a610e73565b3480156103db57600080fd5b5061015a610e82565b3480156103f057600080fd5b5061017e600480360360a081101561040757600080fd5b813591602081013591810190606081016040820135600160201b81111561042d57600080fd5b82018360208201111561043f57600080fd5b803590602001918460018302840111600160201b8311171561046057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b8111156104ba57600080fd5b8201836020820111156104cc57600080fd5b803590602001918460018302840111600160201b831117156104ed57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e91945050505050565b34801561053a57600080fd5b506101da6110ee565b34801561054f57600080fd5b5061015a6110fe565b34801561056457600080fd5b5061017e61110d565b34801561057957600080fd5b506101b361119b565b34801561058e57600080fd5b506103a6600480360360208110156105a557600080fd5b50356001600160a01b03166111aa565b3480156105c157600080fd5b5061017e6111c8565b3480156105d657600080fd5b5061015a611228565b3480156105eb57600080fd5b5061017e600480360360a081101561060257600080fd5b81359167ffffffffffffffff6020820135169160408201359160608101359181019060a081016080820135600160201b81111561063e57600080fd5b82018360208201111561065057600080fd5b803590602001918460018302840111600160201b8311171561067157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611237945050505050565b3480156106be57600080fd5b5061017e600480360360608110156106d557600080fd5b813591602081013591810190606081016040820135600160201b8111156106fb57600080fd5b82018360208201111561070d57600080fd5b803590602001918460018302840111600160201b8311171561072e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113ef945050505050565b34801561077b57600080fd5b5061017e600480360360a081101561079257600080fd5b81359160208101359163ffffffff6040830135169190810190608081016060820135600160201b8111156107c557600080fd5b8201836020820111156107d757600080fd5b803590602001918460018302840111600160201b831117156107f857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506114df915050565b34801561084757600080fd5b5061017e600480360361010081101561085f57600080fd5b604080518082018252833593602081013593838201359360608301359360808401359363ffffffff60a08201351693810192909161010083019160c084019060029083908390808284376000920191909152509194506116389350505050565b6000546001600160a01b031681565b336000908152600c602052604090205460ff16610932576040805162461bcd60e51b815260206004820152601860248201527f43616c6c6572206d7573742062652076616c696461746f720000000000000000604482015290519081900360640190fd5b3360009081526008602052604090208054600654348201928390556001600160801b03161180801561096f57506006546001600160801b03168210155b1561099957600d8054600161ffff62010000808404821692909201160263ffff0000199091161790555b600d5462010000810461ffff90811691161480156109ce57506000600754600160401b900460ff1660038111156109cc57fe5b145b156109e7576007805460ff60401b1916600160401b1790555b5050565b6001600160a01b0381166000908152600860205260409020545b919050565b600d5461ffff1681565b600754600160401b900460ff1690565b6000546001600160a01b03163314610a6d5760405162461bcd60e51b815260040180806020018281038252602d815260200180611f8d602d913960400191505060405180910390fd5b600754600160481b900460ff16610ab55760405162461bcd60e51b8152600401808060200182810382526026815260200180611f676026913960400191505060405180910390fd5b6007805469ff00000000000000000019169055610b1a6001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000205461179990919063ffffffff16565b82356001600160a01b03166000908152600860208181526040832093909355610b52928401356001600160801b031691856001610add565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b03161115610bd05760405162461bcd60e51b8152600401808060200182810382526027815260200180611f406027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b158015610c4a57600080fd5b505af4158015610c5e573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b0316808252918101919091526007548451630823813560e21b815292909516975063208e04d496509194919363ffffffff16928892600490920191829187918190849084905b83811015610cf6578181015183820152602001610cde565b5050505090500184600260200280838360005b83811015610d21578181015183820152602001610d09565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050600060405180830381600087803b158015610d6057600080fd5b505af1158015610d74573d6000803e3d6000fd5b5050505050565b6002546003546004546005546006546007546001600160a01b03909216916001600160801b0382169167ffffffffffffffff600160801b8204811692600160c01b909204169063ffffffff80821691600160201b81049091169060ff600160401b8204811691600160481b9004168b565b600d54600090829061ffff168114610e08576000915050610e6d565b60005b600d5461ffff16811015610e6657600c6000868684818110610e2957fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff16610e5e57600092505050610e6d565b600101610e0b565b5060019150505b92915050565b600a546001600160a01b031681565b6009546001600160a01b031681565b73__$caf066876633ea418098495f1e5bb4c2f5$__635ee899da60023088888888886040518863ffffffff1660e01b815260040180888152602001876001600160a01b03166001600160a01b031681526020018681526020018581526020018060200184815260200180602001838103835286818151815260200191508051906020019080838360005b83811015610f33578181015183820152602001610f1b565b50505050905090810190601f168015610f605780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610f93578181015183820152602001610f7b565b50505050905090810190601f168015610fc05780820380516001836020036101000a031916815260200191505b50995050505050505050505060006040518083038186803b158015610fe457600080fd5b505af4158015610ff8573d6000803e3d6000fd5b5060029250611005915050565b600754600160401b900460ff16600381111561101d57fe5b1415611075576006546005546001600160a01b0316600090815260086020526040902054611059916001600160801b031663ffffffff61179916565b6005546001600160a01b03166000908152600860205260409020555b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63eb49982c6002876040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156110cd57600080fd5b505af41580156110e1573d6000803e3d6000fd5b50505050610d74836117fa565b600d5462010000900461ffff1681565b600b546001600160a01b031681565b600b546001600160a01b03163314611165576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff16600381111561117f57fe5b1415611199576007805460ff60401b1916600160401b1790555b565b6006546001600160801b031690565b6001600160a01b03166000908152600c602052604090205460ff1690565b600b546001600160a01b03163314611220576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6111996119f8565b6001546001600160a01b031681565b73__$caf066876633ea418098495f1e5bb4c2f5$__63b4d866a260023088888888886040518863ffffffff1660e01b815260040180888152602001876001600160a01b03166001600160a01b031681526020018681526020018567ffffffffffffffff1667ffffffffffffffff16815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156112ef5781810151838201526020016112d7565b50505050905090810190601f16801561131c5780820380516001836020036101000a031916815260200191505b509850505050505050505060006040518083038186803b15801561133f57600080fd5b505af4158015611353573d6000803e3d6000fd5b5060029250611360915050565b600754600160401b900460ff16600381111561137857fe5b14156113d0576006546005546001600160a01b03166000908152600860205260409020546113b4916001600160801b031663ffffffff61179916565b6005546001600160a01b03166000908152600860205260409020555b50506007805460ff60401b191668030000000000000000179055505050565b73__$caf066876633ea418098495f1e5bb4c2f5$__63e2d5c52f60028585856040518563ffffffff1660e01b81526004018085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561146c578181015183820152602001611454565b50505050905090810190601f1680156114995780820380516001836020036101000a031916815260200191505b509550505050505060006040518083038186803b1580156114b957600080fd5b505af41580156114cd573d6000803e3d6000fd5b505050506114da816117fa565b505050565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63ce9d5122600287878787876040518763ffffffff1660e01b8152600401808781526020018681526020018581526020018463ffffffff1663ffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561157657818101518382015260200161155e565b50505050905090810190601f1680156115a35780820380516001836020036101000a031916815260200191505b5097505050505050505060006040518083038186803b1580156115c557600080fd5b505af41580156115d9573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054611614935091506001600160801b031663ffffffff61179916565b6005546001600160a01b0316600090815260086020526040902055610d74826117fa565b336000908152600860205260409020546006546001600160801b031611156116915760405162461bcd60e51b8152600401808060200182810382526031815260200180611fba6031913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151636c36f28f60e11b8152600260048201818152602483018c9052604483018b9052606483018a90526084830189905260a4830188905263ffffffff871660c484015273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__9463d86de51e9492938d938d938d938d938d938d938d93909260e401918491908190849084905b8381101561174f578181015183820152602001611737565b505050509050019850505050505050505060006040518083038186803b15801561177857600080fd5b505af415801561178c573d6000803e3d6000fd5b5050505050505050505050565b6000828201838110156117f3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b15801561184057600080fd5b505af1158015611854573d6000803e3d6000fd5b505050506040513d602081101561186a57600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b1580156118bb57600080fd5b505af41580156118cf573d6000803e3d6000fd5b505050506040513d60208110156118e557600080fd5b505181146119305761192c60405180606001604052806119056001611a06565b81526020016119176002800154611a84565b815260200161192584611a84565b9052611b02565b6004555b60015460405163e4eb8c6360e01b81526020600482018181528551602484015285516001600160a01b039094169363e4eb8c6393879383926044909201919085019080838360005b83811015611990578181015183820152602001611978565b50505050905090810190601f1680156119bd5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156119dc57600080fd5b505af11580156119f0573d6000803e3d6000fd5b505050505050565b600b546001600160a01b0316ff5b611a0e611ed8565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611a73565b611a60611ed8565b815260200190600190039081611a585790505b508152600060209091015292915050565b611a8c611ed8565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611af1565b611ade611ed8565b815260200190600190039081611ad65790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b611b26611ed8565b815260200190600190039081611b1e575050805190915060005b81811015611b7857848160038110611b5457fe5b6020020151838281518110611b6557fe5b6020908102919091010152600101611b40565b50611b8282611b8a565b949350505050565b6000600882511115611bda576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611c07578160200160208202803883390190505b50805190915060005b81811015611c6357611c20611f06565b611c3c868381518110611c2f57fe5b6020026020010151611cd6565b90508060000151848381518110611c4f57fe5b602090810291909101015250600101611c10565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611cac578181015183820152602001611c94565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b611cde611f06565b6060820151600c60ff90911610611d30576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16611d5d576040518060200160405280611d548460000151611e0c565b90529050610a05565b606082015160ff1660011415611da4576040518060200160405280611d54846020015160000151856020015160400151866020015160600151876020015160200151611e30565b606082015160ff1660021415611dc95750604080516020810190915281518152610a05565b600360ff16826060015160ff1610158015611ded57506060820151600c60ff909116105b15611e0a576040518060200160405280611d548460400151611b8a565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b60008315611e8a575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611b82565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b604051806080016040528060008152602001611ef2611f18565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274a265627a7a72315820a3ecd7ee59e443a702a2aeed38ed6415902bc44c3b697088416369159973f5fc64736f6c634300050c0032"

// DeployArbChannel deploys a new Ethereum contract, binding an instance of ArbChannel to it.
func DeployArbChannel(auth *bind.TransactOpts, backend bind.ContractBackend, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeManagerAddress common.Address, _globalInboxAddress common.Address, _validatorKeys []common.Address) (common.Address, *types.Transaction, *ArbChannel, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ArbChannelBin = strings.Replace(ArbChannelBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	vMAddr, _, _, _ := DeployVM(auth, backend)
	ArbChannelBin = strings.Replace(ArbChannelBin, "__$8e266570c8a7fb2aaac83b3e040afaf9e1$__", vMAddr.String()[2:], -1)

	unanimousAddr, _, _, _ := DeployUnanimous(auth, backend)
	ArbChannelBin = strings.Replace(ArbChannelBin, "__$caf066876633ea418098495f1e5bb4c2f5$__", unanimousAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbChannelBin = strings.Replace(ArbChannelBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbChannelBin), backend, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _challengeManagerAddress, _globalInboxAddress, _validatorKeys)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbChannel{ArbChannelCaller: ArbChannelCaller{contract: contract}, ArbChannelTransactor: ArbChannelTransactor{contract: contract}, ArbChannelFilterer: ArbChannelFilterer{contract: contract}}, nil
}

// ArbChannel is an auto generated Go binding around an Ethereum contract.
type ArbChannel struct {
	ArbChannelCaller     // Read-only binding to the contract
	ArbChannelTransactor // Write-only binding to the contract
	ArbChannelFilterer   // Log filterer for contract events
}

// ArbChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbChannelSession struct {
	Contract     *ArbChannel       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbChannelCallerSession struct {
	Contract *ArbChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbChannelTransactorSession struct {
	Contract     *ArbChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbChannelRaw struct {
	Contract *ArbChannel // Generic contract binding to access the raw methods on
}

// ArbChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbChannelCallerRaw struct {
	Contract *ArbChannelCaller // Generic read-only contract binding to access the raw methods on
}

// ArbChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbChannelTransactorRaw struct {
	Contract *ArbChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbChannel creates a new instance of ArbChannel, bound to a specific deployed contract.
func NewArbChannel(address common.Address, backend bind.ContractBackend) (*ArbChannel, error) {
	contract, err := bindArbChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbChannel{ArbChannelCaller: ArbChannelCaller{contract: contract}, ArbChannelTransactor: ArbChannelTransactor{contract: contract}, ArbChannelFilterer: ArbChannelFilterer{contract: contract}}, nil
}

// NewArbChannelCaller creates a new read-only instance of ArbChannel, bound to a specific deployed contract.
func NewArbChannelCaller(address common.Address, caller bind.ContractCaller) (*ArbChannelCaller, error) {
	contract, err := bindArbChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbChannelCaller{contract: contract}, nil
}

// NewArbChannelTransactor creates a new write-only instance of ArbChannel, bound to a specific deployed contract.
func NewArbChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbChannelTransactor, error) {
	contract, err := bindArbChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbChannelTransactor{contract: contract}, nil
}

// NewArbChannelFilterer creates a new log filterer instance of ArbChannel, bound to a specific deployed contract.
func NewArbChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbChannelFilterer, error) {
	contract, err := bindArbChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbChannelFilterer{contract: contract}, nil
}

// bindArbChannel binds a generic wrapper to an already deployed contract.
func bindArbChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbChannel *ArbChannelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbChannel.Contract.ArbChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbChannel *ArbChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChannel.Contract.ArbChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbChannel *ArbChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbChannel.Contract.ArbChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbChannel *ArbChannelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbChannel *ArbChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbChannel *ArbChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbChannel.Contract.contract.Transact(opts, method, params...)
}

// ActivatedValidators is a free data retrieval call binding the contract method 0x899b7c74.
//
// Solidity: function activatedValidators() constant returns(uint16)
func (_ArbChannel *ArbChannelCaller) ActivatedValidators(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "activatedValidators")
	return *ret0, err
}

// ActivatedValidators is a free data retrieval call binding the contract method 0x899b7c74.
//
// Solidity: function activatedValidators() constant returns(uint16)
func (_ArbChannel *ArbChannelSession) ActivatedValidators() (uint16, error) {
	return _ArbChannel.Contract.ActivatedValidators(&_ArbChannel.CallOpts)
}

// ActivatedValidators is a free data retrieval call binding the contract method 0x899b7c74.
//
// Solidity: function activatedValidators() constant returns(uint16)
func (_ArbChannel *ArbChannelCallerSession) ActivatedValidators() (uint16, error) {
	return _ArbChannel.Contract.ActivatedValidators(&_ArbChannel.CallOpts)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbChannel *ArbChannelCaller) ChallengeManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "challengeManager")
	return *ret0, err
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbChannel *ArbChannelSession) ChallengeManager() (common.Address, error) {
	return _ArbChannel.Contract.ChallengeManager(&_ArbChannel.CallOpts)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_ArbChannel *ArbChannelCallerSession) ChallengeManager() (common.Address, error) {
	return _ArbChannel.Contract.ChallengeManager(&_ArbChannel.CallOpts)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbChannel *ArbChannelCaller) CurrentDeposit(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "currentDeposit", validator)
	return *ret0, err
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbChannel *ArbChannelSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _ArbChannel.Contract.CurrentDeposit(&_ArbChannel.CallOpts, validator)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_ArbChannel *ArbChannelCallerSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _ArbChannel.Contract.CurrentDeposit(&_ArbChannel.CallOpts, validator)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbChannel *ArbChannelCaller) EscrowRequired(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "escrowRequired")
	return *ret0, err
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbChannel *ArbChannelSession) EscrowRequired() (*big.Int, error) {
	return _ArbChannel.Contract.EscrowRequired(&_ArbChannel.CallOpts)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_ArbChannel *ArbChannelCallerSession) EscrowRequired() (*big.Int, error) {
	return _ArbChannel.Contract.EscrowRequired(&_ArbChannel.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbChannel *ArbChannelCaller) ExitAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "exitAddress")
	return *ret0, err
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbChannel *ArbChannelSession) ExitAddress() (common.Address, error) {
	return _ArbChannel.Contract.ExitAddress(&_ArbChannel.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_ArbChannel *ArbChannelCallerSession) ExitAddress() (common.Address, error) {
	return _ArbChannel.Contract.ExitAddress(&_ArbChannel.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbChannel *ArbChannelCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "getState")
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbChannel *ArbChannelSession) GetState() (uint8, error) {
	return _ArbChannel.Contract.GetState(&_ArbChannel.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_ArbChannel *ArbChannelCallerSession) GetState() (uint8, error) {
	return _ArbChannel.Contract.GetState(&_ArbChannel.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbChannel *ArbChannelCaller) GlobalInbox(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "globalInbox")
	return *ret0, err
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbChannel *ArbChannelSession) GlobalInbox() (common.Address, error) {
	return _ArbChannel.Contract.GlobalInbox(&_ArbChannel.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_ArbChannel *ArbChannelCallerSession) GlobalInbox() (common.Address, error) {
	return _ArbChannel.Contract.GlobalInbox(&_ArbChannel.CallOpts)
}

// IsListedValidator is a free data retrieval call binding the contract method 0xb99738e0.
//
// Solidity: function isListedValidator(address validator) constant returns(bool)
func (_ArbChannel *ArbChannelCaller) IsListedValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "isListedValidator", validator)
	return *ret0, err
}

// IsListedValidator is a free data retrieval call binding the contract method 0xb99738e0.
//
// Solidity: function isListedValidator(address validator) constant returns(bool)
func (_ArbChannel *ArbChannelSession) IsListedValidator(validator common.Address) (bool, error) {
	return _ArbChannel.Contract.IsListedValidator(&_ArbChannel.CallOpts, validator)
}

// IsListedValidator is a free data retrieval call binding the contract method 0xb99738e0.
//
// Solidity: function isListedValidator(address validator) constant returns(bool)
func (_ArbChannel *ArbChannelCallerSession) IsListedValidator(validator common.Address) (bool, error) {
	return _ArbChannel.Contract.IsListedValidator(&_ArbChannel.CallOpts, validator)
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_ArbChannel *ArbChannelCaller) IsValidatorList(opts *bind.CallOpts, _validators []common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "isValidatorList", _validators)
	return *ret0, err
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_ArbChannel *ArbChannelSession) IsValidatorList(_validators []common.Address) (bool, error) {
	return _ArbChannel.Contract.IsValidatorList(&_ArbChannel.CallOpts, _validators)
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_ArbChannel *ArbChannelCallerSession) IsValidatorList(_validators []common.Address) (bool, error) {
	return _ArbChannel.Contract.IsValidatorList(&_ArbChannel.CallOpts, _validators)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbChannel *ArbChannelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbChannel *ArbChannelSession) Owner() (common.Address, error) {
	return _ArbChannel.Contract.Owner(&_ArbChannel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ArbChannel *ArbChannelCallerSession) Owner() (common.Address, error) {
	return _ArbChannel.Contract.Owner(&_ArbChannel.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbChannel *ArbChannelCaller) TerminateAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "terminateAddress")
	return *ret0, err
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbChannel *ArbChannelSession) TerminateAddress() (common.Address, error) {
	return _ArbChannel.Contract.TerminateAddress(&_ArbChannel.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_ArbChannel *ArbChannelCallerSession) TerminateAddress() (common.Address, error) {
	return _ArbChannel.Contract.TerminateAddress(&_ArbChannel.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() constant returns(uint16)
func (_ArbChannel *ArbChannelCaller) ValidatorCount(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _ArbChannel.contract.Call(opts, out, "validatorCount")
	return *ret0, err
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() constant returns(uint16)
func (_ArbChannel *ArbChannelSession) ValidatorCount() (uint16, error) {
	return _ArbChannel.Contract.ValidatorCount(&_ArbChannel.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() constant returns(uint16)
func (_ArbChannel *ArbChannelCallerSession) ValidatorCount() (uint16, error) {
	return _ArbChannel.Contract.ValidatorCount(&_ArbChannel.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbChannel *ArbChannelCaller) Vm(opts *bind.CallOpts) (struct {
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
	err := _ArbChannel.contract.Call(opts, out, "vm")
	return *ret, err
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbChannel *ArbChannelSession) Vm() (struct {
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
	return _ArbChannel.Contract.Vm(&_ArbChannel.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint8 state, bool inChallenge)
func (_ArbChannel *ArbChannelCallerSession) Vm() (struct {
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
	return _ArbChannel.Contract.Vm(&_ArbChannel.CallOpts)
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbChannel *ArbChannelTransactor) ActivateVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "activateVM")
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbChannel *ArbChannelSession) ActivateVM() (*types.Transaction, error) {
	return _ArbChannel.Contract.ActivateVM(&_ArbChannel.TransactOpts)
}

// ActivateVM is a paid mutator transaction binding the contract method 0x94af716b.
//
// Solidity: function activateVM() returns()
func (_ArbChannel *ArbChannelTransactorSession) ActivateVM() (*types.Transaction, error) {
	return _ArbChannel.Contract.ActivateVM(&_ArbChannel.TransactOpts)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbChannel *ArbChannelTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbChannel *ArbChannelSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbChannel.Contract.CompleteChallenge(&_ArbChannel.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_ArbChannel *ArbChannelTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _ArbChannel.Contract.CompleteChallenge(&_ArbChannel.TransactOpts, _players, _rewards)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChannel *ArbChannelTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "confirmDisputableAsserted", _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChannel *ArbChannelSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.ConfirmDisputableAsserted(&_ArbChannel.TransactOpts, _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0xf2204f74.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes _messages, bytes32 _logsAccHash) returns()
func (_ArbChannel *ArbChannelTransactorSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _messages []byte, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.ConfirmDisputableAsserted(&_ArbChannel.TransactOpts, _preconditionHash, _afterHash, _numSteps, _messages, _logsAccHash)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0xe1e5d090.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _afterHash, bytes32 _newInbox, bytes _messages) returns()
func (_ArbChannel *ArbChannelTransactor) ConfirmUnanimousAsserted(opts *bind.TransactOpts, _afterHash [32]byte, _newInbox [32]byte, _messages []byte) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "confirmUnanimousAsserted", _afterHash, _newInbox, _messages)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0xe1e5d090.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _afterHash, bytes32 _newInbox, bytes _messages) returns()
func (_ArbChannel *ArbChannelSession) ConfirmUnanimousAsserted(_afterHash [32]byte, _newInbox [32]byte, _messages []byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.ConfirmUnanimousAsserted(&_ArbChannel.TransactOpts, _afterHash, _newInbox, _messages)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0xe1e5d090.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _afterHash, bytes32 _newInbox, bytes _messages) returns()
func (_ArbChannel *ArbChannelTransactorSession) ConfirmUnanimousAsserted(_afterHash [32]byte, _newInbox [32]byte, _messages []byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.ConfirmUnanimousAsserted(&_ArbChannel.TransactOpts, _afterHash, _newInbox, _messages)
}

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0x812fa865.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _afterHash, bytes32 _newInbox, bytes _messages, bytes32 _logsAccHash, bytes _signatures) returns()
func (_ArbChannel *ArbChannelTransactor) FinalizedUnanimousAssert(opts *bind.TransactOpts, _afterHash [32]byte, _newInbox [32]byte, _messages []byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "finalizedUnanimousAssert", _afterHash, _newInbox, _messages, _logsAccHash, _signatures)
}

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0x812fa865.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _afterHash, bytes32 _newInbox, bytes _messages, bytes32 _logsAccHash, bytes _signatures) returns()
func (_ArbChannel *ArbChannelSession) FinalizedUnanimousAssert(_afterHash [32]byte, _newInbox [32]byte, _messages []byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.FinalizedUnanimousAssert(&_ArbChannel.TransactOpts, _afterHash, _newInbox, _messages, _logsAccHash, _signatures)
}

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0x812fa865.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _afterHash, bytes32 _newInbox, bytes _messages, bytes32 _logsAccHash, bytes _signatures) returns()
func (_ArbChannel *ArbChannelTransactorSession) FinalizedUnanimousAssert(_afterHash [32]byte, _newInbox [32]byte, _messages []byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.FinalizedUnanimousAssert(&_ArbChannel.TransactOpts, _afterHash, _newInbox, _messages, _logsAccHash, _signatures)
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_ArbChannel *ArbChannelTransactor) IncreaseDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "increaseDeposit")
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_ArbChannel *ArbChannelSession) IncreaseDeposit() (*types.Transaction, error) {
	return _ArbChannel.Contract.IncreaseDeposit(&_ArbChannel.TransactOpts)
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_ArbChannel *ArbChannelTransactorSession) IncreaseDeposit() (*types.Transaction, error) {
	return _ArbChannel.Contract.IncreaseDeposit(&_ArbChannel.TransactOpts)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbChannel *ArbChannelTransactor) InitiateChallenge(opts *bind.TransactOpts, _assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "initiateChallenge", _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbChannel *ArbChannelSession) InitiateChallenge(_assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.InitiateChallenge(&_ArbChannel.TransactOpts, _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_ArbChannel *ArbChannelTransactorSession) InitiateChallenge(_assertPreHash [32]byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.InitiateChallenge(&_ArbChannel.TransactOpts, _assertPreHash)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbChannel *ArbChannelTransactor) OwnerShutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "ownerShutdown")
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbChannel *ArbChannelSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbChannel.Contract.OwnerShutdown(&_ArbChannel.TransactOpts)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbChannel *ArbChannelTransactorSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbChannel.Contract.OwnerShutdown(&_ArbChannel.TransactOpts)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbChannel *ArbChannelTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "pendingDisputableAssert", _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbChannel *ArbChannelSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChannel.Contract.PendingDisputableAssert(&_ArbChannel.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0xfec5a286.
//
// Solidity: function pendingDisputableAssert(bytes32 _beforeHash, bytes32 _beforeInbox, bytes32 _afterHash, bytes32 _messagesAccHash, bytes32 _logsAccHash, uint32 _numSteps, uint64[2] _timeBounds) returns()
func (_ArbChannel *ArbChannelTransactorSession) PendingDisputableAssert(_beforeHash [32]byte, _beforeInbox [32]byte, _afterHash [32]byte, _messagesAccHash [32]byte, _logsAccHash [32]byte, _numSteps uint32, _timeBounds [2]uint64) (*types.Transaction, error) {
	return _ArbChannel.Contract.PendingDisputableAssert(&_ArbChannel.TransactOpts, _beforeHash, _beforeInbox, _afterHash, _messagesAccHash, _logsAccHash, _numSteps, _timeBounds)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0xdf949904.
//
// Solidity: function pendingUnanimousAssert(bytes32 _unanRest, uint64 _sequenceNum, bytes32 _messagesAccHash, bytes32 _logsAccHash, bytes _signatures) returns()
func (_ArbChannel *ArbChannelTransactor) PendingUnanimousAssert(opts *bind.TransactOpts, _unanRest [32]byte, _sequenceNum uint64, _messagesAccHash [32]byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _ArbChannel.contract.Transact(opts, "pendingUnanimousAssert", _unanRest, _sequenceNum, _messagesAccHash, _logsAccHash, _signatures)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0xdf949904.
//
// Solidity: function pendingUnanimousAssert(bytes32 _unanRest, uint64 _sequenceNum, bytes32 _messagesAccHash, bytes32 _logsAccHash, bytes _signatures) returns()
func (_ArbChannel *ArbChannelSession) PendingUnanimousAssert(_unanRest [32]byte, _sequenceNum uint64, _messagesAccHash [32]byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.PendingUnanimousAssert(&_ArbChannel.TransactOpts, _unanRest, _sequenceNum, _messagesAccHash, _logsAccHash, _signatures)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0xdf949904.
//
// Solidity: function pendingUnanimousAssert(bytes32 _unanRest, uint64 _sequenceNum, bytes32 _messagesAccHash, bytes32 _logsAccHash, bytes _signatures) returns()
func (_ArbChannel *ArbChannelTransactorSession) PendingUnanimousAssert(_unanRest [32]byte, _sequenceNum uint64, _messagesAccHash [32]byte, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _ArbChannel.Contract.PendingUnanimousAssert(&_ArbChannel.TransactOpts, _unanRest, _sequenceNum, _messagesAccHash, _logsAccHash, _signatures)
}

// ArbChannelConfirmedDisputableAssertionIterator is returned from FilterConfirmedDisputableAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedDisputableAssertion events raised by the ArbChannel contract.
type ArbChannelConfirmedDisputableAssertionIterator struct {
	Event *ArbChannelConfirmedDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChannelConfirmedDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChannelConfirmedDisputableAssertion)
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
		it.Event = new(ArbChannelConfirmedDisputableAssertion)
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
func (it *ArbChannelConfirmedDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChannelConfirmedDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChannelConfirmedDisputableAssertion represents a ConfirmedDisputableAssertion event raised by the ArbChannel contract.
type ArbChannelConfirmedDisputableAssertion struct {
	NewState    [32]byte
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbChannel *ArbChannelFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts) (*ArbChannelConfirmedDisputableAssertionIterator, error) {

	logs, sub, err := _ArbChannel.contract.FilterLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChannelConfirmedDisputableAssertionIterator{contract: _ArbChannel.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_ArbChannel *ArbChannelFilterer) WatchConfirmedDisputableAssertion(opts *bind.WatchOpts, sink chan<- *ArbChannelConfirmedDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChannel.contract.WatchLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChannelConfirmedDisputableAssertion)
				if err := _ArbChannel.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
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
func (_ArbChannel *ArbChannelFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*ArbChannelConfirmedDisputableAssertion, error) {
	event := new(ArbChannelConfirmedDisputableAssertion)
	if err := _ArbChannel.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChannelConfirmedUnanimousAssertionIterator is returned from FilterConfirmedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedUnanimousAssertion events raised by the ArbChannel contract.
type ArbChannelConfirmedUnanimousAssertionIterator struct {
	Event *ArbChannelConfirmedUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChannelConfirmedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChannelConfirmedUnanimousAssertion)
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
		it.Event = new(ArbChannelConfirmedUnanimousAssertion)
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
func (it *ArbChannelConfirmedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChannelConfirmedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChannelConfirmedUnanimousAssertion represents a ConfirmedUnanimousAssertion event raised by the ArbChannel contract.
type ArbChannelConfirmedUnanimousAssertion struct {
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedUnanimousAssertion is a free log retrieval operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_ArbChannel *ArbChannelFilterer) FilterConfirmedUnanimousAssertion(opts *bind.FilterOpts) (*ArbChannelConfirmedUnanimousAssertionIterator, error) {

	logs, sub, err := _ArbChannel.contract.FilterLogs(opts, "ConfirmedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChannelConfirmedUnanimousAssertionIterator{contract: _ArbChannel.contract, event: "ConfirmedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedUnanimousAssertion is a free log subscription operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_ArbChannel *ArbChannelFilterer) WatchConfirmedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *ArbChannelConfirmedUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChannel.contract.WatchLogs(opts, "ConfirmedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChannelConfirmedUnanimousAssertion)
				if err := _ArbChannel.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
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

// ParseConfirmedUnanimousAssertion is a log parse operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_ArbChannel *ArbChannelFilterer) ParseConfirmedUnanimousAssertion(log types.Log) (*ArbChannelConfirmedUnanimousAssertion, error) {
	event := new(ArbChannelConfirmedUnanimousAssertion)
	if err := _ArbChannel.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChannelFinalizedUnanimousAssertionIterator is returned from FilterFinalizedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for FinalizedUnanimousAssertion events raised by the ArbChannel contract.
type ArbChannelFinalizedUnanimousAssertionIterator struct {
	Event *ArbChannelFinalizedUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChannelFinalizedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChannelFinalizedUnanimousAssertion)
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
		it.Event = new(ArbChannelFinalizedUnanimousAssertion)
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
func (it *ArbChannelFinalizedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChannelFinalizedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChannelFinalizedUnanimousAssertion represents a FinalizedUnanimousAssertion event raised by the ArbChannel contract.
type ArbChannelFinalizedUnanimousAssertion struct {
	UnanHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizedUnanimousAssertion is a free log retrieval operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_ArbChannel *ArbChannelFilterer) FilterFinalizedUnanimousAssertion(opts *bind.FilterOpts) (*ArbChannelFinalizedUnanimousAssertionIterator, error) {

	logs, sub, err := _ArbChannel.contract.FilterLogs(opts, "FinalizedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChannelFinalizedUnanimousAssertionIterator{contract: _ArbChannel.contract, event: "FinalizedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchFinalizedUnanimousAssertion is a free log subscription operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_ArbChannel *ArbChannelFilterer) WatchFinalizedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *ArbChannelFinalizedUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChannel.contract.WatchLogs(opts, "FinalizedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChannelFinalizedUnanimousAssertion)
				if err := _ArbChannel.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
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

// ParseFinalizedUnanimousAssertion is a log parse operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_ArbChannel *ArbChannelFilterer) ParseFinalizedUnanimousAssertion(log types.Log) (*ArbChannelFinalizedUnanimousAssertion, error) {
	event := new(ArbChannelFinalizedUnanimousAssertion)
	if err := _ArbChannel.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChannelPendingAssertionCanceledIterator is returned from FilterPendingAssertionCanceled and is used to iterate over the raw logs and unpacked data for PendingAssertionCanceled events raised by the ArbChannel contract.
type ArbChannelPendingAssertionCanceledIterator struct {
	Event *ArbChannelPendingAssertionCanceled // Event containing the contract specifics and raw log

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
func (it *ArbChannelPendingAssertionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChannelPendingAssertionCanceled)
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
		it.Event = new(ArbChannelPendingAssertionCanceled)
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
func (it *ArbChannelPendingAssertionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChannelPendingAssertionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChannelPendingAssertionCanceled represents a PendingAssertionCanceled event raised by the ArbChannel contract.
type ArbChannelPendingAssertionCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPendingAssertionCanceled is a free log retrieval operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_ArbChannel *ArbChannelFilterer) FilterPendingAssertionCanceled(opts *bind.FilterOpts) (*ArbChannelPendingAssertionCanceledIterator, error) {

	logs, sub, err := _ArbChannel.contract.FilterLogs(opts, "PendingAssertionCanceled")
	if err != nil {
		return nil, err
	}
	return &ArbChannelPendingAssertionCanceledIterator{contract: _ArbChannel.contract, event: "PendingAssertionCanceled", logs: logs, sub: sub}, nil
}

// WatchPendingAssertionCanceled is a free log subscription operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_ArbChannel *ArbChannelFilterer) WatchPendingAssertionCanceled(opts *bind.WatchOpts, sink chan<- *ArbChannelPendingAssertionCanceled) (event.Subscription, error) {

	logs, sub, err := _ArbChannel.contract.WatchLogs(opts, "PendingAssertionCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChannelPendingAssertionCanceled)
				if err := _ArbChannel.contract.UnpackLog(event, "PendingAssertionCanceled", log); err != nil {
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

// ParsePendingAssertionCanceled is a log parse operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_ArbChannel *ArbChannelFilterer) ParsePendingAssertionCanceled(log types.Log) (*ArbChannelPendingAssertionCanceled, error) {
	event := new(ArbChannelPendingAssertionCanceled)
	if err := _ArbChannel.contract.UnpackLog(event, "PendingAssertionCanceled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChannelPendingDisputableAssertionIterator is returned from FilterPendingDisputableAssertion and is used to iterate over the raw logs and unpacked data for PendingDisputableAssertion events raised by the ArbChannel contract.
type ArbChannelPendingDisputableAssertionIterator struct {
	Event *ArbChannelPendingDisputableAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChannelPendingDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChannelPendingDisputableAssertion)
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
		it.Event = new(ArbChannelPendingDisputableAssertion)
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
func (it *ArbChannelPendingDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChannelPendingDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChannelPendingDisputableAssertion represents a PendingDisputableAssertion event raised by the ArbChannel contract.
type ArbChannelPendingDisputableAssertion struct {
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
func (_ArbChannel *ArbChannelFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*ArbChannelPendingDisputableAssertionIterator, error) {

	logs, sub, err := _ArbChannel.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChannelPendingDisputableAssertionIterator{contract: _ArbChannel.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe24436.
//
// Solidity: event PendingDisputableAssertion(bytes32[5] fields, address asserter, uint64[2] timeBounds, uint32 numSteps, uint64 deadline)
func (_ArbChannel *ArbChannelFilterer) WatchPendingDisputableAssertion(opts *bind.WatchOpts, sink chan<- *ArbChannelPendingDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChannel.contract.WatchLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChannelPendingDisputableAssertion)
				if err := _ArbChannel.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
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
func (_ArbChannel *ArbChannelFilterer) ParsePendingDisputableAssertion(log types.Log) (*ArbChannelPendingDisputableAssertion, error) {
	event := new(ArbChannelPendingDisputableAssertion)
	if err := _ArbChannel.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChannelPendingUnanimousAssertionIterator is returned from FilterPendingUnanimousAssertion and is used to iterate over the raw logs and unpacked data for PendingUnanimousAssertion events raised by the ArbChannel contract.
type ArbChannelPendingUnanimousAssertionIterator struct {
	Event *ArbChannelPendingUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChannelPendingUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChannelPendingUnanimousAssertion)
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
		it.Event = new(ArbChannelPendingUnanimousAssertion)
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
func (it *ArbChannelPendingUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChannelPendingUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChannelPendingUnanimousAssertion represents a PendingUnanimousAssertion event raised by the ArbChannel contract.
type ArbChannelPendingUnanimousAssertion struct {
	UnanHash    [32]byte
	SequenceNum uint64
	Deadline    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingUnanimousAssertion is a free log retrieval operation binding the contract event 0x4c6950de8aaa67cd052f9e28572dfff2ec4b8badd2f2b4bd8d8479d36987b6a4.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum, uint64 deadline)
func (_ArbChannel *ArbChannelFilterer) FilterPendingUnanimousAssertion(opts *bind.FilterOpts) (*ArbChannelPendingUnanimousAssertionIterator, error) {

	logs, sub, err := _ArbChannel.contract.FilterLogs(opts, "PendingUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChannelPendingUnanimousAssertionIterator{contract: _ArbChannel.contract, event: "PendingUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingUnanimousAssertion is a free log subscription operation binding the contract event 0x4c6950de8aaa67cd052f9e28572dfff2ec4b8badd2f2b4bd8d8479d36987b6a4.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum, uint64 deadline)
func (_ArbChannel *ArbChannelFilterer) WatchPendingUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *ArbChannelPendingUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChannel.contract.WatchLogs(opts, "PendingUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChannelPendingUnanimousAssertion)
				if err := _ArbChannel.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
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

// ParsePendingUnanimousAssertion is a log parse operation binding the contract event 0x4c6950de8aaa67cd052f9e28572dfff2ec4b8badd2f2b4bd8d8479d36987b6a4.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum, uint64 deadline)
func (_ArbChannel *ArbChannelFilterer) ParsePendingUnanimousAssertion(log types.Log) (*ArbChannelPendingUnanimousAssertion, error) {
	event := new(ArbChannelPendingUnanimousAssertion)
	if err := _ArbChannel.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
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
var ArbProtocolBin = "0x610a6e610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100555760003560e01c80624c28f61461005a5780637ddf59d6146100b257806385ecb92a146100f3578063e83f4bfe14610148575b600080fd5b6100a06004803603608081101561007057600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b03166101ee565b60408051918252519081900360200190f35b6100a0600480360360c08110156100c857600080fd5b5080359063ffffffff6020820135169060408101359060608101359060808101359060a001356102e0565b6100a06004803603608081101561010957600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091945050903591506103389050565b6100a06004803603602081101561015e57600080fd5b81019060208101813564010000000081111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460018302840111640100000000831117156101ad57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061038c945050505050565b60408051600480825260a0820190925260009160609190816020015b6102126109d2565b81526020019060019003908161020a579050509050610230866104d1565b8160008151811061023d57fe5b602002602001018190525061025a836001600160a01b031661054f565b8160018151811061026757fe5b602002602001018190525061027b8461054f565b8160028151811061028857fe5b60209081029190910101526102aa6affffffffffffffffffffff19861661054f565b816003815181106102b757fe5b60200260200101819052506102d36102ce826105cd565b61067d565b519150505b949350505050565b6040805160208082019890985260e09690961b6001600160e01b0319168682015260448601949094526064850192909252608484015260a4808401919091528151808403909101815260c49092019052805191012090565b815160209283015160408051808601969096526001600160c01b031960c093841b8116878301529190921b166048850152605080850192909252805180850390920182526070909301909252815191012090565b8051600090819081908190815b818110156104c45773__$d969135829891f807aa9c34494da4ecd99$__6389df40da88866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561040f5781810151838201526020016103f7565b50505050905090810190601f16801561043c5780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561045957600080fd5b505af415801561046d573d6000803e3d6000fd5b505050506040513d604081101561048357600080fd5b50805160209182015160408051808501999099528881018290528051808a038201815260609099019052875197909201969096209594509250600101610399565b509293505050505b919050565b6104d96109d2565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161053e565b61052b6109d2565b8152602001906001900390816105235790505b508152600260209091015292915050565b6105576109d2565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916105bc565b6105a96109d2565b8152602001906001900390816105a15790505b508152600060209091015292915050565b6105d56109d2565b6105df82516107b3565b610630576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b610685610a00565b6060820151600c60ff909116106106d7576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166107045760405180602001604052806106fb84600001516107ba565b905290506104cc565b606082015160ff166001141561074b5760405180602001604052806106fb8460200151600001518560200151604001518660200151606001518760200151602001516107de565b606082015160ff166002141561077057506040805160208101909152815181526104cc565b600360ff16826060015160ff161015801561079457506060820151600c60ff909116105b156107b15760405180602001604052806106fb8460400151610886565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315610838575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206102d8565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b60006008825111156108d6576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610903578160200160208202803883390190505b50805190915060005b8181101561095f5761091c610a00565b61093886838151811061092b57fe5b602002602001015161067d565b9050806000015184838151811061094b57fe5b60209081029190910101525060010161090c565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156109a8578181015183820152602001610990565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6040518060800160405280600081526020016109ec610a12565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a72315820d971e9786c692cf0fa3a7420eb5705fea283f11fc02b843ab5b19410c6c1150c64736f6c634300050c0032"

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
var ArbValueBin = "0x61152b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c806353409fab1161007057806353409fab14610381578063826513e0146103a757806389df40da146103db5780638f3460361461049a578063b2b9dc621461053e576100a8565b80631667b411146100ad5780631f3d4d4e146100dc578063264f384b1461020157806332e6cc211461022d578063364df27714610379575b600080fd5b6100ca600480360360208110156100c357600080fd5b503561056f565b60408051918252519081900360200190f35b610182600480360360408110156100f257600080fd5b810190602081018135600160201b81111561010c57600080fd5b82018360208201111561011e57600080fd5b803590602001918460018302840111600160201b8311171561013f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610595915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101c55781810151838201526020016101ad565b50505050905090810190601f1680156101f25780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100ca6004803603606081101561021757600080fd5b5060ff8135169060208101359060400135610619565b6102d36004803603604081101561024357600080fd5b810190602081018135600160201b81111561025d57600080fd5b82018360208201111561026f57600080fd5b803590602001918460018302840111600160201b8311171561029057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061066b915050565b604051808815151515815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610338578181015183820152602001610320565b50505050905090810190601f1680156103655780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6100ca610869565b6100ca6004803603604081101561039757600080fd5b5060ff81351690602001356108dc565b6100ca600480360360808110156103bd57600080fd5b5060ff81351690602081013515159060408101359060600135610923565b610481600480360360408110156103f157600080fd5b810190602081018135600160201b81111561040b57600080fd5b82018360208201111561041d57600080fd5b803590602001918460018302840111600160201b8311171561043e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506109cc915050565b6040805192835260208301919091528051918290030190f35b6100ca600480360360208110156104b057600080fd5b810190602081018135600160201b8111156104ca57600080fd5b8201836020820111156104dc57600080fd5b803590602001918460018302840111600160201b831117156104fd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a45945050505050565b61055b6004803603602081101561055457600080fd5b5035610ab7565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b600060606000806105a461146f565b6105ae8787610abe565b9194509250905082156105f6576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b8161060a888880840363ffffffff610c4816565b945094505050505b9250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6000806000806000806060600088965060008a888151811061068957fe5b016020015160019098019760f81c9050600781146106bb576106ae8b60018a036109cc565b909850965061085d915050565b6106c58b896109cc565b90985091506106e48b60018c016000198d8c030163ffffffff610c4816565b92508a88815181106106f257fe5b016020015160019098019760f81c90508015610715576106ae8b60018a036109cc565b61071f8b89610cc8565b80995081975050508a888151811061073357fe5b016020015160019098019760f81c90508015610756576106ae8b60018a036109cc565b6107608b89610cc8565b80995081965050508a888151811061077457fe5b016020015160019098019760f81c90508015610797576106ae8b60018a036109cc565b6107a18b89610cc8565b60408051600480825260a0820190925260019c50919a509195506060916020820160808038833901905050905082816000815181106107dc57fe5b6020026020010181815250506107f18761056f565b816001815181106107fe57fe5b6020026020010181815250506108138661056f565b8160028151811061082057fe5b6020026020010181815250506108358561056f565b8160038151811061084257fe5b60200260200101818152505061085781610cef565b97505050505b92959891949750929550565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156108b557818101518382015260200161089d565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000831561097d575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206109c4565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b6000806000806109da61146f565b6109e48787610abe565b919450925090508215610a2c576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b81610a3682610daf565b51909890975095505050505050565b60008080610a5161146f565b610a5c856000610abe565b919450925090508215610aa4576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b610aad81610daf565b5195945050505050565b6008101590565b600080610ac961146f565b84518410610b1e576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110610b3157fe5b016020015160019092019160f81c90506000610b4b61149d565b60ff8316610b7f57610b5d8985610cc8565b9094509150600084610b6e84610ee5565b91985096509450610c419350505050565b60ff831660011415610ba657610b958985610f63565b9094509050600084610b6e836110be565b60ff831660021415610bcd57610bbc8985610cc8565b9094509150600084610b6e8461111e565b600360ff841610801590610be45750600c60ff8416105b15610c2157600219830160606000610bfd838d8961119c565b909850925090508087610c0f84611257565b99509950995050505050505050610c41565b8260ff16612710016000610c356000610ee5565b91985096509450505050505b9250925092565b606081830184511015610c5a57600080fd5b606082158015610c7557604051915060208201604052610cbf565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610cae578051835260209283019201610c96565b5050858452601f01601f1916604052505b50949350505050565b6000808281610cdd868363ffffffff61130716565b60209290920196919550909350505050565b6000600882511115610d3f576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b8151600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610d87578181015183820152602001610d6f565b5050505090500192505050604051602081830303815290604052805190602001209050919050565b610db76114c4565b6060820151600c60ff90911610610e09576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610e36576040518060200160405280610e2d846000015161056f565b90529050610590565b606082015160ff1660011415610e7d576040518060200160405280610e2d846020015160000151856020015160400151866020015160600151876020015160200151610923565b606082015160ff1660021415610ea25750604080516020810190915281518152610590565b600360ff16826060015160ff1610158015610ec657506060820151600c60ff909116105b15610ee3576040518060200160405280610e2d8460400151611323565bfe5b610eed61146f565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610f52565b610f3f61146f565b815260200190600190039081610f375790505b508152600060209091015292915050565b6000610f6d61149d565b60008390506000858281518110610f8057fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110610fa657fe5b016020015160019384019360f89190911c915060009060ff84161415611032576000610fd061146f565b610fda8a87610abe565b90975090925090508115611023576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61102c81610daf565b51925050505b6000611044898663ffffffff61130716565b90506020850194508360ff1660011415611089576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506106129050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6110c661146f565b60408051608081018252600080825260208083018690528351828152908101845291928301919061110d565b6110fa61146f565b8152602001906001900390816110f25790505b508152600160209091015292915050565b61112661146f565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161118b565b61117861146f565b8152602001906001900390816111705790505b508152600260209091015292915050565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156111e757816020015b6111d461146f565b8152602001906001900390816111cc5790505b50905060005b8960ff168160ff161015611241576112058985610abe565b8451859060ff861690811061121657fe5b60209081029190910101529450925082156112395750909450909250905061124e565b6001016111ed565b5060009550919350909150505b93509350939050565b61125f61146f565b6112698251610ab7565b6112ba576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6000816020018351101561131a57600080fd5b50016020015190565b6000600882511115611373576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156113a0578160200160208202803883390190505b50805190915060005b818110156113fc576113b96114c4565b6113d58683815181106113c857fe5b6020026020010151610daf565b905080600001518483815181106113e857fe5b6020908102919091010152506001016113a9565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561144557818101518382015260200161142d565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180608001604052806000815260200161148961149d565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4d61727368616c6c65642076616c7565206d7573742062652076616c69640000a265627a7a723158205e0803af87b5938fad79b90ef8be05d8e5e117e4a54ef586de2435655fbfae2b64736f6c634300050c0032"

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
const ArbitrumVMABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PendingAssertionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeManager\",\"outputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumVM.State\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalPendingInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_messagesAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"enumVM.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"inChallenge\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbitrumVMFuncSigs maps the 4-byte function signature to its string representation.
var ArbitrumVMFuncSigs = map[string]string{
	"94af716b": "activateVM()",
	"023a96fe": "challengeManager()",
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
var ArbitrumVMBin = "0x608060405234801561001057600080fd5b50604051611661380380611661833981810160405260e081101561003357600080fd5b50805160208201516040808401516060850151608086015160a087015160c090970151600180546001600160a01b038084166001600160a01b0319928316179283905560008054828d16931692909217825587517ff39723830000000000000000000000000000000000000000000000000000000081529751999a9899969895979496939492169263f397238392600480820193929182900301818387803b1580156100de57600080fd5b505af11580156100f2573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b15801561017957600080fd5b505af415801561018d573d6000803e3d6000fd5b505050506040513d60208110156101a357600080fd5b50516004555050600680546001600160801b0319166001600160801b039390931692909217909155506007805463ffffffff191663ffffffff9384161763ffffffff60201b19166401000000009290931691909102919091179055506114538061020e6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80636be0022911610097578063cfa8070711610066578063cfa8070714610286578063d489113a1461028e578063f2204f7414610296578063fec5a28614610357576100f5565b80636be00229146102665780638da5cb5b1461026e57806394af716b14610276578063aca0f3721461027e576100f5565b806322c091bc116100d357806322c091bc146101825780632782e87e146101a45780633a768463146101c157806360675a871461025e576100f5565b8063023a96fe146100fa57806308dc89d71461011e5780631865c57d14610156575b600080fd5b6101026103ce565b604080516001600160a01b039092168252519081900360200190f35b6101446004803603602081101561013457600080fd5b50356001600160a01b03166103dd565b60408051918252519081900360200190f35b61015e6103fc565b6040518082600381111561016e57fe5b60ff16815260200191505060405180910390f35b6101a26004803603608081101561019857600080fd5b506040810161040c565b005b6101a2600480360360208110156101ba57600080fd5b503561055f565b6101c9610763565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561023957fe5b60ff1681529115156020830152506040805191829003019a5098505050505050505050f35b6101026107d5565b6101026107e4565b6101026107f3565b6101a2610802565b610144610895565b6101a26108a4565b610102610904565b6101a2600480360360a08110156102ac57600080fd5b81359160208101359163ffffffff60408301351691908101906080810160608201356401000000008111156102e057600080fd5b8201836020820111156102f257600080fd5b8035906020019184600183028401116401000000008311171561031457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610913915050565b6101a2600480360361010081101561036e57600080fd5b604080518082018252833593602081013593838201359360608301359360808401359363ffffffff60a08201351693810192909161010083019160c08401906002908390839080828437600092019190915250919450610a6c9350505050565b6000546001600160a01b031681565b6001600160a01b0381166000908152600860205260409020545b919050565b600754600160401b900460ff1690565b6000546001600160a01b031633146104555760405162461bcd60e51b815260040180806020018281038252602d8152602001806113c1602d913960400191505060405180910390fd5b600754600160481b900460ff1661049d5760405162461bcd60e51b815260040180806020018281038252602681526020018061139b6026913960400191505060405180910390fd5b6007805469ff000000000000000000191690556105026001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002054610bcd90919063ffffffff16565b82356001600160a01b0316600090815260086020818152604083209390935561053a928401356001600160801b0316918560016104c5565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b031611156105b85760405162461bcd60e51b81526004018080602001828103825260278152602001806113746027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b15801561063257600080fd5b505af4158015610646573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b0316808252918101919091526007548451630823813560e21b815292909516975063208e04d496509194919363ffffffff16928892600490920191829187918190849084905b838110156106de5781810151838201526020016106c6565b5050505090500184600260200280838360005b838110156107095781810151838201526020016106f1565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050600060405180830381600087803b15801561074857600080fd5b505af115801561075c573d6000803e3d6000fd5b5050505050565b6002546003546004546005546006546007546001600160a01b03909216916001600160801b0382169167ffffffffffffffff600160801b8204811692600160c01b909204169063ffffffff8082169164010000000081049091169060ff600160401b8204811691600160481b9004168b565b600a546001600160a01b031681565b6009546001600160a01b031681565b600b546001600160a01b031681565b600b546001600160a01b0316331461085a576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff16600381111561087457fe5b1415610893576007805468ff00000000000000001916600160401b1790555b565b6006546001600160801b031690565b600b546001600160a01b031633146108fc576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b610893610c2e565b6001546001600160a01b031681565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63ce9d5122600287878787876040518763ffffffff1660e01b8152600401808781526020018681526020018581526020018463ffffffff1663ffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156109aa578181015183820152602001610992565b50505050905090810190601f1680156109d75780820380516001836020036101000a031916815260200191505b5097505050505050505060006040518083038186803b1580156109f957600080fd5b505af4158015610a0d573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054610a48935091506001600160801b031663ffffffff610bcd16565b6005546001600160a01b031660009081526008602052604090205561075c82610c3c565b336000908152600860205260409020546006546001600160801b03161115610ac55760405162461bcd60e51b81526004018080602001828103825260318152602001806113ee6031913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151636c36f28f60e11b8152600260048201818152602483018c9052604483018b9052606483018a90526084830189905260a4830188905263ffffffff871660c484015273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__9463d86de51e9492938d938d938d938d938d938d938d93909260e401918491908190849084905b83811015610b83578181015183820152602001610b6b565b505050509050019850505050505050505060006040518083038186803b158015610bac57600080fd5b505af4158015610bc0573d6000803e3d6000fd5b5050505050505050505050565b600082820183811015610c27576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600b546001600160a01b0316ff5b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b158015610c8257600080fd5b505af1158015610c96573d6000803e3d6000fd5b505050506040513d6020811015610cac57600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b158015610cfd57600080fd5b505af4158015610d11573d6000803e3d6000fd5b505050506040513d6020811015610d2757600080fd5b50518114610d7257610d6e6040518060600160405280610d476001610e3a565b8152602001610d596002800154610eb8565b8152602001610d6784610eb8565b9052610f36565b6004555b60015460405163e4eb8c6360e01b81526020600482018181528551602484015285516001600160a01b039094169363e4eb8c6393879383926044909201919085019080838360005b83811015610dd2578181015183820152602001610dba565b50505050905090810190601f168015610dff5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610e1e57600080fd5b505af1158015610e32573d6000803e3d6000fd5b505050505050565b610e4261130c565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610ea7565b610e9461130c565b815260200190600190039081610e8c5790505b508152600060209091015292915050565b610ec061130c565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610f25565b610f1261130c565b815260200190600190039081610f0a5790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b610f5a61130c565b815260200190600190039081610f52575050805190915060005b81811015610fac57848160038110610f8857fe5b6020020151838281518110610f9957fe5b6020908102919091010152600101610f74565b50610fb682610fbe565b949350505050565b600060088251111561100e576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825160405190808252806020026020018201604052801561103b578160200160208202803883390190505b50805190915060005b818110156110975761105461133a565b61107086838151811061106357fe5b602002602001015161110a565b9050806000015184838151811061108357fe5b602090810291909101015250600101611044565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156110e05781810151838201526020016110c8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b61111261133a565b6060820151600c60ff90911610611164576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166111915760405180602001604052806111888460000151611240565b905290506103f7565b606082015160ff16600114156111d8576040518060200160405280611188846020015160000151856020015160400151866020015160600151876020015160200151611264565b606082015160ff16600214156111fd57506040805160208101909152815181526103f7565b600360ff16826060015160ff161015801561122157506060820151600c60ff909116105b1561123e5760405180602001604052806111888460400151610fbe565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b600083156112be575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120610fb6565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b60405180608001604052806000815260200161132661134c565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274a265627a7a723158200a610020ba1e092b2238909f12507db53abe801d037c9594d3c13991e04b194464736f6c634300050c0032"

// DeployArbitrumVM deploys a new Ethereum contract, binding an instance of ArbitrumVM to it.
func DeployArbitrumVM(auth *bind.TransactOpts, backend bind.ContractBackend, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _challengeManagerAddress common.Address, _globalInboxAddress common.Address) (common.Address, *types.Transaction, *ArbitrumVM, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrumVMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ArbitrumVMBin = strings.Replace(ArbitrumVMBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

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

// ArbitrumVMPendingAssertionCanceledIterator is returned from FilterPendingAssertionCanceled and is used to iterate over the raw logs and unpacked data for PendingAssertionCanceled events raised by the ArbitrumVM contract.
type ArbitrumVMPendingAssertionCanceledIterator struct {
	Event *ArbitrumVMPendingAssertionCanceled // Event containing the contract specifics and raw log

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
func (it *ArbitrumVMPendingAssertionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumVMPendingAssertionCanceled)
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
		it.Event = new(ArbitrumVMPendingAssertionCanceled)
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
func (it *ArbitrumVMPendingAssertionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrumVMPendingAssertionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrumVMPendingAssertionCanceled represents a PendingAssertionCanceled event raised by the ArbitrumVM contract.
type ArbitrumVMPendingAssertionCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPendingAssertionCanceled is a free log retrieval operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_ArbitrumVM *ArbitrumVMFilterer) FilterPendingAssertionCanceled(opts *bind.FilterOpts) (*ArbitrumVMPendingAssertionCanceledIterator, error) {

	logs, sub, err := _ArbitrumVM.contract.FilterLogs(opts, "PendingAssertionCanceled")
	if err != nil {
		return nil, err
	}
	return &ArbitrumVMPendingAssertionCanceledIterator{contract: _ArbitrumVM.contract, event: "PendingAssertionCanceled", logs: logs, sub: sub}, nil
}

// WatchPendingAssertionCanceled is a free log subscription operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_ArbitrumVM *ArbitrumVMFilterer) WatchPendingAssertionCanceled(opts *bind.WatchOpts, sink chan<- *ArbitrumVMPendingAssertionCanceled) (event.Subscription, error) {

	logs, sub, err := _ArbitrumVM.contract.WatchLogs(opts, "PendingAssertionCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrumVMPendingAssertionCanceled)
				if err := _ArbitrumVM.contract.UnpackLog(event, "PendingAssertionCanceled", log); err != nil {
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

// ParsePendingAssertionCanceled is a log parse operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_ArbitrumVM *ArbitrumVMFilterer) ParsePendingAssertionCanceled(log types.Log) (*ArbitrumVMPendingAssertionCanceled, error) {
	event := new(ArbitrumVMPendingAssertionCanceled)
	if err := _ArbitrumVM.contract.UnpackLog(event, "PendingAssertionCanceled", log); err != nil {
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

// ChannelLauncherABI is the input ABI used to generate the binding from.
const ChannelLauncherABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeManagerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"name\":\"ChannelCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_validatorKeys\",\"type\":\"address[]\"}],\"name\":\"launchChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChannelLauncherFuncSigs maps the 4-byte function signature to its string representation.
var ChannelLauncherFuncSigs = map[string]string{
	"ad7d0fa4": "launchChannel(bytes32,uint32,uint32,uint128,address,address[])",
}

// ChannelLauncherBin is the compiled bytecode used for deploying new contracts.
var ChannelLauncherBin = "0x608060405234801561001057600080fd5b506040516126753803806126758339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556125fb8061007a6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ad7d0fa414610030575b600080fd5b61010e600480360360c081101561004657600080fd5b81359163ffffffff60208201358116926040830135909116916001600160801b03606082013516916001600160a01b03608083013516919081019060c0810160a082013564010000000081111561009c57600080fd5b8201836020820111156100ae57600080fd5b803590602001918460208302840111640100000000831117156100d057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610110945050505050565b005b600154600080546040519192899289928992899289926001600160a01b03928316921690899061013f90610241565b88815263ffffffff80891660208084019190915290881660408301526001600160801b03871660608301526001600160a01b03808716608084015285811660a0840152841660c083015261010060e0830181815284519184019190915283519091610120840191858201910280838360005b838110156101c95781810151838201526020016101b1565b505050509050019950505050505050505050604051809103906000f0801580156101f7573d6000803e3d6000fd5b50604080516001600160a01b038316815290519192507fc625d37dd8b556110d70984e62f74ba35c77422c83c5f548fbd21b697a67ef5c919081900360200190a150505050505050565b6123788061024f8339019056fe60806040523480156200001157600080fd5b50604051620023783803806200237883398181016040526101008110156200003857600080fd5b815160208301516040808501516060860151608087015160a088015160c089015160e08a0180519651989a9799959894979396929591949391820192846401000000008211156200008857600080fd5b9083019060208201858111156200009e57600080fd5b8251866020820283011164010000000082111715620000bc57600080fd5b82525081516020918201928201910280838360005b83811015620000eb578181015183820152602001620000d1565b505050509050016040525050508787878787878780600160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000806101000a8154816001600160a01b0302191690836001600160a01b03160217905550600160009054906101000a90046001600160a01b03166001600160a01b031663f39723836040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156200019d57600080fd5b505af1158015620001b2573d6000803e3d6000fd5b5050600b80546001600160a01b0319166001600160a01b03871617905550506002879055600780546000919060ff60401b19166801000000000000000083021790555073__$d969135829891f807aa9c34494da4ecd99$__63364df2776040518163ffffffff1660e01b815260040160206040518083038186803b1580156200023a57600080fd5b505af41580156200024f573d6000803e3d6000fd5b505050506040513d60208110156200026657600080fd5b50516004555050600680546001600160801b039093166001600160801b031990931692909217909155506007805463ffffffff9283166401000000000263ffffffff60201b199390941663ffffffff19918216179290921692909217909155600d8054845161ffff1692169190911790555060005b600d5461ffff908116908216101562000340576001600c6000848461ffff16815181106200030557fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055600101620002db565b50505050505050505061201f80620003596000396000f3fe6080604052600436106101405760003560e01c8063812fa865116100b6578063cfa807071161006f578063cfa80707146105b5578063d489113a146105ca578063df949904146105df578063e1e5d090146106b2578063f2204f741461076f578063fec5a2861461083b57610140565b8063812fa865146103e4578063899b7c741461052e5780638da5cb5b1461054357806394af716b14610558578063aca0f3721461056d578063b99738e01461058257610140565b806322c091bc1161010857806322c091bc1461022a5780632782e87e146102575780633a76846314610281578063513164fe1461032b57806360675a87146103ba5780636be00229146103cf57610140565b8063023a96fe1461014557806305b050de1461017657806308dc89d7146101805780630f43a677146101c55780631865c57d146101f1575b600080fd5b34801561015157600080fd5b5061015a6108bf565b604080516001600160a01b039092168252519081900360200190f35b61017e6108ce565b005b34801561018c57600080fd5b506101b3600480360360208110156101a357600080fd5b50356001600160a01b03166109eb565b60408051918252519081900360200190f35b3480156101d157600080fd5b506101da610a0a565b6040805161ffff9092168252519081900360200190f35b3480156101fd57600080fd5b50610206610a14565b6040518082600381111561021657fe5b60ff16815260200191505060405180910390f35b34801561023657600080fd5b5061017e6004803603608081101561024d57600080fd5b5060408101610a24565b34801561026357600080fd5b5061017e6004803603602081101561027a57600080fd5b5035610b77565b34801561028d57600080fd5b50610296610d7b565b604080518c8152602081018c90529081018a90526001600160a01b03891660608201526001600160801b038816608082015267ffffffffffffffff80881660a0830152861660c082015263ffffffff80861660e08301528416610100820152610120810183600381111561030657fe5b60ff1681529115156020830152506040805191829003019a5098505050505050505050f35b34801561033757600080fd5b506103a66004803603602081101561034e57600080fd5b810190602081018135600160201b81111561036857600080fd5b82018360208201111561037a57600080fd5b803590602001918460208302840111600160201b8311171561039b57600080fd5b509092509050610dec565b604080519115158252519081900360200190f35b3480156103c657600080fd5b5061015a610e73565b3480156103db57600080fd5b5061015a610e82565b3480156103f057600080fd5b5061017e600480360360a081101561040757600080fd5b813591602081013591810190606081016040820135600160201b81111561042d57600080fd5b82018360208201111561043f57600080fd5b803590602001918460018302840111600160201b8311171561046057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b8111156104ba57600080fd5b8201836020820111156104cc57600080fd5b803590602001918460018302840111600160201b831117156104ed57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e91945050505050565b34801561053a57600080fd5b506101da6110ee565b34801561054f57600080fd5b5061015a6110fe565b34801561056457600080fd5b5061017e61110d565b34801561057957600080fd5b506101b361119b565b34801561058e57600080fd5b506103a6600480360360208110156105a557600080fd5b50356001600160a01b03166111aa565b3480156105c157600080fd5b5061017e6111c8565b3480156105d657600080fd5b5061015a611228565b3480156105eb57600080fd5b5061017e600480360360a081101561060257600080fd5b81359167ffffffffffffffff6020820135169160408201359160608101359181019060a081016080820135600160201b81111561063e57600080fd5b82018360208201111561065057600080fd5b803590602001918460018302840111600160201b8311171561067157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611237945050505050565b3480156106be57600080fd5b5061017e600480360360608110156106d557600080fd5b813591602081013591810190606081016040820135600160201b8111156106fb57600080fd5b82018360208201111561070d57600080fd5b803590602001918460018302840111600160201b8311171561072e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113ef945050505050565b34801561077b57600080fd5b5061017e600480360360a081101561079257600080fd5b81359160208101359163ffffffff6040830135169190810190608081016060820135600160201b8111156107c557600080fd5b8201836020820111156107d757600080fd5b803590602001918460018302840111600160201b831117156107f857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506114df915050565b34801561084757600080fd5b5061017e600480360361010081101561085f57600080fd5b604080518082018252833593602081013593838201359360608301359360808401359363ffffffff60a08201351693810192909161010083019160c084019060029083908390808284376000920191909152509194506116389350505050565b6000546001600160a01b031681565b336000908152600c602052604090205460ff16610932576040805162461bcd60e51b815260206004820152601860248201527f43616c6c6572206d7573742062652076616c696461746f720000000000000000604482015290519081900360640190fd5b3360009081526008602052604090208054600654348201928390556001600160801b03161180801561096f57506006546001600160801b03168210155b1561099957600d8054600161ffff62010000808404821692909201160263ffff0000199091161790555b600d5462010000810461ffff90811691161480156109ce57506000600754600160401b900460ff1660038111156109cc57fe5b145b156109e7576007805460ff60401b1916600160401b1790555b5050565b6001600160a01b0381166000908152600860205260409020545b919050565b600d5461ffff1681565b600754600160401b900460ff1690565b6000546001600160a01b03163314610a6d5760405162461bcd60e51b815260040180806020018281038252602d815260200180611f8d602d913960400191505060405180910390fd5b600754600160481b900460ff16610ab55760405162461bcd60e51b8152600401808060200182810382526026815260200180611f676026913960400191505060405180910390fd5b6007805469ff00000000000000000019169055610b1a6001600160801b038235166008600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000205461179990919063ffffffff16565b82356001600160a01b03166000908152600860208181526040832093909355610b52928401356001600160801b031691856001610add565b6001600160a01b03602093840135166000908152600890935260409092209190915550565b336000908152600860205260409020546006546001600160801b03161115610bd05760405162461bcd60e51b8152600401808060200182810382526027815260200180611f406027913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151632d7c9e3d60e11b81526002600482015260248101849052915173__$2104f4b4ea1fa2fd2334e6605946f6eea1$__92635af93c7a926044808301939192829003018186803b158015610c4a57600080fd5b505af4158015610c5e573d6000803e3d6000fd5b5050600080546040805180820182526005546001600160a01b03908116825233602080840191909152835180850185526006546001600160801b0316808252918101919091526007548451630823813560e21b815292909516975063208e04d496509194919363ffffffff16928892600490920191829187918190849084905b83811015610cf6578181015183820152602001610cde565b5050505090500184600260200280838360005b83811015610d21578181015183820152602001610d09565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050600060405180830381600087803b158015610d6057600080fd5b505af1158015610d74573d6000803e3d6000fd5b5050505050565b6002546003546004546005546006546007546001600160a01b03909216916001600160801b0382169167ffffffffffffffff600160801b8204811692600160c01b909204169063ffffffff80821691600160201b81049091169060ff600160401b8204811691600160481b9004168b565b600d54600090829061ffff168114610e08576000915050610e6d565b60005b600d5461ffff16811015610e6657600c6000868684818110610e2957fe5b602090810292909201356001600160a01b03168352508101919091526040016000205460ff16610e5e57600092505050610e6d565b600101610e0b565b5060019150505b92915050565b600a546001600160a01b031681565b6009546001600160a01b031681565b73__$caf066876633ea418098495f1e5bb4c2f5$__635ee899da60023088888888886040518863ffffffff1660e01b815260040180888152602001876001600160a01b03166001600160a01b031681526020018681526020018581526020018060200184815260200180602001838103835286818151815260200191508051906020019080838360005b83811015610f33578181015183820152602001610f1b565b50505050905090810190601f168015610f605780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610f93578181015183820152602001610f7b565b50505050905090810190601f168015610fc05780820380516001836020036101000a031916815260200191505b50995050505050505050505060006040518083038186803b158015610fe457600080fd5b505af4158015610ff8573d6000803e3d6000fd5b5060029250611005915050565b600754600160401b900460ff16600381111561101d57fe5b1415611075576006546005546001600160a01b0316600090815260086020526040902054611059916001600160801b031663ffffffff61179916565b6005546001600160a01b03166000908152600860205260409020555b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63eb49982c6002876040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156110cd57600080fd5b505af41580156110e1573d6000803e3d6000fd5b50505050610d74836117fa565b600d5462010000900461ffff1681565b600b546001600160a01b031681565b600b546001600160a01b03163314611165576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6000600754600160401b900460ff16600381111561117f57fe5b1415611199576007805460ff60401b1916600160401b1790555b565b6006546001600160801b031690565b6001600160a01b03166000908152600c602052604090205460ff1690565b600b546001600160a01b03163314611220576040805162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015290519081900360640190fd5b6111996119f8565b6001546001600160a01b031681565b73__$caf066876633ea418098495f1e5bb4c2f5$__63b4d866a260023088888888886040518863ffffffff1660e01b815260040180888152602001876001600160a01b03166001600160a01b031681526020018681526020018567ffffffffffffffff1667ffffffffffffffff16815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156112ef5781810151838201526020016112d7565b50505050905090810190601f16801561131c5780820380516001836020036101000a031916815260200191505b509850505050505050505060006040518083038186803b15801561133f57600080fd5b505af4158015611353573d6000803e3d6000fd5b5060029250611360915050565b600754600160401b900460ff16600381111561137857fe5b14156113d0576006546005546001600160a01b03166000908152600860205260409020546113b4916001600160801b031663ffffffff61179916565b6005546001600160a01b03166000908152600860205260409020555b50506007805460ff60401b191668030000000000000000179055505050565b73__$caf066876633ea418098495f1e5bb4c2f5$__63e2d5c52f60028585856040518563ffffffff1660e01b81526004018085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561146c578181015183820152602001611454565b50505050905090810190601f1680156114995780820380516001836020036101000a031916815260200191505b509550505050505060006040518083038186803b1580156114b957600080fd5b505af41580156114cd573d6000803e3d6000fd5b505050506114da816117fa565b505050565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63ce9d5122600287878787876040518763ffffffff1660e01b8152600401808781526020018681526020018581526020018463ffffffff1663ffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561157657818101518382015260200161155e565b50505050905090810190601f1680156115a35780820380516001836020036101000a031916815260200191505b5097505050505050505060006040518083038186803b1580156115c557600080fd5b505af41580156115d9573d6000803e3d6000fd5b50506006546005546001600160a01b0316600090815260086020526040902054611614935091506001600160801b031663ffffffff61179916565b6005546001600160a01b0316600090815260086020526040902055610d74826117fa565b336000908152600860205260409020546006546001600160801b031611156116915760405162461bcd60e51b8152600401808060200182810382526031815260200180611fba6031913960400191505060405180910390fd5b6006543360009081526008602052604080822080546001600160801b039094169093039092558151636c36f28f60e11b8152600260048201818152602483018c9052604483018b9052606483018a90526084830189905260a4830188905263ffffffff871660c484015273__$2104f4b4ea1fa2fd2334e6605946f6eea1$__9463d86de51e9492938d938d938d938d938d938d938d93909260e401918491908190849084905b8381101561174f578181015183820152602001611737565b505050509050019850505050505050505060006040518083038186803b15801561177857600080fd5b505af415801561178c573d6000803e3d6000fd5b5050505050505050505050565b6000828201838110156117f3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001546040805163d106ec1960e01b815290516000926001600160a01b03169163d106ec1991600480830192602092919082900301818787803b15801561184057600080fd5b505af1158015611854573d6000803e3d6000fd5b505050506040513d602081101561186a57600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b1580156118bb57600080fd5b505af41580156118cf573d6000803e3d6000fd5b505050506040513d60208110156118e557600080fd5b505181146119305761192c60405180606001604052806119056001611a06565b81526020016119176002800154611a84565b815260200161192584611a84565b9052611b02565b6004555b60015460405163e4eb8c6360e01b81526020600482018181528551602484015285516001600160a01b039094169363e4eb8c6393879383926044909201919085019080838360005b83811015611990578181015183820152602001611978565b50505050905090810190601f1680156119bd5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156119dc57600080fd5b505af11580156119f0573d6000803e3d6000fd5b505050505050565b600b546001600160a01b0316ff5b611a0e611ed8565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611a73565b611a60611ed8565b815260200190600190039081611a585790505b508152600060209091015292915050565b611a8c611ed8565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611af1565b611ade611ed8565b815260200190600190039081611ad65790505b508152600260209091015292915050565b6040805160038082526080820190925260009160609190816020015b611b26611ed8565b815260200190600190039081611b1e575050805190915060005b81811015611b7857848160038110611b5457fe5b6020020151838281518110611b6557fe5b6020908102919091010152600101611b40565b50611b8282611b8a565b949350505050565b6000600882511115611bda576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611c07578160200160208202803883390190505b50805190915060005b81811015611c6357611c20611f06565b611c3c868381518110611c2f57fe5b6020026020010151611cd6565b90508060000151848381518110611c4f57fe5b602090810291909101015250600101611c10565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611cac578181015183820152602001611c94565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b611cde611f06565b6060820151600c60ff90911610611d30576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16611d5d576040518060200160405280611d548460000151611e0c565b90529050610a05565b606082015160ff1660011415611da4576040518060200160405280611d54846020015160000151856020015160400151866020015160600151876020015160200151611e30565b606082015160ff1660021415611dc95750604080516020810190915281518152610a05565b600360ff16826060015160ff1610158015611ded57506060820151600c60ff909116105b15611e0a576040518060200160405280611d548460400151611b8a565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b60008315611e8a575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611b82565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b604051806080016040528060008152602001611ef2611f18565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e676556616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f7720746f20617373657274a265627a7a72315820a3ecd7ee59e443a702a2aeed38ed6415902bc44c3b697088416369159973f5fc64736f6c634300050c0032a265627a7a723158202176f3215535d4b13840c303ddb2d8cb6731f60a8a992e68b63dca3cda4bf76964736f6c634300050c0032"

// DeployChannelLauncher deploys a new Ethereum contract, binding an instance of ChannelLauncher to it.
func DeployChannelLauncher(auth *bind.TransactOpts, backend bind.ContractBackend, _globalInboxAddress common.Address, _challengeManagerAddress common.Address) (common.Address, *types.Transaction, *ChannelLauncher, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelLauncherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	ChannelLauncherBin = strings.Replace(ChannelLauncherBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	vMAddr, _, _, _ := DeployVM(auth, backend)
	ChannelLauncherBin = strings.Replace(ChannelLauncherBin, "__$8e266570c8a7fb2aaac83b3e040afaf9e1$__", vMAddr.String()[2:], -1)

	unanimousAddr, _, _, _ := DeployUnanimous(auth, backend)
	ChannelLauncherBin = strings.Replace(ChannelLauncherBin, "__$caf066876633ea418098495f1e5bb4c2f5$__", unanimousAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ChannelLauncherBin = strings.Replace(ChannelLauncherBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelLauncherBin), backend, _globalInboxAddress, _challengeManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChannelLauncher{ChannelLauncherCaller: ChannelLauncherCaller{contract: contract}, ChannelLauncherTransactor: ChannelLauncherTransactor{contract: contract}, ChannelLauncherFilterer: ChannelLauncherFilterer{contract: contract}}, nil
}

// ChannelLauncher is an auto generated Go binding around an Ethereum contract.
type ChannelLauncher struct {
	ChannelLauncherCaller     // Read-only binding to the contract
	ChannelLauncherTransactor // Write-only binding to the contract
	ChannelLauncherFilterer   // Log filterer for contract events
}

// ChannelLauncherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelLauncherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelLauncherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelLauncherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelLauncherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelLauncherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelLauncherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelLauncherSession struct {
	Contract     *ChannelLauncher  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChannelLauncherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelLauncherCallerSession struct {
	Contract *ChannelLauncherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ChannelLauncherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelLauncherTransactorSession struct {
	Contract     *ChannelLauncherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ChannelLauncherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelLauncherRaw struct {
	Contract *ChannelLauncher // Generic contract binding to access the raw methods on
}

// ChannelLauncherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelLauncherCallerRaw struct {
	Contract *ChannelLauncherCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelLauncherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelLauncherTransactorRaw struct {
	Contract *ChannelLauncherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannelLauncher creates a new instance of ChannelLauncher, bound to a specific deployed contract.
func NewChannelLauncher(address common.Address, backend bind.ContractBackend) (*ChannelLauncher, error) {
	contract, err := bindChannelLauncher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChannelLauncher{ChannelLauncherCaller: ChannelLauncherCaller{contract: contract}, ChannelLauncherTransactor: ChannelLauncherTransactor{contract: contract}, ChannelLauncherFilterer: ChannelLauncherFilterer{contract: contract}}, nil
}

// NewChannelLauncherCaller creates a new read-only instance of ChannelLauncher, bound to a specific deployed contract.
func NewChannelLauncherCaller(address common.Address, caller bind.ContractCaller) (*ChannelLauncherCaller, error) {
	contract, err := bindChannelLauncher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelLauncherCaller{contract: contract}, nil
}

// NewChannelLauncherTransactor creates a new write-only instance of ChannelLauncher, bound to a specific deployed contract.
func NewChannelLauncherTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelLauncherTransactor, error) {
	contract, err := bindChannelLauncher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelLauncherTransactor{contract: contract}, nil
}

// NewChannelLauncherFilterer creates a new log filterer instance of ChannelLauncher, bound to a specific deployed contract.
func NewChannelLauncherFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelLauncherFilterer, error) {
	contract, err := bindChannelLauncher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelLauncherFilterer{contract: contract}, nil
}

// bindChannelLauncher binds a generic wrapper to an already deployed contract.
func bindChannelLauncher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelLauncherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelLauncher *ChannelLauncherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelLauncher.Contract.ChannelLauncherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelLauncher *ChannelLauncherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelLauncher.Contract.ChannelLauncherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelLauncher *ChannelLauncherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelLauncher.Contract.ChannelLauncherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelLauncher *ChannelLauncherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelLauncher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelLauncher *ChannelLauncherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelLauncher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelLauncher *ChannelLauncherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelLauncher.Contract.contract.Transact(opts, method, params...)
}

// LaunchChannel is a paid mutator transaction binding the contract method 0xad7d0fa4.
//
// Solidity: function launchChannel(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address[] _validatorKeys) returns()
func (_ChannelLauncher *ChannelLauncherTransactor) LaunchChannel(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _ChannelLauncher.contract.Transact(opts, "launchChannel", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _validatorKeys)
}

// LaunchChannel is a paid mutator transaction binding the contract method 0xad7d0fa4.
//
// Solidity: function launchChannel(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address[] _validatorKeys) returns()
func (_ChannelLauncher *ChannelLauncherSession) LaunchChannel(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _ChannelLauncher.Contract.LaunchChannel(&_ChannelLauncher.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _validatorKeys)
}

// LaunchChannel is a paid mutator transaction binding the contract method 0xad7d0fa4.
//
// Solidity: function launchChannel(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _owner, address[] _validatorKeys) returns()
func (_ChannelLauncher *ChannelLauncherTransactorSession) LaunchChannel(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _ChannelLauncher.Contract.LaunchChannel(&_ChannelLauncher.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _owner, _validatorKeys)
}

// ChannelLauncherChannelCreatedIterator is returned from FilterChannelCreated and is used to iterate over the raw logs and unpacked data for ChannelCreated events raised by the ChannelLauncher contract.
type ChannelLauncherChannelCreatedIterator struct {
	Event *ChannelLauncherChannelCreated // Event containing the contract specifics and raw log

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
func (it *ChannelLauncherChannelCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelLauncherChannelCreated)
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
		it.Event = new(ChannelLauncherChannelCreated)
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
func (it *ChannelLauncherChannelCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelLauncherChannelCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelLauncherChannelCreated represents a ChannelCreated event raised by the ChannelLauncher contract.
type ChannelLauncherChannelCreated struct {
	VmAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelCreated is a free log retrieval operation binding the contract event 0xc625d37dd8b556110d70984e62f74ba35c77422c83c5f548fbd21b697a67ef5c.
//
// Solidity: event ChannelCreated(address vmAddress)
func (_ChannelLauncher *ChannelLauncherFilterer) FilterChannelCreated(opts *bind.FilterOpts) (*ChannelLauncherChannelCreatedIterator, error) {

	logs, sub, err := _ChannelLauncher.contract.FilterLogs(opts, "ChannelCreated")
	if err != nil {
		return nil, err
	}
	return &ChannelLauncherChannelCreatedIterator{contract: _ChannelLauncher.contract, event: "ChannelCreated", logs: logs, sub: sub}, nil
}

// WatchChannelCreated is a free log subscription operation binding the contract event 0xc625d37dd8b556110d70984e62f74ba35c77422c83c5f548fbd21b697a67ef5c.
//
// Solidity: event ChannelCreated(address vmAddress)
func (_ChannelLauncher *ChannelLauncherFilterer) WatchChannelCreated(opts *bind.WatchOpts, sink chan<- *ChannelLauncherChannelCreated) (event.Subscription, error) {

	logs, sub, err := _ChannelLauncher.contract.WatchLogs(opts, "ChannelCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelLauncherChannelCreated)
				if err := _ChannelLauncher.contract.UnpackLog(event, "ChannelCreated", log); err != nil {
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

// ParseChannelCreated is a log parse operation binding the contract event 0xc625d37dd8b556110d70984e62f74ba35c77422c83c5f548fbd21b697a67ef5c.
//
// Solidity: event ChannelCreated(address vmAddress)
func (_ChannelLauncher *ChannelLauncherFilterer) ParseChannelCreated(log types.Log) (*ChannelLauncherChannelCreated, error) {
	event := new(ChannelLauncherChannelCreated)
	if err := _ChannelLauncher.contract.UnpackLog(event, "ChannelCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DisputableABI is the input ABI used to generate the binding from.
const DisputableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PendingAssertionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"withinTimeBounds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DisputableFuncSigs maps the 4-byte function signature to its string representation.
var DisputableFuncSigs = map[string]string{
	"ce9d5122": "confirmDisputableAsserted(VM.Data storage,bytes32,bytes32,uint32,bytes,bytes32)",
	"5af93c7a": "initiateChallenge(VM.Data storage,bytes32)",
	"d86de51e": "pendingDisputableAssert(VM.Data storage,bytes32,bytes32,bytes32,bytes32,bytes32,uint32,uint64[2])",
	"42c0787e": "withinTimeBounds(uint64[2])",
}

// DisputableBin is the compiled bytecode used for deploying new contracts.
var DisputableBin = "0x6110e6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806342c0787e1461005b5780635af93c7a146100ba578063ce9d5122146100ec578063d86de51e146101bf575b600080fd5b6100a66004803603604081101561007157600080fd5b604080518082018252918301929181830191839060029083908390808284376000920191909152509194506102489350505050565b604080519115158252519081900360200190f35b8180156100c657600080fd5b506100ea600480360360408110156100dd57600080fd5b508035906020013561027a565b005b8180156100f857600080fd5b506100ea600480360360c081101561010f57600080fd5b81359160208101359160408201359163ffffffff6060820135169181019060a08101608082013564010000000081111561014857600080fd5b82018360208201111561015a57600080fd5b8035906020019184600183028401116401000000008311171561017c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610472915050565b8180156101cb57600080fd5b506100ea60048036036101208110156101e357600080fd5b604080518082018252833593602081013593838201359360608301359360808401359360a08101359363ffffffff60c083013516939082019261012083019160e0840190600290839083908082843760009201919091525091945061082b9350505050565b805160009067ffffffffffffffff1643108015906102745750602082015167ffffffffffffffff164311155b92915050565b60038201546001600160a01b03163314156102c65760405162461bcd60e51b8152600401808060200182810382526021815260200180610f1d6021913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561031557600080fd5b505af4158015610329573d6000803e3d6000fd5b505050506040513d602081101561033f57600080fd5b505161037c5760405162461bcd60e51b8152600401808060200182810382526026815260200180610fc56026913960400191505060405180910390fd5b60026005830154600160401b900460ff16600381111561039857fe5b146103d45760405162461bcd60e51b815260040180806020018281038252602f815260200180610e56602f913960400191505060405180910390fd5b816001015481146104165760405162461bcd60e51b815260040180806020018281038252604d815260200180611018604d913960600191505060405180910390fd5b60006001830181905560058301805460ff60401b1916600160401b1769ff0000000000000000001916600160481b1790556040517f198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d909190a15050565b60026005870154600160401b900460ff16600381111561048e57fe5b146104ca5760405162461bcd60e51b8152600401808060200182810382526022815260200180610ed76022913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5876040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561051957600080fd5b505af415801561052d573d6000803e3d6000fd5b505050506040513d602081101561054357600080fd5b5051156105815760405162461bcd60e51b8152600401808060200182810382526024815260200180610eb36024913960400191505060405180910390fd5b85600101548573__$9836fa7140e5a33041d4b827682e675a30$__637ddf59d68787600073__$9836fa7140e5a33041d4b827682e675a30$__63e83f4bfe8a6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561060c5781810151838201526020016105f4565b50505050905090810190601f1680156106395780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b15801561065657600080fd5b505af415801561066a573d6000803e3d6000fd5b505050506040513d602081101561068057600080fd5b5051604080516001600160e01b031960e088901b168152600481019590955263ffffffff939093166024850152604484019190915260648301526000608483015260a482018790525160c4808301926020929190829003018186803b1580156106e857600080fd5b505af41580156106fc573d6000803e3d6000fd5b505050506040513d602081101561071257600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920190528051910120146107795760405162461bcd60e51b815260040180806020018281038252604d815260200180611065604d913960600191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63eb49982c87866040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156107d057600080fd5b505af41580156107e4573d6000803e3d6000fd5b5050604080518781526020810185905281517f4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb9450908190039091019150a1505050505050565b60016005890154600160401b900460ff16600381111561084757fe5b146108835760405162461bcd60e51b815260040180806020018281038252602d815260200180610feb602d913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__632a3e0a97896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156108d257600080fd5b505af41580156108e6573d6000803e3d6000fd5b505050506040513d60208110156108fc57600080fd5b5051158015610983575073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561095557600080fd5b505af4158015610969573d6000803e3d6000fd5b505050506040513d602081101561097f57600080fd5b5051155b6109be5760405162461bcd60e51b815260040180806020018281038252603e815260200180610f3e603e913960400191505060405180910390fd5b6005880154600160481b900460ff1615610a095760405162461bcd60e51b815260040180806020018281038252602e815260200180610e85602e913960400191505060405180910390fd5b600588015463ffffffff64010000000090910481169083161115610a74576040805162461bcd60e51b815260206004820152601f60248201527f547269656420746f206578656375746520746f6f206d616e7920737465707300604482015290519081900360640190fd5b610a7d81610248565b610ab85760405162461bcd60e51b8152600401808060200182810382526024815260200180610ef96024913960400191505060405180910390fd5b87548714610af75760405162461bcd60e51b8152600401808060200182810382526027815260200180610f9e6027913960400191505060405180910390fd5b87600201548614610b395760405162461bcd60e51b8152600401808060200182810382526022815260200180610f7c6022913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63a3a162cb896040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610b8857600080fd5b505af4158015610b9c573d6000803e3d6000fd5b5050505073__$9836fa7140e5a33041d4b827682e675a30$__6385ecb92a8883896040518463ffffffff1660e01b81526004018084815260200183600260200280838360005b83811015610bfa578181015183820152602001610be2565b50505050905001828152602001935050505060206040518083038186803b158015610c2457600080fd5b505af4158015610c38573d6000803e3d6000fd5b505050506040513d6020811015610c4e57600080fd5b505160408051633eefaceb60e11b81526004810188905263ffffffff8516602482015260006044820181905260648201889052608482015260a48101869052905173__$9836fa7140e5a33041d4b827682e675a30$__91637ddf59d69160c4808301926020929190829003018186803b158015610cca57600080fd5b505af4158015610cde573d6000803e3d6000fd5b505050506040513d6020811015610cf457600080fd5b5051604080516020818101949094528082019290925280518083038201815260609092019052805191012060018901556003880180546001600160a01b031916331790556005880180546002919060ff60401b1916600160401b8302179055507f247e6305d02be2139d3707f34270f5c1e02b6a87abcec6cd099e800dcbe244366040518060a00160405280898152602001888152602001878152602001868152602001858152503383858c60040160109054906101000a900467ffffffffffffffff166040518086600560200280838360005b83811015610de0578181015183820152602001610dc8565b505050506001600160a01b03881692019182525060200184604080838360005b83811015610e18578181015183820152602001610e00565b50505063ffffffff90961692019182525067ffffffffffffffff909216602083015250604080519182900301945092505050a1505050505050505056fe417373657274696f6e206d7573742062652070656e64696e6720746f20696e697469617465206368616c6c656e676543616e206f6e6c792064697370757461626c6520617373657274206966206e6f7420696e206368616c6c656e6765417373657274696f6e206973207374696c6c2070656e64696e67206368616c6c656e6765564d20646f6573206e6f74206861766520617373657274696f6e2070656e64696e67507265636f6e646974696f6e3a206e6f742077697468696e2074696d6520626f756e64734368616c6c656e676520776173206372656174656420627920617373657274657243616e206f6e6c792064697370757461626c6520617373657274206966206d616368696e65206973206e6f74206572726f726564206f722068616c746564507265636f6e646974696f6e3a20696e626f7820646f6573206e6f74206d61746368507265636f6e646974696f6e3a207374617465206861736820646f6573206e6f74206d617463684368616c6c656e676520646964206e6f7420636f6d65206265666f726520646561646c696e6543616e206f6e6c792064697370757461626c65206173736572742066726f6d2077616974696e67207374617465496e697469617465204368616c6c656e67653a20507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6e436f6e6669726d2044697370757461626c653a20507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6ea265627a7a723158204ef267b56b045bb835ac5fca910f2b99cbd55ad948e52cc42ebf56e7edbdc48f64736f6c634300050c0032"

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

// DisputablePendingAssertionCanceledIterator is returned from FilterPendingAssertionCanceled and is used to iterate over the raw logs and unpacked data for PendingAssertionCanceled events raised by the Disputable contract.
type DisputablePendingAssertionCanceledIterator struct {
	Event *DisputablePendingAssertionCanceled // Event containing the contract specifics and raw log

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
func (it *DisputablePendingAssertionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputablePendingAssertionCanceled)
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
		it.Event = new(DisputablePendingAssertionCanceled)
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
func (it *DisputablePendingAssertionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputablePendingAssertionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputablePendingAssertionCanceled represents a PendingAssertionCanceled event raised by the Disputable contract.
type DisputablePendingAssertionCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPendingAssertionCanceled is a free log retrieval operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_Disputable *DisputableFilterer) FilterPendingAssertionCanceled(opts *bind.FilterOpts) (*DisputablePendingAssertionCanceledIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "PendingAssertionCanceled")
	if err != nil {
		return nil, err
	}
	return &DisputablePendingAssertionCanceledIterator{contract: _Disputable.contract, event: "PendingAssertionCanceled", logs: logs, sub: sub}, nil
}

// WatchPendingAssertionCanceled is a free log subscription operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_Disputable *DisputableFilterer) WatchPendingAssertionCanceled(opts *bind.WatchOpts, sink chan<- *DisputablePendingAssertionCanceled) (event.Subscription, error) {

	logs, sub, err := _Disputable.contract.WatchLogs(opts, "PendingAssertionCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputablePendingAssertionCanceled)
				if err := _Disputable.contract.UnpackLog(event, "PendingAssertionCanceled", log); err != nil {
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

// ParsePendingAssertionCanceled is a log parse operation binding the contract event 0x198a6a11b54668b4b20f82518e7c15b46a61eec7893f4f4851a295a9f76f3d90.
//
// Solidity: event PendingAssertionCanceled()
func (_Disputable *DisputableFilterer) ParsePendingAssertionCanceled(log types.Log) (*DisputablePendingAssertionCanceled, error) {
	event := new(DisputablePendingAssertionCanceled)
	if err := _Disputable.contract.UnpackLog(event, "PendingAssertionCanceled", log); err != nil {
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

// IArbChannelABI is the input ABI used to generate the binding from.
const IArbChannelABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"isValidatorList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IArbChannelFuncSigs maps the 4-byte function signature to its string representation.
var IArbChannelFuncSigs = map[string]string{
	"513164fe": "isValidatorList(address[])",
}

// IArbChannel is an auto generated Go binding around an Ethereum contract.
type IArbChannel struct {
	IArbChannelCaller     // Read-only binding to the contract
	IArbChannelTransactor // Write-only binding to the contract
	IArbChannelFilterer   // Log filterer for contract events
}

// IArbChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbChannelSession struct {
	Contract     *IArbChannel      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbChannelCallerSession struct {
	Contract *IArbChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IArbChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbChannelTransactorSession struct {
	Contract     *IArbChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IArbChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbChannelRaw struct {
	Contract *IArbChannel // Generic contract binding to access the raw methods on
}

// IArbChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbChannelCallerRaw struct {
	Contract *IArbChannelCaller // Generic read-only contract binding to access the raw methods on
}

// IArbChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbChannelTransactorRaw struct {
	Contract *IArbChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbChannel creates a new instance of IArbChannel, bound to a specific deployed contract.
func NewIArbChannel(address common.Address, backend bind.ContractBackend) (*IArbChannel, error) {
	contract, err := bindIArbChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbChannel{IArbChannelCaller: IArbChannelCaller{contract: contract}, IArbChannelTransactor: IArbChannelTransactor{contract: contract}, IArbChannelFilterer: IArbChannelFilterer{contract: contract}}, nil
}

// NewIArbChannelCaller creates a new read-only instance of IArbChannel, bound to a specific deployed contract.
func NewIArbChannelCaller(address common.Address, caller bind.ContractCaller) (*IArbChannelCaller, error) {
	contract, err := bindIArbChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChannelCaller{contract: contract}, nil
}

// NewIArbChannelTransactor creates a new write-only instance of IArbChannel, bound to a specific deployed contract.
func NewIArbChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbChannelTransactor, error) {
	contract, err := bindIArbChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChannelTransactor{contract: contract}, nil
}

// NewIArbChannelFilterer creates a new log filterer instance of IArbChannel, bound to a specific deployed contract.
func NewIArbChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbChannelFilterer, error) {
	contract, err := bindIArbChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbChannelFilterer{contract: contract}, nil
}

// bindIArbChannel binds a generic wrapper to an already deployed contract.
func bindIArbChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChannel *IArbChannelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChannel.Contract.IArbChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChannel *IArbChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChannel.Contract.IArbChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChannel *IArbChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChannel.Contract.IArbChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChannel *IArbChannelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChannel *IArbChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChannel *IArbChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChannel.Contract.contract.Transact(opts, method, params...)
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_IArbChannel *IArbChannelCaller) IsValidatorList(opts *bind.CallOpts, _validators []common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IArbChannel.contract.Call(opts, out, "isValidatorList", _validators)
	return *ret0, err
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_IArbChannel *IArbChannelSession) IsValidatorList(_validators []common.Address) (bool, error) {
	return _IArbChannel.Contract.IsValidatorList(&_IArbChannel.CallOpts, _validators)
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_IArbChannel *IArbChannelCallerSession) IsValidatorList(_validators []common.Address) (bool, error) {
	return _IArbChannel.Contract.IsValidatorList(&_IArbChannel.CallOpts, _validators)
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

// SigUtilsABI is the input ABI used to generate the binding from.
const SigUtilsABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"countSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"parseSignature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"recoverAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SigUtilsFuncSigs maps the 4-byte function signature to its string representation.
var SigUtilsFuncSigs = map[string]string{
	"33ae3ad0": "countSignatures(bytes)",
	"b31d63cc": "parseSignature(bytes,uint256)",
	"c655d7aa": "recoverAddress(bytes32,bytes)",
	"f0c8e969": "recoverAddresses(bytes32,bytes)",
}

// SigUtilsBin is the compiled bytecode used for deploying new contracts.
var SigUtilsBin = "0x610764610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806333ae3ad01461005b578063b31d63cc14610111578063c655d7aa146101d9578063f0c8e969146102a0575b600080fd5b6100ff6004803603602081101561007157600080fd5b810190602081018135600160201b81111561008b57600080fd5b82018360208201111561009d57600080fd5b803590602001918460018302840111600160201b831117156100be57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061039b945050505050565b60408051918252519081900360200190f35b6101b76004803603604081101561012757600080fd5b810190602081018135600160201b81111561014157600080fd5b82018360208201111561015357600080fd5b803590602001918460018302840111600160201b8311171561017457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506103c8915050565b6040805160ff9094168452602084019290925282820152519081900360600190f35b610284600480360360408110156101ef57600080fd5b81359190810190604081016020820135600160201b81111561021057600080fd5b82018360208201111561022257600080fd5b803590602001918460018302840111600160201b8311171561024357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610456945050505050565b604080516001600160a01b039092168252519081900360200190f35b61034b600480360360408110156102b657600080fd5b81359190810190604081016020820135600160201b8111156102d757600080fd5b8201836020820111156102e957600080fd5b803590602001918460018302840111600160201b8311171561030a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610589945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561038757818101518382015260200161036f565b505050509050019250505060405180910390f35b600060418251816103a857fe5b06156103b55760006103c2565b60418251816103c057fe5b045b92915050565b604180820283810160208101516040820151919093015160ff169291601b8410156103f457601b840193505b8360ff16601b148061040957508360ff16601c145b61044e576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081886040516020018083805190602001908083835b602083106104cc5780518252601f1990920191602091820191016104ad565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201905282519201919091209250610513915088905060006103c8565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa158015610572573d6000803e3d6000fd5b5050604051601f1901519998505050505050505050565b606060008060008061059a8661039b565b90506060816040519080825280602002602001820160405280156105c8578160200160208202803883390190505b50905060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818a6040516020018083805190602001908083835b6020831061063b5780518252601f19909201916020918201910161061c565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190528251920191909120925060009150505b848110156107205761068b8a826103c8565b6040805160008152602080820180845288905260ff86168284015260608201859052608082018490529151949c50929a5090985060019260a080840193601f198301929081900390910190855afa1580156106ea573d6000803e3d6000fd5b5050506020604051035184828151811061070057fe5b6001600160a01b0390921660209283029190910190910152600101610679565b5091999850505050505050505056fea265627a7a72315820ca13c449a6c354a85861b439d7be32702fa4df9dddb4817404612a1a369b097b64736f6c634300050c0032"

// DeploySigUtils deploys a new Ethereum contract, binding an instance of SigUtils to it.
func DeploySigUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// SigUtils is an auto generated Go binding around an Ethereum contract.
type SigUtils struct {
	SigUtilsCaller     // Read-only binding to the contract
	SigUtilsTransactor // Write-only binding to the contract
	SigUtilsFilterer   // Log filterer for contract events
}

// SigUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsSession struct {
	Contract     *SigUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsCallerSession struct {
	Contract *SigUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTransactorSession struct {
	Contract     *SigUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsRaw struct {
	Contract *SigUtils // Generic contract binding to access the raw methods on
}

// SigUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsCallerRaw struct {
	Contract *SigUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTransactorRaw struct {
	Contract *SigUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtils creates a new instance of SigUtils, bound to a specific deployed contract.
func NewSigUtils(address common.Address, backend bind.ContractBackend) (*SigUtils, error) {
	contract, err := bindSigUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// NewSigUtilsCaller creates a new read-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsCaller, error) {
	contract, err := bindSigUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsCaller{contract: contract}, nil
}

// NewSigUtilsTransactor creates a new write-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTransactor, error) {
	contract, err := bindSigUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTransactor{contract: contract}, nil
}

// NewSigUtilsFilterer creates a new log filterer instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsFilterer, error) {
	contract, err := bindSigUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsFilterer{contract: contract}, nil
}

// bindSigUtils binds a generic wrapper to an already deployed contract.
func bindSigUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.SigUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transact(opts, method, params...)
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_SigUtils *SigUtilsCaller) CountSignatures(opts *bind.CallOpts, _signatures []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SigUtils.contract.Call(opts, out, "countSignatures", _signatures)
	return *ret0, err
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_SigUtils *SigUtilsSession) CountSignatures(_signatures []byte) (*big.Int, error) {
	return _SigUtils.Contract.CountSignatures(&_SigUtils.CallOpts, _signatures)
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_SigUtils *SigUtilsCallerSession) CountSignatures(_signatures []byte) (*big.Int, error) {
	return _SigUtils.Contract.CountSignatures(&_SigUtils.CallOpts, _signatures)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtils *SigUtilsCaller) ParseSignature(opts *bind.CallOpts, _signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	ret := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	out := ret
	err := _SigUtils.contract.Call(opts, out, "parseSignature", _signatures, _pos)
	return *ret, err
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtils *SigUtilsSession) ParseSignature(_signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigUtils.Contract.ParseSignature(&_SigUtils.CallOpts, _signatures, _pos)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtils *SigUtilsCallerSession) ParseSignature(_signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigUtils.Contract.ParseSignature(&_SigUtils.CallOpts, _signatures, _pos)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_SigUtils *SigUtilsCaller) RecoverAddress(opts *bind.CallOpts, _messageHash [32]byte, _signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SigUtils.contract.Call(opts, out, "recoverAddress", _messageHash, _signature)
	return *ret0, err
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_SigUtils *SigUtilsSession) RecoverAddress(_messageHash [32]byte, _signature []byte) (common.Address, error) {
	return _SigUtils.Contract.RecoverAddress(&_SigUtils.CallOpts, _messageHash, _signature)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_SigUtils *SigUtilsCallerSession) RecoverAddress(_messageHash [32]byte, _signature []byte) (common.Address, error) {
	return _SigUtils.Contract.RecoverAddress(&_SigUtils.CallOpts, _messageHash, _signature)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_SigUtils *SigUtilsCaller) RecoverAddresses(opts *bind.CallOpts, _messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _SigUtils.contract.Call(opts, out, "recoverAddresses", _messageHash, _signatures)
	return *ret0, err
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_SigUtils *SigUtilsSession) RecoverAddresses(_messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	return _SigUtils.Contract.RecoverAddresses(&_SigUtils.CallOpts, _messageHash, _signatures)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_SigUtils *SigUtilsCallerSession) RecoverAddresses(_messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	return _SigUtils.Contract.RecoverAddresses(&_SigUtils.CallOpts, _messageHash, _signatures)
}

// UnanimousABI is the input ABI used to generate the binding from.
const UnanimousABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"ConfirmedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unanHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unanHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"PendingUnanimousAssertion\",\"type\":\"event\"}]"

// UnanimousFuncSigs maps the 4-byte function signature to its string representation.
var UnanimousFuncSigs = map[string]string{
	"e2d5c52f": "confirmUnanimousAsserted(VM.Data storage,bytes32,bytes32,bytes)",
	"5ee899da": "finalizedUnanimousAssert(VM.Data storage,IArbChannel,bytes32,bytes32,bytes,bytes32,bytes)",
	"b4d866a2": "pendingUnanimousAssert(VM.Data storage,IArbChannel,bytes32,uint64,bytes32,bytes32,bytes)",
}

// UnanimousBin is the compiled bytecode used for deploying new contracts.
var UnanimousBin = "0x6112e5610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80635ee899da14610050578063b4d866a2146101b1578063e2d5c52f14610298575b600080fd5b81801561005c57600080fd5b506101af600480360360e081101561007357600080fd5b8135916001600160a01b036020820135169160408201359160608101359181019060a081016080820135600160201b8111156100ae57600080fd5b8201836020820111156100c057600080fd5b803590602001918460018302840111600160201b831117156100e157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b81111561013b57600080fd5b82018360208201111561014d57600080fd5b803590602001918460018302840111600160201b8311171561016e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061035c945050505050565b005b8180156101bd57600080fd5b506101af600480360360e08110156101d457600080fd5b8135916001600160a01b03602082013516916040820135916001600160401b036060820135169160808201359160a08101359181019060e0810160c0820135600160201b81111561022457600080fd5b82018360208201111561023657600080fd5b803590602001918460018302840111600160201b8311171561025757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610706945050505050565b8180156102a457600080fd5b506101af600480360360808110156102bb57600080fd5b81359160208101359160408201359190810190608081016060820135600160201b8111156102e857600080fd5b8201836020820111156102fa57600080fd5b803590602001918460018302840111600160201b8311171561031b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b2b945050505050565b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca886040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156103ab57600080fd5b505af41580156103bf573d6000803e3d6000fd5b505050506040513d60208110156103d557600080fd5b505115610429576040805162461bcd60e51b815260206004820152601b60248201527f43616e2774206173736572742068616c746564206d616368696e650000000000604482015290519081900360640190fd5b60016005880154600160401b900460ff16600381111561044557fe5b148061046a575060026005880154600160401b900460ff16600381111561046857fe5b145b8061048e575060036005880154600160401b900460ff16600381111561048c57fe5b145b6104c95760405162461bcd60e51b815260040180806020018281038252602e815260200180611222602e913960400191505060405180910390fd5b60016005880154600160401b900460ff1660038111156104e557fe5b14610553576004870154600160801b90046001600160401b0316431115610553576040805162461bcd60e51b815260206004820152601c60248201527f43616e27742063616e63656c2066696e616c697a656420737461746500000000604482015290519081900360640190fd5b6000806106718989888a60405160200180838152602001828152602001925050506040516020818303038152906040528051906020012060001973__$9836fa7140e5a33041d4b827682e675a30$__63e83f4bfe8b6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156105f45781810151838201526020016105dc565b50505050905090810190601f1680156106215780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b15801561063e57600080fd5b505af4158015610652573d6000803e3d6000fd5b505050506040513d602081101561066857600080fd5b50518989610e6b565b9092509050816106c1576040805162461bcd60e51b8152602060048201526016602482015275125b9d985b1a59081cda59db985d1d5c99481b1a5cdd60521b604482015290519081900360640190fd5b600289018690556040805182815290517f709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa9181900360200190a1505050505050505050565b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca886040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561075557600080fd5b505af4158015610769573d6000803e3d6000fd5b505050506040513d602081101561077f57600080fd5b5051156107d3576040805162461bcd60e51b815260206004820152601b60248201527f43616e2774206173736572742068616c746564206d616368696e650000000000604482015290519081900360640190fd5b60016005880154600160401b900460ff1660038111156107ef57fe5b1480610814575060026005880154600160401b900460ff16600381111561081257fe5b145b80610838575060036005880154600160401b900460ff16600381111561083657fe5b145b6108735760405162461bcd60e51b815260040180806020018281038252602d815260200180611284602d913960400191505060405180910390fd5b60016005880154600160401b900460ff16600381111561088f57fe5b146108fd576004870154600160801b90046001600160401b03164311156108fd576040805162461bcd60e51b815260206004820152601c60248201527f43616e27742063616e63656c2066696e616c697a656420737461746500000000604482015290519081900360640190fd5b60008061090f89898989898989610e6b565b90925090508161095f576040805162461bcd60e51b8152602060048201526016602482015275125b9d985b1a59081cda59db985d1d5c99481b1a5cdd60521b604482015290519081900360640190fd5b600360058a0154600160401b900460ff16600381111561097b57fe5b14156109d75760048901546001600160401b03600160c01b9091048116908716116109d75760405162461bcd60e51b81526004018080602001828103825260428152602001806111a66042913960600191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63a3a162cb8a6040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610a2657600080fd5b505af4158015610a3a573d6000803e3d6000fd5b50505050858960040160186101000a8154816001600160401b0302191690836001600160401b03160217905550848760405160200180838152602001828152602001925050506040516020818303038152906040528051906020012089600101819055507f4c6950de8aaa67cd052f9e28572dfff2ec4b8badd2f2b4bd8d8479d36987b6a481878b60040160109054906101000a90046001600160401b031660405180848152602001836001600160401b03166001600160401b03168152602001826001600160401b03166001600160401b03168152602001935050505060405180910390a1505050505050505050565b60036005850154600160401b900460ff166003811115610b4757fe5b14610b835760405162461bcd60e51b815260040180806020018281038252603a8152602001806111e8603a913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5856040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610bd257600080fd5b505af4158015610be6573d6000803e3d6000fd5b505050506040513d6020811015610bfc57600080fd5b505115610c3a5760405162461bcd60e51b815260040180806020018281038252603e815260200180611168603e913960400191505060405180910390fd5b600184015460405163741fa5ff60e11b815260206004820181815284516024840152845173__$9836fa7140e5a33041d4b827682e675a30$__9363e83f4bfe9387939283926044019185019080838360005b83811015610ca4578181015183820152602001610c8c565b50505050905090810190601f168015610cd15780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610cee57600080fd5b505af4158015610d02573d6000803e3d6000fd5b505050506040513d6020811015610d1857600080fd5b5051604080516020818101879052818301889052825180830384018152606083018452805190820120608083019490945260a0808301949094528251808303909401845260c0909101909152815191012014610da55760405162461bcd60e51b81526004018080602001828103825260348152602001806112506034913960400191505060405180910390fd5b6002840182905560408051633ad2660b60e21b81526004810186905260248101859052905173__$8e266570c8a7fb2aaac83b3e040afaf9e1$__9163eb49982c916044808301926000929190829003018186803b158015610e0557600080fd5b505af4158015610e19573d6000803e3d6000fd5b50505050600484015460408051600160c01b9092046001600160401b03168252517fbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2916020908290030190a150505050565b6000806000878a600001548b60020154898960405160200180868152602001858152602001848152602001836001600160401b03166001600160401b031660c01b815260080182815260200195505050505050604051602081830303815290604052805190602001209050600030828760405160200180846001600160a01b03166001600160a01b031660601b8152601401838152602001828152602001935050505060405160208183030381529060405280519060200120905060008a6001600160a01b031663513164fe73__$59c09a8a68cf3791d03cdf6ed66ba4d742$__63f0c8e969858a6040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610fa5578181015183820152602001610f8d565b50505050905090810190601f168015610fd25780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b158015610ff057600080fd5b505af4158015611004573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561102d57600080fd5b8101908080516040519392919084600160201b82111561104c57600080fd5b90830190602082018581111561106157600080fd5b82518660208202830111600160201b8211171561107d57600080fd5b82525081516020918201928201910280838360005b838110156110aa578181015183820152602001611092565b505050509050016040525050506040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019060200280838360005b838110156111065781810151838201526020016110ee565b505050509050019250505060206040518083038186803b15801561112957600080fd5b505afa15801561113d573d6000803e3d6000fd5b505050506040513d602081101561115357600080fd5b50519c919b5090995050505050505050505056fe43616e206f6e6c7920636f6e6669726d20617373657274696f6e2077686f7365206368616c6c656e676520646561646c696e65206861732070617373656443616e206f6e6c79207375706572736564652070726576696f757320617373657274696f6e207769746820677265617465722073657175656e6365206e756d62657243616e206f6e6c7920636f6e6669726d20696620746865726520697320612070656e64696e6720756e616e696d6f757320617373657274696f6e547269656420746f2066696e616c697a6520756e616e696d6f75732066726f6d20696e76616c696420737461746543616e206f6e6c7920636f6e6669726d20617373657274696f6e20746861742069732063757272656e746c792070656e64696e67547269656420746f2070656e64696e6720756e616e696d6f75732066726f6d20696e76616c6964207374617465a265627a7a7231582051fc001cdb51d1e95eeba6a49bcafd57202cbbadf74bcf2f836105b77b501da964736f6c634300050c0032"

// DeployUnanimous deploys a new Ethereum contract, binding an instance of Unanimous to it.
func DeployUnanimous(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Unanimous, error) {
	parsed, err := abi.JSON(strings.NewReader(UnanimousABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	sigUtilsAddr, _, _, _ := DeploySigUtils(auth, backend)
	UnanimousBin = strings.Replace(UnanimousBin, "__$59c09a8a68cf3791d03cdf6ed66ba4d742$__", sigUtilsAddr.String()[2:], -1)

	vMAddr, _, _, _ := DeployVM(auth, backend)
	UnanimousBin = strings.Replace(UnanimousBin, "__$8e266570c8a7fb2aaac83b3e040afaf9e1$__", vMAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	UnanimousBin = strings.Replace(UnanimousBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnanimousBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Unanimous{UnanimousCaller: UnanimousCaller{contract: contract}, UnanimousTransactor: UnanimousTransactor{contract: contract}, UnanimousFilterer: UnanimousFilterer{contract: contract}}, nil
}

// Unanimous is an auto generated Go binding around an Ethereum contract.
type Unanimous struct {
	UnanimousCaller     // Read-only binding to the contract
	UnanimousTransactor // Write-only binding to the contract
	UnanimousFilterer   // Log filterer for contract events
}

// UnanimousCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnanimousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnanimousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnanimousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnanimousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnanimousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnanimousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnanimousSession struct {
	Contract     *Unanimous        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnanimousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnanimousCallerSession struct {
	Contract *UnanimousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UnanimousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnanimousTransactorSession struct {
	Contract     *UnanimousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UnanimousRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnanimousRaw struct {
	Contract *Unanimous // Generic contract binding to access the raw methods on
}

// UnanimousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnanimousCallerRaw struct {
	Contract *UnanimousCaller // Generic read-only contract binding to access the raw methods on
}

// UnanimousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnanimousTransactorRaw struct {
	Contract *UnanimousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnanimous creates a new instance of Unanimous, bound to a specific deployed contract.
func NewUnanimous(address common.Address, backend bind.ContractBackend) (*Unanimous, error) {
	contract, err := bindUnanimous(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unanimous{UnanimousCaller: UnanimousCaller{contract: contract}, UnanimousTransactor: UnanimousTransactor{contract: contract}, UnanimousFilterer: UnanimousFilterer{contract: contract}}, nil
}

// NewUnanimousCaller creates a new read-only instance of Unanimous, bound to a specific deployed contract.
func NewUnanimousCaller(address common.Address, caller bind.ContractCaller) (*UnanimousCaller, error) {
	contract, err := bindUnanimous(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnanimousCaller{contract: contract}, nil
}

// NewUnanimousTransactor creates a new write-only instance of Unanimous, bound to a specific deployed contract.
func NewUnanimousTransactor(address common.Address, transactor bind.ContractTransactor) (*UnanimousTransactor, error) {
	contract, err := bindUnanimous(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnanimousTransactor{contract: contract}, nil
}

// NewUnanimousFilterer creates a new log filterer instance of Unanimous, bound to a specific deployed contract.
func NewUnanimousFilterer(address common.Address, filterer bind.ContractFilterer) (*UnanimousFilterer, error) {
	contract, err := bindUnanimous(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnanimousFilterer{contract: contract}, nil
}

// bindUnanimous binds a generic wrapper to an already deployed contract.
func bindUnanimous(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnanimousABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unanimous *UnanimousRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Unanimous.Contract.UnanimousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unanimous *UnanimousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unanimous.Contract.UnanimousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unanimous *UnanimousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unanimous.Contract.UnanimousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unanimous *UnanimousCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Unanimous.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unanimous *UnanimousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unanimous.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unanimous *UnanimousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unanimous.Contract.contract.Transact(opts, method, params...)
}

// UnanimousConfirmedUnanimousAssertionIterator is returned from FilterConfirmedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedUnanimousAssertion events raised by the Unanimous contract.
type UnanimousConfirmedUnanimousAssertionIterator struct {
	Event *UnanimousConfirmedUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *UnanimousConfirmedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnanimousConfirmedUnanimousAssertion)
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
		it.Event = new(UnanimousConfirmedUnanimousAssertion)
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
func (it *UnanimousConfirmedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnanimousConfirmedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnanimousConfirmedUnanimousAssertion represents a ConfirmedUnanimousAssertion event raised by the Unanimous contract.
type UnanimousConfirmedUnanimousAssertion struct {
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedUnanimousAssertion is a free log retrieval operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) FilterConfirmedUnanimousAssertion(opts *bind.FilterOpts) (*UnanimousConfirmedUnanimousAssertionIterator, error) {

	logs, sub, err := _Unanimous.contract.FilterLogs(opts, "ConfirmedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &UnanimousConfirmedUnanimousAssertionIterator{contract: _Unanimous.contract, event: "ConfirmedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedUnanimousAssertion is a free log subscription operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) WatchConfirmedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *UnanimousConfirmedUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _Unanimous.contract.WatchLogs(opts, "ConfirmedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnanimousConfirmedUnanimousAssertion)
				if err := _Unanimous.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
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

// ParseConfirmedUnanimousAssertion is a log parse operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) ParseConfirmedUnanimousAssertion(log types.Log) (*UnanimousConfirmedUnanimousAssertion, error) {
	event := new(UnanimousConfirmedUnanimousAssertion)
	if err := _Unanimous.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UnanimousFinalizedUnanimousAssertionIterator is returned from FilterFinalizedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for FinalizedUnanimousAssertion events raised by the Unanimous contract.
type UnanimousFinalizedUnanimousAssertionIterator struct {
	Event *UnanimousFinalizedUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *UnanimousFinalizedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnanimousFinalizedUnanimousAssertion)
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
		it.Event = new(UnanimousFinalizedUnanimousAssertion)
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
func (it *UnanimousFinalizedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnanimousFinalizedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnanimousFinalizedUnanimousAssertion represents a FinalizedUnanimousAssertion event raised by the Unanimous contract.
type UnanimousFinalizedUnanimousAssertion struct {
	UnanHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizedUnanimousAssertion is a free log retrieval operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_Unanimous *UnanimousFilterer) FilterFinalizedUnanimousAssertion(opts *bind.FilterOpts) (*UnanimousFinalizedUnanimousAssertionIterator, error) {

	logs, sub, err := _Unanimous.contract.FilterLogs(opts, "FinalizedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &UnanimousFinalizedUnanimousAssertionIterator{contract: _Unanimous.contract, event: "FinalizedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchFinalizedUnanimousAssertion is a free log subscription operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_Unanimous *UnanimousFilterer) WatchFinalizedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *UnanimousFinalizedUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _Unanimous.contract.WatchLogs(opts, "FinalizedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnanimousFinalizedUnanimousAssertion)
				if err := _Unanimous.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
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

// ParseFinalizedUnanimousAssertion is a log parse operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_Unanimous *UnanimousFilterer) ParseFinalizedUnanimousAssertion(log types.Log) (*UnanimousFinalizedUnanimousAssertion, error) {
	event := new(UnanimousFinalizedUnanimousAssertion)
	if err := _Unanimous.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UnanimousPendingUnanimousAssertionIterator is returned from FilterPendingUnanimousAssertion and is used to iterate over the raw logs and unpacked data for PendingUnanimousAssertion events raised by the Unanimous contract.
type UnanimousPendingUnanimousAssertionIterator struct {
	Event *UnanimousPendingUnanimousAssertion // Event containing the contract specifics and raw log

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
func (it *UnanimousPendingUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnanimousPendingUnanimousAssertion)
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
		it.Event = new(UnanimousPendingUnanimousAssertion)
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
func (it *UnanimousPendingUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnanimousPendingUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnanimousPendingUnanimousAssertion represents a PendingUnanimousAssertion event raised by the Unanimous contract.
type UnanimousPendingUnanimousAssertion struct {
	UnanHash    [32]byte
	SequenceNum uint64
	Deadline    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingUnanimousAssertion is a free log retrieval operation binding the contract event 0x4c6950de8aaa67cd052f9e28572dfff2ec4b8badd2f2b4bd8d8479d36987b6a4.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum, uint64 deadline)
func (_Unanimous *UnanimousFilterer) FilterPendingUnanimousAssertion(opts *bind.FilterOpts) (*UnanimousPendingUnanimousAssertionIterator, error) {

	logs, sub, err := _Unanimous.contract.FilterLogs(opts, "PendingUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &UnanimousPendingUnanimousAssertionIterator{contract: _Unanimous.contract, event: "PendingUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingUnanimousAssertion is a free log subscription operation binding the contract event 0x4c6950de8aaa67cd052f9e28572dfff2ec4b8badd2f2b4bd8d8479d36987b6a4.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum, uint64 deadline)
func (_Unanimous *UnanimousFilterer) WatchPendingUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *UnanimousPendingUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _Unanimous.contract.WatchLogs(opts, "PendingUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnanimousPendingUnanimousAssertion)
				if err := _Unanimous.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
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

// ParsePendingUnanimousAssertion is a log parse operation binding the contract event 0x4c6950de8aaa67cd052f9e28572dfff2ec4b8badd2f2b4bd8d8479d36987b6a4.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum, uint64 deadline)
func (_Unanimous *UnanimousFilterer) ParsePendingUnanimousAssertion(log types.Log) (*UnanimousPendingUnanimousAssertion, error) {
	event := new(UnanimousPendingUnanimousAssertion)
	if err := _Unanimous.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
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
