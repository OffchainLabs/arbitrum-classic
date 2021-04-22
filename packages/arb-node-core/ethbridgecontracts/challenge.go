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
const ChallengeABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AsserterTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChallengerTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContinuedExecutionProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"asserter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asserterTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_merkleRoute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_oldEndHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_gasUsedBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionRest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridges\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengerTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentResponderTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"executors\",\"outputs\":[{\"internalType\":\"contractIOneStepProof\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOneStepProof[]\",\"name\":\"_executors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_resultReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_executionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asserter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_asserterTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengerTimeLeft\",\"type\":\"uint256\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"_sequencerBridge\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"_delayedBridge\",\"type\":\"address\"}],\"name\":\"initializeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMoveBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_merkleRoute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_oldEndHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_initialMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[2]\",\"name\":\"_initialAccs\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_initialState\",\"type\":\"uint256[3]\"},{\"internalType\":\"bytes\",\"name\":\"_executionProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_bufferProof\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"prover\",\"type\":\"uint8\"}],\"name\":\"oneStepProveExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_merkleRoute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengedSegmentLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_oldEndHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_gasUsedBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionRest\",\"type\":\"bytes32\"}],\"name\":\"proveContinuedExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"turn\",\"outputs\":[{\"internalType\":\"enumChallenge.Turn\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFuncSigs = map[string]string{
	"bb4af0b1": "asserter()",
	"9a9e4f44": "asserterTimeLeft()",
	"8e7b84c5": "bisectExecution(bytes32[],uint256,uint256,uint256,bytes32,uint256,bytes32,bytes32[])",
	"a3c44705": "bridges(uint256)",
	"843d5a5c": "challengeState()",
	"534db0e2": "challenger()",
	"41e8510c": "challengerTimeLeft()",
	"8a8cd218": "currentResponder()",
	"e87e3589": "currentResponderTimeLeft()",
	"f97a05df": "executors(uint256)",
	"e0d42b8e": "initializeChallenge(address[],address,bytes32,uint256,address,address,uint256,uint256,address,address)",
	"6f791d29": "isMaster()",
	"925f9a96": "lastMoveBlock()",
	"95979201": "oneStepProveExecution(bytes32[],uint256,uint256,uint256,bytes32,uint256,bytes32[2],uint256[3],bytes,bytes,uint8)",
	"deda4115": "proveContinuedExecution(bytes32[],uint256,uint256,uint256,bytes32,uint256,bytes32)",
	"70dea79a": "timeout()",
	"8b299903": "turn()",
}

// ChallengeBin is the compiled bytecode used for deploying new contracts.
var ChallengeBin = "0x608060405234801561001057600080fd5b506000805460ff19166001179055611fbc8061002d6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063925f9a96116100a2578063bb4af0b111610071578063bb4af0b11461048e578063deda411514610496578063e0d42b8e14610522578063e87e3589146105d5578063f97a05df146105dd5761010b565b8063925f9a961461029057806395979201146102985780639a9e4f4414610469578063a3c44705146104715761010b565b8063843d5a5c116100de578063843d5a5c146101745780638a8cd2181461017c5780638b299903146101845780638e7b84c5146101ad5761010b565b806341e8510c14610110578063534db0e21461012a5780636f791d291461014e57806370dea79a1461016a575b600080fd5b6101186105fa565b60408051918252519081900360200190f35b610132610600565b604080516001600160a01b039092168252519081900360200190f35b61015661060f565b604080519115158252519081900360200190f35b610172610619565b005b610118610772565b610132610778565b61018c610807565b6040518082600281111561019c57fe5b815260200191505060405180910390f35b61017260048036036101008110156101c457600080fd5b810190602081018135600160201b8111156101de57600080fd5b8201836020820111156101f057600080fd5b803590602001918460208302840111600160201b8311171561021157600080fd5b9193909282359260208101359260408201359260608301359260808101359260a082013592909160e081019060c00135600160201b81111561025257600080fd5b82018360208201111561026457600080fd5b803590602001918460208302840111600160201b8311171561028557600080fd5b509092509050610810565b610118610cb6565b61017260048036036101c08110156102af57600080fd5b810190602081018135600160201b8111156102c957600080fd5b8201836020820111156102db57600080fd5b803590602001918460208302840111600160201b831117156102fc57600080fd5b6040805160608181018352949693958335956020850135959385013594818101359460808201359460a08301949193919261014081019260e090910190600390839083908082843760009201919091525091949392602081019250359050600160201b81111561036b57600080fd5b82018360208201111561037d57600080fd5b803590602001918460018302840111600160201b8311171561039e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103f057600080fd5b82018360208201111561040257600080fd5b803590602001918460018302840111600160201b8311171561042357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff169150610cbc9050565b6101186111da565b6101326004803603602081101561048757600080fd5b50356111e0565b6101326111fd565b610172600480360360e08110156104ac57600080fd5b810190602081018135600160201b8111156104c657600080fd5b8201836020820111156104d857600080fd5b803590602001918460208302840111600160201b831117156104f957600080fd5b919350915080359060208101359060408101359060608101359060808101359060a0013561120c565b610172600480360361014081101561053957600080fd5b810190602081018135600160201b81111561055357600080fd5b82018360208201111561056557600080fd5b803590602001918460208302840111600160201b8311171561058657600080fd5b91935091506001600160a01b0381358116916020810135916040820135916060810135821691608082013581169160a08101359160c08201359160e08101358216916101009091013516611487565b6101186115d8565b610132600480360360208110156105f357600080fd5b503561161e565b600a5481565b6007546001600160a01b031681565b60005460ff165b90565b60006106306008544361164590919063ffffffff16565b905061063a6115d8565b81116040518060400160405280601081526020016f54494d454f55545f444541444c494e4560801b815250906106ee5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156106b357818101518382015260200161069b565b50505050905090810190601f1680156106e05780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506001600b5460ff16600281111561070257fe5b141561073e576040517f2b92a4b014281aa2424baba9ea60bf4f26833d1c1fbd873e51cd1a6caeef48f090600090a1610739611687565b61076f565b6040517f4e1f1f06cf69d199fcdb4d87a5a92d5248ca6b540e9fc2d3698927c5002a236a90600090a161076f611706565b50565b600c5481565b60006001600b5460ff16600281111561078d57fe5b14156107a557506006546001600160a01b0316610616565b6002600b5460ff1660028111156107b857fe5b14156107d057506007546001600160a01b0316610616565b6040805162461bcd60e51b81526020600482015260076024820152662727afaa2aa92760c91b604482015290519081900360640190fd5b600b5460ff1681565b610818610778565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b8152509061089b5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b506108a46115d8565b6008546108b2904390611645565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906109255760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b5060008282600019810181811061093857fe5b90506020020135146109855760018611610985576040805162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b604482015290519081900360640190fd5b61099186610190611762565b60010181146109d3576040805162461bcd60e51b815260206004820152600960248201526810d55517d0d3d5539560ba1b604482015290519081900360640190fd5b84828260001981018181106109e457fe5b905060200201351415610a29576040805162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b604482015290519081900360640190fd5b610a33848461177a565b82826000818110610a4057fe5b9050602002013514610a8e576040805162461bcd60e51b81526020600482015260126024820152717365676d656e74207072652d6669656c647360701b604482015290519081900360640190fd5b600082828281610a9a57fe5b905060200201351415610ae8576040805162461bcd60e51b8152602060048201526011602482015270155394915050d21050931157d4d5105495607a1b604482015290519081900360640190fd5b610af287876117a6565b8410610b3e576040805162461bcd60e51b81526020600482015260166024820152750d2dcecc2d8d2c840e6cacedacadce840d8cadccee8d60531b604482015290519081900360640190fd5b6000610b5f888885856000818110610b5257fe5b90506020020135896117f4565b9050610b6d818c8c8c611832565b610bad8383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508c92508b91506118ea9050565b50600c547f0a2bdfea671da507e80b0cbae49dd25100a5bdacc5dff43a9163a3fcbd7c3c7d8989868660405180858152602001848152602001806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f191690920182900397509095505050505050a2506002600b5460ff166002811115610c3857fe5b1415610c7457610c5f610c566008544361164590919063ffffffff16565b600a5490611645565b600a55600b805460ff19166001179055610ca6565b610c95610c8c6008544361164590919063ffffffff16565b60095490611645565b600955600b805460ff191660021790555b5050436008555050505050505050565b60085481565b610cc4610778565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b81525090610d475760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b50610d506115d8565b600854610d5e904390611645565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b81525090610dd15760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b506000806000610ddf611e9e565b60018560ff1681548110610def57fe5b6000918252602090912001546040516323eed0eb60e11b81526001600160a01b03909116906347dda1d6906002908d908d908c908c90600481019060440186825b81546001600160a01b03168152600190910190602001808311610e3057505085815260200184604080828437600081840152601f19601f8201169050808301925050508060200180602001838103835285818151815260200191508051906020019080838360005b83811015610eb0578181015183820152602001610e98565b50505050905090810190601f168015610edd5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610f10578181015183820152602001610ef8565b50505050905090810190601f168015610f3d5780820380516001836020036101000a031916815260200191505b5097505050505050505060c06040518083038186803b158015610f5f57600080fd5b505afa158015610f73573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060c0811015610f9857600080fd5b5080516020820151600554919550935060409091019150821115610ff7576040805162461bcd60e51b8152602060048201526011602482015270544f4f5f4d414e595f4d4553534147455360781b604482015290519081900360640190fd5b6110018d8d6117a6565b885110611040576040805162461bcd60e51b815260206004820152600860248201526713d4d417d0d3d39560c21b604482015290519081900360640190fd5b61104a8d8d6117a6565b61106767ffffffffffffffff85168a60005b6020020151906117a6565b10156110a6576040805162461bcd60e51b815260206004820152600960248201526813d4d417d4d213d49560ba1b604482015290519081900360640190fd5b6110b9893560208b01358a868686611a27565b8b14156110f9576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6111168d8d6111108d8d3560208f01358e88611ac6565b8e6117f4565b9350505050611127818e8e8e611832565b6040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1611158611af7565b506002600b5460ff16600281111561116c57fe5b141561119f5761118a610c566008544361164590919063ffffffff16565b600a55600b805460ff191660011790556111c8565b6111b7610c8c6008544361164590919063ffffffff16565b600955600b805460ff191660021790555b50504360085550505050505050505050565b60095481565b600281600281106111ed57fe5b01546001600160a01b0316905081565b6006546001600160a01b031681565b611214610778565b6001600160a01b0316336001600160a01b0316146040518060400160405280600a8152602001692124a9afa9a2a72222a960b11b815250906112975760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b506112a06115d8565b6008546112ae904390611645565b11156040518060400160405280600c81526020016b4249535f444541444c494e4560a01b815250906113215760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b50600061132e838361177a565b9050600061133e878784886117f4565b905061134c818b8b8b611832565b61135687876117a6565b841015611395576040805162461bcd60e51b81526020600482015260086024820152671393d517d0d3d39560c21b604482015290519081900360640190fd5b848214156113d6576040805162461bcd60e51b815260206004820152600960248201526815d493d391d7d1539160ba1b604482015290519081900360640190fd5b6040517ff62bb8ab32072c0ea3337f57276b8e66418eca0dfcc5e3b8aef4905d43e8f8ca90600090a1611407611af7565b5060029050600b5460ff16600281111561141d57fe5b14156114505761143b610c566008544361164590919063ffffffff16565b600a55600b805460ff19166001179055611479565b611468610c8c6008544361164590919063ffffffff16565b600955600b805460ff191660021790555b505043600855505050505050565b6000600b5460ff16600281111561149a57fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b8152509061150f5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b5061151c60018c8c611ebc565b50600480546001600160a01b03199081166001600160a01b038c8116919091179092556005899055600680548216898416179055600780549091168783161790556009859055600a849055600b805460ff19166002908117909155600c8a90554360085560408051808201909152848316815291831660208301526115a19181611f1f565b506040517f7003482dc89fcecb9f14e280f21ee716bd54187f7f3b0ab5ed78f3648218f2de90600090a15050505050505050505050565b60006001600b5460ff1660028111156115ed57fe5b14156115fc5750600954610616565b6002600b5460ff16600281111561160f57fe5b14156107d05750600a54610616565b6001818154811061162b57fe5b6000918252602090912001546001600160a01b0316905081565b600061167e8383604051806040016040528060148152602001737375627472616374696f6e206f766572666c6f7760601b815250611b25565b90505b92915050565b6004805460075460065460408051637d3c01f360e11b81526001600160a01b039384169581019590955290821660248501525191169163fa7803e691604480830192600092919082900301818387803b1580156116e357600080fd5b505af11580156116f7573d6000803e3d6000fd5b5050505061170433611b7f565b565b6004805460065460075460408051637d3c01f360e11b81526001600160a01b039384169581019590955290821660248501525191169163fa7803e691604480830192600092919082900301818387803b1580156116e357600080fd5b600081831015611773575081611681565b5080611681565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008282018381101561167e576040805162461bcd60e51b81526020600482015260116024820152706164646974696f6e206f766572666c6f7760781b604482015290519081900360640190fd5b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b611872838380806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250859250889150611bff9050565b600c5414604051806040016040528060088152602001672124a9afa82922ab60c11b815250906118e35760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b5050505050565b82516000906000190160608167ffffffffffffffff8111801561190c57600080fd5b50604051908082528060200260200182016040528015611936578160200160208202803683370190505b50905060006119458584611ccd565b9050600086905061198081838a60008151811061195e57fe5b60200260200101518b60018151811061197357fe5b60200260200101516117f4565b8360008151811061198d57fe5b60209081029190910101526119a281836117a6565b90506119ae8685611ceb565b915060015b84811015611a0e576119e382848b84815181106119cc57fe5b60200260200101518c856001018151811061197357fe5b8482815181106119ef57fe5b6020908102919091010152611a0482846117a6565b91506001016119b3565b50611a1883611cfe565b600c5550929695505050505050565b600080611a5083600260200201518914611a42576001611a45565b60005b60ff1687600161105c565b90506000611a7a84600360200201518914611a6c576001611a6f565b60005b60ff1688600261105c565b9050611ab9611a9567ffffffffffffffff881689600061105c565b602086015160408701516060880151611ab4928a929091889088611e53565b61177a565b9998505050505050505050565b8151815160208401516040850151600093611aed939092611ab4928b92918b918b90611e53565b9695505050505050565b6001600b5460ff166002811115611b0a57fe5b1415611b1d57611b18611706565b611704565b611704611687565b60008184841115611b775760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b505050900390565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff1615611bf25760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106b357818101518382015260200161069b565b50806001600160a01b0316ff5b8251600090610100811115611c1357600080fd5b8260005b82811015611cc35760028606611c7057868181518110611c3357fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150611cb5565b81878281518110611c7d57fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101611c17565b5095945050505050565b6000818381611cd857fe5b06828481611ce257fe5b04019392505050565b6000818381611cf657fe5b049392505050565b6000815b600181511115611e365760606002825160010181611d1c57fe5b0467ffffffffffffffff81118015611d3357600080fd5b50604051908082528060200260200182016040528015611d5d578160200160208202803683370190505b50905060005b8151811015611e2e578251816002026001011015611df657828160020281518110611d8a57fe5b6020026020010151838260020260010181518110611da457fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120828281518110611de557fe5b602002602001018181525050611e26565b828160020281518110611e0557fe5b6020026020010151828281518110611e1957fe5b6020026020010181815250505b600101611d63565b509050611d02565b80600081518110611e4357fe5b6020026020010151915050919050565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b60405180608001604052806004906020820280368337509192915050565b828054828255906000526020600020908101928215611f0f579160200282015b82811115611f0f5781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190611edc565b50611f1b929150611f67565b5090565b8260028101928215611f0f579160200282015b82811115611f0f57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611f32565b5b80821115611f1b5780546001600160a01b0319168155600101611f6856fea264697066735822122050c094e5aa58c51208402d38fd44966acb5ae902afca19afa6775b1110128e5e64736f6c634300060c0033"

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

// Bridges is a free data retrieval call binding the contract method 0xa3c44705.
//
// Solidity: function bridges(uint256 ) view returns(address)
func (_Challenge *ChallengeCaller) Bridges(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Challenge.contract.Call(opts, &out, "bridges", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridges is a free data retrieval call binding the contract method 0xa3c44705.
//
// Solidity: function bridges(uint256 ) view returns(address)
func (_Challenge *ChallengeSession) Bridges(arg0 *big.Int) (common.Address, error) {
	return _Challenge.Contract.Bridges(&_Challenge.CallOpts, arg0)
}

// Bridges is a free data retrieval call binding the contract method 0xa3c44705.
//
// Solidity: function bridges(uint256 ) view returns(address)
func (_Challenge *ChallengeCallerSession) Bridges(arg0 *big.Int) (common.Address, error) {
	return _Challenge.Contract.Bridges(&_Challenge.CallOpts, arg0)
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

// InitializeChallenge is a paid mutator transaction binding the contract method 0xe0d42b8e.
//
// Solidity: function initializeChallenge(address[] _executors, address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns()
func (_Challenge *ChallengeTransactor) InitializeChallenge(opts *bind.TransactOpts, _executors []common.Address, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "initializeChallenge", _executors, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// InitializeChallenge is a paid mutator transaction binding the contract method 0xe0d42b8e.
//
// Solidity: function initializeChallenge(address[] _executors, address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns()
func (_Challenge *ChallengeSession) InitializeChallenge(_executors []common.Address, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.InitializeChallenge(&_Challenge.TransactOpts, _executors, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// InitializeChallenge is a paid mutator transaction binding the contract method 0xe0d42b8e.
//
// Solidity: function initializeChallenge(address[] _executors, address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, address _sequencerBridge, address _delayedBridge) returns()
func (_Challenge *ChallengeTransactorSession) InitializeChallenge(_executors []common.Address, _resultReceiver common.Address, _executionHash [32]byte, _maxMessageCount *big.Int, _asserter common.Address, _challenger common.Address, _asserterTimeLeft *big.Int, _challengerTimeLeft *big.Int, _sequencerBridge common.Address, _delayedBridge common.Address) (*types.Transaction, error) {
	return _Challenge.Contract.InitializeChallenge(&_Challenge.TransactOpts, _executors, _resultReceiver, _executionHash, _maxMessageCount, _asserter, _challenger, _asserterTimeLeft, _challengerTimeLeft, _sequencerBridge, _delayedBridge)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0x95979201.
//
// Solidity: function oneStepProveExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _initialMessagesRead, bytes32[2] _initialAccs, uint256[3] _initialState, bytes _executionProof, bytes _bufferProof, uint8 prover) returns()
func (_Challenge *ChallengeTransactor) OneStepProveExecution(opts *bind.TransactOpts, _merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _initialMessagesRead *big.Int, _initialAccs [2][32]byte, _initialState [3]*big.Int, _executionProof []byte, _bufferProof []byte, prover uint8) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "oneStepProveExecution", _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _initialMessagesRead, _initialAccs, _initialState, _executionProof, _bufferProof, prover)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0x95979201.
//
// Solidity: function oneStepProveExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _initialMessagesRead, bytes32[2] _initialAccs, uint256[3] _initialState, bytes _executionProof, bytes _bufferProof, uint8 prover) returns()
func (_Challenge *ChallengeSession) OneStepProveExecution(_merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _initialMessagesRead *big.Int, _initialAccs [2][32]byte, _initialState [3]*big.Int, _executionProof []byte, _bufferProof []byte, prover uint8) (*types.Transaction, error) {
	return _Challenge.Contract.OneStepProveExecution(&_Challenge.TransactOpts, _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _initialMessagesRead, _initialAccs, _initialState, _executionProof, _bufferProof, prover)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0x95979201.
//
// Solidity: function oneStepProveExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _initialMessagesRead, bytes32[2] _initialAccs, uint256[3] _initialState, bytes _executionProof, bytes _bufferProof, uint8 prover) returns()
func (_Challenge *ChallengeTransactorSession) OneStepProveExecution(_merkleNodes [][32]byte, _merkleRoute *big.Int, _challengedSegmentStart *big.Int, _challengedSegmentLength *big.Int, _oldEndHash [32]byte, _initialMessagesRead *big.Int, _initialAccs [2][32]byte, _initialState [3]*big.Int, _executionProof []byte, _bufferProof []byte, prover uint8) (*types.Transaction, error) {
	return _Challenge.Contract.OneStepProveExecution(&_Challenge.TransactOpts, _merkleNodes, _merkleRoute, _challengedSegmentStart, _challengedSegmentLength, _oldEndHash, _initialMessagesRead, _initialAccs, _initialState, _executionProof, _bufferProof, prover)
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
