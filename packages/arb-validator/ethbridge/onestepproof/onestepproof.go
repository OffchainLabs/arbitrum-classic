// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onestepproof

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
const ArbMachineABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"instructionStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"errHandlerHash\",\"type\":\"bytes32\"}],\"name\":\"machineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbMachineFuncSigs maps the 4-byte function signature to its string representation.
var ArbMachineFuncSigs = map[string]string{
	"c1355b59": "machineHash(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32)",
}

// ArbMachineBin is the compiled bytecode used for deploying new contracts.
var ArbMachineBin = "0x6101d6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063c1355b591461003a575b600080fd5b610075600480360360c081101561005057600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610087565b60408051918252519081900360200190f35b604080516101008101825260e081018881528152815160208181018452888252808301919091528251808201845287815282840152825180820184528681526060830152825180820184528581526080830152825190810190925282825260a0810191909152600060c08201819052906101009061010b565b979650505050505050565b600060028260c0015114156101225750600061019c565b60018260c0015114156101375750600161019c565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101205b91905056fea265627a7a723158209ea513e512ebbb00bc77336c9424b4a8c3b3691610fe0a7b62b62d107145e9f464736f6c634300050c0032"

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
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastLogHash\",\"type\":\"bytes32\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint16[]\",\"name\":\"_tokenNums\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes21\",\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[2]\",\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes32\",\"name\":\"_beforeInbox\",\"type\":\"bytes32\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"0f89fbff": "calculateBeforeValues(bytes21[],uint16[],uint256[])",
	"7ddf59d6": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32)",
	"1914612a": "generateLastMessageHash(bytes21[],bytes,uint16[],uint256[],address[])",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"85ecb92a": "generatePreconditionHash(bytes32,uint64[2],bytes32)",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x61115a610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100605760003560e01c80624c28f6146100655780630f89fbff146100bd5780631914612a146102b25780637ddf59d61461055e57806385ecb92a1461059f575b600080fd5b6100ab6004803603608081101561007b57600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b03166105f4565b60408051918252519081900360200190f35b610262600480360360608110156100d357600080fd5b810190602081018135600160201b8111156100ed57600080fd5b8201836020820111156100ff57600080fd5b803590602001918460208302840111600160201b8311171561012057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460208302840111600160201b831117156101a257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101f157600080fd5b82018360208201111561020357600080fd5b803590602001918460208302840111600160201b8311171561022457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506106e6945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561029e578181015183820152602001610286565b505050509050019250505060405180910390f35b6100ab600480360360a08110156102c857600080fd5b810190602081018135600160201b8111156102e257600080fd5b8201836020820111156102f457600080fd5b803590602001918460208302840111600160201b8311171561031557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561036457600080fd5b82018360208201111561037657600080fd5b803590602001918460018302840111600160201b8311171561039757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103e957600080fd5b8201836020820111156103fb57600080fd5b803590602001918460208302840111600160201b8311171561041c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561046b57600080fd5b82018360208201111561047d57600080fd5b803590602001918460208302840111600160201b8311171561049e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104ed57600080fd5b8201836020820111156104ff57600080fd5b803590602001918460208302840111600160201b8311171561052057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108cb945050505050565b6100ab600480360360c081101561057457600080fd5b5080359063ffffffff6020820135169060408101359060608101359060808101359060a00135610b0f565b6100ab600480360360808110156105b557600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194505090359150610b679050565b60408051600480825260a0820190925260009160609190816020015b6106186110be565b81526020019060019003908161061057905050905061063686610bbb565b8160008151811061064357fe5b6020026020010181905250610660836001600160a01b0316610c3b565b8160018151811061066d57fe5b602002602001018190525061068184610c3b565b8160028151811061068e57fe5b60209081029190910101526106b06affffffffffffffffffffff198616610c3b565b816003815181106106bd57fe5b60200260200101819052506106d96106d482610cb9565b610d69565b519150505b949350505050565b60606000835190506060855160405190808252806020026020018201604052801561071b578160200160208202803883390190505b50905060005b828110156108c157600086828151811061073757fe5b60200260200101519050878161ffff168151811061075157fe5b602002602001015160146015811061076557fe5b1a60f81b6001600160f81b0319166107b25785828151811061078357fe5b6020026020010151838261ffff168151811061079b57fe5b6020026020010181815101915081815250506108b8565b828161ffff16815181106107c257fe5b602002602001015160001461081e576040805162461bcd60e51b815260206004820152601d60248201527f43616e277420696e636c756465204e465420746f6b656e207477696365000000604482015290519081900360640190fd5b85828151811061082a57fe5b602002602001015160001415610887576040805162461bcd60e51b815260206004820152601f60248201527f4e465420746f6b656e206d7573742068617665206e6f6e2d7a65726f20696400604482015290519081900360640190fd5b85828151811061089357fe5b6020026020010151838261ffff16815181106108ab57fe5b6020026020010181815250505b50600101610721565b5095945050505050565b60008151835114610919576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8351835114610965576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b825160009081908190815b81811015610b005773__$d969135829891f807aa9c34494da4ecd99$__6389df40da8b866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b838110156109e65781810151838201526020016109ce565b50505050905090810190601f168015610a135780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b158015610a3057600080fd5b505af4158015610a44573d6000803e3d6000fd5b505050506040513d6040811015610a5a57600080fd5b5080516020909101518a519195509350610ac99084908d908c9085908110610a7e57fe5b602002602001015161ffff1681518110610a9457fe5b60200260200101518a8481518110610aa857fe5b60200260200101518a8581518110610abc57fe5b60200260200101516105f4565b6040805160208082019890985280820183905281518082038301815260609091019091528051960195909520949250600101610970565b50929998505050505050505050565b6040805160208082019890985260e09690961b6001600160e01b0319168682015260448601949094526064850192909252608484015260a4808401919091528151808403909101815260c49092019052805191012090565b815160209283015160408051808601969096526001600160c01b031960c093841b8116878301529190921b166048850152605080850192909252805180850390920182526070909301909252815191012090565b610bc36110be565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610c28565b610c156110be565b815260200190600190039081610c0d5790505b508152600260209091015290505b919050565b610c436110be565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610ca8565b610c956110be565b815260200190600190039081610c8d5790505b508152600060209091015292915050565b610cc16110be565b610ccb8251610e9f565b610d1c576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b610d716110ec565b6060820151600c60ff90911610610dc3576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610df0576040518060200160405280610de78460000151610ea6565b90529050610c36565b606082015160ff1660011415610e37576040518060200160405280610de7846020015160000151856020015160400151866020015160600151876020015160200151610eca565b606082015160ff1660021415610e5c5750604080516020810190915281518152610c36565b600360ff16826060015160ff1610158015610e8057506060820151600c60ff909116105b15610e9d576040518060200160405280610de78460400151610f72565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315610f24575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206106de565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6000600882511115610fc2576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610fef578160200160208202803883390190505b50805190915060005b8181101561104b576110086110ec565b61102486838151811061101757fe5b6020026020010151610d69565b9050806000015184838151811061103757fe5b602090810291909101015250600101610ff8565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561109457818101518382015260200161107c565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6040518060800160405280600081526020016110d86110fe565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a7231582040553ef7d22f8e9eef54bca2640e42b7085679542974774647c1dc8cfdfdd88364736f6c634300050c0032"

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

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCaller) CalculateBeforeValues(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
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

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x1914612a.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHash", _tokenTypes, _data, _tokenNums, _amounts, _destinations)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x1914612a.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _data, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x1914612a.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _data, _tokenNums, _amounts, _destinations)
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
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidValueHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeValueHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValidValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"immediate\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueFuncSigs maps the 4-byte function signature to its string representation.
var ArbValueFuncSigs = map[string]string{
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
var ArbValueBin = "0x61111c610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061009d5760003560e01c806353409fab1161007057806353409fab1461022c578063826513e01461025257806389df40da146102865780638f34603614610347578063b2b9dc62146103ed5761009d565b80631667b411146100a25780631f3d4d4e146100d1578063264f384b146101f8578063364df27714610224575b600080fd5b6100bf600480360360208110156100b857600080fd5b503561041e565b60408051918252519081900360200190f35b610179600480360360408110156100e757600080fd5b81019060208101813564010000000081111561010257600080fd5b82018360208201111561011457600080fd5b8035906020019184600183028401116401000000008311171561013657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610444915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101bc5781810151838201526020016101a4565b50505050905090810190601f1680156101e95780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100bf6004803603606081101561020e57600080fd5b5060ff81351690602081013590604001356104c8565b6100bf61051a565b6100bf6004803603604081101561024257600080fd5b5060ff813516906020013561058d565b6100bf6004803603608081101561026857600080fd5b5060ff813516906020810135151590604081013590606001356105d4565b61032e6004803603604081101561029c57600080fd5b8101906020810181356401000000008111156102b757600080fd5b8201836020820111156102c957600080fd5b803590602001918460018302840111640100000000831117156102eb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061067d915050565b6040805192835260208301919091528051918290030190f35b6100bf6004803603602081101561035d57600080fd5b81019060208101813564010000000081111561037857600080fd5b82018360208201111561038a57600080fd5b803590602001918460018302840111640100000000831117156103ac57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106f6945050505050565b61040a6004803603602081101561040357600080fd5b5035610768565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b60006060600080610453611060565b61045d878761076f565b9194509250905082156104a5576040805162461bcd60e51b815260206004820152601e60248201526000805160206110c8833981519152604482015290519081900360640190fd5b816104b9888880840363ffffffff6108f916565b945094505050505b9250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b8381101561056657818101518382015260200161054e565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000831561062e575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120610675565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b166021830152602280830185905283518084039091018152604290920190925280519101205b949350505050565b60008060008061068b611060565b610695878761076f565b9194509250905082156106dd576040805162461bcd60e51b815260206004820152601e60248201526000805160206110c8833981519152604482015290519081900360640190fd5b816106e782610979565b51909890975095505050505050565b60008080610702611060565b61070d85600061076f565b919450925090508215610755576040805162461bcd60e51b815260206004820152601e60248201526000805160206110c8833981519152604482015290519081900360640190fd5b61075e81610979565b5195945050505050565b6008101590565b60008061077a611060565b845184106107cf576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b600084905060008682815181106107e257fe5b016020015160019092019160f81c905060006107fc61108e565b60ff83166108305761080e8985610aaf565b909450915060008461081f84610ad6565b919850965094506108f29350505050565b60ff831660011415610857576108468985610b54565b909450905060008461081f83610caf565b60ff83166002141561087e5761086d8985610aaf565b909450915060008461081f84610d0f565b600360ff8416108015906108955750600c60ff8416105b156108d2576002198301606060006108ae838d89610d8d565b9098509250905080876108c084610e48565b995099509950505050505050506108f2565b8260ff166127100160006108e66000610ad6565b91985096509450505050505b9250925092565b60608183018451101561090b57600080fd5b60608215801561092657604051915060208201604052610970565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561095f578051835260209283019201610947565b5050858452601f01601f1916604052505b50949350505050565b6109816110b5565b6060820151600c60ff909116106109d3576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16610a005760405180602001604052806109f7846000015161041e565b9052905061043f565b606082015160ff1660011415610a475760405180602001604052806109f78460200151600001518560200151604001518660200151606001518760200151602001516105d4565b606082015160ff1660021415610a6c575060408051602081019091528151815261043f565b600360ff16826060015160ff1610158015610a9057506060820151600c60ff909116105b15610aad5760405180602001604052806109f78460400151610ef8565bfe5b6000808281610ac4868363ffffffff61104416565b60209290920196919550909350505050565b610ade611060565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610b43565b610b30611060565b815260200190600190039081610b285790505b508152600060209091015292915050565b6000610b5e61108e565b60008390506000858281518110610b7157fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110610b9757fe5b016020015160019384019360f89190911c915060009060ff84161415610c23576000610bc1611060565b610bcb8a8761076f565b90975090925090508115610c14576040805162461bcd60e51b815260206004820152601e60248201526000805160206110c8833981519152604482015290519081900360640190fd5b610c1d81610979565b51925050505b6000610c35898663ffffffff61104416565b90506020850194508360ff1660011415610c7a576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506104c19050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b610cb7611060565b604080516080810182526000808252602080830186905283518281529081018452919283019190610cfe565b610ceb611060565b815260200190600190039081610ce35790505b508152600160209091015292915050565b610d17611060565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610d7c565b610d69611060565b815260200190600190039081610d615790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610dd857816020015b610dc5611060565b815260200190600190039081610dbd5790505b50905060005b8960ff168160ff161015610e3257610df6898561076f565b8451859060ff8616908110610e0757fe5b6020908102919091010152945092508215610e2a57509094509092509050610e3f565b600101610dde565b5060009550919350909150505b93509350939050565b610e50611060565b610e5a8251610768565b610eab576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b6000600882511115610f48576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610f75578160200160208202803883390190505b50805190915060005b81811015610fd157610f8e6110b5565b610faa868381518110610f9d57fe5b6020026020010151610979565b90508060000151848381518110610fbd57fe5b602090810291909101015250600101610f7e565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561101a578181015183820152602001611002565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6000816020018351101561105757600080fd5b50016020015190565b60405180608001604052806000815260200161107a61108e565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4d61727368616c6c65642076616c7565206d7573742062652076616c69640000a265627a7a7231582070911ce47e4394d9bdb026b7e0433734675bf729f24585fcfd5404a7cd06318c64736f6c634300050c0032"

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
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a9fb49964eecf6140d019ab385582e49921deabebb21f44e37efb5be3dbaaf2b64736f6c634300050c0032"

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
var ChallengeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582015ff87a6136e0af45561bb518db8ecee50aa8be3e12ed77c747784eb8c95f76b64736f6c634300050c0032"

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

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"b32\",\"type\":\"bytes32\"}],\"name\":\"bytes32string\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"out\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// DebugPrintFuncSigs maps the 4-byte function signature to its string representation.
var DebugPrintFuncSigs = map[string]string{
	"252fb38d": "bytes32string(bytes32)",
}

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x610202610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063252fb38d1461003a575b600080fd5b6100576004803603602081101561005057600080fd5b50356100cc565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610091578181015183820152602001610079565b50505050905090810190601f1680156100be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051818152606081810183529182919060208201818038833901905050905060005b602081101561019457600084826020811061010757fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b61012d8261019b565b85856002028151811061013c57fe5b60200101906001600160f81b031916908160001a90535061015c8161019b565b85856002026001018151811061016e57fe5b60200101906001600160f81b031916908160001a90535050600190920191506100f09050565b5092915050565b6000600a60f883901c10156101bb578160f81c60300160f81b90506101c8565b8160f81c60570160f81b90505b91905056fea265627a7a723158202bec80a36bc31e4b6cded2c1a61da54ef42647d739c58b7139ae3670fcf1c46364736f6c634300050c0032"

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

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[7]\",\"name\":\"fields\",\"type\":\"bytes32[7]\"},{\"internalType\":\"uint64[2]\",\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"a49c3308": "oneStepProof(Challenge.Data storage,bytes32[2],uint64[2],bytes32[5],bytes)",
	"c0fee45d": "validateProof(bytes32[7],uint64[2],bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x613c5c610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063a49c330814610045578063c0fee45d14610177575b600080fd5b610175600480360361016081101561005c57600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152505060408051808201825292959493818101939250906002908390839080828437600092019190915250506040805160a0818101909252929594938181019392509060059083908390808284376000920191909152509194939260208101925035905064010000000081111561010057600080fd5b82018360208201111561011257600080fd5b8035906020019184600183028401116401000000008311171561013457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061028c945050505050565b005b61027a600480360361014081101561018e57600080fd5b810190808060e00190600780602002604051908101604052809291908260076020028082843760009201919091525050604080518082018252929594938181019392509060029083908390808284376000920191909152509194939260208101925035905064010000000081111561020557600080fd5b82018360208201111561021757600080fd5b8035906020019184600183028401116401000000008311171561023957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610627945050505050565b60408051918252519081900360200190f35b60016005860154600160601b900460ff1660028111156102a857fe5b146102e45760405162461bcd60e51b8152600401808060200182810382526039815260200180613bef6039913960400191505060405180910390fd5b600585015467ffffffffffffffff16431115610347576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6520737465702070726f6f66206d697373656420646561646c696e650000604482015290519081900360640190fd5b600185015484516020860151604080516342f65c9560e11b81526004810184815273__$9836fa7140e5a33041d4b827682e675a30$__946385ecb92a9490938a9391929160240190849080838360005b838110156103af578181015183820152602001610397565b50505050905001828152602001935050505060206040518083038186803b1580156103d957600080fd5b505af41580156103ed573d6000803e3d6000fd5b505050506040513d602081101561040357600080fd5b50518351602080860151604080880151606089015160808a01518351633eefaceb60e11b815260048101979097526001602488015260448701949094526064860191909152608485015260a48401919091525173__$9836fa7140e5a33041d4b827682e675a30$__92637ddf59d69260c4808301939192829003018186803b15801561048e57600080fd5b505af41580156104a2573d6000803e3d6000fd5b505050506040513d60208110156104b857600080fd5b505160408051602081810194909452808201929092528051808303820181526060909201905280519101201461051f5760405162461bcd60e51b8152600401808060200182810382526026815260200180613acb6026913960400191505060405180910390fd5b60006105d46040518060e001604052808760006002811061053c57fe5b602002015181526020018760016002811061055357fe5b602002015181526020018560006005811061056a57fe5b602002015181526020018560016005811061058157fe5b602002015181526020018560026005811061059857fe5b60200201518152602001856003600581106105af57fe5b60200201518152602001856004600581106105c657fe5b602002015190528584610627565b9050801561061f576040805162461bcd60e51b8152602060048201526013602482015272141c9bdbd9881dd85cc81a5b98dbdc9c9958dd606a1b604482015290519081900360640190fd5b505050505050565b60006106e86040518061012001604052808660006007811061064557fe5b602002015181526020018581526020018660016007811061066257fe5b602002015181526020018660026007811061067957fe5b602002015181526020018660036007811061069057fe5b60200201518152602001866004600781106106a757fe5b60200201518152602001866005600781106106be57fe5b60200201518152602001866006600781106106d557fe5b60200201518152602001848152506106f2565b90505b9392505050565b60008080806060610701613a02565b610709613a02565b610712886115f2565b939950929650909450925090506001600060ff881682141561076857610761838660008151811061073f57fe5b60200260200101518760018151811061075457fe5b6020026020010151611a32565b9150611446565b60ff8816600214156107a757610761838660008151811061078557fe5b60200260200101518760018151811061079a57fe5b6020026020010151611a80565b60ff8816600314156107e65761076183866000815181106107c457fe5b6020026020010151876001815181106107d957fe5b6020026020010151611ac1565b60ff88166004141561082557610761838660008151811061080357fe5b60200260200101518760018151811061081857fe5b6020026020010151611b02565b60ff88166005141561086457610761838660008151811061084257fe5b60200260200101518760018151811061085757fe5b6020026020010151611b53565b60ff8816600614156108a357610761838660008151811061088157fe5b60200260200101518760018151811061089657fe5b6020026020010151611ba4565b60ff8816600714156108e25761076183866000815181106108c057fe5b6020026020010151876001815181106108d557fe5b6020026020010151611bf5565b60ff8816600814156109365761076183866000815181106108ff57fe5b60200260200101518760018151811061091457fe5b60200260200101518860028151811061092957fe5b6020026020010151611c46565b60ff88166009141561098a57610761838660008151811061095357fe5b60200260200101518760018151811061096857fe5b60200260200101518860028151811061097d57fe5b6020026020010151611cb0565b60ff8816600a14156109c95761076183866000815181106109a757fe5b6020026020010151876001815181106109bc57fe5b6020026020010151611d09565b60ff881660101415610a085761076183866000815181106109e657fe5b6020026020010151876001815181106109fb57fe5b6020026020010151611d4a565b60ff881660111415610a47576107618386600081518110610a2557fe5b602002602001015187600181518110610a3a57fe5b6020026020010151611d8b565b60ff881660121415610a86576107618386600081518110610a6457fe5b602002602001015187600181518110610a7957fe5b6020026020010151611dcc565b60ff881660131415610ac5576107618386600081518110610aa357fe5b602002602001015187600181518110610ab857fe5b6020026020010151611e0d565b60ff881660141415610b04576107618386600081518110610ae257fe5b602002602001015187600181518110610af757fe5b6020026020010151611e4e565b60ff881660151415610b2e576107618386600081518110610b2157fe5b6020026020010151611e7a565b60ff881660161415610b6d576107618386600081518110610b4b57fe5b602002602001015187600181518110610b6057fe5b6020026020010151611ec0565b60ff881660171415610bac576107618386600081518110610b8a57fe5b602002602001015187600181518110610b9f57fe5b6020026020010151611f01565b60ff881660181415610beb576107618386600081518110610bc957fe5b602002602001015187600181518110610bde57fe5b6020026020010151611f42565b60ff881660191415610c15576107618386600081518110610c0857fe5b6020026020010151611f83565b60ff8816601a1415610c54576107618386600081518110610c3257fe5b602002602001015187600181518110610c4757fe5b6020026020010151611fb9565b60ff8816601b1415610c93576107618386600081518110610c7157fe5b602002602001015187600181518110610c8657fe5b6020026020010151611ffa565b60ff881660201415610cbd576107618386600081518110610cb057fe5b602002602001015161203b565b60ff881660211415610ce7576107618386600081518110610cda57fe5b6020026020010151612057565b60ff881660301415610d11576107618386600081518110610d0457fe5b6020026020010151612072565b60ff881660311415610d26576107618361207a565b60ff881660321415610d3b576107618361209b565b60ff881660331415610d65576107618386600081518110610d5857fe5b60200260200101516120b4565b60ff881660341415610d8f576107618386600081518110610d8257fe5b60200260200101516120cd565b60ff881660351415610dce576107618386600081518110610dac57fe5b602002602001015187600181518110610dc157fe5b60200260200101516120e3565b60ff881660361415610de3576107618361212b565b60ff881660371415610dfd5761076183856000015161215d565b60ff881660381415610e27576107618386600081518110610e1a57fe5b602002602001015161216f565b60ff881660391415610eb457610e3b613a63565b610e4a8b610100015188612181565b919950975090508715610e8e5760405162461bcd60e51b8152600401808060200182810382526021815260200180613bce6021913960400191505060405180910390fd5b610e9e858263ffffffff61230b16565b610eae848263ffffffff61232d16565b50611446565b60ff8816603a1415610ec9576107618361234a565b60ff8816603b1415610eda57611446565b60ff8816603c1415610eef576107618361236a565b60ff8816603d1415610f19576107618386600081518110610f0c57fe5b6020026020010151612383565b60ff881660401415610f43576107618386600081518110610f3657fe5b60200260200101516123b1565b60ff881660411415610f82576107618386600081518110610f6057fe5b602002602001015187600181518110610f7557fe5b60200260200101516123d3565b60ff881660421415610fd6576107618386600081518110610f9f57fe5b602002602001015187600181518110610fb457fe5b602002602001015188600281518110610fc957fe5b6020026020010151612405565b60ff881660431415611015576107618386600081518110610ff357fe5b60200260200101518760018151811061100857fe5b6020026020010151612447565b60ff88166044141561106957610761838660008151811061103257fe5b60200260200101518760018151811061104757fe5b60200260200101518860028151811061105c57fe5b6020026020010151612459565b60ff8816605014156110a857610761838660008151811061108657fe5b60200260200101518760018151811061109b57fe5b602002602001015161247b565b60ff8816605114156110fc5761076183866000815181106110c557fe5b6020026020010151876001815181106110da57fe5b6020026020010151886002815181106110ef57fe5b60200260200101516124f1565b60ff88166052141561112657610761838660008151811061111957fe5b6020026020010151612569565b60ff88166060141561113b576107618361259c565b60ff88166061141561122557611165838660008151811061115857fe5b60200260200101516125a2565b60e08c015160c08d01516040805160208082019390935280820185905281518082038301815260609091019091528051910120929450909250146111da5760405162461bcd60e51b8152600401808060200182810382526025815260200180613b5a6025913960400191505060405180910390fd5b8960a001518a60800151146112205760405162461bcd60e51b8152600401808060200182810382526027815260200180613b7f6027913960400191505060405180910390fd5b611446565b60ff88166070141561132357600080611252858860008151811061124557fe5b60200260200101516125c6565b809450819550829650839750505050508b60a001518c6080015184604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146112d65760405162461bcd60e51b8152600401808060200182810382526028815260200180613ba66028913960400191505060405180910390fd5b8b60e001518c60c001511461131c5760405162461bcd60e51b8152600401808060200182810382526026815260200180613b136026913960400191505060405180910390fd5b5050611446565b60ff8816607214156113df576040805160028082526060828101909352816020015b61134d613a63565b81526020019060019003908161134557505060208c01519091506113829060005b602002015167ffffffffffffffff166127b2565b8160008151811061138f57fe5b60200260200101819052506113ae8b6020015160016002811061136e57fe5b816001815181106113bb57fe5b6020026020010181905250610eae6113d282612830565b859063ffffffff61232d16565b60ff88166073141561141c5761076183866000815181106113fc57fe5b602002602001015160405180602001604052808e604001518152506128e0565b60ff8816607414156114315760009150611446565b60ff8816607514156114465761144683612952565b806114d7578960a001518a60800151146114915760405162461bcd60e51b8152600401808060200182810382526027815260200180613b7f6027913960400191505060405180910390fd5b8960e001518a60c00151146114d75760405162461bcd60e51b8152600401808060200182810382526026815260200180613b136026913960400191505060405180910390fd5b816115395760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a08401515114156115315761152c8361295c565b611539565b60a083015183525b61154284612966565b8a51146115805760405162461bcd60e51b8152600401808060200182810382526022815260200180613af16022913960400191505060405180910390fd5b61158983612966565b8a60600151146115e0576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b600060606115fe613a02565b611606613a02565b60008080611612613a02565b61161b816129fb565b61162a89610100015184612a05565b909450909250905061163a613a02565b61164382612b0a565b905060008a6101000151858151811061165857fe5b602001015160f81c60f81b60f81c905060008b6101000151866001018151811061167e57fe5b016020015160f81c9050600061169382612b68565b90506060816040519080825280602002602001820160405280156116d157816020015b6116be613a63565b8152602001906001900390816116b65790505b5090506002880197508360ff16600014806116ef57508360ff166001145b611740576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff84166117e3576040805160208082018084528951516353409fab60e01b90915260ff87166024840152604483015291519091829173__$d969135829891f807aa9c34494da4ecd99$__916353409fab916064808601929190818703018186803b1580156117ae57600080fd5b505af41580156117c2573d6000803e3d6000fd5b505050506040513d60208110156117d857600080fd5b50519052865261193a565b6117eb613a63565b6117fa8f61010001518a612181565b909a5090985090508715611855576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561187957808260008151811061186957fe5b6020026020010181905250611889565b611889868263ffffffff61232d16565b604051806020016040528073__$d969135829891f807aa9c34494da4ecd99$__63264f384b876118b886612b82565b518c5151604080516001600160e01b031960e087901b16815260ff909416600485015260248401929092526044830152516064808301926020929190829003018186803b15801561190857600080fd5b505af415801561191c573d6000803e3d6000fd5b505050506040513d602081101561193257600080fd5b505190528752505b60ff84165b828110156119ce576119568f61010001518a612181565b845185908590811061196457fe5b60209081029190910101529950975087156119c6576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b60010161193f565b815115611a1b575060005b8460ff16825103811015611a1b57611a138282600185510303815181106119fc57fe5b60200260200101518861232d90919063ffffffff16565b6001016119d9565b50919d919c50939a50919850939650945050505050565b6000611a3d83612cb8565b1580611a4f5750611a4d82612cb8565b155b15611a5c575060006106eb565b82518251808201611a73878263ffffffff612cc316565b5060019695505050505050565b6000611a8b83612cb8565b1580611a9d5750611a9b82612cb8565b155b15611aaa575060006106eb565b82518251808202611a73878263ffffffff612cc316565b6000611acc83612cb8565b1580611ade5750611adc82612cb8565b155b15611aeb575060006106eb565b82518251808203611a73878263ffffffff612cc316565b6000611b0d83612cb8565b1580611b1f5750611b1d82612cb8565b155b15611b2c575060006106eb565b8251825180611b40576000925050506106eb565b808204611a73878263ffffffff612cc316565b6000611b5e83612cb8565b1580611b705750611b6e82612cb8565b155b15611b7d575060006106eb565b8251825180611b91576000925050506106eb565b808205611a73878263ffffffff612cc316565b6000611baf83612cb8565b1580611bc15750611bbf82612cb8565b155b15611bce575060006106eb565b8251825180611be2576000925050506106eb565b808206611a73878263ffffffff612cc316565b6000611c0083612cb8565b1580611c125750611c1082612cb8565b155b15611c1f575060006106eb565b8251825180611c33576000925050506106eb565b808207611a73878263ffffffff612cc316565b6000611c5184612cb8565b1580611c635750611c6183612cb8565b155b15611c7057506000611ca8565b83518351835180611c875760009350505050611ca8565b6000818385089050611c9f898263ffffffff612cc316565b60019450505050505b949350505050565b6000611cbb84612cb8565b1580611ccd5750611ccb83612cb8565b155b15611cda57506000611ca8565b83518351835180611cf15760009350505050611ca8565b6000818385099050611c9f898263ffffffff612cc316565b6000611d1483612cb8565b1580611d265750611d2482612cb8565b155b15611d33575060006106eb565b8251825180820a611a73878263ffffffff612cc316565b6000611d5583612cb8565b1580611d675750611d6582612cb8565b155b15611d74575060006106eb565b82518251808210611a73878263ffffffff612cc316565b6000611d9683612cb8565b1580611da85750611da682612cb8565b155b15611db5575060006106eb565b82518251808211611a73878263ffffffff612cc316565b6000611dd783612cb8565b1580611de95750611de782612cb8565b155b15611df6575060006106eb565b82518251808212611a73878263ffffffff612cc316565b6000611e1883612cb8565b1580611e2a5750611e2882612cb8565b155b15611e37575060006106eb565b82518251808213611a73878263ffffffff612cc316565b6000611e706113d2611e5f84612b82565b51611e6986612b82565b5114612cd7565b5060019392505050565b6000611e8582612cb8565b611e9f57611e9a83600063ffffffff612cc316565b611eb6565b81518015611eb3858263ffffffff612cc316565b50505b5060015b92915050565b6000611ecb83612cb8565b1580611edd5750611edb82612cb8565b155b15611eea575060006106eb565b82518251808216611a73878263ffffffff612cc316565b6000611f0c83612cb8565b1580611f1e5750611f1c82612cb8565b155b15611f2b575060006106eb565b82518251808217611a73878263ffffffff612cc316565b6000611f4d83612cb8565b1580611f5f5750611f5d82612cb8565b155b15611f6c575060006106eb565b82518251808218611a73878263ffffffff612cc316565b6000611f8e82612cb8565b611f9a57506000611eba565b81518019611fae858263ffffffff612cc316565b506001949350505050565b6000611fc483612cb8565b1580611fd65750611fd482612cb8565b155b15611fe3575060006106eb565b8251825181811a611a73878263ffffffff612cc316565b600061200583612cb8565b1580612017575061201582612cb8565b155b15612024575060006106eb565b8251825181810b611a73878263ffffffff612cc316565b6000611eb661204983612b82565b51849063ffffffff612cc316565b6000611eb661206583612d00565b849063ffffffff61232d16565b600192915050565b6000612093826080015183612d8990919063ffffffff16565b506001919050565b6000612093826060015183612d8990919063ffffffff16565b60006120bf82612b82565b606084015250600192915050565b60006120d882612b82565b835250600192915050565b60006120ee83612d97565b6120fa575060006106eb565b61210382612cb8565b61210f575060006106eb565b815115611e705761211f83612b82565b84525060019392505050565b600061209361215061214361213e612da4565b612b82565b5160208501515114612cd7565b839063ffffffff61232d16565b6000611eb6838363ffffffff612d8916565b6000611eb6838363ffffffff61230b16565b60008061218c613a63565b845184106121e1576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b600084905060008682815181106121f457fe5b016020015160019092019160f81c9050600061220e613a91565b60ff8316612242576122208985612e21565b9094509150600084612231846127b2565b919850965094506123049350505050565b60ff831660011415612269576122588985612e48565b909450905060008461223183612fb5565b60ff8316600214156122905761227f8985612e21565b909450915060008461223184613015565b600360ff8416108015906122a75750600c60ff8416105b156122e4576002198301606060006122c0838d89613093565b9098509250905080876122d284612830565b99509950995050505050505050612304565b8260ff166127100160006122f860006127b2565b91985096509450505050505b9250925092565b612321826040015161231c83612b82565b61314e565b82604001819052505050565b61233e826020015161231c83612b82565b82602001819052505050565b600061209361215061235d61213e612da4565b5160408501515114612cd7565b60006120938260a0015183612d8990919063ffffffff16565b600061238e82612d97565b61239a57506000611eba565b6123a382612b82565b60a084015250600192915050565b60006123c3838363ffffffff61232d16565b611eb6838363ffffffff61232d16565b60006123e5848363ffffffff61232d16565b6123f5848463ffffffff61232d16565b611e70848363ffffffff61232d16565b6000612417858363ffffffff61232d16565b612427858463ffffffff61232d16565b612437858563ffffffff61232d16565b611fae858363ffffffff61232d16565b60006123f5848463ffffffff61232d16565b600061246b858563ffffffff61232d16565b612437858463ffffffff61232d16565b600061248683612cb8565b1580612498575061249682613204565b155b156124a5575060006106eb565b6124ae82613213565b60ff168360000151106124c3575060006106eb565b611e7082604001518460000151815181106124da57fe5b60200260200101518561232d90919063ffffffff16565b60006124fc83613204565b158061250e575061250c84612cb8565b155b1561251b57506000611ca8565b61252483613213565b60ff1684600001511061253957506000611ca8565b81836040015185600001518151811061254e57fe5b6020908102919091010152611fae858463ffffffff61232d16565b600061257482613204565b61258057506000611eba565b611eb661258c83613213565b849060ff1663ffffffff612cc316565b50600190565b6000806125ad613ab8565b6125b684612b82565b51600193509150505b9250929050565b60008060008060008060006125da88613204565b6125ee5760009650945090925090506127a9565b61260f886040015160018151811061260257fe5b6020026020010151612cb8565b6126235760009650945090925090506127a9565b612637886040015160028151811061260257fe5b61264b5760009650945090925090506127a9565b61265f886040015160038151811061260257fe5b6126735760009650945090925090506127a9565b876040015160038151811061268457fe5b60200260200101516000015160001b925087604001516002815181106126a657fe5b602002602001015160000151915073__$9836fa7140e5a33041d4b827682e675a30$__624c28f66126d68a612b82565b6000015185858c604001516001815181106126ed57fe5b6020026020010151600001516040518563ffffffff1660e01b815260040180858152602001846affffffffffffffffffffff19166affffffffffffffffffffff19168152602001838152602001826001600160a01b03166001600160a01b0316815260200194505050505060206040518083038186803b15801561277057600080fd5b505af4158015612784573d6000803e3d6000fd5b505050506040513d602081101561279a57600080fd5b50516001975095509193509150505b92959194509250565b6127ba613a63565b60408051608080820183528482528251908101835260008082526020828101829052828501829052606083018290528084019290925283518181529182018452919283019161281f565b61280c613a63565b8152602001906001900390816128045790505b508152600060209091015292915050565b612838613a63565b6128428251613222565b612893576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b80516000906128ee84612b82565b511415612942576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b611e70848363ffffffff612d8916565b600260c090910152565b600160c090910152565b600060028260c00151141561297d575060006115ed565b60018260c001511415612992575060016115ed565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101206115ed565b600060c090910152565b600080612a10613a02565b612a18613a02565b600060c08201819052612a2b8787613229565b8452965090508015612a435793508492509050612304565b612a4d8787613229565b6020850152965090508015612a685793508492509050612304565b612a728787613229565b6040850152965090508015612a8d5793508492509050612304565b612a978787613229565b6060850152965090508015612ab25793508492509050612304565b612abc8787613229565b6080850152965090508015612ad75793508492509050612304565b612ae18787613229565b60a0850152965090508015612afc5793508492509050612304565b506000969495509392505050565b612b12613a02565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b6000806000612b798460ff16613267565b50949350505050565b612b8a613ab8565b6060820151600c60ff90911610612bdc576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff16612c09576040518060200160405280612c00846000015161370e565b905290506115ed565b606082015160ff1660011415612c50576040518060200160405280612c00846020015160000151856020015160400151866020015160600151876020015160200151613732565b606082015160ff1660021415612c7557506040805160208101909152815181526115ed565b600360ff16826060015160ff1610158015612c9957506060820151600c60ff909116105b15612cb6576040518060200160405280612c0084604001516137da565bfe5b6060015160ff161590565b61233e826020015161231c61213e846127b2565b612cdf613a63565b8115612cf657612cef60016127b2565b90506115ed565b612cef60006127b2565b612d08613a63565b816060015160ff1660021415612d4f5760405162461bcd60e51b8152600401808060200182810382526021815260200180613b396021913960400191505060405180910390fd5b606082015160ff16612d6557612cef60006127b2565b816060015160ff1660011415612d7f57612cef60016127b2565b612cef60036127b2565b61233e82602001518261314e565b6060015160ff1660011490565b612dac613a63565b604080516080808201835260008083528351918201845280825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612e11565b612dfe613a63565b815260200190600190039081612df65790505b5081526003602090910152905090565b6000808281612e36868363ffffffff61392616565b60209290920196919550909350505050565b6000612e52613a91565b60008390506000858281518110612e6557fe5b602001015160f81c60f81b60f81c905081806001019250506000868381518110612e8b57fe5b016020015160019384019360f89190911c915060009060ff84161415612f29576000612eb5613a63565b612ebf8a87612181565b90975090925090508115612f1a576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b612f2381612b82565b51925050505b6000612f3b898663ffffffff61392616565b90506020850194508360ff1660011415612f80576040805160808101825260ff9094168452602084019190915260019083015260608201529193509091506125bf9050565b6040805160808101825260ff909416845260208401919091526000908301819052606083015250919350909150509250929050565b612fbd613a63565b604080516080810182526000808252602080830186905283518281529081018452919283019190613004565b612ff1613a63565b815260200190600190039081612fe95790505b508152600160209091015292915050565b61301d613a63565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191613082565b61306f613a63565b8152602001906001900390816130675790505b508152600260209091015292915050565b60008060606000849050600060608860ff166040519080825280602002602001820160405280156130de57816020015b6130cb613a63565b8152602001906001900390816130c35790505b50905060005b8960ff168160ff161015613138576130fc8985612181565b8451859060ff861690811061310d57fe5b602090810291909101015294509250821561313057509094509092509050613145565b6001016130e4565b5060009550919350909150505b93509350939050565b613156613ab8565b6040805160028082526060828101909352816020015b613174613ab8565b81526020019060019003908161316c579050509050828160008151811061319757fe5b602002602001018190525083816001815181106131b057fe5b602002602001018190525060405180602001604052806131fa60405180604001604052806131e18860000151613015565b81526020016131f38960000151613015565b9052613942565b9052949350505050565b6000611eba82606001516139c1565b6000611eba82606001516139df565b6008101590565b600080613234613ab8565b836000613247878363ffffffff61392616565b604080516020808201909252918252600099930197509550909350505050565b600080600183141561327f5750600290506001613709565b60028314156132945750600290506001613709565b60038314156132a95750600290506001613709565b60048314156132be5750600290506001613709565b60058314156132d35750600290506001613709565b60068314156132e85750600290506001613709565b60078314156132fd5750600290506001613709565b60088314156133125750600390506001613709565b60098314156133275750600390506001613709565b600a83141561333c5750600290506001613709565b60108314156133515750600290506001613709565b60118314156133665750600290506001613709565b601283141561337b5750600290506001613709565b60138314156133905750600290506001613709565b60148314156133a55750600290506001613709565b60158314156133b957506001905080613709565b60168314156133ce5750600290506001613709565b60178314156133e35750600290506001613709565b60188314156133f85750600290506001613709565b601983141561340c57506001905080613709565b601a8314156134215750600290506001613709565b601b8314156134365750600290506001613709565b602083141561344a57506001905080613709565b602183141561345e57506001905080613709565b60308314156134735750600190506000613709565b60318314156134885750600090506001613709565b603283141561349d5750600090506001613709565b60338314156134b25750600190506000613709565b60348314156134c75750600190506000613709565b60358314156134dc5750600290506000613709565b60368314156134f15750600090506001613709565b60378314156135065750600090506001613709565b603883141561351b5750600190506000613709565b60398314156135305750600090506001613709565b603a8314156135455750600090506001613709565b603b83141561355957506000905080613709565b603c83141561356e5750600090506001613709565b603d8314156135835750600190506000613709565b60408314156135985750600190506002613709565b60418314156135ad5750600290506003613709565b60428314156135c25750600390506004613709565b60438314156135d657506002905080613709565b60448314156135ea57506003905080613709565b60508314156135ff5750600290506001613709565b60518314156136145750600390506001613709565b605283141561362857506001905080613709565b606083141561363c57506000905080613709565b60618314156136515750600190506000613709565b60708314156136665750600190506000613709565b607183141561367a57506001905080613709565b607283141561368f5750600090506001613709565b60738314156136a357506001905080613709565b60748314156136b757506000905080613709565b60758314156136cb57506000905080613709565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561378c575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120611ca8565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b600060088251111561382a576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613857578160200160208202803883390190505b50805190915060005b818110156138b357613870613ab8565b61388c86838151811061387f57fe5b6020026020010151612b82565b9050806000015184838151811061389f57fe5b602090810291909101015250600101613860565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156138fc5781810151838201526020016138e4565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6000816020018351101561393957600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b613965613a63565b81526020019060019003908161395d575050805190915060005b818110156139b75784816002811061399357fe5b60200201518382815181106139a457fe5b602090810291909101015260010161397f565b50611ca8826137da565b6000600c60ff8316108015611eba575050600360ff91909116101590565b60006139ea826139c1565b156139fa575060021981016115ed565b5060016115ed565b6040518060e00160405280613a15613ab8565b8152602001613a22613ab8565b8152602001613a2f613ab8565b8152602001613a3c613ab8565b8152602001613a49613ab8565b8152602001613a56613ab8565b8152602001600081525090565b604051806080016040528060008152602001613a7d613a91565b815260606020820152600060409091015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040805160208101909152600081529056fe4f6e6520737465702070726f6f66207769746820696e76616c6964207072657620737461746550726f6f6620686164206e6f6e206d61746368696e672073746172742073746174654c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f64654c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f73656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d657361676550726f6f66206f6620617578706f702068616420626164206175782076616c756543616e206f6e6c79206f6e6520737465702070726f6f6620666f6c6c6f77696e6720612073696e676c652073746570206368616c6c656e6765a265627a7a72315820a6ab08987996d06a38d5268f3937c5d888e3b20a750443acec774b7d1e13bd9464736f6c634300050c0032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

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

// ValidateProof is a free data retrieval call binding the contract method 0xc0fee45d.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCaller) ValidateProof(opts *bind.CallOpts, fields [7][32]byte, timeBounds [2]uint64, proof []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OneStepProof.contract.Call(opts, out, "validateProof", fields, timeBounds, proof)
	return *ret0, err
}

// ValidateProof is a free data retrieval call binding the contract method 0xc0fee45d.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, proof)
}

// ValidateProof is a free data retrieval call binding the contract method 0xc0fee45d.
//
// Solidity: function validateProof(bytes32[7] fields, uint64[2] timeBounds, bytes proof) constant returns(uint256)
func (_OneStepProof *OneStepProofCallerSession) ValidateProof(fields [7][32]byte, timeBounds [2]uint64, proof []byte) (*big.Int, error) {
	return _OneStepProof.Contract.ValidateProof(&_OneStepProof.CallOpts, fields, timeBounds, proof)
}
