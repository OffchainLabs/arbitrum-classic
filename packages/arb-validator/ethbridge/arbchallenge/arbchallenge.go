// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbchallenge

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

// ArbChallengeABI is the input ABI used to generate the binding from.
const ArbChallengeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"machineHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"logAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"challengerWrong\",\"type\":\"bool\"}],\"name\":\"TimedOutChallenge\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"asserterTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_preData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_machineHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_messageAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_logAccs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint32\",\"name\":\"_totalSteps\",\"type\":\"uint32\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"challengerTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assertionToChallenge\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"continueChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vmAddress\",\"type\":\"address\"},{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbChallengeFuncSigs maps the 4-byte function signature to its string representation.
var ArbChallengeFuncSigs = map[string]string{
	"d5345e07": "asserterTimedOut()",
	"a0294172": "bisectAssertion(bytes32,bytes32[],bytes32[],bytes32[],uint32)",
	"635e28a7": "challengerTimedOut()",
	"79d84776": "continueChallenge(uint256,bytes,bytes32,bytes32)",
	"2820245a": "init(address,address[2],uint128[2],uint32,bytes32,bytes32,uint64[2],bytes32)",
	"5552a0ff": "oneStepProof(bytes32,bytes32,uint64[2],bytes32,bytes32,bytes32,bytes32,bytes32,bytes)",
}

// ArbChallengeBin is the compiled bytecode used for deploying new contracts.
var ArbChallengeBin = "0x608060405234801561001057600080fd5b50610ccd806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632820245a146100675780635552a0ff146100bf578063635e28a7146101b957806379d84776146101c1578063a029417214610271578063d5345e0714610425575b600080fd5b6100bd600480360361016081101561007e57600080fd5b506001600160a01b03813516906020810190606081019063ffffffff60a0820135169060c08101359060e081013590610100810190610140013561042d565b005b6100bd60048036036101408110156100d657600080fd5b6040805180820182528335936020810135938101929091608083019180840190600290839083908082843760009201919091525091948335946020850135946040810135945060608101359350608081013592919060c081019060a00135600160201b81111561014557600080fd5b82018360208201111561015757600080fd5b803590602001918460018302840111600160201b8311171561017857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061053f945050505050565b6100bd6106b2565b6100bd600480360360808110156101d757600080fd5b81359190810190604081016020820135600160201b8111156101f857600080fd5b82018360208201111561020a57600080fd5b803590602001918460018302840111600160201b8311171561022b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356107a0565b6100bd600480360360a081101561028757600080fd5b81359190810190604081016020820135600160201b8111156102a857600080fd5b8201836020820111156102ba57600080fd5b803590602001918460208302840111600160201b831117156102db57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561032a57600080fd5b82018360208201111561033c57600080fd5b803590602001918460208302840111600160201b8311171561035d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103ac57600080fd5b8201836020820111156103be57600080fd5b803590602001918460208302840111600160201b831117156103df57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505050903563ffffffff1691506108909050565b6100bd6109e2565b73__$0b6e9c2b3f0d3b9a73d9361721242908cd$__6352426fc360008a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a8152602001896001600160a01b03166001600160a01b0316815260200188600260200280828437600083820152601f01601f191690910190508760408082843760008382015263ffffffff8916601f909101601f191690920191825250602081018690526040808201869052606090910190849080828437600081840152601f19601f820116905080830192505050828152602001995050505050505050505060006040518083038186803b15801561051d57600080fd5b505af4158015610531573d6000803e3d6000fd5b505050505050505050505050565b73__$0b6e9c2b3f0d3b9a73d9361721242908cd$__636e4a55a960008b8b8b8b8b8b8b8b8b6040518b63ffffffff1660e01b8152600401808b81526020018a815260200189815260200188600260200280838360005b838110156105ad578181015183820152602001610595565b5050505090500187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561060b5781810151838201526020016105f3565b50505050905090810190601f1680156106385780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060006040518083038186803b15801561065e57600080fd5b505af4158015610672573d6000803e3d6000fd5b50506040517f117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5925060009150a16106a7610ace565b505050505050505050565b6002600554600160601b900460ff1660028111156106cc57fe5b146107085760405162461bcd60e51b8152600401808060200182810382526030815260200180610c696030913960400191505060405180910390fd5b60055467ffffffffffffffff164311610762576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b604080516000815290517fd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c9181900360200190a161079e610ace565b565b73__$0b6e9c2b3f0d3b9a73d9361721242908cd$__63110112ae6000868686866040518663ffffffff1660e01b81526004018086815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561082457818101518382015260200161080c565b50505050905090810190601f1680156108515780820380516001836020036101000a031916815260200191505b50965050505050505060006040518083038186803b15801561087257600080fd5b505af4158015610886573d6000803e3d6000fd5b5050505050505050565b73__$0b6e9c2b3f0d3b9a73d9361721242908cd$__632716a27b600087878787876040518763ffffffff1660e01b8152600401808781526020018681526020018060200180602001806020018563ffffffff1663ffffffff168152602001848103845288818151815260200191508051906020019060200280838360005b8381101561092657818101518382015260200161090e565b50505050905001848103835287818151815260200191508051906020019060200280838360005b8381101561096557818101518382015260200161094d565b50505050905001848103825286818151815260200191508051906020019060200280838360005b838110156109a457818101518382015260200161098c565b50505050905001995050505050505050505060006040518083038186803b1580156109ce57600080fd5b505af41580156106a7573d6000803e3d6000fd5b6001600554600160601b900460ff1660028111156109fc57fe5b14610a385760405162461bcd60e51b815260040180806020018281038252602e815260200180610c3b602e913960400191505060405180910390fd5b60055467ffffffffffffffff164311610a92576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b604080516001815290517fd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c9181900360200190a161079e610bb2565b60008054604080518082018252600280546001600160801b03808216600160801b909204811692909204011681526020810193909352516308b0246f60e21b81526001600160a01b03909116916322c091bc9160039190600481019060440183825b81546001600160a01b03168152600190910190602001808311610b305750839050604080838360005b83811015610b71578181015183820152602001610b59565b5050505090500192505050600060405180830381600087803b158015610b9657600080fd5b505af1158015610baa573d6000803e3d6000fd5b503392505050ff5b6000805460408051808201825292835260028054600160801b81046001600160801b039081169181169290920401166020840152516308b0246f60e21b8152600380546001600160a01b03908116600480850191825291909416946322c091bc9492939092916044820191602401808311610b30575050825181528260408083836020610b5956fe43616e206f6e6c792074696d65206f7574206173736572746572206966206974206973207468656972207475726e43616e206f6e6c792074696d65206f7574206368616c6c656e676572206966206974206973207468656972207475726ea265627a7a723158201779be083d2658c43020fde72670a8894b10bc43516b7ab94efbc01be3805e0d64736f6c634300050d0032"

// DeployArbChallenge deploys a new Ethereum contract, binding an instance of ArbChallenge to it.
func DeployArbChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbChallenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	challengeImplAddr, _, _, _ := DeployChallengeImpl(auth, backend)
	ArbChallengeBin = strings.Replace(ArbChallengeBin, "__$0b6e9c2b3f0d3b9a73d9361721242908cd$__", challengeImplAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbChallenge{ArbChallengeCaller: ArbChallengeCaller{contract: contract}, ArbChallengeTransactor: ArbChallengeTransactor{contract: contract}, ArbChallengeFilterer: ArbChallengeFilterer{contract: contract}}, nil
}

// ArbChallenge is an auto generated Go binding around an Ethereum contract.
type ArbChallenge struct {
	ArbChallengeCaller     // Read-only binding to the contract
	ArbChallengeTransactor // Write-only binding to the contract
	ArbChallengeFilterer   // Log filterer for contract events
}

// ArbChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbChallengeSession struct {
	Contract     *ArbChallenge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbChallengeCallerSession struct {
	Contract *ArbChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ArbChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbChallengeTransactorSession struct {
	Contract     *ArbChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ArbChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbChallengeRaw struct {
	Contract *ArbChallenge // Generic contract binding to access the raw methods on
}

// ArbChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbChallengeCallerRaw struct {
	Contract *ArbChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ArbChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbChallengeTransactorRaw struct {
	Contract *ArbChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbChallenge creates a new instance of ArbChallenge, bound to a specific deployed contract.
func NewArbChallenge(address common.Address, backend bind.ContractBackend) (*ArbChallenge, error) {
	contract, err := bindArbChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbChallenge{ArbChallengeCaller: ArbChallengeCaller{contract: contract}, ArbChallengeTransactor: ArbChallengeTransactor{contract: contract}, ArbChallengeFilterer: ArbChallengeFilterer{contract: contract}}, nil
}

// NewArbChallengeCaller creates a new read-only instance of ArbChallenge, bound to a specific deployed contract.
func NewArbChallengeCaller(address common.Address, caller bind.ContractCaller) (*ArbChallengeCaller, error) {
	contract, err := bindArbChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbChallengeCaller{contract: contract}, nil
}

// NewArbChallengeTransactor creates a new write-only instance of ArbChallenge, bound to a specific deployed contract.
func NewArbChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbChallengeTransactor, error) {
	contract, err := bindArbChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbChallengeTransactor{contract: contract}, nil
}

// NewArbChallengeFilterer creates a new log filterer instance of ArbChallenge, bound to a specific deployed contract.
func NewArbChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbChallengeFilterer, error) {
	contract, err := bindArbChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbChallengeFilterer{contract: contract}, nil
}

// bindArbChallenge binds a generic wrapper to an already deployed contract.
func bindArbChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbChallenge *ArbChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbChallenge.Contract.ArbChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbChallenge *ArbChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChallenge.Contract.ArbChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbChallenge *ArbChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbChallenge.Contract.ArbChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbChallenge *ArbChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbChallenge *ArbChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbChallenge *ArbChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbChallenge.Contract.contract.Transact(opts, method, params...)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_ArbChallenge *ArbChallengeTransactor) AsserterTimedOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "asserterTimedOut")
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_ArbChallenge *ArbChallengeSession) AsserterTimedOut() (*types.Transaction, error) {
	return _ArbChallenge.Contract.AsserterTimedOut(&_ArbChallenge.TransactOpts)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0xd5345e07.
//
// Solidity: function asserterTimedOut() returns()
func (_ArbChallenge *ArbChallengeTransactorSession) AsserterTimedOut() (*types.Transaction, error) {
	return _ArbChallenge.Contract.AsserterTimedOut(&_ArbChallenge.TransactOpts)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xa0294172.
//
// Solidity: function bisectAssertion(bytes32 _preData, bytes32[] _machineHashes, bytes32[] _messageAccs, bytes32[] _logAccs, uint32 _totalSteps) returns()
func (_ArbChallenge *ArbChallengeTransactor) BisectAssertion(opts *bind.TransactOpts, _preData [32]byte, _machineHashes [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _totalSteps uint32) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "bisectAssertion", _preData, _machineHashes, _messageAccs, _logAccs, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xa0294172.
//
// Solidity: function bisectAssertion(bytes32 _preData, bytes32[] _machineHashes, bytes32[] _messageAccs, bytes32[] _logAccs, uint32 _totalSteps) returns()
func (_ArbChallenge *ArbChallengeSession) BisectAssertion(_preData [32]byte, _machineHashes [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _totalSteps uint32) (*types.Transaction, error) {
	return _ArbChallenge.Contract.BisectAssertion(&_ArbChallenge.TransactOpts, _preData, _machineHashes, _messageAccs, _logAccs, _totalSteps)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0xa0294172.
//
// Solidity: function bisectAssertion(bytes32 _preData, bytes32[] _machineHashes, bytes32[] _messageAccs, bytes32[] _logAccs, uint32 _totalSteps) returns()
func (_ArbChallenge *ArbChallengeTransactorSession) BisectAssertion(_preData [32]byte, _machineHashes [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _totalSteps uint32) (*types.Transaction, error) {
	return _ArbChallenge.Contract.BisectAssertion(&_ArbChallenge.TransactOpts, _preData, _machineHashes, _messageAccs, _logAccs, _totalSteps)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_ArbChallenge *ArbChallengeTransactor) ChallengerTimedOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "challengerTimedOut")
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_ArbChallenge *ArbChallengeSession) ChallengerTimedOut() (*types.Transaction, error) {
	return _ArbChallenge.Contract.ChallengerTimedOut(&_ArbChallenge.TransactOpts)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0x635e28a7.
//
// Solidity: function challengerTimedOut() returns()
func (_ArbChallenge *ArbChallengeTransactorSession) ChallengerTimedOut() (*types.Transaction, error) {
	return _ArbChallenge.Contract.ChallengerTimedOut(&_ArbChallenge.TransactOpts)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ArbChallenge *ArbChallengeTransactor) ContinueChallenge(opts *bind.TransactOpts, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "continueChallenge", _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ArbChallenge *ArbChallengeSession) ContinueChallenge(_assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.ContinueChallenge(&_ArbChallenge.TransactOpts, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0x79d84776.
//
// Solidity: function continueChallenge(uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ArbChallenge *ArbChallengeTransactorSession) ContinueChallenge(_assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.ContinueChallenge(&_ArbChallenge.TransactOpts, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address _vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbChallenge *ArbChallengeTransactor) Init(opts *bind.TransactOpts, _vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "init", _vmAddress, _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address _vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbChallenge *ArbChallengeSession) Init(_vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.Init(&_ArbChallenge.TransactOpts, _vmAddress, _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address _vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _assertionHash) returns()
func (_ArbChallenge *ArbChallengeTransactorSession) Init(_vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.Init(&_ArbChallenge.TransactOpts, _vmAddress, _players, _escrows, _challengePeriod, _beforeHash, _beforeInbox, _timeBounds, _assertionHash)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x5552a0ff.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _afterHash, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, bytes _proof) returns()
func (_ArbChallenge *ArbChallengeTransactor) OneStepProof(opts *bind.TransactOpts, _beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _afterHash [32]byte, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ArbChallenge.contract.Transact(opts, "oneStepProof", _beforeHash, _beforeInbox, _timeBounds, _afterHash, _firstMessage, _lastMessage, _firstLog, _lastLog, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x5552a0ff.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _afterHash, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, bytes _proof) returns()
func (_ArbChallenge *ArbChallengeSession) OneStepProof(_beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _afterHash [32]byte, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.OneStepProof(&_ArbChallenge.TransactOpts, _beforeHash, _beforeInbox, _timeBounds, _afterHash, _firstMessage, _lastMessage, _firstLog, _lastLog, _proof)
}

// OneStepProof is a paid mutator transaction binding the contract method 0x5552a0ff.
//
// Solidity: function oneStepProof(bytes32 _beforeHash, bytes32 _beforeInbox, uint64[2] _timeBounds, bytes32 _afterHash, bytes32 _firstMessage, bytes32 _lastMessage, bytes32 _firstLog, bytes32 _lastLog, bytes _proof) returns()
func (_ArbChallenge *ArbChallengeTransactorSession) OneStepProof(_beforeHash [32]byte, _beforeInbox [32]byte, _timeBounds [2]uint64, _afterHash [32]byte, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ArbChallenge.Contract.OneStepProof(&_ArbChallenge.TransactOpts, _beforeHash, _beforeInbox, _timeBounds, _afterHash, _firstMessage, _lastMessage, _firstLog, _lastLog, _proof)
}

// ArbChallengeBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ArbChallenge contract.
type ArbChallengeBisectedAssertionIterator struct {
	Event *ArbChallengeBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ArbChallengeBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeBisectedAssertion)
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
		it.Event = new(ArbChallengeBisectedAssertion)
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
func (it *ArbChallengeBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeBisectedAssertion represents a BisectedAssertion event raised by the ArbChallenge contract.
type ArbChallengeBisectedAssertion struct {
	MachineHashes [][32]byte
	MessageAccs   [][32]byte
	LogAccs       [][32]byte
	TotalSteps    uint32
	Deadline      uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0xa20d3eb13e8a5d70e9b1e8d380a50a5cc4f2b7c872a24ee1a4cf37bf1de4ac6c.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint32 totalSteps, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) FilterBisectedAssertion(opts *bind.FilterOpts) (*ArbChallengeBisectedAssertionIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeBisectedAssertionIterator{contract: _ArbChallenge.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0xa20d3eb13e8a5d70e9b1e8d380a50a5cc4f2b7c872a24ee1a4cf37bf1de4ac6c.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint32 totalSteps, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ArbChallengeBisectedAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeBisectedAssertion)
				if err := _ArbChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0xa20d3eb13e8a5d70e9b1e8d380a50a5cc4f2b7c872a24ee1a4cf37bf1de4ac6c.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint32 totalSteps, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) ParseBisectedAssertion(log types.Log) (*ArbChallengeBisectedAssertion, error) {
	event := new(ArbChallengeBisectedAssertion)
	if err := _ArbChallenge.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the ArbChallenge contract.
type ArbChallengeContinuedChallengeIterator struct {
	Event *ArbChallengeContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *ArbChallengeContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeContinuedChallenge)
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
		it.Event = new(ArbChallengeContinuedChallenge)
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
func (it *ArbChallengeContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeContinuedChallenge represents a ContinuedChallenge event raised by the ArbChallenge contract.
type ArbChallengeContinuedChallenge struct {
	AssertionIndex *big.Int
	Deadline       uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0xb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a2226.
//
// Solidity: event ContinuedChallenge(uint256 assertionIndex, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) FilterContinuedChallenge(opts *bind.FilterOpts) (*ArbChallengeContinuedChallengeIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeContinuedChallengeIterator{contract: _ArbChallenge.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0xb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a2226.
//
// Solidity: event ContinuedChallenge(uint256 assertionIndex, uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *ArbChallengeContinuedChallenge) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeContinuedChallenge)
				if err := _ArbChallenge.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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
func (_ArbChallenge *ArbChallengeFilterer) ParseContinuedChallenge(log types.Log) (*ArbChallengeContinuedChallenge, error) {
	event := new(ArbChallengeContinuedChallenge)
	if err := _ArbChallenge.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ArbChallenge contract.
type ArbChallengeInitiatedChallengeIterator struct {
	Event *ArbChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ArbChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeInitiatedChallenge)
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
		it.Event = new(ArbChallengeInitiatedChallenge)
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
func (it *ArbChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the ArbChallenge contract.
type ArbChallengeInitiatedChallenge struct {
	Deadline uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f522.
//
// Solidity: event InitiatedChallenge(uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ArbChallengeInitiatedChallengeIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeInitiatedChallengeIterator{contract: _ArbChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f522.
//
// Solidity: event InitiatedChallenge(uint64 deadline)
func (_ArbChallenge *ArbChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ArbChallengeInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeInitiatedChallenge)
				if err := _ArbChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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
func (_ArbChallenge *ArbChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*ArbChallengeInitiatedChallenge, error) {
	event := new(ArbChallengeInitiatedChallenge)
	if err := _ArbChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ArbChallenge contract.
type ArbChallengeOneStepProofCompletedIterator struct {
	Event *ArbChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ArbChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeOneStepProofCompleted)
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
		it.Event = new(ArbChallengeOneStepProofCompleted)
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
func (it *ArbChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the ArbChallenge contract.
type ArbChallengeOneStepProofCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ArbChallenge *ArbChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts) (*ArbChallengeOneStepProofCompletedIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeOneStepProofCompletedIterator{contract: _ArbChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0x117efdf1fdd8be5a6ff0fb3c32333d7033bbd9523924bd0d9ca28f43540516f5.
//
// Solidity: event OneStepProofCompleted()
func (_ArbChallenge *ArbChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ArbChallengeOneStepProofCompleted) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "OneStepProofCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeOneStepProofCompleted)
				if err := _ArbChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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
func (_ArbChallenge *ArbChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*ArbChallengeOneStepProofCompleted, error) {
	event := new(ArbChallengeOneStepProofCompleted)
	if err := _ArbChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbChallengeTimedOutChallengeIterator is returned from FilterTimedOutChallenge and is used to iterate over the raw logs and unpacked data for TimedOutChallenge events raised by the ArbChallenge contract.
type ArbChallengeTimedOutChallengeIterator struct {
	Event *ArbChallengeTimedOutChallenge // Event containing the contract specifics and raw log

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
func (it *ArbChallengeTimedOutChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbChallengeTimedOutChallenge)
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
		it.Event = new(ArbChallengeTimedOutChallenge)
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
func (it *ArbChallengeTimedOutChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbChallengeTimedOutChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbChallengeTimedOutChallenge represents a TimedOutChallenge event raised by the ArbChallenge contract.
type ArbChallengeTimedOutChallenge struct {
	ChallengerWrong bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedOutChallenge is a free log retrieval operation binding the contract event 0xd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c.
//
// Solidity: event TimedOutChallenge(bool challengerWrong)
func (_ArbChallenge *ArbChallengeFilterer) FilterTimedOutChallenge(opts *bind.FilterOpts) (*ArbChallengeTimedOutChallengeIterator, error) {

	logs, sub, err := _ArbChallenge.contract.FilterLogs(opts, "TimedOutChallenge")
	if err != nil {
		return nil, err
	}
	return &ArbChallengeTimedOutChallengeIterator{contract: _ArbChallenge.contract, event: "TimedOutChallenge", logs: logs, sub: sub}, nil
}

// WatchTimedOutChallenge is a free log subscription operation binding the contract event 0xd98fd7f0b64bd4d465d83937d0742c2e61e4ed9357e65cc31936138988178f0c.
//
// Solidity: event TimedOutChallenge(bool challengerWrong)
func (_ArbChallenge *ArbChallengeFilterer) WatchTimedOutChallenge(opts *bind.WatchOpts, sink chan<- *ArbChallengeTimedOutChallenge) (event.Subscription, error) {

	logs, sub, err := _ArbChallenge.contract.WatchLogs(opts, "TimedOutChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbChallengeTimedOutChallenge)
				if err := _ArbChallenge.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
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
func (_ArbChallenge *ArbChallengeFilterer) ParseTimedOutChallenge(log types.Log) (*ArbChallengeTimedOutChallenge, error) {
	event := new(ArbChallengeTimedOutChallenge)
	if err := _ArbChallenge.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbMachineABI is the input ABI used to generate the binding from.
const ArbMachineABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"instructionStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"errHandlerHash\",\"type\":\"bytes32\"}],\"name\":\"machineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbMachineFuncSigs maps the 4-byte function signature to its string representation.
var ArbMachineFuncSigs = map[string]string{
	"c1355b59": "machineHash(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32)",
}

// ArbMachineBin is the compiled bytecode used for deploying new contracts.
var ArbMachineBin = "0x6101d6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063c1355b591461003a575b600080fd5b610075600480360360c081101561005057600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610087565b60408051918252519081900360200190f35b604080516101008101825260e081018881528152815160208181018452888252808301919091528251808201845287815282840152825180820184528681526060830152825180820184528581526080830152825190810190925282825260a0810191909152600060c08201819052906101009061010b565b979650505050505050565b600060028260c0015114156101225750600061019c565b60018260c0015114156101375750600161019c565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101205b91905056fea265627a7a7231582043036cd0ca19bbb2e4fe99d7d83ada2376007667bb60d5411626bcbcc6b4cd9e64736f6c634300050d0032"

// DeployArbMachine deploys a new Ethereum contract, binding an instance of ArbMachine to it.
func DeployArbMachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbMachine, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbMachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbMachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbMachine{ArbMachineCaller: ArbMachineCaller{contract: contract}, ArbMachineTransactor: ArbMachineTransactor{contract: contract}, ArbMachineFilterer: ArbMachineFilterer{contract: contract}}, nil
}

// ArbMachine is an auto generated Go binding around an Ethereum contract.
type ArbMachine struct {
	ArbMachineCaller     // Read-only binding to the contract
	ArbMachineTransactor // Write-only binding to the contract
	ArbMachineFilterer   // Log filterer for contract events
}

// ArbMachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbMachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbMachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbMachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbMachineSession struct {
	Contract     *ArbMachine       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbMachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbMachineCallerSession struct {
	Contract *ArbMachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbMachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbMachineTransactorSession struct {
	Contract     *ArbMachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbMachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbMachineRaw struct {
	Contract *ArbMachine // Generic contract binding to access the raw methods on
}

// ArbMachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbMachineCallerRaw struct {
	Contract *ArbMachineCaller // Generic read-only contract binding to access the raw methods on
}

// ArbMachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbMachineTransactorRaw struct {
	Contract *ArbMachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbMachine creates a new instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachine(address common.Address, backend bind.ContractBackend) (*ArbMachine, error) {
	contract, err := bindArbMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbMachine{ArbMachineCaller: ArbMachineCaller{contract: contract}, ArbMachineTransactor: ArbMachineTransactor{contract: contract}, ArbMachineFilterer: ArbMachineFilterer{contract: contract}}, nil
}

// NewArbMachineCaller creates a new read-only instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineCaller(address common.Address, caller bind.ContractCaller) (*ArbMachineCaller, error) {
	contract, err := bindArbMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbMachineCaller{contract: contract}, nil
}

// NewArbMachineTransactor creates a new write-only instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbMachineTransactor, error) {
	contract, err := bindArbMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbMachineTransactor{contract: contract}, nil
}

// NewArbMachineFilterer creates a new log filterer instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbMachineFilterer, error) {
	contract, err := bindArbMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbMachineFilterer{contract: contract}, nil
}

// bindArbMachine binds a generic wrapper to an already deployed contract.
func bindArbMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbMachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbMachine *ArbMachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbMachine.Contract.ArbMachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbMachine *ArbMachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbMachine.Contract.ArbMachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbMachine *ArbMachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbMachine.Contract.ArbMachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbMachine *ArbMachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbMachine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbMachine *ArbMachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbMachine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbMachine *ArbMachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbMachine.Contract.contract.Transact(opts, method, params...)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineCaller) MachineHash(opts *bind.CallOpts, instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbMachine.contract.Call(opts, out, "machineHash", instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
	return *ret0, err
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _ArbMachine.Contract.MachineHash(&_ArbMachine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineCallerSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _ArbMachine.Contract.MachineHash(&_ArbMachine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// ArbProtocolABI is the input ABI used to generate the binding from.
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"7ddf59d6": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32)",
	"e83f4bfe": "generateLastMessageHash(bytes)",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"85ecb92a": "generatePreconditionHash(bytes32,uint64[2],bytes32)",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x610a6e610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100555760003560e01c80624c28f61461005a5780637ddf59d6146100b257806385ecb92a146100f3578063e83f4bfe14610148575b600080fd5b6100a06004803603608081101561007057600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b03166101ee565b60408051918252519081900360200190f35b6100a0600480360360c08110156100c857600080fd5b5080359063ffffffff6020820135169060408101359060608101359060808101359060a001356102e0565b6100a06004803603608081101561010957600080fd5b60408051808201825283359392830192916060830191906020840190600290839083908082843760009201919091525091945050903591506103389050565b6100a06004803603602081101561015e57600080fd5b81019060208101813564010000000081111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460018302840111640100000000831117156101ad57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061038c945050505050565b60408051600480825260a0820190925260009160609190816020015b6102126109d2565b81526020019060019003908161020a579050509050610230866104d1565b8160008151811061023d57fe5b602002602001018190525061025a836001600160a01b031661054f565b8160018151811061026757fe5b602002602001018190525061027b8461054f565b8160028151811061028857fe5b60209081029190910101526102aa6affffffffffffffffffffff19861661054f565b816003815181106102b757fe5b60200260200101819052506102d36102ce826105cd565b61067d565b519150505b949350505050565b6040805160208082019890985260e09690961b6001600160e01b0319168682015260448601949094526064850192909252608484015260a4808401919091528151808403909101815260c49092019052805191012090565b815160209283015160408051808601969096526001600160c01b031960c093841b8116878301529190921b166048850152605080850192909252805180850390920182526070909301909252815191012090565b8051600090819081908190815b818110156104c45773__$d969135829891f807aa9c34494da4ecd99$__6389df40da88866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561040f5781810151838201526020016103f7565b50505050905090810190601f16801561043c5780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561045957600080fd5b505af415801561046d573d6000803e3d6000fd5b505050506040513d604081101561048357600080fd5b50805160209182015160408051808501999099528881018290528051808a038201815260609099019052875197909201969096209594509250600101610399565b509293505050505b919050565b6104d96109d2565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161053e565b61052b6109d2565b8152602001906001900390816105235790505b508152600260209091015292915050565b6105576109d2565b6040805160808082018352848252825190810183526000808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916105bc565b6105a96109d2565b8152602001906001900390816105a15790505b508152600060209091015292915050565b6105d56109d2565b6105df82516107b3565b610630576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b610685610a00565b6060820151600c60ff909116106106d7576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166107045760405180602001604052806106fb84600001516107ba565b905290506104cc565b606082015160ff166001141561074b5760405180602001604052806106fb8460200151600001518560200151604001518660200151606001518760200151602001516107de565b606082015160ff166002141561077057506040805160208101909152815181526104cc565b600360ff16826060015160ff161015801561079457506060820151600c60ff909116105b156107b15760405180602001604052806106fb8460400151610886565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315610838575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206102d8565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b60006008825111156108d6576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610903578160200160208202803883390190505b50805190915060005b8181101561095f5761091c610a00565b61093886838151811061092b57fe5b602002602001015161067d565b9050806000015184838151811061094b57fe5b60209081029190910101525060010161090c565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156109a8578181015183820152602001610990565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6040518060800160405280600081526020016109ec610a12565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a723158201b90d05c0f27da9305a94b86c30470b81dc58f707f5afd0bd5c7cb8b4160e70664736f6c634300050d0032"

// DeployArbProtocol deploys a new Ethereum contract, binding an instance of ArbProtocol to it.
func DeployArbProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbProtocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbProtocolBin = strings.Replace(ArbProtocolBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// ArbProtocol is an auto generated Go binding around an Ethereum contract.
type ArbProtocol struct {
	ArbProtocolCaller     // Read-only binding to the contract
	ArbProtocolTransactor // Write-only binding to the contract
	ArbProtocolFilterer   // Log filterer for contract events
}

// ArbProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbProtocolSession struct {
	Contract     *ArbProtocol      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbProtocolCallerSession struct {
	Contract *ArbProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ArbProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbProtocolTransactorSession struct {
	Contract     *ArbProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ArbProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbProtocolRaw struct {
	Contract *ArbProtocol // Generic contract binding to access the raw methods on
}

// ArbProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbProtocolCallerRaw struct {
	Contract *ArbProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ArbProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbProtocolTransactorRaw struct {
	Contract *ArbProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbProtocol creates a new instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocol(address common.Address, backend bind.ContractBackend) (*ArbProtocol, error) {
	contract, err := bindArbProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// NewArbProtocolCaller creates a new read-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolCaller(address common.Address, caller bind.ContractCaller) (*ArbProtocolCaller, error) {
	contract, err := bindArbProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolCaller{contract: contract}, nil
}

// NewArbProtocolTransactor creates a new write-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbProtocolTransactor, error) {
	contract, err := bindArbProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolTransactor{contract: contract}, nil
}

// NewArbProtocolFilterer creates a new log filterer instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbProtocolFilterer, error) {
	contract, err := bindArbProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolFilterer{contract: contract}, nil
}

// bindArbProtocol binds a generic wrapper to an already deployed contract.
func bindArbProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.ArbProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transact(opts, method, params...)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x7ddf59d6.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x7ddf59d6.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x7ddf59d6.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _messages []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHash", _messages)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _messages)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0xe83f4bfe.
//
// Solidity: function generateLastMessageHash(bytes _messages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHash(_messages []byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _messages)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x85ecb92a.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox)
}

// ArbValueABI is the input ABI used to generate the binding from.
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"startOffset\",\"type\":\"uint256\"}],\"name\":\"deserializeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidValueHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeValueHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValidValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"immediate\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueFuncSigs maps the 4-byte function signature to its string representation.
var ArbValueFuncSigs = map[string]string{
	"32e6cc21": "deserializeMessage(bytes,uint256)",
	"89df40da": "deserializeValidValueHash(bytes,uint256)",
	"8f346036": "deserializeValueHash(bytes)",
	"1f3d4d4e": "getNextValidValue(bytes,uint256)",
	"826513e0": "hashCodePoint(uint8,bool,bytes32,bytes32)",
	"53409fab": "hashCodePointBasicValue(uint8,bytes32)",
	"264f384b": "hashCodePointImmediateValue(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"1667b411": "hashIntValue(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ArbValueBin is the compiled bytecode used for deploying new contracts.
var ArbValueBin = "0x61152b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c806353409fab1161007057806353409fab14610381578063826513e0146103a757806389df40da146103db5780638f3460361461049a578063b2b9dc621461053e576100a8565b80631667b411146100ad5780631f3d4d4e146100dc578063264f384b1461020157806332e6cc211461022d578063364df27714610379575b600080fd5b6100ca600480360360208110156100c357600080fd5b503561056f565b60408051918252519081900360200190f35b610182600480360360408110156100f257600080fd5b810190602081018135600160201b81111561010c57600080fd5b82018360208201111561011e57600080fd5b803590602001918460018302840111600160201b8311171561013f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610595915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101c55781810151838201526020016101ad565b50505050905090810190601f1680156101f25780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100ca6004803603606081101561021757600080fd5b5060ff8135169060208101359060400135610619565b6102d36004803603604081101561024357600080fd5b810190602081018135600160201b81111561025d57600080fd5b82018360208201111561026f57600080fd5b803590602001918460018302840111600160201b8311171561029057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061066b915050565b604051808815151515815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610338578181015183820152602001610320565b50505050905090810190601f1680156103655780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6100ca610869565b6100ca6004803603604081101561039757600080fd5b5060ff81351690602001356108dc565b6100ca600480360360808110156103bd57600080fd5b5060ff81351690602081013515159060408101359060600135610923565b610481600480360360408110156103f157600080fd5b810190602081018135600160201b81111561040b57600080fd5b82018360208201111561041d57600080fd5b803590602001918460018302840111600160201b8311171561043e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506109cc915050565b6040805192835260208301919091528051918290030190f35b6100ca600480360360208110156104b057600080fd5b810190602081018135600160201b8111156104ca57600080fd5b8201836020820111156104dc57600080fd5b803590602001918460018302840111600160201b831117156104fd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a45945050505050565b61055b6004803603602081101561055457600080fd5b5035610ab7565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b600060606000806105a461146f565b6105ae8787610abe565b9194509250905082156105f6576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b8161060a888880840363ffffffff610c4816565b945094505050505b9250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6000806000806000806060600088965060008a888151811061068957fe5b016020015160019098019760f81c9050600781146106bb576106ae8b60018a036109cc565b909850965061085d915050565b6106c58b896109cc565b90985091506106e48b60018c016000198d8c030163ffffffff610c4816565b92508a88815181106106f257fe5b016020015160019098019760f81c90508015610715576106ae8b60018a036109cc565b61071f8b89610cc8565b80995081975050508a888151811061073357fe5b016020015160019098019760f81c90508015610756576106ae8b60018a036109cc565b6107608b89610cc8565b80995081965050508a888151811061077457fe5b016020015160019098019760f81c90508015610797576106ae8b60018a036109cc565b6107a18b89610cc8565b60408051600480825260a0820190925260019c50919a509195506060916020820160808038833901905050905082816000815181106107dc57fe5b6020026020010181815250506107f18761056f565b816001815181106107fe57fe5b6020026020010181815250506108138661056f565b8160028151811061082057fe5b6020026020010181815250506108358561056f565b8160038151811061084257fe5b60200260200101818152505061085781610cef565b97505050505b92959891949750929550565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156108b557818101518382015260200161089d565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000831561097d575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206109c4565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b6000806000806109da61146f565b6109e48787610abe565b919450925090508215610a2c576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b81610a3682610daf565b51909890975095505050505050565b60008080610a5161146f565b610a5c856000610abe565b919450925090508215610aa4576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b610aad81610daf565b5195945050505050565b6008101590565b600080610ac961146f565b84518410610b1e576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110610b3157fe5b016020015160019092019160f81c90506000610b4b61149d565b60ff8316610b7f57610b5d8985610cc8565b9094509150600084610b6e84610ee5565b91985096509450610c419350505050565b60ff831660011415610ba657610b958985610f63565b9094509050600084610b6e836110be565b60ff831660021415610bcd57610bbc8985610cc8565b9094509150600084610b6e8461111e565b600360ff841610801590610be45750600c60ff8416105b15610c2157600219830160606000610bfd838d8961119c565b909850925090508087610c0f84611257565b99509950995050505050505050610c41565b8260ff16612710016000610c356000610ee5565b91985096509450505050505b9250925092565b606081830184511015610c5a57600080fd5b606082158015610c7557604051915060208201604052610cbf565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610cae578051835260209283019201610c96565b5050858452601f01601f1916604052505b50949350505050565b6000808281610cdd868363ffffffff61130716565b60209290920196919550909350505050565b6000600882511115610d3f576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b8151600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610d87578181015183820152602001610d6f565b5050505090500192505050604051602081830303815290604052805190602001209050919050565b610db76114c4565b6060820151600c60ff90911610610e09576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610e36576040518060200160405280610e2d846000015161056f565b90529050610590565b606082015160ff1660011415610e7d576040518060200160405280610e2d846020015160000151856020015160400151866020015160600151876020015160200151610923565b606082015160ff1660021415610ea25750604080516020810190915281518152610590565b600360ff16826060015160ff1610158015610ec657506060820151600c60ff909116105b15610ee3576040518060200160405280610e2d8460400151611323565bfe5b610eed61146f565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610f52565b610f3f61146f565b815260200190600190039081610f375790505b508152600060209091015292915050565b6000610f6d61149d565b60008390506000858281518110610f8057fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110610fa657fe5b016020015160019384019360f89190911c915060009060ff84161415611032576000610fd061146f565b610fda8a87610abe565b90975090925090508115611023576040805162461bcd60e51b815260206004820152601e60248201526000805160206114d7833981519152604482015290519081900360640190fd5b61102c81610daf565b51925050505b6000611044898663ffffffff61130716565b90506020850194508360ff1660011415611089576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506106129050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6110c661146f565b60408051608081018252600080825260208083018690528351828152908101845291928301919061110d565b6110fa61146f565b8152602001906001900390816110f25790505b508152600160209091015292915050565b61112661146f565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161118b565b61117861146f565b8152602001906001900390816111705790505b508152600260209091015292915050565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156111e757816020015b6111d461146f565b8152602001906001900390816111cc5790505b50905060005b8960ff168160ff161015611241576112058985610abe565b8451859060ff861690811061121657fe5b60209081029190910101529450925082156112395750909450909250905061124e565b6001016111ed565b5060009550919350909150505b93509350939050565b61125f61146f565b6112698251610ab7565b6112ba576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6000816020018351101561131a57600080fd5b50016020015190565b6000600882511115611373576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156113a0578160200160208202803883390190505b50805190915060005b818110156113fc576113b96114c4565b6113d58683815181106113c857fe5b6020026020010151610daf565b905080600001518483815181106113e857fe5b6020908102919091010152506001016113a9565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561144557818101518382015260200161142d565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180608001604052806000815260200161148961149d565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4d61727368616c6c65642076616c7565206d7573742062652076616c69640000a265627a7a72315820184abb01cd16a80d27b9e3125b6e33c2f53720dceececcf7c7182d2f0bcee9ee64736f6c634300050d0032"

// DeployArbValue deploys a new Ethereum contract, binding an instance of ArbValue to it.
func DeployArbValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbValue, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// ArbValue is an auto generated Go binding around an Ethereum contract.
type ArbValue struct {
	ArbValueCaller     // Read-only binding to the contract
	ArbValueTransactor // Write-only binding to the contract
	ArbValueFilterer   // Log filterer for contract events
}

// ArbValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbValueSession struct {
	Contract     *ArbValue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbValueCallerSession struct {
	Contract *ArbValueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbValueTransactorSession struct {
	Contract     *ArbValueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbValueRaw struct {
	Contract *ArbValue // Generic contract binding to access the raw methods on
}

// ArbValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbValueCallerRaw struct {
	Contract *ArbValueCaller // Generic read-only contract binding to access the raw methods on
}

// ArbValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbValueTransactorRaw struct {
	Contract *ArbValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbValue creates a new instance of ArbValue, bound to a specific deployed contract.
func NewArbValue(address common.Address, backend bind.ContractBackend) (*ArbValue, error) {
	contract, err := bindArbValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// NewArbValueCaller creates a new read-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueCaller(address common.Address, caller bind.ContractCaller) (*ArbValueCaller, error) {
	contract, err := bindArbValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueCaller{contract: contract}, nil
}

// NewArbValueTransactor creates a new write-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbValueTransactor, error) {
	contract, err := bindArbValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueTransactor{contract: contract}, nil
}

// NewArbValueFilterer creates a new log filterer instance of ArbValue, bound to a specific deployed contract.
func NewArbValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbValueFilterer, error) {
	contract, err := bindArbValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbValueFilterer{contract: contract}, nil
}

// bindArbValue binds a generic wrapper to an already deployed contract.
func bindArbValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.ArbValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transact(opts, method, params...)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_ArbValue *ArbValueCaller) DeserializeMessage(opts *bind.CallOpts, data []byte, startOffset *big.Int) (struct {
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
	err := _ArbValue.contract.Call(opts, out, "deserializeMessage", data, startOffset)
	return *ret, err
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_ArbValue *ArbValueSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _ArbValue.Contract.DeserializeMessage(&_ArbValue.CallOpts, data, startOffset)
}

// DeserializeMessage is a free data retrieval call binding the contract method 0x32e6cc21.
//
// Solidity: function deserializeMessage(bytes data, uint256 startOffset) constant returns(bool valid, uint256 offset, bytes32 messageHash, uint256 destination, uint256 value, uint256 tokenType, bytes messageData)
func (_ArbValue *ArbValueCallerSession) DeserializeMessage(data []byte, startOffset *big.Int) (struct {
	Valid       bool
	Offset      *big.Int
	MessageHash [32]byte
	Destination *big.Int
	Value       *big.Int
	TokenType   *big.Int
	MessageData []byte
}, error) {
	return _ArbValue.Contract.DeserializeMessage(&_ArbValue.CallOpts, data, startOffset)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValidValueHash(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "deserializeValidValueHash", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValueHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "deserializeValueHash", data)
	return *ret0, err
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCaller) GetNextValidValue(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "getNextValidValue", data, offset)
	return *ret0, *ret1, err
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCallerSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePoint(opts *bind.CallOpts, opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePoint", opcode, immediate, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePoint(&_ArbValue.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePoint is a free data retrieval call binding the contract method 0x826513e0.
//
// Solidity: function hashCodePoint(uint8 opcode, bool immediate, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePoint(opcode uint8, immediate bool, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePoint(&_ArbValue.CallOpts, opcode, immediate, immediateVal, nextCodePoint)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointBasicValue(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointBasicValue", opcode, nextCodePoint)
	return *ret0, err
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointImmediateValue(opts *bind.CallOpts, opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointImmediateValue", opcode, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashIntValue(opts *bind.CallOpts, val *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashIntValue", val)
	return *ret0, err
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCaller) IsValidTupleSize(opts *bind.CallOpts, size *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "isValidTupleSize", size)
	return *ret0, err
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCallerSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820085bb168467a1b1b9144d77c586e271e631429488c107be4f03f7ca8146dd95064736f6c634300050d0032"

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
const ChallengeABI = "[]"

// ChallengeBin is the compiled bytecode used for deploying new contracts.
var ChallengeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820ca2ff07d848317e18aa9377faa0dba77ac4bcdda8cc56590dca63cae921b850264736f6c634300050d0032"

// DeployChallenge deploys a new Ethereum contract, binding an instance of Challenge to it.
func DeployChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Challenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

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

// ChallengeImplABI is the input ABI used to generate the binding from.
const ChallengeImplABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"machineHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"logAccs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assertionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"}]"

// ChallengeImplFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeImplFuncSigs = map[string]string{
	"2716a27b": "bisectAssertion(Challenge.Data storage,bytes32,bytes32[],bytes32[],bytes32[],uint32)",
	"110112ae": "continueChallenge(Challenge.Data storage,uint256,bytes,bytes32,bytes32)",
	"52426fc3": "initializeChallenge(Challenge.Data storage,address,address[2],uint128[2],uint32,bytes32,bytes32,uint64[2],bytes32)",
	"6e4a55a9": "oneStepProof(Challenge.Data storage,bytes32,bytes32,uint64[2],bytes32,bytes32,bytes32,bytes32,bytes32,bytes)",
}

// ChallengeImplBin is the compiled bytecode used for deploying new contracts.
var ChallengeImplBin = "0x6119d4610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c8063110112ae1461005b5780632716a27b1461011f57806352426fc3146102e55780636e4a55a9146103c1575b600080fd5b81801561006757600080fd5b5061011d600480360360a081101561007e57600080fd5b813591602081013591810190606081016040820135600160201b8111156100a457600080fd5b8201836020820111156100b657600080fd5b803590602001918460018302840111600160201b831117156100d757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356104c0565b005b81801561012b57600080fd5b5061011d600480360360c081101561014257600080fd5b813591602081013591810190606081016040820135600160201b81111561016857600080fd5b82018360208201111561017a57600080fd5b803590602001918460208302840111600160201b8311171561019b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101ea57600080fd5b8201836020820111156101fc57600080fd5b803590602001918460208302840111600160201b8311171561021d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561026c57600080fd5b82018360208201111561027e57600080fd5b803590602001918460208302840111600160201b8311171561029f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505050903563ffffffff1691506108f49050565b8180156102f157600080fd5b5061011d600480360361018081101561030957600080fd5b6040805180820182528335936001600160a01b036020820135169381019290916080830191808401906002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525050604080518082018252929563ffffffff853516956020860135958381013595929450909260a08201929160600190600290839083908082843760009201919091525091945050903591506112109050565b61011d60048036036101608110156103d857600080fd5b60408051808201825283359360208101359383820135939082019260a08301916060840190600290839083908082843760009201919091525091948335946020850135946040810135945060608101359350608081013592919060c081019060a00135600160201b81111561044c57600080fd5b82018360208201111561045e57600080fd5b803590602001918460018302840111600160201b8311171561047f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113e8945050505050565b60026005860154600160601b900460ff1660028111156104dc57fe5b1460405180604001604052806009815260200168434f4e5f535441544560b81b815250906105885760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561054d578181015183820152602001610535565b50505050905090810190601f16801561057a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50846001015482146040518060400160405280600881526020016721a7a72fa82922ab60c11b815250906105fd5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b50600585015460408051808201909152600c81526b434f4e5f444541444c494e4560a01b60208201529067ffffffffffffffff1643111561067f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b50600385016001015460408051808201909152600a81526921a7a72fa9a2a72222a960b11b6020820152906001600160a01b031633146107005760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b5073__$800fcb2f4a98daa165a5cdb21a355d7a15$__63b792d767848484886001016040518563ffffffff1660e01b81526004018080602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b83811015610780578181015183820152602001610768565b50505050905090810190601f1680156107ad5780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b1580156107cd57600080fd5b505af41580156107e1573d6000803e3d6000fd5b505050506040513d60208110156107f757600080fd5b505160408051808201909152600981526821a7a72fa82927a7a360b91b6020820152906108655760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b5060058501805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160601b17938404160190811667ffffffffffffffff199290921682179092556001870183905560408051878152602081019290925280517fb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a22269281900390910190a1505050505050565b83518351811461090357600080fd5b8251811461091057600080fd5b6005870154600160601b900460ff16600281111561092a57fe5b600114604051806040016040528060098152602001684249535f535441544560b81b8152509061099b5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b50600587015460408051808201909152600c81526b4249535f444541444c494e4560a01b60208201529067ffffffffffffffff16431115610a1d5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b50600387016000015460408051808201909152600a8152692124a9afa9a2a72222a960b11b6020820152906001600160a01b03163314610a9e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b5086600101548686600081518110610ab257fe5b602002602001015173__$9836fa7140e5a33041d4b827682e675a30$__637ddf59d6896001870381518110610ae357fe5b6020026020010151878a600081518110610af957fe5b60200260200101518b60018a0381518110610b1057fe5b60200260200101518b600081518110610b2557fe5b60200260200101518c60018c0381518110610b3c57fe5b60200260200101516040518763ffffffff1660e01b8152600401808781526020018663ffffffff1663ffffffff168152602001858152602001848152602001838152602001828152602001965050505050505060206040518083038186803b158015610ba757600080fd5b505af4158015610bbb573d6000803e3d6000fd5b505050506040513d6020811015610bd157600080fd5b5051604080516020818101959095528082019390935260608084019290925280518084039092018252608083018082528251929094019190912060c0830190915260088352672124a9afa82922ab60c11b60a090920191909152909114610c795760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b50606081604051908082528060200260200182016040528015610ca6578160200160208202803883390190505b5090508686600081518110610cb757fe5b602002602001015173__$9836fa7140e5a33041d4b827682e675a30$__637ddf59d689600181518110610ce657fe5b60200260200101518663ffffffff168863ffffffff1681610d0357fe5b068763ffffffff168963ffffffff1681610d1957fe5b04018a600081518110610d2857fe5b60200260200101518b600181518110610d3d57fe5b60200260200101518b600081518110610d5257fe5b60200260200101518c600181518110610d6757fe5b60200260200101516040518763ffffffff1660e01b8152600401808781526020018663ffffffff1663ffffffff168152602001858152602001848152602001838152602001828152602001965050505050505060206040518083038186803b158015610dd257600080fd5b505af4158015610de6573d6000803e3d6000fd5b505050506040513d6020811015610dfc57600080fd5b50516040805160208181019590955280820193909352606080840192909252805180840390920182526080909201909152805191012081518290600090610e3f57fe5b602090810291909101015260015b82811015610feb5787878281518110610e6257fe5b602002602001015173__$9836fa7140e5a33041d4b827682e675a30$__637ddf59d68a8560010181518110610e9357fe5b60200260200101518763ffffffff168963ffffffff1681610eb057fe5b048b8781518110610ebd57fe5b60200260200101518c8860010181518110610ed457fe5b60200260200101518c8981518110610ee857fe5b60200260200101518d8a60010181518110610eff57fe5b60200260200101516040518763ffffffff1660e01b8152600401808781526020018663ffffffff1663ffffffff168152602001858152602001848152602001838152602001828152602001965050505050505060206040518083038186803b158015610f6a57600080fd5b505af4158015610f7e573d6000803e3d6000fd5b505050506040513d6020811015610f9457600080fd5b5051604080516020818101959095528082019390935260608084019290925280518084039092018252608090920190915280519101208251839083908110610fd857fe5b6020908102919091010152600101610e4d565b506040516309898dc160e41b815260206004820181815283516024840152835173__$800fcb2f4a98daa165a5cdb21a355d7a15$__93639898dc1093869392839260440191808601910280838360005b8381101561105357818101518382015260200161103b565b505050509050019250505060206040518083038186803b15801561107657600080fd5b505af415801561108a573d6000803e3d6000fd5b505050506040513d60208110156110a057600080fd5b5051600189015560058801805467ffffffffffffffff1960ff60601b19909116600160611b17908116600160401b90910463ffffffff908116430167ffffffffffffffff8116928317909355604080519187166060830152608082019290925260a080825289519082015288517fa20d3eb13e8a5d70e9b1e8d380a50a5cc4f2b7c872a24ee1a4cf37bf1de4ac6c928a928a928a928a9288929091829160208084019284019160c0850191808c01910280838360005b8381101561116e578181015183820152602001611156565b50505050905001848103835288818151815260200191508051906020019060200280838360005b838110156111ad578181015183820152602001611195565b50505050905001848103825287818151815260200191508051906020019060200280838360005b838110156111ec5781810151838201526020016111d4565b505050509050019850505050505050505060405180910390a1505050505050505050565b600060058a0154600160601b900460ff16600281111561122c57fe5b146040518060400160405280600f81526020016e4348414c5f494e49545f535441544560881b815250906112a15760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b5088546001600160a01b0319166001600160a01b0389161789558151602080840151604080516001600160c01b031960c095861b8116828601529290941b909116602884015260308084018790528151808503909101815260508401825280519083012060708401526090830187905260b08084018590528151808503909101815260d09093019052815191012060018a015563ffffffff8516430161134d6002808c01908990611852565b5061135d60038b018960026118f7565b5060058a01805467ffffffffffffffff191667ffffffffffffffff83169081176bffffffff00000000000000001916600160401b63ffffffff8a16021760ff60601b1916600160601b1790915560408051918252517f932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f5229181900360200190a150505050505050505050565b600160058b0154600160601b900460ff16600281111561140457fe5b14604051806040016040528060098152602001684f53505f535441544560b81b815250906114735760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b5060058a015460408051808201909152600c81526b4f53505f444541444c494e4560a01b60208201529067ffffffffffffffff164311156114f55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b5060018a81015488516020808b0151604080516001600160c01b031960c095861b8116828601529290941b909116602884015260308084018e90528151808503909101815260508401808352815191840191909120633eefaceb60e11b909152605484018c90526074840195909552609483018a905260b4830189905260d4830188905260f48301879052519293928d9273__$9836fa7140e5a33041d4b827682e675a30$__92637ddf59d69261011480840193829003018186803b1580156115bd57600080fd5b505af41580156115d1573d6000803e3d6000fd5b505050506040513d60208110156115e757600080fd5b5051604080516020818101959095528082019390935260608084019290925280518084039092018252608083018082528251929094019190912060c08301909152600883526727a9a82fa82922ab60c11b60a09092019190915290911461168f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b50600073__$f55f7f918072b72dcc999cdc8e581605f5$__63d82f8fd58b8a8c8b8b8b8b8b8b6040518a63ffffffff1660e01b8152600401808a815260200189600260200280838360005b838110156116f25781810151838201526020016116da565b5050505090500188815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561175657818101518382015260200161173e565b50505050905090810190601f1680156117835780820380516001836020036101000a031916815260200191505b509a505050505050505050505060206040518083038186803b1580156117a857600080fd5b505af41580156117bc573d6000803e3d6000fd5b505050506040513d60208110156117d257600080fd5b505160408051808201909152600981526827a9a82fa82927a7a360b91b602082015290915081156118445760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561054d578181015183820152602001610535565b505050505050505050505050565b6001830191839082156118e75791602002820160005b838211156118b257835183826101000a8154816001600160801b0302191690836001600160801b031602179055509260200192601001602081600f01049283019260010302611868565b80156118e55782816101000a8154906001600160801b030219169055601001602081600f010492830192600103026118b2565b505b506118f392915061194b565b5090565b826002810192821561193f579160200282015b8281111561193f57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061190a565b506118f392915061197b565b61197891905b808211156118f35780546fffffffffffffffffffffffffffffffff19168155600101611951565b90565b61197891905b808211156118f35780546001600160a01b031916815560010161198156fea265627a7a72315820b09384f4b11ff6c083f4907f4ad4f4f8e4e561e1fe285a88f8d58f57bf13d09564736f6c634300050d0032"

// DeployChallengeImpl deploys a new Ethereum contract, binding an instance of ChallengeImpl to it.
func DeployChallengeImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeImpl, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	merkleLibAddr, _, _, _ := DeployMerkleLib(auth, backend)
	ChallengeImplBin = strings.Replace(ChallengeImplBin, "__$800fcb2f4a98daa165a5cdb21a355d7a15$__", merkleLibAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	ChallengeImplBin = strings.Replace(ChallengeImplBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	oneStepProofAddr, _, _, _ := DeployOneStepProof(auth, backend)
	ChallengeImplBin = strings.Replace(ChallengeImplBin, "__$f55f7f918072b72dcc999cdc8e581605f5$__", oneStepProofAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeImpl{ChallengeImplCaller: ChallengeImplCaller{contract: contract}, ChallengeImplTransactor: ChallengeImplTransactor{contract: contract}, ChallengeImplFilterer: ChallengeImplFilterer{contract: contract}}, nil
}

// ChallengeImpl is an auto generated Go binding around an Ethereum contract.
type ChallengeImpl struct {
	ChallengeImplCaller     // Read-only binding to the contract
	ChallengeImplTransactor // Write-only binding to the contract
	ChallengeImplFilterer   // Log filterer for contract events
}

// ChallengeImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeImplSession struct {
	Contract     *ChallengeImpl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeImplCallerSession struct {
	Contract *ChallengeImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ChallengeImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeImplTransactorSession struct {
	Contract     *ChallengeImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ChallengeImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeImplRaw struct {
	Contract *ChallengeImpl // Generic contract binding to access the raw methods on
}

// ChallengeImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeImplCallerRaw struct {
	Contract *ChallengeImplCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeImplTransactorRaw struct {
	Contract *ChallengeImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeImpl creates a new instance of ChallengeImpl, bound to a specific deployed contract.
func NewChallengeImpl(address common.Address, backend bind.ContractBackend) (*ChallengeImpl, error) {
	contract, err := bindChallengeImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeImpl{ChallengeImplCaller: ChallengeImplCaller{contract: contract}, ChallengeImplTransactor: ChallengeImplTransactor{contract: contract}, ChallengeImplFilterer: ChallengeImplFilterer{contract: contract}}, nil
}

// NewChallengeImplCaller creates a new read-only instance of ChallengeImpl, bound to a specific deployed contract.
func NewChallengeImplCaller(address common.Address, caller bind.ContractCaller) (*ChallengeImplCaller, error) {
	contract, err := bindChallengeImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeImplCaller{contract: contract}, nil
}

// NewChallengeImplTransactor creates a new write-only instance of ChallengeImpl, bound to a specific deployed contract.
func NewChallengeImplTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeImplTransactor, error) {
	contract, err := bindChallengeImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeImplTransactor{contract: contract}, nil
}

// NewChallengeImplFilterer creates a new log filterer instance of ChallengeImpl, bound to a specific deployed contract.
func NewChallengeImplFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeImplFilterer, error) {
	contract, err := bindChallengeImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeImplFilterer{contract: contract}, nil
}

// bindChallengeImpl binds a generic wrapper to an already deployed contract.
func bindChallengeImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeImpl *ChallengeImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeImpl.Contract.ChallengeImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeImpl *ChallengeImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeImpl.Contract.ChallengeImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeImpl *ChallengeImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeImpl.Contract.ChallengeImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeImpl *ChallengeImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeImpl *ChallengeImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeImpl *ChallengeImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeImpl.Contract.contract.Transact(opts, method, params...)
}

// ChallengeImplBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ChallengeImpl contract.
type ChallengeImplBisectedAssertionIterator struct {
	Event *ChallengeImplBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ChallengeImplBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeImplBisectedAssertion)
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
		it.Event = new(ChallengeImplBisectedAssertion)
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
func (it *ChallengeImplBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeImplBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeImplBisectedAssertion represents a BisectedAssertion event raised by the ChallengeImpl contract.
type ChallengeImplBisectedAssertion struct {
	MachineHashes [][32]byte
	MessageAccs   [][32]byte
	LogAccs       [][32]byte
	TotalSteps    uint32
	Deadline      uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0xa20d3eb13e8a5d70e9b1e8d380a50a5cc4f2b7c872a24ee1a4cf37bf1de4ac6c.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint32 totalSteps, uint64 deadline)
func (_ChallengeImpl *ChallengeImplFilterer) FilterBisectedAssertion(opts *bind.FilterOpts) (*ChallengeImplBisectedAssertionIterator, error) {

	logs, sub, err := _ChallengeImpl.contract.FilterLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return &ChallengeImplBisectedAssertionIterator{contract: _ChallengeImpl.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0xa20d3eb13e8a5d70e9b1e8d380a50a5cc4f2b7c872a24ee1a4cf37bf1de4ac6c.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint32 totalSteps, uint64 deadline)
func (_ChallengeImpl *ChallengeImplFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ChallengeImplBisectedAssertion) (event.Subscription, error) {

	logs, sub, err := _ChallengeImpl.contract.WatchLogs(opts, "BisectedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeImplBisectedAssertion)
				if err := _ChallengeImpl.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0xa20d3eb13e8a5d70e9b1e8d380a50a5cc4f2b7c872a24ee1a4cf37bf1de4ac6c.
//
// Solidity: event BisectedAssertion(bytes32[] machineHashes, bytes32[] messageAccs, bytes32[] logAccs, uint32 totalSteps, uint64 deadline)
func (_ChallengeImpl *ChallengeImplFilterer) ParseBisectedAssertion(log types.Log) (*ChallengeImplBisectedAssertion, error) {
	event := new(ChallengeImplBisectedAssertion)
	if err := _ChallengeImpl.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeImplContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the ChallengeImpl contract.
type ChallengeImplContinuedChallengeIterator struct {
	Event *ChallengeImplContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeImplContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeImplContinuedChallenge)
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
		it.Event = new(ChallengeImplContinuedChallenge)
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
func (it *ChallengeImplContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeImplContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeImplContinuedChallenge represents a ContinuedChallenge event raised by the ChallengeImpl contract.
type ChallengeImplContinuedChallenge struct {
	AssertionIndex *big.Int
	Deadline       uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0xb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a2226.
//
// Solidity: event ContinuedChallenge(uint256 assertionIndex, uint64 deadline)
func (_ChallengeImpl *ChallengeImplFilterer) FilterContinuedChallenge(opts *bind.FilterOpts) (*ChallengeImplContinuedChallengeIterator, error) {

	logs, sub, err := _ChallengeImpl.contract.FilterLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return &ChallengeImplContinuedChallengeIterator{contract: _ChallengeImpl.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0xb6e6007d4a77b32247ea255a84a9dd91877135bc31b6b0ff94f2f04b895a2226.
//
// Solidity: event ContinuedChallenge(uint256 assertionIndex, uint64 deadline)
func (_ChallengeImpl *ChallengeImplFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeImplContinuedChallenge) (event.Subscription, error) {

	logs, sub, err := _ChallengeImpl.contract.WatchLogs(opts, "ContinuedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeImplContinuedChallenge)
				if err := _ChallengeImpl.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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
func (_ChallengeImpl *ChallengeImplFilterer) ParseContinuedChallenge(log types.Log) (*ChallengeImplContinuedChallenge, error) {
	event := new(ChallengeImplContinuedChallenge)
	if err := _ChallengeImpl.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeImplInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ChallengeImpl contract.
type ChallengeImplInitiatedChallengeIterator struct {
	Event *ChallengeImplInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeImplInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeImplInitiatedChallenge)
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
		it.Event = new(ChallengeImplInitiatedChallenge)
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
func (it *ChallengeImplInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeImplInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeImplInitiatedChallenge represents a InitiatedChallenge event raised by the ChallengeImpl contract.
type ChallengeImplInitiatedChallenge struct {
	Deadline uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f522.
//
// Solidity: event InitiatedChallenge(uint64 deadline)
func (_ChallengeImpl *ChallengeImplFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*ChallengeImplInitiatedChallengeIterator, error) {

	logs, sub, err := _ChallengeImpl.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &ChallengeImplInitiatedChallengeIterator{contract: _ChallengeImpl.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x932aaaa4c63661003fe646e4fe0c16b945036ea025c7a995b06411b50864f522.
//
// Solidity: event InitiatedChallenge(uint64 deadline)
func (_ChallengeImpl *ChallengeImplFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeImplInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _ChallengeImpl.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeImplInitiatedChallenge)
				if err := _ChallengeImpl.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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
func (_ChallengeImpl *ChallengeImplFilterer) ParseInitiatedChallenge(log types.Log) (*ChallengeImplInitiatedChallenge, error) {
	event := new(ChallengeImplInitiatedChallenge)
	if err := _ChallengeImpl.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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
var DebugPrintBin = "0x610202610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063252fb38d1461003a575b600080fd5b6100576004803603602081101561005057600080fd5b50356100cc565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610091578181015183820152602001610079565b50505050905090810190601f1680156100be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051818152606081810183529182919060208201818038833901905050905060005b602081101561019457600084826020811061010757fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b61012d8261019b565b85856002028151811061013c57fe5b60200101906001600160f81b031916908160001a90535061015c8161019b565b85856002026001018151811061016e57fe5b60200101906001600160f81b031916908160001a90535050600190920191506100f09050565b5092915050565b6000600a60f883901c10156101bb578160f81c60300160f81b90506101c8565b8160f81c60570160f81b90505b91905056fea265627a7a7231582066b66e0573d8b50e4f34a432fb74922dc886db7d72dedc7491b12fc19e449f1664736f6c634300050d0032"

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

// IArbChallengeABI is the input ABI used to generate the binding from.
const IArbChallengeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"internalType\":\"uint32\",\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbChallengeFuncSigs maps the 4-byte function signature to its string representation.
var IArbChallengeFuncSigs = map[string]string{
	"2820245a": "init(address,address[2],uint128[2],uint32,bytes32,bytes32,uint64[2],bytes32)",
}

// IArbChallenge is an auto generated Go binding around an Ethereum contract.
type IArbChallenge struct {
	IArbChallengeCaller     // Read-only binding to the contract
	IArbChallengeTransactor // Write-only binding to the contract
	IArbChallengeFilterer   // Log filterer for contract events
}

// IArbChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbChallengeSession struct {
	Contract     *IArbChallenge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbChallengeCallerSession struct {
	Contract *IArbChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IArbChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbChallengeTransactorSession struct {
	Contract     *IArbChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IArbChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbChallengeRaw struct {
	Contract *IArbChallenge // Generic contract binding to access the raw methods on
}

// IArbChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbChallengeCallerRaw struct {
	Contract *IArbChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// IArbChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbChallengeTransactorRaw struct {
	Contract *IArbChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbChallenge creates a new instance of IArbChallenge, bound to a specific deployed contract.
func NewIArbChallenge(address common.Address, backend bind.ContractBackend) (*IArbChallenge, error) {
	contract, err := bindIArbChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbChallenge{IArbChallengeCaller: IArbChallengeCaller{contract: contract}, IArbChallengeTransactor: IArbChallengeTransactor{contract: contract}, IArbChallengeFilterer: IArbChallengeFilterer{contract: contract}}, nil
}

// NewIArbChallengeCaller creates a new read-only instance of IArbChallenge, bound to a specific deployed contract.
func NewIArbChallengeCaller(address common.Address, caller bind.ContractCaller) (*IArbChallengeCaller, error) {
	contract, err := bindIArbChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChallengeCaller{contract: contract}, nil
}

// NewIArbChallengeTransactor creates a new write-only instance of IArbChallenge, bound to a specific deployed contract.
func NewIArbChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbChallengeTransactor, error) {
	contract, err := bindIArbChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbChallengeTransactor{contract: contract}, nil
}

// NewIArbChallengeFilterer creates a new log filterer instance of IArbChallenge, bound to a specific deployed contract.
func NewIArbChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbChallengeFilterer, error) {
	contract, err := bindIArbChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbChallengeFilterer{contract: contract}, nil
}

// bindIArbChallenge binds a generic wrapper to an already deployed contract.
func bindIArbChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChallenge *IArbChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChallenge.Contract.IArbChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChallenge *IArbChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChallenge.Contract.IArbChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChallenge *IArbChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChallenge.Contract.IArbChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbChallenge *IArbChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbChallenge *IArbChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbChallenge *IArbChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbChallenge.Contract.contract.Transact(opts, method, params...)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 beforeHash, bytes32 beforeInbox, uint64[2] timeBounds, bytes32 _assertionHash) returns()
func (_IArbChallenge *IArbChallengeTransactor) Init(opts *bind.TransactOpts, vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, beforeHash [32]byte, beforeInbox [32]byte, timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IArbChallenge.contract.Transact(opts, "init", vmAddress, _players, _escrows, _challengePeriod, beforeHash, beforeInbox, timeBounds, _assertionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 beforeHash, bytes32 beforeInbox, uint64[2] timeBounds, bytes32 _assertionHash) returns()
func (_IArbChallenge *IArbChallengeSession) Init(vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, beforeHash [32]byte, beforeInbox [32]byte, timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IArbChallenge.Contract.Init(&_IArbChallenge.TransactOpts, vmAddress, _players, _escrows, _challengePeriod, beforeHash, beforeInbox, timeBounds, _assertionHash)
}

// Init is a paid mutator transaction binding the contract method 0x2820245a.
//
// Solidity: function init(address vmAddress, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 beforeHash, bytes32 beforeInbox, uint64[2] timeBounds, bytes32 _assertionHash) returns()
func (_IArbChallenge *IArbChallengeTransactorSession) Init(vmAddress common.Address, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, beforeHash [32]byte, beforeInbox [32]byte, timeBounds [2]uint64, _assertionHash [32]byte) (*types.Transaction, error) {
	return _IArbChallenge.Contract.Init(&_IArbChallenge.TransactOpts, vmAddress, _players, _escrows, _challengePeriod, beforeHash, beforeInbox, timeBounds, _assertionHash)
}

// IArbitrumVMABI is the input ABI used to generate the binding from.
const IArbitrumVMABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_players\",\"type\":\"address[2]\"},{\"internalType\":\"uint128[2]\",\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArbitrumVMFuncSigs maps the 4-byte function signature to its string representation.
var IArbitrumVMFuncSigs = map[string]string{
	"22c091bc": "completeChallenge(address[2],uint128[2])",
}

// IArbitrumVM is an auto generated Go binding around an Ethereum contract.
type IArbitrumVM struct {
	IArbitrumVMCaller     // Read-only binding to the contract
	IArbitrumVMTransactor // Write-only binding to the contract
	IArbitrumVMFilterer   // Log filterer for contract events
}

// IArbitrumVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArbitrumVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbitrumVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArbitrumVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbitrumVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArbitrumVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArbitrumVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArbitrumVMSession struct {
	Contract     *IArbitrumVM      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArbitrumVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArbitrumVMCallerSession struct {
	Contract *IArbitrumVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IArbitrumVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArbitrumVMTransactorSession struct {
	Contract     *IArbitrumVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IArbitrumVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArbitrumVMRaw struct {
	Contract *IArbitrumVM // Generic contract binding to access the raw methods on
}

// IArbitrumVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArbitrumVMCallerRaw struct {
	Contract *IArbitrumVMCaller // Generic read-only contract binding to access the raw methods on
}

// IArbitrumVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArbitrumVMTransactorRaw struct {
	Contract *IArbitrumVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArbitrumVM creates a new instance of IArbitrumVM, bound to a specific deployed contract.
func NewIArbitrumVM(address common.Address, backend bind.ContractBackend) (*IArbitrumVM, error) {
	contract, err := bindIArbitrumVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArbitrumVM{IArbitrumVMCaller: IArbitrumVMCaller{contract: contract}, IArbitrumVMTransactor: IArbitrumVMTransactor{contract: contract}, IArbitrumVMFilterer: IArbitrumVMFilterer{contract: contract}}, nil
}

// NewIArbitrumVMCaller creates a new read-only instance of IArbitrumVM, bound to a specific deployed contract.
func NewIArbitrumVMCaller(address common.Address, caller bind.ContractCaller) (*IArbitrumVMCaller, error) {
	contract, err := bindIArbitrumVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArbitrumVMCaller{contract: contract}, nil
}

// NewIArbitrumVMTransactor creates a new write-only instance of IArbitrumVM, bound to a specific deployed contract.
func NewIArbitrumVMTransactor(address common.Address, transactor bind.ContractTransactor) (*IArbitrumVMTransactor, error) {
	contract, err := bindIArbitrumVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArbitrumVMTransactor{contract: contract}, nil
}

// NewIArbitrumVMFilterer creates a new log filterer instance of IArbitrumVM, bound to a specific deployed contract.
func NewIArbitrumVMFilterer(address common.Address, filterer bind.ContractFilterer) (*IArbitrumVMFilterer, error) {
	contract, err := bindIArbitrumVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArbitrumVMFilterer{contract: contract}, nil
}

// bindIArbitrumVM binds a generic wrapper to an already deployed contract.
func bindIArbitrumVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArbitrumVMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbitrumVM *IArbitrumVMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbitrumVM.Contract.IArbitrumVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbitrumVM *IArbitrumVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbitrumVM.Contract.IArbitrumVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbitrumVM *IArbitrumVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbitrumVM.Contract.IArbitrumVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArbitrumVM *IArbitrumVMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArbitrumVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArbitrumVM *IArbitrumVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArbitrumVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArbitrumVM *IArbitrumVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArbitrumVM.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbitrumVM *IArbitrumVMTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbitrumVM.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbitrumVM *IArbitrumVMSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbitrumVM.Contract.CompleteChallenge(&_IArbitrumVM.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IArbitrumVM *IArbitrumVMTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IArbitrumVM.Contract.CompleteChallenge(&_IArbitrumVM.TransactOpts, _players, _rewards)
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
var MerkleLibBin = "0x610575610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043f565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b6000815b600181511115610422576060600282516001018161031f57fe5b04604051908082528060200260200182016040528015610349578160200160208202803883390190505b50905060005b815181101561041a5782518160020260010110156103e25782816002028151811061037657fe5b602002602001015183826002026001018151811061039057fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d157fe5b602002602001018181525050610412565b8281600202815181106103f157fe5b602002602001015182828151811061040557fe5b6020026020010181815250505b60010161034f565b509050610305565b8060008151811061042f57fe5b6020026020010151915050919050565b600080838160205b88518111610532578089015193506020818a51036020018161046557fe5b0491505b60008211801561047c5750600286066001145b801561048a57508160020a86115b1561049d57600286046001019550610469565b600286066104e85783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104e057fe5b04955061052a565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052357fe5b0460010195505b602001610447565b50509094149594505050505056fea265627a7a72315820dab403e0cf8b5b3cea673ec093051f92ba4d6ae087244559622ae30d50f21ec664736f6c634300050d0032"

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
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"beforeInbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastMessage\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLog\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"d82f8fd5": "validateProof(bytes32,uint64[2],bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x6134ce610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063d82f8fd51461003a575b600080fd5b610138600480360361014081101561005157600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594602085013594604081013594506060810135935060808101359260a082013592909160e081019060c001356401000000008111156100c357600080fd5b8201836020820111156100d557600080fd5b803590602001918460018302840111640100000000831117156100f757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061014a945050505050565b60408051918252519081900360200190f35b60006101946040518061012001604052808c81526020018b81526020018a8152602001898152602001888152602001878152602001868152602001858152602001848152506101a2565b9a9950505050505050505050565b600080808060606101b16132d2565b6101b96132d2565b6101c2886110a6565b939950929650909450925090506001600060ff88168214156102185761021183866000815181106101ef57fe5b60200260200101518760018151811061020457fe5b60200260200101516114e6565b9150610efa565b60ff88166002141561025757610211838660008151811061023557fe5b60200260200101518760018151811061024a57fe5b6020026020010151611536565b60ff88166003141561029657610211838660008151811061027457fe5b60200260200101518760018151811061028957fe5b6020026020010151611577565b60ff8816600414156102d55761021183866000815181106102b357fe5b6020026020010151876001815181106102c857fe5b60200260200101516115b8565b60ff8816600514156103145761021183866000815181106102f257fe5b60200260200101518760018151811061030757fe5b6020026020010151611609565b60ff88166006141561035357610211838660008151811061033157fe5b60200260200101518760018151811061034657fe5b602002602001015161165a565b60ff88166007141561039257610211838660008151811061037057fe5b60200260200101518760018151811061038557fe5b60200260200101516116ab565b60ff8816600814156103e65761021183866000815181106103af57fe5b6020026020010151876001815181106103c457fe5b6020026020010151886002815181106103d957fe5b60200260200101516116fc565b60ff88166009141561043a57610211838660008151811061040357fe5b60200260200101518760018151811061041857fe5b60200260200101518860028151811061042d57fe5b6020026020010151611766565b60ff8816600a141561047957610211838660008151811061045757fe5b60200260200101518760018151811061046c57fe5b60200260200101516117bf565b60ff8816601014156104b857610211838660008151811061049657fe5b6020026020010151876001815181106104ab57fe5b6020026020010151611800565b60ff8816601114156104f75761021183866000815181106104d557fe5b6020026020010151876001815181106104ea57fe5b6020026020010151611841565b60ff88166012141561053657610211838660008151811061051457fe5b60200260200101518760018151811061052957fe5b6020026020010151611882565b60ff88166013141561057557610211838660008151811061055357fe5b60200260200101518760018151811061056857fe5b60200260200101516118c3565b60ff8816601414156105b457610211838660008151811061059257fe5b6020026020010151876001815181106105a757fe5b6020026020010151611904565b60ff8816601514156105de5761021183866000815181106105d157fe5b6020026020010151611930565b60ff88166016141561061d5761021183866000815181106105fb57fe5b60200260200101518760018151811061061057fe5b6020026020010151611976565b60ff88166017141561065c57610211838660008151811061063a57fe5b60200260200101518760018151811061064f57fe5b60200260200101516119b7565b60ff88166018141561069b57610211838660008151811061067957fe5b60200260200101518760018151811061068e57fe5b60200260200101516119f8565b60ff8816601914156106c55761021183866000815181106106b857fe5b6020026020010151611a39565b60ff8816601a14156107045761021183866000815181106106e257fe5b6020026020010151876001815181106106f757fe5b6020026020010151611a6f565b60ff8816601b141561074357610211838660008151811061072157fe5b60200260200101518760018151811061073657fe5b6020026020010151611ab0565b60ff88166020141561076d57610211838660008151811061076057fe5b6020026020010151611af1565b60ff88166021141561079757610211838660008151811061078a57fe5b6020026020010151611b0d565b60ff8816603014156107c15761021183866000815181106107b457fe5b6020026020010151611b28565b60ff8816603114156107d65761021183611b30565b60ff8816603214156107eb5761021183611b51565b60ff88166033141561081557610211838660008151811061080857fe5b6020026020010151611b6a565b60ff88166034141561083f57610211838660008151811061083257fe5b6020026020010151611b83565b60ff88166035141561087e57610211838660008151811061085c57fe5b60200260200101518760018151811061087157fe5b6020026020010151611b99565b60ff8816603614156108935761021183611be1565b60ff8816603714156108ad57610211838560000151611c13565b60ff8816603814156108d75761021183866000815181106108ca57fe5b6020026020010151611c25565b60ff881660391415610964576108eb613333565b6108fa8b610100015188611c37565b91995097509050871561093e5760405162461bcd60e51b81526004018080602001828103825260218152602001806134796021913960400191505060405180910390fd5b61094e858263ffffffff611dc116565b61095e848263ffffffff611de316565b50610efa565b60ff8816603a14156109795761021183611e00565b60ff8816603b141561098a57610efa565b60ff8816603c141561099f5761021183611e20565b60ff8816603d14156109c95761021183866000815181106109bc57fe5b6020026020010151611e39565b60ff8816604014156109f35761021183866000815181106109e657fe5b6020026020010151611e67565b60ff881660411415610a32576102118386600081518110610a1057fe5b602002602001015187600181518110610a2557fe5b6020026020010151611e89565b60ff881660421415610a86576102118386600081518110610a4f57fe5b602002602001015187600181518110610a6457fe5b602002602001015188600281518110610a7957fe5b6020026020010151611ebb565b60ff881660431415610ac5576102118386600081518110610aa357fe5b602002602001015187600181518110610ab857fe5b6020026020010151611efd565b60ff881660441415610b19576102118386600081518110610ae257fe5b602002602001015187600181518110610af757fe5b602002602001015188600281518110610b0c57fe5b6020026020010151611f0f565b60ff881660501415610b58576102118386600081518110610b3657fe5b602002602001015187600181518110610b4b57fe5b6020026020010151611f31565b60ff881660511415610bac576102118386600081518110610b7557fe5b602002602001015187600181518110610b8a57fe5b602002602001015188600281518110610b9f57fe5b6020026020010151611fa7565b60ff881660521415610bd6576102118386600081518110610bc957fe5b602002602001015161201f565b60ff881660601415610beb5761021183612052565b60ff881660611415610ce857610c158386600081518110610c0857fe5b6020026020010151612058565b90925090508115610cdf578960e001518a60c001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610c945760405162461bcd60e51b815260040180806020018281038252602581526020018061342d6025913960400191505060405180910390fd5b8960a001518a6080015114610cda5760405162461bcd60e51b81526004018080602001828103825260278152602001806134526027913960400191505060405180910390fd5b610ce3565b5060005b610efa565b60ff881660701415610dd757610d128386600081518110610d0557fe5b602002602001015161207c565b90925090508115610cdf578960a001518a608001518260405160200180838152602001828152602001925050506040516020818303038152906040528051906020012014610d915760405162461bcd60e51b81526004018080602001828103825260298152602001806133bd6029913960400191505060405180910390fd5b8960e001518a60c0015114610cda5760405162461bcd60e51b81526004018080602001828103825260268152602001806133e66026913960400191505060405180910390fd5b60ff881660711415610e93576040805160028082526060828101909352816020015b610e01613333565b815260200190600190039081610df957505060208c0151909150610e369060005b602002015167ffffffffffffffff16612096565b81600081518110610e4357fe5b6020026020010181905250610e628b60200151600160028110610e2257fe5b81600181518110610e6f57fe5b602002602001018190525061095e610e8682612114565b859063ffffffff611de316565b60ff881660721415610ed0576102118386600081518110610eb057fe5b602002602001015160405180602001604052808e604001518152506121c4565b60ff881660731415610ee55760009150610efa565b60ff881660741415610efa57610efa83612236565b80610f8b578960a001518a6080015114610f455760405162461bcd60e51b81526004018080602001828103825260278152602001806134526027913960400191505060405180910390fd5b8960e001518a60c0015114610f8b5760405162461bcd60e51b81526004018080602001828103825260268152602001806133e66026913960400191505060405180910390fd5b81610fed5760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a0840151511415610fe557610fe083612240565b610fed565b60a083015183525b610ff68461224a565b8a51146110345760405162461bcd60e51b815260040180806020018281038252602281526020018061339b6022913960400191505060405180910390fd5b61103d8361224a565b8a6060015114611094576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b600060606110b26132d2565b6110ba6132d2565b600080806110c66132d2565b6110cf816122df565b6110de896101000151846122e9565b90945090925090506110ee6132d2565b6110f7826123ee565b905060008a6101000151858151811061110c57fe5b602001015160f81c60f81b60f81c905060008b6101000151866001018151811061113257fe5b016020015160f81c905060006111478261244c565b905060608160405190808252806020026020018201604052801561118557816020015b611172613333565b81526020019060019003908161116a5790505b5090506002880197508360ff16600014806111a357508360ff166001145b6111f4576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff8416611297576040805160208082018084528951516353409fab60e01b90915260ff87166024840152604483015291519091829173__$d969135829891f807aa9c34494da4ecd99$__916353409fab916064808601929190818703018186803b15801561126257600080fd5b505af4158015611276573d6000803e3d6000fd5b505050506040513d602081101561128c57600080fd5b5051905286526113ee565b61129f613333565b6112ae8f61010001518a611c37565b909a5090985090508715611309576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561132d57808260008151811061131d57fe5b602002602001018190525061133d565b61133d868263ffffffff611de316565b604051806020016040528073__$d969135829891f807aa9c34494da4ecd99$__63264f384b8761136c86612466565b518c5151604080516001600160e01b031960e087901b16815260ff909416600485015260248401929092526044830152516064808301926020929190829003018186803b1580156113bc57600080fd5b505af41580156113d0573d6000803e3d6000fd5b505050506040513d60208110156113e657600080fd5b505190528752505b60ff84165b828110156114825761140a8f61010001518a611c37565b845185908590811061141857fe5b602090810291909101015299509750871561147a576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b6001016113f3565b8151156114cf575060005b8460ff168251038110156114cf576114c78282600185510303815181106114b057fe5b602002602001015188611de390919063ffffffff16565b60010161148d565b50919d919c50939a50919850939650945050505050565b60006114f18361259c565b158061150357506115018261259c565b155b156115105750600061152f565b82518251808201611527878263ffffffff6125a716565b600193505050505b9392505050565b60006115418361259c565b158061155357506115518261259c565b155b156115605750600061152f565b82518251808202611527878263ffffffff6125a716565b60006115828361259c565b158061159457506115928261259c565b155b156115a15750600061152f565b82518251808203611527878263ffffffff6125a716565b60006115c38361259c565b15806115d557506115d38261259c565b155b156115e25750600061152f565b82518251806115f65760009250505061152f565b808204611527878263ffffffff6125a716565b60006116148361259c565b158061162657506116248261259c565b155b156116335750600061152f565b82518251806116475760009250505061152f565b808205611527878263ffffffff6125a716565b60006116658361259c565b158061167757506116758261259c565b155b156116845750600061152f565b82518251806116985760009250505061152f565b808206611527878263ffffffff6125a716565b60006116b68361259c565b15806116c857506116c68261259c565b155b156116d55750600061152f565b82518251806116e95760009250505061152f565b808207611527878263ffffffff6125a716565b60006117078461259c565b158061171957506117178361259c565b155b156117265750600061175e565b8351835183518061173d576000935050505061175e565b6000818385089050611755898263ffffffff6125a716565b60019450505050505b949350505050565b60006117718461259c565b158061178357506117818361259c565b155b156117905750600061175e565b835183518351806117a7576000935050505061175e565b6000818385099050611755898263ffffffff6125a716565b60006117ca8361259c565b15806117dc57506117da8261259c565b155b156117e95750600061152f565b8251825180820a611527878263ffffffff6125a716565b600061180b8361259c565b158061181d575061181b8261259c565b155b1561182a5750600061152f565b82518251808210611527878263ffffffff6125a716565b600061184c8361259c565b158061185e575061185c8261259c565b155b1561186b5750600061152f565b82518251808211611527878263ffffffff6125a716565b600061188d8361259c565b158061189f575061189d8261259c565b155b156118ac5750600061152f565b82518251808212611527878263ffffffff6125a716565b60006118ce8361259c565b15806118e057506118de8261259c565b155b156118ed5750600061152f565b82518251808213611527878263ffffffff6125a716565b6000611926610e8661191584612466565b5161191f86612466565b51146125bb565b5060019392505050565b600061193b8261259c565b6119555761195083600063ffffffff6125a716565b61196c565b81518015611969858263ffffffff6125a716565b50505b5060015b92915050565b60006119818361259c565b158061199357506119918261259c565b155b156119a05750600061152f565b82518251808216611527878263ffffffff6125a716565b60006119c28361259c565b15806119d457506119d28261259c565b155b156119e15750600061152f565b82518251808217611527878263ffffffff6125a716565b6000611a038361259c565b1580611a155750611a138261259c565b155b15611a225750600061152f565b82518251808218611527878263ffffffff6125a716565b6000611a448261259c565b611a5057506000611970565b81518019611a64858263ffffffff6125a716565b506001949350505050565b6000611a7a8361259c565b1580611a8c5750611a8a8261259c565b155b15611a995750600061152f565b8251825181811a611527878263ffffffff6125a716565b6000611abb8361259c565b1580611acd5750611acb8261259c565b155b15611ada5750600061152f565b8251825181810b611527878263ffffffff6125a716565b600061196c611aff83612466565b51849063ffffffff6125a716565b600061196c611b1b836125e4565b849063ffffffff611de316565b600192915050565b6000611b4982608001518361266d90919063ffffffff16565b506001919050565b6000611b4982606001518361266d90919063ffffffff16565b6000611b7582612466565b606084015250600192915050565b6000611b8e82612466565b835250600192915050565b6000611ba48361267b565b611bb05750600061152f565b611bb98261259c565b611bc55750600061152f565b81511561192657611bd583612466565b84525060019392505050565b6000611b49611c06611bf9611bf4612688565b612466565b51602085015151146125bb565b839063ffffffff611de316565b600061196c838363ffffffff61266d16565b600061196c838363ffffffff611dc116565b600080611c42613333565b84518410611c97576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110611caa57fe5b016020015160019092019160f81c90506000611cc4613361565b60ff8316611cf857611cd68985612705565b9094509150600084611ce784612096565b91985096509450611dba9350505050565b60ff831660011415611d1f57611d0e898561272c565b9094509050600084611ce783612899565b60ff831660021415611d4657611d358985612705565b9094509150600084611ce7846128f9565b600360ff841610801590611d5d5750600c60ff8416105b15611d9a57600219830160606000611d76838d89612977565b909850925090508087611d8884612114565b99509950995050505050505050611dba565b8260ff16612710016000611dae6000612096565b91985096509450505050505b9250925092565b611dd78260400151611dd283612466565b612a32565b82604001819052505050565b611df48260200151611dd283612466565b82602001819052505050565b6000611b49611c06611e13611bf4612688565b51604085015151146125bb565b6000611b498260a001518361266d90919063ffffffff16565b6000611e448261267b565b611e5057506000611970565b611e5982612466565b60a084015250600192915050565b6000611e79838363ffffffff611de316565b61196c838363ffffffff611de316565b6000611e9b848363ffffffff611de316565b611eab848463ffffffff611de316565b611926848363ffffffff611de316565b6000611ecd858363ffffffff611de316565b611edd858463ffffffff611de316565b611eed858563ffffffff611de316565b611a64858363ffffffff611de316565b6000611eab848463ffffffff611de316565b6000611f21858563ffffffff611de316565b611eed858463ffffffff611de316565b6000611f3c8361259c565b1580611f4e5750611f4c82612ae8565b155b15611f5b5750600061152f565b611f6482612af7565b60ff16836000015110611f795750600061152f565b6119268260400151846000015181518110611f9057fe5b602002602001015185611de390919063ffffffff16565b6000611fb283612ae8565b1580611fc45750611fc28461259c565b155b15611fd15750600061175e565b611fda83612af7565b60ff16846000015110611fef5750600061175e565b81836040015185600001518151811061200457fe5b6020908102919091010152611a64858463ffffffff611de316565b600061202a82612ae8565b61203657506000611970565b61196c61204283612af7565b849060ff1663ffffffff6125a716565b50600190565b600080612063613388565b61206c84612466565b51600193509150505b9250929050565b600080600161208a84612466565b51909590945092505050565b61209e613333565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612103565b6120f0613333565b8152602001906001900390816120e85790505b508152600060209091015292915050565b61211c613333565b6121268251612b06565b612177576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b80516000906121d284612466565b511415612226576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b611926848363ffffffff61266d16565b600260c090910152565b600160c090910152565b600060028260c001511415612261575060006110a1565b60018260c001511415612276575060016110a1565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101206110a1565b600060c090910152565b6000806122f46132d2565b6122fc6132d2565b600060c0820181905261230f8787612b0d565b84529650905080156123275793508492509050611dba565b6123318787612b0d565b602085015296509050801561234c5793508492509050611dba565b6123568787612b0d565b60408501529650905080156123715793508492509050611dba565b61237b8787612b0d565b60608501529650905080156123965793508492509050611dba565b6123a08787612b0d565b60808501529650905080156123bb5793508492509050611dba565b6123c58787612b0d565b60a08501529650905080156123e05793508492509050611dba565b506000969495509392505050565b6123f66132d2565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b600080600061245d8460ff16612b4b565b50949350505050565b61246e613388565b6060820151600c60ff909116106124c0576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166124ed5760405180602001604052806124e48460000151612fde565b905290506110a1565b606082015160ff16600114156125345760405180602001604052806124e4846020015160000151856020015160400151866020015160600151876020015160200151613002565b606082015160ff166002141561255957506040805160208101909152815181526110a1565b600360ff16826060015160ff161015801561257d57506060820151600c60ff909116105b1561259a5760405180602001604052806124e484604001516130aa565bfe5b6060015160ff161590565b611df48260200151611dd2611bf484612096565b6125c3613333565b81156125da576125d36001612096565b90506110a1565b6125d36000612096565b6125ec613333565b816060015160ff16600214156126335760405162461bcd60e51b815260040180806020018281038252602181526020018061340c6021913960400191505060405180910390fd5b606082015160ff16612649576125d36000612096565b816060015160ff1660011415612663576125d36001612096565b6125d36003612096565b611df4826020015182612a32565b6060015160ff1660011490565b612690613333565b6040805160808082018352600080835283519182018452808252602082810182905282850182905260608301829052808401929092528351818152918201845291928301916126f5565b6126e2613333565b8152602001906001900390816126da5790505b5081526003602090910152905090565b600080828161271a868363ffffffff6131f616565b60209290920196919550909350505050565b6000612736613361565b6000839050600085828151811061274957fe5b602001015160f81c60f81b60f81c90508180600101925050600086838151811061276f57fe5b016020015160019384019360f89190911c915060009060ff8416141561280d576000612799613333565b6127a38a87611c37565b909750909250905081156127fe576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b61280781612466565b51925050505b600061281f898663ffffffff6131f616565b90506020850194508360ff1660011415612864576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506120759050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b6128a1613333565b6040805160808101825260008082526020808301869052835182815290810184529192830191906128e8565b6128d5613333565b8152602001906001900390816128cd5790505b508152600160209091015292915050565b612901613333565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612966565b612953613333565b81526020019060019003908161294b5790505b508152600260209091015292915050565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156129c257816020015b6129af613333565b8152602001906001900390816129a75790505b50905060005b8960ff168160ff161015612a1c576129e08985611c37565b8451859060ff86169081106129f157fe5b6020908102919091010152945092508215612a1457509094509092509050612a29565b6001016129c8565b5060009550919350909150505b93509350939050565b612a3a613388565b6040805160028082526060828101909352816020015b612a58613388565b815260200190600190039081612a505790505090508281600081518110612a7b57fe5b60200260200101819052508381600181518110612a9457fe5b60200260200101819052506040518060200160405280612ade6040518060400160405280612ac588600001516128f9565b8152602001612ad789600001516128f9565b9052613212565b9052949350505050565b60006119708260600151613291565b600061197082606001516132af565b6008101590565b600080612b18613388565b836000612b2b878363ffffffff6131f616565b604080516020808201909252918252600099930197509550909350505050565b6000806001831415612b635750600290506001612fd9565b6002831415612b785750600290506001612fd9565b6003831415612b8d5750600290506001612fd9565b6004831415612ba25750600290506001612fd9565b6005831415612bb75750600290506001612fd9565b6006831415612bcc5750600290506001612fd9565b6007831415612be15750600290506001612fd9565b6008831415612bf65750600390506001612fd9565b6009831415612c0b5750600390506001612fd9565b600a831415612c205750600290506001612fd9565b6010831415612c355750600290506001612fd9565b6011831415612c4a5750600290506001612fd9565b6012831415612c5f5750600290506001612fd9565b6013831415612c745750600290506001612fd9565b6014831415612c895750600290506001612fd9565b6015831415612c9d57506001905080612fd9565b6016831415612cb25750600290506001612fd9565b6017831415612cc75750600290506001612fd9565b6018831415612cdc5750600290506001612fd9565b6019831415612cf057506001905080612fd9565b601a831415612d055750600290506001612fd9565b601b831415612d1a5750600290506001612fd9565b6020831415612d2e57506001905080612fd9565b6021831415612d4257506001905080612fd9565b6030831415612d575750600190506000612fd9565b6031831415612d6c5750600090506001612fd9565b6032831415612d815750600090506001612fd9565b6033831415612d965750600190506000612fd9565b6034831415612dab5750600190506000612fd9565b6035831415612dc05750600290506000612fd9565b6036831415612dd55750600090506001612fd9565b6037831415612dea5750600090506001612fd9565b6038831415612dff5750600190506000612fd9565b6039831415612e145750600090506001612fd9565b603a831415612e295750600090506001612fd9565b603b831415612e3d57506000905080612fd9565b603c831415612e525750600090506001612fd9565b603d831415612e675750600190506000612fd9565b6040831415612e7c5750600190506002612fd9565b6041831415612e915750600290506003612fd9565b6042831415612ea65750600390506004612fd9565b6043831415612eba57506002905080612fd9565b6044831415612ece57506003905080612fd9565b6050831415612ee35750600290506001612fd9565b6051831415612ef85750600390506001612fd9565b6052831415612f0c57506001905080612fd9565b6060831415612f2057506000905080612fd9565b6061831415612f355750600190506000612fd9565b6070831415612f4a5750600190506000612fd9565b6071831415612f5f5750600090506001612fd9565b6072831415612f7357506001905080612fd9565b6073831415612f8757506000905080612fd9565b6074831415612f9b57506000905080612fd9565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561305c575060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602282018590526042808301859052835180840390910181526062909201909252805191012061175e565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b60006008825111156130fa576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613127578160200160208202803883390190505b50805190915060005b8181101561318357613140613388565b61315c86838151811061314f57fe5b6020026020010151612466565b9050806000015184838151811061316f57fe5b602090810291909101015250600101613130565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156131cc5781810151838201526020016131b4565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6000816020018351101561320957600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b613235613333565b81526020019060019003908161322d575050805190915060005b818110156132875784816002811061326357fe5b602002015183828151811061327457fe5b602090810291909101015260010161324f565b5061175e826130aa565b6000600c60ff8316108015611970575050600360ff91909116101590565b60006132ba82613291565b156132ca575060021981016110a1565b5060016110a1565b6040518060e001604052806132e5613388565b81526020016132f2613388565b81526020016132ff613388565b815260200161330c613388565b8152602001613319613388565b8152602001613326613388565b8152602001600081525090565b60405180608001604052806000815260200161334d613361565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe50726f6f6620686164206e6f6e206d61746368696e6720737461727420737461746573656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d6573736167654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f64654c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a723158201fe579d6582df7186d666b184b3cea7c218493acdff260ab43175545c0e7a37d64736f6c634300050d0032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

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

// ValidateProof is a free data retrieval call binding the contract method 0xd82f8fd5.
//
// Solidity: function validateProof(bytes32 beforeHash, uint64[2] timeBounds, bytes32 beforeInbox, bytes32 afterHash, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, beforeHash [32]byte, timeBounds [2]uint64, beforeInbox [32]byte, afterHash [32]byte, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", beforeHash, timeBounds, beforeInbox, afterHash, firstMessage, lastMessage, firstLog, lastLog, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0xd82f8fd5.
//
// Solidity: function validateProof(bytes32 beforeHash, uint64[2] timeBounds, bytes32 beforeInbox, bytes32 afterHash, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(beforeHash [32]byte, timeBounds [2]uint64, beforeInbox [32]byte, afterHash [32]byte, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBounds, beforeInbox, afterHash, firstMessage, lastMessage, firstLog, lastLog, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0xd82f8fd5.
//
// Solidity: function validateProof(bytes32 beforeHash, uint64[2] timeBounds, bytes32 beforeInbox, bytes32 afterHash, bytes32 firstMessage, bytes32 lastMessage, bytes32 firstLog, bytes32 lastLog, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(beforeHash [32]byte, timeBounds [2]uint64, beforeInbox [32]byte, afterHash [32]byte, firstMessage [32]byte, lastMessage [32]byte, firstLog [32]byte, lastLog [32]byte, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, beforeHash, timeBounds, beforeInbox, afterHash, firstMessage, lastMessage, firstLog, lastLog, proof)
}
