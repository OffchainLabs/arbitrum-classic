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
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209a81e3fc238267b7e44ca5549710241de669e8db2d08f84f90a6b17eb6424fc364736f6c63430005110032"

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
const ValueTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bytesToBytestackHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"}],\"name\":\"bytesToBytestackHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bytestackToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeMessageData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"getERCTokenMsgData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"getEthMsgData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashTestTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"innerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueSize\",\"type\":\"uint256\"}],\"name\":\"hashTuplePreImage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueTesterFuncSigs maps the 4-byte function signature to its string representation.
var ValueTesterFuncSigs = map[string]string{
	"99fb710f": "bytesToBytestackHash(bytes)",
	"b325b7d0": "bytesToBytestackHash(bytes,uint256,uint256)",
	"4d8a591a": "bytestackToBytes(bytes)",
	"98206792": "deserializeHash(bytes,uint256)",
	"2cab3a96": "deserializeMessageData(bytes,uint256)",
	"874c8778": "getERCTokenMsgData(bytes,uint256)",
	"d716661f": "getEthMsgData(bytes,uint256)",
	"364df277": "hashEmptyTuple()",
	"fd5d0c8b": "hashTestTuple()",
	"c6d25c8e": "hashTuplePreImage(bytes32,uint256)",
}

// ValueTesterBin is the compiled bytecode used for deploying new contracts.
var ValueTesterBin = "0x608060405234801561001057600080fd5b50611abc806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806399fb710f1161006657806399fb710f14610450578063b325b7d0146104f4578063c6d25c8e1461059d578063d716661f146105c0578063fd5d0c8b146106965761009e565b80632cab3a96146100a3578063364df2771461017a5780634d8a591a14610194578063874c8778146102ad578063982067921461038a575b600080fd5b610149600480360360408110156100b957600080fd5b810190602081018135600160201b8111156100d357600080fd5b8201836020820111156100e557600080fd5b803590602001918460018302840111600160201b8311171561010657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061069e915050565b6040805194151585526020850193909352838301919091526001600160a01b03166060830152519081900360800190f35b6101826106c0565b60408051918252519081900360200190f35b610238600480360360208110156101aa57600080fd5b810190602081018135600160201b8111156101c457600080fd5b8201836020820111156101d657600080fd5b803590602001918460018302840111600160201b831117156101f757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106cf945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561027257818101518382015260200161025a565b50505050905090810190601f16801561029f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610353600480360360408110156102c357600080fd5b810190602081018135600160201b8111156102dd57600080fd5b8201836020820111156102ef57600080fd5b803590602001918460018302840111600160201b8311171561031057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506106e2915050565b60408051951515865260208601949094526001600160a01b0392831685850152911660608401526080830152519081900360a00190f35b610430600480360360408110156103a057600080fd5b810190602081018135600160201b8111156103ba57600080fd5b8201836020820111156103cc57600080fd5b803590602001918460018302840111600160201b831117156103ed57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610709915050565b604080519315158452602084019290925282820152519081900360600190f35b6101826004803603602081101561046657600080fd5b810190602081018135600160201b81111561048057600080fd5b82018360208201111561049257600080fd5b803590602001918460018302840111600160201b831117156104b357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610745945050505050565b6101826004803603606081101561050a57600080fd5b810190602081018135600160201b81111561052457600080fd5b82018360208201111561053657600080fd5b803590602001918460018302840111600160201b8311171561055757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561075c565b610182600480360360408110156105b357600080fd5b5080359060200135610774565b610666600480360360408110156105d657600080fd5b810190602081018135600160201b8111156105f057600080fd5b82018360208201111561060257600080fd5b803590602001918460018302840111600160201b8311171561062357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610787915050565b60408051941515855260208501939093526001600160a01b03909116838301526060830152519081900360800190f35b610182610797565b6000806000806106ae8686610870565b93509350935093505b92959194509250565b60006106ca61092e565b905090565b60606106da8261094f565b90505b919050565b60008060008060006106f48787610a50565b945094509450945094505b9295509295909350565b6000806000806000610719611a25565b6107238888610b53565b925092509250828261073483610c90565b9550955095505050505b9250925092565b60006106da6107578360008551610d88565b610c90565b600061076c610757858585610d88565b949350505050565b60006107808383610f20565b9392505050565b6000806000806106ae8686610f5a565b60408051600280825260608281019093526000929190816020015b6107ba611a25565b8152602001906001900390816107b25790505090506107d9606f610fa7565b816000815181106107e657fe5b6020026020010181905250610835600060405190808252806020026020018201604052801561082f57816020015b61081c611a25565b8152602001906001900390816108145790505b50611033565b8160018151811061084257fe5b6020026020010181905250610855611a25565b61085e82611033565b905061086981610c90565b9250505090565b60008060008060008060008088905060008a828151811061088d57fe5b016020015160019092019160f81c9050600681146108bd5750600097508896508795508594506106b79350505050565b6108c78b8361112a565b9196509094509150846108ec5750600097508896508795508594506106b79350505050565b6108f68b8361112a565b91965090935091508461091b5750600097508896508795508594506106b79350505050565b5060019a90995091975095509350505050565b604080516000808252602082019092526109498160016111a1565b91505090565b6060600061096483600263ffffffff6111c016565b905060006020601f830104905060608160405190808252806020026020018201604052801561099d578160200160208202803883390190505b509050602360005b838110156109e6576109c0876002840163ffffffff6111c016565b8382815181106109cc57fe5b6020908102919091010152602291909101906001016109a5565b50610a466000858460405160200180828051906020019060200280838360005b83811015610a1e578181015183820152602001610a06565b505050509050019150506040516020818303038152906040526111dc9092919063ffffffff16565b9695505050505050565b6000806000806000806000806000808a905060008c8281518110610a7057fe5b016020015160019092019160f81c905060068114610aa45750600099508a98508997508796508695506106ff945050505050565b610aae8d8361112a565b919750909550915085610ad75750600099508a98508997508796508695506106ff945050505050565b610ae18d8361112a565b919750909450915085610b0a5750600099508a98508997508796508695506106ff945050505050565b610b148d8361112a565b919750909350915085610b3d5750600099508a98508997508796508695506106ff945050505050565b5060019c909b5092995090975095509350505050565b600080610b5e611a25565b84518410610b7e57600084610b736000610fa7565b92509250925061073e565b6000808590506000878281518110610b9257fe5b016020015160019092019160f81c90506000610bac611a59565b60ff8316610be057610bbe8a8561125c565b919650945091508484610bd084610fa7565b975097509750505050505061073e565b60ff831660011415610c0857610bf68a856112a1565b919650945090508484610bd08361141f565b60ff831660021415610c1e57610bd08a85611486565b600360ff841610801590610c355750600c60ff8416105b15610c705760021983016060610c4c828d8861152b565b919850965090508686610c5e83611033565b9950995099505050505050505061073e565b600080610c7d6000610fa7565b9199509750955050505050509250925092565b606081015160009060ff16610cb1578151610caa906115e9565b90506106dd565b606082015160ff1660011415610ce4576020808301518051604082015160608301519290930151610caa9391929061160d565b606082015160ff1660021415610cfd57610caa826116b5565b600360ff16826060015160ff1610158015610d2157506060820151600c60ff909116105b15610d2f57610caa8261171b565b606082015160ff1660641415610d47575080516106dd565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b610d90611a25565b602082046000610d9e61092e565b604080516002808252606080830184529394506001939260208301908038833901905050905060005b84811015610e3f578382600081518110610ddd57fe5b602002602001018181525050610e0c610757610e07836020028b018c6111c090919063ffffffff16565b610fa7565b82600181518110610e1957fe5b602002602001018181525050600283019250610e3582846111a1565b9350600101610dc7565b506020860615610ec5576000610e6189601f198a8a010163ffffffff6111c016565b90506020870660200360080281901b90508382600081518110610e8057fe5b602002602001018181525050610e9861075782610fa7565b82600181518110610ea557fe5b602002602001018181525050600283019250610ec182846111a1565b9350505b610ed161075787610fa7565b81600081518110610ede57fe5b6020026020010181815250508281600181518110610ef857fe5b602002602001018181525050600282019150610f148183611739565b98975050505050505050565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b60008060008060008060008088905060008a8281518110610f7757fe5b016020015160019092019160f81c9050600581146108bd5750600097508896508795508594506106b79350505050565b610faf611a25565b6040805160a08082018352848252825190810183526000808252602082810182905282850182905260608301829052608083018290528084019290925283518181529182018452919283019161101b565b611008611a25565b8152602001906001900390816110005790505b50815260006020820152600160409091015292915050565b61103b611a25565b6110458251611758565b611096576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156110cd578381815181106110b057fe5b60200260200101516080015182019150808060010191505061109b565b506040805160a080820183526000808352835191820184528082526020828101829052828501829052606080840183905260808085019390935290840192909252928201869052945160030160ff16948101949094528301525090565b600080600080855190508481108061114457506021858203105b806111665750600060ff1686868151811061115b57fe5b016020015160f81c14155b1561117b57506000925083915082905061073e565b6001602186016111938888840163ffffffff6111c016565b935093509350509250925092565b60006111ab611a25565b6111b58484611739565b905061076c816116b5565b600081602001835110156111d357600080fd5b50016020015190565b6060818301845110156111ee57600080fd5b60608215801561120957604051915060208201604052611253565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561124257805183526020928301920161122a565b5050858452601f01601f1916604052505b50949350505050565b600080600080855190508481108061127657506020858203105b1561128b57506000925083915082905061073e565b600160208601611193888863ffffffff6111c016565b6000806112ac611a59565b600084905060008682815181106112bf57fe5b602001015160f81c60f81b60f81c9050818060010192505060008783815181106112e557fe5b016020015160019384019360f89190911c9150600090819060ff8516141561137c576000611311611a25565b61131b8c88610b53565b9098509092509050816113675750506040805160a0810182526000808252602082018190529181018290526060810182905260808101829052909850899750955061073e945050505050565b61137081610c90565b93508060800151925050505b600061138e8b8763ffffffff6111c016565b90506020860195508460ff16600114156113df576040805160a08101825260ff90951685526020850191909152600190840181905260608401929092526080830152955091935090915061073e9050565b6040805160a08101825260ff9095168552602085019190915260009084018190526060840181905260808401525060019650929450925050509250925092565b611427611a25565b6040805160a081018252600080825260208083018690528351828152908101845291928301919061146e565b61145b611a25565b8152602001906001900390816114535790505b50815260016020820181905260409091015292915050565b600080611491611a25565b611499611a25565b85516000908190878110806114b057506040888203105b156114c857600088859650965096505050505061073e565b60006114da8a8a63ffffffff6111c016565b90506020890198506114ec8a8a61125c565b909a5094509250821561151757611503818561175f565b60019850899750955061073e945050505050565b60008986975097509750505050505061073e565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561157657816020015b611563611a25565b81526020019060019003908161155b5790505b50905060005b8960ff168160ff1610156115d3576115948985610b53565b8451859060ff86169081106115a557fe5b602090810291909101015294509250826115cb57506000955086945092506115e0915050565b60010161157c565b5060019550919350909150505b93509350939050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315611667575060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602282018590526042808301859052835180840390910181526062909201909252805191012061076c565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b606081015160009060ff1660021461170a576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b815160808301516106da9190610f20565b6000611725611a25565b61172e836117ea565b9050610780816116b5565b611741611a25565b600061174c84611860565b905061076c818461175f565b6008101590565b611767611a25565b6040805160a0808201835285825282519081018352600080825260208281018290528285018290526060830182905260808301829052808401929092528351818152918201845291928301916117d3565b6117c0611a25565b8152602001906001900390816117b85790505b508152600260208201526040019290925250919050565b6117f2611a25565b6117fb82611920565b611841576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b6060611850836040015161192f565b9050610780818460800151611739565b60006008825111156118b0576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156118f45781810151838201526020016118dc565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b60006106da8260600151611a07565b606060088251111561197f576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156119ac578160200160208202803883390190505b50805190915060005b818110156119fe5760006119db8683815181106119ce57fe5b6020026020010151610c90565b9050808483815181106119ea57fe5b6020908102919091010152506001016119b5565b50909392505050565b6000600c60ff83161080156106da575050600360ff91909116101590565b6040518060a0016040528060008152602001611a3f611a59565b815260606020820181905260006040830181905291015290565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a72315820a0d945e13eff08486f3d6ed40306495787e3c02fa1703833c4fced9a663806d464736f6c63430005110032"

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

// BytestackToBytes is a free data retrieval call binding the contract method 0x4d8a591a.
//
// Solidity: function bytestackToBytes(bytes data) pure returns(bytes)
func (_ValueTester *ValueTesterCaller) BytestackToBytes(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ValueTester.contract.Call(opts, out, "bytestackToBytes", data)
	return *ret0, err
}

// BytestackToBytes is a free data retrieval call binding the contract method 0x4d8a591a.
//
// Solidity: function bytestackToBytes(bytes data) pure returns(bytes)
func (_ValueTester *ValueTesterSession) BytestackToBytes(data []byte) ([]byte, error) {
	return _ValueTester.Contract.BytestackToBytes(&_ValueTester.CallOpts, data)
}

// BytestackToBytes is a free data retrieval call binding the contract method 0x4d8a591a.
//
// Solidity: function bytestackToBytes(bytes data) pure returns(bytes)
func (_ValueTester *ValueTesterCallerSession) BytestackToBytes(data []byte) ([]byte, error) {
	return _ValueTester.Contract.BytestackToBytes(&_ValueTester.CallOpts, data)
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

// DeserializeMessageData is a free data retrieval call binding the contract method 0x2cab3a96.
//
// Solidity: function deserializeMessageData(bytes data, uint256 startOffset) pure returns(bool, uint256, uint256, address)
func (_ValueTester *ValueTesterCaller) DeserializeMessageData(opts *bind.CallOpts, data []byte, startOffset *big.Int) (bool, *big.Int, *big.Int, common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _ValueTester.contract.Call(opts, out, "deserializeMessageData", data, startOffset)
	return *ret0, *ret1, *ret2, *ret3, err
}

// DeserializeMessageData is a free data retrieval call binding the contract method 0x2cab3a96.
//
// Solidity: function deserializeMessageData(bytes data, uint256 startOffset) pure returns(bool, uint256, uint256, address)
func (_ValueTester *ValueTesterSession) DeserializeMessageData(data []byte, startOffset *big.Int) (bool, *big.Int, *big.Int, common.Address, error) {
	return _ValueTester.Contract.DeserializeMessageData(&_ValueTester.CallOpts, data, startOffset)
}

// DeserializeMessageData is a free data retrieval call binding the contract method 0x2cab3a96.
//
// Solidity: function deserializeMessageData(bytes data, uint256 startOffset) pure returns(bool, uint256, uint256, address)
func (_ValueTester *ValueTesterCallerSession) DeserializeMessageData(data []byte, startOffset *big.Int) (bool, *big.Int, *big.Int, common.Address, error) {
	return _ValueTester.Contract.DeserializeMessageData(&_ValueTester.CallOpts, data, startOffset)
}

// GetERCTokenMsgData is a free data retrieval call binding the contract method 0x874c8778.
//
// Solidity: function getERCTokenMsgData(bytes data, uint256 startOffset) pure returns(bool, uint256, address, address, uint256)
func (_ValueTester *ValueTesterCaller) GetERCTokenMsgData(opts *bind.CallOpts, data []byte, startOffset *big.Int) (bool, *big.Int, common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(common.Address)
		ret3 = new(common.Address)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _ValueTester.contract.Call(opts, out, "getERCTokenMsgData", data, startOffset)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetERCTokenMsgData is a free data retrieval call binding the contract method 0x874c8778.
//
// Solidity: function getERCTokenMsgData(bytes data, uint256 startOffset) pure returns(bool, uint256, address, address, uint256)
func (_ValueTester *ValueTesterSession) GetERCTokenMsgData(data []byte, startOffset *big.Int) (bool, *big.Int, common.Address, common.Address, *big.Int, error) {
	return _ValueTester.Contract.GetERCTokenMsgData(&_ValueTester.CallOpts, data, startOffset)
}

// GetERCTokenMsgData is a free data retrieval call binding the contract method 0x874c8778.
//
// Solidity: function getERCTokenMsgData(bytes data, uint256 startOffset) pure returns(bool, uint256, address, address, uint256)
func (_ValueTester *ValueTesterCallerSession) GetERCTokenMsgData(data []byte, startOffset *big.Int) (bool, *big.Int, common.Address, common.Address, *big.Int, error) {
	return _ValueTester.Contract.GetERCTokenMsgData(&_ValueTester.CallOpts, data, startOffset)
}

// GetEthMsgData is a free data retrieval call binding the contract method 0xd716661f.
//
// Solidity: function getEthMsgData(bytes data, uint256 startOffset) pure returns(bool, uint256, address, uint256)
func (_ValueTester *ValueTesterCaller) GetEthMsgData(opts *bind.CallOpts, data []byte, startOffset *big.Int) (bool, *big.Int, common.Address, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _ValueTester.contract.Call(opts, out, "getEthMsgData", data, startOffset)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetEthMsgData is a free data retrieval call binding the contract method 0xd716661f.
//
// Solidity: function getEthMsgData(bytes data, uint256 startOffset) pure returns(bool, uint256, address, uint256)
func (_ValueTester *ValueTesterSession) GetEthMsgData(data []byte, startOffset *big.Int) (bool, *big.Int, common.Address, *big.Int, error) {
	return _ValueTester.Contract.GetEthMsgData(&_ValueTester.CallOpts, data, startOffset)
}

// GetEthMsgData is a free data retrieval call binding the contract method 0xd716661f.
//
// Solidity: function getEthMsgData(bytes data, uint256 startOffset) pure returns(bool, uint256, address, uint256)
func (_ValueTester *ValueTesterCallerSession) GetEthMsgData(data []byte, startOffset *big.Int) (bool, *big.Int, common.Address, *big.Int, error) {
	return _ValueTester.Contract.GetEthMsgData(&_ValueTester.CallOpts, data, startOffset)
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
