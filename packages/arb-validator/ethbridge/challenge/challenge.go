// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challenge

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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209209a93bce785e91cc5cbff2f58b8586736753177e41f1b5a490a5098a55b3ae64736f6c634300050e0032"

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

// ChallengeABI is the input ABI used to generate the binding from.
const ChallengeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"machineHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"logAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"gases\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"challengerWrong\",\"type\":\"bool\"}],\"name\":\"TimedOutChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"asserterTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_preData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_machineHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_messageAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_logAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_gases\",\"type\":\"uint64[]\"},{\"internalType\":\"uint32\",\"name\":\"_totalSteps\",\"type\":\"uint32\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"challengerTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assertionToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"continueChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeFuncSigs = map[string]string{
	"d5345e07": "asserterTimedOut()",
	"8dc8f54a": "bisectAssertion(bytes32,bytes32[],bytes32[],bytes32[],uint64[],uint32)",
	"635e28a7": "challengerTimedOut()",
	"79d84776": "continueChallenge(uint256,bytes,bytes32,bytes32)",
	"2820245a": "init(address,address[2],uint128[2],uint32,bytes32,bytes32,uint64[2],bytes32)",
	"452e1137": "oneStepProof(bytes32,bytes32,uint64[2],bytes32,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
}

// ChallengeBin is the compiled bytecode used for deploying new contracts.
var ChallengeBin = "0x608060405234801561001057600080fd5b506120dd806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632820245a14610067578063452e1137146100bf578063635e28a7146101c857806379d84776146101d05780638dc8f54a14610280578063d5345e07146104b6575b600080fd5b6100bd600480360361016081101561007e57600080fd5b506001600160a01b03813516906020810190606081019063ffffffff60a0820135169060c08101359060e08101359061010081019061014001356104be565b005b6100bd60048036036101608110156100d657600080fd5b60408051808201825283359360208101359381019290916080830191808401906002908390839080828437600092019190915250919483359460208501359460408101359450606081013593506080810135926001600160401b0360a08301351692909160e081019060c00135600160201b81111561015457600080fd5b82018360208201111561016657600080fd5b803590602001918460018302840111600160201b8311171561018757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106ca945050505050565b6100bd610b88565b6100bd600480360360808110156101e657600080fd5b81359190810190604081016020820135600160201b81111561020757600080fd5b82018360208201111561021957600080fd5b803590602001918460018302840111600160201b8311171561023a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610c75565b6100bd600480360360c081101561029657600080fd5b81359190810190604081016020820135600160201b8111156102b757600080fd5b8201836020820111156102c957600080fd5b803590602001918460208302840111600160201b831117156102ea57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561033957600080fd5b82018360208201111561034b57600080fd5b803590602001918460208302840111600160201b8311171561036c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103bb57600080fd5b8201836020820111156103cd57600080fd5b803590602001918460208302840111600160201b831117156103ee57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561043d57600080fd5b82018360208201111561044f57600080fd5b803590602001918460208302840111600160201b8311171561047057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505050903563ffffffff16915061105d9050565b6100bd6110a0565b6000600554600160601b900460ff1660028111156104d857fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b8152509061058a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561054f578181015183820152602001610537565b50505050905090810190601f16801561057c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600080546001600160a01b0319166001600160a01b038a16179055604080516001600160c01b0319843560c090811b82166020808501919091528087013590911b909116602883015260308083018790528351808403909101815260508301845280519082012060708301526090820187905260b08083018590528351808403909101815260d0909201909252805191012060015563ffffffff8516430161063560028881611ef6565b506106436003896002611fa4565b506005805467ffffffffffffffff19166001600160401b0383169081176bffffffff00000000000000001916600160401b63ffffffff8a16021760ff60601b1916600160601b1790915560408051918252517f932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f5229181900360200190a1505050505050505050565b6001600554600160601b900460ff1660028111156106e457fe5b14604051806040016040528060098152602001684f53505f535441544560b81b815250906107535760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5060055460408051808201909152600c81526b4f53505f444541444c494e4560a01b6020820152906001600160401b03164311156107d25760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b506001805489516020808c0151604080516001600160c01b031960c095861b8116828601529290941b909116602884015260308084018f90528151808503909101815260508401808352815191840191909120630e6bb2fd60e11b909152605484018d905260748401959095526001600160401b038716609484015260b483018b905260d483018a905260f483018990526101148301889052519293928e9273__$2556963077056ca10a6804584182250fbf$__92631cd765fa9261013480840193829003018186803b1580156108a857600080fd5b505af41580156108bc573d6000803e3d6000fd5b505050506040513d60208110156108d257600080fd5b5051604080516020818101959095528082019390935260608084019290925280518084039092018252608083018082528251929094019190912060c08301909152600883526727a9a82fa82922ab60c11b60a09092019190915290911461097a5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b50600073__$efcad257db8183701794ea8506d55e247c$__6381557a418c8b8d8c8c8c8c8c8c8c6040518b63ffffffff1660e01b8152600401808b81526020018a600260200280838360005b838110156109de5781810151838201526020016109c6565b50505050905001898152602001888152602001878152602001868152602001858152602001848152602001836001600160401b03166001600160401b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a5a578181015183820152602001610a42565b50505050905090810190601f168015610a875780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060206040518083038186803b158015610aad57600080fd5b505af4158015610ac1573d6000803e3d6000fd5b505050506040513d6020811015610ad757600080fd5b505160408051808201909152600981526827a9a82fa82927a7a360b91b60208201529091508115610b495760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b506040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f590600090a1610b7b61118b565b5050505050505050505050565b6002600554600160601b900460ff166002811115610ba257fe5b14610bde5760405162461bcd60e51b81526004018080602001828103825260308152602001806120796030913960400191505060405180910390fd5b6005546001600160401b03164311610c37576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b604080516000815290517fd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c9181900360200190a1610c7361118b565b565b6002600554600160601b900460ff166002811115610c8f57fe5b1460405180604001604052806009815260200168434f4e5f535441544560b81b81525090610cfe5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5060015482146040518060400160405280600881526020016721a7a72fa82922ab60c11b81525090610d715760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5060055460408051808201909152600c81526b434f4e5f444541444c494e4560a01b6020820152906001600160401b0316431115610df05760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5060036001015460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b03163314610e6f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5073__$800fcb2f4a98daa165a5cdb21a355d7a15$__63b792d767848484886001016040518563ffffffff1660e01b81526004018080602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b83811015610eef578181015183820152602001610ed7565b50505050905090810190601f168015610f1c5780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b158015610f3c57600080fd5b505af4158015610f50573d6000803e3d6000fd5b505050506040513d6020811015610f6657600080fd5b505160408051808201909152600981526821a7a72fa82927a7a360b91b602082015290610fd45760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b50600580546001600160401b034363ffffffff600160401b60ff60601b19909416600160601b17938404160190811667ffffffffffffffff19929092168217909255600183905560408051878152602081019290925280517fb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a22269281900390910190a15050505050565b6110986040518060c001604052808881526020018781526020018681526020018581526020018481526020018363ffffffff1681525061126f565b505050505050565b6001600554600160601b900460ff1660028111156110ba57fe5b146110f65760405162461bcd60e51b815260040180806020018281038252602e81526020018061204b602e913960400191505060405180910390fd5b6005546001600160401b0316431161114f576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b604080516001815290517fd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c9181900360200190a1610c73611e6e565b60008054604080518082018252600280546001600160801b03808216600160801b909204811692909204011681526020810193909352516308b0246f60e21b81526001600160a01b03909116916322c091bc9160039190600481019060440183825b81546001600160a01b031681526001909101906020018083116111ed5750839050604080838360005b8381101561122e578181015183820152602001611216565b5050505090500192505050600060405180830381600087803b15801561125357600080fd5b505af1158015611267573d6000803e3d6000fd5b503392505050ff5b6000600182602001515103905081604001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906112f65760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5081606001515181600101146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906113715760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5081608001515181146040518060400160405280600a8152602001692124a9afa4a7282622a760b11b815250906113e95760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b50600554600160601b900460ff16600281111561140257fe5b600114604051806040016040528060098152602001684249535f535441544560b81b815250906114735760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5060055460408051808201909152600c81526b4249535f444541444c494e4560a01b6020820152906001600160401b03164311156114f25760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b5060036000015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b031633146115715760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b506000805b828110156115a7578360800151818151811061158e57fe5b6020026020010151820191508080600101915050611576565b506001548351602085015180516000906115bd57fe5b602002602001015173__$2556963077056ca10a6804584182250fbf$__631cd765fa876020015187815181106115ef57fe5b60200260200101518860a00151878a6040015160008151811061160e57fe5b60200260200101518b604001518b8151811061162657fe5b60200260200101518c6060015160008151811061163f57fe5b60200260200101518d606001518d8151811061165757fe5b60200260200101516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff168152602001866001600160401b03166001600160401b0316815260200185815260200184815260200183815260200182815260200197505050505050505060206040518083038186803b1580156116db57600080fd5b505af41580156116ef573d6000803e3d6000fd5b505050506040513d602081101561170557600080fd5b5051604080516020818101959095528082019390935260608084019290925280518084039092018252608083018082528251929094019190912060c0830190915260088352672124a9afa82922ab60c11b60a0909201919091529091146117ad5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054f578181015183820152602001610537565b506060826040519080825280602002602001820160405280156117da578160200160208202803883390190505b509050836000015184602001516000815181106117f357fe5b602002602001015173__$2556963077056ca10a6804584182250fbf$__631cd765fa876020015160018151811061182657fe5b60200260200101518763ffffffff168960a0015163ffffffff168161184757fe5b068863ffffffff168a60a0015163ffffffff168161186157fe5b0401896080015160008151811061187457fe5b60200260200101518a6040015160008151811061188d57fe5b60200260200101518b604001516001815181106118a657fe5b60200260200101518c606001516000815181106118bf57fe5b60200260200101518d606001516001815181106118d857fe5b60200260200101516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff168152602001866001600160401b03166001600160401b0316815260200185815260200184815260200183815260200182815260200197505050505050505060206040518083038186803b15801561195c57600080fd5b505af4158015611970573d6000803e3d6000fd5b505050506040513d602081101561198657600080fd5b505160408051602081810195909552808201939093526060808401929092528051808403909201825260809092019091528051910120815182906000906119c957fe5b602090810291909101015260015b83811015611bc457845160208601518051839081106119f257fe5b602002602001015173__$2556963077056ca10a6804584182250fbf$__631cd765fa88602001518560010181518110611a2757fe5b60200260200101518863ffffffff168a60a0015163ffffffff1681611a4857fe5b048a608001518781518110611a5957fe5b60200260200101518b604001518881518110611a7157fe5b60200260200101518c604001518960010181518110611a8c57fe5b60200260200101518d606001518a81518110611aa457fe5b60200260200101518e606001518b60010181518110611abf57fe5b60200260200101516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff168152602001866001600160401b03166001600160401b0316815260200185815260200184815260200183815260200182815260200197505050505050505060206040518083038186803b158015611b4357600080fd5b505af4158015611b57573d6000803e3d6000fd5b505050506040513d6020811015611b6d57600080fd5b5051604080516020818101959095528082019390935260608084019290925280518084039092018252608090920190915280519101208251839083908110611bb157fe5b60209081029190910101526001016119d7565b506040516309898dc160e41b815260206004820181815283516024840152835173__$800fcb2f4a98daa165a5cdb21a355d7a15$__93639898dc1093869392839260440191808601910280838360005b83811015611c2c578181015183820152602001611c14565b505050509050019250505060206040518083038186803b158015611c4f57600080fd5b505af4158015611c63573d6000803e3d6000fd5b505050506040513d6020811015611c7957600080fd5b5051600155600580546002919060ff60601b1916600160601b8302179055506000600560089054906101000a900463ffffffff1663ffffffff164301905080600560006101000a8154816001600160401b0302191690836001600160401b031602179055507f8d3e06942cd5f9e0032595e630e8361223df263995845c0364a53ba847a2d84a85602001518660400151876060015188608001518960a001518660405180806020018060200180602001806020018763ffffffff1663ffffffff168152602001866001600160401b03166001600160401b0316815260200185810385528b818151815260200191508051906020019060200280838360005b83811015611d8f578181015183820152602001611d77565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b83811015611dce578181015183820152602001611db6565b50505050905001858103835289818151815260200191508051906020019060200280838360005b83811015611e0d578181015183820152602001611df5565b50505050905001858103825288818151815260200191508051906020019060200280838360005b83811015611e4c578181015183820152602001611e34565b505050509050019a505050505050505050505060405180910390a15050505050565b6000805460408051808201825292835260028054600160801b81046001600160801b039081169181169290920401166020840152516308b0246f60e21b8152600380546001600160a01b03908116600480850191825291909416946322c091bc94929390929160448201916024018083116111ed575050825181528260408083836020611216565b600183019183908215611f945791602002820160005b83821115611f5f5783356001600160801b031683826101000a8154816001600160801b0302191690836001600160801b031602179055509260200192601001602081600f01049283019260010302611f0c565b8015611f925782816101000a8154906001600160801b030219169055601001602081600f01049283019260010302611f5f565b505b50611fa0929150611ff6565b5090565b8260028101928215611fea579160200282015b82811115611fea5781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190611fb7565b50611fa0929150612026565b61202391905b80821115611fa05780546fffffffffffffffffffffffffffffffff19168155600101611ffc565b90565b61202391905b80821115611fa05780546001600160a01b031916815560010161202c56fe43616e206f6e6c792074696d65206f7574206173736572746572206966206974206973207468656972207475726e43616e206f6e6c792074696d65206f7574206368616c6c656e676572206966206974206973207468656972207475726ea265627a7a72315820ac84e53b0a67fc635482c4a5720a9e5561f08418ca246b923d3f6144ee5a939f64736f6c634300050e0032"

// DeployChallenge deploys a new Ethereum contract, binding an instance of Challenge to it.
func DeployChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Challenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	protocolAddr, _, _, _ := DeployProtocol(auth, backend)
	ChallengeBin = strings.Replace(ChallengeBin, "__$2556963077056ca10a6804584182250fbf$__", protocolAddr.String()[2:], -1)

	merkleLibAddr, _, _, _ := DeployMerkleLib(auth, backend)
	ChallengeBin = strings.Replace(ChallengeBin, "__$800fcb2f4a98daa165a5cdb21a355d7a15$__", merkleLibAddr.String()[2:], -1)

	oneStepProofAddr, _, _, _ := DeployOneStepProof(auth, backend)
	ChallengeBin = strings.Replace(ChallengeBin, "__$efcad257db8183701794ea8506d55e247c$__", oneStepProofAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// Challenge is an auto generated Go binding around an Ethereum contract.
type Challenge struct {
	ChallengeCaller     // Read-only binding to the contract
	ChallengeTransactor // Write-only binding to the contract
	ChallengeFilterer   // Log filterer for contract events
}

// ChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeSession struct {
	Contract     *Challenge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeCallerSession struct {
	Contract *ChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTransactorSession struct {
	Contract     *ChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeRaw struct {
	Contract *Challenge // Generic contract binding to access the raw methods on
}

// ChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeCallerRaw struct {
	Contract *ChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTransactorRaw struct {
	Contract *ChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallenge creates a new instance of Challenge, bound to a specific deployed contract.
func NewChallenge(address common.Address, backend bind.ContractBackend) (*Challenge, error) {
	contract, err := bindChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// NewChallengeCaller creates a new read-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeCaller, error) {
	contract, err := bindChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeCaller{contract: contract}, nil
}

// NewChallengeTransactor creates a new write-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTransactor, error) {
	contract, err := bindChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTransactor{contract: contract}, nil
}

// NewChallengeFilterer creates a new log filterer instance of Challenge, bound to a specific deployed contract.
func NewChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFilterer, error) {
	contract, err := bindChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFilterer{contract: contract}, nil
}

// bindChallenge binds a generic wrapper to an already deployed contract.
func bindChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.ChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transact(opts, method, params...)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_Challenge *ChallengeTransactor) AsserterTimedOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "asserterTimedOut")
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_Challenge *ChallengeSession) AsserterTimedOut() (*types.Transaction, error) {
	return _Challenge.Contract.AsserterTimedOut(&_Challenge.TransactOpts)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_Challenge *ChallengeTransactorSession) AsserterTimedOut() (*types.Transaction, error) {
	return _Challenge.Contract.AsserterTimedOut(&_Challenge.TransactOpts)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x8dc8f54a.
//
// Solidity: function bisectAssertion(bytes32 _preData, bytes32[] _machineHashes, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint32 _totalSteps) returns()
func (_Challenge *ChallengeTransactor) BisectAssertion(opts *bind.TransactOpts, _preData [32]byte, _machineHashes [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint32) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "bisectAssertion", _preData, _machineHashes, _messageAccs, _logAccs, _gases, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x8dc8f54a.
//
// Solidity: function bisectAssertion(bytes32 _preData, bytes32[] _machineHashes, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint32 _totalSteps) returns()
func (_Challenge *ChallengeSession) BisectAssertion(_preData [32]byte, _machineHashes [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint32) (*types.Transaction, error) {
	return _Challenge.Contract.BisectAssertion(&_Challenge.TransactOpts, _preData, _machineHashes, _messageAccs, _logAccs, _gases, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x8dc8f54a.
//
// Solidity: function bisectAssertion(bytes32 _preData, bytes32[] _machineHashes, bytes32[] _messageAccs, bytes32[] _logAccs, uint64[] _gases, uint32 _totalSteps) returns()
func (_Challenge *ChallengeTransactorSession) BisectAssertion(_preData [32]byte, _machineHashes [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint32) (*types.Transaction, error) {
	return _Challenge.Contract.BisectAssertion(&_Challenge.TransactOpts, _preData, _machineHashes, _messageAccs, _logAccs, _gases, _totalSteps)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_Challenge *ChallengeTransactor) ChallengerTimedOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "challengerTimedOut")
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_Challenge *ChallengeSession) ChallengerTimedOut() (*types.Transaction, error) {
	return _Challenge.Contract.ChallengerTimedOut(&_Challenge.TransactOpts)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_Challenge *ChallengeTransactorSession) ChallengerTimedOut() (*types.Transaction, error) {
	return _Challenge.Contract.ChallengerTimedOut(&_Challenge.TransactOpts)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_Challenge *ChallengeTransactor) ContinueChallenge(opts *bind.TransactOpts, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "continueChallenge", _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_Challenge *ChallengeSession) ContinueChallenge(_assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.ContinueChallenge(&_Challenge.TransactOpts, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_Challenge *ChallengeTransactorSession) ContinueChallenge(_assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.ContinueChallenge(&_Challenge.TransactOpts, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address _vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_Challenge *ChallengeTransactor) Init(opts *bind.TransactOpts, _vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "init", _vmAddress, _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address _vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_Challenge *ChallengeSession) Init(_vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.Init(&_Challenge.TransactOpts, _vmAddress, _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address _vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_Challenge *ChallengeTransactorSession) Init(_vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _Challenge.Contract.Init(&_Challenge.TransactOpts, _vmAddress, _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x452e1137.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _afterHash, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_Challenge *ChallengeTransactor) OneStepProof(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _afterHash [32]byte, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "oneStepProof", _beforeHash, _beforeInbox, _timeBounds, _afterHash, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x452e1137.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _afterHash, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_Challenge *ChallengeSession) OneStepProof(_beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _afterHash [32]byte, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _Challenge.Contract.OneStepProof(&_Challenge.TransactOpts, _beforeHash, _beforeInbox, _timeBounds, _afterHash, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x452e1137.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _afterHash, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, uint64 _gas, bytes _proof) returns()
func (_Challenge *ChallengeTransactorSession) OneStepProof(_beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _afterHash [32]byte, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) (*types.Transaction, error) {
	return _Challenge.Contract.OneStepProof(&_Challenge.TransactOpts, _beforeHash, _beforeInbox, _timeBounds, _afterHash, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

// ChallengeBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the Challenge contract.
type ChallengeBisectedAssertionIterator struct {
	Event *ChallengeBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ChallengeBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeBisectedAssertion)
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
		it.Event = new(ChallengeBisectedAssertion)
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
func (it *ChallengeBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeBisectedAssertion represents a BisectedAssertion event raised by the Challenge contract.
type ChallengeBisectedAssertion struct {
	MachineHashes [][32]byte
	MessageAccs   [][32]byte
	LogAccs       [][32]byte
	Gases         []uint64
	TotalSteps    uint32
	Deadline      uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0x8d3e06942cd5f9e0032595e630e8361223df263995845c0364a53ba847a2d84a.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint32 totalSteps, uint64 deadline)
func (_Challenge *ChallengeFilterer) FilterBisectedAssertion(opts *bind.FilterOpts) (*ChallengeBisectedAssertionIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return &ChallengeBisectedAssertionIterator{contract: _Challenge.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0x8d3e06942cd5f9e0032595e630e8361223df263995845c0364a53ba847a2d84a.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint32 totalSteps, uint64 deadline)
func (_Challenge *ChallengeFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ChallengeBisectedAssertion) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeBisectedAssertion)
				if err := _Challenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0x8d3e06942cd5f9e0032595e630e8361223df263995845c0364a53ba847a2d84a.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint64[] gases, uint32 totalSteps, uint64 deadline)
func (_Challenge *ChallengeFilterer) ParseBisectedAssertion(log types.Log) (*ChallengeBisectedAssertion, error) {
	event := new(ChallengeBisectedAssertion)
	if err := _Challenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the Challenge contract.
type ChallengeContinuedChallengeIterator struct {
	Event *ChallengeContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeContinuedChallenge)
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
		it.Event = new(ChallengeContinuedChallenge)
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
func (it *ChallengeContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeContinuedChallenge represents a ContinuedChallenge event raised by the Challenge contract.
type ChallengeContinuedChallenge struct {
	AssertionIndex *big.Int
	Deadline       uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0xb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a2226.
//
// Solidity: event ContinuedChallenge(uint256 assertionIndex, uint64 deadline)
func (_Challenge *ChallengeFilterer) FilterContinuedChallenge(opts *bind.FilterOpts) (*ChallengeContinuedChallengeIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return &ChallengeContinuedChallengeIterator{contract: _Challenge.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0xb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a2226.
//
// Solidity: event ContinuedChallenge(uint256 assertionIndex, uint64 deadline)
func (_Challenge *ChallengeFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeContinuedChallenge) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeContinuedChallenge)
				if err := _Challenge.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0xb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a2226.
//
// Solidity: event ContinuedChallenge(uint256 assertionIndex, uint64 deadline)
func (_Challenge *ChallengeFilterer) ParseContinuedChallenge(log types.Log) (*ChallengeContinuedChallenge, error) {
	event := new(ChallengeContinuedChallenge)
	if err := _Challenge.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the Challenge contract.
type ChallengeInitiatedChallengeIterator struct {
	Event *ChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeInitiatedChallenge)
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
		it.Event = new(ChallengeInitiatedChallenge)
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
func (it *ChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the Challenge contract.
type ChallengeInitiatedChallenge struct {
	Deadline uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f522.
//
// Solidity: event InitiatedChallenge(uint64 deadline)
func (_Challenge *ChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ChallengeInitiatedChallengeIterator{contract: _Challenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f522.
//
// Solidity: event InitiatedChallenge(uint64 deadline)
func (_Challenge *ChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeInitiatedChallenge)
				if err := _Challenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f522.
//
// Solidity: event InitiatedChallenge(uint64 deadline)
func (_Challenge *ChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ChallengeInitiatedChallenge, error) {
	event := new(ChallengeInitiatedChallenge)
	if err := _Challenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the Challenge contract.
type ChallengeOneStepProofCompletedIterator struct {
	Event *ChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeOneStepProofCompleted)
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
		it.Event = new(ChallengeOneStepProofCompleted)
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
func (it *ChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the Challenge contract.
type ChallengeOneStepProofCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_Challenge *ChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*ChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &ChallengeOneStepProofCompletedIterator{contract: _Challenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_Challenge *ChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeOneStepProofCompleted)
				if err := _Challenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_Challenge *ChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*ChallengeOneStepProofCompleted, error) {
	event := new(ChallengeOneStepProofCompleted)
	if err := _Challenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeTimedOutChallengeIterator is returned from FilterTimedOutChallenge and is used to iterate over the raw logs and unpacked data for TimedOutChallenge events raised by the Challenge contract.
type ChallengeTimedOutChallengeIterator struct {
	Event *ChallengeTimedOutChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeTimedOutChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeTimedOutChallenge)
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
		it.Event = new(ChallengeTimedOutChallenge)
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
func (it *ChallengeTimedOutChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeTimedOutChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeTimedOutChallenge represents a TimedOutChallenge event raised by the Challenge contract.
type ChallengeTimedOutChallenge struct {
	ChallengerWrong bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedOutChallenge is a free log retrieval operation binding the contract event 0xd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c.
//
// Solidity: event TimedOutChallenge(bool challengerWrong)
func (_Challenge *ChallengeFilterer) FilterTimedOutChallenge(opts *bind.FilterOpts) (*ChallengeTimedOutChallengeIterator, error) {

	logs, sub, err := _Challenge.contract.FilterLogs(opts, "TimedOutChallenge")
	if err != nil {
		return nil, err
	}
	return &ChallengeTimedOutChallengeIterator{contract: _Challenge.contract, event: "TimedOutChallenge", logs: logs, sub: sub}, nil
}

// WatchTimedOutChallenge is a free log subscription operation binding the contract event 0xd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c.
//
// Solidity: event TimedOutChallenge(bool challengerWrong)
func (_Challenge *ChallengeFilterer) WatchTimedOutChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeTimedOutChallenge) (event.Subscription, error) {

	logs, sub, err := _Challenge.contract.WatchLogs(opts, "TimedOutChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeTimedOutChallenge)
				if err := _Challenge.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
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

// ParseTimedOutChallenge is a log parse operation binding the contract event 0xd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c.
//
// Solidity: event TimedOutChallenge(bool challengerWrong)
func (_Challenge *ChallengeFilterer) ParseTimedOutChallenge(log types.Log) (*ChallengeTimedOutChallenge, error) {
	event := new(ChallengeTimedOutChallenge)
	if err := _Challenge.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"b32\",\"type\":\"bytes32\"}],\"name\":\"bytes32string\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"out\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// DebugPrintFuncSigs maps the 4-byte function signature to its string representation.
var DebugPrintFuncSigs = map[string]string{
	"252fb38d": "bytes32string(bytes32)",
}

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x610202610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063252fb38d1461003a575b600080fd5b6100576004803603602081101561005057600080fd5b50356100cc565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610091578181015183820152602001610079565b50505050905090810190601f1680156100be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051818152606081810183529182919060208201818038833901905050905060005b602081101561019457600084826020811061010757fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b61012d8261019b565b85856002028151811061013c57fe5b60200101906001600160f81b031916908160001a90535061015c8161019b565b85856002026001018151811061016e57fe5b60200101906001600160f81b031916908160001a90535050600190920191506100f09050565b5092915050565b6000600a60f883901c10156101bb578160f81c60300160f81b90506101c8565b8160f81c60570160f81b90505b91905056fea265627a7a7231582020e95a9003c67208da26d16be5720ffffc6a3125ff2a26f4c08c700b8ab8defe64736f6c634300050e0032"

// DeployDebugPrint deploys a new Ethereum contract, binding an instance of DebugPrint to it.
func DeployDebugPrint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebugPrint, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DebugPrintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// DebugPrint is an auto generated Go binding around an Ethereum contract.
type DebugPrint struct {
	DebugPrintCaller     // Read-only binding to the contract
	DebugPrintTransactor // Write-only binding to the contract
	DebugPrintFilterer   // Log filterer for contract events
}

// DebugPrintCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebugPrintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebugPrintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebugPrintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebugPrintSession struct {
	Contract     *DebugPrint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebugPrintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebugPrintCallerSession struct {
	Contract *DebugPrintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DebugPrintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebugPrintTransactorSession struct {
	Contract     *DebugPrintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DebugPrintRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebugPrintRaw struct {
	Contract *DebugPrint // Generic contract binding to access the raw methods on
}

// DebugPrintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebugPrintCallerRaw struct {
	Contract *DebugPrintCaller // Generic read-only contract binding to access the raw methods on
}

// DebugPrintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebugPrintTransactorRaw struct {
	Contract *DebugPrintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebugPrint creates a new instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrint(address common.Address, backend bind.ContractBackend) (*DebugPrint, error) {
	contract, err := bindDebugPrint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// NewDebugPrintCaller creates a new read-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintCaller(address common.Address, caller bind.ContractCaller) (*DebugPrintCaller, error) {
	contract, err := bindDebugPrint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintCaller{contract: contract}, nil
}

// NewDebugPrintTransactor creates a new write-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintTransactor(address common.Address, transactor bind.ContractTransactor) (*DebugPrintTransactor, error) {
	contract, err := bindDebugPrint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintTransactor{contract: contract}, nil
}

// NewDebugPrintFilterer creates a new log filterer instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintFilterer(address common.Address, filterer bind.ContractFilterer) (*DebugPrintFilterer, error) {
	contract, err := bindDebugPrint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebugPrintFilterer{contract: contract}, nil
}

// bindDebugPrint binds a generic wrapper to an already deployed contract.
func bindDebugPrint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.DebugPrintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transact(opts, method, params...)
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintCaller) Bytes32string(opts *bind.CallOpts, b32 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DebugPrint.contract.Call(opts, out, "bytes32string", b32)
	return *ret0, err
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintSession) Bytes32string(b32 [32]byte) (string, error) {
	return _DebugPrint.Contract.Bytes32string(&_DebugPrint.CallOpts, b32)
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintCallerSession) Bytes32string(b32 [32]byte) (string, error) {
	return _DebugPrint.Contract.Bytes32string(&_DebugPrint.CallOpts, b32)
}

// IArbBaseABI is the input ABI used to generate the binding from.
const IArbBaseABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbBaseFuncSigs maps the 4-byte function signature to its string representation.
var IArbBaseFuncSigs = map[string]string{
	"22c091bc": "completeChallenge(address[2],uint128[2])",
}

// IArbBase is an auto generated Go binding around an Ethereum contract.
type IArbBase struct {
	IArbBaseCaller     // Read-only binding to the contract
	IArbBaseTransactor // Write-only binding to the contract
	IArbBaseFilterer   // Log filterer for contract events
}

// IArbBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbBaseSession struct {
	Contract     *IArbBase         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbBaseCallerSession struct {
	Contract *IArbBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IArbBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbBaseTransactorSession struct {
	Contract     *IArbBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IArbBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbBaseRaw struct {
	Contract *IArbBase // Generic contract binding to access the raw methods on
}

// IArbBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbBaseCallerRaw struct {
	Contract *IArbBaseCaller // Generic read-only contract binding to access the raw methods on
}

// IArbBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbBaseTransactorRaw struct {
	Contract *IArbBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbBase creates a new instance of IArbBase, bound to a specific deployed contract.
func NewIArbBase(address common.Address, backend bind.ContractBackend) (*IArbBase, error) {
	contract, err := bindIArbBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbBase{IArbBaseCaller: IArbBaseCaller{contract: contract}, IArbBaseTransactor: IArbBaseTransactor{contract: contract}, IArbBaseFilterer: IArbBaseFilterer{contract: contract}}, nil
}

// NewIArbBaseCaller creates a new read-only instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseCaller(address common.Address, caller bind.ContractCaller) (*IArbBaseCaller, error) {
	contract, err := bindIArbBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbBaseCaller{contract: contract}, nil
}

// NewIArbBaseTransactor creates a new write-only instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbBaseTransactor, error) {
	contract, err := bindIArbBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbBaseTransactor{contract: contract}, nil
}

// NewIArbBaseFilterer creates a new log filterer instance of IArbBase, bound to a specific deployed contract.
func NewIArbBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbBaseFilterer, error) {
	contract, err := bindIArbBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbBaseFilterer{contract: contract}, nil
}

// bindIArbBase binds a generic wrapper to an already deployed contract.
func bindIArbBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbBase *IArbBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbBase.Contract.IArbBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbBase *IArbBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbBase.Contract.IArbBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbBase *IArbBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbBase.Contract.IArbBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbBase *IArbBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbBase *IArbBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbBase *IArbBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbBase.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.Contract.CompleteChallenge(&_IArbBase.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbBase *IArbBaseTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbBase.Contract.CompleteChallenge(&_IArbBase.TransactOpts, _players, _rewards)
}

// IChallengeABI is the input ABI used to generate the binding from.
const IChallengeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeFuncSigs = map[string]string{
	"2820245a": "init(address,address[2],uint128[2],uint32,bytes32,bytes32,uint64[2],bytes32)",
}

// IChallenge is an auto generated Go binding around an Ethereum contract.
type IChallenge struct {
	IChallengeCaller     // Read-only binding to the contract
	IChallengeTransactor // Write-only binding to the contract
	IChallengeFilterer   // Log filterer for contract events
}

// IChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeSession struct {
	Contract     *IChallenge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeCallerSession struct {
	Contract *IChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeTransactorSession struct {
	Contract     *IChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeRaw struct {
	Contract *IChallenge // Generic contract binding to access the raw methods on
}

// IChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeCallerRaw struct {
	Contract *IChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeTransactorRaw struct {
	Contract *IChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallenge creates a new instance of IChallenge, bound to a specific deployed contract.
func NewIChallenge(address common.Address, backend bind.ContractBackend) (*IChallenge, error) {
	contract, err := bindIChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallenge{IChallengeCaller: IChallengeCaller{contract: contract}, IChallengeTransactor: IChallengeTransactor{contract: contract}, IChallengeFilterer: IChallengeFilterer{contract: contract}}, nil
}

// NewIChallengeCaller creates a new read-only instance of IChallenge, bound to a specific deployed contract.
func NewIChallengeCaller(address common.Address, caller bind.ContractCaller) (*IChallengeCaller, error) {
	contract, err := bindIChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeCaller{contract: contract}, nil
}

// NewIChallengeTransactor creates a new write-only instance of IChallenge, bound to a specific deployed contract.
func NewIChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeTransactor, error) {
	contract, err := bindIChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeTransactor{contract: contract}, nil
}

// NewIChallengeFilterer creates a new log filterer instance of IChallenge, bound to a specific deployed contract.
func NewIChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeFilterer, error) {
	contract, err := bindIChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeFilterer{contract: contract}, nil
}

// bindIChallenge binds a generic wrapper to an already deployed contract.
func bindIChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallenge *IChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallenge.Contract.IChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallenge *IChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallenge.Contract.IChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallenge *IChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallenge.Contract.IChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallenge *IChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallenge *IChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallenge *IChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallenge.Contract.contract.Transact(opts, method, params...)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 beforeHash, bytes32 beforeInbox, uint64[2] timeBounds, bytes32 _assertionHash) returns()
func (_IChallenge *IChallengeTransactor) Init(opts *bind.TransactOpts, vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, beforeHash [32]byte, beforeInbox [32]byte, timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IChallenge.contract.Transact(opts, "init", vmAddress, _players, _escrows, _challengePeriod, beforeHash, beforeInbox, timeBounds, _assertionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 beforeHash, bytes32 beforeInbox, uint64[2] timeBounds, bytes32 _assertionHash) returns()
func (_IChallenge *IChallengeSession) Init(vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, beforeHash [32]byte, beforeInbox [32]byte, timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IChallenge.Contract.Init(&_IChallenge.TransactOpts, vmAddress, _players, _escrows, _challengePeriod, beforeHash, beforeInbox, timeBounds, _assertionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 beforeHash, bytes32 beforeInbox, uint64[2] timeBounds, bytes32 _assertionHash) returns()
func (_IChallenge *IChallengeTransactorSession) Init(vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, beforeHash [32]byte, beforeInbox [32]byte, timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IChallenge.Contract.Init(&_IChallenge.TransactOpts, vmAddress, _players, _escrows, _challengePeriod, beforeHash, beforeInbox, timeBounds, _assertionHash)
}

// MachineABI is the input ABI used to generate the binding from.
const MachineABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"instructionStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"errHandlerHash\",\"type\":\"bytes32\"}],\"name\":\"machineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MachineFuncSigs maps the 4-byte function signature to its string representation.
var MachineFuncSigs = map[string]string{
	"c1355b59": "machineHash(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32)",
}

// MachineBin is the compiled bytecode used for deploying new contracts.
var MachineBin = "0x6101d6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063c1355b591461003a575b600080fd5b610075600480360360c081101561005057600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610087565b60408051918252519081900360200190f35b604080516101008101825260e081018881528152815160208181018452888252808301919091528251808201845287815282840152825180820184528681526060830152825180820184528581526080830152825190810190925282825260a0810191909152600060c08201819052906101009061010b565b979650505050505050565b600060028260c0015114156101225750600061019c565b60018260c0015114156101375750600161019c565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101205b91905056fea265627a7a72315820e4dfade81e2e9fb7a6677ab6e8c33f3013569e04dea722cc8209d7c1672ee22d64736f6c634300050e0032"

// DeployMachine deploys a new Ethereum contract, binding an instance of Machine to it.
func DeployMachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Machine, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Machine{MachineCaller: MachineCaller{contract: contract}, MachineTransactor: MachineTransactor{contract: contract}, MachineFilterer: MachineFilterer{contract: contract}}, nil
}

// Machine is an auto generated Go binding around an Ethereum contract.
type Machine struct {
	MachineCaller     // Read-only binding to the contract
	MachineTransactor // Write-only binding to the contract
	MachineFilterer   // Log filterer for contract events
}

// MachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineSession struct {
	Contract     *Machine          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineCallerSession struct {
	Contract *MachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineTransactorSession struct {
	Contract     *MachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineRaw struct {
	Contract *Machine // Generic contract binding to access the raw methods on
}

// MachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineCallerRaw struct {
	Contract *MachineCaller // Generic read-only contract binding to access the raw methods on
}

// MachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineTransactorRaw struct {
	Contract *MachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachine creates a new instance of Machine, bound to a specific deployed contract.
func NewMachine(address common.Address, backend bind.ContractBackend) (*Machine, error) {
	contract, err := bindMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Machine{MachineCaller: MachineCaller{contract: contract}, MachineTransactor: MachineTransactor{contract: contract}, MachineFilterer: MachineFilterer{contract: contract}}, nil
}

// NewMachineCaller creates a new read-only instance of Machine, bound to a specific deployed contract.
func NewMachineCaller(address common.Address, caller bind.ContractCaller) (*MachineCaller, error) {
	contract, err := bindMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineCaller{contract: contract}, nil
}

// NewMachineTransactor creates a new write-only instance of Machine, bound to a specific deployed contract.
func NewMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineTransactor, error) {
	contract, err := bindMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineTransactor{contract: contract}, nil
}

// NewMachineFilterer creates a new log filterer instance of Machine, bound to a specific deployed contract.
func NewMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineFilterer, error) {
	contract, err := bindMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineFilterer{contract: contract}, nil
}

// bindMachine binds a generic wrapper to an already deployed contract.
func bindMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machine *MachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Machine.Contract.MachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machine *MachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machine.Contract.MachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machine *MachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machine.Contract.MachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machine *MachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Machine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machine *MachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machine *MachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machine.Contract.contract.Transact(opts, method, params...)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_Machine *MachineCaller) MachineHash(opts *bind.CallOpts, instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Machine.contract.Call(opts, out, "machineHash", instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
	return *ret0, err
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_Machine *MachineSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _Machine.Contract.MachineHash(&_Machine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_Machine *MachineCallerSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _Machine.Contract.MachineHash(&_Machine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// MerkleLibABI is the input ABI used to generate the binding from.
const MerkleLibABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"generateAddressRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"generateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleLibFuncSigs maps the 4-byte function signature to its string representation.
var MerkleLibFuncSigs = map[string]string{
	"6a2dda67": "generateAddressRoot(address[])",
	"9898dc10": "generateRoot(bytes32[])",
	"b792d767": "verifyProof(bytes,bytes32,bytes32,uint256)",
}

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
var MerkleLibBin = "0x610575610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043f565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b6000815b600181511115610422576060600282516001018161031f57fe5b04604051908082528060200260200182016040528015610349578160200160208202803883390190505b50905060005b815181101561041a5782518160020260010110156103e25782816002028151811061037657fe5b602002602001015183826002026001018151811061039057fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d157fe5b602002602001018181525050610412565b8281600202815181106103f157fe5b602002602001015182828151811061040557fe5b6020026020010181815250505b60010161034f565b509050610305565b8060008151811061042f57fe5b6020026020010151915050919050565b600080838160205b88518111610532578089015193506020818a51036020018161046557fe5b0491505b60008211801561047c5750600286066001145b801561048a57508160020a86115b1561049d57600286046001019550610469565b600286066104e85783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104e057fe5b04955061052a565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052357fe5b0460010195505b602001610447565b50509094149594505050505056fea265627a7a7231582049c494227dbc2c0f1b182699d0c750b67090f60abf150ee56a1aec9c60573c1f64736f6c634300050e0032"

// DeployMerkleLib deploys a new Ethereum contract, binding an instance of MerkleLib to it.
func DeployMerkleLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleLib, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// MerkleLib is an auto generated Go binding around an Ethereum contract.
type MerkleLib struct {
	MerkleLibCaller     // Read-only binding to the contract
	MerkleLibTransactor // Write-only binding to the contract
	MerkleLibFilterer   // Log filterer for contract events
}

// MerkleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleLibSession struct {
	Contract     *MerkleLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleLibCallerSession struct {
	Contract *MerkleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MerkleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleLibTransactorSession struct {
	Contract     *MerkleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MerkleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleLibRaw struct {
	Contract *MerkleLib // Generic contract binding to access the raw methods on
}

// MerkleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleLibCallerRaw struct {
	Contract *MerkleLibCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleLibTransactorRaw struct {
	Contract *MerkleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleLib creates a new instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLib(address common.Address, backend bind.ContractBackend) (*MerkleLib, error) {
	contract, err := bindMerkleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// NewMerkleLibCaller creates a new read-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibCaller(address common.Address, caller bind.ContractCaller) (*MerkleLibCaller, error) {
	contract, err := bindMerkleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibCaller{contract: contract}, nil
}

// NewMerkleLibTransactor creates a new write-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleLibTransactor, error) {
	contract, err := bindMerkleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibTransactor{contract: contract}, nil
}

// NewMerkleLibFilterer creates a new log filterer instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleLibFilterer, error) {
	contract, err := bindMerkleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleLibFilterer{contract: contract}, nil
}

// bindMerkleLib binds a generic wrapper to an already deployed contract.
func bindMerkleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.MerkleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transact(opts, method, params...)
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibCaller) GenerateAddressRoot(opts *bind.CallOpts, _addresses []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "generateAddressRoot", _addresses)
	return *ret0, err
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibSession) GenerateAddressRoot(_addresses []common.Address) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateAddressRoot(&_MerkleLib.CallOpts, _addresses)
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibCallerSession) GenerateAddressRoot(_addresses []common.Address) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateAddressRoot(&_MerkleLib.CallOpts, _addresses)
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibCaller) GenerateRoot(opts *bind.CallOpts, _hashes [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "generateRoot", _hashes)
	return *ret0, err
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibSession) GenerateRoot(_hashes [][32]byte) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateRoot(&_MerkleLib.CallOpts, _hashes)
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibCallerSession) GenerateRoot(_hashes [][32]byte) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateRoot(&_MerkleLib.CallOpts, _hashes)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibCaller) VerifyProof(opts *bind.CallOpts, proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "verifyProof", proof, root, hash, index)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibSession) VerifyProof(proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	return _MerkleLib.Contract.VerifyProof(&_MerkleLib.CallOpts, proof, root, hash, index)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibCallerSession) VerifyProof(proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	return _MerkleLib.Contract.VerifyProof(&_MerkleLib.CallOpts, proof, root, hash, index)
}

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"81557a41": "validateProof(bytes32,uint64[2],bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint64,bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x61355b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806381557a411461003a575b600080fd5b610147600480360361016081101561005157600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594602085013594604081013594506060810135935060808101359260a08201359267ffffffffffffffff60c0840135169261010081019060e001356401000000008111156100d257600080fd5b8201836020820111156100e457600080fd5b8035906020019184600183028401116401000000008311171561010657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610159945050505050565b60408051918252519081900360200190f35b60006101b36040518061014001604052808d81526020018c81526020018b81526020018a81526020018981526020018881526020018781526020018681526020018567ffffffffffffffff168152602001848152506101c2565b9b9a5050505050505050505050565b600080808060606101d161335f565b6101d961335f565b6101e288611133565b93995092965090945092509050600160006101fc88611573565b67ffffffffffffffff168a610100015167ffffffffffffffff161461025f576040805162461bcd60e51b815260206004820152601460248201527324b73b30b634b21033b0b99034b710383937b7b360611b604482015290519081900360640190fd5b60ff8816600114156102a55761029e838660008151811061027c57fe5b60200260200101518760018151811061029157fe5b6020026020010151611579565b9150610f87565b60ff8816600214156102e45761029e83866000815181106102c257fe5b6020026020010151876001815181106102d757fe5b60200260200101516115c9565b60ff8816600314156103235761029e838660008151811061030157fe5b60200260200101518760018151811061031657fe5b602002602001015161160a565b60ff8816600414156103625761029e838660008151811061034057fe5b60200260200101518760018151811061035557fe5b602002602001015161164b565b60ff8816600514156103a15761029e838660008151811061037f57fe5b60200260200101518760018151811061039457fe5b602002602001015161169c565b60ff8816600614156103e05761029e83866000815181106103be57fe5b6020026020010151876001815181106103d357fe5b60200260200101516116ed565b60ff88166007141561041f5761029e83866000815181106103fd57fe5b60200260200101518760018151811061041257fe5b602002602001015161173e565b60ff8816600814156104735761029e838660008151811061043c57fe5b60200260200101518760018151811061045157fe5b60200260200101518860028151811061046657fe5b602002602001015161178f565b60ff8816600914156104c75761029e838660008151811061049057fe5b6020026020010151876001815181106104a557fe5b6020026020010151886002815181106104ba57fe5b60200260200101516117f9565b60ff8816600a14156105065761029e83866000815181106104e457fe5b6020026020010151876001815181106104f957fe5b6020026020010151611852565b60ff8816601014156105455761029e838660008151811061052357fe5b60200260200101518760018151811061053857fe5b6020026020010151611893565b60ff8816601114156105845761029e838660008151811061056257fe5b60200260200101518760018151811061057757fe5b60200260200101516118d4565b60ff8816601214156105c35761029e83866000815181106105a157fe5b6020026020010151876001815181106105b657fe5b6020026020010151611915565b60ff8816601314156106025761029e83866000815181106105e057fe5b6020026020010151876001815181106105f557fe5b6020026020010151611956565b60ff8816601414156106415761029e838660008151811061061f57fe5b60200260200101518760018151811061063457fe5b6020026020010151611997565b60ff88166015141561066b5761029e838660008151811061065e57fe5b60200260200101516119c3565b60ff8816601614156106aa5761029e838660008151811061068857fe5b60200260200101518760018151811061069d57fe5b6020026020010151611a09565b60ff8816601714156106e95761029e83866000815181106106c757fe5b6020026020010151876001815181106106dc57fe5b6020026020010151611a4a565b60ff8816601814156107285761029e838660008151811061070657fe5b60200260200101518760018151811061071b57fe5b6020026020010151611a8b565b60ff8816601914156107525761029e838660008151811061074557fe5b6020026020010151611acc565b60ff8816601a14156107915761029e838660008151811061076f57fe5b60200260200101518760018151811061078457fe5b6020026020010151611b02565b60ff8816601b14156107d05761029e83866000815181106107ae57fe5b6020026020010151876001815181106107c357fe5b6020026020010151611b43565b60ff8816602014156107fa5761029e83866000815181106107ed57fe5b6020026020010151611b84565b60ff8816602114156108245761029e838660008151811061081757fe5b6020026020010151611ba0565b60ff88166030141561084e5761029e838660008151811061084157fe5b6020026020010151611bbb565b60ff8816603114156108635761029e83611bc3565b60ff8816603214156108785761029e83611be4565b60ff8816603314156108a25761029e838660008151811061089557fe5b6020026020010151611bfd565b60ff8816603414156108cc5761029e83866000815181106108bf57fe5b6020026020010151611c16565b60ff88166035141561090b5761029e83866000815181106108e957fe5b6020026020010151876001815181106108fe57fe5b6020026020010151611c2c565b60ff8816603614156109205761029e83611c74565b60ff88166037141561093a5761029e838560000151611ca6565b60ff8816603814156109645761029e838660008151811061095757fe5b6020026020010151611cb8565b60ff8816603914156109f1576109786133c0565b6109878b610120015188611cca565b9199509750905087156109cb5760405162461bcd60e51b81526004018080602001828103825260218152602001806135066021913960400191505060405180910390fd5b6109db858263ffffffff611e5416565b6109eb848263ffffffff611e7616565b50610f87565b60ff8816603a1415610a065761029e83611e93565b60ff8816603b1415610a1757610f87565b60ff8816603c1415610a2c5761029e83611eb3565b60ff8816603d1415610a565761029e8386600081518110610a4957fe5b6020026020010151611ecc565b60ff881660401415610a805761029e8386600081518110610a7357fe5b6020026020010151611efa565b60ff881660411415610abf5761029e8386600081518110610a9d57fe5b602002602001015187600181518110610ab257fe5b6020026020010151611f1c565b60ff881660421415610b135761029e8386600081518110610adc57fe5b602002602001015187600181518110610af157fe5b602002602001015188600281518110610b0657fe5b6020026020010151611f4e565b60ff881660431415610b525761029e8386600081518110610b3057fe5b602002602001015187600181518110610b4557fe5b6020026020010151611f90565b60ff881660441415610ba65761029e8386600081518110610b6f57fe5b602002602001015187600181518110610b8457fe5b602002602001015188600281518110610b9957fe5b6020026020010151611fa2565b60ff881660501415610be55761029e8386600081518110610bc357fe5b602002602001015187600181518110610bd857fe5b6020026020010151611fc4565b60ff881660511415610c395761029e8386600081518110610c0257fe5b602002602001015187600181518110610c1757fe5b602002602001015188600281518110610c2c57fe5b602002602001015161203a565b60ff881660521415610c635761029e8386600081518110610c5657fe5b60200260200101516120b2565b60ff881660601415610c785761029e83611573565b60ff881660611415610d7557610ca28386600081518110610c9557fe5b60200260200101516120e5565b90925090508115610d6c578960e001518a60c001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610d215760405162461bcd60e51b81526004018080602001828103825260258152602001806134ba6025913960400191505060405180910390fd5b8960a001518a6080015114610d675760405162461bcd60e51b81526004018080602001828103825260278152602001806134df6027913960400191505060405180910390fd5b610d70565b5060005b610f87565b60ff881660701415610e6457610d9f8386600081518110610d9257fe5b6020026020010151612109565b90925090508115610d6c578960a001518a608001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610e1e5760405162461bcd60e51b815260040180806020018281038252602981526020018061344a6029913960400191505060405180910390fd5b8960e001518a60c0015114610d675760405162461bcd60e51b81526004018080602001828103825260268152602001806134736026913960400191505060405180910390fd5b60ff881660711415610f20576040805160028082526060828101909352816020015b610e8e6133c0565b815260200190600190039081610e8657505060208c0151909150610ec39060005b602002015167ffffffffffffffff16612123565b81600081518110610ed057fe5b6020026020010181905250610eef8b60200151600160028110610eaf57fe5b81600181518110610efc57fe5b60200260200101819052506109eb610f13826121a1565b859063ffffffff611e7616565b60ff881660721415610f5d5761029e8386600081518110610f3d57fe5b602002602001015160405180602001604052808e60400151815250612251565b60ff881660731415610f725760009150610f87565b60ff881660741415610f8757610f87836122c3565b80611018578960a001518a6080015114610fd25760405162461bcd60e51b81526004018080602001828103825260278152602001806134df6027913960400191505060405180910390fd5b8960e001518a60c00151146110185760405162461bcd60e51b81526004018080602001828103825260268152602001806134736026913960400191505060405180910390fd5b8161107a5760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a08401515114156110725761106d836122cd565b61107a565b60a083015183525b611083846122d7565b8a51146110c15760405162461bcd60e51b81526004018080602001828103825260228152602001806134286022913960400191505060405180910390fd5b6110ca836122d7565b8a6060015114611121576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b6000606061113f61335f565b61114761335f565b6000808061115361335f565b61115c8161236c565b61116b89610120015184612376565b909450909250905061117b61335f565b6111848261247b565b905060008a6101200151858151811061119957fe5b602001015160f81c60f81b60f81c905060008b610120015186600101815181106111bf57fe5b016020015160f81c905060006111d4826124d9565b905060608160405190808252806020026020018201604052801561121257816020015b6111ff6133c0565b8152602001906001900390816111f75790505b5090506002880197508360ff166000148061123057508360ff166001145b611281576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff84166113245760408051602080820180845289515163b697e08560e01b90915260ff87166024840152604483015291519091829173__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__9163b697e085916064808601929190818703018186803b1580156112ef57600080fd5b505af4158015611303573d6000803e3d6000fd5b505050506040513d602081101561131957600080fd5b50519052865261147b565b61132c6133c0565b61133b8f61012001518a611cca565b909a5090985090508715611396576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b82156113ba5780826000815181106113aa57fe5b60200260200101819052506113ca565b6113ca868263ffffffff611e7616565b604051806020016040528073__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__633c786053876113f9866124f3565b518c5151604080516001600160e01b031960e087901b16815260ff909416600485015260248401929092526044830152516064808301926020929190829003018186803b15801561144957600080fd5b505af415801561145d573d6000803e3d6000fd5b505050506040513d602081101561147357600080fd5b505190528752505b60ff84165b8281101561150f576114978f61012001518a611cca565b84518590859081106114a557fe5b6020908102919091010152995097508715611507576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b600101611480565b81511561155c575060005b8460ff1682510381101561155c5761155482826001855103038151811061153d57fe5b602002602001015188611e7690919063ffffffff16565b60010161151a565b50919d919c50939a50919850939650945050505050565b50600190565b600061158483612629565b1580611596575061159482612629565b155b156115a3575060006115c2565b825182518082016115ba878263ffffffff61263416565b600193505050505b9392505050565b60006115d483612629565b15806115e657506115e482612629565b155b156115f3575060006115c2565b825182518082026115ba878263ffffffff61263416565b600061161583612629565b1580611627575061162582612629565b155b15611634575060006115c2565b825182518082036115ba878263ffffffff61263416565b600061165683612629565b1580611668575061166682612629565b155b15611675575060006115c2565b8251825180611689576000925050506115c2565b8082046115ba878263ffffffff61263416565b60006116a783612629565b15806116b957506116b782612629565b155b156116c6575060006115c2565b82518251806116da576000925050506115c2565b8082056115ba878263ffffffff61263416565b60006116f883612629565b158061170a575061170882612629565b155b15611717575060006115c2565b825182518061172b576000925050506115c2565b8082066115ba878263ffffffff61263416565b600061174983612629565b158061175b575061175982612629565b155b15611768575060006115c2565b825182518061177c576000925050506115c2565b8082076115ba878263ffffffff61263416565b600061179a84612629565b15806117ac57506117aa83612629565b155b156117b9575060006117f1565b835183518351806117d057600093505050506117f1565b60008183850890506117e8898263ffffffff61263416565b60019450505050505b949350505050565b600061180484612629565b1580611816575061181483612629565b155b15611823575060006117f1565b8351835183518061183a57600093505050506117f1565b60008183850990506117e8898263ffffffff61263416565b600061185d83612629565b158061186f575061186d82612629565b155b1561187c575060006115c2565b8251825180820a6115ba878263ffffffff61263416565b600061189e83612629565b15806118b057506118ae82612629565b155b156118bd575060006115c2565b825182518082106115ba878263ffffffff61263416565b60006118df83612629565b15806118f157506118ef82612629565b155b156118fe575060006115c2565b825182518082116115ba878263ffffffff61263416565b600061192083612629565b1580611932575061193082612629565b155b1561193f575060006115c2565b825182518082126115ba878263ffffffff61263416565b600061196183612629565b1580611973575061197182612629565b155b15611980575060006115c2565b825182518082136115ba878263ffffffff61263416565b60006119b9610f136119a8846124f3565b516119b2866124f3565b5114612648565b5060019392505050565b60006119ce82612629565b6119e8576119e383600063ffffffff61263416565b6119ff565b815180156119fc858263ffffffff61263416565b50505b5060015b92915050565b6000611a1483612629565b1580611a265750611a2482612629565b155b15611a33575060006115c2565b825182518082166115ba878263ffffffff61263416565b6000611a5583612629565b1580611a675750611a6582612629565b155b15611a74575060006115c2565b825182518082176115ba878263ffffffff61263416565b6000611a9683612629565b1580611aa85750611aa682612629565b155b15611ab5575060006115c2565b825182518082186115ba878263ffffffff61263416565b6000611ad782612629565b611ae357506000611a03565b81518019611af7858263ffffffff61263416565b506001949350505050565b6000611b0d83612629565b1580611b1f5750611b1d82612629565b155b15611b2c575060006115c2565b8251825181811a6115ba878263ffffffff61263416565b6000611b4e83612629565b1580611b605750611b5e82612629565b155b15611b6d575060006115c2565b8251825181810b6115ba878263ffffffff61263416565b60006119ff611b92836124f3565b51849063ffffffff61263416565b60006119ff611bae83612671565b849063ffffffff611e7616565b600192915050565b6000611bdc8260800151836126fa90919063ffffffff16565b506001919050565b6000611bdc8260600151836126fa90919063ffffffff16565b6000611c08826124f3565b606084015250600192915050565b6000611c21826124f3565b835250600192915050565b6000611c3783612708565b611c43575060006115c2565b611c4c82612629565b611c58575060006115c2565b8151156119b957611c68836124f3565b84525060019392505050565b6000611bdc611c99611c8c611c87612715565b6124f3565b5160208501515114612648565b839063ffffffff611e7616565b60006119ff838363ffffffff6126fa16565b60006119ff838363ffffffff611e5416565b600080611cd56133c0565b84518410611d2a576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110611d3d57fe5b016020015160019092019160f81c90506000611d576133ee565b60ff8316611d8b57611d698985612792565b9094509150600084611d7a84612123565b91985096509450611e4d9350505050565b60ff831660011415611db257611da189856127b9565b9094509050600084611d7a83612926565b60ff831660021415611dd957611dc88985612792565b9094509150600084611d7a84612986565b600360ff841610801590611df05750600c60ff8416105b15611e2d57600219830160606000611e09838d89612a04565b909850925090508087611e1b846121a1565b99509950995050505050505050611e4d565b8260ff16612710016000611e416000612123565b91985096509450505050505b9250925092565b611e6a8260400151611e65836124f3565b612abf565b82604001819052505050565b611e878260200151611e65836124f3565b82602001819052505050565b6000611bdc611c99611ea6611c87612715565b5160408501515114612648565b6000611bdc8260a00151836126fa90919063ffffffff16565b6000611ed782612708565b611ee357506000611a03565b611eec826124f3565b60a084015250600192915050565b6000611f0c838363ffffffff611e7616565b6119ff838363ffffffff611e7616565b6000611f2e848363ffffffff611e7616565b611f3e848463ffffffff611e7616565b6119b9848363ffffffff611e7616565b6000611f60858363ffffffff611e7616565b611f70858463ffffffff611e7616565b611f80858563ffffffff611e7616565b611af7858363ffffffff611e7616565b6000611f3e848463ffffffff611e7616565b6000611fb4858563ffffffff611e7616565b611f80858463ffffffff611e7616565b6000611fcf83612629565b1580611fe15750611fdf82612b75565b155b15611fee575060006115c2565b611ff782612b84565b60ff1683600001511061200c575060006115c2565b6119b9826040015184600001518151811061202357fe5b602002602001015185611e7690919063ffffffff16565b600061204583612b75565b1580612057575061205584612629565b155b15612064575060006117f1565b61206d83612b84565b60ff16846000015110612082575060006117f1565b81836040015185600001518151811061209757fe5b6020908102919091010152611af7858463ffffffff611e7616565b60006120bd82612b75565b6120c957506000611a03565b6119ff6120d583612b84565b849060ff1663ffffffff61263416565b6000806120f0613415565b6120f9846124f3565b51600193509150505b9250929050565b6000806001612117846124f3565b51909590945092505050565b61212b6133c0565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612190565b61217d6133c0565b8152602001906001900390816121755790505b508152600060209091015292915050565b6121a96133c0565b6121b38251612b93565b612204576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b805160009061225f846124f3565b5114156122b3576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b6119b9848363ffffffff6126fa16565b600260c090910152565b600160c090910152565b600060028260c0015114156122ee5750600061112e565b60018260c0015114156123035750600161112e565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e0909301909252815191012061112e565b600060c090910152565b60008061238161335f565b61238961335f565b600060c0820181905261239c8787612b9a565b84529650905080156123b45793508492509050611e4d565b6123be8787612b9a565b60208501529650905080156123d95793508492509050611e4d565b6123e38787612b9a565b60408501529650905080156123fe5793508492509050611e4d565b6124088787612b9a565b60608501529650905080156124235793508492509050611e4d565b61242d8787612b9a565b60808501529650905080156124485793508492509050611e4d565b6124528787612b9a565b60a085015296509050801561246d5793508492509050611e4d565b506000969495509392505050565b61248361335f565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b60008060006124ea8460ff16612bd8565b50949350505050565b6124fb613415565b6060820151600c60ff9091161061254d576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661257a576040518060200160405280612571846000015161306b565b9052905061112e565b606082015160ff16600114156125c157604051806020016040528061257184602001516000015185602001516040015186602001516060015187602001516020015161308f565b606082015160ff16600214156125e6575060408051602081019091528151815261112e565b600360ff16826060015160ff161015801561260a57506060820151600c60ff909116105b156126275760405180602001604052806125718460400151613137565bfe5b6060015160ff161590565b611e878260200151611e65611c8784612123565b6126506133c0565b8115612667576126606001612123565b905061112e565b6126606000612123565b6126796133c0565b816060015160ff16600214156126c05760405162461bcd60e51b81526004018080602001828103825260218152602001806134996021913960400191505060405180910390fd5b606082015160ff166126d6576126606000612123565b816060015160ff16600114156126f0576126606001612123565b6126606003612123565b611e87826020015182612abf565b6060015160ff1660011490565b61271d6133c0565b604080516080808201835260008083528351918201845280825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612782565b61276f6133c0565b8152602001906001900390816127675790505b5081526003602090910152905090565b60008082816127a7868363ffffffff61328316565b60209290920196919550909350505050565b60006127c36133ee565b600083905060008582815181106127d657fe5b602001015160f81c60f81b60f81c9050818060010192505060008683815181106127fc57fe5b016020015160019384019360f89190911c915060009060ff8416141561289a5760006128266133c0565b6128308a87611cca565b9097509092509050811561288b576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b612894816124f3565b51925050505b60006128ac898663ffffffff61328316565b90506020850194508360ff16600114156128f1576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506121029050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b61292e6133c0565b604080516080810182526000808252602080830186905283518281529081018452919283019190612975565b6129626133c0565b81526020019060019003908161295a5790505b508152600160209091015292915050565b61298e6133c0565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916129f3565b6129e06133c0565b8152602001906001900390816129d85790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015612a4f57816020015b612a3c6133c0565b815260200190600190039081612a345790505b50905060005b8960ff168160ff161015612aa957612a6d8985611cca565b8451859060ff8616908110612a7e57fe5b6020908102919091010152945092508215612aa157509094509092509050612ab6565b600101612a55565b5060009550919350909150505b93509350939050565b612ac7613415565b6040805160028082526060828101909352816020015b612ae5613415565b815260200190600190039081612add5790505090508281600081518110612b0857fe5b60200260200101819052508381600181518110612b2157fe5b60200260200101819052506040518060200160405280612b6b6040518060400160405280612b528860000151612986565b8152602001612b648960000151612986565b905261329f565b9052949350505050565b6000611a03826060015161331e565b6000611a03826060015161333c565b6008101590565b600080612ba5613415565b836000612bb8878363ffffffff61328316565b604080516020808201909252918252600099930197509550909350505050565b6000806001831415612bf05750600290506001613066565b6002831415612c055750600290506001613066565b6003831415612c1a5750600290506001613066565b6004831415612c2f5750600290506001613066565b6005831415612c445750600290506001613066565b6006831415612c595750600290506001613066565b6007831415612c6e5750600290506001613066565b6008831415612c835750600390506001613066565b6009831415612c985750600390506001613066565b600a831415612cad5750600290506001613066565b6010831415612cc25750600290506001613066565b6011831415612cd75750600290506001613066565b6012831415612cec5750600290506001613066565b6013831415612d015750600290506001613066565b6014831415612d165750600290506001613066565b6015831415612d2a57506001905080613066565b6016831415612d3f5750600290506001613066565b6017831415612d545750600290506001613066565b6018831415612d695750600290506001613066565b6019831415612d7d57506001905080613066565b601a831415612d925750600290506001613066565b601b831415612da75750600290506001613066565b6020831415612dbb57506001905080613066565b6021831415612dcf57506001905080613066565b6030831415612de45750600190506000613066565b6031831415612df95750600090506001613066565b6032831415612e0e5750600090506001613066565b6033831415612e235750600190506000613066565b6034831415612e385750600190506000613066565b6035831415612e4d5750600290506000613066565b6036831415612e625750600090506001613066565b6037831415612e775750600090506001613066565b6038831415612e8c5750600190506000613066565b6039831415612ea15750600090506001613066565b603a831415612eb65750600090506001613066565b603b831415612eca57506000905080613066565b603c831415612edf5750600090506001613066565b603d831415612ef45750600190506000613066565b6040831415612f095750600190506002613066565b6041831415612f1e5750600290506003613066565b6042831415612f335750600390506004613066565b6043831415612f4757506002905080613066565b6044831415612f5b57506003905080613066565b6050831415612f705750600290506001613066565b6051831415612f855750600390506001613066565b6052831415612f9957506001905080613066565b6060831415612fad57506000905080613066565b6061831415612fc25750600190506000613066565b6070831415612fd75750600190506000613066565b6071831415612fec5750600090506001613066565b607283141561300057506001905080613066565b607383141561301457506000905080613066565b607483141561302857506000905080613066565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b600083156130e9575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206117f1565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6000600882511115613187576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156131b4578160200160208202803883390190505b50805190915060005b81811015613210576131cd613415565b6131e98683815181106131dc57fe5b60200260200101516124f3565b905080600001518483815181106131fc57fe5b6020908102919091010152506001016131bd565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015613259578181015183820152602001613241565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6000816020018351101561329657600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b6132c26133c0565b8152602001906001900390816132ba575050805190915060005b81811015613314578481600281106132f057fe5b602002015183828151811061330157fe5b60209081029190910101526001016132dc565b506117f182613137565b6000600c60ff8316108015611a03575050600360ff91909116101590565b60006133478261331e565b156133575750600219810161112e565b50600161112e565b6040518060e00160405280613372613415565b815260200161337f613415565b815260200161338c613415565b8152602001613399613415565b81526020016133a6613415565b81526020016133b3613415565b8152602001600081525090565b6040518060800160405280600081526020016133da6133ee565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f64654c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a723158208514ac6ab33b7175bf9d24e8ff479b220a46f53aedd8e186f4f3fa1f36eb1dde64736f6c634300050e0032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	valueAddr, _, _, _ := DeployValue(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__", valueAddr.String()[2:], -1)

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

// ValidateProof is a free data retrieval call binding the contract method 0x81557a41.
//
// Solidity: function validateProof(bytes32 beforeHash, uint64[2] timeBounds, bytes32 beforeInbox, bytes32 afterHash, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, beforeHash [32]byte, timeBounds [2]uint64, beforeInbox [32]byte, afterHash [32]byte, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", beforeHash, timeBounds, beforeInbox, afterHash, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0x81557a41.
//
// Solidity: function validateProof(bytes32 beforeHash, uint64[2] timeBounds, bytes32 beforeInbox, bytes32 afterHash, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(beforeHash [32]byte, timeBounds [2]uint64, beforeInbox [32]byte, afterHash [32]byte, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBounds, beforeInbox, afterHash, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0x81557a41.
//
// Solidity: function validateProof(bytes32 beforeHash, uint64[2] timeBounds, bytes32 beforeInbox, bytes32 afterHash, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, uint64 gas, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(beforeHash [32]byte, timeBounds [2]uint64, beforeInbox [32]byte, afterHash [32]byte, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, gas uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBounds, beforeInbox, afterHash, firstMessage, lastMessage, firstLog, lastLog, gas, proof)
}

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ProtocolFuncSigs = map[string]string{
	"1cd765fa": "generateAssertionHash(bytes32,uint32,uint64,bytes32,bytes32,bytes32,bytes32)",
	"e83f4bfe": "generateLastMessageHash(bytes)",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"85ecb92a": "generatePreconditionHash(bytes32,uint64[2],bytes32)",
}

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x610a93610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100555760003560e01c80624c28f61461005a5780631cd765fa146100b257806385ecb92a14610103578063e83f4bfe14610158575b600080fd5b6100a06004803603608081101561007057600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b03166101fe565b60408051918252519081900360200190f35b6100a0600480360360e08110156100c857600080fd5b5080359063ffffffff6020820135169067ffffffffffffffff6040820135169060608101359060808101359060a08101359060c001356102f0565b6100a06004803603608081101561011957600080fd5b604080518082018252833593928301929160608301919060208401906002908390839080828437600092019190915250919450509035915061035d9050565b6100a06004803603602081101561016e57600080fd5b81019060208101813564010000000081111561018957600080fd5b82018360208201111561019b57600080fd5b803590602001918460018302840111640100000000831117156101bd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103b1945050505050565b60408051600480825260a0820190925260009160609190816020015b6102226109f7565b81526020019060019003908161021a579050509050610240866104f6565b8160008151811061024d57fe5b602002602001018190525061026a836001600160a01b0316610574565b8160018151811061027757fe5b602002602001018190525061028b84610574565b8160028151811061029857fe5b60209081029190910101526102ba6affffffffffffffffffffff198616610574565b816003815181106102c757fe5b60200260200101819052506102e36102de826105f2565b6106a2565b519150505b949350505050565b6040805160208082019990995260e09790971b6001600160e01b0319168782015260c09590951b6001600160c01b0319166044870152604c860193909352606c850191909152608c84015260ac808401919091528151808403909101815260cc9092019052805191012090565b815160209283015160408051808601969096526001600160c01b031960c093841b8116878301529190921b166048850152605080850192909252805180850390920182526070909301909252815191012090565b8051600090819081908190815b818110156104e95773__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__63d36cfac288866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561043457818101518382015260200161041c565b50505050905090810190601f1680156104615780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561047e57600080fd5b505af4158015610492573d6000803e3d6000fd5b505050506040513d60408110156104a857600080fd5b50805160209182015160408051808501999099528881018290528051808a0382018152606090990190528751979092019690962095945092506001016103be565b509293505050505b919050565b6104fe6109f7565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610563565b6105506109f7565b8152602001906001900390816105485790505b508152600260209091015292915050565b61057c6109f7565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916105e1565b6105ce6109f7565b8152602001906001900390816105c65790505b508152600060209091015292915050565b6105fa6109f7565b61060482516107d8565b610655576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6106aa610a25565b6060820151600c60ff909116106106fc576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661072957604051806020016040528061072084600001516107df565b905290506104f1565b606082015160ff1660011415610770576040518060200160405280610720846020015160000151856020015160400151866020015160600151876020015160200151610803565b606082015160ff166002141561079557506040805160208101909152815181526104f1565b600360ff16826060015160ff16101580156107b957506060820151600c60ff909116105b156107d657604051806020016040528061072084604001516108ab565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561085d575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206102e8565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b60006008825111156108fb576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610928578160200160208202803883390190505b50805190915060005b8181101561098457610941610a25565b61095d86838151811061095057fe5b60200260200101516106a2565b9050806000015184838151811061097057fe5b602090810291909101015250600101610931565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156109cd5781810151838201526020016109b5565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b604051806080016040528060008152602001610a11610a37565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a723158205c2051c0600178207eaeab76b25b35eb48f24a857a49b8c9f7da0a7b2d0dae0a64736f6c634300050e0032"

// DeployProtocol deploys a new Ethereum contract, binding an instance of Protocol to it.
func DeployProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Protocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	valueAddr, _, _, _ := DeployValue(auth, backend)
	ProtocolBin = strings.Replace(ProtocolBin, "__$6f0fba43b1e3ecb6a82953dea0ca767fbf$__", valueAddr.String()[2:], -1)

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

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1cd765fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, uint64 _gas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _gas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _gas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1cd765fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, uint64 _gas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _gas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateAssertionHash(&_Protocol.CallOpts, _afterHash, _numSteps, _gas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x1cd765fa.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, uint64 _gas, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _gas uint64, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateAssertionHash(&_Protocol.CallOpts, _afterHash, _numSteps, _gas, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _messages []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateLastMessageHash", _messages)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateLastMessageHash(&_Protocol.CallOpts, _messages)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _Protocol.Contract.GenerateLastMessageHash(&_Protocol.CallOpts, _messages)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _Protocol.Contract.GenerateMessageStubHash(&_Protocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _Protocol.Contract.GenerateMessageStubHash(&_Protocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GeneratePreconditionHash(&_Protocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_Protocol *ProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _Protocol.Contract.GeneratePreconditionHash(&_Protocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeHashed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidHashed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"immediate\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasic\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashInt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueFuncSigs maps the 4-byte function signature to its string representation.
var ValueFuncSigs = map[string]string{
	"3d730ed2": "deserializeHashed(bytes)",
	"32e6cc21": "deserializeMessage(bytes,uint256)",
	"d36cfac2": "deserializeValidHashed(bytes,uint256)",
	"72403aa0": "getNextValid(bytes,uint256)",
	"826513e0": "hashCodePoint(uint8,bool,bytes32,bytes32)",
	"b697e085": "hashCodePointBasic(uint8,bytes32)",
	"3c786053": "hashCodePointImmediate(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"5043dff1": "hashInt(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x61152b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c806372403aa01161007057806372403aa014610300578063826513e014610425578063b2b9dc6214610459578063b697e0851461048a578063d36cfac2146104b0576100a8565b806332e6cc21146100ad578063364df277146101f95780633c786053146102135780633d730ed21461023f5780635043dff1146102e3575b600080fd5b610153600480360360408110156100c357600080fd5b810190602081018135600160201b8111156100dd57600080fd5b8201836020820111156100ef57600080fd5b803590602001918460018302840111600160201b8311171561011057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061056f915050565b604051808815151515815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101b85781810151838201526020016101a0565b50505050905090810190601f1680156101e55780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b61020161076d565b60408051918252519081900360200190f35b6102016004803603606081101561022957600080fd5b5060ff81351690602081013590604001356107e0565b6102016004803603602081101561025557600080fd5b810190602081018135600160201b81111561026f57600080fd5b82018360208201111561028157600080fd5b803590602001918460018302840111600160201b831117156102a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610832945050505050565b610201600480360360208110156102f957600080fd5b50356108a6565b6103a66004803603604081101561031657600080fd5b810190602081018135600160201b81111561033057600080fd5b82018360208201111561034257600080fd5b803590602001918460018302840111600160201b8311171561036357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506108ca915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103e95781810151838201526020016103d1565b50505050905090810190601f1680156104165780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6102016004803603608081101561043b57600080fd5b5060ff8135169060208101351515906040810135906060013561094e565b6104766004803603602081101561046f57600080fd5b50356109f7565b604080519115158252519081900360200190f35b610201600480360360408110156104a057600080fd5b5060ff81351690602001356109fe565b610556600480360360408110156104c657600080fd5b810190602081018135600160201b8111156104e057600080fd5b8201836020820111156104f257600080fd5b803590602001918460018302840111600160201b8311171561051357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610a45915050565b6040805192835260208301919091528051918290030190f35b6000806000806000806060600088965060008a888151811061058d57fe5b016020015160019098019760f81c9050600781146105bf576105b28b60018a03610a45565b9098509650610761915050565b6105c98b89610a45565b90985091506105e88b60018c016000198d8c030163ffffffff610abe16565b92508a88815181106105f657fe5b016020015160019098019760f81c90508015610619576105b28b60018a03610a45565b6106238b89610b3e565b80995081975050508a888151811061063757fe5b016020015160019098019760f81c9050801561065a576105b28b60018a03610a45565b6106648b89610b3e565b80995081965050508a888151811061067857fe5b016020015160019098019760f81c9050801561069b576105b28b60018a03610a45565b6106a58b89610b3e565b60408051600480825260a0820190925260019c50919a509195506060916020820160808038833901905050905082816000815181106106e057fe5b6020026020010181815250506106f5876108a6565b8160018151811061070257fe5b602002602001018181525050610717866108a6565b8160028151811061072457fe5b602002602001018181525050610739856108a6565b8160038151811061074657fe5b60200260200101818152505061075b81610b65565b97505050505b92959891949750929550565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156107b95781810151838201526020016107a1565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6000808061083e61146f565b610849856000610c25565b919450925090508215610891576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61089a81610daf565b5193505050505b919050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060606000806108d961146f565b6108e38787610c25565b91945092509050821561092b576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b8161093f888880840363ffffffff610abe16565b945094505050505b9250929050565b600083156109a8575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206109ef565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b6008101590565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b600080600080610a5361146f565b610a5d8787610c25565b919450925090508215610aa5576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b81610aaf82610daf565b51909890975095505050505050565b606081830184511015610ad057600080fd5b606082158015610aeb57604051915060208201604052610b35565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610b24578051835260209283019201610b0c565b5050858452601f01601f1916604052505b50949350505050565b6000808281610b53868363ffffffff610ee516565b60209290920196919550909350505050565b6000600882511115610bb5576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b8151600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610bfd578181015183820152602001610be5565b5050505090500192505050604051602081830303815290604052805190602001209050919050565b600080610c3061146f565b84518410610c85576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110610c9857fe5b016020015160019092019160f81c90506000610cb261149d565b60ff8316610ce657610cc48985610b3e565b9094509150600084610cd584610f01565b91985096509450610da89350505050565b60ff831660011415610d0d57610cfc8985610f7f565b9094509050600084610cd5836110da565b60ff831660021415610d3457610d238985610b3e565b9094509150600084610cd58461113a565b600360ff841610801590610d4b5750600c60ff8416105b15610d8857600219830160606000610d64838d896111b8565b909850925090508087610d7684611273565b99509950995050505050505050610da8565b8260ff16612710016000610d9c6000610f01565b91985096509450505050505b9250925092565b610db76114c4565b6060820151600c60ff90911610610e09576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610e36576040518060200160405280610e2d84600001516108a6565b905290506108a1565b606082015160ff1660011415610e7d576040518060200160405280610e2d84602001516000015185602001516040015186602001516060015187602001516020015161094e565b606082015160ff1660021415610ea257506040805160208101909152815181526108a1565b600360ff16826060015160ff1610158015610ec657506060820151600c60ff909116105b15610ee3576040518060200160405280610e2d8460400151611323565bfe5b60008160200183511015610ef857600080fd5b50016020015190565b610f0961146f565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610f6e565b610f5b61146f565b815260200190600190039081610f535790505b508152600060209091015292915050565b6000610f8961149d565b60008390506000858281518110610f9c57fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110610fc257fe5b016020015160019384019360f89190911c915060009060ff8416141561104e576000610fec61146f565b610ff68a87610c25565b9097509092509050811561103f576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61104881610daf565b51925050505b6000611060898663ffffffff610ee516565b90506020850194508360ff16600114156110a5576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506109479050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6110e261146f565b604080516080810182526000808252602080830186905283518281529081018452919283019190611129565b61111661146f565b81526020019060019003908161110e5790505b508152600160209091015292915050565b61114261146f565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916111a7565b61119461146f565b81526020019060019003908161118c5790505b508152600260209091015292915050565b60008060606000849050600060608860ff1660405190808252806020026020018201604052801561120357816020015b6111f061146f565b8152602001906001900390816111e85790505b50905060005b8960ff168160ff16101561125d576112218985610c25565b8451859060ff861690811061123257fe5b60209081029190910101529450925082156112555750909450909250905061126a565b600101611209565b5060009550919350909150505b93509350939050565b61127b61146f565b61128582516109f7565b6112d6576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6000600882511115611373576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156113a0578160200160208202803883390190505b50805190915060005b818110156113fc576113b96114c4565b6113d58683815181106113c857fe5b6020026020010151610daf565b905080600001518483815181106113e857fe5b6020908102919091010152506001016113a9565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561144557818101518382015260200161142d565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180608001604052806000815260200161148961149d565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4d61727368616c6c65642076616c7565206d7573742062652076616c69640000a265627a7a7231582002c6ecc6d323338594373e09189f176b7c79457073cee7bd7cd2ed946f380fa664736f6c634300050e0032"

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

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueCaller) DeserializeHashed(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "deserializeHashed", data)
	return *ret0, err
}

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueSession) DeserializeHashed(data []byte) ([32]byte, error) {
	return _Value.Contract.DeserializeHashed(&_Value.CallOpts, data)
}

// DeserializeHashed is a free data retrieval call binding the contract method 0x3d730ed2.
//
// Solidity: function deserializeHashed(bytes data) constant returns(bytes32)
func (_Value *ValueCallerSession) DeserializeHashed(data []byte) ([32]byte, error) {
	return _Value.Contract.DeserializeHashed(&_Value.CallOpts, data)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueCaller) DeserializeMessage(opts *bind.CallOpts, data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	ret := new(struct {
		Valid       bool
		Offset      *big.Int
		MessageHash [32]byte
		Destination *big.Int
		Value       *big.Int
		TokenType   *big.Int
		MessageData []byte
	})
	out := ret
	err := _Value.contract.Call(opts, out, "deserializeMessage", data, startOffset)
	return *ret, err
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _Value.Contract.DeserializeMessage(&_Value.CallOpts, data, startOffset)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_Value *ValueCallerSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _Value.Contract.DeserializeMessage(&_Value.CallOpts, data, startOffset)
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueCaller) DeserializeValidHashed(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Value.contract.Call(opts, out, "deserializeValidHashed", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueSession) DeserializeValidHashed(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _Value.Contract.DeserializeValidHashed(&_Value.CallOpts, data, offset)
}

// DeserializeValidHashed is a free data retrieval call binding the contract method 0xd36cfac2.
//
// Solidity: function deserializeValidHashed(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_Value *ValueCallerSession) DeserializeValidHashed(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _Value.Contract.DeserializeValidHashed(&_Value.CallOpts, data, offset)
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueCaller) GetNextValid(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Value.contract.Call(opts, out, "getNextValid", data, offset)
	return *ret0, *ret1, err
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueSession) GetNextValid(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _Value.Contract.GetNextValid(&_Value.CallOpts, data, offset)
}

// GetNextValid is a free data retrieval call binding the contract method 0x72403aa0.
//
// Solidity: function getNextValid(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_Value *ValueCallerSession) GetNextValid(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _Value.Contract.GetNextValid(&_Value.CallOpts, data, offset)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePoint(opts *bind.CallOpts, opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePoint", opcode, immediate, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePoint(&_Value.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePoint(&_Value.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePointBasic(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePointBasic", opcode, nextCodePoint)
	return *ret0, err
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePointBasic(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointBasic(&_Value.CallOpts, opcode, nextCodePoint)
}

// HashCodePointBasic is a free data retrieval call binding the contract method 0xb697e085.
//
// Solidity: function hashCodePointBasic(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePointBasic(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointBasic(&_Value.CallOpts, opcode, nextCodePoint)
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCaller) HashCodePointImmediate(opts *bind.CallOpts, opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashCodePointImmediate", opcode, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueSession) HashCodePointImmediate(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointImmediate(&_Value.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashCodePointImmediate is a free data retrieval call binding the contract method 0x3c786053.
//
// Solidity: function hashCodePointImmediate(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_Value *ValueCallerSession) HashCodePointImmediate(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _Value.Contract.HashCodePointImmediate(&_Value.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueSession) HashEmptyTuple() ([32]byte, error) {
	return _Value.Contract.HashEmptyTuple(&_Value.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_Value *ValueCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _Value.Contract.HashEmptyTuple(&_Value.CallOpts)
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueCaller) HashInt(opts *bind.CallOpts, val *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "hashInt", val)
	return *ret0, err
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueSession) HashInt(val *big.Int) ([32]byte, error) {
	return _Value.Contract.HashInt(&_Value.CallOpts, val)
}

// HashInt is a free data retrieval call binding the contract method 0x5043dff1.
//
// Solidity: function hashInt(uint256 val) constant returns(bytes32)
func (_Value *ValueCallerSession) HashInt(val *big.Int) ([32]byte, error) {
	return _Value.Contract.HashInt(&_Value.CallOpts, val)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueCaller) IsValidTupleSize(opts *bind.CallOpts, size *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "isValidTupleSize", size)
	return *ret0, err
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _Value.Contract.IsValidTupleSize(&_Value.CallOpts, size)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_Value *ValueCallerSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _Value.Contract.IsValidTupleSize(&_Value.CallOpts, size)
}
