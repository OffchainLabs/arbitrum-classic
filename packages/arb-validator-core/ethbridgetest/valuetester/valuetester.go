// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package valuetester

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

// ValueTesterABI is the input ABI used to generate the binding from.
const ValueTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bytesToBytestackHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"}],\"name\":\"bytesToBytestackHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"bytestackToBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashTestTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"innerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueSize\",\"type\":\"uint256\"}],\"name\":\"hashTuplePreImage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueTesterFuncSigs maps the 4-byte function signature to its string representation.
var ValueTesterFuncSigs = map[string]string{
	"99fb710f": "bytesToBytestackHash(bytes)",
	"b325b7d0": "bytesToBytestackHash(bytes,uint256,uint256)",
	"e4d476f4": "bytestackToBytes(bytes,uint256)",
	"98206792": "deserializeHash(bytes,uint256)",
	"364df277": "hashEmptyTuple()",
	"fd5d0c8b": "hashTestTuple()",
	"c6d25c8e": "hashTuplePreImage(bytes32,uint256)",
}

// ValueTesterBin is the compiled bytecode used for deploying new contracts.
var ValueTesterBin = "0x608060405234801561001057600080fd5b506115f3806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063b325b7d01161005b578063b325b7d014610206578063c6d25c8e146102af578063e4d476f4146102d2578063fd5d0c8b146104025761007d565b8063364df27714610082578063982067921461009c57806399fb710f14610162575b600080fd5b61008a61040a565b60408051918252519081900360200190f35b610142600480360360408110156100b257600080fd5b810190602081018135600160201b8111156100cc57600080fd5b8201836020820111156100de57600080fd5b803590602001918460018302840111600160201b831117156100ff57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610419915050565b604080519315158452602084019290925282820152519081900360600190f35b61008a6004803603602081101561017857600080fd5b810190602081018135600160201b81111561019257600080fd5b8201836020820111156101a457600080fd5b803590602001918460018302840111600160201b831117156101c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610455945050505050565b61008a6004803603606081101561021c57600080fd5b810190602081018135600160201b81111561023657600080fd5b82018360208201111561024857600080fd5b803590602001918460018302840111600160201b8311171561026957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610474565b61008a600480360360408110156102c557600080fd5b508035906020013561048c565b610378600480360360408110156102e857600080fd5b810190602081018135600160201b81111561030257600080fd5b82018360208201111561031457600080fd5b803590602001918460018302840111600160201b8311171561033557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061049f915050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103c55781810151838201526020016103ad565b50505050905090810190601f1680156103f25780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61008a6104bb565b6000610414610594565b905090565b600080600080600061042961155c565b61043388886105b5565b9250925092508282610444836106f2565b9550955095505050505b9250925092565b600061046c61046783600085516107ea565b6106f2565b90505b919050565b60006104846104678585856107ea565b949350505050565b60006104988383610982565b9392505050565b60008060606104ae85856109bc565b9250925092509250925092565b60408051600280825260608281019093526000929190816020015b6104de61155c565b8152602001906001900390816104d65790505090506104fd606f610ade565b8160008151811061050a57fe5b6020026020010181905250610559600060405190808252806020026020018201604052801561055357816020015b61054061155c565b8152602001906001900390816105385790505b50610b6a565b8160018151811061056657fe5b602002602001018190525061057961155c565b61058282610b6a565b905061058d816106f2565b9250505090565b604080516000808252602082019092526105af816001610c61565b91505090565b6000806105c061155c565b845184106105e0576000846105d56000610ade565b92509250925061044e565b60008085905060008782815181106105f457fe5b016020015160019092019160f81c9050600061060e611590565b60ff8316610642576106208a85610c80565b91965094509150848461063284610ade565b975097509750505050505061044e565b60ff83166001141561066a576106588a85610cd3565b91965094509050848461063283610e51565b60ff831660021415610680576106328a85610eb8565b600360ff8416108015906106975750600c60ff8416105b156106d257600219830160606106ae828d88610f5d565b9198509650905086866106c083610b6a565b9950995099505050505050505061044e565b6000806106df6000610ade565b9199509750955050505050509250925092565b606081015160009060ff1661071357815161070c9061101b565b905061046f565b606082015160ff166001141561074657602080830151805160408201516060830151929093015161070c9391929061103f565b606082015160ff166002141561075f5761070c826110e7565b600360ff16826060015160ff161015801561078357506060820151600c60ff909116105b156107915761070c8261114d565b606082015160ff16606414156107a95750805161046f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6107f261155c565b602082046000610800610594565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b848110156108a157838260008151811061083f57fe5b60200260200101818152505061086e610467610869836020028b018c61116b90919063ffffffff16565b610ade565b8260018151811061087b57fe5b6020026020010181815250506002830192506108978284610c61565b9350600101610829565b5060208606156109275760006108c389601f198a8a010163ffffffff61116b16565b90506020870660200360080281901b905083826000815181106108e257fe5b6020026020010181815250506108fa61046782610ade565b8260018151811061090757fe5b6020026020010181815250506002830192506109238284610c61565b9350505b61093361046787610ade565b8160008151811061094057fe5b602002602001018181525050828160018151811061095a57fe5b6020026020010181815250506002820191506109768183611187565b98975050505050505050565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b6000600182016060826109cf86846111a6565b91955093509050836109e557506000925061044e565b60006020601f8301049050606081604051908082528060200260200182016040528015610a1c578160200160208202803883390190505b50905060005b82811015610a7b576001860195506000610a3c8a886111a6565b9199509750905087610a5757506000965061044e9350505050565b8060001b838381518110610a6757fe5b602090810291909101015250600101610a22565b506001856104446000868560405160200180828051906020019060200280838360005b83811015610ab6578181015183820152602001610a9e565b5050505090500191505060405160208183030381529060405261120f9092919063ffffffff16565b610ae661155c565b6040805160a080820183528482528251908101835260008082526020828101829052828501829052606083018290526080830182905280840192909252835181815291820184529192830191610b52565b610b3f61155c565b815260200190600190039081610b375790505b50815260006020820152600160409091015292915050565b610b7261155c565b610b7c825161128f565b610bcd576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015610c0457838181518110610be757fe5b602002602001015160800151820191508080600101915050610bd2565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b6000610c6b61155c565b610c758484611187565b9050610484816110e7565b6000806000808551905084811080610c9a57506020858203105b15610caf57506000925083915082905061044e565b600160208601610cc5888863ffffffff61116b16565b935093509350509250925092565b600080610cde611590565b60008490506000868281518110610cf157fe5b602001015160f81c60f81b60f81c905081806001019250506000878381518110610d1757fe5b016020015160019384019360f89190911c9150600090819060ff85161415610dae576000610d4361155c565b610d4d8c886105b5565b909850909250905081610d995750506040805160a0810182526000808252602082018190529181018290526060810182905260808101829052909850899750955061044e945050505050565b610da2816106f2565b93508060800151925050505b6000610dc08b8763ffffffff61116b16565b90506020860195508460ff1660011415610e11576040805160a08101825260ff90951685526020850191909152600190840181905260608401929092526080830152955091935090915061044e9050565b6040805160a08101825260ff9095168552602085019190915260009084018190526060840181905260808401525060019650929450925050509250925092565b610e5961155c565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190610ea0565b610e8d61155c565b815260200190600190039081610e855790505b50815260016020820181905260409091015292915050565b600080610ec361155c565b610ecb61155c565b8551600090819087811080610ee257506040888203105b15610efa57600088859650965096505050505061044e565b6000610f0c8a8a63ffffffff61116b16565b9050602089019850610f1e8a8a610c80565b909a50945092508215610f4957610f358185611296565b60019850899750955061044e945050505050565b60008986975097509750505050505061044e565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610fa857816020015b610f9561155c565b815260200190600190039081610f8d5790505b50905060005b8960ff168160ff16101561100557610fc689856105b5565b8451859060ff8616908110610fd757fe5b60209081029190910101529450925082610ffd5750600095508694509250611012915050565b600101610fae565b5060019550919350909150505b93509350939050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315611099575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120610484565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff1660021461113c576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b8151608083015161046c9190610982565b600061115761155c565b61116083611321565b9050610498816110e7565b6000816020018351101561117e57600080fd5b50016020015190565b61118f61155c565b600061119a84611397565b90506104848184611296565b60008060008085519050848110806111c057506021858203105b806111e25750600060ff168686815181106111d757fe5b016020015160f81c14155b156111f757506000925083915082905061044e565b600160218601610cc58888840163ffffffff61116b16565b60608183018451101561122157600080fd5b60608215801561123c57604051915060208201604052611286565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561127557805183526020928301920161125d565b5050858452601f01601f1916604052505b50949350505050565b6008101590565b61129e61155c565b6040805160a08082018352858252825190810183526000808252602082810182905282850182905260608301829052608083018290528084019290925283518181529182018452919283019161130a565b6112f761155c565b8152602001906001900390816112ef5790505b508152600260208201526040019290925250919050565b61132961155c565b61133282611457565b611378576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b60606113878360400151611466565b9050610498818460800151611187565b60006008825111156113e7576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561142b578181015183820152602001611413565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b600061046c826060015161153e565b60606008825111156114b6576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156114e3578160200160208202803883390190505b50805190915060005b8181101561153557600061151286838151811061150557fe5b60200260200101516106f2565b90508084838151811061152157fe5b6020908102919091010152506001016114ec565b50909392505050565b6000600c60ff831610801561046c575050600360ff91909116101590565b6040518060a0016040528060008152602001611576611590565b815260606020820181905260006040830181905291015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a72315820415a1101036a6c7753670b6d5e967c63e1022f62b131a3cf003a1ce4e62ad5c064736f6c63430005110032"

// DeployValueTester deploys a new Ethereum contract, binding an instance of ValueTester to it.
func DeployValueTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValueTester, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValueTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValueTester{ValueTesterCaller: ValueTesterCaller{contract: contract}, ValueTesterTransactor: ValueTesterTransactor{contract: contract}, ValueTesterFilterer: ValueTesterFilterer{contract: contract}}, nil
}

// ValueTester is an auto generated Go binding around an Ethereum contract.
type ValueTester struct {
	ValueTesterCaller     // Read-only binding to the contract
	ValueTesterTransactor // Write-only binding to the contract
	ValueTesterFilterer   // Log filterer for contract events
}

// ValueTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueTesterSession struct {
	Contract     *ValueTester      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueTesterCallerSession struct {
	Contract *ValueTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ValueTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueTesterTransactorSession struct {
	Contract     *ValueTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ValueTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueTesterRaw struct {
	Contract *ValueTester // Generic contract binding to access the raw methods on
}

// ValueTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueTesterCallerRaw struct {
	Contract *ValueTesterCaller // Generic read-only contract binding to access the raw methods on
}

// ValueTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueTesterTransactorRaw struct {
	Contract *ValueTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValueTester creates a new instance of ValueTester, bound to a specific deployed contract.
func NewValueTester(address common.Address, backend bind.ContractBackend) (*ValueTester, error) {
	contract, err := bindValueTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValueTester{ValueTesterCaller: ValueTesterCaller{contract: contract}, ValueTesterTransactor: ValueTesterTransactor{contract: contract}, ValueTesterFilterer: ValueTesterFilterer{contract: contract}}, nil
}

// NewValueTesterCaller creates a new read-only instance of ValueTester, bound to a specific deployed contract.
func NewValueTesterCaller(address common.Address, caller bind.ContractCaller) (*ValueTesterCaller, error) {
	contract, err := bindValueTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTesterCaller{contract: contract}, nil
}

// NewValueTesterTransactor creates a new write-only instance of ValueTester, bound to a specific deployed contract.
func NewValueTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueTesterTransactor, error) {
	contract, err := bindValueTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTesterTransactor{contract: contract}, nil
}

// NewValueTesterFilterer creates a new log filterer instance of ValueTester, bound to a specific deployed contract.
func NewValueTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueTesterFilterer, error) {
	contract, err := bindValueTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueTesterFilterer{contract: contract}, nil
}

// bindValueTester binds a generic wrapper to an already deployed contract.
func bindValueTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueTester *ValueTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValueTester.Contract.ValueTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueTester *ValueTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueTester.Contract.ValueTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueTester *ValueTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueTester.Contract.ValueTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueTester *ValueTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValueTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueTester *ValueTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueTester *ValueTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueTester.Contract.contract.Transact(opts, method, params...)
}

// BytesToBytestackHash is a free data retrieval call binding the contract method 0x99fb710f.
//
// Solidity: function bytesToBytestackHash(bytes data) pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) BytesToBytestackHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "bytesToBytestackHash", data)
	return *ret0, err
}

// BytesToBytestackHash is a free data retrieval call binding the contract method 0x99fb710f.
//
// Solidity: function bytesToBytestackHash(bytes data) pure returns(bytes32)
func (_ValueTester *ValueTesterSession) BytesToBytestackHash(data []byte) ([32]byte, error) {
	return _ValueTester.Contract.BytesToBytestackHash(&_ValueTester.CallOpts, data)
}

// BytesToBytestackHash is a free data retrieval call binding the contract method 0x99fb710f.
//
// Solidity: function bytesToBytestackHash(bytes data) pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) BytesToBytestackHash(data []byte) ([32]byte, error) {
	return _ValueTester.Contract.BytesToBytestackHash(&_ValueTester.CallOpts, data)
}

// BytesToBytestackHash0 is a free data retrieval call binding the contract method 0xb325b7d0.
//
// Solidity: function bytesToBytestackHash(bytes data, uint256 startOffset, uint256 dataLength) pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) BytesToBytestackHash0(opts *bind.CallOpts, data []byte, startOffset *big.Int, dataLength *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "bytesToBytestackHash0", data, startOffset, dataLength)
	return *ret0, err
}

// BytesToBytestackHash0 is a free data retrieval call binding the contract method 0xb325b7d0.
//
// Solidity: function bytesToBytestackHash(bytes data, uint256 startOffset, uint256 dataLength) pure returns(bytes32)
func (_ValueTester *ValueTesterSession) BytesToBytestackHash0(data []byte, startOffset *big.Int, dataLength *big.Int) ([32]byte, error) {
	return _ValueTester.Contract.BytesToBytestackHash0(&_ValueTester.CallOpts, data, startOffset, dataLength)
}

// BytesToBytestackHash0 is a free data retrieval call binding the contract method 0xb325b7d0.
//
// Solidity: function bytesToBytestackHash(bytes data, uint256 startOffset, uint256 dataLength) pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) BytesToBytestackHash0(data []byte, startOffset *big.Int, dataLength *big.Int) ([32]byte, error) {
	return _ValueTester.Contract.BytesToBytestackHash0(&_ValueTester.CallOpts, data, startOffset, dataLength)
}

// BytestackToBytes is a free data retrieval call binding the contract method 0xe4d476f4.
//
// Solidity: function bytestackToBytes(bytes data, uint256 offset) pure returns(bool, uint256, bytes)
func (_ValueTester *ValueTesterCaller) BytestackToBytes(opts *bind.CallOpts, data []byte, offset *big.Int) (bool, *big.Int, []byte, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ValueTester.contract.Call(opts, out, "bytestackToBytes", data, offset)
	return *ret0, *ret1, *ret2, err
}

// BytestackToBytes is a free data retrieval call binding the contract method 0xe4d476f4.
//
// Solidity: function bytestackToBytes(bytes data, uint256 offset) pure returns(bool, uint256, bytes)
func (_ValueTester *ValueTesterSession) BytestackToBytes(data []byte, offset *big.Int) (bool, *big.Int, []byte, error) {
	return _ValueTester.Contract.BytestackToBytes(&_ValueTester.CallOpts, data, offset)
}

// BytestackToBytes is a free data retrieval call binding the contract method 0xe4d476f4.
//
// Solidity: function bytestackToBytes(bytes data, uint256 offset) pure returns(bool, uint256, bytes)
func (_ValueTester *ValueTesterCallerSession) BytestackToBytes(data []byte, offset *big.Int) (bool, *big.Int, []byte, error) {
	return _ValueTester.Contract.BytestackToBytes(&_ValueTester.CallOpts, data, offset)
}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(bool, uint256, bytes32)
func (_ValueTester *ValueTesterCaller) DeserializeHash(opts *bind.CallOpts, data []byte, startOffset *big.Int) (bool, *big.Int, [32]byte, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ValueTester.contract.Call(opts, out, "deserializeHash", data, startOffset)
	return *ret0, *ret1, *ret2, err
}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(bool, uint256, bytes32)
func (_ValueTester *ValueTesterSession) DeserializeHash(data []byte, startOffset *big.Int) (bool, *big.Int, [32]byte, error) {
	return _ValueTester.Contract.DeserializeHash(&_ValueTester.CallOpts, data, startOffset)
}

// DeserializeHash is a free data retrieval call binding the contract method 0x98206792.
//
// Solidity: function deserializeHash(bytes data, uint256 startOffset) pure returns(bool, uint256, bytes32)
func (_ValueTester *ValueTesterCallerSession) DeserializeHash(data []byte, startOffset *big.Int) (bool, *big.Int, [32]byte, error) {
	return _ValueTester.Contract.DeserializeHash(&_ValueTester.CallOpts, data, startOffset)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterSession) HashEmptyTuple() ([32]byte, error) {
	return _ValueTester.Contract.HashEmptyTuple(&_ValueTester.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _ValueTester.Contract.HashEmptyTuple(&_ValueTester.CallOpts)
}

// HashTestTuple is a free data retrieval call binding the contract method 0xfd5d0c8b.
//
// Solidity: function hashTestTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) HashTestTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "hashTestTuple")
	return *ret0, err
}

// HashTestTuple is a free data retrieval call binding the contract method 0xfd5d0c8b.
//
// Solidity: function hashTestTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterSession) HashTestTuple() ([32]byte, error) {
	return _ValueTester.Contract.HashTestTuple(&_ValueTester.CallOpts)
}

// HashTestTuple is a free data retrieval call binding the contract method 0xfd5d0c8b.
//
// Solidity: function hashTestTuple() pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) HashTestTuple() ([32]byte, error) {
	return _ValueTester.Contract.HashTestTuple(&_ValueTester.CallOpts)
}

// HashTuplePreImage is a free data retrieval call binding the contract method 0xc6d25c8e.
//
// Solidity: function hashTuplePreImage(bytes32 innerHash, uint256 valueSize) pure returns(bytes32)
func (_ValueTester *ValueTesterCaller) HashTuplePreImage(opts *bind.CallOpts, innerHash [32]byte, valueSize *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "hashTuplePreImage", innerHash, valueSize)
	return *ret0, err
}

// HashTuplePreImage is a free data retrieval call binding the contract method 0xc6d25c8e.
//
// Solidity: function hashTuplePreImage(bytes32 innerHash, uint256 valueSize) pure returns(bytes32)
func (_ValueTester *ValueTesterSession) HashTuplePreImage(innerHash [32]byte, valueSize *big.Int) ([32]byte, error) {
	return _ValueTester.Contract.HashTuplePreImage(&_ValueTester.CallOpts, innerHash, valueSize)
}

// HashTuplePreImage is a free data retrieval call binding the contract method 0xc6d25c8e.
//
// Solidity: function hashTuplePreImage(bytes32 innerHash, uint256 valueSize) pure returns(bytes32)
func (_ValueTester *ValueTesterCallerSession) HashTuplePreImage(innerHash [32]byte, valueSize *big.Int) ([32]byte, error) {
	return _ValueTester.Contract.HashTuplePreImage(&_ValueTester.CallOpts, innerHash, valueSize)
}
