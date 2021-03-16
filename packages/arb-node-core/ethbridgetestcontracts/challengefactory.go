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
const ChallengeFactoryABI = "[{\"inputs\":[{\"internalType\":\"contractIOneStepProof[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"challengeTemplate\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_resultReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_executionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_asserterTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengerTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"executors\",\"outputs\":[{\"internalType\":\"contractIOneStepProof\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeFactoryBin is the compiled bytecode used for deploying new contracts.
var ChallengeFactoryBin = "0x608060405234801561001057600080fd5b5060405161253d38038061253d8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b505050509050016040525050506040516100cb90610122565b604051809103906000f0801580156100e7573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055805161011b90600190602084019061012f565b50506101b3565b611f308061060d83390190565b828054828255906000526020600020908101928215610184579160200282015b8281111561018457825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061014f565b50610190929150610194565b5090565b5b808211156101905780546001600160a01b0319168155600101610195565b61044b806101c26000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806356a44dbb14610046578063777367c9146100b9578063f97a05df146100c1575b600080fd5b61009d600480360361010081101561005d57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135821691608082013581169160a08101359160c08201359160e00135166100de565b604080516001600160a01b039092168252519081900360200190f35b61009d610211565b61009d600480360360208110156100d757600080fd5b5035610220565b6000805481906100f6906001600160a01b0316610247565b9050806001600160a01b03166332f8c24f60018c8c8c8c8c8c8c8c6040518a63ffffffff1660e01b815260040180806020018a6001600160a01b03168152602001898152602001888152602001876001600160a01b03168152602001866001600160a01b03168152602001858152602001848152602001836001600160a01b0316815260200182810382528b81815481526020019150805480156101c357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116101a5575b50509a5050505050505050505050600060405180830381600087803b1580156101eb57600080fd5b505af11580156101ff573d6000803e3d6000fd5b50929c9b505050505050505050505050565b6000546001600160a01b031681565b6001818154811061022d57fe5b6000918252602090912001546001600160a01b0316905081565b6000816001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b15801561028257600080fd5b505afa158015610296573d6000803e3d6000fd5b505050506040513d60208110156102ac57600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b60208201529061035a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561031f578181015183820152602001610307565b50505050905090810190601f16801561034c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061036d826001600160a01b0316610373565b92915050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116610410576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b91905056fea264697066735822122053d3f8ee5649c98dc1d3c104ca89136cec97d129a4d600d6b2237176a48b8f3e64736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff19166001179055611f038061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100db5760003560e01c806332f8c24f146100e057806341e8510c1461018a578063534db0e2146101a45780636f791d29146101c857806370dea79a146101e4578063843d5a5c146101ec5780638a8cd218146101f45780638b299903146101fc5780638e7b84c514610225578063925f9a96146103085780639a9e4f4414610310578063bb4af0b114610318578063deda411514610320578063e08f819e146103ac578063e78cea9214610583578063e87e35891461058b578063f97a05df14610593575b600080fd5b61018860048036036101208110156100f757600080fd5b810190602081018135600160201b81111561011157600080fd5b82018360208201111561012357600080fd5b803590602001918460208302840111600160201b8311171561014457600080fd5b91935091506001600160a01b0381358116916020810135916040820135916060810135821691608082013581169160a08101359160c08201359160e00135166105b0565b005b610192610728565b60408051918252519081900360200190f35b6101ac61072e565b604080516001600160a01b039092168252519081900360200190f35b6101d061073d565b604080519115158252519081900360200190f35b610188610747565b610192610863565b6101ac610869565b6102046108f8565b6040518082600281111561021457fe5b815260200191505060405180910390f35b610188600480360361010081101561023c57600080fd5b810190602081018135600160201b81111561025657600080fd5b82018360208201111561026857600080fd5b803590602001918460208302840111600160201b8311171561028957600080fd5b9193909282359260208101359260408201359260608301359260808101359260a082013592909160e081019060c00135600160201b8111156102ca57600080fd5b8201836020820111156102dc57600080fd5b803590602001918460208302840111600160201b831117156102fd57600080fd5b509092509050610901565b610192610da7565b610192610dad565b6101ac610db3565b610188600480360360e081101561033657600080fd5b810190602081018135600160201b81111561035057600080fd5b82018360208201111561036257600080fd5b803590602001918460208302840111600160201b8311171561038357600080fd5b919350915080359060208101359060408101359060608101359060808101359060a00135610dc2565b61018860048036036101c08110156103c357600080fd5b810190602081018135600160201b8111156103dd57600080fd5b8201836020820111156103ef57600080fd5b803590602001918460208302840111600160201b8311171561041057600080fd5b6040805160608181018352949693958335956020850135959385013594818101359460808201359460a08301359460c084013594929390926101408201929160e00190600390839083908082843760009201919091525091949392602081019250359050600160201b81111561048557600080fd5b82018360208201111561049757600080fd5b803590602001918460018302840111600160201b831117156104b857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561050a57600080fd5b82018360208201111561051c57600080fd5b803590602001918460018302840111600160201b8311171561053d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff16915061103d9050565b6101ac61157a565b610192611589565b6101ac600480360360208110156105a957600080fd5b50356115cf565b6000600a5460ff1660028111156105c357fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b815250906106755760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561063a578181015183820152602001610622565b50505050905090810190601f1680156106675780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061068260018b8b611e2d565b50600380546001600160a01b038a81166001600160a01b0319928316179092556004889055600580548884169083161790556006805487841690831617905560088590556009849055600a8054600260ff199091168117909155600b8a90554360075580549091169183169190911790556040517f7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de90600090a150505050505050505050565b60095481565b6006546001600160a01b031681565b60005460ff165b90565b600061075e600754436115f690919063ffffffff16565b9050610768611589565b81116040518060400160405280601081526020016f54494d454f55545f444541444c494e4560801b815250906107df5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b506001600a5460ff1660028111156107f357fe5b141561082f576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a161082a611638565b610860565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a16108606116b6565b50565b600b5481565b60006001600a5460ff16600281111561087e57fe5b141561089657506005546001600160a01b0316610744565b6002600a5460ff1660028111156108a957fe5b14156108c157506006546001600160a01b0316610744565b6040805162461bcd60e51b81526020600482015260076024820152662727afaa2aa92760c91b604482015290519081900360640190fd5b600a5460ff1681565b610909610869565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b8152509061098c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b50610995611589565b6007546109a39043906115f6565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610a165760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b50600082826000198101818110610a2957fe5b9050602002013514610a765760018611610a76576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b604482015290519081900360640190fd5b610a8286610190611711565b6001018114610ac4576040805162461bcd60e51b815260206004820152600960248201526810d55517d0d3d5539560ba1b604482015290519081900360640190fd5b8482826000198101818110610ad557fe5b905060200201351415610b1a576040805162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b604482015290519081900360640190fd5b610b248484611729565b82826000818110610b3157fe5b9050602002013514610b7f576040805162461bcd60e51b81526020600482015260126024820152717365676d656e74207072652d6669656c647360701b604482015290519081900360640190fd5b600082828281610b8b57fe5b905060200201351415610bd9576040805162461bcd60e51b8152602060048201526011602482015270155394915050d21050931157d4d5105495607a1b604482015290519081900360640190fd5b610be38787611755565b8410610c2f576040805162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840e6cacedacadce840d8cadccee8d60531b604482015290519081900360640190fd5b6000610c50888885856000818110610c4357fe5b90506020020135896117a3565b9050610c5e818c8c8c6117e1565b610c9e8383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508c92508b91506118999050565b50600b547f0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d8989868660405180858152602001848152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f191690920182900397509095505050505050a2506002600a5460ff166002811115610d2957fe5b1415610d6557610d50610d47600754436115f690919063ffffffff16565b600954906115f6565b600955600a805460ff19166001179055610d97565b610d86610d7d600754436115f690919063ffffffff16565b600854906115f6565b600855600a805460ff191660021790555b5050436007555050505050505050565b60075481565b60085481565b6005546001600160a01b031681565b610dca610869565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b81525090610e4d5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b50610e56611589565b600754610e649043906115f6565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610ed75760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b506000610ee48383611729565b90506000610ef4878784886117a3565b9050610f02818b8b8b6117e1565b610f0c8787611755565b841015610f4b576040805162461bcd60e51b81526020600482015260086024820152671393d517d0d3d39560c21b604482015290519081900360640190fd5b84821415610f8c576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6040517ff62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca90600090a1610fbd6119d5565b5060029050600a5460ff166002811115610fd357fe5b141561100657610ff1610d47600754436115f690919063ffffffff16565b600955600a805460ff1916600117905561102f565b61101e610d7d600754436115f690919063ffffffff16565b600855600a805460ff191660021790555b505043600755505050505050565b611045610869565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b815250906110c85760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b506110d1611589565b6007546110df9043906115f6565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906111525760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b506000806000611160611e90565b60018560ff168154811061117057fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b0316639d16dd04600260009054906101000a90046001600160a01b03168d60405180604001604052808f81526020018e8152508b8b6040518663ffffffff1660e01b815260040180866001600160a01b0316815260200185815260200184600260200280838360005b838110156112165781810151838201526020016111fe565b505050509050018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561125a578181015183820152602001611242565b50505050905090810190601f1680156112875780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156112ba5781810151838201526020016112a2565b50505050905090810190601f1680156112e75780820380516001836020036101000a031916815260200191505b5097505050505050505060c06040518083038186803b15801561130957600080fd5b505afa15801561131d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c081101561134257600080fd5b50805160208201516004549195509350604090910191508211156113a1576040805162461bcd60e51b8152602060048201526011602482015270544f4f5f4d414e595f4d4553534147455360781b604482015290519081900360640190fd5b6113ab8e8e611755565b8851106113ea576040805162461bcd60e51b815260206004820152600860248201526713d4d417d0d3d39560c21b604482015290519081900360640190fd5b6113f48e8e611755565b6114106001600160401b0385168a60005b602002015190611755565b101561144f576040805162461bcd60e51b815260206004820152600960248201526813d4d417d4d213d49560ba1b604482015290519081900360640190fd5b61145d8a8a8a868686611a03565b8c141561149d576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6114b58e8e6114af8e8e8e8e88611a84565b8f6117a3565b93505050506114c6818f8f8f6117e1565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a16114f76119d5565b506002600a5460ff16600281111561150b57fe5b141561153e57611529610d47600754436115f690919063ffffffff16565b600955600a805460ff19166001179055611567565b611556610d7d600754436115f690919063ffffffff16565b600855600a805460ff191660021790555b5050436007555050505050505050505050565b6002546001600160a01b031681565b60006001600a5460ff16600281111561159e57fe5b14156115ad5750600854610744565b6002600a5460ff1660028111156115c057fe5b14156108c15750600954610744565b600181815481106115dc57fe5b6000918252602090912001546001600160a01b0316905081565b600061162f8383604051806040016040528060148152602001737375627472616374696f6e206f766572666c6f7760601b815250611ab5565b90505b92915050565b60035460065460055460408051637d3c01f360e11b81526001600160a01b039384166004820152918316602483015251919092169163fa7803e691604480830192600092919082900301818387803b15801561169357600080fd5b505af11580156116a7573d6000803e3d6000fd5b505050506116b433611b0f565b565b60035460055460065460408051637d3c01f360e11b81526001600160a01b039384166004820152918316602483015251919092169163fa7803e691604480830192600092919082900301818387803b15801561169357600080fd5b600081831015611722575081611632565b5080611632565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008282018381101561162f576040805162461bcd60e51b81526020600482015260116024820152706164646974696f6e206f766572666c6f7760781b604482015290519081900360640190fd5b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b611821838380806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250859250889150611b8f9050565b600b5414604051806040016040528060088152602001672124a9afa82922ab60c11b815250906118925760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b5050505050565b8251600090600019016060816001600160401b03811180156118ba57600080fd5b506040519080825280602002602001820160405280156118e4578160200160208202803683370190505b50905060006118f38584611c5d565b9050600086905061192e81838a60008151811061190c57fe5b60200260200101518b60018151811061192157fe5b60200260200101516117a3565b8360008151811061193b57fe5b60209081029190910101526119508183611755565b905061195c8685611c7b565b915060015b848110156119bc5761199182848b848151811061197a57fe5b60200260200101518c856001018151811061192157fe5b84828151811061199d57fe5b60209081029190910101526119b28284611755565b9150600101611961565b506119c683611c8e565b600b5550929695505050505050565b6001600a5460ff1660028111156119e857fe5b14156119fb576119f66116b6565b6116b4565b6116b4611638565b6000611a79611a1c6001600160401b0386168784611405565b60208401516040850151611a74918791611a4b8d8214611a3d576001611a40565b60005b60ff168c6001611405565b6060890151611a6f8e8214611a61576001611a64565b60005b60ff168e6002611405565b611de2565b611729565b979650505050505050565b8151815160208401516040850151600093611aab939092611a74928b92918b918b90611de2565b9695505050505050565b60008184841115611b075760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b505050900390565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff1615611b825760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561063a578181015183820152602001610622565b50806001600160a01b0316ff5b8251600090610100811115611ba357600080fd5b8260005b82811015611c535760028606611c0057868181518110611bc357fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150611c45565b81878281518110611c0d57fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101611ba7565b5095945050505050565b6000818381611c6857fe5b06828481611c7257fe5b04019392505050565b6000818381611c8657fe5b049392505050565b6000815b600181511115611dc55760606002825160010181611cac57fe5b046001600160401b0381118015611cc257600080fd5b50604051908082528060200260200182016040528015611cec578160200160208202803683370190505b50905060005b8151811015611dbd578251816002026001011015611d8557828160020281518110611d1957fe5b6020026020010151838260020260010181518110611d3357fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611d7457fe5b602002602001018181525050611db5565b828160020281518110611d9457fe5b6020026020010151828281518110611da857fe5b6020026020010181815250505b600101611cf2565b509050611c92565b80600081518110611dd257fe5b6020026020010151915050919050565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b828054828255906000526020600020908101928215611e80579160200282015b82811115611e805781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190611e4d565b50611e8c929150611eae565b5090565b60405180608001604052806004906020820280368337509192915050565b5b80821115611e8c5780546001600160a01b0319168155600101611eaf56fea26469706673582212200d4129416fab2e0e9f36a149767a515d1c3e131eb4bc36b5bcb666c682b1304b64736f6c634300060c0033"

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

// CreateChallenge is a paid mutator transaction binding the contract method 0x56a44dbb.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _bridge) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactor) CreateChallenge(opts *bind.TransactOpts, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _bridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.contract.Transact(opts, "createChallenge", _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _bridge)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x56a44dbb.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _bridge) returns(address)
func (_ChallengeFactory *ChallengeFactorySession) CreateChallenge(_resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _bridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _bridge)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x56a44dbb.
//
// Solidity: function createChallenge(address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _bridge) returns(address)
func (_ChallengeFactory *ChallengeFactoryTransactorSession) CreateChallenge(_resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _bridge common.Address) (*types.Transaction, error) {
	return _ChallengeFactory.Contract.CreateChallenge(&_ChallengeFactory.TransactOpts, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _bridge)
}

// IOneStepProofABI is the input ABI used to generate the binding from.
const IOneStepProofABI = "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"totalMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[4]\",\"name\":\"fields\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"accs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStepDebug\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"startMachine\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"afterMachine\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// ExecuteStep is a free data retrieval call binding the contract method 0x9d16dd04.
//
// Solidity: function executeStep(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofCaller) ExecuteStep(opts *bind.CallOpts, bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	var out []interface{}
	err := _IOneStepProof.contract.Call(opts, &out, "executeStep", bridge, initialMessagesRead, accs, proof, bproof)

	outstruct := new(struct {
		Gas               uint64
		TotalMessagesRead *big.Int
		Fields            [4][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gas = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.TotalMessagesRead = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fields = *abi.ConvertType(out[2], new([4][32]byte)).(*[4][32]byte)

	return *outstruct, err

}

// ExecuteStep is a free data retrieval call binding the contract method 0x9d16dd04.
//
// Solidity: function executeStep(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofSession) ExecuteStep(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _IOneStepProof.Contract.ExecuteStep(&_IOneStepProof.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x9d16dd04.
//
// Solidity: function executeStep(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields)
func (_IOneStepProof *IOneStepProofCallerSession) ExecuteStep(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	Gas               uint64
	TotalMessagesRead *big.Int
	Fields            [4][32]byte
}, error) {
	return _IOneStepProof.Contract.ExecuteStep(&_IOneStepProof.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x2ccebb7a.
//
// Solidity: function executeStepDebug(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofCaller) ExecuteStepDebug(opts *bind.CallOpts, bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	var out []interface{}
	err := _IOneStepProof.contract.Call(opts, &out, "executeStepDebug", bridge, initialMessagesRead, accs, proof, bproof)

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

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x2ccebb7a.
//
// Solidity: function executeStepDebug(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofSession) ExecuteStepDebug(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _IOneStepProof.Contract.ExecuteStepDebug(&_IOneStepProof.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}

// ExecuteStepDebug is a free data retrieval call binding the contract method 0x2ccebb7a.
//
// Solidity: function executeStepDebug(address bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) view returns(string startMachine, string afterMachine)
func (_IOneStepProof *IOneStepProofCallerSession) ExecuteStepDebug(bridge common.Address, initialMessagesRead *big.Int, accs [2][32]byte, proof []byte, bproof []byte) (struct {
	StartMachine string
	AfterMachine string
}, error) {
	return _IOneStepProof.Contract.ExecuteStepDebug(&_IOneStepProof.CallOpts, bridge, initialMessagesRead, accs, proof, bproof)
}
