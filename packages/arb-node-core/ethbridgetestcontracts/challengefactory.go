// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

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

// ChallengeFactoryABI is the input ABI used to generate the binding from.
const ChallengeFactoryABI = "[{\"inputs\":[{\"internalType\":\"contractIOneStepProof[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"challengeTemplate\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_resultReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_executionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_asserterTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengerTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"_sequencerBridge\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"_delayedBridge\",\"type\":\"address\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"executors\",\"outputs\":[{\"internalType\":\"contractIOneStepProof\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeFactoryBin is the compiled bytecode used for deploying new contracts.
var ChallengeFactoryBin = "0x608060405234801561001057600080fd5b506040516125df3803806125df8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b505050509050016040525050506040516100cb90610122565b604051809103906000f0801580156100e7573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055805161011b90600190602084019061012f565b50506101b3565b611fb58061062a83390190565b828054828255906000526020600020908101928215610184579160200282015b8281111561018457825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061014f565b50610190929150610194565b5090565b5b808211156101905780546001600160a01b0319168155600101610195565b610468806101c26000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063777367c9146100465780638ecaab111461006a578063f97a05df146100cc575b600080fd5b61004e6100e9565b604080516001600160a01b039092168252519081900360200190f35b61004e600480360361012081101561008157600080fd5b506001600160a01b0381358116916020810135916040820135916060810135821691608082013581169160a08101359160c08201359160e081013582169161010090910135166100f8565b61004e600480360360208110156100e257600080fd5b503561023d565b6000546001600160a01b031681565b600080548190610110906001600160a01b0316610264565b9050806001600160a01b031663e0d42b8e60018d8d8d8d8d8d8d8d8d6040518b63ffffffff1660e01b815260040180806020018b6001600160a01b031681526020018a8152602001898152602001886001600160a01b03168152602001876001600160a01b03168152602001868152602001858152602001846001600160a01b03168152602001836001600160a01b0316815260200182810382528c81815481526020019150805480156101ed57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116101cf575b50509b505050505050505050505050600060405180830381600087803b15801561021657600080fd5b505af115801561022a573d6000803e3d6000fd5b50929d9c50505050505050505050505050565b6001818154811061024a57fe5b6000918252602090912001546001600160a01b0316905081565b6000816001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b15801561029f57600080fd5b505afa1580156102b3573d6000803e3d6000fd5b505050506040513d60208110156102c957600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b6020820152906103775760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561033c578181015183820152602001610324565b50505050905090810190601f1680156103695780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061038a826001600160a01b0316610390565b92915050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b03811661042d576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b91905056fea2646970667358221220cdf2eecced07c96f63d8200029d2c9481424b0f937312e9a77b3886e709903df64736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff19166001179055611f888061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100db5760003560e01c806341e8510c146100e0578063534db0e2146100fa5780636f791d291461011e57806370dea79a1461013a578063843d5a5c146101445780638a8cd2181461014c5780638b299903146101545780638e7b84c51461017d578063925f9a961461026057806395979201146102685780639a9e4f4414610439578063a3c4470514610441578063bb4af0b11461045e578063deda411514610466578063e0d42b8e146104f2578063e87e3589146105a5578063f97a05df146105ad575b600080fd5b6100e86105ca565b60408051918252519081900360200190f35b6101026105d0565b604080516001600160a01b039092168252519081900360200190f35b6101266105df565b604080519115158252519081900360200190f35b6101426105e9565b005b6100e8610742565b610102610748565b61015c6107d7565b6040518082600281111561016c57fe5b815260200191505060405180910390f35b610142600480360361010081101561019457600080fd5b810190602081018135600160201b8111156101ae57600080fd5b8201836020820111156101c057600080fd5b803590602001918460208302840111600160201b831117156101e157600080fd5b9193909282359260208101359260408201359260608301359260808101359260a082013592909160e081019060c00135600160201b81111561022257600080fd5b82018360208201111561023457600080fd5b803590602001918460208302840111600160201b8311171561025557600080fd5b5090925090506107e0565b6100e8610c86565b61014260048036036101c081101561027f57600080fd5b810190602081018135600160201b81111561029957600080fd5b8201836020820111156102ab57600080fd5b803590602001918460208302840111600160201b831117156102cc57600080fd5b6040805160608181018352949693958335956020850135959385013594818101359460808201359460a08301949193919261014081019260e090910190600390839083908082843760009201919091525091949392602081019250359050600160201b81111561033b57600080fd5b82018360208201111561034d57600080fd5b803590602001918460018302840111600160201b8311171561036e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103c057600080fd5b8201836020820111156103d257600080fd5b803590602001918460018302840111600160201b831117156103f357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff169150610c8c9050565b6100e86111a9565b6101026004803603602081101561045757600080fd5b50356111af565b6101026111cc565b610142600480360360e081101561047c57600080fd5b810190602081018135600160201b81111561049657600080fd5b8201836020820111156104a857600080fd5b803590602001918460208302840111600160201b831117156104c957600080fd5b919350915080359060208101359060408101359060608101359060808101359060a001356111db565b610142600480360361014081101561050957600080fd5b810190602081018135600160201b81111561052357600080fd5b82018360208201111561053557600080fd5b803590602001918460208302840111600160201b8311171561055657600080fd5b91935091506001600160a01b0381358116916020810135916040820135916060810135821691608082013581169160a08101359160c08201359160e08101358216916101009091013516611456565b6100e86115a7565b610102600480360360208110156105c357600080fd5b50356115ed565b600a5481565b6007546001600160a01b031681565b60005460ff165b90565b60006106006008544361161490919063ffffffff16565b905061060a6115a7565b81116040518060400160405280601081526020016f54494d454f55545f444541444c494e4560801b815250906106be5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561068357818101518382015260200161066b565b50505050905090810190601f1680156106b05780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506001600b5460ff1660028111156106d257fe5b141561070e576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a1610709611656565b61073f565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a161073f6116d5565b50565b600c5481565b60006001600b5460ff16600281111561075d57fe5b141561077557506006546001600160a01b03166105e6565b6002600b5460ff16600281111561078857fe5b14156107a057506007546001600160a01b03166105e6565b6040805162461bcd60e51b81526020600482015260076024820152662727afaa2aa92760c91b604482015290519081900360640190fd5b600b5460ff1681565b6107e8610748565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b8152509061086b5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b506108746115a7565b600854610882904390611614565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906108f55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b5060008282600019810181811061090857fe5b90506020020135146109555760018611610955576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b604482015290519081900360640190fd5b61096186610190611731565b60010181146109a3576040805162461bcd60e51b815260206004820152600960248201526810d55517d0d3d5539560ba1b604482015290519081900360640190fd5b84828260001981018181106109b457fe5b9050602002013514156109f9576040805162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b604482015290519081900360640190fd5b610a038484611749565b82826000818110610a1057fe5b9050602002013514610a5e576040805162461bcd60e51b81526020600482015260126024820152717365676d656e74207072652d6669656c647360701b604482015290519081900360640190fd5b600082828281610a6a57fe5b905060200201351415610ab8576040805162461bcd60e51b8152602060048201526011602482015270155394915050d21050931157d4d5105495607a1b604482015290519081900360640190fd5b610ac28787611775565b8410610b0e576040805162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840e6cacedacadce840d8cadccee8d60531b604482015290519081900360640190fd5b6000610b2f888885856000818110610b2257fe5b90506020020135896117c3565b9050610b3d818c8c8c611801565b610b7d8383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508c92508b91506118b99050565b50600c547f0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d8989868660405180858152602001848152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f191690920182900397509095505050505050a2506002600b5460ff166002811115610c0857fe5b1415610c4457610c2f610c266008544361161490919063ffffffff16565b600a5490611614565b600a55600b805460ff19166001179055610c76565b610c65610c5c6008544361161490919063ffffffff16565b60095490611614565b600955600b805460ff191660021790555b5050436008555050505050505050565b60085481565b610c94610748565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b81525090610d175760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b50610d206115a7565b600854610d2e904390611614565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610da15760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b506000806000610daf611e6a565b60018560ff1681548110610dbf57fe5b6000918252602090912001546040516323eed0eb60e11b81526001600160a01b03909116906347dda1d6906002908d908d908c908c90600481019060440186825b81546001600160a01b03168152600190910190602001808311610e0057505085815260200184604080828437600081840152601f19601f8201169050808301925050508060200180602001838103835285818151815260200191508051906020019080838360005b83811015610e80578181015183820152602001610e68565b50505050905090810190601f168015610ead5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610ee0578181015183820152602001610ec8565b50505050905090810190601f168015610f0d5780820380516001836020036101000a031916815260200191505b5097505050505050505060c06040518083038186803b158015610f2f57600080fd5b505afa158015610f43573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c0811015610f6857600080fd5b5080516020820151600554919550935060409091019150821115610fc7576040805162461bcd60e51b8152602060048201526011602482015270544f4f5f4d414e595f4d4553534147455360781b604482015290519081900360640190fd5b610fd18d8d611775565b885110611010576040805162461bcd60e51b815260206004820152600860248201526713d4d417d0d3d39560c21b604482015290519081900360640190fd5b61101a8d8d611775565b6110366001600160401b0385168a60005b602002015190611775565b1015611075576040805162461bcd60e51b815260206004820152600960248201526813d4d417d4d213d49560ba1b604482015290519081900360640190fd5b611088893560208b01358a8686866119f5565b8b14156110c8576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6110e58d8d6110df8d8d3560208f01358e88611a93565b8e6117c3565b93505050506110f6818e8e8e611801565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1611127611ac4565b506002600b5460ff16600281111561113b57fe5b141561116e57611159610c266008544361161490919063ffffffff16565b600a55600b805460ff19166001179055611197565b611186610c5c6008544361161490919063ffffffff16565b600955600b805460ff191660021790555b50504360085550505050505050505050565b60095481565b600281600281106111bc57fe5b01546001600160a01b0316905081565b6006546001600160a01b031681565b6111e3610748565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b815250906112665760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b5061126f6115a7565b60085461127d904390611614565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906112f05760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b5060006112fd8383611749565b9050600061130d878784886117c3565b905061131b818b8b8b611801565b6113258787611775565b841015611364576040805162461bcd60e51b81526020600482015260086024820152671393d517d0d3d39560c21b604482015290519081900360640190fd5b848214156113a5576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6040517ff62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca90600090a16113d6611ac4565b5060029050600b5460ff1660028111156113ec57fe5b141561141f5761140a610c266008544361161490919063ffffffff16565b600a55600b805460ff19166001179055611448565b611437610c5c6008544361161490919063ffffffff16565b600955600b805460ff191660021790555b505043600855505050505050565b6000600b5460ff16600281111561146957fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b815250906114de5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b506114eb60018c8c611e88565b50600480546001600160a01b03199081166001600160a01b038c8116919091179092556005899055600680548216898416179055600780549091168783161790556009859055600a849055600b805460ff19166002908117909155600c8a90554360085560408051808201909152848316815291831660208301526115709181611eeb565b506040517f7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de90600090a15050505050505050505050565b60006001600b5460ff1660028111156115bc57fe5b14156115cb57506009546105e6565b6002600b5460ff1660028111156115de57fe5b14156107a05750600a546105e6565b600181815481106115fa57fe5b6000918252602090912001546001600160a01b0316905081565b600061164d8383604051806040016040528060148152602001737375627472616374696f6e206f766572666c6f7760601b815250611af2565b90505b92915050565b6004805460075460065460408051637d3c01f360e11b81526001600160a01b039384169581019590955290821660248501525191169163fa7803e691604480830192600092919082900301818387803b1580156116b257600080fd5b505af11580156116c6573d6000803e3d6000fd5b505050506116d333611b4c565b565b6004805460065460075460408051637d3c01f360e11b81526001600160a01b039384169581019590955290821660248501525191169163fa7803e691604480830192600092919082900301818387803b1580156116b257600080fd5b600081831015611742575081611650565b5080611650565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008282018381101561164d576040805162461bcd60e51b81526020600482015260116024820152706164646974696f6e206f766572666c6f7760781b604482015290519081900360640190fd5b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b611841838380806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250859250889150611bcc9050565b600c5414604051806040016040528060088152602001672124a9afa82922ab60c11b815250906118b25760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b5050505050565b8251600090600019016060816001600160401b03811180156118da57600080fd5b50604051908082528060200260200182016040528015611904578160200160208202803683370190505b50905060006119138584611c9a565b9050600086905061194e81838a60008151811061192c57fe5b60200260200101518b60018151811061194157fe5b60200260200101516117c3565b8360008151811061195b57fe5b60209081029190910101526119708183611775565b905061197c8685611cb8565b915060015b848110156119dc576119b182848b848151811061199a57fe5b60200260200101518c856001018151811061194157fe5b8482815181106119bd57fe5b60209081029190910101526119d28284611775565b9150600101611981565b506119e683611ccb565b600c5550929695505050505050565b600080611a1e83600260200201518914611a10576001611a13565b60005b60ff1687600161102b565b90506000611a4884600360200201518914611a3a576001611a3d565b60005b60ff1688600261102b565b9050611a86611a626001600160401b03881689600061102b565b602086015160408701516060880151611a81928a929091889088611e1f565b611749565b9998505050505050505050565b8151815160208401516040850151600093611aba939092611a81928b92918b918b90611e1f565b9695505050505050565b6001600b5460ff166002811115611ad757fe5b1415611aea57611ae56116d5565b6116d3565b6116d3611656565b60008184841115611b445760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b505050900390565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff1615611bbf5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068357818101518382015260200161066b565b50806001600160a01b0316ff5b8251600090610100811115611be057600080fd5b8260005b82811015611c905760028606611c3d57868181518110611c0057fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150611c82565b81878281518110611c4a57fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101611be4565b5095945050505050565b6000818381611ca557fe5b06828481611caf57fe5b04019392505050565b6000818381611cc357fe5b049392505050565b6000815b600181511115611e025760606002825160010181611ce957fe5b046001600160401b0381118015611cff57600080fd5b50604051908082528060200260200182016040528015611d29578160200160208202803683370190505b50905060005b8151811015611dfa578251816002026001011015611dc257828160020281518110611d5657fe5b6020026020010151838260020260010181518110611d7057fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611db157fe5b602002602001018181525050611df2565b828160020281518110611dd157fe5b6020026020010151828281518110611de557fe5b6020026020010181815250505b600101611d2f565b509050611ccf565b80600081518110611e0f57fe5b6020026020010151915050919050565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b60405180608001604052806004906020820280368337509192915050565b828054828255906000526020600020908101928215611edb579160200282015b82811115611edb5781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190611ea8565b50611ee7929150611f33565b5090565b8260028101928215611edb579160200282015b82811115611edb57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611efe565b5b80821115611ee75780546001600160a01b0319168155600101611f3456fea26469706673582212203320d6ed074c7b13ed6a066daca18c5b354dfeb5ba4763372feeb32cad776bbd64736f6c634300060c0033"

// DeployChallengeFactory deploys a new Ethereum contract, binding an instance of ChallengeFactory to it.
func DeployChallengeFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _executors []common.Address) (common.Address, *types.Transaction, *ChallengeFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeFactoryBin), backend, _executors)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeFactory{ChallengeFactoryCaller: ChallengeFactoryCaller{contract: contract}, ChallengeFactoryTransactor: ChallengeFactoryTransactor{contract: contract}, ChallengeFactoryFilterer: ChallengeFactoryFilterer{contract: contract}}, nil
}

// ChallengeFactory is an auto generated Go binding around an Ethereum contract.
type ChallengeFactory struct {
	ChallengeFactoryCaller     // Read-only binding to the contract
	ChallengeFactoryTransactor // Write-only binding to the contract
	ChallengeFactoryFilterer   // Log filterer for contract events
}

// ChallengeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeFactorySession struct {
	Contract     *ChallengeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeFactoryCallerSession struct {
	Contract *ChallengeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeFactoryTransactorSession struct {
	Contract     *ChallengeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeFactoryRaw struct {
	Contract *ChallengeFactory // Generic contract binding to access the raw methods on
}

// ChallengeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeFactoryCallerRaw struct {
	Contract *ChallengeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeFactoryTransactorRaw struct {
	Contract *ChallengeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeFactory creates a new instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactory(address common.Address, backend bind.ContractBackend) (*ChallengeFactory, error) {
	contract, err := bindChallengeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactory{ChallengeFactoryCaller: ChallengeFactoryCaller{contract: contract}, ChallengeFactoryTransactor: ChallengeFactoryTransactor{contract: contract}, ChallengeFactoryFilterer: ChallengeFactoryFilterer{contract: contract}}, nil
}

// NewChallengeFactoryCaller creates a new read-only instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryCaller(address common.Address, caller bind.ContractCaller) (*ChallengeFactoryCaller, error) {
	contract, err := bindChallengeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryCaller{contract: contract}, nil
}

// NewChallengeFactoryTransactor creates a new write-only instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeFactoryTransactor, error) {
	contract, err := bindChallengeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryTransactor{contract: contract}, nil
}

// NewChallengeFactoryFilterer creates a new log filterer instance of ChallengeFactory, bound to a specific deployed contract.
func NewChallengeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFactoryFilterer, error) {
	contract, err := bindChallengeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFactoryFilterer{contract: contract}, nil
}

// bindChallengeFactory binds a generic wrapper to an already deployed contract.
func bindChallengeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeFactory *ChallengeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeFactory.Contract.ChallengeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeFactory *ChallengeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.ChallengeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeFactory *ChallengeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.ChallengeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeFactory *ChallengeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeFactory *ChallengeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeFactory *ChallengeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.contract.Transact(opts, method, params...)
}

// ChallengeTemplate is a free data retrieval call binding the contract method 0x777367c9.
//
// Solidity: function challengeTemplate() view returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) ChallengeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeFactory.contract.Call(opts, &out, "challengeTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeTemplate is a free data retrieval call binding the contract method 0x777367c9.
//
// Solidity: function challengeTemplate() view returns(address)
func (_ChallengeFactory *ChallengeFactorySession) ChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.ChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// ChallengeTemplate is a free data retrieval call binding the contract method 0x777367c9.
//
// Solidity: function challengeTemplate() view returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) ChallengeTemplate() (common.Address, error) {
	return _ChallengeFactory.Contract.ChallengeTemplate(&_ChallengeFactory.CallOpts)
}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_ChallengeFactory *ChallengeFactoryCaller) Executors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ChallengeFactory.contract.Call(opts, &out, "executors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_ChallengeFactory *ChallengeFactorySession) Executors(arg0 *big.Int) (common.Address, error) {
	return _ChallengeFactory.Contract.Executors(&_ChallengeFactory.CallOpts, arg0)
}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_ChallengeFactory *ChallengeFactoryCallerSession) Executors(arg0 *big.Int) (common.Address, error) {
	return _ChallengeFactory.Contract.Executors(&_ChallengeFactory.CallOpts, arg0)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x8ecaab11.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactor) CreateChallenge(opts *bind.TransactOpts, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.contract.Transact(opts, "createChallenge", _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x8ecaab11.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns(address)
func (_ChallengeFactory *ChallengeFactorySession) CreateChallenge(_resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x8ecaab11.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactorSession) CreateChallenge(_resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// IOneStepProofABI is the input ABI used to generate the binding from.
const IOneStepProofABI = "[{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"bridges\",\"type\":\"address[2]\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"afterMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[4]\",\"name\":\"fields\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"bridges\",\"type\":\"address[2]\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStepDebug\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"startMachine\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"afterMachine\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IOneStepProof is an auto generated Go binding around an Ethereum contract.
type IOneStepProof struct {
	IOneStepProofCaller     // Read-only binding to the contract
	IOneStepProofTransactor // Write-only binding to the contract
	IOneStepProofFilterer   // Log filterer for contract events
}

// IOneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProofSession struct {
	Contract     *IOneStepProof    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProofCallerSession struct {
	Contract *IOneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IOneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProofTransactorSession struct {
	Contract     *IOneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IOneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProofRaw struct {
	Contract *IOneStepProof // Generic contract binding to access the raw methods on
}

// IOneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProofCallerRaw struct {
	Contract *IOneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProofTransactorRaw struct {
	Contract *IOneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProof creates a new instance of IOneStepProof, bound to a specific deployed contract.
func NewIOneStepProof(address common.Address, backend bind.ContractBackend) (*IOneStepProof, error) {
	contract, err := bindIOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProof{IOneStepProofCaller: IOneStepProofCaller{contract: contract}, IOneStepProofTransactor: IOneStepProofTransactor{contract: contract}, IOneStepProofFilterer: IOneStepProofFilterer{contract: contract}}, nil
}

// NewIOneStepProofCaller creates a new read-only instance of IOneStepProof, bound to a specific deployed contract.
func NewIOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProofCaller, error) {
	contract, err := bindIOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofCaller{contract: contract}, nil
}

// NewIOneStepProofTransactor creates a new write-only instance of IOneStepProof, bound to a specific deployed contract.
func NewIOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProofTransactor, error) {
	contract, err := bindIOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofTransactor{contract: contract}, nil
}

// NewIOneStepProofFilterer creates a new log filterer instance of IOneStepProof, bound to a specific deployed contract.
func NewIOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProofFilterer, error) {
	contract, err := bindIOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofFilterer{contract: contract}, nil
}

// bindIOneStepProof binds a generic wrapper to an already deployed contract.
func bindIOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProof *IOneStepProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProof.Contract.IOneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProof *IOneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProof.Contract.IOneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProof *IOneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProof.Contract.IOneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProof *IOneStepProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProof *IOneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProof *IOneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProof.Contract.contract.Transact(opts, method, params...)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x47dda1d6.
//
// Solidity: function executeStep(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 afterMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofCaller) ExecuteStep(opts *bind.CallOpts, bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	AfterMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	var out []interface{}
	err := _IOneStepProof.contract.Call(opts, &out, "executeStep", bridges, initialMessagesRead, accs, proof, bproof)

	outstruct := new(struct {
		Gas               uint64
		AfterMessagesRead *big.Int
		Fields            [4][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gas = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AfterMessagesRead = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fields = *abi.ConvertType(out[2], new([4][32]byte)).(*[4][32]byte)

	return *outstruct, err

}

// ExecuteStep is a free data retrieval call binding the contract method 0x47dda1d6.
//
// Solidity: function executeStep(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 afterMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofSession) ExecuteStep(bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	AfterMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _IOneStepProof.Contract.ExecuteStep(&_IOneStepProof.CallOpts, bridges, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x47dda1d6.
//
// Solidity: function executeStep(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 afterMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofCallerSession) ExecuteStep(bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	AfterMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _IOneStepProof.Contract.ExecuteStep(&_IOneStepProof.CallOpts, bridges, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0xeba67f6e.
//
// Solidity: function executeStepDebug(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofCaller) ExecuteStepDebug(opts *bind.CallOpts, bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	var out []interface{}
	err := _IOneStepProof.contract.Call(opts, &out, "executeStepDebug", bridges, initialMessagesRead, accs, proof, bproof)

	outstruct := new(struct {
		StartMachine string
		AfterMachine string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartMachine = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AfterMachine = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0xeba67f6e.
//
// Solidity: function executeStepDebug(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofSession) ExecuteStepDebug(bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _IOneStepProof.Contract.ExecuteStepDebug(&_IOneStepProof.CallOpts, bridges, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0xeba67f6e.
//
// Solidity: function executeStepDebug(address[2] bridges, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofCallerSession) ExecuteStepDebug(bridges [2]common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _IOneStepProof.Contract.ExecuteStepDebug(&_IOneStepProof.CallOpts, bridges, initialMessagesRead, accs, proof, bproof)
}
