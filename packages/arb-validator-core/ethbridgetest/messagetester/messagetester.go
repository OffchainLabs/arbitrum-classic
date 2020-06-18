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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158204a4406891b86a3dacfb6ea735b92a605972908aec5c506cc187bb6f01ef6a24064736f6c634300050f0032"

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
const MessageTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inboxHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"addMessageToInbox\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inboxTuplePreimage\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxTupleSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageTuplePreimage\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageTupleSize\",\"type\":\"uint256\"}],\"name\":\"addMessageToVMInboxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"erc20Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"erc20MessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"erc721Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"erc721MessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ethHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ethMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"transactionBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"prev\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prevSize\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"transactionMessageBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"transactionMessageBatchHashSingle\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"transactionMessageBatchSingleSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transactionMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MessageTesterFuncSigs maps the 4-byte function signature to its string representation.
var MessageTesterFuncSigs = map[string]string{
	"d504b196": "addMessageToInbox(bytes32,bytes32,uint256,uint256,uint256)",
	"353250a4": "addMessageToVMInboxHash(bytes32,uint256,uint256,uint256,uint256,bytes32,uint256)",
	"da78e176": "erc20Hash(address,address,address,uint256)",
	"655e768d": "erc20MessageHash(address,address,address,uint256)",
	"084f9984": "erc721Hash(address,address,address,uint256)",
	"e8f41e12": "erc721MessageHash(address,address,address,uint256)",
	"89d713b4": "ethHash(address,address,uint256)",
	"57464115": "ethMessageHash(address,address,uint256)",
	"93801dc3": "transactionBatchHash(bytes)",
	"a6c750b9": "transactionHash(address,address,address,uint256,uint256,bytes)",
	"a109e062": "transactionMessageBatchHash(bytes32,uint256,address,bytes,uint256,uint256)",
	"a4d84b9c": "transactionMessageBatchHashSingle(uint256,address,bytes)",
	"fedc217e": "transactionMessageBatchSingleSender(uint256,address,bytes32,bytes)",
	"39977357": "transactionMessageHash(address,address,address,uint256,uint256,bytes)",
}

// MessageTesterBin is the compiled bytecode used for deploying new contracts.
var MessageTesterBin = "0x608060405234801561001057600080fd5b50611dab806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a109e0621161008c578063d504b19611610066578063d504b19614610603578063da78e17614610638578063e8f41e1214610674578063fedc217e146106b0576100ea565b8063a109e062146103b4578063a4d84b9c14610479578063a6c750b914610532576100ea565b806357464115116100c85780635746411514610268578063655e768d1461029e57806389d713b4146102da57806393801dc314610310576100ea565b8063084f9984146100ef578063353250a41461013d578063399773571461017e575b600080fd5b61012b6004803603608081101561010557600080fd5b506001600160a01b0381358116916020810135821691604082013516906060013561078c565b60408051918252519081900360200190f35b61012b600480360360e081101561015357600080fd5b5080359060208101359060408101359060608101359060808101359060a08101359060c001356107a5565b61024f600480360360c081101561019457600080fd5b6001600160a01b038235811692602081013582169260408201359092169160608201359160808101359181019060c0810160a0820135600160201b8111156101db57600080fd5b8201836020820111156101ed57600080fd5b803590602001918460018302840111600160201b8311171561020e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107da945050505050565b6040805192835260208301919091528051918290030190f35b61012b6004803603606081101561027e57600080fd5b506001600160a01b03813581169160208101359091169060400135610825565b61012b600480360360808110156102b457600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610835565b61012b600480360360608110156102f057600080fd5b506001600160a01b03813581169160208101359091169060400135610846565b61012b6004803603602081101561032657600080fd5b810190602081018135600160201b81111561034057600080fd5b82018360208201111561035257600080fd5b803590602001918460018302840111600160201b8311171561037357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610853945050505050565b61012b600480360360c08110156103ca57600080fd5b8135916020810135916001600160a01b036040830135169190810190608081016060820135600160201b81111561040057600080fd5b82018360208201111561041257600080fd5b803590602001918460018302840111600160201b8311171561043357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610866565b61024f6004803603606081101561048f57600080fd5b8135916001600160a01b0360208201351691810190606081016040820135600160201b8111156104be57600080fd5b8201836020820111156104d057600080fd5b803590602001918460018302840111600160201b831117156104f157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610883945050505050565b61012b600480360360c081101561054857600080fd5b6001600160a01b038235811692602081013582169260408201359092169160608201359160808101359181019060c0810160a0820135600160201b81111561058f57600080fd5b8201836020820111156105a157600080fd5b803590602001918460018302840111600160201b831117156105c257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108b4945050505050565b61012b600480360360a081101561061957600080fd5b50803590602081013590604081013590606081013590608001356108cb565b61012b6004803603608081101561064e57600080fd5b506001600160a01b038135811691602081013582169160408201351690606001356108da565b61012b6004803603608081101561068a57600080fd5b506001600160a01b038135811691602081013582169160408201351690606001356108e8565b610770600480360360808110156106c657600080fd5b8135916001600160a01b036020820135169160408201359190810190608081016060820135600160201b8111156106fc57600080fd5b82018360208201111561070e57600080fd5b803590602001918460018302840111600160201b8311171561072f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108f9945050505050565b604080516001600160a01b039092168252519081900360200190f35b600061079a85858585610907565b90505b949350505050565b60006107ce6107c96107b78a8a610917565b8888886107c48989610917565b61099b565b610aca565b98975050505050505050565b6000806107e5611d1b565b60006108098a8a8a8a8a8a805190602001206108048c60008e51610bc0565b610d61565b9150915061081682610aca565b9a909950975050505050505050565b600061079d6107c9858585610e68565b600061079a6107c986868686610f06565b600061079d848484610f1c565b600061085e82610f72565b90505b919050565b6000610876878787878787610ffb565b90505b9695505050505050565b60008061088e611d1b565b600061089b8787876110aa565b915091506108a882610aca565b97909650945050505050565b6000610876878787878787805190602001206111b0565b6000610879868686868661121a565b600061079a85858585611264565b600061079a6107c986868686611274565b600061079a8585858561128a565b600061079a600386868686611340565b61091f611d1b565b6040805160a0810182528481528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610984565b610971611d1b565b8152602001906001900390816109695790505b508152600260208201526040019290925250919050565b6109a3611d1b565b60408051600480825260a08201909252606091816020015b6109c3611d1b565b8152602001906001900390816109bb5790505090506109e1866113ac565b816000815181106109ee57fe5b6020026020010181905250610a02856113ac565b81600181518110610a0f57fe5b6020026020010181905250610a23846113ac565b81600281518110610a3057fe5b60200260200101819052508281600381518110610a4957fe5b60209081029190910101526040805160028082526060828101909352816020015b610a72611d1b565b815260200190600190039081610a6a5790505090508781600081518110610a9557fe5b6020026020010181905250610aa982611431565b81600181518110610ab657fe5b60200260200101819052506107ce81611431565b6000600360090160ff16826060015160ff1610610b22576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610b40578151610b3990611520565b9050610861565b606082015160ff1660011415610b73576020808301518051604082015160608301519290930151610b3993919290611544565b606082015160ff1660021415610b8c57610b39826115ec565b600360ff16826060015160ff1610158015610bb057506060820151600c60ff909116105b15610bbe57610b3982611652565bfe5b610bc8611d1b565b602080830490601f8401046000610bdd611677565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b85811015610c7e578382600081518110610c1c57fe5b602002602001018181525050610c4b6107c9610c46836020028c018d61169890919063ffffffff16565b6113ac565b82600181518110610c5857fe5b602002602001018181525050600283019250610c7482846116b4565b9350600101610c06565b5083851015610d05576000610c9f8a601f198b8b010163ffffffff61169816565b905085602002886020030360080281901b90508382600081518110610cc057fe5b602002602001018181525050610cd86107c9826113ac565b82600181518110610ce557fe5b602002602001018181525050600283019250610d0182846116b4565b9350505b610d116107c9886113ac565b81600081518110610d1e57fe5b6020026020010181815250508281600181518110610d3857fe5b602002602001018181525050600282019150610d5481836116d3565b9998505050505050505050565b610d69611d1b565b60408051600480825260a0820190925260009160609190816020015b610d8d611d1b565b815260200190600190039081610d85579050509050610db4896001600160a01b03166113ac565b81600081518110610dc157fe5b6020026020010181905250610dd5876113ac565b81600181518110610de257fe5b6020026020010181905250610df6866113ac565b81600281518110610e0357fe5b60200260200101819052508381600381518110610e1c57fe5b6020026020010181905250610e2f611d1b565b610e4360008a610e3e85611431565b6116f2565b90506000610e558c8c8c8c8c8c6111b0565b919c919b50909950505050505050505050565b610e70611d1b565b6040805160028082526060828101909352816020015b610e8e611d1b565b815260200190600190039081610e86579050509050610eb5856001600160a01b03166113ac565b81600081518110610ec257fe5b6020026020010181905250610ed6836113ac565b81600181518110610ee357fe5b6020908102919091010152610efd600185610e3e84611431565b95945050505050565b610f0e611d1b565b61079a60028686868661179c565b60408051600160f81b6020808301919091526001600160601b0319606087811b8216602185015286901b166035830152604980830185905283518084039091018152606990920190925280519101209392505050565b6000600682604051602001808360ff1660ff1660f81b815260010182805190602001908083835b60208310610fb85780518252601f199092019160209182019101610f99565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052805190602001209050919050565b6000611005611d1b565b61100f8888610917565b855190915060005b81609782011015611093576000611034888363ffffffff61186a16565b9050828161ffff166097840101111561105b5761105084610aca565b945050505050610879565b611063611d1b565b6000611070848c8c6110aa565b9092509050611082868a8a848661099b565b9550505061ffff1601609701611017565b61109c83610aca565b9a9950505050505050505050565b6110b2611d1b565b6000806110d684609788016110cd828a63ffffffff61186a16565b61ffff16611886565b905060006110e68787848861128a565b90506001600160a01b038116611131576040805162461bcd60e51b815260206004820152600b60248201526a696e76616c69642073696760a81b604482015290519081900360640190fd5b611139611d1b565b61115a8660978a01611151828c63ffffffff61186a16565b61ffff16610bc0565b90506111a1876111738860028c0163ffffffff61188f16565b846111878a60168e0163ffffffff61169816565b61119a8b60368f0163ffffffff61169816565b8887610d61565b94509450505050935093915050565b6040805160006020808301919091526001600160601b03196060998a1b8116602184015297891b881660358301529590971b9095166049870152605d860192909252607d850152609d808501919091528251808503909101815260bd909301909152815191012090565b6000610879868686868660405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001206118b2565b600061079a600286868686611340565b61127c611d1b565b61079a60038686868661179c565b600061079a846112a3846002890163ffffffff61188f16565b6112b68560168a0163ffffffff61169816565b6112c98660368b0163ffffffff61169816565b8760405160200180866001600160a01b03166001600160a01b031660601b8152601401856001600160a01b03166001600160a01b031660601b8152601401848152602001838152602001828152602001955050505050506040516020818303038152906040528051906020012083605688016118de565b604080516001600160f81b031960f888901b166020808301919091526001600160601b0319606088811b8216602185015287811b8216603585015286901b166049830152605d80830185905283518084039091018152607d909201909252805191012095945050505050565b6113b4611d1b565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191611419565b611406611d1b565b8152602001906001900390816113fe5790505b50815260006020820152600160409091015292915050565b611439611d1b565b6114438251611a11565b611494576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156114cb578381815181106114ae57fe5b602002602001015160800151820191508080600101915050611499565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561159e575060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602282018590526042808301859052835180840390910181526062909201909252805191012061079d565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff16600214611641576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161085e9190611a18565b600061165c611d1b565b61166583611a52565b9050611670816115ec565b9392505050565b604080516000808252602082019092526116928160016116b4565b91505090565b600081602001835110156116ab57600080fd5b50016020015190565b60006116be611d1b565b6116c884846116d3565b905061079d816115ec565b6116db611d1b565b60006116e684611ac8565b905061079d8184610917565b6116fa611d1b565b60408051600380825260808201909252606091816020015b61171a611d1b565b815260200190600190039081611712579050509050611738856113ac565b8160008151811061174557fe5b6020026020010181905250611762846001600160a01b03166113ac565b8160018151811061176f57fe5b6020026020010181905250828160028151811061178857fe5b6020026020010181905250610efd81611431565b6117a4611d1b565b60408051600380825260808201909252606091816020015b6117c4611d1b565b8152602001906001900390816117bc5790505090506117eb846001600160a01b03166113ac565b816000815181106117f857fe5b6020026020010181905250611815866001600160a01b03166113ac565b8160018151811061182257fe5b6020026020010181905250611836836113ac565b8160028151811061184357fe5b602002602001018190525061185f8760ff1686610e3e84611431565b979650505050505050565b6000816002018351101561187d57600080fd5b50016002015190565b91016020012090565b600081601401835110156118a257600080fd5b500160200151600160601b900490565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081896040516020018083805190602001908083835b602083106119545780518252601f199092019160209182019101611935565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190528251920191909120925061199a915089905088611b88565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa1580156119f9573d6000803e3d6000fd5b5050604051601f1901519a9950505050505050505050565b6008101590565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b611a5a611d1b565b611a6382611c16565b611aa9576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060611ab88360400151611c25565b90506116708184608001516116d3565b6000600882511115611b18576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611b5c578181015183820152602001611b44565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b818101602081810151604083015160609093015160001a9290918401601b841015611bb457601b840193505b8360ff16601b1480611bc957508360ff16601c145b611c0e576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b600061085e8260600151611cfd565b6060600882511115611c75576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611ca2578160200160208202803883390190505b50805190915060005b81811015611cf4576000611cd1868381518110611cc457fe5b6020026020010151610aca565b905080848381518110611ce057fe5b602090810291909101015250600101611cab565b50909392505050565b6000600c60ff831610801561085e575050600360ff91909116101590565b6040518060a0016040528060008152602001611d35611d4f565b815260606020820181905260006040830181905291015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a723158200a69d6e0f3680b32f178c1746c04b1bd8f85b34da28485b3cdc885242623919b64736f6c634300050f0032"

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

// AddMessageToInbox is a free data retrieval call binding the contract method 0xd504b196.
//
// Solidity: function addMessageToInbox(bytes32 inboxHash, bytes32 messageHash, uint256 blockNumber, uint256 timestamp, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) AddMessageToInbox(opts *bind.CallOpts, inboxHash [32]byte, messageHash [32]byte, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "addMessageToInbox", inboxHash, messageHash, blockNumber, timestamp, messageNum)
	return *ret0, err
}

// AddMessageToInbox is a free data retrieval call binding the contract method 0xd504b196.
//
// Solidity: function addMessageToInbox(bytes32 inboxHash, bytes32 messageHash, uint256 blockNumber, uint256 timestamp, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) AddMessageToInbox(inboxHash [32]byte, messageHash [32]byte, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.AddMessageToInbox(&_MessageTester.CallOpts, inboxHash, messageHash, blockNumber, timestamp, messageNum)
}

// AddMessageToInbox is a free data retrieval call binding the contract method 0xd504b196.
//
// Solidity: function addMessageToInbox(bytes32 inboxHash, bytes32 messageHash, uint256 blockNumber, uint256 timestamp, uint256 messageNum) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) AddMessageToInbox(inboxHash [32]byte, messageHash [32]byte, blockNumber *big.Int, timestamp *big.Int, messageNum *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.AddMessageToInbox(&_MessageTester.CallOpts, inboxHash, messageHash, blockNumber, timestamp, messageNum)
}

// AddMessageToVMInboxHash is a free data retrieval call binding the contract method 0x353250a4.
//
// Solidity: function addMessageToVMInboxHash(bytes32 inboxTuplePreimage, uint256 inboxTupleSize, uint256 blockNumber, uint256 timestamp, uint256 txId, bytes32 messageTuplePreimage, uint256 messageTupleSize) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) AddMessageToVMInboxHash(opts *bind.CallOpts, inboxTuplePreimage [32]byte, inboxTupleSize *big.Int, blockNumber *big.Int, timestamp *big.Int, txId *big.Int, messageTuplePreimage [32]byte, messageTupleSize *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "addMessageToVMInboxHash", inboxTuplePreimage, inboxTupleSize, blockNumber, timestamp, txId, messageTuplePreimage, messageTupleSize)
	return *ret0, err
}

// AddMessageToVMInboxHash is a free data retrieval call binding the contract method 0x353250a4.
//
// Solidity: function addMessageToVMInboxHash(bytes32 inboxTuplePreimage, uint256 inboxTupleSize, uint256 blockNumber, uint256 timestamp, uint256 txId, bytes32 messageTuplePreimage, uint256 messageTupleSize) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) AddMessageToVMInboxHash(inboxTuplePreimage [32]byte, inboxTupleSize *big.Int, blockNumber *big.Int, timestamp *big.Int, txId *big.Int, messageTuplePreimage [32]byte, messageTupleSize *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.AddMessageToVMInboxHash(&_MessageTester.CallOpts, inboxTuplePreimage, inboxTupleSize, blockNumber, timestamp, txId, messageTuplePreimage, messageTupleSize)
}

// AddMessageToVMInboxHash is a free data retrieval call binding the contract method 0x353250a4.
//
// Solidity: function addMessageToVMInboxHash(bytes32 inboxTuplePreimage, uint256 inboxTupleSize, uint256 blockNumber, uint256 timestamp, uint256 txId, bytes32 messageTuplePreimage, uint256 messageTupleSize) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) AddMessageToVMInboxHash(inboxTuplePreimage [32]byte, inboxTupleSize *big.Int, blockNumber *big.Int, timestamp *big.Int, txId *big.Int, messageTuplePreimage [32]byte, messageTupleSize *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.AddMessageToVMInboxHash(&_MessageTester.CallOpts, inboxTuplePreimage, inboxTupleSize, blockNumber, timestamp, txId, messageTuplePreimage, messageTupleSize)
}

// Erc20Hash is a free data retrieval call binding the contract method 0xda78e176.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc20Hash(opts *bind.CallOpts, to common.Address, from common.Address, erc20 common.Address, value *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc20Hash", to, from, erc20, value)
	return *ret0, err
}

// Erc20Hash is a free data retrieval call binding the contract method 0xda78e176.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc20Hash(to common.Address, from common.Address, erc20 common.Address, value *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20Hash(&_MessageTester.CallOpts, to, from, erc20, value)
}

// Erc20Hash is a free data retrieval call binding the contract method 0xda78e176.
//
// Solidity: function erc20Hash(address to, address from, address erc20, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc20Hash(to common.Address, from common.Address, erc20 common.Address, value *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20Hash(&_MessageTester.CallOpts, to, from, erc20, value)
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0x655e768d.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc20MessageHash(opts *bind.CallOpts, to common.Address, from common.Address, erc20 common.Address, value *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc20MessageHash", to, from, erc20, value)
	return *ret0, err
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0x655e768d.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc20MessageHash(to common.Address, from common.Address, erc20 common.Address, value *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20MessageHash(&_MessageTester.CallOpts, to, from, erc20, value)
}

// Erc20MessageHash is a free data retrieval call binding the contract method 0x655e768d.
//
// Solidity: function erc20MessageHash(address to, address from, address erc20, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc20MessageHash(to common.Address, from common.Address, erc20 common.Address, value *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc20MessageHash(&_MessageTester.CallOpts, to, from, erc20, value)
}

// Erc721Hash is a free data retrieval call binding the contract method 0x084f9984.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc721Hash(opts *bind.CallOpts, to common.Address, from common.Address, erc721 common.Address, id *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc721Hash", to, from, erc721, id)
	return *ret0, err
}

// Erc721Hash is a free data retrieval call binding the contract method 0x084f9984.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc721Hash(to common.Address, from common.Address, erc721 common.Address, id *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721Hash(&_MessageTester.CallOpts, to, from, erc721, id)
}

// Erc721Hash is a free data retrieval call binding the contract method 0x084f9984.
//
// Solidity: function erc721Hash(address to, address from, address erc721, uint256 id) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc721Hash(to common.Address, from common.Address, erc721 common.Address, id *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721Hash(&_MessageTester.CallOpts, to, from, erc721, id)
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xe8f41e12.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) Erc721MessageHash(opts *bind.CallOpts, to common.Address, from common.Address, erc721 common.Address, id *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "erc721MessageHash", to, from, erc721, id)
	return *ret0, err
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xe8f41e12.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) Erc721MessageHash(to common.Address, from common.Address, erc721 common.Address, id *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721MessageHash(&_MessageTester.CallOpts, to, from, erc721, id)
}

// Erc721MessageHash is a free data retrieval call binding the contract method 0xe8f41e12.
//
// Solidity: function erc721MessageHash(address to, address from, address erc721, uint256 id) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) Erc721MessageHash(to common.Address, from common.Address, erc721 common.Address, id *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.Erc721MessageHash(&_MessageTester.CallOpts, to, from, erc721, id)
}

// EthHash is a free data retrieval call binding the contract method 0x89d713b4.
//
// Solidity: function ethHash(address to, address from, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) EthHash(opts *bind.CallOpts, to common.Address, from common.Address, value *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "ethHash", to, from, value)
	return *ret0, err
}

// EthHash is a free data retrieval call binding the contract method 0x89d713b4.
//
// Solidity: function ethHash(address to, address from, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) EthHash(to common.Address, from common.Address, value *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthHash(&_MessageTester.CallOpts, to, from, value)
}

// EthHash is a free data retrieval call binding the contract method 0x89d713b4.
//
// Solidity: function ethHash(address to, address from, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) EthHash(to common.Address, from common.Address, value *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthHash(&_MessageTester.CallOpts, to, from, value)
}

// EthMessageHash is a free data retrieval call binding the contract method 0x57464115.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) EthMessageHash(opts *bind.CallOpts, to common.Address, from common.Address, value *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "ethMessageHash", to, from, value)
	return *ret0, err
}

// EthMessageHash is a free data retrieval call binding the contract method 0x57464115.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) EthMessageHash(to common.Address, from common.Address, value *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthMessageHash(&_MessageTester.CallOpts, to, from, value)
}

// EthMessageHash is a free data retrieval call binding the contract method 0x57464115.
//
// Solidity: function ethMessageHash(address to, address from, uint256 value) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) EthMessageHash(to common.Address, from common.Address, value *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.EthMessageHash(&_MessageTester.CallOpts, to, from, value)
}

// TransactionBatchHash is a free data retrieval call binding the contract method 0x93801dc3.
//
// Solidity: function transactionBatchHash(bytes transactions) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionBatchHash(opts *bind.CallOpts, transactions []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionBatchHash", transactions)
	return *ret0, err
}

// TransactionBatchHash is a free data retrieval call binding the contract method 0x93801dc3.
//
// Solidity: function transactionBatchHash(bytes transactions) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionBatchHash(transactions []byte) ([32]byte, error) {
	return _MessageTester.Contract.TransactionBatchHash(&_MessageTester.CallOpts, transactions)
}

// TransactionBatchHash is a free data retrieval call binding the contract method 0x93801dc3.
//
// Solidity: function transactionBatchHash(bytes transactions) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionBatchHash(transactions []byte) ([32]byte, error) {
	return _MessageTester.Contract.TransactionBatchHash(&_MessageTester.CallOpts, transactions)
}

// TransactionHash is a free data retrieval call binding the contract method 0xa6c750b9.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionHash(opts *bind.CallOpts, chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionHash", chain, to, from, seqNumber, value, data)
	return *ret0, err
}

// TransactionHash is a free data retrieval call binding the contract method 0xa6c750b9.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _MessageTester.Contract.TransactionHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data)
}

// TransactionHash is a free data retrieval call binding the contract method 0xa6c750b9.
//
// Solidity: function transactionHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _MessageTester.Contract.TransactionHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data)
}

// TransactionMessageBatchHash is a free data retrieval call binding the contract method 0xa109e062.
//
// Solidity: function transactionMessageBatchHash(bytes32 prev, uint256 prevSize, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) constant returns(bytes32)
func (_MessageTester *MessageTesterCaller) TransactionMessageBatchHash(opts *bind.CallOpts, prev [32]byte, prevSize *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionMessageBatchHash", prev, prevSize, chain, transactions, blockNum, blockTimestamp)
	return *ret0, err
}

// TransactionMessageBatchHash is a free data retrieval call binding the contract method 0xa109e062.
//
// Solidity: function transactionMessageBatchHash(bytes32 prev, uint256 prevSize, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) constant returns(bytes32)
func (_MessageTester *MessageTesterSession) TransactionMessageBatchHash(prev [32]byte, prevSize *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageBatchHash(&_MessageTester.CallOpts, prev, prevSize, chain, transactions, blockNum, blockTimestamp)
}

// TransactionMessageBatchHash is a free data retrieval call binding the contract method 0xa109e062.
//
// Solidity: function transactionMessageBatchHash(bytes32 prev, uint256 prevSize, address chain, bytes transactions, uint256 blockNum, uint256 blockTimestamp) constant returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageBatchHash(prev [32]byte, prevSize *big.Int, chain common.Address, transactions []byte, blockNum *big.Int, blockTimestamp *big.Int) ([32]byte, error) {
	return _MessageTester.Contract.TransactionMessageBatchHash(&_MessageTester.CallOpts, prev, prevSize, chain, transactions, blockNum, blockTimestamp)
}

// TransactionMessageBatchHashSingle is a free data retrieval call binding the contract method 0xa4d84b9c.
//
// Solidity: function transactionMessageBatchHashSingle(uint256 start, address chain, bytes transactions) constant returns(bytes32, bytes32)
func (_MessageTester *MessageTesterCaller) TransactionMessageBatchHashSingle(opts *bind.CallOpts, start *big.Int, chain common.Address, transactions []byte) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MessageTester.contract.Call(opts, out, "transactionMessageBatchHashSingle", start, chain, transactions)
	return *ret0, *ret1, err
}

// TransactionMessageBatchHashSingle is a free data retrieval call binding the contract method 0xa4d84b9c.
//
// Solidity: function transactionMessageBatchHashSingle(uint256 start, address chain, bytes transactions) constant returns(bytes32, bytes32)
func (_MessageTester *MessageTesterSession) TransactionMessageBatchHashSingle(start *big.Int, chain common.Address, transactions []byte) ([32]byte, [32]byte, error) {
	return _MessageTester.Contract.TransactionMessageBatchHashSingle(&_MessageTester.CallOpts, start, chain, transactions)
}

// TransactionMessageBatchHashSingle is a free data retrieval call binding the contract method 0xa4d84b9c.
//
// Solidity: function transactionMessageBatchHashSingle(uint256 start, address chain, bytes transactions) constant returns(bytes32, bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageBatchHashSingle(start *big.Int, chain common.Address, transactions []byte) ([32]byte, [32]byte, error) {
	return _MessageTester.Contract.TransactionMessageBatchHashSingle(&_MessageTester.CallOpts, start, chain, transactions)
}

// TransactionMessageBatchSingleSender is a free data retrieval call binding the contract method 0xfedc217e.
//
// Solidity: function transactionMessageBatchSingleSender(uint256 start, address chain, bytes32 dataHash, bytes transactions) constant returns(address)
func (_MessageTester *MessageTesterCaller) TransactionMessageBatchSingleSender(opts *bind.CallOpts, start *big.Int, chain common.Address, dataHash [32]byte, transactions []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MessageTester.contract.Call(opts, out, "transactionMessageBatchSingleSender", start, chain, dataHash, transactions)
	return *ret0, err
}

// TransactionMessageBatchSingleSender is a free data retrieval call binding the contract method 0xfedc217e.
//
// Solidity: function transactionMessageBatchSingleSender(uint256 start, address chain, bytes32 dataHash, bytes transactions) constant returns(address)
func (_MessageTester *MessageTesterSession) TransactionMessageBatchSingleSender(start *big.Int, chain common.Address, dataHash [32]byte, transactions []byte) (common.Address, error) {
	return _MessageTester.Contract.TransactionMessageBatchSingleSender(&_MessageTester.CallOpts, start, chain, dataHash, transactions)
}

// TransactionMessageBatchSingleSender is a free data retrieval call binding the contract method 0xfedc217e.
//
// Solidity: function transactionMessageBatchSingleSender(uint256 start, address chain, bytes32 dataHash, bytes transactions) constant returns(address)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageBatchSingleSender(start *big.Int, chain common.Address, dataHash [32]byte, transactions []byte) (common.Address, error) {
	return _MessageTester.Contract.TransactionMessageBatchSingleSender(&_MessageTester.CallOpts, start, chain, dataHash, transactions)
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x39977357.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data) constant returns(bytes32, bytes32)
func (_MessageTester *MessageTesterCaller) TransactionMessageHash(opts *bind.CallOpts, chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MessageTester.contract.Call(opts, out, "transactionMessageHash", chain, to, from, seqNumber, value, data)
	return *ret0, *ret1, err
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x39977357.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data) constant returns(bytes32, bytes32)
func (_MessageTester *MessageTesterSession) TransactionMessageHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte) ([32]byte, [32]byte, error) {
	return _MessageTester.Contract.TransactionMessageHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data)
}

// TransactionMessageHash is a free data retrieval call binding the contract method 0x39977357.
//
// Solidity: function transactionMessageHash(address chain, address to, address from, uint256 seqNumber, uint256 value, bytes data) constant returns(bytes32, bytes32)
func (_MessageTester *MessageTesterCallerSession) TransactionMessageHash(chain common.Address, to common.Address, from common.Address, seqNumber *big.Int, value *big.Int, data []byte) ([32]byte, [32]byte, error) {
	return _MessageTester.Contract.TransactionMessageHash(&_MessageTester.CallOpts, chain, to, from, seqNumber, value, data)
}

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582073b40bfb370786f88d08c51ae363931f6fc5e092166bc0110f6c1808495aefff64736f6c634300050f0032"

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

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820f983a223df2637fb0a1be5323c1a6145d6c2c7cad6cf14d5699c7e29a6a8803664736f6c634300050f0032"

// DeployProtocol deploys a new Ethereum contract, binding an instance of Protocol to it.
func DeployProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Protocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// Protocol is an auto generated Go binding around an Ethereum contract.
type Protocol struct {
	ProtocolCaller     // Read-only binding to the contract
	ProtocolTransactor // Write-only binding to the contract
	ProtocolFilterer   // Log filterer for contract events
}

// ProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolSession struct {
	Contract     *Protocol         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolCallerSession struct {
	Contract *ProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolTransactorSession struct {
	Contract     *ProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRaw struct {
	Contract *Protocol // Generic contract binding to access the raw methods on
}

// ProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolCallerRaw struct {
	Contract *ProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolTransactorRaw struct {
	Contract *ProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocol creates a new instance of Protocol, bound to a specific deployed contract.
func NewProtocol(address common.Address, backend bind.ContractBackend) (*Protocol, error) {
	contract, err := bindProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// NewProtocolCaller creates a new read-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolCaller(address common.Address, caller bind.ContractCaller) (*ProtocolCaller, error) {
	contract, err := bindProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolCaller{contract: contract}, nil
}

// NewProtocolTransactor creates a new write-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolTransactor, error) {
	contract, err := bindProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTransactor{contract: contract}, nil
}

// NewProtocolFilterer creates a new log filterer instance of Protocol, bound to a specific deployed contract.
func NewProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolFilterer, error) {
	contract, err := bindProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolFilterer{contract: contract}, nil
}

// bindProtocol binds a generic wrapper to an already deployed contract.
func bindProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.ProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transact(opts, method, params...)
}

// SigUtilsABI is the input ABI used to generate the binding from.
const SigUtilsABI = "[]"

// SigUtilsBin is the compiled bytecode used for deploying new contracts.
var SigUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158204a917da2883ebd014f34b8368e961f23ca9c2f5c48235579f75d635bc878023064736f6c634300050f0032"

// DeploySigUtils deploys a new Ethereum contract, binding an instance of SigUtils to it.
func DeploySigUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// SigUtils is an auto generated Go binding around an Ethereum contract.
type SigUtils struct {
	SigUtilsCaller     // Read-only binding to the contract
	SigUtilsTransactor // Write-only binding to the contract
	SigUtilsFilterer   // Log filterer for contract events
}

// SigUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsSession struct {
	Contract     *SigUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsCallerSession struct {
	Contract *SigUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTransactorSession struct {
	Contract     *SigUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsRaw struct {
	Contract *SigUtils // Generic contract binding to access the raw methods on
}

// SigUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsCallerRaw struct {
	Contract *SigUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTransactorRaw struct {
	Contract *SigUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtils creates a new instance of SigUtils, bound to a specific deployed contract.
func NewSigUtils(address common.Address, backend bind.ContractBackend) (*SigUtils, error) {
	contract, err := bindSigUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// NewSigUtilsCaller creates a new read-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsCaller, error) {
	contract, err := bindSigUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsCaller{contract: contract}, nil
}

// NewSigUtilsTransactor creates a new write-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTransactor, error) {
	contract, err := bindSigUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTransactor{contract: contract}, nil
}

// NewSigUtilsFilterer creates a new log filterer instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsFilterer, error) {
	contract, err := bindSigUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsFilterer{contract: contract}, nil
}

// bindSigUtils binds a generic wrapper to an already deployed contract.
func bindSigUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.SigUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582044b94c18e2bebb1085f90b90eb6ee2dc0fa9bec619bad6c3a7d31981dbb2f15264736f6c634300050f0032"

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
