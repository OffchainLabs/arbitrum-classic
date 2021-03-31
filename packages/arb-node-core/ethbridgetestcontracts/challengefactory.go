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
const ChallengeFactoryABI = "[{\"inputs\":[{\"internalType\":\"contractIOneStepProof[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"challengeTemplate\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_resultReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_executionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSequencerBatchCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_asserterTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengerTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"_sequencerBridge\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"_delayedBridge\",\"type\":\"address\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"executors\",\"outputs\":[{\"internalType\":\"contractIOneStepProof\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeFactoryBin is the compiled bytecode used for deploying new contracts.
var ChallengeFactoryBin = "0x608060405234801561001057600080fd5b506040516126c93803806126c98339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b505050509050016040525050506040516100cb90610122565b604051809103906000f0801580156100e7573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055805161011b90600190602084019061012f565b50506101b3565b6120368061069383390190565b828054828255906000526020600020908101928215610184579160200282015b8281111561018457825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061014f565b50610190929150610194565b5090565b5b808211156101905780546001600160a01b0319168155600101610195565b6104d1806101c26000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063777367c914610046578063820eb6461461006a578063f97a05df146100d1575b600080fd5b61004e6100ee565b604080516001600160a01b039092168252519081900360200190f35b61004e600480360361014081101561008157600080fd5b506001600160a01b038135811691602081013591604082013591606081013591608082013581169160a081013582169160c08201359160e0810135916101008201358116916101200135166100fd565b61004e600480360360208110156100e757600080fd5b5035610288565b6000546001600160a01b031681565b600080548190610115906001600160a01b03166102af565b905061011f61047d565b60405180604001604052808c81526020018b8152509050816001600160a01b03166317e7739260018f8f858e8e8e8e8e8e6040518b63ffffffff1660e01b815260040180806020018b6001600160a01b031681526020018a815260200189600260200280838360005b838110156101a0578181015183820152602001610188565b50505050905001886001600160a01b03168152602001876001600160a01b03168152602001868152602001858152602001846001600160a01b03168152602001836001600160a01b0316815260200182810382528c818154815260200191508054801561023657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610218575b50509b505050505050505050505050600060405180830381600087803b15801561025f57600080fd5b505af1158015610273573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b6001818154811061029557fe5b6000918252602090912001546001600160a01b0316905081565b6000816001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b1580156102ea57600080fd5b505afa1580156102fe573d6000803e3d6000fd5b505050506040513d602081101561031457600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b6020820152906103c25760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561038757818101518382015260200161036f565b50505050905090810190601f1680156103b45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506103d5826001600160a01b03166103db565b92915050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116610478576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b919050565b6040518060400160405280600290602082028036833750919291505056fea2646970667358221220f7f323f6465df95f061c90f4980b7c52ea34ade4c8671d7fd1013360e94a2d7c64736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff191660011790556120098061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100e65760003560e01c806317e77392146100eb5780633e55c0c7146101a057806341e8510c146101c4578063534db0e2146101de5780636f791d29146101e6578063700d37581461020257806370dea79a146103d3578063843d5a5c146103db5780638a8cd218146103e35780638b299903146103eb5780638e7b84c514610414578063925f9a96146104f75780639a9e4f44146104ff578063bb4af0b114610507578063deda41151461050f578063e87e35891461059b578063f51de41b146105a3578063f97a05df146105ab575b600080fd5b61019e600480360361016081101561010257600080fd5b810190602081018135600160201b81111561011c57600080fd5b82018360208201111561012e57600080fd5b803590602001918460208302840111600160201b8311171561014f57600080fd5b91935091506001600160a01b0381358116916020810135916040820191608081013582169160a082013581169160c08101359160e08201359161010081013582169161012090910135166105c8565b005b6101a8610752565b604080516001600160a01b039092168252519081900360200190f35b6101cc610761565b60408051918252519081900360200190f35b6101a8610767565b6101ee610776565b604080519115158252519081900360200190f35b61019e60048036036101e081101561021957600080fd5b810190602081018135600160201b81111561023357600080fd5b82018360208201111561024557600080fd5b803590602001918460208302840111600160201b8311171561026657600080fd5b60408051606081810183529496939583359560208501359593850135948181013594608082019460c08301949193919261016081019261010090910190600390839083908082843760009201919091525091949392602081019250359050600160201b8111156102d557600080fd5b8201836020820111156102e757600080fd5b803590602001918460018302840111600160201b8311171561030857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561035a57600080fd5b82018360208201111561036c57600080fd5b803590602001918460018302840111600160201b8311171561038d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff1691506107809050565b61019e610d47565b6101cc610e63565b6101a8610e69565b6103f3610ef8565b6040518082600281111561040357fe5b815260200191505060405180910390f35b61019e600480360361010081101561042b57600080fd5b810190602081018135600160201b81111561044557600080fd5b82018360208201111561045757600080fd5b803590602001918460208302840111600160201b8311171561047857600080fd5b9193909282359260208101359260408201359260608301359260808101359260a082013592909160e081019060c00135600160201b8111156104b957600080fd5b8201836020820111156104cb57600080fd5b803590602001918460208302840111600160201b831117156104ec57600080fd5b509092509050610f01565b6101cc611395565b6101cc61139b565b6101a86113a1565b61019e600480360360e081101561052557600080fd5b810190602081018135600160201b81111561053f57600080fd5b82018360208201111561055157600080fd5b803590602001918460208302840111600160201b8311171561057257600080fd5b919350915080359060208101359060408101359060608101359060808101359060a001356113b0565b6101cc61162b565b6101a8611671565b6101a8600480360360208110156105c157600080fd5b5035611680565b6000600c5460ff1660028111156105db57fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b8152509061068d5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561065257818101518382015260200161063a565b50505050905090810190601f16801561067f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061069a60018c8c611f15565b50600480546001600160a01b038b81166001600160a01b03199283161790925588356005556020890135600655600780548216898416179055600880548216888416179055600a869055600b859055600c805460ff19166002908117909155600d8b90554360095580548216858416179055600380549091169183169190911790556040517f7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de90600090a15050505050505050505050565b6002546001600160a01b031681565b600b5481565b6008546001600160a01b031681565b60005460ff165b90565b610788610e69565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b8152509061080b5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b5061081461162b565b6009546108229043906116a7565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906108955760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b506000806108a1611f78565b6108a9611f96565b60018560ff16815481106108b957fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b03166392922de5600260009054906101000a90046001600160a01b0316600360009054906101000a90046001600160a01b03168d8d8c8c6040518763ffffffff1660e01b815260040180876001600160a01b03168152602001866001600160a01b0316815260200185600260200280828437600083820152601f01601f1916909101905084604080828437600081840152601f19601f8201169050808301925050508060200180602001838103835285818151815260200191508051906020019080838360005b838110156109bd5781810151838201526020016109a5565b50505050905090810190601f1680156109ea5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610a1d578181015183820152602001610a05565b50505050905090810190601f168015610a4a5780820380516001836020036101000a031916815260200191505b509850505050505050505060e06040518083038186803b158015610a6d57600080fd5b505afa158015610a81573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060e0811015610aa657600080fd5b5080516005546020830180519296509450606090920192501115610b05576040805162461bcd60e51b8152602060048201526011602482015270544f4f5f4d414e595f4d4553534147455360781b604482015290519081900360640190fd5b60065460208301511115610b53576040805162461bcd60e51b815260206004820152601060248201526f544f4f5f4d414e595f4241544348455360801b604482015290519081900360640190fd5b610b5d8d8d6116e9565b885110610b9c576040805162461bcd60e51b815260206004820152600860248201526713d4d417d0d3d39560c21b604482015290519081900360640190fd5b610ba68d8d6116e9565b610bc26001600160401b0385168a60005b6020020151906116e9565b1015610c01576040805162461bcd60e51b815260206004820152600960248201526813d4d417d4d213d49560ba1b604482015290519081900360640190fd5b610c14893560208b01358a868686611737565b8b1415610c54576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b610c718d8d610c6b8d8d3560208f01358e886117dc565b8e611815565b9350505050610c82818e8e8e611853565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1610cb361190b565b506002600c5460ff166002811115610cc757fe5b1415610d0357610cee610ce5600954436116a790919063ffffffff16565b600b54906116a7565b600b55600c805460ff19166001179055610d35565b610d24610d1b600954436116a790919063ffffffff16565b600a54906116a7565b600a55600c805460ff191660021790555b50504360095550505050505050505050565b6000610d5e600954436116a790919063ffffffff16565b9050610d6861162b565b81116040518060400160405280601081526020016f54494d454f55545f444541444c494e4560801b81525090610ddf5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b506001600c5460ff166002811115610df357fe5b1415610e2f576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a1610e2a61193b565b610e60565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a1610e606119b8565b50565b600d5481565b60006001600c5460ff166002811115610e7e57fe5b1415610e9657506007546001600160a01b031661077d565b6002600c5460ff166002811115610ea957fe5b1415610ec157506008546001600160a01b031661077d565b6040805162461bcd60e51b81526020600482015260076024820152662727afaa2aa92760c91b604482015290519081900360640190fd5b600c5460ff1681565b610f09610e69565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b81525090610f8c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b50610f9561162b565b600954610fa39043906116a7565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906110165760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b5060008282600019810181811061102957fe5b90506020020135146110765760018611611076576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b604482015290519081900360640190fd5b61108286610190611a14565b60010181146110c4576040805162461bcd60e51b815260206004820152600960248201526810d55517d0d3d5539560ba1b604482015290519081900360640190fd5b84828260001981018181106110d557fe5b90506020020135141561111a576040805162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b604482015290519081900360640190fd5b6111248484611a2c565b8282600081811061113157fe5b905060200201351461117f576040805162461bcd60e51b81526020600482015260126024820152717365676d656e74207072652d6669656c647360701b604482015290519081900360640190fd5b60008282828161118b57fe5b9050602002013514156111d9576040805162461bcd60e51b8152602060048201526011602482015270155394915050d21050931157d4d5105495607a1b604482015290519081900360640190fd5b6111e387876116e9565b841061122f576040805162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840e6cacedacadce840d8cadccee8d60531b604482015290519081900360640190fd5b600061125088888585600081811061124357fe5b9050602002013589611815565b905061125e818c8c8c611853565b61129e8383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508c92508b9150611a589050565b50600d547f0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d8989868660405180858152602001848152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f191690920182900397509095505050505050a2506002600c5460ff16600281111561132957fe5b141561135c57611347610ce5600954436116a790919063ffffffff16565b600b55600c805460ff19166001179055611385565b611374610d1b600954436116a790919063ffffffff16565b600a55600c805460ff191660021790555b5050436009555050505050505050565b60095481565b600a5481565b6007546001600160a01b031681565b6113b8610e69565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b8152509061143b5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b5061144461162b565b6009546114529043906116a7565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906114c55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b5060006114d28383611a2c565b905060006114e287878488611815565b90506114f0818b8b8b611853565b6114fa87876116e9565b841015611539576040805162461bcd60e51b81526020600482015260086024820152671393d517d0d3d39560c21b604482015290519081900360640190fd5b8482141561157a576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6040517ff62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca90600090a16115ab61190b565b5060029050600c5460ff1660028111156115c157fe5b14156115f4576115df610ce5600954436116a790919063ffffffff16565b600b55600c805460ff1916600117905561161d565b61160c610d1b600954436116a790919063ffffffff16565b600a55600c805460ff191660021790555b505043600955505050505050565b60006001600c5460ff16600281111561164057fe5b141561164f5750600a5461077d565b6002600c5460ff16600281111561166257fe5b1415610ec15750600b5461077d565b6003546001600160a01b031681565b6001818154811061168d57fe5b6000918252602090912001546001600160a01b0316905081565b60006116e08383604051806040016040528060148152602001737375627472616374696f6e206f766572666c6f7760601b815250611b94565b90505b92915050565b6000828201838110156116e0576040805162461bcd60e51b81526020600482015260116024820152706164646974696f6e206f766572666c6f7760781b604482015290519081900360640190fd5b60008061176083600260200201518914611752576001611755565b60005b60ff16876001610bb7565b9050600061178a8460036020020151891461177c57600161177f565b60005b60ff16886002610bb7565b90506117cf6117a46001600160401b038816896000610bb7565b865160208089015190880151604089015160608a01516117ca9493929190899089611bee565b611a2c565b9998505050505050505050565b81518151602080850151604086015160009461180b9490936117ca938c3593918d013592918c91908c90611bee565b9695505050505050565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b611893838380806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250859250889150611c429050565b600d5414604051806040016040528060088152602001672124a9afa82922ab60c11b815250906119045760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b5050505050565b6001600c5460ff16600281111561191e57fe5b14156119315761192c6119b8565b611939565b61193961193b565b565b6004805460085460075460408051637d3c01f360e11b81526001600160a01b039384169581019590955290821660248501525191169163fa7803e691604480830192600092919082900301818387803b15801561199757600080fd5b505af11580156119ab573d6000803e3d6000fd5b5050505061193933611d10565b6004805460075460085460408051637d3c01f360e11b81526001600160a01b039384169581019590955290821660248501525191169163fa7803e691604480830192600092919082900301818387803b15801561199757600080fd5b600081831015611a255750816116e3565b50806116e3565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b8251600090600019016060816001600160401b0381118015611a7957600080fd5b50604051908082528060200260200182016040528015611aa3578160200160208202803683370190505b5090506000611ab28584611d90565b90506000869050611aed81838a600081518110611acb57fe5b60200260200101518b600181518110611ae057fe5b6020026020010151611815565b83600081518110611afa57fe5b6020908102919091010152611b0f81836116e9565b9050611b1b8685611dae565b915060015b84811015611b7b57611b5082848b8481518110611b3957fe5b60200260200101518c8560010181518110611ae057fe5b848281518110611b5c57fe5b6020908102919091010152611b7182846116e9565b9150600101611b20565b50611b8583611dc1565b600d5550929695505050505050565b60008184841115611be65760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b505050900390565b60408051602080820199909952808201979097526060870195909552608086019390935260a085019190915260c084015260e080840191909152815180840390910181526101009092019052805191012090565b8251600090610100811115611c5657600080fd5b8260005b82811015611d065760028606611cb357868181518110611c7657fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150611cf8565b81878281518110611cc057fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101611c5a565b5095945050505050565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff1615611d835760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b50806001600160a01b0316ff5b6000818381611d9b57fe5b06828481611da557fe5b04019392505050565b6000818381611db957fe5b049392505050565b6000815b600181511115611ef85760606002825160010181611ddf57fe5b046001600160401b0381118015611df557600080fd5b50604051908082528060200260200182016040528015611e1f578160200160208202803683370190505b50905060005b8151811015611ef0578251816002026001011015611eb857828160020281518110611e4c57fe5b6020026020010151838260020260010181518110611e6657fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611ea757fe5b602002602001018181525050611ee8565b828160020281518110611ec757fe5b6020026020010151828281518110611edb57fe5b6020026020010181815250505b600101611e25565b509050611dc5565b80600081518110611f0557fe5b6020026020010151915050919050565b828054828255906000526020600020908101928215611f68579160200282015b82811115611f685781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190611f35565b50611f74929150611fb4565b5090565b60405180604001604052806002906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b5b80821115611f745780546001600160a01b0319168155600101611fb556fea2646970667358221220744e12cebdb6a2f264c4bea7b429ce72ebd02faf90899d5b31ad26512ae0ab9464736f6c634300060c0033"

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

// CreateChallenge is a paid mutator transaction binding the contract method 0x820eb646.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, uint256 _maxSequencerBatchCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactor) CreateChallenge(opts *bind.TransactOpts, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _maxSequencerBatchCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.contract.Transact(opts, "createChallenge", _resultReceiver, _executionHash, _maxMessageCount, _maxSequencerBatchCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x820eb646.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, uint256 _maxSequencerBatchCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns(address)
func (_ChallengeFactory *ChallengeFactorySession) CreateChallenge(_resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _maxSequencerBatchCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _resultReceiver, _executionHash, _maxMessageCount, _maxSequencerBatchCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x820eb646.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, uint256 _maxSequencerBatchCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactorSession) CreateChallenge(_resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _maxSequencerBatchCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _resultReceiver, _executionHash, _maxMessageCount, _maxSequencerBatchCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// IOneStepProofABI is the input ABI used to generate the binding from.
const IOneStepProofABI = "[{\"inputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerBridge\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"delayedBridge\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"initialMessagesAndBatchesRead\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint256[2]\",\"name\":\"afterMessagesAndBatchesRead\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[4]\",\"name\":\"fields\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerBridge\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"delayedBridge\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"initialMessagesAndBatchesRead\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStepDebug\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"startMachine\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"afterMachine\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// ExecuteStep is a free data retrieval call binding the contract method 0x92922de5.
//
// Solidity: function executeStep(address sequencerBridge, address delayedBridge, uint256[2] initialMessagesAndBatchesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256[2] afterMessagesAndBatchesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofCaller) ExecuteStep(opts *bind.CallOpts, sequencerBridge common.Address, delayedBridge common.Address, initialMessagesAndBatchesRead [2]*big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas                         uint64
	AfterMessagesAndBatchesRead [2]*big.Int
	Fields                      [4][32]byte
}, error) {
	var out []interface{}
	err := _IOneStepProof.contract.Call(opts, &out, "executeStep", sequencerBridge, delayedBridge, initialMessagesAndBatchesRead, accs, proof, bproof)

	outstruct := new(struct {
		Gas                         uint64
		AfterMessagesAndBatchesRead [2]*big.Int
		Fields                      [4][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gas = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AfterMessagesAndBatchesRead = *abi.ConvertType(out[1], new([2]*big.Int)).(*[2]*big.Int)
	outstruct.Fields = *abi.ConvertType(out[2], new([4][32]byte)).(*[4][32]byte)

	return *outstruct, err

}

// ExecuteStep is a free data retrieval call binding the contract method 0x92922de5.
//
// Solidity: function executeStep(address sequencerBridge, address delayedBridge, uint256[2] initialMessagesAndBatchesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256[2] afterMessagesAndBatchesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofSession) ExecuteStep(sequencerBridge common.Address, delayedBridge common.Address, initialMessagesAndBatchesRead [2]*big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas                         uint64
	AfterMessagesAndBatchesRead [2]*big.Int
	Fields                      [4][32]byte
}, error) {
	return _IOneStepProof.Contract.ExecuteStep(&_IOneStepProof.CallOpts, sequencerBridge, delayedBridge, initialMessagesAndBatchesRead, accs, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x92922de5.
//
// Solidity: function executeStep(address sequencerBridge, address delayedBridge, uint256[2] initialMessagesAndBatchesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256[2] afterMessagesAndBatchesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofCallerSession) ExecuteStep(sequencerBridge common.Address, delayedBridge common.Address, initialMessagesAndBatchesRead [2]*big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas                         uint64
	AfterMessagesAndBatchesRead [2]*big.Int
	Fields                      [4][32]byte
}, error) {
	return _IOneStepProof.Contract.ExecuteStep(&_IOneStepProof.CallOpts, sequencerBridge, delayedBridge, initialMessagesAndBatchesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x1d1bc075.
//
// Solidity: function executeStepDebug(address sequencerBridge, address delayedBridge, uint256[2] initialMessagesAndBatchesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofCaller) ExecuteStepDebug(opts *bind.CallOpts, sequencerBridge common.Address, delayedBridge common.Address, initialMessagesAndBatchesRead [2]*big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	var out []interface{}
	err := _IOneStepProof.contract.Call(opts, &out, "executeStepDebug", sequencerBridge, delayedBridge, initialMessagesAndBatchesRead, accs, proof, bproof)

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

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x1d1bc075.
//
// Solidity: function executeStepDebug(address sequencerBridge, address delayedBridge, uint256[2] initialMessagesAndBatchesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofSession) ExecuteStepDebug(sequencerBridge common.Address, delayedBridge common.Address, initialMessagesAndBatchesRead [2]*big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _IOneStepProof.Contract.ExecuteStepDebug(&_IOneStepProof.CallOpts, sequencerBridge, delayedBridge, initialMessagesAndBatchesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x1d1bc075.
//
// Solidity: function executeStepDebug(address sequencerBridge, address delayedBridge, uint256[2] initialMessagesAndBatchesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofCallerSession) ExecuteStepDebug(sequencerBridge common.Address, delayedBridge common.Address, initialMessagesAndBatchesRead [2]*big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _IOneStepProof.Contract.ExecuteStepDebug(&_IOneStepProof.CallOpts, sequencerBridge, delayedBridge, initialMessagesAndBatchesRead, accs, proof, bproof)
}
