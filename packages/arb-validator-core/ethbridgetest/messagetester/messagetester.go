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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158207d8758ad86f665628dab702fc2bf77ffff692cbcea924b7642d8a631868fabc164736f6c634300050f0032"

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
var MessageTesterBin = "0x608060405234801561001057600080fd5b506119e2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063ec65668c1161005b578063ec65668c1461010e578063f23ba5fc1461012f578063fdaf43c114610142578063fe517bd01461008d57610088565b80636520427f1461008d5780636b0d3519146100b75780639aa86e86146100db578063a3b39209146100fb575b600080fd5b6100a061009b36600461121d565b610155565b6040516100ae9291906117aa565b60405180910390f35b6100ca6100c5366004611252565b610172565b6040516100ae9594939291906117e0565b6100ee6100e936600461130f565b6101b1565b6040516100ae9190611827565b6100ee610109366004611182565b6101d4565b61012161011c36600461121d565b6101e9565b6040516100ae9291906117c5565b6100ee61013d3660046111bc565b6101fc565b6100ee610150366004611288565b610228565b600061015f611055565b61016883610238565b915091505b915091565b6000806000806060600080610185611075565b61018f8a8a6102b5565b80516020820151604090920151939e929d509b50995090975095505050505050565b60006101c96101c48888888888886103a1565b6104c9565b979650505050505050565b60006101e083836105a6565b90505b92915050565b60006101f3611094565b610168836105d9565b600061021d6101c461020e8787610635565b6102188686610635565b6106c0565b90505b949350505050565b60006101c9878787878787610736565b6000610242611055565b604883511015610255576000915061016d565b600c610267848263ffffffff61077516565b6001600160a01b03168252602001610285848263ffffffff61077516565b6001600160a01b031660208301526014016102a6848263ffffffff61079816565b60408301525060019150915091565b6000806102c0611075565b83915060008583815181106102d157fe5b016020015160019093019260f81c90506102e96107b4565b60030160ff168160ff161461030557506000925083915061039a565b600061031187856107b9565b919650945090508461032c57506000935084925061039a9050565b60ff81168352600061033e88866107b9565b919750955090508561035a57506000945085935061039a915050565b6001600160a01b03811660208501526103738886610830565b604087015290965094508561039257506000945085935061039a915050565b506001945050505b9250925092565b6103a96110ab565b60408051600680825260e08201909252606091816020015b6103c96110ab565b8152602001906001900390816103c15790505090506103ea8860ff16610a7c565b816000815181106103f757fe5b602002602001018190525061040b87610a7c565b8160018151811061041857fe5b602002602001018190525061042c86610a7c565b8160028151811061043957fe5b6020026020010181905250610456856001600160a01b0316610a7c565b8160038151811061046357fe5b602002602001018190525061047784610a7c565b8160048151811061048457fe5b602002602001018190525061049c8360008551610b08565b816005815181106104a957fe5b60200260200101819052506104bd81610c9b565b98975050505050505050565b606081015160009060ff166104ea5781516104e390610d5d565b90506105a1565b606082015160ff166001141561051d5760208083015180516040820151606083015192909301516104e393919290610d8d565b606082015160ff1660021415610536576104e382610e00565b600360ff16826060015160ff161015801561055a57506060820151600c60ff909116105b15610568576104e382610e3b565b606082015160ff1660641415610580575080516105a1565b60405162461bcd60e51b815260040161059890611875565b60405180910390fd5b919050565b600082826040516020016105bb92919061163e565b60405160208183030381529060405280519060200120905092915050565b60006105e3611094565b6034835110156105f6576000915061016d565b600c610608848263ffffffff61077516565b6001600160a01b03168252601401610626848263ffffffff61079816565b60208301525060019150915091565b61063d6110ab565b6040805160a0808201835285825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916106a9565b6106966110ab565b81526020019060019003908161068e5790505b508152600260208201526040019290925250919050565b6106c86110ab565b6040805160028082526060828101909352816020015b6106e66110ab565b8152602001906001900390816106de579050509050838160008151811061070957fe5b6020026020010181905250828160018151811061072257fe5b602002602001018190525061022081610c9b565b600086868686868660405160200161075396959493929190611679565b6040516020818303038152906040528051906020012090509695505050505050565b6000816014018351101561078857600080fd5b500160200151600160601b900490565b600081602001835110156107ab57600080fd5b50016020015190565b600390565b60008060008085519050848110806107d357506021858203105b806107f55750600060ff168686815181106107ea57fe5b016020015160f81c14155b1561080a57506000925083915082905061039a565b6001602186016108228888840163ffffffff61079816565b935093509350509250925092565b6000806060839150600085838151811061084657fe5b016020015160019093019260f81c90506005811461086857506000925061039a565b600061087487856107b9565b919650945090508461088c57506000935061039a9050565b60208104601f82166000816108a25760006108a5565b60015b60ff16830190506060836040519080825280602002602001820160405280156108d8578160200160208202803883390190505b5090506060836040519080825280601f01601f191660200182016040528015610908576020820181803883390190505b5090506000805b84811015610a0b578d8b8151811061092357fe5b01602001516001909b019a60f81c98506005891461094e575060009a5061039a975050505050505050565b600061095a8f8d6107b9565b919e509c5090508c61097a575060009b5061039a98505050505050505050565b811580156109885750600087115b156109db578060005b888110156109d4578181602081106109a557fe5b1a60f81b8682815181106109b557fe5b60200101906001600160f81b031916908160001a905350600101610991565b5050610a02565b8060001b858460018b0303815181106109f057fe5b60209081029190910101526001909201915b5060010161090f565b508c8a81518110610a1857fe5b01602001516001909a019960f81c975060038814610a4257506000995061039a9650505050505050565b60018a8484604051602001610a58929190611626565b6040516020818303038152906040529a509a509a5050505050505050509250925092565b610a846110ab565b6040805160a080820183528482528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191610af0565b610add6110ab565b815260200190600190039081610ad55790505b50815260006020820152600160409091015292915050565b610b106110ab565b602082046000610b1e610e59565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b84811015610bbc57610b6f6101c4610b6a8b602085028c0163ffffffff61079816565b610a7c565b82600081518110610b7c57fe5b6020026020010181815250508382600181518110610b9657fe5b602002602001018181525050600283019250610bb28284610e7a565b9350600101610b47565b506020860615610c3e576000610bde89601f198a8a010163ffffffff61079816565b9050602087066020036008021b610bf76101c482610a7c565b82600081518110610c0457fe5b6020026020010181815250508382600181518110610c1e57fe5b602002602001018181525050600283019250610c3a8284610e7a565b9350505b610c4a6101c487610a7c565b81600081518110610c5757fe5b6020026020010181815250508281600181518110610c7157fe5b602002602001018181525050600282019150610c8d8183610e99565b9450505050505b9392505050565b610ca36110ab565b610cad8251610eb8565b610cc95760405162461bcd60e51b815260040161059890611845565b600160005b8351811015610d0057838181518110610ce357fe5b602002602001015160800151820191508080600101915050610cce565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b600081604051602001610d709190611664565b604051602081830303815290604052805190602001209050919050565b60008315610dca576001858484604051602001610dad9493929190611762565b604051602081830303815290604052805190602001209050610220565b60018583604051602001610de093929190611736565b604051602081830303815290604052805190602001209050949350505050565b606081015160009060ff16600214610e2a5760405162461bcd60e51b815260040161059890611855565b815160808301516101e39190610ebf565b6000610e456110ab565b610e4e83610ed7565b9050610c9481610e00565b60408051600080825260208201909252610e74816001610e7a565b91505090565b6000610e846110ab565b610e8e8484610e99565b905061022081610e00565b610ea16110ab565b6000610eac84610f23565b90506102208184610635565b6008101590565b6000600383836040516020016105bb939291906116ff565b610edf6110ab565b610ee882610f7c565b610f045760405162461bcd60e51b815260040161059890611835565b6060610f138360400151610f8b565b9050610c94818460800151610e99565b6000600882511115610f475760405162461bcd60e51b815260040161059890611865565b6000825183604051602001610f5d9291906116e3565b60408051808303601f1901815291905280516020909101209392505050565b60006101e38260600151611037565b6060600882511115610faf5760405162461bcd60e51b815260040161059890611865565b60608251604051908082528060200260200182016040528015610fdc578160200160208202803883390190505b50805190915060005b8181101561102e57600061100b868381518110610ffe57fe5b60200260200101516104c9565b90508084838151811061101a57fe5b602090810291909101015250600101610fe5565b50909392505050565b6000600c60ff83161080156101e3575050600360ff91909116101590565b604080516060810182526000808252602082018190529181019190915290565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b6040518060a00160405280600081526020016110c56110df565b815260606020820181905260006040830181905291015290565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b80356101e381611976565b80356101e38161198d565b600082601f83011261113457600080fd5b8135611147611142826118ac565b611885565b9150808252602083016020830185838301111561116357600080fd5b61116e83828461190c565b50505092915050565b80356101e381611996565b6000806040838503121561119557600080fd5b60006111a18585611118565b92505060206111b285828601611118565b9150509250929050565b600080600080608085870312156111d257600080fd5b60006111de8787611118565b94505060206111ef87828801611118565b935050604061120087828801611118565b925050606061121187828801611118565b91505092959194509250565b60006020828403121561122f57600080fd5b813567ffffffffffffffff81111561124657600080fd5b61022084828501611123565b6000806040838503121561126557600080fd5b823567ffffffffffffffff81111561127c57600080fd5b6111a185828601611123565b60008060008060008060c087890312156112a157600080fd5b60006112ad8989611177565b96505060206112be89828a0161110d565b95505060406112cf89828a01611118565b94505060606112e089828a01611118565b93505060806112f189828a01611118565b92505060a061130289828a01611118565b9150509295509295509295565b60008060008060008060c0878903121561132857600080fd5b60006113348989611177565b965050602061134589828a01611118565b955050604061135689828a01611118565b945050606061136789828a0161110d565b935050608061137889828a01611118565b92505060a087013567ffffffffffffffff81111561139557600080fd5b61130289828a01611123565b60006113ad8383611437565b505060200190565b6113be816118e7565b82525050565b6113be6113d0826118e7565b611944565b60006113e0826118da565b6113ea81856105a1565b93506113f5836118d4565b8060005b8381101561142357815161140d88826113a1565b9750611418836118d4565b9250506001016113f9565b509495945050505050565b6113be816118f2565b6113be816118f7565b6113be61144c826118f7565b6118f7565b600061145c826118da565b61146681856118de565b9350611476818560208601611918565b61147f81611960565b9093019392505050565b6000611494826118da565b61149e81856105a1565b93506114ae818560208601611918565b9290920192915050565b60006114c56012836118de565b714d757374206265205475706c65207479706560701b815260200192915050565b60006114f3601a836118de565b7f5475706c65206d75737420686176652076616c69642073697a65000000000000815260200192915050565b600061152c6013836118de565b7209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b815260200192915050565b600061155b6014836118de565b73092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b815260200192915050565b600061158b6011836118de565b70496e76616c6964207479706520636f646560781b815260200192915050565b805160608301906115bc84826113b5565b5060208201516115cf60208501826113b5565b5060408201516115e26040850182611437565b50505050565b805160408301906115f984826113b5565b5060208201516115e26020850182611437565b6113be81611906565b6113be61162182611906565b611955565b600061163282856113d5565b91506102208284611489565b600061164a8285611440565b60208201915061165a8284611440565b5060200192915050565b60006116708284611440565b50602001919050565b60006116858289611615565b60018201915061169582886113c4565b6014820191506116a58287611440565b6020820191506116b58286611440565b6020820191506116c58285611440565b6020820191506116d58284611440565b506020019695505050505050565b60006116ef8285611615565b60018201915061022082846113d5565b600061170b8286611615565b60018201915061171b8285611440565b60208201915061172b8284611440565b506020019392505050565b60006117428286611615565b6001820191506117528285611615565b60018201915061172b8284611440565b600061176e8287611615565b60018201915061177e8286611615565b60018201915061178e8285611440565b60208201915061179e8284611440565b50602001949350505050565b608081016117b8828561142e565b610c9460208301846115ab565b606081016117d3828561142e565b610c9460208301846115e8565b60a081016117ee828861142e565b6117fb6020830187611437565b611808604083018661160c565b61181560608301856113b5565b81810360808301526101c98184611451565b602081016101e38284611437565b602080825281016101e3816114b8565b602080825281016101e3816114e6565b602080825281016101e38161151f565b602080825281016101e38161154e565b602080825281016101e38161157e565b60405181810167ffffffffffffffff811182821017156118a457600080fd5b604052919050565b600067ffffffffffffffff8211156118c357600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006101e3826118fa565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b60005b8381101561193357818101518382015260200161191b565b838111156115e25750506000910152565b60006101e38260006101e382611970565b60006101e38261196a565b601f01601f191690565b60f81b90565b60601b90565b61197f816118e7565b811461198a57600080fd5b50565b61197f816118f7565b61197f8161190656fea365627a7a723158209032fd1ff874eb694cebf22e1ed34084f90000dbab39d597d32494c4e0d3af206c6578706572696d656e74616cf564736f6c634300050f0040"

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
var MessagesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820288f926630ff090699d165f91860f5c7a9370fc5aaa71d38fc10b008fb025f7b64736f6c634300050f0032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820756e12f9a680c5aa2d821f4845604d9e478c3e9e61704ebc03fff3bdaf85e52964736f6c634300050f0032"

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
