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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820df1cab9f5706f831964c6cfc36e8d53dfb187147670225cce9914aa402e7580164736f6c63430005110032"

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
var MessageTesterBin = "0x608060405234801561001057600080fd5b506118e9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063ec65668c1161005b578063ec65668c1461010e578063f23ba5fc1461012f578063fdaf43c114610142578063fe517bd01461008d57610088565b80636520427f1461008d5780636b0d3519146100b75780639aa86e86146100db578063a3b39209146100fb575b600080fd5b6100a061009b36600461115f565b610155565b6040516100ae9291906116b1565b60405180910390f35b6100ca6100c5366004611194565b610172565b6040516100ae9594939291906116e7565b6100ee6100e9366004611251565b6101b1565b6040516100ae919061172e565b6100ee6101093660046110c4565b6101d4565b61012161011c36600461115f565b6101e9565b6040516100ae9291906116cc565b6100ee61013d3660046110fe565b6101fc565b6100ee6101503660046111ca565b610228565b600061015f610f97565b61016883610238565b915091505b915091565b6000806000806060600080610185610fb7565b61018f8a8a6102b5565b80516020820151604090920151939e929d509b50995090975095505050505050565b60006101c96101c48888888888886103a1565b6104c9565b979650505050505050565b60006101e083836105a6565b90505b92915050565b60006101f3610fd6565b610168836105d9565b600061021d6101c461020e8787610635565b6102188686610635565b6106c0565b90505b949350505050565b60006101c9878787878787610736565b6000610242610f97565b604883511015610255576000915061016d565b6000610267848263ffffffff61077516565b6001600160a01b03168252601401610285848263ffffffff61077516565b6001600160a01b031660208301526014016102a6848263ffffffff61079816565b60408301525060019150915091565b6000806102c0610fb7565b83915060008583815181106102d157fe5b016020015160019093019260f81c90506102e96107b4565b60030160ff168160ff161461030557506000925083915061039a565b600061031187856107b9565b919650945090508461032c57506000935084925061039a9050565b60ff81168352600061033e88866107b9565b919750955090508561035a57506000945085935061039a915050565b6001600160a01b03811660208501526103738886610830565b604087015290965094508561039257506000945085935061039a915050565b506001945050505b9250925092565b6103a9610fed565b60408051600680825260e08201909252606091816020015b6103c9610fed565b8152602001906001900390816103c15790505090506103ea8860ff16610937565b816000815181106103f757fe5b602002602001018190525061040b87610937565b8160018151811061041857fe5b602002602001018190525061042c86610937565b8160028151811061043957fe5b6020026020010181905250610456856001600160a01b0316610937565b8160038151811061046357fe5b602002602001018190525061047784610937565b8160048151811061048457fe5b602002602001018190525061049c83600085516109c3565b816005815181106104a957fe5b60200260200101819052506104bd81610b5d565b98975050505050505050565b606081015160009060ff166104ea5781516104e390610c1f565b90506105a1565b606082015160ff166001141561051d5760208083015180516040820151606083015192909301516104e393919290610c4f565b606082015160ff1660021415610536576104e382610cc2565b600360ff16826060015160ff161015801561055a57506060820151600c60ff909116105b15610568576104e382610cfd565b606082015160ff1660641415610580575080516105a1565b60405162461bcd60e51b81526004016105989061177c565b60405180910390fd5b919050565b600082826040516020016105bb929190611545565b60405160208183030381529060405280519060200120905092915050565b60006105e3610fd6565b6034835110156105f6576000915061016d565b6000610608848263ffffffff61077516565b6001600160a01b03168252601401610626848263ffffffff61079816565b60208301525060019150915091565b61063d610fed565b6040805160a0808201835285825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916106a9565b610696610fed565b81526020019060019003908161068e5790505b508152600260208201526040019290925250919050565b6106c8610fed565b6040805160028082526060828101909352816020015b6106e6610fed565b8152602001906001900390816106de579050509050838160008151811061070957fe5b6020026020010181905250828160018151811061072257fe5b602002602001018190525061022081610b5d565b600086868686868660405160200161075396959493929190611580565b6040516020818303038152906040528051906020012090509695505050505050565b6000816014018351101561078857600080fd5b500160200151600160601b900490565b600081602001835110156107ab57600080fd5b50016020015190565b600390565b60008060008085519050848110806107d357506021858203105b806107f55750600060ff168686815181106107ea57fe5b016020015160f81c14155b1561080a57506000925083915082905061039a565b6001602186016108228888840163ffffffff61079816565b935093509350509250925092565b60006001820160608261084386846107b9565b919550935090508361085957506000925061039a565b60006020601f8301049050606081604051908082528060200260200182016040528015610890578160200160208202803883390190505b50905060005b828110156108ef5760018601955060006108b08a886107b9565b91995097509050876108cb57506000965061039a9350505050565b8060001b8383815181106108db57fe5b602090810291909101015250600101610896565b506001856109276000868560405160200161090a9190611539565b60408051601f19818403018152919052919063ffffffff610d1b16565b9550955095505050509250925092565b61093f610fed565b6040805160a0808201835284825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916109ab565b610998610fed565b8152602001906001900390816109905790505b50815260006020820152600160409091015292915050565b6109cb610fed565b6020820460006109d9610d9b565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b84811015610a7a578382600081518110610a1857fe5b602002602001018181525050610a476101c4610a42836020028b018c61079890919063ffffffff16565b610937565b82600181518110610a5457fe5b602002602001018181525050600283019250610a708284610dbc565b9350600101610a02565b506020860615610b00576000610a9c89601f198a8a010163ffffffff61079816565b90506020870660200360080281901b90508382600081518110610abb57fe5b602002602001018181525050610ad36101c482610937565b82600181518110610ae057fe5b602002602001018181525050600283019250610afc8284610dbc565b9350505b610b0c6101c487610937565b81600081518110610b1957fe5b6020026020010181815250508281600181518110610b3357fe5b602002602001018181525050600282019150610b4f8183610ddb565b9450505050505b9392505050565b610b65610fed565b610b6f8251610dfa565b610b8b5760405162461bcd60e51b81526004016105989061174c565b600160005b8351811015610bc257838181518110610ba557fe5b602002602001015160800151820191508080600101915050610b90565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b600081604051602001610c32919061156b565b604051602081830303815290604052805190602001209050919050565b60008315610c8c576001858484604051602001610c6f9493929190611669565b604051602081830303815290604052805190602001209050610220565b60018583604051602001610ca29392919061163d565b604051602081830303815290604052805190602001209050949350505050565b606081015160009060ff16600214610cec5760405162461bcd60e51b81526004016105989061175c565b815160808301516101e39190610e01565b6000610d07610fed565b610d1083610e19565b9050610b5681610cc2565b606081830184511015610d2d57600080fd5b606082158015610d4857604051915060208201604052610d92565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610d81578051835260209283019201610d69565b5050858452601f01601f1916604052505b50949350505050565b60408051600080825260208201909252610db6816001610dbc565b91505090565b6000610dc6610fed565b610dd08484610ddb565b905061022081610cc2565b610de3610fed565b6000610dee84610e65565b90506102208184610635565b6008101590565b6000600383836040516020016105bb93929190611606565b610e21610fed565b610e2a82610ebe565b610e465760405162461bcd60e51b81526004016105989061173c565b6060610e558360400151610ecd565b9050610b56818460800151610ddb565b6000600882511115610e895760405162461bcd60e51b81526004016105989061176c565b6000825183604051602001610e9f9291906115ea565b60408051808303601f1901815291905280516020909101209392505050565b60006101e38260600151610f79565b6060600882511115610ef15760405162461bcd60e51b81526004016105989061176c565b60608251604051908082528060200260200182016040528015610f1e578160200160208202803883390190505b50805190915060005b81811015610f70576000610f4d868381518110610f4057fe5b60200260200101516104c9565b905080848381518110610f5c57fe5b602090810291909101015250600101610f27565b50909392505050565b6000600c60ff83161080156101e3575050600360ff91909116101590565b604080516060810182526000808252602082018190529181019190915290565b6040805160608082018352600080835260208301529181019190915290565b604080518082019091526000808252602082015290565b6040518060a0016040528060008152602001611007611021565b815260606020820181905260006040830181905291015290565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b80356101e38161187d565b80356101e381611894565b600082601f83011261107657600080fd5b8135611089611084826117b3565b61178c565b915080825260208301602083018583830111156110a557600080fd5b6110b0838284611813565b50505092915050565b80356101e38161189d565b600080604083850312156110d757600080fd5b60006110e3858561105a565b92505060206110f48582860161105a565b9150509250929050565b6000806000806080858703121561111457600080fd5b6000611120878761105a565b94505060206111318782880161105a565b93505060406111428782880161105a565b92505060606111538782880161105a565b91505092959194509250565b60006020828403121561117157600080fd5b813567ffffffffffffffff81111561118857600080fd5b61022084828501611065565b600080604083850312156111a757600080fd5b823567ffffffffffffffff8111156111be57600080fd5b6110e385828601611065565b60008060008060008060c087890312156111e357600080fd5b60006111ef89896110b9565b965050602061120089828a0161104f565b955050604061121189828a0161105a565b945050606061122289828a0161105a565b935050608061123389828a0161105a565b92505060a061124489828a0161105a565b9150509295509295509295565b60008060008060008060c0878903121561126a57600080fd5b600061127689896110b9565b965050602061128789828a0161105a565b955050604061129889828a0161105a565b94505060606112a989828a0161104f565b93505060806112ba89828a0161105a565b92505060a087013567ffffffffffffffff8111156112d757600080fd5b61124489828a01611065565b60006112ef8383611379565b505060200190565b611300816117ee565b82525050565b611300611312826117ee565b61184b565b6000611322826117e1565b61132c81856105a1565b9350611337836117db565b8060005b8381101561136557815161134f88826112e3565b975061135a836117db565b92505060010161133b565b509495945050505050565b611300816117f9565b611300816117fe565b61130061138e826117fe565b6117fe565b600061139e826117e1565b6113a881856117e5565b93506113b881856020860161181f565b6113c181611867565b9093019392505050565b60006113d86012836117e5565b714d757374206265205475706c65207479706560701b815260200192915050565b6000611406601a836117e5565b7f5475706c65206d75737420686176652076616c69642073697a65000000000000815260200192915050565b600061143f6013836117e5565b7209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b815260200192915050565b600061146e6014836117e5565b73092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b815260200192915050565b600061149e6011836117e5565b70496e76616c6964207479706520636f646560781b815260200192915050565b805160608301906114cf84826112f7565b5060208201516114e260208501826112f7565b5060408201516114f56040850182611379565b50505050565b8051604083019061150c84826112f7565b5060208201516114f56020850182611379565b6113008161180d565b6113006115348261180d565b61185c565b6000610b568284611317565b60006115518285611382565b6020820191506115618284611382565b5060200192915050565b60006115778284611382565b50602001919050565b600061158c8289611528565b60018201915061159c8288611306565b6014820191506115ac8287611382565b6020820191506115bc8286611382565b6020820191506115cc8285611382565b6020820191506115dc8284611382565b506020019695505050505050565b60006115f68285611528565b6001820191506102208284611317565b60006116128286611528565b6001820191506116228285611382565b6020820191506116328284611382565b506020019392505050565b60006116498286611528565b6001820191506116598285611528565b6001820191506116328284611382565b60006116758287611528565b6001820191506116858286611528565b6001820191506116958285611382565b6020820191506116a58284611382565b50602001949350505050565b608081016116bf8285611370565b610b5660208301846114be565b606081016116da8285611370565b610b5660208301846114fb565b60a081016116f58288611370565b6117026020830187611379565b61170f604083018661151f565b61171c60608301856112f7565b81810360808301526101c98184611393565b602081016101e38284611379565b602080825281016101e3816113cb565b602080825281016101e3816113f9565b602080825281016101e381611432565b602080825281016101e381611461565b602080825281016101e381611491565b60405181810167ffffffffffffffff811182821017156117ab57600080fd5b604052919050565b600067ffffffffffffffff8211156117ca57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006101e382611801565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b60005b8381101561183a578181015183820152602001611822565b838111156114f55750506000910152565b60006101e38260006101e382611877565b60006101e382611871565b601f01601f191690565b60f81b90565b60601b90565b611886816117ee565b811461189157600080fd5b50565b611886816117fe565b6118868161180d56fea365627a7a72315820d6dcdd6c708d20d0dba7df8e4868ae0a13fd6ca167b161b1ae6fddbb193ad8eb6c6578706572696d656e74616cf564736f6c63430005110040"

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
var MessagesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a47083d0c76797776919ac0aceb4e8eed49fafad9fd418629f5cba2a7e242d1264736f6c63430005110032"

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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820e22c09b651e40f145dc6b2cc67151240f5603c865f346254908942b294c869d664736f6c63430005110032"

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
