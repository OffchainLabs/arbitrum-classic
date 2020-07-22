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

// MessagesERC20Message is an auto generated low-level Go binding around an user-defined struct.
type MessagesERC20Message struct {
	Token common.Address
	Dest  common.Address
	Value *big.Int
}

// MessagesERC721Message is an auto generated low-level Go binding around an user-defined struct.
type MessagesERC721Message struct {
	Token common.Address
	Dest  common.Address
	Id    *big.Int
}

// MessagesEthMessage is an auto generated low-level Go binding around an user-defined struct.
type MessagesEthMessage struct {
	Dest  common.Address
	Value *big.Int
}

// MessageTesterABI is the input ABI used to generate the binding from.
const MessageTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"addMessageToInbox\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inboxTuplePreimage\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxTupleSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageTuplePreimage\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageTupleSize\",\"type\":\"uint256\"}],\"name\":\"addMessageToVMInboxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"messageValueHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseERC20Message\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structMessages.ERC20Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseERC721Message\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"internalType\":\"structMessages.ERC721Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseEthMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structMessages.EthMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"unmarshalOutgoingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MessageTesterFuncSigs maps the 4-byte function signature to its string representation.
var MessageTesterFuncSigs = map[string]string{
	"a3b39209": "addMessageToInbox(bytes32,bytes32)",
	"f23ba5fc": "addMessageToVMInboxHash(bytes32,uint256,bytes32,uint256)",
	"fdaf43c1": "messageHash(uint8,address,uint256,uint256,uint256,bytes32)",
	"9aa86e86": "messageValueHash(uint8,uint256,uint256,address,uint256,bytes)",
	"6520427f": "parseERC20Message(bytes)",
	"fe517bd0": "parseERC721Message(bytes)",
	"ec65668c": "parseEthMessage(bytes)",
	"6b0d3519": "unmarshalOutgoingMessage(bytes,uint256)",
}

// MessageTesterBin is the compiled bytecode used for deploying new contracts.
var MessageTesterBin = "0x608060405234801561001057600080fd5b50611932806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063ec65668c1161005b578063ec65668c1461010e578063f23ba5fc1461012f578063fdaf43c114610142578063fe517bd01461008d57610088565b80636520427f1461008d5780636b0d3519146100b75780639aa86e86146100db578063a3b39209146100fb575b600080fd5b6100a061009b3660046111ea565b610155565b6040516100ae92919061171a565b60405180910390f35b6100ca6100c536600461121f565b610172565b6040516100ae959493929190611750565b6100ee6100e93660046112dc565b6101b1565b6040516100ae9190611797565b6100ee61010936600461114f565b6101d4565b61012161011c3660046111ea565b6101e9565b6040516100ae929190611735565b6100ee61013d366004611189565b6101fc565b6100ee610150366004611255565b610226565b600061015f611050565b61016883610236565b915091505b915091565b6000806000806060600080610185611070565b61018f8a8a6102b3565b80516020820151604090920151939e929d509b50995090975095505050505050565b60006101c96101c488888888888861039f565b6104c7565b979650505050505050565b60006101e083836105cc565b90505b92915050565b60006101f361108f565b610168836105ff565b600061021d6101c461020e878761065b565b610218868661065b565b61070c565b95945050505050565b60006101c987878787878761078a565b6000610240611050565b604883511015610253576000915061016d565b600c610265848263ffffffff6107c916565b6001600160a01b03168252602001610283848263ffffffff6107c916565b6001600160a01b031660208301526014016102a4848263ffffffff6107ec16565b60408301525060019150915091565b6000806102be611070565b83915060008583815181106102cf57fe5b016020015160019093019260f81c90506102e7610808565b60030160ff168160ff1614610303575060009250839150610398565b600061030f878561080d565b919650945090508461032a5750600093508492506103989050565b60ff81168352600061033c888661080d565b9197509550905085610358575060009450859350610398915050565b6001600160a01b0381166020850152610371888661088a565b6040870152909650945085610390575060009450859350610398915050565b506001945050505b9250925092565b6103a76110a6565b60408051600680825260e08201909252606091816020015b6103c76110a6565b8152602001906001900390816103bf5790505090506103e88860ff16610a6a565b816000815181106103f557fe5b602002602001018190525061040987610a6a565b8160018151811061041657fe5b602002602001018190525061042a86610a6a565b8160028151811061043757fe5b6020026020010181905250610454856001600160a01b0316610a6a565b8160038151811061046157fe5b602002602001018190525061047584610a6a565b8160048151811061048257fe5b602002602001018190525061049a8360008551610b1c565b816005815181106104a757fe5b60200260200101819052506104bb81610ca2565b98975050505050505050565b60006104d1610d7f565b60ff16826060015160ff1614156104f45781516104ed90610d84565b90506105c7565b6104fc610db4565b60ff16826060015160ff16141561051a576104ed8260200151610db9565b610522610e54565b60ff16826060015160ff16141561054457815160808301516104ed9190610e59565b61054c610808565b60ff16826060015160ff161415610585576105656110a6565b6105728360400151610e77565b905061057d816104c7565b9150506105c7565b61058d610f74565b60ff16826060015160ff1614156105a6575080516105c7565b60405162461bcd60e51b81526004016105be906117c5565b60405180910390fd5b919050565b600082826040516020016105e19291906115ae565b60405160208183030381529060405280519060200120905092915050565b600061060961108f565b60348351101561061c576000915061016d565b600c61062e848263ffffffff6107c916565b6001600160a01b0316825260140161064c848263ffffffff6107ec16565b60208301525060019150915091565b6106636110a6565b6040805160a08101825284815281516060810183526000808252602082810182905284518281528082018652939490850193908301916106b9565b6106a66110a6565b81526020019060019003908161069e5790505b509052815260408051600080825260208281019093529190920191906106f5565b6106e26110a6565b8152602001906001900390816106da5790505b508152600260208201526040019290925250919050565b6107146110a6565b6040805160028082526060828101909352816020015b6107326110a6565b81526020019060019003908161072a579050509050838160008151811061075557fe5b6020026020010181905250828160018151811061076e57fe5b602002602001018190525061078281610ca2565b949350505050565b60008686868686866040516020016107a7969594939291906115e9565b6040516020818303038152906040528051906020012090509695505050505050565b600081601401835110156107dc57600080fd5b500160200151600160601b900490565b600081602001835110156107ff57600080fd5b50016020015190565b600390565b600080600080855190508481108061082757506021858203105b8061084f5750610835610d7f565b60ff1686868151811061084457fe5b016020015160f81c14155b15610864575060009250839150829050610398565b60016021860161087c8888840163ffffffff6107ec16565b935093509350509250925092565b6000806060600061089b8686610f79565b91955093509050836108b1575060009250610398565b60208104601f82166000816108c75760006108ca565b60015b60ff16830190506060836040519080825280602002602001820160405280156108fd578160200160208202803883390190505b5090506060836040519080825280601f01601f19166020018201604052801561092d576020820181803883390190505b5090506000805b848110156109f85760006109488e8c610f79565b919d509b5090508b610967575060009a50610398975050505050505050565b811580156109755750600087115b156109c8578060005b888110156109c15781816020811061099257fe5b1a60f81b8682815181106109a257fe5b60200101906001600160f81b031916908160001a90535060010161097e565b50506109ef565b8060001b858460018b0303815181106109dd57fe5b60209081029190910101526001909201915b50600101610934565b506000610a058d8b610fdb565b909a509050610a12610808565b60ff168160ff1614610a305750600099506103989650505050505050565b60018a8585604051602001610a46929190611596565b6040516020818303038152906040529a509a509a5050505050505050509250925092565b610a726110a6565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610ac8565b610ab56110a6565b815260200190600190039081610aad5790505b50905281526040805160008082526020828101909352919092019190610b04565b610af16110a6565b815260200190600190039081610ae95790505b50815260006020820152600160409091015292915050565b610b246110a6565b60208204610b306110a6565b610b38611002565b60408051600280825260608281019093529293509091816020015b610b5b6110a6565b815260200190600190039081610b5357905050905060005b83811015610bdc57610b98610b9389602084028a0163ffffffff6107ec16565b610a6a565b82600081518110610ba557fe5b60200260200101819052508282600181518110610bbe57fe5b6020026020010181905250610bd282610e77565b9250600101610b73565b506020850615610c52576000610bfe88601f198989010163ffffffff6107ec16565b9050602086066020036008021b610c1481610a6a565b82600081518110610c2157fe5b60200260200101819052508282600181518110610c3a57fe5b6020026020010181905250610c4e82610e77565b9250505b610c5b85610a6a565b81600081518110610c6857fe5b60200260200101819052508181600181518110610c8157fe5b6020026020010181905250610c9581610e77565b93505050505b9392505050565b610caa6110a6565b610cb48251611049565b610cd05760405162461bcd60e51b81526004016105be906117a5565b600160005b8351811015610d0757838181518110610cea57fe5b602002602001015160800151820191508080600101915050610cd5565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610d61565b610d4e6110a6565b815260200190600190039081610d465790505b50905281526020810194909452600360408501526060909301525090565b600090565b600081604051602001610d9791906115d4565b604051602081830303815290604052805190602001209050919050565b600190565b6000600282604001515110610dca57fe5b604082015151610e1057610ddc610db4565b8251602080850151604051610df3949392016116a6565b6040516020818303038152906040528051906020012090506105c7565b610e18610db4565b8260000151610e3e8460400151600081518110610e3157fe5b60200260200101516104c7565b602080860151604051610d9795949392016116d2565b600290565b6000610e63610808565b83836040516020016105e19392919061166f565b610e7f6110a6565b600882511115610ea15760405162461bcd60e51b81526004016105be906117b5565b60608251604051908082528060200260200182016040528015610ece578160200160208202803883390190505b508051909150600160005b82811015610f3157610ef0868281518110610e3157fe5b848281518110610efc57fe5b602002602001018181525050858181518110610f1457fe5b602002602001015160800151820191508080600101915050610ed9565b506000835184604051602001610f48929190611653565b604051602081830303815290604052805190602001209050610f6a818361065b565b9695505050505050565b606490565b600080600080610f898686610fdb565b9093509050610f96610808565b60020160ff168160ff1614610faf575060009250610398565b610fb9868461080d565b9195509350915083610fcf575060009250610398565b50600192509250925092565b60008082600101848481518110610fee57fe5b016020015190925060f81c90509250929050565b61100a6110a6565b604080516000808252602082019092526110449161103e565b61102b6110a6565b8152602001906001900390816110235790505b50610ca2565b905090565b6008101590565b604080516060810182526000808252602082018190529181019190915290565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b6040518060a00160405280600081526020016110c0611070565b815260606020820181905260006040830181905291015290565b80356101e3816118c6565b80356101e3816118dd565b600082601f83011261110157600080fd5b813561111461110f826117fc565b6117d5565b9150808252602083016020830185838301111561113057600080fd5b61113b83828461185c565b50505092915050565b80356101e3816118e6565b6000806040838503121561116257600080fd5b600061116e85856110e5565b925050602061117f858286016110e5565b9150509250929050565b6000806000806080858703121561119f57600080fd5b60006111ab87876110e5565b94505060206111bc878288016110e5565b93505060406111cd878288016110e5565b92505060606111de878288016110e5565b91505092959194509250565b6000602082840312156111fc57600080fd5b813567ffffffffffffffff81111561121357600080fd5b610782848285016110f0565b6000806040838503121561123257600080fd5b823567ffffffffffffffff81111561124957600080fd5b61116e858286016110f0565b60008060008060008060c0878903121561126e57600080fd5b600061127a8989611144565b965050602061128b89828a016110da565b955050604061129c89828a016110e5565b94505060606112ad89828a016110e5565b93505060806112be89828a016110e5565b92505060a06112cf89828a016110e5565b9150509295509295509295565b60008060008060008060c087890312156112f557600080fd5b60006113018989611144565b965050602061131289828a016110e5565b955050604061132389828a016110e5565b945050606061133489828a016110da565b935050608061134589828a016110e5565b92505060a087013567ffffffffffffffff81111561136257600080fd5b6112cf89828a016110f0565b600061137a8383611404565b505060200190565b61138b81611837565b82525050565b61138b61139d82611837565b611894565b60006113ad8261182a565b6113b781856105c7565b93506113c283611824565b8060005b838110156113f05781516113da888261136e565b97506113e583611824565b9250506001016113c6565b509495945050505050565b61138b81611842565b61138b81611847565b61138b61141982611847565b611847565b60006114298261182a565b611433818561182e565b9350611443818560208601611868565b61144c816118b0565b9093019392505050565b60006114618261182a565b61146b81856105c7565b935061147b818560208601611868565b9290920192915050565b6000611492601a8361182e565b7f5475706c65206d75737420686176652076616c69642073697a65000000000000815260200192915050565b60006114cb60148361182e565b73092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b815260200192915050565b60006114fb60118361182e565b70496e76616c6964207479706520636f646560781b815260200192915050565b8051606083019061152c8482611382565b50602082015161153f6020850182611382565b5060408201516115526040850182611404565b50505050565b805160408301906115698482611382565b5060208201516115526020850182611404565b61138b81611856565b61138b61159182611856565b6118a5565b60006115a282856113a2565b91506107828284611456565b60006115ba828561140d565b6020820191506115ca828461140d565b5060200192915050565b60006115e0828461140d565b50602001919050565b60006115f58289611585565b6001820191506116058288611391565b601482019150611615828761140d565b602082019150611625828661140d565b602082019150611635828561140d565b602082019150611645828461140d565b506020019695505050505050565b600061165f8285611585565b60018201915061078282846113a2565b600061167b8286611585565b60018201915061168b828561140d565b60208201915061169b828461140d565b506020019392505050565b60006116b28286611585565b6001820191506116c28285611585565b60018201915061169b828461140d565b60006116de8287611585565b6001820191506116ee8286611585565b6001820191506116fe828561140d565b60208201915061170e828461140d565b50602001949350505050565b6080810161172882856113fb565b610c9b602083018461151b565b6060810161174382856113fb565b610c9b6020830184611558565b60a0810161175e82886113fb565b61176b6020830187611404565b611778604083018661157c565b6117856060830185611382565b81810360808301526101c9818461141e565b602081016101e38284611404565b602080825281016101e381611485565b602080825281016101e3816114be565b602080825281016101e3816114ee565b60405181810167ffffffffffffffff811182821017156117f457600080fd5b604052919050565b600067ffffffffffffffff82111561181357600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006101e38261184a565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b60005b8381101561188357818101518382015260200161186b565b838111156115525750506000910152565b60006101e38260006101e3826118c0565b60006101e3826118ba565b601f01601f191690565b60f81b90565b60601b90565b6118cf81611837565b81146118da57600080fd5b50565b6118cf81611847565b6118cf8161185656fea365627a7a72315820ff77dd5c878e018b779b979f0472ae02a24f04066cd5c3728d167dab0ab5799d6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployMessageTester deploys a new Ethereum contract, binding an instance of MessageTester to it.
func DeployMessageTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageTester, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessageTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// MessageTester is an auto generated Go binding around an Ethereum contract.
type MessageTester struct {
	MessageTesterCaller     // Read-only binding to the contract
	MessageTesterTransactor // Write-only binding to the contract
	MessageTesterFilterer   // Log filterer for contract events
}

// MessageTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageTesterSession struct {
	Contract     *MessageTester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageTesterCallerSession struct {
	Contract *MessageTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MessageTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTesterTransactorSession struct {
	Contract     *MessageTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MessageTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageTesterRaw struct {
	Contract *MessageTester // Generic contract binding to access the raw methods on
}

// MessageTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageTesterCallerRaw struct {
	Contract *MessageTesterCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTesterTransactorRaw struct {
	Contract *MessageTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageTester creates a new instance of MessageTester, bound to a specific deployed contract.
func NewMessageTester(address common.Address, backend bind.ContractBackend) (*MessageTester, error) {
	contract, err := bindMessageTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// NewMessageTesterCaller creates a new read-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterCaller(address common.Address, caller bind.ContractCaller) (*MessageTesterCaller, error) {
	contract, err := bindMessageTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterCaller{contract: contract}, nil
}

// NewMessageTesterTransactor creates a new write-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTesterTransactor, error) {
	contract, err := bindMessageTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterTransactor{contract: contract}, nil
}

// NewMessageTesterFilterer creates a new log filterer instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageTesterFilterer, error) {
	contract, err := bindMessageTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageTesterFilterer{contract: contract}, nil
}

// bindMessageTester binds a generic wrapper to an already deployed contract.
func bindMessageTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.MessageTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transact(opts, method, params...)
}

// AddMessageToInbox is a free data retrieval call binding the contract method 0xa3b39209.
//
// Solidity: function addMessageToInbox(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) AddMessageToInbox(opts *bind.CallOpts, inbox [32]byte, message [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "addMessageToInbox", inbox, message)
	return *ret0, err
}

// AddMessageToInbox is a free data retrieval call binding the contract method 0xa3b39209.
//
// Solidity: function addMessageToInbox(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) AddMessageToInbox(inbox [32]byte, message [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.AddMessageToInbox(&_MessageTester.CallOpts, inbox, message)
}

// AddMessageToInbox is a free data retrieval call binding the contract method 0xa3b39209.
//
// Solidity: function addMessageToInbox(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) AddMessageToInbox(inbox [32]byte, message [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.AddMessageToInbox(&_MessageTester.CallOpts, inbox, message)
}

// AddMessageToVMInboxHash is a free data retrieval call binding the contract method 0xf23ba5fc.
//
// Solidity: function addMessageToVMInboxHash(bytes32 inboxTuplePreimage, uint256 inboxTupleSize, bytes32 messageTuplePreimage, uint256 messageTupleSize) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) AddMessageToVMInboxHash(opts *bind.CallOpts, inboxTuplePreimage [32]byte, inboxTupleSize *big.Int, messageTuplePreimage [32]byte, messageTupleSize *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "addMessageToVMInboxHash", inboxTuplePreimage, inboxTupleSize, messageTuplePreimage, messageTupleSize)
	return *ret0, err
}

// AddMessageToVMInboxHash is a free data retrieval call binding the contract method 0xf23ba5fc.
//
// Solidity: function addMessageToVMInboxHash(bytes32 inboxTuplePreimage, uint256 inboxTupleSize, bytes32 messageTuplePreimage, uint256 messageTupleSize) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) AddMessageToVMInboxHash(inboxTuplePreimage [32]byte, inboxTupleSize *big.Int, messageTuplePreimage [32]byte, messageTupleSize *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.AddMessageToVMInboxHash(&_MessageTester.CallOpts, inboxTuplePreimage, inboxTupleSize, messageTuplePreimage, messageTupleSize)
}

// AddMessageToVMInboxHash is a free data retrieval call binding the contract method 0xf23ba5fc.
//
// Solidity: function addMessageToVMInboxHash(bytes32 inboxTuplePreimage, uint256 inboxTupleSize, bytes32 messageTuplePreimage, uint256 messageTupleSize) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) AddMessageToVMInboxHash(inboxTuplePreimage [32]byte, inboxTupleSize *big.Int, messageTuplePreimage [32]byte, messageTupleSize *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.AddMessageToVMInboxHash(&_MessageTester.CallOpts, inboxTuplePreimage, inboxTupleSize, messageTuplePreimage, messageTupleSize)
}

// MessageHash is a free data retrieval call binding the contract method 0xfdaf43c1.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint256 blockNumber, uint256 timestamp, uint256 inboxSeqNum, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) MessageHash(opts *bind.CallOpts, messageType uint8, sender common.Address, blockNumber *big.Int, timestamp *big.Int, inboxSeqNum *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "messageHash", messageType, sender, blockNumber, timestamp, inboxSeqNum, messageDataHash)
	return *ret0, err
}

// MessageHash is a free data retrieval call binding the contract method 0xfdaf43c1.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint256 blockNumber, uint256 timestamp, uint256 inboxSeqNum, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) MessageHash(messageType uint8, sender common.Address, blockNumber *big.Int, timestamp *big.Int, inboxSeqNum *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageHash(&_MessageTester.CallOpts, messageType, sender, blockNumber, timestamp, inboxSeqNum, messageDataHash)
}

// MessageHash is a free data retrieval call binding the contract method 0xfdaf43c1.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint256 blockNumber, uint256 timestamp, uint256 inboxSeqNum, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) MessageHash(messageType uint8, sender common.Address, blockNumber *big.Int, timestamp *big.Int, inboxSeqNum *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageHash(&_MessageTester.CallOpts, messageType, sender, blockNumber, timestamp, inboxSeqNum, messageDataHash)
}

// MessageValueHash is a free data retrieval call binding the contract method 0x9aa86e86.
//
// Solidity: function messageValueHash(uint8 messageType, uint256 blockNumber, uint256 timestamp, address sender, uint256 inboxSeqNum, bytes messageData) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) MessageValueHash(opts *bind.CallOpts, messageType uint8, blockNumber *big.Int, timestamp *big.Int, sender common.Address, inboxSeqNum *big.Int, messageData []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "messageValueHash", messageType, blockNumber, timestamp, sender, inboxSeqNum, messageData)
	return *ret0, err
}

// MessageValueHash is a free data retrieval call binding the contract method 0x9aa86e86.
//
// Solidity: function messageValueHash(uint8 messageType, uint256 blockNumber, uint256 timestamp, address sender, uint256 inboxSeqNum, bytes messageData) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) MessageValueHash(messageType uint8, blockNumber *big.Int, timestamp *big.Int, sender common.Address, inboxSeqNum *big.Int, messageData []byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageValueHash(&_MessageTester.CallOpts, messageType, blockNumber, timestamp, sender, inboxSeqNum, messageData)
}

// MessageValueHash is a free data retrieval call binding the contract method 0x9aa86e86.
//
// Solidity: function messageValueHash(uint8 messageType, uint256 blockNumber, uint256 timestamp, address sender, uint256 inboxSeqNum, bytes messageData) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) MessageValueHash(messageType uint8, blockNumber *big.Int, timestamp *big.Int, sender common.Address, inboxSeqNum *big.Int, messageData []byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageValueHash(&_MessageTester.CallOpts, messageType, blockNumber, timestamp, sender, inboxSeqNum, messageData)
}

// ParseERC20Message is a free data retrieval call binding the contract method 0x6520427f.
//
// Solidity: function parseERC20Message(bytes data) pure returns(bool valid, (address,address,uint256) message)
func (_MessageTester *MessageTesterCaller) ParseERC20Message(opts *bind.CallOpts, data []byte) (struct {
	Valid   bool
	Message MessagesERC20Message
}, error) {
	ret := new(struct {
		Valid   bool
		Message MessagesERC20Message
	})
	out := ret
	err := _MessageTester.contract.Call(opts, out, "parseERC20Message", data)
	return *ret, err
}

// ParseERC20Message is a free data retrieval call binding the contract method 0x6520427f.
//
// Solidity: function parseERC20Message(bytes data) pure returns(bool valid, (address,address,uint256) message)
func (_MessageTester *MessageTesterSession) ParseERC20Message(data []byte) (struct {
	Valid   bool
	Message MessagesERC20Message
}, error) {
	return _MessageTester.Contract.ParseERC20Message(&_MessageTester.CallOpts, data)
}

// ParseERC20Message is a free data retrieval call binding the contract method 0x6520427f.
//
// Solidity: function parseERC20Message(bytes data) pure returns(bool valid, (address,address,uint256) message)
func (_MessageTester *MessageTesterCallerSession) ParseERC20Message(data []byte) (struct {
	Valid   bool
	Message MessagesERC20Message
}, error) {
	return _MessageTester.Contract.ParseERC20Message(&_MessageTester.CallOpts, data)
}

// ParseERC721Message is a free data retrieval call binding the contract method 0xfe517bd0.
//
// Solidity: function parseERC721Message(bytes data) pure returns(bool valid, (address,address,uint256) message)
func (_MessageTester *MessageTesterCaller) ParseERC721Message(opts *bind.CallOpts, data []byte) (struct {
	Valid   bool
	Message MessagesERC721Message
}, error) {
	ret := new(struct {
		Valid   bool
		Message MessagesERC721Message
	})
	out := ret
	err := _MessageTester.contract.Call(opts, out, "parseERC721Message", data)
	return *ret, err
}

// ParseERC721Message is a free data retrieval call binding the contract method 0xfe517bd0.
//
// Solidity: function parseERC721Message(bytes data) pure returns(bool valid, (address,address,uint256) message)
func (_MessageTester *MessageTesterSession) ParseERC721Message(data []byte) (struct {
	Valid   bool
	Message MessagesERC721Message
}, error) {
	return _MessageTester.Contract.ParseERC721Message(&_MessageTester.CallOpts, data)
}

// ParseERC721Message is a free data retrieval call binding the contract method 0xfe517bd0.
//
// Solidity: function parseERC721Message(bytes data) pure returns(bool valid, (address,address,uint256) message)
func (_MessageTester *MessageTesterCallerSession) ParseERC721Message(data []byte) (struct {
	Valid   bool
	Message MessagesERC721Message
}, error) {
	return _MessageTester.Contract.ParseERC721Message(&_MessageTester.CallOpts, data)
}

// ParseEthMessage is a free data retrieval call binding the contract method 0xec65668c.
//
// Solidity: function parseEthMessage(bytes data) pure returns(bool valid, (address,uint256) message)
func (_MessageTester *MessageTesterCaller) ParseEthMessage(opts *bind.CallOpts, data []byte) (struct {
	Valid   bool
	Message MessagesEthMessage
}, error) {
	ret := new(struct {
		Valid   bool
		Message MessagesEthMessage
	})
	out := ret
	err := _MessageTester.contract.Call(opts, out, "parseEthMessage", data)
	return *ret, err
}

// ParseEthMessage is a free data retrieval call binding the contract method 0xec65668c.
//
// Solidity: function parseEthMessage(bytes data) pure returns(bool valid, (address,uint256) message)
func (_MessageTester *MessageTesterSession) ParseEthMessage(data []byte) (struct {
	Valid   bool
	Message MessagesEthMessage
}, error) {
	return _MessageTester.Contract.ParseEthMessage(&_MessageTester.CallOpts, data)
}

// ParseEthMessage is a free data retrieval call binding the contract method 0xec65668c.
//
// Solidity: function parseEthMessage(bytes data) pure returns(bool valid, (address,uint256) message)
func (_MessageTester *MessageTesterCallerSession) ParseEthMessage(data []byte) (struct {
	Valid   bool
	Message MessagesEthMessage
}, error) {
	return _MessageTester.Contract.ParseEthMessage(&_MessageTester.CallOpts, data)
}

// UnmarshalOutgoingMessage is a free data retrieval call binding the contract method 0x6b0d3519.
//
// Solidity: function unmarshalOutgoingMessage(bytes data, uint256 startOffset) pure returns(bool, uint256, uint8, address, bytes)
func (_MessageTester *MessageTesterCaller) UnmarshalOutgoingMessage(opts *bind.CallOpts, data []byte, startOffset *big.Int) (bool, *big.Int, uint8, common.Address, []byte, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(uint8)
		ret3 = new(common.Address)
		ret4 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _MessageTester.contract.Call(opts, out, "unmarshalOutgoingMessage", data, startOffset)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// UnmarshalOutgoingMessage is a free data retrieval call binding the contract method 0x6b0d3519.
//
// Solidity: function unmarshalOutgoingMessage(bytes data, uint256 startOffset) pure returns(bool, uint256, uint8, address, bytes)
func (_MessageTester *MessageTesterSession) UnmarshalOutgoingMessage(data []byte, startOffset *big.Int) (bool, *big.Int, uint8, common.Address, []byte, error) {
	return _MessageTester.Contract.UnmarshalOutgoingMessage(&_MessageTester.CallOpts, data, startOffset)
}

// UnmarshalOutgoingMessage is a free data retrieval call binding the contract method 0x6b0d3519.
//
// Solidity: function unmarshalOutgoingMessage(bytes data, uint256 startOffset) pure returns(bool, uint256, uint8, address, bytes)
func (_MessageTester *MessageTesterCallerSession) UnmarshalOutgoingMessage(data []byte, startOffset *big.Int) (bool, *big.Int, uint8, common.Address, []byte, error) {
	return _MessageTester.Contract.UnmarshalOutgoingMessage(&_MessageTester.CallOpts, data, startOffset)
}
