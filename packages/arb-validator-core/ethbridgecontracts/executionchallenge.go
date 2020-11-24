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

// ExecutionChallengeABI is the input ABI used to generate the binding from.
const ExecutionChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"assertionHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"segmentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Continued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_machineHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_inboxAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_messageAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_logAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_outCounts\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_gases\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_totalSteps\",\"type\":\"uint64\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_segmentToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"chooseSegment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oneStepProof\",\"type\":\"address\"}],\"name\":\"connectOneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_firstInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_firstInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"_kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_msgData\",\"type\":\"bytes\"}],\"name\":\"oneStepProofWithMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExecutionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionChallengeFuncSigs = map[string]string{
	"efaa0772": "bisectAssertion(bytes32[],bytes32[],bytes32[],bytes32[],uint64[],uint64[],uint64)",
	"79a9ad85": "chooseSegment(uint256,bytes,bytes32,bytes32)",
	"2cb970f3": "connectOneStepProof(address)",
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
	"6f791d29": "isMaster()",
	"082379bb": "oneStepProof(bytes32,bytes32,bytes32,bytes)",
	"5cd53989": "oneStepProofWithMessage(bytes32,bytes32,bytes32,bytes,uint8,uint256,uint256,address,uint256,bytes)",
	"ced5c1bf": "timeoutChallenge()",
}

// ExecutionChallengeBin is the compiled bytecode used for deploying new contracts.
var ExecutionChallengeBin = "0x60806040526000805460ff191660011790556121e5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80636f791d291161005b5780636f791d291461031257806379a9ad851461032e578063ced5c1bf146103de578063efaa0772146103e657610088565b806302ad1e4e1461008d578063082379bb146100d15780632cb970f3146101885780635cd53989146101ae575b600080fd5b6100cf600480360360a08110156100a357600080fd5b506001600160a01b0381358116916020810135821691604082013516906060810135906080013561071c565b005b6100cf600480360360808110156100e757600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561011457600080fd5b82018360208201111561012657600080fd5b803590602001918460018302840111600160201b8311171561014757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610731945050505050565b6100cf6004803603602081101561019e57600080fd5b50356001600160a01b0316610a14565b6100cf60048036036101408110156101c557600080fd5b81359160208101359160408201359190810190608081016060820135600160201b8111156101f257600080fd5b82018360208201111561020457600080fd5b803590602001918460018302840111600160201b8311171561022557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929560ff85351695602086013595604081013595506001600160a01b0360608201351694506080810135935060c081019060a00135600160201b81111561029e57600080fd5b8201836020820111156102b057600080fd5b803590602001918460018302840111600160201b831117156102d157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a36945050505050565b61031a610d97565b604080519115158252519081900360200190f35b6100cf6004803603608081101561034457600080fd5b81359190810190604081016020820135600160201b81111561036557600080fd5b82018360208201111561037757600080fd5b803590602001918460018302840111600160201b8311171561039857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610da0565b6100cf611061565b6100cf600480360360e08110156103fc57600080fd5b810190602081018135600160201b81111561041657600080fd5b82018360208201111561042857600080fd5b803590602001918460208302840111600160201b8311171561044957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561049857600080fd5b8201836020820111156104aa57600080fd5b803590602001918460208302840111600160201b831117156104cb57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561051a57600080fd5b82018360208201111561052c57600080fd5b803590602001918460208302840111600160201b8311171561054d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561059c57600080fd5b8201836020820111156105ae57600080fd5b803590602001918460208302840111600160201b831117156105cf57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561061e57600080fd5b82018360208201111561063057600080fd5b803590602001918460208302840111600160201b8311171561065157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106a057600080fd5b8201836020820111156106b257600080fd5b803590602001918460208302840111600160201b831117156106d357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550505090356001600160401b031691506111419050565b61072885858585611319565b60065550505050565b60055460ff16600281111561074257fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906107f05760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156107b557818101518382015260200161079d565b50505050905090810190601f1680156107e25780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506003546107fd43611446565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906108705760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146108ec5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5060006108f76120e8565b600754604051630e16f04560e21b81526004810188815260248201889052604482018790526080606483019081528651608484015286516001600160a01b039094169363385bc114938b938b938b938b9360a40190602085019080838360005b8381101561096f578181015183820152602001610957565b50505050905090810190601f16801561099c5780820380516001836020036101000a031916815260200191505b509550505050505060c06040518083038186803b1580156109bc57600080fd5b505afa1580156109d0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c08110156109f557600080fd5b50805192506020019050610a0c828787878561144d565b505050505050565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b60055460ff166002811115610a4757fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090610ab85760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b50600354610ac543611446565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610b385760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610bb45760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b506000610bbf6120e8565b600760009054906101000a90046001600160a01b03166001600160a01b03166396105dce8d8d8d8d8d8d8d8d8d8d6040518b63ffffffff1660e01b8152600401808b81526020018a8152602001898152602001806020018860ff1660ff168152602001878152602001868152602001856001600160a01b03166001600160a01b031681526020018481526020018060200183810383528a818151815260200191508051906020019080838360005b83811015610c85578181015183820152602001610c6d565b50505050905090810190601f168015610cb25780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610ce5578181015183820152602001610ccd565b50505050905090810190601f168015610d125780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060c06040518083038186803b158015610d3957600080fd5b505afa158015610d4d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c0811015610d7257600080fd5b50805192506020019050610d89828d8d8d8561144d565b505050505050505050505050565b60005460ff1690565b60055460ff166002811115610db157fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b81525090610e225760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b50600354610e2f43611446565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b81525090610ea25760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b03163314610f1e5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b81525090610f915760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b50610fa18383838760010161154f565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b8152509061100f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b50600681905561101d611650565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60035461106d43611446565b116110bf576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156110d257fe5b141561110e576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a161110961166c565b61113f565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a161113f6116e9565b565b60055460ff16600281111561115257fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906111c35760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b506003546111d043611446565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906112435760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146112bf5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b506112c8612106565b6040518060e00160405280898152602001888152602001878152602001868152602001858152602001848152602001836001600160401b0316815250905061130f81611745565b5050505050505050565b600060055460ff16600281111561132c57fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b815250906113a15760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5060008054610100600160a81b0319166101006001600160a01b038781169190910291909117909155600180546001600160a01b0319908116868416178255600280549091169285169290921790915560048290556005805460ff1916909117905561140b6118a1565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e80290565b8051602082015160408301516060840151608085015161146b61214c565b60405180610180016040528060016001600160401b031681526020018c6001600160401b031681526020018781526020018681526020018b81526020018581526020018a8152602001848152602001848b146114c85760016114cb565b60005b60ff166001600160401b03168152602001898152602001838152602001838a146114f65760016114f9565b60005b60ff169052905061151161150c826118b3565b6119c3565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a16115426116e9565b5050505050505050505050565b600080838160205b88518111611642578089015193506020818a51036020018161157557fe5b0491505b60008211801561158c5750600286066001145b801561159a57508160020a86115b156115ad57600286046001019550611579565b600286066115f85783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816115f057fe5b04955061163a565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161163357fe5b0460010195505b602001611557565b505090941495945050505050565b600580546001919060ff191682805b021790555061113f6118a1565b600080546002546001546040805163396f51cf60e01b81526001600160a01b0393841660048201529183166024830152516101009093049091169263396f51cf9260448084019382900301818387803b1580156116c857600080fd5b505af11580156116dc573d6000803e3d6000fd5b5050505061113f33611a39565b600080546001546002546040805163396f51cf60e01b81526001600160a01b0393841660048201529183166024830152516101009093049091169263396f51cf9260448084019382900301818387803b1580156116c857600080fd5b8051516000190161175582611ab9565b606081604051908082528060200260200182016040528015611781578160200160208202803883390190505b5090506117a78361179f8560c001516001600160401b031685611dae565b846000611dcc565b816000815181106117b457fe5b602090810291909101015260015b8281101561180b576117ec846117e58660c001516001600160401b031686611e30565b8584611dcc565b8282815181106117f857fe5b60209081029190910101526001016117c2565b5061181581611e43565b61181d611e52565b7f81050542a90cf16d270921d19aeab083e0a9b460a208b224daf345c77cb4c3ce816003546040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b83811015611888578181015183820152602001611870565b50505050905001935050505060405180910390a1505050565b6004546118ad43611446565b01600355565b6000816000015182602001518360400151846060015185608001518660a001518760c001518860e001518961010001518a61012001518b61014001518c6101600151604051602001808d6001600160401b03166001600160401b031660c01b81526008018c6001600160401b03166001600160401b031660c01b81526008018b81526020018a8152602001898152602001888152602001878152602001868152602001856001600160401b03166001600160401b031660c01b8152600801848152602001838152602001826001600160401b03166001600160401b031660c01b81526008019c50505050505050505050505050604051602081830303815290604052805190602001209050919050565b6006548114604051806040016040528060088152602001672124a9afa82922ab60c11b81525090611a355760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5050565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff1615611aac5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b50806001600160a01b0316ff5b8051516020808301515160408051808201909152600a8152692124a9afa4a7282622a760b11b9281019290925260001983019214611b385760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b50604080830151518151808301909252600a8252692124a9afa4a7282622a760b11b60208301526001830114611baf5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5081606001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b81525090611c2a5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b508160a001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b81525090611ca25760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5081608001515181600202146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b81525090611d1d5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156107b557818101518382015260200161079d565b5060008080805b84811015611d8f578560a001518181518110611d3c57fe5b60200260200101518401935085608001518181518110611d5857fe5b602002602001015183019250856080015181860181518110611d7657fe5b6020026020010151820191508080600101915050611d24565b50611da761150c868760c00151600088888888611e66565b5050505050565b6000818381611db957fe5b06828481611dc357fe5b04019392505050565b6000611e27858584856001018960a001518781518110611de857fe5b60200260200101518a608001518881518110611e0057fe5b60200260200101518b60800151898b0181518110611e1a57fe5b6020026020010151611e66565b95945050505050565b6000818381611e3b57fe5b049392505050565b611e4c81611faa565b60065550565b600580546002919060ff191660018361165f565b6000611f9e604051806101800160405280896001600160401b03168152602001866001600160401b031681526020018a600001518981518110611ea557fe5b602002602001015181526020018a600001518881518110611ec257fe5b602002602001015181526020018a602001518981518110611edf57fe5b602002602001015181526020018a602001518881518110611efc57fe5b602002602001015181526020018a604001518981518110611f1957fe5b602002602001015181526020018a604001518881518110611f3657fe5b60200260200101518152602001856001600160401b031681526020018a606001518981518110611f6257fe5b602002602001015181526020018a606001518881518110611f7f57fe5b60200260200101518152602001846001600160401b03168152506118b3565b98975050505050505050565b6000815b6001815111156120cb5760606002825160010181611fc857fe5b04604051908082528060200260200182016040528015611ff2578160200160208202803883390190505b50905060005b81518110156120c357825181600202600101101561208b5782816002028151811061201f57fe5b602002602001015183826002026001018151811061203957fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012082828151811061207a57fe5b6020026020010181815250506120bb565b82816002028151811061209a57fe5b60200260200101518282815181106120ae57fe5b6020026020010181815250505b600101611ff8565b509050611fae565b806000815181106120d857fe5b6020026020010151915050919050565b6040518060a001604052806005906020820280388339509192915050565b6040518060e0016040528060608152602001606081526020016060815260200160608152602001606081526020016060815260200160006001600160401b031681525090565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290526101608101919091529056fea265627a7a723158205f70dc58247e18bc44d35f8fe989930bfd42ab11f5f831790c850a7055ec5caf64736f6c63430005110032"

// DeployExecutionChallenge deploys a new Ethereum contract, binding an instance of ExecutionChallenge to it.
func DeployExecutionChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExecutionChallenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExecutionChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutionChallenge{ExecutionChallengeCaller: ExecutionChallengeCaller{contract: contract}, ExecutionChallengeTransactor: ExecutionChallengeTransactor{contract: contract}, ExecutionChallengeFilterer: ExecutionChallengeFilterer{contract: contract}}, nil
}

// ExecutionChallenge is an auto generated Go binding around an Ethereum contract.
type ExecutionChallenge struct {
	ExecutionChallengeCaller     // Read-only binding to the contract
	ExecutionChallengeTransactor // Write-only binding to the contract
	ExecutionChallengeFilterer   // Log filterer for contract events
}

// ExecutionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionChallengeSession struct {
	Contract     *ExecutionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionChallengeCallerSession struct {
	Contract *ExecutionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ExecutionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionChallengeTransactorSession struct {
	Contract     *ExecutionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ExecutionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionChallengeRaw struct {
	Contract *ExecutionChallenge // Generic contract binding to access the raw methods on
}

// ExecutionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionChallengeCallerRaw struct {
	Contract *ExecutionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionChallengeTransactorRaw struct {
	Contract *ExecutionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionChallenge creates a new instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallenge(address common.Address, backend bind.ContractBackend) (*ExecutionChallenge, error) {
	contract, err := bindExecutionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallenge{ExecutionChallengeCaller: ExecutionChallengeCaller{contract: contract}, ExecutionChallengeTransactor: ExecutionChallengeTransactor{contract: contract}, ExecutionChallengeFilterer: ExecutionChallengeFilterer{contract: contract}}, nil
}

// NewExecutionChallengeCaller creates a new read-only instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeCaller(address common.Address, caller bind.ContractCaller) (*ExecutionChallengeCaller, error) {
	contract, err := bindExecutionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeCaller{contract: contract}, nil
}

// NewExecutionChallengeTransactor creates a new write-only instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionChallengeTransactor, error) {
	contract, err := bindExecutionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeTransactor{contract: contract}, nil
}

// NewExecutionChallengeFilterer creates a new log filterer instance of ExecutionChallenge, bound to a specific deployed contract.
func NewExecutionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionChallengeFilterer, error) {
	contract, err := bindExecutionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeFilterer{contract: contract}, nil
}

// bindExecutionChallenge binds a generic wrapper to an already deployed contract.
func bindExecutionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionChallenge *ExecutionChallengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionChallenge.Contract.ExecutionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionChallenge *ExecutionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ExecutionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionChallenge *ExecutionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ExecutionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionChallenge *ExecutionChallengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionChallenge *ExecutionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionChallenge *ExecutionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.contract.Transact(opts, method, params...)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ExecutionChallenge *ExecutionChallengeCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExecutionChallenge.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ExecutionChallenge *ExecutionChallengeSession) IsMaster() (bool, error) {
	return _ExecutionChallenge.Contract.IsMaster(&_ExecutionChallenge.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ExecutionChallenge *ExecutionChallengeCallerSession) IsMaster() (bool, error) {
	return _ExecutionChallenge.Contract.IsMaster(&_ExecutionChallenge.CallOpts)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xefaa0772.
//
// Solidity: function bisectAssertion(bytes32[] _machineHashes, bytes32[] _inboxAccs, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _outCounts, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) BisectAssertion(opts *bind.TransactOpts, _machineHashes [][32]byte, _inboxAccs [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _outCounts []uint64, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "bisectAssertion", _machineHashes, _inboxAccs, _messageAccs, _logAccs, _outCounts, _gases, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xefaa0772.
//
// Solidity: function bisectAssertion(bytes32[] _machineHashes, bytes32[] _inboxAccs, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _outCounts, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) BisectAssertion(_machineHashes [][32]byte, _inboxAccs [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _outCounts []uint64, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.BisectAssertion(&_ExecutionChallenge.TransactOpts, _machineHashes, _inboxAccs, _messageAccs, _logAccs, _outCounts, _gases, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xefaa0772.
//
// Solidity: function bisectAssertion(bytes32[] _machineHashes, bytes32[] _inboxAccs, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _outCounts, uint64[] _gases, uint64 _totalSteps) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) BisectAssertion(_machineHashes [][32]byte, _inboxAccs [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _outCounts []uint64, _gases []uint64, _totalSteps uint64) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.BisectAssertion(&_ExecutionChallenge.TransactOpts, _machineHashes, _inboxAccs, _messageAccs, _logAccs, _outCounts, _gases, _totalSteps)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) ChooseSegment(opts *bind.TransactOpts, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ChooseSegment(&_ExecutionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ChooseSegment is a paid mutator transaction binding the contract method 0x79a9ad85.
//
// Solidity: function chooseSegment(uint256 _segmentToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) ChooseSegment(_segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ChooseSegment(&_ExecutionChallenge.TransactOpts, _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ConnectOneStepProof is a paid mutator transaction binding the contract method 0x2cb970f3.
//
// Solidity: function connectOneStepProof(address oneStepProof) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) ConnectOneStepProof(opts *bind.TransactOpts, oneStepProof common.Address) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "connectOneStepProof", oneStepProof)
}

// ConnectOneStepProof is a paid mutator transaction binding the contract method 0x2cb970f3.
//
// Solidity: function connectOneStepProof(address oneStepProof) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) ConnectOneStepProof(oneStepProof common.Address) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ConnectOneStepProof(&_ExecutionChallenge.TransactOpts, oneStepProof)
}

// ConnectOneStepProof is a paid mutator transaction binding the contract method 0x2cb970f3.
//
// Solidity: function connectOneStepProof(address oneStepProof) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) ConnectOneStepProof(oneStepProof common.Address) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.ConnectOneStepProof(&_ExecutionChallenge.TransactOpts, oneStepProof)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) InitializeBisection(opts *bind.TransactOpts, _rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "initializeBisection", _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.InitializeBisection(&_ExecutionChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// InitializeBisection is a paid mutator transaction binding the contract method 0x02ad1e4e.
//
// Solidity: function initializeBisection(address _rollupAddress, address _asserter, address _challenger, uint256 _challengePeriodTicks, bytes32 _challengeState) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) InitializeBisection(_rollupAddress common.Address, _asserter common.Address, _challenger common.Address, _challengePeriodTicks *big.Int, _challengeState [32]byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.InitializeBisection(&_ExecutionChallenge.TransactOpts, _rollupAddress, _asserter, _challenger, _challengePeriodTicks, _challengeState)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x082379bb.
//
// Solidity: function oneStepProof(bytes32 _firstInbox, bytes32 _firstMessage, bytes32 _firstLog, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProof(opts *bind.TransactOpts, _firstInbox [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "oneStepProof", _firstInbox, _firstMessage, _firstLog, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x082379bb.
//
// Solidity: function oneStepProof(bytes32 _firstInbox, bytes32 _firstMessage, bytes32 _firstLog, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) OneStepProof(_firstInbox [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.OneStepProof(&_ExecutionChallenge.TransactOpts, _firstInbox, _firstMessage, _firstLog, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x082379bb.
//
// Solidity: function oneStepProof(bytes32 _firstInbox, bytes32 _firstMessage, bytes32 _firstLog, bytes _proof) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) OneStepProof(_firstInbox [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.OneStepProof(&_ExecutionChallenge.TransactOpts, _firstInbox, _firstMessage, _firstLog, _proof)
}

// OneStepProofWithMessage is a paid mutator transaction binding the contract method 0x5cd53989.
//
// Solidity: function oneStepProofWithMessage(bytes32 _firstInbox, bytes32 _firstMessage, bytes32 _firstLog, bytes _proof, uint8 _kind, uint256 _blockNumber, uint256 _timestamp, address _sender, uint256 _inboxSeqNum, bytes _msgData) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProofWithMessage(opts *bind.TransactOpts, _firstInbox [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte, _kind uint8, _blockNumber *big.Int, _timestamp *big.Int, _sender common.Address, _inboxSeqNum *big.Int, _msgData []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "oneStepProofWithMessage", _firstInbox, _firstMessage, _firstLog, _proof, _kind, _blockNumber, _timestamp, _sender, _inboxSeqNum, _msgData)
}

// OneStepProofWithMessage is a paid mutator transaction binding the contract method 0x5cd53989.
//
// Solidity: function oneStepProofWithMessage(bytes32 _firstInbox, bytes32 _firstMessage, bytes32 _firstLog, bytes _proof, uint8 _kind, uint256 _blockNumber, uint256 _timestamp, address _sender, uint256 _inboxSeqNum, bytes _msgData) returns()
func (_ExecutionChallenge *ExecutionChallengeSession) OneStepProofWithMessage(_firstInbox [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte, _kind uint8, _blockNumber *big.Int, _timestamp *big.Int, _sender common.Address, _inboxSeqNum *big.Int, _msgData []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.OneStepProofWithMessage(&_ExecutionChallenge.TransactOpts, _firstInbox, _firstMessage, _firstLog, _proof, _kind, _blockNumber, _timestamp, _sender, _inboxSeqNum, _msgData)
}

// OneStepProofWithMessage is a paid mutator transaction binding the contract method 0x5cd53989.
//
// Solidity: function oneStepProofWithMessage(bytes32 _firstInbox, bytes32 _firstMessage, bytes32 _firstLog, bytes _proof, uint8 _kind, uint256 _blockNumber, uint256 _timestamp, address _sender, uint256 _inboxSeqNum, bytes _msgData) returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) OneStepProofWithMessage(_firstInbox [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte, _kind uint8, _blockNumber *big.Int, _timestamp *big.Int, _sender common.Address, _inboxSeqNum *big.Int, _msgData []byte) (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.OneStepProofWithMessage(&_ExecutionChallenge.TransactOpts, _firstInbox, _firstMessage, _firstLog, _proof, _kind, _blockNumber, _timestamp, _sender, _inboxSeqNum, _msgData)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeTransactor) TimeoutChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionChallenge.contract.Transact(opts, "timeoutChallenge")
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeSession) TimeoutChallenge() (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.TimeoutChallenge(&_ExecutionChallenge.TransactOpts)
}

// TimeoutChallenge is a paid mutator transaction binding the contract method 0xced5c1bf.
//
// Solidity: function timeoutChallenge() returns()
func (_ExecutionChallenge *ExecutionChallengeTransactorSession) TimeoutChallenge() (*types.Transaction, error) {
	return _ExecutionChallenge.Contract.TimeoutChallenge(&_ExecutionChallenge.TransactOpts)
}

// ExecutionChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the ExecutionChallenge contract.
type ExecutionChallengeAsserterTimedOutIterator struct {
	Event *ExecutionChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeAsserterTimedOut)
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
		it.Event = new(ExecutionChallengeAsserterTimedOut)
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
func (it *ExecutionChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the ExecutionChallenge contract.
type ExecutionChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*ExecutionChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeAsserterTimedOutIterator{contract: _ExecutionChallenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeAsserterTimedOut)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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

// ParseAsserterTimedOut is a log parse operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*ExecutionChallengeAsserterTimedOut, error) {
	event := new(ExecutionChallengeAsserterTimedOut)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ExecutionChallenge contract.
type ExecutionChallengeBisectedAssertionIterator struct {
	Event *ExecutionChallengeBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeBisectedAssertion)
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
		it.Event = new(ExecutionChallengeBisectedAssertion)
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
func (it *ExecutionChallengeBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeBisectedAssertion represents a BisectedAssertion event raised by the ExecutionChallenge contract.
type ExecutionChallengeBisectedAssertion struct {
	AssertionHashes [][32]byte
	DeadlineTicks   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0x81050542a90cf16d270921d19aeab083e0a9b460a208b224daf345c77cb4c3ce.
//
// Solidity: event BisectedAssertion(bytes32[] assertionHashes, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterBisectedAssertion(opts *bind.FilterOpts) (*ExecutionChallengeBisectedAssertionIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeBisectedAssertionIterator{contract: _ExecutionChallenge.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0x81050542a90cf16d270921d19aeab083e0a9b460a208b224daf345c77cb4c3ce.
//
// Solidity: event BisectedAssertion(bytes32[] assertionHashes, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeBisectedAssertion) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeBisectedAssertion)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0x81050542a90cf16d270921d19aeab083e0a9b460a208b224daf345c77cb4c3ce.
//
// Solidity: event BisectedAssertion(bytes32[] assertionHashes, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseBisectedAssertion(log types.Log) (*ExecutionChallengeBisectedAssertion, error) {
	event := new(ExecutionChallengeBisectedAssertion)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the ExecutionChallenge contract.
type ExecutionChallengeChallengerTimedOutIterator struct {
	Event *ExecutionChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeChallengerTimedOut)
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
		it.Event = new(ExecutionChallengeChallengerTimedOut)
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
func (it *ExecutionChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the ExecutionChallenge contract.
type ExecutionChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*ExecutionChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeChallengerTimedOutIterator{contract: _ExecutionChallenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeChallengerTimedOut)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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

// ParseChallengerTimedOut is a log parse operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*ExecutionChallengeChallengerTimedOut, error) {
	event := new(ExecutionChallengeChallengerTimedOut)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeContinuedIterator is returned from FilterContinued and is used to iterate over the raw logs and unpacked data for Continued events raised by the ExecutionChallenge contract.
type ExecutionChallengeContinuedIterator struct {
	Event *ExecutionChallengeContinued // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeContinued)
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
		it.Event = new(ExecutionChallengeContinued)
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
func (it *ExecutionChallengeContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeContinued represents a Continued event raised by the ExecutionChallenge contract.
type ExecutionChallengeContinued struct {
	SegmentIndex  *big.Int
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContinued is a free log retrieval operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterContinued(opts *bind.FilterOpts) (*ExecutionChallengeContinuedIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeContinuedIterator{contract: _ExecutionChallenge.contract, event: "Continued", logs: logs, sub: sub}, nil
}

// WatchContinued is a free log subscription operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchContinued(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeContinued) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "Continued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeContinued)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
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

// ParseContinued is a log parse operation binding the contract event 0x1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e4.
//
// Solidity: event Continued(uint256 segmentIndex, uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseContinued(log types.Log) (*ExecutionChallengeContinued, error) {
	event := new(ExecutionChallengeContinued)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "Continued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ExecutionChallenge contract.
type ExecutionChallengeInitiatedChallengeIterator struct {
	Event *ExecutionChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeInitiatedChallenge)
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
		it.Event = new(ExecutionChallengeInitiatedChallenge)
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
func (it *ExecutionChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the ExecutionChallenge contract.
type ExecutionChallengeInitiatedChallenge struct {
	DeadlineTicks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ExecutionChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeInitiatedChallengeIterator{contract: _ExecutionChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeInitiatedChallenge)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0xe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc67.
//
// Solidity: event InitiatedChallenge(uint256 deadlineTicks)
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ExecutionChallengeInitiatedChallenge, error) {
	event := new(ExecutionChallengeInitiatedChallenge)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExecutionChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ExecutionChallenge contract.
type ExecutionChallengeOneStepProofCompletedIterator struct {
	Event *ExecutionChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ExecutionChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionChallengeOneStepProofCompleted)
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
		it.Event = new(ExecutionChallengeOneStepProofCompleted)
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
func (it *ExecutionChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the ExecutionChallenge contract.
type ExecutionChallengeOneStepProofCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ExecutionChallenge *ExecutionChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*ExecutionChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _ExecutionChallenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &ExecutionChallengeOneStepProofCompletedIterator{contract: _ExecutionChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ExecutionChallenge *ExecutionChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ExecutionChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _ExecutionChallenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionChallengeOneStepProofCompleted)
				if err := _ExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ExecutionChallenge *ExecutionChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*ExecutionChallengeOneStepProofCompleted, error) {
	event := new(ExecutionChallengeOneStepProofCompleted)
	if err := _ExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}
