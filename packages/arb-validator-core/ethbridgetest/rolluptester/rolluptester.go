// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rolluptester

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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820557675814df39add0fa761bb6fc1ca2754618e3660ac97035c79af694b750a9264736f6c634300050d0032"

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

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820da8d11bf1947a4571c85676ea8b846317a3049a34118d9ab59be573dd308ff8e64736f6c634300050d0032"

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

// RollupTesterABI is the input ABI used to generate the binding from.
const RollupTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"confNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initalProtoStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"branches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deadlineTicks\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengeNodeData\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"logsAcc\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"vmProtoStateHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"}],\"name\":\"confirm\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"validNodeHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"lastNode\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"initalProtoStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"branches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deadlineTicks\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengeNodeData\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"logsAcc\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"vmProtoStateHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"processValidNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// RollupTesterFuncSigs maps the 4-byte function signature to its string representation.
var RollupTesterFuncSigs = map[string]string{
	"bd912e8e": "confirm(bytes32,bytes32,uint256[],uint256[],bytes32[],bytes32[],bytes32[],uint256[],bytes)",
	"02be0bd0": "generateLastMessageHash(bytes,uint256,uint256)",
	"caf32e44": "processValidNode(bytes32,uint256[],uint256[],bytes32[],bytes32[],bytes32[],uint256[],bytes,uint256,uint256)",
}

// RollupTesterBin is the compiled bytecode used for deploying new contracts.
var RollupTesterBin = "0x608060405234801561001057600080fd5b5061198e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806302be0bd014610046578063bd912e8e14610108578063caf32e4414610520575b600080fd5b6100ef6004803603606081101561005c57600080fd5b810190602081018135600160201b81111561007657600080fd5b82018360208201111561008857600080fd5b803590602001918460018302840111600160201b831117156100a957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356108fb565b6040805192835260208301919091528051918290030190f35b6104c5600480360361012081101561011f57600080fd5b813591602081013591810190606081016040820135600160201b81111561014557600080fd5b82018360208201111561015757600080fd5b803590602001918460208302840111600160201b8311171561017857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101c757600080fd5b8201836020820111156101d957600080fd5b803590602001918460208302840111600160201b831117156101fa57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561024957600080fd5b82018360208201111561025b57600080fd5b803590602001918460208302840111600160201b8311171561027c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102cb57600080fd5b8201836020820111156102dd57600080fd5b803590602001918460208302840111600160201b831117156102fe57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561034d57600080fd5b82018360208201111561035f57600080fd5b803590602001918460208302840111600160201b8311171561038057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103cf57600080fd5b8201836020820111156103e157600080fd5b803590602001918460208302840111600160201b8311171561040257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561045157600080fd5b82018360208201111561046357600080fd5b803590602001918460018302840111600160201b8311171561048457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610915945050505050565b6040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b8381101561050b5781810151838201526020016104f3565b50505050905001935050505060405180910390f35b6108dd600480360361014081101561053757600080fd5b81359190810190604081016020820135600160201b81111561055857600080fd5b82018360208201111561056a57600080fd5b803590602001918460208302840111600160201b8311171561058b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156105da57600080fd5b8201836020820111156105ec57600080fd5b803590602001918460208302840111600160201b8311171561060d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561065c57600080fd5b82018360208201111561066e57600080fd5b803590602001918460208302840111600160201b8311171561068f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106de57600080fd5b8201836020820111156106f057600080fd5b803590602001918460208302840111600160201b8311171561071157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561076057600080fd5b82018360208201111561077257600080fd5b803590602001918460208302840111600160201b8311171561079357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107e257600080fd5b8201836020820111156107f457600080fd5b803590602001918460208302840111600160201b8311171561081557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561086457600080fd5b82018360208201111561087657600080fd5b803590602001918460018302840111600160201b8311171561089757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561096e565b60408051938452602084019290925282820152519081900360600190f35b6000806109098585856109cd565b91509150935093915050565b6060600061095c6040518061010001604052808c81526020018b81526020018a8152602001898152602001888152602001878152602001868152602001858152508c610a90565b91509150995099975050505050505050565b60008060006109b76040518061010001604052808f81526020018e81526020018d81526020018c81526020018b81526020018a8152602001898152602001888152508686610bb2565b9250925092509a509a509a975050505050505050565b60008080806109da6118fe565b8660005b87811015610a80576109f08a83610c3c565b919650909350915084610a43576040805162461bcd60e51b8152602060048201526016602482015275496e76616c6964206f7574707574206d65737361676560501b604482015290519081900360640190fd5b83610a4d84610d7a565b604080516020808201949094528082019290925280518083038201815260609092019052805191012093506001016109de565b5091989197509095505050505050565b602082015151606090600090610aa585610e8d565b60408051828152602080840282010190915260009081908190848015610ad5578160200160208202803883390190505b50885190965060005b85811015610ba45760008a602001518281518110610af857fe5b6020026020010151905060006003821415610b2b57610b188c8887610bb2565b6001909901989196509094509050610b4e565b8b606001518681518110610b3b57fe5b6020026020010151905085806001019650505b610b728b8d604001518581518110610b6257fe5b6020026020010151838588610ee4565b9a506003821415610b9a578a8a6001890381518110610b8d57fe5b6020026020010181815250505b5050600101610ade565b509598969750505050505050565b6000806000806000610be08860e00151878a60c001518a81518110610bd357fe5b60200260200101516109cd565b915091506000610c07838a608001518a81518110610bfa57fe5b6020026020010151610f4c565b905060008960a001518981518110610c1b57fe5b60200260200101519050828282965096509650505050505b93509350939050565b600080610c476118fe565b84518410610c6757600084610c5c6000610f78565b925092509250610d73565b6000808590506000878281518110610c7b57fe5b016020015160019092019160f81c90506000610c95611932565b60ff8316610cc957610ca78a85610ffd565b919650945091508484610cb984610f78565b9750975097505050505050610d73565b60ff831660011415610cf157610cdf8a85611050565b919650945090508484610cb9836111b0565b60ff831660021415610d0757610cb98a85611217565b600360ff841610801590610d1e5750600c60ff8416105b15610d595760021983016060610d35828d886112bc565b919850965090508686610d4783611376565b99509950995050505050505050610d73565b600080610d666000610f78565b9199509750955050505050505b9250925092565b6000600360090160ff16826060015160ff1610610dd2576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610df0578151610de990611465565b9050610e88565b606082015160ff1660011415610e23576020808301518051604082015160608301519290930151610de993919290611489565b606082015160ff1660021415610e5457602080830151015160011415610e4b57508051610e88565b610de982611532565b600360ff16826060015160ff1610158015610e7857506060820151600c60ff909116105b15610e8657610de98261159e565bfe5b919050565b60208101515160c08201515160a0830151518114610eaa57600080fd5b8083608001515114610ebb57600080fd5b8183604001515114610ecc57600080fd5b80820383606001515114610edf57600080fd5b505050565b6040805160208082018490528183018790526060820186905260808083018690528351808403909101815260a08301845280519082012060c0830189905260e08084019190915283518084039091018152610100909201909252805191012095945050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b610f806118fe565b6040805160a0810182528381528151608081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610fe5565b610fd26118fe565b815260200190600190039081610fca5790505b50815260006020820152600160409091015292915050565b600080600080855190508481108061101757506020858203105b1561102c575060009250839150829050610d73565b600160208601611042888863ffffffff6115c316565b935093509350509250925092565b60008061105b611932565b6000849050600086828151811061106e57fe5b602001015160f81c60f81b60f81c90508180600101925050600087838151811061109457fe5b016020015160019384019360f89190911c915060009060ff8416141561111a5760006110be6118fe565b6110c88b87610c3c565b90975090925090508161110c57505060408051608081018252600080825260208201819052918101829052606081018290529097508896509450610d739350505050565b61111581610d7a565b925050505b600061112c8a8663ffffffff6115c316565b90506020850194508360ff1660011415611178576040805160808101825260ff90941684526020840191909152600190830181905260608301919091529550919350909150610d739050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b6111b86118fe565b6040805160a08101825260008082526020808301869052835182815290810184529192830191906111ff565b6111ec6118fe565b8152602001906001900390816111e45790505b50815260016020820181905260409091015292915050565b6000806112226118fe565b61122a6118fe565b855160009081908781108061124157506040888203105b15611259576000888596509650965050505050610d73565b600061126b8a8a63ffffffff6115c316565b905060208901985061127d8a8a610ffd565b909a509450925082156112a85761129481856115df565b600198508997509550610d73945050505050565b600089869750975097505050505050610d73565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561130757816020015b6112f46118fe565b8152602001906001900390816112ec5790505b50905060005b8960ff168160ff161015611364576113258985610c3c565b8451859060ff861690811061133657fe5b6020908102919091010152945092508261135c5750600095508694509250610c33915050565b60010161130d565b50600199929850965090945050505050565b61137e6118fe565b6113888251611663565b6113d9576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b8351811015611410578381815181106113f357fe5b6020026020010151608001518201915080806001019150506113de565b506040805160a08101825260008082528251608080820185528282526020808301849052828601849052606080840194909452840191909152928201869052945160030160ff16948101949094528301525090565b60408051602080820193909352815180820384018152908201909152805191012090565b600083156114e3575060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602282018590526042808301859052835180840390910181526062909201909252805191012061152a565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b606081015160009060ff16600214611587576040805162461bcd60e51b815260206004820152601360248201527209aeae6e840c4ca40a0e4ca92dac2ceca90e6d606b1b604482015290519081900360640190fd5b81516080830151611598919061166a565b92915050565b60006115a86118fe565b6115b1836116a4565b90506115bc81611532565b9392505050565b600081602001835110156115d657600080fd5b50016020015190565b6115e76118fe565b6040805160a081018252848152815160808101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161164c565b6116396118fe565b8152602001906001900390816116315790505b508152600260208201526040019290925250919050565b6008101590565b60408051600360f81b6020808301919091526021820194909452604180820193909352815180820390930183526061019052805191012090565b6116ac6118fe565b6116b58261171a565b6116fb576040805162461bcd60e51b81526020600482015260126024820152714d757374206265205475706c65207479706560701b604482015290519081900360640190fd5b606061170a8360400151611729565b90506115bc818460800151611801565b60006115988260600151611820565b6060600882511115611779576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156117a6578160200160208202803883390190505b50805190915060005b818110156117f85760006117d58683815181106117c857fe5b6020026020010151610d7a565b9050808483815181106117e457fe5b6020908102919091010152506001016117af565b50909392505050565b6118096118fe565b60006118148461183e565b905061152a81846115df565b6000600c60ff8316108015611598575050600360ff91909116101590565b600060088251111561188e576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6000825183604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156118d25781810151838201526020016118ba565b505050509050019250505060405160208183030381529060405280519060200120905080915050919050565b6040518060a0016040528060008152602001611918611932565b815260606020820181905260006040830181905291015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a72315820742db3878b440ce75629c45067263ee2ba0541325fcc5499e0ca427074d5af2064736f6c634300050d0032"

// DeployRollupTester deploys a new Ethereum contract, binding an instance of RollupTester to it.
func DeployRollupTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupTester, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupTester{RollupTesterCaller: RollupTesterCaller{contract: contract}, RollupTesterTransactor: RollupTesterTransactor{contract: contract}, RollupTesterFilterer: RollupTesterFilterer{contract: contract}}, nil
}

// RollupTester is an auto generated Go binding around an Ethereum contract.
type RollupTester struct {
	RollupTesterCaller     // Read-only binding to the contract
	RollupTesterTransactor // Write-only binding to the contract
	RollupTesterFilterer   // Log filterer for contract events
}

// RollupTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupTesterSession struct {
	Contract     *RollupTester     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupTesterCallerSession struct {
	Contract *RollupTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RollupTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTesterTransactorSession struct {
	Contract     *RollupTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RollupTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupTesterRaw struct {
	Contract *RollupTester // Generic contract binding to access the raw methods on
}

// RollupTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupTesterCallerRaw struct {
	Contract *RollupTesterCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTesterTransactorRaw struct {
	Contract *RollupTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupTester creates a new instance of RollupTester, bound to a specific deployed contract.
func NewRollupTester(address common.Address, backend bind.ContractBackend) (*RollupTester, error) {
	contract, err := bindRollupTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupTester{RollupTesterCaller: RollupTesterCaller{contract: contract}, RollupTesterTransactor: RollupTesterTransactor{contract: contract}, RollupTesterFilterer: RollupTesterFilterer{contract: contract}}, nil
}

// NewRollupTesterCaller creates a new read-only instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterCaller(address common.Address, caller bind.ContractCaller) (*RollupTesterCaller, error) {
	contract, err := bindRollupTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTesterCaller{contract: contract}, nil
}

// NewRollupTesterTransactor creates a new write-only instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTesterTransactor, error) {
	contract, err := bindRollupTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTesterTransactor{contract: contract}, nil
}

// NewRollupTesterFilterer creates a new log filterer instance of RollupTester, bound to a specific deployed contract.
func NewRollupTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupTesterFilterer, error) {
	contract, err := bindRollupTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupTesterFilterer{contract: contract}, nil
}

// bindRollupTester binds a generic wrapper to an already deployed contract.
func bindRollupTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTester *RollupTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTester.Contract.RollupTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTester *RollupTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTester.Contract.RollupTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTester *RollupTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTester.Contract.RollupTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTester *RollupTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTester *RollupTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTester *RollupTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTester.Contract.contract.Transact(opts, method, params...)
}

// Confirm is a free data retrieval call binding the contract method 0xbd912e8e.
//
// Solidity: function confirm(bytes32 confNode, bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages) pure returns(bytes32[] validNodeHashes, bytes32 lastNode)
func (_RollupTester *RollupTesterCaller) Confirm(opts *bind.CallOpts, confNode [32]byte, initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte) (struct {
	ValidNodeHashes [][32]byte
	LastNode        [32]byte
}, error) {
	ret := new(struct {
		ValidNodeHashes [][32]byte
		LastNode        [32]byte
	})
	out := ret
	err := _RollupTester.contract.Call(opts, out, "confirm", confNode, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages)
	return *ret, err
}

// Confirm is a free data retrieval call binding the contract method 0xbd912e8e.
//
// Solidity: function confirm(bytes32 confNode, bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages) pure returns(bytes32[] validNodeHashes, bytes32 lastNode)
func (_RollupTester *RollupTesterSession) Confirm(confNode [32]byte, initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte) (struct {
	ValidNodeHashes [][32]byte
	LastNode        [32]byte
}, error) {
	return _RollupTester.Contract.Confirm(&_RollupTester.CallOpts, confNode, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages)
}

// Confirm is a free data retrieval call binding the contract method 0xbd912e8e.
//
// Solidity: function confirm(bytes32 confNode, bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages) pure returns(bytes32[] validNodeHashes, bytes32 lastNode)
func (_RollupTester *RollupTesterCallerSession) Confirm(confNode [32]byte, initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte) (struct {
	ValidNodeHashes [][32]byte
	LastNode        [32]byte
}, error) {
	return _RollupTester.Contract.Confirm(&_RollupTester.CallOpts, confNode, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) pure returns(bytes32, uint256)
func (_RollupTester *RollupTesterCaller) GenerateLastMessageHash(opts *bind.CallOpts, messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RollupTester.contract.Call(opts, out, "generateLastMessageHash", messages, startOffset, length)
	return *ret0, *ret1, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) pure returns(bytes32, uint256)
func (_RollupTester *RollupTesterSession) GenerateLastMessageHash(messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	return _RollupTester.Contract.GenerateLastMessageHash(&_RollupTester.CallOpts, messages, startOffset, length)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x02be0bd0.
//
// Solidity: function generateLastMessageHash(bytes messages, uint256 startOffset, uint256 length) pure returns(bytes32, uint256)
func (_RollupTester *RollupTesterCallerSession) GenerateLastMessageHash(messages []byte, startOffset *big.Int, length *big.Int) ([32]byte, *big.Int, error) {
	return _RollupTester.Contract.GenerateLastMessageHash(&_RollupTester.CallOpts, messages, startOffset, length)
}

// ProcessValidNode is a free data retrieval call binding the contract method 0xcaf32e44.
//
// Solidity: function processValidNode(bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, uint256 validNum, uint256 startOffset) pure returns(uint256, bytes32, bytes32)
func (_RollupTester *RollupTesterCaller) ProcessValidNode(opts *bind.CallOpts, initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, validNum *big.Int, startOffset *big.Int) (*big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _RollupTester.contract.Call(opts, out, "processValidNode", initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, validNum, startOffset)
	return *ret0, *ret1, *ret2, err
}

// ProcessValidNode is a free data retrieval call binding the contract method 0xcaf32e44.
//
// Solidity: function processValidNode(bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, uint256 validNum, uint256 startOffset) pure returns(uint256, bytes32, bytes32)
func (_RollupTester *RollupTesterSession) ProcessValidNode(initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, validNum *big.Int, startOffset *big.Int) (*big.Int, [32]byte, [32]byte, error) {
	return _RollupTester.Contract.ProcessValidNode(&_RollupTester.CallOpts, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, validNum, startOffset)
}

// ProcessValidNode is a free data retrieval call binding the contract method 0xcaf32e44.
//
// Solidity: function processValidNode(bytes32 initalProtoStateHash, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, uint256 validNum, uint256 startOffset) pure returns(uint256, bytes32, bytes32)
func (_RollupTester *RollupTesterCallerSession) ProcessValidNode(initalProtoStateHash [32]byte, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, validNum *big.Int, startOffset *big.Int) (*big.Int, [32]byte, [32]byte, error) {
	return _RollupTester.Contract.ProcessValidNode(&_RollupTester.CallOpts, initalProtoStateHash, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, validNum, startOffset)
}

// RollupTimeABI is the input ABI used to generate the binding from.
const RollupTimeABI = "[]"

// RollupTimeBin is the compiled bytecode used for deploying new contracts.
var RollupTimeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582037f06363d238a70c34530f76a5f2c329256ab1414b55b90256aa470c78073d0964736f6c634300050d0032"

// DeployRollupTime deploys a new Ethereum contract, binding an instance of RollupTime to it.
func DeployRollupTime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupTime, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupTimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupTime{RollupTimeCaller: RollupTimeCaller{contract: contract}, RollupTimeTransactor: RollupTimeTransactor{contract: contract}, RollupTimeFilterer: RollupTimeFilterer{contract: contract}}, nil
}

// RollupTime is an auto generated Go binding around an Ethereum contract.
type RollupTime struct {
	RollupTimeCaller     // Read-only binding to the contract
	RollupTimeTransactor // Write-only binding to the contract
	RollupTimeFilterer   // Log filterer for contract events
}

// RollupTimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupTimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupTimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupTimeSession struct {
	Contract     *RollupTime       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupTimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupTimeCallerSession struct {
	Contract *RollupTimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupTimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTimeTransactorSession struct {
	Contract     *RollupTimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupTimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupTimeRaw struct {
	Contract *RollupTime // Generic contract binding to access the raw methods on
}

// RollupTimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupTimeCallerRaw struct {
	Contract *RollupTimeCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTimeTransactorRaw struct {
	Contract *RollupTimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupTime creates a new instance of RollupTime, bound to a specific deployed contract.
func NewRollupTime(address common.Address, backend bind.ContractBackend) (*RollupTime, error) {
	contract, err := bindRollupTime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupTime{RollupTimeCaller: RollupTimeCaller{contract: contract}, RollupTimeTransactor: RollupTimeTransactor{contract: contract}, RollupTimeFilterer: RollupTimeFilterer{contract: contract}}, nil
}

// NewRollupTimeCaller creates a new read-only instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeCaller(address common.Address, caller bind.ContractCaller) (*RollupTimeCaller, error) {
	contract, err := bindRollupTime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTimeCaller{contract: contract}, nil
}

// NewRollupTimeTransactor creates a new write-only instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTimeTransactor, error) {
	contract, err := bindRollupTime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTimeTransactor{contract: contract}, nil
}

// NewRollupTimeFilterer creates a new log filterer instance of RollupTime, bound to a specific deployed contract.
func NewRollupTimeFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupTimeFilterer, error) {
	contract, err := bindRollupTime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupTimeFilterer{contract: contract}, nil
}

// bindRollupTime binds a generic wrapper to an already deployed contract.
func bindRollupTime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupTimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTime *RollupTimeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTime.Contract.RollupTimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTime *RollupTimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTime.Contract.RollupTimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTime *RollupTimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTime.Contract.RollupTimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupTime *RollupTimeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupTime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupTime *RollupTimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupTime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupTime *RollupTimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupTime.Contract.contract.Transact(opts, method, params...)
}

// RollupUtilsABI is the input ABI used to generate the binding from.
const RollupUtilsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedValidAssertion\",\"type\":\"event\"}]"

// RollupUtilsBin is the compiled bytecode used for deploying new contracts.
var RollupUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158206422665059adef6af28e5fe5c7925f065ae5d0b5f3cd3407b4545c161ec467cf64736f6c634300050d0032"

// DeployRollupUtils deploys a new Ethereum contract, binding an instance of RollupUtils to it.
func DeployRollupUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupUtils{RollupUtilsCaller: RollupUtilsCaller{contract: contract}, RollupUtilsTransactor: RollupUtilsTransactor{contract: contract}, RollupUtilsFilterer: RollupUtilsFilterer{contract: contract}}, nil
}

// RollupUtils is an auto generated Go binding around an Ethereum contract.
type RollupUtils struct {
	RollupUtilsCaller     // Read-only binding to the contract
	RollupUtilsTransactor // Write-only binding to the contract
	RollupUtilsFilterer   // Log filterer for contract events
}

// RollupUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupUtilsSession struct {
	Contract     *RollupUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupUtilsCallerSession struct {
	Contract *RollupUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RollupUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupUtilsTransactorSession struct {
	Contract     *RollupUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RollupUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupUtilsRaw struct {
	Contract *RollupUtils // Generic contract binding to access the raw methods on
}

// RollupUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupUtilsCallerRaw struct {
	Contract *RollupUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// RollupUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupUtilsTransactorRaw struct {
	Contract *RollupUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupUtils creates a new instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtils(address common.Address, backend bind.ContractBackend) (*RollupUtils, error) {
	contract, err := bindRollupUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupUtils{RollupUtilsCaller: RollupUtilsCaller{contract: contract}, RollupUtilsTransactor: RollupUtilsTransactor{contract: contract}, RollupUtilsFilterer: RollupUtilsFilterer{contract: contract}}, nil
}

// NewRollupUtilsCaller creates a new read-only instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsCaller(address common.Address, caller bind.ContractCaller) (*RollupUtilsCaller, error) {
	contract, err := bindRollupUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsCaller{contract: contract}, nil
}

// NewRollupUtilsTransactor creates a new write-only instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupUtilsTransactor, error) {
	contract, err := bindRollupUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsTransactor{contract: contract}, nil
}

// NewRollupUtilsFilterer creates a new log filterer instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupUtilsFilterer, error) {
	contract, err := bindRollupUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsFilterer{contract: contract}, nil
}

// bindRollupUtils binds a generic wrapper to an already deployed contract.
func bindRollupUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupUtils *RollupUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupUtils.Contract.RollupUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupUtils *RollupUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupUtils.Contract.RollupUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupUtils *RollupUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupUtils.Contract.RollupUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupUtils *RollupUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RollupUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupUtils *RollupUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupUtils *RollupUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupUtils.Contract.contract.Transact(opts, method, params...)
}

// RollupUtilsConfirmedValidAssertionIterator is returned from FilterConfirmedValidAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedValidAssertion events raised by the RollupUtils contract.
type RollupUtilsConfirmedValidAssertionIterator struct {
	Event *RollupUtilsConfirmedValidAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUtilsConfirmedValidAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUtilsConfirmedValidAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUtilsConfirmedValidAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUtilsConfirmedValidAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUtilsConfirmedValidAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUtilsConfirmedValidAssertion represents a ConfirmedValidAssertion event raised by the RollupUtils contract.
type RollupUtilsConfirmedValidAssertion struct {
	NodeHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfirmedValidAssertion is a free log retrieval operation binding the contract event 0x89cc5e236414c34f1206c0c14d8ac5b0e5444b669b309aaca16fe3d27749f50e.
//
// Solidity: event ConfirmedValidAssertion(bytes32 indexed nodeHash)
func (_RollupUtils *RollupUtilsFilterer) FilterConfirmedValidAssertion(opts *bind.FilterOpts, nodeHash [][32]byte) (*RollupUtilsConfirmedValidAssertionIterator, error) {

	var nodeHashRule []interface{}
	for _, nodeHashItem := range nodeHash {
		nodeHashRule = append(nodeHashRule, nodeHashItem)
	}

	logs, sub, err := _RollupUtils.contract.FilterLogs(opts, "ConfirmedValidAssertion", nodeHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsConfirmedValidAssertionIterator{contract: _RollupUtils.contract, event: "ConfirmedValidAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedValidAssertion is a free log subscription operation binding the contract event 0x89cc5e236414c34f1206c0c14d8ac5b0e5444b669b309aaca16fe3d27749f50e.
//
// Solidity: event ConfirmedValidAssertion(bytes32 indexed nodeHash)
func (_RollupUtils *RollupUtilsFilterer) WatchConfirmedValidAssertion(opts *bind.WatchOpts, sink chan<- *RollupUtilsConfirmedValidAssertion, nodeHash [][32]byte) (event.Subscription, error) {

	var nodeHashRule []interface{}
	for _, nodeHashItem := range nodeHash {
		nodeHashRule = append(nodeHashRule, nodeHashItem)
	}

	logs, sub, err := _RollupUtils.contract.WatchLogs(opts, "ConfirmedValidAssertion", nodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUtilsConfirmedValidAssertion)
				if err := _RollupUtils.contract.UnpackLog(event, "ConfirmedValidAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedValidAssertion is a log parse operation binding the contract event 0x89cc5e236414c34f1206c0c14d8ac5b0e5444b669b309aaca16fe3d27749f50e.
//
// Solidity: event ConfirmedValidAssertion(bytes32 indexed nodeHash)
func (_RollupUtils *RollupUtilsFilterer) ParseConfirmedValidAssertion(log types.Log) (*RollupUtilsConfirmedValidAssertion, error) {
	event := new(RollupUtilsConfirmedValidAssertion)
	if err := _RollupUtils.contract.UnpackLog(event, "ConfirmedValidAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[]"

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d6218032e387455079786fec723dac80ece2fd829cb6194ef40576341993733164736f6c634300050d0032"

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
