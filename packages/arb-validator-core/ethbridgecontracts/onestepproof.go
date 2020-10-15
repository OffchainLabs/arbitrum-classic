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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IOneStepProof2ABI is the input ABI used to generate the binding from.
const IOneStepProof2ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inboxAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagesAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logsAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IOneStepProof2FuncSigs maps the 4-byte function signature to its string representation.
var IOneStepProof2FuncSigs = map[string]string{
	"1041c884": "executeStep(bytes32,bytes32,bytes32,bytes,bytes)",
}

// IOneStepProof2 is an auto generated Go binding around an Ethereum contract.
type IOneStepProof2 struct {
	IOneStepProof2Caller     // Read-only binding to the contract
	IOneStepProof2Transactor // Write-only binding to the contract
	IOneStepProof2Filterer   // Log filterer for contract events
}

// IOneStepProof2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProof2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProof2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProof2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProof2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProof2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProof2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProof2Session struct {
	Contract     *IOneStepProof2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOneStepProof2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProof2CallerSession struct {
	Contract *IOneStepProof2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IOneStepProof2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProof2TransactorSession struct {
	Contract     *IOneStepProof2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOneStepProof2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProof2Raw struct {
	Contract *IOneStepProof2 // Generic contract binding to access the raw methods on
}

// IOneStepProof2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProof2CallerRaw struct {
	Contract *IOneStepProof2Caller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProof2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProof2TransactorRaw struct {
	Contract *IOneStepProof2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProof2 creates a new instance of IOneStepProof2, bound to a specific deployed contract.
func NewIOneStepProof2(address common.Address, backend bind.ContractBackend) (*IOneStepProof2, error) {
	contract, err := bindIOneStepProof2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProof2{IOneStepProof2Caller: IOneStepProof2Caller{contract: contract}, IOneStepProof2Transactor: IOneStepProof2Transactor{contract: contract}, IOneStepProof2Filterer: IOneStepProof2Filterer{contract: contract}}, nil
}

// NewIOneStepProof2Caller creates a new read-only instance of IOneStepProof2, bound to a specific deployed contract.
func NewIOneStepProof2Caller(address common.Address, caller bind.ContractCaller) (*IOneStepProof2Caller, error) {
	contract, err := bindIOneStepProof2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProof2Caller{contract: contract}, nil
}

// NewIOneStepProof2Transactor creates a new write-only instance of IOneStepProof2, bound to a specific deployed contract.
func NewIOneStepProof2Transactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProof2Transactor, error) {
	contract, err := bindIOneStepProof2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProof2Transactor{contract: contract}, nil
}

// NewIOneStepProof2Filterer creates a new log filterer instance of IOneStepProof2, bound to a specific deployed contract.
func NewIOneStepProof2Filterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProof2Filterer, error) {
	contract, err := bindIOneStepProof2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProof2Filterer{contract: contract}, nil
}

// bindIOneStepProof2 binds a generic wrapper to an already deployed contract.
func bindIOneStepProof2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOneStepProof2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseIOneStepProof2ABI parses the ABI
func ParseIOneStepProof2ABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(IOneStepProof2ABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProof2 *IOneStepProof2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IOneStepProof2.Contract.IOneStepProof2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProof2 *IOneStepProof2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProof2.Contract.IOneStepProof2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProof2 *IOneStepProof2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProof2.Contract.IOneStepProof2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProof2 *IOneStepProof2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IOneStepProof2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProof2 *IOneStepProof2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProof2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProof2 *IOneStepProof2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProof2.Contract.contract.Transact(opts, method, params...)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) constant returns(uint64 gas, bytes32[5] fields)
func (_IOneStepProof2 *IOneStepProof2Caller) ExecuteStep(opts *bind.CallOpts, inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	ret := new(struct {
		Gas    uint64
		Fields [5][32]byte
	})
	out := ret
	err := _IOneStepProof2.contract.Call(opts, out, "executeStep", inboxAcc, messagesAcc, logsAcc, proof, bproof)
	return *ret, err
}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) constant returns(uint64 gas, bytes32[5] fields)
func (_IOneStepProof2 *IOneStepProof2Session) ExecuteStep(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _IOneStepProof2.Contract.ExecuteStep(&_IOneStepProof2.CallOpts, inboxAcc, messagesAcc, logsAcc, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) constant returns(uint64 gas, bytes32[5] fields)
func (_IOneStepProof2 *IOneStepProof2CallerSession) ExecuteStep(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _IOneStepProof2.Contract.ExecuteStep(&_IOneStepProof2.CallOpts, inboxAcc, messagesAcc, logsAcc, proof, bproof)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_IOneStepProof2 *IOneStepProof2Filterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _IOneStepProof2.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// MerkleUtilABI is the input ABI used to generate the binding from.
const MerkleUtilABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"buf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"bufferOp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"buf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof2\",\"type\":\"bytes32[]\"}],\"name\":\"getOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"buf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"nproof1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof2\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"nproof2\",\"type\":\"bytes32[]\"}],\"name\":\"setOp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleUtilFuncSigs maps the 4-byte function signature to its string representation.
var MerkleUtilFuncSigs = map[string]string{
	"585ec5ad": "bufferOp(uint8,bytes32,uint256,uint256,bytes)",
	"1424c514": "getOp(uint8,bytes32,uint256,bytes32[],bytes32[])",
	"c58ed997": "setOp(uint8,bytes32,uint256,uint256,bytes32[],bytes32[],bytes32[],bytes32[])",
}

// MerkleUtilBin is the compiled bytecode used for deploying new contracts.
var MerkleUtilBin = "0x611254610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80631424c51414610050578063585ec5ad1461019b578063c58ed9971461025a575b600080fd5b610189600480360360a081101561006657600080fd5b60ff8235169160208101359160408201359190810190608081016060820135600160201b81111561009657600080fd5b8201836020820111156100a857600080fd5b803590602001918460208302840111600160201b831117156100c957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561011857600080fd5b82018360208201111561012a57600080fd5b803590602001918460208302840111600160201b8311171561014b57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061049d945050505050565b60408051918252519081900360200190f35b610189600480360360a08110156101b157600080fd5b60ff8235169160208101359160408201359160608101359181019060a081016080820135600160201b8111156101e657600080fd5b8201836020820111156101f857600080fd5b803590602001918460018302840111600160201b8311171561021957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610506945050505050565b610189600480360361010081101561027157600080fd5b60ff8235169160208101359160408201359160608101359181019060a081016080820135600160201b8111156102a657600080fd5b8201836020820111156102b857600080fd5b803590602001918460208302840111600160201b831117156102d957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561032857600080fd5b82018360208201111561033a57600080fd5b803590602001918460208302840111600160201b8311171561035b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103aa57600080fd5b8201836020820111156103bc57600080fd5b803590602001918460208302840111600160201b831117156103dd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561042c57600080fd5b82018360208201111561043e57600080fd5b803590602001918460208302840111600160201b8311171561045f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506106cb945050505050565b60008560ff1660a114156104cd576104c66104bc866020870486610741565b6020865b06610826565b90506104fd565b8560ff1660a214156104e5576104c685858585610833565b8560ff1660a314156104fd576104c685858585610992565b95945050505050565b60006060610547838460008151811061051b57fe5b602001015160f81c60f81b8560018151811061053357fe5b01602001516001600160f81b031916610ac5565b90506060610574848560028151811061055c57fe5b602001015160f81c60f81b8660038151811061053357fe5b905060606105a1858660048151811061058957fe5b602001015160f81c60f81b8760058151811061053357fe5b905060606105ce86876006815181106105b657fe5b602001015160f81c60f81b8860078151811061053357fe5b90508960ff1660a11415610602576105f66105ee8a60208b5b0487610741565b60208a6104c0565b94506104fd9350505050565b8960ff1660a2141561061a576105f689898685610833565b8960ff1660a31415610632576105f689898685610992565b8960ff1660a4141561067d57600061064c8a60208b6105e7565b9050600061065e8260208c068b610b48565b90506106708b60208c04838989610b87565b96505050505050506104fd565b8960ff1660a514156106a35761069889898987878787610be1565b9450505050506104fd565b8960ff1660a614156106be5761069889898987878787610d26565b5050505095945050505050565b60008860ff1660a414156107175760006106e98960208a0488610741565b905060006106fb8260208b068a610b48565b905061070e8a60208b5b04838a8a610b87565b92505050610735565b8860ff1660a514156107355761073288888888888888610be1565b90505b98975050505050505050565b600081516000141561076057831561075857600080fd5b50600061081f565b600061077f8360008151811061077257fe5b6020026020010151610e0d565b905060015b83518110156107e95784600116600114156107bd576107b6828583815181106107a957fe5b6020026020010151610e33565b91506107dd565b6107da8482815181106107cc57fe5b602002602001015183610e33565b91505b600194851c9401610784565b508481146107f657600080fd5b831561080657506000905061081f565b8260008151811061081357fe5b60200260200101519150505b9392505050565b601f036008021c60ff1690565b604080516008808252818301909252600091606091906020820181803883390190505090506000610866876020886105e7565b9050602080870660080110610935576000610889886020895b0460010187610741565b905060005b6018601f8916600803018110156108dd576108af838260208b5b0601610826565b60f81b8482815181106108be57fe5b60200101906001600160f81b031916908160001a90535060010161088e565b506018601f8816600803015b600881101561092e576109008260208a84016104c0565b60f81b84828151811061090f57fe5b60200101906001600160f81b031916908160001a9053506001016108e9565b505061097e565b60005b600881101561097c5761094e828260208a6108a8565b60f81b83828151811061095d57fe5b60200101906001600160f81b031916908160001a905350600101610938565b505b61098782610e5f565b979650505050505050565b6040805160208082528183019092526000916060919060208201818038833901905050905060006109c5876020886105e7565b9050602080870660200110610a7e5760006109e28860208961087f565b905060005b601f8816602003811015610a3057610a02838260208b6108a8565b60f81b848281518110610a1157fe5b60200101906001600160f81b031916908160001a9053506001016109e7565b50601f87166008035b602081101561092e57610a508260208a84016104c0565b60f81b848281518110610a5f57fe5b60200101906001600160f81b031916908160001a905350600101610a39565b60005b602081101561097c57610a97828260208a6108a8565b60f81b838281518110610aa657fe5b60200101906001600160f81b031916908160001a905350600101610a81565b6040805160f883811c8083526020808202840101909352606092919085901c908390838015610afe578160200160208202803883390190505b50905060005b83811015610b3d57610b1b88828501602002610ea2565b60001b828281518110610b2a57fe5b6020908102919091010152600101610b04565b509695505050505050565b60006060610b5585610ee8565b90508260f81b818581518110610b6757fe5b60200101906001600160f81b031916908160001a9053506104fd81610e5f565b6000610bd78686868686600081518110610b9d57fe5b602002602001015160001c87600181518110610bb557fe5b602002602001015188600281518110610bca57fe5b6020026020010151610f52565b9695505050505050565b60006060610bee87610ee8565b90506000610c018a60208b5b0489610741565b90506020808a0660080110610cd9576000610c248b60208c5b0460010188610741565b905060005b6018601f8c1660080301811015610c7157610c67838260208e0601868460180181518110610c5357fe5b01602001516001600160f81b03191661113c565b9250600101610c29565b506018601f8b16600803015b6008811015610cd257610c9f8260208d840106868460180181518110610c5357fe5b9150610cb28c60208d5b04858c8c610b87565b9b50610cc88c60208d5b04600101848a8a610b87565b9b50600101610c7d565b5050610d18565b60005b6008811015610d0857610cfe828260208d0601858460180181518110610c5357fe5b9150600101610cdc565b50610d158a60208b610705565b99505b509798975050505050505050565b60006060610d3387610ee8565b90506000610d438a60208b610bfa565b90506020808a0660200110610de1576000610d608b60208c610c1a565b905060005b601f8b16602003811015610d9357610d89838260208e0601868481518110610c5357fe5b9250600101610d65565b50601f8a166020035b6020811015610cd257610dbb8260208d840106868481518110610c5357fe5b9150610dc98c60208d610ca9565b9b50610dd78c60208d610cbc565b9b50600101610d9c565b60005b6020811015610d0857610e03828260208d0601858481518110610c5357fe5b9150600101610de4565b60408051602080820184905282518083038201815291830190925280519101205b919050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080805b8351811015610e9b57600882901b9150838160018651030381518110610e8657fe5b016020015160f81c9190911790600101610e64565b5092915050565b600080805b6020811015610ee057600882901b91508481600186602001030381518110610ecb57fe5b016020015160f81c9190911790600101610ea7565b509392505050565b6040805160208082528183019092526060918391839160208201818038833901905050905060005b8351811015610ee0578260f81b8282601f0381518110610f2c57fe5b60200101906001600160f81b031916908160001a90535060089290921c91600101610f10565b600080610f5e87610e0d565b9050610f6b898988610741565b506060610f76611158565b905060028751901b89111561103d5787610f94578992505050610987565b6000610f9f8a6111f9565b88519091505b60018203811015610fca57610fc08c8483815181106107a957fe5b9b50600101610fa5565b5060015b60018203811015611028578a60011660011415610ffc57610ff5848483815181106107a957fe5b935061101c565b61101983828151811061100b57fe5b602002602001015185610e33565b93505b60019a8b1c9a01610fce565b506110338b84610e33565b9350505050610987565b60015b87518110156110bd5760008a60011660011461105c5783611071565b88828151811061106857fe5b60200260200101515b905060008b6001166001146110995789838151811061108c57fe5b602002602001015161109b565b845b90506110a78282610e33565b60019c8d1c9c9095509290920191506110409050565b5087156110cc57509050610987565b8086815181106110d857fe5b60200260200101518414156110ec57600080fd5b60006110f88686610e33565b905080875b895181101561112057611116838583815181106107a957fe5b91506001016110fd565b5083811461112d57600080fd5b509a9950505050505050505050565b6000606061114985610ee8565b905082818581518110610b6757fe5b604080518181526108208101825260609182919060208201610800803883390190505090506111876000610e0d565b8160008151811061119457fe5b602090810291909101015260015b60408110156111f3576111d48260018303815181106111bd57fe5b60200260200101518360018403815181106107a957fe5b8282815181106111e057fe5b60209081029190910101526001016111a2565b50905090565b60008161120857506000610e2e565b611215600183901c6111f9565b6001019050610e2e56fea265627a7a723158203a4b54fe79dfc5077fe63ce73a1844feb5394dbf3a19879d47f0265e422f61b864736f6c63430005110032"

// DeployMerkleUtil deploys a new Ethereum contract, binding an instance of MerkleUtil to it.
func DeployMerkleUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleUtil, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleUtilABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleUtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleUtil{MerkleUtilCaller: MerkleUtilCaller{contract: contract}, MerkleUtilTransactor: MerkleUtilTransactor{contract: contract}, MerkleUtilFilterer: MerkleUtilFilterer{contract: contract}}, nil
}

// MerkleUtil is an auto generated Go binding around an Ethereum contract.
type MerkleUtil struct {
	MerkleUtilCaller     // Read-only binding to the contract
	MerkleUtilTransactor // Write-only binding to the contract
	MerkleUtilFilterer   // Log filterer for contract events
}

// MerkleUtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleUtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleUtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleUtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleUtilSession struct {
	Contract     *MerkleUtil       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleUtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleUtilCallerSession struct {
	Contract *MerkleUtilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleUtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleUtilTransactorSession struct {
	Contract     *MerkleUtilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleUtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleUtilRaw struct {
	Contract *MerkleUtil // Generic contract binding to access the raw methods on
}

// MerkleUtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleUtilCallerRaw struct {
	Contract *MerkleUtilCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleUtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleUtilTransactorRaw struct {
	Contract *MerkleUtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleUtil creates a new instance of MerkleUtil, bound to a specific deployed contract.
func NewMerkleUtil(address common.Address, backend bind.ContractBackend) (*MerkleUtil, error) {
	contract, err := bindMerkleUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleUtil{MerkleUtilCaller: MerkleUtilCaller{contract: contract}, MerkleUtilTransactor: MerkleUtilTransactor{contract: contract}, MerkleUtilFilterer: MerkleUtilFilterer{contract: contract}}, nil
}

// NewMerkleUtilCaller creates a new read-only instance of MerkleUtil, bound to a specific deployed contract.
func NewMerkleUtilCaller(address common.Address, caller bind.ContractCaller) (*MerkleUtilCaller, error) {
	contract, err := bindMerkleUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilCaller{contract: contract}, nil
}

// NewMerkleUtilTransactor creates a new write-only instance of MerkleUtil, bound to a specific deployed contract.
func NewMerkleUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleUtilTransactor, error) {
	contract, err := bindMerkleUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilTransactor{contract: contract}, nil
}

// NewMerkleUtilFilterer creates a new log filterer instance of MerkleUtil, bound to a specific deployed contract.
func NewMerkleUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleUtilFilterer, error) {
	contract, err := bindMerkleUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilFilterer{contract: contract}, nil
}

// bindMerkleUtil binds a generic wrapper to an already deployed contract.
func bindMerkleUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleUtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseMerkleUtilABI parses the ABI
func ParseMerkleUtilABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleUtilABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleUtil *MerkleUtilRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleUtil.Contract.MerkleUtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleUtil *MerkleUtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleUtil.Contract.MerkleUtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleUtil *MerkleUtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleUtil.Contract.MerkleUtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleUtil *MerkleUtilCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleUtil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleUtil *MerkleUtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleUtil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleUtil *MerkleUtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleUtil.Contract.contract.Transact(opts, method, params...)
}

// BufferOp is a free data retrieval call binding the contract method 0x585ec5ad.
//
// Solidity: function bufferOp(uint8 op, bytes32 buf, uint256 offset, uint256 b, bytes proof) constant returns(bytes32)
func (_MerkleUtil *MerkleUtilCaller) BufferOp(opts *bind.CallOpts, op uint8, buf [32]byte, offset *big.Int, b *big.Int, proof []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleUtil.contract.Call(opts, out, "bufferOp", op, buf, offset, b, proof)
	return *ret0, err
}

// BufferOp is a free data retrieval call binding the contract method 0x585ec5ad.
//
// Solidity: function bufferOp(uint8 op, bytes32 buf, uint256 offset, uint256 b, bytes proof) constant returns(bytes32)
func (_MerkleUtil *MerkleUtilSession) BufferOp(op uint8, buf [32]byte, offset *big.Int, b *big.Int, proof []byte) ([32]byte, error) {
	return _MerkleUtil.Contract.BufferOp(&_MerkleUtil.CallOpts, op, buf, offset, b, proof)
}

// BufferOp is a free data retrieval call binding the contract method 0x585ec5ad.
//
// Solidity: function bufferOp(uint8 op, bytes32 buf, uint256 offset, uint256 b, bytes proof) constant returns(bytes32)
func (_MerkleUtil *MerkleUtilCallerSession) BufferOp(op uint8, buf [32]byte, offset *big.Int, b *big.Int, proof []byte) ([32]byte, error) {
	return _MerkleUtil.Contract.BufferOp(&_MerkleUtil.CallOpts, op, buf, offset, b, proof)
}

// GetOp is a free data retrieval call binding the contract method 0x1424c514.
//
// Solidity: function getOp(uint8 op, bytes32 buf, uint256 offset, bytes32[] proof1, bytes32[] proof2) constant returns(uint256)
func (_MerkleUtil *MerkleUtilCaller) GetOp(opts *bind.CallOpts, op uint8, buf [32]byte, offset *big.Int, proof1 [][32]byte, proof2 [][32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MerkleUtil.contract.Call(opts, out, "getOp", op, buf, offset, proof1, proof2)
	return *ret0, err
}

// GetOp is a free data retrieval call binding the contract method 0x1424c514.
//
// Solidity: function getOp(uint8 op, bytes32 buf, uint256 offset, bytes32[] proof1, bytes32[] proof2) constant returns(uint256)
func (_MerkleUtil *MerkleUtilSession) GetOp(op uint8, buf [32]byte, offset *big.Int, proof1 [][32]byte, proof2 [][32]byte) (*big.Int, error) {
	return _MerkleUtil.Contract.GetOp(&_MerkleUtil.CallOpts, op, buf, offset, proof1, proof2)
}

// GetOp is a free data retrieval call binding the contract method 0x1424c514.
//
// Solidity: function getOp(uint8 op, bytes32 buf, uint256 offset, bytes32[] proof1, bytes32[] proof2) constant returns(uint256)
func (_MerkleUtil *MerkleUtilCallerSession) GetOp(op uint8, buf [32]byte, offset *big.Int, proof1 [][32]byte, proof2 [][32]byte) (*big.Int, error) {
	return _MerkleUtil.Contract.GetOp(&_MerkleUtil.CallOpts, op, buf, offset, proof1, proof2)
}

// SetOp is a free data retrieval call binding the contract method 0xc58ed997.
//
// Solidity: function setOp(uint8 op, bytes32 buf, uint256 offset, uint256 b, bytes32[] proof1, bytes32[] nproof1, bytes32[] proof2, bytes32[] nproof2) constant returns(bytes32)
func (_MerkleUtil *MerkleUtilCaller) SetOp(opts *bind.CallOpts, op uint8, buf [32]byte, offset *big.Int, b *big.Int, proof1 [][32]byte, nproof1 [][32]byte, proof2 [][32]byte, nproof2 [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleUtil.contract.Call(opts, out, "setOp", op, buf, offset, b, proof1, nproof1, proof2, nproof2)
	return *ret0, err
}

// SetOp is a free data retrieval call binding the contract method 0xc58ed997.
//
// Solidity: function setOp(uint8 op, bytes32 buf, uint256 offset, uint256 b, bytes32[] proof1, bytes32[] nproof1, bytes32[] proof2, bytes32[] nproof2) constant returns(bytes32)
func (_MerkleUtil *MerkleUtilSession) SetOp(op uint8, buf [32]byte, offset *big.Int, b *big.Int, proof1 [][32]byte, nproof1 [][32]byte, proof2 [][32]byte, nproof2 [][32]byte) ([32]byte, error) {
	return _MerkleUtil.Contract.SetOp(&_MerkleUtil.CallOpts, op, buf, offset, b, proof1, nproof1, proof2, nproof2)
}

// SetOp is a free data retrieval call binding the contract method 0xc58ed997.
//
// Solidity: function setOp(uint8 op, bytes32 buf, uint256 offset, uint256 b, bytes32[] proof1, bytes32[] nproof1, bytes32[] proof2, bytes32[] nproof2) constant returns(bytes32)
func (_MerkleUtil *MerkleUtilCallerSession) SetOp(op uint8, buf [32]byte, offset *big.Int, b *big.Int, proof1 [][32]byte, nproof1 [][32]byte, proof2 [][32]byte, nproof2 [][32]byte) ([32]byte, error) {
	return _MerkleUtil.Contract.SetOp(&_MerkleUtil.CallOpts, op, buf, offset, b, proof1, nproof1, proof2, nproof2)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_MerkleUtil *MerkleUtilFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _MerkleUtil.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inboxAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagesAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logsAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inboxAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagesAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logsAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"_kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_msgData\",\"type\":\"bytes\"}],\"name\":\"executeStepWithMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"385bc114": "executeStep(bytes32,bytes32,bytes32,bytes)",
	"96105dce": "executeStepWithMessage(bytes32,bytes32,bytes32,bytes,uint8,uint256,uint256,address,uint256,bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x608060405234801561001057600080fd5b50615bb480620000216000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063385bc1141461003b57806396105dce14610106575b600080fd5b6100be6004803603608081101561005157600080fd5b8135916020810135916040820135919081019060808101606082013564010000000081111561007f57600080fd5b82018360208201111561009157600080fd5b803590602001918460018302840111640100000000831117156100b357600080fd5b509092509050610206565b6040516001600160401b0383168152602081018260a080838360005b838110156100f25781810151838201526020016100da565b505050509050019250505060405180910390f35b6100be600480360361014081101561011d57600080fd5b8135916020810135916040820135919081019060808101606082013564010000000081111561014b57600080fd5b82018360208201111561015d57600080fd5b8035906020019184600183028401116401000000008311171561017f57600080fd5b9193909260ff833516926020810135926040820135926001600160a01b0360608401351692608081013592919060c081019060a001356401000000008111156101c757600080fd5b8201836020820111156101d957600080fd5b803590602001918460018302840111640100000000831117156101fb57600080fd5b50909250905061027d565b6000610210615880565b61021861589e565b61025a88888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061037792505050565b90506102658161089b565b61026e81610c03565b92509250509550959350505050565b6000610287615880565b61028f61589e565b6102d18f8f8f8f8f8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061037792505050565b90506102fd8a888b8b8a8a8a6040518083838082843760405192018290039091209350610c6692505050565b8160e001818152505061034a8a8a8a8a8a8a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610cd492505050565b60c08201526103588161089b565b61036181610c03565b92509250509c509c9a5050505050505050505050565b61037f61589e565b60008260008151811061038e57fe5b602001015160f81c60f81b60f81c90506000836001815181106103ad57fe5b602001015160f81c60f81b60f81c905060006002905060608360040160ff1660405190808252806020026020018201604052801561040557816020015b6103f2615943565b8152602001906001900390816103ea5790505b50905060608360040160ff1660405190808252806020026020018201604052801561044a57816020015b610437615943565b81526020019060019003908161042f5790505b50905060005b8560ff16811015610488576104658885610dfc565b845185908490811061047357fe5b60209081029190910101529350600101610450565b5060005b8460ff168110156104c4576104a18885610dfc565b83518490849081106104af57fe5b6020908102919091010152935060010161048c565b506104cd615980565b6104d78885610f8e565b809250819550505060008885815181106104ed57fe5b602001015160f81c60f81b60f81c9050600089866001018151811061050e57fe5b01602001516002969096019560f81c905061052761589e565b6040518061020001604052808581526020016105428661103f565b81526020018f81526020018e81526020018d815260200160006001600160401b031681526020016105716110b4565b81526020016000801b815260200160405180604001604052808c60ff16815260200189815250815260200160405180604001604052808b60ff1681526020018881525081526020018460ff16600114151581526020018360ff1681526020018c815260200188815260200160405180604001604052806000604051908082528060200260200182016040528015610612578160200160208202803883390190505b5081526020016000604051908082528060200260200182016040528015610643578160200160208202803883390190505b509052815260408051600081830181815260608301845282528251908152602080820190935282820152910152905060ff8316158061068557508260ff166001145b6040518060400160405280600b81526020016a04241445f494d4d5f5459560ac1b815250906107325760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156106f75781810151838201526020016106df565b50505050905090810190601f1680156107245780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061073b615943565b60ff8416610758578151516107519084906110fb565b90506107f8565b6000875111604051806040016040528060068152602001654e4f5f494d4d60d01b815250906107c85760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106f75781810151838201526020016106df565b506107f5838360000151600001518960018e0360ff16815181106107e857fe5b602002602001015161115f565b90505b610801816111e3565b82515260005b848b0360ff168110156108465761083e88828151811061082357fe5b6020026020010151846000015161130390919063ffffffff16565b600101610807565b5060005b8960ff168110156108875761087f87828151811061086457fe5b6020026020010151846000015161131d90919063ffffffff16565b60010161084a565b50909e9d5050505050505050505050505050565b60008060006128366108b485610160015160ff16611337565b93509350935093506108c68583611aa5565b156108d45750505050610c00565b61010085015151841115610989576108f26108ed6110b4565b6111e3565b6109038660200151602001516111e3565b146040518060400160405280600d81526020016c535441434b5f4d495353494e4760981b815250906109765760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106f75781810151838201526020016106df565b5061098085611b0f565b50505050610c00565b61012085015151831115610a24576109a26108ed6110b4565b6109b38660200151604001516111e3565b146040518060400160405280600b81526020016a4155585f4d495353494e4760a81b815250906109765760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106f75781810151838201526020016106df565b6000841180610a365750846101400151155b8015610a4757506101008501515184145b80610a6f57508461014001518015610a5d575083155b8015610a6f5750610100850151516001145b6040518060400160405280600a815260200169535441434b5f4d414e5960b01b81525090610ade5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106f75781810151838201526020016106df565b50610120850151516040805180820190915260088152674155585f4d414e5960c01b6020820152908414610b535760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106f75781810151838201526020016106df565b50610b61858263ffffffff16565b60005b61010086015151811015610bad57610ba5866101000151602001518281518110610b8a57fe5b6020026020010151876020015161130390919063ffffffff16565b600101610b64565b5060005b61012086015151811015610bfa57610bf2866101200151602001518281518110610bd757fe5b6020026020010151876020015161131d90919063ffffffff16565b600101610bb1565b50505050505b50565b6000610c0d615880565b8260a001516040518060a00160405280610c2a8660000151611b78565b8152602001610c3c8660200151611b78565b81526020018560400151815260200185606001518152602001856080015181525091509150915091565b6040805160f89790971b6001600160f81b03191660208089019190915260609690961b6bffffffffffffffffffffffff191660218801526035870194909452605586019290925260758501526095808501919091528151808503909101815260b59093019052815191012090565b610cdc615943565b60408051600680825260e08201909252606091816020015b610cfc615943565b815260200190600190039081610cf4579050509050610d1d8860ff16611c52565b81600081518110610d2a57fe5b6020026020010181905250610d3e87611c52565b81600181518110610d4b57fe5b6020026020010181905250610d5f86611c52565b81600281518110610d6c57fe5b6020026020010181905250610d89856001600160a01b0316611c52565b81600381518110610d9657fe5b6020026020010181905250610daa84611c52565b81600481518110610db757fe5b6020026020010181905250610dcf8360008551611d0b565b81600581518110610ddc57fe5b6020026020010181905250610df081611e8f565b98975050505050505050565b6000610e06615943565b83518310610e4c576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b600080610e598686611fa8565b91509150610e65611fcf565b60ff168160ff161415610e99576000610e7e8784611fd4565b909350905082610e8d82611c52565b94509450505050610f87565b610ea1612048565b60ff168160ff161415610ec357610eb8868361204d565b935093505050610f87565b610ecb6120ef565b60ff168160ff161415610ee257610eb886836120f4565b610eea612188565b60ff168160ff1610158015610f0b5750610f0261218d565b60ff168160ff16105b15610f47576000610f1a612188565b820390506060610f2b828986612192565b909450905083610f3a82611e8f565b9550955050505050610f87565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b6000610f98615980565b610fa0615980565b6000610100820181905280610fb58787611fd4565b9096509150610fc487876120f4565b60208501529550610fd587876120f4565b60408501529550610fe68787610dfc565b60608501529550610ff78787610dfc565b608085015295506110088787611fd4565b60a085015295506110198787611fd4565b90965090506110288787610dfc565b60e085015291835260c08301529590945092505050565b611047615980565b60405180610120016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e00151815260200183610100015181525090505b919050565b6110bc615943565b604080516000808252602082019092526110f6916110f0565b6110dd615943565b8152602001906001900390816110d55790505b50611e8f565b905090565b611103615943565b6040805160608101825260ff8516815260208082018590528251600080825291810184526111569383019161114e565b61113b615943565b8152602001906001900390816111335790505b50905261222b565b90505b92915050565b611167615943565b604080516001808252818301909252606091816020015b611186615943565b81526020019060019003908161117e57905050905082816000815181106111a957fe5b60200260200101819052506111da60405180606001604052808760ff1681526020018681526020018381525061222b565b95945050505050565b60006111ed611fcf565b60ff16826080015160ff16141561121057815161120990612299565b90506110af565b611218612048565b60ff16826080015160ff1614156112365761120982602001516122bd565b61123e6120ef565b60ff16826080015160ff16141561126057815160a083015161120991906123ba565b611268612188565b60ff16826080015160ff1614156112a157611281615943565b61128e836040015161240b565b9050611299816111e3565b9150506110af565b6112a961256d565b60ff16826080015160ff1614156112c2575080516110af565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b611311826020015182612572565b82602001819052505050565b61132b826040015182612572565b82604001819052505050565b60008080612836600185148061134d5750600285145b806113585750600385145b1561137257506002925060009150600390506125f0611a9e565b60048514806113815750600685145b1561139b575060029250600091506004905061284f611a9e565b60058514806113aa5750600785145b156113c4575060029250600091506007905061284f611a9e565b60088514806113d35750600985145b156113ed5750600392506000915060049050612921611a9e565b600a85141561140b57506002925060009150601990506125f0611a9e565b600b85141561142957506002925060009150600790506125f0611a9e565b60108514806114385750601185145b806114435750601285145b8061144e5750601385145b806114595750601685145b806114645750601785145b8061146f5750601885145b15611488575060029250600091508290506125f0611a9e565b60148514156114a557506002925060009150829050612a0e611a9e565b60158514156114c257506001925060009150829050612a67611a9e565b60198514156114df57506001925060009150829050612ac3611a9e565b601a8514806114ee5750601b85145b806114f95750601c85145b806115045750601d85145b1561151e57506002925060009150600490506125f0611a9e565b602085141561153c5750600192506000915060079050612b0f611a9e565b602185141561155a5750600192506000915060039050612b3c611a9e565b602285141561157857506002925060009150600890506125f0611a9e565b6023851415611597575060019250600091506102589050612b66611a9e565b60248514156115b55750600392506000915060fa9050612d44611a9e565b60308514156115d257506001925060009150829050612e02611a9e565b60318514156115ef57506000925082915060019050612e10611a9e565b603285141561160c57506000925082915060019050612e27611a9e565b603385141561162a5750600192506000915060029050612e3e611a9e565b60348514156116485750600192506000915060049050612e58611a9e565b60358514156116665750600292506000915060049050612e99611a9e565b603685141561168357506000925082915060029050612f0d611a9e565b60378514156116a057506000925082915060019050612f51611a9e565b60388514156116bd57506001925060009150829050612f6a611a9e565b60398514156116da57506000925060019150819050612f81611a9e565b603a8514156116f757506000925082915060029050612f98611a9e565b603b85141561171457506000925082915060019050610c00611a9e565b603c85141561173157506000925082915060019050612fc6611a9e565b603d85141561174e57506001925060009150829050612fe2611a9e565b604085141561176b57506001925060009150829050613026611a9e565b6041851415611789575060029250600091506001905061305c611a9e565b60428514156117a757506003925060009150600190506130b9611a9e565b60438514156117c5575060029250600091506001905061313d611a9e565b60448514156117e3575060039250600091506001905061317c611a9e565b6050851415611800575060029250600091508290506131e2611a9e565b605185141561181e575060039250600091506028905061327f611a9e565b605285141561183c575060019250600091506002905061333c611a9e565b605385141561185957506001925082915060039050613387611a9e565b60548514156118775750600292506001915060299050613409611a9e565b606085141561189457506000925082915060649050610c00611a9e565b60618514156118b257506001925060009150606490506134c6611a9e565b60708514156118d0575060019250600091506064905061350d611a9e565b60718514156118ee5750600192506000915060289050613595611a9e565b607285141561190b57506000925082915060289050613612611a9e565b607385141561192857506000925082915060059050613672611a9e565b6074851415611945575060009250829150600a905061367b611a9e565b607585141561196257506001925060009150819050613688611a9e565b607685141561197f575060009250829150600190506136c3611a9e565b607785141561199c575060009250829150601990506136dd611a9e565b60788514156119ba575060029250600091506019905061372e611a9e565b60798514156119d857506003925060009150601990506137a6611a9e565b607b8514156119f5575060009250829150600a9050613837611a9e565b6080851415611a1457506004925060009150614e20905061387f611a9e565b6081851415611a3357506004925060009150610dac90506139f8611a9e565b6082851415611a5357506003925060009150620140509050613b35611a9e565b6083851415611a72575060019250600091506103e89050613c30611a9e565b60a0851415611a8f57506001925060009150829050613f31611a9e565b50600092508291508190506136725b9193509193565b60a080830180516001600160401b03908401811690915260208401519091015160009183161115611aef57602083015160001960a090910152611ae783611b0f565b506001611159565b50602082015160a00180516001600160401b038316900390526000611159565b60408051600160f81b6020808301919091526000602183018190526022808401919091528351808403909101815260429092019092528051908201209082015160c001511415611b6b57611b668160200151613f6c565b610c00565b6020015160c08101519052565b600060028261010001511415611b90575060006110af565b60018261010001511415611ba6575060016110af565b81516020830151611bb6906111e3565b611bc384604001516111e3565b611bd085606001516111e3565b611bdd86608001516111e3565b8660a001518760c00151611bf48960e001516111e3565b60405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012090506110af565b611c5a615943565b6040805160c0810182528381528151606081018352600080825260208281018290528451828152808201865293949085019390830191611cb0565b611c9d615943565b815260200190600190039081611c955790505b50905281526040805160008082526020828101909352919092019190611cec565b611cd9615943565b815260200190600190039081611cd15790505b5081526000602082018190526040820152600160609091015292915050565b611d13615943565b60208204611d1f615943565b611d276110b4565b60408051600280825260608281019093529293509091816020015b611d4a615943565b815260200190600190039081611d4257905050905060005b83811015611dcb57611d87611d8289602084028a0163ffffffff613f7716565b611c52565b82600081518110611d9457fe5b60200260200101819052508282600181518110611dad57fe5b6020026020010181905250611dc18261240b565b9250600101611d62565b506020850615611e41576000611ded88601f198989010163ffffffff613f7716565b9050602086066020036008021b611e0381611c52565b82600081518110611e1057fe5b60200260200101819052508282600181518110611e2957fe5b6020026020010181905250611e3d8261240b565b9250505b611e4a85611c52565b81600081518110611e5757fe5b60200260200101819052508181600181518110611e7057fe5b6020026020010181905250611e848161240b565b979650505050505050565b611e97615943565b611ea18251613fd0565b611ef2576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015611f2957838181518110611f0c57fe5b602002602001015160a00151820191508080600101915050611ef7565b506040805160c0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190611f83565b611f70615943565b815260200190600190039081611f685790505b5090528152602081019490945260006040850152600360608501526080909301525090565b60008082600101848481518110611fbb57fe5b016020015190925060f81c90509250929050565b600090565b60008082845110158015611fec575060208385510310155b612029576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b6020830161203d858563ffffffff613f7716565b915091509250929050565b600190565b6000612057615943565b82600080612063615943565b600061206f8986611fa8565b909550935061207e8986611fa8565b9095509250600160ff8516141561209f576120998986610dfc565b90955091505b6120a98986613fd7565b9095509050600160ff851614156120d457846120c684838561115f565b965096505050505050610f87565b846120df84836110fb565b9650965050505050509250929050565b600290565b60006120fe615943565b82845110158015612113575060408385510310155b61214f576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b60008061215c8686613fd7565b909450915061216b8685611fd4565b90945090508361217b8383613fee565b9350935050509250929050565b600390565b600d90565b60006060600083905060608660ff166040519080825280602002602001820160405280156121da57816020015b6121c7615943565b8152602001906001900390816121bf5790505b50905060005b8760ff168160ff16101561221e576121f88784610dfc565b8351849060ff851690811061220957fe5b602090810291909101015292506001016121e0565b5090969095509350505050565b612233615943565b6040805160c081018252600080825260208083018690528351828152908101845291928301919061227a565b612267615943565b81526020019060019003908161225f5790505b5081526000602082015260016040820181905260609091015292915050565b60408051602080820193909352815180820384018152908201909152805191012090565b60006002826040015151106122ce57fe5b604082015151612333576122e0612048565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b9093166021850152602280850191909152825180850390910181526042909301909152815191012090506110af565b61233b612048565b8260000151612361846040015160008151811061235457fe5b60200260200101516111e3565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60006123c4612188565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b612413615943565b600882511115612461576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825160405190808252806020026020018201604052801561248e578160200160208202803883390190505b508051909150600160005b828110156124f1576124b086828151811061235457fe5b8482815181106124bc57fe5b6020026020010181815250508581815181106124d457fe5b602002602001015160a00151820191508080600101915050612499565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561253657818101518382015260200161251e565b50505050905001925050506040516020818303038152906040528051906020012090506125638183613fee565b9695505050505050565b606490565b61257a615943565b6040805160028082526060828101909352816020015b612598615943565b81526020019060019003908161259057905050905082816000815181106125bb57fe5b602002602001018190525083816001815181106125d457fe5b60200260200101819052506125e88161240b565b949350505050565b6125f8615943565b6126068261010001516140a6565b9050612610615943565b61261e8361010001516140a6565b9050612629826140e8565b158061263b5750612639816140e8565b155b1561265057612649836140f3565b5050610c00565b8151815161016085015160009060ff16600114156126715750818101612838565b61016086015160ff166002141561268b5750818102612838565b61016086015160ff16600314156126a55750808203612838565b61016086015160ff16600a14156126bf575080820a612838565b61016086015160ff16600b14156126d9575080820b612838565b61016086015160ff16601014156126f35750808210612838565b61016086015160ff166011141561270d5750808211612838565b61016086015160ff16601214156127275750808212612838565b61016086015160ff16601314156127415750808213612838565b61016086015160ff166016141561275b5750818116612838565b61016086015160ff16601714156127755750818117612838565b61016086015160ff166018141561278f5750818118612838565b61016086015160ff16601a14156127a9575080821a612838565b61016086015160ff16601b14156127c3575080821b612838565b61016086015160ff16601c14156127dd575080821c612838565b61016086015160ff16601d14156127f7575080821d612838565b61016086015160ff1660221415612836575060408051602080820185905281830184905282518083038401815260609092019092528051910120612838565bfe5b610bfa86610100015161284a83611c52565b614112565b612857615943565b6128658261010001516140a6565b905061286f615943565b61287d8361010001516140a6565b9050612888826140e8565b158061289a5750612898816140e8565b155b806128a457508051155b156128b257612649836140f3565b8151815161016085015160009060ff16600414156128d35750808204612838565b61016086015160ff16600514156128ed5750808205612838565b61016086015160ff16600614156129075750808206612838565b61016086015160ff16600714156128365750808207612838565b612929615943565b6129378261010001516140a6565b9050612941615943565b61294f8361010001516140a6565b9050612959615943565b6129678461010001516140a6565b9050612972836140e8565b15806129845750612982826140e8565b155b806129955750612993816140e8565b155b8061299f57508051155b156129b5576129ad846140f3565b505050610c00565b82518251825161016087015160009060ff16600814156129da578183850890506129f2565b61016088015160ff1660091415612836578183850990505b612a0488610100015161284a83611c52565b5050505050505050565b612a16615943565b612a248261010001516140a6565b9050612a2e615943565b612a3c8361010001516140a6565b9050612a6283610100015161284a612a53846111e3565b612a5c866111e3565b1461413c565b505050565b612a6f615943565b612a7d8261010001516140a6565b9050612a88816140e8565b612aa457612a9f82610100015161284a6000611c52565b612abf565b8051610100830151811590612abc9061284a83611c52565b50505b5050565b612acb615943565b612ad98261010001516140a6565b9050612ae4816140e8565b612af757612af1826140f3565b50610c00565b8051610100830151811990612abc9061284a83611c52565b612b17615943565b612b258261010001516140a6565b9050612abf82610100015161284a611d82846111e3565b612b44615943565b612b528261010001516140a6565b9050612abf82610100015161284a8361415e565b612b6e615943565b612b7c8261010001516140a6565b9050612b87816141c7565b1580612b995750806040015151600714155b15612ba757612af1826140f3565b604081015160005b6007811015612be957612bd4828281518110612bc757fe5b60200260200101516140e8565b612be1576129ad846140f3565b600101612baf565b50612bf26159eb565b60005b6019811015612c565760406003821602836004830481518110612c1457fe5b602002602001015160000151901c6001600160401b03168260058381612c3657fe5b04600584066005020160198110612c4957fe5b6020020152600101612bf5565b50612c60816141d4565b604080516007808252610100820190925291925060609190816020015b612c85615943565b815260200190600190039081612c7d57905050905060005b6007811015612ccf57612cb06000611c52565b828281518110612cbc57fe5b6020908102919091010152600101612c9d565b5060005b6019811015612d2a57604060038216028360058304600584066005020160198110612cfa57fe5b6020020151901b826004830481518110612d1057fe5b602090810291909101015180519091179052600101612cd3565b50612d3d85610100015161284a83611e8f565b5050505050565b612d4c615943565b612d5a8261010001516140a6565b9050612d64615943565b612d728361010001516140a6565b9050612d7c615943565b612d8a8461010001516140a6565b9050612d95836140e8565b1580612da75750612da5826140e8565b155b80612db85750612db6816140e8565b155b15612dc6576129ad846140f3565b8251825182516101008701516040805180820190915283815260208101839052612df9919061284a90611d829087614dc5565b50505050505050565b612abf8161010001516140a6565b610c00816101000151826020015160800151614112565b610c00816101000151826020015160600151614112565b612e4c8161010001516140a6565b60209091015160600152565b612e60615943565b612e6e8261010001516140a6565b9050612e79816154e2565b612e8657612af1826140f3565b612e8f816111e3565b6020830151525050565b612ea1615943565b612eaf8261010001516140a6565b9050612eb9615943565b612ec78361010001516140a6565b9050612ed2826154e2565b1580612ee45750612ee2816140e8565b155b15612ef257612649836140f3565b805115612a6257612f02826111e3565b602084015152505050565b61010081015151600090158015612f3d5750612f2a6108ed6110b4565b612f3b8360200151602001516111e3565b145b9050612abf82610100015161284a8361413c565b610100810151815151610c00919061284a9060016154ef565b610c0081610120015161284a8361010001516140a6565b610c0081610100015161284a8361012001516140a6565b61012081015151600090158015612f3d5750612fb56108ed6110b4565b612f3b8360200151604001516111e3565b610c0081610100015161284a836020015160c0015160016154ef565b612fea615943565b612ff88261010001516140a6565b9050613003816154e2565b61301057612af1826140f3565b613019816111e3565b602083015160c001525050565b61302e615943565b61303c8261010001516140a6565b905061304d82610100015182614112565b612abf82610100015182614112565b613064615943565b6130728261010001516140a6565b905061307c615943565b61308a8361010001516140a6565b905061309b83610100015182614112565b6130aa83610100015183614112565b612a6283610100015182614112565b6130c1615943565b6130cf8261010001516140a6565b90506130d9615943565b6130e78361010001516140a6565b90506130f1615943565b6130ff8461010001516140a6565b905061311084610100015182614112565b61311f84610100015183614112565b61312e84610100015184614112565b612abc84610100015182614112565b613145615943565b6131538261010001516140a6565b905061315d615943565b61316b8361010001516140a6565b90506130aa83610100015183614112565b613184615943565b6131928261010001516140a6565b905061319c615943565b6131aa8361010001516140a6565b90506131b4615943565b6131c28461010001516140a6565b90506131d384610100015184614112565b61312e84610100015183614112565b6131ea615943565b6131f88261010001516140a6565b9050613202615943565b6132108361010001516140a6565b905061321b826140e8565b158061322d575061322b816141c7565b155b80613247575061323c816155a7565b60ff16826000015110155b1561325557612649836140f3565b612a62836101000151826040015184600001518151811061327257fe5b6020026020010151614112565b613287615943565b6132958261010001516140a6565b905061329f615943565b6132ad8361010001516140a6565b90506132b7615943565b6132c58461010001516140a6565b90506132d0836140e8565b15806132e257506132e0826141c7565b155b806132fc57506132f1826155a7565b60ff16836000015110155b1561330a576129ad846140f3565b60408201518351815183918391811061331f57fe5b6020026020010181905250612d3d85610100015161284a83611e8f565b613344615943565b6133528261010001516140a6565b905061335d816141c7565b61336a57612af1826140f3565b612abf82610100015161284a61337f846155a7565b60ff16611c52565b61338f615943565b61339d8261010001516140a6565b90506133a7615943565b6133b58361012001516140a6565b90506133c0826140e8565b15806133d257506133d0816141c7565b155b806133ec57506133e1816155a7565b60ff16826000015110155b156133fa57612649836140f3565b61325583610120015182614112565b613411615943565b61341f8261010001516140a6565b9050613429615943565b6134378361010001516140a6565b9050613441615943565b61344f8461012001516140a6565b905061345a816141c7565b158061346c575061346a836140e8565b155b80613486575061347b816155a7565b60ff16836000015110155b15613494576129ad846140f3565b6040810151835181518491839181106134a957fe5b6020026020010181905250612d3d85610120015161284a83611e8f565b80608001516134dc6108ed8361010001516140a6565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120608090910152565b613515615943565b6135238261010001516140a6565b90506127108160a00151118061353f575061353d816155ce565b155b1561354d57612af1826140f3565b816060015161355b826111e3565b6040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208260600181815250505050565b61359d615943565b6135ab8261010001516140a6565b90506135b86108ed6110b4565b6135c9836020015160e001516111e3565b146135e1576135d7826156cc565b602083015160e001525b612abf82610100015161284a6135f6846111e3565b612a5c866020015160e001516040015160018151811061235457fe5b61361d6108ed6110b4565b61362e826020015160e001516111e3565b146136605761364a816101000151826020015160e00151614112565b6136526110b4565b602082015160e00152610c00565b610c0081610100015161284a836156cc565b610c00816140f3565b610c008160200151615765565b613690615943565b61369e8261010001516140a6565b90506136a9816140e8565b6136b657612af1826140f3565b51602082015160a0015250565b610c0081610100015161284a836020015160a00151611c52565b61010081015160408051600160f81b6020808301919091526000602183018190526022808401919091528351808403909101815260429092019092528051910120610c00919061284a9060016154ef565b613736615943565b6137448261010001516140a6565b905061374e615943565b61375c8361010001516140a6565b9050613767826140e8565b15806137795750613777816154e2565b155b1561378757612649836140f3565b612a6283610100015161284a84600001516137a1856111e3565b6110fb565b6137ae615943565b6137bc8261010001516140a6565b90506137c6615943565b6137d48361010001516140a6565b90506137de615943565b6137ec8461010001516140a6565b90506137f7836140e8565b15806138095750613807816154e2565b155b15613817576129ad846140f3565b612abc84610100015161284a8560000151613831856111e3565b8661115f565b6040805160008082526020820190925260609161386a565b613857615943565b81526020019060019003908161384f5790505b509050612abf82610100015161284a83611e8f565b613887615943565b6138958261010001516140a6565b905061389f615943565b6138ad8361010001516140a6565b90506138b7615943565b6138c58461010001516140a6565b90506138cf615943565b6138dd8561010001516140a6565b90506138e8846140e8565b15806138fa57506138f8836140e8565b155b8061390b5750613909826140e8565b155b8061391c575061391a816140e8565b155b1561392a57610980856140f3565b8351835183511580159061394057508351600114155b156139635761395887610100015161284a6000611c52565b505050505050610c00565b83518351604080516000808252602080830180855285905260ff601b9096019586168385015260608301889052608083018790529251909260019260a080820193601f1981019281900390910190855afa1580156139c5573d6000803e3d6000fd5b5050506020604051035190506139ec8a610100015161284a836001600160a01b0316611c52565b50505050505050505050565b613a00615943565b613a0e8261010001516140a6565b9050613a18615943565b613a268361010001516140a6565b9050613a30615943565b613a3e8461010001516140a6565b9050613a48615943565b613a568561010001516140a6565b9050613a61846140e8565b1580613a735750613a71836140e8565b155b80613a845750613a82826140e8565b155b80613a955750613a93816140e8565b155b15613aa357610980856140f3565b613aab615a0a565b5060408051608081018252855181528451602082015283519181019190915281516060820152613ad9615a28565b600060408260808560066107d05a03fa905080613b0557613af9886140f3565b50505050505050610c00565b610100880151613b209061284a8460015b6020020151611c52565b610100880151612a049061284a846000613b16565b613b3d615943565b613b4b8261010001516140a6565b9050613b55615943565b613b638361010001516140a6565b9050613b6d615943565b613b7b8461010001516140a6565b9050613b86836140e8565b1580613b985750613b96826140e8565b155b80613ba95750613ba7816140e8565b155b15613bb7576129ad846140f3565b613bbf615a46565b50604080516060810182528451815283516020820152825191810191909152613be6615a28565b600060408260808560076107d05a03fa905080613c0657613958876140f3565b610100870151613c1b9061284a846001613b16565b610100870151612df99061284a846000613b16565b613c38615943565b613c468261010001516140a6565b9050613c50615a64565b60005b601e811015613ce057613c65836141c7565b613c72576129ad846140f3565b60408301518051613c835750613ce0565b8051600214613c9557610980856140f3565b80600081518110613ca257fe5b60200260200101518383601e8110613cb657fe5b6020020152805181906001908110613cca57fe5b6020908102919091010151935050600101613c53565b613cef846207a1208302611aa5565b15613cfc57505050610c00565b613d05836141c7565b1580613d15575060408301515115155b15613d23576129ad846140f3565b613d2b615a92565b60005b82811015613eed57613d3e615943565b8482601e8110613d4a57fe5b60200201519050613d5a816141c7565b613d6757613958876140f3565b60408101518051600614613d7e57613af9886140f3565b60005b6006811015613dbb57613d99828281518110612bc757fe5b613db357613da6896140f3565b5050505050505050610c00565b600101613d81565b5080600081518110613dc957fe5b602002602001015160000151848460060260b48110613de457fe5b6020020152805181906001908110613df857fe5b602002602001015160000151848460060260010160b48110613e1657fe5b6020020152805181906003908110613e2a57fe5b602002602001015160000151848460060260020160b48110613e4857fe5b6020020152805181906002908110613e5c57fe5b602002602001015160000151848460060260030160b48110613e7a57fe5b6020020152805181906005908110613e8e57fe5b602002602001015160000151848460060260040160b48110613eac57fe5b6020020152805181906004908110613ec057fe5b602002602001015160000151848460060260050160b48110613ede57fe5b60200201525050600101613d2e565b5060c08202613efa615ab1565b6000602082848660086107d05a03fa905080613f1957613af9886140f3565b6101008801518251612a04919061284a90151561413c565b610c0081610100015161284a6000801b6040516020018082815260200191505060405160208183030381529060405280519060200120615770565b600161010090910152565b60008160200183511015613fc7576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6008101590565b6000806020830161203d858563ffffffff613f7716565b613ff6615943565b6040805160c081018252848152815160608101835260008082526020828101829052845182815280820186529394908501939083019161404c565b614039615943565b8152602001906001900390816140315790505b50905281526040805160008082526020828101909352919092019190614088565b614075615943565b81526020019060019003908161406d5790505b50815260006020820152600260408201526060019290925250919050565b6140ae615943565b6140b6615943565b82602001516001846000015103815181106140cd57fe5b60209081029190910101518351600019018452915050919050565b6080015160ff161590565b6140fc81611b0f565b6101008101516000908190526101209091015152565b80826020015183600001518151811061412757fe5b60209081029190910101525080516001019052565b614144615943565b8115614154576112096001611c52565b6112096000611c52565b614166615943565b816080015160ff16600214156141b7576040805162461bcd60e51b8152602060048201526011602482015270696e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b611159826080015160ff16611c52565b6080015160ff1660031490565b6141dc6159eb565b6141e4615880565b6141ec615880565b6141f46159eb565b6141fc615acf565b60405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060008090505b6018811015614dba576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e0151181818189288018390526001600160401b0360028202166001603f1b91829004179092188652510485600260200201516002026001600160401b03161785600060200201511884600160200201526001603f1b85600360200201518161444857fe5b0485600360200201516002026001600160401b03161785600160200201511884600260200201526001603f1b85600460200201518161448357fe5b0485600460200201516002026001600160401b031617856002600581106144a657fe5b602002015118606085015284516001603f1b9086516060808901519390920460029091026001600160401b031617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a01805190911890529084525163100000009060208901516001600160401b03641000000000909102169190041761010084015260408701516001603d1b9060408901516001600160401b03600890910216919004176101608401526060870151628000009060608901516001600160401b036502000000000090910216919004176102608401526080870151654000000000009060808901516001600160401b036204000090910216919004176102c084015260a08701516001603f1b900487600560200201516002026001600160401b0316178360026019811061471057fe5b602002015260c0870151621000008104651000000000009091026001600160401b039081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166001603e1b909104176103008501526101408801516101408901516001600160401b036001603e1b90910216919004176080840152610160870151670400000000000000906101608901516001600160401b036040909102169190041760e084015261018087015162200000906101808901516001600160401b036508000000000090910216919004176101408401526101a08701516602000000000000906101a08901516001600160401b0361800090910216919004176102408401526101c08701516008906101c08901516001600160401b036001603d1b90910216919004176102a08401526101e0870151641000000000906101e08901516001600160401b03631000000090910216919004176020840152610200808801516102008901516001600160401b0366800000000000009091021691900417610120840152610220870151648000000000906102208901516001600160401b036302000000909102169190041761018084015261024087015165080000000000906102408901516001600160401b036220000090910216919004176101e0840152610260870151610100906102608901516001600160401b03600160381b90910216919004176102e0840152610280870151642000000000906102808901516001600160401b036308000000909102169190041760608401526102a087015165100000000000906102a08901516001600160401b0362100000909102169190041760c08401526102c08701516302000000906102c08901516001600160401b0364800000000090910216919004176101c08401526102e0870151600160381b906102e08901516001600160401b036101009091021691900417610220840152610300870151660400000000000090048760186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f6020020152826001602002015183601560200201511916836010602002015118876010602002015282600260200201518360166020020151191683601160200201511887601160200201528260036020020151836017602002015119168360126020020151188760126020020152826004602002015183601860200201511916836013602002015118876013602002015282600560200201518360006020020151191683601460200201511887601460200201528260066020020151836001602002015119168360156020020151188760156020020152826007602002015183600260200201511916836016602002015118876016602002015282600860200201518360036020020151191683601760200201511887601760200201528260096020020151836004602002015119168360186020020151188760186020020152818160188110614da857fe5b60200201518751188752600101614323565b509495945050505050565b6000614dcf615aee565b50604080516108008101825263428a2f9881526371374491602082015263b5c0fbcf9181019190915263e9b5dba56060820152633956c25b60808201526359f111f160a082015263923f82a460c082015263ab1c5ed560e082015263d807aa986101008201526312835b0161012082015263243185be61014082015263550c7dc36101608201526372be5d746101808201526380deb1fe6101a0820152639bdc06a76101c082015263c19bf1746101e082015263e49b69c161020082015263efbe4786610220820152630fc19dc661024082015263240ca1cc610260820152632de92c6f610280820152634a7484aa6102a0820152635cb0a9dc6102c08201526376f988da6102e082015263983e515261030082015263a831c66d61032082015263b00327c861034082015263bf597fc761036082015263c6e00bf361038082015263d5a791476103a08201526306ca63516103c082015263142929676103e08201526327b70a85610400820152632e1b2138610420820152634d2c6dfc6104408201526353380d1361046082015263650a735461048082015263766a0abb6104a08201526381c2c92e6104c08201526392722c856104e082015263a2bfe8a161050082015263a81a664b61052082015263c24b8b7061054082015263c76c51a361056082015263d192e81961058082015263d69906246105a082015263f40e35856105c082015263106aa0706105e08201526319a4c116610600820152631e376c08610620820152632748774c6106408201526334b0bcb561066082015263391c0cb3610680820152634ed8aa4a6106a0820152635b9cca4f6106c082015263682e6ff36106e082015263748f82ee6107008201526378a5636f6107208201526384c87814610740820152638cc702086107608201526390befffa61078082015263a4506ceb6107a082015263bef9a3f76107c082015263c67178f26107e082015261509a615aee565b60005b60088163ffffffff1610156151275763ffffffff6020820260e003168660006020020151901c828263ffffffff16604081106150d557fe5b63ffffffff92831660209182029290920191909152820260e003168660016020020151901c828260080163ffffffff166040811061510f57fe5b63ffffffff909216602092909202015260010161509d565b5060106000805b60408363ffffffff16101561528357600384600f850363ffffffff166040811061515457fe5b602002015163ffffffff16901c61518585600f860363ffffffff166040811061517957fe5b6020020151601261582d565b6151a986600f870363ffffffff166040811061519d57fe5b6020020151600761582d565b18189150600a846002850363ffffffff16604081106151c457fe5b602002015163ffffffff16901c6151f5856002860363ffffffff16604081106151e957fe5b6020020151601361582d565b615219866002870363ffffffff166040811061520d57fe5b6020020151601161582d565b1818905080846007850363ffffffff166040811061523357fe5b602002015183866010870363ffffffff166040811061524e57fe5b6020020151010101848463ffffffff166040811061526857fe5b63ffffffff909216602092909202015260019092019161512e565b61528b615b0d565b600093505b60088463ffffffff1610156152dc578360200260e00363ffffffff1688901c818563ffffffff16600881106152c157fe5b63ffffffff9092166020929092020152600190930192615290565b60008060008096505b60408763ffffffff16101561543157608084015161530490601961582d565b608085015161531490600b61582d565b608086015161532490600661582d565b18189450878763ffffffff166040811061533a57fe5b6020020151898863ffffffff166040811061535157fe5b6020020151608086015160a087015160c088015161537092919061584b565b8787600760200201510101010192506153918460006020020151601661582d565b845161539e90600d61582d565b85516153ab90600261582d565b6040870180516020890180518a5160c08c01805163ffffffff90811660e08f015260a08e018051821690925260808e018051821690925260608e0180518e018216909252808616909152808316909552848116909252808318919091169116189290911892909218818101868101909316875260019990990198975090925090506152e5565b600096505b60088763ffffffff161015615485578660200260e00363ffffffff168b901c848863ffffffff166008811061546757fe5b60200201805163ffffffff9201919091169052600190960195615436565b60008097505b60088863ffffffff1610156154d2578760200260e00363ffffffff16858963ffffffff16600881106154b957fe5b602002015160019099019863ffffffff16901b1761548b565b9c9b505050505050505050505050565b6080015160ff1660011490565b6154f7615943565b6040805160c081018252848152815160608101835260008082526020828101829052845182815280820186529394908501939083019161554d565b61553a615943565b8152602001906001900390816155325790505b50905281526040805160008082526020828101909352919092019190615589565b615576615943565b81526020019060019003908161556e5790505b50815260006020820152606460408201526060019290925250919050565b608081015160009060ff16600314156155c657506040810151516110af565b5060016110af565b608081015160009060ff166155e5575060016110af565b608082015160ff16600114156155fd575060006110af565b608082015160ff1660021415615651576040805162461bcd60e51b81526020600482015260146024820152736d75737420686176652066756c6c2076616c756560601b604482015290519081900360640190fd5b608082015160ff16600314156156b45760408201515160005b818110156156a9576156928460400151828151811061568557fe5b60200260200101516155ce565b6156a1576000925050506110af565b60010161566a565b5060019150506110af565b608082015160ff16606414156112c2575060006110af565b6156d4615943565b60e082015160408051808201909152600981526812539093d617d5905360ba1b6020820152906157455760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156106f75781810151838201526020016106df565b5061575882604001518360e00151615854565b60408301525060c0015190565b600261010090910152565b615778615943565b6040805160c08101825260008082528251606081018452818152602081810183905284518381528082018652939490850193919290830191906157d1565b6157be615943565b8152602001906001900390816157b65790505b5090528152604080516000808252602082810190935291909201919061580d565b6157fa615943565b8152602001906001900390816157f25790505b50815260208101849052600c604082015260016060909101529050919050565b63ffffffff9182166020829003831681901b919092169190911c1790565b82191691161890565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6040518060a001604052806005906020820280388339509192915050565b6040518061020001604052806158b2615980565b81526020016158bf615980565b81526000602082018190526040820181905260608201819052608082015260a0016158e8615943565b8152600060208201526040016158fc615b2c565b8152602001615909615b2c565b815260006020820181905260408201819052606080830152608082015260a001615931615b46565b815260200161593e615b46565b905290565b6040518060c001604052806000815260200161595d615b60565b815260606020820181905260006040830181905290820181905260809091015290565b604080516101208101909152600081526020810161599c615943565b81526020016159a9615943565b81526020016159b6615943565b81526020016159c3615943565b815260006020820181905260408201526060016159de615943565b8152602001600081525090565b6040518061032001604052806019906020820280388339509192915050565b60405180608001604052806004906020820280388339509192915050565b60405180604001604052806002906020820280388339509192915050565b60405180606001604052806003906020820280388339509192915050565b604051806103c00160405280601e905b615a7c615943565b815260200190600190039081615a745790505090565b60405180611680016040528060b4906020820280388339509192915050565b60405180602001604052806001906020820280388339509192915050565b6040518061030001604052806018906020820280388339509192915050565b6040518061080001604052806040906020820280388339509192915050565b6040518061010001604052806008906020820280388339509192915050565b604051806040016040528060008152602001606081525090565b604051806040016040528060608152602001606081525090565b604080516060808201835260008083526020830152918101919091529056fea265627a7a72315820e5aa24eb104c5dd23d5f34c0439e7e7079f1f521af10303a5ee79d10f179ab7f64736f6c63430005110032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// OneStepProof is an auto generated Go binding around an Ethereum contract.
type OneStepProof struct {
	OneStepProofCaller     // Read-only binding to the contract
	OneStepProofTransactor // Write-only binding to the contract
	OneStepProofFilterer   // Log filterer for contract events
}

// OneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofSession struct {
	Contract     *OneStepProof     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofCallerSession struct {
	Contract *OneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofTransactorSession struct {
	Contract     *OneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofRaw struct {
	Contract *OneStepProof // Generic contract binding to access the raw methods on
}

// OneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofCallerRaw struct {
	Contract *OneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofTransactorRaw struct {
	Contract *OneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof creates a new instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProof(address common.Address, backend bind.ContractBackend) (*OneStepProof, error) {
	contract, err := bindOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// NewOneStepProofCaller creates a new read-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofCaller, error) {
	contract, err := bindOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofCaller{contract: contract}, nil
}

// NewOneStepProofTransactor creates a new write-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofTransactor, error) {
	contract, err := bindOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTransactor{contract: contract}, nil
}

// NewOneStepProofFilterer creates a new log filterer instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofFilterer, error) {
	contract, err := bindOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofFilterer{contract: contract}, nil
}

// bindOneStepProof binds a generic wrapper to an already deployed contract.
func bindOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseOneStepProofABI parses the ABI
func ParseOneStepProofABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.OneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transact(opts, method, params...)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x385bc114.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof) constant returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof *OneStepProofCaller) ExecuteStep(opts *bind.CallOpts, inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	ret := new(struct {
		Gas    uint64
		Fields [5][32]byte
	})
	out := ret
	err := _OneStepProof.contract.Call(opts, out, "executeStep", inboxAcc, messagesAcc, logsAcc, proof)
	return *ret, err
}

// ExecuteStep is a free data retrieval call binding the contract method 0x385bc114.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof) constant returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof *OneStepProofSession) ExecuteStep(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _OneStepProof.Contract.ExecuteStep(&_OneStepProof.CallOpts, inboxAcc, messagesAcc, logsAcc, proof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x385bc114.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof) constant returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof *OneStepProofCallerSession) ExecuteStep(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _OneStepProof.Contract.ExecuteStep(&_OneStepProof.CallOpts, inboxAcc, messagesAcc, logsAcc, proof)
}

// ExecuteStepWithMessage is a free data retrieval call binding the contract method 0x96105dce.
//
// Solidity: function executeStepWithMessage(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, uint8 _kind, uint256 _blockNumber, uint256 _timestamp, address _sender, uint256 _inboxSeqNum, bytes _msgData) constant returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof *OneStepProofCaller) ExecuteStepWithMessage(opts *bind.CallOpts, inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, _kind uint8, _blockNumber *big.Int, _timestamp *big.Int, _sender common.Address, _inboxSeqNum *big.Int, _msgData []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	ret := new(struct {
		Gas    uint64
		Fields [5][32]byte
	})
	out := ret
	err := _OneStepProof.contract.Call(opts, out, "executeStepWithMessage", inboxAcc, messagesAcc, logsAcc, proof, _kind, _blockNumber, _timestamp, _sender, _inboxSeqNum, _msgData)
	return *ret, err
}

// ExecuteStepWithMessage is a free data retrieval call binding the contract method 0x96105dce.
//
// Solidity: function executeStepWithMessage(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, uint8 _kind, uint256 _blockNumber, uint256 _timestamp, address _sender, uint256 _inboxSeqNum, bytes _msgData) constant returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof *OneStepProofSession) ExecuteStepWithMessage(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, _kind uint8, _blockNumber *big.Int, _timestamp *big.Int, _sender common.Address, _inboxSeqNum *big.Int, _msgData []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _OneStepProof.Contract.ExecuteStepWithMessage(&_OneStepProof.CallOpts, inboxAcc, messagesAcc, logsAcc, proof, _kind, _blockNumber, _timestamp, _sender, _inboxSeqNum, _msgData)
}

// ExecuteStepWithMessage is a free data retrieval call binding the contract method 0x96105dce.
//
// Solidity: function executeStepWithMessage(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, uint8 _kind, uint256 _blockNumber, uint256 _timestamp, address _sender, uint256 _inboxSeqNum, bytes _msgData) constant returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof *OneStepProofCallerSession) ExecuteStepWithMessage(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, _kind uint8, _blockNumber *big.Int, _timestamp *big.Int, _sender common.Address, _inboxSeqNum *big.Int, _msgData []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _OneStepProof.Contract.ExecuteStepWithMessage(&_OneStepProof.CallOpts, inboxAcc, messagesAcc, logsAcc, proof, _kind, _blockNumber, _timestamp, _sender, _inboxSeqNum, _msgData)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_OneStepProof *OneStepProofFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _OneStepProof.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// PrecompilesABI is the input ABI used to generate the binding from.
const PrecompilesABI = "[]"

// PrecompilesBin is the compiled bytecode used for deploying new contracts.
var PrecompilesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582089a25ee4d0917faa75d52d337f832f58e8fef2b903c6259ee110daea7c41f21264736f6c63430005110032"

// DeployPrecompiles deploys a new Ethereum contract, binding an instance of Precompiles to it.
func DeployPrecompiles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Precompiles, error) {
	parsed, err := abi.JSON(strings.NewReader(PrecompilesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrecompilesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Precompiles{PrecompilesCaller: PrecompilesCaller{contract: contract}, PrecompilesTransactor: PrecompilesTransactor{contract: contract}, PrecompilesFilterer: PrecompilesFilterer{contract: contract}}, nil
}

// Precompiles is an auto generated Go binding around an Ethereum contract.
type Precompiles struct {
	PrecompilesCaller     // Read-only binding to the contract
	PrecompilesTransactor // Write-only binding to the contract
	PrecompilesFilterer   // Log filterer for contract events
}

// PrecompilesCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompilesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompilesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompilesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompilesSession struct {
	Contract     *Precompiles      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrecompilesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompilesCallerSession struct {
	Contract *PrecompilesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PrecompilesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompilesTransactorSession struct {
	Contract     *PrecompilesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PrecompilesRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompilesRaw struct {
	Contract *Precompiles // Generic contract binding to access the raw methods on
}

// PrecompilesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompilesCallerRaw struct {
	Contract *PrecompilesCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompilesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompilesTransactorRaw struct {
	Contract *PrecompilesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompiles creates a new instance of Precompiles, bound to a specific deployed contract.
func NewPrecompiles(address common.Address, backend bind.ContractBackend) (*Precompiles, error) {
	contract, err := bindPrecompiles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Precompiles{PrecompilesCaller: PrecompilesCaller{contract: contract}, PrecompilesTransactor: PrecompilesTransactor{contract: contract}, PrecompilesFilterer: PrecompilesFilterer{contract: contract}}, nil
}

// NewPrecompilesCaller creates a new read-only instance of Precompiles, bound to a specific deployed contract.
func NewPrecompilesCaller(address common.Address, caller bind.ContractCaller) (*PrecompilesCaller, error) {
	contract, err := bindPrecompiles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompilesCaller{contract: contract}, nil
}

// NewPrecompilesTransactor creates a new write-only instance of Precompiles, bound to a specific deployed contract.
func NewPrecompilesTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompilesTransactor, error) {
	contract, err := bindPrecompiles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTransactor{contract: contract}, nil
}

// NewPrecompilesFilterer creates a new log filterer instance of Precompiles, bound to a specific deployed contract.
func NewPrecompilesFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompilesFilterer, error) {
	contract, err := bindPrecompiles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompilesFilterer{contract: contract}, nil
}

// bindPrecompiles binds a generic wrapper to an already deployed contract.
func bindPrecompiles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrecompilesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParsePrecompilesABI parses the ABI
func ParsePrecompilesABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(PrecompilesABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Precompiles *PrecompilesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Precompiles.Contract.PrecompilesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Precompiles *PrecompilesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompiles.Contract.PrecompilesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Precompiles *PrecompilesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Precompiles.Contract.PrecompilesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Precompiles *PrecompilesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Precompiles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Precompiles *PrecompilesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompiles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Precompiles *PrecompilesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Precompiles.Contract.contract.Transact(opts, method, params...)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Precompiles *PrecompilesFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Precompiles.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}
