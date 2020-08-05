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
const ExecutionChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"assertionHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"segmentIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"Continued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_machineHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_inboxAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_messageAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_logAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_outCounts\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_gases\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_totalSteps\",\"type\":\"uint64\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_segmentToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"chooseSegment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oneStepProof\",\"type\":\"address\"}],\"name\":\"connectOneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_challengeState\",\"type\":\"bytes32\"}],\"name\":\"initializeBisection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_firstInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_firstInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"_kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_msgData\",\"type\":\"bytes\"}],\"name\":\"oneStepProofWithMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeoutChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExecutionChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionChallengeFuncSigs = map[string]string{
	"efaa0772": "bisectAssertion(bytes32[],bytes32[],bytes32[],bytes32[],uint64[],uint64[],uint64)",
	"79a9ad85": "chooseSegment(uint256,bytes,bytes32,bytes32)",
	"2cb970f3": "connectOneStepProof(address)",
	"02ad1e4e": "initializeBisection(address,address,address,uint256,bytes32)",
	"082379bb": "oneStepProof(bytes32,bytes32,bytes32,bytes)",
	"5cd53989": "oneStepProofWithMessage(bytes32,bytes32,bytes32,bytes,uint8,uint256,uint256,address,uint256,bytes)",
	"ced5c1bf": "timeoutChallenge()",
}

// ExecutionChallengeBin is the compiled bytecode used for deploying new contracts.
var ExecutionChallengeBin = "0x608060405234801561001057600080fd5b50612116806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80635cd539891161005b5780635cd53989146101a357806379a9ad8514610307578063ced5c1bf146103b7578063efaa0772146103bf5761007d565b806302ad1e4e14610082578063082379bb146100c65780632cb970f31461017d575b600080fd5b6100c4600480360360a081101561009857600080fd5b506001600160a01b038135811691602081013582169160408201351690606081013590608001356106f5565b005b6100c4600480360360808110156100dc57600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561010957600080fd5b82018360208201111561011b57600080fd5b803590602001918460018302840111600160201b8311171561013c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061070a945050505050565b6100c46004803603602081101561019357600080fd5b50356001600160a01b03166109ed565b6100c460048036036101408110156101ba57600080fd5b81359160208101359160408201359190810190608081016060820135600160201b8111156101e757600080fd5b8201836020820111156101f957600080fd5b803590602001918460018302840111600160201b8311171561021a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929560ff85351695602086013595604081013595506001600160a01b0360608201351694506080810135935060c081019060a00135600160201b81111561029357600080fd5b8201836020820111156102a557600080fd5b803590602001918460018302840111600160201b831117156102c657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a0f945050505050565b6100c46004803603608081101561031d57600080fd5b81359190810190604081016020820135600160201b81111561033e57600080fd5b82018360208201111561035057600080fd5b803590602001918460018302840111600160201b8311171561037157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610d70565b6100c4611031565b6100c4600480360360e08110156103d557600080fd5b810190602081018135600160201b8111156103ef57600080fd5b82018360208201111561040157600080fd5b803590602001918460208302840111600160201b8311171561042257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561047157600080fd5b82018360208201111561048357600080fd5b803590602001918460208302840111600160201b831117156104a457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104f357600080fd5b82018360208201111561050557600080fd5b803590602001918460208302840111600160201b8311171561052657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561057557600080fd5b82018360208201111561058757600080fd5b803590602001918460208302840111600160201b831117156105a857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156105f757600080fd5b82018360208201111561060957600080fd5b803590602001918460208302840111600160201b8311171561062a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561067957600080fd5b82018360208201111561068b57600080fd5b803590602001918460208302840111600160201b831117156106ac57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550505090356001600160401b031691506111119050565b610701858585856112e9565b60065550505050565b60055460ff16600281111561071b57fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906107c95760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561078e578181015183820152602001610776565b50505050905090810190601f1680156107bb5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506003546107d643611404565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906108495760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146108c55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5060006108d0612019565b600754604051630e16f04560e21b81526004810188815260248201889052604482018790526080606483019081528651608484015286516001600160a01b039094169363385bc114938b938b938b938b9360a40190602085019080838360005b83811015610948578181015183820152602001610930565b50505050905090810190601f1680156109755780820380516001836020036101000a031916815260200191505b509550505050505060c06040518083038186803b15801561099557600080fd5b505afa1580156109a9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c08110156109ce57600080fd5b508051925060200190506109e5828787878561140b565b505050505050565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b60055460ff166002811115610a2057fe5b600114604051806040016040528060098152602001684249535f535441544560b81b81525090610a915760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b50600354610a9e43611404565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610b115760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610b8d5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b506000610b98612019565b600760009054906101000a90046001600160a01b03166001600160a01b03166396105dce8d8d8d8d8d8d8d8d8d8d6040518b63ffffffff1660e01b8152600401808b81526020018a8152602001898152602001806020018860ff1660ff168152602001878152602001868152602001856001600160a01b03166001600160a01b031681526020018481526020018060200183810383528a818151815260200191508051906020019080838360005b83811015610c5e578181015183820152602001610c46565b50505050905090810190601f168015610c8b5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610cbe578181015183820152602001610ca6565b50505050905090810190601f168015610ceb5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060c06040518083038186803b158015610d1257600080fd5b505afa158015610d26573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c0811015610d4b57600080fd5b50805192506020019050610d62828d8d8d8561140b565b505050505050505050505050565b60055460ff166002811115610d8157fe5b60021460405180604001604052806009815260200168434f4e5f535441544560b81b81525090610df25760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b50600354610dff43611404565b11156040518060400160405280600c81526020016b434f4e5f444541444c494e4560a01b81525090610e725760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5060025460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b03163314610eee5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5060065482146040518060400160405280600881526020016721a7a72fa82922ab60c11b81525090610f615760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b50610f718383838760010161150d565b6040518060400160405280600981526020016821a7a72fa82927a7a360b91b81525090610fdf5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b506006819055610fed61160e565b60035460408051868152602081019290925280517f1e1c1e4e68a25c69a078a396e73975691c071d69ef789015dc16a562957804e49281900390910190a150505050565b60035461103d43611404565b1161108f576040805162461bcd60e51b815260206004820152601760248201527f446561646c696e65206861736e27742065787069726564000000000000000000604482015290519081900360640190fd5b600160055460ff1660028111156110a257fe5b14156110de576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a16110d961162a565b61110f565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a161110f61169e565b565b60055460ff16600281111561112257fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906111935760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b506003546111a043611404565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906112135760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5060015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b0316331461128f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b50611298612037565b6040518060e00160405280898152602001888152602001878152602001868152602001858152602001848152602001836001600160401b031681525090506112df816116f6565b5050505050505050565b600060055460ff1660028111156112fc57fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b815250906113715760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b50600080546001600160a01b038681166001600160a01b03199283161790925560018054868416908316178155600280549386169390921692909217905560048290556005805460ff191690911790556113c9611852565b60035460408051918252517fe070f23072cbc6c0fc7253b8d0120649d5f9d6e19a8aeab79eb50aa6360bcc679181900360200190a150505050565b6103e80290565b8051602082015160408301516060840151608085015161142961207d565b60405180610180016040528060016001600160401b031681526020018c6001600160401b031681526020018781526020018681526020018b81526020018581526020018a8152602001848152602001848b14611486576001611489565b60005b60ff166001600160401b03168152602001898152602001838152602001838a146114b45760016114b7565b60005b60ff16905290506114cf6114ca82611864565b611974565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a161150061169e565b5050505050505050505050565b600080838160205b88518111611600578089015193506020818a51036020018161153357fe5b0491505b60008211801561154a5750600286066001145b801561155857508160020a86115b1561156b57600286046001019550611537565b600286066115b65783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816115ae57fe5b0495506115f8565b82846040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816115f157fe5b0460010195505b602001611515565b505090941495945050505050565b600580546001919060ff191682805b021790555061110f611852565b600080546002546001546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b15801561168257600080fd5b505af1158015611696573d6000803e3d6000fd5b503392505050ff5b600080546001546002546040805163396f51cf60e01b81526001600160a01b039384166004820152918316602483015251919092169263396f51cf926044808201939182900301818387803b15801561168257600080fd5b80515160001901611706826119ea565b606081604051908082528060200260200182016040528015611732578160200160208202803883390190505b509050611758836117508560c001516001600160401b031685611cdf565b846000611cfd565b8160008151811061176557fe5b602090810291909101015260015b828110156117bc5761179d846117968660c001516001600160401b031686611d61565b8584611cfd565b8282815181106117a957fe5b6020908102919091010152600101611773565b506117c681611d74565b6117ce611d83565b7f81050542a90cf16d270921d19aeab083e0a9b460a208b224daf345c77cb4c3ce816003546040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b83811015611839578181015183820152602001611821565b50505050905001935050505060405180910390a1505050565b60045461185e43611404565b01600355565b6000816000015182602001518360400151846060015185608001518660a001518760c001518860e001518961010001518a61012001518b61014001518c6101600151604051602001808d6001600160401b03166001600160401b031660c01b81526008018c6001600160401b03166001600160401b031660c01b81526008018b81526020018a8152602001898152602001888152602001878152602001868152602001856001600160401b03166001600160401b031660c01b8152600801848152602001838152602001826001600160401b03166001600160401b031660c01b81526008019c50505050505050505050505050604051602081830303815290604052805190602001209050919050565b6006548114604051806040016040528060088152602001672124a9afa82922ab60c11b815250906119e65760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5050565b8051516020808301515160408051808201909152600a8152692124a9afa4a7282622a760b11b9281019290925260001983019214611a695760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b50604080830151518151808301909252600a8252692124a9afa4a7282622a760b11b60208301526001830114611ae05760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5081606001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b81525090611b5b5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b508160a001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b81525090611bd35760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5081608001515181600202146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b81525090611c4e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561078e578181015183820152602001610776565b5060008080805b84811015611cc0578560a001518181518110611c6d57fe5b60200260200101518401935085608001518181518110611c8957fe5b602002602001015183019250856080015181860181518110611ca757fe5b6020026020010151820191508080600101915050611c55565b50611cd86114ca868760c00151600088888888611d97565b5050505050565b6000818381611cea57fe5b06828481611cf457fe5b04019392505050565b6000611d58858584856001018960a001518781518110611d1957fe5b60200260200101518a608001518881518110611d3157fe5b60200260200101518b60800151898b0181518110611d4b57fe5b6020026020010151611d97565b95945050505050565b6000818381611d6c57fe5b049392505050565b611d7d81611edb565b60065550565b600580546002919060ff191660018361161d565b6000611ecf604051806101800160405280896001600160401b03168152602001866001600160401b031681526020018a600001518981518110611dd657fe5b602002602001015181526020018a600001518881518110611df357fe5b602002602001015181526020018a602001518981518110611e1057fe5b602002602001015181526020018a602001518881518110611e2d57fe5b602002602001015181526020018a604001518981518110611e4a57fe5b602002602001015181526020018a604001518881518110611e6757fe5b60200260200101518152602001856001600160401b031681526020018a606001518981518110611e9357fe5b602002602001015181526020018a606001518881518110611eb057fe5b60200260200101518152602001846001600160401b0316815250611864565b98975050505050505050565b6000815b600181511115611ffc5760606002825160010181611ef957fe5b04604051908082528060200260200182016040528015611f23578160200160208202803883390190505b50905060005b8151811015611ff4578251816002026001011015611fbc57828160020281518110611f5057fe5b6020026020010151838260020260010181518110611f6a57fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611fab57fe5b602002602001018181525050611fec565b828160020281518110611fcb57fe5b6020026020010151828281518110611fdf57fe5b6020026020010181815250505b600101611f29565b509050611edf565b8060008151811061200957fe5b6020026020010151915050919050565b6040518060a001604052806005906020820280388339509192915050565b6040518060e0016040528060608152602001606081526020016060815260200160608152602001606081526020016060815260200160006001600160401b031681525090565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290526101608101919091529056fea265627a7a7231582086225f7dc90cd34a6ca8ab201fd108c24eef57c4d6605e15b5e4e08f1335c3c264736f6c63430005110032"

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
func (_ExecutionChallenge *ExecutionChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ExecutionChallenge *ExecutionChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
