// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengeRPC

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

// ArbMachineABI is the input ABI used to generate the binding from.
const ArbMachineABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"instructionStackHash\",\"type\":\"bytes32\"},{\"name\":\"dataStackHash\",\"type\":\"bytes32\"},{\"name\":\"auxStackHash\",\"type\":\"bytes32\"},{\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"name\":\"errHandlerHash\",\"type\":\"bytes32\"}],\"name\":\"machineHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbMachineBin is the compiled bytecode used for deploying new contracts.
const ArbMachineBin = `0x6101d6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063c1355b591461003a575b600080fd5b610075600480360360c081101561005057600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610087565b60408051918252519081900360200190f35b604080516101008101825260e081018881528152815160208181018452888252808301919091528251808201845287815282840152825180820184528681526060830152825180820184528581526080830152825190810190925282825260a0810191909152600060c08201819052906101009061010b565b979650505050505050565b600060028260c0015114156101225750600061019c565b60018260c0015114156101375750600161019c565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101205b91905056fea265627a7a7230582037b31eec1eb75344fd449ad6a22d572f0e00f2dee0f6e6f08fa95321f471304064736f6c63430005090032`

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
	ret0 := new([32]byte)

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
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[5]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"unanimousAssertHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"name\":\"_lastLogHash\",\"type\":\"bytes32\"},{\"name\":\"_totalMessageValueAmounts\",\"type\":\"uint256[]\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_dest\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"bytes32\"}],\"name\":\"generateSentMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"countSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_signatures\",\"type\":\"bytes\"},{\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"parseSignature\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageDataHashes\",\"type\":\"bytes32[]\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageValueAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"bytes32[]\"}],\"name\":\"generateLastMessageHashStub\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"bytes32\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"},{\"name\":\"_newMessage\",\"type\":\"bytes32\"}],\"name\":\"appendInboxPendingMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"name\":\"_challengeManagerNum\",\"type\":\"uint16\"},{\"name\":\"_assertKeys\",\"type\":\"address[]\"}],\"name\":\"createVMHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"recoverAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_inboxHash\",\"type\":\"bytes32\"},{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"}],\"name\":\"appendInboxMessages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
const ArbProtocolBin = `0x612210610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100f45760003560e01c8063b31d63cc11610096578063d78d18ea11610070578063d78d18ea14610f6b578063e440673b14610f8e578063f0c8e96914611064578063f11fcc261461110f576100f4565b8063b31d63cc14610bc1578063b327749514610c89578063ccf69dd714610f32576100f4565b806325200160116100d257806325200160146106d45780632a0500d81461098057806333ae3ad0146109bf5780633e28559814610a63576100f4565b8063014bba5b146100f95780630f89fbff146104145780632090372114610609575b600080fd5b610402600480360361018081101561011057600080fd5b810190808060a001906005806020026040519081016040528092919082600560200280828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561018657600080fd5b82018360208201111561019857600080fd5b803590602001918460208302840111600160201b831117156101b957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561020857600080fd5b82018360208201111561021a57600080fd5b803590602001918460018302840111600160201b8311171561023b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561028d57600080fd5b82018360208201111561029f57600080fd5b803590602001918460208302840111600160201b831117156102c057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561030f57600080fd5b82018360208201111561032157600080fd5b803590602001918460208302840111600160201b8311171561034257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561039157600080fd5b8201836020820111156103a357600080fd5b803590602001918460208302840111600160201b831117156103c457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611132945050505050565b60408051918252519081900360200190f35b6105b96004803603606081101561042a57600080fd5b810190602081018135600160201b81111561044457600080fd5b82018360208201111561045657600080fd5b803590602001918460208302840111600160201b8311171561047757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104c657600080fd5b8201836020820111156104d857600080fd5b803590602001918460208302840111600160201b831117156104f957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561054857600080fd5b82018360208201111561055a57600080fd5b803590602001918460208302840111600160201b8311171561057b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506112cb945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156105f55781810151838201526020016105dd565b505050509050019250505060405180910390f35b610402600480360360e081101561061f57600080fd5b81359163ffffffff6020820135169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561066357600080fd5b82018360208201111561067557600080fd5b803590602001918460208302840111600160201b8311171561069657600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061144c945050505050565b610402600480360360a08110156106ea57600080fd5b810190602081018135600160201b81111561070457600080fd5b82018360208201111561071657600080fd5b803590602001918460208302840111600160201b8311171561073757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561078657600080fd5b82018360208201111561079857600080fd5b803590602001918460018302840111600160201b831117156107b957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561080b57600080fd5b82018360208201111561081d57600080fd5b803590602001918460208302840111600160201b8311171561083e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561088d57600080fd5b82018360208201111561089f57600080fd5b803590602001918460208302840111600160201b831117156108c057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561090f57600080fd5b82018360208201111561092157600080fd5b803590602001918460208302840111600160201b8311171561094257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506114b7945050505050565b610402600480360360a081101561099657600080fd5b508035906020810135906001600160581b03196040820135169060608101359060800135611679565b610402600480360360208110156109d557600080fd5b810190602081018135600160201b8111156109ef57600080fd5b820183602082011115610a0157600080fd5b803590602001918460018302840111600160201b83111715610a2257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611857945050505050565b610402600480360360c0811015610a7957600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594909390925060408101915060200135600160201b811115610ace57600080fd5b820183602082011115610ae057600080fd5b803590602001918460208302840111600160201b83111715610b0157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b5057600080fd5b820183602082011115610b6257600080fd5b803590602001918460208302840111600160201b83111715610b8357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611886945050505050565b610c6760048036036040811015610bd757600080fd5b810190602081018135600160201b811115610bf157600080fd5b820183602082011115610c0357600080fd5b803590602001918460018302840111600160201b83111715610c2457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250611971915050565b6040805160ff9094168452602084019290925282820152519081900360600190f35b610402600480360360a0811015610c9f57600080fd5b810190602081018135600160201b811115610cb957600080fd5b820183602082011115610ccb57600080fd5b803590602001918460208302840111600160201b83111715610cec57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d3b57600080fd5b820183602082011115610d4d57600080fd5b803590602001918460208302840111600160201b83111715610d6e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610dbd57600080fd5b820183602082011115610dcf57600080fd5b803590602001918460208302840111600160201b83111715610df057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e3f57600080fd5b820183602082011115610e5157600080fd5b803590602001918460208302840111600160201b83111715610e7257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610ec157600080fd5b820183602082011115610ed357600080fd5b803590602001918460208302840111600160201b83111715610ef457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506119c3945050505050565b61040260048036036080811015610f4857600080fd5b508035906001600160581b03196020820135169060408101359060600135611aa3565b61040260048036036040811015610f8157600080fd5b5080359060200135611b81565b610402600480360360c0811015610fa457600080fd5b63ffffffff82358116926001600160801b036020820135169260408201359092169160608201359161ffff6080820135169181019060c0810160a0820135600160201b811115610ff357600080fd5b82018360208201111561100557600080fd5b803590602001918460208302840111600160201b8311171561102657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611bc5945050505050565b6105b96004803603604081101561107a57600080fd5b81359190810190604081016020820135600160201b81111561109b57600080fd5b8201836020820111156110ad57600080fd5b803590602001918460018302840111600160201b831117156110ce57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c8a945050505050565b6104026004803603604081101561112557600080fd5b5080359060200135611e30565b6000878787878787876040516020018088600560200280838360005b8381101561116657818101518382015260200161114e565b5050505090500187600260200280838360005b83811015611191578181015183820152602001611179565b50505050905001868051906020019060200280838360005b838110156111c15781810151838201526020016111a9565b5050505090500185805190602001908083835b602083106111f35780518252601f1990920191602091820191016111d4565b51815160209384036101000a60001901801990921691161790528751919093019287810192500280838360005b83811015611238578181015183820152602001611220565b50505050905001838051906020019060200280838360005b83811015611268578181015183820152602001611250565b50505050905001828051906020019060200280838360005b83811015611298578181015183820152602001611280565b50505050905001975050505050505050604051602081830303815290604052805190602001209050979650505050505050565b60608084516040519080825280602002602001820160405280156112f9578160200160208202803883390190505b50905060005b8451811015611443578585828151811061131557fe5b602002602001015161ffff168151811061132b57fe5b602002602001015160146015811061133f57fe5b1a60f81b6001600160f81b03191661139f5783818151811061135d57fe5b60200260200101518286838151811061137257fe5b602002602001015161ffff168151811061138857fe5b60200260200101818151019150818152505061143b565b818582815181106113ac57fe5b602002602001015161ffff16815181106113c257fe5b60200260200101516000146113d657600080fd5b8381815181106113e257fe5b6020026020010151600014156113f757600080fd5b83818151811061140357fe5b60200260200101518286838151811061141857fe5b602002602001015161ffff168151811061142e57fe5b6020026020010181815250505b6001016112ff565b50949350505050565b600087878787878787604051602001808881526020018763ffffffff1663ffffffff1660e01b81526004018681526020018581526020018481526020018381526020018280519060200190602002808383600083811015611298578181015183820152602001611280565b600081518351146114c757600080fd5b83518351146114d557600080fd5b60008080805b865181101561166c5773__$0d86abb4a722a612872fb80f4c7e7e95bd$__63615c39b08a856040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561155257818101518382015260200161153a565b50505050905090810190601f16801561157f5780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561159c57600080fd5b505af41580156115b0573d6000803e3d6000fd5b505050506040513d60408110156115c657600080fd5b508051602090910151895191945092506116359083908c908b90859081106115ea57fe5b602002602001015161ffff168151811061160057fe5b602002602001015189848151811061161457fe5b602002602001015189858151811061162857fe5b6020026020010151611aa3565b60408051602080820197909752808201839052815180820383018152606090910190915280519501949094209391506001016114db565b5050505095945050505050565b60408051602080820188905281830187905260608083018690526001600160581b03198716608084015283518084036075018152609584018086528151919093012060048084526101358501909552600094909391929160b5015b6116dc6121a5565b8152602001906001900390816116d45790505090506116fa87611e4a565b8160008151811061170757fe5b602002602001018190525061171b42611ea4565b8160018151811061172857fe5b602002602001018190525061173c43611ea4565b8160028151811061174957fe5b602090810291909101015261175d82611ea4565b8160038151811061176a57fe5b602090810291909101015260408051600480825260a08201909252606091816020015b6117956121a5565b81526020019060019003908161178d5790505090506117b382611efe565b816000815181106117c057fe5b60209081029190910101526117d485611ea4565b816001815181106117e157fe5b60200260200101819052506117f586611ea4565b8160028151811061180257fe5b60209081029190910101526118206001600160581b03198816611ea4565b8160038151811061182d57fe5b602002602001018190525061184961184482611efe565b611f3e565b519998505050505050505050565b6000604182518161186457fe5b061561187157600061187e565b604182518161187c57fe5b045b90505b919050565b600085858260200201518660016020020151868686604051602001808781526020018667ffffffffffffffff1667ffffffffffffffff1660c01b81526008018567ffffffffffffffff1667ffffffffffffffff1660c01b8152600801848152602001838051906020019060200280838360005b838110156119115781810151838201526020016118f9565b50505050905001828051906020019060200280838360005b83811015611941578181015183820152602001611929565b50505050905001965050505050505060405160208183030381529060405280519060200120905095945050505050565b604180820283810160208101516040820151919093015160ff169291601b84101561199d57601b840193505b8360ff16601b14806119b257508360ff16601c145b6119bb57600080fd5b509250925092565b600083518551146119d357600080fd5b82518551146119e157600080fd5b81518551146119ef57600080fd5b600080805b8751811015611a9657611a5f888281518110611a0c57fe5b60200260200101518a898481518110611a2157fe5b602002602001015161ffff1681518110611a3757fe5b6020026020010151888481518110611a4b57fe5b602002602001015188858151811061162857fe5b60408051602080820196909652808201839052815180820383018152606090910190915280519401939093209291506001016119f4565b5090979650505050505050565b60408051600480825260a0820190925260009160609190816020015b611ac76121a5565b815260200190600190039081611abf579050509050611ae586611e4a565b81600081518110611af257fe5b6020908102919091010152611b0683611ea4565b81600181518110611b1357fe5b6020026020010181905250611b2784611ea4565b81600281518110611b3457fe5b6020908102919091010152611b526001600160581b03198616611ea4565b81600381518110611b5f57fe5b6020026020010181905250611b7661184482611efe565b519695505050505050565b6000611bbe6040518060600160405280611b9b6000611ea4565b8152602001611ba986611e4a565b8152602001611bb785611e4a565b9052611ff1565b9392505050565b6000868686868686604051602001808763ffffffff1663ffffffff1660e01b8152600401866001600160801b03166001600160801b031660801b81526010018563ffffffff1663ffffffff1660e01b81526004018481526020018361ffff1661ffff1660f01b8152600201828051906020019060200280838360005b83811015611c59578181015183820152602001611c41565b5050505090500196505050505050506040516020818303038152906040528051906020012090509695505050505050565b6060600080600080611c9b86611857565b9050606081604051908082528060200260200182016040528015611cc9578160200160208202803883390190505b50905060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818a6040516020018083805190602001908083835b60208310611d3c5780518252601f199092019160209182019101611d1d565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190528251920191909120925060009150505b84811015611e2157611d8c8a82611971565b6040805160008152602080820180845288905260ff86168284015260608201859052608082018490529151949c50929a5090985060019260a080840193601f198301929081900390910190855afa158015611deb573d6000803e3d6000fd5b50505060206040510351848281518110611e0157fe5b6001600160a01b0390921660209283029190910190910152600101611d7a565b50919998505050505050505050565b6000611bbe6040518060600160405280611b9b6001611ea4565b611e526121a5565b604080516060810182528381528151600080825260208281019094529192830191611e93565b611e806121a5565b815260200190600190039081611e785790505b508152600260209091015292915050565b611eac6121a5565b604080516060810182528381528151600080825260208281019094529192830191611eed565b611eda6121a5565b815260200190600190039081611ed25790505b508152600060209091015292915050565b611f066121a5565b611f108251612070565b611f1957600080fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b611f466121c9565b6040820151600c60ff90911610611f5c57600080fd5b604082015160ff16611f89576040518060200160405280611f808460000151612077565b90529050611881565b604082015160ff1660021415611fae5750604080516020810190915281518152611881565b600360ff16826040015160ff1610158015611fd257506040820151600c60ff909116105b15611fef576040518060200160405280611f80846020015161209b565bfe5b6040805160038082526080820190925260009160609190816020015b6120156121a5565b81526020019060019003908161200d57905050905060005b81518110156120665783816003811061204257fe5b602002015182828151811061205357fe5b602090810291909101015260010161202d565b50611bbe8161209b565b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b60006008825111156120ac57600080fd5b606082516040519080825280602002602001820160405280156120d9578160200160208202803883390190505b50905060005b8151811015612133576120f06121c9565b61210c8583815181106120ff57fe5b6020026020010151611f3e565b9050806000015183838151811061211f57fe5b6020908102919091010152506001016120df565b508251600360ff160181604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561217c578181015183820152602001612164565b505050509050019250505060405160208183030381529060405280519060200120915050919050565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a7230582098712d4e8ea07ec14c6eb331c474f232a708881439bb6e4d5ca3e490610e1e0164736f6c63430005090032`

// DeployArbProtocol deploys a new Ethereum contract, binding an instance of ArbProtocol to it.
func DeployArbProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbProtocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
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

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) AppendInboxMessages(opts *bind.CallOpts, _inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "appendInboxMessages", _inboxHash, _pendingMessages)
	return *ret0, err
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) AppendInboxMessages(_inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxMessages(&_ArbProtocol.CallOpts, _inboxHash, _pendingMessages)
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) AppendInboxMessages(_inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxMessages(&_ArbProtocol.CallOpts, _inboxHash, _pendingMessages)
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) AppendInboxPendingMessage(opts *bind.CallOpts, _pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "appendInboxPendingMessage", _pendingMessages, _newMessage)
	return *ret0, err
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) AppendInboxPendingMessage(_pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxPendingMessage(&_ArbProtocol.CallOpts, _pendingMessages, _newMessage)
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) AppendInboxPendingMessage(_pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxPendingMessage(&_ArbProtocol.CallOpts, _pendingMessages, _newMessage)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCaller) CalculateBeforeValues(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	ret0 := new([]*big.Int)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "calculateBeforeValues", _tokenTypes, _messageTokenNums, _messageAmounts)
	return *ret0, err
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCallerSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_ArbProtocol *ArbProtocolCaller) CountSignatures(opts *bind.CallOpts, _signatures []byte) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "countSignatures", _signatures)
	return *ret0, err
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_ArbProtocol *ArbProtocolSession) CountSignatures(_signatures []byte) (*big.Int, error) {
	return _ArbProtocol.Contract.CountSignatures(&_ArbProtocol.CallOpts, _signatures)
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_ArbProtocol *ArbProtocolCallerSession) CountSignatures(_signatures []byte) (*big.Int, error) {
	return _ArbProtocol.Contract.CountSignatures(&_ArbProtocol.CallOpts, _signatures)
}

// CreateVMHash is a free data retrieval call binding the contract method 0xe440673b.
//
// Solidity: function createVMHash(uint32 _gracePeriod, uint128 _escrowRequired, uint32 _maxExecutionSteps, bytes32 _vmState, uint16 _challengeManagerNum, address[] _assertKeys) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) CreateVMHash(opts *bind.CallOpts, _gracePeriod uint32, _escrowRequired *big.Int, _maxExecutionSteps uint32, _vmState [32]byte, _challengeManagerNum uint16, _assertKeys []common.Address) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "createVMHash", _gracePeriod, _escrowRequired, _maxExecutionSteps, _vmState, _challengeManagerNum, _assertKeys)
	return *ret0, err
}

// CreateVMHash is a free data retrieval call binding the contract method 0xe440673b.
//
// Solidity: function createVMHash(uint32 _gracePeriod, uint128 _escrowRequired, uint32 _maxExecutionSteps, bytes32 _vmState, uint16 _challengeManagerNum, address[] _assertKeys) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) CreateVMHash(_gracePeriod uint32, _escrowRequired *big.Int, _maxExecutionSteps uint32, _vmState [32]byte, _challengeManagerNum uint16, _assertKeys []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.CreateVMHash(&_ArbProtocol.CallOpts, _gracePeriod, _escrowRequired, _maxExecutionSteps, _vmState, _challengeManagerNum, _assertKeys)
}

// CreateVMHash is a free data retrieval call binding the contract method 0xe440673b.
//
// Solidity: function createVMHash(uint32 _gracePeriod, uint128 _escrowRequired, uint32 _maxExecutionSteps, bytes32 _vmState, uint16 _challengeManagerNum, address[] _assertKeys) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) CreateVMHash(_gracePeriod uint32, _escrowRequired *big.Int, _maxExecutionSteps uint32, _vmState [32]byte, _challengeManagerNum uint16, _assertKeys []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.CreateVMHash(&_ArbProtocol.CallOpts, _gracePeriod, _escrowRequired, _maxExecutionSteps, _vmState, _challengeManagerNum, _assertKeys)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHash", _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x25200160.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _messageDataHashes, uint16[] _messageTokenNum, uint256[] _messageValueAmounts, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHashStub(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageDataHashes [][32]byte, _messageTokenNum []uint16, _messageValueAmounts []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHashStub", _tokenTypes, _messageDataHashes, _messageTokenNum, _messageValueAmounts, _messageDestination)
	return *ret0, err
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _messageDataHashes, uint16[] _messageTokenNum, uint256[] _messageValueAmounts, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _messageDataHashes [][32]byte, _messageTokenNum []uint16, _messageValueAmounts []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _messageDataHashes, _messageTokenNum, _messageValueAmounts, _messageDestination)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xb3277495.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _messageDataHashes, uint16[] _messageTokenNum, uint256[] _messageValueAmounts, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _messageDataHashes [][32]byte, _messageTokenNum []uint16, _messageValueAmounts []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _messageDataHashes, _messageTokenNum, _messageValueAmounts, _messageDestination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0xccf69dd7.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateSentMessageHash(opts *bind.CallOpts, _dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateSentMessageHash", _dest, _data, _tokenType, _value, _sender)
	return *ret0, err
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateSentMessageHash(_dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateSentMessageHash(&_ArbProtocol.CallOpts, _dest, _data, _tokenType, _value, _sender)
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x2a0500d8.
//
// Solidity: function generateSentMessageHash(bytes32 _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, bytes32 _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateSentMessageHash(_dest [32]byte, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateSentMessageHash(&_ArbProtocol.CallOpts, _dest, _data, _tokenType, _value, _sender)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_ArbProtocol *ArbProtocolCaller) ParseSignature(opts *bind.CallOpts, _signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	ret := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	out := ret
	err := _ArbProtocol.contract.Call(opts, out, "parseSignature", _signatures, _pos)
	return *ret, err
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_ArbProtocol *ArbProtocolSession) ParseSignature(_signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _ArbProtocol.Contract.ParseSignature(&_ArbProtocol.CallOpts, _signatures, _pos)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_ArbProtocol *ArbProtocolCallerSession) ParseSignature(_signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _ArbProtocol.Contract.ParseSignature(&_ArbProtocol.CallOpts, _signatures, _pos)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_ArbProtocol *ArbProtocolCaller) RecoverAddresses(opts *bind.CallOpts, _messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	ret0 := new([]common.Address)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "recoverAddresses", _messageHash, _signatures)
	return *ret0, err
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_ArbProtocol *ArbProtocolSession) RecoverAddresses(_messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	return _ArbProtocol.Contract.RecoverAddresses(&_ArbProtocol.CallOpts, _messageHash, _signatures)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_ArbProtocol *ArbProtocolCallerSession) RecoverAddresses(_messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	return _ArbProtocol.Contract.RecoverAddresses(&_ArbProtocol.CallOpts, _messageHash, _signatures)
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) UnanimousAssertHash(opts *bind.CallOpts, _fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "unanimousAssertHash", _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
	return *ret0, err
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) UnanimousAssertHash(_fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.UnanimousAssertHash(&_ArbProtocol.CallOpts, _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x014bba5b.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, bytes32[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) UnanimousAssertHash(_fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination [][32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.UnanimousAssertHash(&_ArbProtocol.CallOpts, _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// ArbValueABI is the input ABI used to generate the binding from.
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"get_next_valid_value\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserialize_valid_value_hash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserialize_value_hash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueBin is the compiled bytecode used for deploying new contracts.
const ArbValueBin = `0x610b6c610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c806353409fab1161006557806353409fab14610221578063615c39b01461024757806392516ac714610308578063b2b9dc62146103ae57610092565b80631667b41114610097578063264f384b146100c6578063364df277146100f25780634d00ef7a146100fa575b600080fd5b6100b4600480360360208110156100ad57600080fd5b50356103df565b60408051918252519081900360200190f35b6100b4600480360360608110156100dc57600080fd5b5060ff8135169060208101359060400135610405565b6100b4610457565b6101a26004803603604081101561011057600080fd5b81019060208101813564010000000081111561012b57600080fd5b82018360208201111561013d57600080fd5b8035906020019184600183028401116401000000008311171561015f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506104ca915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101e55781810151838201526020016101cd565b50505050905090810190601f1680156102125780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100b46004803603604081101561023757600080fd5b5060ff8135169060200135610517565b6102ef6004803603604081101561025d57600080fd5b81019060208101813564010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061055e915050565b6040805192835260208301919091528051918290030190f35b6100b46004803603602081101561031e57600080fd5b81019060208101813564010000000081111561033957600080fd5b82018360208201111561034b57600080fd5b8035906020019184600183028401116401000000008311171561036d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105a1945050505050565b6103cb600480360360208110156103c457600080fd5b50356105dd565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156104a357818101518382015260200161048b565b50505050905001925050506040516020818303038152906040528051906020012091505090565b600060606000806104d9610b01565b6104e387876105e4565b9194509250905082156104f557600080fd5b81610509888880840363ffffffff6106dc16565b945094505050509250929050565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b60008060008061056c610b01565b61057687876105e4565b91945092509050821561058857600080fd5b816105928261075c565b51909890975095505050505050565b600080806105ad610b01565b6105b88560006105e4565b9194509250905082156105ca57600080fd5b6105d38161075c565b5195945050505050565b6008101590565b6000806105ef610b01565b60008585815181106105fd57fe5b016020015160019095019460f81c905060008161063e5761061e878761080f565b909650905060008661062f83610832565b919650945092506106d5915050565b60ff82166002141561066557610654878761080f565b909650905060008661062f8361088c565b600360ff83161080159061067c5750600c60ff8316105b156106b757600219820160606000610695838b8b6108e6565b909a509250905080896106a78461099b565b97509750975050505050506106d5565b8160ff166127100160006106cb6000610832565b9196509450925050505b9250925092565b6060818301845110156106ee57600080fd5b60608215801561070957604051915060208201604052610753565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561074257805183526020928301920161072a565b5050858452601f01601f1916604052505b50949350505050565b610764610b25565b6040820151600c60ff9091161061077a57600080fd5b604082015160ff166107a757604051806020016040528061079e84600001516103df565b90529050610400565b604082015160ff16600214156107cc5750604080516020810190915281518152610400565b600360ff16826040015160ff16101580156107f057506040820151600c60ff909116105b1561080d57604051806020016040528061079e84602001516109db565bfe5b60008080610823858563ffffffff610ae516565b60209490940195939450505050565b61083a610b01565b60408051606081018252838152815160008082526020828101909452919283019161087b565b610868610b01565b8152602001906001900390816108605790505b508152600060209091015292915050565b610894610b01565b6040805160608101825283815281516000808252602082810190945291928301916108d5565b6108c2610b01565b8152602001906001900390816108ba5790505b508152600260209091015292915050565b6000806060600060608760ff1660405190808252806020026020018201604052801561092c57816020015b610919610b01565b8152602001906001900390816109115790505b50905060005b8860ff168160ff1610156109865761094a88886105e4565b8451859060ff861690811061095b57fe5b602090810291909101015297509250821561097e57509093508492509050610992565b600101610932565b50600094508593509150505b93509350939050565b6109a3610b01565b6109ad82516105dd565b6109b657600080fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b60006008825111156109ec57600080fd5b60608251604051908082528060200260200182016040528015610a19578160200160208202803883390190505b50905060005b8151811015610a7357610a30610b25565b610a4c858381518110610a3f57fe5b602002602001015161075c565b90508060000151838381518110610a5f57fe5b602090810291909101015250600101610a1f565b508251600360ff160181604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610abc578181015183820152602001610aa4565b505050509050019250505060405160208183030381529060405280519060200120915050919050565b60008160200183511015610af857600080fd5b50016020015190565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a723058205f519cde41045de65b3c250d163984e9b8cab11f031f580ef558872415345c3c64736f6c63430005090032`

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

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x615c39b0.
//
// Solidity: function deserialize_valid_value_hash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValidValueHash(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "deserialize_valid_value_hash", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x615c39b0.
//
// Solidity: function deserialize_valid_value_hash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x615c39b0.
//
// Solidity: function deserialize_valid_value_hash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x92516ac7.
//
// Solidity: function deserialize_value_hash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValueHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	ret0 := new([32]byte)

	out := ret0
	err := _ArbValue.contract.Call(opts, out, "deserialize_value_hash", data)
	return *ret0, err
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x92516ac7.
//
// Solidity: function deserialize_value_hash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x92516ac7.
//
// Solidity: function deserialize_value_hash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x4d00ef7a.
//
// Solidity: function get_next_valid_value(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCaller) GetNextValidValue(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "get_next_valid_value", data, offset)
	return *ret0, *ret1, err
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x4d00ef7a.
//
// Solidity: function get_next_valid_value(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x4d00ef7a.
//
// Solidity: function get_next_valid_value(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCallerSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointBasicValue(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	ret0 := new([32]byte)

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
	ret0 := new([32]byte)

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
	ret0 := new([32]byte)

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
	ret0 := new([32]byte)

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
	ret0 := new(bool)

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
const BytesLibBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820c6694409c11fc2dfd409c9a86b8959cc4bc5705ed732a15d129c94993240c58b64736f6c63430005090032`

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

// ChallengeManagerABI is the input ABI used to generate the binding from.
const ChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_rootHash\",\"type\":\"bytes32\"},{\"name\":\"_deadline\",\"type\":\"uint64\"}],\"name\":\"asserterTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"name\":\"_challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[3]\"},{\"name\":\"_afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"name\":\"_totalMessageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_totalSteps\",\"type\":\"uint32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"},{\"name\":\"_deadline\",\"type\":\"uint64\"}],\"name\":\"bisectAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_assertionToChallenge\",\"type\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"bytes\"},{\"name\":\"_deadline\",\"type\":\"uint64\"},{\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"continueChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_beforeHashAndInbox\",\"type\":\"bytes32[2]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"},{\"name\":\"_afterHashAndMessages\",\"type\":\"bytes32[5]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_proof\",\"type\":\"bytes\"},{\"name\":\"_deadline\",\"type\":\"uint64\"}],\"name\":\"oneStepProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_rootHash\",\"type\":\"bytes32\"},{\"name\":\"_deadline\",\"type\":\"uint64\"}],\"name\":\"challengerTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vmTracker\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"afterHashAndMessageAndLogsBisections\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"totalSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"totalMessageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assertionIndex\",\"type\":\"uint256\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"proofData\",\"type\":\"bytes32[10]\"}],\"name\":\"OneStepProofDebug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challengerWrong\",\"type\":\"bool\"}],\"name\":\"TimedOutChallenge\",\"type\":\"event\"}]"

// ChallengeManagerBin is the compiled bytecode used for deploying new contracts.
const ChallengeManagerBin = `0x608060405234801561001057600080fd5b5060405161274c38038061274c8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556126e7806100656000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632923a1a7146100675780632b50d42b1461009c578063794614bd146100d5578063bef810f514610374578063c179777714610439578063f2b925ec146106f9575b600080fd5b61009a6004803603606081101561007d57600080fd5b508035906020810135906040013567ffffffffffffffff1661072c565b005b61009a600480360360e08110156100b257600080fd5b508035906020810190606081019063ffffffff60a0820135169060c0013561086d565b61009a60048036036101608110156100ec57600080fd5b810190808060600190600380602002604051908101604052809291908260036020028082843760009201919091525091949392602081019250359050600160201b81111561013957600080fd5b82018360208201111561014b57600080fd5b803590602001918460208302840111600160201b8311171561016c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101bb57600080fd5b8201836020820111156101cd57600080fd5b803590602001918460208302840111600160201b831117156101ee57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080518082018252939663ffffffff8635169690959094606082019450925060200190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561027557600080fd5b82018360208201111561028757600080fd5b803590602001918460208302840111600160201b831117156102a857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102f757600080fd5b82018360208201111561030957600080fd5b803590602001918460208302840111600160201b8311171561032a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505050903567ffffffffffffffff169150610a3f9050565b61009a600480360360c081101561038a57600080fd5b813591602081013591810190606081016040820135600160201b8111156103b057600080fd5b8201836020820111156103c257600080fd5b803590602001918460018302840111600160201b831117156103e357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505067ffffffffffffffff8335169350505060208101359060400135610f76565b61009a60048036036101e081101561045057600080fd5b604080518082018252833593928301929160608301919060208401906002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b8111156104c757600080fd5b8201836020820111156104d957600080fd5b803590602001918460208302840111600160201b831117156104fa57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561054957600080fd5b82018360208201111561055b57600080fd5b803590602001918460208302840111600160201b8311171561057c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160a0818101909252939695948181019493509150600590839083908082843760009201919091525091949392602081019250359050600160201b8111156105f757600080fd5b82018360208201111561060957600080fd5b803590602001918460208302840111600160201b8311171561062a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561067957600080fd5b82018360208201111561068b57600080fd5b803590602001918460018302840111600160201b831117156106ac57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903567ffffffffffffffff1691506112a09050565b61009a6004803603606081101561070f57600080fd5b508035906020810135906040013567ffffffffffffffff166119d3565b60008381526001602090815260409182902080548351808401879052600160f81b818601526001600160c01b031960c087901b16604182015284516029818303018152604990910190945283519390920192909220146107ce576040805162461bcd60e51b8152602060048201526018602482015277496e636f72726563742070726576696f757320737461746560401b604482015290519081900360640190fd5b8167ffffffffffffffff164311610826576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b6108308482611b14565b6040805160018152905185917f4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680919081900360200190a250505050565b6000546001600160a01b031633146108b65760405162461bcd60e51b815260040180806020018281038252602e815260200180612608602e913960400191505060405180910390fd5b600085815260016020526040902054156109015760405162461bcd60e51b81526004018080602001828103825260238152602001806126906023913960400191505060405180910390fd5b60405180608001604052808260018563ffffffff1643016040516020018084815260200183600281111561093157fe5b60ff1660f81b81526001018267ffffffffffffffff1667ffffffffffffffff1660c01b815260080193505050506040516020818303038152906040528051906020012081526020018460028060200260405190810160405280929190826002602002808284376000920191909152505050815260408051808201825260209092019190879060029083908390808284376000920182905250928452505063ffffffff85166020928301528781526001808352604090912083518155918301516109fe918301906002612460565b506040820151610a149060028084019190612505565b50606091909101516004909101805463ffffffff191663ffffffff9092169190911790555050505050565b82511580610a5657508251865181610a5357fe5b06155b610aa0576040805162461bcd60e51b8152602060048201526016602482015275092dcc6dee4e4cac6e840d2dce0eae840d8cadccee8d60531b604482015290519081900360640190fd5b82511580610ac657508251865181610ab457fe5b046003600189510381610ac357fe5b04145b610b10576040805162461bcd60e51b8152602060048201526016602482015275092dcc6dee4e4cac6e840d2dce0eae840d8cadccee8d60531b604482015290519081900360640190fd5b8151835114610b5f576040805162461bcd60e51b8152602060048201526016602482015275092dcc6dee4e4cac6e840d2dce0eae840d8cadccee8d60531b604482015290519081900360640190fd5b60006060610be6604051806101200160405280600360018d510381610b8057fe5b0463ffffffff1681526020018b81526020018a81526020018963ffffffff1681526020018c600160038110610bb157fe5b602002015181526020018881526020018781526020018681526020018c600260038110610bda57fe5b60200201519052611c02565b8b5160009081526001602090815260409182902080548351808401879052600160f81b818601526001600160c01b031960c08b901b166041820152845160298183030181526049909101909452835193909201929092209395509193509114610c96576040805162461bcd60e51b815260206004820152601960248201527f446f6573206e6f74206d61746368207072657620737461746500000000000000604482015290519081900360640190fd5b8367ffffffffffffffff16431115610cf5576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60028101600001546001600160a01b03163314610d59576040805162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f7269676e616c2061737365727465722063616e20626973656374604482015290519081900360640190fd5b6040516309898dc160e41b815260206004820181815284516024840152845173__$a50780cb42d41d2927e39f529dc62d6697$__93639898dc1093879392839260440191808601910280838360005b83811015610dc0578181015183820152602001610da8565b505050509050019250505060206040518083038186803b158015610de357600080fd5b505af4158015610df7573d6000803e3d6000fd5b505050506040513d6020811015610e0d57600080fd5b5051600482015460408051602080820194909452600160f91b8183015263ffffffff909216430160c01b6001600160c01b03191660418301528051808303602901815260499092019052805191012081558a600060200201517fa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef716002830160000160009054906101000a90046001600160a01b03168c8b8d60405180856001600160a01b03166001600160a01b03168152602001806020018463ffffffff1663ffffffff16815260200180602001838103835286818151815260200191508051906020019060200280838360005b83811015610f13578181015183820152602001610efb565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610f52578181015183820152602001610f3a565b50505050905001965050505050505060405180910390a25050505050505050505050565b60008681526001602090815260409182902080548351808401879052600160f91b818601526001600160c01b031960c089901b16604182015284516029818303018152604990910190945283519390920192909220146110075760405162461bcd60e51b815260040180806020018281038252602b815260200180612636602b913960400191505060405180910390fd5b8367ffffffffffffffff16431115611066576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60028101600101546001600160a01b031633146110b45760405162461bcd60e51b815260040180806020018281038252602f815260200180612661602f913960400191505060405180910390fd5b73__$a50780cb42d41d2927e39f529dc62d6697$__63b792d7678685858a6001016040518563ffffffff1660e01b81526004018080602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b8381101561113357818101518382015260200161111b565b50505050905090810190601f1680156111605780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b15801561118057600080fd5b505af4158015611194573d6000803e3d6000fd5b505050506040513d60208110156111aa57600080fd5b50516111fd576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420617373657274696f6e2073656c6563746564000000000000604482015290519081900360640190fd5b6004810154604080516020808201869052600160f81b8284015263ffffffff909316430160c01b6001600160c01b031916604182015281518082036029018152604990910190915280519101208155867ff5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf06002830160010154604080516001600160a01b039092168252602082018a90528051918290030190a250505050505050565b600089815260016020526040902067ffffffffffffffff821643111561130d576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6520737465702070726f6f66206d697373656420646561646c696e650000604482015290519081900360640190fd5b805473__$6b4cc75dad3e0abd6ad83b3d907747c608$__633e2855988b600060200201518b8d600160200201518c8c6040518663ffffffff1660e01b81526004018086815260200185600260200280838360005b83811015611379578181015183820152602001611361565b505050509050018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156113c65781810151838201526020016113ae565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156114055781810151838201526020016113ed565b5050505090500197505050505050505060206040518083038186803b15801561142d57600080fd5b505af4158015611441573d6000803e3d6000fd5b505050506040513d602081101561145757600080fd5b505173__$6b4cc75dad3e0abd6ad83b3d907747c608$__6320903721886000602002015160018a8160200201518b600260200201518c600360200201518d600460200201518d6040518863ffffffff1660e01b8152600401808881526020018763ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156115165781810151838201526020016114fe565b505050509050019850505050505050505060206040518083038186803b15801561153f57600080fd5b505af4158015611553573d6000803e3d6000fd5b505050506040513d602081101561156957600080fd5b505160408051602080820194909452808201929092528051808303820181526060830182528051908401206080830152600160f81b60a08301526001600160c01b031960c087901b1660a18301528051608981840301815260a990920190528051910120146116095760405162461bcd60e51b81526004018080602001828103825260268152602001806125e26026913960400191505060405180910390fd5b600073__$98107b176d4310ec680b0534b46d40334a$__630eca9f136040518060e001604052808d60006002811061163d57fe5b602002015181526020018d60016002811061165457fe5b602002015181526020018960006005811061166b57fe5b602002015181526020018960016005811061168257fe5b602002015181526020018960026005811061169957fe5b60200201518152602001896003600581106116b057fe5b60200201518152602001896004600581106116c757fe5b60200201518152508b8b8b8a8a6040518763ffffffff1660e01b81526004018087600760200280838360005b8381101561170b5781810151838201526020016116f3565b5050505090500186600260200280838360005b8381101561173657818101518382015260200161171e565b5050505090500180602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b8381101561178557818101518382015260200161176d565b50505050905001858103845288818151815260200191508051906020019060200280838360005b838110156117c45781810151838201526020016117ac565b50505050905001858103835287818151815260200191508051906020019060200280838360005b838110156118035781810151838201526020016117eb565b50505050905001858103825286818151815260200191508051906020019080838360005b8381101561183f578181015183820152602001611827565b50505050905090810190601f16801561186c5780820380516001836020036101000a031916815260200191505b509a505050505050505050505060206040518083038186803b15801561189157600080fd5b505af41580156118a5573d6000803e3d6000fd5b505050506040513d60208110156118bb57600080fd5b505190508015611908576040805162461bcd60e51b8152602060048201526013602482015272141c9bdbd9881dd85cc81a5b98dbdc9c9958dd606a1b604482015290519081900360640190fd5b6119128b836123c8565b8a7fb3b8389d4e8e450396d43f12a728c7a73784c6089340d0e904a24fd366ddaaff338660405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561198b578181015183820152602001611973565b50505050905090810190601f1680156119b85780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25050505050505050505050565b60008381526001602090815260409182902080548351808401879052600160f91b818601526001600160c01b031960c087901b1660418201528451602981830301815260499091019094528351939092019290922014611a75576040805162461bcd60e51b8152602060048201526018602482015277496e636f72726563742070726576696f757320737461746560401b604482015290519081900360640190fd5b8167ffffffffffffffff164311611acd576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b611ad784826123c8565b6040805160008152905185917f4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680919081900360200190a250505050565b600080546040805180820182529283526001840154600160801b81046001600160801b0390811660029282168390040116602085015290516363d8463760e01b8152600481018681526001600160a01b03909316936363d84637938793870192606481019084906024015b81546001600160a01b03168152600190910190602001808311611b7f5750839050604080838360005b83811015611bc0578181015183820152602001611ba8565b505050509050019350505050600060405180830381600087803b158015611be657600080fd5b505af1158015611bfa573d6000803e3d6000fd5b505050505050565b60006060611c0e612559565b836000015163ffffffff16604051908082528060200260200182016040528015611c42578160200160208202803883390190505b508160600181905250836000015163ffffffff16846060015163ffffffff1681611c6857fe5b0460010163ffffffff1660a0820152608084015181526000805b855163ffffffff168110156123b157856000015163ffffffff16866060015163ffffffff1681611cae57fe5b0663ffffffff16811415611ccf5760a0830180516000190163ffffffff1690525b8560c0015151604051908082528060200260200182016040528015611cfe578160200160208202803883390190505b506080840152600091505b8560c0015151821015611d63578560400151828760c001515183020181518110611d2f57fe5b602002602001015183608001518381518110611d4757fe5b6020908102919091010180519091019052600190910190611d09565b73__$6b4cc75dad3e0abd6ad83b3d907747c608$__633e28559884600001518860a001518961010001518a60c001518b60e001516040518663ffffffff1660e01b81526004018086815260200185600260200280838360005b83811015611dd4578181015183820152602001611dbc565b505050509050018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015611e21578181015183820152602001611e09565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611e60578181015183820152602001611e48565b5050505090500197505050505050505060206040518083038186803b158015611e8857600080fd5b505af4158015611e9c573d6000803e3d6000fd5b505050506040513d6020811015611eb257600080fd5b50516020840152600091505b8560c0015151821015611f105782608001518281518110611edb57fe5b60200260200101518660e001518381518110611ef357fe5b602090810291909101018051919091039052600190910190611ebe565b826020015173__$6b4cc75dad3e0abd6ad83b3d907747c608$__632090372188602001518481518110611f3f57fe5b60200260200101518660a001518a60200151868c6000015163ffffffff160181518110611f6857fe5b60200260200101518b60200151878d6000015163ffffffff160160010181518110611f8f57fe5b60200260200101518c60200151888e6000015160020260010163ffffffff160181518110611fb957fe5b60200260200101518d60200151898f6000015160020260020163ffffffff160181518110611fe357fe5b60200260200101518b608001516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b8381101561206f578181015183820152602001612057565b505050509050019850505050505050505060206040518083038186803b15801561209857600080fd5b505af41580156120ac573d6000803e3d6000fd5b505050506040513d60208110156120c257600080fd5b50516040805160208181019490945280820192909252805180830382018152606092830190915280519201919091209084015180518390811061210157fe5b6020026020010181815250508560200151818151811061211d57fe5b60209081029190910101518352806123a9578560c001515160405190808252806020026020018201604052801561215e578160200160208202803883390190505b506080840152600091505b8560400151518210156121c8578560400151828151811061218657fe5b602002602001015183608001518760c001515184816121a157fe5b06815181106121ac57fe5b6020908102919091010180519091019052600190910190612169565b826020015173__$6b4cc75dad3e0abd6ad83b3d907747c608$__6320903721886020015160018a600001510363ffffffff168151811061220457fe5b602002602001015189606001518a602001518b6000015163ffffffff168151811061222b57fe5b60200260200101518b602001518c6000015160020263ffffffff168151811061225057fe5b60200260200101518c602001518d6000015160020260010163ffffffff168151811061227857fe5b60200260200101518d6020015160018f6020015151038151811061229857fe5b60200260200101518b608001516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b8381101561232457818101518382015260200161230c565b505050509050019850505050505050505060206040518083038186803b15801561234d57600080fd5b505af4158015612361573d6000803e3d6000fd5b505050506040513d602081101561237757600080fd5b505160408051602081810194909452808201929092528051808303820181526060909201815281519190920120908401525b600101611c82565b505060408101516060909101519092509050915091565b6000805460408051808201825260018501546001600160801b038082166002600160801b909304821683900401168252602082019490945290516363d8463760e01b81526004810186815293850180546001600160a01b039081166024840152909316936363d84637938793909290919060648101906003890190604401808311611b7f575050825181528260408083836020611ba8565b6001830191839082156124f55791602002820160005b838211156124c057835183826101000a8154816001600160801b0302191690836001600160801b031602179055509260200192601001602081600f01049283019260010302612476565b80156124f35782816101000a8154906001600160801b030219169055601001602081600f010492830192600103026124c0565b505b5061250192915061258d565b5090565b826002810192821561254d579160200282015b8281111561254d57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190612518565b506125019291506125bd565b6040805160c0810182526000808252602082018190529181018290526060808201819052608082015260a081019190915290565b6125ba91905b808211156125015780546fffffffffffffffffffffffffffffffff19168155600101612593565b90565b6125ba91905b808211156125015780546001600160a01b03191681556001016125c356fe4f6e6520737465702070726f6f66207769746820696e76616c696420707265762073746174654368616c6c656e6765206d75737420626520666f727761726465642066726f6d206d61696e20636f6e7472616374636f6e74696e75654368616c6c656e67653a20496e636f72726563742070726576696f75732073746174654f6e6c79206f726967696e616c206368616c6c656e6765722063616e20636f6e74696e7565206368616c6c656e67655468657265206d757374206265206e6f206578697374696e67206368616c6c656e6765a265627a7a723058202f0b87f6635a49870d39faee817e77b97add6f6f8fec7b2c6e476454afac78e864736f6c63430005090032`

// DeployChallengeManager deploys a new Ethereum contract, binding an instance of ChallengeManager to it.
func DeployChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend, _vmTracker common.Address) (common.Address, *types.Transaction, *ChallengeManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeManagerBin), backend, _vmTracker)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// ChallengeManager is an auto generated Go binding around an Ethereum contract.
type ChallengeManager struct {
	ChallengeManagerCaller     // Read-only binding to the contract
	ChallengeManagerTransactor // Write-only binding to the contract
	ChallengeManagerFilterer   // Log filterer for contract events
}

// ChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeManagerSession struct {
	Contract     *ChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeManagerCallerSession struct {
	Contract *ChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeManagerTransactorSession struct {
	Contract     *ChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeManagerRaw struct {
	Contract *ChallengeManager // Generic contract binding to access the raw methods on
}

// ChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeManagerCallerRaw struct {
	Contract *ChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactorRaw struct {
	Contract *ChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeManager creates a new instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManager(address common.Address, backend bind.ContractBackend) (*ChallengeManager, error) {
	contract, err := bindChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// NewChallengeManagerCaller creates a new read-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*ChallengeManagerCaller, error) {
	contract, err := bindChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerCaller{contract: contract}, nil
}

// NewChallengeManagerTransactor creates a new write-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeManagerTransactor, error) {
	contract, err := bindChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTransactor{contract: contract}, nil
}

// NewChallengeManagerFilterer creates a new log filterer instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeManagerFilterer, error) {
	contract, err := bindChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerFilterer{contract: contract}, nil
}

// bindChallengeManager binds a generic wrapper to an already deployed contract.
func bindChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.ChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x2923a1a7.
//
// Solidity: function asserterTimedOut(bytes32 _vmId, bytes32 _rootHash, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerTransactor) AsserterTimedOut(opts *bind.TransactOpts, _vmId [32]byte, _rootHash [32]byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "asserterTimedOut", _vmId, _rootHash, _deadline)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x2923a1a7.
//
// Solidity: function asserterTimedOut(bytes32 _vmId, bytes32 _rootHash, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerSession) AsserterTimedOut(_vmId [32]byte, _rootHash [32]byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.AsserterTimedOut(&_ChallengeManager.TransactOpts, _vmId, _rootHash, _deadline)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x2923a1a7.
//
// Solidity: function asserterTimedOut(bytes32 _vmId, bytes32 _rootHash, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) AsserterTimedOut(_vmId [32]byte, _rootHash [32]byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.AsserterTimedOut(&_ChallengeManager.TransactOpts, _vmId, _rootHash, _deadline)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x794614bd.
//
// Solidity: function bisectAssertion(bytes32[3] _fields, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerTransactor) BisectAssertion(opts *bind.TransactOpts, _fields [3][32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "bisectAssertion", _fields, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes, _beforeBalances, _deadline)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x794614bd.
//
// Solidity: function bisectAssertion(bytes32[3] _fields, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerSession) BisectAssertion(_fields [3][32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertion(&_ChallengeManager.TransactOpts, _fields, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes, _beforeBalances, _deadline)
}

// BisectAssertion is a paid mutator transaction binding the contract method 0x794614bd.
//
// Solidity: function bisectAssertion(bytes32[3] _fields, bytes32[] _afterHashAndMessageAndLogsBisections, uint256[] _totalMessageAmounts, uint32 _totalSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) BisectAssertion(_fields [3][32]byte, _afterHashAndMessageAndLogsBisections [][32]byte, _totalMessageAmounts []*big.Int, _totalSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertion(&_ChallengeManager.TransactOpts, _fields, _afterHashAndMessageAndLogsBisections, _totalMessageAmounts, _totalSteps, _timeBounds, _tokenTypes, _beforeBalances, _deadline)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xf2b925ec.
//
// Solidity: function challengerTimedOut(bytes32 _vmId, bytes32 _rootHash, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ChallengerTimedOut(opts *bind.TransactOpts, _vmId [32]byte, _rootHash [32]byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "challengerTimedOut", _vmId, _rootHash, _deadline)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xf2b925ec.
//
// Solidity: function challengerTimedOut(bytes32 _vmId, bytes32 _rootHash, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerSession) ChallengerTimedOut(_vmId [32]byte, _rootHash [32]byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengerTimedOut(&_ChallengeManager.TransactOpts, _vmId, _rootHash, _deadline)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xf2b925ec.
//
// Solidity: function challengerTimedOut(bytes32 _vmId, bytes32 _rootHash, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ChallengerTimedOut(_vmId [32]byte, _rootHash [32]byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengerTimedOut(&_ChallengeManager.TransactOpts, _vmId, _rootHash, _deadline)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xbef810f5.
//
// Solidity: function continueChallenge(bytes32 _vmId, uint256 _assertionToChallenge, bytes _proof, uint64 _deadline, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ContinueChallenge(opts *bind.TransactOpts, _vmId [32]byte, _assertionToChallenge *big.Int, _proof []byte, _deadline uint64, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "continueChallenge", _vmId, _assertionToChallenge, _proof, _deadline, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xbef810f5.
//
// Solidity: function continueChallenge(bytes32 _vmId, uint256 _assertionToChallenge, bytes _proof, uint64 _deadline, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerSession) ContinueChallenge(_vmId [32]byte, _assertionToChallenge *big.Int, _proof []byte, _deadline uint64, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ContinueChallenge(&_ChallengeManager.TransactOpts, _vmId, _assertionToChallenge, _proof, _deadline, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xbef810f5.
//
// Solidity: function continueChallenge(bytes32 _vmId, uint256 _assertionToChallenge, bytes _proof, uint64 _deadline, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ContinueChallenge(_vmId [32]byte, _assertionToChallenge *big.Int, _proof []byte, _deadline uint64, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ContinueChallenge(&_ChallengeManager.TransactOpts, _vmId, _assertionToChallenge, _proof, _deadline, _bisectionRoot, _bisectionHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 _vmId, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, _vmId [32]byte, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "initiateChallenge", _vmId, _players, _escrows, _challengePeriod, _challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 _vmId, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerSession) InitiateChallenge(_vmId [32]byte, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.InitiateChallenge(&_ChallengeManager.TransactOpts, _vmId, _players, _escrows, _challengePeriod, _challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 _vmId, address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) InitiateChallenge(_vmId [32]byte, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.InitiateChallenge(&_ChallengeManager.TransactOpts, _vmId, _players, _escrows, _challengePeriod, _challengeRoot)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xc1797777.
//
// Solidity: function oneStepProof(bytes32 _vmId, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerTransactor) OneStepProof(opts *bind.TransactOpts, _vmId [32]byte, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "oneStepProof", _vmId, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof, _deadline)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xc1797777.
//
// Solidity: function oneStepProof(bytes32 _vmId, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerSession) OneStepProof(_vmId [32]byte, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProof(&_ChallengeManager.TransactOpts, _vmId, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof, _deadline)
}

// OneStepProof is a paid mutator transaction binding the contract method 0xc1797777.
//
// Solidity: function oneStepProof(bytes32 _vmId, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof, uint64 _deadline) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) OneStepProof(_vmId [32]byte, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte, _deadline uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProof(&_ChallengeManager.TransactOpts, _vmId, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof, _deadline)
}

// ChallengeManagerBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ChallengeManager contract.
type ChallengeManagerBisectedAssertionIterator struct {
	Event *ChallengeManagerBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerBisectedAssertion)
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
		it.Event = new(ChallengeManagerBisectedAssertion)
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
func (it *ChallengeManagerBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerBisectedAssertion represents a BisectedAssertion event raised by the ChallengeManager contract.
type ChallengeManagerBisectedAssertion struct {
	VmId                                 [32]byte
	Bisecter                             common.Address
	AfterHashAndMessageAndLogsBisections [][32]byte
	TotalSteps                           uint32
	TotalMessageAmounts                  []*big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0xa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71.
//
// Solidity: event BisectedAssertion(bytes32 indexed vmId, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_ChallengeManager *ChallengeManagerFilterer) FilterBisectedAssertion(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerBisectedAssertionIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "BisectedAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerBisectedAssertionIterator{contract: _ChallengeManager.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0xa4f8cbbd195d8e69d66b332eb24e05c24a35ee00de81da472b89ecb42a70ef71.
//
// Solidity: event BisectedAssertion(bytes32 indexed vmId, address bisecter, bytes32[] afterHashAndMessageAndLogsBisections, uint32 totalSteps, uint256[] totalMessageAmounts)
func (_ChallengeManager *ChallengeManagerFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ChallengeManagerBisectedAssertion, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "BisectedAssertion", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerBisectedAssertion)
				if err := _ChallengeManager.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ChallengeManagerContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the ChallengeManager contract.
type ChallengeManagerContinuedChallengeIterator struct {
	Event *ChallengeManagerContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerContinuedChallenge)
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
		it.Event = new(ChallengeManagerContinuedChallenge)
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
func (it *ChallengeManagerContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerContinuedChallenge represents a ContinuedChallenge event raised by the ChallengeManager contract.
type ChallengeManagerContinuedChallenge struct {
	VmId           [32]byte
	Challenger     common.Address
	AssertionIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0xf5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf0.
//
// Solidity: event ContinuedChallenge(bytes32 indexed vmId, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) FilterContinuedChallenge(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerContinuedChallengeIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ContinuedChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerContinuedChallengeIterator{contract: _ChallengeManager.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0xf5a6d1468c6ce7a03663fa2fd47dad9f6693a9289ebdbcf3309941caa75c7cf0.
//
// Solidity: event ContinuedChallenge(bytes32 indexed vmId, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerContinuedChallenge, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ContinuedChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerContinuedChallenge)
				if err := _ChallengeManager.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompletedIterator struct {
	Event *ChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerOneStepProofCompleted)
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
		it.Event = new(ChallengeManagerOneStepProofCompleted)
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
func (it *ChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompleted struct {
	VmId     [32]byte
	Asserter common.Address
	Proof    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xb3b8389d4e8e450396d43f12a728c7a73784c6089340d0e904a24fd366ddaaff.
//
// Solidity: event OneStepProofCompleted(bytes32 indexed vmId, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerOneStepProofCompletedIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerOneStepProofCompletedIterator{contract: _ChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xb3b8389d4e8e450396d43f12a728c7a73784c6089340d0e904a24fd366ddaaff.
//
// Solidity: event OneStepProofCompleted(bytes32 indexed vmId, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeManagerOneStepProofCompleted, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerOneStepProofCompleted)
				if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ChallengeManagerOneStepProofDebugIterator is returned from FilterOneStepProofDebug and is used to iterate over the raw logs and unpacked data for OneStepProofDebug events raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofDebugIterator struct {
	Event *ChallengeManagerOneStepProofDebug // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerOneStepProofDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerOneStepProofDebug)
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
		it.Event = new(ChallengeManagerOneStepProofDebug)
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
func (it *ChallengeManagerOneStepProofDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerOneStepProofDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerOneStepProofDebug represents a OneStepProofDebug event raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofDebug struct {
	VmId      [32]byte
	ProofData [10][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofDebug is a free log retrieval operation binding the contract event 0x4c7f6075bf7890867239f4808f15728e984938b95252446d263cfdd1f5f6cb21.
//
// Solidity: event OneStepProofDebug(bytes32 indexed vmId, bytes32[10] proofData)
func (_ChallengeManager *ChallengeManagerFilterer) FilterOneStepProofDebug(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerOneStepProofDebugIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "OneStepProofDebug", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerOneStepProofDebugIterator{contract: _ChallengeManager.contract, event: "OneStepProofDebug", logs: logs, sub: sub}, nil
}

// WatchOneStepProofDebug is a free log subscription operation binding the contract event 0x4c7f6075bf7890867239f4808f15728e984938b95252446d263cfdd1f5f6cb21.
//
// Solidity: event OneStepProofDebug(bytes32 indexed vmId, bytes32[10] proofData)
func (_ChallengeManager *ChallengeManagerFilterer) WatchOneStepProofDebug(opts *bind.WatchOpts, sink chan<- *ChallengeManagerOneStepProofDebug, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "OneStepProofDebug", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerOneStepProofDebug)
				if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofDebug", log); err != nil {
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

// ChallengeManagerTimedOutChallengeIterator is returned from FilterTimedOutChallenge and is used to iterate over the raw logs and unpacked data for TimedOutChallenge events raised by the ChallengeManager contract.
type ChallengeManagerTimedOutChallengeIterator struct {
	Event *ChallengeManagerTimedOutChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerTimedOutChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerTimedOutChallenge)
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
		it.Event = new(ChallengeManagerTimedOutChallenge)
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
func (it *ChallengeManagerTimedOutChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerTimedOutChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerTimedOutChallenge represents a TimedOutChallenge event raised by the ChallengeManager contract.
type ChallengeManagerTimedOutChallenge struct {
	VmId            [32]byte
	ChallengerWrong bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedOutChallenge is a free log retrieval operation binding the contract event 0x4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680.
//
// Solidity: event TimedOutChallenge(bytes32 indexed vmId, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) FilterTimedOutChallenge(opts *bind.FilterOpts, vmId [][32]byte) (*ChallengeManagerTimedOutChallengeIterator, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "TimedOutChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTimedOutChallengeIterator{contract: _ChallengeManager.contract, event: "TimedOutChallenge", logs: logs, sub: sub}, nil
}

// WatchTimedOutChallenge is a free log subscription operation binding the contract event 0x4bb5367552b51194ffa95563d907630acc956760789c62c6e92b82dfddf7e680.
//
// Solidity: event TimedOutChallenge(bytes32 indexed vmId, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) WatchTimedOutChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerTimedOutChallenge, vmId [][32]byte) (event.Subscription, error) {
	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "TimedOutChallenge", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerTimedOutChallenge)
				if err := _ChallengeManager.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
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

// IChallengeManagerABI is the input ABI used to generate the binding from.
const IChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"vmId\",\"type\":\"bytes32\"},{\"name\":\"players\",\"type\":\"address[2]\"},{\"name\":\"escrows\",\"type\":\"uint128[2]\"},{\"name\":\"challengePeriod\",\"type\":\"uint32\"},{\"name\":\"challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeManagerBin is the compiled bytecode used for deploying new contracts.
const IChallengeManagerBin = `0x`

// DeployIChallengeManager deploys a new Ethereum contract, binding an instance of IChallengeManager to it.
func DeployIChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IChallengeManager, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IChallengeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initiateChallenge", vmId, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerSession) InitiateChallenge(vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, vmId, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2b50d42b.
//
// Solidity: function initiateChallenge(bytes32 vmId, address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) InitiateChallenge(vmId [32]byte, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, vmId, players, escrows, challengePeriod, challengeRoot)
}

// IVMTrackerABI is the input ABI used to generate the binding from.
const IVMTrackerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"bytes32\"},{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IVMTrackerBin is the compiled bytecode used for deploying new contracts.
const IVMTrackerBin = `0x`

// DeployIVMTracker deploys a new Ethereum contract, binding an instance of IVMTracker to it.
func DeployIVMTracker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IVMTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(IVMTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IVMTrackerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IVMTracker{IVMTrackerCaller: IVMTrackerCaller{contract: contract}, IVMTrackerTransactor: IVMTrackerTransactor{contract: contract}, IVMTrackerFilterer: IVMTrackerFilterer{contract: contract}}, nil
}

// IVMTracker is an auto generated Go binding around an Ethereum contract.
type IVMTracker struct {
	IVMTrackerCaller     // Read-only binding to the contract
	IVMTrackerTransactor // Write-only binding to the contract
	IVMTrackerFilterer   // Log filterer for contract events
}

// IVMTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVMTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVMTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVMTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVMTrackerSession struct {
	Contract     *IVMTracker       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVMTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVMTrackerCallerSession struct {
	Contract *IVMTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IVMTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVMTrackerTransactorSession struct {
	Contract     *IVMTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IVMTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVMTrackerRaw struct {
	Contract *IVMTracker // Generic contract binding to access the raw methods on
}

// IVMTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVMTrackerCallerRaw struct {
	Contract *IVMTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// IVMTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVMTrackerTransactorRaw struct {
	Contract *IVMTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVMTracker creates a new instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTracker(address common.Address, backend bind.ContractBackend) (*IVMTracker, error) {
	contract, err := bindIVMTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVMTracker{IVMTrackerCaller: IVMTrackerCaller{contract: contract}, IVMTrackerTransactor: IVMTrackerTransactor{contract: contract}, IVMTrackerFilterer: IVMTrackerFilterer{contract: contract}}, nil
}

// NewIVMTrackerCaller creates a new read-only instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerCaller(address common.Address, caller bind.ContractCaller) (*IVMTrackerCaller, error) {
	contract, err := bindIVMTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerCaller{contract: contract}, nil
}

// NewIVMTrackerTransactor creates a new write-only instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVMTrackerTransactor, error) {
	contract, err := bindIVMTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerTransactor{contract: contract}, nil
}

// NewIVMTrackerFilterer creates a new log filterer instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVMTrackerFilterer, error) {
	contract, err := bindIVMTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerFilterer{contract: contract}, nil
}

// bindIVMTracker binds a generic wrapper to an already deployed contract.
func bindIVMTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVMTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVMTracker *IVMTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVMTracker.Contract.IVMTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVMTracker *IVMTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVMTracker.Contract.IVMTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVMTracker *IVMTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVMTracker.Contract.IVMTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVMTracker *IVMTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVMTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVMTracker *IVMTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVMTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVMTracker *IVMTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVMTracker.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactor) CompleteChallenge(opts *bind.TransactOpts, _vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.contract.Transact(opts, "completeChallenge", _vmId, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerSession) CompleteChallenge(_vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _vmId, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x63d84637.
//
// Solidity: function completeChallenge(bytes32 _vmId, address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactorSession) CompleteChallenge(_vmId [32]byte, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _vmId, _players, _rewards)
}

// MerkleLibABI is the input ABI used to generate the binding from.
const MerkleLibABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"generateAddressRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"generateRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
const MerkleLibBin = `0x610573610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043d565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b60005b600182511115610421576060600283516001018161031e57fe5b04604051908082528060200260200182016040528015610348578160200160208202803883390190505b50905060005b81518110156104195783518160020260010110156103e15783816002028151811061037557fe5b602002602001015184826002026001018151811061038f57fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d057fe5b602002602001018181525050610411565b8381600202815181106103f057fe5b602002602001015182828151811061040457fe5b6020026020010181815250505b60010161034e565b509150610304565b8160008151811061042e57fe5b60200260200101519050919050565b600080838160205b88518111610530578089015193506020818a51036020018161046357fe5b0491505b60008211801561047a5750600286066001145b801561048857508160020a86115b1561049b57600286046001019550610467565b600286066104e65783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104de57fe5b049550610528565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052157fe5b0460010195505b602001610445565b50509094149594505050505056fea265627a7a7230582060bd4256cc8dc39e31eb5f7eed0ad08ab723bf0337043846f63b965afb197ae664736f6c63430005090032`

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
	ret0 := new([32]byte)

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
	ret0 := new([32]byte)

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
	ret0 := new(bool)

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
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"fields\",\"type\":\"bytes32[7]\"},{\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"beforeValues\",\"type\":\"uint256[]\"},{\"name\":\"messageValue\",\"type\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"instructionStack\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"dataStack\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"auxStack\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"register\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"errHandler\",\"type\":\"bytes32\"}],\"name\":\"SawMachine\",\"type\":\"event\"}]"

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
const OneStepProofBin = `0x61326c610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c80630eca9f131461003a575b600080fd5b6102c160048036036101a081101561005157600080fd5b810190808060e001906007806020026040519081016040528092919082600760200280828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b8111156100c757600080fd5b8201836020820111156100d957600080fd5b803590602001918460208302840111600160201b831117156100fa57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561014957600080fd5b82018360208201111561015b57600080fd5b803590602001918460208302840111600160201b8311171561017c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101cb57600080fd5b8201836020820111156101dd57600080fd5b803590602001918460208302840111600160201b831117156101fe57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561024d57600080fd5b82018360208201111561025f57600080fd5b803590602001918460018302840111600160201b8311171561028057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102d3945050505050565b60408051918252519081900360200190f35b6080860151606087015160009182918291829114801590600019906104185760005b88518167ffffffffffffffff16101561034a57888167ffffffffffffffff168151811061031e57fe5b6020026020010151600014610342578160070b6000191461033e57600080fd5b8091505b6001016102f5565b508060070b6000191461041357878160070b8151811061036657fe5b60200260200101519350898160070b8151811061037f57fe5b6020026020010151945060019250898160070b8151811061039c57fe5b60200260200101516014601581106103b057fe5b1a60f81b6001600160f81b031916600160f81b14156103f05783898260070b815181106103d957fe5b6020026020010151146103eb57600080fd5b610413565b888160070b815181106103ff57fe5b602002602001015184111561041357600080fd5b610462565b60005b88518167ffffffffffffffff16101561046057888167ffffffffffffffff168151811061044457fe5b602002602001015160001461045857600080fd5b60010161041b565b505b6105436040518061018001604052808e60006007811061047e57fe5b602002015181526020018d81526020018e60016007811061049b57fe5b602002015181526020018e6002600781106104b257fe5b602002015181526020018e6003600781106104c957fe5b602002015181526020018e6004600781106104e057fe5b602002015181526020018e6005600781106104f757fe5b602002015181526020018e60066007811061050e57fe5b60200201518152602001876affffffffffffffffffffff19168152602001868152602001851515815260200189815250610553565b9c9b505050505050505050505050565b60008080806060610562613110565b61056a613110565b61057388611386565b939950929650909450925090506001600060ff881661059a57610595836117c8565b6111da565b60ff8816600114156105e0576105d983866000815181106105b757fe5b6020026020010151876001815181106105cc57fe5b60200260200101516117d2565b91506111da565b60ff88166002141561061f576105d983866000815181106105fd57fe5b60200260200101518760018151811061061257fe5b6020026020010151611822565b60ff88166003141561065e576105d9838660008151811061063c57fe5b60200260200101518760018151811061065157fe5b6020026020010151611863565b60ff88166004141561069d576105d9838660008151811061067b57fe5b60200260200101518760018151811061069057fe5b60200260200101516118a4565b60ff8816600514156106dc576105d983866000815181106106ba57fe5b6020026020010151876001815181106106cf57fe5b60200260200101516118e5565b60ff88166006141561071b576105d983866000815181106106f957fe5b60200260200101518760018151811061070e57fe5b6020026020010151611926565b60ff88166007141561075a576105d9838660008151811061073857fe5b60200260200101518760018151811061074d57fe5b6020026020010151611967565b60ff8816600814156107ae576105d9838660008151811061077757fe5b60200260200101518760018151811061078c57fe5b6020026020010151886002815181106107a157fe5b60200260200101516119a8565b60ff881660091415610802576105d983866000815181106107cb57fe5b6020026020010151876001815181106107e057fe5b6020026020010151886002815181106107f557fe5b6020026020010151611a01565b60ff8816600a1415610841576105d9838660008151811061081f57fe5b60200260200101518760018151811061083457fe5b6020026020010151611a49565b60ff8816600b1415610880576105d9838660008151811061085e57fe5b60200260200101518760018151811061087357fe5b6020026020010151611a8a565b60ff8816601014156108bf576105d9838660008151811061089d57fe5b6020026020010151876001815181106108b257fe5b6020026020010151611acb565b60ff8816601114156108fe576105d983866000815181106108dc57fe5b6020026020010151876001815181106108f157fe5b6020026020010151611b0c565b60ff88166012141561093d576105d9838660008151811061091b57fe5b60200260200101518760018151811061093057fe5b6020026020010151611b4d565b60ff88166013141561095a576105d9838660008151811061091b57fe5b60ff881660141415610999576105d9838660008151811061097757fe5b60200260200101518760018151811061098c57fe5b6020026020010151611b8e565b60ff8816601514156109c3576105d983866000815181106109b657fe5b6020026020010151611bba565b60ff881660161415610a02576105d983866000815181106109e057fe5b6020026020010151876001815181106109f557fe5b6020026020010151611c00565b60ff881660171415610a41576105d98386600081518110610a1f57fe5b602002602001015187600181518110610a3457fe5b6020026020010151611c41565b60ff881660181415610a80576105d98386600081518110610a5e57fe5b602002602001015187600181518110610a7357fe5b6020026020010151611c82565b60ff881660191415610aaa576105d98386600081518110610a9d57fe5b6020026020010151611cc3565b60ff8816601a1415610ae9576105d98386600081518110610ac757fe5b602002602001015187600181518110610adc57fe5b6020026020010151611cf9565b60ff881660201415610b13576105d98386600081518110610b0657fe5b6020026020010151611d3a565b60ff881660301415610b3d576105d98386600081518110610b3057fe5b6020026020010151611d56565b60ff881660311415610b52576105d983611d5e565b60ff881660321415610b67576105d983611d7f565b60ff881660331415610b91576105d98386600081518110610b8457fe5b6020026020010151611d98565b60ff881660341415610bce576105d98386600081518110610bae57fe5b602002602001015160405180602001604052808e60400151815250611db1565b60ff881660351415610bf8576105d98386600081518110610beb57fe5b6020026020010151611e23565b60ff881660361415610c37576105d98386600081518110610c1557fe5b602002602001015187600181518110610c2a57fe5b6020026020010151611e39565b60ff881660371415610c4c576105d983611e6c565b60ff881660381415610c61576105d983611e9e565b60ff881660391415610c8b576105d98386600081518110610c7e57fe5b6020026020010151611eb4565b60ff8816603a1415610d1857610c9f613171565b610cae8b610160015188611ec6565b919950975090508715610cf25760405162461bcd60e51b81526004018080602001828103825260218152602001806132176021913960400191505060405180910390fd5b610d02858263ffffffff611fbe16565b610d12848263ffffffff611fe016565b506111da565b60ff8816603c1415610d29576111da565b60ff881660401415610d53576105d98386600081518110610d4657fe5b6020026020010151611ffd565b60ff881660411415610d92576105d98386600081518110610d7057fe5b602002602001015187600181518110610d8557fe5b602002602001015161201f565b60ff881660421415610de6576105d98386600081518110610daf57fe5b602002602001015187600181518110610dc457fe5b602002602001015188600281518110610dd957fe5b6020026020010151612051565b60ff881660431415610e25576105d98386600081518110610e0357fe5b602002602001015187600181518110610e1857fe5b6020026020010151612093565b60ff881660441415610e79576105d98386600081518110610e4257fe5b602002602001015187600181518110610e5757fe5b602002602001015188600281518110610e6c57fe5b60200260200101516120a5565b60ff881660501415610eb8576105d98386600081518110610e9657fe5b602002602001015187600181518110610eab57fe5b60200260200101516120c7565b60ff881660511415610f0c576105d98386600081518110610ed557fe5b602002602001015187600181518110610eea57fe5b602002602001015188600281518110610eff57fe5b602002602001015161213e565b60ff881660521415610f36576105d98386600081518110610f2957fe5b60200260200101516121b7565b60ff881660531415610f60576105d98386600081518110610f5357fe5b60200260200101516121ea565b60ff881660601415610f7d576105d98386600081518110610b3057fe5b60ff88166061141561103057610fa78386600081518110610f9a57fe5b602002602001015161220d565b60e08c015160c08d0151604080516020808201939093528082018590528151808203830181526060909101909152805191012092945090925014610fea57600080fd5b8960a001518a60800151146105955760405162461bcd60e51b81526004018080602001828103825260278152602001806131f06027913960400191505060405180910390fd5b60ff8816607014156110fc5760008061105d858860008151811061105057fe5b602002602001015161222f565b809450819550829650839750505050508b60a001518c6080015184604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146110af57600080fd5b8b60e001518c60c00151146110f55760405162461bcd60e51b81526004018080602001828103825260268152602001806131ca6026913960400191505060405180910390fd5b50506111da565b60ff88166070141561111c5760008061105d858860008151811061105057fe5b60ff8816607314156111da576040805160028082526060828101909352816020015b611146613171565b81526020019060019003908161113e57505060208c015190915061117b9060005b602002015167ffffffffffffffff1661240d565b8160008151811061118857fe5b60200260200101819052506111a78b6020015160016002811061116757fe5b816001815181106111b457fe5b60200260200101819052506111d86111cb82612467565b859063ffffffff611fe016565b505b8061126b578960a001518a60800151146112255760405162461bcd60e51b81526004018080602001828103825260278152602001806131f06027913960400191505060405180910390fd5b8960e001518a60c001511461126b5760405162461bcd60e51b81526004018080602001828103825260268152602001806131ca6026913960400191505060405180910390fd5b816112cd5760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a08401515114156112c5576112c0836124a7565b6112cd565b60a083015183525b6112d6846124b1565b8a51146113145760405162461bcd60e51b81526004018080602001828103825260228152602001806131a86022913960400191505060405180910390fd5b61131d836124b1565b8a6060015114611374576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b60006060611392613110565b61139a613110565b600080806113a6613110565b6113af81612546565b6113be89610160015184612550565b90945090925090506113ce613110565b6113d782612655565b905060008a610160015185815181106113ec57fe5b602001015160f81c60f81b60f81c905060008b6101600151866001018151811061141257fe5b016020015160f81c90506000611427826126b3565b905060608160405190808252806020026020018201604052801561146557816020015b611452613171565b81526020019060019003908161144a5790505b5090506002880197508360ff166000148061148357508360ff166001145b6114d4576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff8416611577576040805160208082018084528951516353409fab60e01b90915260ff87166024840152604483015291519091829173__$0d86abb4a722a612872fb80f4c7e7e95bd$__916353409fab916064808601929190818703018186803b15801561154257600080fd5b505af4158015611556573d6000803e3d6000fd5b505050506040513d602081101561156c57600080fd5b5051905286526116da565b8360ff16600114156116da5761158b613171565b61159a8f61016001518a611ec6565b909a50909850905087156115f5576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561161957808260008151811061160957fe5b6020026020010181905250611629565b611629868263ffffffff611fe016565b604051806020016040528073__$0d86abb4a722a612872fb80f4c7e7e95bd$__63264f384b87611658866126cd565b518c5151604080516001600160e01b031960e087901b16815260ff909416600485015260248401929092526044830152516064808301926020929190829003018186803b1580156116a857600080fd5b505af41580156116bc573d6000803e3d6000fd5b505050506040513d60208110156116d257600080fd5b505190528752505b60ff84165b8281101561176e576116f68f61016001518a611ec6565b845185908590811061170457fe5b6020908102919091010152995097508715611766576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b6001016116df565b5060ff84165b81518110156117b1576117a982826001855103038151811061179257fe5b602002602001015188611fe090919063ffffffff16565b600101611774565b50919d919c50939a50919850939650945050505050565b600260c090910152565b60006117dd83612780565b15806117ef57506117ed82612780565b155b156117fc5750600061181b565b82518251808201611813878263ffffffff61278b16565b600193505050505b9392505050565b600061182d83612780565b158061183f575061183d82612780565b155b1561184c5750600061181b565b82518251808202611813878263ffffffff61278b16565b600061186e83612780565b1580611880575061187e82612780565b155b1561188d5750600061181b565b82518251808203611813878263ffffffff61278b16565b60006118af83612780565b15806118c157506118bf82612780565b155b156118ce5750600061181b565b82518251808204611813878263ffffffff61278b16565b60006118f083612780565b1580611902575061190082612780565b155b1561190f5750600061181b565b82518251808205611813878263ffffffff61278b16565b600061193183612780565b1580611943575061194182612780565b155b156119505750600061181b565b82518251808206611813878263ffffffff61278b16565b600061197283612780565b1580611984575061198282612780565b155b156119915750600061181b565b82518251808207611813878263ffffffff61278b16565b60006119b384612780565b15806119c557506119c383612780565b155b156119d2575060006119f9565b83518351835160008183850890506119f0898263ffffffff61278b16565b60019450505050505b949350505050565b6000611a0c84612780565b1580611a1e5750611a1c83612780565b155b15611a2b575060006119f9565b83518351835160008183850990506119f0898263ffffffff61278b16565b6000611a5483612780565b1580611a665750611a6482612780565b155b15611a735750600061181b565b8251825180820a611813878263ffffffff61278b16565b6000611a9583612780565b1580611aa75750611aa582612780565b155b15611ab45750600061181b565b8251825180820b611813878263ffffffff61278b16565b6000611ad683612780565b1580611ae85750611ae682612780565b155b15611af55750600061181b565b82518251808210611813878263ffffffff61278b16565b6000611b1783612780565b1580611b295750611b2782612780565b155b15611b365750600061181b565b82518251808211611813878263ffffffff61278b16565b6000611b5883612780565b1580611b6a5750611b6882612780565b155b15611b775750600061181b565b82518251808212611813878263ffffffff61278b16565b6000611bb06111cb611b9f846126cd565b51611ba9866126cd565b511461279f565b5060019392505050565b6000611bc582612780565b611bdf57611bda83600063ffffffff61278b16565b611bf6565b81518015611bf3858263ffffffff61278b16565b50505b5060015b92915050565b6000611c0b83612780565b1580611c1d5750611c1b82612780565b155b15611c2a5750600061181b565b82518251808216611813878263ffffffff61278b16565b6000611c4c83612780565b1580611c5e5750611c5c82612780565b155b15611c6b5750600061181b565b82518251808217611813878263ffffffff61278b16565b6000611c8d83612780565b1580611c9f5750611c9d82612780565b155b15611cac5750600061181b565b82518251808218611813878263ffffffff61278b16565b6000611cce82612780565b611cda57506000611bfa565b81518019611cee858263ffffffff61278b16565b506001949350505050565b6000611d0483612780565b1580611d165750611d1482612780565b155b15611d235750600061181b565b8251825180821a611813878263ffffffff61278b16565b6000611bf6611d48836126cd565b51849063ffffffff61278b16565b600192915050565b6000611d778260800151836127c890919063ffffffff16565b506001919050565b6000611d778260600151836127c890919063ffffffff16565b6000611da3826126cd565b606084015250600192915050565b8051600090611dbf846126cd565b511415611e13576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b611bb0848363ffffffff6127c816565b6000611e2e826126cd565b835250600192915050565b6000611e4482612780565b611e505750600061181b565b825115611bb057611e60836126cd565b84525060019392505050565b6000611d77611e91611e84611e7f6127d6565b6126cd565b516020850151511461279f565b839063ffffffff611fe016565b805160009061138190839063ffffffff6127c816565b6000611bfa838363ffffffff611fbe16565b600080611ed1613171565b6000858581518110611edf57fe5b016020015160019095019460f81c9050600081611f2057611f008787612830565b9096509050600086611f118361240d565b91965094509250611fb7915050565b60ff821660021415611f4757611f368787612830565b9096509050600086611f1183612853565b600360ff831610801590611f5e5750600c60ff8316105b15611f9957600219820160606000611f77838b8b6128ad565b909a50925090508089611f8984612467565b9750975097505050505050611fb7565b8160ff16612710016000611fad600061240d565b9196509450925050505b9250925092565b611fd48260400151611fcf836126cd565b612962565b82604001819052505050565b611ff18260200151611fcf836126cd565b82602001819052505050565b600061200f838363ffffffff611fe016565b611bf6838363ffffffff611fe016565b6000612031848363ffffffff611fe016565b612041848463ffffffff611fe016565b611bb0848363ffffffff611fe016565b6000612063858363ffffffff611fe016565b612073858463ffffffff611fe016565b612083858563ffffffff611fe016565b611cee858363ffffffff611fe016565b6000612041848463ffffffff611fe016565b60006120b7858563ffffffff611fe016565b612083858463ffffffff611fe016565b60006120d283612a18565b15806120e457506120e282612780565b155b156120f15750600061181b565b6120fa83612a27565b60ff16826000015111156121105750600061181b565b611bb0836020015183600001518151811061212757fe5b602002602001015185611fe090919063ffffffff16565b600061214984612a18565b158061215b575061215983612780565b155b15612168575060006119f9565b61217184612a27565b60ff1683600001511115612187575060006119f9565b81846020015184600001518151811061219c57fe5b6020908102919091010152611cee858563ffffffff611fe016565b60006121c282612a18565b6121ce57506000611bfa565b611bf66121da83612a27565b849060ff1663ffffffff61278b16565b6000611bf66122006121fb84612a18565b61279f565b849063ffffffff611fe016565b600080612218613195565b612221846126cd565b516001969095509350505050565b600080600080600080600061224388612a18565b612257576000965094509092509050612404565b612278886020015160018151811061226b57fe5b6020026020010151612780565b61228c576000965094509092509050612404565b6122a0886020015160028151811061226b57fe5b6122b4576000965094509092509050612404565b6122c8886020015160038151811061226b57fe5b6122dc576000965094509092509050612404565b87602001516001815181106122ed57fe5b60200260200101516000015160001b9250876020015160028151811061230f57fe5b602002602001015160000151915073__$6b4cc75dad3e0abd6ad83b3d907747c608$__63ccf69dd76123408a6126cd565b6000015185858c6020015160038151811061235757fe5b60200260200101516000015160001b6040518563ffffffff1660e01b815260040180858152602001846affffffffffffffffffffff19166affffffffffffffffffffff1916815260200183815260200182815260200194505050505060206040518083038186803b1580156123cb57600080fd5b505af41580156123df573d6000803e3d6000fd5b505050506040513d60208110156123f557600080fd5b50516001975095509193509150505b92959194509250565b612415613171565b604080516060810182528381528151600080825260208281019094529192830191612456565b612443613171565b81526020019060019003908161243b5790505b508152600060209091015292915050565b61246f613171565b6124798251612a36565b61248257600080fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b600160c090910152565b600060028260c0015114156124c857506000611381565b60018260c0015114156124dd57506001611381565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e09093019092528151910120611381565b600060c090910152565b60008061255b613110565b612563613110565b600060c082018190526125768787612a3d565b845296509050801561258e5793508492509050611fb7565b6125988787612a3d565b60208501529650905080156125b35793508492509050611fb7565b6125bd8787612a3d565b60408501529650905080156125d85793508492509050611fb7565b6125e28787612a3d565b60608501529650905080156125fd5793508492509050611fb7565b6126078787612a3d565b60808501529650905080156126225793508492509050611fb7565b61262c8787612a3d565b60a08501529650905080156126475793508492509050611fb7565b506000969495509392505050565b61265d613110565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b60008060006126c48460ff16612a78565b50949350505050565b6126d5613195565b6040820151600c60ff909116106126eb57600080fd5b604082015160ff1661271857604051806020016040528061270f8460000151612f07565b90529050611381565b604082015160ff166002141561273d5750604080516020810190915281518152611381565b600360ff16826040015160ff161015801561276157506040820151600c60ff909116105b1561277e57604051806020016040528061270f8460200151612f2b565bfe5b6040015160ff161590565b611ff18260200151611fcf611e7f8461240d565b6127a7613171565b81156127be576127b7600161240d565b9050611381565b6127b7600061240d565b611ff1826020015182612962565b6127de613171565b60408051606081018252600080825282518181526020818101909452919283019190612820565b61280d613171565b8152602001906001900390816128055790505b5081526003602090910152905090565b60008080612844858563ffffffff61303516565b60209490940195939450505050565b61285b613171565b60408051606081018252838152815160008082526020828101909452919283019161289c565b612889613171565b8152602001906001900390816128815790505b508152600260209091015292915050565b6000806060600060608760ff166040519080825280602002602001820160405280156128f357816020015b6128e0613171565b8152602001906001900390816128d85790505b50905060005b8860ff168160ff16101561294d576129118888611ec6565b8451859060ff861690811061292257fe5b602090810291909101015297509250821561294557509093508492509050612959565b6001016128f9565b50600094508593509150505b93509350939050565b61296a613195565b6040805160028082526060828101909352816020015b612988613195565b81526020019060019003908161298057905050905082816000815181106129ab57fe5b602002602001018190525083816001815181106129c457fe5b60200260200101819052506040518060200160405280612a0e60405180604001604052806129f58860000151612853565b8152602001612a078960000151612853565b9052613051565b9052949350505050565b6000611bfa82604001516130cf565b6000611bfa82604001516130ed565b6008101590565b600080612a48613195565b6000612a5a868663ffffffff61303516565b60408051602080820190925291825260009896019650949350505050565b60008082612a8b57506000905080612f02565b6001831415612aa05750600290506001612f02565b6002831415612ab55750600290506001612f02565b6003831415612aca5750600290506001612f02565b6004831415612adf5750600290506001612f02565b6005831415612af45750600290506001612f02565b6006831415612b095750600290506001612f02565b6007831415612b1e5750600290506001612f02565b6008831415612b335750600390506001612f02565b6009831415612b485750600390506001612f02565b600a831415612b5d5750600290506001612f02565b600b831415612b725750600290506001612f02565b6010831415612b875750600290506001612f02565b6011831415612b9c5750600290506001612f02565b6012831415612bb15750600290506001612f02565b6013831415612bc65750600290506001612f02565b6014831415612bdb5750600290506001612f02565b6015831415612bef57506001905080612f02565b6016831415612c045750600290506001612f02565b6017831415612c195750600290506001612f02565b6018831415612c2e5750600290506001612f02565b6019831415612c4257506001905080612f02565b601a831415612c575750600290506001612f02565b6020831415612c6b57506001905080612f02565b6030831415612c805750600190506000612f02565b6031831415612c955750600090506001612f02565b6032831415612caa5750600090506001612f02565b6033831415612cbf5750600190506000612f02565b6034831415612cd357506001905080612f02565b6035831415612ce85750600190506000612f02565b6036831415612cfd5750600290506000612f02565b6037831415612d125750600090506001612f02565b6038831415612d275750600090506001612f02565b6039831415612d3c5750600190506000612f02565b603a831415612d515750600090506001612f02565b603b831415612d665750600090506001612f02565b603c831415612d7a57506000905080612f02565b603d831415612d8f5750600090506001612f02565b603e831415612da45750600190506000612f02565b6040831415612db95750600190506002612f02565b6041831415612dce5750600290506003612f02565b6042831415612de35750600390506004612f02565b6043831415612df757506002905080612f02565b6044831415612e0b57506003905080612f02565b6050831415612e205750600290506001612f02565b6051831415612e355750600390506001612f02565b6052831415612e4957506001905080612f02565b6053831415612e5d57506001905080612f02565b6060831415612e7157506000905080612f02565b6061831415612e865750600190506000612f02565b6070831415612e9b5750600190506000612f02565b6070831415612eaf57506001905080612f02565b6073831415612ec45750600090506001612f02565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600882511115612f3c57600080fd5b60608251604051908082528060200260200182016040528015612f69578160200160208202803883390190505b50905060005b8151811015612fc357612f80613195565b612f9c858381518110612f8f57fe5b60200260200101516126cd565b90508060000151838381518110612faf57fe5b602090810291909101015250600101612f6f565b508251600360ff160181604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561300c578181015183820152602001612ff4565b505050509050019250505060405160208183030381529060405280519060200120915050919050565b6000816020018351101561304857600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b613074613171565b81526020019060019003908161306c57905050905060005b81518110156130c5578381600281106130a157fe5b60200201518282815181106130b257fe5b602090810291909101015260010161308c565b5061181b81612f2b565b6000600c60ff8316108015611bfa575050600360ff91909116101590565b60006130f8826130cf565b1561310857506002198101611381565b506001611381565b6040518060e00160405280613123613195565b8152602001613130613195565b815260200161313d613195565b815260200161314a613195565b8152602001613157613195565b8152602001613164613195565b8152602001600081525090565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fe50726f6f6620686164206e6f6e206d61746368696e672073746172742073746174654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f53656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f50726f6f66206f6620617578706f702068616420626164206175782076616c7565a265627a7a72305820c712e1a72994fab79a85402e1bc3b091ed4381fb2004a580ad2617c7be05fe7d64736f6c63430005090032`

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

// ValidateProof is a free data retrieval call binding the contract method 0x0eca9f13.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] beforeValues, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, beforeValues []*big.Int, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	ret0 := new(*big.Int)

	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", fields, timeBounds, tokenTypes, beforeValues, messageValue, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0x0eca9f13.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] beforeValues, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, beforeValues []*big.Int, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, tokenTypes, beforeValues, messageValue, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0x0eca9f13.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes21[] tokenTypes, uint256[] beforeValues, uint256[] messageValue, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, tokenTypes [][21]byte, beforeValues []*big.Int, messageValue []*big.Int, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, tokenTypes, beforeValues, messageValue, proof)
}

// OneStepProofSawMachineIterator is returned from FilterSawMachine and is used to iterate over the raw logs and unpacked data for SawMachine events raised by the OneStepProof contract.
type OneStepProofSawMachineIterator struct {
	Event *OneStepProofSawMachine // Event containing the contract specifics and raw log

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
func (it *OneStepProofSawMachineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneStepProofSawMachine)
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
		it.Event = new(OneStepProofSawMachine)
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
func (it *OneStepProofSawMachineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneStepProofSawMachineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneStepProofSawMachine represents a SawMachine event raised by the OneStepProof contract.
type OneStepProofSawMachine struct {
	InstructionStack [32]byte
	DataStack        [32]byte
	AuxStack         [32]byte
	Register         [32]byte
	StaticHash       [32]byte
	ErrHandler       [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSawMachine is a free log retrieval operation binding the contract event 0x10d11f456a57c1ced446abc92b6cfa3854cd21f069b29f283252b23223d03080.
//
// Solidity: event SawMachine(bytes32 instructionStack, bytes32 dataStack, bytes32 auxStack, bytes32 register, bytes32 staticHash, bytes32 errHandler)
func (_OneStepProof *OneStepProofFilterer) FilterSawMachine(opts *bind.FilterOpts) (*OneStepProofSawMachineIterator, error) {
	logs, sub, err := _OneStepProof.contract.FilterLogs(opts, "SawMachine")
	if err != nil {
		return nil, err
	}
	return &OneStepProofSawMachineIterator{contract: _OneStepProof.contract, event: "SawMachine", logs: logs, sub: sub}, nil
}

// WatchSawMachine is a free log subscription operation binding the contract event 0x10d11f456a57c1ced446abc92b6cfa3854cd21f069b29f283252b23223d03080.
//
// Solidity: event SawMachine(bytes32 instructionStack, bytes32 dataStack, bytes32 auxStack, bytes32 register, bytes32 staticHash, bytes32 errHandler)
func (_OneStepProof *OneStepProofFilterer) WatchSawMachine(opts *bind.WatchOpts, sink chan<- *OneStepProofSawMachine) (event.Subscription, error) {
	logs, sub, err := _OneStepProof.contract.WatchLogs(opts, "SawMachine")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneStepProofSawMachine)
				if err := _OneStepProof.contract.UnpackLog(event, "SawMachine", log); err != nil {
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
