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
const MessageTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"addMessageToInbox\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"messageValueHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseERC20Message\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structMessages.ERC20Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseERC721Message\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"internalType\":\"structMessages.ERC721Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseEthMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structMessages.EthMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"unmarshalOutgoingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MessageTesterFuncSigs maps the 4-byte function signature to its string representation.
var MessageTesterFuncSigs = map[string]string{
	"a3b39209": "addMessageToInbox(bytes32,bytes32)",
	"fdaf43c1": "messageHash(uint8,address,uint256,uint256,uint256,bytes32)",
	"9aa86e86": "messageValueHash(uint8,uint256,uint256,address,uint256,bytes)",
	"6520427f": "parseERC20Message(bytes)",
	"fe517bd0": "parseERC721Message(bytes)",
	"ec65668c": "parseEthMessage(bytes)",
	"6b0d3519": "unmarshalOutgoingMessage(bytes,uint256)",
}

// MessageTesterBin is the compiled bytecode used for deploying new contracts.
var MessageTesterBin = "0x608060405234801561001057600080fd5b50611877806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063a3b392091161005b578063a3b39209146100f0578063ec65668c14610103578063fdaf43c114610124578063fe517bd0146100825761007d565b80636520427f146100825780636b0d3519146100ac5780639aa86e86146100d0575b600080fd5b6100956100903660046110e9565b610137565b6040516100a392919061164f565b60405180910390f35b6100bf6100ba366004611126565b610154565b6040516100a3959493929190611685565b6100e36100de3660046111e3565b610193565b6040516100a391906116cc565b6100e36100fe3660046110af565b6101b6565b6101166101113660046110e9565b6101cb565b6040516100a392919061166a565b6100e361013236600461115c565b6101de565b6000610141610fb0565b61014a836101ee565b915091505b915091565b6000806000806060600080610167610fd0565b6101718a8a61026b565b80516020820151604090920151939e929d509b50995090975095505050505050565b60006101ab6101a6888888888888610357565b61047f565b979650505050505050565b60006101c28383610584565b90505b92915050565b60006101d5610fef565b61014a836105b7565b60006101ab878787878787610613565b60006101f8610fb0565b60488351101561020b576000915061014f565b600c61021d848263ffffffff61065216565b6001600160a01b0316825260200161023b848263ffffffff61065216565b6001600160a01b0316602083015260140161025c848263ffffffff61068816565b60408301525060019150915091565b600080610276610fd0565b839150600085838151811061028757fe5b016020015160019093019260f81c905061029f6106b7565b60030160ff168160ff16146102bb575060009250839150610350565b60006102c787856106bc565b91965094509050846102e25750600093508492506103509050565b60ff8116835260006102f488866106bc565b9197509550905085610310575060009450859350610350915050565b6001600160a01b03811660208501526103298886610739565b6040870152909650945085610348575060009450859350610350915050565b506001945050505b9250925092565b61035f611006565b60408051600680825260e08201909252606091816020015b61037f611006565b8152602001906001900390816103775790505090506103a08860ff16610919565b816000815181106103ad57fe5b60200260200101819052506103c187610919565b816001815181106103ce57fe5b60200260200101819052506103e286610919565b816002815181106103ef57fe5b602002602001018190525061040c856001600160a01b0316610919565b8160038151811061041957fe5b602002602001018190525061042d84610919565b8160048151811061043a57fe5b602002602001018190525061045283600085516109cb565b8160058151811061045f57fe5b602002602001018190525061047381610b51565b98975050505050505050565b6000610489610c2e565b60ff16826060015160ff1614156104ac5781516104a590610c33565b905061057f565b6104b4610c63565b60ff16826060015160ff1614156104d2576104a58260200151610c68565b6104da610d03565b60ff16826060015160ff1614156104fc57815160808301516104a59190610d08565b6105046106b7565b60ff16826060015160ff16141561053d5761051d611006565b61052a8360400151610d26565b90506105358161047f565b91505061057f565b610545610e23565b60ff16826060015160ff16141561055e5750805161057f565b60405162461bcd60e51b81526004016105769061170a565b60405180910390fd5b919050565b600082826040516020016105999291906114e3565b60405160208183030381529060405280519060200120905092915050565b60006105c1610fef565b6034835110156105d4576000915061014f565b600c6105e6848263ffffffff61065216565b6001600160a01b03168252601401610604848263ffffffff61068816565b60208301525060019150915091565b60008686868686866040516020016106309695949392919061151e565b6040516020818303038152906040528051906020012090509695505050505050565b600081601401835110156106785760405162461bcd60e51b8152600401610576906116ea565b500160200151600160601b900490565b600081602001835110156106ae5760405162461bcd60e51b8152600401610576906116ea565b50016020015190565b600390565b60008060008085519050848110806106d657506021858203105b806106fe57506106e4610c2e565b60ff168686815181106106f357fe5b016020015160f81c14155b15610713575060009250839150829050610350565b60016021860161072b8888840163ffffffff61068816565b935093509350509250925092565b6000806060600061074a8686610e28565b9195509350905083610760575060009250610350565b60208104601f8216600081610776576000610779565b60015b60ff16830190506060836040519080825280602002602001820160405280156107ac578160200160208202803883390190505b5090506060836040519080825280601f01601f1916602001820160405280156107dc576020820181803883390190505b5090506000805b848110156108a75760006107f78e8c610e28565b919d509b5090508b610816575060009a50610350975050505050505050565b811580156108245750600087115b15610877578060005b888110156108705781816020811061084157fe5b1a60f81b86828151811061085157fe5b60200101906001600160f81b031916908160001a90535060010161082d565b505061089e565b8060001b858460018b03038151811061088c57fe5b60209081029190910101526001909201915b506001016107e3565b5060006108b48d8b610e8a565b909a5090506108c16106b7565b60ff168160ff16146108df5750600099506103509650505050505050565b60018a85856040516020016108f59291906114cb565b6040516020818303038152906040529a509a509a5050505050505050509250925092565b610921611006565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610977565b610964611006565b81526020019060019003908161095c5790505b509052815260408051600080825260208281019093529190920191906109b3565b6109a0611006565b8152602001906001900390816109985790505b50815260006020820152600160409091015292915050565b6109d3611006565b602082046109df611006565b6109e7610eb1565b60408051600280825260608281019093529293509091816020015b610a0a611006565b815260200190600190039081610a0257905050905060005b83811015610a8b57610a47610a4289602084028a0163ffffffff61068816565b610919565b82600081518110610a5457fe5b60200260200101819052508282600181518110610a6d57fe5b6020026020010181905250610a8182610d26565b9250600101610a22565b506020850615610b01576000610aad88601f198989010163ffffffff61068816565b9050602086066020036008021b610ac381610919565b82600081518110610ad057fe5b60200260200101819052508282600181518110610ae957fe5b6020026020010181905250610afd82610d26565b9250505b610b0a85610919565b81600081518110610b1757fe5b60200260200101819052508181600181518110610b3057fe5b6020026020010181905250610b4481610d26565b93505050505b9392505050565b610b59611006565b610b638251610ef8565b610b7f5760405162461bcd60e51b8152600401610576906116da565b600160005b8351811015610bb657838181518110610b9957fe5b602002602001015160800151820191508080600101915050610b84565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610c10565b610bfd611006565b815260200190600190039081610bf55790505b50905281526020810194909452600360408501526060909301525090565b600090565b600081604051602001610c469190611509565b604051602081830303815290604052805190602001209050919050565b600190565b6000600282604001515110610c7957fe5b604082015151610cbf57610c8b610c63565b8251602080850151604051610ca2949392016115db565b60405160208183030381529060405280519060200120905061057f565b610cc7610c63565b8260000151610ced8460400151600081518110610ce057fe5b602002602001015161047f565b602080860151604051610c469594939201611607565b600290565b6000610d126106b7565b8383604051602001610599939291906115a4565b610d2e611006565b600882511115610d505760405162461bcd60e51b8152600401610576906116fa565b60608251604051908082528060200260200182016040528015610d7d578160200160208202803883390190505b508051909150600160005b82811015610de057610d9f868281518110610ce057fe5b848281518110610dab57fe5b602002602001018181525050858181518110610dc357fe5b602002602001015160800151820191508080600101915050610d88565b506000835184604051602001610df7929190611588565b604051602081830303815290604052805190602001209050610e198183610eff565b9695505050505050565b606490565b600080600080610e388686610e8a565b9093509050610e456106b7565b60020160ff168160ff1614610e5e575060009250610350565b610e6886846106bc565b9195509350915083610e7e575060009250610350565b50600192509250925092565b60008082600101848481518110610e9d57fe5b016020015190925060f81c90509250929050565b610eb9611006565b60408051600080825260208201909252610ef391610eed565b610eda611006565b815260200190600190039081610ed25790505b50610b51565b905090565b6008101590565b610f07611006565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191610f5d565b610f4a611006565b815260200190600190039081610f425790505b50905281526040805160008082526020828101909352919092019190610f99565b610f86611006565b815260200190600190039081610f7e5790505b508152600260208201526040019290925250919050565b604080516060810182526000808252602082018190529181019190915290565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b6040518060a0016040528060008152602001611020610fd0565b815260606020820181905260006040830181905291015290565b80356101c58161180b565b80356101c581611822565b600082601f83011261106157600080fd5b813561107461106f82611741565b61171a565b9150808252602083016020830185838301111561109057600080fd5b61109b8382846117a1565b50505092915050565b80356101c58161182b565b600080604083850312156110c257600080fd5b60006110ce8585611045565b92505060206110df85828601611045565b9150509250929050565b6000602082840312156110fb57600080fd5b813567ffffffffffffffff81111561111257600080fd5b61111e84828501611050565b949350505050565b6000806040838503121561113957600080fd5b823567ffffffffffffffff81111561115057600080fd5b6110ce85828601611050565b60008060008060008060c0878903121561117557600080fd5b600061118189896110a4565b965050602061119289828a0161103a565b95505060406111a389828a01611045565b94505060606111b489828a01611045565b93505060806111c589828a01611045565b92505060a06111d689828a01611045565b9150509295509295509295565b60008060008060008060c087890312156111fc57600080fd5b600061120889896110a4565b965050602061121989828a01611045565b955050604061122a89828a01611045565b945050606061123b89828a0161103a565b935050608061124c89828a01611045565b92505060a087013567ffffffffffffffff81111561126957600080fd5b6111d689828a01611050565b6000611281838361130b565b505060200190565b6112928161177c565b82525050565b6112926112a48261177c565b6117d9565b60006112b48261176f565b6112be818561057f565b93506112c983611769565b8060005b838110156112f75781516112e18882611275565b97506112ec83611769565b9250506001016112cd565b509495945050505050565b61129281611787565b6112928161178c565b6112926113208261178c565b61178c565b60006113308261176f565b61133a8185611773565b935061134a8185602086016117ad565b611353816117f5565b9093019392505050565b60006113688261176f565b611372818561057f565b93506113828185602086016117ad565b9290920192915050565b6000611399601a83611773565b7f5475706c65206d75737420686176652076616c69642073697a65000000000000815260200192915050565b60006113d2601283611773565b7152656164206f7574206f6620626f756e647360701b815260200192915050565b6000611400601483611773565b73092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b815260200192915050565b6000611430601183611773565b70496e76616c6964207479706520636f646560781b815260200192915050565b805160608301906114618482611289565b5060208201516114746020850182611289565b506040820151611487604085018261130b565b50505050565b8051604083019061149e8482611289565b506020820151611487602085018261130b565b6112928161179b565b6112926114c68261179b565b6117ea565b60006114d782856112a9565b915061111e828461135d565b60006114ef8285611314565b6020820191506114ff8284611314565b5060200192915050565b60006115158284611314565b50602001919050565b600061152a82896114ba565b60018201915061153a8288611298565b60148201915061154a8287611314565b60208201915061155a8286611314565b60208201915061156a8285611314565b60208201915061157a8284611314565b506020019695505050505050565b600061159482856114ba565b60018201915061111e82846112a9565b60006115b082866114ba565b6001820191506115c08285611314565b6020820191506115d08284611314565b506020019392505050565b60006115e782866114ba565b6001820191506115f782856114ba565b6001820191506115d08284611314565b600061161382876114ba565b60018201915061162382866114ba565b6001820191506116338285611314565b6020820191506116438284611314565b50602001949350505050565b6080810161165d8285611302565b610b4a6020830184611450565b606081016116788285611302565b610b4a602083018461148d565b60a081016116938288611302565b6116a0602083018761130b565b6116ad60408301866114b1565b6116ba6060830185611289565b81810360808301526101ab8184611325565b602081016101c5828461130b565b602080825281016101c58161138c565b602080825281016101c5816113c5565b602080825281016101c5816113f3565b602080825281016101c581611423565b60405181810167ffffffffffffffff8111828210171561173957600080fd5b604052919050565b600067ffffffffffffffff82111561175857600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006101c58261178f565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b60005b838110156117c85781810151838201526020016117b0565b838111156114875750506000910152565b60006101c58260006101c582611805565b60006101c5826117ff565b601f01601f191690565b60f81b90565b60601b90565b6118148161177c565b811461181f57600080fd5b50565b6118148161178c565b6118148161179b56fea365627a7a7231582084d77d39683577721bf2513052bd213b5a6f17fcf56ae63b6d24cecd0094774d6c6578706572696d656e74616cf564736f6c63430005110040"

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
