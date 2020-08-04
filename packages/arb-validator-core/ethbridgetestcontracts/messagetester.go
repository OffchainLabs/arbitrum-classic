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
var MessageTesterBin = "0x608060405234801561001057600080fd5b50611813806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063a3b392091161005b578063a3b39209146100f0578063ec65668c14610103578063fdaf43c114610124578063fe517bd0146100825761007d565b80636520427f146100825780636b0d3519146100ac5780639aa86e86146100d0575b600080fd5b6100956100903660046110c3565b610137565b6040516100a39291906115fb565b60405180910390f35b6100bf6100ba366004611100565b610154565b6040516100a3959493929190611631565b6100e36100de3660046111bd565b610193565b6040516100a39190611678565b6100e36100fe366004611089565b6101b6565b6101166101113660046110c3565b6101cb565b6040516100a3929190611616565b6100e3610132366004611136565b6101de565b6000610141610f8a565b61014a836101ee565b915091505b915091565b6000806000806060600080610167610faa565b6101718a8a61026b565b80516020820151604090920151939e929d509b50995090975095505050505050565b60006101ab6101a6888888888888610357565b61047f565b979650505050505050565b60006101c28383610584565b90505b92915050565b60006101d5610fc9565b61014a836105b7565b60006101ab878787878787610613565b60006101f8610f8a565b60488351101561020b576000915061014f565b600c61021d848263ffffffff61065216565b6001600160a01b0316825260200161023b848263ffffffff61065216565b6001600160a01b0316602083015260140161025c848263ffffffff61067516565b60408301525060019150915091565b600080610276610faa565b839150600085838151811061028757fe5b016020015160019093019260f81c905061029f610691565b60030160ff168160ff16146102bb575060009250839150610350565b60006102c78785610696565b91965094509050846102e25750600093508492506103509050565b60ff8116835260006102f48886610696565b9197509550905085610310575060009450859350610350915050565b6001600160a01b03811660208501526103298886610713565b6040870152909650945085610348575060009450859350610350915050565b506001945050505b9250925092565b61035f610fe0565b60408051600680825260e08201909252606091816020015b61037f610fe0565b8152602001906001900390816103775790505090506103a08860ff166108f3565b816000815181106103ad57fe5b60200260200101819052506103c1876108f3565b816001815181106103ce57fe5b60200260200101819052506103e2866108f3565b816002815181106103ef57fe5b602002602001018190525061040c856001600160a01b03166108f3565b8160038151811061041957fe5b602002602001018190525061042d846108f3565b8160048151811061043a57fe5b602002602001018190525061045283600085516109a5565b8160058151811061045f57fe5b602002602001018190525061047381610b2b565b98975050505050505050565b6000610489610c08565b60ff16826060015160ff1614156104ac5781516104a590610c0d565b905061057f565b6104b4610c3d565b60ff16826060015160ff1614156104d2576104a58260200151610c42565b6104da610cdd565b60ff16826060015160ff1614156104fc57815160808301516104a59190610ce2565b610504610691565b60ff16826060015160ff16141561053d5761051d610fe0565b61052a8360400151610d00565b90506105358161047f565b91505061057f565b610545610dfd565b60ff16826060015160ff16141561055e5750805161057f565b60405162461bcd60e51b8152600401610576906116a6565b60405180910390fd5b919050565b6000828260405160200161059992919061148f565b60405160208183030381529060405280519060200120905092915050565b60006105c1610fc9565b6034835110156105d4576000915061014f565b600c6105e6848263ffffffff61065216565b6001600160a01b03168252601401610604848263ffffffff61067516565b60208301525060019150915091565b6000868686868686604051602001610630969594939291906114ca565b6040516020818303038152906040528051906020012090509695505050505050565b6000816014018351101561066557600080fd5b500160200151600160601b900490565b6000816020018351101561068857600080fd5b50016020015190565b600390565b60008060008085519050848110806106b057506021858203105b806106d857506106be610c08565b60ff168686815181106106cd57fe5b016020015160f81c14155b156106ed575060009250839150829050610350565b6001602186016107058888840163ffffffff61067516565b935093509350509250925092565b600080606060006107248686610e02565b919550935090508361073a575060009250610350565b60208104601f8216600081610750576000610753565b60015b60ff1683019050606083604051908082528060200260200182016040528015610786578160200160208202803883390190505b5090506060836040519080825280601f01601f1916602001820160405280156107b6576020820181803883390190505b5090506000805b848110156108815760006107d18e8c610e02565b919d509b5090508b6107f0575060009a50610350975050505050505050565b811580156107fe5750600087115b15610851578060005b8881101561084a5781816020811061081b57fe5b1a60f81b86828151811061082b57fe5b60200101906001600160f81b031916908160001a905350600101610807565b5050610878565b8060001b858460018b03038151811061086657fe5b60209081029190910101526001909201915b506001016107bd565b50600061088e8d8b610e64565b909a50905061089b610691565b60ff168160ff16146108b95750600099506103509650505050505050565b60018a85856040516020016108cf929190611477565b6040516020818303038152906040529a509a509a5050505050505050509250925092565b6108fb610fe0565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610951565b61093e610fe0565b8152602001906001900390816109365790505b5090528152604080516000808252602082810190935291909201919061098d565b61097a610fe0565b8152602001906001900390816109725790505b50815260006020820152600160409091015292915050565b6109ad610fe0565b602082046109b9610fe0565b6109c1610e8b565b60408051600280825260608281019093529293509091816020015b6109e4610fe0565b8152602001906001900390816109dc57905050905060005b83811015610a6557610a21610a1c89602084028a0163ffffffff61067516565b6108f3565b82600081518110610a2e57fe5b60200260200101819052508282600181518110610a4757fe5b6020026020010181905250610a5b82610d00565b92506001016109fc565b506020850615610adb576000610a8788601f198989010163ffffffff61067516565b9050602086066020036008021b610a9d816108f3565b82600081518110610aaa57fe5b60200260200101819052508282600181518110610ac357fe5b6020026020010181905250610ad782610d00565b9250505b610ae4856108f3565b81600081518110610af157fe5b60200260200101819052508181600181518110610b0a57fe5b6020026020010181905250610b1e81610d00565b93505050505b9392505050565b610b33610fe0565b610b3d8251610ed2565b610b595760405162461bcd60e51b815260040161057690611686565b600160005b8351811015610b9057838181518110610b7357fe5b602002602001015160800151820191508080600101915050610b5e565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610bea565b610bd7610fe0565b815260200190600190039081610bcf5790505b50905281526020810194909452600360408501526060909301525090565b600090565b600081604051602001610c2091906114b5565b604051602081830303815290604052805190602001209050919050565b600190565b6000600282604001515110610c5357fe5b604082015151610c9957610c65610c3d565b8251602080850151604051610c7c94939201611587565b60405160208183030381529060405280519060200120905061057f565b610ca1610c3d565b8260000151610cc78460400151600081518110610cba57fe5b602002602001015161047f565b602080860151604051610c2095949392016115b3565b600290565b6000610cec610691565b838360405160200161059993929190611550565b610d08610fe0565b600882511115610d2a5760405162461bcd60e51b815260040161057690611696565b60608251604051908082528060200260200182016040528015610d57578160200160208202803883390190505b508051909150600160005b82811015610dba57610d79868281518110610cba57fe5b848281518110610d8557fe5b602002602001018181525050858181518110610d9d57fe5b602002602001015160800151820191508080600101915050610d62565b506000835184604051602001610dd1929190611534565b604051602081830303815290604052805190602001209050610df38183610ed9565b9695505050505050565b606490565b600080600080610e128686610e64565b9093509050610e1f610691565b60020160ff168160ff1614610e38575060009250610350565b610e428684610696565b9195509350915083610e58575060009250610350565b50600192509250925092565b60008082600101848481518110610e7757fe5b016020015190925060f81c90509250929050565b610e93610fe0565b60408051600080825260208201909252610ecd91610ec7565b610eb4610fe0565b815260200190600190039081610eac5790505b50610b2b565b905090565b6008101590565b610ee1610fe0565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191610f37565b610f24610fe0565b815260200190600190039081610f1c5790505b50905281526040805160008082526020828101909352919092019190610f73565b610f60610fe0565b815260200190600190039081610f585790505b508152600260208201526040019290925250919050565b604080516060810182526000808252602082018190529181019190915290565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b6040518060a0016040528060008152602001610ffa610faa565b815260606020820181905260006040830181905291015290565b80356101c5816117a7565b80356101c5816117be565b600082601f83011261103b57600080fd5b813561104e611049826116dd565b6116b6565b9150808252602083016020830185838301111561106a57600080fd5b61107583828461173d565b50505092915050565b80356101c5816117c7565b6000806040838503121561109c57600080fd5b60006110a8858561101f565b92505060206110b98582860161101f565b9150509250929050565b6000602082840312156110d557600080fd5b813567ffffffffffffffff8111156110ec57600080fd5b6110f88482850161102a565b949350505050565b6000806040838503121561111357600080fd5b823567ffffffffffffffff81111561112a57600080fd5b6110a88582860161102a565b60008060008060008060c0878903121561114f57600080fd5b600061115b898961107e565b965050602061116c89828a01611014565b955050604061117d89828a0161101f565b945050606061118e89828a0161101f565b935050608061119f89828a0161101f565b92505060a06111b089828a0161101f565b9150509295509295509295565b60008060008060008060c087890312156111d657600080fd5b60006111e2898961107e565b96505060206111f389828a0161101f565b955050604061120489828a0161101f565b945050606061121589828a01611014565b935050608061122689828a0161101f565b92505060a087013567ffffffffffffffff81111561124357600080fd5b6111b089828a0161102a565b600061125b83836112e5565b505060200190565b61126c81611718565b82525050565b61126c61127e82611718565b611775565b600061128e8261170b565b611298818561057f565b93506112a383611705565b8060005b838110156112d15781516112bb888261124f565b97506112c683611705565b9250506001016112a7565b509495945050505050565b61126c81611723565b61126c81611728565b61126c6112fa82611728565b611728565b600061130a8261170b565b611314818561170f565b9350611324818560208601611749565b61132d81611791565b9093019392505050565b60006113428261170b565b61134c818561057f565b935061135c818560208601611749565b9290920192915050565b6000611373601a8361170f565b7f5475706c65206d75737420686176652076616c69642073697a65000000000000815260200192915050565b60006113ac60148361170f565b73092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b815260200192915050565b60006113dc60118361170f565b70496e76616c6964207479706520636f646560781b815260200192915050565b8051606083019061140d8482611263565b5060208201516114206020850182611263565b50604082015161143360408501826112e5565b50505050565b8051604083019061144a8482611263565b50602082015161143360208501826112e5565b61126c81611737565b61126c61147282611737565b611786565b60006114838285611283565b91506110f88284611337565b600061149b82856112ee565b6020820191506114ab82846112ee565b5060200192915050565b60006114c182846112ee565b50602001919050565b60006114d68289611466565b6001820191506114e68288611272565b6014820191506114f682876112ee565b60208201915061150682866112ee565b60208201915061151682856112ee565b60208201915061152682846112ee565b506020019695505050505050565b60006115408285611466565b6001820191506110f88284611283565b600061155c8286611466565b60018201915061156c82856112ee565b60208201915061157c82846112ee565b506020019392505050565b60006115938286611466565b6001820191506115a38285611466565b60018201915061157c82846112ee565b60006115bf8287611466565b6001820191506115cf8286611466565b6001820191506115df82856112ee565b6020820191506115ef82846112ee565b50602001949350505050565b6080810161160982856112dc565b610b2460208301846113fc565b6060810161162482856112dc565b610b246020830184611439565b60a0810161163f82886112dc565b61164c60208301876112e5565b611659604083018661145d565b6116666060830185611263565b81810360808301526101ab81846112ff565b602081016101c582846112e5565b602080825281016101c581611366565b602080825281016101c58161139f565b602080825281016101c5816113cf565b60405181810167ffffffffffffffff811182821017156116d557600080fd5b604052919050565b600067ffffffffffffffff8211156116f457600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006101c58261172b565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b60005b8381101561176457818101518382015260200161174c565b838111156114335750506000910152565b60006101c58260006101c5826117a1565b60006101c58261179b565b601f01601f191690565b60f81b90565b60601b90565b6117b081611718565b81146117bb57600080fd5b50565b6117b081611728565b6117b08161173756fea365627a7a72315820551af79fe2ac0178ba85ada64847ec9e985e9664006e8113199e30513745f3796c6578706572696d656e74616cf564736f6c63430005110040"

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
