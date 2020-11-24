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

// MessageTesterABI is the input ABI used to generate the binding from.
const MessageTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"addMessageToInbox\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"messageValueHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseERC20Message\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseERC721Message\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseEthMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"unmarshalOutgoingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

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
var MessageTesterBin = "0x608060405234801561001057600080fd5b5061165d806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063a3b392091161005b578063a3b3920914610391578063ec65668c146103b4578063fdaf43c114610480578063fe517bd0146100825761007d565b80636520427f146100825780636b0d35191461015a5780639aa86e86146102b0575b600080fd5b6101266004803603602081101561009857600080fd5b810190602081018135600160201b8111156100b257600080fd5b8201836020820111156100c457600080fd5b803590602001918460018302840111600160201b831117156100e557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104c7945050505050565b6040805194151585526001600160a01b03938416602086015291909216838201526060830191909152519081900360800190f35b6102006004803603604081101561017057600080fd5b810190602081018135600160201b81111561018a57600080fd5b82018360208201111561019c57600080fd5b803590602001918460018302840111600160201b831117156101bd57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506104ff915050565b60405180861515151581526020018581526020018460ff1660ff168152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610271578181015183820152602001610259565b50505050905090810190601f16801561029e5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b61037f600480360360c08110156102c657600080fd5b60ff823516916020810135916040820135916001600160a01b03606082013516916080820135919081019060c0810160a0820135600160201b81111561030b57600080fd5b82018360208201111561031d57600080fd5b803590602001918460018302840111600160201b8311171561033e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061053e945050505050565b60408051918252519081900360200190f35b61037f600480360360408110156103a757600080fd5b5080359060200135610561565b610458600480360360208110156103ca57600080fd5b810190602081018135600160201b8111156103e457600080fd5b8201836020820111156103f657600080fd5b803590602001918460018302840111600160201b8311171561041757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610574945050505050565b6040805193151584526001600160a01b03909216602084015282820152519081900360600190f35b61037f600480360360c081101561049657600080fd5b5060ff813516906001600160a01b036020820135169060408101359060608101359060808101359060a001356105a2565b60008060008060006104d761159e565b6104e0876105b2565b8051602082015160409092015192985096509450925050509193509193565b60008060008060606000806105126115be565b61051c8a8a610630565b80516020820151604090920151939e929d509b50995090975095505050505050565b600061055661055188888888888861071c565b610844565b979650505050505050565b600061056d8383610969565b9392505050565b6000806000806105826115dd565b61058b86610995565b805160209091015191989097509095509350505050565b60006105568787878787876109f1565b60006105bc61159e565b6048835110156105cf576000915061062b565b600c6105e1848263ffffffff610a5f16565b6001600160a01b031682526020016105ff848263ffffffff610a5f16565b6001600160a01b03166020830152601401610620848263ffffffff610abf16565b604083015250600191505b915091565b60008061063b6115be565b839150600085838151811061064c57fe5b016020015160019093019260f81c9050610664610b18565b60030160ff168160ff1614610680575060009250839150610715565b600061068c8785610b1d565b91965094509050846106a75750600093508492506107159050565b60ff8116835260006106b98886610b1d565b91975095509050856106d5575060009450859350610715915050565b6001600160a01b03811660208501526106ee8886610b9a565b604087015290965094508561070d575060009450859350610715915050565b506001945050505b9250925092565b6107246115f4565b60408051600680825260e08201909252606091816020015b6107446115f4565b81526020019060019003908161073c5790505090506107658860ff16610df1565b8160008151811061077257fe5b602002602001018190525061078687610df1565b8160018151811061079357fe5b60200260200101819052506107a786610df1565b816002815181106107b457fe5b60200260200101819052506107d1856001600160a01b0316610df1565b816003815181106107de57fe5b60200260200101819052506107f284610df1565b816004815181106107ff57fe5b60200260200101819052506108178360008551610ea3565b8160058151811061082457fe5b60200260200101819052506108388161101c565b98975050505050505050565b600061084e61112e565b60ff16826060015160ff16141561087157815161086a90611133565b9050610964565b610879611157565b60ff16826060015160ff1614156108975761086a826020015161115c565b61089f611259565b60ff16826060015160ff1614156108c1578151608083015161086a919061125e565b6108c9610b18565b60ff16826060015160ff161415610902576108e26115f4565b6108ef83604001516112af565b90506108fa81610844565b915050610964565b61090a611411565b60ff16826060015160ff16141561092357508051610964565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b919050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600061099f6115dd565b6034835110156109b2576000915061062b565b600c6109c4848263ffffffff610a5f16565b6001600160a01b031682526014016109e2848263ffffffff610abf16565b60208301525060019150915091565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6bffffffffffffffffffffffff191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b60008160140183511015610aaf576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b500160200151600160601b900490565b60008160200183511015610b0f576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b600390565b6000806000808551905084811080610b3757506021858203105b80610b5f5750610b4561112e565b60ff16868681518110610b5457fe5b016020015160f81c14155b15610b74575060009250839150829050610715565b600160218601610b8c8888840163ffffffff610abf16565b935093509350509250925092565b60008060606000610bab8686611416565b9195509350905083610bc1575060009250610715565b60208104601f8216600081610bd7576000610bda565b60015b60ff1683019050606083604051908082528060200260200182016040528015610c0d578160200160208202803883390190505b5090506060836040519080825280601f01601f191660200182016040528015610c3d576020820181803883390190505b5090506000805b84811015610d08576000610c588e8c611416565b919d509b5090508b610c77575060009a50610715975050505050505050565b81158015610c855750600087115b15610cd8578060005b88811015610cd157818160208110610ca257fe5b1a60f81b868281518110610cb257fe5b60200101906001600160f81b031916908160001a905350600101610c8e565b5050610cff565b8060001b858460018b030381518110610ced57fe5b60209081029190910101526001909201915b50600101610c44565b506000610d158d8b611478565b909a509050610d22610b18565b60ff168160ff1614610d405750600099506107159650505050505050565b60018a858560405160200180838051906020019060200280838360005b83811015610d75578181015183820152602001610d5d565b5050505090500182805190602001908083835b60208310610da75780518252601f199092019160209182019101610d88565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529a509a509a5050505050505050509250925092565b610df96115f4565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610e4f565b610e3c6115f4565b815260200190600190039081610e345790505b50905281526040805160008082526020828101909352919092019190610e8b565b610e786115f4565b815260200190600190039081610e705790505b50815260006020820152600160409091015292915050565b610eab6115f4565b60208204610eb76115f4565b610ebf61149f565b60408051600280825260608281019093529293509091816020015b610ee26115f4565b815260200190600190039081610eda57905050905060005b83811015610f6357610f1f610f1a89602084028a0163ffffffff610abf16565b610df1565b82600081518110610f2c57fe5b60200260200101819052508282600181518110610f4557fe5b6020026020010181905250610f59826112af565b9250600101610efa565b506020850615610fd9576000610f8588601f198989010163ffffffff610abf16565b9050602086066020036008021b610f9b81610df1565b82600081518110610fa857fe5b60200260200101819052508282600181518110610fc157fe5b6020026020010181905250610fd5826112af565b9250505b610fe285610df1565b81600081518110610fef57fe5b6020026020010181905250818160018151811061100857fe5b6020026020010181905250610556816112af565b6110246115f4565b61102e82516114e6565b61107f576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156110b65783818151811061109957fe5b602002602001015160800151820191508080600101915050611084565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190611110565b6110fd6115f4565b8152602001906001900390816110f55790505b50905281526020810194909452600360408501526060909301525090565b600090565b60408051602080820193909352815180820384018152908201909152805191012090565b600190565b600060028260400151511061116d57fe5b6040820151516111d25761117f611157565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b909316602185015260228085019190915282518085039091018152604290930190915281519101209050610964565b6111da611157565b826000015161120084604001516000815181106111f357fe5b6020026020010151610844565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b600290565b6000611268610b18565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b6112b76115f4565b600882511115611305576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611332578160200160208202803883390190505b508051909150600160005b82811015611395576113548682815181106111f357fe5b84828151811061136057fe5b60200260200101818152505085818151811061137857fe5b60200260200101516080015182019150808060010191505061133d565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156113da5781810151838201526020016113c2565b505050509050019250505060405160208183030381529060405280519060200120905061140781836114ed565b9695505050505050565b606490565b6000806000806114268686611478565b9093509050611433610b18565b60020160ff168160ff161461144c575060009250610715565b6114568684610b1d565b919550935091508361146c575060009250610715565b50600192509250925092565b6000808260010184848151811061148b57fe5b016020015190925060f81c90509250929050565b6114a76115f4565b604080516000808252602082019092526114e1916114db565b6114c86115f4565b8152602001906001900390816114c05790505b5061101c565b905090565b6008101590565b6114f56115f4565b6040805160a081018252848152815160608101835260008082526020828101829052845182815280820186529394908501939083019161154b565b6115386115f4565b8152602001906001900390816115305790505b50905281526040805160008082526020828101909352919092019190611587565b6115746115f4565b81526020019060019003908161156c5790505b508152600260208201526040019290925250919050565b604080516060810182526000808252602082018190529181019190915290565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b6040518060a001604052806000815260200161160e6115be565b81526060602082018190526000604083018190529101529056fea265627a7a72315820cbffad9d05a79b2bc59887b409693035100d466f6025cb69d4a86079678f5a7764736f6c63430005110032"

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
func (_MessageTester *MessageTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_MessageTester *MessageTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "addMessageToInbox", inbox, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "messageHash", messageType, sender, blockNumber, timestamp, inboxSeqNum, messageDataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "messageValueHash", messageType, blockNumber, timestamp, sender, inboxSeqNum, messageData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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
// Solidity: function parseERC20Message(bytes data) pure returns(bool valid, address token, address dest, uint256 value)
func (_MessageTester *MessageTesterCaller) ParseERC20Message(opts *bind.CallOpts, data []byte) (struct {
	Valid bool
	Token common.Address
	Dest  common.Address
	Value *big.Int
}, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "parseERC20Message", data)

	outstruct := new(struct {
		Valid bool
		Token common.Address
		Dest  common.Address
		Value *big.Int
	})

	outstruct.Valid = out[0].(bool)
	outstruct.Token = out[1].(common.Address)
	outstruct.Dest = out[2].(common.Address)
	outstruct.Value = out[3].(*big.Int)

	return *outstruct, err

}

// ParseERC20Message is a free data retrieval call binding the contract method 0x6520427f.
//
// Solidity: function parseERC20Message(bytes data) pure returns(bool valid, address token, address dest, uint256 value)
func (_MessageTester *MessageTesterSession) ParseERC20Message(data []byte) (struct {
	Valid bool
	Token common.Address
	Dest  common.Address
	Value *big.Int
}, error) {
	return _MessageTester.Contract.ParseERC20Message(&_MessageTester.CallOpts, data)
}

// ParseERC20Message is a free data retrieval call binding the contract method 0x6520427f.
//
// Solidity: function parseERC20Message(bytes data) pure returns(bool valid, address token, address dest, uint256 value)
func (_MessageTester *MessageTesterCallerSession) ParseERC20Message(data []byte) (struct {
	Valid bool
	Token common.Address
	Dest  common.Address
	Value *big.Int
}, error) {
	return _MessageTester.Contract.ParseERC20Message(&_MessageTester.CallOpts, data)
}

// ParseERC721Message is a free data retrieval call binding the contract method 0xfe517bd0.
//
// Solidity: function parseERC721Message(bytes data) pure returns(bool valid, address token, address dest, uint256 id)
func (_MessageTester *MessageTesterCaller) ParseERC721Message(opts *bind.CallOpts, data []byte) (struct {
	Valid bool
	Token common.Address
	Dest  common.Address
	Id    *big.Int
}, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "parseERC721Message", data)

	outstruct := new(struct {
		Valid bool
		Token common.Address
		Dest  common.Address
		Id    *big.Int
	})

	outstruct.Valid = out[0].(bool)
	outstruct.Token = out[1].(common.Address)
	outstruct.Dest = out[2].(common.Address)
	outstruct.Id = out[3].(*big.Int)

	return *outstruct, err

}

// ParseERC721Message is a free data retrieval call binding the contract method 0xfe517bd0.
//
// Solidity: function parseERC721Message(bytes data) pure returns(bool valid, address token, address dest, uint256 id)
func (_MessageTester *MessageTesterSession) ParseERC721Message(data []byte) (struct {
	Valid bool
	Token common.Address
	Dest  common.Address
	Id    *big.Int
}, error) {
	return _MessageTester.Contract.ParseERC721Message(&_MessageTester.CallOpts, data)
}

// ParseERC721Message is a free data retrieval call binding the contract method 0xfe517bd0.
//
// Solidity: function parseERC721Message(bytes data) pure returns(bool valid, address token, address dest, uint256 id)
func (_MessageTester *MessageTesterCallerSession) ParseERC721Message(data []byte) (struct {
	Valid bool
	Token common.Address
	Dest  common.Address
	Id    *big.Int
}, error) {
	return _MessageTester.Contract.ParseERC721Message(&_MessageTester.CallOpts, data)
}

// ParseEthMessage is a free data retrieval call binding the contract method 0xec65668c.
//
// Solidity: function parseEthMessage(bytes data) pure returns(bool valid, address dest, uint256 value)
func (_MessageTester *MessageTesterCaller) ParseEthMessage(opts *bind.CallOpts, data []byte) (struct {
	Valid bool
	Dest  common.Address
	Value *big.Int
}, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "parseEthMessage", data)

	outstruct := new(struct {
		Valid bool
		Dest  common.Address
		Value *big.Int
	})

	outstruct.Valid = out[0].(bool)
	outstruct.Dest = out[1].(common.Address)
	outstruct.Value = out[2].(*big.Int)

	return *outstruct, err

}

// ParseEthMessage is a free data retrieval call binding the contract method 0xec65668c.
//
// Solidity: function parseEthMessage(bytes data) pure returns(bool valid, address dest, uint256 value)
func (_MessageTester *MessageTesterSession) ParseEthMessage(data []byte) (struct {
	Valid bool
	Dest  common.Address
	Value *big.Int
}, error) {
	return _MessageTester.Contract.ParseEthMessage(&_MessageTester.CallOpts, data)
}

// ParseEthMessage is a free data retrieval call binding the contract method 0xec65668c.
//
// Solidity: function parseEthMessage(bytes data) pure returns(bool valid, address dest, uint256 value)
func (_MessageTester *MessageTesterCallerSession) ParseEthMessage(data []byte) (struct {
	Valid bool
	Dest  common.Address
	Value *big.Int
}, error) {
	return _MessageTester.Contract.ParseEthMessage(&_MessageTester.CallOpts, data)
}

// UnmarshalOutgoingMessage is a free data retrieval call binding the contract method 0x6b0d3519.
//
// Solidity: function unmarshalOutgoingMessage(bytes data, uint256 startOffset) pure returns(bool, uint256, uint8, address, bytes)
func (_MessageTester *MessageTesterCaller) UnmarshalOutgoingMessage(opts *bind.CallOpts, data []byte, startOffset *big.Int) (bool, *big.Int, uint8, common.Address, []byte, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "unmarshalOutgoingMessage", data, startOffset)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(uint8), *new(common.Address), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return out0, out1, out2, out3, out4, err

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
