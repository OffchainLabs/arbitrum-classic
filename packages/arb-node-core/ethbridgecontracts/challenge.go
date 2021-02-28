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

// ChallengeABI is the input ABI used to generate the binding from.
const ChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContinuedExecutionProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"asserter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asserterTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_merkleRoute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_oldEndHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_gasUsedBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionRest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengerTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentResponderTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"executors\",\"outputs\":[{\"internalType\":\"contractIOneStepProof\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOneStepProof[]\",\"name\":\"_executors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_resultReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_executionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_asserterTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengerTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initializeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMoveBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_merkleRoute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_oldEndHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_initialSendAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_initialLogAcc\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[3]\",\"name\":\"_initialState\",\"type\":\"uint256[3]\"},{\"internalType\":\"bytes\",\"name\":\"_executionProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_bufferProof\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"prover\",\"type\":\"uint8\"}],\"name\":\"oneStepProveExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_merkleRoute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_oldEndHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_gasUsedBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionRest\",\"type\":\"bytes32\"}],\"name\":\"proveContinuedExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"turn\",\"outputs\":[{\"internalType\":\"enumChallenge.Turn\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFuncSigs = map[string]string{
	"bb4af0b1": "asserter()",
	"9a9e4f44": "asserterTimeLeft()",
	"8e7b84c5": "bisectExecution(bytes32[],uint256,uint256,uint256,bytes32,uint256,bytes32,bytes32[])",
	"e78cea92": "bridge()",
	"843d5a5c": "challengeState()",
	"534db0e2": "challenger()",
	"41e8510c": "challengerTimeLeft()",
	"8a8cd218": "currentResponder()",
	"e87e3589": "currentResponderTimeLeft()",
	"f97a05df": "executors(uint256)",
	"32f8c24f": "initializeChallenge(address[],address,bytes32,uint256,address,address,uint256,uint256,address)",
	"6f791d29": "isMaster()",
	"925f9a96": "lastMoveBlock()",
	"e08f819e": "oneStepProveExecution(bytes32[],uint256,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256[3],bytes,bytes,uint8)",
	"deda4115": "proveContinuedExecution(bytes32[],uint256,uint256,uint256,bytes32,uint256,bytes32)",
	"70dea79a": "timeout()",
	"8b299903": "turn()",
}

// ChallengeBin is the compiled bytecode used for deploying new contracts.
var ChallengeBin = "0x608060405234801561001057600080fd5b506000805460ff19166001179055611f378061002d6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638e7b84c5116100a2578063deda411511610071578063deda411514610350578063e08f819e146103dc578063e78cea92146105b3578063e87e3589146105bb578063f97a05df146105c35761010b565b80638e7b84c514610255578063925f9a96146103385780639a9e4f4414610340578063bb4af0b1146103485761010b565b806370dea79a116100de57806370dea79a14610214578063843d5a5c1461021c5780638a8cd218146102245780638b2999031461022c5761010b565b806332f8c24f1461011057806341e8510c146101ba578063534db0e2146101d45780636f791d29146101f8575b600080fd5b6101b8600480360361012081101561012757600080fd5b810190602081018135600160201b81111561014157600080fd5b82018360208201111561015357600080fd5b803590602001918460208302840111600160201b8311171561017457600080fd5b91935091506001600160a01b0381358116916020810135916040820135916060810135821691608082013581169160a08101359160c08201359160e00135166105e0565b005b6101c2610758565b60408051918252519081900360200190f35b6101dc61075e565b604080516001600160a01b039092168252519081900360200190f35b61020061076d565b604080519115158252519081900360200190f35b6101b8610777565b6101c2610893565b6101dc610899565b610234610928565b6040518082600281111561024457fe5b815260200191505060405180910390f35b6101b8600480360361010081101561026c57600080fd5b810190602081018135600160201b81111561028657600080fd5b82018360208201111561029857600080fd5b803590602001918460208302840111600160201b831117156102b957600080fd5b9193909282359260208101359260408201359260608301359260808101359260a082013592909160e081019060c00135600160201b8111156102fa57600080fd5b82018360208201111561030c57600080fd5b803590602001918460208302840111600160201b8311171561032d57600080fd5b509092509050610931565b6101c2610dd7565b6101c2610ddd565b6101dc610de3565b6101b8600480360360e081101561036657600080fd5b810190602081018135600160201b81111561038057600080fd5b82018360208201111561039257600080fd5b803590602001918460208302840111600160201b831117156103b357600080fd5b919350915080359060208101359060408101359060608101359060808101359060a00135610df2565b6101b860048036036101c08110156103f357600080fd5b810190602081018135600160201b81111561040d57600080fd5b82018360208201111561041f57600080fd5b803590602001918460208302840111600160201b8311171561044057600080fd5b6040805160608181018352949693958335956020850135959385013594818101359460808201359460a08301359460c084013594929390926101408201929160e00190600390839083908082843760009201919091525091949392602081019250359050600160201b8111156104b557600080fd5b8201836020820111156104c757600080fd5b803590602001918460018302840111600160201b831117156104e857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561053a57600080fd5b82018360208201111561054c57600080fd5b803590602001918460018302840111600160201b8311171561056d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff16915061106d9050565b6101dc6115ab565b6101c26115ba565b6101dc600480360360208110156105d957600080fd5b5035611600565b6000600a5460ff1660028111156105f357fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b815250906106a55760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561066a578181015183820152602001610652565b50505050905090810190601f1680156106975780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506106b260018b8b611e61565b50600380546001600160a01b038a81166001600160a01b0319928316179092556004889055600580548884169083161790556006805487841690831617905560088590556009849055600a8054600260ff199091168117909155600b8a90554360075580549091169183169190911790556040517f7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de90600090a150505050505050505050565b60095481565b6006546001600160a01b031681565b60005460ff165b90565b600061078e6007544361162790919063ffffffff16565b90506107986115ba565b81116040518060400160405280601081526020016f54494d454f55545f444541444c494e4560801b8152509061080f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b506001600a5460ff16600281111561082357fe5b141561085f576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a161085a611669565b610890565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a16108906116e7565b50565b600b5481565b60006001600a5460ff1660028111156108ae57fe5b14156108c657506005546001600160a01b0316610774565b6002600a5460ff1660028111156108d957fe5b14156108f157506006546001600160a01b0316610774565b6040805162461bcd60e51b81526020600482015260076024820152662727afaa2aa92760c91b604482015290519081900360640190fd5b600a5460ff1681565b610939610899565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b815250906109bc5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b506109c56115ba565b6007546109d3904390611627565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610a465760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b50600082826000198101818110610a5957fe5b9050602002013514610aa65760018611610aa6576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b604482015290519081900360640190fd5b610ab286610190611742565b6001018114610af4576040805162461bcd60e51b815260206004820152600960248201526810d55517d0d3d5539560ba1b604482015290519081900360640190fd5b8482826000198101818110610b0557fe5b905060200201351415610b4a576040805162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b604482015290519081900360640190fd5b610b54848461175a565b82826000818110610b6157fe5b9050602002013514610baf576040805162461bcd60e51b81526020600482015260126024820152717365676d656e74207072652d6669656c647360701b604482015290519081900360640190fd5b600082828281610bbb57fe5b905060200201351415610c09576040805162461bcd60e51b8152602060048201526011602482015270155394915050d21050931157d4d5105495607a1b604482015290519081900360640190fd5b610c138787611786565b8410610c5f576040805162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840e6cacedacadce840d8cadccee8d60531b604482015290519081900360640190fd5b6000610c80888885856000818110610c7357fe5b90506020020135896117d4565b9050610c8e818c8c8c611812565b610cce8383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508c92508b91506118ca9050565b50600b547f0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d8989868660405180858152602001848152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f191690920182900397509095505050505050a2506002600a5460ff166002811115610d5957fe5b1415610d9557610d80610d776007544361162790919063ffffffff16565b60095490611627565b600955600a805460ff19166001179055610dc7565b610db6610dad6007544361162790919063ffffffff16565b60085490611627565b600855600a805460ff191660021790555b5050436007555050505050505050565b60075481565b60085481565b6005546001600160a01b031681565b610dfa610899565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b81525090610e7d5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b50610e866115ba565b600754610e94904390611627565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610f075760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b506000610f14838361175a565b90506000610f24878784886117d4565b9050610f32818b8b8b611812565b610f3c8787611786565b841015610f7b576040805162461bcd60e51b81526020600482015260086024820152671393d517d0d3d39560c21b604482015290519081900360640190fd5b84821415610fbc576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6040517ff62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca90600090a1610fed611a07565b5060029050600a5460ff16600281111561100357fe5b141561103657611021610d776007544361162790919063ffffffff16565b600955600a805460ff1916600117905561105f565b61104e610dad6007544361162790919063ffffffff16565b600855600a805460ff191660021790555b505043600755505050505050565b611075610899565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b815250906110f85760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b506111016115ba565b60075461110f904390611627565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906111825760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b506000806000611190611ec4565b60018560ff16815481106111a057fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b0316639d16dd04600260009054906101000a90046001600160a01b03168d60405180604001604052808f81526020018e8152508b8b6040518663ffffffff1660e01b815260040180866001600160a01b0316815260200185815260200184600260200280838360005b8381101561124657818101518382015260200161122e565b505050509050018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561128a578181015183820152602001611272565b50505050905090810190601f1680156112b75780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156112ea5781810151838201526020016112d2565b50505050905090810190601f1680156113175780820380516001836020036101000a031916815260200191505b5097505050505050505060c06040518083038186803b15801561133957600080fd5b505afa15801561134d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c081101561137257600080fd5b50805160208201516004549195509350604090910191508211156113d1576040805162461bcd60e51b8152602060048201526011602482015270544f4f5f4d414e595f4d4553534147455360781b604482015290519081900360640190fd5b6113db8e8e611786565b88511061141a576040805162461bcd60e51b815260206004820152600860248201526713d4d417d0d3d39560c21b604482015290519081900360640190fd5b6114248e8e611786565b61144167ffffffffffffffff85168a60005b602002015190611786565b1015611480576040805162461bcd60e51b815260206004820152600960248201526813d4d417d4d213d49560ba1b604482015290519081900360640190fd5b61148e8a8a8a868686611a35565b8c14156114ce576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6114e68e8e6114e08e8e8e8e88611ab7565b8f6117d4565b93505050506114f7818f8f8f611812565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1611528611a07565b506002600a5460ff16600281111561153c57fe5b141561156f5761155a610d776007544361162790919063ffffffff16565b600955600a805460ff19166001179055611598565b611587610dad6007544361162790919063ffffffff16565b600855600a805460ff191660021790555b5050436007555050505050505050505050565b6002546001600160a01b031681565b60006001600a5460ff1660028111156115cf57fe5b14156115de5750600854610774565b6002600a5460ff1660028111156115f157fe5b14156108f15750600954610774565b6001818154811061160d57fe5b6000918252602090912001546001600160a01b0316905081565b60006116608383604051806040016040528060148152602001737375627472616374696f6e206f766572666c6f7760601b815250611ae8565b90505b92915050565b60035460065460055460408051637d3c01f360e11b81526001600160a01b039384166004820152918316602483015251919092169163fa7803e691604480830192600092919082900301818387803b1580156116c457600080fd5b505af11580156116d8573d6000803e3d6000fd5b505050506116e533611b42565b565b60035460055460065460408051637d3c01f360e11b81526001600160a01b039384166004820152918316602483015251919092169163fa7803e691604480830192600092919082900301818387803b1580156116c457600080fd5b600081831015611753575081611663565b5080611663565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600082820183811015611660576040805162461bcd60e51b81526020600482015260116024820152706164646974696f6e206f766572666c6f7760781b604482015290519081900360640190fd5b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b611852838380806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250859250889150611bc29050565b600b5414604051806040016040528060088152602001672124a9afa82922ab60c11b815250906118c35760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b5050505050565b82516000906000190160608167ffffffffffffffff811180156118ec57600080fd5b50604051908082528060200260200182016040528015611916578160200160208202803683370190505b50905060006119258584611c90565b9050600086905061196081838a60008151811061193e57fe5b60200260200101518b60018151811061195357fe5b60200260200101516117d4565b8360008151811061196d57fe5b60209081029190910101526119828183611786565b905061198e8685611cae565b915060015b848110156119ee576119c382848b84815181106119ac57fe5b60200260200101518c856001018151811061195357fe5b8482815181106119cf57fe5b60209081029190910101526119e48284611786565b9150600101611993565b506119f883611cc1565b600b5550929695505050505050565b6001600a5460ff166002811115611a1a57fe5b1415611a2d57611a286116e7565b6116e5565b6116e5611669565b6000611aac611a4f67ffffffffffffffff86168784611436565b60208401516040850151611aa7918791611a7e8d8214611a70576001611a73565b60005b60ff168c6001611436565b6060890151611aa28e8214611a94576001611a97565b60005b60ff168e6002611436565b611e16565b61175a565b979650505050505050565b8151815160208401516040850151600093611ade939092611aa7928b92918b918b90611e16565b9695505050505050565b60008184841115611b3a5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b505050900390565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff1615611bb55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561066a578181015183820152602001610652565b50806001600160a01b0316ff5b8251600090610100811115611bd657600080fd5b8260005b82811015611c865760028606611c3357868181518110611bf657fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150611c78565b81878281518110611c4057fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101611bda565b5095945050505050565b6000818381611c9b57fe5b06828481611ca557fe5b04019392505050565b6000818381611cb957fe5b049392505050565b6000815b600181511115611df95760606002825160010181611cdf57fe5b0467ffffffffffffffff81118015611cf657600080fd5b50604051908082528060200260200182016040528015611d20578160200160208202803683370190505b50905060005b8151811015611df1578251816002026001011015611db957828160020281518110611d4d57fe5b6020026020010151838260020260010181518110611d6757fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611da857fe5b602002602001018181525050611de9565b828160020281518110611dc857fe5b6020026020010151828281518110611ddc57fe5b6020026020010181815250505b600101611d26565b509050611cc5565b80600081518110611e0657fe5b6020026020010151915050919050565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b828054828255906000526020600020908101928215611eb4579160200282015b82811115611eb45781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190611e81565b50611ec0929150611ee2565b5090565b60405180608001604052806004906020820280368337509192915050565b5b80821115611ec05780546001600160a01b0319168155600101611ee356fea2646970667358221220ac0f4f971697748cf173b8c868d95c2c67692179a0ef9d96f982b4ba143154a464736f6c634300060c0033"

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
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// Asserter is a free data retrieval call binding the contract method 0xbb4af0b1.
//
// Solidity: function asserter() view returns(address)
func (_Challenge *ChallengeCaller) Asserter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "asserter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asserter is a free data retrieval call binding the contract method 0xbb4af0b1.
//
// Solidity: function asserter() view returns(address)
func (_Challenge *ChallengeSession) Asserter() (common.Address, error) {
	return _Challenge.Contract.Asserter(&_Challenge.CallOpts)
}

// Asserter is a free data retrieval call binding the contract method 0xbb4af0b1.
//
// Solidity: function asserter() view returns(address)
func (_Challenge *ChallengeCallerSession) Asserter() (common.Address, error) {
	return _Challenge.Contract.Asserter(&_Challenge.CallOpts)
}

// AsserterTimeLeft is a free data retrieval call binding the contract method 0x9a9e4f44.
//
// Solidity: function asserterTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCaller) AsserterTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "asserterTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AsserterTimeLeft is a free data retrieval call binding the contract method 0x9a9e4f44.
//
// Solidity: function asserterTimeLeft() view returns(uint256)
func (_Challenge *ChallengeSession) AsserterTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.AsserterTimeLeft(&_Challenge.CallOpts)
}

// AsserterTimeLeft is a free data retrieval call binding the contract method 0x9a9e4f44.
//
// Solidity: function asserterTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCallerSession) AsserterTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.AsserterTimeLeft(&_Challenge.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Challenge *ChallengeCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Challenge *ChallengeSession) Bridge() (common.Address, error) {
	return _Challenge.Contract.Bridge(&_Challenge.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Challenge *ChallengeCallerSession) Bridge() (common.Address, error) {
	return _Challenge.Contract.Bridge(&_Challenge.CallOpts)
}

// ChallengeState is a free data retrieval call binding the contract method 0x843d5a5c.
//
// Solidity: function challengeState() view returns(bytes32)
func (_Challenge *ChallengeCaller) ChallengeState(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challengeState")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChallengeState is a free data retrieval call binding the contract method 0x843d5a5c.
//
// Solidity: function challengeState() view returns(bytes32)
func (_Challenge *ChallengeSession) ChallengeState() ([32]byte, error) {
	return _Challenge.Contract.ChallengeState(&_Challenge.CallOpts)
}

// ChallengeState is a free data retrieval call binding the contract method 0x843d5a5c.
//
// Solidity: function challengeState() view returns(bytes32)
func (_Challenge *ChallengeCallerSession) ChallengeState() ([32]byte, error) {
	return _Challenge.Contract.ChallengeState(&_Challenge.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_Challenge *ChallengeCaller) Challenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_Challenge *ChallengeSession) Challenger() (common.Address, error) {
	return _Challenge.Contract.Challenger(&_Challenge.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_Challenge *ChallengeCallerSession) Challenger() (common.Address, error) {
	return _Challenge.Contract.Challenger(&_Challenge.CallOpts)
}

// ChallengerTimeLeft is a free data retrieval call binding the contract method 0x41e8510c.
//
// Solidity: function challengerTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCaller) ChallengerTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "challengerTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengerTimeLeft is a free data retrieval call binding the contract method 0x41e8510c.
//
// Solidity: function challengerTimeLeft() view returns(uint256)
func (_Challenge *ChallengeSession) ChallengerTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.ChallengerTimeLeft(&_Challenge.CallOpts)
}

// ChallengerTimeLeft is a free data retrieval call binding the contract method 0x41e8510c.
//
// Solidity: function challengerTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCallerSession) ChallengerTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.ChallengerTimeLeft(&_Challenge.CallOpts)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_Challenge *ChallengeCaller) CurrentResponder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "currentResponder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_Challenge *ChallengeSession) CurrentResponder() (common.Address, error) {
	return _Challenge.Contract.CurrentResponder(&_Challenge.CallOpts)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x8a8cd218.
//
// Solidity: function currentResponder() view returns(address)
func (_Challenge *ChallengeCallerSession) CurrentResponder() (common.Address, error) {
	return _Challenge.Contract.CurrentResponder(&_Challenge.CallOpts)
}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCaller) CurrentResponderTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "currentResponderTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeSession) CurrentResponderTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.CurrentResponderTimeLeft(&_Challenge.CallOpts)
}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xe87e3589.
//
// Solidity: function currentResponderTimeLeft() view returns(uint256)
func (_Challenge *ChallengeCallerSession) CurrentResponderTimeLeft() (*big.Int, error) {
	return _Challenge.Contract.CurrentResponderTimeLeft(&_Challenge.CallOpts)
}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_Challenge *ChallengeCaller) Executors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "executors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_Challenge *ChallengeSession) Executors(arg0 *big.Int) (common.Address, error) {
	return _Challenge.Contract.Executors(&_Challenge.CallOpts, arg0)
}

// Executors is a free data retrieval call binding the contract method 0xf97a05df.
//
// Solidity: function executors(uint256 ) view returns(address)
func (_Challenge *ChallengeCallerSession) Executors(arg0 *big.Int) (common.Address, error) {
	return _Challenge.Contract.Executors(&_Challenge.CallOpts, arg0)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Challenge *ChallengeCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Challenge *ChallengeSession) IsMaster() (bool, error) {
	return _Challenge.Contract.IsMaster(&_Challenge.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_Challenge *ChallengeCallerSession) IsMaster() (bool, error) {
	return _Challenge.Contract.IsMaster(&_Challenge.CallOpts)
}

// LastMoveBlock is a free data retrieval call binding the contract method 0x925f9a96.
//
// Solidity: function lastMoveBlock() view returns(uint256)
func (_Challenge *ChallengeCaller) LastMoveBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "lastMoveBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMoveBlock is a free data retrieval call binding the contract method 0x925f9a96.
//
// Solidity: function lastMoveBlock() view returns(uint256)
func (_Challenge *ChallengeSession) LastMoveBlock() (*big.Int, error) {
	return _Challenge.Contract.LastMoveBlock(&_Challenge.CallOpts)
}

// LastMoveBlock is a free data retrieval call binding the contract method 0x925f9a96.
//
// Solidity: function lastMoveBlock() view returns(uint256)
func (_Challenge *ChallengeCallerSession) LastMoveBlock() (*big.Int, error) {
	return _Challenge.Contract.LastMoveBlock(&_Challenge.CallOpts)
}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() view returns(uint8)
func (_Challenge *ChallengeCaller) Turn(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "turn")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() view returns(uint8)
func (_Challenge *ChallengeSession) Turn() (uint8, error) {
	return _Challenge.Contract.Turn(&_Challenge.CallOpts)
}

// Turn is a free data retrieval call binding the contract method 0x8b299903.
//
// Solidity: function turn() view returns(uint8)
func (_Challenge *ChallengeCallerSession) Turn() (uint8, error) {
	return _Challenge.Contract.Turn(&_Challenge.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0x8e7b84c5.
//
// Solidity: function bisectExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _gasUsedBefore, bytes32 _assertionRest, bytes32[] _chainHashes) returns()
func (_Challenge *ChallengeTransactor) BisectExecution(opts *bind.TransactOpts, _merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _gasUsedBefore *big.Int, _assertionRest [32]byte, _chainHashes [][32]byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "bisectExecution", _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _gasUsedBefore, _assertionRest, _chainHashes)
}

// BisectExecution is a paid mutator transaction binding the contract method 0x8e7b84c5.
//
// Solidity: function bisectExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _gasUsedBefore, bytes32 _assertionRest, bytes32[] _chainHashes) returns()
func (_Challenge *ChallengeSession) BisectExecution(_merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _gasUsedBefore *big.Int, _assertionRest [32]byte, _chainHashes [][32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.BisectExecution(&_Challenge.TransactOpts, _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _gasUsedBefore, _assertionRest, _chainHashes)
}

// BisectExecution is a paid mutator transaction binding the contract method 0x8e7b84c5.
//
// Solidity: function bisectExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _gasUsedBefore, bytes32 _assertionRest, bytes32[] _chainHashes) returns()
func (_Challenge *ChallengeTransactorSession) BisectExecution(_merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _gasUsedBefore *big.Int, _assertionRest [32]byte, _chainHashes [][32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.BisectExecution(&_Challenge.TransactOpts, _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _gasUsedBefore, _assertionRest, _chainHashes)
}

// InitializeChallenge is a paid mutator transaction binding the contract method 0x32f8c24f.
//
// Solidity: function initializeChallenge(address[] _executors, address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _bridge) returns()
func (_Challenge *ChallengeTransactor) InitializeChallenge(opts *bind.TransactOpts, _executors []common.Address, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _bridge common.Address) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "initializeChallenge", _executors, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _bridge)
}

// InitializeChallenge is a paid mutator transaction binding the contract method 0x32f8c24f.
//
// Solidity: function initializeChallenge(address[] _executors, address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _bridge) returns()
func (_Challenge *ChallengeSession) InitializeChallenge(_executors []common.Address, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _bridge common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.InitializeChallenge(&_Challenge.TransactOpts, _executors, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _bridge)
}

// InitializeChallenge is a paid mutator transaction binding the contract method 0x32f8c24f.
//
// Solidity: function initializeChallenge(address[] _executors, address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _bridge) returns()
func (_Challenge *ChallengeTransactorSession) InitializeChallenge(_executors []common.Address, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _bridge common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.InitializeChallenge(&_Challenge.TransactOpts, _executors, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _bridge)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xe08f819e.
//
// Solidity: function oneStepProveExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _initialMessagesRead, bytes32 _initialSendAcc, bytes32 _initialLogAcc, uint256[3] _initialState, bytes _executionProof, bytes _bufferProof, uint8 prover) returns()
func (_Challenge *ChallengeTransactor) OneStepProveExecution(opts *bind.TransactOpts, _merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _initialMessagesRead *big.Int, _initialSendAcc [32]byte, _initialLogAcc [32]byte, _initialState [3]*big.Int, _executionProof []byte, _bufferProof []byte, prover uint8) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "oneStepProveExecution", _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _initialMessagesRead, _initialSendAcc, _initialLogAcc, _initialState, _executionProof, _bufferProof, prover)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xe08f819e.
//
// Solidity: function oneStepProveExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _initialMessagesRead, bytes32 _initialSendAcc, bytes32 _initialLogAcc, uint256[3] _initialState, bytes _executionProof, bytes _bufferProof, uint8 prover) returns()
func (_Challenge *ChallengeSession) OneStepProveExecution(_merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _initialMessagesRead *big.Int, _initialSendAcc [32]byte, _initialLogAcc [32]byte, _initialState [3]*big.Int, _executionProof []byte, _bufferProof []byte, prover uint8) (*types.Transaction, error) {
	return _Challenge.Contract.OneStepProveExecution(&_Challenge.TransactOpts, _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _initialMessagesRead, _initialSendAcc, _initialLogAcc, _initialState, _executionProof, _bufferProof, prover)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xe08f819e.
//
// Solidity: function oneStepProveExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _initialMessagesRead, bytes32 _initialSendAcc, bytes32 _initialLogAcc, uint256[3] _initialState, bytes _executionProof, bytes _bufferProof, uint8 prover) returns()
func (_Challenge *ChallengeTransactorSession) OneStepProveExecution(_merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _initialMessagesRead *big.Int, _initialSendAcc [32]byte, _initialLogAcc [32]byte, _initialState [3]*big.Int, _executionProof []byte, _bufferProof []byte, prover uint8) (*types.Transaction, error) {
	return _Challenge.Contract.OneStepProveExecution(&_Challenge.TransactOpts, _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _initialMessagesRead, _initialSendAcc, _initialLogAcc, _initialState, _executionProof, _bufferProof, prover)
}

// ProveContinuedExecution is a paid mutator transaction binding the contract method 0xdeda4115.
//
// Solidity: function proveContinuedExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _gasUsedBefore, bytes32 _assertionRest) returns()
func (_Challenge *ChallengeTransactor) ProveContinuedExecution(opts *bind.TransactOpts, _merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _gasUsedBefore *big.Int, _assertionRest [32]byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "proveContinuedExecution", _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _gasUsedBefore, _assertionRest)
}

// ProveContinuedExecution is a paid mutator transaction binding the contract method 0xdeda4115.
//
// Solidity: function proveContinuedExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _gasUsedBefore, bytes32 _assertionRest) returns()
func (_Challenge *ChallengeSession) ProveContinuedExecution(_merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _gasUsedBefore *big.Int, _assertionRest [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.ProveContinuedExecution(&_Challenge.TransactOpts, _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _gasUsedBefore, _assertionRest)
}

// ProveContinuedExecution is a paid mutator transaction binding the contract method 0xdeda4115.
//
// Solidity: function proveContinuedExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _gasUsedBefore, bytes32 _assertionRest) returns()
func (_Challenge *ChallengeTransactorSession) ProveContinuedExecution(_merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _gasUsedBefore *big.Int, _assertionRest [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.ProveContinuedExecution(&_Challenge.TransactOpts, _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _gasUsedBefore, _assertionRest)
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Challenge *ChallengeTransactor) Timeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "timeout")
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Challenge *ChallengeSession) Timeout() (*types.Transaction, error) {
	return _Challenge.Contract.Timeout(&_Challenge.TransactOpts)
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Challenge *ChallengeTransactorSession) Timeout() (*types.Transaction, error) {
	return _Challenge.Contract.Timeout(&_Challenge.TransactOpts)
}

// ChallengeAsserterTimedOutIterator is returned from FilterAsserterTimedOut and is used to iterate over the raw logs and unpacked data for AsserterTimedOut events raised by the Challenge contract.
type ChallengeAsserterTimedOutIterator struct {
	Event *ChallengeAsserterTimedOut // Event containing the contract specifics and raw log

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
func (it *ChallengeAsserterTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeAsserterTimedOut)
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
		it.Event = new(ChallengeAsserterTimedOut)
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
func (it *ChallengeAsserterTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeAsserterTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeAsserterTimedOut represents a AsserterTimedOut event raised by the Challenge contract.
type ChallengeAsserterTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAsserterTimedOut is a free log retrieval operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_Challenge *ChallengeFilterer) FilterAsserterTimedOut(opts *bind.FilterOpts) (*ChallengeAsserterTimedOutIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return &ChallengeAsserterTimedOutIterator{contract: _Challenge.contract, event: "AsserterTimedOut", logs: logs, sub: sub}, nil
}

// WatchAsserterTimedOut is a free log subscription operation binding the contract event 0x2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f0.
//
// Solidity: event AsserterTimedOut()
func (_Challenge *ChallengeFilterer) WatchAsserterTimedOut(opts *bind.WatchOpts, sink chan<- *ChallengeAsserterTimedOut) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "AsserterTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeAsserterTimedOut)
				if err := _Challenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
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
func (_Challenge *ChallengeFilterer) ParseAsserterTimedOut(log types.Log) (*ChallengeAsserterTimedOut, error) {
	event := new(ChallengeAsserterTimedOut)
	if err := _Challenge.contract.UnpackLog(event, "AsserterTimedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the Challenge contract.
type ChallengeBisectedIterator struct {
	Event *ChallengeBisected // Event containing the contract specifics and raw log

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
func (it *ChallengeBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeBisected)
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
		it.Event = new(ChallengeBisected)
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
func (it *ChallengeBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeBisected represents a Bisected event raised by the Challenge contract.
type ChallengeBisected struct {
	ChallengeRoot           [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	ChainHashes             [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d.
//
// Solidity: event Bisected(bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_Challenge *ChallengeFilterer) FilterBisected(opts *bind.FilterOpts, challengeRoot [][32]byte) (*ChallengeBisectedIterator, error) {

	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "Bisected", challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeBisectedIterator{contract: _Challenge.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d.
//
// Solidity: event Bisected(bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_Challenge *ChallengeFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *ChallengeBisected, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "Bisected", challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeBisected)
				if err := _Challenge.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d.
//
// Solidity: event Bisected(bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_Challenge *ChallengeFilterer) ParseBisected(log types.Log) (*ChallengeBisected, error) {
	event := new(ChallengeBisected)
	if err := _Challenge.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeChallengerTimedOutIterator is returned from FilterChallengerTimedOut and is used to iterate over the raw logs and unpacked data for ChallengerTimedOut events raised by the Challenge contract.
type ChallengeChallengerTimedOutIterator struct {
	Event *ChallengeChallengerTimedOut // Event containing the contract specifics and raw log

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
func (it *ChallengeChallengerTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeChallengerTimedOut)
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
		it.Event = new(ChallengeChallengerTimedOut)
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
func (it *ChallengeChallengerTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeChallengerTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeChallengerTimedOut represents a ChallengerTimedOut event raised by the Challenge contract.
type ChallengeChallengerTimedOut struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengerTimedOut is a free log retrieval operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_Challenge *ChallengeFilterer) FilterChallengerTimedOut(opts *bind.FilterOpts) (*ChallengeChallengerTimedOutIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return &ChallengeChallengerTimedOutIterator{contract: _Challenge.contract, event: "ChallengerTimedOut", logs: logs, sub: sub}, nil
}

// WatchChallengerTimedOut is a free log subscription operation binding the contract event 0x4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a.
//
// Solidity: event ChallengerTimedOut()
func (_Challenge *ChallengeFilterer) WatchChallengerTimedOut(opts *bind.WatchOpts, sink chan<- *ChallengeChallengerTimedOut) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "ChallengerTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeChallengerTimedOut)
				if err := _Challenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
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
func (_Challenge *ChallengeFilterer) ParseChallengerTimedOut(log types.Log) (*ChallengeChallengerTimedOut, error) {
	event := new(ChallengeChallengerTimedOut)
	if err := _Challenge.contract.UnpackLog(event, "ChallengerTimedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeContinuedExecutionProvenIterator is returned from FilterContinuedExecutionProven and is used to iterate over the raw logs and unpacked data for ContinuedExecutionProven events raised by the Challenge contract.
type ChallengeContinuedExecutionProvenIterator struct {
	Event *ChallengeContinuedExecutionProven // Event containing the contract specifics and raw log

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
func (it *ChallengeContinuedExecutionProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeContinuedExecutionProven)
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
		it.Event = new(ChallengeContinuedExecutionProven)
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
func (it *ChallengeContinuedExecutionProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeContinuedExecutionProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeContinuedExecutionProven represents a ContinuedExecutionProven event raised by the Challenge contract.
type ChallengeContinuedExecutionProven struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContinuedExecutionProven is a free log retrieval operation binding the contract event 0xf62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca.
//
// Solidity: event ContinuedExecutionProven()
func (_Challenge *ChallengeFilterer) FilterContinuedExecutionProven(opts *bind.FilterOpts) (*ChallengeContinuedExecutionProvenIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "ContinuedExecutionProven")
	if err != nil {
		return nil, err
	}
	return &ChallengeContinuedExecutionProvenIterator{contract: _Challenge.contract, event: "ContinuedExecutionProven", logs: logs, sub: sub}, nil
}

// WatchContinuedExecutionProven is a free log subscription operation binding the contract event 0xf62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca.
//
// Solidity: event ContinuedExecutionProven()
func (_Challenge *ChallengeFilterer) WatchContinuedExecutionProven(opts *bind.WatchOpts, sink chan<- *ChallengeContinuedExecutionProven) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "ContinuedExecutionProven")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeContinuedExecutionProven)
				if err := _Challenge.contract.UnpackLog(event, "ContinuedExecutionProven", log); err != nil {
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

// ParseContinuedExecutionProven is a log parse operation binding the contract event 0xf62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca.
//
// Solidity: event ContinuedExecutionProven()
func (_Challenge *ChallengeFilterer) ParseContinuedExecutionProven(log types.Log) (*ChallengeContinuedExecutionProven, error) {
	event := new(ChallengeContinuedExecutionProven)
	if err := _Challenge.contract.UnpackLog(event, "ContinuedExecutionProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the Challenge contract.
type ChallengeInitiatedChallengeIterator struct {
	Event *ChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeInitiatedChallenge)
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
		it.Event = new(ChallengeInitiatedChallenge)
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
func (it *ChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the Challenge contract.
type ChallengeInitiatedChallenge struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de.
//
// Solidity: event InitiatedChallenge()
func (_Challenge *ChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ChallengeInitiatedChallengeIterator{contract: _Challenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de.
//
// Solidity: event InitiatedChallenge()
func (_Challenge *ChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeInitiatedChallenge)
				if err := _Challenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de.
//
// Solidity: event InitiatedChallenge()
func (_Challenge *ChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ChallengeInitiatedChallenge, error) {
	event := new(ChallengeInitiatedChallenge)
	if err := _Challenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the Challenge contract.
type ChallengeOneStepProofCompletedIterator struct {
	Event *ChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeOneStepProofCompleted)
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
		it.Event = new(ChallengeOneStepProofCompleted)
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
func (it *ChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the Challenge contract.
type ChallengeOneStepProofCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_Challenge *ChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*ChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &ChallengeOneStepProofCompletedIterator{contract: _Challenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_Challenge *ChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeOneStepProofCompleted)
				if err := _Challenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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
func (_Challenge *ChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*ChallengeOneStepProofCompleted, error) {
	event := new(ChallengeOneStepProofCompleted)
	if err := _Challenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
