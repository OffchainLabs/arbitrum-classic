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
var ChallengeFactoryBin = "0x608060405234801561001057600080fd5b506040516128c83803806128c88339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b505050509050016040525050506040516100cb90610122565b604051809103906000f0801580156100e7573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055805161011b90600190602084019061012f565b50506101bb565b6121d2806106f683390190565b828054828255906000526020600020908101928215610184579160200282015b8281111561018457825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061014f565b50610190929150610194565b5090565b6101b891905b808211156101905780546001600160a01b031916815560010161019a565b90565b61052c806101ca6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063777367c9146100465780638ecaab111461006a578063f97a05df146100cc575b600080fd5b61004e6100e9565b604080516001600160a01b039092168252519081900360200190f35b61004e600480360361012081101561008157600080fd5b506001600160a01b0381358116916020810135916040820135916060810135821691608082013581169160a08101359160c08201359160e081013582169161010090910135166100f8565b61004e600480360360208110156100e257600080fd5b503561026a565b6000546001600160a01b031681565b600080548190610110906001600160a01b0316610291565b9050806001600160a01b031663e0d42b8e60018d8d8d8d8d8d8d8d8d6040518b63ffffffff1660e01b815260040180806020018b6001600160a01b03166001600160a01b031681526020018a8152602001898152602001886001600160a01b03166001600160a01b03168152602001876001600160a01b03166001600160a01b03168152602001868152602001858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182810382528c818154815260200191508054801561021a57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116101fc575b50509b505050505050505050505050600060405180830381600087803b15801561024357600080fd5b505af1158015610257573d6000803e3d6000fd5b50929d9c50505050505050505050505050565b6001818154811061027757fe5b6000918252602090912001546001600160a01b0316905081565b6000816102a6816001600160a01b031661044f565b604051806040016040528060188152602001772727afa1a7a72a2920a1aa2fa1a627a722afa6a0a9aa22a960411b815250906103605760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561032557818101518382015260200161030d565b50505050905090810190601f1680156103525780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b15801561039a57600080fd5b505afa1580156103ae573d6000803e3d6000fd5b505050506040513d60208110156103c457600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b6020820152906104355760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561032557818101518382015260200161030d565b50610448836001600160a01b0316610459565b9392505050565b803b15155b919050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116610454576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fdfea26469706673582212207935c5c813f4a3a4fd36132242c06161c407735cf55fc15f220c26a0751e4f3d64736f6c634300060b0033608060405234801561001057600080fd5b506000805460ff191660011790556121a58061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063925f9a9611610092578063925f9a9614610285578063959792011461028d5780639a9e4f441461045e578063a3c4470514610466578063bb4af0b114610483578063deda41151461048b578063e0d42b8e14610517578063e87e3589146105ca578063f97a05df146105d2576100f5565b806214ebe7146100fa57806341e8510c14610104578063534db0e21461011e5780636f791d291461014257806370dea79a1461015e578063843d5a5c146101665780638a8cd2181461016e5780638b299903146101765780638e7b84c5146101a2575b600080fd5b6101026105ef565b005b61010c61064c565b60408051918252519081900360200190f35b610126610652565b604080516001600160a01b039092168252519081900360200190f35b61014a610661565b604080519115158252519081900360200190f35b61010261066b565b61010c6107c4565b6101266107ca565b61017e610859565b6040518082600281111561018e57fe5b60ff16815260200191505060405180910390f35b61010260048036036101008110156101b957600080fd5b810190602081018135600160201b8111156101d357600080fd5b8201836020820111156101e557600080fd5b803590602001918460208302840111600160201b8311171561020657600080fd5b9193909282359260208101359260408201359260608301359260808101359260a082013592909160e081019060c00135600160201b81111561024757600080fd5b82018360208201111561025957600080fd5b803590602001918460208302840111600160201b8311171561027a57600080fd5b509092509050610862565b61010c610e1e565b61010260048036036101c08110156102a457600080fd5b810190602081018135600160201b8111156102be57600080fd5b8201836020820111156102d057600080fd5b803590602001918460208302840111600160201b831117156102f157600080fd5b6040805160608181018352949693958335956020850135959385013594818101359460808201359460a08301949193919261014081019260e090910190600390839083908082843760009201919091525091949392602081019250359050600160201b81111561036057600080fd5b82018360208201111561037257600080fd5b803590602001918460018302840111600160201b8311171561039357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103e557600080fd5b8201836020820111156103f757600080fd5b803590602001918460018302840111600160201b8311171561041857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff169150610e249050565b61010c6113cb565b6101266004803603602081101561047c57600080fd5b50356113d1565b6101266113ee565b610102600480360360e08110156104a157600080fd5b810190602081018135600160201b8111156104bb57600080fd5b8201836020820111156104cd57600080fd5b803590602001918460208302840111600160201b831117156104ee57600080fd5b919350915080359060208101359060408101359060608101359060808101359060a001356113fd565b610102600480360361014081101561052e57600080fd5b810190602081018135600160201b81111561054857600080fd5b82018360208201111561055a57600080fd5b803590602001918460208302840111600160201b8311171561057b57600080fd5b91935091506001600160a01b0381358116916020810135916040820135916060810135821691608082013581169160a08101359160c08201359160e081013582169161010090910135166116f5565b61010c611846565b610126600480360360208110156105e857600080fd5b503561188c565b6004546001600160a01b03163314610641576040805162461bcd60e51b815260206004820152601060248201526f2727aa2fa922a9afa922a1a2a4ab22a960811b604482015290519081900360640190fd5b61064a336118b3565b565b600a5481565b6007546001600160a01b031681565b60005460ff165b90565b60006106826008544361193390919063ffffffff16565b905061068c611846565b81116040518060400160405280601081526020016f54494d454f55545f444541444c494e4560801b815250906107405760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156107055781810151838201526020016106ed565b50505050905090810190601f1680156107325780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506001600b5460ff16600281111561075457fe5b1415610790576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a161078b611995565b6107c1565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a16107c1611a12565b50565b600c5481565b60006001600b5460ff1660028111156107df57fe5b14156107f757506006546001600160a01b0316610668565b6002600b5460ff16600281111561080a57fe5b141561082257506007546001600160a01b0316610668565b6040805162461bcd60e51b81526020600482015260076024820152662727afaa2aa92760c91b604482015290519081900360640190fd5b600b5460ff1681565b61086a6107ca565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b815250906108ed5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b506108f6611846565b60085461090a90439063ffffffff61193316565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b8152509061097d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b5060008282600019810181811061099057fe5b90506020020135146109dd57600186116109dd576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b604482015290519081900360640190fd5b6000600460009054906101000a90046001600160a01b03166001600160a01b031663dc72a33b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a2d57600080fd5b505afa158015610a41573d6000803e3d6000fd5b505050506040513d6020811015610a5757600080fd5b50519050610a658782611a6e565b6001018214610aa7576040805162461bcd60e51b815260206004820152600960248201526810d55517d0d3d5539560ba1b604482015290519081900360640190fd5b8583836000198101818110610ab857fe5b905060200201351415610afd576040805162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b604482015290519081900360640190fd5b610b078585611a86565b83836000818110610b1457fe5b9050602002013514610b62576040805162461bcd60e51b81526020600482015260126024820152717365676d656e74207072652d6669656c647360701b604482015290519081900360640190fd5b600083838281610b6e57fe5b905060200201351415610bbc576040805162461bcd60e51b8152602060048201526011602482015270155394915050d21050931157d4d5105495607a1b604482015290519081900360640190fd5b610bcc888863ffffffff611ab216565b8510610c18576040805162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840e6cacedacadce840d8cadccee8d60531b604482015290519081900360640190fd5b6000610c39898986866000818110610c2c57fe5b905060200201358a611b13565b9050610c4a600c54828e8e8e611b51565b604051806040016040528060088152602001672124a9afa82922ab60c11b81525090610cb75760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b506000610cfa8585808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508d9150611b9f9050565b905080600c81905550807f0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d8b8b888860405180858152602001848152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f191690920182900397509095505050505050a25060029150610d839050565b600b5460ff166002811115610d9457fe5b1415610dd657610dc1610db26008544361193390919063ffffffff16565b600a549063ffffffff61193316565b600a55600b805460ff19166001179055610e0e565b610dfd610dee6008544361193390919063ffffffff16565b6009549063ffffffff61193316565b600955600b805460ff191660021790555b5050436008555050505050505050565b60085481565b610e2c6107ca565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b81525090610eaf5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b50610eb8611846565b600854610ecc90439063ffffffff61193316565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610f3f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b506000806000610f4d612082565b60018560ff1681548110610f5d57fe5b6000918252602090912001546040516323eed0eb60e11b81526001600160a01b03909116906347dda1d6906002908d908d908c908c90600481019060440186825b81546001600160a01b03168152600190910190602001808311610f9e57505085815260200184604080828437600081840152601f19601f8201169050808301925050508060200180602001838103835285818151815260200191508051906020019080838360005b8381101561101e578181015183820152602001611006565b50505050905090810190601f16801561104b5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561107e578181015183820152602001611066565b50505050905090810190601f1680156110ab5780820380516001836020036101000a031916815260200191505b5097505050505050505060c06040518083038186803b1580156110cd57600080fd5b505afa1580156110e1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c081101561110657600080fd5b5080516020820151600554919550935060409091019150821115611165576040805162461bcd60e51b8152602060048201526011602482015270544f4f5f4d414e595f4d4553534147455360781b604482015290519081900360640190fd5b6111758d8d63ffffffff611ab216565b8851106111b4576040805162461bcd60e51b815260206004820152600860248201526713d4d417d0d3d39560c21b604482015290519081900360640190fd5b6111c48d8d63ffffffff611ab216565b6111e767ffffffffffffffff85168a60005b60200201519063ffffffff611ab216565b1015611226576040805162461bcd60e51b815260206004820152600960248201526813d4d417d4d213d49560ba1b604482015290519081900360640190fd5b611239893560208b01358a868686611ce5565b8b1415611279576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6112968d8d6112908d8d3560208f01358e88611d84565b8e611b13565b93505050506112aa600c54828f8f8f611b51565b604051806040016040528060088152602001672124a9afa82922ab60c11b815250906113175760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b506040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1611349611db5565b506002600b5460ff16600281111561135d57fe5b14156113905761137b610db26008544361193390919063ffffffff16565b600a55600b805460ff191660011790556113b9565b6113a8610dee6008544361193390919063ffffffff16565b600955600b805460ff191660021790555b50504360085550505050505050505050565b60095481565b600281600281106113de57fe5b01546001600160a01b0316905081565b6006546001600160a01b031681565b6114056107ca565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b815250906114885760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b50611491611846565b6008546114a590439063ffffffff61193316565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906115185760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b5060006115258383611a86565b9050600061153587878488611b13565b9050611546600c54828c8c8c611b51565b604051806040016040528060088152602001672124a9afa82922ab60c11b815250906115b35760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b506115c4878763ffffffff611ab216565b841015611603576040805162461bcd60e51b81526020600482015260086024820152671393d517d0d3d39560c21b604482015290519081900360640190fd5b84821415611644576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6040517ff62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca90600090a1611675611db5565b5060029050600b5460ff16600281111561168b57fe5b14156116be576116a9610db26008544361193390919063ffffffff16565b600a55600b805460ff191660011790556116e7565b6116d6610dee6008544361193390919063ffffffff16565b600955600b805460ff191660021790555b505043600855505050505050565b6000600b5460ff16600281111561170857fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b8152509061177d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b5061178a60018c8c6120a0565b50600480546001600160a01b03199081166001600160a01b038c8116919091179092556005899055600680548216898416179055600780549091168783161790556009859055600a849055600b805460ff19166002908117909155600c8a905543600855604080518082019091528483168152918316602083015261180f9181612103565b506040517f7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de90600090a15050505050505050505050565b60006001600b5460ff16600281111561185b57fe5b141561186a5750600954610668565b6002600b5460ff16600281111561187d57fe5b14156108225750600a54610668565b6001818154811061189957fe5b6000918252602090912001546001600160a01b0316905081565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff16156119265760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107055781810151838201526020016106ed565b50806001600160a01b0316ff5b60008282111561198a576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b508082035b92915050565b6004805460075460065460408051637d3c01f360e11b81526001600160a01b039384169581019590955290821660248501525191169163fa7803e691604480830192600092919082900301818387803b1580156119f157600080fd5b505af1158015611a05573d6000803e3d6000fd5b5050505061064a336118b3565b6004805460065460075460408051637d3c01f360e11b81526001600160a01b039384169581019590955290821660248501525191169163fa7803e691604480830192600092919082900301818387803b1580156119f157600080fd5b600081831015611a7f57508161198f565b508061198f565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600082820183811015611b0c576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b6000611b93848480806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250869250899150611de39050565b90951495945050505050565b82516000906000190160608167ffffffffffffffff81118015611bc157600080fd5b50604051908082528060200260200182016040528015611beb578160200160208202803683370190505b5090506000611bfa8584611eb1565b90506000869050611c3581838a600081518110611c1357fe5b60200260200101518b600181518110611c2857fe5b6020026020010151611b13565b83600081518110611c4257fe5b6020908102919091010152611c5d818363ffffffff611ab216565b9050611c698685611ecf565b915060015b84811015611ccf57611c9e82848b8481518110611c8757fe5b60200260200101518c8560010181518110611c2857fe5b848281518110611caa57fe5b6020908102919091010152611cc5828463ffffffff611ab216565b9150600101611c6e565b50611cd983611ee2565b98975050505050505050565b600080611d0e83600260200201518914611d00576001611d03565b60005b60ff168760016111d6565b90506000611d3884600360200201518914611d2a576001611d2d565b60005b60ff168860026111d6565b9050611d77611d5367ffffffffffffffff88168960006111d6565b602086015160408701516060880151611d72928a929091889088612037565b611a86565b9998505050505050505050565b8151815160208401516040850151600093611dab939092611d72928b92918b918b90612037565b9695505050505050565b6001600b5460ff166002811115611dc857fe5b1415611ddb57611dd6611a12565b61064a565b61064a611995565b8251600090610100811115611df757600080fd5b8260005b82811015611ea75760028606611e5457868181518110611e1757fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150611e99565b81878281518110611e6157fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101611dfb565b5095945050505050565b6000818381611ebc57fe5b06828481611ec657fe5b04019392505050565b6000818381611eda57fe5b049392505050565b6000815b60018151111561201a5760606002825160010181611f0057fe5b0467ffffffffffffffff81118015611f1757600080fd5b50604051908082528060200260200182016040528015611f41578160200160208202803683370190505b50905060005b8151811015612012578251816002026001011015611fda57828160020281518110611f6e57fe5b6020026020010151838260020260010181518110611f8857fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611fc957fe5b60200260200101818152505061200a565b828160020281518110611fe957fe5b6020026020010151828281518110611ffd57fe5b6020026020010181815250505b600101611f47565b509050611ee6565b8060008151811061202757fe5b6020026020010151915050919050565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b60405180608001604052806004906020820280368337509192915050565b8280548282559060005260206000209081019282156120f3579160200282015b828111156120f35781546001600160a01b0319166001600160a01b038435161782556020909201916001909101906120c0565b506120ff92915061214b565b5090565b82600281019282156120f3579160200282015b828111156120f357825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190612116565b61066891905b808211156120ff5780546001600160a01b031916815560010161215156fea2646970667358221220a635d0c106bd09edb3a7a088c7855dc437bfeec5e238c21f1aa0ce301812363b64736f6c634300060b0033"

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
