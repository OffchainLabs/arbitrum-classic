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

// RollupCreatorNoProxyABI is the input ABI used to generate the binding from.
const RollupCreatorNoProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeCreator\",\"outputs\":[{\"internalType\":\"contractBridgeCreatorNoProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerDelaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"name\":\"createRollupNoProxy\",\"outputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTemplate\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBridgeCreatorNoProxy\",\"name\":\"_bridgeCreator\",\"type\":\"address\"},{\"internalType\":\"contractICloneable\",\"name\":\"_rollupTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"}],\"name\":\"setTemplates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupCreatorNoProxyBin is the compiled bytecode used for deploying new contracts.
var RollupCreatorNoProxyBin = "0x608060405234801561001057600080fd5b5060006100246001600160e01b0361007316565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350610077565b3390565b610c8d806100866000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c80635dbaf68b14610088578063715018a6146100ac5780638689d996146100b65780638babc4d9146100be5780638da5cb5b1461017f578063c8a7cb2114610187578063d93fe9c4146101c5578063f2fde38b146101cd578063f860cefa146101f3575b600080fd5b6100906101fb565b604080516001600160a01b039092168252519081900360200190f35b6100b461020a565b005b6100906102b6565b61009060048036036101608110156100d557600080fd5b8135916020810135916040820135916060810135916080820135916001600160a01b0360a082013581169260c083013582169260e081013590921691610100810135916101208201359190810190610160810161014082013564010000000081111561014057600080fd5b82018360208201111561015257600080fd5b8035906020019184600183028401116401000000008311171561017457600080fd5b5090925090506102c5565b61009061037b565b6100b46004803603608081101561019d57600080fd5b506001600160a01b03813581169160208101358216916040820135811691606001351661038a565b61009061043c565b6100b4600480360360208110156101e357600080fd5b50356001600160a01b031661044b565b61009061054d565b6003546001600160a01b031681565b61021261055c565b6001600160a01b031661022361037b565b6001600160a01b03161461026c576040805162461bcd60e51b81526020600482018190526024820152600080516020610c38833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6002546001600160a01b031681565b600061036a6040518061016001604052808f81526020018e81526020018d81526020018c81526020018b81526020018a6001600160a01b03168152602001896001600160a01b03168152602001886001600160a01b0316815260200187815260200186815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050915250610560565b9d9c50505050505050505050505050565b6000546001600160a01b031690565b61039261055c565b6001600160a01b03166103a361037b565b6001600160a01b0316146103ec576040805162461bcd60e51b81526020600482018190526024820152600080516020610c38833981519152604482015290519081900360640190fd5b600180546001600160a01b039586166001600160a01b0319918216179091556002805494861694821694909417909355600380549285169284169290921790915560048054919093169116179055565b6004546001600160a01b031681565b61045361055c565b6001600160a01b031661046461037b565b6001600160a01b0316146104ad576040805162461bcd60e51b81526020600482018190526024820152600080516020610c38833981519152604482015290519081900360640190fd5b6001600160a01b0381166104f25760405162461bcd60e51b8152600401808060200182810382526026815260200180610c126026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b3390565b600061056a610bd5565b60025461057f906001600160a01b0316610974565b6001600160a01b0390811660c0830181905260015460e08601516101008701516101208801516040805163058553e960e01b815260048101969096529286166024860152604485019190915260648401525192169163058553e99160848082019260a0929091908290030181600087803b1580156105fc57600080fd5b505af1158015610610573d6000803e3d6000fd5b505050506040513d60a081101561062657600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050856020018660400187606001886080018960a001856001600160a01b03166001600160a01b0316815250856001600160a01b03166001600160a01b0316815250856001600160a01b03166001600160a01b0316815250856001600160a01b03166001600160a01b0316815250856001600160a01b03166001600160a01b031681525050505050508060c001516001600160a01b031663fdaf5797846000015185602001518660400151876060015188608001518960a001518a60c001518b61014001516040518060c001604052808c602001516001600160a01b03166001600160a01b031681526020018c604001516001600160a01b03166001600160a01b031681526020018c60a001516001600160a01b03166001600160a01b031681526020018c608001516001600160a01b03166001600160a01b03168152602001600360009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600460009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152506040518a63ffffffff1660e01b8152600401808a8152602001898152602001888152602001878152602001868152602001856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b031681526020018060200183600660200280838360005b8381101561087e578181015183820152602001610866565b50505050905001828103825284818151815260200191508051906020019080838360005b838110156108ba5781810151838201526020016108a2565b50505050905090810190601f1680156108e75780820380516001836020036101000a031916815260200191505b509a5050505050505050505050600060405180830381600087803b15801561090e57600080fd5b505af1158015610922573d6000803e3d6000fd5b50505060c0820151604080516001600160a01b039092168252517f84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c92509081900360200190a160c0015190505b919050565b600081610989816001600160a01b0316610b32565b604051806040016040528060188152602001772727afa1a7a72a2920a1aa2fa1a627a722afa6a0a9aa22a960411b81525090610a435760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a085781810151838201526020016109f0565b50505050905090810190601f168015610a355780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b158015610a7d57600080fd5b505afa158015610a91573d6000803e3d6000fd5b505050506040513d6020811015610aa757600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b602082015290610b185760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610a085781810151838201526020016109f0565b50610b2b836001600160a01b0316610b38565b9392505050565b3b151590565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b03811661096f576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091529056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a26469706673582212205c882277c8ea0bc11a0eb0f6311fe58275747c29808a4230f4acd7aeabf75eee64736f6c634300060b0033"

// DeployRollupCreatorNoProxy deploys a new Ethereum contract, binding an instance of RollupCreatorNoProxy to it.
func DeployRollupCreatorNoProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupCreatorNoProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorNoProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCreatorNoProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCreatorNoProxy{RollupCreatorNoProxyCaller: RollupCreatorNoProxyCaller{contract: contract}, RollupCreatorNoProxyTransactor: RollupCreatorNoProxyTransactor{contract: contract}, RollupCreatorNoProxyFilterer: RollupCreatorNoProxyFilterer{contract: contract}}, nil
}

// RollupCreatorNoProxy is an auto generated Go binding around an Ethereum contract.
type RollupCreatorNoProxy struct {
	RollupCreatorNoProxyCaller     // Read-only binding to the contract
	RollupCreatorNoProxyTransactor // Write-only binding to the contract
	RollupCreatorNoProxyFilterer   // Log filterer for contract events
}

// RollupCreatorNoProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCreatorNoProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorNoProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCreatorNoProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorNoProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCreatorNoProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorNoProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCreatorNoProxySession struct {
	Contract     *RollupCreatorNoProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupCreatorNoProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCreatorNoProxyCallerSession struct {
	Contract *RollupCreatorNoProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RollupCreatorNoProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCreatorNoProxyTransactorSession struct {
	Contract     *RollupCreatorNoProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RollupCreatorNoProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCreatorNoProxyRaw struct {
	Contract *RollupCreatorNoProxy // Generic contract binding to access the raw methods on
}

// RollupCreatorNoProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCreatorNoProxyCallerRaw struct {
	Contract *RollupCreatorNoProxyCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCreatorNoProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCreatorNoProxyTransactorRaw struct {
	Contract *RollupCreatorNoProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCreatorNoProxy creates a new instance of RollupCreatorNoProxy, bound to a specific deployed contract.
func NewRollupCreatorNoProxy(address common.Address, backend bind.ContractBackend) (*RollupCreatorNoProxy, error) {
	contract, err := bindRollupCreatorNoProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxy{RollupCreatorNoProxyCaller: RollupCreatorNoProxyCaller{contract: contract}, RollupCreatorNoProxyTransactor: RollupCreatorNoProxyTransactor{contract: contract}, RollupCreatorNoProxyFilterer: RollupCreatorNoProxyFilterer{contract: contract}}, nil
}

// NewRollupCreatorNoProxyCaller creates a new read-only instance of RollupCreatorNoProxy, bound to a specific deployed contract.
func NewRollupCreatorNoProxyCaller(address common.Address, caller bind.ContractCaller) (*RollupCreatorNoProxyCaller, error) {
	contract, err := bindRollupCreatorNoProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyCaller{contract: contract}, nil
}

// NewRollupCreatorNoProxyTransactor creates a new write-only instance of RollupCreatorNoProxy, bound to a specific deployed contract.
func NewRollupCreatorNoProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCreatorNoProxyTransactor, error) {
	contract, err := bindRollupCreatorNoProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyTransactor{contract: contract}, nil
}

// NewRollupCreatorNoProxyFilterer creates a new log filterer instance of RollupCreatorNoProxy, bound to a specific deployed contract.
func NewRollupCreatorNoProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCreatorNoProxyFilterer, error) {
	contract, err := bindRollupCreatorNoProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyFilterer{contract: contract}, nil
}

// bindRollupCreatorNoProxy binds a generic wrapper to an already deployed contract.
func bindRollupCreatorNoProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorNoProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreatorNoProxy.Contract.RollupCreatorNoProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.RollupCreatorNoProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.RollupCreatorNoProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreatorNoProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.contract.Transact(opts, method, params...)
}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) BridgeCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "bridgeCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) BridgeCreator() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.BridgeCreator(&_RollupCreatorNoProxy.CallOpts)
}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) BridgeCreator() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.BridgeCreator(&_RollupCreatorNoProxy.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) ChallengeFactory() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.ChallengeFactory(&_RollupCreatorNoProxy.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) ChallengeFactory() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.ChallengeFactory(&_RollupCreatorNoProxy.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) NodeFactory() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.NodeFactory(&_RollupCreatorNoProxy.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) NodeFactory() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.NodeFactory(&_RollupCreatorNoProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) Owner() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.Owner(&_RollupCreatorNoProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) Owner() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.Owner(&_RollupCreatorNoProxy.CallOpts)
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCaller) RollupTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreatorNoProxy.contract.Call(opts, &out, "rollupTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) RollupTemplate() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.RollupTemplate(&_RollupCreatorNoProxy.CallOpts)
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyCallerSession) RollupTemplate() (common.Address, error) {
	return _RollupCreatorNoProxy.Contract.RollupTemplate(&_RollupCreatorNoProxy.CallOpts)
}

// CreateRollupNoProxy is a paid mutator transaction binding the contract method 0x8babc4d9.
//
// Solidity: function createRollupNoProxy(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactor) CreateRollupNoProxy(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.contract.Transact(opts, "createRollupNoProxy", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// CreateRollupNoProxy is a paid mutator transaction binding the contract method 0x8babc4d9.
//
// Solidity: function createRollupNoProxy(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) CreateRollupNoProxy(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.CreateRollupNoProxy(&_RollupCreatorNoProxy.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// CreateRollupNoProxy is a paid mutator transaction binding the contract method 0x8babc4d9.
//
// Solidity: function createRollupNoProxy(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorSession) CreateRollupNoProxy(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.CreateRollupNoProxy(&_RollupCreatorNoProxy.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.RenounceOwnership(&_RollupCreatorNoProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.RenounceOwnership(&_RollupCreatorNoProxy.TransactOpts)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactor) SetTemplates(opts *bind.TransactOpts, _bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.contract.Transact(opts, "setTemplates", _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) SetTemplates(_bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.SetTemplates(&_RollupCreatorNoProxy.TransactOpts, _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorSession) SetTemplates(_bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.SetTemplates(&_RollupCreatorNoProxy.TransactOpts, _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.TransferOwnership(&_RollupCreatorNoProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreatorNoProxy *RollupCreatorNoProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreatorNoProxy.Contract.TransferOwnership(&_RollupCreatorNoProxy.TransactOpts, newOwner)
}

// RollupCreatorNoProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RollupCreatorNoProxy contract.
type RollupCreatorNoProxyOwnershipTransferredIterator struct {
	Event *RollupCreatorNoProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RollupCreatorNoProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorNoProxyOwnershipTransferred)
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
		it.Event = new(RollupCreatorNoProxyOwnershipTransferred)
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
func (it *RollupCreatorNoProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorNoProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorNoProxyOwnershipTransferred represents a OwnershipTransferred event raised by the RollupCreatorNoProxy contract.
type RollupCreatorNoProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupCreatorNoProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreatorNoProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyOwnershipTransferredIterator{contract: _RollupCreatorNoProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupCreatorNoProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreatorNoProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorNoProxyOwnershipTransferred)
				if err := _RollupCreatorNoProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) ParseOwnershipTransferred(log types.Log) (*RollupCreatorNoProxyOwnershipTransferred, error) {
	event := new(RollupCreatorNoProxyOwnershipTransferred)
	if err := _RollupCreatorNoProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorNoProxyRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the RollupCreatorNoProxy contract.
type RollupCreatorNoProxyRollupCreatedIterator struct {
	Event *RollupCreatorNoProxyRollupCreated // Event containing the contract specifics and raw log

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
func (it *RollupCreatorNoProxyRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorNoProxyRollupCreated)
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
		it.Event = new(RollupCreatorNoProxyRollupCreated)
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
func (it *RollupCreatorNoProxyRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorNoProxyRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorNoProxyRollupCreated represents a RollupCreated event raised by the RollupCreatorNoProxy contract.
type RollupCreatorNoProxyRollupCreated struct {
	RollupAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*RollupCreatorNoProxyRollupCreatedIterator, error) {

	logs, sub, err := _RollupCreatorNoProxy.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &RollupCreatorNoProxyRollupCreatedIterator{contract: _RollupCreatorNoProxy.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupCreatorNoProxyRollupCreated) (event.Subscription, error) {

	logs, sub, err := _RollupCreatorNoProxy.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorNoProxyRollupCreated)
				if err := _RollupCreatorNoProxy.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreatorNoProxy *RollupCreatorNoProxyFilterer) ParseRollupCreated(log types.Log) (*RollupCreatorNoProxyRollupCreated, error) {
	event := new(RollupCreatorNoProxyRollupCreated)
	if err := _RollupCreatorNoProxy.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
