// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package messagetester

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

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205130ee8069d4d45b671687a9bd9471ae9b0f29d9a9560a567b2a039c58a80e1a64736f6c63430005110032"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
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
var MessageTesterBin = "0x608060405234801561001057600080fd5b50611989806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063ec65668c1161005b578063ec65668c1461010e578063f23ba5fc1461012f578063fdaf43c114610142578063fe517bd01461008d57610088565b80636520427f1461008d5780636b0d3519146100b75780639aa86e86146100db578063a3b39209146100fb575b600080fd5b6100a061009b366004611203565b610155565b6040516100ae929190611761565b60405180910390f35b6100ca6100c5366004611238565b610172565b6040516100ae959493929190611797565b6100ee6100e93660046112f5565b6101b1565b6040516100ae91906117de565b6100ee610109366004611168565b6101d4565b61012161011c366004611203565b6101e9565b6040516100ae92919061177c565b6100ee61013d3660046111a2565b6101fc565b6100ee61015036600461126e565b610226565b600061015f611069565b61016883610236565b915091505b915091565b6000806000806060600080610185611089565b61018f8a8a6102b3565b80516020820151604090920151939e929d509b50995090975095505050505050565b60006101c96101c488888888888861039f565b6104c7565b979650505050505050565b60006101e083836105b1565b90505b92915050565b60006101f36110a8565b610168836105e4565b600061021d6101c461020e8787610640565b6102188686610640565b6106f1565b95945050505050565b60006101c987878787878761076f565b6000610240611069565b604883511015610253576000915061016d565b600c610265848263ffffffff6107ae16565b6001600160a01b03168252602001610283848263ffffffff6107ae16565b6001600160a01b031660208301526014016102a4848263ffffffff6107d116565b60408301525060019150915091565b6000806102be611089565b83915060008583815181106102cf57fe5b016020015160019093019260f81c90506102e76107ed565b60030160ff168160ff1614610303575060009250839150610398565b600061030f87856107f2565b919650945090508461032a5750600093508492506103989050565b60ff81168352600061033c88866107f2565b9197509550905085610358575060009450859350610398915050565b6001600160a01b03811660208501526103718886610869565b6040870152909650945085610390575060009450859350610398915050565b506001945050505b9250925092565b6103a76110bf565b60408051600680825260e08201909252606091816020015b6103c76110bf565b8152602001906001900390816103bf5790505090506103e88860ff16610ab5565b816000815181106103f557fe5b602002602001018190525061040987610ab5565b8160018151811061041657fe5b602002602001018190525061042a86610ab5565b8160028151811061043757fe5b6020026020010181905250610454856001600160a01b0316610ab5565b8160038151811061046157fe5b602002602001018190525061047584610ab5565b8160048151811061048257fe5b602002602001018190525061049a8360008551610b67565b816005815181106104a757fe5b60200260200101819052506104bb81610cfa565b98975050505050505050565b606081015160009060ff166104e85781516104e190610ddc565b90506105ac565b606082015160ff1660011415610505576104e18260200151610e0c565b606082015160ff166002141561052657815160808301516104e19190610e9d565b600360ff16826060015160ff161015801561054a57506060820151600c60ff909116105b15610573576105576110bf565b61056083610eb5565b905061056b816104c7565b9150506105ac565b606082015160ff166064141561058b575080516105ac565b60405162461bcd60e51b81526004016105a39061181c565b60405180910390fd5b919050565b600082826040516020016105c69291906115f5565b60405160208183030381529060405280519060200120905092915050565b60006105ee6110a8565b603483511015610601576000915061016d565b600c610613848263ffffffff6107ae16565b6001600160a01b03168252601401610631848263ffffffff6107d116565b60208301525060019150915091565b6106486110bf565b6040805160a081018252848152815160608101835260008082526020828101829052845182815280820186529394908501939083019161069e565b61068b6110bf565b8152602001906001900390816106835790505b509052815260408051600080825260208281019093529190920191906106da565b6106c76110bf565b8152602001906001900390816106bf5790505b508152600260208201526040019290925250919050565b6106f96110bf565b6040805160028082526060828101909352816020015b6107176110bf565b81526020019060019003908161070f579050509050838160008151811061073a57fe5b6020026020010181905250828160018151811061075357fe5b602002602001018190525061076781610cfa565b949350505050565b600086868686868660405160200161078c96959493929190611630565b6040516020818303038152906040528051906020012090509695505050505050565b600081601401835110156107c157600080fd5b500160200151600160601b900490565b600081602001835110156107e457600080fd5b50016020015190565b600390565b600080600080855190508481108061080c57506021858203105b8061082e5750600060ff1686868151811061082357fe5b016020015160f81c14155b15610843575060009250839150829050610398565b60016021860161085b8888840163ffffffff6107d116565b935093509350509250925092565b6000806060839150600085838151811061087f57fe5b016020015160019093019260f81c9050600581146108a1575060009250610398565b60006108ad87856107f2565b91965094509050846108c55750600093506103989050565b60208104601f82166000816108db5760006108de565b60015b60ff1683019050606083604051908082528060200260200182016040528015610911578160200160208202803883390190505b5090506060836040519080825280601f01601f191660200182016040528015610941576020820181803883390190505b5090506000805b84811015610a44578d8b8151811061095c57fe5b01602001516001909b019a60f81c985060058914610987575060009a50610398975050505050505050565b60006109938f8d6107f2565b919e509c5090508c6109b3575060009b5061039898505050505050505050565b811580156109c15750600087115b15610a14578060005b88811015610a0d578181602081106109de57fe5b1a60f81b8682815181106109ee57fe5b60200101906001600160f81b031916908160001a9053506001016109ca565b5050610a3b565b8060001b858460018b030381518110610a2957fe5b60209081029190910101526001909201915b50600101610948565b508c8a81518110610a5157fe5b01602001516001909a019960f81c975060038814610a7b5750600099506103989650505050505050565b60018a8484604051602001610a919291906115dd565b6040516020818303038152906040529a509a509a5050505050505050509250925092565b610abd6110bf565b6040805160a0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191610b13565b610b006110bf565b815260200190600190039081610af85790505b50905281526040805160008082526020828101909352919092019190610b4f565b610b3c6110bf565b815260200190600190039081610b345790505b50815260006020820152600160409091015292915050565b610b6f6110bf565b602082046000610b7d610f91565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b84811015610c1b57610bce6101c4610bc98b602085028c0163ffffffff6107d116565b610ab5565b82600081518110610bdb57fe5b6020026020010181815250508382600181518110610bf557fe5b602002602001018181525050600283019250610c118284610fb2565b9350600101610ba6565b506020860615610c9d576000610c3d89601f198a8a010163ffffffff6107d116565b9050602087066020036008021b610c566101c482610ab5565b82600081518110610c6357fe5b6020026020010181815250508382600181518110610c7d57fe5b602002602001018181525050600283019250610c998284610fb2565b9350505b610ca96101c487610ab5565b81600081518110610cb657fe5b6020026020010181815250508281600181518110610cd057fe5b602002602001018181525050600282019150610cec8183610fbd565b9450505050505b9392505050565b610d026110bf565b610d0c8251610fdc565b610d285760405162461bcd60e51b81526004016105a3906117fc565b600160005b8351811015610d5f57838181518110610d4257fe5b602002602001015160800151820191508080600101915050610d2d565b506040805160a0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190610db9565b610da66110bf565b815260200190600190039081610d9e5790505b509052815260208101859052935160030160ff1660408501526060909301525090565b600081604051602001610def919061161b565b604051602081830303815290604052805190602001209050919050565b6000600282604001515110610e1d57fe5b604082015151610e5f578151602080840151604051610e4293600193909291016116ed565b6040516020818303038152906040528051906020012090506105ac565b60018260000151610e878460400151600081518110610e7a57fe5b60200260200101516104c7565b602080860151604051610def9594939201611719565b6000600383836040516020016105c6939291906116b6565b610ebd6110bf565b610ec682610fe3565b610ee25760405162461bcd60e51b81526004016105a3906117ec565b60088260400151511115610f085760405162461bcd60e51b81526004016105a39061180c565b6060826040015151604051908082528060200260200182016040528015610f39578160200160208202803883390190505b50805190915060005b81811015610f82576000610f5f86604001518381518110610e7a57fe5b905080848381518110610f6e57fe5b602090810291909101015250600101610f42565b50610767828560800151610fbd565b60408051600080825260208201909252610fac816001610fb2565b91505090565b60006101e06101c484845b610fc56110bf565b6000610fd084610ff2565b90506107678184610640565b6008101590565b60006101e3826060015161104b565b60006008825111156110165760405162461bcd60e51b81526004016105a39061180c565b600082518360405160200161102c92919061169a565b60408051808303601f1901815291905280516020909101209392505050565b6000600c60ff83161080156101e3575050600360ff91909116101590565b604080516060810182526000808252602082018190529181019190915290565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b6040518060a00160405280600081526020016110d9611089565b815260606020820181905260006040830181905291015290565b80356101e38161191d565b80356101e381611934565b600082601f83011261111a57600080fd5b813561112d61112882611853565b61182c565b9150808252602083016020830185838301111561114957600080fd5b6111548382846118b3565b50505092915050565b80356101e38161193d565b6000806040838503121561117b57600080fd5b600061118785856110fe565b9250506020611198858286016110fe565b9150509250929050565b600080600080608085870312156111b857600080fd5b60006111c487876110fe565b94505060206111d5878288016110fe565b93505060406111e6878288016110fe565b92505060606111f7878288016110fe565b91505092959194509250565b60006020828403121561121557600080fd5b813567ffffffffffffffff81111561122c57600080fd5b61076784828501611109565b6000806040838503121561124b57600080fd5b823567ffffffffffffffff81111561126257600080fd5b61118785828601611109565b60008060008060008060c0878903121561128757600080fd5b6000611293898961115d565b96505060206112a489828a016110f3565b95505060406112b589828a016110fe565b94505060606112c689828a016110fe565b93505060806112d789828a016110fe565b92505060a06112e889828a016110fe565b9150509295509295509295565b60008060008060008060c0878903121561130e57600080fd5b600061131a898961115d565b965050602061132b89828a016110fe565b955050604061133c89828a016110fe565b945050606061134d89828a016110f3565b935050608061135e89828a016110fe565b92505060a087013567ffffffffffffffff81111561137b57600080fd5b6112e889828a01611109565b6000611393838361141d565b505060200190565b6113a48161188e565b82525050565b6113a46113b68261188e565b6118eb565b60006113c682611881565b6113d081856105ac565b93506113db8361187b565b8060005b838110156114095781516113f38882611387565b97506113fe8361187b565b9250506001016113df565b509495945050505050565b6113a481611899565b6113a48161189e565b6113a46114328261189e565b61189e565b600061144282611881565b61144c8185611885565b935061145c8185602086016118bf565b61146581611907565b9093019392505050565b600061147a82611881565b61148481856105ac565b93506114948185602086016118bf565b9290920192915050565b60006114ab601283611885565b714d757374206265205475706c65207479706560701b815260200192915050565b60006114d9601a83611885565b7f5475706c65206d75737420686176652076616c69642073697a65000000000000815260200192915050565b6000611512601483611885565b73092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b815260200192915050565b6000611542601183611885565b70496e76616c6964207479706520636f646560781b815260200192915050565b80516060830190611573848261139b565b506020820151611586602085018261139b565b506040820151611599604085018261141d565b50505050565b805160408301906115b0848261139b565b506020820151611599602085018261141d565b6113a4816118ad565b6113a46115d8826118ad565b6118fc565b60006115e982856113bb565b9150610767828461146f565b60006116018285611426565b6020820191506116118284611426565b5060200192915050565b60006116278284611426565b50602001919050565b600061163c82896115cc565b60018201915061164c82886113aa565b60148201915061165c8287611426565b60208201915061166c8286611426565b60208201915061167c8285611426565b60208201915061168c8284611426565b506020019695505050505050565b60006116a682856115cc565b60018201915061076782846113bb565b60006116c282866115cc565b6001820191506116d28285611426565b6020820191506116e28284611426565b506020019392505050565b60006116f982866115cc565b60018201915061170982856115cc565b6001820191506116e28284611426565b600061172582876115cc565b60018201915061173582866115cc565b6001820191506117458285611426565b6020820191506117558284611426565b50602001949350505050565b6080810161176f8285611414565b610cf36020830184611562565b6060810161178a8285611414565b610cf3602083018461159f565b60a081016117a58288611414565b6117b2602083018761141d565b6117bf60408301866115c3565b6117cc606083018561139b565b81810360808301526101c98184611437565b602081016101e3828461141d565b602080825281016101e38161149e565b602080825281016101e3816114cc565b602080825281016101e381611505565b602080825281016101e381611535565b60405181810167ffffffffffffffff8111828210171561184b57600080fd5b604052919050565b600067ffffffffffffffff82111561186a57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006101e3826118a1565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b60005b838110156118da5781810151838201526020016118c2565b838111156115995750506000910152565b60006101e38260006101e382611917565b60006101e382611911565b601f01601f191690565b60f81b90565b60601b90565b6119268161188e565b811461193157600080fd5b50565b6119268161189e565b611926816118ad56fea365627a7a723158203fcd09be852155b5a81eeb2eef45e4a121bc4b5c606081695c46de34baffc6a86c6578706572696d656e74616cf564736f6c63430005110040"

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

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a1f456343425d4a1ac63ee2c237ed76613d9d2e51071989d30561b052f16401c64736f6c63430005110032"

// DeployMessages deploys a new Ethereum contract, binding an instance of Messages to it.
func DeployMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Messages, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// Messages is an auto generated Go binding around an Ethereum contract.
type Messages struct {
	MessagesCaller     // Read-only binding to the contract
	MessagesTransactor // Write-only binding to the contract
	MessagesFilterer   // Log filterer for contract events
}

// MessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagesSession struct {
	Contract     *Messages         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagesCallerSession struct {
	Contract *MessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagesTransactorSession struct {
	Contract     *MessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagesRaw struct {
	Contract *Messages // Generic contract binding to access the raw methods on
}

// MessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagesCallerRaw struct {
	Contract *MessagesCaller // Generic read-only contract binding to access the raw methods on
}

// MessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagesTransactorRaw struct {
	Contract *MessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessages creates a new instance of Messages, bound to a specific deployed contract.
func NewMessages(address common.Address, backend bind.ContractBackend) (*Messages, error) {
	contract, err := bindMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// NewMessagesCaller creates a new read-only instance of Messages, bound to a specific deployed contract.
func NewMessagesCaller(address common.Address, caller bind.ContractCaller) (*MessagesCaller, error) {
	contract, err := bindMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesCaller{contract: contract}, nil
}

// NewMessagesTransactor creates a new write-only instance of Messages, bound to a specific deployed contract.
func NewMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagesTransactor, error) {
	contract, err := bindMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesTransactor{contract: contract}, nil
}

// NewMessagesFilterer creates a new log filterer instance of Messages, bound to a specific deployed contract.
func NewMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagesFilterer, error) {
	contract, err := bindMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagesFilterer{contract: contract}, nil
}

// bindMessages binds a generic wrapper to an already deployed contract.
func bindMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.MessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820142891281d8f03ec873930437a8ca5e0c61de65eedd6218afc3087c76cc8fbbd64736f6c63430005110032"

// DeployValue deploys a new Ethereum contract, binding an instance of Value to it.
func DeployValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Value, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// Value is an auto generated Go binding around an Ethereum contract.
type Value struct {
	ValueCaller     // Read-only binding to the contract
	ValueTransactor // Write-only binding to the contract
	ValueFilterer   // Log filterer for contract events
}

// ValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueSession struct {
	Contract     *Value            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueCallerSession struct {
	Contract *ValueCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueTransactorSession struct {
	Contract     *ValueTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueRaw struct {
	Contract *Value // Generic contract binding to access the raw methods on
}

// ValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueCallerRaw struct {
	Contract *ValueCaller // Generic read-only contract binding to access the raw methods on
}

// ValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueTransactorRaw struct {
	Contract *ValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValue creates a new instance of Value, bound to a specific deployed contract.
func NewValue(address common.Address, backend bind.ContractBackend) (*Value, error) {
	contract, err := bindValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// NewValueCaller creates a new read-only instance of Value, bound to a specific deployed contract.
func NewValueCaller(address common.Address, caller bind.ContractCaller) (*ValueCaller, error) {
	contract, err := bindValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueCaller{contract: contract}, nil
}

// NewValueTransactor creates a new write-only instance of Value, bound to a specific deployed contract.
func NewValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueTransactor, error) {
	contract, err := bindValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTransactor{contract: contract}, nil
}

// NewValueFilterer creates a new log filterer instance of Value, bound to a specific deployed contract.
func NewValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueFilterer, error) {
	contract, err := bindValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueFilterer{contract: contract}, nil
}

// bindValue binds a generic wrapper to an already deployed contract.
func bindValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.ValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.contract.Transact(opts, method, params...)
}
